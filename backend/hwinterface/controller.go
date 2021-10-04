package hwinterface

import (
	"bytes"
	"fmt"
	"image/jpeg"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/CapsLock-Studio/go-webpbin"
	"github.com/golang/glog"
	"github.com/idalmasso/stepperphotogopi/backend/hwinterface/drivers"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

type piController struct {
	degreesForPhoto        float64
	processing             bool
	motor                  *drivers.StepperMotorDriver
	gearTransmissionDriver *drivers.TransmissionFromStepperMotorDriver
	camera                 *drivers.CameraDriver
	buttonInput            *gpio.ButtonDriver
	ledOk                  *gpio.LedDriver
	mutex                  sync.RWMutex
	actualProcessName      string
	buttonPressFunc        func()
}

func (c *piController) SetDegreesMovement(degrees float64) error {
	if glog.V(3) {
		glog.Infoln("piController - SetDegreesMovement called w value", degrees)
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing {
		return ProcessingError{Operation: "Set #degrees for photo"}
	}
	c.degreesForPhoto = degrees
	return nil
}

//StartProcess actually starts the real process of making photo 360
func (c *piController) StartProcess(imagePath string) error {
	if glog.V(3) {
		glog.Infoln("piController - StartProcess called")
	}
	if !c.canSetStartProcess() {
		return ProcessingError{Operation: "Start Process"}
	}
	go func() {
		defer c.setProcessing(false)
		if c.degreesForPhoto == 0 {
			if glog.V(1) {
				glog.Errorln("piController - Set to 0 degrees")
			}
			return
		}
		c.gearTransmissionDriver.ResetActualAngle()
		newpath := imagePath
		if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
			if glog.V(1) {
				glog.Errorln("piController - StartProcess error on create public folder", err.Error())
			}
			return
		}
		t := time.Now()
		c.actualProcessName = fmt.Sprintf("%04d%02d%02d%02d%02d%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
		newpath = filepath.Join(newpath, c.actualProcessName)
		if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
			if glog.V(1) {
				glog.Errorln("piController - StartProcess error on create folder", newpath, err.Error())
			}
			return
		}

		for actualAngle, numPhoto := 0.0, 1; actualAngle < 360; actualAngle, numPhoto = actualAngle+c.degreesForPhoto, numPhoto+1 {
			if !c.isProcessing() {
				if glog.V(2) {
					glog.Warningln("piController - StartProcess interrupted")
				}
				return
			}
			if err := c.gearTransmissionDriver.GoToAngle(actualAngle); err != nil {
				if glog.V(1) {
					glog.Errorln("piController - CameraSnapshot gotoangle error", err.Error())
				}
				return
			}

			var bytes bytes.Buffer
			if err := c.CameraSnapshot(&bytes); err != nil {
				if glog.V(1) {
					glog.Errorln("piController - CameraSnapshot error", err.Error())
				}
				return
			}
			image, err := jpeg.Decode(&bytes)
			if err != nil {
				if glog.V(1) {
					glog.Errorln("piController - CameraSnapshot error", err.Error())
				}
				return
			}
			file, err := os.Create(filepath.Join(newpath, strconv.FormatInt(int64(numPhoto), 10)+".webp"))
			if err != nil {
				if glog.V(1) {
					glog.Errorln("piController - CameraSnapshot error on create photo file", newpath+"-"+strconv.FormatInt(int64(numPhoto), 10)+".webp", err.Error())
				}
				return
			}

			if err = webpbin.Encode(file, image); err != nil {
				if glog.V(1) {
					glog.Errorln("piController - CameraSnapshot webpbin ", err.Error())
				}
				file.Close()
				return
			}
			file.Close()
		}
	}()

	return nil
}

//StopProcess should stop the process at any time
func (c *piController) StopProcess() error {
	if glog.V(3) {
		glog.Infoln("piController - StopProcess called")
	}

	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing {
		c.ledOk.On()
		c.motor.Sleep()
		c.processing = false
		c.actualProcessName = ""
	}
	return nil
}
func (c *piController) MoveMotor() error {
	if glog.V(3) {
		glog.Infoln("piController - MoveMotor called")
	}
	if !c.canSetStartProcess() {
		return ProcessingError{Operation: "MoveMotor"}
	}
	if c.degreesForPhoto == 0 {
		return ProcessingError{Operation: "MoveMotor 0 degrees"}
	}
	c.moveMotorWork(int(c.degreesForPhoto / c.motor.DegreesPerStep()))
	return nil
}

func (c *piController) moveMotorWork(numSteps int) {
	if glog.V(3) {
		glog.Infoln("piController - moveMotorWork doing steps", numSteps)
	}
	for stepsDone := 0; stepsDone < numSteps; {
		if c.isProcessing() {
			c.motor.DoSteps(4)
			stepsDone += 4
		} else {
			return
		}
	}
	c.setProcessing(false)
}

// Return true if the machine is actually doing a work and so can be stopped but cannot start another one
func (c *piController) IsWorking() bool {
	return c.isProcessing()
}

func (c *piController) isProcessing() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.processing
}

func (c *piController) canSetStartProcess() bool {
	if glog.V(4) {
		glog.Infoln("piController -  canSetStartProcess canStartProcess")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing {
		return false
	} else {
		c.ledOk.Off()
		c.motor.Awake()

		c.processing = true

		if glog.V(4) {
			glog.Infoln("piController - canSetStartProcess start processing")
		}
		return true
	}
}

func (c *piController) GetActualProcessName() string { return c.actualProcessName }
func (c *piController) setProcessing(value bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing && !value {
		if glog.V(4) {
			glog.Infoln("piController - setProcessing stop processing")
		}
		c.actualProcessName = ""
	} else if !c.processing && value {
		if glog.V(4) {
			glog.Infoln("piController - setProcessing start processing")
		}
	}
	c.processing = value
	if c.processing {
		c.ledOk.Off()
		c.motor.Awake()
	} else {
		c.ledOk.On()
		c.motor.Sleep()
	}
}

//Writes a snapshot into the writer passed
func (c *piController) CameraSnapshot(w io.Writer) (err error) {
	if glog.V(3) {
		glog.Infoln("piController - piController.CameraSnapshot - start")
	}
	return c.camera.DoPhoto(w)
}

func (c *piController) SetMotorDegreePerStep(degrees float64) {
	if glog.V(3) {
		glog.Infoln("piController - piController.SetMotorDegreePerStep - start")
	}
	if !c.isProcessing() {
		c.motor.SetDegreesPerStep(degrees)
	}
}

func (c *piController) SetWaitForStep(waitTimeMillis int) {
	if glog.V(3) {
		glog.Infoln("piController - piController.SetWaitForStep - start")
	}
	if !c.isProcessing() {
		c.motor.SetWaitTimeBetweenSteps(waitTimeMillis)
	}
}
func (c *piController) SetGearRatio(ratio float64) {
	if glog.V(3) {
		glog.Infoln("piController - piController.SetGearRatio - start")
	}
	if !c.isProcessing() {
		c.gearTransmissionDriver.SetRatio(ratio)
	}
}
func (c *piController) GetMotorDegreePerStep() float64 {
	if glog.V(3) {
		glog.Infoln("piController - piController.GetMotorDegreePerStep - start")
	}
	return c.motor.DegreesPerStep()
}

func (c *piController) GetWaitForStep() int {
	if glog.V(3) {
		glog.Infoln("piController - piController.GetWaitForStep - start")
	}
	return c.motor.WaitTimeBetweenSteps()
}
func (c *piController) GetGearRatio() float64 {
	if glog.V(3) {
		glog.Infoln("piController - piController.GetGearRatio - start")
	}
	return c.gearTransmissionDriver.GetRatio()
}

func (c *piController) SetCameraWidth(width int)           { c.camera.SetWidth(width) }
func (c *piController) SetCameraHeight(height int)         { c.camera.SetHeight(height) }
func (c *piController) SetCameraContrast(contrast int)     { c.camera.SetContrast(contrast) }
func (c *piController) SetCameraSharpness(sharpness int)   { c.camera.SetSharpness(sharpness) }
func (c *piController) SetCameraBrightness(brightness int) { c.camera.SetBrightness(brightness) }
func (c *piController) SetCameraSaturation(saturation int) { c.camera.SetSaturation(saturation) }
func (c *piController) SetCameraAWB(awbMode string)        { c.camera.SetAWBMode(awbMode) }
func (c *piController) SetOnButtonPress(callback func()) {
	c.buttonPressFunc = callback
}
func (c *piController) buttonPressed(interface{}) {
	c.buttonPressFunc()
}
func NewController() *piController {

	r := raspi.NewAdaptor()
	r.Connect()

	motor := drivers.NewStepperMotorDriver(r, "38", "40", "36", 1, 1)
	motor.Start()
	camera := drivers.NewCameraDriver()
	camera.Start()
	//TODO: Update the ratio here!
	buttonInput := gpio.NewButtonDriver(r, "15", time.Duration(10*time.Millisecond))

	buttonInput.Start()

	gearTransmissionDriver := drivers.NewTransmissionStepperMotorDriver(motor, 1)
	gearTransmissionDriver.Start()
	ledOk := gpio.NewLedDriver(r, "13")
	ledOk.Start()
	ledOk.On()
	pi := piController{motor: motor, camera: camera, gearTransmissionDriver: gearTransmissionDriver, buttonInput: buttonInput, ledOk: ledOk}
	buttonInput.On(gpio.ButtonRelease, pi.buttonPressed)
	return &pi
}

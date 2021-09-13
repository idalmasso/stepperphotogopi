package hwinterface

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/idalmasso/stepperphotogopi/backend/hwinterface/drivers"
	"gobot.io/x/gobot/platforms/raspi"
)

type piController struct {
	degreesForPhoto        float64
	processing             bool
	motor                  *drivers.StepperMotorDriver
	gearTransmissionDriver *drivers.TransmissionFromStepperMotorDriver
	camera                 *drivers.CameraDriver
	mutex                  sync.RWMutex
	actualProcessName      string
}

func (c *piController) SetDegreesMovement(degrees float64) error {
	if glog.V(3) {
		glog.Infoln("piController - SetDegreesMovement called")
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
					glog.Errorln("piController - Gotoangle error on create public folder", err.Error())
				}
				return
			}
			file, err := os.Create(filepath.Join(newpath, strconv.FormatInt(int64(numPhoto), 10)+".jpg"))
			if err != nil {
				if glog.V(1) {
					glog.Errorln("piController - Gotoangle error on create photo file", newpath+"-"+strconv.FormatInt(int64(numPhoto), 10)+".jpg", err.Error())
				}
				return
			}
			if err := c.CameraSnapshot(file); err != nil {
				if glog.V(1) {
					glog.Errorln("piController - CameraSnapshot error", err.Error())
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

	//RAMP here! and then, each time
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

func NewController() piController {

	r := raspi.NewAdaptor()
	r.Connect()

	motor := drivers.NewStepperMotorDriver(r, "40", "39", 1, 1)
	motor.Start()
	camera := drivers.NewCameraDriver()
	camera.Start()
	//TODO: Update the ratio here!

	gearTransmissionDriver := drivers.NewTransmissionStepperMotorDriver(motor, 1)
	gearTransmissionDriver.Start()
	return piController{motor: motor, camera: camera, gearTransmissionDriver: gearTransmissionDriver}
}

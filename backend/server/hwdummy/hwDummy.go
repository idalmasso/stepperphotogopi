package hwdummy

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/golang/glog"
)

type dummyController struct {
	degreesForPhoto   float64
	processing        bool
	mutex             sync.RWMutex
	actualProcessName string
}

func (c *dummyController) SetDegreesMovement(degrees float64) error {
	if glog.V(3) {
		glog.Infoln("dummyController - SetDegreesMovement called")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing {
		return ProcessingError{Operation: "Set #degrees for photo"}
	}
	c.degreesForPhoto = degrees
	return nil
}

//StopProcess should stop the process at any time
func (c *dummyController) StopProcess() error {
	if glog.V(3) {
		glog.Infoln("dummyController - StopProcess called")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing {
		c.processing = false
		c.actualProcessName = ""
	}
	return nil
}
func (c *dummyController) MoveMotor() error {
	if glog.V(3) {
		glog.Infoln("dummyController - MoveMotor called")
	}
	if !c.canSetStartProcess() {
		return ProcessingError{Operation: "MoveMotor"}
	}
	if c.degreesForPhoto == 0 {
		return ProcessingError{Operation: "MoveMotor 0 degrees"}
	}
	c.moveMotorWork(int(c.degreesForPhoto / 1.8))
	return nil
}

//StartProcess actually starts the real process of making photo 360
func (c *dummyController) StartProcess(imagePath string) error {
	if glog.V(3) {
		glog.Infoln("dummyController - StartProcess called")
	}
	if !c.canSetStartProcess() {
		return ProcessingError{Operation: "Start Process"}
	}
	go func() {
		defer c.setProcessing(false)
		if c.degreesForPhoto == 0 {
			if glog.V(1) {
				glog.Errorln("dummyController - Set to 0 degrees")
			}
			return
		}

		newpath := imagePath
		if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
			if glog.V(1) {
				glog.Errorln("dummyController - StartProcess error on create public folder", err.Error())
			}
			return
		}
		t := time.Now()
		c.actualProcessName = fmt.Sprintf("%04d%02d%02d%02d%02d%02d", t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), t.Second())
		newpath = filepath.Join(newpath, c.actualProcessName)
		if err := os.MkdirAll(newpath, os.ModePerm); err != nil {
			if glog.V(1) {
				glog.Errorln("dummyController - StartProcess error on create folder", newpath, err.Error())
			}
			return
		}

		for actualAngle, numPhoto := 0.0, 1; actualAngle < 360; actualAngle, numPhoto = actualAngle+c.degreesForPhoto, numPhoto+1 {
			if !c.isProcessing() {
				if glog.V(2) {
					glog.Warningln("dummyController - StartProcess interrupted")
				}
				return
			}
			time.Sleep(500 * time.Millisecond)
			file, err := os.Create(filepath.Join(newpath, strconv.FormatInt(int64(numPhoto), 10)+".jpg"))
			if err != nil {
				if glog.V(1) {
					glog.Errorln("dummyController - Gotoangle error on create photo file", newpath+"-"+strconv.FormatInt(int64(numPhoto), 10)+".jpg", err.Error())
				}
				return
			}
			if err := c.CameraSnapshot(file); err != nil {
				if glog.V(1) {
					glog.Errorln("dummyController - CameraSnapshot error", err.Error())
				}
				file.Close()
				return
			}
			file.Close()
		}
	}()

	return nil
}

func (c *dummyController) moveMotorWork(numSteps int) {
	if glog.V(3) {
		glog.Infoln("dummyController - moveMotorWork doing steps", numSteps)
	}

	//RAMP here! and then, each time
	for stepsDone := 0; stepsDone < numSteps; {
		if c.isProcessing() {
			time.Sleep(time.Duration(4) * time.Second)
			stepsDone += 4
		} else {
			return
		}
	}
	c.setProcessing(false)
}

// Return true if the machine is actually doing a work and so can be stopped but cannot start another one
func (c *dummyController) IsWorking() bool {
	return c.isProcessing()
}

func (c *dummyController) isProcessing() bool {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.processing
}

func (c *dummyController) canSetStartProcess() bool {
	if glog.V(3) {
		glog.Infoln("dummyController - canSetStartProcess")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing {
		return false
	} else {
		if glog.V(3) {
			glog.Infoln("dummyController - start processing")
		}
		c.processing = true
		return true
	}
}
func (c *dummyController) setProcessing(value bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing && !value {
		if glog.V(3) {
			glog.Infoln("dummyController - stop processing")
		}
		c.actualProcessName = ""
	} else if !c.processing && value {
		if glog.V(3) {
			glog.Infoln("dummyController - start processing")
		}
	}
	c.processing = value
}
func (c *dummyController) SetMotorDegreePerStep(degrees float64) {
	if glog.V(3) {
		glog.Infoln("dummyController - dummyController.SetMotorDegreePerStep - start")
	}

}

func (c *dummyController) SetWaitForStep(waitTimeMillis int) {
	if glog.V(3) {
		glog.Infoln("dummyController - dummyController.SetWaitForStep - start")
	}

}
func (c *dummyController) SetGearRatio(ratio float64) {
	if glog.V(3) {
		glog.Infoln("dummyController - dummyController.SetGearRatio - start")
	}

}

func (c *dummyController) GetMotorDegreePerStep() float64 {
	if glog.V(3) {
		glog.Infoln("dummyController - dummyController.GetMotorDegreePerStep - start")
	}
	return 0
}

func (c *dummyController) GetWaitForStep() int {
	if glog.V(3) {
		glog.Infoln("dummyController - dummyController.GetWaitForStep - start")
	}
	return 0
}
func (c *dummyController) GetGearRatio() float64 {
	if glog.V(3) {
		glog.Infoln("dummyController - dummyController.GetGearRatio - start")
	}
	return 0
}

func (c *dummyController) GetActualProcessName() string { return c.actualProcessName }

//Writes a snapshot into the writer passed
func (c *dummyController) CameraSnapshot(w io.Writer) (err error) {
	m := image.NewRGBA(image.Rect(0, 0, 240, 240))
	blue := color.RGBA{uint8(rand.Intn(255)), uint8(rand.Intn(255)), uint8(rand.Intn(255)), 255}
	draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.Point{1, 1}, draw.Src)
	buffer := new(bytes.Buffer)
	if err := jpeg.Encode(buffer, m, nil); err != nil {
		if glog.V(1) {
			glog.Errorln("unable to encode image.")
		}
	}
	if _, err = w.Write(buffer.Bytes()); err != nil {
		if glog.V(1) {
			glog.Errorln("unable to write image.")
		}
	}
	return
}
func NewController() dummyController {
	return dummyController{}
}

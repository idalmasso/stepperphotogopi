package hwinterface

import (
	"io"
	"sync"
	"time"

	"github.com/golang/glog"
	"github.com/idalmasso/stepperphotogopi/backend/hwinterface/drivers"
	"gobot.io/x/gobot/platforms/raspi"
)

type piController struct {
	degreesForPhoto float64
	processing      bool
	motor           *drivers.StepperMotorDriver
	camera					*drivers.CameraDriver
	mutex           sync.RWMutex
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
func (c *piController) StartProcess() error {
	if glog.V(3) {
		glog.Infoln("piController - StartProcess called")
	}
	if !c.canSetStartProcess() {
		return ProcessingError{Operation: "Start Process"}
	}
	c.processWork(10000)
	c.setProcessing(true)

	return nil
}
func (c *piController) StopProcess() error {
	if glog.V(3) {
		glog.Infoln("piController - StopProcess called")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing {
		c.processing = false
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
func (c *piController) processWork(numSteps int) {
	if glog.V(3) {
		glog.Infoln("piController - processWork")
	}
	
	time.Sleep(time.Millisecond * 100)

}
func (c *piController) moveMotorWork(numSteps int) {
	if glog.V(3) {
		glog.Infoln("piController - moveMotorWork doing steps", numSteps)
	}
	
	//RAMP here! and then, each time
	for stepsDone:=0;stepsDone<numSteps;{
		if c.isProcessing(){
			c.motor.DoSteps(4)
			stepsDone+=4
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
func (c *piController) setProcessing(value bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing && !value {
		if glog.V(4) {
			glog.Infoln("piController - setProcessing stop processing")
		}
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
func NewController() piController {

	r := raspi.NewAdaptor()
	r.Connect()

	motor := drivers.NewStepperMotorDriver(r, "40", "39")
	motor.Start()
	camera := drivers.NewCameraDriver()
	return piController{ motor: motor, camera: camera}
}


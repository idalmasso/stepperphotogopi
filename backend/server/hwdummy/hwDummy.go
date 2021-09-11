package hwdummy

import (
	"sync"
	"time"

	"github.com/golang/glog"
)

type dummyController struct {
	degreesForPhoto float64
	processing      bool
	mutex           sync.RWMutex
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
func (c *dummyController) StartProcess() error {
	if glog.V(3) {
		glog.Infoln("dummyController - StartProcess called")
	}
	if !c.canSetStartProcess() {
		return ProcessingError{Operation: "Start Process"}
	}
	c.processWork(10000)

	return nil
}
func (c *dummyController) StopProcess() error {
	if glog.V(3) {
		glog.Infoln("dummyController - StopProcess called")
	}
	c.setProcessing(false)
	return nil
}
func (c *dummyController) MoveMotor() error {
	if glog.V(3) {
		glog.Infoln("piController - MoveMotor called")
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
func (c *dummyController) processWork(numSteps int) {

	if glog.V(3) {
		glog.Infoln("dummyController - processWork")
	}
	c.moveMotorWork(numSteps)

}
func (c *dummyController) moveMotorWork(numSteps int) {
		if glog.V(3) {
		glog.Infoln("piController - moveMotorWork doing steps", numSteps)
	}
	
	//RAMP here! and then, each time
	for stepsDone:=0;stepsDone<numSteps;{
		if c.isProcessing(){
			time.Sleep(time.Duration(4)*time.Second)
			stepsDone+=4
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
	} else if !c.processing && value {
		if glog.V(3) {
			glog.Infoln("dummyController - start processing")
		}
	}
	c.processing = value
}
func NewController() dummyController {
	return dummyController{}
}

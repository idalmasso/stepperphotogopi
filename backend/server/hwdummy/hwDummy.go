package hwdummy

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/glog"
)

type dummyController struct {
	degreesForPhoto float64
	processing      bool
	endWork         chan bool
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
	go func() {
		for {
			select {
			case <-c.endWork:
				return
			default:
				c.processWork(10000)
			}
		}
	}()
	c.setProcessing(true)

	return nil
}
func (c *dummyController) StopProcess() error {
	if glog.V(3) {
		glog.Infoln("dummyController - StopProcess called")
	}
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.processing {
		c.processing = false
		c.endWork <- true
	}
	return nil
}
func (c *dummyController) MoveMotor() error {
	if glog.V(3) {
		glog.Infoln("dummyController - MoveMotor called")
	}
	if c.isProcessing() {
		return fmt.Errorf("Cannot move MoveMotor while already processing")
	}

	c.setProcessing(true)
	c.moveMotorWork(int(c.degreesForPhoto / 1.8))
	c.setProcessing(false)
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
		glog.Infoln("dummyController - moveMotorWork for steps", numSteps)
	}
	time.Sleep(time.Duration(numSteps) * time.Second)
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
	return dummyController{endWork: make(chan bool)}
}

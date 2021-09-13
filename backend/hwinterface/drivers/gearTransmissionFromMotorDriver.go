package drivers

import "gobot.io/x/gobot"

//TransmissionFromStepperMotorDriver is a utility class wrapping the normal MotorDriver. It just multiply the steps
type TransmissionFromStepperMotorDriver struct {
	motorDriver       *StepperMotorDriver
	transmissionRatio float64
	actualAngle       float64
	gobot.Commander
}

// NewTransmissionStepperMotorDriver return a new TransmissionStepperMotorDriver given a Motor and a transmissionRatio.

func NewTransmissionStepperMotorDriver(motor *StepperMotorDriver, ratio float64) *TransmissionFromStepperMotorDriver {
	l := &TransmissionFromStepperMotorDriver{
		motorDriver:       motor,
		transmissionRatio: ratio,
		actualAngle:       0,
		Commander:         gobot.NewCommander(),
	}
	return l
}
func (d *TransmissionFromStepperMotorDriver) ResetActualAngle() {
	d.actualAngle = 0
}
func (d *TransmissionFromStepperMotorDriver) SetRatio(newRatio float64) {
	d.transmissionRatio = newRatio
}
func (d *TransmissionFromStepperMotorDriver) GetRatio() float64 {
	return d.transmissionRatio
}
func (d *TransmissionFromStepperMotorDriver) GoToAngle(angle float64) (err error) {
	if angle < d.actualAngle && d.motorDriver.forward || angle > d.actualAngle && !d.motorDriver.forward {
		if err = d.motorDriver.SwapDirection(); err != nil {
			return err
		}
	}
	var toDoAngle float64
	if d.motorDriver.forward {
		toDoAngle = angle - d.actualAngle
	} else {
		toDoAngle = d.actualAngle - angle
	}
	toDoSteps := (toDoAngle / d.transmissionRatio) / d.motorDriver.DegreesPerStep()
	steps, err := d.motorDriver.DoSteps(int(toDoSteps))
	if d.motorDriver.forward {
		d.actualAngle += float64(steps) * d.transmissionRatio * d.motorDriver.DegreesPerStep()
	} else {
		d.actualAngle -= float64(steps) * d.transmissionRatio * d.motorDriver.DegreesPerStep()
	}
	return
}

// Start implements the Driver interface
func (l *TransmissionFromStepperMotorDriver) Start() (err error) { return }

// Halt implements the Driver interface
func (l *TransmissionFromStepperMotorDriver) Halt() (err error) { return }

// Name returns the TransmissionFromStepperMotorDriver name (=to the motorDriver)
func (l *TransmissionFromStepperMotorDriver) Name() string { return l.motorDriver.Name() }

// SetName sets the TransmissionFromStepperMotorDriver name (=to the motorDriver)
func (l *TransmissionFromStepperMotorDriver) SetName(n string) { l.motorDriver.SetName(n) }

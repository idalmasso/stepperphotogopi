package drivers

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

// StepperMotorDriver represents a Stepper motor. TBD> Use with PWM
type StepperMotorDriver struct {
	pin            string
	directionPin   string
	name           string
	connection     gpio.DigitalWriter
	steps          int
	forward        bool
	millisWait     int
	degreesPerStep float64
	gobot.Commander
}

// NewLedDriver return a new LedDriver given a DigitalWriter and pin.
//
// Adds the following API Commands:
//	"Brightness" - See LedDriver.Brightness
//	"Toggle" - See LedDriver.Toggle
//	"On" - See LedDriver.On
//	"Off" - See LedDriver.Off
func NewStepperMotorDriver(a gpio.DigitalWriter, pin, directionPin string) *StepperMotorDriver {
	l := &StepperMotorDriver{
		name:           gobot.DefaultName("STEPPERMOTOR"),
		pin:            pin,
		directionPin:   directionPin,
		connection:     a,
		steps:          0,
		forward:        true,
		millisWait:     20,
		degreesPerStep: 1.8,
		Commander:      gobot.NewCommander(),
	}

	l.AddCommand("DoSteps", func(params map[string]interface{}) interface{} {
		numSteps := params["numSteps"].(int)
		return l.DoSteps(numSteps)
	})

	/*
		l.AddCommand("Toggle", func(params map[string]interface{}) interface{} {
			return l.Toggle()
		})

		l.AddCommand("On", func(params map[string]interface{}) interface{} {
			return l.On()
		})

		l.AddCommand("Off", func(params map[string]interface{}) interface{} {
			return l.Off()
		})
	*/
	return l
}

// Start implements the Driver interface
func (l *StepperMotorDriver) Start() (err error) { return }

// Halt implements the Driver interface
func (l *StepperMotorDriver) Halt() (err error) { return }

// Name returns the StepperMotorDriver name
func (l *StepperMotorDriver) Name() string { return l.name }

// SetName sets the StepperMotorDriver name
func (l *StepperMotorDriver) SetName(n string) { l.name = n }

// Pin returns the StepperMotorDriver pin name
func (l *StepperMotorDriver) Pin() string { return l.pin }

// DirectionPin returns the StepperMotorDriver direction pinname
func (l *StepperMotorDriver) DirectionPin() string { return l.directionPin }

// Connection returns the StepperMotorDriver Connection
func (l *StepperMotorDriver) Connection() gobot.Connection {
	return l.connection.(gobot.Connection)
}

// NumSteps return the number of steps done
func (l *StepperMotorDriver) NumSteps() int {
	return l.steps
}

// DegreesPerStep return the number of degrees made with 1 step done
func (l *StepperMotorDriver) DegreesPerStep() float64 {
	return l.degreesPerStep
}

// DoSteps does the actual number of requested steps in the set direction
func (l *StepperMotorDriver) DoSteps(numSteps int) (err error) {
	for l.steps = 0; l.steps < numSteps; l.steps++ {
		if err = l.doSingleStep(); err != nil {
			return
		}
	}
	return
}

// On sets the led to a high state.
func (l *StepperMotorDriver) doSingleStep() (err error) {
	if err = l.connection.DigitalWrite(l.Pin(), 1); err != nil {
		return
	}

	time.Sleep(time.Duration(l.millisWait) * time.Millisecond)
	if err = l.connection.DigitalWrite(l.Pin(), 0); err != nil {
		return
	}
	time.Sleep(time.Duration(l.millisWait) * time.Millisecond)
	return
}

// SwapDirection change direction of the motor
func (l *StepperMotorDriver) SwapDirection() (err error) {
	if l.forward {
		return l.SetBackward()
	} else {
		return l.SetForward()
	}
}

//SetBackward set the motor to go backward
func (l *StepperMotorDriver) SetBackward() (err error) {
	err = l.connection.DigitalWrite(l.DirectionPin(), 0)
	return
}

//SetForward set the motor to go forward
func (l *StepperMotorDriver) SetForward() (err error) {
	err = l.connection.DigitalWrite(l.DirectionPin(), 1)
	return
}

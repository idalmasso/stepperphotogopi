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
	notSleepPin    string
	name           string
	connection     gpio.DigitalWriter
	steps          int
	forward        bool
	millisWait     int
	degreesPerStep float64
	sleeping       bool
	gobot.Commander
}

// NewStepperMotorDriver return a new StepperMotorDriver given a DigitalWriter and pin and a directionpin.

func NewStepperMotorDriver(a gpio.DigitalWriter, pin, directionPin string, notSleepPin string, degreesPerStep float64, millisWait int) *StepperMotorDriver {
	l := &StepperMotorDriver{
		name:           gobot.DefaultName("STEPPERMOTOR"),
		pin:            pin,
		directionPin:   directionPin,
		notSleepPin:    notSleepPin,
		connection:     a,
		steps:          0,
		forward:        true,
		sleeping:       true,
		millisWait:     millisWait,
		degreesPerStep: degreesPerStep,
		Commander:      gobot.NewCommander(),
	}

	l.AddCommand("DoSteps", func(params map[string]interface{}) interface{} {
		numSteps := params["numSteps"].(int)
		_, err := l.DoSteps(numSteps)
		return err
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
func (l *StepperMotorDriver) Start() (err error) {
	err = l.Sleep()
	return
}

// Halt implements the Driver interface
func (l *StepperMotorDriver) Halt() (err error) {
	err = l.Sleep()
	return
}

// Name returns the StepperMotorDriver name
func (l *StepperMotorDriver) Name() string { return l.name }

// SetName sets the StepperMotorDriver name
func (l *StepperMotorDriver) SetName(n string) { l.name = n }

// Pin returns the StepperMotorDriver pin name
func (l *StepperMotorDriver) Pin() string { return l.pin }

// NotSleepPin returns the StepperMotorDriver sleep pinname  (enabled at 0 logic value)
func (l *StepperMotorDriver) NotSleepPin() string { return l.notSleepPin }

// DirectionPin returns the StepperMotorDriver direction pinname
func (l *StepperMotorDriver) DirectionPin() string { return l.directionPin }

// Connection returns the StepperMotorDriver Connection
func (l *StepperMotorDriver) Connection() gobot.Connection {
	return l.connection.(gobot.Connection)
}

//Sleep puts the motor in sleep state
func (l *StepperMotorDriver) Sleep() (err error) {
	err = l.connection.DigitalWrite(l.NotSleepPin(), 0)
	l.sleeping = true
	return
}

//Awake remove the motor from sleep state
func (l *StepperMotorDriver) Awake() (err error) {
	err = l.connection.DigitalWrite(l.NotSleepPin(), 1)
	time.Sleep(100 * time.Millisecond)
	l.sleeping = false
	return
}

func (l *StepperMotorDriver) IsSleeping() bool {
	return l.sleeping
}

// NumSteps return the number of steps done
func (l *StepperMotorDriver) NumSteps() int {
	return l.steps
}

// DegreesPerStep return the number of degrees made with 1 step done
func (l *StepperMotorDriver) DegreesPerStep() float64 {
	return l.degreesPerStep
}

// SetDegreesPerStep sets the number of degrees done with 1 step
func (l *StepperMotorDriver) SetDegreesPerStep(degrees float64) {
	l.degreesPerStep = degrees
}

// WaitTimeBetweenSteps return the number of millis to wait between steps
func (l *StepperMotorDriver) WaitTimeBetweenSteps() int {
	return l.millisWait
}

// SetWaitTimeBetweenSteps sets   the number of millis to wait between steps
func (l *StepperMotorDriver) SetWaitTimeBetweenSteps(wait int) {
	l.millisWait = wait
}

// DoSteps does the actual number of requested steps in the set direction
func (l *StepperMotorDriver) DoSteps(numSteps int) (stepsDone int, err error) {
	stepsDone = 0
	for l.steps = 0; l.steps < numSteps; l.steps++ {
		if err = l.doSingleStep(); err != nil {
			return
		}
		stepsDone++
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

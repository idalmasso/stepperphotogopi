package drivers

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/dhowden/raspicam"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

// CameraDriver represents a Camera
type CameraDriver struct {
	name string
	connection     	gpio.DigitalWriter
	secondsWait    	int
	still 					*raspicam.Still
	gobot.Commander
}

// NewCameraDriver return a new CameraDriver.
//
func NewCameraDriver() *CameraDriver {
	l := &CameraDriver{
		name:           gobot.DefaultName("CAMERA"),
		secondsWait: 		1,
		Commander:      gobot.NewCommander(),
		
		
	}
	
	return l
}

// Start implements the Driver interface
func (l *CameraDriver) Start() (err error) { 
	l.still = raspicam.NewStill()
	l.still.Timeout = 1 * time.Second
	
	l.still.Width=2000
	l.still.Height=1500
	l.still.Camera.Brightness = 50
	l.still.Camera.Contrast = 0
	l.still.Camera.Sharpness = 0
	return 
}

// Halt implements the Driver interface
func (l *CameraDriver) Halt() (err error) { return }

// Name returns the CameraDriver name
func (l *CameraDriver) Name() string { return l.name }

// SetName sets the CameraDriver name
func (l *CameraDriver) SetName(n string) { l.name = n }


// Connection returns the CameraDriver Connection
func (l *CameraDriver) Connection() gobot.Connection {
	return l.connection.(gobot.Connection)
}

// SecondsWait return the number of Seconds to wait before photo
func (l *CameraDriver) SecondsWait() int {
	return l.secondsWait
}

// DoSteps does the actual number of requested steps in the set direction
func (l *CameraDriver) DoPhoto(w io.Writer) (err error) {
	errCh := make(chan error)
	go func() {
		for x := range errCh {
			fmt.Fprintf(os.Stderr, "%v\n", x)
		}
	}()
	log.Println("Capturing image...")
	raspicam.Capture(l.still, w, errCh)
	return
}


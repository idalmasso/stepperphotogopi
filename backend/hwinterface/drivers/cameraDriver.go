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
	name        string
	connection  gpio.DigitalWriter
	secondsWait int
	still       *raspicam.Still
	gobot.Commander
}

// NewCameraDriver return a new CameraDriver.
//
func NewCameraDriver() *CameraDriver {
	l := &CameraDriver{
		name:        gobot.DefaultName("CAMERA"),
		secondsWait: 1,
		Commander:   gobot.NewCommander(),
		still:       raspicam.NewStill(),
	}
	l.still.Timeout = 1 * time.Second
	l.still.Width = 2000
	l.still.Height = 1500
	l.still.Camera.Brightness = 50
	l.still.Camera.Contrast = 0
	l.still.Camera.Sharpness = 0
	l.still.Camera.Saturation = 0
	l.still.Camera.AWBMode = raspicam.AWBAuto

	return l
}

// Start implements the Driver interface
func (l *CameraDriver) Start() (err error) { return }

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

func (l *CameraDriver) SetWidth(width int) {
	l.still.Width = width
}
func (l *CameraDriver) SetHeight(height int) {
	l.still.Height = height
}
func (l *CameraDriver) SetBrightness(brightness int) {
	l.still.Camera.Brightness = brightness
}
func (l *CameraDriver) SetContrast(contrast int) {
	l.still.Camera.Contrast = contrast
}
func (l *CameraDriver) SetSharpness(sharpness int) {
	l.still.Camera.Sharpness = sharpness
}
func (l *CameraDriver) SetSaturation(saturation int) {
	l.still.Camera.Saturation = saturation
}
func (l *CameraDriver) SetAWBMode(awbMode string) {
	switch awbMode {
	case "off":
		l.still.Camera.AWBMode = raspicam.AWBOff
	case "auto":
		l.still.Camera.AWBMode = raspicam.AWBAuto
	case "sun":
		l.still.Camera.AWBMode = raspicam.AWBSunlight
	case "cloud":
		l.still.Camera.AWBMode = raspicam.AWBCloudy
	case "shade":
		l.still.Camera.AWBMode = raspicam.AWBShade
	case "tungsten":
		l.still.Camera.AWBMode = raspicam.AWBTungsten
	case "fluorescent":
		l.still.Camera.AWBMode = raspicam.AWBFluorescent
	case "incandescent":
		l.still.Camera.AWBMode = raspicam.AWBIncandescent
	case "flash":
		l.still.Camera.AWBMode = raspicam.AWBFlash
	case "horizon":
		l.still.Camera.AWBMode = raspicam.AWBHorizon
	default:
		l.still.Camera.AWBMode = raspicam.AWBAuto
	}
}

// DoPhoto does a photo and write in the passed writer
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

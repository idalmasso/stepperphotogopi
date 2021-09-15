package server

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang/glog"
	"github.com/idalmasso/stepperphotogopi/backend/config"
)

type controllerMachine interface {
	SetDegreesMovement(degrees float64) error
	StartProcess(folder string) error
	StopProcess() error
	MoveMotor() error
	IsWorking() bool
	CameraSnapshot(w io.Writer) error
	GetActualProcessName() string
	SetGearRatio(ratio float64)
	SetWaitForStep(waitTimeMillis int)
	SetMotorDegreePerStep(degrees float64)
	GetGearRatio() float64
	GetWaitForStep() int
	GetMotorDegreePerStep() float64
	SetCameraWidth(int)
	SetCameraHeight(int)
	SetCameraContrast(int)
	SetCameraSharpness(int)
	SetCameraBrightness(int)
}

//PiServer
type MachineServer struct {
	configuration *config.Config
	initialized   bool
	Router        chi.Router
	machine       controllerMachine
}

//ListenAndServe is the main server procedure that only wraps http.ListenAndServe
func (s *MachineServer) ListenAndServe() {
	if glog.V(3) {
		glog.Infoln("MachineServer -  MachineServer.ListenAndServe start")
	}
	if !s.initialized {
		panic("Server not initialized")
	}
	if glog.V(3) {
		glog.Infoln("MachineServer -  MachineServer.starting on port", s.configuration.Server.Port)
	}
	if err := http.ListenAndServe(":"+s.configuration.Server.Port, s.Router); err != nil {
		panic("Cannot listen on server: " + err.Error())
	}
}

//Init initialize the server router and set the controllerMachine needed to do the work
func (s *MachineServer) Init(machine controllerMachine) {
	if glog.V(3) {
		glog.Infoln("MachineServer -  MachineServer.Init start")
	}
	s.configuration = &config.Config{}
	if err := s.configuration.ReadFromFile("configuration.yaml"); err != nil {
		panic("cannot read configuration file")
	}
	s.machine = machine
	s.machine.SetMotorDegreePerStep(s.configuration.Hardware.MotorDegreePerStep)
	s.machine.SetGearRatio(s.configuration.Hardware.GearRatio)
	s.machine.SetWaitForStep(s.configuration.Hardware.WaitForStep)
	s.machine.SetCameraHeight(s.configuration.Hardware.Camera.Height)
	s.machine.SetCameraWidth(s.configuration.Hardware.Camera.Width)
	s.machine.SetCameraBrightness(s.configuration.Hardware.Camera.Brightness)
	s.machine.SetCameraContrast(s.configuration.Hardware.Camera.Contrast)
	s.machine.SetCameraSharpness(s.configuration.Hardware.Camera.Sharpness)

	s.Router = chi.NewRouter()
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))

	FileServer(s.Router.(*chi.Mux), s.configuration.Server.DistributionDirectory)
	FileServerImages(s.Router.(*chi.Mux), s.configuration.Server.PhotoDirectory)
	s.Router.Route("/api", func(router chi.Router) {
		router.Route("/processes", func(processRouter chi.Router) {
			processRouter.Get("/", s.getListProcessDone)
			processRouter.Post("/", s.startProcess)
			processRouter.Delete("/{process}", s.deleteProcessDone)
			processRouter.Get("/{process}", s.getZipProcess)
		})
		router.Get("/get-snapshot", s.cameraSnapshot)
		router.Get("/machine-status", s.getMachineStatus)

		router.Post("/move-motor", s.moveMotor)
		router.Post("/stop-process", s.stopProcess)
		router.Route("/configuration", func(configRouter chi.Router) {
			configRouter.Get("/", s.getConfig)
			configRouter.Put("/", s.updateConfig)
		})
	})

	s.initialized = true
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
// FileServer is serving static files.
func FileServer(router *chi.Mux, root string) {
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
// FileServer is serving static files.
func FileServerImages(router *chi.Mux, root string) {
	if err := os.MkdirAll(root, os.ModePerm); err != nil {
		if glog.V(1) {
			glog.Errorln("FileServerImages - Cnnot create public folder", err.Error())
		}
		return
	}
	//fs := http.FileServer(http.Dir(root))
	fileServer := http.FileServer(http.Dir(root))
	router.Handle("/process-images/*", http.StripPrefix("/process-images", fileServer))

}

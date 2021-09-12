package server

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang/glog"
)

type controllerMachine interface {
	SetDegreesMovement(degrees float64) error
	StartProcess() error
	StopProcess() error
	MoveMotor() error
	IsWorking() bool
	CameraSnapshot(w io.Writer) error
}

//PiServer
type MachineServer struct {
	initialized bool
	Router      chi.Router
	machine     controllerMachine
}

//ListenAndServe is the main server procedure that only wraps http.ListenAndServe
func (s *MachineServer) ListenAndServe() {
	if glog.V(3) {
		glog.Infoln("MachineServer -  MachineServer.ListenAndServe start")
	}
	if !s.initialized {
		panic("Server not initialized")
	}
	http.ListenAndServe(":3333", s.Router)
}

//Init initialize the server router and set the controllerMachine needed to do the work
func (s *MachineServer) Init(machine controllerMachine) {
	if glog.V(3) {
		glog.Infoln("MachineServer -  MachineServer.Init start")
	}
	s.machine = machine
	s.Router = chi.NewRouter()
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))

	FileServer(s.Router.(*chi.Mux))
	s.Router.Route("/api", func(router chi.Router) {
		router.Post("/move-motor", s.moveMotor)
		router.Get("/machine-status", s.getMachineStatus)
		router.Post("/stop-process", s.stopProcess)
		router.Post("/start-process", s.startProcess)
		router.Get("/get-snapshot", s.cameraSnapshot)
	})
	s.initialized = true
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
// FileServer is serving static files.
func FileServer(router *chi.Mux) {
	root := "../../frontend/dist"
	fs := http.FileServer(http.Dir(root))

	router.Get("/*", func(w http.ResponseWriter, r *http.Request) {
		if _, err := os.Stat(root + r.RequestURI); os.IsNotExist(err) {
			http.StripPrefix(r.RequestURI, fs).ServeHTTP(w, r)
		} else {
			fs.ServeHTTP(w, r)
		}
	})
}

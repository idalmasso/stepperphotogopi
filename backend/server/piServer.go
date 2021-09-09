package server

import (
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/golang/glog"
)


type PiServer struct {
	Router chi.Router
}

func (s *PiServer) StartAndListen() {
	defer glog.Flush()
	http.ListenAndServe(":3333", s.Router)
}

func (s *PiServer) Init() {
	s.Router = chi.NewRouter()
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))

	FileServer(s.Router.(*chi.Mux))
	s.Router.Route("/api", func(router chi.Router) {
		router.Post("/move-motor", s.moveMotor)
	})
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

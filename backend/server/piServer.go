package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type PiServer struct {
	Router chi.Router
}

func (s *PiServer) startAndListen() {
	http.ListenAndServe(":3333", s.Router)
}

func (s *PiServer) init() {
	s.Router = chi.NewRouter()
	s.Router.Use(middleware.RequestID)
	s.Router.Use(middleware.RealIP)
	s.Router.Use(middleware.Logger)
	s.Router.Use(middleware.Recoverer)
	s.Router.Use(middleware.Timeout(60 * time.Second))

}

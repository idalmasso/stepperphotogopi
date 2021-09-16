package server

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
)

func (s *MachineServer) stopProcess(w http.ResponseWriter, r *http.Request) {
	if err := s.machine.StopProcess(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}

func (s *MachineServer) startProcess(w http.ResponseWriter, r *http.Request) {
	if glog.V(3) {
		glog.Infoln("Start process called with num photos per process=", s.configuration.Hardware.Camera.NumPhotosPerProcess)
	}
	s.machine.SetDegreesMovement(float64(360) / float64(s.configuration.Hardware.Camera.NumPhotosPerProcess))
	if err := s.machine.StartProcess(s.configuration.Server.PhotoDirectory); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}

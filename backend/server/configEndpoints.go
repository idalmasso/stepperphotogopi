package server

import (
	"encoding/json"
	"net/http"
)

func (s *MachineServer) getConfig(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(s.configuration); err != nil {
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *MachineServer) updateConfig(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(s.configuration); err != nil {
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if err := s.configuration.SaveToFile("configuration.yaml"); err != nil {
		json.NewEncoder(w).Encode(errorMessage{Message: "Cannot update config file: " + err.Error()})
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}

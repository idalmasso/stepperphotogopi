package server

import (
	"encoding/json"
	"net/http"
)

func (s *MachineServer) getConfig(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(s.configuration); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (s *MachineServer) updateConfig(w http.ResponseWriter, r *http.Request) {

	if err := json.NewDecoder(r.Body).Decode(s.configuration); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := s.configuration.SaveToFile("configuration.yaml"); err != nil {
		http.Error(w, "cannot update file config "+err.Error(), http.StatusInternalServerError)
		return
	}
	s.updateMachineFromConfig()
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}

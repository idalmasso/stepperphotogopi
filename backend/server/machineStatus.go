package server

import (
	"encoding/json"
	"net/http"
)

type responseMachineStatus struct {
	Value string `json:"value"`
}

func (s *MachineServer) getMachineStatus(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var response responseMachineStatus
	if s.machine.IsWorking() {
		response.Value = "working"
	} else {
		response.Value = "waiting"
	}
	json.NewEncoder(w).Encode(response)
}

package server

import (
	"encoding/json"
	"net/http"
)


func (s *MachineServer)stopProcess(w http.ResponseWriter, r *http.Request){
	if err:=s.machine.StopProcess(); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}

func (s *MachineServer)startProcess(w http.ResponseWriter, r *http.Request){
	s.machine.SetDegreesMovement(9)
	if err:=s.machine.StartProcess(); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}

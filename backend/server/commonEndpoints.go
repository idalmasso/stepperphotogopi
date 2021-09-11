package server

import (
	"encoding/json"
	"net/http"
)


func (s *MachineServer)StopProcess(w http.ResponseWriter, r *http.Request){
	if err:=s.machine.StopProcess(); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
	}
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}

package server

import (
	"encoding/json"
	"net/http"
)


func (s *MachineServer)cameraSnapshot(w http.ResponseWriter, r *http.Request){
	if err:=s.machine.CameraSnapshot(w); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
	}
	
	w.WriteHeader(http.StatusOK)
}

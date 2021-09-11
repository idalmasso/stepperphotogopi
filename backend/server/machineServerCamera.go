package server

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
)


func (s *MachineServer)cameraSnapshot(w http.ResponseWriter, r *http.Request){
	if glog.V(3) {
		glog.Infoln("MachineServer -  MachineServer.cameraSnapshot start ")
	}
	if err:=s.machine.CameraSnapshot(w); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
	}
	
	w.WriteHeader(http.StatusOK)
}

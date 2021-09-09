package server

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
)

type moveMotorData struct{
	Degrees float32 `json:"degrees"` 
}

func  (s PiServer) moveMotor(w http.ResponseWriter, r *http.Request){
	if glog.V(3) {
		glog.Infoln("DEBUG - PiServer.moveMotor - start")
	}
	var moveMotorStruct moveMotorData
	if err:=json.NewDecoder(r.Body).Decode(&moveMotorStruct); err!=nil{
		w.WriteHeader(http.StatusBadRequest)
		errResponse:=errorMessage{Message: "Cannot decode the degrees"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	if glog.V(3) {
		glog.Infoln("DEBUG - PiServer.moveMotor - move" ,moveMotorStruct.Degrees," degrees" )
	}
	
	w.WriteHeader(http.StatusOK)
}

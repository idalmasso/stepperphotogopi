package server

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
)

type moveMotorData struct {
	Degrees json.Number `json:"degrees"`
}

func (s MachineServer) moveMotor(w http.ResponseWriter, r *http.Request) {
	if glog.V(3) {
		glog.Infoln("MachineServer - PiServer.moveMotor - start")
	}
	var moveMotorStruct moveMotorData
	if err := json.NewDecoder(r.Body).Decode(&moveMotorStruct); err != nil {
		if glog.V(2) {
			glog.Infoln("MachineServer - PiServer.moveMotor error", err.Error())
		}
		w.WriteHeader(http.StatusBadRequest)
		errResponse := errorMessage{Message: "Cannot decode the degrees"}
		json.NewEncoder(w).Encode(errResponse)
		return
	}
	if glog.V(3) {
		glog.Infoln("MachineServer - PiServer.moveMotor - move", moveMotorStruct.Degrees, " degrees")
	}
	if degrees, err := moveMotorStruct.Degrees.Float64(); err != nil {

	} else {
		if err := s.machine.SetDegreesMovement(degrees); err != nil {
			if glog.V(2) {
				glog.Infoln("MachineServer -  MachineServer.moveMotor machineController.SetDegreesMovement returned ", err.Error())
			}
			w.WriteHeader(http.StatusFailedDependency)
			errResponse := errorMessage{Message: err.Error()}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		if err := s.machine.MoveMotor(); err != nil {
			if glog.V(2) {
				glog.Infoln("MachineServer -  MachineServer.moveMotor machineController.MoveMotor returned ", err.Error())
			}
			w.WriteHeader(http.StatusFailedDependency)
			errResponse := errorMessage{Message: err.Error()}
			json.NewEncoder(w).Encode(errResponse)
			return
		}
		json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
		w.WriteHeader(http.StatusOK)

	}
}

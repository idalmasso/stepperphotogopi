package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)


func (s *MachineServer)getListProcessDone(w http.ResponseWriter, r *http.Request){
	if values, err:=os.ReadDir("../../images"); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		return
	} else {
		var list valueListStringResponse
		for _,value:=range(values){
			if value.IsDir() && value.Name()!="." &&  value.Name()!=".." && value.Name()!=s.machine.GetActualProcessName(){
				list.Value = append(list.Value, value.Name())
			}
		}
		json.NewEncoder(w).Encode(list)
		w.WriteHeader(http.StatusOK)
	}
	
	
}

func (s *MachineServer)deleteProcessDone(w http.ResponseWriter, r *http.Request){
	val:=r.URL.Query().Get("process")
	if val!=""	{
		if s.machine.GetActualProcessName()==val{
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorMessage{Message: "cannot delete the process actually processing"})
			return
		}
		if _, err := os.Stat(filepath.Join("../../images", val)); !os.IsNotExist(err) {
			os.RemoveAll(filepath.Join("../../images", val))
		}
		
	}
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}
func (s *MachineServer)getImagesProcessDone(w http.ResponseWriter, r *http.Request){
	s.machine.SetDegreesMovement(9)
	if err:=s.machine.StartProcess(); err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		return
	}
	json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
	w.WriteHeader(http.StatusOK)
}

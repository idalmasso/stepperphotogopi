package server

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func (s *MachineServer) getListProcessDone(w http.ResponseWriter, r *http.Request) {
	if values, err := os.ReadDir(s.configuration.PhotoDirectory); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
		return
	} else {
		var list valueListStringResponse
		for _, value := range values {
			if value.IsDir() && value.Name() != "." && value.Name() != ".." && value.Name() != s.machine.GetActualProcessName() {
				list.Value = append(list.Value, value.Name())
			}
		}
		json.NewEncoder(w).Encode(list)
		w.WriteHeader(http.StatusOK)
	}

}

func (s *MachineServer) deleteProcessDone(w http.ResponseWriter, r *http.Request) {
	val := chi.URLParam(r, "process");
	if val != "" {
		if s.machine.GetActualProcessName()  == val && s.machine.IsWorking(){
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorMessage{Message: "cannot delete the process actually processing"})
			return
		}
		if _, err := os.Stat(filepath.Join(s.configuration.PhotoDirectory, val)); !os.IsNotExist(err) {
			if err=os.RemoveAll(filepath.Join(s.configuration.PhotoDirectory, val)); err!=nil{
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
				return
			}
			json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
			w.WriteHeader(http.StatusOK)
			return
		}

	}
	json.NewEncoder(w).Encode(errorMessage{Message: "not found"})
	w.WriteHeader(http.StatusNotFound)
}

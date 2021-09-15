package server

import (
	"archive/zip"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
)

func (s *MachineServer) getListProcessDone(w http.ResponseWriter, r *http.Request) {
	if values, err := os.ReadDir(s.configuration.Server.PhotoDirectory); err != nil {
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
	val := chi.URLParam(r, "process")
	if val != "" {
		if s.machine.GetActualProcessName() == val && s.machine.IsWorking() {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorMessage{Message: "cannot delete the process actually processing"})
			return
		}
		if _, err := os.Stat(filepath.Join(s.configuration.Server.PhotoDirectory, val)); !os.IsNotExist(err) {
			if err = os.RemoveAll(filepath.Join(s.configuration.Server.PhotoDirectory, val)); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
				return
			}
			if _, err := os.Stat(filepath.Join(s.configuration.Server.PhotoDirectory, val+".zip")); !os.IsNotExist(err) {
				if err = os.Remove(filepath.Join(s.configuration.Server.PhotoDirectory, val+".zip")); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
					return
				}
			}
			json.NewEncoder(w).Encode(valueResponse{Value: "ok"})
			w.WriteHeader(http.StatusOK)
			return
		}

	}
	json.NewEncoder(w).Encode(errorMessage{Message: "not found"})
	w.WriteHeader(http.StatusNotFound)
}

func (s *MachineServer) getZipProcess(w http.ResponseWriter, r *http.Request) {
	processId := chi.URLParam(r, "process")
	if processId != "" {
		if s.machine.GetActualProcessName() == processId && s.machine.IsWorking() {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(errorMessage{Message: "cannot get zip for the process actually processing"})
			return
		}
		zipFileName := filepath.Join(s.configuration.Server.PhotoDirectory, processId+".zip")
		if _, err := os.Stat(zipFileName); os.IsNotExist(err) {
			zipFile, err := os.Create(zipFileName)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
				return
			}
			//defer zipFile.Close() //cannot use defer, I want to do something if something is wrong
			writer := zip.NewWriter(zipFile)

			walker := func(path string, info os.FileInfo, err error) error {

				if info.IsDir() {
					return nil
				}
				file, err := os.Open(path)
				if err != nil {
					return err
				}

				defer file.Close()

				f, err := writer.Create(info.Name())
				if err != nil {
					return err
				}

				_, err = io.Copy(f, file)
				if err != nil {
					return err
				}

				return nil
			}
			err = filepath.Walk(filepath.Join(s.configuration.Server.PhotoDirectory, processId), walker)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(errorMessage{Message: err.Error()})
				writer.Close()
				zipFile.Close()
				os.Remove(zipFileName)
				return
			}
			writer.Close()
			zipFile.Close()
		}
		zipFileName = strings.Replace(zipFileName, filepath.Join(s.configuration.Server.PhotoDirectory), "/process-images", 1)
		json.NewEncoder(w).Encode(valueResponse{Value: zipFileName})
		w.WriteHeader(http.StatusOK)

		return

	}
	json.NewEncoder(w).Encode(errorMessage{Message: "not found"})
	w.WriteHeader(http.StatusNotFound)
}

package utils

import (
	"encoding/json"
	"net/http"
)

type AppError struct {
	Error string `json:"error"`
}

func WriteErrorJson(w http.ResponseWriter, err error, status int) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)
		raw, _ := json.Marshal(AppError{Error: err.Error()})
		w.Write(raw)
}


func WriteJson(w http.ResponseWriter, data any, status int) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(status)
		raw, _ := json.Marshal(data)
		w.Write(raw)	
}

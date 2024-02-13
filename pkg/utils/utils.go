package utils

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

type AppError struct {
	Error string `json:"error"`
}

var AppName string = "local_app"

func WriteErrorJson(w http.ResponseWriter, err error, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("x-status-code", strconv.Itoa(status))
	w.Header().Add("x-error-msg", err.Error())
	w.WriteHeader(status)
	raw, _ := json.Marshal(AppError{Error: err.Error()})
	w.Write(raw)
}

func WriteJson(w http.ResponseWriter, data any, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("x-status-code", strconv.Itoa(status))
	w.WriteHeader(status)
	raw, _ := json.Marshal(data)
	w.Write(raw)
}

func SetAppName() {
	appName := os.Getenv("APP_NAME")
	if appName != "" {
		AppName = appName
	}
}

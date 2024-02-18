package main

import (
	"fmt"
	"github.com/Gilmardealcantara/rinha/pkg/data"
	"github.com/Gilmardealcantara/rinha/pkg/statement"
	"github.com/Gilmardealcantara/rinha/pkg/transactions"
	"github.com/Gilmardealcantara/rinha/pkg/utils"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

const VERSION = "0.1.2"

func main() {
	bindAddress := os.Getenv("BIND_ADDRESS")
	if bindAddress == "" {
		bindAddress = "localhost"
	}
	bindPort := os.Getenv("BIND_PORT")
	if bindPort == "" {
		bindPort = "8000"
	}
	utils.SetAppName()

	storage := data.NewStorage()
	err := storage.CleanUp()
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /clientes/{id}/transacoes", transactions.Create(storage))
	mux.HandleFunc("GET /clientes/{id}/extrato", statement.GetStatement(storage))
	// mux.Handle("POST /clientes/{id}/transacoes", logMiddleware(transactions.Create(storage)))
	// mux.Handle("GET /clientes/{id}/extrato", logMiddleware(statement.GetStatement(storage)))
	slog.Info(
		"server " + utils.AppName + ": " + VERSION + " running in " + bindAddress + ":" + bindPort,
	)
	err = http.ListenAndServe(bindAddress+":"+bindPort, mux)
	if err != nil {
		slog.Error("Error in start server", slog.Any("error", err))
	}
}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
		message := fmt.Sprintf("%s %s", r.Method, r.URL.Path)
		statusCode, _ := strconv.Atoi(w.Header().Get("x-status-code"))
		if statusCode > 399 {
			errMsg := w.Header().Get("x-error-msg")
			slog.Info(message, slog.Int("status_code", statusCode), slog.String("err_msg", errMsg))
		}
	})
}

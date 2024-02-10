package main

import (
	"log/slog"
	"net/http"

	"github.com/Gilmardealcantara/rinha/pkg/transactions"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /clientes/{id}/transacoes", transactions.Create)
	mux.HandleFunc("GET /clientes/{id}/extrato", transactions.Create)

	slog.Info("starging server...")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		slog.Error("Error in start server", slog.Any("error", err))
	}
}

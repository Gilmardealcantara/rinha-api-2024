package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Gilmardealcantara/rinha/pkg/data"
	"github.com/Gilmardealcantara/rinha/pkg/statement"
	"github.com/Gilmardealcantara/rinha/pkg/transactions"
	"github.com/Gilmardealcantara/rinha/pkg/utils"
)

func main() {
	bindAddress := os.Getenv("BIND_ADDRESS")
	if bindAddress == "" {bindAddress = "localhost"}
	bindPort := os.Getenv("BIND_PORT")
	if bindPort == "" {bindPort = "8000"}	
	utils.SetAppName()

	storage := data.NewStorage()

	mux := http.NewServeMux()
	mux.HandleFunc("POST /clientes/{id}/transacoes", transactions.Create(storage))
	mux.HandleFunc("GET /clientes/{id}/extrato", statement.GetStatement(storage))


	slog.Info("server "+utils.AppName+" running in " + bindAddress + ":" + bindPort)
	err := http.ListenAndServe(bindAddress + ":" + bindPort, mux)
	if err != nil {
		slog.Error("Error in start server", slog.Any("error", err))
	}
}

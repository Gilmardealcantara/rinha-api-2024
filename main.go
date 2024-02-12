package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/Gilmardealcantara/rinha/pkg/statement"
	"github.com/Gilmardealcantara/rinha/pkg/transactions"
	"github.com/Gilmardealcantara/rinha/pkg/utils"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /clientes/{id}/transacoes", transactions.Create)
	mux.HandleFunc("GET /clientes/{id}/extrato", statement.GetStatement)

	bindAddress := os.Getenv("BIND_ADDRESS")
	if bindAddress == "" {bindAddress = "localhost"}
	bindPort := os.Getenv("BIND_PORT")
	if bindPort == "" {bindPort = "3000"}
	
	utils.SetAppName()
	
	slog.Info("server "+utils.AppName+" running in " + bindAddress + ":" + bindPort)
	err := http.ListenAndServe(bindAddress + ":" + bindPort, mux)
	if err != nil {
		slog.Error("Error in start server", slog.Any("error", err))
	}
}

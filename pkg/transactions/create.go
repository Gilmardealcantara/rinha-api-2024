package transactions

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Gilmardealcantara/rinha/pkg/data"
	"github.com/Gilmardealcantara/rinha/pkg/utils"
)

type Response struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}

func Create(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		utils.WriteErrorJson(w, err, 400)
		return
	}

	var payload CreateRequest
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteErrorJson(w, err, 400)
		return
	}

	if err := payload.Validate(); err != nil {
		utils.WriteErrorJson(w, err, 400)
		return
	}

	client, err := data.FindClient(id)
	if err != nil {
		utils.WriteErrorJson(w, err, 500)
		return
	}

	if client == nil {
		utils.WriteErrorJson(w, errors.New("client not found: "+idPath), 400)
		return
	}

	if err := client.PerformTransaction(payload.Type, payload.Value); err != nil {
		utils.WriteErrorJson(w, err, 422)
		return
	}

	if err = data.Save(*client); err != nil {
		utils.WriteErrorJson(w, err, 500)
		return
	}

	result := Response{
		Balance: client.Balance,
		Limit:   client.Limit,
	}

	slog.Info("CreateTransaction: id: "+idPath, slog.Any("payload", payload), slog.Any("client", client))
	utils.WriteJson(w, result, 200)
}

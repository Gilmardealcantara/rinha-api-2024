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

	var payload data.Transaction
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		utils.WriteErrorJson(w, err, 400)
		return
	}

	if err := payload.Validate(); err != nil {
		utils.WriteErrorJson(w, err, 400)
		return
	}

	storage := data.NewStorage()

	acc, err := storage.FindAccount(id)
	if err != nil {
		utils.WriteErrorJson(w, err, 500)
		return
	}

	if acc == nil {
		utils.WriteErrorJson(w, errors.New("account not found: "+idPath), 400)
		return
	}

	if err := acc.PerformTransaction(payload); err != nil {
		utils.WriteErrorJson(w, err, 422)
		return
	}

	payload.ClientId = acc.ClientId
	if err = storage.Save(*acc, payload); err != nil {
		utils.WriteErrorJson(w, err, 500)
		return
	}

	slog.Info("CreateTransaction: id: "+idPath, slog.Any("transaction", payload), slog.Any("account", acc))
	result := Response{
		Balance: acc.Balance,
		Limit:   acc.Limit,
	}

	utils.WriteJson(w, result, 200)
}

package transactions

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Gilmardealcantara/rinha/pkg/data"
	"github.com/Gilmardealcantara/rinha/pkg/utils"
)

type Response struct {
	Limit   int64 `json:"limite"`
	Balance int64 `json:"saldo"`
}

func Create(storage data.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idPath := r.PathValue("id")
		id, err := strconv.Atoi(idPath)
		if err != nil {
			utils.WriteErrorJson(
				w,
				errors.Join(errors.New("invalid id:"+idPath), err),
				http.StatusUnprocessableEntity,
			)
			return
		}

		var payload data.Transaction
		err = json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			utils.WriteErrorJson(
				w,
				errors.Join(errors.New("error to decode r.Body"), err),
				http.StatusUnprocessableEntity,
			)
			return
		}

		if err := payload.Validate(); err != nil {
			raw, _ := json.Marshal(payload)
			utils.WriteErrorJson(
				w,
				errors.Join(errors.New("payload validate fail:"+string(raw)), err),
				http.StatusUnprocessableEntity,
			)
			return
		}

		acc, err := storage.FindAccount(id)
		if err != nil {
			utils.WriteErrorJson(w, err, 500)
			return
		}

		if acc == nil {
			utils.WriteErrorJson(
				w,
				errors.New("account not found: "+idPath),
				http.StatusUnprocessableEntity,
			)
			return
		}

		if err := acc.PerformTransaction(&payload); err != nil {
			utils.WriteErrorJson(w, err, http.StatusUnprocessableEntity)
			return
		}

		payload.ClientId = acc.ClientId
		if err = storage.Save(*acc, payload); err != nil {
			utils.WriteErrorJson(w, err, 500)
			return
		}

		// slog.Info("CreateTransaction: id: "+idPath, slog.String("app_name", utils.AppName), slog.Any("transaction", payload), slog.Any("account", acc))
		result := Response{
			Balance: acc.Balance,
			Limit:   acc.Limit,
		}

		utils.WriteJson(w, result, 200)
	}
}

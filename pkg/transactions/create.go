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
	Limit   int `json:"limite"`
	Balance int `json:"saldo"`
}

func Create(storage data.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idPath := r.PathValue("id")
		id, err := strconv.Atoi(idPath)
		if err != nil {
			utils.WriteErrorJson(w, errors.Join(errors.New("invalid id:"+idPath), err), http.StatusUnprocessableEntity)
			return
		}

		var payload data.Transaction
		err = json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			utils.WriteErrorJson(w, errors.Join(errors.New("error to decode r.Body"), err), http.StatusUnprocessableEntity)
			return
		}

		if err := payload.Validate(); err != nil {
			raw, _ := json.Marshal(payload)
			utils.WriteErrorJson(w, errors.Join(errors.New("payload validate fail:"+string(raw)), err), http.StatusUnprocessableEntity)
			return
		}

		acc, derr := storage.SaveSafety(id, payload)
		// acc, derr := storage.SaveOptimistic(id, payload)
		if derr != nil {
			utils.WriteErrorJson(w, derr, derr.Code)
			return
		}

		result := Response{
			Balance: acc.Balance,
			Limit:   acc.Limit,
		}

		utils.WriteJson(w, result, 200)
	}
}

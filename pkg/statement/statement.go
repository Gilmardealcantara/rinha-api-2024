package statement

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Gilmardealcantara/rinha/pkg/data"
	"github.com/Gilmardealcantara/rinha/pkg/utils"
)

func GetStatement(w http.ResponseWriter, r *http.Request) {
	idPath := r.PathValue("id")
	id, err := strconv.Atoi(idPath)
	if err != nil {
		fmt.Println("ID : ", idPath, ", err: ", err.Error())
		utils.WriteErrorJson(w, err, http.StatusBadRequest)
		return
	}

	storage := data.NewStorage()

	acc, err := storage.FindAccount(id)
	if err != nil {
		utils.WriteErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	if acc == nil {
		utils.WriteErrorJson(w, errors.New("account not found"), http.StatusNotFound)
		return
	}

	transaction, err := storage.GetTransactions(acc.ClientId)
	if err != nil {
		utils.WriteErrorJson(w, err, http.StatusInternalServerError)
		return
	}

	result := Response{
		Balance: BalanceResult{
			Total: acc.Balance,
			Date:  getTimeStr(AppNow()),
			Limit: acc.Limit,
		},
		LastTransactions: transaction,
	}
	utils.WriteJson(w, result, http.StatusOK)
}

var AppNow = func() time.Time {
	return time.Now()
}

func getTimeStr(t time.Time) string {
	return t.Format(time.RFC3339Nano)
}

type Response struct {
	Balance          BalanceResult      `json:"saldo"`
	LastTransactions []data.Transaction `json:"ultimas_transacoes"`
}

type BalanceResult struct {
	Total int64  `json:"total"`
	Date  string `json:"data_extrato"`
	Limit int64  `json:"limite"`
}

package statement

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"slices"
	"strconv"
	"time"

	"github.com/Gilmardealcantara/rinha/pkg/data"
	"github.com/Gilmardealcantara/rinha/pkg/utils"
)

func GetStatement(storage data.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idPath := r.PathValue("id")
		id, err := strconv.Atoi(idPath)
		if err != nil {
			fmt.Println("ID : ", idPath, ", err: ", err.Error())
			utils.WriteErrorJson(w, err, http.StatusBadRequest)
			return
		}

		acc, err := storage.FindAccount(id)
		if err != nil {
			slog.Error("GetStatement FindAccount error ", slog.String("error", err.Error()))
			utils.WriteErrorJson(w, err, http.StatusInternalServerError)
			return
		}

		if acc == nil {
			utils.WriteErrorJson(
				w,
				errors.New("account not found id: "+idPath),
				http.StatusNotFound,
			)
			return
		}

		// slog.Info("GetStatement: client_id: "+idPath, slog.String("app_name", utils.AppName), slog.Any("account", acc))

		transactions, err := storage.GetTransactions(acc.ClientId)
		if err != nil {
			slog.Error("GetStatement GetTransactions error ", slog.String("error", err.Error()))
			utils.WriteErrorJson(w, err, http.StatusInternalServerError)
			return
		}

		slices.SortFunc(transactions, func(a, b data.Transaction) int {
			return b.CreatedAt.Compare(a.CreatedAt)
		})

		result := Response{
			Balance: BalanceResult{
				Total: acc.Balance,
				Date:  getTimeStr(AppNow()),
				Limit: acc.Limit,
			},
			LastTransactions: transactions,
		}
		utils.WriteJson(w, result, http.StatusOK)
	}
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

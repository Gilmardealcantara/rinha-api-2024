package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Gilmardealcantara/rinha/pkg/data"
	"github.com/Gilmardealcantara/rinha/pkg/statement"
	"github.com/Gilmardealcantara/rinha/pkg/transactions"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	cases := []struct {
		name    string
		payload string
	}{{
		name:    "invalid type",
		payload: `{"valor": 1000,"tipo":"x","descricao" : "descricao"}`,
	}, {
		name:    "invalid value",
		payload: `{"valor": 0,"tipo":"c","descricao" : "descricao"}`,
	}, {
		name:    "invalid description",
		payload: `{"valor": 1000,"tipo":"x","descricao" : "12345678901"}`,
	}}

	storage := setupStorage(t)

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			request := httptest.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(c.payload)))
			request.SetPathValue("id", "1")
			recorder := httptest.NewRecorder()
			transactions.Create(storage)(recorder, request)
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		})
	}

	t.Run("with sucess credit", func(t *testing.T) {
		payload := `{"valor": 1000,"tipo":"c","descricao" : "descricao"}`
		request := httptest.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(payload)))
		request.SetPathValue("id", "1")
		recorder := httptest.NewRecorder()
		transactions.Create(storage)(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, recorder.Body.String())
		assert.Equal(t, `{"limite":100000,"saldo":1000}`, recorder.Body.String())
	})

	t.Run("with sucess debit", func(t *testing.T) {
		payload := `{"valor": 1000,"tipo":"d","descricao" : "descricao"}`
		request := httptest.NewRequest("POST", "/clientes/2/transacoes", bytes.NewBuffer([]byte(payload)))
		request.SetPathValue("id", "2")
		recorder := httptest.NewRecorder()
		transactions.Create(storage)(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, recorder.Body.String())
		assert.Equal(t, `{"limite":80000,"saldo":-1000}`, recorder.Body.String())
	})

	t.Run("with unprocessable debit", func(t *testing.T) {
		payload := `{"valor": 1000001,"tipo":"d","descricao" : "descricao"}`
		request := httptest.NewRequest("POST", "/clientes/3/transacoes", bytes.NewBuffer([]byte(payload)))
		request.SetPathValue("id", "3")
		recorder := httptest.NewRecorder()
		transactions.Create(storage)(recorder, request)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code, recorder.Body.String())
		assert.Equal(t, `{"error":"insufficient limit"}`, recorder.Body.String())
	})
}

func TestGetStatment(t *testing.T) {
	storage := setupStorage(t)
	t.Run("with sucess", func(t *testing.T) {
		mockTimeNow(t, "2024-01-12T11:45:26.371Z")
		request := httptest.NewRequest("GET", "/clientes/1/extrato", nil)
		request.SetPathValue("id", "1")
		recorder := httptest.NewRecorder()

		statement.GetStatement(storage)(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, recorder.Body.String())
		expectedValue := `{"saldo":{"total":0,"data_extrato":"2024-01-12T11:45:26.371Z","limite":100000},"ultimas_transacoes":[]}`
		assert.Equal(t, expectedValue, recorder.Body.String())
	})
}

func mockTimeNow(t *testing.T, str string) {
	now, err := time.Parse(time.RFC3339Nano, str)
	assert.NoError(t, err)
	statement.AppNow = func() time.Time { return now }
}

func setupStorage(t *testing.T) data.Storage {
	storage := data.NewStorage()
	err := storage.CleanUp()
	if err != nil {
		assert.FailNow(t, err.Error())
	}
	return storage
}

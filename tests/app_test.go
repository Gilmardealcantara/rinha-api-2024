package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

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

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			request := httptest.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(c.payload)))
			request.SetPathValue("id", "1")
			recorder := httptest.NewRecorder()
			transactions.Create(recorder, request)
			assert.Equal(t, http.StatusBadRequest, recorder.Code)
		})
	}

	t.Run("with sucess credit", func(t *testing.T) {
		payload := `{"valor": 1000,"tipo":"c","descricao" : "descricao"}`
		request := httptest.NewRequest("POST", "/clientes/1/transacoes", bytes.NewBuffer([]byte(payload)))
		request.SetPathValue("id", "1")
		recorder := httptest.NewRecorder()
		transactions.Create(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, recorder.Body.String())
		assert.Equal(t, `{"limite":100000,"saldo":1000}`, recorder.Body.String())
	})

	t.Run("with sucess debit", func(t *testing.T) {
		payload := `{"valor": 1000,"tipo":"d","descricao" : "descricao"}`
		request := httptest.NewRequest("POST", "/clientes/2/transacoes", bytes.NewBuffer([]byte(payload)))
		request.SetPathValue("id", "2")
		recorder := httptest.NewRecorder()
		transactions.Create(recorder, request)

		assert.Equal(t, http.StatusOK, recorder.Code, recorder.Body.String())
		assert.Equal(t, `{"limite":80000,"saldo":-1000}`, recorder.Body.String())
	})

	t.Run("with unprocessable debit", func(t *testing.T) {
		payload := `{"valor": 1000001,"tipo":"d","descricao" : "descricao"}`
		request := httptest.NewRequest("POST", "/clientes/3/transacoes", bytes.NewBuffer([]byte(payload)))
		request.SetPathValue("id", "3")
		recorder := httptest.NewRecorder()
		transactions.Create(recorder, request)

		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code, recorder.Body.String())
		assert.Equal(t, `{"error":"insufficient limit"}`, recorder.Body.String())
	})

}

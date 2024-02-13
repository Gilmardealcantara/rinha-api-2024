package data

import (
	"errors"
	"time"

	"gopkg.in/validator.v2"
)

type Transaction struct {
	Id          int       `json:"-"`
	ClientId    int       `json:"client_id"`
	Value       int64     `json:"valor" validate:"min=1"`
	Type        string    `json:"tipo"`
	Description string    `json:"descricao" validate:"min=1,max=10"`
	CreatedAt   time.Time `json:"realizada_em"`
}

func (r Transaction) Validate() error {
	if err := validator.Validate(r); err != nil {
		return err
	}
	if r.Type != "c" && r.Type != "d" {
		return errors.New("invalid type: " + r.Type)
	}
	return nil
}

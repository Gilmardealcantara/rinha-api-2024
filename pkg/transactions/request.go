package transactions

import (
	"errors"
	"gopkg.in/validator.v2"
)

type CreateRequest struct {
	Value       int64 `json:"valor" validate:"min=1"`
	Type        string `json:"tipo"`
	Description string `json:"descricao" validate:"min=1,max=10"`
}

func (r CreateRequest) Validate() error {
	if err := validator.Validate(r); err != nil 	{
		return err
	}
	if r.Type != "c" && r.Type != "d" {
		return errors.New("invaid type: " + r.Type)
	} 
	return nil 
}

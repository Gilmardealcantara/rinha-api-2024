package data

import (
	"errors"
	"time"
)

type Account struct {
	ClientId   int
	ClientName string
	Limit      int
	Balance    int
}

func (c *Account) PerformTransaction(payload *Transaction) error {
	payload.CreatedAt = time.Now()
	if payload.Type == "c" {
		return c.credit(payload.Value)
	}
	return c.debit(payload.Value)
}

func (c *Account) credit(value int) error {
	c.Balance += value
	return nil
}

func (c *Account) debit(value int) error {
	newBalance := c.Balance - value
	if newBalance+c.Limit < 0 {
		return errors.New("insufficient limit")
	}
	c.Balance = newBalance
	return nil
}

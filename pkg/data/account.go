package data

import (
	"errors"
	"time"
)

type Account struct {
	ClientId   int
	ClientName string
	Limit      int64
	Balance    int64
}

func (c *Account) PerformTransaction(payload *Transaction) error {
	payload.CreatedAt = time.Now()
	if payload.Type == "c" {
		return c.credit(payload.Value)
	}

	return c.debit(payload.Value)
}

func (c *Account) credit(value int64) error {
	c.Balance += value
	return nil
}

func (c *Account) debit(value int64) error {
	newBalance := c.Balance - value
	if c.Limit+newBalance < 0 {
		return errors.New("insufficient limit")
	}
	c.Balance = newBalance
	return nil
}

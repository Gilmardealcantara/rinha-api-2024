package data

import (
	"errors"
	"time"
)

type Client struct {
	Id      int
	Limit   int64
	Balance int64
}

func(c *Client) PerformTransaction(payload Transaction) error {
	payload.CreatedAt = time.Now()
	if payload.Type == "c" {
		return c.credit(payload.Value)
	}

	return c.debit(payload.Value)
}

func (c *Client)credit(value int64) error {
	c.Balance += value
	return nil
}


func (c *Client)debit(value int64) error {
	newBalance := c.Balance -  value
	if newBalance < (c.Limit*-1) {
		return errors.New("insufficient limit")
	}
	c.Balance = newBalance 
	return nil
}



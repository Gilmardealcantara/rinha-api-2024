package data

import "errors"

type Client struct {
	Id      int
	Limit   int64
	Balance int64
}

func(c *Client) PerformTransaction(typ string, value int64) error {
	if typ == "c" {
		return c.credit(value)
	}

	return c.debit(value)
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

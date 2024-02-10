package data

import (
	"slices"
)

type imemory struct {
	clients []Client
}

func newImemoryStorage() *imemory {
	var clients = []Client{
		{Id: 1, Limit: 100_000, Balance: 0},
		{Id: 2, Limit: 80_000, Balance: 0},
		{Id: 3, Limit: 1_000_000, Balance: 0},
		{Id: 4, Limit: 10_000_000, Balance: 0},
		{Id: 5, Limit: 500_000, Balance: 0},
	}
	return &imemory{clients}
}

func (s imemory) FindClient(id int) (*Client, error) {
	index := slices.IndexFunc(s.clients, func(c Client) bool { return c.Id == id })
	if index < 0 {
		return nil, nil
	}
	return &s.clients[index], nil
}

func (s imemory)GetTransactions(clientId int) ([]Transaction, error) {
	return []Transaction{}, nil
}

func (s imemory)Save(client Client, t Transaction) error {
	return nil
}

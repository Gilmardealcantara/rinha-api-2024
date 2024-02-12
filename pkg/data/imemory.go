package data

import (
	"slices"
)

type imemory struct {
	clients []Account
}

func newImemoryStorage() *imemory {
	clients := []Account{
		{ClientId: 1, Limit: 100_000, Balance: 0},
		{ClientId: 2, Limit: 80_000, Balance: 0},
		{ClientId: 3, Limit: 1_000_000, Balance: 0},
		{ClientId: 4, Limit: 10_000_000, Balance: 0},
		{ClientId: 5, Limit: 500_000, Balance: 0},
	}
	return &imemory{clients}
}

func (s imemory) FindAccount(id int) (*Account, error) {
	index := slices.IndexFunc(s.clients, func(c Account) bool { return c.ClientId == id })
	if index < 0 {
		return nil, nil
	}
	return &s.clients[index], nil
}

func (s imemory) GetTransactions(clientId int) ([]Transaction, error) {
	return []Transaction{}, nil
}

func (s imemory) Save(client Account, t Transaction) error {
	return nil
}

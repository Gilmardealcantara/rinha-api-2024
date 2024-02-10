package data

import (
	"slices"
)

var clients = []Client{
	{Id: 1, Limit: 100_000, Balance: 0},
	{Id: 2, Limit: 80_000, Balance: 0},
	{Id: 3, Limit: 1_000_000, Balance: 0},
	{Id: 4, Limit: 10_000_000, Balance: 0},
	{Id: 5, Limit: 500_000, Balance: 0},
}

func FindClient(id int) (*Client, error) {
	index := slices.IndexFunc(clients, func(c Client) bool { return c.Id == id })
	if index < 0 {
		return nil, nil
	}
	return &clients[index], nil
}

func Save(client Client, t Transaction) error {
	return nil
}


func GetTransactions(clientId int) ([]Transaction, error) {
	return []Transaction{}, nil
}

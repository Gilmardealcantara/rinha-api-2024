package data


type Storage interface {
	FindClient(id int) (*Client, error)
	GetTransactions(clientId int) ([]Transaction, error)
	Save(client Client, t Transaction) error
}


func NewStorage() Storage {
	return newImemoryStorage()
}



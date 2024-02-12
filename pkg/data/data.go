package data

type Storage interface {
	FindAccount(id int) (*Account, error)
	GetTransactions(clientId int) ([]Transaction, error)
	Save(client Account, t Transaction) error
}

func NewStorage() Storage {
	return newImemoryStorage()
	// return newPgImpl()
}

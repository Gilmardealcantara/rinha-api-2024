package data

type Storage interface {
	FindAccount(id int) (*Account, error)
	GetTransactions(clientId int) ([]Transaction, error)
	Save(client Account, t Transaction) error
	SaveSafety(accId int, t Transaction) (acc Account, derr *DataError)
	CleanUp() error
}

func NewStorage() Storage {
	// return newImemoryStorage()
	return newPgImpl()
}

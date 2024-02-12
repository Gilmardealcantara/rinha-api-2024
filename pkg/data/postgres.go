package data

import (
	"context"
	"database/sql"
	"log/slog"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type pgImpl struct {
	dbpool *pgxpool.Pool
}

func newPgImpl() Storage {
	dbconn, err := pgxpool.NewWithConfig(context.Background(), Config())
	if err != nil {
		panic(err)
	}	

	err = dbconn.Ping(context.Background())
	if err != nil {
		panic(err)
	}

	return &pgImpl{dbconn}
}

func (i *pgImpl) FindAccount(id int) (*Account, error) {
	var acc Account
	err := i.dbpool.QueryRow(context.Background(), "select id, nome, limite, saldo from clientes where id=$1", id).
		Scan(&acc.ClientId, &acc.ClientName, &acc.Limit, &acc.Balance)	
	if err != nil && err == sql.ErrNoRows {
		return nil, nil
	}
	return &acc, err
}

func (i *pgImpl) GetTransactions(clientId int) ([]Transaction, error) {
	return []Transaction{}, nil
}

func (i *pgImpl) Save(client Account, t Transaction) error {
	return nil
}

func Config() *pgxpool.Config {
	const defaultMaxConns = int32(4)
	const defaultMinConns = int32(0)
	const defaultMaxConnLifetime = time.Hour
	const defaultMaxConnIdleTime = time.Minute * 30
	const defaultHealthCheckPeriod = time.Minute
	const defaultConnectTimeout = time.Second * 5

	// Your own Database URL
	var DATABASE_URL string = os.Getenv("DATABASE_URL")
	if DATABASE_URL == "" {
		DATABASE_URL = "postgres://admin:123@localhost:5433/rinha?"
	}

	dbConfig, err := pgxpool.ParseConfig(DATABASE_URL)
	if err != nil {
		slog.Error("Failed to create a config, error: ", err)
	}

	dbConfig.MaxConns = defaultMaxConns
	dbConfig.MinConns = defaultMinConns
	dbConfig.MaxConnLifetime = defaultMaxConnLifetime
	dbConfig.MaxConnIdleTime = defaultMaxConnIdleTime
	dbConfig.HealthCheckPeriod = defaultHealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = defaultConnectTimeout

	// dbConfig.BeforeAcquire = func(ctx context.Context, c *pgx.Conn) bool {
	// 	slog.Info("Before acquiring the connection pool to the database!!")
	// 	return true
	// }
	//
	// dbConfig.AfterRelease = func(c *pgx.Conn) bool {
	// 	slog.Info("After releasing the connection pool to the database!!")
	// 	return true
	// }
	//
	// dbConfig.BeforeClose = func(c *pgx.Conn) {
	// 	slog.Println("Closed the connection pool to the database!!")
	// }

	return dbConfig
}

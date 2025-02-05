package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
}

// Store provides all function excute bd queries and transactions
type SQLStore struct {
	*Queries
	connPool *pgxpool.Pool
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) *SQLStore {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

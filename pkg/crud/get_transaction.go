package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

// GetTransaction ..
func GetTransaction(ctx context.Context, transactionID string) (*ent.Transaction, error) {
	return db.Client().
		Transaction.
		Query().
		Where(
			transaction.TransactionIDEQ(transactionID),
		).
		Only(ctx)
}

func GetTransactionExist(ctx context.Context, transactionID string) (bool, error) {
	return db.Client().
		Transaction.
		Query().
		Where(
			transaction.TransactionIDEQ(transactionID),
		).Exist(ctx)
}

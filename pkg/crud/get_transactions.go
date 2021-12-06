package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
)

// GetTransactions ..
func GetTransactions(ctx context.Context) ([]*ent.Transaction, error) {
	return db.Client().
		Transaction.
		Query().
		Where(
			transaction.StateIn(transaction.StateWait, transaction.StateSign),
		).
		Order(ent.Asc(transaction.FieldCreatedAt, transaction.FieldNonce)).
		Limit(constant.DefaultPageSize).
		All(ctx)
}

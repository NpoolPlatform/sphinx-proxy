package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
)

type GetTransactionsParams struct {
	State transaction.State
}

// GetTransactions ..
func GetTransactions(ctx context.Context, param GetTransactionsParams) ([]*ent.Transaction, error) {
	return db.Client().
		Transaction.
		Query().
		Where(
			transaction.StateEQ(param.State),
		).
		Order(ent.Asc(transaction.FieldCreateAt, transaction.FieldNonce)).
		Limit(constant.DefaultPageSize).
		All(ctx)
}

package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
)

// GetTransactions ..
func GetTransactions(ctx context.Context) ([]*ent.Transaction, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
	return client.
		Transaction.
		Query().
		Where(
			transaction.StateIn(
				uint8(sphinxproxy.TransactionState_TransactionStateWait),
				uint8(sphinxproxy.TransactionState_TransactionStateSign),
				uint8(sphinxproxy.TransactionState_TransactionStateSync),
			),
		).
		Order(ent.Asc(transaction.FieldCreatedAt, transaction.FieldNonce)).
		Limit(constant.DefaultPageSize).
		All(ctx)
}

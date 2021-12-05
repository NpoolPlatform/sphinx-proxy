package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

type CreateTransactionParam struct {
	CoinType      sphinxplugin.CoinType
	TransactionID string
	From          string
	To            string
	Value         float64
}

func CreateTransaction(ctx context.Context, t CreateTransactionParam) error {
	_, err := db.
		Client().
		Transaction.
		Create().
		SetCoinType(int32(t.CoinType)).
		SetTransactionID(t.TransactionID).
		SetFrom(t.From).
		SetTo(t.To).
		SetValue(t.Value).
		SetState(transaction.StateWait).
		Save(ctx)
	return err
}

package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	constant "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
)

type GetTransactionsParam struct {
	CoinType         sphinxplugin.CoinType
	TransactionState sphinxproxy.TransactionState
}

// GetTransactions ..
func GetTransactions(ctx context.Context, param GetTransactionsParam) ([]*ent.Transaction, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}
	return client.
		Transaction.
		Query().
		Where(
			transaction.CoinTypeEQ(int32(param.CoinType)),
			transaction.StateEQ(uint8(param.TransactionState)),
		).
		Order(ent.Asc(transaction.FieldCreatedAt)).
		Limit(constant.DefaultPageSize).
		All(ctx)
}

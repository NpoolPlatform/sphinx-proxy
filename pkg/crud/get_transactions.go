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
func GetTransactions(ctx context.Context, params GetTransactionsParam) ([]*ent.Transaction, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	stmt := client.
		Transaction.
		Query().
		Where(
			transaction.StateEQ(uint8(params.TransactionState)),
		)

	if params.CoinType != sphinxplugin.CoinType_CoinTypeUnKnow {
		stmt = stmt.Where(transaction.CoinTypeEQ(int32(params.CoinType)))
	}

	return stmt.Order(ent.Asc(transaction.FieldCreatedAt)).
		Limit(constant.DefaultPageSize).
		All(ctx)
}

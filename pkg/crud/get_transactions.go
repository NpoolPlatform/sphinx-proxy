package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type GetTransactionsParam struct {
	CoinType         sphinxplugin.CoinType
	TransactionState sphinxproxy.TransactionState
}

// GetTransactions ..
func GetTransactions(ctx context.Context, params GetTransactionsParam) ([]*ent.Transaction, error) {
	_, span := otel.Tracer(sconst.ServiceName).Start(ctx, "GetTransactions")
	defer span.End()

	span.SetAttributes(
		attribute.String("CoinType", utils.TruncateCoinTypePrefix(params.CoinType)),
		attribute.Int64("TransactionState", int64(params.TransactionState)),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
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
		Limit(sconst.DefaultPageSize).
		All(ctx)
}

package crud

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/utils"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

type CreateTransactionParam struct {
	CoinType         sphinxplugin.CoinType
	TransactionState sphinxproxy.TransactionState
	TransactionID    string
	Name             string
	From             string
	To               string
	Value            float64
	Memo             string
}

func CreateTransaction(ctx context.Context, t *CreateTransactionParam) error {
	_, span := otel.Tracer(sconst.ServiceName).Start(ctx, "CreateTransaction")
	defer span.End()

	span.SetAttributes(
		attribute.String("CoinType", utils.TruncateCoinTypePrefix(t.CoinType)),
		attribute.String("TransactionID", t.TransactionID),
		attribute.String("Name", t.Name),
		attribute.String("From", t.From),
		attribute.String("To", t.To),
		attribute.Float64("Value", t.Value),
		attribute.String("Memo", t.Memo),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return err
	}
	_, err = client.Transaction.Create().
		SetCoinType(int32(t.CoinType)).
		SetTransactionID(t.TransactionID).
		SetName(t.Name).
		SetFrom(t.From).
		SetTo(t.To).
		SetMemo(t.Memo).
		SetAmount(price.VisualPriceToDBPrice(t.Value)).
		SetState(uint8(t.TransactionState)).
		Save(ctx)
	return err
}

package crud

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
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
	client, err := db.Client()
	if err != nil {
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

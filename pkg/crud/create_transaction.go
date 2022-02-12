package crud

import (
	"context"

	"github.com/NpoolPlatform/go-service-framework/pkg/price"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
)

type CreateTransactionParam struct {
	CoinType      sphinxplugin.CoinType
	TransactionID string
	From          string
	To            string
	Value         float64
}

func CreateTransaction(ctx context.Context, t CreateTransactionParam) error {
	client, err := db.Client()
	if err != nil {
		return err
	}
	_, err = client.Transaction.Create().
		SetCoinType(int32(t.CoinType)).
		SetTransactionID(t.TransactionID).
		SetFrom(t.From).
		SetTo(t.To).
		SetAmount(price.VisualPriceToDBPrice(t.Value)).
		SetState(uint8(sphinxproxy.TransactionState_TransactionStateWait)).
		Save(ctx)
	return err
}

package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// update nonce/utxo and state
type UpdateTransactionParams struct {
	TransactionID string
	State         sphinxproxy.TransactionState
	NextState     sphinxproxy.TransactionState
	Payload       []byte
	Cid           string
	ExitCode      int64
}

// UpdateTransaction update transaction info
func UpdateTransaction(ctx context.Context, t *UpdateTransactionParams) error {
	_, span := otel.Tracer(sconst.ServiceName).Start(ctx, "UpdateTransaction")
	defer span.End()

	span.SetAttributes(
		attribute.String("TransactionID", t.TransactionID),
		attribute.Int64("State", int64(t.State)),
		attribute.Int64("NextState", int64(t.NextState)),
		attribute.String("Cid", t.Cid),
		attribute.Int64("ExitCode", t.ExitCode),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return err
	}

	stmt := client.
		Transaction.
		Update().
		Where(
			transaction.TransactionIDEQ(t.TransactionID),
			transaction.StateEQ(uint8(t.State)),
		).
		SetPayload(t.Payload).
		SetState(uint8(t.NextState)).
		SetExitCode(t.ExitCode)

	if t.Cid != "" {
		stmt.
			SetCid(t.Cid)
	}

	return stmt.Exec(ctx)
}

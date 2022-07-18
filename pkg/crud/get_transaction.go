package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
	sconst "github.com/NpoolPlatform/sphinx-proxy/pkg/message/const"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
)

// GetTransaction ..
func GetTransaction(ctx context.Context, transactionID string) (*ent.Transaction, error) {
	_, span := otel.Tracer(sconst.ServiceName).Start(ctx, "GetTransaction")
	defer span.End()

	span.SetAttributes(
		attribute.String("TransactionID", transactionID),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return nil, err
	}

	return client.
		Transaction.
		Query().
		Select(
			transaction.FieldTransactionID,
			transaction.FieldCoinType,
			transaction.FieldFrom,
			transaction.FieldTo,
			transaction.FieldAmount,
			transaction.FieldCid,
			transaction.FieldExitCode,
			transaction.FieldPayload,
			transaction.FieldState,
			transaction.FieldCreatedAt,
			transaction.FieldUpdatedAt,
		).
		Where(
			transaction.TransactionIDEQ(transactionID),
		).
		Only(ctx)
}

type GetTransactionExistParam struct {
	TransactionID    string
	TransactionState sphinxproxy.TransactionState
}

func GetTransactionExist(ctx context.Context, params GetTransactionExistParam) (bool, error) {
	_, span := otel.Tracer(sconst.ServiceName).Start(ctx, "GetTransactionExist")
	defer span.End()

	span.SetAttributes(
		attribute.String("TransactionID", params.TransactionID),
		attribute.Int64("TransactionState", int64(params.TransactionState)),
	)

	client, err := db.Client()
	if err != nil {
		span.SetStatus(codes.Error, "get db client fail")
		span.RecordError(err)
		return false, err
	}

	stm := client.
		Transaction.
		Query().
		Select(transaction.FieldID).
		Where(
			transaction.TransactionIDEQ(params.TransactionID),
		)

	if params.TransactionState != sphinxproxy.TransactionState_TransactionStateUnKnow {
		stm.Where(transaction.StateEQ(uint8(params.TransactionState)))
	}

	return stm.Exist(ctx)
}

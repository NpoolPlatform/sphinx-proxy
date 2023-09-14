package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

// GetTransaction ..
func GetTransaction(ctx context.Context, transactionID string) (*ent.Transaction, error) {
	client, err := db.Client()
	if err != nil {
		return nil, err
	}

	return client.
		Transaction.
		Query().
		Select(
			transaction.FieldTransactionID,
			transaction.FieldCoinType,
			transaction.FieldName,
			transaction.FieldFrom,
			transaction.FieldTo,
			transaction.FieldMemo,
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
	client, err := db.Client()
	if err != nil {
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

package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

// update nonce/utxo and state
type UpdateTransactionParams struct {
	TransactionID string
	State         sphinxproxy.TransactionState
	Nonce         uint64
	Cid           string
	ExitCode      int64
}

// UpdateTransaction update transaction info
func UpdateTransaction(ctx context.Context, t UpdateTransactionParams) error {
	stm := db.
		Client().
		Transaction.
		Update().
		Where(transaction.TransactionIDEQ(t.TransactionID))

	switch t.State {
	case sphinxproxy.TransactionState_TransactionStateSign:
		stm.SetNonce(t.Nonce)
		stm.Where(transaction.StateEQ(uint8(sphinxproxy.TransactionState_TransactionStateWait)))
	case sphinxproxy.TransactionState_TransactionStateSync:
		stm.SetCid(t.Cid)
		stm.Where(transaction.StateEQ(uint8(sphinxproxy.TransactionState_TransactionStateSign)))
	case sphinxproxy.TransactionState_TransactionStateDone:
		stm.SetExitCode(t.ExitCode)
		stm.Where(transaction.StateEQ(uint8(sphinxproxy.TransactionState_TransactionStateSync)))
	}

	_, err := stm.SetState(uint8(t.State)).
		Save(ctx)
	return err
}

func ConfirmTransaction(ctx context.Context, transactionID string) error {
	return db.
		Client().
		Transaction.
		Update().
		SetState(uint8(sphinxproxy.TransactionState_TransactionStateWait)).
		Where(transaction.TransactionIDEQ(transactionID)).Exec(ctx)
}

func RejectTransaction(ctx context.Context, transactionID string) error {
	return db.
		Client().
		Transaction.
		Update().
		SetState(uint8(sphinxproxy.TransactionState_TransactionStateRejected)).
		Where(transaction.TransactionIDEQ(transactionID)).Exec(ctx)
}

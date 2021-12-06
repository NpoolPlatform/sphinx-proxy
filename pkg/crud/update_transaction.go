package crud

import (
	"context"

	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

// update nonce/utxo and state
type UpdateTransactionParams struct {
	TransactionID string
	State         transaction.State
	Nonce         uint64
	Cid           string
}

// UpdateTransaction update transaction info
func UpdateTransaction(ctx context.Context, t UpdateTransactionParams) error {
	stm := db.
		Client().
		Transaction.
		Update().
		Where(transaction.TransactionIDEQ(t.TransactionID))

	switch t.State {
	case transaction.StateSign:
		stm.SetNonce(t.Nonce)
		stm.Where(transaction.StateEQ(transaction.StateWait))
	case transaction.StateDone:
		stm.SetCid(t.Cid)
		stm.Where(transaction.StateEQ(transaction.StateSign))
	case transaction.StateFail:
	}

	_, err := stm.SetState(t.State).
		Save(ctx)
	return err
}

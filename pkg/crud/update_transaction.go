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
	NextState     sphinxproxy.TransactionState
	Payload       []byte
	Cid           string
	ExitCode      int64
}

// UpdateTransaction update transaction info
func UpdateTransaction(ctx context.Context, t *UpdateTransactionParams) error {
	client, err := db.Client()
	if err != nil {
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

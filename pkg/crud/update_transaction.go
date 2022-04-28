package crud

import (
	"context"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
	"github.com/NpoolPlatform/sphinx-plugin/pkg/plugin/eth"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db"
	"github.com/NpoolPlatform/sphinx-proxy/pkg/db/ent/transaction"
)

// update nonce/utxo and state
type UpdateTransactionParams struct {
	TransactionID string
	State         sphinxproxy.TransactionState
	RecentBhash   string
	Nonce         uint64
	// TODO optimize
	UTXO     []*sphinxplugin.Unspent
	Pre      *eth.PreSignInfo
	Cid      string
	TxData   []byte
	ExitCode int64
}

// UpdateTransaction update transaction info
func UpdateTransaction(ctx context.Context, t *UpdateTransactionParams) error {
	client, err := db.Client()
	if err != nil {
		return err
	}
	stm := client.
		Transaction.
		Update().
		Where(transaction.TransactionIDEQ(t.TransactionID))

	switch t.State {
	case sphinxproxy.TransactionState_TransactionStateSign:
		stm.SetNonce(t.Nonce)
		stm.SetUtxo(t.UTXO)
		stm.SetPre(t.Pre)
		stm.SetRecentBhash(t.RecentBhash)
		stm.SetTxData(t.TxData)
		stm.Where(transaction.StateEQ(uint8(sphinxproxy.TransactionState_TransactionStateWait)))
	case sphinxproxy.TransactionState_TransactionStateSync:
		stm.SetCid(t.Cid)
		stm.Where(transaction.StateEQ(uint8(sphinxproxy.TransactionState_TransactionStateSign)))
	case sphinxproxy.TransactionState_TransactionStateDone,
		sphinxproxy.TransactionState_TransactionStateFail:
		stm.SetExitCode(t.ExitCode)
		stm.Where(transaction.StateEQ(uint8(sphinxproxy.TransactionState_TransactionStateSync)))
	}

	_, err = stm.SetState(uint8(t.State)).
		Save(ctx)
	return err
}

package check

import (
	"errors"

	"github.com/NpoolPlatform/message/npool/sphinxproxy"
)

var (
	// ErrCoinTypeNotSupport ..
	ErrCoinTypeNotSupport = errors.New("coin type not support")
	// ErrTransactionTypeNotSupport ..
	ErrTransactionTypeNotSupport = errors.New("transaction type not support")
)

// TransactionType ..
func TransactionType(transType sphinxproxy.TransactionType) error {
	switch transType {
	case sphinxproxy.TransactionType_WalletNew,
		sphinxproxy.TransactionType_TransactionNew,
		sphinxproxy.TransactionType_Balance:
	default:
		return ErrTransactionTypeNotSupport
	}
	return nil
}

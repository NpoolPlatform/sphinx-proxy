package check

import (
	"errors"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
)

var (
	// ErrCoinTypeNotSupport ..
	ErrCoinTypeNotSupport = errors.New("coin type not support")
	// ErrTransactionTypeNotSupport ..
	ErrTransactionTypeNotSupport = errors.New("transaction type not support")
)

// CoinType ..
func CoinType(coinType sphinxplugin.CoinType) error {
	switch coinType {
	case sphinxplugin.CoinType_CoinTypeBTC, sphinxplugin.CoinType_CoinTypeFIL:
	default:
		return ErrCoinTypeNotSupport
	}
	return nil
}

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

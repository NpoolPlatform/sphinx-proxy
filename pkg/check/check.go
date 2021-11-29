package check

import (
	"errors"
	"strings"

	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/sphinxplugin"
)

var (
	ErrCoinTypeNotSupport        = errors.New("coin type not support")        // nolint
	ErrTransactionTypeNotSupport = errors.New("transaction type not support") // nolint
)

func CoinType(coinType sphinxplugin.CoinType) error {
	switch coinType {
	case sphinxplugin.CoinType_CoinTypeBTC, sphinxplugin.CoinType_CoinTypeFIL:
	default:
		return ErrCoinTypeNotSupport
	}
	return nil
}

func TransactionType(transType signproxy.TransactionType) error {
	switch transType {
	case signproxy.TransactionType_WalletNew,
		signproxy.TransactionType_TransactionNew,
		signproxy.TransactionType_Balance:
	default:
		return ErrTransactionTypeNotSupport
	}
	return nil
}

func TruncateCoinTypePrefix(ct sphinxplugin.CoinType) string {
	return strings.TrimPrefix(ct.String(), "CoinType")
}

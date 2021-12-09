package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
	"github.com/NpoolPlatform/message/npool/sphinxproxy"
)

var (
	// ErrCoinTypeUnKnow ..
	ErrCoinTypeUnKnow = errors.New("coin type unknow")
	// ErrTransactionStateKnow ..
	ErrTransactionStateKnow = errors.New("transaction state unknow")
)

const (
	coinTypePrefix         = "CoinType"
	transactionStatePrefix = "TransactionState"
)

// TruncateCoinTypePrefix ..
func TruncateCoinTypePrefix(ct sphinxplugin.CoinType) string {
	return strings.TrimPrefix(ct.String(), coinTypePrefix)
}

// ToCoinType ..
func ToCoinType(coinType string) (sphinxplugin.CoinType, error) {
	_coinType, ok := sphinxplugin.CoinType_value[fmt.Sprintf("%s%s", coinTypePrefix, coinType)]
	if !ok {
		return sphinxplugin.CoinType_CoinTypeUnKnow, ErrCoinTypeUnKnow
	}
	return sphinxplugin.CoinType(_coinType), nil
}

// TruncateTransactionStatePrefix ..
func TruncateTransactionStatePrefix(ct sphinxproxy.TransactionState) string {
	return strings.TrimPrefix(ct.String(), transactionStatePrefix)
}

// ToTransactionState ..
func ToTransactionState(transactionState string) (sphinxproxy.TransactionState, error) {
	_transactionState, ok := sphinxproxy.TransactionState_value[fmt.Sprintf("%s%s", transactionStatePrefix, transactionState)]
	if !ok {
		return sphinxproxy.TransactionState_TransactionStateUnKnow, ErrTransactionStateKnow
	}
	return sphinxproxy.TransactionState(_transactionState), nil
}

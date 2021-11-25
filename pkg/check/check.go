package check

import (
	"errors"
	"strings"

	"github.com/NpoolPlatform/message/npool/sphinxplugin"
)

var ErrCoinTypeNotSupport = errors.New("coin type not support") // nolint

func CoinType(coinType sphinxplugin.CoinType) error {
	switch coinType {
	case sphinxplugin.CoinType_CoinTypeBTC, sphinxplugin.CoinType_CoinTypeFIL:
	default:
		return ErrCoinTypeNotSupport
	}
	return nil
}

func TruncateCoinTypePrefix(ct sphinxplugin.CoinType) string {
	return strings.TrimPrefix(ct.String(), "CoinType")
}

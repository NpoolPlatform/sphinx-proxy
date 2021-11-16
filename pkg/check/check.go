package check

import (
	"errors"

	"github.com/NpoolPlatform/message/npool/signproxy"
)

var ErrCoinTypeNotSupport = errors.New("coin type not support")

func CoinType(coinType signproxy.CoinType) error {
	switch coinType {
	case signproxy.CoinType_CoinTypeBTC, signproxy.CoinType_CoinTypeFIL:
	default:
		return ErrCoinTypeNotSupport
	}
	return nil
}

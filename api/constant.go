package api

import "fmt"

const (
	PRE = "proxy::%s::nonce::%s"
)

var PreSignKey = func(coinType, transactionID string) string {
	return fmt.Sprintf(PRE,
		coinType,
		transactionID,
	)
}

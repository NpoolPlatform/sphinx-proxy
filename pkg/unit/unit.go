package unit

import (
	"github.com/filecoin-project/lotus/build"
	"github.com/shopspring/decimal"
)

// FIL2AttoFIL not used, at sphinx sign deal
func FIL2AttoFIL(value float64) (float64, bool) {
	return decimal.NewFromFloat(value).
		Mul(decimal.NewFromInt(int64(build.FilecoinPrecision))).
		Float64()
}

func AttoFIL2FIL(value float64) (float64, bool) {
	return decimal.NewFromFloat(value).
		Div(decimal.NewFromInt(int64(build.FilecoinPrecision))).
		Float64()
}

package unit

import (
	"github.com/shopspring/decimal"
)

const (
	FIL      = 1 // nolint
	MilliFIL = FIL * 1000 // nolint
	MicroFIL = MilliFIL * 1000 // nolint
	NanoFIL  = MicroFIL * 1000 // nolint
	PicoFIL  = NanoFIL * 1000 // nolint
	FemtoFIL = PicoFIL * 1000 // nolint
	AttoFIL  = FemtoFIL * 1000 // nolint
)

func FIL2AttoFIL(value float64) string {
	return decimal.NewFromFloat(value).Mul(decimal.NewFromInt(AttoFIL)).String()
}

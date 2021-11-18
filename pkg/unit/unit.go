package unit

import (
	"github.com/shopspring/decimal"
)

const (
	FIL      = 1
	MilliFIL = FIL * 1000
	MicroFIL = MilliFIL * 1000
	NanoFIL  = MicroFIL * 1000
	PicoFIL  = NanoFIL * 1000
	FemtoFIL = PicoFIL * 1000
	AttoFIL  = FemtoFIL * 1000
)

func FIL2AttoFIL(value float64) string {
	return decimal.NewFromFloat(value).Mul(decimal.NewFromInt(AttoFIL)).String()
}

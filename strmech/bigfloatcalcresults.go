package strmech

import "math/big"

type BigFloatCalcStats struct {
	BaseFloatNum big.Float

	BasePrec uint

	BaseNumIntDigits int64

	BaseNumFracDigits int64

	CalcResult big.Float

	CalcResultPrec uint

	CalcResultNumIntDigits int64

	CalcResultNumFracDigits int64
}

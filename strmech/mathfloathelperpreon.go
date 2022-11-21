package strmech

import (
	"math/big"
	"sync"
)

type mathFloatHelperPreon struct {
	lock *sync.Mutex
}

// precisionToDigitsFactor
//
// Returns an instance of *big.Float configured with the
// "Precision To Digits" conversion factor.
//
// Precision bits are used in the configuration of
// big.Float types. The conversion factor is
// "3.3219789....".
//
//		Precision Bits / Conversion Factor =
//				Numeric Digit Capacity
//
//	Conversely:
//
//		Conversion Factor  x  Numeric Digit Capacity =
//				Precision Bits
//
//	Precision, as used in connection with type big.Float,
//	specifies the mantissa precision of a number in bits.
//
//	Also, remember that the number of numeric digits
//	identified using this conversion factor
//	For information on 'precision bits' and their
//	relevance to type big.Float, reference:
//
//	https://pkg.go.dev/math/big#Float
//
//	Bear in mind that this conversion factor may only be
//	used to generate an estimate of numeric digits
//	associated with a give precision bits value. This
//	estimate may vary from the actual number of numeric
//	digits. This estimate has a margin of error of plus
//	or minus five (+ or - 5).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	None
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*big.Float
//
//		This method returns a pointer to an instance of
//		big.Float configured with the conversion factor
//		used to convert precision bits to the number of
//		equivalent numeric digits.
func (floatHelperPreon *mathFloatHelperPreon) precisionToDigitsFactor() *big.Float {

	if floatHelperPreon.lock == nil {
		floatHelperPreon.lock = new(sync.Mutex)
	}

	floatHelperPreon.lock.Lock()

	defer floatHelperPreon.lock.Unlock()

	conversionStrValue := new(MathConstantsFloat).
		PrecisionToDigitsFactorStr()

	precisionToDigitsFactor,
		_ := new(big.Float).
		SetInt64(0).
		SetMode(big.AwayFromZero).SetString(
		conversionStrValue)

	return precisionToDigitsFactor
}

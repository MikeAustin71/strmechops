package strmech

import (
	"math/big"
	"sync"
)

type mathFloatHelperPreon struct {
	lock *sync.Mutex
}

//	estimateDigitsToPrecision
//
//	Computes an estimate of the number of precision
//	bits required in order to store a given number
//	of numeric digits in a type big.Float, floating
//	point number.
//
//	Precision bits are used in the configuration of
//	big.Float types. The conversion factor is
//	"3.3219789132197891321978913219789".
//
//		Conversion Factor  x  Numeric Digit Capacity =
//				Precision Bits
//			(margin of error +/- 16)
//
//	The number of precision bits returned is an
//	estimate with a margin of error of plus or minus
//	sixteen (+ or - 16).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numNumericDigitsRequired	int64
//
//		The number of numeric digits to be stored and
//		processed by a type big.Float floating point
//		numeric value. This value represents the desired
//		capacity for a big.Float number. This number of
//		numeric digits should include both integer and
//		fractional numeric digits.
//
//		If this value is less than one (+1), this
//		method will return a value of zero, thereby
//		signaling an error.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	uint
//
//		If input parameter 'numNumericDigitsRequired'
//		has a value less than one (+1), this parameter
//		will return a value of zero (0) signaling an
//		error.
//
//		Otherwise, the value returned will represent the
//		estimated number of precision bits required for
//		the mantissa of a big.Float value i
//		be stored given the value of input parameter,
//		'precisionBits'. This estimate has a margin of
//		error of plus or minus sixteen bits (+ or - 16).
func (floatHelperPreon *mathFloatHelperPreon) estimateDigitsToPrecision(
	numNumericDigitsRequired int64) uint {

	if floatHelperPreon.lock == nil {
		floatHelperPreon.lock = new(sync.Mutex)
	}

	floatHelperPreon.lock.Lock()

	defer floatHelperPreon.lock.Unlock()

	if numNumericDigitsRequired < 1 {

		return uint(0)
	}

	conversionStrValue := new(MathConstantsFloat).
		PrecisionToDigitsFactorStr()

	precisionToDigitsFactor,
		_ := new(big.Float).
		SetMode(big.AwayFromZero).
		SetString(
			conversionStrValue)

	precisionToDigitsFactor.SetPrec(
		precisionToDigitsFactor.Prec())

	numOfDigitsFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(numNumericDigitsRequired)

	numOfPrecisionBitsFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		Mul(
			numOfDigitsFloat, precisionToDigitsFactor)

	precisionBitsInt64,
		_ :=
		numOfPrecisionBitsFloat.Uint64()

	precisionBitsUint := uint(precisionBitsInt64)

	return precisionBitsUint
}

//	estimatePrecisionToDigits
//
//	Computes an estimates of the number of numerical
//	digits which can be stored given the number of
//	precision bits configured for a type big.Float,
//	floating point number.
//
// Precision bits are used in the configuration of
// big.Float types. The conversion factor is
// "3.3219789132197891321978913219789".
//
//		Precision Bits / Conversion Factor =
//				Numeric Digit Capacity
//			(margin of error +/- 3)
//
//	The number of numerical digits returned is an
//	estimate with a margin of error of plus or minus
//	three (+ or - 3).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	precisionBits				uint
//
//		The number of bits of precision in the mantissa
//		of a big.Float floating point numeric value.
//
//		If this value is less than four (+4), an invalid
//		value of minus one (-1) will be returned.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	int64
//
//		If input parameter 'precisionBits' has a value
//		less than four (+4), this parameter will return
//		a value of minus one (-1) signaling an error.
//
//		Otherwise, the value returned will represent the
//		estimated number of numerical digits which can
//		be stored given the value of input parameter,
//		'precisionBits'. This estimate has a margin of
//		error of plus or minus three numeric digits
//		(+ or - 3).
func (floatHelperPreon *mathFloatHelperPreon) estimatePrecisionToDigits(
	precisionBits uint) int64 {

	if floatHelperPreon.lock == nil {
		floatHelperPreon.lock = new(sync.Mutex)
	}

	floatHelperPreon.lock.Lock()

	defer floatHelperPreon.lock.Unlock()

	if precisionBits < 4 {

		return int64(-1)
	}

	conversionStrValue := new(MathConstantsFloat).
		PrecisionToDigitsFactorStr()

	precisionToDigitsFactor,
		_ := new(big.Float).
		SetMode(big.AwayFromZero).
		SetString(
			conversionStrValue)

	precisionToDigitsFactor.SetPrec(
		precisionToDigitsFactor.Prec())

	precisionBitsFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		SetInt64(int64(precisionBits))

	numOfDigitsFloat := new(big.Float).
		SetMode(big.AwayFromZero).
		Quo(
			precisionBitsFloat, precisionToDigitsFactor)

	numOfDigitsInt64,
		_ :=
		numOfDigitsFloat.Int64()

	return numOfDigitsInt64
}

// precisionToDigitsFactor
//
// Returns an instance of *big.Float configured with the
// "Precision To Digits" conversion factor.
//
// Precision bits are used in the configuration of
// big.Float types. The conversion factor is
// "3.3219789132197891321978913219789".
//
//		Precision Bits / Conversion Factor =
//				Numeric Digit Capacity
//			(margin of error +/- 3)
//
//	Conversely:
//
//		Conversion Factor  x  Numeric Digit Capacity =
//				Precision Bits
//			(margin of error +/- 16)
//
//	Precision, as used in connection with type big.Float,
//	specifies the mantissa precision of a number in bits.
//
//	Also, remember that the number of numeric digits
//	identified using this conversion factor includes
//	both integer and fractional digits.
//
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
//	or minus five (+ or - 3).
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
		SetMode(big.AwayFromZero).
		SetString(
			conversionStrValue)

	return precisionToDigitsFactor
}

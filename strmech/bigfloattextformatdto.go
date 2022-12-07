package strmech

import "math/big"

// BigFloatTextFormatDto
//
// This data transfer object contains the specifications
// necessary to format a big.Float floating point number
// for output as a text string.
type BigFloatTextFormatDto struct {
	BigFloatNum big.Float
	// The big.Float floating point number to
	// be formatted for output as a text string.

	RoundingMode big.RoundingMode
	// The rounding mode used to round 'BigFloatNum'
	// to the number of fractional digits specified
	// by parameter, 'NumOfFractionalDigits'.
	//
	// Rounding Modes are defined in Golang as follows:
	//
	//	ToNearestEven RoundingMode == IEEE 754-2008 roundTiesToEven
	//	ToNearestAway == IEEE 754-2008 roundTiesToAway
	//	ToZero        == IEEE 754-2008 roundTowardZero
	//	AwayFromZero  == no IEEE 754-2008 equivalent
	//	ToNegativeInf == IEEE 754-2008 roundTowardNegative
	//	ToPositiveInf == IEEE 754-2008 roundTowardPositive

	NumOfFractionalDigits int
	// The number of digits to the right of the radix
	// point (a.k.a. decimal point) which will be
	// displayed in the formatted text string for the
	// big.Float floating point number, 'BigFloatNum'.
}

// GetFormattedText
//
// Returns a string containing the formatted text
// representation of the big.Float floating point number
// encapsulated in the current instance of
// BigFloatTextFormatDto.
func (bFloatFmtDto *BigFloatTextFormatDto) GetFormattedText() string {

	bFloatFmtDto.BigFloatNum.SetMode(
		bFloatFmtDto.RoundingMode)

	return bFloatFmtDto.BigFloatNum.Text(
		'f', bFloatFmtDto.NumOfFractionalDigits)
}

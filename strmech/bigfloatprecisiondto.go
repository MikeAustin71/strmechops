package strmech

// BigFloatPrecisionDto
//
// This Data Transfer Object (DTO) is used to transmit
// and store specifications related to the precision bits
// of a big.Float floating point numeric value.
type BigFloatPrecisionDto struct {
	NumIntegerDigits int64
	//	The actual or estimated number of integer digits
	//	in a big.Float floating point numeric value. The
	//	number of integer digits in a floating point
	//	number is one of the elements used to calculate
	//	the precision bits required to store that
	//	floating point number.

	NumFractionalDigits int64
	//	The actual or estimated number of fractional
	//	digits in a big.Float floating point numeric
	//	value. The number of fractional digits in a
	//	floating point number is one of the elements used
	//	to calculate the precision bits required to store
	//	that floating point number.

	NumOfExtraDigitsBuffer int64
	//	When estimating the number of precision necessary
	//	to store or process big.Float floating point
	//	values, is generally a good idea to include a
	//	safety margin consisting of excess numeric digits.
	//
	//	This parameter stores the number of extra numeric
	//	digits used in a calculation of total require
	//	precision bits.

	PrecisionBitsSpec uint
	//	This parameter represents the estimated number of
	//	bits required to store a specific floating point
	//	numeric value in an instance of type big.Float.
	//
	//	The 'PrecisionBitsSpec' value is usually generated
	//	by an internal calculation based on the estimated
	//	number of integer and fractional digits contained
	//	in a big.Float floating point number. However,
	//	users have the option to specify an arbitrary
	//	precision bits value.
}

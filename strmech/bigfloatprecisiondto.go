package strmech

// BigFloatPrecisionDto
//
// This Data Transfer Object (DTO) is used to transmit
// and store specifications related to the precision bits
// of a big.Float floating point numeric value.
type BigFloatPrecisionDto struct {
  NumIntegerDigits uint64
  //	The actual or estimated number of integer digits
  //	in a big.Float floating point numeric value. The
  //	number of integer digits in a floating point
  //	number is one of the elements used to calculate
  //	the precision bits required to store that
  //	floating point number.

  NumFractionalDigits uint64
  //	The actual or estimated number of fractional
  //	digits in a big.Float floating point numeric
  //	value. The number of fractional digits in a
  //	floating point number is one of the elements used
  //	to calculate the precision bits required to store
  //	that floating point number.

  NumOfExtraDigitsBuffer uint64
  //	When estimating the amount of precision necessary
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
  //	numeric value in an instance of the type big.Float.
  //
  //	The 'PrecisionBitsSpec' value is usually generated
  //	by an internal calculation based on the estimated
  //	number of integer and fractional digits contained
  //	in a big.Float floating point number. However,
  //	users can specify an arbitrary precision bits value.
}

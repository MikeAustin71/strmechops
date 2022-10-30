package strmech

import (
	"sync"
)

type BigDecimal struct {
	integerDigits Int8ArrayDto
	//	Contains an array of int8 integers which
	//	comprise the BigDecimal integer digits.

	fractionalDigits Int8ArrayDto
	//	Contains an array of int8 integers which
	//	comprise the BigDecimal fractional digits.

	numberSign NumericSignValueType
	// An enumeration specifying the number sign associated
	// with the numeric value represented by this instance
	// of BigDecimal.  Possible values are listed as
	// follows:
	//      NumSignVal.None() - Invalid Value
	//      NumSignVal.Negative() = -1
	//      NumSignVal.Zero()     =  0
	//      NumSignVal.Positive() =  1

	numberStrFormat NumStrFmtCountryCultureSpec
	//	Required for Number String Formatting. Includes
	//	details about the country and culture associated
	//	with this number string formatting specification.
	//
	//	This number string formatting specification is
	//	used to format the BigDecimal numeric values for
	//	text presentations.

	numericValueType NumericValueType
	// This enumeration value specifies the type of
	// numeric value contained in the current instance
	// of BigDecimal. The contained numeric value
	// is classified either as an integer or a floating
	// point value.
	//
	// Possible enumeration values are listed as
	// follows:
	//  NumValType.None()
	//  NumValType.FloatingPoint()
	//  NumValType.Integer()

	lock *sync.Mutex
}

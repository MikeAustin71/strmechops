package strmech

import "sync"

// NumberStrStatsDto
//
// A data transport type designed to store and transmit
// information on a numeric value.
type NumberStrStatsDto struct {
	NumOfIntegerDigits uint64
	//	The total number of integer digits to the left of
	//	the radix point or, decimal point, in the subject
	//	numeric value.

	NumOfSignificantIntegerDigits uint64
	//	The number of nonzero integer digits to the left
	//	of the radix point, or decimal point, in the
	//	subject numeric value.

	NumOfFractionalDigits uint64
	//	The total number of fractional digits to the
	//	right of the radix point, or decimal point, in
	//	the subject numeric value.

	NumOfSignificantFractionalDigits uint64
	//	The number of nonzero fractional digits to the
	//	right of the radix point, or decimal point, in
	//	the subject numeric value.

	NumberValueType NumericValueType
	//	This enumeration value specifies whether the
	//	subject numeric value is classified either as an
	//	integer or a floating point value.
	//
	//	Possible enumeration values are listed as
	//	follows:
	//  	NumValType.None()
	//  	NumValType.FloatingPoint()
	//  	NumValType.Integer()

	NumberSign NumericSignValueType
	//	An enumeration specifying the number sign
	//	associated with the numeric value. Possible
	//	values are listed as follows:
	//      NumSignVal.None()		= Invalid Value
	//      NumSignVal.Negative()	= -1
	//      NumSignVal.Zero()		=  0
	//      NumSignVal.Positive()	=  1

	IsZeroValue bool
	//	If 'false', the Numeric Value is greater than or
	//	less than zero ('0').
	//
	//	If 'true', the Numeric Value is equal to zero.

	lock *sync.Mutex
}

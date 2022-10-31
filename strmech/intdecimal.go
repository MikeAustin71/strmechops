package strmech

import "sync"

// IntDecimal
//
// This type stores and performs computations on numeric
// values stored as scientific notation using an array
// of integer digits.
//
// The numeric value is computed by multiplying the
// 'significand' by 10 to the power of 'exponent'.
//
//	numeric value = significand x 10^exponent
type IntDecimal struct {
	significand Int8ArrayDto

	exponent Int8ArrayDto

	numberSign NumericSignValueType
	//	An enumeration specifying the number sign associated
	//	with the numeric value represented by the numeric
	//	digits contained within the member variable,
	//	'Int8Array'
	//
	//	Possible values are listed as follows:
	//
	//      NumSignVal.None() - Invalid Value
	//      NumSignVal.Negative() = -1
	//      NumSignVal.Zero()     =  0
	//      NumSignVal.Positive() =  1

	Description1 string
	//	Optional. A name, label or narrative text used to
	//	describe the current instance of Int8ArrayDto.

	Description2 string
	//	Optional. A name, label or narrative text used to
	//	describe the current instance of Int8ArrayDto.

	lock *sync.Mutex
}

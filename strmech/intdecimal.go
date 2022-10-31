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

	Description1 string
	//	Optional. A name, label or narrative text used to
	//	describe the current instance of Int8ArrayDto.

	Description2 string
	//	Optional. A name, label or narrative text used to
	//	describe the current instance of Int8ArrayDto.

	lock *sync.Mutex
}

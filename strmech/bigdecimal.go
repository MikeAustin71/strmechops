package strmech

import (
	"sync"
)

type BigDecimal struct {
	numericValue Int8ArrayDto
	//	Contains an array of int8 integers which
	//	comprise the BigDecimal numeric value.

	numberStrFormat NumStrFmtCountryCultureSpec
	//	Required for Number String Formatting. Includes
	//	details about the country and culture associated
	//	with this number string formatting specification.
	//
	//	This number string formatting specification is
	//	used to format the BigDecimal numeric values for
	//	text presentations.

	lock *sync.Mutex
}

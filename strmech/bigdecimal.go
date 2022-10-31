package strmech

import (
	"math/big"
	"sync"
)

type BigDecimal struct {
	significand *big.Int

	exponent *big.Int

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

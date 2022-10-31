package strmech

import (
	"sync"
)

type BigDecimal struct {
	significand NumberStrKernel

	exponent NumberStrKernel

	lock *sync.Mutex
}

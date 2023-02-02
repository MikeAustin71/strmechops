package strmech

import (
	"math/big"
	"sync"
)

type BigIntNum struct {
	significand big.Int

	exponent big.Int

	lock *sync.Mutex
}

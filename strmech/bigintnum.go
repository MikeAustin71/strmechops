package strmech

import (
	"math/big"
	"sync"
)

type BigIntValue struct {
	significand *big.Int

	exponent *big.Int

	lock *sync.Mutex
}

type BigIntNum struct {
	significand BigIntValue

	exponent BigIntValue

	lock *sync.Mutex
}

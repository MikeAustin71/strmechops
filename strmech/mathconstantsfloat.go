package strmech

import (
	"sync"
)

type MathConstantsFloat struct {
	lock *sync.Mutex
}

// PrecisionToDigitsFactorStr
//
// Returns a string containing the numerical value used
// to convert bits of precision to numerical digits for
// big.Float numbers.
//
// This factor is used when converting precision to
// numerical digits or numerical digits to precision.
//
// The values produced with this constant represent an
// estimate with an error of plus or minus 5.
func (mathConstantsFloat *MathConstantsFloat) PrecisionToDigitsFactorStr() string {

	if mathConstantsFloat.lock == nil {
		mathConstantsFloat.lock = new(sync.Mutex)
	}

	mathConstantsFloat.lock.Lock()

	defer mathConstantsFloat.lock.Unlock()

	return "3.3219789132197891321978913219789"
}

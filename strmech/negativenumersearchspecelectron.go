package strmech

import (
	"sync"
)

type negNumSearchSpecElectron struct {
	lock *sync.Mutex
}

func (negNumSearchElectron *negNumSearchSpecElectron) emptyProcessingFlags(
	negNumSignSpec *NegativeNumberSearchSpec) {

	if negNumSearchElectron.lock == nil {
		negNumSearchElectron.lock = new(sync.Mutex)
	}

	negNumSearchElectron.lock.Lock()

	defer negNumSearchElectron.lock.Unlock()

	if negNumSignSpec == nil {
		return
	}

	negNumSignSpec.foundFirstNumericDigitInNumStr = false
	negNumSignSpec.foundNegNumSignSymbols = false
	negNumSignSpec.foundLeadingNegNumSign = false
	negNumSignSpec.foundLeadingNegNumSignIndex = -1
	negNumSignSpec.foundTrailingNegNumSign = false
	negNumSignSpec.foundTrailingNegNumSignIndex = -1

}

// ptr - Returns a pointer to a new instance of
// negNumSearchSpecElectron.
//
func (negNumSearchElectron negNumSearchSpecElectron) ptr() *negNumSearchSpecElectron {

	if negNumSearchElectron.lock == nil {
		negNumSearchElectron.lock = new(sync.Mutex)
	}

	negNumSearchElectron.lock.Lock()

	defer negNumSearchElectron.lock.Unlock()

	return &negNumSearchSpecElectron{
		lock: new(sync.Mutex),
	}
}

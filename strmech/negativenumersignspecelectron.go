package strmech

import (
	"sync"
)

type negNumSignSpecElectron struct {
	lock *sync.Mutex
}

func (negNumSignElectron *negNumSignSpecElectron) emptyProcessingFlags(
	negNumSignSpec *NegativeNumberSignSpec) {

	if negNumSignElectron.lock == nil {
		negNumSignElectron.lock = new(sync.Mutex)
	}

	negNumSignElectron.lock.Lock()

	defer negNumSignElectron.lock.Unlock()

	if negNumSignSpec == nil {
		return
	}

	negNumSignSpec.negNumSignTargetSearchChars = nil
	negNumSignSpec.foundFirstNumericDigitInNumStr = false
	negNumSignSpec.foundNegNumSignSymbols = false
	negNumSignSpec.foundLeadingNegNumSign = false
	negNumSignSpec.foundLeadingNegNumSignIndex = -1
	negNumSignSpec.foundTrailingNegNumSign = false
	negNumSignSpec.foundTrailingNegNumSignIndex = -1

}

// ptr - Returns a pointer to a new instance of
// negNumSignSpecElectron.
//
func (negNumSignElectron negNumSignSpecElectron) ptr() *negNumSignSpecElectron {

	if negNumSignElectron.lock == nil {
		negNumSignElectron.lock = new(sync.Mutex)
	}

	negNumSignElectron.lock.Lock()

	defer negNumSignElectron.lock.Unlock()

	return &negNumSignSpecElectron{
		lock: new(sync.Mutex),
	}
}

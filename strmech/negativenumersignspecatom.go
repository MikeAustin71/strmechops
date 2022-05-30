package strmech

import "sync"

type negNumSignSpecAtom struct {
	lock *sync.Mutex
}

func (negNumSignAtom *negNumSignSpecAtom) empty(
	negNumSignSpec *NegativeNumberSignSpec) {

	if negNumSignAtom.lock == nil {
		negNumSignAtom.lock = new(sync.Mutex)
	}

	negNumSignAtom.lock.Lock()

	defer negNumSignAtom.lock.Unlock()

	if negNumSignSpec == nil {
		return
	}

	negNumSignSpec.negNumSignPosition = NSignSymPos.None()
	negNumSignSpec.leadingNegNumSignSymbols = nil
	negNumSignSpec.trailingNegNumSignSymbols = nil

	negNumSignSpecElectron{}.ptr().
		emptyProcessingFlags(negNumSignSpec)

	return
}

// ptr - Returns a pointer to a new instance of
// negNumSignSpecAtom.
//
func (negNumSignAtom negNumSignSpecAtom) ptr() *negNumSignSpecAtom {

	if negNumSignAtom.lock == nil {
		negNumSignAtom.lock = new(sync.Mutex)
	}

	negNumSignAtom.lock.Lock()

	defer negNumSignAtom.lock.Unlock()

	return &negNumSignSpecAtom{
		lock: new(sync.Mutex),
	}
}

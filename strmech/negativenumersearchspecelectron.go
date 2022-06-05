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

// equal - Receives a pointer to two instances of
// NegativeNumberSearchSpec and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (negNumSearchElectron *negNumSearchSpecElectron) equal(
	negNumSearchSpec01 *NegativeNumberSearchSpec,
	negNumSearchSpec02 *NegativeNumberSearchSpec) bool {

	if negNumSearchElectron.lock == nil {
		negNumSearchElectron.lock = new(sync.Mutex)
	}

	negNumSearchElectron.lock.Lock()

	defer negNumSearchElectron.lock.Unlock()

	if negNumSearchSpec01 == nil ||
		negNumSearchSpec02 == nil {

		return false
	}

	if negNumSearchSpec01.negNumSignPosition !=
		negNumSearchSpec02.negNumSignPosition {

		return false
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		negNumSearchSpec01.leadingNegNumSignSymbols,
		negNumSearchSpec02.leadingNegNumSignSymbols)

	if !areEqual {
		return false
	}

	areEqual = sMechPreon.equalRuneArrays(
		negNumSearchSpec01.trailingNegNumSignSymbols,
		negNumSearchSpec02.trailingNegNumSignSymbols)

	if !areEqual {
		return false
	}

	if negNumSearchSpec01.foundFirstNumericDigitInNumStr !=
		negNumSearchSpec02.foundFirstNumericDigitInNumStr {

		return false
	}

	if negNumSearchSpec01.foundNegNumSignSymbols !=
		negNumSearchSpec02.foundNegNumSignSymbols {

		return false
	}

	if negNumSearchSpec01.foundLeadingNegNumSign !=
		negNumSearchSpec02.foundLeadingNegNumSign {

		return false
	}

	if negNumSearchSpec01.foundLeadingNegNumSignIndex !=
		negNumSearchSpec02.foundLeadingNegNumSignIndex {

		return false
	}

	if negNumSearchSpec01.foundTrailingNegNumSign !=
		negNumSearchSpec02.foundTrailingNegNumSign {

		return false
	}

	if negNumSearchSpec01.foundTrailingNegNumSignIndex !=
		negNumSearchSpec02.foundTrailingNegNumSignIndex {

		return false
	}

	return true
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

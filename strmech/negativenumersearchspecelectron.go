package strmech

import (
	"sync"
)

type negNumSearchSpecElectron struct {
	lock *sync.Mutex
}

// emptyProcessingFlags - Sets NegativeNumberSearchSpec internal
// member variables used as Processing Flags to their initial or
// zero values.
//
// Internal Processing flags are used by Number String parsing
// functions to identify Negative Number Sign Symbols in strings of
// numeric digits called 'Number Strings'. Number String parsing
// functions review strings of text characters containing numeric
// digits and convert those numeric digits to numeric values.
//
func (negNumSearchElectron *negNumSearchSpecElectron) emptyProcessingFlags(
	negNumSearchSpec *NegativeNumberSearchSpec) {

	if negNumSearchElectron.lock == nil {
		negNumSearchElectron.lock = new(sync.Mutex)
	}

	negNumSearchElectron.lock.Lock()

	defer negNumSearchElectron.lock.Unlock()

	if negNumSearchSpec == nil {
		return
	}

	negNumSearchSpec.foundFirstNumericDigitInNumStr = false
	negNumSearchSpec.foundNegNumSignSymbols = false
	negNumSearchSpec.foundLeadingNegNumSign = false
	negNumSearchSpec.foundLeadingNegNumSignIndex = -1
	negNumSearchSpec.foundTrailingNegNumSign = false
	negNumSearchSpec.foundTrailingNegNumSignIndex = -1

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
		negNumSearchSpec01.leadingNegNumSignSymbols.CharsArray,
		negNumSearchSpec02.leadingNegNumSignSymbols.CharsArray)

	if !areEqual {
		return false
	}

	areEqual = sMechPreon.equalRuneArrays(
		negNumSearchSpec01.trailingNegNumSignSymbols.CharsArray,
		negNumSearchSpec02.trailingNegNumSignSymbols.CharsArray)

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

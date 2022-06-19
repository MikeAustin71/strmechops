package strmech

import "sync"

// decimalSepSpecElectron - Provides helper methods for type
// DecimalSeparatorSpec.
//
type decimalSepSpecElectron struct {
	lock *sync.Mutex
}

// emptyProcessingFlags - Sets NegativeNumberSearchSpec internal
// member variables used as Processing Flags to their initial or
// zero values.
//
// Internal Processing flags are used by Number String parsing
// functions to identify decimal separators in strings of numeric
// digits called 'Number Strings'. Number String parsing functions
// review strings of text characters containing numeric digits and
// convert those numeric digits to numeric values.
//
// Decimal Separators are comprised of a text character or
// characters and are used to separate integer digits from
// fractional digits in floating point numeric values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (decSepSpecElectron *decimalSepSpecElectron) emptyProcessingFlags(
	decSepSpec *DecimalSeparatorSpec) {

	if decSepSpecElectron.lock == nil {
		decSepSpecElectron.lock = new(sync.Mutex)
	}

	decSepSpecElectron.lock.Lock()

	defer decSepSpecElectron.lock.Unlock()

	decSepSpec.foundDecimalSeparatorSymbols = false

	decSepSpec.foundDecimalSeparatorIndex = -1

}

// equal - Receives a pointer to two instances of
// DecimalSeparatorSpec and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (decSepSpecElectron *decimalSepSpecElectron) equal(
	decSepSpec01 *DecimalSeparatorSpec,
	decSepSpec02 *DecimalSeparatorSpec) (
	areEqual bool) {

	if decSepSpecElectron.lock == nil {
		decSepSpecElectron.lock = new(sync.Mutex)
	}

	decSepSpecElectron.lock.Lock()

	defer decSepSpecElectron.lock.Unlock()

	areEqual = false

	if decSepSpec01 == nil ||
		decSepSpec02 == nil {

		return areEqual
	}

	areEqual = strMechPreon{}.ptr().
		equalRuneArrays(
			decSepSpec01.decimalSeparatorChars.CharsArray,
			decSepSpec02.decimalSeparatorChars.CharsArray)

	if !areEqual {
		return areEqual
	}

	areEqual = false

	if decSepSpec01.foundDecimalSeparatorSymbols !=
		decSepSpec02.foundDecimalSeparatorSymbols {

		return areEqual
	}

	if decSepSpec01.foundDecimalSeparatorIndex !=
		decSepSpec02.foundDecimalSeparatorIndex {

		return areEqual
	}

	areEqual = true

	return areEqual
}

// ptr - Returns a pointer to a new instance of
// decimalSepSpecElectron.
//
func (decSepSpecElectron decimalSepSpecElectron) ptr() *decimalSepSpecElectron {

	if decSepSpecElectron.lock == nil {
		decSepSpecElectron.lock = new(sync.Mutex)
	}

	decSepSpecElectron.lock.Lock()

	defer decSepSpecElectron.lock.Unlock()

	return &decimalSepSpecElectron{
		lock: new(sync.Mutex),
	}
}

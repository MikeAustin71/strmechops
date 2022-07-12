package strmech

import "sync"

// decimalSepSpecElectron - Provides helper methods for type
// DecimalSeparatorSpec.
//
type decimalSepSpecElectron struct {
	lock *sync.Mutex
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

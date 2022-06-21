package strmech

import "sync"

// numberStrKernelElectron - Provides helper methods for type
// NumberStrKernel.
//
type numberStrKernelElectron struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NumberStrKernel and proceeds to reset the data values
// for member variables to their initial or zero values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'numStrKernel' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrKernel               *NumberStrKernel
//     - A pointer to an instance of NumberStrKernel. All
//       the internal member variables contained in this instance
//       will be deleted and reset to their zero values.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (numStrKernelElectron *numberStrKernelElectron) empty(
	numStrKernel *NumberStrKernel) {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	if numStrKernel == nil {
		return
	}

	numStrKernel.integerDigits.Empty()

	numStrKernel.fractionalDigits.Empty()

	numStrKernel.numericValueType =
		NumValType.None()

	numStrKernel.numberSign = NumSignVal.None()

	numStrKernel.isNonZeroValue = false
}

// equal - Receives a pointer to two instances of
// NumberStrKernel and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (numStrKernelElectron *numberStrKernelElectron) equal(
	numStrKernel1 *NumberStrKernel,
	numStrKernel2 *NumberStrKernel) bool {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	if numStrKernel1 == nil ||
		numStrKernel2 == nil {

		return false
	}

	if !numStrKernel1.integerDigits.Equal(
		&numStrKernel2.integerDigits) {

		return false
	}

	if !numStrKernel1.fractionalDigits.Equal(
		&numStrKernel2.fractionalDigits) {

		return false
	}

	if numStrKernel1.numericValueType !=
		numStrKernel2.numericValueType {

		return false
	}

	if numStrKernel1.numberSign !=
		numStrKernel2.numberSign {

		return false
	}

	if numStrKernel1.isNonZeroValue !=
		numStrKernel2.isNonZeroValue {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelElectron.
//
func (numStrKernelElectron numberStrKernelElectron) ptr() *numberStrKernelElectron {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	return &numberStrKernelElectron{
		lock: new(sync.Mutex),
	}
}

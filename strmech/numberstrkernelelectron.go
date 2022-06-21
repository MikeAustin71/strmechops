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

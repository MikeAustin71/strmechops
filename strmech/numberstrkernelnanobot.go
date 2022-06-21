package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelNanobot - Provides helper methods for type
// NumberStrKernel.
//
type numberStrKernelNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingNumStrKernel' to input parameter 'targetNumStrKernel'.
// Both instances are of type NumberStrKernel.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'targetNumStrKernel' will be
// overwritten.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetNumStrKernel        *NumberStrKernel
//     - A pointer to a NumberStrKernel instance. All the
//       member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'incomingNumStrKernel'.
//
//       'targetNumStrKernel' is the target of this copy
//       operation.
//
//
//  incomingNumStrKernel      *NumberStrKernel
//     - A pointer to another NumberStrKernel instance. All
//       the member variable data values from this object will
//       be copied to corresponding member variables in
//       'targetNumStrKernel'.
//
//       'incomingNumStrKernel' is the source for this copy
//       operation.
//
//       If 'incomingNumStrKernel' is determined to be invalid,
//       an error will be returned.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (numStrKernelNanobot *numberStrKernelNanobot) copyIn(
	targetNumStrKernel *NumberStrKernel,
	incomingNumStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetNumStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingNumStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	numStrKernelAtom := numberStrKernelAtom{}

	_,
		err2 =
		numStrKernelAtom.testValidityOfNumStrKernel(
			incomingNumStrKernel,
			nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'incomingNumStrKernel' failed!\n"+
			"This instance of NumberStrKernel is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	numStrKernelElectron := numberStrKernelElectron{}

	numStrKernelElectron.empty(
		targetNumStrKernel)

	err = targetNumStrKernel.integerDigits.CopyIn(
		&incomingNumStrKernel.integerDigits,
		ePrefix.XCpy(
			"targetNumStrKernel.integerDigits"+
				"<-incomingNumStrKernel.integerDigits"))

	if err != nil {
		return err
	}

	err = targetNumStrKernel.fractionalDigits.CopyIn(
		&incomingNumStrKernel.fractionalDigits,
		ePrefix.XCpy(
			"targetNumStrKernel.fractionalDigits"+
				"<-incomingNumStrKernel.fractionalDigits"))

	if err != nil {
		return err
	}

	targetNumStrKernel.numericValueType =
		incomingNumStrKernel.numericValueType

	targetNumStrKernel.numberSign =
		incomingNumStrKernel.numberSign

	targetNumStrKernel.isNonZeroValue =
		incomingNumStrKernel.isNonZeroValue

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'numStrKernel'. a pointer to an instance of
// NumberStrKernel.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The input parameter 'numStrKernel' is determined to be invalid,
// this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrKernel           *NumberStrKernel
//     - A pointer to an instance of NumberStrKernel. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of NumberStrKernel.
//
//       If the member variable data values encapsulated by
//       'numStrKernel' are found to be invalid, this method will
//       return an error
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  deepCopyNumStrKernel       NumberStrKernel
//     - If this method completes successfully, a deep copy of
//       input parameter 'numStrKernel' will be created and
//       returned in a new instance of NumberStrKernel.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (numStrKernelNanobot *numberStrKernelNanobot) copyOut(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyNumStrKernel NumberStrKernel,
	err error) {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	numberStrKernelElectron{}.ptr().
		empty(&deepCopyNumStrKernel)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyNumStrKernel, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyNumStrKernel, err
	}

	var err2 error

	_,
		err2 = numberStrKernelAtom{}.ptr().
		testValidityOfNumStrKernel(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'numStrKernel' failed!\n"+
			"This instance of NumberStrKernel is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return deepCopyNumStrKernel, err
	}

	err = deepCopyNumStrKernel.integerDigits.CopyIn(
		&numStrKernel.integerDigits,
		ePrefix.XCpy(
			"deepCopyNumStrKernel.integerDigits"+
				"<-numStrKernel.integerDigits"))

	if err != nil {
		return deepCopyNumStrKernel, err
	}

	err = deepCopyNumStrKernel.fractionalDigits.CopyIn(
		&numStrKernel.fractionalDigits,
		ePrefix.XCpy(
			"deepCopyNumStrKernel.fractionalDigits"+
				"<-numStrKernel.fractionalDigits"))

	if err != nil {
		return deepCopyNumStrKernel, err
	}

	deepCopyNumStrKernel.numericValueType =
		numStrKernel.numericValueType

	deepCopyNumStrKernel.numberSign =
		numStrKernel.numberSign

	deepCopyNumStrKernel.isNonZeroValue =
		numStrKernel.isNonZeroValue

	return deepCopyNumStrKernel, err
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelNanobot.
//
func (numStrKernelNanobot numberStrKernelNanobot) ptr() *numberStrKernelNanobot {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	return &numberStrKernelNanobot{
		lock: new(sync.Mutex),
	}
}

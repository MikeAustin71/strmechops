package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// decimalSepSpecNanobot - Provides helper methods for type
// DecimalSeparatorSpec.
//
type decimalSepSpecNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingDecSepSpec' to input parameter
// 'targetDecSepSpec'. Both instances are of type
// DecimalSeparatorSpec.
//
// IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'targetDecSepSpec' will
// be deleted and overwritten.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetDecSepSpec           *DecimalSeparatorSpec
//     - A pointer to a DecimalSeparatorSpec instance. All the
//       member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'incomingDecSepSpec'.
//
//       'targetDecSepSpec' is the target of this copy
//       operation.
//
//
//  incomingDecSepSpec         *DecimalSeparatorSpec
//     - A pointer to another DecimalSeparatorSpec instance. All
//       the member variable data values from this object will
//       be copied to corresponding member variables in
//       'targetDecSepSpec'.
//
//       'incomingDecSepSpec' is the source for this copy
//       operation.
//
//       If 'incomingDecSepSpec' is determined to be invalid,
//       an error will be returned.
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
func (decSepSpecNanobot *decimalSepSpecNanobot) copyIn(
	targetDecSepSpec *DecimalSeparatorSpec,
	incomingDecSepSpec *DecimalSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if decSepSpecNanobot.lock == nil {
		decSepSpecNanobot.lock = new(sync.Mutex)
	}

	decSepSpecNanobot.lock.Lock()

	defer decSepSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"decimalSepSpecNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetDecSepSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingDecSepSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingNegNumSearchSpec' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	decSepSpecAtom := decimalSeparatorSpecAtom{}

	_,
		err2 = decSepSpecAtom.testValidityOfDecSepSearchSpec(
		incomingDecSepSpec,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'incomingDecSepSpec' failed!\n"+
			"This instance of DecimalSeparatorSpec is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	decSepSpecAtom.empty(
		targetDecSepSpec)

	lenDecSepChars :=
		len(incomingDecSepSpec.decimalSeparatorChars.CharsArray)

	if lenDecSepChars == 0 {

		incomingDecSepSpec.decimalSeparatorChars.Empty()
		targetDecSepSpec.decimalSeparatorChars.Empty()

	} else {

		err = targetDecSepSpec.decimalSeparatorChars.CopyIn(
			&incomingDecSepSpec.decimalSeparatorChars,
			ePrefix.XCpy(
				"targetDecSepSpec<-incomingDecSepSpec"))
	}

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'depSecSpec', a pointer to an instance of
// DecimalSeparatorSpec.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The input parameter 'depSecSpec' is determined to be
// invalid, this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  depSecSpec           *DecimalSeparatorSpec
//     - A pointer to an instance of DecimalSeparatorSpec. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of DecimalSeparatorSpec.
//
//       If the member variable data values encapsulated by
//       'depSecSpec' are found to be invalid, this method will
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
//  copyOfNegNumSearchSpec     DecimalSeparatorSpec
//     - If this method completes successfully, a deep copy of
//       input parameter 'depSecSpec' will be created and returned
//       in a new instance of DecimalSeparatorSpec.
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
func (decSepSpecNanobot *decimalSepSpecNanobot) copyOut(
	decSepSpec *DecimalSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	copyOfDecSepSpec DecimalSeparatorSpec,
	err error) {

	if decSepSpecNanobot.lock == nil {
		decSepSpecNanobot.lock = new(sync.Mutex)
	}

	decSepSpecNanobot.lock.Lock()

	defer decSepSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"decimalSepSpecNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return copyOfDecSepSpec, err

	}

	if decSepSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decSepSpec' is a nil pointer!\n",
			ePrefix.String())

		return copyOfDecSepSpec, err
	}

	var err2 error

	decSepSpecAtom := decimalSeparatorSpecAtom{}

	_,
		err2 = decSepSpecAtom.testValidityOfDecSepSearchSpec(
		decSepSpec,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'decSepSpec' failed!\n"+
			"This instance of DecimalSeparatorSpec is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return copyOfDecSepSpec, err
	}

	decSepSpecAtom.empty(
		&copyOfDecSepSpec)

	lenDecSepChars :=
		len(decSepSpec.decimalSeparatorChars.CharsArray)

	if lenDecSepChars == 0 {
		decSepSpec.decimalSeparatorChars.Empty()
		copyOfDecSepSpec.decimalSeparatorChars.Empty()

	} else {

		copyOfDecSepSpec.decimalSeparatorChars,
			err =
			decSepSpec.decimalSeparatorChars.CopyOut(
				ePrefix.XCpy(
					"copyOfDecSepSpec.decimalSeparatorChars<-" +
						"decSepSpec.decimalSeparatorChars"))
	}

	return copyOfDecSepSpec, err
}

// ptr - Returns a pointer to a new instance of
// decimalSepSpecNanobot.
//
func (decSepSpecNanobot decimalSepSpecNanobot) ptr() *decimalSepSpecNanobot {

	if decSepSpecNanobot.lock == nil {
		decSepSpecNanobot.lock = new(sync.Mutex)
	}

	decSepSpecNanobot.lock.Lock()

	defer decSepSpecNanobot.lock.Unlock()

	return &decimalSepSpecNanobot{
		lock: new(sync.Mutex),
	}
}

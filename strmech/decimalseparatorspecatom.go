package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// decimalSeparatorSpecAtom - Provides helper methods for type
// DecimalSeparatorSpec.
//
type decimalSeparatorSpecAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// decimalSeparatorSpecAtom.
//
func (decSepSpecAtom decimalSeparatorSpecAtom) ptr() *decimalSeparatorSpecAtom {

	if decSepSpecAtom.lock == nil {
		decSepSpecAtom.lock = new(sync.Mutex)
	}

	decSepSpecAtom.lock.Lock()

	defer decSepSpecAtom.lock.Unlock()

	return &decimalSeparatorSpecAtom{
		lock: new(sync.Mutex),
	}
}

// empty - Receives a pointer to an instance of
// DecimalSeparatorSpec and proceeds to reset the data values for
// member values to their initial or zero values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'decSepSpec' will be deleted and reset to their zero values.
//
// All member variables containing zero based index values will be
// set to minus one (-1).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  decSepSpec           *DecimalSeparatorSpec
//     - A pointer to an instance of DecimalSeparatorSpec. All
//       the internal member variables contained in this instance
//       will be deleted and reset to their zero values. All member
//       variables containing zero based index values will be set
//       to minus one (-1).
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (decSepSpecAtom *decimalSeparatorSpecAtom) empty(
	decSepSpec *DecimalSeparatorSpec) {

	if decSepSpecAtom.lock == nil {
		decSepSpecAtom.lock = new(sync.Mutex)
	}

	decSepSpecAtom.lock.Lock()

	defer decSepSpecAtom.lock.Unlock()

	if decSepSpec == nil {
		return
	}

	decSepSpec.decimalSeparatorChars = nil

	decSepSpec.foundFirstNumericDigitInNumStr = false

	decSepSpec.foundDecimalSeparatorSymbols = false

	decSepSpec.foundDecimalSeparatorIndex = -1

	return
}

// testValidityOfDepSecSpec - Receives a pointer to an
// instance of DecimalSeparatorSpec and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'depSecSpec' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'depSecSpec' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  decSepSpec                 *DecimalSeparatorSpec
//     - A pointer to an instance of DecimalSeparatorSpec. This
//       object will be subjected to diagnostic analysis in order
//       to determine if all the member variables contain valid
//       values.
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
//  isValid                    bool
//     - If input parameter 'depSecSpec' is judged to be valid in
//       all respects, this return parameter will be set to 'true'.
//
//     - If input parameter 'depSecSpec' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'depSecSpec' is judged to be valid in
//       all respects, this return parameter will be set to'nil'.
//
//       If input parameter, 'depSecSpec' is found to be invalid,
//       this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (decSepSpecAtom *decimalSeparatorSpecAtom) testValidityOfDecSepSearchSpec(
	decSepSpec *DecimalSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if decSepSpecAtom.lock == nil {
		decSepSpecAtom.lock = new(sync.Mutex)
	}

	decSepSpecAtom.lock.Lock()

	defer decSepSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"testValidityOfNegNumSearchSpec()",
		"")

	if err != nil {

		return isValid, err
	}

	if decSepSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decSepSpec' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if len(decSepSpec.decimalSeparatorChars) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'DecimalSeparatorSpec' is invalid!\n"+
			"The length of the Decimal Separators Character Array is zero.\n"+
			"len(decSepSpec.decimalSeparatorChars) == 0\n",
			ePrefix.String())

		return isValid, err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		decSepSpec.decimalSeparatorChars,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'DecimalSeparatorSpec' is invalid!\n"+
			"The Decimal Separators Character Array contains invalid characters.\n"+
			"Internal member variable 'decSepSpec.decimalSeparatorChars' returned\n"+
			"the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return isValid, err
	}

	return isValid, err
}

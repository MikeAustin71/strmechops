package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type negNumSearchSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NegativeNumberSearchSpec and proceeds to reset the data values
// for member values to their initial or zero values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'negNumSearchSpec' will be deleted and reset to their zero
// values.
//
// All member variables containing zero based index values will be
// set to minus one (-1).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchSpec           *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. All
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
func (negNumSearchAtom *negNumSearchSpecAtom) empty(
	negNumSearchSpec *NegativeNumberSearchSpec) {

	if negNumSearchAtom.lock == nil {
		negNumSearchAtom.lock = new(sync.Mutex)
	}

	negNumSearchAtom.lock.Lock()

	defer negNumSearchAtom.lock.Unlock()

	if negNumSearchSpec == nil {
		return
	}

	negNumSearchSpec.negNumSignPosition = NumSignSymPos.None()
	negNumSearchSpec.leadingNegNumSignSymbols.Empty()
	negNumSearchSpec.trailingNegNumSignSymbols.Empty()

	negNumSearchSpecElectron{}.ptr().
		emptyProcessingFlags(negNumSearchSpec)

	return
}

// ptr - Returns a pointer to a new instance of
// negNumSearchSpecAtom.
//
func (negNumSearchAtom negNumSearchSpecAtom) ptr() *negNumSearchSpecAtom {

	if negNumSearchAtom.lock == nil {
		negNumSearchAtom.lock = new(sync.Mutex)
	}

	negNumSearchAtom.lock.Lock()

	defer negNumSearchAtom.lock.Unlock()

	return &negNumSearchSpecAtom{
		lock: new(sync.Mutex),
	}
}

// testValidityOfNegNumSearchSpec - Receives a pointer to an
// instance of NegativeNumberSearchSpec and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'negNumSignSpec' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'negNumSignSpec' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchSpec           *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
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
//     - If input parameter 'negNumSignSpec' is judged to be valid
//       in all respects, this return parameter will be set to
//       'true'.
//
//     - If input parameter 'negNumSignSpec' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'negNumSignSpec' is judged to be valid
//       in all respects, this return parameter will be set to
//       'nil'.
//
//       If input parameter, 'negNumSignSpec' is found to be
//       invalid, this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (negNumSearchAtom *negNumSearchSpecAtom) testValidityOfNegNumSearchSpec(
	negNumSearchSpec *NegativeNumberSearchSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if negNumSearchAtom.lock == nil {
		negNumSearchAtom.lock = new(sync.Mutex)
	}

	negNumSearchAtom.lock.Lock()

	defer negNumSearchAtom.lock.Unlock()

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

	if negNumSearchSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if !negNumSearchSpec.negNumSignPosition.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
			"The internal member variable 'negNumSignPosition' is NOT configured.\n"+
			"negNumSignPosition = %v\n",
			ePrefix.String(),
			negNumSearchSpec.negNumSignPosition.String())

		return isValid, err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	if negNumSearchSpec.negNumSignPosition == NumSignSymPos.Before() {

		if len(negNumSearchSpec.trailingNegNumSignSymbols.CharsArray) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, it contains Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return isValid, err

		}

		if len(negNumSearchSpec.leadingNegNumSignSymbols.CharsArray) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, no Leading Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return isValid, err
		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.leadingNegNumSignSymbols.CharsArray,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign Symbol.\n"+
				"Internal member variable 'leadingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return isValid, err
		}

		err = negNumSearchSpec.leadingNegNumSignSymbols.
			SetCharacterSearchType(
				CharSearchType.LinearTargetStartingIndex(),
				ePrefix.XCpy(
					"negNumSearchSpec.leadingNegNumSignSymbols"))

		if err != nil {

			return isValid, err
		}

		isValid = true

		return isValid, err
	}

	if negNumSearchSpec.negNumSignPosition == NumSignSymPos.After() {

		if len(negNumSearchSpec.leadingNegNumSignSymbols.CharsArray) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return isValid, err

		}

		if len(negNumSearchSpec.trailingNegNumSignSymbols.CharsArray) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, no Trailing Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return isValid, err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.trailingNegNumSignSymbols.CharsArray,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'trailingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return isValid, err
		}

		err = negNumSearchSpec.trailingNegNumSignSymbols.
			SetCharacterSearchType(
				CharSearchType.LinearTargetStartingIndex(),
				ePrefix.XCpy(
					"negNumSearchSpec.trailingNegNumSignSymbols"))

		if err != nil {

			return isValid, err
		}

		isValid = true

		return isValid, err
	}

	if negNumSearchSpec.negNumSignPosition == NumSignSymPos.BeforeAndAfter() {

		if len(negNumSearchSpec.leadingNegNumSignSymbols.CharsArray) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign.\n"+
				"However, it contains NO Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return isValid, err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.leadingNegNumSignSymbols.CharsArray,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'leadingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return isValid, err

		}

		if len(negNumSearchSpec.trailingNegNumSignSymbols.CharsArray) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains NO Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return isValid, err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.trailingNegNumSignSymbols.CharsArray,
			nil)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign Symbol.\n"+
				"Internal member variable 'trailingNegNumSignSymbols' returned\n"+
				"the following validation error:\n"+
				"%v\n",
				ePrefix.String(),
				err2.Error())

			return isValid, err
		}

		err = negNumSearchSpec.leadingNegNumSignSymbols.
			SetCharacterSearchType(
				CharSearchType.LinearTargetStartingIndex(),
				ePrefix.XCpy(
					"negNumSearchSpec.leadingNegNumSignSymbols"))

		if err != nil {

			return isValid, err
		}

		err = negNumSearchSpec.trailingNegNumSignSymbols.
			SetCharacterSearchType(
				CharSearchType.LinearTargetStartingIndex(),
				ePrefix.XCpy(
					"negNumSearchSpec.leadingNegNumSignSymbols"))

		if err != nil {

			return isValid, err
		}

		isValid = true

		return isValid, err

	}

	// NegativeNumberSearchSpec is invalid!
	return isValid, err
}

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
// 'negNumSearchSpec' will be deleted and reset to their zero values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchSpec           *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. All
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

	negNumSearchSpec.negNumSignPosition = NSignSymPos.None()
	negNumSearchSpec.leadingNegNumSignSymbols = nil
	negNumSearchSpec.trailingNegNumSignSymbols = nil

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

// leadingNegSignSymSearch - Performs a search for Leading Negative
// Number Sign Symbols in a number string. The text characters to
// be searched were previously configured and saved in internal
// member variable,
//   'NegativeNumberSearchSpec.negNumSignTargetSearchChars'
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSignSpec                  *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
//       instance will be configured as a Leading and Trailing
//       Negative Number Sign Specification. All previous
//       configuration data will be deleted and replaced with a new
//       Leading and Trailing Negative Number Sign configuration.
//
//  foundFirstNumericDigitInNumStr  bool
//     - This boolean value serves as a status flag signaling
//       whether the first numeric digit in a number string
//       has already been located. If this is set to 'true',
//       this method will take no action and exit without
//       error. The reason for this behavior is based on the
//       logic:
//          Once the first Numeric Digit is discovered in
//          a Number String parsing operation, it is
//          impossible to find valid Leading Negative
//          Number Symbols. By definition, Leading Negative
//          Number Symbols always occur before the first
//          numeric digit in a number string.
//
//
//  startingSearchIndex             int
//        The target search characters which will be searched for
//        Leading Negative Number Symbols must be previously
//        configured and saved in the NegativeNumberSearchSpec
//        internal member variable 'negNumSignTargetSearchChars'.
//
//        The 'startingSearchIndex' parameter specifies the index
//        in the Target Search Characters array at which the search
//        for Leading Negative Number Symbols will commence.
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
// -----------------------------------------------------------------
//
// Return Values
//
//  foundNegNumSignSymbols          bool
//     - If this method completes successfully, this parameter will
//       signal whether the search for Leading Negative Number
//       Symbols was successful.
//
//       A return value of 'false' signals that the search for
//       Leading Negative Number Symbols was unsuccessful and
//       the Symbols were NOT located in the Target Search
//       Characters.
//
//       A return value of 'true' signals that the Leading
//       Negative Number Symbols were located in the Target Search
//       Characters and the search was therefore successful.
//
//
//  lastIndex                       int
//       If the search for Leading Negative Number Symbols was
//       unsuccessful, the value of 'lastIndex' will be set to
//       'startingSearchIndex'.
//
//       However, if the Leading Negative Number Symbols were
//       located in the Target Search Characters, the value of
//       'lastIndex' is set to last index of the Leading Negative
//       Number Symbols.
//
//
//  err                             error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchAtom *negNumSearchSpecAtom) leadingNegSignSymSearch(
	negNumSearchSpec *NegativeNumberSearchSpec,
	targetSearchString *TargetSearchStringDto,
	foundFirstNumericDigitInNumStr bool,
	startingSearchIndex int,
	errPrefDto *ePref.ErrPrefixDto) (
	foundNegNumSignSymbols bool,
	lastIndex int,
	err error) {

	if negNumSearchAtom.lock == nil {
		negNumSearchAtom.lock = new(sync.Mutex)
	}

	negNumSearchAtom.lock.Lock()

	defer negNumSearchAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundNegNumSignSymbols = false

	// This assumes startingSearchIndex has
	// already been validated
	lastIndex = startingSearchIndex

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecElectron."+
			"leadingNegSignSymSearch()",
		"")

	if err != nil {

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	if negNumSearchSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	if targetSearchString == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSearchString' is a nil pointer!\n",
			ePrefix.String())

		return foundNegNumSignSymbols,
			lastIndex,
			err

	}

	if negNumSearchSpec.foundFirstNumericDigitInNumStr {

		// Nothing to do. Already found the first
		// numeric digit. Further, Leading Neg Num Sign
		// search is pointless.
		//
		// Once the first Numeric Digit is discovered in
		// a Number String parsing operation, it is
		// impossible to find valid Leading Negative
		// Number Symbols.
		return foundNegNumSignSymbols,
			lastIndex,
			err

	}

	if negNumSearchSpec.foundLeadingNegNumSign == true {

		foundNegNumSignSymbols = true

		// Nothing to do.Found Leading Neg Num Sign
		// Symbols on a previous cycle.
		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	lenNegNumSignTargetSearchChars := len(targetSearchString.CharsToSearch)

	if lenNegNumSignTargetSearchChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'targetSearchString' is empty and invalid!\n"+
			"'targetSearchString.CharsToSearch' has an array length of zero.\n"+
			"There are no target search characters in which to search\n"+
			"for Negative Number Sign symbols.\n",
			ePrefix.String())

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	if startingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid!\n"+
			"startingSearchIndex has a length less than zero.\n"+
			"startingSearchIndex length = '%v'\n",
			ePrefix.String(),
			startingSearchIndex)

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	if startingSearchIndex >= lenNegNumSignTargetSearchChars {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid!\n"+
			"startingSearchIndex has a length less than zero.\n"+
			"startingSearchIndex length = '%v'\n",
			ePrefix.String(),
			startingSearchIndex)

		return foundNegNumSignSymbols,
			lastIndex,
			err

	}

	lenLeadingNegNumSymbols := len(negNumSearchSpec.leadingNegNumSignSymbols)

	if lenLeadingNegNumSymbols == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NegativeNumberSearchSpec is invalid!\n"+
			"No Leading Negative Number Sign Symbols have been previously\n"+
			"configured. The length of the negNumSearchSpec.leadingNegNumSignSymbols"+
			"array is zero. The array is empty.\n",
			ePrefix.String())

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	leadingNegNumSymbolsLastIdx := lenLeadingNegNumSymbols - 1

	// Set internal processing flag
	// foundFirstNumericDigitInNumStr
	negNumSearchSpec.foundFirstNumericDigitInNumStr =
		foundFirstNumericDigitInNumStr

	leadNegNumSymIdx := 0

	for i := startingSearchIndex; i < lenNegNumSignTargetSearchChars; i++ {

		if negNumSearchSpec.leadingNegNumSignSymbols[leadNegNumSymIdx] !=
			targetSearchString.CharsToSearch[i] {

			// The Leading Negative Number Symbols were
			// NOT found in this search string.
			return foundNegNumSignSymbols,
				lastIndex,
				err

		}

		leadNegNumSymIdx++

		if leadNegNumSymIdx > leadingNegNumSymbolsLastIdx {
			// Found the Neg Num Sign Symbol
			lastIndex = i

			negNumSearchSpec.foundLeadingNegNumSign = true

			negNumSearchSpec.foundLeadingNegNumSignIndex =
				startingSearchIndex

			foundNegNumSignSymbols = true

			return foundNegNumSignSymbols,
				lastIndex,
				err
		}
	}

	return foundNegNumSignSymbols,
		lastIndex,
		err
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

	if negNumSearchSpec.negNumSignPosition == NSignSymPos.Before() {

		if len(negNumSearchSpec.trailingNegNumSignSymbols) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, it contains Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return isValid, err

		}

		if len(negNumSearchSpec.leadingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading Negative Number Sign.\n"+
				"However, no Leading Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return isValid, err
		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.leadingNegNumSignSymbols,
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

		isValid = true

		return isValid, err
	}

	if negNumSearchSpec.negNumSignPosition == NSignSymPos.After() {

		if len(negNumSearchSpec.leadingNegNumSignSymbols) > 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return isValid, err

		}

		if len(negNumSearchSpec.trailingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, no Trailing Negative Number Sign characters are configured.\n",
				ePrefix.String())

			return isValid, err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.trailingNegNumSignSymbols,
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

		isValid = true

		return isValid, err
	}

	if negNumSearchSpec.negNumSignPosition == NSignSymPos.BeforeAndAfter() {

		if len(negNumSearchSpec.leadingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Leading and Trailing Negative Number Sign.\n"+
				"However, it contains NO Leading Negative Number Sign characters.\n",
				ePrefix.String())

			return isValid, err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.leadingNegNumSignSymbols,
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

		if len(negNumSearchSpec.trailingNegNumSignSymbols) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of 'NegativeNumberSearchSpec' is invalid!\n"+
				"It is configured as a Trailing Negative Number Sign.\n"+
				"However, it contains NO Trailing Negative Number Sign characters.\n",
				ePrefix.String())

			return isValid, err

		}

		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			negNumSearchSpec.trailingNegNumSignSymbols,
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

		isValid = true

		return isValid, err

	}

	// NegativeNumberSearchSpec is invalid!
	return isValid, err
}

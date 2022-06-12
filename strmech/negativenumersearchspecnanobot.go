package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type negNumSignSearchNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingNegNumSearchSpec' to input parameter
// 'targetNegNumSearchSpec'. Both instances are of type
// NegativeNumberSearchSpec.
//
// IMPORTANT
// -----------------------------------------------------------------
// Be advised that the data fields in 'targetNegNumSearchSpec' will be
// overwritten.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetNegNumSearchSpec        *NegativeNumberSearchSpec
//     - A pointer to a NegativeNumberSearchSpec instance. All the
//       member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'incomingNegNumSearchSpec'.
//
//       'targetNegNumSearchSpec' is the target of this copy
//       operation.
//
//
//  incomingNegNumSearchSpec      *NegativeNumberSearchSpec
//     - A pointer to another NegativeNumberSearchSpec instance. All
//       the member variable data values from this object will
//       be copied to corresponding member variables in
//       'targetNegNumSearchSpec'.
//
//       'incomingNegNumSearchSpec' is the source for this copy
//       operation.
//
//       If 'incomingNegNumSearchSpec' is determined to be invalid,
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
func (negNumSearchNanobot *negNumSignSearchNanobot) copyIn(
	targetNegNumSearchSpec *NegativeNumberSearchSpec,
	incomingNegNumSearchSpec *NegativeNumberSearchSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetNegNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingNegNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	negNumSearchAtom := negNumSearchSpecAtom{}

	_,
		err2 =
		negNumSearchAtom.
			testValidityOfNegNumSearchSpec(
				incomingNegNumSearchSpec,
				nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'incomingNegNumSearchSpec' failed!\n"+
			"This instance of NegativeNumberSearchSpec is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	// Reset all targetNegNumSearchSpec member
	//  variables to their zero values
	negNumSearchAtom.
		empty(targetNegNumSearchSpec)

	targetNegNumSearchSpec.negNumSignPosition =
		incomingNegNumSearchSpec.negNumSignPosition

	var lenLeadingNegNumSignSymbols,
		lenTrailingNegNumSignSymbols int

	if targetNegNumSearchSpec.negNumSignPosition ==
		NSignSymPos.Before() {

		lenLeadingNegNumSignSymbols =
			len(incomingNegNumSearchSpec.leadingNegNumSignSymbols)

		targetNegNumSearchSpec.leadingNegNumSignSymbols =
			make([]rune, lenLeadingNegNumSignSymbols)

		for i := 0; i < lenLeadingNegNumSignSymbols; i++ {
			targetNegNumSearchSpec.leadingNegNumSignSymbols[i] =
				incomingNegNumSearchSpec.leadingNegNumSignSymbols[i]
		}

		targetNegNumSearchSpec.foundLeadingNegNumSign =
			incomingNegNumSearchSpec.foundLeadingNegNumSign

		targetNegNumSearchSpec.foundLeadingNegNumSignIndex =
			incomingNegNumSearchSpec.foundLeadingNegNumSignIndex

	} else if targetNegNumSearchSpec.negNumSignPosition ==
		NSignSymPos.BeforeAndAfter() {

		lenTrailingNegNumSignSymbols =
			len(incomingNegNumSearchSpec.trailingNegNumSignSymbols)

		targetNegNumSearchSpec.trailingNegNumSignSymbols =
			make([]rune, lenTrailingNegNumSignSymbols)

		for i := 0; i < lenTrailingNegNumSignSymbols; i++ {
			targetNegNumSearchSpec.trailingNegNumSignSymbols[i] =
				incomingNegNumSearchSpec.trailingNegNumSignSymbols[i]
		}

		targetNegNumSearchSpec.foundTrailingNegNumSign =
			incomingNegNumSearchSpec.foundTrailingNegNumSign

		targetNegNumSearchSpec.foundTrailingNegNumSignIndex =
			incomingNegNumSearchSpec.foundTrailingNegNumSignIndex

	} else {
		// Must be targetNegNumSearchSpec.negNumSignPosition ==
		//            NSignSymPos.After()

		// Leading data elements
		lenLeadingNegNumSignSymbols =
			len(incomingNegNumSearchSpec.leadingNegNumSignSymbols)

		targetNegNumSearchSpec.leadingNegNumSignSymbols =
			make([]rune, lenLeadingNegNumSignSymbols)

		for i := 0; i < lenLeadingNegNumSignSymbols; i++ {
			targetNegNumSearchSpec.leadingNegNumSignSymbols[i] =
				incomingNegNumSearchSpec.leadingNegNumSignSymbols[i]
		}

		targetNegNumSearchSpec.foundLeadingNegNumSign =
			incomingNegNumSearchSpec.foundLeadingNegNumSign

		targetNegNumSearchSpec.foundLeadingNegNumSignIndex =
			incomingNegNumSearchSpec.foundLeadingNegNumSignIndex

		// Trailing Data Elements
		lenTrailingNegNumSignSymbols =
			len(incomingNegNumSearchSpec.trailingNegNumSignSymbols)

		targetNegNumSearchSpec.trailingNegNumSignSymbols =
			make([]rune, lenTrailingNegNumSignSymbols)

		for i := 0; i < lenTrailingNegNumSignSymbols; i++ {
			targetNegNumSearchSpec.trailingNegNumSignSymbols[i] =
				incomingNegNumSearchSpec.trailingNegNumSignSymbols[i]
		}

		targetNegNumSearchSpec.foundTrailingNegNumSign =
			incomingNegNumSearchSpec.foundTrailingNegNumSign

		targetNegNumSearchSpec.foundTrailingNegNumSignIndex =
			incomingNegNumSearchSpec.foundTrailingNegNumSignIndex

	}

	targetNegNumSearchSpec.foundFirstNumericDigitInNumStr =
		incomingNegNumSearchSpec.foundFirstNumericDigitInNumStr

	targetNegNumSearchSpec.foundNegNumSignSymbols =
		incomingNegNumSearchSpec.foundNegNumSignSymbols

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'negNumSearchSpec'. a pointer to an instance of
// NegativeNumberSearchSpec.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The input parameter 'negNumSearchSpec' is determined to be
// invalid, this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchSpec           *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of NegativeNumberSearchSpec.
//
//       If the member variable data values encapsulated by
//       'negNumSearchSpec' are found to be invalid, this method will
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
//  copyOfNegNumSearchSpec     NegativeNumberSearchSpec
//     - If this method completes successfully, a deep copy of
//       input parameter 'negNumSearchSpec' will be created and returned
//       in a new instance of NegativeNumberSearchSpec.
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
func (negNumSearchNanobot *negNumSignSearchNanobot) copyOut(
	negNumSearchSpec *NegativeNumberSearchSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	copyOfNegNumSearchSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return copyOfNegNumSearchSpec, err

	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return copyOfNegNumSearchSpec, err
	}

	var err2 error

	_,
		err2 =
		negNumSearchSpecAtom{}.ptr().
			testValidityOfNegNumSearchSpec(
				negNumSearchSpec,
				nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'negNumSearchSpec' failed!\n"+
			"This instance of NegativeNumberSearchSpec is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return copyOfNegNumSearchSpec, err
	}

	copyOfNegNumSearchSpec.negNumSignPosition =
		negNumSearchSpec.negNumSignPosition

	var lenLeadingNegNumSignSymbols,
		lenTrailingNegNumSignSymbols int

	if copyOfNegNumSearchSpec.negNumSignPosition ==
		NSignSymPos.Before() {

		lenLeadingNegNumSignSymbols =
			len(negNumSearchSpec.leadingNegNumSignSymbols)

		copyOfNegNumSearchSpec.leadingNegNumSignSymbols =
			make([]rune, lenLeadingNegNumSignSymbols)

		for i := 0; i < lenLeadingNegNumSignSymbols; i++ {
			copyOfNegNumSearchSpec.leadingNegNumSignSymbols[i] =
				negNumSearchSpec.leadingNegNumSignSymbols[i]
		}

		copyOfNegNumSearchSpec.foundLeadingNegNumSign =
			negNumSearchSpec.foundLeadingNegNumSign

		copyOfNegNumSearchSpec.foundLeadingNegNumSignIndex =
			negNumSearchSpec.foundLeadingNegNumSignIndex

	} else if copyOfNegNumSearchSpec.negNumSignPosition ==
		NSignSymPos.BeforeAndAfter() {

		lenTrailingNegNumSignSymbols =
			len(negNumSearchSpec.trailingNegNumSignSymbols)

		copyOfNegNumSearchSpec.trailingNegNumSignSymbols =
			make([]rune, lenTrailingNegNumSignSymbols)

		for i := 0; i < lenTrailingNegNumSignSymbols; i++ {
			copyOfNegNumSearchSpec.trailingNegNumSignSymbols[i] =
				negNumSearchSpec.trailingNegNumSignSymbols[i]
		}

		copyOfNegNumSearchSpec.foundTrailingNegNumSign =
			negNumSearchSpec.foundTrailingNegNumSign

		copyOfNegNumSearchSpec.foundTrailingNegNumSignIndex =
			negNumSearchSpec.foundTrailingNegNumSignIndex

	} else {
		// Must be copyOfNegNumSearchSpec.negNumSignPosition ==
		//            NSignSymPos.After()

		// Leading data elements
		lenLeadingNegNumSignSymbols =
			len(negNumSearchSpec.leadingNegNumSignSymbols)

		copyOfNegNumSearchSpec.leadingNegNumSignSymbols =
			make([]rune, lenLeadingNegNumSignSymbols)

		for i := 0; i < lenLeadingNegNumSignSymbols; i++ {
			copyOfNegNumSearchSpec.leadingNegNumSignSymbols[i] =
				negNumSearchSpec.leadingNegNumSignSymbols[i]
		}

		copyOfNegNumSearchSpec.foundLeadingNegNumSign =
			negNumSearchSpec.foundLeadingNegNumSign

		copyOfNegNumSearchSpec.foundLeadingNegNumSignIndex =
			negNumSearchSpec.foundLeadingNegNumSignIndex

		// Trailing Data Elements
		lenTrailingNegNumSignSymbols =
			len(negNumSearchSpec.trailingNegNumSignSymbols)

		copyOfNegNumSearchSpec.trailingNegNumSignSymbols =
			make([]rune, lenTrailingNegNumSignSymbols)

		for i := 0; i < lenTrailingNegNumSignSymbols; i++ {
			copyOfNegNumSearchSpec.trailingNegNumSignSymbols[i] =
				negNumSearchSpec.trailingNegNumSignSymbols[i]
		}

		copyOfNegNumSearchSpec.foundTrailingNegNumSign =
			negNumSearchSpec.foundTrailingNegNumSign

		copyOfNegNumSearchSpec.foundTrailingNegNumSignIndex =
			negNumSearchSpec.foundTrailingNegNumSignIndex

	}

	copyOfNegNumSearchSpec.foundFirstNumericDigitInNumStr =
		negNumSearchSpec.foundFirstNumericDigitInNumStr

	copyOfNegNumSearchSpec.foundNegNumSignSymbols =
		negNumSearchSpec.foundNegNumSignSymbols

	return copyOfNegNumSearchSpec, err
}

// ptr - Returns a pointer to a new instance of
// negNumSignSearchNanobot.
//
func (negNumSearchNanobot negNumSignSearchNanobot) ptr() *negNumSignSearchNanobot {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	return &negNumSignSearchNanobot{
		lock: new(sync.Mutex),
	}
}

// leadingNegSignSymSearch - Performs a search for Leading Negative
// Number Sign Symbols in a number string. The text characters to
// be searched are passed via input parameter 'targetSearchString'.
//
// Leading Negative Number Symbols are positioned to the left of
// numeric digits within a Number String.
//
// Leading Negative Number Symbols are used by various countries
// including the United States and Canada.
//     Examples:  -127.45   -654
//
// This method is almost exclusively used by Number String parsing
// functions.
//
// Number String parsing functions will attempt to identify a
// Negative Number Sign Symbol or Symbols in strings of numeric
// digits called 'Number Strings'. Number String parsing functions
// review strings of text characters containing numeric digits and
// convert those numeric digits to numeric values. The presence or
// absence of Negative Number Sign Symbols determines whether a
// numeric value is positive or negative.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSignSpec                  *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. The
//       Target Search String ('targetSearchString') will examined
//       in an attempt to locate Leading Negative Number Sign
//       Symbols specified by 'negNumSignSpec'.
//
//
//  targetSearchString              *RuneArrayDto
//     - A pointer to a RuneArrayDto. Type
//       RuneArrayDto contains the string of text
//       characters which will be searched for the presence of a
//       Leading Negative Number Sign Symbols specified by
//       'negNumSignSpec', an instance of NegativeNumberSearchSpec.
//
//			  type RuneArrayDto struct {
//                 CharsArray []rune
//			  }
//
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
//     - The 'startingSearchIndex' parameter specifies the zero
//       based index in the Target Search Characters String
//       ('targetSearchString') from which the search for Leading
//       Negative Number Symbols will commence.
//
//       If this value is less than zero or greater than the
//       length of 'targetSearchString' minus one, an error will be
//       returned.
//
//
//  errPrefDto                      *ePref.ErrPrefixDto
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
//       However, if the Leading Negative Number Symbols are
//       located in the Target Search Characters, the value of
//       'lastIndex' will be set to the index in the Target Search
//       String ('targetSearchString') occupied by the last text
//       character in the specified Leading Negative Number Sign
//       Symbols.
//
//         Example:
//
//           Target Search String: "xx(-)567890"
//
//           Leading Negative Number Sign Symbols (3-characters):
//                   "(-)"
//
//           Note: "(-)" is a negative number sign used in the UK.
//
//           lastIndex = 4  The ")" in Target Search String.
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
func (negNumSearchNanobot *negNumSignSearchNanobot) leadingNegSignSymSearch(
	negNumSearchSpec *NegativeNumberSearchSpec,
	targetSearchString *RuneArrayDto,
	foundFirstNumericDigitInNumStr bool,
	startingSearchIndex int,
	errPrefDto *ePref.ErrPrefixDto) (
	foundNegNumSignSymbols bool,
	lastIndex int,
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundNegNumSignSymbols = false

	// This assumes startingSearchIndex has
	// already been validated
	lastIndex = startingSearchIndex

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
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

	var err2 error
	negNumSignAtom := negNumSearchSpecAtom{}

	_,
		err2 = negNumSignAtom.testValidityOfNegNumSearchSpec(
		negNumSearchSpec,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: The current instance of NegativeNumberSearchSpec (negNumSearchSpec)\n"+
			"is invalid. The Number String parsing operation has been aborted.\n"+
			"Validation checks returned the following error for this intance of\n"+
			"NegativeNumberSearchSpec:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

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

	lenNegNumSignTargetSearchChars := len(targetSearchString.CharsArray)

	if lenNegNumSignTargetSearchChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'targetSearchString' is empty and invalid!\n"+
			"'targetSearchString.CharsArray' has an array length of zero.\n"+
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

	// Set internal processing flag
	// foundFirstNumericDigitInNumStr
	negNumSearchSpec.foundFirstNumericDigitInNumStr =
		foundFirstNumericDigitInNumStr

	j := 0

	for i := startingSearchIndex; i < lenNegNumSignTargetSearchChars; i++ {

		if negNumSearchSpec.leadingNegNumSignSymbols[j] !=
			targetSearchString.CharsArray[i] {

			// The Leading Negative Number Symbols were
			// NOT found in this search string.
			return foundNegNumSignSymbols,
				lastIndex,
				err

		}

		j++

		if j >= lenLeadingNegNumSymbols {
			// Search Was SUCCESSFUL!
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

// setLeadingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Leading Negative Number Sign Specification. All internal
// member variables are then configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted before applying the
// new configuration specifications.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchSpec           *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
//       instance will be configured as a Leading Negative Number
//       Sign Specification. All previous configuration data will be
//       deleted and replaced with a new Leading Negative Number
//       Sign configuration.
//
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance,
//       'negNumSearchSpec'.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
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
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchNanobot *negNumSignSearchNanobot) setLeadingNegNumSearchSpec(
	negNumSearchSpec *NegativeNumberSearchSpec,
	leadingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setLeadingNegNumSearchSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(leadingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	negNumSignAtom := negNumSearchSpecAtom{}

	negNumSignAtom.empty(
		negNumSearchSpec)

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		leadingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSearchSpec.leadingNegNumSignSymbols,
		&leadingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSearchSpec<-leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSearchSpec.negNumSignPosition = NSignSymPos.Before()

	return err
}

// setLeadingAndTrailingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Leading and Trailing Negative Number Sign Specification.
// All internal member variables are then configured using the
// input parameter 'leadingNegNumSignSymbols' and
// 'trailingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted and replaced with the new
// configuration specifications.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, in the US and Canada
// parentheses "()" are used to indicate negative numeric
// values. Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchSpec           *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
//       instance will be configured as a Leading and Trailing
//       Negative Number Sign Specification. All previous
//       configuration data will be deleted and replaced with a new
//       Leading and Trailing Negative Number Sign configuration.
//
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the character or characters
//       which comprise the Leading Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance,
//       'negNumSearchSpec'.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance,
//       'negNumSearchSpec'.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
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
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchNanobot *negNumSignSearchNanobot) setLeadingAndTrailingNegNumSearchSpec(
	negNumSearchSpec *NegativeNumberSearchSpec,
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setLeadingAndTrailingNegNumSearchSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(leadingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	if len(trailingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		leadingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNegNumSignSymbols' is invalid!\n"+
			"'leadingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		trailingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSearchSpec.leadingNegNumSignSymbols,
		&leadingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSearchSpec<-leadingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSearchSpec.trailingNegNumSignSymbols,
		&trailingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSearchSpec<-trailingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSearchSpec.negNumSignPosition = NSignSymPos.BeforeAndAfter()

	return err
}

// setTrailingNegNumSearchSpec - Receives an instance of
// NegativeNumberSearchSpec and proceeds to configure that instance
// as a Trailing Negative Number Sign Specification. All internal
// member variables are then configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Any previous configuration data associated with this instance of
// NegativeNumberSearchSpec will be deleted before applying the
// new configuration specifications.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchSpec           *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
//       instance will be configured as a Trailing Negative Number
//       Sign Specification. All previous configuration data will
//       be deleted and replaced with a new Trailing Negative Number
//       Sign configuration.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the character or characters
//       which comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance,
//       'negNumSearchSpec'.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
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
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchNanobot *negNumSignSearchNanobot) setTrailingNegNumSearchSpec(
	negNumSearchSpec *NegativeNumberSearchSpec,
	trailingNegNumSignSymbols []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"setTrailingNegNumSearchSpec()",
		"")

	if err != nil {
		return err
	}

	if negNumSearchSpec == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'negNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(trailingNegNumSignSymbols) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' is an empty array. The array length\n"+
			"is zero (0)!\n",
			ePrefix.String())

		return err
	}

	negNumSearchSpecAtom{}.ptr().empty(
		negNumSearchSpec)

	sMechPreon := strMechPreon{}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		trailingNegNumSignSymbols,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNegNumSignSymbols' is invalid!\n"+
			"'trailingNegNumSignSymbols' returned the following validation error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	err = sMechPreon.copyRuneArrays(
		&negNumSearchSpec.trailingNegNumSignSymbols,
		&trailingNegNumSignSymbols,
		true,
		ePrefix.XCpy(
			"negNumSearchSpec<-trailingNegNumSignSymbols"))

	if err != nil {
		return err
	}

	negNumSearchSpec.negNumSignPosition = NSignSymPos.After()

	return err
}

// trailingNegSignSymSearch - Performs a search for Trailing Negative
// Number Sign Symbols in a number string. The text characters to
// be searched are passed via input parameter 'targetSearchString'.
//
// Trailing Negative Number Symbols are positioned to the right of
// numeric digits within a Number String.
//
// Trailing Negative Number Symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is almost exclusively used by Number String parsing
// functions.
//
// Number String parsing functions will attempt to identify a
// Negative Number Sign Symbol or Symbols in strings of numeric
// digits called 'Number Strings'. Number String parsing functions
// review strings of text characters containing numeric digits and
// convert those numeric digits to numeric values. The presence or
// absence of Negative Number Sign Symbols determines whether a
// numeric value is positive or negative.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  negNumSignSpec                  *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. The
//       Target Search String ('targetSearchString') will examined
//       in an attempt to locate Trailing Negative Number Sign
//       Symbols specified by 'negNumSignSpec'.
//
//
//  targetSearchString              *RuneArrayDto
//     - A pointer to a RuneArrayDto. Type
//       RuneArrayDto contains the string of text
//       characters which will be searched for the presence of a
//       Trailing Negative Number Sign Symbols specified by
//       'negNumSignSpec', an instance of NegativeNumberSearchSpec.
//
//			  type RuneArrayDto struct {
//                 CharsArray []rune
//			  }
//
//
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
//          impossible to find valid Trailing Negative
//          Number Symbols. By definition, Trailing Negative
//          Number Symbols always occur before the first
//          numeric digit in a number string.
//
//
//  startingSearchIndex             int
//     - The 'startingSearchIndex' parameter specifies the zero
//       based index in the Target Search Characters String
//       ('targetSearchString') from which the search for Trailing
//       Negative Number Symbols will commence.
//
//       If this value is less than zero or greater than the
//       length of 'targetSearchString' minus one, an error will be
//       returned.
//
//
//  errPrefDto                      *ePref.ErrPrefixDto
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
//       signal whether the search for Trailing Negative Number
//       Symbols was successful.
//
//       A return value of 'false' signals that the search for
//       Trailing Negative Number Symbols was unsuccessful and
//       the Symbols were NOT located in the Target Search
//       Characters.
//
//       A return value of 'true' signals that the Trailing
//       Negative Number Symbols were located in the Target Search
//       Characters and the search was therefore successful.
//
//
//  lastIndex                       int
//       If the search for Trailing Negative Number Symbols was
//       unsuccessful, the value of 'lastIndex' will be set to
//       'startingSearchIndex'.
//
//       However, if the Trailing Negative Number Symbols are
//       located in the Target Search Characters, the value of
//       'lastIndex' will be set to the index in the Target Search
//       String ('targetSearchString') occupied by the last text
//       character in the specified Trailing Negative Number Sign
//       Symbols.
//
//         Example:
//
//           Target Search String: "xx(-)567890"
//
//           Trailing Negative Number Sign Symbols (3-characters):
//                   "(-)"
//
//           Note: "(-)" is a negative number sign used in the UK.
//
//           lastIndex = 4  The ")" in Target Search String.
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
func (negNumSearchNanobot *negNumSignSearchNanobot) trailingNegSignSymSearch(
	negNumSearchSpec *NegativeNumberSearchSpec,
	targetSearchString *RuneArrayDto,
	foundFirstNumericDigitInNumStr bool,
	startingSearchIndex int,
	errPrefDto *ePref.ErrPrefixDto) (
	foundNegNumSignSymbols bool,
	lastIndex int,
	err error) {

	if negNumSearchNanobot.lock == nil {
		negNumSearchNanobot.lock = new(sync.Mutex)
	}

	negNumSearchNanobot.lock.Lock()

	defer negNumSearchNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundNegNumSignSymbols = false

	// This assumes startingSearchIndex has
	// already been validated
	lastIndex = startingSearchIndex

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSearchNanobot."+
			"trailingNegSignSymSearch()",
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

	var err2 error
	negNumSignAtom := negNumSearchSpecAtom{}

	_,
		err2 = negNumSignAtom.testValidityOfNegNumSearchSpec(
		negNumSearchSpec,
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: The current instance of NegativeNumberSearchSpec (negNumSearchSpec)\n"+
			"is invalid. The Number String parsing operation has been aborted.\n"+
			"Validation checks returned the following error for this intance of\n"+
			"NegativeNumberSearchSpec:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

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
		// numeric digit. Further, Trailing Neg Num Sign
		// search is pointless.
		//
		// Once the first Numeric Digit is discovered in
		// a Number String parsing operation, it is
		// impossible to find valid Trailing Negative
		// Number Symbols.
		return foundNegNumSignSymbols,
			lastIndex,
			err

	}

	if negNumSearchSpec.foundTrailingNegNumSign == true {

		foundNegNumSignSymbols = true

		// Nothing to do.Found Trailing Neg Num Sign
		// Symbols on a previous cycle.
		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	lenNegNumSignTargetSearchChars := len(targetSearchString.CharsArray)

	if lenNegNumSignTargetSearchChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'targetSearchString' is empty and invalid!\n"+
			"'targetSearchString.CharsArray' has an array length of zero.\n"+
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

	lenTrailingNegNumSymbols := len(negNumSearchSpec.trailingNegNumSignSymbols)

	if lenTrailingNegNumSymbols == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NegativeNumberSearchSpec is invalid!\n"+
			"No Trailing Negative Number Sign Symbols have been previously\n"+
			"configured. The length of the negNumSearchSpec.trailingNegNumSignSymbols"+
			"array is zero. The array is empty.\n",
			ePrefix.String())

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	// Set internal processing flag
	// foundFirstNumericDigitInNumStr
	negNumSearchSpec.foundFirstNumericDigitInNumStr =
		foundFirstNumericDigitInNumStr

	j := 0

	for i := startingSearchIndex; i < lenNegNumSignTargetSearchChars; i++ {

		if negNumSearchSpec.trailingNegNumSignSymbols[j] !=
			targetSearchString.CharsArray[i] {

			// The Trailing Negative Number Symbols were
			// NOT found in this search string.
			return foundNegNumSignSymbols,
				lastIndex,
				err

		}

		j++

		if j >= lenTrailingNegNumSymbols {
			// Search Was SUCCESSFUL!
			// Found the Neg Num Sign Symbol
			lastIndex = i

			negNumSearchSpec.foundTrailingNegNumSign = true

			negNumSearchSpec.foundTrailingNegNumSignIndex =
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

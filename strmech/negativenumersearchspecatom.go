package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type negNumSearchSpecAtom struct {
	lock *sync.Mutex
}

func (negNumSearchAtom *negNumSearchSpecAtom) empty(
	negNumSignSpec *NegativeNumberSearchSpec) {

	if negNumSearchAtom.lock == nil {
		negNumSearchAtom.lock = new(sync.Mutex)
	}

	negNumSearchAtom.lock.Lock()

	defer negNumSearchAtom.lock.Unlock()

	if negNumSignSpec == nil {
		return
	}

	negNumSignSpec.negNumSignPosition = NSignSymPos.None()
	negNumSignSpec.leadingNegNumSignSymbols = nil
	negNumSignSpec.trailingNegNumSignSymbols = nil

	negNumSearchSpecElectron{}.ptr().
		emptyProcessingFlags(negNumSignSpec)

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
	negNumSignSpec *NegativeNumberSearchSpec,
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

	if negNumSignSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negNumSignSpec' is a nil pointer!\n",
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

	if negNumSignSpec.foundFirstNumericDigitInNumStr {

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

	if negNumSignSpec.foundLeadingNegNumSign == true {

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

	lenLeadingNegNumSymbols := len(negNumSignSpec.leadingNegNumSignSymbols)

	if lenLeadingNegNumSymbols == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NegativeNumberSearchSpec is invalid!\n"+
			"No Leading Negative Number Sign Symbols have been previously\n"+
			"configured. The length of the negNumSignSpec.leadingNegNumSignSymbols"+
			"array is zero. The array is empty.\n",
			ePrefix.String())

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	leadingNegNumSymbolsLastIdx := lenLeadingNegNumSymbols - 1

	// Set internal processing flag
	// foundFirstNumericDigitInNumStr
	negNumSignSpec.foundFirstNumericDigitInNumStr =
		foundFirstNumericDigitInNumStr

	leadNegNumSymIdx := 0

	for i := startingSearchIndex; i < lenNegNumSignTargetSearchChars; i++ {

		if negNumSignSpec.leadingNegNumSignSymbols[leadNegNumSymIdx] !=
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

			negNumSignSpec.foundLeadingNegNumSign = true

			negNumSignSpec.foundLeadingNegNumSignIndex =
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

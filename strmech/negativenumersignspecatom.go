package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type negNumSignSpecAtom struct {
	lock *sync.Mutex
}

func (negNumSignAtom *negNumSignSpecAtom) empty(
	negNumSignSpec *NegativeNumberSignSpec) {

	if negNumSignAtom.lock == nil {
		negNumSignAtom.lock = new(sync.Mutex)
	}

	negNumSignAtom.lock.Lock()

	defer negNumSignAtom.lock.Unlock()

	if negNumSignSpec == nil {
		return
	}

	negNumSignSpec.negNumSignPosition = NSignSymPos.None()
	negNumSignSpec.leadingNegNumSignSymbols = nil
	negNumSignSpec.trailingNegNumSignSymbols = nil

	negNumSignSpecElectron{}.ptr().
		emptyProcessingFlags(negNumSignSpec)

	return
}

// ptr - Returns a pointer to a new instance of
// negNumSignSpecAtom.
//
func (negNumSignAtom negNumSignSpecAtom) ptr() *negNumSignSpecAtom {

	if negNumSignAtom.lock == nil {
		negNumSignAtom.lock = new(sync.Mutex)
	}

	negNumSignAtom.lock.Lock()

	defer negNumSignAtom.lock.Unlock()

	return &negNumSignSpecAtom{
		lock: new(sync.Mutex),
	}
}

func (negNumSignAtom *negNumSignSpecAtom) beforeNegSignSymSearch(
	negNumSignSpec *NegativeNumberSignSpec,
	foundFirstNumericDigitInNumStr bool,
	startingSearchIndex int,
	errPrefDto *ePref.ErrPrefixDto) (
	foundNegNumSignSymbols bool,
	lastIndex int,
	err error) {

	if negNumSignAtom.lock == nil {
		negNumSignAtom.lock = new(sync.Mutex)
	}

	negNumSignAtom.lock.Lock()

	defer negNumSignAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundNegNumSignSymbols = false

	// This assumes startingSearchIndex has
	// already been validated
	lastIndex = startingSearchIndex

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSignSpecElectron."+
			"beforeNegSignSymSearch()",
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

	lenNegNumSignTargetSearchChars := len(negNumSignSpec.negNumSignTargetSearchChars)

	if lenNegNumSignTargetSearchChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NegativeNumberSignSpec has NOT"+
			"been properly configured for a number string parsing operation."+
			"negNumSignSpec.negNumSignTargetSearchChars has a length of zero!"+
			"There are no target search characters in which to search for "+
			"Negative Number Sign symbols.\n",
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

	if negNumSignSpec.foundFirstNumericDigitInNumStr {

		// Nothing to do. Already found the first
		// numeric digit. Further, Leading Neg Num Sign
		// search is pointless.
		return foundNegNumSignSymbols,
			lastIndex,
			err

	}

	if negNumSignSpec.foundLeadingNegNumSign == true {

		foundNegNumSignSymbols = true

		// Nothing to do. Already found Leading Neg Num Sign
		// Symbols
		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	lenLeadingNegNumSymbols := len(negNumSignSpec.leadingNegNumSignSymbols)

	if lenLeadingNegNumSymbols == 0 {

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
			negNumSignSpec.negNumSignTargetSearchChars[i] {

			return foundNegNumSignSymbols,
				lastIndex,
				err

		}

		leadNegNumSymIdx++

		if leadNegNumSymIdx == leadingNegNumSymbolsLastIdx {
			// Found the Neg Num Sign Symbol
			lastIndex = i

			return foundNegNumSignSymbols,
				lastIndex,
				err
		}
	}

	return foundNegNumSignSymbols,
		lastIndex,
		err
}

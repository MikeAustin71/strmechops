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
	searchTargetChars *[]rune,
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

	if searchTargetChars == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'searchTargetChars' is a nil pointer!\n",
			ePrefix.String())

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	if *searchTargetChars == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'searchTargetChars' is empty and\n"+
			"has a length of zero!\n",
			ePrefix.String())

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

	lenSrcRuneAry := len(*searchTargetChars)

	if startingSearchIndex >= lenSrcRuneAry {

		return foundNegNumSignSymbols,
			lastIndex,
			err

	}

	if startingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid!\n"+
			"'startingSearchIndex' has a value less than zero!\n"+
			"startingSearchIndex = '%v'\n",
			ePrefix.String(),
			startingSearchIndex)

		return foundNegNumSignSymbols,
			lastIndex,
			err

	}

	leadNegNumSymIdx := 0

	concreteSearchTarget := make([]rune, lenSrcRuneAry)

	itemsCopied := copy(concreteSearchTarget, *searchTargetChars)

	if itemsCopied <= 0 {

		err = fmt.Errorf("%v\n"+
			"Error: 'searchTargetChars' copy operation failed!\n"+
			"copy(concreteSearchTarget, *searchTargetChars)\n"+
			"itemsCopied = '%v'\n",
			ePrefix.String(),
			itemsCopied)

		return foundNegNumSignSymbols,
			lastIndex,
			err

	}

	for i := startingSearchIndex; i < lenSrcRuneAry; i++ {

		if negNumSignSpec.leadingNegNumSignSymbols[leadNegNumSymIdx] !=
			concreteSearchTarget[i] {

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

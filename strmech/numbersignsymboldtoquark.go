package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numberSignSymbolDtoQuark struct {
	lock *sync.Mutex
}

// isNumberSignSymbolFoundInNumber - Receives a pointer to an
// instance of NumberSignSymbolDto and performs an analysis of
// internal data elements to determine if the entire number sign
// symbol has been located in a target number or number string.
//
// A return value of 'true' signals that the entire number sign
// symbol has been located in the target number or number string.
//
func (nSignSymQuark *numberSignSymbolDtoQuark) isNumberSignSymbolFoundInNumber(
	numSignSymbol *NumberSignSymbolDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isNumSignFoundInNumber bool,
	err error) {

	if nSignSymQuark.lock == nil {
		nSignSymQuark.lock = new(sync.Mutex)
	}

	nSignSymQuark.lock.Lock()

	defer nSignSymQuark.lock.Unlock()

	isNumSignFoundInNumber = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberSignSymbolDtoElectron.isNumberSignSymbolFoundInNumber()",
		"")

	if err != nil {
		return isNumSignFoundInNumber, err
	}

	if numSignSymbol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSignSymbol' is a 'nil' pointer!\n",
			ePrefix.String())

		return isNumSignFoundInNumber, err
	}

	// isNumSignFoundInNumber == false
	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).Before() {

		isNumSignFoundInNumber =
			numSignSymbol.leadingNumSignFoundInNumber

	} else if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).After() {

		isNumSignFoundInNumber =
			numSignSymbol.trailingNumSignFoundInNumber
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).BeforeAndAfter() {

		if numSignSymbol.leadingNumSignFoundInNumber &&
			numSignSymbol.trailingNumSignFoundInNumber {

			isNumSignFoundInNumber = true
		}
	}

	return isNumSignFoundInNumber, err
}

// ptr - Returns a pointer to a new instance of
// numberSignSymbolDtoQuark.
//
func (nSignSymQuark numberSignSymbolDtoQuark) ptr() *numberSignSymbolDtoQuark {

	if nSignSymQuark.lock == nil {
		nSignSymQuark.lock = new(sync.Mutex)
	}

	nSignSymQuark.lock.Lock()

	defer nSignSymQuark.lock.Unlock()

	return &numberSignSymbolDtoQuark{
		lock: new(sync.Mutex),
	}
}

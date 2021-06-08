package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numberSignSymbolMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// numberSignSymbolMechanics.
//
func (nSignSymMechanics numberSignSymbolMechanics) ptr() *numberSignSymbolMechanics {

	if nSignSymMechanics.lock == nil {
		nSignSymMechanics.lock = new(sync.Mutex)
	}

	nSignSymMechanics.lock.Lock()

	defer nSignSymMechanics.lock.Unlock()

	newNumSignSymbolMech := numberSignSymbolMechanics{}
	newNumSignSymbolMech.lock = new(sync.Mutex)

	return &newNumSignSymbolMech
}

// setNumberSignSymbol - Receives a pointer to an instance of
// NumberSignSymbol and proceeds to populate the internal data
// elements based on the input parameter values.
//
//
func (nSignSymMechanics *numberSignSymbolMechanics) setNumberSignSymbol(
	numSignSymbol *NumberSignSymbol,
	leadingNumberSign string,
	trailingNumberSign string,
	isNegativeValue bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nSignSymMechanics.lock == nil {
		nSignSymMechanics.lock = new(sync.Mutex)
	}

	nSignSymMechanics.lock.Lock()

	defer nSignSymMechanics.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberSignSymbolMechanics.setNumberSignSymbol()",
		"")

	if err != nil {
		return err
	}

	if numSignSymbol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSignSymbol' is a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	leadingNumSignRunes := []rune(leadingNumberSign)

	trailingNumSignRunes := []rune(trailingNumberSign)

	lenLeadingNumSignRunes := len(leadingNumSignRunes)
	lenTrailingNumSignRunes := len(trailingNumSignRunes)

	if lenLeadingNumSignRunes == 0 &&
		lenTrailingNumSignRunes == 0 {

		err = fmt.Errorf("%v\n" +
			"Error: Input parameters 'leadingNumberSign' and " +
			"'trailingNumberSign' are zero length strings!\n" +
			ePrefix.String())
		return err

	}

	err = numberSignSymbolElectron{}.ptr().emptyNumSignSymbol(
		numSignSymbol,
		ePrefix)

	if err != nil {
		return err
	}

	if lenLeadingNumSignRunes > 0 &&
		lenTrailingNumSignRunes < 0 {

		numSignSymbol.numSignPosition =
			NumSymPos.Before()

		numSignSymbol.leadingNumSignChars =
			make([]rune, lenLeadingNumSignRunes)

		copy(numSignSymbol.leadingNumSignChars,
			leadingNumSignRunes)

	} else if lenLeadingNumSignRunes == 0 &&
		lenTrailingNumSignRunes > 0 {

		numSignSymbol.numSignPosition =
			NumSymPos.After()

		numSignSymbol.trailingNumSignChars =
			make([]rune, lenTrailingNumSignRunes)

		copy(numSignSymbol.trailingNumSignChars,
			trailingNumSignRunes)

	} else {
		// Must be lenLeadingNumSignRunes > 0 &&
		// lenTrailingNumSignRunes > 0

		numSignSymbol.numSignPosition =
			NumSymPos.BeforeAndAfter()

		numSignSymbol.leadingNumSignChars =
			make([]rune, lenLeadingNumSignRunes)

		copy(numSignSymbol.leadingNumSignChars,
			leadingNumSignRunes)

		numSignSymbol.trailingNumSignChars =
			make([]rune, lenTrailingNumSignRunes)

		copy(numSignSymbol.trailingNumSignChars,
			trailingNumSignRunes)

	}

	if isNegativeValue {
		numSignSymbol.numSignType = NumSignVal.Negative()
	} else {
		numSignSymbol.numSignType = NumSignVal.Positive()
	}

	return err
}

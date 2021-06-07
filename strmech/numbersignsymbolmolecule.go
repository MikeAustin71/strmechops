package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numberSignSymbolMolecule struct {
	lock *sync.Mutex
}

// CopyIn - Copies string information from input parameter
// 'incomingNumSignSymbol' to input parameter
// 'targetNumSignSymbol'.
//
// Be advised that the data fields in 'targetNumSignSymbol' will be
// overwritten.
//
func (nSignSymMolecule *numberSignSymbolMolecule) copyIn(
	targetNumSignSymbol *NumberSignSymbol,
	incomingNumSignSymbol *NumberSignSymbol,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nSignSymMolecule.lock == nil {
		nSignSymMolecule.lock = new(sync.Mutex)
	}

	nSignSymMolecule.lock.Lock()

	defer nSignSymMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberSignSymbolMolecule.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetNumSignSymbol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNumSignSymbol' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if incomingNumSignSymbol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNumSignSymbol' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	if targetNumSignSymbol.lock == nil {
		targetNumSignSymbol.lock = new(sync.Mutex)
	}

	targetNumSignSymbol.symFoundInNumber =
		incomingNumSignSymbol.symFoundInNumber

	targetNumSignSymbol.numSignPosition =
		incomingNumSignSymbol.numSignPosition

	targetNumSignSymbol.numSignType =
		incomingNumSignSymbol.numSignType

	lenLeadingSignChars := len(incomingNumSignSymbol.leadingNumSignChars)

	lenTrailingSignChars := len(incomingNumSignSymbol.trailingNumSignChars)

	if lenLeadingSignChars == 0 &&
		lenTrailingSignChars == 0 {
		targetNumSignSymbol.leadingNumSignChars = nil
		targetNumSignSymbol.trailingNumSignChars = nil
		return err
	}

	if lenLeadingSignChars > 0 {
		targetNumSignSymbol.leadingNumSignChars =
			make([]rune, lenLeadingSignChars)

		copy(targetNumSignSymbol.leadingNumSignChars,
			incomingNumSignSymbol.leadingNumSignChars)

	}

	if lenTrailingSignChars > 0 {
		targetNumSignSymbol.trailingNumSignChars =
			make([]rune, lenTrailingSignChars)

		copy(targetNumSignSymbol.trailingNumSignChars,
			incomingNumSignSymbol.trailingNumSignChars)
	}

	return nil
}

// ptr - Returns a pointer to a new instance of
// numberSignSymbolMolecule.
//
func (nSignSymMolecule numberSignSymbolMolecule) ptr() *numberSignSymbolMolecule {

	if nSignSymMolecule.lock == nil {
		nSignSymMolecule.lock = new(sync.Mutex)
	}

	nSignSymMolecule.lock.Lock()

	defer nSignSymMolecule.lock.Unlock()

	return &numberSignSymbolMolecule{
		lock: new(sync.Mutex),
	}
}

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
// If the input parameter 'incomingNumSignSymbol' contains invalid
// data elements, this method will return an error and leave
// 'targetNumSignSymbol' unaltered in its original state.
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
			"Error: Input parameter 'incomingNumSignSymbol' is a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err = numberSignSymbolElectron{}.ptr().
		testValidityOfNumSignSymbol(
			incomingNumSignSymbol,
			ePrefix.XCtx("incomingNumSignSymbol"))

	if err != nil {
		return err
	}

	if targetNumSignSymbol.lock == nil {
		targetNumSignSymbol.lock = new(sync.Mutex)
	}

	targetNumSignSymbol.leadingNumSignFoundInNumber =
		incomingNumSignSymbol.leadingNumSignFoundInNumber

	targetNumSignSymbol.trailingNumSignFoundInNumber =
		incomingNumSignSymbol.trailingNumSignFoundInNumber

	targetNumSignSymbol.leadingNumSignFoundIndex =
		incomingNumSignSymbol.leadingNumSignFoundIndex

	targetNumSignSymbol.trailingNumSignFoundIndex =
		incomingNumSignSymbol.trailingNumSignFoundIndex

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

// copyOut - Returns a deep copy of input parameter
// 'incomingNumSignSymbol' styled as a new instance
// of NumberSignSymbol.
//
// If input parameter 'incomingNumSignSymbol' is judged to be
// invalid, this method will return an error.
//
func (nSignSymMolecule *numberSignSymbolMolecule) copyOut(
	incomingNumSignSymbol *NumberSignSymbol,
	errPrefDto *ePref.ErrPrefixDto) (
	newNumSignSymbol NumberSignSymbol,
	err error) {

	nSignSymMolecule.lock.Lock()

	defer nSignSymMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberSignSymbolMolecule.copyOut()",
		"")

	if err != nil {
		return newNumSignSymbol, err
	}

	if incomingNumSignSymbol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingNumSignSymbol' is a 'nil' pointer!\n",
			ePrefix.String())
		return newNumSignSymbol, err
	}

	_,
		err = numberSignSymbolElectron{}.ptr().
		testValidityOfNumSignSymbol(
			incomingNumSignSymbol,
			ePrefix.XCtx("incomingNumSignSymbol"))

	if err != nil {
		return newNumSignSymbol, err
	}

	lenRuneAry :=
		len(incomingNumSignSymbol.leadingNumSignChars)

	if lenRuneAry > 0 {

		newNumSignSymbol.leadingNumSignChars =
			make([]rune, lenRuneAry)

		copy(newNumSignSymbol.leadingNumSignChars,
			incomingNumSignSymbol.leadingNumSignChars)
	}

	lenRuneAry = len(incomingNumSignSymbol.trailingNumSignChars)

	if lenRuneAry > 0 {

		newNumSignSymbol.trailingNumSignChars =
			make([]rune, lenRuneAry)

		copy(newNumSignSymbol.trailingNumSignChars,
			incomingNumSignSymbol.trailingNumSignChars)
	}

	newNumSignSymbol.leadingNumSignFoundInNumber =
		incomingNumSignSymbol.leadingNumSignFoundInNumber

	newNumSignSymbol.trailingNumSignFoundInNumber =
		incomingNumSignSymbol.trailingNumSignFoundInNumber

	newNumSignSymbol.leadingNumSignFoundIndex =
		incomingNumSignSymbol.leadingNumSignFoundIndex

	newNumSignSymbol.trailingNumSignFoundIndex =
		incomingNumSignSymbol.trailingNumSignFoundIndex

	newNumSignSymbol.numSignType =
		incomingNumSignSymbol.numSignType

	newNumSignSymbol.numSignPosition =
		incomingNumSignSymbol.numSignPosition

	return newNumSignSymbol, err
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

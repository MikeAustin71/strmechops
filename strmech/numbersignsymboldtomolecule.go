package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numberSignSymbolDtoMolecule struct {
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
func (nSignSymMolecule *numberSignSymbolDtoMolecule) copyIn(
	targetNumSignSymbol *NumberSignSymbolDto,
	incomingNumSignSymbol *NumberSignSymbolDto,
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
		"numberSignSymbolDtoMolecule.copyIn()",
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
		err = numberSignSymbolDtoElectron{}.ptr().
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

	targetNumSignSymbol.numSignValueType =
		incomingNumSignSymbol.numSignValueType

	targetNumSignSymbol.numSymbolClass =
		NumericSymbolClass(0).NumberSign()

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
// of NumberSignSymbolDto.
//
// If input parameter 'incomingNumSignSymbol' is judged to be
// invalid, this method will return an error.
//
func (nSignSymMolecule *numberSignSymbolDtoMolecule) copyOut(
	incomingNumSignSymbol *NumberSignSymbolDto,
	errPrefDto *ePref.ErrPrefixDto) (
	newNumSignSymbol NumberSignSymbolDto,
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
		"numberSignSymbolDtoMolecule.copyOut()",
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
		err = numberSignSymbolDtoElectron{}.ptr().
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

	newNumSignSymbol.numSignValueType =
		incomingNumSignSymbol.numSignValueType

	newNumSignSymbol.numSymbolClass =
		NumericSymbolClass(0).NumberSign()

	newNumSignSymbol.numSignPosition =
		incomingNumSignSymbol.numSignPosition

	return newNumSignSymbol, err
}

// ptr - Returns a pointer to a new instance of
// numberSignSymbolDtoMolecule.
//
func (nSignSymMolecule numberSignSymbolDtoMolecule) ptr() *numberSignSymbolDtoMolecule {

	if nSignSymMolecule.lock == nil {
		nSignSymMolecule.lock = new(sync.Mutex)
	}

	nSignSymMolecule.lock.Lock()

	defer nSignSymMolecule.lock.Unlock()

	return &numberSignSymbolDtoMolecule{
		lock: new(sync.Mutex),
	}
}

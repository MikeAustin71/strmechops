package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numberSignSymbolElectron struct {
	lock *sync.Mutex
}

// emptyNumSignSymbol - Receives a pointer to an NumberSignSymbol
// object an proceeds to set all internal member variables to their
// zero or uninitialized state.
//
func (nSignSymElectron *numberSignSymbolElectron) emptyNumSignSymbol(
	numSignSymbol *NumberSignSymbol,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nSignSymElectron.lock == nil {
		nSignSymElectron.lock = new(sync.Mutex)
	}

	nSignSymElectron.lock.Lock()

	defer nSignSymElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberSignSymbolElectron.emptyNumSignSymbol()",
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

	numSignSymbol.leadingNumSignChars = nil
	numSignSymbol.trailingNumSignChars = nil
	numSignSymbol.symFoundInNumber = false
	numSignSymbol.numSignPosition = NumSignSymbolPosition(0).None()
	numSignSymbol.numSignType = NumericSignValueType(0).None()

	return nil
}

// equalNumSignSymbolObjects - Receives pointers to two
// NumberSignSymbol objects and proceeds to determine if their data
// values are equal in all respects.
//
// If the data values of the two NumberSignSymbol instances are
// equivalent, this method returns 'true'. If they are not equal,
// the method returns false.
//
func (nSignSymElectron *numberSignSymbolElectron) equalNumSignSymbolObjects(
	numSignSymbolOne *NumberSignSymbol,
	numSignSymbolTwo *NumberSignSymbol) (
	areEqual bool) {

	if nSignSymElectron.lock == nil {
		nSignSymElectron.lock = new(sync.Mutex)
	}

	nSignSymElectron.lock.Lock()

	defer nSignSymElectron.lock.Unlock()

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		numSignSymbolOne.leadingNumSignChars,
		numSignSymbolTwo.leadingNumSignChars) {
		return false
	}

	if !sMechPreon.equalRuneArrays(
		numSignSymbolOne.trailingNumSignChars,
		numSignSymbolTwo.trailingNumSignChars) {
		return false
	}

	if numSignSymbolOne.symFoundInNumber !=
		numSignSymbolTwo.symFoundInNumber {
		return false
	}

	if numSignSymbolOne.numSignPosition !=
		numSignSymbolTwo.numSignPosition {
		return false
	}

	if numSignSymbolOne.numSignType !=
		numSignSymbolTwo.numSignType {
		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// numberSignSymbolElectron.
//
func (nSignSymElectron numberSignSymbolElectron) ptr() *numberSignSymbolElectron {

	if nSignSymElectron.lock == nil {
		nSignSymElectron.lock = new(sync.Mutex)
	}

	nSignSymElectron.lock.Lock()

	defer nSignSymElectron.lock.Unlock()

	return &numberSignSymbolElectron{
		lock: new(sync.Mutex),
	}
}

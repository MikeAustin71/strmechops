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

// testValidityOfNumSignSymbol - Performs a diagnostic review of
// the input parameter 'numSignSymbol', an instance of
// NumberSignSymbol. The purpose of this diagnostic review is to
// determine whether this NumberSignSymbol instance is valid in all
// respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numSignSymbol       *NumberSignSymbol
//     - A pointer to an instance of NumberSignSymbol. This object
//       will be evaluated to determine whether or not it is a
//       valid instance.
//
//
//  errPrefDto          *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods listed
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
//  isValid             bool
//     - This returned boolean value will signal whether the input
//       parameter, 'numSignSymbol', is valid, or not. If the
//       NumberSignSymbol object contains valid data, this method
//       returns 'true'. If the data is invalid, this method will
//       return 'false'.
//
//  err                 error
//     - If the input parameter object, 'errPrefixDto', contains
//       invalid data, a detailed error message will be returned
//       identifying the invalid data item.
//
//       If the input parameter object, 'numSignSymbol', is valid,
//       this error parameter will be set to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type ('err') will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (nSignSymElectron *numberSignSymbolElectron) testValidityOfNumSignSymbol(
	numSignSymbol *NumberSignSymbol,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if nSignSymElectron.lock == nil {
		nSignSymElectron.lock = new(sync.Mutex)
	}

	nSignSymElectron.lock.Lock()

	defer nSignSymElectron.lock.Unlock()

	isValid = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberSignSymbolElectron.testValidityOfNumSignSymbol()",
		"")

	if err != nil {
		return isValid, err
	}

	if numSignSymbol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSignSymbol' is a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	lenLeadingChars :=
		len(numSignSymbol.leadingNumSignChars)

	lenTrailingChars :=
		len(numSignSymbol.trailingNumSignChars)

	if lenLeadingChars == 0 &&
		lenTrailingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Both Leading and Trailing Number Sign Characters are empty!\n"+
			"This NumberSignSymbol object is invalid!\n",
			ePrefix.String())

		return isValid, err
	}

	if !numSignSymbol.numSignType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"The Number Sign Type for this NumberSignSymbol object\n"+
			"is invalid!\n"+
			"Number Sign Type='%v'\n"+
			"Number Sign Type Integer Value='%v'\n",
			ePrefix.String(),
			numSignSymbol.numSignType.String(),
			numSignSymbol.numSignType.XValueInt())

		return isValid, err
	}

	if !numSignSymbol.numSignPosition.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"The Number Sign Position for this NumberSignSymbol object\n"+
			"is invalid!\n"+
			"Number Sign Position='%v'\n"+
			"Number Sign Position Integer Value='%v'\n",
			ePrefix.String(),
			numSignSymbol.numSignPosition.String(),
			numSignSymbol.numSignPosition.XValueInt())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).Before() &&
		lenTrailingChars > 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbol object is invalid!\n"+
			"The Number Sign Symbol Position == 'Before' but\n"+
			"the Trailing Characters Array is populated!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).Before() &&
		lenLeadingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbol object is invalid!\n"+
			"The Number Sign Symbol Position == 'Before' but\n"+
			"the Leading Characters Array is empty!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).After() &&
		lenTrailingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbol object is invalid!\n"+
			"The Number Sign Symbol Position == 'After' but\n"+
			"the Trailing Characters Array is empty!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).After() &&
		lenLeadingChars > 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbol object is invalid!\n"+
			"The Number Sign Symbol Position == 'After' but\n"+
			"the Leading Characters Array is populated!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).BeforeAndAfter() &&
		lenTrailingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbol object is invalid!\n"+
			"The Number Sign Symbol Position == 'BeforeAndAfter' but\n"+
			"the Trailing Characters Array is empty!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).BeforeAndAfter() &&
		lenLeadingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbol object is invalid!\n"+
			"The Number Sign Symbol Position == 'BeforeAndAfter' but\n"+
			"the Leading Characters Array is empty!\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true

	return isValid, err
}

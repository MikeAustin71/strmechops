package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numberSignSymbolDtoElectron struct {
	lock *sync.Mutex
}

// emptyNumSignSymbol - Receives a pointer to a NumberSignSymbolDto
// object and proceeds to set all internal member variables to their
// zero or uninitialized state.
//
func (nSignSymElectron *numberSignSymbolDtoElectron) emptyNumSignSymbol(
	numSignSymbol *NumberSignSymbolDto,
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
		"numberSignSymbolDtoElectron.emptyNumSignSymbol()",
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
	numSignSymbol.leadingNumSignFoundInNumber = false
	numSignSymbol.trailingNumSignFoundInNumber = false
	numSignSymbol.leadingNumSignFoundIndex = 0
	numSignSymbol.trailingNumSignFoundIndex = 0
	numSignSymbol.numSignPosition = NumSignSymbolPosition(0).None()
	numSignSymbol.numSignValueType = NumericSignValueType(0).None()
	numSignSymbol.numSymbolClass = NumericSymbolClass(0).None()
	numSignSymbol.numSymbolDisplayMode = NumSignSymbolDisplayMode(0).None()

	return nil
}

// equalNumSignRuneArrays - Receives a pointer to two
// NumberSignSymbolDto objects and proceeds to compare their number
// sign rune arrays. If the two number sign rune arrays are
// equivalent, this method returns 'true'.
//
// If the two number sign rune arrays differ in any respect, this
// method will return false.
//
func (nSignSymElectron *numberSignSymbolDtoElectron) equalNumSignRuneArrays(
	numSignSymbolOne *NumberSignSymbolDto,
	numSignSymbolTwo *NumberSignSymbolDto) (
	areEqual bool) {

	if nSignSymElectron.lock == nil {
		nSignSymElectron.lock = new(sync.Mutex)
	}

	nSignSymElectron.lock.Lock()

	defer nSignSymElectron.lock.Unlock()

	if numSignSymbolOne == nil {
		return false
	}

	if numSignSymbolTwo == nil {
		return false
	}

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

	return true
}

// equalNumSignSymbolObjects - Receives pointers to two
// NumberSignSymbolDto objects and proceeds to determine if their data
// values are equal in all respects.
//
// If the data values of the two NumberSignSymbolDto instances are
// equivalent, this method returns 'true'. If they are not equal,
// the method returns false.
//
func (nSignSymElectron *numberSignSymbolDtoElectron) equalNumSignSymbolObjects(
	numSignSymbolOne *NumberSignSymbolDto,
	numSignSymbolTwo *NumberSignSymbolDto) (
	areEqual bool) {

	if nSignSymElectron.lock == nil {
		nSignSymElectron.lock = new(sync.Mutex)
	}

	nSignSymElectron.lock.Lock()

	defer nSignSymElectron.lock.Unlock()

	if numSignSymbolOne == nil {
		return false
	}

	if numSignSymbolTwo == nil {
		return false
	}

	sMechPreon := strMechPreon{}

	areEqual = false

	if !sMechPreon.equalRuneArrays(
		numSignSymbolOne.leadingNumSignChars,
		numSignSymbolTwo.leadingNumSignChars) {
		return areEqual
	}

	if !sMechPreon.equalRuneArrays(
		numSignSymbolOne.trailingNumSignChars,
		numSignSymbolTwo.trailingNumSignChars) {
		return areEqual
	}

	if numSignSymbolOne.leadingNumSignFoundInNumber !=
		numSignSymbolTwo.leadingNumSignFoundInNumber {
		return areEqual
	}

	if numSignSymbolOne.trailingNumSignFoundInNumber !=
		numSignSymbolTwo.trailingNumSignFoundInNumber {
		return areEqual
	}

	if numSignSymbolOne.leadingNumSignFoundIndex !=
		numSignSymbolTwo.leadingNumSignFoundIndex {
		return areEqual
	}

	if numSignSymbolOne.trailingNumSignFoundIndex !=
		numSignSymbolTwo.trailingNumSignFoundIndex {
		return areEqual
	}

	if numSignSymbolOne.numSignPosition !=
		numSignSymbolTwo.numSignPosition {
		return areEqual
	}

	if numSignSymbolOne.numSignValueType !=
		numSignSymbolTwo.numSignValueType {
		return areEqual
	}

	if numSignSymbolOne.numSymbolClass !=
		numSignSymbolTwo.numSymbolClass {
		return areEqual
	}

	if numSignSymbolOne.numSymbolDisplayMode !=
		numSignSymbolTwo.numSymbolDisplayMode {
		return areEqual
	}

	areEqual = true

	return areEqual
}

// ptr - Returns a pointer to a new instance of
// numberSignSymbolDtoElectron.
//
func (nSignSymElectron numberSignSymbolDtoElectron) ptr() *numberSignSymbolDtoElectron {

	if nSignSymElectron.lock == nil {
		nSignSymElectron.lock = new(sync.Mutex)
	}

	nSignSymElectron.lock.Lock()

	defer nSignSymElectron.lock.Unlock()

	return &numberSignSymbolDtoElectron{
		lock: new(sync.Mutex),
	}
}

// testValidityOfNumSignSymbol - Performs a diagnostic review of
// the input parameter 'numSignSymbol', an instance of
// NumberSignSymbolDto. The purpose of this diagnostic review is to
// determine whether this NumberSignSymbolDto instance is valid in all
// respects.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numSignSymbol       *NumberSignSymbolDto
//     - A pointer to an instance of NumberSignSymbolDto. This object
//       will be evaluated to determine whether it is a valid
//       instance.
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
//       NumberSignSymbolDto object contains valid data, this method
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
func (nSignSymElectron *numberSignSymbolDtoElectron) testValidityOfNumSignSymbol(
	numSignSymbol *NumberSignSymbolDto,
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
		"numberSignSymbolDtoElectron.testValidityOfNumSignSymbol()",
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
			"This NumberSignSymbolDto object is invalid!\n",
			ePrefix.String())

		return isValid, err
	}

	if !numSignSymbol.numSignPosition.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"The Number Sign Position for this NumberSignSymbolDto object\n"+
			"is invalid!\n"+
			"Number Sign Position='%v'\n"+
			"Number Sign Position Integer Value='%v'\n",
			ePrefix.String(),
			numSignSymbol.numSignPosition.String(),
			numSignSymbol.numSignPosition.XValueInt())

		return isValid, err
	}

	if !numSignSymbol.numSignValueType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"The Number Sign Type for this NumberSignSymbolDto object\n"+
			"is invalid!\n"+
			"Number Sign Type='%v'\n"+
			"Number Sign Type Integer Value='%v'\n",
			ePrefix.String(),
			numSignSymbol.numSignValueType.String(),
			numSignSymbol.numSignValueType.XValueInt())

		return isValid, err
	}

	numSignSymbol.numSymbolClass = NumericSymbolClass(0).NumberSign()

	if !numSignSymbol.numSymbolDisplayMode.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"The Number Sign Symbol Display Mode for this NumberSignSymbolDto\n"+
			"object is invalid!\n"+
			"Number Sign Symbol Display Mode='%v'\n"+
			"Number Sign Symbol Display Mode Integer Value='%v'\n",
			ePrefix.String(),
			numSignSymbol.numSymbolDisplayMode.String(),
			numSignSymbol.numSymbolDisplayMode.XValueInt())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).Before() &&
		lenTrailingChars > 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbolDto object is invalid!\n"+
			"The Number Sign Symbol Position == 'Before' but\n"+
			"the Trailing Characters Array is populated!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).Before() &&
		lenLeadingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbolDto object is invalid!\n"+
			"The Number Sign Symbol Position == 'Before' but\n"+
			"the Leading Characters Array is empty!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).After() &&
		lenTrailingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbolDto object is invalid!\n"+
			"The Number Sign Symbol Position == 'After' but\n"+
			"the Trailing Characters Array is empty!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).After() &&
		lenLeadingChars > 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbolDto object is invalid!\n"+
			"The Number Sign Symbol Position == 'After' but\n"+
			"the Leading Characters Array is populated!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).BeforeAndAfter() &&
		lenTrailingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbolDto object is invalid!\n"+
			"The Number Sign Symbol Position == 'BeforeAndAfter' but\n"+
			"the Trailing Characters Array is empty!\n",
			ePrefix.String())

		return isValid, err
	}

	if numSignSymbol.numSignPosition ==
		NumSignSymbolPosition(0).BeforeAndAfter() &&
		lenLeadingChars == 0 {

		err = fmt.Errorf("%v\n"+
			"This NumberSignSymbolDto object is invalid!\n"+
			"The Number Sign Symbol Position == 'BeforeAndAfter' but\n"+
			"the Leading Characters Array is empty!\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true

	return isValid, err
}

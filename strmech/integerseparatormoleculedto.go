package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorDtoMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies the data fields from input parameter
// 'incomingNStrIntSeparator' to input parameter
// 'targetNStrIntSeparator'.
//
// Be advised - All data fields in 'targetNStrIntSeparator'
// will be overwritten.
//
// If input parameter 'incomingNStrIntSeparator' is judged
// to be invalid, this method will return an error.
//
func (nStrIntSepMolecule *integerSeparatorDtoMolecule) copyIn(
	targetNStrIntSeparator *IntegerSeparatorDto,
	incomingNStrIntSeparator *IntegerSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMolecule."+
			"copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetNStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetNStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingNStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingNStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err =
		integerSeparatorDtoQuark{}.ptr().
			testValidityOfNumStrIntSeparator(
				incomingNStrIntSeparator,
				ePrefix.XCtx(
					"Testing validity of 'incomingNStrIntSeparator'."))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&targetNStrIntSeparator.intSeparatorChars,
		&incomingNStrIntSeparator.intSeparatorChars,
		true,
		ePrefix.XCtx(
			"incomingNStrIntSeparator.intSeparatorChars->"+
				"targetNStrIntSeparator.intSeparatorChars"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyUnsignedIntArrays(
		&targetNStrIntSeparator.intGroupingSequence,
		&incomingNStrIntSeparator.intGroupingSequence,
		true,
		ePrefix.XCtx(
			"incomingNStrIntSeparator.intGroupingSequence->"+
				"targetNStrIntSeparator.intGroupingSequence"))

	if err != nil {
		return err
	}

	targetNStrIntSeparator.restartIntGroupingSequence =
		incomingNStrIntSeparator.restartIntGroupingSequence

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'numStrIntSeparator' styled as a new instance
// of IntegerSeparatorDto.
//
// If input parameter 'numStrIntSeparator' is judged to be
// invalid, this method will return an error.
//
func (nStrIntSepMolecule *integerSeparatorDtoMolecule) copyOut(
	numStrIntSeparator *IntegerSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (
	newNumSrIntSeparator IntegerSeparatorDto,
	err error) {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMolecule."+
			"copyOut()",
		"")

	if err != nil {
		return newNumSrIntSeparator, err
	}

	if numStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return newNumSrIntSeparator, err
	}

	_,
		err =
		integerSeparatorDtoQuark{}.ptr().
			testValidityOfNumStrIntSeparator(
				numStrIntSeparator,
				ePrefix.XCtx(
					"Testing validity of 'numStrIntSeparator'."))

	if err != nil {
		return newNumSrIntSeparator, err
	}

	lIntSepChars :=
		len(numStrIntSeparator.intSeparatorChars)

	if lIntSepChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'numStrIntSeparator.intSeparatorChars' "+
			"is invalid!\n"+
			"'numStrIntSeparator.intSeparatorChars' is a zero "+
			"length array.\n",
			ePrefix.XCtxEmpty().String())

		return newNumSrIntSeparator, err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&newNumSrIntSeparator.intSeparatorChars,
		&numStrIntSeparator.intSeparatorChars,
		true,
		ePrefix.XCtx(
			"numStrIntSeparator.intSeparatorChars->"+
				"newNumSrIntSeparator.intSeparatorChars"))

	if err != nil {
		return newNumSrIntSeparator, err
	}

	err = sMechPreon.copyUnsignedIntArrays(
		&newNumSrIntSeparator.intGroupingSequence,
		&numStrIntSeparator.intGroupingSequence,
		true,
		ePrefix.XCtx(
			"numStrIntSeparator.intGroupingSequence->"+
				"newNumSrIntSeparator.intGroupingSequence"))

	if err != nil {
		return newNumSrIntSeparator, err
	}

	newNumSrIntSeparator.restartIntGroupingSequence =
		numStrIntSeparator.restartIntGroupingSequence

	newNumSrIntSeparator.lock = new(sync.Mutex)

	return newNumSrIntSeparator, err
}

// equal - Receives two IntegerSeparatorDto objects and proceeds to
// determine whether all data elements in the first object are
// equal to corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrIntSepOne       *IntegerSeparatorDto
//     - A pointer to the first IntegerSeparatorDto object. This
//       method will compare all data elements in this object to
//       corresponding data elements in the second
//       IntegerSeparatorDto object in order determine equivalency.
//
//
//  nStrIntSepTwo       *IntegerSeparatorDto
//     - A pointer to the second IntegerSeparatorDto object. This
//       method will compare all data elements in the first
//       IntegerSeparatorDto object to corresponding data elements in
//       this second IntegerSeparatorDto object in order determine
//       equivalency.
//
//
//  errPrefDto          *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isEqual             bool
//     - If all the data elements in 'nStrIntSepOne' are equal to
//       all the corresponding data elements in 'nStrIntSepTwo',
//       this return parameter will be set to 'true'. If all the
//       data elements are NOT equal, this return parameter will be
//       set to 'false'.
//
//
//  err                 error
//     - If all the data elements in 'nStrIntSepOne' are equal to
//       all the corresponding data elements in 'nStrIntSepTwo',
//       this return parameter will be set to 'nil'.
//
//       If the corresponding data elements are not equal, a
//       detailed error message identifying the unequal elements
//       will be returned.
//
func (nStrIntSepMolecule *integerSeparatorDtoMolecule) equal(
	nStrIntSepOne *IntegerSeparatorDto,
	nStrIntSepTwo *IntegerSeparatorDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isEqual bool,
	err error) {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMolecule."+
			"equal()",
		"")

	if err != nil {
		return isEqual, err
	}

	isEqual = false

	if nStrIntSepOne == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSepOne' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if nStrIntSepTwo == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSepTwo' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return isEqual, err
	}

	if nStrIntSepOne.intSeparatorChars == nil &&
		nStrIntSepTwo.intSeparatorChars != nil {
		err = fmt.Errorf("%v\n"+
			"Error: nStrIntSepOne.intSeparatorChars == nil\n"+
			"nStrIntSepTwo.intSeparatorChars != nil\n"+
			"nStrIntSepOne.intSeparatorChars='nil'\n"+
			"nStrIntSepTwo.intSeparatorChars='%v'\n",
			ePrefix.String(),
			string(nStrIntSepTwo.intSeparatorChars))

		return isEqual, err
	}

	if nStrIntSepOne.intGroupingSequence != nil &&
		nStrIntSepTwo.intGroupingSequence == nil {
		err = fmt.Errorf("%v\n"+
			"Error: nStrIntSepOne.intGroupingSequence != nil\n"+
			"nStrIntSepTwo.intGroupingSequence == nil\n"+
			"nStrIntSepOne.intGroupingSequence='%v'\n"+
			"nStrIntSepTwo.intGroupingSequence='nil'\n",
			ePrefix.String(),
			nStrIntSepOne.intGroupingSequence)

		return isEqual, err
	}

	sMechPreon := strMechPreon{}

	isEqual = sMechPreon.equalRuneArrays(
		nStrIntSepOne.intSeparatorChars,
		nStrIntSepTwo.intSeparatorChars)

	if !isEqual {
		return isEqual, err
	}

	isEqual = sMechPreon.equalUintArrays(
		nStrIntSepOne.intGroupingSequence,
		nStrIntSepTwo.intGroupingSequence)

	if !isEqual {
		return isEqual, err
	}

	if nStrIntSepOne.restartIntGroupingSequence !=
		nStrIntSepTwo.restartIntGroupingSequence {

		isEqual = false

		err = fmt.Errorf("%v\n"+
			"Error: nStrIntSepOne.restartIntGroupingSequence !=\n"+
			"nStrIntSepTwo.restartIntGroupingSequence\n"+
			"nStrIntSepOne.restartIntGroupingSequence='%v'\n"+
			"nStrIntSepTwo.restartIntGroupingSequence='%v'\n",
			ePrefix.String(),
			nStrIntSepOne.restartIntGroupingSequence,
			nStrIntSepTwo.restartIntGroupingSequence)

		return isEqual, err
	}

	isEqual = true

	return isEqual, err
}

// ptr - Returns a pointer to a new instance of
// integerSeparatorDtoMolecule.
func (nStrIntSepMolecule integerSeparatorDtoMolecule) ptr() *integerSeparatorDtoMolecule {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	newIntSepMolecule := new(integerSeparatorDtoMolecule)

	newIntSepMolecule.lock = new(sync.Mutex)

	return newIntSepMolecule
}

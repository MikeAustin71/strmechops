package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numberIntSeparatorMolecule struct {
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
func (nStrIntSepMolecule *numberIntSeparatorMolecule) copyIn(
	targetNStrIntSeparator *NumberIntSeparator,
	incomingNStrIntSeparator *NumberIntSeparator,
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
		"numberIntSeparatorMolecule."+
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
		numberIntSeparatorQuark{}.ptr().
			testValidityOfNumStrIntSeparator(
				incomingNStrIntSeparator,
				ePrefix.XCtx(
					"Testing validity of 'incomingNStrIntSeparator'."))

	if err != nil {
		return err
	}

	lIntSepChars :=
		len(incomingNStrIntSeparator.intSeparatorChars)

	if lIntSepChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'incomingNStrIntSeparator.intSeparatorChars' "+
			"is invalid!\n"+
			"'incomingNStrIntSeparator.intSeparatorChars' is a zero "+
			"length array.\n",
			ePrefix.XCtxEmpty().String())

		return err
	}

	targetNStrIntSeparator.intSeparatorChars = nil

	targetNStrIntSeparator.intSeparatorChars =
		make([]rune, lIntSepChars, lIntSepChars+5)

	copy(targetNStrIntSeparator.intSeparatorChars,
		incomingNStrIntSeparator.intSeparatorChars)

	targetNStrIntSeparator.intSeparatorGrouping =
		incomingNStrIntSeparator.intSeparatorGrouping

	targetNStrIntSeparator.intSeparatorRepetitions =
		incomingNStrIntSeparator.intSeparatorRepetitions

	targetNStrIntSeparator.restartIntGroupingSequence =
		incomingNStrIntSeparator.restartIntGroupingSequence

	return err
}

// copyOut - Returns a deep copy of input parameter
// 'numStrIntSeparator' styled as a new instance
// of NumberIntSeparator.
//
// If input parameter 'numStrIntSeparator' is judged to be
// invalid, this method will return an error.
//
func (nStrIntSepMolecule *numberIntSeparatorMolecule) copyOut(
	numStrIntSeparator *NumberIntSeparator,
	errPrefDto *ePref.ErrPrefixDto) (
	newNumSrIntSeparator NumberIntSeparator,
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
		"numberIntSeparatorMolecule."+
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
		numberIntSeparatorQuark{}.ptr().
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

	newNumSrIntSeparator.intSeparatorChars =
		make([]rune, lIntSepChars, lIntSepChars+5)

	copy(newNumSrIntSeparator.intSeparatorChars,
		numStrIntSeparator.intSeparatorChars)

	newNumSrIntSeparator.intSeparatorGrouping =
		numStrIntSeparator.intSeparatorGrouping

	newNumSrIntSeparator.intSeparatorRepetitions =
		numStrIntSeparator.intSeparatorRepetitions

	newNumSrIntSeparator.restartIntGroupingSequence =
		numStrIntSeparator.restartIntGroupingSequence

	newNumSrIntSeparator.lock = new(sync.Mutex)

	return newNumSrIntSeparator, err
}

// equal - Receives two NumberIntSeparator objects and proceeds to
// determine whether all data elements in the first object are
// equal to corresponding data elements in the second object.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrIntSepOne       *NumberIntSeparator
//     - A pointer to the first NumberIntSeparator object. This
//       method will compare all data elements in this object to
//       corresponding data elements in the second
//       NumberIntSeparator object in order determine equivalency.
//
//
//  nStrIntSepTwo       *NumberIntSeparator
//     - A pointer to the second NumberIntSeparator object. This
//       method will compare all data elements in the first
//       NumberIntSeparator object to corresponding data elements in
//       this second NumberIntSeparator object in order determine
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
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  equal             bool
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
func (nStrIntSepMolecule *numberIntSeparatorMolecule) equal(
	nStrIntSepOne *NumberIntSeparator,
	nStrIntSepTwo *NumberIntSeparator,
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
		"numberIntSeparatorMolecule."+
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

	if nStrIntSepOne.intSeparatorChars != nil &&
		nStrIntSepTwo.intSeparatorChars == nil {
		err = fmt.Errorf("%v\n"+
			"Error: nStrIntSepOne.intSeparatorChars != nil\n"+
			"nStrIntSepTwo.intSeparatorChars == nil\n"+
			"nStrIntSepOne.intSeparatorChars='%v'\n"+
			"nStrIntSepTwo.intSeparatorChars='nil'\n",
			ePrefix.String(),
			string(nStrIntSepOne.intSeparatorChars))

		return isEqual, err
	}

	lenIntSeps1 := len(nStrIntSepOne.intSeparatorChars)

	if lenIntSeps1 != len(nStrIntSepTwo.intSeparatorChars) {

		err = fmt.Errorf("%v\n"+
			"nStrIntSepOne.intSeparatorChars !="+
			"nStrIntSepTwo.intSeparatorChars\n"+
			"nStrIntSepOne.intSeparatorChars='%v'\n"+
			"nStrIntSepTwo.intSeparatorChars='%v'\n",
			ePrefix.String(),
			string(nStrIntSepOne.intSeparatorChars),
			string(nStrIntSepTwo.intSeparatorChars))

		return isEqual, err
	}

	for i := 0; i < lenIntSeps1; i++ {
		if nStrIntSepOne.intSeparatorChars[i] !=
			nStrIntSepTwo.intSeparatorChars[i] {
			err = fmt.Errorf("%v\n"+
				"nStrIntSepOne.intSeparatorChars !="+
				"nStrIntSepTwo.intSeparatorChars\n"+
				"nStrIntSepOne.intSeparatorChars='%v'\n"+
				"nStrIntSepTwo.intSeparatorChars='%v'\n"+
				"Reference index '%v'\n",
				ePrefix.String(),
				string(nStrIntSepOne.intSeparatorChars),
				string(nStrIntSepTwo.intSeparatorChars),
				i)

			return isEqual, err
		}

	}

	if nStrIntSepOne.intSeparatorGrouping !=
		nStrIntSepTwo.intSeparatorGrouping {

		err = fmt.Errorf("%v\n"+
			"Error: nStrIntSepOne.intSeparatorGrouping !=\n"+
			"nStrIntSepTwo.intSeparatorGrouping\n"+
			"nStrIntSepOne.intSeparatorGrouping='%v'\n"+
			"nStrIntSepTwo.intSeparatorGrouping='%v'\n",
			ePrefix.String(),
			nStrIntSepOne.intSeparatorGrouping,
			nStrIntSepTwo.intSeparatorGrouping)

		return isEqual, err
	}

	if nStrIntSepOne.intSeparatorRepetitions !=
		nStrIntSepTwo.intSeparatorRepetitions {

		err = fmt.Errorf("%v\n"+
			"Error: nStrIntSepOne.intSeparatorRepetitions !=\n"+
			"nStrIntSepTwo.intSeparatorRepetitions\n"+
			"nStrIntSepOne.intSeparatorRepetitions='%v'\n"+
			"nStrIntSepTwo.intSeparatorRepetitions='%v'\n",
			ePrefix.String(),
			nStrIntSepOne.intSeparatorRepetitions,
			nStrIntSepTwo.intSeparatorRepetitions)

		return isEqual, err
	}

	if nStrIntSepOne.restartIntGroupingSequence !=
		nStrIntSepTwo.restartIntGroupingSequence {

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
// numberIntSeparatorMolecule.
func (nStrIntSepMolecule numberIntSeparatorMolecule) ptr() *numberIntSeparatorMolecule {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	newIntSepMolecule := new(numberIntSeparatorMolecule)

	newIntSepMolecule.lock = new(sync.Mutex)

	return newIntSepMolecule
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type integerSeparatorDtoMolecule struct {
	lock *sync.Mutex
}

// applyIntSeparators - Receives an array of runes which consists
// entirely of integer digit characters '0' to '9' inclusive. Input
// parameter, 'nStrIntSeparator' supplies the information and
// format parameters necessary to insert integer separators into
// the integer digits supplied by input parameter, 'pureNumRunes'.
//
// Integer separators are often referred to as thousands separators.
// The result is returned as an runes correctly formatted with
// integer separators.
//
// Example:
//  pureNumRunes = 123456789012345
//  integer separator character = ','
//  integer grouping for thousands = 3
//  result = 123,456,789,012,345
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  nStrIntSeparator           *IntegerSeparatorDto
//     - A pointer to an IntegerSeparatorDto object which contains
//       the integer separation format parameters which will be
//       used to insert integer separators.
//
//       Integer separators consist of a character, or series of
//       characters, used to separate integer digits in a number
//       string. These characters are commonly known as the
//       'thousands separator'. A 'thousands separator' is used to
//       separate groups of integer digits to the left of the
//       decimal separator (a.k.a. decimal point). In the United
//       States, the standard integer digits separator is the
//       single comma character (',').
//             United States Example:  1,000,000,000
//
//       In many European countries, a single period ('.') is used
//       as the integer separator character.
//             European Example: 1.000.000.000
//
//       Other countries and cultures use spaces, apostrophes or
//       multiple characters to separate integers.
//
//       For additional details, reference the source code
//       documentation for type  IntegerSeparatorDto.
//
//
//  pureNumRunes               []rune
//     - An array of runes consisting entirely of integer digit
//       characters from '0' to '9' inclusive. If any character
//       within this array is NOT an integer digit, this method
//       will return an error.
//
//
//  errPrefDto                 *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  numStrWithIntSeps          []rune
//     - If this method completes successfully, an array of runes
//       will be returned containing the integer digits supplied by
//       input parameter 'pureNumRunes' properly formatted with
//       integer digit separators (a.k.a. thousands separators).
//       Example:
//         pureNumRunes = 123456789012345
//         integer separator character = ','
//         integer grouping for thousands = 3
//         numStrWithIntSeps = 123,456,789,012,345
//
//
//  err                        error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'.
//
//       If errors are encountered during processing, the returned
//       error Type will encapsulate an error message. This
//       returned error message will incorporate the method chain
//       and text passed by input parameter, 'errPrefDto'. The
//       'errPrefDto' text will be attached to the beginning of the
//       error message.
//
func (nStrIntSepMolecule *integerSeparatorDtoMolecule) applyIntSeparators(
	nStrIntSeparator *IntegerSeparatorDto,
	pureNumRunes []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	numStrWithIntSeps []rune,
	err error) {

	if nStrIntSepMolecule.lock == nil {
		nStrIntSepMolecule.lock = new(sync.Mutex)
	}

	nStrIntSepMolecule.lock.Lock()

	defer nStrIntSepMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	numStrWithIntSeps = nil

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"integerSeparatorDtoMolecule."+
			"copyIn()",
		"")

	if err != nil {
		return numStrWithIntSeps, err
	}

	if nStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSeparator' (*IntegerSeparatorDto) is invalid!\n"+
			"'nStrIntSeparator' is a 'nil' pointer.\n",
			ePrefix.String())

		return numStrWithIntSeps, err
	}

	if pureNumRunes == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumRunes' ([]rune) is invalid!\n"+
			"'pureNumRunes' is 'nil'!\n",
			ePrefix.String())

		return numStrWithIntSeps, err
	}

	lenRawNumRunes := len(pureNumRunes)

	if lenRawNumRunes == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumRunes' is invalid!\n"+
			"'pureNumRunes' is a zero length array.\n",
			ePrefix)

		return numStrWithIntSeps, err
	}

	_,
		err = integerSeparatorDtoQuark{}.ptr().
		testValidityOfNumStrIntSeparator(
			nStrIntSeparator,
			ePrefix.XCtx("nStrIntSeparator->"))

	if err != nil {
		return numStrWithIntSeps, err
	}

	lenIGrpSeq := len(nStrIntSeparator.intGroupingSequence)

	if lenIGrpSeq == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerGroupingSequence' is invalid!\n"+
			"'integerGroupingSequence' is a ZERO length array.\n",
			ePrefix)

		return numStrWithIntSeps, err
	}

	lenIntSeparatorChars := len(nStrIntSeparator.intSeparatorChars)

	if lenIntSeparatorChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSeparator.intSeparatorChars' is invalid!\n"+
			"The Integer Separator Characters rune array is a ZERO length array.\n",
			ePrefix)

		return numStrWithIntSeps, err
	}

	lenOutRunes := lenRawNumRunes * 2 * lenIntSeparatorChars

	outRunes := make([]rune, lenOutRunes)

	outIdx := lenOutRunes - 1

	groupCnt := uint(0)
	maxGroupCnt := nStrIntSeparator.intGroupingSequence[0]
	currGroupCntIdx := 0
	lastGroupCntIdx := lenIGrpSeq - 1

	for i := lenRawNumRunes - 1; i >= 0; i-- {

		if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9' {

			groupCnt++
			outRunes[outIdx] = pureNumRunes[i]
			outIdx--

			if groupCnt == maxGroupCnt && i != 0 {

				groupCnt = 0

				copy(outRunes[outIdx:], nStrIntSeparator.intSeparatorChars)
				outIdx = outIdx - lenIntSeparatorChars

				if currGroupCntIdx+1 > lastGroupCntIdx {

					maxGroupCnt =
						nStrIntSeparator.intGroupingSequence[currGroupCntIdx]

				} else {

					currGroupCntIdx++

					maxGroupCnt = nStrIntSeparator.intGroupingSequence[currGroupCntIdx]

				}

			} // End of if groupCnt == maxGroupCnt && i != 0

			// End Of if pureNumRunes[i] >= '0' && pureNumRunes[i] <= '9'
		} else {

			err = fmt.Errorf("%v\n"+
				"Error: Input parameter 'pureNumRunes' is invalid!\n"+
				" The 'pureNumRunes' contains a character which is an integer digit!\n"+
				"pureNumRunes[%v] = '%v'\n",
				ePrefix.String(),
				i,
				string(pureNumRunes[i]))

			return numStrWithIntSeps, err
		}

	} // End of for i := lenRawNumRunes - 1; i >= 0; i--

	outputLen := outIdx + 1

	numStrWithIntSeps = make([]rune, outputLen)

	charsCopied := copy(numStrWithIntSeps, outRunes[outIdx:])

	if charsCopied != outputLen {
		err = fmt.Errorf("%v\n"+
			"Calculation Error: The number of output characters copied is incorrect!\n"+
			"Output Length is '%v'\n"+
			"Characters Coped is '%v'\n"+
			"outRunes[outIdx:] = '%v'\n",
			ePrefix.String(),
			outputLen,
			charsCopied,
			string(outRunes[outIdx:]))
	}

	return numStrWithIntSeps, err
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

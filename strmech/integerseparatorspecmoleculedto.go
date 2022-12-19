package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"sync"
)

type integerSeparatorSpecMolecule struct {
	lock *sync.Mutex
}

// applyIntSeparators - Inserts integer separators into
// an array of integer digits.
//
// An integer separator is a character, or series of characters,
// used to separate integer digits in a number string. These
// characters are commonly known as the 'thousands separator'.
// A 'thousands separator' is used to separate groups of integer
// digits to the left of the decimal separator (a.k.a. decimal
// point).
//
// In the United States, the standard integer digits separator is
// the single comma character (',').
//
//	United States Example:  1,000,000,000
//
// In many European countries, a single period ('.') is used as
// the integer separator character.
//
//	European Example: 1.000.000.000
//
// Other countries and cultures use spaces, apostrophes or
// multiple characters to separate integers.
//
// This method receives an array of runes which consists
// entirely of integer digit characters '0' (0x30) to '9' (0x39)
// inclusive. Input parameter, 'nStrIntSeparator' supplies the
// information and format parameters necessary to insert integer
// separators into the sequence of integer digits supplied by input
// parameter, 'pureNumRunes'. The result is a returned rune array
// consisting of integer digits properly formatted with integer
// separators.
//
// Example:
//
//	pureNumRunes = 123456789012345
//	integer separator character = ','
//	integer grouping for thousands = 3
//	result = 123,456,789,012,345
//
// The complexity inherent in the IntegerSeparatorSpec type, used as
// an input parameter, is necessary in order to provide support for
// multinational and multicultural integer separation algorithms.
// For more information on integer grouping sequence, reference the
// source documentation for type, IntegerSeparatorSpec.
//
// Be advised - This method will never return a number sign for the
// returned formatted integer separated numeric value. The numeric
// sign value (positive or negative) of the returned numeric value
// (result) must be tracked externally.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrIntSeparator           *IntegerSeparatorSpec
//
//		A pointer to an IntegerSeparatorSpec object which contains
//		the integer separation format parameters which will be
//		used to insert integer separators.
//
//		Integer separators consist of a character, or series of
//		characters, used to separate integer digits in a number
//		string. These characters are commonly known as the
//		'thousands separator'. A 'thousands separator' is used to
//		separate groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the United
//		States, the standard integer digits separator is the
//		single comma character (',').
//
//		      United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//
//		      European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		The complexity inherent in the IntegerSeparatorSpec type is
//		necessary in order to provide support for multinational
//		and multicultural integer separation algorithms. For
//		additional details, reference the source code
//		documentation for type  IntegerSeparatorSpec.
//
//		If this input parameter contains a zero length
//		string, no error will be returned and integer
//		separation will be turned off. As a result,
//		integer digits will be displayed as a single
//		string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	     If 'nStrIntSeparator' is invalid, this method will return
//	     an error.
//
//	pureNumRunes               []rune
//
//		An array of runes consisting entirely of integer digit
//		characters from '0' (0x30) to '9' (0x39) inclusive. If any
//		character within this array is NOT an integer digit, this
//		method will return an error.
//
//
//	errPrefDto                 *ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the names of the calling method or methods.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	numStrWithIntSeps          []rune
//
//		If this method completes successfully, an array of runes
//		will be returned containing the integer digits supplied by
//		input parameter 'pureNumRunes' properly formatted with
//		integer digit separators (a.k.a. thousands separators).
//		Example:
//		  pureNumRunes = 123456789012345
//		  integer separator character = ','
//		  integer grouping for thousands = 3
//		  numStrWithIntSeps = 123,456,789,012,345
//
//		Be advised - This method will never return a number sign
//		for the formatted integer separated numeric value. The
//		numeric sign value (positive or negative) of this numeric
//		value must be tracked externally.
//
//
//	err                        error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errPrefDto'. The
//		'errPrefDto' text will be attached to the beginning of the
//		error message.
func (nStrIntSepMolecule *integerSeparatorSpecMolecule) applyIntSeparators(
	nStrIntSeparator *IntegerSeparatorSpec,
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
		"integerSeparatorSpecMolecule."+
			"applyIntSeparators()",
		"")

	if err != nil {
		return numStrWithIntSeps, err
	}

	if nStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrIntSeparator' (*IntegerSeparatorSpec) is invalid!\n"+
			"'nStrIntSeparator' is a 'nil' pointer.\n",
			ePrefix.String())

		return numStrWithIntSeps, err
	}

	lenPureNumRunes := len(pureNumRunes)

	if len(nStrIntSeparator.intSeparatorChars) == 0 ||
		nStrIntSeparator.turnOffIntegerSeparation == true {
		// This is a NOP condition. Integer separation
		// IS NOT APPLIED
		numStrWithIntSeps = make(
			[]rune, lenPureNumRunes)

		for i := 0; i < lenPureNumRunes; i++ {
			numStrWithIntSeps[i] =
				pureNumRunes[i]

		}

		return numStrWithIntSeps, err
	}

	_,
		err = new(strMechPreon).
		testValidityOfRuneIntArray(
			pureNumRunes,
			ePrefix.XCpy(
				"pureNumRunes"))

	if err != nil {
		return numStrWithIntSeps, err
	}

	// If lenPureNumRunes is zero, an error would
	// have been thrown by
	// testValidityOfRuneIntArray()

	_,
		err = new(integerSeparatorSpecQuark).
		testValidityOfNumStrIntSeparator(
			nStrIntSeparator,
			ePrefix.XCpy("nStrIntSeparator->"))

	if err != nil {
		return numStrWithIntSeps, err
	}

	// If this is zero, an error would have been thrown by
	// testValidityOfNumStrIntSeparator()
	lenIGrpSeq := len(nStrIntSeparator.intSeparatorGrouping)

	var minimumGroupLength uint = math.MaxUint32

	for i := 0; i < lenIGrpSeq; i++ {

		if nStrIntSeparator.intSeparatorGrouping[i] < minimumGroupLength {
			minimumGroupLength = nStrIntSeparator.intSeparatorGrouping[i]
		}

	}

	// If this is zero, an error would have been thrown by
	// testValidityOfNumStrIntSeparator()
	lenIntSeparatorChars := len(nStrIntSeparator.intSeparatorChars)

	var maximumIntSepCharLen = 0

	for i := 0; i < lenIntSeparatorChars; i++ {
		lenIntSep := len(nStrIntSeparator.intSeparatorChars)

		if lenIntSep > maximumIntSepCharLen {
			maximumIntSepCharLen = lenIntSep
		}
	}

	// 'Output Buffer Length ( lenOutRunes ) =
	//  (((Length of pure integer digits/Minimum Grouping Length) + 1) X
	//     Length of Maximum Integer Separation Chars) +
	//         Length of pure integer digits
	lenOutRunes := ((lenPureNumRunes/int(minimumGroupLength) + 1) * maximumIntSepCharLen) + lenPureNumRunes

	outRunes := make([]rune, lenOutRunes)

	outIdx := lenOutRunes - 1

	/*
		fmt.Printf("Length Pure Num Runes: %v\n",
			lenPureNumRunes)

		fmt.Printf("Beginning lenOutRunes: %v\n",
			lenOutRunes)

		fmt.Printf("Beginning outIdx: %v\n",
			outIdx)
	*/

	groupCnt := uint(0)
	maxGroupCnt := nStrIntSeparator.intSeparatorGrouping[0]
	currGroupCntIdx := 0
	lastGroupCntIdx := lenIGrpSeq - 1

	for i := lenPureNumRunes - 1; i >= 0; i-- {

		groupCnt++
		outRunes[outIdx] = pureNumRunes[i]
		outIdx--

		if groupCnt == maxGroupCnt && i != 0 {

			groupCnt = 0

			copy(outRunes[outIdx:], nStrIntSeparator.intSeparatorChars)
			outIdx = outIdx - lenIntSeparatorChars

			// Is this the last integer grouping index?
			if currGroupCntIdx+1 > lastGroupCntIdx {

				if nStrIntSeparator.restartIntGroupingSequence == true {
					// Restart at the first Group Sequence Index
					currGroupCntIdx = 0

					maxGroupCnt =
						nStrIntSeparator.intSeparatorGrouping[currGroupCntIdx]

				} else {
					// Repeat the Last Integer Grouping Index
					maxGroupCnt =
						nStrIntSeparator.intSeparatorGrouping[currGroupCntIdx]
				}

			} else {
				// Go to the next Integer Grouping Index
				currGroupCntIdx++

				maxGroupCnt = nStrIntSeparator.intSeparatorGrouping[currGroupCntIdx]

			}

		} // End of if groupCnt == maxGroupCnt

	} // End of for i := lenPureNumRunes - 1; i >= 0; i--

	// Adjust to last filled index
	outIdx++

	// Actual Output Length =
	//   Original Length Of OutRunes - Final outIdx

	outputLen := lenOutRunes - outIdx

	numStrWithIntSeps = make([]rune, outputLen)

	// Copy only the characters extracted!
	charsCopied := copy(numStrWithIntSeps, outRunes[outIdx:])

	// fmt.Printf("Final outputLen= '%v'\n",	outputLen)
	// fmt.Printf("Final outIdx= '%v'\n", outIdx)

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

//	copyIntSepSpec
//
//	Copies the data fields from input parameter
//	'incomingNStrIntSeparator' to input parameter
//	'targetNStrIntSeparator'.
//
//	If input parameter 'incomingNStrIntSeparator'
//	is judged to be invalid, this method will
//	return an error.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that all data fields in
//	'targetNStrIntSeparator' will be overwritten.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	destinationNStrIntSeparator		*IntegerSeparatorSpec
//
//		A pointer to a IntegerSeparatorSpec instance. All
//		the member variable data fields in this object will
//		be replaced by data values copied from input
//		parameter 'sourceNStrIntSeparator'.
//
//		'destinationPosNumSignSpec' is the destination for
//		this copy operation.
//
//
//	sourceNStrIntSeparator			*IntegerSeparatorSpec
//
//		A pointer to another IntegerSeparatorSpec instance. All
//		the member variable data values from this object will
//		be copied to corresponding member variables in
//		'destinationNStrIntSeparator'.
//
//		'sourceNStrIntSeparator' is the source for this copy
//		operation.
//
//		If 'sourceNStrIntSeparator' is found to be invalid,
//		an error will be returned.
//
//
//	errPrefDto						*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//		     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrIntSepMolecule *integerSeparatorSpecMolecule) copyIntSepSpec(
	destinationNStrIntSeparator *IntegerSeparatorSpec,
	sourceNStrIntSeparator *IntegerSeparatorSpec,
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
		"integerSeparatorSpecMolecule."+
			"copyIn()",
		"")

	if err != nil {
		return err
	}

	if destinationNStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceNStrIntSeparator == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNStrIntSeparator' is"+
			" a 'nil' pointer!\n",
			ePrefix.String())
		return err
	}

	_,
		err =
		new(integerSeparatorSpecQuark).
			testValidityOfNumStrIntSeparator(
				sourceNStrIntSeparator,
				ePrefix.XCpy(
					"Testing validity of 'sourceNStrIntSeparator'."))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&destinationNStrIntSeparator.intSeparatorChars,
		&sourceNStrIntSeparator.intSeparatorChars,
		true,
		ePrefix.XCpy(
			"sourceNStrIntSeparator.intSeparatorChars->"+
				"destinationNStrIntSeparator.intSeparatorChars"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyUnsignedIntArrays(
		&destinationNStrIntSeparator.intSeparatorGrouping,
		&sourceNStrIntSeparator.intSeparatorGrouping,
		true,
		ePrefix.XCpy(
			"sourceNStrIntSeparator.intSeparatorGrouping->"+
				"destinationNStrIntSeparator.intSeparatorGrouping"))

	if err != nil {
		return err
	}

	destinationNStrIntSeparator.restartIntGroupingSequence =
		sourceNStrIntSeparator.restartIntGroupingSequence

	destinationNStrIntSeparator.turnOffIntegerSeparation =
		sourceNStrIntSeparator.turnOffIntegerSeparation

	return err
}

// equal - Receives two IntegerSeparatorSpec objects and proceeds to
// determine whether all data elements in the first object are
// equal to corresponding data elements in the second object.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrIntSepOne       *IntegerSeparatorSpec
//	   - A pointer to the first IntegerSeparatorSpec object. This
//	     method will compare all data elements in this object to
//	     corresponding data elements in the second
//	     IntegerSeparatorSpec object in order determine equivalency.
//
//
//	nStrIntSepTwo       *IntegerSeparatorSpec
//	   - A pointer to the second IntegerSeparatorSpec object. This
//	     method will compare all data elements in the first
//	     IntegerSeparatorSpec object to corresponding data elements in
//	     this second IntegerSeparatorSpec object in order determine
//	     equivalency.
//
//
//	errPrefDto          *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	isEqual             bool
//	   - If all the data elements in 'nStrIntSepOne' are equal to
//	     all the corresponding data elements in 'nStrIntSepTwo',
//	     this return parameter will be set to 'true'. If all the
//	     data elements are NOT equal, this return parameter will be
//	     set to 'false'.
//
//
//	err                 error
//	   - If all the data elements in 'nStrIntSepOne' are equal to
//	     all the corresponding data elements in 'nStrIntSepTwo',
//	     this return parameter will be set to 'nil'.
//
//	     If the corresponding data elements are not equal, a
//	     detailed error message identifying the unequal elements
//	     will be returned.
func (nStrIntSepMolecule *integerSeparatorSpecMolecule) equal(
	nStrIntSepOne *IntegerSeparatorSpec,
	nStrIntSepTwo *IntegerSeparatorSpec,
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
		"integerSeparatorSpecMolecule."+
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

	if nStrIntSepOne.intSeparatorGrouping != nil &&
		nStrIntSepTwo.intSeparatorGrouping == nil {

		err = fmt.Errorf("%v\n"+
			"Error: nStrIntSepOne.intSeparatorGrouping != nil\n"+
			"nStrIntSepTwo.intSeparatorGrouping == nil\n"+
			"nStrIntSepOne.intSeparatorGrouping='%v'\n"+
			"nStrIntSepTwo.intSeparatorGrouping='nil'\n",
			ePrefix.String(),
			nStrIntSepOne.intSeparatorGrouping)

		return isEqual, err
	}

	sMechPreon := strMechPreon{}

	isEqual = sMechPreon.equalRuneArrays(
		nStrIntSepOne.intSeparatorChars,
		nStrIntSepTwo.intSeparatorChars)

	if !isEqual {

		err = fmt.Errorf("%v\n" +
			"Error: 'intSeparatorChars' member variables not equal.\n" +
			ePrefix.String())

		return isEqual, err
	}

	isEqual = sMechPreon.equalUintArrays(
		nStrIntSepOne.intSeparatorGrouping,
		nStrIntSepTwo.intSeparatorGrouping)

	if !isEqual {

		err = fmt.Errorf("%v\n" +
			"Error: 'intSeparatorGrouping' member variables are not equal.\n" +
			ePrefix.String())

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

	if nStrIntSepOne.turnOffIntegerSeparation !=
		nStrIntSepTwo.turnOffIntegerSeparation {

		err = fmt.Errorf("%v\n" +
			"Error: 'turnOffIntegerSeparation' member variables are not equal.\n" +
			ePrefix.String())

		return isEqual, err

	}

	isEqual = true

	return isEqual, err
}

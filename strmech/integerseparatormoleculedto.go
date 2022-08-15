package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"sync"
)

type integerSeparatorDtoMolecule struct {
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
//	  United States Example:  1,000,000,000
//
//	In many European countries, a single period ('.') is used as
//	the integer separator character.
//	  European Example: 1.000.000.000
//
//	Other countries and cultures use spaces, apostrophes or
//	multiple characters to separate integers.
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
// The complexity inherent in the IntegerSeparatorDto type, used as
// an input parameter, is necessary in order to provide support for
// multinational and multicultural integer separation algorithms.
// For more information on integer grouping sequence, reference the
// source documentation for type, IntegerSeparatorDto.
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
//	nStrIntSeparator           *IntegerSeparatorDto
//	   - A pointer to an IntegerSeparatorDto object which contains
//	     the integer separation format parameters which will be
//	     used to insert integer separators.
//
//	     Integer separators consist of a character, or series of
//	     characters, used to separate integer digits in a number
//	     string. These characters are commonly known as the
//	     'thousands separator'. A 'thousands separator' is used to
//	     separate groups of integer digits to the left of the
//	     decimal separator (a.k.a. decimal point). In the United
//	     States, the standard integer digits separator is the
//	     single comma character (',').
//	           United States Example:  1,000,000,000
//
//	     In many European countries, a single period ('.') is used
//	     as the integer separator character.
//	           European Example: 1.000.000.000
//
//	     Other countries and cultures use spaces, apostrophes or
//	     multiple characters to separate integers.
//
//	     The complexity inherent in the IntegerSeparatorDto type is
//	     necessary in order to provide support for multinational
//	     and multicultural integer separation algorithms. For
//	     additional details, reference the source code
//	     documentation for type  IntegerSeparatorDto.
//
//	     If 'nStrIntSeparator' is invalid, this method will return
//	     an error.
//
//
//	pureNumRunes               []rune
//	   - An array of runes consisting entirely of integer digit
//	     characters from '0' (0x30) to '9' (0x39) inclusive. If any
//	     character within this array is NOT an integer digit, this
//	     method will return an error.
//
//
//	errPrefDto                 *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// Return Values
//
//	numStrWithIntSeps          []rune
//	   - If this method completes successfully, an array of runes
//	     will be returned containing the integer digits supplied by
//	     input parameter 'pureNumRunes' properly formatted with
//	     integer digit separators (a.k.a. thousands separators).
//	     Example:
//	       pureNumRunes = 123456789012345
//	       integer separator character = ','
//	       integer grouping for thousands = 3
//	       numStrWithIntSeps = 123,456,789,012,345
//
//	     Be advised - This method will never return a number sign
//	     for the formatted integer separated numeric value. The
//	     numeric sign value (positive or negative) of this numeric
//	     value must be tracked externally.
//
//
//	err                        error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'.
//
//	     If errors are encountered during processing, the returned
//	     error Type will encapsulate an error message. This
//	     returned error message will incorporate the method chain
//	     and text passed by input parameter, 'errPrefDto'. The
//	     'errPrefDto' text will be attached to the beginning of the
//	     error message.
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

	_,
		err = new(strMechPreon).
		testValidityOfRuneIntArray(
			pureNumRunes,
			ePrefix.XCpy(
				"pureNumRunes"))

	if err != nil {
		return numStrWithIntSeps, err
	}

	// If this is zero, an error would have been
	// thrown by testValidityOfRuneIntArray()
	lenPureNumRunes := len(pureNumRunes)

	_,
		err = new(integerSeparatorDtoQuark).
		testValidityOfNumStrIntSeparator(
			nStrIntSeparator,
			ePrefix.XCpy("nStrIntSeparator->"))

	if err != nil {
		return numStrWithIntSeps, err
	}

	// If this is zero, an error would have been thrown by
	// testValidityOfNumStrIntSeparator()
	lenIGrpSeq := len(nStrIntSeparator.intGroupingSequence)

	var minimumGroupLength uint = math.MaxUint32

	for i := 0; i < lenIGrpSeq; i++ {

		if nStrIntSeparator.intGroupingSequence[i] < minimumGroupLength {
			minimumGroupLength = nStrIntSeparator.intGroupingSequence[i]
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
	maxGroupCnt := nStrIntSeparator.intGroupingSequence[0]
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
						nStrIntSeparator.intGroupingSequence[currGroupCntIdx]

				} else {
					// Repeat the Last Integer Grouping Index
					maxGroupCnt =
						nStrIntSeparator.intGroupingSequence[currGroupCntIdx]
				}

			} else {
				// Go to the next Integer Grouping Index
				currGroupCntIdx++

				maxGroupCnt = nStrIntSeparator.intGroupingSequence[currGroupCntIdx]

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

// copyIntSepDto - Copies the data fields from input parameter
// 'incomingNStrIntSeparator' to input parameter
// 'targetNStrIntSeparator'.
//
// Be advised - All data fields in 'targetNStrIntSeparator'
// will be overwritten.
//
// If input parameter 'incomingNStrIntSeparator' is judged
// to be invalid, this method will return an error.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		destinationNStrIntSeparator     *IntegerSeparatorDto
//		   - A pointer to a IntegerSeparatorDto instance. All the
//		     member variable data fields in this object will be
//		     replaced by data values copied from input parameter
//		     'sourceNStrIntSeparator'.
//
//		     'destinationPosNumSignSpec' is the destination for this
//		     copy operation.
//
//
//		sourceNStrIntSeparator          *IntegerSeparatorDto
//		   - A pointer to another IntegerSeparatorDto instance. All
//		     the member variable data values from this object will
//		     be copied to corresponding member variables in
//		     'destinationNStrIntSeparator'.
//
//		     'sourceNStrIntSeparator' is the source for this copy
//		     operation.
//
//	      If 'sourceNStrIntSeparator' is found to be invalid,
//	      an error will be returned.
//
//
//		errPrefDto          *ErrPrefixDto
//		   - This object encapsulates an error prefix string which is
//		     included in all returned error messages. Usually, it
//		     contains the names of the calling method or methods.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     Type ErrPrefixDto is included in the 'errpref' software
//		     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                 error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrIntSepMolecule *integerSeparatorDtoMolecule) copyIntSepDto(
	destinationNStrIntSeparator *IntegerSeparatorDto,
	sourceNStrIntSeparator *IntegerSeparatorDto,
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
		new(integerSeparatorDtoQuark).
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
		&destinationNStrIntSeparator.intGroupingSequence,
		&sourceNStrIntSeparator.intGroupingSequence,
		true,
		ePrefix.XCpy(
			"sourceNStrIntSeparator.intGroupingSequence->"+
				"destinationNStrIntSeparator.intGroupingSequence"))

	if err != nil {
		return err
	}

	destinationNStrIntSeparator.restartIntGroupingSequence =
		sourceNStrIntSeparator.restartIntGroupingSequence

	return err
}

// equal - Receives two IntegerSeparatorDto objects and proceeds to
// determine whether all data elements in the first object are
// equal to corresponding data elements in the second object.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrIntSepOne       *IntegerSeparatorDto
//	   - A pointer to the first IntegerSeparatorDto object. This
//	     method will compare all data elements in this object to
//	     corresponding data elements in the second
//	     IntegerSeparatorDto object in order determine equivalency.
//
//
//	nStrIntSepTwo       *IntegerSeparatorDto
//	   - A pointer to the second IntegerSeparatorDto object. This
//	     method will compare all data elements in the first
//	     IntegerSeparatorDto object to corresponding data elements in
//	     this second IntegerSeparatorDto object in order determine
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

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelQuark - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelQuark struct {
	lock *sync.Mutex
}

//	compareNumStrKernelValues
//
//	This method receives pointers to two instances of
//	NumberStrKernel, 'numStrKernel01' and
//	'numStrKernel02'.
//
//	The numeric value of 'numStrKernel01' is compared to
//	that of 'numStrKernel01'. The comparison results are
//	returned as one of three integer values:
//
//		-1	= numStrKernel01 is less than numStrKernel02
//		 0	= numStrKernel01 is equal to numStrKernel02
//		+1	= numStrKernel01 is greater than numStrKernel02
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//  1. This method assumes the integer and fractional
//     digit arrays contained in input parameters
//     'numStrKernel01' and 'numStrKernel02' are equal
//     in length.
//
//     If the integer digit array length and fractional
//     digit array lengths of 'numStrKernel01' are NOT
//     equal to the corresponding array lengths in
//     'numStrKernel02', an error will be returned.
//
//  2. This method assumes that the number signs for
//     'numStrKernel01' and 'numStrKernel02' are equal.
//     If 'numStrKernel01' and 'numStrKernel02' ARE NOT
//     equal, an error will be returned.
//
//     Possible values for number sign are listed as
//     follows:
//
//     NumSignVal.Negative() = -1
//     NumSignVal.Zero()     =  0
//     NumSignVal.Positive() =  1
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel01				*NumberStrKernel
//
//		The numeric value of numStrKernel01 will be
//		compared to that of numStrKernel02. The
//		comparison results will be returned as an integer
//		value.
//
//	numStrKernel02				*NumberStrKernel
//
//		The numeric value of numStrKernel01 will be
//		compared to that of this parameter,
//		numStrKernel02. The comparison results will be
//		returned as an integer value.
//
//	errPrefDto					*ePref.ErrPrefixDto
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
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	comparisonValue				int
//
//		This parameter will return the results of numeric
//		value comparisons for input parameters,
//		'numStrKernel01' and 'numStrKernel02'. The
//		integer comparison result will be set to one of
//		three values:
//
//		-1	= numStrKernel01 is less than numStrKernel02
//		 0	= numStrKernel01 is equal to numStrKernel02
//		+1	= numStrKernel01 is greater than numStrKernel02
//
//	err							error
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
func (numStrKernelQuark *numberStrKernelQuark) compareNumStrKernelValues(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	comparisonValue int,
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"equalizeNStrFracDigitsLengths()",
		"")

	if err != nil {

		return comparisonValue, err
	}

	if numStrKernel01 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel01' is a nil pointer!\n",
			ePrefix.String())

		return comparisonValue, err
	}

	if numStrKernel02 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel02' is a nil pointer!\n",
			ePrefix.String())

		return comparisonValue, err
	}

	if numStrKernel01.numberSign !=
		numStrKernel02.numberSign {

		err = fmt.Errorf("%v\n"+
			"ERROR: The number signs for 'numStrKernel01'"+
			"and 'numStrKernel02' ARE NOT EQUAL!\n"+
			"numStrKernel01.numberSign = '%v'\n"+
			"numStrKernel02.numberSign = '%v'\n",
			ePrefix.String(),
			numStrKernel01.numberSign.String(),
			numStrKernel02.numberSign.String())

		return comparisonValue, err
	}

	if numStrKernel01.numberSign == NumSignVal.Zero() {

		comparisonValue = 0

		return comparisonValue, err
	}

	lenIntDigits01 :=
		len(numStrKernel01.integerDigits.CharsArray)

	lenFracDigits01 :=
		len(numStrKernel01.fractionalDigits.CharsArray)

	lenIntDigits02 :=
		len(numStrKernel02.integerDigits.CharsArray)

	lenFracDigits02 :=
		len(numStrKernel02.fractionalDigits.CharsArray)

	if lenIntDigits01 != lenIntDigits02 {

		err = fmt.Errorf("%v\n"+
			"Error: numStrKernel01.integerDigits array length is\n"+
			"NOT EQUAL to numStrKernel02.integerDigits array length!\n,"+
			"numStrKernel01.integerDigits array length = '%v'\n"+
			"numStrKernel02.integerDigits array length = '%v'\n",
			ePrefix.String(),
			lenIntDigits01,
			lenIntDigits02)

		return comparisonValue, err
	}

	if lenFracDigits01 != lenFracDigits02 {

		err = fmt.Errorf("%v\n"+
			"Error: numStrKernel01.fractionalDigits array length is\n"+
			"NOT EQUAL to numStrKernel02.fractionalDigits array length!\n,"+
			"numStrKernel01.fractionalDigits array length = '%v'\n"+
			"numStrKernel02.fractionalDigits array length = '%v'\n",
			ePrefix.String(),
			lenFracDigits01,
			lenFracDigits02)

		return comparisonValue, err
	}

	for i := 0; i < lenIntDigits01; i++ {

		if numStrKernel01.integerDigits.CharsArray[i] >
			numStrKernel02.integerDigits.CharsArray[i] {

			// Remember, numStrKernel01 & numStrKernel02
			// number signs are equal
			if numStrKernel01.numberSign == NumSignVal.Positive() {

				comparisonValue = 1

			} else {
				//	MUST BE -
				//	numStrKernel01.numberSign ==
				//		NumSignVal.Negative()
				comparisonValue = -1

			}

			return comparisonValue, err

		}

		if numStrKernel02.integerDigits.CharsArray[i] >
			numStrKernel01.integerDigits.CharsArray[i] {

			// Remember, numStrKernel01 & numStrKernel02
			// number signs are equal
			if numStrKernel02.numberSign == NumSignVal.Positive() {

				comparisonValue = -1

			} else {
				//	MUST BE -
				//	numStrKernel02.numberSign ==
				//		NumSignVal.Negative()
				comparisonValue = 1

			}

			return comparisonValue, err
		}
	}

	// Integer Digits ARE EQUAL
	// Now test Fractional Digits

	for k := 0; k < lenFracDigits01; k++ {

		if numStrKernel01.fractionalDigits.CharsArray[k] >
			numStrKernel02.fractionalDigits.CharsArray[k] {

			// Remember, numStrKernel01 & numStrKernel02
			// number signs are equal
			if numStrKernel01.numberSign == NumSignVal.Positive() {

				comparisonValue = 1

			} else {
				//	MUST BE -
				//	numStrKernel01.numberSign ==
				//		NumSignVal.Negative()
				comparisonValue = -1

			}

			return comparisonValue, err
		}

		if numStrKernel02.fractionalDigits.CharsArray[k] >
			numStrKernel01.fractionalDigits.CharsArray[k] {

			// Remember, numStrKernel01 & numStrKernel02
			// number signs are equal
			if numStrKernel02.numberSign == NumSignVal.Positive() {

				comparisonValue = -1

			} else {
				//	MUST BE -
				//	numStrKernel02.numberSign ==
				//		NumSignVal.Negative()
				comparisonValue = 1

			}

			return comparisonValue, err
		}
	}

	// MUST BE -
	//	numStrKernel01 and numStrKernel02 have
	//	equal numerical values
	comparisonValue = 0

	return comparisonValue, err
}

// extractNumRunes
//
// Receives an array of runes and extracts the numeric
// digits as text characters. Set 'characterSearchLength'
// to minus one (-1) for end-of-string search.
func (numStrKernelQuark *numberStrKernelQuark) extractNumRunes(
	targetSearchString RuneArrayDto,
	targetSearchStringName string,
	startingSearchIndex int,
	breakOnCharSearchLength int,
	negativeNumSearchSpecsCol NegNumSearchSpecCollection,
	decimalSeparatorSpec DecimalSeparatorSpec,
	breakOnCharDelimiters RuneArrayCollection,
	requestRemainderRunesString bool,
	ePrefDto *ePref.ErrPrefixDto) (
	searchResults CharSearchNumStrParseResultsDto,
	numStrKernel NumberStrKernel,
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"numStrBuilderElectron."+
			"extractNumRunes()",
		"")

	if err != nil {

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms := CharSearchTargetInputParametersDto{}.New()

	targetInputParms.TargetString = &targetSearchString

	targetInputParms.TargetStringLength = targetSearchString.GetRuneArrayLength()

	if targetInputParms.TargetStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSearchString' is invalid.\n"+
			"'targetSearchString' has an array length of zero!\n",
			ePrefix.String())

		return searchResults,
			numStrKernel,
			err

	}

	sMechPreon := strMechPreon{}
	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		targetInputParms.TargetString.CharsArray,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid.\n"+
			"'%v' should contain valid characters.\n"+
			"A validity test on this rune array produced the following error:\n"+
			"%v\n",
			ePrefix.String(),
			targetSearchStringName,
			targetSearchStringName,
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.TargetInputParametersName = "Extract Number Runes"
	targetInputParms.TargetStringDescription2 =
		"strMechMolecule.extractNumRunes()"

	if len(targetSearchStringName) == 0 {
		targetInputParms.TargetStringName = "targetSearchString"
	} else {
		targetInputParms.TargetStringName = targetSearchStringName
	}

	targetInputParms.TargetStringLengthName =
		targetInputParms.TargetStringName + "Length"

	targetInputParms.TargetStringStartingSearchIndexName =
		targetInputParms.TargetStringName + "StartingSearchIndex"

	targetInputParms.TargetStringSearchLength = -1

	targetInputParms.FoundFirstNumericDigitInNumStr = false

	targetInputParms.TargetStringStartingSearchIndex =
		startingSearchIndex

	if targetInputParms.TargetStringStartingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid.\n"+
			"'startingSearchIndex' has a value less than zero!\n"+
			"startingSearchIndex = '%v'\n",
			ePrefix.String(),
			targetInputParms.TargetStringStartingSearchIndex)

		return searchResults,
			numStrKernel,
			err
	}

	if targetInputParms.TargetStringStartingSearchIndex >=
		targetInputParms.TargetStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid.\n"+
			"'startingSearchIndex' has a value greater than the last index\n"+
			"of %v!\n"+
			"startingSearchIndex = '%v'\n"+
			"%v last index = %v\n",
			ePrefix.String(),
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringStartingSearchIndex,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringLength-1)

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.TargetStringSearchLength = breakOnCharSearchLength

	err = targetInputParms.ValidateTargetParameters(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {

		return searchResults,
			numStrKernel,
			err

	}

	err2 = negativeNumSearchSpecsCol.IsValidInstanceError(
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeNumSearchSpecsCol' is invalid!\n"+
			"The following validation error was returned:\n"+
			"%v\n",
			ePrefix.XCpy(
				"negativeNumSearchSpecsCol"),
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	err2 = decimalSeparatorSpec.IsValidInstanceError(
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparatorSpec' is invalid!\n"+
			"The following validation error was returned:\n"+
			"%v\n",
			ePrefix.XCpy(
				"decimalSeparatorSpec"),
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.RequestRemainderString = false
	targetInputParms.RequestFoundTestCharacters = false
	targetInputParms.RequestReplacementString = false

	searchResults.Empty()

	searchResults.SearchResultsName = "Number Runes Extraction Results"

	searchResults.SearchResultsFunctionChain =
		ePrefix.String()

	searchResults.FoundNumericDigits = false
	searchResults.FoundNonZeroValue = false
	searchResults.FoundDecimalSeparatorSymbols = false
	searchResults.FoundDecimalDigits = false
	searchResults.NumSignValue = NumSignVal.Zero()
	searchResults.NumValueType = NumValType.None()
	searchResults.RemainderString.Empty()
	searchResults.DecimalSeparatorSearchResults.Empty()
	searchResults.NegativeNumberSymbolSearchResults.Empty()
	searchResults.ParsingTerminatorSearchResults.Empty()

	err = searchResults.TargetSearchString.CopyIn(
		&targetSearchString,
		ePrefix.XCpy(
			"searchResults.TargetSearchString"+
				"<-targetSearchString"))

	if err != nil {

		return searchResults,
			numStrKernel,
			err
	}

	searchResults.TargetStringSearchLength =
		targetInputParms.TargetStringSearchLength

	searchResults.TargetStringAdjustedSearchLength =
		targetInputParms.TargetStringAdjustedSearchLength

	searchResults.TargetStringStartingSearchIndex =
		targetInputParms.TargetStringStartingSearchIndex

	searchResults.TargetStringLastSearchIndex = -1

	searchResults.TargetStringNextSearchIndex = -1

	searchResults.ReasonForSearchTermination =
		CharSearchTermType.ProcessError()

	// Processing Flags

	// Number Parsing Setup
	searchResults.ParsingTerminatorSearchResults.IsNOP =
		breakOnCharDelimiters.IsNOP()

	searchResults.ParsingTerminatorSearchResults.SearchResultsName =
		"Number Parsing Terminator Search Results"

	numParsingTerminatorsIsNOP :=
		searchResults.ParsingTerminatorSearchResults.IsNOP

	// Decimal Separator Setup
	searchResults.DecimalSeparatorSearchResults.IsNOP =
		decimalSeparatorSpec.IsNOP()

	searchResults.DecimalSeparatorSearchResults.SearchResultsName =
		"Decimal Separator Search Results"

	decSeparatorIsNOP :=
		searchResults.DecimalSeparatorSearchResults.IsNOP

	var nextIdx int
	var tempDecSepSearchResults CharSearchDecimalSeparatorResultsDto

	for i := targetInputParms.TargetStringStartingSearchIndex; i < targetInputParms.TargetStringAdjustedSearchLength; i++ {

		targetInputParms.TargetStringCurrentSearchIndex = i

		if targetSearchString.CharsArray[i] >= '0' &&
			targetSearchString.CharsArray[i] <= '9' {

			searchResults.FoundNumericDigits = true
			targetInputParms.FoundFirstNumericDigitInNumStr = true

			if targetSearchString.CharsArray[i] > '0' {

				searchResults.FoundNonZeroValue = true
				targetInputParms.FoundNonZeroValue = true

				if searchResults.NumSignValue == NumSignVal.Zero() {
					searchResults.NumSignValue = NumSignVal.Positive()
				}
			}

			if !searchResults.DecimalSeparatorSearchResults.
				FoundDecimalSeparatorSymbols {

				err = numStrKernel.AddIntegerDigit(
					targetSearchString.CharsArray[i],
					ePrefix.XCpy(
						fmt.Sprintf(
							"targetSearchString.CharsArray[%v]",
							i)))

				if err != nil {
					return searchResults,
						numStrKernel,
						err
				}

				searchResults.FoundIntegerDigits = true

			} else {

				err = numStrKernel.AddFractionalDigit(
					targetSearchString.CharsArray[i],
					ePrefix.XCpy(
						fmt.Sprintf(
							"targetSearchString.CharsArray[%v]",
							i)))

				if err != nil {
					return searchResults,
						numStrKernel,
						err
				}

				searchResults.FoundDecimalDigits = true
			}

			continue
		}

		// Check for Parsing Terminators
		// All Parsing Operations Cease if Delimiter is Found
		if !numParsingTerminatorsIsNOP &&
			targetInputParms.FoundFirstNumericDigitInNumStr {

			searchResults.ParsingTerminatorSearchResults,
				err = breakOnCharDelimiters.SearchForTextCharacters(
				targetInputParms,
				ePrefix.XCpy(
					"breakOnCharDelimiters"))

			if searchResults.ParsingTerminatorSearchResults.FoundSearchTarget {

				i = searchResults.ParsingTerminatorSearchResults.
					TargetStringCurrentSearchIndex

				targetInputParms.TargetStringCurrentSearchIndex = i

				searchResults.ReasonForSearchTermination = CharSearchTermType.TerminationDelimiters()

				searchResults.ParsingTerminatorSearchResults.SearchResultsName =
					"Number Parsing Terminator Search Results"

				goto computeExitStats
			}
		}

		// Check for Negative Number Sign Symbol
		if !searchResults.NegativeNumberSymbolSearchResults.
			FoundNegativeNumberSymbols {

			searchResults.NegativeNumberSymbolSearchResults,
				err = negativeNumSearchSpecsCol.
				SearchForNegNumSignSymbols(
					targetInputParms,
					ePrefix.XCpy(
						"negativeNumSearchSpecsCol"))

			if searchResults.NegativeNumberSymbolSearchResults.
				FoundNegativeNumberSymbols {

				i =
					searchResults.NegativeNumberSymbolSearchResults.
						TargetStringCurrentSearchIndex

				targetInputParms.TargetStringCurrentSearchIndex = i

				targetInputParms.TargetStringNextSearchIndex = i + 1

				if targetInputParms.TargetStringNextSearchIndex >=
					targetInputParms.TargetStringLength {

					targetInputParms.TargetStringNextSearchIndex = -1

				}

				searchResults.NumSignValue = NumSignVal.Negative()

				if searchResults.NegativeNumberSymbolSearchResults.
					PrimaryNumSignPosition == NumSignSymPos.BeforeAndAfter() ||
					searchResults.NegativeNumberSymbolSearchResults.
						PrimaryNumSignPosition == NumSignSymPos.After() {

					goto computeExitStats
				}

				continue
			}
		}

		// Check for Decimal Separators
		if decSeparatorIsNOP == false &&
			targetInputParms.FoundDecimalSeparatorSymbols == false {

			if targetInputParms.FoundFirstNumericDigitInNumStr == false {

				nextIdx = i + 1

				if nextIdx < targetInputParms.TargetStringAdjustedSearchLength &&
					targetSearchString.CharsArray[nextIdx] >= '0' &&
					targetSearchString.CharsArray[nextIdx] <= '9' {

					tempDecSepSearchResults,
						err = decimalSeparatorSpec.SearchForDecimalSeparator(
						targetInputParms,
						ePrefix.XCpy(
							"decimalSeparatorSpec"))

					if err != nil {

						return searchResults,
							numStrKernel,
							err

					}

					if tempDecSepSearchResults.FoundDecimalSeparatorSymbols == true {

						err =
							searchResults.DecimalSeparatorSearchResults.CopyIn(
								&tempDecSepSearchResults,
								ePrefix.XCpy(
									"No Int Digits: "+
										"searchResults<-tempDecSepSearchResults"))

						if err != nil {

							return searchResults,
								numStrKernel,
								err

						}

						searchResults.FoundDecimalSeparatorSymbols = true

						targetInputParms.FoundDecimalSeparatorSymbols = true

						i = searchResults.DecimalSeparatorSearchResults.
							TargetStringLastSearchIndex

						targetInputParms.TargetStringCurrentSearchIndex = i

						targetInputParms.TargetStringNextSearchIndex = i + 1

						if targetInputParms.TargetStringNextSearchIndex >=
							targetInputParms.TargetStringLength {

							targetInputParms.TargetStringNextSearchIndex = -1

						}

						err = numStrKernel.AddIntegerDigit(
							'0',
							ePrefix.XCpy(
								fmt.Sprintf(
									"targetSearchString.CharsArray[%v]",
									i)))

						if err != nil {
							return searchResults,
								numStrKernel,
								err
						}

						searchResults.FoundNumericDigits = true
						targetInputParms.FoundFirstNumericDigitInNumStr = true
						searchResults.FoundIntegerDigits = true

						continue
					}

				}

			} else {
				// MUST BE
				// targetInputParms.FoundFirstNumericDigitInNumStr == true

				tempDecSepSearchResults,
					err = decimalSeparatorSpec.SearchForDecimalSeparator(
					targetInputParms,
					ePrefix.XCpy(
						"decimalSeparatorSpec"))

				if err != nil {

					return searchResults,
						numStrKernel,
						err

				}

				if tempDecSepSearchResults.FoundDecimalSeparatorSymbols == true {

					err =
						searchResults.DecimalSeparatorSearchResults.CopyIn(
							&tempDecSepSearchResults,
							ePrefix.XCpy(
								"Found1stNumDigit: "+
									"searchResults<-tempDecSepSearchResults"))

					if err != nil {

						return searchResults,
							numStrKernel,
							err

					}

					searchResults.FoundDecimalSeparatorSymbols = true

					targetInputParms.FoundDecimalSeparatorSymbols = true

					i = searchResults.DecimalSeparatorSearchResults.
						TargetStringLastSearchIndex

					targetInputParms.TargetStringCurrentSearchIndex = i

					targetInputParms.TargetStringNextSearchIndex = i + 1

					if targetInputParms.TargetStringNextSearchIndex >=
						targetInputParms.TargetStringLength {

						targetInputParms.TargetStringNextSearchIndex = -1

					}

					continue
				}

			}
		}
	}

	// End Of Target String Standard Loop Completion

	if targetInputParms.TargetStringAdjustedSearchLength <
		targetInputParms.TargetStringLength {

		searchResults.ReasonForSearchTermination =
			CharSearchTermType.SearchLengthLimit()

	} else {

		searchResults.ReasonForSearchTermination =
			CharSearchTermType.EndOfTargetString()

	}

computeExitStats:

	targetInputParms.TargetStringNextSearchIndex =
		targetInputParms.TargetStringCurrentSearchIndex + 1

	if targetInputParms.TargetStringNextSearchIndex >=
		targetInputParms.TargetStringLength {

		targetInputParms.TargetStringNextSearchIndex = -1

		searchResults.TargetStringNextSearchIndex = -1

		searchResults.TargetStringLastSearchIndex =
			targetInputParms.TargetStringLength - 1

	} else {
		// targetInputParms.TargetStringNextSearchIndex <
		// targetInputParms.TargetStringLength

		searchResults.TargetStringNextSearchIndex =
			targetInputParms.TargetStringNextSearchIndex

		searchResults.TargetStringLastSearchIndex =
			targetInputParms.TargetStringNextSearchIndex - 1

	}

	// Slice Examples
	//arr := []int{1,2,3,4,5}
	//
	//fmt.Println(arr[:2])        // [1,2]
	//
	//fmt.Println(arr[2:])        // [3,4,5]
	//
	//fmt.Println(arr[2:3])        // [3]
	//
	//fmt.Println(arr[:])            // [1,2,3,4,5]

	searchResults.RemainderString.Empty()

	if searchResults.FoundNumericDigits == true {

		if requestRemainderRunesString {

			if targetInputParms.TargetStringNextSearchIndex == -1 {
				// All characters have been searched.
				// There is no Remainder String
			} else {
				// All characters HAVE NOT BEEN SEARCHED
				// There is a Remainder String

				searchResults.RemainderString.CharsArray =
					append(
						searchResults.RemainderString.CharsArray,
						targetSearchString.
							CharsArray[targetInputParms.TargetStringCurrentSearchIndex:]...)
			}

		}

		if numStrKernel.GetNumberOfIntegerDigits() > 0 {

			searchResults.IdentifiedIntegerDigits =
				numStrKernel.GetIntegerString()
		}

		if numStrKernel.GetNumberOfFractionalDigits() > 0 {

			searchResults.IdentifiedFractionalDigits =
				numStrKernel.GetFractionalString()
		}

		if numStrKernel.GetNumberOfFractionalDigits() > 0 &&
			numStrKernel.GetNumberOfIntegerDigits() == 0 {

			// Rationalize Fractional Digits
			err = numStrKernel.AddIntegerDigit(
				'0',
				ePrefix.XCpy("Add '0' int digit"))

			if err != nil {

				return searchResults,
					numStrKernel,
					err
			}

		}

		if searchResults.FoundDecimalDigits {

			searchResults.NumValueType = NumValType.FloatingPoint()

		} else {

			searchResults.NumValueType = NumValType.Integer()

		}

		if searchResults.FoundNonZeroValue == true {
			// Value is Nonzero

			if searchResults.NumSignValue == NumSignVal.Negative() {
				numStrKernel.numberSign = NumSignVal.Negative()
			} else {
				numStrKernel.numberSign = NumSignVal.Positive()
			}

		} else {
			// Value is Zero
			numStrKernel.numberSign = NumSignVal.Zero()
		}

	} else {
		// searchResults.FoundSearchTarget == false
		// Didn't find any numeric digits!
		numStrKernel.Empty()
		searchResults.FoundNumericDigits = false
		searchResults.FoundDecimalDigits = false
		searchResults.FoundNonZeroValue = false

		searchResults.RemainderString.CharsArray =
			targetSearchString.CharsArray[:]

	}

	return searchResults,
		numStrKernel,
		err
}

//	getSetNumValueType
//
//	Sets and returns the current NumericValueType for the
//	instance of NumberStrKernel passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		method will examine the internal member variables
//		contained in this instance and set the correct
//		value for Numeric Value Type.
//
//		NumericValueType is an enumeration value specifying
//		the type of numeric value contained in the
//		'numStrKernel' instance.
//
//		Possible NumericValueType enumeration values are
//		listed as follows:
//			NumValType.None()
//			NumValType.FloatingPoint()
//			NumValType.Integer()
//
//		The internal variable contained in 'numStrKernel'
//		which will be configured is:
//
//			NumberStrKernel.numberValueType
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelQuark *numberStrKernelQuark) getSetNumValueType(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	NumericValueType,
	error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumericValueType := NumValType.None()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"getSetNumValueType()",
		"")

	if err != nil {

		return newNumericValueType, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return newNumericValueType, err
	}

	lenIntegerDigits :=
		numStrKernel.integerDigits.GetRuneArrayLength()

	lenFracDigits :=
		numStrKernel.fractionalDigits.GetRuneArrayLength()

	if lenIntegerDigits == 0 &&
		lenFracDigits == 0 {

		newNumericValueType = NumValType.None()

	} else if lenIntegerDigits > 0 &&
		lenFracDigits == 0 {

		newNumericValueType = NumValType.Integer()

	} else {

		// MUST BE lenFracDigits > 0

		newNumericValueType = NumValType.FloatingPoint()

	}

	numStrKernel.numberValueType = newNumericValueType

	return newNumericValueType, err
}

//	equalizeNStrIntDigitsLengths
//
//	Receives pointers to two instances of
//	NumberStrKernel, 'numStrKernel01' and
//	'numStrKernel02'. This method will ensure that the
//	integer arrays contained in both instances have
//	equal array lengths.
//
//	If the integer arrays do not have equal array
//	lengths, leading zero characters ('0') will be added
//	to configure their array lengths as equal.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel01				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal integer array contained in this instance
//		will be	compared to that of input parameter,
//		'numStrKernel02'. If the 'numStrKernel01' integer
//		array length is shorter than that of
//		'numStrKernel02', leading zero characters ('0')
//		will be added to achieve an equal integer array
//		length with the integer array contained in
//		'numStrKernel02'.
//
//	numStrKernel02				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal integer array contained in this instance
//		will be	compared to that of input parameter,
//		'numStrKernel01'. If the 'numStrKernel02' integer
//		array length is shorter than that of
//		'numStrKernel01', leading zero characters ('0')
//		will be added to achieve an equal integer array
//		length with the integer array contained in
//		'numStrKernel01'.
//
//	errPrefDto					*ePref.ErrPrefixDto
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
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
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
func (numStrKernelQuark *numberStrKernelQuark) equalizeNStrIntDigitsLengths(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"equalizeNStrIntDigitsLengths()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel01 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel01' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if numStrKernel02 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel02' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenIntArray01 :=
		len(numStrKernel01.integerDigits.CharsArray)

	lenIntArray02 :=
		len(numStrKernel02.integerDigits.CharsArray)

	if lenIntArray01 == lenIntArray02 {

		// Nothing to do
		return err
	}

	// Integer Array Lengths Are NOT Equal

	var numOfCharsToAdd int

	if lenIntArray01 > lenIntArray02 {

		numOfCharsToAdd =
			lenIntArray01 - lenIntArray02

		err =
			numStrKernel02.integerDigits.ExtendRuneArray(
				'0',
				numOfCharsToAdd,
				false,
				ePrefix.XCpy(
					"numStrKernel02.integerDigits"))

		if err != nil {

			return err

		}

	} else {
		// MUST BE
		// lenIntArray02 > lenIntArray01

		numOfCharsToAdd =
			lenIntArray02 - lenIntArray01

		err =
			numStrKernel01.integerDigits.ExtendRuneArray(
				'0',
				numOfCharsToAdd,
				false,
				ePrefix.XCpy(
					"numStrKernel01.integerDigits"))

		if err != nil {

			return err

		}

	}

	return err
}

//	equalizeNStrFracDigitsLengths
//
//	Receives pointers to two instances of
//	NumberStrKernel, 'numStrKernel01' and
//	'numStrKernel01'. This method will ensure that the
//	fractional arrays contained in both instances have
//	equal array lengths.
//
//	If the fractional arrays do not have equal array
//	lengths, trailing zero characters ('0') will be added
//	to configure their array lengths as equal.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel01				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal fractional array contained in this instance
//		will be	compared to that of input parameter,
//		'numStrKernel02'. If the 'numStrKernel01' fractional
//		array length is shorter than that of
//		'numStrKernel02', trailing zero characters ('0')
//		will be added to achieve an equal fractional array
//		length with the fractional array contained in
//		'numStrKernel02'.
//
//	numStrKernel02				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal fractional array contained in this instance
//		will be	compared to that of input parameter,
//		'numStrKernel01'. If the 'numStrKernel02' fractional
//		array length is shorter than that of
//		'numStrKernel01', trailing zero characters ('0')
//		will be added to achieve an equal fractional array
//		length with the fractional array contained in
//		'numStrKernel01'.
//
//	errPrefDto					*ePref.ErrPrefixDto
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
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
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
func (numStrKernelQuark *numberStrKernelQuark) equalizeNStrFracDigitsLengths(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"equalizeNStrFracDigitsLengths()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel01 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel01' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if numStrKernel02 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel02' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenFracArray01 :=
		len(numStrKernel01.fractionalDigits.CharsArray)

	lenFracArray02 :=
		len(numStrKernel02.fractionalDigits.CharsArray)

	if lenFracArray01 == lenFracArray02 {

		// Nothing to do
		return err
	}

	// Fractional Digit Array Lengths Are NOT Equal

	var numOfCharsToAdd int

	if lenFracArray01 > lenFracArray02 {

		numOfCharsToAdd =
			lenFracArray01 - lenFracArray02

		err =
			numStrKernel02.fractionalDigits.ExtendRuneArray(
				'0',
				numOfCharsToAdd,
				true,
				ePrefix.XCpy(
					"numStrKernel02.fractionalDigits"))

		if err != nil {

			return err

		}

	} else {
		// MUST BE
		// lenFracArray02 > lenFracArray01

		numOfCharsToAdd =
			lenFracArray02 - lenFracArray01

		err =
			numStrKernel01.fractionalDigits.ExtendRuneArray(
				'0',
				numOfCharsToAdd,
				true,
				ePrefix.XCpy(
					"numStrKernel01.fractionalDigits"))

		if err != nil {

			return err

		}

	}

	return err
}

//	parsePureNumStr
//
//	Receives a pure number string and proceeds to return the
//	extracted numeric value as a type NumberStrKernel.
//
//	This method is particularly useful when numeric values
//	are converted to string using 'fmt.Sprintf()' and
//	similar formatting algorithms.
//
//	A "Pure Number String" is defined as follows:
//
//		1.	Consists of numeric character digits
//			zero through nine inclusive (0-9).
//
//		2.	Option: A Pure Number String may include
//			a radix point or decimal separator. The
//			decimal separator may consist of one or
//			more characters.
//
//			In the US, UK, Australia and most of Canada,
//			the decimal separator is the period
//			character ('.') known as the decimal point.
//
//		3.	Optional: A Pure Number String may include a
//			negative number sign symbol consisting of a
//			minus sign ('-'). Only the minus sign ('-')
//			classifies the numeric value as a negative
//			number in Pure Number String.
//
//			If the leading or trailing minus sign ('-')
//			is NOT present, the numeric value is assumed
//			to be positive.
//
//		4.	Only numeric characters, the decimal
//			separator and the minus sign will be
//			processed by the number string parsing
//			algorithm. All other characters will be
//			ignored.
//
//		5.	Pure Number Strings consist of a single
//			numeric value. The entire Pure Number String
//			will be parsed, or processed, and only one
//			numeric value per Pure Number String will
//			be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberString			RuneArrayDto
//
//		This rune array contains the character digits from
//		which the numeric value will be extracted and
//		returned as a NumberStrKernel.
//
//		A "Pure Number String" is defined as follows:
//
//			1.	Consists of numeric character digits
//				zero through nine inclusive (0-9).
//
//			2.	Option: A Pure Number String may include
//				a radix point or decimal separator. The
//				decimal separator may consist of one or
//				more characters.
//
//				In the US, UK, Australia and most of Canada,
//				the decimal separator is the period
//				character ('.') known as the decimal point.
//
//			3.	Optional: A Pure Number String may include a
//				negative number sign symbol consisting of a
//				minus sign ('-'). Only the minus sign ('-')
//				classifies the numeric value as a negative
//				number in Pure Number String.
//
//				If the leading or trailing minus sign ('-')
//				is NOT present, the numeric value is assumed
//				to be positive.
//
//			4.	Only numeric characters, the decimal
//				separator and the minus sign will be
//				processed by the number string parsing
//				algorithm. All other characters will be
//				ignored.
//
//			5.	Pure Number Strings consist of a single
//				numeric value. The entire Pure Number String
//				will be parsed, or processed, and only one
//				numeric value per Pure Number String will
//				be returned.
//
//		If parameter 'pureNumberString' fails to include
//		numeric character digits, an error will be returned.
//
//	decSeparatorSpec				DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingNumSymbols 			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format. Number Symbols refers to
//		the negative number sign or minus sign ('-') used
//		in classifying negative numeric values.
//
//		When set to 'true', the Pure Number String
//		parsing algorithm will search for leading minus
//		signs ('-') at the beginning of the Pure Number
//		String. Leading minus signs represent the
//		standard for designating negative numeric values
//		in the US, UK, Australia and most of Canada.
//
//		Example Leading Number Symbols:
//			"-123.456"
//
//		When set to 'false', the Pure Number String
//		parsing algorithm will search for trailing minus
//		signs ('-') at the end of Pure Number String.
//		Trailing minus signs represent the standard for
//		France, Germany and many countries in the
//		European Union.
//
//		Example Trailing Number Symbols:
//			"123.456-"
//
//	errPrefDto          		*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numStrKernel				NumberStrKernel
//
//		If this method completes successfully, an instance of
//		NumberStrKernel containing the numeric value extracted
//		from parameter 'pureNumberString' will be returned.
//
//	err							error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error Type
//		will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelQuark *numberStrKernelQuark) parsePureNumStr(
	pureNumberString RuneArrayDto,
	decSeparatorSpec DecimalSeparatorSpec,
	leadingNumSymbols bool,
	ePrefDto *ePref.ErrPrefixDto) (
	numStrKernel NumberStrKernel,
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"numberStrKernelQuark."+
			"parsePureNumStr()",
		"")

	if err != nil {

		return numStrKernel, err
	}

	lenPureNStr := pureNumberString.GetRuneArrayLength()

	if lenPureNStr == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'pureNumberString' is empty\n"+
			"and contains zero characters.\n",
			ePrefix.String())

		return numStrKernel, err
	}

	err = numStrKernel.integerDigits.SetCharacterSearchType(
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy("integerDigits"))

	if err != nil {

		return numStrKernel, err
	}

	err = numStrKernel.fractionalDigits.SetCharacterSearchType(
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy("fractionalDigits"))

	if err != nil {

		return numStrKernel, err
	}

	targetInputParms := CharSearchTargetInputParametersDto{}.New()

	targetInputParms.TargetString = &pureNumberString

	targetInputParms.TargetStringLength = lenPureNStr

	if targetInputParms.TargetStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'pureNumberString' is invalid.\n"+
			"'pureNumberString' has an array length of zero!\n",
			ePrefix.String())

		return numStrKernel, err

	}

	targetInputParms.TargetStringName = "pureNumberString"

	sMechPreon := strMechPreon{}
	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		targetInputParms.TargetString.CharsArray,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid.\n"+
			"'%v' should contain valid characters.\n"+
			"A validity test on this rune array produced the following error:\n"+
			"%v\n",
			ePrefix.String(),
			pureNumberString,
			pureNumberString,
			err2.Error())

		return numStrKernel, err
	}

	targetInputParms.TargetInputParametersName = "Extract Number Runes"
	targetInputParms.TargetStringDescription2 =
		"strMechMolecule.extractNumRunes()"

	targetInputParms.TargetStringLengthName =
		targetInputParms.TargetStringName + "Length"

	targetInputParms.TargetStringStartingSearchIndexName =
		targetInputParms.TargetStringName + "StartingSearchIndex"

	targetInputParms.TargetStringSearchLength = lenPureNStr

	targetInputParms.FoundFirstNumericDigitInNumStr = false

	targetInputParms.TargetStringStartingSearchIndex = 0

	if targetInputParms.TargetStringStartingSearchIndex >=
		targetInputParms.TargetStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid.\n"+
			"'startingSearchIndex' has a value greater than the last index\n"+
			"of %v!\n"+
			"startingSearchIndex = '%v'\n"+
			"%v last index = %v\n",
			ePrefix.String(),
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringStartingSearchIndex,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringLength-1)

		return numStrKernel, err
	}

	targetInputParms.RequestRemainderString = false
	targetInputParms.RequestFoundTestCharacters = false
	targetInputParms.RequestReplacementString = false

	err = targetInputParms.ValidateTargetParameters(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {

		return numStrKernel, err

	}

	decSeparatorIsNOP := decSeparatorSpec.IsNOP()

	var decSepSearchResults CharSearchDecimalSeparatorResultsDto

	foundFirstNumericChar := false
	foundRadixPoint := false
	foundMinusSign := false
	isNonZero := false

	for i := 0; i < lenPureNStr; i++ {

		targetInputParms.TargetStringCurrentSearchIndex = i

		if pureNumberString.CharsArray[i] == '-' {

			if leadingNumSymbols {
				// MUST BE A LEADING MINUS SIGN

				if !foundFirstNumericChar {

					foundMinusSign = true
				}

			} else {
				// MUST BE A TRAILING MINUS SIGN

				if foundFirstNumericChar {

					foundMinusSign = true
				}
			}

			continue
		}

		testStr := string(pureNumberString.CharsArray[i])

		//if decSeparatorIsNOP == false &&
		//	!foundRadixPoint &&
		//	(pureNumberString.CharsArray[i] < '0' ||
		//		pureNumberString.CharsArray[i] > '9')
		// Test for Radix Point
		if decSeparatorIsNOP == false &&
			!foundRadixPoint &&
			(testStr < "0" ||
				testStr > "9") {

			decSepSearchResults,
				err = decSeparatorSpec.SearchForDecimalSeparator(
				targetInputParms,
				ePrefix.XCpy(
					"decSeparatorSpec"))

			if err != nil {

				return numStrKernel, err

			}

			if decSepSearchResults.FoundDecimalSeparatorSymbols == true {

				foundRadixPoint = true

				i = decSepSearchResults.TargetStringLastSearchIndex

			}

			continue
		}

		if pureNumberString.CharsArray[i] <= '9' &&
			pureNumberString.CharsArray[i] >= '0' {

			if pureNumberString.CharsArray[i] != '0' {
				isNonZero = true
			}

			foundFirstNumericChar = true

			if foundRadixPoint == true {

				err = numStrKernel.AddFractionalDigit(
					pureNumberString.CharsArray[i],
					ePrefix.XCpy("pureNumberString"))

				if err != nil {

					return numStrKernel, err
				}

			} else {

				err = numStrKernel.AddIntegerDigit(
					pureNumberString.CharsArray[i],
					ePrefix.XCpy("pureNumberString"))

				if err != nil {

					return numStrKernel, err
				}

			}

			continue
		}

	}

	if foundFirstNumericChar == false {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'pureNumberString' is invalid!\n"+
			"It contains zero numeric digit characters.\n",
			ePrefix.String())

		return numStrKernel, err

	}

	numStrKernel.RationalizeFractionalIntegerDigits()

	if isNonZero && foundMinusSign {

		err = numStrKernel.SetNumberSign(
			NumSignVal.Negative(),
			ePrefix.XCpy("numStrKernel"))
	}

	if isNonZero && !foundMinusSign {

		err = numStrKernel.SetNumberSign(
			NumSignVal.Positive(),
			ePrefix.XCpy("numStrKernel"))
	}

	if !isNonZero {
		// MUST BE ZERO
		err = numStrKernel.SetNumberSign(
			NumSignVal.Zero(),
			ePrefix.XCpy(
				"numStrKernel"))
	}

	return numStrKernel, err
}

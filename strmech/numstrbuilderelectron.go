package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrBuilderElectron struct {
	lock *sync.Mutex
}

// extractNumRunes - Receives an array of runes and extracts the
// numeric digits as text characters.
// Set 'characterSearchLength' to minus one (-1) for end-of-string
// search.
func (nStrBuilderElectron *numStrBuilderElectron) extractNumRunes(
	targetSearchString RuneArrayDto,
	targetSearchStringName string,
	startingSearchIndex int,
	characterSearchLength int,
	negativeNumSearchSpecsCol NegNumSearchSpecCollection,
	decimalSeparatorSpec DecimalSeparatorSpec,
	numParsingTerminators RuneArrayCollection,
	requestRemainderRunesString bool,
	ePrefDto *ePref.ErrPrefixDto) (
	searchResults CharSearchNumStrParseResultsDto,
	numStrKernel NumberStrKernel,
	err error) {

	if nStrBuilderElectron.lock == nil {
		nStrBuilderElectron.lock = new(sync.Mutex)
	}

	nStrBuilderElectron.lock.Lock()

	defer nStrBuilderElectron.lock.Unlock()

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

	targetInputParms.TargetStringSearchLength = characterSearchLength

	if targetInputParms.TargetStringSearchLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'characterSearchLength' is invalid.\n"+
			"'characterSearchLength' has a value of zero.\n",
			ePrefix.String())

		return searchResults,
			numStrKernel,
			err
	}

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
		numParsingTerminators.IsNOP()

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
				err = numParsingTerminators.SearchForTextCharacters(
				targetInputParms,
				ePrefix.XCpy(
					"numParsingTerminators"))

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

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// charSearchRuneArrayResultsDtoNanobot - Provides helper methods for
// type charSearchRuneArrayResultsDtoNanobot.
//
type charSearchRuneArrayResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceRuneSearchResults' to input parameter
// 'destinationRuneSearchResults'. Both instances are of type
// CharSearchRuneArrayResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationRuneSearchResults'
// will be overwritten.
//
// Also, NO data validation is performed on 'sourceRuneSearchResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  destinationRuneSearchResults     *CharSearchRuneArrayResultsDto
//     - A pointer to a CharSearchRuneArrayResultsDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceRuneSearchResults'.
//
//       'destinationRuneSearchResults' is the destination for this
//       copy operation.
//
//
//  sourceRuneSearchResults          *CharSearchRuneArrayResultsDto
//     - A pointer to another CharSearchRuneArrayResultsDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationRuneSearchResults'.
//
//       'sourceRuneSearchResults' is the source for this copy
//       operation.
//
//       No data validation is performed on 'sourceRuneSearchResults'.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchRunesResultsDtoNanobot charSearchRuneArrayResultsDtoNanobot) copyIn(
	destinationRuneSearchResults *CharSearchRuneArrayResultsDto,
	sourceRuneSearchResults *CharSearchRuneArrayResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchRunesResultsDtoNanobot.lock == nil {
		searchRunesResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchRunesResultsDtoNanobot.lock.Lock()

	defer searchRunesResultsDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchRuneArrayResultsDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationRuneSearchResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationRuneSearchResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceRuneSearchResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceRuneSearchResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchRuneArrayResultsDtoAtom{}.ptr().
		empty(destinationRuneSearchResults)

	destinationRuneSearchResults.SearchResultsName =
		sourceRuneSearchResults.SearchResultsName

	destinationRuneSearchResults.SearchResultsFunctionChain =
		sourceRuneSearchResults.SearchResultsFunctionChain

	destinationRuneSearchResults.FoundSearchTarget =
		sourceRuneSearchResults.FoundSearchTarget

	destinationRuneSearchResults.FoundFirstNumericDigitInNumStr =
		sourceRuneSearchResults.FoundFirstNumericDigitInNumStr

	destinationRuneSearchResults.FoundDecimalSeparatorSymbols =
		sourceRuneSearchResults.FoundDecimalSeparatorSymbols

	destinationRuneSearchResults.FoundNonZeroValue =
		sourceRuneSearchResults.FoundNonZeroValue

	destinationRuneSearchResults.TargetInputParametersName =
		sourceRuneSearchResults.TargetInputParametersName

	destinationRuneSearchResults.TargetStringLength =
		sourceRuneSearchResults.TargetStringLength

	destinationRuneSearchResults.TargetStringSearchLength =
		sourceRuneSearchResults.TargetStringSearchLength

	destinationRuneSearchResults.TargetStringAdjustedSearchLength =
		sourceRuneSearchResults.TargetStringAdjustedSearchLength

	destinationRuneSearchResults.TargetStringStartingSearchIndex =
		sourceRuneSearchResults.TargetStringStartingSearchIndex

	destinationRuneSearchResults.TargetStringCurrentSearchIndex =
		sourceRuneSearchResults.TargetStringCurrentSearchIndex

	destinationRuneSearchResults.TargetStringFirstFoundIndex =
		sourceRuneSearchResults.TargetStringFirstFoundIndex

	destinationRuneSearchResults.TargetStringLastFoundIndex =
		sourceRuneSearchResults.TargetStringLastFoundIndex

	destinationRuneSearchResults.TargetStringLastSearchIndex =
		sourceRuneSearchResults.TargetStringLastSearchIndex

	destinationRuneSearchResults.TargetStringNextSearchIndex =
		sourceRuneSearchResults.TargetStringNextSearchIndex

	destinationRuneSearchResults.TargetStringDescription1 =
		sourceRuneSearchResults.TargetStringDescription1

	destinationRuneSearchResults.TargetStringDescription2 =
		sourceRuneSearchResults.TargetStringDescription2

	destinationRuneSearchResults.TestInputParametersName =
		sourceRuneSearchResults.TestInputParametersName

	destinationRuneSearchResults.TestStringName =
		sourceRuneSearchResults.TestStringName

	destinationRuneSearchResults.TestStringLength =
		sourceRuneSearchResults.TestStringLength

	destinationRuneSearchResults.TestStringLengthName =
		sourceRuneSearchResults.TestStringLengthName

	destinationRuneSearchResults.TestStringStartingIndex =
		sourceRuneSearchResults.TestStringStartingIndex

	destinationRuneSearchResults.TestStringStartingIndexName =
		sourceRuneSearchResults.TestStringStartingIndexName

	destinationRuneSearchResults.TestStringFirstFoundIndex =
		sourceRuneSearchResults.TestStringFirstFoundIndex

	destinationRuneSearchResults.TestStringLastFoundIndex =
		sourceRuneSearchResults.TestStringLastFoundIndex

	destinationRuneSearchResults.TestStringDescription1 =
		sourceRuneSearchResults.TestStringDescription1

	destinationRuneSearchResults.TestStringDescription2 =
		sourceRuneSearchResults.TestStringDescription2

	destinationRuneSearchResults.CollectionTestObjIndex =
		sourceRuneSearchResults.CollectionTestObjIndex

	destinationRuneSearchResults.TextCharSearchType =
		sourceRuneSearchResults.TextCharSearchType

	err = destinationRuneSearchResults.ReplacementString.CopyIn(
		&sourceRuneSearchResults.ReplacementString,
		ePrefix.XCpy(
			"destinationRuneSearchResults.ReplacementString<-"+
				"sourceRuneSearchResults"))

	if err != nil {
		return err
	}

	err = destinationRuneSearchResults.RemainderString.CopyIn(
		&sourceRuneSearchResults.RemainderString,
		ePrefix.XCpy(
			"destinationRuneSearchResults.RemainderString<-"+
				"sourceRuneSearchResults"))

	if err != nil {
		return err
	}

	err = destinationRuneSearchResults.FoundCharacters.CopyIn(
		&sourceRuneSearchResults.FoundCharacters,
		ePrefix.XCpy(
			"destinationRuneSearchResults.FoundCharacters<-"+
				"sourceRuneSearchResults"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'runeSearchResultsDto', a pointer to an instance of
// CharSearchRuneArrayResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'runeSearchResultsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  runeSearchResultsDto       *CharSearchRuneArrayResultsDto
//     - A pointer to an instance of CharSearchRuneArrayResultsDto.
//       A deep copy of the internal member variables will be
//       created and returned in a new instance of
//       CharSearchRuneArrayResultsDto.
//
//       No data validation is performed on 'runeSearchResultsDto'.
//
//
//  errPrefDto                      *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  deepCopyRuneSearchResults  CharSearchRuneArrayResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'runeSearchResultsDto' will be created and
//       returned in a new instance of
//       CharSearchRuneArrayResultsDto.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchRunesResultsDtoNanobot charSearchRuneArrayResultsDtoNanobot) copyOut(
	runeSearchResultsDto *CharSearchRuneArrayResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyRuneSearchResults CharSearchRuneArrayResultsDto,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchRuneArrayResultsDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyRuneSearchResults, err

	}

	if runeSearchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyRuneSearchResults, err
	}

	charSearchRuneArrayResultsDtoAtom{}.ptr().
		empty(&deepCopyRuneSearchResults)

	deepCopyRuneSearchResults.SearchResultsName =
		runeSearchResultsDto.SearchResultsName

	deepCopyRuneSearchResults.SearchResultsFunctionChain =
		runeSearchResultsDto.SearchResultsFunctionChain

	deepCopyRuneSearchResults.FoundSearchTarget =
		runeSearchResultsDto.FoundSearchTarget

	deepCopyRuneSearchResults.FoundFirstNumericDigitInNumStr =
		runeSearchResultsDto.FoundFirstNumericDigitInNumStr

	deepCopyRuneSearchResults.FoundDecimalSeparatorSymbols =
		runeSearchResultsDto.FoundDecimalSeparatorSymbols

	deepCopyRuneSearchResults.FoundNonZeroValue =
		runeSearchResultsDto.FoundNonZeroValue

	deepCopyRuneSearchResults.TargetInputParametersName =
		runeSearchResultsDto.TargetInputParametersName

	deepCopyRuneSearchResults.TargetStringLength =
		runeSearchResultsDto.TargetStringLength

	deepCopyRuneSearchResults.TargetStringSearchLength =
		runeSearchResultsDto.TargetStringSearchLength

	deepCopyRuneSearchResults.TargetStringAdjustedSearchLength =
		runeSearchResultsDto.TargetStringAdjustedSearchLength

	deepCopyRuneSearchResults.TargetStringStartingSearchIndex =
		runeSearchResultsDto.TargetStringStartingSearchIndex

	deepCopyRuneSearchResults.TargetStringCurrentSearchIndex =
		runeSearchResultsDto.TargetStringCurrentSearchIndex

	deepCopyRuneSearchResults.TargetStringFirstFoundIndex =
		runeSearchResultsDto.TargetStringFirstFoundIndex

	deepCopyRuneSearchResults.TargetStringLastFoundIndex =
		runeSearchResultsDto.TargetStringLastFoundIndex

	deepCopyRuneSearchResults.TargetStringLastSearchIndex =
		runeSearchResultsDto.TargetStringLastSearchIndex

	deepCopyRuneSearchResults.TargetStringNextSearchIndex =
		runeSearchResultsDto.TargetStringNextSearchIndex

	deepCopyRuneSearchResults.TargetStringDescription1 =
		runeSearchResultsDto.TargetStringDescription1

	deepCopyRuneSearchResults.TargetStringDescription2 =
		runeSearchResultsDto.TargetStringDescription2

	deepCopyRuneSearchResults.TestInputParametersName =
		runeSearchResultsDto.TestInputParametersName

	deepCopyRuneSearchResults.TestStringName =
		runeSearchResultsDto.TestStringName

	deepCopyRuneSearchResults.TestStringLength =
		runeSearchResultsDto.TestStringLength

	deepCopyRuneSearchResults.TestStringLengthName =
		runeSearchResultsDto.TestStringLengthName

	deepCopyRuneSearchResults.TestStringStartingIndex =
		runeSearchResultsDto.TestStringStartingIndex

	deepCopyRuneSearchResults.TestStringStartingIndexName =
		runeSearchResultsDto.TestStringStartingIndexName

	deepCopyRuneSearchResults.TestStringFirstFoundIndex =
		runeSearchResultsDto.TestStringFirstFoundIndex

	deepCopyRuneSearchResults.TestStringLastFoundIndex =
		runeSearchResultsDto.TestStringLastFoundIndex

	deepCopyRuneSearchResults.TestStringDescription1 =
		runeSearchResultsDto.TestStringDescription1

	deepCopyRuneSearchResults.TestStringDescription2 =
		runeSearchResultsDto.TestStringDescription2

	deepCopyRuneSearchResults.CollectionTestObjIndex =
		runeSearchResultsDto.CollectionTestObjIndex

	deepCopyRuneSearchResults.TextCharSearchType =
		runeSearchResultsDto.TextCharSearchType

	err = deepCopyRuneSearchResults.ReplacementString.CopyIn(
		&runeSearchResultsDto.ReplacementString,
		ePrefix.XCpy(
			"deepCopyRuneSearchResults.ReplacementString<-"+
				"runeSearchResultsDto"))

	if err != nil {
		return deepCopyRuneSearchResults, err
	}

	err = deepCopyRuneSearchResults.RemainderString.CopyIn(
		&runeSearchResultsDto.RemainderString,
		ePrefix.XCpy(
			"deepCopyRuneSearchResults.RemainderString<-"+
				"runeSearchResultsDto"))

	if err != nil {
		return deepCopyRuneSearchResults, err
	}

	err = deepCopyRuneSearchResults.FoundCharacters.CopyIn(
		&runeSearchResultsDto.FoundCharacters,
		ePrefix.XCpy(
			"deepCopyRuneSearchResults.FoundCharacters<-"+
				"runeSearchResultsDto"))

	return deepCopyRuneSearchResults, err
}

// ptr - Returns a pointer to a new instance of
// charSearchResultsDtoNanobot.
//
func (searchRunesResultsDtoNanobot charSearchRuneArrayResultsDtoNanobot) ptr() *charSearchRuneArrayResultsDtoNanobot {

	if searchRunesResultsDtoNanobot.lock == nil {
		searchRunesResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchRunesResultsDtoNanobot.lock.Lock()

	defer searchRunesResultsDtoNanobot.lock.Unlock()

	return &charSearchRuneArrayResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

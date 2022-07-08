package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// charSearchDecimalSeparatorResultsDtoNanobot - Provides helper
// methods for type CharSearchDecimalSeparatorResultsDto.
//
type charSearchDecimalSeparatorResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceDecSepResults' to input parameter
// 'destinationDecSepResults'. Both instances are of type
// CharSearchDecimalSeparatorResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationDecSepResults'
// will be overwritten.
//
// Also, NO data validation is performed on 'sourceDecSepResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  destinationDecSepRes   *CharSearchDecimalSeparatorResultsDto
//     - A pointer to a CharSearchDecimalSeparatorResultsDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceDecSepResults'.
//
//       'destinationDecSepResults' is the destination for this
//       copy operation.
//
//
//  sourceDecSepResults    *CharSearchDecimalSeparatorResultsDto
//     - A pointer to another CharSearchDecimalSeparatorResultsDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationDecSepResults'.
//
//       'sourceDecSepResults' is the source for this copy
//       operation.
//
//       No data validation is performed on 'sourceDecSepResults'.
//
//
//  errPrefDto             *ePref.ErrPrefixDto
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
func (searchDecimalSepResultsNanobot *charSearchDecimalSeparatorResultsDtoNanobot) copyIn(
	destinationDecSepResults *CharSearchDecimalSeparatorResultsDto,
	sourceDecSepResults *CharSearchDecimalSeparatorResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchDecimalSepResultsNanobot.lock == nil {
		searchDecimalSepResultsNanobot.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsNanobot.lock.Lock()

	defer searchDecimalSepResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchDecimalSeparatorResultsDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationDecSepResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationDecSepResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceDecSepResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceDecSepResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchDecimalSeparatorResultsDtoAtom{}.ptr().empty(
		destinationDecSepResults)

	destinationDecSepResults.SearchResultsName =
		sourceDecSepResults.SearchResultsName

	destinationDecSepResults.SearchResultsFunctionChain =
		sourceDecSepResults.SearchResultsFunctionChain

	destinationDecSepResults.FoundDecimalSeparatorSymbols =
		sourceDecSepResults.FoundDecimalSeparatorSymbols

	destinationDecSepResults.FoundDecimalSepSymbolsOnPreviousSearch =
		sourceDecSepResults.FoundDecimalSeparatorSymbols

	destinationDecSepResults.FoundFirstNumericDigitInNumStr =
		sourceDecSepResults.FoundFirstNumericDigitInNumStr

	destinationDecSepResults.FoundNonZeroValue =
		sourceDecSepResults.FoundNonZeroValue

	destinationDecSepResults.TargetInputParametersName =
		sourceDecSepResults.TargetInputParametersName

	destinationDecSepResults.TargetStringLength =
		sourceDecSepResults.TargetStringLength

	destinationDecSepResults.TargetStringSearchLength =
		sourceDecSepResults.TargetStringSearchLength

	destinationDecSepResults.TargetStringAdjustedSearchLength =
		sourceDecSepResults.TargetStringAdjustedSearchLength

	destinationDecSepResults.TargetStringStartingSearchIndex =
		sourceDecSepResults.TargetStringStartingSearchIndex

	destinationDecSepResults.TargetStringCurrentSearchIndex =
		sourceDecSepResults.TargetStringCurrentSearchIndex

	destinationDecSepResults.TargetStringFirstFoundIndex =
		sourceDecSepResults.TargetStringFirstFoundIndex

	destinationDecSepResults.TargetStringLastFoundIndex =
		sourceDecSepResults.TargetStringLastFoundIndex

	destinationDecSepResults.TargetStringLastSearchIndex =
		sourceDecSepResults.TargetStringLastSearchIndex

	destinationDecSepResults.TargetStringNextSearchIndex =
		sourceDecSepResults.TargetStringNextSearchIndex

	destinationDecSepResults.TargetStringDescription1 =
		sourceDecSepResults.TargetStringDescription1

	destinationDecSepResults.TargetStringDescription2 =
		sourceDecSepResults.TargetStringDescription2

	destinationDecSepResults.TestInputParametersName =
		sourceDecSepResults.TestInputParametersName

	destinationDecSepResults.TestStringName =
		sourceDecSepResults.TestStringName

	destinationDecSepResults.TestStringLength =
		sourceDecSepResults.TestStringLength

	destinationDecSepResults.TestStringLengthName =
		sourceDecSepResults.TestStringLengthName

	destinationDecSepResults.TestStringStartingIndex =
		sourceDecSepResults.TestStringStartingIndex

	destinationDecSepResults.TestStringStartingIndexName =
		sourceDecSepResults.TestStringStartingIndexName

	destinationDecSepResults.TestStringFirstFoundIndex =
		sourceDecSepResults.TestStringFirstFoundIndex

	destinationDecSepResults.TestStringLastFoundIndex =
		sourceDecSepResults.TestStringLastFoundIndex

	destinationDecSepResults.TestStringDescription1 =
		sourceDecSepResults.TestStringDescription1

	destinationDecSepResults.TestStringDescription2 =
		sourceDecSepResults.TestStringDescription2

	destinationDecSepResults.CollectionTestObjIndex =
		sourceDecSepResults.CollectionTestObjIndex

	destinationDecSepResults.NumValueType =
		sourceDecSepResults.NumValueType

	destinationDecSepResults.NumSymbolLocation =
		sourceDecSepResults.NumSymbolLocation

	destinationDecSepResults.TextCharSearchType =
		sourceDecSepResults.TextCharSearchType

	err = destinationDecSepResults.DecimalSeparatorSymbolsSpec.CopyIn(
		&sourceDecSepResults.DecimalSeparatorSymbolsSpec,
		ePrefix.XCpy(
			"destinationDecSepResults<-sourceDecSepResults"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'decimalSeparatorResults', a pointer to an instance of
// CharSearchDecimalSeparatorResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'decimalSeparatorResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorResults  *CharSearchDecimalSeparatorResultsDto
//     - A pointer to an instance of
//       CharSearchDecimalSeparatorResultsDto. A deep copy of the
//       internal member variables contained in this instance will
//       be created and returned in a new instance of
//       CharSearchDecimalSeparatorResultsDto.
//
//       No data validation is performed on
//       'decimalSeparatorResults'.
//
//
//  errPrefDto               *ePref.ErrPrefixDto
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
//  deepCopyNegNumResultsDto CharSearchDecimalSeparatorResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'decimalSeparatorResults' will be created
//       and returned in a new instance of
//       CharSearchDecimalSeparatorResultsDto.
//
//
//  err                      error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchDecimalSepResultsNanobot *charSearchDecimalSeparatorResultsDtoNanobot) copyOut(
	decimalSeparatorResults *CharSearchDecimalSeparatorResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyDecSepResults CharSearchDecimalSeparatorResultsDto,
	err error) {

	if searchDecimalSepResultsNanobot.lock == nil {
		searchDecimalSepResultsNanobot.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsNanobot.lock.Lock()

	defer searchDecimalSepResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchDecimalSeparatorResultsDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyDecSepResults, err

	}

	if decimalSeparatorResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decimalSeparatorResults' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyDecSepResults, err
	}

	charSearchDecimalSeparatorResultsDtoAtom{}.ptr().empty(
		&deepCopyDecSepResults)

	deepCopyDecSepResults.SearchResultsName =
		decimalSeparatorResults.SearchResultsName

	deepCopyDecSepResults.SearchResultsFunctionChain =
		decimalSeparatorResults.SearchResultsFunctionChain

	deepCopyDecSepResults.FoundDecimalSeparatorSymbols =
		decimalSeparatorResults.FoundDecimalSeparatorSymbols

	deepCopyDecSepResults.FoundDecimalSepSymbolsOnPreviousSearch =
		decimalSeparatorResults.FoundDecimalSeparatorSymbols

	deepCopyDecSepResults.FoundFirstNumericDigitInNumStr =
		decimalSeparatorResults.FoundFirstNumericDigitInNumStr

	deepCopyDecSepResults.FoundNonZeroValue =
		decimalSeparatorResults.FoundNonZeroValue

	deepCopyDecSepResults.TargetInputParametersName =
		decimalSeparatorResults.TargetInputParametersName

	deepCopyDecSepResults.TargetStringLength =
		decimalSeparatorResults.TargetStringLength

	deepCopyDecSepResults.TargetStringSearchLength =
		decimalSeparatorResults.TargetStringSearchLength

	deepCopyDecSepResults.TargetStringAdjustedSearchLength =
		decimalSeparatorResults.TargetStringAdjustedSearchLength

	deepCopyDecSepResults.TargetStringStartingSearchIndex =
		decimalSeparatorResults.TargetStringStartingSearchIndex

	deepCopyDecSepResults.TargetStringCurrentSearchIndex =
		decimalSeparatorResults.TargetStringCurrentSearchIndex

	deepCopyDecSepResults.TargetStringFirstFoundIndex =
		decimalSeparatorResults.TargetStringFirstFoundIndex

	deepCopyDecSepResults.TargetStringLastFoundIndex =
		decimalSeparatorResults.TargetStringLastFoundIndex

	deepCopyDecSepResults.TargetStringLastSearchIndex =
		decimalSeparatorResults.TargetStringLastSearchIndex

	deepCopyDecSepResults.TargetStringNextSearchIndex =
		decimalSeparatorResults.TargetStringNextSearchIndex

	deepCopyDecSepResults.TargetStringDescription1 =
		decimalSeparatorResults.TargetStringDescription1

	deepCopyDecSepResults.TargetStringDescription2 =
		decimalSeparatorResults.TargetStringDescription2

	deepCopyDecSepResults.TestInputParametersName =
		decimalSeparatorResults.TestInputParametersName

	deepCopyDecSepResults.TestStringName =
		decimalSeparatorResults.TestStringName

	deepCopyDecSepResults.TestStringLength =
		decimalSeparatorResults.TestStringLength

	deepCopyDecSepResults.TestStringLengthName =
		decimalSeparatorResults.TestStringLengthName

	deepCopyDecSepResults.TestStringStartingIndex =
		decimalSeparatorResults.TestStringStartingIndex

	deepCopyDecSepResults.TestStringStartingIndexName =
		decimalSeparatorResults.TestStringStartingIndexName

	deepCopyDecSepResults.TestStringFirstFoundIndex =
		decimalSeparatorResults.TestStringFirstFoundIndex

	deepCopyDecSepResults.TestStringLastFoundIndex =
		decimalSeparatorResults.TestStringLastFoundIndex

	deepCopyDecSepResults.TestStringDescription1 =
		decimalSeparatorResults.TestStringDescription1

	deepCopyDecSepResults.TestStringDescription2 =
		decimalSeparatorResults.TestStringDescription2

	deepCopyDecSepResults.CollectionTestObjIndex =
		decimalSeparatorResults.CollectionTestObjIndex

	deepCopyDecSepResults.NumValueType =
		decimalSeparatorResults.NumValueType

	deepCopyDecSepResults.NumSymbolLocation =
		decimalSeparatorResults.NumSymbolLocation

	deepCopyDecSepResults.TextCharSearchType =
		decimalSeparatorResults.TextCharSearchType

	err = deepCopyDecSepResults.DecimalSeparatorSymbolsSpec.CopyIn(
		&decimalSeparatorResults.DecimalSeparatorSymbolsSpec,
		ePrefix.XCpy(
			"deepCopyDecSepResults<-decimalSeparatorResults"))
	return deepCopyDecSepResults, err
}

// ptr - Returns a pointer to a new instance of
// charSearchDecimalSeparatorResultsDtoNanobot.
//
func (searchDecimalSepResultsNanobot charSearchDecimalSeparatorResultsDtoNanobot) ptr() *charSearchDecimalSeparatorResultsDtoNanobot {

	if searchDecimalSepResultsNanobot.lock == nil {
		searchDecimalSepResultsNanobot.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsNanobot.lock.Lock()

	defer searchDecimalSepResultsNanobot.lock.Unlock()

	return &charSearchDecimalSeparatorResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

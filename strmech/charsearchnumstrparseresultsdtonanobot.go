package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// charSearchNumStrParseResultsDtoNanobot - Provides helper
// methods for type CharSearchNumStrParseResultsDto.
//
type charSearchNumStrParseResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceNumStrParseResults' to input parameter
// 'destinationNumStrParseResults'. Both instances are of type
// CharSearchNumStrParseResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in
// 'destinationNumStrParseResults' will be overwritten.
//
// Also, NO data validation is performed on
// 'sourceNumStrParseResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  destinationNumStrParseResults  *CharSearchNumStrParseResultsDto
//     - A pointer to a CharSearchNumStrParseResultsDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceNumStrParseResults'.
//
//       'destinationNumStrParseResults' is the destination for this
//       copy operation.
//
//
//  sourceNumStrParseResults       *CharSearchNumStrParseResultsDto
//     - A pointer to another CharSearchNumStrParseResultsDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationNumStrParseResults'.
//
//       'sourceNumStrParseResults' is the source for this copy
//       operation.
//
//       No data validation is performed on 'sourceNumStrParseResults'.
//
//
//  errPrefDto                     *ePref.ErrPrefixDto
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
func (searchNumStrParseResultsNanobot charSearchNumStrParseResultsDtoNanobot) copyIn(
	destinationNumStrParseResults *CharSearchNumStrParseResultsDto,
	sourceNumStrParseResults *CharSearchNumStrParseResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchNumStrParseResultsNanobot.lock == nil {
		searchNumStrParseResultsNanobot.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsNanobot.lock.Lock()

	defer searchNumStrParseResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNumStrParseResultsDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationNumStrParseResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationNumStrParseResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceNumStrParseResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceNumStrParseResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchNumStrParseResultsDtoAtom{}.ptr().
		empty(destinationNumStrParseResults)

	destinationNumStrParseResults.SearchResultsName =
		sourceNumStrParseResults.SearchResultsName

	destinationNumStrParseResults.SearchResultsFunctionChain =
		sourceNumStrParseResults.SearchResultsFunctionChain

	destinationNumStrParseResults.FoundNumericDigits =
		sourceNumStrParseResults.FoundNumericDigits

	destinationNumStrParseResults.FoundNonZeroValue =
		sourceNumStrParseResults.FoundNonZeroValue

	destinationNumStrParseResults.FoundDecimalSeparatorSymbols =
		sourceNumStrParseResults.FoundDecimalSeparatorSymbols

	destinationNumStrParseResults.FoundDecimalDigits =
		sourceNumStrParseResults.FoundDecimalDigits

	destinationNumStrParseResults.NumSignValue =
		sourceNumStrParseResults.NumSignValue

	destinationNumStrParseResults.NumValueType =
		sourceNumStrParseResults.NumValueType

	err = destinationNumStrParseResults.RemainderString.
		CopyIn(
			&sourceNumStrParseResults.RemainderString,
			ePrefix.XCpy(
				"destinationNumStrParseResults"+
					"<-sourceNumStrParseResults"))

	if err != nil {
		return err
	}

	err = destinationNumStrParseResults.DecimalSeparatorSearchResults.
		CopyIn(
			&sourceNumStrParseResults.DecimalSeparatorSearchResults,
			ePrefix.XCpy(
				"destinationNumStrParseResults"+
					"<-sourceNumStrParseResults"))

	if err != nil {
		return err
	}

	err = destinationNumStrParseResults.NegativeNumberSymbolSearchResults.
		CopyIn(
			&sourceNumStrParseResults.NegativeNumberSymbolSearchResults,
			ePrefix.XCpy(
				"destinationNumStrParseResults"+
					"<-sourceNumStrParseResults"))

	if err != nil {
		return err
	}

	err = destinationNumStrParseResults.ParsingTerminatorSearchResults.
		CopyIn(
			&sourceNumStrParseResults.ParsingTerminatorSearchResults,
			ePrefix.XCpy(
				"destinationNumStrParseResults"+
					"<-sourceNumStrParseResults"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'numStrParseResults', a pointer to an instance of
// CharSearchNumStrParseResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'numStrParseResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrParseResults            *CharSearchNumStrParseResultsDto
//     - A pointer to an instance of
//       CharSearchNumStrParseResultsDto. A deep copy of the
//       internal member variables contained in this instance will
//       be created and returned in a new instance of
//       CharSearchNumStrParseResultsDto.
//
//       No data validation is performed on 'numStrParseResults'.
//
//
//  errPrefDto                    *ePref.ErrPrefixDto
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
//  deepCopyNumStrParseResults    CharSearchNumStrParseResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'numStrParseResults' will be created and
//       returned in a new instance of
//       CharSearchNumStrParseResultsDto.
//
//
//  err                           error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchNumStrParseResultsNanobot charSearchNumStrParseResultsDtoNanobot) copyOut(
	numStrParseResults *CharSearchNumStrParseResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyNumStrParseResults CharSearchNumStrParseResultsDto,
	err error) {

	if searchNumStrParseResultsNanobot.lock == nil {
		searchNumStrParseResultsNanobot.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsNanobot.lock.Lock()

	defer searchNumStrParseResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNumStrParseResultsDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyNumStrParseResults, err

	}

	if numStrParseResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrParseResults' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyNumStrParseResults, err
	}

	charSearchNumStrParseResultsDtoAtom{}.ptr().
		empty(&deepCopyNumStrParseResults)

	deepCopyNumStrParseResults.SearchResultsName =
		numStrParseResults.SearchResultsName

	deepCopyNumStrParseResults.SearchResultsFunctionChain =
		numStrParseResults.SearchResultsFunctionChain

	deepCopyNumStrParseResults.FoundNumericDigits =
		numStrParseResults.FoundNumericDigits

	deepCopyNumStrParseResults.FoundNonZeroValue =
		numStrParseResults.FoundNonZeroValue

	deepCopyNumStrParseResults.FoundDecimalSeparatorSymbols =
		numStrParseResults.FoundDecimalSeparatorSymbols

	deepCopyNumStrParseResults.FoundDecimalDigits =
		numStrParseResults.FoundDecimalDigits

	deepCopyNumStrParseResults.NumSignValue =
		numStrParseResults.NumSignValue

	deepCopyNumStrParseResults.NumValueType =
		numStrParseResults.NumValueType

	err = deepCopyNumStrParseResults.RemainderString.
		CopyIn(
			&numStrParseResults.RemainderString,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-numStrParseResults"))

	if err != nil {
		return deepCopyNumStrParseResults, err
	}

	err = deepCopyNumStrParseResults.DecimalSeparatorSearchResults.
		CopyIn(
			&numStrParseResults.DecimalSeparatorSearchResults,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-numStrParseResults"))

	if err != nil {
		return deepCopyNumStrParseResults, err
	}

	err = deepCopyNumStrParseResults.NegativeNumberSymbolSearchResults.
		CopyIn(
			&numStrParseResults.NegativeNumberSymbolSearchResults,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-numStrParseResults"))

	if err != nil {
		return deepCopyNumStrParseResults, err
	}

	err = deepCopyNumStrParseResults.ParsingTerminatorSearchResults.
		CopyIn(
			&numStrParseResults.ParsingTerminatorSearchResults,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-numStrParseResults"))

	return deepCopyNumStrParseResults, err
}

// ptr - Returns a pointer to a new instance of
// charSearchNegNumResultsDtoNanobot.
//
func (searchNumStrParseResultsNanobot charSearchNumStrParseResultsDtoNanobot) ptr() *charSearchNumStrParseResultsDtoNanobot {

	if searchNumStrParseResultsNanobot.lock == nil {
		searchNumStrParseResultsNanobot.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsNanobot.lock.Lock()

	defer searchNumStrParseResultsNanobot.lock.Unlock()

	return &charSearchNumStrParseResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

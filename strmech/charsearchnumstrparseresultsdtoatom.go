package strmech

import "sync"

// charSearchNumStrParseResultsDtoAtom - Provides helper methods for
// type CharSearchNumStrParseResultsDto.
//
type charSearchNumStrParseResultsDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchNumStrParseResultsDto and proceeds to reset the
// data values for member variables to their initial or zero
// values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'numStrParseResults' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrParseResults     *CharSearchNumStrParseResultsDto
//     - A pointer to an instance of
//       CharSearchNumStrParseResultsDto. All the internal
//       member variables contained in this instance will be
//       deleted and reset to their zero values.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchNumStrParseResultsAtom *charSearchNumStrParseResultsDtoAtom) empty(
	numStrParseResults *CharSearchNumStrParseResultsDto) {

	if searchNumStrParseResultsAtom.lock == nil {
		searchNumStrParseResultsAtom.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsAtom.lock.Lock()

	defer searchNumStrParseResultsAtom.lock.Unlock()

	if numStrParseResults == nil {
		return
	}

	numStrParseResults.SearchResultsName = ""

	numStrParseResults.SearchResultsFunctionChain = ""

	numStrParseResults.FoundNumericDigits = false

	numStrParseResults.FoundNonZeroValue = false

	numStrParseResults.FoundDecimalSeparatorSymbols = false

	numStrParseResults.FoundDecimalDigits = false

	numStrParseResults.NumSignValue = NumSignVal.None()

	numStrParseResults.NumValueType = NumValType.None()

	numStrParseResults.RemainderString.Empty()

	numStrParseResults.DecimalSeparatorSearchResults.Empty()

	numStrParseResults.NegativeNumberSymbolSearchResults.Empty()

	numStrParseResults.ParsingTerminatorSearchResults.Empty()

	return
}

// ptr - Returns a pointer to a new instance of
// charSearchNumStrParseResultsDtoAtom.
//
func (searchNumStrParseResultsAtom charSearchNumStrParseResultsDtoAtom) ptr() *charSearchNumStrParseResultsDtoAtom {

	if searchNumStrParseResultsAtom.lock == nil {
		searchNumStrParseResultsAtom.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsAtom.lock.Lock()

	defer searchNumStrParseResultsAtom.lock.Unlock()

	return &charSearchNumStrParseResultsDtoAtom{
		lock: new(sync.Mutex),
	}
}

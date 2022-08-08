package strmech

import "sync"

// charSearchNumStrParseResultsDtoAtom - Provides helper methods for
// type CharSearchNumStrParseResultsDto.
type charSearchNumStrParseResultsDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchNumStrParseResultsDto and proceeds to reset the
// data values for member variables to their initial or zero
// values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'numStrParseResults' will be deleted and reset to their zero
// values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	numStrParseResults     *CharSearchNumStrParseResultsDto
//	   - A pointer to an instance of
//	     CharSearchNumStrParseResultsDto. All the internal
//	     member variables contained in this instance will be
//	     deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
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

	numStrParseResults.TargetSearchString.Empty()

	numStrParseResults.TargetStringSearchLength = -1

	numStrParseResults.TargetStringAdjustedSearchLength = -1

	numStrParseResults.TargetStringStartingSearchIndex = -1

	numStrParseResults.TargetStringLastSearchIndex = -1

	numStrParseResults.TargetStringLastSearchIndex = -1

	numStrParseResults.ReasonForSearchTermination =
		CharSearchTermType.None()

	numStrParseResults.FoundNumericDigits = false

	numStrParseResults.FoundNonZeroValue = false

	numStrParseResults.FoundDecimalSeparatorSymbols = false

	numStrParseResults.FoundIntegerDigits = false

	numStrParseResults.IdentifiedIntegerDigits = ""

	numStrParseResults.FoundDecimalDigits = false

	numStrParseResults.IdentifiedFractionalDigits = ""

	numStrParseResults.NumSignValue = NumSignVal.None()

	numStrParseResults.NumValueType = NumValType.None()

	numStrParseResults.RemainderString.Empty()

	numStrParseResults.DecimalSeparatorSearchResults.Empty()

	numStrParseResults.NegativeNumberSymbolSearchResults.Empty()

	numStrParseResults.ParsingTerminatorSearchResults.Empty()

	return
}

// equal - Receives a pointer to two instances of
// CharSearchNumStrParseResultsDto and proceeds to compare
// their member variables in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numStrParseResults1    *CharSearchNumStrParseResultsDto
//	   - An instance of CharSearchNumStrParseResultsDto.
//	     Internal member variables from 'numStrParseResults1'
//	     will be compared to those of 'numStrParseResults2' to
//	     determine if both instances are equivalent.
//
//
//	numStrParseResults2    *CharSearchNumStrParseResultsDto
//	   - An instance of CharSearchNumStrParseResultsDto.
//	     Internal member variables from 'numStrParseResults2'
//	     will be compared to those of 'numStrParseResults1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'numStrParseResults1' and
//	     'numStrParseResults2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling function.
func (searchNumStrParseResultsAtom *charSearchNumStrParseResultsDtoAtom) equal(
	numStrParseResults1 *CharSearchNumStrParseResultsDto,
	numStrParseResults2 *CharSearchNumStrParseResultsDto) bool {

	if searchNumStrParseResultsAtom.lock == nil {
		searchNumStrParseResultsAtom.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsAtom.lock.Lock()

	defer searchNumStrParseResultsAtom.lock.Unlock()

	if numStrParseResults1.SearchResultsName !=
		numStrParseResults2.SearchResultsName {

		return false
	}

	if numStrParseResults1.SearchResultsFunctionChain !=
		numStrParseResults2.SearchResultsFunctionChain {

		return false
	}

	if !numStrParseResults1.TargetSearchString.Equal(
		&numStrParseResults2.TargetSearchString) {

		return false
	}

	if numStrParseResults1.TargetStringSearchLength !=
		numStrParseResults2.TargetStringSearchLength {

		return false
	}

	if numStrParseResults1.TargetStringAdjustedSearchLength !=
		numStrParseResults2.TargetStringAdjustedSearchLength {

		return false
	}

	if numStrParseResults1.TargetStringStartingSearchIndex !=
		numStrParseResults2.TargetStringStartingSearchIndex {

		return false
	}

	if numStrParseResults1.TargetStringLastSearchIndex !=
		numStrParseResults2.TargetStringLastSearchIndex {

		return false
	}

	if numStrParseResults1.TargetStringNextSearchIndex !=
		numStrParseResults2.TargetStringNextSearchIndex {

		return false
	}

	if numStrParseResults1.ReasonForSearchTermination !=
		numStrParseResults2.ReasonForSearchTermination {

		return false
	}

	if numStrParseResults1.FoundNumericDigits !=
		numStrParseResults2.FoundNumericDigits {

		return false
	}

	if numStrParseResults1.FoundNonZeroValue !=
		numStrParseResults2.FoundNonZeroValue {

		return false
	}

	if numStrParseResults1.FoundDecimalSeparatorSymbols !=
		numStrParseResults2.FoundDecimalSeparatorSymbols {

		return false
	}

	if numStrParseResults1.FoundIntegerDigits !=
		numStrParseResults2.FoundIntegerDigits {

		return false
	}

	if numStrParseResults1.IdentifiedIntegerDigits !=
		numStrParseResults2.IdentifiedIntegerDigits {

		return false
	}

	if numStrParseResults1.FoundDecimalDigits !=
		numStrParseResults2.FoundDecimalDigits {

		return false
	}

	if numStrParseResults1.IdentifiedFractionalDigits !=
		numStrParseResults2.IdentifiedFractionalDigits {

		return false
	}

	if numStrParseResults1.NumSignValue !=
		numStrParseResults2.NumSignValue {

		return false
	}

	if numStrParseResults1.NumValueType !=
		numStrParseResults2.NumValueType {

		return false
	}

	if !numStrParseResults1.RemainderString.Equal(
		&numStrParseResults2.RemainderString) {

		return false
	}

	if !numStrParseResults1.DecimalSeparatorSearchResults.Equal(
		&numStrParseResults2.DecimalSeparatorSearchResults) {

		return false
	}

	if !numStrParseResults1.NegativeNumberSymbolSearchResults.Equal(
		&numStrParseResults1.NegativeNumberSymbolSearchResults) {

		return false
	}

	if !numStrParseResults1.ParsingTerminatorSearchResults.Equal(
		&numStrParseResults1.ParsingTerminatorSearchResults) {

		return false
	}

	return true
}

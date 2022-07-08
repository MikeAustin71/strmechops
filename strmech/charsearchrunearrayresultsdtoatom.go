package strmech

import "sync"

// charSearchRuneArrayResultsDtoAtom - Provides helper methods for
// type charSearchRuneArrayResultsDtoAtom.
//
type charSearchRuneArrayResultsDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchRuneArrayResultsDto and proceeds to reset the data
// values for member variables to their initial or zero values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'runeSearchResults' will be deleted and reset to their zero
// values. All Array Index values will be set to minus one (-1).
// Valid Array Indexes have values greater than minus one (-1).
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  runeSearchResults          *CharSearchRuneArrayResultsDto
//     - A pointer to an instance of
//       CharSearchRuneArrayResultsDto. All the internal member
//       variables contained in this instance will be deleted and
//       reset to their zero values.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchRunesResultsDtoAtom *charSearchRuneArrayResultsDtoAtom) empty(
	runeSearchResults *CharSearchRuneArrayResultsDto) {

	if searchRunesResultsDtoAtom.lock == nil {
		searchRunesResultsDtoAtom.lock = new(sync.Mutex)
	}

	searchRunesResultsDtoAtom.lock.Lock()

	defer searchRunesResultsDtoAtom.lock.Unlock()

	if runeSearchResults == nil {
		return
	}

	runeSearchResults.SearchResultsName = ""

	runeSearchResults.SearchResultsFunctionChain = ""

	runeSearchResults.FoundSearchTarget = false

	runeSearchResults.FoundFirstNumericDigitInNumStr = false

	runeSearchResults.FoundDecimalSeparatorSymbols = false

	runeSearchResults.FoundNonZeroValue = false

	runeSearchResults.TargetInputParametersName = ""

	runeSearchResults.TargetStringLength = -1

	runeSearchResults.TargetStringSearchLength = -1

	runeSearchResults.TargetStringAdjustedSearchLength = -1

	runeSearchResults.TargetStringStartingSearchIndex = -1

	runeSearchResults.TargetStringCurrentSearchIndex = -1

	runeSearchResults.TargetStringFirstFoundIndex = -1

	runeSearchResults.TargetStringLastFoundIndex = -1

	runeSearchResults.TargetStringLastSearchIndex = -1

	runeSearchResults.TargetStringNextSearchIndex = -1

	runeSearchResults.TargetStringDescription1 = ""

	runeSearchResults.TargetStringDescription2 = ""

	runeSearchResults.TestInputParametersName = ""

	runeSearchResults.TestStringName = ""

	runeSearchResults.TestStringLength = -1

	runeSearchResults.TestStringLengthName = ""

	runeSearchResults.TestStringStartingIndex = -1

	runeSearchResults.TestStringStartingIndexName = ""

	runeSearchResults.TestStringFirstFoundIndex = -1

	runeSearchResults.TestStringLastFoundIndex = -1

	runeSearchResults.TestStringDescription1 = ""

	runeSearchResults.TestStringDescription2 = ""

	runeSearchResults.CollectionTestObjIndex = -1

	runeSearchResults.TextCharSearchType =
		CharSearchType.None()

	runeSearchResults.ReplacementString.Empty()

	runeSearchResults.RemainderString.Empty()

	runeSearchResults.FoundCharacters.Empty()

	return
}

// equal - Receives a pointer to two instances of
// CharSearchRuneArrayResultsDto and proceeds to compare their
// member variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  runeSearchResults1         *CharSearchRuneArrayResultsDto
//     - A pointer to an instance of CharSearchRuneArrayResultsDto.
//       Internal member variables from 'runeSearchResults1' will
//       be compared to those of 'runeSearchResults2' to determine
//       if both instances are equivalent.
//
//
//  runeSearchResults2         *CharSearchRuneArrayResultsDto
//     - A pointer to an instance of CharSearchRuneArrayResultsDto.
//       Internal member variables from 'runeSearchResults2' will
//       be compared to those of 'runeSearchResults1' to determine
//       if both instances are equivalent.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of 'runeSearchResults1' and
//       'runeSearchResults2' shows that all internal member
//       variables are equivalent, this method will return a
//       boolean value of 'true'.
//
//       If the two instances are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchRunesResultsDtoAtom *charSearchRuneArrayResultsDtoAtom) equal(
	runeSearchResults1 *CharSearchRuneArrayResultsDto,
	runeSearchResults2 *CharSearchRuneArrayResultsDto) bool {

	if searchRunesResultsDtoAtom.lock == nil {
		searchRunesResultsDtoAtom.lock = new(sync.Mutex)
	}

	searchRunesResultsDtoAtom.lock.Lock()

	defer searchRunesResultsDtoAtom.lock.Unlock()

	if runeSearchResults1 == nil ||
		runeSearchResults2 == nil {
		return false
	}

	if runeSearchResults1.SearchResultsName !=
		runeSearchResults2.SearchResultsName {

		return false
	}

	if runeSearchResults1.SearchResultsFunctionChain !=
		runeSearchResults2.SearchResultsFunctionChain {

		return false
	}

	if runeSearchResults1.FoundSearchTarget !=
		runeSearchResults2.FoundSearchTarget {

		return false
	}

	if runeSearchResults1.FoundFirstNumericDigitInNumStr !=
		runeSearchResults2.FoundFirstNumericDigitInNumStr {

		return false
	}

	if runeSearchResults1.FoundDecimalSeparatorSymbols !=
		runeSearchResults2.FoundDecimalSeparatorSymbols {

		return false
	}

	if runeSearchResults1.FoundNonZeroValue !=
		runeSearchResults2.FoundNonZeroValue {

		return false
	}

	if runeSearchResults1.TargetInputParametersName !=
		runeSearchResults2.TargetInputParametersName {

		return false
	}

	if runeSearchResults1.TargetStringLength !=
		runeSearchResults2.TargetStringLength {

		return false
	}

	if runeSearchResults1.TargetStringSearchLength !=
		runeSearchResults2.TargetStringSearchLength {

		return false
	}

	if runeSearchResults1.TargetStringAdjustedSearchLength !=
		runeSearchResults2.TargetStringAdjustedSearchLength {

		return false
	}

	if runeSearchResults1.TargetStringStartingSearchIndex !=
		runeSearchResults2.TargetStringStartingSearchIndex {

		return false
	}

	if runeSearchResults1.TargetStringCurrentSearchIndex !=
		runeSearchResults2.TargetStringCurrentSearchIndex {

		return false
	}

	if runeSearchResults1.TargetStringFirstFoundIndex !=
		runeSearchResults2.TargetStringFirstFoundIndex {

		return false
	}

	if runeSearchResults1.TargetStringLastFoundIndex !=
		runeSearchResults2.TargetStringLastFoundIndex {

		return false
	}

	if runeSearchResults1.TargetStringLastSearchIndex !=
		runeSearchResults2.TargetStringLastSearchIndex {

		return false
	}

	if runeSearchResults1.TargetStringNextSearchIndex !=
		runeSearchResults2.TargetStringNextSearchIndex {

		return false
	}

	if runeSearchResults1.TargetStringDescription1 !=
		runeSearchResults2.TargetStringDescription1 {

		return false
	}

	if runeSearchResults1.TargetStringDescription2 !=
		runeSearchResults2.TargetStringDescription2 {

		return false
	}

	if runeSearchResults1.TestInputParametersName !=
		runeSearchResults2.TestInputParametersName {

		return false
	}

	if runeSearchResults1.TestStringName !=
		runeSearchResults2.TestStringName {

		return false
	}

	if runeSearchResults1.TestStringLength !=
		runeSearchResults2.TestStringLength {

		return false
	}

	if runeSearchResults1.TestStringLengthName !=
		runeSearchResults2.TestStringLengthName {

		return false
	}

	if runeSearchResults1.TestStringStartingIndex !=
		runeSearchResults2.TestStringStartingIndex {

		return false
	}

	if runeSearchResults1.TestStringStartingIndexName !=
		runeSearchResults2.TestStringStartingIndexName {

		return false
	}

	if runeSearchResults1.TestStringFirstFoundIndex !=
		runeSearchResults2.TestStringFirstFoundIndex {

		return false
	}

	if runeSearchResults1.TestStringLastFoundIndex !=
		runeSearchResults2.TestStringLastFoundIndex {

		return false
	}

	if runeSearchResults1.TestStringDescription1 !=
		runeSearchResults2.TestStringDescription1 {

		return false
	}

	if runeSearchResults1.TestStringDescription2 !=
		runeSearchResults2.TestStringDescription2 {

		return false
	}

	if runeSearchResults1.CollectionTestObjIndex !=
		runeSearchResults2.CollectionTestObjIndex {

		return false
	}

	if runeSearchResults1.TextCharSearchType !=
		runeSearchResults2.TextCharSearchType {

		return false
	}

	if !runeSearchResults1.ReplacementString.Equal(
		&runeSearchResults2.ReplacementString) {

		return false
	}

	if !runeSearchResults1.RemainderString.Equal(
		&runeSearchResults2.RemainderString) {

		return false
	}

	if !runeSearchResults1.FoundCharacters.Equal(
		&runeSearchResults2.FoundCharacters) {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// charSearchResultsDtoAtom.
//
func (searchRunesResultsDtoAtom charSearchRuneArrayResultsDtoAtom) ptr() *charSearchRuneArrayResultsDtoAtom {

	if searchRunesResultsDtoAtom.lock == nil {
		searchRunesResultsDtoAtom.lock = new(sync.Mutex)
	}

	searchRunesResultsDtoAtom.lock.Lock()

	defer searchRunesResultsDtoAtom.lock.Unlock()

	return &charSearchRuneArrayResultsDtoAtom{
		lock: new(sync.Mutex),
	}
}

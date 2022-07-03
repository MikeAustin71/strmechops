package strmech

import "sync"

// charSearchResultsDtoAtom - Provides helper methods for type
// CharSearchResultsDto.
//
type charSearchResultsDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchResultsDto and proceeds to reset the
// data values for member variables to their initial or zero
// values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'searchResultsDto' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto           *CharSearchResultsDto
//     - A pointer to an instance of
//       CharSearchResultsDto. All the internal
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
func (searchResultsDtoAtom *charSearchResultsDtoAtom) empty(
	searchResultsDto *CharSearchResultsDto) {

	if searchResultsDtoAtom.lock == nil {
		searchResultsDtoAtom.lock = new(sync.Mutex)
	}

	searchResultsDtoAtom.lock.Lock()

	defer searchResultsDtoAtom.lock.Unlock()

	if searchResultsDto == nil {
		return
	}

	searchResultsDto.SearchResultsName = ""

	searchResultsDto.SearchResultsFunctionChain = ""

	searchResultsDto.FoundSearchTarget = false

	searchResultsDto.FoundSearchTargetOnPreviousSearch = false

	searchResultsDto.FoundFirstNumericDigitInNumStr = false

	searchResultsDto.TargetInputParametersName = ""

	searchResultsDto.TargetStringLength = -1

	searchResultsDto.TargetStringSearchLength = -1

	searchResultsDto.TargetStringStartingSearchIndex = -1

	searchResultsDto.TargetStringFirstFoundIndex = -1

	searchResultsDto.TargetStringLastFoundIndex = -1

	searchResultsDto.TargetStringLastSearchIndex = -1

	searchResultsDto.TargetStringDescription1 = ""

	searchResultsDto.TargetStringDescription2 = ""

	searchResultsDto.TestInputParametersName = ""

	searchResultsDto.TestStringLength = 0

	searchResultsDto.TestStringStartingIndex = -1

	searchResultsDto.TestStringFirstFoundIndex = -1

	searchResultsDto.TestStringLastFoundIndex = -1

	searchResultsDto.TestStringDescription1 = ""

	searchResultsDto.TestStringDescription2 = ""

	searchResultsDto.CollectionTestObjIndex = -1

	charSearchResultsDtoElectron{}.ptr().
		emptyReplacementStrings(
			searchResultsDto)

	searchResultsDto.NumValueType = NumValType.None()

	searchResultsDto.NumStrFormatType = NumStrFmtType.None()

	searchResultsDto.NumSymLocation = NumSymLocation.None()

	searchResultsDto.NumSymbolClass = NumSymClass.None()

	searchResultsDto.NumSignValue = NumSignVal.None()

	searchResultsDto.PrimaryNumSignPosition =
		NumSignSymPos.None()

	searchResultsDto.SecondaryNumSignPosition =
		NumSignSymPos.None()

	searchResultsDto.TextCharSearchType = CharSearchType.None()

	return
}

// equal - Receives a pointer to two instances of
// CharSearchResultsDto and proceeds to compare their
// member variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto1          *CharSearchResultsDto
//     - An instance of CharSearchResultsDto.
//       Internal member variables from 'searchResultsDto1' will be
//       compared to those of 'searchResultsDto2' to determine if
//       both instances are equivalent.
//
//
//  searchResultsDto2          *CharSearchResultsDto
//     - An instance of CharSearchResultsDto.
//       Internal member variables from 'searchResultsDto2' will be
//       compared to those of 'searchResultsDto1' to determine if
//       both instances are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of 'searchResultsDto1' and
//       'searchResultsDto2' shows that all internal member
//       variables are equivalent, this method will return a
//       boolean value of 'true'.
//
//       If the two instances are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchResultsDtoAtom *charSearchResultsDtoAtom) equal(
	searchResultsDto1 *CharSearchResultsDto,
	searchResultsDto2 *CharSearchResultsDto) bool {

	if searchResultsDtoAtom.lock == nil {
		searchResultsDtoAtom.lock = new(sync.Mutex)
	}

	searchResultsDtoAtom.lock.Lock()

	defer searchResultsDtoAtom.lock.Unlock()

	if searchResultsDto1 == nil ||
		searchResultsDto2 == nil {
		return false
	}

	if searchResultsDto1.SearchResultsName !=
		searchResultsDto2.SearchResultsName {

		return false
	}

	if searchResultsDto1.SearchResultsFunctionChain !=
		searchResultsDto2.SearchResultsFunctionChain {

		return false
	}

	if searchResultsDto1.FoundSearchTarget !=
		searchResultsDto2.FoundSearchTarget {

		return false
	}

	if searchResultsDto1.FoundSearchTargetOnPreviousSearch !=
		searchResultsDto2.FoundSearchTargetOnPreviousSearch {

		return false
	}

	if searchResultsDto1.FoundFirstNumericDigitInNumStr !=
		searchResultsDto2.FoundFirstNumericDigitInNumStr {

		return false
	}

	if searchResultsDto1.TargetInputParametersName !=
		searchResultsDto2.TargetInputParametersName {

		return false
	}

	if searchResultsDto1.TargetStringLength !=
		searchResultsDto2.TargetStringLength {

		return false
	}

	if searchResultsDto1.TargetStringSearchLength !=
		searchResultsDto2.TargetStringSearchLength {

		return false
	}

	if searchResultsDto1.TargetStringStartingSearchIndex !=
		searchResultsDto2.TargetStringStartingSearchIndex {

		return false
	}

	if searchResultsDto1.TargetStringFirstFoundIndex !=
		searchResultsDto2.TargetStringFirstFoundIndex {

		return false
	}

	if searchResultsDto1.TargetStringLastFoundIndex !=
		searchResultsDto2.TargetStringLastFoundIndex {

		return false
	}

	if searchResultsDto1.TargetStringLastSearchIndex !=
		searchResultsDto2.TargetStringLastSearchIndex {

		return false
	}

	if searchResultsDto1.TargetStringDescription1 !=
		searchResultsDto2.TargetStringDescription1 {

		return false
	}

	if searchResultsDto1.TargetStringDescription2 !=
		searchResultsDto2.TargetStringDescription2 {

		return false
	}

	if searchResultsDto1.TestInputParametersName !=
		searchResultsDto2.TestInputParametersName {

		return false
	}

	if searchResultsDto1.TestStringLength !=
		searchResultsDto2.TestStringLength {

		return false
	}

	if searchResultsDto1.TestStringStartingIndex !=
		searchResultsDto2.TestStringStartingIndex {

		return false
	}

	if searchResultsDto1.TestStringFirstFoundIndex !=
		searchResultsDto2.TestStringFirstFoundIndex {

		return false
	}

	if searchResultsDto1.TestStringLastFoundIndex !=
		searchResultsDto2.TestStringLastFoundIndex {

		return false
	}

	if searchResultsDto1.TestStringDescription1 !=
		searchResultsDto2.TestStringDescription1 {

		return false
	}

	if searchResultsDto1.TestStringDescription2 !=
		searchResultsDto2.TestStringDescription2 {

		return false
	}

	if searchResultsDto1.CollectionTestObjIndex !=
		searchResultsDto2.CollectionTestObjIndex {

		return false
	}

	areEqual := charSearchResultsDtoElectron{}.ptr().
		equalReplacementStrings(
			searchResultsDto1,
			searchResultsDto2)

	if !areEqual {
		return false
	}

	if searchResultsDto1.NumValueType !=
		searchResultsDto2.NumValueType {

		return false
	}

	if searchResultsDto1.NumStrFormatType !=
		searchResultsDto2.NumStrFormatType {

		return false
	}

	if searchResultsDto1.NumSymLocation !=
		searchResultsDto2.NumSymLocation {

		return false
	}

	if searchResultsDto1.NumSymbolClass !=
		searchResultsDto2.NumSymbolClass {

		return false
	}

	if searchResultsDto1.NumSignValue !=
		searchResultsDto2.NumSignValue {

		return false
	}

	if searchResultsDto1.PrimaryNumSignPosition !=
		searchResultsDto2.PrimaryNumSignPosition {

		return false
	}

	if searchResultsDto1.SecondaryNumSignPosition !=
		searchResultsDto2.SecondaryNumSignPosition {

		return false
	}

	if searchResultsDto1.TextCharSearchType !=
		searchResultsDto2.TextCharSearchType {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// charSearchResultsDtoAtom.
//
func (searchResultsDtoAtom charSearchResultsDtoAtom) ptr() *charSearchResultsDtoAtom {

	if searchResultsDtoAtom.lock == nil {
		searchResultsDtoAtom.lock = new(sync.Mutex)
	}

	searchResultsDtoAtom.lock.Lock()

	defer searchResultsDtoAtom.lock.Unlock()

	return &charSearchResultsDtoAtom{
		lock: new(sync.Mutex),
	}
}

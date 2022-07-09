package strmech

import "sync"

// charSearchNegativeNumberResultsDtoAtom - Provides helper methods for type
// CharSearchNegativeNumberResultsDto.
//
type charSearchNegativeNumberResultsDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchNegativeNumberResultsDto and proceeds to reset the
// data values for member variables to their initial or zero
// values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'negNumSearchResultsDto' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  negNumSearchResultsDto     *CharSearchNegativeNumberResultsDto
//     - A pointer to an instance of
//       CharSearchNegativeNumberResultsDto. All the internal
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
func (searchNegNumResultsAtom *charSearchNegativeNumberResultsDtoAtom) empty(
	negNumSearchResultsDto *CharSearchNegativeNumberResultsDto) {

	if searchNegNumResultsAtom.lock == nil {
		searchNegNumResultsAtom.lock = new(sync.Mutex)
	}

	searchNegNumResultsAtom.lock.Lock()

	defer searchNegNumResultsAtom.lock.Unlock()

	if negNumSearchResultsDto == nil {
		return
	}

	negNumSearchResultsDto.SearchResultsName = ""

	negNumSearchResultsDto.SearchResultsFunctionChain = ""

	negNumSearchResultsDto.FoundNegativeNumberSymbols = false

	negNumSearchResultsDto.FoundNegNumSymbolsOnPreviousSearch =
		false

	negNumSearchResultsDto.FoundLeadingNegNumSymbols =
		false

	negNumSearchResultsDto.FoundTrailingNegNumSymbols =
		false

	negNumSearchResultsDto.FoundFirstNumericDigitInNumStr = false

	negNumSearchResultsDto.FoundDecimalSeparatorSymbols = false

	negNumSearchResultsDto.FoundNonZeroValue = false

	negNumSearchResultsDto.TargetInputParametersName = ""

	negNumSearchResultsDto.TargetStringLength = -1

	negNumSearchResultsDto.TargetStringSearchLength = -1

	negNumSearchResultsDto.TargetStringAdjustedSearchLength = -1

	negNumSearchResultsDto.TargetStringStartingSearchIndex = -1

	negNumSearchResultsDto.TargetStringCurrentSearchIndex = -1

	negNumSearchResultsDto.TargetStringFirstFoundIndex = -1

	negNumSearchResultsDto.TargetStringLastFoundIndex = -1

	negNumSearchResultsDto.TargetStringLastSearchIndex = -1

	negNumSearchResultsDto.TargetStringNextSearchIndex = -1

	negNumSearchResultsDto.TargetStringDescription1 = ""

	negNumSearchResultsDto.TargetStringDescription2 = ""

	negNumSearchResultsDto.TestInputParametersName = ""

	negNumSearchResultsDto.TestStringName = ""

	negNumSearchResultsDto.TestStringLength = -1

	negNumSearchResultsDto.TestStringLengthName = ""

	negNumSearchResultsDto.TestStringStartingIndex = -1

	negNumSearchResultsDto.TestStringStartingIndexName = ""

	negNumSearchResultsDto.TestStringFirstFoundIndex = -1

	negNumSearchResultsDto.TestStringLastFoundIndex = -1

	negNumSearchResultsDto.TestStringDescription1 = ""

	negNumSearchResultsDto.TestStringDescription2 = ""

	negNumSearchResultsDto.CollectionTestObjIndex = -1

	negNumSearchResultsDto.NumSignValue = NumSignVal.None()

	negNumSearchResultsDto.PrimaryNumSignPosition =
		NumSignSymPos.None()

	negNumSearchResultsDto.SecondaryNumSignPosition =
		NumSignSymPos.None()

	negNumSearchResultsDto.TextCharSearchType =
		CharSearchType.None()

	negNumSearchResultsDto.NegativeNumberSymbolsSpec.Empty()

	negNumSearchResultsDto.ReplacementString.Empty()

	negNumSearchResultsDto.ReplacementString.Empty()

	negNumSearchResultsDto.ReplacementString.Empty()

}

// equal - Receives a pointer to two instances of
// CharSearchNegativeNumberResultsDto and proceeds to compare
// their member variables in order to determine if they are
// equivalent.
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
//  negNumSearchResultsDto1    *CharSearchNegativeNumberResultsDto
//     - An instance of CharSearchNegativeNumberResultsDto.
//       Internal member variables from 'negNumSearchResultsDto1'
//       will be compared to those of 'negNumSearchResultsDto2' to
//       determine if both instances are equivalent.
//
//
//  negNumSearchResultsDto2    *CharSearchNegativeNumberResultsDto
//     - An instance of CharSearchNegativeNumberResultsDto.
//       Internal member variables from 'negNumSearchResultsDto2'
//       will be compared to those of 'negNumSearchResultsDto1' to
//       determine if both instances are equivalent.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of 'negNumSearchResultsDto1' and
//       'negNumSearchResultsDto2' shows that all internal member
//       variables are equivalent, this method will return a
//       boolean value of 'true'.
//
//       If the two instances are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchNegNumResultsAtom *charSearchNegativeNumberResultsDtoAtom) equal(
	negNumSearchResultsDto1 *CharSearchNegativeNumberResultsDto,
	negNumSearchResultsDto2 *CharSearchNegativeNumberResultsDto) bool {

	if searchNegNumResultsAtom.lock == nil {
		searchNegNumResultsAtom.lock = new(sync.Mutex)
	}

	searchNegNumResultsAtom.lock.Lock()

	defer searchNegNumResultsAtom.lock.Unlock()

	if negNumSearchResultsDto1 == nil ||
		negNumSearchResultsDto2 == nil {
		return false
	}

	if negNumSearchResultsDto1.SearchResultsName !=
		negNumSearchResultsDto2.SearchResultsName {

		return false
	}

	if negNumSearchResultsDto1.SearchResultsFunctionChain !=
		negNumSearchResultsDto2.SearchResultsFunctionChain {

		return false
	}

	if negNumSearchResultsDto1.FoundNegativeNumberSymbols !=
		negNumSearchResultsDto2.FoundNegativeNumberSymbols {

		return false
	}

	if negNumSearchResultsDto1.FoundNegNumSymbolsOnPreviousSearch !=
		negNumSearchResultsDto2.FoundNegNumSymbolsOnPreviousSearch {

		return false
	}

	if negNumSearchResultsDto1.FoundLeadingNegNumSymbols !=
		negNumSearchResultsDto2.FoundLeadingNegNumSymbols {

		return false
	}

	if negNumSearchResultsDto1.FoundTrailingNegNumSymbols !=
		negNumSearchResultsDto2.FoundTrailingNegNumSymbols {

		return false
	}

	if negNumSearchResultsDto1.FoundFirstNumericDigitInNumStr !=
		negNumSearchResultsDto2.FoundFirstNumericDigitInNumStr {

		return false
	}

	if negNumSearchResultsDto1.FoundDecimalSeparatorSymbols !=
		negNumSearchResultsDto2.FoundDecimalSeparatorSymbols {

		return false
	}

	if negNumSearchResultsDto1.FoundNonZeroValue !=
		negNumSearchResultsDto2.FoundNonZeroValue {

		return false
	}

	if negNumSearchResultsDto1.TargetInputParametersName !=
		negNumSearchResultsDto2.TargetInputParametersName {

		return false
	}

	if negNumSearchResultsDto1.TargetStringLength !=
		negNumSearchResultsDto2.TargetStringLength {

		return false
	}

	if negNumSearchResultsDto1.TargetStringSearchLength !=
		negNumSearchResultsDto2.TargetStringSearchLength {

		return false
	}

	if negNumSearchResultsDto1.TargetStringAdjustedSearchLength !=
		negNumSearchResultsDto2.TargetStringAdjustedSearchLength {

		return false
	}

	if negNumSearchResultsDto1.TargetStringStartingSearchIndex !=
		negNumSearchResultsDto2.TargetStringStartingSearchIndex {

		return false
	}

	if negNumSearchResultsDto1.TargetStringCurrentSearchIndex !=
		negNumSearchResultsDto2.TargetStringCurrentSearchIndex {

		return false
	}

	if negNumSearchResultsDto1.TargetStringFirstFoundIndex !=
		negNumSearchResultsDto2.TargetStringFirstFoundIndex {

		return false
	}

	if negNumSearchResultsDto1.TargetStringLastFoundIndex !=
		negNumSearchResultsDto2.TargetStringLastFoundIndex {

		return false
	}

	if negNumSearchResultsDto1.TargetStringLastSearchIndex !=
		negNumSearchResultsDto2.TargetStringLastSearchIndex {

		return false
	}

	if negNumSearchResultsDto1.TargetStringNextSearchIndex !=
		negNumSearchResultsDto2.TargetStringNextSearchIndex {

		return false
	}

	if negNumSearchResultsDto1.TargetStringDescription1 !=
		negNumSearchResultsDto2.TargetStringDescription1 {

		return false
	}

	if negNumSearchResultsDto1.TargetStringDescription2 !=
		negNumSearchResultsDto2.TargetStringDescription2 {

		return false
	}

	if negNumSearchResultsDto1.TestInputParametersName !=
		negNumSearchResultsDto2.TestInputParametersName {

		return false
	}

	if negNumSearchResultsDto1.TestStringName !=
		negNumSearchResultsDto2.TestStringName {

		return false
	}

	if negNumSearchResultsDto1.TestStringLength !=
		negNumSearchResultsDto2.TestStringLength {

		return false
	}

	if negNumSearchResultsDto1.TestStringLengthName !=
		negNumSearchResultsDto2.TestStringLengthName {

		return false
	}

	if negNumSearchResultsDto1.TestStringStartingIndex !=
		negNumSearchResultsDto2.TestStringStartingIndex {

		return false
	}

	if negNumSearchResultsDto1.TestStringStartingIndexName !=
		negNumSearchResultsDto2.TestStringStartingIndexName {

		return false
	}

	if negNumSearchResultsDto1.TestStringFirstFoundIndex !=
		negNumSearchResultsDto2.TestStringFirstFoundIndex {

		return false
	}

	if negNumSearchResultsDto1.TestStringLastFoundIndex !=
		negNumSearchResultsDto2.TestStringLastFoundIndex {

		return false
	}

	if negNumSearchResultsDto1.TestStringDescription1 !=
		negNumSearchResultsDto2.TestStringDescription1 {

		return false
	}

	if negNumSearchResultsDto1.TestStringDescription2 !=
		negNumSearchResultsDto2.TestStringDescription2 {

		return false
	}

	if negNumSearchResultsDto1.CollectionTestObjIndex !=
		negNumSearchResultsDto2.CollectionTestObjIndex {

		return false
	}

	if negNumSearchResultsDto1.NumSignValue !=
		negNumSearchResultsDto2.NumSignValue {

		return false
	}

	if negNumSearchResultsDto1.PrimaryNumSignPosition !=
		negNumSearchResultsDto2.PrimaryNumSignPosition {

		return false
	}

	if negNumSearchResultsDto1.SecondaryNumSignPosition !=
		negNumSearchResultsDto2.SecondaryNumSignPosition {

		return false
	}

	if negNumSearchResultsDto1.TextCharSearchType !=
		negNumSearchResultsDto2.TextCharSearchType {

		return false
	}

	if !negNumSearchResultsDto1.NegativeNumberSymbolsSpec.Equal(
		&negNumSearchResultsDto2.NegativeNumberSymbolsSpec) {

		return false
	}

	if !negNumSearchResultsDto1.ReplacementString.Equal(
		&negNumSearchResultsDto2.ReplacementString) {

		return false
	}

	if !negNumSearchResultsDto1.RemainderString.Equal(
		&negNumSearchResultsDto2.RemainderString) {

		return false
	}

	if !negNumSearchResultsDto1.FoundRuneArrayChars.Equal(
		&negNumSearchResultsDto2.FoundRuneArrayChars) {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// charSearchNegativeNumberResultsDtoAtom.
//
func (searchNegNumResultsAtom charSearchNegativeNumberResultsDtoAtom) ptr() *charSearchNegativeNumberResultsDtoAtom {

	if searchNegNumResultsAtom.lock == nil {
		searchNegNumResultsAtom.lock = new(sync.Mutex)
	}

	searchNegNumResultsAtom.lock.Lock()

	defer searchNegNumResultsAtom.lock.Unlock()

	return &charSearchNegativeNumberResultsDtoAtom{
		lock: new(sync.Mutex),
	}
}

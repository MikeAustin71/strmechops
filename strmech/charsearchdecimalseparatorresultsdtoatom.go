package strmech

import "sync"

// charSearchDecimalSeparatorResultsDtoAtom - Provides helper
// methods for type CharSearchDecimalSeparatorResultsDto.
//
type charSearchDecimalSeparatorResultsDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// charSearchDecimalSeparatorResultsDtoAtom and proceeds to reset
// the data values for member variables to their initial or zero
// values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'targetInputParms' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  decimalSepResults          *CharSearchDecimalSeparatorResultsDto
//     - A pointer to an instance of
//       CharSearchDecimalSeparatorResultsDto. All internal member
//       variables contained in this instance will be deleted and
//       reset to their zero values. Array index value will be set
//       minus one (-1). Valid array indexes have values greater
//       than minus one (-1).
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchDecimalSepResultsAtom *charSearchDecimalSeparatorResultsDtoAtom) empty(
	decimalSepResults *CharSearchDecimalSeparatorResultsDto) {

	if searchDecimalSepResultsAtom.lock == nil {
		searchDecimalSepResultsAtom.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsAtom.lock.Lock()

	defer searchDecimalSepResultsAtom.lock.Unlock()

	if decimalSepResults == nil {
		return
	}

	decimalSepResults.SearchResultsName = ""

	decimalSepResults.SearchResultsFunctionChain = ""

	decimalSepResults.IsNOP = false

	decimalSepResults.FoundDecimalSeparatorSymbols = false

	decimalSepResults.FoundDecimalSepSymbolsOnPreviousSearch =
		false

	decimalSepResults.FoundFirstNumericDigitInNumStr = false

	decimalSepResults.FoundNonZeroValue = false

	decimalSepResults.TargetInputParametersName = ""

	decimalSepResults.TargetStringLength = -1

	decimalSepResults.TargetStringSearchLength = -1

	decimalSepResults.TargetStringAdjustedSearchLength = -1

	decimalSepResults.TargetStringStartingSearchIndex = -1

	decimalSepResults.TargetStringCurrentSearchIndex = -1

	decimalSepResults.TargetStringFirstFoundIndex = -1

	decimalSepResults.TargetStringLastFoundIndex = -1

	decimalSepResults.TargetStringLastSearchIndex = -1

	decimalSepResults.TargetStringNextSearchIndex = -1

	decimalSepResults.TargetStringDescription1 = ""

	decimalSepResults.TargetStringDescription2 = ""

	decimalSepResults.TestInputParametersName = ""

	decimalSepResults.TestStringName = ""

	decimalSepResults.TestStringLength = -1

	decimalSepResults.TestStringLengthName = ""

	decimalSepResults.TestStringStartingIndex = -1

	decimalSepResults.TestStringStartingIndexName = ""

	decimalSepResults.TestStringFirstFoundIndex = -1

	decimalSepResults.TestStringLastFoundIndex = -1

	decimalSepResults.TestStringDescription1 = ""

	decimalSepResults.TestStringDescription2 = ""

	decimalSepResults.CollectionTestObjIndex = -1

	decimalSepResults.NumValueType = NumValType.None()

	decimalSepResults.NumSymbolLocation = NumSymLocation.None()

	decimalSepResults.TextCharSearchType =
		CharSearchType.None()

	decimalSepResults.DecimalSeparatorSymbolsSpec.Empty()

	decimalSepResults.ReplacementString.Empty()

	decimalSepResults.RemainderString.Empty()

	decimalSepResults.FoundRuneArrayChars.Empty()

	return
}

// equal - Receives a pointer to two instances of
// CharSearchDecimalSeparatorResultsDto and proceeds to compare their
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
//  decimalSepResults1  *CharSearchDecimalSeparatorResultsDto
//     - A pointer to an instance of
//       CharSearchDecimalSeparatorResultsDto. Internal member
//       variables from 'decimalSepResults1' will be compared to
//       those of 'decimalSepResults2' to determine if both
//       instances are equivalent.
//
//
//  decimalSepResults2  *CharSearchDecimalSeparatorResultsDto
//     - A pointer to an instance of
//       CharSearchDecimalSeparatorResultsDto. Internal member
//       variables from 'decimalSepResults2' will be compared to
//       those of 'decimalSepResults1' to determine if both
//       instances are equivalent.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of 'decimalSepResults1' and
//       'decimalSepResults2' shows that all internal member
//       variables are equivalent, this method will return a
//       boolean value of 'true'.
//
//       If the two instances are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchDecimalSepResultsAtom *charSearchDecimalSeparatorResultsDtoAtom) equal(
	decimalSepResults1 *CharSearchDecimalSeparatorResultsDto,
	decimalSepResults2 *CharSearchDecimalSeparatorResultsDto) bool {

	if searchDecimalSepResultsAtom.lock == nil {
		searchDecimalSepResultsAtom.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsAtom.lock.Lock()

	defer searchDecimalSepResultsAtom.lock.Unlock()

	if decimalSepResults1 == nil ||
		decimalSepResults2 == nil {

		return false
	}

	if decimalSepResults1.SearchResultsName !=
		decimalSepResults2.SearchResultsName {

		return false
	}

	if decimalSepResults1.SearchResultsFunctionChain !=
		decimalSepResults2.SearchResultsFunctionChain {

		return false
	}

	if decimalSepResults1.IsNOP !=
		decimalSepResults2.IsNOP {

		return false
	}

	if decimalSepResults1.FoundDecimalSeparatorSymbols !=
		decimalSepResults2.FoundDecimalSeparatorSymbols {

		return false
	}

	if decimalSepResults1.FoundDecimalSepSymbolsOnPreviousSearch !=
		decimalSepResults2.FoundDecimalSeparatorSymbols {

		return false
	}

	if decimalSepResults1.FoundFirstNumericDigitInNumStr !=
		decimalSepResults2.FoundFirstNumericDigitInNumStr {

		return false
	}

	if decimalSepResults1.FoundNonZeroValue !=
		decimalSepResults2.FoundNonZeroValue {

		return false
	}

	if decimalSepResults1.TargetInputParametersName !=
		decimalSepResults2.TargetInputParametersName {

		return false
	}

	if decimalSepResults1.TargetStringLength !=
		decimalSepResults2.TargetStringLength {

		return false
	}

	if decimalSepResults1.TargetStringSearchLength !=
		decimalSepResults2.TargetStringSearchLength {

		return false
	}

	if decimalSepResults1.TargetStringAdjustedSearchLength !=
		decimalSepResults2.TargetStringAdjustedSearchLength {

		return false
	}

	if decimalSepResults1.TargetStringStartingSearchIndex !=
		decimalSepResults2.TargetStringStartingSearchIndex {

		return false
	}

	if decimalSepResults1.TargetStringCurrentSearchIndex !=
		decimalSepResults2.TargetStringCurrentSearchIndex {

		return false
	}

	if decimalSepResults1.TargetStringFirstFoundIndex !=
		decimalSepResults2.TargetStringFirstFoundIndex {

		return false
	}

	if decimalSepResults1.TargetStringLastFoundIndex !=
		decimalSepResults2.TargetStringLastFoundIndex {

		return false
	}

	if decimalSepResults1.TargetStringLastSearchIndex !=
		decimalSepResults2.TargetStringLastSearchIndex {

		return false
	}

	if decimalSepResults1.TargetStringNextSearchIndex !=
		decimalSepResults2.TargetStringNextSearchIndex {

		return false
	}

	if decimalSepResults1.TargetStringDescription1 !=
		decimalSepResults2.TargetStringDescription1 {

		return false
	}

	if decimalSepResults1.TargetStringDescription2 !=
		decimalSepResults2.TargetStringDescription2 {

		return false
	}

	if decimalSepResults1.TestInputParametersName !=
		decimalSepResults2.TestInputParametersName {

		return false
	}

	if decimalSepResults1.TestStringName !=
		decimalSepResults2.TestStringName {

		return false
	}

	if decimalSepResults1.TestStringLength !=
		decimalSepResults2.TestStringLength {

		return false
	}

	if decimalSepResults1.TestStringLengthName !=
		decimalSepResults2.TestStringLengthName {

		return false
	}

	if decimalSepResults1.TestStringStartingIndex !=
		decimalSepResults2.TestStringStartingIndex {

		return false
	}

	if decimalSepResults1.TestStringStartingIndexName !=
		decimalSepResults2.TestStringStartingIndexName {

		return false
	}

	if decimalSepResults1.TestStringFirstFoundIndex !=
		decimalSepResults2.TestStringFirstFoundIndex {

		return false
	}

	if decimalSepResults1.TestStringLastFoundIndex !=
		decimalSepResults2.TestStringLastFoundIndex {

		return false
	}

	if decimalSepResults1.TestStringDescription1 !=
		decimalSepResults2.TestStringDescription1 {

		return false
	}

	if decimalSepResults1.TestStringDescription2 !=
		decimalSepResults2.TestStringDescription2 {

		return false
	}

	if decimalSepResults1.CollectionTestObjIndex !=
		decimalSepResults2.CollectionTestObjIndex {

		return false
	}

	if decimalSepResults1.NumValueType !=
		decimalSepResults2.NumValueType {

		return false
	}

	if decimalSepResults1.NumSymbolLocation !=
		decimalSepResults2.NumSymbolLocation {

		return false
	}

	if decimalSepResults1.TextCharSearchType !=
		decimalSepResults2.TextCharSearchType {

		return false
	}

	if !decimalSepResults1.DecimalSeparatorSymbolsSpec.Equal(
		&decimalSepResults2.DecimalSeparatorSymbolsSpec) {

		return false
	}

	if !decimalSepResults1.ReplacementString.Equal(
		&decimalSepResults2.ReplacementString) {

		return false
	}

	if !decimalSepResults1.RemainderString.Equal(
		&decimalSepResults2.RemainderString) {

		return false
	}

	if !decimalSepResults1.FoundRuneArrayChars.Equal(
		&decimalSepResults2.FoundRuneArrayChars) {

		return false
	}

	return true
}

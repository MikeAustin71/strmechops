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
func (searchDecimalSepResultsAtom charSearchDecimalSeparatorResultsDtoAtom) empty(
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

	decimalSepResults.FoundDecimalSeparatorSymbols = false

	decimalSepResults.FoundDecimalSepSymbolsOnPreviousSearch =
		false

	decimalSepResults.FoundFirstNumericDigitInNumStr = false

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

}

// ptr - Returns a pointer to a new instance of
// charSearchDecimalSeparatorResultsDtoAtom.
//
func (searchDecimalSepResultsAtom charSearchDecimalSeparatorResultsDtoAtom) ptr() *charSearchDecimalSeparatorResultsDtoAtom {

	if searchDecimalSepResultsAtom.lock == nil {
		searchDecimalSepResultsAtom.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsAtom.lock.Lock()

	defer searchDecimalSepResultsAtom.lock.Unlock()

	return &charSearchDecimalSeparatorResultsDtoAtom{
		lock: new(sync.Mutex),
	}
}

package strmech

import "sync"

// charSearchTestConfigDtoAtom - Provides helper methods for type
// CharSearchTestConfigDto.
//
type charSearchTestConfigDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchTestConfigDto and proceeds to reset the data values
// for member variables to their initial or zero values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'searchTestConfigDto' will be deleted and reset to their zero
// values. Array index values will be set to minus one (-1). Valid
// array indexes have a value greater than minus one (-1).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  searchTestConfigDto        *CharSearchTestConfigDto
//     - A pointer to an instance of CharSearchTestConfigDto. All
//       the internal member variables contained in this instance
//       will be deleted and reset to their zero values. Array
//       index values will be set to minus one (-1). Valid array
//       indexes have a value greater than minus one (-1).
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchTestConfigAtom *charSearchTestConfigDtoAtom) empty(
	searchTestConfigDto *CharSearchTestConfigDto) {

	if searchTestConfigAtom.lock == nil {
		searchTestConfigAtom.lock = new(sync.Mutex)
	}

	searchTestConfigAtom.lock.Lock()

	defer searchTestConfigAtom.lock.Unlock()

	if searchTestConfigDto == nil {
		return
	}

	searchTestConfigDto.TestInputParametersName = ""

	searchTestConfigDto.TestStringName = ""

	searchTestConfigDto.TestStringLengthName = ""

	searchTestConfigDto.TestStringStartingIndex = -1

	searchTestConfigDto.TestStringStartingIndexName = ""

	searchTestConfigDto.TestStringDescription1 = ""

	searchTestConfigDto.TestStringDescription2 = ""

	searchTestConfigDto.CollectionTestObjIndex = -1

	searchTestConfigDto.NumValueType = NumValType.None()

	searchTestConfigDto.NumStrFormatType = NumStrFmtType.None()

	searchTestConfigDto.NumSymbolLocation = NumSymLocation.None()

	searchTestConfigDto.NumSymbolClass = NumSymClass.None()

	searchTestConfigDto.NumSignValue = NumSignVal.None()

	searchTestConfigDto.PrimaryNumSignPosition =
		NumSignSymPos.None()

	searchTestConfigDto.SecondaryNumSignPosition =
		NumSignSymPos.None()

	searchTestConfigDto.TextCharSearchType =
		CharSearchType.None()

	searchTestConfigDto.RequestFoundTestCharacters = false

	searchTestConfigDto.RequestRemainderString = false

	searchTestConfigDto.RequestReplacementString = false

	return
}

// equal - Receives a pointer to two instances of
// CharSearchTestConfigDto and proceeds to compare their
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
//  searchTestCfgDto1          *CharSearchTestConfigDto
//     - An instance of CharSearchTestConfigDto.
//       Internal member variables from 'searchTestCfgDto1' will be
//       compared to those of 'searchTestCfgDto2' to determine if
//       both instances are equivalent.
//
//
//  searchTestCfgDto2          *CharSearchTestConfigDto
//     - An instance of CharSearchTestConfigDto.
//       Internal member variables from 'searchTestCfgDto2' will be
//       compared to those of 'searchTestCfgDto1' to determine if
//       both instances are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of 'searchTestCfgDto1' and
//       'searchTestCfgDto2' shows that all internal member
//       variables are equivalent, this method will return a
//       boolean value of 'true'.
//
//       If the two instances are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchTestConfigAtom *charSearchTestConfigDtoAtom) equal(
	searchTestCfgDto1 *CharSearchTestConfigDto,
	searchTestCfgDto2 *CharSearchTestConfigDto) bool {

	if searchTestConfigAtom.lock == nil {
		searchTestConfigAtom.lock = new(sync.Mutex)
	}

	searchTestConfigAtom.lock.Lock()

	defer searchTestConfigAtom.lock.Unlock()

	if searchTestCfgDto1 == nil ||
		searchTestCfgDto2 == nil {

		return false
	}

	if searchTestCfgDto1.TestInputParametersName !=
		searchTestCfgDto2.TestInputParametersName {

		return false
	}

	if searchTestCfgDto1.TestStringName !=
		searchTestCfgDto2.TestStringName {

		return false
	}

	if searchTestCfgDto1.TestStringLengthName !=
		searchTestCfgDto2.TestStringLengthName {

		return false
	}

	if searchTestCfgDto1.TestStringStartingIndex !=
		searchTestCfgDto2.TestStringStartingIndex {

		return false
	}

	if searchTestCfgDto1.TestStringStartingIndexName !=
		searchTestCfgDto2.TestStringStartingIndexName {

		return false
	}

	if searchTestCfgDto1.TestStringDescription1 !=
		searchTestCfgDto2.TestStringDescription1 {

		return false
	}

	if searchTestCfgDto1.TestStringDescription2 !=
		searchTestCfgDto2.TestStringDescription2 {

		return false
	}

	if searchTestCfgDto1.CollectionTestObjIndex !=
		searchTestCfgDto2.CollectionTestObjIndex {

		return false
	}

	if searchTestCfgDto1.NumValueType !=
		searchTestCfgDto2.NumValueType {

		return false
	}

	if searchTestCfgDto1.NumStrFormatType !=
		searchTestCfgDto2.NumStrFormatType {

		return false
	}

	if searchTestCfgDto1.NumSymbolLocation !=
		searchTestCfgDto2.NumSymbolLocation {

		return false
	}

	if searchTestCfgDto1.NumSymbolClass !=
		searchTestCfgDto2.NumSymbolClass {

		return false
	}

	if searchTestCfgDto1.NumSignValue !=
		searchTestCfgDto2.NumSignValue {

		return false
	}

	if searchTestCfgDto1.PrimaryNumSignPosition !=
		searchTestCfgDto2.PrimaryNumSignPosition {

		return false
	}

	if searchTestCfgDto1.SecondaryNumSignPosition !=
		searchTestCfgDto2.SecondaryNumSignPosition {

		return false
	}

	if searchTestCfgDto1.TextCharSearchType !=
		searchTestCfgDto2.TextCharSearchType {

		return false
	}

	if searchTestCfgDto1.RequestFoundTestCharacters !=
		searchTestCfgDto2.RequestFoundTestCharacters {

		return false
	}

	if searchTestCfgDto1.RequestRemainderString !=
		searchTestCfgDto2.RequestRemainderString {

		return false
	}

	if searchTestCfgDto1.RequestReplacementString !=
		searchTestCfgDto2.RequestReplacementString {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// charSearchTestConfigDtoAtom.
//
func (searchTestConfigAtom charSearchTestConfigDtoAtom) ptr() *charSearchTestConfigDtoAtom {

	if searchTestConfigAtom.lock == nil {
		searchTestConfigAtom.lock = new(sync.Mutex)
	}

	searchTestConfigAtom.lock.Lock()

	defer searchTestConfigAtom.lock.Unlock()

	return &charSearchTestConfigDtoAtom{
		lock: new(sync.Mutex),
	}
}

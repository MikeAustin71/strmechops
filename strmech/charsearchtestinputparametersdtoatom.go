package strmech

import "sync"

// charSearchTestInputParametersDtoAtom - Provides helper methods for type
// CharSearchTestInputParametersDto.
//
type charSearchTestInputParametersDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchTestInputParametersDto and proceeds to reset the
// data values for member variables to their initial or zero
// values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'testInputParms' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  testInputParms           *CharSearchTestInputParametersDto
//     - A pointer to an instance of
//       CharSearchTestInputParametersDto. All the internal
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
func (searchTestInputParmsAtom charSearchTestInputParametersDtoAtom) empty(
	testInputParms *CharSearchTestInputParametersDto) {

	if searchTestInputParmsAtom.lock == nil {
		searchTestInputParmsAtom.lock = new(sync.Mutex)
	}

	searchTestInputParmsAtom.lock.Lock()

	defer searchTestInputParmsAtom.lock.Unlock()

	if testInputParms == nil {
		return
	}

	testInputParms.TestInputParametersName = ""

	charSearchTestInputParametersDtoElectron{}.ptr().
		emptyTestStrings(testInputParms)

	testInputParms.TestStringName = ""

	testInputParms.TestStringLength = -1

	testInputParms.TestStringLengthName = ""

	testInputParms.TestStringStartingIndex = -1

	testInputParms.TestStringStartingIndexName = ""

	testInputParms.TestStringDescription1 = ""

	testInputParms.TestStringDescription2 = ""

	testInputParms.CollectionTestObjIndex = -1

	testInputParms.NumValueType = NumValType.None()

	testInputParms.NumStrFormatType = NumStrFmtType.None()

	testInputParms.NumSymLocation = NumSymLocation.None()

	testInputParms.NumSymbolClass = NumSymClass.None()

	testInputParms.NumSignValue = NumSignVal.None()

	testInputParms.PrimaryNumSignPosition = NumSignSymPos.None()

	testInputParms.SecondaryNumSignPosition = NumSignSymPos.None()

	testInputParms.TextCharSearchType = CharSearchType.None()

	return
}

// equal - Receives a pointer to two instances of
// CharSearchTestInputParametersDto and proceeds to compare their
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
//  testInputParms1          *CharSearchTestInputParametersDto
//     - An instance of CharSearchTestInputParametersDto.
//       Internal member variables from 'testInputParms1' will be
//       compared to those of 'testInputParms2' to determine if
//       both instances are equivalent.
//
//
//  testInputParms2          *CharSearchTestInputParametersDto
//     - An instance of CharSearchTestInputParametersDto.
//       Internal member variables from 'testInputParms2' will be
//       compared to those of 'testInputParms1' to determine if
//       both instances are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of 'testInputParms1' and
//       'testInputParms2' shows that all internal member
//       variables are equivalent, this method will return a
//       boolean value of 'true'.
//
//       If the two instances are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchTestInputParmsAtom charSearchTestInputParametersDtoAtom) equal(
	testInputParms1 *CharSearchTestInputParametersDto,
	testInputParms2 *CharSearchTestInputParametersDto) bool {

	if searchTestInputParmsAtom.lock == nil {
		searchTestInputParmsAtom.lock = new(sync.Mutex)
	}

	searchTestInputParmsAtom.lock.Lock()

	defer searchTestInputParmsAtom.lock.Unlock()

	if testInputParms1 == nil ||
		testInputParms2 == nil {

		return false
	}

	if testInputParms1.TestInputParametersName !=
		testInputParms2.TestInputParametersName {

		return false
	}

	testInputParmsElectron := charSearchTestInputParametersDtoElectron{}

	if !testInputParmsElectron.equalTestStrings(
		testInputParms1, testInputParms2) {

		return false
	}

	if testInputParms1.TestStringName !=
		testInputParms2.TestStringName {

		return false
	}

	if testInputParms1.TestStringLengthName !=
		testInputParms2.TestStringLengthName {

		return false
	}

	if testInputParms1.TestStringStartingIndex !=
		testInputParms2.TestStringStartingIndex {

		return false
	}

	if testInputParms1.TestStringStartingIndexName !=
		testInputParms2.TestStringStartingIndexName {

		return false
	}

	if testInputParms1.TestStringDescription1 !=
		testInputParms2.TestStringDescription1 {

		return false
	}

	if testInputParms1.CollectionTestObjIndex !=
		testInputParms2.CollectionTestObjIndex {

		return false
	}

	if testInputParms1.NumValueType !=
		testInputParms2.NumValueType {

		return false
	}

	if testInputParms1.NumStrFormatType !=
		testInputParms2.NumStrFormatType {

		return false
	}

	if testInputParms1.NumSymLocation !=
		testInputParms2.NumSymLocation {

		return false
	}

	if testInputParms1.NumSymbolClass !=
		testInputParms2.NumSymbolClass {

		return false
	}

	if testInputParms1.NumSignValue !=
		testInputParms2.NumSignValue {

		return false
	}

	if testInputParms1.PrimaryNumSignPosition !=
		testInputParms2.PrimaryNumSignPosition {

		return false
	}

	if testInputParms1.SecondaryNumSignPosition !=
		testInputParms2.SecondaryNumSignPosition {

		return false
	}

	if testInputParms1.TextCharSearchType !=
		testInputParms2.TextCharSearchType {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// charSearchTestInputParametersDtoAtom.
//
func (searchTestInputParmsAtom charSearchTestInputParametersDtoAtom) ptr() *charSearchTestInputParametersDtoAtom {

	if searchTestInputParmsAtom.lock == nil {
		searchTestInputParmsAtom.lock = new(sync.Mutex)
	}

	searchTestInputParmsAtom.lock.Lock()

	defer searchTestInputParmsAtom.lock.Unlock()

	return &charSearchTestInputParametersDtoAtom{
		lock: new(sync.Mutex),
	}
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

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

// testValidityOfTestInputParms - Receives a pointer to an
// instance of CharSearchTestInputParametersDto,
// 'targetInputParms', and performs a diagnostic analysis to
// determine if that instance is valid in all respects.
//
// If the input parameter 'targetInputParms' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'targetInputParms' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
// ----------------------------------------------------------------
//
// Be Advised
//
// In addition to performing validation diagnostics on input
// parameter 'targetInputParms', this method will proceed to set
// all empty member variable labels or name strings to their
// default values.
//
// Type CharSearchTestInputParametersDto contains a number of
// string variables which are used to label, name or otherwise
// describe other operational member variables. If any of these
// label strings are empty when this method is called, those empty
// label strings will be set to their default values.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms           *CharSearchTestInputParametersDto
//     - A pointer to an instance of
//       CharSearchTestInputParametersDto. This object will be
//       subjected to diagnostic analysis in order to determine if
//       all the member variables contain valid values.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
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
//  isValid                    bool
//     - If input parameter 'targetInputParms' is judged to be
//       valid in all respects, this return parameter will be set
//       to 'true'.
//
//     - If input parameter 'targetInputParms' is found to be
//       invalid, this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'targetInputParms' is judged to be
//       valid in all respects, this return parameter will be set
//       to 'nil'.
//
//       If input parameter, 'targetInputParms' is found to be
//       invalid, this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchTestInputParmsAtom charSearchTestInputParametersDtoAtom) testValidityOfTestInputParms(
	testInputParms *CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if searchTestInputParmsAtom.lock == nil {
		searchTestInputParmsAtom.lock = new(sync.Mutex)
	}

	searchTestInputParmsAtom.lock.Lock()

	defer searchTestInputParmsAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoAtom."+
			"testValidityOfTestInputParms()",
		"")

	if err != nil {

		return isValid, err
	}

	if testInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'testInputParms' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if len(testInputParms.TestInputParametersName) == 0 {
		testInputParms.TestInputParametersName =
			"TestInputParameters"
	}

	if len(testInputParms.TestStringName) == 0 {
		testInputParms.TestStringName = "TestString"
	}

	if testInputParms.TestString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			testInputParms.TestStringName)

		return isValid, err
	}

	if len(testInputParms.TestStringLengthName) == 0 {
		testInputParms.TestStringLengthName =
			"TestStringLengthName"
	}

	testInputParms.TestStringLength =
		len(testInputParms.TestString.CharsArray)

	if testInputParms.TestStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			testInputParms.TestStringName,
			testInputParms.TestStringName,
			testInputParms.TestStringName)

		return isValid, err
	}

	if testInputParms.TestStringStartingIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' Starting Index is invalid!\n"+
			"The '%v' Starting Index is less than Zero (0).\n"+
			"%v Starting Index = '%v'.\n",
			ePrefix.String(),
			testInputParms.TestStringName,
			testInputParms.TestStringName,
			testInputParms.TestStringName,
			testInputParms.TestStringStartingIndex)

		return isValid, err

	}

	if len(testInputParms.TestStringStartingIndexName) == 0 {
		testInputParms.TestStringStartingIndexName =
			"TestStringStartingIndex"
	}

	if testInputParms.TestStringStartingIndex >=
		testInputParms.TestStringLength {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The '%v' starting index value is greater than the last index\n"+
			"in the '%v' character array.\n"+
			"%v Last Character Index = '%v'.\n"+
			"%v Starting Index = '%v'\n",
			ePrefix.String(),
			testInputParms.TestStringStartingIndexName,
			testInputParms.TestStringStartingIndexName,
			testInputParms.TestStringName,
			testInputParms.TestStringName,
			testInputParms.TestStringLength-1,
			testInputParms.TestStringStartingIndexName,
			testInputParms.TestStringStartingIndex)

		return isValid, err

	}

	return isValid, err
}

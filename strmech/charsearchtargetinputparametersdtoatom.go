package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// charSearchTargetInputParametersDtoAtom - Provides helper methods for type
// CharSearchTargetInputParametersDto.
//
type charSearchTargetInputParametersDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchTargetInputParametersDto and proceeds to reset the
// data values for member variables to their initial or zero
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
//  targetInputParms           *CharSearchTargetInputParametersDto
//     - A pointer to an instance of
//       CharSearchTargetInputParametersDto. All the internal
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
func (searchTargetInputParmsAtom *charSearchTargetInputParametersDtoAtom) empty(
	targetInputParms *CharSearchTargetInputParametersDto) {

	if searchTargetInputParmsAtom.lock == nil {
		searchTargetInputParmsAtom.lock = new(sync.Mutex)
	}

	searchTargetInputParmsAtom.lock.Lock()

	defer searchTargetInputParmsAtom.lock.Unlock()

	if targetInputParms == nil {
		return
	}

	targetInputParms.TargetInputParametersName = ""

	charSearchTargetInputParametersDtoElectron{}.ptr().
		emptyTargetStrings(targetInputParms)

	targetInputParms.TargetStringName = ""

	targetInputParms.TargetStringLength = -1

	targetInputParms.TargetStringLengthName = ""

	targetInputParms.TargetStringStartingSearchIndex = -1

	targetInputParms.TargetStringCurrentSearchIndex = -1

	targetInputParms.TargetStringNextSearchIndex = -1

	targetInputParms.TargetStringStartingSearchIndexName = ""

	targetInputParms.TargetStringSearchLength = -2

	targetInputParms.TargetStringSearchLengthName = ""

	targetInputParms.TargetStringAdjustedSearchLength = -1

	targetInputParms.TargetStringDescription1 = ""

	targetInputParms.TargetStringDescription2 = ""

	targetInputParms.FoundFirstNumericDigitInNumStr = false

	targetInputParms.FoundDecimalSeparatorSymbols = false

	targetInputParms.FoundNonZeroValue = false

	targetInputParms.TextCharSearchType = CharSearchType.None()

	targetInputParms.RequestFoundTestCharacters = false

	targetInputParms.RequestRemainderString = false

	targetInputParms.RequestReplacementString = false

	return
}

// equal - Receives a pointer to two instances of
// CharSearchTargetInputParametersDto and proceeds to compare their
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
//  targetInputParms1          *CharSearchTargetInputParametersDto
//     - An instance of CharSearchTargetInputParametersDto.
//       Internal member variables from 'targetInputParms1' will be
//       compared to those of 'targetInputParms2' to determine if
//       both instances are equivalent.
//
//
//  targetInputParms2          *CharSearchTargetInputParametersDto
//     - An instance of CharSearchTargetInputParametersDto.
//       Internal member variables from 'targetInputParms2' will be
//       compared to those of 'targetInputParms1' to determine if
//       both instances are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of 'targetInputParms1' and
//       'targetInputParms2' shows that all internal member
//       variables are equivalent, this method will return a
//       boolean value of 'true'.
//
//       If the two instances are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchTargetInputParmsAtom *charSearchTargetInputParametersDtoAtom) equal(
	targetInputParms1 *CharSearchTargetInputParametersDto,
	targetInputParms2 *CharSearchTargetInputParametersDto) bool {

	if searchTargetInputParmsAtom.lock == nil {
		searchTargetInputParmsAtom.lock = new(sync.Mutex)
	}

	searchTargetInputParmsAtom.lock.Lock()

	defer searchTargetInputParmsAtom.lock.Unlock()

	if targetInputParms1 == nil ||
		targetInputParms2 == nil {
		return false
	}

	if targetInputParms1.TargetInputParametersName !=
		targetInputParms2.TargetInputParametersName {

		return false
	}

	targetInputParmsElectron := charSearchTargetInputParametersDtoElectron{}

	if !targetInputParmsElectron.equalTargetStrings(
		targetInputParms1, targetInputParms2) {
		return false
	}

	if targetInputParms1.TargetStringName !=
		targetInputParms2.TargetStringName {

		return false
	}

	if targetInputParms1.TargetStringLength !=
		targetInputParms2.TargetStringLength {

		return false
	}

	if targetInputParms1.TargetStringLengthName !=
		targetInputParms2.TargetStringLengthName {

		return false
	}

	if targetInputParms1.TargetStringStartingSearchIndex !=
		targetInputParms2.TargetStringStartingSearchIndex {

		return false
	}

	if targetInputParms1.TargetStringCurrentSearchIndex !=
		targetInputParms2.TargetStringCurrentSearchIndex {

		return false
	}

	if targetInputParms1.TargetStringNextSearchIndex !=
		targetInputParms2.TargetStringNextSearchIndex {

		return false
	}

	if targetInputParms1.TargetStringStartingSearchIndexName !=
		targetInputParms2.TargetStringStartingSearchIndexName {

		return false
	}

	if targetInputParms1.TargetStringSearchLength !=
		targetInputParms2.TargetStringSearchLength {

		return false
	}

	if targetInputParms1.TargetStringSearchLengthName !=
		targetInputParms2.TargetStringSearchLengthName {

		return false
	}

	if targetInputParms1.TargetStringAdjustedSearchLength !=
		targetInputParms2.TargetStringAdjustedSearchLength {

		return false
	}

	if targetInputParms1.TargetStringDescription1 !=
		targetInputParms2.TargetStringDescription1 {

		return false
	}

	if targetInputParms1.TargetStringDescription2 !=
		targetInputParms2.TargetStringDescription2 {

		return false
	}

	if targetInputParms1.FoundFirstNumericDigitInNumStr !=
		targetInputParms2.FoundFirstNumericDigitInNumStr {

		return false
	}

	if targetInputParms1.FoundDecimalSeparatorSymbols !=
		targetInputParms2.FoundDecimalSeparatorSymbols {

		return false
	}

	if targetInputParms1.FoundNonZeroValue !=
		targetInputParms2.FoundNonZeroValue {

		return false
	}

	if targetInputParms1.TextCharSearchType !=
		targetInputParms2.TextCharSearchType {

		return false
	}

	if targetInputParms1.RequestFoundTestCharacters !=
		targetInputParms2.RequestFoundTestCharacters {

		return false
	}

	if targetInputParms1.RequestRemainderString !=
		targetInputParms2.RequestRemainderString {

		return false
	}

	if targetInputParms1.RequestReplacementString !=
		targetInputParms2.RequestReplacementString {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// charSearchTargetInputParametersDtoAtom.
//
func (searchTargetInputParmsAtom charSearchTargetInputParametersDtoAtom) ptr() *charSearchTargetInputParametersDtoAtom {

	if searchTargetInputParmsAtom.lock == nil {
		searchTargetInputParmsAtom.lock = new(sync.Mutex)
	}

	searchTargetInputParmsAtom.lock.Lock()

	defer searchTargetInputParmsAtom.lock.Unlock()

	return &charSearchTargetInputParametersDtoAtom{
		lock: new(sync.Mutex),
	}
}

// testValidityOfTargetInputParms - Receives a pointer to an
// instance of CharSearchTargetInputParametersDto,
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
// Type CharSearchTargetInputParametersDto contains a number of
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
//  targetInputParms           *CharSearchTargetInputParametersDto
//     - A pointer to an instance of
//       CharSearchTargetInputParametersDto. This object will be
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
func (searchTargetInputParmsAtom *charSearchTargetInputParametersDtoAtom) testValidityOfTargetInputParms(
	targetInputParms *CharSearchTargetInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if searchTargetInputParmsAtom.lock == nil {
		searchTargetInputParmsAtom.lock = new(sync.Mutex)
	}

	searchTargetInputParmsAtom.lock.Lock()

	defer searchTargetInputParmsAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoAtom."+
			"testValidityOfTargetInputParms()",
		"")

	if err != nil {

		return isValid, err
	}

	if targetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if len(targetInputParms.TargetInputParametersName) == 0 {
		targetInputParms.TargetInputParametersName =
			"TargetInputParameters"
	}

	if len(targetInputParms.TargetStringName) == 0 {
		targetInputParms.TargetStringName =
			"TargetString"
	}

	if targetInputParms.TargetString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			targetInputParms.TargetStringName)

		return isValid, err
	}

	if len(targetInputParms.TargetStringLengthName) == 0 {
		targetInputParms.TargetStringLengthName =
			"TargetStringLength"
	}

	targetInputParms.TargetStringLength =
		len(targetInputParms.TargetString.CharsArray)

	if targetInputParms.TargetStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			targetInputParms.TargetStringLengthName,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringName)

		return isValid, err
	}

	if len(targetInputParms.TargetStringStartingSearchIndexName) == 0 {
		targetInputParms.TargetStringStartingSearchIndexName =
			"TargetStringStartingSearchIndex"
	}

	if targetInputParms.TargetStringStartingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter %v is invalid!\n"+
			"%v is less than zero (0)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetInputParms.TargetStringStartingSearchIndexName,
			targetInputParms.TargetStringStartingSearchIndexName,
			targetInputParms.TargetStringStartingSearchIndexName,
			targetInputParms.TargetStringStartingSearchIndex)

		return isValid, err
	}

	if targetInputParms.TargetStringStartingSearchIndex >=
		targetInputParms.TargetStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value greater than the last\n"+
			"index in '%v.CharsArray'.\n"+
			"Last Index in %v.CharsArray = '%v'\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetInputParms.TargetStringStartingSearchIndexName,
			targetInputParms.TargetStringStartingSearchIndexName,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringLength-1,
			targetInputParms.TargetStringStartingSearchIndexName,
			targetInputParms.TargetStringStartingSearchIndex)

		return isValid, err
	}

	if len(targetInputParms.TargetStringSearchLengthName) == 0 {
		targetInputParms.TargetStringSearchLengthName =
			"TargetStringSearchLength"
	}

	if targetInputParms.TargetStringSearchLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value less than minus one (-1)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetInputParms.TargetStringSearchLengthName,
			targetInputParms.TargetStringSearchLengthName,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringSearchLength)

		return isValid, err
	}

	if targetInputParms.TargetStringSearchLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value of Zero (0)\n",
			ePrefix.String(),
			targetInputParms.TargetStringSearchLengthName,
			targetInputParms.TargetStringSearchLengthName)

		return isValid, err
	}

	if targetInputParms.TargetStringSearchLength == -1 {

		targetInputParms.TargetStringAdjustedSearchLength =
			targetInputParms.TargetStringLength -
				targetInputParms.TargetStringStartingSearchIndex
	} else {

		targetInputParms.TargetStringAdjustedSearchLength =
			targetInputParms.TargetStringSearchLength

	}

	targetInputParms.TargetStringAdjustedSearchLength =
		targetInputParms.TargetStringStartingSearchIndex +
			targetInputParms.TargetStringAdjustedSearchLength

	if targetInputParms.TargetStringAdjustedSearchLength >
		targetInputParms.TargetStringLength {

		targetInputParms.TargetStringAdjustedSearchLength =
			targetInputParms.TargetStringLength

	}

	if targetInputParms.TargetStringCurrentSearchIndex <
		targetInputParms.TargetStringStartingSearchIndex {

		targetInputParms.TargetStringCurrentSearchIndex =
			targetInputParms.TargetStringStartingSearchIndex
	}

	if targetInputParms.TargetStringCurrentSearchIndex >=
		targetInputParms.TargetStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%vCurrentSearchIndex has a value greater than\n"+
			"or equal to %vLength\n"+
			"%vLength = '%v'\n"+
			"%vCurrentSearchIndex   = '%v'\n",
			ePrefix.String(),
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringAdjustedSearchLength,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringCurrentSearchIndex)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

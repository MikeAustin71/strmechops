package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// CharSearchTestInputParametersDto - Test Input Parameters are
// more easily understood in the context of text character search
// operations.
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//
// Text Character Search algorithms typically perform comparisons
// between two strings or groups of text characters to determine
// the search outcome. A successful search outcome usually involves
// finding one or more text characters from one string inside a
// second string. A successful search outcome is often referred to
// as a 'Match' condition because characters in one string were
// compared and matched with characters in another string.
//
// Character Search algorithms using the Character Search Type
// ('CharacterSearchType') rely on a framework consisting of a
// 'Target Search String' and a 'Test String'.
//
//    Target String        - A string character or characters which
//                           will be searched for the occurrence of
//                           another predefined character or
//                           characters referred to as a Test
//                           String.
//
//
//    Test String          - A string character or characters which
//                           will be used to search for matching
//                           characters in a Target Search String.
//
// A comparison of text characters contained in the Target Search
// String and the Test String serves as the basis for determining
// a 'Match' condition or successful outcome from a text character
// search algorithm. The specific criterion for determining a
// 'Match' condition vary between the different Character Search
// Types.
//
// ----------------------------------------------------------------
//
// The Character Search Test Input Parameters Data Transfer
// Object type (CharSearchTestInputParametersDto) is used to
// transmit Test String input parameters to methods performing
// search operations.
//
type CharSearchTestInputParametersDto struct {
	TestInputParametersName string
	// The Name, Label or descriptive Tag associated with this
	// instance of CharSearchTestInputParametersDto. If empty,
	// this string will be defaulted to "TestInputParameters"

	TestString *RuneArrayDto
	// A pointer to the Rune Array Data Transfer
	// Object containing the Test Characters to be
	// used in a text character search algorithm.
	// Target Characters are compared against Test
	// Characters to determine if a 'Match' condition
	// exists.

	TestStringName string
	// The label or name of the 'TestString' parameter.
	// Used in error and informational messages.

	TestStringLength int
	// Actual number of text characters in the entire Test
	// String ('TestString').

	TestStringLengthName string
	// The label or name of the 'TestStringLength'
	// parameter. Used in error and informational
	// messages.

	TestStringStartingIndex int
	// The starting index in the Test String where the
	// search operation will begin.

	TestStringStartingIndexName string
	// The label or name of the TestStringStartingIndex
	// parameter. Used in error and informational messages.

	TestStringDescription1 string
	// First of two optional description strings
	// describing the Test String in the context of the
	// current search operation.

	TestStringDescription2 string
	// Second of two optional description strings
	// describing the Test String in the context of the
	// current search operation.

	CollectionTestObjIndex int
	// If the Test String object resides in a collection
	// of Test String objects, this parameter will record
	// the array index of the current Test String object
	// in the collection.

	NumValueType NumericValueType
	// Optional. This enumeration value specifies the type
	// of numeric value for this Test Parameter as either
	// an integer or floating point value.
	//
	// Possible values are listed as follows:
	//  NumValType.None()
	//  NumValType.FloatingPoint()
	//  NumValType.Integer()

	NumStrFormatType NumStrFormatTypeCode
	// Optional. This enumeration value specifies the
	// Output Format Type for a number.
	//
	// Possible values are listed as follows:
	//  NumStrFmtType.None()
	//  NumStrFmtType.AbsoluteValue()
	//  NumStrFmtType.Binary()
	//  NumStrFmtType.CountryCulture()
	//  NumStrFmtType.Currency()
	//  NumStrFmtType.Binary()
	//  NumStrFmtType.Hexadecimal()
	//  NumStrFmtType.Octal()
	//  NumStrFmtType.ScientificNotation()

	NumSymLocation NumericSymbolLocation
	// Optional. This enumeration value specifies the
	// relative location of a numeric symbol.
	//
	// Possible values are listed as follows:
	//  NumSymLocation.None()
	//  NumSymLocation.Before()
	//  NumSymLocation.Interior()
	//  NumSymLocation.After()

	NumSymbolClass NumericSymbolClass
	// Optional. This enumeration value specifies the
	// Number Symbol Classification for a text character.
	//
	// Possible values are listed as follows:
	//  NumSymClass.None()
	//  NumSymClass.NumberSign()
	//  NumSymClass.CurrencySign()
	//  NumSymClass.IntegerSeparator()
	//  NumSymClass.DecimalSeparator()

	NumSignValue NumericSignValueType
	// Optional. This enumeration value specifies the
	// number sign value.
	//
	// Possible values are listed as follows:
	//  NumSignVal.None()
	//  NumSignVal.Negative()
	//  NumSignVal.Zero()
	//  NumSignVal.Positive()

	PrimaryNumSignPosition NumSignSymbolPosition
	// Optional This enumeration value specifies the
	// relative position of positive and negative number
	// sign symbols in a number string. This is the
	// Primary Type Code for Number Signs. Cases involving
	// both 'Leading' and 'Trailing' symbols also make use
	// of the 'SecondaryNumSignPosition'.
	//
	// Possible values are listed as follows:
	//  NumSignSymPos.None()
	//  NumSignSymPos.Before()
	//  NumSignSymPos.After()
	//  NumSignSymPos.BeforeAndAfter()

	SecondaryNumSignPosition NumSignSymbolPosition
	// Optional. This enumeration value specifies the
	// relative position of positive and negative number
	// sign symbols in a number string. This value is used
	// in searches involving number signs which occur both
	// before and after the numeric value.
	//
	// Possible values are listed as follows:
	//  NumSignSymPos.None()
	//  NumSignSymPos.Before()
	//  NumSignSymPos.After()
	//  NumSignSymPos.BeforeAndAfter()

	TextCharSearchType CharacterSearchType
	// Required. An enumeration value signaling the type
	// of text character search algorithm used to conduct
	// this search operation.
	//
	// Possible values are listed as follows:
	//  CharSearchType.None()                      - Invalid Value
	//  CharSearchType.LinearTargetStartingIndex() - Valid Default
	//  CharSearchType.SingleTargetChar()          - Valid
	//  CharSearchType.LinearEndOfString()         - Valid

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// CharSearchTestInputParametersDto ('sourceTestInputParms') to
// the data fields of the current
// CharSearchTestInputParametersDto instance
// ('searchTestInputParmsDto').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current
// CharSearchTestInputParametersDto instance
// ('testSearchInputParms') will be modified and overwritten.
//
// Also, NO data validation is performed on input parameter
// 'sourceTestInputParms'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  sourceTestInputParms     *CharSearchTestInputParametersDto
//     - A pointer to an instance of
//       CharSearchTestInputParametersDto. This method will NOT
//       change the values of internal member variables contained
//       in this instance.
//
//       All data values in this CharSearchTestInputParametersDto
//       instance will be copied to current
//       CharSearchTestInputParametersDto instance
//       ('searchTestInputParmsDto').
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) CopyIn(
	sourceTestInputParms *CharSearchTestInputParametersDto,
	errorPrefix interface{}) error {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTestInputParametersDto."+
			"CopyIn()",
		"")

	if err != nil {

		return err

	}

	testInputParamsNanobot := charSearchTestInputParametersDtoNanobot{}

	return testInputParamsNanobot.
		copyIn(
			testSearchInputParms,
			sourceTestInputParms,
			ePrefix.XCpy(
				"testSearchInputParms"+
					"<-sourceTestInputParms"))
}

// CopyOut - Returns a deep copy of the current
// CharSearchTestInputParametersDto instance.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO data validation is performed on the current instance of
// CharSearchTestInputParametersDto prior to the creation of
// the deep copy returned to the calling function.
//
// It may be necessary to call the validation method
//   CharSearchTestInputParametersDto.IsValidInstanceError()
// before calling this method.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  deepCopyTestInputParms     CharSearchTestInputParametersDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current CharSearchTestInputParametersDto instance.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) CopyOut(
	errorPrefix interface{}) (
	deepCopyTestInputParms CharSearchTestInputParametersDto,
	err error) {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTestInputParametersDto."+
			"CopyOut()",
		"")

	if err != nil {

		return deepCopyTestInputParms, err

	}

	deepCopyTestInputParms,
		err = charSearchTestInputParametersDtoNanobot{}.
		ptr().
		copyOut(
			testSearchInputParms,
			ePrefix.XCpy(
				"deepCopyTestInputParms"+
					"<-testSearchInputParms"))

	return deepCopyTestInputParms, err
}

// Empty - Resets all internal member variables for the current
// instance of CharSearchTestInputParametersDto to their zero or
// uninitialized states. This method will leave the current
// instance of CharSearchTestInputParametersDto in an invalid state
// and unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in the
// current instance of CharSearchTestInputParametersDto. All member
// variable data values will be reset to their zero or
// uninitialized states.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (testSearchInputParms *CharSearchTestInputParametersDto) Empty() {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	charSearchTestInputParametersDtoAtom{}.ptr().
		empty(testSearchInputParms)

	testSearchInputParms.lock.Unlock()

	testSearchInputParms.lock = nil
}

// EmptyTestString - Resets the Empty Test String contained in the
// internal member variable 'TestString' to a nil value thereby
// deleting the previous contents.
//
// The Test String stores the text character or characters used in
// text character search operations. The Test String is configured
// as an internal member variable of type RuneArrayDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (testSearchInputParms *CharSearchTestInputParametersDto) EmptyTestString() {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	charSearchTestInputParametersDtoElectron{}.ptr().
		emptyTestStrings(testSearchInputParms)

	return
}

// Equal - Receives a pointer to another instance of
// CharSearchTestInputParametersDto and proceeds to compare its
// internal member variables to those of the current
// CharSearchTestInputParametersDto instance in order to
// determine if they are equivalent.
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
//  incomingTestInputParms     *CharSearchTestInputParametersDto
//     - A pointer to an instance of
//       CharSearchTestInputParametersDto. The internal member
//       variable data values in this instance will be compared to
//       those in the current instance of
//       CharSearchTestInputParametersDto. The results of this
//       comparison will be returned to the calling functions as a
//       boolean value.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the internal member variable data values contained in
//       input parameter 'incomingTestInputParms' are equivalent
//       in all respects to those contained in the current instance
//       of CharSearchTestInputParametersDto, this return value
//       will be set to 'true'.
//
//       Otherwise, this method will return 'false'.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) Equal(
	incomingTestInputParms *CharSearchTestInputParametersDto) bool {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	return charSearchTestInputParametersDtoAtom{}.ptr().
		equal(
			testSearchInputParms,
			incomingTestInputParms)
}

// EqualTestStrings - Receives a pointer to another instance of
// CharSearchTestInputParametersDto and proceeds to compare the
// internal member variable 'TestString' to the same internal
// member variable 'TestString' contained in the current
// CharSearchTestInputParametersDto instance in order to
// determine if the two 'TestString' member variables are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the 'TestString' member variables for both
// instances are equal in all respects, this flag is set to 'true'.
// Otherwise, this method returns a boolean value of 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTestInputParms   *CharSearchTestInputParametersDto
//     - A pointer to an instance of
//       CharSearchTestInputParametersDto. The internal member
//       variable 'TestString' contained in this instance will be
//       compared to the same member variable 'TestString'
//       contained in the current instance of
//       CharSearchTestInputParametersDto. The results of this
//       comparison will be returned to the calling functions as a
//       boolean value.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the internal member variable 'TestString' contained
//       in input parameter 'incomingTestInputParms' is
//       equivalent in all respects to the same member variable
//       'TestString' contained in the current instance of
//       CharSearchTestInputParametersDto, this return value
//       will be set to 'true'.
//
//       Otherwise, this method will return a boolean value of
//       'false'.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) EqualTestStrings(
	incomingTestInputParms *CharSearchTestInputParametersDto) bool {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	return charSearchTestInputParametersDtoElectron{}.ptr().
		equalTestStrings(
			testSearchInputParms,
			incomingTestInputParms)
}

// GetFormattedText - Returns a formatted text string detailing all
// internal member variables and their values for the current
// instance of CharSearchTestInputParametersDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  string
//     - If this method completes successfully, this string will
//       contain a detailed listing of all internal member
//       variables and their values for the current instance of
//       CharSearchTestInputParametersDto.
//
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) GetFormattedText(
	errorPrefix interface{}) (
	string,
	error) {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTestInputParametersDto."+
			"GetFormattedText()",
		"")

	if err != nil {

		return "", err

	}

	var strBuilder strings.Builder

	strBuilder,
		err = charSearchTestInputParametersDtoNanobot{}.ptr().
		getFormattedText(
			testSearchInputParms,
			ePrefix.XCpy(
				"strBuilder<-Formatted Text"))

	if err != nil {
		return "", err
	}

	return strBuilder.String(), err
}

// IsValidInstance - Performs a diagnostic review of the member
// variable data values encapsulated in the current instance of
// CharSearchTestInputParametersDto to determine if they are
// valid.
//
// If all data element evaluate as valid, this method returns
// 'true'. If any data element is invalid, this method returns
// 'false'.
//
// ----------------------------------------------------------------
//
// Be Advised
//
// In addition to performing validation diagnostics on the current
// instance of CharSearchTestInputParametersDto, this method will
// proceed to set all empty member variable labels or name strings
// to their default values.
//
// Type CharSearchTestInputParametersDto contains a number of
// string variables which are used to label, name or otherwise
// describe other operational member variables. These labels are
// used for error or informational messages. If any of these label
// strings are empty when this method is called, those empty label
// strings will be set to their default values.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isValid             bool
//     - If all data elements encapsulated by the current instance
//       of CharSearchTestInputParametersDto are valid, this
//       returned boolean value is set to 'true'. If any data
//       values are invalid, this return parameter is set to
//       'false'.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) IsValidInstance() (
	isValid bool) {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	isValid,
		_ = charSearchTestInputParametersDtoAtom{}.ptr().
		testValidityOfTestInputParms(
			testSearchInputParms,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current instance of
// CharSearchTestInputParametersDto to determine if they are all
// valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
// This method is functionally equivalent to method:
//  CharSearchTestInputParametersDto.ValidateTestParameters()
//
// This method, IsValidInstanceError(), is included for convenience
// and continuity as many other types in this package also use a
// method of this name when checking the validity of internal member
// variables.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of CharSearchTestInputParametersDto
//       are found to be invalid, this method will return an error.
//       If the member data variables are determined to be valid,
//       this error return parameter will be set to 'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) IsValidInstanceError(
	errorPrefix interface{}) error {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchInputParametersDto."+
			"IsValidInstanceError()",
		"")

	if err != nil {

		return err

	}

	_,
		err = charSearchTestInputParametersDtoAtom{}.ptr().
		testValidityOfTestInputParms(
			testSearchInputParms,
			ePrefix.XCpy(
				"testSearchInputParms"))

	return err
}

// New - Returns a new uninitialized instance of
// CharSearchTestInputParametersDto
//
func (testSearchInputParms CharSearchTestInputParametersDto) New() CharSearchTestInputParametersDto {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	newEmptyTestInputParms := CharSearchTestInputParametersDto{}

	return newEmptyTestInputParms
}

// String - Returns a formatted text string detailing all the
// internal member variables in the current instance of
// CharSearchTestInputParametersDto and their values.
//
// If an error is encountered, the error message is included in the
// string returned by this method.
//
// This method implements the Stringer Interface.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) String() string {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"CharSearchTestInputParametersDto."+
			"String()",
		"")

	if err != nil {
		errOut := fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			"CharSearchTestInputParametersDto.String()",
			err.Error())

		return errOut
	}

	var strBuilder strings.Builder

	strBuilder,
		err = charSearchTestInputParametersDtoNanobot{}.ptr().
		getFormattedText(
			testSearchInputParms,
			ePrefix.XCpy(
				"strBuilder<-Formatted Text"))

	if err != nil {
		return ""
	}

	return strBuilder.String()
}

// ValidateTestParameters - Validates internal member variables
// contained in the current instance of
// CharSearchTestInputParametersDto.
//
// This method is functionally equivalent to method:
//   CharSearchTestInputParametersDto.IsValidInstanceError()
//
//
// ----------------------------------------------------------------
//
// Be Advised
//
// In addition to performing validation diagnostics on the current
// instance of CharSearchTestInputParametersDto, this method will
// proceed to set all empty member variable labels or name strings
// to their default values.
//
// Type CharSearchTestInputParametersDto contains a number of
// string variables which are used to label, name or otherwise
// describe other operational member variables. These labels are
// used for error or informational messages. If any of these label
// strings are empty when this method is called, those empty label
// strings will be set to their default values.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//          containing error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of
//                          ErrPrefixDto. ErrorPrefixInfo from this
//                          object will be copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of CharSearchTestInputParametersDto
//       are found to be invalid, this method will return an error.
//       If the member data variables are determined to be valid,
//       this error return parameter will be set to 'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) ValidateTestParameters(
	errorPrefix interface{}) error {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchInputParametersDto."+
			"ValidateTestParameters()",
		"")

	if err != nil {

		return err

	}

	_,
		err = charSearchTestInputParametersDtoAtom{}.ptr().
		testValidityOfTestInputParms(
			testSearchInputParms,
			ePrefix.XCpy(
				"testSearchInputParms"))

	return err
}

// ValidateCharSearchType - Validates the member variable
// 'TextCharSearchType'. This member variables contains the
// Character Search Type enumeration value which specifies the type
// of text character search algorithm which will be applied in text
// character search operations.
//
// 'TextCharSearchType' is of type CharacterSearchType and is a
// required parameter for all text character search operations.
//
// Possible enumeration values are listed as follows:
//  CharSearchType.None()                      - Invalid Value
//  CharSearchType.LinearTargetStartingIndex() - Valid - Default
//  CharSearchType.SingleTargetChar()          - Valid
//  CharSearchType.LinearEndOfString()         - Valid
//
// For more information, see the documentation for
// type CharacterSearchType.
//
// The validation diagnostics performed by this method will return
// an error if the current value of 'TextCharSearchType' is found
// to be invalid.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the member variable 'TextCharSearchType' contained in
//       the current instance of CharSearchTestInputParametersDto
//       is found to be invalid, this method will return an error.
//
//       If the member data variable 'TextCharSearchType' is found
//       to be valid, this error return parameter will be set to
//       'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
//
func (testSearchInputParms *CharSearchTestInputParametersDto) ValidateCharSearchType(
	errorPrefix interface{}) error {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTestInputParametersDto."+
			"ValidateCharSearchType()",
		"")

	if err != nil {

		return err

	}

	if !testSearchInputParms.TextCharSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: The Character Search Type is invalid!\n"+
			"Character Search Type must be set to one of these\n"+
			"enumeration values:\n"+
			"  CharacterSearchType(0).LinearTargetStartingIndex()\n"+
			"  CharacterSearchType(0).SingleTargetChar()\n"+
			"  CharacterSearchType(0).LinearEndOfString()\n"+
			"The invalid Input Character Search Type is currently\n"+
			"configured as:\n"+
			" Character Search Type   String Name: %v\n"+
			" Character Search Type Integer Value: %v\n",
			ePrefix.String(),
			testSearchInputParms.TextCharSearchType.String(),
			testSearchInputParms.TextCharSearchType.XValueInt())

	}

	return err

}

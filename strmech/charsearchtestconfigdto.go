package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// CharSearchTestConfigDto - This data transfer object transmits
// identification and configuration data to methods creating Test
// Strings and Character Search Test String input parameters.
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
// Type CharSearchTestConfigDto contains data structures used in
// configuring Test Strings. The information contained in this data
// transfer object is usually transferred to a Type
// CharSearchTestInputParametersDto during final configuration of
// Test String input parameters.
//
type CharSearchTestConfigDto struct {
	TestInputParametersName string
	// The Name, Label or descriptive Tag associated with this
	// instance of CharSearchTestInputParametersDto. If empty,
	// this string will be defaulted to "TestInputParameters"

	TestStringName string
	// The label or name of the 'TestString' parameter.
	// Used in error and informational messages.

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

	NumSymbolLocation NumericSymbolLocation
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
// CharSearchTestConfigDto ('sourceSearchTestCfgDto') to the data
// fields of the current CharSearchTestConfigDto instance
// ('searchTestConfigDto').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current CharSearchTestConfigDto instance
// ('searchTestConfigDto') will be modified and overwritten.
//
// NO data validation is performed on input parameter
// 'sourceSearchTestCfgDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  sourceSearchTestCfgDto     *CharSearchTestConfigDto
//     - A pointer to an instance of CharSearchTestConfigDto. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
//
//       All data values in this CharSearchTestConfigDto instance
//       will be copied to current CharSearchTestConfigDto instance
//       ('searchTestConfigDto').
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
// ----------------------------------------------------------------
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
func (searchTestConfigDto *CharSearchTestConfigDto) CopyIn(
	sourceSearchTestCfgDto *CharSearchTestConfigDto,
	errorPrefix interface{}) error {

	if searchTestConfigDto.lock == nil {
		searchTestConfigDto.lock = new(sync.Mutex)
	}

	searchTestConfigDto.lock.Lock()

	defer searchTestConfigDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTestConfigDto."+
			"CopyIn()",
		"")

	if err != nil {

		return err

	}

	return charSearchTestConfigDtoNanobot{}.ptr().
		copyIn(
			searchTestConfigDto,
			sourceSearchTestCfgDto,
			ePrefix.XCpy(
				"searchTestConfigDto<-"+
					"sourceSearchTestCfgDto"))
}

// CopyOut - Returns a deep copy of the current
// CharSearchTestConfigDto instance.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO data validation is performed on the current instance of
// CharSearchTestConfigDto prior to the creation of
// the deep copy returned to the calling function.
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
//  deepCopySearchTestCfgDto   CharSearchTestConfigDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current CharSearchTestConfigDto instance.
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
func (searchTestConfigDto *CharSearchTestConfigDto) CopyOut(
	errorPrefix interface{}) (
	deepCopySearchTestCfgDto CharSearchTestConfigDto,
	err error) {

	if searchTestConfigDto.lock == nil {
		searchTestConfigDto.lock = new(sync.Mutex)
	}

	searchTestConfigDto.lock.Lock()

	defer searchTestConfigDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTestConfigDto."+
			"CopyOut()",
		"")

	if err != nil {

		return deepCopySearchTestCfgDto, err

	}

	deepCopySearchTestCfgDto,
		err = charSearchTestConfigDtoNanobot{}.ptr().
		copyOut(
			searchTestConfigDto,
			ePrefix.XCpy(
				"deepCopySearchTestCfgDto<-"+
					"searchTestConfigDto"))

	return deepCopySearchTestCfgDto, err
}

// New - Returns a new, empty, uninitialized instance of
// CharSearchTestConfigDto. All member variable data elements in
// the returned instance of CharSearchTestConfigDto will be set to
// their zero values. Array index data elements will be set to
// minus one (-1).
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
//  CharSearchTestConfigDto
//     - This method will return a new instance of
//       CharSearchTestConfigDto. All member variable data elements
//       in this new instance will be set to their zero or
//       uninitialized values. Array index data elements will be
//       set to minus one (-1).
//
func (searchTestConfigDto CharSearchTestConfigDto) New() CharSearchTestConfigDto {

	if searchTestConfigDto.lock == nil {
		searchTestConfigDto.lock = new(sync.Mutex)
	}

	searchTestConfigDto.lock.Lock()

	defer searchTestConfigDto.lock.Unlock()

	newTestCfgDto := CharSearchTestConfigDto{}

	charSearchTestConfigDtoAtom{}.ptr().empty(
		&newTestCfgDto)

	return newTestCfgDto
}

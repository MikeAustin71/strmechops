package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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
	//  CharSearchType.None() - Invalid value
	//  CharSearchType.LinearTargetStartingIndex() - Default
	//  CharSearchType.SingleTargetChar()
	//  CharSearchType.LinearEndOfString()

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
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
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
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
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
		"CharSearchInputParametersDto."+
			"ValidateTestString()",
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

func (testSearchInputParms CharSearchTestInputParametersDto) New() CharSearchTestInputParametersDto {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	newEmptyTestInputParms := CharSearchTestInputParametersDto{}

	return newEmptyTestInputParms
}

// ValidateTestParameters - Validates the Test String and related
// member variables contained in the current instance of
// CharSearchInputParametersDto.
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
			"ValidateTestString()",
		"")

	if err != nil {

		return err

	}

	if len(testSearchInputParms.TestInputParametersName) == 0 {
		testSearchInputParms.TestInputParametersName =
			"TestInputParameters"
	}

	if len(testSearchInputParms.TestStringName) == 0 {
		testSearchInputParms.TestStringName = "TestString"
	}

	if testSearchInputParms.TestString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			testSearchInputParms.TestStringName)

		return err
	}

	if len(testSearchInputParms.TestStringLengthName) == 0 {
		testSearchInputParms.TestStringLengthName =
			"TestStringLengthName"
	}

	testSearchInputParms.TestStringLength =
		len(testSearchInputParms.TestString.CharsArray)

	if testSearchInputParms.TestStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName)

		return err
	}

	if testSearchInputParms.TestStringStartingIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' Starting Index is invalid!\n"+
			"The '%v' Starting Index is less than Zero (0).\n"+
			"%v Starting Index = '%v'.\n",
			ePrefix.String(),
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringStartingIndex)

		return err

	}

	if len(testSearchInputParms.TestStringStartingIndexName) == 0 {
		testSearchInputParms.TestStringStartingIndexName =
			"TestStringStartingIndex"
	}

	if testSearchInputParms.TestStringStartingIndex >=
		testSearchInputParms.TestStringLength {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The '%v' starting index value is greater than the last index\n"+
			"in the '%v' character array.\n"+
			"%v Last Character Index = '%v'.\n"+
			"%v Starting Index = '%v'\n",
			ePrefix.String(),
			testSearchInputParms.TestStringStartingIndexName,
			testSearchInputParms.TestStringStartingIndexName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringLength-1,
			testSearchInputParms.TestStringStartingIndexName,
			testSearchInputParms.TestStringStartingIndex)

		return err

	}

	return err
}

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

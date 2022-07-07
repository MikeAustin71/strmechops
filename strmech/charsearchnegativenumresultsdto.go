package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// CharSearchNegativeNumberResultsDto - Contains parameters
// detailing the results of a text character search for negative
// number symbols.
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
// When a 'Match' condition or successful search outcome is
// identified, statistical data describing the 'Match' condition
// is bundled and returned to the calling function.
//
//     Number String        - As used here, a Number String is a
//                            string of text characters which
//                            contain numeric digit characters.
//
// Number String Parsing functions represent a specific type of
// text character search. They are designed to review a string of
// text characters searching for numeric digits. The numeric digits
// are extracted to form numeric values. Number string parsing
// functions therefore convert numeric text characters to valid
// numeric values.
//
// A key feature of Number String Parsing operations is the
// classification of numeric values as positive or negative values.
// This classification logic assumes that converted numeric values
// are positive unless a Negative Number Sign Symbol or Symbols are
// detected within the number string.
//
// ----------------------------------------------------------------
//
// The Character Search Negative Number Results Data Transfer
// Object type (CharSearchNegativeNumberResultsDto) is used to
// identify and transmit the results or outcomes of a search
// operation for negative number symbols in a number string.
//
// Methods performing elements of this text character search
// operation use this wrapper type to encapsulate and return key
// data variables describing the search outcome.
//
type CharSearchNegativeNumberResultsDto struct {
	SearchResultsName string
	// Optional. The Name, Label or descriptive Tag associated with
	// the current instance of CharSearchResultsDto.

	SearchResultsFunctionChain string
	// Optional. A listing of the functions which led to the
	// creation of this instance of CharSearchResultsDto.

	FoundNegativeNumberSymbols bool
	// Signals a successful search outcome. If set to 'true' the
	// Negative Number Symbol character or characters were found
	// in the Target Search String.

	FoundNegNumSymbolsOnPreviousSearch bool
	// Signals that Negative Number Symbols were located in a
	// previous search operation.

	FoundFirstNumericDigitInNumStr bool
	// When set to 'true' this signals that a previous search
	// operation has identified the first numeric digit in a
	// string of text characters.

	TargetInputParametersName string
	// The Name, Label or descriptive Tag associated with an
	// instance of CharSearchTargetInputParametersDto.

	TargetStringLength int
	// Actual number of text characters in the entire
	// Target Search String ('TargetString').

	TargetStringSearchLength int
	// The actual number of characters within the Target
	// Search String that are included in the search
	// operation. This value may be less than the actual
	// length of the Target Search String.

	TargetStringAdjustedSearchLength int
	// The adjusted or corrected Target String Search
	// Length. This value is guaranteed to be equal to or
	// less than the actual Target String Length.

	TargetStringStartingSearchIndex int
	// The index in 'TargetString' at which the search
	// operation begins.

	TargetStringCurrentSearchIndex int
	// The index in 'TargetString' currently being searched.

	TargetStringFirstFoundIndex int
	// The index of the first character position in the
	// Target Search String occupied by the first
	// character in the Test String.

	TargetStringLastFoundIndex int
	// The last character position in the Target Search
	// String occupied by the last character in the Test
	// String.

	TargetStringLastSearchIndex int
	// The index in Target Search String occupied by the
	// last Target character searched. If the Search
	// Target was found, this value is equal to the
	// 'TargetStringLastFoundIndex'. If the Search Target
	// was NOT found this value is equal to the
	// 'TargetStringStartingSearchIndex'. This value is
	// useful in computing the next index to be searched
	// in the Target String.

	TargetStringNextSearchIndex int
	// The starting point for the next search operation.
	// If the entire Target String was included in the
	// last search, this value is set to -1.
	//
	//  Example-1:
	//  String = "Hello"
	//  String Length = 5
	//  Last Search Index = 4
	//  TargetStringNextSearchIndex = -1
	//
	//  Example-2:
	//  String = "Hello"
	//  String Length = 5
	//  Last Search Index = 2
	//  TargetStringNextSearchIndex = 3

	TargetStringDescription1 string
	// First of two optional description strings
	// describing the Target Search String in the context
	// of the current search operation.

	TargetStringDescription2 string
	// Second of two optional description strings
	// describing the Target Search String in the context
	// of the current search operation.

	TestInputParametersName string
	// The Name, Label or descriptive Tag associated with an
	// instance of CharSearchTestInputParametersDto.

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

	TestStringFirstFoundIndex int
	// The index number in Test String of the first test
	// character to be located in the Target Search String.

	TestStringLastFoundIndex int
	// The index number in the Test String occupied by the
	// last Test Character to be located in the Target
	// String.

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

	NegativeNumberSymbolsSpec NegativeNumberSearchSpec
	// If negative number symbols were found in the current search
	// operation, they will be stored in this instance of
	// NegativeNumberSearchSpec

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// CharSearchNegativeNumberResultsDto ('incomingNegNumResultsDto')
// to the data fields of the current
// CharSearchNegativeNumberResultsDto instance
// ('negNumSearchResults').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in the current
// CharSearchNegativeNumberResultsDto instance
// ('negNumSearchResults') will be deleted and overwritten.
//
// No Data Validation will be performed on
// 'incomingNegNumResultsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNegNumResultsDto   *CharSearchNegativeNumberResultsDto
//     - A pointer to an instance of
//       CharSearchNegativeNumberResultsDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this CharSearchNegativeNumberResultsDto
//       instance ('incomingNegNumResultsDto') will be copied to
//       the current CharSearchNegativeNumberResultsDto instance
//       ('negNumSearchResults').
//
//       No Data Validation will be performed on
//       'incomingNegNumResultsDto'.
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
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchResults *CharSearchNegativeNumberResultsDto) CopyIn(
	incomingNegNumResultsDto *CharSearchNegativeNumberResultsDto,
	errorPrefix interface{}) error {

	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	defer negNumSearchResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchNegativeNumberResultsDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return charSearchNegNumResultsDtoNanobot{}.ptr().copyIn(
		negNumSearchResults,
		incomingNegNumResultsDto,
		ePrefix.XCpy(
			"negNumSearchResults<-incomingNegNumResultsDto"))

}

// CopyOut - Returns a deep copy of the current
// CharSearchNegativeNumberResultsDto instance.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// No Data Validation will be performed on the current instance
// of CharSearchResultsDto.
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
//  CharSearchNegativeNumberResultsDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current CharSearchNegativeNumberResultsDto instance.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchResults *CharSearchNegativeNumberResultsDto) CopyOut(
	errorPrefix interface{}) (
	CharSearchNegativeNumberResultsDto,
	error) {

	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	defer negNumSearchResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchNegativeNumberResultsDto."+
			"CopyOut()",
		"")

	if err != nil {
		return CharSearchNegativeNumberResultsDto{}, err
	}

	return charSearchNegNumResultsDtoNanobot{}.ptr().copyOut(
		negNumSearchResults,
		ePrefix.XCpy(
			"<-negNumSearchResults"))
}

// Empty - Resets all internal member variables for the current
// instance of CharSearchNegativeNumberResultsDto to their zero or
// uninitialized states. This method will leave the current
// instance of CharSearchNegativeNumberResultsDto in an
// invalid state and unavailable for immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in the
// current instance of CharSearchNegativeNumberResultsDto. All
// member variable data values will be reset to their zero or
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
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (negNumSearchResults *CharSearchNegativeNumberResultsDto) Empty() {
	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	charSearchNegativeNumberResultsDtoAtom{}.ptr().
		empty(negNumSearchResults)

	negNumSearchResults.lock.Unlock()

	negNumSearchResults.lock = nil
}

// Equal - Receives a pointer to another instance of
// CharSearchNegativeNumberResultsDto and proceeds to compare the member
// variables to those of the current CharSearchNegativeNumberResultsDto
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNegNumSearchResultsDto  *CharSearchNegativeNumberResultsDto
//     - A pointer to an incoming instance of
//       CharSearchNegativeNumberResultsDto. This method will
//       compare all member variable data values in this instance
//       against those contained in the current instance of
//       CharSearchNegativeNumberResultsDto. If the data values in
//       both instances are found to be equal in all respects, this
//       method will return a boolean value of 'true'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the member variable data values contained in input
//       parameter 'incomingNegNumSearchResultsDto' are equal in
//       all respects to those contained in the current instance of
//       CharSearchNegativeNumberResultsDto, this method will
//       return a boolean value of 'true'. Otherwise, a value of
//       'false' will be returned to the calling function.
//
func (negNumSearchResults *CharSearchNegativeNumberResultsDto) Equal(
	incomingNegNumSearchResultsDto *CharSearchNegativeNumberResultsDto) bool {

	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	defer negNumSearchResults.lock.Unlock()

	return charSearchNegativeNumberResultsDtoAtom{}.ptr().
		equal(
			negNumSearchResults,
			incomingNegNumSearchResultsDto)
}

// GetParameterTextListing - Returns formatted text output
// detailing the member variable names and their corresponding
// values contained in the current instance of
// CharSearchNegativeNumberResultsDto ('negNumSearchResults').
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
//  strings.Builder
//     - If this method completes successfully, an instance of
//       strings.Builder will be returned. This instance contains
//       the formatted text output listing the member variable
//       names and their corresponding values for the current
//       instance of CharSearchNegativeNumberResultsDto. This
//       formatted text can then be used for text displays, file
//       output or printing.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (negNumSearchResults *CharSearchNegativeNumberResultsDto) GetParameterTextListing(
	errorPrefix interface{}) (
	strings.Builder,
	error) {

	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	defer negNumSearchResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchNegativeNumberResultsDto."+
			"GetParameterTextListing()",
		"")

	if err != nil {
		return strings.Builder{}, err
	}

	return charSearchNegNumResultsDtoNanobot{}.ptr().
		getParameterTextListing(
			negNumSearchResults,
			ePrefix.XCpy(
				"negNumSearchResults->Parameter Listing"))

}

// LoadTargetBaseInputParameters - Receives Target String data from
// input parameter 'targetInputParms' and proceeds to transfer key
// data for the search operation to the current instance of
// CharSearchNegativeNumberResultsDto.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method assumes that the input data elements contained in
// 'testInputParms' have been validated.
//
func (negNumSearchResults *CharSearchNegativeNumberResultsDto) LoadTargetBaseInputParameters(
	targetInputParms CharSearchTargetInputParametersDto) {

	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	defer negNumSearchResults.lock.Unlock()

	negNumSearchResults.TargetInputParametersName =
		targetInputParms.TargetInputParametersName

	negNumSearchResults.FoundFirstNumericDigitInNumStr =
		targetInputParms.FoundFirstNumericDigitInNumStr

	negNumSearchResults.TargetStringLength =
		targetInputParms.TargetStringLength

	negNumSearchResults.TargetStringSearchLength =
		targetInputParms.TargetStringSearchLength

	negNumSearchResults.TargetStringAdjustedSearchLength =
		targetInputParms.TargetStringAdjustedSearchLength

	negNumSearchResults.TargetStringStartingSearchIndex =
		targetInputParms.TargetStringStartingSearchIndex

	negNumSearchResults.TargetStringCurrentSearchIndex =
		targetInputParms.TargetStringCurrentSearchIndex

	negNumSearchResults.TargetStringDescription1 =
		targetInputParms.TargetStringDescription1

	negNumSearchResults.TargetStringDescription2 =
		targetInputParms.TargetStringDescription2

}

// LoadTestBaseInputParameters - Receives Target String data from
// input parameter 'testInputParms' and proceeds to transfer key
// data for the search operation to the current instance of
// CharSearchNegativeNumberResultsDto.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method assumes that the input data elements contained in
// 'testInputParms' have been validated.
//
func (negNumSearchResults *CharSearchNegativeNumberResultsDto) LoadTestBaseInputParameters(
	testInputParms CharSearchTestInputParametersDto) {

	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	defer negNumSearchResults.lock.Unlock()

	negNumSearchResults.TestInputParametersName =
		testInputParms.TestInputParametersName

	negNumSearchResults.TestStringName =
		testInputParms.TestStringName

	negNumSearchResults.TestStringLength =
		testInputParms.TestStringLength

	negNumSearchResults.TestStringLengthName =
		testInputParms.TestStringLengthName

	negNumSearchResults.TestStringStartingIndex =
		testInputParms.TestStringStartingIndex

	negNumSearchResults.TestStringStartingIndexName =
		testInputParms.TestStringStartingIndexName

	negNumSearchResults.TestStringDescription1 =
		testInputParms.TestStringDescription1

	negNumSearchResults.TestStringDescription2 =
		testInputParms.TestStringDescription2

	negNumSearchResults.CollectionTestObjIndex =
		testInputParms.CollectionTestObjIndex

	negNumSearchResults.NumSignValue =
		testInputParms.NumSignValue

	negNumSearchResults.PrimaryNumSignPosition =
		testInputParms.PrimaryNumSignPosition

	negNumSearchResults.SecondaryNumSignPosition =
		testInputParms.SecondaryNumSignPosition

	negNumSearchResults.TextCharSearchType =
		testInputParms.TextCharSearchType

}

// New - Returns a new and uninitialized instance of
// CharSearchNegativeNumberResultsDto.
//
// All member variables in this returned instance are set to their
// zero or uninitialized states. Array index values are set to a
// value of minus one (-1). All valid array indexes have values
// greater than minus one (-1).
//
func (negNumSearchResults CharSearchNegativeNumberResultsDto) New() CharSearchNegativeNumberResultsDto {

	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	defer negNumSearchResults.lock.Unlock()

	newNegNumResultsDto := CharSearchNegativeNumberResultsDto{}

	newNegNumResultsDto.Empty()

	return newNegNumResultsDto
}

// String - Returns a formatted text string detailing all the
// internal member variable names and their corresponding values
// for the current instance of CharSearchNegativeNumberResultsDto.
//
// If an error is encountered, the error message is included in the
// string returned by this method.
//
// This method implements the Stringer Interface.
//
func (negNumSearchResults *CharSearchNegativeNumberResultsDto) String() string {

	if negNumSearchResults.lock == nil {
		negNumSearchResults.lock = new(sync.Mutex)
	}

	negNumSearchResults.lock.Lock()

	defer negNumSearchResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"CharSearchNegativeNumberResultsDto."+
			"String()",
		"")

	if err != nil {
		errOut := fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			"CharSearchNegativeNumberResultsDto.String()",
			err.Error())

		return errOut
	}

	var strBuilder strings.Builder

	strBuilder,
		err = charSearchNegNumResultsDtoNanobot{}.ptr().
		getParameterTextListing(
			negNumSearchResults,
			ePrefix.XCpy(
				"negNumSearchResults->Parameter Listing"))

	if err != nil {
		errOut := fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			ePrefix.String(),
			err.Error())

		return errOut
	}

	return strBuilder.String()
}

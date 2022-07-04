package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// CharSearchResultsDto - Text character search results are more
// easily understood in the context of text character search
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
// When a 'Match' condition or successful search outcome is
// identified, statistical data describing the 'Match' condition
// is bundled and returned to the calling function.
//
// ----------------------------------------------------------------
//
// The Character Search Results Data Transfer Object type
// (CharSearchResultsDto) is used to identify and transmit the
// results or outcome of a search operation.
//
// Methods performing elements of the text character search
// operation use this wrapper type to encapsulate and return key
// data variables describing the search outcome.
//
type CharSearchResultsDto struct {
	SearchResultsName string
	// Optional. The Name, Label or descriptive Tag associated with
	// the current instance of CharSearchResultsDto.

	SearchResultsFunctionChain string
	// Optional. A listing of the functions which led to the
	// creation of this instance of CharSearchResultsDto.

	FoundSearchTarget bool
	// Signals a successful search outcome. If set to 'true' the
	// Test String character or characters were found in the Target
	// Search String.

	FoundSearchTargetOnPreviousSearch bool
	// Signals that the Search Target was located in a previous
	// search operation.

	FoundFirstNumericDigitInNumStr bool
	// When set to 'true' this signals that the search operation
	// has identified the first numeric digit in a string of text
	// characters.

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

	TestStringLength int
	// Actual number of text characters in the entire Test
	// String ('TestString').

	TestStringStartingIndex int
	// The starting index in the Test String where the
	// search operation will begin.

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

	ReplacementString *RuneArrayDto
	// A pointer to the Rune Array Data Transfer Object
	// containing the Replacement Characters to be
	// substituted for existing characters in a Target
	// String.

	RemainderString *RuneArrayDto
	// A pointer to the Rune Array Data Transfer Object
	// containing the remaining characters in a Target
	// String which were NOT included in the search
	// operation and which remain to be searched in future
	// search operations. This string is also used in 'cut'
	// operations where Target String is divided based on
	// string delimiters.

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
// CharSearchResultsDto ('incomingSearchResults') to the data
// fields of the current CharSearchResultsDto instance
// ('charSearchResults').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current CharSearchResultsDto instance
// ('charSearchResults') will be deleted and overwritten.
//
// No Data Validation will be performed on 'incomingSearchResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingSearchResults     *CharSearchResultsDto
//     - A pointer to an instance of CharSearchResultsDto. This
//       method will NOT change the data values of member variables
//       contained in this instance.
//
//       All data values in this CharSearchResultsDto instance
//       ('incomingSearchResults') will be copied to the current
//       CharSearchResultsDto instance ('charSearchResults').
//
//       No Data Validation will be performed on
//       'incomingSearchResults'.
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
func (charSearchResults *CharSearchResultsDto) CopyIn(
	incomingSearchResults *CharSearchResultsDto,
	errorPrefix interface{}) error {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	defer charSearchResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchResultsDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return charSearchResultsDtoNanobot{}.ptr().
		copyIn(
			charSearchResults,
			incomingSearchResults,
			ePrefix.XCpy(
				"charSearchResults<-"+
					"incomingSearchResults"))
}

// CopyOut - Returns a deep copy of the current
// CharSearchResultsDto instance.
//
// If the current CharSearchResultsDto instance contains invalid
// member variables, this method will return an error.
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
//  CharSearchResultsDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current CharSearchResultsDto instance.
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
func (charSearchResults *CharSearchResultsDto) CopyOut(
	errorPrefix interface{}) (
	CharSearchResultsDto,
	error) {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	defer charSearchResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchResultsDto."+
			"CopyOut()",
		"")

	if err != nil {
		return CharSearchResultsDto{}, err
	}

	return charSearchResultsDtoNanobot{}.ptr().
		copyOut(
			charSearchResults,
			ePrefix.XCpy(
				"<-CharSearchResultsDto"))

}

// Empty - Resets all internal member variables for the current
// instance of CharSearchResultsDto to their zero or uninitialized
// states. This method will leave the current instance of
// CharSearchResultsDto in an invalid state and unavailable for
// immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of CharSearchResultsDto. All member variable
// data values will be reset to their zero or uninitialized states.
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
func (charSearchResults *CharSearchResultsDto) Empty() {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	charSearchResultsDtoAtom{}.ptr().empty(
		charSearchResults)

	charSearchResults.lock.Unlock()

	charSearchResults.lock = nil
}

// EmptyRemainderString - Resets the internal member variable
// 'RemainderString' for the current instance of
// CharSearchResultsDto to its zero or uninitialized state.
//
// Only the internal member variable
// 'CharSearchResultsDto.RemainderString' is deleted and reset to a
// value of 'nil'.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method deletes the member variable data value contained in
// 'CharSearchResultsDto.RemainderString' and resets this pointer
// to 'nil'.
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
func (charSearchResults *CharSearchResultsDto) EmptyRemainderString() {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	charSearchResultsDtoAtom{}.ptr().empty(
		charSearchResults)

	charSearchResults.lock.Unlock()

	charSearchResults.lock = nil

	charSearchResultsDtoElectron{}.ptr().
		emptyRemainderStrings(
			charSearchResults)
}

// EmptyReplacementString - Resets the internal member variable
// 'ReplacementString' for the current instance of
// CharSearchResultsDto to its zero or uninitialized state.
//
// Only the internal member variable
// CharSearchResultsDto.ReplacementString is deleted and reset to a
// value of 'nil'.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method deletes the member variable data value contained in
// CharSearchResultsDto.ReplacementString and resets this pointer
// to 'nil'.
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
func (charSearchResults *CharSearchResultsDto) EmptyReplacementString() {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	charSearchResultsDtoAtom{}.ptr().empty(
		charSearchResults)

	charSearchResults.lock.Unlock()

	charSearchResults.lock = nil

	charSearchResultsDtoElectron{}.ptr().
		emptyReplacementStrings(
			charSearchResults)
}

// Equal - Receives a pointer to another instance of
// CharSearchResultsDto and proceeds to compare the member
// variables to those of the current CharSearchResultsDto
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
//  incomingSearchResults    *CharSearchResultsDto
//     - A pointer to an incoming instance of
//       CharSearchResultsDto. This method will compare all member
//       variable data values in this instance against those
//       contained in the current instance of CharSearchResultsDto.
//       If the data values in both instances are found to be equal
//       in all respects, this method will return a boolean value
//       of 'true'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the member variable data values contained in input
//       parameter 'incomingSearchResults' are equal in all
//       respects to those contained in the current instance of
//       CharSearchResultsDto, this method will return a boolean
//       value of 'true'. Otherwise a value of 'false' will be
//       returned to the calling function.
//
func (charSearchResults *CharSearchResultsDto) Equal(
	incomingSearchResults *CharSearchResultsDto) bool {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	defer charSearchResults.lock.Unlock()

	return charSearchResultsDtoAtom{}.ptr().equal(
		charSearchResults,
		incomingSearchResults)
}

// New - Returns a new and uninitialized instance of
// CharSearchResultsDto
//
func (charSearchResults CharSearchResultsDto) New() CharSearchResultsDto {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	defer charSearchResults.lock.Unlock()

	newEmptySearchResults := CharSearchResultsDto{}

	newEmptySearchResults.Empty()

	return newEmptySearchResults
}

// LoadTargetBaseInputParameters - Receives Target String data from
// input parameter 'targetInputParms' and proceeds to transfer key
// data for the search operation to the current instance of
// CharSearchResultsDto.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method assumes that the input data elements contained in
// 'testInputParms' have been validated.
//
func (charSearchResults *CharSearchResultsDto) LoadTargetBaseInputParameters(
	targetInputParms CharSearchTargetInputParametersDto) {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	defer charSearchResults.lock.Unlock()

	charSearchResults.TargetInputParametersName =
		targetInputParms.TargetInputParametersName

	charSearchResults.FoundFirstNumericDigitInNumStr =
		targetInputParms.FoundFirstNumericDigitInNumStr

	charSearchResults.TargetStringLength =
		targetInputParms.TargetStringLength

	charSearchResults.TargetStringSearchLength =
		targetInputParms.TargetStringSearchLength

	charSearchResults.TargetStringAdjustedSearchLength =
		targetInputParms.TargetStringAdjustedSearchLength

	charSearchResults.TargetStringStartingSearchIndex =
		targetInputParms.TargetStringStartingSearchIndex

	charSearchResults.TargetStringDescription1 =
		targetInputParms.TargetStringDescription1

	charSearchResults.TargetStringDescription2 =
		targetInputParms.TargetStringDescription2

}

// LoadTestBaseInputParameters - Receives Target String data from
// input parameter 'testInputParms' and proceeds to transfer key
// data for the search operation to the current instance of
// CharSearchResultsDto.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method assumes that the input data elements contained in
// 'testInputParms' have been validated.
//
func (charSearchResults *CharSearchResultsDto) LoadTestBaseInputParameters(
	testInputParms CharSearchTestInputParametersDto) {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	defer charSearchResults.lock.Unlock()

	charSearchResults.TestInputParametersName =
		testInputParms.TestInputParametersName

	charSearchResults.TestStringLength =
		testInputParms.TestStringLength

	charSearchResults.TestStringStartingIndex =
		testInputParms.TestStringStartingIndex

	charSearchResults.TestStringDescription1 =
		testInputParms.TestStringDescription1

	charSearchResults.TestStringDescription2 =
		testInputParms.TestStringDescription2

	charSearchResults.CollectionTestObjIndex =
		testInputParms.CollectionTestObjIndex

	charSearchResults.NumValueType =
		testInputParms.NumValueType

	charSearchResults.NumStrFormatType =
		testInputParms.NumStrFormatType

	charSearchResults.NumSymLocation =
		testInputParms.NumSymLocation

	charSearchResults.NumSymbolClass =
		testInputParms.NumSymbolClass

	charSearchResults.NumSignValue =
		testInputParms.NumSignValue

	charSearchResults.PrimaryNumSignPosition =
		testInputParms.PrimaryNumSignPosition

	charSearchResults.SecondaryNumSignPosition =
		testInputParms.SecondaryNumSignPosition

	charSearchResults.TextCharSearchType =
		testInputParms.TextCharSearchType

}

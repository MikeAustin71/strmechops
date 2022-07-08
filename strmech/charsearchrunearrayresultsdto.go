package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type CharSearchRuneArrayResultsDto struct {
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

	FoundFirstNumericDigitInNumStr bool
	// When set to 'true' this signals that the search operation
	// has identified the first numeric digit in a string of text
	// characters.

	FoundDecimalSeparatorSymbols bool
	// When set to 'true' this signals that a Decimal
	// Separator Symbol character or characters have been
	// identified in the text characters specified by
	// 'TargetString'

	FoundNonZeroValue bool
	// When set to 'true' this signals that the search operation
	// has detected a nonzero numeric digit.

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

	ReplacementString RuneArrayDto
	// A Rune Array Data Transfer Object containing the
	// Replacement Characters to be substituted for
	// existing characters in a Target String.

	RemainderString RuneArrayDto
	// A Rune Array Data Transfer Object containing the
	// remaining text characters at the end of the Target
	// String which were NOT included in the most recent
	// search operation.

	FoundCharacters RuneArrayDto
	// A Rune Array Data Transfer Object containing the
	// text characters located in the Target String
	// by the most recent search operation.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// CharSearchRuneArrayResultsDto ('incomingRuneSearchResults')
// to the data fields of the current CharSearchRuneArrayResultsDto
// instance ('runesSearchResultsDto').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in the current CharSearchRuneArrayResultsDto
// instance ('runesSearchResultsDto') will be deleted and
// overwritten.
//
// No Data Validation will be performed on
// input parameter 'incomingRuneSearchResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingRuneSearchResults   *CharSearchRuneArrayResultsDto
//     - A pointer to an instance of
//       CharSearchRuneArrayResultsDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this CharSearchRuneArrayResultsDto
//       instance ('incomingRuneSearchResults') will be copied to
//       the current CharSearchRuneArrayResultsDto instance
//       ('runesSearchResultsDto').
//
//       No Data Validation will be performed on
//       'incomingRuneSearchResults'.
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
func (runesSearchResultsDto *CharSearchRuneArrayResultsDto) CopyIn(
	incomingRuneSearchResults *CharSearchRuneArrayResultsDto,
	errorPrefix interface{}) error {

	if runesSearchResultsDto.lock == nil {
		runesSearchResultsDto.lock = new(sync.Mutex)
	}

	runesSearchResultsDto.lock.Lock()

	defer runesSearchResultsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchRuneArrayResultsDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return charSearchRuneArrayResultsDtoNanobot{}.ptr().copyIn(
		runesSearchResultsDto,
		incomingRuneSearchResults,
		ePrefix.XCpy(
			"runesSearchResultsDto<-incomingRuneSearchResults"))
}

// CopyOut - Returns a deep copy of the current
// CharSearchRuneArrayResultsDto instance.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// No Data Validation will be performed on the current instance
// of CharSearchRuneArrayResultsDto.
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
//  CharSearchRuneArrayResultsDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current CharSearchRuneArrayResultsDto instance.
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
func (runesSearchResultsDto *CharSearchRuneArrayResultsDto) CopyOut(
	errorPrefix interface{}) (
	CharSearchRuneArrayResultsDto,
	error) {

	if runesSearchResultsDto.lock == nil {
		runesSearchResultsDto.lock = new(sync.Mutex)
	}

	runesSearchResultsDto.lock.Lock()

	defer runesSearchResultsDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchRuneArrayResultsDto."+
			"CopyIn()",
		"")

	if err != nil {
		return CharSearchRuneArrayResultsDto{}, err
	}

	return charSearchRuneArrayResultsDtoNanobot{}.ptr().copyOut(
		runesSearchResultsDto,
		ePrefix.XCpy(
			"<-runesSearchResultsDto"))

}

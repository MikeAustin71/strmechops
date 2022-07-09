package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
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

// Empty - Resets all internal member variables for the current
// instance of CharSearchRuneArrayResultsDto to their zero or
// uninitialized states. This method will leave the current
// instance of CharSearchRuneArrayResultsDto in an
// invalid state and unavailable for immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in the
// current instance of CharSearchRuneArrayResultsDto. All member
// variable data values will be reset to their zero or
// uninitialized states. Array Index values will be set to minus
// one (-1). All valid Array Index values are greater than minus
// one (-1).
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
func (runesSearchResultsDto *CharSearchRuneArrayResultsDto) Empty() {

	if runesSearchResultsDto.lock == nil {
		runesSearchResultsDto.lock = new(sync.Mutex)
	}

	runesSearchResultsDto.lock.Lock()

	charSearchRuneArrayResultsDtoAtom{}.ptr().
		empty(runesSearchResultsDto)

	runesSearchResultsDto.lock.Unlock()

	runesSearchResultsDto.lock = nil

	return
}

// Equal - Receives a pointer to another instance of
// CharSearchRuneArrayResultsDto and proceeds to compare the member
// variables to those of the current CharSearchRuneArrayResultsDto
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
//  incomingRuneSearchResults  *CharSearchRuneArrayResultsDto
//     - A pointer to an incoming instance of
//       CharSearchRuneArrayResultsDto. This method will compare
//       all member variable data values in this instance against
//       those contained in the current instance of
//       CharSearchRuneArrayResultsDto. If the data values in both
//       instances are found to be equal in all respects, this
//       method will return a boolean value of 'true'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the member variable data values contained in input
//       parameter 'incomingRuneSearchResults' are equal in all
//       respects to those contained in the current instance of
//       CharSearchRuneArrayResultsDto, this method will return a
//       boolean value of 'true'. Otherwise, a value of 'false'
//       will be returned to the calling function.
//
func (runesSearchResultsDto *CharSearchRuneArrayResultsDto) Equal(
	incomingRuneSearchResults *CharSearchRuneArrayResultsDto) bool {

	if runesSearchResultsDto.lock == nil {
		runesSearchResultsDto.lock = new(sync.Mutex)
	}

	runesSearchResultsDto.lock.Lock()

	defer runesSearchResultsDto.lock.Unlock()

	return charSearchRuneArrayResultsDtoAtom{}.ptr().
		equal(runesSearchResultsDto,
			incomingRuneSearchResults)
}

// GetParameterTextListing - Returns formatted text output
// detailing the member variable names and their corresponding
// values contained in the current instance of
// CharSearchRuneArrayResultsDto ('runesSearchResultsDto').
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
//       instance of CharSearchRuneArrayResultsDto. This
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
func (runesSearchResultsDto *CharSearchRuneArrayResultsDto) GetParameterTextListing(
	errorPrefix interface{}) (
	strings.Builder,
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
			"GetParameterTextListing()",
		"")

	if err != nil {
		return strings.Builder{}, err
	}

	return charSearchRuneArrayResultsDtoNanobot{}.ptr().
		getParameterTextListing(
			runesSearchResultsDto,
			ePrefix.XCpy("runesSearchResultsDto"))
}

// LoadTargetBaseInputParameters - Receives Target String data from
// input parameter 'targetInputParms' and proceeds to transfer key
// data for the search operation to the current instance of
// CharSearchRuneArrayResultsDto.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method assumes that the input data elements contained in
// 'targetInputParms' have been validated.
//
// This method therefore performs NO DATA VALIDATION on input
// parameter 'targetInputParms'.
//
func (runesSearchResultsDto *CharSearchRuneArrayResultsDto) LoadTargetBaseInputParameters(
	targetInputParms CharSearchTargetInputParametersDto) {

	if runesSearchResultsDto.lock == nil {
		runesSearchResultsDto.lock = new(sync.Mutex)
	}

	runesSearchResultsDto.lock.Lock()

	defer runesSearchResultsDto.lock.Unlock()

	runesSearchResultsDto.TargetInputParametersName =
		targetInputParms.TargetInputParametersName

	runesSearchResultsDto.FoundFirstNumericDigitInNumStr =
		targetInputParms.FoundFirstNumericDigitInNumStr

	runesSearchResultsDto.FoundDecimalSeparatorSymbols =
		targetInputParms.FoundDecimalSeparatorSymbols

	runesSearchResultsDto.FoundNonZeroValue =
		targetInputParms.FoundNonZeroValue

	runesSearchResultsDto.TargetStringLength =
		targetInputParms.TargetStringLength

	runesSearchResultsDto.TargetStringSearchLength =
		targetInputParms.TargetStringSearchLength

	runesSearchResultsDto.TargetStringAdjustedSearchLength =
		targetInputParms.TargetStringAdjustedSearchLength

	runesSearchResultsDto.TargetStringStartingSearchIndex =
		targetInputParms.TargetStringStartingSearchIndex

	runesSearchResultsDto.TargetStringCurrentSearchIndex =
		targetInputParms.TargetStringCurrentSearchIndex

	runesSearchResultsDto.TargetStringDescription1 =
		targetInputParms.TargetStringDescription1

	runesSearchResultsDto.TargetStringDescription2 =
		targetInputParms.TargetStringDescription2

}

// LoadTestBaseInputParameters - Receives Test String data from
// input parameter 'testInputParms' and proceeds to transfer key
// data for the search operation to the current instance of
// CharSearchRuneArrayResultsDto.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method assumes that the input data elements contained in
// 'testInputParms' have been validated. Therefore, NO DATA
// VALIDATION is performed on input parameter, 'testInputParms'.
//
func (runesSearchResultsDto *CharSearchRuneArrayResultsDto) LoadTestBaseInputParameters(
	testInputParms CharSearchTestInputParametersDto) {

	if runesSearchResultsDto.lock == nil {
		runesSearchResultsDto.lock = new(sync.Mutex)
	}

	runesSearchResultsDto.lock.Lock()

	defer runesSearchResultsDto.lock.Unlock()

	runesSearchResultsDto.TestInputParametersName =
		testInputParms.TestInputParametersName

	runesSearchResultsDto.TestStringName =
		testInputParms.TestStringName

	runesSearchResultsDto.TestStringLength =
		testInputParms.TestStringLength

	runesSearchResultsDto.TestStringLengthName =
		testInputParms.TestStringLengthName

	runesSearchResultsDto.TestStringStartingIndex =
		testInputParms.TestStringStartingIndex

	runesSearchResultsDto.TestStringStartingIndexName =
		testInputParms.TestStringStartingIndexName

	runesSearchResultsDto.TestStringDescription1 =
		testInputParms.TestStringDescription1

	runesSearchResultsDto.TestStringDescription2 =
		testInputParms.TestStringDescription2

	runesSearchResultsDto.CollectionTestObjIndex =
		testInputParms.CollectionTestObjIndex

	runesSearchResultsDto.TextCharSearchType =
		testInputParms.TextCharSearchType

}

// New - Returns a new and uninitialized instance of
// CharSearchRuneArrayResultsDto.
//
// All member variables in this returned instance are set to their
// zero or uninitialized states. Array index values are set to a
// value of minus one (-1) to differentiate them from valid array
// indexes which have values greater than minus one (-1).
//
func (runesSearchResultsDto *CharSearchRuneArrayResultsDto) New() CharSearchRuneArrayResultsDto {

	if runesSearchResultsDto.lock == nil {
		runesSearchResultsDto.lock = new(sync.Mutex)
	}

	runesSearchResultsDto.lock.Lock()

	defer runesSearchResultsDto.lock.Unlock()

	newRunesSearchResults := CharSearchRuneArrayResultsDto{}

	charSearchRuneArrayResultsDtoAtom{}.ptr().
		empty(&newRunesSearchResults)

	return newRunesSearchResults
}

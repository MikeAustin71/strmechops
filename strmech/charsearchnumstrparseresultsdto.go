package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// CharSearchNumStrParseResultsDto - Conveys the results of a
// number string parsing operation. The member variables of
// CharSearchNumStrParseResultsDto contain detailed information on
// the results of a specific type of character search known as
// a number string parsing operation.
//
// ----------------------------------------------------------------
//
// # TERMINOLOGY
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
// ('TextCharSearchType') rely on a framework consisting of a
// 'Target Search String' and a 'Test String'.
//
//	Target String        - A string character or characters which
//	                       will be searched for the occurrence of
//	                       another predefined character or
//	                       characters referred to as a Test
//	                       String.
//
//
//	Test String          - A string character or characters which
//	                       will be used to search for matching
//	                       characters in a Target Search String.
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
//	Number String        - As used here, a Number String is a
//	                       string of text characters which
//	                       contain numeric digit characters.
//
// Number String Parsing functions represent a specific type of
// text character search. They are designed to review a string of
// text characters searching for numeric digits. The numeric digits
// are extracted to form numeric values. Number string parsing
// functions therefore convert numeric text characters to valid
// numeric values.
//
// ----------------------------------------------------------------
//
// The Character Search Number String Results Data Transfer
// type (CharSearchNumStrParseResultsDto) is used to
// identify, summarize, and transmit the results or outcomes of a
// search operation for numeric digit characters within a number
// string.
//
// This is the primary or master type used to convey the results
// of a number string parsing operation.
type CharSearchNumStrParseResultsDto struct {
	SearchResultsName string
	// Optional. The Name, Label or descriptive Tag
	// associated with the current instance of
	// CharSearchNumStrParseResultsDto.

	SearchResultsFunctionChain string
	// Optional. A listing of the functions which led to the
	// creation of this instance of
	// CharSearchNumStrParseResultsDto.

	TargetSearchString RuneArrayDto
	// This instance of RuneArrayDto contains the original
	// Target String. The Target String contains the numeric
	// digits which are parsed and extracted by the Number
	// String Parsing algorithm.

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

	TargetStringLastSearchIndex int
	// The index in Target Search String occupied by the
	// last Target character searched.

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

	ReasonForSearchTermination CharSearchTerminationType
	// This enumeration will be set to explain why the
	// text search operation was terminated.
	//
	// Possible values are listed as follows:
	//	CharSearchTermType.None() - Invalid Value
	//	CharSearchTermType.ProcessError()
	//	CharSearchTermType.EndOfTargetString()
	//	CharSearchTermType.SearchLengthLimit()
	//	CharSearchTermType.TerminationDelimiters()
	//	CharSearchTermType.FoundSearchTarget()

	FoundNumericDigits bool
	// Signals a successful Number String Parsing operation.
	// When set to 'true', this means one or more numeric
	// digit characters ('0' through '9' inclusive) were
	// located in the Target String.

	FoundNonZeroValue bool
	// If numeric digit characters were located in the
	// Target String (see 'FoundNumericDigits' above), this
	// signals whether the numeric value is zero or nonzero.
	// When set to 'true', the found numeric value is some
	// value other than zero. If set to 'false', this means
	// that the found numeric value is zero.

	FoundDecimalSeparatorSymbols bool
	// When set to 'true', this signals that one or more
	// Decimal Separator Symbol characters have been
	// identified in the 'Target String'.

	FoundIntegerDigits bool
	// When set to 'true', this signals that one or more
	// integer digits ('0' through '9' inclusive) were
	// located in the Target String to the left of the
	// Decimal Separator Symbol (a.k.a. Decimal Point).

	IdentifiedIntegerDigits string
	// If Integer Digits (to the left of the decimal
	// separator) were found during the character search,
	// those digits will be stored in this string.

	FoundDecimalDigits bool
	// When set to 'true', this signals that one or more
	// numeric digit characters ('0' through '9' inclusive)
	// were located in the Target String to the right of
	// the Decimal Separator Symbol (a.k.a Decimal Point).
	//
	// Decimal Digits are Fractional Digits.

	IdentifiedFractionalDigits string
	// If Fractional Digits (to the right of the decimal
	// separator) were found during the character search,
	// those digits will be stored in this string.

	NumSignValue NumericSignValueType
	// If a numeric value was extracted from the number string
	// (see 'FoundNumericDigits' above), this parameter
	// specifies the number sign associated with that value.
	//
	// Possible values are listed as follows:
	//  NumSignVal.None() - Signals no numeric value was found
	//  NumSignVal.Negative()
	//  NumSignVal.Zero()
	//  NumSignVal.Positive()

	NumValueType NumericValueType
	// If a numeric value was extracted from the number string
	// (see 'FoundNumericDigits' above), this parameter signals
	// whether the numeric value is an integer or a floating
	// point (digits to the right of the decimal) value.
	//
	// Possible values are listed as follows:
	//  NumValType.None() - Signals no numeric value was found
	//  NumValType.FloatingPoint()
	//  NumValType.Integer()

	RemainderString RuneArrayDto
	// A Rune Array Data Transfer Object containing the
	// remaining characters in a Target String which were
	// NOT included in the search operation and which remain
	// to be searched in future search operations.
	//
	// The 'RemainderString' is only returned if requested.

	DecimalSeparatorSearchResults CharSearchDecimalSeparatorResultsDto
	// An instance of CharSearchDecimalSeparatorResultsDto detailing
	// the results of a search for a decimal separator in the Target
	// String.

	NegativeNumberSymbolSearchResults CharSearchNegativeNumberResultsDto
	// An instance of CharSearchNegativeNumberResultsDto detailing the
	// results of a search for negative number symbols in the Target
	// String.

	ParsingTerminatorSearchResults CharSearchRuneArrayResultsDto
	// An instance of CharSearchRuneArrayResultsDto detailing the
	// results of a search for a Number String Parsing Delimiter
	// character or characters. When a Parsing Terminator Delimiter
	// is located in the Target String, the Number Parsing operation
	// is immediately terminated.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// CharSearchNumStrParseResultsDto ('incomingNumStrParseResults')
// to the data fields of the current
// CharSearchNumStrParseResultsDto instance
// ('searchNumStrParseResults').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in the current
// CharSearchNumStrParseResultsDto instance
// ('searchNumStrParseResults') will be deleted and overwritten.
//
// No Data Validation will be performed on
// 'incomingNumStrParseResults'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumStrParseResults   *CharSearchNumStrParseResultsDto
//	   - A pointer to an instance of
//	     CharSearchNumStrParseResultsDto. This method will NOT
//	     change the data values of member variables contained in
//	     this instance.
//
//	     All data values in this CharSearchNumStrParseResultsDto
//	     instance ('incomingNumStrParseResults') will be copied to
//	     the current CharSearchNumStrParseResultsDto instance
//	     ('searchNumStrParseResults').
//
//	     No Data Validation will be performed on
//	     'incomingNumStrParseResults'.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (searchNumStrParseResults *CharSearchNumStrParseResultsDto) CopyIn(
	incomingNumStrParseResults *CharSearchNumStrParseResultsDto,
	errorPrefix interface{}) error {

	if searchNumStrParseResults.lock == nil {
		searchNumStrParseResults.lock = new(sync.Mutex)
	}

	searchNumStrParseResults.lock.Lock()

	defer searchNumStrParseResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchNumStrParseResultsDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(charSearchNumStrParseResultsDtoNanobot).
		copyNumStrParseResultsDto(
			searchNumStrParseResults,
			incomingNumStrParseResults,
			ePrefix.XCpy(
				"searchNumStrParseResults<-"+
					"incomingNumStrParseResults"))
}

// CopyOut - Returns a deep copy of the current
// CharSearchNumStrParseResultsDto instance.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// No Data Validation will be performed on the current instance
// of CharSearchNumStrParseResultsDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	CharSearchNumStrParseResultsDto
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current CharSearchNumStrParseResultsDto instance.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message occurs, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (searchNumStrParseResults *CharSearchNumStrParseResultsDto) CopyOut(
	errorPrefix interface{}) (
	deepCopyNumStrParseResults CharSearchNumStrParseResultsDto,
	err error) {

	if searchNumStrParseResults.lock == nil {
		searchNumStrParseResults.lock = new(sync.Mutex)
	}

	searchNumStrParseResults.lock.Lock()

	defer searchNumStrParseResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchNumStrParseResultsDto."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNumStrParseResults, err
	}

	err = new(charSearchNumStrParseResultsDtoNanobot).
		copyNumStrParseResultsDto(
			&deepCopyNumStrParseResults,
			searchNumStrParseResults,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-searchNumStrParseResults"))

	return deepCopyNumStrParseResults, err
}

// Empty - Resets all internal member variables for the current
// instance of CharSearchNumStrParseResultsDto to their zero or
// uninitialized states. This method will leave the current
// instance of CharSearchDecimalSeparatorResultsDto in an invalid
// state and unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in the
// current instance of CharSearchNumStrParseResultsDto. All member
// variable data values will be reset to their zero or
// uninitialized states.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (searchNumStrParseResults *CharSearchNumStrParseResultsDto) Empty() {

	if searchNumStrParseResults.lock == nil {
		searchNumStrParseResults.lock = new(sync.Mutex)
	}

	searchNumStrParseResults.lock.Lock()

	new(charSearchNumStrParseResultsDtoAtom).
		empty(searchNumStrParseResults)

	searchNumStrParseResults.lock.Unlock()

	searchNumStrParseResults.lock = nil

	return
}

// Equal - Receives a pointer to another instance of
// CharSearchNumStrParseResultsDto and proceeds to compare the
// member variables to those of the current
// CharSearchNumStrParseResultsDto instance in order to
// determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumStrParseResults *CharSearchNumStrParseResultsDto
//	   - A pointer to an incoming instance of
//	     CharSearchNumStrParseResultsDto. This method will
//	     compare all member variable data values in this instance
//	     against those contained in the current instance of
//	     CharSearchNumStrParseResultsDto. If the data values
//	     in both instances are found to be equal in all respects,
//	     this method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingNumStrParseResults' are equal in all
//	     respects to those contained in the current instance of
//	     CharSearchNumStrParseResultsDto, this method will
//	     return a boolean value of 'true'. Otherwise, a value of
//	     'false' will be returned to the calling function.
func (searchNumStrParseResults *CharSearchNumStrParseResultsDto) Equal(
	incomingNumStrParseResults *CharSearchNumStrParseResultsDto) bool {

	if searchNumStrParseResults.lock == nil {
		searchNumStrParseResults.lock = new(sync.Mutex)
	}

	searchNumStrParseResults.lock.Lock()

	defer searchNumStrParseResults.lock.Unlock()

	return new(charSearchNumStrParseResultsDtoAtom).
		equal(
			searchNumStrParseResults,
			incomingNumStrParseResults)
}

// GetParameterTextListing - Returns formatted text output
// detailing the member variable names and their corresponding
// values contained in the current instance of
// CharSearchNumStrParseResultsDto ('searchNumStrParseResults').
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strBuilder                 *strings.Builder
//	   - A pointer to an instance of *strings.Builder. The
//	     formatted text characters produced by this method will be
//	     written to this instance of strings.Builder.
//
//
//	displayFunctionChain       bool
//	   - Set 'displayFunctionChain' to 'true' and a list of the
//	     functions which led to this result will be included in
//	     the formatted text output.
//
//
//	printDetail                bool
//	   - If this parameter is set to 'true', detail information for
//	     subsidiary types RemainderString,
//	     DecimalSeparatorSearchResults,
//	     NegativeNumberSymbolSearchResults and
//	     ParsingTerminatorSearchResults will be included in the
//	     formatted text output.
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of
//	                        ErrPrefixDto. ErrorPrefixInfo from this
//	                        object will be copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (searchNumStrParseResults *CharSearchNumStrParseResultsDto) GetParameterTextListing(
	strBuilder *strings.Builder,
	displayFunctionChain bool,
	printDetail bool,
	errorPrefix interface{}) error {

	if searchNumStrParseResults.lock == nil {
		searchNumStrParseResults.lock = new(sync.Mutex)
	}

	searchNumStrParseResults.lock.Lock()

	defer searchNumStrParseResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchNumStrParseResultsDto."+
			"GetParameterTextListing()",
		"")

	if err != nil {
		return err
	}

	return new(charSearchNumStrParseResultsDtoNanobot).
		getParameterTextListing(
			strBuilder,
			searchNumStrParseResults,
			displayFunctionChain,
			printDetail,
			ePrefix.XCpy(
				"searchNumStrParseResults"))
}

// New - Returns a new and uninitialized instance of
// CharSearchNumStrParseResultsDto.
//
// All member variables in this returned instance are set to their
// zero or uninitialized states.
func (searchNumStrParseResults CharSearchNumStrParseResultsDto) New() CharSearchNumStrParseResultsDto {

	if searchNumStrParseResults.lock == nil {
		searchNumStrParseResults.lock = new(sync.Mutex)
	}

	searchNumStrParseResults.lock.Lock()

	defer searchNumStrParseResults.lock.Unlock()

	newNumStrParseResults := CharSearchNumStrParseResultsDto{}

	new(charSearchNumStrParseResultsDtoAtom).
		empty(&newNumStrParseResults)

	return newNumStrParseResults
}

// String - Returns a formatted text string detailing all the
// internal member variable names and their corresponding values
// for the current instance of
// CharSearchNumStrParseResultsDto.
//
// If an error is encountered, the error message is included in the
// string returned by this method.
//
// This method implements the Stringer Interface.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// This method will NOT include the detail information on subsidiary
// types RemainderString, DecimalSeparatorSearchResults,
// NegativeNumberSymbolSearchResults and
// ParsingTerminatorSearchResults. If this detail information is
// required in the formatted text output, call method:
//
//	CharSearchNumStrParseResultsDto.GetParameterTextListing()
func (searchNumStrParseResults *CharSearchNumStrParseResultsDto) String() string {

	if searchNumStrParseResults.lock == nil {
		searchNumStrParseResults.lock = new(sync.Mutex)
	}

	searchNumStrParseResults.lock.Lock()

	defer searchNumStrParseResults.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		nil,
		"CharSearchNumStrParseResultsDto."+
			"String()",
		"")

	if err != nil {
		errOut := fmt.Sprintf("%v\n"+
			"Error Message:\n"+
			"%v",
			"CharSearchNumStrParseResultsDto."+
				"String()",
			err.Error())

		return errOut
	}

	strBuilder := strings.Builder{}

	err = new(charSearchNumStrParseResultsDtoNanobot).
		getParameterTextListing(
			&strBuilder,
			searchNumStrParseResults,
			true,
			false,
			ePrefix.XCpy(
				"searchNumStrParseResults"))

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

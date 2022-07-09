package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
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
// ('TextCharSearchType') rely on a framework consisting of a
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
// ----------------------------------------------------------------
//
// The Character Search Number String Results Data Transfer
// type (CharSearchNumStrParseResultsDto) is used to
// identify, sumarize, and transmit the results or outcomes of a
// search operation for numeric digit characters within a number
// string.
//
// This is the primary or master type used to convey the results
// of a number string parsing operation.
//
type CharSearchNumStrParseResultsDto struct {
	SearchResultsName string
	// Optional. The Name, Label or descriptive Tag
	// associated with the current instance of
	// CharSearchNumStrParseResultsDto.

	SearchResultsFunctionChain string
	// Optional. A listing of the functions which led to the
	// creation of this instance of
	// CharSearchNumStrParseResultsDto.

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

	FoundDecimalDigits bool
	// When set to 'true', this signals that one or more
	// numeric digit characters ('0' through '9' inclusive)
	// were located in the Target String to the right of
	// the Decimal Separator Symbol (a.k.a Decimal Point).

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
// IMPORTANT
//
// All the data fields in the current
// CharSearchNumStrParseResultsDto instance
// ('searchNumStrParseResults') will be deleted and overwritten.
//
// No Data Validation will be performed on
// 'incomingNumStrParseResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNumStrParseResults   *CharSearchNumStrParseResultsDto
//     - A pointer to an instance of
//       CharSearchNumStrParseResultsDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this CharSearchNumStrParseResultsDto
//       instance ('incomingNumStrParseResults') will be copied to
//       the current CharSearchNumStrParseResultsDto instance
//       ('searchNumStrParseResults').
//
//       No Data Validation will be performed on
//       'incomingNumStrParseResults'.
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

	return charSearchNumStrParseResultsDtoNanobot{}.ptr().
		copyIn(
			searchNumStrParseResults,
			incomingNumStrParseResults,
			ePrefix.XCpy(
				"searchNumStrParseResults<-"+
					"incomingNumStrParseResults"))
}

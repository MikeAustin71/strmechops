package strmech

import "sync"

// CharSearchDecimalSeparatorResultsDto  - Contains parameters
// detailing the results of a text character search for decimal
// separator characters or symbols.
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
// classification of numeric values integer or floating point
// values. This classification logic searches for a specific
// decimal separator character or characters which separate integer
// numeric digits from fractional numeric digits.
//
// ----------------------------------------------------------------
//
// The Character Search Decimal Separator Results Data Transfer
// Object type (CharSearchDecimalSeparatorResultsDto) is used to
// identify and transmit the results or outcomes of a search
// operation for decimal separator symbols in a number string.
//
// Methods performing elements of this text character search
// operation use this wrapper type to encapsulate and return key
// data variables describing the search outcome.
//
type CharSearchDecimalSeparatorResultsDto struct {
	SearchResultsName string
	// Optional. The Name, Label or descriptive Tag associated with
	// the current instance of CharSearchResultsDto.

	SearchResultsFunctionChain string
	// Optional. A listing of the functions which led to the
	// creation of this instance of CharSearchResultsDto.

	FoundDecimalSeparatorSymbols bool
	// Signals a successful search outcome. If set to 'true' the
	// Decimal Separator Symbol character or characters were found
	// in the Target Search String.

	FoundDecimalSepSymbolsOnPreviousSearch bool
	// Signals that Decimal Separator Symbols were located in a
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

	NumValueType NumericValueType
	// Optional. This enumeration value specifies the type
	// of numeric value as either an integer or floating
	// point value.
	//
	// Possible values are listed as follows:
	//  NumValType.None()
	//  NumValType.FloatingPoint()
	//  NumValType.Integer()

	NumSymbolLocation NumericSymbolLocation
	// Optional. This enumeration value specifies the
	// relative location of a numeric symbol. If integer
	// digits were found before the Decimal Separator,
	// this value is set to NumSymLocation.Interior().
	// If no integer digits were found before the Decimal
	// Separator, this value is set to
	// NumSymLocation.Before().
	//
	// Possible values are listed as follows:
	//  NumSymLocation.None()
	//  NumSymLocation.Before()
	//  NumSymLocation.Interior()
	//  NumSymLocation.After()

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

	DecimalSeparatorSymbolsSpec DecimalSeparatorSpec
	// If Decimal Separator symbols were found in the current
	// search operation, they will be stored in this instance
	// of DecimalSeparatorSpec

	lock *sync.Mutex
}

// Empty - Resets all internal member variables for the current
// instance of CharSearchDecimalSeparatorResultsDto to their zero
// or uninitialized states. This method will leave the current
// instance of CharSearchDecimalSeparatorResultsDto in an
// invalid state and unavailable for immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in the
// current instance of CharSearchDecimalSeparatorResultsDto. All
// member variable data values will be reset to their zero or
// uninitialized states. Array index values will be set to minus
// one (-1). Valid array indexes have values greater than minus one
// (-1).
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
func (decSepSearchResultsDto *CharSearchDecimalSeparatorResultsDto) Empty() {

	if decSepSearchResultsDto.lock == nil {
		decSepSearchResultsDto.lock = new(sync.Mutex)
	}

	decSepSearchResultsDto.lock.Lock()

	charSearchDecimalSeparatorResultsDtoAtom{}.ptr().
		empty(decSepSearchResultsDto)

	decSepSearchResultsDto.lock.Unlock()

	decSepSearchResultsDto.lock = nil
}

// New - Returns a new and uninitialized instance of
// CharSearchDecimalSeparatorResultsDto.
//
// All member variables in this returned instance are set to their
// zero or uninitialized states. Array index values are set to a
// value of minus one (-1). All valid array indexes have values
// greater than minus one (-1).
//
func (decSepSearchResultsDto CharSearchDecimalSeparatorResultsDto) New() CharSearchDecimalSeparatorResultsDto {

	if decSepSearchResultsDto.lock == nil {
		decSepSearchResultsDto.lock = new(sync.Mutex)
	}

	decSepSearchResultsDto.lock.Lock()

	defer decSepSearchResultsDto.lock.Unlock()

	newDecSepSearchResults := CharSearchDecimalSeparatorResultsDto{}

	return newDecSepSearchResults
}

package strmech

import (
	"sync"
)

type CharSearchResultsDto struct {
	FoundSearchTarget bool
	// Signals a successful search outcome. If set to
	// 'true' the Test String character or characters were
	// found in the Target Search String.

	FoundSearchTargetOnPreviousSearch bool
	// Signals that the Search Target was located in a previous
	// search operation.

	FoundFirstNumericDigitInNumStr bool
	//            - When set to 'true' this signals that the search
	//              operation has already identified the first numeric
	//              digit in a string of text characters.

	TargetStringLength int
	// Actual number of text characters in the entire
	// Target Search String ('TargetString').

	TargetStringSearchLength int
	// The actual number of characters within the Target
	// Search String that are included in the search
	// operation. This value may be less than the actual
	// length of the Target Search String.

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

	TargetStringDescription1 string
	// First of two optional description strings
	// describing the Target Search String in the context
	// of the current search operation.

	TargetStringDescription2 string
	// Second of two optional description strings
	// describing the Target Search String in the context
	// of the current search operation.

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
	//  TextCharSearchType.None() - Invalid value
	//  TextCharSearchType.LinearTargetStartingIndex() - Default
	//  TextCharSearchType.SingleTargetChar()
	//  TextCharSearchType.LinearEndOfString()

	lock *sync.Mutex
}

func (charSearchResults *CharSearchResultsDto) Empty() {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	charSearchResults.FoundSearchTarget = false

	charSearchResults.FoundSearchTargetOnPreviousSearch = false

	charSearchResults.TargetStringLength = -1

	charSearchResults.TargetStringStartingSearchIndex = -1

	charSearchResults.TargetStringFirstFoundIndex = -1

	charSearchResults.TargetStringLastFoundIndex = -1

	charSearchResults.TargetStringLastSearchIndex = -1

	charSearchResults.TargetStringDescription1 = ""

	charSearchResults.TargetStringDescription2 = ""

	charSearchResults.TestStringLength = -1

	charSearchResults.TestStringStartingIndex = -1

	charSearchResults.TestStringFirstFoundIndex = -1

	charSearchResults.TestStringLastFoundIndex = -1

	charSearchResults.TestStringDescription1 = ""

	charSearchResults.TestStringDescription2 = ""

	charSearchResults.CollectionTestObjIndex = -1

	charSearchResults.NumValueType = NumValType.None()

	charSearchResults.NumStrFormatType = NumStrFmtType.None()

	charSearchResults.NumSymLocation = NumSymLocation.None()

	charSearchResults.NumSymbolClass = NumSymClass.None()

	charSearchResults.NumSignValue = NumSignVal.None()

	charSearchResults.PrimaryNumSignPosition = NumSignSymPos.None()

	charSearchResults.SecondaryNumSignPosition = NumSignSymPos.None()

	charSearchResults.TextCharSearchType =
		CharSearchType.None()

	charSearchResults.FoundFirstNumericDigitInNumStr = false

	charSearchResults.lock.Unlock()

	charSearchResults.lock = nil
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

func (charSearchResults *CharSearchResultsDto) LoadTargetBaseInputParameters(
	targetInputParms CharSearchTargetInputParametersDto) {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	defer charSearchResults.lock.Unlock()

	charSearchResults.FoundFirstNumericDigitInNumStr =
		targetInputParms.FoundFirstNumericDigitInNumStr

	charSearchResults.TargetStringLength =
		targetInputParms.TargetStringLength

	charSearchResults.TargetStringSearchLength =
		targetInputParms.TargetStringAdjustedSearchLength

	charSearchResults.TargetStringStartingSearchIndex =
		targetInputParms.TargetStringStartingSearchIndex

	charSearchResults.TargetStringDescription1 =
		targetInputParms.TargetStringDescription1

	charSearchResults.TargetStringDescription2 =
		targetInputParms.TargetStringDescription2

}

func (charSearchResults *CharSearchResultsDto) LoadTestBaseInputParameters(
	testInputParms CharSearchTestInputParametersDto) {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	defer charSearchResults.lock.Unlock()

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

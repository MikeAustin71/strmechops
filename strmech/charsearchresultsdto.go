package strmech

import (
	"sync"
)

type CharSearchResultsDto struct {
	FoundSearchTarget bool // Signals a successful search
	//                        outcome. If 'true' the Test
	//                        String character or characters
	//                        were found in the Target Search
	//                        String.

	FoundSearchTargetOnPreviousSearch bool // The Search Target was located
	//                                        on a previous search operation.

	TargetStringLength int // Actual Full Length of the Target Search
	//                        String

	TargetStringSearchLength int // The actual number of characters within
	//                              Target Search String that are included
	//                              in the search specification.

	TargetStringStartingSearchIndex int // The index at which the search
	//                               	   operation commenced.

	TargetStringFirstFoundIndex int // The first character position in the
	//                                 Target Search String occupied by the
	//                                 first character in the Test String.

	TargetStringLastFoundIndex int // The last character position in
	//                                the Target Search String occupied
	//                                by the last character in the Test
	//                                String.

	TargetStringLastSearchIndex int // The index in Target Search String
	//                                 occupied by the last character searched.
	//                                 If the Search Target was found this
	//                                 value is equal to the 'TargetStringLastFoundIndex'.
	//                                 If the Search Target was NOT found this
	//                                 value is equal to the 'TargetStringStartingSearchIndex'

	TargetStringDescription1 string // First optional description string
	//                                  describing the Target Search String
	//                                  used in this search

	TargetStringDescription2 string // Second Optional description string
	//                                  describing the Target Search String
	//                                  used in this search

	TestStringLength int // The Length of the Test String used in
	//                      this search operation.

	TestStringStartingIndex int // The Test String beginning index used in the
	//                             search operation. Usually zero.

	TestStringFirstFoundIndex int // The index number in Test String of the first
	//                               test character to be located in the Target
	//                               Search String.

	TestStringLastFoundIndex int // The index number in the Test String occupied
	//                              by the last Test Character to be located in the
	//                              Target Search String.

	TestStringDescription1 string // Optional description string describing
	//                               Test Characters used in this search

	TestStringDescription2 string // Optional description string describing
	//                               Test Characters used in the search

	CollectionTestObjIndex int // The collection index of the object
	//                            containing the Test String which was
	//                            found in Target Search String

	NumValueType NumericValueType // Specifies the numeric value
	//                               as an integer or floating point value.
	//                                 NumValType.None(),
	//                                 NumValType.FloatingPoint(),
	//                                 NumValType.Integer(),

	NumStrFormatType NumStrFormatTypeCode // Specifies Output Format Type for
	//                                       a numeric value.
	//                                         NumStrFmtType.None()
	//                                         NumStrFmtType.AbsoluteValue()
	//                                         NumStrFmtType.Binary()
	//                                         NumStrFmtType.CountryCulture()
	//                                         NumStrFmtType.Currency()
	//                                         NumStrFmtType.Binary()
	//                                         NumStrFmtType.Hexadecimal()
	//                                         NumStrFmtType.Octal()
	//                                         NumStrFmtType.ScientificNotation()

	NumSymLocation NumericSymbolLocation // Specifies the relative location of a
	//                                        numeric symbol.
	//                                          NumSymLocation.None(),
	//                                          NumSymLocation.Before(),
	//                                          NumSymLocation.Interior(),
	//                                          NumSymLocation.After(),

	NumSymbolClass NumericSymbolClass // Number Symbol Classification
	//                                      NumSymClass.None(),
	//                                      NumSymClass.NumberSign(),
	//                                      NumSymClass.CurrencySign(),
	//                                      NumSymClass.IntegerSeparator(),
	//                                      NumSymClass.DecimalSeparator(),

	NumSignValue NumericSignValueType // An enumeration value classifying the
	//                                   number sign.
	//                                     NumSignVal.None()
	//                                     NumSignVal.Negative()
	//                                     NumSignVal.Zero()
	//                                     NumSignVal.Positive()

	PrimaryNumSignPosition NumSignSymbolPosition // Used in searches involving
	//                                       positive and negative number
	//                                       signs.
	//                                        NumSignSymPos.None()
	//                                        NumSignSymPos.Before()
	//                                        NumSignSymPos.After()
	//                                        NumSignSymPos.BeforeAndAfter()

	SecondaryNumSignPosition NumSignSymbolPosition // Used in searches involving
	//                                                signs which occur both before
	//                                                and after the numeric value.
	//                                                 NumSignSymPos.None()
	//                                                 NumSignSymPos.Before()
	//                                                 NumSignSymPos.After()

	CharSearchType CharacterSearchType // An enumeration value signaling
	//                                the type of search algorithm which
	//                                was used to conduct this search.
	//                                 CharSearchType.None()
	//                                 CharSearchType.LinearTargetStartingIndex() - Default
	//                                 CharSearchType.SingleTargetChar()
	//                                 CharSearchType.LinearEndOfString()

	FoundFirstNumericDigitInNumStr bool // When set to 'true' this signals
	//                                     that the first numeric digit has
	//                                     been identified in a string of text
	//                                     characters.

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

	charSearchResults.CharSearchType =
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

	charSearchResults.CharSearchType =
		testInputParms.CharSearchType

}

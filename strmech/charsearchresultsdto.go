package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type CharSearchResultsDto struct {
	FoundSearchTarget bool // Signals a successful search
	//                        outcome. If 'true' the Test
	//                        String character or characters
	//                        were found in the Target Search
	//                        String.

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

	NumSignPosition NumSignSymbolPosition // Used in searches involving
	//                                       positive and negative number
	//                                       signs.
	//                                        NumSignSymPos.None()
	//                                        NumSignSymPos.Before()
	//                                        NumSignSymPos.After()
	//                                        NumSignSymPos.BeforeAndAfter()

	CharSearchType CharacterSearchType // An enumeration value signaling
	//                                the type of search algorithm which
	//                                was used to conduct this search.
	//                                 CharSearchType.None()
	//                                 CharSearchType.LinearTargetStartingIndex() - Default
	//                                 CharSearchType.SingleTargetChar()
	//                                 CharSearchType.LinearEndOfString()

	lock *sync.Mutex
}

func (charSearchResults *CharSearchResultsDto) Empty() {

	if charSearchResults.lock == nil {
		charSearchResults.lock = new(sync.Mutex)
	}

	charSearchResults.lock.Lock()

	charSearchResults.FoundSearchTarget = false

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

	charSearchResults.NumSignPosition = NumSignSymPos.None()

	charSearchResults.CharSearchType =
		CharSearchType.None()

	charSearchResults.lock.Unlock()

	charSearchResults.lock = nil
}

func (charSearchResults *CharSearchResultsDto) LoadBaseCharSearchInputParameters(
	inputParms CharSearchInputParametersDto,
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
		"RuneArrayDto."+
			"SetCharacterSearchType()",
		"")

	if err != nil {
		return err
	}

	err = inputParms.ValidateTestString(
		ePrefix)

	if err != nil {
		return err
	}

	err = inputParms.ValidateTargetSearchString(
		ePrefix)

	if err != nil {
		return err
	}

	charSearchResults.TestStringLength =
		inputParms.TestStringLength

	charSearchResults.TestStringStartingIndex =
		inputParms.TestStringStartingIndex

	charSearchResults.TestStringDescription1 =
		inputParms.TestStringDescription1

	charSearchResults.TestStringDescription2 =
		inputParms.TestStringDescription2

	charSearchResults.TargetStringLength =
		inputParms.TargetSearchStringLength

	charSearchResults.TargetStringSearchLength =
		inputParms.TargetStringAdjustedSearchLength

	charSearchResults.TargetStringStartingSearchIndex =
		inputParms.TargetStringStartingSearchIndex

	charSearchResults.TargetStringDescription1 =
		inputParms.TargetStringDescription1

	charSearchResults.TargetStringDescription2 =
		inputParms.TargetStringDescription2

	charSearchResults.CollectionTestObjIndex =
		inputParms.CollectionTestObjIndex

	charSearchResults.NumValueType =
		inputParms.NumValueType

	charSearchResults.NumStrFormatType =
		inputParms.NumStrFormatType

	charSearchResults.NumSymLocation =
		inputParms.NumSymLocation

	charSearchResults.NumSymbolClass =
		inputParms.NumSymbolClass

	charSearchResults.NumSignValue =
		inputParms.NumSignValue

	charSearchResults.CharSearchType =
		inputParms.CharSearchType

	return err
}

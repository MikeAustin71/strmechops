package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// CharSearchInputParametersDto - The Character Search Input
// Parameters Data Transfer Object is used to transmit input
// parameters for character search algorithms.
//
type CharSearchInputParametersDto struct {
	TestString *RuneArrayDto // The Rune Array containing the Test
	//                          Characters to be used in a search
	//                          algorithm.

	TestStringName string // The label or name of the TestString.

	TestStringLength int // The Length of the Test String used in
	//                      this search operation.

	TestStringLengthName string // The label or name of the TestStringLength
	//                              parameter.

	TestStringStartingIndex int // The starting index for the Test
	//                              String. Search Comparisons will begin
	//                              at this point in the Test String.

	TestStringDescription1 string // Optional description string describing
	//                               Test Characters used in this search

	TestStringDescription2 string // Optional description string describing
	//                               Test Characters used in the search

	TargetSearchString *RuneArrayDto

	TargetSearchStringName string // The label or name of the TargetSearchString
	//                               parameter.

	TargetSearchStringLength int // Actual Full Length of the Target Search
	//                              String.

	TargetSearchStringLengthName string // The label or name of the
	//                                     TargetSearchStringLength parameter.

	TargetStringStartingSearchIndex int // The index at which the search
	//                               	   operation commenced.

	TargetStringStartingSearchIndexName string // The label or name of the
	//                                            TargetStringStartingSearchIndex
	//                                            parameter.

	TargetStringSearchLength int // The actual number of characters within
	//                              Target Search String that are included
	//                              in the search specification.

	TargetStringSearchLengthName string // The label or name of the
	//                                     TargetStringSearchLength parameter.

	TargetStringAdjustedSearchLength int // The adjusted Target String Search Length
	//                                   guaranteed to be equal to or less than
	//                                   the actual Target String Length.

	TargetStringDescription1 string // First optional description string
	//                                  describing the Target Search String
	//                                  used in this search

	TargetStringDescription2 string // Second Optional description string
	//                                  describing the Target Search String
	//                                  used in this search

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

func (searchInputParms *CharSearchInputParametersDto) Empty() {

	if searchInputParms.lock == nil {
		searchInputParms.lock = new(sync.Mutex)
	}

	searchInputParms.lock.Lock()

	searchInputParms.TestString = nil

	searchInputParms.TestStringName = ""

	searchInputParms.TestStringLength = -1

	searchInputParms.TestStringLengthName = ""

	searchInputParms.TestStringStartingIndex = -1

	searchInputParms.TestStringDescription1 = ""

	searchInputParms.TestStringDescription2 = ""

	searchInputParms.TargetSearchString = nil

	searchInputParms.TargetSearchStringName = ""

	searchInputParms.TargetSearchStringLength = -1

	searchInputParms.TargetSearchStringLengthName = ""

	searchInputParms.TargetStringStartingSearchIndex = -1

	searchInputParms.TargetStringStartingSearchIndexName = ""

	searchInputParms.TargetStringSearchLength = -2

	searchInputParms.TargetStringSearchLengthName = ""

	searchInputParms.TargetStringAdjustedSearchLength = -2

	searchInputParms.TargetStringDescription1 = ""

	searchInputParms.TargetStringDescription2 = ""

	searchInputParms.CollectionTestObjIndex = -1

	searchInputParms.NumValueType = NumValType.None()

	searchInputParms.NumStrFormatType = NumStrFmtType.None()

	searchInputParms.NumSymLocation = NumSymLocation.None()

	searchInputParms.NumSymbolClass = NumSymClass.None()

	searchInputParms.NumSignValue = NumSignVal.None()

	searchInputParms.NumSignPosition = NumSignSymPos.None()

	searchInputParms.CharSearchType = CharSearchType.None()

	searchInputParms.lock.Unlock()

	searchInputParms.lock = nil
}

// ValidateTargetSearchString - Validates the Target Search String
// and related member variables contained in the current instance
// of CharSearchInputParametersDto.
//
func (searchInputParms *CharSearchInputParametersDto) ValidateTargetSearchString(
	errorPrefix interface{}) error {

	if searchInputParms.lock == nil {
		searchInputParms.lock = new(sync.Mutex)
	}

	searchInputParms.lock.Lock()

	defer searchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchInputParametersDto."+
			"ValidateTargetSearchString()",
		"")

	if err != nil {

		return err

	}

	if len(searchInputParms.TargetSearchStringName) == 0 {
		searchInputParms.TargetSearchStringName =
			"TargetSearchString"
	}

	if searchInputParms.TargetSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			searchInputParms.TargetSearchStringName)

		return err
	}

	if len(searchInputParms.TargetSearchStringLengthName) == 0 {
		searchInputParms.TargetSearchStringLengthName =
			"TargetSearchStringLength"
	}

	searchInputParms.TargetSearchStringLength =
		len(searchInputParms.TargetSearchString.CharsArray)

	if searchInputParms.TargetSearchStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			searchInputParms.TargetSearchStringLengthName,
			searchInputParms.TargetSearchStringName,
			searchInputParms.TargetSearchStringName)

		return err
	}

	if len(searchInputParms.TargetStringStartingSearchIndexName) == 0 {
		searchInputParms.TargetStringStartingSearchIndexName =
			"TargetStringStartingSearchIndex"
	}

	if searchInputParms.TargetStringStartingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter %v is invalid!\n"+
			"%v is less than zero (0)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			searchInputParms.TargetStringStartingSearchIndexName,
			searchInputParms.TargetStringStartingSearchIndexName,
			searchInputParms.TargetStringStartingSearchIndexName,
			searchInputParms.TargetStringStartingSearchIndex)

		return err
	}

	if searchInputParms.TargetStringStartingSearchIndex >=
		searchInputParms.TargetSearchStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value greater than the last\n"+
			"index in '%v.CharsArray'.\n"+
			"Last Index in %v.CharsArray = '%v'\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			searchInputParms.TargetStringStartingSearchIndexName,
			searchInputParms.TargetStringStartingSearchIndexName,
			searchInputParms.TargetSearchStringName,
			searchInputParms.TargetSearchStringName,
			searchInputParms.TargetSearchStringLength-1,
			searchInputParms.TargetStringStartingSearchIndexName,
			searchInputParms.TargetStringStartingSearchIndex)

		return err
	}

	if len(searchInputParms.TargetStringSearchLengthName) == 0 {
		searchInputParms.TargetStringSearchLengthName =
			"TargetStringSearchLength"
	}

	if searchInputParms.TargetStringSearchLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value less than minus one (-1)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			searchInputParms.TargetStringSearchLengthName,
			searchInputParms.TargetStringSearchLengthName,
			searchInputParms.TargetSearchStringName,
			searchInputParms.TargetStringSearchLength)

		return err
	}

	if searchInputParms.TargetStringSearchLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value of Zero (0)\n",
			ePrefix.String(),
			searchInputParms.TargetStringSearchLengthName,
			searchInputParms.TargetStringSearchLengthName)

	}

	if searchInputParms.TargetStringSearchLength == -1 {

		searchInputParms.TargetStringAdjustedSearchLength =
			searchInputParms.TargetSearchStringLength -
				searchInputParms.TargetStringStartingSearchIndex
	} else {

		searchInputParms.TargetStringAdjustedSearchLength =
			searchInputParms.TargetStringSearchLength

	}

	searchInputParms.TargetStringAdjustedSearchLength =
		searchInputParms.TargetStringStartingSearchIndex +
			searchInputParms.TargetStringAdjustedSearchLength

	if searchInputParms.TargetStringAdjustedSearchLength >
		searchInputParms.TargetSearchStringLength {

		searchInputParms.TargetStringAdjustedSearchLength =
			searchInputParms.TargetSearchStringLength

	}

	return err
}

// ValidateTestString - Validates the Test String and related
// member variables contained in the current instance of
// CharSearchInputParametersDto.
//
func (searchInputParms *CharSearchInputParametersDto) ValidateTestString(
	errorPrefix interface{}) error {

	if searchInputParms.lock == nil {
		searchInputParms.lock = new(sync.Mutex)
	}

	searchInputParms.lock.Lock()

	defer searchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchInputParametersDto."+
			"ValidateTestString()",
		"")

	if err != nil {

		return err

	}

	if len(searchInputParms.TestStringName) == 0 {
		searchInputParms.TestStringName = "TestString"
	}

	if searchInputParms.TestString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			searchInputParms.TestStringName)

		return err
	}

	if len(searchInputParms.TestStringLengthName) == 0 {
		searchInputParms.TestStringLengthName =
			"TestStringLengthName"
	}

	searchInputParms.TestStringLength =
		len(searchInputParms.TestString.CharsArray)

	if searchInputParms.TestStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			searchInputParms.TestStringName,
			searchInputParms.TestStringName,
			searchInputParms.TestStringName)

		return err
	}

	if searchInputParms.TestStringStartingIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' Starting Index is invalid!\n"+
			"The '%v' Starting Index is less than Zero (0).\n"+
			"%v Starting Index = '%v'.\n",
			ePrefix.String(),
			searchInputParms.TestStringName,
			searchInputParms.TestStringName,
			searchInputParms.TestStringName,
			searchInputParms.TestStringStartingIndex)

		return err

	}

	if searchInputParms.TestStringStartingIndex >=
		searchInputParms.TestStringLength {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' Starting Index is invalid!\n"+
			"The '%v' Starting Index is greater than the last index\n"+
			"in the '%v' character array.\n"+
			"%v Last String Index = '%v'.\n"+
			"%v Starting Index = '%v'\n",
			ePrefix.String(),
			searchInputParms.TestStringName,
			searchInputParms.TestStringName,
			searchInputParms.TestStringName,
			searchInputParms.TestStringName,
			searchInputParms.TestStringLength-1,
			searchInputParms.TestStringName,
			searchInputParms.TestStringStartingIndex)

		return err

	}

	return err
}

func (searchInputParms *CharSearchInputParametersDto) ValidateCharSearchType(
	errorPrefix interface{}) error {

	if searchInputParms.lock == nil {
		searchInputParms.lock = new(sync.Mutex)
	}

	searchInputParms.lock.Lock()

	defer searchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchInputParametersDto."+
			"ValidateCharSearchType()",
		"")

	if err != nil {

		return err

	}

	if !searchInputParms.CharSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: The Character Search Type is invalid!\n"+
			"Character Search Type must be set to one of these\n"+
			"enumeration values:\n"+
			"  CharacterSearchType(0).LinearTargetStartingIndex()\n"+
			"  CharacterSearchType(0).SingleTargetChar()\n"+
			"  CharacterSearchType(0).LinearEndOfString()\n"+
			"The invalid Input Character Search Type is currently\n"+
			"configured as:\n"+
			" Character Search Type   String Name: %v\n"+
			" Character Search Type Integer Value: %v\n",
			ePrefix.String(),
			searchInputParms.CharSearchType.String(),
			searchInputParms.CharSearchType.XValueInt())

	}

	return err
}

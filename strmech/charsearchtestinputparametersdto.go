package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type CharSearchTestInputParametersDto struct {
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
	//                                       signs. This is the Primary
	//                                       Type Code for Number Signs.
	//                                       Cases involving 'Leading' and
	//                                       'Trailing' symbols also make
	//                                       use of the 'NumSignSubPosition'.
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

	lock *sync.Mutex
}

func (testSearchInputParms CharSearchTestInputParametersDto) New() CharSearchTestInputParametersDto {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	newEmptyTestInputParms := CharSearchTestInputParametersDto{}

	return newEmptyTestInputParms
}

// ValidateTestParameters - Validates the Test String and related
// member variables contained in the current instance of
// CharSearchInputParametersDto.
//
func (testSearchInputParms *CharSearchTestInputParametersDto) ValidateTestParameters(
	errorPrefix interface{}) error {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

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

	if len(testSearchInputParms.TestStringName) == 0 {
		testSearchInputParms.TestStringName = "TestString"
	}

	if testSearchInputParms.TestString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			testSearchInputParms.TestStringName)

		return err
	}

	if len(testSearchInputParms.TestStringLengthName) == 0 {
		testSearchInputParms.TestStringLengthName =
			"TestStringLengthName"
	}

	testSearchInputParms.TestStringLength =
		len(testSearchInputParms.TestString.CharsArray)

	if testSearchInputParms.TestStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName)

		return err
	}

	if testSearchInputParms.TestStringStartingIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' Starting Index is invalid!\n"+
			"The '%v' Starting Index is less than Zero (0).\n"+
			"%v Starting Index = '%v'.\n",
			ePrefix.String(),
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringStartingIndex)

		return err

	}

	if testSearchInputParms.TestStringStartingIndex >=
		testSearchInputParms.TestStringLength {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' Starting Index is invalid!\n"+
			"The '%v' Starting Index is greater than the last index\n"+
			"in the '%v' character array.\n"+
			"%v Last String Index = '%v'.\n"+
			"%v Starting Index = '%v'\n",
			ePrefix.String(),
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringLength-1,
			testSearchInputParms.TestStringName,
			testSearchInputParms.TestStringStartingIndex)

		return err

	}

	return err
}

func (testSearchInputParms *CharSearchTestInputParametersDto) ValidateCharSearchType(
	errorPrefix interface{}) error {

	if testSearchInputParms.lock == nil {
		testSearchInputParms.lock = new(sync.Mutex)
	}

	testSearchInputParms.lock.Lock()

	defer testSearchInputParms.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"CharSearchTestInputParametersDto."+
			"ValidateCharSearchType()",
		"")

	if err != nil {

		return err

	}

	if !testSearchInputParms.CharSearchType.XIsValid() {

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
			testSearchInputParms.CharSearchType.String(),
			testSearchInputParms.CharSearchType.XValueInt())

	}

	return err

}

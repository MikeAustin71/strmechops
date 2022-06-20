package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoElectron struct {
	lock *sync.Mutex
}

// linearEndOfStringSearch - This low level function executes a
// string search operation proceeds through the entire length of
// Target Search String.
//
// As a low level function, very little validation is performed on
// the input parameters for this function. It is assumed that
// input parameter validation was previously performed by a higher
// level function.
//
func (runeDtoElectron *runeArrayDtoElectron) linearEndOfStringSearch(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	searchResults := CharSearchResultsDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoElectron."+
			"linearEndOfStringSearch()",
		"")

	if err != nil {
		return searchResults, err
	}

	err = targetInputParms.ValidateTargetParameters(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {
		return searchResults, err
	}

	err = testInputParms.ValidateTestParameters(
		ePrefix.XCpy(
			"testInputParms"))

	if err != nil {
		return searchResults, err
	}

	searchResults.LoadTargetBaseInputParameters(
		targetInputParms)

	searchResults.LoadTestBaseInputParameters(
		testInputParms)

	j := 0
	k := 0

	searchResults.TestStringStartingIndex = 0

	for i := searchResults.TargetStringStartingSearchIndex; i < searchResults.TargetStringSearchLength; i++ {

		k = i
		j = 0

		for true {

			if testInputParms.TestString.CharsArray[j] !=
				targetInputParms.TargetString.CharsArray[k] {

				break
			}

			j++

			if j == searchResults.TestStringLength {

				// Search Was SUCCESSFUL!
				// All Test characters found!
				// EXIT HERE!

				searchResults.FoundSearchTarget = true

				searchResults.TestStringStartingIndex = 0

				searchResults.TestStringFirstFoundIndex = 0

				searchResults.TestStringLastFoundIndex =
					searchResults.TestStringLength - 1

				searchResults.TargetStringLastFoundIndex = k

				searchResults.TargetStringFirstFoundIndex =
					searchResults.TargetStringLastFoundIndex -
						searchResults.TestStringLength +
						1

				return searchResults, err
			}

			k++

			if k == searchResults.TargetStringSearchLength {
				break
			}
		}

	}

	return searchResults, err
}

// linearTargetStartingIndexSearch
func (runeDtoElectron *runeArrayDtoElectron) linearTargetStartingIndexSearch(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	searchResults := CharSearchResultsDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoElectron."+
			"linearTargetStartingIndexSearch()",
		"")

	if err != nil {

		return searchResults, err

	}

	err = targetInputParms.ValidateTargetParameters(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {
		return searchResults, err
	}

	err = testInputParms.ValidateTestParameters(
		ePrefix.XCpy(
			"testInputParms"))

	if err != nil {
		return searchResults, err
	}

	searchResults.LoadTargetBaseInputParameters(
		targetInputParms)

	searchResults.LoadTestBaseInputParameters(
		testInputParms)

	j := 0

	for i := searchResults.TargetStringStartingSearchIndex; i < searchResults.TargetStringSearchLength; i++ {

		if testInputParms.TestString.CharsArray[j] !=
			targetInputParms.TargetString.CharsArray[i] {

			searchResults.FoundSearchTarget = false

			// Search Failed. No Match!
			// Exit Here!
			return searchResults, err
		}

		j++

		if j == searchResults.TestStringLength {

			// Search Was SUCCESSFUL!
			// All Test Characters found!
			// EXIT HERE!

			searchResults.FoundSearchTarget = true

			searchResults.TestStringStartingIndex = 0

			searchResults.TestStringFirstFoundIndex = 0

			searchResults.TestStringLastFoundIndex =
				searchResults.TestStringLength - 1

			searchResults.TargetStringLastFoundIndex = i

			searchResults.TargetStringFirstFoundIndex =
				searchResults.TargetStringLastFoundIndex -
					searchResults.TestStringLength +
					1

			return searchResults, err
		}

	}

	return searchResults, err
}

// ptr - Returns a pointer to a new instance of
// runeArrayDtoElectron.
//
func (runeDtoElectron runeArrayDtoElectron) ptr() *runeArrayDtoElectron {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	return &runeArrayDtoElectron{
		lock: new(sync.Mutex),
	}
}

// singleCharacterSearch - Performs a Single Target Character
// Search on the target string (rune array) specified by input
// parameter 'targetInputParms'.
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//
//  Character Search algorithms using the Character Search Type
//  ('CharacterSearchType') rely on a framework consisting of a
//  'Target Search String' and a 'Test String'.
//
//    Target Search String - A string character or characters which
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
//  A comparison of text characters contained in the Target Search
//  String and the Test String serves as the basis for determining
//  a 'Match' condition or successful outcome from a text character
//  search algorithm. The specific criterion for determining a
//  'Match' condition vary between the different Character Search
//  Types.
//
// ----------------------------------------------------------------
//
// A Single Character Search Type means that a single character in
// the Target Search String will be compared to all characters in
// the Test String.
//
// If a single Target String character equals any character in the
// Test String, a 'Match' or successful search outcome will be
// declared.
//
// The search operation is limited to a single designated Target
// Search String character. Each and every one of the Test String
// Characters will be compared to this single designated Target
// String Search Character. The search operation will terminate
// when a matching character is first identified in the Test String
// or when the end of the Test String is encountered.
//
//   Example #1
//                              1         2         3
//             Index  0123456789012345678901234567890
//    Target String: "Hey, Xray-4 is the call sign."
//    Target String Starting Index: 5
//      Test String: "ZFXyURJK"
//
// In this example of a Single Target Character Search, the search
// will begin and end at Target Search String index number 5. Since
// one of the Test String Characters ('X') matches the 'X'
// character at index number 5 in the Target Search String, the
// search operation is classified as a success. A matching
// character was found.
//
//   Example #2
//                              1         2         3
//             Index  0123456789012345678901234567890
//    Target String: "Hey, Xray-4 is the call sign."
//    Target String Starting Index: 0
//      Test String: "ZFXyURJK"
//
// In this second example of a Single Target Character Search, the
// search will begin and end at Target Search String index number
// zero (0). Since NONE of the Test String Characters matches the
// 'H' character at index number zero (0) in the Target Search
// String, the search operation is classified as a failure. No
// matching character was found.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms           CharSearchTargetInputParametersDto
//     - This Dto or Data Transfer Object, contains the input
//       parameters and specifications associated with the Target
//       Search String which will be used in completing the Single
//       Target Character Search operation.
//
//       The CharSearchTargetInputParametersDto structure consists
//       of the following member variable data elements.
//
//       type CharSearchTargetInputParametersDto struct {
//
//       	TargetString                        *RuneArrayDto
//            - A pointer to the RuneArrayDto containing the Target
//              Search String text characters used in the search
//              algorithm. Target Characters are compared against
//              Test Characters to determine if a 'Match' condition
//              exists.
//
//       	TargetStringName                    string
//            - The label or name of the 'TargetString' parameter.
//              Used in error and informational messages.
//
//       	TargetStringLength                  int
//            - Actual number of text characters in the entire
//              Target Search String ('TargetString').
//
//       	TargetStringLengthName              string
//            - The label or name of the 'TargetStringLength' parameter.
//              Used in error and informational messages.
//
//       	TargetStringStartingSearchIndex     int
//            - The index in 'TargetString' at which the search
//              operation begins.
//
//       	TargetStringStartingSearchIndexName string
//            - The label or name of the
//              TargetStringStartingSearchIndex parameter.
//              Used in error and informational messages.
//
//       	TargetStringSearchLength            int
//            - The actual number of characters within the Target
//              Search String that are included in the search
//              operation. This value may be less than the actual
//              length of the Target Search String.
//
//       	TargetStringSearchLengthName        string
//            - The label or name of the TargetStringSearchLength
//              parameter. Used in error and informational
//              messages.
//
//       	TargetStringAdjustedSearchLength    int
//            - The adjusted or corrected Target String Search
//              Length. This value is guaranteed to be equal to or
//              less than the actual Target String Length.
//
//       	TargetStringDescription1            string
//            - First of two optional description strings
//              describing the Target Search String in the context
//              of the current search operation.
//
//       	TargetStringDescription2            string
//            - Second of two optional description strings
//              describing the Target Search String in the context
//              of the current search operation.
//
//       	FoundFirstNumericDigitInNumStr      bool
//            - When set to 'true' this signals that the first
//              numeric digit has been identified in the text
//              characters specified by 'TargetString'
//
//
//       	TextCharSearchType                  CharacterSearchType
//            - Optional. An enumeration value signaling the type
//              of text character search algorithm used to conduct
//              this search operation. When set to a valid value,
//              this specification will override the search
//              specification contained in the Test Input
//              Parameters Data Transfer Object.
//
//              Valid CharSearch Type values are listed as follows:
//       	     TextCharSearchType.None() - Invalid Value
//       	     TextCharSearchType.LinearTargetStartingIndex() - Default
//       	     TextCharSearchType.SingleTargetChar()
//       	     TextCharSearchType.LinearEndOfString()
//
//       }
//
//
//  testInputParms             CharSearchTestInputParametersDto
//     - This Dto or Data Transfer Object, contains the input
//       parameters and specifications associated with the Test
//       String which will be used in completing the Single Target
//       Character Search operation.
//
//       The CharSearchTestInputParametersDto structure consists
//       of the following member variable data elements.
//
//       type CharSearchTestInputParametersDto struct {
//       	TestString                          *RuneArrayDto
//            - A pointer to the Rune Array containing the Test
//              Characters to be used in a text character search
//              algorithm. Target Characters are compared against
//              Test Characters to determine if a 'Match' condition
//              exists.
//
//       	TestStringName                      string
//            - The label or name of the 'TestString' parameter.
//              Used in error and informational messages.
//
//       	TestStringLength                    int
//            - Actual number of text characters in the entire Test
//              String ('TestString').
//
//       	TestStringLengthName                string
//            - The label or name of the 'TestStringLength'
//              parameter. Used in error and informational
//              messages.
//
//       	TestStringStartingIndex             int
//            - The starting index in the Test String where the
//              search operation will begin.
//
//       	TestStringDescription1              string
//            - First of two optional description strings
//              describing the Test String in the context of the
//              current search operation.
//
//       	TestStringDescription2              string
//            - Second of two optional description strings
//              describing the Test String in the context of the
//              current search operation.
//
//       	CollectionTestObjIndex              int
//            - If the Test String object resides in a collection
//              of Test String objects, this parameter will record
//              the array index of the current Test String object
//              in the collection.
//
//       	NumValueType                        NumericValueType
//            - Optional. This enumeration value specifies the type
//              of numeric value for this Test Parameter as either
//              an integer or floating point value.
//
//              Possible values are listed as follows:
//       	     NumValType.None()
//       	     NumValType.FloatingPoint()
//       	     NumValType.Integer()
//
//       	NumStrFormatType                    NumStrFormatTypeCode
//            - Optional. This enumeration value specifies the
//              Output Format Type for a number.
//
//              Possible values are listed as follows:
//       	     NumStrFmtType.None()
//       	     NumStrFmtType.AbsoluteValue()
//       	     NumStrFmtType.Binary()
//       	     NumStrFmtType.CountryCulture()
//       	     NumStrFmtType.Currency()
//       	     NumStrFmtType.Binary()
//       	     NumStrFmtType.Hexadecimal()
//       	     NumStrFmtType.Octal()
//       	     NumStrFmtType.ScientificNotation()
//
//       	NumSymLocation                      NumericSymbolLocation
//            - Optional. This enumeration value specifies the
//              relative location of a numeric symbol.
//
//              Possible values are listed as follows:
//       	     NumSymLocation.None()
//       	     NumSymLocation.Before()
//       	     NumSymLocation.Interior()
//       	     NumSymLocation.After()
//
//       	NumSymbolClass                      NumericSymbolClass
//            - Optional. This enumeration value specifies the
//              Number Symbol Classification for a text character.
//
//              Possible values are listed as follows:
//       	     NumSymClass.None()
//       	     NumSymClass.NumberSign()
//       	     NumSymClass.CurrencySign()
//       	     NumSymClass.IntegerSeparator()
//       	     NumSymClass.DecimalSeparator()
//
//       	NumSignValue                        NumericSignValueType
//            - Optional. This enumeration value specifies the
//              number sign value.
//
//              Possible values are listed as follows:
//       	     NumSignVal.None()
//       	     NumSignVal.Negative()
//       	     NumSignVal.Zero()
//       	     NumSignVal.Positive()
//
//       	PrimaryNumSignPosition              NumSignSymbolPosition
//            - Optional This enumeration value specifies the
//              relative position of positive and negative number
//              sign symbols in a number string. This is the
//              Primary Type Code for Number Signs. Cases involving
//              both 'Leading' and 'Trailing' symbols also make use
//              of the 'SecondaryNumSignPosition'.
//
//              Possible values are listed as follows:
//       	     NumSignSymPos.None()
//       	     NumSignSymPos.Before()
//       	     NumSignSymPos.After()
//       	     NumSignSymPos.BeforeAndAfter()
//
//       	SecondaryNumSignPosition            NumSignSymbolPosition
//            - Optional. This enumeration value specifies the
//              relative position of positive and negative number
//              sign symbols in a number string. This value is used
//              in searches involving number signs which occur both
//              before and after the numeric value.
//
//              Possible values are listed as follows:
//       	     NumSignSymPos.None()
//       	     NumSignSymPos.Before()
//       	     NumSignSymPos.After()
//       	     NumSignSymPos.BeforeAndAfter()
//
//       	TextCharSearchType                  CharacterSearchType
//            - Required. An enumeration value signaling the type
//              of text character search algorithm used to conduct
//              this search operation.
//
//              Possible values are listed as follows:
//       	     TextCharSearchType.None() - Invalid value
//       	     TextCharSearchType.LinearTargetStartingIndex() - Default
//       	     TextCharSearchType.SingleTargetChar()
//       	     TextCharSearchType.LinearEndOfString()
//       }
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  CharSearchResultsDto
//     - If this method completes successfully, a populated
//       instance of this parameter will be returned to the calling
//       function. This Dto or Data Transfer Object, contains the
//       results of the text character search operation performed
//       by this method.
//
//       The CharSearchResultsDto structure consists
//       of the following member variable data elements.
//
//       type CharSearchResultsDto struct {
//
//       	FoundSearchTarget                   bool
//            - Signals a successful search outcome. If set to
//              'true' the Test String character or characters were
//              found in the Target Search String.
//
//       	FoundSearchTargetOnPreviousSearch   bool
//            - The Search Target was located in a previous search
//              operation.
//
//       	TargetStringLength                  int
//            - Actual number of text characters in the entire
//              Target Search String ('TargetString').
//
//       	TargetStringSearchLength            int
//            - The actual number of characters within the Target
//              Search String that are included in the search
//              operation. This value may be less than the actual
//              length of the Target Search String.
//
//       	TargetStringStartingSearchIndex     int
//            - The index in 'TargetString' at which the search
//              operation begins.
//
//       	TargetStringFirstFoundIndex         int
//            - The index of the first character position in the
//              Target Search String occupied by the first
//              character in the Test String.
//
//       	TargetStringLastFoundIndex          int
//            - The last character position in the Target Search
//              String occupied by the last character in the Test
//       	    String.
//
//       	TargetStringLastSearchIndex int
//      	  - The index in Target Search String occupied by the
//        	    last Target character searched. If the Search
//       	    Target was found, this value is equal to the
//      	    'TargetStringLastFoundIndex'. If the Search Target
//     	        was NOT found this value is equal to the
//    	        'TargetStringStartingSearchIndex'. This value is
//   	        useful in computing the next index to be searched
//  	        in the Target String/
//
//       	TargetStringDescription1            string
//            - First of two optional description strings
//              describing the Target Search String in the context
//              of the current search operation.
//
//       	TargetStringDescription2            string
//            - Second of two optional description strings
//              describing the Target Search String in the context
//              of the current search operation.
//
//       	TestStringLength                    int
//            - Actual number of text characters in the entire Test
//              String ('TestString').
//
//       	TestStringStartingIndex             int
//            - The starting index in the Test String where the
//              search operation will begin.
//
//       	TestStringFirstFoundIndex           int
//      	  - The index number in Test String of the first test
//              character to be located in the Target Search String.
//
//       	TestStringLastFoundIndex            int
//            - The index number in the Test String occupied by the
//              last Test Character to be located in the Target
//              String.
//
//       	TestStringDescription1              string
//            - First of two optional description strings
//              describing the Test String in the context of the
//              current search operation.
//
//       	TestStringDescription2              string
//            - Second of two optional description strings
//              describing the Test String in the context of the
//              current search operation.
//
//       	CollectionTestObjIndex              int
//            - If the Test String object resides in a collection
//              of Test String objects, this parameter will record
//              the array index of the current Test String object
//              in the collection.
//
//       	NumValueType                        NumericValueType
//            - Optional. This enumeration value specifies the type
//              of numeric value for this Test Parameter as either
//              an integer or floating point value.
//
//              Possible values are listed as follows:
//       	     NumValType.None()
//       	     NumValType.FloatingPoint()
//       	     NumValType.Integer()
//
//       	NumStrFormatType                    NumStrFormatTypeCode
//            - Optional. This enumeration value specifies the
//              Output Format Type for a number.
//
//              Possible values are listed as follows:
//       	     NumStrFmtType.None()
//       	     NumStrFmtType.AbsoluteValue()
//       	     NumStrFmtType.Binary()
//       	     NumStrFmtType.CountryCulture()
//       	     NumStrFmtType.Currency()
//       	     NumStrFmtType.Binary()
//       	     NumStrFmtType.Hexadecimal()
//       	     NumStrFmtType.Octal()
//       	     NumStrFmtType.ScientificNotation()
//
//       	NumSymLocation                      NumericSymbolLocation
//            - Optional. This enumeration value specifies the
//              relative location of a numeric symbol.
//
//              Possible values are listed as follows:
//       	     NumSymLocation.None()
//       	     NumSymLocation.Before()
//       	     NumSymLocation.Interior()
//       	     NumSymLocation.After()
//
//       	NumSymbolClass                      NumericSymbolClass
//            - Optional. This enumeration value specifies the
//              Number Symbol Classification for a text character.
//
//              Possible values are listed as follows:
//       	     NumSymClass.None()
//       	     NumSymClass.NumberSign()
//       	     NumSymClass.CurrencySign()
//       	     NumSymClass.IntegerSeparator()
//       	     NumSymClass.DecimalSeparator()
//
//       	NumSignValue                        NumericSignValueType
//            - Optional. This enumeration value specifies the
//              number sign value.
//
//              Possible values are listed as follows:
//       	     NumSignVal.None()
//       	     NumSignVal.Negative()
//       	     NumSignVal.Zero()
//       	     NumSignVal.Positive()
//
//       	PrimaryNumSignPosition              NumSignSymbolPosition
//            - Optional This enumeration value specifies the
//              relative position of positive and negative number
//              sign symbols in a number string. This is the
//              Primary Type Code for Number Signs. Cases involving
//              both 'Leading' and 'Trailing' symbols also make use
//              of the 'SecondaryNumSignPosition'.
//
//              Possible values are listed as follows:
//       	     NumSignSymPos.None()
//       	     NumSignSymPos.Before()
//       	     NumSignSymPos.After()
//       	     NumSignSymPos.BeforeAndAfter()
//
//       	SecondaryNumSignPosition            NumSignSymbolPosition
//            - Optional. This enumeration value specifies the
//              relative position of positive and negative number
//              sign symbols in a number string. This value is used
//              in searches involving number signs which occur both
//              before and after the numeric value.
//
//              Possible values are listed as follows:
//       	     NumSignSymPos.None()
//       	     NumSignSymPos.Before()
//       	     NumSignSymPos.After()
//       	     NumSignSymPos.BeforeAndAfter()
//
//       	TextCharSearchType                  CharacterSearchType
//            - Required. An enumeration value signaling the type
//              of text character search algorithm used to conduct
//              this search operation.
//
//              Possible values are listed as follows:
//       	     TextCharSearchType.None() - Invalid value
//       	     TextCharSearchType.LinearTargetStartingIndex() - Default
//       	     TextCharSearchType.SingleTargetChar()
//       	     TextCharSearchType.LinearEndOfString()
//       }
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
func (runeDtoElectron *runeArrayDtoElectron) singleCharacterSearch(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	searchResults := CharSearchResultsDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoElectron."+
			"singleCharacterSearch()",
		"")

	err = targetInputParms.ValidateTargetParameters(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {
		return searchResults, err
	}

	err = testInputParms.ValidateTestParameters(
		ePrefix.XCpy(
			"testInputParms"))

	if err != nil {
		return searchResults, err
	}

	searchResults.LoadTargetBaseInputParameters(
		targetInputParms)

	searchResults.LoadTestBaseInputParameters(
		testInputParms)

	targetChar :=
		targetInputParms.TargetString.CharsArray[targetInputParms.TargetStringStartingSearchIndex]

	for j := testInputParms.TestStringStartingIndex; j < testInputParms.TestStringLength; j++ {

		if testInputParms.TestString.CharsArray[j] == targetChar {
			// Search SUCCESSFUL! SINGLE CHARACTER MATCH!
			// Exit Here!

			searchResults.FoundSearchTarget = true

			searchResults.TestStringStartingIndex =
				testInputParms.TestStringStartingIndex

			searchResults.TestStringFirstFoundIndex = j

			searchResults.TestStringLastFoundIndex = j

			searchResults.TargetStringLastFoundIndex =
				targetInputParms.TargetStringStartingSearchIndex

			searchResults.TargetStringFirstFoundIndex =
				targetInputParms.TargetStringStartingSearchIndex

			return searchResults, err
		}
	}

	return searchResults, err
}

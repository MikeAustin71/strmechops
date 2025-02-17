package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoElectron struct {
	lock *sync.Mutex
}

// emptyCharsArray - Empties or deletes all the characters residing
// in the internal member variable 'CharsArray' contained in the
// instance of RuneArrayDto passed as input parameter
// 'runeArrayDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto               *RuneArrayDto
//	   - A pointer to an instance of RuneArrayDto. All
//	     the characters contained in the internal member rune array
//	     'CharsArray' will be deleted. The internal member rune
//	     array will then be set to a value of 'nil'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (runeDtoElectron *runeArrayDtoElectron) emptyCharsArray(
	runeArrayDto *RuneArrayDto) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	if runeArrayDto == nil {
		return
	}

	runeArrayDto.CharsArray = nil

	return
}

// equalCharArrays - Receives pointers to two instances of
// RuneArrayDto and proceeds to compare the internal rune arrays
// contained in the two instances.
//
// If the internal member variables 'CharsArray' presented by
// the two instances of RuneArrayDto are equal in all respects,
// this method will return a boolean value of 'true'.
//
// If the two rune arrays are unequal, this method will return a
// boolean value of 'false'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto1              *RuneArrayDto
//	   - A pointer to first of two instances of RuneArrayDto. The
//	     text characters contained in internal member variable
//	     rune array, 'CharsArray', will be compared to those
//	     contained in the second instance of RuneArrayDto,
//	     'runeArrayDto2'. If they are equal in all respects,
//	     this method will return a boolean value of 'true'.
//
//
//	runeArrayDto2              *RuneArrayDto
//	   - A pointer to the second of two instances of RuneArrayDto.
//	     The text characters contained in internal member variable
//	     rune array, 'CharsArray', will be compared to those
//	     contained in the first instance of RuneArrayDto,
//	     'runeArrayDto1'. If they are equal in all respects,
//	     this method will return a boolean value of 'true'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - This boolean flag will transmit the results of a
//	     comparison between the character rune arrays contained in
//	     input parameters 'runeArrayDto1' and 'runeArrayDto2'. If
//	     the comparison shows that the member variables
//	     'CharsArray' are equivalent in all respects, this return
//	     value will be set to 'true'.
//
//	     If the member variable 'CharsArray' differs in any way
//	     between instances 'runeArrayDto1' and 'runeArrayDto2',
//	     this will return 'false'.
func (runeDtoElectron *runeArrayDtoElectron) equalCharArrays(
	runeArrayDto1 *RuneArrayDto,
	runeArrayDto2 *RuneArrayDto) bool {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	if runeArrayDto1 == nil ||
		runeArrayDto2 == nil {

		return false
	}

	lenOfCharsArray := len(runeArrayDto1.CharsArray)

	if lenOfCharsArray !=
		len(runeArrayDto2.CharsArray) {

		return false
	}

	// Length of Character Arrays
	// MUST be equal
	if lenOfCharsArray == 0 {
		return true
	}

	for i := 0; i < lenOfCharsArray; i++ {

		if runeArrayDto1.CharsArray[i] !=
			runeArrayDto2.CharsArray[i] {
			return false
		}

	}

	return true
}

// linearEndOfStringSearch - Performs a Linear End Of String Search
// for Text Characters. This means that each character in the
// Target Search String will be compared to each corresponding
// character in the Test String beginning at a specified starting
// index in the Target Search String and proceeding through the
// entire length of the Target Search String.
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
// ('CharacterSearchType') rely on a framework consisting of a
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
// Character Search Type is a series of enumeration values
// specifying the type of text character search algorithm applied
// by the current instance of RuneArrayDto.
//
// For this method, the Character Search Type must be set to one of
// the following enumeration values:
//
//	TextCharSearchType.LinearTargetStartingIndex()
//	TextCharSearchType.SingleTargetChar()
//	TextCharSearchType.LinearEndOfString()
//
// ----------------------------------------------------------------
//
// With the Linear End Of String Search algorithm, the entire
// Target Search String will be searched from left to right for the
// first occurrence of the Test String. Other search algorithms
// terminate after searching a designated subset of the Target
// Search String
//
// The search will begin the Target String Starting Index and
// proceed left to right until (1) an instance of the entire Test
// String is located or (2) the end of the Target Search String is
// encountered.
//
// This is a linear search, so a 'Match' requires that each
// character in Target Search String must correspond to a matching
// character in the Test String.
//
//	Example
//	                           1         2         3
//	          Index  0123456789012345678901234567890
//	 Target String: "Hey, Xray-4 is the call sign."
//	 Target String Starting Index: 0
//	   Test String: "Xray-4"
//
// In this example of a Linear End of String Search, the search
// operation will begin comparing corresponding characters in the
// Target Search String and the Test String beginning at index
// zero. The comparison will fail at index zero, but the search
// algorithm will continue attempting to find the Test String at
// indexes 1,2, 3 & 4. The Test String will be found beginning at
// index number 5 and the search algorithm will terminate at that
// point with a successful outcome or 'Match' result.
//
// -----------------------------------------------------------------
//
// # Be Advised
//
// This method is a low level function. As such, very little
// validation is performed on the input parameters for this
// function. It is assumed that input parameter validation was
// previously performed by a higher level function.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	targetInputParms           CharSearchTargetInputParametersDto
//	   - This Dto, or Data Transfer Object, contains the input
//	     parameters and specifications associated with the Target
//	     Search String which will be used in completing the Linear
//	     Target Starting Index Search operation.
//
//	     The CharSearchTargetInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchTargetInputParametersDto struct {
//
//	      TargetInputParametersName           string
//	       - The Name, Label or descriptive Tag associated with this
//	         instance of CharSearchTargetInputParametersDto. If empty,
//	         this string will be defaulted to "TargetInputParameters"
//
//	      TargetString                        *RuneArrayDto
//	       - A pointer to the RuneArrayDto containing the Target
//	         Search String text characters used in the search
//	         algorithm. Target Characters are compared against
//	         Test Characters to determine if a 'Match' condition
//	         exists.
//
//	      TargetStringName                    string
//	       - The label or name of the 'TargetString' parameter.
//	         Used in error and informational messages.
//
//	      TargetStringLength                  int
//	       - Actual number of text characters in the entire
//	         Target Search String ('TargetString').
//
//	      TargetStringLengthName              string
//	       - The label or name of the 'TargetStringLength' parameter.
//	         Used in error and informational messages.
//
//	      TargetStringStartingSearchIndex     int
//	       - The index in 'TargetString' at which the search
//	         operation begins.
//
//	      TargetStringStartingSearchIndexName string
//	       - The label or name of the
//	         TargetStringStartingSearchIndex parameter.
//	         Used in error and informational messages.
//
//	      TargetStringSearchLength            int
//	       - The actual number of characters within the Target
//	         Search String that are included in the search
//	         operation. This value may be less than the actual
//	         length of the Target Search String.
//
//	      TargetStringSearchLengthName        string
//	       - The label or name of the TargetStringSearchLength
//	         parameter. Used in error and informational
//	         messages.
//
//	      TargetStringAdjustedSearchLength    int
//	       - The adjusted or corrected Target String Search
//	         Length. This value is guaranteed to be equal to or
//	         less than the actual Target String Length.
//
//	      TargetStringDescription1            string
//	       - First of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      TargetStringDescription2            string
//	       - Second of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      FoundFirstNumericDigitInNumStr      bool
//	       - When set to 'true' this signals that the first
//	         numeric digit has been identified in the text
//	         characters specified by 'TargetString'
//
//	      CollectionTestObjIndex int
//	       - If the Test String object resides in a collection
//	         of Test String objects, this parameter will record
//	         the array index of the current Test String object
//	         in the collection.
//
//	      TextCharSearchType                  CharacterSearchType
//	       - Optional. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation. When set to a valid value,
//	         this specification will override the search
//	         specification contained in the Test Input
//	         Parameters Data Transfer Object.
//
//	        Valid CharSearch Type values are listed as follows:
//	        TextCharSearchType.None() - Invalid Value
//	        TextCharSearchType.LinearTargetStartingIndex() - Default
//	        TextCharSearchType.SingleTargetChar()
//	        TextCharSearchType.LinearEndOfString()
//	     }
//
//
//	testInputParms      CharSearchTestInputParametersDto
//	   - This Dto, or Data Transfer Object, contains the input
//	     parameters and specifications associated with the Test
//	     String which will be used in completing the Linear Target
//	     Starting Index Search operation. This algorithm compares
//	     text characters in the Target Search String
//	     ('targetInputParms') with those contained in the Test
//	     String specified below.
//
//	     The CharSearchTestInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchTestInputParametersDto struct {
//
//	       TestInputParametersName string
//	        - The Name, Label or descriptive Tag associated with this
//	          instance of CharSearchTestInputParametersDto. If empty,
//	          this string will be defaulted to "TestInputParameters"
//
//	       TestString *RuneArrayDto
//	        - A pointer to the Rune Array Data Transfer
//	          Object containing the Test Characters to be
//	          used in a text character search algorithm.
//	          Target Characters are compared against Test
//	          Characters to determine if a 'Match' condition
//	          exists.
//
//	       TestStringName string
//	        - The label or name of the 'TestString' parameter.
//	          Used in error and informational messages.
//
//	       TestStringLength int
//	        - Actual number of text characters in the entire Test
//	          String ('TestString').
//
//	       TestStringLengthName string
//	        - The label or name of the 'TestStringLength'
//	          parameter. Used in error and informational
//	          messages.
//
//	       TestStringStartingIndex int
//	        - The starting index in the Test String where the
//	          search operation will begin.
//
//	       TestStringStartingIndexName string
//	        - The label or name of the TestStringStartingIndex
//	          parameter. Used in error and informational messages.
//
//	       TestStringDescription1 string
//	        - First of two optional description strings
//	          describing the Test String in the context of the
//	          current search operation.
//
//	       TestStringDescription2 string
//	        - Second of two optional description strings
//	          describing the Test String in the context of the
//	          current search operation.
//
//	       CollectionTestObjIndex int
//	        - If the Test String object resides in a collection
//	          of Test String objects, this parameter will record
//	          the array index of the current Test String object
//	          in the collection.
//
//	       NumValueType NumericValueType
//	        - Optional. This enumeration value specifies the type
//	          of numeric value for this Test Parameter as either
//	          an integer or floating point value.
//
//	          Possible values are listed as follows:
//	           NumValType.None()
//	           NumValType.FloatingPoint()
//	           NumValType.Integer()
//
//	       NumStrFormatType NumStrFormatTypeCode
//	        - Optional. This enumeration value specifies the
//	          Output Format Type for a number.
//
//	          Possible values are listed as follows:
//	           NumStrFmtType.None()
//	           NumStrFmtType.AbsoluteValue()
//	           NumStrFmtType.Binary()
//	           NumStrFmtType.CountryCulture()
//	           NumStrFmtType.Currency()
//	           NumStrFmtType.Binary()
//	           NumStrFmtType.Hexadecimal()
//	           NumStrFmtType.Octal()
//	           NumStrFmtType.ScientificNotation()
//
//	       NumSymLocation NumericSymbolLocation
//	        - Optional. This enumeration value specifies the
//	          relative location of a numeric symbol.
//
//	          Possible values are listed as follows:
//	           NumSymLocation.None()
//	           NumSymLocation.Before()
//	           NumSymLocation.Interior()
//	           NumSymLocation.After()
//
//	       NumSymbolClass NumericSymbolClass
//	        - Optional. This enumeration value specifies the
//	          Number Symbol Classification for a text character.
//
//	          Possible values are listed as follows:
//	           NumSymClass.None()
//	           NumSymClass.NumberSign()
//	           NumSymClass.CurrencySign()
//	           NumSymClass.IntegerSeparator()
//	           NumSymClass.DecimalSeparator()
//
//	       NumSignValue NumericSignValueType
//	        - Optional. This enumeration value specifies the
//	          number sign value.
//
//	          Possible values are listed as follows:
//	           NumSignVal.None()
//	           NumSignVal.Negative()
//	           NumSignVal.Zero()
//	           NumSignVal.Positive()
//
//	       PrimaryNumSignPosition NumSignSymbolPosition
//	       - Optional This enumeration value specifies the
//	         relative position of positive and negative number
//	         sign symbols in a number string. This is the
//	         Primary Type Code for Number Signs. Cases involving
//	         both 'Leading' and 'Trailing' symbols also make use
//	         of the 'SecondaryNumSignPosition'.
//
//	         Possible values are listed as follows:
//	          NumSignSymPos.None()
//	          NumSignSymPos.Before()
//	          NumSignSymPos.After()
//	          NumSignSymPos.BeforeAndAfter()
//
//	       SecondaryNumSignPosition NumSignSymbolPosition
//	        - Optional. This enumeration value specifies the
//	          relative position of positive and negative number
//	          sign symbols in a number string. This value is used
//	          in searches involving number signs which occur both
//	          before and after the numeric value.
//
//	          Possible values are listed as follows:
//	           NumSignSymPos.None()
//	           NumSignSymPos.Before()
//	           NumSignSymPos.After()
//	           NumSignSymPos.BeforeAndAfter()
//
//	       TextCharSearchType CharacterSearchType
//	       - Required. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation.
//
//	         Possible values are listed as follows:
//	          CharSearchType.None()                      - Invalid Value
//	          CharSearchType.LinearTargetStartingIndex() - Valid Default
//	          CharSearchType.SingleTargetChar()          - Valid
//	          CharSearchType.LinearEndOfString()         - Valid
//
//	     }
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	CharSearchRuneArrayResultsDto
//	   - If this method completes successfully, a fully populated
//	     instance of CharSearchRuneArrayResultsDto will be returned to the
//	     calling function.
//
//	     This structure contains information and data elements
//	     detailing the results of the text character search
//	     operation performed by this method.
//
//
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
func (runeDtoElectron *runeArrayDtoElectron) linearEndOfStringSearch(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchRuneArrayResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	searchResults := CharSearchRuneArrayResultsDto{}.New()

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

	searchResults.SearchResultsFunctionChain =
		ePrefix.String()

	j := 0
	k := 0

	searchResults.TestStringStartingIndex = 0

	for i := targetInputParms.TargetStringCurrentSearchIndex; i < targetInputParms.TargetStringLength; i++ {

		k = i
		j = 0

		for {

			if testInputParms.TestString.CharsArray[j] !=
				targetInputParms.TargetString.CharsArray[k] {

				// This is an End Of String Search
				// Keep looking for a match until
				// End-Of-String
				break
			}

			j++

			if j == searchResults.TestStringLength {

				// Search Was SUCCESSFUL!
				// All Test characters found!
				// EXIT HERE!

				searchResults.FoundSearchTarget = true

				// Set Test String Parameters
				searchResults.TestStringFirstFoundIndex = 0

				searchResults.TestStringLastFoundIndex =
					searchResults.TestStringLength - 1

				// Set Target String Parameters
				searchResults.TargetStringLastFoundIndex = k

				searchResults.TargetStringLastSearchIndex = k

				searchResults.TargetStringFirstFoundIndex =
					searchResults.TargetStringLastFoundIndex -
						searchResults.TestStringLength +
						1

				searchResults.TargetStringCurrentSearchIndex = k

				searchResults.TargetStringNextSearchIndex =
					k + 1

				if searchResults.TargetStringNextSearchIndex >=
					targetInputParms.TargetStringLength {

					searchResults.TargetStringNextSearchIndex =
						-1
				}

				if testInputParms.RequestFoundTestCharacters {

					err = searchResults.FoundRuneArrayChars.CopyIn(
						testInputParms.TestString,
						ePrefix.XCpy(
							"searchResults.FoundRuneArrayChars<-"+
								"testInputParms.TestString"))

					if err != nil {
						return searchResults, err
					}
				}

				if testInputParms.RequestRemainderString {

					if searchResults.TargetStringNextSearchIndex >
						-1 {

						searchResults.RemainderString.Empty()

						searchResults.RemainderString.CharsArray =
							append(
								searchResults.RemainderString.CharsArray,
								targetInputParms.TargetString.CharsArray[searchResults.TargetStringCurrentSearchIndex:]...)
					}
				}

				return searchResults, err
			}

			k++

			if k == searchResults.TargetStringSearchLength {
				break
			}
		}

	}

	// Search Failed!
	// No MATCHING Text Characters were
	// Found.
	return searchResults, err
}

// linearTargetStartingIndexSearch - Performs a Linear Target
// Starting Index Search for Text Characters. This means that each
// character in the Target Search String will be compared to each
// corresponding character in the Test String beginning at a
// specified starting index in the Target Search String.
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
// ('CharacterSearchType') rely on a framework consisting of a
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
// Character Search Type is a series of enumeration values
// specifying the type of text character search algorithm applied
// by the current instance of RuneArrayDto.
//
// For this method, the Character Search Type must be set to one of
// the following enumeration values:
//
//	TextCharSearchType.LinearTargetStartingIndex()
//	TextCharSearchType.SingleTargetChar()
//	TextCharSearchType.LinearEndOfString()
//
// ----------------------------------------------------------------
//
// The Linear Starting Index Text Search will proceed for from left
// to right in Test Character Sequence.
//
// If the Test Characters are NOT found in the Target Search
// String beginning at the designated Target String Starting
// Index, the search outcome will be unsuccessful, NO match
// will be declared and the search operation will terminate.
//
// A 'Match', or successful search outcome, is defined as the
// case where each character in the Target String matches each
// corresponding character in the Test String beginning at the
// designated Target String Starting Index.
//
//	Example 1
//	                          1         2         3
//	         Index  0123456789012345678901234567890
//	Target String: "Hey, Xray-4 is the call sign."
//	Target String Starting Index: 5
//	  Test String: "Xray"
//
// In this example of a Linear Target Starting Index Search, a
// match between the Target String and Test String will be
// declared, if and only if, the search begins at Target String
// index number 5. If the search begins at an any index other
// than 5, no match will be declared and the search will be
// classified as unsuccessful.
//
//	Example 2
//	                          1         2         3
//	         Index  0123456789012345678901234567890
//	Target String: "Hey, Xray-4 is the call sign."
//	Target String Starting Index: 0
//	  Test String: "Xray"
//
// In this second example of a Linear Target Starting Index
// Search, the search operation will fail, no match will be
// found and the search operation will terminate. This result
// follows because the character at staring index number zero
// (0) 'H', does NOT match the first character in the Test
// String 'Xray'. The search operation will therefore terminate
// immediately after the comparison between first characters
// 'H' and 'X' fails to produce a match.
//
// -----------------------------------------------------------------
//
// # Be Advised
//
// This method is a low level function. As such, very little
// validation is performed on the input parameters for this
// function. It is assumed that input parameter validation was
// previously performed by a higher level function.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	targetInputParms           CharSearchTargetInputParametersDto
//	   - This Dto, or Data Transfer Object, contains the input
//	     parameters and specifications associated with the Target
//	     Search String which will be used in completing the Linear
//	     Target Starting Index Search operation.
//
//	     The CharSearchTargetInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchTargetInputParametersDto struct {
//
//	      TargetInputParametersName           string
//	       - The Name, Label or descriptive Tag associated with this
//	         instance of CharSearchTargetInputParametersDto. If empty,
//	         this string will be defaulted to "TargetInputParameters"
//
//	      TargetString                        *RuneArrayDto
//	       - A pointer to the RuneArrayDto containing the Target
//	         Search String text characters used in the search
//	         algorithm. Target Characters are compared against
//	         Test Characters to determine if a 'Match' condition
//	         exists.
//
//	      TargetStringName                    string
//	       - The label or name of the 'TargetString' parameter.
//	         Used in error and informational messages.
//
//	      TargetStringLength                  int
//	       - Actual number of text characters in the entire
//	         Target Search String ('TargetString').
//
//	      TargetStringLengthName              string
//	       - The label or name of the 'TargetStringLength' parameter.
//	         Used in error and informational messages.
//
//	      TargetStringStartingSearchIndex     int
//	       - The index in 'TargetString' at which the search
//	         operation begins.
//
//	      TargetStringStartingSearchIndexName string
//	       - The label or name of the
//	         TargetStringStartingSearchIndex parameter.
//	         Used in error and informational messages.
//
//	      TargetStringSearchLength            int
//	       - The actual number of characters within the Target
//	         Search String that are included in the search
//	         operation. This value may be less than the actual
//	         length of the Target Search String.
//
//	      TargetStringSearchLengthName        string
//	       - The label or name of the TargetStringSearchLength
//	         parameter. Used in error and informational
//	         messages.
//
//	      TargetStringAdjustedSearchLength    int
//	       - The adjusted or corrected Target String Search
//	         Length. This value is guaranteed to be equal to or
//	         less than the actual Target String Length.
//
//	      TargetStringDescription1            string
//	       - First of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      TargetStringDescription2            string
//	       - Second of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      FoundFirstNumericDigitInNumStr      bool
//	       - When set to 'true' this signals that the first
//	         numeric digit has been identified in the text
//	         characters specified by 'TargetString'
//
//	      CollectionTestObjIndex int
//	       - If the Test String object resides in a collection
//	         of Test String objects, this parameter will record
//	         the array index of the current Test String object
//	         in the collection.
//
//	      TextCharSearchType                  CharacterSearchType
//	       - Optional. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation. When set to a valid value,
//	         this specification will override the search
//	         specification contained in the Test Input
//	         Parameters Data Transfer Object.
//
//	        Valid CharSearch Type values are listed as follows:
//	        TextCharSearchType.None() - Invalid Value
//	        TextCharSearchType.LinearTargetStartingIndex() - Default
//	        TextCharSearchType.SingleTargetChar()
//	        TextCharSearchType.LinearEndOfString()
//	     }
//
//
//	testInputParms      CharSearchTestInputParametersDto
//	   - This Dto, or Data Transfer Object, contains the input
//	     parameters and specifications associated with the Test
//	     String which will be used in completing the Linear Target
//	     Starting Index Search operation. This algorithm compares
//	     text characters in the Target Search String
//	     ('targetInputParms') with those contained in the Test
//	     String specified below.
//
//	     The CharSearchTestInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchTestInputParametersDto struct {
//
//	       TestInputParametersName string
//	        - The Name, Label or descriptive Tag associated with this
//	          instance of CharSearchTestInputParametersDto. If empty,
//	          this string will be defaulted to "TestInputParameters"
//
//	       TestString *RuneArrayDto
//	        - A pointer to the Rune Array Data Transfer
//	          Object containing the Test Characters to be
//	          used in a text character search algorithm.
//	          Target Characters are compared against Test
//	          Characters to determine if a 'Match' condition
//	          exists.
//
//	       TestStringName string
//	        - The label or name of the 'TestString' parameter.
//	          Used in error and informational messages.
//
//	       TestStringLength int
//	        - Actual number of text characters in the entire Test
//	          String ('TestString').
//
//	       TestStringLengthName string
//	        - The label or name of the 'TestStringLength'
//	          parameter. Used in error and informational
//	          messages.
//
//	       TestStringStartingIndex int
//	        - The starting index in the Test String where the
//	          search operation will begin.
//
//	       TestStringStartingIndexName string
//	        - The label or name of the TestStringStartingIndex
//	          parameter. Used in error and informational messages.
//
//	       TestStringDescription1 string
//	        - First of two optional description strings
//	          describing the Test String in the context of the
//	          current search operation.
//
//	       TestStringDescription2 string
//	        - Second of two optional description strings
//	          describing the Test String in the context of the
//	          current search operation.
//
//	       CollectionTestObjIndex int
//	        - If the Test String object resides in a collection
//	          of Test String objects, this parameter will record
//	          the array index of the current Test String object
//	          in the collection.
//
//	       NumValueType NumericValueType
//	        - Optional. This enumeration value specifies the type
//	          of numeric value for this Test Parameter as either
//	          an integer or floating point value.
//
//	          Possible values are listed as follows:
//	           NumValType.None()
//	           NumValType.FloatingPoint()
//	           NumValType.Integer()
//
//	       NumStrFormatType NumStrFormatTypeCode
//	        - Optional. This enumeration value specifies the
//	          Output Format Type for a number.
//
//	          Possible values are listed as follows:
//	           NumStrFmtType.None()
//	           NumStrFmtType.AbsoluteValue()
//	           NumStrFmtType.Binary()
//	           NumStrFmtType.CountryCulture()
//	           NumStrFmtType.Currency()
//	           NumStrFmtType.Binary()
//	           NumStrFmtType.Hexadecimal()
//	           NumStrFmtType.Octal()
//	           NumStrFmtType.ScientificNotation()
//
//	       NumSymLocation NumericSymbolLocation
//	        - Optional. This enumeration value specifies the
//	          relative location of a numeric symbol.
//
//	          Possible values are listed as follows:
//	           NumSymLocation.None()
//	           NumSymLocation.Before()
//	           NumSymLocation.Interior()
//	           NumSymLocation.After()
//
//	       NumSymbolClass NumericSymbolClass
//	        - Optional. This enumeration value specifies the
//	          Number Symbol Classification for a text character.
//
//	          Possible values are listed as follows:
//	           NumSymClass.None()
//	           NumSymClass.NumberSign()
//	           NumSymClass.CurrencySign()
//	           NumSymClass.IntegerSeparator()
//	           NumSymClass.DecimalSeparator()
//
//	       NumSignValue NumericSignValueType
//	        - Optional. This enumeration value specifies the
//	          number sign value.
//
//	          Possible values are listed as follows:
//	           NumSignVal.None()
//	           NumSignVal.Negative()
//	           NumSignVal.Zero()
//	           NumSignVal.Positive()
//
//	       PrimaryNumSignPosition NumSignSymbolPosition
//	       - Optional This enumeration value specifies the
//	         relative position of positive and negative number
//	         sign symbols in a number string. This is the
//	         Primary Type Code for Number Signs. Cases involving
//	         both 'Leading' and 'Trailing' symbols also make use
//	         of the 'SecondaryNumSignPosition'.
//
//	         Possible values are listed as follows:
//	          NumSignSymPos.None()
//	          NumSignSymPos.Before()
//	          NumSignSymPos.After()
//	          NumSignSymPos.BeforeAndAfter()
//
//	       SecondaryNumSignPosition NumSignSymbolPosition
//	        - Optional. This enumeration value specifies the
//	          relative position of positive and negative number
//	          sign symbols in a number string. This value is used
//	          in searches involving number signs which occur both
//	          before and after the numeric value.
//
//	          Possible values are listed as follows:
//	           NumSignSymPos.None()
//	           NumSignSymPos.Before()
//	           NumSignSymPos.After()
//	           NumSignSymPos.BeforeAndAfter()
//
//	       TextCharSearchType CharacterSearchType
//	       - Required. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation.
//
//	         Possible values are listed as follows:
//	          CharSearchType.None()                      - Invalid Value
//	          CharSearchType.LinearTargetStartingIndex() - Valid Default
//	          CharSearchType.SingleTargetChar()          - Valid
//	          CharSearchType.LinearEndOfString()         - Valid
//
//	     }
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	CharSearchResultsDto
//	   - If this method completes successfully, a fully populated
//	     instance of CharSearchResultsDto will be returned to the
//	     calling function.
//
//	     This structure contains information and data elements
//	     detailing the results of the text character search
//	     operation performed by this method.
//
//	     The CharSearchTestInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchResultsDto struct {
//
//	      SearchResultsName string
//	       - Optional. The Name, Label or descriptive Tag associated with
//	         the current instance of CharSearchResultsDto.
//
//	      SearchResultsFunctionChain string
//	       - Optional. A listing of the functions which led to the
//	         creation of this instance of CharSearchResultsDto.
//
//	      FoundSearchTarget bool
//	       - Signals a successful search outcome. If set to 'true' the
//	         Test String character or characters were found in the Target
//	         Search String.
//
//	      FoundSearchTargetOnPreviousSearch bool
//	       - Signals that the Search Target was located in a previous
//	         search operation.
//
//	      FoundFirstNumericDigitInNumStr bool
//	       - When set to 'true' this signals that the search operation
//	         has identified the first numeric digit in a string of text
//	         characters.
//
//	      TargetInputParametersName string
//	       - The Name, Label or descriptive Tag associated with an
//	         instance of CharSearchTargetInputParametersDto.
//
//	      TargetStringLength int
//	       - Actual number of text characters in the entire
//	         Target Search String ('TargetString').
//
//	      TargetStringSearchLength int
//	       - The actual number of characters within the Target
//	         Search String that are included in the search
//	         operation. This value may be less than the actual
//	         length of the Target Search String.
//
//	      TargetStringAdjustedSearchLength int
//	       - The adjusted or corrected Target String Search
//	         Length. This value is guaranteed to be equal to or
//	         less than the actual Target String Length.
//
//	      TargetStringStartingSearchIndex int
//	       - The index in 'TargetString' at which the search
//	         operation begins.
//
//	      TargetStringFirstFoundIndex int
//	       - The index of the first character position in the
//	         Target Search String occupied by the first
//	         character in the Test String.
//
//	      TargetStringLastFoundIndex int
//	       - The last character position in the Target Search
//	         String occupied by the last character in the Test
//	         String.
//
//	      TargetStringLastSearchIndex int
//	       - The index in Target Search String occupied by the
//	         last Target character searched. If the Search
//	         Target was found, this value is equal to the
//	         'TargetStringLastFoundIndex'. If the Search Target
//	         was NOT found this value is equal to the
//	         'TargetStringStartingSearchIndex'. This value is
//	         useful in computing the next index to be searched
//	         in the Target String.
//
//	      TargetStringNextSearchIndex int
//	       - The starting point for the next search operation.
//	         If the entire Target String was included in the
//	         last search, this value is set to -1.
//
//	          Example-1:
//	          String = "Hello"
//	          String Length = 5
//	          Last Search Index = 4
//	          TargetStringNextSearchIndex = -1
//
//	          Example-2:
//	          String = "Hello"
//	          String Length = 5
//	          Last Search Index = 2
//	          TargetStringNextSearchIndex = 3
//
//	      TargetStringDescription1 string
//	       - First of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      TargetStringDescription2 string
//	       - Second of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      TestInputParametersName string
//	       - The Name, Label or descriptive Tag associated with an
//	         instance of CharSearchTestInputParametersDto.
//
//	      TestStringLength int
//	       - Actual number of text characters in the entire Test
//	         String ('TestString').
//
//	      TestStringStartingIndex int
//	       - The starting index in the Test String where the
//	         search operation will begin.
//
//	      TestStringFirstFoundIndex int
//	       - The index number in Test String of the first test
//	         character to be located in the Target Search String.
//
//	      TestStringLastFoundIndex int
//	       - The index number in the Test String occupied by the
//	         last Test Character to be located in the Target
//	         String.
//
//	      TestStringDescription1 string
//	       - First of two optional description strings
//	         describing the Test String in the context of the
//	         current search operation.
//
//	      TestStringDescription2 string
//	       - Second of two optional description strings
//	         describing the Test String in the context of the
//	         current search operation.
//
//	      CollectionTestObjIndex int
//	       - If the Test String object resides in a collection
//	         of Test String objects, this parameter will record
//	         the array index of the current Test String object
//	         in the collection.
//
//	      ReplacementString *RuneArrayDto
//	       - A pointer to the Rune Array Data Transfer Object
//	         containing the Replacement Characters to be
//	         substituted for existing characters in a Target
//	         String.
//
//	      RemainderString *RuneArrayDto
//	       - A pointer to the Rune Array Data Transfer Object
//	         containing the remaining characters in a Target
//	         String which were NOT included in the search
//	         operation and which remain to be searched in future
//	         search operations. This string is also used in 'cut'
//	         operations where Target String is divided based on
//	         string delimiters.
//
//	      NumValueType NumericValueType
//	       - Optional. This enumeration value specifies the type
//	         of numeric value for this Test Parameter as either
//	         an integer or floating point value.
//
//	         Possible values are listed as follows:
//	          NumValType.None()
//	          NumValType.FloatingPoint()
//	          NumValType.Integer()
//
//	      NumStrFormatType NumStrFormatTypeCode
//	       - Optional. This enumeration value specifies the
//	         Output Format Type for a number.
//
//	         Possible values are listed as follows:
//	          NumStrFmtType.None()
//	          NumStrFmtType.AbsoluteValue()
//	          NumStrFmtType.Binary()
//	          NumStrFmtType.CountryCulture()
//	          NumStrFmtType.Currency()
//	          NumStrFmtType.Binary()
//	          NumStrFmtType.Hexadecimal()
//	          NumStrFmtType.Octal()
//	          NumStrFmtType.ScientificNotation()
//
//	      NumSymLocation NumericSymbolLocation
//	       - Optional. This enumeration value specifies the
//	         relative location of a numeric symbol.
//
//	         Possible values are listed as follows:
//	          NumSymLocation.None()
//	          NumSymLocation.Before()
//	          NumSymLocation.Interior()
//	          NumSymLocation.After()
//
//	      NumSymbolClass NumericSymbolClass
//	       - Optional. This enumeration value specifies the
//	         Number Symbol Classification for a text character.
//
//	         Possible values are listed as follows:
//	          NumSymClass.None()
//	          NumSymClass.NumberSign()
//	          NumSymClass.CurrencySign()
//	          NumSymClass.IntegerSeparator()
//	          NumSymClass.DecimalSeparator()
//
//	      NumSignValue NumericSignValueType
//	       - Optional. This enumeration value specifies the
//	         number sign value.
//
//	         Possible values are listed as follows:
//	          NumSignVal.None()
//	          NumSignVal.Negative()
//	          NumSignVal.Zero()
//	          NumSignVal.Positive()
//
//	      PrimaryNumSignPosition NumSignSymbolPosition
//	       - Optional This enumeration value specifies the
//	         relative position of positive and negative number
//	         sign symbols in a number string. This is the
//	         Primary Type Code for Number Signs. Cases involving
//	         both 'Leading' and 'Trailing' symbols also make use
//	         of the 'SecondaryNumSignPosition'.
//
//	         Possible values are listed as follows:
//	          NumSignSymPos.None()
//	          NumSignSymPos.Before()
//	          NumSignSymPos.After()
//	          NumSignSymPos.BeforeAndAfter()
//
//	      SecondaryNumSignPosition NumSignSymbolPosition
//	       - Optional. This enumeration value specifies the
//	         relative position of positive and negative number
//	         sign symbols in a number string. This value is used
//	         in searches involving number signs which occur both
//	         before and after the numeric value.
//
//	         Possible values are listed as follows:
//	          NumSignSymPos.None()
//	          NumSignSymPos.Before()
//	          NumSignSymPos.After()
//	          NumSignSymPos.BeforeAndAfter()
//
//	      TextCharSearchType CharacterSearchType
//	       - Required. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation.
//
//	         Possible values are listed as follows:
//	          CharSearchType.None() - Invalid value
//	          CharSearchType.LinearTargetStartingIndex() - Default
//	          CharSearchType.SingleTargetChar()
//	          CharSearchType.LinearEndOfString()
//	     }
//
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
func (runeDtoElectron *runeArrayDtoElectron) linearTargetStartingIndexSearch(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchRuneArrayResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	searchResults := CharSearchRuneArrayResultsDto{}.New()

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

	searchResults.SearchResultsFunctionChain =
		ePrefix.String()

	j := 0

	for i := targetInputParms.TargetStringCurrentSearchIndex; i < targetInputParms.TargetStringAdjustedSearchLength; i++ {

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

			// Target String Parameters
			searchResults.TargetStringLastFoundIndex = i

			searchResults.TargetStringLastSearchIndex = i

			searchResults.TargetStringFirstFoundIndex =
				searchResults.TargetStringLastFoundIndex -
					searchResults.TestStringLength +
					1

			searchResults.TargetStringCurrentSearchIndex = i

			searchResults.TargetStringNextSearchIndex =
				searchResults.TargetStringCurrentSearchIndex + 1

			if searchResults.TargetStringNextSearchIndex >=
				targetInputParms.TargetStringLength {

				searchResults.TargetStringNextSearchIndex =
					-1
			}

			if testInputParms.RequestFoundTestCharacters {
				err = searchResults.FoundRuneArrayChars.CopyIn(
					testInputParms.TestString,
					ePrefix.XCpy(
						"searchResults.FoundRuneArrayChars<-"+
							"testInputParms.TestString"))

				if err != nil {
					return searchResults, err
				}
			}

			if testInputParms.RequestRemainderString {
				if searchResults.TargetStringNextSearchIndex >
					-1 {

					searchResults.RemainderString.Empty()

					searchResults.RemainderString.CharsArray =
						append(
							searchResults.RemainderString.CharsArray,
							targetInputParms.TargetString.CharsArray[searchResults.TargetStringCurrentSearchIndex:]...)
				}
			}

			return searchResults, err
		}

	}

	// Search Failed!
	// No MATCHING Text Characters were
	// Found.
	return searchResults, err
}

// singleCharacterSearch - Performs a Single Target Character
// Search on the target string (rune array) specified by input
// parameter 'targetInputParms'.
//
// ----------------------------------------------------------------
//
// # TERMINOLOGY
//
// Text Character Search algorithms typically perform comparisons
// between two strings or groups of text characters to determine
// the search outcome. A successful search outcome usual involves
// finding one or more text characters from one string inside a
// second string. A successful search outcome is often referred to
// as a 'Match' condition because characters in one string were
// compared and matched with characters in another string.
//
// Character Search algorithms using the Character Search Type
// ('CharacterSearchType') rely on a framework consisting of a
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
//	Example #1
//	                           1         2         3
//	          Index  0123456789012345678901234567890
//	 Target String: "Hey, Xray-4 is the call sign."
//	 Target String Starting Index: 5
//	   Test String: "ZFXyURJK"
//
// In this example of a Single Target Character Search, the search
// will begin and end at Target Search String index number 5. Since
// one of the Test String Characters ('X') matches the 'X'
// character at index number 5 in the Target Search String, the
// search operation is classified as a success. A matching
// character was found.
//
//	Example #2
//	                           1         2         3
//	          Index  0123456789012345678901234567890
//	 Target String: "Hey, Xray-4 is the call sign."
//	 Target String Starting Index: 0
//	   Test String: "ZFXyURJK"
//
// In this second example of a Single Target Character Search, the
// search will begin and end at Target Search String index number
// zero (0). Since NONE of the Test String Characters matches the
// 'H' character at index number zero (0) in the Target Search
// String, the search operation is classified as a failure. No
// matching character was found.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	targetInputParms           CharSearchTargetInputParametersDto
//	   - This Dto, or Data Transfer Object, contains the input
//	     parameters and specifications associated with the Target
//	     Search String which will be used in completing the Linear
//	     Target Starting Index Search operation.
//
//	     The CharSearchTargetInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchTargetInputParametersDto struct {
//
//	      TargetInputParametersName           string
//	       - The Name, Label or descriptive Tag associated with this
//	         instance of CharSearchTargetInputParametersDto. If empty,
//	         this string will be defaulted to "TargetInputParameters"
//
//	      TargetString                        *RuneArrayDto
//	       - A pointer to the RuneArrayDto containing the Target
//	         Search String text characters used in the search
//	         algorithm. Target Characters are compared against
//	         Test Characters to determine if a 'Match' condition
//	         exists.
//
//	      TargetStringName                    string
//	       - The label or name of the 'TargetString' parameter.
//	         Used in error and informational messages.
//
//	      TargetStringLength                  int
//	       - Actual number of text characters in the entire
//	         Target Search String ('TargetString').
//
//	      TargetStringLengthName              string
//	       - The label or name of the 'TargetStringLength' parameter.
//	         Used in error and informational messages.
//
//	      TargetStringStartingSearchIndex     int
//	       - The index in 'TargetString' at which the search
//	         operation begins.
//
//	      TargetStringStartingSearchIndexName string
//	       - The label or name of the
//	         TargetStringStartingSearchIndex parameter.
//	         Used in error and informational messages.
//
//	      TargetStringSearchLength            int
//	       - The actual number of characters within the Target
//	         Search String that are included in the search
//	         operation. This value may be less than the actual
//	         length of the Target Search String.
//
//	      TargetStringSearchLengthName        string
//	       - The label or name of the TargetStringSearchLength
//	         parameter. Used in error and informational
//	         messages.
//
//	      TargetStringAdjustedSearchLength    int
//	       - The adjusted or corrected Target String Search
//	         Length. This value is guaranteed to be equal to or
//	         less than the actual Target String Length.
//
//	      TargetStringDescription1            string
//	       - First of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      TargetStringDescription2            string
//	       - Second of two optional description strings
//	         describing the Target Search String in the context
//	         of the current search operation.
//
//	      FoundFirstNumericDigitInNumStr      bool
//	       - When set to 'true' this signals that the first
//	         numeric digit has been identified in the text
//	         characters specified by 'TargetString'
//
//	      CollectionTestObjIndex int
//	       - If the Test String object resides in a collection
//	         of Test String objects, this parameter will record
//	         the array index of the current Test String object
//	         in the collection.
//
//	      TextCharSearchType                  CharacterSearchType
//	       - Optional. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation. When set to a valid value,
//	         this specification will override the search
//	         specification contained in the Test Input
//	         Parameters Data Transfer Object.
//
//	        Valid CharSearch Type values are listed as follows:
//	        TextCharSearchType.None() - Invalid Value
//	        TextCharSearchType.LinearTargetStartingIndex() - Default
//	        TextCharSearchType.SingleTargetChar()
//	        TextCharSearchType.LinearEndOfString()
//	     }
//
//
//	testInputParms      CharSearchTestInputParametersDto
//	   - This Dto, or Data Transfer Object, contains the input
//	     parameters and specifications associated with the Test
//	     String which will be used in completing the Linear Target
//	     Starting Index Search operation. This algorithm compares
//	     text characters in the Target Search String
//	     ('targetInputParms') with those contained in the Test
//	     String specified below.
//
//	     The CharSearchTestInputParametersDto structure consists
//	     of the following member variable data elements.
//
//	     type CharSearchTestInputParametersDto struct {
//
//	       TestInputParametersName string
//	        - The Name, Label or descriptive Tag associated with this
//	          instance of CharSearchTestInputParametersDto. If empty,
//	          this string will be defaulted to "TestInputParameters"
//
//	       TestString *RuneArrayDto
//	        - A pointer to the Rune Array Data Transfer
//	          Object containing the Test Characters to be
//	          used in a text character search algorithm.
//	          Target Characters are compared against Test
//	          Characters to determine if a 'Match' condition
//	          exists.
//
//	       TestStringName string
//	        - The label or name of the 'TestString' parameter.
//	          Used in error and informational messages.
//
//	       TestStringLength int
//	        - Actual number of text characters in the entire Test
//	          String ('TestString').
//
//	       TestStringLengthName string
//	        - The label or name of the 'TestStringLength'
//	          parameter. Used in error and informational
//	          messages.
//
//	       TestStringStartingIndex int
//	        - The starting index in the Test String where the
//	          search operation will begin.
//
//	       TestStringStartingIndexName string
//	        - The label or name of the TestStringStartingIndex
//	          parameter. Used in error and informational messages.
//
//	       TestStringDescription1 string
//	        - First of two optional description strings
//	          describing the Test String in the context of the
//	          current search operation.
//
//	       TestStringDescription2 string
//	        - Second of two optional description strings
//	          describing the Test String in the context of the
//	          current search operation.
//
//	       CollectionTestObjIndex int
//	        - If the Test String object resides in a collection
//	          of Test String objects, this parameter will record
//	          the array index of the current Test String object
//	          in the collection.
//
//	       NumValueType NumericValueType
//	        - Optional. This enumeration value specifies the type
//	          of numeric value for this Test Parameter as either
//	          an integer or floating point value.
//
//	          Possible values are listed as follows:
//	           NumValType.None()
//	           NumValType.FloatingPoint()
//	           NumValType.Integer()
//
//	       NumStrFormatType NumStrFormatTypeCode
//	        - Optional. This enumeration value specifies the
//	          Output Format Type for a number.
//
//	          Possible values are listed as follows:
//	           NumStrFmtType.None()
//	           NumStrFmtType.AbsoluteValue()
//	           NumStrFmtType.Binary()
//	           NumStrFmtType.CountryCulture()
//	           NumStrFmtType.Currency()
//	           NumStrFmtType.Binary()
//	           NumStrFmtType.Hexadecimal()
//	           NumStrFmtType.Octal()
//	           NumStrFmtType.ScientificNotation()
//
//	       NumSymbolLocation NumericSymbolLocation
//	        - Optional. This enumeration value specifies the
//	          relative location of a numeric symbol.
//
//	          Possible values are listed as follows:
//	           NumSymLocation.None()
//	           NumSymLocation.Before()
//	           NumSymLocation.Interior()
//	           NumSymLocation.After()
//
//	       NumSymbolClass NumericSymbolClass
//	        - Optional. This enumeration value specifies the
//	          Number Symbol Classification for a text character.
//
//	          Possible values are listed as follows:
//	           NumSymClass.None()
//	           NumSymClass.NumberSign()
//	           NumSymClass.CurrencySign()
//	           NumSymClass.IntegerSeparator()
//	           NumSymClass.DecimalSeparator()
//
//	       NumSignValue NumericSignValueType
//	        - Optional. This enumeration value specifies the
//	          number sign value.
//
//	          Possible values are listed as follows:
//	           NumSignVal.None()
//	           NumSignVal.Negative()
//	           NumSignVal.Zero()
//	           NumSignVal.Positive()
//
//	       PrimaryNumSignPosition NumSignSymbolPosition
//	       - Optional This enumeration value specifies the
//	         relative position of positive and negative number
//	         sign symbols in a number string. This is the
//	         Primary Type Code for Number Signs. Cases involving
//	         both 'Leading' and 'Trailing' symbols also make use
//	         of the 'SecondaryNumSignPosition'.
//
//	         Possible values are listed as follows:
//	          NumSignSymPos.None()
//	          NumSignSymPos.Before()
//	          NumSignSymPos.After()
//	          NumSignSymPos.BeforeAndAfter()
//
//	       SecondaryNumSignPosition NumSignSymbolPosition
//	        - Optional. This enumeration value specifies the
//	          relative position of positive and negative number
//	          sign symbols in a number string. This value is used
//	          in searches involving number signs which occur both
//	          before and after the numeric value.
//
//	          Possible values are listed as follows:
//	           NumSignSymPos.None()
//	           NumSignSymPos.Before()
//	           NumSignSymPos.After()
//	           NumSignSymPos.BeforeAndAfter()
//
//	       TextCharSearchType CharacterSearchType
//	       - Required. An enumeration value signaling the type
//	         of text character search algorithm used to conduct
//	         this search operation.
//
//	         Possible values are listed as follows:
//	          CharSearchType.None()                      - Invalid Value
//	          CharSearchType.LinearTargetStartingIndex() - Valid Default
//	          CharSearchType.SingleTargetChar()          - Valid
//	          CharSearchType.LinearEndOfString()         - Valid
//
//	     }
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	CharSearchRuneArrayResultsDto
//	   - If this method completes successfully, a fully populated
//	     instance of CharSearchRuneArrayResultsDto will be returned to the
//	     calling function.
//
//	     This structure contains information and data elements
//	     detailing the results of the text character search
//	     operation performed by this method.
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
func (runeDtoElectron *runeArrayDtoElectron) singleCharacterSearch(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchRuneArrayResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	searchResults := CharSearchRuneArrayResultsDto{}.New()

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

	searchResults.SearchResultsFunctionChain =
		ePrefix.String()

	targetChar :=
		targetInputParms.TargetString.CharsArray[targetInputParms.TargetStringCurrentSearchIndex]

	for j := testInputParms.TestStringStartingIndex; j < testInputParms.TestStringLength; j++ {

		if testInputParms.TestString.CharsArray[j] == targetChar {
			// Search SUCCESSFUL! SINGLE CHARACTER MATCH!
			// Exit Here!

			searchResults.FoundSearchTarget = true

			// Set Test String Parameters
			searchResults.TestStringStartingIndex =
				testInputParms.TestStringStartingIndex

			searchResults.TestStringFirstFoundIndex = j

			searchResults.TestStringLastFoundIndex = j

			// Set Target String Parameters
			searchResults.TargetStringLastFoundIndex =
				targetInputParms.TargetStringCurrentSearchIndex

			searchResults.TargetStringFirstFoundIndex =
				targetInputParms.TargetStringCurrentSearchIndex

			searchResults.TargetStringLastSearchIndex =
				targetInputParms.TargetStringCurrentSearchIndex

			searchResults.TargetStringCurrentSearchIndex =
				targetInputParms.TargetStringCurrentSearchIndex

			searchResults.TargetStringNextSearchIndex =
				searchResults.TargetStringCurrentSearchIndex + 1

			if searchResults.TargetStringNextSearchIndex >=
				targetInputParms.TargetStringLength {

				searchResults.TargetStringNextSearchIndex =
					-1
			}

			if testInputParms.RequestFoundTestCharacters {

				err = searchResults.FoundRuneArrayChars.CopyIn(
					testInputParms.TestString,
					ePrefix.XCpy(
						"searchResults.FoundRuneArrayChars<-"+
							"testInputParms.TestString"))

				if err != nil {
					return searchResults, err
				}
			}

			if testInputParms.RequestRemainderString {

				if searchResults.TargetStringNextSearchIndex >
					-1 {

					searchResults.RemainderString.Empty()

					searchResults.RemainderString.CharsArray =
						append(
							searchResults.RemainderString.CharsArray,
							testInputParms.TestString.CharsArray[j])
				}
			}

			return searchResults, err
		}
	}

	return searchResults, err
}

// setRuneArray - Receives a pointer to an instance of
// RuneArrayDto ('runeArrayDto') and a rune array ('newChars').
// This method then proceeds to copy the 'newChars' rune array
// to the internal member variable rune array ('CharsArray')
// maintained by the 'runeArrayDto' instance.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// For the input parameter 'runeArrayDto', all pre-existing data in
// the internal rune array member variable, 'charArray' will be
// deleted and overwritten with new data.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto               *RuneArrayDto
//	   - A pointer to an instance of RuneArrayDto. The internal
//	     rune array maintained by this instance, 'CharsArray',
//	     will be deleted and overwritten.
//
//	     New data taken from input parameter, 'charArray', will
//	     then be copied to 'runeArrayDto.CharsArray'
//
//
//	charArray                  []rune
//	   - An array of runes used to populate the internal member
//	     variable rune array for the 'runeArrayDto' instance of
//	     RuneArrayDto.
//
//	     If this array is empty or has a zero length, an error will
//	     be returned.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
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
func (runeDtoElectron *runeArrayDtoElectron) setRuneArray(
	runeArrayDto *RuneArrayDto,
	newChars []rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	searchResults := CharSearchResultsDto{}

	searchResults.Empty()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoElectron."+
			"setRuneArray()",
		"")

	if err != nil {

		return err

	}

	if runeArrayDto == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeArrayDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenOfNewCharsArray := len(newChars)

	if lenOfNewCharsArray == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'newChars' is invalid!\n"+
			"'newChars' is empty and has a length of zero.",
			ePrefix.String())

		return err
	}

	runeArrayDto.CharsArray =
		make([]rune, lenOfNewCharsArray)

	for i := 0; i < lenOfNewCharsArray; i++ {
		runeArrayDto.CharsArray[i] =
			newChars[i]
	}

	return err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrHelperQuark
//
// Provides number string utility methods
type numStrHelperQuark struct {
	lock *sync.Mutex
}

//	extractNumericDigits
//
//	Examines an input parameter 'targetStr' to identify
//	and extract the first instance of a number string.
//	The number string will be comprised if one or more
//	consecutive numeric digits (0-9) and may include
//	leading, trailing or interior non-numeric characters
//	as specified by input parameters.
//
//	The search for this number string will be started at
//	the index specified by input parameter 'startIdx'.
//	Beginning at 'startIdx' the 'targetStr' will be
//	searched to identify and extract the first instance
//	of a number string.
//
//	A number string is usually defined a string of
//	consecutive numeric digits. However, this method
//	allows the caller to include additional non-numeric
//	characters as identified by input parameters:
//		'keepLeadingChars'
//		'keepInteriorChars'
//		'keepTrailingChars'
//
//	'keepLeadingChars' is a string of characters which
//	will be prefixed to the number string if those
//	characters exist in 'targetStr' and immediately
//	precede the number string.
//
//	'keepInteriorChars' is a string of characters which,
//	if they exist within the number string, will be
//	retained and presented in the final extracted number
//	string.
//
//	'keepTrailingChars' is a string of characters which
//	will be suffixed to the end of the final extracted
//	number string.  To qualify, the designated
//	'keepTrailingChars' must immediately follow the
//	number string contained in 'targetStr'.
//
//	If successfully located within 'targetStr', the first
//	instance of a number string along with
//	characteristics describing that number string are
//	returned in a Type 'NumStrProfileDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	targetStr           string    - The target string to be searched for the first instance of
//	                                a number string. A number string is usually defined as a
//	                                string comprised of one or more consecutive numeric digits.
//	                                Additional parameters provided by this method will allow
//	                                the caller to insert specified non-numeric characters at
//	                                the beginning, end or interior of a number string.
//
//	startIdx               int    - The starting index in input parameter 'targetStr'
//	                                from which the search for a number string will be
//	                                initiated. This useful in extracting multiple number
//	                                strings form a single 'targetStr'.
//
//	keepLeadingChars    string    - This string contains non-numeric characters which will be
//	                                retained as a prefix to the final number string extracted
//	                                from the 'targetStr' parameter. To be included, these characters
//	                                must exist in 'targetStr' and must immediately precede the
//	                                first instance of a number string.
//
//	                                For example, if the target string is "Hello $123789 world" and
//	                                parameter 'keepLeadingChars' includes the USA currency character,
//	                                '$', the returned number string would be '$123789'.  If no currency
//	                                character was included in 'keepLeadingChars', the returned number
//	                                string would be '123789'. It is worth noting that if the target
//	                                string was '$ 123789' and a currency symbol, '$', was included
//	                                in 'keepLeadingChars', the returned number string would still be
//	                                '123789' because 'keepLeadingChars' characters must immediately
//	                                precede the string of numeric digits in 'targetStr'.
//
//	                                Specifically, if the plus ('+') and minus ('-') sign are NOT
//	                                included in 'keepLeadingChars' those leading number signs will
//	                                never be included in the final number string.
//
//	                                Leading characters will not be repeated. If for some reason you
//	                                wanted to retain two leading currency symbols ("$$") it would be
//	                                necessary to include two currency characters in 'keepLeadingChars'.
//
//	keepInteriorChars   string    - This string contains non-numeric characters which will be retained
//	                                as valid characters within the final extracted number string. The
//	                                characters must exist withing the first instance of a number string
//	                                located in 'targetStr'. Such interior characters might include
//	                                thousands separators (commas) or decimal points (periods).
//
//	                                For example, if a comma and a period are included in 'keepInteriorChars'
//	                                and the target string is "Hello word 123,456,789.25 !", the returned
//	                                number string would be "123,456,789.25".  If the comma character was
//	                                NOT included in the 'keepInteriorChars' string, the returned number
//	                                string would be '123', since the number string extraction parser
//	                                would break on the comma, a non-numeric digit.
//
//	                                'keepInteriorChars' will NOT allow multiple non-numeric characters
//	                                to exist within the interior of the final extracted number string.
//	                                Only single non-numeric characters are allowed within a number string.
//
//	keepTrailingChars   string    - This string contains non-numeric characters which should be retained
//	                                at the end of the final number string. By default, a non-numeric
//	                                character will mark the end of a number string. However, if the caller
//	                                elects to use parameter 'keepTrailingChars' to retain non-numeric
//	                                characters such as a trailing right-parenthesis, then those non-numeric
//	                                characters will be retained in the final extracted number string.
//
//	                                Trailing characters will not be repeated. If for some reason you
//	                                wanted to retain two closing parentheses symbols ("))") it would be
//	                                necessary to include closing parentheses characters in 'keepTrailingChars'.
//
//	                                It should be emphasized that 'keepTrailingChars' must immediately
//	                                follow the first instance of a number string in parameter, 'targetStr'.
//
//	                                Example #1:
//	                                  Target String = "Hello world, (1234). Today is new day."
//	                                  keepLeadingChars = "("
//	                                  keepInteriorChars = ""
//	                                  keepTrailingChars= ")"
//	                                  Extracted Number String = "(1234)"
//
//	                                Example #2:
//	                                  Target String = "Hello world, USA GDP growth is projected at 1.8%."
//	                                  keepLeadingChars = ""
//	                                  keepInteriorChars = "."
//	                                  keepTrailingChars= "%"
//	                                  Extracted Number String = "1.8%"
//
//
//	ePrefix             *ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the names of the calling method or methods listed
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
//	NumStrProfileDto    - If successful, this method will return a type 'NumStrProfileDto'
//	                      populated with the extracted number string and additional profile
//	                      information related to the extracted number string.
//
//	   type NumStrProfileDto struct {
//
//	         TargetStr            string   //  The original target string which is scanned for a
//	                                       //    number string
//
//	         TargetStrStartIndex  int      //  The starting index in 'TargetStr' from which the
//	                                       //    number string search was initiated.
//
//	         LeadingSignIndex     int      //  The string index of a leading sign in 'NumStr' below. If a
//	                                       //    leading sign character is NOT present in 'NumStr' this
//	                                       //    value is set to -1
//
//	         LeadingSignChar      string   //  If a leading sign character (plus '+' or minus '-')
//	                                       //    exists in data field 'NumStr' (below), it is stored
//	                                       //    in this string.
//
//	         FirstNumCharIndex    int      //  The index in 'TargetStr' (above) where the first character
//	                                       //    of the extracted number string is located.
//
//	         NextTargetStrIndex   int      //  The index of the next character in 'TargetStr' immediately
//	                                       //    following the extracted number string.
//
//	         NumStrLen            int      //  The length of the extracted number string.
//
//	         NumStr               string   //  The number string extracted from 'TargetStr'.
//	   }
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered this return value will contain
//	     an appropriate error message.
//
//	     If 'startIndex' is less than zero or if 'startIndex'
//	     exceeds the last character index in 'targetStr', an error
//	     will be returned.
//
//	     If an error message is returned, the text value for input
//	     parameter 'ePrefix' (error prefix) will be inserted or
//	     prefixed at the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//	ePrefix := "TestStrOps_ExtractNumericDigits_06() "
//	targetStr := "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"
//	startIndex := 0
//	keepLeadingChars := "$("
//	keepInteriorChars := ",."
//	keepTrailingChars := ")"
//
//
//	nStrDto,
//	err :=
//	  new(StrMech).
//	       ExtractNumericDigits(
//	           targetStr,
//	           startIndex,
//	           keepLeadingChars,
//	           keepInteriorChars,
//	           keepTrailingChars,
//	           ePrefix)
//
//	----------------------------------------
//
//	nStrDto.NumStr is now equal to:
//	   "$(1,250,364.33)"
func (nStrHelperQuark *numStrHelperQuark) extractNumericDigits(
	targetStr string,
	startIndex int,
	keepLeadingChars string,
	keepInteriorChars string,
	keepTrailingChars string,
	ePrefix *ePref.ErrPrefixDto) (
	NumStrProfileDto,
	error) {

	if nStrHelperQuark.lock == nil {
		nStrHelperQuark.lock = new(sync.Mutex)
	}

	nStrHelperQuark.lock.Lock()

	defer nStrHelperQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"numStrHelperQuark." +
			"extractNumericDigits()")

	nStrDto := NumStrProfileDto{}.New()
	nStrDto.TargetStr = targetStr
	nStrDto.StartIndex = startIndex

	lenTargetStr := len(targetStr)

	if lenTargetStr == 0 {
		return nStrDto,
			fmt.Errorf("%v\n"+
				"ERROR: Input parameter 'targetStr' is "+
				"an empty string!\n",
				ePrefix.String())
	}

	if startIndex < 0 {
		return nStrDto, fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'startIndex' is less than zero!\n"+
			"startIndex='%v'",
			ePrefix.String(),
			startIndex)
	}

	if startIndex >= lenTargetStr {
		return nStrDto, fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'startIndex' is INVALID!\n"+
			"'startIndex' exceeds the last character index in "+
			"'targetStr'\n"+
			"startIndex='%v'\tlast character index='%v'\n"+
			"targetStr='%v'\n",
			ePrefix.String(),
			startIndex, lenTargetStr-1, targetStr)
	}

	targetRunes := []rune(targetStr)
	lenTargetStr = len(targetRunes)

	keepLeadingRunes := make([]rune, 0)
	lenKeepLeadingRunes := 0

	keepInteriorRunes := make([]rune, 0)
	lenKeepInteriorRunes := 0

	keepTrailingRunes := make([]rune, 0)
	lenKeepTrailingRunes := 0

	if len(keepLeadingChars) > 0 {

		// Remove any numeric characters
		for a := 0; a < len(keepLeadingChars); a++ {

			if keepLeadingChars[a] >= '0' &&
				keepLeadingChars[a] <= '9' {
				continue
			}

			keepLeadingRunes = append(keepLeadingRunes, rune(keepLeadingChars[a]))

		}

		lenKeepLeadingRunes = len(keepLeadingRunes)
	}

	if len(keepInteriorChars) > 0 {

		// Remove any numeric characters
		for a := 0; a < len(keepInteriorChars); a++ {

			if keepInteriorChars[a] >= '0' &&
				keepInteriorChars[a] <= '9' {
				continue
			}

			keepInteriorRunes = append(keepInteriorRunes, rune(keepInteriorChars[a]))

		}

		lenKeepInteriorRunes = len(keepInteriorRunes)
	}

	if len(keepTrailingChars) > 0 {

		// Remove any numeric characters
		for a := 0; a < len(keepTrailingChars); a++ {

			if keepTrailingChars[a] >= '0' &&
				keepTrailingChars[a] <= '9' {
				continue
			}

			keepTrailingRunes = append(keepTrailingRunes, rune(keepTrailingChars[a]))

		}

		lenKeepTrailingRunes = len(keepTrailingRunes)
	}

	numberRunesCaptured := make([]rune, 0, 20)
	lenNumberRunesCaptured := 0

	leadingCharRunesCaptured := make([]rune, 0, 20)
	lenLeadingCharRunesCaptured := 0

	firstNumericDigitIdx := -1

	for e := startIndex; e < lenTargetStr; e++ {

		if targetRunes[e] >= '0' &&
			targetRunes[e] <= '9' &&
			firstNumericDigitIdx == -1 {
			// Target has at least one numeric
			// digit - and we found it.
			firstNumericDigitIdx = e
			break
		}
	}

	if firstNumericDigitIdx == -1 {
		// There are no numeric digits
		// in this target string.
		// EXIT HERE!!!
		return nStrDto, nil
	}

	firstNumStrCharIdx := -1
	leadingSignChar := ""

	// Check for leading non-numeric characters that
	// need to be retained at the front of the number
	// string.
	if lenKeepLeadingRunes > 0 &&
		startIndex < firstNumericDigitIdx {

		for f := firstNumericDigitIdx - 1; f >= startIndex; f-- {

			for g := 0; g < lenKeepLeadingRunes; g++ {

				if keepLeadingRunes[g] == targetRunes[f] {

					if keepLeadingRunes[g] == '+' ||
						keepLeadingRunes[g] == '-' {

						// This is a leading sign char
						leadingSignChar = string(targetRunes[f])

						leadingCharRunesCaptured = append(leadingCharRunesCaptured, targetRunes[f])
						// Delete Leading Sign character. It will not be repeated in
						// future searches. Only one leading sign char per number string.

						keepLeadingRunes = append(keepLeadingRunes[0:g], keepLeadingRunes[g+1:]...)
						lenKeepLeadingRunes--

						firstNumStrCharIdx = f

						// Now delete the alternative leading sign character.
						// There are only two - plus or minus
						nextSignChar := '-'

						if leadingSignChar == "-" {
							nextSignChar = '+'
						}

						// Leading sign char has been found. Now delete the
						// alternative lead sign char to avoid duplications
						for m := 0; m < lenKeepLeadingRunes; m++ {
							if keepLeadingRunes[m] == nextSignChar {
								keepLeadingRunes = append(keepLeadingRunes[0:m], keepLeadingRunes[m+1:]...)
								lenKeepLeadingRunes--
							}
						}

						break

					} else {

						// Standard Keep Leading Rune character found
						leadingCharRunesCaptured = append(leadingCharRunesCaptured, targetRunes[f])
						// Delete Leading Rune character. It will not be repeated in
						// future searches

						firstNumStrCharIdx = f

						keepLeadingRunes = append(keepLeadingRunes[0:g], keepLeadingRunes[g+1:]...)
						lenKeepLeadingRunes--
						break
					}
				}
			}

			t := len(leadingCharRunesCaptured)

			if t > lenLeadingCharRunesCaptured {
				lenLeadingCharRunesCaptured = t
				continue
			}

			break
		}
	}

	leadingSignIndex := -1

	if lenLeadingCharRunesCaptured > 0 {

		for h := lenLeadingCharRunesCaptured - 1; h >= 0; h-- {

			if leadingCharRunesCaptured[h] == '+' ||
				leadingCharRunesCaptured[h] == '-' {

				numberRunesCaptured = append(numberRunesCaptured, leadingCharRunesCaptured[h])
				leadingSignIndex = lenNumberRunesCaptured
				lenNumberRunesCaptured++

			} else {
				numberRunesCaptured = append(numberRunesCaptured, leadingCharRunesCaptured[h])
				lenNumberRunesCaptured++
			}
		}
	}

	// Main Number String Extraction Loop
	isEndOfNumStr := false

	for i := firstNumericDigitIdx; i < lenTargetStr; i++ {

		if !isEndOfNumStr {

			if targetRunes[i] >= '0' && targetRunes[i] <= '9' {

				numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])
				continue
			}

			for j := 0; j < lenKeepInteriorRunes; j++ {

				if targetRunes[i] == keepInteriorRunes[j] {

					if i+1 >= lenTargetStr ||
						(targetRunes[i+1] < '0' || targetRunes[i+1] > '9') {
						// We are either at the end of string or the next char
						// is NOT a numeric character.
						goto trailChar
					}

					numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])

					goto numDigitLoop
				}
			}

		}

	trailChar:
		isEndOfNumStr = true

		for k := 0; k < lenKeepTrailingRunes; k++ {

			if targetRunes[i] == keepTrailingRunes[k] {
				numberRunesCaptured = append(numberRunesCaptured, targetRunes[i])
				// Only one instance of a keep trailing rune character is captured.
				// Delete the keep trailing rune character to prevent repeat captures.
				keepTrailingRunes = append(keepLeadingRunes[0:k], keepTrailingRunes[k+1:]...)
				lenKeepTrailingRunes--
				goto numDigitLoop
			}

		}

		// Non-numeric character and Non-Trailing Character: Exit the Loop
		break

	numDigitLoop:
	}

	if len(numberRunesCaptured) > 0 {
		nStrDto.NumStr = string(numberRunesCaptured)

		if firstNumStrCharIdx > -1 {
			nStrDto.FirstNumCharIndex = firstNumStrCharIdx
		} else {
			nStrDto.FirstNumCharIndex = firstNumericDigitIdx
		}

		nStrDto.NumStrLen = len(nStrDto.NumStr)
		nStrDto.LeadingSignChar = leadingSignChar
		nStrDto.LeadingSignIndex = leadingSignIndex
		nStrDto.NextTargetStrIndex =
			nStrDto.FirstNumCharIndex + nStrDto.NumStrLen

		if nStrDto.NextTargetStrIndex >= len(targetStr) {
			nStrDto.NextTargetStrIndex = -1
		}
	}

	return nStrDto, nil
}

// extractNumRunes
//
// Receives an array of runes and extracts the numeric
// digits as text characters. Set
// 'characterSearchLength' to minus one (-1) for
// end-of-string search.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	targetSearchString				RuneArrayDto
//
//		A instance of RuneArrayDto. Type RuneArrayDto
//		contains the string of text characters which will
//		be searched for the presence of a numeric digit
//		characters.
//
//			type RuneArrayDto struct {
//				CharsArray []rune
//			}
//
//	targetSearchStringName			string
//
//		The name of the target search string which will
//		be displayed in error or informational messages
//		returned by this method.
//
//	startingSearchIndex				int
//
//		The index in the target string rune array
//		('targetSearchString') where the search for
//		numeric digit characters will begin.
//
//	breakOnCharSearchLength			int
//
//		The actual number of characters within the Raw
//		Number String Rune Array ('targetSearchString')
//		that are included in the search for numeric
//		character digits.
//
//		If this value is set to -1, the search length
//		will be configured to include the last index in
//		'targetSearchString'. In other words the search
//		will proceed	to the end of
//		'targetSearchString' array.
//
//	negativeNumSearchSpecsCol		NegNumSearchSpecCollection
//
//		Type NegNumSearchSpecCollection is a collection
//		of NegativeNumberSearchSpec object. This
//		collection represents all the negative number
//		signs which might be located with the Target
//		Search String ('targetSearchString').
//
//		This collection allows the flexibility for
//		locating such negative number signs as leading
//		minus signs ('-'), trailing minus signs ('-'),
//		parentheses ('()') or any other negative number
//		sign characters which might be associated with
//		a particular national or cultural usage.
//
//	decimalSeparator				DecimalSeparatorSpec
//
//		Type DecimalSeparatorSpec is used to specify the
//		radix point or decimal separator which will
//		separate integer and fractional digits in the
//		Target Number String Rune Array
//		('targetSearchString').
//
//		In the US, Canada and Australia, the period
//		('.'), or decimal point, separates integer and
//		fractional digits within a floating point numeric
//		value.
//
//		Many countries in Europe use the comma (',') to
//		separate integer and fractional digits within a
//		number string.
//
//	breakOnCharDelimiters			RuneArrayCollection
//
//		A collection of rune arrays used to specify
//		text characters which will automatically trigger
//		termination of the search and number parsing
//		operation.
//
//		If any one of these Terminator characters are
//		encountered while searching the Raw Number
//		String Rune Array ('rawNumStrRunes'), the
//		search operation will be immediately terminated.
//
//	requestRemainderRunesString		bool
//
//		If this parameter is set to 'true', the Number
//		String Search Results ('searchResults')
//		returned by this method	will include the
//		remaining string characters	which were NOT
//		included in the search operation.
//
//		This feature may be useful when either
//		'breakOnCharSearchLength' or
//		'breakOnCharDelimiters' parameters are used to
//		arbitrarily terminate the search operation.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	searchResults				CharSearchNumStrParseResultsDto
//
//		This structure contains public member variables
//		containing operational data describing the
//		success or failure of the number string parsing
//		operation and the search for numeric digits in
//		the Target Search String Rune Array,
//		'targetSearchString'.
//
//		Users are advised to check the following member
//		variable values after a number string parsing
//		operation:
//
//			FoundNumericDigits		bool
//
//				Signals a successful Number String Parsing
//				operation. When set to 'true', this means
//				one or more numeric	digit characters ('0'
//				through '9' inclusive) were	located in
//				the Target Search String Rune Array,
//				('targetSearchString').
//
//			FoundNonZeroValue		bool
//
//				When set to 'true' this signals that the
//				search operation has detected one or more
//				nonzero numeric digits.
//
//	numStrKernel				NumberStrKernel
//
//		If this method completes successfully, a new
//		instance of NumberStrKernel will be returned
//		configured with the numeric value parsed from the
//		Target Search String Rune Array,
//		('targetSearchString').
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (nStrHelperQuark *numStrHelperQuark) extractNumRunes(
	targetSearchString RuneArrayDto,
	targetSearchStringName string,
	startingSearchIndex int,
	breakOnCharSearchLength int,
	negativeNumSearchSpecsCol NegNumSearchSpecCollection,
	decimalSeparator DecimalSeparatorSpec,
	breakOnCharDelimiters RuneArrayCollection,
	requestRemainderRunesString bool,
	ePrefDto *ePref.ErrPrefixDto) (
	searchResults CharSearchNumStrParseResultsDto,
	numStrKernel NumberStrKernel,
	err error) {

	if nStrHelperQuark.lock == nil {
		nStrHelperQuark.lock = new(sync.Mutex)
	}

	nStrHelperQuark.lock.Lock()

	defer nStrHelperQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"numStrHelperQuark."+
			"extractNumRunes()",
		"")

	if err != nil {

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms := CharSearchTargetInputParametersDto{}.New()

	targetInputParms.TargetString = &targetSearchString

	targetInputParms.TargetStringLength = targetSearchString.GetRuneArrayLength()

	if targetInputParms.TargetStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetSearchString' is invalid.\n"+
			"'targetSearchString' has an array length of zero!\n",
			ePrefix.String())

		return searchResults,
			numStrKernel,
			err

	}

	sMechPreon := strMechPreon{}
	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		targetInputParms.TargetString.CharsArray,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid.\n"+
			"'%v' should contain valid characters.\n"+
			"A validity test on this rune array produced the following error:\n"+
			"%v\n",
			ePrefix.String(),
			targetSearchStringName,
			targetSearchStringName,
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.TargetInputParametersName = "Extract Number Runes"
	targetInputParms.TargetStringDescription2 =
		"strMechMolecule.extractNumRunes()"

	if len(targetSearchStringName) == 0 {
		targetInputParms.TargetStringName = "targetSearchString"
	} else {
		targetInputParms.TargetStringName = targetSearchStringName
	}

	targetInputParms.TargetStringLengthName =
		targetInputParms.TargetStringName + "Length"

	targetInputParms.TargetStringStartingSearchIndexName =
		targetInputParms.TargetStringName + "StartingSearchIndex"

	targetInputParms.TargetStringSearchLength = -1

	targetInputParms.FoundFirstNumericDigitInNumStr = false

	targetInputParms.TargetStringStartingSearchIndex =
		startingSearchIndex

	if targetInputParms.TargetStringStartingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid.\n"+
			"'startingSearchIndex' has a value less than zero!\n"+
			"startingSearchIndex = '%v'\n",
			ePrefix.String(),
			targetInputParms.TargetStringStartingSearchIndex)

		return searchResults,
			numStrKernel,
			err
	}

	if targetInputParms.TargetStringStartingSearchIndex >=
		targetInputParms.TargetStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'startingSearchIndex' is invalid.\n"+
			"'startingSearchIndex' has a value greater than the last index\n"+
			"of %v!\n"+
			"startingSearchIndex = '%v'\n"+
			"%v last index = %v\n",
			ePrefix.String(),
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringStartingSearchIndex,
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringLength-1)

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.TargetStringSearchLength = breakOnCharSearchLength

	err = targetInputParms.ValidateTargetParameters(
		ePrefix.XCpy(
			"targetInputParms"))

	if err != nil {

		return searchResults,
			numStrKernel,
			err

	}

	err2 = negativeNumSearchSpecsCol.IsValidInstanceError(
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeNumSearchSpecsCol' is invalid!\n"+
			"The following validation error was returned:\n"+
			"%v\n",
			ePrefix.XCpy(
				"negativeNumSearchSpecsCol"),
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	err2 = decimalSeparator.IsValidInstanceError(
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decimalSeparator' is invalid!\n"+
			"The following validation error was returned:\n"+
			"%v\n",
			ePrefix.XCpy(
				"decimalSeparator"),
			err2.Error())

		return searchResults,
			numStrKernel,
			err
	}

	targetInputParms.RequestRemainderString = false
	targetInputParms.RequestFoundTestCharacters = false
	targetInputParms.RequestReplacementString = false

	searchResults.Empty()

	searchResults.SearchResultsName = "Number Runes Extraction Results"

	searchResults.SearchResultsFunctionChain =
		ePrefix.String()

	searchResults.FoundNumericDigits = false
	searchResults.FoundNonZeroValue = false
	searchResults.FoundDecimalSeparatorSymbols = false
	searchResults.FoundDecimalDigits = false
	searchResults.NumSignValue = NumSignVal.Zero()
	searchResults.NumValueType = NumValType.None()
	searchResults.RemainderString.Empty()
	searchResults.DecimalSeparatorSearchResults.Empty()
	searchResults.NegativeNumberSymbolSearchResults.Empty()
	searchResults.ParsingTerminatorSearchResults.Empty()

	err = searchResults.TargetSearchString.CopyIn(
		&targetSearchString,
		ePrefix.XCpy(
			"searchResults.TargetSearchString"+
				"<-targetSearchString"))

	if err != nil {

		return searchResults,
			numStrKernel,
			err
	}

	searchResults.TargetStringSearchLength =
		targetInputParms.TargetStringSearchLength

	searchResults.TargetStringAdjustedSearchLength =
		targetInputParms.TargetStringAdjustedSearchLength

	searchResults.TargetStringStartingSearchIndex =
		targetInputParms.TargetStringStartingSearchIndex

	searchResults.TargetStringLastSearchIndex = -1

	searchResults.TargetStringNextSearchIndex = -1

	searchResults.ReasonForSearchTermination =
		CharSearchTermType.ProcessError()

	// Processing Flags

	// Number Parsing Setup
	searchResults.ParsingTerminatorSearchResults.IsNOP =
		breakOnCharDelimiters.IsNOP()

	searchResults.ParsingTerminatorSearchResults.SearchResultsName =
		"Number Parsing Terminator Search Results"

	numParsingTerminatorsIsNOP :=
		searchResults.ParsingTerminatorSearchResults.IsNOP

	// Decimal Separator Setup
	searchResults.DecimalSeparatorSearchResults.IsNOP =
		decimalSeparator.IsNOP()

	searchResults.DecimalSeparatorSearchResults.SearchResultsName =
		"Decimal Separator Search Results"

	decSeparatorIsNOP :=
		searchResults.DecimalSeparatorSearchResults.IsNOP

	var nextIdx int
	var tempDecSepSearchResults CharSearchDecimalSeparatorResultsDto

	for i := targetInputParms.TargetStringStartingSearchIndex; i < targetInputParms.TargetStringAdjustedSearchLength; i++ {

		targetInputParms.TargetStringCurrentSearchIndex = i

		if targetSearchString.CharsArray[i] >= '0' &&
			targetSearchString.CharsArray[i] <= '9' {

			searchResults.FoundNumericDigits = true
			targetInputParms.FoundFirstNumericDigitInNumStr = true

			if targetSearchString.CharsArray[i] > '0' {

				searchResults.FoundNonZeroValue = true
				targetInputParms.FoundNonZeroValue = true

				if searchResults.NumSignValue == NumSignVal.Zero() {
					searchResults.NumSignValue = NumSignVal.Positive()
				}
			}

			if !searchResults.DecimalSeparatorSearchResults.
				FoundDecimalSeparatorSymbols {

				err = numStrKernel.AddIntegerDigit(
					targetSearchString.CharsArray[i],
					ePrefix.XCpy(
						fmt.Sprintf(
							"targetSearchString.CharsArray[%v]",
							i)))

				if err != nil {
					return searchResults,
						numStrKernel,
						err
				}

				searchResults.FoundIntegerDigits = true

			} else {

				err = numStrKernel.AddFractionalDigit(
					targetSearchString.CharsArray[i],
					ePrefix.XCpy(
						fmt.Sprintf(
							"targetSearchString.CharsArray[%v]",
							i)))

				if err != nil {
					return searchResults,
						numStrKernel,
						err
				}

				searchResults.FoundDecimalDigits = true
			}

			continue
		}

		// Check for Parsing Terminators
		// All Parsing Operations Cease if Delimiter is Found
		if !numParsingTerminatorsIsNOP &&
			targetInputParms.FoundFirstNumericDigitInNumStr {

			searchResults.ParsingTerminatorSearchResults,
				err = breakOnCharDelimiters.SearchForTextCharacters(
				targetInputParms,
				ePrefix.XCpy(
					"breakOnCharDelimiters"))

			if searchResults.ParsingTerminatorSearchResults.FoundSearchTarget {

				i = searchResults.ParsingTerminatorSearchResults.
					TargetStringCurrentSearchIndex

				targetInputParms.TargetStringCurrentSearchIndex = i

				searchResults.ReasonForSearchTermination = CharSearchTermType.TerminationDelimiters()

				searchResults.ParsingTerminatorSearchResults.SearchResultsName =
					"Number Parsing Terminator Search Results"

				goto computeExitStats
			}
		}

		// Check for Negative Number Sign Symbol
		if !searchResults.NegativeNumberSymbolSearchResults.
			FoundNegativeNumberSymbols {

			searchResults.NegativeNumberSymbolSearchResults,
				err = negativeNumSearchSpecsCol.
				SearchForNegNumSignSymbols(
					targetInputParms,
					ePrefix.XCpy(
						"negativeNumSearchSpecsCol"))

			if searchResults.NegativeNumberSymbolSearchResults.
				FoundNegativeNumberSymbols {

				i =
					searchResults.NegativeNumberSymbolSearchResults.
						TargetStringCurrentSearchIndex

				targetInputParms.TargetStringCurrentSearchIndex = i

				targetInputParms.TargetStringNextSearchIndex = i + 1

				if targetInputParms.TargetStringNextSearchIndex >=
					targetInputParms.TargetStringLength {

					targetInputParms.TargetStringNextSearchIndex = -1

				}

				searchResults.NumSignValue = NumSignVal.Negative()

				if searchResults.NegativeNumberSymbolSearchResults.
					PrimaryNumSignPosition == NumSignSymPos.BeforeAndAfter() ||
					searchResults.NegativeNumberSymbolSearchResults.
						PrimaryNumSignPosition == NumSignSymPos.After() {

					goto computeExitStats
				}

				continue
			}
		}

		// Check for Decimal Separators
		if decSeparatorIsNOP == false &&
			targetInputParms.FoundDecimalSeparatorSymbols == false {

			if targetInputParms.FoundFirstNumericDigitInNumStr == false {

				nextIdx = i + 1

				if nextIdx < targetInputParms.TargetStringAdjustedSearchLength &&
					targetSearchString.CharsArray[nextIdx] >= '0' &&
					targetSearchString.CharsArray[nextIdx] <= '9' {

					tempDecSepSearchResults,
						err = decimalSeparator.SearchForDecimalSeparator(
						targetInputParms,
						ePrefix.XCpy(
							"decimalSeparator"))

					if err != nil {

						return searchResults,
							numStrKernel,
							err

					}

					if tempDecSepSearchResults.FoundDecimalSeparatorSymbols == true {

						err =
							searchResults.DecimalSeparatorSearchResults.CopyIn(
								&tempDecSepSearchResults,
								ePrefix.XCpy(
									"No Int Digits: "+
										"searchResults<-tempDecSepSearchResults"))

						if err != nil {

							return searchResults,
								numStrKernel,
								err

						}

						searchResults.FoundDecimalSeparatorSymbols = true

						targetInputParms.FoundDecimalSeparatorSymbols = true

						i = searchResults.DecimalSeparatorSearchResults.
							TargetStringLastSearchIndex

						targetInputParms.TargetStringCurrentSearchIndex = i

						targetInputParms.TargetStringNextSearchIndex = i + 1

						if targetInputParms.TargetStringNextSearchIndex >=
							targetInputParms.TargetStringLength {

							targetInputParms.TargetStringNextSearchIndex = -1

						}

						err = numStrKernel.AddIntegerDigit(
							'0',
							ePrefix.XCpy(
								fmt.Sprintf(
									"targetSearchString.CharsArray[%v]",
									i)))

						if err != nil {
							return searchResults,
								numStrKernel,
								err
						}

						searchResults.FoundNumericDigits = true
						targetInputParms.FoundFirstNumericDigitInNumStr = true
						searchResults.FoundIntegerDigits = true

						continue
					}

				}

			} else {
				// MUST BE
				// targetInputParms.FoundFirstNumericDigitInNumStr == true

				tempDecSepSearchResults,
					err = decimalSeparator.SearchForDecimalSeparator(
					targetInputParms,
					ePrefix.XCpy(
						"decimalSeparator"))

				if err != nil {

					return searchResults,
						numStrKernel,
						err

				}

				if tempDecSepSearchResults.FoundDecimalSeparatorSymbols == true {

					err =
						searchResults.DecimalSeparatorSearchResults.CopyIn(
							&tempDecSepSearchResults,
							ePrefix.XCpy(
								"Found1stNumDigit: "+
									"searchResults<-tempDecSepSearchResults"))

					if err != nil {

						return searchResults,
							numStrKernel,
							err

					}

					searchResults.FoundDecimalSeparatorSymbols = true

					targetInputParms.FoundDecimalSeparatorSymbols = true

					i = searchResults.DecimalSeparatorSearchResults.
						TargetStringLastSearchIndex

					targetInputParms.TargetStringCurrentSearchIndex = i

					targetInputParms.TargetStringNextSearchIndex = i + 1

					if targetInputParms.TargetStringNextSearchIndex >=
						targetInputParms.TargetStringLength {

						targetInputParms.TargetStringNextSearchIndex = -1

					}

					continue
				}

			}
		}
	}

	// End Of Target String Standard Loop Completion

	if targetInputParms.TargetStringAdjustedSearchLength <
		targetInputParms.TargetStringLength {

		searchResults.ReasonForSearchTermination =
			CharSearchTermType.SearchLengthLimit()

	} else {

		searchResults.ReasonForSearchTermination =
			CharSearchTermType.EndOfTargetString()

	}

computeExitStats:

	targetInputParms.TargetStringNextSearchIndex =
		targetInputParms.TargetStringCurrentSearchIndex + 1

	if targetInputParms.TargetStringNextSearchIndex >=
		targetInputParms.TargetStringLength {

		targetInputParms.TargetStringNextSearchIndex = -1

		searchResults.TargetStringNextSearchIndex = -1

		searchResults.TargetStringLastSearchIndex =
			targetInputParms.TargetStringLength - 1

	} else {
		// targetInputParms.TargetStringNextSearchIndex <
		// targetInputParms.TargetStringLength

		searchResults.TargetStringNextSearchIndex =
			targetInputParms.TargetStringNextSearchIndex

		searchResults.TargetStringLastSearchIndex =
			targetInputParms.TargetStringNextSearchIndex - 1

	}

	// Slice Examples
	//arr := []int{1,2,3,4,5}
	//
	//fmt.Println(arr[:2])        // [1,2]
	//
	//fmt.Println(arr[2:])        // [3,4,5]
	//
	//fmt.Println(arr[2:3])        // [3]
	//
	//fmt.Println(arr[:])            // [1,2,3,4,5]

	searchResults.RemainderString.Empty()

	if searchResults.FoundNumericDigits == true {

		if requestRemainderRunesString {

			if targetInputParms.TargetStringNextSearchIndex == -1 {
				// All characters have been searched.
				// There is no Remainder String
			} else {
				// All characters HAVE NOT BEEN SEARCHED
				// There is a Remainder String

				searchResults.RemainderString.CharsArray =
					append(
						searchResults.RemainderString.CharsArray,
						targetSearchString.
							CharsArray[targetInputParms.TargetStringCurrentSearchIndex:]...)
			}

		}

		if numStrKernel.GetNumberOfIntegerDigits() > 0 {

			searchResults.IdentifiedIntegerDigits =
				numStrKernel.GetIntegerString()
		}

		if numStrKernel.GetNumberOfFractionalDigits() > 0 {

			searchResults.IdentifiedFractionalDigits =
				numStrKernel.GetFractionalString()
		}

		if numStrKernel.GetNumberOfFractionalDigits() > 0 &&
			numStrKernel.GetNumberOfIntegerDigits() == 0 {

			// Rationalize Fractional Digits
			err = numStrKernel.AddIntegerDigit(
				'0',
				ePrefix.XCpy("Add '0' int digit"))

			if err != nil {

				return searchResults,
					numStrKernel,
					err
			}

		}

		if searchResults.FoundDecimalDigits {

			searchResults.NumValueType = NumValType.FloatingPoint()

		} else {

			searchResults.NumValueType = NumValType.Integer()

		}

		if searchResults.FoundNonZeroValue == true {
			// Value is Nonzero

			if searchResults.NumSignValue == NumSignVal.Negative() {
				numStrKernel.numberSign = NumSignVal.Negative()
			} else {
				numStrKernel.numberSign = NumSignVal.Positive()
			}

		} else {
			// Value is Zero
			numStrKernel.numberSign = NumSignVal.Zero()
		}

	} else {
		// searchResults.FoundSearchTarget == false
		// Didn't find any numeric digits!
		numStrKernel.Empty()
		searchResults.FoundNumericDigits = false
		searchResults.FoundDecimalDigits = false
		searchResults.FoundNonZeroValue = false

		searchResults.RemainderString.CharsArray =
			targetSearchString.CharsArray[:]

	}

	return searchResults,
		numStrKernel,
		err
}

//	parsePureNumStr
//
//	Receives a pure number string and proceeds to return the
//	extracted numeric value as a type NumberStrKernel.
//
//	This method is particularly useful when numeric values
//	are converted to string using 'fmt.Sprintf()' and
//	similar formatting algorithms.
//
//	A "Pure Number String" is defined as follows:
//
//		1.	Consists of numeric character digits
//			zero through nine inclusive (0-9).
//
//		2.	Option: A Pure Number String may include
//			a radix point or decimal separator. The
//			decimal separator may consist of one or
//			more characters.
//
//			In the US, UK, Australia and most of Canada,
//			the decimal separator is the period
//			character ('.') known as the decimal point.
//
//		3.	Optional: A Pure Number String may include a
//			negative number sign symbol consisting of a
//			minus sign ('-'). Only the minus sign ('-')
//			classifies the numeric value as a negative
//			number in Pure Number String.
//
//			If the leading or trailing minus sign ('-')
//			is NOT present, the numeric value is assumed
//			to be positive.
//
//		4.	Only numeric characters, the decimal
//			separator and the minus sign will be
//			processed by the number string parsing
//			algorithm. All other characters will be
//			ignored.
//
//		5.	Pure Number Strings consist of a single
//			numeric value. The entire Pure Number String
//			will be parsed, or processed, and only one
//			numeric value per Pure Number String will
//			be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberString			RuneArrayDto
//
//		This rune array contains the character digits from
//		which the numeric value will be extracted and
//		returned as a NumberStrKernel.
//
//		A "Pure Number String" is defined as follows:
//
//			1.	Consists of numeric character digits
//				zero through nine inclusive (0-9).
//
//			2.	Option: A Pure Number String may include
//				a radix point or decimal separator.
//				Decimal separators separate integer and
//				fractional numeric digits in a pure
//				number string. The decimal separator may
//				consist of one or more text characters.
//
//				In the US, UK, Australia, most of Canada
//				and many other countries, the decimal
//				separator is the period character ('.')
//				known as the decimal point.
//
//				In France, Germany and many countries in
//				the European Union, the Decimal Separator
//				is the comma character (',').
//
//			3.	Optional: A Pure Number String may
//				include a negative number sign symbol
//				consisting of a minus sign ('-'). The
//				minus sign will identify the numeric
//				value contained in the pure number string
//				as a negative number. Only the minus sign
//				('-') classifies a numeric value as a
//				negative number in a Pure Number String.
//
//				If a leading or trailing minus sign ('-')
//				is NOT present in the pure number string,
//				the numeric value is assumed to be
//				positive.
//
//			4.	Only numeric characters, the decimal
//				separator and the minus sign will be
//				processed by the pure number string
//				parsing algorithm. All other characters
//				will be	ignored.
//
//			5.	Pure Number Strings consist of a single
//				numeric value. The entire Pure Number String
//				will be parsed, or processed, and only one
//				numeric value per Pure Number String will
//				be returned.
//
//	decSeparatorSpec				DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingMinusSign			bool
//
//		In pure number strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		pure number string parsing algorithm will search
//		for a leading minus sign ('-') at the beginning of
//		the number string. Leading minus signs represent
//		the standard means for designating negative
//		numeric values in the US, UK, Australia, most of
//		Canada and many other parts of world.
//
//		Example Leading Minus Sign:
//			"-123.456" or "- 123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		pure number string parsing algorithm will search
//		for trailing minus signs ('-') located at the end
//		of the number string. Trailing minus signs
//		represent the standard for France, Germany and
//		many countries in the European Union.
//
//		NOTE: Identification of a trailing minus sign in
//		the pure number string input parameter,
//		'pureNumberString', will immediately terminate
//		the search for numeric characters.
//
//		Example Trailing Number Symbols:
//			"123.456-" or "123.456 -"
//
//	errPrefDto          		*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numStrKernel				NumberStrKernel
//
//		If this method completes successfully, an instance of
//		NumberStrKernel containing the numeric value extracted
//		from parameter 'pureNumberString' will be returned.
//
//	err							error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error Type
//		will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (nStrHelperQuark *numStrHelperQuark) parsePureNumStr(
	pureNumberString RuneArrayDto,
	decSeparatorSpec DecimalSeparatorSpec,
	leadingMinusSign bool,
	ePrefDto *ePref.ErrPrefixDto) (
	numStrKernel NumberStrKernel,
	err error) {

	if nStrHelperQuark.lock == nil {
		nStrHelperQuark.lock = new(sync.Mutex)
	}

	nStrHelperQuark.lock.Lock()

	defer nStrHelperQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"numStrHelperQuark."+
			"parsePureNumStr()",
		"")

	if err != nil {

		return numStrKernel, err
	}

	lenPureNStr := pureNumberString.GetRuneArrayLength()

	if lenPureNStr == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'pureNumberString' is empty\n"+
			"and contains zero characters.\n",
			ePrefix.String())

		return numStrKernel, err
	}

	var numberStats NumberStrStatsDto

	numberStats,
		err = new(numStrMathQuark).
		pureNumStrToRunes(
			string(pureNumberString.CharsArray),
			&numStrKernel.integerDigits,
			&numStrKernel.fractionalDigits,
			&decSeparatorSpec.decimalSeparatorChars,
			leadingMinusSign,
			ePrefix)

	if err != nil {

		return numStrKernel, err
	}

	numStrKernel.numberSign = numberStats.NumberSign

	numStrKernel.numberValueType = numberStats.NumberValueType

	numStrKernel.isNonZeroValue = !numberStats.IsZeroValue

	return numStrKernel, err
}

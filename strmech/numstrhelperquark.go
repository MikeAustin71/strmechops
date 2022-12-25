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

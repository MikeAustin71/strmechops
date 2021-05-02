package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"regexp"
	"sort"
	"strings"
	"sync"
	"unicode/utf8"
)

type strMechQuark struct {
	lock *sync.Mutex
}

// convertNonPrintableChars - Receives a string containing
// non-printable characters and converts them to 'printable'
// characters returned in a string.
//
// Examples of non-printable characters are '\n', '\t' or 0x06
// (Acknowledge). These example characters would be translated into
// printable string characters as: "\\n", "\\t" and "[ACK]".
//
// Space characters are typically translated as " ". However, if
// the input parameter 'convertSpace' is set to 'true' then all
// spaces are converted to "[SPACE]" in the returned string.
//
// Reference:
//    https://www.juniper.net/documentation/en_US/idp5.1/topics/reference/general/intrusion-detection-prevention-custom-attack-object-extended-ascii.html
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  nonPrintableChars   []rune
//     - An array of runes containing non-printable characters.
//       The non-printable characters will be converted to
//       printable characters.
//
//  convertSpace        bool
//     - Space or white space characters (0x20) are by default
//       translated as " ". However, if this parameter is set to
//       'true', space characters will be converted to "[SPACE]".
//
//
//  ePrefix             *ErrPrefixDto
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
//  printableChars      string
//     - This returned string is identical to input parameter
//       'nonPrintableChars' with the exception that non-printable
//       characters are translated into printable characters.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  testStr := "Hello world!\n"
//  testRunes := []rune(testStr)
//
//  actualStr :=
//    StrMech{}.NewPtr().
//      ConvertNonPrintableChars(testRunes, true)
//
//  ----------------------------------------------------
//  'actualStr' is now equal to:
//     "Hello[SPACE]world!\\n"
//
func (sMechQuark *strMechQuark) convertNonPrintableChars(
	nonPrintableChars []rune,
	convertSpace bool,
	ePrefix *ePref.ErrPrefixDto) (
	printableChars string,
	err error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"convertNonPrintableChars()")

	lenNonPrintableChars := len(nonPrintableChars)

	if lenNonPrintableChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nonPrintableChars' is "+
			"invalid!\n"+
			"'nonPrintableChars' is a zero length array of "+
			"runes.\n",
			ePrefix.String())

		return "[EMPTY]", err
	}

	var b strings.Builder
	b.Grow(lenNonPrintableChars * 5)

	for i := 0; i < lenNonPrintableChars; i++ {
		cRune := nonPrintableChars[i]

		switch cRune {
		case 0:
			b.WriteString("[NULL]") // 0x00 NULL
		case 1:
			b.WriteString("[SOH]") // 0x01 State of Heading
		case 2:
			b.WriteString("[STX]") // 0x02 State of Text
		case 3:
			b.WriteString("[ETX]") // 0x03 End of Text
		case 4:
			b.WriteString("[EOT]") // 0X04 End of Transmission
		case 5:
			b.WriteString("[ENQ]") // 0x05 Enquiry
		case 6:
			b.WriteString("[ACK]") // 0x06 Acknowledge
		case '\a':
			b.WriteString("\\a") // U+0007 alert or bell
		case '\b':
			b.WriteString("\\b") // U+0008 backspace
		case '\t':
			b.WriteString("\\t") // U+0009 horizontal tab
		case '\n':
			b.WriteString("\\n") // U+000A line feed or newline
		case '\v':
			b.WriteString("\\v") // U+000B vertical tab
		case '\f':
			b.WriteString("\\f") // U+000C form feed
		case '\r':
			b.WriteString("\\r") // U+000D carriage return
		case 14:
			b.WriteString("[SO]") // U+000E Shift Out
		case 15:
			b.WriteString("[SI]") // U+000F Shift In
		case '\\':
			b.WriteString("\\") // U+005c backslash
		case ' ':
			// 0X20 Space character
			if convertSpace {
				b.WriteString("[SPACE]")
			} else {
				b.WriteRune(' ')
			}

		default:
			b.WriteRune(cRune)

		}

	}

	return b.String(), err

}

// convertPrintableChars - Converts selected printable characters
// to their non-printable or native equivalent. For example,
// instances of '\\n' in a string will be converted to '\n'.
//
// Additional examples of converted printable string characters
// are: "\\n", "\\t" and "[ACK]". These printable characters be
// converted into their native, non-printable state: '\n', '\t' or
// 0x06 (Acknowledge).
//
// This method, strMechQuark.convertPrintableChars(), performs the
// mirror operation to that performed by method
// strMechQuark.convertNonPrintableChars().
//
// strMechQuark.convertNonPrintableChars() converts non-printable
// characters into printable characters.
//
// strMechQuark.convertPrintableChars() preforms in just the
// opposite manner. It converts printable characters back into
// non-printable characters.
//
// If strMechQuark.convertNonPrintableChars() is called on a string
// containing non-printable characters, calling
// strMechQuark.convertPrintableChars() on the resulting string will
// reverse the operation and return the string to its original
// content.
//
// Reference:
//    https://www.juniper.net/documentation/en_US/idp5.1/topics/reference/general/intrusion-detection-prevention-custom-attack-object-extended-ascii.html
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  printableChars      string
//     - A string which may contain non-printable characters converted
//       to their printable equivalents. These printable characters will
//       be converted back to their native, non-printable values.
//
//
//  ePrefix             *ErrPrefixDto
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
//  nonPrintableChars   []rune
//     - An array of runes containing non-printable characters.
//       The non-printable characters were be converted from the
//       printable characters contained in input parameter
//       'printableChars'.
//
//
//  err                 error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  testStr := "Hello[SPACE]world!\\n"
//  ePrefix := "theCallingFunction()"
//
//  ePrefQuark := errPrefQuark{}
//
//  actualRuneArray :=
//    ePrefQuark.
//      convertPrintableChars(
//           testStr,
//           ePrefix)
//
//  ----------------------------------------------------
//  'actualRuneArray' is now equal to:
//     "Hello world!\n"
//
func (sMechQuark *strMechQuark) convertPrintableChars(
	printableChars string,
	ePrefix *ePref.ErrPrefixDto) (
	nonPrintableChars []rune,
	err error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"convertPrintableChars()")

	nonPrintableChars = make([]rune, 0)

	lenPrintableChars := len(printableChars)

	if lenPrintableChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'printableChars' is invalid!\n"+
			"'printableChars' is an empty or zero lenght string.\n",
			ePrefix)

		return nonPrintableChars, err
	}

	printableChars = strings.Replace(
		printableChars, "[SPACE]", " ", -1)

	printableChars = strings.Replace(
		printableChars, "[NULL]", string(rune(0x00)), -1) // 0x00 NULL

	printableChars = strings.Replace(
		printableChars, "[SOH]", string(rune(0x01)), -1) // 0x01 State of Heading

	printableChars = strings.Replace(
		printableChars, "[STX]", string(rune(0x02)), -1) // 0x02 State of Text

	printableChars = strings.Replace(
		printableChars, "[ETX]", string(rune(0x03)), -1) // 0x03 End of Text

	printableChars = strings.Replace(
		printableChars, "[EOT]", string(rune(0x04)), -1) // 0X04 End of Transmission

	printableChars = strings.Replace(
		printableChars, "[ENQ]", string(rune(0x05)), -1) // 0x05 Enquiry

	printableChars = strings.Replace(
		printableChars, "[ACK]", string(rune(0x06)), -1) // 0x06 Acknowledge

	printableChars = strings.Replace(
		printableChars, "\\a", string(rune(0x07)), -1) // U+0007 alert or bell

	printableChars = strings.Replace(
		printableChars, "\\b", string(rune(0x08)), -1) // U+0008 backspace

	printableChars = strings.Replace(
		printableChars, "\\t", string(rune(0x09)), -1) // U+0009 horizontal tab

	printableChars = strings.Replace(
		printableChars, "\\n", string(rune(0x0A)), -1) // U+000A line feed or newline

	printableChars = strings.Replace(
		printableChars, "\\v", string(rune(0x0B)), -1) // U+000B vertical tab

	printableChars = strings.Replace(
		printableChars, "\\f", string(rune(0x0C)), -1) // U+000C form feed

	printableChars = strings.Replace(
		printableChars, "\\r", string(rune(0x0D)), -1) // U+000D carriage return

	printableChars = strings.Replace(
		printableChars, "[SO]", string(rune(0x0E)), -1) // U+000E Shift Out

	printableChars = strings.Replace(
		printableChars, "[SI]", string(rune(0x0F)), -1) // U+000F Shift In

	printableChars = strings.Replace(
		printableChars, "\\", string(rune(0x5c)), -1) // U+005c backslash

	nonPrintableChars = []rune(printableChars)

	return nonPrintableChars, err
}

// doesLastCharExist - returns true if the last character (rune) of
// input string 'testStr' is equal to input parameter 'lastChar' which
// is of type 'rune'.
//
func (sMechQuark *strMechQuark) doesLastCharExist(
	testStr string,
	lastChar rune) bool {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()
	testStrLen := len(testStr)

	if testStrLen == 0 {
		return false
	}

	strLastChar := rune(testStr[testStrLen-1])

	if strLastChar == lastChar {
		return true
	}

	return false

}

// findLastNonSpaceChar - Returns the string index of the last non-space character in a
// string segment.  The string to be searched is input parameter, 'targetStr'. The
// string segment is further defined by input parameters 'startIdx' and  'endIdx'. These
// indexes define a segment within 'targetStr' which will be searched to identify the last
// non-space character.
//
// The search is a backwards search, from right to left, conducted within the defined
// 'targetStr' segment. The search therefore starts at 'endIdx' and proceeds towards
// 'startIdx' until the last non-space character in the string segment is identified.
//
//
// If the last non-space character is found, that string index is returned. If the string
// segment consists entirely of space characters, the return value is -1.
//
// if 'targetStr' is a zero length string, an error will be triggered. Likewise, if 'startIdx'
// of 'endIdx' are invalid, an error will be returned.
//
func (sMechQuark *strMechQuark) findLastNonSpaceChar(
	targetStr string,
	startIdx,
	endIdx int,
	ePrefix *ePref.ErrPrefixDto) (
	int,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"findLastNonSpaceChar()")

	targetStrLen := len(targetStr)

	if targetStrLen == 0 {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'targetStr' is a ZERO LENGTH STRING!\n",
				ePrefix.String())
	}

	if startIdx < 0 {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'startIdx' is LESS THAN ZERO!\n"+
				"startIdx='%v'\n",
				ePrefix.String(),
				startIdx)
	}

	if endIdx < 0 {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'endIdx' is LESS THAN ZERO!\n"+
				"startIdx='%v'\n",
				ePrefix.String(),
				startIdx)
	}

	if endIdx >= targetStrLen {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'endIdx' is greater than target string length.\n"+
				"INDEX OUT OF RANGE!\n"+
				"endIdx='%v' target string length='%v'\n",
				ePrefix.String(),
				endIdx,
				targetStrLen)
	}

	if startIdx > endIdx {
		return -1, fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter.\n"+
			"'startIdx' is GREATER THAN 'endIdx'\n"+
			"startIdx='%v' endIdx='%v'\n",
			ePrefix.String(),
			startIdx, endIdx)
	}

	floor := startIdx - 1

	for endIdx > floor {

		if targetStr[endIdx] != ' ' {
			return endIdx, nil
		}

		endIdx--
	}

	return -1, nil
}

// FindLastSpace - Returns a string index indicating the last space character (' ') in
// a string segment. The string segment is defined by input parameters, 'startIdx' and
// 'endIdx'.
//
// The string segment search proceeds backwards, from right to left. The search therefore
// starts at 'endIdx' and proceeds towards 'startIdx' until the last space character in
// the string segment is identified.
//
// If a valid index for the last space character is found in the string segment, that
// index value is returned. If a space character is NOT found in the specified string
// segment, a value of -1 is returned.
//
// if 'targetStr' is a zero length string, an error will be triggered. Likewise, if 'startIdx'
// of 'endIdx' are invalid, an error will be returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string segment to be searched for the last space
//       character resides in this string.
//
//  startIdx            int
//     - Defines the actual ending index for the string segment to
//       be searched. The search is a 'backwards' search and
//       proceeds from right to left. Therefore, the starting point
//       for the string segment is input parameter 'endIdx' while
//       the ending point for the string segment is this 'startIdx'.
//
//
//  endIdx              int
//     - Defines the actual beginning index of the string segment
//       to be searched. The search is a 'backwards' search and
//       proceeds from right to left. Therefore, the starting point
//       for the string segment is defined by this 'endIdx'
//       parameter while the ending point for the string segment is
//       marked by the input parameter, 'startIdx'.
//
//
//  ePrefix             *ErrPrefixDto
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
//  int
//     - If this method completes successfully, this returned
//       integer value will constitute the string index of the last
//       space character in the string segment marked by input
//       parameters 'startIdx' and 'endIdx'.
//
//       If a space character is NOT found in the specified string
//       segment, a value of minus one (-1) is returned.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) findLastSpace(
	targetStr string,
	startIdx int,
	endIdx int,
	ePrefix *ePref.ErrPrefixDto) (
	int,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"findLastSpace()")

	targetStrLen := len(targetStr)

	if targetStrLen == 0 {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'targetStr' is a ZERO LENGTH STRING!\n",
				ePrefix.String())
	}

	if startIdx < 0 {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'startIdx' is LESS THAN ZERO!\n"+
				"startIdx='%v'\n",
				ePrefix.String(),
				startIdx)
	}

	if endIdx < 0 {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'endIdx' is LESS THAN ZERO!\n"+
				"startIdx='%v'\n",
				ePrefix.String(),
				startIdx)
	}

	if endIdx >= targetStrLen {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter. 'endIdx' is greater than target string length.\n"+
				"INDEX OUT OF RANGE!\n"+
				"endIdx='%v'\n"+
				"target string length='%v'\n",
				ePrefix.String(),
				endIdx,
				targetStrLen)
	}

	if startIdx > endIdx {
		return -1,
			fmt.Errorf("%v\n"+
				"ERROR: Invalid input parameter.\n"+
				"'startIdx' is GREATER THAN 'endIdx'\n"+
				"startIdx='%v' endIdx='%v'\n",
				ePrefix.String(),
				startIdx, endIdx)
	}

	for endIdx >= startIdx {

		if targetStr[endIdx] == ' ' {
			return endIdx, nil
		}

		endIdx--
	}

	return -1, nil
}

// findLastWord - Returns the beginning and ending indexes of
// the last word in a target string segment. A 'word' is defined here
// as a contiguous set of non-space characters delimited by spaces or
// the beginning and ending indexes of the target string segment. Note,
// for purposes of this method, a 'word' my consist of a single non-space
// character such as an article 'a' or a punctuation mark '.'
//
// ------------------------------------------------------------------------
//
// Examples
//
//
//   Example (1)
//     In the text string segment:
//
//     "The cow jumped over the moon."
//
//     The last word would be defined as "moon."
//
//     Example (2)
//       In the text string segment:
//
//       "  somewhere over the rainbow  "
//
//       The last word would be defined as "rainbow"
//
// ------------------------------------------------------------------------
//
// The string to be searched is contained in input parameter, 'targetStr'.
// The string segment within 'targetStr' is defined by input parameters
// 'startIndex' and 'endIndex'.
//
// If the entire string segment is classified as a 'word', meaning that
// there are no space characters in the string segment, the returned
// values for 'beginWrdIdx' and 'endWrdIdx' will be equal to the input
// parameters 'startIndex' and 'endIndex'.
//
// If the string segment is consists entirely of space characters, the
// returned 'beginWrdIdx' and 'endWrdIdx' will be set equal to -1 and
// the returned value, 'isAllSpaces' will be set to 'true'.
//
// If 'targetStr' is an empty string, an error will be returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetStr      string
//     - The string containing the string segment which
//       will be searched to identify the last word
//       in the string segment.
//
//
//  startIndex     int
//     - The index marking the beginning of the string
//       segment in 'targetStr'.
//
//
//  endIndex       int
//     - The index marking the end of the string segment
//       in 'targetStr'.
//
//
//  ePrefix             *ErrPrefixDto
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
//  beginWrdIdx         int
//     - The index marking the beginning of the last word
//       in the string segment identified by input parameters
//       'startIndex' and 'endIndex'. If the string segment
//       consists of all spaces or is empty, this value is
//       set to minus one (-1).
//
//  endWrdIdx           int
//     - The index marking the end of the last word in the
//       string segment identified by input parameters 'startIndex'
//       and 'endIndex'. If the string segment consists of all
//       spaces or is empty, this value is set to -1.
//
//  isAllOneWord        bool
//     - If the string segment identified by input parameters
//       'startIndex' and 'endIndex' consists entirely of non-space
//       characters (characters other than ' '), this value is set
//       to 'true'.
//
//  isAllSpaces         bool
//     - If the string segment identified by input parameters
//       'startIndex' and 'endIndex' consists entirely of space
//       characters (character = ' '), this value is set to 'true'.
//
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) findLastWord(
	targetStr string,
	startIndex int,
	endIndex int,
	ePrefix *ePref.ErrPrefixDto) (
	beginWrdIdx int,
	endWrdIdx int,
	isAllOneWord bool,
	isAllSpaces bool,
	err error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"findLastWord()")

	beginWrdIdx = -1
	endWrdIdx = -1
	isAllOneWord = false
	isAllSpaces = false

	targetStrLen := len(targetStr)

	if targetStrLen == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStr' is an EMPTY STRING!\n",
			ePrefix.String())

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if startIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter.\n"+
			"'startIndex' is LESS THAN ZERO!\n"+
			"startIndex='%v'\n",
			ePrefix.String(),
			startIndex)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if endIndex < 0 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter.\n"+
			"'endIndex' is LESS THAN ZERO!\n"+
			"startIndex='%v'\n",
			ePrefix.String(),
			startIndex)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if endIndex >= targetStrLen {

		err = fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter. 'endIndex' is greater than\n"+
			"target string length. INDEX OUT OF RANGE!\n"+
			"endIndex='%v'\n"+
			"target string length='%v'\n",
			ePrefix.String(),
			endIndex,
			targetStrLen)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if startIndex > endIndex {
		err = fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter.\n"+
			"'startIndex' is GREATER THAN 'endIndex'.\n"+
			"startIndex='%v' endIndex='%v'\n",
			ePrefix.String(),
			startIndex,
			endIndex)

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	beginWrdIdx = startIndex
	endWrdIdx = endIndex

	idx := endIndex

	var endingIdxFound bool

	isAllSpaces = true
	isAllOneWord = true

	if startIndex == endIndex {

		beginWrdIdx = startIndex
		endWrdIdx = startIndex

		if targetStr[startIndex] == ' ' {
			isAllSpaces = true
			isAllOneWord = false
		} else {
			isAllSpaces = false
			isAllOneWord = true
		}

		err = nil

		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	for idx >= startIndex {

		if targetStr[idx] != ' ' {
			isAllSpaces = false
		} else {
			isAllOneWord = false
		}

		if !endingIdxFound &&
			targetStr[idx] != ' ' {

			endWrdIdx = idx
			endingIdxFound = true
			idx--
			continue
		}

		if endingIdxFound &&
			targetStr[idx] == ' ' {

			beginWrdIdx = idx + 1
			break
		}

		idx--
	}

	if isAllSpaces {
		isAllOneWord = false
		beginWrdIdx = -1
		endWrdIdx = -1
		err = nil
		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	if isAllOneWord {
		beginWrdIdx = startIndex
		endWrdIdx = endIndex
		isAllSpaces = false
		err = nil
		return beginWrdIdx,
			endWrdIdx,
			isAllOneWord,
			isAllSpaces,
			err
	}

	err = nil

	return beginWrdIdx,
		endWrdIdx,
		isAllOneWord,
		isAllSpaces,
		err
}

// findRegExIndex - returns a two-element slice of integers defining the location
// of the leftmost match in targetStr of the regular expression (regex).
//
// ------------------------------------------------------------------------
//
// Return Value
//
// The return value is an array of integers. If no match is found the return
// value is 'nil'.  If regular expression is successfully matched, the match
// will be located at targetStr[loc[0]:loc[1]]. Again, a return value of 'nil'
// signals that no match was found.
//
func (sMechQuark *strMechQuark) findRegExIndex(
	targetStr string,
	regex string) []int {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	re := regexp.MustCompile(regex)

	return re.FindStringIndex(targetStr)
}

// getCharCountInStr - Uses the 'len' method to return the number
// of rune characters in a string.
//
// There is a difference in counting characters or bytes in a string
// and counting 'runes' in a string.
//
// For information on an alternative counting method, see
// strMechQuark.getRuneCountInStr() and research the 'utf8' Rune
// Count In String function at:
//
//   https://golang.org/pkg/unicode/utf8/#RuneCountInString
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The number of 'runes' in this string will be counted and
//       that count will be returned to the calling function.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This returned integer value will be set to the number of
//       of 'runes' counted in 'targetStr'.
//
func (sMechQuark *strMechQuark) getCharCountInStr(
	targetStr string) int {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	return len([]rune(targetStr))
}

// getRuneCountInStr - Uses utf8 Rune Count function to return the
// number of characters, or runes, within a string.
//
// There is a difference in counting characters or bytes in a string
// and counting 'runes' in a string.
//
// For more information on the 'utf8' Rune Count In String function,
// Reference:
//   https://golang.org/pkg/unicode/utf8/#RuneCountInString
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The number of 'runes' in this string will be counted and
//       that count will be returned to the calling function.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  int
//     - This returned integer value will be set to the number of
//       of 'runes' counted in 'targetStr'.
//
func (sMechQuark *strMechQuark) getRuneCountInStr(
	targetStr string) int {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	return utf8.RuneCountInString(targetStr)
}

// getValidBytes - Receives an array of 'targetBytes' which will be examined to determine
// the validity of individual bytes or characters. Each character (byte) in input array
// 'targetBytes' will be compared to input parameter 'validBytes', another array of bytes.
// If a character in 'targetBytes' also exists in 'validBytes' it will be considered valid
// and included in the returned array of bytes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetBytes         [] byte
//     - An array of characters (bytes) which will be examined
//       for valid characters. The list of valid characters is
//       found in input parameter 'validBytes'. Valid characters
//       in targetBytes will be returned by this method as an
//       array of bytes. Invalid characters will be discarded.
//
//
//  validBytes          [] byte
//     - An array of bytes containing valid characters. If a character
//       (byte) in 'targetBytes' is also present in 'validBytes' it will
//       be classified as 'valid' and included in the returned array of
//       bytes. Invalid characters will be discarded.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods listed
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
//  [] byte
//     - An array of bytes which contains bytes that are
//       present in both 'targetBytes' and 'validBytes'.
//       Note: If all characters in 'targetBytes' are classified
//       as 'invalid', the returned array of bytes will be a zero
//       length array.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) getValidBytes(
	targetBytes []byte,
	validBytes []byte,
	ePrefix *ePref.ErrPrefixDto) (
	[]byte,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"getValidBytes()")

	lenTargetBytes := len(targetBytes)

	output := make([]byte, 0, lenTargetBytes+10)

	if lenTargetBytes == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetBytes' is a "+
				"ZERO LENGTH ARRAY!\n",
				ePrefix.String())
	}

	lenValidBytes := len(validBytes)

	if lenValidBytes == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'validBytes' is a "+
				"ZERO LENGTH ARRAY!\n",
				ePrefix.String())
	}

	for i := 0; i < lenTargetBytes; i++ {

		for j := 0; j < lenValidBytes; j++ {
			if targetBytes[i] == validBytes[j] {
				output = append(output, targetBytes[i])
				break
			}
		}

	}

	return output, nil
}

// getValidRunes - Receives an array of 'targetRunes' which will be examined to determine
// the validity of individual runes or characters. Each character (rune) in input array
// 'targetRunes' will be compared to input parameter 'validRunes', another array of runes.
// If a character in 'targetRunes' also exists in 'validRunes', that character will be considered
// valid and included in the returned array of runes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetRunes    [] rune  - An array of characters (runes) which will be examined
//                            for valid characters. The list of valid characters is
//                            found in input parameter 'validRunes'. Valid characters
//                            in targetRunes will be returned by this method as an
//                            array of runes. Invalid characters will be discarded.
//
//  validRunes    [] rune  - An array of runes containing valid characters. If a character
//                           (rune) in targetRunes is also present in 'validRunes' it will
//                           be classified as 'valid' and included in the returned array of
//                           runes. Invalid characters will be discarded.
//
//
//  ePrefix             *ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the names of the calling method or methods listed
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
//  [] rune  - An array of runes which contains runes that are present in 'targetRunes' and
//             'validRunes'. Note: If all characters in 'targetRunes' are classified as
//             'invalid', the returned array of runes will be a zero length array.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) getValidRunes(
	targetRunes []rune,
	validRunes []rune,
	ePrefix *ePref.ErrPrefixDto) (
	[]rune,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"getValidRunes()")

	lenTargetRunes := len(targetRunes)

	output := make([]rune, 0, lenTargetRunes+10)

	if lenTargetRunes == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetRunes' is a "+
				"ZERO LENGTH ARRAY!\n",
				ePrefix.String())
	}

	lenValidRunes := len(validRunes)

	if lenValidRunes == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'validRunes' is a "+
				"ZERO LENGTH ARRAY!\n",
				ePrefix.String())
	}

	for i := 0; i < lenTargetRunes; i++ {

		for j := 0; j < lenValidRunes; j++ {
			if targetRunes[i] == validRunes[j] {
				output = append(output, targetRunes[i])
				break
			}
		}
	}

	return output, nil
}

// isEmptyOrWhiteSpace - Analyzes the incoming string and returns
// 'true' if the strings is empty or consists of all white space.
//
func (sMechQuark *strMechQuark) isEmptyOrWhiteSpace(
	targetStr string) bool {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	targetLen := len(targetStr)

	for i := 0; i < targetLen; i++ {
		if targetStr[i] != ' ' {
			return false
		}
	}

	return true
}

// LowerCaseFirstLetter - Finds the first alphabetic character
// in a string (a-z A-Z) and converts it to lower case.
//
func (sMechQuark *strMechQuark) lowerCaseFirstLetter(
	str string) string {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if len(str) == 0 {
		return str
	}

	runeStr := []rune(str)

	for i := 0; i < len(runeStr); i++ {

		if runeStr[i] == ' ' {
			continue
		}

		if runeStr[i] >= 'A' &&

			runeStr[i] <= 'Z' {

			runeStr[i] += 32

			break

		} else if runeStr[i] >= 'a' &&

			runeStr[i] <= 'z' {

			break
		}

	}

	return string(runeStr)
}

// MakeSingleCharString - Creates a string of length 'strLen' consisting of
// a single character passed through input parameter, 'charRune' as type
// 'rune'.
//
//
// Example Usage:
//
//     sUtil := StrMech{}
//     requestedLen := 5
//     charRune := '='
//     outputStr, err := sUtil.MakeSingleCharString(charRune, requestedLen)
//
//     outputStr is now equal to "====="
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  charRune            rune
//     - The text character which will be repeated throughout the
//       length of the returned string.
//
//
//  strLen              int
//     - The length of the returned string containing the repeated
//       characters.
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - This returned string will have a length of 'strLen' and
//       contain 'strLen' characters all of which will be equal to
//       'charRune'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) makeSingleCharString(
	charRune rune,
	strLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"makeSingleCharString()")

	if strLen < 1 {
		return "",
			fmt.Errorf("%s"+
				"Error: Input parameter 'strLen' "+
				"MUST BE GREATER THAN '1'.\n"+
				"strLen='%v'\n",
				ePrefix.String(),
				strLen)
	}

	if charRune == 0 {
		return "",
			fmt.Errorf("%s\n"+
				"Error: Input parameter 'charRune' IS INVALID!\n"+
				"charRune='%v'\n",
				ePrefix.String(),
				charRune)
	}

	var b strings.Builder
	b.Grow(strLen + 1)

	for i := 0; i < strLen; i++ {

		_, err := b.WriteRune(charRune)

		if err != nil {
			return "",
				fmt.Errorf("%s\n"+
					"Error returned by b.WriteRune(charRune).\n"+
					"charRune='%v'\n"+
					"Error='%v'\n",
					ePrefix.String(),
					charRune,
					err.Error())
		}
	}

	return b.String(), nil
}

// ptr - Returns a pointer to a new instance of
// strMechQuark.
//
func (sMechQuark strMechQuark) ptr() *strMechQuark {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	return &strMechQuark{
		lock: new(sync.Mutex),
	}
}

// removeStringChar - Removes or deletes a specified character
// from a string and returns a new string.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetStr                  string
//     - The target string containing the character to be removed.
//       If this is a zero length or empty string, an error will
//       be returned.
//
//
//  charToRemove               rune
//     - The character which will be removed from 'targetStr'. If
//       this is an empty character or zero value rune, an error
//       will be returned.
//
//
//  maxNumOfCharDeletions      int
//     - If this parameter is set to minus one (-1), all instances
//       of 'charToRemove' in 'targetStr' will be deleted in the
//       returned string. If this parameter is greater than zero,
//       it will limit the maximum number of character deletions
//       in this operation. Remember that the search for
//       'charToRemove' proceeds from left to right starting at
//       index zero (0).
//
//       If this parameter is set to zero, an error will be
//       returned.
//
//
//  ePrefix             *ErrPrefixDto
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
// -----------------------------------------------------------------
//
// Return Values
//
//  newStr                     string
//     - If this method completes successfully, a new string
//       will be returned containing all the characters in
//       input parameter 'targetStr' except those removed by
//       the deletion operation.
//
//
//  numOfDeletions             int
//     - This parameter will record the number of character
//       deletions performed by this operation.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) removeStringChar(
	targetStr string,
	charToRemove rune,
	maxNumOfCharDeletions int,
	ePrefix *ePref.ErrPrefixDto) (
	newStr string,
	numOfDeletions int,
	err error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"replaceStringChar()")

	if len(targetStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStr' is an "+
			"empty string!\n",
			ePrefix.String())

		return newStr, numOfDeletions, err
	}

	if charToRemove == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charToRemove' is an "+
			"empty character!\n"+
			"charToReplace == 0\n",
			ePrefix.String())

		return newStr, numOfDeletions, err
	}

	if maxNumOfCharDeletions == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxNumOfReplacements' is "+
			"zero!\n",
			ePrefix.String())

		return newStr, numOfDeletions, err
	}

	runeArray := make([]rune, 0, 100)

	for _, r := range targetStr {

		if maxNumOfCharDeletions > -1 &&
			numOfDeletions >= maxNumOfCharDeletions {
			runeArray = append(runeArray, r)
			continue
		}

		if r == charToRemove {
			numOfDeletions++
			continue
		}

		runeArray = append(runeArray, r)
	}

	newStr = string(runeArray)

	return newStr, numOfDeletions, err
}

// removeSubString - Removes or deletes a sub-string
// ('targetSubStr') from a parent string ('strToSearch'). The
// returned string is identical to the original string with the
// sole exception that all specified instances of the target
// sub-string have been deleted and removed.
//
// The number of 'targetSubStr' instances deleted from 'strToSearch'
// is controlled by input parameter, 'maxRemovalCount'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToSearch                string
//     - This is the parent string which will be searched for
//       instances of the target sub-string to be deleted.
//
//
//  targetSubStr               string
//     - This is the sub-string which will be targeted for deletion
//       and removal when located in 'strToSearch'.
//
//
//  maxRemovalCount            int
//     - Specifies the maximum number of 'targetSubStr' instances
//       which will be deleted and removed from 'strToSearch'.
//
//       If 'maxRemovalCount' is set to a value less than one (+1),
//       it specifies that all instances of 'targetSubStr' which
//       exist in 'strToSearch' will be deleted and removed.
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - If this method completes successfully, this string will be
//       populated with an identical copy of 'strToSearch' minus the
//       target sub-strings deleted and removed.
//
//
//  int
//     - If this method completes successfully, this integer value
//       will record the number of target sub-strings removed from
//       'strToSearch'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) removeSubString(
	strToSearch string,
	targetSubStr string,
	maxRemovalCount int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	int,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"removeSubString()")

	var err error

	if len(strToSearch) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: input parameter 'strToSearch' is"+
			" a zero length or empty string!\n",
			ePrefix.String())

		return "", 0, err
	}

	if len(targetSubStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: input parameter 'targetSubStr' is"+
			" a zero length or empty string!\n",
			ePrefix.String())

		return "", 0, err
	}

	if maxRemovalCount < 1 {
		maxRemovalCount = math.MaxInt32
	}

	newRunes := make([]rune, 0, 100)

	targetRunesSubArray := []rune(targetSubStr)
	lenTargetRunes := len(targetRunesSubArray)
	runesToSearch := []rune(strToSearch)
	lenRunesToSearch := len(runesToSearch)

	isCorrectCharCnt := 0
	removalCnt := 0

	for i := 0; i < lenRunesToSearch; i++ {

		if runesToSearch[i] == targetRunesSubArray[0] &&
			i+lenTargetRunes <= lenRunesToSearch &&
			removalCnt < maxRemovalCount {

			isCorrectCharCnt = 1

			for j := 1; j < lenTargetRunes; j++ {

				if runesToSearch[i+j] == targetRunesSubArray[j] {
					isCorrectCharCnt++
				} else {
					break
				}
			}

			if isCorrectCharCnt == lenTargetRunes {
				i += lenTargetRunes - 1
				removalCnt++
			} else {
				newRunes = append(newRunes, runesToSearch[i])
			}

			continue

		} else {
			newRunes = append(newRunes, runesToSearch[i])
		}
	}

	return string(newRunes), removalCnt, err
}

// replaceRunes - Replaces individual characters in a target array
// of runes ([]rune) with those specified in a two-dimensional slice
// of runes, 'replacementRunes[][]'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetRunes         []rune
//     - The rune array which will be examined. If target characters
//       ('runes') eligible for replacement are identified by
//       replacementRunes[i][0], they will be replaced by the
//       character specified in replacementRunes[i][1].
//
//
//  replacementRunes    [][2]rune
//     - A two dimensional slice of type 'rune'. Element [i][0]
//       contains the target character to locate in 'targetRunes'.
//       Element[i][1] contains the replacement character which will
//       replace the target character in 'targetRunes'. If the
//       replacement character element [i][1] is a zero value, the
//       target character will not be replaced. Instead, it will be
//       eliminated or removed from the returned rune array
//       ([]rune).
//
//
//  ePrefix             *ErrPrefixDto
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
//  []rune
//     - The returned rune array containing the characters replaced
//       from the original 'targetRunes' array.
//
//
//  error
//     - If the method completes successfully this value is 'nil'.
//       If an error is encountered this value will contain the
//       error message. Examples of possible errors include a zero
//       length 'targetRunes' array or 'replacementRunes' array.
//
//       In addition, if any of the replacementRunes[][x] 2nd
//       dimension elements have a length less than two, an
//       error will be returned.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) replaceRunes(
	targetRunes []rune,
	replacementRunes [][2]rune,
	ePrefix *ePref.ErrPrefixDto) (
	[]rune,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"replaceRunes()")

	output := make([]rune, 0, 100)

	targetLen := len(targetRunes)

	if targetLen == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetRunes' is a "+
				"zero length array!\n",
				ePrefix.String())
	}

	baseReplaceLen := len(replacementRunes)

	if baseReplaceLen == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'replacementRunes' is a "+
				"zero length array!\n",
				ePrefix.String())
	}

	for i := 0; i < targetLen; i++ {

		foundReplacement := false

		for k := 0; k < baseReplaceLen; k++ {

			if targetRunes[i] == replacementRunes[k][0] {

				if replacementRunes[k][1] != 0 {
					output = append(output, replacementRunes[k][1])
				}

				foundReplacement = true
				break
			}
		}

		if !foundReplacement {
			output = append(output, targetRunes[i])
		}

	}

	return output, nil
}

// replaceStringChar - Replaces a specific character found anywhere
// in a string with another specified substitute character.
//
// The replacement operation proceeds from left to right within the
// 'targetStr' beginning with index zero (0).
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr                  string
//     - The string containing the character to be replaced.
//       If this is an empty string, an error will be returned.
//
//
//  charToReplace              rune
//     - The character within input parameter string 'targetStr'
//       which will be replaced. If this parameter is set to zero
//       signaling an empty character, this method will return an
//       error.
//
//
//  replacementChar            rune
//     - The character which will replace 'charToReplace' in
//       'targetStr'. If this parameter is set to zero signaling an
//       empty character, this method will return an error.
//
//
//  maxNumOfReplacements       int
//     - The maximum number of replacements allowed for this
//       operation. If this parameter is set to minus one (-1), all
//       instances of 'charToReplace' in 'targetStr' will be
//       replaced with 'replacementChar'.
//
//       If this parameter is set to zero ('0'), an error will be
//       returned.
//
//       The replacement operation proceeds from left to right
//       within the 'targetStr' beginning with index zero (0).
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - If this method completes successfully, a new string will be
//       returned with the designated replacement characters.
//
//
//  int
//     - The integer value records the number of characters
//       replaced in 'targetStr'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) replaceStringChar(
	targetStr string,
	charToReplace rune,
	replacementChar rune,
	maxNumOfReplacements int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	int,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"replaceStringChar()")

	var err error

	if len(targetStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStr' is an empty "+
			"string!\n",
			ePrefix.String())

		return "", 0, err
	}

	if charToReplace == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charToReplace' is an "+
			"empty character!\n"+
			"charToReplace == 0\n",
			ePrefix.String())

		return "", 0, err
	}

	if replacementChar == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replacementChar' is an "+
			"empty character!\n"+
			"replacementChar == 0\n",
			ePrefix.String())

		return "", 0, err
	}

	if maxNumOfReplacements == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxNumOfReplacements' is "+
			"zero!\n",
			ePrefix.String())

		return "", 0, err
	}

	runeArray := []rune(targetStr)

	numOfReplacements := 0

	for i, r := range runeArray {

		if maxNumOfReplacements > -1 &&
			numOfReplacements >= maxNumOfReplacements {
			break
		}

		if r == charToReplace {
			runeArray[i] = replacementChar
			numOfReplacements++
		}

	}

	return string(runeArray), numOfReplacements, err
}

// replaceSubString - Replaces instances of a sub-string
// ('targetSubStr') found in a parent string ('strToSearch') with
// a replacement string ('replacementStr').
//
// The number of instances located and replaced in 'strToSearch' is
// controlled by input parameter 'maxReplacementCount'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToSearch                string
//     - This is the parent string which will be searched for
//       instances of the target sub-string to be replaced.
//
//
//  targetSubStr               string
//     - This is the sub-string which will be targeted for
//       replacement when located in 'strToSearch'.
//
//
//  replacementStr             string
//     - This is the replacement string which will replace instances
//       of 'targetSubStr' located in 'strToSearch'.
//
//
//  maxReplacementCount        int
//     - Specifies the maximum number of 'targetSubStr' instances
//       which will be replaced with 'replacementStr' in the parent
//       string, 'strToSearch'.
//
//       If 'maxReplacementCount' is set to a value less than one
//       (+1), it specifies that all instances of 'targetSubStr'
//       which exist in 'strToSearch' will be replaced with the
//       replacement string, 'replacementStr' .
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - If this method completes successfully, this string will be
//       populated with an identical copy of 'strToSearch' where
//       instances of the target sub-string have been replaced with
//       the replacement string, 'replacementStr'.
//
//
//  int
//     - If this method completes successfully, this integer value
//       will record the number of target sub-strings replaced
//       within 'strToSearch'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechQuark *strMechQuark) replaceSubString(
	strToSearch string,
	targetSubStr string,
	replacementStr string,
	maxReplacementCount int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	int,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"replaceSubString()")

	var err error

	err = nil

	if len(strToSearch) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: input parameter 'strToSearch' is"+
			" a zero length or empty string!\n",
			ePrefix.String())

		return "", 0, err
	}

	if len(targetSubStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: input parameter 'targetSubStr' is"+
			" a zero length or empty string!\n",
			ePrefix.String())

		return "", 0, err
	}

	if len(replacementStr) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: input parameter 'replacementStr' is"+
			" a zero length or empty string!\n",
			ePrefix.String())

		return "", 0, err
	}

	if maxReplacementCount < 1 {
		maxReplacementCount = math.MaxInt32
	}

	newRunes := make([]rune, 0, 100)

	targetRunes := []rune(targetSubStr)
	lenTargetRunes := len(targetRunes)
	replacementRunes := []rune(replacementStr)
	runesToSearch := []rune(strToSearch)
	lenRunesToSearch := len(runesToSearch)

	isCorrectCharCnt := 0
	replacementCnt := 0

	for i := 0; i < lenRunesToSearch; i++ {

		if runesToSearch[i] == targetRunes[0] &&
			i+lenTargetRunes <= lenRunesToSearch &&
			replacementCnt < maxReplacementCount {

			isCorrectCharCnt = 1

			for j := 1; j < lenTargetRunes; j++ {

				if runesToSearch[i+j] == targetRunes[j] {
					isCorrectCharCnt++
				} else {
					break
				}
			}

			if isCorrectCharCnt == lenTargetRunes {
				i += lenTargetRunes - 1
				replacementCnt++
				// copy(newRunes[i:], replacementRunes)
				newRunes = append(newRunes, replacementRunes...)

			} else {
				newRunes = append(newRunes, runesToSearch[i])
			}

			continue

		} else {
			newRunes = append(newRunes, runesToSearch[i])
		}
	}

	return string(newRunes), replacementCnt, err
}

// stripBadChars - Removes/deletes specified sub-strings from a
// parent or host string. The targeted sub-strings are deleted
// wherever found in the parent or host string.
//
// The sub-strings to be removed are identified in a string array
// passed as input parameter, 'badChars'.
//
// All instances of 'badChars' sub-strings are deleted from the
// target string which is passed as input parameter, 'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be searched for the sub-strings
//       identified in the 'badChars' array for deletion.
//
//
//  badChars            []string
//     - A one dimensional array of strings which contains the
//       sub-strings to be deleted from input parameter,
//       'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  cleanStr            string
//     - This returned string is a copy of 'targetStr' minus the
//       sub-strings identified in the 'badChars' array which are
//       deleted.
//
//  strLen              int
//     - This integer value contains the length of the newly
//       generated, 'cleanStr', described above.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//   badChars := []string{"@@"}
//
//   expectedStrLen := len(expectedStr)
//                  12    123456789    12  12
//   testString := "@@Some@@@@@@@@@Stri@@ng@@"
//
//   actualString, actualStrLen :=
//         StrMech{}.Ptr().StripBadChars(
//                            testString,
//                            badChars)
//
//  -----------------------------------------------
//   actualString is now equal to "Some@String"
//
//
func (sMechQuark *strMechQuark) stripBadChars(
	targetStr string,
	badChars []string) (
	cleanStr string,
	strLen int) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	cleanStr = targetStr
	strLen = len(cleanStr)

	if strLen == 0 {
		return cleanStr, strLen
	}

	lenBadChars := len(badChars)

	if lenBadChars == 0 {
		return cleanStr, strLen
	}

	sort.Sort(SortStrLengthHighestToLowest(badChars))

	cycleWhereStringRemoved := 0
	k := -1

	for {

		k++

		for i := 0; i < lenBadChars; i++ {

			for {

				badCharIdx := strings.Index(cleanStr, badChars[i])

				if badCharIdx == -1 {
					break
				}

				lastCleanStrIdx := strLen - 1
				lChar := len(badChars[i])
				nextIdx := badCharIdx + lChar

				if nextIdx > lastCleanStrIdx {
					cleanStr = cleanStr[0:badCharIdx]

				} else {

					cleanStr = cleanStr[0:badCharIdx] + cleanStr[nextIdx:]
				}

				cycleWhereStringRemoved = k
			}

			strLen = len(cleanStr)

			if strLen == 0 {
				goto Done
			}
		}

		if k-cycleWhereStringRemoved > 3 || k > 1000000 {
			goto Done
		}
	}

Done:

	return cleanStr, strLen
}

// stripLeadingChars - Strips or deletes sub-strings from the front
// of a parent or host string labeled here as 'targetStr'. The
// sub-strings to be deleted are specified by input parameter and
// string array, 'badChars'.
//
// The method then returns a string which does not contain leading
// sub-strings identified in the 'badChars' array. In addition, the
// length of the final 'clean string' is also returned.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be searched for leading sub-strings
//       identified in the 'badChars' array for deletion.
//
//
//  badChars            []string
//     - A one dimensional array of strings which contains the
//       sub-strings to be deleted from input parameter,
//       'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  cleanStr            string
//     - This returned string is a copy of 'targetStr' minus the
//       sub-strings identified in the 'badChars' array which are
//       deleted from the front of 'targetStr'.
//
//  strLen              int
//     - This integer value contains the length of the newly
//       generated, 'cleanStr', described above.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  badChars := []string{
//                 " ", // A single white-space character
//                 "/",
//                 "//",
//                 "\\\\",
//                 "\\",
//                 ".\\",
//                 "../",
//                 ".",
//                 "..\\",
//                 "\\\\\\",
//                 "..",
//                 "./",
//                 "//",
//                 "///",
//                 "////",
//                 "..."}
//
//  testString :=
//   "..........      ./../.\\.\\..\\////   SomeString"
//
//  actualString, actualStrLen :=
//      StrMech{}.Ptr().StripLeadingChars(
//                       testString,
//                       badChars)
//
//  ----------------------------------------------------
//  actualString is now equal to "SomeString"
//  actualStrLen is now equal to 10
//
//
func (sMechQuark *strMechQuark) stripLeadingChars(
	targetStr string,
	badChars []string) (
	cleanStr string,
	strLen int) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	cleanStr = targetStr
	strLen = len(cleanStr)

	lenBadChars := len(badChars)

	if lenBadChars == 0 {
		return cleanStr, strLen
	}

	if strLen == 0 {
		return cleanStr, strLen
	}

	sort.Sort(SortStrLengthHighestToLowest(badChars))

	cycleWhereStringRemoved := 0
	k := -1

	for {

		k++

		for i := 0; i < lenBadChars; i++ {

			for {

				if !strings.HasPrefix(cleanStr, badChars[i]) {
					break
				}

				cleanStr = cleanStr[len(badChars[i]):]

				cycleWhereStringRemoved = k
			}

			strLen = len(cleanStr)

			if strLen == 0 {
				goto Done
			}
		}

		if k-cycleWhereStringRemoved > 3 || k > 1000000 {
			goto Done
		}
	}

Done:

	return cleanStr, strLen
}

// StripTrailingChars - Strips or deletes sub-strings from the
// end of a parent or host string. The sub-strings to be deleted
// are identified in a string array input parameter labeled,
// 'badChars'.  The parent string to be searched is passed as input
// parameter, 'targetStr'. The targeted sub-strings are only deleted
// if they exist at the end of 'targetStr'.
//
// Upon completion, this method returns the cleaned string and the
// length of the cleaned string to the caller.  The cleaned string
// is equivalent to input parameter, 'targetStr', minus the trailing
// sub-strings identified by string array 'badChars'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The parent or host string which will be searched for
//       instances of trailing sub-strings identified in the
//       'badChars' string array for deletion.
//
//
//  badChars            []string
//     - A one dimensional array of strings which contains the
//       sub-strings to be deleted from the end of 'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  cleanStr            string
//     - This returned string is a copy of 'targetStr' minus the
//       trailing sub-strings identified for deletion in the
//       'badChars' array.
//
//  strLen              int
//     - This integer value contains the length of the newly
//       generated, 'cleanStr', described above.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  badChars := []string{
//                 " ", // Single white space character
//                 "/",
//                 "//",
//                 "\\\\",
//                 "\\",
//                 ".\\",
//                 "../",
//                 ".",
//                 "..\\",
//                 "\\\\\\",
//                 "..",
//                 "./",
//                 "//",
//                 "///",
//                 "////",
//                 "..."}
//
//  testString :=
//   "SomeString..........      ./../.\\.\\..\\////   "
//
//  actualString, actualStrLen :=
//    StrMech{}.Ptr().StripTrailingChars(
//                      testString,
//                      badChars)
//
//  -------------------------------------------------------------
//
//                                1234567890
//  actualString is now equal to "SomeString"
//  actualStrLen is now equal to 10
//
func (sMechQuark *strMechQuark) stripTrailingChars(
	targetStr string,
	badChars []string) (
	cleanStr string,
	strLen int) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	cleanStr = targetStr
	strLen = len(cleanStr)

	lenBadChars := len(badChars)

	if lenBadChars == 0 {
		return cleanStr, strLen
	}

	if strLen == 0 {
		return cleanStr, strLen
	}

	k := -1
	cycleWhereStringRemoved := 0

	for {

		k++

		for i := 0; i < lenBadChars; i++ {

			for {

				if !strings.HasSuffix(cleanStr, badChars[i]) {
					break
				}

				strLen = len(cleanStr) - len(badChars[i])

				cleanStr = cleanStr[0:strLen]

				cycleWhereStringRemoved = k

			}

			if len(cleanStr) == 0 {
				goto Done
			}
		}

		strLen = len(cleanStr)

		if strLen == 0 {
			goto Done
		}

		if k-cycleWhereStringRemoved > 3 || k > 1000000 {
			goto Done
		}

	}

Done:

	return cleanStr, strLen
}

// SwapRune - Swaps, or replaces, instances of the 'oldRune'
// character with the 'newRune' character in the input parameter
// target string ('targetStr'). The number of character
// substitutions performed is controlled by input parameter
// 'maxNumOfSwaps'.
//
// If input parameter 'targetStr' is an empty string, no error will
// be generated and the empty 'targetStr' will be returned.
//
// For more information on 'runes', reference:
//   https://www.geeksforgeeks.org/rune-in-golang/
//   https://golangbyexample.com/understanding-rune-in-golang/
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - This string will be searched for instances of 'oldRune'.
//       When 'oldRune' is located in this string, it will be
//       replaced with 'newRune'.
//
//
//  oldRune             rune
//     - This rune or text character will be used search the parent
//       string, 'targetStr'. When instances of 'oldRune' are
//       located in 'targetStr', they will be replaced with
//       'newRune'.
//
//
//  newRune             rune
//     - This rune or text character will be used to replaced
//       instances of 'oldRune' located in a search of 'targetStr'.
//
//
//  maxNumOfSwaps       int
//     - This integer value constitutes the upper limit for the
//       number of character substitutions performed by this method.
//       If this value is less than one (+1), no limit will be
//       imposed and all instances of 'oldRune' found in 'targetStr'
//       will be replaced with 'newRune'.
//
//
//  ePrefix             *ErrPrefixDto
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
//  string
//     - This string will be populated with the contents of
//       'targetStr' after instances of 'oldRune' have been
//        replaced by 'newRune'. Effectively, this string is the
//        result of the character substitution operation.
//
//  int
//     - The value of this returned integer records the number of
//       character substitutions performed on 'targetStr'.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     tStr := "  Hello   World  "
//     ePrefix := "theCallingMethodName "
//
//     su := StrMech{}
//        resultStr,
//        numOfReplacements,
//        err := su.SwapRune(
//                 tStr,
//                 ' ',
//                 '!',
//                 -1,
//                 ePrefix)
//
//
//    resultSt is now equal to "!!Hello!!!World!!"
//    numOfReplacements is equal to '7'
//
func (sMechQuark *strMechQuark) swapRune(
	targetStr string,
	oldRune rune,
	newRune rune,
	maxNumOfSwaps int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	int,
	error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"swapRune()")

	if targetStr == "" {
		return targetStr, 0, nil
	}

	if maxNumOfSwaps < 1 {
		maxNumOfSwaps = math.MaxInt32
	}

	rStr := []rune(targetStr)

	lrStr := len(rStr)
	numOfSwapsCnt := 0

	for i := 0; i < lrStr; i++ {

		if numOfSwapsCnt >= maxNumOfSwaps {
			break
		}

		if rStr[i] == oldRune {
			numOfSwapsCnt++
			rStr[i] = newRune
		}

	}

	return string(rStr), numOfSwapsCnt, nil
}

// TrimMultipleChars - Performs the following operations on strings:
//
//  1. Trims Right and Left ends of 'targetStr' for all instances
//     of 'trimChar'
//
//  2. Within the interior of a string, multiple instances of
//     'trimChar' are reduced to a single instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The parent or host string which will be searched for
//       instances of the character 'trimChar'.
//
//
//  trimChar            rune
//     - 'targetStr' will be searched for instances of this
//       character. If this character is a leading or trailing
//       character in 'targetStr', those instances will be deleted.
//       If this character is found in the interior of 'targetStr',
//       it will be reduced to one single instance of the
//       character.
//
//
//  ePrefix             *ErrPrefixDto
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
//  rStr                string
//     - This is the 'result' string. It is comprised of all the
//       characters in the original 'targetStr' minus those
//       'trimChar' characters deleted in the 'trim' operation.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "theCallingFunctionName() "
//  targetStr = "       Hello          World        "
//  trimChar  = ' ' (One Space Character)
//  sops := StrMech{}
//
//  rStr,
//  err := sops.TrimMultipleChars(
//           targetStr,
//           trimChar,
//           ePrefix)
//
//  returned string (rStr) now equal to "Hello World"
//
//
func (sMechQuark *strMechQuark) trimMultipleChars(
	targetStr string,
	trimChar rune,
	ePrefix *ePref.ErrPrefixDto) (
	rStr string,
	err error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"trimMultipleChars()")

	rStr = ""
	err = nil

	if targetStr == "" {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStr' is an "+
			"EMPTY STRING!\n",
			ePrefix.String())

		return rStr, err
	}

	if trimChar == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trimChar' is ZERO!\n",
			ePrefix.String())

		return rStr, err
	}

	fStr := []rune(targetStr)
	lenTargetStr := len(fStr)
	outputStr := make([]rune, lenTargetStr)
	lenTargetStr--
	idx := lenTargetStr
	foundFirstChar := false

	for i := lenTargetStr; i >= 0; i-- {

		if !foundFirstChar && fStr[i] == trimChar {
			continue
		}

		if i > 0 && fStr[i] == trimChar && fStr[i-1] == trimChar {
			continue
		}

		if i == 0 && fStr[i] == trimChar {
			continue
		}

		foundFirstChar = true
		outputStr[idx] = fStr[i]
		idx--
	}

	if idx != lenTargetStr {
		idx++
	}

	if outputStr[idx] == trimChar {
		idx++
	}

	rStr = string(outputStr[idx:])
	err = nil

	return rStr, err
}

// trimStringEnds - Removes all instances of input parameter
// 'trimChar' from the beginning and end of input parameter string
// 'targetStr'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The parent or host string which will be searched for
//       instances of the character 'trimChar'.
//
//
//  trimChar            rune
//     - 'targetStr' will be searched for instances of this
//       character. If the character is found to be either a
//       trailing character or a leading character, it will be
//       deleted. If this character exists in the interior of
//       'targetStr' it will be ignored and NOT deleted.
//
//
//  ePrefix             *ErrPrefixDto
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
//  rStr                string
//     - This is the 'result' string. It is comprised of all the
//       characters in the original 'targetStr' minus those
//       'trimChar' characters deleted in the 'trim' operation.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//     ePrefix := "TestStrOps_TrimStringEnds_01() "
//     tStr := "XXXHello WorlXdXXX"
//     trimChar := 'X'
//
//     sops := StrMech{}
//
//     result,
//     err := sops.TrimStringEnds(
//              tStr,
//              trimChar,
//              ePrefix)
//
//  result is now equal to "Hello WorlXd"
//
//
func (sMechQuark *strMechQuark) trimStringEnds(
	targetStr string,
	trimChar rune,
	ePrefix *ePref.ErrPrefixDto) (
	rStr string,
	err error) {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechQuark." +
			"trimStringEnds()")

	rStr = ""
	err = nil

	targetStrLen := len(targetStr)

	if targetStrLen == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetStr' is an "+
			"EMPTY STRING!\n",
			ePrefix.String())

		return rStr, err
	}

	if trimChar == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trimChar' is ZERO!\n",
			ePrefix.String())

		return rStr, err
	}

	foundGoodChar := false
	firstIdx := 0

	for !foundGoodChar {

		if rune(targetStr[firstIdx]) == trimChar {
			firstIdx++
		} else {
			foundGoodChar = true
		}

		if firstIdx >= targetStrLen {
			rStr = ""
			return rStr, err
		}
	}

	if firstIdx >= targetStrLen {

		rStr = targetStr

		return rStr, err
	}

	foundGoodChar = false
	lastIdx := targetStrLen - 1

	for !foundGoodChar {

		if rune(targetStr[lastIdx]) == trimChar {
			lastIdx--
		} else {
			foundGoodChar = true
		}
	}

	if lastIdx < 0 {
		rStr = targetStr[firstIdx:]
		return rStr, err
	}

	lastIdx++
	rStr = targetStr[firstIdx:lastIdx]

	return rStr, err
}

// upperCaseFirstLetter - Finds the first alphabetic character in a string
// (a-z A-Z) and converts it to upper case.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  str                 string
//     - The first character in this string will be converted to
//       an Upper Case Letter.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string              string
//     - This is copy of input parameter 'str' with the first
//       character converted to upper case (a.k.a. a capital
//       letter).
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//   str := "how now brown cow."
//
//   sMech := StrMech{}
//
//   actualStr := sMech.upperCaseFirstLetter(str)
//
//  'actualStr' is now equal to "How now brown cow."
//
//
func (sMechQuark *strMechQuark) upperCaseFirstLetter(
	str string) string {

	if sMechQuark.lock == nil {
		sMechQuark.lock = new(sync.Mutex)
	}

	sMechQuark.lock.Lock()

	defer sMechQuark.lock.Unlock()

	if len(str) == 0 {
		return str
	}

	runesStr := []rune(str)

	for i := 0; i < len(runesStr); i++ {

		// Skip leading spaces
		if runesStr[i] == ' ' {
			continue
		}

		// Find the first alphabetic character and
		// convert to upper case.
		if runesStr[i] >= 'a' && runesStr[i] <= 'z' {

			runesStr[i] -= 32
			break

		} else if runesStr[i] >= 'A' && runesStr[i] <= 'Z' {
			break
		}

	}

	return string(runesStr)
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"sync"
)

type strMechElectron struct {
	lock *sync.Mutex
}

// findFirstNonSpaceChar - Returns the string index of the first non-space character in
// a string segment. The string to be searched is input parameter 'targetStr'. The string
// segment which will be searched from left to right in 'targetStr' is defined by the
// starting index ('startIndex') and the ending index ('endIndex').
//
// Searching from left to right, this method identifies the first non-space character
// (any character that is NOT a space ' ') in the target string segment and returns
// the index associated with that non-space character.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string to be searched for the first non-space character.
//
//
//  startIdx            int
//     - Since the search is forwards from left to right, this is
//       the starting index for the search.
//
//
//  endIdx              int
//     - Since the search is forwards from left to right, this is
//       the ending index for the search.
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
//  int
//     - This method returns the index of the first non-space
//       character in the target string segment using a left
//       to right search. If the target string is an empty string
//       or consists of entirely of space characters, this method
//       returns a value of minus one (-1).
//
//
//  error
//     - If the method completes successfully this value is
//       'nil'. If an error is encountered this value will
//       contain the error message.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechElectron *strMechElectron) findFirstNonSpaceChar(
	targetStr string,
	startIndex,
	endIndex int,
	ePrefix *ePref.ErrPrefixDto) (
	int,
	error) {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechElectron." +
			"findFirstNonSpaceChar()")

	sMechQuark := strMechQuark{}

	if sMechQuark.isEmptyOrWhiteSpace(targetStr) {
		return -1, nil
	}

	targetStrLen := len(targetStr)

	if startIndex < 0 {
		return -1, fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter. 'startIndex' is LESS THAN ZERO! "+
			"startIndex='%v' ",
			ePrefix.String(),
			startIndex)
	}

	if endIndex < 0 {
		return -1, fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter. 'endIndex' is LESS THAN ZERO! "+
			"startIndex='%v' ",
			ePrefix.String(),
			startIndex)
	}

	if endIndex >= targetStrLen {
		return -1, fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter. 'endIndex' is greater than "+
			"target string length. INDEX OUT OF RANGE! endIndex='%v' target string length='%v' ",
			ePrefix.String(),
			endIndex,
			targetStrLen)
	}

	if startIndex > endIndex {
		return -1, fmt.Errorf("%v\n"+
			"ERROR: Invalid input parameter. 'startIndex' is GREATER THAN 'endIndex' "+
			"startIndex='%v' endIndex='%v' ",
			ePrefix.String(),
			startIndex,
			endIndex)
	}

	idx := startIndex

	for idx <= endIndex {

		if targetStr[idx] != ' ' {
			return idx, nil
		}

		idx++
	}

	return -1, nil

}

// GetValidString - Validates the individual characters in input parameter string,
// 'targetStr'. To identify valid characters, the characters in 'targetStr' are
// compared against input parameter 'validRunes', an array of type rune. If a character
// exists in both 'targetStr' and 'validRunes' it is deemed valid and returned in
// an output string.
//
// ------------------------------------------------------------------------
//
// Input Parameter
//
//  targetStr           string
//     - The string which will be screened for valid characters.
//
//
//  validRunes          []rune
//     - An array of type rune containing valid characters. Characters
//       which exist in both 'targetStr' and 'validRunes' will be
//       returned as a new string. Invalid characters are discarded.
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
//  string
//     - This string will be returned containing valid characters extracted
//       from 'targetStr'. A character is considered valid if it exists in
//       both 'targetStr' and 'validRunes'. Invalid characters are discarded.
//       This means that if no valid characters are identified, a zero length
//       string will be returned.
//
//  error
//     - If the method completes successfully this value is 'nil'. If an error is
//       encountered this value will contain the error message. Examples of possible
//       errors include a zero length 'targetStr' (string) or a zero length
//       'validRunes' array.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechElectron *strMechElectron) getValidString(
	targetStr string,
	validRunes []rune,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechElectron." +
			"getValidString()")

	if len(targetStr) == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetStr' is a "+
				"ZERO LENGTH STRING!\n",
				ePrefix.String())
	}

	if len(validRunes) == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'validRunes' is a "+
				"ZERO LENGTH ARRAY!\n",
				ePrefix.String())
	}

	actualValidRunes, err :=
		strMechQuark{}.ptr().
			getValidRunes(
				[]rune(targetStr),
				validRunes,
				ePrefix)

	if err != nil {
		return "", err
	}

	return string(actualValidRunes), err
}

// ptr - Returns a pointer to a new instance of
// strMechElectron.
//
func (sMechElectron strMechElectron) ptr() *strMechElectron {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	return &strMechElectron{
		lock: new(sync.Mutex),
	}
}

// readBytes - Implements io.Reader interface for type StrMech.
// 'readBytes' reads up to len(p) bytes into 'p'. This
// method supports buffered 'read' operations.
//
// The internal member string variable, 'StrMech.stringData'
// is written into 'p'. When the end of 'StrMech.stringData'
// is written to 'p', the method returns error = 'io.EOF'.
//
// 'StrMech.stringData' can be accessed through Getter an Setter methods,
// StrMech.GetStringData() and StrMech.SetStringData()
//
func (sMechElectron *strMechElectron) readBytes(
	strOpsInstance *StrMech,
	p []byte,
	ePrefix *ePref.ErrPrefixDto) (
	n int,
	err error) {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechElectron." +
			"readBytes()")

	n = 0
	err = nil

	if strOpsInstance == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'strOpsInstance' is a 'nil' pointer!\n",
			ePrefix)

		return n, err
	}

	n = len(p)
	err = io.EOF

	if n == 0 {
		strOpsInstance.cntBytesRead = 0
		err = fmt.Errorf("%v\n"+
			"Error: Input byte array 'p' is zero length!\n",
			ePrefix)

		return 0, err
	}

	strData := strOpsInstance.stringData

	w := []byte(strData)

	lenW := uint64(len(w))

	cntBytesRead := strOpsInstance.cntBytesRead

	if lenW == 0 ||
		cntBytesRead >= lenW {
		strOpsInstance.cntBytesRead = 0
		n = 0
		return n, err
	}

	startReadIdx := cntBytesRead

	remainingBytesToRead := lenW - cntBytesRead

	if uint64(n) < remainingBytesToRead {
		remainingBytesToRead = startReadIdx + uint64(n)
		err = nil
	} else {
		remainingBytesToRead += startReadIdx
		err = io.EOF
	}

	n = 0

	for i := startReadIdx; i < remainingBytesToRead; i++ {
		p[n] = w[i]
		n++
	}

	cntBytesRead += uint64(n)

	if cntBytesRead >= lenW {
		strOpsInstance.cntBytesRead = 0
	} else {
		strOpsInstance.cntBytesRead = cntBytesRead
	}

	return n, err
}

// readStringFromBytes - Receives a byte array and retrieves a string. The beginning of
// the string is designated by input parameter 'startIdx'. The end of the string is determined
// when a carriage return ('\r'), vertical tab ('\v') or a new line character ('\n') is encountered.
//
// The parsed string is returned to the caller along with 'nextStartIdx', which is the byte
// array index of the beginning of the next string.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  bytes          []byte - An array of bytes from which a string will be extracted
//                          and returned.
//
//  startIdx          int - The starting index in input parameter 'bytes' where the string
//                          extraction will begin. The string extraction will cease when
//                          a carriage return ('\r'), a vertical tab ('\v') or a new line
//                          character ('\n') is encountered.
//
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  extractedStr   string - The string extracted from input parameter 'bytes' beginning
//                          at the index in 'bytes' indicated by input parameter 'startIdx'.
//
//  nextStartIdx      int - The index of the beginning of the next string in the byte array
//                          'bytes' after 'extractedString'. If no more strings exist in the
//                          the byte array, 'nextStartIdx' will be set to -1.
//
func (sMechElectron *strMechElectron) readStringFromBytes(
	bytes []byte,
	startIdx int) (
	extractedStr string,
	nextStartIdx int) {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	extractedStr = ""
	nextStartIdx = -1

	bLen := len(bytes)

	if bLen == 0 {
		return extractedStr, nextStartIdx
	}

	if startIdx >= bLen || startIdx < 0 {
		return extractedStr, nextStartIdx
	}

	nextStartIdx = -1

	runeAry := make([]rune, 0, bLen+5)

	for i := startIdx; i < bLen; i++ {

		if bytes[i] == '\r' ||
			bytes[i] == '\n' ||
			bytes[i] == '\v' {

			if i+1 < bLen {

				j := 0

				for j = i + 1; j < bLen; j++ {
					if bytes[j] == '\r' ||
						bytes[j] == '\v' ||
						bytes[j] == '\n' {
						continue
					} else {
						break
					}
				}

				if j >= bLen {
					nextStartIdx = -1
				} else {
					nextStartIdx = j
				}

				break

			} else {
				nextStartIdx = -1
			}

			break
		}

		runeAry = append(runeAry, rune(bytes[i]))
	}

	extractedStr = string(runeAry)

	return extractedStr, nextStartIdx
}

// ReplaceBytes - Replaces characters in a target array of bytes ([]bytes) with those specified in
// a two dimensional slice of bytes.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//
//  targetBytes         []byte
//     - The byte array which will be examined. If characters ('bytes') eligible
//       for replacement are identified by replacementBytes[i][0] they will be
//       replaced by the character specified in replacementBytes[i][1].
//
//  replacementBytes    [][]byte
//     - A two dimensional slice of type byte. Element [i][0] contains the
//       target character to locate in 'targetBytes'. Element[i][1] contains
//       the replacement character which will replace the target character
//       in 'targetBytes'. If the replacement character element [i][1] is
//       a zero value, the target character will not be replaced. Instead,
//       it will be eliminated or removed from the returned byte array ([]byte).
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
//  []byte
//     - The returned byte array containing the characters and replaced characters
//       from the original 'targetBytes' array.
//
//  error
//     - If the method completes successfully this value is 'nil'. If an error is
//       encountered this value will contain the error message. Examples of possible
//       errors include a zero length targetBytes[] array or replacementBytes[][] array.
//       In addition, if any of the replacementBytes[][x] 2nd dimension elements have
//       a length less than two, an error will be returned.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechElectron *strMechElectron) replaceBytes(
	targetBytes []byte,
	replacementBytes [][]byte,
	ePrefix *ePref.ErrPrefixDto) (
	[]byte,
	error) {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechElectron." +
			"replaceBytes()")

	output := make([]byte, 0, 100)

	targetLen := len(targetBytes)

	if targetLen == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetBytes' is a "+
				"zero length array!\n",
				ePrefix.String())
	}

	baseReplaceLen := len(replacementBytes)

	if baseReplaceLen == 0 {
		return output,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'replacementBytes' is "+
				"a zero length array!\n",
				ePrefix.String())
	}

	for h := 0; h < baseReplaceLen; h++ {

		if len(replacementBytes[h]) < 2 {
			return output,
				fmt.Errorf("%s\n"+
					"Error: Invalid Replacement Array Element.\n"+
					"replacementBytes[%v] has "+
					"a length less than two.\n",
					ePrefix.String(),
					h)
		}

	}

	for i := 0; i < targetLen; i++ {

		foundReplacement := false

		for k := 0; k < baseReplaceLen; k++ {

			if targetBytes[i] == replacementBytes[k][0] {

				if replacementBytes[k][1] != 0 {
					output = append(output, replacementBytes[k][1])
				}

				foundReplacement = true
				break
			}
		}

		if !foundReplacement {
			output = append(output, targetBytes[i])
		}

	}

	return output, nil
}

// ReplaceMultipleStrs - Replaces all instances of string replaceArray[i][0] with
// replacement string from replaceArray[i][1] in 'targetStr'.
//
// Note: The original 'targetStr' is NOT altered.
//
// Input parameter 'replaceArray' should be passed as a two-dimensional slice.
// If the length of the 'replaceArray' second dimension is less than '2', an
// error will be returned.
//
func (sMechElectron *strMechElectron) replaceMultipleStrs(
	targetStr string,
	replaceArray [][]string,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechElectron." +
			"replaceMultipleStrs()")

	if targetStr == "" {
		return targetStr,
			fmt.Errorf("%v\n"+
				"Input parameter 'targetStr' is an "+
				"EMPTY STRING.\n",
				ePrefix.String())
	}

	if len(replaceArray) == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Length of first dimension [X][] in two "+
				"dimensional array\n"+
				"'replaceArray' is ZERO!\n",
				ePrefix.String())
	}

	newString := targetStr

	for aIdx, aVal := range replaceArray {

		if len(aVal) < 2 {
			return "",
				fmt.Errorf("%s\n"+
					"Length of second dimension [][X] in two dimensional array\n"+
					"'replaceArray' is Less Than 2!\n"+
					"replaceArray[%v][]\n",
					ePrefix.String(),
					aIdx)
		}

		newString = strings.Replace(
			newString,
			replaceArray[aIdx][0],
			replaceArray[aIdx][1],
			-1)

	}

	return newString, nil
}

// ReplaceStringChars - Replaces string characters in a target string ('targetStr') with those
// specified in a two dimensional slice of runes, 'replacementRunes[][]'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetStr           string
//     - The string which will be examined. If target string characters
//       eligible for replacement are identified by replacementRunes[i][0],
//       they will be replaced by the character specified in
//       replacementRunes[i][1].
//
//  replacementRunes    [][2]rune
//     - A two dimensional slice of type 'rune'. Element [i][0] contains
//       the target character to locate in 'targetStr'. Element[i][1]
//       contains the replacement character which will replace the target
//       character in 'targetStr'. If the replacement character
//       element [i][1] is a zero value, the target character will not
//       be replaced. Instead, it will be eliminated or removed from the
//       returned string.
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
//  string
//     - The returned string containing the characters and replaced characters
//       from the original target string, ('targetStr').
//
//  error
//     - If the method completes successfully this value is 'nil'. If an error is
//       encountered this value will contain the error message. Examples of possible
//       errors include a zero length 'targetStr' or 'replacementRunes[][]' array.
//       In addition, if any of the replacementRunes[][x] 2nd dimension elements have
//       a length less than two, an error will be returned.
//
//       If an error message is returned, the text value for input
//       parameter 'ePrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (sMechElectron *strMechElectron) replaceStringChars(
	targetStr string,
	replacementRunes [][2]rune,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechElectron." +
			"replaceStringChars()")

	if len(targetStr) == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'targetStr' "+
				"is an EMPTY STRING!\n",
				ePrefix.String())
	}

	if len(replacementRunes) == 0 {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'replacementRunes' is an EMPTY STRING!\n",
				ePrefix)
	}

	outputStr,
		err :=
		strMechQuark{}.ptr().
			replaceRunes(
				[]rune(targetStr),
				replacementRunes,
				ePrefix)

	if err != nil {

		return "",
			fmt.Errorf("%s\n"+
				"Error returned by sOpsQuark.replaceRunes([]rune("+
				"targetStr), replacementRunes).\n"+
				"Error='%v' ",
				ePrefix.String(),
				err.Error())
	}

	return string(outputStr), nil

}

// write - Implements the io.Writer interface.
// Write writes len(p) bytes from p to the underlying
// data stream. In this case the underlying data stream
// is private member variable string, 'StrMech.stringData'.
//
// Receives a byte array 'p' and writes the contents to
// a string, private structure data element 'StrMech.stringData'.
//
// 'StrMech.stringData' can be accessed through 'Getter' and
// 'Setter' methods, 'GetStringData()' and 'SetStringData()'.
//
func (sMechElectron *strMechElectron) write(
	strOpsInstance *StrMech,
	p []byte,
	ePrefix *ePref.ErrPrefixDto) (
	n int,
	err error) {

	if sMechElectron.lock == nil {
		sMechElectron.lock = new(sync.Mutex)
	}

	sMechElectron.lock.Lock()

	defer sMechElectron.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechElectron." +
			"write()")

	n = 0
	err = nil

	if strOpsInstance.cntBytesWritten == 0 {
		strOpsInstance.stringData = ""
	}

	n = len(p)

	if n == 0 {

		strOpsInstance.cntBytesWritten = 0

		err = fmt.Errorf("%v\n"+
			"Error: Input byte array 'p' is ZERO LENGTH!\n",
			ePrefix.String())

		return n, err
	}

	w := strings.Builder{}
	w.Grow(n + 5)
	cnt := 0

	endOfString := false

	for i := 0; i < n; i++ {

		if p[i] == 0 {
			endOfString = true
			break
		}

		w.WriteByte(p[i])
		cnt++
	}

	n = cnt

	strOpsInstance.stringData += w.String()

	if endOfString {
		strOpsInstance.cntBytesWritten = 0
	} else {
		strOpsInstance.cntBytesWritten += uint64(n)
	}

	return n, err
}

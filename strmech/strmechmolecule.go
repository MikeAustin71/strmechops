package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type strMechMolecule struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// strMechMolecule.
//
func (sMechMolecule strMechMolecule) ptr() *strMechMolecule {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	return &strMechMolecule{
		lock: new(sync.Mutex),
	}
}

// strCenterInStr - returns a string which includes a left pad blank string plus
// the original string ('strToCenter'), plus a right pad blank string.
//
// The returned string will effectively center the original string ('strToCenter')
// in a field of specified length ('fieldLen').
//
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToCenter         string
//     - This string will be centered in a text field. The text
//       field length is defined by input parameter, 'fieldLen'.
//
//
//  fieldLen            int
//     - Defines the length of a text field in which 'strToCenter'
//       will be centered.
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
//     - This returned string contains 'strToCenter' with the
//       necessary left-pad and right-pad number of spaces
//       required for centering. The total length of this string
//       will be equal to input parameter, 'fieldLen'.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//  ePrefix := "TestStrOps_StrCenterInStr_02() "
//  strToCenter := "Hello"
//  fieldLen := 15
//
//  su := StrMech{}
//  centeredStr, err := su.strCenterInStr(
//  strToCenter,
//  fieldLen,
//  ePrefix)
//
//  ---------------------------------------------
//                               123456789012345
//  centeredStr is now equal to "     Hello     "
//  'Hello' is centered in a field of length 15
//  with left and right pad of 5-spaces.
//
func (sMechMolecule *strMechMolecule) strCenterInStr(
	strToCenter string,
	fieldLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref("strMechMolecule." +
		"strCenterInStr()")

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToCenter) {
		return strToCenter,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToCenter' is All White "+
				"Space or an EMPTY String!\n",
				ePrefix.String())
	}

	sLen := len(strToCenter)

	if sLen > fieldLen {
		return strToCenter,
			fmt.Errorf("%v\n"+
				"Error: 'fieldLen' = '%v' strToCenter Length= '%v'.\n"+
				"'fieldLen' is shorter than 'strToCenter' Length!\n",
				ePrefix.String(),
				fieldLen,
				sLen)
	}

	if sLen == fieldLen {
		return strToCenter, nil
	}

	leftPadCnt := (fieldLen - sLen) / 2

	leftPadStr := strings.Repeat(" ", leftPadCnt)

	rightPadCnt := fieldLen - sLen - leftPadCnt

	rightPadStr := ""

	if rightPadCnt > 0 {
		rightPadStr = strings.Repeat(" ", rightPadCnt)
	}

	return leftPadStr + strToCenter + rightPadStr, nil
}

// strLeftJustify - Creates a new string, left-justified, within a
// a wider text field or output string. The text to be left
// justified is specified by input parameter 'strToJustify'. The
// length of the output string is defined by input parameter,
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the left side of the
// output string and spaces are padded to the right in order to
// create a string with total length of 'fieldLen'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//     - The text content which will be left-justified in the
//       output string returned by this method.
//
//  fieldLen            int
//     - Defines the length of the output string in which input
//       parameter 'strToJustify' will be left-justified.
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
//     - The output string resulting from the 'left-justify'
//       operation. Input parameter, 'strToJustify' will be
//       left-justified in this output string which will have a
//       total string length as defined by input parameter,
//       'fieldLen'.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  ePrefix := "TestStrOps_StrLeftJustify_01() "
//  fieldLen = 15
//  strToJustify    = "Hello World"
//  su := StrMech{}
//  justifiedStr, err := su.strLeftJustify(
//                           strToJustify,
//                           fieldLen,
//                           ePrefix)
//
//  ------------------------------------------------
//                                  123456789012345
//  'justifiedStr' is now equal to "Hello World    "
//  The string length of 'justifiedStr' is 15
//
//
func (sMechMolecule *strMechMolecule) strLeftJustify(
	strToJustify string,
	fieldLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechMolecule." +
			"strLeftJustify()")

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToJustify) {
		return strToJustify,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToJustify' is "+
				"All White Space or an EMPTY String!\n",
				ePrefix.String())
	}

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify,
			fmt.Errorf("%v\n"+
				"Error: Length of string to left justify is '%v'.\n"+
				"However, 'fieldLen' is less.\n"+
				"'fieldLen'= '%v'\n",
				ePrefix.String(),
				strLen, fieldLen)
	}

	rightPadLen := fieldLen - strLen

	rightPadStr := strings.Repeat(" ", rightPadLen)

	return strToJustify + rightPadStr, nil
}

// strPadLeftToCenter - Returns a blank string which allows
// centering of the target string in a fixed length field. A blank
// string is defined here as a string consisting entirely of white
// space characters (' ' 0x20). This string can be prefixed or
// added to the beginning of a text string to achieve a centering
// effect when displayed.
//
// Therefore the blank string returned by this method only
// constitutes the left-padding of white space necessary to center
// the input parameter, 'strToCenter'. It does not include the
// right padding of white space.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToCenter         string
//    - The content or text string which will be centered.
//
//
//  fieldLen            int
//     - The total length of the text field in which 'strToCenter'
//       will be centered.
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
//     - The output string resulting from the string centering
//       operation. This string will consist entirely of white
//       space (' ' 0x20 characters). It will represent the left
//       padding necessary to center the text string,
//       'strToCenter'. See 'Example Usage' section below.
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
//
// ------------------------------------------------------------------------
//
// Example Usage
//
// Assume that total field length ('fieldlen') is 70. Further
// assume that the string to Center ('strToCenter') is
// 10-characters in length. In order to center a 10-character
// string in a 70-character field, 30-space characters would need
// to be positioned on each side of the string to center. This
// method only returns the left margin, or in this example, a
// string consisting of 30-spaces.
//
//
//  ePrefix := "TestStrOps_StrPadLeftToCenter_02() "
//  //              12345
//  strToCenter := "Hello"
//  fieldLen := 15
//  su := StrMech{}
//  padStr, err := su.StrPadLeftToCenter(
//                   strToCenter,
//                   fieldLen,
//                   ePrefix)
//
//  -------------------------------------------------------------
//                          12345
//  'padStr' is now equal to "     "
//  'padStr' consists of 5-spaces.
//  padStr + strToCenter will yield a centered string.
//
//
func (sMechMolecule *strMechMolecule) strPadLeftToCenter(
	strToCenter string,
	fieldLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechMolecule." +
			"strPadLeftToCenter()")

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToCenter) {
		return strToCenter,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToCenter' is "+
				"All White Space or an EMPTY String!\n",
				ePrefix.String())
	}

	sLen := sOpsQuark.getRuneCountInStr(strToCenter)

	if sLen > fieldLen {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input Parameter String To Center "+
				"('strToCenter')\n"+
				"is longer than Field Length.\n",
				ePrefix.String())
	}

	if sLen == fieldLen {
		return "", nil
	}

	margin := (fieldLen - sLen) / 2

	return strings.Repeat(" ", margin), nil
}

// strRightJustify - Creates a new string, right-justified, within
// a wider text field or output string. The text to be right
// justified is specified by input parameter 'strToJustify'. The
// length of the output string is defined by input parameter,
// 'fieldLen'.
//
// Input parameter 'strToJustify' is placed on the right side of
// the output string and spaces are padded to the left in order to
// create a string with total length of 'fieldLen'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//    - The content or text string which will be right justified.
//
//
//  fieldLen            int
//     - The total length of the text field in which 'strToCenter'
//       will be right-justified.
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
//     - The output string resulting from the 'right-justify'
//       operation. Input parameter, 'strToJustify' will be
//       right-justified in this output string which will have a
//       total string length as defined by input parameter,
//       'fieldLen'.
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
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  If the total field length ('fieldLen') is specified as
//  10-characters and the length of string to justify
//  ('strToJustify') is 5-characters, then this method would return
//  a string consisting of 5-space characters plus the
//  'strToJustify'.
//
//  ePrefix := "TestStrOps_StrRightJustify_05() "
//  strToJustify := "12345"
//  fieldLen := 10
//
//  su := StrMech{}
//  strRightJustified, err :=
//   su.StrRightJustify(
//               strToJustify,
//               fieldLen,
//               ePrefix)
//
//  --------------------------------------------------------
//                                       1234567890
//  'strRightJustified' is now equal to "     12345"
//  The string length of 'strRightJustified' is 10
//
//
func (sMechMolecule *strMechMolecule) strRightJustify(
	strToJustify string,
	fieldLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechMolecule.lock == nil {
		sMechMolecule.lock = new(sync.Mutex)
	}

	sMechMolecule.lock.Lock()

	defer sMechMolecule.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"sMechMolecule." +
			"strRightJustify()")

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToJustify) {
		return strToJustify,
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToJustify' is "+
				"All White Space or an EMPTY String!\n",
				ePrefix.String())
	}

	strLen := len(strToJustify)

	if fieldLen == strLen {
		return strToJustify, nil
	}

	if fieldLen < strLen {
		return strToJustify,
			fmt.Errorf("%s\n"+
				"Error: Length of string to right justify is '%v'.\n"+
				"However, 'fieldLen' is less.\n"+
				"'fieldLen'= '%v'\n",
				ePrefix.String(),
				strLen, fieldLen)
	}

	// fieldLen must be greater than strLen
	lefPadCnt := fieldLen - strLen

	leftPadStr := strings.Repeat(" ", lefPadCnt)

	return leftPadStr + strToJustify, nil
}

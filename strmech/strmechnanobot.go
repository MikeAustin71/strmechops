package strmech

import (
	"fmt"
	"strings"
	"sync"
)

type strOpsNanobot struct {
	lock *sync.Mutex
}

// strCenterInStrLeft - returns a string which includes a left pad blank string
// plus the original string. It does NOT include the Right pad blank string.
//
// Nevertheless, the complete string will effectively center the original string
// in a field of specified length.
//
// Example:
//   In this example the total field length is 15. The length of the test string,
//   "Hello", is 5. In order to center the test string in a field with length of 15,
//   there will be 5-spaces on the left and 5-spaces on the right. This method will
//   compute the left-pad spaces necessary to center the string in a field with length
//   of 15, but will only include the padded left margin of 5-spaces. It will NOT
//   include the trailing 5-spaces on the right.
//
//   In the following example, the final result string will substitute the'@' character
//   for the white space character (0x20) in order to illustrate the padding added by
//   this method.
//
//    strToCenter     = "Hello"
//    fieldLen        = 15
//    Returned String = "@@@@@Hello" or "     Hello"
//
func (sOpsNanobot *strOpsNanobot) strCenterInStrLeft(
	strToCenter string,
	fieldLen int,
	ePrefix string) (
	string,
	error) {

	if sOpsNanobot.lock == nil {
		sOpsNanobot.lock = new(sync.Mutex)
	}

	sOpsNanobot.lock.Lock()

	defer sOpsNanobot.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsNanobot.strCenterInStrLeft()\n"

	sOpsQuark := strOpsQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToCenter) {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToCenter' is All White Space or an EMPTY String!\n",
				ePrefix)
	}

	if fieldLen < len(strToCenter) {
		return "",
			fmt.Errorf(ePrefix+
				"\n"+
				"Error: Input parameter 'fieldLen' is less than length of 'strToCenter'.\n"+
				"strToCenter length='%v'\n"+
				"fieldLen='%v'\n",
				len(strToCenter), fieldLen)
	}

	sOpsMolecule := strOpsMolecule{}

	pad, err := sOpsMolecule.strPadLeftToCenter(
		strToCenter,
		fieldLen,
		ePrefix)

	if err != nil {
		return "",
			fmt.Errorf(ePrefix+
				"\n"+
				"Error returned by sops.StrPadLeftToCenter(strToCenter, fieldLen).\n"+
				"Error='%v'\n", err.Error())
	}

	return pad + strToCenter, nil

}

// justifyTextInStrField - Creates a and returns a new string text
// field with text 'strToJustify' positioned inside that new string
// in accordance with the string justification formatting passed in
// input parameter, 'textJustify'.
//
// 'textJustify' will specify either 'Right-Justify', 'Left-Justify'
// or 'Center'. Based on this specification, the newly created and
// returned text string will contain 'strToJustify' properly
// formatted as 'Right-Justified', 'Left-Justified' or 'Centered'.
//
// If the length of 'strToJustify' is greater than the output field
// length value, 'fieldLen', this method will increase the value of
// 'fieldLen' to match the length of 'strToJustify'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  strToJustify        string
//    - The string content or text which will be positioned and
//      formatted inside the returned string ('justifiedStr')
//      according to the text justification specification contained
//      in input parameter, 'textJustify'.
//
//      If 'strToJustify' is a zero length empty string, this
//      method will return a string of white-space characters
//      equal in length to the value of input parameter,
//      'fieldLen'.
//
//
//  fieldLen            int
//     - The total length of the returned text field
//       ('justifiedStr') in which 'strToJustify' will be positioned
//       and formatted. If this 'fieldLen' value is less than the
//       length of 'strToJustify', 'fieldLen' will be automatically
//       increased to a value equal to the length of 'strToJustify'.
//
//
//  textJustify         TextJustify
//     - An enumeration value used to specify the type of text
//       formatting which will be applied to 'strToJustify' when
//       it is positioned inside of the returned text string,
//       'justifiedStr'. This enumeration value must be one
//       of the three following format specifications:
//
//       1. Left   - Signals that the text justification format is
//                   set to 'Left-Justify'. Strings within text
//                   fields will be flush with the left margin.
//                          Example: "TextString      "
//
//       2. Right  - Signals that the text justification format is
//                   set to 'Right-Justify'. Strings within text
//                   fields will terminate at the right margin.
//                          Example: "      TextString"
//
//       3. Center - Signals that the text justification format is
//                   is set to 'Centered'. Strings will be positioned
//                   in the center of the text field equidistant
//                   from the left and right margins.
//                           Example: "   TextString   "
//
//
//  ePrefix             string
//     - This is an error prefix which is included in all returned
//       error messages. Usually, it contains the names of the calling
//       method or methods. Be sure to leave a space at the end
//       of 'ePrefix'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  justifiedStr        string
//     - The output string resulting from the text justification
//       operation described above. Input parameter, 'strToJustify'
//       will be formatted in this output string according to the
//       format specification defined by input parameter,
//       'textJustify'. The length of this output string is
//       controlled by input parameter, 'fieldLen'.
//
//
//  err                 error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the input parameter
//       'ePrefix' (error prefix) will be inserted or prefixed at
//       the beginning of the error message.
//
//
// ------------------------------------------------------------------------
//
// Example Usage
//
//
//  ePrefix := "TestStrOps_StrJustify_01() "
//  strToJustify := "12345"
//  fieldLen := 10
//
//  su := StrMech{}
//  strJustified, err :=
//   su.justifyTextInStrField(
//               strToJustify,
//               fieldLen,
//               TextJustify(0).Right()
//               ePrefix)
//
//  --------------------------------------------------------
//                                  1234567890
//  'strJustified' is now equal to "     12345"
//  The string length of 'strJustified' is 10
//
func (sOpsNanobot *strOpsNanobot) justifyTextInStrField(
	strToJustify string,
	fieldLen int,
	textJustify TextJustify,
	ePrefix string) (
	justifiedStr string,
	err error) {

	if sOpsNanobot.lock == nil {
		sOpsNanobot.lock = new(sync.Mutex)
	}

	sOpsNanobot.lock.Lock()

	defer sOpsNanobot.lock.Unlock()

	if len(ePrefix) > 0 {
		ePrefix += "\n"
	}

	ePrefix += "strOpsNanobot.justifyTextInStrField()\n "
	justifiedStr = ""

	if !textJustify.XIsValid() {
		err = fmt.Errorf(ePrefix+
			"Error: Input parameter 'textJustify' is invalid!\n"+
			"textJustify integer value == '%v'\n",
			textJustify.XValueInt())

		return justifiedStr, err
	}

	lenStrToJustify := len(strToJustify)

	if fieldLen < 1 &&
		lenStrToJustify == 0 {

		err = fmt.Errorf(ePrefix +
			"Error: Input parameters 'strToJustify' and 'fieldLen' are invalid!\n")

		return justifiedStr, err

	} else if fieldLen < 1 &&
		lenStrToJustify > 0 {

		justifiedStr = strToJustify

		return justifiedStr, err

	} else if fieldLen > 0 && lenStrToJustify == 0 {

		justifiedStr = strings.Repeat(justifiedStr, fieldLen)

		return justifiedStr, err

	}

	if lenStrToJustify > fieldLen {
		justifiedStr = strToJustify
		return justifiedStr, err
	}

	sOpsMolecule := strOpsMolecule{}

	switch textJustify {
	case TextJustify(0).Left():

		justifiedStr,
			err = sOpsMolecule.strLeftJustify(
			strToJustify,
			fieldLen,
			ePrefix)

	case TextJustify(0).Right():

		justifiedStr,
			err = sOpsMolecule.strRightJustify(
			strToJustify,
			fieldLen,
			ePrefix)

	case TextJustify(0).Center():

		justifiedStr,
			err = sOpsMolecule.strCenterInStr(
			strToJustify,
			fieldLen,
			ePrefix)

	default:
		err = fmt.Errorf(ePrefix+""+
			"Invalid 'textJustify' value!\n"+
			"'textJustify' integer value == '%v'\n",
			textJustify.XValueInt())
	}

	return justifiedStr, err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type strMechNanobot struct {
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
func (sMechNanobot *strMechNanobot) strCenterInStrLeft(
	strToCenter string,
	fieldLen int,
	ePrefix *ePref.ErrPrefixDto) (
	string,
	error) {

	if sMechNanobot.lock == nil {
		sMechNanobot.lock = new(sync.Mutex)
	}

	sMechNanobot.lock.Lock()

	defer sMechNanobot.lock.Unlock()

	ePrefix.SetEPref(
		"strMechNanobot." +
			"strCenterInStrLeft()")

	sOpsQuark := strMechQuark{}

	if sOpsQuark.isEmptyOrWhiteSpace(strToCenter) {
		return "",
			fmt.Errorf("%v\n"+
				"Error: Input parameter 'strToCenter' is All White Space or an EMPTY String!\n",
				ePrefix.String())
	}

	if fieldLen < len(strToCenter) {
		return "",
			fmt.Errorf("%s\n"+
				"Error: Input parameter 'fieldLen' is less than length of 'strToCenter'.\n"+
				"strToCenter='%v'\n"+
				"strToCenter length='%v'\n"+
				"fieldLen='%v'\n",
				ePrefix.String(),
				strToCenter,
				len(strToCenter),
				fieldLen)
	}

	pad, err := strMechMolecule{}.ptr().
		strPadLeftToCenter(
			strToCenter,
			fieldLen,
			ePrefix.XCpy(
				fmt.Sprintf(
					"\nstrToCenter='%v'\n"+
						"fieldLen='%v'\n",
					strToCenter,
					fieldLen)))

	if err != nil {
		return "", err
	}

	return pad + strToCenter, err
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
//      according to the field length and text justification
//      specifications contained in input parameters, 'fieldLen'
//      and 'textJustify'.
//
//      If 'strToJustify' is a zero length or empty string, this
//      method will return an error.
//
//
//  fieldLen            int
//     - The total length of the returned text field
//       ('justifiedStr') in which 'strToJustify' will be
//       positioned and formatted. If this 'fieldLen' value is less
//       than the length of 'strToJustify', 'fieldLen' will be
//       automatically increased to a value equal to the length of
//       'strToJustify'.
//
//       To automatically set the value of 'fieldLen' to the length
//       of 'strToJustify', set this parameter to a value of minus one
//       (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textJustify         TextJustify
//     - An enumeration which specifies the justification of the
//       'strToJustify' within the field specified by 'fieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a string to justify, field length and a 'textJustify'
//       object of type TextJustify. This is because a string to
//       justify with a field length equal to or less than the
//       length of the string to justify will never use text
//       justification. In these cases, text justification is
//       completely ignored.
//
//       However, if the field length is greater than the length of
//       the string to justify, text justification MUST be equal to
//       one of these three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
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
func (sMechNanobot *strMechNanobot) justifyTextInStrField(
	strToJustify string,
	fieldLen int,
	textJustify TextJustify,
	ePrefix *ePref.ErrPrefixDto) (
	justifiedStr string,
	err error) {

	if sMechNanobot.lock == nil {
		sMechNanobot.lock = new(sync.Mutex)
	}

	sMechNanobot.lock.Lock()

	defer sMechNanobot.lock.Unlock()

	if ePrefix == nil {
		ePrefix = ePref.ErrPrefixDto{}.Ptr()
	} else {
		ePrefix = ePrefix.CopyPtr()
	}

	ePrefix.SetEPref(
		"strMechNanobot." +
			"justifyTextInStrField()")

	justifiedStr = ""

	lenStrToJustify := len(strToJustify)

	if lenStrToJustify == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strToJustify' is an empty string\n"+
			"with zero string length!\n",
			ePrefix.String())

		return justifiedStr, err
	}

	if fieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is less than minus one (-1)!\n"+
			"'fieldLen' controls the length of the Text Output Field,\n"+
			"'justifiedStr'\n",
			ePrefix.String())

		return justifiedStr, err
	}

	if fieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLen' is greater than one-million (1,000,000)!\n"+
			"'fieldLen' controls the length of the Text Output Field,\n"+
			"'justifiedStr'\n",
			ePrefix.String())

		return justifiedStr, err
	}

	if fieldLen <= lenStrToJustify {
		justifiedStr = strToJustify
		return justifiedStr, err
	}

	if !textJustify.XIsValid() {
		err = fmt.Errorf("%v\n"+
			"Error: Text Justification is INVALID!\n"+
			"Text Justification MUST be set to\n"+
			"Left, Right or Center.\n"+
			"String To Justify Value = '%v'\n"+
			"Field Length = '%v'\n"+
			"'textJustify'  String Value = '%v'\n"+
			"'textJustify' Integer Value = '%v'\n",
			ePrefix.String(),
			strToJustify,
			fieldLen,
			textJustify.String(),
			textJustify.XValueInt())

		return justifiedStr, err
	}

	sMechMolecule := strMechMolecule{}

	ePrefix.SetCtx(fmt.Sprintf(
		"\n"+
			"strToJustify='%v'\n"+
			"fieldLen='%v'\n",
		strToJustify,
		fieldLen))

	switch textJustify {
	case TextJustify(0).Left():

		justifiedStr,
			err = sMechMolecule.strLeftJustify(
			strToJustify,
			fieldLen,
			ePrefix)

	case TextJustify(0).Right():

		justifiedStr,
			err = sMechMolecule.strRightJustify(
			strToJustify,
			fieldLen,
			ePrefix)

	case TextJustify(0).Center():

		justifiedStr,
			err = sMechMolecule.strCenterInStr(
			strToJustify,
			fieldLen,
			ePrefix)

	default:
		err = fmt.Errorf("%v\n"+
			"Invalid 'textJustify' value!\n"+
			"'textJustify'  String Value = '%v'\n"+
			"'textJustify' Integer Value = '%v'\n",
			ePrefix.String(),
			textJustify.String(),
			textJustify.XValueInt())
	}

	return justifiedStr, err
}

// ptr - Returns a pointer to a new instance of
// strMechNanobot.
//
func (sMechNanobot strMechNanobot) ptr() *strMechNanobot {

	if sMechNanobot.lock == nil {
		sMechNanobot.lock = new(sync.Mutex)
	}

	sMechNanobot.lock.Lock()

	defer sMechNanobot.lock.Unlock()

	return &strMechNanobot{
		lock: new(sync.Mutex),
	}
}

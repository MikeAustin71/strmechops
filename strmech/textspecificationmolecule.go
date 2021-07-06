package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type textSpecificationMolecule struct {
	lock *sync.Mutex
}

// getFormattedText - Formats text using text string, field
// length and text justification values.
//
// If the length of the text label string is zero and the field
// length is zero this method returns an empty string.
//
// If the length of the text label string is zero and the field
// length is greater than zero, this method returns a string with
// a length equal to field length and content equal to white space
// (the space character " " x field length).
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textStr             string
//     - Contains the text context which will be formatted within a
//       text field specified by input parameter 'fieldLen'.
//
//  fieldLen            int
//     - The length of the text field within which input parameter
//       'textStr' will be positioned.
//
//
//  textJustify         TextJustify
//     - Specifies the text justification which will be applied
//       when positioning 'textStr' within the text field specified
//       by 'fieldLen'. 'textJustify' MUST BE set to one of the
//       three following values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//
//  errPrefDto          *ePref.ErrPrefixDto
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
//     - The formatted text is returned as a string. If an error
//       occurs, the error message is included in this string.
//
func (txtSpecMolecule *textSpecificationMolecule) getFormattedText(
	textStr string,
	fieldLen int,
	textJustify TextJustify,
	errPrefDto *ePref.ErrPrefixDto) string {

	if txtSpecMolecule.lock == nil {
		txtSpecMolecule.lock = new(sync.Mutex)
	}

	txtSpecMolecule.lock.Lock()

	defer txtSpecMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textSpecificationMolecule.getFormattedText()",
		"")

	if err != nil {
		return fmt.Sprintf("%v\n", err.Error())
	}

	lenTxtLabel := len(textStr)

	if lenTxtLabel == 0 &&
		fieldLen <= 0 {
		return ""
	}

	if lenTxtLabel == 0 &&
		fieldLen > 0 {
		return strings.Repeat(" ", fieldLen)
	}

	if !textJustify.XIsValid() {
		return fmt.Sprintf("%v\n"+
			"Text Justification Specification is INVALID!\n"+
			"Text Justification Integer Value = '%v'\n",
			ePrefix.String(),
			textJustify.XValueInt())
	}

	var result string

	result, err = strMechNanobot{}.ptr().
		justifyTextInStrField(
			textStr,
			fieldLen,
			textJustify,
			ePrefix)

	if err != nil {
		return fmt.Sprintf("%v", err.Error())
	}

	return result
}

// ptr - Returns a pointer to a new instance of
// textSpecificationMolecule.
//
func (txtSpecMolecule textSpecificationMolecule) ptr() *textSpecificationMolecule {

	if txtSpecMolecule.lock == nil {
		txtSpecMolecule.lock = new(sync.Mutex)
	}

	txtSpecMolecule.lock.Lock()

	defer txtSpecMolecule.lock.Unlock()

	return &textSpecificationMolecule{
		lock: new(sync.Mutex),
	}
}

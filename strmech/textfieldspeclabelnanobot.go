package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecLabelNanobot struct {
	lock *sync.Mutex
}

// setTextFieldLabel - Receives a pointer to an instance of
// TextFieldSpecLabel and proceeds to reset the data values
// based on the input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The pre-existing data fields for input parameter
// 'textFieldLabel' will be deleted and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	textFieldLabel             *TextFieldSpecLabel
//	   - A pointer to an instance of TextFieldSpecLabel. All the
//	     internal member variable data values will be deleted and
//	     reset based on the following input parameters.
//
//
//	textLabel                  []rune
//	   - An array of runes or text characters which is used to
//	     generate string content for display as a text label.
//
//	     If this parameter is submitted as a 'nil' value or a zero
//	    length array, an error will be returned.
//
//
//	fieldLen                   int
//	   - The length of the text field in which the 'textLabel' will
//	     be displayed. If 'fieldLen' is less than the length of the
//	     'textLabel' rune array, it will be automatically set equal
//	     to the 'textLabel' rune array length.
//
//	     To automatically set the value of 'fieldLen' to the length
//	     of 'textLabel', set this parameter to a value of minus one
//	     (-1).
//
//	     If this parameter is submitted with a value less than
//	     minus one (-1) or greater than 1-million (1,000,000), an
//	     error will be returned.
//
//
//	textJustification          TextJustify
//	   - An enumeration which specifies the justification of the
//	     'textLabel' within the field specified by 'fieldLen'.
//
//	     Text justification can only be evaluated in the context of
//	     a text label, field length and a 'textJustification'
//	     object of type TextJustify. This is because text labels
//	     with a field length equal to or less than the length of
//	     the text label never use text justification. In these
//	     cases, text justification is completely ignored.
//
//	     If the field length is greater than the length of the text
//	     label, text justification must be equal to one of these
//	     three valid values:
//	         TextJustify(0).Left()
//	         TextJustify(0).Right()
//	         TextJustify(0).Center()
//
//	     You can also use the abbreviated text justification
//	     enumeration syntax as follows:
//
//	         TxtJustify.Left()
//	         TxtJustify.Right()
//	         TxtJustify.Center()
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
func (txtFieldLabelNanobot textFieldSpecLabelNanobot) setTextFieldLabel(
	textFieldLabel *TextFieldSpecLabel,
	textLabel []rune,
	fieldLen int,
	textJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFieldLabelNanobot.lock == nil {
		txtFieldLabelNanobot.lock = new(sync.Mutex)
	}

	txtFieldLabelNanobot.lock.Lock()

	defer txtFieldLabelNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecLabelNanobot."+
			"setTextFieldLabel()",
		"")

	if err != nil {
		return err
	}

	if textFieldLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFieldLabel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if textLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is invalid!\n"+
			"The 'textLabel' rune array is empty and has a value of 'nil'.\n",
			ePrefix.String())

		return err
	}

	if len(textLabel) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLabel' is invalid!\n"+
			"The 'textLabel' rune array is empty with a length of zero.\n",
			ePrefix.String())

		return err
	}

	txtLabelElectron := textFieldSpecLabelElectron{}

	_,
		err = txtLabelElectron.isTextLabelValid(
		textLabel,
		ePrefix.XCpy("textLabel"))

	if err != nil {
		return err
	}

	err = txtLabelElectron.isFieldLengthValid(
		fieldLen,
		ePrefix.XCpy("fieldLen"))

	if err != nil {
		return err
	}

	err = txtLabelElectron.isTextJustificationValid(
		textLabel,
		fieldLen,
		textJustification,
		ePrefix.XCpy("textJustification"))

	if err != nil {
		return err
	}

	err = strMechPreon{}.ptr().copyRuneArrays(
		&textFieldLabel.textLabel,
		&textLabel,
		true,
		ePrefix.XCpy(
			"textLabel->textFieldLabel.textLabel"))

	if err != nil {
		return err
	}

	textFieldLabel.fieldLen = fieldLen

	textFieldLabel.textJustification = textJustification

	textFieldLabel.textLineReader = nil

	return err
}

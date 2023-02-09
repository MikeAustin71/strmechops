package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

type textFieldSpecDateTimeMechanics struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecDateTimeMechanics.
func (txtFieldDateTimeMechanics textFieldSpecDateTimeMechanics) ptr() *textFieldSpecDateTimeMechanics {

	if txtFieldDateTimeMechanics.lock == nil {
		txtFieldDateTimeMechanics.lock = new(sync.Mutex)
	}

	txtFieldDateTimeMechanics.lock.Lock()

	defer txtFieldDateTimeMechanics.lock.Unlock()

	return &textFieldSpecDateTimeMechanics{
		lock: new(sync.Mutex),
	}
}

// setTextFieldDateTime - Receives a pointer to an instance of
// TextFieldSpecDateTime and proceeds to reset the data values
// based on the input parameters.
//
// The passed instance of TextFieldSpecDateTime serves as a Text
// Field Specification. Text Field Specifications are designed to
// be configured within a single line of text. That line of text
// can then be used for text displays, file output or printing.
//
// Type TextLineSpecStandardLine can be used to compose a line of
// text consisting of multiple Text Field Specifications like
// TextFieldSpecDateTime. Text Field Specifications like
// TextFieldSpecDateTime are therefore used as the components or
// building blocks for constructing a single lines of text.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The pre-existing data fields for input parameter
// 'dateTimeTxtField' will be deleted and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	dateTimeTxtField           *TextFieldSpecDateTime
//	   - A pointer to an instance of TextFieldSpecDateTime. All the
//	     internal member variable data values will be deleted and
//	     reset based on the following input parameters.
//
//
//	dateTime                   time.Time
//	   - A valid date time value which is used to generate a
//	     formatted date/time text string. Type time.Time is part of
//	     the Golang time package:
//	            https://pkg.go.dev/time.
//
//	     If this parameter is submitted as a zero value, an error
//	     will be returned.
//
//
//	fieldLen                   int
//	   - The length of the text field in which the formatted
//	     'dateTime' value will be displayed.
//
//	     If 'fieldLen' is less than the length of the formatted
//	     'dateTime' string, it will be automatically set equal to
//	     the formatted 'dateTime' string length.
//
//	     If 'fieldLen' is greater than the length of the formatted
//	     'dateTime' string, 'dateTime' will be positioned within a
//	     text field with a length equal to 'fieldLen'. In this
//	     case, the position of the 'dateTime' string within the
//	     text field will be controlled by the text justification
//	     value contained in parameter, 'textJustification'.
//
//	     To automatically set the value of 'fieldLen' to the length
//	     of the formatted 'dateTime', set this parameter to a value
//	     of minus one (-1).
//
//	     If this parameter is submitted with a value less than
//	     minus one (-1) or greater than 1-million (1,000,000), an
//	     error will be returned.
//
//
//	dateTimeFormat             string
//	   - This string holds the date/time format parameters used to
//	     format the 'dateTime' value when generating a 'dateTime'
//	     text string. The formatted 'dateTime' text string is
//	     generated by type TextFieldSpecDateTime for use in text
//	     displays, file output or printing.
//
//	     The date/time format is documented in the Golang time.Time
//	     package, https://pkg.go.dev/time. The format operations are
//	     documented at https://pkg.go.dev/time#Time.Format .
//
//	     If this parameter is submitted as an empty string,
//	     parameter 'dateTimeFormat' will be assigned a default
//	     value of "2006-01-02 15:04:05.000000000 -0700 MST".
//
//
//	textJustification          TextJustify
//	   - An enumeration which specifies the justification of the
//	     'dateTime' string within a text field with a field length
//	     specified by parameter 'fieldLen'.
//
//	     Text justification can only be evaluated in the context of
//	     a 'dateTime' text string, field length and a
//	     'textJustification' object of type TextJustify. This is
//	     because a field length ('fieldLen') value equal to or less
//	     than the length of the 'dateTime' text string will never
//	     use text justification. In these cases, text justification
//	     is completely ignored because the length of the text field
//	     ('fieldLen') is automatically set equal to the length of
//	     the 'dateTime' text string.
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
func (txtFieldDateTimeMechanics *textFieldSpecDateTimeMechanics) setTextFieldDateTime(
	dateTimeTxtField *TextFieldSpecDateTime,
	dateTime time.Time,
	fieldLen int,
	dateTimeFormat string,
	textJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtFieldDateTimeMechanics.lock == nil {
		txtFieldDateTimeMechanics.lock = new(sync.Mutex)
	}

	txtFieldDateTimeMechanics.lock.Lock()

	defer txtFieldDateTimeMechanics.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecDateTimeMechanics."+
			"setTextFieldDateTime()",
		"")

	if err != nil {
		return err
	}

	if dateTimeTxtField == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dateTimeTxtField' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if dateTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dateTime' has "+
			"a Zero Value!\n",
			ePrefix.String())

		return err
	}

	txtLabelElectron := textFieldSpecLabelElectron{}

	if len(dateTimeFormat) == 0 {

		dateTimeFormat =
			new(textSpecificationMolecule).
				getDefaultDateTimeFormat()
	}

	dateTimeStr := dateTime.Format(dateTimeFormat)

	lenDateTimeStr := len(dateTimeStr)

	if fieldLen >= -1 && fieldLen <= lenDateTimeStr {
		fieldLen = lenDateTimeStr
	}

	err = txtLabelElectron.isFieldLengthValid(
		fieldLen,
		ePrefix.XCpy("fieldLen"))

	if err != nil {
		return err
	}

	err = txtLabelElectron.isTextJustificationValid(
		[]rune(dateTimeStr),
		fieldLen,
		textJustification,
		ePrefix.XCpy("textJustification"))

	if err != nil {
		return err
	}

	dateTimeTxtField.dateTime = dateTime

	dateTimeTxtField.fieldLen = fieldLen

	dateTimeTxtField.dateTimeFormat = dateTimeFormat

	dateTimeTxtField.textJustification = textJustification

	dateTimeTxtField.textLineReader = nil

	return err
}

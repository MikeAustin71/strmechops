package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

type TextFieldDateTimeDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed.
	//
	// For the TextFieldDateTimeDto Format Type, this value
	// should always be set to:
	//   TxtFieldType.DateTime()       - Valid

	LeftMarginStr string
	// The contents of the string will be used as the left margin
	// for the Text Field.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	FieldDateTime time.Time
	// This time value will be converted to a string and used
	// to populate the Date Time Text Field.

	FieldDateTimeFormat string
	// This string will be used to format the Date/Time value
	// contained in parameter 'FieldDateTime' as text.
	//
	// If 'FieldDateTime' is set to a value greater than zero
	// and this 'FieldDateTimeFormat' string is empty (has a
	// zero length), a default Date/Time format string will be
	// applied as follows:
	//         "2006-01-02 15:04:05.000000000 -0700 MST"

	FieldLength int
	// Used to format Text Fields. This is the length of the
	// text field in which the date time text field will be
	// displayed. If 'FieldLength' is less than the length of
	// the date time text string, it will be automatically set
	// equal to the date time text string length.
	//
	//  If 'FieldLength' is greater than the length of the date
	//  time text string, the 'FieldJustify' parameter will be
	//  used to configure or justify the text within the
	//  boundaries the text field defined by 'FieldLength'.
	//
	// To automatically set the value of 'FieldLength' to the
	// length of the date time text string, set this parameter
	// to a value of minus one (-1).
	//
	// If this parameter is submitted with a value less than
	// minus one (-1) or greater than 1-million (1,000,000),
	// an error will be returned.

	FieldJustify TextJustify
	// An enumeration which specifies the justification of the
	// 'FieldText' string within the text field specified by
	// 'FieldLength'.
	//
	// Text justification can only be evaluated in the context of
	// a text label, field length and a Text Justification object
	// of type TextJustify. This is because text labels with a
	// field length equal to or less than the length of the text
	// label never use text justification. In these cases, text
	// justification is completely ignored.
	//
	// If the field length is greater than the length of the text
	// label, text justification must be equal to one of these
	// three valid values:
	//     TextJustify(0).Left()
	//     TextJustify(0).Right()
	//     TextJustify(0).Center()
	//
	// You can also use the abbreviated text justification
	// enumeration syntax as follows:
	//
	//     TxtJustify.Left()
	//     TxtJustify.Right()
	//     TxtJustify.Center()

	RightMarginStr string
	// The contents of the string will be used as the right margin
	// for the Text Field.
	//
	// If no right margin is required, set 'RightMarginStr' to a
	// zero length or empty string, and no right margin will be
	// created.

	LineTerminator string
	// This string holds the character or characters which will be
	// used to terminate the formatted line of text output.
	//
	// The most common usage sets this string to a new line
	// character ("\n").
	//
	// If no Line Terminator is required, set 'lineTerminator' to
	// a zero length or empty string and no line termination
	// characters will be created.

	MaxLineLength int
	// Set this parameter to minus one -1 to specify an
	// unlimited line length for this text line.

	TurnAutoLineLengthBreaksOn bool
	// When this parameter is set to 'true', text fields which
	// extend beyond the maximum line length ('MaxLineLength')
	// will be terminated with a new line character ('\n') and
	// placed on the following line of text.

	MultiLineLeftMarginStr string
	// The left margin used when a text string exceeds the
	// maximum line length and is separated into multiple text
	// lines. This left margin is applied to the second and
	// all subsequent lines of a multi-line text display. This
	// parameter is only valid when 'TurnAutoLineLengthBreaksOn'
	// is set to 'true' and the initial text string exceeds the
	// maximum line length.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldDateTimeDto ('incomingDateTimeDto') to the data fields
// of the current TextFieldDateTimeDto instance
// ('txtDateTimeDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextFieldDateTimeDto instance
// ('txtDateTimeDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingDateTimeDto'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingDateTimeDto        TextFieldDateTimeDto
//	   - An instance of TextFieldDateTimeDto. This method will NOT
//	     change the data values of member variables contained in
//	     this instance.
//
//	     All data values in this TextFieldDateTimeDto instance
//	     ('incomingDateTimeDto') will be copied to the current
//	     TextFieldDateTimeDto instance ('txtDateTimeDto').
//
//	     No data validation is performed on input parameter,
//	     'incomingDateTimeDto'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtDateTimeDto *TextFieldDateTimeDto) CopyIn(
	incomingDateTimeDto TextFieldDateTimeDto) {

	if txtDateTimeDto.lock == nil {
		txtDateTimeDto.lock = new(sync.Mutex)
	}

	txtDateTimeDto.lock.Lock()

	defer txtDateTimeDto.lock.Unlock()

	_ = textFieldDateTimeDtoNanobot{}.ptr().
		copy(
			txtDateTimeDto,
			&incomingDateTimeDto,
			nil)

	return
}

// CopyOut - Returns a deep copy of the current
// TextFieldDateTimeDto instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFieldDateTimeDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyDateTimeDto        TextFieldDateTimeDto
//	   - This parameter will return a deep copy of the current
//	     TextFieldDateTimeDto instance.
func (txtDateTimeDto *TextFieldDateTimeDto) CopyOut() (
	deepCopyDateTimeDto TextFieldDateTimeDto) {

	if txtDateTimeDto.lock == nil {
		txtDateTimeDto.lock = new(sync.Mutex)
	}

	txtDateTimeDto.lock.Lock()

	defer txtDateTimeDto.lock.Unlock()

	_ = textFieldDateTimeDtoNanobot{}.ptr().
		copy(
			&deepCopyDateTimeDto,
			txtDateTimeDto,
			nil)

	return deepCopyDateTimeDto
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldDateTimeDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextFieldDateTimeDto in an invalid state and unavailable for
// immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFieldDateTimeDto. All member variable
// data values will be reset to their zero or uninitialized states.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtDateTimeDto *TextFieldDateTimeDto) Empty() {

	if txtDateTimeDto.lock == nil {
		txtDateTimeDto.lock = new(sync.Mutex)
	}

	txtDateTimeDto.lock.Lock()

	txtDateTimeDto.FormatType = TxtFieldType.None()

	txtDateTimeDto.LeftMarginStr = ""

	txtDateTimeDto.FieldDateTime = time.Time{}

	txtDateTimeDto.FieldDateTimeFormat = ""

	txtDateTimeDto.FieldLength = -99

	txtDateTimeDto.FieldJustify = TxtJustify.None()

	txtDateTimeDto.RightMarginStr = ""

	txtDateTimeDto.LineTerminator = ""

	txtDateTimeDto.MaxLineLength = -99

	txtDateTimeDto.TurnAutoLineLengthBreaksOn = false

	txtDateTimeDto.MultiLineLeftMarginStr = ""

	txtDateTimeDto.lock.Unlock()

	txtDateTimeDto.lock = nil

}

// Equal - Receives another instance of TextFieldDateTimeDto and
// proceeds to compare the member variables to those of the current
// TextFieldDateTimeDto instance in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingDateTimeDto        TextFieldDateTimeDto
//	   - An incoming instance of TextFieldDateTimeDto. This method
//	     will compare all member variable data values in this
//	     instance against those contained in the current instance
//	     of TextFieldDateTimeDto. If the data values in both
//	     instances are found to be equal in all respects, this
//	     method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingDateTimeDto' are equal in all respects
//	     to those contained in the current instance of
//	     TextFieldDateTimeDto, this method will return a boolean
//	     value of 'true'. Otherwise, a value of 'false' will be
//	     returned to the calling function.
func (txtDateTimeDto *TextFieldDateTimeDto) Equal(
	incomingDateTimeDto TextFieldDateTimeDto) bool {

	if txtDateTimeDto.lock == nil {
		txtDateTimeDto.lock = new(sync.Mutex)
	}

	txtDateTimeDto.lock.Lock()

	defer txtDateTimeDto.lock.Unlock()

	if txtDateTimeDto.FormatType !=
		incomingDateTimeDto.FormatType {

		return false
	}

	if txtDateTimeDto.LeftMarginStr !=
		incomingDateTimeDto.LeftMarginStr {

		return false
	}

	if txtDateTimeDto.FieldDateTime !=
		incomingDateTimeDto.FieldDateTime {

		return false
	}

	if txtDateTimeDto.FieldDateTimeFormat !=
		incomingDateTimeDto.FieldDateTimeFormat {

		return false
	}

	if txtDateTimeDto.FieldLength !=
		incomingDateTimeDto.FieldLength {

		return false
	}

	if txtDateTimeDto.FieldJustify !=
		incomingDateTimeDto.FieldJustify {

		return false
	}

	if txtDateTimeDto.RightMarginStr !=
		incomingDateTimeDto.RightMarginStr {

		return false
	}

	if txtDateTimeDto.LineTerminator !=
		incomingDateTimeDto.LineTerminator {

		return false
	}

	if txtDateTimeDto.MaxLineLength !=
		incomingDateTimeDto.MaxLineLength {

		return false
	}

	if txtDateTimeDto.TurnAutoLineLengthBreaksOn !=
		incomingDateTimeDto.TurnAutoLineLengthBreaksOn {

		return false
	}

	if txtDateTimeDto.MultiLineLeftMarginStr !=
		incomingDateTimeDto.MultiLineLeftMarginStr {

		return false
	}

	return true
}

// textFieldDateTimeDtoNanobot - Provides helper methods for
// TextFieldDateTimeDto.
type textFieldDateTimeDtoNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFieldDateTimeDto to a destination instance of
// TextFieldDateTimeDto.
func (txtDateTimeDtoNanobot textFieldDateTimeDtoNanobot) copy(
	sourceTxtDateTimeDto *TextFieldDateTimeDto,
	destinationTxtDateTimeDto *TextFieldDateTimeDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtDateTimeDtoNanobot.lock == nil {
		txtDateTimeDtoNanobot.lock = new(sync.Mutex)
	}

	txtDateTimeDtoNanobot.lock.Lock()

	defer txtDateTimeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldDateTimeDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtDateTimeDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtDateTimeDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationTxtDateTimeDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtDateTimeDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	destinationTxtDateTimeDto.FormatType =
		sourceTxtDateTimeDto.FormatType

	destinationTxtDateTimeDto.LeftMarginStr =
		sourceTxtDateTimeDto.LeftMarginStr

	destinationTxtDateTimeDto.FieldDateTime =
		sourceTxtDateTimeDto.FieldDateTime

	destinationTxtDateTimeDto.FieldDateTimeFormat =
		sourceTxtDateTimeDto.FieldDateTimeFormat

	destinationTxtDateTimeDto.FieldLength =
		sourceTxtDateTimeDto.FieldLength

	destinationTxtDateTimeDto.FieldJustify =
		sourceTxtDateTimeDto.FieldJustify

	destinationTxtDateTimeDto.RightMarginStr =
		sourceTxtDateTimeDto.RightMarginStr

	destinationTxtDateTimeDto.LineTerminator =
		sourceTxtDateTimeDto.LineTerminator

	destinationTxtDateTimeDto.MaxLineLength =
		sourceTxtDateTimeDto.MaxLineLength

	destinationTxtDateTimeDto.TurnAutoLineLengthBreaksOn =
		sourceTxtDateTimeDto.TurnAutoLineLengthBreaksOn

	destinationTxtDateTimeDto.MultiLineLeftMarginStr =
		sourceTxtDateTimeDto.MultiLineLeftMarginStr

	return err
}

// ptr - Returns a pointer to a new instance of
// textFieldDateTimeDtoNanobot.
func (txtDateTimeDtoNanobot textFieldDateTimeDtoNanobot) ptr() *textFieldDateTimeDtoNanobot {

	if txtDateTimeDtoNanobot.lock == nil {
		txtDateTimeDtoNanobot.lock = new(sync.Mutex)
	}

	txtDateTimeDtoNanobot.lock.Lock()

	defer txtDateTimeDtoNanobot.lock.Unlock()

	return &textFieldDateTimeDtoNanobot{
		lock: new(sync.Mutex),
	}
}

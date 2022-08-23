package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextFieldLabelDto - The Text Field Label Data Transfer Object is
// used to transmit data parameters and specifications required to
// construct a text label string. A text label contains a string of
// text characters.
//
// Typically, Text Label Fields are designed to be configured
// within a line of text. However, with the TextFieldLabelDto,
// users have the option of configuring a Text Label Field as a
// separate stand-alone line of text.
//
// The text string or text label is positioned inside a text field
// with a given field length. Text Justification within this text
// field is controlled by the Text Justification specification
// value which may be set to 'Left', 'Right' or 'Center'.
//
// Text Label Justification Examples:
//   'Hello World  ' - Left Justified, Field Length= 13
//   '  Hello World' - Right Justified, Field Length= 13
//   ' Hello World ' - Centered, Field Length= 13
//
// Text Label Usage Example 1:
//   LeftMarginStr = "" // Empty String
//   FieldText = "Hello"  // Length = 5 characters
//   FieldLength = 7
//   FieldJustify = TxtJustify.Center()
//   RightMarginStr = "" // Empty String
//   LineTerminator = "" // Empty String
//   MaxLineLength = -1
//   TurnAutoLineLengthBreaksOn = false
//   Final Text Label string = " Hello "
//
// Text Label Usage Example 2:
//   LeftMarginStr = "" // Empty String
//   FieldText = "Hello"  // Length = 5 characters
//   FieldLength = 7
//   FieldJustify = TxtJustify.Center()
//   RightMarginStr = "" // Empty String
//   LineTerminator = "\n" // Empty String
//   MaxLineLength = -1
//   TurnAutoLineLengthBreaksOn = false
//   Final Text Label string = " Hello \n"
//
// Text Field Specifications are designed to be configured within a
// line of text. Those lines of text can then be formatted for text
// displays, file output or printing. The type
// TextLineSpecStandardLine can be used to compose a line of text
// consisting of multiple Text Field Specifications like
// TextFieldSpecLabel. Text Field Specifications are therefore
// used as the components or building blocks for single lines of
// text.

type TextFieldLabelDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed.
	//
	// For TextFieldLabelDto Format Type, this value
	// should always be set to:
	//   TxtFieldType.Label()

	LeftMarginStr string
	// The contents of the string will be used as the left margin
	// for the Text Field.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	FieldText string
	// The Text Field string or contents. If this string is empty
	// (has a zero (0) length) it will be automatically converted
	// to a single white space character (" ").
	//
	// This string represents the contents of the Text Field.

	FieldLength int
	// Used to format FieldText string. This is the length of the
	// text field in which the 'FieldText' will be displayed. If
	// 'FieldLength' is less than the length of the 'FieldText'
	// string, it will be automatically set equal to the
	// 'FieldText' string length.
	//
	// If 'FieldLength' is greater than the length of the
	// 'FieldText' text string, the 'FieldJustify' parameter will
	//  be used to configure or justify the text within the
	//  boundaries of the text field defined by 'FieldLength'.
	//
	// To automatically set the value of 'FieldLength' to the
	// length of 'FieldText', set this parameter to a value of
	// minus one (-1).
	//
	// If this parameter is submitted with a value less than
	// minus one (-1) or greater than 1-million (1,000,000), an
	// error will be returned.

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
	// The contents of the string will be used as the right
	// margin for the Text Field.
	//
	// If no right margin is required, set 'RightMarginStr' to
	// a zero length or empty string, and no right margin will
	// be created.

	LineTerminator string
	// This string holds the character or characters which will be
	// used to terminate the formatted line of text output thereby
	// converting this text element into a valid line of text. Line
	// Termination is optional. Populate this string only if this
	// text output should be formatted as a separate line of text.
	//
	// The most common usage sets this string to a new line
	// character ("\n").
	//
	// If no Line Terminator is required, set 'LineTerminator'
	// to a zero length or empty string and no line termination
	// characters will be created.

	MaxLineLength int
	// The maximum length of the line on which this label text
	// will be presented.
	//
	// Set this parameter to minus one (-1) to specify an unlimited
	// line length for this text line.
	//
	// 'MaxLineLength' is used in conjunction with parameter
	// 'TurnAutoLineLengthBreaksOn' to automatically place text
	// fields on separate text lines when that text exceeds the
	// maximum text line length ('MaxLineLength'). Therefore,
	// paramter 'turnAutoLineLengthBreaksOn' controls whether
	// automatic line breaks using 'MaxLineLength' will be
	// applied.
	//
	// If the value of 'maxLineLength' is less than one (1), it
	// will be automatically converted to minus one (-1).
	//
	// Set this parameter to minus one (-1) to specify an unlimited
	// line length for this text line.

	TurnAutoLineLengthBreaksOn bool
	// This parameter controls whether text lines which exceed the
	// maximum line length ('MaxLineLength') are broken up and
	// presented on the following line.
	//
	// To apply automatic line breaking at the maximum line length,
	// set the value of this parameter to 'true'.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldLabelDto ('incomingTxtLabelDto') to the data fields
// of the current TextFieldLabelDto instance
// ('txtLabelDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextFieldLabelDto instance
// ('txtLabelDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingTxtLabelDto'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTxtLabelDto        TextFieldLabelDto
//	   - An instance of TextFieldLabelDto. This method will NOT
//	     change the data values of member variables contained in
//	     this instance.
//
//	     All data values in this TextFieldLabelDto instance
//	     ('incomingTxtLabelDto') will be copied to the current
//	     TextFieldLabelDto instance ('txtLabelDto').
//
//	     No data validation is performed on input parameter,
//	     'incomingTxtLabelDto'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtLabelDto *TextFieldLabelDto) CopyIn(
	incomingTxtLabelDto TextFieldLabelDto) {

	if txtLabelDto.lock == nil {
		txtLabelDto.lock = new(sync.Mutex)
	}

	txtLabelDto.lock.Lock()

	defer txtLabelDto.lock.Unlock()

	_ = textFieldLabelDtoNanobot{}.ptr().copy(
		txtLabelDto,
		&incomingTxtLabelDto,
		nil)

	return
}

// CopyOut - Returns a deep copy of the current TextFieldLabelDto
// instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFieldLabelDto.
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
//	deepCopyTxtLabelDto        TextFieldLabelDto
//	   - This parameter will return a deep copy of the current
//	     TextFieldLabelDto instance.
func (txtLabelDto *TextFieldLabelDto) CopyOut() (
	deepCopyTxtLabelDto TextFieldLabelDto) {

	if txtLabelDto.lock == nil {
		txtLabelDto.lock = new(sync.Mutex)
	}

	txtLabelDto.lock.Lock()

	defer txtLabelDto.lock.Unlock()

	_ = textFieldLabelDtoNanobot{}.ptr().copy(
		&deepCopyTxtLabelDto,
		txtLabelDto,
		nil)

	return deepCopyTxtLabelDto
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldLabelDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextFieldLabelDto in an invalid state and unavailable for
// immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFieldLabelDto. All member variable data
// values will be reset to their zero or uninitialized states.
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
func (txtLabelDto *TextFieldLabelDto) Empty() {

	if txtLabelDto.lock == nil {
		txtLabelDto.lock = new(sync.Mutex)
	}

	txtLabelDto.lock.Lock()

	textFieldLabelDtoAtom{}.ptr().
		empty(
			txtLabelDto)

	txtLabelDto.lock.Unlock()

	txtLabelDto.lock = nil

	return
}

// Equal - Receives another instance of TextFieldLabelDto and
// proceeds to compare the member variables to those of the current
// TextFieldLabelDto instance in order to determine if they are
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
//	incomingTxtLabelDto        TextFieldLabelDto
//	   - An incoming instance of TextFieldLabelDto. This method
//	     will compare all member variable data values in this
//	     instance against those contained in the current instance
//	     of TextFieldLabelDto. If the data values in both
//	     instances are found to be equal in all respects, this
//	     method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingTxtLabelDto' are equal in all respects
//	     to those contained in the current instance of
//	     TextFieldLabelDto, this method will return a boolean
//	     value of 'true'. Otherwise, a value of 'false' will be
//	     returned to the calling function.
func (txtLabelDto *TextFieldLabelDto) Equal(
	incomingTxtLabelDto TextFieldLabelDto) bool {

	if txtLabelDto.lock == nil {
		txtLabelDto.lock = new(sync.Mutex)
	}

	txtLabelDto.lock.Lock()

	defer txtLabelDto.lock.Unlock()

	return textFieldLabelDtoAtom{}.ptr().
		equal(
			txtLabelDto,
			&incomingTxtLabelDto)

}

// textFieldLabelDtoNanobot - Provides helper methods for
// TextFieldLabelDto
type textFieldLabelDtoNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFieldLabelDto to a destination instance of
// TextFieldLabelDto.
func (txtLabelDtoNanobot *textFieldLabelDtoNanobot) copy(
	destinationTxtLabelDto *TextFieldLabelDto,
	sourceTxtLabelDto *TextFieldLabelDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLabelDtoNanobot.lock == nil {
		txtLabelDtoNanobot.lock = new(sync.Mutex)
	}

	txtLabelDtoNanobot.lock.Lock()

	defer txtLabelDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldLabelDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtLabelDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtLabelDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationTxtLabelDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtLabelDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	textFieldLabelDtoAtom{}.ptr().empty(
		destinationTxtLabelDto)

	destinationTxtLabelDto.FormatType =
		sourceTxtLabelDto.FormatType

	destinationTxtLabelDto.LeftMarginStr =
		sourceTxtLabelDto.LeftMarginStr

	destinationTxtLabelDto.FieldText =
		sourceTxtLabelDto.FieldText

	destinationTxtLabelDto.FieldLength =
		sourceTxtLabelDto.FieldLength

	destinationTxtLabelDto.FieldJustify =
		sourceTxtLabelDto.FieldJustify

	destinationTxtLabelDto.RightMarginStr =
		sourceTxtLabelDto.RightMarginStr

	destinationTxtLabelDto.LineTerminator =
		sourceTxtLabelDto.LineTerminator

	destinationTxtLabelDto.MaxLineLength =
		sourceTxtLabelDto.MaxLineLength

	destinationTxtLabelDto.TurnAutoLineLengthBreaksOn =
		sourceTxtLabelDto.TurnAutoLineLengthBreaksOn

	return err
}

// ptr - Returns a pointer to a new instance of
// textFieldLabelDtoNanobot.
func (txtLabelDtoNanobot textFieldLabelDtoNanobot) ptr() *textFieldLabelDtoNanobot {

	if txtLabelDtoNanobot.lock == nil {
		txtLabelDtoNanobot.lock = new(sync.Mutex)
	}

	txtLabelDtoNanobot.lock.Lock()

	defer txtLabelDtoNanobot.lock.Unlock()

	return &textFieldLabelDtoNanobot{
		lock: new(sync.Mutex),
	}
}

// textFieldLabelDtoAtom - Provides helper methods for
// TextFieldLabelDto
type textFieldLabelDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of TextFieldLabelDto
// and proceeds to set all the internal member variables to their
// zero or uninitialized states.
//
// This method will therefore delete all data currently held
// by this instance of TextFieldLabelDto.
func (txtLabelDtoAtom *textFieldLabelDtoAtom) empty(
	txtLabelDto *TextFieldLabelDto) {

	if txtLabelDtoAtom.lock == nil {
		txtLabelDtoAtom.lock = new(sync.Mutex)
	}

	txtLabelDtoAtom.lock.Lock()

	defer txtLabelDtoAtom.lock.Unlock()

	if txtLabelDto == nil {

		return
	}

	txtLabelDto.FormatType = TxtFieldType.None()

	txtLabelDto.LeftMarginStr = ""

	txtLabelDto.FieldText = ""

	txtLabelDto.FieldLength = -99

	txtLabelDto.FieldJustify = TxtJustify.None()

	txtLabelDto.RightMarginStr = ""

	txtLabelDto.LineTerminator = ""

	txtLabelDto.MaxLineLength = -99

	txtLabelDto.TurnAutoLineLengthBreaksOn = false

	return
}

// equal - Receives pointers to two instances of TextFieldLabelDto
// and proceeds to compare all the member data variables for both
// instances.
//
// If the two instances of TextFieldLabelDto are found to be equal
// in all respects, this method will return a boolean value of
// 'true'.
func (txtLabelDtoAtom *textFieldLabelDtoAtom) equal(
	txtLabelDto1 *TextFieldLabelDto,
	txtLabelDto2 *TextFieldLabelDto) bool {

	if txtLabelDtoAtom.lock == nil {
		txtLabelDtoAtom.lock = new(sync.Mutex)
	}

	txtLabelDtoAtom.lock.Lock()

	defer txtLabelDtoAtom.lock.Unlock()

	if txtLabelDto1 == nil ||
		txtLabelDto2 == nil {

		return false
	}

	if txtLabelDto1.FormatType !=
		txtLabelDto2.FormatType {

		return false
	}

	if txtLabelDto1.LeftMarginStr !=
		txtLabelDto2.LeftMarginStr {

		return false
	}

	if txtLabelDto1.FieldText !=
		txtLabelDto2.FieldText {

		return false
	}

	if txtLabelDto1.FieldLength !=
		txtLabelDto2.FieldLength {

		return false
	}

	if txtLabelDto1.FieldJustify !=
		txtLabelDto2.FieldJustify {

		return false
	}

	if txtLabelDto1.RightMarginStr !=
		txtLabelDto2.RightMarginStr {

		return false
	}

	if txtLabelDto1.LineTerminator !=
		txtLabelDto2.LineTerminator {

		return false
	}

	if txtLabelDto1.MaxLineLength !=
		txtLabelDto2.MaxLineLength {

		return false
	}

	if txtLabelDto1.TurnAutoLineLengthBreaksOn !=
		txtLabelDto2.TurnAutoLineLengthBreaksOn {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// textFieldLabelDtoAtom.
func (txtLabelDtoAtom textFieldLabelDtoAtom) ptr() *textFieldLabelDtoAtom {

	if txtLabelDtoAtom.lock == nil {
		txtLabelDtoAtom.lock = new(sync.Mutex)
	}

	txtLabelDtoAtom.lock.Lock()

	defer txtLabelDtoAtom.lock.Unlock()

	return &textFieldLabelDtoAtom{
		lock: new(sync.Mutex),
	}
}

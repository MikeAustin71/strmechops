package strmech

import "sync"

type TextFieldLabelDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed.
	//
	// For TextFieldLabelDto Format Type, this value
	// should always be set to:
	//   TxtFieldType.Label()       - Valid

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
// IMPORTANT
//
// All the data fields in current TextFieldLabelDto instance
// ('txtLabelDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingTxtLabelDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtLabelDto        TextFieldLabelDto
//     - An instance of TextFieldLabelDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this TextFieldLabelDto instance
//       ('incomingTxtLabelDto') will be copied to the current
//       TextFieldLabelDto instance ('txtLabelDto').
//
//       No data validation is performed on input parameter,
//       'incomingTxtLabelDto'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (txtLabelDto *TextFieldLabelDto) CopyIn(
	incomingTxtLabelDto TextFieldLabelDto) {

	if txtLabelDto.lock == nil {
		txtLabelDto.lock = new(sync.Mutex)
	}

	txtLabelDto.lock.Lock()

	defer txtLabelDto.lock.Unlock()

	txtLabelDto.FormatType =
		incomingTxtLabelDto.FormatType

	txtLabelDto.LeftMarginStr =
		incomingTxtLabelDto.LeftMarginStr

	txtLabelDto.FieldText =
		incomingTxtLabelDto.FieldText

	txtLabelDto.FieldLength =
		incomingTxtLabelDto.FieldLength

	txtLabelDto.FieldJustify =
		incomingTxtLabelDto.FieldJustify

	txtLabelDto.RightMarginStr =
		incomingTxtLabelDto.RightMarginStr

	txtLabelDto.LineTerminator =
		incomingTxtLabelDto.LineTerminator

	txtLabelDto.MaxLineLength =
		incomingTxtLabelDto.MaxLineLength

	txtLabelDto.TurnAutoLineLengthBreaksOn =
		incomingTxtLabelDto.TurnAutoLineLengthBreaksOn

	return
}

// CopyOut - Returns a deep copy of the current TextFieldLabelDto
// instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFieldLabelDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  deepCopyTxtLabelDto        TextFieldLabelDto
//     - This parameter will return a deep copy of the current
//       TextFieldLabelDto instance.
//
func (txtLabelDto *TextFieldLabelDto) CopyOut() (
	deepCopyTxtLabelDto TextFieldLabelDto) {

	if txtLabelDto.lock == nil {
		txtLabelDto.lock = new(sync.Mutex)
	}

	txtLabelDto.lock.Lock()

	defer txtLabelDto.lock.Unlock()

	deepCopyTxtLabelDto.FormatType =
		txtLabelDto.FormatType

	deepCopyTxtLabelDto.LeftMarginStr =
		txtLabelDto.LeftMarginStr

	deepCopyTxtLabelDto.FieldText =
		txtLabelDto.FieldText

	deepCopyTxtLabelDto.FieldLength =
		txtLabelDto.FieldLength

	deepCopyTxtLabelDto.FieldJustify =
		txtLabelDto.FieldJustify

	deepCopyTxtLabelDto.RightMarginStr =
		txtLabelDto.RightMarginStr

	deepCopyTxtLabelDto.LineTerminator =
		txtLabelDto.LineTerminator

	deepCopyTxtLabelDto.MaxLineLength =
		txtLabelDto.MaxLineLength

	deepCopyTxtLabelDto.TurnAutoLineLengthBreaksOn =
		txtLabelDto.TurnAutoLineLengthBreaksOn

	return deepCopyTxtLabelDto
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldLabelDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextFieldLabelDto in an invalid state and unavailable for
// immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFieldLabelDto. All member variable data
// values will be reset to their zero or uninitialized states.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (txtLabelDto *TextFieldLabelDto) Empty() {

	if txtLabelDto.lock == nil {
		txtLabelDto.lock = new(sync.Mutex)
	}

	txtLabelDto.lock.Lock()

	txtLabelDto.FormatType = TxtFieldType.None()

	txtLabelDto.LeftMarginStr = ""

	txtLabelDto.FieldText = ""

	txtLabelDto.FieldLength = -99

	txtLabelDto.FieldJustify = TxtJustify.None()

	txtLabelDto.RightMarginStr = ""

	txtLabelDto.LineTerminator = ""

	txtLabelDto.MaxLineLength = -99

	txtLabelDto.TurnAutoLineLengthBreaksOn = false

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
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtLabelDto        TextFieldLabelDto
//     - An incoming instance of TextFieldLabelDto. This method
//       will compare all member variable data values in this
//       instance against those contained in the current instance
//       of TextFieldLabelDto. If the data values in both
//       instances are found to be equal in all respects, this
//       method will return a boolean value of 'true'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the member variable data values contained in input
//       parameter 'incomingTxtLabelDto' are equal in all respects
//       to those contained in the current instance of
//       TextFieldLabelDto, this method will return a boolean
//       value of 'true'. Otherwise a value of 'false' will be
//       returned to the calling function.
//
func (txtLabelDto *TextFieldLabelDto) Equal(
	incomingTxtLabelDto TextFieldLabelDto) bool {

	if txtLabelDto.lock == nil {
		txtLabelDto.lock = new(sync.Mutex)
	}

	txtLabelDto.lock.Lock()

	defer txtLabelDto.lock.Unlock()

	if txtLabelDto.FormatType !=
		incomingTxtLabelDto.FormatType {

		return false
	}

	if txtLabelDto.LeftMarginStr !=
		incomingTxtLabelDto.LeftMarginStr {

		return false
	}

	if txtLabelDto.FieldText !=
		incomingTxtLabelDto.FieldText {

		return false
	}

	if txtLabelDto.FieldLength !=
		incomingTxtLabelDto.FieldLength {

		return false
	}

	if txtLabelDto.FieldJustify !=
		incomingTxtLabelDto.FieldJustify {

		return false
	}

	if txtLabelDto.RightMarginStr !=
		incomingTxtLabelDto.RightMarginStr {

		return false
	}

	if txtLabelDto.LineTerminator !=
		incomingTxtLabelDto.LineTerminator {

		return false
	}

	if txtLabelDto.MaxLineLength !=
		incomingTxtLabelDto.MaxLineLength {

		return false
	}

	if txtLabelDto.TurnAutoLineLengthBreaksOn !=
		incomingTxtLabelDto.TurnAutoLineLengthBreaksOn {

		return false
	}

	return true
}

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

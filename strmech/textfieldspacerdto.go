package strmech

import "sync"

// TextFieldSpacerDto - This type is used to transmit parameters
// required for the creation of Text Spacer Fields. Text Spacer
// Fields consist of one or more white space characters (" ").
//
type TextFieldSpacerDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed.
	//
	// For the TextFieldSpacerDto Format Type, this value
	// should always be set to:
	//   TxtFieldType.Spacer()       - Valid

	LeftMarginStr string
	// The contents of the string will be used as the left margin
	// for the Text Field.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	FieldLength int
	//  An integer value greater than zero and less than 1,000,001
	//  which is used to specify the number of white space
	//  characters in the Text Spacer Field.
	//
	//  Examples:
	//     fieldLen = 1 produces text field " "
	//     fieldLen = 2 produces text field "  "
	//     fieldLen = 5 produces text field "     "

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

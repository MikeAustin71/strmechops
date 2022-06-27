package strmech

import (
	"sync"
	"time"
)

// TextFieldDto - The Text Field Data Transfer Object specifies the
// parameters necessary to construct a formatted line of text
// consisting of a left margin, a text field, a right margin and a
// line termination sequence.
//
// Example Layout
//
//         " " + "Inflation Rates" + " " + "\n"
//
type TextFieldDto struct {
	LeftMarginStr string
	// The contents of the string will be used as the left margin
	// for the Text Field.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	FieldDateTime time.Time
	// If this Text Field is designated as a 'DateTime' Field, this
	// time value will be used to populate the Text Field.

	FieldDateTimeFormat string
	// If this Text Field is designated as a 'DateTime' Field, this
	// string will be used to format the Date/Time.
	//
	// If 'FieldDateTime' is set to a value greater than zero and
	// this 'ParamValueDateTimeFormat' string is empty (has a zero length), a
	// default Date/Time format string will be applied as follows:
	//         "2006-01-02 15:04:05.000000000 -0700 MST"

	FieldText string
	// The Text Field string or contents. If this string is empty
	// (has a zero (0) length) and is designated as a Label, Filler
	// or Spacer Text Field, an error will be generated.
	//
	// This string represents the contents of the Text Field.

	FieldLength int
	// Used to format Label Text Fields. This is the length of the
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
	//
	// If TextFieldType is set to 'Filler' or 'Spacer',
	// 'FieldLength' will be used to specify the number of Filler
	// or Spacer characters in the line.
	//
	// If TextFieldType is set to 'BlankLine', 'FieldLength' will
	// be used to specify the number of blank lines created.

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

	FieldType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Field Specification which will to configure the final text
	// field output. If this parameter is invalid, an error will be
	// generated.
	//
	// Possible values are listed as follows:
	//   TxtFieldType.None()      - Invalid
	//   TxtFieldType.Label()     - Valid
	//   TxtFieldType.DateTime()  - Valid
	//   TxtFieldType.Filler()    - Valid
	//   TxtFieldType.Spacer()    - Valid
	//   TxtFieldType.BlankLine() - Valid

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

	lock *sync.Mutex
}

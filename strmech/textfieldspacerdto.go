package strmech

// TextFieldSpacerDto - This type is used to transmit parameters
// required for the creation of Text Spacer Fields. Text Spacer
// Fields consist of one or more white space characters (" ").
//
type TextFieldSpacerDto struct {
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

}

package strmech

import "sync"

type TextLineSolidLineDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed.
	//
	// Possible values are listed as follows:
	//   TxtFieldType.None()        - Invalid
	//   TxtFieldType.Label()       - Valid
	//   TxtFieldType.DateTime()    - Valid
	//   TxtFieldType.Filler()      - Valid
	//   TxtFieldType.Spacer()      - Valid
	//   TxtFieldType.BlankLine()   - Valid
	//   TxtFieldType.SolidLine()   - Valid
	//   TxtFieldType.Line1Column() - Valid
	//   TxtFieldType.Line2Column() - Valid
	//   TxtFieldType.Line3Column() - Valid
	//   TxtFieldType.Line4Column() - Valid
	//   TxtFieldType.Line5Column() - Valid
	//   TxtFieldType.Line6Column() - Valid
	//
	// For TextLineSolidDto 'FormatType' should be
	// set to: TxtFieldType.SolidLine()

	LeftMarginStr string
	// A string containing the text characters to be positioned
	// on the Left side of the Solid Line.
	//
	// If no Left margin is required, set this parameter to an
	// empty string.
	//
	// Example:
	//   solidLineChars = "*"
	//   solidLineCharsRepeatCount = 5
	//   LeftMarginStr  = "   " // 3-spaces
	//   rightMarginStr = "" // Empty string
	//   Solid line = "   *****"
	//
	// If the 'LeftMarginStr' string length is greater than
	// one-million (1,000,000), an error will be returned.

	SolidLineChars string
	// This string specifies the character or characters which
	// will comprise the solid line output for screen display or
	// printing.
	//
	// Example:
	//   solidLineChars = "*"
	//   solidLineCharsRepeatCount = 5
	//   Solid line = "*****"
	//
	// If this parameter is submitted as a zero length string, an
	// error will be returned.

	SolidLineCharRepeatCount int
	// This integer value specifies the number of times that
	// parameter 'solidLineChars' will be repeated in
	// constructing the solid line.
	//
	// If this parameter is submitted with a value less than one
	// (1), an error will be returned.
	//
	// Example:
	//   solidLineChars = "*"
	//   solidLineCharsRepeatCount = 5
	//   Solid line = "*****"

	RightMarginStr string
	// A string containing the text characters to positioned on
	// the Right side of the Solid Line.
	//
	// If no Right margin is required, set this parameter to an
	// empty string.
	//
	// Example:
	//   solidLineChars = "*"
	//   solidLineCharsRepeatCount = 5
	//   leftMarginStr = "" // Empty string
	//   RightMarginStr = "   " // 3-spaces
	//   Solid line = "*****   "
	//
	// If the 'RightMarginStr' string length is greater than
	// one-million (1,000,000), an error will be returned.

	LineTerminator string
	// This string holds the character or characters which will be
	// used to terminate the formatted line of text output.
	//
	// The most common usage sets this string to a new line
	// character ("\n").
	//
	// If no Line Terminator is required, set 'LineTerminator' to
	// a zero length or empty string and no line termination
	// characters will be created.

	lock *sync.Mutex
}

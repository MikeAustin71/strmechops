package strmech

import "sync"

type TextFieldFillerDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed.
	//
	// For the TextFieldFillerDto Format Type, this value
	// should always be set to:
	//   TxtFieldType.Filler()       - Valid

	LeftMarginStr string
	// The contents of the string will be used as the left margin
	// for the Text Field.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	FillerCharacters string
	// A string containing the text characters which will be
	// included in the Text Filler Field. The final Text Filler
	// Field will be constructed from the filler characters
	// repeated one or more times as specified by the
	// 'fillerCharsRepeatCount' parameter.
	//
	//  Text Field Filler Length =
	//    Length of fillerCharacters X fillerCharsRepeatCount
	//
	//    Example #1: fillerCharacters = "-*"
	//                fillerRepeatCount = 3
	//                Final Text Filler Field = "-*-*-*"
	//
	//    Example #2: fillerCharacters = "-"
	//                fillerRepeatCount = 3
	//                Final Text Filler Field = "---"
	//

	FillerCharsRepeatCount int
	// Controls the number of times 'fillerCharacters' is
	// repeated when constructing the final Text Filler Field
	// returned by this method. The actual length of the string
	// which will populated the completed Text Filler Field is
	// equal to the length of 'fillerCharacters' times the value
	// of 'fillerCharsRepeatCount'.
	//
	//  Text Field Filler Length =
	//    Length of fillerCharacters X fillerCharsRepeatCount
	//
	//    Example #1: fillerCharacters = "-*"
	//                fillerRepeatCount = 3
	//                Final Text Filler Field = "-*-*-*"
	//
	//    Example #2: fillerCharacters = "-"
	//                fillerRepeatCount = 3
	//                Final Text Filler Field = "---"
	//
	// If 'fillerCharsRepeatCount' has a value less than one (1) or
	// greater than one-million (1,000,000), an error will be
	// returned.

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
// TextFieldFillerDto ('incomingTxtFillerDto') to the data fields
// of the current TextFieldFillerDto instance
// ('txtFillerDto').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextFieldFillerDto instance
// ('txtFillerDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingTxtFillerDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtFillerDto        TextFieldFillerDto
//     - An instance of TextFieldFillerDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this TextFieldFillerDto instance
//       ('incomingTxtFillerDto') will be copied to the current
//       TextFieldFillerDto instance ('txtFillerDto').
//
//       No data validation is performed on input parameter,
//       'incomingTxtFillerDto'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (txtFillerDto *TextFieldFillerDto) CopyIn(
	incomingTxtFillerDto TextFieldFillerDto) {

	if txtFillerDto.lock == nil {
		txtFillerDto.lock = new(sync.Mutex)
	}

	txtFillerDto.lock.Lock()

	defer txtFillerDto.lock.Unlock()

	txtFillerDto.FormatType =
		incomingTxtFillerDto.FormatType

	txtFillerDto.LeftMarginStr =
		incomingTxtFillerDto.LeftMarginStr

	txtFillerDto.FillerCharacters =
		incomingTxtFillerDto.FillerCharacters

	txtFillerDto.FillerCharsRepeatCount =
		incomingTxtFillerDto.FillerCharsRepeatCount

	txtFillerDto.RightMarginStr =
		incomingTxtFillerDto.RightMarginStr

	txtFillerDto.LineTerminator =
		incomingTxtFillerDto.LineTerminator

	txtFillerDto.MaxLineLength =
		incomingTxtFillerDto.MaxLineLength

	txtFillerDto.TurnAutoLineLengthBreaksOn =
		incomingTxtFillerDto.TurnAutoLineLengthBreaksOn

	return
}

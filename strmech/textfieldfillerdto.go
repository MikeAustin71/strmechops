package strmech

import "sync"

// TextFieldFillerDto - The Text Field Filler Data Transfer Object
// is designed to transfer data parameters and specifications
// required to construct a Text Filler Field.
//
// Typically, Text Filler Fields are designed to be configured
// within a line of text. However, with the TextFieldFillerDto,
// users have the option of configuring a Text Filler Field as a
// separate stand-alone line of text.
//
// Text Filler Fields are commonly used as margins containing
// multiple white space characters, or line separators containing
// multiple dashes, equal signs or underscore characters. Text
// Filler Fields consist of filler characters ('FillerCharacters')
// and the filler characters repeat count
// ('FillerCharsRepeatCount'). A filler field is made up of one or
// more filler characters. These filler characters are repeated one
// or more times in order to construct the complete filler field as
// shown in the following examples:
//
//  Example 1:
//   Filler Characters = "-"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "---"
//
//  Example 2:
//   Filler Characters = "-*"
//   Filler Characters Repeat Count = 3
//   Formatted Text = "-*-*-*"
//
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
	// 'FillerCharsRepeatCount' parameter.
	//
	//  Text Field Filler Length =
	//    Length of fillerCharacters X fillerCharsRepeatCount
	//
	//    Example #1: FillerCharacters = "-*"
	//                FillerCharsRepeatCount = 3
	//                Final Text Filler Field = "-*-*-*"
	//
	//    Example #2: FillerCharacters = "-"
	//                FillerCharsRepeatCount = 3
	//                Final Text Filler Field = "---"
	//

	FillerCharsRepeatCount int
	// Controls the number of times 'FillerCharacters' is
	// repeated when constructing the final Text Filler Field.
	// The actual length of the string which will populate the
	// completed Text Filler Field is equal to the length of
	// 'FillerCharacters' times the value of
	// 'FillerCharsRepeatCount'.
	//
	//  Text Field Filler Length =
	//    Length of FillerCharacters X FillerCharsRepeatCount
	//
	//    Example #1: FillerCharacters = "-*"
	//                FillerRepeatCount = 3
	//                Final Text Filler Field = "-*-*-*"
	//
	//    Example #2: FillerCharacters = "-"
	//                FillerRepeatCount = 3
	//                Final Text Filler Field = "---"
	//
	// If 'FillerCharsRepeatCount' has a value less than one (1) or
	// greater than one-million (1,000,000), an error will be
	// returned.

	RightMarginStr string
	// The contents of the string will be used as the right margin
	// for the Text Filler Field.
	//
	// If no right margin is required, set 'RightMarginStr' to a
	// zero length or empty string, and no right margin will be
	// created.

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

// CopyOut - Returns a deep copy of the current TextFieldFillerDto
// instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFieldFillerDto.
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
//  deepCopyTxtFillerDto        TextFieldFillerDto
//     - This parameter will return a deep copy of the current
//       TextFieldFillerDto instance.
//
func (txtFillerDto *TextFieldFillerDto) CopyOut() (
	deepCopyTxtFillerDto TextFieldFillerDto) {

	if txtFillerDto.lock == nil {
		txtFillerDto.lock = new(sync.Mutex)
	}

	txtFillerDto.lock.Lock()

	defer txtFillerDto.lock.Unlock()

	deepCopyTxtFillerDto.FormatType =
		txtFillerDto.FormatType

	deepCopyTxtFillerDto.LeftMarginStr =
		txtFillerDto.LeftMarginStr

	deepCopyTxtFillerDto.FillerCharacters =
		txtFillerDto.FillerCharacters

	deepCopyTxtFillerDto.FillerCharsRepeatCount =
		txtFillerDto.FillerCharsRepeatCount

	deepCopyTxtFillerDto.RightMarginStr =
		txtFillerDto.RightMarginStr

	deepCopyTxtFillerDto.LineTerminator =
		txtFillerDto.LineTerminator

	deepCopyTxtFillerDto.MaxLineLength =
		txtFillerDto.MaxLineLength

	deepCopyTxtFillerDto.TurnAutoLineLengthBreaksOn =
		txtFillerDto.TurnAutoLineLengthBreaksOn

	return deepCopyTxtFillerDto
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldFillerDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextFieldFillerDto in an invalid state and unavailable for
// immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFieldFillerDto. All member variable data
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
func (txtFillerDto *TextFieldFillerDto) Empty() {

	if txtFillerDto.lock == nil {
		txtFillerDto.lock = new(sync.Mutex)
	}

	txtFillerDto.lock.Lock()

	txtFillerDto.FormatType = TxtFieldType.None()

	txtFillerDto.LeftMarginStr = ""

	txtFillerDto.FillerCharacters = ""

	txtFillerDto.FillerCharsRepeatCount = 0

	txtFillerDto.RightMarginStr = ""

	txtFillerDto.LineTerminator = ""

	txtFillerDto.MaxLineLength = -99

	txtFillerDto.TurnAutoLineLengthBreaksOn = false

	txtFillerDto.lock.Unlock()

	txtFillerDto.lock = nil

}

// Equal - Receives another instance of TextFieldFillerDto and
// proceeds to compare the member variables to those of the current
// TextFieldFillerDto instance in order to determine if they are
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
//  incomingTxtFillerDto        TextFieldFillerDto
//     - An incoming instance of TextFieldFillerDto. This method
//       will compare all member variable data values in this
//       instance against those contained in the current instance
//       of TextFieldFillerDto. If the data values in both
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
//       parameter 'incomingTxtFillerDto' are equal in all respects
//       to those contained in the current instance of
//       TextFieldFillerDto, this method will return a boolean
//       value of 'true'. Otherwise a value of 'false' will be
//       returned to the calling function.
//
func (txtFillerDto *TextFieldFillerDto) Equal(
	incomingTxtFillerDto TextFieldFillerDto) bool {

	if txtFillerDto.lock == nil {
		txtFillerDto.lock = new(sync.Mutex)
	}

	txtFillerDto.lock.Lock()

	defer txtFillerDto.lock.Unlock()

	if txtFillerDto.FormatType !=
		incomingTxtFillerDto.FormatType {

		return false
	}

	if txtFillerDto.LeftMarginStr !=
		incomingTxtFillerDto.LeftMarginStr {

		return false
	}

	if txtFillerDto.FillerCharacters !=
		incomingTxtFillerDto.FillerCharacters {

		return false
	}

	if txtFillerDto.FillerCharsRepeatCount !=
		incomingTxtFillerDto.FillerCharsRepeatCount {

		return false
	}

	if txtFillerDto.RightMarginStr !=
		incomingTxtFillerDto.RightMarginStr {

		return false
	}

	if txtFillerDto.LineTerminator !=
		incomingTxtFillerDto.LineTerminator {

		return false
	}

	if txtFillerDto.MaxLineLength !=
		incomingTxtFillerDto.MaxLineLength {

		return false
	}

	if txtFillerDto.TurnAutoLineLengthBreaksOn !=
		incomingTxtFillerDto.TurnAutoLineLengthBreaksOn {

		return false
	}

	return true
}

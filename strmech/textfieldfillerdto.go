package strmech

type TextFieldFillerDto struct {
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

}

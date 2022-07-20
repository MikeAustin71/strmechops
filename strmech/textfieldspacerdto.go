package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextFieldSpacerDto - This type is used to transmit data
// parameters and specifications required to create Text Spacer
// Fields.
//
// Text Spacer Fields consist of one or more white space
// characters (" ").
//
// Text Fields are typically configured to reside within a line
// of text. As such, they are NOT generally terminated with a new
// line character ('\n') thereby creating a separate stand-alone
// line of text. However, users have the option of specifying
// line termination characters ('LineTerminator') which WILL
// configure Text Spacer Fields as separator stand-alone lines
// of text.
//
// Text Spacer Field Example-1:
//
//  LeftMarginStr = ""
//  FieldLength = 3
//  RightMarginStr = ""
//  LineTerminator = ""
//  Final Text Spacer string = "   " // 3-white spaces
//
// Text Spacer Field Example-2:
//
//  LeftMarginStr = ""
//  FieldLength = 3
//  RightMarginStr = ""
//  LineTerminator = "\n"
//  Final Text Spacer string = "   \n" // 3-white spaces and
//                                     // 1-new line character
//                                     // ('\n')
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
	// for the Text Spacer Field.
	//
	// If no left margin is required, set 'LeftMarginStr' to a zero
	// length or empty string, and no left margin will be created.

	FieldLength int
	//  An integer value used to specify the number of white space
	//  characters in the Text Spacer Field.
	//
	//  Values less than zero and greater than one million
	//  (1,000,001) will generate errors.
	//
	//  Examples:
	//   fieldLen = 1 produces text field " "  // 1-white space
	//   fieldLen = 2 produces text field "  " // 2-white spaces
	//   fieldLen = 5 produces text field "     " // 5-white spaces

	RightMarginStr string
	// The contents of the string will be used as the right margin
	// for the Text Spacer Field.
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
	// If no Line Terminator is required, set 'lineTerminator'
	// to a zero length or empty string and no line termination
	// characters will be created.

	MaxLineLength int
	// The maximum length of the line on which this label Text
	// Field Spacer string will be presented.
	//
	// Set this parameter to minus one (-1) to specify an
	// unlimited line length for this text line.
	//
	// 'MaxLineLength' is used in conjunction with parameter
	// 'TurnAutoLineLengthBreaksOn' to automatically place text
	// fields on separate text lines when that text exceeds the
	// maximum text line length ('MaxLineLength'). Therefore,
	// paramter 'TurnAutoLineLengthBreaksOn' controls whether
	// automatic line breaks using 'MaxLineLength' will be
	// applied.
	//
	// If the value of 'MaxLineLength' is less than one (1), it
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
	//
	// When this parameter is set to 'true', text fields which extend
	// beyond the maximum line length 'MaxLineLength' will be
	// formatted as a separate line of text on the following line.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldSpacerDto ('incomingTxtSpacerDto') to the data fields
// of the current TextFieldSpacerDto instance
// ('txtSpacerDto').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextFieldSpacerDto instance
// ('txtSpacerDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingTxtSpacerDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtSpacerDto        TextFieldSpacerDto
//     - An instance of TextFieldSpacerDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this TextFieldSpacerDto instance
//       ('incomingTxtSpacerDto') will be copied to the current
//       TextFieldSpacerDto instance ('txtSpacerDto').
//
//       No data validation is performed on input parameter,
//       'incomingTxtSpacerDto'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (txtSpacerDto *TextFieldSpacerDto) CopyIn(
	incomingTxtSpacerDto TextFieldSpacerDto) {

	if txtSpacerDto.lock == nil {
		txtSpacerDto.lock = new(sync.Mutex)
	}

	txtSpacerDto.lock.Lock()

	defer txtSpacerDto.lock.Unlock()

	_ = textFieldSpacerDtoNanobot{}.ptr().copy(
		txtSpacerDto,
		&incomingTxtSpacerDto,
		nil)

	return
}

// CopyOut - Returns a deep copy of the current TextFieldSpacerDto
// instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFieldSpacerDto.
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
//  deepCopyTxtSpacerDto        TextFieldSpacerDto
//     - This parameter will return a deep copy of the current
//       TextFieldSpacerDto instance.
//
func (txtSpacerDto *TextFieldSpacerDto) CopyOut() (
	deepCopyTxtSpacerDto TextFieldSpacerDto) {

	if txtSpacerDto.lock == nil {
		txtSpacerDto.lock = new(sync.Mutex)
	}

	txtSpacerDto.lock.Lock()

	defer txtSpacerDto.lock.Unlock()

	_ = textFieldSpacerDtoNanobot{}.ptr().copy(
		&deepCopyTxtSpacerDto,
		txtSpacerDto,
		nil)

	return deepCopyTxtSpacerDto
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldSpacerDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextFieldSpacerDto in an invalid state and unavailable for
// immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFieldSpacerDto. All member variable data
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
func (txtSpacerDto *TextFieldSpacerDto) Empty() {

	if txtSpacerDto.lock == nil {
		txtSpacerDto.lock = new(sync.Mutex)
	}

	txtSpacerDto.lock.Lock()

	txtSpacerDto.FormatType = TxtFieldType.None()

	txtSpacerDto.LeftMarginStr = ""

	txtSpacerDto.FieldLength = -99

	txtSpacerDto.RightMarginStr = ""

	txtSpacerDto.LineTerminator = ""

	txtSpacerDto.MaxLineLength = -99

	txtSpacerDto.TurnAutoLineLengthBreaksOn = false

	txtSpacerDto.lock.Unlock()

	txtSpacerDto.lock = nil

	return
}

// Equal - Receives another instance of TextFieldSpacerDto and
// proceeds to compare the member variables to those of the current
// TextFieldSpacerDto instance in order to determine if they are
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
//  incomingTxtSpacerDto       TextFieldSpacerDto
//     - An incoming instance of TextFieldSpacerDto. This method
//       will compare all member variable data values in this
//       instance against those contained in the current instance
//       of TextFieldSpacerDto. If the data values in both
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
//       parameter 'incomingTxtSpacerDto' are equal in all respects
//       to those contained in the current instance of
//       TextFieldSpacerDto, this method will return a boolean
//       value of 'true'. Otherwise a value of 'false' will be
//       returned to the calling function.
//
func (txtSpacerDto *TextFieldSpacerDto) Equal(
	incomingTxtSpacerDto TextFieldSpacerDto) bool {

	if txtSpacerDto.lock == nil {
		txtSpacerDto.lock = new(sync.Mutex)
	}

	txtSpacerDto.lock.Lock()

	defer txtSpacerDto.lock.Unlock()

	if txtSpacerDto.FormatType !=
		incomingTxtSpacerDto.FormatType {

		return false
	}

	if txtSpacerDto.LeftMarginStr !=
		incomingTxtSpacerDto.LeftMarginStr {

		return false
	}

	if txtSpacerDto.FieldLength !=
		incomingTxtSpacerDto.FieldLength {

		return false
	}

	if txtSpacerDto.RightMarginStr !=
		incomingTxtSpacerDto.RightMarginStr {

		return false
	}

	if txtSpacerDto.LineTerminator !=
		incomingTxtSpacerDto.LineTerminator {

		return false
	}

	if txtSpacerDto.MaxLineLength !=
		incomingTxtSpacerDto.MaxLineLength {

		return false
	}

	if txtSpacerDto.TurnAutoLineLengthBreaksOn !=
		incomingTxtSpacerDto.TurnAutoLineLengthBreaksOn {

		return false
	}

	return true
}

// textFieldSpacerDtoNanobot - Provides helper methods for
// TextFieldSpacerDto.
type textFieldSpacerDtoNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFieldSpacerDto to a destination instance of
// TextFieldSpacerDto.
func (txtSpacerDtoNanobot *textFieldSpacerDtoNanobot) copy(
	destinationTxtSpacerDto *TextFieldSpacerDto,
	sourceTxtSpacerDto *TextFieldSpacerDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtSpacerDtoNanobot.lock == nil {
		txtSpacerDtoNanobot.lock = new(sync.Mutex)
	}

	txtSpacerDtoNanobot.lock.Lock()

	defer txtSpacerDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpacerDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtSpacerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtSpacerDto' is a nil pointer!\n",
			ePrefix.String())

	}

	if destinationTxtSpacerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtSpacerDto' is a nil pointer!\n",
			ePrefix.String())

	}

	destinationTxtSpacerDto.FormatType =
		sourceTxtSpacerDto.FormatType

	destinationTxtSpacerDto.LeftMarginStr =
		sourceTxtSpacerDto.LeftMarginStr

	destinationTxtSpacerDto.FieldLength =
		sourceTxtSpacerDto.FieldLength

	destinationTxtSpacerDto.RightMarginStr =
		sourceTxtSpacerDto.RightMarginStr

	destinationTxtSpacerDto.LineTerminator =
		sourceTxtSpacerDto.LineTerminator

	destinationTxtSpacerDto.MaxLineLength =
		sourceTxtSpacerDto.MaxLineLength

	destinationTxtSpacerDto.TurnAutoLineLengthBreaksOn =
		sourceTxtSpacerDto.TurnAutoLineLengthBreaksOn

	return err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpacerDtoNanobot.
//
func (txtSpacerDtoNanobot textFieldSpacerDtoNanobot) ptr() *textFieldSpacerDtoNanobot {

	if txtSpacerDtoNanobot.lock == nil {
		txtSpacerDtoNanobot.lock = new(sync.Mutex)
	}

	txtSpacerDtoNanobot.lock.Lock()

	defer txtSpacerDtoNanobot.lock.Unlock()

	return &textFieldSpacerDtoNanobot{
		lock: new(sync.Mutex),
	}
}

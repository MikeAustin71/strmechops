package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextLineSolidDto - The Text Line Solid Data Transfer Object
// is used to transfer data parameters required for the creation
// of a solid line of text characters formatted for screen display,
// file output or printing.
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// by member variable, 'SolidLineCharRepeatCount'.
//
// Therefore, the number of solid line characters produced is
// equal to:
//
//	SolidLineChars x SolidLineCharRepeatCount
//
// By default, all solid lines are terminated by a new line
// character ('\n'). This means that each instance of a solid will
// be formatted as a single line of text. Users have the option
// to 'turn-off' this behavior and control the content and
// application of line termination characters through two member
// variables: 'TurnLineTerminationOff' and 'LineTerminator'.
//
//	Example-1:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = false
//	  LineTerminator = ""
//	  Final Solid Line String = "*****\n"
//
//	Example-2:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = false
//	  LineTerminator = "\n\n"
//	  Final Solid Line String = "*****\n\n"
//
//	Example-3:
//	  SolidLineChars = "*"
//	  SolidLineCharRepeatCount = 5
//	  LeftMarginStr = ""
//	  RightMarginStr = ""
//	  TurnLineTerminationOff = true
//	  LineTerminator = "\n\n"
//	  Final Solid Line String = "*****"
type TextLineSolidDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of
	// Text Format Operation to be performed.
	//
	// For TextLineSolidDto 'FormatType' this parameter should
	// be set to: TxtFieldType.SolidLine()

	LeftMarginStr string
	// A string containing the text characters to be positioned
	// on the Left side of the Solid Line.
	//
	// If no Left margin is required, set this parameter to an
	// empty string.
	//
	// Example:
	//   LeftMarginStr = "xxx"
	//   SolidLineChars = "*"
	//   SolidLineCharRepeatCount = 5
	//   RightMarginStr = "" // Empty string
	//   Solid Line = "xxx*****"
	//
	// If the 'LeftMarginStr' string length is greater than
	// one-million (1,000,000), an error will be returned.

	SolidLineChars string
	// This string specifies the character or characters which
	// will comprise the solid line string for screen display,
	// file output or printing.
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
	//   RightMarginStr = "xxx"
	//   Solid line = "*****xxx"
	//
	// If the 'RightMarginStr' string length is greater than
	// one-million (1,000,000), an error will be returned.

	TurnLineTerminationOff bool
	// If this parameter is set to 'true' no Line Termination
	// Sequence will be applied for this text line (Example-3).
	//
	//       Example-1:
	//         SolidLineChars = "*"
	//         SolidLineCharRepeatCount = 5
	//         LeftMarginStr = ""
	//         RightMarginStr = ""
	//         TurnLineTerminationOff = false
	//         LineTerminator = ""
	//         Final Solid Line String = "*****\n"
	//
	//       Example-2:
	//         SolidLineChars = "*"
	//         SolidLineCharRepeatCount = 5
	//         LeftMarginStr = ""
	//         RightMarginStr = ""
	//         TurnLineTerminationOff = false
	//         LineTerminator = "\n\n"
	//         Final Solid Line String = "*****\n\n"
	//
	//       Example-3:
	//         SolidLineChars = "*"
	//         SolidLineCharRepeatCount = 5
	//         LeftMarginStr = ""
	//         RightMarginStr = ""
	//         TurnLineTerminationOff = true
	//         LineTerminator = "\n\n"
	//         Final Solid Line String = "*****"

	LineTerminator string
	// This string holds the character or characters which
	// will be used to terminate the formatted line of text
	// output.
	//
	// The most common usage sets this string to a new line
	// character ("\n").
	//
	// If 'LineTerminator' is configured as an empty string
	// (string length zero), a single new line character ('\n')
	// will be automatically applied to produce line termination,
	// depending on the setting for parameter
	// 'TurnLineTerminationOff'.
	//
	// LineTerminator works in conjunction with member variable
	// 'TurnLineTerminationOff'. 'TurnLineTerminationOff'
	// controls the application of a line terminator. Setting
	// 'TurnLineTerminationOff' to 'true' means that NO line
	// terminator will be applied to this instance of Text Solid
	// Line.
	//
	//       Example-1:
	//         SolidLineChars = "*"
	//         SolidLineCharRepeatCount = 5
	//         LeftMarginStr = ""
	//         RightMarginStr = ""
	//         TurnLineTerminationOff = false
	//         LineTerminator = ""
	//         Final Solid Line String = "*****\n"
	//
	//       Example-2:
	//         SolidLineChars = "*"
	//         SolidLineCharRepeatCount = 5
	//         LeftMarginStr = ""
	//         RightMarginStr = ""
	//         TurnLineTerminationOff = false
	//         LineTerminator = "\n\n"
	//         Final Solid Line String = "*****\n\n"
	//
	//       Example-3:
	//         SolidLineChars = "*"
	//         SolidLineCharRepeatCount = 5
	//         LeftMarginStr = ""
	//         RightMarginStr = ""
	//         TurnLineTerminationOff = true
	//         LineTerminator = "\n\n"
	//         Final Solid Line String = "*****"

	MaxLineLength int
	// The maximum length of the line on which the solid line
	// text characters will be presented.
	//
	// Set this parameter to minus one (-1) to specify an
	// unlimited line length for this text line.
	//
	// If the value of 'MaxLineLength' is less than one (1),
	// it will be automatically converted to minus one (-1).
	//
	// 'MaxLineLength' is used in conjunction with parameter
	// 'TurnAutoLineLengthBreaksOn' to automatically place text
	// on separate text lines when that text exceeds the maximum
	// text line length ('MaxLineLength'). Therefore, paramter
	// 'TurnAutoLineLengthBreaksOn' controls whether automatic
	// line breaks using 'MaxLineLength' will be applied.

	TurnAutoLineLengthBreaksOn bool
	// This parameter controls whether text lines which exceed
	// the maximum line length ('MaxLineLength') are
	// positioned on the following line as a separate line of
	// text.
	//
	// To apply automatic line breaking at the maximum line
	// length, set the value of this parameter to 'true'.
	//
	// When this parameter is set to 'true', text fields which
	// extend beyond the maximum line length 'MaxLineLength'
	// will be formatted as a separate line of text on the
	// following line.

	MultiLineLeftMarginStr string
	// The left margin used when a text string exceeds the
	// maximum line length and is separated into multiple text
	// lines. This left margin is applied to the second and
	// all subsequent lines of a multi-line text display. This
	// parameter is only valid when 'TurnAutoLineLengthBreaksOn'
	// is set to 'true' and the initial text string exceeds the
	// maximum line length.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextLineSolidDto ('incomingSolidLineDto') to the data fields
// of the current TextLineSolidDto instance
// ('txtSolidLineDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextLineSolidDto instance
// ('txtSolidLineDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingSolidLineDto'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingSolidLineDto        TextLineSolidDto
//	   - An instance of TextLineSolidDto. This method will NOT
//	     change the data values of member variables contained in
//	     this instance.
//
//	     All data values in this TextLineSolidDto instance
//	     ('incomingSolidLineDto') will be copied to the current
//	     TextLineSolidDto instance ('txtSolidLineDto').
//
//	     No data validation is performed on input parameter,
//	     'incomingSolidLineDto'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtSolidLineDto *TextLineSolidDto) CopyIn(
	incomingSolidLineDto TextLineSolidDto) {

	if txtSolidLineDto.lock == nil {
		txtSolidLineDto.lock = new(sync.Mutex)
	}

	txtSolidLineDto.lock.Lock()

	defer txtSolidLineDto.lock.Unlock()

	_ = textLineSolidLineDtoNanobot{}.ptr().copyData(
		txtSolidLineDto,
		&incomingSolidLineDto,
		nil)

	return
}

// CopyOut - Returns a deep copy of the current TextLineSolidDto
// instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextLineSolidDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyTxtLineSolidDto    TextLineSolidDto
//	   - This parameter will return a deep copy of the current
//	     TextLineSolidDto instance.
func (txtSolidLineDto *TextLineSolidDto) CopyOut() (
	deepCopyTxtLineSolidDto TextLineSolidDto) {

	if txtSolidLineDto.lock == nil {
		txtSolidLineDto.lock = new(sync.Mutex)
	}

	txtSolidLineDto.lock.Lock()

	defer txtSolidLineDto.lock.Unlock()

	_ = textLineSolidLineDtoNanobot{}.ptr().copyData(
		&deepCopyTxtLineSolidDto,
		txtSolidLineDto,
		nil)

	return deepCopyTxtLineSolidDto
}

// Empty - Resets all internal member variables for the current
// instance of TextLineSolidDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextLineSolidDto in an invalid state and unavailable for
// immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextLineSolidDto. All member variable data
// values will be reset to their zero or uninitialized states.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtSolidLineDto *TextLineSolidDto) Empty() {

	if txtSolidLineDto.lock == nil {
		txtSolidLineDto.lock = new(sync.Mutex)
	}

	txtSolidLineDto.lock.Lock()

	textLineSolidLineDtoMolecule{}.ptr().empty(
		txtSolidLineDto)

	txtSolidLineDto.lock.Unlock()

	txtSolidLineDto.lock = nil

	return
}

// Equal - Receives another instance of TextLineSolidDto and
// proceeds to compare the member variables to those of the current
// TextLineSolidDto instance in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTxtLineSolidDto       TextLineSolidDto
//	   - An incoming instance of TextLineSolidDto. This method
//	     will compare all member variable data values in this
//	     instance against those contained in the current instance
//	     of TextLineSolidDto. If the data values in both
//	     instances are found to be equal in all respects, this
//	     method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingTxtLineSolidDto' are equal in all
//	     respects to those contained in the current instance of
//	     TextLineSolidDto, this method will return a boolean
//	     value of 'true'. Otherwise, a value of 'false' will be
//	     returned to the calling function.
func (txtSolidLineDto *TextLineSolidDto) Equal(
	incomingTxtLineSolidDto TextLineSolidDto) bool {

	if txtSolidLineDto.lock == nil {
		txtSolidLineDto.lock = new(sync.Mutex)
	}

	txtSolidLineDto.lock.Lock()

	defer txtSolidLineDto.lock.Unlock()

	return textLineSolidLineDtoMolecule{}.ptr().equal(
		txtSolidLineDto,
		&incomingTxtLineSolidDto)
}

// textLineSolidLineDtoNanobot - Provides helper methods for
// TextLineSolidDto.
type textLineSolidLineDtoNanobot struct {
	lock *sync.Mutex
}

// copyData - Copies all data from a source instance of
// TextLineSolidDto to a destination instance of
// TextLineSolidDto.
func (txtSolidLineDtoNanobot *textLineSolidLineDtoNanobot) copyData(
	destinationSolidLineDto *TextLineSolidDto,
	sourceSolidLineDto *TextLineSolidDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtSolidLineDtoNanobot.lock == nil {
		txtSolidLineDtoNanobot.lock = new(sync.Mutex)
	}

	txtSolidLineDtoNanobot.lock.Lock()

	defer txtSolidLineDtoNanobot.lock.Unlock()

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

	if sourceSolidLineDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceSolidLineDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationSolidLineDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationSolidLineDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	textLineSolidLineDtoMolecule{}.ptr().empty(
		destinationSolidLineDto)

	destinationSolidLineDto.FormatType =
		sourceSolidLineDto.FormatType

	destinationSolidLineDto.LeftMarginStr =
		sourceSolidLineDto.LeftMarginStr

	destinationSolidLineDto.SolidLineChars =
		sourceSolidLineDto.SolidLineChars

	destinationSolidLineDto.SolidLineCharRepeatCount =
		sourceSolidLineDto.SolidLineCharRepeatCount

	destinationSolidLineDto.RightMarginStr =
		sourceSolidLineDto.RightMarginStr

	destinationSolidLineDto.TurnLineTerminationOff =
		sourceSolidLineDto.TurnLineTerminationOff

	destinationSolidLineDto.LineTerminator =
		sourceSolidLineDto.LineTerminator

	destinationSolidLineDto.MaxLineLength =
		sourceSolidLineDto.MaxLineLength

	destinationSolidLineDto.TurnAutoLineLengthBreaksOn =
		sourceSolidLineDto.TurnAutoLineLengthBreaksOn

	return err
}

// ptr - Returns a pointer to a new instance of
// textLineSolidLineDtoNanobot.
func (txtSolidLineDtoNanobot textLineSolidLineDtoNanobot) ptr() *textLineSolidLineDtoNanobot {

	if txtSolidLineDtoNanobot.lock == nil {
		txtSolidLineDtoNanobot.lock = new(sync.Mutex)
	}

	txtSolidLineDtoNanobot.lock.Lock()

	defer txtSolidLineDtoNanobot.lock.Unlock()

	return &textLineSolidLineDtoNanobot{
		lock: new(sync.Mutex),
	}
}

// textLineSolidLineDtoMolecule - Provides helper methods for
// TextLineSolidDto.
type textLineSolidLineDtoMolecule struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of TextLineSolidDto
// and proceeds to set all the internal member variables to their
// zero or uninitialized states.
//
// This method will therefore delete all data currently held
// by this instance of TextLineSolidDto.
func (txtSolidLineDtoMolecule *textLineSolidLineDtoMolecule) empty(
	solidLineDto *TextLineSolidDto) {

	if txtSolidLineDtoMolecule.lock == nil {
		txtSolidLineDtoMolecule.lock = new(sync.Mutex)
	}

	txtSolidLineDtoMolecule.lock.Lock()

	defer txtSolidLineDtoMolecule.lock.Unlock()

	solidLineDto.FormatType = TxtFieldType.None()

	solidLineDto.LeftMarginStr = ""

	solidLineDto.SolidLineChars = ""

	solidLineDto.SolidLineCharRepeatCount = 0

	solidLineDto.RightMarginStr = ""

	solidLineDto.TurnLineTerminationOff = false

	solidLineDto.LineTerminator = ""

	solidLineDto.MaxLineLength = -99

	solidLineDto.TurnAutoLineLengthBreaksOn = false

	return
}

// equal - Receives pointers to two instances of TextLineSolidDto
// and proceeds to compare all the member data variables for both
// instances.
//
// If the two instances of TextLineSolidDto are found to be equal
// in all respects, this method will return a boolean value of
// 'true'.
func (txtSolidLineDtoMolecule *textLineSolidLineDtoMolecule) equal(
	solidLineDto1 *TextLineSolidDto,
	solidLineDto2 *TextLineSolidDto) bool {

	if txtSolidLineDtoMolecule.lock == nil {
		txtSolidLineDtoMolecule.lock = new(sync.Mutex)
	}

	txtSolidLineDtoMolecule.lock.Lock()

	defer txtSolidLineDtoMolecule.lock.Unlock()

	if solidLineDto1 == nil ||
		solidLineDto2 == nil {
		return false
	}

	if solidLineDto1.FormatType !=
		solidLineDto2.FormatType {

		return false
	}

	if solidLineDto1.LeftMarginStr !=
		solidLineDto2.LeftMarginStr {

		return false
	}

	if solidLineDto1.SolidLineChars !=
		solidLineDto2.SolidLineChars {

		return false
	}

	if solidLineDto1.SolidLineCharRepeatCount !=
		solidLineDto2.SolidLineCharRepeatCount {

		return false
	}

	if solidLineDto1.RightMarginStr !=
		solidLineDto2.RightMarginStr {

		return false
	}

	if solidLineDto1.TurnLineTerminationOff !=
		solidLineDto2.TurnLineTerminationOff {

		return false
	}

	if solidLineDto1.LineTerminator !=
		solidLineDto2.LineTerminator {

		return false
	}

	if solidLineDto1.MaxLineLength !=
		solidLineDto2.MaxLineLength {

		return false
	}

	if solidLineDto1.TurnAutoLineLengthBreaksOn !=
		solidLineDto2.TurnAutoLineLengthBreaksOn {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// textFieldSpacerDtoMolecule.
func (txtSolidLineDtoMolecule textLineSolidLineDtoMolecule) ptr() *textLineSolidLineDtoMolecule {

	if txtSolidLineDtoMolecule.lock == nil {
		txtSolidLineDtoMolecule.lock = new(sync.Mutex)
	}

	txtSolidLineDtoMolecule.lock.Lock()

	defer txtSolidLineDtoMolecule.lock.Unlock()

	return &textLineSolidLineDtoMolecule{
		lock: new(sync.Mutex),
	}
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextAdHocDto - Text Ad Hoc Data Transfer Object is used to
// transfer data parameters required for to inject a user
// generated string of text into the stream of text characters
// being formatted for screen display, file output or printing.
//
// Except for line breaks configured at the user's discretion,
// no additional formatting is performed on this text, and it is
// inserted "as is" in to the final output of formatted text.
type TextAdHocDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of
	// Text Format Operation to be performed.
	//
	// For TextAdHocDto 'FormatType' this parameter should
	// be set to: TxtFieldType.TextAdHoc()

	LeftMarginStr string
	// A string containing the text characters to be positioned
	// on the Left side of the Ad Hoc Text.
	//
	// If no Left margin is required, set this parameter to an
	// empty string.

	AdHocText string
	// This strings holds the raw ad hoc text.

	RightMarginStr string
	// The contents of the string will be used as the right
	// margin for the Text Ad Hoc string.
	//
	// If no right margin is required, set 'RightMarginStr' to
	// a zero length or empty string, and no right margin will
	// be created.

	TurnLineTerminationOff bool
	// By default, a new line string terminator ('\n') will be
	// appended to the Ad Hoc text ('AdHocText'). If this
	// parameter is set to 'true', no line termination sequence
	// will be applied.
	//
	// This parameter controls the operation of parameter
	// 'LineTerminator'. If 'TurnLineTerminationOff' is set
	// to 'true', 'LineTerminator' will be completely
	// ignored and have no effect.

	LineTerminator string
	// This string holds the character or characters which
	// will be used to terminate the formatted line of text
	// output, if parameter 'TurnLineTerminationOff' is set
	// to 'false'.
	//
	// The most common usage sets this string to a new line
	// character ("\n").
	//
	// If 'LineTerminator' is configured as an empty string
	// (string length zero), a single new line character ('\n')
	// will be automatically applied to produce line termination
	// depending on the setting for parameter
	// 'TurnLineTerminationOff'.
	//
	// LineTerminator works in conjunction with member variable
	// 'TurnLineTerminationOff'. 'TurnLineTerminationOff'
	// controls the application of a line terminator. Setting
	// 'TurnLineTerminationOff' to 'true' means that NO line
	// terminator will be applied.
	//
	// Setting 'TurnLineTerminationOff' to 'true' means that
	// parameter 'LineTerminator' will be completely ignored
	// and have no effect.

	MaxLineLength int
	// The maximum length of the line on which the ad hoc
	// text characters ('AdHocText') will be presented.
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

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextAdHocDto ('incomingTxtAdHocDto') to the data fields
// of the current TextAdHocDto instance
// ('textAdHocDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextAdHocDto instance
// ('textAdHocDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingTxtAdHocDto'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTxtAdHocDto        TextAdHocDto
//	   - An instance of TextAdHocDto. This method will NOT
//	     change the data values of member variables contained in
//	     this instance.
//
//	     All data values in this TextAdHocDto instance
//	     ('incomingTxtAdHocDto') will be copied to the current
//	     TextAdHocDto instance ('textAdHocDto').
//
//	     No data validation is performed on input parameter,
//	     'incomingTxtAdHocDto'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (textAdHocDto *TextAdHocDto) CopyIn(
	incomingTxtAdHocDto TextAdHocDto) {

	if textAdHocDto.lock == nil {
		textAdHocDto.lock = new(sync.Mutex)
	}

	textAdHocDto.lock.Lock()

	defer textAdHocDto.lock.Unlock()

	_ = textAdHocDtoNanobot{}.ptr().copyData(
		textAdHocDto,
		&incomingTxtAdHocDto,
		nil)

	return
}

// CopyOut - Returns a deep copy of the current TextAdHocDto
// instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextAdHocDto.
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
//	deepCopyTxtAdHocDto        TextAdHocDto
//	   - This parameter will return a deep copy of the current
//	     TextAdHocDto instance.
func (textAdHocDto *TextAdHocDto) CopyOut() (
	deepCopyTxtAdHocDto TextAdHocDto) {

	if textAdHocDto.lock == nil {
		textAdHocDto.lock = new(sync.Mutex)
	}

	textAdHocDto.lock.Lock()

	defer textAdHocDto.lock.Unlock()

	_ = textAdHocDtoNanobot{}.ptr().copyData(
		&deepCopyTxtAdHocDto,
		textAdHocDto,
		nil)

	return deepCopyTxtAdHocDto
}

// Empty - Resets all internal member variables for the current
// instance of TextAdHocDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextAdHocDto in an invalid state and unavailable for
// immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextAdHocDto. All member variable data
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
func (textAdHocDto *TextAdHocDto) Empty() {

	if textAdHocDto.lock == nil {
		textAdHocDto.lock = new(sync.Mutex)
	}

	textAdHocDto.lock.Lock()

	textAdHocDtoMolecule{}.ptr().empty(
		textAdHocDto)

	textAdHocDto.lock.Unlock()

	textAdHocDto.lock = nil

}

// Equal - Receives another instance of TextAdHocDto and
// proceeds to compare the member variables to those of the current
// TextAdHocDto instance in order to determine if they are
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
//	incomingTxtAdHocDto			TextAdHocDto
//
//		An incoming instance of TextAdHocDto. This method
//		will compare all member variable data values in this
//		instance against those contained in the current instance
//		of TextAdHocDto. If the data values in both
//		instances are found to be equal in all respects, this
//		method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//		If the member variable data values contained in input
//		parameter 'incomingTxtAdHocDto' are equal in all
//		respects to those contained in the current instance of
//		TextAdHocDto, this method will return a boolean
//		value of 'true'. Otherwise, a value of 'false' will
//		be returned to the calling function.
func (textAdHocDto *TextAdHocDto) Equal(
	incomingTxtAdHocDto TextAdHocDto) bool {

	if textAdHocDto.lock == nil {
		textAdHocDto.lock = new(sync.Mutex)
	}

	textAdHocDto.lock.Lock()

	defer textAdHocDto.lock.Unlock()

	return textAdHocDtoMolecule{}.ptr().equal(
		textAdHocDto,
		&incomingTxtAdHocDto)
}

// textAdHocDtoNanobot - Provides helper methods for
// TextAdHocDto.
type textAdHocDtoNanobot struct {
	lock *sync.Mutex
}

// copyData - Copies all data from a source instance of
// TextAdHocDto to a destination instance of
// TextAdHocDto.
func (txtSolidLineDtoNanobot *textAdHocDtoNanobot) copyData(
	destinationTextAdHocDto *TextAdHocDto,
	sourceTextAdHocDto *TextAdHocDto,
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

	if sourceTextAdHocDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTextAdHocDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationTextAdHocDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTextAdHocDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	textAdHocDtoMolecule{}.ptr().empty(
		destinationTextAdHocDto)

	destinationTextAdHocDto.FormatType =
		sourceTextAdHocDto.FormatType

	destinationTextAdHocDto.LeftMarginStr =
		sourceTextAdHocDto.LeftMarginStr

	destinationTextAdHocDto.AdHocText =
		sourceTextAdHocDto.AdHocText

	destinationTextAdHocDto.RightMarginStr =
		sourceTextAdHocDto.RightMarginStr

	destinationTextAdHocDto.TurnLineTerminationOff =
		sourceTextAdHocDto.TurnLineTerminationOff

	destinationTextAdHocDto.LineTerminator =
		sourceTextAdHocDto.LineTerminator

	destinationTextAdHocDto.MaxLineLength =
		sourceTextAdHocDto.MaxLineLength

	destinationTextAdHocDto.TurnAutoLineLengthBreaksOn =
		sourceTextAdHocDto.TurnAutoLineLengthBreaksOn

	return err
}

// ptr - Returns a pointer to a new instance of
// textAdHocDtoNanobot.
func (txtSolidLineDtoNanobot textAdHocDtoNanobot) ptr() *textAdHocDtoNanobot {

	if txtSolidLineDtoNanobot.lock == nil {
		txtSolidLineDtoNanobot.lock = new(sync.Mutex)
	}

	txtSolidLineDtoNanobot.lock.Lock()

	defer txtSolidLineDtoNanobot.lock.Unlock()

	return &textAdHocDtoNanobot{
		lock: new(sync.Mutex),
	}
}

// textAdHocDtoMolecule - Provides helper methods for
// TextAdHocDto.
type textAdHocDtoMolecule struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of TextAdHocDto
// and proceeds to set all the internal member variables to their
// zero or uninitialized states.
//
// This method will therefore delete all data currently held
// by this instance of TextAdHocDto.
func (txtAdHocDtoMolecule *textAdHocDtoMolecule) empty(
	txtAdHocDto *TextAdHocDto) {

	if txtAdHocDtoMolecule.lock == nil {
		txtAdHocDtoMolecule.lock = new(sync.Mutex)
	}

	txtAdHocDtoMolecule.lock.Lock()

	defer txtAdHocDtoMolecule.lock.Unlock()

	txtAdHocDto.FormatType = TxtFieldType.None()

	txtAdHocDto.LeftMarginStr = ""

	txtAdHocDto.AdHocText = ""

	txtAdHocDto.RightMarginStr = ""

	txtAdHocDto.TurnLineTerminationOff = false

	txtAdHocDto.LineTerminator = ""

	txtAdHocDto.MaxLineLength = -99

	txtAdHocDto.TurnAutoLineLengthBreaksOn = false

	return
}

// equal - Receives pointers to two instances of TextAdHocDto
// and proceeds to compare all the member data variables for both
// instances.
//
// If the two instances of TextAdHocDto are found to be equal in
// all respects, this method will return a boolean value of 'true'.
func (txtAdHocDtoMolecule *textAdHocDtoMolecule) equal(
	txtAdHocDto1 *TextAdHocDto,
	txtAdHocDto2 *TextAdHocDto) bool {

	if txtAdHocDtoMolecule.lock == nil {
		txtAdHocDtoMolecule.lock = new(sync.Mutex)
	}

	txtAdHocDtoMolecule.lock.Lock()

	defer txtAdHocDtoMolecule.lock.Unlock()

	if txtAdHocDto1 == nil ||
		txtAdHocDto2 == nil {
		return false
	}

	if txtAdHocDto1.FormatType !=
		txtAdHocDto2.FormatType {

		return false
	}

	if txtAdHocDto1.LeftMarginStr !=
		txtAdHocDto2.LeftMarginStr {

		return false
	}

	if txtAdHocDto1.AdHocText !=
		txtAdHocDto2.AdHocText {

		return false
	}

	if txtAdHocDto1.RightMarginStr !=
		txtAdHocDto2.RightMarginStr {

		return false
	}

	if txtAdHocDto1.TurnLineTerminationOff !=
		txtAdHocDto2.TurnLineTerminationOff {

		return false
	}

	if txtAdHocDto1.LineTerminator !=
		txtAdHocDto2.LineTerminator {

		return false
	}

	if txtAdHocDto1.MaxLineLength !=
		txtAdHocDto2.MaxLineLength {

		return false
	}

	if txtAdHocDto1.TurnAutoLineLengthBreaksOn !=
		txtAdHocDto2.TurnAutoLineLengthBreaksOn {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// textFieldSpacerDtoMolecule.
func (txtAdHocDtoMolecule textAdHocDtoMolecule) ptr() *textAdHocDtoMolecule {

	if txtAdHocDtoMolecule.lock == nil {
		txtAdHocDtoMolecule.lock = new(sync.Mutex)
	}

	txtAdHocDtoMolecule.lock.Lock()

	defer txtAdHocDtoMolecule.lock.Unlock()

	return &textAdHocDtoMolecule{
		lock: new(sync.Mutex),
	}
}

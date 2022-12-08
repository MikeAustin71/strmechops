package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type TextFormatterDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format Operation to be performed. Depending on this value,
	// other member variables in the type will be accessed and used
	//
	// generated.
	//
	// Possible values are listed as follows:
	//   TxtFieldType.None()      - Invalid
	//  TxtFieldType.Label()
	//  TxtFieldType.DateTime()
	//  TxtFieldType.Filler()
	//  TxtFieldType.Spacer()
	//  TxtFieldType.BlankLine()
	//  TxtFieldType.SolidLine()
	//  TxtFieldType.LineColumns()
	//  TxtFieldType.TextAdHoc()

	DateTime TextFieldDateTimeDto
	// A structure containing data elements necessary for the
	// creation of Text Date/Time Fields.

	Filler TextFieldFillerDto
	// A structure containing data elements necessary for the
	// creation of Text Filler Fields.

	Label TextFieldLabelDto
	// A structure containing data elements necessary for the
	// creation of Text Label Fields.

	Spacer TextFieldSpacerDto
	// A structure containing data elements necessary for the
	// creation of Text Spacer Fields.

	BlankLine TextLineBlankDto
	// A structure containing data elements necessary for the
	// creation of one or more Blank Lines. A Blank Line
	// typically consists of one or more new line characters
	// ('\n').

	SolidLine TextLineSolidDto
	// A structure containing data elements necessary for the
	// creation of a solid line of text characters.
	//
	// A solid line, as used here, consists of a single character
	// or multiple characters used in a repeating sequence to
	// construct a solid line. Typically, solid lines consist of
	// dashes ("---"), underscore characters ("____"), equal signs
	// ("====="), asterisks ("*****") and other similar line break
	// presentations. The length of a solid line is specified by the
	// member variable, 'SolidLineCharsRepeatCount'.

	LineColumns TextLineColumnsDto
	// A structure containing data elements necessary for the
	// creation of a text line containing one or more Text Columns.

	LinesTimerStartStop TextLineTimerStartStopDto
	// A structure containing data elements necessary for the
	// creation of a series of text lines describing a timer
	// event containing a start time, end time and time
	// duration or elapsed time.

	TextAdHoc TextAdHocDto
	// The Text Ad Hoc Data Transfer Object is used to transfer
	// data parameters required for to inject a user generated
	// string of text into the stream of text characters being
	// formatted for screen display, file output or printing.
	//
	// Except for line breaks configured at the user's discretion,
	// no additional formatting is performed on this text, and it is
	// inserted "as is" in to the final output of formatted text.
	//

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFormatterDto ('incomingTxtFmtDto') to the data fields
// of the current TextFormatterDto instance
// ('txtFmtDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextFormatterDto instance
// ('txtFmtDto') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingTxtFmtDto'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingTxtFmtDto        TextFormatterDto
//	   - An instance of TextFormatterDto. This method will NOT
//	     change the data values of member variables contained in
//	     this instance.
//
//	     All data values in this TextFormatterDto instance
//	     ('incomingTxtFmtDto') will be copied to the current
//	     TextFormatterDto instance ('txtFmtDto').
//
//	     No data validation is performed on input parameter,
//	     'incomingTxtFmtDto'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (txtFmtDto *TextFormatterDto) CopyIn(
	incomingTxtFmtDto TextFormatterDto) {

	if txtFmtDto.lock == nil {
		txtFmtDto.lock = new(sync.Mutex)
	}

	txtFmtDto.lock.Lock()

	defer txtFmtDto.lock.Unlock()

	_ = textFormatterDtoNanobot{}.ptr().
		copyData(
			txtFmtDto,
			&incomingTxtFmtDto,
			nil)

	return
}

// CopyOut - Returns a deep copy of the current
// TextFormatterDto instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFormatterDto.
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
//	deepCopyTxtFmtDto     TextFormatterDto
//	   - This parameter will return a deep copy of the current
//	     TextFormatterDto instance.
func (txtFmtDto *TextFormatterDto) CopyOut() (
	deepCopyTxtFmtDto TextFormatterDto) {

	if txtFmtDto.lock == nil {
		txtFmtDto.lock = new(sync.Mutex)
	}

	txtFmtDto.lock.Lock()

	defer txtFmtDto.lock.Unlock()

	_ = textFormatterDtoNanobot{}.ptr().
		copyData(
			&deepCopyTxtFmtDto,
			txtFmtDto,
			nil)

	return deepCopyTxtFmtDto
}

// Empty - Resets all internal member variables for the current
// instance of TextFormatterDto to their zero or
// uninitialized states. This method will leave the current
// instance of TextFormatterDto in an invalid state and
// unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFormatterDto. All member
// variable data values will be reset to their zero or
// uninitialized states.
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
func (txtFmtDto *TextFormatterDto) Empty() {

	if txtFmtDto.lock == nil {
		txtFmtDto.lock = new(sync.Mutex)
	}

	txtFmtDto.lock.Lock()

	textFormatterDtoMolecule{}.ptr().
		empty(txtFmtDto)

	txtFmtDto.lock.Unlock()

	txtFmtDto.lock = nil

	return
}

// Equal - Receives another instance of TextFormatterDto and
// proceeds to compare the member variables to those of the current
// TextFormatterDto instance in order to determine if they are
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
//	incomingTxtFmtDto          TextFormatterDto
//	   - An incoming instance of TextFormatterDto. This method
//	     will compare all member variable data values in this
//	     instance against those contained in the current instance
//	     of TextFormatterDto. If the data values in both
//	     instances are found to be equal in all respects, this
//	     method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingTxtFmtDto' are equal in all respects
//	     to those contained in the current instance of
//	     TextFormatterDto, this method will return a boolean
//	     value of 'true'. Otherwise, a value of 'false' will be
//	     returned to the calling function.
func (txtFmtDto *TextFormatterDto) Equal(
	incomingTxtFmtDto TextFormatterDto) bool {

	if txtFmtDto.lock == nil {
		txtFmtDto.lock = new(sync.Mutex)
	}

	txtFmtDto.lock.Lock()

	defer txtFmtDto.lock.Unlock()

	return textFormatterDtoMolecule{}.ptr().
		equal(
			txtFmtDto,
			&incomingTxtFmtDto)
}

// textFormatterDtoNanobot - Provides helper methods for
// TextLineColumnsDto.
type textFormatterDtoNanobot struct {
	lock *sync.Mutex
}

// copyData - Copies all data from a source instance of
// TextFmtParamsLineColumnsDto to a destination instance of
// TextFmtParamsLineColumnsDto.
func (textFMtDtoNanobot *textFormatterDtoNanobot) copyData(
	destinationTxtFormatterDto *TextFormatterDto,
	sourceTxtFormatterDto *TextFormatterDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if textFMtDtoNanobot.lock == nil {
		textFMtDtoNanobot.lock = new(sync.Mutex)
	}

	textFMtDtoNanobot.lock.Lock()

	defer textFMtDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFormatterDtoNanobot."+
			"copyData()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtFormatterDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtFormatterDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationTxtFormatterDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtFormatterDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	textFormatterDtoMolecule{}.ptr().
		empty(destinationTxtFormatterDto)

	destinationTxtFormatterDto.FormatType =
		sourceTxtFormatterDto.FormatType

	destinationTxtFormatterDto.DateTime.CopyIn(
		sourceTxtFormatterDto.DateTime)

	destinationTxtFormatterDto.Filler.CopyIn(
		sourceTxtFormatterDto.Filler)

	destinationTxtFormatterDto.Label.CopyIn(
		sourceTxtFormatterDto.Label)

	destinationTxtFormatterDto.Spacer.CopyIn(
		sourceTxtFormatterDto.Spacer)

	destinationTxtFormatterDto.BlankLine.CopyIn(
		sourceTxtFormatterDto.BlankLine)

	destinationTxtFormatterDto.SolidLine.CopyIn(
		sourceTxtFormatterDto.SolidLine)

	destinationTxtFormatterDto.LineColumns.CopyIn(
		sourceTxtFormatterDto.LineColumns)

	destinationTxtFormatterDto.LinesTimerStartStop.CopyIn(
		sourceTxtFormatterDto.LinesTimerStartStop)

	destinationTxtFormatterDto.TextAdHoc.CopyIn(
		sourceTxtFormatterDto.TextAdHoc)

	return err
}

// ptr - Returns a pointer to a new instance of
// textFormatterDtoNanobot.
func (textFMtDtoNanobot textFormatterDtoNanobot) ptr() *textFormatterDtoNanobot {

	if textFMtDtoNanobot.lock == nil {
		textFMtDtoNanobot.lock = new(sync.Mutex)
	}

	textFMtDtoNanobot.lock.Lock()

	defer textFMtDtoNanobot.lock.Unlock()

	return &textFormatterDtoNanobot{
		lock: new(sync.Mutex),
	}
}

// textFormatterDtoMolecule - Provides helper methods for
// TextLineColumnsDto.
type textFormatterDtoMolecule struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of TextFormatterDto
// and proceeds to set all the internal member variables to their
// zero or uninitialized states.
//
// This method will therefore delete all data currently held
// by this instance of TextFormatterDto.
func (textFmtDtoMolecule *textFormatterDtoMolecule) empty(
	txtFormatterDto *TextFormatterDto) {

	if textFmtDtoMolecule.lock == nil {
		textFmtDtoMolecule.lock = new(sync.Mutex)
	}

	textFmtDtoMolecule.lock.Lock()

	defer textFmtDtoMolecule.lock.Unlock()

	if txtFormatterDto == nil {
		return
	}

	txtFormatterDto.FormatType = TxtFieldType.None()

	txtFormatterDto.DateTime.Empty()

	txtFormatterDto.Filler.Empty()

	txtFormatterDto.Label.Empty()

	txtFormatterDto.Spacer.Empty()

	txtFormatterDto.BlankLine.Empty()

	txtFormatterDto.SolidLine.Empty()

	txtFormatterDto.LineColumns.Empty()

	txtFormatterDto.LinesTimerStartStop.Empty()

	txtFormatterDto.TextAdHoc.Empty()

	return
}

// equal - Receives pointers to two instances of TextFormatterDto
// and proceeds to compare all the member data variables for both
// instances.
//
// If the two instances of TextFormatterDto are found to be equal
// in all respects, this method will return a boolean value of
// 'true'.
func (textFmtDtoMolecule *textFormatterDtoMolecule) equal(
	txtFormatterDto1 *TextFormatterDto,
	txtFormatterDto2 *TextFormatterDto) bool {

	if textFmtDtoMolecule.lock == nil {
		textFmtDtoMolecule.lock = new(sync.Mutex)
	}

	textFmtDtoMolecule.lock.Lock()

	defer textFmtDtoMolecule.lock.Unlock()

	if txtFormatterDto1 == nil ||
		txtFormatterDto2 == nil {
		return false
	}

	if txtFormatterDto1.FormatType !=
		txtFormatterDto2.FormatType {

		return false
	}

	if !txtFormatterDto1.DateTime.Equal(
		txtFormatterDto2.DateTime) {

		return false
	}

	if !txtFormatterDto1.Filler.Equal(
		txtFormatterDto2.Filler) {

		return false
	}

	if !txtFormatterDto1.Label.Equal(
		txtFormatterDto2.Label) {

		return false
	}

	if !txtFormatterDto1.Spacer.Equal(
		txtFormatterDto2.Spacer) {

		return false
	}

	if !txtFormatterDto1.BlankLine.Equal(
		txtFormatterDto2.BlankLine) {

		return false
	}

	if !txtFormatterDto1.SolidLine.Equal(
		txtFormatterDto2.SolidLine) {

		return false
	}

	if !txtFormatterDto1.LineColumns.Equal(
		txtFormatterDto2.LineColumns) {

		return false
	}

	if !txtFormatterDto1.LinesTimerStartStop.Equal(
		txtFormatterDto2.LinesTimerStartStop) {

		return false
	}

	if !txtFormatterDto1.TextAdHoc.Equal(
		txtFormatterDto2.TextAdHoc) {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// textFormatterDtoMolecule.
func (textFmtDtoMolecule textFormatterDtoMolecule) ptr() *textFormatterDtoMolecule {

	if textFmtDtoMolecule.lock == nil {
		textFmtDtoMolecule.lock = new(sync.Mutex)
	}

	textFmtDtoMolecule.lock.Lock()

	defer textFmtDtoMolecule.lock.Unlock()

	return &textFormatterDtoMolecule{
		lock: new(sync.Mutex),
	}
}

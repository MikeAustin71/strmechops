package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextFmtParamsLineColumnsDto - Contains Text Format Parameters for a
// Line containing one or more columns of text.
type TextFmtParamsLineColumnsDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format/Parameter Operation to be performed. For the
	// Line1Column Parameter Type, this value is always set to:
	//   TxtFieldType.Line1Column()

	FieldFormatParams []TextFieldFmtParamsDto
	// An array of Text Field Format Parameters.
	// The number of elements in this array must
	// be equal to the number of text columns to
	// be formatted. If this array is empty, the
	// current instance of TextFmtParamsLineColumnsDto
	// is invalid.

	TurnLineTerminationOff bool
	// If this parameter is set to 'true' no Line Termination
	// Sequence will be applied for this text line.

	LineTerminator string
	// If this parameter is submitted as an empty string,
	// the default new line terminator ('\n') will be applied.
	// If this parameter is populated, this character sequence
	// will be used as the Line Terminator for this text line.
	// Remember that the line termination operation will only
	// be performed if paramter 'TurnLineTerminationOff' is set
	// to 'false'.

	MaxLineLength int
	// Set this parameter to minus one -1 to specify an unlimited
	// line length for this text line.

	TurnAutoLineLengthBreaksOn bool
	// When this parameter is set to 'true', text fields which extend
	// beyond the maximum line length 'MaxLineLength' will be placed
	// on the following line of text.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFmtParamsLineColumnsDto ('incomingParams') to the data fields
// of the current TextFmtParamsLineColumnsDto instance
// ('paramsLineCol').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextFmtParamsLineColumnsDto instance
// ('paramsLineCol') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingParams'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingParams        TextFmtParamsLineColumnsDto
//	   - An instance of TextFmtParamsLineColumnsDto. This method will NOT
//	     change the data values of member variables contained in
//	     this instance.
//
//	     All data values in this TextFmtParamsLineColumnsDto instance
//	     ('incomingParams') will be copied to the current
//	     TextFmtParamsLineColumnsDto instance ('paramsLineCol').
//
//	     No data validation is performed on input parameter,
//	     'incomingParams'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NONE
func (paramsLineCol *TextFmtParamsLineColumnsDto) CopyIn(
	incomingParams TextFmtParamsLineColumnsDto) {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	_ = new(textFmtParamsLineColsNanobot).
		copy(paramsLineCol,
			&incomingParams,
			nil)

	return
}

// CopyOut - Returns a deep copy of the current
// TextFmtParamsLineColumnsDto instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFmtParamsLineColumnsDto.
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
//	deepCopyTxtFmtLineCols     TextFmtParamsLineColumnsDto
//	   - This parameter will return a deep copy of the current
//	     TextFmtParamsLineColumnsDto instance.
func (paramsLineCol *TextFmtParamsLineColumnsDto) CopyOut() (
	deepCopyTxtFmtLineCols TextFmtParamsLineColumnsDto) {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	_ = new(textFmtParamsLineColsNanobot).
		copy(
			&deepCopyTxtFmtLineCols,
			paramsLineCol,
			nil)

	return deepCopyTxtFmtLineCols
}

// Empty - Resets all internal member variables for the current
// instance of TextFmtParamsLineColumnsDto to their zero or
// uninitialized states. This method will leave the current
// instance of TextFmtParamsLineColumnsDto in an invalid state and
// unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFmtParamsLineColumnsDto. All member
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
func (paramsLineCol *TextFmtParamsLineColumnsDto) Empty() {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	new(textFmtParamsLineColsMolecule).
		empty(paramsLineCol)

	paramsLineCol.lock.Unlock()

	paramsLineCol.lock = nil
}

// Equal - Receives another instance of TextFieldFmtParamsDto and
// proceeds to compare the member variables to those of the current
// TextFieldFmtParamsDto instance in order to determine if they are
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
//	incomingFieldParams        TextFieldFmtParamsDto
//	   - An incoming instance of TextFieldFmtParamsDto. This method
//	     will compare all member variable data values in this
//	     instance against those contained in the current instance
//	     of TextFieldFmtParamsDto. If the data values in both
//	     instances are found to be equal in all respects, this
//	     method will return a boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the member variable data values contained in input
//	     parameter 'incomingFieldParams' are equal in all respects
//	     to those contained in the current instance of
//	     TextFieldFmtParamsDto, this method will return a boolean
//	     value of 'true'. Otherwise a value of 'false' will be
//	     returned to the calling function.
func (paramsLineCol *TextFmtParamsLineColumnsDto) Equal(
	incomingFmtLineCols TextFmtParamsLineColumnsDto) bool {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	return new(textFmtParamsLineColsMolecule).
		equal(
			paramsLineCol,
			&incomingFmtLineCols)
}

// GetTextParamsType - Returns the Text FormatType Type.
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
//	TextFieldType
//	   - This method returns an enumeration value specifying the
//	     Text Format Type associated with this instance of
//	     TextFmtParamsLineColumnsDto:
//	        TextFmtParamsLineColumnsDto.FormatType
func (paramsLineCol *TextFmtParamsLineColumnsDto) GetTextParamsType() TextFieldType {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	return paramsLineCol.FormatType
}

// GetNumOfFieldFmtParams - Returns the number of
// TextFieldFmtParams Data Transfer Objects contained in the Field
// Format Parameters array maintained by this instance of
// TextFmtParamsLineColumnsDto.
//
// The number of elements in the Text Field Format Parameters array
// must be equal to the number of text columns to be formatted.
//
// If this array is empty, the current instance of
// TextFmtParamsLineColumnsDto is invalid.
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
//	int
//	   - This method returns an integer value specifying the number
//	     of array objects in the Format Parameters array maintained
//	     by this instance of TextFmtParamsLineColumnsDto:
//	        TextFmtParamsLineColumnsDto.FieldFormatParams
func (paramsLineCol *TextFmtParamsLineColumnsDto) GetNumOfFieldFmtParams() int {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	return len(paramsLineCol.FieldFormatParams)
}

// IsValid - If the number of elements in the Text
// Field Format Parameters array is empty, this
// method will return false.
//
// The number of elements in the Text Field Format
// Parameters array must be equal to the number of
// text columns to be formatted.
//
// If this array contains one or more elements, this
// method will return 'true'.
func (paramsLineCol *TextFmtParamsLineColumnsDto) IsValid() bool {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	if len(paramsLineCol.FieldFormatParams) == 0 {
		return false
	}

	return true
}

// textFmtParamsLineColsNanobot - Provides helper methods for
// TextFmtParamsLineColumnsDto.
type textFmtParamsLineColsNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFmtParamsLineColumnsDto to a destination instance of
// TextFmtParamsLineColumnsDto.
func (lineColsNanobot *textFmtParamsLineColsNanobot) copy(
	destinationTxtParamsLineColsDto *TextFmtParamsLineColumnsDto,
	sourceTxtParamsLineColsDto *TextFmtParamsLineColumnsDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if lineColsNanobot.lock == nil {
		lineColsNanobot.lock = new(sync.Mutex)
	}

	lineColsNanobot.lock.Lock()

	defer lineColsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFmtParamsLineColsNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtParamsLineColsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtParamsLineColsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationTxtParamsLineColsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtParamsLineColsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(textFmtParamsLineColsMolecule).
		empty(destinationTxtParamsLineColsDto)

	destinationTxtParamsLineColsDto.FormatType =
		sourceTxtParamsLineColsDto.FormatType

	lenItems := len(sourceTxtParamsLineColsDto.FieldFormatParams)

	if lenItems > 0 {

		destinationTxtParamsLineColsDto.FieldFormatParams =
			make([]TextFieldFmtParamsDto, lenItems)

		for i := 0; i < lenItems; i++ {

			destinationTxtParamsLineColsDto.FieldFormatParams[i].CopyIn(
				sourceTxtParamsLineColsDto.FieldFormatParams[i])

		}
	}

	destinationTxtParamsLineColsDto.TurnLineTerminationOff =
		sourceTxtParamsLineColsDto.TurnLineTerminationOff

	destinationTxtParamsLineColsDto.LineTerminator =
		sourceTxtParamsLineColsDto.LineTerminator

	destinationTxtParamsLineColsDto.MaxLineLength =
		sourceTxtParamsLineColsDto.MaxLineLength

	destinationTxtParamsLineColsDto.TurnAutoLineLengthBreaksOn =
		sourceTxtParamsLineColsDto.TurnAutoLineLengthBreaksOn

	return err
}

// textFmtParamsLineColsMolecule - Provides helper methods for
// TextFmtParamsLineColumnsDto.
type textFmtParamsLineColsMolecule struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// TextFmtParamsLineColumnsDto and proceeds to set all the member
// variables to their zero or uninitialized states.
//
// Are previous data values contained in this instance of
// TextFmtParamsLineColumnsDto will be deleted.
func (lineColsMolecule textFmtParamsLineColsMolecule) empty(
	targetTxtParamsLineColsDto *TextFmtParamsLineColumnsDto) {

	if lineColsMolecule.lock == nil {
		lineColsMolecule.lock = new(sync.Mutex)
	}

	lineColsMolecule.lock.Lock()

	defer lineColsMolecule.lock.Unlock()

	if targetTxtParamsLineColsDto == nil {
		return
	}

	targetTxtParamsLineColsDto.FormatType = TxtFieldType.None()

	lenItems := len(targetTxtParamsLineColsDto.FieldFormatParams)

	for i := 0; i < lenItems; i++ {
		targetTxtParamsLineColsDto.FieldFormatParams[i].Empty()
	}

	targetTxtParamsLineColsDto.FieldFormatParams = nil

	targetTxtParamsLineColsDto.TurnLineTerminationOff = false

	targetTxtParamsLineColsDto.LineTerminator = ""

	targetTxtParamsLineColsDto.MaxLineLength = -1

	targetTxtParamsLineColsDto.TurnAutoLineLengthBreaksOn = false

	return
}

// equal - Compares two instances of TextFmtParamsLineColumnsDto
// and returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextFmtParamsLineColumnsDto are equal,
// this method returns 'true'.
func (lineColsMolecule textFmtParamsLineColsMolecule) equal(
	txtFmtParamsDto1 *TextFmtParamsLineColumnsDto,
	txtFmtParamsDto2 *TextFmtParamsLineColumnsDto) bool {

	if lineColsMolecule.lock == nil {
		lineColsMolecule.lock = new(sync.Mutex)
	}

	lineColsMolecule.lock.Lock()

	defer lineColsMolecule.lock.Unlock()

	if txtFmtParamsDto1 == nil ||
		txtFmtParamsDto2 == nil {

		return false
	}

	if txtFmtParamsDto1.FormatType !=
		txtFmtParamsDto2.FormatType {

		return false
	}

	lenItems1 := len(txtFmtParamsDto1.FieldFormatParams)

	lenItems2 := len(txtFmtParamsDto2.FieldFormatParams)

	if lenItems1 != lenItems2 {
		return false
	}

	for i := 0; i < lenItems1; i++ {

		if !txtFmtParamsDto1.FieldFormatParams[i].Equal(
			txtFmtParamsDto2.FieldFormatParams[i]) {

			return false
		}
	}

	if txtFmtParamsDto1.LineTerminator !=
		txtFmtParamsDto2.LineTerminator {

		return false
	}

	if txtFmtParamsDto1.TurnAutoLineLengthBreaksOn !=
		txtFmtParamsDto2.TurnAutoLineLengthBreaksOn {

		return false
	}

	return true
}

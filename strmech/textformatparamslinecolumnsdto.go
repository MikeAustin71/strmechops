package strmech

import "sync"

// TextFmtParamsLineColumnsDto - Contains Text Format Parameters for a
// Line containing one or more columns of text.
type TextFmtParamsLineColumnsDto struct {
	FormatType TextFieldType
	// Required. This enumeration value specifies the type of Text
	// Format/Parameter Operation to be performed. For the
	// Line1Column Parameter Type, this value is always set to:
	//   TxtFieldType.Line1Column()

	FieldFormatParams []TextFieldFmtParamsDto

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

	isValid bool

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFmtParamsLineColumnsDto ('incomingParams') to the data fields
// of the current TextFmtParamsLineColumnsDto instance
// ('paramsLineCol').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextFmtParamsLineColumnsDto instance
// ('paramsLineCol') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingParams'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingParams        TextFmtParamsLineColumnsDto
//     - An instance of TextFmtParamsLineColumnsDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this TextFmtParamsLineColumnsDto instance
//       ('incomingParams') will be copied to the current
//       TextFmtParamsLineColumnsDto instance ('paramsLineCol').
//
//       No data validation is performed on input parameter,
//       'incomingParams'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (paramsLineCol *TextFmtParamsLineColumnsDto) CopyIn(
	incomingParams TextFmtParamsLineColumnsDto) {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	paramsLineCol.FormatType = incomingParams.FormatType

	paramsLineCol.FieldFormatParams = nil

	lenCol := len(incomingParams.FieldFormatParams)

	if lenCol > 0 {
		paramsLineCol.FieldFormatParams =
			make([]TextFieldFmtParamsDto, lenCol)

		for i := 0; i < lenCol; i++ {

			paramsLineCol.FieldFormatParams[i].CopyIn(
				incomingParams.FieldFormatParams[i])

		}
	}

	paramsLineCol.TurnLineTerminationOff =
		incomingParams.TurnLineTerminationOff

	paramsLineCol.LineTerminator =
		incomingParams.LineTerminator

	paramsLineCol.MaxLineLength =
		incomingParams.MaxLineLength

	paramsLineCol.TurnAutoLineLengthBreaksOn =
		incomingParams.TurnAutoLineLengthBreaksOn

	paramsLineCol.isValid = incomingParams.isValid

	return
}

// CopyOut - Returns a deep copy of the current
// TextFmtParamsLineColumnsDto instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFmtParamsLineColumnsDto.
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
//  deepCopyTxtFmtLineCols     TextFmtParamsLineColumnsDto
//     - This parameter will return a deep copy of the current
//       TextFmtParamsLineColumnsDto instance.
//
func (paramsLineCol *TextFmtParamsLineColumnsDto) CopyOut() (
	deepCopyTxtFmtLineCols TextFmtParamsLineColumnsDto) {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	deepCopyTxtFmtLineCols.FormatType = paramsLineCol.FormatType

	deepCopyTxtFmtLineCols.FieldFormatParams = nil

	lenCol := len(paramsLineCol.FieldFormatParams)

	if lenCol > 0 {
		deepCopyTxtFmtLineCols.FieldFormatParams =
			make([]TextFieldFmtParamsDto, lenCol)

		for i := 0; i < lenCol; i++ {

			deepCopyTxtFmtLineCols.FieldFormatParams[i].CopyIn(
				paramsLineCol.FieldFormatParams[i])

		}
	}

	deepCopyTxtFmtLineCols.TurnLineTerminationOff =
		paramsLineCol.TurnLineTerminationOff

	deepCopyTxtFmtLineCols.LineTerminator =
		paramsLineCol.LineTerminator

	deepCopyTxtFmtLineCols.MaxLineLength =
		paramsLineCol.MaxLineLength

	deepCopyTxtFmtLineCols.TurnAutoLineLengthBreaksOn =
		paramsLineCol.TurnAutoLineLengthBreaksOn

	deepCopyTxtFmtLineCols.isValid = paramsLineCol.isValid

	return deepCopyTxtFmtLineCols
}

// Empty - Resets all internal member variables for the current
// instance of TextFmtParamsLineColumnsDto to their zero or
// uninitialized states. This method will leave the current
// instance of TextFmtParamsLineColumnsDto in an invalid state and
// unavailable for immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFmtParamsLineColumnsDto. All member
// variable data values will be reset to their zero or
// uninitialized states.
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
func (paramsLineCol *TextFmtParamsLineColumnsDto) Empty() {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	paramsLineCol.FormatType = TxtFieldType.None()

	lenCol := len(paramsLineCol.FieldFormatParams)

	for i := 0; i < lenCol; i++ {

		paramsLineCol.FieldFormatParams[i].Empty()

	}

	paramsLineCol.FieldFormatParams = nil

	paramsLineCol.TurnLineTerminationOff = false

	paramsLineCol.LineTerminator = ""

	paramsLineCol.MaxLineLength = -1

	paramsLineCol.TurnAutoLineLengthBreaksOn = false

	paramsLineCol.isValid = false

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
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFieldParams        TextFieldFmtParamsDto
//     - An incoming instance of TextFieldFmtParamsDto. This method
//       will compare all member variable data values in this
//       instance against those contained in the current instance
//       of TextFieldFmtParamsDto. If the data values in both
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
//       parameter 'incomingFieldParams' are equal in all respects
//       to those contained in the current instance of
//       TextFieldFmtParamsDto, this method will return a boolean
//       value of 'true'. Otherwise a value of 'false' will be
//       returned to the calling function.
//
func (paramsLineCol *TextFmtParamsLineColumnsDto) Equal(
	incomingFmtLineCols TextFmtParamsLineColumnsDto) bool {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	if paramsLineCol.FormatType !=
		incomingFmtLineCols.FormatType {

		return false
	}

	lenCol := len(incomingFmtLineCols.FieldFormatParams)

	if len(paramsLineCol.FieldFormatParams) !=
		lenCol {

		return false
	}

	if lenCol > 0 {

		for i := 0; i < lenCol; i++ {

			if !paramsLineCol.FieldFormatParams[i].Equal(
				incomingFmtLineCols.FieldFormatParams[i]) {

				return false
			}

		}
	}

	if paramsLineCol.TurnLineTerminationOff !=
		incomingFmtLineCols.TurnLineTerminationOff {

		return false
	}

	if paramsLineCol.LineTerminator !=
		incomingFmtLineCols.LineTerminator {

		return false
	}

	if paramsLineCol.MaxLineLength !=
		incomingFmtLineCols.MaxLineLength {

		return false
	}

	if paramsLineCol.TurnAutoLineLengthBreaksOn !=
		incomingFmtLineCols.TurnAutoLineLengthBreaksOn {

		return false
	}

	if paramsLineCol.isValid != incomingFmtLineCols.isValid {

		return false
	}

	return true
}

// GetTextParamsType - Returns the Text FormatType Type.
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
//  TextFieldType
//     - This method returns an enumeration value specifying the
//       Text Format Type associated with this instance of
//       TextFmtParamsLineColumnsDto:
//          TextFmtParamsLineColumnsDto.FormatType
//
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
//  int
//     - This method returns an integer value specifying the number
//       of array objects in the Format Parameters array maintained
//       by this instance of TextFmtParamsLineColumnsDto:
//          TextFmtParamsLineColumnsDto.FieldFormatParams
//
func (paramsLineCol *TextFmtParamsLineColumnsDto) GetNumOfFieldFmtParams() int {

	if paramsLineCol.lock == nil {
		paramsLineCol.lock = new(sync.Mutex)
	}

	paramsLineCol.lock.Lock()

	defer paramsLineCol.lock.Unlock()

	return len(paramsLineCol.FieldFormatParams)
}

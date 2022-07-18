package strmech

import "sync"

type TextFieldFmtParamsDto struct {
	LeftMarginStr  string
	FieldLength    int
	FieldJustify   TextJustify
	DateTimeFormat string
	RightMarginStr string

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldFmtParamsDto ('incomingFieldParams') to the data fields
// of the current TextFieldFmtParamsDto instance
// ('txtFieldFmtParams').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextFieldFmtParamsDto instance
// ('txtFieldFmtParams') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingFieldParams'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFieldParams        TextFieldFmtParamsDto
//     - An instance of TextFieldFmtParamsDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this TextFieldFmtParamsDto instance
//       ('incomingFieldParams') will be copied to the current
//       TextFieldFmtParamsDto instance ('txtFieldFmtParams').
//
//       No data validation is performed on input parameter,
//       'incomingFieldParams'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (txtFieldFmtParams *TextFieldFmtParamsDto) CopyIn(
	incomingFieldParams TextFieldFmtParamsDto) {

	if txtFieldFmtParams.lock == nil {
		txtFieldFmtParams.lock = new(sync.Mutex)
	}

	txtFieldFmtParams.lock.Lock()

	defer txtFieldFmtParams.lock.Unlock()

	txtFieldFmtParams.LeftMarginStr =
		incomingFieldParams.LeftMarginStr

	txtFieldFmtParams.FieldLength =
		incomingFieldParams.FieldLength

	txtFieldFmtParams.FieldJustify =
		incomingFieldParams.FieldJustify

	txtFieldFmtParams.DateTimeFormat =
		incomingFieldParams.DateTimeFormat

	txtFieldFmtParams.RightMarginStr =
		incomingFieldParams.RightMarginStr

	return
}

// CopyOut - Returns a deep copy of the current
// TextFieldFmtParamsDto instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFieldFmtParamsDto.
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
//  deepCopyTxtFieldParams     TextFieldFmtParamsDto
//     - This parameter will return a deep copy of the current
//       TextFieldFmtParamsDto instance.
//
func (txtFieldFmtParams *TextFieldFmtParamsDto) CopyOut() (
	deepCopyTxtFieldParams TextFieldFmtParamsDto) {

	if txtFieldFmtParams.lock == nil {
		txtFieldFmtParams.lock = new(sync.Mutex)
	}

	txtFieldFmtParams.lock.Lock()

	defer txtFieldFmtParams.lock.Unlock()

	deepCopyTxtFieldParams.LeftMarginStr =
		txtFieldFmtParams.LeftMarginStr

	deepCopyTxtFieldParams.FieldLength =
		txtFieldFmtParams.FieldLength

	deepCopyTxtFieldParams.FieldJustify =
		txtFieldFmtParams.FieldJustify

	deepCopyTxtFieldParams.DateTimeFormat =
		txtFieldFmtParams.DateTimeFormat

	deepCopyTxtFieldParams.RightMarginStr =
		txtFieldFmtParams.RightMarginStr

	return deepCopyTxtFieldParams
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldFmtParamsDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextFieldFmtParamsDto in an invalid state and unavailable for
// immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFieldFmtParamsDto. All member variable
// data values will be reset to their zero or uninitialized states.
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
func (txtFieldFmtParams *TextFieldFmtParamsDto) Empty() {

	if txtFieldFmtParams.lock == nil {
		txtFieldFmtParams.lock = new(sync.Mutex)
	}

	txtFieldFmtParams.lock.Lock()

	txtFieldFmtParams.LeftMarginStr = ""

	txtFieldFmtParams.FieldLength = -1

	txtFieldFmtParams.FieldJustify = TxtJustify.None()

	txtFieldFmtParams.DateTimeFormat = ""

	txtFieldFmtParams.RightMarginStr = ""

	txtFieldFmtParams.lock.Unlock()

	txtFieldFmtParams.lock = nil

	return
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
func (txtFieldFmtParams *TextFieldFmtParamsDto) Equal(
	incomingFieldParams TextFieldFmtParamsDto) bool {

	if txtFieldFmtParams.lock == nil {
		txtFieldFmtParams.lock = new(sync.Mutex)
	}

	txtFieldFmtParams.lock.Lock()

	defer txtFieldFmtParams.lock.Unlock()

	if txtFieldFmtParams.LeftMarginStr !=
		incomingFieldParams.LeftMarginStr {

		return false
	}

	if txtFieldFmtParams.FieldLength !=
		incomingFieldParams.FieldLength {

		return false
	}

	if txtFieldFmtParams.FieldJustify !=
		incomingFieldParams.FieldJustify {

		return false
	}

	if txtFieldFmtParams.DateTimeFormat !=
		incomingFieldParams.DateTimeFormat {

		return false
	}

	if txtFieldFmtParams.RightMarginStr !=
		incomingFieldParams.RightMarginStr {

		return false
	}

	return true
}

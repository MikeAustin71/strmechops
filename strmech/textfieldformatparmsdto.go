package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

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

	_ = textFieldFmtParamsDtoNanobot{}.ptr().
		copy(
			&deepCopyTxtFieldParams,
			txtFieldFmtParams,
			nil)

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

	textFieldFmtParamsDtoNanobot{}.ptr().
		empty(
			txtFieldFmtParams)

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

	return textFieldFmtParamsDtoNanobot{}.ptr().
		equal(txtFieldFmtParams,
			&incomingFieldParams)
}

// textFieldFmtParamsDtoNanobot - Provides helper methods for
// TextFieldFmtParamsDto.
type textFieldFmtParamsDtoNanobot struct {
	lock *sync.Mutex
}

// copy - Copies all data from a source instance of
// TextFieldFmtParamsDto to a destination instance of
// TextFieldFmtParamsDto.
func (txtFmtParamsNanobot *textFieldFmtParamsDtoNanobot) copy(
	destinationTxtFmtParamsDto *TextFieldFmtParamsDto,
	sourceTxtFmtParamsDto *TextFieldFmtParamsDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFmtParamsNanobot.lock == nil {
		txtFmtParamsNanobot.lock = new(sync.Mutex)
	}

	txtFmtParamsNanobot.lock.Lock()

	defer txtFmtParamsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldFmtParamsDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtFmtParamsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtFmtParamsDto' is a nil pointer!\n",
			ePrefix.String())

	}

	if destinationTxtFmtParamsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtFmtParamsDto' is a nil pointer!\n",
			ePrefix.String())

	}

	destinationTxtFmtParamsDto.LeftMarginStr =
		sourceTxtFmtParamsDto.LeftMarginStr

	destinationTxtFmtParamsDto.FieldLength =
		sourceTxtFmtParamsDto.FieldLength

	destinationTxtFmtParamsDto.FieldJustify =
		sourceTxtFmtParamsDto.FieldJustify

	destinationTxtFmtParamsDto.DateTimeFormat =
		sourceTxtFmtParamsDto.DateTimeFormat

	destinationTxtFmtParamsDto.RightMarginStr =
		sourceTxtFmtParamsDto.RightMarginStr

	return err
}

// empty - Receives a pointer to an instance of
// TextFieldFmtParamsDto and proceeds to set all the member
// variables to their zero or uninitialized states.
//
// Are previous data values contained in this instance of
// TextFieldFmtParamsDto will be deleted.
//
func (txtFmtParamsNanobot *textFieldFmtParamsDtoNanobot) empty(
	targetTxtFmtParamsDto *TextFieldFmtParamsDto) {

	if txtFmtParamsNanobot.lock == nil {
		txtFmtParamsNanobot.lock = new(sync.Mutex)
	}

	txtFmtParamsNanobot.lock.Lock()

	defer txtFmtParamsNanobot.lock.Unlock()

	if targetTxtFmtParamsDto == nil {
		return
	}

	targetTxtFmtParamsDto.LeftMarginStr = ""

	targetTxtFmtParamsDto.FieldLength = -1

	targetTxtFmtParamsDto.FieldJustify = TxtJustify.None()

	targetTxtFmtParamsDto.DateTimeFormat = ""

	targetTxtFmtParamsDto.RightMarginStr = ""

}

// equal - Compares two instances of TextFieldFmtParamsDto
// and returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextFieldFmtParamsDto are equal,
// this method returns 'true'.
//
func (txtFmtParamsNanobot *textFieldFmtParamsDtoNanobot) equal(
	txtFmtParamsDto1 *TextFieldFmtParamsDto,
	txtFmtParamsDto2 *TextFieldFmtParamsDto) bool {

	if txtFmtParamsDto1 == nil ||
		txtFmtParamsDto2 == nil {

		return false
	}

	if txtFmtParamsDto1.LeftMarginStr !=
		txtFmtParamsDto2.LeftMarginStr {

		return false
	}

	if txtFmtParamsDto1.FieldLength !=
		txtFmtParamsDto2.FieldLength {

		return false
	}

	if txtFmtParamsDto1.FieldJustify !=
		txtFmtParamsDto2.FieldJustify {

		return false
	}

	if txtFmtParamsDto1.DateTimeFormat !=
		txtFmtParamsDto2.DateTimeFormat {

		return false
	}

	if txtFmtParamsDto1.RightMarginStr !=
		txtFmtParamsDto2.RightMarginStr {

		return false
	}

	return true

}

// ptr - Returns a pointer to a new instance of
// textFieldFmtParamsDtoNanobot.
//
func (txtFmtParamsNanobot textFieldFmtParamsDtoNanobot) ptr() *textFieldFmtParamsDtoNanobot {

	if txtFmtParamsNanobot.lock == nil {
		txtFmtParamsNanobot.lock = new(sync.Mutex)
	}

	txtFmtParamsNanobot.lock.Lock()

	defer txtFmtParamsNanobot.lock.Unlock()

	return &textFieldFmtParamsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

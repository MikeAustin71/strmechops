package strmech

import (
	"sync"
	"time"
)

type TextFieldsContentDto struct {
	TextFieldString string
	// Either TextFieldString or TextFieldDateTime should
	// be configured, but NOT both.

	TextFieldDateTime time.Time
	// Either TextFieldString or TextFieldDateTime should
	// be configured, but NOT both.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextFieldsContentDto ('incomingFieldContent') to the data fields
// of the current TextFieldsContentDto instance
// ('txtFieldsContent').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextFieldsContentDto instance
// ('txtFieldsContent') will be deleted and overwritten.
//
// NO DATA VALIDATION IS performed on input parameter,
// 'incomingFieldContent'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingFieldContent       TextFieldsContentDto
//     - An instance of TextFieldsContentDto. This method will NOT
//       change the data values of member variables contained in
//       this instance.
//
//       All data values in this TextFieldsContentDto instance
//       ('incomingFieldContent') will be copied to the current
//       TextFieldsContentDto instance ('txtFieldsContent').
//
//       No data validation is performed on input parameter,
//       'incomingFieldContent'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (txtFieldsContent *TextFieldsContentDto) CopyIn(
	incomingFieldContent TextFieldsContentDto) {

	if txtFieldsContent.lock == nil {
		txtFieldsContent.lock = new(sync.Mutex)
	}

	txtFieldsContent.lock.Lock()

	defer txtFieldsContent.lock.Unlock()

	txtFieldsContent.TextFieldString =
		incomingFieldContent.TextFieldString

	txtFieldsContent.TextFieldDateTime =
		incomingFieldContent.TextFieldDateTime

	return
}

// CopyOut - Returns a deep copy of the current
// TextFieldsContentDto instance.
//
// NO DATA VALIDATION is performed on the current instance of
// TextFieldsContentDto.
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
//  deepCopyTxtFieldsContent   TextFieldsContentDto
//     - This parameter will return a deep copy of the current
//       TextFieldsContentDto instance.
//
func (txtFieldsContent *TextFieldsContentDto) CopyOut() (
	deepCopyTxtFieldsContent TextFieldsContentDto) {

	if txtFieldsContent.lock == nil {
		txtFieldsContent.lock = new(sync.Mutex)
	}

	txtFieldsContent.lock.Lock()

	defer txtFieldsContent.lock.Unlock()

	deepCopyTxtFieldsContent.TextFieldString =
		txtFieldsContent.TextFieldString

	deepCopyTxtFieldsContent.TextFieldDateTime =
		txtFieldsContent.TextFieldDateTime

	return deepCopyTxtFieldsContent
}

// Empty - Resets all internal member variables for the current
// instance of TextFieldsContentDto to their zero or uninitialized
// states. This method will leave the current instance of
// TextFieldsContentDto in an invalid state and unavailable for
// immediate reuse.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all member variable data values in this
// current instance of TextFieldsContentDto. All member variable
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
func (txtFieldsContent *TextFieldsContentDto) Empty() {

	if txtFieldsContent.lock == nil {
		txtFieldsContent.lock = new(sync.Mutex)
	}

	txtFieldsContent.lock.Lock()

	txtFieldsContent.TextFieldString = ""

	txtFieldsContent.TextFieldDateTime = time.Time{}

	txtFieldsContent.lock.Unlock()

	txtFieldsContent.lock = nil

}

// Equal - Receives another instance of TextFieldsContentDto and
// proceeds to compare the member variables to those of the current
// TextFieldsContentDto instance in order to determine if they are
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
//  incomingFieldContent       TextFieldsContentDto
//     - An incoming instance of TextFieldsContentDto. This method
//       will compare all member variable data values in this
//       instance against those contained in the current instance
//       of TextFieldsContentDto. If the data values in both
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
//       parameter 'incomingFieldContent' are equal in all respects
//       to those contained in the current instance of
//       TextFieldsContentDto, this method will return a boolean
//       value of 'true'. Otherwise a value of 'false' will be
//       returned to the calling function.
//
func (txtFieldsContent *TextFieldsContentDto) Equal(
	incomingFieldContent TextFieldsContentDto) bool {

	if txtFieldsContent.lock == nil {
		txtFieldsContent.lock = new(sync.Mutex)
	}

	txtFieldsContent.lock.Lock()

	defer txtFieldsContent.lock.Unlock()

	if txtFieldsContent.TextFieldString !=
		incomingFieldContent.TextFieldString {

		return false
	}

	if txtFieldsContent.TextFieldDateTime !=
		incomingFieldContent.TextFieldDateTime {

		return false
	}

	return true
}

package strmech

import (
	"strings"
	"sync"
	"time"
)

type TextFieldSpecDateTime struct {
	dateTime time.Time // The content of the datetime text.
	fieldLen int       // The length of the text field in which
	//               //  the text label will be positioned.
	dateTimeFormat string // Holds the format parameters used to
	//                        //  format the datetime for text presentation.
	textJustification TextJustify // The specification which controls
	//                            //  how the datetime text will be positioned
	//                            //  within the text field: 'Left', 'Right'
	//                            //  or 'Center'.
	textLineReader *strings.Reader
	lock           *sync.Mutex
}

// Empty - Resets all internal member variables to their initial
// or zero states.
//
func (txtDateTimeField *TextFieldSpecDateTime) Empty() {

	if txtDateTimeField.lock == nil {
		txtDateTimeField.lock = new(sync.Mutex)
	}

	txtDateTimeField.lock.Lock()

	textFieldSpecDateTimeAtom{}.ptr().empty(
		txtDateTimeField)

	txtDateTimeField.lock.Unlock()

	txtDateTimeField.lock = nil
}

// Equal - Receives a pointer to another instance of
// TextFieldSpecDateTime and proceeds to compare the member
// variables to those of the current TextFieldSpecDateTime instance
// in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables are equal in all respects,
// this flag is set to 'true'. Otherwise, this method returns
// 'false'.
//
func (txtDateTimeField *TextFieldSpecDateTime) Equal(
	incomingTxtFieldDateTime *TextFieldSpecDateTime) bool {

	if txtDateTimeField.lock == nil {
		txtDateTimeField.lock = new(sync.Mutex)
	}

	txtDateTimeField.lock.Lock()

	defer txtDateTimeField.lock.Unlock()

	return textFieldSpecDateTimeAtom{}.ptr().equal(
		txtDateTimeField,
		incomingTxtFieldDateTime)
}

// New - Returns a new concrete instance of TextFieldSpecDateTime.
// This returned instance is empty and unpopulated. All the member
// variables contained in this new instance are set to their
// uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecDateTime instance is
// invalid. Therefore, in order to use this TextFieldSpecDateTime
// instance, users must later call the setter methods on this type
// in order to configure valid and meaningful meaningful member
// variable data values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  TextFieldSpecDateTime
//     - This parameter returns a new and empty concrete instance
//       of TextFieldSpecDateTime. Member variable data values are
//       set to their initial or zero values.
//
func (txtDateTimeField TextFieldSpecDateTime) New() TextFieldSpecDateTime {

	if txtDateTimeField.lock == nil {
		txtDateTimeField.lock = new(sync.Mutex)
	}

	txtDateTimeField.lock.Lock()

	defer txtDateTimeField.lock.Unlock()

	newTxtFieldDateTime := TextFieldSpecDateTime{}

	newTxtFieldDateTime.textJustification = TextJustify(0).None()

	newTxtFieldDateTime.textLineReader = nil

	newTxtFieldDateTime.lock = new(sync.Mutex)

	return newTxtFieldDateTime
}

// NewPtr - Returns a pointer to a new unpopulated instance of
// TextFieldSpecDateTime. All the member variables contained in
// this new instance are set to their uninitialized or zero values.
//
// Be advised that setting member variables to their zero values
// means that the returned TextFieldSpecDateTime instance is
// invalid. Therefore, in order to use this TextFieldSpecDateTime
// instance, users must later call the setter methods on this type
// in order to configure valid and meaningful member variable data
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  --- NONE ---
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  *TextFieldSpecDateTime
//     - This parameter returns a pointer to a new, empty instance
//       of TextFieldSpecDateTime. Member variable data values are
//       set to their initial or zero values.
//
func (txtDateTimeField TextFieldSpecDateTime) NewPtr() *TextFieldSpecDateTime {

	if txtDateTimeField.lock == nil {
		txtDateTimeField.lock = new(sync.Mutex)
	}

	txtDateTimeField.lock.Lock()

	defer txtDateTimeField.lock.Unlock()

	newTxtFieldDateTime := TextFieldSpecDateTime{}

	newTxtFieldDateTime.textJustification = TextJustify(0).None()

	newTxtFieldDateTime.textLineReader = nil

	newTxtFieldDateTime.lock = new(sync.Mutex)

	return &newTxtFieldDateTime
}

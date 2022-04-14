package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
	"time"
)

type textFieldSpecDateTimeAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// TextFieldSpecDateTime and proceeds to set all the internal
// member variables to their uninitialized or zero states.
//
// IMPORTANT
// ----------------------------------------------------------------
// The values of all member variables contained in input parameter
// 'dateTimeTxtField' will be overwritten and replaced.
//
func (txtFieldDateTimeAtom *textFieldSpecDateTimeAtom) empty(
	dateTimeTxtField *TextFieldSpecDateTime) {

	if txtFieldDateTimeAtom.lock == nil {
		txtFieldDateTimeAtom.lock = new(sync.Mutex)
	}

	txtFieldDateTimeAtom.lock.Lock()

	defer txtFieldDateTimeAtom.lock.Unlock()

	if dateTimeTxtField == nil {
		return
	}

	dateTimeTxtField.dateTime = time.Time{}

	dateTimeTxtField.fieldLen = 0

	dateTimeTxtField.dateTimeFormat = ""

	dateTimeTxtField.textJustification = TextJustify(0).None()

	dateTimeTxtField.textLineReader = nil

	return
}

// equal - Receives pointers to two instances of
// TextFieldSpecDateTime and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (txtFieldDateTimeAtom *textFieldSpecDateTimeAtom) equal(
	dateTimeTxtFieldOne *TextFieldSpecDateTime,
	dateTimeTxtFieldTwo *TextFieldSpecDateTime) bool {

	if txtFieldDateTimeAtom.lock == nil {
		txtFieldDateTimeAtom.lock = new(sync.Mutex)
	}

	txtFieldDateTimeAtom.lock.Lock()

	defer txtFieldDateTimeAtom.lock.Unlock()

	if dateTimeTxtFieldOne == nil ||
		dateTimeTxtFieldTwo == nil {
		return false
	}

	if !dateTimeTxtFieldOne.dateTime.Equal(dateTimeTxtFieldTwo.dateTime) {
		return false
	}

	if dateTimeTxtFieldOne.fieldLen !=
		dateTimeTxtFieldTwo.fieldLen {
		return false
	}

	if dateTimeTxtFieldOne.dateTimeFormat !=
		dateTimeTxtFieldTwo.dateTimeFormat {
		return false
	}

	if dateTimeTxtFieldOne.textJustification !=
		dateTimeTxtFieldTwo.textJustification {
		return false
	}

	return true
}

// isValidTextFieldLabel - This method receives a pointer to
// an instance of TextFieldSpecDateTime and performs a diagnostic
// on the object to determine if it is valid or not.
//
// This method returns two parameters, 'isValid' and 'err'. If
// input parameter 'dateTimeTxtField' is valid, the return
// parameter 'isValid' is set to 'true' and 'err' is set to 'nil'.
//
// If input parameter 'dateTimeTxtField' is invalid, the return
// parameter 'isValid' is set to 'false' and 'err' is configured
// with an appropriate error message.
//
func (txtFieldDateTimeAtom *textFieldSpecDateTimeAtom) isValidTextFieldDateTime(
	dateTimeTxtField *TextFieldSpecDateTime,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	isValid = false
	err = nil

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecDateTimeAtom.isValidTextFieldDateTime()",
		"")

	if err != nil {
		return isValid, err
	}

	if dateTimeTxtField == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dateTimeTxtField' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if dateTimeTxtField.dateTime.IsZero() {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dateTimeTxtField.dateTime'\n"+
			"contains a zero value!\n",
			ePrefix.String())

		return isValid, err
	}

	txtLabelElectron := textFieldSpecLabelElectron{}

	err = txtLabelElectron.isFieldLengthValid(
		dateTimeTxtField.fieldLen,
		ePrefix.XCpy("dateTimeTxtField.fieldLen"))

	if err != nil {
		return isValid, err
	}

	if len(dateTimeTxtField.dateTimeFormat) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'dateTimeTxtField.dateTimeFormat'\n"+
			"is a zero length string!\n",
			ePrefix.String())

		return isValid, err
	}

	tempTxtLabel := dateTimeTxtField.dateTime.Format(dateTimeTxtField.dateTimeFormat)

	err = txtLabelElectron.isTextJustificationValid(
		[]rune(tempTxtLabel),
		dateTimeTxtField.fieldLen,
		dateTimeTxtField.textJustification,
		ePrefix.XCpy("dateTimeTxtField.textJustification"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecDateTimeAtom.
//
func (txtFieldDateTimeAtom textFieldSpecDateTimeAtom) ptr() *textFieldSpecDateTimeAtom {

	if txtFieldDateTimeAtom.lock == nil {
		txtFieldDateTimeAtom.lock = new(sync.Mutex)
	}

	txtFieldDateTimeAtom.lock.Lock()

	defer txtFieldDateTimeAtom.lock.Unlock()

	return &textFieldSpecDateTimeAtom{
		lock: new(sync.Mutex),
	}
}

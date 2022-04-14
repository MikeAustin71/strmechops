package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecLabelAtom struct {
	lock *sync.Mutex
}

// isValidTextFieldLabel - This method receives a pointer to
// an instance of TextFieldSpecLabel and performs a diagnostic on
// the object to determine if it is valid or not.
//
// This method returns two parameters, 'isValid' and 'err'. If
// input parameter 'txtFieldLabel' is valid, the return parameter
// 'isValid' is set to 'true' and 'err' is set to 'nil'.
//
// If input parameter 'txtFieldLabel' is invalid, the return
// parameter 'isValid' is set to 'false' and 'err' is configured
// with an appropriate error message.
//
func (txtFieldLabelAtom *textFieldSpecLabelAtom) isValidTextFieldLabel(
	txtFieldLabel *TextFieldSpecLabel,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtFieldLabelAtom.lock == nil {
		txtFieldLabelAtom.lock = new(sync.Mutex)
	}

	txtFieldLabelAtom.lock.Lock()

	defer txtFieldLabelAtom.lock.Unlock()

	isValid = false
	err = nil

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecLabelAtom.isValidTextFieldLabel()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtFieldLabel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldLabel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	txtLabelElectron := textFieldSpecLabelElectron{}

	_,
		err = txtLabelElectron.isTextLabelValid(
		txtFieldLabel.textLabel,
		ePrefix.XCpy("txtFieldLabel.textLabel"))

	if err != nil {
		return isValid, err
	}

	err = txtLabelElectron.isFieldLengthValid(
		txtFieldLabel.fieldLen,
		ePrefix.XCpy("txtFieldLabel.fieldLen"))

	if err != nil {
		return isValid, err
	}

	err = txtLabelElectron.isTextJustificationValid(
		txtFieldLabel.textLabel,
		txtFieldLabel.fieldLen,
		txtFieldLabel.textJustification,
		ePrefix.XCpy("txtFieldLabel.textJustification"))

	if err != nil {
		return isValid, err
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecLabelAtom.
//
func (txtFieldLabelAtom textFieldSpecLabelAtom) ptr() *textFieldSpecLabelAtom {

	if txtFieldLabelAtom.lock == nil {
		txtFieldLabelAtom.lock = new(sync.Mutex)
	}

	txtFieldLabelAtom.lock.Lock()

	defer txtFieldLabelAtom.lock.Unlock()

	return &textFieldSpecLabelAtom{
		lock: new(sync.Mutex),
	}
}

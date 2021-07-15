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
// on the object to determine if it is valid or not.
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

	lenTextLabel := len(txtFieldLabel.textLabel)

	if lenTextLabel == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Text Label is a zero length string!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtFieldLabel.fieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Text Label Field Length is less than minus one (-1)!\n"+
			"txtFieldLabel.fieldLen='%v'\n",
			ePrefix.String(),
			txtFieldLabel.fieldLen)

		return isValid, err
	}

	if txtFieldLabel.fieldLen > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Text Label Field Length is greater than one-million (1,000,000)!\n"+
			"txtFieldLabel.fieldLen='%v'\n",
			ePrefix.String(),
			txtFieldLabel.fieldLen)

		return isValid, err
	}

	txtJustificationIsValid := txtFieldLabel.textJustification.XIsValid()

	if txtFieldLabel.fieldLen > lenTextLabel &&
		!txtJustificationIsValid {
		err = fmt.Errorf("%v\n"+
			"Error: Text Justification is INVALID!\n"+
			"Text Lable Length = '%v'\n"+
			"Field Length = '%v'\n"+
			"Text Justification Integer Value = '%v'\n",
			ePrefix.String(),
			lenTextLabel,
			txtFieldLabel.fieldLen,
			txtFieldLabel.textJustification.XValueInt())

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecLabelMolecule.
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

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecFillerAtom struct {
	lock *sync.Mutex
}

// isValidTextFieldSpecFiller - This method receives a pointer to
// an instance of TextFieldSpecFiller and performs a diagnostic on
// on the object to determine if it is valid or not.
//
// This method returns two parameters, 'isValid' and 'err'. If
// input parameter 'txtFieldFiller' is valid, the return parameter
// 'isValid' is set to 'true' and 'err' is set to 'nil'.
//
// If input parameter 'txtFieldFiller' is invalid, the return
// parameter 'isValid' is set to 'false' and 'err' is configured
// with an appropriate error message.
//
func (txtFieldFillerAtom *textFieldSpecFillerAtom) isValidTextFieldSpecFiller(
	txtFieldFiller *TextFieldSpecFiller,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtFieldFillerAtom.lock == nil {
		txtFieldFillerAtom.lock = new(sync.Mutex)
	}

	txtFieldFillerAtom.lock.Lock()

	defer txtFieldFillerAtom.lock.Unlock()

	isValid = false
	err = nil

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecFillerAtom.isValidTextFieldSpecFiller()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtFieldFiller == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFieldFiller' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtFieldFiller.fillerCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: 'incomingTxtFiller.fillerCharsRepeatCount' is "+
			"less than one (+1)!\n"+
			"incomingTxtFiller.fillerCharsRepeatCount='%v'\n",
			ePrefix.String(),
			txtFieldFiller.fillerCharsRepeatCount)

		return isValid, err
	}

	if txtFieldFiller.fillerCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: 'incomingTxtFiller.fillerCharsRepeatCount' is "+
			"less than one (+1)!\n"+
			"incomingTxtFiller.fillerCharsRepeatCount='%v'\n",
			ePrefix.String(),
			txtFieldFiller.fillerCharsRepeatCount)

		return isValid, err
	}

	if len(txtFieldFiller.fillerCharacters) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtFieldFiller.fillerCharacters' is a zero "+
			"length array!\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecFillerAtom.
//
func (txtFieldFillerAtom textFieldSpecFillerAtom) ptr() *textFieldSpecFillerAtom {

	if txtFieldFillerAtom.lock == nil {
		txtFieldFillerAtom.lock = new(sync.Mutex)
	}

	txtFieldFillerAtom.lock.Lock()

	defer txtFieldFillerAtom.lock.Unlock()

	return &textFieldSpecFillerAtom{
		lock: new(sync.Mutex),
	}
}

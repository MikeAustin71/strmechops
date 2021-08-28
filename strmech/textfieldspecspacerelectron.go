package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textFieldSpecSpacerElectron struct {
	lock *sync.Mutex
}

// isFieldLenValidError - This method receives an integer value
// specifying the field length used to configure an instance of
// TextFieldSpecSpacer.
//
// This method returns two parameters, 'isValid' and 'err'. If
// input parameter 'fieldLen' is valid, the return parameter
// 'isValid' is set to 'true' and 'err' is set to 'nil'.
//
// If input parameter 'fieldLen' is invalid, the return
// parameter 'isValid' is set to 'false' and 'err' is configured
// with an appropriate error message.
//
func (txtFieldSpacerElectron *textFieldSpecSpacerElectron) isFieldLenValidError(
	fieldLen int,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtFieldSpacerElectron.lock == nil {
		txtFieldSpacerElectron.lock = new(sync.Mutex)
	}

	txtFieldSpacerElectron.lock.Lock()

	defer txtFieldSpacerElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldSpecSpacerElectron."+
			"isFieldLenValidError()",
		"")

	if err != nil {
		return isValid, err
	}

	if fieldLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter Field Length ('fieldLen') is invalid!\n"+
			"'fieldLen' is less than one (+1).\n"+
			"fieldLen = '%v'\n",
			ePrefix.String(),
			fieldLen)

		return isValid, err
	}

	if fieldLen > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter Field Length ('fieldLen') is invalid!\n"+
			"'fieldLen' is greater than one-million (+1,000,000).\n"+
			"fieldLen = '%v'\n",
			ePrefix.String(),
			fieldLen)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// textFieldSpecSpacerElectron.
//
func (txtFieldSpacerElectron textFieldSpecSpacerElectron) ptr() *textFieldSpecSpacerElectron {

	if txtFieldSpacerElectron.lock == nil {
		txtFieldSpacerElectron.lock = new(sync.Mutex)
	}

	txtFieldSpacerElectron.lock.Lock()

	defer txtFieldSpacerElectron.lock.Unlock()

	return &textFieldSpecSpacerElectron{
		lock: new(sync.Mutex),
	}
}

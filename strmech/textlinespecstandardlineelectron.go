package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecStandardLineElectron struct {
	lock *sync.Mutex
}

// testValidityOfTextLineSpecStdLine - Receives a pointer to an
// instance of TextLineSpecStandardLine and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'txtStdLine' is determined to be invalid,
// this method will return a boolean flag ('isValid') of 'false'.
// In addition, an instance of type error ('err') will be returned
// configured with an appropriate error message.
//
// If the input parameter 'txtStdLine' is valid, this method will
// return a boolean flag ('isValid') of 'true' and the returned
// error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtStdLine                 *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine. This
//       object will be subjected to diagnostic analysis in order
//       to determine if all the member variables contain valid
//       values.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  isValid                    bool
//     - If input parameter 'txtStdLine' is judged to be valid in
//       all respects, this return parameter will be set to 'true'.
//
//     - If input parameter 'txtStdLine' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'txtStdLine' is judged to be valid in
//       all respects, this return parameter will be set to 'nil'.
//
//       If input parameter, 'txtStdLine' is found to be invalid,
//       this return parameter will be configured with an appropriate
//       error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtStdLineElectron *textLineSpecStandardLineElectron) testValidityOfTextLineSpecStdLine(
	txtStdLine *TextLineSpecStandardLine,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecStandardLineElectron.testValidityOfTextLineSpecStdLine()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtStdLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtStdLine.numOfStdLines < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: The number of standard lines is less than one ('1')!\n"+
			"This means that no lines will be generate by this specification.\n",
			ePrefix.String())
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtStdLine' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	lenTextFields := len(txtStdLine.textFields)

	if lenTextFields == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: No Text Fields have been configured for\n"+
			"this standard line specification!\n",
			ePrefix.String())

		return isValid, err
	}

	var err2 error

	for i := 0; i < lenTextFields; i++ {

		err2 = txtStdLine.textFields[i].IsValidInstanceError(ePrefix)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: Text Field Element [%v] is invalid!\n"+
				"Text Field Element Error = \n%v\n",
				ePrefix.String(),
				i,
				err2.Error())

			return isValid, err
		}
	}

	if len(txtStdLine.newLineChars) == 0 {
		txtStdLine.newLineChars = []rune{'\n'}
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecStandardLineElectron.
//
func (txtStdLineElectron textLineSpecStandardLineElectron) ptr() *textLineSpecStandardLineElectron {

	if txtStdLineElectron.lock == nil {
		txtStdLineElectron.lock = new(sync.Mutex)
	}

	txtStdLineElectron.lock.Lock()

	defer txtStdLineElectron.lock.Unlock()

	return &textLineSpecStandardLineElectron{
		lock: new(sync.Mutex),
	}
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecSolidLineAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// textLineSpecSolidLineAtom.
//
func (txtSolidLineAtom textLineSpecSolidLineAtom) ptr() *textLineSpecSolidLineAtom {

	if txtSolidLineAtom.lock == nil {
		txtSolidLineAtom.lock = new(sync.Mutex)
	}

	txtSolidLineAtom.lock.Lock()

	defer txtSolidLineAtom.lock.Unlock()

	return &textLineSpecSolidLineAtom{
		lock: new(sync.Mutex),
	}
}

// testValidityOfTxtSpecTimerLines - Receives a pointer to an
// instance of TextLineSpecSolidLine and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'txtSolidLine' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'txtSolidLine' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtSolidLine               *TextLineSpecSolidLine
//     - A pointer to an instance of TextLineSpecSolidLine. This
//       object will be subjected to diagnostic analysis in order
//       to determine if all the member variables contain valid
//       values.
//
//       If 'txtSolidLine.leftMargin' is less than zero (0), this
//       method will assign a default value of zero (0) to
//       'txtSolidLine.leftMargin'.
//
//       If 'txtSolidLine.newLineChars' is a zero length rune
//       array, or if 'txtSolidLine.newLineChars' contains invalid
//       zero value characters, it will be set to the default new
//       line character ('\n').
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
//     - If input parameter 'txtTimerLines' is judged to be valid
//       in all respects, this return parameter will be set to
//       'true'.
//
//     - If input parameter 'txtTimerLines' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'txtTimerLines' is judged to be valid
//       in all respects, this return parameter will be set to
//       'nil'.
//
//       If input parameter, 'txtTimerLines' is found to be
//       invalid, this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSolidLineAtom *textLineSpecSolidLineAtom) testValidityOfTextSpecSolidLine(
	txtSolidLine *TextLineSpecSolidLine,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtSolidLineAtom.lock == nil {
		txtSolidLineAtom.lock = new(sync.Mutex)
	}

	txtSolidLineAtom.lock.Lock()

	defer txtSolidLineAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecSolidLineAtom.testValidityOfTextSpecSolidLine()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtSolidLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtSolidLine' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		txtSolidLine.solidLineChars,
		ePrefix.XCtx(
			"txtSolidLine.solidLineChars is invalid!"))

	if err != nil {
		return isValid, err
	}

	if txtSolidLine.solidLineCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtSolidLine.solidLineCharsRepeatCount' is invalid!\n"+
			"The value of 'txtSolidLine.solidLineCharsRepeatCount' is "+
			"less than one ('1').\n"+
			"txtSolidLine.solidLineCharsRepeatCount = '%v'.\n",
			ePrefix.XCtxEmpty().String(),
			txtSolidLine.solidLineCharsRepeatCount)

		return isValid, err
	}

	if txtSolidLine.leftMargin < 0 {
		txtSolidLine.leftMargin = 0
	}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		txtSolidLine.newLineChars,
		nil)

	if err != nil {
		txtSolidLine.newLineChars = []rune{'\n'}
		err = nil
	}

	isValid = true

	return isValid, err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecSolidLineAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// TextLineSpecSolidLine and proceeds to set all the internal
// member variables to their uninitialized or zero states.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The data values of all member variables contained in input
// parameter 'txtSolidLine' will be overwritten and deleted.
//
func (txtSolidLineAtom *textLineSpecSolidLineAtom) empty(
	txtSolidLine *TextLineSpecSolidLine) {

	if txtSolidLineAtom.lock == nil {
		txtSolidLineAtom.lock = new(sync.Mutex)
	}

	txtSolidLineAtom.lock.Lock()

	defer txtSolidLineAtom.lock.Unlock()

	if txtSolidLine == nil {
		return
	}

	txtSolidLine.leftMargin = 0
	txtSolidLine.rightMargin = 0
	txtSolidLine.solidLineChars = nil
	txtSolidLine.solidLineCharsRepeatCount = 0
	txtSolidLine.newLineChars = nil
	txtSolidLine.turnLineTerminatorOff = false
	txtSolidLine.textLineReader = nil

	return
}

// equal - Receives pointers to two TextLineSpecSolidLine
// instances and proceeds to compare the member data elements to
// determine whether they are equal.
//
// If the data elements of both input parameters 'txtSolidLineOne'
// and 'txtSolidLineTwo' are equal in all respects, this method
// returns a boolean value of 'true'. Otherwise, this method returns
// 'false'.
//
func (txtSolidLineAtom *textLineSpecSolidLineAtom) equal(
	txtSolidLineOne *TextLineSpecSolidLine,
	txtSolidLineTwo *TextLineSpecSolidLine) bool {

	if txtSolidLineAtom.lock == nil {
		txtSolidLineAtom.lock = new(sync.Mutex)
	}

	txtSolidLineAtom.lock.Lock()

	defer txtSolidLineAtom.lock.Unlock()

	if txtSolidLineOne == nil ||
		txtSolidLineTwo == nil {
		return false
	}

	if txtSolidLineOne.leftMargin !=
		txtSolidLineTwo.leftMargin {
		return false
	}

	if txtSolidLineOne.rightMargin !=
		txtSolidLineTwo.rightMargin {
		return false
	}

	if txtSolidLineOne.turnLineTerminatorOff !=
		txtSolidLineTwo.turnLineTerminatorOff {
		return false
	}

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		txtSolidLineOne.solidLineChars,
		txtSolidLineTwo.solidLineChars) {
		return false
	}

	if txtSolidLineOne.solidLineCharsRepeatCount !=
		txtSolidLineTwo.solidLineCharsRepeatCount {
		return false
	}

	// nil arrays and zero length arrays
	// are treated as 'equal'
	if !sMechPreon.equalRuneArrays(
		txtSolidLineOne.newLineChars,
		txtSolidLineTwo.newLineChars) {
		return false
	}

	return true
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

	if len(txtSolidLine.solidLineChars) == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"'solidLineChars' is empty. Zero array length.\n",
			ePrefix.String())

		return isValid, err
	}

	var err2 error

	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		txtSolidLine.solidLineChars,
		ePrefix.XCpy(
			"txtSolidLine.solidLineChars is invalid!"))

	if err2 != nil {

		err = fmt.Errorf("%v - ERROR\n"+
			"txtSolidLine.solidLineChars contains invalid runes!\n"+
			"%v\n"+
			"txtSolidLine.solidLineChars='%v'\n",
			ePrefix.String(),
			err2.Error(),
			txtSolidLine.solidLineChars)

		return isValid, err
	}

	if txtSolidLine.solidLineCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtSolidLine.solidLineCharsRepeatCount' is invalid!\n"+
			"The value of 'txtSolidLine.solidLineCharsRepeatCount' is "+
			"less than one ('1').\n"+
			"txtSolidLine.solidLineCharsRepeatCount = '%v'.\n",
			ePrefix.String(),
			txtSolidLine.solidLineCharsRepeatCount)

		return isValid, err
	}

	if txtSolidLine.leftMargin < 0 {
		txtSolidLine.leftMargin = 0
	}

	if txtSolidLine.leftMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtSolidLine.leftMargin' is invalid!\n"+
			"The integer value of 'txtSolidLine.leftMargin' is greater than 1,000,000.\n"+
			"txtSolidLine.leftMargin='%v'\n",
			ePrefix.String(),
			txtSolidLine.leftMargin)

		return isValid, err
	}

	if txtSolidLine.rightMargin < 0 {
		txtSolidLine.rightMargin = 0
	}

	if txtSolidLine.rightMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: 'txtSolidLine.rightMargin' is invalid!\n"+
			"The integer value of 'txtSolidLine.rightMargin' is greater than 1,000,000.\n"+
			"txtSolidLine.rightMargin='%v'\n",
			ePrefix.String(),
			txtSolidLine.rightMargin)

		return isValid, err
	}

	lenNewLineChars := len(txtSolidLine.newLineChars)

	if lenNewLineChars == 0 {

		txtSolidLine.newLineChars = []rune{'\n'}

	} else if lenNewLineChars > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: 'txtSolidLine.newLineChars' is invalid!\n"+
			"The length of 'txtSolidLine.newLineChars' is greater than one-million (1,000,000).\n"+
			"txtSolidLine.newLineChars length ='%v'\n",
			ePrefix.String(),
			lenNewLineChars)

		return isValid, err

	} else {
		_,
			err2 = sMechPreon.testValidityOfRuneCharArray(
			txtSolidLine.newLineChars,
			nil)

		if err2 != nil {

			err = fmt.Errorf("%v\n"+
				"Error: 'txtSolidLine.newLineChars' is invalid!\n"+
				"txtSolidLine.newLineChars='%v'\n"+
				"Error Msg: %v\n",
				ePrefix.String(),
				txtSolidLine.rightMargin,
				err2.Error())

			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}

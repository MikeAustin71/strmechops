package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecPlainTextAtom struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// textLineSpecPlainTextAtom.
//
func (txtLinePlainTextAtom textLineSpecPlainTextAtom) ptr() *textLineSpecPlainTextAtom {

	if txtLinePlainTextAtom.lock == nil {
		txtLinePlainTextAtom.lock = new(sync.Mutex)
	}

	txtLinePlainTextAtom.lock.Lock()

	defer txtLinePlainTextAtom.lock.Unlock()

	return &textLineSpecPlainTextAtom{
		lock: new(sync.Mutex),
	}
}

// testValidityOfTextLineSpecPlainText - Receives a pointer to an
// instance of TextLineSpecPlainText and performs a diagnostic
// analysis to determine if that instance is valid in all respects.
//
// If the input parameter 'plainTextLine' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'plainTextLine' is valid, this method
// will return a boolean flag ('isValid') of 'true' and the
// returned error type ('err') will be set to 'nil'.
//
// If plainTextLine.newLineChars is a zero length array, this method
// will automatically set this value to the default new line
// character ('\n').
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  plainTextLine              *TextLineSpecPlainText
//     - A pointer to an instance of TextLineSpecPlainText. This
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
//     - If input parameter 'plainTextLine' is judged to be valid
//       in all respects, this return parameter will be set to
//       'true'.
//
//     - If input parameter 'plainTextLine' is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If input parameter 'plainTextLine' is judged to be valid
//       in all respects, this return parameter will be set to
//       'nil'.
//
//       If input parameter, 'plainTextLine' is found to be
//       invalid, this return parameter will be configured with an
//       appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtLinePlainTextAtom *textLineSpecPlainTextAtom) testValidityOfTextLineSpecPlainText(
	plainTextLine *TextLineSpecPlainText,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtLinePlainTextAtom.lock == nil {
		txtLinePlainTextAtom.lock = new(sync.Mutex)
	}

	txtLinePlainTextAtom.lock.Lock()

	defer txtLinePlainTextAtom.lock.Unlock()

	isValid = false

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextAtom."+
			"testValidityOfTextLineSpecPlainText()",
		"")

	if err != nil {
		return isValid, err
	}

	if plainTextLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'plainTextLine' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if len(plainTextLine.textString) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: plainTextLine.textString is an empty string!\n"+
			"No Text String have been configured for\n"+
			"this Plain Text Line Specification!\n",
			ePrefix.String())

		return isValid, err
	}

	if len(plainTextLine.newLineChars) == 0 {

		err =
			textSpecificationMolecule{}.ptr().
				setDefaultNewLineChars(
					&plainTextLine.newLineChars,
					ePrefix.XCtx(
						"plainTextLine.newLineChars"))

		if err != nil {
			return isValid, err
		}

	}

	sMechPreon := strMechPreon{}

	if len(plainTextLine.leftMarginChars) > 0 {
		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				plainTextLine.leftMarginChars,
				ePrefix.XCtx(
					"plainTextLine.leftMarginChars invalid!"))

		if err != nil {
			return isValid, err
		}

	}

	if len(plainTextLine.rightMarginChars) > 0 {
		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				plainTextLine.rightMarginChars,
				ePrefix.XCtx(
					"plainTextLine.rightMarginChars invalid!"))

		if err != nil {
			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}

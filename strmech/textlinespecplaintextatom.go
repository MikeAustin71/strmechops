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

// setPlainTextSpec - Receives a pointer to an instance of
// TextLineSpecPlainText and proceeds to reset all the member
// variable data values using the values passed through input
// parameters.
//
// The TextLineSpecPlainText type provides formatting
// specifications for a simple line of text characters for text
// display, file output or printing.
//
// The plain text line consists of a left margin, the text string,
// the right margin and a line termination character or characters.
// The line termination character is usually a new line character
// ('\n').
//
// Left and right margins consist of zero or more characters
// customized and provided by the calling function.
//
// This method is similar to
// textLineSpecPlainTextNanobot.setPlainTextSpecRunes() with the
// sole exception being that this method receives input parameter
// 'textString' as a string instead of an array of runes.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leftMarginRunes            []rune
//     - An array of runes containing the text characters which
//       will be used to construct the left margin of the plain
//       text line generated by the current instance of
//       TextLineSpecPlainText.
//
//       If this parameter is submitted as a zero length or empty
//       array, NO error will be generated and the existing left
//       margin will be effectively eliminated from the plain text
//       line generated by this instance of TextLineSpecPlainText.
//
//       If the 'leftMarginRunes' array exceeds a length of
//       one-million array elements, an error will be returned. If
//       any of the array elements has a rune value of zero ('0'),
//       an error will be returned.
//
//
//  rightMarginRunes           []rune
//     - An array of runes containing the text characters which
//       will be used to construct the right margin of the plain
//       text line generated by the current instance of
//       TextLineSpecPlainText.
//
//       If this parameter is submitted as a zero length or empty
//       array, NO error will be generated and the existing left
//       margin will be effectively eliminated from the plain text
//       line generated by this instance of TextLineSpecPlainText.
//
//       If the 'rightMarginRunes' array exceeds a length of
//       one-million array elements, an error will be returned. If
//       any of the array elements has a rune value of zero ('0'),
//       an error will be returned.
//
//
//  textString                 string
//     - A string of text which will be used to configure the text
//       characters generated by the current instance of
//       TextLineSpecPlainText.
//
//       If this parameter is submitted as a zero length or empty
//       string, an error will be returned.
//
//       If this string contains more than one-million characters,
//       an error will be returned.
//
//
//  newLineChars               []rune
//     - An array of runes containing the character or characters
//       used to terminate each line of text generated by the
//       current instance of TextLineSpecPlainText.
//
//       If this parameter is submitted as a zero length or empty
//       rune array, it will be reset to the default new line value
//       ('\n').
//
//       If the rune array contains invalid zero rune values, an
//       error will be returned.
//
//
//  turnLineTerminatorOff      bool
//     - The 'turnLineTerminatorOff' flag controls whether a line
//       termination character or characters will be automatically
//       appended to each line of text produced by
//       TextLineSpecPlainText.
//
//       When the boolean flag 'turnLineTerminatorOff' is set to
//       'false', line terminators as defined by parameter
//       'newLineChars' WILL BE applied as a line termination
//       sequence for each line of text produced by
//       TextLineSpecPlainText.
//
//       When this boolean value is set to 'true', it turns off or
//       cancels the automatic generation of line terminators for
//       each line of text produced by TextLineSpecStandardLine.
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
func (txtLinePlainTextAtom *textLineSpecPlainTextAtom) setPlainTextSpec(
	plainTxtLine *TextLineSpecPlainText,
	leftMarginChars []rune,
	rightMarginChars []rune,
	textString string,
	newLineChars []rune,
	turnLineTerminatorOff bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLinePlainTextAtom.lock == nil {
		txtLinePlainTextAtom.lock = new(sync.Mutex)
	}

	txtLinePlainTextAtom.lock.Lock()

	defer txtLinePlainTextAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextAtom."+
			"setPlainTextSpec()",
		"")

	if err != nil {
		return err
	}

	if plainTxtLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'plainTxtLine' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(textString) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textString' is an empty or zero length string!\n",
			ePrefix.String())

		return err
	}

	lenLeftMargin := len(leftMarginChars)

	if lenLeftMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: The 'leftMarginsChars' rune array exceeds\n"+
			"one-million (1,000,000) characters in length.\n"+
			"Length of 'leftMarginsChars' = '%v'\n",
			ePrefix.String(),
			lenLeftMargin)

		return err
	}

	lenRightMargin := len(rightMarginChars)

	if lenRightMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: The 'rightMarginChars' rune array exceeds\n"+
			"one-million (1,000,000) characters in length.\n"+
			"Length of 'rightMarginChars' = '%v'\n",
			ePrefix.String(),
			lenRightMargin)

		return err
	}

	sMechPreon := strMechPreon{}

	if len(leftMarginChars) > 0 {
		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				leftMarginChars,
				ePrefix.XCtx(
					"input parameter leftMarginChars invalid!"))

		if err != nil {
			return err
		}

		err = sMechPreon.copyRuneArrays(
			&plainTxtLine.leftMarginChars,
			&leftMarginChars,
			true,
			ePrefix.XCtx(
				"leftMarginChars->"+
					"plainTxtLine.leftMarginChars"))

		if err != nil {
			return err
		}

	} else {
		// len(leftMarginChars) == 0
		plainTxtLine.leftMarginChars = nil
	}

	if len(rightMarginChars) > 0 {

		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				rightMarginChars,
				ePrefix.XCtx(
					"input parameter rightMarginChars invalid!"))

		if err != nil {
			return err
		}

		err = sMechPreon.copyRuneArrays(
			&plainTxtLine.rightMarginChars,
			&rightMarginChars,
			true,
			ePrefix.XCtx(
				"rightMarginChars->"+
					"plainTxtLine.rightMarginChars"))

		if err != nil {
			return err
		}

	} else {
		// len(rightMarginChars) == 0

		plainTxtLine.rightMarginChars = nil

	}

	if len(newLineChars) > 0 {

		_,
			err =
			sMechPreon.testValidityOfRuneCharArray(
				newLineChars,
				ePrefix.XCtx(
					"input parameter newLineChars invalid!"))

		if err != nil {
			return err
		}

		err = sMechPreon.copyRuneArrays(
			&plainTxtLine.newLineChars,
			&newLineChars,
			true,
			ePrefix.XCtx(
				"newLineChars->"+
					"plainTxtLine.newLineChars"))

	} else {
		// len(newLineChars) == 0
		newLineChars = []rune{'\n'}

	}

	plainTxtLine.textString =
		textString

	plainTxtLine.turnLineTerminatorOff =
		turnLineTerminatorOff

	return err
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

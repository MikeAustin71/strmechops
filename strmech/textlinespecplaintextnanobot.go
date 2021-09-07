package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecPlainTextNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingPlainTextLine' to input parameter
// 'targetPlainTextLine'.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// The pre-existing data fields for input parameter 'targetStdLine'
// will be overwritten and deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetPlainTextLine        *TextLineSpecPlainText
//     - A pointer to an instance of TextLineSpecPlainText. Data
//       extracted from input parameter 'incomingPlainTextLine'
//       will be copied to this input parameter,
//       'targetPlainTextLine'. If this method completes
//       successfully, all member data variables encapsulated in
//       'targetPlainTextLine' will be identical to those contained
//       in input parameter, 'incomingPlainTextLine'.
//
//       Be advised that the pre-existing data fields in input
//       parameter 'targetPlainTextLine' will be overwritten and
//       deleted.
//
//
//  incomingPlainTextLine      *TextLineSpecPlainText
//     - A pointer to an instance of TextLineSpecPlainText.
//
//       All data values in this TextLineSpecPlainText instance
//       will be copied to input parameter 'targetPlainTextLine'.
//
//       The original member variable data values encapsulated in
//       'incomingPlainTextLine' will remain unchanged and will NOT
//       be overwritten or deleted.
//
//       If 'incomingPlainTextLine' contains invalid member
//       variable data values, this method will return an error.
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
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtLinePlainTextNanobot *textLineSpecPlainTextNanobot) copyIn(
	targetPlainTextLine *TextLineSpecPlainText,
	incomingPlainTextLine *TextLineSpecPlainText,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLinePlainTextNanobot.lock == nil {
		txtLinePlainTextNanobot.lock = new(sync.Mutex)
	}

	txtLinePlainTextNanobot.lock.Lock()

	defer txtLinePlainTextNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextNanobot."+
			"copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetPlainTextLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetPlainTextLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingPlainTextLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingPlainTextLine' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textLineSpecPlainTextAtom{}.ptr().
		testValidityOfTextLineSpecPlainText(
			incomingPlainTextLine,
			ePrefix.XCtx(
				"incomingPlainTextLine"))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&targetPlainTextLine.leftMarginChars,
		&incomingPlainTextLine.leftMarginChars,
		true,
		ePrefix.XCtx(
			"incomingPlainTextLine.leftMarginChars->"+
				"targetPlainTextLine.leftMarginChars"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&targetPlainTextLine.rightMarginChars,
		&incomingPlainTextLine.rightMarginChars,
		true,
		ePrefix.XCtx(
			"incomingPlainTextLine.rightMarginChars->"+
				"targetPlainTextLine.rightMarginChars"))

	if err != nil {
		return err
	}

	targetPlainTextLine.textString =
		incomingPlainTextLine.textString

	targetPlainTextLine.turnLineTerminatorOff =
		incomingPlainTextLine.turnLineTerminatorOff

	err = sMechPreon.copyRuneArrays(
		&targetPlainTextLine.newLineChars,
		&incomingPlainTextLine.newLineChars,
		true,
		ePrefix.XCtx(
			"incomingPlainTextLine.newLineChars->"+
				"targetPlainTextLine.newLineChars"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'plainTxtLine'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  plainTxtLine          *TextLineSpecPlainText
//     - A pointer to an instance of TextLineSpecPlainText. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of TextLineSpecPlainText.
//
//       If the member variable data values encapsulated by
//       'plainTxtLine' are found to be invalid, this method will
//       return an error
//
//
//  errPrefDto          *ePref.ErrPrefixDto
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
//  TextLineSpecPlainText
//     - If this method completes successfully, a deep copy of
//       input parameter 'plainTxtLine' will be created and returned
//       in a new instance of TextLineSpecPlainText.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtLinePlainTextNanobot *textLineSpecPlainTextNanobot) copyOut(
	plainTxtLine *TextLineSpecPlainText,
	errPrefDto *ePref.ErrPrefixDto) (
	TextLineSpecPlainText, error) {

	if txtLinePlainTextNanobot.lock == nil {
		txtLinePlainTextNanobot.lock = new(sync.Mutex)
	}

	txtLinePlainTextNanobot.lock.Lock()

	defer txtLinePlainTextNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextNanobot.copyOut()",
		"")

	if err != nil {
		return TextLineSpecPlainText{}, err
	}

	if plainTxtLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'plainTxtLine' is a nil pointer!\n",
			ePrefix.String())

		return TextLineSpecPlainText{}, err
	}

	_,
		err = textLineSpecPlainTextAtom{}.ptr().
		testValidityOfTextLineSpecPlainText(
			plainTxtLine,
			ePrefix.XCtx(
				"plainTxtLine"))

	if err != nil {
		return TextLineSpecPlainText{}, err
	}

	newPlainTxtLine := TextLineSpecPlainText{}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&newPlainTxtLine.leftMarginChars,
		&plainTxtLine.leftMarginChars,
		true,
		ePrefix.XCtx(
			"plainTxtLine.leftMarginChars->"+
				"newPlainTxtLine.leftMarginChars"))

	if err != nil {
		return TextLineSpecPlainText{}, err
	}

	err = sMechPreon.copyRuneArrays(
		&newPlainTxtLine.rightMarginChars,
		&plainTxtLine.rightMarginChars,
		true,
		ePrefix.XCtx(
			"plainTxtLine.rightMarginChars->"+
				"newPlainTxtLine.rightMarginChars"))

	if err != nil {
		return TextLineSpecPlainText{}, err
	}

	newPlainTxtLine.textString =
		plainTxtLine.textString

	newPlainTxtLine.turnLineTerminatorOff =
		plainTxtLine.turnLineTerminatorOff

	err = sMechPreon.copyRuneArrays(
		&newPlainTxtLine.newLineChars,
		&plainTxtLine.newLineChars,
		true,
		ePrefix.XCtx(
			"plainTxtLine.newLineChars->"+
				"newPlainTxtLine.newLineChars"))

	return newPlainTxtLine, err
}

// getFormattedText - Returns the formatted text generated by this
// Text Line Specification, 'plainTxtLine', for output and printing.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  plainTxtLine               *TextLineSpecPlainText
//     - A pointer to an instance of TextLineSpecPlainText. The
//       member variables encapsulated by this instance will be
//       used to generate formatted text for text display, file
//       output and printing.
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
//  formattedText              string
//     - If this method completes successfully, a string of
//       formatted text will be generated from the data provided by
//       input parameter 'plainTxtLine'.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtLinePlainTextNanobot *textLineSpecPlainTextNanobot) getFormattedText(
	plainTxtLine *TextLineSpecPlainText,
	errPrefDto *ePref.ErrPrefixDto) (
	formattedText string,
	err error) {

	if txtLinePlainTextNanobot.lock == nil {
		txtLinePlainTextNanobot.lock = new(sync.Mutex)
	}

	txtLinePlainTextNanobot.lock.Lock()

	defer txtLinePlainTextNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	formattedText = ""

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextNanobot."+
			"getFormattedText()",
		"")

	if err != nil {
		return formattedText, err
	}

	if plainTxtLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'plainTxtLine' is a nil pointer!\n",
			ePrefix.String())

		return formattedText, err
	}

	_,
		err = textLineSpecPlainTextAtom{}.ptr().
		testValidityOfTextLineSpecPlainText(
			plainTxtLine,
			ePrefix.XCtx(
				"plainTxtLine"))

	if err != nil {
		return formattedText, err
	}

	if len(plainTxtLine.leftMarginChars) > 0 {
		formattedText += string(plainTxtLine.leftMarginChars)
	}

	formattedText += plainTxtLine.textString

	if len(plainTxtLine.rightMarginChars) > 0 {
		formattedText += string(plainTxtLine.rightMarginChars)
	}

	if plainTxtLine.turnLineTerminatorOff == true {
		return formattedText, err
	}

	formattedText += string(plainTxtLine.newLineChars)

	return formattedText, err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecPlainTextNanobot.
//
func (txtLinePlainTextNanobot textLineSpecPlainTextNanobot) ptr() *textLineSpecPlainTextNanobot {

	if txtLinePlainTextNanobot.lock == nil {
		txtLinePlainTextNanobot.lock = new(sync.Mutex)
	}

	txtLinePlainTextNanobot.lock.Lock()

	defer txtLinePlainTextNanobot.lock.Unlock()

	return &textLineSpecPlainTextNanobot{
		lock: new(sync.Mutex),
	}
}

// setPlainTextSpecRunes - Receives a pointer to an instance of
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
// textLineSpecPlainTextAtom.setPlainTextSpec() with the sole
// exception being that this method receives input parameter
// 'textRunes as an array of runes instead of a string.
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
//  textRunes                []rune
//     - An array of runes which will be used to configure the text
//       characters generated by the current instance of
//       TextLineSpecPlainText.
//
//       If this parameter is submitted as a zero length or empty
//       array, an error will be returned.
//
//       If this array contains more than one-million characters,
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
func (txtLinePlainTextNanobot *textLineSpecPlainTextNanobot) setPlainTextSpecRunes(
	plainTxtLine *TextLineSpecPlainText,
	leftMarginChars []rune,
	rightMarginChars []rune,
	textRunes []rune,
	newLineChars []rune,
	turnLineTerminatorOff bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLinePlainTextNanobot.lock == nil {
		txtLinePlainTextNanobot.lock = new(sync.Mutex)
	}

	txtLinePlainTextNanobot.lock.Lock()

	defer txtLinePlainTextNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecPlainTextNanobot."+
			"setPlainTextSpecRunes()",
		"")

	lenTextRunes := len(textRunes)

	if lenTextRunes > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textRunes' string exceeds\n"+
			"one-million (1,000,000) characters in length.\n"+
			"Length of 'textRunes' = '%v'\n",
			ePrefix.String(),
			lenTextRunes)

		return err
	}

	if lenTextRunes == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textRunes' is empty and\n"+
			"contains zero characters.\n",
			ePrefix.String())

		return err
	}

	_,
		err = strMechPreon{}.ptr().
		testValidityOfRuneCharArray(
			textRunes,
			ePrefix.XCtx(
				"textRunes invalid!"))

	if err != nil {
		return err
	}

	textString := string(textRunes)

	return textLineSpecPlainTextAtom{}.ptr().
		setPlainTextSpec(
			plainTxtLine,
			leftMarginChars,
			rightMarginChars,
			textString,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix)
}

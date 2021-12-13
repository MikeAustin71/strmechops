package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type textLineSpecSolidLineMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingTxtSolidLine' to input parameter 'targetTxtSolidLine'.
//
// Be advised that the pre-existing data fields in input parameter
// 'targetTxtSolidLine' will be overwritten and deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetTxtSolidLine         *TextLineSpecSolidLine
//     - A pointer to an instance of TextLineSpecSolidLine. Data
//       extracted from input parameter 'incomingTxtSolidLine' will
//       be copied to this input parameter, 'targetTxtSolidLine'.
//       If this method completes successfully, all member data
//       variables encapsulated in 'targetTxtSolidLine' will be
//       identical to those contained in input parameter,
//       'incomingTxtSolidLine'.
//
//       Be advised that the pre-existing data fields in input
//       parameter 'targetTxtSolidLine' will be overwritten and
//       deleted.
//
//
//  incomingTxtSolidLine       *TextLineSpecSolidLine
//     - A pointer to an instance of TextLineSpecSolidLine.
//
//       All data values in this TextLineSpecSolidLine instance
//       will be copied to input parameter 'targetTxtSolidLine'.
//
//       The original member variable data values encapsulated in
//       'incomingTxtSolidLine' will remain unchanged and will NOT be
//       overwritten or deleted.
//
//       If 'incomingTxtSolidLine' contains invalid member data
//       variables, this method will return an error.
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
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSolidLineMolecule *textLineSpecSolidLineMolecule) copyIn(
	targetTxtSolidLine *TextLineSpecSolidLine,
	incomingTxtSolidLine *TextLineSpecSolidLine,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtSolidLineMolecule.lock == nil {
		txtSolidLineMolecule.lock = new(sync.Mutex)
	}

	txtSolidLineMolecule.lock.Lock()

	defer txtSolidLineMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecSolidLineMolecule."+
			"copyIn()",
		"")

	if err != nil {
		return err
	}

	if incomingTxtSolidLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingTxtSolidLine' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if targetTxtSolidLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetTxtSolidLine' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = textLineSpecSolidLineAtom{}.ptr().
		testValidityOfTextSpecSolidLine(
			incomingTxtSolidLine,
			ePrefix.XCtx(
				"'incomingTxtSolidLine' is invalid!"))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&targetTxtSolidLine.solidLineChars,
		&incomingTxtSolidLine.solidLineChars,
		true,
		ePrefix.XCtx(
			"incomingTxtSolidLine.solidLineChars->"+
				"targetTxtSolidLine.solidLineChars"))

	targetTxtSolidLine.leftMargin =
		incomingTxtSolidLine.leftMargin

	targetTxtSolidLine.rightMargin =
		incomingTxtSolidLine.rightMargin

	targetTxtSolidLine.solidLineCharsRepeatCount =
		incomingTxtSolidLine.solidLineCharsRepeatCount

	targetTxtSolidLine.turnLineTerminatorOff =
		incomingTxtSolidLine.turnLineTerminatorOff

	targetTxtSolidLine.textLineReader = nil

	err = sMechPreon.copyRuneArrays(
		&targetTxtSolidLine.newLineChars,
		&incomingTxtSolidLine.newLineChars,
		true,
		ePrefix.XCtx(
			"targetTxtSolidLine.newLineChars->"+
				"incomingTxtSolidLine.newLineChars"))

	return err
}

// copyOut - Returns a deep copy of the TextLineSpecSolidLine
// input parameter 'txtSolidLine'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtSolidLine               *TextLineSpecSolidLine
//     - A pointer to an instance of TextLineSpecSolidLine. A deep
//       copy of the internal member variables will be created
//       and returned in a new instance of TextLineSpecSolidLine.
//
//       If the member variable data values encapsulated by this
//       'txtSolidLine' parameter are found to be invalid, an error
//       will be returned.
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
//  TextLineSpecSolidLine
//     - If this method completes successfully, a deep copy of
//       input parameter 'txtSolidLine' will be created and
//       returned in a new instance of TextLineSpecSolidLine.
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
func (txtSolidLineMolecule *textLineSpecSolidLineMolecule) copyOut(
	txtSolidLine *TextLineSpecSolidLine,
	errPrefDto *ePref.ErrPrefixDto) (
	newTxtSolidLine TextLineSpecSolidLine,
	err error) {

	if txtSolidLineMolecule.lock == nil {
		txtSolidLineMolecule.lock = new(sync.Mutex)
	}

	txtSolidLineMolecule.lock.Lock()

	defer txtSolidLineMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecSolidLineMolecule."+
			"copyOut()",
		"")

	if err != nil {
		return newTxtSolidLine, err
	}

	if txtSolidLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtSolidLine' is a nil pointer!\n",
			ePrefix.String())

		return newTxtSolidLine, err
	}

	_,
		err = textLineSpecSolidLineAtom{}.ptr().
		testValidityOfTextSpecSolidLine(
			txtSolidLine,
			ePrefix.XCtx(
				"'txtSolidLine' is invalid!"))

	if err != nil {
		return newTxtSolidLine, err
	}

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&newTxtSolidLine.solidLineChars,
		&txtSolidLine.solidLineChars,
		true,
		ePrefix.XCtx(
			"txtSolidLine.solidLineChars->"+
				"newTxtSolidLine.solidLineChars"))

	newTxtSolidLine.leftMargin =
		txtSolidLine.leftMargin

	newTxtSolidLine.rightMargin =
		txtSolidLine.rightMargin

	newTxtSolidLine.solidLineCharsRepeatCount =
		txtSolidLine.solidLineCharsRepeatCount

	newTxtSolidLine.turnLineTerminatorOff =
		txtSolidLine.turnLineTerminatorOff

	err = sMechPreon.copyRuneArrays(
		&newTxtSolidLine.newLineChars,
		&txtSolidLine.newLineChars,
		true,
		ePrefix.XCtx(
			"txtSolidLine.newLineChars->"+
				"newTxtSolidLine.newLineChars"))

	return newTxtSolidLine, err
}

// getFormattedText - Returns the formatted text generated by an
// instance of TextLineSpecSolidLine passed as an input parameter.
//
// Input parameter 'txtSolidLine' will provide the necessary
// information for generating a solid line of text for text output
// display or printing.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtSolidLine               *TextLineSpecSolidLine
//     - A pointer to an instance of TextLineSpecSolidLine. The
//       member variables encapsulated by this instance will be
//       used to generate a solid line of text for text output
//       display or printing.
//
//       If the member variable data values encapsulated by this
//       TextLineSpecSolidLine instance are found to be invalid,
//       this method will return an error.
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
//       input parameter 'txtSolidLine'.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSolidLineMolecule *textLineSpecSolidLineMolecule) getFormattedText(
	txtSolidLine *TextLineSpecSolidLine,
	errPrefDto *ePref.ErrPrefixDto) (
	formattedText string,
	err error) {

	if txtSolidLineMolecule.lock == nil {
		txtSolidLineMolecule.lock = new(sync.Mutex)
	}

	txtSolidLineMolecule.lock.Lock()

	defer txtSolidLineMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecSolidLineMolecule."+
			"getFormattedText()",
		"")

	if err != nil {
		return formattedText, err
	}

	if txtSolidLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtSolidLine' is a nil pointer!\n",
			ePrefix.String())

		return formattedText, err
	}

	_,
		err = textLineSpecSolidLineAtom{}.ptr().
		testValidityOfTextSpecSolidLine(
			txtSolidLine,
			ePrefix.XCtx(
				"'txtSolidLine' is invalid!"))

	sb := strings.Builder{}

	sb.Grow(256)

	for i := 0; i < txtSolidLine.leftMargin; i++ {
		sb.WriteString(" ")
	}

	str := string(txtSolidLine.solidLineChars)

	for i := 0; i < txtSolidLine.solidLineCharsRepeatCount; i++ {
		sb.WriteString(str)
	}

	for i := 0; i < txtSolidLine.rightMargin; i++ {
		sb.WriteString(" ")
	}

	if !txtSolidLine.turnLineTerminatorOff {
		sb.WriteString(string(txtSolidLine.newLineChars))
	}

	formattedText = sb.String()

	return formattedText, err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecSolidLineMolecule.
//
func (txtSolidLineMolecule textLineSpecSolidLineMolecule) ptr() *textLineSpecSolidLineMolecule {

	if txtSolidLineMolecule.lock == nil {
		txtSolidLineMolecule.lock = new(sync.Mutex)
	}

	txtSolidLineMolecule.lock.Lock()

	defer txtSolidLineMolecule.lock.Unlock()

	return &textLineSpecSolidLineMolecule{
		lock: new(sync.Mutex),
	}
}

// setTxtSolidLine - Sets the member variable data values for an
// instance of TextLineSpecSolidLine passed as input parameter
// 'txtSolidLine'.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display and
// printing. A solid line consists of a single character or
// multiple characters which constitute a solid line and are often
// used for line breaks. Typically, solid lines consist of dashes
// ("---"), underscore characters ("____"), equal signs ("=====")
// or asterisks ("*****"). Multiple characters may be used to
// produce different line sequences ("--*--*--*"). The length of a
// solid is specified by the calling function using input
// parameter 'solidLineCharsRepeatCount'.
//
// By default, all solid lines of text are terminated with a new
// line character ('\n). This means that each instance of a solid
// will be formatted as a single line of text. The new line
// character or characters may be customized by the calling
// function.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtSolidLine               *TextLineSpecSolidLine
//     - A pointer to an instance of TextLineSpecSolidLine. The
//       member variables data values encapsulated in this object
//       will be set to new values extracted from the following
//       input parameters.
//
//
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0). If this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  rightMargin                 int
//     - The number of white space characters appended to the
//       end, or right side, of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0). If this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  solidLineChars             []rune
//     - The array of runes specifies the character or characters
//       which will comprise the solid line output for text display
//       or printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
///
//       If this parameter is submitted as a zero length rune
//       array, an error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//
//  newLineChars               []rune
//     - This rune array contains one or more characters which will
//       be used to terminate the solid text line.
//
//       Example:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         newLineChars = []rune{'??\n')
//         Solid line = "*****??\n"
//
//       If this parameter is submitted as a zero length rune
//       array, or if 'newLineChars' contains invalid zero value
//       characters, it will be set to the default new line
//       character ('\n').
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
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSolidLineMolecule *textLineSpecSolidLineMolecule) setTxtSolidLine(
	txtSolidLine *TextLineSpecSolidLine,
	leftMargin int,
	rightMargin int,
	solidLineChars []rune,
	solidLineCharsRepeatCount int,
	newLineChars []rune,
	turnLineTerminatorOff bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtSolidLineMolecule.lock == nil {
		txtSolidLineMolecule.lock = new(sync.Mutex)
	}

	txtSolidLineMolecule.lock.Lock()

	defer txtSolidLineMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecSolidLineMolecule."+
			"setTxtSolidLine()",
		"")

	if err != nil {
		return err
	}

	if txtSolidLine == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtSolidLine' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(solidLineChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineChars' is empty (zero length)!\n",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		solidLineChars,
		ePrefix.XCtx(
			"Error: Input parameter 'solidLineChars'"+
				" is invalid!"))

	if err != nil {
		return err
	}

	if solidLineCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"The value of 'solidLineCharsRepeatCount' is less than one ('1').\n"+
			"solidLineCharsRepeatCount = '%v'.\n",
			ePrefix.String(),
			solidLineCharsRepeatCount)

		return err
	}

	if leftMargin < 0 {
		leftMargin = 0
	}

	if leftMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leftMargin' is invalid!\n"+
			"The integer value of 'leftMargin' is greater than 1,000,000.\n"+
			"leftMargin='%v'\n",
			ePrefix.String(),
			leftMargin)

		return err
	}

	if rightMargin < 0 {
		rightMargin = 0
	}

	if rightMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'rightMargin' is invalid!\n"+
			"The integer value of 'rightMargin' is greater than 1,000,000.\n"+
			"leftMargin='%v'\n",
			ePrefix.String(),
			leftMargin)

		return err
	}

	lenNewLineChars := len(newLineChars)

	if lenNewLineChars == 0 ||
		lenNewLineChars > 1000000 {
		newLineChars = []rune{'\n'}
	}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		newLineChars,
		ePrefix.XCtx(
			"Testing Validity of 'newLineChars'"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&txtSolidLine.solidLineChars,
		&solidLineChars,
		true,
		ePrefix.XCtx(
			"solidLineChars->"+
				"txtSolidLine.solidLineChars"))

	if err != nil {
		return err
	}

	txtSolidLine.leftMargin = leftMargin

	txtSolidLine.rightMargin = rightMargin

	txtSolidLine.solidLineCharsRepeatCount =
		solidLineCharsRepeatCount

	txtSolidLine.turnLineTerminatorOff =
		turnLineTerminatorOff

	err = sMechPreon.copyRuneArrays(
		&txtSolidLine.newLineChars,
		&newLineChars,
		true,
		ePrefix.XCtx(
			"newLineChars->"+
				"txtSolidLine.newLineChars"))

	return err
}

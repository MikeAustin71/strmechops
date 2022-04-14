package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecBlankLinesMolecule struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter 'incomingBlkLines'
// to input parameter 'targetBlkLines'.
//
// Be advised that the pre-existing data fields in input parameter
// 'targetBlkLines' will be overwritten and deleted.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetBlkLines             *TextLineSpecBlankLines
//     - A pointer to an instance of TextLineSpecBlankLines. Data
//       extracted from input parameter 'incomingBlkLines' will
//       be copied to this input parameter, 'targetBlkLines'. If
//       this method completes successfully, all member data
//       variables encapsulated in 'targetBlkLines' will be
//       identical to those contained in input parameter,
//       'incomingTimerLines'.
//
//       Be advised that the pre-existing data fields in input
//       parameter 'targetBlkLines' will be overwritten and
//       deleted.
//
//
//  incomingBlkLines           *TextLineSpecBlankLines
//     - A pointer to an instance of TextLineSpecBlankLines.
//
//       All data values in this TextLineSpecBlankLines instance
//       will be copied to input parameter 'targetBlkLines'.
//
//       The original member variables data values encapsulated in
//       'incomingBlkLines' will remain unchanged and will NOT be
//       overwritten or deleted.
//
//       If 'incomingBlkLines' contains invalid member data
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
func (txtBlankLinesMolecule *textLineSpecBlankLinesMolecule) copyIn(
	targetBlkLines *TextLineSpecBlankLines,
	incomingBlkLines *TextLineSpecBlankLines,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtBlankLinesMolecule.lock == nil {
		txtBlankLinesMolecule.lock = new(sync.Mutex)
	}

	txtBlankLinesMolecule.lock.Lock()

	defer txtBlankLinesMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecBlankLinesMolecule.copyIn()",
		"")

	if err != nil {
		return err
	}

	if targetBlkLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'targetBlkLines' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingBlkLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingBlkLines' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	targetBlkLines.textLineReader = nil

	_,
		err = textLineSpecBlankLinesAtom{}.ptr().
		testValidityOfTextLineSpecBlankLines(
			incomingBlkLines,
			ePrefix.XCpy("incomingBlkLines"))

	if err != nil {
		return err
	}

	targetBlkLines.numBlankLines =
		incomingBlkLines.numBlankLines

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&targetBlkLines.newLineChars,
		&incomingBlkLines.newLineChars,
		true,
		ePrefix.XCpy(
			"incomingBlkLines.newLineChars->"+
				"targetBlkLines.newLineChars"))

	return err
}

// copyOut - Returns a deep copy of the TextLineSpecBlankLines
// input parameter 'txtBlankLines'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtBlankLines              *TextLineSpecBlankLines
//     - A pointer to an instance of TextLineSpecBlankLines. A deep
//       copy of the internal member variables will be created
//       and returned in a new instance of TextLineSpecBlankLines.
//
//       If the member variable data values encapsulated by this
//       'txtBlankLines' are found to be invalid, this method will
//       return an error
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
//  TextLineSpecBlankLines
//     - If this method completes successfully, a deep copy of
//       input parameter 'txtBlankLines' will be created and
//       returned in a new instance of TextLineSpecBlankLines.
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
func (txtBlankLinesMolecule *textLineSpecBlankLinesMolecule) copyOut(
	txtBlankLines *TextLineSpecBlankLines,
	errPrefDto *ePref.ErrPrefixDto) (
	TextLineSpecBlankLines,
	error) {

	if txtBlankLinesMolecule.lock == nil {
		txtBlankLinesMolecule.lock = new(sync.Mutex)
	}

	txtBlankLinesMolecule.lock.Lock()

	defer txtBlankLinesMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecBlankLinesMolecule.copyOut()",
		"")

	if err != nil {
		return TextLineSpecBlankLines{}, err
	}

	if txtBlankLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtBlankLines' is a nil pointer!\n",
			ePrefix.String())

		return TextLineSpecBlankLines{}, err
	}

	_,
		err = textLineSpecBlankLinesAtom{}.ptr().
		testValidityOfTextLineSpecBlankLines(
			txtBlankLines,
			ePrefix.XCpy("txtBlankLines"))

	if err != nil {
		return TextLineSpecBlankLines{}, err
	}

	newBlankLinesSpec := TextLineSpecBlankLines{}

	newBlankLinesSpec.numBlankLines = txtBlankLines.numBlankLines

	newBlankLinesSpec.textLineReader = nil

	sMechPreon := strMechPreon{}

	err = sMechPreon.copyRuneArrays(
		&newBlankLinesSpec.newLineChars,
		&txtBlankLines.newLineChars,
		true,
		ePrefix.XCpy(
			"txtBlankLines.newLineChars->"+
				"newBlankLinesSpec.newLineChars"))

	return newBlankLinesSpec, err
}

// empty - Receives a pointer to an instance of
// TextLineSpecBlankLines and proceeds to set all the internal
// member variables to their uninitialized or zero states.
//
// IMPORTANT
// ----------------------------------------------------------------
// The values of all member variables contained in input parameter
// 'textSpecBlkLines' will be overwritten and deleted.
//
func (txtBlankLinesMolecule *textLineSpecBlankLinesMolecule) empty(
	textSpecBlkLines *TextLineSpecBlankLines) {

	if txtBlankLinesMolecule.lock == nil {
		txtBlankLinesMolecule.lock = new(sync.Mutex)
	}

	txtBlankLinesMolecule.lock.Lock()

	defer txtBlankLinesMolecule.lock.Unlock()

	if textSpecBlkLines == nil {
		return
	}

	textSpecBlkLines.numBlankLines = 0

	textSpecBlkLines.textLineReader = nil

	textSpecBlkLines.newLineChars = nil

	return
}

// equal - Receives pointers to two TextLineSpecBlankLines
// instances and proceeds to compare the member data elements to
// determine whether they are equal.
//
// If the data elements of both input parameters 'blkLines' and
// 'incomingBlkLines' are equal in all respects, this method
// returns a boolean value of 'true'. Otherwise, this method
// returns 'false'.
//
func (txtBlankLinesMolecule *textLineSpecBlankLinesMolecule) equal(
	txtBlankLinesOne *TextLineSpecBlankLines,
	txtBlankLinesTwo *TextLineSpecBlankLines) bool {

	if txtBlankLinesMolecule.lock == nil {
		txtBlankLinesMolecule.lock = new(sync.Mutex)
	}

	txtBlankLinesMolecule.lock.Lock()

	defer txtBlankLinesMolecule.lock.Unlock()

	if txtBlankLinesOne == nil {
		return false
	}

	if txtBlankLinesTwo == nil {
		return false
	}

	if txtBlankLinesOne.numBlankLines !=
		txtBlankLinesTwo.numBlankLines {

		return false
	}

	sMechPreon := strMechPreon{}

	if !sMechPreon.equalRuneArrays(
		txtBlankLinesOne.newLineChars,
		txtBlankLinesTwo.newLineChars) {
		return false
	}

	return true
}

// getFormattedText - Receives a pointer to an instance of
// TextLineSpecBlankLines and generates formatted text for
// output display and printing.
//
// The value of 'blkLines.newLineChars' will be replicated
// multiple times as specified by 'blkLines.numBlankLines'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  txtBlankLines              *TextLineSpecBlankLines
//     - A pointer to an instance of TextLineSpecBlankLines. The
//       member variables encapsulated in this object will provide
//       the data necessary to generate blank lines to text output
//       display and printing.
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
//     - If this method completes successfully, formatted text from
//       input parameter 'txtBlankLines' will be generated and
//       returned in this parameter.
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
func (txtBlankLinesMolecule *textLineSpecBlankLinesMolecule) getFormattedText(
	txtBlankLines *TextLineSpecBlankLines,
	errPrefDto *ePref.ErrPrefixDto) (
	formattedText string,
	err error) {

	if txtBlankLinesMolecule.lock == nil {
		txtBlankLinesMolecule.lock = new(sync.Mutex)
	}

	txtBlankLinesMolecule.lock.Lock()

	defer txtBlankLinesMolecule.lock.Unlock()

	formattedText = ""

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecBlankLinesMolecule.getFormattedText()",
		"")

	if err != nil {
		return formattedText, err
	}

	if txtBlankLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtBlankLines' is a nil pointer!\n",
			ePrefix.String())

		return formattedText, err
	}

	_,
		err = textLineSpecBlankLinesAtom{}.ptr().
		testValidityOfTextLineSpecBlankLines(
			txtBlankLines,
			ePrefix.XCpy("txtBlankLines"))

	if err != nil {
		return formattedText, err
	}

	outStr := string(txtBlankLines.newLineChars)

	for i := 0; i < txtBlankLines.numBlankLines; i++ {
		formattedText += outStr
	}

	return formattedText, err
}

// ptr - Returns a pointer to a new instance of
// textLineSpecBlankLinesMolecule.
//
func (txtBlankLinesMolecule textLineSpecBlankLinesMolecule) ptr() *textLineSpecBlankLinesMolecule {

	if txtBlankLinesMolecule.lock == nil {
		txtBlankLinesMolecule.lock = new(sync.Mutex)
	}

	txtBlankLinesMolecule.lock.Lock()

	defer txtBlankLinesMolecule.lock.Unlock()

	return &textLineSpecBlankLinesMolecule{
		lock: new(sync.Mutex),
	}
}

// setTextLinesSpecBlankLines - Receives a pointer to an instance
// of TextLineSpecBlankLines and proceeds to configure that
// instance according to the input parameters, 'numOfBlankLines'
// and 'newLineChars'.
//
// Type TextLineSpecBlankLines is used to configure one or more
// blank lines of text for display, printing for file output.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  txtBlankLines              *TextLineSpecBlankLines
//     - A pointer to an instance of TextLineSpecBlankLines.
//       Member data variables for this instance will be configured
//       in accordance with the following input parameters.
//
//       All pre-existing member data variables in this instance
//       will be deleted and overwritten with the new data values
//       specified in the following input parameters.
//
//
//  numOfBlankLines            int
//     - The number of blank lines which will be generated by an
//       instance of TextLineSpecBlankLines. This value will be
//       tested for validity and configured for input parameter
//       'txtBlankLines'.
//
//
//  newLineRunes               []rune
//     - This rune array holds the text character or characters
//       which will be used as the line termination sequence for
//       input parameter 'txtBlankLines', an instance of
//       TextLineSpecBlankLines.
//
//       If this parameter is submitted as a 'nil' value or a zero
//       length rune array, an error will be returned.
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
func (txtBlankLinesMolecule *textLineSpecBlankLinesMolecule) setTextLinesSpecBlankLines(
	txtBlankLines *TextLineSpecBlankLines,
	numOfBlankLines int,
	newLineRunes []rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtBlankLinesMolecule.lock == nil {
		txtBlankLinesMolecule.lock = new(sync.Mutex)
	}

	txtBlankLinesMolecule.lock.Lock()

	defer txtBlankLinesMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecBlankLinesMolecule."+
			"setTextLinesSpecBlankLines()",
		"")

	if err != nil {
		return err
	}

	if txtBlankLines == nil {
		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'txtBlankLines' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtBlankLinesElectron :=
		textLineSpecBlankLinesElectron{}

	err = txtBlankLinesElectron.
		testValidityNumOfBlankLines(
			numOfBlankLines,
			ePrefix.XCpy(
				"numOfBlankLines"))

	if err != nil {
		return err
	}

	err = txtBlankLinesElectron.
		testValidityNewLinesChars(
			newLineRunes,
			ePrefix.XCpy(
				"newLineRunes"))

	if err != nil {
		return err
	}

	txtBlankLines.numBlankLines = numOfBlankLines

	return strMechPreon{}.ptr().copyRuneArrays(
		&txtBlankLines.newLineChars,
		&newLineRunes,
		true,
		ePrefix.XCpy(
			"newLineRunes->txtBlankLines.newLineChars"))
}

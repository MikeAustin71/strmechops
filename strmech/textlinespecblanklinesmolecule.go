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
// Be advised that the data fields in 'targetBlkLines' will be
// overwritten.
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

	_,
		err = textLineSpecBlankLinesAtom{}.ptr().
		testValidityOfTextLineSpecBlankLines(
			incomingBlkLines,
			ePrefix.XCtx("incomingBlkLines"))

	if err != nil {
		return err
	}

	if len(incomingBlkLines.newLineChars) == 0 {
		incomingBlkLines.newLineChars = []rune{'\n'}
	}

	if incomingBlkLines.numBlankLines < 0 {
		incomingBlkLines.numBlankLines = 0
	}

	targetBlkLines.numBlankLines =
		incomingBlkLines.numBlankLines

	lenIncomingChars := len(incomingBlkLines.newLineChars)

	targetBlkLines.newLineChars =
		make([]rune, lenIncomingChars)

	for i := 0; i < lenIncomingChars; i++ {
		targetBlkLines.newLineChars[i] =
			incomingBlkLines.newLineChars[i]
	}

	return nil
}

// copyOut - Returns a deep copy of the TextLineSpecBlankLines
// input parameter 'blkLines'.
//
func (txtBlankLinesMolecule *textLineSpecBlankLinesMolecule) copyOut(
	txtBlankLines *TextLineSpecBlankLines,
	errPrefDto *ePref.ErrPrefixDto) (
	TextLineSpecBlankLines, error) {

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
			ePrefix.XCtx("txtBlankLines"))

	if err != nil {
		return TextLineSpecBlankLines{}, err
	}

	if len(txtBlankLines.newLineChars) == 0 {
		txtBlankLines.newLineChars = []rune{'\n'}
	}

	if txtBlankLines.numBlankLines < 0 {
		txtBlankLines.numBlankLines = 0
	}

	newBlankLinesSpec := TextLineSpecBlankLines{}

	newBlankLinesSpec.numBlankLines = txtBlankLines.numBlankLines

	lenBlkLineChars := len(txtBlankLines.newLineChars)

	newBlankLinesSpec.newLineChars = make([]rune, lenBlkLineChars)

	copy(newBlankLinesSpec.newLineChars,
		txtBlankLines.newLineChars)

	return newBlankLinesSpec, nil
}

// empty - Receives a pointer to an instance of
// TextLineSpecBlankLines and proceeds to set all of the internal
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

	textSpecBlkLines.newLineChars = nil

	return
}

// equal - Receives pointers to two TextLineSpecBlankLines
// instances and proceeds to compare the member data elements to
// determine whether they are equal.
//
// If the data elements of both input parameters 'blkLines' and
// 'incomingBlkLines' are equal in all respects, this method
// returns a boolean value of 'true'. Otherwise this method returns
// 'false'.
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

	lenCurrBlkLineChars := len(txtBlankLinesOne.newLineChars)

	if lenCurrBlkLineChars !=
		len(txtBlankLinesTwo.newLineChars) {
		return false
	}

	if lenCurrBlkLineChars > 0 {

		for i := 0; i < lenCurrBlkLineChars; i++ {

			if txtBlankLinesOne.newLineChars[i] !=
				txtBlankLinesTwo.newLineChars[i] {
				return false
			}
		}
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
			ePrefix.XCtx("txtBlankLines"))

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

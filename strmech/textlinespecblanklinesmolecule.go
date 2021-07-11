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
	blkLines *TextLineSpecBlankLines,
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

	if blkLines == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'blkLines' is a nil pointer!\n",
			ePrefix.String())

		return TextLineSpecBlankLines{}, err
	}

	if len(blkLines.newLineChars) == 0 {
		blkLines.newLineChars = []rune{'\n'}
	}

	if blkLines.numBlankLines < 0 {
		blkLines.numBlankLines = 0
	}

	newBlankLinesSpec := TextLineSpecBlankLines{}

	newBlankLinesSpec.numBlankLines = blkLines.numBlankLines

	lenBlkLineChars := len(blkLines.newLineChars)

	newBlankLinesSpec.newLineChars = make([]rune, lenBlkLineChars)

	copy(newBlankLinesSpec.newLineChars,
		blkLines.newLineChars)

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
	blkLinesOne *TextLineSpecBlankLines,
	blkLinesTwo *TextLineSpecBlankLines) bool {

	if txtBlankLinesMolecule.lock == nil {
		txtBlankLinesMolecule.lock = new(sync.Mutex)
	}

	txtBlankLinesMolecule.lock.Lock()

	defer txtBlankLinesMolecule.lock.Unlock()

	if blkLinesOne == nil {
		return false
	}

	if blkLinesTwo == nil {
		return false
	}

	if blkLinesOne.numBlankLines !=
		blkLinesTwo.numBlankLines {

		return false
	}

	lenCurrBlkLineChars := len(blkLinesOne.newLineChars)

	if lenCurrBlkLineChars !=
		len(blkLinesTwo.newLineChars) {
		return false
	}

	if lenCurrBlkLineChars > 0 {

		for i := 0; i < lenCurrBlkLineChars; i++ {

			if blkLinesOne.newLineChars[i] !=
				blkLinesTwo.newLineChars[i] {
				return false
			}
		}
	}

	return true
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

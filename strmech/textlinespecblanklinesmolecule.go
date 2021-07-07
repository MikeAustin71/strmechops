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
func (txtBlankLinesMolecule textLineSpecBlankLinesMolecule) copyIn(
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
		"textLineSpecBlankLinesMolecule.copyOut()",
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
func (txtBlankLinesMolecule textLineSpecBlankLinesMolecule) copyOut(
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

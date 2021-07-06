package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecBlankLinesMolecule struct {
	lock *sync.Mutex
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

	newBlankLinesSpec := TextLineSpecBlankLines{}

	newBlankLinesSpec.numBlankLines = blkLines.numBlankLines

	lenBlkLineChars := len(blkLines.newLineChars)

	newBlankLinesSpec.newLineChars = make([]rune, lenBlkLineChars)

	copyLen := copy(newBlankLinesSpec.newLineChars,
		blkLines.newLineChars)

	if copyLen != lenBlkLineChars {
		err = fmt.Errorf("%v\n"+
			"Error: Copy blkLines.newLineChars ->!\n"+
			"  newBlankLinesSpec.newLineChars FAILED!\n"+
			"Expected %v-elements to be copied. However,\n"+
			"only %v-elements were copied!\n",
			ePrefix.String(),
			lenBlkLineChars,
			copyLen)

		return TextLineSpecBlankLines{}, err
	}

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

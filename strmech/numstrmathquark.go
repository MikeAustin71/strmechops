package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathQuark struct {
	lock *sync.Mutex
}

func (nStrMathQuark *numStrMathQuark) extendRunesRight(
	charsToExtend *RuneArrayDto,
	fillChar rune,
	newCharArrayLength int,
	errPrefDto *ePref.ErrPrefixDto) error {

	nStrMathQuark.lock.Lock()

	defer nStrMathQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathQuark."+
			"extendRunesRight()",
		"")

	if err != nil {
		return err
	}

	if charsToExtend == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charArrayLen := charsToExtend.GetRuneArrayLength()

	if charArrayLen >= newCharArrayLength {
		// Nothing to do.
		return err
	}

	lenghtDelta := newCharArrayLength - charArrayLen

	newFillArray := make([]rune, lenghtDelta)

	for i := 0; i < lenghtDelta; i++ {
		newFillArray[i] = fillChar
	}

	charsToExtend.CharsArray =
		append(
			charsToExtend.CharsArray,
			newFillArray...)

	return err
}

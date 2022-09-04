package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathRoundingQuark struct {
	lock *sync.Mutex
}

func (nStrMathRoundQuark *numStrMathRoundingQuark) preRoundingValidation(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	roundToFractionalDigits int,
	numberSign NumericSignValueType,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathRoundQuark.lock == nil {
		nStrMathRoundQuark.lock = new(sync.Mutex)
	}

	nStrMathRoundQuark.lock.Lock()

	defer nStrMathRoundQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"nStrMathRoundQuark."+
			"preRoundingValidation()",
		"")

	if err != nil {
		return err
	}

	if integerDigits == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigits' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if fractionalDigits == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fractionalDigits' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if roundToFractionalDigits < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'roundToFractionalDigits' is "+
			"less than zero (0)!\n"+
			"roundToFractionalDigits = %v\n",
			ePrefix.String(),
			roundToFractionalDigits)

		return err

	}

	if !numberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberSign' is invalid!\n"+
			"numberSign String Value  = '%v'\n"+
			"numberSign Integer Value = '%v'\n",
			ePrefix.String(),
			numberSign.String(),
			numberSign.XValueInt())

		return err

	}

	if integerDigits.GetRuneArrayLength() == 0 {

		integerDigits.CharsArray = make([]rune, 1)

		integerDigits.CharsArray[0] = '0'
	}

	return err
}

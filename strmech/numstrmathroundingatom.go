package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathRoundingAtom struct {
	lock *sync.Mutex
}

func (nStrMathRoundAtom *numStrMathRoundingAtom) roundHalfAwayFromZero(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathRoundAtom.lock == nil {
		nStrMathRoundAtom.lock = new(sync.Mutex)
	}

	nStrMathRoundAtom.lock.Lock()

	defer nStrMathRoundAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"nStrMathRoundAtom."+
			"roundHalfAwayFromZero()",
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

	lenFracDigits := fractionalDigits.GetRuneArrayLength()

	if roundToFractionalDigits > lenFracDigits {
		// Not much to do. Extend with zeroes to
		// to the right and exit!
		return new(numStrMathQuark).extendRunes(
			fractionalDigits,
			fractionalDigits,
			'0',
			roundToFractionalDigits,
			true,
			ePrefix.XCpy(
				fmt.Sprintf("roundToFractionalDigits= %v",
					roundToFractionalDigits)))

	}

	if lenFracDigits == roundToFractionalDigits {
		// Nothing to do. Already rounded!
		// Exit Here.

		return err
	}

	// roundToFractionalDigits MUST BE
	//  less than existingNumOfFracDigits
	//roundIdx := roundToFractionalDigits

	roundUp := true

	if fractionalDigits.CharsArray[roundToFractionalDigits] <
		'5' {

		roundUp = false
	}

	// Truncate Fractional Digits array!
	// Last Index = roundToFractionalDigits - 1
	fractionalDigits.CharsArray =
		fractionalDigits.CharsArray[:roundToFractionalDigits]

	if roundUp == false {
		return err
	}

	// MUST BE ROUND UP!

	return err

}

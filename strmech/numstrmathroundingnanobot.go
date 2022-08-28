package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathRoundingNanobot struct {
	lock *sync.Mutex
}

func (nStrMathRoundNanobot *numStrMathRoundingNanobot) roundNumStrKernel(
	numStrKernel *NumberStrKernel,
	numStrRoundingSpec NumStrRoundingSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathRoundNanobot.lock == nil {
		nStrMathRoundNanobot.lock = new(sync.Mutex)
	}

	nStrMathRoundNanobot.lock.Lock()

	defer nStrMathRoundNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathRoundingNanobot."+
			"roundNumStrKernel()",
		"")

	if err != nil {
		return err
	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var roundingType NumberRoundingType

	roundingType = numStrRoundingSpec.GetRoundingType()

	if !roundingType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrRoundingSpec Rounding Type' is invalid!\n"+
			"'roundingType' string  value = '%v'\n"+
			"'roundingType' integer value = '%v'\n",
			ePrefix.String(),
			roundingType.String(),
			roundingType.XValueInt())

		return err

	}

	if roundingType == NumRoundType.NoRounding() {
		// Nothing to do

		return err
	}

	var roundToFractionalDigits int

	roundToFractionalDigits =
		numStrRoundingSpec.GetRoundToFractionalDigits()

	if roundToFractionalDigits < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrRoundingSpec RoundToFractionalDigits' is invalid!\n"+
			"'roundToFractionalDigits' has a value which is less than zero (0).\n"+
			"roundToFractionalDigits = '%v'\n",
			ePrefix.String(),
			roundToFractionalDigits)

		return err
	}

	numOfFracDigits := numStrKernel.GetNumberOfFractionalDigits()

	if roundToFractionalDigits >
		numOfFracDigits {

		return new(numStrMathQuark).extendRunes(
			&numStrKernel.fractionalDigits,
			&numStrKernel.fractionalDigits,
			'0',
			roundToFractionalDigits,
			true,
			ePrefix.XCpy(
				fmt.Sprintf("roundToFractionalDigits= %v",
					roundToFractionalDigits)))

	}

	if roundToFractionalDigits ==
		numOfFracDigits {
		// Nothing to do. Already rounded

		return err
	}

	switch roundingType {

	case NumRoundType.HalfAwayFromZero():

		err = new(numStrMathMolecule).
			roundHalfAwayFromZero(
				numStrKernel,
				roundToFractionalDigits,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	default:

		err = fmt.Errorf("%v\n"+
			"Error: This rounding algorithm is not supported!\n"+
			"Rounding Type string value  = '%v'\n"+
			"Rounding Type integer value = '%v'\n",
			ePrefix.String(),
			roundingType.String(),
			roundingType.XValueInt())

	}

	return err
}

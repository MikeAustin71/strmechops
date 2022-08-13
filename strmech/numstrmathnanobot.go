package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathNanobot struct {
	lock *sync.Mutex
}

func (nStrMathNanobot *numStrMathNanobot) roundNumStrKernel(
	numStrKernel *NumberStrKernel,
	numStrRoundingSpec NumStrRoundingSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrMathNanobot.lock == nil {
		nStrMathNanobot.lock = new(sync.Mutex)
	}

	nStrMathNanobot.lock.Lock()

	defer nStrMathNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathNanobot."+
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

	if !numStrRoundingSpec.RoundFractionalDigits {
		// Nothing to do

		return err
	}

	if numStrRoundingSpec.RoundToFractionalDigits < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrRoundingSpec.RoundToFractionalDigits' is invalid!\n"+
			"'numStrRoundingSpec.RoundToFractionalDigits' has a value which is less than zero (0).\n"+
			"numStrRoundingSpec.RoundToFractionalDigits = '%v'\n",
			ePrefix.String(),
			numStrRoundingSpec.RoundToFractionalDigits)

		return err
	}

	numOfFracDigits := numStrKernel.GetNumberOfFractionalDigits()

	if numStrRoundingSpec.RoundToFractionalDigits >
		numOfFracDigits {

		return new(numStrMathAtom).extendFractionalDigits(
			numStrKernel,
			numStrRoundingSpec.RoundToFractionalDigits,
			ePrefix.XCpy(
				fmt.Sprintf("roundToFractionalDigits= %v",
					numStrRoundingSpec.RoundToFractionalDigits)))

	}

	if numStrRoundingSpec.RoundToFractionalDigits ==
		numOfFracDigits {
		// Nothing to do. Already rounded

		return err
	}

	switch numStrRoundingSpec.RoundingType {

	case NumRoundType.None():

		err = fmt.Errorf("%v\n"+
			"Error: Rounding was specified, but the\n"+
			"Number Rounding Type is NumRoundType.None().\n",
			ePrefix.String())

	case NumRoundType.HalfAwayFromZero():

		err = new(numStrMathMolecule).
			roundHalfAwayFromZero(
				numStrKernel,
				numStrRoundingSpec.
					RoundToFractionalDigits,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							RoundToFractionalDigits)))

		if err != nil {
			return err
		}

	default:

		err = fmt.Errorf("%v\n"+
			"Error: This rounding algorithm is not supported!\n"+
			"Rounding Type = '%v'\n",
			ePrefix.String(),
			numStrRoundingSpec.RoundingType.String())

	}

	return err
}

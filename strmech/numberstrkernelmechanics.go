package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelMechanics
//
// Provides helper methods for type NumberStrKernel.
type numberStrKernelMechanics struct {
	lock *sync.Mutex
}

//	convertToSciNotation
//
//	Receives a pointer to an instance of numStrKernel and
//	proceeds to convert the intrinsic numeric value to
//	Scientific Notation before returning an instance of
//	SciNotationKernel.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		converted to Scientific Notation and returned
//		as an instance of SciNotationKernel.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	sciNotKernel				SciNotationKernel
//
//		This returned instance of SciNotationKernel will
//		be configured with the numeric value contained in
//		input parameter 'numStrKernel'.
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernelMech *numberStrKernelMechanics) convertToSciNotation(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	sciNotKernel SciNotationKernel,
	err error) {

	if numStrKernelMech.lock == nil {
		numStrKernelMech.lock = new(sync.Mutex)
	}

	numStrKernelMech.lock.Lock()

	defer numStrKernelMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelMechanics."+
			"convertToSciNotation()",
		"")

	if err != nil {

		return sciNotKernel, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return sciNotKernel, err
	}

	intArrayLen := numStrKernel.integerDigits.GetRuneArrayLength()

	fracArrayLen := numStrKernel.fractionalDigits.GetRuneArrayLength()

	nStrKernelNanobot := numberStrKernelNanobot{}

	if (intArrayLen == 0 &&
		fracArrayLen == 0) ||
		(numStrKernel.integerDigits.IsAllNumericZeros() &&
			numStrKernel.fractionalDigits.IsAllNumericZeros()) {

		err = nStrKernelNanobot.setWithRunes(
			&sciNotKernel.significand,
			[]rune{'0'},
			[]rune{'0'},
			NumSignVal.Zero(),
			ePrefix)

		if err != nil {

			return sciNotKernel, err
		}

		err = nStrKernelNanobot.setWithRunes(
			&sciNotKernel.exponent,
			[]rune{'0'},
			[]rune{},
			NumSignVal.Zero(),
			ePrefix)

		return sciNotKernel, err
	}

	var intArray RuneArrayDto

	intArray,
		err = numStrKernel.integerDigits.CopyOut(
		ePrefix.XCpy(
			"numStrKernel.integerDigits"))

	if err != nil {

		return sciNotKernel, err
	}

	var zerosCount uint64

	zerosCount = intArray.GetCountTrailingZeros()

	var deleteTrailingChars bool

	deleteTrailingChars = false

	// Delete all leading intArray Zeros
	err = intArray.DeleteLeadingTrailingChars(
		zerosCount,
		deleteTrailingChars,
		ePrefix.XCpy(
			fmt.Sprintf(
				"deleteTrailingChars='%v'"+
					" intArray zerosCount='%v'",
				deleteTrailingChars,
				zerosCount)))

	if err != nil {

		return sciNotKernel, err
	}

	intArrayLen = intArray.GetRuneArrayLength()

	var fracArray RuneArrayDto

	fracArray,
		err = numStrKernel.fractionalDigits.CopyOut(
		ePrefix.XCpy(
			"fracArray<-"))

	if err != nil {

		return sciNotKernel, err
	}

	zerosCount = fracArray.GetCountTrailingZeros()

	deleteTrailingChars = true

	// Delete Trailing Fractional Zeros
	err = fracArray.DeleteLeadingTrailingChars(
		zerosCount,
		deleteTrailingChars,
		ePrefix.XCpy(
			fmt.Sprintf(
				"deleteTrailingChars='%v'"+
					" fracArray zerosCount='%v'",
				deleteTrailingChars,
				zerosCount)))

	if err != nil {

		return sciNotKernel, err
	}

	fracArrayLen = fracArray.GetRuneArrayLength()

	// Compute Significand and Exponent

	var newIntRuneArray, newFracRuneArray []rune

	var exponent int64

	nStrKernelMolecule := numberStrKernelMolecule{}

	if intArrayLen > 0 && fracArrayLen == 0 {

		err = nStrKernelNanobot.setWithRunes(
			&sciNotKernel.significand,
			intArray.CharsArray,
			[]rune{},
			numStrKernel.numberSign,
			ePrefix)

		if err != nil {

			return sciNotKernel, err
		}

		exponent = 0

		err = nStrKernelMolecule.convertNumberToKernel(
			&sciNotKernel.exponent,
			exponent,
			NumSignVal.Zero(),
			ePrefix.XCpy(
				fmt.Sprintf("sciNotKernel.exponent='%v'",
					exponent)))

		if err != nil {

			return sciNotKernel, err
		}

	} else if intArrayLen >= 1 {

		newIntRuneArray = make([]rune, 1)
		newIntRuneArray[0] = intArray.CharsArray[0]

		newFracRuneArray = append(
			intArray.CharsArray[1:],
			fracArray.CharsArray...)

		err = nStrKernelNanobot.setWithRunes(
			&sciNotKernel.significand,
			newIntRuneArray,
			newFracRuneArray,
			numStrKernel.numberSign,
			ePrefix)

		if err != nil {

			return sciNotKernel, err
		}

		exponent = int64(intArrayLen - 1)

		err = nStrKernelMolecule.convertNumberToKernel(
			&sciNotKernel.exponent,
			exponent,
			NumSignVal.Positive(),
			ePrefix.XCpy(
				fmt.Sprintf("sciNotKernel.exponent='%v'",
					exponent)))

		if err != nil {

			return sciNotKernel, err
		}

	} else if intArrayLen == 1 {

		newIntRuneArray = make([]rune, 1)
		newIntRuneArray[0] = intArray.CharsArray[0]

		err = nStrKernelNanobot.setWithRunes(
			&sciNotKernel.significand,
			newIntRuneArray,
			fracArray.CharsArray,
			numStrKernel.numberSign,
			ePrefix)

		if err != nil {

			return sciNotKernel, err
		}

		exponent = 0

		err = nStrKernelMolecule.convertNumberToKernel(
			&sciNotKernel.exponent,
			exponent,
			NumSignVal.Zero(),
			ePrefix.XCpy(
				fmt.Sprintf("sciNotKernel.exponent='%v'",
					exponent)))

		if err != nil {

			return sciNotKernel, err
		}

	} else {
		// MUST BE intArrayLen <= 0 &&
		//	fracArrayLen > 0

		leadingFracZerosCount := fracArray.GetCountLeadingZeros()

		if leadingFracZerosCount > 0 {
			// Delete Leading Fractional Zeros
			deleteTrailingChars = false
			err = fracArray.DeleteLeadingTrailingChars(
				zerosCount,
				deleteTrailingChars,
				ePrefix.XCpy(
					fmt.Sprintf(
						"deleteTrailingChars='%v'"+
							" fracArray zerosCount='%v'",
						deleteTrailingChars,
						zerosCount)))

			if err != nil {

				return sciNotKernel, err
			}

			newIntRuneArray = make([]rune, 1)
			newIntRuneArray[0] = fracArray.CharsArray[0]

			newFracRuneArray = append(
				newFracRuneArray,
				intArray.CharsArray[1:]...)

			err = nStrKernelNanobot.setWithRunes(
				&sciNotKernel.significand,
				newIntRuneArray,
				newFracRuneArray,
				numStrKernel.numberSign,
				ePrefix)

			if err != nil {

				return sciNotKernel, err
			}

			exponent = int64(zerosCount + 1)

			err = nStrKernelMolecule.convertNumberToKernel(
				&sciNotKernel.exponent,
				exponent,
				NumSignVal.Zero(),
				ePrefix.XCpy(
					fmt.Sprintf("sciNotKernel.exponent='%v'",
						exponent)))

			if err != nil {

				return sciNotKernel, err
			}

		} else {
			// MUST BE leadingFracZerosCount <= 0

			newIntRuneArray = make([]rune, 1)
			newIntRuneArray[0] = fracArray.CharsArray[0]

			newFracRuneArray = append(
				newFracRuneArray,
				intArray.CharsArray[1:]...)

			err = nStrKernelNanobot.setWithRunes(
				&sciNotKernel.significand,
				newIntRuneArray,
				newFracRuneArray,
				numStrKernel.numberSign,
				ePrefix)

			if err != nil {

				return sciNotKernel, err
			}

			exponent = int64(zerosCount + 1)

			err = nStrKernelMolecule.convertNumberToKernel(
				&sciNotKernel.exponent,
				exponent,
				NumSignVal.Zero(),
				ePrefix.XCpy(
					fmt.Sprintf("sciNotKernel.exponent='%v'",
						exponent)))

			if err != nil {

				return sciNotKernel, err
			}

		}

	} // END OF Compute Significand and Exponent

	return sciNotKernel, err
}

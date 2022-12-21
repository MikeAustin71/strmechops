package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathRoundingNanobot struct {
	lock *sync.Mutex
}

// roundNumStrKernel
//
// Receives an instance of NumberStrKernel and a rounding
// specification (NumStrRoundingSpec). The method then
// proceeds to apply the selected rounding algorithm to
// the numeric value contained in the NumberStrKernel.
//
// The Number String Rounding Specification allows users to
// apply numeric value formatting algorithms such as
// 'Truncate', 'Floor' and 'Ceiling'.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
//	This method will modify the data values contained in
//	the input parameter, 'numStrKernel'. The rounded
//	numeric values produced by this method will be stored
//	in the 'numStrKernel' instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		This instance of NumberStrKernel contains the
//		numeric value to be rounded.
//
//		This instance will be modified to reflect the
//		rounded numeric value produced by this rounding
//		operation.
//
//	numStrRoundingSpec			*NumStrRoundingSpec
//
//		This data transfer object contains all the
//		parameters required to configure a rounding
//		algorithm for a floating point number string.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which
//		is included in all returned error messages. Usually,
//		it contains the name of the calling method or methods
//		listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during	processing, the returned error
//		Type will encapsulate an error	message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
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

	err = numStrRoundingSpec.IsValidInstanceError(
		ePrefix.XCpy(
			"numStrRoundingSpec"))

	if err != nil {
		return err
	}

	var roundingType NumberRoundingType

	roundingType = numStrRoundingSpec.roundingType

	if roundingType == NumRoundType.NoRounding() {
		// Nothing to do

		return err
	}

	var roundToFractionalDigits int

	roundToFractionalDigits =
		numStrRoundingSpec.roundToFractionalDigits

	numOfFracDigits := len(numStrKernel.fractionalDigits.CharsArray)

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

	case NumRoundType.Ceiling():

		err = new(numStrMathRoundingAtom).
			ceiling(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				numStrKernel.numberSign,
				ePrefix)

	case NumRoundType.Floor():

		err = new(numStrMathRoundingAtom).
			floor(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				numStrKernel.numberSign,
				ePrefix)

	case NumRoundType.HalfUpWithNegNums():

		err = new(numStrMathRoundingAtom).
			roundHalfUpWithNegNums(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				roundToFractionalDigits,
				numStrKernel.numberSign,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	case NumRoundType.HalfDownWithNegNums():

		err = new(numStrMathRoundingAtom).
			roundHalfDownWithNegNums(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				roundToFractionalDigits,
				numStrKernel.numberSign,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	case NumRoundType.HalfAwayFromZero():

		err = new(numStrMathRoundingAtom).
			roundHalfAwayFromZero(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				roundToFractionalDigits,
				numStrKernel.numberSign,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	case NumRoundType.HalfTowardsZero():

		err = new(numStrMathRoundingAtom).
			roundHalfTowardsZero(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				roundToFractionalDigits,
				numStrKernel.numberSign,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	case NumRoundType.HalfToEven():

		err = new(numStrMathRoundingAtom).
			roundHalfToEven(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				roundToFractionalDigits,
				numStrKernel.numberSign,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	case NumRoundType.HalfToOdd():

		err = new(numStrMathRoundingAtom).
			roundHalfToOdd(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				roundToFractionalDigits,
				numStrKernel.numberSign,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	case NumRoundType.Randomly():

		err = new(numStrMathRoundingAtom).
			roundRandomly(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				roundToFractionalDigits,
				numStrKernel.numberSign,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	case NumRoundType.Truncate():

		err = new(numStrMathRoundingAtom).
			truncate(
				&numStrKernel.integerDigits,
				&numStrKernel.fractionalDigits,
				roundToFractionalDigits,
				numStrKernel.numberSign,
				ePrefix.XCpy(
					fmt.Sprintf("newNumStrKernel<-"+
						"RoundTo %v-digits",
						numStrRoundingSpec.
							roundToFractionalDigits)))

	default:

		err = fmt.Errorf("%v\n"+
			"Error: This rounding algorithm selected is invalid!\n"+
			"Rounding Type string value  = '%v'\n"+
			"Rounding Type integer value = '%v'\n",
			ePrefix.String(),
			roundingType.String(),
			roundingType.XValueInt())

	}

	return err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathRoundingAtom struct {
	lock *sync.Mutex
}

// roundHalfUpWithNegNums
//
// Round Half Round Up Including Negative Numbers.
// This method is intuitive but may produce unexpected
// results when applied to negative numbers.
//
// 'HalfUpWithNegNums' rounds 0.5 up.
//
//	Examples of 'HalfUpWithNegNums'
//
//	7.6 rounds up to 8
//	7.5 rounds up to 8
//	7.4 rounds down to 7
//	-7.4 rounds up to -7
//	-7.5 rounds up to -7
//	-7.6 rounds down to -8
//
// Reference:
//
//	https://www.mathsisfun.com/numbers/rounding-methods.html
//	https://en.wikipedia.org/wiki/Rounding
//	https://www.mathsisfun.com/rounding-numbers.html
//	https://www.vedantu.com/maths/rounding-methods
//	https://rounding.to/the-most-common-rounding-methods/
//	https://www.wikihow.com/Round-Numbers
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of integer numeric
//		digits which may be used to the 'Round Half Away From
//		Zero' algorithm.
//
//		These integer numeric digits comprise the integer
//		portion of a floating point numeric value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional numeric
//		digits which may be used to the 'Round Half Away From
//		Zero' algorithm.
//
//		These fractional numeric digits comprise the fractional
//		portion of a floating point numeric value.
//
//	roundToFractionalDigits		int
//
//		The number of digits to the right of the decimal
//		separator (a.k.a. decimal point) which will remain
//		after this rounding algorithm is applied.
//
//		Example:
//			Floating Point Number  0.12345
//			roundToFractionalDigits 3
//			Floating Point Number after rounding: 0.123
//
//	numberSign					NumericSignValueType
//
//		Type NumericSignValueType is an enumeration of possible
//		number sign values listed as follows:
//			NumSignVal.None()     = -2 - Invalid Value
//			NumSignVal.Negative() = -1 - Valid Value
//			NumSignVal.Zero()     =  0 - Valid Value
//			NumSignVal.Positive() =  1 - Valid Value
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
// ------------------------------------------------------------------------
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) roundHalfUpWithNegNums(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	roundToFractionalDigits int,
	numberSign NumericSignValueType,
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
			"roundHalfUpWithNegNums()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrMathRoundingQuark).preRoundingValidation(
		integerDigits,
		fractionalDigits,
		roundToFractionalDigits,
		numberSign,
		ePrefix)

	if err != nil {
		return err
	}

	lenFracDigits := fractionalDigits.GetRuneArrayLength()

	if roundToFractionalDigits == 0 &&
		lenFracDigits == 0 {
		return err
	}

	if roundToFractionalDigits > lenFracDigits {
		// Not much to do. Extend with zeroes to
		// the right and exit!
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

	roundUp := false

	if numberSign == NumSignVal.Zero() {

		roundUp = false

	} else if numberSign == NumSignVal.Positive() {

		if fractionalDigits.CharsArray[roundToFractionalDigits] >=
			'5' {

			roundUp = true
		}

	} else {
		// MUST BE
		// numberSign ==  NumSignVal.Negative()

		if fractionalDigits.CharsArray[roundToFractionalDigits] >
			'5' {

			roundUp = true
		}

	}

	if roundUp == false {

		return new(numStrMathRoundingElectron).roundDown(
			integerDigits,
			fractionalDigits,
			roundToFractionalDigits,
			numberSign,
			ePrefix)

	}

	return new(numStrMathRoundingElectron).roundUp(
		integerDigits,
		fractionalDigits,
		roundToFractionalDigits,
		numberSign,
		ePrefix)
}

// roundHalfAwayFromZero - Performs a rounding operation on
// the integer and fractional numeric digit arrays contained
// in two instances of RuneArrayDto passed as input
// parameters.
//
// This method performs and applies the 'Round Half Away From
// Zero' algorithm.
//
// Examples of 'HalfAwayFromZero' Rounding Algorithm
//
//	7.6 rounds away to 8
//	7.5 rounds away to 8
//	7.4 rounds to 7
//	-7.4 rounds to -7
//	-7.5 rounds away to -8
//	-7.6 rounds away to -8
//
// This algorithm corresponds to NumberRoundingType:
//
//	NumRoundType.HalfAwayFromZero
//
// Reference:
//
//	https://www.mathsisfun.com/numbers/rounding-methods.html
//	https://en.wikipedia.org/wiki/Rounding
//	https://www.mathsisfun.com/rounding-numbers.html
//	https://www.vedantu.com/maths/rounding-methods
//	https://rounding.to/the-most-common-rounding-methods/
//	https://www.wikihow.com/Round-Numbers
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of integer numeric
//		digits which may be used to the 'Round Half Away From
//		Zero' algorithm.
//
//		These integer numeric digits comprise the integer
//		portion of a floating point numeric value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional numeric
//		digits which may be used to the 'Round Half Away From
//		Zero' algorithm.
//
//		These fractional numeric digits comprise the fractional
//		portion of a floating point numeric value.
//
//	roundToFractionalDigits		int
//
//		The number of digits to the right of the decimal
//		separator (a.k.a. decimal point) which will remain
//		after this rounding algorithm is applied.
//
//		Example:
//			Floating Point Number  0.12345
//			roundToFractionalDigits 3
//			Floating Point Number after rounding: 0.123
//
//	numberSign					NumericSignValueType
//
//		Type NumericSignValueType is an enumeration of possible
//		number sign values listed as follows:
//			NumSignVal.None()     = -2 - Invalid Value
//			NumSignVal.Negative() = -1 - Valid Value
//			NumSignVal.Zero()     =  0 - Valid Value
//			NumSignVal.Positive() =  1 - Valid Value
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
// ------------------------------------------------------------------------
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) roundHalfAwayFromZero(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	roundToFractionalDigits int,
	numberSign NumericSignValueType,
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

	err = new(numStrMathRoundingQuark).preRoundingValidation(
		integerDigits,
		fractionalDigits,
		roundToFractionalDigits,
		numberSign,
		ePrefix)

	if err != nil {
		return err
	}

	lenFracDigits := fractionalDigits.GetRuneArrayLength()

	if roundToFractionalDigits == 0 &&
		lenFracDigits == 0 {
		return err
	}

	if roundToFractionalDigits > lenFracDigits {
		// Not much to do. Extend with zeroes to
		// the right and exit!
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

	if fractionalDigits.CharsArray[roundToFractionalDigits] <
		'5' {
		// ROUND DOWN!
		return new(numStrMathRoundingElectron).roundDown(
			integerDigits,
			fractionalDigits,
			roundToFractionalDigits,
			numberSign,
			ePrefix)
	}

	// MUST BE ROUND UP!

	return new(numStrMathRoundingElectron).roundUp(
		integerDigits,
		fractionalDigits,
		roundToFractionalDigits,
		numberSign,
		ePrefix)
}

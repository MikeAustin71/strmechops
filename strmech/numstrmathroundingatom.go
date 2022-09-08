package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathRoundingAtom struct {
	lock *sync.Mutex
}

// ceiling
//
// Yields the nearest integer up. Ceiling does not apply
// any special treatment to 0.5.
//
//	Ceiling Function: The least integer that is greater than or
//	                  equal to x.
//
// Source:
//
//	https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//	The ceiling function maps x to the least integer greater than
//	or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//
// Source:
//
//	https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//		Examples of Ceiling
//
//		  Number    Ceiling
//		   2           2
//		   2.4         3
//		   2.9         3
//		  -2.5        -2
//		  -2.7        -2
//		  -2          -2
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.Ceiling()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of integer numeric
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
//
//		These fractional numeric digits comprise the fractional
//		portion of a floating point numeric value.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) ceiling(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
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
			"ceiling()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrMathRoundingQuark).preRoundingValidation(
		integerDigits,
		fractionalDigits,
		0,
		numberSign,
		ePrefix)

	if err != nil {
		return err
	}

	lenFracDigits := fractionalDigits.GetRuneArrayLength()

	if lenFracDigits == 0 {
		return err
	}

	fractionalDigits.CharsArray = nil

	if numberSign == NumSignVal.Negative() {

		return err

	}

	// MUST BE
	// numberSign == NumSignVal.Positive()
	//            OR
	// numberSign == NumSignVal.Zero()

	_,
		err = new(numStrMathAtom).addOneToRunes(
		integerDigits,
		integerDigits,
		true,
		ePrefix.XCpy(
			"integerDigits<- +1"))

	return err

}

// floor
//
// Yields the nearest integer down. Floor does not apply
// any special treatment to 0.5.
//
//	Floor Function: The greatest integer that is less than or
//	                equal to x
//
// Source:
//
//	https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
// In mathematics and computer science, the floor function is the
// function that takes as input a real number x, and gives as
// output the greatest integer less than or equal to x, denoted
// floor(x) or ⌊x⌋.
//
// Source:
//
//	https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//		Examples of Floor
//
//		  Number     Floor
//		   2           2
//		   2.4         2
//		   2.9         2
//		  -2.5        -3
//		  -2.7        -3
//		  -2          -2
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.Floor()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of integer numeric
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
//
//		These fractional numeric digits comprise the fractional
//		portion of a floating point numeric value.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) floor(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
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
			"floor()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrMathRoundingQuark).preRoundingValidation(
		integerDigits,
		fractionalDigits,
		0,
		numberSign,
		ePrefix)

	if err != nil {
		return err
	}

	lenFracDigits := fractionalDigits.GetRuneArrayLength()

	if lenFracDigits == 0 {
		return err
	}

	fractionalDigits.CharsArray = nil

	if numberSign == NumSignVal.Zero() ||
		numberSign == NumSignVal.Positive() {

		return err

	}

	// MUST BE
	// numberSign == NumSignVal.Negative()

	_,
		err = new(numStrMathAtom).addOneToRunes(
		integerDigits,
		integerDigits,
		true,
		ePrefix.XCpy(
			"integerDigits<- +1"))

	return err
}

// roundHalfUpWithNegNums
//
// Performs a rounding operation on the integer and fractional
// numeric digit arrays contained in two instances of
// RuneArrayDto passed as input parameters.
//
// Round Half Round Up Including Negative Numbers.
// This method is intuitive but may produce unexpected
// results when applied to negative numbers.
//
// 'HalfUpWithNegNums' rounds 0.5 up.
//
//	Examples of 'HalfUpWithNegNums'
//
//		7.6 rounds up to 8
//		7.5 rounds up to 8
//		7.4 rounds down to 7
//		-7.4 rounds up to -7
//		-7.5 rounds up to -7
//		-7.6 rounds down to -8
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.HalfUpWithNegNums()
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
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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

// roundHalfDownWithNegNums
//
// Performs a rounding operation on the integer and fractional
// numeric digit arrays contained in two instances of
// RuneArrayDto passed as input parameters.
//
// Round Half Round Down Including Negative Numbers.
//
// This method is considered intuitive but may produce
// unexpected results when applied to negative numbers.
//
// 'HalfDownWithNegNums' rounds 0.5 down.
//
//	Examples of 'HalfDownWithNegNums'
//
//		7.6 rounds up to 8
//		7.5 rounds down to 7
//		7.4 rounds down to 7
//		-7.4 rounds up to -7
//		-7.5 rounds down to -8
//		-7.6 rounds down to -8
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.HalfDownWithNegNums()
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
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) roundHalfDownWithNegNums(
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
			"roundHalfDownWithNegNums()",
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

		if fractionalDigits.CharsArray[roundToFractionalDigits] >
			'5' {

			roundUp = true
		}

	} else {
		// MUST BE
		// numberSign ==  NumSignVal.Negative()

		if fractionalDigits.CharsArray[roundToFractionalDigits] >
			'4' {

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

// roundHalfAwayFromZero
//
// Performs a rounding operation on the integer and fractional
// numeric digit arrays contained in two instances of
// RuneArrayDto passed as input parameters.
//
// This method performs and applies the 'Round Half Away From
// Zero' algorithm.
//
//	Examples of 'HalfAwayFromZero' Rounding Algorithm
//
//		7.6 rounds away to 8
//		7.5 rounds away to 8
//		7.4 rounds to 7
//		-7.4 rounds to -7
//		-7.5 rounds away to -8
//		-7.6 rounds away to -8
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.HalfAwayFromZero()
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
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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

// roundHalfTowardsZero
//
// Performs a rounding operation on the integer and fractional
// numeric digit arrays contained in two instances of
// RuneArrayDto passed as input parameters.
//
// Round Half Towards Zero.  'HalfTowardsZero' rounds 0.5 closer
// to zero. It provides clear and consistent behavior when
// dealing with negative numbers.
//
// This method performs and applies the 'Round Half Away Towards
// Zero' algorithm.
//
//	Examples of HalfTowardsZero
//
//		7.6 rounds away to 8
//		7.5 rounds to 7
//		7.4 rounds to 7
//		-7.4 rounds to -7
//		-7.5 rounds to -7
//		-7.6 rounds away to -8
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.HalfTowardsZero()
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
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) roundHalfTowardsZero(
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
			"roundHalfTowardsZero()",
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

	if fractionalDigits.CharsArray[roundToFractionalDigits] <=
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

// roundHalfToEven
//
// Performs a rounding operation on the integer and fractional
// numeric digit arrays contained in two instances of
// RuneArrayDto passed as input parameters.
//
// This method performs and applies the 'Round To Even'
// algorithm.
//
// 'Round to Even' is known as (Banker's Rounding). It rounds
// 0.5 to the nearest even digit.
//
//	Examples of HalfToEven
//
//	7.5 rounds up to 8 (because 8 is an even number)
//	but 6.5 rounds down to 6 (because 6 is an even number)
//
//	HalfToEven only applies to 0.5. Other numbers (not ending
//	in 0.5) round to nearest as usual, so:
//
//		7.6 rounds up to 8
//		7.5 rounds up to 8 (because 8 is an even number)
//		7.4 rounds down to 7
//		6.6 rounds up to 7
//		6.5 rounds down to 6 (because 7 is an odd number)
//		6.4 rounds down to 6
//		-23.5 rounds to -24 (because 4 is an even number)
//		-24.5 rounds to -24 (because 5 is an odd number)
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.HalfToEven()
//
// Reference:
//
//	https://en.wikipedia.org/wiki/Rounding#Round_half_to_even
//	https://rounding.to/understanding-the-bankers-rounding/
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
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) roundHalfToEven(
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
			"roundHalfToEven()",
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

	if fractionalDigits.CharsArray[roundToFractionalDigits] <
		'5' {

		// ROUND DOWN!
		roundUp = false

	} else if fractionalDigits.CharsArray[roundToFractionalDigits] >
		'5' {
		// ROUND UP!
		roundUp = true

	} else {
		// MUST BE
		// fractionalDigits.CharsArray[roundToFractionalDigits] ==
		// '5'

		testIdx := roundToFractionalDigits - 1

		var testDigit rune

		if testIdx < 0 {

			testIdx = integerDigits.GetRuneArrayLength() - 1

			if testIdx < 0 {

				err = fmt.Errorf("%v\n"+
					"Error: 'integerDigits' array is zero length and empty!\n",
					ePrefix.String())

				return err
			}

			testDigit =
				integerDigits.CharsArray[testIdx]

		} else {

			testDigit =
				fractionalDigits.CharsArray[testIdx]

		}

		remainder := (int(testDigit-48) + 1) % 2

		if remainder == 0 {

			roundUp = true

		} else {

			roundUp = false
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

// roundHalfToOdd
//
// Performs a rounding operation on the integer and fractional
// numeric digit arrays contained in two instances of
// RuneArrayDto passed as input parameters.
//
// Round Half to Odd Numbers is Similar to 'HalfToEven',
// but in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
// This method performs and applies the 'Round To Odd'
// algorithm.
//
//	Examples of HalfToOdd
//
//		HalfToOdd only applies to 0.5. Other numbers
//		(not ending in 0.5) round to nearest as usual.
//
//		7.5 rounds down to 7 (because 8 is an even number)
//
//		6.5 rounds up to 7 (because 7 is an odd number)
//
//		7.6 rounds up to 8
//		7.5 rounds down to 7 (because 8 is an even number)
//		7.4 rounds down to 7
//		6.6 rounds up to 7
//		6.5 rounds up to 7 (because 7 is an odd number)
//		6.4 rounds down to 6
//		-23.5 rounds down to -23 (because 4 is an even number)
//		-24.5 rounds up to -25 (because 5 is an odd number)
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.HalfToOdd()
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
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) roundHalfToOdd(
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
			"roundHalfToOdd()",
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

	if fractionalDigits.CharsArray[roundToFractionalDigits] <
		'5' {

		// ROUND DOWN!
		roundUp = false

	} else if fractionalDigits.CharsArray[roundToFractionalDigits] >
		'5' {
		// ROUND UP!
		roundUp = true

	} else {
		// MUST BE
		// fractionalDigits.CharsArray[roundToFractionalDigits] ==
		// '5'

		testIdx := roundToFractionalDigits - 1

		var testDigit rune

		if testIdx < 0 {

			testIdx = integerDigits.GetRuneArrayLength() - 1

			if testIdx < 0 {

				err = fmt.Errorf("%v\n"+
					"Error: 'integerDigits' array is zero length and empty!\n",
					ePrefix.String())

				return err
			}

			testDigit =
				integerDigits.CharsArray[testIdx]

		} else {

			testDigit =
				fractionalDigits.CharsArray[testIdx]

		}

		remainder := (int(testDigit-48) + 1) % 2

		if remainder == 0 {

			roundUp = false

		} else {

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

// roundRandomly
//
// Round Half Randomly. Uses a Random Number Generator
// to choose between rounding 0.5 up or down.
//
// All numbers other than 0.5 round to the nearest as usual.
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.Randomly()
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
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) roundRandomly(
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
			"roundRandomly()",
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

	if fractionalDigits.CharsArray[roundToFractionalDigits] <
		'5' {

		// ROUND DOWN!
		roundUp = false

	} else if fractionalDigits.CharsArray[roundToFractionalDigits] >
		'5' {
		// ROUND UP!
		roundUp = true

	} else {
		// MUST BE
		// fractionalDigits.CharsArray[roundToFractionalDigits] ==
		// '5'

		randomNum := new(numStrMathQuark).randomInt()

		remainder := randomNum % 2

		if remainder == 0 {

			roundUp = true

		} else {

			roundUp = false
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

// truncate
//
// Applies NO Rounding whatsoever. The Round From Digit is
// dropped or deleted. The Round To Digit is NEVER changed.
//
//	Examples of Truncate
//
//	Example-1
//		Number: 23.14567
//		Objective: Round to two decimal places to
//		           the right of the decimal point.
//		Rounding Method:	Truncate
//		Round To Digit:		4
//		Round From Digit:	5
//		Rounded Number:		23.14 - The Round From Digit
//									is dropped.
//
//	Example-2
//		Number:				-23.14567
//		Objective:			Round to two decimal places to
//							the right of the decimal point.
//		Rounding Method:	Truncate
//		Round To Digit:		4
//		Round From Digit:	5
//		Rounded Number:		-23.14 - The Round From Digit
//							is dropped.
//
// This algorithm corresponds to the following NumberRoundingType:
//
//	NumRoundType.Truncate()
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
//		digits used to construct the base integer numeric
//		value to which the rounding algorithm is applied.
//
//		These integer numeric digits comprise the integer
//		portion of an integer or floating point numeric
//		value.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance contains the rune array of fractional
//		numeric digits used to construct the base floating
//		point numeric value to which the rounding algorithm
//		is applied.
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
//		This parameter designates the number sign of the
//		numeric value represented by the integer and fractional
//		digits contained in input parameters, 'integerDigits'
//		and 'fractionalDigits'.
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
func (nStrMathRoundAtom *numStrMathRoundingAtom) truncate(
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
			"truncate()",
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

		// Nothing to do
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

	// Truncate Fractional Digits array!
	// Last Index = roundToFractionalDigits - 1

	fractionalDigits.CharsArray =
		fractionalDigits.CharsArray[:roundToFractionalDigits]

	return err
}

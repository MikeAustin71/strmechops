package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelAtom - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelAtom struct {
	lock *sync.Mutex
}

//	addFractionalDigit
//
//	Appends a single numeric digit to the end of the internal
//	fractional digits rune array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance contains the fractional digit rune array
//		to which the 'fractionalDigit' rune will be appended.
//
//	fractionalDigit				rune
//
//		A rune with a numeric character between '0' (zero)
//		and '9' (nine) inclusive. This numeric digit will
//		be appended to the end of the internal member
//		variable 'NumberStrKernel.fractionalDigits'
//		contained within the NumberStrKernel input
//		parameter, 'numStrKernel'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelAtom *numberStrKernelAtom) addFractionalDigit(
	numStrKernel *NumberStrKernel,
	fractionalDigit rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"addFractionalDigit()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if fractionalDigit < '0' ||
		fractionalDigit > '9' {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fractionalDigit' is invalid!\n"+
			"Fractional Rune characters must represent a numberic character between\n"+
			"'0' and '9', inclusive. 'fractionalDigit' fails to meet this criterion.\n"+
			"The rune value of 'fractionalDigit' is %v\n"+
			"The string value of 'fractionalDigit' is %v\n",
			ePrefix.String(),
			fractionalDigit,
			string(fractionalDigit))

		return err
	}

	err = numStrKernel.fractionalDigits.AddChar(
		fractionalDigit,
		true,
		ePrefix.XCpy(
			"numStrKernel.fractionalDigits<-"))

	if err != nil {
		return err
	}

	if numStrKernel.numberValueType !=
		NumValType.FloatingPoint() {

		numStrKernel.numberValueType =
			NumValType.FloatingPoint()
	}

	if fractionalDigit != '0' {
		numStrKernel.isNonZeroValue = true
	}

	return err
}

// addIntegerDigit - Adds a single numeric digit to the internal
// integer digits rune array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance contains the integer digit rune array to
//		which the 'integerDigit' rune will be appended.
//
//	integerDigit            rune
//
//		A rune with a numeric character between '0' (zero) and
//		'9' (nine) inclusive. This numeric digit will be
//		appended to the internal member variable
//		'NumberStrKernel.integerDigits' for NumberStrKernel
//		input parameter 'numStrKernel'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelAtom *numberStrKernelAtom) addIntegerDigit(
	numStrKernel *NumberStrKernel,
	integerDigit rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"addIntegerDigit()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if integerDigit < '0' ||
		integerDigit > '9' {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigit' is invalid!\n"+
			"Integer Runes must represent a numberic character between\n"+
			"'0' and '9', inclusive. 'integerDigit' fails to meet this criterion.\n"+
			"The rune value of 'integerDigit' is %v\n"+
			"The string value of 'integerDigit' is %v\n",
			ePrefix.String(),
			integerDigit,
			string(integerDigit))

		return err
	}

	err = numStrKernel.integerDigits.AddChar(
		integerDigit,
		true,
		ePrefix.XCpy(
			"numStrKernel.integerDigits<-"))

	if err != nil {
		return err
	}

	if numStrKernel.numberValueType !=
		NumValType.FloatingPoint() {

		numStrKernel.numberValueType =
			NumValType.Integer()
	}

	if integerDigit != '0' {
		numStrKernel.isNonZeroValue = true
	}

	return err
}

//	calcNumStrKernelStats
//
//	Receives a pointer to an instance of NumberStrKernel
//	and proceeds to analyze that instance to produce
//	statistical information on the encapsulated numeric
//	value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		analyzed to produce a statistics on the nature of
//		the numeric value encapsulated by this instance.
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
//	numStrStatsDto				NumberStrStatsDto
//
//		This data transfer object will return key
//		statistics on the numeric value encapsulated
//		by input parameter, 'numStrKernel'.
//
//		type NumberStrStatsDto struct {
//
//		NumOfIntegerDigits					uint64
//
//			The total number of integer digits to the
//			left of the radix point or, decimal point, in
//			the subject numeric value.
//
//		NumOfSignificantIntegerDigits		uint64
//
//			The number of nonzero integer digits to the
//			left of the radix point or, decimal point, in
//			the subject numeric value.
//
//		NumOfFractionalDigits				uint64
//
//			The total number of fractional digits to the
//			right of the radix point or, decimal point,
//			in the subject numeric value.
//
//		NumOfSignificantFractionalDigits	uint64
//
//			The number of nonzero fractional digits to
//			the right of the radix point or, decimal
//			point, in the subject numeric value.
//
//		NumberValueType 					NumericValueType
//
//			This enumeration value specifies whether the
//			subject numeric value is classified either as
//			an integer or a floating point number.
//
//			Possible enumeration values are listed as
//			follows:
//				NumValType.None()
//				NumValType.FloatingPoint()
//				NumValType.Integer()
//
//		NumberSign							NumericSignValueType
//
//			An enumeration specifying the number sign
//			associated with the numeric value. Possible
//			values are listed as follows:
//				NumSignVal.None()		= Invalid Value
//				NumSignVal.Negative()	= -1
//				NumSignVal.Zero()		=  0
//				NumSignVal.Positive()	=  1
//
//		IsZeroValue							bool
//
//			If 'false', the subject numeric value is
//			greater than or less than zero ('0').
//
//			If 'true', the Numeric Value is equal to
//			zero.
//		}
//
//	error
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
func (numStrKernelAtom *numberStrKernelAtom) calcNumStrKernelStats(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	numStrStatsDto NumberStrStatsDto,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"calcNumStrKernelStats()",
		"")

	if err != nil {

		return numStrStatsDto, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return numStrStatsDto, err
	}

	numStrStatsDto.NumOfIntegerDigits =
		uint64(len(numStrKernel.integerDigits.CharsArray))

	lenZeros := numStrKernel.integerDigits.GetCountLeadingZeros()

	numStrStatsDto.NumOfSignificantIntegerDigits =
		numStrStatsDto.NumOfIntegerDigits - lenZeros

	numStrStatsDto.NumOfFractionalDigits =
		uint64(len(numStrKernel.fractionalDigits.CharsArray))

	lenZeros = numStrKernel.fractionalDigits.GetCountTrailingZeros()

	numStrStatsDto.NumOfSignificantFractionalDigits =
		numStrStatsDto.NumOfFractionalDigits - lenZeros

	numStrStatsDto.NumberValueType,
		err = new(numberStrKernelQuark).getSetNumValueType(
		numStrKernel,
		ePrefix.XCpy(
			"numStrStatsDto.NumberValueType<-"+
				"numStrKernel"))

	if err != nil {

		return numStrStatsDto, err

	}

	numStrStatsDto.NumberSign =
		numStrKernel.numberSign

	numStrStatsDto.IsZeroValue,
		err = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {

		return numStrStatsDto, err

	}

	numStrStatsDto.IsZeroValue = !numStrStatsDto.IsZeroValue

	return numStrStatsDto, err
}

//	emptyFractionalDigits
//
//	Receives an instance of NumberStrKernel and proceeds
//	to set the internal fractional digits array to 'nil'.
//	This effectively deletes the fractional part of the
//	numeric value contained within the passed instance
//	of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance contains the fractional digit rune array
//		which will be deleted and set to 'nil'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelAtom *numberStrKernelAtom) emptyFractionalDigits(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"emptyFractionalDigits()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	numStrKernel.fractionalDigits.Empty()

	numStrKernel.fractionalDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	var isNonZero bool

	isNonZero,
		err = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		ePrefix)

	if err != nil {

		return err
	}

	if isNonZero == false {

		numStrKernel.numberSign = NumSignVal.Zero()

		numStrKernel.numberValueType = NumValType.None()

		return err
	}

	// integer digits exist

	numStrKernel.numberValueType = NumValType.Integer()

	return err
}

//	emptyIntegerDigits
//
//	Receives an instance of NumberStrKernel and proceeds to set
//	the internal integer digits array to 'nil'. This effectively
//	deletes the integer part of the numeric value contained
//	within the passed instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance contains the integer digit rune array
//		which will be deleted and set to 'nil'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelAtom *numberStrKernelAtom) emptyIntegerDigits(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"emptyIntegerDigits()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	numStrKernel.integerDigits.Empty()

	numStrKernel.integerDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	var isNonZero bool

	isNonZero,
		err = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		ePrefix)

	if err != nil {

		return err
	}

	if isNonZero == false {

		numStrKernel.numberSign = NumSignVal.Zero()

		numStrKernel.numberValueType = NumValType.None()

		return err
	}

	// fractional digits exist

	numStrKernel.numberValueType = NumValType.FloatingPoint()

	return err
}

//	formatNumStrElements
//
//	Creates and returns a fully formatted Number String
//	generated from Number String formatting components
//	passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		formatted and returned as a Number String.
//
//	roundingSpec 				NumStrRoundingSpec
//
//		The Number String Rounding Specification
//		contains all the parameters required to
//		configure a rounding algorithm for a
//		floating point number string.
//
//		type NumStrRoundingSpec struct {
//
//			roundingType NumberRoundingType
//
//			This enumeration parameter is used to specify the type
//			of rounding algorithm that will be applied for the
//			rounding of fractional digits in a number string.
//
//			Possible values are listed as follows:
//				NumRoundType.None()
//				NumRoundType.NoRounding()
//				NumRoundType.HalfUpWithNegNums()
//				NumRoundType.HalfDownWithNegNums()
//				NumRoundType.HalfAwayFromZero()
//				NumRoundType.HalfTowardsZero()
//				NumRoundType.HalfToEven()
//				NumRoundType.HalfToOdd()
//				NumRoundType.Randomly()
//				NumRoundType.Floor()
//				NumRoundType.Ceiling()
//				NumRoundType.Truncate()
//
//			NoRounding
//
//				Signals that no rounding operation will be performed
//				on fractional digits contained in a number string.
//				The fractional digits will therefore remain unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers. This method
//				is intuitive but may produce unexpected results when
//				applied to negative numbers.
//
//				'HalfUpWithNegNums' rounds .5 up.
//
//					Examples of 'HalfUpWithNegNums'
//					7.6 rounds up to 8
//					7.5 rounds up to 8
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds up to -7
//					-7.6 rounds down to -8
//
//			HalfDownWithNegNums
//
//				Half Round Down Including Negative Numbers. This method
//				is also considered intuitive but may produce unexpected
//				results when applied to negative numbers.
//
//				'HalfDownWithNegNums' rounds .5 down.
//
//					Examples of HalfDownWithNegNums
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7
//					7.4 rounds down to 7
//					-7.4 rounds up to -7
//					-7.5 rounds down to -8
//					-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				Round Half Away From Zero. This rounding method is treated
//				as the default and this value is returned by method:
//				NumberRoundingType(0).XGetDefaultRoundingType()
//
//				The 'HalfAwayFromZero' method rounds .5 further away from zero.
//				It provides clear and consistent behavior when dealing with
//				negative numbers.
//
//					Examples of HalfAwayFromZero
//
//					7.6 rounds away to 8
//					7.5 rounds away to 8
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds away to -8
//					-7.6 rounds away to -8
//
//			HalfTowardsZero
//
//				Round Half Towards Zero. 'HalfTowardsZero' rounds 0.5
//				closer to zero. It provides clear and consistent behavior
//				when dealing with negative numbers.
//
//					Examples of HalfTowardsZero
//
//					7.6 rounds away to 8
//					7.5 rounds to 7
//					7.4 rounds to 7
//					-7.4 rounds to -7
//					-7.5 rounds to -7
//					-7.6 rounds away to -8
//
//			HalfToEven
//
//				Round Half To Even Numbers. 'HalfToEven' is also called
//				Banker's Rounding. This method rounds 0.5 to the nearest
//				even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even number)
//					but 6.5 rounds down to 6 (because 6 is an even number)
//
//					HalfToEven only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual, so:
//
//					7.6 rounds up to 8
//					7.5 rounds up to 8 (because 8 is an even number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds down to 6 (because 6 is an even number)
//					6.4 rounds down to 6
//
//			HalfToOdd
//
//				Round Half to Odd Numbers. Similar to 'HalfToEven', but
//				in this case 'HalfToOdd' rounds 0.5 towards odd numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers (not ending
//					in 0.5) round to nearest as usual.
//
//					7.5 rounds down to 7 (because 7 is an odd number)
//
//					6.5 rounds up to 7 (because 7 is an odd number)
//
//					7.6 rounds up to 8
//					7.5 rounds down to 7 (because 7 is an odd number)
//					7.4 rounds down to 7
//					6.6 rounds up to 7
//					6.5 rounds up to 7 (because 7 is an odd number)
//					6.4 rounds down to 6
//
//			Randomly
//
//				Round Half Randomly. Uses a Random Number Generator to choose
//				between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply any
//				special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than or
//				equal to x
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function is
//				the function that takes as input a real number x, and gives
//				as output the greatest integer less than or equal to x,
//				denoted floor(x) or ⌊x⌋.
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Floor
//
//					Number     Floor
//					2           2
//					2.4         2
//					2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not apply any
//				special treatment to 0.5.
//
//				Ceiling Function: The least integer that is greater than or
//				equal to x.
//				Source: https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer greater than
//				or equal to x, denoted ceil(x) or ⌈x⌉.[1]
//				Source: https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//					Number    Ceiling
//					2           2
//					2.4         3
//					2.9         3
//					-2.5        -2
//					-2.7        -2
//					-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit is dropped
//				or deleted. The Round To Digit is NEVER changed.
//
//					Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit is dropped.
//
//			roundToFractionalDigits int
//
//				When set to a positive integer value, this
//				parameter controls the number of digits to
//				the right of the radix point or decimal
//				separator (a.k.a. decimal point) which will
//				remain after completion of the number rounding
//				operation.
//		}
//
//	decSeparatorSpec				DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Number String Integer Separator Specification. This
//		type encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//		Type IntegerSeparatorSpec is designed to manage
//		integer separators, primarily thousands separators,
//		for different countries and cultures. The term
//		'integer separators' is used because this type
//		manages both integer grouping and the characters
//		used to separate integer groups.
//
//		In the USA and many other countries, integer
//		numbers are often separated by commas thereby
//		grouping the number into thousands.
//
//		Examples:
//
//			IntGroupingType.None()
//			(a.k.a Integer Separation Turned Off)
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
//
//		Other countries and cultures use characters other
//		than the comma to separate integers into thousands.
//		Some countries and cultures do not use thousands
//		separation and instead rely on multiple integer
//		separation characters and grouping sequences for a
//		single integer number. Notable examples of this
//		are found in the 'India Number System' and
//		'Chinese Numerals'.
//
//		Reference:
//			https://en.wikipedia.org/wiki/Indian_numbering_system
//			https://en.wikipedia.org/wiki/Chinese_numerals
//			https://en.wikipedia.org/wiki/Decimal_separator
//
//		The IntegerSeparatorSpec type provides the
//		flexibility necessary to process these complex
//		number separation formats.
//
//		If integer separation is turned off, no error
//		will be returned and integer digits will be
//		displayed as a single string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	negativeNumberSign				NumStrNumberSymbolSpec
//
//		The Number String Negative Number Sign
//		Specification is used to configure negative
//		number sign symbols for negative numeric
//		values formatted and displayed in number
//		stings.
//
//		If this parameter is submitted as an empty or
//		invalid Negative Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	positiveNumberSign 				NumStrNumberSymbolSpec
//
//		Positive number signs are commonly implied
//		and not specified. However, the user has
//		the option to specify a positive number sign
//		character or characters for positive numeric
//		values using this input parameter.
//
//		If this parameter is submitted as an empty or
//		invalid Positive Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//	zeroNumberSign					NumStrNumberSymbolSpec
//
//		The Number String Zero Number Sign
//		Specification is used to configure number
//		sign symbols for zero numeric values formatted
//		and displayed in number stings. Zero number signs
//		are commonly omitted because zero does not
//		technically qualify as either a positive or
//		negative value. However, the user has the option
//		to configure number sign symbols for zero values
//		if necessary.
//
//		If this parameter is submitted as an empty or
//		invalid Zero Number Sign Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+ "
//			Trailing Symbols: ""
//			Number String:   "+ 0.00"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: "+"
//			Number String:   "0.00+"
//
//	currencySymbol					NumStrNumberSymbolSpec
//
//		A Currency Symbol next to a number shows the
//		number is a monetary amount.
//
//		The Number String Currency Symbol Specification
//		is used to configure currency symbols for
//		positive, negative and zero numeric values
//		formatted and displayed in number stings.
//
//		If this parameter is submitted as an empty or
//		invalid Currency Symbol Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Examples of Currency Symbols include the Dollar
//		sign ('$'), Euro sign ('€') or Pound sign ('£').
//
//		This instance of NumStrNumberSymbolSpec is used
//		to configure leading Currency Symbols, trailing
//		Currency Symbols or both leading and trailing
//		Currency Symbols.
//
//		Example-1: Leading Currency Symbols
//
//			Leading Currency Symbols: "$ "
//			Number String:   "$ 123.456"
//
//		Example-2: Leading Currency Symbols
//
//			Leading Currency Symbols: "$"
//			Number String:   "$123.456"
//
//		Example-3: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: "€"
//			Number String:   "123.456€"
//
//		Example-4: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: " €"
//			Number String:   "123.456 €"
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
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
//	numStr						string
//
//		If this method completes successfully, the
//		numeric	value represented by input parameters
//		'integerDigits' and 'fractionalDigits' will be
//		returned as a formatted Number String, 'numStr'.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (numStrKernelAtom *numberStrKernelAtom) formatNumStrElements(
	numStrKernel *NumberStrKernel,
	roundingSpec NumStrRoundingSpec,
	decSeparator DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	currencySymbol NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	numStr string,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"formatNumStrElements()",
		"")

	if err != nil {

		return numStr, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return numStr, err
	}

	if len(numStrKernel.integerDigits.CharsArray) == 0 &&
		len(numStrKernel.fractionalDigits.CharsArray) == 0 {

		numStr = "0"

		return numStr, err
	}

	err = roundingSpec.IsValidInstanceError(
		ePrefix.XCpy(
			"roundingSpec"))

	if err != nil {

		return numStr, err
	}

	var newNumStrKernel NumberStrKernel

	err = new(numberStrKernelNanobot).copy(
		&newNumStrKernel,
		numStrKernel,
		ePrefix.XCpy(
			"newNumStrKernel<-numStrKernel"))

	if err != nil {
		return numStr, err
	}

	// Performing fractional digit rounding
	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&newNumStrKernel,
		roundingSpec,
		ePrefix.XCpy(
			"newNumStrKernel Rounding"))

	if err != nil {
		return numStr, err
	}

	var numOfFracDigits int

	numOfFracDigits = newNumStrKernel.GetNumberOfFractionalDigits()

	if numOfFracDigits > 0 &&
		decSeparator.GetNumberOfSeparatorChars() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This is a floating point number and the number\n"+
			"of decimal separator characters specified is zero.\n"+
			"Input parameter 'decSeparator'\n"+
			"is invalid!\n",
			ePrefix.String())

		return numStr, err
	}

	return new(numStrHelperNanobot).formatNumStrElements(
		&newNumStrKernel.integerDigits,
		&newNumStrKernel.fractionalDigits,
		newNumStrKernel.numberSign,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		currencySymbol,
		numberFieldSpec,
		ePrefix.XCpy(
			"<-newNumStrKernel"))
}

/*
func (numStrKernelAtom *numberStrKernelAtom) formatNumStrElements(
	numStrKernel *NumberStrKernel,
	decSeparator DecimalSeparatorSpec,
	intSeparatorDto IntegerSeparatorSpec,
	roundingSpec NumStrRoundingSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	numStr string,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"formatNumStrElements()",
		"")

	if err != nil {

		return numStr, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return numStr, err
	}

	if len(numStrKernel.integerDigits.CharsArray) == 0 &&
		len(numStrKernel.fractionalDigits.CharsArray) == 0 {

		numStr = "0"

		return numStr, err
	}

	err = roundingSpec.IsValidInstanceError(
		ePrefix.XCpy(
			"roundingSpec"))

	if err != nil {

		return numStr, err
	}

	var newNumStrKernel NumberStrKernel

	err = new(numberStrKernelNanobot).copy(
		&newNumStrKernel,
		numStrKernel,
		ePrefix.XCpy(
			"newNumStrKernel<-numStrKernel"))

	if err != nil {
		return numStr, err
	}

	// Performing fractional digit rounding
	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&newNumStrKernel,
		roundingSpec,
		ePrefix.XCpy(
			"newNumStrKernel Rounding"))

	if err != nil {
		return numStr, err
	}

	var numOfFracDigits int

	numOfFracDigits = newNumStrKernel.GetNumberOfFractionalDigits()

	if numOfFracDigits > 0 &&
		decSeparator.GetNumberOfSeparatorChars() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This is a floating point number and the number\n"+
			"of decimal separator characters specified is zero.\n"+
			"Input parameter 'decSeparator'\n"+
			"is invalid!\n",
			ePrefix.String())

		return numStr, err
	}

	var numStrWithIntSeps []rune

	numStrWithIntSeps,
		err = new(integerSeparatorSpecMolecule).applyIntSeparators(
		&intSeparatorDto,
		newNumStrKernel.GetIntegerRuneArray(),
		ePrefix.XCpy("intSeparatorDto"))

	if err != nil {
		return numStr, err
	}

	tempNumStr := string(numStrWithIntSeps)

	if numOfFracDigits > 0 {

		tempNumStr += decSeparator.GetDecimalSeparatorStr()

		tempNumStr += newNumStrKernel.GetFractionalString()

	}

	leadingNumSym := ""

	trailingNumSym := ""

	var leadingNumSymPosition, trailingNumSymPosition NumberFieldSymbolPosition

	if newNumStrKernel.numberSign == NumSignVal.Negative() {

		if !negativeNumberSign.IsNOP() {

			leadingNumSym =
				negativeNumberSign.GetLeadingNumberSymbolStr()

			leadingNumSymPosition =
				negativeNumberSign.GetLeadingNumberSymbolPosition()

			trailingNumSym =
				negativeNumberSign.GetTrailingNumberSymbolStr()

			trailingNumSymPosition =
				negativeNumberSign.GetTrailingNumberSymbolPosition()

		}

	}

	if newNumStrKernel.numberSign == NumSignVal.Positive() {

		if !positiveNumberSign.IsNOP() {

			leadingNumSym =
				positiveNumberSign.GetLeadingNumberSymbolStr()

			leadingNumSymPosition =
				positiveNumberSign.GetLeadingNumberSymbolPosition()

			trailingNumSym =
				positiveNumberSign.GetTrailingNumberSymbolStr()

			trailingNumSymPosition =
				positiveNumberSign.GetTrailingNumberSymbolPosition()

		}

	}

	if newNumStrKernel.numberSign == NumSignVal.Zero() {

		if !zeroNumberSign.IsNOP() {

			leadingNumSym =
				zeroNumberSign.GetLeadingNumberSymbolStr()

			leadingNumSymPosition =
				zeroNumberSign.GetLeadingNumberSymbolPosition()

			trailingNumSym =
				zeroNumberSign.GetTrailingNumberSymbolStr()

			trailingNumSymPosition =
				zeroNumberSign.GetTrailingNumberSymbolPosition()

		}

	}

	lenLeadingNumSymbol := len(leadingNumSym)
	lenTrailingNumSymbol := len(trailingNumSym)

	if lenLeadingNumSymbol > 0 &&
		leadingNumSymPosition == NumFieldSymPos.InsideNumField() {

		tempNumStr = leadingNumSym + tempNumStr
	}

	if lenTrailingNumSymbol > 0 &&
		trailingNumSymPosition == NumFieldSymPos.InsideNumField() {

		tempNumStr = tempNumStr + trailingNumSym

	}

	numStr,
		err = new(strMechNanobot).justifyTextInStrField(
		tempNumStr,
		numberFieldSpec.GetNumFieldLength(),
		numberFieldSpec.GetNumFieldJustification(),
		ePrefix.XCpy("numStr<-tempNumStr"))

	if err != nil {

		return numStr, err
	}

	if lenLeadingNumSymbol > 0 &&
		leadingNumSymPosition == NumFieldSymPos.OutsideNumField() {

		numStr = leadingNumSym + numStr
	}

	if lenTrailingNumSymbol > 0 &&
		trailingNumSymPosition == NumFieldSymPos.OutsideNumField() {

		numStr = numStr + trailingNumSym

	}

	return numStr, err
}
*/

//	prepareCompareNumStrKernels
//
//	This method receives pointers to two instances of
//	NumberStrKernel, 'numStrKernel01' and
//	'numStrKernel02'.
//
//	These two instances will be revamped to ensure that
//	their integer digits and fractional digits are of
//	equal lengths. Next, this method will proceed to
//	compare the numeric values of 'numStrKernel01' and
//	'numStrKernel02'. The comparison results will be
//	returned as one of three integer values:
//
//		-1	= numStrKernel01 is less than numStrKernel02
//		 0	= numStrKernel01 is equal to numStrKernel02
//		+1	= numStrKernel01 is greater than numStrKernel02
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel01				*NumberStrKernel
//
//		The numeric value of numStrKernel01 will be
//		compared to that of numStrKernel02. The
//		comparison results will be returned as an integer
//		value.
//
//	numStrKernel02				*NumberStrKernel
//
//		The numeric value of numStrKernel01 will be
//		compared to that of this parameter,
//		numStrKernel02. The comparison results will be
//		returned as an integer value.
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
//	comparisonValue				int
//
//		This parameter will return the results of numeric
//		value comparisons for input parameters,
//		'numStrKernel01' and 'numStrKernel02'. The
//		integer comparison result will be set to one of
//		three values:
//
//		-1	= numStrKernel01 is less than numStrKernel02
//		 0	= numStrKernel01 is equal to numStrKernel02
//		+1	= numStrKernel01 is greater than numStrKernel02
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
func (numStrKernelAtom *numberStrKernelAtom) prepareCompareNumStrKernels(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	comparisonValue int,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"calcNumStrKernelStats()",
		"")

	if err != nil {

		return comparisonValue, err

	}

	if numStrKernel01 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel01' is a nil pointer!\n",
			ePrefix.String())

		return comparisonValue, err
	}

	if numStrKernel02 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel02' is a nil pointer!\n",
			ePrefix.String())

		return comparisonValue, err
	}

	var numStrKernel01V2, numStrKernel02V2 NumberStrKernel

	numStrKernel01V2,
		err = numStrKernel01.CopyOut(
		ePrefix.XCpy(
			"numStrKernel01V2<-numStrKernel01"))

	if err != nil {

		return comparisonValue, err

	}

	numStrKernel02V2,
		err = numStrKernel02.CopyOut(
		ePrefix.XCpy(
			"numStrKernel02V2<-numStrKernel02"))

	if err != nil {

		return comparisonValue, err

	}

	lenIntDigits01V2 :=
		len(numStrKernel01V2.integerDigits.CharsArray)

	lenFracDigits01V2 :=
		len(numStrKernel01V2.fractionalDigits.CharsArray)

	lenIntDigits02V2 :=
		len(numStrKernel02V2.integerDigits.CharsArray)

	lenFracDigits02V2 :=
		len(numStrKernel02V2.fractionalDigits.CharsArray)

	numStrKernelElectron := numberStrKernelElectron{}

	if lenIntDigits01V2 != lenIntDigits02V2 ||
		lenFracDigits01V2 != lenFracDigits02V2 {

		err = numStrKernelElectron.
			equalizeNumStrDigitsLengths(
				&numStrKernel01V2,
				&numStrKernel02V2,
				ePrefix.XCpy(
					"numStrKernel01V2 - numStrKernel02V2"))

		if err != nil {

			return comparisonValue, err

		}
	}

	// Integer Digits and Fractional Digits are NOW
	// equivalent between:
	//    numStrKernel01V2 and numStrKernel02V2

	comparisonValue,
		err = new(numberStrKernelQuark).
		compareNumStrKernelValues(
			&numStrKernel01V2,
			&numStrKernel02V2,
			ePrefix.XCpy(
				"numStrKernel01V2 & numStrKernel02V2"))

	return comparisonValue, err
}

// testValidityOfNumStrKernel - Receives a pointer to an instance
// of NumberStrKernel and performs a diagnostic analysis to
// determine if that instance is valid in all respects.
//
// If the input parameter 'numStrKernel' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'numStrKernel' is valid, this method will
// return a boolean flag ('isValid') of 'true' and the returned
// error type ('err') will be set to 'nil'.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		object will be subjected to diagnostic analysis in
//		order to determine if all the member variables
//		contain valid values.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	isValid						bool
//
//		If input parameter 'numStrKernel' is judged to be valid in
//		all respects, this return parameter will be set to 'true'.
//
//		If input parameter 'numStrKernel' is found to be invalid,
//		this return parameter will be set to 'false'.
//
//	err							error
//
//		If this method completes successfully and
//	 	'numStrKernel' evaluates as valid, this returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered or if 'numStrKernel'
//		evaluates as invalid, the returned error Type
//		will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernelAtom *numberStrKernelAtom) testValidityOfNumStrKernel(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"testValidityOfNumStrKernel()",
		"")

	if err != nil {

		return isValid, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	lenIntDigits := numStrKernel.integerDigits.GetRuneArrayLength()

	lenFracDigits := numStrKernel.fractionalDigits.GetRuneArrayLength()

	if lenIntDigits == 0 &&
		lenFracDigits == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!\n"+
			"Both Integer Digits and Fractional Digits are empty\n"+
			"and contain zero digits.\n",
			ePrefix.String())

		return isValid, err
	}

	if lenIntDigits == 0 &&
		lenFracDigits > 0 {

		err = numStrKernel.AddIntegerDigit(
			'0',
			ePrefix.XCpy("Adding Missing Zero Digit"))

		if err != nil {
			return isValid, err
		}
	}

	if !numStrKernel.numberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!\n"+
			"The Number Sign Value is invalid.\n"+
			"Valid Number Sign Values are:\n"+
			"   NumSignVal.Negative()\n"+
			"   NumSignVal.Zero()\n"+
			"   NumSignVal.Positive()\n"+
			"The current Number Sign Value is:\n"+
			"   Number Sign String Value = '%v'\n"+
			"  Number Sign Integer Value = '%v\n",
			ePrefix.String(),
			numStrKernel.numberSign.String(),
			numStrKernel.numberSign.XValueInt())

		return isValid, err
	}

	numValHasNonZeroVal := false

	for i := 0; i < lenIntDigits; i++ {

		if numStrKernel.integerDigits.CharsArray[i] < '0' &&
			numStrKernel.integerDigits.CharsArray[i] > '9' {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!\n"+
				"The Integer Digits rune array contains non-numeric characters.\n",
				ePrefix.String())

			return isValid, err
		}
	}

	for i := 0; i < lenIntDigits; i++ {

		if numStrKernel.integerDigits.CharsArray[i] >= '1' &&
			numStrKernel.integerDigits.CharsArray[i] <= '9' {
			numValHasNonZeroVal = true
			break
		}
	}

	for j := 0; j < lenFracDigits; j++ {

		if numStrKernel.fractionalDigits.CharsArray[j] < '0' &&
			numStrKernel.fractionalDigits.CharsArray[j] > '9' {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!\n"+
				"The Fractional Digits rune array contains non-numeric characters.\n",
				ePrefix.String())

			return isValid, err
		}

	}

	if !numValHasNonZeroVal {

		for j := 0; j < lenFracDigits; j++ {

			if numStrKernel.fractionalDigits.CharsArray[j] >= '1' &&
				numStrKernel.fractionalDigits.CharsArray[j] <= '9' {
				numValHasNonZeroVal = true
				break
			}

		}

	}

	if numValHasNonZeroVal &&
		numStrKernel.numberSign == NumSignVal.Zero() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!\n"+
			"The Number Sign Value is invalid.\n"+
			"NumberStrKernel has a non-zero numeric value.\n"+
			"However, Number Sign is equal to Zero.\n"+
			"Number Sign = NumSignVal.Zero()\n",
			ePrefix.String())

		return isValid, err
	}

	if numValHasNonZeroVal != numStrKernel.isNonZeroValue {

		if numValHasNonZeroVal == false &&
			numStrKernel.isNonZeroValue == true {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!\n"+
				"The Number Sign Value is invalid.\n"+
				"NumberStrKernel has a zero numeric value.\n"+
				"However, internal flag numStrKernel.isNonZeroValue\n"+
				"is set to 'true'.\n",
				ePrefix.String())

			return isValid, err

		}

		if numValHasNonZeroVal == true &&
			numStrKernel.isNonZeroValue == false {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!\n"+
				"The Number Sign Value is invalid.\n"+
				"NumberStrKernel has a non-zero numeric value.\n"+
				"However, internal flag numStrKernel.isNonZeroValue\n"+
				"is set to 'false'.\n",
				ePrefix.String())

			return isValid, err

		}
	}

	if new(numStrFmtSpecNanobot).isNOP(
		&numStrKernel.numStrFormatSpec) {

		// This is a NOP!
		// numStrKernel.numStrFormatSpec is invalid.
		// Set default numStrKernel.numStrFormatSpec
		// to stand US Signed Number String Format
		// Specification.
		numStrKernel.numStrFormatSpec,
			err = new(NumStrFormatSpec).NewSignedNumDefaultsUSMinus(
			NumStrNumberFieldSpec{
				fieldLength:        -1,
				fieldJustification: TxtJustify.Right(),
			},
			ePrefix.XCpy(
				"numStrKernel.numStrFormatSpec "+
					"Default US NumStrNumberFieldSpec"))

		if err != nil {
			return isValid, err
		}
	}

	isValid = true

	return isValid, err
}

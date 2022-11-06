package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

// numberStrKernelAtom - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelAtom struct {
	lock *sync.Mutex
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
//		input parameter, 'numStrKernel'.
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

//	convertKernelToBigFloat
//
//	Converts an instance of NumberStrKernel to a floating
//	point numeric value of type *big.Float.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		converted to a floating point number of type
//		*big.Float.
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		'roundingType' is only applied in cases where the
//		current NumberStrKernel instance consists of a
//		floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
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
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
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
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
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
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
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
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
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
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
//
//	roundToFractionalDigits int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		radix point or decimal separator (a.k.a.
//		decimal point).
//
//		After completion of a number rounding operation, the
//		value of roundToFractionalDigits will be equal to the
//		number of digits to the right of the decimal point.
//
//		If this parameter is set to a value less than zero,
//		an error will be returned.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	*big.Float
//
//		If this method completes successfully, the
//		numeric	value represented by the NumberStrKernel
//		instance, 'numStrKernel', will be returned as a
//		type *big.Float.
//
//	int
//
//		The number of fractional digits contained in the
//		numeric value passed through return parameter
//		'*big.Float'.
//
//		For returned type *big.Float, the number of
//		fractional digits can be used to improve accuracy
//		in conversions from *big.Float to character strings.
//
//	error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If errors
//		are	encountered during processing, the returned
//		error Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelAtom *numberStrKernelAtom) convertKernelToBigFloat(
	numStrKernel *NumberStrKernel,
	roundingType NumberRoundingType,
	roundToFactionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	*big.Float,
	int,
	error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var t big.Float
	var err error
	var numOfFractionalDigits int

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"convertKernelToBigFloat()",
		"")

	if err != nil {

		return &t, numOfFractionalDigits, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return &t, numOfFractionalDigits, err
	}

	if !roundingType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrRoundingSpec Rounding Type' is invalid!\n"+
			"'roundingType' string  value = '%v'\n"+
			"'roundingType' integer value = '%v'\n",
			ePrefix.String(),
			roundingType.String(),
			roundingType.XValueInt())

		return &t, numOfFractionalDigits, err

	}

	if roundToFactionalDigits < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'roundToFactionalDigits' is invalid!\n"+
			"'roundToFactionalDigits' has an integer value less than zero (0).\n"+
			"'roundToFactionalDigits' integer value = '%v'\n",
			ePrefix.String(),
			roundToFactionalDigits)

		return &t, numOfFractionalDigits, err

	}

	nStrKernelElectron := numberStrKernelElectron{}

	_,
		err = nStrKernelElectron.getSetIsNonZeroValue(
		numStrKernel,
		ePrefix.XCpy("numStrKernel"))

	if err != nil {

		return &t, numOfFractionalDigits, err

	}

	err = nStrKernelElectron.rationalizeFractionalIntegerDigits(
		numStrKernel,
		ePrefix)

	if err != nil {

		return &t, numOfFractionalDigits, err

	}

	var ok bool

	var newNumStrKernel NumberStrKernel

	newNumStrKernel,
		err = new(numberStrKernelNanobot).copyOut(
		numStrKernel,
		ePrefix.XCpy(
			"numStrKernel->newNumStrKernel"))

	if err != nil {

		return &t, numOfFractionalDigits, err

	}

	err = nStrKernelElectron.setUninitializedKernelToZero(
		&newNumStrKernel,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {

		return &t, numOfFractionalDigits, err

	}

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err = new(NumStrRoundingSpec).NewRoundingSpec(
		roundingType,
		roundToFactionalDigits,
		ePrefix.XCpy(
			"->numStrRoundingSpec"))

	if err != nil {

		return &t, numOfFractionalDigits, err

	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&newNumStrKernel,
		numStrRoundingSpec,
		ePrefix.XCpy(
			"->newNumStrKernel"))

	if err != nil {

		return &t, numOfFractionalDigits, err

	}

	numOfFractionalDigits =
		newNumStrKernel.fractionalDigits.GetRuneArrayLength()

	var numValueStr string

	if newNumStrKernel.numberSign == NumSignVal.Negative() {
		numValueStr += "-"
	}

	numValueStr += newNumStrKernel.integerDigits.GetCharacterString()

	if newNumStrKernel.fractionalDigits.GetRuneArrayLength() > 0 {

		numValueStr += "."
		numValueStr += newNumStrKernel.fractionalDigits.GetCharacterString()

	}

	_,
		ok = t.SetString(numValueStr)

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error Converting floating point number string to *big.Float!\n"+
			"The following string of numeric digits from 'numStrKernel'\n"+
			"generated an error.\n"+
			"Numeric Digits = '%v'\n",
			ePrefix.String(),
			numValueStr)

		return &t, numOfFractionalDigits, err
	}

	t.SetPrec(t.MinPrec())
	t.SetMode(big.AwayFromZero)

	if t.Acc() != big.Exact {

		err = fmt.Errorf("%v\n"+
			"Error Converting floating point number string to *big.Float!\n"+
			"Could NOT generate Accuracy = 'Exact' for the following string of\n"+
			"numeric digits from 'numStrKernel'.\n"+
			"Numeric Digits = '%v'\n",
			ePrefix.String(),
			numValueStr)

	}

	return &t, numOfFractionalDigits, err
}

//	convertKernelToBigInt
//
//	Converts an instance of NumberStrKernel to an integer
//	value of type *big.Int.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		converted to an integer of type *big.Int.
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
//
//		'roundingType' is only applied in cases where the
//		current NumberStrKernel instance consists of a
//		floating point numeric value.
//
//		If in doubt as to a suitable rounding method,
//		'HalfAwayFromZero' is recommended.
//
//		Possible values are listed as follows:
//			NumRoundType.None()	- Invalid Value
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//		Definitions:
//
//			NoRounding
//
//				Signals that no rounding operation will be
//				performed on fractional digits. The
//				fractional digits will therefore remain
//				unchanged.
//
//			HalfUpWithNegNums
//
//				Half Round Up Including Negative Numbers.
//				This method is intuitive but may produce
//				unexpected results when applied to negative
//				numbers.
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
//			Half Round Down Including Negative Numbers. This
//			method is also considered intuitive but may
//			produce unexpected results when applied to
//			negative numbers.
//
//			'HalfDownWithNegNums' rounds .5 down.
//
//				Examples of HalfDownWithNegNums
//
//				7.6 rounds up to 8
//				7.5 rounds down to 7
//				7.4 rounds down to 7
//				-7.4 rounds up to -7
//				-7.5 rounds down to -8
//				-7.6 rounds down to -8
//
//			HalfAwayFromZero
//
//				The 'HalfAwayFromZero' method rounds .5 further
//				away from zero.	It provides clear and consistent
//				behavior when dealing with negative numbers.
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
//				Round Half Towards Zero. 'HalfTowardsZero' rounds
//				0.5	closer to zero. It provides clear and
//				consistent behavior	when dealing with negative
//				numbers.
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
//				Round Half To Even Numbers. 'HalfToEven' is
//				also called	Banker's Rounding. This method
//				rounds 0.5 to the nearest even digit.
//
//					Examples of HalfToEven
//
//					7.5 rounds up to 8 (because 8 is an even
//					number)	but 6.5 rounds down to 6 (because
//					6 is an even number)
//
//					HalfToEven only applies to 0.5. Other
//					numbers (not ending	in 0.5) round to
//					nearest as usual, so:
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
//				Round Half to Odd Numbers. Similar to 'HalfToEven',
//				but in this case 'HalfToOdd' rounds 0.5 towards odd
//				numbers.
//
//					Examples of HalfToOdd
//
//					HalfToOdd only applies to 0.5. Other numbers
//					(not ending	in 0.5) round to nearest as usual.
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
//				Round Half Randomly. Uses a Random Number Generator
//				to choose between rounding 0.5 up or down.
//
//				All numbers other than 0.5 round to the nearest as
//				usual.
//
//			Floor
//
//				Yields the nearest integer down. Floor does not apply
//				any	special treatment to 0.5.
//
//				Floor Function: The greatest integer that is less than
//				or equal to x
//
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				In mathematics and computer science, the floor function
//				is the function that takes as input a real number x,
//				and gives as output the greatest integer less than or
//				equal to x,	denoted floor(x) or ⌊x⌋.
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//				Examples of Floor
//
//					Number     Floor
//					 2           2
//					 2.4         2
//					 2.9         2
//					-2.5        -3
//					-2.7        -3
//					-2          -2
//
//			Ceiling
//
//				Yields the nearest integer up. Ceiling does not
//				apply any special treatment to 0.5.
//
//				Ceiling Function: The least integer that is
//				greater than or	equal to x.
//				Source:
//					https://www.mathsisfun.com/sets/function-floor-ceiling.html
//
//				The ceiling function maps x to the least integer
//				greater than or equal to x, denoted ceil(x) or
//				⌈x⌉.[1]
//
//				Source:
//					https://en.wikipedia.org/wiki/Floor_and_ceiling_functions
//
//					Examples of Ceiling
//
//						Number    Ceiling
//						 2           2
//						 2.4         3
//						 2.9         3
//						-2.5        -2
//						-2.7        -2
//						-2          -2
//
//			Truncate
//
//				Apply NO Rounding whatsoever. The Round From Digit
//				is dropped or deleted. The Round To Digit is NEVER
//				changed.
//
//				Examples of Truncate
//
//					Example-1
//					Number: 23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:   23.14 - The Round From Digit
//					is dropped.
//
//					Example-2
//					Number: -23.14567
//					Objective: Round to two decimal places to
//					the right of the decimal point.
//					Rounding Method: Truncate
//					Round To Digit:   4
//					Round From Digit: 5
//					Rounded Number:  -23.14 - The Round From Digit
//					is dropped.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	*big.Int
//
//		If this method completes successfully, the
//		numeric	value represented by the current instance
//		of	NumberStrKernel will be returned as a type
//		*big.Int.
//
//	error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If errors
//		are	encountered during processing, the returned
//		error Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelAtom *numberStrKernelAtom) convertKernelToBigInt(
	numStrKernel *NumberStrKernel,
	roundingType NumberRoundingType,
	errPrefDto *ePref.ErrPrefixDto) (
	*big.Int,
	error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	bigIntValue := big.NewInt(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelElectron."+
			"convertKernelToBigInt()",
		"")

	if err != nil {

		return bigIntValue, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return bigIntValue, err
	}

	if !roundingType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrRoundingSpec Rounding Type' is invalid!\n"+
			"'roundingType' string  value = '%v'\n"+
			"'roundingType' integer value = '%v'\n",
			ePrefix.String(),
			roundingType.String(),
			roundingType.XValueInt())

		return bigIntValue, err

	}

	nStrKernelElectron := numberStrKernelElectron{}

	_,
		err = nStrKernelElectron.getSetIsNonZeroValue(
		numStrKernel,
		ePrefix.XCpy("numStrKernel"))

	if err != nil {

		return bigIntValue, err

	}

	err = nStrKernelElectron.rationalizeFractionalIntegerDigits(
		numStrKernel,
		ePrefix)

	if err != nil {

		return bigIntValue, err

	}

	var ok bool

	var copyNStrKernel NumberStrKernel

	copyNStrKernel,
		err = new(numberStrKernelNanobot).copyOut(
		numStrKernel,
		ePrefix.XCpy(
			"copyNStrKernel<-numStrKernel"))

	if err != nil {

		return bigIntValue, err

	}

	err = nStrKernelElectron.setUninitializedKernelToZero(
		&copyNStrKernel,
		ePrefix.XCpy(
			"copyNStrKernel"))

	if err != nil {

		return bigIntValue, err

	}

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err =
		new(NumStrRoundingSpec).NewRoundingSpec(
			roundingType,
			0,
			ePrefix)

	if err != nil {
		return bigIntValue, err
	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&copyNStrKernel,
		numStrRoundingSpec,
		ePrefix)

	if err != nil {
		return bigIntValue, err
	}

	var numberString string

	if copyNStrKernel.numberSign == NumSignVal.Negative() {

		numberString += "-"

	}

	numberString +=
		copyNStrKernel.integerDigits.GetCharacterString()

	_,
		ok = bigIntValue.SetString(
		numberString,
		10)

	if !ok {
		err = fmt.Errorf("%v\n"+
			"Error Converting Rounded Integer string to *big.Int!\n"+
			"The following integerDigits string generated an error.\n"+
			"numberString = '%v'\n",
			ePrefix.String(),
			numberString)
	}

	return bigIntValue, err
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

//	formatNumStrComponents
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
//	decSeparator				DecimalSeparatorSpec
//
//		This structure contains the radix point or decimal
//		separator character(s) (a.k.a. decimal point)
//		which be used to separate integer and fractional
//		digits within a formatted Number String.
//
//	intSeparatorDto				IntegerSeparatorSpec
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
//		Example: 1,000,000,000
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
//		The IntegerSeparatorSpec type provides the flexibility
//		necessary to process these complex number separation
//		formats.
//
//	roundingSpec 				NumStrRoundingSpec
//
//		Numeric Value Rounding Specification. This
//		specification contains all the parameters
//		required to configure and apply a rounding
//		algorithm for floating point Number Strings.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with negative numeric values.
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with positive numeric values.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with zero numeric values.
//
//	numberFieldSpec			NumStrNumberFieldSpec
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	numStr						string
//
//		If this method completes successfully, the
//		numeric	value represented by the NumberStrKernel
//		instance, 'numStrKernel', will be returned as a
//		formatted Number String, 'numStr'.
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If errors
//		are	encountered during processing, the returned
//		error Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelAtom *numberStrKernelAtom) formatNumStrComponents(
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
		"negNumSearchSpecAtom."+
			"testValidityOfNegNumSearchSpec()",
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

	if numStrKernel.GetNumberOfIntegerDigits() == 0 &&
		numStrKernel.GetNumberOfFractionalDigits() == 0 {
		numStr = "0"

		return numStr, err
	}

	var newNumStrKernel NumberStrKernel

	newNumStrKernel,
		err = new(numberStrKernelNanobot).copyOut(
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
			"Input parameter 'nStrFormatSpec.DecSeparator'\n"+
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
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
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
			"testValidityOfNegNumSearchSpec()",
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
			"Error: This instance of NumberStrKernel is invalid!"+
			"Both Integer Digits and Fractional Digits are empty"+
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
			"Error: This instance of NumberStrKernel is invalid!"+
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
				"Error: This instance of NumberStrKernel is invalid!"+
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
				"Error: This instance of NumberStrKernel is invalid!"+
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
			"Error: This instance of NumberStrKernel is invalid!"+
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
				"Error: This instance of NumberStrKernel is invalid!"+
				"The Number Sign Value is invalid.\n"+
				"NumberStrKernel has a zero numeric value.\n"+
				"However, internal flag numStrKernel.getSetIsNonZeroValue\n"+
				"is set to 'true'.\n",
				ePrefix.String())

			return isValid, err

		}

		if numValHasNonZeroVal == true &&
			numStrKernel.isNonZeroValue == false {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!"+
				"The Number Sign Value is invalid.\n"+
				"NumberStrKernel has a non-zero numeric value.\n"+
				"However, internal flag numStrKernel.getSetIsNonZeroValue\n"+
				"is set to 'false'.\n",
				ePrefix.String())

			return isValid, err

		}
	}

	isValid = true

	return isValid, err
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelAtom.
func (numStrKernelAtom numberStrKernelAtom) ptr() *numberStrKernelAtom {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	return &numberStrKernelAtom{
		lock: new(sync.Mutex),
	}
}

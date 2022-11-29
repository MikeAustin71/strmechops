package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type mathFloatHelperQuark struct {
	lock *sync.Mutex
}

//	raiseToFloatPositiveExponent
//
//	Receives a pointer to a big.Float floating point
//	number and raises that number to the power specified
//	by input parameter 'exponent'.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
//
//	This method will only process positive exponents.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	base						*big.Float
//
//		This floating point value will be raised to the
//		power of 'exponent' and returned to the calling
//		function.
//
//	exponent					int64
//
//		This value will be used to raise 'base' to the
//		power of 'exponent'.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
//
//		If this value is less than zero, an error will be
//		returned.
//
//	precisionBits				uint
//
//		The number of bits in the mantissa of the result
//		'raisedToExponent'. Effectively, this parameter
//		controls the precision and accuracy for the
//		calculation of 'base' raised to the power of
//		'exponent'.
//
//		If in doubt as to this number, identify the
//		total number of calculation result integer and
//		fractional digits required to store an accurate
//		result and multiply this number times four (+4).
//		Be sure to add a safety buffer of extra numerical
//		digits (maybe 50-digits) to handle processing
//		requirements.
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
//	*big.Float
//
//		If this method completes successfully, this will
//		return 'base' value raised to the power of the
//		'exponent' value.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
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
func (floatHelperQuark *mathFloatHelperQuark) raiseToFloatPositiveExponent(
	base *big.Float,
	exponent int64,
	precisionBits uint,
	errPrefDto *ePref.ErrPrefixDto) (
	*big.Float,
	error) {

	if floatHelperQuark.lock == nil {
		floatHelperQuark.lock = new(sync.Mutex)
	}

	floatHelperQuark.lock.Lock()

	defer floatHelperQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperQuark."+
			"raiseToFloatPositiveExponent()",
		"")

	if err != nil {
		return big.NewFloat(0), err
	}

	if exponent < 0 {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'exponent' is invalid!\n"+
			"'exponent' is less than zero and negative.\n"+
			"exponent = '%v'\n",
			ePrefix.String(),
			exponent)

		return big.NewFloat(0), err
	}

	if base == nil {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'base' is invalid!\n"+
			"'base' is a nil pointer.\n",
			ePrefix.String())

		return big.NewFloat(0), err
	}

	if exponent == 0 {

		return big.NewFloat(0), err
	}

	baseStr := base.Text('f', -1)

	fmt.Printf("baseStr = %v\n",
		baseStr)
	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.

	newBase,
		ok := big.NewFloat(0).
		SetMode(big.AwayFromZero).
		SetPrec(precisionBits).
		SetString(baseStr)

	if !ok {

		err = fmt.Errorf("\n%v\n"+
			"Error: newBase.SetString(baseStr) Failed!\n"+
			"baseStr = %v\n",
			ePrefix.String(),
			base.Text('f', 80))

		return big.NewFloat(0), err
	}

	raisedToExponent :=
		big.NewFloat(0).
			SetPrec(precisionBits).
			SetMode(big.AwayFromZero).
			SetInt64(1)

	for i := int64(0); i < exponent; i++ {

		raisedToExponent.Mul(raisedToExponent, newBase)
	}

	raisedToExponent.SetPrec(raisedToExponent.MinPrec())

	if raisedToExponent.Acc() != big.Exact {

		err = fmt.Errorf("\n%v\n"+
			"Error: Final 'raisedToExponent' Accuracy is NOT equal to 'Exact'!\n"+
			"Accuracy may be compromised.\n"+
			"'raisedToExponent' Accuracy = %v\n"+
			"raisedToExponent = %v\n",
			ePrefix.String(),
			raisedToExponent.Acc(),
			raisedToExponent.Text('f', -1))

	}

	return raisedToExponent, err
}

//	raiseToIntPositiveExponent
//
//	Receives a pointer to a big.Float floating point
//	number and raises that number to the power specified
//	by input parameter 'exponent'.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
//
//	The floating point precision value required to
//	support the calculation result ('raisedToExponent')
//	is computed internally.
//
//	This method will only process positive exponents.
//
//	This method employs integer mathematics and type
//	big.Int to compute the base floating point value
//	raised to the power of exponent. As such it
//	produces highly accurate results.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	base						*big.Float
//
//		This floating point value will be raised to the
//		power of 'exponent' and returned to the calling
//		function.
//
//	exponent					int64
//
//		This value will be used to raise 'base' to the
//		power of 'exponent'.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
//
//		If this value is less than zero, an error will be
//		returned.
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
//	*big.Float
//
//		If this method completes successfully, this
//		parameter will return 'base' value raised to the
//		power of the 'exponent' value.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
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
func (floatHelperQuark *mathFloatHelperQuark) raiseToIntPositiveExponent(
	base *big.Float,
	exponent int64,
	errPrefDto *ePref.ErrPrefixDto) (
	*big.Float,
	error) {

	if floatHelperQuark.lock == nil {
		floatHelperQuark.lock = new(sync.Mutex)
	}

	floatHelperQuark.lock.Lock()

	defer floatHelperQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperQuark."+
			"raiseToIntPositiveExponent()",
		"")

	if err != nil {
		return big.NewFloat(0), err
	}

	if base == nil {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'base' is invalid!\n"+
			"'base' is a nil pointer.\n",
			ePrefix.String())

		return big.NewFloat(0), err
	}

	if exponent == 0 {

		return big.NewFloat(0), err
	}

	if exponent < 0 {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'exponent' is INVALID!\n"+
			"'exponent' has a value less than zero.\n"+
			"exponent = %v\n",
			ePrefix.String(),
			exponent)

		return big.NewFloat(0), err
	}

	// base.SetPrec(base.MinPrec())

	var pureNumStrStats PureNumberStrComponents

	pureNumStrStats,
		err = new(numStrMathAtom).
		pureNumStrToComponents(
			base.Text('f', -1),
			".",
			true,
			ePrefix.XCpy(
				"<-base"))

	if err != nil {

		return big.NewFloat(0), err
	}

	if pureNumStrStats.NumStrStats.IsZeroValue == true {

		// base is zero.
		//	zero^exponent = zero
		return big.NewFloat(0), err
	}

	var ok bool

	var numStr string

	if pureNumStrStats.NumStrStats.NumberSign ==
		NumSignVal.Negative() {

		numStr += "-"

	}

	numStr += pureNumStrStats.AllIntegerDigitsNumStr

	bigIntBase,
		ok := big.NewInt(0).
		SetString(
			numStr,
			10)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: bigIntBase=SetString(numStr)\n"+
			"SetString Failed!\n"+
			"numStr = %v\n",
			ePrefix,
			numStr)

		return big.NewFloat(0), err
	}

	bigIntExponent := big.NewInt(0).SetInt64(exponent)

	bigIntBase.Exp(bigIntBase, bigIntExponent, nil)

	numStr =
		bigIntBase.Text(10)

	lenNumStr := uint64(len(numStr))

	negativeAdjustment := uint64(0)

	if pureNumStrStats.NumStrStats.NumberSign ==
		NumSignVal.Negative() {

		negativeAdjustment = 1
	}

	raisedToPowerStats := PureNumberStrComponents{}

	raisedToPowerStats.NumStrStats.NumOfFractionalDigits =
		pureNumStrStats.NumStrStats.NumOfFractionalDigits *
			uint64(exponent)

	raisedToPowerStats.NumStrStats.NumOfIntegerDigits =
		lenNumStr -
			raisedToPowerStats.NumStrStats.NumOfFractionalDigits -
			negativeAdjustment

	numStr =
		numStr[0:(raisedToPowerStats.NumStrStats.NumOfIntegerDigits+
			negativeAdjustment)] +
			"." +
			numStr[raisedToPowerStats.NumStrStats.NumOfIntegerDigits+
				negativeAdjustment:]

	var precisionBits uint

	precisionBits,
		err = new(mathFloatHelperAtom).precisionBitsFromRequiredDigits(
		int64(raisedToPowerStats.NumStrStats.NumOfIntegerDigits),
		int64(raisedToPowerStats.NumStrStats.NumOfFractionalDigits),
		5,
		ePrefix)

	if err != nil {

		return big.NewFloat(0), err
	}

	var raisedToPowerFloat *big.Float

	raisedToPowerFloat,
		ok = big.NewFloat(0).
		SetPrec(precisionBits).
		SetMode(big.AwayFromZero).
		SetString(numStr)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error: raisedToPowerFloat=SetString(numStr)\n"+
			"SetString Failed!\n"+
			"numStr = %v\n",
			ePrefix.String(),
			numStr)

	}

	raisedToPowerFloat.SetPrec(raisedToPowerFloat.MinPrec())

	fmt.Printf("raisedToPower = %v\n"+
		"raisedToPower Precision = %v\n"+
		"    Estimated Precision = %v\n",
		raisedToPowerFloat.Text('f', -1),
		raisedToPowerFloat.Prec(),
		precisionBits)

	return raisedToPowerFloat, err
}

//	roundBigFloat
//
//	This method will round a big.Float numeric value
//	based on rounding specifications passed by input
//	parameters, 'roundingType' and
//	'roundToFractionalDigits'.
//
//	The big.Float numeric value to be rounded is passed
//	by input parameter, 'numberToRound'. The final
//	rounded value is returned through input parameter,
//	'roundedNumber'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberToRound				*big.Float
//
//		A pointer to an instance of big.Float. This is
//		the numeric value which will be rounded.
//
//
//	roundedNumber				*big.Float
//
//		A pointer to an instance of big.Float. The result
//		of the rounding operation performed on input
//		parameter,'numberToRound', will be stored here.
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		current instance of NumberStrKernel.
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
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the radix
//		point or decimal separator (a.k.a. decimal point). This
//		controls the number of fractional digits remaining after
//		completion of the number rounding operation.
//
//		For the purposes of this method, any
//		'roundToFractionalDigits' value greater than 20,000 will
//		trigger an error return.
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
func (floatHelperQuark *mathFloatHelperQuark) roundBigFloat(
	numberToRound *big.Float,
	roundedNumber *big.Float,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if floatHelperQuark.lock == nil {
		floatHelperQuark.lock = new(sync.Mutex)
	}

	floatHelperQuark.lock.Lock()

	defer floatHelperQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperQuark."+
			"roundBigFloat()",
		"")

	if err != nil {
		return err
	}

	if numberToRound == nil {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'numberToRound' is invalid!\n"+
			"'numberToRound' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if roundedNumber == nil {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'numberToRound' is invalid!\n"+
			"'numberToRound' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	precision := numberToRound.Prec()

	numStrKernel := NumberStrKernel{}

	numStrKernel.numberValueType = NumValType.FloatingPoint()

	var numberStats NumberStrStatsDto

	numberStats,
		err = new(mathFloatHelperMechanics).
		floatNumToIntFracRunes(
			numberToRound,
			&numStrKernel.integerDigits,
			&numStrKernel.fractionalDigits,
			ePrefix.XCpy("numStrKernel<-numberToRound"))

	if err != nil {
		return err
	}

	numStrKernel.isNonZeroValue = !numberStats.IsZeroValue

	numStrKernel.numberSign = numberStats.NumberSign

	numStrKernel.numberValueType = numberStats.NumberValueType

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err =
		new(NumStrRoundingSpec).NewRoundingSpec(
			roundingType,
			roundToFractionalDigits,
			ePrefix)

	if err != nil {
		return err
	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&numStrKernel,
		numStrRoundingSpec,
		ePrefix)

	if err != nil {
		return err
	}

	roundedNumber.SetPrec(precision)

	pureNumberStr := numStrKernel.GetPureNumberStr(
		".",
		true)

	var ok bool
	_,
		ok = roundedNumber.SetString(
		pureNumberStr)

	if !ok {

		err = fmt.Errorf("\n%v\n"+
			"Error: roundedNumber.SetString(pureNumberStr) FAILED!\n"+
			"big.Float was unable to set the number string value\n"+
			"for input parameter 'roundedNumber'.\n"+
			"pureNumberStr = %v",
			ePrefix.String(),
			pureNumberStr)

	}

	return err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"sync"
)

// numberStrKernelMolecule - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelMolecule struct {
	lock *sync.Mutex
}

//	convertKernelToFloatNum
//
//	Converts an instance of NumberStrKernel to a floating
//	point numeric value.
//
//	The type of floating point number returned by this
//	conversion operation is controlled by the 'numericValue'
//	parameter which MUST be set to one of the following types:
//
//			*float32
//			*float64
//			*big.Float
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		converted to an integer of type int.
//
//	numericValue 				interface{}
//
//		This empty interface MUST be convertible to one
//		of the following types:
//
//			*float32
//			*float64
//			*big.Float
//
//		This parameter will receive the converted numeric
//		value of the 'numStrKernel' instance cast to one
//		of the supported types listed above.
//
//		if 'float32' return type is specified, be advised
//		that 'float32' capacity is approximately 7 to 8
//		digits including integer and fractional digits.
//
//		if 'float64' capacity is approximately 15 to 17
//		digits including integer and fractional digits.
//
//		Source:
//		https://en.wikipedia.org/wiki/Floating-point_arithmetic#IEEE_754:_floating_point_in_modern_computers
//
//		Type	Sign  Exponent	Significand	  Total   Exponent	 Bits	     Number of
//								   field       bits     bias    precision	decimal digits
//		         ----  --------	-----------   -----   --------  ---------	--------------
//		Single	 1		  8			23			32	     127	   24		    ~7.2
//		Double	 1		 11			52			64	    1023	   53		   ~15.9
//
//		Type *big.Float provides the most accurate representation of
//		floating point numeric values with large numbers of integer
//		and fractional digits.
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
//		radix point or decimal separator (a.k.a. decimal point).
//		This value is equal to the number fractional digits which
//		will remain in the floating point number after completion
//		of the number rounding operation.
//
//		If parameter 'roundingType' is set to
//		NumRoundType.NoRounding(), 'roundToFractionalDigits' is
//		ignored and has no effect.
//
//		if 'roundToFractionalDigits' is set to a value greater
//		than the number of fractional digits in 'numStrKernel',
//		the number of fractional digits will be extended with
//		zero values and reflected in the numeric value returned
//		through parameter 'numericValue'.
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
//		Type ErrPrefixDto is included in the 'errpref'
//		software package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numOfFractionalDigits		int
//
//		The number of fractional digits contained in the
//		returned numeric value passed through parameter
//		'numericValue'.
//
//		For returned types *float32 and *float64, those
//		types may not have the capacity to store all the
//		fractional digits generated by the Number String
//		Kernel ('numStrKernel').
//
//		For returned type *big.Float, the
//		'numOfFractionalDigits' parameter can be used to
//		improve accuracy in conversions from *big.Float
//		to character strings.
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
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernelMolecule *numberStrKernelMolecule) convertKernelToFloatNum(
	numStrKernel *NumberStrKernel,
	numericValue interface{},
	roundingType NumberRoundingType,
	roundToFactionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	numOfFractionalDigits int,
	err error) {

	if numStrKernelMolecule.lock == nil {
		numStrKernelMolecule.lock = new(sync.Mutex)
	}

	numStrKernelMolecule.lock.Lock()

	defer numStrKernelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelMolecule."+
			"convertKernelToFloatNum()",
		"")

	if err != nil {

		return numOfFractionalDigits, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return numOfFractionalDigits, err
	}

	var bigFloatNum *big.Float

	bigFloatNum,
		numOfFractionalDigits,
		err = new(numberStrKernelAtom).convertKernelToBigFloat(
		numStrKernel,
		roundingType,
		roundToFactionalDigits,
		ePrefix.XCpy(
			"->bigFloatNum"))

	if err != nil {

		return numOfFractionalDigits, err

	}

	var ok bool
	var maxFloatValue *big.Float

	switch numericValue.(type) {

	case *float32:

		maxFloatValue = big.NewFloat(math.MaxFloat32)

		if bigFloatNum.Cmp(maxFloatValue) > 0 {

			err = fmt.Errorf("%v\n"+
				"ERROR: Numeric Value Out Of Range for type 'float32'!\n"+
				"The numeric value of the NumStrKernel (numStrKernel)\n"+
				"exceeds the maximum capacity of type 'float32'.\n"+
				"NumStrKernel Numeric Value = %v\n"+
				"Max Float32 Capactiy Value = %v\n",
				ePrefix.String(),
				bigFloatNum.Text('f', -1),
				maxFloatValue.Text('f', -1))

			return numOfFractionalDigits, err

		}

		var float32Value *float32

		float32Value, ok = numericValue.(*float32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *float32 cast to 'float32Value' failed!\n",
				ePrefix.String())

			return numOfFractionalDigits, err

		}

		var float64Value float64

		float64Value,
			_ = bigFloatNum.Float64()

		*float32Value = float32(float64Value)

	case *float64:

		maxFloatValue = big.NewFloat(math.MaxFloat64)

		if bigFloatNum.Cmp(maxFloatValue) > 0 {

			err = fmt.Errorf("%v\n"+
				"ERROR: Numeric Value Out Of Range for type 'float64'!\n"+
				"The numeric value of the NumStrKernel (numStrKernel)\n"+
				"exceeds the maximum capacity of type 'float64'.\n"+
				"NumStrKernel Numeric Value = %v\n"+
				"Max Float64 Capacity Value = %v\n",
				ePrefix.String(),
				bigFloatNum.Text('f', -1),
				maxFloatValue.Text('f', -1))

			return numOfFractionalDigits, err

		}

		var float64Value *float64

		float64Value, ok = numericValue.(*float64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *float64 cast to 'float64Value' failed!\n",
				ePrefix.String())

			return numOfFractionalDigits, err

		}

		*float64Value,
			_ = bigFloatNum.Float64()

	case *big.Float:

		var bigFloatVal *big.Float

		bigFloatVal, ok = numericValue.(*big.Float)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Float cast to 'bigFloatVal' failed!\n",
				ePrefix.String())

			return numOfFractionalDigits, err

		}

		bigFloatVal.Set(bigFloatNum)

		bigFloatVal.SetPrec(bigFloatVal.MinPrec())
		bigFloatVal.SetMode(big.AwayFromZero)

		accuracy := bigFloatVal.Acc()

		if accuracy != big.Exact {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Float Accuracy is NOT equal to 'Exact'!\n"+
				"Accuracy = '%v'\n",
				ePrefix.String(),
				accuracy.String())

			return numOfFractionalDigits, err

		}

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numericValue' is an invalid type!\n"+
			"'numericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", numericValue))

	}

	return numOfFractionalDigits, err
}

//	convertKernelToIntNum
//
//	Converts an instance of NumberStrKernel to an integer
//	numeric value.
//
//	The type of integer returned by this conversion operation
//	is controlled by the 'numericValue' parameter which MUST
//	be set to one of the following types:
//
//			*int
//			*int32
//			*int64
//			*big.Int
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		converted to an integer of type int.
//
//	numericValue 				interface{}
//
//		This empty interface MUST be convertible to one
//		of the following types:
//
//			*int
//			*int32
//			*int64
//			*big.Int
//
//		This parameter will receive the converted numeric
//		value of the 'numStrKernel' instance cast to one
//		of the supported types listed above.
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
//		This object encapsulates an error prefix string which
//		is included in all returned error messages. Usually,
//		it contains the name of the calling method or methods
//		listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
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
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernelMolecule *numberStrKernelMolecule) convertKernelToIntNum(
	numStrKernel *NumberStrKernel,
	numericValue interface{},
	roundingType NumberRoundingType,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelMolecule.lock == nil {
		numStrKernelMolecule.lock = new(sync.Mutex)
	}

	numStrKernelMolecule.lock.Lock()

	defer numStrKernelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelMolecule."+
			"convertKernelToIntNum()",
		"")

	if err != nil {

		return err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var bigIntNum *big.Int

	bigIntNum,
		err = new(numberStrKernelAtom).convertKernelToBigInt(
		numStrKernel,
		roundingType,
		ePrefix)

	if err != nil {

		return err

	}

	var ok bool
	var maxIntValue *big.Int

	switch numericValue.(type) {

	case *int:

		maxIntValue = big.NewInt(int64(math.MaxInt))

		if bigIntNum.Cmp(maxIntValue) == 1 {

			err = fmt.Errorf("%v\n"+
				"ERROR: Numeric Value Out Of Range for type 'int'!\n"+
				"The numeric value of the NumStrKernel (numStrKernel)\n"+
				"exceeds the maximum capacity of type 'int'.\n"+
				"Numeric Value = %v\n",
				ePrefix.String(),
				bigIntNum.Text(10))

			return err

		}

		var intValue *int

		intValue, ok = numericValue.(*int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int cast to 'intValue' failed!\n",
				ePrefix.String())

			return err

		}

		*intValue = int(bigIntNum.Int64())

	case *int32:

		maxIntValue = big.NewInt(int64(math.MaxInt32))

		if bigIntNum.Cmp(maxIntValue) == 1 {

			err = fmt.Errorf("%v\n"+
				"ERROR: Numeric Value Out Of Range for type 'int32'!\n"+
				"The numeric value of the NumStrKernel (numStrKernel)\n"+
				"exceeds the maximum capacity of type 'int32'.\n"+
				"Numeric Value = %v\n",
				ePrefix.String(),
				bigIntNum.Text(10))

			return err

		}

		var int32Value *int32

		int32Value, ok = numericValue.(*int32)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int32 cast to 'int32Value' failed!\n",
				ePrefix.String())

			return err

		}

		*int32Value = int32(bigIntNum.Int64())

	case *int64:

		maxIntValue = big.NewInt(math.MaxInt64)

		if bigIntNum.Cmp(maxIntValue) == 1 {

			err = fmt.Errorf("%v\n"+
				"ERROR: Numeric Value Out Of Range for type 'int64'!\n"+
				"The numeric value of the NumStrKernel (numStrKernel)\n"+
				"exceeds the maximum capacity of type 'int64'.\n"+
				"NumStrKernel Numeric Value = %v\n",
				ePrefix.String(),
				bigIntNum.Text(10))

			return err

		}

		var int64Value *int64

		int64Value, ok = numericValue.(*int64)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *int32 cast to 'int32Value' failed!\n",
				ePrefix.String())

			return err

		}

		*int64Value = bigIntNum.Int64()

	case *big.Int:

		var bigIntValue *big.Int

		bigIntValue, ok = numericValue.(*big.Int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Int cast to 'bigIntValue' failed!\n",
				ePrefix.String())

			return err

		}

		bigIntValue.Set(bigIntNum)

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numericValue' is an invalid type!\n"+
			"'numericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", numericValue))

	}

	return err
}

//	convertSignedIntToKernel
//
//	Receives an empty interface which is assumed to be an
//	integer numeric value configured as one of the following
//	types:
//
//		int8
//		int16
//		int32
//		int	(equivalent to int32)
//		int64
//
//	This integer numeric value is then converted to a
//	type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	signedIntNumericValue		interface{}
//
//		This empty interface is assumed to encapsulate a signed
//		integer	numeric value comprised of one of the following
//		types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//
//		This numeric value will be used to populate the instance
//		of NumberStrKernel passed by parameter, 'numStrKernel'.
//
//		If the object passed by this empty interface is NOT one
//		of the types listed above, an error will be returned.
//
//	errPrefDto          *ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
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
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrKernelMolecule *numberStrKernelMolecule) convertSignedIntToKernel(
	numStrKernel *NumberStrKernel,
	signedIntNumericValue interface{},
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelMolecule.lock == nil {
		numStrKernelMolecule.lock = new(sync.Mutex)
	}

	numStrKernelMolecule.lock.Lock()

	defer numStrKernelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelMolecule."+
			"convertSignedIntToKernel()",
		"")

	if err != nil {

		return err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	switch signedIntNumericValue.(type) {

	case int8, int16, int, int32, int64:

		goto intNumericValueProcessing

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'signedIntNumericValue' is an invalid type!\n"+
			"Only signed integer types are suppored.\n"+
			"'signedIntNumericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", signedIntNumericValue))

		return err

	}

intNumericValueProcessing:

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		signedIntNumericValue)

	runeArrayDto := RuneArrayDto{
		CharsArray:   []rune(numberStr),
		Description1: "",
		Description2: "",
	}

	err = runeArrayDto.SetCharacterSearchType(
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy("runeArrayDto"))

	if err != nil {

		return err

	}

	*numStrKernel,
		err = new(numStrBuilderElectron).parsePurNumStr(
		runeArrayDto,
		ePrefix.XCpy(
			numberStr))

	if err != nil {
		return err
	}

	return err
}

//	convertBigIntToKernel
//
//	Receives a parameter of type empty interface which is
//	assumed to be a type *big.Int. If the empty interface
//	is NOT convertible to a type *big.Int an error will be
//	returned.
//
//	This *big.Int integer numeric value is then converted to
//	a type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	bigFloatValue				*big.Float
//
//		The numeric value this method will use to configure
//		parameter 'numStrKernel'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this parameter
//		to 'nil'.
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
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrKernelMolecule *numberStrKernelMolecule) convertBigFloatToKernel(
	numStrKernel *NumberStrKernel,
	bigFloatValue *big.Float,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelMolecule.lock == nil {
		numStrKernelMolecule.lock = new(sync.Mutex)
	}

	numStrKernelMolecule.lock.Lock()

	defer numStrKernelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelMolecule."+
			"convertBigFloatToKernel()",
		"")

	if err != nil {

		return err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(numberStrKernelElectron).empty(
		numStrKernel)

	precision := bigFloatValue.Prec()

	if precision > uint(math.MaxInt) {

		err = fmt.Errorf("%v\n"+
			"Error: Precision Out-Of-Range!\n"+
			"The precision specified by parameter 'bigFloatValue'\n"+
			"exceeds the maximum value for an integer and therefore\n"+
			"cannot be converted to a string value.\n"+
			"'bigFloatValue' precision = '%v'\n"+
			"Maximum allowed conversion precision = '%v'\n",
			ePrefix.String(),
			precision,
			math.MaxInt)

		return err
	}

	numberStr := fmt.Sprintf("%v",
		bigFloatValue.Text('f', int(precision)))

	var searchResults CharSearchNumStrParseResultsDto
	decimalSeparatorSpec := DecimalSeparatorSpec{}
	numParsingTerminators := RuneArrayCollection{}
	negativeNumSearchSpecs := NegNumSearchSpecCollection{}

	err = decimalSeparatorSpec.SetDecimalSeparatorStr(
		".",
		ePrefix.XCpy("Decimal Point '.'"))

	if err != nil {
		return err
	}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading Minus Sign '-'"))

	if err != nil {
		return err
	}

	runeArrayDto := RuneArrayDto{
		CharsArray:   []rune(numberStr),
		Description1: "",
		Description2: "",
	}

	searchResults,
		*numStrKernel,
		err = new(numStrBuilderElectron).extractNumRunes(
		runeArrayDto,
		"integerNumberStr",
		0,
		-1,
		negativeNumSearchSpecs,
		decimalSeparatorSpec,
		numParsingTerminators,
		false,
		ePrefix.XCpy(
			numberStr))

	if err != nil {
		return err
	}

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	return err
}

//	convertBigIntToKernel
//
//	Receives a parameter of type empty interface which is
//	assumed to be a type *big.Int. If the empty interface
//	is NOT convertible to a type *big.Int an error will be
//	returned.
//
//	This *big.Int integer numeric value is then converted to
//	a type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	bigIntValue					*big.Int
//
//		The numeric value this method will use to configure
//		parameter 'numStrKernel'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this parameter
//		to 'nil'.
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
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrKernelMolecule *numberStrKernelMolecule) convertBigIntToKernel(
	numStrKernel *NumberStrKernel,
	bigIntValue *big.Int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelMolecule.lock == nil {
		numStrKernelMolecule.lock = new(sync.Mutex)
	}

	numStrKernelMolecule.lock.Lock()

	defer numStrKernelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelMolecule."+
			"convertBigIntToKernel()",
		"")

	if err != nil {

		return err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if bigIntValue == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'bigIntValue' is a nil pointer!\n",
			ePrefix.String())

		return err

	}

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		bigIntValue.Text(10))

	var searchResults CharSearchNumStrParseResultsDto
	decimalSeparatorSpec := DecimalSeparatorSpec{}
	numParsingTerminators := RuneArrayCollection{}
	negativeNumSearchSpecs := NegNumSearchSpecCollection{}

	err = decimalSeparatorSpec.SetDecimalSeparatorStr(
		".",
		ePrefix.XCpy("Decimal Point '.'"))

	if err != nil {
		return err
	}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading Minus Sign '-'"))

	if err != nil {
		return err
	}

	runeArrayDto := RuneArrayDto{
		CharsArray:   []rune(numberStr),
		Description1: "",
		Description2: "",
	}

	searchResults,
		*numStrKernel,
		err = new(numStrBuilderElectron).extractNumRunes(
		runeArrayDto,
		"integerNumberStr",
		0,
		-1,
		negativeNumSearchSpecs,
		decimalSeparatorSpec,
		numParsingTerminators,
		false,
		ePrefix.XCpy(
			numberStr))

	if err != nil {
		return err
	}

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	return err
}

//	convertFloatToKernel
//
//	Receives an empty interface which is assumed to be a
//	floating point numeric value configured as one of the
//	following types:
//
//		float32
//		float64
//
//	This floating point numeric value is then converted to
//	a type of 'NumberStrKernel' and returned to the calling
//	function via input parameter 'numStrKernel'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	floatNumericValue			interface{}
//
//		This empty interface is assumed to encapsulate a
//		floating point numeric value comprised of one of the
//		following types:
//
//		float32
//		float64
//
//		This numeric value will be used to populate the instance
//		of NumberStrKernel passed by parameter, 'numStrKernel'.
//
//		If the object passed by this empty interface is NOT one
//		of the types listed above, an error will be returned.
//
//	errPrefDto          *ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
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
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrKernelMolecule *numberStrKernelMolecule) convertFloatToKernel(
	numStrKernel *NumberStrKernel,
	floatNumericValue interface{},
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelMolecule.lock == nil {
		numStrKernelMolecule.lock = new(sync.Mutex)
	}

	numStrKernelMolecule.lock.Lock()

	defer numStrKernelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelMolecule."+
			"convertFloatToKernel()",
		"")

	if err != nil {

		return err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	switch floatNumericValue.(type) {

	case float32, float64:

		goto floatNumericValueProcessing

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'floatNumericValue' is an invalid type!\n"+
			"Only floating point types are suppored (float32, float64).\n"+
			"'floatNumericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", floatNumericValue))

		return err

	}

floatNumericValueProcessing:

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		floatNumericValue)

	var searchResults CharSearchNumStrParseResultsDto
	decimalSeparatorSpec := DecimalSeparatorSpec{}
	numParsingTerminators := RuneArrayCollection{}
	negativeNumSearchSpecs := NegNumSearchSpecCollection{}

	err = decimalSeparatorSpec.SetDecimalSeparatorStr(
		".",
		ePrefix.XCpy("Decimal Point '.'"))

	if err != nil {
		return err
	}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading Minus Sign '-'"))

	if err != nil {
		return err
	}

	runeArrayDto := RuneArrayDto{
		CharsArray:   []rune(numberStr),
		Description1: "",
		Description2: "",
	}

	searchResults,
		*numStrKernel,
		err = new(numStrBuilderElectron).extractNumRunes(
		runeArrayDto,
		"integerNumberStr",
		0,
		-1,
		negativeNumSearchSpecs,
		decimalSeparatorSpec,
		numParsingTerminators,
		false,
		ePrefix.XCpy(
			numberStr))

	if err != nil {
		return err
	}

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	return err
}

//	convertUnsignedInteger
//
//	Receives an empty interface which is assumed to be an
//	unsigned integer numeric value configured as one of the
//	following types:
//
//		uint8
//		uint16
//		uint32
//		uint	(equivalent to uint32)
//		uint64
//
//	This unsigned integer numeric value is then converted to
//	a type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	intNumericValue	interface{}
//
//		This empty interface is assumed to encapsulate an unsigned
//		integer numeric value comprised of one of the following
//		types:
//
//			uint8
//			uint16
//			uint32
//			uint (equivalent to uint32)
//			uint64
//
//		If the object passed by this empty interface is NOT one of
//		the types listed above, an error will be returned.
//
//	numberSign					NumericSignValueType
//
//		The Number Sign is specified by means of a
//		NumericSignValueType enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = -2 - Infer From Number
//			NumSignVal.Negative() = -1 - Valid Value
//			NumSignVal.Zero()     =  0 - Valid Value
//			NumSignVal.Positive() =  1 - Valid Value
//
//		Unsigned integer values are by default converted as
//		positive numeric values. If this parameter is set
//		to NumSignVal.Negative(), the numeric value returned
//		through parameter 'numStrKernel' will be classified
//		as a negative value.
//
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored and
//		the final number sign for the converted numeric
//		value will be set to 'positive'.
//
//	errPrefDto          *ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this parameter
//		to 'nil'.
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
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrKernelMolecule *numberStrKernelMolecule) convertUnsignedInteger(
	numStrKernel *NumberStrKernel,
	unsignedIntValue interface{},
	numberSign NumericSignValueType,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelMolecule.lock == nil {
		numStrKernelMolecule.lock = new(sync.Mutex)
	}

	numStrKernelMolecule.lock.Lock()

	defer numStrKernelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelMolecule."+
			"convertUnsignedInteger()",
		"")

	if err != nil {

		return err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	switch unsignedIntValue.(type) {

	case uint8, uint16, uint, uint32, uint64:

		goto unsignedIntProcessing

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'unsignedIntValue' is an invalid type!\n"+
			"Only unsigned integer types are suppored.\n"+
			"'unsignedIntValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", unsignedIntValue))

		return err

	}

unsignedIntProcessing:

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		unsignedIntValue)

	var searchResults CharSearchNumStrParseResultsDto
	decimalSeparatorSpec := DecimalSeparatorSpec{}
	numParsingTerminators := RuneArrayCollection{}
	negativeNumSearchSpecs := NegNumSearchSpecCollection{}

	err = decimalSeparatorSpec.SetDecimalSeparatorStr(
		".",
		ePrefix.XCpy("Decimal Point '.'"))

	if err != nil {
		return err
	}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading Minus Sign '-'"))

	if err != nil {
		return err
	}

	runeArrayDto := RuneArrayDto{
		CharsArray:   []rune(numberStr),
		Description1: "",
		Description2: "",
	}

	searchResults,
		*numStrKernel,
		err = new(numStrBuilderElectron).extractNumRunes(
		runeArrayDto,
		"unsignedIntegerNumberStr",
		0,
		-1,
		negativeNumSearchSpecs,
		decimalSeparatorSpec,
		numParsingTerminators,
		false,
		ePrefix.XCpy(
			numberStr))

	if err != nil {
		return err
	}

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	if numStrKernel.numberSign == NumSignVal.Zero() {

		return err
	}

	if numberSign == NumSignVal.Negative() {
		numStrKernel.numberSign = NumSignVal.Negative()
	}

	return err
}

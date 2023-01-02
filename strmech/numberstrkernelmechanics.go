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

//	compareNumStrKernels
//
//	Receives pointers to two instances of NumberStrKernel,
//	'numStrKernel01' and 'numStrKernel02'.
//
//	The numeric value of 'numStrKernel01' is compared to
//	that of 'numStrKernel01'. The comparison results are
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
func (numStrKernelMech *numberStrKernelMechanics) compareNumStrKernels(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	comparisonValue int,
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
			"compareNumStrKernels()",
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

	var numStrStats01, numStrStats02 NumberStrStatsDto

	nStrKernelAtom := numberStrKernelAtom{}

	numStrStats01,
		err = nStrKernelAtom.calcNumStrKernelStats(
		numStrKernel01,
		ePrefix.XCpy(
			"numStrStats01<-numStrKernel01"))

	if err != nil {

		return comparisonValue, err
	}

	numStrStats02,
		err = nStrKernelAtom.calcNumStrKernelStats(
		numStrKernel02,
		ePrefix.XCpy(
			"numStrStats02<-numStrKernel02"))

	if err != nil {

		return comparisonValue, err
	}

	if !numStrStats01.NumberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: numStrStats01.NumberSign is Invalid!\n"+
			"numStrStats01.NumberSign integer value ='%v'\n"+
			"numStrStats01.NumberSign string value  ='%v'\n",
			ePrefix.String(),
			numStrStats01.NumberSign.XValueInt(),
			numStrStats01.NumberSign.String())

		return comparisonValue, err

	}

	if !numStrStats02.NumberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: numStrStats02.NumberSign is Invalid!\n"+
			"numStrStats02.NumberSign integer value ='%v'\n"+
			"numStrStats02.NumberSign string value  ='%v'\n",
			ePrefix.String(),
			numStrStats02.NumberSign.XValueInt(),
			numStrStats02.NumberSign.String())

		return comparisonValue, err
	}

	// Are both values zero?
	if numStrStats01.IsZeroValue &&
		numStrStats02.IsZeroValue {

		comparisonValue = 0

		return comparisonValue, err

	}

	if numStrStats01.NumberSign ==
		numStrStats02.NumberSign {
		// MUST BE Equal Number Signs

		if numStrStats01.NumberSign == NumSignVal.Zero() {

			comparisonValue = 0

			return comparisonValue, err

		} else {
			// MUST BE -
			//	Number Sign is NumSignVal.Positive() OR
			//	NumSignVal.Negative() AND the Number Signs
			//	are equal.

			if numStrStats01.NumOfSignificantIntegerDigits >
				numStrStats02.NumOfSignificantIntegerDigits {

				if numStrStats01.NumberSign == NumSignVal.Positive() {

					comparisonValue = 1

				} else {
					// MUST BE
					// numStrStats01.NumberSign == NumSignVal.Negative()

					comparisonValue = -1
				}

				return comparisonValue, err

			} else if numStrStats01.NumOfSignificantIntegerDigits <
				numStrStats02.NumOfSignificantIntegerDigits {

				if numStrStats01.NumberSign == NumSignVal.Positive() {

					comparisonValue = -1

				} else {
					// MUST BE
					// numStrStats01.NumberSign == NumSignVal.Negative()

					comparisonValue = 1
				}

				return comparisonValue, err

			} else {
				// MUST BE EQUAL Number Signs And EQUAL
				//		Significant Integer Digits
				// numStrStats01.NumOfSignificantIntegerDigits ==
				// 		numStrStats02.NumOfSignificantIntegerDigits

				if numStrStats01.NumOfIntegerDigits ==
					numStrStats02.NumOfIntegerDigits &&
					numStrStats01.NumOfFractionalDigits ==
						numStrStats02.NumOfFractionalDigits {
					// Int and Frac arrays have equal lengths

					comparisonValue,
						err = new(numberStrKernelQuark).
						compareNumStrKernelValues(
							numStrKernel01,
							numStrKernel02,
							ePrefix.XCpy(
								"numStrKernel01 vs "+
									"numStrKernel02"))

					return comparisonValue, err
				}

				// Int and Frac arrays have unequal lengths

				comparisonValue,
					err = new(numberStrKernelAtom).
					prepareCompareNumStrKernels(
						numStrKernel01,
						numStrKernel02,
						ePrefix.XCpy(
							"numStrKernel01 vs "+
								"numStrKernel02"))

				return comparisonValue, err
			}

		}

	} // End Of
	// if numStrStats01.NumberSign == numStrStats02.NumberSign

	// Number Signs Are NOT Equal
	//	Each sign is either Positive, Negative or Zero.
	if numStrStats01.NumberSign == NumSignVal.Positive() {

		comparisonValue = 1

	} else if numStrStats01.NumberSign == NumSignVal.Negative() {

		comparisonValue = -1

	} else if numStrStats01.NumberSign == NumSignVal.Zero() {

		if numStrStats02.NumberSign == NumSignVal.Positive() {

			comparisonValue = -1

		} else {
			// MUST BE
			//	numStrStats02.NumberSign == NumSignVal.Negative()

			comparisonValue = 1

		}
	}

	return comparisonValue, err
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
//	significandRoundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm which will be applied
//		when rounding fractional digits contained in the
//		significand of the returned Scientific Notation
//		value.
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
//	significandRoundToFactionalDigits	int
//
//		This parameter in conjunction with
//		'significandRoundingType' to specify the number of digits
//		to the right of the radix point, or decimal point, which
//		will be returned as the Scientific Notation significand
//		value.
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		radix point or decimal separator (a.k.a. decimal point).
//		This value is equal to the number fractional digits which
//		will remain in the floating point number after completion
//		of the number rounding operation.
//
//		If parameter 'roundingType' is set to NumRoundType.NoRounding(),
//		'significandRoundToFactionalDigits' is ignored and has no
//		effect.
//
//		if 'significandRoundToFactionalDigits' is set to a value
//		greater than the number of fractional digits in the
//		'significand', the number of fractional digits will be
//		extended with zero values and reflected in the numeric
//		value returned through parameter 'numericValue'.
//
//		NOTE: Rounding the significand to zero is considered bad
//		form. Common practice always retains at least one digit
//		to the right of the decimal point in Scientific Notation.
//
//
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
	significandRoundingType NumberRoundingType,
	significandRoundToFactionalDigits int,
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

		err = nStrKernelMolecule.
			convertNumericValueToKernel(
				&sciNotKernel.exponent,
				exponent,
				ePrefix.XCpy(
					fmt.Sprintf("sciNotKernel.exponent='%v'",
						exponent)))

		if err != nil {

			return sciNotKernel, err
		}

	} else if intArrayLen > 1 {

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

		err = nStrKernelMolecule.
			convertNumericValueToKernel(
				&sciNotKernel.exponent,
				exponent,
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

		err = nStrKernelMolecule.
			convertNumericValueToKernel(
				&sciNotKernel.exponent,
				exponent,
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

			err = nStrKernelMolecule.
				convertNumericValueToKernel(
					&sciNotKernel.exponent,
					exponent,
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

			err = nStrKernelMolecule.
				convertNumericValueToKernel(
					&sciNotKernel.exponent,
					exponent,
					ePrefix.XCpy(
						fmt.Sprintf("sciNotKernel.exponent='%v'",
							exponent)))

			if err != nil {

				return sciNotKernel, err
			}

		}

	} // END OF Compute Significand and Exponent

	if significandRoundingType == NumRoundType.NoRounding() {

		// Nothing more to do.
		return sciNotKernel, err
	}

	// Round Significand if necessary
	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err = new(NumStrRoundingSpec).NewRoundingSpec(
		significandRoundingType,
		significandRoundToFactionalDigits,
		ePrefix.XCpy(
			"numStrRoundingSpec<-"))

	if err != nil {

		return sciNotKernel, err

	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&sciNotKernel.significand,
		numStrRoundingSpec,
		ePrefix.XCpy(
			"sciNotKernel.significand"))

	return sciNotKernel, err
}

// setNumStrKernelFromRoundedDirtyNumStr
//
// Receives a Dirty Number String, extracts a valid
// Native Number String and proceeds to reconfigure the
// NumberStrKernel instance passed as input parameter
// 'numStrKernel' with the resulting numeric value.
//
// A "Dirty Number String" is a malformed number string
// containing numeric digits which will be converted to a
// properly formatted Native Number String and used to
// configure the NumberStrKernel instance 'numStrKernel'.
//
//	Examples Of Dirty Number Strings
//
//		$1,254.65
//		1 000 000,00 €
//		1.000.000,00 €
//		6,78,90,00,00,00,00,000
//		6,7890,0000,0000,0000
//
// All the examples shown above are valid number string
// formats used by different countries and cultures.
// There is nothing wrong with these formats. The term
// "Dirty" simply distinguishes these formats from the
// Native Number String format required by many
// functions and packages in the Go Programming
// Language. The Native Number String format is one
// which is used as a standard format for numeric
// conversions performed in the Go Programming Language
// as well as many other programming languages.
//
// The 'Dirty Number String' passed as input parameter
// 'dirtyNumberStr' is expected to comply with the
// following requirements:
//
//  1. The dirty number string must contain numeric
//     digit characters zero to nine inclusive (0-9).
//
//  2. The dirty number string must contain a radix
//     point or decimal separator to separate
//     integer and fractional digits in a floating
//     point numeric value. This decimal separator
//     is specified by input parameter,
//     'decimalSeparator'.
//
//     If no decimal separator is identified in the
//     dirty number string, the numeric value is
//     assumed to be an integer value.
//
//  3. The dirty number string must designate
//     negative numeric values using one of the
//     following three negative number symbols:
//
//     (a)	A Leading Minus Sign ('-').
//     Example: -123.45
//
//     (b)	A Trailing Minus Sign ('-').
//     Example: 123.45-
//
//     (c) A combination of leading and trailing
//     Parentheses ('()').
//     Example: (123.45)
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the NumberStrKernel instance passed as
//	input parameter 'numStrKernel'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		NumberStrKernel instance will be reconfigured
//		with the numeric value extracted from input
//		parameter 'dirtyNumStr'.
//
//	dirtyNumberStr				string
//
//		This number string contains the numeric digits
//		which will be extracted, converted to a valid
//		numeric value and used to configure the
//		NumberStrKernel instance passed as input
//		parameter 'numStrKernel'.
//
//		The 'dirtyNumberStr' is expected to comply with
//		the following requirements:
//
//		1.	The dirty number string must contain numeric
//			digit characters zero to nine inclusive (0-9).
//
//		2.	The dirty number string must contain a radix
//			point or decimal separator to separate
//			integer and fractional digits in a floating
//			point numeric value. This decimal separator
//			is specified by input parameter,
//			'decimalSeparator'.
//
//			If no decimal separator is identified in the
//			dirty number string, the numeric value is
//			assumed to be an integer value.
//
//		3.	The dirty number string must designate
//			negative numeric values using one of the
//			following three negative number symbols:
//
//			(a)	A Leading Minus Sign ('-').
//				Example: -123.45
//
//			(b)	A Trailing Minus Sign ('-').
//				Example: 123.45-
//
//			(c) A combination of leading and trailing
//				Parentheses ('()').
//				Example: (123.45)
//
//		If 'dirtyNumberStr' does not contain any numeric
//		digits, an error will be returned.
//
//	decimalSeparator			string
//
//		Type DecimalSeparatorSpec is used to specify the
//		radix point or decimal separator which will
//		separate integer and fractional digits in the
//		dirty number string passed as input parameter
//		'dirtyNumberStr'.
//
//		The decimal separator will typically consist of
//		one or more non-numeric characters.
//
//		If 'decimalSeparator' consists of an empty
//		or zero length sting, it is assumed that the
//		numeric value contained in input parameter
//		'dirtyNumberStr' is an integer value.
//
//		In the US, Australia, UK, most of Canada and many
//		other countries the period ('.'), or decimal
//		point, separates integer and fractional digits
//		within a floating point numeric value.
//
//		Other countries, including many in the European
//		Union, use the comma (',') to separate integer
//		and fractional digits within a number string.
//
//		If 'decimalSeparator' contains any one of the
//		following invalid characters, an error will be
//		returned.
//
//			Invalid Decimal Separator Characters
//							'-'
//							'('
//							')'
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		reconfigured instance of NumberStrKernel
//		(numStrKernel).
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
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). This controls the number
//		of fractional digits remaining after completion
//		of the number rounding operation.
//
//		If input parameter 'roundingType' is set to
//		NumRoundType.NoRounding(),
//		'roundToFractionalDigits' is ignored and no
//		rounding operation is performed.
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
//		If this method completes successfully, an
//		instance of NumberStrStatsDto will be returned
//		containing a profile and key statistics on the
//		numeric value encapsulated in the
//		NumberStrKernel instance passed as input
//		parameter, 'nStrKernel'.
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
//			If 'true', the subject numeric value is equal
//			to zero ('0').
//
//			If 'false', the subject numeric value is
//			greater than or less than zero ('0').
//		}
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
func (numStrKernelMech *numberStrKernelMechanics) setNumStrKernelFromRoundedDirtyNumStr(
	numStrKernel *NumberStrKernel,
	dirtyNumberStr string,
	decimalSeparator string,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	numStrStatsDto NumberStrStatsDto,
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
			"setNumStrKernelFromRoundedDirtyNumStr()",
		"")

	if err != nil {

		return numStrStatsDto, err
	}

	var nativeNumStr string

	nativeNumStr,
		err = new(NumStrHelper).DirtyToNativeNumStr(
		dirtyNumberStr,
		decimalSeparator,
		ePrefix.XCpy(
			"nativeNumStr<-"+
				"dirtyNumberStr"))

	if err != nil {

		return numStrStatsDto, err

	}

	err = new(numberStrKernelQuark).
		setNumStrKernelFromNativeNumStr(
			numStrKernel,
			nativeNumStr,
			ePrefix.XCpy(
				"numStrKernel<-"+
					"nativeNumStr"))

	if err != nil {

		return numStrStatsDto, err

	}

	if roundingType != NumRoundType.NoRounding() {

		var numStrRoundingSpec NumStrRoundingSpec

		numStrRoundingSpec,
			err =
			new(NumStrRoundingSpec).NewRoundingSpec(
				roundingType,
				roundToFractionalDigits,
				ePrefix)

		if err != nil {

			return numStrStatsDto, err
		}

		err = new(numStrMathRoundingNanobot).roundNumStrKernel(
			numStrKernel,
			numStrRoundingSpec,
			ePrefix)

		if err != nil {

			return numStrStatsDto, err
		}

	}

	numStrStatsDto,
		err = new(numberStrKernelAtom).
		calcNumStrKernelStats(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	return numStrStatsDto, err
}

// setNumStrKernelFromRoundedNativeNumStr
//
// Receives a Native Number String, extracts the numeric
// value contained therein and proceeds to reconfigure
// the NumberStrKernel instance passed as input parameter
// 'numStrKernel' with the that calculated numeric value.
//
// The term 'Native Number String' means that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math'
// and 'big' (big.Int, big.Float, big.Rat) routinely
// parse and convert this type of number string to
// numeric values. In addition, Native Number Strings are
// frequently consumed by external library functions such
// as this one (String Mechanics 'strmech') to convert
// strings to numeric values and numeric values to
// strings.
//
// While this format is inconsistent with many national
// and cultural formatting conventions, number strings
// which fail to implement this standardized formatting
// protocol will generate errors in some Golang library
// functions.
//
//	Examples Of Native Number Strings
//		1000000
//		12.5483
//		-1000000
//		-12.5483
//
// A valid Native Number String must conform to the
// standardized formatting criteria defined below:
//
//  1. A Native Number String Consists of numeric
//     character digits zero through nine inclusive
//     (0-9).
//
//  2. A Native Number String will include a period
//     or decimal point ('.') to separate integer and
//     fractional digits within a number string.
//
//     Native Number String Floating Point Value:
//     123.1234
//
//  3. A Native Number String will always format
//     negative numeric values with a leading minus sign
//     ('-').
//
//     Native Number String Negative Value:
//     -123.2
//
//  4. A Native Number String WILL NEVER include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Native Number String: 1000000
//
//  5. Native Number Strings will only consist of:
//
//     (a)	Numeric digits zero through nine inclusive (0-9).
//
//     (b)	A decimal point ('.') for floating point
//     numbers.
//
//     (c)	A leading minus sign ('-') in the case of
//     negative numeric values.
//
//  6. A Native Number String will NEVER include
//     currency symbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the NumberStrKernel instance passed as
//	input parameter 'numStrKernel'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		NumberStrKernel instance will be reconfigured
//		with the numeric value extracted from input
//		parameter 'nativeNumStr'.
//
//	nativeNumStr				string
//
//		A Native Number String containing the numeric
//		character digits which will be converted to, and
//		stored in, the NumberStrKernel object passed as
//		input parameter 'numStrKernel'.
//
//		The term 'Native Number String' means that the
//		number string format is designed to interoperate
//		with the Golang programming language library
//		functions and packages. Types like 'strconv',
//		'strings', 'math' and 'big' (big.Int, big.Float,
//		big.Rat) routinely parse and convert this type of
//		number string to generate numeric values. In
//		addition, Native Number Strings are frequently
//		consumed by external library functions such	as
//		this one (String Mechanics 'strmech') to convert
//		strings to numeric values and numeric values to
//		strings.
//
//		If 'nativeNumStr' fails to meet the formatting
//		criteria for a Native Number String defined
//		below, an error will be returned.
//
//		A valid Native Number String must conform to the
//		standardized formatting criteria defined below:
//
//	 	1.	A Native Number String Consists of numeric
//	 	  	character digits zero through nine inclusive
//	 	  	(0-9).
//
//	 	2.	A Native Number String will include a period
//	 	  	or decimal point ('.') to separate integer and
//	 	  	fractional digits within a number string.
//
//	 	  	Native Number String Floating Point Value:
//	 	   				123.1234
//
//	 	3.	A Native Number String will always format
//	 	  	negative numeric values with a leading minus sign
//	 	  	('-').
//
//	 	  	Native Number String Negative Value:
//	 	  					-123.2
//
//	 	4.	A Native Number String WILL NEVER include integer
//			separators such as commas (',') to separate
//			integer digits by thousands.
//
//	 	   					NOT THIS: 1,000,000
//	 	   		Native Number String: 1000000
//
//	 	5.	Native Number Strings will only consist of:
//
//			(a)	Numeric digits zero through nine inclusive (0-9).
//
//			(b)	A decimal point ('.') for floating point
//				numbers.
//
//			(c)	A leading minus sign ('-') in the case of
//				negative numeric values.
//
//		6.	A Native Number String will NEVER include
//			currency symbols.
//
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the
//		type of rounding algorithm that will be applied for
//		the	rounding of fractional digits contained in the
//		reconfigured instance of NumberStrKernel
//		(numStrKernel).
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
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). This controls the number
//		of fractional digits remaining after completion
//		of the number rounding operation.
//
//		If input parameter 'roundingType' is set to
//		NumRoundType.NoRounding(),
//		'roundToFractionalDigits' is ignored and no
//		rounding operation is performed.
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
//		If this method completes successfully, an
//		instance of NumberStrStatsDto will be returned
//		containing a profile and key statistics on the
//		numeric value encapsulated in the
//		NumberStrKernel instance passed as input
//		parameter, 'nStrKernel'.
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
//			If 'true', the subject numeric value is equal
//			to zero ('0').
//
//			If 'false', the subject numeric value is
//			greater than or less than zero ('0').
//		}
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
func (numStrKernelMech *numberStrKernelMechanics) setNumStrKernelFromRoundedNativeNumStr(
	numStrKernel *NumberStrKernel,
	nativeNumStr string,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	numStrStatsDto NumberStrStatsDto,
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
			"setNumStrKernelFromRoundedNativeNumStr()",
		"")

	if err != nil {

		return numStrStatsDto, err
	}

	nativeNumStr,
		_,
		err = new(NumStrHelper).NormalizeNativeNumStr(
		nativeNumStr,
		ePrefix.XCpy(
			"nativeNumStr<-nativeNumStr"))

	if err != nil {

		return numStrStatsDto, err

	}

	err = new(numberStrKernelQuark).
		setNumStrKernelFromNativeNumStr(
			numStrKernel,
			nativeNumStr,
			ePrefix.XCpy(
				"numStrKernel<-"+
					"nativeNumStr"))

	if err != nil {

		return numStrStatsDto, err

	}

	if roundingType != NumRoundType.NoRounding() {

		var numStrRoundingSpec NumStrRoundingSpec

		numStrRoundingSpec,
			err =
			new(NumStrRoundingSpec).NewRoundingSpec(
				roundingType,
				roundToFractionalDigits,
				ePrefix)

		if err != nil {

			return numStrStatsDto, err
		}

		err = new(numStrMathRoundingNanobot).roundNumStrKernel(
			numStrKernel,
			numStrRoundingSpec,
			ePrefix)

		if err != nil {

			return numStrStatsDto, err
		}

	}

	numStrStatsDto,
		err = new(numberStrKernelAtom).
		calcNumStrKernelStats(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	return numStrStatsDto, err
}

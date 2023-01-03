package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelQuark - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelQuark struct {
	lock *sync.Mutex
}

//	compareNumStrKernelValues
//
//	This method receives pointers to two instances of
//	NumberStrKernel, 'numStrKernel01' and
//	'numStrKernel02'.
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
// # IMPORTANT
//
//  1. This method assumes the integer and fractional
//     digit arrays contained in input parameters
//     'numStrKernel01' and 'numStrKernel02' are equal
//     in length.
//
//     If the integer digit array length and fractional
//     digit array lengths of 'numStrKernel01' are NOT
//     equal to the corresponding array lengths in
//     'numStrKernel02', an error will be returned.
//
//  2. This method assumes that the number signs for
//     'numStrKernel01' and 'numStrKernel02' are equal.
//     If 'numStrKernel01' and 'numStrKernel02' ARE NOT
//     equal, an error will be returned.
//
//     Possible values for number sign are listed as
//     follows:
//
//     NumSignVal.Negative() = -1
//     NumSignVal.Zero()     =  0
//     NumSignVal.Positive() =  1
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
func (numStrKernelQuark *numberStrKernelQuark) compareNumStrKernelValues(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	comparisonValue int,
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"equalizeNStrFracDigitsLengths()",
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

	if numStrKernel01.numberSign !=
		numStrKernel02.numberSign {

		err = fmt.Errorf("%v\n"+
			"ERROR: The number signs for 'numStrKernel01'"+
			"and 'numStrKernel02' ARE NOT EQUAL!\n"+
			"numStrKernel01.numberSign = '%v'\n"+
			"numStrKernel02.numberSign = '%v'\n",
			ePrefix.String(),
			numStrKernel01.numberSign.String(),
			numStrKernel02.numberSign.String())

		return comparisonValue, err
	}

	if numStrKernel01.numberSign == NumSignVal.Zero() {

		comparisonValue = 0

		return comparisonValue, err
	}

	lenIntDigits01 :=
		len(numStrKernel01.integerDigits.CharsArray)

	lenFracDigits01 :=
		len(numStrKernel01.fractionalDigits.CharsArray)

	lenIntDigits02 :=
		len(numStrKernel02.integerDigits.CharsArray)

	lenFracDigits02 :=
		len(numStrKernel02.fractionalDigits.CharsArray)

	if lenIntDigits01 != lenIntDigits02 {

		err = fmt.Errorf("%v\n"+
			"Error: numStrKernel01.integerDigits array length is\n"+
			"NOT EQUAL to numStrKernel02.integerDigits array length!\n,"+
			"numStrKernel01.integerDigits array length = '%v'\n"+
			"numStrKernel02.integerDigits array length = '%v'\n",
			ePrefix.String(),
			lenIntDigits01,
			lenIntDigits02)

		return comparisonValue, err
	}

	if lenFracDigits01 != lenFracDigits02 {

		err = fmt.Errorf("%v\n"+
			"Error: numStrKernel01.fractionalDigits array length is\n"+
			"NOT EQUAL to numStrKernel02.fractionalDigits array length!\n,"+
			"numStrKernel01.fractionalDigits array length = '%v'\n"+
			"numStrKernel02.fractionalDigits array length = '%v'\n",
			ePrefix.String(),
			lenFracDigits01,
			lenFracDigits02)

		return comparisonValue, err
	}

	for i := 0; i < lenIntDigits01; i++ {

		if numStrKernel01.integerDigits.CharsArray[i] >
			numStrKernel02.integerDigits.CharsArray[i] {

			// Remember, numStrKernel01 & numStrKernel02
			// number signs are equal
			if numStrKernel01.numberSign == NumSignVal.Positive() {

				comparisonValue = 1

			} else {
				//	MUST BE -
				//	numStrKernel01.numberSign ==
				//		NumSignVal.Negative()
				comparisonValue = -1

			}

			return comparisonValue, err

		}

		if numStrKernel02.integerDigits.CharsArray[i] >
			numStrKernel01.integerDigits.CharsArray[i] {

			// Remember, numStrKernel01 & numStrKernel02
			// number signs are equal
			if numStrKernel02.numberSign == NumSignVal.Positive() {

				comparisonValue = -1

			} else {
				//	MUST BE -
				//	numStrKernel02.numberSign ==
				//		NumSignVal.Negative()
				comparisonValue = 1

			}

			return comparisonValue, err
		}
	}

	// Integer Digits ARE EQUAL
	// Now test Fractional Digits

	for k := 0; k < lenFracDigits01; k++ {

		if numStrKernel01.fractionalDigits.CharsArray[k] >
			numStrKernel02.fractionalDigits.CharsArray[k] {

			// Remember, numStrKernel01 & numStrKernel02
			// number signs are equal
			if numStrKernel01.numberSign == NumSignVal.Positive() {

				comparisonValue = 1

			} else {
				//	MUST BE -
				//	numStrKernel01.numberSign ==
				//		NumSignVal.Negative()
				comparisonValue = -1

			}

			return comparisonValue, err
		}

		if numStrKernel02.fractionalDigits.CharsArray[k] >
			numStrKernel01.fractionalDigits.CharsArray[k] {

			// Remember, numStrKernel01 & numStrKernel02
			// number signs are equal
			if numStrKernel02.numberSign == NumSignVal.Positive() {

				comparisonValue = -1

			} else {
				//	MUST BE -
				//	numStrKernel02.numberSign ==
				//		NumSignVal.Negative()
				comparisonValue = 1

			}

			return comparisonValue, err
		}
	}

	// MUST BE -
	//	numStrKernel01 and numStrKernel02 have
	//	equal numerical values
	comparisonValue = 0

	return comparisonValue, err
}

// getNativeNumStr
//
// Returns a Native Number String representing the
// numeric value of a NumberStrKernel instance passed as
// an input parameter.
//
// The term 'Native' means that the number string format
// is designed to interoperate with the Golang
// programming language library functions and packages.
// Types like 'strconv', 'strings', 'math' and 'big'
// (big.Int, big.Float, big.Rat) routinely parse and
// convert this type of number string to numeric values.
// In addition, Native Number Strings are frequently
// consumed by external library functions such as this one
// (String Mechanics 'strmech') to convert strings to
// numeric values and numeric values to strings.
//
// While this format is inconsistent with many national
// and cultural formatting conventions, number strings
// which fail to implement this standardized formatting
// protocol will generate errors in some Golang library
// functions.
//
// The numeric value represented by the returned Native
// Number String will be rounded as specified by input
// parameters, 'roundingType' and
// 'roundToFractionalDigits'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel.
//		This instance contains the numeric value which
//		will be used to generate and return a Native
//		Number String.
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
//	nativeNumStr				string
//
//		If this method completes successfully, a Native
//		Number String representing the numeric value
//		encapsulated in the NumberStrKernel instance
//		passed as input parameter 'numStrKernel' will be
//		returned.
//
//		The 'Native' Number String returned by this
//		method implements a standardized format defined
//		as follows:
//
//		1.	A Native Number String Consists of numeric
//		  	character digits zero through nine inclusive
//		  	(0-9).
//
//		2.	A Native Number String will include a period
//		  	or decimal point ('.') to separate integer and
//		  	fractional digits within a number string.
//
//				Native Number String Floating Point Value:
//								123.1234
//
//		3.	A Native Number String will always format
//				negative numeric values with a leading minus sign
//				('-').
//
//				Native Number String Negative Value:
//							-123.2
//
//		4.	A Native Number String WILL NEVER include integer
//		  	separators such as commas (',') to separate
//		  	integer digits by thousands.
//
//		    				NOT THIS: 1,000,000
//				Native Number String: 1000000
//
//		5.	Native Number Strings will only consist of:
//
//			(a)	Numeric digits (0-9).
//
//			(b)	A decimal point ('.') for floating point
//				numbers.
//
//			(c)	A leading minus sign ('-') in the case of
//				negative numeric values.
//
//
//		The numeric value represented by the returned
//		Native Number String will be rounded as specified
//		by input parameters, 'roundingType' and
//		'roundToFractionalDigits'.
//
//	nativeNumStrStats			NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the integer and fractional digits contained
//		in the return parameter 'nativeNumStr'.
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
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelQuark *numberStrKernelQuark) getNativeNumStr(
	numStrKernel *NumberStrKernel,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	nativeNumStr string,
	nativeNumStrStats NumberStrStatsDto,
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"getSetNumValueType()",
		"")

	if err != nil {

		return nativeNumStr,
			nativeNumStrStats,
			err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return nativeNumStr,
			nativeNumStrStats,
			err
	}

	var err2 error
	_,
		err2 = new(numberStrKernelAtom).
		testValidityOfNumStrKernel(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	if err2 != nil {

		err = fmt.Errorf("Error: 'numStrKernel' is invalid!\n"+
			"This instance of NumberStrKernel failed validity tests.\n"+
			"Validation Error: \n%v\n",
			err2.Error())

		return nativeNumStr,
			nativeNumStrStats,
			err
	}

	if roundingType == NumRoundType.NoRounding() {

		nativeNumStr,
			nativeNumStrStats,
			err = new(numberStrKernelBoson).
			createNativeNumStrFromNumStrKernel(
				numStrKernel,
				ePrefix.XCpy(
					"numStrKernel"))

		return nativeNumStr,
			nativeNumStrStats,
			err
	}

	var deepCopyNumStrKernel NumberStrKernel

	err = new(numberStrKernelNanobot).
		copy(
			&deepCopyNumStrKernel,
			numStrKernel,
			ePrefix.XCpy(
				"deepCopyNumStrKernel<-"+
					"numStrKernel"))

	if err != nil {

		return nativeNumStr,
			nativeNumStrStats,
			err
	}

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err =
		new(NumStrRoundingSpec).NewRoundingSpec(
			roundingType,
			roundToFractionalDigits,
			ePrefix)

	if err != nil {

		return nativeNumStr,
			nativeNumStrStats,
			err
	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&deepCopyNumStrKernel,
		numStrRoundingSpec,
		ePrefix)

	if err != nil {

		return nativeNumStr,
			nativeNumStrStats,
			err
	}

	nativeNumStr,
		nativeNumStrStats,
		err = new(numberStrKernelBoson).
		createNativeNumStrFromNumStrKernel(
			&deepCopyNumStrKernel,
			ePrefix.XCpy(
				"deepCopyNumStrKernel"))

	return nativeNumStr,
		nativeNumStrStats,
		err
}

// getPureNumStr
//
// Receives an instance of NumberStrKernel and extracts
// the numeric value contained therein in order to
// create and return a formatted Pure Number String.
//
// A Pure Number String differs from a Native Number
// String in that it offers more options for
// customization. Pure Number Strings are better able to
// match multinational and multicultural number
// formatting conventions. Users have the option to
// specify custom radix points or decimal separator
// characters as well as designating leading or trailing
// minus signs for negative numbers.
//
// A Pure Number String is defined as follows:
//
//  1. A pure number string consists entirely of numeric
//     digit characters (0-9).
//
//  2. A pure number string will separate integer and
//     fractional digits with a radix point or decimal
//     separator. This could be, but is not limited to,
//     a decimal point ('.'). For example, many European
//     countries use the comma (',') as a radix point.
//
//  3. A pure number string will designate negative values
//     with a minus sign ('-'). This minus sign could be
//     positioned as a leading or trailing minus sign.
//
//  4. A pure number string will NEVER include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
//  5. A pure number string will NEVER include currency
//     symbols.
//
// The numeric value extracted from the NumberStrKernel
// instance passed as input parameter 'numStrKernel' will
// first be rounded according to input parameter
// specifications before generating the final Pure Number
// String.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will NOT change or modify the data values
//	contained in the NumberStrKernel instance passed as
//	input parameter 'numStrKernel'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel.
//		This instance contains the numeric value which
//		will be used to generate and return a Pure Number
//		String.
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol.
//
//		The Decimal Separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a floating point number
//		string.
//
//		In the US, UK, Australia, most of Canada and many
//		other countries the Decimal Separator is the
//		period character ('.') known as the decimal
//		point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//		If this parameter is submitted as an empty or
//		zero length string, an error will be returned.
//
//	leadingMinusSign			bool
//
//		In pure number strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		pure number string returned by this method will
//		format negative numeric values with a leading
//		minus sign ('-') at the beginning of the number
//		string.
//
//		Leading minus signs represent the standard means
//		for designating negative numeric values in the
//		US, UK, Australia, most of Canada and many other
//		countries.
//
//		Example Leading Minus Sign:
//			"-123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		pure number string returned by this method will
//		format negative numeric values with a trailing
//		minus signs ('-') located at the end of the
//		number string.
//
//		Trailing minus signs represent the standard for
//		France, Germany and many countries in the
//		European Union.
//
//		Example Trailing Number Symbols:
//			"123.456-"
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
//		pureNumberStr				string
//
//			If this method completes successfully, a Pure
//			number string containing the numeric value
//			extracted from the current instance of
//			NumberStrKernel will be returned.
//
//			A Pure Number String is defined as follows:
//
//	 	1.	A pure number string consists entirely of numeric
//	 	  	digit characters (0-9).
//
//	 	2.	A pure number string will separate integer and
//	 	  	fractional digits with a radix point or decimal
//	 	  	separator. This could be, but is not limited to,
//	 	  	a decimal point ('.'). For example, many European
//	 	  	countries use the comma (',') as a radix point.
//
//	 	3.	A pure number string will designate negative values
//	 	  	with a minus sign ('-'). This minus sign could be
//	 	  	positioned as a leading or trailing minus sign.
//
//	 	4.	A pure number string will NEVER include integer
//	 	  	separators such as commas (',') to separate
//	 	  	integer digits by thousands.
//
//	 	  				  NOT THIS: 1,000,000
//	 	  		Pure Number String: 1000000
//
//	 	5.	A pure number string will NEVER include currency
//				symbols.
//
//		pureNumStrComponents		PureNumberStrComponents
//
//			If this method completes successfully, this
//			parameter will return an instance of
//			PureNumberStrComponents. This data structure
//			contains an analysis and profile information on
//			the Pure Number String returned by paramter,
//			'pureNumberStr'.
//
//			type PureNumberStrComponents struct {
//
//				NumStrStats NumberStrStatsDto
//
//					This data transfer object will return key
//					statistics on the numeric value encapsulated
//					by the current instance of NumberStrKernel.
//
//					type NumberStrStatsDto struct {
//
//					NumOfIntegerDigits					uint64
//
//						The total number of integer digits to the
//						left of the radix point or, decimal point, in
//						the subject numeric value.
//
//					NumOfSignificantIntegerDigits		uint64
//
//						The number of nonzero integer digits to the
//						left of the radix point or, decimal point, in
//						the subject numeric value.
//
//					NumOfFractionalDigits				uint64
//
//						The total number of fractional digits to the
//						right of the radix point or, decimal point,
//						in the subject numeric value.
//
//					NumOfSignificantFractionalDigits	uint64
//
//						The number of nonzero fractional digits to
//						the right of the radix point or, decimal
//						point, in the subject numeric value.
//
//					NumberValueType 					NumericValueType
//
//						This enumeration value specifies whether the
//						subject numeric value is classified either as
//						an integer or a floating point number.
//
//						Possible enumeration values are listed as
//						follows:
//							NumValType.None()
//							NumValType.FloatingPoint()
//							NumValType.Integer()
//
//					NumberSign							NumericSignValueType
//
//						An enumeration specifying the number sign
//						associated with the numeric value. Possible
//						values are listed as follows:
//							NumSignVal.None()		= Invalid Value
//							NumSignVal.Negative()	= -1
//							NumSignVal.Zero()		=  0
//							NumSignVal.Positive()	=  1
//
//					IsZeroValue							bool
//
//						If 'true', the subject numeric value is equal
//						to zero ('0').
//
//						If 'false', the subject numeric value is
//						greater than or less than zero ('0').
//					}
//
//
//
//				AbsoluteValueNumStr string
//
//				The number string expressed as an absolute value.
//				Be advised, this number string may be a floating
//				point number string containing fractional digits.
//
//				AbsoluteValAllIntegerDigitsNumStr string
//
//				Integer and fractional digits are combined
//				in a single number string without a decimal
//				point separating integer and fractional digits.
//				This string DOES NOT contain a leading number
//				sign (a.k.a. minus sign ('-')
//
//				SignedAllIntegerDigitsNumStr string
//
//				Integer and fractional digits are combined
//				in a single number string without a decimal
//				point separating integer and fractional digits.
//				If the numeric value is negative, a leading
//				minus sign will be prefixed at the beginning
//				of the number string.
//
//				NativeNumberStr string
//
//				A Native Number String representing the base
//				numeric value used to generate these profile
//				number string statistics.
//
//				A valid Native Number String must conform to the
//				standardized formatting criteria defined below:
//
//				 	1. A Native Number String Consists of numeric
//				 	   character digits zero through nine inclusive
//				 	   (0-9).
//
//				 	2. A Native Number String will include a period
//				 	   or decimal point ('.') to separate integer and
//				 	   fractional digits within a number string.
//
//				 	   Native Number String Floating Point Value:
//				 	   				123.1234
//
//				 	3. A Native Number String will always format
//				 	   negative numeric values with a leading minus sign
//				 	   ('-').
//
//				 	   Native Number String Negative Value:
//				 	   				-123.2
//
//				 	4. A Native Number String WILL NEVER include integer
//				 	   separators such as commas (',') to separate
//				 	   integer digits by thousands.
//
//				 	   					NOT THIS: 1,000,000
//				 	   		Native Number String: 1000000
//
//				 	5. Native Number Strings will only consist of:
//
//				 	   (a)	Numeric digits zero through nine inclusive (0-9).
//
//				 	   (b)	A decimal point ('.') for floating point
//				 	   		numbers.
//
//				 	   (c)	A leading minus sign ('-') in the case of
//				 	   		negative numeric values.
//
//			}
//
//		err							error
//
//			If this method completes successfully, the
//			returned error Type is set equal to 'nil'. If
//			errors are encountered during processing, the
//			returned error Type will encapsulate an error
//			message.
//
//			If an error message is returned, the text value
//			for input parameter 'errPrefDto' (error prefix)
//			will be prefixed or attached at the beginning of
//			the error message.
func (numStrKernelQuark *numberStrKernelQuark) getPureNumStr(
	numStrKernel *NumberStrKernel,
	decSeparatorChars string,
	leadingMinusSign bool,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	pureNumberStr string,
	pureNumStrComponents PureNumberStrComponents,
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"getPureNumStr()",
		"")

	if err != nil {

		return pureNumberStr,
			pureNumStrComponents,
			err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return pureNumberStr,
			pureNumStrComponents,
			err
	}

	var err2 error
	_,
		err2 = new(numberStrKernelAtom).
		testValidityOfNumStrKernel(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: 'numStrKernel' is invalid!\n"+
			"This instance of NumberStrKernel failed validity tests.\n"+
			"Validation Error: \n%v\n",
			ePrefix.String(),
			err2.Error())

		return pureNumberStr,
			pureNumStrComponents,
			err
	}

	if roundingType == NumRoundType.NoRounding() {

		pureNumberStr,
			pureNumStrComponents,
			err = new(numberStrKernelBoson).
			createPureNumStrFromNumStrKernel(
				numStrKernel,
				decSeparatorChars,
				leadingMinusSign,
				ePrefix.XCpy(
					"numStrKernel"))

		return pureNumberStr,
			pureNumStrComponents,
			err
	}

	var deepCopyNumStrKernel NumberStrKernel

	err = new(numberStrKernelNanobot).
		copy(
			&deepCopyNumStrKernel,
			numStrKernel,
			ePrefix.XCpy(
				"deepCopyNumStrKernel<-"+
					"numStrKernel"))

	if err != nil {

		return pureNumberStr,
			pureNumStrComponents,
			err
	}

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err =
		new(NumStrRoundingSpec).NewRoundingSpec(
			roundingType,
			roundToFractionalDigits,
			ePrefix)

	if err != nil {

		return pureNumberStr,
			pureNumStrComponents,
			err
	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&deepCopyNumStrKernel,
		numStrRoundingSpec,
		ePrefix)

	if err != nil {

		return pureNumberStr,
			pureNumStrComponents,
			err
	}

	pureNumberStr,
		pureNumStrComponents,
		err = new(numberStrKernelBoson).
		createPureNumStrFromNumStrKernel(
			&deepCopyNumStrKernel,
			decSeparatorChars,
			leadingMinusSign,
			ePrefix.XCpy(
				"deepCopyNumStrKernel"))

	return pureNumberStr,
		pureNumStrComponents,
		err
}

//	getSetNumValueType
//
//	Sets and returns the current NumericValueType for the
//	instance of NumberStrKernel passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		method will examine the internal member variables
//		contained in this instance and set the correct
//		value for Numeric Value Type.
//
//		NumericValueType is an enumeration value specifying
//		the type of numeric value contained in the
//		'numStrKernel' instance.
//
//		Possible NumericValueType enumeration values are
//		listed as follows:
//			NumValType.None()
//			NumValType.FloatingPoint()
//			NumValType.Integer()
//
//		The internal variable contained in 'numStrKernel'
//		which will be configured is:
//
//			NumberStrKernel.numberValueType
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
func (numStrKernelQuark *numberStrKernelQuark) getSetNumValueType(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	NumericValueType,
	error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumericValueType := NumValType.None()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"getSetNumValueType()",
		"")

	if err != nil {

		return newNumericValueType, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return newNumericValueType, err
	}

	lenIntegerDigits :=
		numStrKernel.integerDigits.GetRuneArrayLength()

	lenFracDigits :=
		numStrKernel.fractionalDigits.GetRuneArrayLength()

	if lenIntegerDigits == 0 &&
		lenFracDigits == 0 {

		newNumericValueType = NumValType.None()

	} else if lenIntegerDigits > 0 &&
		lenFracDigits == 0 {

		newNumericValueType = NumValType.Integer()

	} else {

		// MUST BE lenFracDigits > 0

		newNumericValueType = NumValType.FloatingPoint()

	}

	numStrKernel.numberValueType = newNumericValueType

	return newNumericValueType, err
}

//	equalizeNStrIntDigitsLengths
//
//	Receives pointers to two instances of
//	NumberStrKernel, 'numStrKernel01' and
//	'numStrKernel02'. This method will ensure that the
//	integer arrays contained in both instances have
//	equal array lengths.
//
//	If the integer arrays do not have equal array
//	lengths, leading zero characters ('0') will be added
//	to configure their array lengths as equal.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel01				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal integer array contained in this instance
//		will be	compared to that of input parameter,
//		'numStrKernel02'. If the 'numStrKernel01' integer
//		array length is shorter than that of
//		'numStrKernel02', leading zero characters ('0')
//		will be added to achieve an equal integer array
//		length with the integer array contained in
//		'numStrKernel02'.
//
//	numStrKernel02				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal integer array contained in this instance
//		will be	compared to that of input parameter,
//		'numStrKernel01'. If the 'numStrKernel02' integer
//		array length is shorter than that of
//		'numStrKernel01', leading zero characters ('0')
//		will be added to achieve an equal integer array
//		length with the integer array contained in
//		'numStrKernel01'.
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
func (numStrKernelQuark *numberStrKernelQuark) equalizeNStrIntDigitsLengths(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"equalizeNStrIntDigitsLengths()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel01 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel01' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if numStrKernel02 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel02' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenIntArray01 :=
		len(numStrKernel01.integerDigits.CharsArray)

	lenIntArray02 :=
		len(numStrKernel02.integerDigits.CharsArray)

	if lenIntArray01 == lenIntArray02 {

		// Nothing to do
		return err
	}

	// Integer Array Lengths Are NOT Equal

	var numOfCharsToAdd int

	if lenIntArray01 > lenIntArray02 {

		numOfCharsToAdd =
			lenIntArray01 - lenIntArray02

		err =
			numStrKernel02.integerDigits.ExtendRuneArray(
				'0',
				numOfCharsToAdd,
				false,
				ePrefix.XCpy(
					"numStrKernel02.integerDigits"))

		if err != nil {

			return err

		}

	} else {
		// MUST BE
		// lenIntArray02 > lenIntArray01

		numOfCharsToAdd =
			lenIntArray02 - lenIntArray01

		err =
			numStrKernel01.integerDigits.ExtendRuneArray(
				'0',
				numOfCharsToAdd,
				false,
				ePrefix.XCpy(
					"numStrKernel01.integerDigits"))

		if err != nil {

			return err

		}

	}

	return err
}

//	equalizeNStrFracDigitsLengths
//
//	Receives pointers to two instances of
//	NumberStrKernel, 'numStrKernel01' and
//	'numStrKernel01'. This method will ensure that the
//	fractional arrays contained in both instances have
//	equal array lengths.
//
//	If the fractional arrays do not have equal array
//	lengths, trailing zero characters ('0') will be added
//	to configure their array lengths as equal.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel01				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal fractional array contained in this instance
//		will be	compared to that of input parameter,
//		'numStrKernel02'. If the 'numStrKernel01' fractional
//		array length is shorter than that of
//		'numStrKernel02', trailing zero characters ('0')
//		will be added to achieve an equal fractional array
//		length with the fractional array contained in
//		'numStrKernel02'.
//
//	numStrKernel02				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal fractional array contained in this instance
//		will be	compared to that of input parameter,
//		'numStrKernel01'. If the 'numStrKernel02' fractional
//		array length is shorter than that of
//		'numStrKernel01', trailing zero characters ('0')
//		will be added to achieve an equal fractional array
//		length with the fractional array contained in
//		'numStrKernel01'.
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
func (numStrKernelQuark *numberStrKernelQuark) equalizeNStrFracDigitsLengths(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"equalizeNStrFracDigitsLengths()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel01 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel01' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if numStrKernel02 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel02' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenFracArray01 :=
		len(numStrKernel01.fractionalDigits.CharsArray)

	lenFracArray02 :=
		len(numStrKernel02.fractionalDigits.CharsArray)

	if lenFracArray01 == lenFracArray02 {

		// Nothing to do
		return err
	}

	// Fractional Digit Array Lengths Are NOT Equal

	var numOfCharsToAdd int

	if lenFracArray01 > lenFracArray02 {

		numOfCharsToAdd =
			lenFracArray01 - lenFracArray02

		err =
			numStrKernel02.fractionalDigits.ExtendRuneArray(
				'0',
				numOfCharsToAdd,
				true,
				ePrefix.XCpy(
					"numStrKernel02.fractionalDigits"))

		if err != nil {

			return err

		}

	} else {
		// MUST BE
		// lenFracArray02 > lenFracArray01

		numOfCharsToAdd =
			lenFracArray02 - lenFracArray01

		err =
			numStrKernel01.fractionalDigits.ExtendRuneArray(
				'0',
				numOfCharsToAdd,
				true,
				ePrefix.XCpy(
					"numStrKernel01.fractionalDigits"))

		if err != nil {

			return err

		}

	}

	return err
}

//	roundNumStrKernel
//
//	This method receives a pointer to an instance of
//	numStrKernel and proceeds to round the numeric
//	value according to the specifications passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will modify the numeric value contained
//	in the NumberStrKernel instance passed as input
//	parameter 'numStrKernel'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		modified and rounded according to the rounding
//		specifications contained in the following input
//		parameters.
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
func (numStrKernelQuark *numberStrKernelQuark) roundNumStrKernel(
	numStrKernel *NumberStrKernel,
	roundingType NumberRoundingType,
	roundToFactionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"roundNumStrKernel()",
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

	if roundingType == NumRoundType.NoRounding() {

		// Nothing to do!
		return err

	}

	if roundToFactionalDigits < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'roundToFactionalDigits' is invalid!\n"+
			"'roundToFactionalDigits' has an integer value less than zero (0).\n"+
			"'roundToFactionalDigits' integer value = '%v'\n",
			ePrefix.String(),
			roundToFactionalDigits)

		return err
	}

	nStrKernelElectron := numberStrKernelElectron{}

	err = nStrKernelElectron.setUninitializedKernelToZero(
		numStrKernel,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {

		return err

	}

	err = nStrKernelElectron.rationalizeFractionalIntegerDigits(
		numStrKernel,
		ePrefix)

	if err != nil {

		return err

	}

	_,
		err = nStrKernelElectron.getSetIsNonZeroValue(
		numStrKernel,
		ePrefix.XCpy("numStrKernel"))

	if err != nil {

		return err

	}

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err = new(NumStrRoundingSpec).NewRoundingSpec(
		roundingType,
		roundToFactionalDigits,
		ePrefix.XCpy(
			"->numStrRoundingSpec"))

	if err != nil {

		return err

	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		numStrKernel,
		numStrRoundingSpec,
		ePrefix.XCpy(
			"->newNumStrKernel"))

	return err
}

// setNumStrKernelFromNativeNumStr
//
// Receives an instance of NumberStrKernel and proceeds
// to reconfigure the internal data elements with the
// numeric value extracted from the Native Number String
// passed as input paramter, 'nativeNumStr'.
//
// The term 'Native' means that the number string format
// is designed to interoperate with the Golang
// programming language library functions and packages.
// Types like 'strconv', 'strings', 'math' and 'big'
// (big.Int, big.Float, big.Rat) routinely parse and
// convert this type of number string to numeric values.
// In addition, Native Number Strings are frequently
// consumed by external library functions such as this
// one (String Mechanics 'strmech') to convert strings to
// numeric values and numeric values to strings.
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
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will delete and overwrite all
//	pre-existing data values in the instance of
//	NumberStrKernel passed as input parameter
//	'numStrKernel'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance will be reconfigured using the numeric
//		value extracted from input parameter, 'nativeNumStr'.
//
//	nativeNumStr				string
//
//		A Native Number String containing the numeric
//		character digits which will be converted to the
//		numeric value used to reconfigure the
//		NumberStrKernel instance passed as input paramter,
//		'numStrKernel'.
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
//		criteria for a Native Number String, an error
//		will be returned.
//
//		A valid Native Number String must conform to the
//		standardized formatting criteria defined below:
//
//	 	1. A Native Number String Consists of numeric
//	 	   character digits zero through nine inclusive
//	 	   (0-9).
//
//	 	2. A Native Number String will include a period
//	 	   or decimal point ('.') to separate integer and
//	 	   fractional digits within a number string.
//
//	 	   Native Number String Floating Point Value:
//	 	   				123.1234
//
//	 	3. A Native Number String will always format
//	 	   negative numeric values with a leading minus sign
//	 	   ('-').
//
//	 	   Native Number String Negative Value:
//	 	   				-123.2
//
//	 	4. A Native Number String WILL NEVER include integer
//	 	   separators such as commas (',') to separate
//	 	   integer digits by thousands.
//
//	 	   					NOT THIS: 1,000,000
//	 	   		Native Number String: 1000000
//
//	 	5. Native Number Strings will only consist of:
//
//	 	   (a)	Numeric digits zero through nine inclusive (0-9).
//
//	 	   (b)	A decimal point ('.') for floating point
//	 	   		numbers.
//
//	 	   (c)	A leading minus sign ('-') in the case of
//	 	   		negative numeric values.
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
func (numStrKernelQuark *numberStrKernelQuark) setNumStrKernelFromNativeNumStr(
	numStrKernel *NumberStrKernel,
	nativeNumStr string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"setNumStrKernelFromNativeNumStr()",
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

	err = new(NumStrHelper).
		IsValidNativeNumStrError(
			nativeNumStr,
			ePrefix.XCpy(
				"nativeNumStr"))

	if err != nil {

		return err
	}

	var numberStats NumberStrStatsDto

	numberStats,
		err = new(numStrMathQuark).
		nativeNumStrToRunes(
			nativeNumStr,
			&numStrKernel.integerDigits,
			&numStrKernel.fractionalDigits,
			ePrefix.XCpy(
				"numberStats<-nativeNumStr"))

	if err != nil {

		return err

	}

	numStrKernel.numberSign = numberStats.NumberSign

	numStrKernel.numberValueType = numberStats.NumberValueType

	numStrKernel.isNonZeroValue = !numberStats.IsZeroValue

	var err2 error
	_,
		err2 = new(numberStrKernelAtom).
		testValidityOfNumStrKernel(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: The new NumberStrKernel configuration failed validity tests.\n"+
			"One or more data values were classified as 'invalid'.\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return err
}

// setNumStrKernelFromPureNumStr
//
// Receives an instance of NumberStrKernel and proceeds
// to reconfigure that instance based on the numeric
// value extracted from a Pure Number String.
//
// A Pure Number String differs from a Native Number
// String in that it offers more options for
// customization. Pure Number Strings are better able to
// match multinational and multicultural number
// formatting conventions. Users have the option to
// specify custom radix points or decimal separator
// characters as well as designating leading or trailing
// minus signs for negative numbers.
//
// A Pure Number String is defined as follows:
//
//  1. A pure number string consists entirely of numeric
//     digit characters (0-9).
//
//  2. A pure number string will separate integer and
//     fractional digits with a radix point or decimal
//     separator. This could be, but is not limited to,
//     a decimal point ('.'). For example, many European
//     countries use the comma (',') as a radix point.
//
//  3. A pure number string will designate negative values
//     with a minus sign ('-'). This minus sign could be
//     positioned as a leading or trailing minus sign.
//
//  4. A pure number string will NEVER include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
//  5. A pure number string will NEVER include currency
//     symbols.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	This method will delete and overwrite all
//	pre-existing data values in the instance of
//	NumberStrKernel passed as input parameter
//	'numStrKernel'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance will be reconfigured using the numeric
//		value extracted from the Pure Number String input
//		parameter, 'pureNumStr'.
//
//	pureNumberStr				string
//
//		This strings contains the numeric character
//		digits from	which a numeric value will be
//		extracted. This numeric value will be used
//		to reconfigure the instance of NumberStrKernel
//		passed as input parameter 'numStrKernel'.
//
//		A "Pure Number String" is defined as follows:
//
//			1.	Consists of numeric character digits
//				zero through nine inclusive (0-9).
//
//			2.	Option: A Pure Number String may include
//				a radix point or decimal separator.
//				Decimal separators separate integer and
//				fractional numeric digits in a pure
//				number string. The decimal separator may
//				consist of one or more text characters.
//
//				In the US, UK, Australia, most of Canada
//				and many other countries, the decimal
//				separator is the period character ('.')
//				known as the decimal point.
//
//				In France, Germany and many countries in
//				the European Union, the Decimal Separator
//				is the comma character (',').
//
//			3.	Optional: A Pure Number String may
//				include a negative number sign symbol
//				consisting of a minus sign ('-'). The
//				minus sign will identify the numeric
//				value contained in the pure number string
//				as a negative number. Only the minus sign
//				('-') classifies a numeric value as a
//				negative number in a Pure Number String.
//
//				If a leading or trailing minus sign ('-')
//				is NOT present in the pure number string,
//				the numeric value is assumed to be
//				positive.
//
//			4.	Only numeric characters, the decimal
//				separator and the minus sign will be
//				processed by the pure number string
//				parsing algorithm. All other characters
//				will be	ignored.
//
//			5.	Pure Number Strings consist of a single
//				numeric value. The entire Pure Number String
//				will be parsed, or processed, and only one
//				numeric value per Pure Number String will
//				be returned.
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol.
//
//		The Decimal Separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted, floating
//		point Number String.
//
//		In the US, UK, Australia, most of Canada and many
//		other countries the Decimal Separator is the
//		period character ('.') known as the decimal
//		point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//		The Decimal Separator is used to identify and
//		separate integer and fractional digits while
//		parsing the Pure Number String.
//
//	leadingMinusSign			bool
//
//		In pure number strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		pure number string parsing algorithm will search
//		for a leading minus sign ('-') at the beginning of
//		the number string. Leading minus signs represent
//		the standard means for designating negative
//		numeric values in the US, UK, Australia, most of
//		Canada and many other parts of world.
//
//		Example Leading Minus Sign:
//			"-123.456" or "- 123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		pure number string parsing algorithm will search
//		for trailing minus signs ('-') located at the end
//		of the number string. Trailing minus signs
//		represent the standard for France, Germany and
//		many countries in the European Union.
//
//		NOTE: Identification of a trailing minus sign in
//		the pure number string input parameter,
//		'pureNumberString', will immediately terminate
//		the search for numeric characters.
//
//		Example Trailing Number Symbols:
//			"123.456-" or "123.456 -"
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
//	pureNumStrComponents		PureNumberStrComponents
//
//		If this method completes successfully, this
//		parameter will return an instance of
//		PureNumberStrComponents. This data structure
//		contains an analysis and profile information on
//		the reconfigured instance of NumberStrKernel
//		passed as input parameter 'numStrKernel'.
//
//		type PureNumberStrComponents struct {
//
//			NumStrStats NumberStrStatsDto
//
//				This data transfer object will return key
//				statistics on the numeric value encapsulated
//				by the current instance of NumberStrKernel.
//
//				type NumberStrStatsDto struct {
//
//				NumOfIntegerDigits					uint64
//
//					The total number of integer digits to the
//					left of the radix point or, decimal point, in
//					the subject numeric value.
//
//				NumOfSignificantIntegerDigits		uint64
//
//					The number of nonzero integer digits to the
//					left of the radix point or, decimal point, in
//					the subject numeric value.
//
//				NumOfFractionalDigits				uint64
//
//					The total number of fractional digits to the
//					right of the radix point or, decimal point,
//					in the subject numeric value.
//
//				NumOfSignificantFractionalDigits	uint64
//
//					The number of nonzero fractional digits to
//					the right of the radix point or, decimal
//					point, in the subject numeric value.
//
//				NumberValueType 					NumericValueType
//
//					This enumeration value specifies whether the
//					subject numeric value is classified either as
//					an integer or a floating point number.
//
//					Possible enumeration values are listed as
//					follows:
//						NumValType.None()
//						NumValType.FloatingPoint()
//						NumValType.Integer()
//
//				NumberSign							NumericSignValueType
//
//					An enumeration specifying the number sign
//					associated with the numeric value. Possible
//					values are listed as follows:
//						NumSignVal.None()		= Invalid Value
//						NumSignVal.Negative()	= -1
//						NumSignVal.Zero()		=  0
//						NumSignVal.Positive()	=  1
//
//				IsZeroValue							bool
//
//					If 'true', the subject numeric value is equal
//					to zero ('0').
//
//					If 'false', the subject numeric value is
//					greater than or less than zero ('0').
//				}
//
//
//
//			AbsoluteValueNumStr string
//
//			The number string expressed as an absolute value.
//			Be advised, this number string may be a floating
//			point number string containing fractional digits.
//
//			AbsoluteValAllIntegerDigitsNumStr string
//
//			Integer and fractional digits are combined
//			in a single number string without a decimal
//			point separating integer and fractional digits.
//			This string DOES NOT contain a leading number
//			sign (a.k.a. minus sign ('-')
//
//			SignedAllIntegerDigitsNumStr string
//
//			Integer and fractional digits are combined
//			in a single number string without a decimal
//			point separating integer and fractional digits.
//			If the numeric value is negative, a leading
//			minus sign will be prefixed at the beginning
//			of the number string.
//
//			NativeNumberStr string
//
//			A Native Number String representing the base
//			numeric value used to generate these profile
//			number string statistics.
//
//			A valid Native Number String must conform to the
//			standardized formatting criteria defined below:
//
//			 	1. A Native Number String Consists of numeric
//			 	   character digits zero through nine inclusive
//			 	   (0-9).
//
//			 	2. A Native Number String will include a period
//			 	   or decimal point ('.') to separate integer and
//			 	   fractional digits within a number string.
//
//			 	   Native Number String Floating Point Value:
//			 	   				123.1234
//
//			 	3. A Native Number String will always format
//			 	   negative numeric values with a leading minus sign
//			 	   ('-').
//
//			 	   Native Number String Negative Value:
//			 	   				-123.2
//
//			 	4. A Native Number String WILL NEVER include integer
//			 	   separators such as commas (',') to separate
//			 	   integer digits by thousands.
//
//			 	   					NOT THIS: 1,000,000
//			 	   		Native Number String: 1000000
//
//			 	5. Native Number Strings will only consist of:
//
//			 	   (a)	Numeric digits zero through nine inclusive (0-9).
//
//			 	   (b)	A decimal point ('.') for floating point
//			 	   		numbers.
//
//			 	   (c)	A leading minus sign ('-') in the case of
//			 	   		negative numeric values.
//
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
func (numStrKernelQuark *numberStrKernelQuark) setNumStrKernelFromPureNumStr(
	numStrKernel *NumberStrKernel,
	pureNumStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	errPrefDto *ePref.ErrPrefixDto) (
	pureNumStrComponents PureNumberStrComponents,
	err error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelQuark."+
			"setNumStrKernelFromPureNumStr()",
		"")

	if err != nil {

		return pureNumStrComponents,
			err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return pureNumStrComponents,
			err
	}

	var numberStats NumberStrStatsDto

	var decSepCharsRuneArray RuneArrayDto

	decSepCharsRuneArray,
		err = new(RuneArrayDto).NewString(
		decSeparatorChars,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"<-decSepCharsRuneArray"))

	if err != nil {

		return pureNumStrComponents,
			err
	}

	numberStats,
		err = new(numStrMathQuark).
		pureNumStrToRunes(
			pureNumStr,
			&numStrKernel.integerDigits,
			&numStrKernel.fractionalDigits,
			&decSepCharsRuneArray,
			leadingMinusSign,
			ePrefix.XCpy(
				"pureNumStr"))

	if err != nil {

		return pureNumStrComponents,
			err
	}

	var nativeNumStr string

	if numberStats.NumberSign == NumSignVal.Negative() {
		nativeNumStr = "-"
	}

	nativeNumStr +=
		string(numStrKernel.integerDigits.CharsArray)

	if len(numStrKernel.fractionalDigits.CharsArray) > 0 {

		nativeNumStr += "."

		nativeNumStr +=
			string(numStrKernel.fractionalDigits.CharsArray)

	}

	pureNumStrComponents,
		err = new(NumStrMath).
		PureNumStrToComponents(
			nativeNumStr,
			".",
			true,
			ePrefix.XCpy(
				"<-base"))

	if err != nil {

		return pureNumStrComponents,
			err
	}

	if !pureNumStrComponents.NumStrStats.Equal(&numberStats) {

		err = fmt.Errorf("%v\n"+
			"Error: Actual Number Stats NOT EQUAL to Computed Number Stats!\n"+
			"numberStats = \n%v\n"+
			"pureNumStrComponents = \n%v\n",
			ePrefix.String(),
			numberStats.String(),
			pureNumStrComponents.String())

		return pureNumStrComponents,
			err

	}

	err = pureNumStrComponents.NumStrStats.CopyIn(
		&numberStats,
		ePrefix.XCpy(
			"pureNumStrComponents.NumStrStats<-numberStats"))

	if err != nil {

		return pureNumStrComponents,
			err
	}

	numStrKernel.numberSign = pureNumStrComponents.NumStrStats.NumberSign

	numStrKernel.numberValueType = pureNumStrComponents.NumStrStats.NumberValueType

	numStrKernel.isNonZeroValue =
		!pureNumStrComponents.NumStrStats.IsZeroValue

	var err2 error

	_,
		err2 = new(numberStrKernelAtom).
		testValidityOfNumStrKernel(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: The new NumberStrKernel configuration failed validity tests.\n"+
			"One or more data values were classified as 'invalid'.\n"+
			"Error=\n%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return pureNumStrComponents, err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type MathFloatHelper struct {
	lock *sync.Mutex
}

//	FloatToIntFracRunes
//
//	Receives one of several types of floating point
//	values and converts that value to an integer digit
//	rune array and a fractional digit rune array.
//
//	The integer and fractional digit rune arrays
//	represent and absolute values extracted from the
//	original floating point number.
//
//	The returned integer and fractional digits are stored
//	in input parameters 'intDigits' and 'fracDigits'.
//
//	The positive or negative number sign for the returned
//	numeric digits can be determined by examining the
//	statistics returned by parameter 'numberStats'
//	(numberStats.NumberSign).
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	floatingPointNumber 		interface{}
//
//		Numeric values passed by means of this empty
//		interface MUST BE convertible to one of the
//		following types:
//
//			float32
//			float64
//			*big.Float
//
//		If 'floatingPointNumber' is NOT convertible to
//		one of the types listed above, an error will be
//		returned.
//
//	intDigits					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		integer numeric digits extracted from
//		'floatingPointNumber' will be stored as text
//		characters in the rune array encapsulated by
//		this RuneArrayDto object.
//
//		The positive or negative number sign for the
//		extracted integer digits, can be determined by
//		examining the statistics returned by parameter
//		'numberStats' (numberStats.NumberSign).
//
//	fracDigits					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		fractional numeric digits extracted from
//		'floatingPointNumber' will be stored as text
//		characters in the rune array encapsulated by
//		this RuneArrayDto object.
//
//		The positive or negative number sign for the
//		extracted integer digits, can be determined by
//		examining the statistics returned by parameter
//		'numberStats' (numberStats.NumberSign).
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numberStats					NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the integer and fractional digits extracted
//		from 'floatingPointNumber' and stored in the
//		'intDigits' and 'fracDigits' RuneArrayDto
//		objects.
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
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathFloatHelper *MathFloatHelper) FloatNumToIntFracRunes(
	floatingPointNumber interface{},
	intDigits *RuneArrayDto,
	fracDigits *RuneArrayDto,
	errorPrefix interface{}) (
	numberStats NumberStrStatsDto,
	err error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"FloatNumToIntFracRunes()",
		"")

	if err != nil {
		return numberStats, err
	}

	numberStats,
		err = new(mathFloatHelperMechanics).
		floatNumToIntFracRunes(
			floatingPointNumber,
			intDigits,
			fracDigits,
			ePrefix)

	return numberStats, err
}

//	PiTo20k
//
//	Returns an instance of *big.Float configured for Pi
//	up to 20k fractional digits.
//
//	Pi to 20,001 digits. Including the integer '3' this
//	is 20,001 digits. There are 20,000 fractional digits.
//
//	OEIS A000796
//
//	https://oeis.org/A000796
//	https://oeis.org/A000796/b000796.txt
//
//	If the user sets input parameter 'roundingType' to
//	NumRoundType.NoRounding(), the entire 20,000
//	fractional digits will be configured and returned
//	as an instance of *big.Float.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*big.Float
//
//		A pointer to an instance of big.Float. If this
//		method completes successfully, this instance will
//		be configured with the value of Pi out to the
//		specified number of decimal places.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathFloatHelper *MathFloatHelper) PiTo20k(
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	*big.Float,
	error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	pi20k := new(big.Float).
		SetInt64(0).
		SetPrec(66504).
		SetMode(big.AwayFromZero).
		SetInt64(0)

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"PiTo20k()",
		"")

	if err != nil {
		return pi20k, err
	}

	if roundToFractionalDigits > 20000 {

		err = fmt.Errorf("\n%v\n"+
			"Error: Input parameter 'roundToFractionalDigits' is invalid!\n"+
			"'roundToFractionalDigits' exceeds the maximum limit of 20,000.\n"+
			"roundToFractionalDigits = '%v'\n",
			ePrefix.String(),
			roundToFractionalDigits)

		return pi20k, err
	}

	numStrKernel := NumberStrKernel{}

	numStrKernel.numberValueType = NumValType.FloatingPoint()

	numStrKernel.numberSign = NumSignVal.Positive()

	numStrKernel.isNonZeroValue = true

	numStrKernel.integerDigits,
		numStrKernel.fractionalDigits = new(MathConstantsFloat).
		Pi20KRunes(roundToFractionalDigits)

	if roundingType != NumRoundType.NoRounding() &&
		roundingType != NumRoundType.None() {

		err = numStrKernel.Round(
			roundingType,
			roundToFractionalDigits,
			ePrefix.XCpy("Pi20KRunes"))

		if err != nil {

			return pi20k, err

		}
	}

	var ok bool
	_,
		ok = pi20k.SetString(
		numStrKernel.GetPureNumberStr(
			".",
			true))

	if !ok {

		err = fmt.Errorf("\n%v\n"+
			"Error: pi20k.SetString(numStrKernel.GetPureNumberStr()) FAILED!\n"+
			"big.Float was unable to set the Pi value to\n"+
			"%v fractional digits.\n"+
			"numStrKernel.GetPureNumberStr() = %v",
			ePrefix.String(),
			roundToFractionalDigits,
			numStrKernel.GetPureNumberStr(
				".",
				true))
	}

	return pi20k, err
}

//	DigitsToPrecisionEstimate
//
//	Computes an estimate of the number of precision
//	bits required in order to store a given number
//	of numeric digits in a type big.Float, floating
//	point number.
//
//	Precision bits are used in the configuration of
//	big.Float types. The conversion factor is
//	"3.3219789132197891321978913219789".
//
//		Conversion Factor  x  Numeric Digit Capacity =
//				Precision Bits
//			(margin of error +/- 16)
//
//	The number of precision bits returned is an
//	estimate with a margin of error of plus or minus
//	sixteen (+ or - 16).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numNumericDigitsRequired	int64
//
//		The number of numeric digits to be stored and
//		processed by a type big.Float floating point
//		numeric value. This value represents the desired
//		capacity for a big.Float number. This number of
//		numeric digits should include both integer and
//		fractional numeric digits as well as a buffer
//		of extra digits necessary to perform accurate
//		calculations. The number of buffer digits will
//		vary depending on the complexity of pending
//		calculations.
//
//		If this value is less than one (+1), an error
//		will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	precisionBits				uint
//
//		Precision bits defines the number of bits in the
//		mantissa of a big.Float numeric value. The number
//		of precision bits controls the number of integer
//		and fractional numeric digits that can be stored
//		in an instance of big.Float.
//
//		If this method completes successfully, the value
//		returned will represent the estimated number of
//		precision bits required to store and process
//		the number of numerical digits specified by input
//		parameter, 'numNumericDigitsRequired'.
//
//		This estimate for precision bits has a margin of
//		error of plus or minus sixteen bits (+ or - 16).
//
//		The value of 'precisionBits' returned by this
//		method will always be a multiple of eight (+8).
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
func (mathFloatHelper *MathFloatHelper) DigitsToPrecisionEstimate(
	numNumericDigitsRequired int64,
	errorPrefix interface{}) (
	precisionBits uint,
	err error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"DigitsToPrecisionEstimate()",
		"")

	if err != nil {

		return precisionBits, err
	}

	precisionBits,
		err = new(mathFloatHelperPreon).
		estimateDigitsToPrecision(
			numNumericDigitsRequired,
			ePrefix)

	return precisionBits, err
}

//	PrecisionBitsFromRequiredDigits
//
//	Generates the number of precision bits in the
//	mantissa of a big.Float number based on the
//	number of numerical digits required to produce
//	an accurate calculation result.
//
//	Be advised that the number of mantissa precision bits
//	required to store a process an accurate numeric value
//	includes both integer and fractional numeric digits.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//
//	requiredIntegerDigits		int64
//
//		The number of integer digits required for the
//		pending calculation.
//
//		If this parameter has a value less than zero,
//		an error will be returned.
//
//		If the sum of parameters 'requiredIntegerDigits'
//		and 'requiredFractionalDigits' is equal to zero,
//		an error will be returned.
//
//	requiredFractionalDigits	int64
//
//		The number of fractional digits required to
//		ensure accuracy for the pending calculation.
//
//		If this parameter has a value less than zero,
//		an error will be returned.
//
//		If the sum of parameters 'requiredIntegerDigits'
//		and 'requiredFractionalDigits' is equal to zero,
//		an error will be returned.
//
//	requestedBufferDigits		int64
//
//		The number of extra numerical digits required to
//		ensure accuracy for the pending calculation. It
//		is generally a good idea to add space for extra
//		numerical digits to accommodate rounding and/or
//		complex numerical calculations.
//
//		If this parameter has a value less than zero,
//		an error will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	precisionBits				uint
//
//		If this method completes successfully, this
//		parameter will return the number of precision
//		bits required to store and accurately process
//		the number of numerical digits identified by
//		input parameters, 'requiredIntegerDigits',
//		'requiredFractionalDigits' and
//		'requestedBufferDigits'.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathFloatHelper *MathFloatHelper) PrecisionBitsFromRequiredDigits(
	requiredIntegerDigits,
	requiredFractionalDigits,
	requestedBufferDigits int64,
	errorPrefix interface{}) (
	precisionBits uint,
	err error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"PrecisionBitsFromRequiredDigits()",
		"")

	if err != nil {

		return precisionBits, err
	}

	precisionBits,
		err = new(mathFloatHelperAtom).precisionBitsFromRequiredDigits(
		requiredIntegerDigits,
		requiredFractionalDigits,
		requestedBufferDigits,
		ePrefix)

	return precisionBits, err
}

//	PrecisionToDigitsEstimate
//
//	Computes an estimate of the number of numerical
//	digits which can be stored given the number of
//	precision bits configured for a type big.Float,
//	floating point number.
//
//	Precision bits are used in the configuration of
//	big.Float types. The conversion factor is:
//		"3.3219789132197891321978913219789"
//
//		Precision Bits / Conversion Factor =
//				Numeric Digit Capacity
//			(margin of error +/- 3)
//
//	The number of numerical digits returned is an
//	estimate with a margin of error of plus or minus
//	three (+ or - 3) numeric digits.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	precisionBits				uint
//
//		The number of bits of precision in the mantissa
//		of a big.Float floating point numeric value.
//
//		If this value is less than eight (+8), an error
//		will be returned.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	totalNumOfNumericalDigits	int64
//
//		If this method completes successfully, the value
//		returned will represent the estimated total
//		number of numerical digits which can be stored
//		in a big.Float floating point number mantissa
//		configured for the number of Precision Bits
//		specified by input parameter 'precisionBits'.
//
//		This estimate has a margin of error of plus or
//		minus three (+ or - 3) numeric digits.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathFloatHelper *MathFloatHelper) PrecisionToDigitsEstimate(
	precisionBits uint,
	errorPrefix interface{}) (
	totalNumOfNumericalDigits int64,
	err error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"PrecisionToDigitsEstimate()",
		"")

	if err != nil {

		return totalNumOfNumericalDigits, err
	}

	totalNumOfNumericalDigits,
		err = new(mathFloatHelperPreon).
		estimatePrecisionToDigits(
			precisionBits,
			ePrefix)

	return totalNumOfNumericalDigits, err
}

// PrecisionToDigitsFactor
//
// Returns an instance of *big.Float configured with the
// "Precision To Digits" conversion factor.
//
// Precision bits are used in the configuration of
// big.Float types. The conversion factor is
// "3.3219789132197891321978913219789".
//
//		Precision Bits / Conversion Factor =
//				Numeric Digit Capacity
//			(margin of error +/- 3)
//
//	Conversely:
//
//		Conversion Factor  x  Numeric Digit Capacity =
//				Precision Bits
//			(margin of error +/- 16)
//
//	Precision, as used in connection with type big.Float,
//	specifies the mantissa precision of a number in bits.
//
//	Also, remember that the number of numeric digits
//	identified using this conversion factor includes
//	both integer and fractional digits.
//
//	For information on 'precision bits' and their
//	relevance to type big.Float, reference:
//
//	https://pkg.go.dev/math/big#Float
//
//	Bear in mind that this conversion factor may only be
//	used to generate an estimate of numeric digits
//	associated with a give precision bits value. This
//	estimate may vary from the actual number of numeric
//	digits. This estimate has a margin of error of plus
//	or minus five (+ or - 3).
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	None
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*big.Float
//
//		This method returns a pointer to an instance of
//		big.Float configured with the conversion factor
//		used to convert precision bits to the number of
//		equivalent numeric digits.
func (mathFloatHelper *MathFloatHelper) PrecisionToDigitsFactor() *big.Float {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	return new(mathFloatHelperPreon).
		precisionToDigitsFactor()
}

// NativeNumStrToBigFloat
//
// Receives a Native Number String and converts that
// string to a big.Float value.
//
// The term 'Native' applies in the sense that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math' and
// 'big' (big.Int, big.Float, big.Rat) routinely parse
// and convert this type of number string to numeric
// values. In addition, Native Number Strings are
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
// The input parameter 'nativeNumStr' must be formatted
// as a Native Number String in accordance with the
// following criteria:
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
//     (a)	Numeric digits zero through nine inclusive
//     (0-9).
//
//     (b)	A decimal point ('.') for floating point
//     numbers.
//
//     (c)	A leading minus sign ('-') in the case of
//     negative numeric values.
//
//     If the input parameter 'nativeNumStr' does NOT meet
//     these criteria, an error will be returned.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If the Native Number String ('nativeNumStr') fails to
//	comply with Native Number String formatting
//	requirements try the following method as a means of
//	converting a 'dirty' number string to a valid Native
//	Number String:
//
//			NumStrHelper.DirtyToNativeNumStr()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr				string
//
//		This string contains the Native Number String which
//		will be parsed to produce and return a big.Float
//		value.
//
//		If 'nativeNumStr' fails to meet the criteria for
//		a Native Number String, an error will be
//		returned.
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
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	big.Float
//
//		If this method completes successfully, the pure
//		number string passed as input value 'pureNumStr'
//		will be converted and returned as a big.Float
//		floating point value.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathFloatHelper *MathFloatHelper) NativeNumStrToBigFloat(
	nativeNumStr string,
	errorPrefix interface{}) (
	big.Float,
	error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"NativeNumStrToBigFloat()",
		"")

	if err != nil {
		return big.Float{}, err
	}

	var bigFloatDto BigFloatDto

	bigFloatDto,
		err = new(mathFloatHelperBoson).
		bigFloatDtoFromPureNumStr(
			nativeNumStr,
			".",
			true,
			2,
			0,
			big.ToNearestEven,
			ePrefix)

	if err != nil {
		return big.Float{}, err
	}

	bigFloatNum := big.Float{}

	bigFloatNum.Copy(&bigFloatDto.Value)

	return bigFloatNum, err

	/*
		return new(mathFloatHelperNanobot).
			nativeNumStrToBigFloat(
				nativeNumStr,
				ePrefix.XCpy(
					"nativeNumStr"))
	*/

}

// NativeNumStrToBigFloatDto
//
// Receives a Native Number String containing a numeric
// value which will be converted and returned as a
// big.Float floating point value encapsulated within
// an instance of BigFloatDto.
//
// The term 'Native' applies in the sense that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math' and
// 'big' (big.Int, big.Float, big.Rat) routinely parse
// and convert this type of number string to numeric
// values. In addition, Native Number Strings are
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
// The input parameter 'nativeNumStr' must be formatted
// as a Native Number String in accordance with the
// following criteria:
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
//     (a)	Numeric digits zero through nine inclusive
//     (0-9).
//
//     (b)	A decimal point ('.') for floating point
//     numbers.
//
//     (c)	A leading minus sign ('-') in the case of
//     negative numeric values.
//
//     If the input parameter 'nativeNumStr' does NOT meet
//     these criteria, an error will be returned.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If the Native Number String ('nativeNumStr') fails to
//	comply with Native Number String formatting
//	requirements try the following method as a means of
//	converting a 'dirty' number string to a valid Native
//	Number String:
//
//			NumStrHelper.DirtyToNativeNumStr()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr				string
//
//		This Native Number String contains the numeric
//		character digits which will be analyzed and
//		converted to a big.Float value in the returned
//		instance of 'BigFloatDto'.
//
//		If 'nativeNumStr' fails to meet the criteria for
//		a Native Number String, an error will be
//		returned.
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
//	numOfExtraDigitsBuffer		int64
//
//		When configuring the big.Float numeric value
//		returned by the BigFloatDto instance, the number
//		of big.Float precision bits will be calculated
//		based on the number of integer and fractional
//		numeric digits contained in the Native Number
//		String ('nativeNumStr'). To deal with
//		contingencies and requirements often found in
//		complex floating point operations, users have
//		the option to arbitrarily increase the number
//		of precision bits by specifying additional
//		numeric digits via parameter,
//		'numOfExtraDigitsBuffer'.
//
//		Note: The user has the option of overriding the
//		automatic precision bits calculation by specifying
//		a precision bits value directly through parameter,
//		'precisionBitsOverride'.
//
//	precisionBitsOverride		uint
//
//		The term 'precision bits' refers to the number of
//		bits in the mantissa of a big.Float floating point
//		number. Effectively, 'precision bits' controls the
//		precision, accuracy and numerical digit storage
//		capacity for a big.Float floating point number.
//
//		Typically, this method will automatically
//		calculate the value of big.Float precision bits
//		using the parameter 'numOfExtraDigitsBuffer'
//		listed above. However, if 'precisionBitsOverride'
//		has a value greater than zero, the automatic
//		precision bit calculation will be overridden and
//		big.Float precision bits will be set to the value
//		of this	precision bits specification
//		('precisionBitsOverride').
//
//	roundingMode 				big.RoundingMode
//
//		Specifies the rounding algorithm which will be used
//		internally to calculate the base value raised to the
//		power of exponent.
//
//		Each instance of big.Float is configured with a
//		rounding mode. Input parameter 'roundingMode'
//		controls this configuration for the calculation
//		and the big.Float value returned by this method.
//
//		The constant values available for big.Float
//		rounding mode are listed as follows:
//
//		big.ToNearestEven  		// == IEEE 754-2008 roundTiesToEven
//		big.ToNearestAway       // == IEEE 754-2008 roundTiesToAway
//		big.ToZero              // == IEEE 754-2008 roundTowardZero
//		big.AwayFromZero        // no IEEE 754-2008 equivalent
//		big.ToNegativeInf       // == IEEE 754-2008 roundTowardNegative
//		big.ToPositiveInf       // == IEEE 754-2008 roundTowardPositive
//
//		If in doubt as this setting, 'big.AwayFromZero' or
//		'big.ToNearestEven' are common selections for rounding mode.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	BigFloatDto
//
//		If this method completes successfully, a fully
//		populated instance of BigFloatDto will be
//		returned containing the big.Float value
//		generated from the Native Number String parameter,
//		'nativeNumStr'.
//
//		type BigFloatDto struct {
//			Value big.Float
//				The actual value of the big.Float instance.
//
//			NumStrComponents PureNumberStrComponents
//				This parameter profiles the actual big.Float
//				floating point numeric value identified by
//				structure element 'Value'.
//
//				type PureNumberStrComponents struct {
//
//					NumStrStats NumberStrStatsDto
//
//						This data transfer object will return key
//						statistics on the numeric value encapsulated
//						by the current instance of NumberStrKernel.
//
//							type NumberStrStatsDto struct {
//
//								NumOfIntegerDigits					uint64
//
//									The total number of integer digits to the
//									left of the radix point or, decimal point, in
//									the subject numeric value.
//
//								NumOfSignificantIntegerDigits		uint64
//
//									The number of nonzero integer digits to the
//									left of the radix point or, decimal point, in
//									the subject numeric value.
//
//								NumOfFractionalDigits				uint64
//
//									The total number of fractional digits to the
//									right of the radix point or, decimal point,
//									in the subject numeric value.
//
//								NumOfSignificantFractionalDigits	uint64
//
//									The number of nonzero fractional digits to
//									the right of the radix point or, decimal
//									point, in the subject numeric value.
//
//								NumberValueType 					NumericValueType
//
//									This enumeration value specifies whether the
//									subject numeric value is classified either as
//									an integer or a floating point number.
//
//									Possible enumeration values are listed as
//									follows:
//										NumValType.None()
//										NumValType.FloatingPoint()
//										NumValType.Integer()
//
//								NumberSign							NumericSignValueType
//
//									An enumeration specifying the number sign
//									associated with the numeric value. Possible
//									values are listed as follows:
//										NumSignVal.None()		= Invalid Value
//										NumSignVal.Negative()	= -1
//										NumSignVal.Zero()		=  0
//										NumSignVal.Positive()	=  1
//
//								IsZeroValue							bool
//
//									If 'true', the subject numeric value is equal
//									to zero ('0').
//
//									If 'false', the subject numeric value is
//									greater than or less than zero ('0').
//							}
//
//
//
//					AbsoluteValueNumStr string
//					The number string expressed as an absolute value.
//
//					AllIntegerDigitsNumStr string
//					Integer and fractional digits are combined
//					in a single number string without a decimal
//					point separating integer and fractional digits.
//					This string DOES NOT contain a leading number
//					sign (a.k.a. minus sign ('-')
//				}
//
//			EstimatedPrecisionBits BigFloatPrecisionDto
//
//			This structure stores the components and final
//			results value for a precision bits calculation.
//			The number of precision bits configured for a
//			big.Float floating point numeric value determines
//			the storage capacity for a specific floating
//			point number. As such, the calculation of a
//			correct and adequate precision bits value can
//			affect the accuracy of floating point calculations.
//
//			type BigFloatPrecisionDto struct {
//
//					NumIntegerDigits			int64
//
//						The actual or estimated number of integer digits
//						in a big.Float floating point numeric value. The
//						number of integer digits in a floating point
//						number is one of the elements used to calculate
//						the precision bits required to store that
//						floating point number.
//
//					NumFractionalDigits			int64
//
//						The actual or estimated number of fractional
//						digits in a big.Float floating point numeric
//						value. The number of fractional digits in a
//						floating point number is one of the elements used
//						to calculate the precision bits required to store
//						that floating point number.
//
//					NumOfExtraDigitsBuffer		int64
//
//						When estimating the number of precision necessary
//						to store or process big.Float floating point
//						values, is generally a good idea to include a
//						safety margin consisting of excess numeric digits.
//
//						This parameter stores the number of extra numeric
//						digits used in a calculation of total require
//						precision bits.
//
//					PrecisionBitsSpec uint
//						This parameter represents the estimated number of
//						bits required to store a specific floating point
//						numeric value in an instance of type big.Float.
//
//						The 'PrecisionBitsSpec' value is usually generated
//						by an internal calculation based on the estimated
//						number of integer and fractional digits contained
//						in a big.Float floating point number. However,
//						users have the option to specify an arbitrary
//						precision bits value.
//			}
//
//
//	error
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
func (mathFloatHelper *MathFloatHelper) NativeNumStrToBigFloatDto(
	nativeNumStr string,
	numOfExtraDigitsBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	errorPrefix interface{}) (
	BigFloatDto,
	error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"NativeNumStrToBigFloatDto()",
		"")

	if err != nil {
		return BigFloatDto{}, err
	}

	return new(mathFloatHelperBoson).
		bigFloatDtoFromPureNumStr(
			nativeNumStr,
			".",
			true,
			numOfExtraDigitsBuffer,
			precisionBitsOverride,
			roundingMode,
			ePrefix)
}

// NativeNumStrToFloat64
//
// Receives a Native Number string and converts that
// string to a float64 floating point value.
//
// The term 'Native' applies in the sense that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math' and
// 'big' (big.Int, big.Float, big.Rat) routinely parse
// and convert this type of number string to numeric
// values. In addition, Native Number Strings are
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
// The input parameter 'nativeNumStr' must be formatted
// as a Native Number String in accordance with the
// following criteria:
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
//     (a)	Numeric digits zero through nine inclusive
//     (0-9).
//
//     (b)	A decimal point ('.') for floating point
//     numbers.
//
//     (c)	A leading minus sign ('-') in the case of
//     negative numeric values.
//
//     If the input parameter 'nativeNumStr' does NOT meet
//     these criteria, an error will be returned.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If the Native Number String ('nativeNumStr') fails to
//	comply with Native Number String formatting
//	requirements try the following method as a means of
//	converting a 'dirty' number string to a valid Native
//	Number String:
//
//			NumStrHelper.DirtyToNativeNumStr()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr				string
//
//		This string contains the Native Number String which
//		will be parsed to produce and return a big.Float
//		value.
//
//		If 'nativeNumStr' fails to meet the criteria for
//		a Native Number String, an error will be
//		returned.
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
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	float64
//
//		If this method completes successfully, the pure
//		number string passed as input value 'pureNumStr'
//		will be converted and returned as a float64
//		floating point value.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathFloatHelper *MathFloatHelper) NativeNumStrToFloat64(
	nativeNumStr string,
	errorPrefix interface{}) (
	float64,
	error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var float64Num float64

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"NativeNumStrToFloat64()",
		"")

	if err != nil {

		return float64Num, err
	}

	return new(mathFloatHelperBoson).
		pureNumStrToFloat64(
			nativeNumStr,
			ePrefix.XCpy(
				"nativeNumStr"))
}

//	PureNumStrToBigFloatDto
//
//	Receives a Pure Number String containing a numeric
//	value which will be converted and returned as a
//	big.Float floating point value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberStr				string
//
//		This Pure Number String contains the numeric
//		character digits which will be analyzed and
//		reported in the returned instance of
//		'BigFloatDto'.
//
//		A "Pure Number String" is defined as follows:
//
//			1.	Consists of numeric character digits
//				zero through nine inclusive (0-9).
//
//			2.	Optional: A Pure Number String may
//				include a radix point or decimal
//				separator. Decimal separators separate
//				integer and fractional numeric digits in
//				a Pure Number String. The decimal
//				separator may consist of one or more text
//				characters.
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
//				processed by the Pure Number String
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
//		other countries, the Decimal Separator is the
//		period character ('.') known as the decimal
//		point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingMinusSign			bool
//
//		In Pure Number Strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		Pure Number String parsing algorithm will search
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
//		Pure Number String parsing algorithm will search
//		for trailing minus signs ('-') located at the end
//		of the number string. Trailing minus signs
//		represent the standard for France, Germany and
//		many countries in the European Union.
//
//		NOTE: Identification of a trailing minus sign in
//		the Pure Number String input parameter,
//		'pureNumberString', will immediately terminate
//		the search for additional numeric characters.
//
//		Example Trailing Number Symbols:
//			"123.456-" or "123.456 -"
//
//	numOfExtraDigitsBuffer		int64
//
//		When configuring the big.Float numeric value
//		returned by the BigFloatDto instance,
//		the number of big.Float precision bits will be
//		calculated based on the number of integer and
//		fractional numeric digits contained in the Pure
//		Number String ('pureNumberStr'). To deal with
//		contingencies and requirements often found in
//		complex floating point operations, users have
//		the option to arbitrarily increase the number
//		of precision bits by specifying additional
//		precision bits via parameter,
//		'numOfExtraDigitsBuffer'.
//
//		Note: The user has the option of overriding the
//		automatic precision bits calculation by specifying
//		a precision bits value directly through parameter,
//		'precisionBitsOverride'.
//
//		Specifying a margin of 5-10 digits per 100-digits
//		of string length is recommended.
//
//	precisionBitsOverride		uint
//
//		The term 'precision bits' refers to the number of
//		bits in the mantissa of a big.Float floating point
//		number. Effectively, 'precision bits' controls the
//		precision, accuracy and numerical digit storage
//		capacity for a big.Float floating point number.
//
//		Typically, this method will automatically
//		calculate the value of big.Float precision bits
//		using the parameter 'numOfExtraDigitsBuffer'
//		listed above. However, if 'precisionBitsOverride'
//		has a value greater than zero, the automatic
//		precision bit calculation will be overridden and
//		big.Float precision bits will be set to the value
//		of this	precision bits specification
//		('precisionBitsOverride').
//
//	roundingMode 				big.RoundingMode
//
//		Specifies the rounding algorithm which will be used
//		internally to calculate the base value raised to the
//		power of exponent.
//
//		Each instance of big.Float is configured with a
//		rounding mode. Input parameter 'roundingMode'
//		controls this configuration for the calculation
//		and the big.Float value returned by this method.
//
//		The constant values available for big.Float
//		rounding mode are listed as follows:
//
//		big.ToNearestEven  		// == IEEE 754-2008 roundTiesToEven
//		big.ToNearestAway       // == IEEE 754-2008 roundTiesToAway
//		big.ToZero              // == IEEE 754-2008 roundTowardZero
//		big.AwayFromZero        // no IEEE 754-2008 equivalent
//		big.ToNegativeInf       // == IEEE 754-2008 roundTowardNegative
//		big.ToPositiveInf       // == IEEE 754-2008 roundTowardPositive
//
//		If in doubt as this setting, 'big.AwayFromZero' or
//		'big.ToNearestEven' are common selections for rounding mode.
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
//	BigFloatDto
//
//		If this method completes successfully, a fully
//		populated instance of BigFloatDto will be
//		returned containing the big.Float value
//		generated from the Pure Number String parameter,
//		'pureNumberStr'.
//
//		type BigFloatDto struct {
//			Value big.Float
//				The actual value of the big.Float instance.
//
//			NumStrComponents PureNumberStrComponents
//				This parameter profiles the actual big.Float
//				floating point numeric value identified by
//				structure element 'Value'.
//
//				type PureNumberStrComponents struct {
//
//					NumStrStats NumberStrStatsDto
//
//						This data transfer object will return key
//						statistics on the numeric value encapsulated
//						by the current instance of NumberStrKernel.
//
//							type NumberStrStatsDto struct {
//
//								NumOfIntegerDigits					uint64
//
//									The total number of integer digits to the
//									left of the radix point or, decimal point, in
//									the subject numeric value.
//
//								NumOfSignificantIntegerDigits		uint64
//
//									The number of nonzero integer digits to the
//									left of the radix point or, decimal point, in
//									the subject numeric value.
//
//								NumOfFractionalDigits				uint64
//
//									The total number of fractional digits to the
//									right of the radix point or, decimal point,
//									in the subject numeric value.
//
//								NumOfSignificantFractionalDigits	uint64
//
//									The number of nonzero fractional digits to
//									the right of the radix point or, decimal
//									point, in the subject numeric value.
//
//								NumberValueType 					NumericValueType
//
//									This enumeration value specifies whether the
//									subject numeric value is classified either as
//									an integer or a floating point number.
//
//									Possible enumeration values are listed as
//									follows:
//										NumValType.None()
//										NumValType.FloatingPoint()
//										NumValType.Integer()
//
//								NumberSign							NumericSignValueType
//
//									An enumeration specifying the number sign
//									associated with the numeric value. Possible
//									values are listed as follows:
//										NumSignVal.None()		= Invalid Value
//										NumSignVal.Negative()	= -1
//										NumSignVal.Zero()		=  0
//										NumSignVal.Positive()	=  1
//
//								IsZeroValue							bool
//
//									If 'true', the subject numeric value is equal
//									to zero ('0').
//
//									If 'false', the subject numeric value is
//									greater than or less than zero ('0').
//							}
//
//
//
//					AbsoluteValueNumStr string
//					The number string expressed as an absolute value.
//
//					AllIntegerDigitsNumStr string
//					Integer and fractional digits are combined
//					in a single number string without a decimal
//					point separating integer and fractional digits.
//					This string DOES NOT contain a leading number
//					sign (a.k.a. minus sign ('-')
//				}
//
//			EstimatedPrecisionBits BigFloatPrecisionDto
//
//			This structure stores the components and final
//			results value for a precision bits calculation.
//			The number of precision bits configured for a
//			big.Float floating point numeric value determines
//			the storage capacity for a specific floating
//			point number. As such, the calculation of a
//			correct and adequate precision bits value can
//			affect the accuracy of floating point calculations.
//
//			type BigFloatPrecisionDto struct {
//
//					NumIntegerDigits			int64
//
//						The actual or estimated number of integer digits
//						in a big.Float floating point numeric value. The
//						number of integer digits in a floating point
//						number is one of the elements used to calculate
//						the precision bits required to store that
//						floating point number.
//
//					NumFractionalDigits			int64
//
//						The actual or estimated number of fractional
//						digits in a big.Float floating point numeric
//						value. The number of fractional digits in a
//						floating point number is one of the elements used
//						to calculate the precision bits required to store
//						that floating point number.
//
//					NumOfExtraDigitsBuffer		int64
//
//						When estimating the number of precision necessary
//						to store or process big.Float floating point
//						values, is generally a good idea to include a
//						safety margin consisting of excess numeric digits.
//
//						This parameter stores the number of extra numeric
//						digits used in a calculation of total require
//						precision bits.
//
//					PrecisionBitsSpec uint
//						This parameter represents the estimated number of
//						bits required to store a specific floating point
//						numeric value in an instance of type big.Float.
//
//						The 'PrecisionBitsSpec' value is usually generated
//						by an internal calculation based on the estimated
//						number of integer and fractional digits contained
//						in a big.Float floating point number. However,
//						users have the option to specify an arbitrary
//						precision bits value.
//			}
//
//
//	error
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
func (mathFloatHelper *MathFloatHelper) PureNumStrToBigFloatDto(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	numOfExtraDigitsBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	errorPrefix interface{}) (
	BigFloatDto,
	error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"PureNumStrToBigFloatDto()",
		"")

	if err != nil {
		return BigFloatDto{}, err
	}

	return new(mathFloatHelperBoson).
		bigFloatDtoFromPureNumStr(
			pureNumberStr,
			decSeparatorChars,
			leadingMinusSign,
			numOfExtraDigitsBuffer,
			precisionBitsOverride,
			roundingMode,
			ePrefix)
}

//	RaiseToFloatPositiveExponent
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
//	This method employs floating point mathematics and
//	type big.Float to compute the base floating point
//	value raised to the power of exponent.
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
//	numOfExtraDigitsBuffer		int64
//
//		The term 'precision bits' refers to the number of
//		bits in the mantissa of a big.Float floating
//		point number. Effectively, 'precision bits'
//		controls the precision, accuracy and numerical
//		digit storage capacity for a big.Float floating
//		point number
//
//		When configuring the big.Float numeric value
//		returned by this method, the number of big.Float
//		precision bits will be calculated based on the
//		estimated number of integer and fractional
//		numeric digits contained in the base floating
//		point value ('base'). To deal with contingencies
//		and requirements often found in complex floating
//		point operations, users have the option to
//		arbitrarily increase the number of precision bits
//		by specifying additional numeric digits via
//		parameter, 'numOfExtraDigitsBuffer'.
//
//		The automatic precision bits calculation will add
//		the number of integer digits, fractional digits and
//		'numOfExtraDigitsBuffer' to compute the estimated
//		number of precision bits. As such this parameter
//		provides of safety in terms of ensuring accurate
//		calculation results.
//
//		The greater the value of this parameter, the
//		greater the number of accurate significant digits
//		produced in the calculation result.
//
//		Note: The user has the option of overriding the
//		automatic precision bits calculation by specifying
//		a precision bits value directly through parameter,
//		'precisionBitsOverride'.
//
//	precisionBitsOverride		uint
//
//		The term 'precision bits' refers to the number of
//		bits in the mantissa of a big.Float floating point
//		number. Effectively, 'precision bits' controls the
//		precision, accuracy and numerical digit storage
//		capacity for a big.Float floating point number.
//
//		Typically, this method will automatically
//		calculate the value of big.Float precision bits
//		using the parameter 'numOfExtraDigitsBuffer'
//		listed above. However, if 'precisionBitsOverride'
//		has a value greater than zero, the automatic
//		precision bit calculation will be overridden and
//		big.Float precision bits will be set to the value
//		of this	precision bits specification
//		('precisionBitsOverride').
//
//	roundingMode 				big.RoundingMode
//
//		Specifies the rounding algorithm which will be used
//		internally to calculate the base value raised to the
//		power of exponent.
//
//		Each instance of big.Float is configured with a
//		rounding mode. Input parameter 'roundingMode'
//		controls this configuration for the calculation
//		and the big.Float value returned by this method.
//
//		The constant values available for big.Float
//		rounding mode are listed as follows:
//
//		big.ToNearestEven  		// == IEEE 754-2008 roundTiesToEven
//		big.ToNearestAway       // == IEEE 754-2008 roundTiesToAway
//		big.ToZero              // == IEEE 754-2008 roundTowardZero
//		big.AwayFromZero        // no IEEE 754-2008 equivalent
//		big.ToNegativeInf       // == IEEE 754-2008 roundTowardNegative
//		big.ToPositiveInf       // == IEEE 754-2008 roundTowardPositive
//
//		If in doubt as this setting, 'big.AwayFromZero' or
//		'big.ToNearestEven' are common selections for rounding mode.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (mathFloatHelper *MathFloatHelper) RaiseToFloatPositiveExponent(
	base *big.Float,
	exponent int64,
	numOfExtraDigitsBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	errorPrefix interface{}) (
	*big.Float,
	error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"RaiseToFloatPositiveExponent()",
		"")

	if err != nil {
		return big.NewFloat(0), err
	}

	return new(mathFloatHelperMechanics).
		raiseToFloatExponentConfig(
			base,
			exponent,
			numOfExtraDigitsBuffer,
			precisionBitsOverride,
			roundingMode,
			ePrefix)
}

//	RaiseToIntPositiveExponent
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
//	raised to the power of an exponent. As such it
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
//	numOfExtraDigitsBuffer		int64
//
//		The term 'precision bits' refers to the number of
//		bits in the mantissa of a big.Float floating
//		point number. Effectively, 'precision bits'
//		controls the precision, accuracy and numerical
//		digit storage capacity for a big.Float floating
//		point number.
//
//		When configuring the big.Float numeric value
//		returned by this method, the number of big.Float
//		precision bits will be automatically calculated
//		based on the number of integer and fractional
//		numeric digits contained in the base floating
//		point value ('base'). To deal with contingencies
//		and requirements often found in complex floating
//		point operations, users have the option to
//		arbitrarily increase the number of automatically
//		calculated precision bits by specifying additional
//		numeric digits via parameter,
//		'numOfExtraDigitsBuffer'.
//
//		The automatic precision bits calculation will add
//		the number of integer digits, fractional digits and
//		'numOfExtraDigitsBuffer' to compute the estimated
//		number of precision bits.
//
//		The greater the value of this parameter, the
//		greater the number of accurate significant digits
//		produced in the calculation result.
//
//		Note: The user has the option of overriding the
//		automatic precision bits calculation by specifying
//		a precision bits value directly through parameter,
//		'precisionBitsOverride'.
//
//	precisionBitsOverride		uint
//
//		The term 'precision bits' refers to the number of
//		bits in the mantissa of a big.Float floating point
//		number. Effectively, 'precision bits' controls the
//		precision, accuracy and numerical digit storage
//		capacity for a big.Float floating point number.
//
//		Typically, this method will automatically
//		calculate the value of big.Float precision bits
//		using the parameter 'numOfExtraDigitsBuffer'
//		listed above. However, if 'precisionBitsOverride'
//		has a value greater than zero, the automatic
//		precision bit calculation will be overridden and
//		big.Float precision bits will be set to the value
//		of this	precision bits specification
//		('precisionBitsOverride').
//
//		If in doubt as to this number, identify the
//		total number of integer and fractional digits
//		required to store an accurate result and
//		multiply this number times four (+4) to generate
//		an estimate of precision bits required.
//
//	roundingMode 				big.RoundingMode
//
//		Specifies the rounding algorithm which will be used
//		internally to calculate the base value raised to the
//		power of exponent.
//
//		Each instance of big.Float is configured with a
//		rounding mode. Input parameter 'roundingMode'
//		controls this configuration for the calculation
//		and the big.Float value returned by this method.
//
//		The constant values available for big.Float
//		rounding mode are listed as follows:
//
//		big.ToNearestEven  		// == IEEE 754-2008 roundTiesToEven
//		big.ToNearestAway       // == IEEE 754-2008 roundTiesToAway
//		big.ToZero              // == IEEE 754-2008 roundTowardZero
//		big.AwayFromZero        // no IEEE 754-2008 equivalent
//		big.ToNegativeInf       // == IEEE 754-2008 roundTowardNegative
//		big.ToPositiveInf       // == IEEE 754-2008 roundTowardPositive
//
//		If in doubt as this setting, 'big.AwayFromZero' is a
//		common selection for rounding mode.
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
func (mathFloatHelper *MathFloatHelper) RaiseToIntPositiveExponent(
	base *big.Float,
	exponent int64,
	numOfExtraDigitsBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	errorPrefix interface{}) (
	raisedToExponent *big.Float,
	err error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"RaiseToIntPositiveExponent()",
		"")

	if err != nil {
		return big.NewFloat(0), err
	}

	return new(mathFloatHelperQuark).
		raiseToIntPositiveExponent(
			base,
			exponent,
			numOfExtraDigitsBuffer,
			precisionBitsOverride,
			roundingMode,
			ePrefix)
}

//	RoundBigFloat
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
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (mathFloatHelper *MathFloatHelper) RoundBigFloat(
	numberToRound *big.Float,
	roundedNumber *big.Float,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	err error) {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathFloatHelper."+
			"RoundBigFloat()",
		"")

	if err != nil {
		return err
	}

	err = new(mathFloatHelperQuark).
		roundBigFloat(
			numberToRound,
			roundedNumber,
			roundingType,
			roundToFractionalDigits,
			ePrefix)

	return err
}

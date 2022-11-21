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
			"GetNumericValueStats()",
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
//		fractional numeric digits.
//
//		If this value is less than one (+1), this
//		method will return a value of zero, thereby
//		signaling an error.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	uint
//
//		If input parameter 'numNumericDigitsRequired'
//		has a value less than one (+1), this parameter
//		will return a value of zero (0) signaling an
//		error.
//
//		Otherwise, the value returned will represent the
//		estimated number of precision bits required for
//		the mantissa of a big.Float value i
//		be stored given the value of input parameter,
//		'precisionBits'. This estimate has a margin of
//		error of plus or minus sixteen bits (+ or - 16).
func (mathFloatHelper *MathFloatHelper) DigitsToPrecisionEstimate(
	numNumericDigitsRequired int64) uint {

	if mathFloatHelper.lock == nil {
		mathFloatHelper.lock = new(sync.Mutex)
	}

	mathFloatHelper.lock.Lock()

	defer mathFloatHelper.lock.Unlock()

	return new(mathFloatHelperPreon).
		estimateDigitsToPrecision(
			numNumericDigitsRequired)
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

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

// mathFloatHelperMechanics
//
// Provides helper methods for type MathFloatHelper
type mathFloatHelperMechanics struct {
	lock *sync.Mutex
}

//	floatNumToIntFracRunes
//
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
func (mathFloatHelpMech *mathFloatHelperMechanics) floatNumToIntFracRunes(
	floatingPointNumber interface{},
	intDigits *RuneArrayDto,
	fracDigits *RuneArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	numberStats NumberStrStatsDto,
	err error) {

	if mathFloatHelpMech.lock == nil {
		mathFloatHelpMech.lock = new(sync.Mutex)
	}

	mathFloatHelpMech.lock.Lock()

	defer mathFloatHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperMechanics."+
			"floatNumToIntFracRunes()",
		"")

	if err != nil {

		return numberStats, err
	}

	if intDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'intDigits' is a nil pointer!\n",
			ePrefix.String())

		return numberStats, err
	}

	if fracDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'fracDigits' is a nil pointer!\n",
			ePrefix.String())

		return numberStats, err
	}

	var signedPureNumStr string

	signedPureNumStr,
		err = new(mathFloatHelperAtom).
		floatNumToSignedPureNumStr(
			floatingPointNumber,
			ePrefix.XCpy(
				"floatingPointNumber"))

	if err != nil {

		return numberStats, err
	}

	var decSeparatorChars RuneArrayDto

	decSeparatorChars,
		err = new(RuneArrayDto).NewString(
		".",
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix)

	if err != nil {

		return numberStats, err
	}

	numberStats,
		err = new(numStrMathQuark).
		pureNumStrToRunes(
			signedPureNumStr,
			intDigits,
			fracDigits,
			&decSeparatorChars,
			true,
			ePrefix)

	return numberStats, err
}

//	raiseToFloatExponentConfig
//
//	Receives a pointer to a big.Float floating point
//	number and raises that number to the power specified
//	by input parameter, 'exponent'.
//
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
//
//	This method performs formatting and setup
//	preparations on the 'base' floating point value
//	to ensure accurate results.
//
//	Currently, this method will only process positive
//	exponent. Therefore, if input parameter 'exponent'
//	is less than zero, an error will be returned.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	base							*big.Float
//
//		This floating point value will be raised to the
//		power of 'exponent' and returned to the calling
//		function.
//
//	exponent						int64
//
//		This value will be used to raise 'base' to the
//		power of 'exponent'.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
//
//		Currently, this method will only process positive
//		exponent. Therefore, if input parameter 'exponent'
//		is less than zero, an error will be returned.
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
//		parameter will return 'base' value raised to
//		the power of the 'exponent' value.
//
//		Example:	3.2 ^ 4 = 104.8576
//					base ^ exponent = raisedToExponent
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
func (mathFloatHelpMech *mathFloatHelperMechanics) raiseToFloatExponentConfig(
	base *big.Float,
	exponent int64,
	numOfExtraDigitsBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	errPrefDto *ePref.ErrPrefixDto) (
	*big.Float,
	error) {

	if mathFloatHelpMech.lock == nil {
		mathFloatHelpMech.lock = new(sync.Mutex)
	}

	mathFloatHelpMech.lock.Lock()

	defer mathFloatHelpMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperMechanics."+
			"raiseToFloatExponentConfig()",
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

	baseStr := base.Text('f', -1)

	var pureNumStrComponents PureNumberStrComponents

	pureNumStrComponents,
		err = new(numStrMathAtom).
		pureNumStrToComponents(
			baseStr,
			".",
			true,
			ePrefix.XCpy(
				"<-base"))

	if err != nil {

		return big.NewFloat(0), err
	}

	if pureNumStrComponents.NumStrStats.IsZeroValue {

		// zero^exponent = 0
		return big.NewFloat(0), err

	}

	if exponent == 0 {

		// num^0 = 1
		return big.NewFloat(1), err

	}

	var bFloatDto BigFloatDto

	bFloatDto,
		err = new(mathFloatHelperBoson).
		bigFloatFromPureNumStr(
			baseStr,
			".",
			true,
			numOfExtraDigitsBuffer,
			precisionBitsOverride,
			roundingMode,
			ePrefix)

	bFloatStr := bFloatDto.Value.Text('f', -1)
	bFloatPrec := bFloatDto.Value.Prec()

	fmt.Printf("bFloatStr #1: %v\n"+
		"bFloat Prec: %v\n",
		bFloatStr,
		bFloatPrec)

	requiredIntegerDigits :=
		int64(bFloatDto.NumStrComponents.NumStrStats.NumOfIntegerDigits) *
			exponent

	requiredFractionalDigits :=
		int64(bFloatDto.NumStrComponents.NumStrStats.NumOfFractionalDigits) *
			exponent

	var estimatedPrecisionBits uint

	if precisionBitsOverride == 0 {

		estimatedPrecisionBits,
			err = new(mathFloatHelperAtom).
			precisionBitsFromRequiredDigits(
				requiredIntegerDigits,
				requiredFractionalDigits,
				numOfExtraDigitsBuffer,
				ePrefix)

	} else {

		estimatedPrecisionBits = precisionBitsOverride
	}

	return new(mathFloatHelperQuark).
		raiseToFloatPositiveExponent(
			&bFloatDto.Value,
			exponent,
			estimatedPrecisionBits,
			roundingMode,
			ePrefix)
}

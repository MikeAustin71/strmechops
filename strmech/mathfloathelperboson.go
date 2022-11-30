package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

type mathFloatHelperBoson struct {
	lock *sync.Mutex
}

//	bigFloatFromPureNumStr
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
//	BigFloatDto
//
//		If this method completes successfully, fully
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
func (floatHelperBoson *mathFloatHelperBoson) bigFloatFromPureNumStr(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	numOfExtraDigitsBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	errPrefDto *ePref.ErrPrefixDto) (
	BigFloatDto,
	error) {

	if floatHelperBoson.lock == nil {
		floatHelperBoson.lock = new(sync.Mutex)
	}

	floatHelperBoson.lock.Lock()

	defer floatHelperBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var bFloatDto BigFloatDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperBoson."+
			"bigFloatFromPureNumStr()",
		"")

	if err != nil {
		return bFloatDto, err
	}

	if len(pureNumberStr) == 0 {
		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'pureNumberStr'\n"+
			"is a zero length string and INVALID!\n",
			ePrefix.String())

		return bFloatDto, err
	}

	bFloatDto.NumStrComponents,
		err = new(numStrMathAtom).
		pureNumStrToComponents(
			pureNumberStr,
			decSeparatorChars,
			leadingMinusSign,
			ePrefix)

	if err != nil {
		return bFloatDto, err
	}

	if precisionBitsOverride == 0 {

		bFloatDto.EstimatedPrecisionBits.PrecisionBitsSpec,
			err = new(mathFloatHelperAtom).precisionBitsFromRequiredDigits(
			int64(bFloatDto.NumStrComponents.NumStrStats.NumOfIntegerDigits),
			int64(bFloatDto.NumStrComponents.NumStrStats.NumOfFractionalDigits),
			numOfExtraDigitsBuffer,
			ePrefix)

	} else {

		bFloatDto.EstimatedPrecisionBits.PrecisionBitsSpec =
			precisionBitsOverride
	}

	bFloatDto.Value.SetInt64(0).
		SetPrec(bFloatDto.EstimatedPrecisionBits.
			PrecisionBitsSpec).
		SetMode(roundingMode)

	var ok bool
	_,
		ok = bFloatDto.Value.SetString(pureNumberStr)

	if !ok {

		err = fmt.Errorf("\n%v\n"+
			"Error: bFloatDto.Value.SetString(pureNumberStr) Failed!\n"+
			"pureNumberStr = %v\n",
			ePrefix.String(),
			pureNumberStr)

		return bFloatDto, err
	}

	bFloatDto.Value.SetPrec(bFloatDto.Value.MinPrec())

	if bFloatDto.Value.Acc() != big.Exact {

		err = fmt.Errorf("\n%v\n"+
			"Accuracy Error\n"+
			"An exact floating pointing number value could NOT\n"+
			"be calculated accurately from the Pure Number string\n"+
			"input parameter 'pureNumberStr',\n"+
			"pureNumberValueStr = %v\n"+
			"Accuracy = %v",
			ePrefix.String(),
			pureNumberStr,
			bFloatDto.Value.Acc())

	}

	return bFloatDto, err
}

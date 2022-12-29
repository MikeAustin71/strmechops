package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strconv"
	"sync"
)

type mathFloatHelperBoson struct {
	lock *sync.Mutex
}

//	bigFloatDtoFromPureNumStr
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
//		If in doubt, it recommended that this parameter
//		be set to two ('2').
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
//		If in doubt as to this setting, 'big.ToNearestEven'
//		is the default since its enumeration value is zero.
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
func (floatHelperBoson *mathFloatHelperBoson) bigFloatDtoFromPureNumStr(
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
			"bigFloatDtoFromPureNumStr()",
		"")

	if err != nil {
		return bFloatDto, err
	}

	lenPureNumStr := len(pureNumberStr)

	if lenPureNumStr == 0 {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'pureNumberStr' is invalid!\n"+
			"'pureNumberStr' is a zero length string.\n",
			ePrefix.String())

		return bFloatDto, err
	}

	if numOfExtraDigitsBuffer < 0 {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'numOfExtraDigitsBuffer' is invalid!\n"+
			"'numOfExtraDigitsBuffer' is a less than zero!\n"+
			"numOfExtraDigitsBuffer = %v\n",
			ePrefix.String(),
			numOfExtraDigitsBuffer)

		return bFloatDto, err

	}

	isValidPureNumStr := false

	for i := 0; i < lenPureNumStr; i++ {

		if pureNumberStr[i] >= '0' &&
			pureNumberStr[i] <= '9' {

			isValidPureNumStr = true

			break
		}

	}

	if !isValidPureNumStr {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'pureNumberStr' is INVALID!\n"+
			"'pureNumberStr' contains NO Numeric Digit Characters.\n",
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

		if err != nil {
			return bFloatDto, err
		}

	} else {
		// MUST BE -
		// precisionBitsOverride > 0

		bFloatDto.EstimatedPrecisionBits.PrecisionBitsSpec =
			precisionBitsOverride
	}

	bFloatDto.Value = big.Float{}

	bFloatDto.Value.
		SetPrec(bFloatDto.EstimatedPrecisionBits.PrecisionBitsSpec)

	var ok bool
	_,
		ok = bFloatDto.Value.SetString(pureNumberStr)

	if !ok {

		err = fmt.Errorf("\n%v\n"+
			"Error: bFloatDto.Value.SetString(pureNumberStr) Failed!\n"+
			"pureNumberStr = %v\n"+
			"roundingMode = '%v'\n",
			ePrefix.String(),
			pureNumberStr,
			bFloatDto.Value.Mode())

		return bFloatDto, err
	}

	if !bFloatDto.Value.IsInt() {

		bFloatDto.Value.SetPrec(
			bFloatDto.Value.MinPrec())
	}

	bFloatDto.Value.SetMode(roundingMode)

	return bFloatDto, err
}

// nativeNumStrToBigFloat
//
// Receives a Native Number String ('nativeNumStr') and
// then converts and returns that string as a big.Float
// floating point numeric value.
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
// as a Native Number String which is defined as follows:
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
//     If the input parameter 'nativeNumStr' does NOT meet these
//     criteria, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr					string
//
//		This string contains the Native Number String which
//		will be parsed to produce and return a big.Float
//		value.
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
//	numOfExtraDigitsBuffer		int64
//
//		When configuring the big.Float numeric value
//		returned by the BigFloatDto instance,
//		the number of big.Float precision bits will be
//		calculated based on the number of integer and
//		fractional numeric digits contained in the Pure
//		Number String ('nativeNumStr'). To deal with
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
//		If in doubt, it recommended that this parameter
//		be set to two ('2').
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
//		Specifies the rounding algorithm which will be
//		used internally to configure the returned
//		big.Float value.
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
//		If in doubt as to this setting, the default applied
//		by Golang is big.ToNearestEven because in has an
//		enumeration integer value of zero.
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
//	big.Float
//
//		If this method completes successfully, the pure
//		number string passed as input value 'nativeNumStr'
//		will be converted and returned as a big.Float
//		floating point value.
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
func (floatHelperBoson *mathFloatHelperBoson) nativeNumStrToBigFloat(
	nativeNumStr string,
	numOfExtraDigitsBuffer int64,
	precisionBitsOverride uint,
	roundingMode big.RoundingMode,
	errPrefDto *ePref.ErrPrefixDto) (
	big.Float,
	error) {

	if floatHelperBoson.lock == nil {
		floatHelperBoson.lock = new(sync.Mutex)
	}

	floatHelperBoson.lock.Lock()

	defer floatHelperBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	bigFloatNum := big.Float{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperBoson."+
			"nativeNumStrToBigFloat()",
		"")

	if err != nil {
		return bigFloatNum, err
	}

	lenNativeNumStr := len(nativeNumStr)

	if lenNativeNumStr == 0 {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'nativeNumStr' is INVALID!\n"+
			"'nativeNumStr' is a zero length string.\n",
			ePrefix.String())

		return bigFloatNum, err
	}

	if numOfExtraDigitsBuffer < 0 {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'numOfExtraDigitsBuffer' is invalid!\n"+
			"'numOfExtraDigitsBuffer' is a less than zero!\n"+
			"numOfExtraDigitsBuffer = %v\n",
			ePrefix.String(),
			numOfExtraDigitsBuffer)

		return bigFloatNum, err
	}

	isValidNativeNumStr := false

	for i := 0; i < lenNativeNumStr; i++ {

		if nativeNumStr[i] >= '0' &&
			nativeNumStr[i] <= '9' {

			isValidNativeNumStr = true

			break
		}

	}

	if !isValidNativeNumStr {

		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'nativeNumStr' is INVALID!\n"+
			"'nativeNumStr' contains NO Numeric Digit Characters.\n",
			ePrefix.String())

		return bigFloatNum, err

	}

	var pureNumStrComponents PureNumberStrComponents

	pureNumStrComponents,
		err = new(numStrMathAtom).
		pureNumStrToComponents(
			nativeNumStr,
			".",
			true,
			ePrefix)

	if err != nil {
		return bigFloatNum, err
	}

	var precisionBitsSpec uint

	if precisionBitsOverride == 0 {
		precisionBitsSpec,
			err = new(mathFloatHelperAtom).precisionBitsFromRequiredDigits(
			int64(pureNumStrComponents.NumStrStats.NumOfIntegerDigits),
			int64(pureNumStrComponents.NumStrStats.NumOfFractionalDigits),
			numOfExtraDigitsBuffer,
			ePrefix)

		if err != nil {
			return bigFloatNum, err
		}

	} else {
		// MUST BE -
		// precisionBitsOverride > 0

		precisionBitsSpec = precisionBitsOverride
	}

	bigFloatNum.SetPrec(precisionBitsSpec)

	var ok bool

	_,
		ok = bigFloatNum.SetString(nativeNumStr)

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error:  bigFloatNum.SetString(nativeNumStr)\n"+
			"SetString() FAILED to generate a valid big.Float value.\n"+
			"nativeNumStr = '%v'\n"+
			"roundingMode = '%v'\n",
			ePrefix.String(),
			nativeNumStr,
			bigFloatNum.Mode())

		return bigFloatNum, err
	}

	if !bigFloatNum.IsInt() {

		bigFloatNum.SetPrec(
			bigFloatNum.MinPrec())
	}

	bigFloatNum.SetMode(roundingMode)

	return bigFloatNum, err
}

//	pureNumStrToFloat64
//
//	Receives a pure number string and converts that
//	string to a float64 floating point value.
//
//	The input parameter 'pureNumStr' must be formatted
//	as a pure number string which is defined as follows:
//
//		1.	The pure number string must consist entirely
//			of numeric digit characters (0-9), with
//			following exceptions.
//
//		2.	For floating point values, the pure number
//			string must separate integer and fractional
//			digits with a decimal point ('.').
//
//		3.	The pure number string must designate
//			negative values with a leading minus sign
//			('-').
//
//		4.	The pure number string must NOT include integer
//			separators such as commas (',') to separate
//			integer digits by thousands.
//
//						  NOT THIS: 1,000,000
//				Pure Number String: 1000000
//
//	If the input parameter 'pureNumStr' does NOT meet these
//	criteria, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumStr					string
//
//		This string contains the pure number string which
//		will be parsed to produce and return a float64
//		floating point value.
//
//		The input parameter 'pureNumStr' must be formatted
//		as a pure number string which is defined as follows:
//
//			1.	The pure number string must consist entirely
//				of numeric digit characters (0-9), with
//				following exceptions.
//
//			2.	For floating point values, the pure number
//				string must separate integer and fractional
//				digits with a decimal point ('.').
//
//			3.	The pure number string must designate
//				negative values with a leading minus sign
//				('-').
//
//			4.	The pure number string must NOT include integer
//				separators such as commas (',') to separate
//				integer digits by thousands.
//
//							  NOT THIS: 1,000,000
//					Pure Number String: 1000000
//
//		If the input parameter 'pureNumStr' does NOT meet these
//		criteria, an error will be returned.
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (floatHelperBoson *mathFloatHelperBoson) pureNumStrToFloat64(
	pureNumStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	float64,
	error) {

	if floatHelperBoson.lock == nil {
		floatHelperBoson.lock = new(sync.Mutex)
	}

	floatHelperBoson.lock.Lock()

	defer floatHelperBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperBoson."+
			"pureNumStrToFloat64()",
		"")

	if err != nil {
		return 0.0, err
	}

	var err2 error
	var float64Num float64

	float64Num,
		err2 = strconv.ParseFloat(
		pureNumStr,
		64)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error returned by strconv.ParseFloat(pureNumStr)\n"+
			"pureNumStr = '%v'\n"+
			"Error = \n'%v'\n",
			ePrefix.String(),
			pureNumStr,
			err2.Error())

		return 0.0, err
	}

	return float64Num, err
}

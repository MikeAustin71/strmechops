package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrMath
//
// Consists of methods designed to assist in the
// mathematics of number strings.
type NumStrMath struct {
	lock *sync.Mutex
}

//	NativeNumStrStats
//
// Receives and analyzes the numeric digits configured
// in a Native Number String. The results of this
// analysis, including number sign, number type and key
// statistics relating to the numeric value contained in
// the Native Number String will be returned by an
// instance of type, NumberStrStatsDto.
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
// # Input Parameters
//
//	nativeNumStr				string
//
//		A Native Number String. The numeric digits in
//		this string will be analyzed, profiled and
//		described in the returned instance of
//		'NumberStrStatsDto'.
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
//	NumberStrStatsDto
//
//		If this method completes successfully, this
//		parameter will return an instance of
//		NumberStrStatsDto. This data structure contains
//		an analysis and detail information on the Native
//		Number String passed as input paramter,
//		'nativeNumStr'.
//
//		type NumberStrStatsDto struct {
//
//			NumOfIntegerDigits					uint64
//
//				The total number of integer digits to the
//				left of the radix point or, decimal point, in
//				the subject numeric value.
//
//			NumOfSignificantIntegerDigits		uint64
//
//				The number of integer digits to the left of
//				the radix point, excluding leading zeros, in
//				the subject numeric value.
//
//			NumOfFractionalDigits				uint64
//
//				The total number of fractional digits to the
//				right of the radix point or, decimal point,
//				in the subject numeric value.
//
//			NumOfSignificantFractionalDigits	uint64
//
//				The number of nonzero fractional digits to
//				the right of the radix point or, decimal
//				point, in the subject numeric value.
//
//			NumberValueType 					NumericValueType
//
//				This enumeration value specifies whether the
//				subject numeric value is classified either as
//				an integer or a floating point number.
//
//				Possible enumeration values are listed as
//				follows:
//					NumValType.None()
//					NumValType.FloatingPoint()
//					NumValType.Integer()
//
//			NumberSign							NumericSignValueType
//
//				An enumeration specifying the number sign
//				associated with the numeric value. Possible
//				values are listed as follows:
//					NumSignVal.None()		= Invalid Value
//					NumSignVal.Negative()	= -1
//					NumSignVal.Zero()		=  0
//					NumSignVal.Positive()	=  1
//
//			IsZeroValue							bool
//
//				If 'true', the subject numeric value is equal
//				to zero ('0').
//
//				If 'false', the subject numeric value is
//				greater than or less than zero ('0').
//		}
//
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
func (nStrMath *NumStrMath) NativeNumStrStats(
	nativeNumStr string,
	errorPrefix interface{}) (
	NumberStrStatsDto,
	error) {

	if nStrMath.lock == nil {
		nStrMath.lock = new(sync.Mutex)
	}

	nStrMath.lock.Lock()

	defer nStrMath.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrMath."+
			"NativeNumStrStats()",
		"")

	if err != nil {
		return NumberStrStatsDto{}, err
	}

	return new(numStrMathAtom).pureNumStrStats(
		nativeNumStr,
		".",
		true,
		ePrefix)
}

// PureNumStrStats
//
//	Receives and analyzes the numeric digits configured
//	in a Pure Number String. The results of this
//	analysis, including number sign, number type and key
//	statistics relating to the numeric value contained in
//	the Pure Number String will be returned by an
//	instance of type, NumberStrStatsDto.
//
//	A Pure Number String consists of numeric digits, an
//	optional radix point (a.k.a. decimal point) and an
//	optional leading or trailing minus ('-') sign.
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
//		PureNumberStrComponents.
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
//	NumberStrStatsDto
//
//		If this method completes successfully, this
//		parameter will return an instance of
//		NumberStrStatsDto. This data structure contains
//		an analysis and detail information on the Pure
//		Number String passed as input paramter,
//		'pureNumberStr'.
//
//		type NumberStrStatsDto struct {
//
//			NumOfIntegerDigits					uint64
//
//				The total number of integer digits to the
//				left of the radix point or, decimal point, in
//				the subject numeric value.
//
//			NumOfSignificantIntegerDigits		uint64
//
//				The number of integer digits to the left of
//				the radix point, excluding leading zeros, in
//				the subject numeric value.
//
//			NumOfFractionalDigits				uint64
//
//				The total number of fractional digits to the
//				right of the radix point or, decimal point,
//				in the subject numeric value.
//
//			NumOfSignificantFractionalDigits	uint64
//
//				The number of nonzero fractional digits to
//				the right of the radix point or, decimal
//				point, in the subject numeric value.
//
//			NumberValueType 					NumericValueType
//
//				This enumeration value specifies whether the
//				subject numeric value is classified either as
//				an integer or a floating point number.
//
//				Possible enumeration values are listed as
//				follows:
//					NumValType.None()
//					NumValType.FloatingPoint()
//					NumValType.Integer()
//
//			NumberSign							NumericSignValueType
//
//				An enumeration specifying the number sign
//				associated with the numeric value. Possible
//				values are listed as follows:
//					NumSignVal.None()		= Invalid Value
//					NumSignVal.Negative()	= -1
//					NumSignVal.Zero()		=  0
//					NumSignVal.Positive()	=  1
//
//			IsZeroValue							bool
//
//				If 'true', the subject numeric value is equal
//				to zero ('0').
//
//				If 'false', the subject numeric value is
//				greater than or less than zero ('0').
//		}
//
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
func (nStrMath *NumStrMath) PureNumStrStats(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	errorPrefix interface{}) (
	NumberStrStatsDto,
	error) {

	if nStrMath.lock == nil {
		nStrMath.lock = new(sync.Mutex)
	}

	nStrMath.lock.Lock()

	defer nStrMath.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrMath."+
			"PureNumStrStats()",
		"")

	if err != nil {
		return NumberStrStatsDto{}, err
	}

	return new(numStrMathAtom).pureNumStrStats(
		pureNumberStr,
		decSeparatorChars,
		leadingMinusSign,
		ePrefix)
}

//	PureNumStrToComponents
//
//	Receives and analyzes the numeric digits configured
//	in a Pure Number String. The results of this analysis,
//	including number sign, number type, absolute value
//	number string and key statistics relating to the
//	numeric value contained in the Pure Number String
//	will be returned by an instance of type,
//	PureNumberStrComponents.
//
//	A Pure Number String consists of numeric digits, an
//	optional radix point (a.k.a. decimal point) and an
//	optional leading or trailing minus ('-') sign.
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
//		PureNumberStrComponents.
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
//	PureNumberStrComponents
//
//		If this method completes successfully, this
//		parameter will return an instance of
//		PureNumberStrComponents. This data structure
//		contains an analysis and detail information on
//		the Pure Number String passed as input paramter,
//		'pureNumberStr'.
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
func (nStrMath *NumStrMath) PureNumStrToComponents(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	errorPrefix interface{}) (
	PureNumberStrComponents,
	error) {

	if nStrMath.lock == nil {
		nStrMath.lock = new(sync.Mutex)
	}

	nStrMath.lock.Lock()

	defer nStrMath.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrMath."+
			"PureNumStrToComponents()",
		"")

	if err != nil {
		return PureNumberStrComponents{}, err
	}

	return new(numStrMathAtom).pureNumStrToComponents(
		pureNumberStr,
		decSeparatorChars,
		leadingMinusSign,
		ePrefix)
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelBoson - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelBoson struct {
	lock *sync.Mutex
}

// createNativeNumStrFromNumStrKernel
//
// This low level function extracts a Native Number
// String from an instance of NumberStrKernel passed as
// input parameter 'numStrKernel'.
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
//		This method assumes that 'numStrKernel' has
//		already been tested for validity and is in fact
//		a valid instance of numStrKernel.
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
func (numStrKernelBoson *numberStrKernelBoson) createNativeNumStrFromNumStrKernel(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	nativeNumStr string,
	nativeNumStrStats NumberStrStatsDto,
	err error) {

	if numStrKernelBoson.lock == nil {
		numStrKernelBoson.lock = new(sync.Mutex)
	}

	numStrKernelBoson.lock.Lock()

	defer numStrKernelBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelBoson."+
			"createNativeNumStrFromNumStrKernel()",
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

	if numStrKernel.numberSign == NumSignVal.Negative() {

		nativeNumStr += "-"
	}

	if len(numStrKernel.integerDigits.CharsArray) > 0 {

		nativeNumStr +=
			string(numStrKernel.integerDigits.CharsArray)

	} else {

		nativeNumStr += "0"
	}

	if len(numStrKernel.fractionalDigits.CharsArray) > 0 {

		nativeNumStr += "."

		nativeNumStr +=
			string(numStrKernel.fractionalDigits.CharsArray)
	}

	nativeNumStr,
		nativeNumStrStats,
		err = new(NumStrHelper).
		NormalizeNativeNumStr(
			nativeNumStr,
			ePrefix.XCpy(
				"<-nativeNumStr"))

	return nativeNumStr,
		nativeNumStrStats,
		err
}

// createPureNumStrFromNumStrKernel
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
// No rounding is performed on the numeric value
// extracted from the NumberStrKernel instance passed as
// 'numStrKernel'.
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
func (numStrKernelBoson *numberStrKernelBoson) createPureNumStrFromNumStrKernel(
	numStrKernel *NumberStrKernel,
	decSeparatorChars string,
	leadingMinusSign bool,
	errPrefDto *ePref.ErrPrefixDto) (
	pureNumberStr string,
	pureNumStrComponents PureNumberStrComponents,
	err error) {

	if numStrKernelBoson.lock == nil {
		numStrKernelBoson.lock = new(sync.Mutex)
	}

	numStrKernelBoson.lock.Lock()

	defer numStrKernelBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelBoson."+
			"createPureNumStrFromNumStrKernel()",
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

	if len(decSeparatorChars) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'decSeparatorChars' is invalid!\n"+
			"'decSeparatorChars' is a zero length or empty string.",
			ePrefix.String())

		return pureNumberStr,
			pureNumStrComponents,
			err
	}

	if leadingMinusSign == true &&
		numStrKernel.numberSign == NumSignVal.Negative() {

		pureNumberStr += "-"
	}

	pureNumberStr +=
		numStrKernel.integerDigits.GetCharacterString()

	if len(numStrKernel.fractionalDigits.CharsArray) > 0 {

		pureNumberStr += decSeparatorChars

		pureNumberStr +=
			numStrKernel.fractionalDigits.GetCharacterString()
	}

	if leadingMinusSign == false &&
		numStrKernel.numberSign == NumSignVal.Negative() {

		pureNumberStr += "-"
	}

	pureNumStrComponents,
		err = new(NumStrMath).
		PureNumStrToComponents(
			pureNumberStr,
			".",
			true,
			ePrefix.XCpy(
				"<-base"))

	return pureNumberStr,
		pureNumStrComponents,
		err
}

// normalizeNumericDigits
//
// Removes leading integer zeros and trailing fractional
// zeros from an instance of NumberStrKernel.
func (numStrKernelBoson *numberStrKernelBoson) normalizeNumericDigits(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelBoson.lock == nil {
		numStrKernelBoson.lock = new(sync.Mutex)
	}

	numStrKernelBoson.lock.Lock()

	defer numStrKernelBoson.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelBoson."+
			"normalizeNumericDigits()",
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

	badZerosCount :=
		numStrKernel.integerDigits.GetCountLeadingZeros()

	if badZerosCount > 0 {

		err = numStrKernel.integerDigits.
			DeleteLeadingTrailingChars(
				badZerosCount,
				false,
				ePrefix.XCpy(
					"numStrKernel.integerDigits"))

		if err != nil {

			return err

		}
	}

	badZerosCount =
		numStrKernel.fractionalDigits.GetCountTrailingZeros()

	if badZerosCount > 0 {

		err = numStrKernel.fractionalDigits.
			DeleteLeadingTrailingChars(
				badZerosCount,
				true,
				ePrefix.XCpy(
					"numStrKernel.fractionalDigits"))

		if err != nil {

			return err

		}

	}

	return err
}

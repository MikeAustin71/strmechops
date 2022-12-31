package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrHelperElectron
//
// Provides number string utility methods
type numStrHelperElectron struct {
	lock *sync.Mutex
}

// normalizeNativeNumStr
//
// Removes leading integer zeros and trailing fractional
// zeros from a Native Number String and returns a clean
// or 'normalized' version of the Native Number String.
//
// The term 'Native Number String' means that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math' and
// 'big' (big.Int, big.Float, big.Rat) routinely parse
// and convert this type of number string to numeric
// values. In addition, Native Number Strings are
// frequently consumed by external library functions such
// as this one (String Mechanics 'strmech') in order to
// convert strings to numeric values and numeric values
// to strings.
//
// A Native Number String which has been normalized,
// will contain no leading integer zeros and no trailing
// fractional zeros.
//
//	Examples:
//		Non-Standard Native Number Strings:
//					0001234
//					1234.56780000
//					0001234.5678000
//
//		Normalized Native Number String:
//					1234.5678
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nonStandardNativeNumStr		string
//
//		A Native Number String which may contain leading
//		integer zeros and or trailing fractional zeros.
//			Examples:
//				Non-Standard Native Number Strings:
//						0001234
//						1234.56780000
//						0001234.5678000
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
//		This method will analyze input parameter
//		'nonStandardNativeNumStr' and return a
//		clean or 'normalized' version of the Native
//		Number String by deleting all leading integer
//		zeros and trailing fractional zeros.
//
//		If 'nonStandardNativeNumStr' fails to meet the
//		formatting criteria for a Native Number String,
//		an error will be returned.
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
//	errPrefDto						*ePref.ErrPrefixDto
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
//	normalizedNativeNumStr		string
//
//		If this method completes successfully, a
//		normalized Native Number String extracted
//		from input parameter 'nonStandardNativeNumStr'
//		will be returned.
//
//		A valid 'normalized' Native Number String will
//		meet the following criteria:
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
//		6.	A 'Normalized' Native Number String will
//			contain no leading integer zeros.
//
//										 NOT THIS: 0001234
//				Normalized Native Number String: 1234
//
//		7.	A 'Normalized' Native Number String will
//			contain no trailing fractional zeros.
//
//										 NOT THIS: 12.34000
//				Normalized Native Number String: 12.34
//
//	normalizedNativeNumStrStats	NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the integer and fractional digits contained
//		in the return parameter
//		'normalizedNativeNumStr'.
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
//	err								error
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
func (nStrHelperElectron *numStrHelperElectron) normalizeNativeNumStr(
	nonStandardNativeNumStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	normalizedNativeNumStr string,
	normalizedNativeNumStrStats NumberStrStatsDto,
	err error) {

	if nStrHelperElectron.lock == nil {
		nStrHelperElectron.lock = new(sync.Mutex)
	}

	nStrHelperElectron.lock.Lock()

	defer nStrHelperElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrHelperElectron."+
			"normalizeNativeNumStr()",
		"")

	if err != nil {

		return normalizedNativeNumStr,
			normalizedNativeNumStrStats,
			err
	}

	_,
		err = new(numStrHelperQuark).
		testValidityOfNativeNumStr(
			nonStandardNativeNumStr,
			"nonStandardNativeNumStr",
			ePrefix.XCpy(
				"nonStandardNativeNumStr"))

	if err != nil {

		return normalizedNativeNumStr,
			normalizedNativeNumStrStats,
			err
	}

	var integerDigits RuneArrayDto
	var fractionalDigits RuneArrayDto

	normalizedNativeNumStrStats,
		err = new(numStrMathQuark).
		nativeNumStrToRunes(
			nonStandardNativeNumStr,
			&integerDigits,
			&fractionalDigits,
			ePrefix.XCpy(
				"nonStandardNativeNumStr"))

	if err != nil {

		return normalizedNativeNumStr,
			normalizedNativeNumStrStats,
			err
	}

	reformatStr := false

	if normalizedNativeNumStrStats.NumOfIntegerDigits !=
		normalizedNativeNumStrStats.NumOfSignificantIntegerDigits {

		reformatStr = true

		var numOfLeadingZerosCount uint64

		numOfLeadingZerosCount =
			integerDigits.GetCountLeadingZeros()

		err = integerDigits.DeleteLeadingTrailingChars(
			numOfLeadingZerosCount,
			false,
			ePrefix.XCpy(
				"integerDigits<-numOfLeadingZerosCount"))

		if err != nil {

			return normalizedNativeNumStr,
				normalizedNativeNumStrStats,
				err
		}

	}

	if normalizedNativeNumStrStats.NumOfFractionalDigits !=
		normalizedNativeNumStrStats.NumOfSignificantFractionalDigits {

		reformatStr = true

		var numOfTrailingZerosCount uint64

		numOfTrailingZerosCount =
			fractionalDigits.GetCountTrailingZeros()

		err = fractionalDigits.DeleteLeadingTrailingChars(
			numOfTrailingZerosCount,
			true,
			ePrefix.XCpy(
				"fractionalDigits<-numOfTrailingZerosCount"))

		if err != nil {

			return normalizedNativeNumStr,
				normalizedNativeNumStrStats,
				err
		}

	}

	if reformatStr == false {

		normalizedNativeNumStr = nonStandardNativeNumStr

		return normalizedNativeNumStr,
			normalizedNativeNumStrStats,
			err

	}

	// MUST BE -
	// reformatStr == true

	if normalizedNativeNumStrStats.NumberSign ==
		NumSignVal.Negative() {

		normalizedNativeNumStr = "-"
	}

	if len(integerDigits.CharsArray) == 0 {

		normalizedNativeNumStr += "0"

	} else {

		normalizedNativeNumStr +=
			string(integerDigits.CharsArray)

	}

	if len(fractionalDigits.CharsArray) > 0 {

		normalizedNativeNumStr += "."

		normalizedNativeNumStr +=
			string(fractionalDigits.CharsArray)

	}

	normalizedNativeNumStrStats,
		err = new(numStrMathQuark).
		nativeNumStrToRunes(
			normalizedNativeNumStr,
			&integerDigits,
			&fractionalDigits,
			ePrefix.XCpy(
				"normalizedNativeNumStr"))

	return normalizedNativeNumStr,
		normalizedNativeNumStrStats,
		err
}

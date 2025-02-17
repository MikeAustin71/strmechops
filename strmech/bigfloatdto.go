package strmech

import (
	"math/big"
)

// BigFloatDto
//
// This Data Transfer Object (DTO) is used to transmit
// specifications and information related to floating
// point numeric value of type big.Float.
type BigFloatDto struct {
	Value big.Float
	//	The actual value of the big.Float instance.

	NumStrComponents PureNumberStrComponents
	//	This parameter profiles the actual big.Float
	//	floating point numeric value identified by
	//	structure element 'Value'.
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
	//
	//
	//		}

	EstimatedPrecisionBits BigFloatPrecisionDto
	//	This structure stores the components and final
	//	results value for a precision bits calculation.
	//	The number of precision bits configured for a
	//	big.Float floating point numeric value determines
	//	the storage capacity for a specific floating
	//	point number. As such, the calculation of a
	//	correct and adequate precision bits value can
	//	affect the accuracy of floating point calculations.
	//
	//	type BigFloatPrecisionDto struct {
	//
	//		NumIntegerDigits			int64
	//
	//			The actual or estimated number of integer digits
	//			in a big.Float floating point numeric value. The
	//			number of integer digits in a floating point
	//			number is one of the elements used to calculate
	//			the precision bits required to store that
	//			floating point number.
	//
	//		NumFractionalDigits			int64
	//
	//			The actual or estimated number of fractional
	//			digits in a big.Float floating point numeric
	//			value. The number of fractional digits in a
	//			floating point number is one of the elements used
	//			to calculate the precision bits required to store
	//			that floating point number.
	//
	//		NumOfExtraDigitsBuffer		int64
	//
	//			When estimating the number of precision necessary
	//			to store or process big.Float floating point
	//			values, is generally a good idea to include a
	//			safety margin consisting of excess numeric digits.
	//
	//			This parameter stores the number of extra numeric
	//			digits used in a calculation of total require
	//			precision bits.
	//
	//		PrecisionBitsSpec	uint
	//
	//			This parameter stores the estimated number of
	//			bits required to store a specific floating point
	//			numeric value in an instance of type big.Float.
	//
	//			This value is used to specify the number of
	//			precision bits configured for a big.Float floating
	//			point numeric value.
	//	}
	//
}

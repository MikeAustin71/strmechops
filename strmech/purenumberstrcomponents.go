package strmech

type PureNumberStrComponents struct {
	NumStrStats NumberStrStatsDto
	//
	//		This data transfer object will return key
	//		statistics on the numeric value encapsulated
	//		in a number string.
	//
	//		type NumberStrStatsDto struct {
	//
	//		NumOfIntegerDigits					uint64
	//
	//			The total number of integer digits to the
	//			left of the radix point or, decimal point, in
	//			the subject number string.
	//
	//		NumOfSignificantIntegerDigits		uint64
	//
	//			The number of integer digits to the left of
	//			the radix point, excluding leading zeros, in
	//			the subject number string.
	//
	//		NumOfFractionalDigits				uint64
	//
	//			The total number of fractional digits to the
	//			right of the radix point or, decimal point,
	//			in the subject number string.
	//
	//		NumOfSignificantFractionalDigits	uint64
	//
	//			The number of fractional digits to the right
	//			of the radix point, excluding trailing zeros,
	//			in the subject number string.
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

	AbsoluteValueNumStr string
	//	The number string expressed as an absolute value.
	//	Be advised, this number string may be a floating
	//	point number string containing fractional digits.

	AbsoluteValAllIntegerDigitsNumStr string
	//	Integer and fractional digits are combined
	//	in a single number string without a decimal
	//	point separating integer and fractional digits.
	//	This string DOES NOT contain a leading number
	//	sign (a.k.a. minus sign ('-'). It is therefore
	//	rendered as an absolute or positive value.

	SignedAllIntegerDigitsNumStr string
	//	Integer and fractional digits are combined
	//	in a single number string without a decimal
	//	point separating integer and fractional digits.
	//	If the numeric value is negative, a leading
	//	minus sign will be prefixed at the beginning
	//	of the number string.

	NativeNumberStr string
	//	A Native Number String representing the base
	//	numeric value used to generate these profile
	//	number string statistics.
	//
	//	A valid Native Number String must conform to the
	//	standardized formatting criteria defined below:
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

}

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

	AllIntegerDigitsNumStr string
	//	Integer and fractional digits are combined
	//	in a single number string without a decimal
	//	point separating integer and fractional digits.
	//	This string DOES NOT contain a leading number
	//	sign (a.k.a. minus sign ('-')
}

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

	ActualNumStrComponents PureNumberStrComponents
	//	This parameter profiles the big.Float floating
	//	point numeric value identified by structure
	//	element 'Value'.
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
	//			The number string expressed as an absolute value.
	//
	//			AllIntegerDigitsNumStr string
	//			Integer and fractional digits are combined
	//			in a single number string without a decimal
	//			point separating integer and fractional digits.
	//			This string DOES NOT contain a leading number
	//			sign (a.k.a. minus sign ('-')
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
}

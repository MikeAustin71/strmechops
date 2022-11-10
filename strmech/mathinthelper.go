package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type MathIntHelper struct {
	lock *sync.Mutex
}

//	IntegerToAbsValueNumStr
//
//	Receives one of several types of integer values
//	and converts that value to a pure number string
//	containing the absolute value of the original integer
//	input value.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	intNumericValue				interface{}
//
//		Integer numeric values passed by means of this
//		empty interface MUST BE convertible to one of the
//		following types:
//
//			int8
//			int16
//			int32
//			int	(currently equivalent to int32)
//			int64
//			uint8
//			uint16
//			uint32
//			uint (currently equivalent to uint32)
//			uint64
//			*big.Int
//
//		If parameter 'intNumericValue' is NOT convertible
//		to one of the types listed above, an error will
//		be returned.
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
//	absValueNumberStr			string
//
//		The integer input parameter, 'intNumericValue',
//		will be converted to an absolute value and
//		returned by 'absValueNumberStr' as a string of
//		numeric digits.
//
//	numberStrStats				NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the returned absolute value number string,
//		'absValueNumberStr'.
//
//		Most notably, the number sign, positive or
//		negative, associated with return parameter,
//		'absValueNumberStr', will be specified by
//		'numberStrStats.NumberSign'.
//
//		Since 'absValueNumberStr' represents an integer
//		value, statistics on fractional digits are not
//		relevant.
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
func (mathIntHelper *MathIntHelper) IntegerToAbsValueNumStr(
	intNumericValue interface{},
	errorPrefix interface{}) (
	absValueNumberStr string,
	numberStrStats NumberStrStatsDto,
	err error) {

	if mathIntHelper.lock == nil {
		mathIntHelper.lock = new(sync.Mutex)
	}

	mathIntHelper.lock.Lock()

	defer mathIntHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathIntHelper."+
			"IntegerToAbsValueNumStr()",
		"")

	if err != nil {
		return absValueNumberStr, numberStrStats, err
	}

	absValueNumberStr,
		numberStrStats,
		err = new(mathIntHelperMechanics).
		intToAbsoluteValueStr(
			intNumericValue,
			ePrefix)

	return absValueNumberStr, numberStrStats, err
}

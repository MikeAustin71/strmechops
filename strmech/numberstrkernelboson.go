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
		RationalizeNativeNumStr(
			nativeNumStr,
			ePrefix.XCpy(
				"<-nativeNumStr"))

	return nativeNumStr,
		nativeNumStrStats,
		err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

// mathBigRatHelperMechanics
//
// Provides helper methods for type MathBigRatHelper.
type mathBigRatHelperMechanics struct {
	lock *sync.Mutex
}

// bigRatToNativeNumStr
//
// Receives a pointer to a big.Rat numeric value and
// converts that value to a Native Number String.
//
// The term 'Native Number String' means that the number
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
// The 'Native' Number String returned by this method
// therefore implements a standardized format defined as
// follows:
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
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Before being converted to Native Number String, the
//	big.Rat numeric value will be rounded by the Golang
//	'big' package function:
//
//			func (*Rat) FloatString
//
//	The rounding algorithm used is described as:
//
//		The last digit is rounded to nearest, with halves
//		rounded away from zero.
//			https://pkg.go.dev/math/big#Rat
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bigRatNum					*big.Rat
//
//		A pointer to an instance of the numeric value
//		type big.Rat. This numeric value will be
//		converted to, and returned as, a Native Number
//		String.
//
//		Before being converted to a Native Number String,
//		this numeric value will be rounded by the Golang
//		'big' package functions as specified by input
//		parameter, 'roundToFractionalDigits'.
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this
//		parameter controls the number of digits to the
//		right of the radix point or decimal separator
//		(a.k.a. decimal point). Effectively this defines
//		the number of fractional digits remaining after
//		completion of the number rounding operation
//		performed by the Golang package functions.
//
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
//	nativeNumStr			string
//
//		If this method completes successfully, a Native
//		Number String representing the numeric value
//		passed as input	parameter 'numericValue' will be
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
//					Native Number String Floating Point Value:
//									123.1234
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
func (mathRatHelperMech *mathBigRatHelperMechanics) bigRatToNativeNumStr(
	bigRatNum *big.Rat,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	nativeNumStr string,
	err error) {

	if mathRatHelperMech.lock == nil {
		mathRatHelperMech.lock = new(sync.Mutex)
	}

	mathRatHelperMech.lock.Lock()

	defer mathRatHelperMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathBigRatHelperMechanics."+
			"bigRatToNativeNumStr()",
		"")

	if err != nil {

		return nativeNumStr, err
	}

	if bigRatNum == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'bigRatNum' is a nil pointer!\n",
			ePrefix.String())

		return nativeNumStr, err
	}

	if roundToFractionalDigits < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'roundToFractionalDigits' is invalid!\n"+
			"'roundToFractionalDigits' has a value less than zero (0).\n"+
			"roundToFractionalDigits = %v\n",
			ePrefix.String(),
			roundToFractionalDigits)

		return nativeNumStr, err

	}

	nativeNumStr =
		bigRatNum.FloatString(roundToFractionalDigits)

	return nativeNumStr, err
}

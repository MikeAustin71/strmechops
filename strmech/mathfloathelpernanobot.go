package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

// mathFloatHelperNanobot
//
// Provides helper methods for type MathFloatHelper
type mathFloatHelperNanobot struct {
	lock *sync.Mutex
}

// nativeNumStrToBigFloat
//
// Receives a Native Number String ('nativeNumStr') and
// then converts and returns that string as a big.Float
// floating point numeric value.
//
// The term 'Native' applies in the sense that the number
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
// The input parameter 'nativeNumStr' must be formatted
// as a Native Number String which is defined as follows:
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
//     If the input parameter 'nativeNumStr' does NOT meet these
//     criteria, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr					string
//
//		This string contains the Native Number String which
//		will be parsed to produce and return a big.Float
//		value.
//
//		If 'nativeNumStr' fails to meet the formatting
//		criteria for a Native Number String, an error
//		will be returned.
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
//	big.Float
//
//		If this method completes successfully, the pure
//		number string passed as input value 'nativeNumStr'
//		will be converted and returned as a big.Float
//		floating point value.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (mathFloatHelperNanobot *mathFloatHelperNanobot) nativeNumStrToBigFloat(
	nativeNumStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	big.Float,
	error) {

	if mathFloatHelperNanobot.lock == nil {
		mathFloatHelperNanobot.lock = new(sync.Mutex)
	}

	mathFloatHelperNanobot.lock.Lock()

	defer mathFloatHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"mathFloatHelperNanobot."+
			"nativeNumStrToBigFloat()",
		"")

	if err != nil {
		return big.Float{}, err
	}

	bigFloatNum := big.Float{}

	var ok bool

	_,
		ok = bigFloatNum.SetString(nativeNumStr)

	if !ok {

		err = fmt.Errorf("%v\n"+
			"Error: bigFloatNum.SetString(nativeNumStr)\n"+
			"Parsing 'nativeNumStr' failed!\n"+
			"nativeNumStr = '%v'\n",
			ePrefix.String(),
			nativeNumStr)
	}

	return bigFloatNum, err
}

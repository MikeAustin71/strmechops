package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

//	MathBigRatHelper
//
//	This type provides helper methods for Rational
//	Numbers created with the Go Programming Language
//	Big Math package.
//
//	The Big Math package defines type 'Rat' or Rational
//	Number as follows:
//
//	"A Rat represents a quotient a/b of arbitrary
//	precision. The zero value for a Rat represents the
//	value 0."
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://pkg.go.dev/math/big#Rat
type MathBigRatHelper struct {
	lock *sync.Mutex
}

// BigRatToNativeNumStr
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
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
func (mathRatHelper *MathBigRatHelper) BigRatToNativeNumStr(
	bigRatNum *big.Rat,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	nativeNumStr string,
	err error) {

	if mathRatHelper.lock == nil {
		mathRatHelper.lock = new(sync.Mutex)
	}

	mathRatHelper.lock.Lock()

	defer mathRatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathBigRatHelper."+
			"BigRatToNativeNumStr()",
		"")

	if err != nil {

		return nativeNumStr, err
	}

	return new(mathBigRatHelperMechanics).
		bigRatToNativeNumStr(
			bigRatNum,
			roundToFractionalDigits,
			ePrefix.XCpy("bigRatNum"))
}

//	BigRatToBigFloat
//
//	Converts a rational number (type big.Rat) to a
//	floating point number of type big.Float.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	bigRatNum					*big.Rat
//
//		The rational number which will be converted to a
//		big.Float numeric value.
//
//	roundToFractionalDigits 	int
//
//		Controls the number of fractional digits returned
//		as a big float number ('bigFloatNum').
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
//	bigFloatNum					*big.Float
//
//		If this method completes successfully, this
//		parameter will return a big.Float number
//		representing the numeric value of the rational
//		number passed through input parameter,
//		'bigRatNum'.
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
func (mathRatHelper *MathBigRatHelper) BigRatToBigFloat(
	bigRatNum *big.Rat,
	roundToFractionalDigits int,
	errorPrefix interface{}) (
	bigFloatNum *big.Float,
	err error) {

	if mathRatHelper.lock == nil {
		mathRatHelper.lock = new(sync.Mutex)
	}

	mathRatHelper.lock.Lock()

	defer mathRatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	bigFloatNum = big.NewFloat(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathBigRatHelper."+
			"BigRatToBigFloat()",
		"")

	if err != nil {
		return bigFloatNum, err
	}

	bigFloatNum,
		err = new(mathBigRatHelperAtom).
		ratToBigFloat(
			bigRatNum,
			roundToFractionalDigits,
			ePrefix)

	return bigFloatNum, err
}

// NativeNumStrToBigRatValue
//
// Receives a Native Number String and converts for
// storage in a *big.Rat number.
//
// Input parameter 'ptrBigRatNum' is a pointer to a
// big.Rat number. The numeric value converted from
// the Native Number String ('ptrBigRatNum') will be
// stored in the big.Rat number pointed to by the
// *big.Rat type 'ptrBigRatNum'.
//
// The term 'Native Number String' means that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math'
// and 'big' (big.Int, big.Float, big.Rat) routinely
// parse and convert this type of number string to
// numeric values. In addition, Native Number Strings are
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
//	Examples Of Native Number Strings
//		1000000
//		12.5483
//		-1000000
//		-12.5483
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The numeric value pointed to by the input parameter
//	'ptrBigRatNum' (Type *big.Rat) will be deleted and
//	overwritten with a new value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr				string
//
//		A Native Number String containing the numeric
//		character digits which will be converted to, and
//		stored in, the numeric value passed as input
//		parameter 'numericValue'.
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
//	ptrBigRatNum				*big.Rat
//
//		A pointer to a  big.Rat number. The numeric value
//		extracted from input parameter 'nativeNumStr' will
//		be stored in number pointed to by the big.Rat
//		pointer.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
//	error
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
func (mathRatHelper *MathBigRatHelper) NativeNumStrToBigRatValue(
	nativeNumStr string,
	ptrBigRatNum *big.Rat,
	errorPrefix interface{}) error {

	if mathRatHelper.lock == nil {
		mathRatHelper.lock = new(sync.Mutex)
	}

	mathRatHelper.lock.Lock()

	defer mathRatHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathBigRatHelper."+
			"NativeNumStrToBigRatValue()",
		"")

	if err != nil {
		return err
	}

	return new(mathBigRatHelperMechanics).
		nativeNumStrToBigRatValue(
			nativeNumStr,
			ptrBigRatNum,
			ePrefix.XCpy(
				"ptrBigRatNum<-nativeNumStr"))
}

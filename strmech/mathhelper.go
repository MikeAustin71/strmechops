package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// MathHelper
//
// Provides math utility methods
type MathHelper struct {
	lock *sync.Mutex
}

// NativeNumStrToNumericValue
//
// Receives a Native Number String and converts it to a
// numeric value passed as an empty interface through
// input parameter 'numericValue'.
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
// The 'numericValue' input parameter supports pointers
// to specific concrete types which will be configured
// with the numeric value extracted from the Native
// Number String ('nativeNumStr').
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If the Native Number String ('nativeNumStr') fails to
//	comply with Native Number String formatting
//	requirements try the following method as a means of
//	converting a 'dirty' number string to a valid Native
//	Number String:
//
//			NumStrHelper.DirtyToNativeNumStr()
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
//		The term 'Native' applies in the sense that the
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
//	numericValue				interface{}
//
//		The numeric value generated from input parameter
//		'nativeNumStr' will be stored in a numeric value
//		type passed through this empty interface.
//
//		The supported Numeric Value Types are listed as
//		follows:
//
//				*float32, *float64, *big.Float
//				*BigFloatDto
//				*TextFieldFormatDtoFloat64
//				*TextFieldFormatDtoBigFloat
//				*int8, *int16, *int, *int32, *int64, *big.Int
//				*uint8, *uint16, *uint, *uint32, *uint64
//				*NumberStrKernel
//
//		Any type passed through this empty interface which
//		is not listed above will generate an error.
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
func (mathHelper *MathHelper) NativeNumStrToNumericValue(
	nativeNumStr string,
	numericValue interface{},
	errorPrefix interface{}) error {

	if mathHelper.lock == nil {
		mathHelper.lock = new(sync.Mutex)
	}

	mathHelper.lock.Lock()

	defer mathHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathHelper."+
			"NativeNumStrToNumericValue()",
		"")

	if err != nil {

		return err
	}

	return new(mathHelperNanobot).
		nativeNumStrToNumericValue(
			nativeNumStr,
			numericValue,
			ePrefix.XCpy(
				"numericValue<-nativeNumStr"))
}

// NumericValueToNativeNumStr
//
// Receives a numeric value as an empty interface and
// converts that value to a Native Number String.
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
// # Input Parameters
//
//		numericValue				interface{}
//
//			An empty interface containing the numeric value
//			which will be converted and returned as a Native
//			Number String.
//
//			An error will be returned if the concrete type
//			passed through this parameter does not match one
//			of the supported types below.
//
//			Supported Numeric Value ('numericValue') Types:
//
//	     		float32, float64, big.Float
//				*float32, *float64, *big.Float
//				*BigFloatDto, BigFloatDto
//				int8, int16, int, int32, int64, big.Int
//				*int8, *int16, *int, *int32, *int64, *big.Int
//				uint8, uint16, uint, uint32, uint64
//				*uint8, *uint16, *uint, *uint32, *uint64
//				*TextFieldFormatDtoFloat64, TextFieldFormatDtoFloat64
//				*TextFieldFormatDtoBigFloat, TextFieldFormatDtoBigFloat
//				*NumberStrKernel, NumberStrKernel
//
//		errorPrefix					interface{}
//
//			This object encapsulates error prefix text which
//			is included in all returned error messages.
//			Usually, it contains the name of the calling
//			method or methods listed as a method or function
//			chain of execution.
//
//			If no error prefix information is needed, set this
//			parameter to 'nil'.
//
//			This empty interface must be convertible to one of
//			the following types:
//
//			1.	nil
//					A nil value is valid and generates an
//					empty collection of error prefix and
//					error context information.
//
//			2.	string
//					A string containing error prefix
//					information.
//
//			3.	[]string
//					A one-dimensional slice of strings
//					containing error prefix information.
//
//			4.	[][2]string
//					A two-dimensional slice of strings
//			   		containing error prefix and error
//			   		context information.
//
//			5.	ErrPrefixDto
//					An instance of ErrPrefixDto.
//					Information from this object will
//					be copied for use in error and
//					informational messages.
//
//			6.	*ErrPrefixDto
//					A pointer to an instance of
//					ErrPrefixDto. Information from
//					this object will be copied for use
//					in error and informational messages.
//
//			7.	IBasicErrorPrefix
//					An interface to a method
//					generating a two-dimensional slice
//					of strings containing error prefix
//					and error context information.
//
//			If parameter 'errorPrefix' is NOT convertible
//			to one of the valid types listed above, it will
//			be considered invalid and trigger the return of
//			an error.
//
//			Types ErrPrefixDto and IBasicErrorPrefix are
//			included in the 'errpref' software package:
//				"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
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
func (mathHelper *MathHelper) NumericValueToNativeNumStr(
	numericValue interface{},
	errorPrefix interface{}) (
	string,
	error) {

	if mathHelper.lock == nil {
		mathHelper.lock = new(sync.Mutex)
	}

	mathHelper.lock.Lock()

	defer mathHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathHelper."+
			"NumericValueToNativeNumStr()",
		"")

	if err != nil {

		return "", err
	}

	return new(mathHelperNanobot).
		numericValueToNativeNumStr(
			numericValue,
			ePrefix.XCpy("<-numericValue"))

}

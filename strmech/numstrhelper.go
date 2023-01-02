package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrHelper
//
// Provides number string utility methods.
type NumStrHelper struct {
	lock *sync.Mutex
}

// DirtyToNativeNumStr
//
// Converts a "dirty" or malformed string containing
// numeric digits to a well-formed and properly
// formatted Native Number String.
//
// The 'Dirty Number String' passed as input parameter
// 'dirtyNumberStr' is expected to comply with the
// following requirements:
//
//  1. The dirty number string must contain numeric
//     digit characters zero to nine inclusive (0-9).
//
//  2. The dirty number string must contain a radix
//     point or decimal separator to separate
//     integer and fractional digits in a floating
//     point numeric value. This decimal separator
//     is specified by input parameter,
//     'decimalSeparator'.
//
//     If no decimal separator is identified in the
//     dirty number string, the numeric value is
//     assumed to be an integer value.
//
//  3. The dirty number string must designate
//     negative numeric values using one of the
//     following three negative number symbols:
//
//     (a)	A Leading Minus Sign ('-').
//     Example: -123.45
//
//     (b)	A Trailing Minus Sign ('-').
//     Example: 123.45-
//
//     (c) A combination of leading and trailing
//     Parentheses ('()').
//     Example: (123.45)
//
//     Examples Of Dirty Number Strings
//
//     $1,254.65
//     1 000 000,00 €
//     1.000.000,00 €
//     6,78,90,00,00,00,00,000
//     6,7890,0000,0000,0000
//
// All the examples shown above are valid number string
// formats used by different countries and cultures.
// There is nothing wrong with these formats. The term
// "Dirty" simply distinguishes these formats from the
// Native Number String format required by many
// functions and packages in the Go Programming
// Language. The Native Number String format is one
// which is used as a standard format for numeric
// conversions performed in the Go Programming Language
// as well as many other programming languages.
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
//     (a)	Numeric digits zero through nine inclusive
//     (0-9).
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
//	dirtyNumberStr				string
//
//		This number string contains the numeric digits
//		which will be converted and returned as a Native
//		Number String.
//
//		The 'dirtyNumberStr' is expected to comply with
//		the following requirements:
//
//		1.	The dirty number string must contain numeric
//			digit characters zero to nine inclusive (0-9).
//
//		2.	The dirty number string must contain a radix
//			point or decimal separator to separate
//			integer and fractional digits in a floating
//			point numeric value. This decimal separator
//			is specified by input parameter,
//			'decimalSeparator'.
//
//			If no decimal separator is identified in the
//			dirty number string, the numeric value is
//			assumed to be an integer value.
//
//		3.	The dirty number string must designate
//			negative numeric values using one of the
//			following three negative number symbols:
//
//			(a)	A Leading Minus Sign ('-').
//				Example: -123.45
//
//			(b)	A Trailing Minus Sign ('-').
//				Example: 123.45-
//
//			(c) A combination of leading and trailing
//				Parentheses ('()').
//				Example: (123.45)
//
//		If 'dirtyNumberStr' does not contain any numeric
//		digits, an error will be returned.
//
//	decimalSeparator			string
//
//		Type DecimalSeparatorSpec is used to specify the
//		radix point or decimal separator which will
//		separate integer and fractional digits in the
//		dirty number string passed as input parameter
//		'dirtyNumberStr'.
//
//		The decimal separator will typically consist of
//		one or more non-numeric characters.
//
//		If 'decimalSeparator' consists of an empty
//		or zero length sting, it is assumed that the
//		numeric value contained in input parameter
//		'dirtyNumberStr' is an integer value.
//
//		In the US, Australia, UK, most of Canada and many
//		other countries the period ('.'), or decimal
//		point, separates integer and fractional digits
//		within a floating point numeric value.
//
//		Other countries, including many in the European
//		Union, use the comma (',') to separate integer
//		and fractional digits within a number string.
//
//		If 'decimalSeparator' contains any one of the
//		following invalid characters, an error will be
//		returned.
//
//			Invalid Decimal Separator Characters
//							'-'
//							'('
//							')'
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
//	nativeNumStr				string
//
//		If this method completes successfully, a Native
//		Number String will be returned.
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
func (numStrHelper *NumStrHelper) DirtyToNativeNumStr(
	dirtyNumberStr string,
	decimalSeparator string,
	errorPrefix interface{}) (
	nativeNumStr string,
	err error) {

	if numStrHelper.lock == nil {
		numStrHelper.lock = new(sync.Mutex)
	}

	numStrHelper.lock.Lock()

	defer numStrHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"MathHelper."+
			"DirtyToNativeNumStr()",
		"")

	if err != nil {

		return nativeNumStr, err
	}

	var decSeparator DecimalSeparatorSpec

	if len(decimalSeparator) > 0 {

		decSeparator,
			err = new(DecimalSeparatorSpec).
			NewStr(
				decimalSeparator,
				ePrefix.XCpy(
					"decSeparator<-decimalSeparator"))

		if err != nil {

			return nativeNumStr, err
		}

	}

	var dirtyNumRunes RuneArrayDto

	dirtyNumRunes,
		err = new(RuneArrayDto).
		NewString(
			dirtyNumberStr,
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				"dirtyNumRunes<-dirtyNumberStr"))

	if err != nil {

		return nativeNumStr, err
	}

	var nativeNumRunes RuneArrayDto

	nativeNumRunes,
		err = new(numStrHelperAtom).dirtyToNativeNumRunes(
		&dirtyNumRunes,
		"dirtyNumberStr",
		decSeparator,
		ePrefix.XCpy(
			"nativeNumRunes<-dirtyNumberStr"))

	if err != nil {

		return nativeNumStr, err
	}

	nativeNumStr = nativeNumRunes.GetCharacterString()

	return nativeNumStr, err
}

// IsValidNativeNumStr
//
// Receives a Native Number String and performs a
// diagnostic analysis to determine if string conforms
// to the specifications required for a properly
// formatted Native Number String.
//
// The term 'Native' means that the number string format
// is designed to interoperate with the Golang
// programming language library functions and packages.
// Types like 'strconv', 'strings', 'math' and 'big'
// (big.Int, big.Float, big.Rat) routinely parse and
// convert this type of number string to numeric values.
// In addition, Native Number Strings are frequently
// consumed by external library functions such as this
// one (String Mechanics 'strmech') to convert strings to
// numeric values and numeric values to strings.
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
// If the Native Number String fails to meet the criteria
// for a Native Number String, it is classified as
// invalid and a boolean value of 'false' will be
// returned.
//
// If the Native Number String is determined to valid in
// all respects, this method returns a boolean value of
// 'true'.
//
// This method is functionally equivalent to
// NumStrHelper.IsValidNativeNumStrError() with the sole
// exception being that this method returns a boolean
// value instead of an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr	string
//
//		A string of characters formatted as a Native
//		Number String.
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
//		This method will analyze the Native Number String
//		passed through input parameter 'nativeNumStr' to
//		determine if meets the required formatting
//		criteria and is valid in all respects.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the Native Number String passed as input
//		parameter 'nativeNumStr' fails the meet the
//		criteria for a valid Native Number String, this
//		method will return a boolean value of 'false'.
//
//		If parameter 'nativeNumStr' is determined to be
//		valid in all respects, this method returns a boolean
//		value of 'true'.
func (numStrHelper *NumStrHelper) IsValidNativeNumStr(
	nativeNumStr string) bool {

	if numStrHelper.lock == nil {
		numStrHelper.lock = new(sync.Mutex)
	}

	numStrHelper.lock.Lock()

	defer numStrHelper.lock.Unlock()

	isValid,
		_ := new(numStrHelperQuark).
		testValidityOfNativeNumStr(
			nativeNumStr,
			"nativeNumStr",
			nil)

	return isValid
}

// IsValidNativeNumStrError
//
// Receives a Native Number String and performs a
// diagnostic analysis to determine if string conforms
// to the specifications required for a properly
// formatted Native Number String.
//
// The term 'Native' means that the number string format
// is designed to interoperate with the Golang
// programming language library functions and packages.
// Types like 'strconv', 'strings', 'math' and 'big'
// (big.Int, big.Float, big.Rat) routinely parse and
// convert this type of number string to numeric values.
// In addition, Native Number Strings are frequently
// consumed by external library functions such as this
// one (String Mechanics 'strmech') to convert strings to
// numeric values and numeric values to strings.
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
// If the Native Number String fails to meet the criteria
// for a properly formatted Native Number String, it is
// classified as invalid and an error will be returned
// containing an appropriate error message.
//
// If the Native Number String is determined to be valid
// in all respects, this method returns an error value of
// 'nil'.
//
// This method is functionally equivalent to
// NumStrHelper.IsValidNativeNumStr() with the sole
// exception being that this method returns an error
// instead of a boolean value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nativeNumStr	string
//
//		A string of characters formatted as a Native
//		Number String.
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
//		This method will analyze the Native Number String
//		passed through input parameter 'nativeNumStr' to
//		determine if meets the required formatting
//		criteria and is valid in all respects.
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
//		If the Native Number String passed as input
//		parameter 'nativeNumStr' fails the meet the
//		criteria for a valid Native Number String, this
//		method will return an error containing an
//		appropriate error message.
//
//		If parameter 'nativeNumStr' is determined to be
//		valid in all respects, this method returns an
//		error value of 'nil'.
//
//		If an error is returned, that error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (numStrHelper *NumStrHelper) IsValidNativeNumStrError(
	nativeNumStr string,
	errorPrefix interface{}) error {

	if numStrHelper.lock == nil {
		numStrHelper.lock = new(sync.Mutex)
	}

	numStrHelper.lock.Lock()

	defer numStrHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrHelper."+
			"IsValidNativeNumStrError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(numStrHelperQuark).
		testValidityOfNativeNumStr(
			nativeNumStr,
			"nativeNumStr",
			ePrefix.XCpy(
				"nativeNumStr"))

	return err
}

// NormalizeNativeNumStr
//
// Removes leading integer zeros and trailing fractional
// zeros from a Native Number String and returns a clean
// or 'normalized' version of the Native Number String.
//
// The term 'Native Number String' means that the number
// string format is designed to interoperate with the
// Golang programming language library functions and
// packages. Types like 'strconv', 'strings', 'math' and
// 'big' (big.Int, big.Float, big.Rat) routinely parse
// and convert this type of number string to numeric
// values. In addition, Native Number Strings are
// frequently consumed by external library functions such
// as this one (String Mechanics 'strmech') in order to
// convert strings to numeric values and numeric values
// to strings.
//
// A Native Number String which has been normalized,
// will contain no leading integer zeros and no trailing
// fractional zeros.
//
//	Examples:
//		Un-Rationalized Native Number Strings:
//					0001234
//					1234.56780000
//					0001234.5678000
//
//		Rationalized Native Number String:
//					1234.5678
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nonStandardNativeNumStr		string
//
//		A raw Native Number String which may contain
//		leading integer zeros and or trailing fractional
//		zeros.
//			Examples:
//				Non-Standard Native Number Strings:
//						0001234
//						1234.56780000
//						0001234.5678000
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
//		This method will analyze input parameter
//		'unRationalizedNativeNumStr' and return a
//		clean or 'normalized' version of the Native
//		Number String by deleting all leading integer
//		zeros and trailing fractional zeros.
//
//		If 'nonStandardNativeNumStr' fails to meet the
//		formatting criteria for a Native Number String,
//		an error will be returned.
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
//	errPrefDto						*ePref.ErrPrefixDto
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
//	normalizedNativeNumStr		string
//
//		If this method completes successfully, a
//		normalized Native Number String extracted
//		from input parameter 'nonStandardNativeNumStr'
//		will be returned.
//
//		A valid 'normalized' Native Number String will
//		meet the following criteria:
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
//		6.	A 'Normalized' Native Number String will
//			contain no leading integer zeros.
//
//										 NOT THIS: 0001234
//				Rationalized Native Number String: 1234
//
//		7.	A 'Normalized' Native Number String will
//			contain no trailing fractional zeros.
//
//								   NOT THIS: 12.34000
//			Normalized Native Number String: 12.34
//
//	normalizedNativeNumStrStats	NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the integer and fractional digits contained
//		in the return parameter
//		'normalizedNativeNumStr'.
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
//	err								error
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
func (numStrHelper *NumStrHelper) NormalizeNativeNumStr(
	nonStandardNativeNumStr string,
	errorPrefix interface{}) (
	normalizedNativeNumStr string,
	normalizedNativeNumStrStats NumberStrStatsDto,
	err error) {

	if numStrHelper.lock == nil {
		numStrHelper.lock = new(sync.Mutex)
	}

	numStrHelper.lock.Lock()

	defer numStrHelper.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrHelper."+
			"NormalizeNativeNumStr()",
		"")

	if err != nil {

		return normalizedNativeNumStr,
			normalizedNativeNumStrStats,
			err
	}

	normalizedNativeNumStr,
		normalizedNativeNumStrStats,
		err = new(numStrHelperElectron).
		normalizeNativeNumStr(
			nonStandardNativeNumStr,
			ePrefix.XCpy(
				"nonStandardNativeNumStr"))

	return normalizedNativeNumStr,
		normalizedNativeNumStrStats,
		err
}

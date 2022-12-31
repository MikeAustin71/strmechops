package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrHelperQuark
//
// Provides number string utility methods
type numStrHelperAtom struct {
	lock *sync.Mutex
}

// dirtyToNativeNumRunes
//
// Converts a string containing numeric digits to a well
// formatted Native Number String.
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
//	dirtyNumberRunes			RuneArrayDto
//
//		This instance of RuneArrayDto encapsulates a rune
//		array containing the numeric digits which will be
//		converted and returned as a Native Number rune
//		array.
//
//		The 'dirtyNumberRunes' rune array is expected to
//		comply with the following requirements:
//
//		1.	The dirty number rune array must contain
//			numeric digit characters zero to nine
//			inclusive (0-9).
//
//		2.	The dirty number rune array must contain a
//			radix point or decimal separator to separate
//			integer and fractional digits in a floating
//			point numeric value. This decimal separator
//			is specified by input parameter,
//			'decimalSeparator'.
//
//			If no decimal separator is identified in the
//			dirty number string, the numeric value is
//			assumed to be an integer value.
//
//		4.	The dirty number string must designate
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
//	decimalSeparator			DecimalSeparatorSpec
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
//		or zero length rune array, it is assumed that
//		the numeric value contained in input parameter
//		'dirtyNumberRunes' is an integer value.
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
func (nStrHelperAtom *numStrHelperAtom) dirtyToNativeNumRunes(
	dirtyNumberRunes *RuneArrayDto,
	dirtyNumberRunesLabel string,
	decimalSeparator DecimalSeparatorSpec,
	ePrefDto *ePref.ErrPrefixDto) (
	nativeNumStr RuneArrayDto,
	err error) {

	if nStrHelperAtom.lock == nil {
		nStrHelperAtom.lock = new(sync.Mutex)
	}

	nStrHelperAtom.lock.Lock()

	defer nStrHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"numStrHelperAtom."+
			"dirtyToNativeNumRunes()",
		"")

	if err != nil {

		return nativeNumStr, err
	}

	if dirtyNumberRunes == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			dirtyNumberRunesLabel)

		return nativeNumStr, err
	}

	lenOfDirtyNumRunes := len(dirtyNumberRunes.CharsArray)

	if lenOfDirtyNumRunes == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is empty\n"+
			"with a length of zero!\n",
			ePrefix.String(),
			dirtyNumberRunesLabel)

		return nativeNumStr, err
	}

	var decSepChars *RuneArrayDto

	decSepChars = &decimalSeparator.decimalSeparatorChars

	decSepIsNOP := false

	lenOfDecSeps := len(decSepChars.CharsArray)

	if lenOfDecSeps == 0 {

		decSepIsNOP = true

	}

	for i := 0; i < lenOfDecSeps; i++ {

		if decSepChars.CharsArray[i] == '-' ||
			decSepChars.CharsArray[i] == '(' ||
			decSepChars.CharsArray[i] == ')' {

			err = fmt.Errorf("%v\n"+
				"ERROR: Input parameter Decimal Separators is invalid.\n"+
				"Decimal Separators contains an invalid character!\n"+
				"The following three characters are invalid\n"+
				"decimal separator characters: '-','(', ')'\n"+
				"The Decimal Separators parameter contains\n"+
				"invalid character '%v'\n",
				ePrefix.String(),
				string(decSepChars.CharsArray[i]))

			return nativeNumStr, err
		}
	}

	var foundNegNumSymbol, foundLeadingParen,
		foundFirstNumericDigit, foundDecimalSep bool

	numOfNumericDigits := 0

	for i := 0; i < lenOfDirtyNumRunes; i++ {

		if dirtyNumberRunes.CharsArray[i] >= '0' &&
			dirtyNumberRunes.CharsArray[i] <= '9' {

			nativeNumStr.CharsArray =
				append(
					nativeNumStr.CharsArray,
					dirtyNumberRunes.CharsArray[i])

			foundFirstNumericDigit = true

			numOfNumericDigits++

			continue
		}

		if foundNegNumSymbol == false &&
			dirtyNumberRunes.CharsArray[i] == '-' {

			if foundFirstNumericDigit == true {

				if !foundNegNumSymbol {
					foundNegNumSymbol = true
				}

			} else {
				// MUST BE -
				// Have NOT found first numeric digit yet

				if !foundNegNumSymbol {
					foundNegNumSymbol = true
				}

			}

			continue
		}

		if foundNegNumSymbol == false &&
			dirtyNumberRunes.CharsArray[i] == '(' {

			if foundFirstNumericDigit == true {

				continue

			} else {
				// MUST BE -
				// Have NOT found first numeric digit yet

				foundLeadingParen = true

			}

		}

		if foundNegNumSymbol == false &&
			foundLeadingParen == true &&
			dirtyNumberRunes.CharsArray[i] == ')' {

			if foundFirstNumericDigit == true {

				if foundLeadingParen == true {
					foundNegNumSymbol = true
				}

			} else {
				// MUST BE -
				// Have NOT found first numeric digit yet

				continue

			}

		}

		if decSepIsNOP == false &&
			foundDecimalSep == false &&
			dirtyNumberRunes.CharsArray[i] ==
				decSepChars.CharsArray[0] {

			if i+lenOfDecSeps > lenOfDirtyNumRunes {

				continue

			}

			j := i

			for k := 0; k < lenOfDecSeps; k++ {

				if decSepChars.CharsArray[k] !=
					dirtyNumberRunes.CharsArray[j] {

					continue
				}

				j++

			}

			// Found Decimal Separator
			nativeNumStr.CharsArray =
				append(
					nativeNumStr.CharsArray,
					'.')

			foundDecimalSep = true

			i = i + lenOfDecSeps - 1

		}

	}

	if numOfNumericDigits == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' contains zero numeric digits!\n"+
			"%v must contain at least one numeric character digit (0-9).\n",
			ePrefix.String(),
			dirtyNumberRunesLabel,
			dirtyNumberRunesLabel)

		return nativeNumStr, err
	}

	if foundNegNumSymbol {
		nativeNumStr.CharsArray = append(
			[]rune{'-'},
			nativeNumStr.CharsArray...)
	}

	return nativeNumStr, err
}

//	parsePureNumStr
//
//	Receives a pure number string and proceeds to return the
//	extracted numeric value as a type NumberStrKernel.
//
//	This method is particularly useful when numeric values
//	are converted to string using 'fmt.Sprintf()' and
//	similar formatting algorithms.
//
//	A "Pure Number String" is defined as follows:
//
//		1.	Consists of numeric character digits
//			zero through nine inclusive (0-9).
//
//		2.	Option: A Pure Number String may include
//			a radix point or decimal separator. The
//			decimal separator may consist of one or
//			more characters.
//
//			In the US, UK, Australia and most of Canada,
//			the decimal separator is the period
//			character ('.') known as the decimal point.
//
//		3.	Optional: A Pure Number String may include a
//			negative number sign symbol consisting of a
//			minus sign ('-'). Only the minus sign ('-')
//			classifies the numeric value as a negative
//			number in Pure Number String.
//
//			If the leading or trailing minus sign ('-')
//			is NOT present, the numeric value is assumed
//			to be positive.
//
//		4.	Only numeric characters, the decimal
//			separator and the minus sign will be
//			processed by the number string parsing
//			algorithm. All other characters will be
//			ignored.
//
//		5.	Pure Number Strings consist of a single
//			numeric value. The entire Pure Number String
//			will be parsed, or processed, and only one
//			numeric value per Pure Number String will
//			be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberString			RuneArrayDto
//
//		This rune array contains the character digits from
//		which the numeric value will be extracted and
//		returned as a NumberStrKernel.
//
//		A "Pure Number String" is defined as follows:
//
//			1.	Consists of numeric character digits
//				zero through nine inclusive (0-9).
//
//			2.	Optional: A Pure Number String may
//				include a radix point or decimal
//				separator. Decimal separators separate
//				integer and fractional numeric digits in
//				a pure number string. The decimal
//				separator may consist of one or more text
//				characters.
//
//				In the US, UK, Australia, most of Canada
//				and many other countries, the decimal
//				separator is the period character ('.')
//				known as the decimal point.
//
//				In France, Germany and many countries in
//				the European Union, the Decimal Separator
//				is the comma character (',').
//
//			3.	Optional: A Pure Number String may
//				include a negative number sign symbol
//				consisting of a minus sign ('-'). The
//				minus sign will identify the numeric
//				value contained in the pure number string
//				as a negative number. Only the minus sign
//				('-') classifies a numeric value as a
//				negative number in a Pure Number String.
//
//				If a leading or trailing minus sign ('-')
//				is NOT present in the pure number string,
//				the numeric value is assumed to be
//				positive.
//
//			4.	Only numeric characters, the decimal
//				separator and the minus sign will be
//				processed by the pure number string
//				parsing algorithm. All other characters
//				will be	ignored.
//
//			5.	Pure Number Strings consist of a single
//				numeric value. The entire Pure Number String
//				will be parsed, or processed, and only one
//				numeric value per Pure Number String will
//				be returned.
//
//	decSeparatorSpec				DecimalSeparatorSpec
//
//		This structure contains the radix point or
//		decimal separator character(s) which will be used
//		to separate integer and fractional digits within
//		a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingMinusSign			bool
//
//		In pure number strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		pure number string parsing algorithm will search
//		for a leading minus sign ('-') at the beginning of
//		the number string. Leading minus signs represent
//		the standard means for designating negative
//		numeric values in the US, UK, Australia, most of
//		Canada and many other parts of world.
//
//		Example Leading Minus Sign:
//			"-123.456" or "- 123.456"
//
//		When 'leadingMinusSign' is set to 'false', the
//		pure number string parsing algorithm will search
//		for trailing minus signs ('-') located at the end
//		of the number string. Trailing minus signs
//		represent the standard for France, Germany and
//		many countries in the European Union.
//
//		NOTE: Identification of a trailing minus sign in
//		the pure number string input parameter,
//		'pureNumberString', will immediately terminate
//		the search for numeric characters.
//
//		Example Trailing Number Symbols:
//			"123.456-" or "123.456 -"
//
//	errPrefDto          		*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	numStrKernel				NumberStrKernel
//
//		If this method completes successfully, an instance of
//		NumberStrKernel containing the numeric value extracted
//		from parameter 'pureNumberString' will be returned.
//
//	err							error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error Type
//		will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (nStrHelperAtom *numStrHelperAtom) parsePureNumStr(
	pureNumberString RuneArrayDto,
	decSeparatorSpec DecimalSeparatorSpec,
	leadingMinusSign bool,
	ePrefDto *ePref.ErrPrefixDto) (
	numStrKernel NumberStrKernel,
	err error) {

	if nStrHelperAtom.lock == nil {
		nStrHelperAtom.lock = new(sync.Mutex)
	}

	nStrHelperAtom.lock.Lock()

	defer nStrHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		ePrefDto,
		"numStrHelperAtom."+
			"parsePureNumStr()",
		"")

	if err != nil {

		return numStrKernel, err
	}

	lenPureNStr := pureNumberString.GetRuneArrayLength()

	if lenPureNStr == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'pureNumberString' is empty\n"+
			"and contains zero characters.\n",
			ePrefix.String())

		return numStrKernel, err
	}

	var numberStats NumberStrStatsDto

	numberStats,
		err = new(numStrMathQuark).
		pureNumStrToRunes(
			string(pureNumberString.CharsArray),
			&numStrKernel.integerDigits,
			&numStrKernel.fractionalDigits,
			&decSeparatorSpec.decimalSeparatorChars,
			leadingMinusSign,
			ePrefix)

	if err != nil {

		return numStrKernel, err
	}

	numStrKernel.numberSign = numberStats.NumberSign

	numStrKernel.numberValueType = numberStats.NumberValueType

	numStrKernel.isNonZeroValue = !numberStats.IsZeroValue

	return numStrKernel, err
}

// rationalizeNativeNumStr
//
// Removes leading integer zeros and trailing fractional
// zeros from a Native Number String and returns a clean
// or 'rationalized' version of the Native Number String.
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
// A Native Number String which has been rationalized,
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
//	unRationalizedNativeNumStr		string
//
//		A Native Number String which may contain leading
//		integer zeros and or trailing fractional zeros.
//			Examples:
//				Un-Rationalized Native Number Strings:
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
//		clean or 'rationalized' version of the Native
//		Number String by deleting all leading integer
//		zeros and trailing fractional zeros.
//
//		If 'unRationalizedNativeNumStr' fails to meet the
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
//	rationalizedNativeNumStr		string
//
//		If this method completes successfully, a
//		rationalized Native Number String extracted
//		from input parameter 'unRationalizedNativeNumStr'
//		will be returned.
//
//		A valid 'rationalized' Native Number String will
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
//		6.	A 'Rationalized' Native Number String will
//			contain no leading integer zeros.
//
//										 NOT THIS: 0001234
//				Rationalized Native Number String: 1234
//
//		7.	A 'Rationalized' Native Number String will
//			contain no trailing fractional zeros.
//
//										 NOT THIS: 12.34000
//				Rationalized Native Number String: 12.34
//
//	rationalizedNativeNumStrStats	NumberStrStatsDto
//
//		This data transfer object will return critical
//		statistics on the numeric value represented
//		by the integer and fractional digits contained
//		in the return parameter
//		'rationalizedNativeNumStr'.
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
func (nStrHelperAtom *numStrHelperAtom) rationalizeNativeNumStr(
	unRationalizedNativeNumStr string,
	errPrefDto *ePref.ErrPrefixDto) (
	rationalizedNativeNumStr string,
	rationalizedNativeNumStrStats NumberStrStatsDto,
	err error) {

	if nStrHelperAtom.lock == nil {
		nStrHelperAtom.lock = new(sync.Mutex)
	}

	nStrHelperAtom.lock.Lock()

	defer nStrHelperAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrHelperAtom."+
			"rationalizeNativeNumStr()",
		"")

	if err != nil {

		return rationalizedNativeNumStr,
			rationalizedNativeNumStrStats,
			err
	}

	_,
		err = new(numStrHelperQuark).
		testValidityOfNativeNumStr(
			unRationalizedNativeNumStr,
			"unRationalizedNativeNumStr",
			ePrefix.XCpy(
				"unRationalizedNativeNumStr"))

	if err != nil {

		return rationalizedNativeNumStr,
			rationalizedNativeNumStrStats,
			err
	}

	var integerDigits RuneArrayDto
	var fractionalDigits RuneArrayDto

	rationalizedNativeNumStrStats,
		err = new(numStrMathQuark).
		nativeNumStrToRunes(
			unRationalizedNativeNumStr,
			&integerDigits,
			&fractionalDigits,
			ePrefix.XCpy(
				"unRationalizedNativeNumStr"))

	if err != nil {

		return rationalizedNativeNumStr,
			rationalizedNativeNumStrStats,
			err
	}

	reformatStr := false

	if rationalizedNativeNumStrStats.NumOfIntegerDigits !=
		rationalizedNativeNumStrStats.NumOfSignificantIntegerDigits {

		reformatStr = true

		var numOfLeadingZerosCount uint64

		numOfLeadingZerosCount =
			integerDigits.GetCountLeadingZeros()

		err = integerDigits.DeleteLeadingTrailingChars(
			numOfLeadingZerosCount,
			false,
			ePrefix.XCpy(
				"integerDigits<-numOfLeadingZerosCount"))

		if err != nil {

			return rationalizedNativeNumStr,
				rationalizedNativeNumStrStats,
				err
		}

	}

	if rationalizedNativeNumStrStats.NumOfFractionalDigits !=
		rationalizedNativeNumStrStats.NumOfSignificantFractionalDigits {

		reformatStr = true

		var numOfTrailingZerosCount uint64

		numOfTrailingZerosCount =
			fractionalDigits.GetCountTrailingZeros()

		err = fractionalDigits.DeleteLeadingTrailingChars(
			numOfTrailingZerosCount,
			true,
			ePrefix.XCpy(
				"fractionalDigits<-numOfTrailingZerosCount"))

		if err != nil {

			return rationalizedNativeNumStr,
				rationalizedNativeNumStrStats,
				err
		}

	}

	if reformatStr == false {

		rationalizedNativeNumStr = unRationalizedNativeNumStr

		return rationalizedNativeNumStr,
			rationalizedNativeNumStrStats,
			err

	}

	// MUST BE -
	// reformatStr == true

	if rationalizedNativeNumStrStats.NumberSign ==
		NumSignVal.Negative() {

		rationalizedNativeNumStr = "-"
	}

	if len(integerDigits.CharsArray) == 0 {

		rationalizedNativeNumStr += "0"

	} else {

		rationalizedNativeNumStr +=
			string(integerDigits.CharsArray)

	}

	if len(fractionalDigits.CharsArray) > 0 {

		rationalizedNativeNumStr += "."

		rationalizedNativeNumStr +=
			string(fractionalDigits.CharsArray)

	}

	rationalizedNativeNumStrStats,
		err = new(numStrMathQuark).
		nativeNumStrToRunes(
			rationalizedNativeNumStr,
			&integerDigits,
			&fractionalDigits,
			ePrefix.XCpy(
				"rationalizedNativeNumStr"))

	return rationalizedNativeNumStr,
		rationalizedNativeNumStrStats,
		err
}

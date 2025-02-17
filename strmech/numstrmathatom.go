package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type numStrMathAtom struct {
	lock *sync.Mutex
}

// addRunes - Adds two rune arrays of numeric digits and
// returns the total in a RuneArrayDto instance passed as
// an input parameter.
//
// The computed total of these two rune arrays will be
// expressed as numeric characters returned in a 'total'
// rune array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	addArray01					*RuneArrayDto
//		A pointer to an instance of RuneArrayDto. This
//		instance contains a rune array of numeric digits
//		which will be added to those contained in input
//		parameter, 'addArray02'.
//
//	addArray02					*RuneArrayDto
//		A pointer to an instance of RuneArrayDto. This
//		instance contains a rune array of numeric digits
//		which will be added to those contained in input
//		parameter, 'addArray01'.
//
//	totalArray 					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance will be populated with the results of
//		the addition operation performed on input
//		parameters 'addArray01' and 'addArray02'.
//
//	extendTotalForCarry			bool
//
//		If this parameter is set to 'true' it signals
//		that carry values from the last add operation
//		will cause the total array to be extended by
//		one array element and populated with a value
//		of one (1).
//
//		If this parameter is set to 'false', a carry
//		value of one (1) generated by the last addition
//		operation will NOT be added to the total array.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which
//		is included in all returned error messages. Usually,
//		it contains the name of the calling method or methods
//		listed as a function chain.
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
//	isCarry 					bool
//
//		If this return parameter is set to 'true', it signals
//		that the last addition operation produced a 'carry'
//		value of '1' which has NOT been added to the total
//		array.
//
//	err							error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during	processing, the returned error
//		Type will encapsulate an error	message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (nStrMathAtom *numStrMathAtom) addRunes(
	addArray01 *RuneArrayDto,
	addArray02 *RuneArrayDto,
	totalArray *RuneArrayDto,
	extendTotalForCarry bool,
	errPrefDto *ePref.ErrPrefixDto) (
	isCarry bool,
	err error) {

	if nStrMathAtom.lock == nil {
		nStrMathAtom.lock = new(sync.Mutex)
	}

	nStrMathAtom.lock.Lock()

	defer nStrMathAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathAtom."+
			"addRunes()",
		"")

	if err != nil {
		return isCarry, err
	}

	if addArray01 == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'addArray01' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return isCarry, err
	}

	if addArray02 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'addArray02' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return isCarry, err
	}

	if totalArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'totalArray' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return isCarry, err
	}

	idxAdd1 := addArray01.GetRuneArrayLength() - 1

	if idxAdd1 < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'addArray01' is "+
			"a zero length rune array!\n",
			ePrefix.String())

		return isCarry, err
	}

	idxAdd2 := addArray02.GetRuneArrayLength() - 1

	if idxAdd2 < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'addArray02' is "+
			"a zero length rune array!\n",
			ePrefix.String())

		return isCarry, err
	}

	maxLastIdx := idxAdd1

	if idxAdd2 > maxLastIdx {
		maxLastIdx = idxAdd2
	}

	var num1, num2, factor rune

	isCarry = false

	total := RuneArrayDto{}

	total.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	total.CharsArray = make([]rune, maxLastIdx+1)

	for i := maxLastIdx; i > -1; i-- {

		if idxAdd1 < 0 {
			num1 = '0'
		} else {
			num1 = addArray01.CharsArray[idxAdd1]
		}

		idxAdd1--

		if idxAdd2 < 0 {
			num2 = '0'
		} else {
			num2 = addArray02.CharsArray[idxAdd2]
		}

		idxAdd2--

		if isCarry {
			factor = 47
			isCarry = false
		} else {
			factor = 48
		}

		total.CharsArray[i] = num1 + num2 - factor

		if total.CharsArray[i] > '9' {
			total.CharsArray[i] -= 10
			isCarry = true
		}

	}

	if isCarry && extendTotalForCarry {

		total.CharsArray = append(
			[]rune{'1'},
			total.CharsArray...)

		isCarry = false

	}

	err = totalArray.CopyIn(
		&total,
		ePrefix.XCpy(
			"totalArray<-total"))

	return isCarry, err
}

// addOneToRunes - Adds a value of one (1) to a series
// of numeric digits contained in a rune array passed
// as an input parameter.
//
// The computed total of these rune array of numeric
// digits plus one (1) will be expressed as numeric
// characters returned in a 'total' rune array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	addArray					*RuneArrayDto
//		A pointer to an instance of RuneArrayDto. This
//		instance contains a rune array of numeric digits
//		which will be added to a numeric value of one (1).
//
//	totalArray 					*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. This
//		instance will be populated with the results of
//		the addition operation performed by adding a
//		value of one (1) to input parameter, 'addArray'.
//
//	extendTotalForCarry			bool
//
//		If this parameter is set to 'true' it signals
//		that carry values from the last add operation
//		will cause the total array to be extended by
//		one array element and populated with a value
//		of one (1).
//
//		If this parameter is set to 'false', a carry
//		value of one (1) generated by the last addition
//		operation will NOT be added to the total array.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which
//		is included in all returned error messages. Usually,
//		it contains the name of the calling method or methods
//		listed as a function chain.
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
//	isCarry 					bool
//
//		If this return parameter is set to 'true', it signals
//		that the last addition operation produced a 'carry'
//		value of '1' which has NOT been added to the total
//		array.
//
//	err							error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during	processing, the returned error
//		Type will encapsulate an error	message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (nStrMathAtom *numStrMathAtom) addOneToRunes(
	addArray *RuneArrayDto,
	totalArray *RuneArrayDto,
	extendTotalForCarry bool,
	errPrefDto *ePref.ErrPrefixDto) (
	isCarry bool,
	err error) {

	if nStrMathAtom.lock == nil {
		nStrMathAtom.lock = new(sync.Mutex)
	}

	nStrMathAtom.lock.Lock()

	defer nStrMathAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathAtom."+
			"addOneToRunes()",
		"")

	if err != nil {
		return isCarry, err
	}

	if addArray == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'addArray' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return isCarry, err
	}

	if totalArray == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'totalArray' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return isCarry, err
	}

	idxAdd1 := addArray.GetRuneArrayLength() - 1

	if idxAdd1 < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'addArray' is "+
			"a zero length rune array!\n",
			ePrefix.String())

		return isCarry, err
	}

	maxLastIdx := idxAdd1

	var num1, num2, factor rune

	isCarry = false

	total := RuneArrayDto{}

	total.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	total.CharsArray = make([]rune, maxLastIdx+1)

	for i := maxLastIdx; i > -1; i-- {

		num1 = addArray.CharsArray[idxAdd1]

		if idxAdd1 == maxLastIdx {
			num2 = '1'
		} else {
			num2 = '0'
		}

		idxAdd1--

		if isCarry {
			factor = 47
			isCarry = false
		} else {
			factor = 48
		}

		total.CharsArray[i] = num1 + num2 - factor

		if total.CharsArray[i] > '9' {
			total.CharsArray[i] -= 10
			isCarry = true
		}

	}

	if isCarry && extendTotalForCarry {

		total.CharsArray = append(
			[]rune{'1'},
			total.CharsArray...)

		isCarry = false

	}

	err = totalArray.CopyIn(
		&total,
		ePrefix.XCpy(
			"totalArray<-total"))

	return isCarry, err
}

//	pureNumStrStats
//
//	Receives and analyzes the numeric digits configured
//	in a Pure Number String. The results of this
//	analysis, including number sign, number type and key
//	statistics relating to the numeric value contained in
//	the Pure Number String will be returned by an
//	instance of type, NumberStrStatsDto.
//
//	A Pure Number String consists of numeric digits, an
//	optional radix point (a.k.a. decimal point) and an
//	optional leading or trailing minus ('-') sign.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberStr				string
//
//		This Pure Number String contains the numeric
//		character digits which will be analyzed and
//		reported in the returned instance of
//		PureNumberStrComponents.
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
//				a Pure Number String. The decimal
//				separator may consist of one or more text
//				characters.
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
//				processed by the Pure Number String
//				parsing algorithm. All other characters
//				will be	ignored.
//
//			5.	Pure Number Strings consist of a single
//				numeric value. The entire Pure Number String
//				will be parsed, or processed, and only one
//				numeric value per Pure Number String will
//				be returned.
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol.
//
//		The Decimal Separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted, floating
//		point Number String.
//
//		In the US, UK, Australia, most of Canada and many
//		other countries, the Decimal Separator is the
//		period character ('.') known as the decimal
//		point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingMinusSign			bool
//
//		In Pure Number Strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		Pure Number String parsing algorithm will search
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
//		Pure Number String parsing algorithm will search
//		for trailing minus signs ('-') located at the end
//		of the number string. Trailing minus signs
//		represent the standard for France, Germany and
//		many countries in the European Union.
//
//		NOTE: Identification of a trailing minus sign in
//		the Pure Number String input parameter,
//		'pureNumberString', will immediately terminate
//		the search for additional numeric characters.
//
//		Example Trailing Number Symbols:
//			"123.456-" or "123.456 -"
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
//	numStrStats				NumberStrStatsDto
//
//		If this method completes successfully, this
//		parameter will return an instance of
//		NumberStrStatsDto. This data structure contains
//		an analysis and detail information on the Pure
//		Number String passed as input paramter,
//		'pureNumberStr'.
//
//		type NumberStrStatsDto struct {
//
//			NumOfIntegerDigits					uint64
//
//				The total number of integer digits to the
//				left of the radix point or, decimal point, in
//				the subject numeric value.
//
//			NumOfSignificantIntegerDigits		uint64
//
//				The number of integer digits to the left of
//				the radix point, excluding leading zeros, in
//				the subject numeric value.
//
//			NumOfFractionalDigits				uint64
//
//				The total number of fractional digits to the
//				right of the radix point or, decimal point,
//				in the subject numeric value.
//
//			NumOfSignificantFractionalDigits	uint64
//
//				The number of nonzero fractional digits to
//				the right of the radix point or, decimal
//				point, in the subject numeric value.
//
//			NumberValueType 					NumericValueType
//
//				This enumeration value specifies whether the
//				subject numeric value is classified either as
//				an integer or a floating point number.
//
//				Possible enumeration values are listed as
//				follows:
//					NumValType.None()
//					NumValType.FloatingPoint()
//					NumValType.Integer()
//
//			NumberSign							NumericSignValueType
//
//				An enumeration specifying the number sign
//				associated with the numeric value. Possible
//				values are listed as follows:
//					NumSignVal.None()		= Invalid Value
//					NumSignVal.Negative()	= -1
//					NumSignVal.Zero()		=  0
//					NumSignVal.Positive()	=  1
//
//			IsZeroValue							bool
//
//				If 'true', the subject numeric value is equal
//				to zero ('0').
//
//				If 'false', the subject numeric value is
//				greater than or less than zero ('0').
//		}
//
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrMathAtom *numStrMathAtom) pureNumStrStats(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	errPrefDto *ePref.ErrPrefixDto) (
	numStrStats NumberStrStatsDto,
	err error) {

	if nStrMathAtom.lock == nil {
		nStrMathAtom.lock = new(sync.Mutex)
	}

	nStrMathAtom.lock.Lock()

	defer nStrMathAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathAtom."+
			"pureNumStrStats()",
		"")

	if err != nil {
		return numStrStats, err
	}

	if len(pureNumberStr) == 0 {
		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'pureNumberStr'\n"+
			"is a zero length string and INVALID!\n",
			ePrefix.String())

		return numStrStats, err
	}

	var intDigits,
		fracDigits RuneArrayDto

	var radixPoint RuneArrayDto

	radixPoint,
		err = new(RuneArrayDto).NewString(
		decSeparatorChars,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix)

	if err != nil {
		return numStrStats, err
	}

	numStrStats,
		err = new(numStrMathQuark).
		pureNumStrToRunes(
			pureNumberStr,
			&intDigits,
			&fracDigits,
			&radixPoint,
			leadingMinusSign,
			ePrefix)

	return numStrStats, err
}

//	pureNumStrToComponents
//
//	Receives and analyzes the numeric digits configured
//	in a Pure Number String. The results of this analysis,
//	including number sign, number type, absolute value
//	number string and key statistics relating to the
//	numeric value contained in the Pure Number String
//	will be returned by an instance of type,
//	PureNumberStrComponents.
//
//	A Pure Number String consists of numeric digits, an
//	optional radix point and an optional leading or
//	trailing minus ('-') sign.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	pureNumberStr				string
//
//		This Pure Number String contains the numeric
//		character digits which will be analyzed and
//		reported in the returned instance of
//		PureNumberStrComponents.
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
//				a Pure Number String. The decimal
//				separator may consist of one or more text
//				characters.
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
//				processed by the Pure Number String
//				parsing algorithm. All other characters
//				will be	ignored.
//
//			5.	Pure Number Strings consist of a single
//				numeric value. The entire Pure Number String
//				will be parsed, or processed, and only one
//				numeric value per Pure Number String will
//				be returned.
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol.
//
//		The Decimal Separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted, floating
//		point Number String.
//
//		In the US, UK, Australia, most of Canada and many
//		other countries, the Decimal Separator is the
//		period character ('.') known as the decimal
//		point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingMinusSign			bool
//
//		In Pure Number Strings, a minus sign ('-')
//		identifies a number as a negative numeric value.
//
//		When 'leadingMinusSign' is set to 'true', the
//		Pure Number String parsing algorithm will search
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
//		Pure Number String parsing algorithm will search
//		for trailing minus signs ('-') located at the end
//		of the number string. Trailing minus signs
//		represent the standard for France, Germany and
//		many countries in the European Union.
//
//		NOTE: Identification of a trailing minus sign in
//		the Pure Number String input parameter,
//		'pureNumberString', will immediately terminate
//		the search for additional numeric characters.
//
//		Example Trailing Number Symbols:
//			"123.456-" or "123.456 -"
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
//	pureNumStrComponents		PureNumberStrComponents
//
//		If this method completes successfully, this
//		parameter will return an instance of
//		PureNumberStrComponents. This data structure
//		contains an analysis and detail information on
//		the Pure Number String passed as input paramter,
//		'pureNumberStr'.
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
//
//			The number string expressed as an absolute value.
//			Be advised, this number string may be a floating
//			point number string containing fractional digits.
//
//			AbsoluteValAllIntegerDigitsNumStr string
//
//			Integer and fractional digits are combined
//			in a single number string without a decimal
//			point separating integer and fractional digits.
//			This string DOES NOT contain a leading number
//			sign (a.k.a. minus sign ('-')
//
//			SignedAllIntegerDigitsNumStr string
//
//			Integer and fractional digits are combined
//			in a single number string without a decimal
//			point separating integer and fractional digits.
//			If the numeric value is negative, a leading
//			minus sign will be prefixed at the beginning
//			of the number string.
//
//			NativeNumberStr string
//
//			A Native Number String representing the base
//			numeric value used to generate these profile
//			number string statistics.
//
//			A valid Native Number String must conform to the
//			standardized formatting criteria defined below:
//
//			 	1. A Native Number String Consists of numeric
//			 	   character digits zero through nine inclusive
//			 	   (0-9).
//
//			 	2. A Native Number String will include a period
//			 	   or decimal point ('.') to separate integer and
//			 	   fractional digits within a number string.
//
//			 	   Native Number String Floating Point Value:
//			 	   				123.1234
//
//			 	3. A Native Number String will always format
//			 	   negative numeric values with a leading minus sign
//			 	   ('-').
//
//			 	   Native Number String Negative Value:
//			 	   				-123.2
//
//			 	4. A Native Number String WILL NEVER include integer
//			 	   separators such as commas (',') to separate
//			 	   integer digits by thousands.
//
//			 	   					NOT THIS: 1,000,000
//			 	   		Native Number String: 1000000
//
//			 	5. Native Number Strings will only consist of:
//
//			 	   (a)	Numeric digits zero through nine inclusive (0-9).
//
//			 	   (b)	A decimal point ('.') for floating point
//			 	   		numbers.
//
//			 	   (c)	A leading minus sign ('-') in the case of
//			 	   		negative numeric values.
//
//		}
//
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrMathAtom *numStrMathAtom) pureNumStrToComponents(
	pureNumberStr string,
	decSeparatorChars string,
	leadingMinusSign bool,
	errPrefDto *ePref.ErrPrefixDto) (
	pureNumStrComponents PureNumberStrComponents,
	err error) {

	if nStrMathAtom.lock == nil {
		nStrMathAtom.lock = new(sync.Mutex)
	}

	nStrMathAtom.lock.Lock()

	defer nStrMathAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrMathAtom."+
			"pureNumStrToComponents()",
		"")

	if err != nil {
		return pureNumStrComponents, err
	}

	if len(pureNumberStr) == 0 {
		err = fmt.Errorf("\n\n%v\n"+
			"Error: Input parameter 'pureNumberStr'\n"+
			"is a zero length string and INVALID!\n",
			ePrefix.String())

		return pureNumStrComponents, err
	}

	var intDigits,
		fracDigits RuneArrayDto

	var radixPoint RuneArrayDto

	radixPoint,
		err = new(RuneArrayDto).NewString(
		decSeparatorChars,
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix)

	if err != nil {
		return pureNumStrComponents, err
	}

	pureNumStrComponents.NumStrStats,
		err = new(numStrMathQuark).
		pureNumStrToRunes(
			pureNumberStr,
			&intDigits,
			&fracDigits,
			&radixPoint,
			leadingMinusSign,
			ePrefix)

	if err != nil {
		return pureNumStrComponents, err
	}

	if pureNumStrComponents.NumStrStats.NumberSign ==
		NumSignVal.Zero() {

		if pureNumStrComponents.NumStrStats.NumOfFractionalDigits > 0 {

			pureNumStrComponents.NumStrStats.NumOfIntegerDigits = 1

			pureNumStrComponents.AbsoluteValueNumStr =
				"0.0"

			pureNumStrComponents.AbsoluteValAllIntegerDigitsNumStr =
				"00"

			pureNumStrComponents.SignedAllIntegerDigitsNumStr =
				"00"

			pureNumStrComponents.NativeNumberStr =
				"0.0"

		} else {

			pureNumStrComponents.AbsoluteValueNumStr =
				"0"

			pureNumStrComponents.AbsoluteValAllIntegerDigitsNumStr =
				"0"

			pureNumStrComponents.SignedAllIntegerDigitsNumStr =
				"0"

			pureNumStrComponents.NativeNumberStr =
				"0"

			pureNumStrComponents.NumStrStats.NumOfIntegerDigits = 1

		}

	} else {

		// MUST BE -
		// pureNumStrComponents.NumStrStats.NumberSign is
		//	NOT EQUAL TO ZERO.

		pureNumStrComponents.AbsoluteValueNumStr =
			string(intDigits.CharsArray) +
				"." +
				string(fracDigits.CharsArray)

		pureNumStrComponents.AbsoluteValAllIntegerDigitsNumStr =
			string(intDigits.CharsArray) +
				string(fracDigits.CharsArray)

		if pureNumStrComponents.NumStrStats.NumberSign ==
			NumSignVal.Negative() {

			pureNumStrComponents.SignedAllIntegerDigitsNumStr =
				"-"

			pureNumStrComponents.NativeNumberStr = "-"

		}

		pureNumStrComponents.SignedAllIntegerDigitsNumStr +=
			pureNumStrComponents.AbsoluteValAllIntegerDigitsNumStr

		pureNumStrComponents.NativeNumberStr +=
			pureNumStrComponents.AbsoluteValueNumStr
	}

	pureNumStrComponents.NativeNumberStr,
		_,
		err = new(NumStrHelper).NormalizeNativeNumStr(
		pureNumStrComponents.NativeNumberStr,
		ePrefix.XCpy("pureNumStrComponents.NativeNumberStr"))

	return pureNumStrComponents, err
}

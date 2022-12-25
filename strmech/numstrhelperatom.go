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
//			2.	Option: A Pure Number String may include
//				a radix point or decimal separator.
//				Decimal separators separate integer and
//				fractional numeric digits in a pure
//				number string. The decimal separator may
//				consist of one or more text characters.
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

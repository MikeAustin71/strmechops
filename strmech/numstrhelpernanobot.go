package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrHelperNanobot
//
// Provides number string utility methods
type numStrHelperNanobot struct {
	lock *sync.Mutex
}

func (nStrHelperNanobot *numStrHelperNanobot) formatNumStrComponents(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	numberSign NumericSignValueType,
	decSeparator DecimalSeparatorSpec,
	intSeparatorDto IntegerSeparatorSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	currencySymbol NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	numStr string,
	err error) {

	if nStrHelperNanobot.lock == nil {
		nStrHelperNanobot.lock = new(sync.Mutex)
	}

	nStrHelperNanobot.lock.Lock()

	defer nStrHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"formatNumStrComponents()",
		"")

	if err != nil {

		return numStr, err
	}

	if integerDigits == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigits' is a nil pointer!\n",
			ePrefix.String())

		return numStr, err
	}

	if fractionalDigits == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigits' is a nil pointer!\n",
			ePrefix.String())

		return numStr, err
	}

	if numberSign.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberSign' is invalid!\n"+
			"numberSign  String Value = %v\n"+
			"numberSign Integer Value = %v\n",
			ePrefix.String(),
			numberSign.String(),
			numberSign.XValueInt())

		return numStr, err
	}

	numOfBadZeros := integerDigits.GetCountLeadingZeros()

	if numOfBadZeros > 0 {

		err = integerDigits.DeleteLeadingTrailingChars(
			numOfBadZeros,
			false,
			ePrefix.XCpy("integerDigits-leading zeros"))

		if err != nil {
			return numStr, err
		}

	}

	numOfFracDigits := fractionalDigits.GetRuneArrayLength()
	numOfIntDigits := integerDigits.GetRuneArrayLength()

	if numOfFracDigits > 0 &&
		decSeparator.GetNumberOfSeparatorChars() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This is a floating point number and the number\n"+
			"of decimal separator characters specified is zero.\n"+
			"Input parameter 'decSeparator' is invalid!\n"+
			"No decimal separator or radix point characters\n"+
			"have been specified.\n",
			ePrefix.String())

		return numStr, err
	}
	var numStrWithIntSeps []rune

	if numOfIntDigits > 0 {

		numStrWithIntSeps,
			err = new(integerSeparatorSpecMolecule).applyIntSeparators(
			&intSeparatorDto,
			integerDigits.GetRuneArray(),
			ePrefix.XCpy("intSeparatorDto<-integerDigits"))

		if err != nil {
			return numStr, err
		}

	} else {

		numStrWithIntSeps = []rune{'0'}
	}

	tempNumStr := string(numStrWithIntSeps)

	if numOfFracDigits > 0 {

		tempNumStr += decSeparator.GetDecimalSeparatorStr()

		tempNumStr += fractionalDigits.GetCharacterString()

	}

	return numStr, err
}

//	formatNumStrComponents
//
//	Creates and returns a fully formatted Number String
//	generated from Number String formatting components
//	passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		formatted and returned as a Number String.
//
//	decSeparator				DecimalSeparatorSpec
//
//		This structure contains the radix point or decimal
//		separator character(s) (a.k.a. decimal point)
//		which be used to separate integer and fractional
//		digits within a formatted Number String.
//
//	intSeparatorDto				IntegerSeparatorSpec
//
//		Type IntegerSeparatorSpec is designed to manage
//		integer separators, primarily thousands separators,
//		for different countries and cultures. The term
//		'integer separators' is used because this type
//		manages both integer grouping and the characters
//		used to separate integer groups.
//
//		In the USA and many other countries, integer
//		numbers are often separated by commas thereby
//		grouping the number into thousands.
//
//		Example: 1,000,000,000
//
//		Other countries and cultures use characters other
//		than the comma to separate integers into thousands.
//		Some countries and cultures do not use thousands
//		separation and instead rely on multiple integer
//		separation characters and grouping sequences for a
//		single integer number. Notable examples of this
//		are found in the 'India Number System' and
//		'Chinese Numerals'.
//
//		Reference:
//			https://en.wikipedia.org/wiki/Indian_numbering_system
//			https://en.wikipedia.org/wiki/Chinese_numerals
//			https://en.wikipedia.org/wiki/Decimal_separator
//
//		The IntegerSeparatorSpec type provides the
//		flexibility necessary to process these complex
//		number separation formats.
//
//		If integer separation is turned off, no error
//		will be returned and integer digits will be
//		displayed as a single string of numeric digits:
//
//			Integer Separation Turned Off: 1000000000
//
//	roundingSpec 				NumStrRoundingSpec
//
//		Numeric Value Rounding Specification. This
//		specification contains all the parameters
//		required to configure and apply a rounding
//		algorithm for floating point Number Strings.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with negative numeric values.
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with positive numeric values.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with zero numeric values.
//
//	numberFieldSpec			NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	numStr						string
//
//		If this method completes successfully, the
//		numeric	value represented by the NumberStrKernel
//		instance, 'numStrKernel', will be returned as a
//		formatted Number String, 'numStr'.
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If errors
//		are	encountered during processing, the returned
//		error Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.

/*
func (nStrHelperNanobot *numStrHelperNanobot) formatNumStrComponents(
	numStrKernel *NumberStrKernel,
	decSeparator DecimalSeparatorSpec,
	intSeparatorDto IntegerSeparatorSpec,
	roundingSpec NumStrRoundingSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	numStr string,
	err error) {

	if nStrHelperNanobot.lock == nil {
		nStrHelperNanobot.lock = new(sync.Mutex)
	}

	nStrHelperNanobot.lock.Lock()

	defer nStrHelperNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"formatNumStrComponents()",
		"")

	if err != nil {

		return numStr, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return numStr, err
	}

	if len(numStrKernel.integerDigits.CharsArray) == 0 &&
		len(numStrKernel.fractionalDigits.CharsArray) == 0 {

		numStr = "0"

		return numStr, err
	}

	err = roundingSpec.IsValidInstanceError(
		ePrefix.XCpy(
			"roundingSpec"))

	if err != nil {

		return numStr, err
	}

	var newNumStrKernel NumberStrKernel

	err = new(numberStrKernelNanobot).copy(
		&newNumStrKernel,
		numStrKernel,
		ePrefix.XCpy(
			"newNumStrKernel<-numStrKernel"))

	if err != nil {
		return numStr, err
	}

	// Performing fractional digit rounding
	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&newNumStrKernel,
		roundingSpec,
		ePrefix.XCpy(
			"newNumStrKernel Rounding"))

	if err != nil {
		return numStr, err
	}

	var numOfFracDigits int

	numOfFracDigits = newNumStrKernel.GetNumberOfFractionalDigits()

	if numOfFracDigits > 0 &&
		decSeparator.GetNumberOfSeparatorChars() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This is a floating point number and the number\n"+
			"of decimal separator characters specified is zero.\n"+
			"Input parameter 'nStrFormatSpec.DecSeparator'\n"+
			"is invalid!\n",
			ePrefix.String())

		return numStr, err
	}

	var numStrWithIntSeps []rune

	numStrWithIntSeps,
		err = new(integerSeparatorSpecMolecule).applyIntSeparators(
		&intSeparatorDto,
		newNumStrKernel.GetIntegerRuneArray(),
		ePrefix.XCpy("intSeparatorDto"))

	if err != nil {
		return numStr, err
	}

	tempNumStr := string(numStrWithIntSeps)

	if numOfFracDigits > 0 {

		tempNumStr += decSeparator.GetDecimalSeparatorStr()

		tempNumStr += newNumStrKernel.GetFractionalString()

	}

	leadingNumSym := ""

	trailingNumSym := ""

	var leadingNumSymPosition, trailingNumSymPosition NumberFieldSymbolPosition

	if newNumStrKernel.numberSign == NumSignVal.Negative() {

		if !negativeNumberSign.IsNOP() {

			leadingNumSym =
				negativeNumberSign.GetLeadingNumberSymbolStr()

			leadingNumSymPosition =
				negativeNumberSign.GetLeadingNumberSymbolPosition()

			trailingNumSym =
				negativeNumberSign.GetTrailingNumberSymbolStr()

			trailingNumSymPosition =
				negativeNumberSign.GetTrailingNumberSymbolPosition()

		}

	}

	if newNumStrKernel.numberSign == NumSignVal.Positive() {

		if !positiveNumberSign.IsNOP() {

			leadingNumSym =
				positiveNumberSign.GetLeadingNumberSymbolStr()

			leadingNumSymPosition =
				positiveNumberSign.GetLeadingNumberSymbolPosition()

			trailingNumSym =
				positiveNumberSign.GetTrailingNumberSymbolStr()

			trailingNumSymPosition =
				positiveNumberSign.GetTrailingNumberSymbolPosition()

		}

	}

	if newNumStrKernel.numberSign == NumSignVal.Zero() {

		if !zeroNumberSign.IsNOP() {

			leadingNumSym =
				zeroNumberSign.GetLeadingNumberSymbolStr()

			leadingNumSymPosition =
				zeroNumberSign.GetLeadingNumberSymbolPosition()

			trailingNumSym =
				zeroNumberSign.GetTrailingNumberSymbolStr()

			trailingNumSymPosition =
				zeroNumberSign.GetTrailingNumberSymbolPosition()

		}

	}

	lenLeadingNumSymbol := len(leadingNumSym)
	lenTrailingNumSymbol := len(trailingNumSym)

	if lenLeadingNumSymbol > 0 &&
		leadingNumSymPosition == NumFieldSymPos.InsideNumField() {

		tempNumStr = leadingNumSym + tempNumStr
	}

	if lenTrailingNumSymbol > 0 &&
		trailingNumSymPosition == NumFieldSymPos.InsideNumField() {

		tempNumStr = tempNumStr + trailingNumSym

	}

	numStr,
		err = new(strMechNanobot).justifyTextInStrField(
		tempNumStr,
		numberFieldSpec.GetNumFieldLength(),
		numberFieldSpec.GetNumFieldJustification(),
		ePrefix.XCpy("numStr<-tempNumStr"))

	if lenLeadingNumSymbol > 0 &&
		leadingNumSymPosition == NumFieldSymPos.OutsideNumField() {

		numStr = leadingNumSym + numStr
	}

	if lenTrailingNumSymbol > 0 &&
		trailingNumSymPosition == NumFieldSymPos.OutsideNumField() {

		numStr = numStr + trailingNumSym

	}

	return numStr, err
}
*/

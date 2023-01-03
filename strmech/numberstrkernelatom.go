package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelAtom - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelAtom struct {
	lock *sync.Mutex
}

//	addFractionalDigit
//
//	Appends a single numeric digit to the end of the internal
//	fractional digits rune array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance contains the fractional digit rune array
//		to which the 'fractionalDigit' rune will be appended.
//
//	fractionalDigit				rune
//
//		A rune with a numeric character between '0' (zero)
//		and '9' (nine) inclusive. This numeric digit will
//		be appended to the end of the internal member
//		variable 'NumberStrKernel.fractionalDigits'
//		contained within the NumberStrKernel input
//		parameter, 'numStrKernel'.
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
//	error
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
func (numStrKernelAtom *numberStrKernelAtom) addFractionalDigit(
	numStrKernel *NumberStrKernel,
	fractionalDigit rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"addFractionalDigit()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if fractionalDigit < '0' ||
		fractionalDigit > '9' {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fractionalDigit' is invalid!\n"+
			"Fractional Rune characters must represent a numberic character between\n"+
			"'0' and '9', inclusive. 'fractionalDigit' fails to meet this criterion.\n"+
			"The rune value of 'fractionalDigit' is %v\n"+
			"The string value of 'fractionalDigit' is %v\n",
			ePrefix.String(),
			fractionalDigit,
			string(fractionalDigit))

		return err
	}

	err = numStrKernel.fractionalDigits.AddChar(
		fractionalDigit,
		true,
		ePrefix.XCpy(
			"numStrKernel.fractionalDigits<-"))

	if err != nil {
		return err
	}

	if numStrKernel.numberValueType !=
		NumValType.FloatingPoint() {

		numStrKernel.numberValueType =
			NumValType.FloatingPoint()
	}

	if fractionalDigit != '0' {
		numStrKernel.isNonZeroValue = true
	}

	return err
}

// addIntegerDigit - Adds a single numeric digit to the internal
// integer digits rune array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance contains the integer digit rune array to
//		which the 'integerDigit' rune will be appended.
//
//	integerDigit            rune
//
//		A rune with a numeric character between '0' (zero) and
//		'9' (nine) inclusive. This numeric digit will be
//		appended to the internal member variable
//		'NumberStrKernel.integerDigits' for NumberStrKernel
//		input parameter 'numStrKernel'.
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
//	error
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
func (numStrKernelAtom *numberStrKernelAtom) addIntegerDigit(
	numStrKernel *NumberStrKernel,
	integerDigit rune,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"addIntegerDigit()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if integerDigit < '0' ||
		integerDigit > '9' {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigit' is invalid!\n"+
			"Integer Runes must represent a numberic character between\n"+
			"'0' and '9', inclusive. 'integerDigit' fails to meet this criterion.\n"+
			"The rune value of 'integerDigit' is %v\n"+
			"The string value of 'integerDigit' is %v\n",
			ePrefix.String(),
			integerDigit,
			string(integerDigit))

		return err
	}

	err = numStrKernel.integerDigits.AddChar(
		integerDigit,
		true,
		ePrefix.XCpy(
			"numStrKernel.integerDigits<-"))

	if err != nil {
		return err
	}

	if numStrKernel.numberValueType !=
		NumValType.FloatingPoint() {

		numStrKernel.numberValueType =
			NumValType.Integer()
	}

	if integerDigit != '0' {
		numStrKernel.isNonZeroValue = true
	}

	return err
}

//	calcNumStrKernelStats
//
//	Receives a pointer to an instance of NumberStrKernel
//	and proceeds to analyze that instance to produce
//	statistical information on the encapsulated numeric
//	value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		numeric value contained in this instance will be
//		analyzed to produce a statistics on the nature of
//		the numeric value encapsulated by this instance.
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
//	numStrStatsDto				NumberStrStatsDto
//
//		This data transfer object will return key
//		statistics on the numeric value encapsulated
//		by input parameter, 'numStrKernel'.
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
//			If 'false', the subject numeric value is
//			greater than or less than zero ('0').
//
//			If 'true', the Numeric Value is equal to
//			zero.
//		}
//
//	error
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
func (numStrKernelAtom *numberStrKernelAtom) calcNumStrKernelStats(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	numStrStatsDto NumberStrStatsDto,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"calcNumStrKernelStats()",
		"")

	if err != nil {

		return numStrStatsDto, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return numStrStatsDto, err
	}

	numStrStatsDto.NumOfIntegerDigits =
		uint64(len(numStrKernel.integerDigits.CharsArray))

	lenZeros := numStrKernel.integerDigits.GetCountLeadingZeros()

	numStrStatsDto.NumOfSignificantIntegerDigits =
		numStrStatsDto.NumOfIntegerDigits - lenZeros

	numStrStatsDto.NumOfFractionalDigits =
		uint64(len(numStrKernel.fractionalDigits.CharsArray))

	lenZeros = numStrKernel.fractionalDigits.GetCountTrailingZeros()

	numStrStatsDto.NumOfSignificantFractionalDigits =
		numStrStatsDto.NumOfFractionalDigits - lenZeros

	numStrStatsDto.NumberValueType,
		err = new(numberStrKernelQuark).getSetNumValueType(
		numStrKernel,
		ePrefix.XCpy(
			"numStrStatsDto.NumberValueType<-"+
				"numStrKernel"))

	if err != nil {

		return numStrStatsDto, err

	}

	numStrStatsDto.NumberSign =
		numStrKernel.numberSign

	numStrStatsDto.IsZeroValue,
		err = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		ePrefix.XCpy(
			"numStrKernel"))

	if err != nil {

		return numStrStatsDto, err

	}

	numStrStatsDto.IsZeroValue = !numStrStatsDto.IsZeroValue

	return numStrStatsDto, err
}

//	emptyFractionalDigits
//
//	Receives an instance of NumberStrKernel and proceeds
//	to set the internal fractional digits array to 'nil'.
//	This effectively deletes the fractional part of the
//	numeric value contained within the passed instance
//	of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance contains the fractional digit rune array
//		which will be deleted and set to 'nil'.
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
//	error
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
func (numStrKernelAtom *numberStrKernelAtom) emptyFractionalDigits(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"emptyFractionalDigits()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	numStrKernel.fractionalDigits.Empty()

	numStrKernel.fractionalDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	var isNonZero bool

	isNonZero,
		err = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		ePrefix)

	if err != nil {

		return err
	}

	if isNonZero == false {

		numStrKernel.numberSign = NumSignVal.Zero()

		numStrKernel.numberValueType = NumValType.None()

		return err
	}

	// integer digits exist

	numStrKernel.numberValueType = NumValType.Integer()

	return err
}

//	emptyIntegerDigits
//
//	Receives an instance of NumberStrKernel and proceeds to set
//	the internal integer digits array to 'nil'. This effectively
//	deletes the integer part of the numeric value contained
//	within the passed instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		instance contains the integer digit rune array
//		which will be deleted and set to 'nil'.
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
//	error
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
func (numStrKernelAtom *numberStrKernelAtom) emptyIntegerDigits(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"emptyIntegerDigits()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	numStrKernel.integerDigits.Empty()

	numStrKernel.integerDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	var isNonZero bool

	isNonZero,
		err = new(numberStrKernelElectron).getSetIsNonZeroValue(
		numStrKernel,
		ePrefix)

	if err != nil {

		return err
	}

	if isNonZero == false {

		numStrKernel.numberSign = NumSignVal.Zero()

		numStrKernel.numberValueType = NumValType.None()

		return err
	}

	// fractional digits exist

	numStrKernel.numberValueType = NumValType.FloatingPoint()

	return err
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
func (numStrKernelAtom *numberStrKernelAtom) formatNumStrComponents(
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

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

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

//	prepareCompareNumStrKernels
//
//	This method receives pointers to two instances of
//	NumberStrKernel, 'numStrKernel01' and
//	'numStrKernel02'.
//
//	These two instances will be revamped to ensure that
//	their integer digits and fractional digits are of
//	equal lengths. Next, this method will proceed to
//	compare the numeric values of 'numStrKernel01' and
//	'numStrKernel02'. The comparison results will be
//	returned as one of three integer values:
//
//		-1	= numStrKernel01 is less than numStrKernel02
//		 0	= numStrKernel01 is equal to numStrKernel02
//		+1	= numStrKernel01 is greater than numStrKernel02
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel01				*NumberStrKernel
//
//		The numeric value of numStrKernel01 will be
//		compared to that of numStrKernel02. The
//		comparison results will be returned as an integer
//		value.
//
//	numStrKernel02				*NumberStrKernel
//
//		The numeric value of numStrKernel01 will be
//		compared to that of this parameter,
//		numStrKernel02. The comparison results will be
//		returned as an integer value.
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
//	comparisonValue				int
//
//		This parameter will return the results of numeric
//		value comparisons for input parameters,
//		'numStrKernel01' and 'numStrKernel02'. The
//		integer comparison result will be set to one of
//		three values:
//
//		-1	= numStrKernel01 is less than numStrKernel02
//		 0	= numStrKernel01 is equal to numStrKernel02
//		+1	= numStrKernel01 is greater than numStrKernel02
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
func (numStrKernelAtom *numberStrKernelAtom) prepareCompareNumStrKernels(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	comparisonValue int,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelAtom."+
			"calcNumStrKernelStats()",
		"")

	if err != nil {

		return comparisonValue, err

	}

	if numStrKernel01 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel01' is a nil pointer!\n",
			ePrefix.String())

		return comparisonValue, err
	}

	if numStrKernel02 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel02' is a nil pointer!\n",
			ePrefix.String())

		return comparisonValue, err
	}

	var numStrKernel01V2, numStrKernel02V2 NumberStrKernel

	numStrKernel01V2,
		err = numStrKernel01.CopyOut(
		ePrefix.XCpy(
			"numStrKernel01V2<-numStrKernel01"))

	if err != nil {

		return comparisonValue, err

	}

	numStrKernel02V2,
		err = numStrKernel02.CopyOut(
		ePrefix.XCpy(
			"numStrKernel02V2<-numStrKernel02"))

	if err != nil {

		return comparisonValue, err

	}

	lenIntDigits01V2 :=
		len(numStrKernel01V2.integerDigits.CharsArray)

	lenFracDigits01V2 :=
		len(numStrKernel01V2.fractionalDigits.CharsArray)

	lenIntDigits02V2 :=
		len(numStrKernel02V2.integerDigits.CharsArray)

	lenFracDigits02V2 :=
		len(numStrKernel02V2.fractionalDigits.CharsArray)

	numStrKernelElectron := numberStrKernelElectron{}

	if lenIntDigits01V2 != lenIntDigits02V2 ||
		lenFracDigits01V2 != lenFracDigits02V2 {

		err = numStrKernelElectron.
			equalizeNumStrDigitsLengths(
				&numStrKernel01V2,
				&numStrKernel02V2,
				ePrefix.XCpy(
					"numStrKernel01V2 - numStrKernel02V2"))

		if err != nil {

			return comparisonValue, err

		}
	}

	// Integer Digits and Fractional Digits are NOW
	// equivalent between:
	//    numStrKernel01V2 and numStrKernel02V2

	comparisonValue,
		err = new(numberStrKernelQuark).
		compareNumStrKernelValues(
			&numStrKernel01V2,
			&numStrKernel02V2,
			ePrefix.XCpy(
				"numStrKernel01V2 & numStrKernel02V2"))

	return comparisonValue, err
}

// testValidityOfNumStrKernel - Receives a pointer to an instance
// of NumberStrKernel and performs a diagnostic analysis to
// determine if that instance is valid in all respects.
//
// If the input parameter 'numStrKernel' is determined to be
// invalid, this method will return a boolean flag ('isValid') of
// 'false'. In addition, an instance of type error ('err') will be
// returned configured with an appropriate error message.
//
// If the input parameter 'numStrKernel' is valid, this method will
// return a boolean flag ('isValid') of 'true' and the returned
// error type ('err') will be set to 'nil'.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		object will be subjected to diagnostic analysis in
//		order to determine if all the member variables
//		contain valid values.
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
//	isValid						bool
//
//		If input parameter 'numStrKernel' is judged to be valid in
//		all respects, this return parameter will be set to 'true'.
//
//		If input parameter 'numStrKernel' is found to be invalid,
//		this return parameter will be set to 'false'.
//
//	err							error
//
//		If this method completes successfully and
//	 	'numStrKernel' evaluates as valid, this returned
//		error Type is set equal to 'nil'.
//
//		If errors are encountered or if 'numStrKernel'
//		evaluates as invalid, the returned error Type
//		will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
//
//		NOTE: If the numeric value of 'numStrKernel' exceeds
//		the maximum value for type int, an error will be
//		returned.
func (numStrKernelAtom *numberStrKernelAtom) testValidityOfNumStrKernel(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if numStrKernelAtom.lock == nil {
		numStrKernelAtom.lock = new(sync.Mutex)
	}

	numStrKernelAtom.lock.Lock()

	defer numStrKernelAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"testValidityOfNumStrKernel()",
		"")

	if err != nil {

		return isValid, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	lenIntDigits := numStrKernel.integerDigits.GetRuneArrayLength()

	lenFracDigits := numStrKernel.fractionalDigits.GetRuneArrayLength()

	if lenIntDigits == 0 &&
		lenFracDigits == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!\n"+
			"Both Integer Digits and Fractional Digits are empty\n"+
			"and contain zero digits.\n",
			ePrefix.String())

		return isValid, err
	}

	if lenIntDigits == 0 &&
		lenFracDigits > 0 {

		err = numStrKernel.AddIntegerDigit(
			'0',
			ePrefix.XCpy("Adding Missing Zero Digit"))

		if err != nil {
			return isValid, err
		}
	}

	if !numStrKernel.numberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!\n"+
			"The Number Sign Value is invalid.\n"+
			"Valid Number Sign Values are:\n"+
			"   NumSignVal.Negative()\n"+
			"   NumSignVal.Zero()\n"+
			"   NumSignVal.Positive()\n"+
			"The current Number Sign Value is:\n"+
			"   Number Sign String Value = '%v'\n"+
			"  Number Sign Integer Value = '%v\n",
			ePrefix.String(),
			numStrKernel.numberSign.String(),
			numStrKernel.numberSign.XValueInt())

		return isValid, err
	}

	numValHasNonZeroVal := false

	for i := 0; i < lenIntDigits; i++ {

		if numStrKernel.integerDigits.CharsArray[i] < '0' &&
			numStrKernel.integerDigits.CharsArray[i] > '9' {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!\n"+
				"The Integer Digits rune array contains non-numeric characters.\n",
				ePrefix.String())

			return isValid, err
		}
	}

	for i := 0; i < lenIntDigits; i++ {

		if numStrKernel.integerDigits.CharsArray[i] >= '1' &&
			numStrKernel.integerDigits.CharsArray[i] <= '9' {
			numValHasNonZeroVal = true
			break
		}
	}

	for j := 0; j < lenFracDigits; j++ {

		if numStrKernel.fractionalDigits.CharsArray[j] < '0' &&
			numStrKernel.fractionalDigits.CharsArray[j] > '9' {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!\n"+
				"The Fractional Digits rune array contains non-numeric characters.\n",
				ePrefix.String())

			return isValid, err
		}

	}

	if !numValHasNonZeroVal {

		for j := 0; j < lenFracDigits; j++ {

			if numStrKernel.fractionalDigits.CharsArray[j] >= '1' &&
				numStrKernel.fractionalDigits.CharsArray[j] <= '9' {
				numValHasNonZeroVal = true
				break
			}

		}

	}

	if numValHasNonZeroVal &&
		numStrKernel.numberSign == NumSignVal.Zero() {

		err = fmt.Errorf("%v\n"+
			"Error: This instance of NumberStrKernel is invalid!\n"+
			"The Number Sign Value is invalid.\n"+
			"NumberStrKernel has a non-zero numeric value.\n"+
			"However, Number Sign is equal to Zero.\n"+
			"Number Sign = NumSignVal.Zero()\n",
			ePrefix.String())

		return isValid, err
	}

	if numValHasNonZeroVal != numStrKernel.isNonZeroValue {

		if numValHasNonZeroVal == false &&
			numStrKernel.isNonZeroValue == true {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!\n"+
				"The Number Sign Value is invalid.\n"+
				"NumberStrKernel has a zero numeric value.\n"+
				"However, internal flag numStrKernel.isNonZeroValue\n"+
				"is set to 'true'.\n",
				ePrefix.String())

			return isValid, err

		}

		if numValHasNonZeroVal == true &&
			numStrKernel.isNonZeroValue == false {

			err = fmt.Errorf("%v\n"+
				"Error: This instance of NumberStrKernel is invalid!\n"+
				"The Number Sign Value is invalid.\n"+
				"NumberStrKernel has a non-zero numeric value.\n"+
				"However, internal flag numStrKernel.isNonZeroValue\n"+
				"is set to 'false'.\n",
				ePrefix.String())

			return isValid, err

		}
	}

	if new(numStrFmtSpecNanobot).isNOP(
		&numStrKernel.numStrFormatSpec) {

		// This is a NOP!
		// numStrKernel.numStrFormatSpec is invalid.
		// Set default numStrKernel.numStrFormatSpec
		// to stand US Signed Number String Format
		// Specification.
		numStrKernel.numStrFormatSpec,
			err = new(NumStrFormatSpec).NewSignedNumFmtUS(
			NumStrNumberFieldSpec{
				fieldLength:        -1,
				fieldJustification: TxtJustify.Right(),
			},
			ePrefix.XCpy(
				"numStrKernel.numStrFormatSpec "+
					"Default US NumStrNumberFieldSpec"))

		if err != nil {
			return isValid, err
		}
	}

	isValid = true

	return isValid, err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// numStrHelperNanobot
//
// Provides number string utility methods
type numStrHelperNanobot struct {
	lock *sync.Mutex
}

//	formatNumStrElements
//
//	Creates and returns a fully formatted Number String
//	generated from Number String formatting components
//	passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	integerDigits				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto which
//		contains a rune array consisting of the integer
//		numeric digits to be included in the final
//		returned Number String.
//
//		If this rune array is empty and contains zero
//		characters, it will be automatically configured
//		with a single zero ('0') character.
//
//		If this rune array contains any non-numeric
//		characters, an error will be returned.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto which
//		contains a rune array consisting of the
//		fractional numeric digits to be included in the
//		final returned Number String.
//
//		If this rune array is empty and contains zero
//		characters, the final returned Number String will
//		be configured as an integer number.
//
//		If this rune array contains any non-numeric
//		characters, an error will be returned.
//
//	numberSign					NumericSignValueType
//
//		An enumeration specifying the number sign
//		associated with the numeric value represented by
//		the integer and fractional numeric digits passed
//		as input parameters.
//
//		Possible values are listed as follows:
//
//	 	NumSignVal.None() - Invalid Value
//	 	NumSignVal.Negative() = -1
//	 	NumSignVal.Zero()     =  0
//	 	NumSignVal.Positive() =  1
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
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Number String Integer Separator Specification. This
//		type encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
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
//		Examples:
//
//			IntGroupingType.None()
//			(a.k.a Integer Separation Turned Off)
//				'1000000000'
//
//			IntGroupingType.Thousands()
//					'1,000,000,000'
//
//			IntGroupingType.IndiaNumbering()
//				'6,78,90,00,00,00,00,000'
//
//			IntGroupingType.ChineseNumbering()
//				'6,7890,0000,0000,0000'
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
//	negativeNumberSign				NumStrNumberSymbolSpec
//
//		The Number String Negative Number Sign
//		Specification is used to configure negative
//		number sign symbols for negative numeric
//		values formatted and displayed in number
//		stings.
//
//		If this parameter is submitted as an empty or
//		invalid Negative Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Negative
//			Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Negative
//			Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	positiveNumberSign 				NumStrNumberSymbolSpec
//
//		Positive number signs are commonly implied
//		and not specified. However, the user has
//		the option to specify a positive number sign
//		character or characters for positive numeric
//		values using this input parameter.
//
//		If this parameter is submitted as an empty or
//		invalid Positive Number Sign Specification, it
//		will be automatically converted to a 'NOP' or
//		empty placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Positive
//			Values
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Positive
//			Values
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//	zeroNumberSign					NumStrNumberSymbolSpec
//
//		The Number String Zero Number Sign
//		Specification is used to configure number
//		sign symbols for zero numeric values formatted
//		and displayed in number stings. Zero number signs
//		are commonly omitted because zero does not
//		technically qualify as either a positive or
//		negative value. However, the user has the option
//		to configure number sign symbols for zero values
//		if necessary.
//
//		If this parameter is submitted as an empty or
//		invalid Zero Number Sign Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Example-1: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		Example-2: Leading Number Sign Symbols
//			Leading Number Sign Symbols for Zero Values
//
//			Leading Symbols: "+ "
//			Trailing Symbols: ""
//			Number String:   "+ 0.00"
//
//		Example-3: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		Example-4: Trailing Number Sign Symbols
//			Trailing Number Sign Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: "+"
//			Number String:   "0.00+"
//
//	currencySymbol					NumStrNumberSymbolSpec
//
//		A Currency Symbol next to a number shows the
//		number is a monetary amount.
//
//		The Number String Currency Symbol Specification
//		is used to configure currency symbols for
//		positive, negative and zero numeric values
//		formatted and displayed in number stings.
//
//		If this parameter is submitted as an empty or
//		invalid Currency Symbol Specification, it will
//		be automatically converted to a 'NOP' or empty
//		placeholder which will be ignored by Number
//		String formatting algorithms. 'NOP' is a computer
//		science term meaning 'No Operation'.
//
//		Examples of Currency Symbols include the Dollar
//		sign ('$'), Euro sign ('€') or Pound sign ('£').
//
//		This instance of NumStrNumberSymbolSpec is used
//		to configure leading Currency Symbols, trailing
//		Currency Symbols or both leading and trailing
//		Currency Symbols.
//
//		Example-1: Leading Currency Symbols
//
//			Leading Currency Symbols: "$ "
//			Number String:   "$ 123.456"
//
//		Example-2: Leading Currency Symbols
//
//			Leading Currency Symbols: "$"
//			Number String:   "$123.456"
//
//		Example-3: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: "€"
//			Number String:   "123.456€"
//
//		Example-4: Trailing Currency Symbols
//			Trailing Currency Symbols for Positive Values
//
//			Trailing Currency Symbols: " €"
//			Number String:   "123.456 €"
//
//	numberFieldSpec				NumStrNumberFieldSpec
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
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
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
//	numStr						string
//
//		If this method completes successfully, the
//		numeric	value represented by input parameters
//		'integerDigits' and 'fractionalDigits' will be
//		returned as a formatted Number String, 'numStr'.
//
//	err							error
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
func (nStrHelperNanobot *numStrHelperNanobot) formatNumStrElements(
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	numberSign NumericSignValueType,
	decSeparator DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
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
			"formatNumStrElements()",
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

	numOfIntDigits := integerDigits.GetRuneArrayLength()

	if numOfIntDigits == 0 {

		err = integerDigits.AddChar('0',
			true,
			ePrefix.XCpy("integerDigits"))

		if err != nil {

			return numStr, err
		}
	}

	numOfFracDigits := fractionalDigits.GetRuneArrayLength()

	if integerDigits.IsAllNumericDigits() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'integerDigits' is invalid!\n"+
			"'integerDigits' contains non-numeric digits\n"+
			"integerDigits = %v\n",
			ePrefix.String(),
			integerDigits.GetCharacterString())

		return numStr, err

	}

	if numOfFracDigits > 0 &&
		fractionalDigits.IsAllNumericDigits() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fractionalDigits' is invalid!\n"+
			"'fractionalDigits' contains non-numeric digits\n"+
			"fractionalDigits = %v\n",
			ePrefix.String(),
			fractionalDigits.GetCharacterString())

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

	if numberSign == NumSignVal.Negative() &&
		negativeNumberSign.IsNOP() {

		err = fmt.Errorf("%v\n"+
			"Error: The numeric value is negative however\n"+
			"no negative number sign has been configured.\n",
			ePrefix.String())

		return numStr, err
	}

	if !numberFieldSpec.IsValidInstance() {

		numberFieldSpec.SetNOP()
	}

	var err2 error

	err2 = numberFieldSpec.IsValidInstanceError(
		ePrefix.XCpy(
			"numberFieldSpec"))

	if err2 != nil {

		err = fmt.Errorf(
			"Error: Input parameter 'numberFieldSpec' is invalid!\n"+
				"Error = \n%v\n",
			err2.Error())

		return numStr, err
	}

	numOfBadZeros := integerDigits.GetCountLeadingZeros()

	if numOfBadZeros > 1 {

		numOfBadZeros--

		err = integerDigits.DeleteLeadingTrailingChars(
			numOfBadZeros,
			false,
			ePrefix.XCpy("integerDigits-leading zeros"))

		if err != nil {
			return numStr, err
		}

	}

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

	numOfIntDigits = integerDigits.GetRuneArrayLength()

	if numOfIntDigits > 0 {

		numStrWithIntSeps,
			err = new(integerSeparatorSpecMolecule).applyIntSeparators(
			&intSeparatorSpec,
			integerDigits.GetRuneArray(),
			ePrefix.XCpy("intSeparatorSpec<-integerDigits"))

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

	var outsideNumFieldLeadingSymbols,
		outsideNumFieldTrailingSymbols,
		insideNumFieldLeadingSymbols,
		insideNumFieldTrailingSymbols string

	var leadingNumFieldSymPos,
		trailingNumFieldSymPos NumberFieldSymbolPosition

	var numSignSymbolSpec NumStrNumberSymbolSpec

	if err != nil {

		return numStr, err
	}

	if numberSign == NumSignVal.Negative() {

		if negativeNumberSign.IsNOP() == true {

			err = fmt.Errorf("%v\n"+
				"Error: The Negative Number Sign Symbol Specification is invalid!\n"+
				"'numberSign' specifies a negative numeric value, but no negative\n"+
				"number sign symbols have been configured in input parameter,\n"+
				"'negativeNumberSign'.\n",
				ePrefix.String())

			return numStr, err

		}

		numSignSymbolSpec,
			err = negativeNumberSign.CopyOut(
			ePrefix.XCpy(
				"numSignSymbolSpec<-negativeNumberSign"))

		if err != nil {

			return numStr, err
		}

	} else if numberSign == NumSignVal.Positive() {

		if positiveNumberSign.IsNOP() == true {

			numSignSymbolSpec.SetNOP()

		} else {

			numSignSymbolSpec,
				err = positiveNumberSign.CopyOut(
				ePrefix.XCpy(
					"numSignSymbolSpec<-positiveNumberSign"))

			if err != nil {

				return numStr, err
			}
		}

	} else {
		// MUST BE !
		//	numberSign == NumSignVal.Zero()

		if zeroNumberSign.IsNOP() == true {

			numSignSymbolSpec.SetNOP()

		} else {

			numSignSymbolSpec,
				err = zeroNumberSign.CopyOut(
				ePrefix.XCpy(
					"numSignSymbolSpec<-zeroNumberSign"))

			if err != nil {

				return numStr, err
			}
		}

	}

	if numSignSymbolSpec.IsNOP() == false {

		if numSignSymbolSpec.leadingNumberSymbols.
			GetRuneArrayLength() > 0 {

			leadingNumFieldSymPos =
				numSignSymbolSpec.
					leadingNumberFieldSymbolPosition

			if leadingNumFieldSymPos ==
				NumFieldSymPos.InsideNumField() {

				insideNumFieldLeadingSymbols =
					numSignSymbolSpec.
						leadingNumberSymbols.
						GetCharacterString()

			} else if leadingNumFieldSymPos ==
				NumFieldSymPos.OutsideNumField() {

				outsideNumFieldLeadingSymbols =
					numSignSymbolSpec.
						leadingNumberSymbols.
						GetCharacterString()

			} else {

				err = fmt.Errorf("%v\n"+
					"Error: %v Leading Number Sign Symbol Position is invalid!"+
					"%vNumberSign.leadingNumberFieldSymbolPosition  String Value %v\n"+
					"%vNumberSign.leadingNumberFieldSymbolPosition Integer Value %v\n",
					ePrefix.String(),
					numberSign.String(),
					strings.ToLower(numberSign.String()),
					strings.ToLower(numberSign.String()),
					leadingNumFieldSymPos.String(),
					leadingNumFieldSymPos.XValueInt())

				return numStr, err
			}

		}

		if numSignSymbolSpec.trailingNumberSymbols.
			GetRuneArrayLength() > 0 {

			trailingNumFieldSymPos =
				numSignSymbolSpec.
					trailingNumberFieldSymbolPosition

			if trailingNumFieldSymPos ==
				NumFieldSymPos.InsideNumField() {

				insideNumFieldTrailingSymbols =
					numSignSymbolSpec.
						trailingNumberSymbols.
						GetCharacterString()

			} else if trailingNumFieldSymPos ==
				NumFieldSymPos.OutsideNumField() {

				outsideNumFieldTrailingSymbols =
					numSignSymbolSpec.
						trailingNumberSymbols.
						GetCharacterString()

			} else {

				err = fmt.Errorf("%v\n"+
					"Error: %v Trailing Number Sign Symbol Position is invalid!"+
					"%vNumberSign.trailingNumberFieldSymbolPosition  String Value %v\n"+
					"%vNumberSign.trailingNumberFieldSymbolPosition Integer Value %v\n",
					ePrefix.String(),
					numberSign.String(),
					strings.ToLower(numberSign.String()),
					strings.ToLower(numberSign.String()),
					trailingNumFieldSymPos.String(),
					trailingNumFieldSymPos.XValueInt())

				return numStr, err
			}

		}

	}

	var currencySymbolStr string

	if currencySymbol.IsNOP() == false {
		// Currency Symbols are configured

		if currencySymbol.
			leadingNumberSymbols.
			GetRuneArrayLength() > 0 {

			// We have a leading currency symbol
			//	$ 123.45
			currencySymbolStr =
				currencySymbol.
					leadingNumberSymbols.
					GetCharacterString()

			if leadingNumFieldSymPos ==
				currencySymbol.
					leadingNumberFieldSymbolPosition {
				// Both Leading Number sign and Leading
				//	Currency Symbol Field Positions
				//	are EQUAL!

				if leadingNumFieldSymPos ==
					NumFieldSymPos.InsideNumField() {
					// We have a Leading Number sign and a
					// Leading Currency Symbol.
					//
					// Both Symbols will go inside the Number Field

					if currencySymbol.
						currencyNumSignRelativePos ==
						CurrNumSignRelPos.InsideNumSign() {

						// We have a Leading Number sign and a
						// Leading Currency Symbol.
						//
						// Both Symbols will go inside the
						//	Number Field.
						//
						// Leading Currency Symbol goes inside
						// the Leading Number sign '-$'

						// " - $123.45"
						insideNumFieldLeadingSymbols =
							insideNumFieldLeadingSymbols +
								currencySymbolStr

					} else if currencySymbol.
						currencyNumSignRelativePos ==
						CurrNumSignRelPos.OutsideNumSign() {
						// We have a Leading Number sign and a
						// Leading Currency Symbol.
						//
						// Both Symbols will go inside the
						// Number Field
						//
						// Leading Currency Symbol goes outside
						// the Leading Number sign '$-'

						//  " $ -123.45"
						insideNumFieldLeadingSymbols =
							currencySymbolStr +
								insideNumFieldLeadingSymbols

					} else {

						err = fmt.Errorf("%v\n"+
							"Error: Currency Number Sign Relative Position is invalid!\n"+
							"currencySymbol.currencyNumSignRelativePos  String Value = %v\n"+
							"currencySymbol.currencyNumSignRelativePos Integer Value = %v\n",
							ePrefix.String(),
							currencySymbol.currencyNumSignRelativePos.String(),
							currencySymbol.currencyNumSignRelativePos.XValueInt())

						return numStr, err
					}

				} else {
					// leadingNumFieldSymPos ==
					//	NumFieldSymPos.OutsideNumField()
					//
					// We have a Leading Number sign and
					//	a Leading Currency Symbol.
					//
					// Both Leading Currency Symbol and
					//	Leading Number Sign Symbols will go
					//	outside the Number Field
					//		$- " 123.45"

					if currencySymbol.
						currencyNumSignRelativePos ==
						CurrNumSignRelPos.InsideNumSign() {
						// We have a Leading Number sign and
						//	a Leading Currency Symbol.
						//
						// Both Leading Currency Symbol and
						//	Leading Number Sign Symbols will go
						//	outside the Number Field
						//
						// Leading Currency Symbol is inside
						//	the Leading Number Sign
						//	-$" 123.45"
						outsideNumFieldLeadingSymbols =
							outsideNumFieldLeadingSymbols +
								currencySymbolStr

					} else if currencySymbol.
						currencyNumSignRelativePos ==
						CurrNumSignRelPos.OutsideNumSign() {
						// We have a Leading Number sign and
						//	a Leading Currency Symbol.
						//
						// Both Leading Currency Symbol and
						//	Leading Number Sign Symbols will go
						//	outside the Number Field
						//
						// Leading Currency Symbol is outside
						//	the Leading Number Sign
						//	$-" 123.45"

						outsideNumFieldLeadingSymbols =
							currencySymbolStr +
								outsideNumFieldLeadingSymbols

					} else {

						err = fmt.Errorf("%v\n"+
							"Error: Currency Number Sign Relative Position is invalid!\n"+
							"currencySymbol.currencyNumSignRelativePos  String Value = %v\n"+
							"currencySymbol.currencyNumSignRelativePos Integer Value = %v\n",
							ePrefix.String(),
							currencySymbol.currencyNumSignRelativePos.String(),
							currencySymbol.currencyNumSignRelativePos.XValueInt())

						return numStr, err

					}

				}

			} else {
				// MUST BE !
				// leadingNumFieldSymPos NOT EQUAL TO
				//	currencySymbol.leadingNumberFieldSymbolPosition
				//
				// Currency and Number Sign Field Positions
				//	ARE NOT EQUAL!

				if currencySymbol.
					leadingNumberFieldSymbolPosition ==
					NumFieldSymPos.InsideNumField() {

					// Currency and Number Sign Field Positions
					//	ARE NOT EQUAL!
					//
					// Leading Currency Symbol is inside the
					//	Number Field. Leading Number Sign is
					//	outside the Number Field.
					//		-" $123.45"
					insideNumFieldLeadingSymbols +=
						currencySymbolStr

				} else if currencySymbol.
					leadingNumberFieldSymbolPosition ==
					NumFieldSymPos.OutsideNumField() {

					// Currency and Number Sign Field Positions
					//	ARE NOT EQUAL!
					//
					// Leading Currency Symbol is outside the
					//	Number Field. Leading Number Sign is
					//	inside the Number Field.
					//		$" -123.45"

					outsideNumFieldLeadingSymbols +=
						currencySymbolStr

				} else {

					err = fmt.Errorf("%v\n"+
						"Error: Currency Symbol Leading Field Symbol Position is invalid!\n"+
						"currencySymbol.leadingNumberFieldSymbolPosition  String Value = %v\n"+
						"currencySymbol.leadingNumberFieldSymbolPosition Integer Value = %v\n",
						ePrefix.String(),
						currencySymbol.leadingNumberFieldSymbolPosition.String(),
						currencySymbol.leadingNumberFieldSymbolPosition.XValueInt())

					return numStr, err
				}
			}

		}

		if currencySymbol.
			trailingNumberSymbols.
			GetRuneArrayLength() > 0 {

			currencySymbolStr =
				currencySymbol.
					trailingNumberSymbols.
					GetCharacterString()

			// We have a trailing currency symbol
			// 	123.45 €

			if trailingNumFieldSymPos ==
				currencySymbol.
					trailingNumberFieldSymbolPosition {
				// Both Trailing Number sign and Trailing
				//	Currency Symbol Field Positions
				//	are EQUAL!

				if trailingNumFieldSymPos ==
					NumFieldSymPos.InsideNumField() {
					// We have a Trailing Number Sign and a
					// Trailing Currency Symbol.
					//
					// Both Symbols will go inside the Number Field

					if currencySymbol.
						currencyNumSignRelativePos ==
						CurrNumSignRelPos.InsideNumSign() {

						// We have a Trailing Number sign and a
						// Trailing Currency Symbol.
						//
						// Both Symbols will go inside the Number Field
						//
						// Trailing Currency Symbol goes inside
						// the Trailing Number sign '€-'

						//  "123.45 € -"
						insideNumFieldTrailingSymbols =
							currencySymbolStr +
								insideNumFieldTrailingSymbols

					} else if currencySymbol.
						currencyNumSignRelativePos ==
						CurrNumSignRelPos.OutsideNumSign() {
						// We have a Trailing Number sign and a
						// Trailing Currency Symbol.
						//
						// Both Symbols will go inside the
						// Number Field
						//
						// Trailing Currency Symbol goes outside
						// the Trailing Number sign '-€'

						//  "123.45 - €"
						outsideNumFieldTrailingSymbols =
							outsideNumFieldTrailingSymbols +
								currencySymbolStr

					} else {

						err = fmt.Errorf("%v\n"+
							"Error: Currency Number Sign Relative Position is invalid!\n"+
							"currencySymbol.currencyNumSignRelativePos  String Value = %v\n"+
							"currencySymbol.currencyNumSignRelativePos Integer Value = %v\n",
							ePrefix.String(),
							currencySymbol.currencyNumSignRelativePos.String(),
							currencySymbol.currencyNumSignRelativePos.XValueInt())

						return numStr, err
					}

				} else {
					// MUST BE!
					//
					// Both Trailing Number sign and Trailing
					//	Currency Symbol Field Positions
					//	are EQUAL!
					//
					// trailingNumFieldSymPos ==
					//	 NumFieldSymPos.OutsideNumField()
					//
					// We have a Trailing Number sign and
					//	a Trailing Currency Symbol.
					//
					// Both Trailing Currency Symbol and
					//	Trailing Number Sign Symbols will go
					//	outside the Number Field
					//		" 123.45" € -

					if currencySymbol.
						currencyNumSignRelativePos ==
						CurrNumSignRelPos.InsideNumSign() {
						// Both Trailing Number sign and Trailing
						//	Currency Symbol Field Positions
						//	are EQUAL!
						//
						// trailingNumFieldSymPos ==
						//	 NumFieldSymPos.OutsideNumField()
						//
						// We have a Trailing Number sign and
						//	a Trailing Currency Symbol.
						//
						// Both Trailing Currency Symbol and
						//	Trailing Number Sign Symbols will go
						//	outside the Number Field
						//
						// The Trailing Currency symbol goes
						//	inside the Trailing Number Sign
						// "123.45" € -

						outsideNumFieldTrailingSymbols =
							currencySymbolStr +
								outsideNumFieldTrailingSymbols

					} else if currencySymbol.
						currencyNumSignRelativePos ==
						CurrNumSignRelPos.OutsideNumSign() {
						// Both Trailing Number sign and Trailing
						//	Currency Symbol Field Positions
						//	are EQUAL!
						//
						// trailingNumFieldSymPos ==
						//	 NumFieldSymPos.OutsideNumField()
						//
						// We have a Trailing Number sign and
						//	a Trailing Currency Symbol.
						//
						// Both Trailing Currency Symbol and
						//	Trailing Number Sign Symbols will go
						//	outside the Number Field
						//
						// The Trailing Currency symbol goes
						//	outside the Trailing Number Sign
						// "123.45"  - €

						outsideNumFieldTrailingSymbols =
							outsideNumFieldTrailingSymbols +
								currencySymbolStr

					} else {

						err = fmt.Errorf("%v\n"+
							"Error: Currency Number Sign Relative Position is invalid!\n"+
							"currencySymbol.currencyNumSignRelativePos  String Value = %v\n"+
							"currencySymbol.currencyNumSignRelativePos Integer Value = %v\n",
							ePrefix.String(),
							currencySymbol.currencyNumSignRelativePos.String(),
							currencySymbol.currencyNumSignRelativePos.XValueInt())

						return numStr, err

					}

				}

			} else {
				// MUST BE !
				// trailingNumFieldSymPos NOT EQUAL TO
				//	currencySymbol.trailingNumberFieldSymbolPosition
				//
				// Currency and Number Sign Field Positions
				//	ARE NOT EQUAL!

				if currencySymbol.
					trailingNumberFieldSymbolPosition ==
					NumFieldSymPos.InsideNumField() {

					// Currency and Number Sign Field Positions
					//	ARE NOT EQUAL!
					//
					// Trailing Currency Symbol is inside the
					//	Number Field. Trailing Number Sign is
					//	outside the Number Field.
					//		" $123.45€" -
					insideNumFieldTrailingSymbols +=
						currencySymbolStr

				} else if currencySymbol.
					trailingNumberFieldSymbolPosition ==
					NumFieldSymPos.OutsideNumField() {

					// Currency and Number Sign Field Positions
					//	ARE NOT EQUAL!
					//
					// Trailing Currency Symbol is outside the
					//	Number Field. Trailing Number Sign is
					//	inside the Number Field.
					//		" 123.45-" €

					outsideNumFieldTrailingSymbols +=
						currencySymbolStr

				} else {

					err = fmt.Errorf("%v\n"+
						"Error: Currency Symbol Trailing Field Symbol Position is invalid!\n"+
						"currencySymbol.trailingNumberFieldSymbolPosition  String Value = %v\n"+
						"currencySymbol.trailingNumberFieldSymbolPosition Integer Value = %v\n",
						ePrefix.String(),
						currencySymbol.trailingNumberFieldSymbolPosition.String(),
						currencySymbol.trailingNumberFieldSymbolPosition.XValueInt())

					return numStr, err
				}
			}

		}

	}

	tempNumStr =
		insideNumFieldLeadingSymbols +
			tempNumStr +
			insideNumFieldTrailingSymbols

	numStr,
		err = new(strMechNanobot).justifyTextInStrField(
		tempNumStr,
		numberFieldSpec.GetNumFieldLength(),
		numberFieldSpec.GetNumFieldJustification(),
		ePrefix.XCpy("numStr<-tempNumStr"))

	if err != nil {

		return numStr, err
	}

	numStr =
		outsideNumFieldLeadingSymbols +
			numStr +
			outsideNumFieldTrailingSymbols

	return numStr, err
}

//	formatNumStrElements
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
func (nStrHelperNanobot *numStrHelperNanobot) formatNumStrElements(
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
			"formatNumStrElements()",
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

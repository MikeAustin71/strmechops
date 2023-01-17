package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrFmtSpecMechanics - This type provides
// helper methods for NumStrFormatSpec
type numStrFmtSpecMechanics struct {
	lock *sync.Mutex
}

// setCurrencyNStrFmtBasic
//
// Receives an instance of NumStrFormatSpec and proceeds
// to reconfigure that instance using a 'basic' set of
// Number String Format Specification parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//
//		This instance will be configured with new
//		NumStrFormatSpec data values extracted from the
//		following input parameters.
//
//	decSeparatorChars				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
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
//	leadingCurrencySymbol			string
//
//		The character or characters which comprise the
//		leading Currency Symbol. The leading Currency
//		Symbol will be positioned at the beginning or
//		left side of the number string.
//
//			Example: $ 123.45
//
//		Note:	If a space between the currency symbol
//				and the first digit of the number string
//				is required, be sure to include the space
//				in the currency symbol string.
//					Example:
//						Leading Currency Symbol: "$ "
//						Formatted Number String: "$ 123.45"
//
//	trailingCurrencySymbol			string
//
//		The character or characters which comprise the
//		trailing Currency Symbol. The trailing Currency
//		Symbol will be positioned at the end of, or
//		right side of, the number string.
//
//			Example: 123.45 €
//
//		Note:	If a space between the last digit of the
//				number string and the currency symbol
//				is required, be sure to include the space
//				in the currency symbol string.
//					Example:
//						Trailing Currency Symbol: " €"
//						Formatted Number String: "123.45 €"
//
//	currencyInsideNumSymbol			bool
//
//		This parameter determines whether the currency
//		symbol will be positioned inside or outside
//		the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example Outside - 1
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example Outside - 2
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 1
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 2
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numSymbolFieldPosition			NumberFieldSymbolPosition
//
//		Defines the position of the Currency and number
//		sign characters relative to a Number Field in
//		which a number string is displayed.
//
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example InsideNumField - 1:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$ -123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example InsideNumField - 2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45- €"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				Example InsideNumField - 3:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example InsideNumField - 4:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "  123.45 €"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example OutsideNumField - 5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$ -  123.45"
//					Number Field Index:       01234567890
//					Total Number String Length: 11
//
//				Example OutsideNumField - 6:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45- €"
//					Number Field Index:       01234567890
//					Total Number String Length: 11
//
//				Example OutsideNumField - 7:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example OutsideNumField - 6:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45 €"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
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
// -----------------------------------------------------------------
//
// # Return Values
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
func (nStrFmtSpecMechanics *numStrFmtSpecMechanics) setCurrencyNStrFmtBasic(
	numStrFmtSpec *NumStrFormatSpec,
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingCurrencySymbol string,
	trailingCurrencySymbol string,
	currencyInsideNumSymbol bool,
	leadingNegativeNumSign string,
	trailingNegativeNumSign string,
	numSymbolFieldPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecMechanics.lock == nil {
		nStrFmtSpecMechanics.lock = new(sync.Mutex)
	}

	nStrFmtSpecMechanics.lock.Lock()

	defer nStrFmtSpecMechanics.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecMechanics."+
			"setCurrencyNStrFmtBasic()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenLeadingCurrencySym := len(leadingCurrencySymbol)

	lenTrailingCurrencySym := len(trailingCurrencySymbol)

	lenLeadingNegativeNumSign := len(leadingNegativeNumSign)

	lenTrailingNegativeNumSign := len(trailingNegativeNumSign)

	if lenLeadingCurrencySym == 0 &&
		lenTrailingCurrencySym == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Leading and Trailing Currency Symbols are invalid!\n"+
			"Both leadingCurrencySymbol and trailingCurrencySymbol strings\n"+
			"are both empty and have a zero length\n",
			ePrefix.String())

		return err

	}

	if lenLeadingNegativeNumSign == 0 &&
		lenTrailingNegativeNumSign == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Leading and Trailing Negative Number Signs are invalid!\n"+
			"Both leadingNegativeNumSign and trailingNegativeNumSign strings\n"+
			"are both empty and have a zero length\n",
			ePrefix.String())

		return err

	}

	if numSymbolFieldPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSymbolFieldPosition' is invalid!\n"+
			"The Number Symbol Field Position is invalid. This value indicates\n"+
			"whether number signs and currency symbols are displayed inside or\n"+
			"outside of the number field.\n"+
			"'numSymbolFieldPosition'  string value = %v\n"+
			"'numSymbolFieldPosition' integer value = %v\n",
			numSymbolFieldPosition.String(),
			numSymbolFieldPosition.XValueInt(),
			ePrefix.String())

		return err
	}

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        numFieldLength,
		fieldJustification: numFieldJustification,
	}

	err = numberFieldSpec.IsValidInstanceError(
		ePrefix.XCpy(
			"numberFieldSpec"))

	if err != nil {
		return err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewStr(
		decSeparatorChars,
		ePrefix.XCpy("decSeparator<-decSeparatorChars"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewIntGroupEnum(
		intGroupingType,
		intSeparatorChars,
		ePrefix.XCpy("intSeparatorSpec<-intSeparatorChars"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	if currencyInsideNumSymbol == true {
		// Currency Symbol is Inside Negative Number Symbol

		if lenLeadingNegativeNumSign > 0 {

			err = negativeNumberSign.
				leadingNumberSymbols.AddRuneArrays(
				ePrefix.XCpy("NegLeadingNumberSymbols"+
					"<-leadingNegativeNumSign"),
				[]rune(leadingNegativeNumSign))

			if err != nil {
				return err
			}

			negativeNumberSign.leadingNumberSymbols.charSearchType =
				CharSearchType.LinearTargetStartingIndex()

			negativeNumberSign.leadingNumberFieldSymbolPosition =
				numSymbolFieldPosition
		}

		if lenLeadingCurrencySym > 0 {

			err = negativeNumberSign.
				leadingNumberSymbols.AddRuneArrays(
				ePrefix.XCpy("NegLeadingNumberSymbols<-"+
					"leadingCurrencySymbol"),
				[]rune(leadingCurrencySymbol))

			if err != nil {
				return err
			}

			negativeNumberSign.leadingNumberSymbols.charSearchType =
				CharSearchType.LinearTargetStartingIndex()
		}

		if lenTrailingCurrencySym > 0 {
			err = negativeNumberSign.
				trailingNumberSymbols.AddRuneArrays(
				ePrefix.XCpy("NegLeadingNumberSymbols<-"+
					"trailingCurrencySymbol"),
				[]rune(trailingCurrencySymbol))

			if err != nil {
				return err
			}

			negativeNumberSign.trailingNumberSymbols.charSearchType =
				CharSearchType.LinearTargetStartingIndex()
		}

		if lenTrailingNegativeNumSign > 0 {

			err = negativeNumberSign.
				trailingNumberSymbols.AddRuneArrays(
				ePrefix.XCpy("NegLeadingNumberSymbols"+
					"<-trailingNegativeNumSign"),
				[]rune(trailingNegativeNumSign))

			if err != nil {
				return err
			}

			negativeNumberSign.trailingNumberFieldSymbolPosition =
				numSymbolFieldPosition

			negativeNumberSign.trailingNumberSymbols.charSearchType =
				CharSearchType.LinearTargetStartingIndex()
		}

	} else {
		// Currency Symbol is Outside Negative Number Symbol

		if lenLeadingCurrencySym > 0 {

			err = negativeNumberSign.
				leadingNumberSymbols.AddRuneArrays(
				ePrefix.XCpy("NegLeadingNumberSymbols<-"+
					"leadingCurrencySymbol"),
				[]rune(leadingCurrencySymbol))

			if err != nil {
				return err
			}

			negativeNumberSign.leadingNumberSymbols.charSearchType =
				CharSearchType.LinearTargetStartingIndex()
		}

		if lenLeadingNegativeNumSign > 0 {

			err = negativeNumberSign.
				leadingNumberSymbols.AddRuneArrays(
				ePrefix.XCpy("NegLeadingNumberSymbols"+
					"<-leadingNegativeNumSign"),
				[]rune(leadingNegativeNumSign))

			if err != nil {
				return err
			}

			negativeNumberSign.leadingNumberFieldSymbolPosition =
				numSymbolFieldPosition

			negativeNumberSign.leadingNumberSymbols.charSearchType =
				CharSearchType.LinearTargetStartingIndex()
		}

		if lenTrailingNegativeNumSign > 0 {

			err = negativeNumberSign.
				trailingNumberSymbols.AddRuneArrays(
				ePrefix.XCpy("NegLeadingNumberSymbols"+
					"<-trailingNegativeNumSign"),
				[]rune(trailingNegativeNumSign))

			if err != nil {
				return err
			}

			negativeNumberSign.trailingNumberFieldSymbolPosition =
				numSymbolFieldPosition

			negativeNumberSign.trailingNumberSymbols.charSearchType =
				CharSearchType.LinearTargetStartingIndex()
		}

		if lenTrailingCurrencySym > 0 {

			err = negativeNumberSign.
				trailingNumberSymbols.AddRuneArrays(
				ePrefix.XCpy("NegTrailingNumberSymbols<-"+
					"trailingCurrencySymbol"),
				[]rune(trailingCurrencySymbol))

			if err != nil {
				return err
			}

			negativeNumberSign.trailingNumberSymbols.charSearchType =
				CharSearchType.LinearTargetStartingIndex()
		}

	}

	var positiveNumberSign NumStrNumberSymbolSpec

	if lenLeadingCurrencySym > 0 {

		err = positiveNumberSign.
			leadingNumberSymbols.AddRuneArrays(
			ePrefix.XCpy("PositiveLeadingNumberSymbols<-"+
				"leadingCurrencySymbol"),
			[]rune(leadingCurrencySymbol))

		if err != nil {
			return err
		}

		positiveNumberSign.leadingNumberFieldSymbolPosition =
			numSymbolFieldPosition

		positiveNumberSign.leadingNumberSymbols.charSearchType =
			CharSearchType.LinearTargetStartingIndex()

	}

	if lenTrailingCurrencySym > 0 {

		err = positiveNumberSign.
			trailingNumberSymbols.AddRuneArrays(
			ePrefix.XCpy("PositiveTrailingNumberSymbols<-"+
				"trailingCurrencySymbol"),
			[]rune(trailingCurrencySymbol))

		if err != nil {
			return err
		}

		positiveNumberSign.trailingNumberFieldSymbolPosition =
			numSymbolFieldPosition

		positiveNumberSign.trailingNumberSymbols.charSearchType =
			CharSearchType.LinearTargetStartingIndex()

	}

	var zeroNumberSign NumStrNumberSymbolSpec

	if lenLeadingCurrencySym > 0 {

		err = zeroNumberSign.
			leadingNumberSymbols.AddRuneArrays(
			ePrefix.XCpy("ZeroLeadingNumberSymbols<-"+
				"leadingCurrencySymbol"),
			[]rune(leadingCurrencySymbol))

		if err != nil {
			return err
		}

		zeroNumberSign.leadingNumberFieldSymbolPosition =
			numSymbolFieldPosition

		zeroNumberSign.leadingNumberSymbols.charSearchType =
			CharSearchType.LinearTargetStartingIndex()

	}

	if lenTrailingCurrencySym > 0 {

		err = zeroNumberSign.
			trailingNumberSymbols.AddRuneArrays(
			ePrefix.XCpy("ZeroTrailingNumberSymbols<-"+
				"trailingCurrencySymbol"),
			[]rune(trailingCurrencySymbol))

		if err != nil {
			return err
		}

		zeroNumberSign.trailingNumberFieldSymbolPosition =
			numSymbolFieldPosition

		zeroNumberSign.trailingNumberSymbols.charSearchType =
			CharSearchType.LinearTargetStartingIndex()

	}

	var numSymbols NumStrNumberSymbolGroup

	err = numSymbols.negativeNumberSign.CopyIn(
		&negativeNumberSign,
		ePrefix.XCpy(
			"numSymbols.negativeNumberSign<-"))

	if err != nil {
		return err
	}

	err = numSymbols.positiveNumberSign.CopyIn(
		&positiveNumberSign,
		ePrefix.XCpy(
			"numSymbols.positiveNumberSign<-"))

	if err != nil {
		return err
	}

	err = numSymbols.zeroNumberSign.CopyIn(
		&zeroNumberSign,
		ePrefix.XCpy(
			"numSymbols.zeroNumberSign<-"))

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		numSymbols,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

// setSignedBasicNumStr
//
// This method reconfigures the instance of
// NumStrFormatSpec passed as input parameter,
// 'numberStrFmtSpec'.
//
// The input parameters for this method provide
// additional options for multiple integer grouping
// types including India and Chinese Numbering.
//
// Negative number sign symbols may be configured as
// leading negative number signs, trailing negative
// number signs or both. The combination of leading
// and trailing negative number signs allows for the
// configuration symbols like parentheses for the
// formatting of negative numbers.
//
//	Example: (125.34)
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberStrFmtSpec			*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//
//		This instance will be configured with new
//		NumStrFormatSpec data values extracted from the
//		following input parameters.
//
//	decSeparator					string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and fractional
//		digits within a formatted Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//		Note:	Setting 'intGroupingType' to a value of
//				IntGroupingType.None() effectively turns
//				off integer separation.
//
//		Examples:
//
//			IntGroupingType.None()
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
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a leading
//		negative number sign would be a leading minus
//		sign ('-').
//
//		Another option is to configure a single
//		parenthesis ("(") to be matched by a trailing
//		negative number sign with the closing parenthesis
//		(")"). This combination would effectively enclose
//		negative numbers in parentheses.
//			Example "(125.67)"
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		The most common configuration for a trailing
//		negative number sign would be a trailing minus
//		sign ('-').
//
//		Another option is to configure a single
//		closing parenthesis (")") to be matched by a
//		leading negative number sign with the opening
//		parenthesis ("("). This combination would
//		effectively enclose negative numbers in
//		parentheses.
//			Example "(125.67)"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be
//		displayed within a number string.
//
//		If 'numFieldLength' is less than the length of
//		the numeric value string, it will be
//		automatically set equal to the length of that
//		numeric value string.
//
//		To automatically set the value of numFieldLength
//		to the string length of the numeric value, set
//		this parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//		Field Length Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field
//		length equal to or less than the length of the
//		numeric value string never use text
//		justification. In these cases, text justification
//		is completely ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//		Text Justification Examples
//
//			Example-1
//	         FieldContents String = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Center()
//				Text Field String =
//					"   1234.5678   "
//
//			Example-2
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 15
//				numFieldJustification = TxtJustify.Right()
//				Text Field String =
//					"      1234.5678"
//
//			Example-3
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = -1
//				numFieldJustification = TxtJustify.Center()
//					// Justification Ignored. Field Length
//					// Equals -1
//				Text Field String =
//					"1234.5678"
//
//			Example-4
//	         FieldContents = "1234.5678"
//				FieldContents String Length = 9
//				numFieldLength = 2
//				numFieldJustification = TxtJustify.Center()
//					// Ignored, because FieldLength Less
//					// Than FieldContents String Length.
//				Text Field String =
//					"1234.5678"
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
func (nStrFmtSpecMechanics *numStrFmtSpecMechanics) setSignedBasicNumStr(
	numberStrFmtSpec *NumStrFormatSpec,
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingNegativeNumSign []rune,
	trailingNegativeNumSign []rune,
	numFieldLength int,
	numFieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrFmtSpecMechanics.lock == nil {
		nStrFmtSpecMechanics.lock = new(sync.Mutex)
	}

	nStrFmtSpecMechanics.lock.Lock()

	defer nStrFmtSpecMechanics.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecMechanics."+
			"setSignedBasicNumStr()",
		"")

	if err != nil {
		return err
	}

	if numberStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberStrFmtSpec' is invalid!\n"+
			"'numberStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err

	}

	if len(leadingNegativeNumSign) == 0 &&
		len(trailingNegativeNumSign) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Leading and Trailing Negative Number Signs are invalid!\n"+
			"Both leadingNegativeNumSign and trailingNegativeNumSign strings\n"+
			"are both empty and have a zero length\n",
			ePrefix.String())

		return err

	}

	err = new(numStrFmtSpecNanobot).
		setNStrNumberFieldSpec(
			numberStrFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			[]rune{},
			[]rune{},
			NumFieldSymPos.InsideNumField(),
			leadingNegativeNumSign,
			trailingNegativeNumSign,
			NumFieldSymPos.InsideNumField(),
			[]rune{},
			[]rune{},
			NumFieldSymPos.InsideNumField(),
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newNumFmtSpec<-"))

	return err
}

//	setUSDefaultSignedNumStrFmtIfNeeded
//
//	Receives a pointer to an instance of NumStrFormatSpec
//	determines if the instance is invalid.
//
//	If the NumStrFormatSpec instance is valid, the method
//	will take no action and return without error.
//
//	If the NumStrFormatSpec instance is invalid, the
//	instance will be reconfigured using the US
//	(United States) default values for Signed Number
//	String Formatting Specifications.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberStrFmtSpec			*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//
//		If this instance is invalid or uninitialized,
//		it will be reconfigured with standard US
//		(United States) Signed Number String Formatting
//		Specifications.
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
//				Field Length Examples
//
//					Example-1
//			        Number String = "5672.1234567"
//						Number String Length = 12
//						fieldLength = 18
//						fieldJustification = TxtJustify.Center()
//						Text Field String =
//							"   5672.1234567   "
//
//					Example-2
//			        Number String = "5672.1234567"
//						Number String Length = 12
//						fieldLength = 18
//						fieldJustification = TxtJustify.Center()
//						Text Field String =
//							"      5672.1234567"
//
//					Example-3
//			        Number String = "5672.1234567"
//						Number String Length = 12
//						fieldLength = -1
//						fieldJustification = TxtJustify.Center()
//							// Text Justification Ignored. Field
//							// Length Equals Title Line String Length
//						Text Field String =
//							"5672.1234567"
//
//					Example-4
//			        Number String = "5672.1234567"
//						Number String Length = 12
//						fieldLength = 2
//						fieldJustification = TxtJustify.Center()
//							// Justification Ignored because Field
//							// Length Less Than Title Line String Length.
//						Text Field String =
//							"5672.1234567"
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
//
//				Text Justification Examples
//
//					Example-1
//			        Number String = "5672.1234567"
//						Number String Length = 12
//						fieldLength = 18
//						fieldJustification = TxtJustify.Center()
//						Text Field String =
//							"   5672.1234567   "
//
//					Example-2
//			        Number String = "5672.1234567"
//						Number String Length = 12
//						fieldLength = 18
//						fieldJustification = TxtJustify.Center()
//						Text Field String =
//							"      5672.1234567"
//
//					Example-3
//			        Number String = "5672.1234567"
//						Number String Length = 12
//						fieldLength = -1
//						fieldJustification = TxtJustify.Center()
//							// Text Justification Ignored. Field
//							// Length Equals Title Line String Length
//						Text Field String =
//							"5672.1234567"
//
//					Example-4
//			        Number String = "5672.1234567"
//						Number String Length = 12
//						fieldLength = 2
//						fieldJustification = TxtJustify.Center()
//							// Justification Ignored because Field
//							// Length Less Than Title Line String Length.
//						Text Field String =
//							"5672.1234567"
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
func (nStrFmtSpecMechanics *numStrFmtSpecMechanics) setUSDefaultSignedNumStrFmtIfNeeded(
	numberStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrFmtSpecMechanics.lock == nil {
		nStrFmtSpecMechanics.lock = new(sync.Mutex)
	}

	nStrFmtSpecMechanics.lock.Lock()

	defer nStrFmtSpecMechanics.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecMechanics."+
			"setUSDefaultSignedNumStrFmtIfNeeded()",
		"")

	if err != nil {
		return err
	}

	if numberStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numberStrFmtSpec' is invalid!\n"+
			"'numberStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err

	}

	isValid,
		_ := new(numStrFmtSpecElectron).
		testValidityNumStrFormatSpec(
			numberStrFmtSpec,
			nil)

	if isValid {
		return err
	}

	err = new(numStrFmtSpecNanobot).
		setSignedNStrFmtUS(
			numberStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy(
				"numberStrFmtSpec<-Signed US Format"))

	return err
}

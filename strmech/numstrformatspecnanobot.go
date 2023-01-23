package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrFmtSpecNanobot - This type provides
// helper methods for NumStrFormatSpec
type numStrFmtSpecNanobot struct {
	lock *sync.Mutex
}

// copySignedNumberFormatSpec - Copies all data from input parameter
// 'sourceSignedNumFmtSpec' to input parameter
// 'destinationSignedNumFmtSpec'. Both instances are of type
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields in
// 'destinationSignedNumFmtSpec' will be deleted and overwritten.
//
// Also, NO data validation is performed on 'sourceSignedNumFmtSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationSignedNumFmtSpec		*NumStrFormatSpec
//		A pointer to a NumStrFormatSpec instance.
//		All the member variable data fields in this object will be
//		replaced by data values copied from input parameter
//		'sourceSignedNumFmtSpec'.
//
//		'destinationSignedNumFmtSpec' is the destination for this
//		copy operation.
//
//	sourceSignedNumFmtSpec			*NumStrFormatSpec
//		A pointer to another NumStrFormatSpec
//		instance. All the member variable data values from this
//		object will be copied to corresponding member variables in
//		'destinationSignedNumFmtSpec'.
//
//		'sourceSignedNumFmtSpec' is the source for this copy
//		operation.
//
//		No data validation is performed on 'sourceSignedNumFmtSpec'.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) copySignedNumberFormatSpec(
	destinationSignedNumFmtSpec *NumStrFormatSpec,
	sourceSignedNumFmtSpec *NumStrFormatSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"copySignedNumberFormatSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationSignedNumFmtSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationSignedNumFmtSpec' is invalid!\n"+
			"'destinationSignedNumFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceSignedNumFmtSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceSignedNumFmtSpec' is invalid!\n"+
			"'sourceSignedNumFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(numStrFmtSpecElectron).
		testValidityNumStrFormatSpec(
			sourceSignedNumFmtSpec,
			ePrefix.XCpy(
				"sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	new(numStrFmtSpecAtom).empty(
		destinationSignedNumFmtSpec)

	err = destinationSignedNumFmtSpec.decSeparator.CopyIn(
		&sourceSignedNumFmtSpec.decSeparator,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.decSeparator"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.intSeparatorSpec.CopyIn(
		&sourceSignedNumFmtSpec.intSeparatorSpec,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.intSeparatorSpec"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.numberFieldSpec.CopyIn(
		&sourceSignedNumFmtSpec.numberFieldSpec,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.numberFieldSpec"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.numberSymbolsSpec.CopyIn(
		&sourceSignedNumFmtSpec.numberSymbolsSpec,
		ePrefix.XCpy(
			" destinationSignedNumFmtSpec.numberSymbols"+
				"<-sourceSignedNumFmtSpec"))

	return err
}

// isNOP
//
//	Stands for 'Is No Operation'. This method returns
//	a boolean value signaling whether the instance of
//	NumStrFormatSpec passed as input parameter
//	'numberStrFmtSpec' is engaged, valid, fully
//	configured and capable of performing number string
//	formatting operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberStrFmtSpec			*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		This instance will be reviewed and analyzed to
//		determine if it capable of performing number
//		string formatting operations.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//		If this method returns 'true', it signals
//		that the Number String Format Specification
//		('numberStrFmtSpec') is simply an empty
//		placeholder and is NOT capable of performing
//		number string formatting operations.
//
//		If this method returns 'false', it signals that
//		the 'NumStrFormatSpec' instance
//		('numberStrFmtSpec') is fully configured, valid,
//		functional and ready to perform number string
//		formatting operations.
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) isNOP(
	numberStrFmtSpec *NumStrFormatSpec) bool {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	if numberStrFmtSpec == nil {
		return true
	}

	isValid,
		_ := new(numStrFmtSpecElectron).
		testValidityNumStrFormatSpec(
			numberStrFmtSpec,
			nil)

	return !isValid
}

//	setCurrencyNStrFmtFrance
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Reconfigures 'numStrFmtSpec' using Currency Number
//	String formatting conventions typically applied in France.
//
//	Within in the European Union many, if not most, of
//	the member countries subscribe to the Number String
//	Currency formatting standards implemented by either
//	France or Germany.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the instance
//	of NumStrFormatSpec passed as input parameter,
//	'numStrFmtSpec', will be deleted and replaced by Currency
//	Number String formatting parameters typically applied in
//	France.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1:
//			123,45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2:
//		1 000 000 000
//
//	Currency Symbols
//
//	The currency symbol used in the France is the
//	trailing Euro Currency Symbol ('€').
//
//		French Example-3:
//			Positive Numeric Currency Value
//				1 000 000,00 €
//
//	Positive Numeric Values
//
//	The positive number sign is implied. No positive
//	number is applied, only the trailing Euro Currency
//	Symbol.
//
//		French Example-4:
//			Positive Numeric Currency Value
//				1 000 000,00 €
//
//	Zero Numeric Values
//
//	The zero number format has no number sign, but the
//	currency symbol is set to a trailing Euro Currency
//	Symbol.
//
//		French Example-5:
//			Zero Numeric Currency Value
//				0,00 €
//
//	Negative Numeric Values
//
//	The negative number sign is set to leading minus
//	sign ('-') and a trailing Euro Currency Symbol
//	("€").
//
//		French Example-6:
//			Negative Numeric Currency Value
//				-1 000 000,00 €
//
//	The negative signed number symbol is configured with a
//	leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign and the currency
//	symbol will be positioned inside the number field:
//
//		French Example-7:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123,45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " -123,45 €"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
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
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtFrance(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtFrance()",
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

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewFrance(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewFrenchDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var numSymbolsGroup NumStrNumberSymbolGroup

	numSymbolsGroup,
		err = new(NumStrNumberSymbolGroup).NewCurrencyDefaultsFrance(
		ePrefix.XCpy(
			"<-numSymbolsGroup"))

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		numSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	setCurrencyNStrFmtGermany
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Reconfigures the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec' using Number
//	String Currency formatting conventions typically
//	applied in the Germany.
//
//	Within in the European Union, many member countries
//	subscribe to the Number String Currency formatting
//	standards implemented by either Germany or France.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	NumStrFormatSpec input	parameter, 'numStrFmtSpec',
//	will be deleted and replaced by Number String Currency
//	formatting parameters typically applied in Germany.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//			123,45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//		1.000.000,00
//
//	Currency Symbols
//
//	The currency symbol used in the Germany is the
//	trailing Euro symbol ('€').
//
//		German Example-3:
//		1.000.000,00 €
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		German Example-4:
//		1.000.000 €
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		German Example-5:
//			0,00 €
//
//	Negative Numeric Values
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-6:
//		1.000.000- €
//
//	The negative signed number symbol is configured with a
//	trailing minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a trailing minus
//	sign ('-'). The negative number sign and the currency
//	symbol will be positioned inside the number field:
//
//		German Example-5:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: 123,45-
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " 123,45- €"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
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
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtGermany(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtGermany()",
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

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewGermany(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewGermanDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var numSymbolsGroup NumStrNumberSymbolGroup

	numSymbolsGroup,
		err = new(NumStrNumberSymbolGroup).
		NewCurrencyDefaultsGermany(
			ePrefix.XCpy(
				"numSymbolsGroup<-"))

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		numSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	setCurrencyNStrFmtUKMinusOutside
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the UK (United Kingdom).
//
//	The default UK currency symbol is a leading Pound
//	sign ('£').
//
//		UK Example:	Positive Numeric Currency Value
//			£ 123.45
//
//	The term "MinusOutside" in the method name means that
//	the Pound sign ('£') configured for negative numeric
//	values will be outside, or to the left of, the Pound
//	sign ('£').
//
//	UK Example:	Negative Numeric Currency Value
//			- £123.45
//
//	Default values will be used to configure the returned
//	instance of NumStrNumberSymbolGroup with UK Currency
//	Number formatting specifications. New data values
//	will be configured for the positive, zero and negative
//	number sign symbols as well as the currency symbol.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://learn.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	instance of NumStrFormatSpec passed as an input
//	parameter will be deleted and replaced by Currency
//	Number String formatting parameters typically
//	applied the in UK (United Kingdom).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	Decimal Separator
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		UK Example-1:
//			123.45 (The fractional digits are "45")
//
//	Integer Separator
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		UK Example-2:
//			1,000,000
//
//	Currency Symbol
//
//	The default currency symbol used in the UK is the
//	leading Pound symbol ('£').
//
//	Positive Numeric Values
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		UK Example-3:
//			Positive Numeric Currency Value
//				£ 123.45  Positive Value
//
//	Zero Numeric Values
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		UK Example-4:
//			Zero Numeric Currency Value
//				£ 0.00
//
//	Negative Numeric Values
//
//	The negative number sign is set to a leading minus
//	sign ('-').
//
//	This method will configure the Pound sign ('£')
//	such that any minus sign configured for negative
//	numeric values will be outside, or to the left of,
//	the Pound sign ('£').
//
//		UK Example-5:
//			Negative Numeric Currency Value
//				- £123.45
//
//	The negative signed number symbol is configured with a
//	leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign and the currency
//	symbol will be positioned inside the number field:
//
//		UK Example-6:
//			NumFieldSymPos.InsideNumField()
//				Number Field Length: 10
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Right Justified
//				Formatted Number String: " - £123.45"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
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
//				You can also use the abbreviated text justification
//				enumeration syntax as follows:
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
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtUKMinusOutside(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtUKMinusOutside()",
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

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewUS(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var numSymbolsGroup NumStrNumberSymbolGroup

	numSymbolsGroup,
		err = new(NumStrNumberSymbolGroup).
		NewCurrencyDefaultsUKMinusOutside(
			ePrefix.XCpy(
				"numSymbolsGroup<-"))

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		numSymbolsGroup,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	setCurrencyNStrFmtUSParen
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Negative currency values will be displayed with a
//	leading minus sign.
//
//		Negative Currency Example
//			$ -1,000,000.00
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the US (United States).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in input
//	parameter 'numStrFmtSpec' will be deleted and
//	replaced by Currency Number String formatting
//	parameters typically applied the in US (United
//	States).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//		123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United States Example-2
//				1,000,000
//
//	The currency symbol used in the United States is the
//	Dollar Sign symbol ('$').
//
//		United States Example-3
//			$ 1,000,000.00
//
//	The negative number sign is set to a leading minus
//	sign ("-").
//
//		United States Example-4
//			$ -1,000,000.00
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			$ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-6
//				$ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
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
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtUSMinus(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtUSMinus()",
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

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewUS(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' ', '-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		lock: nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' '},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' '},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
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

//	setCurrencyNStrFmtUSParen
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Negative currency values will be displayed
//	surrounded by Parentheses ('()').
//
//		Negative Currency Example
//			$ (1,000,000.00)
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the US (United States).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in input
//	parameter 'numStrFmtSpec' will be deleted and
//	replaced by Currency Number String formatting
//	parameters typically applied the in US (United
//	States).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//		123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United States Example-2
//				1,000,000
//
//	The currency symbol used in the United States is the
//	Dollar Sign symbol ('$').
//
//		United States Example-3
//			$ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United States Example-4
//			$ (1,000,000.00)
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			$ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-6
//				$ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
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
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtUSParen(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtUSParen()",
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

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewUS(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' ', '('},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{')'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		lock: nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' '},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' '},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
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

//	setNStrNumberFieldSpec
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in input
//	parameter 'signedNumFmtSpec' will be deleted and
//	replaced by values generated from the listed input
//	parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrFmtSpec							*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All  member
//		variable data fields in this object will be replaced
//		by data values configured from the input parameter
//		described below.
//
//
//	decSeparator					[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the
//		Decimal Separator Symbol or Symbols for the
//		current instance of NumStrFormatSpec.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted Number
//		String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	intSeparatorChars				[]rune
//
//		A rune array containing one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator in the United States. It is used to
//		separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In
//	  	the United States, the standard	integer digits
//	  	separator is the comma (",").
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
//		rune array and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be
//		returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the
//		type of IntegerSeparatorSpec which will be
//		returned. The enumeration IntegerGroupingType must
//		be set to one of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//	leadingPosNumSign				[]rune
//
//		A rune array containing the leading positive
//		number sign character or characters used to
//		configure a Positive Number Sign Symbol in a
//		number string with a positive numeric value.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
//
//	trailingPosNumSign				[]rune
//
//		A rune array containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " +123.45"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45+"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " -123.45"
//					Number Field Index:  01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				[]rune
//
//		A rune array containing the leading zero number
//		sign character or characters used to configure
//		Number Sign Symbols in a number string with a
//		zero numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+') and/or currency symbols ('$').
//
//	trailingZeroNumSign				[]rune
//
//		A rune array containing the trailing zero number
//		sign character or characters used to configure
//		Number Sign Symbols in a number string with a
//		zero numeric value.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	zeroNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " +123.45"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45+"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
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
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
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
//	errPrefDto						*ePref.ErrPrefixDto
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setNStrNumberFieldSpec(
	numStrFmtSpec *NumStrFormatSpec,
	decSeparatorChars []rune,
	intGroupingChars []rune,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign []rune,
	trailingZeroNumSign []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setNStrNumberFieldSpec()",
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

	signedNumFmtSpecAtom := numStrFmtSpecAtom{}

	signedNumFmtSpecAtom.empty(
		numStrFmtSpec)

	err = signedNumFmtSpecAtom.setDecimalSeparatorParams(
		numStrFmtSpec,
		decSeparatorChars,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"decSeparatorChars"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setIntegerGroupingParams(
		numStrFmtSpec,
		intGroupingChars,
		intGroupingType,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"intGroupingParams"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setPositiveNumberSign(
		numStrFmtSpec,
		leadingPosNumSign,
		positiveNumFieldSymPosition,
		trailingPosNumSign,
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"Positive Number Sign Params"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setNegativeNumberSign(
		numStrFmtSpec,
		leadingNegNumSign,
		negativeNumFieldSymPosition,
		trailingNegNumSign,
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"Negative Number Sign Params"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setZeroNumberSign(
		numStrFmtSpec,
		leadingZeroNumSign,
		zeroNumFieldSymPosition,
		trailingZeroNumSign,
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"Zero Number Sign Params"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setNumberFieldParams(
		numStrFmtSpec,
		numFieldLength,
		numFieldJustification,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"Number Field Parameters"))

	return err
}

//	setSignedNStrFmtFrance
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'signedNumFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the France.
//
//	Within in the European Union many, if not most, of
//	the member countries subscribe to the Signed Number
//	String formatting standards implemented by either
//	France or Germany.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	NumStrFormatSpec input	parameter, 'numStrFmtSpec', will
//	be deleted and replaced by Number String formatting
//	parameters typically applied in France.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		French Example-3
//		-1 000 000 000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		French Example-4
//		1 000 000 000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		French Example-5
//			0,0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
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
//	errPrefDto						*ePref.ErrPrefixDto
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setSignedNStrFmtFrance(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setSignedNStrFmtFrance()",
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

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewFrance(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewFrenchDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		leadingNumberFieldSymbolPosition:  NumFieldSymPos.InsideNumField(),
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols:             RuneArrayDto{},
		leadingNumberFieldSymbolPosition: 0,

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols:             RuneArrayDto{},
		leadingNumberFieldSymbolPosition: 0,

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
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

//	setSignedNStrFmtGermany
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'signedNumFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the Germany.
//
//	Within in the European Union, many, if not most, of
//	the member countries subscribe to the Signed Number
//	String formatting standards implemented by either
//	France or Germany.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
//	https://docs.oracle.com/cd/E19455-01/806-0169/overview-9/index.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	NumStrFormatSpec input	parameter, 'numStrFmtSpec', will
//	be deleted and replaced by Number String formatting
//	parameters typically applied in Germany.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2
//		1.000.000.000
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-3
//		1.000.000-
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		German Example-4
//		1.000.000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		German Example-5
//			0,00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
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
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setSignedNStrFmtGermany(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setSignedNStrFmtGermany()",
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

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewGermany(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewGermanDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols:             RuneArrayDto{},
		leadingNumberFieldSymbolPosition: 0,

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),
		lock:                              nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols:              RuneArrayDto{},
		leadingNumberFieldSymbolPosition:  0,
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols:              RuneArrayDto{},
		leadingNumberFieldSymbolPosition:  0,
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
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

// setSignedNStrFmtUS
//
// Deletes and resets the member variable data values
// stored in the instance of NumStrFormatSpec passed
// as input parameter 'signedNumFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the US (United States).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in the US (United States).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United States Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United States Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
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
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setSignedNStrFmtUS(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setSignedNStrFmtUS()",
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

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewUS(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		leadingNumberFieldSymbolPosition:  NumFieldSymPos.InsideNumField(),
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols:              RuneArrayDto{},
		leadingNumberFieldSymbolPosition:  0,
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols:              RuneArrayDto{},
		leadingNumberFieldSymbolPosition:  0,
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
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

// setSignedPureNStrSpec
//
// Receives an instance of NumStrFormatSpec and
// configures that instance with specifications for
// generating a pure number string.
//
// In this context, a Signed Floating Point Pure Number
// String is defined as follows:
//
//  1. A pure number string consists entirely of numeric
//     digit characters.
//
//  2. A pure number string will separate integer and
//     fractional digits with a radix point. This
//     could be, but is not limited to, a decimal point
//     ('.').
//
//  3. A pure number string will designate negative values
//     with a minus sign ('-'). This minus sign could be
//     positioned as a leading or trailing minus sign.
//
//  4. A pure number string will NOT include integer
//     separators such as commas (',') to separate
//     integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. The
//		data specifications contained in this instance
//		will be reconfigured to format and generate a
//		pure number string representation of a floating
//		point numeric value.
//
//	decSeparatorChars			string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the NumStrFormatSpec
//		parameter, 'numStrFmtSpec'.
//
//		The decimal separator is also known as the radix
//		point and is used to separate integer and
//		fractional digits within a formatted floating
//		point Number String.
//
//		In the US, UK, Australia and most of Canada, the
//		decimal separator is the period character ('.')
//		also known as the decimal point.
//
//		In France, Germany and many countries in the
//		European Union, the Decimal Separator is the
//		comma character (',').
//
//	leadingNumSymbols			bool
//
//		In Pure Number Strings, positive numeric values
//		do NOT have leading or trailing plus signs ('+').
//		Negative values on the other hand are always
//		designated by leading or trailing minus sign
//		('-').
//
//		This parameter, 'leadingNumSymbols', controls
//		the positioning of minus signs for negative
//		numeric values within a	Number String.
//
//		When set to 'true', the NumStrFormatSpec instance,
//		'numStrFmtSpec', will configure minus signs for
//		negative numbers at the beginning of, or on the
//		left side of, the numeric value. In these cases,
//		the minus sign is said to be configured as a
//		leading minus sign. This is the positioning
//		format used in the US, UK, Australia and most of
//		Canada. In addition, library functions in 'Go'
//		and other programming languages generally expect
//		leading minus signs for negative numbers.
//
//			Example Leading Minus Sign:
//				"-123.456"
//
//		When parameter 'leadingNumSymbols' is set to
//		'false', the NumStrFormatSpec instance,
//		'numStrFmtSpec', will configure minus signs for
//		negative numbers at the end of, or on the right
//		side of, the numeric value. With this positioning
//		format, the minus sign is said to be configured
//		as a trailing minus sign. This is the positioning
//		format used in France, Germany and many countries
//		in the European Union.
//
//			Example Trailing Minus Sign:
//				"123.456-"
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
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
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
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
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setSignedPureNStrSpec(
	numStrFmtSpec *NumStrFormatSpec,
	decSeparatorChars string,
	leadingNumSymbols bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setSignedPureNStrSpec()",
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

	var decSeparatorSpec DecimalSeparatorSpec

	decSeparatorSpec,
		err = new(DecimalSeparatorSpec).NewStr(
		decSeparatorChars,
		ePrefix.XCpy(
			"decSeparatorSpec<-"))

	if err != nil {

		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec =
		new(IntegerSeparatorSpec).NewNoIntegerSeparation()

	if err != nil {

		return err
	}

	var numSymbolsSpec NumStrNumberSymbolGroup

	numSymbolsSpec,
		err = new(NumStrNumberSymbolGroup).NewSimpleSignedNumber(
		leadingNumSymbols,
		ePrefix.XCpy(
			"numSymbolsSpec<-"))

	if err != nil {

		return err
	}

	var nStrNumFieldSpec NumStrNumberFieldSpec

	nStrNumFieldSpec,
		err = new(NumStrNumberFieldSpec).NewFieldSpec(
		numFieldLength,
		numFieldJustification,
		ePrefix.XCpy(
			"nStrNumFieldSpec<-"))

	if err != nil {

		return err
	}

	nStrFmtSpecAtom := numStrFmtSpecAtom{}

	err = nStrFmtSpecAtom.setDecimalSeparatorSpec(
		numStrFmtSpec,
		decSeparatorSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-"))

	if err != nil {

		return err
	}

	err = nStrFmtSpecAtom.setIntegerGroupingSpec(
		numStrFmtSpec,
		intSeparatorSpec,
		ePrefix.XCpy(
			""))

	if err != nil {

		return err
	}

	err = nStrFmtSpecAtom.setNumberFieldSpec(
		numStrFmtSpec,
		nStrNumFieldSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-nStrNumFieldSpec"))

	if err != nil {

		return err
	}

	err = nStrFmtSpecAtom.setNumberSymbolSpec(
		numStrFmtSpec,
		numSymbolsSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-numSymbolsSpec"))

	return err
}

//	setSimpleFormatSpec
//
//	Deletes all the data values and reconfigures an
//	instance of NumStrFormatSpec for currency Number
//	String formatting.
//
//	This method provides a simplified means of creating
//	type NumStrFormatSpec using default values. The
//	passed instance of NumStrFormatSpec will be
//	configured for either Currency, or Signed Number
//	formatting.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	instance of NumStrFormatSpec passed as an input
//	parameter will be deleted and replaced by with
//	either Currency or Signed Number formatting as
//	specified by the input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameters described below.
//
//	decSeparator				string
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
//	intSeparatorChars			string
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
//		string, an error will be returned.
//
//	currencySymbols				string
//
//		The symbol or symbols used to format currency.
//		This currency formatting will be configured in
//		the new instance of NumStrFormatSpec returned by
//		this method.
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrFormatSpec will configure Number Symbols on
//		the left side of the numeric value. Such Number
//		Symbols are therefore configured as leading
//		Number Symbols. This is the positioning format
//		used in the US, UK, Australia and most of Canada.
//
//		Example Number Strings:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted between
//				the currency symbol and the minus sign.
//
//		When set to 'false', the returned instance of
//		NumStrFormatSpec will configure Number Symbols on
//		the right side of the numeric value. Such Number
//		Symbols are therefore configured as trailing
//		Number Symbols. This is the positioning format
//		used in France, Germany and many countries in
//		the European Union.
//
//		Example Number Strings:
//			"123.456- €"
//
//		NOTE:	A space is automatically inserted between
//				the minus sign and the currency symbol.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
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
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
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
func (nStrFmtSpecNanobot *numStrFmtSpecNanobot) setSimpleFormatSpec(
	numStrFmtSpec *NumStrFormatSpec,
	decSeparatorChars string,
	intSeparatorChars string,
	currencySymbols string,
	leadingNumSymbols bool,
	numFieldLength int,
	numFieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrFmtSpecNanobot.lock == nil {
		nStrFmtSpecNanobot.lock = new(sync.Mutex)
	}

	nStrFmtSpecNanobot.lock.Lock()

	defer nStrFmtSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setSimpleFormatSpec()",
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

	var decSeparatorSpec DecimalSeparatorSpec

	decSeparatorSpec,
		err = new(DecimalSeparatorSpec).NewStr(
		decSeparatorChars,
		ePrefix.XCpy(
			"decSeparatorSpec<-"))

	if err != nil {

		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	if len(intSeparatorChars) > 0 {

		intSeparatorSpec,
			err = new(IntegerSeparatorSpec).NewThousands(
			intSeparatorChars,
			ePrefix.XCpy(
				"intSeparatorSpec<-"))

		if err != nil {

			return err
		}

	} else {

		intSeparatorSpec = new(IntegerSeparatorSpec).
			NewNoIntegerSeparation()
	}

	var numSymbolsSpec NumStrNumberSymbolGroup

	numSymbolsSpec,
		err = new(NumStrNumberSymbolGroup).NewSimpleCurrency(
		currencySymbols,
		leadingNumSymbols,
		ePrefix.XCpy(
			"numSymbolsSpec<-"))

	if err != nil {

		return err
	}

	var nStrNumFieldSpec NumStrNumberFieldSpec

	nStrNumFieldSpec,
		err = new(NumStrNumberFieldSpec).NewFieldSpec(
		numFieldLength,
		numFieldJustification,
		ePrefix.XCpy(
			"nStrNumFieldSpec<-"))

	if err != nil {

		return err
	}

	nStrFmtSpecAtom := numStrFmtSpecAtom{}

	err = nStrFmtSpecAtom.setDecimalSeparatorSpec(
		numStrFmtSpec,
		decSeparatorSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-"))

	if err != nil {

		return err
	}

	err = nStrFmtSpecAtom.setIntegerGroupingSpec(
		numStrFmtSpec,
		intSeparatorSpec,
		ePrefix.XCpy(
			""))

	if err != nil {

		return err
	}

	err = nStrFmtSpecAtom.setNumberFieldSpec(
		numStrFmtSpec,
		nStrNumFieldSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-nStrNumFieldSpec"))

	if err != nil {

		return err
	}

	err = nStrFmtSpecAtom.setNumberSymbolSpec(
		numStrFmtSpec,
		numSymbolsSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-numSymbolsSpec"))

	return err
}

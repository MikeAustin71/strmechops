package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrNumberSymbolSpecMechanics - This type provides
// helper methods for NumStrNumberSymbolSpec
type numStrNumberSymbolSpecMechanics struct {
	lock *sync.Mutex
}

//	setCurrencyBasic
//
//	Receives an instance of NumStrNumberSymbolSpec,
//	deletes all internal member data values and proceeds
//	to configure that instance with a basic set of
//	new currency symbol specifications.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	NumStrNumberSymbolSpec passed as input parameter
//	'currencySymbolSpecs'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencySymbolSpecs			*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		This instance will be reconfigured with the
//		basic currency symbol specifications extracted
//		from the following input parameters.
//
//	leadingCurrencySymbol		[]rune
//
//		This rune array contains a character or
//		characters which comprise the leading Currency
//		Symbol. The leading Currency Symbol will be
//		positioned at the beginning or left side of the
//		number string.
//
//			Example: $ 123.45
//
//		If a space between the currency symbol
//		and the first digit of the number string
//		is required, be sure to include the space
//		in the currency symbol rune array.
//			Example:
//				Leading Currency Symbol: "$ "
//				Formatted Number String: "$ 123.45"
//
//		If both the leading and trailing Currency Symbol
//		input parameters are empty, an error will be
//		returned.
//
//	trailingCurrencySymbol		[]rune
//
//		This rune array contains a character or
//		characters which comprise the trailing Currency
//		Symbol. The trailing Currency Symbol will be
//		positioned at the end of, or right side of, the
//		number string.
//
//			Example: 123.45 €
//
//		If a space between the last digit of the
//		number string and the currency symbol
//		is required, be sure to include the space
//		in the currency symbol string.
//			Example:
//				Trailing Currency Symbol: " €"
//				Formatted Number String: "123.45 €"
//
//	currencyInsideNumSymbol			bool
//
//		This boolean parameter determines whether the
//		currency symbol will be positioned inside or
//		outside the negative number sign symbol.
//
//		If this parameter is set to 'false', the
//		currency symbol will be positioned outside
//		the negative number sign symbol.
//
//			Example-1 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "$ -123.45"
//
//			Example-2 Outside:
//				currencyInsideNumSymbol = false
//				Number String = "  123.45- €"
//
//		If this parameter is set to 'true', the
//		currency symbol will be positioned inside
//		the negative number sign symbol.
//
//			Example - 3 Inside:
//				currencyInsideNumSymbol = true
//				Number String = " - $123.45"
//
//			Example - 4 Inside:
//				currencyInsideNumSymbol = true
//				Number String = "  123.45€ -"
//
//	numSymbolFieldPosition			NumberFieldSymbolPosition
//
//		Defines the position of the Currency symbols
//		relative to a Number Field in which a number
//		string is displayed.
//
//		Possible valid values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//
//				Example-1 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: -123.45
//					Number Sign Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$ -123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-2 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: -123.45
//					Number Sign Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45- €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				Example-3 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//
//				Example-5 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: -123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$ -  123.45"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-6 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: -123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45- €"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-7 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-8 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setCurrencyBasic(
	currencySymbolSpecs *NumStrNumberSymbolSpec,
	leadingCurrencySymbol []rune,
	trailingCurrencySymbol []rune,
	currencyInsideNumSymbol bool,
	numSymbolFieldPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setCurrencyBasic()",
		"")

	if err != nil {
		return err
	}

	if currencySymbolSpecs == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbolSpecs' is invalid!\n"+
			"'currencySymbolSpecs' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	lenLeadingCurrSym := len(leadingCurrencySymbol)

	lenTrailingCurrSym := len(trailingCurrencySymbol)

	if lenLeadingCurrSym == 0 &&
		lenTrailingCurrSym == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'leadingCurrencySymbol' and\n"+
			"'trailingCurrencySymbol' are invalid!\n"+
			"Both 'leadingCurrencySymbol' and 'trailingCurrencySymbol'\n"+
			"are empty and contain zero text characters.\n",
			ePrefix.String())

		return err
	}

	if numSymbolFieldPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSymbolFieldPosition' is invalid!\n"+
			"'numSymbolFieldPosition' string value  = '%v'\n"+
			"'numSymbolFieldPosition' integer value = '%v'\n",
			ePrefix.String(),
			numSymbolFieldPosition.String(),
			numSymbolFieldPosition.XValueInt())

		return err

	}

	var currencyNumSignRelPos CurrencyNumSignRelativePosition

	if currencyInsideNumSymbol == true {

		currencyNumSignRelPos =
			CurrNumSignRelPos.InsideNumSign()

	} else {

		currencyNumSignRelPos =
			CurrNumSignRelPos.OutsideNumSign()
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

	if lenLeadingCurrSym > 0 &&
		lenTrailingCurrSym == 0 {

		err = nStrNumSymSpecNanobot.setLeadingCurrencySymbol(
			currencySymbolSpecs,
			leadingCurrencySymbol,
			numSymbolFieldPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"currencySymbolSpecs<-"))

	} else if lenLeadingCurrSym == 0 &&
		lenTrailingCurrSym > 0 {

		err = nStrNumSymSpecNanobot.setTrailingCurrencySymbol(
			currencySymbolSpecs,
			leadingCurrencySymbol,
			numSymbolFieldPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"currencySymbolSpecs<-"))

	} else {
		// MUST BE -
		// lenLeadingCurrSym > 0 &&
		//		lenTrailingCurrSym > 0

		err = nStrNumSymSpecNanobot.setLeadingTrailingCurrencySymbols(
			currencySymbolSpecs,
			leadingCurrencySymbol,
			trailingCurrencySymbol,
			numSymbolFieldPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"currencySymbolSpecs<-"))

	}

	return err
}

// setCurrencyDefaultsEU
//
// Receives an instance of NumStrNumberSymbolSpec and
// configures it with the default European Union (EU)
// currency symbol.
//
// The default EU currency symbol is a trailing Euro
// symbol ('€').
//
//	Example:
//		1.000.000,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencySymbols				*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		This instance will be reconfigured with the
//		default European Union (EU) currency symbol.
//
//		The default EU currency symbol is a trailing
//		Euro symbol.
//
//			Example:
//				1.000.000,00 €
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setCurrencyDefaultsEU(
	currencySymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setCurrencyDefaultsEU()",
		"")

	if err != nil {
		return err
	}

	if currencySymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"'currencySymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setTrailingCurrencySymbol(
			currencySymbols,
			[]rune{' ', '€'},
			NumFieldSymPos.InsideNumField(),
			CurrNumSignRelPos.OutsideNumSign(),
			ePrefix.XCpy(
				"currencySymbols<-Trailing Euro Sign"))
}

// setCurrencyDefaultsUKMinusInside
//
// Receives an instance of NumStrNumberSymbolSpec and
// configures it with the default UK (United Kingdom)
// currency symbol. The default UK currency symbol is
// a leading pound sign ('£').
//
//	Example:
//		£ 123.45 - Positive Value
//
// This method will configure the pound sign ('£')
// such that any minus sign configured for negative
// numeric values will be inside the pound sign ('£').
//
//	Example:
//		£ -123.45 - Negative Value
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencySymbols				*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		This instance will be reconfigured with the
//		default UK (United Kingdom) currency symbol.
//
//		The default UK currency symbol is a leading
//		pound sign ('£').
//
//			Example:
//				£ 123.45 - Positive Value
//
//		This method will configure the pound sign ('£')
//		such that any minus sign configured for negative
//		numeric values will be positioned inside the
//		pound sign ('£').
//
//			Example:
//				£ -123.45 - Negative Value
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setCurrencyDefaultsUKMinusInside(
	currencySymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setCurrencyDefaultsUKMinusInside()",
		"")

	if err != nil {
		return err
	}

	if currencySymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"'currencySymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			currencySymbols,
			[]rune{'£', ' '},
			NumFieldSymPos.InsideNumField(),
			CurrNumSignRelPos.OutsideNumSign(),
			ePrefix.XCpy(
				"currencySymbols<-Leading Pound Sign"))
}

// setCurrencyDefaultsUKMinusOutside
//
// Receives an instance of NumStrNumberSymbolSpec and
// configures it with the default UK (United Kingdom)
// currency symbol. The default UK currency symbol is
// a leading pound sign ('£').
//
//	Example:
//		£ 123.45 - Positive Value
//
// This method will configure the pound sign ('£')
// such that any minus sign configured for negative
// numeric values will be outside the pound sign ('£').
//
//	Example:
//		-£ 123.45 - Negative Value
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencySymbols				*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		This instance will be reconfigured with the
//		default UK (United Kingdom) currency symbol.
//
//		The default UK currency symbol is a leading
//		pound sign ('£').
//
//			Example:
//				£ 123.45 - Positive Value
//
//		This method will configure the pound sign ('£')
//		such that any minus sign configured for negative
//		numeric values will be positioned outside the
//		pound sign ('£').
//
//			Example:
//				-£ 123.45 - Negative Value
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setCurrencyDefaultsUKMinusOutside(
	currencySymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setCurrencyDefaultsUKMinusOutside()",
		"")

	if err != nil {
		return err
	}

	if currencySymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"'currencySymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			currencySymbols,
			[]rune{'£', ' '},
			NumFieldSymPos.InsideNumField(),
			CurrNumSignRelPos.InsideNumSign(),
			ePrefix.XCpy(
				"currencySymbols<-Leading Pound Sign"))
}

// setCurrencyDefaultsUS
//
// Receives an instance of NumStrNumberSymbolSpec and
// configures it with the default US (United States)
// currency symbol. The default US currency symbol is
// a leading dollar sign.
//
//	Example:
//		$ 123.45
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencySymbols				*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		This instance will be reconfigured with the
//		default US (United States) currency symbol.
//
//		The default US currency symbol is a leading
//		dollar sign.
//
//			Example:
//				$ 123.45
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setCurrencyDefaultsUS(
	currencySymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setCurrencyDefaultsUS()",
		"")

	if err != nil {
		return err
	}

	if currencySymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"'currencySymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			currencySymbols,
			[]rune{'$', ' '},
			NumFieldSymPos.InsideNumField(),
			CurrNumSignRelPos.OutsideNumSign(),
			ePrefix.XCpy(
				"currencySymbols<-Leading Dollar Sign"))
}

//	setSignedNumSymbolsBasic
//
//	Reconfigures three NumStrNumberSymbolSpec input
//	parameters for positive, zero and negative number
//	sign symbol specifications.
//
//	Under the basic signed number symbol configuration
//	protocol, number signs for positive and zero numeric
//	values are assumed to be implicit. Consequently,
//	number sign symbols are not displayed in number
//	strings for	positive and zero numeric values.
//
//	By default, the positive and zero input parameters
//	are therefore set as empty or 'NOP' placeholder
//	instances. 'NOP' stands for 'No	Operation'.
//
//	The negative number sign symbol specification is
//	actively configured using the
//	'leadingNegativeNumSign' and 'trailingNegativeNumSign'
//	input parameters.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If both leading and trailing negative number signs
//	are required, be sure to populate both
//	'leadingNegativeNumSign' and
//	'trailingNegativeNumSign' input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the
//	NumStrNumberSymbolSpec input parameters
//	'positiveNumberSignSymbols', 'zeroNumberSignSymbols'
//	and 'negativeNumberSignSymbols'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNegativeNumSign		[]rune
//
//		This rune array contains a character or
//		characters which comprise the leading Negative
//		Number Sign Symbol. The leading Negative Number
//		Sign Symbol will be	positioned at the beginning
//		or left side of the	number string for negative
//		numeric values.
//
//			Example: -123.45
//
//		If a space between the Negative Number Sign
//		Symbol and the first digit of the number string
//		is required, be sure to include the space
//		in the 'leadingNegativeNumSign' rune array.
//
//			Example:
//				Leading Currency Symbol: "- "
//				Formatted Number String: "- 123.45"
//
//		If both the leading and trailing Negative Number
//		Sign input parameters are empty, an error will be
//		returned.
//
//	trailingNegativeNumSign		[]rune
//
//		This rune array contains a character or
//		characters which comprise the trailing Negative
//		Number Sign Symbol. The trailing Negative Number
//		Sign Symbol will be	positioned at the end or
//		the number string for negative numeric values.
//
//			Example: 123.45-
//
//		If a space between the Negative Number Sign
//		Symbol and the first digit of the number string
//		is required, be sure to include the space
//		in the 'trailingNegativeNumSign' rune array.
//
//			Example:
//				Leading Currency Symbol: " -"
//				Formatted Number String: "123.45 -"
//
//		If both the leading and trailing Negative Number
//		Sign input parameters are empty, an error will be
//		returned.
//
//	numSymbolFieldPosition		NumberFieldSymbolPosition
//
//		Defines the position of the number sign symbols
//		relative to a Number Field in which a number
//		string is displayed.
//
//		Since input parameters 'positiveNumberSignSymbols'
//		and 'zeroNumberSignSymbols' are set to empty,
//		'NOP' placeholders by default,
//		'numSymbolFieldPosition' is used exclusively to
//		configure NumStrNumberSymbolSpec instance
//		'negativeNumberSignSymbols'.
//
//		Possible valid values for 'numSymbolFieldPosition'
//		are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//
//				Example-1 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: -123.45
//					Number Sign Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$ -123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-2 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: -123.45
//					Number Sign Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45- €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				Example-3 InsideNumField:
//					Number Field Length: 9
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4 InsideNumField:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//
//				Example-5 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: -123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$ -  123.45"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-6 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: -123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45- €"
//					Number Field Index:------>01234567890
//					Total Number String Length: 11
//
//				Example-7 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "$  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-8 OutsideNumField:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//					Number Sign Symbol: None - Value is Positive
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45 €"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//
//	positiveNumberSignSymbols	*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured as an empty or 'NOP' placeholder.
//		'NOP' stands for 'No Operation.
//
//		The positive signed number symbol is configured
//		as empty or blank because under the Basic Signed
//		Number formatting protocol, positive number signs
//		are implied and not specifically displayed for
//		positive numeric values.
//
//	zeroNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured as an empty or 'NOP' placeholder.
//		'NOP' stands for 'No Operation.
//
//		The zero signed number symbol is configured
//		as empty or blank because under the Basic Signed
//		Number formatting protocol, zero number signs
//		are implied and not specifically displayed for
//		zero numeric values.
//
//	negativeNumberSignSymbols	*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		actively configured with signed number symbols
//		associated with negative numeric values.
//
//		Leading negative number sign symbols will be
//		extracted from input parameter
//		'leadingNegativeNumSign'.
//
//		Trailing negative number sign symbols will be
//		extracted from input parameter
//		'trailingNegativeNumSign'.
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setSignedNumSymbolsBasic(
	leadingNegativeNumSign []rune,
	trailingNegativeNumSign []rune,
	numSymbolFieldPosition NumberFieldSymbolPosition,
	positiveNumberSignSymbols *NumStrNumberSymbolSpec,
	zeroNumberSignSymbols *NumStrNumberSymbolSpec,
	negativeNumberSignSymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setSignedNumSymbolsBasic()",
		"")

	if err != nil {
		return err
	}

	if positiveNumberSignSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveNumberSignSymbols' is invalid!\n"+
			"'positiveNumberSignSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if zeroNumberSignSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'zeroNumberSignSymbols' is invalid!\n"+
			"'zeroNumberSignSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if negativeNumberSignSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeNumberSignSymbols' is invalid!\n"+
			"'negativeNumberSignSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	numStrNumSymSpecNanobot.empty(
		positiveNumberSignSymbols)

	numStrNumSymSpecNanobot.empty(
		zeroNumberSignSymbols)

	numStrNumSymSpecNanobot.empty(
		negativeNumberSignSymbols)

	lenLeadingNegNumSign := len(leadingNegativeNumSign)

	lenTrailingNegNumSign := len(trailingNegativeNumSign)

	if lenLeadingNegNumSign == 0 &&
		lenTrailingNegNumSign == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'leadingNegativeNumSign' and\n"+
			"'trailingNegativeNumSign' are invalid!\n"+
			"Both 'leadingNegativeNumSign' and 'trailingNegativeNumSign'\n"+
			"are empty and contain zero text characters.\n",
			ePrefix.String())

		return err
	}

	if numSymbolFieldPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSymbolFieldPosition' is invalid!\n"+
			"'numSymbolFieldPosition' string value  = '%v'\n"+
			"'numSymbolFieldPosition' integer value = '%v'\n",
			ePrefix.String(),
			numSymbolFieldPosition.String(),
			numSymbolFieldPosition.XValueInt())

		return err

	}

	nStrNumSymSpecMolecule := numStrNumberSymbolSpecMolecule{}

	if lenLeadingNegNumSign > 0 &&
		lenTrailingNegNumSign == 0 {

		err = nStrNumSymSpecMolecule.setLeadingNStrNumSymbolSpec(
			negativeNumberSignSymbols,
			leadingNegativeNumSign,
			numSymbolFieldPosition,
			ePrefix.XCpy(
				"negativeNumberSignSymbols<-"))

	} else if lenLeadingNegNumSign == 0 &&
		lenTrailingNegNumSign > 0 {

		err = nStrNumSymSpecMolecule.setTrailingNStrNumSymbolSpec(
			negativeNumberSignSymbols,
			trailingNegativeNumSign,
			numSymbolFieldPosition,
			ePrefix.XCpy(
				"negativeNumberSignSymbols<-"))

	} else {
		// MUST BE -
		// lenLeadingCurrSym > 0 &&
		//		lenTrailingCurrSym > 0

		err = nStrNumSymSpecMolecule.setLeadingNStrNumSymbolSpec(
			negativeNumberSignSymbols,
			leadingNegativeNumSign,
			numSymbolFieldPosition,
			ePrefix.XCpy(
				"negativeNumberSignSymbols<-"))

		if err != nil {
			return err
		}

		err = nStrNumSymSpecMolecule.setTrailingNStrNumSymbolSpec(
			negativeNumberSignSymbols,
			trailingNegativeNumSign,
			numSymbolFieldPosition,
			ePrefix.XCpy(
				"negativeNumberSignSymbols<-"))

	}

	return err
}

// setSignedNumSymbolsDefaultFrance
//
// Reconfigures three NumStrNumberSymbolSpec input
// parameters with default signed number symbols
// commonly applied in France.
//
// This method applies default number symbol
// specifications configuring negative numeric values
// with leading minus signs ('-').
//
//	Example: -123.34
//
// These number symbol specifications are designed to
// format number strings containing signed numeric
// values. Currency symbols ARE NOT included in these
// configured number symbol specifications.
//
// The three configured NumStrNumberSymbolSpec instances
// are therefore configured with French specifications
// for positive signed number symbols, zero value
// symbols, and negative signed number symbols.
//
// The positive signed number symbol is configured as
// empty or blank because under French formatting
// standards, positive number signs are implied and not
// specifically displayed. Therefore, no leading plus
// ('+') symbol is required.
//
// Likewise, the zero signed number symbol is also
// configured as empty or blank because under French
// formatting standards, zero numeric values have no
// number sign symbols.
//
// The negative signed number symbol is configured with a
// trailing minus sign ('-') meaning that all negative
// numeric values will be suffixed with a trailing minus
// sign ('-'). The negative number sign will be
// positioned inside the number field:
//
//	NumFieldSymPos.InsideNumField()
//		Example:
//			Number Field Length: 9
//			Numeric Value: -123.45
//			Number Symbol: trailing minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: "  -123.45"
//			Number Field Index:       012345678
//			Total Number String Length: 9
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSignSymbols	*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with positive numeric values.
//
//		The positive signed number symbol is configured
//		as empty or blank because under French formatting
//		standards, positive number signs are implied and
//		not specifically displayed. Therefore, no leading
//		plus ('+') symbol is required.
//
//	zeroNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with zero numeric values.
//
//		The zero signed number symbol is configured as
//		empty or blank because under French formatting
//		standards, zero numeric values have	no number
//		sign symbols.
//
//	negativeNumberSignSymbols	*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with negative numeric values.
//
//		Negative numeric values will be	configured with
//		trailing minus signs ('-') in accordance with
//		French number string formatting standards.
//
//			Example: 123.34-
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setSignedNumSymbolsDefaultFrance(
	positiveNumberSignSymbols *NumStrNumberSymbolSpec,
	zeroNumberSignSymbols *NumStrNumberSymbolSpec,
	negativeNumberSignSymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setSignedNumSymbolsDefaultFrance()",
		"")

	if err != nil {
		return err
	}

	if positiveNumberSignSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveNumberSignSymbols' is invalid!\n"+
			"'positiveNumberSignSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if zeroNumberSignSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'zeroNumberSignSymbols' is invalid!\n"+
			"'zeroNumberSignSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if negativeNumberSignSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeNumberSignSymbols' is invalid!\n"+
			"'negativeNumberSignSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	numStrNumSymSpecNanobot.empty(
		positiveNumberSignSymbols)

	numStrNumSymSpecNanobot.empty(
		zeroNumberSignSymbols)

	numStrNumSymSpecNanobot.empty(
		negativeNumberSignSymbols)

	err = numStrNumSymSpecNanobot.setTrailingNStrNumSymbolSpec(
		negativeNumberSignSymbols,
		[]rune{'-'},
		NumFieldSymPos.InsideNumField(),
		ePrefix.XCpy(
			"negativeNumberSignSymbols"))

	return err

}

// setSignedNumSymbolsDefaultGermany
//
// Reconfigures three NumStrNumberSymbolSpec input
// parameters with default signed number symbols
// commonly applied in Germany.
//
// This method applies default number symbol
// specifications configuring negative numeric values
// with trailing minus signs ('-').
//
//	Example: 123.34-
//
// These number symbol specifications are designed to
// format number strings containing signed numeric
// values. Currency symbols ARE NOT included in these
// configured number symbol specifications.
//
// The three configured NumStrNumberSymbolSpec instances
// are therefore configured with German specifications
// for positive signed number symbols, zero value
// symbols, and negative signed number symbols.
//
// The positive signed number symbol is configured as
// empty or blank because under German formatting
// standards, positive number signs are implied and not
// specifically displayed. Therefore, no leading plus
// ('+') symbol is required.
//
// Likewise, the zero signed number symbol is also
// configured as empty or blank because under German
// formatting standards, zero numeric values have no
// number sign symbols.
//
// The negative signed number symbol is configured with a
// trailing minus sign ('-') meaning that all negative
// numeric values will be suffixed with a trailing minus
// sign ('-'). The negative number sign will be
// positioned inside the number field:
//
//	NumFieldSymPos.InsideNumField()
//		Example:
//			Number Field Length: 9
//			Numeric Value: -123.45
//			Number Symbol: trailing minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: "  123.45-"
//			Number Field Index:       012345678
//			Total Number String Length: 9
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveSignedNumberSymbols *NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with positive numeric values.
//
//		The positive signed number symbol is configured
//		as empty or blank because under German formatting
//		standards, positive number signs are implied and
//		not specifically displayed. Therefore, no leading
//		plus ('+') symbol is required.
//
//	zeroSignedNumberSymbols		*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with zero numeric values.
//
//		The zero signed number symbol is configured as
//		empty or blank because under German formatting
//		standards, zero numeric values have	no number
//		sign symbols.
//
//	negativeSignedNumberSymbols *NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with negative numeric values.
//
//		Negative numeric values will be	configured with
//		trailing minus signs ('-') in accordance with
//		German number string formatting standards.
//
//			Example: 123.34-
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setSignedNumSymbolsDefaultGermany(
	positiveSignedNumberSymbols *NumStrNumberSymbolSpec,
	zeroSignedNumberSymbols *NumStrNumberSymbolSpec,
	negativeSignedNumberSymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setSignedNumSymbolsDefaultGermany()",
		"")

	if err != nil {
		return err
	}

	if positiveSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveSignedNumberSymbols' is invalid!\n"+
			"'positiveSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if zeroSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'zeroSignedNumberSymbols' is invalid!\n"+
			"'zeroSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if negativeSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeSignedNumberSymbols' is invalid!\n"+
			"'negativeSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	numStrNumSymSpecNanobot.empty(
		positiveSignedNumberSymbols)

	numStrNumSymSpecNanobot.empty(
		zeroSignedNumberSymbols)

	numStrNumSymSpecNanobot.empty(
		negativeSignedNumberSymbols)

	err = numStrNumSymSpecNanobot.setTrailingNStrNumSymbolSpec(
		negativeSignedNumberSymbols,
		[]rune{'-'},
		NumFieldSymPos.InsideNumField(),
		ePrefix.XCpy(
			"negativeSignedNumberSymbols"))

	return err

}

// setSignedNumSymbolsDefaultUSMinus
//
// Reconfigures three NumStrNumberSymbolSpec instances
// with default signed number symbols commonly applied in
// the United States (US).
//
// This method applies default number symbol
// specifications configuring negative numeric values
// with leading minus signs ('-').
//
//	Example: -123.34
//
// These number symbol specifications are designed to
// format number strings containing signed numeric
// values. Currency symbols ARE NOT included in these
// configured number symbol specifications.
//
// The three configured NumStrNumberSymbolSpec instances
// are therefore configured with US specifications for
// positive signed number symbols, zero value symbols,
// and negative signed number symbols.
//
// The positive signed number symbol is configured as
// empty or blank because under United States formatting
// standards, positive number signs are implied and not
// specifically displayed. Therefore, no leading plus
// ('+') symbol is required.
//
// Likewise, the zero signed number symbol is also
// configured as empty or blank because under United
// States formatting standards, zero numeric values have
// no number sign symbols.
//
// The negative signed number symbol is configured with a
// leading minus sign ('-') meaning that all negative
// numeric values will be prefixed with a leading minus
// sign ('-'). The negative number sign will be
// positioned inside the number field:
//
//	NumFieldSymPos.InsideNumField()
//		Example:
//			Number Field Length: 9
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: "  -123.45"
//			Number Field Index:       012345678
//			Total Number String Length: 9
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveSignedNumberSymbols *NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with positive numeric values.
//
//		The positive signed number symbol is configured
//		as empty or blank because under United States
//		formatting standards, positive number signs are
//		implied and not specifically displayed.
//		Therefore, no leading plus ('+') symbol is
//		required.
//
//	zeroSignedNumberSymbols		*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with zero numeric values.
//
//		The zero signed number symbol is configured as
//		empty or blank because under United States
//		formatting standards, zero numeric values have
//		no number sign symbols.
//
//	negativeSignedNumberSymbols *NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with negative numeric values.
//
//		Negative numeric values will be	configured with
//		leading minus signs ('-') in accordance with US
//		number string formatting standards.
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setSignedNumSymbolsDefaultUSMinus(
	positiveSignedNumberSymbols *NumStrNumberSymbolSpec,
	zeroSignedNumberSymbols *NumStrNumberSymbolSpec,
	negativeSignedNumberSymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setSignedNumSymbolsDefaultUSMinus()",
		"")

	if err != nil {
		return err
	}

	if positiveSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveSignedNumberSymbols' is invalid!\n"+
			"'positiveSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if zeroSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'zeroSignedNumberSymbols' is invalid!\n"+
			"'zeroSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if negativeSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeSignedNumberSymbols' is invalid!\n"+
			"'negativeSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	numStrNumSymSpecNanobot.empty(
		positiveSignedNumberSymbols)

	numStrNumSymSpecNanobot.empty(
		zeroSignedNumberSymbols)

	numStrNumSymSpecNanobot.empty(
		negativeSignedNumberSymbols)

	err = numStrNumSymSpecNanobot.setLeadingNStrNumSymbolSpec(
		negativeSignedNumberSymbols,
		[]rune{'-'},
		NumFieldSymPos.InsideNumField(),
		ePrefix.XCpy(
			"negativeSignedNumberSymbols"))

	return err
}

// setSignedNumSymbolsDefaultUSParen
//
// Reconfigures three NumStrNumberSymbolSpec instances
// with default signed number symbols commonly applied in
// the United States (US).
//
// This method applies default number symbol
// specifications configuring negative numeric values
// with leading and trailing parentheses ("()")
//
//	Example: (123.34)
//
// These number symbol specifications are designed to
// format number strings containing signed numeric
// values. Currency symbols ARE NOT included in these
// configured number symbol specifications.
//
// The three configured NumStrNumberSymbolSpec instances
// are therefore configured with US specifications for
// positive signed number symbols, zero value symbols,
// and negative signed number symbols.
//
// The positive signed number symbol is configured as
// empty or blank because under United States formatting
// standards, positive number signs are implied and not
// specifically displayed. Therefore, no leading plus
// ('+') symbol is required.
//
// Likewise, the zero signed number symbol is also
// configured as empty or blank because under United
// States formatting standards, zero numeric values have
// no number sign symbols.
//
// The negative signed number symbols are configured with
// surrounding parentheses ("()"). The negative number
// sign will be positioned inside the number field:
//
//	NumFieldSymPos.InsideNumField()
//		Example:
//			Number Field Length: 9
//			Numeric Value: -123.45
//			Number Symbol:
//				Leading and Trailing parenthesis ("()")
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " (123.45)"
//			Number Field Index:       012345678
//			Total Number String Length: 9
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveSignedNumberSymbols *NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with positive numeric values.
//
//		The positive signed number symbol is configured
//		as empty or blank because under United States
//		formatting standards, positive number signs are
//		implied and not specifically displayed.
//		Therefore, no leading plus ('+') symbol is
//		required.
//
//	zeroSignedNumberSymbols		*NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with zero numeric values.
//
//		The zero signed number symbol is configured as
//		empty or blank because under United States
//		formatting standards, zero numeric values have
//		no number sign symbols.
//
//	negativeSignedNumberSymbols *NumStrNumberSymbolSpec
//
//		This instance of NumStrNumberSymbolSpec will be
//		configured with signed number symbols associated
//		with negative numeric values.
//
//		Negative numeric values will be	configured with
//		surrounding parentheses ("()") in accordance with
//		US number string formatting standards.
//
//				Example: (123.45)
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setSignedNumSymbolsDefaultUSParen(
	positiveSignedNumberSymbols *NumStrNumberSymbolSpec,
	zeroSignedNumberSymbols *NumStrNumberSymbolSpec,
	negativeSignedNumberSymbols *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMech.lock == nil {
		nStrNumSymSpecMech.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMech.lock.Lock()

	defer nStrNumSymSpecMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMechanics."+
			"setSignedNumSymbolsDefaultUSParen()",
		"")

	if err != nil {
		return err
	}

	if positiveSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'positiveSignedNumberSymbols' is invalid!\n"+
			"'positiveSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if zeroSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'zeroSignedNumberSymbols' is invalid!\n"+
			"'zeroSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	if negativeSignedNumberSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'negativeSignedNumberSymbols' is invalid!\n"+
			"'negativeSignedNumberSymbols' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	numStrNumSymSpecNanobot.empty(
		positiveSignedNumberSymbols)

	numStrNumSymSpecNanobot.empty(
		zeroSignedNumberSymbols)

	numStrNumSymSpecNanobot.empty(
		negativeSignedNumberSymbols)

	err = numStrNumSymSpecNanobot.setLeadingNStrNumSymbolSpec(
		negativeSignedNumberSymbols,
		[]rune{'('},
		NumFieldSymPos.InsideNumField(),
		ePrefix.XCpy(
			"negativeSignedNumberSymbols"))

	if err != nil {
		return err
	}

	err = numStrNumSymSpecNanobot.setTrailingNStrNumSymbolSpec(
		negativeSignedNumberSymbols,
		[]rune{')'},
		NumFieldSymPos.InsideNumField(),
		ePrefix.XCpy(
			"negativeSignedNumberSymbols"))

	return err
}

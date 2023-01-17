package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrNumberSymbolSpecNanobot - This type provides
// helper methods for NumStrNumberSymbolSpec
type numStrNumberSymbolSpecNanobot struct {
	lock *sync.Mutex
}

// setLeadingCurrencySymbol
//
// Receives a pointer to an instance of
// NumStrNumberSymbolSpec and reconfigures that instance
// with a leading currency symbol.
//
// The Number String Number Symbol Specification type
// (NumStrNumberSymbolSpec) is designed to assist in
// formatting numeric values as number strings for
// screen displays, printing or file output.
//
// Examples of Currency Symbol characters include such
// symbols as the dollar sign ('$'), Euro sign ('€') or
// Pound sign ('£').
//
// Leading Currency Symbols are prefixed or prepended to
// the beginning of number strings containing currency
// numeric values.
//
//	Example: $125.34
//
// Currency symbol characters are intended for use in
// formatting currency numeric values displayed in number
// strings.
//
// This method will configure and store Leading
// Currency Symbols for the instance of
// NumStrNumberSymbolSpec passed as input parameter,
// 'currencySymbols'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencySymbols					*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec. The data values contained
//		in this instance will be deleted and reconfigured
//		with a leading currency symbol.
//
//	leadingCurrencySymbol			[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the instance of NumStrNumberSymbolSpec passed
//		as input parameter 'currencySymbols'.
//
//		Leading Currency Symbol characters can include
//		such symbols as the Dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//				Example: $125.34
//
//		If this parameter is submitted as an empty string,
//		an error will be returned.
//
//	currencyFieldSymbolPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Currency
//		Symbol ('leadingCurrencySymbol') relative to a
//		Number Field in which a number string is
//		displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right
//				Formatted Number String: "  $123.45"
//				Number Field Index:-------012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 11
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45  "
//				Number Field Index:-------01234567890
//				Total Number String Length: 11
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//			    Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "$  123.45"
//				Number Field Index:-------012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  "
//				Number Field Index:-------01234567890
//				Total Number String Length: 11
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	currencyNumSignRelPos			CurrencyNumSignRelativePosition
//
//		Currency Symbols have the option of being
//		positioned either inside or outside number sign
//		symbols formatted with numeric values in a
//		number string.
//
//		Examples of number sign symbols include minus
//		signs ('-'), plus signs ('+') and surrounding
//		parentheses ("()").
//
//		Parameter 'currencyNumSignRelPos' is an instance
//		of type CurrencyNumSignRelativePosition which
//		serves as an enumeration. This enumeration has
//		three possible values, only two of which are
//		valid:
//
//			CurrNumSignRelPos.None()			- Invalid
//			CurrNumSignRelPos.OutsideNumSign()	- Valid
//			CurrNumSignRelPos.InsideNumSign()	- Valid
//
//		'CurrNumSignRelPos' is global constant used to
//		abbreviate the syntax for invoking these
//		enumeration	values. The formal syntax is:
//
//			CurrencyNumSignRelativePosition(0).OutsideNumSign()
//			CurrencyNumSignRelativePosition(0).InsideNumSign()
//
//		Examples CurrNumSignRelPos.OutsideNumSign()
//				"$ -123.45"
//				"123.45- €"
//				"£ -123.45"
//
//		Examples CurrNumSignRelPos.InsideNumSign()
//
//			Examples:
//				"- $123.45"
//				"123.45€ -"
//				"- £123.45"
//
//		NumberFieldSymbolPosition Conflicts
//
//		When formatting a number string, the
//		NumberFieldSymbolPosition values for both the
//		Currency Symbol and the Number Sign Symbol
//		MUST BE EQUAL before the Currency Number Sign
//		Relative Position parameter,
//		('currencyNumSignRelPos'), will be activated
//		and applied to the number string formatting
//		algorithm.
//
//		If the NumberFieldSymbolPosition values for both
//		the	Currency Symbol and the Number Sign Symbol
//		ARE NOT EQUAL, the NumberFieldSymbolPosition
//		parameter controls and the Currency Number Sign
//		Relative Position parameter,
//		('currencyNumSignRelPos'), will be ignored.
//
//		Example:
//			-- NumberFieldSymbolPosition Values NOT EQUAL --
//
//			Number Field Length: 8
//		  	Numeric Value: -123.45
//			Minus Sign NumberFieldSymbolPosition:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//				 Number Field Index:  012345678
//			Total Number String Length: 9
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
func (nStrNumSymSpecNanobot *numStrNumberSymbolSpecNanobot) setLeadingCurrencySymbol(
	currencySymbols *NumStrNumberSymbolSpec,
	leadingCurrencySymbol []rune,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecNanobot.lock == nil {
		nStrNumSymSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymSpecNanobot.lock.Lock()

	defer nStrNumSymSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecNanobot."+
			"setLeadingCurrencySymbol()",
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

	if currencyFieldSymbolPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencyFieldSymbolPosition' is invalid!\n"+
			"'currencyFieldSymbolPosition' string value  = '%v'\n"+
			"'currencyFieldSymbolPosition' integer value = '%v'\n",
			ePrefix.String(),
			currencyFieldSymbolPosition.String(),
			currencyFieldSymbolPosition.XValueInt())

		return err

	}

	if currencyNumSignRelPos.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: The Currency Number Sign Relative Position\n"+
			"input parameter, 'currencyNumSignRelPos' is invalid!\n"+
			" currencyNumSignRelPos String Value = %v\n"+
			"currencyNumSignRelPos Integer Value = %v\n",
			ePrefix.String(),
			currencyNumSignRelPos.String(),
			currencyNumSignRelPos.XValueInt())

		return err
	}

	if len(leadingCurrencySymbol) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingCurrencySymbol' is invalid!\n"+
			"'leadingCurrencySymbol' is empty and has a character length of zero.\n",
			ePrefix.String())

		return err
	}

	nStrNumSymSpecMolecule := numStrNumberSymbolSpecMolecule{}

	nStrNumSymSpecMolecule.empty(
		currencySymbols)

	err = nStrNumSymSpecMolecule.
		setLeadingNStrNumSymbolSpec(
			currencySymbols,
			leadingCurrencySymbol,
			currencyFieldSymbolPosition,
			ePrefix.XCpy(
				"currencySymbols<-"+
					"leadingCurrencySymbol"))

	if err != nil {
		return err
	}

	currencySymbols.currencyNumSignRelativePos =
		currencyNumSignRelPos

	return err
}

// setLeadingTrailingCurrencySymbols
//
// Receives a pointer to an instance of
// NumStrNumberSymbolSpec and reconfigures that instance
// with a leading and trailing currency symbols.
//
// The Number String Number Symbol Specification type
// (NumStrNumberSymbolSpec) is designed to assist in
// formatting numeric values as number strings for
// screen displays, printing or file output.
//
// Examples of Currency Symbol characters include such
// symbols as the dollar sign ('$'), Euro sign ('€') or
// Pound sign ('£').
//
// This method will configure and store leading and
// trailing currency symbols for the instance of
// NumStrNumberSymbolSpec passed as input parameter,
// 'currencySymbols'.
//
// Leading currency symbols are prefixed or prepended
// at the beginning of a number string while trailing
// currency symbols are suffixed or appended at the end
// of a number string.
//
//	Example Leading Currency Sign : $123.45
//	Example Trailing Currency Sign: 123.45€
//
// Although most currency symbols are configured as
// either leading or trailing symbols, cases may arise
// where currency symbols are required at both ends of
// a number string. This method configures a Number
// String Number Symbol Specification
// (NumStrNumberSymbolSpec) with both Leading and
// Trailing Currency Symbols.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencySymbols					*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec. The data values contained
//		in this instance will be deleted and reconfigured
//		with both Leading and Trailing Currency symbols.
//
//	leadingCurrencySymbol			[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the instance of NumStrNumberSymbolSpec passed
//		as input parameter 'currencySymbols'.
//
//		Leading Currency Symbol characters can include
//		such symbols as the Dollar sign ('$'), Euro sign
//	 	('€') or Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//		If this parameter is submitted as an empty string,
//		an error will be returned.
//
//	trailingCurrencySymbol			[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the instance of NumStrNumberSymbolSpec passed
//		as input parameter 'currencySymbols'.
//
//		Trailing Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') or Pound sign ('£').
//
//		Trailing Currency Symbols are suffixed or
//		appended to the end of number strings containing
//		currency numeric values.
//
//		If this parameter is submitted as an empty string,
//		an error will be returned.
//
//	currencyFieldSymbolPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading and Trailing
//		Currency Symbols ('leadingCurrencySymbol' and
//		'trailingCurrencySymbol') relative to a	Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 11
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ("$")
//				Trailing Currency Symbol: Dollar sign ("$")
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right
//				Formatted Number String: "  $123.45$"
//				Number Field Index:-------0123456790
//				Total Number String Length: 11
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ("$")
//				Trailing Currency Symbol: Dollar sign ("$")
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:-------012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//			    Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ("$")
//				Trailing Currency Symbol: Dollar sign ("$")
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:-------0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ("$")
//				Trailing Currency Symbol: Dollar sign ("$")
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:-------012345678901
//				Total Number String Length: 12
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	currencyNumSignRelPos			CurrencyNumSignRelativePosition
//
//		Currency Symbols have the option of being
//		positioned either inside or outside number sign
//		symbols formatted with numeric values in a
//		number string.
//
//		Examples of number sign symbols include minus
//		signs ('-'), plus signs ('+') and surrounding
//		parentheses ("()").
//
//		The Currency Number Sign Relative Position
//		('currencyNumSignRelPos') determines whether the
//		leading and trailing currency symbols positioned
//		inside or outside number signs formatted in
//		the number string.
//
//		Parameter 'currencyNumSignRelPos' is an instance
//		of type CurrencyNumSignRelativePosition which
//		serves as an enumeration. This enumeration has
//		three possible values, only two of which are
//		valid:
//
//			CurrNumSignRelPos.None()			- Invalid
//			CurrNumSignRelPos.OutsideNumSign()	- Valid
//			CurrNumSignRelPos.InsideNumSign()	- Valid
//
//		'CurrNumSignRelPos' is global constant used to
//		abbreviate the syntax for invoking these
//		enumeration	values. The formal syntax is:
//
//			CurrencyNumSignRelativePosition(0).OutsideNumSign()
//			CurrencyNumSignRelativePosition(0).InsideNumSign()
//
//		Examples CurrNumSignRelPos.OutsideNumSign()
//				"$ -123.45$"
//				"€123.45- €"
//				"£ -123.45£"
//
//		Examples CurrNumSignRelPos.InsideNumSign()
//
//			Examples:
//				"- $123.45$"
//				"€123.45€ -"
//				"- £123.45£"
//
//		NumberFieldSymbolPosition Conflicts
//
//		When formatting a number string, the
//		NumberFieldSymbolPosition values for both the
//		Leading and Trailing Currency Symbols and the
//		Number Sign Symbol MUST BE EQUAL before the
//		Currency Number Sign Relative Position parameter,
//		('currencyNumSignRelPos'), will be activated and
//		applied to the number string formatting
//		algorithm.
//
//		If the NumberFieldSymbolPosition values for both
//		the	Currency Symbol and the Number Sign Symbol
//		ARE NOT EQUAL, the NumberFieldSymbolPosition
//		parameter controls and the Currency Number Sign
//		Relative Position parameter,
//		('currencyNumSignRelPos'), will be ignored.
//
//		Example:
//			-- NumberFieldSymbolPosition Values NOT EQUAL --
//
//			Number Field Length: 8
//		  	Numeric Value: -123.45
//			Minus Sign NumberFieldSymbolPosition:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Trailing Currency Symbol: Dollar sign ("$")
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45$"
//				 Number Field Index:  0123456789
//			Total Number String Length: 10
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
func (nStrNumSymSpecNanobot *numStrNumberSymbolSpecNanobot) setLeadingTrailingCurrencySymbols(
	currencySymbols *NumStrNumberSymbolSpec,
	leadingCurrencySymbol []rune,
	trailingCurrencySymbol []rune,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecNanobot.lock == nil {
		nStrNumSymSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymSpecNanobot.lock.Lock()

	defer nStrNumSymSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecNanobot."+
			"setLeadingTrailingCurrencySymbols()",
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

	if currencyFieldSymbolPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencyFieldSymbolPosition' is invalid!\n"+
			"'currencyFieldSymbolPosition' string value  = '%v'\n"+
			"'currencyFieldSymbolPosition' integer value = '%v'\n",
			ePrefix.String(),
			currencyFieldSymbolPosition.String(),
			currencyFieldSymbolPosition.XValueInt())

		return err

	}

	if currencyNumSignRelPos.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: The Currency Number Sign Relative Position\n"+
			"input parameter, 'currencyNumSignRelPos' is invalid!\n"+
			" currencyNumSignRelPos String Value = %v\n"+
			"currencyNumSignRelPos Integer Value = %v\n",
			ePrefix.String(),
			currencyNumSignRelPos.String(),
			currencyNumSignRelPos.XValueInt())

		return err
	}

	if len(leadingCurrencySymbol) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingCurrencySymbol' is invalid!\n"+
			"'leadingCurrencySymbol' is empty and has a character length of zero.\n",
			ePrefix.String())

		return err
	}

	if len(trailingCurrencySymbol) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingCurrencySymbol' is invalid!\n"+
			"'trailingCurrencySymbol' is empty and has a character length of zero.\n",
			ePrefix.String())

		return err
	}

	nStrNumSymSpecMolecule := numStrNumberSymbolSpecMolecule{}

	nStrNumSymSpecMolecule.empty(
		currencySymbols)

	err = nStrNumSymSpecMolecule.
		setLeadingNStrNumSymbolSpec(
			currencySymbols,
			leadingCurrencySymbol,
			currencyFieldSymbolPosition,
			ePrefix.XCpy(
				"currencySymbols<-"+
					"leadingCurrencySymbol"))

	if err != nil {
		return err
	}

	err = nStrNumSymSpecMolecule.
		setTrailingNStrNumSymbolSpec(
			currencySymbols,
			trailingCurrencySymbol,
			currencyFieldSymbolPosition,
			ePrefix.XCpy(
				"currencySymbols<-"+
					"trailingCurrencySymbol"))

	if err != nil {
		return err
	}

	currencySymbols.currencyNumSignRelativePos =
		currencyNumSignRelPos

	return err
}

// setTrailingCurrencySymbol
//
// Receives a pointer to an instance of
// NumStrNumberSymbolSpec and reconfigures that instance
// with a trailing currency symbol.
//
// The Number String Number Symbol Specification type
// (NumStrNumberSymbolSpec) is designed to assist in
// formatting numeric values as number strings for
// screen displays, printing or file output.
//
// Trailing Currency Symbols are suffixed or appended
// to the end of a number string.
//
//	Example: 125.34€
//
// This method will configure and store Trailing
// Currency Symbols for the instance of
// NumStrNumberSymbolSpec passed as input parameter,
// 'currencySymbols'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	currencySymbols					*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec. The data values contained
//		in this instance will be deleted and reconfigured
//		with a trailing currency symbol.
//
//	trailingCurrencySymbol			[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the instance of NumStrNumberSymbolSpec passed
//		as input parameter 'currencySymbols'.
//
//		Trailing Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') or Pound sign ('£').
//
//		Trailing Currency Symbols are suffixed and
//		appended to the end of number strings containing
//		currency numeric values.
//
//				Example: 125.34€
//
//		If this parameter is submitted as an empty string,
//		an error will be returned.
//
//	currencyFieldSymbolPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Currency
//		Symbol ('trailingCurrencySymbol') relative to a
//		Number Field in which a number string is
//		displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Euro sign ('€')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right
//				Formatted Number String: "  123.45€"
//				Number Field Index:-------012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 11
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  123.45€  "
//				Number Field Index:-------01234567890
//				Total Number String Length: 11
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//			    Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "  123.45€"
//				Number Field Index:-------012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "  123.45  €"
//				Number Field Index:-------01234567890
//				Total Number String Length: 11
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	currencyNumSignRelPos			CurrencyNumSignRelativePosition
//
//		Currency Symbols have the option of being
//		positioned either inside or outside number sign
//		symbols formatted with numeric values in a
//		number string.
//
//		Examples of number sign symbols include minus
//		signs ('-'), plus signs ('+') and surrounding
//		parentheses ("()").
//
//		Parameter 'currencyNumSignRelPos' is an instance
//		of type CurrencyNumSignRelativePosition which
//		serves as an enumeration. This enumeration has
//		three possible values, only two of which are
//		valid:
//
//			CurrNumSignRelPos.None()			- Invalid
//			CurrNumSignRelPos.OutsideNumSign()	- Valid
//			CurrNumSignRelPos.InsideNumSign()	- Valid
//
//		'CurrNumSignRelPos' is global constant used to
//		abbreviate the syntax for invoking these
//		enumeration	values. The formal syntax is:
//
//			CurrencyNumSignRelativePosition(0).OutsideNumSign()
//			CurrencyNumSignRelativePosition(0).InsideNumSign()
//
//		Examples CurrNumSignRelPos.OutsideNumSign()
//				"123.45- €"
//
//		Examples CurrNumSignRelPos.InsideNumSign()
//
//			Examples:
//				"123.45€ -"
//
//		NumberFieldSymbolPosition Conflicts
//
//		When formatting a number string, the
//		NumberFieldSymbolPosition values for both the
//		Currency Symbol and the Number Sign Symbol
//		MUST BE EQUAL before the Currency Number Sign
//		Relative Position parameter,
//		('currencyNumSignRelPos'), will be activated
//		and applied to the number string formatting
//		algorithm.
//
//		If the NumberFieldSymbolPosition values for both
//		the	Currency Symbol and the Number Sign Symbol
//		ARE NOT EQUAL, the NumberFieldSymbolPosition
//		parameter controls and the Currency Number Sign
//		Relative Position parameter,
//		('currencyNumSignRelPos'), will be ignored.
//
//		Example:
//			-- NumberFieldSymbolPosition Values NOT EQUAL --
//
//			Number Field Length: 8
//		  	Numeric Value: 123.45-
//			Minus Sign NumberFieldSymbolPosition:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Trailing Currency Symbol: Euro sign ('€')
//			Number Text Justification: Right
//			Formatted Number String: " 123.45-€"
//				 Number Field Index:  012345678
//			Total Number String Length: 9
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
func (nStrNumSymSpecNanobot *numStrNumberSymbolSpecNanobot) setTrailingCurrencySymbol(
	currencySymbols *NumStrNumberSymbolSpec,
	trailingCurrencySymbol []rune,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecNanobot.lock == nil {
		nStrNumSymSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymSpecNanobot.lock.Lock()

	defer nStrNumSymSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecNanobot."+
			"setTrailingCurrencySymbol()",
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

	if currencyFieldSymbolPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencyFieldSymbolPosition' is invalid!\n"+
			"'currencyFieldSymbolPosition' string value  = '%v'\n"+
			"'currencyFieldSymbolPosition' integer value = '%v'\n",
			ePrefix.String(),
			currencyFieldSymbolPosition.String(),
			currencyFieldSymbolPosition.XValueInt())

		return err

	}

	if currencyNumSignRelPos.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: The Currency Number Sign Relative Position\n"+
			"input parameter, 'currencyNumSignRelPos' is invalid!\n"+
			" currencyNumSignRelPos String Value = %v\n"+
			"currencyNumSignRelPos Integer Value = %v\n",
			ePrefix.String(),
			currencyNumSignRelPos.String(),
			currencyNumSignRelPos.XValueInt())

		return err
	}

	if len(trailingCurrencySymbol) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingCurrencySymbol' is invalid!\n"+
			"'trailingCurrencySymbol' is empty and has a character length of zero.\n",
			ePrefix.String())

		return err
	}

	nStrNumSymSpecMolecule := numStrNumberSymbolSpecMolecule{}

	nStrNumSymSpecMolecule.empty(
		currencySymbols)

	err = nStrNumSymSpecMolecule.
		setTrailingNStrNumSymbolSpec(
			currencySymbols,
			trailingCurrencySymbol,
			currencyFieldSymbolPosition,
			ePrefix.XCpy(
				"currencySymbols<-"+
					"trailingCurrencySymbol"))

	if err != nil {
		return err
	}

	currencySymbols.currencyNumSignRelativePos =
		currencyNumSignRelPos

	return err
}

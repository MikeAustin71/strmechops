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
	currencyInsideNumSymbol bool,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
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

	return nStrNumSymSpecMolecule.
		setLeadingCurrencySymbolSpec(
			currencySymbols,
			leadingCurrencySymbol,
			currencyInsideNumSymbol,
			currencyFieldSymbolPosition,
			ePrefix.XCpy(
				"currencySymbols<-"+
					"leadingCurrencySymbol"))
}

//	setLeadingTrailingCurrencySymbols
//
//	Receives a pointer to an instance of
//	NumStrNumberSymbolSpec and reconfigures that instance
//	with a leading and trailing currency symbols.
//
//	The Number String Number Symbol Specification type
//	(NumStrNumberSymbolSpec) is designed to assist in
//	formatting numeric values as number strings for
//	screen displays, printing or file output.
//
//	Examples of Currency Symbol characters include such
//	symbols as the dollar sign ('$'), Euro sign ('€') or
//	Pound sign ('£').
//
//	This method will configure and store leading and
//	trailing currency symbols for the instance of
//	NumStrNumberSymbolSpec passed as input parameter,
//	'currencySymbols'.
//
//	Leading currency symbols are prefixed or prepended
//	at the beginning of a number string while trailing
//	currency symbols are suffixed or appended at the end
//	of a number string.
//
//		Example Leading Currency Sign : $123.45
//		Example Trailing Currency Sign: 123.45€
//
//	Although most currency symbols are configured as
//	either leading or trailing symbols, cases may arise
//	where currency symbols are required at both ends of
//	a number string. This method configures a Number
//	String Number Symbol Specification
//	(NumStrNumberSymbolSpec) with both Leading and
//	Trailing Currency Symbols.
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
	currencyInsideNumSymbol bool,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
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

	lenLeadingCurrSym := len(leadingCurrencySymbol)

	lenTrailingCurrSym := len(trailingCurrencySymbol)

	if lenLeadingCurrSym == 0 &&
		lenTrailingCurrSym == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingCurrencySymbol' and\n"+
			"'trailingCurrencySymbol' are invalid!\n"+
			"Both 'leadingCurrencySymbol' and 'trailingCurrencySymbol'\n"+
			" are empty and have zero text characters.\n",
			ePrefix.String())

		return err
	}

	nStrNumSymSpecMolecule := numStrNumberSymbolSpecMolecule{}

	nStrNumSymSpecMolecule.empty(
		currencySymbols)

	if lenLeadingCurrSym > 0 &&
		lenTrailingCurrSym == 0 {

		err = nStrNumSymSpecMolecule.
			setLeadingCurrencySymbolSpec(
				currencySymbols,
				leadingCurrencySymbol,
				currencyInsideNumSymbol,
				currencyFieldSymbolPosition,
				ePrefix.XCpy(
					"currencySymbols<-"+
						"leadingCurrencySymbol"))

	} else if lenLeadingCurrSym == 0 &&
		lenTrailingCurrSym > 0 {

		err = nStrNumSymSpecMolecule.
			setTrailingCurrencySymbolSpec(
				currencySymbols,
				trailingCurrencySymbol,
				currencyInsideNumSymbol,
				currencyFieldSymbolPosition,
				ePrefix.XCpy(
					"currencySymbols<-"+
						"trailingCurrencySymbol"))
	} else {
		// MUST BE -
		//  lenLeadingCurrSym > 0 &&
		//		lenTrailingCurrSym > 0

		err = nStrNumSymSpecMolecule.
			setLeadingCurrencySymbolSpec(
				currencySymbols,
				leadingCurrencySymbol,
				currencyInsideNumSymbol,
				currencyFieldSymbolPosition,
				ePrefix.XCpy(
					"currencySymbols<-"+
						"leadingCurrencySymbol"))

		if err != nil {
			return err
		}

		err = nStrNumSymSpecMolecule.
			setTrailingCurrencySymbolSpec(
				currencySymbols,
				trailingCurrencySymbol,
				currencyInsideNumSymbol,
				currencyFieldSymbolPosition,
				ePrefix.XCpy(
					"currencySymbols<-"+
						"trailingCurrencySymbol"))

	}

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
	currencyInsideNumSymbol bool,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
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
		setTrailingCurrencySymbolSpec(
			currencySymbols,
			trailingCurrencySymbol,
			currencyInsideNumSymbol,
			currencyFieldSymbolPosition,
			ePrefix.XCpy(
				"currencySymbols<-"+
					"trailingCurrencySymbol"))

	if err != nil {
		return err
	}

	return err
}

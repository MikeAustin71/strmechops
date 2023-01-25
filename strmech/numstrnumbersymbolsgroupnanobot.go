package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrNumberSymbolGroupNanobot
//
// Provides helper methods for NumStrNumberSymbolGroup.
type numStrNumberSymbolGroupNanobot struct {
	lock *sync.Mutex
}

//	empty
//
//	Deletes and resets all data values contained in
//	paramter 'nStrNumSymbols' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data contained in the Positive, Negative and
//	Zero Number Symbol Specification will be deleted
//	and reset to their zero or uninitialized states.
//	The affected member variables are identified as
//	follows:
//
//		nStrNumSymbols.positiveNumberSign
//		nStrNumSymbols.negativeNumberSign
//		nStrNumSymbols.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) empty(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.positiveNumberSign.Empty()

	nStrNumSymbols.negativeNumberSign.Empty()

	nStrNumSymbols.zeroNumberSign.Empty()

	nStrNumSymbols.currencySymbol.Empty()
}

//	equal
//
//	Receives a pointer to two instances of
//	NumStrNumberSymbolGroup and proceeds to compare their
//	member variables in order to determine if they are
//	equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables for both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nNumSymbols1			*NumStrNumberSymbolGroup
//
//		An instance of NumStrNumberSymbolGroup. Internal
//		member variables from 'nNumSymbols1' will be
//		compared to those of 'nNumSymbols2' to
//		determine if both instances are equivalent.
//
//
//	nNumSymbols2			*NumStrNumberSymbolGroup
//
//		An instance of NumStrNumberSymbolGroup. Internal
//		member variables from 'nNumSymbols1' will
//		be compared to those of 'nNumSymbols2' to
//		determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the comparison of 'nNumSymbols1' and
//		'nNumSymbols2' shows that all internal
//		member variables are equivalent, this method
//		will return a boolean value of 'true'.
//
//		If the two instances are NOT equal, this method
//		will return a boolean value of 'false' to the
//		calling function.
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) equal(
	nNumSymbols1 *NumStrNumberSymbolGroup,
	nNumSymbols2 *NumStrNumberSymbolGroup) bool {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	if nNumSymbols1 == nil ||
		nNumSymbols2 == nil {

		return false
	}

	if !nNumSymbols1.negativeNumberSign.Equal(
		&nNumSymbols2.negativeNumberSign) {

		return false
	}

	if !nNumSymbols1.positiveNumberSign.Equal(
		&nNumSymbols2.positiveNumberSign) {

		return false
	}

	if !nNumSymbols1.zeroNumberSign.Equal(
		&nNumSymbols2.zeroNumberSign) {

		return false
	}

	if !nNumSymbols1.currencySymbol.Equal(
		&nNumSymbols2.currencySymbol) {

		return false
	}

	return true
}

//	setCurrencySymbolRunes
//
//	Receives two rune arrays used to configure leading
//	currency symbols, trailing currency symbols or
//	leading and trailing currency symbols.
//
//	Examples of Currency Symbol characters include such
//	symbols as the Dollar sign ('$'), Euro sign ('€') or
//	Pound sign ('£').
//
//	Input parameter 'nStrNumSymbols' is an instance of
//	NumStrNumberSymbolGroup which will be configured with
//	new data values for Currency Symbols.
//
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method configures currency symbols to be
//	displayed with positive, negative and zero numeric
//	values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Currency Symbol member variable data fields for
//	'nStrNumSymbols' according to the Currency input
//	parameters.
//
//	The nStrNumSymbols Currency Symbol member variable
//	is:
//
//		nStrNumSymbols.currencySymbols
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols					*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		The Negative Number Sign Symbol Specifications
//		for this instance will be deleted and reset to
//		the values provided by the following input
//		parameters.
//
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the NumStrNumberSymbolSpec instance,
//		'nStrNumSymbols'.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//		If 'leadingCurrencySymbol' and
//		'trailingCurrencySymbols' parameters are both
//		submitted as empty rune arrays with zero text
//		characters, an error will be returned.
//
//	trailingCurrencySymbol			[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the NumStrNumberSymbolSpec instance,
//		'nStrNumSymbols'.
//
//		Trailing Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Trailing Currency Symbols are suffixed or
//		appended to the end of number strings containing
//		currency numeric values.
//
//				Example: 125.34€
//
//		If the 'leadingCurrencySymbol' and
//		'trailingCurrencySymbols' parameters are both
//		submitted as empty rune arrays with zero text
//		characters, an error will be returned.
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
//	currencyNumFieldSymPosition 	NumberFieldSymbolPosition
//
//		Defines the position of the Currency Symbol
//		('leadingCurrencySymbol' and 'trailingCurrencySymbols' ) relative to a
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
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 11
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45  "
//				Number Field Index:------>01234567890
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
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  "
//				Number Field Index:------>01234567890
//				Total Number String Length: 11
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
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
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) setCurrencySymbolRunes(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	leadingCurrencySymbols []rune,
	trailingCurrencySymbol []rune,
	currencyInsideNumSymbol bool,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupNanobot."+
			"setCurrencySymbolRunes()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbols' is invalid!\n"+
			"'nStrNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenLeadingCurrSyms := len(leadingCurrencySymbols)

	lenTrailingCurrSyms := len(trailingCurrencySymbol)

	if lenLeadingCurrSyms == 0 &&
		lenTrailingCurrSyms == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'leadingCurrencySymbols' and\n"+
			"'trailingCurrencySymbol' are invalid!\n"+
			"Both parameters are empty and contain zero text characters.\n",
			ePrefix.String())

		return err

	}

	if lenLeadingCurrSyms > 0 &&
		lenTrailingCurrSyms > 0 {

		err = nStrNumSymbols.currencySymbol.
			SetCurrencyLeadingTrailingSymbolRunes(
				leadingCurrencySymbols,
				trailingCurrencySymbol,
				currencyInsideNumSymbol,
				currencyNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.currencySymbol<-"))

	} else if lenLeadingCurrSyms > 0 &&
		lenTrailingCurrSyms == 0 {

		err = nStrNumSymbols.currencySymbol.
			SetCurrencyLeadingSymbolRunes(
				leadingCurrencySymbols,
				currencyInsideNumSymbol,
				currencyNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.currencySymbol<-"))

	} else {
		// MUST BE -
		//  lenLeadingCurrSyms == 0  &&
		//		lenTrailingCurrSyms > 0

		err = nStrNumSymbols.currencySymbol.
			SetCurrencyTrailingSymbolRunes(
				trailingCurrencySymbol,
				currencyInsideNumSymbol,
				currencyNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.currencySymbol<-"))
	}

	return err
}

//	setCurrencySymbolSpec
//
//	Receives a pointer to an instance of
//	NumStrNumberSymbolSpec and proceeds to delete and
//	reconfigure the internal Currency Symbols member
//	variable:
//
//		NumStrNumberSymbolSpec.currencySymbol
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Currency Symbols member variable data fields
//	contained in input paramter, 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		The Currency Symbol Specifications for this
//		instance will be deleted and reset to the values
//		provided by input parameter 'currencySymbolSpec'.
//
//	currencySymbolSpec			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec containing
//		the Currency Symbol format specifications which
//		will be copied to the Currency Symbol member
//		variable contained in the NumStrNumberSymbolGroup
//		instance, 'nStrNumSymbols'.
//
//		If this value is empty, np error will be returned
//		and the 'nStrNumSymbols' will be set to a NOP.
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
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) setCurrencySymbolSpec(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	currencySymbolSpec NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupNanobot."+
			"setNegativeNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbols' is invalid!\n"+
			"'nStrNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if currencySymbolSpec.IsNOP() {

		nStrNumSymbols.currencySymbol.SetNOP()

		return err
	}

	err = nStrNumSymbols.currencySymbol.CopyIn(
		&currencySymbolSpec,
		ePrefix.XCpy(
			"nStrNumSymbols.currencySymbol"))

	return err
}

//	setNegativeNumSignRunes
//
//	Receives two rune arrays used to configure
//	leading negative number signs, trailing negative
//	number signs or leading and trailing negative number
//	signs.
//
//	Input parameter 'nStrNumSymbols' is an instance of
//	NumStrNumberSymbolGroup. The Negative Number data
//	in this instance will be configured with new
//	data values for Negative Number Sign Symbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Negative Number Sign Symbol member variable data fields
//	as specified by the rune arrays passed as input
//	parameters.
//
//		nStrNumSymbols.negativeNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols					*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		The Negative Number Sign Symbol Specifications
//		for this instance will be deleted and reset to
//		the values provided by the following input
//		parameters.
//
//	leadingNegativeNumberSymbols	[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') or leading and trailing parentheses ('()').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	trailingNegativeNumberSymbols	[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') or leading and trailing parentheses ('()').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed.
//
//		A Number Field defines the length of a text
//		string in which a numeric value is formatted.
//		When applied, a Number Field is usually longer
//		than the numeric value string and typically
//		justifies that numeric value string as "Left",
//		"Right" or "Centered".
//
//			Number Field Example (Centered):
//				"  123.45  "
//
//
//		Possible valid 'negativeNumFieldSymPosition'
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
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
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
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
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) setNegativeNumSignRunes(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	leadingNegativeNumberSymbols []rune,
	trailingNegativeNumberSymbols []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupNanobot."+
			"setNegativeNumSignRunes()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbols' is invalid!\n"+
			"'nStrNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenLeadingNumSyms := len(leadingNegativeNumberSymbols)

	lenTrailingNumSyms := len(trailingNegativeNumberSymbols)

	if lenLeadingNumSyms == 0 &&
		lenTrailingNumSyms == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'leadingNegativeNumberSymbols' and\n"+
			"'trailingNegativeNumberSymbols' are invalid!\n"+
			"Both parameters are empty and contain zero text characters.\n",
			ePrefix.String())

		return err

	}

	if lenLeadingNumSyms > 0 &&
		lenTrailingNumSyms > 0 {

		err = nStrNumSymbols.negativeNumberSign.
			SetNumberSignLeadingTrailingSymbolRunes(
				leadingNegativeNumberSymbols,
				negativeNumFieldSymPosition,
				trailingNegativeNumberSymbols,
				negativeNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.negativeNumberSign"))

	} else if lenLeadingNumSyms > 0 &&
		lenTrailingNumSyms == 0 {

		err = nStrNumSymbols.negativeNumberSign.
			SetNumberSignLeadingSymbolRunes(
				leadingNegativeNumberSymbols,
				negativeNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.negativeNumberSign"))

	} else {
		// MUST BE -
		//  lenLeadingNumSyms == 0  &&
		//		lenTrailingNumSyms > 0

		err = nStrNumSymbols.negativeNumberSign.
			SetNumberSignTrailingSymbolRunes(
				trailingNegativeNumberSymbols,
				negativeNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.negativeNumberSign"))

	}

	return err
}

//	setNegativeNumSignSpec
//
//	Receives a single NumStrNumberSymbolSpec object
//	configured as a Zero Number Sign Symbol.
//
//	This method then proceeds to reset the
//	corresponding Negative Number Sign Symbol member
//	variable data value for the NumStrNumberSymbolGroup
//	input paramter 'nStrNumSymbols'.
//
//	The Negative Number Sign Symbol internal member
//	variable is:
//
//		nStrNumSymbols.negativeNumberSign
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Be advised that this method will delete and reset
//		the Negative Number Sign Symbol member variable
//		data fields contained in input paramter,
//		'nStrNumSymbols':
//
//			nStrNumSymbols.negativeNumberSign
//
//	(2)	If input parameter 'negativeNumberSign' is
//		invalid or a NOP, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		The Negative Number Sign Symbol Specifications
//		for this instance will be deleted and reset to
//		the values provided by input parameter
//		'negativeNumberSign'.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Negative Number Sign Symbol Specification
//		will be copied to the corresponding Negative
//		Symbol Specification in input paramter,
//		'nStrNumSymbols':
//
//			nStrNumSymbols.negativeNumberSign
//
//		If input parameter 'negativeNumberSign' is
//		invalid or a NOP, an error will be returned.
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
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) setNegativeNumSignSpec(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	negativeNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupNanobot."+
			"setNegativeNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbols' is invalid!\n"+
			"'nStrNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = negativeNumberSign.IsValidInstanceError(
		ePrefix.XCpy("nStrNumSymbols.negativeNumberSign"))

	if err != nil {

		err2 := fmt.Errorf("numStrNumberSymbolGroupNanobot"+
			"setNegativeNumSignSpec()\n"+
			"Error: Input parameter 'negativeNumberSign' is invalid.\n\n"+
			"%v\n",
			err.Error())

		return err2
	}

	err = nStrNumSymbols.negativeNumberSign.CopyIn(
		&negativeNumberSign,
		ePrefix.XCpy(
			"nStrNumSymbols.negativeNumberSign<-"+
				"negativeNumberSign"))

	return err
}

//	setPositiveNumSignRunes
//
//	Receives a series of rune arrays used to configure
//	the input parameter 'nStrNumSymbols' with new data
//	values for the Positive Number Sign Symbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Positive Number Sign Symbol member variable data fields
//	as specified by the rune arrays passed as input
//	parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols					*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		The Positive Number Sign Symbol Specifications
//		for this instance will be deleted and reset to
//		the values provided by the following input
//		parameters.
//
//	leadingPositiveNumberSymbols	[]rune
//
//		A rune array containing the leading positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "$+"
//			Number String:   "$+123.456"
//
//		Example-3: Leading Number Symbols
//			Leading Number Symbols for Positive Values
//
//			Leading Symbols: "$"
//			Number String:   "$123.456"
//
//	trailingPositiveNumberSymbols	[]rune
//
//		A rune array containing the trailing positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "+$"
//			Number String:   "123.456+$"
//
//		Example-3: Trailing Number Symbols
//			Trailing Number Symbols for Positive Values
//
//			Trailing Symbols: "$"
//			Number String:   "123.456$"
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
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
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
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
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
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
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) setPositiveNumSignRunes(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	leadingPositiveNumberSymbols []rune,
	trailingPositiveNumberSymbols []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupNanobot."+
			"setPositiveNumSignRunes()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbols' is invalid!\n"+
			"'nStrNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenLeadingNumSyms := len(leadingPositiveNumberSymbols)

	lenTrailingNumSyms := len(trailingPositiveNumberSymbols)

	if lenLeadingNumSyms == 0 &&
		lenTrailingNumSyms == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'leadingPositiveNumberSymbols' and\n"+
			"'trailingPositiveNumberSymbols' are invalid!\n"+
			"Both parameters are empty and contain zero text characters.\n",
			ePrefix.String())

		return err

	}

	if lenLeadingNumSyms > 0 &&
		lenTrailingNumSyms > 0 {

		err = nStrNumSymbols.positiveNumberSign.
			SetNumberSignLeadingTrailingSymbolRunes(
				leadingPositiveNumberSymbols,
				positiveNumFieldSymPosition,
				trailingPositiveNumberSymbols,
				positiveNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.positiveNumberSign"))

	} else if lenLeadingNumSyms > 0 &&
		lenTrailingNumSyms == 0 {

		err = nStrNumSymbols.positiveNumberSign.
			SetNumberSignLeadingSymbolRunes(
				leadingPositiveNumberSymbols,
				positiveNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.positiveNumberSign"))

	} else {
		// MUST BE -
		//  lenLeadingNumSyms == 0  &&
		//		lenTrailingNumSyms > 0

		err = nStrNumSymbols.positiveNumberSign.
			SetNumberSignTrailingSymbolRunes(
				trailingPositiveNumberSymbols,
				positiveNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.positiveNumberSign"))

	}

	return err
}

//	setPositiveNumSignSpec
//
//	Receives a single NumStrNumberSymbolSpec object
//	configured as a Positive Number Sign Symbol.
//
//	This method then proceeds to reset the corresponding
//	Positive Number Sign Symbol internal member variable
//	data value for the NumStrNumberSymbolGroup input
//	paramter 'nStrNumSymbols'.
//
//	The Positive Number Sign Symbol internal member
//	variable is:
//
//		nStrNumSymbols.positiveNumberSign
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Be advised that this method will delete and reset
//		the Positive Number Sign Symbol member variable
//		data fields contained in input paramter,
//		'nStrNumSymbols':
//
//			nStrNumSymbols.positiveNumberSign
//
//	(2)	If input parameter 'positiveNumberSign' is invalid
//		or a NOP, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		The Positive Number Sign Symbol Specification
//		for this instance will be deleted and reset to
//		the values provided by input parameter
//		'positiveNumberSign'.
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Positive Number Sign Symbol Specification
//		will be copied to the corresponding Positive
//		Symbol Specification in input paramter,
//		'nStrNumSymbols':
//
//			nStrNumSymbols.zeroNumberSign
//
//		If input parameter 'positiveNumberSign' is
//		invalid or a NOP, an error will be returned.
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
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) setPositiveNumSignSpec(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	positiveNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupNanobot."+
			"setPositiveNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbols' is invalid!\n"+
			"'nStrNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = positiveNumberSign.IsValidInstanceError(
		ePrefix.XCpy("nStrNumSymbols.positiveNumberSign"))

	if err != nil {

		err2 := fmt.Errorf("numStrNumberSymbolGroupNanobot"+
			"setPositiveNumSignSpec()\n"+
			"Error: Input parameter 'positiveNumberSign' is invalid.\n\n"+
			"%v\n",
			err.Error())

		return err2
	}

	err = nStrNumSymbols.positiveNumberSign.CopyIn(
		&positiveNumberSign,
		ePrefix.XCpy(
			"nStrNumSymbols.positiveNumberSign<-"+
				"positiveNumberSign"))

	return err
}

//	setZeroNumSignRunes
//
//	Receives a series of rune arrays used to configure
//	the input parameter 'nStrNumSymbols' with new data
//	values for the Zero Number Sign Symbols.
//
//	Then Zero Number Sign Symbols internal member
//	variable is:
//
//		nStrNumSymbols.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Zero Number Sign Symbol member variable data fields
//	as specified by the rune arrays passed as input
//	parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols					*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		The Zero Number Sign Symbol Specifications
//		for this instance will be deleted and reset to
//		the values provided by the following input
//		parameters.
//
//	leadingZeroNumberSymbols	[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, type
//		NumStrNumberSymbolGroup gives the user the
//		option to configure both leading and trailing
//		Zero Number Sign Symbols.
//
//		Leading zero number sign symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Zero Number Sign Symbols
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//	trailingZeroNumberSymbols	[]rune
//
//		A rune array containing the trailing zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, type
//		NumStrNumberSymbolGroup gives the user the
//		option to configure both leading and trailing
//		Zero Number Sign Symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Trailing Zero Number Sign Symbols
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
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
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45+"
//					Number Field Index:------>01234567
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
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
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
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) setZeroNumSignRunes(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	leadingZeroNumberSymbols []rune,
	trailingZeroNumberSymbols []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupNanobot."+
			"setZeroNumSignRunes()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbols' is invalid!\n"+
			"'nStrNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenLeadingNumSyms := len(leadingZeroNumberSymbols)

	lenTrailingNumSyms := len(trailingZeroNumberSymbols)

	if lenLeadingNumSyms == 0 &&
		lenTrailingNumSyms == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameters 'leadingZeroNumberSymbols' and\n"+
			"'trailingZeroNumberSymbols' are invalid!\n"+
			"Both parameters are empty and contain zero text characters.\n",
			ePrefix.String())

		return err

	}

	if lenLeadingNumSyms > 0 &&
		lenTrailingNumSyms > 0 {

		err = nStrNumSymbols.zeroNumberSign.
			SetNumberSignLeadingTrailingSymbolRunes(
				leadingZeroNumberSymbols,
				zeroNumFieldSymPosition,
				trailingZeroNumberSymbols,
				zeroNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.zeroNumberSign"))

	} else if lenLeadingNumSyms > 0 &&
		lenTrailingNumSyms == 0 {

		err = nStrNumSymbols.zeroNumberSign.
			SetNumberSignLeadingSymbolRunes(
				leadingZeroNumberSymbols,
				zeroNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.zeroNumberSign"))

	} else {
		// MUST BE -
		//  lenLeadingNumSyms == 0  &&
		//		lenTrailingNumSyms > 0

		err = nStrNumSymbols.zeroNumberSign.
			SetNumberSignTrailingSymbolRunes(
				trailingZeroNumberSymbols,
				zeroNumFieldSymPosition,
				ePrefix.XCpy(
					"nStrNumSymbols.zeroNumberSign"))

	}

	return err
}

//	setZeroNumSignSpec
//
//	Receives a single NumStrNumberSymbolSpec object
//	configured as a Zero Number Sign Symbol.
//
//	This method then proceeds to reset the
//	corresponding Zero Number Sign Symbol member
//	variable data value for the NumStrNumberSymbolGroup
//	input paramter 'nStrNumSymbols'.
//
//	The Zero Number Sign Symbol internal member
//	variable is:
//
//		nStrNumSymbols.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Be advised that this method will delete and reset
//		the Zero Number Sign Symbol member variable
//		data fields contained in input paramter,
//		'nStrNumSymbols':
//
//			nStrNumSymbols.zeroNumberSign
//
//	(2)	If input parameter 'zeroNumberSign' is invalid
//		or a NOP, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		The Zero Number Sign for this instance will be
//		deleted and reset to the values	provided by input
//		parameter 'zeroNumberSign'.
//
//	zeroNumberSign				NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Symbol Specification in input paramter,
//		'nStrNumSymbols':
//
//			nStrNumSymbols.zeroNumberSign
//
//		If input parameter 'zeroNumberSign' is invalid
//		or a NOP, an error will be returned.
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
func (nStrNumSymbolGroupNanobot *numStrNumberSymbolGroupNanobot) setZeroNumSignSpec(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	zeroNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolGroupNanobot.lock == nil {
		nStrNumSymbolGroupNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupNanobot.lock.Lock()

	defer nStrNumSymbolGroupNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupNanobot."+
			"setZeroNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbols' is invalid!\n"+
			"'nStrNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = zeroNumberSign.IsValidInstanceError(
		ePrefix.XCpy("nStrNumSymbols.zeroNumberSign"))

	if err != nil {

		err2 := fmt.Errorf("numStrNumberSymbolGroupNanobot"+
			"setZeroNumSignSpec()\n"+
			"Error: Input parameter 'zeroNumberSign' is invalid.\n\n"+
			"%v\n",
			err.Error())

		return err2
	}

	err = nStrNumSymbols.zeroNumberSign.CopyIn(
		&zeroNumberSign,
		ePrefix.XCpy(
			"nStrNumSymbols.zeroNumberSign<-"+
				"zeroNumberSign"))

	return err
}

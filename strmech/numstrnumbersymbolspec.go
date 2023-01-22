package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrNumberSymbolSpec
//
// The Number String Number Symbol Specification,
// 'NumStrNumberSymbolSpec' contains specifications
// for number symbols used to construct and format
// numeric values in number strings for screen displays,
// file output and printing.
//
// Each valid instance of Type 'NumStrNumberSymbolSpec'
// will specify one of two types of number symbols:
//
//	(1)	Number Signs
//
//		Number signs such as the minus sign
//		('-'), plus sign ('+') or surrounding
//		parentheses ('()') used to designate negative
//		numeric values.
//
//		Number Sign Symbols are configured with methods
//		labeled "Number Sign"
//
//						OR
//
//	(2)	Currency Symbols
//
//		Currency Symbols such as the the Dollar sign
//		('$'), Euro sign ('€') or Pound sign ('£').
//
//		Currency Symbols are configured with methods
//		labeled "Currency"
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Radix points, or Decimal Separators, are used to
//		separate integer and fractional digits in floating
//		point numeric values.
//
//		In the United States, the decimal point or period
//		('.') serves as the decimal separator.
//
//			United States Example: 127.54
//
//		In various European countries, the comma (',') is
//		used as the decimal separator.
//
//			European Example: 127,54
//
//		Radix points or Decimal Separators are NOT
//		considered legitimate Number String Number Symbol
//		Specifications.
//
//		DO NOT USE Type NumStrNumberSymbolSpec to
//		configure Radix Points or Decimal Separators.
//
//		Radix Points, or Decimal Separators, should be
//		configured using Type DecimalSeparatorSpec.
//
//	(2)	To better understand the role of
//		NumStrNumberSymbolSpec in number string
//		formatting, reference Type
//		NumStrNumberSymbolGroup.
//
// ----------------------------------------------------------------
//
// # Background
//
// The Number String Number Symbol Specification
// (Type 'NumStrNumberSymbolSpec') defines either
// Currency Symbols or Number Sign Symbols.
//
// Examples of Currency Symbols include dollar signs
// ('$'), Euro signs ('€') and pound signs ('£').
//
// Examples of Number Sign Symbols include plus signs
// ('+'), minus signs ('-'), and surrounding parentheses
// ('()').
//
// Number String Number Symbols may be defined as leading
// character symbols, trailing character symbols, or leading
// and trailing symbols.
//
// Leading symbols are prepended at the beginning of a
// number string while trailing symbols are appended to the
// end of a number string.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example-1:
//		Leading Number Sign Symbols for Positive Values
//
//		Leading Number Sign Symbols: "+ "
//		Number String:   "+ 123.456"
//
//	Example-2: Leading Number Symbols
//		Leading Number Sign Symbols for Positive Values
//
//		Leading Number Sign Symbols: "+"
//		Number String:   "+123.456"
//
//	Example-3: Leading Number Symbols
//		Leading Number Sign Symbols for Negative Values
//
//		Leading Number Sign Symbols: "- "
//		Number String:   "- 123.456"
//
//	Example-4: Leading Number Symbols
//		Leading Number Sign Symbols for Negative Values
//
//		Leading Number Sign Symbols: "-"
//		Number String:   "-123.456"
//
//	Example-5: Leading Currency Symbols
//
//		Leading Currency Symbols: "$ "
//		Number String:   "$ 123.456"
//
//	Example-6: Leading Currency Symbols
//
//		Leading Currency Symbols: "$"
//		Number String:   "$123.456"
//
//	Example-7: Trailing Number Sign Symbols
//		Trailing Number Sign Symbols for Positive Values
//
//		Trailing Number Sign Symbols: " +"
//		Number String:   "123.456 +"
//
//	Example-8: Trailing Number Symbols
//		Trailing Number Sign Symbols for Positive Values
//
//		Trailing Number Sign Symbols: "+"
//		Number String:   "123.456+"
//
//	Example-9: Trailing Number Symbols
//		Trailing Number Sign Symbols for Negative Values
//
//		Trailing Number Sign Symbols: " -"
//		Number String:   "123.456 -"
//
//	Example-10: Trailing Number Symbols
//		Trailing Number Sign Symbols for Negative Values
//
//		Trailing Number Sign Symbols: "-"
//		Number String:   "123.456-"
//
//	Example-11: Trailing Currency Symbols
//		Trailing Currency Symbols for Positive Values
//
//		Trailing Currency Symbols: "€"
//		Number String:   "123.456€"
//
//	Example-12: Trailing Currency Symbols
//		Trailing Currency Symbols for Positive Values
//
//		Trailing Currency Symbols: " €"
//		Number String:   "123.456 €"
//
//	Example-13: Trailing Currency Symbols
//		Trailing Currency Symbols for Negative Values
//
//		Trailing Currency Symbols: " €"
//		Number String:   "123.456- €"
//
//	Example-14: Leading AND Trailing Number Sign Symbols
//		Leading AND Trailing Number Sign Symbols for
//		Negative Values
//
//		Leading Number Sign Symbols: "("
//		Trailing Number Sign Symbols: ")"
//		Number String:   "(123.456)"
type NumStrNumberSymbolSpec struct {
	leadingNumberSymbols RuneArrayDto
	// Contains the character or characters which
	// will be formatted and displayed in front of
	// a numeric value in a number string. Any
	// character or combination of characters
	// can be configured, to include currency
	// symbols.

	leadingNumberFieldSymbolPosition NumberFieldSymbolPosition
	// NumberFieldSymbolPosition is an enumeration
	// used to specify the position of a number symbol
	// relative to number field in which the number
	// string is formatted and displayed. Possible
	// values are listed as follows:
	//
	//	NumFieldSymPos.None() - Invalid Value
	// 	NumFieldSymPos.InsideNumField() - Valid Value
	//	NumFieldSymPos.OutsideNumField() - Valid Value
	//
	// 'leadingNumberFieldSymbolPosition' defines the
	// position of the Leading Number Symbol relative
	// to a Number Field in which a number string is
	// displayed. Valid values with examples are shown
	// below:
	//
	//	NumFieldSymPos.InsideNumField()
	//		Example-1:
	//			Number Field Length: 8
	//			Numeric Value: 123.45
	//			Number Symbol: leading minus sign ('-')
	//			Number Symbol Position: Inside Number Field
	//			Formatted Number String: " -123.45"
	//			Number Field Index:------>01234567
	//			Total Number String Length: 8
	//
	//		Example-2:
	//			Number Field Length: 10
	//			Numeric Value: 123.45
	//			Number Symbol: before and after parentheses  ('()')
	//			Number Symbol Position: Outside Number Field
	//          Number Text Justification: Centered
	//			Formatted Number String: " (123.45) "
	//			Number Field Index:------>0123456789
	//			Total Number String Length: 10
	//
	//		In this case the final length of the number string
	//		is defined by the Number Field length.
	//
	//	NumFieldSymPos.OutsideNumField()
	//		Example-3:
	//			Number Field Length: 8
	//	     	Numeric Value: 123.45
	//	     	Number Symbol: leading minus sign ('-')
	//	     	Number Symbol Position: Outside Number Field
	//	     	Formatted Number String: "-  123.45"
	//			Number Field Index:------>012345678
	//			Total Number String Length: 9
	//
	//		Example-4:
	//			Number Field Length: 8
	//			Numeric Value: 123.45
	//			Number Symbol: before and after parentheses  ('()')
	//			Number Symbol Position: Outside Number Field
	//			Formatted Number String: "( 123.45 )"
	//			Number Field Index:------>0123456789
	//			Total Number String Length: 10
	//
	//		In this case the final length of the number string
	//		is greater than the Number Field length.

	trailingNumberSymbols RuneArrayDto
	// Contains the character or characters which
	// will be formatted and displayed after a
	// numeric value in a number string. Any
	// character or combination of characters
	// can be configured to include currency
	// symbols.

	trailingNumberFieldSymbolPosition NumberFieldSymbolPosition
	// NumberFieldSymbolPosition is an enumeration
	// used to specify the position of a number symbol
	// relative to number field in which the number
	// string is formatted and displayed. Possible
	// values are listed as follows:
	//
	//	NumFieldSymPos.None() - Invalid Value
	// 	NumFieldSymPos.InsideNumField() - Valid Value
	//	NumFieldSymPos.OutsideNumField() - Valid Value
	//
	// 'trailingNumberFieldSymbolPosition' defines the
	// position of the Trailing Number Symbol relative
	// to a Number Field in which  a number string is
	// displayed. Valid values with examples are shown
	// below:
	//
	//	NumFieldSymPos.InsideNumField()
	//		Example-1:
	//			Number Field Length: 8
	//			Numeric Value: 123.45
	//			Number Symbol: trailing minus sign ('-')
	//			Number Symbol Position: Inside Number Field
	//          Number Text Justification: Right
	//			Formatted Number String: " 123.45-"
	//			Number Field Index:------>01234567
	//			Total Number String Length: 8
	//
	//		Example-2:
	//			Number Field Length: 10
	//			Numeric Value: 123.45
	//			Number Symbol: before and after parentheses  ('()')
	//			Number Symbol Position: Outside Number Field
	//          Number Text Justification: Centered
	//			Formatted Number String: " (123.45) "
	//			Number Field Index:------>0123456789
	//			Total Number String Length: 10
	//
	//		In this case the final length of the number string
	//		is defined by the Number Field length.
	//
	//	NumFieldSymPos.OutsideNumField()
	//		Example-3:
	//			Number Field Length: 8
	//	     	Numeric Value: 123.45
	//	     	Number Symbol: trailing minus sign ('-')
	//	     	Number Symbol Position: Outside Number Field
	//          Number Text Justification: Right
	//	     	Formatted Number String: "  123.45-"
	//			Number Field Index:------>012345678
	//			Total Number String Length: 9
	//
	//		Example-4:
	//			Number Field Length: 8
	//			Numeric Value: 123.45
	//			Number Symbol: before and after parentheses  ('()')
	//			Number Symbol Position: Outside Number Field
	//          Number Text Justification: Centered
	//			Formatted Number String: "( 123.45 )"
	//			Number Field Index:------>0123456789
	//			Total Number String Length: 10
	//
	//		In this case the final length of the number string
	//		is greater than the Number Field length.

	currencyNumSignRelativePos CurrencyNumSignRelativePosition
	// The Currency Number Sign Relative Position is used
	// exclusively by Currency Symbol Specifications and
	// only applies to currency symbols such as dollar
	// signs ('$'), Euro signs ('€') or Pound signs
	// ('£').
	//
	// This enumeration has three possible values, only
	// two of which are valid:
	//
	//	CurrNumSignRelPos.None()			- Invalid
	//	CurrNumSignRelPos.OutsideNumSign()	- Valid
	//	CurrNumSignRelPos.InsideNumSign()	- Valid
	//
	//	CurrNumSignRelPos is global constant used to
	//	abbreviate the syntax for invoking enumeration
	//	values. The formal syntax is:
	//
	//		CurrencyNumSignRelativePosition(0).OutsideNumSign()
	//		CurrencyNumSignRelativePosition(0).InsideNumSign()
	//
	// Currency Symbols have the option of being
	// positioned either inside or outside number sign
	// symbols formatted with numeric values in a number
	// string.
	//
	// Examples CurrNumSignRelPos.OutsideNumSign()
	//	Currency Symbol Outside of Number Sign
	//
	//		"$ -123.45"
	//		"123.45- €"
	//
	// Examples CurrNumSignRelPos.InsideNumSign()
	//	Currency Symbol Inside Of Number Sign
	//
	//	Examples:
	//		"- $123.45"
	//		"123.45€ -"
	//
	// Be Advised -
	//
	//		NumberFieldSymbolPosition Conflicts
	//
	//		When formatting a number string, the
	//		NumberFieldSymbolPosition values for both the
	//		Currency Symbol and the Number Sign Symbol
	//		MUST BE EQUAL before the Currency Number Sign
	//		Relative Position parameter,
	//		('currencyNumSignRelativePos'), will be
	//		activated and applied to the number string
	//		formatting algorithm.
	//
	//		If the NumberFieldSymbolPosition values for
	//		both the Currency Symbol and the Number Sign
	//		Symbol ARE NOT EQUAL, the
	//		NumberFieldSymbolPosition parameter controls
	//		and the Currency Number Sign Relative
	//		Position parameter,
	//		('currencyNumSignRelativePos'), will be
	//		ignored.
	//
	//		Example:
	//
	//			-- NumberFieldSymbolPosition Values NOT EQUAL --
	//
	//			Number Field Length: 8
	//		  	Numeric Value: -123.45
	//			Minus Sign Number Field Symbol Position:
	//				NumFieldSymPos.InsideNumField()
	//			Currency Number Field Symbol Position:
	//				NumFieldSymPos.OutsideNumField()
	//			Currency Number Sign Relative Position:
	//				CurrNumSignRelPos.InsideNumSign()
	//			Leading Currency Symbol: Dollar sign ('$')
	//			Number Text Justification: Right
	//			Formatted Number String: "$ -123.45"
	//			Number Field Index:------>012345678
	//			Total Number String Length: 9
	//
	//			Currency Symbol is Formatted OUTSIDE
	//			the Number Field.

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumStrNumberSymbolSpec ('incomingNStrNumSymSpec')
//	to the data fields of the current NumStrNumberSymbolSpec
//	instance ('nStrNumberSymbolSpec').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the data fields in current NumStrNumberSymbolSpec
//	instance ('nStrNumberSymbolSpec') will be deleted and
//	overwritten.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingNStrNumSymSpec	*NumStrNumberSymbolSpec
//
//		A pointer to an instance of NumStrNumberSymbolSpec.
//		This method will NOT change the values of internal member
//		variables contained in this instance.
//
//		All data values in this NumStrNumberSymbolSpec instance
//		will be copied to current NumStrNumberSymbolSpec
//		instance ('nStrNumberSymbolSpec').
//
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) CopyIn(
	incomingNStrNumSymSpec *NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecMolecule).copyNStrNumberSymbolSpec(
		nStrNumberSymbolSpec,
		incomingNStrNumSymSpec,
		ePrefix.XCpy(
			"nStrNumberSymbolSpec<-"+
				"incomingNStrNumSymSpec"))
}

//	CopyOut
//
//	Returns a deep copy of the current
//	NumStrNumberSymbolSpec instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	deepCopyNStrNumSymSpec		NumStrNumberSymbolSpec
//
//		If this method completes successfully and no errors are
//		encountered, this parameter will return a deep copy of the
//		current NumStrNumberSymbolSpec instance.
//
//
//	err							error
//
//		If the method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopyNStrNumSymSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNStrNumSymSpec, err
	}

	err = new(numStrNumberSymbolSpecMolecule).copyNStrNumberSymbolSpec(
		&deepCopyNStrNumSymSpec,
		nStrNumberSymbolSpec,
		ePrefix.XCpy(
			"deepCopyNStrNumSymSpec<-"+
				"nStrNumberSymbolSpec"))

	return deepCopyNStrNumSymSpec, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of NumStrNumberSymbolSpec to their zero or
//	uninitialized states. This method will leave the
//	current instance of NumStrNumberSymbolSpec in an
//	invalid state and unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all member variable data
//	values in the current instance of
//	NumStrNumberSymbolSpec. All member variable data
//	values will be reset to their zero or uninitialized
//	states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) Empty() {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	new(numStrNumberSymbolSpecMolecule).empty(
		nStrNumberSymbolSpec)

	nStrNumberSymbolSpec.lock.Unlock()

	nStrNumberSymbolSpec.lock = nil
}

// EmptyLeadingNStrNumSymbol - Resets the member variable
// data for the leading number symbol contained in the
// current instance of NumStrNumberSymbolSpec to an
// initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the leading number symbol
// contained in the current instance of NumStrNumberSymbolSpec
// will be deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) EmptyLeadingNStrNumSymbol() {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	new(nStrNumberSymbolSpecAtom).emptyLeadingNStrNumSymbol(
		nStrNumberSymbolSpec)
}

// EmptyTrailingNStrNumSymbol - Resets the member variable data
// for the trailing number symbol contained in the current
// instance of NumStrNumberSymbolSpec to an initial or
// zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the trailing number
// symbol contained in the current instance of
// NumStrNumberSymbolSpec will be deleted and reset to
// an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) EmptyTrailingNStrNumSymbol() {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	new(nStrNumberSymbolSpecAtom).emptyTrailingNStrNumSymbol(
		nStrNumberSymbolSpec)
}

// Equal - Receives a pointer to an NumStrNumberSymbolSpec
// object and proceeds to determine whether all data elements in
// this object are equal to all corresponding data elements in
// the current instance of NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNStrNumSymSpec *NumStrNumberSymbolSpec
//		This method will compare all data elements in the
//		current NumStrNumberSymbolSpec object to
//		corresponding data elements in this second
//		NumStrNumberSymbolSpec object in order
//		to determine equivalency.
//
//
//	errorPrefix					interface{}
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//		If all the data elements in the current
//		NumStrNumberSymbolSpec instance are equal to all the
//		corresponding data elements in 'incomingNStrNumSymSpec',
//		this return parameter will be set to 'true'. If all the data
//		elements are NOT equal, this return parameter will be set to
//		'false'.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) Equal(
	incomingNStrNumSymSpec *NumStrNumberSymbolSpec) bool {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return new(nStrNumberSymbolSpecAtom).equal(
		nStrNumberSymbolSpec,
		incomingNStrNumSymSpec)
}

// GetCurrencyNumSignRelativePosition
//
// Returns the current value of NumStrNumberSymbolSpec
// internal member variable,
// 'currencyNumSignRelativePos'.
//
// The Currency Number Sign Relative Position is used
// exclusively by Currency Symbol Specifications.
// This enumeration has three values, only two of
// which are valid:
//
//		CurrNumSignRelPos.None()			- Invalid
//	 	If this value is returned it signals that
//			the configured number symbol is NOT a
//			Currency Symbol.
//
//		CurrNumSignRelPos.OutsideNumSign()	- Valid
//		CurrNumSignRelPos.InsideNumSign()	- Valid
//
// Currency Symbols have the option of being
// positioned either inside or outside number sign
// symbols formatted with numeric values in a number
// string.
//
// Examples CurrNumSignRelPos.OutsideNumSign()
//
//	"$ -123.45"
//	"123.45- €"
//
// Examples CurrNumSignRelPos.InsideNumSign()
//
//	Examples:
//		"- $123.45"
//		"123.45€ -"
//
// The Currency Number Sign Relative Position therefore
// determines the location of the currency symbol
// relative to a number sign in a number string.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetCurrencyNumSignRelativePosition() CurrencyNumSignRelativePosition {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.currencyNumSignRelativePos
}

// GetLeadingNumberSymbolRunesDto
//
// Returns a RuneArrayDto containing the leading number
// symbol character or characters configured for the
// current instance of NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	RuneArrayDto
//
//		If this method completes successfully, this
//		parameter will be returned with an encapsulated
//		rune array containing the Leading Number Symbol
//		characters for the current instance of
//		NumStrNumberSymbolSpec.
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetLeadingNumberSymbolRunesDto(
	errorPrefix interface{}) (
	RuneArrayDto,
	error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var leadNumSymRunesDto RuneArrayDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"GetLeadingNumberSymbolRunesDto()",
		"")

	if err != nil {
		return leadNumSymRunesDto, err
	}

	leadNumSymRunesDto,
		err = nStrNumberSymbolSpec.leadingNumberSymbols.CopyOut(
		ePrefix.XCpy(
			"leadNumSymRunesDto<-"))

	return leadNumSymRunesDto, err
}

// GetLeadingNumberSymbolRunes
//
// Returns a rune array containing the leading number
// symbol character or characters configured for the
// current instance of NumStrNumberSymbolSpec.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetLeadingNumberSymbolRunes() []rune {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.leadingNumberSymbols.GetRuneArray()
}

// GetLeadingNumberSymbolStr
//
// Returns a string containing the leading number symbol
// character or characters configured for the current
// instance of NumStrNumberSymbolSpec.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetLeadingNumberSymbolStr() string {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.leadingNumberSymbols.GetCharacterString()
}

// GetLeadingNumberSymbolPosition
//
// Returns an instance of NumberFieldSymbolPosition.
//
// This enumeration value defines the position
// of the Leading Number Symbol relative to a
// Number Field in which a number string is
// displayed. Possible valid values are listed
// as follows:
//
//	NumFieldSymPos.InsideNumField()
//
//		Example-1:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
//		Example-2:
//			Number Field Length: 10
//			Numeric Value: 123.45
//			Number Symbol: before and after parentheses  ('()')
//			Number Symbol Position: Outside Number Field
//	     	Number Text Justification: Centered
//			Formatted Number String: " (123.45) "
//			Number Field Index:------>0123456789
//			Total Number String Length: 10
//
//		In this case the final length of the number string
//		is defined by the Number Field length.
//
//	NumFieldSymPos.OutsideNumField()
//
//		Example-3:
//			Number Field Length: 8
//	     	Numeric Value: 123.45
//	     	Number Symbol: leading minus sign ('-')
//	     	Number Symbol Position: Outside Number Field
//	     	Formatted Number String: "-  123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//		Example-4:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: before and after parentheses  ('()')
//			Number Symbol Position: Outside Number Field
//			Formatted Number String: "( 123.45 )"
//			Number Field Index:------>0123456789
//			Total Number String Length: 10
//
//		In this case the final length of the number string
//		is greater than the Number Field length.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NumberFieldSymbolPosition
//		If this method completes successfully, an
//		instance of NumberFieldSymbolPosition for
//		the Leading Number Symbol will be returned.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetLeadingNumberSymbolPosition() NumberFieldSymbolPosition {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.leadingNumberFieldSymbolPosition
}

// GetNumStrSymbolPosition
//
// Returns an instance of NumSignSymbolPosition showing
// the position or location of the Number Symbol in
// relation to the numeric value configured in a number
// string.
//
// The returned type, 'NumSignSymbolPosition' will be
// set to one of four values:
//
//	NumSignSymPos.None()
//		Signals that the current instance of
//		NumStrNumberSymbolSpec is a NOP (No Operation)
//		instance configured with empty or uninitialized
//		values.
//
//	NumSignSymPos.Before()
//		Signals that the number symbol configured for
//		this current instance of NumStrNumberSymbolSpec
//		is a leading number symbol positioned at the
//		beginning of a number string.
//
//	NumSignSymPos.After()
//		Signals that the number symbol configured for
//		this current instance of NumStrNumberSymbolSpec
//		is a trailing number symbol positioned at the end
//		of a number string.
//
//	NumSignSymPos.BeforeAndAfter()
//		Signals that this current instance of
//		NumStrNumberSymbolSpec contains two number
//		symbols. One configured as a leading number
//		symbol and the other configured as a trailing
//		number symbol. One example of this combination
//		would be surrounding parentheses ("()") for
//		negative numeric values.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetNumStrSymbolPosition() NumSignSymbolPosition {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	lenLeadingSym :=
		nStrNumberSymbolSpec.leadingNumberSymbols.GetRuneArrayLength()

	lenTrailingSym :=
		nStrNumberSymbolSpec.trailingNumberSymbols.GetRuneArrayLength()

	if lenLeadingSym > 0 &&
		lenTrailingSym > 0 {

		return NumSignSymPos.BeforeAndAfter()
	}

	if lenLeadingSym > 0 &&
		lenTrailingSym == 0 {

		return NumSignSymPos.Before()
	}

	if lenLeadingSym == 0 &&
		lenTrailingSym > 0 {

		return NumSignSymPos.After()
	}

	// MUST BE
	//  lenLeadingSym == 0 &&
	//	lenTrailingSym == 0

	return NumSignSymPos.None()
}

// GetTrailingNumberSymbolRunesDto
//
// Returns a RuneArrayDto containing the trailing number
// symbol character or characters configured for the
// current instance of NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	RuneArrayDto
//
//		If this method completes successfully, this
//		parameter will be returned with an encapsulated
//		rune array containing the Trailing Number Symbol
//		characters for the current instance of
//		NumStrNumberSymbolSpec.
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetTrailingNumberSymbolRunesDto(
	errorPrefix interface{}) (
	RuneArrayDto,
	error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var leadNumSymRunesDto RuneArrayDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyDefaultsEU()",
		"")

	if err != nil {
		return leadNumSymRunesDto, err
	}

	leadNumSymRunesDto,
		err = nStrNumberSymbolSpec.trailingNumberSymbols.CopyOut(
		ePrefix.XCpy(
			"leadNumSymRunesDto<-"))

	return leadNumSymRunesDto, err
}

// GetTrailingNumberSymbolRunes
// Returns a rune array containing the trailing number
// symbol character or characters configured for the
// current instance of NumStrNumberSymbolSpec.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetTrailingNumberSymbolRunes() []rune {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.trailingNumberSymbols.GetRuneArray()
}

// GetTrailingNumberSymbolStr
//
// Returns a string containing the trailing number symbol
// character or characters configured for the current
// instance of NumStrNumberSymbolSpec.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetTrailingNumberSymbolStr() string {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.trailingNumberSymbols.GetCharacterString()
}

// GetTrailingNumberSymbolPosition
//
// Returns an instance of NumberFieldSymbolPosition.
//
// This enumeration value defines the position
// of the Trailing Number Symbol relative to a
// Number Field in which a number string is
// displayed. Possible valid values are listed
// as follows:
//
//	NumFieldSymPos.InsideNumField()
//
//		Example-1:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: trailing minus sign ('-')
//			Number Symbol Position: Inside Number Field
//	     Number Text Justification: Right
//			Formatted Number String: " 123.45-"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
//		Example-2:
//			Number Field Length: 10
//			Numeric Value: 123.45
//			Number Symbol: before and after parentheses  ('()')
//			Number Symbol Position: Outside Number Field
//	     Number Text Justification: Centered
//			Formatted Number String: " (123.45) "
//			Number Field Index:------>0123456789
//			Total Number String Length: 10
//
//		In this case the final length of the number string
//		is defined by the Number Field length.
//
//	NumFieldSymPos.OutsideNumField()
//
//		Example-3:
//			Number Field Length: 8
//	     	Numeric Value: 123.45
//	     	Number Symbol: trailing minus sign ('-')
//	     	Number Symbol Position: Outside Number Field
//	     Number Text Justification: Right
//	     	Formatted Number String: "  123.45-"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//		Example-4:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: before and after parentheses  ('()')
//			Number Symbol Position: Outside Number Field
//	     Number Text Justification: Centered
//			Formatted Number String: "( 123.45 )"
//			Number Field Index:------>0123456789
//			Total Number String Length: 10
//
//		In this case the final length of the number string
//		is greater than the Number Field length.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NumberFieldSymbolPosition
//		If this method completes successfully, an
//		instance of NumberFieldSymbolPosition for
//		the Trailing Number Symbol will be returned.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetTrailingNumberSymbolPosition() NumberFieldSymbolPosition {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.trailingNumberFieldSymbolPosition
}

//	IsNOP
//
//	Stands for 'Is No Operation'. This method returns a
//	boolean value signaling whether this instance of
//	NumStrNumberSymbolSpec is engaged, valid and
//	operational with respect to the application of a
//	number symbol.
//
//	NOP is a computer science term which stands for
//	'No Operation' meaning it performs no operations
//	and serves an empty placeholder.
//
//	If 'IsNOP' is set to 'true', it signals that this
//	Number String Number Symbol Specification is simply
//	an empty placeholder and performs no active role in,
//	and is completely ignored by, Number String Formatting
//	algorithms. When 'IsNOP' set to 'true', no Number
//	Symbol will be inserted or formatted as part of a
//	number sign in a Number String text presentation.
//
//	If this method returns 'false', it signals that the
//	current instance of NumStrNumberSymbolSpec is fully
//	populated, valid and functional. Number Strings
//	Formatting operations will therefore include the
//	specified number symbol in formatted number strings.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If this method returns 'true', it signals that the
//		current instance of NumStrNumberSymbolSpec is a
//		'NOP' or No Operation. This condition signals that
//		the Number String Number Symbol Specification is
//		simply an empty placeholder and performs no active
//		role in, and is completely ignored by, Number
//		String Formatting algorithms. In this case, no
//		number symbol will be inserted in formatted number
//		strings.
//
//		If this method returns 'false', it signals that
//		the current instance of NumStrNumberSymbolSpec
//		is fully populated, valid and functional. Number
//		Strings Formatting operations will therefore
//		include the specified number symbol in formatted
//		number strings.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) IsNOP() bool {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	isValid,
		_ := new(nStrNumberSymbolSpecAtom).
		testValidityNumStrNumberSymbolSpec(
			nStrNumberSymbolSpec,
			nil)

	return !isValid
}

//	IsValidInstance
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current NumStrNumberSymbolSpec
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this
//	method will return a boolean value of 'false'.
//
//	If all data elements are determined to be valid,
//	this method returns a boolean value of 'true'.
//
//	This method is functionally equivalent to
//	NumStrNumberSymbolSpec.IsValidInstanceError() with
//	the sole exceptions being that this method takes
//	no input parameters and returns a boolean value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If any of the internal data values contained in
//		the current instance of NumStrNumberSymbolSpec
//		are found to be invalid, this method will return
//		a boolean value	of 'false'.
//
//		If all internal member data values contained in
//		the current instance of NumStrNumberSymbolSpec
//		are found to be valid, this method returns a
//		boolean value of 'true'.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) IsValidInstance() bool {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	isValid,
		_ := new(nStrNumberSymbolSpecAtom).
		testValidityNumStrNumberSymbolSpec(
			nStrNumberSymbolSpec,
			nil)

	return isValid
}

// IsValidInstanceError
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current NumStrNumberSymbolSpec
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this
//	method will return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
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
//		If any of the internal member data variables
//		contained in the current instance of
//		NumStrNumberSymbolSpec are found to be invalid,
//		this method will return an error containing an
//		appropriate message identifying the invalid
//		member data variable.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) IsValidInstanceError(
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(nStrNumberSymbolSpecAtom).
		testValidityNumStrNumberSymbolSpec(
			nStrNumberSymbolSpec,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec"))

	return err
}

// NewCurrencyDefaultsEU
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with the default
// currency symbol used by the European Union (EU).
//
// The default EU currency symbol is a trailing Euro
// sign ('€').
//
//	Example:
//		123.45 €
//
// The Number String Number Symbol Specification type
// (NumStrNumberSymbolSpec) is designed to assist in
// formatting numeric values as number strings for
// screen displays, printing or file output.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	euCurrencySymbols			NumStrNumberSymbolSpec
//
//		If this method completes successfully, an
//		instance of NumStrNumberSymbolSpec will be
//		returned configured with the default EU
//		(European Union) Currency Symbol.
//
//
//		The default EU currency symbol is a trailing Euro
//		sign ('€').
//
//			Example:
//				123.45 €
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyDefaultsEU(
	errorPrefix interface{}) (
	euCurrencySymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyDefaultsEU()",
		"")

	if err != nil {
		return euCurrencySymbols, err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setCurrencyDefaultsEU(
			&euCurrencySymbols,
			ePrefix.XCpy(
				"euCurrencySymbols"))

	return euCurrencySymbols, err
}

// NewCurrencyDefaultsUKMinusInside
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with the default UK
// (United Kingdom) currency symbol.
//
// The Number String Number Symbol Specification type
// (NumStrNumberSymbolSpec) is designed to assist in
// formatting numeric values as number strings for
// screen displays, printing or file output.
//
// The default UK currency symbol is a leading pound
// sign ('£').
//
//	Example:
//		£ 123.45  Positive Value
//
// This method will configure the pound sign ('£')
// such that any minus sign configured for negative
// numeric values will be inside the pound sign ('£').
//
//	Example:
//		£ -123.45  Negative Value
//
// To generate a default configuration where the minus
// sign is outside the pound sign ('£'), use method:
//
//	NumStrNumberSymbolSpec.NewCurrencyDefaultsUKMinusOutside()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	ukCurrencySymbols			NumStrNumberSymbolSpec
//
//		If this method completes successfully, an
//		instance of NumStrNumberSymbolSpec will be
//		returned configured with the default UK
//		(United Kingdom) Currency Symbol.
//
//		The default UK currency symbol is a leading
//		pound sign ('£').
//
//		Minus signs configured in association with
//		this currency symbol will be positioned
//		inside the pound sign ('£').
//
//			Examples:
//				£ 123.45 	Positive Value
//				£ -123.45	Negative Value
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyDefaultsUKMinusInside(
	errorPrefix interface{}) (
	ukCurrencySymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyDefaultsUKMinusInside()",
		"")

	if err != nil {
		return ukCurrencySymbols, err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setCurrencyDefaultsUKMinusInside(
			&ukCurrencySymbols,
			ePrefix.XCpy(
				"ukCurrencySymbols"))

	return ukCurrencySymbols, err
}

// NewCurrencyDefaultsUKMinusOutside
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with the default UK
// (United Kingdom) currency symbol.
//
// The Number String Number Symbol Specification type
// (NumStrNumberSymbolSpec) is designed to assist in
// formatting numeric values as number strings for
// screen displays, printing or file output.
//
// The default Uk currency symbol is a leading pound
// sign ('£').
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
// To generate a default configuration where the minus
// sign is inside the pound sign ('£'), use method:
//
//	NumStrNumberSymbolSpec.NewCurrencyDefaultsUKMinusInside()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	ukCurrencySymbols			NumStrNumberSymbolSpec
//
//		If this method completes successfully, an
//		instance of NumStrNumberSymbolSpec will be
//		returned configured with the default UK
//		(United Kingdom) Currency Symbol.
//
//		Minus signs configured in association with
//		this currency symbol will be positioned
//		outside the pound sign ('£').
//
//		The default US currency symbol is a leading
//		dollar sign.
//
//			Examples:
//				£ 123.45 	- Positive Value
//				-£ 123.45	- Negative Value
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyDefaultsUKMinusOutside(
	errorPrefix interface{}) (
	ukCurrencySymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyDefaultsUKMinusOutside()",
		"")

	if err != nil {
		return ukCurrencySymbols, err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setCurrencyDefaultsUKMinusOutside(
			&ukCurrencySymbols,
			ePrefix.XCpy(
				"ukCurrencySymbols"))

	return ukCurrencySymbols, err
}

// NewCurrencyDefaultsUS
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with the default US
// (United States) currency symbol.
//
// The default US currency symbol is a leading dollar
// sign.
//
//	Example:
//		$ 123.45
//
// The Number String Number Symbol Specification type
// (NumStrNumberSymbolSpec) is designed to assist in
// formatting numeric values as number strings for
// screen displays, printing or file output.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	usCurrencySymbols			NumStrNumberSymbolSpec
//
//		If this method completes successfully, an
//		instance of NumStrNumberSymbolSpec will be
//		returned configured with the default US
//		Currency Symbol.
//
//		The default US currency symbol is a leading
//		dollar sign.
//
//			Example:
//				$ 123.45
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyDefaultsUS(
	errorPrefix interface{}) (
	usCurrencySymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyDefaultsUS()",
		"")

	if err != nil {
		return usCurrencySymbols, err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setCurrencyDefaultsUS(
			&usCurrencySymbols,
			ePrefix.XCpy(
				"usCurrencySymbols"))

	return usCurrencySymbols, err
}

// NewCurrencyLeadingSymbol
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with a Leading
// Currency Symbol character or characters.
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
// Currency Symbols for the current instance of
// NumStrNumberSymbolSpec.
//
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading Currency Symbol.
//
//		Example:
//			Leading Currency Symbol: "$ "
//			Formatted Number String: "$ 123.45"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingCurrencySymbol     		string
//
//		A string containing one or more Leading
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of NumStrNumberSymbolSpec, configured
//		with Leading Currency Symbols, will be returned.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Currency Symbols
//
//	Leading Currency Symbols: "$ "
//	Number String:   "$ 123.456"
//
// Example-2: Leading Currency Symbols
//
//	Leading Currency Symbols: "$"
//	Number String:   "$123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyLeadingSymbol(
	leadingCurrencySymbol string,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyLeadingSymbol()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			&newNStrNumberSymbolSpec,
			[]rune(leadingCurrencySymbol),
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingCurrencySymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewCurrencyLeadingSymbolRunes
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with a Leading
// Currency Symbol character or characters.
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
// Currency Symbols for the current instance of
// NumStrNumberSymbolSpec.
//
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading Currency Symbol.
//
//		Example:
//			Leading Currency Symbol: "$ "
//			Formatted Number String: "$ 123.45"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of NumStrNumberSymbolSpec, configured
//		with Leading Currency Symbols, will be returned.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Currency Symbols
//
//	Leading Currency Symbols: []rune{'$',' '}
//	Number String:   "$ 123.456"
//
// Example-2: Leading Currency Symbols
//
//	Leading Currency Symbols: []rune{'$'}
//	Number String:   "$123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyLeadingSymbolRunes(
	leadingCurrencySymbol []rune,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyLeadingSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			&newNStrNumberSymbolSpec,
			leadingCurrencySymbol,
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingCurrencySymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewCurrencyLeadingTrailingSymbols
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with both a Leading
// Currency Symbol and a Trailing Currency Symbol.
//
// The Number String Number Symbol Specification type
// (NumStrNumberSymbolSpec) is designed to assist in
// formatting numeric values as number strings for
// screen displays, printing or file output.
//
// Examples of Currency Symbol characters include such
// symbols as the Dollar sign ('$'), Euro sign ('€') or
// Pound sign ('£').
//
// Leading currency symbols are prefixed or prepended
// at the beginning of a number string while trailing
// currency symbols are suffixed or appended at the end
// of a number string.
//
//	Example Leading Currency Sign : $123.45
//
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
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading and/or Trailing Currency Symbols.
//
//		Example:
//			Leading Currency Symbol: "$ "
//			Trailing Currency Symbol: " $"
//			Formatted Number String: "$ 123.45 $"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingCurrencySymbol     		string
//
//		A string containing one or more Leading
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//		If this parameter is submitted as an empty string,
//		an error will be returned.
//
//	trailingCurrencySymbol     		string
//
//		A string containing one or more Trailing
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
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
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right
//				Formatted Number String: " $123.45$"
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:------>012345678901
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
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:------>012345678901
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//	 	instance of NumStrNumberSymbolSpec, configured
//	 	with Leading and Trailing Currency Symbols, will
//	 	be returned.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading and Trailing Currency Symbols
//
//	Leading Currency Symbols: "$ "
//	Trailing Currency Symbols: " $"
//	Number String:   "$ 123.456 $"
//
// Example-2: Leading and Trailing Currency Symbols
//
//	Leading Currency Symbols: "$"
//	Trailing Currency Symbols: "$"
//	Number String:   "$123.456$"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyLeadingTrailingSymbols(
	leadingCurrencySymbol string,
	trailingCurrencySymbol string,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyLeadingTrailingSymbols()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setLeadingTrailingCurrencySymbols(
			&newNStrNumberSymbolSpec,
			[]rune(leadingCurrencySymbol),
			[]rune(trailingCurrencySymbol),
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"))

	return newNStrNumberSymbolSpec, err
}

// NewCurrencyLeadingTrailingSymbolRunes
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with both a Leading
// Currency Symbol and a Trailing Currency Symbol.
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
// Leading currency symbols are prefixed or prepended
// at the beginning of a number string while trailing
// currency symbols are suffixed or appended at the end
// of a number string.
//
//	Example Leading Currency Sign : $123.45
//
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
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading and/or Trailing Currency Symbols.
//
//		Example:
//			Leading Currency Symbol: "$ "
//			Trailing Currency Symbol: " $"
//			Formatted Number String: "$ 123.45 $"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//		If this parameter is submitted as an empty string,
//		an error will be returned.
//
//	trailingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
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
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right
//				Formatted Number String: " $123.45$"
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:------>012345678901
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
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:------>012345678901
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//	 	instance of NumStrNumberSymbolSpec, configured
//	 	with Leading and Trailing Currency Symbols, will
//	 	be returned.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading and Trailing Currency Symbols
//
//	Leading Currency Symbols: []rune{'$', ' '}
//	Trailing Currency Symbols: []rune{' ', '$'}
//	Number String:   "$ 123.456 $"
//
// Example-2: Leading and Trailing Currency Symbols
//
//	Leading Currency Symbols: []rune{'$'}
//	Trailing Currency Symbols: []rune{'$'}
//	Number String:   "$123.456$"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyLeadingTrailingSymbolRunes(
	leadingCurrencySymbol string,
	trailingCurrencySymbol string,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyLeadingTrailingSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setLeadingTrailingCurrencySymbols(
			&newNStrNumberSymbolSpec,
			[]rune(leadingCurrencySymbol),
			[]rune(trailingCurrencySymbol),
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"))

	return newNStrNumberSymbolSpec, err
}

// NewCurrencyTrailingSymbol
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with a Trailing
// Currency Symbol character or characters.
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
// Trailing Currency Symbols are suffixed or appended
// to the end of a number string.
//
//	Example: 125.34€
//
// Currency symbol characters are intended for use in
// formatting currency numeric values displayed in number
// strings.
//
// This method will configure and store Trailing Currency
// Symbols for the current instance of
// NumStrNumberSymbolSpec.
//
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Trailing Currency Symbol.
//
//		Example:
//			Trailing Currency Symbol: " €"
//			Formatted Number String: "123.45 €"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingCurrencySymbol     		string
//
//		A string containing one or more Trailing
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
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
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 11
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  123.45€  "
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
//				Trailing Currency Symbol: Dollar sign ('€')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "  123.45€"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "  123.45  €"
//				Number Field Index:------>01234567890
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
//		  	Numeric Value: -123.45
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of NumStrNumberSymbolSpec, configured
//		with Trailing Currency Symbols, will be returned.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Currency Symbols
//
//	Trailing Currency Symbols: " €"
//	Number String:   "123.456 €"
//
// Example-2: Trailing Currency Symbols
//
//	Trailing Currency Symbols: "€"
//	Number String:   "123.456€"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyTrailingSymbol(
	trailingCurrencySymbol string,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyTrailingSymbol()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setTrailingCurrencySymbol(
			&newNStrNumberSymbolSpec,
			[]rune(trailingCurrencySymbol),
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingCurrencySymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewCurrencyTrailingSymbolRunes
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with a Trailing
// Currency Symbol character or characters.
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
// Trailing Currency Symbols are suffixed or appended
// to the end of a number string.
//
//	Example: 125.34€
//
// Currency symbol characters are intended for use in
// formatting currency numeric values displayed in number
// strings.
//
// This method will configure and store Trailing
// Currency Symbols for the current instance of
// NumStrNumberSymbolSpec.
//
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Trailing Currency Symbol.
//
//		Example:
//			Trailing Currency Symbol: " €"
//			Formatted Number String: "123.45 €"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
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
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 11
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  123.45€  "
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
//				Trailing Currency Symbol: Dollar sign ('€')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "  123.45€"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "  123.45  €"
//				Number Field Index:------>01234567890
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
//		  	Numeric Value: -123.45
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of NumStrNumberSymbolSpec, configured
//		with Trailing Currency Symbols, will be returned.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Currency Symbols
//
//	Trailing Currency Symbols: []rune{' ', '€'}
//	Number String:   "123.456 €"
//
// Example-2: Trailing Currency Symbols
//
//	Trailing Currency Symbols: []rune{'€'}
//	Number String:   "123.456€"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyTrailingSymbolRunes(
	trailingCurrencySymbol []rune,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewCurrencyTrailingSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setTrailingCurrencySymbol(
			&newNStrNumberSymbolSpec,
			trailingCurrencySymbol,
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingCurrencySymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewNOP - Returns a new, empty instance of
// NumStrNumberSymbolSpec. This instance is designed to
// serve as an empty placeholder or NOP.
//
// NOP is a computer science term which stands for
// 'No Operation' meaning it performs no operations and
// serves an empty placeholder.
//
// This method is often used to configure implied
// positive number signs in number string formatting. The
// positive number sign, plus ('+'), is often implied for
// positive numeric values and is not displayed in the
// number string.
//
// Call this method when creating implied positive number
// signs for number strings.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewNOP() NumStrNumberSymbolSpec {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	newNOPNumSymSpec := NumStrNumberSymbolSpec{}

	new(numStrNumberSymbolSpecMolecule).empty(
		&newNOPNumSymSpec)

	return newNOPNumSymSpec
}

// NewNumberSignLeadingSymbol
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with a leading
// number sign symbol character or characters.
//
// Leading number sign symbol characters can include such
// symbols as plus signs ('+') or minus signs ('-').
//
// Leading number symbol characters are intended for use
// in formatting numeric values displayed in number
// strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol     		string
//
//		A string containing the leading number sign
//		symbol character or characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading number sign symbol characters can include
//		such symbols as plus signs ('+') or minus signs
//		('-').
//
//		Example-1:
//			Numeric Value: -123.45
//			Number Sign Symbol: leading minus sign ("-")
//			Formatted Number String: "-123.45"
//
//		Example-2:
//			Numeric Value: 123.45
//			Number Sign Symbol: leading plus sign ("+")
//			Formatted Number String: "+123.45"
//
//		Example-3:
//			Numeric Value: -123.45
//			Number Sign Symbol: leading minus sign ("- ")
//			Formatted Number String: "- 123.45"
//
//		Example-4:
//			Numeric Value: 123.45
//			Number Sign Symbol: leading plus sign ("+ ")
//			Formatted Number String: "+ 123.45"
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		'NumFieldSymPos' is global constant used to
//		abbreviate the syntax for invoking these
//		enumeration	values. The formal syntax is:
//
//			NumberFieldSymbolPosition(0).InsideNumField()
//			NumberFieldSymbolPosition(0).OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//		Example-1:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Sign Symbol: leading minus sign ("-")
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
//		Example-2:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Sign Symbol: leading plus sign ("+")
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right
//			Formatted Number String: " +123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
//			For the 'NumFieldSymPos.InsideNumField()'
//			specification, the final length of the number
//			string is defined by the Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//		Example-3:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Outside Number Field
//			Number Text Justification: Right
//			Formatted Number String: "-  123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//		Example-4:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading plus sign ('+')
//			Number Symbol Position: Outside Number Field
//			Number Text Justification: Right
//			Formatted Number String: "+  123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			For the 'NumFieldSymPos.OutsideNumField()'
//			specification, the final length of the number
//			string is greater than the Number Field
//			length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of	NumStrNumberSymbolSpec, configured
//		with leading sign number symbol characters, will
//		be returned through	this parameter.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Sign Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Sign Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewNumberSignLeadingSymbol(
	leadingNumberSymbol string,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewNumberSignLeadingSymbol()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		setLeadingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			[]rune(leadingNumberSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewNumberSignLeadingSymbolRunes
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with a leading
// number sign symbol character or characters.
//
// Leading number sign symbol characters can include such
// symbols as plus signs ('+') or minus signs ('-').
//
// Leading number symbol characters are intended for use
// in formatting numeric values displayed in number
// strings.
//
// This method is similar to NewNumberSignLeadingSymbol()
// with sole exception being that this method receives
// rune array input parameters instead of strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSignSymbol			[]rune
//
//		A rune array containing the leading number sign
//		symbol character or characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading number sign symbol characters can include
//		such symbols as plus signs ('+') or minus signs
//		('-').
//
//		Example-1:
//			Numeric Value: -123.45
//			Number Sign Symbol: leading minus sign ("-")
//			Formatted Number String: "-123.45"
//
//		Example-2:
//			Numeric Value: 123.45
//			Number Sign Symbol: leading plus sign ("+")
//			Formatted Number String: "+123.45"
//
//		Example-3:
//			Numeric Value: -123.45
//			Number Sign Symbol: leading minus sign ("- ")
//			Formatted Number String: "- 123.45"
//
//		Example-4:
//			Numeric Value: 123.45
//			Number Sign Symbol: leading plus sign ("+ ")
//			Formatted Number String: "+ 123.45"
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//			NumFieldSymPos.OutsideNumField()
//
//		'NumFieldSymPos' is global constant used to
//		abbreviate the syntax for invoking these
//		enumeration	values. The formal syntax is:
//
//			NumberFieldSymbolPosition(0).InsideNumField()
//			NumberFieldSymbolPosition(0).OutsideNumField()
//
//		Examples NumFieldSymPos.InsideNumField()
//
//		Example-1:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Sign Symbol: leading minus sign ("-")
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
//		Example-2:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Sign Symbol: leading plus sign ("+")
//			Number Symbol Position: Inside Number Field
//		   	Number Text Justification: Right
//			Formatted Number String: " +123.45"
//		    Number Field Index:------>01234567
//			Total Number String Length: 8
//
//			For the 'NumFieldSymPos.InsideNumField()'
//			specification, the final length of the number
//			string is defined by the Number Field length.
//
//		Examples NumFieldSymPos.OutsideNumField()
//
//		Example-3:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Outside Number Field
//			Number Text Justification: Right
//			Formatted Number String: "-  123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//		Example-4:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading plus sign ('+')
//			Number Symbol Position: Outside Number Field
//			Number Text Justification: Right
//			Formatted Number String: "+  123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			For the 'NumFieldSymPos.OutsideNumField()'
//			specification, the final length of the number
//			string is greater than the Number Field
//			length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of NumStrNumberSymbolSpec, configured
//		with leading number sign symbol characters, will
//		be returned through	this parameter.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Sign Symbols
//
//	Leading Symbols: []rune{'+',' '}
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Sign Symbols
//
//	Leading Symbols: []rune{'-',' '}
//	Number String:   "- 123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewNumberSignLeadingSymbolRunes(
	leadingNumberSignSymbol []rune,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewNumberSignLeadingSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		setLeadingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			leadingNumberSignSymbol,
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingNumberSignSymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewNumberSignLeadingTrailingSymbol
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with both leading
// and trailing number sign symbols.
//
// Leading and Trailing number symbol characters can
// include such symbols as plus signs ('+') or minus
// signs ('-').
//
// Leading and Trailing number symbols are intended for
// use in formatting numeric values displayed in number
// strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol     		string
//
//		A string containing the leading number sign
//		symbol character or characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading number symbol characters can include such
//		symbols as plus signs ('+') or minus signs ('-').
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
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
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	trailingNumberSymbol     		string
//
//		A string containing the trailing number sign
//		symbol character or characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Trailing number sign symbol characters can
//		include such symbols as plus signs ('+') or minus
//		signs ('-').
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of NumStrNumberSymbolSpec, configured
//		with both leading and trailing number sign symbol
//		characters, will be returned through this
//		parameter.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Sign Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Sign Symbols
//
//	Leading Symbols: "+"
//	Number String:   "+123.456"
//
// Example-3: Leading Number Sign Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-4: Trailing Number Sign Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-5: Trailing Number Sign Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-6: Trailing Number Sign Symbols
//
//	Trailing Symbols: "-"
//	Number String:   "123.456-"
//
// Example-7: Trailing Number Sign Symbols
//
//	 Leading Symbols: "("
//		Trailing Symbols: ")"
//		Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewNumberSignLeadingTrailingSymbol(
	leadingNumberSymbol string,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	trailingNumberSymbol string,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewNumberSignLeadingTrailingSymbol()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	err = nStrNumSymSpecNanobot.
		setLeadingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			[]rune(leadingNumberSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = nStrNumSymSpecNanobot.
		setTrailingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			[]rune(trailingNumberSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewNumberSignLeadingTrailingSymbolRunes
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with both leading
// and trailing number sign symbols.
//
// Leading and Trailing number sign symbol characters
// can include such symbols as plus signs ('+') or minus
// signs ('-').
//
// Leading and Trailing number symbols are intended for
// use in formatting numeric values displayed in number
// strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol				[]rune
//
//		An array of runes containing the leading number
//		sign symbol character or characters used to
//		configure the returned instance of
//		NumStrNumberSymbolSpec.
//
//		Leading number symbol characters can include such
//		symbols as plus signs ('+') or minus signs ('-').
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
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
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	trailingNumberSymbol			[]rune
//
//		An array of runes containing the trailing number
//		sign symbol character or characters used to
//		configure the returned instance of
//		NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+') or minus signs ('-').
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of NumStrNumberSymbolSpec, configured
//		with both leading and trailing number sign symbol
//		characters, will be returned through this
//		parameter.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Sign Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Sign Symbols
//
//	Leading Symbols: "+"
//	Number String:   "+123.456"
//
// Example-3: Leading Number Sign Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-4: Trailing Number Sign Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-5: Trailing Number Sign Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-6: Trailing Number Sign Symbols
//
//	Trailing Symbols: "-"
//	Number String:   "123.456-"
//
// Example-7: Trailing Number Sign Symbols
//
//	 Leading Symbols: "("
//		Trailing Symbols: ")"
//		Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewNumberSignLeadingTrailingSymbolRunes(
	leadingNumberSymbol string,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	trailingNumberSymbol string,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewNumberSignLeadingTrailingSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	err = nStrNumSymSpecNanobot.
		setLeadingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			[]rune(leadingNumberSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = nStrNumSymSpecNanobot.
		setTrailingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			[]rune(trailingNumberSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewNumberSignTrailingSymbol
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with a trailing
// number sign symbol character or characters.
//
// Trailing number symbol characters can include such
// symbols as plus signs ('+') or minus signs ('-').
//
// Trailing number symbol characters are intended for
// use in formatting numeric values displayed in number
// strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingNumberSymbol     		string
//
//		A string containing the trailing number sign
//		symbol character or characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include
//		such symbols as plus signs ('+') or minus signs
//		('-').
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance ofNumStrNumberSymbolSpec, configured
//		with trailing number sign symbol characters, will
//		be returned through this parameter.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Number Sign Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-2: Trailing Number Sign Symbols
//
//	Trailing Symbols: "+"
//	Number String:   "123.456+"
//
// Example-3: Trailing Number Sign Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-4: Trailing Number Sign Symbols
//
//	Trailing Symbols: "-"
//	Number String:   "123.456-"
//
// Example-5: Trailing Number Sign Symbols
//
//	 Leading Symbols: "("
//		Trailing Symbols: ")"
//		Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewNumberSignTrailingSymbol(
	trailingNumberSymbol string,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewNumberSignTrailingSymbol()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		setTrailingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			[]rune(trailingNumberSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewNumberSignTrailingSymbolRunes
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with a trailing
// number sign symbol character or characters.
//
// Trailing number sign symbol characters can include
// such symbols as plus signs ('+') or minus signs ('-').
//
// Trailing number sign symbol characters are intended
// for use in formatting numeric values displayed in
// number strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingNumberSymbol			[]rune
//
//		An array of runes containing the trailing number
//		sign symbol character or characters used to
//		configure the returned instance of
//		NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include
//		such symbols as plus signs ('+') or minus signs
//		('-').
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new
//		instance of NumStrNumberSymbolSpec, configured
//		with trailing number sign symbol characters, will
//		be returned through this parameter.
//
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Number Sign Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-2: Trailing Number Sign Symbols
//
//	Trailing Symbols: "+"
//	Number String:   "123.456+"
//
// Example-3: Trailing Number Sign Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-4: Trailing Number Sign Symbols
//
//	Trailing Symbols: "-"
//	Number String:   "123.456-"
//
// Example-5: Trailing Number Sign Symbols
//
//	 Leading Symbols: "("
//		Trailing Symbols: ")"
//		Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewNumberSignTrailingSymbolRunes(
	trailingNumberSymbol []rune,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	newNStrNumberSymbolSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewNumberSignTrailingSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		setTrailingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			trailingNumberSymbol,
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
}

//	NewSignedNumDefaultsFrance
//
//	Creates and returns three new instances of default
//	positive signed number symbols, zero value symbols
//	and negative signed number symbols configured in
//	accordance with formatting standards commonly applied
//	in France.
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The positive signed number symbol is empty or blank
//	because under French formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		French Example-1:	123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under French formatting standards,
//	zero numeric values have no number sign symbols.
//
//		French Example-2:	0
//
//	The negative signed number symbol is configured with
//	a leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign will be
//	positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 8
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " -123.45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
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
// ----------------------------------------------------------------
//
// # Defaults
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		French Example-1
//		-1 000 000 000
//
//	The number sign is implied for positive numeric values.
//	Therefore, the positive number sign symbol is set to a
//	blank or empty string ("").
//
//		French Example-3
//		1 000 000 000
//
//	The zero number format has no number sign. Therefore,
//	the positive number sign symbol is set to a blank or
//	empty string ("").
//
//		French Example-4
//			0,0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	positiveNumberSignSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under French signed number
//		formatting standards, the positive number sign
//		('+') is implied and not explicitly displayed.
//
//	zeroNumberSignSymbols			NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under French signed number
//	 	formatting standards, zero numeric values do not
//		have an associated number sign.
//
//	negativeNumberSignSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is configured
//		with a leading minus sign ('-') which will be
//	 	formatted and displayed for all negative number
//		values. This negative number sign will be displayed
//		inside the number field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: -123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewSignedNumDefaultsFrance(
	errorPrefix interface{}) (
	positiveNumberSignSymbols NumStrNumberSymbolSpec,
	zeroNumberSignSymbols NumStrNumberSymbolSpec,
	negativeNumberSignSymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewSignedNumDefaultsFrance()",
		"")

	if err != nil {
		return positiveNumberSignSymbols,
			zeroNumberSignSymbols,
			negativeNumberSignSymbols,
			err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultFrance(
			&positiveNumberSignSymbols,
			&zeroNumberSignSymbols,
			&negativeNumberSignSymbols,
			ePrefix)

	return positiveNumberSignSymbols,
		zeroNumberSignSymbols,
		negativeNumberSignSymbols,
		err
}

// NewSignedNumDefaultsGermany
//
// Creates and returns three new instances of default
// positive signed number symbols, zero value symbols
// and negative signed number symbols configured in
// accordance with formatting standards commonly applied
// in Germany.
//
// The positive signed number symbol is empty or blank
// because under German formatting standards, positive
// number signs are implied and not explicitly
// displayed. Therefore, no leading plus ('+') symbol is
// required.
//
//	German Example:	Positive Numeric Value
//				1.000.000
//
// Likewise, the zero signed number symbol is also empty
// or blank because under German formatting standards,
// zero numeric values have no number sign symbols.
//
//	German Example:	Zero Numeric Value
//					0
//
// The negative signed number symbol is configured with a
// trailing minus sign ('-') meaning that all negative
// numeric values will be suffixed with a trailing minus
// sign ('-'). The negative number sign will be
// positioned inside the number field:
//
//	NumFieldSymPos.InsideNumField()
//		Example:
//			Number Field Length: 11
//			Numeric Value: -123.45
//			Number Symbol: trailing minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " 1.000.000-"
//			Number Field Index:------>01234567890
//			Total Number String Length: 11
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://freeformatter.com/germany-standards-code-snippets.html
//
// https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-1
//		1.000.000-
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		German Example-2
//		1.000.000
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		German Example-3
//			0,00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	positiveNumberSignSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under German signed number
//		formatting standards, the positive number sign
//		('+') is implied and not explicitly displayed.
//
//	zeroNumberSignSymbols			NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under German signed number
//	 	formatting standards, zero numeric values do not
//		have an associated number sign.
//
//	negativeNumberSignSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is configured
//		with a leading minus sign ('-') which will be
//	 	formatted and displayed for all negative number
//		values. This negative number sign will be displayed
//		inside the number field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewSignedNumDefaultsGermany(
	errorPrefix interface{}) (
	positiveNumberSignSymbols NumStrNumberSymbolSpec,
	zeroNumberSignSymbols NumStrNumberSymbolSpec,
	negativeNumberSignSymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewSignedNumDefaultsGermany()",
		"")

	if err != nil {
		return positiveNumberSignSymbols,
			zeroNumberSignSymbols,
			negativeNumberSignSymbols,
			err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultGermany(
			&positiveNumberSignSymbols,
			&zeroNumberSignSymbols,
			&negativeNumberSignSymbols,
			ePrefix)

	return positiveNumberSignSymbols,
		zeroNumberSignSymbols,
		negativeNumberSignSymbols,
		err
}

// NewSignedNumDefaultsUKMinus
//
// Creates and returns three new instances of default
// positive signed number symbols, zero value symbols
// and negative signed number symbols in accordance with
// formatting standards commonly applied in the United
// Kingdom (UK).
//
// The positive signed number symbol is empty or blank
// because under United Kingdom formatting standards,
// positive number signs are implied and not specifically
// displayed. Therefore, no leading plus ('+') symbol is
// required.
//
// Likewise, the zero signed number symbol is also empty
// or blank because under United Kingdom formatting
// standards, zero numeric values have no number sign
// symbols.
//
// The negative signed number symbol is configured with a
// leading minus sign ('-') meaning that all negative
// numeric values will be prefixed with a leading minus
// sign ('-'). The negative number sign will be
// positioned inside the number field:
//
//	NumFieldSymPos.InsideNumField()
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	positiveNumberSignSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under the United States signed number
//		formatting standards, the positive number sign ('+')
//		is implied and not displayed.
//
//	zeroNumberSignSymbols			NumStrNumberSymbolSpec
//
//		One of three returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under the United States signed number
//		formatting standards, zero numeric values do not
//		have an associated number sign.
//
//	negativeNumberSignSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is configured
//		with a leading minus sign ('-') which will be
//	 	formatted and displayed for all negative number
//		values. This negative number sign will be displayed
//		inside the number field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewSignedNumDefaultsUKMinus(
	errorPrefix interface{}) (
	positiveNumberSignSymbols NumStrNumberSymbolSpec,
	zeroNumberSignSymbols NumStrNumberSymbolSpec,
	negativeNumberSignSymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewSignedNumDefaultsUKMinus()",
		"")

	if err != nil {
		return positiveNumberSignSymbols,
			zeroNumberSignSymbols,
			negativeNumberSignSymbols,
			err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultUSMinus(
			&positiveNumberSignSymbols,
			&zeroNumberSignSymbols,
			&negativeNumberSignSymbols,
			ePrefix)

	return positiveNumberSignSymbols,
		zeroNumberSignSymbols,
		negativeNumberSignSymbols,
		err
}

// NewSignedNumDefaultsUSMinus
//
// Creates and returns three new instances of default
// positive signed number symbols, zero value symbols
// and negative signed number symbols in accordance with
// formatting parameters commonly applied in the United
// States.
//
// The positive signed number symbol is empty or blank
// because under United States formatting standards,
// positive number signs are implied and not specifically
// displayed. Therefore, no leading plus ('+') symbol is
// required.
//
// Likewise, the zero signed number symbol is also empty
// or blank because under United States formatting
// standards, zero numeric values have no number sign
// symbols.
//
// The negative signed number symbol is configured with a
// leading minus sign ('-') meaning that all negative
// numeric values will be prefixed with a leading minus
// sign ('-'). The negative number sign will be
// positioned inside the number field:
//
//	NumFieldSymPos.InsideNumField()
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	positiveSignedNumberSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under the United States signed number
//		formatting standards, the positive number sign ('+')
//		is implied and not displayed.
//
//	zeroSignedNumberSymbols			NumStrNumberSymbolSpec
//
//		One of three returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under the United States signed number
//		formatting standards, zero numeric values do not
//		have an associated number sign.
//
//	negativeSignedNumberSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is configured
//		with a leading minus sign ('-') which will be
//	 	formatted and displayed for all negative number
//		values. This negative number sign will be displayed
//		inside the number field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewSignedNumDefaultsUSMinus(
	errorPrefix interface{}) (
	positiveSignedNumberSymbols NumStrNumberSymbolSpec,
	zeroSignedNumberSymbols NumStrNumberSymbolSpec,
	negativeSignedNumberSymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewSignedNumDefaultsUSMinus()",
		"")

	if err != nil {
		return positiveSignedNumberSymbols,
			zeroSignedNumberSymbols,
			negativeSignedNumberSymbols,
			err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultUSMinus(
			&positiveSignedNumberSymbols,
			&zeroSignedNumberSymbols,
			&negativeSignedNumberSymbols,
			ePrefix)

	return positiveSignedNumberSymbols,
		zeroSignedNumberSymbols,
		negativeSignedNumberSymbols,
		err
}

// NewSignedNumDefaultsUSParen
//
// Creates and returns three new instances of default
// positive signed number symbols, zero value symbols
// and negative signed number symbols in accordance with
// formatting parameters commonly applied in the United
// States.
//
// The negative signed number symbols will be configured
// with surrounding parentheses.
//
//	Example: (123.45)
//
// The positive signed number symbol is empty or blank
// because under United States formatting standards,
// positive number signs are implied and not specifically
// displayed. Therefore, no leading plus ('+') symbol is
// required.
//
// Likewise, the zero signed number symbol is also empty
// or blank because under United States formatting
// standards, zero numeric values have no number sign
// symbols.
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
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	positiveSignedNumberSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under the United States signed number
//		formatting standards, the positive number sign ('+')
//		is implied and not displayed.
//
//	zeroSignedNumberSymbols			NumStrNumberSymbolSpec
//
//		One of three returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under the United States signed number
//		formatting standards, zero numeric values do not
//		have an associated number sign.
//
//	negativeSignedNumberSymbols		NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is
//	 	configured with a leading and trailing
//	  	parenthesis ("()") which will be formatted and
//	  	displayed for all negative number values. The
//	  	surrounding parentheses will be displayed
//		inside the number field.
//
//			NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 9
//				Numeric Value: -123.45
//				Number Symbol:
//					Leading and Trailing parenthesis ("()")
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " (123.45)"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//	err								error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewSignedNumDefaultsUSParen(
	errorPrefix interface{}) (
	positiveSignedNumberSymbols NumStrNumberSymbolSpec,
	zeroSignedNumberSymbols NumStrNumberSymbolSpec,
	negativeSignedNumberSymbols NumStrNumberSymbolSpec,
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"NewSignedNumDefaultsUSParen()",
		"")

	if err != nil {
		return positiveSignedNumberSymbols,
			zeroSignedNumberSymbols,
			negativeSignedNumberSymbols,
			err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultUSParen(
			&positiveSignedNumberSymbols,
			&zeroSignedNumberSymbols,
			&negativeSignedNumberSymbols,
			ePrefix)

	return positiveSignedNumberSymbols,
		zeroSignedNumberSymbols,
		negativeSignedNumberSymbols,
		err
}

//	SetCurrencyDefaultEU
//
//	Deletes and reconfigures the data values in the
//	current instance of NumStrNumberSymbolSpec with
//	default currency specifications commonly applied
//	in the European Union.
//
//	The currency symbol used in the European Union is
//	the trailing Euro symbol ('€').
//
//		Example:
//			125 €
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec. All data values will be
//	configured with default currency specifications
//	commonly applied in the European Union.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyDefaultEU(
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyDefaultEU()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecMechanics).
		setCurrencyDefaultsEU(
			nStrNumberSymbolSpec,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec"))
}

//	SetCurrencyDefaultsUKMinusInside
//
//	Deletes and reconfigures the data values in the
//	current instance of NumStrNumberSymbolSpec with
//	the default Currency Symbol specification applied
//	in the United Kingdom (UK).
//
//	The default UK currency symbol is a leading pound
//	sign ('£').
//
//		Example:
//			£ 123.45	Positive Value
//
//	The phrase "MinusInside" in the method name signals
//	that the minus sign will be placed to the right of,
//	or inside, the pound sign ('£').
//
//		Example:
//			£ -123.45	Negative Value
//
//	To generate a default configuration where the minus
//	sign is outside the pound sign ('£'), use method:
//
//	NumStrNumberSymbolSpec.SetCurrencyDefaultsUKMinusOutside()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyDefaultsUKMinusInside(
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyDefaultsUKMinusInside()",
		"")

	if err != nil {
		return err

	}

	return new(numStrNumberSymbolSpecMechanics).
		setCurrencyDefaultsUKMinusInside(
			nStrNumberSymbolSpec,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec"))
}

//	SetCurrencyDefaultsUKMinusOutside
//
//	Deletes and reconfigures the data values in the
//	current instance of NumStrNumberSymbolSpec with
//	the default Currency Symbol specification applied
//	in the United Kingdom (UK).
//
//	The default UK currency symbol is a leading pound
//	sign ('£').
//
//		Example:
//			£ 123.45	Positive Value
//
//	The phrase "MinusOutside" in the method name signals
//	that, for negative values, the minus sign will be
//	placed to the left of, or outside, the pound sign
//	('£').
//
//		Example:
//			-£ 123.45	Negative Value
//
//	To generate a default configuration where the minus
//	sign is inside the pound sign ('£'), use method:
//
//	NumStrNumberSymbolSpec.SetCurrencyDefaultsUKMinusInside()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyDefaultsUKMinusOutside(
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyDefaultsUKMinusOutside()",
		"")

	if err != nil {
		return err

	}

	return new(numStrNumberSymbolSpecMechanics).
		setCurrencyDefaultsUKMinusOutside(
			nStrNumberSymbolSpec,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec"))
}

// SetCurrencyDefaultsUS
//
// Deletes and reconfigures the data values in the
// current instance of NumStrNumberSymbolSpec with
// default currency specifications commonly applied
// in the United States (US).
//
// The currency symbol used in the United States is
// the leading Dollar sign symbol ('$').
//
//	Example:
//		$ 125
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec. All data values will be
//	configured with default currency specifications
//	commonly applied in the United States.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyDefaultsUS(
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyDefaultsUS()",
		"")

	if err != nil {
		return err

	}

	return new(numStrNumberSymbolSpecMechanics).
		setCurrencyDefaultsUS(
			nStrNumberSymbolSpec,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec"))

}

// SetCurrencyLeadingSymbol
//
// Resets and configures a Leading Currency Symbol
// character or characters for the current instance of
// NumStrNumberSymbolSpec.
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
// This method will configure and store Leading Currency
// Symbols for the current instance of
// NumStrNumberSymbolSpec.
//
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading Currency Symbol.
//
//		Example:
//			Leading Currency Symbol: "$ "
//			Formatted Number String: "$ 123.45"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingCurrencySymbol     		string
//
//		A string containing one or more Leading
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Currency Symbols
//
//	Leading Currency Symbols: "$ "
//	Number String:   "$ 123.456"
//
// Example-2: Leading Currency Symbols
//
//	Leading Currency Symbols: "$"
//	Number String:   "$123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyLeadingSymbol(
	leadingCurrencySymbol string,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyLeadingSymbol()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			nStrNumberSymbolSpec,
			[]rune(leadingCurrencySymbol),
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingCurrencySymbol"))
}

// SetCurrencyLeadingSymbolRunes
//
// Resets and configures a Leading Currency Symbol
// character or characters for the current instance of
// NumStrNumberSymbolSpec.
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
// This method will configure and store Leading Currency
// Symbols for the current instance of
// NumStrNumberSymbolSpec.
//
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading Currency Symbol.
//
//		Example:
//			Leading Currency Symbol: []rune{'$',' '}
//			Formatted Number String: "$ 123.45"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
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
//		This empty interface must be convertible to one of
//		the following types:
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Currency Symbols
//
//	Leading Currency Symbols: []rune{'$',' '}
//	Number String:   "$ 123.456"
//
// Example-2: Leading Currency Symbols
//
//	Leading Currency Symbols: []rune{'$'}
//	Number String:   "$123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyLeadingSymbolRunes(
	leadingCurrencySymbol []rune,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyLeadingSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			nStrNumberSymbolSpec,
			leadingCurrencySymbol,
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingCurrencySymbol"))
}

// SetCurrencyLeadingTrailingSymbols
//
// Resets and configures both Leading and Trailing
// Currency Symbol characters for the current instance of
// NumStrNumberSymbolSpec.
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
// Leading currency symbols are prefixed or prepended
// at the beginning of a number string while trailing
// currency symbols are suffixed or appended at the end
// of a number string.
//
//	Example Leading Currency Sign : $123.45
//
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
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading and/or Trailing Currency Symbols.
//
//		Example:
//			Leading Currency Symbol: "$ "
//			Trailing Currency Symbol: " $"
//			Formatted Number String: "$ 123.45 $"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingCurrencySymbol     		string
//
//		A string containing one or more Leading
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
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
//		'trailingCurrencySymbol' are both submitted as
//		empty strings (zero character length), an error
//		will be returned.
//
//	trailingCurrencySymbol     		string
//
//		A string containing one or more Trailing
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
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
//
//		If 'leadingCurrencySymbol' and
//		'trailingCurrencySymbol' are both submitted as
//		empty strings (zero character length), an error
//		will be returned.
//
//	currencyFieldSymbolPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Currency
//		Symbol ('leadingCurrencySymbol') and the Trailing
//		Currency Symbol ('trailingCurrencySymbol')
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
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
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right
//				Formatted Number String: " $123.45$"
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:------>012345678901
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
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:------>012345678901
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading and Trailing Currency Symbols
//
//	Leading Currency Symbols: "$ "
//	Trailing Currency Symbols: " $"
//	Number String:   "$ 123.456 $"
//
// Example-2: Leading and Trailing Currency Symbols
//
//	Leading Currency Symbols: "$"
//	Trailing Currency Symbols: "$"
//	Number String:   "$123.456$"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyLeadingTrailingSymbols(
	leadingCurrencySymbol string,
	trailingCurrencySymbol string,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyLeadingSymbol()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingTrailingCurrencySymbols(
			nStrNumberSymbolSpec,
			[]rune(leadingCurrencySymbol),
			[]rune(trailingCurrencySymbol),
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingCurrencySymbol-trailingCurrencySymbol"))
}

// SetCurrencyLeadingTrailingSymbolRunes
//
// Resets and configures both Leading and Trailing
// Currency Symbol characters for the current instance of
// NumStrNumberSymbolSpec.
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
// Leading currency symbols are prefixed or prepended
// at the beginning of a number string while trailing
// currency symbols are suffixed or appended at the end
// of a number string.
//
//	Example Leading Currency Sign : $123.45
//
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
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading and/or Trailing Currency Symbols.
//
//		Example:
//			Leading Currency Symbol: []rune{'$',' '}
//			Trailing Currency Symbol: []rune{' ','$'}
//			Formatted Number String: "$ 123.45 $"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Leading Currency Symbol characters can include
//		such symbols as the dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Leading Currency Symbols are prefixed or
//		prepended to the beginning of number strings
//		containing currency numeric values.
//
//				Example: $125.34
//
//		If both 'leadingCurrencySymbol' and
//		'trailingCurrencySymbol' are submitted as empty
//		or zero length rune arrays, an error will be
//		returned.
//
//	trailingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
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
//
//		If both 'leadingCurrencySymbol' and
//		'trailingCurrencySymbol' are submitted as empty
//		or zero length rune arrays, an error will be
//		returned.
//
//	currencyFieldSymbolPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Currency
//		Symbol ('leadingCurrencySymbol') and Trailing
//		Currency Symbol ('trailingCurrencySymbol')
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
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
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//			    Number Text Justification: Right
//				Formatted Number String: " $123.45$"
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 12
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  $123.45$  "
//				Number Field Index:------>012345678901
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
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "$  123.45  $"
//				Number Field Index:------>012345678901
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading and Trailing Currency Symbols
//
//	Leading Currency Symbols: []rune{'$',' '}
//	Trailing Currency Symbols: []rune{' ','$'}
//	Number String:   "$ 123.456 $"
//
// Example-2: Leading and Trailing Currency Symbols
//
//	Leading Currency Symbols: []rune{'$'}
//	Trailing Currency Symbols: []rune{'$'}
//	Number String:   "$123.456$"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyLeadingTrailingSymbolRunes(
	leadingCurrencySymbol []rune,
	trailingCurrencySymbol []rune,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyLeadingTrailingSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingTrailingCurrencySymbols(
			nStrNumberSymbolSpec,
			leadingCurrencySymbol,
			trailingCurrencySymbol,
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingCurrencySymbol-trailingCurrencySymbol"))
}

// SetCurrencyNumSignRelPos
//
// Deletes and resets the value of the value of the
// Currency Number Sign Relative Position member variable
// contained in an instance of NumStrNumberSymbolSpec
// passed as an input parameter.
//
// Currency Number Sign Relative Position controls the
// positioning of currency symbols relative to number
// signs associated with numeric values formatted in a
// number string.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numSymbolSpec				*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		The Currency Number Sign Relative Position member
//		variable contained in this instance
//		(currencyNumSignRelativePos) will be deleted
//		and reset to the value specified by input parameter,
//		'currencyNumSignRelPos'.
//
//	currencyNumSignRelPos		CurrencyNumSignRelativePosition
//
//		This parameter is used exclusively by Currency
//		Symbol Specifications.
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyNumSignRelPos(
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyNumSignRelPos()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecMolecule).
		setCurrencyNumSignRelPos(
			nStrNumberSymbolSpec,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec"))
}

// SetCurrencyTrailingSymbol
//
// Resets and configures the Trailing Currency Symbol
// characters for the current instance of
// NumStrNumberSymbolSpec.
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
// Trailing Currency Symbols are suffixed or appended
// to the end of a number string.
//
//	Example: 125.34€
//
// Currency symbol characters are intended for use in
// formatting currency numeric values displayed in number
// strings.
//
// This method will configure and store Trailing Currency
// Symbols for the current instance of
// NumStrNumberSymbolSpec.
//
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Trailing Currency Symbol.
//
//		Example:
//			Trailing Currency Symbol: " €"
//			Formatted Number String: "123.45 €"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingCurrencySymbol     		string
//
//		A string containing one or more Trailing
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
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
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 11
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  123.45€  "
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
//				Trailing Currency Symbol: Dollar sign ('€')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "  123.45€"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "  123.45  €"
//				Number Field Index:------>01234567890
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
//		  	Numeric Value: -123.45
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Currency Symbols
//
//	Trailing Currency Symbols: " €"
//	Number String:   "123.456 €"
//
// Example-2: Trailing Currency Symbols
//
//	Trailing Currency Symbols: "€"
//	Number String:   "123.456€"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyTrailingSymbol(
	trailingCurrencySymbol string,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyTrailingSymbol()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setTrailingCurrencySymbol(
			nStrNumberSymbolSpec,
			[]rune(trailingCurrencySymbol),
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingCurrencySymbol"))

}

// SetCurrencyTrailingSymbolRunes
//
// Resets and configures the Trailing Currency Symbol
// characters for the current instance of
// NumStrNumberSymbolSpec.
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
// Trailing Currency Symbols are suffixed or appended
// to the end of a number string.
//
//	Example: 125.34€
//
// Currency symbol characters are intended for use in
// formatting currency numeric values displayed in number
// strings.
//
// This method will configure and store Trailing Currency
// Symbols for the current instance of
// NumStrNumberSymbolSpec.
//
// This method is designed exclusively for configuring
// Currency Symbols and NOT Number Sign Symbols.
//
// To configure Number Sign Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	(1)	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Trailing Currency Symbol.
//
//		Example:
//			Trailing Currency Symbol: " €"
//			Formatted Number String: "123.45 €"
//
//	(2)	NumberFieldSymbolPosition Conflicts
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
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
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
//				Number Field Index:------>012345679
//				Total Number String Length: 10
//
//			Example-2:
//				Number Field Length: 11
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Inside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: "  123.45€  "
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
//				Trailing Currency Symbol: Dollar sign ('€')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "  123.45€"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Trailing Currency Symbol: Dollar sign ('€')
//				Number Symbol Position: Outside Number Field
//			    Number Text Justification: Centered
//				Formatted Number String: "  123.45  €"
//				Number Field Index:------>01234567890
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
//		  	Numeric Value: -123.45
//			Minus Sign Number Field Symbol Position:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Field Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:------>012345678
//			Total Number String Length: 9
//
//			Currency Symbol is Formatted OUTSIDE
//			the Number Field.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Currency Symbols
//
//	Trailing Currency Symbols: []rune{' ','€'}
//	Number String:   "123.456 €"
//
// Example-2: Trailing Currency Symbols
//
//	Trailing Currency Symbols: []rune{'€'}
//	Number String:   "123.456€"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetCurrencyTrailingSymbolRunes(
	trailingCurrencySymbol []rune,
	currencyFieldSymbolPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetCurrencyTrailingSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecNanobot).
		setTrailingCurrencySymbol(
			nStrNumberSymbolSpec,
			trailingCurrencySymbol,
			currencyFieldSymbolPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingCurrencySymbol"))

}

// SetNOP - Resets and configures the current instance
// of NumStrNumberSymbolSpec as a NOP.
//
// NOP is a computer science term which stands for
// 'No Operation' meaning it performs no operations and
// serves an empty placeholder.
//
// This method is often used to configure implied
// positive number signs in number string formatting. The
// positive number sign, plus ('+'), is often implied for
// positive numeric values and is not displayed in the
// number string.
//
// Call this method when creating implied positive number
// signs for number strings.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNOP() {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	new(numStrNumberSymbolSpecMolecule).empty(
		nStrNumberSymbolSpec)

	return
}

// SetNumberSignLeadingSymbol
//
// Resets and configures a leading number sign symbol
// character or characters for the current instance of
// NumStrNumberSymbolSpec.
//
// Leading number sign symbol characters can include such
// symbols as plus signs ('+') or minus signs ('-').
//
// Leading number sign symbol characters are intended for
// use in formatting numeric values displayed in number
// strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading Number Sign Symbol.
//
//		Example:
//			Leading Number Sign Symbol: "- "
//			Formatted Number String: "- 123.45"
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSignSymbol			string
//
//		A string containing the leading number sign
//		symbol character or characters used to configure
//	 	the current instance of NumStrNumberSymbolSpec.
//
//		Leading number sign symbol characters can include
//		such symbols as plus signs ('+') or minus signs
//		('-').
//
//		Currency symbols are configured separately using
//		other methods.
//
//		If 'leadingNumberSignSymbol' is submitted as an
//		empty string containing zero characters, an error
//		will be returned.
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//		NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Right
//				Formatted Number String: " -123.45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		NumFieldSymPos.OutsideNumField()
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: leading minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Right
//		     	Formatted Number String: "-  123.45"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
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
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Sign Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Sign Symbols
//
//	Leading Symbols: "+"
//	Number String:   "+123.456"
//
// Example-3: Leading Number Sign Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-4: Leading Number Sign Symbols
//
//	Leading Symbols: "-"
//	Number String:   "-123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignLeadingSymbol(
	leadingNumberSignSymbol string,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetNumberSignLeadingSymbol()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		setLeadingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			[]rune(leadingNumberSignSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSignSymbol"))

	return err
}

// SetNumberSignLeadingSymbolRunes
//
// Resets and configures a leading number sign symbol
// character or characters for the current instance of
// NumStrNumberSymbolSpec.
//
// Leading number sign symbol characters can include such
// symbols as plus signs ('+') or minus signs ('-').
//
// Leading number symbol characters are intended for use
// in formatting numeric values displayed in number
// strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading Number Sign Symbol.
//
//		Example:
//			Leading Number Sign Symbol: "- "
//			Formatted Number String: "- 123.45"
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSignSymbol			[]rune
//
//		An array of runes containing the leading number
//	 	sign symbol character or characters used to
//	 	configure the current instance of
//	 	NumStrNumberSymbolSpec.
//
//		Leading number sign symbol characters can include
//		such symbols as plus signs ('+') or minus signs
//		('-').
//
//		Currency symbols are configured separately using
//		other methods.
//
//		If 'leadingNumberSignSymbol' is submitted as an
//		empty string containing zero characters, an error
//		will be returned.
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
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
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:------>0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Sign Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Sign Symbols
//
//	Leading Symbols: "+"
//	Number String:   "+123.456"
//
// Example-3: Leading Number Sign Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-4: Leading Number Sign Symbols
//
//	Leading Symbols: "-"
//	Number String:   "-123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignLeadingSymbolRunes(
	leadingNumberSignSymbol []rune,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetNumberSignLeadingSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		setLeadingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			leadingNumberSignSymbol,
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSignSymbol"))

	return err
}

// SetNumberSignLeadingTrailingSymbol
//
// Resets and configures both Leading and Trailing number
// sign symbol characters for the current instance of
// NumStrNumberSymbolSpec.
//
// Leading and Trailing number sign symbol characters can
// include such symbols as plus signs ('+') or minus
// signs ('-').
//
// Leading and Trailing number symbols are intended for
// use in formatting numeric values displayed in number
// strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading and/or Trailing Number Sign Symbols.
//
//		Example-1:
//			Leading Number Sign Symbol: "- "
//			Formatted Number String: "- 123.45"
//
//		Example-2:
//			Trailing Number Sign Symbol: " -"
//			Formatted Number String: "123.45 -"
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSignSymbol     	string
//
//		A string containing the leading number sign
//		symbol character or characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Leading number sign symbol characters can include
//		such symbols as plus signs ('+') or minus signs
//		('-').
//
//		Currency symbols are configured separately using
//		other methods.
//
//		If 'leadingNumberSignSymbol' is submitted as an
//		empty string containing zero characters, an error
//		will be returned.
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Sign Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//		NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Right
//				Formatted Number String: " -123.45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: leading minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Right
//		     	Formatted Number String: "-  123.45"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.OutsideNumField()' specification,
//			the final length of the number string is greater than
//			the Number Field length.
//
//	trailingNumberSignSymbol     	string
//
//		A string containing the trailing number sign
//		symbol character or characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Trailing number sign symbol characters can
//		include	such symbols as plus signs ('+') or minus
//		signs ('-').
//
//		Currency symbols are configured separately using
//		other methods.
//
//		If 'trailingNumberSignSymbol' is submitted as an
//		empty string containing zero characters, an error
//		will be returned.
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//		NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Right
//				Formatted Number String: " 123.45-"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.InsideNumField()' specification,
//			the final length of the number string is defined by the
//			Number Field length.
//
//		NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: trailing minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Right
//		     	Formatted Number String: "  123.45-"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
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
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Sign Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Sign Symbols
//
//	Leading Symbols: "+"
//	Number String:   "+123.456"
//
// Example-3: Leading Number Sign Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-4: Leading Number Sign Symbols
//
//	Leading Symbols: "-"
//	Number String:   "-123.456"
//
// Example-5: Trailing Number Sign Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-6: Trailing Number Sign Symbols
//
//	Trailing Symbols: "+"
//	Number String:   "123.456+"
//
// Example-7: Trailing Number Sign Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-8: Trailing Number Symbols
//
//	Trailing Symbols: "-"
//	Number String:   "123.456-"
//
// Example-9: Trailing Number Symbols
//
//	Leading Number Sign Symbols: "("
//	Trailing Number Sign Symbols: ")"
//	Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignLeadingTrailingSymbol(
	leadingNumberSignSymbol string,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	trailingNumberSignSymbol string,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetNumberSignLeadingTrailingSymbol()",
		"")

	if err != nil {
		return err
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	err = nStrNumSymSpecNanobot.
		setLeadingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			[]rune(leadingNumberSignSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSignSymbol"))

	if err != nil {
		return err
	}

	err = nStrNumSymSpecNanobot.
		setTrailingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			[]rune(trailingNumberSignSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSignSymbol"))

	return err
}

// SetNumberSignLeadingTrailingSymbolRunes
//
// Resets and configures both Leading and Trailing number
// sign symbol characters for the current instance of
// NumStrNumberSymbolSpec.
//
// Leading and Trailing number sign symbol characters can
// include such symbols as plus signs ('+') or minus
// signs ('-').
//
// Leading and Trailing number symbols are intended for
// use in formatting numeric values displayed in number
// strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Leading and/or Trailing Number Sign Symbols.
//
//		Example-1:
//			Leading Number Sign Symbol: "- "
//			Formatted Number String: "- 123.45"
//
//		Example-2:
//			Trailing Number Sign Symbol: " -"
//			Formatted Number String: "123.45 -"
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSignSymbol     	[]rune
//
//		A string containing the leading number sign
//		symbol character or characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Leading number sign symbol characters can include
//		such symbols as plus signs ('+') or minus signs
//		('-').
//
//		Currency symbols are configured separately using
//		other methods.
//
//		If 'leadingNumberSignSymbol' is submitted as an
//		empty string containing zero characters, an error
//		will be returned.
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//		NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Right
//				Formatted Number String: " -123.45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//				Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.InsideNumField()'
//			specification, the final length of the
//			number string is defined by the Number
//			Field length.
//
//		NumFieldSymPos.OutsideNumField()
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: leading minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Right
//		     	Formatted Number String: "-  123.45"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.OutsideNumField()'
//			specification, the final length of the
//			number string is greater than the Number
//			Field length.
//
//	trailingNumberSignSymbol		[]rune
//
//		An array of runes containing the trailing number
//		sign symbol character or characters used to
//		configure the returned instance of
//		NumStrNumberSymbolSpec.
//
//		Trailing number sign symbol characters can
//		include	such symbols as plus signs ('+') or minus
//		signs ('-').
//
//		Currency symbols are configured separately using
//		other methods.
//
//		If 'trailingNumberSignSymbol' is submitted as an
//		empty string containing zero characters, an error
//		will be returned.
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//		NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Right
//				Formatted Number String: " 123.45-"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.InsideNumField()'
//			specification, the final length of the number
//			string is defined by the Number Field length.
//
//		NumFieldSymPos.OutsideNumField()
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: trailing minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Right
//		     	Formatted Number String: "  123.45-"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.OutsideNumField()'
//			specification, the final length of the number
//			string is greater than the Number Field length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Sign Symbols
//
//	Leading Symbols: []rune{'+', ' '}
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Sign Symbols
//
//	Leading Symbols: []rune{'+'}
//	Number String:   "+123.456"
//
// Example-3: Leading Number Sign Symbols
//
//	Leading Symbols: []rune{'-', ' '}
//	Number String:   "- 123.456"
//
// Example-4: Leading Number Sign Symbols
//
//	Leading Symbols: []rune{'-'}
//	Number String:   "-123.456"
//
// Example-5: Trailing Number Sign Symbols
//
//	Trailing Symbols: []rune{' ', '+'}
//	Number String:   "123.456 +"
//
// Example-6: Trailing Number Sign Symbols
//
//	Trailing Symbols: []rune{'+'}
//	Number String:   "123.456+"
//
// Example-7: Trailing Number Sign Symbols
//
//	Trailing Symbols: []rune{' ', '-'}
//	Number String:   "123.456 -"
//
// Example-8: Trailing Number Symbols
//
//	Trailing Symbols: []rune{'-'}
//	Number String:   "123.456-"
//
// Example-9: Trailing Number Symbols
//
//	Leading Number Sign Symbols: []rune{'('}
//	Trailing Number Sign Symbols: []rune{')'}
//	Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignLeadingTrailingSymbolRunes(
	leadingNumberSignSymbol []rune,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	trailingNumberSignSymbol []rune,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetNumberSignLeadingTrailingSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

	err = nStrNumSymSpecNanobot.
		setLeadingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			leadingNumberSignSymbol,
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSignSymbol"))

	if err != nil {
		return err
	}

	err = nStrNumSymSpecNanobot.
		setTrailingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			trailingNumberSignSymbol,
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSignSymbol"))

	return err
}

// SetNumberSignTrailingSymbol
//
// Resets and configures a trailing number sign symbol
// character or characters for the current instance of
// NumStrNumberSymbolSpec
//
// Trailing number sign symbol characters can include
// such symbols as plus signs ('+') or minus signs ('-').
//
// Trailing number sign symbol characters are intended
// for use in formatting numeric values displayed in
// number strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Trailing Number Sign Symbols.
//
//		Example-1:
//			Trailing Number Sign Symbol: " -"
//			Formatted Number String: "123.45 -"
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	trailingNumberSignSymbol     	string
//
//		A string containing the trailing number sign
//		symbol character or characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Trailing number sign symbol characters can
//		include	such symbols as plus signs ('+') or minus
//		signs ('-').
//
//		Currency symbols are configured separately using
//		other methods.
//
//		If 'trailingNumberSignSymbol' is submitted as an
//		empty string containing zero characters, an error
//		will be returned.
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//		NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Right
//				Formatted Number String: " 123.45-"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.InsideNumField()'
//			specification, the final length of the number
//			string is defined by the Number Field length.
//
//		NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: trailing minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Right
//		     	Formatted Number String: "  123.45-"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.OutsideNumField()'
//			specification, the final length of the number
//			string is greater than the Number Field
//			length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Number Sign Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-2: Trailing Number Sign Symbols
//
//	Trailing Symbols: "+"
//	Number String:   "123.456+"
//
// Example-3: Trailing Number Sign Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-4: Trailing Number Symbols
//
//	Trailing Symbols: "-"
//	Number String:   "123.456-"
//
// Example-5: Trailing Number Symbols
//
//	Leading Number Sign Symbols: "("
//	Trailing Number Sign Symbols: ")"
//	Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignTrailingSymbol(
	trailingNumberSignSymbol string,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetNumberSignTrailingSymbol()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		setTrailingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			[]rune(trailingNumberSignSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"trailingNumberSignSymbol"))

	return err
}

// SetNumberSignTrailingSymbolRunes
//
// Resets and configures a trailing number sign symbol
// character or characters for the current instance of
// NumStrNumberSymbolSpec.
//
// Trailing number sign symbol characters can include
// such symbols as plus signs ('+') or minus signs ('-').
//
// Trailing number sign symbol characters are intended
// for use in formatting numeric values displayed in
// number strings.
//
// This method is designed exclusively for configuring
// Number Sign Symbols and NOT Currency Symbols.
//
// To configure Currency Symbols, use other methods.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Interior Spaces
//
//		It may be advantageous to include spaces with the
//		Trailing Number Sign Symbols.
//
//		Example-1:
//			Trailing Number Sign Symbol: " -"
//			Formatted Number String: "123.45 -"
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingNumberSignSymbol		[]rune
//
//		An array of runes containing the trailing number
//		sign symbol character or characters used to
//		configure the current instance of
//		NumStrNumberSymbolSpec.
//
//		Trailing number sign symbol characters can
//		include	such symbols as plus signs ('+') or minus
//		signs ('-').
//
//		Currency symbols are configured separately using
//		other methods.
//
//		If 'trailingNumberSignSymbol' is submitted as an
//		empty string containing zero characters, an error
//		will be returned.
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//		NumFieldSymPos.InsideNumField()
//
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Right
//				Formatted Number String: " 123.45-"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Inside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.InsideNumField()'
//			specification, the final length of the number
//			string is defined by the Number Field length.
//
//		NumFieldSymPos.OutsideNumField()
//
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: trailing minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Right
//		     	Formatted Number String: "  123.45-"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//		     	Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:------>0123456789
//				Total Number String Length: 10
//
//			For the 'NumFieldSymPos.OutsideNumField()'
//			specification, the final length of the number
//			string is greater than the Number Field
//			length.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
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
//	err								error
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Number Sign Symbols
//
//	Trailing Symbols: []rune{' ', '+'}
//	Number String:   "123.456 +"
//
// Example-2: Trailing Number Sign Symbols
//
//	Trailing Symbols: []rune{'+'}
//	Number String:   "123.456+"
//
// Example-3: Trailing Number Sign Symbols
//
//	Trailing Symbols: []rune{' ', '-'}
//	Number String:   "123.456 -"
//
// Example-4: Trailing Number Symbols
//
//	Trailing Symbols: []rune{'-'}
//	Number String:   "123.456-"
//
// Example-5: Trailing Number Symbols
//
//	Leading Number Sign Symbols: []rune{'('}
//	Trailing Number Sign Symbols: []rune{')'}
//	Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignTrailingSymbolRunes(
	trailingNumberSignSymbol []rune,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	err error) {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetNumberSignTrailingSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecMolecule).
		setTrailingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			trailingNumberSignSymbol,
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"trailingNumberSignSymbol"))

	return err
}

//	SetSignedNumDefaultsFrance
//
//	Receives pointers to three instances of
//	NumStrNumberSymbolSpec and proceeds to configure
//	positive, zero and negative number sign symbols
//	commonly applied to formatting signed number values
//	according to French standards.
//
//	The number sign symbols are configured using default
//	French signed number formatting values.
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
// ----------------------------------------------------------------
//
// # Defaults
//
//	The negative signed number symbol is configured with a
//	leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign will be
//	positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 8
//				Numeric Value: -123,45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " -123,45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		French Example:	Positive Numeric Value
//		1 000 000 000
//
//	The zero number format has no number sign. Therefore,
//	the positive number sign symbol is set to a blank or
//	empty string ("").
//
//		French Example:	Zero Numeric Value
//			0,0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance will be configured as empty or
//		blank because under French signed number
//		formatting standards, the positive number sign
//		('+') is implied and not explicitly displayed.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	zeroNumberSignSymbols			*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance is configured as empty or blank
//		because under French signed number formatting
//		standards, zero numeric values do not have an
//		associated number sign.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	negativeNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance is configured with a leading minus
//		sign ('-') which will be formatted and displayed
//		for all negative number	values. The negative
//		number sign will be displayed inside the number
//		field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: -123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as a leading
//		minus sign ('-') in accordance with French
//		Signed Number formatting standards.
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetSignedNumDefaultsFrance(
	positiveNumberSignSymbols *NumStrNumberSymbolSpec,
	zeroNumberSignSymbols *NumStrNumberSymbolSpec,
	negativeNumberSignSymbols *NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetSignedNumDefaultsFrance()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultFrance(
			positiveNumberSignSymbols,
			zeroNumberSignSymbols,
			negativeNumberSignSymbols,
			ePrefix)
}

//	SetSignedNumDefaultsGermany
//
//	Receives pointers to three instances of
//	NumStrNumberSymbolSpec and proceeds to configure
//	positive, zero and negative number sign symbols
//	commonly applied to formatting signed number values
//	according to German standards.
//
//	The number sign symbols are configured using default
//	German signed number formatting values.
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
// # Defaults
//
//	The negative signed number symbol is configured with a
//	trailing minus sign ('-') meaning that all negative
//	numeric values will be suffixed with a trailing minus
//	sign ('-'). The negative number sign will be
//	positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 11
//				Numeric Value: -123.45
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " 1.000.000-"
//				Number Field Index:------>01234567890
//				Total Number String Length: 11
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		German Example:	Positive Numeric Value
//			1.000.000
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		German Example:	Zero Numeric Value
//			0,00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance will be configured as empty or
//		blank because under German signed number
//		formatting standards, the positive number sign
//		('+') is implied and not explicitly displayed.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	zeroNumberSignSymbols			*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance is configured as empty or blank
//		because under German signed number formatting
//		standards, zero numeric values do not have an
//		associated number sign.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	negativeNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance is configured with a trailing minus
//		sign ('-') which will be formatted and displayed
//		for all negative number	values. The negative
//		number sign will be displayed inside the number
//		field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: -123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as a trailing
//		minus sign ('-') in accordance with German
//		Signed Number formatting standards.
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetSignedNumDefaultsGermany(
	positiveNumberSignSymbols *NumStrNumberSymbolSpec,
	zeroNumberSignSymbols *NumStrNumberSymbolSpec,
	negativeNumberSignSymbols *NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetSignedNumDefaultsGermany()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultGermany(
			positiveNumberSignSymbols,
			zeroNumberSignSymbols,
			negativeNumberSignSymbols,
			ePrefix)
}

//	SetSignedNumDefaultsUKMinus
//
//	Receives pointers to three instances of
//	NumStrNumberSymbolSpec and proceeds to configure
//	positive, zero and negative number sign symbols
//	commonly applied to formatting signed number values
//	according to the United Kingdom (UK) standards.
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		Example: -123//
//
// //	The number sign symbols are configured using default
// //	UK signed number formatting values.
//
//	The positive signed number symbol is empty or blank
//	because under United Kingdom formatting standards,
//	positive number signs are implied and not specifically
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		UK Example:	Positive Numeric Value
//					123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under United Kingdom formatting
//	standards, zero numeric values have no number sign
//	symbols.
//
//		UK Example:	Zero Numeric Value
//						0
//
//	The negative signed number symbol is configured with a
//	leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign will be
//	positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 8
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " -123.45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
// ----------------------------------------------------------------
//
// # Reference:
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
// # Defaults
//
//	The negative signed number symbol is configured with a
//	leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign will be
//	positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 8
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " -123.45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		UK Example:	Positive Numeric Value
//				1,000,000
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		UK Example:	Zero Numeric Value
//						0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance will be configured as empty or
//		blank because under UK signed number formatting
//		standards, the positive number sign ('+') is
//		implied and not explicitly displayed.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	zeroNumberSignSymbols			*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance is configured as empty or blank
//		because under UK signed number formatting
//		standards, zero numeric values do not have an
//		associated number sign.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	negativeNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//
//
//		This instance is configured with a leading minus
//		sign ('-') which will be formatted and displayed
//		for all negative number values. This negative
//		number sign will be displayed inside the number
//		field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: -123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//	errorPrefix					interface{}
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetSignedNumDefaultsUKMinus(
	positiveNumberSignSymbols *NumStrNumberSymbolSpec,
	zeroNumberSignSymbols *NumStrNumberSymbolSpec,
	negativeNumberSignSymbols *NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetSignedNumDefaultsUKMinus()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultUSMinus(
			positiveNumberSignSymbols,
			zeroNumberSignSymbols,
			negativeNumberSignSymbols,
			ePrefix)
}

//	SetSignedNumDefaultsUSMinus
//
//	Receives pointers to three instances of
//	NumStrNumberSymbolSpec and proceeds to configure
//	positive, zero and negative number sign symbols
//	commonly applied to formatting signed number values
//	according to the United States (US) standards.
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		Example Signed Number Negative Value: -123
//
//	The number sign symbols are configured using default
//	US signed number formatting values.
//
//	The positive signed number symbol is empty or blank
//	because under United States formatting standards,
//	positive number signs are implied and not specifically
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		Example Signed Number Positive Value: 123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under United States formatting
//	standards, zero numeric values have no number sign
//	symbols.
//
//		Example Signed Number Zero Value: 0
//
//	The negative signed number symbol is configured with a
//	leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign will be
//	positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " -123.45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The negative signed number symbol is configured with a
//	leading minus sign ('-') meaning that all negative
//	numeric values will be prefixed with a leading minus
//	sign ('-'). The negative number sign will be
//	positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 8
//				Numeric Value: -123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " -123.45"
//				Number Field Index:------>01234567
//				Total Number String Length: 8
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example:	Positive Numeric Value
//				1,000,000
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example:	Zero Numeric Value
//						0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance will be configured as empty or
//		blank because under US signed number formatting
//		standards, the positive number sign ('+') is
//		implied and not explicitly displayed.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	zeroNumberSignSymbols			*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance is configured as empty or blank
//		because under US signed number formatting
//		standards, zero numeric values do not have an
//		associated number sign.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	negativeNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//
//
//		This instance is configured with a leading minus
//		sign ('-') which will be formatted and displayed
//		for all negative number values. This negative
//		number sign will be displayed inside the number
//		field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: -123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//	errorPrefix					interface{}
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetSignedNumDefaultsUSMinus(
	positiveNumberSignSymbols *NumStrNumberSymbolSpec,
	zeroNumberSignSymbols *NumStrNumberSymbolSpec,
	negativeNumberSignSymbols *NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetSignedNumDefaultsUSMinus()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultUSMinus(
			positiveNumberSignSymbols,
			zeroNumberSignSymbols,
			negativeNumberSignSymbols,
			ePrefix)
}

//	SetSignedNumDefaultsUSParen
//
//	Receives pointers to three instances of
//	NumStrNumberSymbolSpec and proceeds to configure
//	positive, zero and negative number sign symbols
//	commonly applied to formatting signed number values
//	according to the United States (US) standards.
//
//	The word 'Paren' in the method name signals that
//	negative numeric values will be configured with a
//	surrounding parentheses ('()').
//
//		Example Signed Number Negative Value: (123)
//
//	The number sign symbols are configured using default
//	US signed number formatting values.
//
//	The positive signed number symbol is empty or blank
//	because under United States formatting standards,
//	positive number signs are implied and not specifically
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		Example Signed Number Positive Value: 123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under United States formatting
//	standards, zero numeric values have no number sign
//	symbols.
//
//		Example Signed Number Zero Value: 0
//
//	The negative signed number symbol is configured with
//	surrounding parentheses ('()') meaning that all
//	negative numeric values will be surrounded with a
//	leading parenthesis sign ('(') and trailing closing
//	parenthesis sing (')'). The negative number signs
//	will be positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 9
//				Numeric Value: -123.45
//				Number Symbol: Surrounding Parentheses ('()')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " (123.45)"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The negative signed number symbol is configured with
//	surrounding parentheses ('()') meaning that all
//	negative numeric values will be surrounded with a
//	leading parenthesis sign ('(') and trailing closing
//	parenthesis sing (')'). The negative number signs
//	will be positioned inside the number field:
//
//		NumFieldSymPos.InsideNumField()
//			Example:
//				Number Field Length: 9
//				Numeric Value: -123.45
//				Number Symbol: Surrounding Parentheses ('()')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " (123.45)"
//				Number Field Index:------>012345678
//				Total Number String Length: 9
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example:	Positive Numeric Value
//				1,000,000
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example:	Zero Numeric Value
//						0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance will be configured as empty or
//		blank because under US signed number formatting
//		standards, the positive number sign ('+') is
//		implied and not explicitly displayed.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	zeroNumberSignSymbols			*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//		This instance is configured as empty or blank
//		because under US signed number formatting
//		standards, zero numeric values do not have an
//		associated number sign.
//
//		All pre-existing data values in this instance
//		will be deleted and reconfigured as an empty
//		or NOP instance.
//
//	negativeNumberSignSymbols		*NumStrNumberSymbolSpec
//
//		A pointer to an instance of
//		NumStrNumberSymbolSpec.
//
//
//
//		This instance is configured with a leading minus
//		sign ('-') which will be formatted and displayed
//		for all negative number values. This negative
//		number sign will be displayed inside the number
//		field.
//
//			NumFieldSymPos.InsideNumField()
//				Example:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//	errorPrefix					interface{}
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetSignedNumDefaultsUSParen(
	positiveNumberSignSymbols *NumStrNumberSymbolSpec,
	zeroNumberSignSymbols *NumStrNumberSymbolSpec,
	negativeNumberSignSymbols *NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolSpec."+
			"SetSignedNumDefaultsUSParen()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultUSParen(
			positiveNumberSignSymbols,
			zeroNumberSignSymbols,
			negativeNumberSignSymbols,
			ePrefix)
}

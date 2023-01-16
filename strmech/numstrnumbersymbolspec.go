package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrNumberSymbolSpec
//
// The Number String Number Symbol Specification,
// 'NumStrNumberSymbolSpec' contains specifications
// used to construct and format numeric values in
// number strings for screen displays, file output
// and printing.
//
// # Background
//
// The Number String Number Symbol Specification
// (Type 'NumStrNumberSymbolSpec') typically defines
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
//		Leading Currency Symbols for Positive Values
//
//		Leading Currency Symbols: "$ "
//		Number String:   "$ 123.456"
//
//	Example-6: Leading Currency Symbols
//		Leading Currency Symbols for Positive Values
//
//		Leading Currency Symbols: "$"
//		Number String:   "$123.456"
//
//	Example-7: Trailing Number Symbols
//		Trailing Number Sign Symbols for Positive Values
//
//		Trailing Number Sign Symbols: " +"
//		Number String:   "123.456 +"
//
//	Example-8: Trailing Number Symbols
//		Trailing Number Symbols for Positive Values
//
//		Trailing Number Sign Symbols: "+"
//		Number String:   "123.456+"
//
//	Example-9: Trailing Number Symbols
//		Trailing Number Symbols for Negative Values
//
//		Trailing Number Sign Symbols: " -"
//		Number String:   "123.456 -"
//
//	Example-10: Trailing Number Symbols
//		Trailing Number Symbols for Negative Values
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
//	Example-12: Leading AND Trailing Number Symbols
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
	//			Number Field Index:       01234567
	//			Total Number String Length: 8
	//
	//		Example-2:
	//			Number Field Length: 10
	//			Numeric Value: 123.45
	//			Number Symbol: before and after parentheses  ('()')
	//			Number Symbol Position: Outside Number Field
	//          Number Text Justification: Centered
	//			Formatted Number String: " (123.45) "
	//			Number Field Index:       0123456789
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
	//			Number Field Index:  012345678
	//			Total Number String Length: 9
	//
	//		Example-4:
	//			Number Field Length: 8
	//			Numeric Value: 123.45
	//			Number Symbol: before and after parentheses  ('()')
	//			Number Symbol Position: Outside Number Field
	//			Formatted Number String: "( 123.45 )"
	//			Number Field Index:  0123456789
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
	//			Number Field Index:       01234567
	//			Total Number String Length: 8
	//
	//		Example-2:
	//			Number Field Length: 10
	//			Numeric Value: 123.45
	//			Number Symbol: before and after parentheses  ('()')
	//			Number Symbol Position: Outside Number Field
	//          Number Text Justification: Centered
	//			Formatted Number String: " (123.45) "
	//			Number Field Index:       0123456789
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
	//			Number Field Index:       012345678
	//			Total Number String Length: 9
	//
	//		Example-4:
	//			Number Field Length: 8
	//			Numeric Value: 123.45
	//			Number Symbol: before and after parentheses  ('()')
	//			Number Symbol Position: Outside Number Field
	//          Number Text Justification: Centered
	//			Formatted Number String: "( 123.45 )"
	//			Number Field Index:       0123456789
	//			Total Number String Length: 10
	//
	//		In this case the final length of the number string
	//		is greater than the Number Field length.

	currencyNumSignRelativePosition CurrencyNumSignRelativePosition
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
	//		"$ -123.45"
	//		"123.45- €"
	//
	// Examples CurrNumSignRelPos.InsideNumSign()
	//
	//	Examples:
	//		"- $123.45"
	//		"123.45€ -"
	//
	// Be Advised -
	//	If the currency symbol is formatted Outside a
	//	Number Field and the number sign symbol is
	//	formatted Inside a Number Field, this
	//	parameter will be ignored.

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
// 'currencyNumSignRelativePosition'.
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

	return nStrNumberSymbolSpec.currencyNumSignRelativePosition
}

// GetLeadingNumberSymbolRunesDto
// Returns a RuneArrayDto containing the leading number symbol
// character or characters.
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
			"NewCurrencyDefaultsEU()",
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
// Returns a rune array containing the leading number symbol
// character or characters.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetLeadingNumberSymbolRunes() []rune {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.leadingNumberSymbols.GetRuneArray()
}

// GetLeadingNumberSymbolStr - Returns a string containing the
// leading number symbol character or characters.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetLeadingNumberSymbolStr() string {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.leadingNumberSymbols.GetCharacterString()
}

//		GetLeadingNumberSymbolPosition
//
//		Returns an instance of NumberFieldSymbolPosition.
//
//		This enumeration value defines the position
//		of the Leading Number Symbol relative to a
//		Number Field in which a number string is
//		displayed. Possible valid values are listed
//		as follows:
//
//		NumFieldSymPos.InsideNumField()
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//				Formatted Number String: " -123.45"
//				Number Field Index:-------01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//	         Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:-------0123456789
//				Total Number String Length: 10
//
//			In this case the final length of the number string
//			is defined by the Number Field length.
//
//		NumFieldSymPos.OutsideNumField()
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: leading minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//		     	Formatted Number String: "-  123.45"
//				Number Field Index:-------012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:-------0123456789
//				Total Number String Length: 10
//
//			In this case the final length of the number string
//			is greater than the Number Field length.
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
// Returns a RuneArrayDto containing the trailing number symbol
// character or characters.
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
// Returns a rune array containing the trailing number symbol
// character or characters.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetTrailingNumberSymbolRunes() []rune {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.trailingNumberSymbols.GetRuneArray()
}

// GetTrailingNumberSymbolStr - Returns a string containing the
// trailing number symbol character or characters.
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) GetTrailingNumberSymbolStr() string {

	if nStrNumberSymbolSpec.lock == nil {
		nStrNumberSymbolSpec.lock = new(sync.Mutex)
	}

	nStrNumberSymbolSpec.lock.Lock()

	defer nStrNumberSymbolSpec.lock.Unlock()

	return nStrNumberSymbolSpec.trailingNumberSymbols.GetCharacterString()
}

//		GetTrailingNumberSymbolPosition
//
//		Returns an instance of NumberFieldSymbolPosition.
//
//		This enumeration value defines the position
//		of the Trailing Number Symbol relative to a
//		Number Field in which a number string is
//		displayed. Possible valid values are listed
//		as follows:
//
//		NumFieldSymPos.InsideNumField()
//			Example-1:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: trailing minus sign ('-')
//				Number Symbol Position: Inside Number Field
//	         Number Text Justification: Right
//				Formatted Number String: " 123.45-"
//				Number Field Index:-------01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//	         Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:-------0123456789
//				Total Number String Length: 10
//
//			In this case the final length of the number string
//			is defined by the Number Field length.
//
//		NumFieldSymPos.OutsideNumField()
//			Example-3:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: trailing minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//	         Number Text Justification: Right
//		     	Formatted Number String: "  123.45-"
//				Number Field Index:-------012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//	         Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:-------0123456789
//				Total Number String Length: 10
//
//			In this case the final length of the number string
//			is greater than the Number Field length.
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

	if nStrNumberSymbolSpec.leadingNumberSymbols.GetRuneArrayLength() == 0 &&
		nStrNumberSymbolSpec.trailingNumberSymbols.GetRuneArrayLength() == 0 {

		return true
	}

	return false
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

// NewCurrencyDefaultsUK
//
// Creates and returns a new instance of
// NumStrNumberSymbolSpec configured with the default UK
// (United Kingdom) currency symbol.
//
// The default Uk currency symbol is a leading pound
// sign.
//
//	Example:
//		£ 123.45
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewCurrencyDefaultsUK(
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
			"NewCurrencyDefaultsUK()",
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
//			Minus Sign NumberFieldSymbolPosition:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:-------012345678
//			Total Number String Length: 9
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with Leading
//		Currency Symbols, will be returned.
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
//			Minus Sign NumberFieldSymbolPosition:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:-------012345678
//			Total Number String Length: 9
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
//				Number Field Index:-------012345679
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
//				Leading Currency Symbol: Dollar sign ('$')
//				Trailing Currency Symbol: Dollar sign ('$')
//			    Number Symbol Position: Outside Number Field
//			    Number Text Justification: Right
//			    Formatted Number String: "$  123.45$"
//				Number Field Index:-------0123456789
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with Leading
//		Currency Symbols, will be returned.
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
//			Minus Sign NumberFieldSymbolPosition:
//				NumFieldSymPos.InsideNumField()
//			Currency Number Symbol Position:
//				NumFieldSymPos.OutsideNumField()
//			Currency Number Sign Relative Position:
//				CurrNumSignRelPos.InsideNumSign()
//			Leading Currency Symbol: Dollar sign ('$')
//			Number Text Justification: Right
//			Formatted Number String: "$ -123.45"
//			Number Field Index:-------012345678
//			Total Number String Length: 9
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with Leading
//		Currency Symbols, will be returned.
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
// This method will configure and store Trailing
// Currency Symbols for the current instance of
// NumStrNumberSymbolSpec.
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
//			Number Field Index:-------012345678
//			Total Number String Length: 9
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with Trailing
//		Currency Symbols, will be returned.
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
//			Number Field Index:-------012345678
//			Total Number String Length: 9
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with Trailing
//		Currency Symbols, will be returned.
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
// NumStrNumberSymbolSpec. This instance is designed to serve as
// an empty placeholder or NOP.
//
// NOP is a computer science term which stands for 'No Operation'
// meaning it performs no operations and serves an empty
// placeholder.
//
// This method is often used to configure implied positive number
// signs in number string formatting. The positive number sign,
// plus ('+'), is often implied for positive numeric values and
// is not displayed in the number string.
//
// Call this method when creating implied positive number signs
// for number strings.
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
// Leading number symbol characters are intended for use in
// formatting numeric values displayed in number strings.
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
//		such symbols as plus signs ('+') and minus signs
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
//			Number Field Index:-------01234567
//			Total Number String Length: 8
//
//		Example-2:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Sign Symbol: leading plus sign ("+")
//			Number Symbol Position: Inside Number Field
//			Number Text Justification: Right
//			Formatted Number String: " +123.45"
//			Number Field Index:-------01234567
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
//			Number Field Index:-------012345678
//			Total Number String Length: 9
//
//		Example-4:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading plus sign ('+')
//			Number Symbol Position: Outside Number Field
//			Number Text Justification: Right
//			Formatted Number String: "+  123.45"
//			Number Field Index:-------012345678
//			Total Number String Length: 9
//
//			For the 'NumFieldSymPos.OutsideNumField()'
//			specification, the final length of the number
//			string is greater than the Number Field
//			length.
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with leading
//		number symbol characters, will be returned through
//		this parameter.
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
//		such symbols as plus signs ('+') and minus signs
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
//			Number Field Index:-------01234567
//			Total Number String Length: 8
//
//		Example-2:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Sign Symbol: leading plus sign ("+")
//			Number Symbol Position: Inside Number Field
//		   	Number Text Justification: Right
//			Formatted Number String: " +123.45"
//		    Number Field Index:-------01234567
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
//			Number Field Index:-------012345678
//			Total Number String Length: 9
//
//		Example-4:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading plus sign ('+')
//			Number Symbol Position: Outside Number Field
//			Number Text Justification: Right
//			Formatted Number String: "+  123.45"
//			Number Field Index:-------012345678
//			Total Number String Length: 9
//
//			For the 'NumFieldSymPos.OutsideNumField()'
//			specification, the final length of the number
//			string is greater than the Number Field
//			length.
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with leading
//		number symbol characters, will be returned through
//		this parameter.
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

// NewNumberSignLeadingTrailingSymbol - Creates and returns a new
// instance of NumStrNumberSymbolSpec configured with both
// leading and trailing number symbols.
//
// Leading and Trailing number symbol characters can include
// such symbols as plus signs ('+'), minus signs ('-') and/or
// currency symbols ('$').
//
// Leading and Trailing number symbols are intended for use
// in formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol     		string
//
//		A string containing the leading number symbol
//		character or characters used to configure the
//		returned instance of NumStrNumberSymbolSpec.
//
//		Leading number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//		and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	trailingNumberSymbol     		string
//
//		A string containing the trailing number symbol
//		character or characters used to configure the returned
//		instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with both leading
//		and trailing number symbol characters, will be returned
//	 	through this parameter.
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Symbols
//
//	Leading Symbols: "$+"
//	Number String:   "$+123.456"
//
// Example-3: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-4: Leading Number Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-5: Leading Number Symbols
//
//	Leading Symbols: "$-"
//	Number String:   "$-123.456"
//
// Example-6: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-7: Trailing Number Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-8: Trailing Number Symbols
//
//	Trailing Symbols: "+$"
//	Number String:   "123.456+$"
//
// Example-9: Trailing Number Symbols
//
//	Trailing Symbols: "$"
//	Number String:   "123.456$"
//
// Example-10: Trailing Number Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-11: Trailing Number Symbols
//
//	Trailing Symbols: "-$"
//	Number String:   "123.456-$"
//
// Example-12: Trailing Number Symbols
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

// NewNumberSignLeadingTrailingSymbolRunes - Creates and returns a
// new instance of NumStrNumberSymbolSpec configured with
// both leading and trailing number symbols.
//
// Leading and Trailing number symbol characters can include
// such symbols as plus signs ('+'), minus signs ('-') and/or
// currency symbols ('$').
//
// Leading and Trailing number symbols are intended for use
// in formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol				[]rune
//
//		An array of runes containing the leading number
//		symbol character or characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//		and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	trailingNumberSymbol			[]rune
//
//		An array of runes containing the trailing number symbol
//		character or characters used to configure the returned
//		instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with both leading
//		and trailing number symbol characters, will be returned
//	 	through this parameter.
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Symbols
//
//	Leading Symbols: "$+"
//	Number String:   "$+123.456"
//
// Example-3: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-4: Leading Number Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-5: Leading Number Symbols
//
//	Leading Symbols: "$-"
//	Number String:   "$-123.456"
//
// Example-6: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-7: Trailing Number Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-8: Trailing Number Symbols
//
//	Trailing Symbols: "+$"
//	Number String:   "123.456+$"
//
// Example-9: Trailing Number Symbols
//
//	Trailing Symbols: "$"
//	Number String:   "123.456$"
//
// Example-10: Trailing Number Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-11: Trailing Number Symbols
//
//	Trailing Symbols: "-$"
//	Number String:   "123.456-$"
//
// Example-12: Trailing Number Symbols
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

// NewNumberSignTrailingSymbol - Creates and returns a new instance
// of NumStrNumberSymbolSpec configured with a trailing
// number symbol character or characters.
//
// Trailing number symbol characters can include such symbols as
// plus signs ('+'), minus signs ('-') and/or currency symbols
// ('$').
//
// Trailing number symbol characters are intended for use in
// formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingNumberSymbol     		string
//
//		A string containing the trailing number symbol
//		character or characters used to configure the returned
//		instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with trailing
//		number symbol characters, will be returned through this
//		parameter.
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Number Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-2: Trailing Number Symbols
//
//	Trailing Symbols: "+$"
//	Number String:   "123.456+$"
//
// Example-3: Trailing Number Symbols
//
//	Trailing Symbols: "$"
//	Number String:   "123.456$"
//
// Example-4: Trailing Number Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-5: Trailing Number Symbols
//
//	Trailing Symbols: "-$"
//	Number String:   "123.456-$"
//
// Example-6: Trailing Number Symbols
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

// NewNumberSignTrailingSymbolRunes - Creates and returns a new
// instance of NumStrNumberSymbolSpec configured with a
// trailing number symbol character or characters.
//
// Trailing number symbol characters can include such symbols as
// plus signs ('+'), minus signs ('-') and/or currency symbols
// ('$').
//
// Trailing number symbol characters are intended for use in
// formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingNumberSymbol			[]rune
//
//		An array of runes containing the trailing number symbol
//		character or characters used to configure the returned
//		instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with trailing
//		number symbol characters, will be returned through this
//		parameter.
//
//	err								error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Number Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-2: Trailing Number Symbols
//
//	Trailing Symbols: "+$"
//	Number String:   "123.456+$"
//
// Example-3: Trailing Number Symbols
//
//	Trailing Symbols: "$"
//	Number String:   "123.456$"
//
// Example-4: Trailing Number Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-5: Trailing Number Symbols
//
//	Trailing Symbols: "-$"
//	Number String:   "123.456-$"
//
// Example-6: Trailing Number Symbols
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

// NewSignedNumDefaultsFrance
//
// Creates and returns three new instances of default
// positive signed number symbols, zero value symbols
// and negative signed number symbols configured in
// accordance with formatting standards commonly applied
// in France.
//
// The positive signed number symbol is empty or blank
// because under French formatting standards, positive
// number signs are implied and not explicitly
// displayed. Therefore, no leading plus ('+') symbol is
// required.
//
// Likewise, the zero signed number symbol is also empty
// or blank because under French formatting standards,
// zero numeric values have no number sign symbols.
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
//			     Number Field Index:  01234567
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
//		blank because under French signed number
//		formatting standards, the positive number sign
//		('+') is implied and not explicitly displayed.
//
//	zeroSignedNumberSymbols			NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under French signed number
//	 	formatting standards, zero numeric values do not
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
//					     Number Field Index:  01234567
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
			"NewSignedNumDefaultsFrance()",
		"")

	if err != nil {
		return positiveSignedNumberSymbols,
			zeroSignedNumberSymbols,
			negativeSignedNumberSymbols,
			err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultFrance(
			&positiveSignedNumberSymbols,
			&zeroSignedNumberSymbols,
			&negativeSignedNumberSymbols,
			ePrefix)

	return positiveSignedNumberSymbols,
		zeroSignedNumberSymbols,
		negativeSignedNumberSymbols,
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
// Likewise, the zero signed number symbol is also empty
// or blank because under German formatting standards,
// zero numeric values have no number sign symbols.
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
//			     Number Field Index:  01234567890
//			Total Number String Length: 11
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
//		blank because under German signed number
//		formatting standards, the positive number sign
//		('+') is implied and not explicitly displayed.
//
//	zeroSignedNumberSymbols			NumStrNumberSymbolSpec
//
//		One of three new returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under German signed number
//	 	formatting standards, zero numeric values do not
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
//					     Number Field Index:  01234567
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
			"NewSignedNumDefaultsGermany()",
		"")

	if err != nil {
		return positiveSignedNumberSymbols,
			zeroSignedNumberSymbols,
			negativeSignedNumberSymbols,
			err

	}

	err = new(numStrNumberSymbolSpecMechanics).
		setSignedNumSymbolsDefaultGermany(
			&positiveSignedNumberSymbols,
			&zeroSignedNumberSymbols,
			&negativeSignedNumberSymbols,
			ePrefix)

	return positiveSignedNumberSymbols,
		zeroSignedNumberSymbols,
		negativeSignedNumberSymbols,
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
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			     Number Field Index:  01234567
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
//					     Number Field Index:  01234567
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
//			     Number Field Index:  012345678
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
//				     Number Field Index:  012345678
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
//		(currencyNumSignRelativePosition) will be deleted
//		and reset to the value specified by input parameter,
//		'currencyNumSignRelativePosition'.
//
//	currencyNumSignRelPos		CurrencyNumSignRelativePosition
//
//		This parameter is used exclusively by Currency
//		Symbol Specifications.
//
//		Type CurrencyNumSignRelativePosition is an
//		enumeration which has three values, only two of
//		which are valid:
//
//			CurrNumSignRelPos.None()			- Invalid
//			CurrNumSignRelPos.OutsideNumSign()	- Valid
//			CurrNumSignRelPos.InsideNumSign()	- Valid
//
//		Currency Symbols have the option of being
//		positioned either inside or outside number sign
//		symbols formatted with numeric values in a number
//		string.
//
//		Examples CurrNumSignRelPos.OutsideNumSign()
//				"$ -123.45"
//				"123.45- €"
//
//		Examples CurrNumSignRelPos.InsideNumSign()
//
//			Examples:
//				"- $123.45"
//				"123.45€ -"
//
//		Be Advised -
//			If the currency symbol is formatted Outside a
//			Number Field and the number sign symbol is
//			formatted Inside a Number Field, this
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

// SetNOP - Resets and configures the current instance
// of NumStrNumberSymbolSpec as a NOP.
//
// NOP is a computer science term which stands for 'No Operation'
// meaning it performs no operations and serves an empty
// placeholder.
//
// This method is often used to configure implied positive number
// signs in number string formatting. The positive number sign,
// plus ('+'), is often implied for positive numeric values and
// is not displayed in the number string.
//
// Call this method when creating implied positive number signs
// for number strings.
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

// SetNumberSignLeadingSymbol - Resets and configures a
// leading number symbol character or characters for the
// current instance of NumStrNumberSymbolSpec.
//
// Leading number symbol characters can include such
// symbols as plus signs ('+'), minus signs ('-') and
// currency symbols ('$').
//
// Leading number symbol characters are intended for use
// in formatting numeric values displayed in number
// strings.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the leading
// number symbol data value in the current instance of
// NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol				string
//
//		A string containing the leading number symbol
//		character or characters used to configure the current
//		instance of NumStrNumberSymbolSpec.
//
//		Leading number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Symbols
//
//	Leading Symbols: "$+"
//	Number String:   "$+123.456"
//
// Example-3: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-4: Leading Number Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-5: Leading Number Symbols
//
//	Leading Symbols: "$-"
//	Number String:   "$-123.456"
//
// Example-6: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignLeadingSymbol(
	leadingNumberSymbol string,
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
			[]rune(leadingNumberSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	return err
}

// SetNumberSignLeadingSymbolRunes - Resets and configures a leading
// number symbol character or characters for the current
// instance of NumStrNumberSymbolSpec.
//
// Leading number symbol characters can include such symbols as
// plus signs ('+'), minus signs ('-') and/or currency symbols
// ('$').
//
// Leading number symbol characters are intended for use in
// formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the leading number
// symbol data value in the current instance of
// NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol				[]rune
//
//		An array of runes containing the leading number
//	 	symbol character or characters used to configure
//	  	the current instance of NumStrNumberSymbolSpec.
//
//		Leading number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Symbols
//
//	Leading Symbols: "$+"
//	Number String:   "$+123.456"
//
// Example-3: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-4: Leading Number Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-5: Leading Number Symbols
//
//	Leading Symbols: "$-"
//	Number String:   "$-123.456"
//
// Example-6: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignLeadingSymbolRunes(
	leadingNumberSymbol []rune,
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
			leadingNumberSymbol,
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	return err
}

// SetNumberSignLeadingTrailingSymbol - Resets and configures
// both Leading and Trailing number symbol characters for
// the current instance of NumStrNumberSymbolSpec.
//
// Leading and Trailing number symbol characters can include
// such symbols as plus signs ('+'), minus signs ('-') and/or
// currency symbols ('$').
//
// Leading and Trailing number symbols are intended for use
// in formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the Leading and
// Trailing number symbol data value in the current instance
// of NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol     		string
//
//		A string containing the leading number symbol
//		character or characters used to configure the
//		returned instance of NumStrNumberSymbolSpec.
//
//		Leading number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//		and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	trailingNumberSymbol     		string
//
//		A string containing the trailing number symbol
//		character or characters used to configure the returned
//		instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Symbols
//
//	Leading Symbols: "$+"
//	Number String:   "$+123.456"
//
// Example-3: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-4: Leading Number Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-5: Leading Number Symbols
//
//	Leading Symbols: "$-"
//	Number String:   "$-123.456"
//
// Example-6: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-7: Trailing Number Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-8: Trailing Number Symbols
//
//	Trailing Symbols: "+$"
//	Number String:   "123.456+$"
//
// Example-9: Trailing Number Symbols
//
//	Trailing Symbols: "$"
//	Number String:   "123.456$"
//
// Example-10: Trailing Number Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-11: Trailing Number Symbols
//
//	Trailing Symbols: "-$"
//	Number String:   "123.456-$"
//
// Example-12: Trailing Number Symbols
//
//	 Leading Symbols: "("
//		Trailing Symbols: ")"
//		Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignLeadingTrailingSymbol(
	leadingNumberSymbol string,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	trailingNumberSymbol string,
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
			[]rune(leadingNumberSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	if err != nil {
		return err
	}

	err = nStrNumSymSpecNanobot.
		setTrailingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			[]rune(trailingNumberSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return err
}

// SetNumberSignLeadingTrailingSymbolRunes - Resets and configures
// both Leading and Trailing number symbol characters for the
// current instance of NumStrNumberSymbolSpec.
//
// Leading and Trailing number symbol characters can include
// such symbols as plus signs ('+'), minus signs ('-') and/or
// currency symbols ('$').
//
// Leading and Trailing number symbols are intended for use
// in formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the Leading and
// Trailing number symbol data value in the current instance
// of NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNumberSymbol				[]rune
//
//		An array of runes containing the leading number
//		symbol character or characters used to configure
//		the returned instance of NumStrNumberSymbolSpec.
//
//		Leading number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//		and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	trailingNumberSymbol			[]rune
//
//		An array of runes containing the trailing number symbol
//		character or characters used to configure the returned
//		instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Leading Number Symbols
//
//	Leading Symbols: "+ "
//	Number String:   "+ 123.456"
//
// Example-2: Leading Number Symbols
//
//	Leading Symbols: "$+"
//	Number String:   "$+123.456"
//
// Example-3: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-4: Leading Number Symbols
//
//	Leading Symbols: "- "
//	Number String:   "- 123.456"
//
// Example-5: Leading Number Symbols
//
//	Leading Symbols: "$-"
//	Number String:   "$-123.456"
//
// Example-6: Leading Number Symbols
//
//	Leading Symbols: "$"
//	Number String:   "$123.456"
//
// Example-7: Trailing Number Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-8: Trailing Number Symbols
//
//	Trailing Symbols: "+$"
//	Number String:   "123.456+$"
//
// Example-9: Trailing Number Symbols
//
//	Trailing Symbols: "$"
//	Number String:   "123.456$"
//
// Example-10: Trailing Number Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-11: Trailing Number Symbols
//
//	Trailing Symbols: "-$"
//	Number String:   "123.456-$"
//
// Example-12: Trailing Number Symbols
//
//	 Leading Symbols: "("
//		Trailing Symbols: ")"
//		Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignLeadingTrailingSymbolRunes(
	leadingNumberSymbol []rune,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	trailingNumberSymbol []rune,
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
			leadingNumberSymbol,
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	if err != nil {
		return err
	}

	err = nStrNumSymSpecNanobot.
		setTrailingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			trailingNumberSymbol,
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return err
}

// SetNumberSignTrailingSymbol - Resets and configures a trailing
// number symbol character or characters for the current
// instance of NumStrNumberSymbolSpec
//
// Trailing number symbol characters can include such symbols as
// plus signs ('+'), minus signs ('-') and/or currency symbols
// ('$').
//
// Trailing number symbol characters are intended for use in
// formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the trailing number
// symbol data values in the current instance of
// NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	trailingNumberSymbol			string
//
//		A string containing the trailing number symbol
//		character or characters used to configure the current
//		instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//	err								error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Number Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-2: Trailing Number Symbols
//
//	Trailing Symbols: "+$"
//	Number String:   "123.456+$"
//
// Example-3: Trailing Number Symbols
//
//	Trailing Symbols: "$"
//	Number String:   "123.456$"
//
// Example-4: Trailing Number Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-5: Trailing Number Symbols
//
//	Trailing Symbols: "-$"
//	Number String:   "123.456-$"
//
// Example-6: Trailing Number Symbols
//
//	 Leading Symbols: "("
//		Trailing Symbols: ")"
//		Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignTrailingSymbol(
	trailingNumberSymbol string,
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
			[]rune(trailingNumberSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return err
}

// SetNumberSignTrailingSymbolRunes - Resets and configures a
// trailing number symbol character or characters for the
// current instance of NumStrNumberSymbolSpec.
//
// Trailing number symbol characters can include such symbols
// as plus signs ('+'), minus signs ('-') and/or currency
// symbols ('$').
//
// Trailing number symbol characters are intended for use in
// formatting numeric values displayed in number strings.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the trailing number
// symbol data value in the current instance of
// NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	trailingNumberSymbol			[]rune
//
//		An array of runes containing the trailing number symbol
//		character or characters used to configure the current
//		instance of NumStrNumberSymbolSpec.
//
//		Trailing number symbol characters can include such
//		symbols as plus signs ('+'), minus signs ('-')
//	 	and/or currency symbols ('$').
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
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
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
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
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
//	newNStrNumberSymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with a trailing
//		number symbol, will be returned through this
//		parameter.
//
//	err								error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
//
// ----------------------------------------------------------------
//
// # Usage
//
// Example-1: Trailing Number Symbols
//
//	Trailing Symbols: " +"
//	Number String:   "123.456 +"
//
// Example-2: Trailing Number Symbols
//
//	Trailing Symbols: "+$"
//	Number String:   "123.456+$"
//
// Example-3: Trailing Number Symbols
//
//	Trailing Symbols: "$"
//	Number String:   "123.456$"
//
// Example-4: Trailing Number Symbols
//
//	Trailing Symbols: " -"
//	Number String:   "123.456 -"
//
// Example-5: Trailing Number Symbols
//
//	Trailing Symbols: "-$"
//	Number String:   "123.456-$"
//
// Example-6: Trailing Number Symbols
//
//	Leading Symbols: "("
//	Trailing Symbols: ")"
//	Number String:   "(123.456)"
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetNumberSignTrailingSymbolRunes(
	trailingNumberSymbol []rune,
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
			trailingNumberSymbol,
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return err
}

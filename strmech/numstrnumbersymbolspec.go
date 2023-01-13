package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

//	NumStrNumberSymbolSpec
//
//	Contains one or more characters to be inserted and
//	displayed in a number string.
//
// # Background
//
//	The Number String Number Symbol Specification
//	includes specifications for both leading and
//	trailing characters.
//
//	As such, 'Number Symbols' can include plus signs
//	('+'), minus signs ('-'), parentheses ('()') as
//	well as currency ('$') and other symbols.
//
//	The user has the option to configure both leading and
//	trailing number symbols, leading number symbols only
//	or trailing number symbols only.
//
//	A Number Symbol is defined as one or more text
//	characters.
//
// # Usage
//
//	Example-1:
//		Leading Number Symbols for Positive Values
//
//		Leading Symbols: "+ "
//		Number String:   "+ 123.456"
//
//	Example-2: Leading Number Symbols
//		Leading Number Symbols for Positive Values
//
//		Leading Symbols: "+"
//		Number String:   "+123.456"
//
//	Example-3: Leading Number Symbols
//		Leading Number Symbols for Negative Values
//
//		Leading Symbols: "- "
//		Number String:   "- 123.456"
//
//	Example-4: Leading Number Symbols
//		Leading Number Symbols for Negative Values
//
//		Leading Symbols: "-"
//		Number String:   "-123.456"
//
//	Example-5: Trailing Number Symbols
//		Trailing Number Symbols for Positive Values
//
//		Trailing Symbols: " +"
//		Number String:   "123.456 +"
//
//	Example-6: Trailing Number Symbols
//		Trailing Number Symbols for Positive Values
//
//		Trailing Symbols: "+"
//		Number String:   "123.456+"
//
//	Example-7: Trailing Number Symbols
//		Trailing Number Symbols for Negative Values
//
//		Trailing Symbols: " -"
//		Number String:   "123.456 -"
//
//	Example-8: Trailing Number Symbols
//		Trailing Number Symbols for Negative Values
//
//		Trailing Symbols: "-"
//		Number String:   "123.456-"
//
//	Example-9: Trailing Number Symbols
//		Trailing Number Symbols for Negative Values
//
//		Leading Symbols: "("
//		Trailing Symbols: ")"
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
	// exclusively by Currency Symbol Specifications.
	// This enumeration has three values, only two of
	// which are valid:
	//
	//	CurrNumSignRelPos.None()			- Invalid
	//	CurrNumSignRelPos.OutsideNumSign()	- Valid
	//	CurrNumSignRelPos.InsideNumSign()	- Valid
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

	return new(numStrNumberSymbolSpecNanobot).copyNStrNumberSymbolSpec(
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

	err = new(numStrNumberSymbolSpecNanobot).copyNStrNumberSymbolSpec(
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

	new(numStrNumberSymbolSpecNanobot).empty(
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
//	CurrNumSignRelPos.None()			- Invalid
//	CurrNumSignRelPos.OutsideNumSign()	- Valid
//	CurrNumSignRelPos.InsideNumSign()	- Valid
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
//				Number Field Index:       01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//	         Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:       0123456789
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
//				Number Field Index:  012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:  0123456789
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
//				Number Field Index:       01234567
//				Total Number String Length: 8
//
//			Example-2:
//				Number Field Length: 10
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//	         Number Text Justification: Centered
//				Formatted Number String: " (123.45) "
//				Number Field Index:       0123456789
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
//				Number Field Index:       012345678
//				Total Number String Length: 9
//
//			Example-4:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//	         Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:       0123456789
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

// NewLeadingNumberSymbol - Creates and returns a new instance
// of NumStrNumberSymbolSpec configured with a leading
// number symbol character or characters.
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewLeadingNumberSymbol(
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
			"NewLeadingNumberSymbol()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setLeadingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			[]rune(leadingNumberSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewLeadingNumberSymbolRunes - Creates and returns a new
// instance of NumStrNumberSymbolSpec configured with
// a leading number symbol character or characters.
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
// # Input Parameters
//
//	leadingNumberSymbol		[]rune
//
//		An array of runes containing the leading number
//	 	symbol character or characters used to configure
//	  	the returned instance of NumStrNumberSymbolSpec.
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
//	newNStrNumberSymbolSpec		NumStrNumberSymbolSpec
//
//		If this method completes successfully, a new instance of
//		NumStrNumberSymbolSpec, configured with leading
//		number symbol characters, will be returned through
//		this parameter.
//
//	err							error
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewLeadingNumberSymbolRunes(
	leadingNumberSymbol []rune,
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
			"NewLeadingNumberSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setLeadingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			leadingNumberSymbol,
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewLeadingTrailingNumSymbol - Creates and returns a new
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewLeadingTrailingNumSymbol(
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
			"NewLeadingTrailingNumSymbol()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

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

// NewLeadingTrailingNumSymbolRunes - Creates and returns a
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewLeadingTrailingNumSymbolRunes(
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
			"NewLeadingTrailingNumSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

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

	new(numStrNumberSymbolSpecNanobot).empty(
		&newNOPNumSymSpec)

	return newNOPNumSymSpec
}

// NewTrailingNumberSymbol - Creates and returns a new instance
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewTrailingNumberSymbol(
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
			"NewTrailingNumberSymbol()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setTrailingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			[]rune(trailingNumberSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
}

// NewTrailingNumberSymbolRunes - Creates and returns a new
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) NewTrailingNumberSymbolRunes(
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
			"NewTrailingNumberSymbolRunes()",
		"")

	if err != nil {
		return newNStrNumberSymbolSpec, err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setTrailingNStrNumSymbolSpec(
			&newNStrNumberSymbolSpec,
			trailingNumberSymbol,
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"newNStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return newNStrNumberSymbolSpec, err
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
//			Number Field Index:  01234567
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
//		blank because under United States' signed number
//		formatting standards, the positive number sign ('+')
//		is implied and not displayed.
//
//	zeroSignedNumberSymbols			NumStrNumberSymbolSpec
//
//		One of two three returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under United States' signed number
//		formatting standards, zero numeric values do not
//		have an associated number sign.
//
//	negativeSignedNumberSymbols		NumStrNumberSymbolSpec
//
//		One of two new returned instances of
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
//					Number Field Index:  01234567
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
//			Number Field Index:       012345678
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
//		blank because under United States' signed number
//		formatting standards, the positive number sign ('+')
//		is implied and not displayed.
//
//	zeroSignedNumberSymbols			NumStrNumberSymbolSpec
//
//		One of three returned instances of
//		NumStrNumberSymbolSpec. This instance is empty or
//		blank because under United States' signed number
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
//				Number Field Index:       012345678
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

	return new(numStrNumberSymbolSpecNanobot).
		setCurrencyNumSignRelPos(
			nStrNumberSymbolSpec,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec"))
}

// SetLeadingNumberSymbol - Resets and configures a
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetLeadingNumberSymbol(
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
			"SetLeadingNumberSymbol()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setLeadingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			[]rune(leadingNumberSymbol),
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	return err
}

// SetLeadingNumberSymbolRunes - Resets and configures a leading
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetLeadingNumberSymbolRunes(
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
			"SetLeadingNumberSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setLeadingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			leadingNumberSymbol,
			leadingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"leadingNumberSymbol"))

	return err
}

// SetLeadingTrailingNumberSymbol - Resets and configures
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetLeadingTrailingNumberSymbol(
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
			"SetLeadingTrailingNumberSymbol()",
		"")

	if err != nil {
		return err
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

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

// SetLeadingTrailingNumberSymbolRunes - Resets and configures
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetLeadingTrailingNumberSymbolRunes(
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
			"SetLeadingTrailingNumberSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	nStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

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

	new(numStrNumberSymbolSpecNanobot).empty(
		nStrNumberSymbolSpec)

	return
}

// SetTrailingNumberSymbol - Resets and configures a trailing
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetTrailingNumberSymbol(
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
			"SetTrailingNumberSymbol()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setTrailingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			[]rune(trailingNumberSymbol),
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return err
}

// SetTrailingNumberSymbolRunes - Resets and configures a
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
func (nStrNumberSymbolSpec *NumStrNumberSymbolSpec) SetTrailingNumberSymbolRunes(
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
			"SetTrailingNumberSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		setTrailingNStrNumSymbolSpec(
			nStrNumberSymbolSpec,
			trailingNumberSymbol,
			trailingNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumberSymbolSpec<-"+
					"trailingNumberSymbol"))

	return err
}

// numStrNumberSymbolSpecMechanics - This type provides
// helper methods for NumStrNumberSymbolSpec
type numStrNumberSymbolSpecMechanics struct {
	lock *sync.Mutex
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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

	numStrNumSymSpecNanobot.empty(
		currencySymbols)

	err = numStrNumSymSpecNanobot.setTrailingNStrNumSymbolSpec(
		currencySymbols,
		[]rune{' ', '€'},
		NumFieldSymPos.InsideNumField(),
		ePrefix.XCpy(
			"currencySymbols"))

	if err != nil {
		return err
	}

	return numStrNumSymSpecNanobot.setCurrencyNumSignRelPos(
		currencySymbols,
		CurrNumSignRelPos.OutsideNumSign(),
		ePrefix.XCpy(
			"currencySymbols"))
}

// setCurrencyDefaultsUK
//
// Receives an instance of NumStrNumberSymbolSpec and
// configures it with the default UK (United Kingdom)
// currency symbol. The default UK currency symbol is
// a leading pound sign.
//
//	Example:
//		£ 123.45
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
//		pound sign.
//
//			Example:
//				£ 123.45
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setCurrencyDefaultsUK(
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
			"setCurrencyDefaultsUK()",
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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

	numStrNumSymSpecNanobot.empty(
		currencySymbols)

	err = numStrNumSymSpecNanobot.setLeadingNStrNumSymbolSpec(
		currencySymbols,
		[]rune{'£', ' '},
		NumFieldSymPos.InsideNumField(),
		ePrefix.XCpy(
			"currencySymbols"))

	if err != nil {
		return err
	}

	return numStrNumSymSpecNanobot.setCurrencyNumSignRelPos(
		currencySymbols,
		CurrNumSignRelPos.InsideNumSign(),
		ePrefix.XCpy(
			"currencySymbols"))
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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

	numStrNumSymSpecNanobot.empty(
		currencySymbols)

	err = numStrNumSymSpecNanobot.setLeadingNStrNumSymbolSpec(
		currencySymbols,
		[]rune{'$', ' '},
		NumFieldSymPos.InsideNumField(),
		ePrefix.XCpy(
			"currencySymbols"))

	if err != nil {
		return err
	}

	return numStrNumSymSpecNanobot.setCurrencyNumSignRelPos(
		currencySymbols,
		CurrNumSignRelPos.OutsideNumSign(),
		ePrefix.XCpy(
			"currencySymbols"))
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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecNanobot{}

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

// numStrNumberSymbolSpecNanobot - This type provides
// helper methods for NumStrNumberSymbolSpec
type numStrNumberSymbolSpecNanobot struct {
	lock *sync.Mutex
}

//	copyNStrNumberSymbolSpec
//
//	Copies all data from input parameter
//	'sourceNumSymbolSpec' to input parameter
//
// 'destinationNumSymbolSpec'. Both instances are of
//
//	type NumStrNumberSymbolSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields in
//	'destinationNumSymbolSpec' will be deleted and overwritten.
//
//	Also, NO data validation is performed on 'sourceNumSymbolSpec'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationNumSymbolSpec	*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		All the member variable data fields in this object will be
//		replaced by data values copied from input parameter
//		'sourceNumSymbolSpec'.
//
//		'destinationNumSymbolSpec' is the destination for this
//		copy operation.
//
//
//	sourceNumSymbolSpec			*NumStrNumberSymbolSpec
//
//		A pointer to another NumStrNumberSymbolSpec
//		instance. All the member variable data values from this
//		object will be copied to corresponding member variables in
//		'destinationNumSymbolSpec'.
//
//		'sourceNumSymbolSpec' is the source for this copy
//		operation.
//
//		No data validation is performed on 'sourceNumSymbolSpec'.
//
//	errPrefDto		*ePref.ErrPrefixDto
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
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrNumSymSpecNanobot *numStrNumberSymbolSpecNanobot) copyNStrNumberSymbolSpec(
	destinationNumSymbolSpec *NumStrNumberSymbolSpec,
	sourceNumSymbolSpec *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumSymSpecNanobot.lock == nil {
		nStrNumSymSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymSpecNanobot.lock.Lock()

	defer nStrNumSymSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecNanobot."+
			"copyNStrNumberSymbolSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationNumSymbolSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNumSymbolSpec' is invalid!\n"+
			"'destinationNumSymbolSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceNumSymbolSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNumSymbolSpec' is invalid!\n"+
			"'sourceNumSymbolSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	nStrNumSymSpecAtom := nStrNumberSymbolSpecAtom{}

	nStrNumSymSpecAtom.emptyLeadingNStrNumSymbol(
		destinationNumSymbolSpec)

	nStrNumSymSpecAtom.emptyTrailingNStrNumSymbol(
		destinationNumSymbolSpec)

	err = destinationNumSymbolSpec.leadingNumberSymbols.
		CopyIn(
			&sourceNumSymbolSpec.leadingNumberSymbols,
			ePrefix.XCpy(
				"destinationNumSymbolSpec.leadingNumberSymbols<-"+
					"sourceNumSymbolSpec"))

	if err != nil {
		return err
	}

	err = destinationNumSymbolSpec.trailingNumberSymbols.
		CopyIn(
			&sourceNumSymbolSpec.trailingNumberSymbols,
			ePrefix.XCpy(
				"destinationNumSymbolSpec.trailingNumberSymbols<-"+
					"sourceNumSymbolSpec"))

	return err
}

// empty - Receives a pointer to an instance of
// NumStrNumberSymbolSpec and proceeds to reset the
// data values for all member variables to their initial or
// zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'nStrNumSymbolSpec' will be deleted and reset to their zero values.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	nStrNumSymbolSpec           *NumStrNumberSymbolSpec
//	   - A pointer to an instance of NumStrNumberSymbolSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (nStrNumSymSpecNanobot *numStrNumberSymbolSpecNanobot) empty(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec) {

	if nStrNumSymSpecNanobot.lock == nil {
		nStrNumSymSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymSpecNanobot.lock.Lock()

	defer nStrNumSymSpecNanobot.lock.Unlock()

	if nStrNumSymbolSpec == nil {
		return
	}

	nStrNumSymSpecAtom := nStrNumberSymbolSpecAtom{}

	nStrNumSymSpecAtom.emptyLeadingNStrNumSymbol(
		nStrNumSymbolSpec)

	nStrNumSymSpecAtom.emptyTrailingNStrNumSymbol(
		nStrNumSymbolSpec)

	nStrNumSymSpecAtom.emptyCurrNumSignRelPos(
		nStrNumSymbolSpec)

	return
}

// setCurrencyNumSignRelPos
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
//		The Currency Number Sign Relative Position value
//		contained in this instance
//		(currencyNumSignRelativePosition) will be deleted
//		and reset to the value specified by input parameter,
//		'currencyNumSignRelPos'.
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
//			parameter will be ignored.
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
func (nStrNumSymSpecNanobot *numStrNumberSymbolSpecNanobot) setCurrencyNumSignRelPos(
	numSymbolSpec *NumStrNumberSymbolSpec,
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
			"setCurrencyNumSignRelPos()",
		"")

	if err != nil {
		return err
	}

	if numSymbolSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSymbolSpec' is invalid!\n"+
			"'numSymbolSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if currencyNumSignRelPos.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencyNumSignRelPos' is invalid!\n"+
			" currencyNumSignRelativePosition String Value = %v\n"+
			"currencyNumSignRelativePosition Integer Value = %v\n",
			ePrefix.String(),
			currencyNumSignRelPos.String(),
			currencyNumSignRelPos.XValueInt())

		return err
	}

	numSymbolSpec.currencyNumSignRelativePosition =
		currencyNumSignRelPos

	return err
}

// setLeadingNStrNumSymbolSpec - Deletes and resets the data
// value of the Leading Number Symbol contained in an
// instance of NumStrNumberSymbolSpec passed as an input
// parameter.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numSymbolSpec				*NumStrNumberSymbolSpec
//		A pointer to a NumStrNumberSymbolSpec instance.
//		The Leading Number Symbol contained in this
//		instance will be deleted and reset to the value
//		specified by input parameter,
//		'leadingNumberSymbols'.
//
//	leadingNumberSymbols		[]rune
//		An array of runes specifying the character or
//		characters which will be copied to the Leading
//		Number Symbol contained in input parameter,
//		'posNumSignSpec'.
//
//	leadingNumFieldSymPosition	NumberFieldSymbolPosition
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
//	errPrefDto					*ePref.ErrPrefixDto
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
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrNumSymSpecNanobot *numStrNumberSymbolSpecNanobot) setLeadingNStrNumSymbolSpec(
	numSymbolSpec *NumStrNumberSymbolSpec,
	leadingNumberSymbol []rune,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumSymSpecNanobot.lock == nil {
		nStrNumSymSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymSpecNanobot.lock.Lock()

	defer nStrNumSymSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecNanobot."+
			"setLeadingNStrNumSymbolSpec()",
		"")

	if err != nil {
		return err
	}

	if numSymbolSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSymbolSpec' is invalid!\n"+
			"'numSymbolSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if !leadingNumFieldSymPosition.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNumFieldSymPosition' is invalid!\n"+
			"'leadingNumFieldSymPosition' string value  = '%v'\n"+
			"'leadingNumFieldSymPosition' integer value = '%v'\n",
			ePrefix.String(),
			leadingNumFieldSymPosition.String(),
			leadingNumFieldSymPosition.XValueInt())

		return err

	}

	new(nStrNumberSymbolSpecAtom).emptyLeadingNStrNumSymbol(
		numSymbolSpec)

	if len(leadingNumberSymbol) > 0 {

		err = numSymbolSpec.leadingNumberSymbols.SetRuneArray(
			leadingNumberSymbol,
			ePrefix.XCpy(
				"numSymbolSpec.leadingNumberSymbols"+
					"<-leadingNumberSymbols"))

		if err != nil {
			return err
		}
	}

	numSymbolSpec.leadingNumberFieldSymbolPosition =
		leadingNumFieldSymPosition

	return err
}

// setTrailingNStrNumSymbolSpec - Deletes and resets the data
// value of the Trailing Number Symbol contained in an
// instance of NumStrNumberSymbolSpec passed as an input
// parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numSignSymbolSpec				*NumStrNumberSymbolSpec
//		A pointer to a NumStrNumberSymbolSpec instance.
//		The Trailing Number Symbol contained in this
//		instance will be deleted and reset to the value
//		specified by input parameter,
//		'trailingNumberSymbols'.
//
//	trailingNumberSymbols			[]rune
//		   - An array of runes specifying the character or
//		     characters which will be copied to the Trailing
//		     Number Symbol contained in input parameter,
//	      'posNumSignSpec'.
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
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
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrNumSymSpecNanobot *numStrNumberSymbolSpecNanobot) setTrailingNStrNumSymbolSpec(
	numSignSymbolSpec *NumStrNumberSymbolSpec,
	trailingNumberSymbol []rune,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumSymSpecNanobot.lock == nil {
		nStrNumSymSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymSpecNanobot.lock.Lock()

	defer nStrNumSymSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecNanobot."+
			"setTrailingNStrNumSymbolSpec()",
		"")

	if err != nil {
		return err
	}

	if numSignSymbolSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numSignSymbolSpec' is invalid!\n"+
			"'numSignSymbolSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if !trailingNumFieldSymPosition.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNumFieldSymPosition' is invalid!\n"+
			"'trailingNumFieldSymPosition' string value  = '%v'\n"+
			"'trailingNumFieldSymPosition' integer value = '%v'\n",
			ePrefix.String(),
			trailingNumFieldSymPosition.String(),
			trailingNumFieldSymPosition.XValueInt())

		return err

	}

	new(nStrNumberSymbolSpecAtom).emptyTrailingNStrNumSymbol(
		numSignSymbolSpec)

	if len(trailingNumberSymbol) > 0 {

		err = numSignSymbolSpec.trailingNumberSymbols.SetRuneArray(
			trailingNumberSymbol,
			ePrefix.XCpy(
				"numSignSymbolSpec.trailingNumberSymbols"+
					"<-trailingNumberSymbols"))

		if err != nil {
			return err
		}
	}

	numSignSymbolSpec.trailingNumberFieldSymbolPosition =
		trailingNumFieldSymPosition

	return err
}

// nStrNumberSymbolSpecAtom - This type provides
// helper methods for NumStrNumberSymbolSpec
type nStrNumberSymbolSpecAtom struct {
	lock *sync.Mutex
}

// emptyLeadingNStrNumSymbol - Receives a pointer to an
// instance of NumStrNumberSymbolSpec and proceeds to
// reset the member variable data for the leading number
// symbol to an initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the leading number
// symbol contained in input parameter 'nStrNumSymbolSpec'
// will be deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNumSymbolSpec           *NumStrNumberSymbolSpec
//
//		A pointer to an instance of NumStrNumberSymbolSpec.
//		The Leading Number Symbol contained in this
//		instance will be deleted and reset to an empty or zero
//		value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) emptyLeadingNStrNumSymbol(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec) {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	if nStrNumSymbolSpec == nil {
		return
	}

	nStrNumSymbolSpec.leadingNumberSymbols.Empty()

	nStrNumSymbolSpec.leadingNumberFieldSymbolPosition =
		NumFieldSymPos.InsideNumField()

	return
}

// emptyTrailingNStrNumSymbol - Receives a pointer to an instance
// of NumStrNumberSymbolSpec and proceeds to reset the
// member variable data for the trailing number symbol
// to an initial or zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data for the trailing number symbol
// contained in input parameter 'nStrNumSymbolSpec' will be
// deleted and reset to an empty or zero value.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNumSymbolSpec           *NumStrNumberSymbolSpec
//	   - A pointer to an instance of NumStrNumberSymbolSpec.
//	     The Trailing Number Symbol contained in this
//	     instance will be deleted and reset to an empty or zero
//	     value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) emptyTrailingNStrNumSymbol(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec) {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	if nStrNumSymbolSpec == nil {
		return
	}

	nStrNumSymbolSpec.trailingNumberSymbols.Empty()

	nStrNumSymbolSpec.trailingNumberFieldSymbolPosition =
		NumFieldSymPos.InsideNumField()

	return
}
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) emptyCurrNumSignRelPos(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec) {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	if nStrNumSymbolSpec == nil {
		return
	}

	nStrNumSymbolSpec.currencyNumSignRelativePosition =
		CurrNumSignRelPos.None()
}

// equal - Receives a pointer to two instances of
// NumStrNumberSymbolSpec and proceeds to compare their
// member variables in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are
// equal in all respects, this flag is set to 'true'. Otherwise,
// this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNumSymbolSpec1    *NumStrNumberSymbolSpec
//	   - An instance of NumStrNumberSymbolSpec.
//	     Internal member variables from 'nStrNumSymbolSpec1'
//	     will be compared to those of 'nStrNumSymbolSpec2' to
//	     determine if both instances are equivalent.
//
//
//	nStrNumSymbolSpec2    *NumStrNumberSymbolSpec
//	   - An instance of NumStrNumberSymbolSpec.
//	     Internal member variables from 'nStrNumSymbolSpec2'
//	     will be compared to those of 'nStrNumSymbolSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'nStrNumSymbolSpec1' and
//	     'nStrNumSymbolSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (nStrNumSymSpecAtom *nStrNumberSymbolSpecAtom) equal(
	nStrNumSymbolSpec1 *NumStrNumberSymbolSpec,
	nStrNumSymbolSpec2 *NumStrNumberSymbolSpec) bool {

	if nStrNumSymSpecAtom.lock == nil {
		nStrNumSymSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumSymSpecAtom.lock.Lock()

	defer nStrNumSymSpecAtom.lock.Unlock()

	if nStrNumSymbolSpec1 == nil ||
		nStrNumSymbolSpec2 == nil {
		return false
	}

	if !nStrNumSymbolSpec1.leadingNumberSymbols.Equal(
		&nStrNumSymbolSpec2.leadingNumberSymbols) {

		return false
	}

	if nStrNumSymbolSpec1.leadingNumberFieldSymbolPosition !=
		nStrNumSymbolSpec2.leadingNumberFieldSymbolPosition {

		return false

	}

	if !nStrNumSymbolSpec1.trailingNumberSymbols.Equal(
		&nStrNumSymbolSpec2.trailingNumberSymbols) {

		return false
	}

	if nStrNumSymbolSpec1.trailingNumberFieldSymbolPosition !=
		nStrNumSymbolSpec2.trailingNumberFieldSymbolPosition {

		return false
	}

	if nStrNumSymbolSpec1.currencyNumSignRelativePosition !=
		nStrNumSymbolSpec2.currencyNumSignRelativePosition {

		return false
	}

	return true
}

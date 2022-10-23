package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

//	NumStrNumberSymbols
//
//	Contains three instances of NumStrNumberSymbolSpec
//	defining the Number Symbols to be used with positive
//	numeric values, negative numeric values and zero
//	numeric values.
//
// # Background
//
//	Each Number String Number Symbol Specification
//	contained within this type includes specifications
//	for both leading and trailing characters.
//
//	As such, 'Number Symbols' can include plus signs
//	('+'), minus signs ('-'), parentheses ('()') as well
//	as currency ('$') and other symbols.
//
//	The user has the option to configure both leading and
//	trailing number symbols, leading number symbols only
//	or trailing number symbols only.
//
//	A Number Symbol is defined as one or more text
//	characters.
type NumStrNumberSymbols struct {
	negativeNumberSign NumStrNumberSymbolSpec
	//	The Number String Negative Number Sign
	//	Specification is used to configure negative
	//	number sign symbols for negative numeric
	//	values formatted and displayed in number
	//	stings.
	//
	//	For currency presentations, the currency
	//	symbol is combined with the negative number
	//	sign.
	//
	//	Example-1: Leading Number Symbols
	//		Leading Number Symbols for Negative Values
	//
	//		Leading Symbols: "- "
	//		Number String:   "- 123.456"
	//
	//	Example-2: Leading Number Symbols With Currency
	//		Leading Number Symbols for Negative Values
	//
	//		Leading Symbols: "$-"
	//		Number String:   "$-123.456"
	//
	//
	//	Example-3: Trailing Number Symbols
	//		Trailing Number Symbols for Negative Values
	//
	//		Trailing Symbols: " -"
	//		Number String:   "123.456 -"
	//
	//	Example-4: Trailing Number Symbols
	//		Trailing Number Symbols for Negative Values
	//
	//		Trailing Symbols: "-$"
	//		Number String:   "123.456-$"

	positiveNumberSign NumStrNumberSymbolSpec
	//	Positive number signs are commonly implied
	//	and not specified. However, the user has
	//	the option to specify a positive number sign
	//	character or characters for positive numeric
	//	values using a Number String Positive Number
	//	Sign Specification.
	//
	//	For currency presentations, the currency
	//	symbol is combined with the positive number
	//	sign.
	//
	//	Example-1: Leading Number Symbols
	//		Leading Number Symbols for Positive Values
	//
	//		Leading Symbols: "+ "
	//		Number String:   "+ 123.456"
	//
	//	Example-2: Leading Number Symbols
	//		Leading Number Symbols for Positive Values
	//
	//		Leading Symbols: "$+"
	//		Number String:   "$+123.456"
	//
	//	Example-3: Leading Number Symbols
	//		Leading Number Symbols for Positive Values
	//
	//		Leading Symbols: "$"
	//		Number String:   "$123.456"
	//
	//	Example-4: Trailing Number Symbols
	//		Trailing Number Symbols for Positive Values
	//
	//		Trailing Symbols: " +"
	//		Number String:   "123.456 +"
	//
	//	Example-5: Trailing Number Symbols
	//		Trailing Number Symbols for Positive Values
	//
	//		Trailing Symbols: "+$"
	//		Number String:   "123.456+$"
	//
	//	Example-6: Trailing Number Symbols
	//		Trailing Number Symbols for Positive Values
	//
	//		Trailing Symbols: "$"
	//		Number String:   "123.456$"

	zeroNumberSign NumStrNumberSymbolSpec
	//	The Number String Zero Number Symbol
	//	Specification is used to configure number
	//	symbols for zero numeric values formatted
	//	and displayed in number stings. Zero number
	//	signs are commonly omitted because zero
	//	does not technically qualify as either a
	//	positive or negative value. However,
	//	currency symbols may be required for zero
	//	values.
	//
	//	For currency presentations, the currency
	//	symbol is often used as either a leading
	//	or trailing symbol for zero numeric
	//	values.
	//
	//	Example-1: Leading Number Symbols
	//		Leading Number Symbols for Zero Values
	//
	//		Leading Symbols: "$"
	//		Trailing Symbols: ""
	//		Number String:   "$0.00"
	//
	//	Example-2: Trailing Number Symbols
	//		Trailing Number Symbols for Zero Values
	//
	//		Leading Symbols: ""
	//		Trailing Symbols: " $"
	//		Number String:   "0.00 $"

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumStrNumberSymbols ('incomingNumSymbols')
//	to the data fields of the current NumStrNumberSymbols
//	instance ('nStrNumSym').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the data fields in current NumStrNumberSymbols
//	instance ('nStrNumSym') will be deleted and
//	overwritten.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	incomingNumSymbols				*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
//		This method will NOT change the data values of
//		internal member variables contained in this
//		instance.
//
//		All data values in this NumStrNumberSymbols
//		instance will be copied to current
//		NumStrNumberSymbols instance ('nStrNumSym').
//
//	 errorPrefix                	interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) CopyIn(
	incomingNumSymbols *NumStrNumberSymbols,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsMechanics).copyNumSymbols(
		nStrNumSym,
		incomingNumSymbols,
		ePrefix.XCpy(
			"nStrNumSym<-incomingNumSymbols"))
}

//	CopyOut
//
//	Returns a deep copy of the current NumStrNumberSymbols
//	instance.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
//	deepCopyNumSymbols			NumStrNumberSymbols
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the
//		current NumStrNumberSymbols instance.
//
//	err							error
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
func (nStrNumSym *NumStrNumberSymbols) CopyOut(
	errorPrefix interface{}) (
	deepCopyNumSymbols NumStrNumberSymbols,
	err error) {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNumSymbols, err
	}

	err = new(numStrNumberSymbolsMechanics).copyNumSymbols(
		&deepCopyNumSymbols,
		nStrNumSym,
		ePrefix.XCpy(
			"deepCopyNumSymbols<-nStrNumSym"))

	return deepCopyNumSymbols, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of NumStrNumberSymbols to their zero or
//	uninitialized states.
//
//	This method will leave the current instance of
//	NumStrNumberSymbols in an invalid state and
//	unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all member variable data
//	values in the current instance of NumStrNumberSymbols.
//	All member variable data values will be reset to their
//	zero or uninitialized states.
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
func (nStrNumSym *NumStrNumberSymbols) Empty() {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	new(numStrNumberSymbolsNanobot).empty(
		nStrNumSym)

	nStrNumSym.lock.Unlock()

	nStrNumSym.lock = nil
}

//	EmptyNegativeNumSymbols
//
//	Deletes and resets the Negative Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The target Negative Number Sign Symbol Specification
//	object is a member variable contained in the current
//	instance of NumStrNumberSymbols.
//
//	This action will put the Negative Number Sign Symbol
//	Specification into a NOP or No Operation state. In
//	this state, the Negative Number Sign Symbol
//	Specification will be treated as a placeholder and
//	ignored by Number String Formatting operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will modify the current instance of
//	NumStrNumberSymbols.
//
//	The Negative Number Sign Symbol member variable
//	data values will be reset to their zero or
//	uninitialized states.
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
func (nStrNumSym *NumStrNumberSymbols) EmptyNegativeNumSymbols() {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	new(numStrNumberSymbolsAtom).emptyNegativeNumSymbols(
		nStrNumSym)
}

//	EmptyPositiveNumSymbols
//
//	Deletes and resets the Positive Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The target Positive Number Sign Symbol Specification
//	object is a member variable contained in the current
//	instance of NumStrNumberSymbols.
//
//	This action will put the Positive Number Sign Symbol
//	Specification into a NOP or No Operation state. In
//	this state, the Positive Number Sign Symbol
//	Specification will be treated as a placeholder and
//	ignored by Number String Formatting operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will modify the current instance of
//	NumStrNumberSymbols.
//
//	The Positive Number Sign Symbol member variable
//	data values will be reset to their zero or
//	uninitialized states.
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
func (nStrNumSym *NumStrNumberSymbols) EmptyPositiveNumSymbols() {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	new(numStrNumberSymbolsAtom).emptyPositiveNumSymbols(
		nStrNumSym)
}

//	EmptyZeroNumSymbols
//
//	Deletes and resets the Zero Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The target Zero Number Sign Symbol Specification
//	object is a member variable contained in the current
//	instance of NumStrNumberSymbols.
//
//	This action will put the Zero Number Sign Symbol
//	Specification into a NOP or No Operation state. In
//	this state, the Zero Number Sign Symbol
//	Specification will be treated as a placeholder and
//	ignored by Number String Formatting operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will modify the current instance of
//	NumStrNumberSymbols.
//
//	The Zero Number Sign Symbol member variable
//	data values will be reset to their zero or
//	uninitialized states.
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
func (nStrNumSym *NumStrNumberSymbols) EmptyZeroNumSymbols() {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	new(numStrNumberSymbolsAtom).emptyZeroNumSymbols(
		nStrNumSym)
}

// IsNOP
//
//	'NOP' stands for 'No Operation'. This method signals
//	whether the three Number Symbol Specifications
//	contained in the current instance of
//	NumStrNumberSymbols are all valid, engaged and fully
//	operational with respect to Number String Formatting
//	operations.
//
//	Type NumStrNumberSymbols encapsulates the Negative,
//	Positive and Zero Number Symbol Specifications.
//
//	If this method returns 'true', it signals that all
//	three Number Symbol Specifications are simply empty
//	placeholders and perform no active role in, and are
//	completely ignored by, Number String Formatting
//	algorithms.
//
//	When 'NOP' is 'true', none of the three Number Symbol
//	Specifications are valid and all are used as empty
//	placeholders when formatting Number Strings.
//
//	If this method returns 'false', it signals that one
//	or more of the three Number Symbol Specifications
//	are in valid, engaged and fully operational with
//	respect to Number String Formatting operations.
func (nStrNumSym *NumStrNumberSymbols) IsNOP() bool {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	if nStrNumSym.negativeNumberSign.IsNOP() &&
		nStrNumSym.positiveNumberSign.IsNOP() &&
		nStrNumSym.zeroNumberSign.IsNOP() {

		return true
	}

	return false
}

//	IsNOPNegativeNumSymbols
//
//	'NOP' stands for 'No Operation'. This method
//	returns a boolean value signaling whether the
//	Negative Number Sign Symbol Specification contained
//	in the current NumStrNumberSymbols instance is
//	engaged, valid and fully operational with respect to
//	Number String Formatting.
//
//	If this method returns 'true', it signals that the
//	Negative Number Symbol Specification is simply an
//	empty placeholder and performs no active role in,
//	and is completely ignored by, Number String Formatting
//	algorithms. When 'NOP' is 'true', no Negative
//	Number Symbols will be inserted or formatted as part
//	of a Number String formatting operation.
//
//	If this method returns 'false', it signals that the
//	Negative Number Sign Symbol Specification is fully
//	populated, valid and functional. When 'NOP' is
//	'false', Number String Formatting operations will
//	include Negative Number Symbols in formatted number
//	strings.
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
//	bool
//
//		If this method returns 'true', it signals that
//		the Negative Number Symbol Specification is
//		simply an empty placeholder and performs no
//		active role in, and is completely ignored by,
//		Number String Formatting algorithms. This means,
//		no Negative Number Symbols will be inserted
//		or formatted as part of a Number String
//		formatting operation.
//
//		If this method returns 'false', it signals that
//		the Negative Number Sign Symbol Specification is
//		fully populated, valid and functional. Number
//		String Formatting operations will therefore
//		include these Negative Number Symbols in
//		formatted number strings.
func (nStrNumSym *NumStrNumberSymbols) IsNOPNegativeNumSymbols() bool {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	return nStrNumSym.negativeNumberSign.IsNOP()
}

//	IsNOPPositiveNumSymbols
//
//	'NOP' stands for 'No Operation'. This method
//	returns a boolean value signaling whether the
//	Positive Number Sign Symbol Specification contained
//	in the current NumStrNumberSymbols instance is
//	engaged, valid and fully operational with respect to
//	Number String Formatting.
//
//	If this method returns 'true', it signals that the
//	Positive Number Symbol Specification is simply an
//	empty placeholder and performs no active role in,
//	and is completely ignored by, Number String Formatting
//	algorithms. When 'NOP' is 'true', no Positive
//	Number Symbols will be inserted or formatted as part
//	of a Number String formatting operation.
//
//	If this method returns 'false', it signals that the
//	Positive Number Sign Symbol Specification is fully
//	populated, valid and functional. When 'NOP' is
//	'false', Number String Formatting operations will
//	include Positive Number Symbols in formatted number
//	strings.
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
//	bool
//
//		If this method returns 'true', it signals that
//		the Positive Number Symbol Specification is
//		simply an empty placeholder and performs no
//		active role in, and is completely ignored by,
//		Number String Formatting algorithms. This means,
//		no Positive Number Symbols will be inserted
//		or formatted as part of a Number String
//		formatting operation.
//
//		If this method returns 'false', it signals that
//		the Positive Number Sign Symbol Specification is
//		fully populated, valid and functional. Number
//		String Formatting operations will therefore
//		include these Positive Number Symbols in
//		formatted number strings.
func (nStrNumSym *NumStrNumberSymbols) IsNOPPositiveNumSymbols() bool {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	return nStrNumSym.positiveNumberSign.IsNOP()
}

//	IsNOPZeroNumSymbols
//
//	'NOP' stands for 'No Operation'. This method
//	returns a boolean value signaling whether the
//	Zero Number Sign Symbol Specification contained
//	in the current NumStrNumberSymbols instance is
//	engaged, valid and fully operational with respect to
//	Number String Formatting.
//
//	If this method returns 'true', it signals that the
//	Zero Number Symbol Specification is simply an
//	empty placeholder and performs no active role in,
//	and is completely ignored by, Number String Formatting
//	algorithms. When 'NOP' is 'true', no Zero
//	Number Symbols will be inserted or formatted as part
//	of a Number String formatting operation.
//
//	If this method returns 'false', it signals that the
//	Zero Number Sign Symbol Specification is fully
//	populated, valid and functional. When 'NOP' is
//	'false', Number String Formatting operations will
//	include Zero Number Symbols in formatted number
//	strings.
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
//	bool
//
//		If this method returns 'true', it signals that
//		the Zero Number Symbol Specification is
//		simply an empty placeholder and performs no
//		active role in, and is completely ignored by,
//		Number String Formatting algorithms. This means,
//		no Zero Number Symbols will be inserted
//		or formatted as part of a Number String
//		formatting operation.
//
//		If this method returns 'false', it signals that
//		the Zero Number Sign Symbol Specification is
//		fully populated, valid and functional. Number
//		String Formatting operations will therefore
//		include these Zero Number Symbols in
//		formatted number strings.
func (nStrNumSym *NumStrNumberSymbols) IsNOPZeroNumSymbols() bool {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	return nStrNumSym.zeroNumberSign.IsNOP()
}

//	NewSymbolsRunes
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	leadingNegativeNumberSymbols	[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "$-"
//			Number String:   "$-123.456"
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
//		('-') and/or currency symbols ('$').
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
//			Trailing Symbols: "-$"
//			Number String:   "123.456-$"
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
//	leadingZeroNumberSymbols	[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Leading Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
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
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Trailing Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " $"
//			Number String:   "0.00 $"
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
//	NumStrNumberSymbols
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbols configured
//		with the Positive, Negative and Zero Number
//		Sign Symbol Specification objects passed as
//		input parameters.
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
func (nStrNumSym *NumStrNumberSymbols) NewSymbolsRunes(
	leadingPositiveNumberSymbols []rune,
	trailingPositiveNumberSymbols []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegativeNumberSymbols []rune,
	trailingNegativeNumberSymbols []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumberSymbols []rune,
	trailingZeroNumberSymbols []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	NumStrNumberSymbols,
	error) {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumberSymbols NumStrNumberSymbols

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"NewSymbolsRunes()",
		"")

	if err != nil {
		return newNumberSymbols, err
	}

	nStrNumSymNanobot := numStrNumberSymbolsNanobot{}

	err = nStrNumSymNanobot.setPositiveNumSignRunes(
		&newNumberSymbols,
		leadingPositiveNumberSymbols,
		trailingPositiveNumberSymbols,
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))

	if err != nil {
		return newNumberSymbols, err
	}

	err = nStrNumSymNanobot.setNegativeNumSignRunes(
		&newNumberSymbols,
		leadingNegativeNumberSymbols,
		trailingNegativeNumberSymbols,
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))

	if err != nil {
		return newNumberSymbols, err
	}

	err = nStrNumSymNanobot.setZeroNumSignRunes(
		&newNumberSymbols,
		leadingZeroNumberSymbols,
		trailingZeroNumberSymbols,
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))

	return newNumberSymbols, err
}

//	NewSymbolsStrings
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingPositiveNumberSymbols	string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		a Positive Number Sign Symbol in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
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
//	trailingPositiveNumberSymbols	string
//
//		A string containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
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
//	leadingNegativeNumberSymbols	string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "$-"
//			Number String:   "$-123.456"
//
//	trailingNegativeNumberSymbols	string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
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
//			Trailing Symbols: "-$"
//			Number String:   "123.456-$"
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
//	leadingZeroNumberSymbols		string
//
//		A string containing the leading zero number sign
//		character or characters used to configure a	Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Leading Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//	trailingZeroNumberSymbols		string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Trailing Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//	 	currency symbols ('$').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " $"
//			Number String:   "0.00 $"
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
//	NumStrNumberSymbols
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbols configured
//		with the Positive, Negative and Zero Number
//		Sign Symbol Specification objects passed as
//		input parameters.
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
func (nStrNumSym *NumStrNumberSymbols) NewSymbolsStrings(
	leadingPositiveNumberSymbols string,
	trailingPositiveNumberSymbols string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegativeNumberSymbols string,
	trailingNegativeNumberSymbols string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumberSymbols string,
	trailingZeroNumberSymbols string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) (
	NumStrNumberSymbols,
	error) {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumberSymbols NumStrNumberSymbols

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"NewSymbolsStrings()",
		"")

	if err != nil {
		return newNumberSymbols, err
	}

	nStrNumSymNanobot := numStrNumberSymbolsNanobot{}

	err = nStrNumSymNanobot.setPositiveNumSignRunes(
		&newNumberSymbols,
		[]rune(leadingPositiveNumberSymbols),
		[]rune(trailingPositiveNumberSymbols),
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))

	if err != nil {
		return newNumberSymbols, err
	}

	err = nStrNumSymNanobot.setNegativeNumSignRunes(
		&newNumberSymbols,
		[]rune(leadingNegativeNumberSymbols),
		[]rune(trailingNegativeNumberSymbols),
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))

	if err != nil {
		return newNumberSymbols, err
	}

	err = nStrNumSymNanobot.setZeroNumSignRunes(
		&newNumberSymbols,
		[]rune(leadingZeroNumberSymbols),
		[]rune(trailingZeroNumberSymbols),
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))

	return newNumberSymbols, err
}

//	NewSymbolsSpecs
//
//	Creates and returns a new instance of
//	NumStrNumberSymbols. The new instance is generated
//	from the Positive, Negative and Zero Number Sign
//	Symbol Specification objects passed as input
//	parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Positive Number Sign Symbol Specification
//		will be copied to the corresponding Positive
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbols.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Negative Number Sign Symbol Specification
//		will be copied to the corresponding Negative
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbols.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbols.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
//	NumStrNumberSymbols
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbols configured
//		with the Positive, Negative and Zero Number
//		Sign Symbol Specification objects passed as
//		input parameters.
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
func (nStrNumSym *NumStrNumberSymbols) NewSymbolsSpecs(
	positiveNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) (
	NumStrNumberSymbols,
	error) {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumberSymbols NumStrNumberSymbols

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"NewSymbolsSpecs()",
		"")

	if err != nil {
		return newNumberSymbols, err
	}

	err = new(numStrNumberSymbolsMechanics).
		setNumSymbolSpecs(
			&newNumberSymbols,
			positiveNumberSign,
			negativeNumberSign,
			zeroNumberSign,
			ePrefix.XCpy(
				"newNumberSymbols"))

	return newNumberSymbols, err
}

//	SetNegativeNumSignSpec
//
//	Reconfigures the current instance of
//	NumStrNumberSymbols based on the Negative Number
//	Sign Symbol Specification object passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Negative Number Sign Symbol member variable data
//	fields contained in the current instance of
//	NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Negative Number Sign Symbol Specification
//		will be copied to the corresponding Negative
//		Symbol Specification in the current instance of
//		NumStrNumberSymbols.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetNegativeNumSignSpec(
	negativeNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetNegativeNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).
		setNegativeNumSignSpec(
			nStrNumSym,
			negativeNumberSign,
			ePrefix.XCpy(
				"nStrNumSym"))

}

//	SetNegativeSymbolsRunes
//
//	Deletes and resets the Negative Number Symbols data
//	fields for the current instance of NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Negative Number Sign Symbol member variable data
//	fields as specified by the Negative Number Symbol
//	input parameters passed as rune arrays.
//
//	The Negative Number Sign Symbol member variable for
//	the current instance of NumStrNumberSymbols is:
//
//		NumStrNumberSymbols.negativeNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		('-') and/or currency symbols ('$').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "$+"
//			Number String:   "$+123.456"
//
//		Example-3: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "$"
//			Number String:   "$123.456"
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
//		('-') and/or currency symbols ('$').
//
//		Example-1: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "+$"
//			Number String:   "123.456+$"
//
//		Example-3: Trailing Number Symbols
//			Trailing Number Symbols for Negative Values
//
//			Trailing Symbols: "$"
//			Number String:   "123.456$"
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetNegativeSymbolsRunes(
	leadingNegativeNumberSymbols []rune,
	trailingNegativeNumberSymbols []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetNegativeSymbolsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).setNegativeNumSignRunes(
		nStrNumSym,
		leadingNegativeNumberSymbols,
		trailingNegativeNumberSymbols,
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))
}

//	SetNegativeSymbolsStrings
//
//	Deletes and resets the Negative Number Symbols data
//	fields for the current instance of NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Negative Number Sign Symbol member variable data
//	fields as specified by the Negative Number Symbol
//	input parameters passed as strings.
//
//	The Negative Number Sign Symbol member variable for
//	the current instance of NumStrNumberSymbols is:
//
//		NumStrNumberSymbols.negativeNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingNegativeNumberSymbols	string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "$-"
//			Number String:   "$-123.456"
//
//	trailingNegativeNumberSymbols	string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
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
//			Trailing Symbols: "-$"
//			Number String:   "123.456-$"
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetNegativeSymbolsStrings(
	leadingNegativeNumberSymbols string,
	trailingNegativeNumberSymbols string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetNegativeSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).setNegativeNumSignRunes(
		nStrNumSym,
		[]rune(leadingNegativeNumberSymbols),
		[]rune(trailingNegativeNumberSymbols),
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))
}

//	SetPositiveNumSignSpec
//
//	Reconfigures the current instance of
//	NumStrNumberSymbols based on the Positive Number
//	Sign Symbol Specification object passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Positive Number Sign Symbol member variable data
//	fields contained in the current instance of
//	NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Positive Number Sign Symbol Specification
//		will be copied to the corresponding Positive
//		Symbol Specification in the current instance of
//		NumStrNumberSymbols.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetPositiveNumSignSpec(
	positiveNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetPositiveNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).
		setPositiveNumSignSpec(
			nStrNumSym,
			positiveNumberSign,
			ePrefix.XCpy(
				"nStrNumSym"))

}

//	SetPositiveSymbolsRunes
//
//	Deletes and resets the Positive Number Symbols data
//	fields for the current instance of NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Positive Number Sign Symbol member variable data
//	fields as specified by the Positive Number Symbol
//	input parameters passed as rune arrays.
//
//	The Positive Number Sign Symbol member variable for
//	the current instance of NumStrNumberSymbols is:
//
//		NumStrNumberSymbols.positiveNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetPositiveSymbolsRunes(
	leadingPositiveNumberSymbols []rune,
	trailingPositiveNumberSymbols []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetPositiveSymbolsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).setPositiveNumSignRunes(
		nStrNumSym,
		leadingPositiveNumberSymbols,
		trailingPositiveNumberSymbols,
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))
}

//	SetPositiveSymbolsStrings
//
//	Deletes and resets the Positive Number Symbols data
//	fields for the current instance of NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Positive Number Sign Symbol member variable data
//	fields as specified by the Positive Number Symbol
//	input parameters passed as strings.
//
//	The Positive Number Sign Symbol member variable for
//	the current instance of NumStrNumberSymbols is:
//
//		NumStrNumberSymbols.positiveNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingPositiveNumberSymbols	string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		a Positive Number Sign Symbol in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
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
//	trailingPositiveNumberSymbols	string
//
//		A string containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetPositiveSymbolsStrings(
	leadingPositiveNumberSymbols string,
	trailingPositiveNumberSymbols string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetPositiveSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).setPositiveNumSignRunes(
		nStrNumSym,
		[]rune(leadingPositiveNumberSymbols),
		[]rune(trailingPositiveNumberSymbols),
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))
}

//	SetZeroNumSignSpec
//
//	Reconfigures the current instance of
//	NumStrNumberSymbols based on the Zero Number
//	Sign Symbol Specification object passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Zero Number Sign Symbol member variable data
//	fields contained in the current instance of
//	NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	zeroNumberSign				NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Symbol Specification in the current instance of
//		NumStrNumberSymbols.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetZeroNumSignSpec(
	positiveNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetZeroNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).
		setZeroNumSignSpec(
			nStrNumSym,
			positiveNumberSign,
			ePrefix.XCpy(
				"nStrNumSym"))

}

//	SetSymbolsRunes
//
//	Reconfigures the current instance of
//	NumStrNumberSymbols based on the Positive, Negative
//	and Zero Number	Sign Symbol rune arrays passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbols will be
//	deleted and replaced with the Positive, Negative
//	and Zero Number	Sign Symbol input parameters
//	passed as rune arrays.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	leadingNegativeNumberSymbols	[]rune
//
//		A rune array containing the leading negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "$-"
//			Number String:   "$-123.456"
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
//		('-') and/or currency symbols ('$').
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
//			Trailing Symbols: "-$"
//			Number String:   "123.456-$"
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
//	leadingZeroNumberSymbols	[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Leading Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
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
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Trailing Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " $"
//			Number String:   "0.00 $"
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetSymbolsRunes(
	leadingPositiveNumberSymbols []rune,
	trailingPositiveNumberSymbols []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegativeNumberSymbols []rune,
	trailingNegativeNumberSymbols []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumberSymbols []rune,
	trailingZeroNumberSymbols []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetSymbolsRunes()",
		"")

	if err != nil {
		return err
	}

	nStrNumSymNanobot := numStrNumberSymbolsNanobot{}

	err = nStrNumSymNanobot.setPositiveNumSignRunes(
		nStrNumSym,
		leadingPositiveNumberSymbols,
		trailingPositiveNumberSymbols,
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))

	if err != nil {
		return err
	}

	err = nStrNumSymNanobot.setNegativeNumSignRunes(
		nStrNumSym,
		leadingNegativeNumberSymbols,
		trailingNegativeNumberSymbols,
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))

	if err != nil {
		return err
	}

	return nStrNumSymNanobot.setZeroNumSignRunes(
		nStrNumSym,
		leadingZeroNumberSymbols,
		trailingZeroNumberSymbols,
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))

}

//	SetSymbolsSpecs
//
//	Reconfigures the current instance of
//	NumStrNumberSymbols based on the Positive, Negative
//	and Zero Number	Sign Symbol Specification objects
//	passed as input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbols will be
//	deleted and replaced with the Positive, Negative
//	and Zero Number	Sign Symbol Specifications passed
//	as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Positive Number Sign Symbol Specification
//		will be copied to the corresponding Positive
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbols.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Negative Number Sign Symbol Specification
//		will be copied to the corresponding Negative
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbols.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbols.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetSymbolsSpecs(
	positiveNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetSymbolsSpecs()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsMechanics).
		setNumSymbolSpecs(
			nStrNumSym,
			positiveNumberSign,
			negativeNumberSign,
			zeroNumberSign,
			ePrefix.XCpy(
				"nStrNumSym"))

}

//	SetSymbolsStrings
//
//	Reconfigures the current instance of
//	NumStrNumberSymbols based on the Positive, Negative
//	and Zero Number	Sign Symbol strings passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbols will be
//	deleted and replaced with the Positive, Negative
//	and Zero Number	Sign Symbol input parameters
//	passed as strings.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingPositiveNumberSymbols	string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		a Positive Number Sign Symbol in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
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
//	trailingPositiveNumberSymbols	string
//
//		A string containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
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
//	leadingNegativeNumberSymbols	string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "$-"
//			Number String:   "$-123.456"
//
//	trailingNegativeNumberSymbols	string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
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
//			Trailing Symbols: "-$"
//			Number String:   "123.456-$"
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
//	leadingZeroNumberSymbols		string
//
//		A string containing the leading zero number sign
//		character or characters used to configure a	Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Leading Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//	trailingZeroNumberSymbols		string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Trailing Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//	 	currency symbols ('$').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " $"
//			Number String:   "0.00 $"
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
//	NumStrNumberSymbols
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbols configured
//		with the Positive, Negative and Zero Number
//		Sign Symbol Specification objects passed as
//		input parameters.
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
func (nStrNumSym *NumStrNumberSymbols) SetSymbolsStrings(
	leadingPositiveNumberSymbols string,
	trailingPositiveNumberSymbols string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegativeNumberSymbols string,
	trailingNegativeNumberSymbols string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumberSymbols string,
	trailingZeroNumberSymbols string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumberSymbols NumStrNumberSymbols

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	nStrNumSymNanobot := numStrNumberSymbolsNanobot{}

	err = nStrNumSymNanobot.setPositiveNumSignRunes(
		&newNumberSymbols,
		[]rune(leadingPositiveNumberSymbols),
		[]rune(trailingPositiveNumberSymbols),
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))

	if err != nil {
		return err
	}

	err = nStrNumSymNanobot.setNegativeNumSignRunes(
		&newNumberSymbols,
		[]rune(leadingNegativeNumberSymbols),
		[]rune(trailingNegativeNumberSymbols),
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))

	if err != nil {
		return err
	}

	return nStrNumSymNanobot.setZeroNumSignRunes(
		&newNumberSymbols,
		[]rune(leadingZeroNumberSymbols),
		[]rune(trailingZeroNumberSymbols),
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))
}

//	SetZeroSymbolsRunes
//
//	Deletes and resets the Zero Number Symbols data
//	fields for the current instance of NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Zero Number Sign Symbol member variable data
//	fields as specified by the Zero Number Symbol
//	input parameters passed as rune arrays.
//
//	The Zero Number Sign Symbol member variable for
//	the current instance of NumStrNumberSymbols is:
//
//		NumStrNumberSymbols.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Leading Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
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
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Trailing Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " $"
//			Number String:   "0.00 $"
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetZeroSymbolsRunes(
	leadingZeroNumberSymbols []rune,
	trailingZeroNumberSymbols []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetZeroSymbolsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).setZeroNumSignRunes(
		nStrNumSym,
		leadingZeroNumberSymbols,
		trailingZeroNumberSymbols,
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))
}

//	SetZeroSymbolsStrings
//
//	Deletes and resets the Zero Number Symbols data
//	fields for the current instance of NumStrNumberSymbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Zero Number Sign Symbol member variable data
//	fields as specified by the Zero Number Symbol
//	input parameters passed as strings.
//
//	The Zero Number Sign Symbol member variable for
//	the current instance of NumStrNumberSymbols is:
//
//		NumStrNumberSymbols.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingZeroNumberSymbols		string
//
//		A string containing the leading zero number sign
//		character or characters used to configure a	Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Leading Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//	trailingZeroNumberSymbols		string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Trailing Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//	 	currency symbols ('$').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " $"
//			Number String:   "0.00 $"
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
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
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
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
func (nStrNumSym *NumStrNumberSymbols) SetZeroSymbolsStrings(
	leadingZeroNumberSymbols string,
	trailingZeroNumberSymbols string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSym.lock == nil {
		nStrNumSym.lock = new(sync.Mutex)
	}

	nStrNumSym.lock.Lock()

	defer nStrNumSym.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbols."+
			"SetZeroSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolsNanobot).setZeroNumSignRunes(
		nStrNumSym,
		[]rune(leadingZeroNumberSymbols),
		[]rune(trailingZeroNumberSymbols),
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))
}

// numStrNumberSymbolsMechanics
//
// Provides helper methods for NumStrNumberSymbols.
type numStrNumberSymbolsMechanics struct {
	lock *sync.Mutex
}

//	copyNumSymbols
//
//	Copies all Number Symbol Specification data
//	from input parameter 'sourceNumSymbols' to
//	input parameter 'destinationNumSymbols'.
//	Both instances are of type
//	NumStrNumberSymbols.
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
//	# Input Parameters
//
//	destinationNumSymbols			*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
//		All the member variable data fields in this object will be
//		replaced by data values copied from input parameter
//		'sourceNumSymbolSpec'.
//
//		'destinationNumSymbolSpec' is the destination for this
//		copy operation.
//
//	sourceNumSymbols				*NumStrNumberSymbols
//
//		A pointer to another instance of NumStrNumberSymbols.
//		All the member variable data values from this object
//		will be copied to corresponding member variables in
//		'destinationNumSymbols'.
//
//		'sourceNumSymbolSpec' is the source for this copy
//		operation.
//
//		No data validation is performed on 'sourceNumSymbols'.
//
//		No data values contained in 'sourceNumSymbols' will
//		be modified.
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
func (nStrNumSymMech *numStrNumberSymbolsMechanics) copyNumSymbols(
	destinationNumSymbols *NumStrNumberSymbols,
	sourceNumSymbols *NumStrNumberSymbols,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymMech.lock == nil {
		nStrNumSymMech.lock = new(sync.Mutex)
	}

	nStrNumSymMech.lock.Lock()

	defer nStrNumSymMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolsMechanics."+
			"copyNumSymbols()",
		"")

	if err != nil {
		return err
	}

	if destinationNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNumSymbols' is invalid!\n"+
			"'destinationNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceNumSymbols == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNumSymbols' is invalid!\n"+
			"'sourceNumSymbols' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrNumberSymbolsNanobot).empty(destinationNumSymbols)

	err = destinationNumSymbols.positiveNumberSign.CopyIn(
		&sourceNumSymbols.positiveNumberSign,
		ePrefix.XCpy(
			"destinationNumSymbols<-"+
				"sourceNumSymbols.positiveNumberSign"))

	if err != nil {
		return err
	}

	err = destinationNumSymbols.negativeNumberSign.CopyIn(
		&sourceNumSymbols.negativeNumberSign,
		ePrefix.XCpy(
			"destinationNumSymbols<-"+
				"sourceNumSymbols.negativeNumberSign"))

	if err != nil {
		return err
	}

	err = destinationNumSymbols.zeroNumberSign.CopyIn(
		&sourceNumSymbols.zeroNumberSign,
		ePrefix.XCpy(
			"destinationNumSymbols<-"+
				"sourceNumSymbols.zeroNumberSign"))

	return err
}

//	setNumSymbolSpecs
//
//	Receives three NumStrNumberSymbolSpec objects and
//	proceeds to reset the corresponding member variable
//	data values for the NumStrNumberSymbols input
//	paramter 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset all
//	the member variable data fields in input paramter,
//	'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
//		The Positive Number Sign, Negative Number Sign and
//		Zero Number Sign Symbol Specifications for this
//		instance will be deleted and reset to the values
//		provided by input parameters 'positiveNumberSign',
//		'negativeNumberSign' and 'zeroNumberSign'.
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Positive Number Sign Symbol Specification
//		will be copied to the corresponding Positive
//		Symbol Specification in input paramter,
//		'nStrNumSymbols'.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Negative Number Sign Symbol Specification
//		will be copied to the corresponding Negative
//		Symbol Specification in input paramter,
//		'nStrNumSymbols'.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Symbol Specification in input paramter,
//		'nStrNumSymbols'.
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
func (nStrNumSymMech *numStrNumberSymbolsMechanics) setNumSymbolSpecs(
	nStrNumSymbols *NumStrNumberSymbols,
	positiveNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymMech.lock == nil {
		nStrNumSymMech.lock = new(sync.Mutex)
	}

	nStrNumSymMech.lock.Lock()

	defer nStrNumSymMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolsMechanics."+
			"setNumSymbolSpecs()",
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

	err = nStrNumSymbols.positiveNumberSign.CopyIn(
		&positiveNumberSign,
		ePrefix.XCpy(
			"nStrNumSymbols.positiveNumberSign<-"+
				"positiveNumberSign"))

	if err != nil {
		return err
	}

	err = nStrNumSymbols.negativeNumberSign.CopyIn(
		&negativeNumberSign,
		ePrefix.XCpy(
			"nStrNumSymbols.negativeNumberSign<-"+
				"negativeNumberSign"))

	if err != nil {
		return err
	}

	err = nStrNumSymbols.zeroNumberSign.CopyIn(
		&zeroNumberSign,
		ePrefix.XCpy(
			"nStrNumSymbols.zeroNumberSign<-"+
				"zeroNumberSign"))

	return err
}

// numStrNumberSymbolsNanobot
//
// Provides helper methods for NumStrNumberSymbols.
type numStrNumberSymbolsNanobot struct {
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
func (nStrNumSymNanobot *numStrNumberSymbolsNanobot) empty(
	nStrNumSymbols *NumStrNumberSymbols) {

	if nStrNumSymNanobot.lock == nil {
		nStrNumSymNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymNanobot.lock.Lock()

	defer nStrNumSymNanobot.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.positiveNumberSign.Empty()

	nStrNumSymbols.negativeNumberSign.Empty()

	nStrNumSymbols.zeroNumberSign.Empty()
}

//	setNegativeNumSignRunes
//
//	Receives a series of rune arrays used to configure
//	the input parameter 'nStrNumSymbols' with new data
//	values for the Negative Number Sign Symbols.
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
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols					*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
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
//		('-') and/or currency symbols ('$').
//
//		Example-1: Leading Number Symbols
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Number Symbols With Currency
//			Leading Number Symbols for Negative Values
//
//			Leading Symbols: "$-"
//			Number String:   "$-123.456"
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
//		('-') and/or currency symbols ('$').
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
//			Trailing Symbols: "-$"
//			Number String:   "123.456-$"
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
func (nStrNumSymNanobot *numStrNumberSymbolsNanobot) setNegativeNumSignRunes(
	nStrNumSymbols *NumStrNumberSymbols,
	leadingNegativeNumberSymbols []rune,
	trailingNegativeNumberSymbols []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymNanobot.lock == nil {
		nStrNumSymNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymNanobot.lock.Lock()

	defer nStrNumSymNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolsNanobot."+
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

	err = nStrNumSymbols.negativeNumberSign.
		SetLeadingTrailingNumberSymbolRunes(
			leadingNegativeNumberSymbols,
			negativeNumFieldSymPosition,
			trailingNegativeNumberSymbols,
			negativeNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumSymbols.negativeNumberSign"))

	return err
}

//	setNegativeNumSignSpec
//
//	Receives a single NumStrNumberSymbolSpec object
//	configured as a Negative Number Sign Symbol.
//
//	This method then proceeds to reset the
//	corresponding Negative Number Sign Symbol member
//	variable data value for the NumStrNumberSymbols
//	input paramter 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Negative Number Sign Symbol member variable data fields
//	contained in input paramter, 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
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
//		'nStrNumSymbols'.
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
func (nStrNumSymNanobot *numStrNumberSymbolsNanobot) setNegativeNumSignSpec(
	nStrNumSymbols *NumStrNumberSymbols,
	negativeNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymNanobot.lock == nil {
		nStrNumSymNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymNanobot.lock.Lock()

	defer nStrNumSymNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolsNanobot."+
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
//	nStrNumSymbols					*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
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
func (nStrNumSymNanobot *numStrNumberSymbolsNanobot) setPositiveNumSignRunes(
	nStrNumSymbols *NumStrNumberSymbols,
	leadingPositiveNumberSymbols []rune,
	trailingPositiveNumberSymbols []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymNanobot.lock == nil {
		nStrNumSymNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymNanobot.lock.Lock()

	defer nStrNumSymNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolsNanobot."+
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

	err = nStrNumSymbols.positiveNumberSign.
		SetLeadingTrailingNumberSymbolRunes(
			leadingPositiveNumberSymbols,
			positiveNumFieldSymPosition,
			trailingPositiveNumberSymbols,
			positiveNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumSymbols.positiveNumberSign"))

	return err
}

//	setPositiveNumSignSpec
//
//	Receives a single NumStrNumberSymbolSpec object
//	configured as a Positive Number Sign Symbol.
//
//	This method then proceeds to reset the
//	corresponding Positive Number Sign Symbol member
//	variable data value for the NumStrNumberSymbols
//	input paramter 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Positive Number Sign Symbol member variable data fields
//	contained in input paramter, 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
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
//		'nStrNumSymbols'.
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
func (nStrNumSymNanobot *numStrNumberSymbolsNanobot) setPositiveNumSignSpec(
	nStrNumSymbols *NumStrNumberSymbols,
	positiveNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymNanobot.lock == nil {
		nStrNumSymNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymNanobot.lock.Lock()

	defer nStrNumSymNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolsNanobot."+
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
//	nStrNumSymbols					*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
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
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Leading Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
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
//		positive or negative value. However, currency
//		symbols may be required for zero values.
//		Therefore, Trailing Zero Number Symbols almost
//		always consist of currency symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " $"
//			Number String:   "0.00 $"
//
//	zeroNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Zero Number Sign
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
func (nStrNumSymNanobot *numStrNumberSymbolsNanobot) setZeroNumSignRunes(
	nStrNumSymbols *NumStrNumberSymbols,
	leadingZeroNumberSymbols []rune,
	trailingZeroNumberSymbols []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymNanobot.lock == nil {
		nStrNumSymNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymNanobot.lock.Lock()

	defer nStrNumSymNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolsNanobot."+
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

	err = nStrNumSymbols.zeroNumberSign.
		SetLeadingTrailingNumberSymbolRunes(
			leadingZeroNumberSymbols,
			zeroNumFieldSymPosition,
			trailingZeroNumberSymbols,
			zeroNumFieldSymPosition,
			ePrefix.XCpy(
				"nStrNumSymbols.zeroNumberSign"))

	return err
}

//	setZeroNumSignSpec
//
//	Receives a single NumStrNumberSymbolSpec object
//	configured as a Zero Number Sign Symbol.
//
//	This method then proceeds to reset the
//	corresponding Zero Number Sign Symbol member
//	variable data value for the NumStrNumberSymbols
//	input paramter 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Zero Number Sign Symbol member variable data fields
//	contained in input paramter, 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbols
//
//		A pointer to an instance of NumStrNumberSymbols.
//		The Zero Number Sign for this instance will be
//		deleted and reset to the values	provided by input
//		parameter 'zeroNumberSign'.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Symbol Specification in input paramter,
//		'nStrNumSymbols'.
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
func (nStrNumSymNanobot *numStrNumberSymbolsNanobot) setZeroNumSignSpec(
	nStrNumSymbols *NumStrNumberSymbols,
	zeroNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymNanobot.lock == nil {
		nStrNumSymNanobot.lock = new(sync.Mutex)
	}

	nStrNumSymNanobot.lock.Lock()

	defer nStrNumSymNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolsNanobot."+
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

	err = nStrNumSymbols.zeroNumberSign.CopyIn(
		&zeroNumberSign,
		ePrefix.XCpy(
			"nStrNumSymbols.zeroNumberSign<-"+
				"zeroNumberSign"))

	return err
}

// numStrNumberSymbolsAtom
//
// Provides helper methods for NumStrNumberSymbols.
type numStrNumberSymbolsAtom struct {
	lock *sync.Mutex
}

//	emptyNegativeNumSymbols
//
//	Deletes and resets the Negative Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The Negative Number Sign Symbol Specification object
//	is a member variable in the 'nStrNumSymbols' instance
//	passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values contained in the  Negative Number
//	Sign Symbol Specification will be deleted and reset
//	to their zero or uninitialized values. This Negative
//	Number Sign Symbol Specification member variable
//	is identified as:
//
//		nStrNumSymbols.negativeNumberSign
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
func (nStrNumSymbolsAtom *numStrNumberSymbolsAtom) emptyNegativeNumSymbols(
	nStrNumSymbols *NumStrNumberSymbols) {

	if nStrNumSymbolsAtom.lock == nil {
		nStrNumSymbolsAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolsAtom.lock.Lock()

	defer nStrNumSymbolsAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.negativeNumberSign.Empty()
}

//	emptyPositiveNumSymbols
//
//	Deletes and resets the Positive Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The Positive Number Sign Symbol Specification object
//	is a member variable in the 'nStrNumSymbols' instance
//	passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values contained in the  Positive Number
//	Sign Symbol Specification will be deleted and reset
//	to their zero or uninitialized values. This Positive
//	Number Sign Symbol Specification member variable
//	is identified as:
//
//		nStrNumSymbols.positiveNumberSign
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
func (nStrNumSymbolsAtom *numStrNumberSymbolsAtom) emptyPositiveNumSymbols(
	nStrNumSymbols *NumStrNumberSymbols) {

	if nStrNumSymbolsAtom.lock == nil {
		nStrNumSymbolsAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolsAtom.lock.Lock()

	defer nStrNumSymbolsAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.positiveNumberSign.Empty()
}

//	emptyZeroNumSymbols
//
//	Deletes and resets the Zero Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The Zero Number Sign Symbol Specification object
//	is a member variable in the 'nStrNumSymbols' instance
//	passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values contained in the Zero Number
//	Sign Symbol Specification will be deleted and reset
//	to their zero or uninitialized values. This Zero
//	Number Sign Symbol Specification member variable
//	is identified as:
//
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
func (nStrNumSymbolsAtom *numStrNumberSymbolsAtom) emptyZeroNumSymbols(
	nStrNumSymbols *NumStrNumberSymbols) {

	if nStrNumSymbolsAtom.lock == nil {
		nStrNumSymbolsAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolsAtom.lock.Lock()

	defer nStrNumSymbolsAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.zeroNumberSign.Empty()
}

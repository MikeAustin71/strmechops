package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

//	NumStrNumberSymbolGroup
//
//	This type is used to configure Number Symbols
//	required in converting numeric values to Number
//	Strings.
//
//	NumStrNumberSymbolGroup contains four instances of
//	NumStrNumberSymbolSpec defining the Number Symbols to
//	be used with positive numeric values, negative numeric
//	values, zero numeric values and currency.
//
// ----------------------------------------------------------------
//
// # Background
//
//	Each Number String Number Symbol Specification
//	contained within this type includes specifications
//	for both leading and trailing characters.
//
//	As such, 'Number Symbols' can include number signs
//	such as plus signs ('+'), minus signs ('-') and
//	parentheses ('()'),	as well as currency symbols such
//	as the Dollar Signs ('$'), Euro sign ('€') and Pound
//	sign ('£').
//
//	The user has the option to configure both leading and
//	trailing number symbols, leading number symbols only
//	or trailing number symbols only.
//
//	A Number Symbol is defined as one or more text
//	characters.
//
//	For Number Symbol examples, see the source code
//	documentation for NumStrNumberSymbolGroup member
//	variables listed below.
//
// ----------------------------------------------------------------
// # Usage
//
//	Configuring formatting for Number Stings while
//	supporting multinational and multicultural standards
//	necessarily requires a complex series of parameters
//	and specifications.
//
//	Configuring the Number Symbols which comprise part of
//	the Number String formatting process also involves
//	complexity.
//
//	Typically, instances of NumStrNumberSymbolGroup are
//	created, or constructed, using the 'New' methods
//	documented below. Many of these methods provide
//	input parameters capable of detailing all the
//	Number Symbol features required to support
//	any multinational and multicultural Number Symbol
//	formatting requirement.
//
//	For those only interested in a quick and simple means
//	of generating Number Symbol formatting, the following
//	methods provide defaults which greatly simplify the
//	Number Symbols creation process:
//
//		NumStrNumberSymbolGroup.NewCurrencySimple()
//		NumStrNumberSymbolGroup.NewSignedSimpleNumber()
//
//	If more granular control is required to meet
//	specialized requirements for multinational or
//	multicultural number symbol formatting, consider
//	using one of the following methods:
//
//		NumStrNumberSymbolGroup.NewSymbolsRunes()
//		NumStrNumberSymbolGroup.NewSymbolsStrings()
//		NumStrNumberSymbolGroup.NewSymbolsSpecs()
//		NumStrNumberSymbolGroup.SetSymbolsRunes()
//		NumStrNumberSymbolGroup.SetSymbolsStrings()
//		NumStrNumberSymbolGroup.SetSymbolsSpecs()
type NumStrNumberSymbolGroup struct {
	negativeNumberSign NumStrNumberSymbolSpec
	//	The Number String Negative Number Sign
	//	Specification is used to configure negative
	//	number sign symbols for negative numeric
	//	values formatted and displayed in number
	//	stings.
	//
	//	Example-1: Leading Number Sign Symbols
	//		Leading Number Sign Symbols for Negative
	//		Values
	//
	//		Leading Symbols: "- "
	//		Number String:   "- 123.456"
	//
	//	Example-2: Leading Number Sign Symbols
	//		Leading Number Sign Symbols for Negative
	//		Values
	//
	//		Leading Symbols: "-"
	//		Number String:   "-123.456"
	//
	//	Example-3: Trailing Number Sign Symbols
	//		Trailing Number Sign Symbols for Negative
	//		Values
	//
	//		Trailing Symbols: " -"
	//		Number String:   "123.456 -"
	//
	//	Example-4: Trailing Number Sign Symbols
	//		Trailing Number Sign Symbols for Negative
	//		Values
	//
	//		Trailing Symbols: "-"
	//		Number String:   "123.456-"

	positiveNumberSign NumStrNumberSymbolSpec
	//	Positive number signs are commonly implied
	//	and not specified. However, the user has
	//	the option to specify a positive number sign
	//	character or characters for positive numeric
	//	values using a Number String Positive Number
	//	Sign Specification.
	//
	//	Example-1: Leading Number Sign Symbols
	//		Leading Number Sign Symbols for Positive
	//		Values
	//
	//		Leading Symbols: "+ "
	//		Number String:   "+ 123.456"
	//
	//	Example-2: Leading Number Sign Symbols
	//		Leading Number Sign Symbols for Positive
	//		Values
	//
	//		Leading Symbols: "+"
	//		Number String:   "+123.456"
	//
	//	Example-3: Trailing Number Sign Symbols
	//		Trailing Number Sign Symbols for Positive
	//		Values
	//
	//		Trailing Symbols: " +"
	//		Number String:   "123.456 +"
	//
	//	Example-4: Trailing Number Sign Symbols
	//		Trailing Number Sign Symbols for Positive
	//		Values
	//
	//		Trailing Symbols: "+"
	//		Number String:   "123.456+"

	zeroNumberSign NumStrNumberSymbolSpec
	//	The Number String Zero Number Sign
	//	Specification is used to configure number
	//	sign symbols for zero numeric values formatted
	//	and displayed in number stings. Zero number
	//	signs are commonly omitted because zero
	//	does not technically qualify as either a
	//	positive or negative value. However,
	//	the user has the option to configure number
	//	sign symbols for zero values if necessary.
	//
	//	Example-1: Leading Number Sign Symbols
	//		Leading Number Sign Symbols for Zero Values
	//
	//		Leading Symbols: "+"
	//		Trailing Symbols: ""
	//		Number String:   "+0.00"
	//
	//	Example-2: Leading Number Sign Symbols
	//		Leading Number Sign Symbols for Zero Values
	//
	//		Leading Symbols: "+ "
	//		Trailing Symbols: ""
	//		Number String:   "+ 0.00"
	//
	//	Example-3: Trailing Number Sign Symbols
	//		Trailing Number Sign Symbols for Zero Values
	//
	//		Leading Symbols: ""
	//		Trailing Symbols: " +"
	//		Number String:   "0.00 +"
	//
	//	Example-4: Trailing Number Sign Symbols
	//		Trailing Number Sign Symbols for Zero Values
	//
	//		Leading Symbols: ""
	//		Trailing Symbols: "+"
	//		Number String:   "0.00+"

	currencySymbol NumStrNumberSymbolSpec
	//	A Currency Symbol next to a number shows the
	//	number is a monetary amount.
	//
	//	Examples of Currency Symbols include the Dollar
	//	sign ('$'), Euro sign ('€') or Pound sign ('£').
	//
	//	This instance of NumStrNumberSymbolSpec is used
	//	to configure leading Currency Symbols, trailing
	//	Currency Symbols or both leading and trailing
	//	Currency Symbols.
	//
	//	Example-1: Leading Currency Symbols
	//
	//		Leading Currency Symbols: "$ "
	//		Number String:   "$ 123.456"
	//
	//	Example-2: Leading Currency Symbols
	//
	//		Leading Currency Symbols: "$"
	//		Number String:   "$123.456"
	//
	//	Example-3: Trailing Currency Symbols
	//		Trailing Currency Symbols for Positive Values
	//
	//		Trailing Currency Symbols: "€"
	//		Number String:   "123.456€"
	//
	//	Example-4: Trailing Currency Symbols
	//		Trailing Currency Symbols for Positive Values
	//
	//		Trailing Currency Symbols: " €"
	//		Number String:   "123.456 €"

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the data fields from an incoming instance of
//	NumStrNumberSymbolGroup ('incomingNumSymbols')
//	to the data fields of the current NumStrNumberSymbolGroup
//	instance ('nStrNumSym').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the data fields in current NumStrNumberSymbolGroup
//	instance ('nStrNumSym') will be deleted and
//	overwritten.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	incomingNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		This method will NOT change the data values of
//		internal member variables contained in this
//		instance.
//
//		All data values in this NumStrNumberSymbolGroup
//		instance will be copied to current
//		NumStrNumberSymbolGroup instance ('nStrNumSym').
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) CopyIn(
	incomingNumSymbols *NumStrNumberSymbolGroup,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).copyNumSymbols(
		nStrNumSymbolsGroup,
		incomingNumSymbols,
		ePrefix.XCpy(
			"nStrNumSymbolsGroup<-incomingNumSymbols"))
}

//	CopyOut
//
//	Returns a deep copy of the current NumStrNumberSymbolGroup
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
//	deepCopyNumSymbols			NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the
//		current NumStrNumberSymbolGroup instance.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) CopyOut(
	errorPrefix interface{}) (
	deepCopyNumSymbols NumStrNumberSymbolGroup,
	err error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).copyNumSymbols(
		&deepCopyNumSymbols,
		nStrNumSymbolsGroup,
		ePrefix.XCpy(
			"deepCopyNumSymbols<-nStrNumSymbolsGroup"))

	return deepCopyNumSymbols, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of NumStrNumberSymbolGroup to their zero or
//	uninitialized states.
//
//	This method will leave the current instance of
//	NumStrNumberSymbolGroup in an invalid state and
//	unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all member variable data
//	values in the current instance of NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) Empty() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	new(numStrNumberSymbolGroupNanobot).empty(
		nStrNumSymbolsGroup)

	nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.lock = nil
}

//	EmptyNegativeNumSymbols
//
//	Deletes and resets the Negative Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The target Negative Number Sign Symbol Specification
//	object is a member variable contained in the current
//	instance of NumStrNumberSymbolGroup.
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
//	NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) EmptyNegativeNumSymbols() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	new(numStrNumberSymbolGroupAtom).emptyNegativeNumSymbols(
		nStrNumSymbolsGroup)
}

//	EmptyPositiveNumSymbols
//
//	Deletes and resets the Positive Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The target Positive Number Sign Symbol Specification
//	object is a member variable contained in the current
//	instance of NumStrNumberSymbolGroup.
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
//	NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) EmptyPositiveNumSymbols() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	new(numStrNumberSymbolGroupAtom).emptyPositiveNumSymbols(
		nStrNumSymbolsGroup)
}

//	EmptyZeroNumSymbols
//
//	Deletes and resets the Zero Number Sign Symbol
//	Specification to its zero or uninitialized state.
//
//	The target Zero Number Sign Symbol Specification
//	object is a member variable contained in the current
//	instance of NumStrNumberSymbolGroup.
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
//	NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) EmptyZeroNumSymbols() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	new(numStrNumberSymbolGroupAtom).emptyZeroNumSymbols(
		nStrNumSymbolsGroup)
}

//	Equal
//
//	Receives a pointer to another instance of
//	NumStrNumberSymbolGroup and proceeds to compare its
//	internal member variables to those of the current
//	NumStrNumberSymbolGroup instance in order to determine if
//	they are equivalent.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables for both
//	instances are equal in all respects, this flag is set
//	to 'true'. Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingNumSymbols		*NumStrNumberSymbolGroup
//
//		A pointer to an external instance of
//		NumStrNumberSymbolGroup. The member variable data
//		values in this instance will be compared to those
//		in the current instance of NumStrNumberSymbolGroup.
//		The results of this comparison will be returned
//		to the calling function as a boolean value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//
//		If the internal member variable data values contained in
//		input parameter 'incomingNumSymbols' are equivalent
//		in all respects to those contained in the current
//		instance of 'NumStrNumberSymbolGroup', this return value
//		will be set to 'true'.
//
//		Otherwise, this method will return 'false'.
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) Equal(
	incomingNumSymbols *NumStrNumberSymbolGroup) bool {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	return new(numStrNumberSymbolGroupNanobot).equal(
		nStrNumSymbolsGroup,
		incomingNumSymbols)
}

//	GetNegativeNumberSignSpec
//
//	Returns a deep copy of the NumStrNumberSymbolSpec
//	configured for negative number signs.
//
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
//		Number String:   "123.456-€"
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	negativeNumSignSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolSpec configured
//		with the Negative Number Sign Symbol configured
//		for the current instance of
//		NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) GetNegativeNumberSignSpec(
	errorPrefix interface{}) (
	negativeNumSignSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"GetNegativeNumberSignSpec()",
		"")

	if err != nil {
		return negativeNumSignSpec, err
	}

	negativeNumSignSpec,
		err = nStrNumSymbolsGroup.negativeNumberSign.CopyOut(
		ePrefix.XCpy(
			"negativeNumSignSpec<-"))

	return negativeNumSignSpec, err
}

//	GetPositiveNumberSignSpec
//
//	Returns a deep copy of the NumStrNumberSymbolSpec
//	configured for positive number signs.
//
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
//		Trailing Symbols: "+€"
//		Number String:   "123.456+€"
//
//	Example-6: Trailing Number Symbols
//		Trailing Number Symbols for Positive Values
//
//		Trailing Symbols: " €"
//		Number String:   "123.456 €"
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	positiveNumSignSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolSpec configured
//		with the Positive Number Sign Symbol configured
//		for the current instance of
//		NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) GetPositiveNumberSignSpec(
	errorPrefix interface{}) (
	positiveNumSignSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"GetPositiveNumberSignSpec()",
		"")

	if err != nil {
		return positiveNumSignSpec, err
	}

	positiveNumSignSpec,
		err = nStrNumSymbolsGroup.positiveNumberSign.CopyOut(
		ePrefix.XCpy(
			"positiveNumSignSpec<-"))

	return positiveNumSignSpec, err
}

//	GetZeroNumberSignSpec
//
//	Returns a deep copy of the
//	NumStrNumberSymbolSpec configured for zero
//	number signs.
//
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
//		Trailing Symbols: " €"
//		Number String:   "0.00 €"
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	zeroNumSignSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolSpec configured
//		with the Zero Number Sign Symbol configured
//		for the current instance of
//		NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) GetZeroNumberSignSpec(
	errorPrefix interface{}) (
	zeroNumSignSpec NumStrNumberSymbolSpec,
	err error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"GetZeroNumberSignSpec()",
		"")

	if err != nil {
		return zeroNumSignSpec, err
	}

	zeroNumSignSpec,
		err = nStrNumSymbolsGroup.zeroNumberSign.CopyOut(
		ePrefix.XCpy(
			"zeroNumSignSpec<-"))

	return zeroNumSignSpec, err
}

// IsNOP
//
//	'NOP' stands for 'No Operation'. This method signals
//	whether the three Number Symbol Specifications
//	contained in the current instance of
//	NumStrNumberSymbolGroup are all valid, engaged and fully
//	operational with respect to Number String Formatting
//	operations.
//
//	Type NumStrNumberSymbolGroup encapsulates the Negative,
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
//		all Number Symbol Specifications contained in
//		the current instance of NumStrNumberSymbolGroup are
//		simply empty placeholders and perform no
//		active role in, and are completely ignored by,
//		Number String Formatting algorithms. This means,
//		no Number Symbols will be inserted or formatted
//		as part of Number String Formatting operations.
//
//		If this method returns 'false', it signals that
//		all Number Sign Symbol Specifications are fully
//		populated, valid and functional. Number String
//		Formatting operations will therefore include all
//		specified Number Symbols in formatted number strings.
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) IsNOP() bool {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	if nStrNumSymbolsGroup.negativeNumberSign.IsNOP() &&
		nStrNumSymbolsGroup.positiveNumberSign.IsNOP() &&
		nStrNumSymbolsGroup.zeroNumberSign.IsNOP() {

		return true
	}

	return false
}

//	IsNOPCurrencySymbols
//
//	'NOP' stands for 'No Operation'.
//
//	This method returns a boolean value signaling whether
//	the Currency Symbol Specification contained in the
//	current NumStrNumberSymbolGroup instance is engaged,
//	valid and fully operational with respect to Number
//	String Formatting operations.
//
//	If this method returns 'true', it signals that the
//	Currency Symbol Specification is simply an empty
//	placeholder and performs no active role in, and is
//	completely ignored by, Number String Formatting
//	algorithms.
//
//	When 'NOP' is 'true', no Currency Symbols will be
//	inserted or formatted as part of a Number String
//	formatting operation.
//
//	If this method returns 'false', it signals that the
//	Currency Symbol Specification is fully populated,
//	valid and functional. When 'NOP' is 'false', Number
//	String Formatting operations WILL INCLUDE Currency
//	Symbols in formatted number strings.
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
//		the Currency Symbol Specification is simply an
//		empty placeholder and performs no active role in,
//		and is completely ignored by, Number String
//		Formatting algorithms. This means, no Negative
//		Number Symbols will be inserted or formatted as
//		part of a Number String formatting operation.
//
//		If this method returns 'false', it signals that
//		the Currency Symbol Specification is fully
//		populated, valid and functional. In this case,
//		Number String Formatting WILL INCLUDE Currency
//		Symbols in formatted number strings.
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) IsNOPCurrencySymbols() bool {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	return nStrNumSymbolsGroup.currencySymbol.IsNOP()
}

//	IsNOPNegativeNumSymbols
//
//	'NOP' stands for 'No Operation'.
//
//	This method returns a boolean value signaling whether
//	the Negative Number Sign Symbol Specification
//	contained in the current NumStrNumberSymbolGroup instance
//	is engaged, valid and fully operational with respect
//	to Number String Formatting.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) IsNOPNegativeNumSymbols() bool {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	return nStrNumSymbolsGroup.negativeNumberSign.IsNOP()
}

//	IsNOPPositiveNumSymbols
//
//	'NOP' stands for 'No Operation'. This method
//	returns a boolean value signaling whether the
//	Positive Number Sign Symbol Specification contained
//	in the current NumStrNumberSymbolGroup instance is
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) IsNOPPositiveNumSymbols() bool {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	return nStrNumSymbolsGroup.positiveNumberSign.IsNOP()
}

//	IsNOPZeroNumSymbols
//
//	'NOP' stands for 'No Operation'. This method
//	returns a boolean value signaling whether the
//	Zero Number Sign Symbol Specification contained
//	in the current NumStrNumberSymbolGroup instance is
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) IsNOPZeroNumSymbols() bool {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	return nStrNumSymbolsGroup.zeroNumberSign.IsNOP()
}

//	NewNOP Creates and returns a new instance of
//	NumStrNumberSymbolGroup configured as a NOP.
//
//	'NOP' stands for 'No Operation'.
//
//	Type NumStrNumberSymbolGroup encapsulates the Negative,
//	Positive and Zero Number Symbol Specifications.
//
//	Configured as a NOP, the returned instance of
//	NumStrNumberSymbolGroup will contain three invalid
//	Number Symbol Specifications which are simple empty
//	placeholders. As such these Number Symbol
//	Specifications perform no active role in, and are
//	completely ignored by, Number String Formatting
//	algorithms.
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
//	NumStrNumberSymbolGroup
//
//		This method returns a new instance of
//		NumStrNumberSymbolGroup configured as a NOP.
//
//		'NOP' stands for 'No Operation'.
//
//		This means that	all Number Symbol Specifications
//		contained in the returned instance of
//		NumStrNumberSymbolGroup are	empty placeholders and
//		perform no active role in, and are completely
//		ignored by, Number String Formatting algorithms.
//
//		Therefore, none of the Number Symbols contained
//		in this returned NumStrNumberSymbolGroup instance
//		will be inserted or formatted as part of Number
//		String Formatting operations.
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewNOP() NumStrNumberSymbolGroup {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	newNStrNumSym := NumStrNumberSymbolGroup{}

	newNStrNumSym.negativeNumberSign.SetNOP()

	newNStrNumSym.positiveNumberSign.SetNOP()

	newNStrNumSym.zeroNumberSign.SetNOP()

	return newNStrNumSym
}

//	NewCurrencySimple
//
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup configured for currency
//	Number Symbol formatting.
//
//	If currency number symbol formatting is NOT
//	required, see method:
//
//		NumStrNumberSymbolGroup.NewSignedSimpleNumber()
//
//	Type NumStrNumberSymbolGroup is used to configure
//	Number Symbols required in converting numeric
//	values to formatted Number Strings.
//
//	NumStrNumberSymbolGroup contains three instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values and zero numeric values.
//
//	This method provides a simplified means of creating
//	type NumStrNumberSymbolGroup using default values. The
//	generated returned instance of NumStrNumberSymbolGroup
//	will be configured with currency number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	symbol creation, use one of the more advanced
//	constructor or 'New' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrNumberSymbolGroup.NewSymbolsRunes()
//		NumStrNumberSymbolGroup.NewSymbolsStrings()
//		NumStrNumberSymbolGroup.NewSymbolsSpecs()
//
// ----------------------------------------------------------------
//
// # Currency Defaults
//
//	Currency-Negative Symbol Position:
//		Currency Symbol defaults to 'outside' the
//		minus sign.
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Negative Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Positive Number Symbol:
//		No Positive Number Sign Symbol. Positive
//		values are assumed.
//
//		Positive Numeric Value Currency Examples:
//			European Number String: "123.456 €"
//			US Number String: "$ 123.456"
//
//	Zero Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//		Zero Numeric Value Currency Examples:
//			European Number String: "0.00 €"
//			US Number String: "$ 0.00"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	currencySymbols				string
//
//		The symbol or symbols used to format currency. This
//		currency formatting will be configured in the new
//		instance of NumStrNumberSymbolGroup returned by this
//		method.
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Number Symbols
//		on the left side of the numeric value. Such
//		Number Symbols are therefore configured as
//		leading Number Symbols. This is the positioning
//		format used in the US, UK, Australia and most of
//		Canada.
//
//		Example Number Strings:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted between
//				the currency symbol and the minus sign.
//
//		When set to 'false', the returned instance of
//		NumStrNumberSymbolGroup will configure Number Symbols
//		on the right side of the numeric value. Such
//		Number Symbols are therefore configured as
//		trailing Number Symbols. This is the positioning
//		format used in France, Germany and many countries
//		in the European Union.
//
//		Example Number Strings:
//			"123.456- €"
//
//		NOTE:	A space is automatically inserted between
//				the minus sign and the currency symbol.
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
//	NumStrNumberSymbolGroup
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with the Positive, Negative and Zero Number
//		Sign Symbol Specification objects.
//
//		This returned NumStrNumberSymbolGroup instance will
//		be configured with currency symbols for Number
//		String formatting.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSimpleCurrency(
	currencySymbols string,
	leadingNumSymbols bool,
	errorPrefix interface{}) (
	NumStrNumberSymbolGroup,
	error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNStrNumSymbols NumStrNumberSymbolGroup

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"NewCurrencySimple()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).setSimpleNumSymbolsConfig(
		&newNStrNumSymbols,
		currencySymbols,
		leadingNumSymbols,
		ePrefix.XCpy(
			"newNStrNumSymbols<-"))

	return newNStrNumSymbols, err
}

//	NewSignedSimpleNumber
//
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup configured for Signed
//	Number formatting. Signed numbers do NOT
//	contain currency symbols.
//
//	If currency number symbol formatting IS
//	required, see method:
//
//		NumStrNumberSymbolGroup.NewCurrencySimple()
//
//	Type NumStrNumberSymbolGroup is used to configure
//	Number Symbols required in converting numeric
//	values to formatted Number Strings.
//
//	NumStrNumberSymbolGroup contains three instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values and zero numeric values.
//
//	This method provides a simplified means of creating
//	type NumStrNumberSymbolGroup using default values and
//	generating number symbols suitable for signed numeric
//	values. Signed Number Symbols do NOT contain currency
//	symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over signed number symbol
//	creation, use one of the more advanced constructor or
//	'New' methods to achieve specialized multinational or
//	multicultural signed number symbol formatting
//	requirements:
//
//		NumStrNumberSymbolGroup.NewSymbolsRunes()
//		NumStrNumberSymbolGroup.NewSymbolsStrings()
//		NumStrNumberSymbolGroup.NewSymbolsSpecs()
//
// ----------------------------------------------------------------
//
// # Signed Number Defaults
//
//	Negative Signed Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456-"
//			US Number String: "-123.456"
//
//	Positive Signed Number Symbol:
//		No Positive Number Sign Symbol. Positive values
//		are assumed.
//
//			Positive Value Number String: "123.456"
//
//	Zero Signed Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//			Zero Value Number String: "123.456"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Number Symbols
//		on the left side of the numeric value. Such
//		Number Symbols are therefore configured as
//		leading Number Symbols. This is the positioning
//		format used in the US, UK, Australia and most of
//		Canada.
//
//		Example Number String with Leading Number Symbols:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted between
//				the currency symbol and the minus sign.
//
//		When set to 'false', the returned instance of
//		NumStrNumberSymbolGroup will configure Number Symbols
//		on the right side of the numeric value. Such
//		Number Symbols are therefore configured as
//		trailing Number Symbols. This is the positioning
//		format used in France, Germany and many countries
//		in the European Union.
//
//		Example Number String with Trailing Number Symbols:
//			"123.456- €"
//
//		NOTE:	A space is automatically inserted between
//				the minus sign and the currency symbol.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with the Positive, Negative and Zero Number
//		Sign Symbol Specification objects.
//
//		This returned NumStrNumberSymbolGroup instance will
//		be configured with symbols suitable for Signed
//		Number String formatting.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSimpleSignedNumber(
	leadingNumSymbols bool,
	errorPrefix interface{}) (
	NumStrNumberSymbolGroup,
	error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNStrNumSymbols NumStrNumberSymbolGroup

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"NewSimpleSignedNumber()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).setSimpleNumSymbolsConfig(
		&newNStrNumSymbols,
		"",
		leadingNumSymbols,
		ePrefix.XCpy(
			"newNStrNumSymbols<-"))

	return newNStrNumSymbols, err
}

//	NewSymbolsRunes
//
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup.
//
//	This type is used to configure Number Symbols
//	required in converting numeric values to Number
//	Strings.
//
//	Type NumStrNumberSymbolGroup contains three instances of
//	NumStrNumberSymbolSpec defining the Number Symbols to
//	be used with positive numeric values, negative numeric
//	values and zero numeric values.
//
//	This method generates a new instance of
//	NumStrNumberSymbolGroup using rune arrays as input
//	parameters.
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSymbolsRunes(
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
	NumStrNumberSymbolGroup,
	error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumberSymbols NumStrNumberSymbolGroup

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"NewSymbolsRunes()",
		"")

	if err != nil {
		return newNumberSymbols, err
	}

	nStrNumSymNanobot := numStrNumberSymbolGroupNanobot{}

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
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup.
//
//	This type is used to configure Number Symbols
//	required in converting numeric values to Number
//	Strings.
//
//	Type NumStrNumberSymbolGroup contains three instances of
//	NumStrNumberSymbolSpec defining the Number Symbols to
//	be used with positive numeric values, negative numeric
//	values and zero numeric values.
//
//	This method generates a new instance of
//	NumStrNumberSymbolGroup using strings as input parameters.
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSymbolsStrings(
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
	NumStrNumberSymbolGroup,
	error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumberSymbols NumStrNumberSymbolGroup

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"NewSymbolsStrings()",
		"")

	if err != nil {
		return newNumberSymbols, err
	}

	nStrNumSymNanobot := numStrNumberSymbolGroupNanobot{}

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
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup.
//
//	This type is used to configure Number Symbols
//	required in converting numeric values to Number
//	Strings.
//
//	Type NumStrNumberSymbolGroup contains three instances of
//	NumStrNumberSymbolSpec defining the Number Symbols to
//	be used with positive numeric values, negative numeric
//	values and zero numeric values.
//
//	This method generates a new instance of
//	NumStrNumberSymbolGroup using Positive, Negative and Zero
//	Number Sign Symbol Specification objects passed as
//	input parameters.
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
//		instance of NumStrNumberSymbolGroup.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Negative Number Sign Symbol Specification
//		will be copied to the corresponding Negative
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbolGroup.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbolGroup.
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSymbolsSpecs(
	positiveNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) (
	NumStrNumberSymbolGroup,
	error) {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumberSymbols NumStrNumberSymbolGroup

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"NewSymbolsSpecs()",
		"")

	if err != nil {
		return newNumberSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).
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
//	NumStrNumberSymbolGroup based on the Negative Number
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
//	NumStrNumberSymbolGroup.
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
//		NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNegativeNumSignSpec(
	negativeNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetNegativeNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).
		setNegativeNumSignSpec(
			nStrNumSymbolsGroup,
			negativeNumberSign,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup"))

}

//	SetNegativeSymbolsRunes
//
//	Deletes and resets the Negative Number Symbols data
//	fields for the current instance of NumStrNumberSymbolGroup.
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
//	the current instance of NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.negativeNumberSign
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNegativeSymbolsRunes(
	leadingNegativeNumberSymbols []rune,
	trailingNegativeNumberSymbols []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetNegativeSymbolsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).setNegativeNumSignRunes(
		nStrNumSymbolsGroup,
		leadingNegativeNumberSymbols,
		trailingNegativeNumberSymbols,
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))
}

//	SetNegativeSymbolsStrings
//
//	Deletes and resets the Negative Number Symbols data
//	fields for the current instance of NumStrNumberSymbolGroup.
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
//	the current instance of NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.negativeNumberSign
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNegativeSymbolsStrings(
	leadingNegativeNumberSymbols string,
	trailingNegativeNumberSymbols string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetNegativeSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).setNegativeNumSignRunes(
		nStrNumSymbolsGroup,
		[]rune(leadingNegativeNumberSymbols),
		[]rune(trailingNegativeNumberSymbols),
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))
}

//	SetNOP
//
//	'NOP' stands for 'No Operation'.
//
//	When called, this method will convert the current
//	instance of NumStrNumberSymbolGroup to NOP or 'No
//	Operation'. This means that all data values
//	contained in the current NumStrNumberSymbolGroup
//	instance will be deleted and reset to their
//	zero or uninitialized state.
//
//	Type NumStrNumberSymbolGroup encapsulates the Negative,
//	Positive and Zero Number Symbol Specifications.
//
//	Calling this method will convert all three Number
//	Symbol Specifications to empty placeholders and
//	ensure that they perform no active role in, and are
//	completely ignored by, Number String Formatting
//	algorithms. In this state, none of the three Number
//	Symbol Specifications are valid and all are used as
//	empty placeholders when formatting Number Strings.
//
//	To determine the status of NOP for the current
//	instance of NumStrNumberSymbolGroup, use the following
//	methods:
//
//		NumStrNumberSymbolGroup.IsNOP()
//		NumStrNumberSymbolGroup.IsNOPNegativeNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPPositiveNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPZeroNumSymbols()
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNOP() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.negativeNumberSign.SetNOP()

	nStrNumSymbolsGroup.positiveNumberSign.SetNOP()

	nStrNumSymbolsGroup.zeroNumberSign.SetNOP()

	nStrNumSymbolsGroup.currencySymbol.SetNOP()

	return
}

//	SetNOPCurrencySymbol
//
//	'NOP' stands for 'No Operation'.
//
//	When called, this method will convert the Currency
//	Symbol Specification in the current instance of
//	NumStrNumberSymbolGroup to a NOP or 'No Operation'.
//	This means that the Currency Symbol Specification is
//	simply an empty placeholder and performs no active
//	role in, and is completely ignored by, Number String
//	Formatting algorithms.
//
//	After conversion to NOP, no Currency Symbols will be
//	inserted or formatted as part of a Number String
//	formatting operation.
//
//	To determine the status of NOP for the current
//	instance of NumStrNumberSymbolGroup, use the following
//	methods:
//
//		NumStrNumberSymbolGroup.IsNOP()
//		NumStrNumberSymbolGroup.IsNOPNegativeNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPPositiveNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPZeroNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPCurrencySymbol()
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNOPCurrencySymbol() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.currencySymbol.SetNOP()

	return
}

//	SetNOPNegativeNumSymbol
//
//	'NOP' stands for 'No Operation'.
//
//	When called, this method will convert the Negative
//	Number Symbol Specification in the current instance
//	of NumStrNumberSymbolGroup to a NOP or
//	'No Operation'. This means that the Negative Number
//	Symbol Specification is simply an empty placeholder
//	and performs no active role in, and is completely
//	ignored by, Number String Formatting algorithms.
//
//	After conversion to NOP, no Negative Number Symbols
//	will be inserted or formatted as part of a Number
//	String formatting operation.
//
//	To determine the status of NOP for the current
//	instance of NumStrNumberSymbolGroup, use the following
//	methods:
//
//		NumStrNumberSymbolGroup.IsNOP()
//		NumStrNumberSymbolGroup.IsNOPNegativeNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPPositiveNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPZeroNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPCurrencySymbol()
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNOPNegativeNumSymbol() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.negativeNumberSign.SetNOP()

	return
}

//	SetNOPPositiveNumSymbol
//
//	'NOP' stands for 'No Operation'.
//
//	When called, this method will convert the Positive
//	Number Symbol Specification in the current instance
//	of NumStrNumberSymbolGroup to a NOP or
//	'No Operation'. This means that the Positive Number
//	Symbol Specification is simply an empty placeholder
//	and performs no active role in, and is completely
//	ignored by, Number String Formatting algorithms.
//
//	After conversion to NOP, no Positive Number Symbols
//	will be inserted or formatted as part of a Number
//	String formatting operation.
//
//	To determine the status of NOP for the current
//	instance of NumStrNumberSymbolGroup, use the
//	following methods:
//
//		NumStrNumberSymbolGroup.IsNOP()
//		NumStrNumberSymbolGroup.IsNOPNegativeNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPPositiveNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPZeroNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPCurrencySymbol()
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNOPPositiveNumSymbol() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.positiveNumberSign.SetNOP()

	return
}

//	SetNOPZeroNumSymbol
//
//	'NOP' stands for 'No Operation'.
//
//	When called, this method will convert the Zero
//	Number Symbol Specification in the current instance
//	of NumStrNumberSymbolGroup to a NOP or
//	'No Operation'. This means that the Zero Number
//	Symbol Specification is simply an empty placeholder
//	and performs no active role in, and is completely
//	ignored by, Number String Formatting algorithms.
//
//	After conversion to NOP, no Zero Number Symbols will
//	be inserted or formatted as part of a Number String
//	formatting operation.
//
//	To determine the status of NOP for the current
//	instance of NumStrNumberSymbolGroup, use the
//	following methods:
//
//		NumStrNumberSymbolGroup.IsNOP()
//		NumStrNumberSymbolGroup.IsNOPNegativeNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPPositiveNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPZeroNumSymbols()
//		NumStrNumberSymbolGroup.IsNOPCurrencySymbol()
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNOPZeroNumSymbol() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.zeroNumberSign.SetNOP()

	return
}

//	SetPositiveNumSignSpec
//
//	Reconfigures the current instance of
//	NumStrNumberSymbolGroup based on the Positive Number
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
//	NumStrNumberSymbolGroup.
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
//		NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetPositiveNumSignSpec(
	positiveNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetPositiveNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).
		setPositiveNumSignSpec(
			nStrNumSymbolsGroup,
			positiveNumberSign,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup"))

}

//	SetPositiveSymbolsRunes
//
//	Deletes and resets the Positive Number Symbols data
//	fields for the current instance of NumStrNumberSymbolGroup.
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
//	the current instance of NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.positiveNumberSign
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetPositiveSymbolsRunes(
	leadingPositiveNumberSymbols []rune,
	trailingPositiveNumberSymbols []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetPositiveSymbolsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).setPositiveNumSignRunes(
		nStrNumSymbolsGroup,
		leadingPositiveNumberSymbols,
		trailingPositiveNumberSymbols,
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))
}

//	SetPositiveSymbolsStrings
//
//	Deletes and resets the Positive Number Symbols data
//	fields for the current instance of NumStrNumberSymbolGroup.
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
//	the current instance of NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.positiveNumberSign
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetPositiveSymbolsStrings(
	leadingPositiveNumberSymbols string,
	trailingPositiveNumberSymbols string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetPositiveSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).setPositiveNumSignRunes(
		nStrNumSymbolsGroup,
		[]rune(leadingPositiveNumberSymbols),
		[]rune(trailingPositiveNumberSymbols),
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))
}

// SetSimpleCurrency
//
// Deletes and reconfigures the data values contained in
// the current instance of NumStrNumberSymbolGroup with
// currency number symbol parameters.
//
// If currency number symbol formatting is NOT required,
// see method:
//
//	NumStrNumberSymbolGroup.SetSignedSimpleNumber()
//
// Type NumStrNumberSymbolGroup is used to configure
// Number Symbols required in converting numeric
// values to formatted Number Strings.
//
// NumStrNumberSymbolGroup contains three instances of
// type NumStrNumberSymbolSpec defining the Number
// Symbols to be used with positive numeric values,
// negative numeric values and zero numeric values.
//
// This method provides a simplified means of
// reconfiguring the current instance of
// NumStrNumberSymbolGroup with currency number symbols
// using default values.
//
// Upon completion, this method will reconfigure the
// current instance of NumStrNumberSymbolGroup with
// currency number symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	symbol creation, use one of the more advanced
//	'Setter' methods to achieve specialized
//	multinational or multicultural currency number
//	symbol formatting requirements:
//
//		NumStrNumberSymbolGroup.SetSymbolsRunes()
//		NumStrNumberSymbolGroup.SetSymbolsStrings()
//		NumStrNumberSymbolGroup.SetSymbolsSpecs()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and
//	reconfigure all data values contained in the current
//	instance of NumStrNumberSymbolGroup.
//
// ----------------------------------------------------------------
//
// # Currency Defaults
//
//	Currency-Negative Symbol Position:
//		Currency Symbol defaults to 'outside' the
//		minus sign.
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Negative Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Positive Number Symbol:
//		No Positive Number Sign Symbol. Positive
//		values are assumed.
//
//		Positive Numeric Value Currency Examples:
//			European Number String: "123.456 €"
//			US Number String: "$ 123.456"
//
//	Zero Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//		Zero Numeric Value Currency Examples:
//			European Number String: "0.00 €"
//			US Number String: "$ 0.00"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	currencySymbols				string
//
//		The symbol or symbols used to format currency. The
//		current instance of NumStrNumberSymbolGroup will be
//		reconfigured with the currency symbols contained
//		in this input parameter.
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', the current instance of
//		NumStrNumberSymbolGroup will configure Number Symbols
//		on the left side of the numeric value. Such
//		Number Symbols are therefore configured as
//		leading Number Symbols. This is the positioning
//		format used in the US, UK, Australia and most of
//		Canada.
//
//		Example Number Strings:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted between
//				the currency symbol and the minus sign.
//
//		When 'leadingNumSymbols' is set to 'false', the
//		current instance of NumStrNumberSymbolGroup will
//		configure Number Symbols on the right side of the
//		numeric value. Such Number Symbols are therefore
//		configured as trailing Number Symbols. This is
//		the positioning format used in France, Germany
//		and many countries in the European Union.
//
//		Example Number Strings:
//			"123.456- €"
//
//		NOTE:	A space is automatically inserted between
//				the minus sign and the currency symbol.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSimpleCurrency(
	currencySymbols string,
	leadingNumSymbols bool,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetSimpleCurrency()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).setSimpleNumSymbolsConfig(
		nStrNumSymbolsGroup,
		currencySymbols,
		leadingNumSymbols,
		ePrefix.XCpy(
			"nStrNumSymbolsGroup<-"))
}

//	SetSignedSimpleNumber
//
//	Deletes and reconfigures the data values contained in
//	the current instance of NumStrNumberSymbolGroup with
//	signed number symbol parameters. Signed numbers do
//	NOT contain currency symbols.
//
//	If currency number symbol formatting IS required,
//	see method:
//
//	NumStrNumberSymbolGroup.SetSimpleCurrency()
//
//	Type NumStrNumberSymbolGroup is used to configure Number
//	Symbols required in converting numeric values to
//	formatted Number Strings.
//
//	NumStrNumberSymbolGroup contains three instances of type
//	NumStrNumberSymbolSpec defining the Number Symbols to
//	be used with positive numeric values, negative
//	numeric values and zero numeric values.
//
//	This method provides a simplified means for
//	reconfiguring the current instance of
//	NumStrNumberSymbolGroup using default values and
//	generating number symbols suitable for signed numeric
//	values. Signed Number Symbols do NOT contain currency
//	symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over signed number symbol
//	creation, use one of the more advanced 'Setter'
//	methods to achieve specialized multinational or
//	multicultural signed number symbol formatting
//	requirements:
//
//		NumStrNumberSymbolGroup.SetSymbolsRunes()
//		NumStrNumberSymbolGroup.SetSymbolsStrings()
//		NumStrNumberSymbolGroup.SetSymbolsSpecs()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and
//	reconfigure all data values contained in the current
//	instance of NumStrNumberSymbolGroup.
//
// ----------------------------------------------------------------
//
// # Signed Number Defaults
//
//	Negative Signed Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456-"
//			US Number String: "-123.456"
//
//	Positive Signed Number Symbol:
//		No Positive Number Sign Symbol. Positive values
//		are assumed.
//
//			Positive Value Number String: "123.456"
//
//	Zero Signed Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//			Zero Value Number String: "123.456"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', the current instance of
//		NumStrNumberSymbolGroup will configure Number Symbols
//		on the left side of the numeric value. Such
//		Number Symbols are therefore configured as
//		leading Number Symbols. This is the positioning
//		format used in the US, UK, Australia and most of
//		Canada.
//
//		Example Number Strings:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted between
//				the currency symbol and the minus sign.
//
//		When 'leadingNumSymbols' is set to 'false', the
//		current instance of NumStrNumberSymbolGroup will
//		configure Number Symbols on the right side of the
//		numeric value. Such Number Symbols are therefore
//		configured as trailing Number Symbols. This is
//		the positioning format used in France, Germany
//		and many countries in the European Union.
//
//		Example Number Strings:
//			"123.456- €"
//
//		NOTE:	A space is automatically inserted between
//				the minus sign and the currency symbol.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSimpleSignedNumber(
	leadingNumSymbols bool,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetSignedSimpleNumber()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).setSimpleNumSymbolsConfig(
		nStrNumSymbolsGroup,
		"",
		leadingNumSymbols,
		ePrefix.XCpy(
			"nStrNumSymbolsGroup<-"))
}

//	SetSymbolsRunes
//
//	Reconfigures the current instance of
//	NumStrNumberSymbolGroup based on the Positive, Negative
//	and Zero Number	Sign Symbol rune arrays passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbolGroup will be
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSymbolsRunes(
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

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetSymbolsRunes()",
		"")

	if err != nil {
		return err
	}

	nStrNumSymNanobot := numStrNumberSymbolGroupNanobot{}

	err = nStrNumSymNanobot.setPositiveNumSignRunes(
		nStrNumSymbolsGroup,
		leadingPositiveNumberSymbols,
		trailingPositiveNumberSymbols,
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-PositiveNumSyms"))

	if err != nil {
		return err
	}

	err = nStrNumSymNanobot.setNegativeNumSignRunes(
		nStrNumSymbolsGroup,
		leadingNegativeNumberSymbols,
		trailingNegativeNumberSymbols,
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-NegativeNumSyms"))

	if err != nil {
		return err
	}

	return nStrNumSymNanobot.setZeroNumSignRunes(
		nStrNumSymbolsGroup,
		leadingZeroNumberSymbols,
		trailingZeroNumberSymbols,
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))

}

//	SetSymbolsSpecs
//
//	Reconfigures the current instance of
//	NumStrNumberSymbolGroup based on the Positive, Negative
//	and Zero Number	Sign Symbol Specification objects
//	passed as input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbolGroup will be
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
//		instance of NumStrNumberSymbolGroup.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Negative Number Sign Symbol Specification
//		will be copied to the corresponding Negative
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbolGroup.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Symbol Specification in the new, returned
//		instance of NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSymbolsSpecs(
	positiveNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetSymbolsSpecs()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setNumSymbolSpecs(
			nStrNumSymbolsGroup,
			positiveNumberSign,
			negativeNumberSign,
			zeroNumberSign,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup"))

}

//	SetSymbolsStrings
//
//	Reconfigures the current instance of
//	NumStrNumberSymbolGroup based on the Positive, Negative
//	and Zero Number	Sign Symbol strings passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbolGroup will be
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSymbolsStrings(
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

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newNumberSymbols NumStrNumberSymbolGroup

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	nStrNumSymNanobot := numStrNumberSymbolGroupNanobot{}

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
//	fields for the current instance of NumStrNumberSymbolGroup.
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
//	the current instance of NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.zeroNumberSign
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetZeroSymbolsRunes(
	leadingZeroNumberSymbols []rune,
	trailingZeroNumberSymbols []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetZeroSymbolsRunes()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).setZeroNumSignRunes(
		nStrNumSymbolsGroup,
		leadingZeroNumberSymbols,
		trailingZeroNumberSymbols,
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))
}

//	SetZeroNumSignSpec
//
//	Reconfigures the current instance of
//	NumStrNumberSymbolGroup based on the Zero Number
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
//	NumStrNumberSymbolGroup.
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
//		NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetZeroNumSignSpec(
	positiveNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetZeroNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).
		setZeroNumSignSpec(
			nStrNumSymbolsGroup,
			positiveNumberSign,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup"))

}

//	SetZeroSymbolsStrings
//
//	Deletes and resets the Zero Number Symbols data
//	fields for the current instance of NumStrNumberSymbolGroup.
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
//	the current instance of NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.zeroNumberSign
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetZeroSymbolsStrings(
	leadingZeroNumberSymbols string,
	trailingZeroNumberSymbols string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
	errorPrefix interface{}) error {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberSymbolGroup."+
			"SetZeroSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).setZeroNumSignRunes(
		nStrNumSymbolsGroup,
		[]rune(leadingZeroNumberSymbols),
		[]rune(trailingZeroNumberSymbols),
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"nuStrNumSym<-ZeroNumSyms"))
}

// numStrNumberSymbolGroupMechanics
//
// Provides helper methods for NumStrNumberSymbolGroup.
type numStrNumberSymbolGroupMechanics struct {
	lock *sync.Mutex
}

//	copyNumSymbols
//
//	Copies all Number Symbol Specification data
//	from input parameter 'sourceNumSymbols' to
//	input parameter 'destinationNumSymbols'.
//	Both instances are of type
//	NumStrNumberSymbolGroup.
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
//	destinationNumSymbols			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		All the member variable data fields in this object will be
//		replaced by data values copied from input parameter
//		'sourceNumSymbolSpec'.
//
//		'destinationNumSymbolSpec' is the destination for this
//		copy operation.
//
//	sourceNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to another instance of NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) copyNumSymbols(
	destinationNumSymbols *NumStrNumberSymbolGroup,
	sourceNumSymbols *NumStrNumberSymbolGroup,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolsGroupMech.lock == nil {
		nStrNumSymbolsGroupMech.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroupMech.lock.Lock()

	defer nStrNumSymbolsGroupMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupMechanics."+
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

	new(numStrNumberSymbolGroupNanobot).empty(destinationNumSymbols)

	if sourceNumSymbols.IsNOP() {
		// Nothing to do. All data values
		// are set to their zero states.
		return err
	}

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

	if err != nil {
		return err
	}

	err = destinationNumSymbols.currencySymbol.CopyIn(
		&sourceNumSymbols.currencySymbol,
		ePrefix.XCpy(
			"destinationNumSymbols<-"+
				"sourceNumSymbols.currencySymbol"))

	return err
}

//	setNumSymbolSpecs
//
//	Receives three NumStrNumberSymbolSpec objects and
//	proceeds to reset the corresponding member variable
//	data values for the NumStrNumberSymbolGroup input
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
//	nStrNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setNumSymbolSpecs(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	positiveNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolsGroupMech.lock == nil {
		nStrNumSymbolsGroupMech.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroupMech.lock.Lock()

	defer nStrNumSymbolsGroupMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupMechanics."+
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

//	setSimpleNumSymbolsConfig
//
//	Receives a pointer to an instance of
//	NumStrNumberSymbolGroup and proceeds to reconfigure that
//	instance for currency number symbols or signed number
//	symbols.
//
//	Type NumStrNumberSymbolGroup is used to configure
//	Number Symbols required in converting numeric
//	values to formatted Number Strings.
//
//	NumStrNumberSymbolGroup contains three instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values and zero numeric values.
//
//	This method provides a simplified means of
//	configuring type NumStrNumberSymbolGroup using default
//	values. The generated returned instance of
//	NumStrNumberSymbolGroup will be configured for either
//	Currency Number Symbols or Signed Number Symbols,
//	depending on the input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and
//	reconfigure all data values contained in input
//	parameter, 'nStrNumSymbols'.
//
// ----------------------------------------------------------------
//
// # Currency Defaults
//
//	Currency-Negative Symbol Position:
//		Currency Symbol defaults to 'outside' the
//		minus sign.
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Negative Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Positive Number Symbol:
//		No Positive Number Sign Symbol. Positive
//		values are assumed.
//
//		Positive Numeric Value Currency Examples:
//			European Number String: "123.456 €"
//			US Number String: "$ 123.456"
//
//	Zero Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//		Zero Numeric Value Currency Examples:
//			European Number String: "0.00 €"
//			US Number String: "$ 0.00"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
// # Signed Number Defaults
//
//	Negative Signed Number Symbol:
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456-"
//			US Number String: "-123.456"
//
//	Positive Signed Number Symbol:
//		No Positive Number Sign Symbol. Positive values
//		are assumed.
//
//			Positive Value Number String: "123.456"
//
//	Zero Signed Number Symbol:
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative.
//
//			Zero Value Number String: "123.456"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:       01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	nStrNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured as either
//		new currency symbol specifications or new signed
//		number specifications.
//
//	currencySymbols				string
//
//		The symbol or symbols used to format currency.
//		This currency formatting will be used to
//		reconfigure the NumStrNumberSymbolGroup instance
//		passed by input parameter, 'nStrNumSymbols'.
//
//		If this string is empty, 'nStrNumSymbols' will be
//		reconfigured with signed number symbols suitable
//		for a number string containing a signed numeric
//		value.
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', 'nStrNumSymbols' will be
//		reconfigured with Number Symbols on the left side
//		of the numeric value. Such Number Symbols are
//		therefore configured as leading	Number Symbols.
//		This is the positioning format used in the US,
//		UK, Australia and most of Canada.
//
//		Example Number Strings:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted between
//				the currency symbol and the minus sign.
//
//		When set to 'false', the returned instance of
//		NumStrNumberSymbolGroup will configure Number Symbols
//		on the right side of the numeric value. Such
//		Number Symbols are therefore configured as
//		trailing Number Symbols. This is the positioning
//		format used in France, Germany and many countries
//		in the European Union.
//
//		Example Number Strings:
//			"123.456- €"
//
//		NOTE:	A space is automatically inserted between
//				the minus sign and the currency symbol.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setSimpleNumSymbolsConfig(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	currencySymbols string,
	leadingNumSymbols bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymbolsGroupMech.lock == nil {
		nStrNumSymbolsGroupMech.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroupMech.lock.Lock()

	defer nStrNumSymbolsGroupMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolGroupMechanics."+
			"setSimpleNumSymbolsConfig()",
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

	new(numStrNumberSymbolGroupNanobot).empty(
		nStrNumSymbols)

	var numSymStr string

	if leadingNumSymbols {
		// Leading Number Symbols

		if len(currencySymbols) == 0 {

			numSymStr = ""

		} else {

			numSymStr = currencySymbols + " "
		}

		err = nStrNumSymbols.SetPositiveSymbolsRunes(
			[]rune(numSymStr),
			nil,
			NumFieldSymPos.InsideNumField(),
			nil)

		if err != nil {
			return err
		}

		err = nStrNumSymbols.SetZeroSymbolsRunes(
			[]rune(numSymStr),
			nil,
			NumFieldSymPos.InsideNumField(),
			nil)

		if err != nil {
			return err
		}

		if len(currencySymbols) == 0 {

			numSymStr = "-"

		} else {
			numSymStr = currencySymbols + " -"
		}

		err = nStrNumSymbols.SetNegativeSymbolsRunes(
			[]rune(numSymStr),
			nil,
			NumFieldSymPos.InsideNumField(),
			nil)

		if err != nil {
			return err
		}

	} else {
		// Trailing Number Symbols

		if len(currencySymbols) == 0 {

			numSymStr = ""

		} else {

			numSymStr = " " + currencySymbols
		}

		err = nStrNumSymbols.SetPositiveSymbolsRunes(
			nil,
			[]rune(numSymStr),
			NumFieldSymPos.InsideNumField(),
			nil)

		if err != nil {
			return err
		}

		err = nStrNumSymbols.SetZeroSymbolsRunes(
			nil,
			[]rune(numSymStr),
			NumFieldSymPos.InsideNumField(),
			nil)

		if err != nil {
			return err
		}

		if len(currencySymbols) == 0 {

			numSymStr = "-"

		} else {
			numSymStr = "- " + currencySymbols
		}

		err = nStrNumSymbols.SetNegativeSymbolsRunes(
			nil,
			[]rune(numSymStr),
			NumFieldSymPos.InsideNumField(),
			nil)

		if err != nil {
			return err
		}

	}

	return err
}

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

	err = nStrNumSymbols.negativeNumberSign.
		SetNumberSignLeadingTrailingSymbolRunes(
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
//	variable data value for the NumStrNumberSymbolGroup
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

	err = nStrNumSymbols.positiveNumberSign.
		SetNumberSignLeadingTrailingSymbolRunes(
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
//	variable data value for the NumStrNumberSymbolGroup
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

	err = nStrNumSymbols.zeroNumberSign.
		SetNumberSignLeadingTrailingSymbolRunes(
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
//	variable data value for the NumStrNumberSymbolGroup
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
//	nStrNumSymbols				*NumStrNumberSymbolGroup
//
//		A pointer to an instance of NumStrNumberSymbolGroup.
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

	err = nStrNumSymbols.zeroNumberSign.CopyIn(
		&zeroNumberSign,
		ePrefix.XCpy(
			"nStrNumSymbols.zeroNumberSign<-"+
				"zeroNumberSign"))

	return err
}

// numStrNumberSymbolGroupAtom
//
// Provides helper methods for NumStrNumberSymbolGroup.
type numStrNumberSymbolGroupAtom struct {
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
func (nStrNumSymbolGroupAtom *numStrNumberSymbolGroupAtom) emptyNegativeNumSymbols(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupAtom.lock == nil {
		nStrNumSymbolGroupAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupAtom.lock.Lock()

	defer nStrNumSymbolGroupAtom.lock.Unlock()

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
func (nStrNumSymbolGroupAtom *numStrNumberSymbolGroupAtom) emptyPositiveNumSymbols(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupAtom.lock == nil {
		nStrNumSymbolGroupAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupAtom.lock.Lock()

	defer nStrNumSymbolGroupAtom.lock.Unlock()

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
func (nStrNumSymbolGroupAtom *numStrNumberSymbolGroupAtom) emptyZeroNumSymbols(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupAtom.lock == nil {
		nStrNumSymbolGroupAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupAtom.lock.Lock()

	defer nStrNumSymbolGroupAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.zeroNumberSign.Empty()
}

//	emptyCurrencySymbols
//
//	Deletes and resets the Currency Symbol Specification
//	to its zero or uninitialized state.
//
//	The Currency Symbol Specification object is a member
//	variable in the 'nStrNumSymbols' instance passed as
//	an input parameter.
//
//		nStrNumSymbols.currencySymbol
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All data values contained in the Currency Symbol
//	Specification will be deleted and reset to their
//	zero or uninitialized values. This Currency Symbol
//	Specification member variable is identified as:
//
//		nStrNumSymbols.currencySymbol
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
func (nStrNumSymbolGroupAtom *numStrNumberSymbolGroupAtom) emptyCurrencySymbols(
	nStrNumSymbols *NumStrNumberSymbolGroup) {

	if nStrNumSymbolGroupAtom.lock == nil {
		nStrNumSymbolGroupAtom.lock = new(sync.Mutex)
	}

	nStrNumSymbolGroupAtom.lock.Lock()

	defer nStrNumSymbolGroupAtom.lock.Unlock()

	if nStrNumSymbols == nil {

		return
	}

	nStrNumSymbols.currencySymbol.Empty()
}

package strmech

import (
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

//	EmptyCurrencySymbols
//
//	Deletes and resets the Currency Symbol Specification
//	to its zero or uninitialized state.
//
//	The target Currency Symbol Specification object is a
//	member variable contained in the current instance of
//	NumStrNumberSymbolGroup.
//
//	This action will put the Currency Symbol
//	Specification into a NOP or No Operation state. In
//	this state, the Currency Symbol Specification will be
//	treated as an empty placeholder and ignored by Number
//	String Formatting operations.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will modify the current instance of
//	NumStrNumberSymbolGroup.
//
//	The Currency Symbol Specification member variable
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) EmptyCurrencySymbols() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	new(numStrNumberSymbolGroupAtom).emptyCurrencySymbols(
		nStrNumSymbolsGroup)
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

//	GetCurrencySymbolSpec
//
//	Returns a deep copy of the Currency Symbol
//	Specification configured for the current instance of
//	NumStrNumberSymbolGroup.
//
//	The Currency Symbol	Specification is used to define
//	and format currency symbols in number strings.
//
//	Example-1: Leading Currency Symbols
//
//		Leading Symbols: "$"
//		Number String:   "$123.456"
//
//	Example-2: Leading Currency Symbols
//
//		Leading Symbols: "$ "
//		Number String:   "$ 123.456"
//
//	Example-3: Trailing Currency Symbols
//
//		Trailing Symbols: "€"
//		Number String:   "123.456€"
//
//	Example-4: Trailing Currency Symbols
//
//		Trailing Symbols: " €"
//		Number String:   "123.456 €"
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
//	currencySymbolSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the
//		Currency Symbol	Specification contained in the
//		current instance of	NumStrNumberSymbolGroup.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) GetCurrencySymbolSpec(
	errorPrefix interface{}) (
	currencySymbolSpec NumStrNumberSymbolSpec,
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
			"GetCurrencySymbolSpec()",
		"")

	if err != nil {
		return currencySymbolSpec, err
	}

	currencySymbolSpec,
		err = nStrNumSymbolsGroup.negativeNumberSign.CopyOut(
		ePrefix.XCpy(
			"currencySymbolSpec<-"))

	return currencySymbolSpec, err
}

//	GetNegativeNumberSignSpec
//
//	Returns an instance of NumStrNumberSymbolSpec
//	representing deep copy of the Negative Number Sign
//	Specification contained in the current instance of
//	NumStrNumberSymbolGroup.
//
//	The Negative Number Sign Specification is used to
//	configure negative number sign symbols for negative
//	numeric values formatted and displayed in number
//	stings.
//
//	Example-1: Leading Negative Number Sign Symbols
//
//		Leading Symbols: "- "
//		Number String:   "- 123.456"
//
//	Example-2: Leading Negative Number Sign Symbols
//
//		Leading Symbols: "-"
//		Number String:   "-123.456"
//
//	Example-3: Trailing Negative Number Sign Symbols
//
//		Trailing Symbols: " -"
//		Number String:   "123.456 -"
//
//	Example-4: Trailing Negative Number Sign Symbols
//
//		Trailing Symbols: "-"
//		Number String:   "123.456-"
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
//	negativeNumSignSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the Negative
//		Number Sign Symbol Specification configured for
//		the current instance of NumStrNumberSymbolGroup.
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
//	Returns a deep copy of the Positive Number Sign
//	Specification contained in the current instance of
//	NumStrNumberSymbolGroup.
//
//	Positive number signs are commonly implied
//	and not specified. However, the user has
//	the option to specify a positive number sign
//	character or characters for positive numeric
//	values using a Positive Number Sign Specification.
//
//	Example-1: Leading Positive Number Sign Symbols
//
//		Leading Symbols: "+ "
//		Number String:   "+ 123.456"
//
//	Example-2: Leading Positive Number Sign Symbols
//
//		Leading Symbols: "+"
//		Number String:   "+123.456"
//
//	Example-3: Trailing Positive Number Sign Symbols
//
//		Trailing Symbols: " +"
//		Number String:   "123.456 +"
//
//	Example-4: Trailing Positive Number Sign Symbols
//
//		Trailing Symbols: "+"
//		Number String:   "123.456+"
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
//	positiveNumSignSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the Positive
//		Number Sign	Symbol configured for the current
//		instance of	NumStrNumberSymbolGroup.
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
//	Returns a deep copy of the Zero Number Sign
//	Specification contained in the current instance
//	of NumStrNumberSymbolSpec.
//
//	The Zero Number Sign Specification is used to
//	configure number symbols for zero numeric values
//	formatted and displayed in number stings. Zero
//	number signs are commonly omitted because zero
//	does not technically qualify as either a positive or
//	negative value. However, the user does have the
//	option to configure leading and/or trailing Zero
//	Number Sign Symbols if needed.
//
//	Example-1: Leading Zero Number Sign Symbols
//
//		Leading Symbols: "+"
//		Trailing Symbols: ""
//		Number String:   "+0.00"
//
//	Example-1: Leading Zero Number Sign Symbols
//
//		Leading Symbols: "+ "
//		Trailing Symbols: ""
//		Number String:   "+ 0.00"
//
//	Example-3: Trailing Zero Number Sign Symbols
//
//		Leading Symbols: ""
//		Trailing Symbols: "+"
//		Number String:   "0.00+"
//
//	Example-4: Trailing Zero Number Sign Symbols
//
//		Leading Symbols: ""
//		Trailing Symbols: " +"
//		Number String:   "0.00 +"
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
//	zeroNumSignSpec			NumStrNumberSymbolSpec
//
//		If this method completes successfully, this
//		parameter will return a deep copy of the Zero
//		Number Sign Symbol Specification contained in the
//		current instance of NumStrNumberSymbolGroup.
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

//	NewCurrencyDefaultFrance
//
//	Creates and returns a new instance of
//	NumStrNumberSymbolGroup configured for French
//	Currency using default parameters.
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
//	The currency symbol used in the France is the
//	trailing Euro Currency Symbol ('€').
//
//		French Example-1
//		1 000 000,00 €
//
//	The positive number sign is implied. No positive
//	number is applied, only the trailing Euro Currency
//	Symbol.
//
//		French Example-2
//		1 000 000,00 €
//
//	The zero number format has no number sign, but the
//	currency symbol is set to a trailing Euro Currency
//	Symbol.
//
//		French Example-3
//			0,00 €
//
//	The negative number sign is set to leading minus
//	sign ('-') and a trailing Euro Currency Symbol
//	("€").
//
//		French Example-4
//		-1 000 000,00 €
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with Positive, Negative, Zero and Currency
//		Specification objects in compliance with French
//		Currency Formatting Standards.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewCurrencyDefaultFrance(
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
			"NewCurrencyDefaultFrance()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).
		setCurrencyDefaultsFrance(
			&newNStrNumSymbols,
			ePrefix.XCpy(
				"newNStrNumSymbols"))

	return newNStrNumSymbols, err
}

//	NewCurrencyDefaultGermany
//
//	Creates and returns a new instance of
//	NumStrNumberSymbolGroup configured for German
//	Currency using default parameters.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
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
//	The currency symbol used in the Germany is the
//	trailing Euro symbol ('€').
//
//		German Example-1
//		1.000.000,00 €
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-2
//		1.000.000- €
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		German Example-3
//		1.000.000 €
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		German Example-4
//			0,00 €
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with Positive, Negative, Zero and Currency
//		Specification objects in compliance with German
//		Currency Formatting Standards.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewCurrencyDefaultGermany(
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
			"NewCurrencyDefaultGermany()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).
		setCurrencyDefaultsGermany(
			&newNStrNumSymbols,
			ePrefix.XCpy(
				"newNStrNumSymbols"))

	return newNStrNumSymbols, err
}

//	NewNOP Creates and returns a new instance of
//	NumStrNumberSymbolGroup configured as a NOP.
//
//	'NOP' stands for 'No Operation'.
//
//	NOP is a computer science term which stands for
//	'No Operation' meaning it performs no operations
//	and serves an empty placeholder.
//
//	Type NumStrNumberSymbolGroup encapsulates four
//	instances of NumStrNumberSymbolSpec:
//
//		(1)	Negative Number Sign Specification
//
//		(2)	Positive Number Sign Specification
//
//		(3) Zero Number Sign Specification
//
//		(4)	Currency Symbol Specification
//
//	Configured as a NOP, the returned instance of
//	NumStrNumberSymbolGroup will contain four invalid
//	Number Symbol Specifications which function as
//	simple, empty placeholders. As such, these Number
//	Symbol Specifications perform no active role in,
//	and are completely ignored by, Number String
//	Formatting algorithms.
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
//		will be used in Number String Formatting
//		operations.
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

	newNStrNumSym.currencySymbol.SetNOP()

	return newNStrNumSym
}

//	NewSignedNumDefaultsFrance
//
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup configured for signed number
//	formatting in accordance with number string
//	formatting standards used in France.
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications. The Currency member
//	variable in the returned NumStrNumberSymbolGroup will
//	be empty and configured as a 'NOP' or empty
//	placeholder. 'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under French formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		French Example: Positive Numeric Value
//					123,45
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under French formatting standards,
//	zero numeric values have no number sign symbols.
//
//		French Example: Zero Numeric Value
//						0
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with Positive, Negative and Zero Number Sign
//		Symbol Specification objects conforming to
//		French Signed Number formatting standards.
//
//		This returned instance WILL NOT INCLUDE
//		Currency Symbols.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSignedNumDefaultsFrance(
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
			"NewSignedNumDefaultsFrance()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsFrance(
			&newNStrNumSymbols,
			ePrefix.XCpy(
				"newNStrNumSymbols<-"))

	return newNStrNumSymbols, err
}

//	NewSignedNumDefaultsGermany
//
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup configured for signed number
//	formatting in accordance with number string
//	formatting standards used in Germany.
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications. The Currency member
//	variable in the returned NumStrNumberSymbolGroup will
//	be empty and configured as a 'NOP' or empty
//	placeholder. 'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under German formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		German Example:	Positive Numeric Value
//					1.000.000
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under German formatting standards,
//	zero numeric values have no number sign symbols.
//
//		German Example:	Zero Numeric Value
//						0
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with Positive, Negative and Zero Number Sign
//		Symbol Specification objects conforming to
//		German Signed Number formatting standards.
//
//		This returned instance WILL NOT INCLUDE
//		Currency Symbols.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSignedNumDefaultsGermany(
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
			"NewSignedNumDefaultsGermany()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsGermany(
			&newNStrNumSymbols,
			ePrefix.XCpy(
				"newNStrNumSymbols<-"))

	return newNStrNumSymbols, err
}

//	NewSignedNumDefaultsUKMinus
//
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup configured for signed number
//	formatting in accordance with number string
//	formatting standards used in the United Kingdom (UK).
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		UK Example: Negative Numeric Value
//				-123
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications. The Currency member
//	variable in the returned NumStrNumberSymbolGroup will
//	be empty and configured as a 'NOP' or empty
//	placeholder. 'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under UK formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		UK Example:	Positive Numeric Value
//					123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under UK formatting standards,
//	zero numeric values have no number sign symbols.
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with Positive, Negative and Zero Number Sign
//		Symbol Specification objects conforming to
//		the United Kingdom Signed Number formatting
//		standards.
//
//		This returned instance WILL NOT INCLUDE
//		Currency Symbols.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSignedNumDefaultsUKMinus(
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
			"NewSignedNumDefaultsUKMinus()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsUKMinus(
			&newNStrNumSymbols,
			ePrefix.XCpy(
				"newNStrNumSymbols<-"))

	return newNStrNumSymbols, err
}

//	NewSignedNumDefaultsUSMinus
//
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup configured for signed number
//	formatting in accordance with number string
//	formatting standards used in the United States (US).
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		US Example: Negative Numeric Value
//				-123
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications. The Currency member
//	variable in the returned NumStrNumberSymbolGroup will
//	be empty and configured as a 'NOP' or empty
//	placeholder. 'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under US formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		US Example:	Positive Numeric Value
//					123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under US formatting standards,
//	zero numeric values have no number sign symbols.
//
//		US Example:	Zero Numeric Value
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with Positive, Negative and Zero Number Sign
//		Symbol Specification objects conforming to the
//		United States Signed Number formatting standards.
//
//		This returned instance WILL NOT INCLUDE
//		Currency Symbols.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSignedNumDefaultsUSMinus(
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
			"NewSignedNumDefaultsUSMinus()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsUSMinus(
			&newNStrNumSymbols,
			ePrefix.XCpy(
				"newNStrNumSymbols<-"))

	return newNStrNumSymbols, err
}

//	NewSignedNumDefaultsUSParen
//
//	Creates and returns and instance of
//	NumStrNumberSymbolGroup configured for signed number
//	formatting in accordance with number string
//	formatting standards used in the United States (US).
//
//
//	The term 'Paren' in the method name signals that
//	negative numeric values will be configured with a
//	surrounding parentheses ('()').
//
//		US Example: Negative Numeric Value
//					(123)
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The new, returned instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values.
//
//	Currency Symbols WILL NOT BE INCLUDED in the returned
//	number symbol specifications. The Currency member
//	variable in the returned NumStrNumberSymbolGroup will
//	be empty and configured as a 'NOP' or empty
//	placeholder. 'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under US formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		US Example:	Positive Numeric Value
//					123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under US formatting standards,
//	zero numeric values have no number sign symbols.
//
//		US Example:	Zero Numeric Value
//						0
//
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with Positive, Negative and Zero Number Sign
//		Symbol Specification objects conforming to the
//		United States Signed Number formatting standards.
//
//		This returned instance WILL NOT INCLUDE
//		Currency Symbols.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) NewSignedNumDefaultsUSParen(
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
			"NewSignedNumDefaultsUSParen()",
		"")

	if err != nil {
		return newNStrNumSymbols, err
	}

	err = new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsUSParen(
			&newNStrNumSymbols,
			ePrefix.XCpy(
				"newNStrNumSymbols<-"))

	return newNStrNumSymbols, err
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
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method provides a simplified means for
//	configuring the returned instance of
//	NumStrNumberSymbolGroup with currency and number sign
//	symbols. Using default values, this method will
//	generate Currency Symbols to be used with positive
//	numeric values, negative numeric values and zero
//	numeric values.
//
//	If the default configuration values fail to provide
//	sufficient granular control over currency number
//	symbol creation, use one of the more granular
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
//			European Number String:	"123.456- €"
//			US Number String:		"$ -123.456"
//			UK Number String:		"£ -123.45"
//
//	Negative Number Symbol:
//		The default Negative Number Symbol is the minus
//		sign ('-'). Negative numeric values will be
//		designated with the minus sign ('-').
//
//		The minus sign will be configured as a leading or
//		trailing minus sign depending on the value of
//		input parameter 'leadingNumSymbols'.
//
//		Examples:
//
//			Leading Minus Sign: "$ -123.456"
//			Trailing Minus Sign: "123.456-"
//
//	Positive Number Symbol:
//		No Positive Number Sign Symbol. Positive
//		values number signs are assumed and implicit. No
//		Number Signs will be formatted for positive
//		numeric values
//
//		Positive Numeric Value Example:
//					"123.456"
//
//	Zero Number Symbol:
//		No Zero Number Sign Symbol. Technically a zero
//		value is neither positive nor negative.
//		Consequently, no number sign is included with
//		zero numeric values.
//
//		Zero Numeric Value Example:
//					"0.00"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//			The minus sign is 'inside' the Number Field.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	currencySymbols				string
//
//		The symbol or symbols used to specify currency.
//		This currency symbol will be configured in the
//		new instance of NumStrNumberSymbolGroup returned
//		by this method.
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Number
//		Symbols at the beginning or left side of the
//		number string. Such Number Symbols are therefore
//		configured as leading Number Symbols. This is the
//		positioning format used in the US, UK, Australia
//		and most of Canada.
//
//		Example Number Strings:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted between
//				the currency symbol and the minus sign.
//
//		When 'leadingNumSymbols' is set to 'false', the
//		returned instance of NumStrNumberSymbolGroup will
//		configure Number Symbols on the right side of the
//		number string. Such	Number Symbols are therefore
//		configured as trailing Number Symbols. This is
//		the positioning format used in France, Germany
//		and many other countries in the European Union.
//
//			Example Number Strings:
//				"123.456- €"
//
//		NOTE:	A space is automatically inserted between
//				the minus sign and the currency symbol.
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
//	NumStrNumberSymbolGroup
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrNumberSymbolGroup configured
//		with Positive, Negative, Zero and Currency
//		Specification objects.
//
//		This returned NumStrNumberSymbolGroup instance
//		will therefore be configured with currency
//		symbols for inclusion in Number String
//		formatting operations.
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

	err = new(numStrNumberSymbolGroupMechanics).setSimpleCurrencySymbolsConfig(
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
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method provides a simplified means of creating
//	a new instance of type NumStrNumberSymbolGroup using
//	default values and generating number symbols suitable
//	in formatting signed numeric values.
//
//	Again, signed Number Symbols do NOT contain currency
//	symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over signed number symbol
//	creation, use one of the more advanced constructor or
//	'New' methods to implement specialized multinational
//	or multicultural signed number symbol formatting:
//
//		NumStrNumberSymbolGroup.NewSymbolsRunes()
//		NumStrNumberSymbolGroup.NewSymbolsStrings()
//		NumStrNumberSymbolGroup.NewSymbolsSpecs()
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The returned instance of NumStrNumberSymbolGroup will
//	NOT include currency symbols.
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
//
//		No Positive Number Sign Symbol. Positive
//		values are implicit and assumed. Therefore,
//		no plus is formatted in the number string.
//
//			Positive Value Number String: "123.456"
//
//	Zero Signed Number Symbol:
//
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative. Therefore,
//		no plus is formatted in the number string.
//
//			Zero Value Number String: "123.456"
//
//	Number Field Symbol Position:
//
//		Defaults to "Inside Number Field"
//
//			A Number Field defines the length of a
//			text string in which a numeric value is
//			formatted. When applied, a Number Field
//			is usually longer than the numeric value
//			string and typically justifies that numeric
//			value string as "Left", "Right" or "Centered".
//
//			Number Field Example (Centered):
//				"  123.45  "
//
//		Example Number Field Symbol Position:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Sign Symbols
//		in a Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Number
//		Sign Symbols on the left side of the numeric
//		value. Such Number Symbols are therefore
//		configured as leading Number Symbols. This is
//		the positioning format used in the US, UK,
//		Australia and most of Canada.
//
//		Example Number String with Leading Number Symbols:
//			"-123.456"
//
//		When set to 'false', the returned instance of
//		NumStrNumberSymbolGroup will configure Number
//		Symbols on the right side of the numeric value.
//		Such Number Symbols are therefore configured as
//		trailing Number Symbols. This is the positioning
//		format used in France, Germany and many other
//		countries in the European Union.
//
//		Example Number String with Trailing Number Symbols:
//			"123.456-"
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
//		Currency Symbols will NOT be included.
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

	err = new(numStrNumberSymbolGroupMechanics).
		setSimpleNumberSignSymbolsConfig(
			&newNStrNumSymbols,
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
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method generates a new instance of
//	NumStrNumberSymbolGroup using rune array input
//	parameters for positive number symbols, zero number
//	symbols, negative number symbols and currency
//	symbols.
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
//		combination of characters such as plus signs
//		('+').
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
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	trailingPositiveNumberSymbols	[]rune
//
//		A rune array containing the trailing positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
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
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
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
//		('-').
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
//		('-').
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
//	leadingZeroNumberSymbols	[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
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
//		positive or negative value. However, user have
//		the option to configure any combination of
//		symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
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
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
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
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'leadingCurrencySymbol' to 'nil'.
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
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'trailingCurrencySymbol' to 'nil'.
//
//	currencyNumFieldSymPosition		NumberFieldSymbolPosition
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
	leadingCurrencySymbols []rune,
	trailingCurrencySymbols []rune,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
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

	err = new(numStrNumberSymbolGroupMechanics).
		setNumberSymbolRunes(
			&newNumberSymbols,
			leadingPositiveNumberSymbols,
			trailingPositiveNumberSymbols,
			positiveNumFieldSymPosition,
			leadingNegativeNumberSymbols,
			trailingNegativeNumberSymbols,
			negativeNumFieldSymPosition,
			leadingZeroNumberSymbols,
			trailingZeroNumberSymbols,
			zeroNumFieldSymPosition,
			leadingCurrencySymbols,
			trailingCurrencySymbols,
			currencyNumFieldSymPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNumberSymbols<-"))

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
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method generates a new instance of
//	NumStrNumberSymbolGroup using string input
//	parameters for positive number symbols, zero number
//	symbols, negative number symbols and currency
//	symbols.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Currency symbol input parameters are optional and NOT
//	required.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingPositiveNumberSymbols	string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		Positive Number Sign Symbols in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
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
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
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
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
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
//		('-').
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
//			Leading Symbols: "-"
//			Number String:   "-123.456"
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
//		('-').
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
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to empty an empty string.
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
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to an empty string.
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
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'leadingCurrencySymbol' to 'nil'.
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
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'trailingCurrencySymbol' to 'nil'.
//
//	currencyNumFieldSymPosition		NumberFieldSymbolPosition
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
	leadingCurrencySymbols string,
	trailingCurrencySymbols string,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
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

	err = new(numStrNumberSymbolGroupMechanics).
		setNumberSymbolRunes(
			&newNumberSymbols,
			[]rune(leadingPositiveNumberSymbols),
			[]rune(trailingPositiveNumberSymbols),
			positiveNumFieldSymPosition,
			[]rune(leadingNegativeNumberSymbols),
			[]rune(trailingNegativeNumberSymbols),
			negativeNumFieldSymPosition,
			[]rune(leadingZeroNumberSymbols),
			[]rune(trailingZeroNumberSymbols),
			zeroNumFieldSymPosition,
			[]rune(leadingCurrencySymbols),
			[]rune(trailingCurrencySymbols),
			currencyNumFieldSymPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNumberSymbols<-"))

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
//
//	Type NumStrNumberSymbolGroup contains four instances
//	of type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method generates a new instance of
//	NumStrNumberSymbolGroup using Positive, Negative, Zero
//	Number Sign Symbols and Currency Symbol Specification
//	objects passed as input parameters.
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
//	currencySymbols				NumStrNumberSymbolSpec
//
//		This Currency Symbols Specification	will be
//		copied to the corresponding Currency Symbols
//		Specification in the new, returned instance
//		of NumStrNumberSymbolGroup.
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
	currencySymbols NumStrNumberSymbolSpec,
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
		setNumberSymbolSpecs(
			&newNumberSymbols,
			positiveNumberSign,
			negativeNumberSign,
			zeroNumberSign,
			currencySymbols,
			ePrefix.XCpy(
				"newNumberSymbols"))

	return newNumberSymbols, err
}

//	SetCurrencyDefaultFrance
//
//	Deletes all data values in the current instance of
//	NumStrNumberSymbolGroup and proceeds to reconfigure
//	that instance with number sign and currency symbols
//	commonly applied to French currency numeric values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolGroup.
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
//	The currency symbol used in the France is the
//	trailing Euro Currency Symbol ('€').
//
//		French Example-1
//		1 000 000,00 €
//
//	The positive number sign is implied. No positive
//	number is applied, only the trailing Euro Currency
//	Symbol.
//
//		French Example-2
//		1 000 000,00 €
//
//	The zero number format has no number sign, but the
//	currency symbol is set to a trailing Euro Currency
//	Symbol.
//
//		French Example-3
//			0,00 €
//
//	The negative number sign is set to leading minus
//	sign ('-') and a trailing Euro Currency Symbol
//	("€").
//
//		French Example-4
//		-1 000 000,00 €
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetCurrencyDefaultFrance(
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
			"SetCurrencyDefaultFrance()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setCurrencyDefaultsFrance(
			nStrNumSymbolsGroup,
			ePrefix.XCpy(
				"newNStrNumSymbols"))
}

//	SetCurrencyDefaultGermany
//
//	Deletes all data values in the current instance of
//	NumStrNumberSymbolGroup and proceeds to reconfigure
//	that instance with number sign and currency symbols
//	commonly applied to German currency numeric values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the current instance of
//	NumStrNumberSymbolGroup.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
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
//	The currency symbol used in the Germany is the
//	trailing Euro symbol ('€').
//
//		German Example-1
//		1.000.000,00 €
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-2
//		1.000.000- €
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		German Example-3
//		1.000.000 €
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		German Example-4
//			0,00 €
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetCurrencyDefaultGermany(
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
			"SetCurrencyDefaultGermany()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setCurrencyDefaultsGermany(
			nStrNumSymbolsGroup,
			ePrefix.XCpy(
				"newNStrNumSymbols"))
}

//	SetNegativeNumSignRunes
//
//	Deletes and resets the Negative Number Symbols data
//	fields for the current instance of
//	NumStrNumberSymbolGroup:
//
//		NumStrNumberSymbolGroup.negativeNumberSign
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
//		('-').
//
//		Example-1: Leading Negative Number Symbols
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Negative Number Symbols
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
//		('-').
//
//		Example-1: Trailing Negative Number Symbols
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Negative Number Symbols
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNegativeNumSignRunes(
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
			"SetNegativeNumSignRunes()",
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

//	SetNegativeNumSignSpec
//
//
//	Receives a single NumStrNumberSymbolSpec object
//	configured as a Negative Number Sign Symbol.
//
//	This method then proceeds to reset the
//	corresponding Negative Number Sign Symbol member
//	variable data value for the current instance of
//	NumStrNumberSymbolGroup:
//
//		NumStrNumberSymbolGroup.negativeNumberSign
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Be advised that this method will delete and reset
//		the Negative Number Sign Symbol member variable
//		data fields contained in the current instance of
//		NumStrNumberSymbolGroup:
//
//			NumStrNumberSymbolGroup.negativeNumberSign
//
//	(2)	If input parameter 'negativeNumberSign' is
//		invalid or a NOP, an error will be returned.
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
//		NumStrNumberSymbolGroup:
//
//			NumStrNumberSymbolSpec.negativeNumberSign
//
//		If input parameter 'negativeNumberSign' is
//		invalid or a NOP, an error will be returned.
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

//	SetNegativeNumSignStrings
//
//	Deletes and resets the Negative Number Symbols data
//	fields for the current instance of
//	NumStrNumberSymbolGroup.
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
//		('-').
//
//		Example-1: Leading Negative Number Symbols
//
//			Leading Symbols: "- "
//			Number String:   "- 123.456"
//
//		Example-2: Leading Negative Number Symbols
//
//			Leading Symbols: "-"
//			Number String:   "-123.456"
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
//		('-').
//
//		Example-1: Trailing Negative Number Symbols
//
//			Trailing Symbols: " -"
//			Number String:   "123.456 -"
//
//		Example-2: Trailing Negative Number Symbols
//
//			Trailing Symbols: "-"
//			Number String:   "123.456-"
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNegativeNumSignStrings(
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
			"SetNegativeNumSignStrings()",
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
//	instance of NumStrNumberSymbolGroup to a NOP or 'No
//	Operation'. This means that all data values contained
//	in the current NumStrNumberSymbolGroup instance will
//	be deleted and reset to their zero or uninitialized
//	states.
//
//	Type NumStrNumberSymbolGroup encapsulates the
//	Negative, Positive, Zero and Currency Symbol
//	Specifications.
//
//	Calling this method will convert all four Number
//	Symbol Specifications to empty placeholders and
//	ensure that they perform no active role in, and are
//	completely ignored by, Number String Formatting
//	algorithms. In this state, none of the four Number
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
//		NumStrNumberSymbolGroup.IsNOPCurrencySymbols()
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

//	SetNOPNegativeNumSignSymbol
//
//	'NOP' stands for 'No Operation'.
//
//	When called, this method will convert the Negative
//	Number Sign Symbol Specification in the current
//	instance of NumStrNumberSymbolGroup to a NOP or
//	'No Operation'. This means that the Negative Number
//	Sign Symbol Specification is simply an empty
//	placeholder and performs no active role in, and is
//	completely ignored by, Number String Formatting
//	algorithms.
//
//	Examples of Negative Number Signs include the minus
//	sign ('-') and surrounding parentheses ('()').
//
//		Example Negative Number Signs
//
//				"-123.45"
//				"123.45-"
//				"(123.45)"
//
//	After conversion to a NOP, no Negative Number Sign
//	Symbols will be inserted or formatted as part of
//	a Number String formatting operation.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNOPNegativeNumSignSymbol() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.negativeNumberSign.SetNOP()

	return
}

//	SetNOPPositiveNumSignSymbol
//
//	'NOP' stands for 'No Operation'.
//
//	When called, this method will convert the Positive
//	Number Sign Symbol Specification in the current
//	instance of NumStrNumberSymbolGroup to a NOP or
//	'No Operation'. This means that the Positive Number
//	Sign Symbol Specification is simply an empty
//	placeholder and performs no active role in, and is
//	completely ignored by, Number String Formatting
//	algorithms.
//
//	Positive number signs are usually omitted from
//	positive numeric values when displayed in number
//	strings. The number sign is considered implicit
//	or assumed for positive numeric values. However,
//	Type NumStrNumberSymbolGroup allows users the option
//	of configuring both leading and trailing positive
//	number sign symbols.
//
//		Examples of Positive Number Sign Symbols
//
//				"+123.45"
//				"123.45+"
//
//	After conversion to a NOP, no Positive Number Sign
//	Symbols	will be inserted or formatted as part of a
//	Number String formatting operation.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNOPPositiveNumSignSymbol() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.positiveNumberSign.SetNOP()

	return
}

//	SetNOPZeroNumSignSymbol
//
//	'NOP' stands for 'No Operation'.
//
//	When called, this method will convert the Zero
//	Number Sign Symbol Specification in the current
//	instance of NumStrNumberSymbolGroup to a NOP or
//	'No Operation'. This means that the Zero Number
//	Sign Symbol Specification is simply an empty
//	placeholder and performs no active role in, and is
//	completely ignored by, Number String Formatting
//	algorithms.
//
//	Zero Number Sign Symbols are usually omitted from
//	zero numeric values displayed in number strings.
//	Technically, the value zero is neither positive nor
//	negative and therefore number signs are unnecessary.
//	However, for Type NumStrNumberSymbolGroup, users have
//	the option of configuring both leading and trailing
//	zero number sign symbols.
//
//		Example Zero Number Signs
//
//				"+0.00"
//				"0.00+"
//
//	After conversion to a NOP, no Zero Number Sign
//	Symbols will be inserted or formatted as part of a
//	Number String formatting operation.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetNOPZeroNumSignSymbol() {

	if nStrNumSymbolsGroup.lock == nil {
		nStrNumSymbolsGroup.lock = new(sync.Mutex)
	}

	nStrNumSymbolsGroup.lock.Lock()

	defer nStrNumSymbolsGroup.lock.Unlock()

	nStrNumSymbolsGroup.zeroNumberSign.SetNOP()

	return
}

//	SetPositiveNumSignRunes
//
//	Deletes and resets the Positive Number Symbols data
//	fields for the current instance of
//	NumStrNumberSymbolGroup.
//
//	Positive number signs are usually omitted from
//	positive numeric values when displayed in number
//	strings. The number sign is considered implicit
//	or assumed for positive numeric values. However,
//	Type NumStrNumberSymbolGroup allows users the option
//	of configuring both leading and trailing positive
//	number sign symbols.
//
//		Examples of Positive Number Sign Symbols
//
//				"+123.45"
//				"123.45+"
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Positive Number Sign Symbol member variable data
//	fields as specified by the Positive Number Sign
//	Symbol input parameters passed as rune arrays.
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
//		Leading positive number sign symbols can include
//		any combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Positive Number Sign Symbols
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Positive Number Sign Symbols
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//	trailingPositiveNumberSymbols	[]rune
//
//		A rune array containing the trailing positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Trailing positive number sign symbols can include
//		any combination of characters such as plus signs
//		('+').
//
//		Example-1: Trailing Positive Number Sign Symbols
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Positive Number Sign Symbols
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetPositiveNumSignRunes(
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
			"SetPositiveNumSignRunes()",
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
//	(1)	Be advised that this method will delete and reset
//		the Positive Number Sign Symbol member variable
//		data fields contained in the current instance of
//		NumStrNumberSymbolGroup:
//
//			NumStrNumberSymbolGroup.positiveNumberSign
//
//	(2)	If input parameter 'positiveNumberSign' is invalid
//		or a NOP, an error will be returned.
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
//		NumStrNumberSymbolGroup:
//
//			NumStrNumberSymbolGroup.positiveNumberSign
//
//		If input parameter 'positiveNumberSign' is
//		invalid or a NOP, an error will be returned.
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

//	SetPositiveNumSignStrings
//
//	Deletes and resets the Positive Number Sign Symbols
//	data fields for the current instance of
//	NumStrNumberSymbolGroup.
//
//	Positive number signs are usually omitted from
//	positive numeric values when displayed in number
//	strings. The number sign is considered implicit
//	or assumed for positive numeric values. However,
//	Type NumStrNumberSymbolGroup allows users the option
//	of configuring both leading and trailing positive
//	number sign symbols.
//
//		Examples of Positive Number Sign Symbols
//
//				"+123.45"
//				"123.45+"
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
//		Leading positive number sign symbols can include
//		any combination of characters such as plus signs
//		('+').
//
//		Example-1: Leading Positive Number Sign Symbols
//
//			Leading Symbols: "+ "
//			Number String:   "+ 123.456"
//
//		Example-2: Leading Positive Number Sign Symbols
//
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//	trailingPositiveNumberSymbols	string
//
//		A string containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing positive number sign symbols can include
//		any combination of characters such as plus signs
//		('+').
//
//		Example-1: Trailing Positive Number Sign Symbols
//
//			Trailing Symbols: " +"
//			Number String:   "123.456 +"
//
//		Example-2: Trailing Positive Number Sign Symbols
//
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetPositiveNumSignStrings(
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
			"SetPositiveNumSignStrings()",
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
// currency symbol parameters.
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
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method provides a simplified means for
//	reconfiguring the current instance of
//	NumStrNumberSymbolGroup with currency and number sign
//	symbols. Using default values, this method will
//	generate Currency Symbols to be used with positive
//	numeric values, negative numeric values and zero
//	numeric values.
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
//			European Number String:	"123.456- €"
//			US Number String:		"$ -123.456"
//			UK Number String:		"£ -123.45"
//
//	Negative Number Symbol:
//		The default Negative Number Symbol is the minus
//		sign ('-'). Negative numeric values will be
//		designated with the minus sign ('-').
//
//		The minus sign will be configured as a leading or
//		trailing minus sign depending on the value of
//		input parameter 'leadingNumSymbols'.
//
//		Examples:
//
//			Leading Minus Sign: "$ -123.456"
//			Trailing Minus Sign: "123.456-"
//
//	Positive Number Symbol:
//		No Positive Number Sign Symbol. Positive
//		values number signs are assumed and implicit. No
//		Number Signs will be formatted for positive
//		numeric values
//
//		Positive Numeric Value Example:
//					"123.456"
//
//	Zero Number Symbol:
//		No Zero Number Sign Symbol. Technically a zero
//		value is neither positive nor negative.
//		Consequently, no number sign is included with
//		zero numeric values.
//
//		Zero Numeric Value Example:
//					"0.00"
//
//	Number Field Symbol Position:
//		Defaults to "Inside Number Field"
//
//		Example:
//			Number Field Length: 8
//			Numeric Value: -123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//			The minus sign is 'inside' the Number Field.
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
//		and many other countries in the European Union.
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

	return new(numStrNumberSymbolGroupMechanics).
		setSimpleCurrencySymbolsConfig(
			nStrNumSymbolsGroup,
			currencySymbols,
			leadingNumSymbols,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup<-"))
}

//	SetSignedNumDefaultsFrance
//
//	Deletes and reconfigures all the data values
//	contained in the current instance of
//	NumStrNumberSymbolGroup and the proceeds to
//	reconfigure that instance with default Signed
//	Number formatting specifications typically applied
//	in France.
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The reconfigured instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values conforming to French formatting standards.
//
//	Currency Symbols WILL NOT BE INCLUDED in the
//	reconfigured number symbol specifications. The
//	Currency member variable in NumStrNumberSymbolGroup
//	will be configured as a 'NOP' or empty placeholder.
//	'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under French formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		French Example: Positive Numeric Value
//					123,45
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under French formatting standards,
//	zero numeric values have no number sign symbols.
//
//		French Example: Zero Numeric Value
//						0
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSignedNumDefaultsFrance(
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
			"SetSignedNumDefaultsFrance()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsFrance(
			nStrNumSymbolsGroup,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup<-"))

}

//	SetSignedNumDefaultsGermany
//
//	Deletes and reconfigures all the data values
//	contained in the current instance of
//	NumStrNumberSymbolGroup and the proceeds to
//	reconfigure that instance with default Signed Number
//	formatting specifications typically applied in
//	Germany.
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The reconfigured instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values conforming to German formatting standards.
//
//	Currency Symbols WILL NOT BE INCLUDED in the
//	reconfigured number symbol specifications. The
//	Currency member variable in NumStrNumberSymbolGroup
//	will be configured as a 'NOP' or empty placeholder.
//	'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under German formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		German Example:	Positive Numeric Value
//					1.000.000
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under German formatting standards,
//	zero numeric values have no number sign symbols.
//
//		German Example:	Zero Numeric Value
//						0
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSignedNumDefaultsGermany(
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
			"SetSignedNumDefaultsGermany()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsGermany(
			nStrNumSymbolsGroup,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup<-"))

}

//	SetSignedNumDefaultsUKMinus
//
//	Deletes and reconfigures all the data values
//	contained in the current instance of
//	NumStrNumberSymbolGroup and the proceeds to
//	reconfigure that instance with default Signed Number
//	formatting specifications typically applied in
//	the United Kingdom (UK).
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		UK Example: Negative Numeric Value
//				-123
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The reconfigured instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values conforming to UK formatting standards.
//
//	Currency Symbols WILL NOT BE INCLUDED in the
//	reconfigured number symbol specifications. The
//	Currency member variable in NumStrNumberSymbolGroup
//	will be configured as a 'NOP' or empty placeholder.
//	'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under UK formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		UK Example:	Positive Numeric Value
//					123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under UK formatting standards,
//	zero numeric values have no number sign symbols.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSignedNumDefaultsUKMinus(
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
			"SetSignedNumDefaultsUKMinus()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsUKMinus(
			nStrNumSymbolsGroup,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup<-"))

}

//	SetSignedNumDefaultsUSMinus
//
//	Deletes and reconfigures all the data values
//	contained in the current instance of
//	NumStrNumberSymbolGroup and the proceeds to
//	reconfigure that instance with default Signed Number
//	formatting specifications typically applied in
//	the United States (US).
//
//	The word 'Minus' in the method name signals that
//	negative numeric values will be configured with a
//	leading minus sign ('-').
//
//		US Example: Negative Numeric Value
//				-123
//
//	A signed number is a numeric value formatted in a
//	number string which does NOT contain currency
//	symbols.
//
//	The reconfigured instance of
//	NumStrNumberSymbolGroup will include signed number
//	symbols for positive, zero and negative numeric
//	values conforming to US formatting standards.
//
//	Currency Symbols WILL NOT BE INCLUDED in the
//	reconfigured number symbol specifications. The
//	Currency member variable in NumStrNumberSymbolGroup
//	will be configured as a 'NOP' or empty placeholder.
//	'NOP' stands for 'No Operation'.
//
//	The positive signed number symbol is empty or blank
//	because under US formatting standards, positive
//	number signs are implied and not explicitly
//	displayed. Therefore, no leading plus ('+') symbol is
//	required.
//
//		US Example:	Positive Numeric Value
//					123
//
//	Likewise, the zero signed number symbol is also empty
//	or blank because under US formatting standards,
//	zero numeric values have no number sign symbols.
//
//		US Example:	Zero Numeric Value
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSignedNumDefaultsUSMinus(
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
			"SetSignedNumDefaultsUSMinus()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setSignedNumDefaultsUSMinus(
			nStrNumSymbolsGroup,
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
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method provides a simplified means for
//	reconfiguring the current instance of
//	NumStrNumberSymbolGroup using default values and
//	generating number symbols suitable in formatting
//	signed numeric values.
//
//	Again, Signed Number Symbols do NOT contain currency
//	symbols.
//
//	If the default configuration values fail to provide
//	sufficient granular control over signed number symbol
//	creation, use one of the more advanced 'Setter'
//	methods to implement specialized multinational or
//	multicultural signed number symbol formatting:
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
//
//		No Positive Number Sign Symbol. Positive
//		values are implicit and assumed. Therefore,
//		no plus is formatted in the number string.
//
//			Positive Value Number String: "123.456"
//
//	Zero Signed Number Symbol:
//
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative. Therefore,
//		no plus is formatted in the number string.
//
//			Zero Value Number String: "123.456"
//
//	Number Field Symbol Position:
//
//		Defaults to "Inside Number Field"
//
//			A Number Field defines the length of a
//			text string in which a numeric value is
//			formatted. When applied, a Number Field
//			is usually longer than the numeric value
//			string and typically justifies that numeric
//			value string as "Left", "Right" or "Centered".
//
//			Number Field Example (Centered):
//				"  123.45  "
//
//		Example Number Field Symbol Position:
//			Number Field Length: 8
//			Numeric Value: 123.45
//			Number Symbol: leading minus sign ('-')
//			Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			Number Field Index:------>01234567
//			Total Number String Length: 8
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Sign Symbols
//		in a Number String Format.
//
//		When set to 'true', the returned instance of
//		NumStrNumberSymbolGroup will configure Number
//		Sign Symbols on the left side of the numeric
//		value. Such Number Symbols are therefore
//		configured as leading Number Symbols. This is
//		the positioning format used in the US, UK,
//		Australia and most of Canada.
//
//		Example Number String with Leading Number Symbols:
//			"-123.456"
//
//		When set to 'false', the returned instance of
//		NumStrNumberSymbolGroup will configure Number
//		Symbols on the right side of the numeric value.
//		Such Number Symbols are therefore configured as
//		trailing Number Symbols. This is the positioning
//		format used in France, Germany and many other
//		countries in the European Union.
//
//		Example Number String with Trailing Number Symbols:
//			"123.456-"
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
			"SetSimpleSignedNumber()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setSimpleNumberSignSymbolsConfig(
			nStrNumSymbolsGroup,
			leadingNumSymbols,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup<-"))
}

//	SetSymbolsRunes
//
//	Reconfigures the current instance of
//	NumStrNumberSymbolGroup based on the Positive,
//	Negative, Zero and Currency symbol rune arrays
//	passed as input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbolGroup will be
//	deleted and replaced with new data values passed as
//	input parameters.
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
//		combination of characters such as plus signs
//		('+').
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
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
//
//	trailingPositiveNumberSymbols	[]rune
//
//		A rune array containing the trailing positive
//		number sign character or characters used to
//		configure Positive Number Sign Symbols in a
//		number string with a positive numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
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
//			Trailing Symbols: "+"
//			Number String:   "123.456+"
//
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to 'nil' for an empty rune array.
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
//		('-').
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
//		('-').
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
//	leadingZeroNumberSymbols	[]rune
//
//		A rune array containing the leading zero
//		number sign character or characters used to
//		configure Zero Number Sign Symbols in a
//		number string with a zero numeric value.
//
//		Zero number signs are commonly omitted because
//		zero does not technically qualify as either a
//		positive or negative value. However, users have
//		the option of configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "+"
//			Trailing Symbols: ""
//			Number String:   "+0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
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
//		positive or negative value. However, user have
//		the option to configure any combination of
//		symbols.
//
//		Trailing number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to 'nil' for an empty rune
//		array.
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
//	leadingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Leading
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
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
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'leadingCurrencySymbol' to 'nil'.
//
//	trailingCurrencySymbol     		[]rune
//
//		A rune array containing one or more Trailing
//		Currency Symbol characters used to configure
//		the current instance of NumStrNumberSymbolSpec.
//
//		Trailing Currency Symbol characters can include
//		such symbols as the Dollar sign ('$'), Euro sign
//	 	('€') and Pound sign ('£').
//
//		Trailing Currency Symbols are suffixed or
//		appended to the end of number strings containing
//		currency numeric values.
//
//				Example: 125.34€
//
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'trailingCurrencySymbol' to 'nil'.
//
//	currencyNumFieldSymPosition		NumberFieldSymbolPosition
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
	leadingCurrencySymbols []rune,
	trailingCurrencySymbols []rune,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
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

	return new(numStrNumberSymbolGroupMechanics).
		setNumberSymbolRunes(
			nStrNumSymbolsGroup,
			leadingPositiveNumberSymbols,
			trailingPositiveNumberSymbols,
			positiveNumFieldSymPosition,
			leadingNegativeNumberSymbols,
			trailingNegativeNumberSymbols,
			negativeNumFieldSymPosition,
			leadingZeroNumberSymbols,
			trailingZeroNumberSymbols,
			zeroNumFieldSymPosition,
			leadingCurrencySymbols,
			trailingCurrencySymbols,
			currencyNumFieldSymPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNumberSymbols<-"))
}

//	SetSymbolsSpecs
//
//	Reconfigures the current instance of
//	NumStrNumberSymbolGroup based on the Positive,
//	Negative Zero Number Sign and Currency Symbol
//	Specification objects passed as input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbolGroup will be
//	deleted and replaced with the Positive, Negative
//	Zero Number Sign and Currency Symbol Specifications
//	passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Positive Number Sign Symbol Specification
//		will be copied to the corresponding Positive
//		Number Sign Symbol Specification in the current
//		instance of NumStrNumberSymbolGroup.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Negative Number Sign Symbol Specification
//		will be copied to the corresponding Negative
//		Number Sign Symbol Specification in the current
//		instance of NumStrNumberSymbolGroup.
//
//	zeroNumberSign				NumStrNumberSymbolSpec
//
//		This Zero Number Sign Symbol Specification
//		will be copied to the corresponding Zero
//		Number Sign Symbol Specification in the current
//		instance of NumStrNumberSymbolGroup.
//
//	currencySymbols				NumStrNumberSymbolSpec
//
//		This Currency Symbols Specification	will be
//		copied to the corresponding Currency Symbols
//		Specification in the current instance of
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetSymbolsSpecs(
	positiveNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	currencySymbols NumStrNumberSymbolSpec,
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
		setNumberSymbolSpecs(
			nStrNumSymbolsGroup,
			positiveNumberSign,
			negativeNumberSign,
			zeroNumberSign,
			currencySymbols,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup"))

}

//	SetSymbolsStrings
//
//	Reconfigures the current instance of
//	NumStrNumberSymbolGroup based on the Positive,
//	Negative, Zero and Currency symbol strings passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrNumberSymbolGroup will be
//	deleted and replaced with new data values passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	leadingPositiveNumberSymbols	string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		Positive Number Sign Symbols in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
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
//			Leading Symbols: "+"
//			Number String:   "+123.456"
//
//		Leading number symbols are often omitted for
//		positive numeric values. If leading positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
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
//		Trailing number symbols are often omitted for
//		positive numeric values. If trailing positive
//		number symbols are NOT required, set this
//		parameter to an empty string.
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
//		('-').
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
//			Leading Symbols: "-"
//			Number String:   "-123.456"
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
//		('-').
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
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+').
//
//		Example: Leading Number Symbols
//			Leading Number Symbols for Zero Values
//
//			Leading Symbols: "$"
//			Trailing Symbols: ""
//			Number String:   "$0.00"
//
//		If leading zero number symbols are NOT required,
//		set this parameter to empty an empty string.
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
//		positive or negative value. However, users have
//		the option to configure any combination of
//		symbols for zero numeric values.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+').
//
//		Example: Trailing Number Symbols
//			Trailing Number Symbols for Zero Values
//
//			Leading Symbols: ""
//			Trailing Symbols: " +"
//			Number String:   "0.00 +"
//
//		If trailing zero number symbols are NOT required,
//		set this parameter to an empty string.
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
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'leadingCurrencySymbol' to 'nil'.
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
//		Currency Symbols are optional. If Currency
//		Symbols are not required, set
//		'trailingCurrencySymbol' to 'nil'.
//
//	currencyNumFieldSymPosition		NumberFieldSymbolPosition
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
	leadingCurrencySymbols string,
	trailingCurrencySymbols string,
	currencyNumFieldSymPosition NumberFieldSymbolPosition,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
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
			"NewSymbolsStrings()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupMechanics).
		setNumberSymbolRunes(
			nStrNumSymbolsGroup,
			[]rune(leadingPositiveNumberSymbols),
			[]rune(trailingPositiveNumberSymbols),
			positiveNumFieldSymPosition,
			[]rune(leadingNegativeNumberSymbols),
			[]rune(trailingNegativeNumberSymbols),
			negativeNumFieldSymPosition,
			[]rune(leadingZeroNumberSymbols),
			[]rune(trailingZeroNumberSymbols),
			zeroNumFieldSymPosition,
			[]rune(leadingCurrencySymbols),
			[]rune(trailingCurrencySymbols),
			currencyNumFieldSymPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"newNumberSymbols<-"))
}

//	SetZeroNumSignRunes
//
//	Deletes and resets the Zero Number Sign Symbols data
//	fields for the current instance of
//	NumStrNumberSymbolGroup.
//
//	The Zero Number Sign Symbols member variable for the
//	current instance of NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Zero Number Sign Symbol member variable data fields
//	as specified by the Zero Number Sign Symbol input
//	parameters passed as rune arrays.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetZeroNumSignRunes(
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
			"SetZeroNumSignRunes()",
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
//	Receives a  Zero Number Sign Symbol Specification
//	object passed as an input parameter
//	('zeroNumberSign').
//
//	This method then proceeds to reset the
//	corresponding Zero Number Sign Symbol member
//	variable data value for the current instance of
//	NumStrNumberSymbolGroup.
//
//	The Zero Number Sign Symbol internal member
//	variable for the current instance of
//	NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	Be advised that this method will delete and reset
//		the Zero Number Sign Symbol member variable data
//		fields contained in the current instance of
//		NumStrNumberSymbolGroup:
//
//			NumStrNumberSymbolGroup.zeroNumberSign
//
//	(2)	If input parameter 'zeroNumberSign' is invalid
//		or a NOP, an error will be returned.
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
//		NumStrNumberSymbolGroup:
//
//			NumStrNumberSymbolGroup.zeroNumberSign
//
//		If input parameter 'zeroNumberSign' is invalid
//		or a NOP, an error will be returned.
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
			"SetZeroNumSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberSymbolGroupNanobot).
		setZeroNumSignSpec(
			nStrNumSymbolsGroup,
			zeroNumberSign,
			ePrefix.XCpy(
				"nStrNumSymbolsGroup"))

}

//	SetZeroNumSignStrings
//
//	Deletes and resets the Zero Number Sign Symbols data
//	fields for the current instance of
//	NumStrNumberSymbolGroup.
//
//	The Zero Number Sign Symbols member variable for the
//	current instance of NumStrNumberSymbolGroup is:
//
//		NumStrNumberSymbolGroup.zeroNumberSign
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and reset the
//	Zero Number Sign Symbol member variable data
//	fields as specified by the Zero Number Sign Symbol
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
//	trailingZeroNumberSymbols		string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
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
func (nStrNumSymbolsGroup *NumStrNumberSymbolGroup) SetZeroNumSignStrings(
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
			"SetZeroNumSignStrings()",
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

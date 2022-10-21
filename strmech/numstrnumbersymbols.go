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

//	NewSymbolSpecs
//
//	Creates and returns a new instance of
//	NumStrNumberSymbols. The new instance is generated
//	from the Positive, Negative and Zero Number Sign
//	Symbol objects passed as input parameters.
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
//		Sign Symbol objects passed as input parameters.
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
func (nStrNumSym *NumStrNumberSymbols) NewSymbolSpecs(
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
			"NewSymbolSpecs()",
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

type numStrNumberSymbolsMechanics struct {
	lock *sync.Mutex
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

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

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

//	setCurrencyDefaultsFrance
//
//	Receives an instance of NumStrNumberSymbolGroup,
//	deletes the pre-existing data values and proceeds to
//	reconfigure the instance according to French currency
//	formatting standards.
//
//	Default values will be used to reconfigure the
//	NumStrNumberSymbolGroup input parameter,
//	'nStrNumSymbolGroup', for French Currency formatting.
//	New data values will be configured for the positive,
//	zero and negative number sign symbols as well as the
//	currency symbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	NumStrNumberSymbolGroup passed as input parameter
//	'nStrNumSymbolGroup'.
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
//	The currency symbol used in the France is the
//	trailing Euro symbol ('€').
//
//		French Example-1
//		1 000 000,00 €
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		French Example-2
//		-1 000 000 000
//
//	The number sign is implied for positive numeric values.
//	Therefore, the positive number sign symbol is set to a
//	blank or empty string ("").
//
//		French Example-3
//		1 000 000 000 €
//
//	The zero number format has no number sign. Therefore,
//	the positive number sign symbol is set to a blank or
//	empty string ("").
//
//		French Example-4
//			0,0 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured using
//		the default parameters for French Currency number
//		formatting standards.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setCurrencyDefaultsFrance(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setCurrencyDefaultsFrance()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = nStrNumSymbolGroup.currencySymbol.
		SetCurrencyDefaultEU(
			ePrefix.XCpy(
				"nStrNumSymbolGroup.currencySymbol"))

	if err != nil {
		return err
	}

	return new(NumStrNumberSymbolSpec).
		SetSignedNumDefaultsFrance(
			&nStrNumSymbolGroup.positiveNumberSign,
			&nStrNumSymbolGroup.zeroNumberSign,
			&nStrNumSymbolGroup.negativeNumberSign,
			ePrefix)
}

//	setCurrencyDefaultsGermany
//
//	Receives an instance of NumStrNumberSymbolGroup,
//	deletes the pre-existing data values and proceeds to
//	reconfigure the instance according to German currency
//	formatting standards.
//
//	Default values will be used to reconfigure the
//	NumStrNumberSymbolGroup input parameter,
//	'nStrNumSymbolGroup', for German Currency formatting.
//	New data values will be configured for the positive,
//	zero and negative number sign symbols as well as the
//	currency symbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	NumStrNumberSymbolGroup passed as input parameter
//	'nStrNumSymbolGroup'.
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
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured using
//		the default parameters for German Currency number
//		formatting standards.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setCurrencyDefaultsGermany(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setCurrencyDefaultsGermany()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = nStrNumSymbolGroup.currencySymbol.
		SetCurrencyDefaultEU(
			ePrefix.XCpy(
				"nStrNumSymbolGroup.currencySymbol"))

	if err != nil {
		return err
	}

	return new(NumStrNumberSymbolSpec).
		SetSignedNumDefaultsGermany(
			&nStrNumSymbolGroup.positiveNumberSign,
			&nStrNumSymbolGroup.zeroNumberSign,
			&nStrNumSymbolGroup.negativeNumberSign,
			ePrefix)
}

//	setCurrencyDefaultsUKMinusInside
//
//	Receives an instance of NumStrNumberSymbolGroup,
//	deletes the pre-existing data values and proceeds to
//	reconfigure the instance according to the United
//	Kingdom (UK) currency formatting standards.
//
//	Default values will be used to reconfigure the
//	NumStrNumberSymbolGroup input parameter,
//	'nStrNumSymbolGroup', for UK Currency formatting.
//	New data values will be configured for the positive,
//	zero and negative number sign symbols as well as the
//	currency symbols.
//
//	The phrase 'Minus Inside' in the method name signals
//	that for negative values, the minus sign ('-') will
//	be to the right of, or inside, the currency symbol.
//
//		UK Example
//			£ -123.45  Negative Value
//
//	To configure the minus sign outside the currency
//	symbol see method:
//
//		numStrNumberSymbolGroupMechanics.
//			setCurrencyDefaultsUKMinusOutside()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	NumStrNumberSymbolGroup passed as input parameter
//	'nStrNumSymbolGroup'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The default UK currency symbol is a leading pound
//	sign ('£').
//
//			UK Example-1:
//				£ 123.45  Positive Value
//
//		The negative number sign is set to a leading minus
//		sign ('-').
//
//			UK Example-2
//				-123.45  Negative Value
//
//	This method will configure the pound sign ('£')
//	such that any minus sign configured for negative
//	numeric values will be inside the pound sign ('£').
//
//			UK Example-3
//				£ -123.45  Negative Value
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		UK Example-4
//			£ 123.45
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		UK Example-5
//			£ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured using
//		the default parameters for UK Currency number
//		formatting standards.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setCurrencyDefaultsUKMinusInside(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setCurrencyDefaultsUKMinusInside()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = nStrNumSymbolGroup.currencySymbol.
		SetCurrencyDefaultsUKMinusInside(
			ePrefix.XCpy(
				"nStrNumSymbolGroup.currencySymbol"))

	if err != nil {
		return err
	}

	return new(NumStrNumberSymbolSpec).
		SetSignedNumDefaultsUKMinus(
			&nStrNumSymbolGroup.positiveNumberSign,
			&nStrNumSymbolGroup.zeroNumberSign,
			&nStrNumSymbolGroup.negativeNumberSign,
			ePrefix)
}

//	setCurrencyDefaultsUKMinusOutside
//
//	Receives an instance of NumStrNumberSymbolGroup,
//	deletes the pre-existing data values and proceeds to
//	reconfigure the instance according to the United
//	Kingdom (UK) currency formatting standards.
//
//	Default values will be used to reconfigure the
//	NumStrNumberSymbolGroup input parameter,
//	'nStrNumSymbolGroup', for UK Currency formatting.
//	New data values will be configured for the positive,
//	zero and negative number sign symbols as well as the
//	currency symbols.
//
//	The phrase 'Minus Outside' in the method name signals
//	that for negative values, the minus sign ('-') will
//	be to the left of, or outside, the currency symbol.
//
//		UK Example
//			- £123.45  Negative Value
//
//	To configure the minus sign outside the currency
//	symbol see method:
//
//		numStrNumberSymbolGroupMechanics.
//			setCurrencyDefaultsUKMinusOutside()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	NumStrNumberSymbolGroup passed as input parameter
//	'nStrNumSymbolGroup'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The default UK currency symbol is a leading pound
//	sign ('£').
//
//			UK Example-1:
//				£ 123.45  Positive Value
//
//		The negative number sign is set to a leading minus
//		sign ('-').
//
//			UK Example-2
//				-123.45  Negative Value
//
//	This method will configure the pound sign ('£')
//	such that any minus sign configured for negative
//	numeric values will be outside or the left of the
//	pound sign ('£').
//
//			UK Example-3
//				- £123.45  Negative Value
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		UK Example-4
//			£ 123.45
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		UK Example-5
//			£ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured using
//		the default parameters for UK Currency number
//		formatting standards.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setCurrencyDefaultsUKMinusOutside(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setCurrencyDefaultsUKMinusOutside()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = nStrNumSymbolGroup.currencySymbol.
		SetCurrencyDefaultsUKMinusOutside(
			ePrefix.XCpy(
				"nStrNumSymbolGroup.currencySymbol"))

	if err != nil {
		return err
	}

	return new(NumStrNumberSymbolSpec).
		SetSignedNumDefaultsUKMinus(
			&nStrNumSymbolGroup.positiveNumberSign,
			&nStrNumSymbolGroup.zeroNumberSign,
			&nStrNumSymbolGroup.negativeNumberSign,
			ePrefix)
}

//	setCurrencyDefaultsUSMinus
//
//	Receives an instance of NumStrNumberSymbolGroup,
//	deletes the pre-existing data values and proceeds to
//	reconfigure the instance according to the United
//	States (US) currency formatting standards.
//
//	Default values will be used to reconfigure the
//	NumStrNumberSymbolGroup input parameter,
//	'nStrNumSymbolGroup', for US Currency formatting.
//	New data values will be configured for the positive,
//	zero and negative number sign symbols as well as the
//	currency symbols.
//
//	The term 'Minus' in the method name signals that a
//	leading minus sign ('-') will be used to designate
//	negative numeric values.
//
//		US Example
//			 -123.45  Negative Value
//
//	To configure surrounding parentheses ('()') as the
//	negative number symbol see method:
//
//		numStrNumberSymbolGroupMechanics.
//			setCurrencyDefaultsUSParen()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	NumStrNumberSymbolGroup passed as input parameter
//	'nStrNumSymbolGroup'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The default US currency symbol is a leading Dollar
//	sign ('$').
//
//			US Example-1:
//				$ 123.45  Positive Value
//
//	The negative number sign is set to a leading minus
//	sign ('-').
//
//			US Example-2
//				-123.45  Negative Value
//
//	This method will configure the leading Dollar sign
//	('$') such that any minus sign configured for negative
//	numeric values will be inside or to the left of the
//	Dollar sign ('$').
//
//			US Example-3
//				$ -123.45  Negative Value
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-4
//			$ 123.45
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-5
//			$ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured using
//		the default parameters for US Currency number
//		formatting standards.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setCurrencyDefaultsUSMinus(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setCurrencyDefaultsUSMinus()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = nStrNumSymbolGroup.currencySymbol.
		SetCurrencyDefaultsUS(
			ePrefix.XCpy(
				"nStrNumSymbolGroup.currencySymbol"))

	if err != nil {
		return err
	}

	return new(NumStrNumberSymbolSpec).
		SetSignedNumDefaultsUSMinus(
			&nStrNumSymbolGroup.positiveNumberSign,
			&nStrNumSymbolGroup.zeroNumberSign,
			&nStrNumSymbolGroup.negativeNumberSign,
			ePrefix)
}

//	setCurrencyDefaultsUSParen
//
//	Receives an instance of NumStrNumberSymbolGroup,
//	deletes the pre-existing data values and proceeds to
//	reconfigure the instance according to the United
//	States (US) currency formatting standards.
//
//	Default values will be used to reconfigure the
//	NumStrNumberSymbolGroup input parameter,
//	'nStrNumSymbolGroup', for US Currency formatting.
//	New data values will be configured for the positive,
//	zero and negative number sign symbols as well as the
//	currency symbols.
//
//	The term 'Paren' in the method name signals that a
//	surrounding parentheses ('()') will be used to designate
//	negative numeric values.
//
//		US Example
//			 (123.45)  Negative Value
//
//	To configure minus signs ('-') as the
//	negative number symbol see method:
//
//		numStrNumberSymbolGroupMechanics.
//			setCurrencyDefaultsUSMinus()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete, overwrite and reset all
//	pre-existing data values in the instance of
//	NumStrNumberSymbolGroup passed as input parameter
//	'nStrNumSymbolGroup'.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.evertype.com/standards/euro/formats.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://www.thefinancials.com/Default.aspx?SubSectionID=curformat
//
//	https://www.codeproject.com/articles/78175/international-number-formats
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The default US currency symbol is a leading Dollar
//	sign ('$').
//
//			US Example-1:
//				$ 123.45  Positive Value
//
//	The negative number sign is set to surrounding
//	parentheses ('()').
//
//			US Example-2
//				(123.45)  Negative Value
//
//	This method will configure the leading Dollar sign
//	('$') such that any parentheses configured for
//	negative numeric values will be inside or to the left
//	of the Dollar sign ('$').
//
//			US Example-3
//				$ (123.45)  Negative Value
//
//	The positive number sign is implied for positive
//	numeric values. Therefore, the positive number sign
//	symbol is set to a blank or empty string ("").
//
//		US Example-4
//			$ 123.45
//
//	Zero numeric values have no number sign. Therefore,
//	the zero number symbol is set to a blank or empty
//	string ("").
//
//		US Example-5
//			$ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured using
//		the default parameters for US Currency number
//		formatting standards.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setCurrencyDefaultsUSParen(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setCurrencyDefaultsUSParen()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = nStrNumSymbolGroup.currencySymbol.
		SetCurrencyDefaultsUS(
			ePrefix.XCpy(
				"nStrNumSymbolGroup.currencySymbol"))

	if err != nil {
		return err
	}

	return new(NumStrNumberSymbolSpec).
		SetSignedNumDefaultsUSParen(
			&nStrNumSymbolGroup.positiveNumberSign,
			&nStrNumSymbolGroup.zeroNumberSign,
			&nStrNumSymbolGroup.negativeNumberSign,
			ePrefix)
}

//	setNumberSymbolRunes
//
//	Receives a pointer to an instance of
//	NumStrNumberSymbolGroup ('nStrNumSymbolGroup') and
//	reconfigures that instance using input parameters
//	for positive number symbols, zero number symbols,
//	negative number symbols and currency symbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reconfigure all
//	pre-existing data values in the
//	NumStrNumberSymbolGroup instance passed as input
//	parameter 'nStrNumSymbolGroup'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured using
//		the following input parameters for positive
//		number symbols, zero number symbols, negative
//		number symbols and currency symbols.
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
//	trailingZeroNumberSymbols	[]rune
//
//		A rune array containing the trailing zero
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setNumberSymbolRunes(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setNumberSymbolRunes()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	nStrNumSymNanobot := numStrNumberSymbolGroupNanobot{}

	nStrNumSymNanobot.empty(
		nStrNumSymbolGroup)

	if len(leadingPositiveNumberSymbols) > 0 ||
		len(trailingPositiveNumberSymbols) > 0 {
		err = nStrNumSymNanobot.setPositiveNumSignRunes(
			nStrNumSymbolGroup,
			leadingPositiveNumberSymbols,
			trailingPositiveNumberSymbols,
			positiveNumFieldSymPosition,
			ePrefix.XCpy(
				"nuStrNumSym<-PositiveNumSyms"))

		if err != nil {
			return err
		}

	} else {

		nStrNumSymbolGroup.positiveNumberSign.SetNOP()
	}

	if len(leadingNegativeNumberSymbols) > 0 ||
		len(trailingNegativeNumberSymbols) > 0 {

		err = nStrNumSymNanobot.setNegativeNumSignRunes(
			nStrNumSymbolGroup,
			leadingNegativeNumberSymbols,
			trailingNegativeNumberSymbols,
			negativeNumFieldSymPosition,
			ePrefix.XCpy(
				"nuStrNumSym<-NegativeNumSyms"))

		if err != nil {
			return err
		}

	} else {

		nStrNumSymbolGroup.negativeNumberSign.SetNOP()
	}

	if len(leadingZeroNumberSymbols) > 0 ||
		len(trailingZeroNumberSymbols) > 0 {

		err = nStrNumSymNanobot.setZeroNumSignRunes(
			nStrNumSymbolGroup,
			leadingZeroNumberSymbols,
			trailingZeroNumberSymbols,
			zeroNumFieldSymPosition,
			ePrefix.XCpy(
				"nuStrNumSym<-ZeroNumSyms"))

		if err != nil {
			return err
		}

	} else {

		nStrNumSymbolGroup.zeroNumberSign.SetNOP()
	}

	if len(leadingCurrencySymbols) > 0 ||
		len(trailingCurrencySymbols) > 0 {

		err = nStrNumSymNanobot.setCurrencySymbolRunes(
			nStrNumSymbolGroup,
			leadingCurrencySymbols,
			trailingCurrencySymbols,
			currencyNumFieldSymPosition,
			currencyNumSignRelPos,
			ePrefix.XCpy(
				"nuStrNumSym<-ZeroNumSyms"))
	}

	return err
}

//	setNumberSymbolSpecs
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setNumberSymbolSpecs(
	nStrNumSymbols *NumStrNumberSymbolGroup,
	positiveNumberSign NumStrNumberSymbolSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	currencySymbols NumStrNumberSymbolSpec,
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
			"setNumberSymbolSpecs()",
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

	if err != nil {
		return err
	}

	err = nStrNumSymbols.currencySymbol.CopyIn(
		&currencySymbols,
		ePrefix.XCpy(
			"nStrNumSymbols.currencySymbol<-"+
				"currencySymbols"))

	return err
}

//	setSignedNumDefaultsFrance
//
//	Receives an instance of NumStrNumberSymbolGroup,
//	deletes the pre-existing data values and proceeds to
//	reconfigure the instance according to signed number
//	formatting standards used in France.
//
//	Default values will be used to reconfigure the
//	NumStrNumberSymbolGroup input parameter,
//	'nStrNumSymbolGroup', for French Signed Number
//	formatting. New data values will be configured for
//	the positive, zero and negative number sign symbols.
//
//	No Currency Symbols will be included in the final
//	configuration.
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
// # Input Parameters
//
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured using
//		the default parameters for French Signed Number
//		formatting standards.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setSignedNumDefaultsFrance(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setSignedNumDefaultsFrance()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	nStrNumSymbolGroup.currencySymbol.Empty()

	return new(NumStrNumberSymbolSpec).
		SetSignedNumDefaultsFrance(
			&nStrNumSymbolGroup.positiveNumberSign,
			&nStrNumSymbolGroup.zeroNumberSign,
			&nStrNumSymbolGroup.negativeNumberSign,
			ePrefix)
}

//	setSimpleCurrencySymbolsConfig
//
//	Receives a pointer to an instance of
//	NumStrNumberSymbolGroup and proceeds to reconfigure
//	that instance for currency symbols.
//
//	Examples of Currency Symbol characters include such
//	symbols as the Dollar sign ('$'), Euro sign ('€') or
//	Pound sign ('£').
//
//	The NumStrNumberSymbolGroup instance is passed as
//	input parameter 'nStrNumSymbolGroup'. It is used
//	to configure Number Symbols required in converting
//	numeric values to formatted Number Strings.
//
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	All Number Symbol data values contained in the
//	'nStrNumSymbolGroup' instance will be deleted and
//	reconfigured as currency symbol specifications. These
//	specifications include Currency Symbols to be used
//	with positive numeric values, negative numeric values
//	and zero numeric values.
//
//	This method provides a simplified means for
//	reconfiguring input parameter 'nStrNumSymbolGroup',
//	with currency and number sign symbols. Using default
//	values, this method will generate Currency Symbols to
//	be used with positive numeric values, negative
//	numeric values and zero numeric values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and
//	reconfigure all data values contained in input
//	parameter, 'nStrNumSymbolGroup'.
//
//	All Number Symbol data values contained in
//	'nStrNumSymbolGroup' will be deleted and
//	reconfigured as currency symbol specifications.
//	These specifications include Currency Symbols to be
//	used with positive numeric values, negative numeric
//	values and zero numeric values.
//
// ----------------------------------------------------------------
//
// # Currency Defaults
//
//	Currency Symbol Padding Space:
//
//		If Currency Symbol string does NOT include a
//		blank space between the Currency Symbol and the
//		Number Sign Symbol, one blank space will be
//		added:
//
//		Examples:
//			European Number Strings:
//				"123.456- €"
//				"123.456 €"
//
//			US Number Strings:
//				"$ -123.456"
//				"$ 123.456"
//
//	Currency Negative Value Symbol Position:
//
//		Currency Symbol defaults to 'outside' the
//		minus sign.
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Negative Number Symbol:
//
//		The default Negative Number Symbol is the
//		minus sign ('-').
//
//		Examples:
//			European Number String: "123.456- €"
//			US Number String: "$ -123.456"
//
//	Positive Number Symbol:
//
//		No Positive Number Sign Symbol. Positive
//		values are implicit and assumed. Therefore,
//		no plus is formatted in the number string.
//
//		Positive Numeric Value Currency Examples:
//
//			European Number String: "123.456 €"
//			US Number String: "$ 123.456"
//
//	Zero Number Symbol:
//
//		No Number Sign Symbol. Technically a zero value
//		is neither positive nor negative. Therefore,
//		no plus is formatted in the number string.
//
//		Zero Numeric Value Currency Examples:
//
//			European Number String: "0.00 €"
//			US Number String: "$ 0.00"
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
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured as
//		currency symbol specifications. These
//		specifications include Currency Symbols to be used
//		with positive numeric values, negative numeric
//		values and zero numeric values.
//
//	currencySymbols				string
//
//		The symbol or symbols used to format currency.
//		This currency formatting will be used to
//		reconfigure the NumStrNumberSymbolGroup instance
//		passed by input parameter, 'nStrNumSymbols'.
//
//		Examples of Currency Symbol characters include
//		such symbols as the Dollar sign ('$'), Euro sign
//		('€') or Pound sign ('£').
//
//		If this string is empty, an error will be
//		returned.
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Symbols in a
//		Number String Format.
//
//		When set to 'true', 'nStrNumSymbols' will be
//		reconfigured with Number Symbols on the left side
//		of the numeric value. Such Number Symbols are
//		therefore configured as Leading Number Symbols.
//		This is the positioning format used in the US,
//		UK, Australia and most of Canada.
//
//		Example Number Strings:
//			"$ -123.456"
//
//		NOTE:	A space is automatically inserted after
//				the currency symbol, between the currency
//				symbol and the minus sign.
//
//		When set to 'false', the returned instance of
//		NumStrNumberSymbolGroup will configure Number
//		Symbols on the right side of the numeric value.
//		Such Number Symbols are therefore configured as
//		Trailing Number Symbols. This is the positioning
//		format used in France, Germany and many other
//		countries in the European Union.
//
//		Example Number Strings:
//			"123.456- €"
//
//		NOTE:	A space is automatically inserted after
//				the minus sign, between the minus sign
//				and the currency symbol.
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setSimpleCurrencySymbolsConfig(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setSimpleCurrencySymbolsConfig()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenCurrencySymbols := len(currencySymbols)

	if lenCurrencySymbols == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'currencySymbols' is invalid!\n"+
			"'currencySymbols' is an empty string with a zero\n"+
			"character length.\n",
			ePrefix.String())

		return err

	}

	if leadingNumSymbols == true &&
		currencySymbols[lenCurrencySymbols-1] != ' ' {

		currencySymbols += " "

	}

	if leadingNumSymbols == false &&
		currencySymbols[0] != ' ' {

		currencySymbols = " " + currencySymbols

	}

	new(numStrNumberSymbolGroupNanobot).empty(
		nStrNumSymbolGroup)

	nStrNumSymbolGroup.positiveNumberSign.SetNOP()

	nStrNumSymbolGroup.zeroNumberSign.SetNOP()

	var numSymStr string

	numSymStr = "-"

	if leadingNumSymbols {
		// Leading Number Symbols

		err = nStrNumSymbolGroup.negativeNumberSign.SetNumberSignLeadingSymbolRunes(
			[]rune(numSymStr),
			NumFieldSymPos.InsideNumField(),
			ePrefix.XCpy(
				"nStrNumSymbolGroup"))

		if err != nil {
			return err
		}

		err = nStrNumSymbolGroup.currencySymbol.
			SetCurrencyLeadingSymbolRunes(
				[]rune(currencySymbols),
				NumFieldSymPos.InsideNumField(),
				CurrNumSignRelPos.OutsideNumSign(),
				ePrefix.XCpy(
					"nStrNumSymbolGroup."+
						"currencySymbol<-currencySymbols"))

		if err != nil {
			return err
		}

	} else {
		// Trailing Number Symbols

		err = nStrNumSymbolGroup.negativeNumberSign.SetNumberSignTrailingSymbolRunes(
			[]rune(numSymStr),
			NumFieldSymPos.InsideNumField(),
			ePrefix.XCpy(
				"nStrNumSymbolGroup"))

		if err != nil {
			return err
		}

		err = nStrNumSymbolGroup.currencySymbol.
			SetCurrencyTrailingSymbolRunes(
				[]rune(currencySymbols),
				NumFieldSymPos.InsideNumField(),
				CurrNumSignRelPos.OutsideNumSign(),
				ePrefix.XCpy(
					"nStrNumSymbolGroup."+
						"currencySymbol<-currencySymbols"))

		if err != nil {
			return err
		}

	}

	return err
}

//	setSimpleNumberSignSymbolsConfig
//
//	Deletes and reconfigures the input parameter instance
//	of NumStrNumberSymbolGroup with a simple default
//	number sign symbol configuration.
//
//	Examples of number sign symbols include the plus
//	sign ('+') for positive numeric values and the minus sign
//	('-') or parentheses ('()') for negative numeric values.
//
//	Number sign symbols Do NOT contain currency symbols.
//
//	If currency symbol formatting is required,	see method:
//
//	numStrNumberSymbolGroupMechanics.setSimpleCurrencyConfig()
//
//	Type NumStrNumberSymbolGroup is used to configure
//	Number Symbols required in converting numeric values
//	to formatted Number Strings.
//
//	NumStrNumberSymbolGroup contains four instances of
//	type NumStrNumberSymbolSpec defining the Number
//	Symbols to be used with positive numeric values,
//	negative numeric values, zero numeric values and
//	currency values.
//
//	This method provides a simplified means for
//	reconfiguring input parameter 'nStrNumSymbolGroup',
//	with number sign symbols. Using default values, this
//	method will generate number sign symbols for positive
//	numeric values, zero numeric values and negative
//	numeric values.
//
//	Again, Number Sign Symbols DO NOT contain currency
//	symbols.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that this method will delete and
//	reconfigure all data values contained in the instance
//	of NumStrNumberSymbolGroup passed as input parameter
//	'nStrNumSymbolGroup'.
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
//	nStrNumSymbolGroup			*NumStrNumberSymbolGroup
//
//		A pointer to an instance of
//		NumStrNumberSymbolGroup.
//
//		All Number Symbol data values contained in this
//		object will be deleted and reconfigured as new
//		Number Sign Symbol specifications.
//
//		Examples of number sign symbols include the plus
//		sign ('+') for positive numeric values and the
//		minus sign ('-') or parentheses ('()') for
//		negative numeric values.
//
//	leadingNumSymbols			bool
//
//		Controls the positioning of Number Sign Symbols
//		in a Number String Format.
//
//		Examples of number sign symbols include the plus
//		sign ('+') for positive numeric values and the
//		minus sign ('-') or parentheses ('()') for
//		negative numeric values.
//
//		When set to 'true', the current instance of
//		NumStrNumberSymbolGroup will configure Number
//		Sign Symbols on the left side of the numeric
//		value. Such Number Symbols are therefore
//		configured as leading Number Symbols. This is the
//		positioning format used in the US, UK, Australia
//		and most of Canada.
//
//			Example Number Strings:
//				"-123.456"
//
//		When 'leadingNumSymbols' is set to 'false', the
//		current instance of NumStrNumberSymbolGroup will
//		configure Number Sign Symbols on the right side
//		of the numeric value. Such Number Sign Symbols
//		are therefore configured as trailing Number Sign
//		Symbols. This is the positioning format used in
//		France, Germany and many other countries in the
//		European Union.
//
//		Example Number Strings:
//			"123.456-"
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
func (nStrNumSymbolsGroupMech *numStrNumberSymbolGroupMechanics) setSimpleNumberSignSymbolsConfig(
	nStrNumSymbolGroup *NumStrNumberSymbolGroup,
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
			"setSimpleNumberSignSymbolsConfig()",
		"")

	if err != nil {
		return err
	}

	if nStrNumSymbolGroup == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumSymbolGroup' is invalid!\n"+
			"'nStrNumSymbolGroup' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrNumberSymbolGroupNanobot).empty(
		nStrNumSymbolGroup)

	nStrNumSymbolGroup.positiveNumberSign.SetNOP()

	nStrNumSymbolGroup.zeroNumberSign.SetNOP()

	var numSymStr string

	numSymStr = "-"

	if leadingNumSymbols {
		// Leading Number Sign Symbols

		err = nStrNumSymbolGroup.negativeNumberSign.SetNumberSignLeadingSymbolRunes(
			[]rune(numSymStr),
			NumFieldSymPos.InsideNumField(),
			ePrefix.XCpy(
				"nStrNumSymbolGroup"))

	} else {

		// Trailing Number Sign Symbols

		err = nStrNumSymbolGroup.negativeNumberSign.SetNumberSignTrailingSymbolRunes(
			[]rune(numSymStr),
			NumFieldSymPos.InsideNumField(),
			ePrefix.XCpy(
				"nStrNumSymbolGroup"))

	}

	return err
}

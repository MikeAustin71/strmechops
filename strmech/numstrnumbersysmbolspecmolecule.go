package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numStrNumberSymbolSpecMolecule - This type provides
// helper methods for NumStrNumberSymbolSpec
type numStrNumberSymbolSpecMolecule struct {
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
func (nStrNumSymSpecMolecule *numStrNumberSymbolSpecMolecule) copyNStrNumberSymbolSpec(
	destinationNumSymbolSpec *NumStrNumberSymbolSpec,
	sourceNumSymbolSpec *NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumSymSpecMolecule.lock == nil {
		nStrNumSymSpecMolecule.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMolecule.lock.Lock()

	defer nStrNumSymSpecMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMolecule."+
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

	new(numStrNumberSymbolSpecMolecule).empty(
		destinationNumSymbolSpec)

	isValid,
		_ := new(nStrNumberSymbolSpecAtom).
		testValidityNumStrNumberSymbolSpec(
			sourceNumSymbolSpec,
			nil)

	if !isValid {
		return err
	}

	err = destinationNumSymbolSpec.leadingNumberSymbols.
		CopyIn(
			&sourceNumSymbolSpec.leadingNumberSymbols,
			ePrefix.XCpy(
				"destinationNumSymbolSpec.leadingNumberSymbols<-"+
					"sourceNumSymbolSpec"))

	if err != nil {
		return err
	}

	destinationNumSymbolSpec.leadingNumberFieldSymbolPosition =
		sourceNumSymbolSpec.leadingNumberFieldSymbolPosition

	err = destinationNumSymbolSpec.trailingNumberSymbols.
		CopyIn(
			&sourceNumSymbolSpec.trailingNumberSymbols,
			ePrefix.XCpy(
				"destinationNumSymbolSpec.trailingNumberSymbols<-"+
					"sourceNumSymbolSpec"))

	destinationNumSymbolSpec.trailingNumberFieldSymbolPosition =
		sourceNumSymbolSpec.trailingNumberFieldSymbolPosition

	destinationNumSymbolSpec.currencyNumSignRelativePos =
		sourceNumSymbolSpec.currencyNumSignRelativePos

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
func (nStrNumSymSpecMolecule *numStrNumberSymbolSpecMolecule) empty(
	nStrNumSymbolSpec *NumStrNumberSymbolSpec) {

	if nStrNumSymSpecMolecule.lock == nil {
		nStrNumSymSpecMolecule.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMolecule.lock.Lock()

	defer nStrNumSymSpecMolecule.lock.Unlock()

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
//		(currencyNumSignRelativePos) will be deleted and
//		reset to the value specified by input parameter,
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
func (nStrNumSymSpecMolecule *numStrNumberSymbolSpecMolecule) setCurrencyNumSignRelPos(
	numSymbolSpec *NumStrNumberSymbolSpec,
	currencyNumSignRelPos CurrencyNumSignRelativePosition,
	errPrefDto *ePref.ErrPrefixDto) error {

	if nStrNumSymSpecMolecule.lock == nil {
		nStrNumSymSpecMolecule.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMolecule.lock.Lock()

	defer nStrNumSymSpecMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMolecule."+
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
			" currencyNumSignRelativePos String Value = %v\n"+
			"currencyNumSignRelativePos Integer Value = %v\n",
			ePrefix.String(),
			currencyNumSignRelPos.String(),
			currencyNumSignRelPos.XValueInt())

		return err
	}

	numSymbolSpec.currencyNumSignRelativePos =
		currencyNumSignRelPos

	return err
}

//	setLeadingCurrencySymbolSpec
//
//	Deletes and resets the data value of the Leading
//	Currency Symbol contained in an instance of
//	NumStrNumberSymbolSpec passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numSymbolSpec					*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		The Leading Number Symbol contained in this
//		instance will be deleted and reset to the value
//		specified by input parameter,
//		'leadingCurrencySymbols'.
//
//	leadingCurrencySymbols			[]rune
//
//		An array of runes specifying the currency
//		character or characters which will be copied to
//		the Leading Number Symbol contained in input
//		parameter, 'numSymbolSpec'.
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
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Currency
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: -123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " -123.45"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: -123.45
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
//			     	Numeric Value: -123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: -123.45
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
//	errPrefDto						*ePref.ErrPrefixDto
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
//	err								error
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
func (nStrNumSymSpecMolecule *numStrNumberSymbolSpecMolecule) setLeadingCurrencySymbolSpec(
	numSymbolSpec *NumStrNumberSymbolSpec,
	leadingNumberSymbol []rune,
	currencyInsideNumSymbol bool,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumSymSpecMolecule.lock == nil {
		nStrNumSymSpecMolecule.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMolecule.lock.Lock()

	defer nStrNumSymSpecMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMolecule."+
			"setLeadingCurrencySymbolSpec()",
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

	if leadingNumFieldSymPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leadingNumFieldSymPosition' is invalid!\n"+
			"'leadingNumFieldSymPosition' string value  = '%v'\n"+
			"'leadingNumFieldSymPosition' integer value = '%v'\n",
			ePrefix.String(),
			leadingNumFieldSymPosition.String(),
			leadingNumFieldSymPosition.XValueInt())

		return err

	}

	var currencyNumSignRelPos CurrencyNumSignRelativePosition

	if currencyInsideNumSymbol == true {

		currencyNumSignRelPos =
			CurrNumSignRelPos.InsideNumSign()

	} else {

		currencyNumSignRelPos =
			CurrNumSignRelPos.OutsideNumSign()
	}

	new(nStrNumberSymbolSpecAtom).emptyLeadingNStrNumSymbol(
		numSymbolSpec)

	if len(leadingNumberSymbol) > 0 {

		err = numSymbolSpec.leadingNumberSymbols.SetRuneArray(
			leadingNumberSymbol,
			ePrefix.XCpy(
				"numSymbolSpec.leadingCurrencySymbols"+
					"<-leadingCurrencySymbols"))

		if err != nil {
			return err
		}

		numSymbolSpec.leadingNumberFieldSymbolPosition =
			leadingNumFieldSymPosition

		numSymbolSpec.currencyNumSignRelativePos =
			currencyNumSignRelPos
	}

	return err
}

//	setLeadingNStrNumSymbolSpec
//
//	Deletes and resets the data value of the Leading
//	Number Sign Symbol contained in an instance of
//	NumStrNumberSymbolSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numSymbolSpec				*NumStrNumberSymbolSpec
//
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
//	err							error
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
func (nStrNumSymSpecMolecule *numStrNumberSymbolSpecMolecule) setLeadingNStrNumSymbolSpec(
	numSymbolSpec *NumStrNumberSymbolSpec,
	leadingNumberSymbol []rune,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumSymSpecMolecule.lock == nil {
		nStrNumSymSpecMolecule.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMolecule.lock.Lock()

	defer nStrNumSymSpecMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMolecule."+
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

		numSymbolSpec.leadingNumberFieldSymbolPosition =
			leadingNumFieldSymPosition

	}

	return err
}

//	setTrailingCurrencySymbolSpec
//
//	Deletes and resets the data value of the Trailing
//	Currency Number Symbol contained in an instance of
//	NumStrNumberSymbolSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numSymbolSpec					*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		The Trailing Number Symbol contained in this
//		instance will be deleted and reset to the value
//		specified by input parameter,
//		'trailingNumberSymbols'.
//
//	trailingCurrencySymbols			[]rune
//
//		An array of runes specifying the currency
//		character or characters which will be copied to
//		the Trailing Number Symbol contained in input
//		parameter, 'numSymbolSpec'.
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
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Currency
//		Number Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: -123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: -123.45
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
//			     	Numeric Value: -123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: -123.45
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
func (nStrNumSymSpecMolecule *numStrNumberSymbolSpecMolecule) setTrailingCurrencySymbolSpec(
	numSymbolSpec *NumStrNumberSymbolSpec,
	trailingNumberSymbol []rune,
	currencyInsideNumSymbol bool,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumSymSpecMolecule.lock == nil {
		nStrNumSymSpecMolecule.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMolecule.lock.Lock()

	defer nStrNumSymSpecMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMolecule."+
			"setTrailingCurrencySymbolSpec()",
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

	if trailingNumFieldSymPosition.XIsValid() == false {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'trailingNumFieldSymPosition' is invalid!\n"+
			"'trailingNumFieldSymPosition' string value  = '%v'\n"+
			"'trailingNumFieldSymPosition' integer value = '%v'\n",
			ePrefix.String(),
			trailingNumFieldSymPosition.String(),
			trailingNumFieldSymPosition.XValueInt())

		return err

	}

	var currencyNumSignRelPos CurrencyNumSignRelativePosition

	if currencyInsideNumSymbol == true {

		currencyNumSignRelPos =
			CurrNumSignRelPos.InsideNumSign()

	} else {

		currencyNumSignRelPos =
			CurrNumSignRelPos.OutsideNumSign()
	}

	new(nStrNumberSymbolSpecAtom).emptyTrailingNStrNumSymbol(
		numSymbolSpec)

	if len(trailingNumberSymbol) > 0 {

		err = numSymbolSpec.trailingNumberSymbols.SetRuneArray(
			trailingNumberSymbol,
			ePrefix.XCpy(
				"numSymbolSpec.trailingCurrencySymbols"+
					"<-trailingCurrencySymbols"))

		if err != nil {
			return err
		}

		numSymbolSpec.trailingNumberFieldSymbolPosition =
			trailingNumFieldSymPosition

		numSymbolSpec.currencyNumSignRelativePos =
			currencyNumSignRelPos
	}

	return err
}

// setTrailingNStrNumSymbolSpec
//
//	Deletes and resets the data value of the Trailing
//	Number Symbol contained in an instance of
//	NumStrNumberSymbolSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numSymbolSpec				*NumStrNumberSymbolSpec
//
//		A pointer to a NumStrNumberSymbolSpec instance.
//		The Trailing Number Symbol contained in this
//		instance will be deleted and reset to the value
//		specified by input parameter,
//		'trailingNumberSymbols'.
//
//	trailingNumberSymbols			[]rune
//
//		An array of runes specifying the number sign
//		character or characters which will be copied to
//		the Trailing Number Symbol contained in input
//		parameter, 'numSymbolSpec'.
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
//					Numeric Value: -123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:------>01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 10
//					Numeric Value: -123.45
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
//			     	Numeric Value: -123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:------>012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//					Numeric Value: -123.45
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
func (nStrNumSymSpecMolecule *numStrNumberSymbolSpecMolecule) setTrailingNStrNumSymbolSpec(
	numSymbolSpec *NumStrNumberSymbolSpec,
	trailingNumberSymbol []rune,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumSymSpecMolecule.lock == nil {
		nStrNumSymSpecMolecule.lock = new(sync.Mutex)
	}

	nStrNumSymSpecMolecule.lock.Lock()

	defer nStrNumSymSpecMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberSymbolSpecMolecule."+
			"setTrailingNStrNumSymbolSpec()",
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
		numSymbolSpec)

	if len(trailingNumberSymbol) > 0 {

		err = numSymbolSpec.trailingNumberSymbols.SetRuneArray(
			trailingNumberSymbol,
			ePrefix.XCpy(
				"numSymbolSpec.trailingNumberSymbols"+
					"<-trailingNumberSymbols"))

		if err != nil {
			return err
		}

		numSymbolSpec.trailingNumberFieldSymbolPosition =
			trailingNumFieldSymPosition

	}

	return err
}

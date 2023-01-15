package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

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

	return new(numStrNumberSymbolSpecNanobot).
		setTrailingCurrencySymbol(
			currencySymbols,
			[]rune{' ', '€'},
			NumFieldSymPos.InsideNumField(),
			CurrNumSignRelPos.OutsideNumSign(),
			ePrefix.XCpy(
				"currencySymbols<-Trailing Euro Sign"))
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

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			currencySymbols,
			[]rune{'£', ' '},
			NumFieldSymPos.InsideNumField(),
			CurrNumSignRelPos.InsideNumSign(),
			ePrefix.XCpy(
				"currencySymbols<-Leading Pound Sign"))
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

	return new(numStrNumberSymbolSpecNanobot).
		setLeadingCurrencySymbol(
			currencySymbols,
			[]rune{'$', ' '},
			NumFieldSymPos.InsideNumField(),
			CurrNumSignRelPos.OutsideNumSign(),
			ePrefix.XCpy(
				"currencySymbols<-Leading Dollar Sign"))
}

// setSignedNumSymbolsDefaultFrance
//
// Reconfigures three NumStrNumberSymbolSpec input
// parameters with default signed number symbols
// commonly applied in France.
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
// are therefore configured with French specifications
// for positive signed number symbols, zero value
// symbols, and negative signed number symbols.
//
// The positive signed number symbol is configured as
// empty or blank because under French formatting
// standards, positive number signs are implied and not
// specifically displayed. Therefore, no leading plus
// ('+') symbol is required.
//
// Likewise, the zero signed number symbol is also
// configured as empty or blank because under French
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
//		as empty or blank because under French formatting
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
//		empty or blank because under French formatting
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
//		French number string formatting standards.
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
func (nStrNumSymSpecMech *numStrNumberSymbolSpecMechanics) setSignedNumSymbolsDefaultFrance(
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
			"setSignedNumSymbolsDefaultFrance()",
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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

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

	numStrNumSymSpecNanobot := numStrNumberSymbolSpecMolecule{}

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

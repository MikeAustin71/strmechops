package strmech

import "sync"

// Lock lockCurrencyNumSignRelativePosition before
// accessing these 'maps'.

var mapCurrencyNumSignRelPosCodeToString = map[CurrencyNumSignRelativePosition]string{
	CurrencyNumSignRelativePosition(0): "None",
	CurrencyNumSignRelativePosition(1): "OutsideNumSign",
	CurrencyNumSignRelativePosition(2): "InsideNumSign",
}

var mapCurrencyNumSignRelPosStringToCode = map[string]CurrencyNumSignRelativePosition{
	"None":           CurrencyNumSignRelativePosition(0),
	"OutsideNumSign": CurrencyNumSignRelativePosition(1),
	"InsideNumSign":  CurrencyNumSignRelativePosition(2),
}

var mapCurrencyNumSignRelPosLwrCaseStringToCode = map[string]CurrencyNumSignRelativePosition{
	"none":           CurrencyNumSignRelativePosition(0),
	"outsidenumsign": CurrencyNumSignRelativePosition(1),
	"insidenumsign":  CurrencyNumSignRelativePosition(2),
}

// CurrencyNumSignRelativePosition
//
// An enumeration used to define the position of currency
// symbols relative to number sign symbols in number
// strings. Number strings are strings of text comprised
// of numeric digits formatted for the text display of
// numeric values.
//
// ----------------------------------------------------------------
//
// # BACKGROUND
//
// Often, these number strings contain number signs
// associated with the numeric values displayed as text.
//
// These number signs might include minus signs ('-') for
// negative values or plus signs ('+') for positive values.
// Some negative values are alternatively formatted with
// opening and closing parentheses ("()") which also
// qualify as number signs.
//
// When currency symbols and number signs are displayed
// in the same number string, the question arises as
// to which comes first, the currency symbol or the
// number sign.
//
// The CurrencyNumSignRelativePosition enumeration is
// used to specify the position of a currency symbol
// relative to a number string.
//
// Since the Go Programming Language does not directly
// support enumerations, the
// 'CurrencyNumSignRelativePosition' type has been
// adapted to function in a manner similar to classic
// enumerations.
//
// ----------------------------------------------------------------
//
// # USAGE
//
// CurrencyNumSignRelativePosition is declared as a type
// 'int' and includes two types of methods:
//
//	Enumeration Methods
//	      and
//	Utility Methods
//
// Enumeration methods have names which collectively
// represent an enumeration of different currency
// positions relative to number signs which may used in
// formatting currency symbols within number strings.
//
//	  Examples Of Enumeration Method Names:
//		CurrencyNumSignRelativePosition(0).None()
//		CurrencyNumSignRelativePosition(0).OutsideNumSign()
//		CurrencyNumSignRelativePosition(0).InsideNumSign()
//
// Enumeration methods return an integer value used to
// designate a specific Currency Symbol Position relative
// to a number sign within a number string.
//
// Utility methods make up the second type of method
// included in type CurrencyNumSignRelativePosition. These
// methods are NOT part of the enumeration but instead
// provide needed supporting services.
//
// All utility methods, with the sole exception of method
// String(), have names beginning with 'X' to separate
// them from standard enumeration methods.
//
//	  Examples:
//	    XIsValid()
//	    XParseString()
//	    XValue()
//	    XValueInt()
//
//	The utility method 'String()' supports the Stringer
//	Interface and is not part of the standard enumeration.
//
// For easy access to these enumeration values, use the
// global constant 'CurrNumSignRelPos'.
//
//	Example:
//		CurrNumSignRelPos.OutsideNumSign()
//
// Otherwise you will need to use the longer formal syntax.
//
//	Example:
//		CurrencyNumSignRelativePosition(0).OutsideNumSign()
//
// Depending on your editor, intellisense (a.k.a.
// intelligent code completion) may not list the
// CurrencyNumSignRelativePosition methods in
// alphabetical order.
//
// ----------------------------------------------------------------
//
// # ENUMERATION METHODS
//
// The CurrencyNumSignRelativePosition enumeration methods
// are described below:
//
// Method                     Integer
//
//	Name                       Value
//
// ------                     -------
//
//	None						 0
//
//	Signals that the CurrencyNumSignRelativePosition Type
//	is uninitialized. This is an error condition.
//
//	OutsideNumSign				 1
//
//	Signals that the currency symbol will be positioned
//	outside the number sign in a number string.
//
//	Examples:
//		"$ -123.45"
//		"123.45- €"
//
//	InsideNumSign				 2
//
//	Signals that the currency symbol will be positioned
//	inside the number sign in a number string.
//
//	Examples:
//		"- $123.45"
//		"123.45€ -"
type CurrencyNumSignRelativePosition int

var lockCurrencyNumSignRelativePosition sync.Mutex

// None
//
// Signals that the CurrencyNumSignRelativePosition Type
// is uninitialized.
//
// This is an error condition.
//
// This method is part of the standard enumeration.
func (currencyNumSingRelPos CurrencyNumSignRelativePosition) None() CurrencyNumSignRelativePosition {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	return CurrencyNumSignRelativePosition(0)
}

// OutsideNumSign
//
// Signals that the currency symbol will be positioned
// outside the number sign in a number string.
//
// Examples:
//
//	"$ -123.45"
//	"123.45- €"
//
// This is a valid option for Type
// CurrencyNumSignRelativePosition.
//
// This method is part of the standard enumeration.
func (currencyNumSingRelPos CurrencyNumSignRelativePosition) OutsideNumSign() CurrencyNumSignRelativePosition {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	return CurrencyNumSignRelativePosition(1)
}

// InsideNumSign
//
// Signals that the currency symbol will be positioned
// inside the number sign in a number string.
//
// Examples:
//
//	"- $123.45"
//	"123.45€ -"
//
// This is a valid option for Type
// CurrencyNumSignRelativePosition.
//
// This method is part of the standard enumeration.
func (currencyNumSingRelPos CurrencyNumSignRelativePosition) InsideNumSign() CurrencyNumSignRelativePosition {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	return CurrencyNumSignRelativePosition(2)
}

// String
//
// Returns a string with the name of the enumeration
// associated with this instance of
// 'CurrencyNumSignRelativePosition'.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Usage
//
//		Example-1
//		This example uses the formal complete syntax for
//	 invoking a CurrencyNumSignRelativePosition method.
//
//		t:= CurrencyNumSignRelativePosition(0).OutsideNumField()
//		str := t.String()
//
//		str is now equal to "OutsideNumSign"
//
//		Example-2
//		This example uses the abbreviated shorthand syntax
//		for invoking a CurrencyNumSignRelativePosition method.
//
//		t:= CurrNumSignRelPos.OutsideNumSign()
//		str := t.String()
//
//		str is now equal to "OutsideNumSign"
func (currencyNumSingRelPos CurrencyNumSignRelativePosition) String() string {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	result, ok :=
		mapCurrencyNumSignRelPosCodeToString[currencyNumSingRelPos]

	if !ok {
		return "Error: CurrencyNumSignRelativePosition code UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// NumberFieldSymbolPosition value is valid.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Usage
//
//		Example-1
//		This examples uses the formal complete syntax to invoke
//	 a CurrencyNumSignRelativePosition method.
//		currSymRelPos := CurrencyNumSignRelativePosition(0).InsideNumField()
//
//		isValid := currSymRelPos.XIsValid()
//			"isValid = true"
//
//		currSymRelPos = CurrencyNumSignRelativePosition(0).None()
//
//		isValid := currSymRelPos.XIsValid()
//			"isValid = false"
//
//		Example-2
//		This examples uses the abbreviated shorthand syntax
//		to invoke a CurrencyNumSignRelativePosition method.
//
//		currSymRelPos := CurrNumSignRelPos.InsideNumField()
//
//		isValid := currSymRelPos.XIsValid()
//			"isValid = true"
//
//		currSymRelPos = CurrNumSignRelPos.None()
//
//		isValid := currSymRelPos.XIsValid()
//			"isValid = false"
func (currencyNumSingRelPos CurrencyNumSignRelativePosition) XIsValid() bool {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	return new(CurrencyNumSignRelativePositionNanobot).
		isValidCurrNumSignRelPos(currencyNumSingRelPos)
}

// CurrNumSignRelPos - Public global constant of type
// CurrencyNumSignRelativePosition.
//
// This variable serves as an easier, shorthand technique for
// accessing CurrencyNumSignRelativePosition values.
//
// For easy access to these enumeration values, use the global
// variable CurrNumSignRelPos.
//
//	Example: CurrNumSignRelPos.InsideNumField()
//
// Otherwise you will need to use the formal syntax.
//
//	Example:
//		CurrencyNumSignRelativePosition(0).InsideNumSign()
//
// Usage:
//
//	CurrNumSignRelPos.None()
//	CurrNumSignRelPos.OutsideNumSign()
//	CurrNumSignRelPos.InsideNumSign()
const CurrNumSignRelPos = CurrencyNumSignRelativePosition(0)

// CurrencyNumSignRelativePositionNanobot - Provides helper methods for
// enumeration CurrencyNumSignRelativePosition.
type CurrencyNumSignRelativePositionNanobot struct {
	lock *sync.Mutex
}

// isValidCurrNumSignRelPos - Receives an instance of
// CurrencyNumSignRelativePosition and returns a boolean
// value signaling whether that
// CurrencyNumSignRelativePosition instance is valid.
//
// If the passed instance of
// CurrencyNumSignRelativePosition is valid, this method
// returns 'true'.
//
// Be advised, the enumeration value "None" is considered
// NOT VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of
// the valid CurrencyNumSignRelativePosition
// enumeration.
func (currencyNumSignRelPosNanobot *CurrencyNumSignRelativePositionNanobot) isValidCurrNumSignRelPos(
	currencyNumSignRelPosValue CurrencyNumSignRelativePosition) bool {

	if currencyNumSignRelPosNanobot.lock == nil {
		currencyNumSignRelPosNanobot.lock = new(sync.Mutex)
	}

	currencyNumSignRelPosNanobot.lock.Lock()

	defer currencyNumSignRelPosNanobot.lock.Unlock()

	if currencyNumSignRelPosValue < 1 ||
		currencyNumSignRelPosValue > 2 {

		return false
	}

	return true
}

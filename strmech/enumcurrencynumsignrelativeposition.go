package strmech

import (
	"fmt"
	"strings"
	"sync"
)

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
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) None() CurrencyNumSignRelativePosition {

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
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) OutsideNumSign() CurrencyNumSignRelativePosition {

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
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) InsideNumSign() CurrencyNumSignRelativePosition {

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
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) String() string {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	result, ok :=
		mapCurrencyNumSignRelPosCodeToString[currencyNumSignRelPos]

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
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) XIsValid() bool {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	return new(CurrencyNumSignRelativePositionNanobot).
		isValidCurrNumSignRelPos(currencyNumSignRelPos)
}

// XParseString
//
// Receives a string and attempts to match it with the
// string value of a supported enumeration. If
// successful, a new instance of
// CurrencyNumSignRelativePosition is returned set to the
// value of the associated enumeration.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
// valueString			string
//
//	A string which will be matched against the
//	enumeration string values. If 'valueString'
//	is equal to one of the enumeration names, this
//	method will proceed to successful completion
//	and return the correct enumeration value.
//
// caseSensitive		bool
//
//	If 'true' the search for enumeration names
//	will be case-sensitive and will require an
//	exact match. Therefore, 'insidenumsign' will
//	NOT	match the enumeration name, 'InsideNumSign'.
//
//	If 'false', a case-insensitive search is conducted
//	for the enumeration name. In this case,
//	'insidenumsign' will match the enumeration name
//	'InsideNumSign'.
//
// ------------------------------------------------------------------------
//
// # Return Values
//
// CurrencyNumSignRelativePosition
//
//	Upon successful completion, this method will return a
//	new instance of CurrencyNumSignRelativePosition set
//	to the value of the enumeration matched by the string
//	search performed on input parameter, 'valueString'.
//
// error
//
//	If this method completes successfully, the returned
//	error Type is set equal to 'nil'. If an error
//	condition is encountered, this method will return an
//	error type which encapsulates an appropriate error
//	message.
//
// ------------------------------------------------------------------------
//
// # Usage
//
// t, err :=
//
//		CurrencyNumSignRelativePosition(0).XParseString(
//		"InsideNumSign", true)
//
//	t is now equal to
//		CurrencyNumSignRelativePosition(0).InsideNumSign()
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) XParseString(
	valueString string,
	caseSensitive bool) (
	CurrencyNumSignRelativePosition,
	error) {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	ePrefix := "CurrencyNumSignRelativePosition.XParseString() "

	var ok bool
	var currSymNumSignRelPos CurrencyNumSignRelativePosition

	if caseSensitive {

		currSymNumSignRelPos,
			ok =
			mapCurrencyNumSignRelPosStringToCode[valueString]

		if !ok {
			return CurrencyNumSignRelativePosition(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CurrencyNumSignRelativePosition Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		currSymNumSignRelPos,
			ok = mapCurrencyNumSignRelPosLwrCaseStringToCode[strings.
			ToLower(valueString)]

		if !ok {
			return CurrencyNumSignRelativePosition(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CurrencyNumSignRelativePosition Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return currSymNumSignRelPos, nil
}

// XReturnNoneIfInvalid
//
// Provides a standardized value for invalid instances of
// enumeration CurrencyNumSignRelativePosition.
//
// If the current instance of
// CurrencyNumSignRelativePosition is invalid, this
// method will always return a value of
// CurrencyNumSignRelativePosition(0).None().
//
// # Background
//
// Enumeration CurrencyNumSignRelativePosition has an
// underlying type of integer (int). This means the type
// could conceivably be set to any integer value. This
// method ensures that all invalid
// CurrencyNumSignRelativePosition instances are
// consistently classified as 'None'
// (CurrencyNumSignRelativePosition(0).None()). Remember
// that 'None' is considered an invalid value.
//
// For example, assume that
// CurrencyNumSignRelativePosition was set to an integer
// value of -848972. Calling this method on a
// CurrencyNumSignRelativePosition with this invalid
// integer value will return an integer value of zero or
// the equivalent of
// CurrencyNumSignRelativePosition(0).None(). This
// conversion is useful in generating text strings for
// meaningful informational and error messages.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) XReturnNoneIfInvalid() CurrencyNumSignRelativePosition {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	isValid := new(CurrencyNumSignRelativePositionNanobot).
		isValidCurrNumSignRelPos(currencyNumSignRelPos)

	if !isValid {
		return CurrencyNumSignRelativePosition(0)
	}

	return currencyNumSignRelPos
}

// XValue
//
// This method returns the enumeration value of the
// current CurrencyNumSignRelativePosition instance.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) XValue() CurrencyNumSignRelativePosition {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	return currencyNumSignRelPos
}

// XValueInt
//
// This method returns the integer value of the current
// CurrencyNumSignRelativePosition instance.
//
// This is a standard utility method and is not part of
// the valid enumerations for this type.
func (currencyNumSignRelPos CurrencyNumSignRelativePosition) XValueInt() int {

	lockCurrencyNumSignRelativePosition.Lock()

	defer lockCurrencyNumSignRelativePosition.Unlock()

	return int(currencyNumSignRelPos)
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

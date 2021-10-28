package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockNumericSymbolClass before accessing these
// 'maps'.

var mapNumericSymClassCodeToString = map[NumericSymbolClass]string{
	NumericSymbolClass(0): "None",
	NumericSymbolClass(1): "NumberSign",
	NumericSymbolClass(2): "CurrencySign",
	NumericSymbolClass(3): "IntegerSeparator",
	NumericSymbolClass(4): "DecimalSeparator",
}

var mapNumericSymClassStringToCode = map[string]NumericSymbolClass{
	"None":             NumericSymbolClass(0),
	"NumberSign":       NumericSymbolClass(1),
	"CurrencySign":     NumericSymbolClass(2),
	"IntegerSeparator": NumericSymbolClass(3),
	"DecimalSeparator": NumericSymbolClass(4),
}

var mapNumericSymClassLwrCaseStringToCode = map[string]NumericSymbolClass{
	"none":             NumericSymbolClass(0),
	"numbersign":       NumericSymbolClass(1),
	"currencysign":     NumericSymbolClass(2),
	"integerseparator": NumericSymbolClass(3),
	"decimalseparator": NumericSymbolClass(4),
}

// NumericSymbolClass - An enumeration of numeric symbol
// classifications. Numeric symbols, usually found in number
// strings, are classified as Number Sign Symbols, Currency
// Symbols, Integer Separator Symbols (a.k.a. thousands separators),
// or Decimal Separator Symbols (a.k.a. integer/fractional
// separators).
//
// Since the Go Programming Language does not directly support
// enumerations, the NumericSymbolClass has been adapted to
// function in a manner similar to classic enumerations.
//
// NumericSignValueType is declared as a type 'int'. The method
// names effectively represent an enumeration of numeric symbol
// classifications. These methods are listed as follows:
//
//  None             (0) - Signals that 'NumericSymbolClass' has
//                         not been initialized and therefore has
//                         no value. This is an error condition.
//
//  NumberSign       (1) - Signals that the numeric symbol is
//                         classified as a number sign such as a
//                         plus ('+') or minus ('-').
//                         (Example: -842)
//
//  CurrencySign     (2) - Signals that the numeric symbol is
//                         classified as a currency sign such as
//                         the USA Dollar Sign ('$').
//                            (Example: $256.00)
//
//  IntegerSeparator (3) - Signals that the numeric symbol is
//                         classified as an integer separator or
//                         thousands separator. In the USA, integer
//                         digits are commonly separated into groups
//                         of three (thousands) separated by a comma.
//                         The Integer Separator in this case is a
//                         comma.
//                            (Example:  1,000,000,000)
//
//  DecimalSeparator (4) - Signals that the numeric symbol is
//                         classified as a decimal separator.
//                         Decimal separators are used to separate
//                         integer and fractional portions of a
//                         floating point number. In the USA,
//                         integer and fractional components of a
//                         floating point number are separated by a
//                         decimal point.
//                            (Example: 256.32)
//
// For easy access to these enumeration values, use the global
// constant NumSymClass. Example: NumSymClass.DecimalSeparator()
//
// Otherwise you will need to use the formal syntax.
// Example: NumericSymbolClass(0).DecimalSeparator()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumericSymbolClass methods in
// alphabetical order. Be advised that all NumericSymbolClass
// methods beginning with 'X', as well as the method 'String()',
// are utility methods, and NOT part of the enumeration values.
//
type NumericSymbolClass int

var lockNumericSymbolClass sync.Mutex

// None - Signals that NumericSymbolClass has not been set or
// initialized and therefore has no value. This is an error
// condition.
//
func (nSymbolClass NumericSymbolClass) None() NumericSymbolClass {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	return NumericSymbolClass(0)
}

// NumberSign - Signals that the numeric symbol is classified as a
// number sign such as a plus ('+') or minus ('-').
//   (Example: -842)
//
func (nSymbolClass NumericSymbolClass) NumberSign() NumericSymbolClass {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	return NumericSymbolClass(1)
}

// CurrencySign - Signals that the numeric symbol is classified as
// a currency sign such as the USA Dollar Sign ('$').
//  (Example: $256.00)
//
func (nSymbolClass NumericSymbolClass) CurrencySign() NumericSymbolClass {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	return NumericSymbolClass(2)
}

// IntegerSeparator - Signals that the numeric symbol is classified
// as an integer separator or thousands separator. In the USA,
// integer digits are commonly separated into groups of three
// (thousands) separated by a comma. The Integer Separator in this
// case is a comma.
//  (Example:  1,000,000,000)
//
func (nSymbolClass NumericSymbolClass) IntegerSeparator() NumericSymbolClass {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	return NumericSymbolClass(3)
}

// DecimalSeparator - Signals that the numeric symbol is classified
// as a decimal separator. Decimal separators are used to separate
// integer and fractional parts of a floating point number. In the
// USA, integer and fractional components of a floating point
// number are separated by a decimal point.
//  (Example: 256.32)
//
func (nSymbolClass NumericSymbolClass) DecimalSeparator() NumericSymbolClass {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	return NumericSymbolClass(4)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of NumericSymbolClass.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  t:= NumericSymbolClass(0).CurrencySign()
//  str := t.String()
//     str is now equal to 'CurrencySign'
//
func (nSymbolClass NumericSymbolClass) String() string {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	result, ok :=
		mapNumericSymClassCodeToString[nSymbolClass]

	if !ok {
		return "Error: NumericSymbolClass code UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// NumericSymbolClass value is valid.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  numSymClass := NumericSymbolClass(0).DecimalSeparator()
//
//  isValid := numSymClass.XIsValid() // isValid == true
//
//  numSymClass = NumericSignValueType(0).None()
//
//  isValid = numSymClass.XIsValid() // isValid == false
//
func (nSymbolClass NumericSymbolClass) XIsValid() bool {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	if nSymbolClass > 4 ||
		nSymbolClass < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of NumericSymbolClass is returned set to the
// value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// valueString          string
//     - A string which will be matched against the enumeration
//       string values. If 'valueString' is equal to one of the
//       enumeration names, this method will proceed to successful
//       completion and return the correct enumeration value.
//
// caseSensitive        bool
//     - If 'true' the search for enumeration names will be
//       case-sensitive and will require an exact match. Therefore,
//       "currencysign" will NOT match the enumeration name,
//       "CurrencySign".
//
//       If 'false' a case-insensitive search is conducted for the
//       enumeration name. In this case, 'currencysign' will match
//       the enumeration name 'CurrencySign'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// NumericSymbolClass
//     - Upon successful completion, this method will return a new
//       instance of NumericSymbolClass set to the value of the
//       enumeration matched by the string search performed on input
//       parameter, 'valueString'.
//
// error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is
//       encountered, this method will return an error type which
//       encapsulates an appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t, err := NumericSymbolClass(0).XParseString("IntegerSeparator", true)
//
//     t is now equal to NumericSymbolClass(0).IntegerSeparator()
//
func (nSymbolClass NumericSymbolClass) XParseString(
	valueString string,
	caseSensitive bool) (NumericSymbolClass, error) {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	ePrefix := "NumericSymbolClass.XParseString() "

	if len(valueString) < 4 {
		return NumericSymbolClass(0),
			fmt.Errorf(ePrefix+"\n"+
				"Input parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var strNumSymClass NumericSymbolClass

	if caseSensitive {

		strNumSymClass, ok =
			mapNumericSymClassStringToCode[valueString]

		if !ok {
			return NumericSymbolClass(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSymbolClass Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strNumSymClass, ok =
			mapNumericSymClassLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumericSymbolClass(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSymbolClass Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strNumSymClass, nil
}

// XValue - This method returns the enumeration value of the current
// NumericSymbolClass instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (nSymbolClass NumericSymbolClass) XValue() NumericSymbolClass {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	return nSymbolClass
}

// XValueInt - This method returns the integer value of the current
// NumericSymbolClass instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (nSymbolClass NumericSymbolClass) XValueInt() int {

	lockNumericSymbolClass.Lock()

	defer lockNumericSymbolClass.Unlock()

	return int(nSymbolClass)
}

// NumSymClass - public global constant of type
// NumericSymbolClass.
//
// This variable serves as an easier, shorthand technique for
// accessing NumericSymbolClass values.
//
// Usage:
// NumSymClass.None(),
// NumSymClass.NumberSign(),
// NumSymClass.CurrencySign(),
// NumSymClass.IntegerSeparator(),
// NumSymClass.DecimalSeparator(),
//
const NumSymClass = NumericSymbolClass(0)

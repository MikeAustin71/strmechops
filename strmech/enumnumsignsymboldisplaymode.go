package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockNumSignSymbolDisplayMode before accessing these
// 'maps'.

var mapNumSignSymbolDisplayModeCodeToString = map[NumSignSymbolDisplayMode]string{
	NumSignSymbolDisplayMode(0): "None",
	NumSignSymbolDisplayMode(1): "Explicit",
	NumSignSymbolDisplayMode(2): "Implicit",
}

var mapNumSignSymbolDisplayModeStringToCode = map[string]NumSignSymbolDisplayMode{
	"None":     NumSignSymbolDisplayMode(0),
	"Explicit": NumSignSymbolDisplayMode(1),
	"Implicit": NumSignSymbolDisplayMode(2),
}

var mapNumSignSymbolDisplayModeLwrCaseStringToCode = map[string]NumSignSymbolDisplayMode{
	"none":     NumSignSymbolDisplayMode(0),
	"explicit": NumSignSymbolDisplayMode(1),
	"implicit": NumSignSymbolDisplayMode(2),
}

// NumSignSymbolDisplayMode - An enumeration of display modes used
// describe the display of Number Signs in number strings.
//
// Number Sign Symbols refer to characters like the plus sign ('+')
// or minus sign ('-') used to define positive, negative or zero
// numeric values. Often these Number Sign Symbols are displayed
// as part of the numeric value (Examples: '-123' or '+123'). This
// visible display of a Number Sign Symbol in conjunction with a
// numeric value is referred to as an 'explicit' Number Sign Symbol
// display.
//
// In contrast, numeric values are often presented without Number
// Sign Symbols (Example: '123'). In this case the Number Sign
// Symbol is implied or 'implicit'. This means that the Number Sign
// Symbol is NOT displayed and the numeric value is assumed to be
// positive. For example the numeric value '123' has an assumed or
// implicit numeric sign of plus ('+') meaning the the value '123'
// is equal to the positive value '+123'.
//
// The NumSignSymbolDisplayMode enumeration is used to signal whether
// the number sign for a numeric value is 'explicit' or 'implicit'.
//
// Since the Go Programming Language does not directly support
// enumerations, the NumSignSymbolDisplayMode type has been adapted
// to function in a manner similar to classic enumerations.
// NumSignSymbolDisplayMode is declared as a type 'int'. The method
// names effectively represent an enumeration of number sign symbol
// position types. These methods are listed as follows:
//
//  None           (0) - Signals that the NumSignSymbolDisplayMode
//                       has not been initialized and therefore
//                       has no value. This is an error condition.
//
//  Explicit       (1) - Signals that the Number Sign Symbol will
//                       be explicitly displayed with the
//                       associated numeric value in a number
//                       string (Example: '+123').
//
//  Implicit       (2) - Signals that the Number Sign Symbol will
//                       NOT be explicitly displayed with the
//                       associated numeric value in a number
//                       string. In this mode, the Number Sign
//                       Symbol is assumed or implied.
//                       Example: '123' is assumed to be a positive
//                                value of '+123'. The plus sign
//                                ('+') is assumed or implied.
//
//
// For easy access to these enumeration values, use the global
// constant NSignSymDisplayMode.
//   Example: NSignSymDisplayMode.Explicit()
//
// Otherwise you will need to use the formal syntax.
//   Example: NumSignSymbolDisplayMode(0).Explicit()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumSignSymbolDisplayMode methods in
// alphabetical order. Be advised that all NumSignSymbolDisplayMode
// methods beginning with 'X', as well as the method 'String()',
// are utility methods, and NOT part of the enumeration values.
//
type NumSignSymbolDisplayMode int

var lockNumSignSymbolDisplayMode sync.Mutex

// None - Signals that the NumSignSymbolDisplayMode has not been
// initialized and therefore has no value.
//
// This is an error condition.
//
func (nSignSymDisMode NumSignSymbolDisplayMode) None() NumSignSymbolDisplayMode {

	lockNumSignSymbolDisplayMode.Lock()

	defer lockNumSignSymbolDisplayMode.Unlock()

	return NumSignSymbolDisplayMode(0)
}

// Explicit - Signals that the Number Sign Symbol will be
// explicitly displayed with the associated numeric value in a
// number string (Example: '+123').
//
func (nSignSymDisMode NumSignSymbolDisplayMode) Explicit() NumSignSymbolDisplayMode {

	lockNumSignSymbolDisplayMode.Lock()

	defer lockNumSignSymbolDisplayMode.Unlock()

	return NumSignSymbolDisplayMode(1)
}

// Implicit - Signals that the Number Sign Symbol will NOT be
// explicitly displayed with the associated numeric value in a
// number string. In this mode, the Number Sign Symbol is assumed
// or implied. Example: '123' is assumed to be a positive value
// of '+123'. The plus sign ('+') is assumed or implied.
//
// A value of '123' would have a Number Sign Symbol Display Mode
// of 'Implicit' and a positive numeric value of '+123'.
//
func (nSignSymDisMode NumSignSymbolDisplayMode) Implicit() NumSignSymbolDisplayMode {

	lockNumSignSymbolDisplayMode.Lock()

	defer lockNumSignSymbolDisplayMode.Unlock()

	return NumSignSymbolDisplayMode(2)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of NumSignSymbolDisplayMode.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  t:= NumSignSymbolDisplayMode(0).Implicit()
//  str := t.String()
//     str is now equal to 'Implicit'
//
func (nSignSymDisMode NumSignSymbolDisplayMode) String() string {

	lockNumSignSymbolDisplayMode.Lock()

	defer lockNumSignSymbolDisplayMode.Unlock()

	result, ok :=
		mapNumSignSymbolDisplayModeCodeToString[nSignSymDisMode]

	if !ok {
		return "Error: NumSignSymbolDisplayMode code UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// NumSignSymbolDisplayMode value is valid.
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
//  numSignSymDisMode := NumSignSymbolDisplayMode(0).Explicit()
//
//  isValid := numSignSymDisMode.XIsValid() // isValid == true
//
//  numSignSymDisMode = NumSignSymbolDisplayMode(0).None()
//
//  isValid = numSignSymDisMode.XIsValid() // isValid == false
//
func (nSignSymDisMode NumSignSymbolDisplayMode) XIsValid() bool {

	lockNumSignSymbolDisplayMode.Lock()

	defer lockNumSignSymbolDisplayMode.Unlock()

	if nSignSymDisMode > 2 ||
		nSignSymDisMode < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of NumSignSymbolDisplayMode is returned set to the
// value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  valueString   string - A string which will be matched against
//                         the enumeration string values. If
//                         'valueString' is equal to one of the
//                         enumeration names, this method will
//                         proceed to successful completion and
//                         return the correct enumeration value.
//
//
//  caseSensitive   bool - If 'true' the search for enumeration
//                         names will be case sensitive and will
//                         require an exact match. Therefore,
//                         'explicit' will NOT match the enumeration
//                         name, 'Explicit'.
//
//                         If 'false' a case insensitive search is
//                         conducted for the enumeration name. In
//                         this case, 'explicit' will match the
//                         enumeration name 'Explicit'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NumSignSymbolDisplayMode
//     - Upon successful completion, this method will return a new
//       instance of NumSignSymbolDisplayMode set to the value of
//       the enumeration matched by the string search performed on
//       input parameter, 'valueString'.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is
//       encountered, this method will return an error type which
//       encapsulates an appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
// t, err := NumSignSymbolDisplayMode(0).XParseString("Explicit", true)
//
//     t is now equal to NumSignSymbolDisplayMode(0).Explicit()
//
// t, err = NumSignSymbolDisplayMode(0).XParseString("explicit", false)
//
//     t is now equal to NumSignSymbolDisplayMode(0).Explicit()
//
func (nSignSymDisMode NumSignSymbolDisplayMode) XParseString(
	valueString string,
	caseSensitive bool) (NumSignSymbolDisplayMode, error) {

	lockNumSignSymbolDisplayMode.Lock()

	defer lockNumSignSymbolDisplayMode.Unlock()

	ePrefix := "NumSignSymbolDisplayMode.XParseString() "

	if len(valueString) < 4 {
		return NumSignSymbolDisplayMode(0),
			fmt.Errorf(ePrefix+"\n"+
				"Input parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var strNumSignSymDisMode NumSignSymbolDisplayMode

	if caseSensitive {

		strNumSignSymDisMode, ok =
			mapNumSignSymbolDisplayModeStringToCode[valueString]

		if !ok {
			return NumSignSymbolDisplayMode(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumSignSymbolDisplayMode Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strNumSignSymDisMode, ok =
			mapNumSignSymbolDisplayModeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumSignSymbolDisplayMode(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumSignSymbolDisplayMode Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strNumSignSymDisMode, nil
}

// XValue - This method returns the enumeration value of the
// current NumSignSymbolDisplayMode instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (nSignSymDisMode NumSignSymbolDisplayMode) XValue() NumSignSymbolDisplayMode {

	lockNumSignSymbolDisplayMode.Lock()

	defer lockNumSignSymbolDisplayMode.Unlock()

	return nSignSymDisMode
}

// XValueInt - This method returns the integer value of the current
// NumSignSymbolDisplayMode instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (nSignSymDisMode NumSignSymbolDisplayMode) XValueInt() int {

	lockNumSignSymbolDisplayMode.Lock()

	defer lockNumSignSymbolDisplayMode.Unlock()

	return int(nSignSymDisMode)
}

// NSignSymDisplayMode - public global constant of type
// NumSignSymbolDisplayMode.
//
// This variable serves as an easier, short hand technique for
// accessing NumSignSymbolDisplayMode values.
//
// For easy access to these enumeration values, use this global
// constant NSignSymDisplayMode.
//   Example: NSignSymDisplayMode.Explicit()
//
// Otherwise you will need to use the formal syntax.
//   Example: NumSignSymbolDisplayMode(0).Explicit()
//
// Usage:
// NSignSymDisplayMode.None(),
// NSignSymDisplayMode.Explicit(),
// NSignSymDisplayMode.Implicit(),
//
const NSignSymDisplayMode = NumSignSymbolDisplayMode(0)

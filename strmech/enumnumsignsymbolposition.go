package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockNumSignSymbolPosition before accessing these
// 'maps'.

var mapNumSignSymbolPosCodeToString = map[NumSignSymbolPosition]string{
	NumSignSymbolPosition(0): "None",
	NumSignSymbolPosition(1): "Before",
	NumSignSymbolPosition(2): "After",
	NumSignSymbolPosition(3): "BeforeAndAfter",
}

var mapNumSignSymbolPosStringToCode = map[string]NumSignSymbolPosition{
	"None":           NumSignSymbolPosition(0),
	"Before":         NumSignSymbolPosition(1),
	"After":          NumSignSymbolPosition(2),
	"BeforeAndAfter": NumSignSymbolPosition(3),
}

var mapNumSignSymbolPosLwrCaseStringToCode = map[string]NumSignSymbolPosition{
	"none":           NumSignSymbolPosition(0),
	"before":         NumSignSymbolPosition(1),
	"after":          NumSignSymbolPosition(2),
	"beforeandafter": NumSignSymbolPosition(3),
}

// NumSignSymbolPosition - An enumeration of type designations used
// to define the position of a number sign symbol within a number
// string. Number sign symbols describe numeric values contained
// in number strings as positive or negative values. Examples of
// number sign symbols are the plus sign ('+') and minus sign ('-').
//
// Within number strings, number sign symbols may be positioned
// before the numeric value, after the numeric value or before and
// after numeric values. An example of this later position
// combination is the USA case of parentheses used to define
// negative numbers. The opening and closing parenthesis characters
// ('()') are positioned on both sides of the numeric value
// "(25.32)".
//
// Since the Go Programming Language does not directly support
// enumerations, the NumSignSymbolPosition type has been adapted
// to function in a manner similar to classic enumerations.
// NumSignSymbolPosition is declared as a type 'int'. The method
// names effectively represent an enumeration of number sign symbol
// position types. These methods are listed as follows:
//
//  None             (0) - Signals that 'NumSignSymbolPosition' has
//                         not been initialized and therefore has no
//                         value. This is an error condition.
//
//  Before           (1) - Signals that the Numeric Symbol is
//                         positioned 'Before' the numeric value.
//
//  After            (2) - Signals that the Numeric Symbol is
//                         positioned 'After' the numeric value.
//
//  BeforeAndAfter   (3) - Signals that the Numeric Symbol is
//                         comprised of two symbols, one positioned
//                         'Before' the numeric value and one
//                         positioned 'After' the numeric value.
//
//
// For easy access to these enumeration values, use the global constant
// 'NumSymPos'. Example: NumSymPos.Before()
//
// Otherwise you will need to use the formal syntax.
// Example: NumSignSymbolPosition(0).Before()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumSignSymbolPosition methods in
// alphabetical order. Be advised that all NumSignSymbolPosition
// methods beginning with 'X', as well as the method 'String()',
// are utility methods, and NOT part of the enumeration values.
//
type NumSignSymbolPosition int

var lockNumSignSymbolPosition sync.Mutex

// None - Signals that 'NumSignSymbolPosition' has not been
// initialized and therefore contains no value. This is an
// error condition.
//
func (nSignSymPos NumSignSymbolPosition) None() NumSignSymbolPosition {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return NumSignSymbolPosition(0)
}

// Before - Signals that the Numeric Symbol is positioned 'Before'
// the numeric value.
//
func (nSignSymPos NumSignSymbolPosition) Before() NumSignSymbolPosition {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return NumSignSymbolPosition(1)
}

// After - Signals that the Numeric Symbol is positioned 'After'
// the numeric value.
//
func (nSignSymPos NumSignSymbolPosition) After() NumSignSymbolPosition {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return NumSignSymbolPosition(2)
}

// BeforeAndAfter - Signals that the Numeric Symbol is comprised of
// two symbols, one positioned 'Before' the numeric value and one
// positioned 'After' the numeric value.
//
func (nSignSymPos NumSignSymbolPosition) BeforeAndAfter() NumSignSymbolPosition {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return NumSignSymbolPosition(3)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of NumSignSymbolPosition.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  t:= NumSignSymbolPosition(0).After()
//  str := t.String()
//     str is now equal to 'After'
//
func (nSignSymPos NumSignSymbolPosition) String() string {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	result, ok :=
		mapNumSignSymbolPosCodeToString[nSignSymPos]

	if !ok {
		return "Error: NumSignSymbolPosition code UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// NumSignSymbolPosition value is valid.
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
//  numSignSymPos := NumSignSymbolPosition(0).Before()
//
//  isValid := numSignSymPos.XIsValid() // isValid == true
//
//  numSignSymPos = NumSignSymbolPosition(0).None)
//
//  isValid = numSignSymPos.XIsValid() // isValid == false
//
func (nSignSymPos NumSignSymbolPosition) XIsValid() bool {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	if nSignSymPos > 3 ||
		nSignSymPos < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of NumSignSymbolPosition is returned set to the
// value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
// valueString   string - A string which will be matched against the
//                        enumeration string values. If 'valueString'
//                        is equal to one of the enumeration names, this
//                        method will proceed to successful completion
//                        and return the correct enumeration value.
//
// caseSensitive   bool - If 'true' the search for enumeration names
//                        will be case sensitive and will require an
//                        exact match. Therefore, 'before' will NOT
//                        match the enumeration name, 'Before'.
//
//                        If 'false' a case insensitive search is conducted
//                        for the enumeration name. In this case, 'before'
//                        will match the enumeration name 'Before'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// NumSignSymbolPosition
//     - Upon successful completion, this method will return a new
//       instance of NumSignSymbolPosition set to the value of the
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
// t, err := NumSignSymbolPosition(0).XParseString("Before", true)
//
//     t is now equal to NumSignSymbolPosition(0).Before()
//
func (nSignSymPos NumSignSymbolPosition) XParseString(
	valueString string,
	caseSensitive bool) (NumSignSymbolPosition, error) {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	ePrefix := "NumSignSymbolPosition.XParseString() "

	if len(valueString) < 4 {
		return NumSignSymbolPosition(0),
			fmt.Errorf(ePrefix+"\n"+
				"Input parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var strNumSignSymPos NumSignSymbolPosition

	if caseSensitive {

		strNumSignSymPos, ok =
			mapNumSignSymbolPosStringToCode[valueString]

		if !ok {
			return NumSignSymbolPosition(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumSignSymbolPosition Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strNumSignSymPos, ok =
			mapNumSignSymbolPosLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumSignSymbolPosition(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumSignSymbolPosition Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strNumSignSymPos, nil
}

// XValue - This method returns the enumeration value of the current
// NumSignSymbolPosition instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (nSignSymPos NumSignSymbolPosition) XValue() NumSignSymbolPosition {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return nSignSymPos
}

// XValueInt - This method returns the integer value of the current
// NumSignSymbolPosition instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
//
func (nSignSymPos NumSignSymbolPosition) XValueInt() int {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return int(nSignSymPos)
}

// NumSymPos - public global constant of type
// NumSignSymbolPosition.
//
// This variable serves as an easier, short hand technique for
// accessing NumSignSymbolPosition values.
//
// Usage:
// NumSymPos.None(),
// NumSymPos.Before(),
// NumSymPos.After(),
// NumSymPos.BeforeAndAfter(),
//
const NumSymPos = NumSignSymbolPosition(0)

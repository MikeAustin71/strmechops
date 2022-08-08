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
//	None             (0) - Signals that 'NumSignSymbolPosition' has
//	                       not been initialized and therefore has no
//	                       value. This is an error condition.
//
//	Before           (1) - Signals that the Numeric Symbol is
//	                       positioned 'Before' the numeric value.
//
//	After            (2) - Signals that the Numeric Symbol is
//	                       positioned 'After' the numeric value.
//
//	BeforeAndAfter   (3) - Signals that the Numeric Symbol
//	                       comprises two symbols, one positioned
//	                       'Before' the numeric value and one
//	                       positioned 'After' the numeric value.
//
// For easy access to these enumeration values, use the global constant
// NumSignSymPos. Example: NumSignSymPos.Before()
//
// Otherwise you will need to use the formal syntax.
// Example: NumSignSymbolPosition(0).Before()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumSignSymbolPosition methods in
// alphabetical order. Be advised that all NumSignSymbolPosition
// methods beginning with 'X', as well as the method 'String()',
// are utility methods, and NOT part of the enumeration values.
type NumSignSymbolPosition int

var lockNumSignSymbolPosition sync.Mutex

// None - Signals that 'NumSignSymbolPosition' has not been
// initialized and therefore contains no value. This is an
// error condition.
func (nSignSymPos NumSignSymbolPosition) None() NumSignSymbolPosition {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return NumSignSymbolPosition(0)
}

// Before - Signals that the Numeric Symbol is positioned 'Before'
// the numeric value.
func (nSignSymPos NumSignSymbolPosition) Before() NumSignSymbolPosition {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return NumSignSymbolPosition(1)
}

// After - Signals that the Numeric Symbol is positioned 'After'
// the numeric value.
func (nSignSymPos NumSignSymbolPosition) After() NumSignSymbolPosition {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return NumSignSymbolPosition(2)
}

// BeforeAndAfter - Signals that the Numeric Symbol consists of
// two symbols, one positioned 'Before' the numeric value and one
// positioned 'After' the numeric value.
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
//	t:= NumSignSymbolPosition(0).After()
//	str := t.String()
//	   str is now equal to 'After'
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
//	numSignSymPos := NumSignSymbolPosition(0).Before()
//
//	isValid := numSignSymPos.XIsValid() // isValid == true
//
//	numSignSymPos = NumSignSymbolPosition(0).None()
//
//	isValid = numSignSymPos.XIsValid() // isValid == false
func (nSignSymPos NumSignSymbolPosition) XIsValid() bool {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return new(numSignSymbolPosNanobot).
		isValidNumSignSymbolPosition(nSignSymPos)
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
// # Input Parameters
//
// valueString   string - A string which will be matched against the
//
//	enumeration string values. If 'valueString'
//	is equal to one of the enumeration names, this
//	method will proceed to successful completion
//	and return the correct enumeration value.
//
// caseSensitive   bool - If 'true' the search for enumeration names
//
//	will be case-sensitive and will require an
//	exact match. Therefore, 'before' will NOT
//	match the enumeration name, 'Before'.
//
//	If 'false' a case-insensitive search is conducted
//	for the enumeration name. In this case, 'before'
//	will match the enumeration name 'Before'.
//
// ------------------------------------------------------------------------
//
// # Return Values
//
// NumSignSymbolPosition
//   - Upon successful completion, this method will return a new
//     instance of NumSignSymbolPosition set to the value of the
//     enumeration matched by the string search performed on input
//     parameter, 'valueString'.
//
// error
//   - If this method completes successfully, the returned error
//     Type is set equal to 'nil'. If an error condition is
//     encountered, this method will return an error type which
//     encapsulates an appropriate error message.
//
// ------------------------------------------------------------------------
//
// # Usage
//
// t, err := NumSignSymbolPosition(0).XParseString("Before", true)
//
//	t is now equal to NumSignSymbolPosition(0).Before()
//
// t, err := NumSignSymbolPosition(0).XParseString("before", false)
//
//	t is now equal to NumSignSymbolPosition(0).Before()
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

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration NumSignSymbolPosition.
//
// If the current instance of NumSignSymbolPosition is invalid,
// this method will always return a value of
// NumSignSymbolPosition(0).None().
//
// # Background
//
// Enumeration NumSignSymbolPosition has an underlying type of
// integer (int). This means the type could conceivably be set to
// any integer value. This method ensures that all invalid
// NumSignSymbolPosition instances are consistently classified as
// 'None' (NumSignSymbolPosition(0).None()). Remember that 'None'
// is considered an invalid value.
//
// For example, assume that NumSignSymbolPosition was set to an
// invalid integer value of -848972. Calling this method on a
// NumSignSymbolPosition with this invalid integer value will
// return an integer value of zero or the equivalent of
// NumSignSymbolPosition(0).None(). This conversion is useful in
// generating text strings for meaningful informational and error
// messages.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (nSignSymPos NumSignSymbolPosition) XReturnNoneIfInvalid() NumSignSymbolPosition {
	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	isValid := new(numSignSymbolPosNanobot).
		isValidNumSignSymbolPosition(nSignSymPos)

	if !isValid {
		return NumSignSymbolPosition(0)
	}

	return nSignSymPos
}

// XValue - This method returns the enumeration value of the current
// NumSignSymbolPosition instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
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
func (nSignSymPos NumSignSymbolPosition) XValueInt() int {

	lockNumSignSymbolPosition.Lock()

	defer lockNumSignSymbolPosition.Unlock()

	return int(nSignSymPos)
}

// NumSignSymPos - public global constant of type
// NumSignSymbolPosition.
//
// This variable serves as an easier, shorthand technique for
// accessing NumSignSymbolPosition values.
//
// Usage:
// NumSignSymPos.None(),
// NumSignSymPos.Before(),
// NumSignSymPos.After(),
// NumSignSymPos.BeforeAndAfter(),
const NumSignSymPos = NumSignSymbolPosition(0)

// numSignSymbolPosNanobot - Provides helper methods for
// enumeration NumSignSymbolPosition.
type numSignSymbolPosNanobot struct {
	lock *sync.Mutex
}

// isValidNumSignSymbolPosition - Receives an instance of
// NumSignSymbolPosition and returns a boolean value signaling
// whether that NumSignSymbolPosition instance is valid.
//
// If the passed instance of NumSignSymbolPosition is valid, this
// method returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// NumSignSymbolPosition enumeration.
func (textFieldNanobot *numSignSymbolPosNanobot) isValidNumSignSymbolPosition(
	numSignSymbolPos NumSignSymbolPosition) bool {

	if textFieldNanobot.lock == nil {
		textFieldNanobot.lock = new(sync.Mutex)
	}

	textFieldNanobot.lock.Lock()

	defer textFieldNanobot.lock.Unlock()

	if numSignSymbolPos < 1 ||
		numSignSymbolPos > 3 {

		return false
	}

	return true
}

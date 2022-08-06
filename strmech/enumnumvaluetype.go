package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Do NOT access these maps without first getting
// the lock on 'lockNumericValueType'.

var mNumericValueTypeCodeToString = map[NumericValueType]string{
	NumericValueType(0): "None",
	NumericValueType(1): "FloatingPoint",
	NumericValueType(2): "Integer",
}

var mNumericValueTypeStringToCode = map[string]NumericValueType{
	"None":           NumericValueType(0),
	"FloatingPoint":  NumericValueType(1),
	"Floating Point": NumericValueType(1),
	"Float":          NumericValueType(1),
	"Integer":        NumericValueType(2),
	"Int":            NumericValueType(2),
}

var mNumericValueTypeLwrCaseStringToCode = map[string]NumericValueType{
	"none":           NumericValueType(0),
	"floatingpoint":  NumericValueType(1),
	"floating point": NumericValueType(1),
	"float":          NumericValueType(1),
	"integer":        NumericValueType(2),
	"int":            NumericValueType(2),
}

// NumericValueType - The 'Numeric Value Type' is an enumeration of
// type codes for classification of numeric values.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'NumericValueType' type has been adapted to
// function in a manner similar to classic enumerations.
//
// 'NumericValueType' is declared as a type 'int'. The method
// names effectively represent an enumeration of numeric value
// types. These methods are listed as
// follows:
//
// Method             Integer
// Name                Value
// ------             -------
//
// None                 (0)
//   - Signals that the Numeric Value type is empty and not
//     initialized. This is an error condition.
//
// FloatingPoint        (1)
//
//   - Designates the numeric value as a decimal floating point
//     number. These numeric value consists of both an integer
//     component and a fractional component separated by a
//     decimal separator character. In the USA, the decimal
//     separator character is a period or decimal point ('.').
//     Floating point numeric values will always include integer
//     digits to the left of the decimal separator and fractional
//     digits to the right of the decimal separator.
//
//     Reference:
//     https://en.wikipedia.org/wiki/Decimal_floating_point
//     https://www.differencebetween.info/difference-between-integer-and-float
//
//     Examples:
//     123.456
//     0.5
//     2.12345678
//     1,765,432.12
//
// Integer              (2)
//
//   - Designates the numeric value as a whole number containing
//     integer digits and no fractional digits.
//
//     Reference:
//     https://en.wikipedia.org/wiki/Integer
//     https://www.differencebetween.info/difference-between-integer-and-float
//
//     Examples:
//     0
//     123
//     9,875,432
//
// For easy access to these enumeration values, use the global
// constant 'NumValType'.
//
//	Example: NumValType.Integer()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumericValueType(0).Integer()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumericValueType methods in
// alphabetical order.
//
// Be advised that all 'NumericValueType' methods beginning with
// 'X', as well as the method 'String()', are utility methods and
// not part of the enumeration.
type NumericValueType int

var lockNumericValueType sync.Mutex

// None - Signals that the NumericValueType specification is empty
// or uninitialized. This is an error condition
//
// This method is part of the standard enumeration.
func (numValType NumericValueType) None() NumericValueType {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	return NumericValueType(0)
}

// FloatingPoint - Designates the numeric value as a decimal
// floating point number. These numeric value consists of both an
// integer component and a fractional component separated by a
// decimal separator character. In the USA, the decimal separator
// character is a period or decimal point ('.').
//
// Floating point numeric values will always include integer digits
// to the left of the decimal separator and fractional digits to
// the right of the decimal separator.
//
//	Examples:
//	         123.456
//	           0.5
//	           2.12345678
//	   1,765,432.12
//
// Reference:
//
//	https://en.wikipedia.org/wiki/Decimal_floating_point
//	https://www.differencebetween.info/difference-between-integer-and-float
func (numValType NumericValueType) FloatingPoint() NumericValueType {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	return NumericValueType(1)
}

// Integer - Designates the numeric value as a whole number
// containing integer digits, but no fractional digits.
//
//	Examples:
//	        0
//	      123
//	9,875,432
//
// Reference:
//
//	https://en.wikipedia.org/wiki/Integer
//	https://www.differencebetween.info/difference-between-integer-and-float
func (numValType NumericValueType) Integer() NumericValueType {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	return NumericValueType(2)
}

// String - Returns a string with the name of the enumeration
// associated with this current instance of 'NumericValueType'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Usage
//
// t:= NumericValueType(0).Integer()
// str := t.String()
//
//	str is now equal to 'Integer'
func (numValType NumericValueType) String() string {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	result, ok := mNumericValueTypeCodeToString[numValType]

	if !ok {
		return "Error: Numeric Value Type Specification UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// Numeric Value Type Specification (NumericValueType) is valid.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// Usage
//
//	nStrValueSpec := NumericValueType(0).FloatingPoint()
//
//	isValid := nStrValueSpec.XIsValid()
//
//	In this case the boolean value of 'isValid' is 'true'.
//
//	Be advised, the value NumericValueType(0).None() is
//	classified as an 'invalid' value.
func (numValType NumericValueType) XIsValid() bool {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	return new(numericValueTypeNanobot).
		isValidNumericValueType(
			numValType)
}

// XParseString - Receives a string and attempts to match it with the
// string value of a supported enumeration. If successful, a new
// instance of NumericValueType is returned set to the value of the
// associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
// valueString   string
//   - A string which will be matched against the enumeration string
//     values. If 'valueString' is equal to one of the enumeration
//     names, this method will proceed to successful completion and
//     return the correct enumeration value.
//
// caseSensitive   bool
//
//   - If 'true' the search for enumeration names will be
//     case-sensitive and will require an exact match. Therefore,
//     'integer' will NOT match the enumeration name, 'Integer'.
//
//     A case-sensitive search will match any of the following strings:
//     "None"
//     "FloatingPoint"
//     "Floating Point"
//     "Float"
//     "Integer"
//     "Int"
//
//     If 'false', a case-insensitive search is conducted for the
//     enumeration name. In this example, 'integer' WILL MATCH
//     the enumeration name, 'Integer'.
//
//     A case-insensitive search will match any of the following
//     lower case names:
//     "none"
//     "floatingpoint"
//     "floating point"
//     "float"
//     "integer"
//     "int"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NumericValueType
//	   - Upon successful completion, this method will return a new
//	     instance of NumericValueType set to the value of the
//	     enumeration matched by the string search performed on
//	     input parameter, 'valueString'.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'. If an error condition is encountered,
//	     this method will return an error type which encapsulates an
//	     appropriate error message.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t, err := NumericValueType(0).XParseString("Integer", true)
//
//	   t is now equal to NumericValueType(0).Integer()
func (numValType NumericValueType) XParseString(
	valueString string,
	caseSensitive bool) (NumericValueType, error) {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	ePrefix := "NumericValueType.XParseString() "

	if len(valueString) < 3 {
		return NumericValueType(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '3'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var numValueType NumericValueType

	if caseSensitive {

		numValueType, ok = mNumericValueTypeStringToCode[valueString]

		if !ok {
			return NumericValueType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericValueType Specification.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		numValueType, ok = mNumericValueTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumericValueType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericValueType Specification.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return numValueType, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration NumericValueType.
//
// If the current instance of NumericValueType is invalid, this
// method will always return a value of NumericValueType(0).None().
//
// # Background
//
// Enumeration NumericValueType has an underlying type of integer
// (int). This means the type could conceivably be set to any
// integer value. This method ensures that all invalid
// NumericValueType instances are consistently classified as 'None'
// (NumericValueType(0).None()). Remember that 'None' is considered
// an invalid value.
//
// For example, assume that NumericValueType was set to an integer
// value of -848972. Calling this method on a NumericValueType with
// this invalid integer value will return an integer value of zero
// or the equivalent of NumericValueType(0).None(). This conversion
// is useful in generating text strings for meaningful
// informational and error messages.
func (numValType NumericValueType) XReturnNoneIfInvalid() NumericValueType {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	isValid := new(numericValueTypeNanobot).
		isValidNumericValueType(numValType)

	if !isValid {
		return NumericValueType(0).None()
	}

	return numValType
}

// XValue - This method returns the enumeration value of the current
// NumericValueType instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (numValType NumericValueType) XValue() NumericValueType {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	return numValType
}

// XValueInt - This method returns the integer value of the current
// NumericValueType.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (numValType NumericValueType) XValueInt() int {

	lockNumericValueType.Lock()

	defer lockNumericValueType.Unlock()

	return int(numValType)
}

// NumValType - public global constant of type NumericValueType.
//
// This variable serves as an easier, shorthand technique for
// accessing NumericValueType values.
//
// For easy access to these enumeration values, use the
// global variable NumValType.
//
//	Example: NumValType.FloatingPoint()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumericValueType(0).FloatingPoint()
//
// Usage:
//
//	NumValType.None()
//	NumValType.FloatingPoint()
//	NumValType.Integer()
const NumValType = NumericValueType(0)

// numericValueTypeNanobot - Provides helper methods for
// enumeration NumericValueType.
type numericValueTypeNanobot struct {
	lock *sync.Mutex
}

// isValidNumericValueType - Receives an instance of
// NumericValueType and returns a boolean value signaling whether
// that NumericValueType instance is valid.
//
// If the passed instance of NumericValueType is valid, this method
// returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// NumericValueType enumeration.
func (enumNumValueTypeNanobot *numericValueTypeNanobot) isValidNumericValueType(
	enumNumericValType NumericValueType) bool {

	if enumNumValueTypeNanobot.lock == nil {
		enumNumValueTypeNanobot.lock = new(sync.Mutex)
	}

	enumNumValueTypeNanobot.lock.Lock()

	defer enumNumValueTypeNanobot.lock.Unlock()

	if enumNumericValType < 1 ||
		enumNumericValType > 2 {
		return false
	}

	return true
}

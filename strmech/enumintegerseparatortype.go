package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Do NOT access these maps without first getting
// the lock on 'lockIntegerSeparatorTypeCode'.

var mIntegerSeparatorTypeCodeToString = map[IntegerSeparatorType]string{
	IntegerSeparatorType(0): "None",
	IntegerSeparatorType(1): "Thousands",
	IntegerSeparatorType(2): "IndiaNumbering",
	IntegerSeparatorType(3): "ChineseNumbering",
}

var mIntegerSeparatorTypeStringToCode = map[string]IntegerSeparatorType{
	"None":             IntegerSeparatorType(0),
	"Thousands":        IntegerSeparatorType(1),
	"IndiaNumbering":   IntegerSeparatorType(2),
	"ChineseNumbering": IntegerSeparatorType(3),
}

var mIntegerSeparatorTypeLwrCaseStringToCode = map[string]IntegerSeparatorType{
	"none":             IntegerSeparatorType(0),
	"thousands":        IntegerSeparatorType(1),
	"indianumbering":   IntegerSeparatorType(2),
	"chinesenumbering": IntegerSeparatorType(3),
}

// IntegerSeparatorType - The Integer Separator Type enumeration
// is used to specify the type of integer separation in creating
// groups of integer digits for number string displays.
//
// The most common type of integer separation is the 'thousands'
// separation.
//
//	Example: 1,000,000,000
//
// In this example, integer numeric digits are separated by commas
// into groups of three digits or thousands.
//
// In most western countries, integer digits to the left of
// the decimal separator (a.k.a. decimal point) are separated
// into groups of three digits representing a grouping of
// 'thousands' like this: '1,000,000,000'.
//
// In some countries and cultures other integer groupings are
// used. In India, for example, a number might be formatted
// like this: '6,78,90,00,00,00,00,000'.
//
// Chinese Numerals have an integer grouping value of four
// and are formatted like this: '12,3456,7890,2345'.
//
// The 'IntegerSeparatorType' designates the type of integer
// separation to be used when displaying number strings.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'IntegerSeparatorType' type has been
// adapted to function in a manner similar to classic
// enumerations.
//
// 'IntegerSeparatorType' is declared as a type 'int'. The
// method names effectively represent an enumeration of
// integer separation display specifications. These methods
// are listed as follows:
//
// Method             Integer
// Name                Value
// ------             -------
//
// None                 (0)
//   - Signals that Integer Separation is undefined and will
//     not be applied to the text presentation of a number
//     string.
//     Example: 6789000000000000
//
// Thousands            (1)
//   - Signals that Integer Separation will be applied.
//     Integer digits will be separated into groups of three
//     digits.
//     Example: 6,789,000,000,000,000
//
// IndiaNumbering       (2)
//   - Signals that India Number integer separation will
//     be applied. The India Numbering system groups
//     integers for text display like this:
//     Example: 6,78,90,00,00,00,00,000
//
// ChineseNumbering     (3)
//   - Signals that Chinese Number integer separation will
//     be applied. The Chinese Numbering System separates
//     integers into groups of 4-numeric digits.
//     Example: 6789,0000,0000,0000
//
// ----------------------------------------------------------------
//
// For easy access to these enumeration values, use the global
// constant 'IntSeparatorType'.
//
//	Example: IntSeparatorType.Thousands()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: IntegerSeparatorType(0).Thousands()
//
// Depending on your editor, intellisense (a.k.a. intelligent
// code completion) may not list the IntegerSeparatorType
// methods in alphabetical order.
//
// Be advised that all 'IntegerSeparatorType' methods beginning
// with 'X', as well as the method 'String()', are utility
// methods and not part of the enumeration.
type IntegerSeparatorType int

var lockIntegerSeparatorTypeCode sync.Mutex

// None - Signals that Integer Separation is undefined and will
// not be applied to the text presentation of a number string.
//
//	Example: 6789000000000000
//
// None is considered a valid choice for integer separation.
//
// This method is part of the standard enumeration.
func (intSeparatorType IntegerSeparatorType) None() IntegerSeparatorType {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	return IntegerSeparatorType(0)
}

// Thousands - Signals that Integer Separation will be applied.
// Integer digits will be separated into groups of three digits.
//
//	Example: 6,789,000,000,000,000
//
// This method is part of the standard enumeration.
func (intSeparatorType IntegerSeparatorType) Thousands() IntegerSeparatorType {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	return IntegerSeparatorType(1)
}

// IndiaNumbering - Signals that India Number integer separation
// will be applied. The India Numbering system groups integers
// for text display like this:
//
//	Example: 6,78,90,00,00,00,00,000
//
// This method is part of the standard enumeration.
func (intSeparatorType IntegerSeparatorType) IndiaNumbering() IntegerSeparatorType {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	return IntegerSeparatorType(2)
}

// ChineseNumbering - Signals that Chinese Number integer
// separation will be applied. The Chinese Numbering System
// separates integers into groups of 4-numeric digits.
//
//	Example: 6789,0000,0000,0000
//
// This method is part of the standard enumeration.
func (intSeparatorType IntegerSeparatorType) ChineseNumbering() IntegerSeparatorType {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	return IntegerSeparatorType(3)
}

// String - Returns a string with the name of the enumeration
// associated with this current instance of
// 'IntegerSeparatorType'.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Usage
//
// t:= IntegerSeparatorType(0).Thousands()
// str := t.String()
//
//	str is now equal to 'Thousands'
func (intSeparatorType IntegerSeparatorType) String() string {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	result, ok :=
		mIntegerSeparatorTypeCodeToString[intSeparatorType]

	if !ok {
		return "Error: Integer Separator Type INVALID!"

	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// IntegerSeparatorType value is valid.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	 intSepType :=
//				IntegerSeparatorType(0).Thousands()
//
//	 isValid := intSepType.XIsValid() // isValid == true
//
//	 intSepType = IntegerSeparatorType(-99)
//
//	 isValid = intSepType.XIsValid() // isValid == false
func (intSeparatorType IntegerSeparatorType) XIsValid() bool {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	return new(integerSeparatorTypeNanobot).
		isValidIntegerSeparatorType(intSeparatorType)
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of IntegerSeparatorType is returned set to the value
// of the associated enumeration.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
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
//     'thousands' will NOT match the enumeration name,
//     'Thousands'.
//
//     A case-sensitive search will match any of the following
//     strings:
//     "None"
//     "Thousands"
//     "IndiaNumbering"
//     "ChineseNumbering"
//
//     If 'false', a case-insensitive search is conducted for the
//     enumeration name. In this example, 'Thousands'  WILL MATCH
//     the enumeration name, 'thousands'.
//
//     A case-insensitive search will match any of the following
//     lower case names:
//     "none"
//     "thousands"
//     "indianumbering"
//     "chinesenumbering"
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	IntegerSeparatorType
//	   - Upon successful completion, this method will return a new
//	     instance of IntegerSeparatorType set to the value of the
//	     enumeration matched by the string search performed on
//	     input parameter, 'valueString'.
//
//	error
//	   - If this method completes successfully, the returned error
//	     Type is set equal to 'nil'. If an error condition is
//	     encountered, this method will return an error type which
//	     encapsulates an appropriate error message.
//
// ----------------------------------------------------------------
//
// Usage
//
//	t, err := IntegerSeparatorType(0).
//	             XParseString("Thousands", true)
//
//	t is now equal to IntegerSeparatorType(0).Thousands()
func (intSeparatorType IntegerSeparatorType) XParseString(
	valueString string,
	caseSensitive bool) (
	IntegerSeparatorType,
	error) {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	ePrefix := "IntegerSeparatorType.XParseString() "

	var ok bool
	var integerSeparatorType IntegerSeparatorType

	if caseSensitive {

		integerSeparatorType, ok = mIntegerSeparatorTypeStringToCode[valueString]

		if !ok {
			return IntegerSeparatorType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid IntegerSeparatorType Specification.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		integerSeparatorType, ok = mIntegerSeparatorTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return IntegerSeparatorType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid IntegerSeparatorType Specification.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return integerSeparatorType, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration IntegerSeparatorType.
//
// If the current instance of IntegerSeparatorType is invalid, this
// method will always return a value of
// IntegerSeparatorType(0).None().
//
// # Background
//
// Enumeration IntegerSeparatorType has an underlying type of
// integer (int). This means the type could conceivably be set
// to any integer value. This method ensures that all invalid
// IntegerSeparatorType instances are consistently classified
// as 'None' (IntegerSeparatorType(0).None()). Remember that
// 'None' is considered a valid value.
//
// For example, assume that IntegerSeparatorType was set to an
// invalid integer value of -848972. Calling this method on a
// IntegerSeparatorType with this invalid integer value will
// return an integer value of zero (0) or the equivalent of
// IntegerSeparatorType(0).None(). This conversion is useful in
// generating text strings for meaningful informational and error
// messages.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (intSeparatorType IntegerSeparatorType) XReturnNoneIfInvalid() IntegerSeparatorType {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	isValid := new(integerSeparatorTypeNanobot).
		isValidIntegerSeparatorType(intSeparatorType)

	if !isValid {
		return IntegerSeparatorType(0)
	}

	return intSeparatorType
}

// XValue - This method returns the enumeration value of the
// current IntegerSeparatorType instance.
//
// This is a standard utility method and is NOT part of the
// valid enumerations for this type.
func (intSeparatorType IntegerSeparatorType) XValue() IntegerSeparatorType {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	return intSeparatorType
}

// XValueInt - This method returns the integer value of the
// current IntegerSeparatorType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
func (intSeparatorType IntegerSeparatorType) XValueInt() int {

	lockIntegerSeparatorTypeCode.Lock()

	defer lockIntegerSeparatorTypeCode.Unlock()

	return int(intSeparatorType)
}

// IntSeparatorType - public global constant of type
// IntegerSeparatorType.
//
// This variable serves as an easier, shorthand technique for
// accessing IntegerSeparatorType values.
//
// For easy access to these enumeration values, use the global
// variable IntSeparatorType.
//
//	Example: IntSeparatorType.Thousands()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: IntegerSeparatorType(0).Thousands()
//
// Usage:
//
//	IntSeparatorType.None(),
//	IntSeparatorType.Thousands(),
//	IntSeparatorType.IndiaNumbering(),
//	IntSeparatorType.ChineseNumbering(),
const IntSeparatorType = IntegerSeparatorType(0)

// integerSeparatorTypeNanobot - Provides helper methods for
// enumeration IntegerSeparatorType.
type integerSeparatorTypeNanobot struct {
	lock *sync.Mutex
}

// isValidIntegerSeparatorType - Receives an instance of
// IntegerSeparatorType and returns a boolean value signaling
// whether that IntegerSeparatorType instance is valid.
//
// If the passed instance of IntegerSeparatorType is valid, this method
// returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// IntegerSeparatorType enumeration.
func (intSepTypeNanobot *integerSeparatorTypeNanobot) isValidIntegerSeparatorType(
	enumNumericValType IntegerSeparatorType) bool {

	if intSepTypeNanobot.lock == nil {
		intSepTypeNanobot.lock = new(sync.Mutex)
	}

	intSepTypeNanobot.lock.Lock()

	defer intSepTypeNanobot.lock.Unlock()

	if enumNumericValType < 0 ||
		enumNumericValType > 3 {
		return false
	}

	return true
}

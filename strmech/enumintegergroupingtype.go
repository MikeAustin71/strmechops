package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Do NOT access these maps without first getting
// the lock on 'lockIntegerGroupingTypeCode'.

var mIntegerGroupingTypeCodeToString = map[IntegerGroupingType]string{
	IntegerGroupingType(0): "None",
	IntegerGroupingType(1): "Thousands",
	IntegerGroupingType(2): "IndiaNumbering",
	IntegerGroupingType(3): "ChineseNumbering",
}

var mIntegerGroupingTypeStringToCode = map[string]IntegerGroupingType{
	"None":             IntegerGroupingType(0),
	"Thousands":        IntegerGroupingType(1),
	"IndiaNumbering":   IntegerGroupingType(2),
	"ChineseNumbering": IntegerGroupingType(3),
}

var mIntegerGroupingTypeLwrCaseStringToCode = map[string]IntegerGroupingType{
	"none":             IntegerGroupingType(0),
	"thousands":        IntegerGroupingType(1),
	"indianumbering":   IntegerGroupingType(2),
	"chinesenumbering": IntegerGroupingType(3),
}

// IntegerGroupingType - The Integer Grouping Type enumeration
// is used to specify the type of integer digit grouping when
// creating formatting integer digits number string text
// displays.
//
// The most common type of integer grouping is the 'thousands'
// separation.
//
//	Example: 1,000,000,000
//
// In this example, integer numeric digits are separated by
// commas into groups of three digits ('thousands'). The
// integer separator character in this example is the comma
// (','). However, this character separator will vary among
// different countries and cultures. This enumeration focuses
// strictly on specifying the manner and type of integer
// grouping.
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
// The 'IntegerGroupingType' designates the type of integer
// separation to be used when displaying number strings.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'IntegerGroupingType' type has been
// adapted to function in a manner similar to classic
// enumerations.
//
// 'IntegerGroupingType' is declared as a type 'int'. The
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
// constant 'IntGroupingType'.
//
//	Example: IntGroupingType.Thousands()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: IntegerGroupingType(0).Thousands()
//
// Depending on your editor, intellisense (a.k.a. intelligent
// code completion) may not list the IntegerGroupingType
// methods in alphabetical order.
//
// Be advised that all 'IntegerGroupingType' methods beginning
// with 'X', as well as the method 'String()', are utility
// methods and not part of the enumeration.
type IntegerGroupingType int

var lockIntegerGroupingTypeCode sync.Mutex

// None - Signals that Integer Separation is undefined and will
// not be applied to the text presentation of a number string.
//
//	Example: 6789000000000000
//
// None is considered a valid choice for integer separation.
//
// This method is part of the standard enumeration.
func (intGroupingType IntegerGroupingType) None() IntegerGroupingType {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	return IntegerGroupingType(0)
}

// Thousands - Signals that Integer Separation will be applied.
// Integer digits will be separated into groups of three digits.
//
//	Example: 6,789,000,000,000,000
//
// This method is part of the standard enumeration.
func (intGroupingType IntegerGroupingType) Thousands() IntegerGroupingType {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	return IntegerGroupingType(1)
}

// IndiaNumbering - Signals that India Number integer separation
// will be applied. The India Numbering system groups integers
// for text display like this:
//
//	Example: 6,78,90,00,00,00,00,000
//
// This method is part of the standard enumeration.
func (intGroupingType IntegerGroupingType) IndiaNumbering() IntegerGroupingType {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	return IntegerGroupingType(2)
}

// ChineseNumbering - Signals that Chinese Number integer
// separation will be applied. The Chinese Numbering System
// separates integers into groups of 4-numeric digits.
//
//	Example: 6789,0000,0000,0000
//
// This method is part of the standard enumeration.
func (intGroupingType IntegerGroupingType) ChineseNumbering() IntegerGroupingType {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	return IntegerGroupingType(3)
}

// String - Returns a string with the name of the enumeration
// associated with this current instance of
// 'IntegerGroupingType'.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Usage
//
// t:= IntegerGroupingType(0).Thousands()
// str := t.String()
//
//	str is now equal to 'Thousands'
func (intGroupingType IntegerGroupingType) String() string {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	result, ok :=
		mIntegerGroupingTypeCodeToString[intGroupingType]

	if !ok {
		return "Error: Integer Grouping Type INVALID!"

	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// IntegerGroupingType value is valid.
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
//				IntegerGroupingType(0).Thousands()
//
//	 isValid := intSepType.XIsValid() // isValid == true
//
//	 intSepType = IntegerGroupingType(-99)
//
//	 isValid = intSepType.XIsValid() // isValid == false
func (intGroupingType IntegerGroupingType) XIsValid() bool {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	return new(integerGroupingTypeNanobot).
		isValidIntegerGroupingType(intGroupingType)
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of IntegerGroupingType is returned set to the value
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
//	IntegerGroupingType
//	   - Upon successful completion, this method will return a new
//	     instance of IntegerGroupingType set to the value of the
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
//	t, err := IntegerGroupingType(0).
//	             XParseString("Thousands", true)
//
//	t is now equal to IntegerGroupingType(0).Thousands()
func (intGroupingType IntegerGroupingType) XParseString(
	valueString string,
	caseSensitive bool) (
	IntegerGroupingType,
	error) {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	ePrefix := "IntegerGroupingType.XParseString() "

	var ok bool
	var integerGroupingType IntegerGroupingType

	if caseSensitive {

		integerGroupingType, ok = mIntegerGroupingTypeStringToCode[valueString]

		if !ok {
			return IntegerGroupingType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid IntegerGroupingType Specification.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		integerGroupingType, ok = mIntegerGroupingTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return IntegerGroupingType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid IntegerGroupingType Specification.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return integerGroupingType, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration IntegerGroupingType.
//
// If the current instance of IntegerGroupingType is invalid, this
// method will always return a value of
// IntegerGroupingType(0).None().
//
// # Background
//
// Enumeration IntegerGroupingType has an underlying type of
// integer (int). This means the type could conceivably be set
// to any integer value. This method ensures that all invalid
// IntegerGroupingType instances are consistently classified
// as 'None' (IntegerGroupingType(0).None()). Remember that
// 'None' is considered a valid value.
//
// For example, assume that IntegerGroupingType was set to an
// invalid integer value of -848972. Calling this method on a
// IntegerGroupingType with this invalid integer value will
// return an integer value of zero (0) or the equivalent of
// IntegerGroupingType(0).None(). This conversion is useful in
// generating text strings for meaningful informational and error
// messages.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (intGroupingType IntegerGroupingType) XReturnNoneIfInvalid() IntegerGroupingType {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	isValid := new(integerGroupingTypeNanobot).
		isValidIntegerGroupingType(intGroupingType)

	if !isValid {
		return IntegerGroupingType(0)
	}

	return intGroupingType
}

// XValue - This method returns the enumeration value of the
// current IntegerGroupingType instance.
//
// This is a standard utility method and is NOT part of the
// valid enumerations for this type.
func (intGroupingType IntegerGroupingType) XValue() IntegerGroupingType {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	return intGroupingType
}

// XValueInt - This method returns the integer value of the
// current IntegerGroupingType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
func (intGroupingType IntegerGroupingType) XValueInt() int {

	lockIntegerGroupingTypeCode.Lock()

	defer lockIntegerGroupingTypeCode.Unlock()

	return int(intGroupingType)
}

// IntGroupingType - public global constant of type
// IntegerGroupingType.
//
// This variable serves as an easier, shorthand technique for
// accessing IntegerGroupingType values.
//
// For easy access to these enumeration values, use the global
// variable IntGroupingType.
//
//	Example: IntGroupingType.Thousands()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: IntegerGroupingType(0).Thousands()
//
// Usage:
//
//	IntGroupingType.None()
//	IntGroupingType.Thousands()
//	IntGroupingType.IndiaNumbering()
//	IntGroupingType.ChineseNumbering()
const IntGroupingType = IntegerGroupingType(0)

// integerGroupingTypeNanobot - Provides helper methods for
// enumeration IntegerGroupingType.
type integerGroupingTypeNanobot struct {
	lock *sync.Mutex
}

// isValidIntegerGroupingType - Receives an instance of
// IntegerGroupingType and returns a boolean value signaling
// whether that IntegerGroupingType instance is valid.
//
// If the passed instance of IntegerGroupingType is valid, this method
// returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// IntegerGroupingType enumeration.
func (intGroupingTypeNanobot *integerGroupingTypeNanobot) isValidIntegerGroupingType(
	enumNumericValType IntegerGroupingType) bool {

	if intGroupingTypeNanobot.lock == nil {
		intGroupingTypeNanobot.lock = new(sync.Mutex)
	}

	intGroupingTypeNanobot.lock.Lock()

	defer intGroupingTypeNanobot.lock.Unlock()

	if enumNumericValType < 0 ||
		enumNumericValType > 3 {
		return false
	}

	return true
}

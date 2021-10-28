package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockNumericSignValueType before accessing these
// 'maps'.

var mapNumSignValueTypeCodeToString = map[NumericSignValueType]string{
	NumericSignValueType(0): "None",
	NumericSignValueType(1): "Negative",
	NumericSignValueType(2): "Zero",
	NumericSignValueType(3): "Positive",
}

var mapNumSignValueTypeStringToCode = map[string]NumericSignValueType{
	"None":     NumericSignValueType(0),
	"Negative": NumericSignValueType(1),
	"Zero":     NumericSignValueType(2),
	"Positive": NumericSignValueType(3),
}

var mapNumSignValueTypeLwrCaseStringToCode = map[string]NumericSignValueType{
	"none":     NumericSignValueType(0),
	"negative": NumericSignValueType(1),
	"zero":     NumericSignValueType(2),
	"positive": NumericSignValueType(3),
}

// NumericSignValueType - An enumeration of numeric sign values.
// Any number can be positive, negative or zero.
//
// Positive numeric values are greater than zero. Negative numeric
// values are less than zero and zero is neither positive nor
// negative.
//
// Since the Go Programming Language does not directly support
// enumerations, the NumericSignValueType has been adapted to
// function in a manner similar to classic enumerations.
//
// NumericSignValueType is declared as a type 'int'. The method
// names effectively represent an enumeration of numeric sign
// value types. These methods are listed as follows:
//
//  None            (0) - Signals that 'NumericSignValueType' has
//                        not been initialized and therefore has
//                        no value. This is an error condition.
//
//  Negative        (1) - Signals that the numeric value is
//                        negative meaning that it has a value
//                        less than zero.
//
//  Zero            (2) - Signals that the numeric value is zero.
//
//  Positive        (3) - Signals that the numeric value is
//                        greater than zero.
//
// Note that these numeric equivalent values (0 - 3) are styled for
// data management purposes. In arithmetic calculations, numeric
// sign values are typically represented as integer values:
//         Negative= -1, Zero= 0 and/ Positive= 1.
//
// To convert enumeration values for use in numeric computations,
// call the utility method, XArithmeticValue().
//
// For easy access to these enumeration values, use the global
// constant NumSignVal. Example: NumSignVal.Positive()
//
// Otherwise you will need to use the formal syntax.
// Example: NumericSignValueType(0).Positive()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumericSignValueType methods in
// alphabetical order. Be advised that all NumericSignValueType
// methods beginning with 'X', as well as the method 'String()',
// are utility methods, and NOT part of the enumeration values.
//
type NumericSignValueType int

var lockNumericSignValueType sync.Mutex

// None - Signals that 'NumericSignValueType' has not been
// initialized and therefore has no value. This is an error
// condition.
//
func (nSignValue NumericSignValueType) None() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return NumericSignValueType(0)
}

// Negative - Signals that the numeric value is negative meaning
// that it has a value less than zero.
//
func (nSignValue NumericSignValueType) Negative() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return NumericSignValueType(1)
}

// Zero - Signals that the numeric value is zero.
//
func (nSignValue NumericSignValueType) Zero() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return NumericSignValueType(2)
}

// Positive - Signals that the numeric value is greater than zero.
//
func (nSignValue NumericSignValueType) Positive() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return NumericSignValueType(3)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of NumericSignValueType.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  t:= NumericSignValueType(0).Negative()
//  str := t.String()
//     str is now equal to 'Negative'
//
func (nSignValue NumericSignValueType) String() string {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	result, ok :=
		mapNumSignValueTypeCodeToString[nSignValue]

	if !ok {
		return "Error: NumericSignValueType code UNKNOWN!"
	}

	return result
}

// XArithmeticValue - The integer values assigned to the
// NumericSignValueType enumeration are styled for data management.
//
// As such integer values for the enumeration are assigned as
// follows: None = 0, Negative = 1, Zero = 2 and Positive = 3.
// These values are useful for data management purpose but do not
// conform to the conventional values assigned as numeric sign
// values for purposes of arithmetic calculations.
//
// With arithmetic calculations, numeric sign values are typically
// represented as integer values with Negative= -1, Zero= 0 and
// Positive = +1.
//
// The purpose of this method is to convert enumeration values to
// conventional numeric sign values for use in arithmetic
// calculations.
//
// If the value of the current NumericSignValueType is invalid and
// not equal to Negative, Zero or Positive, this method will return
// an integer value of -99.
//
// If the value of the current NumericSignValueType is valid, the
// following integer values will be returned:
//     Negative == -1
//     Zero     ==  0
//     Positive ==  1
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//  nSignVal := NumericSignValueType(0).Positive()
//  // nSignVal now has integer value of '3'
//
//  numericSignValue := nSignVal.XArithmeticValue()
//  // numericSignValue has an integer value of '1'
//
func (nSignValue NumericSignValueType) XArithmeticValue() int {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	if nSignValue > 3 ||
		nSignValue < 1 {
		return -99
	}

	return int(nSignValue) - 2
}

// XIsPositiveOrNegative - There are cases where it is valuable to
// know whether the numeric sign value is positive or negative.
//
// This method answers that question by returning a boolean value.
// If the returned value is 'true' it signals that the current
// NumericSignValueType type is equal to one of the two following
// values:
//           NumericSignValueType(0).Positive()
//                          Or
//           NumericSignValueType(0).Negative()
//
// If the NumericSignValueType is equal to any value other than the
// two shown above, a value of 'false' is returned.
//
// Specifically, this means that if the NumericSignValueType value is
// 'Zero' or 'None', a boolean value of 'false' is returned.
//
func (nSignValue NumericSignValueType) XIsPositiveOrNegative() bool {
	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	if nSignValue == 1 ||
		nSignValue == 3 {
		return true
	}

	return false
}

// XIsValid - Returns a boolean value signaling whether the current
// NumericSignValueType value is valid.
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
//  numSignVal := NumericSignValueType(0).Positive()
//
//  isValid := numSignVal.XIsValid() // isValid == true
//
//  numSignVal = NumericSignValueType(0).None()
//
//  isValid = numSignVal.XIsValid() // isValid == false
//
func (nSignValue NumericSignValueType) XIsValid() bool {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	if nSignValue > 3 ||
		nSignValue < 1 {
		return false
	}

	return true
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of NumericSignValueType is returned set to the
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
//                        will be case-sensitive and will require an
//                        exact match. Therefore, 'negative' will NOT
//                        match the enumeration name, 'Negative'.
//
//                        If 'false' a case-insensitive search is conducted
//                        for the enumeration name. In this case, 'negative'
//                        will match the enumeration name 'Negative'.
//
// ------------------------------------------------------------------------
//
// Return Values
//
// NumericSignValueType
//     - Upon successful completion, this method will return a new
//       instance of NumericSignValueType set to the value of the
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
// t, err := NumericSignValueType(0).XParseString("Positive", true)
//
//     t is now equal to NumericSignValueType(0).Positive()
//
func (nSignValue NumericSignValueType) XParseString(
	valueString string,
	caseSensitive bool) (NumericSignValueType, error) {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	ePrefix := "NumericSignValueType.XParseString() "

	if len(valueString) < 4 {
		return NumericSignValueType(0),
			fmt.Errorf(ePrefix+"\n"+
				"Input parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var strNumSignValType NumericSignValueType

	if caseSensitive {

		strNumSignValType, ok =
			mapNumSignValueTypeStringToCode[valueString]

		if !ok {
			return NumericSignValueType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSignValueType Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strNumSignValType, ok =
			mapNumSignValueTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumericSignValueType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSignValueType Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strNumSignValType, nil
}

// XValue - This method returns the enumeration value of the current
// NumericSignValueType instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (nSignValue NumericSignValueType) XValue() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return nSignValue
}

// XValueInt - This method returns the integer value of the current
// NumericSignValueType instance.
//
// Bear in mind that the returned integer value is the enumeration
// integer value which cannot be used in arithmetic calculations.
// To access the equivalent arithmetic number sign value, call the
// method XArithmeticValue().
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
func (nSignValue NumericSignValueType) XValueInt() int {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return int(nSignValue)
}

// NumSignVal - public global constant of type
// NumericSignValueType.
//
// This variable serves as an easier, shorthand technique for
// accessing NumericSignValueType values.
//
// Usage:
// NumSignVal.None(),
// NumSignVal.Negative(),
// NumSignVal.Zero(),
// NumSignVal.Positive(),
//
const NumSignVal = NumericSignValueType(0)

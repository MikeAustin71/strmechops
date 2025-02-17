package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockNumericSignValueType before accessing these
// 'maps'.

var mapNumSignValueTypeCodeToString = map[NumericSignValueType]string{
	NumericSignValueType(-2): "None",
	NumericSignValueType(-1): "Negative",
	NumericSignValueType(0):  "Zero",
	NumericSignValueType(1):  "Positive",
}

var mapNumSignValueTypeStringToCode = map[string]NumericSignValueType{
	"None":     NumericSignValueType(-2),
	"Negative": NumericSignValueType(-1),
	"Zero":     NumericSignValueType(0),
	"Positive": NumericSignValueType(1),
}

var mapNumSignValueTypeLwrCaseStringToCode = map[string]NumericSignValueType{
	"none":     NumericSignValueType(-2),
	"negative": NumericSignValueType(-1),
	"zero":     NumericSignValueType(0),
	"positive": NumericSignValueType(1),
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
// ----------------------------------------------------------------
//
// Method        Integer
//
//	Name          Value
//
// ------        -------
//
//	None           (-2) - Signals that 'NumericSignValueType' has
//	                      not been initialized and therefore has
//	                      no value. This is an error condition.
//
//	Negative       (-1) - Signals that the numeric value is
//	                      negative meaning that it has a value
//	                      less than zero.
//
//	Zero            (0) - Signals that the numeric value is zero.
//
//	Positive        (1) - Signals that the numeric value is
//	                      greater than zero.
//
// To convert enumeration values for use in numeric computations,
// call the utility method, XArithmeticValue() or XValueInt.
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
type NumericSignValueType int

var lockNumericSignValueType sync.Mutex

// None - Signals that 'NumericSignValueType' has not been
// initialized and therefore has no value. This is an error
// condition.
func (nSignValue NumericSignValueType) None() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return NumericSignValueType(-2)
}

// Negative - Signals that the numeric value is negative meaning
// that it has a value less than zero.
func (nSignValue NumericSignValueType) Negative() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return NumericSignValueType(-1)
}

// Zero - Signals that the numeric value is zero.
func (nSignValue NumericSignValueType) Zero() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return NumericSignValueType(0)
}

// Positive - Signals that the numeric value is greater than zero.
func (nSignValue NumericSignValueType) Positive() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return NumericSignValueType(1)
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
//	t:= NumericSignValueType(0).Negative()
//	str := t.String()
//	   str is now equal to 'Negative'
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

// XArithmeticValue - Returns the arithmetic value of the current
// NumericSignValueType instance.
//
// NumericSignValueType(0).Positive() = +1
// NumericSignValueType(0).Zero() = 0
// NumericSignValueType(0).Negative() = -1
// NumericSignValueType(0).None() = -2
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	nSignVal := NumericSignValueType(0).Positive()
//	// nSignVal now has integer value of '1'
//
//	numericSignValue := nSignVal.XArithmeticValue()
//	// numericSignValue has an integer value of '1'
func (nSignValue NumericSignValueType) XArithmeticValue() int {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return int(nSignValue)
}

// XIsPositiveOrNegative - There are cases where it is valuable to
// know whether the numeric sign value is positive or negative.
//
// This method answers that question by returning a boolean value.
// If the returned value is 'true' it signals that the current
// NumericSignValueType type is equal to one of the two following
// values:
//
//	NumericSignValueType(0).Positive()
//	               Or
//	NumericSignValueType(0).Negative()
//
// If the NumericSignValueType is equal to any value other than the
// two shown above, a value of 'false' is returned.
//
// Specifically, this means that if the NumericSignValueType value is
// 'Zero' or 'None', a boolean value of 'false' is returned.
func (nSignValue NumericSignValueType) XIsPositiveOrNegative() bool {
	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	if nSignValue == -1 ||
		nSignValue == 1 {
		return true
	}

	return false
}

// XIsZero - There are cases where it is valuable to know whether
// the numeric sign value is zero or not.
//
// This method answers that question by returning a boolean value.
// If the returned value is 'true' it signals that the current
// NumericSignValueType type is equal to zero:
//
//	NumericSignValueType(0).Zero()
//
// If the NumericSignValueType is equal to any value other than the
// zero, a value of 'false' is returned.
//
// Specifically, this means that if the NumericSignValueType value is
// non-zero a boolean value of 'false' is returned.
func (nSignValue NumericSignValueType) XIsZero() bool {
	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	if nSignValue == 0 {
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
//	numSignVal := NumericSignValueType(0).Positive()
//
//	isValid := numSignVal.XIsValid() // isValid == true
//
//	numSignVal = NumericSignValueType(0).None()
//
//	isValid = numSignVal.XIsValid() // isValid == false
func (nSignValue NumericSignValueType) XIsValid() bool {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	return new(numericSignValueTypeNanobot).
		isValidNumSignValueType(
			nSignValue)
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
//	exact match. Therefore, 'negative' will NOT
//	match the enumeration name, 'Negative'.
//
//	If 'false' a case-insensitive search is conducted
//	for the enumeration name. In this case, 'negative'
//	will match the enumeration name 'Negative'.
//
// ------------------------------------------------------------------------
//
// # Return Values
//
// NumericSignValueType
//   - Upon successful completion, this method will return a new
//     instance of NumericSignValueType set to the value of the
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
// t, err := NumericSignValueType(0).XParseString("Positive", true)
//
//	t is now equal to NumericSignValueType(0).Positive()
func (nSignValue NumericSignValueType) XParseString(
	valueString string,
	caseSensitive bool) (NumericSignValueType, error) {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	ePrefix := "NumericSignValueType.XParseString() "

	if len(valueString) < 4 {
		return NumericSignValueType(-2),
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
			return NumericSignValueType(-2),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSignValueType Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		strNumSignValType, ok =
			mapNumSignValueTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumericSignValueType(-2),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSignValueType Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return strNumSignValType, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration NumericSignValueType.
//
// If the current instance of NumericSignValueType is invalid, this
// method will always return a value of
// NumericSignValueType(0).None().
//
// # Background
//
// Enumeration NumericSignValueType has an underlying type of
// integer (int). This means the type could conceivably be set to
// any integer value. This method ensures that all invalid
// NumericSignValueType instances are consistently classified as
// 'None' (NumericSignValueType(0).None()). Remember that 'None' is
// considered an invalid value.
//
// For example, assume that NumericSignValueType was set to an
// integer value of -848972. Calling this method on a
// NumericSignValueType with this invalid integer value will return
// an integer value of zero or the equivalent of
// NumericSignValueType(0).None(). This conversion is useful in
// generating text strings for meaningful informational and error
// messages.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (nSignValue NumericSignValueType) XReturnNoneIfInvalid() NumericSignValueType {

	lockNumericSignValueType.Lock()

	defer lockNumericSignValueType.Unlock()

	isValid := new(numericSignValueTypeNanobot).
		isValidNumSignValueType(
			nSignValue)

	if !isValid {
		return NumericSignValueType(-2)
	}

	return nSignValue
}

// XValue - This method returns the enumeration value of the current
// NumericSignValueType instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
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
const NumSignVal = NumericSignValueType(-2)

// numericSignValueTypeNanobot - Provides helper methods for
// enumeration NumericSignValueType.
type numericSignValueTypeNanobot struct {
	lock *sync.Mutex
}

// isValidNumSignValueType - Receives an instance of
// NumericSignValueType and returns a boolean value signaling
// whether that NumericSignValueType instance is valid.
//
// If the passed instance of NumericSignValueType is valid, this
// method returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// NumericSignValueType enumeration.
func (numSignValTypeNanobot *numericSignValueTypeNanobot) isValidNumSignValueType(
	numSignValueType NumericSignValueType) bool {

	if numSignValTypeNanobot.lock == nil {
		numSignValTypeNanobot.lock = new(sync.Mutex)
	}

	numSignValTypeNanobot.lock.Lock()

	defer numSignValTypeNanobot.lock.Unlock()

	if numSignValueType < -1 ||
		numSignValueType > 1 {

		return false
	}

	return true
}

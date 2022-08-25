package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockCarryResultType before accessing these maps!

var mCarryResultTypeCodeToString = map[CarryResultType]string{
	CarryResultType(-1): "MinusOne",
	CarryResultType(0):  "Zero",
	CarryResultType(1):  "PlusOne",
}

var mCarryResultTypeStringToCode = map[string]CarryResultType{
	"MinusOne": CarryResultType(-1),
	"Zero":     CarryResultType(0),
	"PlusOne":  CarryResultType(1),
}

var mCarryResultTypeLwrCaseStringToCode = map[string]CarryResultType{
	"minusone": CarryResultType(-1),
	"zero":     CarryResultType(0),
	"plusone":  CarryResultType(1),
}

// CarryResultType - The 'Carry Result Type' is an
// enumeration of types codes used to classify the results
// of an addition or subtraction operation.
//
// These codes are designed for use by low level routines
// performing addition or subtraction operations.
//
// For example, when an addition operation is performed on
// two numbers and the result is greater than 9, a carry
// value of plus-one (+1) results. With subtraction
// operations, the same principle applies in that value of
// minus-one is the carry result, and it is subtracted from
// the next digit in sequence.
//
// Example Addition:
//
//			23
//	    + 7
//	    ---
//	     30
//	 In this example, value of plus one is carried to the
//
// 10's digit, generating a total value of thirty (30).
//
// Example Subtraction:
//
//	 23
//	- 4
//	---
//	 19
//
// In this subtraction example, four (4) is subtracted from
// thirteen (13) meaning that a carry value of one was
// subtracted from two (2) to generate a net value of
// nineteen (19).
//
// ----------------------------------------------------------------
//
// # BACKGROUND
//
// Thr 'Carry Result Type' is styled as an enumeration. Since the
// Go Programming Language does not directly support enumerations,
// type CarryResultType has been adapted to function in a manner
// similar to classic enumerations.
//
// CarryResultType is declared as a type 'int' and includes two
// types of methods:
//
//	Enumeration Methods
//	      and
//	Utility Methods
//
// Enumeration methods have names which collectively represent
// an enumeration of the carry result accruing from a single
// addition or subtraction operation.
//
//		  Examples Of Enumeration Method Names:
//		      MinusOne()
//		      Zero()
//		      PlusOne()
//
//		Enumeration methods return an integer value used to designate
//		a specific addition or subtraction carry result.
//
//		Utility methods make up the second type of method included
//		in CarryResultType. These methods are NOT part of the
//		enumeration	but instead provide needed supporting services.
//		All utility methods, with the sole exception of method
//	 String(), have names beginning with 'X' to separate them
//	 from standard enumeration methods.
//		  Examples:
//		    XIsValid()
//		    XParseString()
//		    XValue()
//		    XValueInt()
//
//		The utility method 'String()' supports the Stringer Interface
//		and is not part of the standard enumeration.
//
// ----------------------------------------------------------------
//
// # Enumeration Methods
//
// The CarryResultType enumeration methods are described
// below:
//
//	Method					Integer
//
//	 Name					 Value
//
//	------					-------
//
//	MinusOne					-1
//
//	Signals that the carry result of a
//	subtraction operation is minus one (-1)
//
//	Zero						0
//
//	Signals that the carry result of an
//	addition operation or subtraction
//	operation is zero (0).
//
//	PlusOne						1
//
//	Signals that the carry result of an
//	addition operation is plus one (1)
//
// ----------------------------------------------------------------
//
// # USAGE
//
// For easy access to these enumeration values, use the global
// constant 'CarryType'.
//
//	Example: CarryType.Zero()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: CarryResultType(0).Zero()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the CarryResultType methods in
// alphabetical order.
//
// Be advised that all 'CarryResultType' methods beginning with
// 'X', as well as the method 'String()', are utility methods and
// not part of the enumeration.
type CarryResultType int

var lockCarryResultType sync.Mutex

// MinusOne
//
// Signals that the carry result of a subtraction operation
// is minus one (-1).
//
// This method is part of the standard enumeration.
func (carryResultType CarryResultType) MinusOne() CarryResultType {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	return CarryResultType(-1)
}

// Zero
//
// Signals that the carry result of an/ addition operation
// or subtraction operation is zero (0).
//
// This method is part of the standard enumeration.
func (carryResultType CarryResultType) Zero() CarryResultType {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	return CarryResultType(0)
}

// PlusOne
//
// Signals that the carry result of an addition operation
// is plus one (1)
//
// This method is part of the standard enumeration.
func (carryResultType CarryResultType) PlusOne() CarryResultType {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	return CarryResultType(1)
}

// String - Returns a string with the name of the enumeration
// associated with this current instance of 'CarryResultType'.
//
// This is a standard utility method and is not part of the
// valid enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Usage
//
// t:= CarryResultType(0).PlusOne()
// str := t.String()
//
//	str is now equal to 'PlusOne'
func (carryResultType CarryResultType) String() string {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	result, ok := mCarryResultTypeCodeToString[carryResultType]

	if !ok {

		return "Error: Number Rounding Type Specification UNKNOWN!"

	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// NumberRoundingType value is valid.
//
// Be advised, the enumeration value "None" is considered a VALID
// selection for 'NumberRoundingType'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	 roundingType :=
//				NumberRoundingType(0).HalfAwayFromZero()
//
//	 isValid := roundingType.XIsValid() // isValid == true
//
//	 roundingType = NumberRoundingType(-999)
//
//	 isValid = roundingType.XIsValid() // isValid == false
func (carryResultType CarryResultType) XIsValid() bool {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	return new(carryResultTypeNanobot).
		isValidCarryResultType(
			carryResultType)
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of CarryResultType is returned set to the value
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
//     'plusone' will NOT match the enumeration name,
//     'PlusOne'.
//
//     A case-sensitive search will match any of the following
//     strings:
//     "MinusOne"
//     "Zero"
//     "PlusOne"
//
//     If 'false', a case-insensitive search is conducted for the
//     enumeration name. In this example, 'plusone'
//     WILL MATCH the enumeration name, 'PlusOne'.
//
//     A case-insensitive search will match any of the following
//     lower case names:
//     "minusone"
//     "zero"
//     "plusone"
//
// ----------------------------------------------------------------
//
// Return Values
//
//	CarryResultType
//	   - Upon successful completion, this method will return a new
//	     instance of CarryResultType set to the value of the
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
//	t, err := CarryResultType(0).
//	             XParseString("Zero", true)
//
//	t is now equal to CarryResultType(0).Zero()
func (carryResultType CarryResultType) XParseString(
	valueString string,
	caseSensitive bool) (CarryResultType, error) {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	ePrefix := "CarryResultType.XParseString() "

	var ok bool
	var carryResult CarryResultType

	if caseSensitive {

		carryResult, ok =
			mCarryResultTypeStringToCode[valueString]

		if !ok {
			return CarryResultType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CarryResultType Specification.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		carryResult, ok =
			mCarryResultTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return CarryResultType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CarryResultType Specification.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return carryResult, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration CarryResultType.
//
// If the current instance of CarryResultType is invalid, this
// method will always return a value of
// CarryResultType(0).None().
//
// # Background
//
// Enumeration CarryResultType has an underlying type of
// integer (int). This means the type could conceivably be set
// to any integer value. This method ensures that all invalid
// CarryResultType instances are consistently classified as
// 'Zero' (CarryResultType(0).Zero()). Remember that 'None'
// is considered a VALID selection for 'CarryResultType'.
//
// For example, assume that CarryResultType was set to an
// integer value of -848972. Calling this method on a
// CarryResultType with this invalid integer value will
// return an integer value of zero or the equivalent of
// CarryResultType(0).Zero(). This conversion is useful in
// generating text strings for meaningful informational and
// error messages.
//
// This is a standard utility method and is not part of the
// valid enumerations for this type.
func (carryResultType CarryResultType) XReturnNoneIfInvalid() CarryResultType {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	isValid := new(carryResultTypeNanobot).
		isValidCarryResultType(carryResultType)

	if !isValid {
		return CarryResultType(0)
	}

	return carryResultType
}

// XValue - This method returns the enumeration value of the
// current CarryResultType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
func (carryResultType CarryResultType) XValue() CarryResultType {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	return carryResultType
}

// XValueInt - This method returns the integer value of the current
// carryResultType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
func (carryResultType CarryResultType) XValueInt() int {

	lockCarryResultType.Lock()

	defer lockCarryResultType.Unlock()

	return int(carryResultType)
}

// CarryType - Public global constant of type
// CarryResultType.
//
// This variable serves as an easier, shorthand technique for
// accessing CarryResultType values.
//
// For easy access to these enumeration values, use the global
// variable CarryType.
//
//	Example: CarryType.PlusOne()
//
// Otherwise, you will need to use the formal syntax.
//
//	Example: CarryResultType(0).PlusOne()
//
// Usage:
//
//	CarryType.MinusOne(),
//	CarryType.Zero(),
//	CarryType.PlusOne(),
const CarryType = CarryResultType(0)

// carryResultTypeNanobot - Provides helper methods for
// enumeration CarryResultType.
type carryResultTypeNanobot struct {
	lock *sync.Mutex
}

// isValidCarryResultType - Receives an instance of
// CarryResultType and returns a boolean value signaling whether
// that CarryResultType instance is valid.
//
// If the passed instance of CarryResultType is valid, this
// method returns 'true'.
//
// Be advised, the enumeration value "None" is considered a
// VALID selection for 'CarryResultType'.
//
// This is a standard utility method and is not part of the valid
// CarryResultType enumeration.
func (carryResultNanobot *carryResultTypeNanobot) isValidCarryResultType(
	carryResultType CarryResultType) bool {

	if carryResultNanobot.lock == nil {
		carryResultNanobot.lock = new(sync.Mutex)
	}

	carryResultNanobot.lock.Lock()

	defer carryResultNanobot.lock.Unlock()

	if carryResultType < -1 ||
		carryResultType > 1 {

		return false
	}

	return true
}

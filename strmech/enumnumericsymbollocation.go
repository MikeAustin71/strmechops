package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockNumericSymbolLocation before accessing these
// 'maps'.

var mapNumericSymbolLocationCodeToString = map[NumericSymbolLocation]string{
	NumericSymbolLocation(0): "None",
	NumericSymbolLocation(1): "Before",
	NumericSymbolLocation(2): "Interior",
	NumericSymbolLocation(3): "After",
	NumericSymbolLocation(4): "BeforeAndAfter",
}

var mapNumericSymbolLocationStringToCode = map[string]NumericSymbolLocation{
	"None":           NumericSymbolLocation(0),
	"Before":         NumericSymbolLocation(1),
	"Interior":       NumericSymbolLocation(2),
	"After":          NumericSymbolLocation(3),
	"BeforeAndAfter": NumericSymbolLocation(4),
}

var mapNumericSymbolLocationLwrCaseStringToCode = map[string]NumericSymbolLocation{
	"none":           NumericSymbolLocation(0),
	"before":         NumericSymbolLocation(1),
	"interior":       NumericSymbolLocation(2),
	"after":          NumericSymbolLocation(3),
	"beforeandafter": NumericSymbolLocation(4),
}

// NumericSymbolLocation - Describes the location of a numeric
// symbol, such as a currency sign or decimal point, within a
// number string.
//
// Numeric symbols refer to characters typically found in strings
// of numbers. These include currency signs and decimal separators.
//
// These numeric symbols are located at various positions within a
// number string. For instance, currency symbols are located either
// before or after the numeric value. Whereas decimal points are
// always located in the interior of a number string, between
// numeric digits which comprise the number.
//
// The NumericSymbolLocation enumeration is used to define the
// location of a numeric symbol within a string of numbers.
//
// Since the Go Programming Language does not directly support
// enumerations, the NumericSymbolLocation type has been adapted to
// function in a manner similar to classic enumerations.
//
// NumericSymbolLocation is declared as a type 'int'. The method
// names effectively represent an enumeration of number sign symbol
// location positions. These methods are listed as follows:
//
//		None           (0)
//		- Signals that the NumericSymbolLocation
//	   has not been initialized and therefore
//	   has no value. This is an error condition.
//
//		Before         (1)
//		- Signals that the Numeric Symbol is located
//		  before the first numeric digit in the number
//		  string.
//
//		Interior       (2)
//		- Signals that the Numeric Symbol is located
//		  within the number string. In other words,
//		  it is located between the individual numeric
//		  digits which make up the number string.
//
//		After          (3)
//		- Signals that the Numeric Symbol is located
//		  after the last numeric digit in the number
//		  string.
//
// BeforeAndAfter  (4)
//   - Signals that the Numeric Symbol is located both before the
//     first numeric digit after the last numeric digit in the
//     number string.
//
// ----------------------------------------------------------------
//
// # USAGE
//
// For easy access to these enumeration values, use the global
// constant NumSymLocation.
//
//	Example: NumSymLocation.Before()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumericSymbolLocation(0).Before()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the NumericSymbolLocation methods in
// alphabetical order.
//
// Be advised that all NumericSymbolLocation methods beginning with
// 'X', as well as the method 'String()', are utility methods, and
// are NOT part of the enumeration values.
type NumericSymbolLocation int

var lockNumericSymbolLocation sync.Mutex

// None - Signals that the NumericSymbolLocation has not been
// initialized and therefore has no value.
//
// This is an error condition.
func (nSymLocation NumericSymbolLocation) None() NumericSymbolLocation {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	return NumericSymbolLocation(0)
}

// Before - Signals that the Numeric Symbol is located before the
// first numeric digit in the number string.
//
// This method is part of the standard enumeration.
func (nSymLocation NumericSymbolLocation) Before() NumericSymbolLocation {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	return NumericSymbolLocation(1)
}

// Interior - Signals that the Numeric Symbol is located within the
// number string. In other words, it is located between the
// individual numeric digits which make up the number string.
//
// This method is part of the standard enumeration.
func (nSymLocation NumericSymbolLocation) Interior() NumericSymbolLocation {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	return NumericSymbolLocation(2)
}

// After - Signals that the Numeric Symbol is located after the
// last numeric digit in the number string.
//
// This method is part of the standard enumeration.
func (nSymLocation NumericSymbolLocation) After() NumericSymbolLocation {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	return NumericSymbolLocation(3)
}

// BeforeAndAfter - Signals that the Numeric Symbol is located both
// before the first numeric digit after the last numeric digit in
// the number string.
//
// This method is part of the standard enumeration.
func (nSymLocation NumericSymbolLocation) BeforeAndAfter() NumericSymbolLocation {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	return NumericSymbolLocation(4)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of NumericSymbolLocation.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// Usage
//
//	t:= NumericSymbolLocation(0).After()
//	str := t.String()
//	   str is now equal to 'After'
func (nSymLocation NumericSymbolLocation) String() string {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	result, ok :=
		mapNumericSymbolLocationCodeToString[nSymLocation]

	if !ok {
		return "Error: NumericSymbolLocation code UNKNOWN!"
	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// NumericSymbolLocation value is valid.
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
//	numSymLocation := NumericSymbolLocation(0).Before()
//
//	isValid := numSymLocation.XIsValid() // isValid == true
//
//	numSymLocation = NumericSymbolLocation(0).None()
//
//	isValid = numSymLocation.XIsValid() // isValid == false
func (nSymLocation NumericSymbolLocation) XIsValid() bool {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	return new(numSymbolLocNanobot).isValidTextField(
		nSymLocation)
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of NumericSymbolLocation is returned set to the
// value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	valueString   string
//	   - A string which will be matched against the enumeration
//	     string values. If 'valueString' is equal to one of the
//	     enumeration names, this method will proceed to successful
//	     completion and return the correct enumeration value.
//
//
//	caseSensitive   bool
//	   - If 'true' the search for enumeration names will be
//	     case-sensitive and will require an exact match. Therefore,
//	     'before' will NOT match the enumeration name, 'Before'.
//
//	     If 'false' a case-insensitive search is conducted for the
//	     enumeration name. In this case, 'before' will match the
//	     enumeration name 'Before'.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	NumericSymbolLocation
//	   - Upon successful completion, this method will return a new
//	     instance of NumericSymbolLocation set to the value of
//	     the enumeration matched by the string search performed on
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
// # Usage
//
// t, err := NumericSymbolLocation(0).XParseString("Interior", true)
//
//	t is now equal to NumericSymbolLocation(0).Interior()
//
// t, err = NumericSymbolLocation(0).XParseString("interior", false)
//
//	t is now equal to NumericSymbolLocation(0).Explicit()
func (nSymLocation NumericSymbolLocation) XParseString(
	valueString string,
	caseSensitive bool) (NumericSymbolLocation, error) {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	ePrefix := "NumericSymbolLocation.XParseString() "

	if len(valueString) < 4 {
		return NumericSymbolLocation(0),
			fmt.Errorf(ePrefix+"\n"+
				"Input parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var numSymLocation NumericSymbolLocation

	if caseSensitive {

		numSymLocation, ok =
			mapNumericSymbolLocationStringToCode[valueString]

		if !ok {
			return NumericSymbolLocation(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSymbolLocation Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		numSymLocation, ok =
			mapNumericSymbolLocationLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return NumericSymbolLocation(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumericSymbolLocation Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return numSymLocation, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration NumericSymbolLocation.
//
// If the current instance of NumericSymbolLocation is invalid,
// this method will always return a value of
// NumericSymbolLocation(0).None().
//
// # Background
//
// Enumeration NumericSymbolLocation has an underlying type of
// integer (int). This means the type could conceivably be set to
// any integer value. This method ensures that all invalid
// NumericSymbolLocation instances are consistently classified as
// 'None' (NumericSymbolLocation(0).None()). Remember that 'None'
// is considered an invalid value.
//
// For example, assume that NumericSymbolLocation was set to an
// invalid integer value of -848972. Calling this method on a
// NumericSymbolLocation with this invalid integer value will
// return an integer value of zero or the equivalent of
// NumericSymbolLocation(0).None(). This conversion is useful in
// generating text strings for meaningful informational and error
// messages.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (nSymLocation NumericSymbolLocation) XReturnNoneIfInvalid() NumericSymbolLocation {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	isValid := new(numSymbolLocNanobot).
		isValidTextField(nSymLocation)

	if !isValid {
		return NumericSymbolLocation(0)
	}

	return nSymLocation
}

// XValue - This method returns the enumeration value of the
// current NumericSymbolLocation instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (nSymLocation NumericSymbolLocation) XValue() NumericSymbolLocation {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	return nSymLocation
}

// XValueInt - This method returns the integer value of the current
// nSymLocation NumericSymbolLocation instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (nSymLocation NumericSymbolLocation) XValueInt() int {

	lockNumericSymbolLocation.Lock()

	defer lockNumericSymbolLocation.Unlock()

	return int(nSymLocation)
}

// NumSymLocation - public global constant of type
// NumericSymbolLocation.
//
// This variable serves as an easier, shorthand technique for
// accessing NumericSymbolLocation values.
//
// For easy access to these enumeration values, use this global
// constant NumSymLocation.
//
//	Example: NumSymLocation.Before()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumericSymbolLocation(0).Before()
//
// Usage:
// NumSymLocation.None(),
// NumSymLocation.Before(),
// NumSymLocation.Interior(),
// NumSymLocation.After(),
// NumSymLocation.BeforeAndAfter(),
const NumSymLocation = NumericSymbolLocation(0)

// numSymbolLocNanobot - Provides helper methods for
// enumeration NumericSymbolLocation.
type numSymbolLocNanobot struct {
	lock *sync.Mutex
}

// isValidTextField - Receives an instance of NumericSymbolLocation
// and returns a boolean value signaling whether that
// NumericSymbolLocation instance is valid.
//
// If the passed instance of NumericSymbolLocation is valid, this
// method returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// NumericSymbolLocation enumeration.
func (numSymbolLocNanobot *numSymbolLocNanobot) isValidTextField(
	numSymbolLoc NumericSymbolLocation) bool {

	if numSymbolLocNanobot.lock == nil {
		numSymbolLocNanobot.lock = new(sync.Mutex)
	}

	numSymbolLocNanobot.lock.Lock()

	defer numSymbolLocNanobot.lock.Unlock()

	if numSymbolLoc < 1 ||
		numSymbolLoc > 4 {

		return false
	}

	return true
}

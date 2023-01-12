package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockNumberFieldSymbolPosition before accessing these
// 'maps'.

var mapNumFieldNumSymbolPosCodeToString = map[NumberFieldSymbolPosition]string{
	NumberFieldSymbolPosition(0): "None",
	NumberFieldSymbolPosition(1): "InsideNumField",
	NumberFieldSymbolPosition(2): "OutsideNumField",
}

var mapNumFieldNumSymbolPosStringToCode = map[string]NumberFieldSymbolPosition{
	"None":            NumberFieldSymbolPosition(0),
	"InsideNumField":  NumberFieldSymbolPosition(1),
	"OutsideNumField": NumberFieldSymbolPosition(2),
}

var mapNumFieldNumSymbolPosLwrCaseStringToCode = map[string]NumberFieldSymbolPosition{
	"none":            NumberFieldSymbolPosition(0),
	"insidenumfield":  NumberFieldSymbolPosition(1),
	"outsidenumfield": NumberFieldSymbolPosition(2),
}

// NumberFieldSymbolPosition - An enumeration used to define
// the position of number symbols relative to number fields
// containing numeric values formatted for text display in
// number strings.
//
// ----------------------------------------------------------------
//
// # TERMINOLOGY
//
// Number symbols such as plus signs ('+'), minus signs ('-'),
// parentheses ('()') currency signs ('$') and other numeric
// symbols are positioned either inside of number fields or
// outside number fields containing the numeric values.
//
//		Example-1
//			Number Field Length: 8
//	     Numeric Value: 123.45
//	     Number Symbol: leading minus sign ('-')
//	     Number Symbol Position: Inside Number Field
//	     Formatted Number String: " -123.45"
//	          Number Field Index:  01234567
//	  Total Number String Length: 8
//
//		Example-2
//			Number Field Length: 8
//	     Numeric Value: 123.45
//	     Number Symbol: leading minus sign ('-')
//	     Number Symbol Position: Outside Number Field
//	     Formatted Number String: "-  123.45"
//	          Number Field Index:  012345678
//	  Total Number String Length: 9
//
//		Example-3
//			Number Field Length: 8
//	     Numeric Value: 123.45
//	     Number Symbol: before and after parentheses  ('()')
//	     Number Symbol Position: Outside Number Field
//	     Formatted Number String: "( 123.45 )"
//	          Number Field Index:  0123456789
//	  Total Number String Length: 10
//
// Since the Go Programming Language does not directly
// support enumerations, the 'NumberFieldSymbolPosition'
// type has been adapted to function in a manner similar to
// classic enumerations.
//
// ----------------------------------------------------------------
//
// # BACKGROUND
//
// Number string consist of numeric values converted to their
// text character equivalents. These numeric digit strings
// are commonly formatted within number fields.
//
// Number fields are defined with a field length and numeric
// digits are then positioned within that number field. The
// placement of the numeric digits is achieved through text
// justification which means that the numeric digits are
// either centered within the number field, left justified
// in the number field or right justified in the number
// field. In such cases the number field length is longer
// than the length of the numeric digits string.
//
//		Example-4: Centered
//			Number Field Length: 8
//	     Numeric Value: 123.45
//	     Formatted Number String: " 123.45 "
//	          Number Field Index:  01234567
//	  Total Number String Length: 8
//
//		Example-5: Left Justified
//			Number Field Length: 8
//	     Numeric Value: 123.45
//	     Formatted Number String: "123.45  "
//	          Number Field Index:  01234567
//	  Total Number String Length: 8
//
//		Example-6: Right Justified
//			Number Field Length: 8
//	     Numeric Value: 123.45
//	     Formatted Number String: "  123.45"
//	          Number Field Index:  01234567
//	  Total Number String Length: 8
//
// The NumberFieldSymbolPosition enumeration specifies
// where other number symbols like minus signs ('-'), plus
// signs ('+') or currency symbols ('$') are placed within
// the number field or outside the number field.
//
// ----------------------------------------------------------------
//
// # USAGE
//
// Type NumberFieldSymbolPosition is styled as an
// enumeration. Since the Go Programming Language does not
// directly support enumerations, type
// NumberFieldSymbolPosition has been adapted to function
// in a manner similar to classic enumerations.
//
// NumberFieldSymbolPosition is declared as a type 'int'
// and includes two types of methods:
//
//	Enumeration Methods
//	      and
//	Utility Methods
//
// Enumeration methods have names which collectively
// represent an enumeration of different Number Field
// Symbol positions which may be applied in formatting
// numeric values in text strings.
//
//	  Examples Of Enumeration Method Names:
//		NumberFieldSymbolPosition(0).None()
//		NumberFieldSymbolPosition(0).InsideNumField()
//		NumberFieldSymbolPosition(0).OutsideNumField()
//
//
//	Enumeration methods return an integer value used to designate
//	a specific rounding methodology.
//
//	Utility methods make up the second type of method included in
//	NumberFieldSymbolPosition. These methods are NOT part of the
//	enumeration but instead provide needed supporting services.
//	All utility methods, with the sole exception of method
//	String(), have names beginning with 'X' to separate them
//	from standard enumeration methods.
//	  Examples:
//	    XIsValid()
//	    XParseString()
//	    XValue()
//	    XValueInt()
//
//	The utility method 'String()' supports the Stringer Interface
//	and is not part of the standard enumeration.
//
// For easy access to these enumeration values, use the global
// constant 'NumFieldSymPos'.
//
//	Example: NumFieldSymPos.InsideNumField()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumberFieldSymbolPosition(0).InsideNumField()
//
// Depending on your editor, intellisense (a.k.a. intelligent
// code completion) may not list the NumberFieldSymbolPosition
// methods in alphabetical order.
//
// ----------------------------------------------------------------
//
// # ENUMERATION METHODS
//
// The NumberFieldSymbolPosition enumeration methods are described
// below:
//
// Method                     Integer
//
//	Name                       Value
//
// ------                     -------
//
// None							(0)
//
//	Signals that the Number Field Symbol Position
//	(NumberFieldSymbolPosition) Type is not initialized.
//	This is an error condition.
//
// InsideNumField				(1)
//
//		Signals that the Number Symbol will be positioned
//		inside the Number Field.
//			Example-7:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: leading minus sign ('-')
//				Number Symbol Position: Inside Number Field
//	         Number Text Justification: Right
//				Formatted Number String: " -123.45"
//				Number Field Index:       01234567
//				Total Number String Length: 8
//
//		In this case the final length of the number string is
//		defined by the Number Field length.
//
// OutsideNumField				(2)
//
//		Signals that the Number Symbol we be positioned
//		outside the Number Field.
//
//			Example-8:
//				Number Field Length: 8
//		     	Numeric Value: 123.45
//		     	Number Symbol: leading minus sign ('-')
//		     	Number Symbol Position: Outside Number Field
//	         Number Text Justification: Right
//		     	Formatted Number String: "-  123.45"
//				Number Field Index:       012345678
//				Total Number String Length: 9
//
//			Example-9:
//				Number Field Length: 8
//				Numeric Value: 123.45
//				Number Symbol: before and after parentheses  ('()')
//				Number Symbol Position: Outside Number Field
//	         Number Text Justification: Centered
//				Formatted Number String: "( 123.45 )"
//				Number Field Index:       0123456789
//				Total Number String Length: 10
//
//		In this case the final length of the number string is
//		greater than the Number Field length.
type NumberFieldSymbolPosition int

var lockNumberFieldSymbolPosition sync.Mutex

// None - Signals that the NumberFieldSymbolPosition Type is
// uninitialized.
//
// This is an error condition.
//
// This method is part of the standard enumeration.
func (numFieldSymbolPos NumberFieldSymbolPosition) None() NumberFieldSymbolPosition {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	return NumberFieldSymbolPosition(0)
}

// InsideNumField - Signals that the Number Symbol will be positioned
//
//	inside the Number Field.
//		Example-7:
//			    Number Field Length: 8
//			          Numeric Value: 123.45
//			          Number Symbol: leading minus sign ('-')
//			 Number Symbol Position: Inside Number Field
//			Formatted Number String: " -123.45"
//			     Number Field Index:  01234567
//		 Total Number String Length: 8
//
// In this case the final length of the number string is defined
// by the Number Field length.
//
// This method is part of the standard enumeration.
func (numFieldSymbolPos NumberFieldSymbolPosition) InsideNumField() NumberFieldSymbolPosition {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	return NumberFieldSymbolPosition(1)
}

// OutsideNumField - Signals that the Number Symbol we be
// positioned outside the Number Field.
//
//		Example-8
//			Number Field Length: 8
//	     Numeric Value: 123.45
//	     Number Symbol: leading minus sign ('-')
//	     Number Symbol Position: Outside Number Field
//	     Formatted Number String: "-  123.45"
//	          Number Field Index:  012345678
//	  Total Number String Length: 9
//
//		Example-9
//			Number Field Length: 8
//	     Numeric Value: 123.45
//	     Number Symbol: before and after parentheses  ('()')
//	     Number Symbol Position: Outside Number Field
//	     Formatted Number String: "( 123.45 )"
//	          Number Field Index:  0123456789
//	  Total Number String Length: 10
//
//	In this case the final length of the number string is
//	greater than the Number Field length.
//
// This method is part of the standard enumeration.
func (numFieldSymbolPos NumberFieldSymbolPosition) OutsideNumField() NumberFieldSymbolPosition {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	return NumberFieldSymbolPosition(2)
}

// String - Returns a string with the name of the enumeration associated
// with this instance of 'NumberFieldSymbolPosition'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Usage
//
// t:= NumberFieldSymbolPosition(0).OutsideNumField()
// str := t.String()
//
//	str is now equal to "OutsideNumField"
func (numFieldSymbolPos NumberFieldSymbolPosition) String() string {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	result, ok :=
		mapNumFieldNumSymbolPosCodeToString[numFieldSymbolPos]

	if !ok {
		return "Error: NumberFieldSymbolPosition code UNKNOWN!"
	}

	return result

}

// XIsValid - Returns a boolean value signaling whether the current
// NumberFieldSymbolPosition value is valid.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Usage
//
//	numFieldSymPos := NumberFieldSymbolPosition(0).InsideNumField()
//
//	isValid := numFieldSymPos.XIsValid()
func (numFieldSymbolPos NumberFieldSymbolPosition) XIsValid() bool {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	return new(NumberFieldSymbolPositionNanobot).
		isValidNumberFieldSymbolPosition(
			numFieldSymbolPos)

}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of NumberFieldSymbolPosition is returned set to
// the value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
// valueString			string
//
//	A string which will be matched against the
//	enumeration string values. If 'valueString'
//	is equal to one of the enumeration names, this
//	method will proceed to successful completion
//	and return the correct enumeration value.
//
// caseSensitive		bool
//
//	If 'true' the search for enumeration names
//	will be case-sensitive and will require an
//	exact match. Therefore, 'insidenumfield' will
//	NOT	match the enumeration name, 'InsideNumField'.
//
//	If 'false' a case-insensitive search is conducted
//	for the enumeration name. In this case,
//	'insidenumfield' will match the enumeration name
//	'InsideNumField'.
//
// ------------------------------------------------------------------------
//
// # Return Values
//
// NumberFieldSymbolPosition
//
//	Upon successful completion, this method will return a new
//	instance of NumberFieldSymbolPosition set to the value of
//	the enumeration	matched by the string search performed on
//	input parameter, 'valueString'.
//
// error
//
//	If this method completes successfully, the returned error
//	Type is set equal to 'nil'. If an error condition is
//	encountered, this method will return an error type which
//	encapsulates an	appropriate error message.
//
// ------------------------------------------------------------------------
//
// # Usage
//
// t, err :=
//
//		NumberFieldSymbolPosition(0).XParseString(
//		"InsideNumField", true)
//
//	t is now equal to NumberFieldSymbolPosition(0).InsideNumField()
func (numFieldSymbolPos NumberFieldSymbolPosition) XParseString(
	valueString string,
	caseSensitive bool) (
	NumberFieldSymbolPosition,
	error) {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	ePrefix := "NumberFieldSymbolPosition.XParseString() "

	var ok bool
	var numFieldSymPos NumberFieldSymbolPosition

	if caseSensitive {

		numFieldSymPos,
			ok =
			mapNumFieldNumSymbolPosStringToCode[valueString]

		if !ok {
			return NumberFieldSymbolPosition(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumberFieldSymbolPosition Value.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		numFieldSymPos,
			ok = mapNumFieldNumSymbolPosLwrCaseStringToCode[strings.
			ToLower(valueString)]

		if !ok {
			return NumberFieldSymbolPosition(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid NumberFieldSymbolPosition Value.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return numFieldSymPos, nil
}

// XReturnNoneIfInvalid - Provides a standardized value for invalid
// instances of enumeration NumberFieldSymbolPosition.
//
// If the current instance of NumberFieldSymbolPosition is invalid,
// this method will always return a value of
// NumberFieldSymbolPosition(0).None().
//
// # Background
//
// Enumeration NumberFieldSymbolPosition has an underlying type of
// integer (int). This means the type could conceivably be set to
// any integer value. This method ensures that all invalid
// NumberFieldSymbolPosition instances are consistently classified
// as 'None' (NumberFieldSymbolPosition(0).None()). Remember that
// 'None' is considered an invalid value.
//
// For example, assume that NumberFieldSymbolPosition was set to
// an integer value of -848972. Calling this method on a
// NumberFieldSymbolPosition with this invalid integer value will
// return an integer value of zero or the equivalent of
// NumberFieldSymbolPosition(0).None(). This conversion is useful
// in generating text strings for meaningful informational and
// error messages.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (numFieldSymbolPos NumberFieldSymbolPosition) XReturnNoneIfInvalid() NumberFieldSymbolPosition {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	isValid := new(NumberFieldSymbolPositionNanobot).
		isValidNumberFieldSymbolPosition(numFieldSymbolPos)

	if !isValid {
		return NumberFieldSymbolPosition(0)
	}

	return numFieldSymbolPos
}

// XValue - This method returns the enumeration value of the current
// NumberFieldSymbolPosition instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (numFieldSymbolPos NumberFieldSymbolPosition) XValue() NumberFieldSymbolPosition {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	return numFieldSymbolPos
}

// XValueInt - This method returns the integer value of the current
// NumberFieldSymbolPosition instance.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
func (numFieldSymbolPos NumberFieldSymbolPosition) XValueInt() int {

	lockNumberFieldSymbolPosition.Lock()

	defer lockNumberFieldSymbolPosition.Unlock()

	return int(numFieldSymbolPos)
}

// NumFieldSymPos - Public global constant of type
// NumberFieldSymbolPosition.
//
// This variable serves as an easier, shorthand technique for
// accessing NumberFieldSymbolPosition values.
//
// For easy access to these enumeration values, use the global
// variable NumFieldSymPos.
//
//	Example: NumFieldSymPos.InsideNumField()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: NumberFieldSymbolPosition(0).InsideNumField()
//
// Usage:
//
//	NumFieldSymPos.None(),
//	NumFieldSymPos.InsideNumField(),
//	NumFieldSymPos.OutsideNumField(),
const NumFieldSymPos = NumberFieldSymbolPosition(0)

// NumberFieldSymbolPositionNanobot - Provides helper methods for
// enumeration NumberFieldSymbolPosition.
type NumberFieldSymbolPositionNanobot struct {
	lock *sync.Mutex
}

// isValidNumberFieldSymbolPosition - Receives an instance of
// NumberFieldSymbolPosition and returns a boolean value
// signaling whether that NumberFieldSymbolPosition instance
// is valid.
//
// If the passed instance of NumberFieldSymbolPosition is
// valid, this method returns 'true'.
//
// Be advised, the enumeration value "None" is considered NOT
// VALID. "None" represents an error condition.
//
// This is a standard utility method and is not part of the valid
// NumberFieldSymbolPosition enumeration.
func (numberFieldSymPosNanobot *NumberFieldSymbolPositionNanobot) isValidNumberFieldSymbolPosition(
	numFldSymPosValue NumberFieldSymbolPosition) bool {

	if numberFieldSymPosNanobot.lock == nil {
		numberFieldSymPosNanobot.lock = new(sync.Mutex)
	}

	numberFieldSymPosNanobot.lock.Lock()

	defer numberFieldSymPosNanobot.lock.Unlock()

	if numFldSymPosValue < 1 ||
		numFldSymPosValue > 2 {

		return false
	}

	return true
}

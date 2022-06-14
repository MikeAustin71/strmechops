package strmech

import (
	"fmt"
	"strings"
	"sync"
)

var mCharacterSearchTypeCodeToString = map[CharacterSearchType]string{
	CharacterSearchType(-1): "None",
	CharacterSearchType(1):  "LinearTargetChars",
	CharacterSearchType(2):  "SingleTargetChar",
}

var mCharacterSearchTypeStringToCode = map[string]CharacterSearchType{
	"None":              CharacterSearchType(-1),
	"LinearTargetChars": CharacterSearchType(1),
	"Linear":            CharacterSearchType(1),
	"SingleTargetChar":  CharacterSearchType(2),
	"Single":            CharacterSearchType(2),
}

var mCharacterSearchTypeLwrCaseStringToCode = map[string]CharacterSearchType{
	"none":              CharacterSearchType(-1),
	"lineartargetchars": CharacterSearchType(1),
	"linear":            CharacterSearchType(1),
	"singletargetchar":  CharacterSearchType(2),
	"single":            CharacterSearchType(2),
}

// CharacterSearchType - The 'Character Search Type' is an
// enumeration of type codes used for classification of string or
// text character search operations.
//
// Since the Go Programming Language does not directly support
// enumerations, the 'CharacterSearchType' type has been adapted to
// function in a manner similar to classic enumerations.
//
// 'CharacterSearchType' is declared as a type 'int'. The method
// names effectively represent an enumeration of text character
// search types. These methods are listed as follows:
//
// Method             Integer
// Name                Value
// ------             -------
//
// None                 (-1)
//  - Signals that the Character Search Type is empty and not
//    initialized. This is an error condition.
//
//
// LinearTargetChars     (0)
//  - Designates the search type as a Linear Target Character
//    Search Type. This means that each character in the Target
//    Search String will be compared to each character in the Test
//    String. The search will proceed for from left to right in
//    Test Character Sequence.
//
//    A 'Match' or successful search outcome is defined as the case
//    where each character in the Target String matches each
//    corresponding character in the Test String.
//
//    Example
//                             1         2         3
//              Index  0123456789012345678901234567890
//     Target String: "Hey, Xray-4 is the call sign."
//       Test String: "Xray"
//
//    In this example of a Linear Target Characters Search, a match
//    between the Target String and Test String will be declared,
//    if and only if, the search begins at Target String index number
//    5.
//
//    NOTE: Linear Target Characters Search is the default search type.
//
//
// SingleTargetChar      (1)
//  - Designates the search type as a Single Target Character
//    Search Type. This means that a single character in the Target
//    Search String will be compared to all characters in the Test
//    String.
//
//    If a single Target String character equals any character in
//    the Test String, a 'Match' or successful search outcome will
//    be declared.
//
//    The search will proceed from left to right in the Target
//    String. Each Target String Character will be compared to all
//    characters in the Test String looking for the first matching
//    Test String Character.
//
//
//      Example
//                                 1         2         3
//                Index  0123456789012345678901234567890
//       Target String: "Hey, Xray-4 is the call sign."
//         Test String: "XZyURJK"
//
//    In this example of a Single Target Character Search, two
//    matches between the Target String and Test String will be
//    declared at Target String index numbers 5 and 8.
//
//
// For easy access to these enumeration values, use the global
// constant 'CharSearchType'.
//     Example: CharSearchType.Integer()
//
// Otherwise you will need to use the formal syntax.
//     Example: CharacterSearchType(0).Integer()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the CharacterSearchType methods in
// alphabetical order.
//
// Be advised that all 'CharacterSearchType' methods beginning with
// 'X', as well as the method 'String()', are utility methods and
// not part of the enumeration.
//
type CharacterSearchType int

var lockCharacterSearchType sync.Mutex

// None - Signals that the CharacterSearchType specification is
// empty or uninitialized. This is an error condition
//
// This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) None() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(-1)
}

// LinearTargetChars - Signals that each character in the Target
// Search String will be compared to each character in the Test
// String. The search will proceed for from left to right in
// Test Character Sequence.
//
// A 'Match' or successful search outcome is defined as the case
// where each character in the Target String matches each
// corresponding character in the Test String.
//
//  Example
//                             1         2         3
//            Index  0123456789012345678901234567890
//   Target String: "Hey, Xray-4 is the call sign."
//     Test String: "Xray"
//
//   In this example of a Linear Target Characters Search, a match
//   between the Target String and Test String will be declared, if
//   and only if, the search begins at Target String index number
//   5.
//
//   NOTE: Linear Target Characters Search is the default search
//   type.
//
// This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) LinearTargetChars() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(0)
}

// SingleTargetChar - Signals that a single character in the Target
// Search String will be compared to all characters in the Test
// String.
//
// If a single Target String character equals any character in the
// Test String, a 'Match' or successful search outcome will
// be declared.
//
// The search will proceed from left to right in the Target String.
// Each Target String Character will be compared to all characters
// in the Test String looking for the first matching Test String
// Character.
//
//
//  Example
//                             1         2         3
//            Index  0123456789012345678901234567890
//   Target String: "Hey, Xray-4 is the call sign."
//     Test String: "XZyURJK"
//
//   In this example of a Single Target Character Search, two
//   matches between the Target String and Test String will be
//   declared at Target String index numbers 5 and 8.
//
// This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) SingleTargetChar() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(1)
}

// String - Returns a string with the name of the enumeration
// associated with this current instance of 'CharacterSearchType'.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// Usage
//
// t:= CharacterSearchType(0).SingleTargetChar()
// str := t.String()
//     str is now equal to 'SingleTargetChar'
//
func (charSearchType CharacterSearchType) String() string {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()
	result, ok := mCharacterSearchTypeCodeToString[charSearchType]

	if !ok {

		return "Error: Character Search Type Specification UNKNOWN!"

	}

	return result
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of CharacterSearchType is returned set to the value
// of the associated enumeration.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
// valueString   string
//     - A string which will be matched against the enumeration string
//       values. If 'valueString' is equal to one of the enumeration
//       names, this method will proceed to successful completion and
//       return the correct enumeration value.
//
// caseSensitive   bool
//     - If 'true' the search for enumeration names will be
//       case-sensitive and will require an exact match. Therefore,
//       'lineartargetchars' will NOT match the enumeration name,
//       'LinearTargetChars'.
//
//       A case-sensitive search will match any of the following
//       strings:
//           "None"
//           "LinearTargetChars"
//           "Linear"
//           "SingleTargetChar"
//           "Single"
//
//       If 'false', a case-insensitive search is conducted for the
//       enumeration name. In this example, 'lineartargetchars'
//       WILL MATCH the enumeration name, 'LinearTargetChars'.
//
//       A case-insensitive search will match any of the following
//       lower case names:
//           "none"
//           "lineartargetchars"
//           "linear"
//           "singletargetchar"
//           "single"
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  CharacterSearchType
//     - Upon successful completion, this method will return a new
//       instance of CharacterSearchType set to the value of the
//       enumeration matched by the string search performed on
//       input parameter, 'valueString'.
//
//  error
//     - If this method completes successfully, the returned error
//       Type is set equal to 'nil'. If an error condition is
//       encountered, this method will return an error type which
//       encapsulates an appropriate error message.
//
// ----------------------------------------------------------------
//
// Usage
//
//  t, err := CharacterSearchType(0).
//               XParseString("LinearTargetChars", true)
//
//     t is now equal to CharacterSearchType(0).LinearTargetChars()
//
func (charSearchType CharacterSearchType) XParseString(
	valueString string,
	caseSensitive bool) (CharacterSearchType, error) {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	ePrefix := "CharacterSearchType.XParseString() "

	if len(valueString) < 4 {
		return CharacterSearchType(-1),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var characterSearchType CharacterSearchType

	if caseSensitive {

		characterSearchType, ok = mCharacterSearchTypeStringToCode[valueString]

		if !ok {
			return CharacterSearchType(-1),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CharacterSearchType Specification.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		characterSearchType, ok = mCharacterSearchTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return CharacterSearchType(-1),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid CharacterSearchType Specification.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return characterSearchType, nil
}

// XValue - This method returns the enumeration value of the current
// CharacterSearchType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
func (charSearchType CharacterSearchType) XValue() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return charSearchType
}

// XValueInt - This method returns the integer value of the current
// CharacterSearchType instance.
//
// This is a standard utility method and is NOT part of the valid
// enumerations for this type.
//
func (charSearchType CharacterSearchType) XValueInt() int {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return int(charSearchType)
}

// CharSearchType - public global constant of type
// CharacterSearchType.
//
// This variable serves as an easier, shorthand technique for
// accessing CharacterSearchType values.
//
// For easy access to these enumeration values, use the global
// variable CharSearchType.
//  Example: CharSearchType.LinearTargetChars()
//
// Otherwise you will need to use the formal syntax.
//  Example: CharacterSearchType(0).LinearTargetChars()
//
// Usage:
//  CharSearchType.None(),
//  CharSearchType.LinearTargetChars(),
//  CharSearchType.SingleTargetChar(),
//
const CharSearchType = CharacterSearchType(-1)

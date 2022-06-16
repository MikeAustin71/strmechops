package strmech

import (
	"fmt"
	"strings"
	"sync"
)

var mCharacterSearchTypeCodeToString = map[CharacterSearchType]string{
	CharacterSearchType(-1): "None",
	CharacterSearchType(0):  "LinearTargetStartingIndex",
	CharacterSearchType(1):  "SingleTargetChar",
	CharacterSearchType(2):  "LinearEndOfString",
}

var mCharacterSearchTypeStringToCode = map[string]CharacterSearchType{
	"None":                      CharacterSearchType(-1),
	"LinearTargetStartingIndex": CharacterSearchType(0),
	"LinearStartingIndex":       CharacterSearchType(0),
	"StartingIndex":             CharacterSearchType(0),
	"SingleTargetChar":          CharacterSearchType(1),
	"Single":                    CharacterSearchType(1),
	"TargetChar":                CharacterSearchType(1),
	"LinearEndOfString":         CharacterSearchType(2),
	"EndOfString":               CharacterSearchType(2),
}

var mCharacterSearchTypeLwrCaseStringToCode = map[string]CharacterSearchType{
	"none":                      CharacterSearchType(-1),
	"lineartargetstartingindex": CharacterSearchType(0),
	"linearstartingindex":       CharacterSearchType(0),
	"startingindex":             CharacterSearchType(0),
	"singletargetchar":          CharacterSearchType(1),
	"single":                    CharacterSearchType(1),
	"targetchar":                CharacterSearchType(1),
	"linearendofstring":         CharacterSearchType(2),
	"endofstring":               CharacterSearchType(2),
}

// CharacterSearchType - The 'Character Search Type' is an
// enumeration of type codes used for classification of string or
// text character search operations.
//
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//    The Character Search Type relies on a framework consisting of
//    a 'Target Search String' and a 'Test String'.
//
//    All the valid Search Types defined below use the following
//    definitions and terminology.
//
//    Target Search String - A string character or characters which
//                           will be searched for the occurrence of
//                           another predefined character or
//                           characters referred to as a Test
//                           String.
//
//
//    Test String          - A string character or characters which
//                           will be used to search for matching
//                           characters in a Target Search String.
//
//
// ----------------------------------------------------------------
//
// Since the Go Programming Language does not directly support
// enumerations, the 'CharacterSearchType' type has been adapted to
// function in a manner similar to classic enumerations.
//
// 'CharacterSearchType' is declared as a type 'int'. The method
// names effectively represent an enumeration of text character
// search types. These methods are listed as follows:
//
//
// ----------------------------------------------------------------
//
// Method                   Integer
//  Name                     Value
// ------                   -------
//
// None                       (-1)
//  - Signals that the Character Search Type is empty and not
//    initialized. This is an error condition.
//
//
// LinearTargetStartingIndex   (0)
//  - Designates the search type as a Linear Target Starting Index
//    Search Type. This means that each character in the Target
//    Search String will be compared to each corresponding
//    character in the Test String beginning at a specified
//    starting index in the Target Search String.
//
//    The search will proceed for from left to right in Test
//    Character Sequence.
//
//    If the Test Characters are NOT found in the Target Search
//    String beginning at the designated Target String Starting
//    Index, the search outcome will be unsuccessful and NO match
//    will be declared.
//
//    A 'Match', or successful search outcome, is defined as the
//    case where each character in the Target String matches each
//    corresponding character in the Test String beginning at the
//    designated Target String Starting Index.
//
//
//      Example
//                                1         2         3
//               Index  0123456789012345678901234567890
//      Target String: "Hey, Xray-4 is the call sign."
//      Target String Starting Index: 5
//        Test String: "Xray"
//
//    In this example of a Linear Target Starting Index Search, a
//    match between the Target String and Test String will be
//    declared, if and only if, the search begins at Target String
//    index number 5. If the search begins at an any index other
//    than 5, no match will be declared and the search will be
//    classified as unsuccessful.
//
//    NOTE: Linear Target Starting Index is the default search
//          type.
//
//
// SingleTargetChar            (1)
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
//    Test String Character. The search will terminate when a
//    matching character is first identified or when the end of the
//    Target String is encountered.
//
//
//      Example
//                                 1         2         3
//                Index  0123456789012345678901234567890
//       Target String: "Hey, Xray-4 is the call sign."
//       Target String Starting Index: 0
//         Test String: "ZFXyURJK"
//
//    In this example of a Single Target Character Search, the
//    search will terminate at Target String index numbers 5
//    because it is the first Target String index to match one
//    of the Test String Characters ('X').
//
//
// LinearEndOfString           (2)
//  - Designates the search type as a Linear End Of String Search.
//    With this type of search operation, the entire Target Search
//    String will be searched from left to right for the
//    first occurrence of the Test String.
//
//    The search will begin the Target String Starting Index and
//    proceed left to right until (1) an instance of the entire
//    Test String is located or (2) the end of the Target Search
//    String is encountered.
//
//    This is a linear search, so a 'Match' requires that each
//    character in Target Search String must correspond to a
//    matching character in the Test String.
//
//         Example
//                                    1         2         3
//                   Index  0123456789012345678901234567890
//          Target String: "Hey, Xray-4 is the call sign."
//          Target String Starting Index: 0
//            Test String: "Xray-4"
//
//    In this example of a Linear End of String Search, the search
//    operation will begin comparing corresponding characters in
//    the Target Search String and the Test String beginning at
//    index zero. The comparison will fail at index zero, but the
//    search algorithm will continue attempting to find the Test
//    String at indexes 1,2, 3 & 4. The Test String will be found
//    beginning at index number 5 and the search algorithm will
//    terminate at that point with a successful outcome or 'Match'
//    result.
//
//
// ----------------------------------------------------------------
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

// LinearTargetStartingIndex - Signals that each character in the
// Target Search String will be compared to each corresponding
// character in the Test String beginning at a specified starting
// index in the Target Search String.
//
// The search will proceed for from left to right in Test Character
// Sequence.
//
// If the Test Characters are NOT found in the Target Search String
// beginning at the designated Target String Starting Index, the
// search outcome will be unsuccessful and NO match will be
// declared.
//
// A 'Match', or successful search outcome, is defined as the case
// where each character in the Target String matches each
// corresponding character in the Test String beginning at the
// designated Target String Starting Index.
//
//  Example
//                             1         2         3
//            Index  0123456789012345678901234567890
//   Target String: "Hey, Xray-4 is the call sign."
//   Target String Starting Index: 5
//     Test String: "Xray"
//
//   In this example of a Linear Target Starting Index Search, a
//   match between the Target String and Test String will be
//   declared, if and only if, the search begins at Target String
//   index number 5. If the search begins an any index other than
//   5, no match will be declared and the search will be classified
//   as unsuccessful.
//
// NOTE: Linear Target Starting Index is the default search
//       type.
//
//  This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) LinearTargetStartingIndex() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(0)
}

// SingleTargetChar - Signals that a single character in the Target
// Search String will be compared to all characters in the Test
// String.
//
// If a single Target String character equals any character in the
// Test String, a 'Match' or successful search outcome will be
// declared.
//
// The search will proceed from left to right in the Target String.
// Each Target String Character will be compared to all characters
// in the Test String looking for the first matching Test String
// Character. The search will terminate when a matching character
// is first identified or when the end of the Target String is
// encountered.
//
//
//   Example
//                              1         2         3
//             Index  0123456789012345678901234567890
//    Target String: "Hey, Xray-4 is the call sign."
//    Target String Starting Index: 0
//      Test String: "ZFXyURJK"
//
// In this example of a Single Target Character Search, the search
// will terminate at Target String index numbers 5 because it is
// the first Target String index to match one of the Test String
// Characters ('X').
//
// This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) SingleTargetChar() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(1)
}

// LinearEndOfString - Designates the search type as a Linear End
// Of String Search. With this type of search operation, the entire
// Target Search String will be searched from left to right for the
// first occurrence of the Test String.
//
// The search will begin the Target String Starting Index and
// proceed left to right until (1) an instance of the entire Test
// String is located or (2) the end of the Target Search String is
// encountered.
//
// This is a linear search, so a 'Match' requires that each
// character in Target Search String must correspond to a matching
// character in the Test String.
//
//      Example
//                                 1         2         3
//                Index  0123456789012345678901234567890
//       Target String: "Hey, Xray-4 is the call sign."
//       Target String Starting Index: 0
//         Test String: "Xray-4"
//
// In this example of a Linear End of String Search, the search
// operation will begin comparing corresponding characters in the
// Target Search String and the Test String beginning at index zero.
// The comparison will fail at index zero, but the search algorithm
// will continue attempting to find the Test String at indexes 1,2,
// 3 & 4. The Test String will be found beginning at index number 5
// and the search algorithm will terminate at that point with a
// successful outcome or 'Match' result.
//
// This method is part of the standard enumeration.
//
func (charSearchType CharacterSearchType) LinearEndOfString() CharacterSearchType {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	return CharacterSearchType(2)

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

// XIsValid - Returns a boolean value signaling whether the current
// CharacterSearchType value is valid.
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
//  charSearchType :=
// 			CharacterSearchType(0).LinearTargetStartingIndex()
//
//  isValid := charSearchType.XIsValid() // isValid == true
//
//  charSearchType = CharacterSearchType(0).None()
//
//  isValid = charSearchType.XIsValid() // isValid == false
//
func (charSearchType CharacterSearchType) XIsValid() bool {

	lockCharacterSearchType.Lock()

	defer lockCharacterSearchType.Unlock()

	if charSearchType > 2 ||
		charSearchType < 0 {

		return false
	}

	return true
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
//       'LinearTargetStartingIndex'.
//
//       A case-sensitive search will match any of the following
//       strings:
//           "None"
//           "LinearTargetStartingIndex"
//           "LinearStartingIndex"
//           "StartingIndex"
//           "SingleTargetChar"
//           "Single"
//           "TargetChar"
//           "LinearEndOfString"
//           "EndOfString"
//
//       If 'false', a case-insensitive search is conducted for the
//       enumeration name. In this example, 'lineartargetchars'
//       WILL MATCH the enumeration name, 'LinearTargetStartingIndex'.
//
//       A case-insensitive search will match any of the following
//       lower case names:
//           "none"
//           "lineartargetstartingindex"
//           "linearstartingindex"
//           "startingindex"
//           "singletargetchar"
//           "single"
//           "targetchar"
//           "linearendofstring"
//           "endofstring"
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
//               XParseString("LinearTargetStartingIndex", true)
//
//  t is now equal to CharacterSearchType(0).LinearTargetStartingIndex()
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
//  Example: CharSearchType.LinearEndOfString()
//
// Otherwise you will need to use the formal syntax.
//  Example: CharacterSearchType(0).LinearEndOfString()
//
// Usage:
//  CharSearchType.None(),
//  CharSearchType.LinearTargetStartingIndex(),
//  CharSearchType.SingleTargetChar(),
//  CharSearchType.LinearEndOfString(),
//
const CharSearchType = CharacterSearchType(-1)

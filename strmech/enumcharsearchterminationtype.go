package strmech

import (
	"fmt"
	"strings"
	"sync"
)

// Lock lockCharSearchTerminationType before accessing these maps!

var mCharSearchTerminationTypeCodeToString = map[CharSearchTerminationType]string{
	CharSearchTerminationType(0): "None",
	CharSearchTerminationType(1): "ProcessError",
	CharSearchTerminationType(2): "EndOfTargetString",
	CharSearchTerminationType(3): "SearchLengthLimit",
	CharSearchTerminationType(4): "TerminationDelimiters",
	CharSearchTerminationType(5): "FoundSearchTarget",
}

var mCharSearchTerminationTypeStringToCode = map[string]CharSearchTerminationType{
	"None":                   CharSearchTerminationType(0),
	"ProcessError":           CharSearchTerminationType(1),
	"EndOfTargetString":      CharSearchTerminationType(2),
	"End Of TargetString":    CharSearchTerminationType(2),
	"SearchLengthLimit":      CharSearchTerminationType(3),
	"Search Length Limit":    CharSearchTerminationType(3),
	"TerminationDelimiters":  CharSearchTerminationType(4),
	"Termination Delimiters": CharSearchTerminationType(4),
	"FoundSearchTarget":      CharSearchTerminationType(5),
	"Found Search Target":    CharSearchTerminationType(5),
}

var mCharSearchTerminationTypeLwrCaseStringToCode = map[string]CharSearchTerminationType{
	"none":                   CharSearchTerminationType(0),
	"processerror":           CharSearchTerminationType(1),
	"process error":          CharSearchTerminationType(1),
	"endoftargetstring":      CharSearchTerminationType(2),
	"end of target string":   CharSearchTerminationType(2),
	"searchlengthlimit":      CharSearchTerminationType(3),
	"search length limit":    CharSearchTerminationType(3),
	"terminationdelimiters":  CharSearchTerminationType(4),
	"termination delimiters": CharSearchTerminationType(4),
	"foundsearchtarget":      CharSearchTerminationType(5),
	"found search target":    CharSearchTerminationType(5),
}

// CharSearchTerminationType - The 'Character Search Termination
// Type' is an enumeration classifying and identifying the
// reasons for the termination of a text character search
// operation.
//
// ----------------------------------------------------------------
//
// TERMINOLOGY
//
//		A Text Character Search is designed to search inside a target
//	 string of characters looking for a specified Test Character or
//	 Test String. A Text Character Search operation therefore relies
//	 on a framework consisting of a 'Target Search String' and a
//	 'Test String'.
//
//		Target Search String - A string character or characters which
//		                       will be searched for the occurrence of
//		                       another predefined character or
//		                       characters referred to as a Test
//		                       String.
//
//
//		Test String          - A string character or characters which
//		                       will be used to search for matching
//		                       characters in a Target Search String.
//
// ----------------------------------------------------------------
//
//	Once a Text Character Search operation is under way, it can be
//	terminated for one of five reasons:
//
//	 1. A Processing Error may terminate the search.
//
//	 2. The search may reach the end of the 'Target Search
//	    String'.
//
//	 3. The search may be limited by an arbitrary search length
//	    which is less than the string length of the 'Target Search
//	    String'.
//
//	 4. The search operation may encounter one or more characters
//	    in the Target Search String which are designated as
//	    Termination Delimiters. When the character or characters
//	    which comprise a Termination Delimiter are encountered, the
//	    search operation is immediately terminated.
//
//	 5. Depending on the nature of the character search,
//	    termination may occur when the object or target of the
//	    search operation is located.
//
//	The reason for ending a Text Character Search operation should
//	be documented in a clean an unambiguous manner. The enumeration,
//	'CharSearchTerminationType', is designed to fulfill this
//	requirement.
//
// ----------------------------------------------------------------
//
// Since the Go Programming Language does not directly support
// enumerations, the 'CharSearchTerminationType' has been adapted
// to function in a manner similar to classic enumerations.
//
// 'CharSearchTerminationType' is declared as a type 'int'. The
// method names effectively represent an enumeration of reasons for
// text character search termination. These methods are listed as
// follows:
//
// ----------------------------------------------------------------
//
// Method                   Integer
//
//	Name                     Value
//
// ------                   -------
//
// None                       (0)
//   - Signals that the Character Search Termination Reason is
//     empty and not initialized. This is an invalid or error
//     condition.
//
// ProcessError               (1)
//   - Signals that a processing error caused the Text Character
//     Search operation to terminate.
//
// EndOfTargetString          (2)
//   - Signals that the text search operation reached the end of
//     the 'Target Search String' and was forced to terminate.
//
// SearchLengthLimit          (3)
//   - Signals that the search operation was terminated because it
//     reached the end of an arbitrary search length limit. Search
//     Length Limits are less than the length of the Target Search
//     String.
//
// TerminationDelimiters      (4)
//   - Signals that the search operation was terminated because it
//     encountered one or more characters defined as termination
//     delimiters.
//
// FoundSearchTarget          (5)
//   - Signals that the search operation was terminated because the
//     object or target of the search was identified and located.
//
// ----------------------------------------------------------------
//
// For easy access to these enumeration values, use the global
// constant 'CharSearchTermType'.
//
//	Example: CharSearchTermType.TerminationDelimiters()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: CharSearchTerminationType(0).TerminationDelimiters()
//
// Depending on your editor, intellisense (a.k.a. intelligent code
// completion) may not list the 'CharSearchTerminationType' methods
// in alphabetical order.
//
// Be advised that all 'CharSearchTerminationType' methods
// beginning with 'X', as well as the method 'String()', are
// utility methods and not part of the enumeration.
type CharSearchTerminationType int

var lockCharSearchTerminationType sync.Mutex

// None - Signals that the Character Search Termination Reason is
// empty and not initialized. This is an invalid or error condition.
//
// This method is part of the standard enumeration.
func (charSearchTerm CharSearchTerminationType) None() CharSearchTerminationType {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	return CharSearchTerminationType(0)
}

// ProcessError - Signals that a processing error caused the Text
// Character Search operation to terminate.
//
// This method is part of the standard enumeration.
func (charSearchTerm CharSearchTerminationType) ProcessError() CharSearchTerminationType {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	return CharSearchTerminationType(1)

}

// EndOfTargetString - Signals that the text search operation
// reached the end of the 'Target Search String' and was forced to
// terminate.
//
// This method is part of the standard enumeration.
func (charSearchTerm CharSearchTerminationType) EndOfTargetString() CharSearchTerminationType {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	return CharSearchTerminationType(2)

}

// SearchLengthLimit - Signals that the search operation was
// terminated because it reached the end of an arbitrary search
// length limit. Search Length Limits are less than the length of
// the Target Search String.
//
// This method is part of the standard enumeration.
func (charSearchTerm CharSearchTerminationType) SearchLengthLimit() CharSearchTerminationType {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	return CharSearchTerminationType(3)

}

// TerminationDelimiters - Signals that the search operation was
// terminated because it encountered one or more characters defined
// as termination delimiters.
//
// This method is part of the standard enumeration.
func (charSearchTerm CharSearchTerminationType) TerminationDelimiters() CharSearchTerminationType {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	return CharSearchTerminationType(4)

}

// FoundSearchTarget - Signals that the search operation was
// terminated because the object or target of the search was
// identified and located.
//
// This method is part of the standard enumeration.
func (charSearchTerm CharSearchTerminationType) FoundSearchTarget() CharSearchTerminationType {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	return CharSearchTerminationType(5)

}

// String - Returns a string with the name of the enumeration
// associated with this current instance of
// 'CharSearchTerminationType'.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// # Usage
//
// t:= CharSearchTerminationType(0).EndOfTargetString()
// str := t.String()
//
//	str is now equal to 'EndOfTargetString'
func (charSearchTerm CharSearchTerminationType) String() string {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	result, ok := mCharSearchTerminationTypeCodeToString[charSearchTerm]

	if !ok {

		return "Error: Character Search Termination Type UNKNOWN!"

	}

	return result
}

// XIsValid - Returns a boolean value signaling whether the current
// Character Search Termination Type (CharSearchTerminationType) is
// valid.
//
// This is a standard utility method and is not part of the valid
// enumerations for this type.
//
// ----------------------------------------------------------------
//
// Usage
//
//	nStrValueSpec :=
//	    CharSearchTerminationType(0).TerminationDelimiters()
//
//	isValid := nStrValueSpec.XIsValid()
//
//	In this case the boolean value of 'isValid' is 'true'.
//
//	Be advised, the value CharSearchTerminationType(0).None() is
//	classified as an 'invalid' value.
func (charSearchTerm CharSearchTerminationType) XIsValid() bool {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	return new(charSearchTerminationTypeNanobot).
		isValidCharSearchTermType(
			charSearchTerm)
}

// XParseString - Receives a string and attempts to match it with
// the string value of a supported enumeration. If successful, a
// new instance of CharSearchTerminationType is returned set to the
// value of the associated enumeration.
//
// This is a standard utility method and is not part of the valid
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
//     'processerror' will NOT match the enumeration name,
//     'ProcessError'.
//
//     A case-sensitive search will match any of the following strings:
//     "None"
//     "ProcessError"
//     "EndOfTargetString"
//     "End Of TargetString"
//     "SearchLengthLimit"
//     "Search Length Limit"
//     "TerminationDelimiters"
//     "Termination Delimiters"
//     "FoundSearchTarget"
//     "Found Search Target"
//
//     If 'false', a case-insensitive search is conducted for the
//     enumeration name. In this example, 'processerror' WILL MATCH
//     the enumeration name, 'ProcessError'.
//
//     A case-insensitive search will match any of the following
//     lower case names:
//     "none"
//     "processerror"
//     "process error"
//     "endoftargetstring"
//     "end of target string"
//     "searchlengthlimit"
//     "search length limit"
//     "terminationdelimiters"
//     "termination delimiters"
//     "foundsearchtarget"
//     "found search target"
//
// ----------------------------------------------------------------
//
// Return Values
//
//	CharSearchTerminationType
//	   - Upon successful completion, this method will return a new
//	     instance of CharSearchTerminationType set to the value of
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
// Usage
//
//	t, err :=
//	  CharSearchTerminationType(0).
//	    XParseString("EndOfTargetString", true)
//
//	   t is now equal to
//	     CharSearchTerminationType(0).EndOfTargetString()
func (charSearchTerm CharSearchTerminationType) XParseString(
	valueString string,
	caseSensitive bool) (
	CharSearchTerminationType,
	error) {

	lockCharSearchTerminationType.Lock()

	defer lockCharSearchTerminationType.Unlock()

	ePrefix := "CharSearchTerminationType.XParseString() "

	if len(valueString) < 4 {

		return CharSearchTerminationType(0),
			fmt.Errorf(ePrefix+
				"\nInput parameter 'valueString' is INVALID!\n"+
				"String length is less than '4'.\n"+
				"valueString='%v'\n", valueString)
	}

	var ok bool
	var searchTermType CharSearchTerminationType

	if caseSensitive {

		searchTermType, ok = mCharSearchTerminationTypeStringToCode[valueString]

		if !ok {
			return CharSearchTerminationType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid Character Search"+
					" Termination Type Specification.\n"+
					"valueString='%v'\n", valueString)
		}

	} else {

		searchTermType, ok = mCharSearchTerminationTypeLwrCaseStringToCode[strings.ToLower(valueString)]

		if !ok {
			return CharSearchTerminationType(0),
				fmt.Errorf(ePrefix+
					"\n'valueString' did NOT MATCH a valid Character "+
					"Search Termination Type Specification.\n"+
					"valueString='%v'\n", valueString)
		}
	}

	return searchTermType, nil

}

// CharSearchTermType - public global constant of type
// CharSearchTerminationType.
//
// This variable serves as an easier, shorthand technique for
// accessing CharSearchTerminationType values.
//
// For easy access to these enumeration values, use the
// global variable CharSearchTermType.
//
//	Example: CharSearchTermType.SearchLengthLimit()
//
// Otherwise you will need to use the formal syntax.
//
//	Example: CharSearchTerminationType(0).SearchLengthLimit()
//
// Usage:
//
//	CharSearchTermType.None()
//	CharSearchTermType.ProcessError()
//	CharSearchTermType.EndOfTargetString()
//	CharSearchTermType.SearchLengthLimit()
//	CharSearchTermType.EndOfTargetString()
//	CharSearchTermType.TerminationDelimiters()
//	CharSearchTermType.FoundSearchTarget()
const CharSearchTermType = CharSearchTerminationType(0)

// numericValueTypeNanobot - Provides helper methods for
// enumeration NumericValueType.
type charSearchTerminationTypeNanobot struct {
	lock *sync.Mutex
}

func (charSearchTermNanobot *charSearchTerminationTypeNanobot) isValidCharSearchTermType(
	charSearchTermType CharSearchTerminationType) bool {

	if charSearchTermNanobot.lock == nil {
		charSearchTermNanobot.lock = new(sync.Mutex)
	}

	charSearchTermNanobot.lock.Lock()

	defer charSearchTermNanobot.lock.Unlock()

	if charSearchTermType < 1 ||
		charSearchTermType > 5 {
		return false
	}

	return true
}

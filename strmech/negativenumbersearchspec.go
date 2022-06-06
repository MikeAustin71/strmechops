package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NegativeNumberSearchSpec - Negative Number Search Specification.
// This type is designed for use by number string parsing
// functions. These functions review a string of text characters
// searching for numeric digits. The numeric digits are extracted
// to form numeric values. Number string parsing functions
// therefore convert numeric text characters to valid numeric
// values.
//
// Type NegativeNumberSearchSpec allows users to configure search
// parameters for identifying negative numeric values within number
// strings when extracting or parsing numeric digits.
//
// Parsing functions used in converting strings of numeric
// characters into numeric values assume that those values are
// positive unless a Negative Number Sign Symbol or Symbols are
// present in the number string.
//
// Users configure the NegativeNumberSearchSpec type to define the
// criterion for identifying those Negative Number Sign Symbols.
// Number string parsing functions then apply this criterion when
// searching for Negative Number Sign Symbols in number strings.
//
// Examples of Negative Number Sign Symbols:
//
//   "-"   The Minus Sign ('-'). Depending on the country or
//         culture, the Minus Signs could be positioned before or
//         after a string of numeric digits.
//               -127.54
//               - 127.54
//               127.54-
//               127.54 -
//
//
//   "(-)"   These three characters are often used in Europe and
//           the United Kingdom to classify a numeric value as
//           negative.
//               (-) 127.54
//               (-)127.54
//               127.54(-)
//               127.54 (-)
//
//   "()"   Opposing parenthesis characters are frequently used in
//          the United States to classify numeric values as
//          negative.
//               (127.54)
//               ( 127.54 )
//
//
type NegativeNumberSearchSpec struct {
	negNumSignPosition NumSignSymbolPosition // Before(), After(), BeforeAndAfter()
	//                                                   Negative Number Signs are classified
	//                                                   by their location relative to the
	//                                                   numeric digits in a number string.
	leadingNegNumSignSymbols  []rune
	trailingNegNumSignSymbols []rune

	// Processing flags

	foundFirstNumericDigitInNumStr bool // Indicates first numeric digit in
	//                                       the number string has been found
	foundNegNumSignSymbols bool // Indicates all negative number sign symbols
	//                               in this specification have been found
	foundLeadingNegNumSign       bool
	foundLeadingNegNumSignIndex  int
	foundTrailingNegNumSign      bool
	foundTrailingNegNumSignIndex int
	lock                         *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NegativeNumberSearchSpec ('incomingNegNumSearchSpec') to the
// data fields of the current NegativeNumberSearchSpec instance
// ('negNumSearchSpec').
//
// IMPORTANT
// All the data fields in current NegativeNumberSearchSpec instance
// ('negNumSearchSpec') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNegNumSearchSpec   *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
//
//       All data values in this NegativeNumberSearchSpec instance
//       will be copied to current NegativeNumberSearchSpec
//       instance ('negNumSearchSpec').
//
//       If parameter 'incomingNegNumSearchSpec' is determined to
//       be invalid, an error will be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) CopyIn(
	incomingNegNumSearchSpec *NegativeNumberSearchSpec,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	negNumSearchNanobot := negNumSignSearchNanobot{}

	err = negNumSearchNanobot.copyIn(
		negNumSearchSpec,
		incomingNegNumSearchSpec,
		ePrefix.XCpy(
			"negNumSearchSpec<-incomingNegNumSearchSpec"))

	return err
}

// CopyOut - Returns a deep copy of the current
// NegativeNumberSearchSpec instance.
//
// If the current NegativeNumberSearchSpec instance contains
// invalid member variables, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  copyOfNegNumSearchSpec     NegativeNumberSearchSpec
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current NegativeNumberSearchSpec instance.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) CopyOut(
	errorPrefix interface{}) (
	copyOfNegNumSearchSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return copyOfNegNumSearchSpec, err
	}

	copyOfNegNumSearchSpec,
		err =
		negNumSignSearchNanobot{}.ptr().
			copyOut(
				negNumSearchSpec,
				ePrefix.XCpy(
					"copyOfNegNumSearchSpec<-negNumSearchSpec"))

	return copyOfNegNumSearchSpec, err
}

// Empty - Resets all internal member variables for the current
// instance of NegativeNumberSearchSpec to their initial or zero
// states.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of
// NegativeNumberSearchSpec.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (negNumSearchSpec *NegativeNumberSearchSpec) Empty() {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	negNumSearchSpecAtom{}.ptr().empty(
		negNumSearchSpec)

	negNumSearchSpec.lock.Unlock()

	negNumSearchSpec.lock = nil
}

// EmptyProcessingFlags - Resets all the internal processing flags
// to their initial or zero states.
//
// Internal Processing flags are used by Number String parsing
// functions to identify a Negative Number Sign Symbol or Symbols
// in strings of numeric digits called 'Number Strings'. Number
// String parsing functions review strings of text characters
// containing numeric digits and convert those numeric digits to
// numeric values.
//
// The NegativeNumberSearchSpec type includes a series of flags
// which are used to identify a Negative Numeric Sign Symbol or
// Symbols within Number Strings. Number String parsing functions
// use these internal processing flags to record the status of a
// search for a Negative Number Sign Symbol or Symbols defined by
// the current instance of NegativeNumberSearchSpec.
//
// Calling this method will effectively clear all of these internal
// processing flags and prepare the current instance of
// NegativeNumberSearchSpec for a new number string parsing operation.
//
// This method will only reset the internal processing flags:
//  NegativeNumberSearchSpec.foundFirstNumericDigitInNumStr
//  NegativeNumberSearchSpec.foundNegNumSignSymbols
//  NegativeNumberSearchSpec.foundLeadingNegNumSign
//  NegativeNumberSearchSpec.foundLeadingNegNumSignIndex
//  NegativeNumberSearchSpec.foundTrailingNegNumSign
//  NegativeNumberSearchSpec.foundTrailingNegNumSignIndex
//
// This method will not alter the Negative Number Sign Symbols
// configured for the current instance of DecimalSeparatorSpec.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (negNumSearchSpec *NegativeNumberSearchSpec) EmptyProcessingFlags() {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	negNumSearchSpecElectron{}.ptr().emptyProcessingFlags(
		negNumSearchSpec)

}

// Equal - Receives a pointer to another instance of
// NegativeNumberSearchSpec and proceeds to compare its internal
// member variables to those of the current
// NegativeNumberSearchSpec instance in order to determine if they
// are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingNegNumSearchSpec   *NegativeNumberSearchSpec
//     - A pointer to an instance of NegativeNumberSearchSpec. The
//       internal member variable data values in this instance will
//       be compared to those in the current instance of
//       NegativeNumberSearchSpec. The results of this comparison
//       will be returned to the calling functions as a boolean
//       value.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the internal member variable data values contained in
//       input parameter 'incomingNegNumSearchSpec' are equivalent
//       in all respects to those contained in the current instance
//       of NegativeNumberSearchSpec, this return value will be set
//       to 'true'.
//
//       Otherwise, this method will return 'false'.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) Equal(
	incomingNegNumSearchSpec *NegativeNumberSearchSpec) bool {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpecElectron{}.ptr().equal(
		negNumSearchSpec,
		incomingNegNumSearchSpec)
}

// GetFoundLeadingNegNumSign - This boolean flag is set internally
// during a number string parsing operation.
//
// This boolean value signals whether valid Leading Negative Number
// Sign Symbols were located during a number string parsing operation.
//
// This method returns the current value of this boolean value in the
// form of internal member variable:
//   'NegativeNumberSearchSpec.foundLeadingNegNumSign'
//
// If this returned value is set to 'true', it means that valid
// Leading Negative Number Symbol(s) were located in the target
// number string.
//
// If the Leading Negative Number Symbol(s) are NOT present in the
// target number string, this return value is set to 'false'.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetFoundLeadingNegNumSign() bool {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.foundLeadingNegNumSign
}

// GetFoundLeadingNegNumSignIndex - This integer value is set
// internally during a number string parsing operation.
//
// If Leading Negative Number Symbol(s) are present in a number
// string, this value is set to the beginning zero based index of
// the Leading Negative Number Symbol(s).
//
// Effectively, this zero based index marks the beginning of the
// Leading Negative Number Symbol(s) found in the target number
// string.
//
// This method returns the current integer value of this index in
// the form of internal member variable:
//   'NegativeNumberSearchSpec.foundLeadingNegNumSignIndex'
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// Before using this index value, be sure to call method:
//  NegativeNumberSearchSpec.GetFoundLeadingNegNumSign()
//
// The result of this method will signal whether the
// 'foundLeadingNegNumSignIndex' is valid.
//
// If Leading Negative Number Symbol(s) have not yet been located
// in the target number string, this index value is invalid.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetFoundLeadingNegNumSignIndex() int {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.foundLeadingNegNumSignIndex
}

// GetFoundFirstNumericDigit - This boolean flag is set internally
// during a number string parsing operation.
//
// If the first numeric digit in a numeric value has been
// identified in the string parsing operation, the internal member
// variable 'foundFirstNumericDigitInNumStr' is set to the boolean
// value of 'true'. This member variable is typically set by the
// number string parsing routine.
//
// If the first numeric digit has not yet been located in the
// parsing operation, 'foundFirstNumericDigitInNumStr' is set to
// 'false'.
//
// This method returns the status flag
// ('foundFirstNumericDigitInNumStr') indicating whether the first
// numeric digit has been located in the number string parsing
// operation.
//
// Type NegativeNumberSearchSpec uses this flag internally to
// determine if searches for Leading Negative Number Sign Symbols
// are required. If the First Numeric Digit in a number string has
// already been located, then Leading Negative Number Sign Symbols
// are not present in the number string.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetFoundFirstNumericDigit() bool {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.foundFirstNumericDigitInNumStr
}

// GetFoundNegNumSignSymbols - This processing flag is set during a
// number string parsing operation.
//
// If the all the symbols comprising the Negative Number Search
// Specification defined by the current instance of
// NegativeNumberSearchSpec have been located within a number
// string, the internal member variable, 'foundNegNumSignSymbols'
// is set to 'true'.
//
// Otherwise, 'foundNegNumSignSymbols' is set to false signaling
// that a negative number sign matching that defined by the current
// NegativeNumberSearchSpec instance has not yet been identified in
// the target number string.
//
// This internal member variable is typically set by the number
// string parsing routine.
//
// This method returns the status flag ('foundNegNumSignSymbols')
// indicating whether the Negative Number Sign Symbols defined by
// the current NegativeNumberSearchSpec instance have been located by
// the number string parsing routine.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetFoundNegNumSignSymbols() bool {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.foundNegNumSignSymbols
}

// GetFoundTrailingNegNumSign - This boolean flag is set internally
// during a number string parsing operation.
//
// This boolean value signals whether valid Trailing Negative
// Number Sign Symbol(s) were located during a number string
// parsing operation.
//
// This method returns the current value of this boolean value in
// the form of internal member variable:
//   'NegativeNumberSearchSpec.foundTrailingNegNumSign'
//
// If this returned value is set to 'true', it means that valid
// Trailing Negative Number Symbol(s) were located in the target
// number string.
//
// If the Trailing Negative Number Symbol(s) are NOT present in the
// target number string, this return value is set to 'false'.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetFoundTrailingNegNumSign() bool {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.foundTrailingNegNumSign
}

// GetFoundTrailingNegNumSignIndex - This integer value is set
// internally during a number string parsing operation.
//
// If Trailing Negative Number Symbol(s) are present in a number
// string, this value is set to the beginning zero based index of
// the Trailing Negative Number Symbol(s).
//
// Effectively, this zero based index marks the beginning of the
// Trailing Negative Number Symbol(s) found in the target number
// string.
//
// This method returns the current integer value of this index in
// the form of internal member variable:
//   'NegativeNumberSearchSpec.foundTrailingNegNumSignIndex'
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// Before using this index value, be sure to call method:
//  NegativeNumberSearchSpec.GetFoundTrailingNegNumSign()
//
// The result of this method will signal whether the
// 'foundTrailingNegNumSignIndex' is valid.
//
// If Trailing Negative Number Symbol(s) have not yet been located
// in the target number string, this index value is invalid.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetFoundTrailingNegNumSignIndex() int {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.foundTrailingNegNumSignIndex
}

// GetNegNumSignPosition - Returns the position of the Negative
// Number Sign Symbol defined the current instance of
// NegativeNumberSearchSpec.
//
// Negative Number Sign Symbols are positioned in a string of text
// characters relative to the number digits comprising a numeric
// value. As such, Negative Number Sign Symbols will be positioned
// 'Before', 'After' or 'Before and After' the numeric digits.
//
// If the current instance of NumSignSymbolPosition is valid and
// properly configured, this method will return a type
// NumSignSymbolPosition set to one of the following valid values:
//           Value                   Example
//    NSignSymPos.Before()           -123.45
//    NSignSymPos.After()             123.45-
//    NSignSymPos.BeforeAndAfter()   (123.45)
//
func (negNumSearchSpec *NegativeNumberSearchSpec) GetNegNumSignPosition() NumSignSymbolPosition {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	return negNumSearchSpec.negNumSignPosition
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current NegativeNumberSearchSpec
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return a boolean value of 'false'.
//
// If all data elements are determined to be valid, this method
// returns a boolean value of 'true'.
//
// This method is functionally equivalent to
// NegativeNumberSearchSpec.IsValidInstanceError() with the sole
// exceptions being that this method takes no input parameters and
// returns a boolean value.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  -- NONE --
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If any of the internal member data variables contained in
//       the current instance of NegativeNumberSearchSpec are found
//       to be invalid, this method will return a boolean value of
//       'false'.
//
//       If all internal member data variables contained in the
//       current instance of NegativeNumberSearchSpec are found to be
//       valid, this method returns a boolean value of 'true'.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) IsValidInstance() bool {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	isValid,
		_ :=
		negNumSearchSpecAtom{}.ptr().
			testValidityOfNegNumSearchSpec(
				negNumSearchSpec,
				nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current NegativeNumberSearchSpec
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix         interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If any of the internal member data variables contained in
//       the current instance of NegativeNumberSearchSpec are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) IsValidInstanceError(
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err =
		negNumSearchSpecAtom{}.ptr().
			testValidityOfNegNumSearchSpec(
				negNumSearchSpec,
				ePrefix)

	return err
}

// NewLeadingNegNumSearchRunes - Returns a fully populated
// configuration for a Leading Negative Number Search
// Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
//
// This method is identical in function to the method:
//  NegativeNumberSearchSpec.NewLeadingNegNumSearchStr()
//
// The only difference between the two methods is that this method
// receives a rune array as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Leading Negative Number
//       Symbols used in configuring the NegativeNumberSearchSpec
//       instance returned to the calling function.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec   NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Leading Negative Number Search Specification and returned to
//       the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewLeadingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	newLeadingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingNegNumSearchRunes()",
		"")

	if err != nil {
		return newLeadingNegNumSignSpec, err
	}

	negNumSignNanobot := negNumSignSearchNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSearchSpec(
		&newLeadingNegNumSignSpec,
		leadingNegNumSignSymbols,
		ePrefix.XCpy(
			"newLeadingNegNumSignSpec<-leadingNegNumSignSymbols"))

	return newLeadingNegNumSignSpec, err
}

// NewLeadingNegNumSearchStr - Returns a fully populated
// configuration for a Leading Negative Number Search
// Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
//
// This method is identical in function to the method:
//  NegativeNumberSearchSpec.NewLeadingNegNumSearchRunes()
//
// The only difference between the two methods is that this method
// receives a string as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A string identifying the text character or characters
//       which comprise the Leading Negative Number Symbols used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this string is empty (has a zero length) an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Leading Negative Number Search Specification and returned to
//       the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewLeadingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	errorPrefix interface{}) (
	newLeadingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingNegNumSearchStr()",
		"")

	if err != nil {
		return newLeadingNegNumSignSpec, err
	}

	negNumSignNanobot := negNumSignSearchNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSearchSpec(
		&newLeadingNegNumSignSpec,
		[]rune(leadingNegNumSignSymbols),
		ePrefix.XCpy(
			"newLeadingNegNumSignSpec<-leadingNegNumSignSymbols"))

	return newLeadingNegNumSignSpec, err
}

// NewLeadingAndTrailingNegNumSearchRunes - Returns a fully
// populated configuration for a Leading and Trailing Negative
// Number Search Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameters
// 'leadingNegNumSignSymbols' and 'trailingNegNumSignSymbols'.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, the US and Canada use
// parentheses "()" to indicate negative numeric values.
//    Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.NewLeadingAndTrailingNegNumSearchStr()
//
// The only difference between to the two methods is that this
// method receives rune arrays as input parameters.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Leading Negative Number
//       Symbols used in configuring the NegativeNumberSearchSpec
//       instance returned to the calling function.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Trailing Negative Number
//       Symbols used in configuring the NegativeNumberSearchSpec
//       instance returned to the calling function.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Leading and Trailing Negative Number Search Specification
//       and returned to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewLeadingAndTrailingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	leadingAndTrailingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingAndTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingAndTrailingNegNumSearchSpec(
			&leadingAndTrailingNegNumSignSpec,
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return leadingAndTrailingNegNumSignSpec, err
}

// NewLeadingAndTrailingNegNumSearchStr - Returns a fully populated
// configuration for a Leading and Trailing Negative Number Search
// Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameters
// 'leadingNegNumSignSymbols' and 'trailingNegNumSignSymbols '.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, in the US and Canada
// parentheses "()" are used to indicate negative numeric
// values. Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.NewLeadingAndTrailingNegNumSearchRunes()
//
// The only difference between to the two methods is that this
// method receives strings as input parameters.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A string identifying the text character or characters
//       which comprise the Leading Negative Number Symbols used
//       in configuring the NegativeNumberSearchSpec instance
//       returned to the calling function.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the character or characters which
//       comprise the Trailing Negative Number Symbol used in
//       configuring the NegativeNumberSearchSpec instance returned
//       to the calling function.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Leading and Trailing Negative Number Search Specification
//       and returned to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewLeadingAndTrailingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	leadingAndTrailingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewLeadingAndTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return leadingAndTrailingNegNumSignSpec, err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingAndTrailingNegNumSearchSpec(
			&leadingAndTrailingNegNumSignSpec,
			[]rune(leadingNegNumSignSymbols),
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return leadingAndTrailingNegNumSignSpec, err
}

// NewTrailingNegNumSearchRunes - Returns a fully populated
// configuration for a Trailing Negative Number Search
// Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.NewTrailingNegNumSearchStr()
//
// The only difference between to the two methods is that this
// method receives a rune array as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Trailing Negative Number
//       Symbols used in configuring the NegativeNumberSearchSpec
//       instance returned to the calling function.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Trailing Negative Number Search Specification and returned
//       to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewTrailingNegNumSearchRunes(
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	trailingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return trailingNegNumSignSpec, err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setTrailingNegNumSearchSpec(
			&trailingNegNumSignSpec,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"trailingNegNumSignSpec<-trailingNegNumSignSymbols"))

	return trailingNegNumSignSpec, err
}

// NewTrailingNegNumSearchStr - Returns a fully populated
// configuration for a Trailing Negative Number Search
// Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.NewTrailingNegNumSearchRunes()
//
// The only difference between to the two methods is that this
// method receives a string as an input parameter.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the text character or characters
//       which comprise the Trailing Negative Number Symbols used
//       in configuring the NegativeNumberSearchSpec instance
//       returned to the calling function.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as a
//       Trailing Negative Number Search Specification and returned
//       to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec NegativeNumberSearchSpec) NewTrailingNegNumSearchStr(
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	trailingNegNumSignSpec NegativeNumberSearchSpec,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"NewTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return trailingNegNumSignSpec, err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setTrailingNegNumSearchSpec(
			&trailingNegNumSignSpec,
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"trailingNegNumSignSpec<-trailingNegNumSignSymbols"))

	return trailingNegNumSignSpec, err
}

// SearchForNegNumSignSymbols - This method is typically called by
// a number string parsing routine attempting to determine if the
// characters in a search string match the Negative Number Sign
// Symbol defined by this current instance of
// NegativeNumberSearchSpec.
//
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SearchForNegNumSignSymbols(
	targetSearchString *TargetSearchStringDto,
	foundFirstNumericDigitInNumStr bool,
	startingSearchIndex int,
	errorPrefix interface{}) (
	foundNegNumSignSymbols bool,
	lastIndex int,
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundNegNumSignSymbols = false
	lastIndex = startingSearchIndex

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SearchForNegNumSignSymbols()",
		"")

	if err != nil {

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	var err2 error

	err2 = negNumSearchSpec.IsValidInstanceError(
		nil)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: The current instance of NegativeNumberSearchSpec\n"+
			"is invalid. The Number String parsing operation has been aborted.\n"+
			"Validation checks returned the following error for this intance of\n"+
			"NegativeNumberSearchSpec:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return foundNegNumSignSymbols,
			lastIndex,
			err
	}

	negNumSignAtom := negNumSearchSpecAtom{}

	if negNumSearchSpec.negNumSignPosition == NSignSymPos.Before() {

		foundNegNumSignSymbols,
			lastIndex,
			err =
			negNumSignAtom.leadingNegSignSymSearch(
				negNumSearchSpec,
				targetSearchString,
				foundFirstNumericDigitInNumStr,
				startingSearchIndex,
				ePrefix)

	} else if negNumSearchSpec.negNumSignPosition == NSignSymPos.After() {

	} else {
		// Must be: NSignSymPos.BeforeAndAfter()

	}

	return foundNegNumSignSymbols,
		lastIndex,
		err
}

// SetFoundNegNumSignSymbols - Sets the processing flag describing
// the results of a number string parsing operation.
//
// If the all the symbols comprising the Negative Number Sign
// defined by the current instance of NegativeNumberSearchSpec have
// been located within a number string, the internal member
// variable, 'foundNegNumSignSymbols' is set to 'true'.
//
// Otherwise, 'foundNegNumSignSymbols' is set to 'false' signaling
// that a negative number sign matching that defined by the
// current NegativeNumberSearchSpec instance has not yet been
// identified in the target number string.
//
// This internal member variable is typically set by the number
// string parsing routine when calling method:
//    NegativeNumberSearchSpec.SearchForNegNumSignSymbols()
//
// This method sets the processing status flag
// ('foundNegNumSignSymbols') indicating whether the Negative
// Number Sign Symbols defined by the current
// NegativeNumberSearchSpec instance have been located by the number
// string parsing routine.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetFoundNegNumSignSymbols(
	foundNegNumSignSymbols bool) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	negNumSearchSpec.foundNegNumSignSymbols =
		foundNegNumSignSymbols
}

// SetFoundFirstNumericDigit - Sets the internal member variable,
// 'foundFirstNumericDigitInNumStr'.
//
// This flag is typically set during a number string parsing
// operation. If the first numeric digit in a numeric value has
// been identified in the string parsing operation, this internal
// member variable is set to the boolean value of 'true'. Again,
// this member variable is typically set by the number string
// parsing routine.
//
// If the first numeric digit in a number string has not yet been
// identified, this flag is set to 'false'.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetFoundFirstNumericDigit(
	foundFirstNumericDigitInNumStr bool) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	negNumSearchSpec.foundFirstNumericDigitInNumStr =
		foundFirstNumericDigitInNumStr
}

// SetLeadingNegNumSearchRunes - Reconfigures the current instance of
// NegativeNumberSearchSpec as a Leading Negative Number Search
// Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
//
// This method is identical in function to the method:
//  NegativeNumberSearchSpec.SetLeadingNegNumSearchStr()
//
// The only difference between the two methods is that this method
// receives a rune array as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Leading Negative Number Search Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Leading Negative Number
//       Symbols used in configuring the current
//       NegativeNumberSearchSpec instance as a Leading Negative
//       Number Search Specification.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetLeadingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetLeadingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	negNumSignNanobot := negNumSignSearchNanobot{}

	err = negNumSignNanobot.setLeadingNegNumSearchSpec(
		negNumSearchSpec,
		leadingNegNumSignSymbols,
		ePrefix.XCpy(
			"negNumSearchSpec<-leadingNegNumSignSymbols"))

	return err
}

// SetLeadingNegNumSearchStr  - Reconfigures the current instance of
// NegativeNumberSearchSpec as a Leading Negative Number Search
// Specification.
//
// All internal member variables in the returned instance of
// NegativeNumberSearchSpec are configured using the input parameter
// 'leadingNegNumSignSymbols'.
//
// Leading Negative Number symbols are used by many countries
// including the US and Canada. Examples: -123.45  -6,432
//
// This method is identical in function to the method:
//  NegativeNumberSearchSpec.SetLeadingNegNumSearchRunes()
//
// The only difference between the two methods is that this method
// receives a string as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Leading Negative Number Search Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A strung identifying the text character or characters
//       which comprise the Leading Negative Number Symbols used in
//       configuring the current NegativeNumberSearchSpec instance
//       as a Leading Negative Number Search Specification.
//
//       If this string is empty (has a zero length) an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetLeadingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetLeadingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingNegNumSearchSpec(
			negNumSearchSpec,
			[]rune(leadingNegNumSignSymbols),
			ePrefix.XCpy(
				"negNumSearchSpec<-leadingNegNumSignSymbols"))

	return err
}

// SetLeadingAndTrailingNegNumSearchRunes - Reconfigures the current
// instance of NegativeNumberSearchSpec as a Leading and Trailing
// Negative Number Search Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input
// parameters 'leadingNegNumSignSymbols' and
// 'trailingNegNumSignSymbols'.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, the US and Canada use
// parentheses "()" to indicate negative numeric values.
//    Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.SetLeadingAndTrailingNegNumSearchStr()
//
// The only difference between to the two methods is that this
// method receives rune arrays as input parameters.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Leading and Trailing Negative Number Search
// Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Leading Negative Number
//       Symbols used in configuring the current instance of
//       NegativeNumberSearchSpec.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Trailing Negative Number
//       Symbols used in configuring the current instance of
//       NegativeNumberSearchSpec.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetLeadingAndTrailingNegNumSearchRunes(
	leadingNegNumSignSymbols []rune,
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetLeadingAndTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingAndTrailingNegNumSearchSpec(
			negNumSearchSpec,
			leadingNegNumSignSymbols,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"leadingAndTrailingNegNumSignSpec"))

	return err
}

// SetLeadingAndTrailingNegNumSearchStr - Reconfigures the current
// instance of NegativeNumberSearchSpec as a Leading and Trailing
// Negative Number Search Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input
// parameters 'leadingNegNumSignSymbols' and
// 'trailingNegNumSignSymbols'.
//
// In certain nations and cultures, a pair of symbols is used to
// designate a numeric value as negative. These pairs of symbols
// are described here as a Leading and Trailing Negative Number
// Sign Specification. As an example, the US and Canada use
// parentheses "()" to indicate negative numeric values.
//    Examples: (127.45) = -127.45  (4,654.00) = -4,654.00
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.SetLeadingAndTrailingNegNumSearchRunes()
//
// The only difference between to the two methods is that this
// method receives strings as input parameters.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Leading and Trailing Negative Number Search
// Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  leadingNegNumSignSymbols   string
//     - A string identifying the text character or characters
//       which comprise the Leading Negative Number Symbols used in
//       configuring the current instance of
//       NegativeNumberSearchSpec.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the text character or characters
//       which comprise the Trailing Negative Number Symbols used
//       in configuring the current instance of
//       NegativeNumberSearchSpec.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetLeadingAndTrailingNegNumSearchStr(
	leadingNegNumSignSymbols string,
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetLeadingAndTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setLeadingAndTrailingNegNumSearchSpec(
			negNumSearchSpec,
			[]rune(leadingNegNumSignSymbols),
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"negNumSearchSpec"))

	return err
}

// SetTrailingNegNumSearchRunes - Reconfigures the current instance
// of NegativeNumberSearchSpec as a Trailing Negative Number Search
// Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.SetTrailingNegNumSearchStr()
//
// The only difference between to the two methods is that this
// method receives a rune array as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Trailing Negative Number Search Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  trailingNegNumSignSymbols  []rune
//     - An array of runes identifying the text character or
//       characters which comprise the Trailing Negative Number
//       Symbols used in configuring the current
//       NegativeNumberSearchSpec instance as a Trailing Negative
//       Number Search Specification.
//
//       If this array is empty (zero length) or includes array
//       elements containing a zero value, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetTrailingNegNumSearchRunes(
	trailingNegNumSignSymbols []rune,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetTrailingNegNumSearchRunes()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setTrailingNegNumSearchSpec(
			negNumSearchSpec,
			trailingNegNumSignSymbols,
			ePrefix.XCpy(
				"negNumSearchSpec<-trailingNegNumSignSymbols"))

	return err
}

// SetTrailingNegNumSearchStr - Reconfigures the current instance of
// NegativeNumberSearchSpec as a Trailing Negative Number Search
// Specification.
//
// All internal member variables in the current instance of
// NegativeNumberSearchSpec are reconfigured using the input parameter
// 'trailingNegNumSignSymbols'.
//
// Trailing negative number symbols are used by various European
// Union countries. Examples:  127.45-   654-
//
// This method is identical in function to method:
//    NegativeNumberSearchSpec.SetTrailingNegNumSearchRunes()
//
// The only difference between to the two methods is that this
// method receives a string as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All internal member variable data values will be deleted and
// replaced when configuring the current NegativeNumberSearchSpec
// instance as a Trailing Negative Number Search Specification.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  trailingNegNumSignSymbols  string
//     - A string identifying the text character or characters
//       which comprise the Trailing Negative Number Symbols used
//       in configuring the current NegativeNumberSearchSpec
//       instance as a Trailing Negative Number Search Specification.
//
//       If this string is empty (has a zero length), an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// -----------------------------------------------------------------
//
// Return Values
//
//  NegativeNumberSearchSpec     NegativeNumberSearchSpec
//     - If the method completes successfully, a fully populated
//       instance of NegativeNumberSearchSpec will be configured as
//       a Trailing Negative Number Search Specification and
//       returned to the calling function.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (negNumSearchSpec *NegativeNumberSearchSpec) SetTrailingNegNumSearchStr(
	trailingNegNumSignSymbols string,
	errorPrefix interface{}) (
	err error) {

	if negNumSearchSpec.lock == nil {
		negNumSearchSpec.lock = new(sync.Mutex)
	}

	negNumSearchSpec.lock.Lock()

	defer negNumSearchSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NegativeNumberSearchSpec."+
			"SetTrailingNegNumSearchStr()",
		"")

	if err != nil {
		return err
	}

	err = negNumSignSearchNanobot{}.ptr().
		setTrailingNegNumSearchSpec(
			negNumSearchSpec,
			[]rune(trailingNegNumSignSymbols),
			ePrefix.XCpy(
				"negNumSearchSpec<-trailingNegNumSignSymbols"))

	return err
}

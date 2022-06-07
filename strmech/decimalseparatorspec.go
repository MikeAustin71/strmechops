package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// DecimalSeparatorSpec - Decimal Separator Specification.
//
// A decimal separator is one or more text characters used to
// separate integer digits from fractional digits within a
// number string.
//
// This type performs two major functions.
//
// First, it is used by number string parsing functions to search
// for decimal separators within a number string or string of
// numeric digits. Number string parsing functions are designed to
// convert strings of numeric text characters into numeric values.
//
// Second, the DecimalSeparatorSpec type is used to format number
// strings. Number string formatting functions likewise use the
// Decimal Separator Specification to separate integer and
// fractional numeric digits when formatting a number string
// comprised of a floating point numeric value.
//
// Decimal Separators are comprised of a text character or
// characters and are used to separate integer digits from
// fractional digits in floating point numeric values.
//
// The specific characters used as decimal separators vary by
// country and culture.
//
// For example, in the United States, the decimal point or period
// ('.') serves as the decimal separator. Example: 127.54
//
// In various European countries, the comma (',') is used as a
// decimal separator. Example: 127,54
//
// Type DecimalSeparatorSpec allows the user to configure a
// detailed specification for a Decimal Separator character or
// characters.
//
type DecimalSeparatorSpec struct {
	decimalSeparatorChars []rune // Contains the character or characters
	//                                        which comprise the Decimal
	//                                        Separator.

	// Processing flags
	//
	// Internal Processing flags are used by Number String parsing
	// functions to identify decimal separators in strings of numeric
	// digits called 'Number Strings'. These Number String parsing
	// functions review strings of text characters containing numeric
	// digits and convert those numeric digits to numeric values.

	foundFirstNumericDigitInNumStr bool // Indicates first numeric digit in
	//                                       the number string has been found
	foundDecimalSeparatorSymbols bool // Indicates that the decimal separator
	//                                       characters have been found in the
	//                                       number string.
	foundDecimalSeparatorIndex int // Holds the zero based index of the
	//                                       number where the beginning Decimal
	//                                       Separator Character was
	//                                       found in the number string.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// DecimalSeparatorSpec ('incomingDecSepSpec') to the data fields
// of the current DecimalSeparatorSpec instance
// ('decSeparatorSpec').
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current DecimalSeparatorSpec instance
// ('decSeparatorSpec') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingDecSepSpec   *DecimalSeparatorSpec
//     - A pointer to an instance of DecimalSeparatorSpec. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
//
//       All data values in this DecimalSeparatorSpec instance
//       will be copied to current DecimalSeparatorSpec
//       instance ('decSeparatorSpec').
//
//       If parameter 'incomingDecSepSpec' is determined to
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
func (decSeparatorSpec *DecimalSeparatorSpec) CopyIn(
	incomingDecSepSpec *DecimalSeparatorSpec,
	errorPrefix interface{}) (
	err error) {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DecimalSeparatorSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	err = decimalSepSpecNanobot{}.ptr().
		copyIn(
			decSeparatorSpec,
			incomingDecSepSpec,
			ePrefix.XCpy(
				"decSeparatorSpec<-incomingDecSepSpec"))

	return err
}

// CopyOut - Returns a deep copy of the current
// DecimalSeparatorSpec instance.
//
// If the current DecimalSeparatorSpec instance contains
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
//  copyOfDecSepSpec           DecimalSeparatorSpec
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current DecimalSeparatorSpec instance.
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
func (decSeparatorSpec *DecimalSeparatorSpec) CopyOut(
	errorPrefix interface{}) (
	copyOfDecSepSpec DecimalSeparatorSpec,
	err error) {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DecimalSeparatorSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return copyOfDecSepSpec, err
	}

	copyOfDecSepSpec,
		err = decimalSepSpecNanobot{}.ptr().
		copyOut(
			decSeparatorSpec,
			ePrefix.XCpy(
				"copyOfDecSepSpec<-decSeparatorSpec"))

	return copyOfDecSepSpec, err
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
// DecimalSeparatorSpec.
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
func (decSeparatorSpec *DecimalSeparatorSpec) Empty() {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	decimalSeparatorSpecAtom{}.ptr().
		empty(
			decSeparatorSpec)

	decSeparatorSpec.lock.Unlock()

	decSeparatorSpec.lock = nil
}

// EmptyProcessingFlags - Resets all the internal processing flags
// to their initial or zero states.
//
// Internal Processing flags are used by Number String parsing
// functions to identify a Decimal Separator Symbol or Symbols in
// strings of numeric digits called 'Number Strings'. Number String
// parsing functions review strings of text characters containing
// numeric digits and convert those numeric digits to numeric
// values.
//
// The DecimalSeparatorSpec type includes a series of flags which
// are used to identify a Decimal Separator Symbol or Symbols
// within Number Strings. Number String parsing functions use these
// internal processing flags to record the status of a search for
// a Decimal Separator Symbol or Symbols defined by the current
// instance of DecimalSeparatorSpec.
//
// Calling this method will effectively clear all of these internal
// processing flags and prepare the current instance of
// DecimalSeparatorSpec for a new number string parsing operation.
//
// This method will only reset the internal processing flags:
//   DecimalSeparatorSpec.foundFirstNumericDigitInNumStr
//   DecimalSeparatorSpec.foundDecimalSeparatorSymbols
//   DecimalSeparatorSpec.foundDecimalSeparatorIndex
//
// This method will not alter the Decimal Separator Characters
// configured for the current instance of DecimalSeparatorSpec.
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
func (decSeparatorSpec *DecimalSeparatorSpec) EmptyProcessingFlags() {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	decSepSpecElectron := decimalSepSpecElectron{}

	decSepSpecElectron.emptyProcessingFlags(
		decSeparatorSpec)

	return
}

// Equal - Receives a pointer to another instance of
// DecimalSeparatorSpec and proceeds to compare its internal member
// variables to those of the current DecimalSeparatorSpec instance
// in order to determine if they are equivalent.
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
//  incomingDepSepSpec         *DecimalSeparatorSpec
//     - A pointer to an instance of DecimalSeparatorSpec. The
//       internal member variable data values in this instance will
//       be compared to those in the current instance of
//       DecimalSeparatorSpec. The results of this comparison will
//       be returned to the calling functions as a boolean value.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the internal member variable data values contained in
//       input parameter 'incomingDepSepSpec' are equivalent in all
//       respects to those contained in the current instance of
//       DecimalSeparatorSpec, this return value will be set to
//       'true'.
//
//       Otherwise, this method will return 'false'.
//
func (decSeparatorSpec *DecimalSeparatorSpec) Equal(
	incomingDepSepSpec *DecimalSeparatorSpec) bool {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return decimalSepSpecElectron{}.ptr().
		equal(
			decSeparatorSpec,
			incomingDepSepSpec)
}

// GetDecimalSeparatorRunes - Returns the currently configured
// Decimal Separator character or characters as an array of runes.
//
// Decimal Separators are comprised of a text character or
// characters and are used to separate integer digits from
// fractional digits in floating point numeric values.
//
// The specific characters used as decimal separators vary by
// country and culture.
//
// For example, in the United States, the decimal point or period
// ('.') serves as the decimal separator. Example: 127.54
//
// In various European countries, the comma (',') is used as a
// decimal separator. Example: 127,54
//
// This method will return a string containing the configured
// Decimal Separator text character or characters for the current
// instance of DecimalSeparatorSpec.
//
// If Decimal Separator character(s) have not yet been configured,
// this method will return 'nil'.
//
//
// ----------------------------------------------------------------
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
//  []rune
//     - An array of runes containing the text character or
//       characters configured for the current instance of
//       DecimalSeparatorSpec.
//
//       If Decimal Separator character(s) have not yet been
//       configured, this method will return 'nil'.
//
func (decSeparatorSpec *DecimalSeparatorSpec) GetDecimalSeparatorRunes() []rune {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return decSeparatorSpec.decimalSeparatorChars
}

// GetDecimalSeparatorStr - Returns the currently configured Decimal Separator
// character or characters as a string.
//
// Decimal Separators are comprised of a text character or
// characters and are used to separate integer digits from
// fractional digits in floating point numeric values.
//
// The specific characters used as decimal separators vary by
// country and culture.
//
// For example, in the United States, the decimal point or period
// ('.') serves as the decimal separator. Example: 127.54
//
// In various European countries, the comma (',') is used as a
// decimal separator. Example: 127,54
//
// This method will return a string containing the configured
// Decimal Separator text character or characters for the current
// instance of DecimalSeparatorSpec.
//
// If Decimal Separator character(s) have not yet been configured,
// this method will return an empty string.
//
//
// ----------------------------------------------------------------
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
//  string
//     - A string containing the text character or characters
//       configured for the current instance of
//       DecimalSeparatorSpec.
//
//       If Decimal Separator character(s) have not yet been
//       configured, this method will return an empty string.
//
func (decSeparatorSpec *DecimalSeparatorSpec) GetDecimalSeparatorStr() string {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return string(decSeparatorSpec.decimalSeparatorChars)
}

// GetFoundDecimalSeparatorIndex - This integer value is set
// internally during a number string parsing operation.
//
// As such it is almost exclusively used by Number String parsing
// functions. Users will typically have little or no use for this
// boolean processing flag.
//
// Internal Processing flags like internal member variable
// 'foundDecimalSeparatorIndex' are used by Number String parsing
// functions to identify a Decimal Separator Symbol or Symbols in
// strings of numeric digits called 'Number Strings'. Number String
// parsing functions review strings of text characters containing
// numeric digits and convert those numeric digits to numeric
// values.
//
// If the Decimal Separator character or characters configured for
// the current instance of DecimalSeparatorSpec is located in a
// Number String during a number string parsing operation, this
// 'foundDecimalSeparatorIndex' will store the zero based string
// index marking the beginning of the Decimal Separator character
// or characters within the Number String.
//
// The value of this index is only valid if another internal
// processing flag, 'foundDecimalSeparatorSymbols' is set to
// 'true'. For more information on 'foundDecimalSeparatorSymbols'
// see method:
//    DecimalSeparatorSpec.GetFoundDecimalSeparatorSymbols()
//
// This method returns the internal processing status flag
// 'foundDecimalSeparatorIndex' identifying the zero based index
// location of the first Decimal Separator character in a Number
// String.
//
func (decSeparatorSpec *DecimalSeparatorSpec) GetFoundDecimalSeparatorIndex() int {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return decSeparatorSpec.foundDecimalSeparatorIndex
}

// GetFoundDecimalSeparatorSymbols - This boolean flag is set internally
// during a number string parsing operation.
//
// As such it is almost exclusively used by Number String parsing
// functions. Users will typically have little or no use for this
// boolean processing flag.
//
// Internal Processing flags like internal member variable
// 'foundDecimalSeparatorSymbols' are used by Number String parsing
// functions to identify a Decimal Separator Symbol or Symbols in
// strings of numeric digits called 'Number Strings'. Number String
// parsing functions review strings of text characters containing
// numeric digits and convert those numeric digits to numeric
// values.
//
// If the Decimal Separator character or characters configured for
// the current instance of DecimalSeparatorSpec is located in a
// Number String during a number string parsing operation, this
// boolean value is set to 'true'.
//
// If the subject Decimal Separator character or characters has not
// yet been located in the Number String, this value is set to
// 'false'.
//
// This method returns the internal processing status flag
// 'foundDecimalSeparatorSymbols' indicating whether the Decimal
// Separator character or characters has been located in a number
// string parsing operation.
//
func (decSeparatorSpec *DecimalSeparatorSpec) GetFoundDecimalSeparatorSymbols() bool {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return decSeparatorSpec.foundDecimalSeparatorSymbols
}

// GetFoundFirstNumericDigit - This boolean flag is set internally
// during a number string parsing operation.
//
// As such it is almost exclusively used by Number String parsing
// functions. Users will typically have little or no use for this
// boolean processing flag.
//
// Internal Processing flags like internal member variable
// 'foundFirstNumericDigitInNumStr' are used by Number String
// parsing functions to identify a Decimal Separator Symbol or
// Symbols in strings of numeric digits called 'Number Strings'.
// Number String parsing functions review strings of text
// characters containing numeric digits and convert those numeric
// digits to numeric values.
//
// If the first numeric digit in a numeric value has been
// identified in a Number String parsing operation, the internal
// member variable 'foundFirstNumericDigitInNumStr' is set to the
// boolean value of 'true'. Again This member variable is typically
// set and managed by Number String parsing functions.
//
// If the first numeric digit has not yet been located in the
// parsing operation, 'foundFirstNumericDigitInNumStr' is set to
// 'false'.
//
// This method returns the internal processing status flag
// 'foundFirstNumericDigitInNumStr' indicating whether the first
// numeric digit has been located in a number string parsing
// operation.
//
func (decSeparatorSpec *DecimalSeparatorSpec) GetFoundFirstNumericDigit() bool {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return decSeparatorSpec.foundFirstNumericDigitInNumStr
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current DecimalSeparatorSpec instance
// to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return a boolean value of 'false'.
//
// If all data elements are determined to be valid, this method
// returns a boolean value of 'true'.
//
// This method is functionally equivalent to
// DecimalSeparatorSpec.IsValidInstanceError() with the sole
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
//       the current instance of DecimalSeparatorSpec are found to
//       be invalid, this method will return a boolean value of
//       'false'.
//
//       If all internal member data variables contained in the
//       current instance of DecimalSeparatorSpec are found to be
//       valid, this method returns a boolean value of 'true'.
//
func (decSeparatorSpec *DecimalSeparatorSpec) IsValidInstance() bool {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	isValid,
		_ := decimalSeparatorSpecAtom{}.ptr().
		testValidityOfDecSepSearchSpec(
			decSeparatorSpec,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current DecimalSeparatorSpec
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
//       the current instance of DecimalSeparatorSpec are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (decSeparatorSpec *DecimalSeparatorSpec) IsValidInstanceError(
	errorPrefix interface{}) (
	err error) {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DecimalSeparatorSpec."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = decimalSeparatorSpecAtom{}.ptr().
		testValidityOfDecSepSearchSpec(
			decSeparatorSpec,
			ePrefix.XCpy(
				"decSeparatorSpec"))

	return err
}

// SetDecimalSeparatorRunes - Sets the Decimal Separator Symbols
// for the current instance of DecimalSeparatorSpec.
//
// Decimal Separators are comprised of a text character or
// characters and are used to separate integer digits from
// fractional digits in floating point numeric values.
//
// Type DecimalSeparatorSpec performs two major functions.
//
// First, it is used by number string parsing functions to search
// for decimal separators within a number string or string of
// numeric digits. Number string parsing functions are designed to
// convert strings of numeric text characters into numeric values.
//
// Second, the DecimalSeparatorSpec type is used to format number
// strings. Number string formatting functions likewise use the
// Decimal Separator Specification to separate integer and
// fractional numeric digits when formatting a number string
// comprised of a floating point numeric value for presentation ond
// display.
//
// The specific characters used as decimal separators vary by
// country and culture.
//
// For example, in the United States, the decimal point or period
// ('.') serves as the decimal separator. Example: 127.54
//
// In various European countries, the comma (',') is used as a
// decimal separator. Example: 127,54
//
// This method is identical in function to method:
//   DecimalSeparatorSpec.SetDecimalSeparatorStr()
//
// The sole difference between the two methods is that this method
// receives an array of runes as an input parameter.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current DecimalSeparatorSpec instance
// ('decSeparatorSpec') will be deleted and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decSeparator               string
//     - This string contains the character or characters which
//       will be configured as the Decimal Separator Symbol or
//       Symbols for the current instance of DecimalSeparatorSpec,
//       the Decimal Separator Specification.
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
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (decSeparatorSpec *DecimalSeparatorSpec) SetDecimalSeparatorRunes(
	decSeparator []rune,
	errorPrefix interface{}) (
	err error) {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DecimalSeparatorSpec."+
			"SetDecimalSeparatorStr()",
		"")

	if err != nil {
		return err
	}

	lenDecSepRunes := len(decSeparator)

	if lenDecSepRunes == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decSeparator' is invalid!"+
			"'decSeparator' is an empty array with a length of zero.\n",
			ePrefix.String())

		return err

	}

	sMechPreon := strMechPreon{}

	var err2 error
	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		decSeparator,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter 'decSeparator' is invalid!\n"+
			"This rune array contains invalid characters.\n"+
			"'decSeparator' returned the following validation\n"+
			"error:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	decimalSeparatorSpecAtom{}.ptr().
		empty(decSeparatorSpec)

	decSeparatorSpec.decimalSeparatorChars =
		make([]rune, lenDecSepRunes)

	for i := 0; i < lenDecSepRunes; i++ {

		decSeparatorSpec.decimalSeparatorChars[i] =
			decSeparator[i]

	}

	return err
}

// SetDecimalSeparatorStr - Sets the Decimal Separator Symbols for
// the current instance of DecimalSeparatorSpec.
//
// Decimal Separators are comprised of a text character or
// characters and are used to separate integer digits from
// fractional digits in floating point numeric values.
//
// This type performs two major functions.
//
// First, it is used by number string parsing functions to search
// for decimal separators within a number string or string of
// numeric digits. Number string parsing functions are designed to
// convert strings of numeric text characters into numeric values.
//
// Second, the DecimalSeparatorSpec type is used to format number
// strings. Number string formatting functions likewise use the
// Decimal Separator Specification to separate integer and
// fractional numeric digits when formatting a number string
// comprised of a floating point numeric value for presentation ond
// display.
//
// The specific characters used as decimal separators vary by
// country and culture.
//
// For example, in the United States, the decimal point or period
// ('.') serves as the decimal separator. Example: 127.54
//
// In various European countries, the comma (',') is used as a
// decimal separator. Example: 127,54
//
// This method is identical in function to method:
//   DecimalSeparatorSpec.SetDecimalSeparatorRunes()
//
// The sole difference between the two methods is that this method
// receives a string as an input parameter.
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current DecimalSeparatorSpec instance
// ('decSeparatorSpec') will be deleted and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decSeparator               string
//     - This string contains the character or characters which
//       will be configured as the Decimal Separator Symbol or
//       Symbols for the current instance of DecimalSeparatorSpec,
//       the Decimal Separator Specification.
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
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (decSeparatorSpec *DecimalSeparatorSpec) SetDecimalSeparatorStr(
	decSeparator string,
	errorPrefix interface{}) (
	err error) {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DecimalSeparatorSpec."+
			"SetDecimalSeparatorStr()",
		"")

	if err != nil {
		return err
	}

	decSepRunes := []rune(decSeparator)

	lenDecSepRunes := len(decSepRunes)

	if lenDecSepRunes == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decSeparator' is invalid!"+
			"'decSeparator' is an empty string with a length of zero.\n",
			ePrefix.String())

		return err

	}

	decimalSeparatorSpecAtom{}.ptr().
		empty(decSeparatorSpec)

	decSeparatorSpec.decimalSeparatorChars =
		make([]rune, lenDecSepRunes)

	for i := 0; i < lenDecSepRunes; i++ {

		decSeparatorSpec.decimalSeparatorChars[i] =
			decSepRunes[i]

	}

	return err
}

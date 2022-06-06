package strmech

import (
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
// convert strings to numeric characters into numeric values.
//
// Second, the DecimalSeparatorSpec type is used to format number
// strings. Number string formatting functions likewise use the
// Decimal Separator Specification to separate integer and
// fractional numeric digits when formatting a number string
// comprised of a floating point numeric value.
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
// Type DecimalSeparatorSpec allows the user configure a detailed
// specification for a Decimal Separator character or characters.
//
type DecimalSeparatorSpec struct {
	decimalSeparatorChars []rune // Contains the character or characters
	//                                        which comprise the Decimal
	//                                        Separator.

	// Processing flags
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

	err = decimalSepSearchNanobot{}.ptr().
		copyIn(
			decSeparatorSpec,
			incomingDecSepSpec,
			ePrefix.XCpy(
				"decSeparatorSpec<-incomingDecSepSpec"))

	return err
}

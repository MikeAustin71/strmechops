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
type DecimalSeparatorSpec struct {
	decimalSeparatorChars RuneArrayDto
	// Contains the character or characters
	// which comprise the Decimal Separator.

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
// ----------------------------------------------------------------
//
// Input Parameters
//
//		incomingDecSepSpec   *DecimalSeparatorSpec
//		   - A pointer to an instance of DecimalSeparatorSpec. This
//		     method will NOT change the values of internal member
//		     variables contained in this instance.
//
//		     All data values in this DecimalSeparatorSpec instance
//		     will be copied to current DecimalSeparatorSpec
//		     instance ('decSeparatorSpec').
//
//		     If parameter 'incomingDecSepSpec' is determined to
//		     be invalid, an error will be returned.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

	err = new(decimalSepSpecNanobot).
		copyDecimalSeparator(
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	copyOfDecSepSpec           DecimalSeparatorSpec
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current DecimalSeparatorSpec instance.
//
//
//	err                        error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

	err = new(decimalSepSpecNanobot).
		copyDecimalSeparator(
			&copyOfDecSepSpec,
			decSeparatorSpec,
			ePrefix.XCpy(
				"copyOfDecSepSpec<-decSeparatorSpec"))

	return copyOfDecSepSpec, err
}

// Empty - Resets all internal member variables for the current
// instance of DecimalSeparatorSpec to their initial or zero
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of
// DecimalSeparatorSpec.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (decSeparatorSpec *DecimalSeparatorSpec) Empty() {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	new(decimalSeparatorSpecAtom).
		empty(
			decSeparatorSpec)

	decSeparatorSpec.lock.Unlock()

	decSeparatorSpec.lock = nil
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingDepSepSpec         *DecimalSeparatorSpec
//	   - A pointer to an instance of DecimalSeparatorSpec. The
//	     internal member variable data values in this instance will
//	     be compared to those in the current instance of
//	     DecimalSeparatorSpec. The results of this comparison will
//	     be returned to the calling function as a boolean value.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the internal member variable data values contained in
//	     input parameter 'incomingDepSepSpec' are equivalent in all
//	     respects to those contained in the current instance of
//	     DecimalSeparatorSpec, this return value will be set to
//	     'true'.
//
//	     Otherwise, this method will return 'false'.
func (decSeparatorSpec *DecimalSeparatorSpec) Equal(
	incomingDepSepSpec *DecimalSeparatorSpec) bool {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return new(decimalSepSpecElectron).
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	[]rune
//	   - An array of runes containing the text character or
//	     characters configured for the current instance of
//	     DecimalSeparatorSpec.
//
//	     If Decimal Separator character(s) have not yet been
//	     configured, this method will return 'nil'.
func (decSeparatorSpec *DecimalSeparatorSpec) GetDecimalSeparatorRunes() []rune {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return decSeparatorSpec.decimalSeparatorChars.CharsArray
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	string
//	   - A string containing the text character or characters
//	     configured for the current instance of
//	     DecimalSeparatorSpec.
//
//	     If Decimal Separator character(s) have not yet been
//	     configured, this method will return an empty string.
func (decSeparatorSpec *DecimalSeparatorSpec) GetDecimalSeparatorStr() string {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return string(decSeparatorSpec.decimalSeparatorChars.CharsArray)
}

// GetNumberOfSeparatorChars - Returns an integer value specifying
// the number of decimal separator characters contained in this
// instance of DecimalSeparatorSpec.
func (decSeparatorSpec *DecimalSeparatorSpec) GetNumberOfSeparatorChars() int {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	return decSeparatorSpec.decimalSeparatorChars.GetRuneArrayLength()
}

// IsNOP - Stands for 'Is No Operation'. This method returns a
// boolean value signaling whether this instance of
// DecimalSeparatorSpec is engaged, valid and operational with
// respect to the current search algorithm.
//
// If 'IsNOP' is set to 'true', it signals that this Decimal
// Separator Specification is simply an empty placeholder and
// performs no active role in, and is completely ignored by,
// the search algorithm. With 'IsNOP' set to 'true', no search
// for decimal separator characters will ever be conducted.
//
// If this method returns 'false', it signals that the current
// instance of DecimalSeparatorSpec is fully populated, valid,
// functional and ready to perform the search for decimal
// separator characters.
func (decSeparatorSpec *DecimalSeparatorSpec) IsNOP() bool {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	if decSeparatorSpec.decimalSeparatorChars.GetRuneArrayLength() == 0 {
		return true
	}

	return false
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	-- NONE --
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If any of the internal member data variables contained in
//	     the current instance of DecimalSeparatorSpec are found to
//	     be invalid, this method will return a boolean value of
//	     'false'.
//
//	     If all internal member data variables contained in the
//	     current instance of DecimalSeparatorSpec are found to be
//	     valid, this method returns a boolean value of 'true'.
func (decSeparatorSpec *DecimalSeparatorSpec) IsValidInstance() bool {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	isValid,
		_ := new(decimalSeparatorSpecAtom).
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
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If any of the internal member data variables contained in
//	     the current instance of DecimalSeparatorSpec are found
//	     to be invalid, this method will return an error.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' (error prefix) will be inserted or
//	     prefixed at the beginning of the error message.
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
		err = new(decimalSeparatorSpecAtom).
		testValidityOfDecSepSearchSpec(
			decSeparatorSpec,
			ePrefix.XCpy(
				"decSeparatorSpec"))

	return err
}

// NewRunes - Creates and returns a new instance of
// DecimalSeparatorSpec populated with the decimal
// separator character or characters contained in an
// array runes passed by input parameter 'decSeparator'.
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	decSeparator               []rune
//	   - An array of runes containing the character or
//	     characters which will be configured as the
//	     Decimal Separator Symbol or Symbols for the
//	     returned instance of DecimalSeparatorSpec, the
//	     Decimal Separator Specification.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	newDecimalSeparator        DecimalSeparatorSpec
//	   - If this method completes successfully, this return
//	     parameter is configured with the decimal separator
//	     characters specified by input parameter, 'decSeparator'.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (decSeparatorSpec *DecimalSeparatorSpec) NewRunes(
	decSeparator []rune,
	errorPrefix interface{}) (
	newDecimalSeparator DecimalSeparatorSpec,
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
			"NewStr()",
		"")

	if err != nil {
		return newDecimalSeparator, err
	}

	lenDecSepRunes := len(decSeparator)

	if lenDecSepRunes == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decSeparator' is invalid!"+
			"'decSeparator' is an empty array with a length of zero.\n",
			ePrefix.String())

		return newDecimalSeparator, err
	}

	err = newDecimalSeparator.decimalSeparatorChars.SetRuneArray(
		decSeparator,
		ePrefix.XCpy(
			"decSeparatorSpec<-decSepRunes"))

	return newDecimalSeparator, err
}

// NewStr - Creates and returns a new instance of
// DecimalSeparatorSpec populated with the decimal
// separator character or characters passed by input
// parameter string, 'decSeparator'.
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 decSeparator               string
//	   - This string contains the character or characters
//	     which will be configured as the Decimal Separator
//	     Symbol or Symbols for the returned instance of
//	     DecimalSeparatorSpec, the Decimal Separator
//	     Specification.
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	newDecimalSeparator        DecimalSeparatorSpec
//	   - If this method completes successfully, this return
//	     parameter is configured with the decimal separator
//	     characters specified by input parameter, 'decSeparator'.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (decSeparatorSpec *DecimalSeparatorSpec) NewStr(
	decSeparator string,
	errorPrefix interface{}) (
	newDecimalSeparator DecimalSeparatorSpec,
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
			"NewStr()",
		"")

	if err != nil {
		return newDecimalSeparator, err
	}

	decSepRunes := []rune(decSeparator)

	lenDecSepRunes := len(decSepRunes)

	if lenDecSepRunes == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decSeparator' is invalid!"+
			"'decSeparator' is an empty string with a length of zero.\n",
			ePrefix.String())

		return newDecimalSeparator, err
	}

	err = newDecimalSeparator.decimalSeparatorChars.SetRuneArray(
		decSepRunes,
		ePrefix.XCpy(
			"decSeparatorSpec<-decSepRunes"))

	return newDecimalSeparator, err
}

// SearchForDecimalSeparator - Searches a target text character
// string for the presence of a Decimal Separator Symbol or
// Symbols specified by the current instance of
// DecimalSeparatorSpec.
//
// This method is typically called by Number String parsing
// functions.
//
// Number String parsing functions search for a Decimal Separator
// Symbol or Symbols in strings of numeric digits called
// 'Number Strings'. Number String parsing functions review strings
// of text characters containing numeric digits and convert those
// numeric digits to numeric values. If a Decimal Separator Symbol
// or Symbols are located in a Number String the digits to the
// right of the Decimal Separator are considered fractional numeric
// values and the digits to the left are considered integer numeric
// values.
//
// If the Decimal Separator specified by the current instance of
// DecimalSeparatorSpec is located in a Target Search String, this
// method will return a boolean flag signalling success and the last
// zero based string index of the Decimal Separator found in the
// Target Search String.
//
// This method will also configure member variables used as
// internal processing flags for the current instance of
// DecimalSeparatorSpec:
//
//	DecimalSeparatorSpec.foundDecimalSeparatorSymbols
//	DecimalSeparatorSpec.foundDecimalSeparatorIndex
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Before starting the first iteration of a search for a Decimal
// Separator, make certain you clear all the internal processing
// flags for this instance of DecimalSeparatorSpec. Call method:
//
//	DecimalSeparatorSpec.EmptyProcessingFlags()
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//		 targetSearchString           *RuneArrayDto
//		    - A pointer to a RuneArrayDto. Type
//		      RuneArrayDto contains the string of text
//		      characters which will be searched for the presence of a
//		      Decimal Separator Symbol or Symbols specified by the
//		      current instance of DecimalSeparatorSpec.
//
//					  type RuneArrayDto struct {
//		                CharsArray []rune
//					  }
//
//
//	 errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this parameter
//		     to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	foundDecimalSeparatorSymbols bool
//	   - A boolean status flag which signals the search of a target
//	     character string located the Decimal Separator Symbol or
//	     Symbols specified by the current instance of
//	     DecimalSeparatorSpec.
//
//	     If the search operation successfully located the Decimal
//	     Separator Symbol(s), the return parameter will be set to
//	     'true'.
//
//	     If the search operation failed to locate the specified
//	     Decimal Separator Symbol(s), this return parameter will be
//	     set to 'false'.
//
//
//	lastSearchIndex              int
//	   - If the search operation successfully located the specified
//	     Decimal Separator Symbol(s), the return parameter will be
//	     set to the index in the target search string
//	     ('targetSearchString') occupied by the last text character
//	     in the specified Decimal Separator Symbol(s).
//	       Example:
//	         Target Search String: "0123.56"
//	         Decimal Separator Symbol: "."
//	         lastIndex = 4
//
//
//	err                          error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (decSeparatorSpec *DecimalSeparatorSpec) SearchForDecimalSeparator(
	targetInputParms CharSearchTargetInputParametersDto,
	errorPrefix interface{}) (
	CharSearchDecimalSeparatorResultsDto,
	error) {

	if decSeparatorSpec.lock == nil {
		decSeparatorSpec.lock = new(sync.Mutex)
	}

	decSeparatorSpec.lock.Lock()

	defer decSeparatorSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	decimalSearchResults := CharSearchDecimalSeparatorResultsDto{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"DecimalSeparatorSpec."+
			"SearchForDecimalSeparator()",
		"")

	if err != nil {
		return decimalSearchResults, err
	}

	var err2 error
	_,
		err2 = new(decimalSeparatorSpecAtom).
		testValidityOfDecSepSearchSpec(
			decSeparatorSpec,
			ePrefix.XCpy(
				"decSeparatorSpec"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: The current instance of DecimalSeparatorSpec\n"+
			"is invalid. The Search operation has been aborted.\n"+
			"Validation checks returned the following error for this intance of\n"+
			"DecimalSeparatorSpec:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return decimalSearchResults, err
	}

	err = targetInputParms.ValidateTargetParameters(
		ePrefix.XCpy(
			"targetInputParms invalid!"))

	if err != nil {
		return decimalSearchResults, err
	}

	sMechPreon := strMechPreon{}
	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		targetInputParms.TargetString.CharsArray,
		nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input Parameter '%v' is invalid!\n"+
			"This rune array contains invalid characters.\n"+
			"'%v' returned the following validation\n"+
			"error:\n"+
			"%v\n",
			ePrefix.String(),
			targetInputParms.TargetStringName,
			targetInputParms.TargetStringName,
			err2.Error())

		return decimalSearchResults, err
	}

	testInputParms := CharSearchTestInputParametersDto{}.New()
	testInputParms.TestStringLength =
		decSeparatorSpec.decimalSeparatorChars.GetRuneArrayLength()

	// Nothing to do. Already Found the Decimal
	// Separator on a previous cycle.
	if targetInputParms.FoundDecimalSeparatorSymbols == true {

		decimalSearchResults.FoundDecimalSeparatorSymbols = false
		decimalSearchResults.FoundDecimalSepSymbolsOnPreviousSearch =
			true

		decimalSearchResults.TargetStringFirstFoundIndex =
			-1

		decimalSearchResults.TargetStringLastFoundIndex =
			-1

		return decimalSearchResults, err
	}

	testConfigDto := CharSearchTestConfigDto{}.New()
	testConfigDto.TestStringName = "Decimal Separator"
	testConfigDto.TestInputParametersName = "Decimal Separator Search Parameters"
	testConfigDto.TestStringLengthName = "Decimal Separator Chars Length"
	testConfigDto.TestStringStartingIndex = 0
	testConfigDto.TestStringDescription1 = "Decimal Separator"
	testConfigDto.CollectionTestObjIndex = -1
	testConfigDto.NumSymbolClass = NumSymClass.DecimalSeparator()
	testConfigDto.TextCharSearchType = CharSearchType.LinearTargetStartingIndex()

	decimalSearchResults.Empty()
	runeArraySearchResults := CharSearchRuneArrayResultsDto{}

	runeArraySearchResults,
		err =
		decSeparatorSpec.decimalSeparatorChars.SearchForTextCharacterString(
			targetInputParms,
			testConfigDto,
			ePrefix.XCpy(
				"decSeparatorSpec.decimalSeparatorChars"))

	if err != nil {
		return decimalSearchResults, err
	}

	if runeArraySearchResults.FoundSearchTarget {

		decimalSearchResults.LoadRuneArraySearchResults(
			runeArraySearchResults)

		decimalSearchResults.FoundDecimalSeparatorSymbols = true

		decimalSearchResults.NumValueType = NumValType.FloatingPoint()

		if targetInputParms.FoundFirstNumericDigitInNumStr == true {

			decimalSearchResults.NumSymbolLocation = NumSymLocation.Interior()

		} else {

			decimalSearchResults.NumSymbolLocation = NumSymLocation.Before()

		}

		decimalSearchResults.SearchResultsName = "Decimal Separator Search Results"

		err = decimalSearchResults.FoundRuneArrayChars.CopyIn(
			&decSeparatorSpec.decimalSeparatorChars,
			ePrefix.XCpy("decSeparatorSpec.decimalSeparatorChars"))

		if err != nil {
			return decimalSearchResults, err
		}

		decimalSearchResults.SearchResultsFunctionChain = ePrefix.String()

		err = decimalSearchResults.DecimalSeparatorSymbolsSpec.CopyIn(
			decSeparatorSpec,
			ePrefix.XCpy("decSeparatorSpec"))

		if err != nil {
			return decimalSearchResults, err
		}

	}

	return decimalSearchResults, err
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
//
//	DecimalSeparatorSpec.SetDecimalSeparatorStr()
//
// The sole difference between the two methods is that this method
// receives an array of runes as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current DecimalSeparatorSpec instance
// ('decSeparatorSpec') will be deleted and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//		decSeparator               []rune
//			   - This array of runes contains the character or
//	          characters which will be configured as the Decimal
//	          Separator Symbol for the current instance of
//	          DecimalSeparatorSpec, the Decimal Separator
//	          Specification.
//
//	          If 'decSeparator' is submitted as an empty or nil
//	          array, it will be accepted, the current instance of
//	          DecimalSeparatorSpec will be configured with an empty
//	          decimal separator character array, and no error will
//	          be returned.
//
//
//		 errorPrefix                interface{}
//			   - This object encapsulates error prefix text which is
//			     included in all returned error messages. Usually, it
//			     contains the name of the calling method or methods
//			     listed as a method or function chain of execution.
//
//			     If no error prefix information is needed, set this parameter
//			     to 'nil'.
//
//			     This empty interface must be convertible to one of the
//			     following types:
//
//			     1. nil - A nil value is valid and generates an empty
//			        collection of error prefix and error context
//			        information.
//
//			     2. string - A string containing error prefix information.
//
//			     3. []string A one-dimensional slice of strings containing
//			        error prefix information
//
//			     4. [][2]string A two-dimensional slice of strings
//			        containing error prefix and error context information.
//
//			     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//			        from this object will be copied for use in error and
//			        informational messages.
//
//			     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//			        Information from this object will be copied for use in
//			        error and informational messages.
//
//			     7. IBasicErrorPrefix - An interface to a method generating
//			        a two-dimensional slice of strings containing error
//			        prefix and error context information.
//
//			     If parameter 'errorPrefix' is NOT convertible to one of
//			     the valid types listed above, it will be considered
//			     invalid and trigger the return of an error.
//
//			     Types ErrPrefixDto and IBasicErrorPrefix are included in
//			     the 'errpref' software package,
//			     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
			"SetDecimalSeparatorRunes()",
		"")

	if err != nil {
		return err
	}

	lenDecSepRunes := len(decSeparator)

	if lenDecSepRunes == 0 {

		decSeparatorSpec.decimalSeparatorChars.Empty()

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

	err = decSeparatorSpec.decimalSeparatorChars.SetRuneArray(
		decSeparator,
		ePrefix.XCpy("decSeparatorSpec"))

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
//
//	DecimalSeparatorSpec.SetDecimalSeparatorRunes()
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//		decSeparator               string
//			   - This string contains the character or characters which
//			     will be configured as the Decimal Separator Symbol or
//			     Symbols for the current instance of
//			     DecimalSeparatorSpec, the Decimal Separator
//			     Specification.
//
//	          If 'decSeparator' is submitted as an empty string,
//	          it will be accepted, the current instance of
//	          DecimalSeparatorSpec will be configured with an empty
//	          decimal separator character array, and no error will
//	          be returned.
//
//
//		 errorPrefix                interface{}
//			   - This object encapsulates error prefix text which is
//			     included in all returned error messages. Usually, it
//			     contains the name of the calling method or methods
//			     listed as a method or function chain of execution.
//
//			     If no error prefix information is needed, set this parameter
//			     to 'nil'.
//
//			     This empty interface must be convertible to one of the
//			     following types:
//
//			     1. nil - A nil value is valid and generates an empty
//			        collection of error prefix and error context
//			        information.
//
//			     2. string - A string containing error prefix information.
//
//			     3. []string A one-dimensional slice of strings containing
//			        error prefix information
//
//			     4. [][2]string A two-dimensional slice of strings
//			        containing error prefix and error context information.
//
//			     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//			        from this object will be copied for use in error and
//			        informational messages.
//
//			     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//			        Information from this object will be copied for use in
//			        error and informational messages.
//
//			     7. IBasicErrorPrefix - An interface to a method generating
//			        a two-dimensional slice of strings containing error
//			        prefix and error context information.
//
//			     If parameter 'errorPrefix' is NOT convertible to one of
//			     the valid types listed above, it will be considered
//			     invalid and trigger the return of an error.
//
//			     Types ErrPrefixDto and IBasicErrorPrefix are included in
//			     the 'errpref' software package,
//			     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

	if len(decSeparator) == 0 {

		decSeparatorSpec.decimalSeparatorChars.Empty()

		return err

	}

	decSepRunes := []rune(decSeparator)

	new(decimalSeparatorSpecAtom).
		empty(decSeparatorSpec)

	err = decSeparatorSpec.decimalSeparatorChars.SetRuneArray(
		decSepRunes,
		ePrefix.XCpy(
			"decSeparatorSpec<-decSepRunes"))

	return err
}

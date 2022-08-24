package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrFormatSpec - The Number String Format Specification
// contains parameters used to format signed numbers and
// currency numeric values as number strings.
type NumStrFormatSpec struct {
	decSeparator DecimalSeparatorSpec
	// Contains the decimal separator character or
	// characters which will separate integer and
	// fractional digits in a floating point number

	intGroupingSpec NumStrIntegerGroupingSpec
	// Integer Grouping Specification. This parameter
	// specifies the type of integer grouping and
	// integer separator characters which will be
	// applied to the number string formatting
	// operations.

	roundingSpec NumStrRoundingSpec
	// Controls the rounding algorithm applied to
	// floating point numbers.

	positiveNumberSign NumStrNumberSymbolSpec
	// Positive number signs are commonly implied and
	// not specified. However, the user as the option
	// to specify a positive number sign character or
	// characters for positive numeric values using a
	// Number String Positive Number Sign
	// Specification.

	negativeNumberSign NumStrNumberSymbolSpec
	// The Number String Negative Number Sign
	// Specification is used to configure negative
	// number sign symbols for negative numeric values
	// formatted and displayed in number stings.

	numberFieldSpec NumStrNumberFieldSpec
	// This Number String Number Field Specification
	// contains the field length and text justification
	// parameter necessary to display a numeric value
	// within a number field for display as a number
	// string.

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrFormatSpec ('incomingSignedNumFmt')
// to the data fields of the current NumStrFormatSpec
// instance ('signedNumFmtSpec').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
// All the member variable data values in the current
// NumStrFormatSpec instance ('signedNumFmtSpec') will
// be deleted and replaced.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingSignedNumFmt		*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		This method will NOT change the values of internal
//		member variables contained in this instance.
//
//		All data values in this NumStrFormatSpec instance
//		will be copied to current NumStrFormatSpec
//		instance ('signedNumFmtSpec').
//
//		If parameter 'incomingSignedNumFmt' is determined to
//		be invalid, an error will be returned.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) CopyIn(
	incomingSignedNumFmt *NumStrFormatSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		copySignedNumberFormatSpec(
			numStrFmtSpec,
			incomingSignedNumFmt,
			ePrefix.XCpy(
				"numStrFmtSpec<-"+
					"incomingSignedNumFmt"))
}

// CopyOut - Returns a deep copy of the current
// NumStrFormatSpec instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopySignedNumFmtSpec	NumStrFormatSpec
//
//		If this method completes successfully and no errors are
//		encountered, this parameter will return a deep copy of
//		the current NumStrFormatSpec instance.
//
//
//	err							error
//
//		If the method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopySignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopySignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		copySignedNumberFormatSpec(
			&deepCopySignedNumFmtSpec,
			numStrFmtSpec,
			ePrefix.XCpy(
				"deepCopySignedNumFmtSpec<-"+
					"numStrFmtSpec"))

	return deepCopySignedNumFmtSpec, err
}

// Empty - Resets all internal member variables for the current
// instance of NumStrFormatSpec to their initial or zero
// values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of
// NumStrFormatSpec.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (numStrFmtSpec *NumStrFormatSpec) Empty() {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	new(numStrFmtSpecAtom).empty(
		numStrFmtSpec)

	numStrFmtSpec.lock.Unlock()

	numStrFmtSpec.lock = nil

}

// Equal - Receives a pointer to another instance of
// NumStrFormatSpec and proceeds to compare its
// internal member variables to those of the current
// NumStrFormatSpec instance in order to determine
// if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are
// equal in all respects, this flag is set to 'true'.
// Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingSignedNumFmt		*NumStrFormatSpec
//
//		A pointer to an external instance of
//		NumStrFormatSpec. The internal member variable
//		data values in this instance will be compared to those
//		in the current instance of NumStrFormatSpec. The
//		results of this comparison will be returned to the
//		calling function as a boolean value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//
//		If the internal member variable data values contained in
//		input parameter 'incomingSignedNumFmt' are equivalent
//		in all respects to those contained in the current
//		instance of 'NumStrFormatSpec', this return value
//		will be set to 'true'.
//
//		Otherwise, this method will return 'false'.
func (numStrFmtSpec *NumStrFormatSpec) Equal(
	incomingSignedNumFmt *NumStrFormatSpec) bool {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return new(numStrFmtSpecAtom).equal(
		numStrFmtSpec,
		incomingSignedNumFmt)
}

// GetDecSeparatorRunes - Returns an array of runes containing the
// Decimal Separator character or characters configured for the
// current instance of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	[]rune
//
//		An array of runes containing the Decimal Separator character
//		or characters configured for the current instance of
//		NumStrFormatSpec.
//
//		If Decimal Separator character(s) have not yet been
//		configured, this method will return 'nil'.
func (numStrFmtSpec *NumStrFormatSpec) GetDecSeparatorRunes() []rune {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return numStrFmtSpec.decSeparator.GetDecimalSeparatorRunes()
}

// GetDecSeparatorSpec - Returns a deep copy of the Decimal
// Separator Specification configured for the current instance
// of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	DecimalSeparatorSpec
//
//		If this method completes successfully, a deep copy of
//		the Decimal Separator Specification configured for the
//		current instance of NumStrFormatSpec will be
//		returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetDecSeparatorSpec(
	errorPrefix interface{}) (
	DecimalSeparatorSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetDecSeparatorSpec()",
		"")

	if err != nil {
		return DecimalSeparatorSpec{}, err
	}

	return numStrFmtSpec.decSeparator.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.decSeparator"))
}

// GetDecSeparatorStr - Returns a string containing the Decimal
// Separator character or characters configured for the current
// instance of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		A string containing the Decimal Separator character or
//	 	characters configured for the current instance of
//	 	NumStrFormatSpec.
//
//		If Decimal Separator character(s) have not yet been
//		configured, this method will return an empty string.
func (numStrFmtSpec *NumStrFormatSpec) GetDecSeparatorStr() string {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return numStrFmtSpec.decSeparator.GetDecimalSeparatorStr()
}

// GetIntGroupingSpec - Returns a deep copy of the Integer
// Grouping Specification configured for the current instance
// of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrIntegerGroupingSpec
//
//		If this method completes successfully, a deep copy of
//		the Integer Grouping Specification configured for the
//		current instance of NumStrFormatSpec will be
//		returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetIntGroupingSpec(
	errorPrefix interface{}) (
	NumStrIntegerGroupingSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetIntGroupingSpec()",
		"")

	if err != nil {
		return NumStrIntegerGroupingSpec{}, err
	}

	return numStrFmtSpec.intGroupingSpec.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.intGroupingSpec"))
}

// GetIntSeparatorChars - Returns a string containing the Integer
// Separator character or characters configured for the current
// instance of NumStrFormatSpec.
//
// Integer Separator Characters consist of one or more text
// characters used to separate groups of integers. This
// separator is also known as the 'thousands' separator in
// the United States. It is used to separate groups of integer
// digits to the left of the decimal separator (a.k.a. decimal
// point). In the United States, the standard integer digits
// separator is the comma (",").
//
//	United States Example:  1,000,000,000
//
// In many European countries, a single period (".") is used
// as the integer separator character.
//
//	European Example: 1.000.000.000
//
// Other countries and cultures use spaces, apostrophes or
// multiple characters to separate integers.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method will return a string containing the Integer
//		Separator character or characters configured for the
//	 	current instance of NumStrFormatSpec.
func (numStrFmtSpec *NumStrFormatSpec) GetIntSeparatorChars() string {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return numStrFmtSpec.intGroupingSpec.GetIntegerSeparatorChars()
}

// GetIntegerSeparatorDto - Returns an instance of
// IntegerSeparatorDto based on the configuration parameters
// contained within the current instance of
// NumStrFormatSpec.
//
// IntegerSeparatorDto is used by low level number string
// formatting functions to complete the number string generation
// operation.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	IntegerSeparatorDto
//
//		If this method completes successfully, a copy
//		of the integer grouping specification configured
//		for the current NumStrFormatSpec instance
//		will be returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetIntegerSeparatorDto(
	errorPrefix interface{}) (
	IntegerSeparatorDto,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetIntegerSeparatorDto()",
		"")

	if err != nil {
		return IntegerSeparatorDto{}, err
	}

	return numStrFmtSpec.intGroupingSpec.GetIntegerSeparatorDto(
		ePrefix.XCpy(
			"<-numStrFmtSpec.intGroupingSpec"))
}

// GetIntSeparatorRunes - Returns a rune array containing the
// Integer Separator character or characters configured for the
// current instance of NumStrFormatSpec.
//
// Integer Separator Characters consist of one or more text
// characters used to separate groups of integers. This
// separator is also known as the 'thousands' separator in
// the United States. It is used to separate groups of integer
// digits to the left of the decimal separator (a.k.a. decimal
// point). In the United States, the standard integer digits
// separator is the comma (',').
//
//	United States Example:  1,000,000,000
//
// In many European countries, a single period ('.') is used
// as the integer separator character.
//
//	European Example: 1.000.000.000
//
// Other countries and cultures use spaces, apostrophes or
// multiple characters to separate integers.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	[]rune
//
//		This method will return a rune array containing the
//		Integer Separator character or characters configured
//		for the	current instance of NumStrFormatSpec.
func (numStrFmtSpec *NumStrFormatSpec) GetIntSeparatorRunes() []rune {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	return numStrFmtSpec.intGroupingSpec.GetIntegerSeparatorRunes()
}

// GetNegativeNumSymSpec - Returns the Negative Number Symbol
// Specification currently configured for this instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrNumberSymbolSpec
//
//		If this method completes successfully, a copy
//		of the negative number sign specification configured
//		for the current NumStrFormatSpec instance
//		will be returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetNegativeNumSymSpec(
	errorPrefix interface{}) (
	NumStrNumberSymbolSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetNegativeNumSymSpec()",
		"")

	if err != nil {
		return NumStrNumberSymbolSpec{}, err
	}

	return numStrFmtSpec.negativeNumberSign.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.negativeNumberSign"))
}

// GetNumberFieldSpec - Returns the Number Field Specification
// currently configured for this instance of
// NumStrFormatSpec.
//
// The Number Field Specification includes parameters for
// field length and text justification ('Right, Center,
// Left').
//
// Numeric digits are formatted using the Number Field
// Specification within a number field specified by field
// length. The numeric digits string is then justified
// within the number field according to a text justification
// specification of 'Right', 'Center' or 'Left'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrNumberFieldSpec
//
//		If this method completes successfully, a copy
//		of the Number Field Specification configured
//		for the current NumStrFormatSpec instance
//		will be returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetNumberFieldSpec(
	errorPrefix interface{}) (
	NumStrNumberFieldSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetNumberFieldSpec()",
		"")

	if err != nil {
		return NumStrNumberFieldSpec{}, err
	}

	return numStrFmtSpec.numberFieldSpec.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.numberFieldSpec"))
}

// GetPositiveNumSymSpec - Returns the Positive Number Symbol
// Specification currently configured for this instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrNumberSymbolSpec
//
//		If this method completes successfully, a copy
//		of the positive number sign specification configured
//		for the current NumStrFormatSpec instance
//		will be returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetPositiveNumSymSpec(
	errorPrefix interface{}) (
	NumStrNumberSymbolSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetPositiveNumSymSpec()",
		"")

	if err != nil {
		return NumStrNumberSymbolSpec{}, err
	}

	return numStrFmtSpec.positiveNumberSign.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.positiveNumberSign"))
}

// GetRoundingSpec - Returns the Rounding Specification
// configured for the current instance of
// NumStrFormatSpec.
//
// The Rounding Specification determines if rounding will be
// applied to the fractional digits in a number string as
// well as the type of rounding algorithm applied and the
// number of fractional digits remaining after the rounding
// operation is completed.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrRoundingSpec
//
//		If this method completes successfully, a copy
//		of the Rounding Specification configured for
//		the current NumStrFormatSpec instance
//		will be returned.
//
//	error
//
//		If this method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (numStrFmtSpec *NumStrFormatSpec) GetRoundingSpec(
	errorPrefix interface{}) (
	NumStrRoundingSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetRoundingSpec()",
		"")

	if err != nil {
		return NumStrRoundingSpec{}, err
	}

	return numStrFmtSpec.roundingSpec.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.roundingSpec"))
}

// NewSignedNumFmtParams - Creates and returns a new instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		DecimalSeparatorSpec, the Decimal Separator
//		Specification.
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the 'thousands'
//		separator. It is used to separate groups of integer digits
//		to the left of the decimal separator
//		(a.k.a. decimal point). In the United States, the standard
//		integer digits separator is the comma (",").
//		United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//		European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length string
//		and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	intGroupingType				IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorDto which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	roundingType				NumberRoundingType
//
//		This parameter will populate the
//		'NumStrRoundingSpec.roundingType' member variable
//		data value contained in returned instance of
//		NumStrRoundingSpec.
//
//		NumberRoundingType is an enumeration specifying the
//		rounding algorithm to be applied in the fractional
//		digit rounding operation. Valid values are listed
//		as follows:
//
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		decimal separator (a.k.a. decimal point) which will
//		remain after completion of the number rounding
//		operation.
//
//	leadingPositiveNumSign		string
//
//		A string containing the leading positive number sign
//		character or characters used to configure a Positive
//		Number Sign Symbol in a number string.
//
//		Leading number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//		currency symbols ('$').
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " +123.45"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45+"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	trailingPositiveNumSign     string
//
//		A string containing the trailing positive number sign
//		character or characters used to configure a Positive
//		Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	leadingNegativeNumSign		string
//
//		A string containing the leading negative number sign
//		character or characters used to configure a Negative
//		Number Sign Symbol in a number string.
//
//		Leading number symbols can include any combination of
//		characters to include minus signs ('-') and/or
//	 	currency symbols ('$').
//
//	trailingNegativeNumSign		string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure a
//		Negative Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any	combination
//		of characters to include minus signs ('-') and/or
//		currency symbols ('$').
//
//	negativeNumFieldSymPosition				NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " -123.45"
//					Number Field Index:  01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	fieldLength					int
//
//		This parameter defines the length of the text field in
//		which the numeric value will be displayed within a
//		number string.
//
//		If 'fieldLength' is less than the length of the numeric
//		value string, it will be automatically set equal to the
//		length of that numeric value string.
//
//		To automatically set the value of fieldLength to the string
//		length of the numeric value, set this parameter to a value
//		of minus one (-1).
//
//		If this parameter is submitted with a value less than minus
//		one (-1) or greater than 1-million (1,000,000), an error will
//		be returned.
//
//	fieldLength					int
//
//		This parameter defines the length of the text field in
//		which the numeric value will be displayed within a
//		number string.
//
//		If 'fieldLength' is less than the length of the numeric
//		value string, it will be automatically set equal to the
//		length of that numeric value string.
//
//		To automatically set the value of fieldLength to the string
//		length of the numeric value, set this parameter to a value
//		of minus one (-1).
//
//		If this parameter is submitted with a value less than minus
//		one (-1) or greater than 1-million (1,000,000), an error will
//		be returned.
//
//	fieldJustification			TextJustify
//
//		An enumeration which specifies the justification of the
//		numeric value within the number field length specified
//		by input parameter 'fieldLength'.
//
//		Text justification can only be evaluated in the context of
//		a number string, field length and a 'textJustification'
//		object of type TextJustify. This is because number strings
//		with a field length equal to or less than the length of the
//		numeric value string never use text justification. In these
//		cases, text justification is completely ignored.
//
//		If the field length parameter ('fieldLength') is greater
//		than the length of the numeric value string, text
//		justification must be equal to one of these
//		three valid values:
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumFmtParams(
	decSeparatorChars string,
	intGroupingChars string,
	intGroupingType IntegerGroupingType,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	leadingPosNumSign string,
	trailingPosNumSign string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign string,
	trailingNegNumSign string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumFmtParams()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setNStrNumberFieldSpec(
			&newSignedNumFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intGroupingChars),
			intGroupingType,
			roundingType,
			roundToFractionalDigits,
			[]rune(leadingPosNumSign),
			[]rune(trailingPosNumSign),
			positiveNumFieldSymPosition,
			[]rune(leadingNegNumSign),
			[]rune(trailingNegNumSign),
			negativeNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

// NewSignedNumFmtParamsRunes - Creates and returns a new instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator				[]rune
//
//		This rune array contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		DecimalSeparatorSpec, the Decimal Separator
//		Specification.
//
//	intSeparatorChars			[]rune
//
//		A rune array containing one or more characters used to
//		separate groups of integers. This separator is also known
//	 	as the 'thousands' separator in the United States. It is
//	 	used to separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In the
//	  	United States, the standard	integer digits separator is
//	  	the comma (",").
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length rune
//		array and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	intGroupingType				IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorDto which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	roundingType				NumberRoundingType
//
//		This parameter will populate the
//		'NumStrRoundingSpec.roundingType' member variable
//		data value contained in returned instance of
//		NumStrRoundingSpec.
//
//		NumberRoundingType is an enumeration specifying the
//		rounding algorithm to be applied in the fractional
//		digit rounding operation. Valid values are listed
//		as follows:
//
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		decimal separator (a.k.a. decimal point) which will
//		remain after completion of the number rounding
//		operation.
//
//	leadingPositiveNumSign		[]rune
//
//		A rune array containing the leading positive number
//		sign character or characters used to configure a
//	 	Positive Number Sign Symbol in a number string.
//
//		Leading number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//		currency symbols ('$').
//
//	trailingPositiveNumSign     []rune
//
//		A rune array containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " +123.45"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45+"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	leadingNegativeNumSign		[]rune
//
//		A rune array containing the leading negative number
//		sign character or characters used to configure a
//		Negative Number Sign Symbol in a number string.
//
//		Leading number symbols can include any combination of
//		characters to include minus signs ('-') and/or
//	 	currency symbols ('$').
//
//	trailingNegativeNumSign		[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to configure
//		a Negative Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any	combination
//		of characters to include minus signs ('-') and/or
//		currency symbols ('$').
//
//	negativeNumFieldSymPosition				NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " -123.45"
//					Number Field Index:  01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	fieldLength					int
//
//		This parameter defines the length of the text field in
//		which the numeric value will be displayed within a
//		number string.
//
//		If 'fieldLength' is less than the length of the numeric
//		value string, it will be automatically set equal to the
//		length of that numeric value string.
//
//		To automatically set the value of fieldLength to the string
//		length of the numeric value, set this parameter to a value
//		of minus one (-1).
//
//		If this parameter is submitted with a value less than minus
//		one (-1) or greater than 1-million (1,000,000), an error will
//		be returned.
//
//	fieldJustification			TextJustify
//
//		An enumeration which specifies the justification of the
//		numeric value within the number field length specified
//		by input parameter 'fieldLength'.
//
//		Text justification can only be evaluated in the context of
//		a number string, field length and a 'textJustification'
//		object of type TextJustify. This is because number strings
//		with a field length equal to or less than the length of the
//		numeric value string never use text justification. In these
//		cases, text justification is completely ignored.
//
//		If the field length parameter ('fieldLength') is greater
//		than the length of the numeric value string, text
//		justification must be equal to one of these
//		three valid values:
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec.
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumFmtParamsRunes(
	decSeparatorChars []rune,
	intGroupingChars []rune,
	intGroupingType IntegerGroupingType,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	newSignedNumFmtSpec NumStrFormatSpec,
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"NewSignedNumFmtParamsRunes()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setNStrNumberFieldSpec(
			&newSignedNumFmtSpec,
			decSeparatorChars,
			intGroupingChars,
			intGroupingType,
			roundingType,
			roundToFractionalDigits,
			leadingPosNumSign,
			trailingPosNumSign,
			positiveNumFieldSymPosition,
			leadingNegNumSign,
			trailingNegNumSign,
			negativeNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

// SetDecimalSeparatorSpec - Deletes and replaces the Decimal
// Separator Specification for the current instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparatorSpec			*DecimalSeparatorSpec
//
//		An instance of DecimalSeparatorSpec. The member
//		variable data values contained in this instance
//		will be copied to the current
//		NumStrFormatSpec member variable:
//			'NumStrFormatSpec.decSeparator'.
//
//		In the United States, the decimal separator is
//		referred to as the decimal point.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetDecimalSeparatorSpec(
	decSeparatorSpec DecimalSeparatorSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetDecimalSeparatorSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setDecimalSeparatorSpec(
		numStrFmtSpec,
		decSeparatorSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-decSeparatorSpec"))
}

// SetIntegerGroupingSpec - Deletes and replaces the Integer
// Grouping Specification for the current instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	intGroupingSpec				*NumStrIntegerGroupingSpec
//
//		An instance of NumStrIntegerGroupingSpec. The member
//		variable data values contained in this instance
//		will be copied to the current instance of
//		NumStrFormatSpec :
//			'NumStrFormatSpec.intGroupingSpec'.
//
//		In the United States, the Integer Group Specification
//		character is a comma (',') and integer grouped in
//		thousands.
//			United States Example: 1,000,000,000
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetIntegerGroupingSpec(
	intGroupingSpec NumStrIntegerGroupingSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetIntegerGroupingSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setIntegerGroupingSpec(
		numStrFmtSpec,
		intGroupingSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-intGroupingSpec"))
}

// SetNegativeNumberSignSpec - Deletes and replaces the Negative
// Number Sign Specification for the current instance of
// NumStrFormatSpec:
//
//	NumStrFormatSpec.negativeNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to the current
//		NumStrFormatSpec member variable:
//			'NumStrFormatSpec.negativeNumberSign'.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetNegativeNumberSignSpec(
	negativeNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetNegativeNumberSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).
		setNegativeNumberSignSpec(
			numStrFmtSpec,
			negativeNumberSign,
			ePrefix.XCpy(
				"numStrFmtSpec<-"+
					"negativeNumberSign"))
}

// SetNumberFieldSpec - Deletes and replaces the Number Field
// Specification for the current instance of
// NumStrFormatSpec.
//
// The Number Field Specification includes parameters for
// field length and text justification ('Right, Center,
// Left').
//
// Numeric digits are formatted using the Number Field
// Specification within a number field specified by field
// length. The numeric digits string is then justified
// within the number field according to a text justification
// specification of 'Right', 'Center' or 'Left'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		An instance of NumStrIntegerGroupingSpec. The member
//		variable data values contained in this instance
//		will be copied to the current instance of
//		NumStrFormatSpec :
//			'NumStrFormatSpec.numberFieldSpec'.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetNumberFieldSpec(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setNumberFieldSpec(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-numberFieldSpec"))
}

// SetPositiveNumberSignSpec - Deletes and replaces the Positive
// Number Sign Specification for the current instance of
// NumStrFormatSpec:
//
//	NumStrFormatSpec.positiveNumberSign
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to the current
//		NumStrFormatSpec member variable:
//			'NumStrFormatSpec.positiveNumberSign'.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetPositiveNumberSignSpec(
	positiveNumberSign NumStrNumberSymbolSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetPositiveNumberSignSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).
		setPositiveNumberSignSpec(
			numStrFmtSpec,
			positiveNumberSign,
			ePrefix.XCpy(
				"numStrFmtSpec<-"+
					"positiveNumberSign"))
}

// SetRoundingSpec - Deletes and replaces the Rounding
// Specification contained in the current instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	roundingSpec				NumStrRoundingSpec
//
//		An instance of NumStrRoundingSpec. All data values in
//		this NumStrRoundingSpec instance will be copied to the
//		member variable 'NumStrFormatSpec.roundingSpec'
//		contained in the current instance of
//		NumStrFormatSpec.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetRoundingSpec(
	roundingSpec NumStrRoundingSpec,
	errorPrefix interface{}) error {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetRoundingSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).setRoundingSpec(
		numStrFmtSpec,
		roundingSpec,
		ePrefix.XCpy(
			""))
}

// SetSignedNumFmtParams - Deletes and resets all the member variable
// data values stored in the current instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields contained in the current
// instance of  will be deleted and replaced by values
// generated from the listed input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator				string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		DecimalSeparatorSpec, the Decimal Separator
//		Specification.
//
//	intSeparatorChars			string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the 'thousands'
//		separator. It is used to separate groups of integer digits
//		to the left of the decimal separator
//		(a.k.a. decimal point). In the United States, the standard
//		integer digits separator is the comma (",").
//		United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//		European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length string
//		and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	intGroupingType				IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorDto which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	roundingType				NumberRoundingType
//
//		This parameter will populate the
//		'NumStrRoundingSpec.roundingType' member variable
//		data value contained in returned instance of
//		NumStrRoundingSpec.
//
//		NumberRoundingType is an enumeration specifying the
//		rounding algorithm to be applied in the fractional
//		digit rounding operation. Valid values are listed
//		as follows:
//
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		decimal separator (a.k.a. decimal point) which will
//		remain after completion of the number rounding
//		operation.
//
//	leadingPositiveNumSign		string
//
//		A string containing the leading positive number sign
//		character or characters used to configure a Positive
//		Number Sign Symbol in a number string.
//
//		Leading number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//		currency symbols ('$').
//
//	trailingPositiveNumSign     string
//
//		A string containing the trailing positive number sign
//		character or characters used to configure a Positive
//		Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " +123.45"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45+"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	leadingNegativeNumSign		string
//
//		A string containing the leading negative number sign
//		character or characters used to configure a Negative
//		Number Sign Symbol in a number string.
//
//		Leading number symbols can include any combination of
//		characters to include minus signs ('-') and/or
//	 	currency symbols ('$').
//
//	trailingNegativeNumSign		string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure a
//		Negative Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any	combination
//		of characters to include minus signs ('-') and/or
//		currency symbols ('$').
//
//	negativeNumFieldSymPosition				NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " -123.45"
//					Number Field Index:  01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	fieldLength					int
//
//		This parameter defines the length of the text field in
//		which the numeric value will be displayed within a
//		number string.
//
//		If 'fieldLength' is less than the length of the numeric
//		value string, it will be automatically set equal to the
//		length of that numeric value string.
//
//		To automatically set the value of fieldLength to the string
//		length of the numeric value, set this parameter to a value
//		of minus one (-1).
//
//		If this parameter is submitted with a value less than minus
//		one (-1) or greater than 1-million (1,000,000), an error will
//		be returned.
//
//	fieldJustification			TextJustify
//
//		An enumeration which specifies the justification of the
//		numeric value within the number field length specified
//		by input parameter 'fieldLength'.
//
//		Text justification can only be evaluated in the context of
//		a number string, field length and a 'textJustification'
//		object of type TextJustify. This is because number strings
//		with a field length equal to or less than the length of the
//		numeric value string never use text justification. In these
//		cases, text justification is completely ignored.
//
//		If the field length parameter ('fieldLength') is greater
//		than the length of the numeric value string, text
//		justification must be equal to one of these
//		three valid values:
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumFmtParams(
	decSeparatorChars string,
	intGroupingChars string,
	intGroupingType IntegerGroupingType,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	leadingPosNumSign string,
	trailingPosNumSign string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign string,
	trailingNegNumSign string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumFmtParams()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrFmtSpecNanobot).
		setNStrNumberFieldSpec(
			numStrFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intGroupingChars),
			intGroupingType,
			roundingType,
			roundToFractionalDigits,
			[]rune(leadingPosNumSign),
			[]rune(trailingPosNumSign),
			positiveNumFieldSymPosition,
			[]rune(leadingNegNumSign),
			[]rune(trailingNegNumSign),
			negativeNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return err
}

// SetSignedNumFmtSpecRunes - Deletes and resets all the member variable
// data values stored in the current instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields contained in the current
// instance of  will be deleted and replaced by values
// generated from the listed input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator				[]rune
//
//		This rune array contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		DecimalSeparatorSpec, the Decimal Separator
//		Specification.
//
//	intSeparatorChars			[]rune
//
//		A rune array containing one or more characters used to
//		separate groups of integers. This separator is also known
//	 	as the 'thousands' separator in the United States. It is
//	 	used to separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In the
//	  	United States, the standard	integer digits separator is
//	  	the comma (",").
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length rune
//		array and 'intGroupingType' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	intGroupingType				IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorDto which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	roundingType				NumberRoundingType
//
//		This parameter will populate the
//		'NumStrRoundingSpec.roundingType' member variable
//		data value contained in returned instance of
//		NumStrRoundingSpec.
//
//		NumberRoundingType is an enumeration specifying the
//		rounding algorithm to be applied in the fractional
//		digit rounding operation. Valid values are listed
//		as follows:
//
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		decimal separator (a.k.a. decimal point) which will
//		remain after completion of the number rounding
//		operation.
//
//	leadingPositiveNumSign		[]rune
//
//		A rune array containing the leading positive number
//		sign character or characters used to configure a
//	 	Positive Number Sign Symbol in a number string.
//
//		Leading number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//		currency symbols ('$').
//
//	trailingPositiveNumSign     []rune
//
//		A rune array containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any combination
//		of characters to include plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " +123.45"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45+"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	leadingNegativeNumSign		[]rune
//
//		A rune array containing the leading negative number
//		sign character or characters used to configure a
//		Negative Number Sign Symbol in a number string.
//
//		Leading number symbols can include any combination of
//		characters to include minus signs ('-') and/or
//	 	currency symbols ('$').
//
//	trailingNegativeNumSign		[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to configure
//		a Negative Number Sign Symbol in a number string.
//
//		Trailing number symbols can include any	combination
//		of characters to include minus signs ('-') and/or
//		currency symbols ('$').
//
//	negativeNumFieldSymPosition				NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " -123.45"
//					Number Field Index:  01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	fieldLength					int
//
//		This parameter defines the length of the text field in
//		which the numeric value will be displayed within a
//		number string.
//
//		If 'fieldLength' is less than the length of the numeric
//		value string, it will be automatically set equal to the
//		length of that numeric value string.
//
//		To automatically set the value of fieldLength to the string
//		length of the numeric value, set this parameter to a value
//		of minus one (-1).
//
//		If this parameter is submitted with a value less than minus
//		one (-1) or greater than 1-million (1,000,000), an error will
//		be returned.
//
//	fieldJustification			TextJustify
//
//		An enumeration which specifies the justification of the
//		numeric value within the number field length specified
//		by input parameter 'fieldLength'.
//
//		Text justification can only be evaluated in the context of
//		a number string, field length and a 'textJustification'
//		object of type TextJustify. This is because number strings
//		with a field length equal to or less than the length of the
//		numeric value string never use text justification. In these
//		cases, text justification is completely ignored.
//
//		If the field length parameter ('fieldLength') is greater
//		than the length of the numeric value string, text
//		justification must be equal to one of these
//		three valid values:
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix
//			information.
//
//		3. []string A one-dimensional slice of strings
//			containing error prefix information.
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context
//		   information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto.
//			Information from this object will be copied for use
//			in error and informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of
//			ErrPrefixDto. Information from this object will be
//			copied for use in error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method
//			generating a two-dimensional slice of strings
//			containing error prefix and error context
//			information.
//
//		If parameter 'errorPrefix' is NOT convertible to one
//		of the valid types listed above, it will be
//		considered invalid and trigger the return of an
//		error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included
//		in the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumFmtSpecRunes(
	decSeparatorChars []rune,
	intGroupingChars []rune,
	intGroupingType IntegerGroupingType,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errorPrefix interface{}) (
	err error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"SetSignedNumFmtSpecRunes()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrFmtSpecNanobot).
		setNStrNumberFieldSpec(
			numStrFmtSpec,
			decSeparatorChars,
			intGroupingChars,
			intGroupingType,
			roundingType,
			roundToFractionalDigits,
			leadingPosNumSign,
			trailingPosNumSign,
			positiveNumFieldSymPosition,
			leadingNegNumSign,
			trailingNegNumSign,
			negativeNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return err
}

// numStrFmtSpecNanobot - This type provides
// helper methods for NumStrFormatSpec
type numStrFmtSpecNanobot struct {
	lock *sync.Mutex
}

// copySignedNumberFormatSpec - Copies all data from input parameter
// 'sourceSignedNumFmtSpec' to input parameter
// 'destinationSignedNumFmtSpec'. Both instances are of type
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields in
// 'destinationSignedNumFmtSpec' will be deleted and overwritten.
//
// Also, NO data validation is performed on 'sourceSignedNumFmtSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationSignedNumFmtSpec		*NumStrFormatSpec
//		A pointer to a NumStrFormatSpec instance.
//		All the member variable data fields in this object will be
//		replaced by data values copied from input parameter
//		'sourceSignedNumFmtSpec'.
//
//		'destinationSignedNumFmtSpec' is the destination for this
//		copy operation.
//
//	sourceSignedNumFmtSpec			*NumStrFormatSpec
//		A pointer to another NumStrFormatSpec
//		instance. All the member variable data values from this
//		object will be copied to corresponding member variables in
//		'destinationSignedNumFmtSpec'.
//
//		'sourceSignedNumFmtSpec' is the source for this copy
//		operation.
//
//		No data validation is performed on 'sourceSignedNumFmtSpec'.
//
//	errPrefDto						*ePref.ErrPrefixDto
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) copySignedNumberFormatSpec(
	destinationSignedNumFmtSpec *NumStrFormatSpec,
	sourceSignedNumFmtSpec *NumStrFormatSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumberFieldSpecNanobot.lock == nil {
		nStrNumberFieldSpecNanobot.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"copySignedNumberFormatSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationSignedNumFmtSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationSignedNumFmtSpec' is invalid!\n"+
			"'destinationSignedNumFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceSignedNumFmtSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceSignedNumFmtSpec' is invalid!\n"+
			"'sourceSignedNumFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrFmtSpecAtom).empty(
		destinationSignedNumFmtSpec)

	err = destinationSignedNumFmtSpec.decSeparator.CopyIn(
		&sourceSignedNumFmtSpec.decSeparator,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.decSeparator"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.intGroupingSpec.CopyIn(
		&sourceSignedNumFmtSpec.intGroupingSpec,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.intGroupingSpec"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.roundingSpec.CopyIn(
		&sourceSignedNumFmtSpec.roundingSpec,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.roundingSpec"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.positiveNumberSign.CopyIn(
		&sourceSignedNumFmtSpec.positiveNumberSign,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.positiveNumberSign"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.negativeNumberSign.CopyIn(
		&sourceSignedNumFmtSpec.negativeNumberSign,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.negativeNumberSign"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.numberFieldSpec.CopyIn(
		&sourceSignedNumFmtSpec.numberFieldSpec,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.numberFieldSpec"+
				"<-sourceSignedNumFmtSpec"))

	return err
}

// setNStrNumberFieldSpec - Deletes and resets the member variable
// data values stored in the instance of NumStrFormatSpec
// passed as input parameter 'signedNumFmtSpec'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields contained in input parameter
// 'signedNumFmtSpec' will be deleted and replaced by values
// generated from the listed input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmtSpec						*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance.
//		All the member variable data fields in this object will be
//		replaced by data values configured from the input parameter
//		described below.
//
//		'destinationSignedNumFmtSpec' is the destination for this
//		copy operation.
//
//
//	decSeparator							[]rune
//
//		This rune array contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		DecimalSeparatorSpec, the Decimal Separator
//		Specification.
//
//			Example: []rune('.')
//
//	intSeparatorChars						[]rune
//
//		A rune array containing one or more characters used to
//		separate groups of integers. This separator is also known
//	 	as the 'thousands' separator in the United States. It is
//	 	used to separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In the
//	  	United States, the standard	integer digits separator is
//	  	the comma (",").
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length rune
//		array and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	intGroupingType							IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorDto which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
//
//	roundingType							NumberRoundingType
//
//		This parameter will populate the
//		'NumStrRoundingSpec.roundingType' member variable
//		data value contained in returned instance of
//		NumStrRoundingSpec.
//
//		NumberRoundingType is an enumeration specifying the
//		rounding algorithm to be applied in the fractional
//		digit rounding operation. Valid values are listed
//		as follows:
//
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//	roundToFractionalDigits					int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the
//		decimal separator (a.k.a. decimal point) which will
//		remain after completion of the number rounding
//		operation.
//
//	leadingPositiveNumSign					[]rune
//
//		A rune array containing the leading positive number
//		sign character or characters used to configure a
//	 	Positive Number Sign Symbol in a number string.
//
//	trailingPositiveNumSign     			[]rune
//
//		A rune array containing the trailing positive number
//	 	sign character or characters used to configure a
//	  	Positive Number Sign Symbol in a number string.
//
//	positiveNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Positive Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " +123.45"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing plus sign ('+')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45+"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-3:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "+  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing plus sign ('+')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45+"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	leadingNegativeNumSign					[]rune
//
//		A rune array containing the leading negative number
//		sign character or characters used to configure a
//		Negative Number Sign Symbol in a number string.
//
//	trailingNegativeNumSign					[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to configure
//		a Negative Number Sign Symbol in a number string.
//
//	negativeNumFieldSymPosition				NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number string
//		is displayed. Possible valid values are listed as
//		follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " -123.45"
//					Number Field Index:  01234567
//					Total Number String Length: 8
//
//				Example-2:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Right
//					Formatted Number String: " 123.45-"
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				Example-3:
//					Number Field Length: 10
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Inside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: " (123.45) "
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.InsideNumField()' specification,
//				the final length of the number string is defined by the
//				Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-4:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-5:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-6:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     	Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				For the 'NumFieldSymPos.OutsideNumField()' specification,
//				the final length of the number string is greater than
//				the Number Field length.
//
//	fieldLength								int
//
//		This parameter defines the length of the text field in
//		which the numeric value will be displayed within a
//		number string.
//
//		If 'fieldLength' is less than the length of the numeric
//		value string, it will be automatically set equal to the
//		length of that numeric value string.
//
//		To automatically set the value of fieldLength to the string
//		length of the numeric value, set this parameter to a value
//		of minus one (-1).
//
//		If this parameter is submitted with a value less than minus
//		one (-1) or greater than 1-million (1,000,000), an error will
//		be returned.
//
//	fieldJustification						TextJustify
//
//		An enumeration which specifies the justification of the
//		numeric value within the number field length specified
//		by input parameter 'fieldLength'.
//
//		Text justification can only be evaluated in the context of
//		a number string, field length and a 'textJustification'
//		object of type TextJustify. This is because number strings
//		with a field length equal to or less than the length of the
//		numeric value string never use text justification. In these
//		cases, text justification is completely ignored.
//
//		If the field length parameter ('fieldLength') is greater
//		than the length of the numeric value string, text
//		justification must be equal to one of these
//		three valid values:
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	errorPrefix								interface{}
//
//		This object encapsulates error prefix text which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods
//		listed as a method or function chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of the
//		following types:
//
//		1. nil - A nil value is valid and generates an empty
//		   collection of error prefix and error context
//		   information.
//
//		2. string - A string containing error prefix information.
//
//		3. []string A one-dimensional slice of strings containing
//		   error prefix information
//
//		4. [][2]string A two-dimensional slice of strings
//		   containing error prefix and error context information.
//
//		5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		   from this object will be copied for use in error and
//		   informational messages.
//
//		6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		   Information from this object will be copied for use in
//		   error and informational messages.
//
//		7. IBasicErrorPrefix - An interface to a method generating
//		   a two-dimensional slice of strings containing error
//		   prefix and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible to one of
//		the valid types listed above, it will be considered
//		invalid and trigger the return of an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are included in
//		the 'errpref' software package,
//		"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	error
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setNStrNumberFieldSpec(
	signedNumFmtSpec *NumStrFormatSpec,
	decSeparatorChars []rune,
	intGroupingChars []rune,
	intGroupingType IntegerGroupingType,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	numFieldLength int,
	numFieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setNStrNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmtSpec' is invalid!\n"+
			"'signedNumFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmtSpecAtom := numStrFmtSpecAtom{}

	signedNumFmtSpecAtom.empty(
		signedNumFmtSpec)

	err = signedNumFmtSpecAtom.setDecimalSeparatorParams(
		signedNumFmtSpec,
		decSeparatorChars,
		ePrefix.XCpy(
			"signedNumFmtSpec<-"+
				"decSeparatorChars"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setIntegerGroupingParams(
		signedNumFmtSpec,
		intGroupingChars,
		intGroupingType,
		ePrefix.XCpy(
			"signedNumFmtSpec<-"+
				"intGroupingParams"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setRoundingParams(
		signedNumFmtSpec,
		roundingType,
		roundToFractionalDigits,
		ePrefix.XCpy(
			"signedNumFmtSpec<-"+
				"RoundingParams"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setPositiveNumberSign(
		signedNumFmtSpec,
		leadingPosNumSign,
		positiveNumFieldSymPosition,
		trailingPosNumSign,
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"signedNumFmtSpec<-"+
				"Positive Number Sign Params"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setNegativeNumberSign(
		signedNumFmtSpec,
		leadingNegNumSign,
		negativeNumFieldSymPosition,
		trailingNegNumSign,
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"signedNumFmtSpec<-"+
				"Negative Number Sign Params"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setNumberFieldParams(
		signedNumFmtSpec,
		numFieldLength,
		numFieldJustification,
		ePrefix.XCpy(
			"signedNumFmtSpec<-"+
				"Number Field Parameters"))

	return err
}

// numStrFmtSpecAtom - This type provides
// helper methods for NumStrFormatSpec
type numStrFmtSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NumStrFormatSpec and proceeds to reset the
// data values for all member variables to their
// initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'SignedNumFmtSpec' will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	SignedNumFmtSpec           *NumStrFormatSpec
//	   - A pointer to an instance of NumStrFormatSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) empty(
	signedNumFmtSpec *NumStrFormatSpec) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	if signedNumFmtSpec == nil {
		return
	}

	signedNumFmtSpec.decSeparator.Empty()

	signedNumFmtSpec.intGroupingSpec.Empty()

	signedNumFmtSpec.roundingSpec.Empty()

	signedNumFmtSpec.positiveNumberSign.Empty()

	signedNumFmtSpec.negativeNumberSign.Empty()

	signedNumFmtSpec.numberFieldSpec.Empty()

}

// equal - Receives a pointer to two instances of
// NumStrFormatSpec and proceeds to compare their
// member variables in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are
// equal in all respects, this flag is set to 'true'. Otherwise,
// this method returns 'false'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmtSpec1    *NumStrFormatSpec
//	   - An instance of NumStrFormatSpec.
//	     Internal member variables from 'signedNumFmtSpec1'
//	     will be compared to those of 'signedNumFmtSpec2' to
//	     determine if both instances are equivalent.
//
//
//	signedNumFmtSpec2    *NumStrFormatSpec
//	   - An instance of NumStrFormatSpec.
//	     Internal member variables from 'signedNumFmtSpec2'
//	     will be compared to those of 'signedNumFmtSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'signedNumFmtSpec1' and
//	     'signedNumFmtSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) equal(
	signedNumFmtSpec1 *NumStrFormatSpec,
	signedNumFmtSpec2 *NumStrFormatSpec) bool {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	if signedNumFmtSpec1 == nil ||
		signedNumFmtSpec2 == nil {
		return false
	}

	if !signedNumFmtSpec1.decSeparator.Equal(
		&signedNumFmtSpec2.decSeparator) {

		return false
	}

	if !signedNumFmtSpec1.intGroupingSpec.Equal(
		&signedNumFmtSpec2.intGroupingSpec) {

		return false
	}

	if !signedNumFmtSpec1.roundingSpec.Equal(
		&signedNumFmtSpec2.roundingSpec) {

		return false
	}

	if !signedNumFmtSpec1.positiveNumberSign.Equal(
		&signedNumFmtSpec2.positiveNumberSign) {

		return false
	}

	if !signedNumFmtSpec1.negativeNumberSign.Equal(
		&signedNumFmtSpec2.negativeNumberSign) {

		return false
	}

	if !signedNumFmtSpec1.numberFieldSpec.Equal(
		&signedNumFmtSpec2.numberFieldSpec) {

		return false
	}

	return true
}

// setDecimalSeparatorParams - Deletes and resets the member
// variable data value for 'NumStrFormatSpec.decSeparator'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//		A pointer to an instance of NumStrFormatSpec.
//		This instance contains the member variable
//		'signedNumFmt.decSeparator' which will be reset
//		to the value of the decimal separator specified
//		by input parameter, 'decSeparator'.
//
//
//	decSeparator				[]rune
//		This rune array contains the decimal separator
//		character or characters used to separate integer
//		digits from fractional digits in floating point
//		number strings.
//
//		In the United States, the decimal separator is
//		referred to as the decimal point.
//
//	errPrefDto					*ePref.ErrPrefixDto
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setDecimalSeparatorParams(
	signedNumFmt *NumStrFormatSpec,
	decSeparator []rune,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setDecimalSeparatorParams()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.decSeparator.Empty()

	err = signedNumFmt.decSeparator.SetDecimalSeparatorRunes(
		decSeparator,
		ePrefix.XCpy(
			"signedNumFmt.decSeparator<-"+
				"decSeparator"))

	return err
}

// setDecimalSeparatorSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.decSeparator'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// This method receives an instance of 'DecimalSeparatorSpec' and
// copies the member variable data values to
// 'signedNumFmt.decSeparator'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.decSeparator'
//		will be reset to the values provided by the
//		following input parameters.
//
//
//	decSeparatorSpec			DecimalSeparatorSpec
//
//		An instance of DecimalSeparatorSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.decSeparator'.
//
//		In the United States, the decimal separator is
//		referred to as the decimal point.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setDecimalSeparatorSpec(
	signedNumFmt *NumStrFormatSpec,
	decSeparatorSpec DecimalSeparatorSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setDecimalSeparatorSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.decSeparator.Empty()

	err = signedNumFmt.decSeparator.CopyIn(
		&decSeparatorSpec,
		ePrefix.XCpy(
			"signedNumFmt.decSeparatorSpec<-"+
				"decSeparatorSpec"))

	return err
}

// setIntegerGroupingParams - Deletes and resets the member
// variable data value for 'NumStrFormatSpec.intGroupingSpec'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.intGroupingSpec'
//		will be reset to the values provided by the
//		following input parameters.
//
//	intGroupingChars			[]rune
//		One or more characters used to separate groups of
//		integer digits. This separator is also known as the
//		'thousands' separator. It is used to separate groups of
//		integer digits to the left of the decimal separator
//		(a.k.a. decimal point). In the United States, the
//		standard integer digits separator is the comma (",").
//		United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.') is used
//		as the integer separator character.
//		European Example: 1.000.000.000
//
//		Other countries and cultures use spaces, apostrophes or
//		multiple characters to separate integers.
//
//		If this input parameter contains a zero length array
//		and 'intGroupingSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	errPrefDto					*ePref.ErrPrefixDto
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setIntegerGroupingParams(
	signedNumFmt *NumStrFormatSpec,
	intGroupingChars []rune,
	intGroupingType IntegerGroupingType,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setIntegerGroupingParams()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.intGroupingSpec.Empty()

	err = signedNumFmt.intGroupingSpec.SetRunes(
		intGroupingType,
		intGroupingChars,
		ePrefix.XCpy(
			"signedNumFmt.intGroupingSpec<-"))

	return err
}

// setIntegerGroupingSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.intGroupingSpec'
// contained in the instance of NumStrFormatSpec
// passed as an input parameter.
//
// This method receives an instance of
// 'NumStrIntegerGroupingSpec' and copies the member variable
// data values to 'signedNumFmt.intGroupingSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		All the member variable data values in this instance
//		will be deleted and reset to the values contained
//		in the Integer Grouping Specification supplied by
//		input parameter, 'intGroupingSpec'.
//
//
//	intGroupingSpec				NumStrIntegerGroupingSpec
//
//		An instance of NumStrIntegerGroupingSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.intGroupingSpec'.
//
//		In the United States, the Integer Group Specification
//		character is a comma (',') and integer grouped in
//		thousands.
//			United States Example: 1,000,000,000
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setIntegerGroupingSpec(
	signedNumFmt *NumStrFormatSpec,
	intGroupingSpec NumStrIntegerGroupingSpec,
	errPrefDto *ePref.ErrPrefixDto) error {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setIntegerGroupingSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.intGroupingSpec.Empty()

	return signedNumFmt.intGroupingSpec.CopyIn(
		&intGroupingSpec,
		ePrefix.XCpy(
			"signedNumFmt.intGroupingSpec<-"+
				"intGroupingSpec"))
}

// setNegativeNumberSign - Deletes and resets the member variable
// data value for 'NumStrFormatSpec.negativeNumberSign'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt					*NumStrFormatSpec
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.roundingSpec'
//		will be reset to the values provided by the
//		following input parameters.
//
//	leadingNegNumSign				[]rune
//		An array of runes containing the character or
//		characters which will be formatted and displayed
//		in front of a negative numeric value in a number
//		string.
//
//			Example: Leading Number Symbols
//			Leading Symbols: "-"
//			Number String:   "-123.456"
//
//	leadingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Leading Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: leading minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " -123.45"
//					Number Field Index:  01234567
//					Total Number String Length: 8
//
//				In this case the final length of the number string
//				is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-2:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: leading minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     	Formatted Number String: "-  123.45"
//					Number Field Index:  012345678
//					Total Number String Length: 9
//
//				Example-3:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:  0123456789
//					Total Number String Length: 10
//
//				In this case the final length of the number string
//				is greater than the Number Field length.
//
//	trailingNegNumSign				[]rune
//		An array of runes containing the character or
//		characters which will be formatted and displayed
//		after a negative numeric value in a number
//		string.
//
//			Example: Trailing Number Symbols
//				Trailing Symbols: "-"
//				Number String:   "123.456-"
//
//	trailingNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Trailing Number
//		Symbol relative to a Number Field in which
//		a number string is displayed. Possible valid
//		values are listed as follows:
//
//			NumFieldSymPos.InsideNumField()
//				Example-1:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: trailing minus sign ('-')
//					Number Symbol Position: Inside Number Field
//					Formatted Number String: " 123.45-"
//			     Number Text Justification: Right
//					Number Field Index:       01234567
//					Total Number String Length: 8
//
//				In this case the final length of the number string
//				is defined by the Number Field length.
//
//			NumFieldSymPos.OutsideNumField()
//				Example-2:
//					Number Field Length: 8
//			     	Numeric Value: 123.45
//			     	Number Symbol: trailing minus sign ('-')
//			     	Number Symbol Position: Outside Number Field
//			     Number Text Justification: Right
//			     	Formatted Number String: "  123.45-"
//					Number Field Index:       012345678
//					Total Number String Length: 9
//
//				Example-3:
//					Number Field Length: 8
//					Numeric Value: 123.45
//					Number Symbol: before and after parentheses  ('()')
//					Number Symbol Position: Outside Number Field
//			     Number Text Justification: Centered
//					Formatted Number String: "( 123.45 )"
//					Number Field Index:       0123456789
//					Total Number String Length: 10
//
//				In this case the final length of the number string
//				is greater than the Number Field length.
//
//	errPrefDto						*ePref.ErrPrefixDto
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err								error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setNegativeNumberSign(
	signedNumFmt *NumStrFormatSpec,
	leadingNegNumSign []rune,
	leadingNumFieldSymPosition NumberFieldSymbolPosition,
	trailingNegNumSign []rune,
	trailingNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setNegativeNumberSign()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenLeadingNegNumSign := len(leadingNegNumSign)

	lenTrailingNegNumSign := len(trailingNegNumSign)

	if lenLeadingNegNumSign == 0 &&
		lenTrailingNegNumSign == 0 {

		signedNumFmt.negativeNumberSign.SetNOP()

		return err
	}

	signedNumFmt.negativeNumberSign.Empty()

	if lenLeadingNegNumSign > 0 {

		err = signedNumFmt.negativeNumberSign.
			SetLeadingNumberSymbolRunes(
				leadingNegNumSign,
				leadingNumFieldSymPosition,
				ePrefix.XCpy(
					"signedNumFmt.negativeNumberSign"+
						"<-leadingNegNumSign"))

		if err != nil {
			return err
		}
	}

	if lenTrailingNegNumSign > 0 {

		err = signedNumFmt.negativeNumberSign.
			SetTrailingNumberSymbolRunes(
				trailingNegNumSign,
				trailingNumFieldSymPosition,
				ePrefix.XCpy(
					"signedNumFmt.negativeNumberSign<-"+
						"trailingNegNumSign"))

		if err != nil {
			return err
		}
	}

	return err
}

// setNegativeNumberSignSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.negativeNumberSign'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// This method receives an instance of 'NumStrNumberSymbolSpec' and
// copies the member variable data values to
// 'signedNumFmt.negativeNumberSign'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.decSeparator'
//		will be reset to the values provided by the
//		following input parameters.
//
//
//	negativeNumberSign 			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.negativeNumberSign'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setNegativeNumberSignSpec(
	signedNumFmt *NumStrFormatSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setNegativeNumberSignSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if negativeNumberSign.IsNOP() {

		signedNumFmt.negativeNumberSign.SetNOP()

		return err
	}

	err = signedNumFmt.negativeNumberSign.CopyIn(
		&negativeNumberSign,
		ePrefix.XCpy(
			"signedNumFmt.negativeNumberSign<-"+
				"negativeNumberSign"))

	return err
}

// setNumberFieldParams - Deletes and resets the member variable data
// value for 'NumStrFormatSpec.numberFieldSpec'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.numberFieldSpec'
//		will be reset to the values provided by the
//		following input parameters.
//
//	fieldLength					int
//		This parameter defines the length of the text field in
//		which the numeric value will be displayed within a
//		number string.
//
//		If 'fieldLength' is less than the length of the numeric
//		value string, it will be automatically set equal to the
//		length of that numeric value string.
//
//		To automatically set the value of fieldLength to the string
//		length of the numeric value, set this parameter to a value
//		of minus one (-1).
//
//		If this parameter is submitted with a value less than minus
//		one (-1) or greater than 1-million (1,000,000), an error will
//		be returned.
//
//	fieldJustification			TextJustify
//		An enumeration which specifies the justification of the
//		numeric value string within the number field length specified
//		by input parameter 'fieldLength'.
//
//		Text justification can only be evaluated in the context of
//		a number string, field length and a 'textJustification'
//		object of type TextJustify. This is because number strings
//		with a field length equal to or less than the length of the
//		numeric value string never use text justification. In these
//		cases, text justification is completely ignored.
//
//		If the field length parameter ('fieldLength') is greater
//		than the length of the numeric value string, text
//		justification must be equal to one of these
//		three valid values:
//		          TextJustify(0).Left()
//		          TextJustify(0).Right()
//		          TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//		          TxtJustify.Left()
//		          TxtJustify.Right()
//		          TxtJustify.Center()
//
//	errPrefDto					*ePref.ErrPrefixDto
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setNumberFieldParams(
	signedNumFmt *NumStrFormatSpec,
	fieldLength int,
	fieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setNumberFieldParams()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.numberFieldSpec.Empty()

	err = signedNumFmt.numberFieldSpec.SetFieldSpec(
		fieldLength,
		fieldJustification,
		ePrefix.XCpy(
			"signedNumFmt.numberFieldSpec<-"))

	return err
}

// setNumberFieldSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.numberFieldSpec'
// contained in the instance of NumStrFormatSpec
// passed as an input parameter.
//
// This method receives an instance of
// 'NumStrNumberFieldSpec' and copies the member variable
// data values to 'signedNumFmt.numberFieldSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		All the member variable data values in this instance
//		will be deleted and reset to the values contained
//		in the Number Field Specification supplied by
//		input parameter, 'numberFieldSpec'.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		An instance of NumStrNumberFieldSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.numberFieldSpec'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setNumberFieldSpec(
	signedNumFmt *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.numberFieldSpec.Empty()

	return signedNumFmt.numberFieldSpec.CopyIn(
		&numberFieldSpec,
		ePrefix.XCpy(
			"signedNumFmt<-"+
				"numberFieldSpec"))
}

// setPositiveNumberSign - Deletes and resets the member variable
// data value for 'NumStrFormatSpec.positiveNumberSign'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.roundingSpec'
//		will be reset to the values provided by the
//		following input parameters.
//
//	leadingPosNumSymbols			[]rune
//
//		An array of runes containing the character or
//		characters which will be formatted and displayed
//		in front of a positive numeric value in a number
//		string.
//
//	trailingPosNumSymbols			[]rune
//
//		An array of runes containing the character or
//		characters which will be formatted and displayed
//		after a positive numeric value in a number
//		string.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setPositiveNumberSign(
	signedNumFmt *NumStrFormatSpec,
	leadingPosNumSign []rune,
	leadingPosNumFieldSymPosition NumberFieldSymbolPosition,
	trailingPosNumSign []rune,
	trailingPosNumFieldSymPosition NumberFieldSymbolPosition,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setPositiveNumberSign()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	lenLeadingPosNumSign := len(leadingPosNumSign)
	lenTrailingPosNumSign := len(trailingPosNumSign)

	if lenLeadingPosNumSign == 0 &&
		lenTrailingPosNumSign == 0 {

		signedNumFmt.positiveNumberSign.SetNOP()

		return err
	}

	signedNumFmt.positiveNumberSign.Empty()

	if lenLeadingPosNumSign > 0 {

		err = signedNumFmt.positiveNumberSign.
			SetLeadingNumberSymbolRunes(
				leadingPosNumSign,
				leadingPosNumFieldSymPosition,
				ePrefix.XCpy(
					"signedNumFmt.positiveNumberSign"+
						"<-leadingPosNumSymbols"))

		if err != nil {
			return err
		}
	}

	if lenTrailingPosNumSign > 0 {

		err = signedNumFmt.positiveNumberSign.
			SetTrailingNumberSymbolRunes(
				trailingPosNumSign,
				trailingPosNumFieldSymPosition,
				ePrefix.XCpy(
					"signedNumFmt.positiveNumberSign<-"+
						"trailingPosNumSymbols"))

		if err != nil {
			return err
		}
	}

	return err
}

// setPositiveNumberSignSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.positiveNumberSign'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// This method receives an instance of 'NumStrNumberSymbolSpec' and
// copies the member variable data values to
// 'signedNumFmt.positiveNumberSign'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.decSeparator'
//		will be reset to the values provided by the
//		following input parameters.
//
//
//	positiveNumberSign 			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.positiveNumberSign'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setPositiveNumberSignSpec(
	signedNumFmt *NumStrFormatSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setPositiveNumberSignSpec()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if positiveNumberSign.IsNOP() {

		signedNumFmt.positiveNumberSign.SetNOP()

		return err
	}

	err = signedNumFmt.positiveNumberSign.CopyIn(
		&positiveNumberSign,
		ePrefix.XCpy(
			"signedNumFmt.positiveNumberSign<-"+
				"positiveNumberSign"))

	return err
}

// setRoundingParams - Deletes and resets the member variable data
// value for 'NumStrFormatSpec.intGroupingSpec'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.roundingSpec'
//		will be reset to the values provided by the
//		following input parameters.
//
//	roundingType				NumberRoundingType
//
//		This enumeration parameter is used to specify the type
//		of rounding algorithm that will be applied for the
//		rounding of fractional digits in a number string.
//
//		Possible values are listed as follows:
//
//			NumRoundType.NoRounding()
//			NumRoundType.HalfUpWithNegNums()
//			NumRoundType.HalfDownWithNegNums()
//			NumRoundType.HalfAwayFromZero()
//			NumRoundType.HalfTowardsZero()
//			NumRoundType.HalfToEven()
//			NumRoundType.HalfToOdd()
//			NumRoundType.Randomly()
//			NumRoundType.Floor()
//			NumRoundType.Ceiling()
//			NumRoundType.Truncate()
//
//	roundToFractionalDigits		int
//
//		When set to a positive integer value, this parameter
//		controls the number of digits to the right of the decimal
//	 	separator (a.k.a. decimal point) which will remain after
//		completion of the number rounding operation.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setRoundingParams(
	signedNumFmt *NumStrFormatSpec,
	roundingType NumberRoundingType,
	roundToFractionalDigits int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setRoundingParams()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.roundingSpec.Empty()

	err = signedNumFmt.roundingSpec.SetRoundingSpec(
		roundingType,
		roundToFractionalDigits,
		ePrefix.XCpy(
			"signedNumFmt.roundingSpec<-"))

	return err

}

// setRoundingSpec - Deletes and resets the member variable
// data value for 'signedNumFmt.roundingSpec'
// contained in the instance of NumStrFormatSpec passed
// as an input parameter.
//
// This method receives an instance of 'NumStrRoundingSpec'
// and copies the member variable data values to
// 'signedNumFmt.roundingSpec'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.roundingSpec'
//		will be reset to the values provided by the
//		following input parameters.
//
//	roundingSpec				NumStrRoundingSpec
//		An instance of NumStrRoundingSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.roundingSpec'.
//
//	errPrefDto					*ePref.ErrPrefixDto
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will encapsulate
//		an error message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setRoundingSpec(
	signedNumFmt *NumStrFormatSpec,
	roundingSpec NumStrRoundingSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if signedNumFmtSpecAtom.lock == nil {
		signedNumFmtSpecAtom.lock = new(sync.Mutex)
	}

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setRoundingParams()",
		"")

	if err != nil {
		return err
	}

	if signedNumFmt == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'signedNumFmt' is invalid!\n"+
			"'signedNumFmt' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmt.roundingSpec.Empty()

	err = signedNumFmt.roundingSpec.CopyIn(
		&roundingSpec,
		ePrefix.XCpy(
			"signedNumFmt.roundingSpec<-"+
				"roundingSpec"))

	return err
}

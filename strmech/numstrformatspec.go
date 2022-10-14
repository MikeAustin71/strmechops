package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrFormatSpec
//
// The Number String Format Specification contains
// parameters used to format signed numbers and
// currency numeric values as number strings.
type NumStrFormatSpec struct {
	decSeparator DecimalSeparatorSpec
	// Contains the decimal separator character or
	// characters which will separate integer and
	// fractional digits in a floating point number.

	intSeparatorSpec IntegerSeparatorSpec
	// Integer Separator Specification. This parameter
	// specifies the type of integer grouping and
	// integer separator characters which will be
	// applied to the number string formatting
	// operations.

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
	//
	//		type NumStrNumberFieldSpec struct {
	//
	//			fieldLength int
	//
	//				This parameter defines the length of the
	//				text field in which the numeric value will
	//				be displayed within a number string.
	//
	//				If 'fieldLength' is less than the length
	//				of the numeric value string, it will be
	//				automatically set equal to the length of
	//				that numeric value string.
	//
	//				To automatically set the value of
	//				'fieldLength' to the string length of the
	//				numeric value, set this parameter to a
	//				value of minus one (-1).
	//
	//				If this parameter is submitted with a
	//				value less than minus one (-1) or greater
	//				than 1-million (1,000,000), an error will
	//				be returned.
	//
	//			fieldJustification TextJustify
	//
	//				An enumeration which specifies the
	//				justification of the numeric value string
	//				within the number field length specified
	//				by data field 'fieldLength'.
	//
	//				Text justification can only be evaluated in
	//				the context of a number string, field length
	//				and a 'textJustification' object of type
	//				TextJustify. This is because number strings
	//				with a field length equal to or less than the
	//				length of the numeric value string never use
	//				text justification. In these cases, text
	//				justification is completely ignored.
	//
	//				If the field length parameter ('fieldLength')
	//				is greater than the length of the numeric
	//				value string, text justification must be equal
	//				to one of these three valid values:
	//
	//				          TextJustify(0).Left()
	//				          TextJustify(0).Right()
	//				          TextJustify(0).Center()
	//
	//				You can also use the abbreviated text
	//				justification enumeration syntax as follows:
	//
	//				          TxtJustify.Left()
	//				          TxtJustify.Right()
	//				          TxtJustify.Center()
	//		}

	positiveNumberSign NumStrNumberSymbolSpec
	// Positive number signs are commonly implied and
	// not specified. However, the user as the option
	// to specify a positive number sign character or
	// characters for positive numeric values using a
	// Number String Positive Number Sign
	// Specification.

	zeroNumberSign NumStrNumberSymbolSpec
	// The Number String Zero Number Symbol
	// Specification is used to configure number
	// symbols for zero numeric values formatted
	//and displayed in number stings.

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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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

// GetIntSeparatorSpec - Returns a deep copy of the Integer
// Grouping Specification configured for the current instance
// of NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	IntegerSeparatorSpec
//
//		If this method completes successfully, a deep copy of
//		the Integer Separator Specification configured for the
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
func (numStrFmtSpec *NumStrFormatSpec) GetIntSeparatorSpec(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
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
			"GetIntSeparatorSpec()",
		"")

	if err != nil {
		return IntegerSeparatorSpec{}, err
	}

	return numStrFmtSpec.intSeparatorSpec.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.intSeparatorSpec"))
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

	str,
		_ := numStrFmtSpec.intSeparatorSpec.GetIntSeparatorStr(
		nil)

	return str
}

// GetIntegerSeparatorSpec - Returns an instance of
// IntegerSeparatorSpec based on the configuration parameters
// contained within the current instance of
// NumStrFormatSpec.
//
// IntegerSeparatorSpec is used by low level number string
// formatting functions to complete the number string generation
// operation.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	IntegerSeparatorSpec
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
func (numStrFmtSpec *NumStrFormatSpec) GetIntegerSeparatorSpec(
	errorPrefix interface{}) (
	IntegerSeparatorSpec,
	error) {

	if numStrFmtSpec.lock == nil {
		numStrFmtSpec.lock = new(sync.Mutex)
	}

	numStrFmtSpec.lock.Lock()

	defer numStrFmtSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	var newIntSeparatorSpec IntegerSeparatorSpec

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrFormatSpec."+
			"GetIntegerSeparatorSpec()",
		"")

	if err != nil {
		return newIntSeparatorSpec, err
	}

	newIntSeparatorSpec,
		err =
		numStrFmtSpec.intSeparatorSpec.CopyOut(
			ePrefix.XCpy(
				"<-numStrFmtSpec.intSeparatorSpec"))

	return newIntSeparatorSpec, err
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

	intSepRunes,
		_ := numStrFmtSpec.intSeparatorSpec.GetIntSeparatorChars(
		nil)

	return intSepRunes
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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

// GetZeroNumSymSpec - Returns the Zero Number Symbol
// Specification currently configured for this instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NumStrNumberSymbolSpec
//
//		If this method completes successfully, a copy
//		of the zero number sign specification configured
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
func (numStrFmtSpec *NumStrFormatSpec) GetZeroNumSymSpec(
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
			"GetZeroNumSymSpec()",
		"")

	if err != nil {
		return NumStrNumberSymbolSpec{}, err
	}

	return numStrFmtSpec.zeroNumberSign.CopyOut(
		ePrefix.XCpy(
			"<-numStrFmtSpec.zeroNumberSign"))
}

//	NewCurrencyNumFmtFrance
//
//	Returns a new instance of NumStrFormatSpec
//	configured with Currency Number String
//	formatting conventions typically applied
//	in France.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Currency Number String formatting
//	standards implemented by either France or
//	Germany.
//
//	For information on German Number String
//	formatting conventions, see method:
//
//		NumStrFormatSpec.NewCurrencyNumFmtGermany()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewNumFmtParams()
//		NumStrFormatSpec.NewNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The currency symbol used in the France is the
//	Euro symbol ('€').
//
//		France Example-3
//		1 000 000,00 €
//
//	The negative number sign is set to leading minus
//	sign ('-') and a trailing Euro symbol ("€").
//
//		France Example-4
//		-1 000 000,00 €
//
//	The positive number sign is set to a trailing
//	Euro symbol.
//
//		France Example-5
//		1 000 000,00 €
//
//	The zero number format is set to a trailing
//	Euro symbol.
//
//		France Example-6
//			0,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec					NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                	interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newFrenchCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured for
//		Currency Number String formatting parameters
//		typically applied in France.
//
//	err								error
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyNumFmtFrance(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newFrenchCurrencyNumFmtSpec NumStrFormatSpec,
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
			"NewCurrencyNumFmtFrance()",
		"")

	if err != nil {
		return newFrenchCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyNStrFmtFrance(
		&newFrenchCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newFrenchCurrencyNumFmtSpec<-"))

	return newFrenchCurrencyNumFmtSpec, err
}

//	NewCurrencyNumFmtGermany
//
//	Returns a new instance of NumStrFormatSpec
//	configured with Currency Number String
//	formatting conventions typically applied
//	in Germany.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Currency Number String formatting
//	standards implemented by either Germany or
//	France.
//
//	For information on French Number String
//	formatting conventions, see method:
//
//		NumStrFormatSpec.NewCurrencyNumFmtFrance()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewNumFmtParams()
//		NumStrFormatSpec.NewNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://freeformatter.com/germany-standards-code-snippets.html
//
// https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//		1.000.000,00
//
//	The currency symbol used in the Germany is the
//	Euro symbol ('€').
//
//		German Example-3
//		1.000.000,00 €
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-4
//		1.000.000,00- €
//
//	The positive number sign is set to a trailing
//	Euro symbol.
//
//		Germany Example-5
//		1.000.000,00 €
//
//	The zero number format is set to a trailing
//	Euro symbol.
//
//		Germany Example-6
//			0,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec					NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                	interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newGermanCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number String formatting parameters
//		typically applied in Germany.
//
//	err								error
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyNumFmtGermany(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newGermanCurrencyNumFmtSpec NumStrFormatSpec,
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
			"NewCurrencyNumFmtGermany()",
		"")

	if err != nil {
		return newGermanCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyNStrFmtGermany(
		&newGermanCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newGermanCurrencyNumFmtSpec<-"))

	return newGermanCurrencyNumFmtSpec, err
}

//	NewCurrencyNumFmtUK
//
//	Returns a new instance of NumStrFormatSpec
//	configured for UK (United Kingdom) Currency
//	Number String formatting conventions.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewNumFmtParams()
//		NumStrFormatSpec.NewNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//		123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United Kingdom Example-2
//				1,000,000
//
//	The currency symbol used in the United Kingdom is the
//	Pound Sterling symbol ('£').
//
//		United Kingdom Example-3
//			£ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United Kingdom Example-4
//			£ -1,000,000.00
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			£ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-6
//				£ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newUKCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number String formatting parameters
//		typically applied in the United Kingdom.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyNumFmtUK(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newUKCurrencyNumFmtSpec NumStrFormatSpec,
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
			"NewCurrencyNumFmtUK()",
		"")

	if err != nil {
		return newUKCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyNStrFmtUK(
		&newUKCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newUKCurrencyNumFmtSpec<-"))

	return newUKCurrencyNumFmtSpec, err
}

//	NewCurrencyNumFmtUS
//
//	Returns a new instance of NumStrFormatSpec
//	configured for US (United States) Currency
//	Number String formatting conventions.
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewNumFmtParams()
//		NumStrFormatSpec.NewNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United States Example-2
//			1,000,000
//
//	The currency symbol used in the United States is the
//	Dollar Sign symbol ('$').
//
//		United States Example-3
//			$ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United States Example-4
//			$ (1,000,000.00)
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			$ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-6
//			$ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newUSCurrencyNumFmtSpec		NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of NumStrFormatSpec configured with
//		Currency Number String formatting parameters
//		typically applied in the United States.
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
func (numStrFmtSpec *NumStrFormatSpec) NewCurrencyNumFmtUS(
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (
	newUSCurrencyNumFmtSpec NumStrFormatSpec,
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
			"NewCurrencyNumFmtUS()",
		"")

	if err != nil {
		return newUSCurrencyNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setCurrencyNStrFmtUS(
		&newUSCurrencyNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newUSCurrencyNumFmtSpec<-"))

	return newUSCurrencyNumFmtSpec, err
}

//	NewNumFmtComponents
//
//	Creates and returns a new instance of NumStrFormatSpec
//	generated from Number String formatting input
//	components passed as input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator				DecimalSeparatorSpec
//
//		This structure contains the radix point or decimal
//		separator character(s) (a.k.a. decimal point)
//		which be used to separate integer and fractional
//		digits within a formatted Number String.
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Integer Separator Specification. This type
//		encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//	negativeNumberSign			IntegerSeparatorSpec
//
//		This Integer Separator Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with negative numeric values.
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with positive numeric values.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with zero numeric values.
//
//	numberFieldSpec			NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this
//		parameter will return a new, fully populated
//		instance of	NumStrFormatSpec.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the	method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (numStrFmtSpec *NumStrFormatSpec) NewNumFmtComponents(
	decSeparator DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
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
			"NewNumFmtParams()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecAtom).setNStrFmtComponents(
		&newSignedNumFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

// NewNumFmtParams - Creates and returns a new instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					string
//
//		This string contains the character or characters
//		which will be configured as the Decimal Separator
//		Symbol or Symbols for the returned instance of
//		DecimalSeparatorSpec, the Decimal Separator
//		Specification.
//
//	intSeparatorChars				string
//
//		One or more characters used to separate groups of
//		integers. This separator is also known as the
//		'thousands' separator. It is used to separate
//		groups of integer digits to the left of the
//		decimal separator (a.k.a. decimal point). In the
//		United States, the standard integer digits
//		separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		string and parameter 'intGroupingType' is NOT
//		equal to 'IntGroupingType.None()', an error will
//		be returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//	leadingPosNumSign				string
//
//		A string containing the leading positive number
//		sign character or characters used to configure
//		a Positive Number Sign Symbol in a number string
//		with a positive numeric value.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
//
//	trailingPosNumSign				string
//
//		A string containing the trailing positive number
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
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				string
//
//		A string containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//	trailingNegNumSign				string
//
//		A string containing the trailing negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				string
//
//		A string containing the leading zero number sign
//		character or characters used to configure a	Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
//
//	trailingZeroNumSign				string
//
//		A string containing the trailing zero number sign
//		character or characters used to configure a Number
//		Sign Symbol in a number string with a zero
//		numeric value.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	zeroNumFieldSymPosition			NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	numFieldLength					int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) NewNumFmtParams(
	decSeparatorChars string,
	intSeparatorChars string,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign string,
	trailingPosNumSign string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign string,
	trailingNegNumSign string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign string,
	trailingZeroNumSign string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
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
			"NewNumFmtParams()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setNStrNumberFieldSpec(
			&newSignedNumFmtSpec,
			[]rune(decSeparatorChars),
			[]rune(intSeparatorChars),
			intGroupingType,
			[]rune(leadingPosNumSign),
			[]rune(trailingPosNumSign),
			positiveNumFieldSymPosition,
			[]rune(leadingNegNumSign),
			[]rune(trailingNegNumSign),
			negativeNumFieldSymPosition,
			[]rune(leadingZeroNumSign),
			[]rune(trailingZeroNumSign),
			zeroNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

// NewNumFmtParamsRunes - Creates and returns a new instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator					[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the Decimal
//		Separator Symbol or Symbols for the returned
//		instance of	DecimalSeparatorSpec, the Decimal
//		Separator Specification.
//
//	intSeparatorChars				[]rune
//
//		A rune array containing one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator in the United States. It is used to
//		separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In
//	  	the United States, the standard	integer digits
//	  	separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		rune array and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be
//		returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the
//		type of IntegerSeparatorSpec which will be
//		returned. The enumeration IntegerGroupingType must
//		be set to one of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//	leadingPosNumSign				[]rune
//
//		A rune array containing the leading positive
//		number sign character or characters used to
//		configure a Positive Number Sign Symbol in a
//		number string with a positive numeric value.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
//
//	trailingPosNumSign				[]rune
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
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				[]rune
//
//		A rune array containing the leading zero number
//		sign character or characters used to configure
//		Number Sign Symbols in a number string with a
//		zero numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+') and/or currency symbols ('$').
//
//	trailingZeroNumSign				[]rune
//
//		A rune array containing the trailing zero number
//		sign character or characters used to configure
//		Number Sign Symbols in a number string with a
//		zero numeric value.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	zeroNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) NewNumFmtParamsRunes(
	decSeparatorChars []rune,
	intSeparatorChars []rune,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign []rune,
	trailingZeroNumSign []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
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
			"NewNumFmtParamsRunes()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).
		setNStrNumberFieldSpec(
			&newSignedNumFmtSpec,
			decSeparatorChars,
			intSeparatorChars,
			intGroupingType,
			leadingPosNumSign,
			trailingPosNumSign,
			positiveNumFieldSymPosition,
			leadingNegNumSign,
			trailingNegNumSign,
			negativeNumFieldSymPosition,
			leadingZeroNumSign,
			trailingZeroNumSign,
			zeroNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumFmtFrance
//
//	Returns a new instance of NumStrFormatSpec
//	configured for a Signed Number using French
//	Number String formatting conventions.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on German Signed Number
//	String formatting conventions, see method:
//
//		NumStrFormatSpec.NewSignedNumFmtGermany()
//
//	If custom decimal separator, integer separator
//	and negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewNumFmtParams()
//		NumStrFormatSpec.NewNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		Example-1:
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		French Example-3:
//		-1 000 000 000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//	The zero number format is set to a blank or empty
//	string ("").
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec configured with French number
//		formatting specifications.
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumFmtFrance(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"NewSignedNumFmtFrance()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNStrFmtFrance(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumFmtGermany
//
//	Returns a new instance of NumStrFormatSpec
//	configured for a Signed Number using German
//	Number String formatting conventions.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on French Signed Number
//	String formatting conventions, see method:
//
//		NumStrFormatSpec.NewSignedNumFmtFrance()
//
//	If custom decimal separator, integer separator
//	and negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewNumFmtParams()
//		NumStrFormatSpec.NewNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		Example-1:
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		Example-2:
//		1.000.000.000
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		Example-3:
//		1.000.000-
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//	The zero number format is set to a blank or empty
//	string ("").
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	newSignedNumFmtSpec			NumStrFormatSpec
//
//		If this method completes successfully, this parameter
//		will return a new, fully populated instance of
//		NumStrFormatSpec configured with German number
//		formatting specifications.
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumFmtGermany(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"NewSignedNumFmtGermany()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNStrFmtGermany(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumFmtUK
//
//	Returns a new instance of NumStrFormatSpec
//	configured for a Signed Number using UK
//	(United Kingdom) Number String formatting
//	conventions.
//
//	If custom decimal separator, integer separators
//	or negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewNumFmtParams()
//		NumStrFormatSpec.NewNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://freeformatter.com/united-kingdom-standards-code-snippets.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United Kingdom Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United Kingdom Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumFmtUK(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"NewSignedNumFmtUK()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNStrFmtUS(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	NewSignedNumFmtUS
//
//	Returns a new instance of NumStrFormatSpec
//	configured for a Signed Number using US
//	(United States) Number String formatting
//	conventions.
//
//	If custom decimal separator, integer separators
//	or negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.NewNumFmtComponents()
//		NumStrFormatSpec.NewNumFmtParams()
//		NumStrFormatSpec.NewNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United States Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United States Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) NewSignedNumFmtUS(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"NewSignedNumFmtUS()",
		"")

	if err != nil {
		return newSignedNumFmtSpec, err
	}

	err = new(numStrFmtSpecNanobot).setSignedNStrFmtUS(
		&newSignedNumFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("newSignedNumFmtSpec<-"))

	return newSignedNumFmtSpec, err
}

//	SetCurrencyFmtFrance
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using Currency Number
//	String formatting conventions typically
//	applied in France.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Currency Number String formatting
//	standards implemented by either France or
//	Germany.
//
//	For information on German Number String
//	formatting conventions, see method:
//
//		NumStrFormatSpec.SetCurrencyFmtGermany()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetNumFmtParams()
//		NumStrFormatSpec.SetNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied in the France.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://www.ibm.com/support/pages/english-and-french-currency-formats
//
// https://freeformatter.com/france-standards-code-snippets.html
//
// https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The currency symbol used in the France is the
//	Euro symbol ('€').
//
//		France Example-3
//		1 000 000,00 €
//
//	The negative number sign is set to leading minus
//	sign ('-') and a trailing Euro symbol ("€").
//
//		France Example-4
//		-1 000 000,00 €
//
//	The positive number sign is set to a trailing
//	Euro symbol.
//
//		France Example-5
//		1 000 000,00 €
//
//	The zero number format is set to a trailing
//	Euro symbol.
//
//		France Example-6
//			0,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyFmtFrance(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetCurrencyFmtFrance()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyNStrFmtFrance(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyFmtGermany
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using Currency Number
//	String formatting conventions typically
//	applied in Germany.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Currency Number String formatting
//	standards implemented by either Germany or
//	France.
//
//	For information on French Number String
//	formatting conventions, see method:
//
//		NumStrFormatSpec.SetCurrencyFmtFrance()
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetNumFmtParams()
//		NumStrFormatSpec.SetNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied in the Germany.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://freeformatter.com/germany-standards-code-snippets.html
//
// https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//		1.000.000,00
//
//	The currency symbol used in the Germany is the
//	Euro symbol ('€').
//
//		German Example-3
//		1.000.000,00 €
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-4
//		1.000.000,00- €
//
//	The positive number sign is set to a trailing
//	Euro symbol.
//
//		Germany Example-5
//		1.000.000,00 €
//
//	The zero number format is set to a trailing
//	Euro symbol.
//
//		Germany Example-6
//			0,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger text Number Field. In addition
//		to specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyFmtGermany(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetCurrencyFmtGermany()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyNStrFmtGermany(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyFmtUK
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the UK (United Kingdom).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied in the UK (United Kingdom).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//		123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United Kingdom Example-2
//				1,000,000
//
//	The currency symbol used in the United Kingdom is the
//	Pound Sterling symbol ('£').
//
//		United Kingdom Example-3
//			£ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United Kingdom Example-4
//			£ -1,000,000.00
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			£ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-6
//				£ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyFmtUK(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetCurrencyFmtUK()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyNStrFmtUK(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetCurrencyFmtUS
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the US (United States).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied in the US (United States).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United States Example-2
//			1,000,000
//
//	The currency symbol used in the United States is the
//	Dollar Sign symbol ('$').
//
//		United States Example-3
//			$ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United States Example-4
//			$ (1,000,000.00)
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			$ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-6
//			$ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetCurrencyFmtUS(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetCurrencyFmtUS()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).
		setCurrencyNStrFmtUS(
			numStrFmtSpec,
			numberFieldSpec,
			ePrefix.XCpy("numStrFmtSpec<-"))
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		An instance of Integer Separator Specification.
//		The member variable data values contained in
//		this structure will be copied to the current
//		instance of	NumStrFormatSpec :
//
//			'NumStrFormatSpec.intSeparatorSpec'.
//
//		In the United States, the Integer Separator
//		Specification character is a comma (',') with
//		integers grouped in	thousands.
//
//			United States Example: 1,000,000
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
	intSeparatorSpec IntegerSeparatorSpec,
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
		intSeparatorSpec,
		ePrefix.XCpy(
			"numStrFmtSpec<-intSeparatorSpec"))
}

//	SetNegativeNumberFmtSpec
//
//	Deletes and replaces the Negative Number Format
//	Specification for the current instance of
//	NumStrFormatSpec:
//
//		NumStrFormatSpec.negativeNumberSign
//
//	If the current instance of NumStrFormatSpec is set
//	to a negative numeric value, this formatting
//	specification will be applied.
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetNegativeNumberFmtSpec(
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
			"SetNegativeNumberFmtSpec()",
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
//		An instance of NumStrNumberFieldSpec. The member
//		variable data values contained in this instance
//		will be copied to the current instance of
//		NumStrFormatSpec :
//			'NumStrFormatSpec.numberFieldSpec'.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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

//	SetPositiveNumberFmtSpec
//
//	Deletes and replaces the Positive Number Sign Format
//	Specification for the current instance of
//	NumStrFormatSpec:
//
//		NumStrFormatSpec.positiveNumberSign
//
//	If the current instance of NumStrFormatSpec is set
//	to a positive numeric value, this formatting
//	specification will be applied.
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetPositiveNumberFmtSpec(
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
			"SetPositiveNumberFmtSpec()",
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

//	SetNumFmtComponents
//
//	Reconfigures the current instance of NumStrFormatSpec
//	based on Number String formatting components passed as
//	input parameters.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by values generated from the listed input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	decSeparator				DecimalSeparatorSpec
//
//		This structure contains the radix point or decimal
//		separator character(s) (a.k.a. decimal point)
//		which be used to separate integer and fractional
//		digits within a formatted Number String.
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Integer Separator Specification. This type
//		encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with negative numeric values.
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with positive numeric values.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with zero numeric values.
//
//	numberFieldSpec			NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetNumFmtComponents(
	decSeparator DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errorPrefix interface{}) (err error) {

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
			"SetNumFmtComponents()",
		"")

	if err != nil {
		return err
	}

	err = new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))

	return err
}

// SetNumFmtParams - Deletes and resets all the member variable
// data values stored in the current instance of
// NumStrFormatSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	current instance of NumStrFormatSpec will be deleted
//	and replaced by values generated from the listed
//	input parameters.
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
//		and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be returned.
//
//	intGroupingType				IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the type
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetNumFmtParams(
	decSeparatorChars string,
	intGroupingChars string,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign string,
	trailingPosNumSign string,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign string,
	trailingNegNumSign string,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign string,
	trailingZeroNumSign string,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
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
			"SetNumFmtParams()",
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
			[]rune(leadingPosNumSign),
			[]rune(trailingPosNumSign),
			positiveNumFieldSymPosition,
			[]rune(leadingNegNumSign),
			[]rune(trailingNegNumSign),
			negativeNumFieldSymPosition,
			[]rune(leadingZeroNumSign),
			[]rune(trailingZeroNumSign),
			zeroNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return err
}

// SetNumFmtParamsRunes - Deletes and resets all the member variable
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
//		of IntegerSeparatorSpec which will be returned. The
//		enumeration IntegerGroupingType must be set to one
//		of the following values:
//		IntGroupingType.None()
//		IntGroupingType.Thousands()
//		IntGroupingType.IndiaNumbering()
//		IntGroupingType.ChineseNumbering()
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
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetNumFmtParamsRunes(
	decSeparatorChars []rune,
	intGroupingChars []rune,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign []rune,
	trailingZeroNumSign []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
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
			"SetNumFmtParamsRunes()",
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
			leadingPosNumSign,
			trailingPosNumSign,
			positiveNumFieldSymPosition,
			leadingNegNumSign,
			trailingNegNumSign,
			negativeNumFieldSymPosition,
			leadingZeroNumSign,
			trailingZeroNumSign,
			zeroNumFieldSymPosition,
			numFieldLength,
			numFieldJustification,
			ePrefix.XCpy(
				"newSignedNumFmtSpec<-"))

	return err
}

//	SetSignedNumFmtFrance
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using Number String
//	formatting conventions typically
//	applied in France.
//
//	Within in the European Union many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on German Signed Number
//	String formatting conventions, see method:
//
//		NumStrFormatSpec.SetSignedNumFmtGermany()
//
//	If custom decimal separator, integer separator
//	and negative number sign characters are required,
//	see methods:
//
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetNumFmtParams()
//		NumStrFormatSpec.SetNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in France.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		Example-1:
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		Example-2:
//		1 000 000 000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		Example-3:
//		-1 000 000 000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//	The zero number format is set to a blank or empty
//	string ("").
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumFmtFrance(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetSignedNumFmtFrance()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNStrFmtFrance(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetSignedNumFmtGermany
//
//	Reconfigures the current instance of
//	NumStrFormatSpec using Number String
//	formatting conventions typically
//	applied in Germany.
//
//	Within in the European Union, many, if not
//	most, of the member countries subscribe to
//	the Signed Number String formatting standards
//	implemented by either France or Germany.
//
//	For information on French Signed Number
//	String formatting conventions, see method:
//
//		NumStrFormatSpec.SetSignedNumFmtFrance()
//
//	If custom decimal separator, integer separator
//	and negative number sign characters are required,
//	see method:
//
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetNumFmtParams()
//		NumStrFormatSpec.SetNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://freeformatter.com/germany-standards-code-snippets.html
//
//	https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in Germany.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1:
//
//		123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//		1.000.000.000
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		Example-3:
//		1.000.000-
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//	The zero number format is set to a blank or empty
//	string ("").
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumFmtGermany(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetSignedNumFmtGermany()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNStrFmtGermany(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetSignedNumFmtUK
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the UK (United Kingdom).
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetNumFmtParams()
//		NumStrFormatSpec.SetNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
//	https://freeformatter.com/united-kingdom-standards-code-snippets.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in the UK (United Kingdom).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United Kingdom Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United Kingdom Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumFmtUK(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetSignedNumFmtUK()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNStrFmtUS(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetSignedNumFmtUS
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the US (United States).
//
//	If custom decimal separator, integer separators,
//	negative number sign characters or currency
//	symbols are required, see methods:
//
//		NumStrFormatSpec.SetNumFmtComponents()
//		NumStrFormatSpec.SetNumFmtParams()
//		NumStrFormatSpec.SetNumFmtParamsRunes()
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in the US (United States).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United States Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United States Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetSignedNumFmtUS(
	numberFieldSpec NumStrNumberFieldSpec,
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
			"SetSignedNumFmtUS()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecNanobot).setSignedNStrFmtUS(
		numStrFmtSpec,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	SetZeroNumberFmtSpec
//
//	Deletes and replaces the Zero Number Format
//	Specification for the current instance of
//	NumStrFormatSpec:
//
//		NumStrFormatSpec.zeroNumberSign
//
//	If the current instance of NumStrFormatSpec is set
//	to a zero numeric value, this formatting specification
//	will be applied.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to the current
//		NumStrFormatSpec member variable:
//			'NumStrFormatSpec.zeroNumberSign'.
//
//	 errorPrefix                interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.  IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
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
func (numStrFmtSpec *NumStrFormatSpec) SetZeroNumberFmtSpec(
	zeroNumberSign NumStrNumberSymbolSpec,
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
			"SetZeroNumberFmtSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrFmtSpecAtom).
		setZeroNumberSignSpec(
			numStrFmtSpec,
			zeroNumberSign,
			ePrefix.XCpy(
				"numStrFmtSpec<-"+
					"zeroNumberSign"))

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
//
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

	err = new(decimalSepSpecNanobot).
		copyDecimalSeparator(
			&destinationSignedNumFmtSpec.decSeparator,
			&sourceSignedNumFmtSpec.decSeparator,
			ePrefix.XCpy(
				"destinationSignedNumFmtSpec.decSeparator"+
					"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = destinationSignedNumFmtSpec.intSeparatorSpec.CopyIn(
		&sourceSignedNumFmtSpec.intSeparatorSpec,
		ePrefix.XCpy(
			"destinationSignedNumFmtSpec.intSeparatorSpec"+
				"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	var nStrNumSymSpecNanobot numStrNumberSymbolSpecNanobot

	err = nStrNumSymSpecNanobot.
		copyNStrNumberSymbolSpec(
			&destinationSignedNumFmtSpec.positiveNumberSign,
			&sourceSignedNumFmtSpec.positiveNumberSign,
			ePrefix.XCpy(
				"destinationSignedNumFmtSpec.positiveNumberSign"+
					"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = nStrNumSymSpecNanobot.
		copyNStrNumberSymbolSpec(
			&destinationSignedNumFmtSpec.negativeNumberSign,
			&sourceSignedNumFmtSpec.negativeNumberSign,
			ePrefix.XCpy(
				"destinationSignedNumFmtSpec.negativeNumberSign"+
					"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = new(numStrNumberFieldSpecNanobot).
		copyNStrNumberFieldSpec(
			&destinationSignedNumFmtSpec.numberFieldSpec,
			&sourceSignedNumFmtSpec.numberFieldSpec,
			ePrefix.XCpy(
				"destinationSignedNumFmtSpec.numberFieldSpec"+
					"<-sourceSignedNumFmtSpec"))

	if err != nil {
		return err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		copyNStrNumberSymbolSpec(
			&destinationSignedNumFmtSpec.zeroNumberSign,
			&sourceSignedNumFmtSpec.zeroNumberSign,
			ePrefix.XCpy(
				"destinationSignedNumFmtSpec.zeroNumberSign"+
					"<-sourceSignedNumFmtSpec"))

	return err
}

//	setCurrencyNStrFmtFrance
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Reconfigures 'numStrFmtSpec' using Currency Number
//	String formatting conventions typically applied in France.
//
//	Within in the European Union many, if not most, of
//	the member countries subscribe to the Number String
//	Currency formatting standards implemented by either
//	France or Germany.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the instance
//	of NumStrFormatSpec passed as input parameter,
//	'numStrFmtSpec', will be deleted and replaced by Currency
//	Number String formatting parameters typically applied in
//	France.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The currency symbol used in the France is the
//	Euro symbol ('€').
//
//		France Example-3
//		1 000 000,00 €
//
//	The negative number sign is set to leading minus
//	sign ('-') and a trailing Euro symbol ("€").
//
//		France Example-4
//		-1 000 000,00 €
//
//	The positive number sign is set to a trailing
//	Euro symbol.
//
//		France Example-5
//		1 000 000,00 €
//
//	The zero number format is set to a trailing
//	Euro symbol.
//
//		France Example-6
//			0,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameter described below.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtFrance(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtFrance()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewFrance(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewFrenchDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{' ', '\U000020ac'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		lock: nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     nil,
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.None(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{' ', '\U000020ac'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     nil,
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.None(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{' ', '\U000020ac'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),
		lock:                              nil,
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	setCurrencyNStrFmtGermany
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Reconfigures the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec' using Number
//	String Currency formatting conventions typically
//	applied in the Germany.
//
//	Within in the European Union, many member countries
//	subscribe to the Number String Currency formatting
//	standards implemented by either Germany or France.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://freeformatter.com/germany-standards-code-snippets.html
//
// https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	NumStrFormatSpec input	parameter, 'numStrFmtSpec',
//	will be deleted and replaced by Number String Currency
//	formatting parameters typically applied in Germany.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a period character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//		1.000.000,00
//
//	The currency symbol used in the Germany is the
//	Euro symbol ('€').
//
//		German Example-3
//		1.000.000,00 €
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-4
//		1.000.000,00- €
//
//	The positive number sign is set to a trailing
//	Euro symbol.
//
//		Germany Example-5
//		1.000.000,00 €
//
//	The zero number format is set to a trailing
//	Euro symbol.
//
//		Germany Example-6
//			0,00 €
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameter described below.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtGermany(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtGermany()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewGermany(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewGermanDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     nil,
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.None(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'-', ' ', '\U000020ac'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		lock: nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{' ', '\U000020ac'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{' ', '\U000020ac'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),
		lock:                              nil,
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	setCurrencyNStrFmtUK
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the UK (United Kingdom).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied the in UK (United Kingdom).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United Kingdom Example-1
//		123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United Kingdom Example-2
//				1,000,000
//
//	The currency symbol used in the United Kingdom is the
//	Pound Sterling symbol ('£').
//
//		United Kingdom Example-3
//			£ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United Kingdom Example-4
//			£ -1,000,000.00
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-5
//			£ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United Kingdom Example-6
//				£ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameter described below.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text justification
//				enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtUK(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtUK()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewUS(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'\U000000a3', ' ', '-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{},

		trailingNumberFieldSymbolPosition: 0,

		lock: nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'\U000000a3', ' '},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'\U000000a3', ' '},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	setCurrencyNStrFmtUS
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Currency Number String formatting conventions
//	typically applied in the US (United States).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Currency Number String formatting parameters typically
//	applied the in US (United States).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//		123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits
//	each:
//
//		United States Example-2
//				1,000,000
//
//	The currency symbol used in the United States is the
//	Dollar Sign symbol ('$').
//
//		United States Example-3
//			$ 1,000,000.00
//
//	The negative number sign is set to leading and
//	trailing parentheses ("()").
//
//		United States Example-4
//			$ (1,000,000.00)
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			$ 1,000,000.00
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-6
//				$ 0.00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameter described below.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setCurrencyNStrFmtUS(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setCurrencyNStrFmtUS()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewUS(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' ', '('},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{')'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		lock: nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' '},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'$', ' '},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},

		leadingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	setNStrNumberFieldSpec
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'numStrFmtSpec'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in input
//	parameter 'signedNumFmtSpec' will be deleted and
//	replaced by values generated from the listed input
//	parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrFmtSpec							*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All  member
//		variable data fields in this object will be replaced
//		by data values configured from the input parameter
//		described below.
//
//
//	decSeparator					[]rune
//
//		This rune array contains the character or
//		characters which will be configured as the Decimal
//		Separator Symbol or Symbols for the returned
//		instance of	DecimalSeparatorSpec, the Decimal
//		Separator Specification.
//
//	intSeparatorChars				[]rune
//
//		A rune array containing one or more characters
//		used to separate groups of integers. This
//		separator is also known as the 'thousands'
//		separator in the United States. It is used to
//		separate groups of integer digits to the left of
//	  	the decimal separator (a.k.a. decimal point). In
//	  	the United States, the standard	integer digits
//	  	separator is the comma (",").
//
//			United States Example:  1,000,000,000
//
//		In many European countries, a single period ('.')
//		is used as the integer separator character.
//
//			European Example: 1.000.000.000
//
//		Other countries and cultures use spaces,
//		apostrophes or multiple characters to separate
//		integers.
//
//		If this input parameter contains a zero length
//		rune array and 'intSeparatorSpec' is NOT equal to
//		'IntGroupingType.None()', an error will be
//		returned.
//
//	intGroupingType					IntegerGroupingType
//
//		This instance of IntegerGroupingType defines the
//		type of IntegerSeparatorSpec which will be
//		returned. The enumeration IntegerGroupingType must
//		be set to one of the following values:
//
//			IntGroupingType.None()
//			IntGroupingType.Thousands()
//			IntGroupingType.IndiaNumbering()
//			IntGroupingType.ChineseNumbering()
//
//	leadingPosNumSign				[]rune
//
//		A rune array containing the leading positive
//		number sign character or characters used to
//		configure a Positive Number Sign Symbol in a
//		number string with a positive numeric value.
//
//		Leading number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//		currency symbols ('$').
//
//	trailingPosNumSign				[]rune
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
//		character, or characters, relative to a Number
//		Field in which a number string is displayed.
//		Possible valid values are listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the
//				number string is greater than the Number
//				Field length.
//
//	leadingNegNumSign				[]rune
//
//		A rune array containing the leading negative number
//		sign character or characters used to configure
//		Negative Number Sign Symbols in a number string
//		with a negative numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//	trailingNegNumSign				[]rune
//
//		A rune array containing the trailing negative
//		number sign character or characters used to
//		configure Negative Number Sign Symbols in a
//		number string with a negative numeric value.
//
//		Trailing number symbols can include any
//		combination of characters such as minus signs
//		('-') and/or currency symbols ('$').
//
//	negativeNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the Negative Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the number
//				string is defined by the Number Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	leadingZeroNumSign				[]rune
//
//		A rune array containing the leading zero number
//		sign character or characters used to configure
//		Number Sign Symbols in a number string with a
//		zero numeric value.
//
//		Leading number symbols can include any
//		combination of characters such as plus signs
//		('+') and/or currency symbols ('$').
//
//	trailingZeroNumSign				[]rune
//
//		A rune array containing the trailing zero number
//		sign character or characters used to configure
//		Number Sign Symbols in a number string with a
//		zero numeric value.
//
//		Trailing number symbols can include any combination
//		of characters such as plus signs ('+') and/or
//	 	currency symbols ('$').
//
//	zeroNumFieldSymPosition		NumberFieldSymbolPosition
//
//		Defines the position of the zero Number Sign
//		relative to a Number Field in which a number
//		string is displayed. Possible valid values are
//		listed as follows:
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
//				For the 'NumFieldSymPos.InsideNumField()'
//				specification, the final length of the
//				number string is defined by the Number
//				Field length.
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
//				For the 'NumFieldSymPos.OutsideNumField()'
//				specification, the final length of the number
//				string is greater than the Number Field length.
//
//	numFieldLength				int
//
//		This parameter defines the length of the text
//		field in which the numeric value will be displayed
//		within a number string.
//
//		If 'numFieldLength' is less than the length of the
//		numeric value string, it will be automatically set
//		equal to the length of that numeric value string.
//
//		To automatically set the value of fieldLength to
//		the string length of the numeric value, set this
//		parameter to a value of minus one (-1).
//
//		If this parameter is submitted with a value less
//		than minus one (-1) or greater than 1-million
//		(1,000,000), an error will be returned.
//
//	numFieldJustification		TextJustify
//
//		An enumeration which specifies the justification
//		of the numeric value within the number field
//		length specified by input parameter
//		'numFieldLength'.
//
//		Text justification can only be evaluated in the
//		context of a number string, field length and a
//		'textJustification' object of type TextJustify.
//		This is because number strings with a field length
//		equal to or less than the length of the numeric
//		value string never use text justification. In
//		these cases, text justification is completely
//		ignored.
//
//		If the field length parameter ('numFieldLength')
//		is greater than the length of the numeric value
//		string, text justification must be equal to one
//		of these three valid values:
//
//			TextJustify(0).Left()
//			TextJustify(0).Right()
//			TextJustify(0).Center()
//
//		You can also use the abbreviated text justification
//		enumeration syntax as follows:
//
//			TxtJustify.Left()
//			TxtJustify.Right()
//			TxtJustify.Center()
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setNStrNumberFieldSpec(
	numStrFmtSpec *NumStrFormatSpec,
	decSeparatorChars []rune,
	intGroupingChars []rune,
	intGroupingType IntegerGroupingType,
	leadingPosNumSign []rune,
	trailingPosNumSign []rune,
	positiveNumFieldSymPosition NumberFieldSymbolPosition,
	leadingNegNumSign []rune,
	trailingNegNumSign []rune,
	negativeNumFieldSymPosition NumberFieldSymbolPosition,
	leadingZeroNumSign []rune,
	trailingZeroNumSign []rune,
	zeroNumFieldSymPosition NumberFieldSymbolPosition,
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

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	signedNumFmtSpecAtom := numStrFmtSpecAtom{}

	signedNumFmtSpecAtom.empty(
		numStrFmtSpec)

	err = signedNumFmtSpecAtom.setDecimalSeparatorParams(
		numStrFmtSpec,
		decSeparatorChars,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"decSeparatorChars"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setIntegerGroupingParams(
		numStrFmtSpec,
		intGroupingChars,
		intGroupingType,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"intGroupingParams"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setPositiveNumberSign(
		numStrFmtSpec,
		leadingPosNumSign,
		positiveNumFieldSymPosition,
		trailingPosNumSign,
		positiveNumFieldSymPosition,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"Positive Number Sign Params"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setNegativeNumberSign(
		numStrFmtSpec,
		leadingNegNumSign,
		negativeNumFieldSymPosition,
		trailingNegNumSign,
		negativeNumFieldSymPosition,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"Negative Number Sign Params"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setZeroNumberSign(
		numStrFmtSpec,
		leadingZeroNumSign,
		zeroNumFieldSymPosition,
		trailingZeroNumSign,
		zeroNumFieldSymPosition,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"Zero Number Sign Params"))

	if err != nil {
		return err
	}

	err = signedNumFmtSpecAtom.setNumberFieldParams(
		numStrFmtSpec,
		numFieldLength,
		numFieldJustification,
		ePrefix.XCpy(
			"numStrFmtSpec<-"+
				"Number Field Parameters"))

	return err
}

//	setSignedNStrFmtFrance
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'signedNumFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the France.
//
//	Within in the European Union many, if not most, of
//	the member countries subscribe to the Signed Number
//	String formatting standards implemented by either
//	France or Germany.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	NumStrFormatSpec input	parameter, 'numStrFmtSpec', will
//	be deleted and replaced by Number String formatting
//	parameters typically applied in France.
//
// ----------------------------------------------------------------
//
// # Reference:
//
//	https://www.ibm.com/support/pages/english-and-french-currency-formats
//
//	https://freeformatter.com/france-standards-code-snippets.html
//
//	https://docs.microsoft.com/en-us/globalization/locale/currency-formatting
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		French Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a space character
//	(' ').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		French Example-2
//		1 000 000 000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		French Example-3
//		-1 000 000 000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		France Example-4
//		1 000 000 000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		France Example-6
//			0,00
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//	A pointer to a NumStrFormatSpec instance. All  member
//	variable data fields in this object will be replaced
//	by data values configured from the input parameter
//	described below.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setSignedNStrFmtFrance(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setSignedNStrFmtFrance()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewFrance(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewFrenchDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		leadingNumberFieldSymbolPosition:  NumFieldSymPos.InsideNumField(),
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols:             RuneArrayDto{},
		leadingNumberFieldSymbolPosition: 0,

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols:             RuneArrayDto{},
		leadingNumberFieldSymbolPosition: 0,

		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

//	setSignedNStrFmtGermany
//
//	Deletes and resets the member variable data values
//	stored in the instance of NumStrFormatSpec passed
//	as input parameter 'signedNumFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the Germany.
//
//	Within in the European Union, many, if not most, of
//	the member countries subscribe to the Signed Number
//	String formatting standards implemented by either
//	France or Germany.
//
// ----------------------------------------------------------------
//
// # Reference:
//
// https://freeformatter.com/germany-standards-code-snippets.html
//
// https://www.evertype.com/standards/euro/formats.html
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	NumStrFormatSpec input	parameter, 'numStrFmtSpec', will
//	be deleted and replaced by Number String formatting
//	parameters typically applied in Germany.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	comma character (','):
//
//		German Example-1
//			123,45 (The fractional digits are "45")
//
//	The integer group separator is a period character
//	('.').
//
//	The integer group specification is set to
//	'thousands'. This means that integer digits will be
//	separated into 'thousands' with each group containing
//	three digits each:
//
//		German Example-2:
//
//		1.000.000.000,00
//
//	The negative number sign is set to a trailing minus
//	sign ('-').
//
//		German Example-4
//		1.000.000,00-
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//	The zero number format is set to a blank or empty
//	string ("").
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//	A pointer to a NumStrFormatSpec instance. All  member
//	variable data fields in this object will be replaced
//	by data values configured from the input parameter
//	described below.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setSignedNStrFmtGermany(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setSignedNStrFmtGermany()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewGermany(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewGermanDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{

		leadingNumberSymbols:             RuneArrayDto{},
		leadingNumberFieldSymbolPosition: 0,

		trailingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		trailingNumberFieldSymbolPosition: NumFieldSymPos.InsideNumField(),
		lock:                              nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols:              RuneArrayDto{},
		leadingNumberFieldSymbolPosition:  0,
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols:              RuneArrayDto{},
		leadingNumberFieldSymbolPosition:  0,
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
}

// setSignedNStrFmtUS
//
// Deletes and resets the member variable data values
// stored in the instance of NumStrFormatSpec passed
// as input parameter 'signedNumFmtSpec'.
//
//	Reconfigures the current instance of NumStrFormatSpec
//	using Number String formatting conventions typically
//	applied in the US (United States).
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the current
//	instance of NumStrFormatSpec will be deleted and replaced
//	by Number String formatting parameters typically applied
//	in the US (United States).
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	The radix point or decimal separator is set to the
//	period character ('.').
//
//		United States Example-1
//			123.45 (The fractional digits are "45")
//
//	The integer group separator is a comma character
//	(',').
//
//	The integer group specification is set to 'thousands'.
//	This means that integer digits will be separated into
//	'thousands' with each group containing three digits each:
//
//		United States Example-2
//			1,000,000,000
//
//	The negative number sign is set to a leading minus sign
//	('-').
//
//		United States Example-3
//			-1,000,000,000
//
//	The positive number sign is set to a blank or empty
//	string ("").
//
//		United States Example-4
//			1,000,000,000
//
//	The zero number format is set to a blank or empty
//	string ("").
//
//		United States Example-5
//			0
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//		A pointer to a NumStrFormatSpec instance. All
//		member variable data fields in this object will
//		be replaced by data values configured from the
//		input parameter described below.
//
//	numberFieldSpec				NumStrNumberFieldSpec
//
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// -----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (nStrNumberFieldSpecNanobot *numStrFmtSpecNanobot) setSignedNStrFmtUS(
	numStrFmtSpec *NumStrFormatSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	nStrNumberFieldSpecNanobot.lock.Lock()

	defer nStrNumberFieldSpecNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecNanobot."+
			"setSignedNStrFmtUS()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var decSeparator DecimalSeparatorSpec

	decSeparator,
		err = new(DecimalSeparatorSpec).NewUS(
		ePrefix.XCpy("decSeparator"))

	if err != nil {
		return err
	}

	var intSeparatorSpec IntegerSeparatorSpec

	intSeparatorSpec,
		err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
		ePrefix.XCpy("intSeparatorSpec"))

	if err != nil {
		return err
	}

	var negativeNumberSign NumStrNumberSymbolSpec

	negativeNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols: RuneArrayDto{
			CharsArray:     []rune{'-'},
			Description1:   "",
			Description2:   "",
			charSearchType: CharSearchType.LinearTargetStartingIndex(),
			lock:           nil,
		},
		leadingNumberFieldSymbolPosition:  NumFieldSymPos.InsideNumField(),
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var positiveNumberSign NumStrNumberSymbolSpec

	positiveNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols:              RuneArrayDto{},
		leadingNumberFieldSymbolPosition:  0,
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	var zeroNumberSign NumStrNumberSymbolSpec

	zeroNumberSign = NumStrNumberSymbolSpec{
		leadingNumberSymbols:              RuneArrayDto{},
		leadingNumberFieldSymbolPosition:  0,
		trailingNumberSymbols:             RuneArrayDto{},
		trailingNumberFieldSymbolPosition: 0,
		lock:                              nil,
	}

	return new(numStrFmtSpecAtom).setNStrFmtComponents(
		numStrFmtSpec,
		decSeparator,
		intSeparatorSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		numberFieldSpec,
		ePrefix.XCpy("numStrFmtSpec<-"))
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

	signedNumFmtSpec.intSeparatorSpec.Empty()

	signedNumFmtSpec.positiveNumberSign.Empty()

	signedNumFmtSpec.negativeNumberSign.Empty()

	signedNumFmtSpec.numberFieldSpec.Empty()

	signedNumFmtSpec.zeroNumberSign.Empty()

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

	areEqual,
		_ := signedNumFmtSpec1.intSeparatorSpec.Equal(
		&signedNumFmtSpec2.intSeparatorSpec,
		nil)

	if !areEqual {

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

	if !signedNumFmtSpec1.zeroNumberSign.Equal(
		&signedNumFmtSpec2.zeroNumberSign) {

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

	err = new(decimalSepSpecNanobot).
		copyDecimalSeparator(
			&signedNumFmt.decSeparator,
			&decSeparatorSpec,
			ePrefix.XCpy(
				"signedNumFmt.decSeparatorSpec<-"+
					"decSeparatorSpec"))

	return err
}

// setIntegerGroupingParams - Deletes and resets the member
// variable data value for 'NumStrFormatSpec.intSeparatorSpec'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	signedNumFmt				*NumStrFormatSpec
//		A pointer to an instance of NumStrFormatSpec.
//		The member variable 'signedNumFmt.intSeparatorSpec'
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
//		and 'intSeparatorSpec' is NOT equal to
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

	signedNumFmt.intSeparatorSpec.Empty()

	err = signedNumFmt.intSeparatorSpec.SetIntGroupEnumRunes(
		intGroupingType,
		intGroupingChars,
		ePrefix.XCpy(
			"signedNumFmt.intSeparatorSpec<-"))

	return err
}

// setIntegerGroupingSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.intSeparatorSpec'
// contained in the instance of NumStrFormatSpec
// passed as an input parameter.
//
// This method receives an instance of IntegerSeparatorSpec
// and copies the member variable data values to
// 'signedNumFmt.intSeparatorSpec'.
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
//		in the Integer Separator Specification. supplied by
//		input parameter, 'intSeparatorSpec'.
//
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		An instance of IntegerSeparatorSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.intSeparatorSpec'.
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
	intSeparatorSpec IntegerSeparatorSpec,
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

	return signedNumFmt.intSeparatorSpec.CopyIn(
		&intSeparatorSpec,
		ePrefix.XCpy(
			"signedNumFmt.intSeparatorSpec<-"+
				"intSeparatorSpec"))

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

	err = new(numStrNumberSymbolSpecNanobot).
		copyNStrNumberSymbolSpec(
			&signedNumFmt.negativeNumberSign,
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
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
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

	return new(numStrNumberFieldSpecNanobot).
		copyNStrNumberFieldSpec(
			&signedNumFmt.numberFieldSpec,
			&numberFieldSpec,
			ePrefix.XCpy(
				"signedNumFmt<-"+
					"numberFieldSpec"))
}

// setNStrFmtComponents
//
// Deletes and resets the member variable data values
// stored in the instance of NumStrFormatSpec passed
// as input parameter 'signedNumFmtSpec'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that the data fields contained in the
//	NumStrFormatSpec instance passed by input parameter
//	'numStrFmtSpec'  will be deleted and replaced by
//	values generated from the other input parameters.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
// numStrFmtSpec				*NumStrFormatSpec
//
//	A pointer to a NumStrFormatSpec instance. All  member
//	variable data fields in this object will be replaced
//	by data values configured from the input parameter
//	described below.
//
//	decSeparator				DecimalSeparatorSpec
//
//		This structure contains the radix point or decimal
//		separator character(s) (a.k.a. decimal point)
//		which be used to separate integer and fractional
//		digits within a formatted Number String.
//
//	intSeparatorSpec				IntegerSeparatorSpec
//
//		Number String Integer Separator Specification. This
//		type encapsulates the parameters required to format
//		integer grouping and separation within a Number
//		String.
//
//	negativeNumberSign			NumStrNumberSymbolSpec
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with negative numeric values.
//
//	positiveNumberSign			NumStrNumberSymbolSpec
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with positive numeric values.
//
//	zeroNumberSign			NumStrNumberSymbolSpec
//		This Number String Symbol Specification contains
//		all the characters used to format number sign
//		symbols and currency symbols for Number Strings
//		with zero numeric values.
//
//	numberFieldSpec			NumStrNumberFieldSpec
//		This Number Field Specification contains all
//		parameters necessary to format a Number String
//		within a larger Number Field. In addition to
//		specifying the length of number field, this
//		object contains justification specifications
//		for centering, left justifying or right
//		justifying a Number String within a Number
//		Field.
//
//		type NumStrNumberFieldSpec struct {
//
//			fieldLength int
//
//				This parameter defines the length of the
//				text field in which the numeric value will
//				be displayed within a number string.
//
//				If 'fieldLength' is less than the length
//				of the numeric value string, it will be
//				automatically set equal to the length of
//				that numeric value string.
//
//				To automatically set the value of
//				'fieldLength' to the string length of the
//				numeric value, set this parameter to a
//				value of minus one (-1).
//
//				If this parameter is submitted with a
//				value less than minus one (-1) or greater
//				than 1-million (1,000,000), an error will
//				be returned.
//
//			fieldJustification TextJustify
//
//				An enumeration which specifies the
//				justification of the numeric value string
//				within the number field length specified
//				by data field 'fieldLength'.
//
//				Text justification can only be evaluated in
//				the context of a number string, field length
//				and a 'textJustification' object of type
//				TextJustify. This is because number strings
//				with a field length equal to or less than the
//				length of the numeric value string never use
//				text justification. In these cases, text
//				justification is completely ignored.
//
//				If the field length parameter ('fieldLength')
//				is greater than the length of the numeric
//				value string, text justification must be equal
//				to one of these three valid values:
//
//				          TextJustify(0).Left()
//				          TextJustify(0).Right()
//				          TextJustify(0).Center()
//
//				You can also use the abbreviated text
//				justification enumeration syntax as follows:
//
//				          TxtJustify.Left()
//				          TxtJustify.Right()
//				          TxtJustify.Center()
//		}
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
// -----------------------------------------------------------------
//
// # Return Values
//
//	err								error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setNStrFmtComponents(
	numStrFmtSpec *NumStrFormatSpec,
	decSeparator DecimalSeparatorSpec,
	intSeparatorSpec IntegerSeparatorSpec,
	negativeNumberSign NumStrNumberSymbolSpec,
	positiveNumberSign NumStrNumberSymbolSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
	numberFieldSpec NumStrNumberFieldSpec,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	signedNumFmtSpecAtom.lock.Lock()

	defer signedNumFmtSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrFmtSpecAtom."+
			"setNStrFmtComponents()",
		"")

	if err != nil {
		return err
	}

	if numStrFmtSpec == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrFmtSpec' is invalid!\n"+
			"'numStrFmtSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	err = numStrFmtSpec.decSeparator.CopyIn(
		&decSeparator,
		ePrefix.XCpy(
			"decSeparator->"))

	if err != nil {
		return err
	}

	err = numStrFmtSpec.intSeparatorSpec.CopyIn(
		&intSeparatorSpec,
		ePrefix.XCpy(
			"intSeparatorSpec->"))

	if err != nil {
		return err
	}

	err = numStrFmtSpec.negativeNumberSign.CopyIn(
		&negativeNumberSign,
		ePrefix.XCpy(
			"negativeNumberSign->"))

	if err != nil {
		return err
	}

	err = numStrFmtSpec.positiveNumberSign.CopyIn(
		&positiveNumberSign,
		ePrefix.XCpy(
			"positiveNumberSign->"))

	if err != nil {
		return err
	}

	err = numStrFmtSpec.zeroNumberSign.CopyIn(
		&zeroNumberSign,
		ePrefix.XCpy(
			"zeroNumberSign->"))

	if err != nil {
		return err
	}

	err = numStrFmtSpec.numberFieldSpec.CopyIn(
		&numberFieldSpec,
		ePrefix.XCpy(
			"numberFieldSpec->"))

	return err
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

	err = new(numStrNumberSymbolSpecNanobot).
		copyNStrNumberSymbolSpec(
			&signedNumFmt.positiveNumberSign,
			&positiveNumberSign,
			ePrefix.XCpy(
				"signedNumFmt.positiveNumberSign<-"+
					"positiveNumberSign"))

	return err
}

// setZeroNumberSign - Deletes and resets the member variable
// data value for 'NumStrFormatSpec.zeroNumberSign'
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
//	leadingZeroNumSymbols			[]rune
//
//		An array of runes containing the character or
//		characters which will be formatted and displayed
//		in front of a zero numeric value in a number
//		string.
//
//	trailingZeroNumSymbols			[]rune
//
//		An array of runes containing the character or
//		characters which will be formatted and displayed
//		after a zero numeric value in a number
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
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setZeroNumberSign(
	signedNumFmt *NumStrFormatSpec,
	leadingZeroNumSign []rune,
	leadingZeroNumFieldSymPosition NumberFieldSymbolPosition,
	trailingZeroNumSign []rune,
	trailingZeroNumFieldSymPosition NumberFieldSymbolPosition,
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
			"setZeroNumberSign()",
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

	lenLeadingZeroNumSign := len(leadingZeroNumSign)
	lenTrailingZeroNumSign := len(trailingZeroNumSign)

	if lenLeadingZeroNumSign == 0 &&
		lenTrailingZeroNumSign == 0 {

		signedNumFmt.zeroNumberSign.SetNOP()

		return err
	}

	signedNumFmt.zeroNumberSign.Empty()

	if lenLeadingZeroNumSign > 0 {

		err = signedNumFmt.zeroNumberSign.
			SetLeadingNumberSymbolRunes(
				leadingZeroNumSign,
				leadingZeroNumFieldSymPosition,
				ePrefix.XCpy(
					"signedNumFmt.zeroNumberSign"+
						"<-leadingZeroNumSymbols"))

		if err != nil {
			return err
		}
	}

	if lenTrailingZeroNumSign > 0 {

		err = signedNumFmt.zeroNumberSign.
			SetTrailingNumberSymbolRunes(
				trailingZeroNumSign,
				trailingZeroNumFieldSymPosition,
				ePrefix.XCpy(
					"signedNumFmt.zeroNumberSign<-"+
						"trailingZeroNumSymbols"))

		if err != nil {
			return err
		}
	}

	return err
}

// setZeroNumberSignSpec - Deletes and resets the member
// variable data value for 'signedNumFmt.zeroNumberSign'
// contained in the instance of NumStrFormatSpec passed as
// an input parameter.
//
// This method receives an instance of 'NumStrNumberSymbolSpec' and
// copies the member variable data values to
// 'signedNumFmt.zeroNumberSign'.
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
//	zeroNumberSign 			NumStrNumberSymbolSpec
//
//		An instance of NumStrNumberSymbolSpec. The member
//		variable data values contained in this instance
//		will be copied to:
//			'signedNumFmt.zeroNumberSign'.
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
func (signedNumFmtSpecAtom *numStrFmtSpecAtom) setZeroNumberSignSpec(
	signedNumFmt *NumStrFormatSpec,
	zeroNumberSign NumStrNumberSymbolSpec,
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
			"setZeroNumberSignSpec()",
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

	if zeroNumberSign.IsNOP() {

		signedNumFmt.zeroNumberSign.SetNOP()

		return err
	}

	err = new(numStrNumberSymbolSpecNanobot).
		copyNStrNumberSymbolSpec(
			&signedNumFmt.zeroNumberSign,
			&zeroNumberSign,
			ePrefix.XCpy(
				"signedNumFmt.zeroNumberSign<-"+
					"zeroNumberSign"))

	return err
}

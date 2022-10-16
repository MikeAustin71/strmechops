package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// NumStrNumberFieldSpec
//
// Number String Number Field Specification. This type
// contains all the parameters required to format a
// numeric value withing a text field for display as a
// number string.
type NumStrNumberFieldSpec struct {
	fieldLength int
	//	This parameter defines the length of the text
	//	field in which the numeric value will be
	//	displayed within a number string.
	//
	//	If 'fieldLength' is less than the length of the
	//	numeric value string, it will be automatically
	//	set equal to the length of that numeric value
	//	string.
	//
	//	To automatically set the value of 'fieldLength'
	//	to the string length of the numeric value, set
	//	this parameter to a value of minus one (-1).
	//
	//	If this parameter is submitted with a value less
	//	than minus one (-1) or greater than 1-million
	//	(1,000,000), an error will be returned.

	fieldJustification TextJustify
	//	An enumeration which specifies the
	//	justification of the numeric value string within
	//	the number field length specified by data field
	//	'fieldLength'.
	//
	//	Text justification can only be evaluated in the
	//	context of a number string, field length and a
	//	'textJustification' object of type TextJustify.
	//	This is because number strings with a field length
	//	equal to or less than the length of the numeric
	//	value string never use text justification. In
	//	these cases, text justification is completely
	//	ignored.
	//
	//	If the field length parameter ('fieldLength') is
	//	greater than the length of the numeric value
	//	string, text justification must be equal to one of
	//	these three valid values:
	//
	//	          TextJustify(0).Left()
	//	          TextJustify(0).Right()
	//	          TextJustify(0).Center()
	//
	//	You can also use the abbreviated text justification
	//	enumeration syntax as follows:
	//
	//	          TxtJustify.Left()
	//	          TxtJustify.Right()
	//	          TxtJustify.Center()

	lock *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// NumStrNumberFieldSpec ('incomingNStrNumFieldSpec')
// to the data fields of the current NumStrNumberFieldSpec
// instance ('nStrNumberFieldSpec').
//
// ----------------------------------------------------------------
//
// IMPORTANT
// All the data fields in current NumStrNumberFieldSpec
// instance ('nStrNumberFieldSpec') will be deleted and
// overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//			 incomingNStrNumFieldSpec *NumStrNumberFieldSpec
//			   - A pointer to an instance of NumStrNumberFieldSpec.
//			     This method will NOT change the values of internal member
//			     variables contained in this instance.
//
//			     All data values in this NumStrNumberFieldSpec instance
//			     will be copied to current NumStrNumberFieldSpec
//			     instance ('nStrNumberFieldSpec').
//
//			     If parameter 'incomingNStrNumFieldSpec' is determined to
//			     be invalid, an error will be returned.
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
// ----------------------------------------------------------------
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
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) CopyIn(
	incomingNStrNumFieldSpec *NumStrNumberFieldSpec,
	errorPrefix interface{}) error {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberFieldSpec."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberFieldSpecNanobot).
		copyNStrNumberFieldSpec(
			nStrNumberFieldSpec,
			incomingNStrNumFieldSpec,
			ePrefix.XCpy(
				"nStrNumberFieldSpec<-"+
					"incomingNStrNumFieldSpec"))
}

// CopyOut - Returns a deep copy of the current
// NumStrNumberFieldSpec instance.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 errorPrefix                interface{}
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
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyNumFieldSpec		NumStrNumberFieldSpec
//		If this method completes successfully and no errors are
//		encountered, this parameter will return a deep copy of
//		the current NumStrNumberFieldSpec instance.
//
//
//	err							error
//		If the method completes successfully and no errors are
//		encountered this return value is set to 'nil'. Otherwise,
//		if errors are encountered, this return value will contain
//		an appropriate error message.
//
//		If an error message is returned, the text value of input
//		parameter 'errorPrefix' will be inserted or prefixed at
//		the beginning of the error message.
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) CopyOut(
	errorPrefix interface{}) (
	deepCopyNumFieldSpec NumStrNumberFieldSpec,
	err error) {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberFieldSpec."+
			"CopyOut()",
		"")

	if err != nil {
		return deepCopyNumFieldSpec, err
	}

	err = new(numStrNumberFieldSpecNanobot).
		copyNStrNumberFieldSpec(
			&deepCopyNumFieldSpec,
			nStrNumberFieldSpec,
			ePrefix.XCpy(
				"deepCopyNumFieldSpec<-"+
					"nStrNumberFieldSpec"))

	return deepCopyNumFieldSpec, err
}

// Empty - Resets all internal member variables for the current
// instance of NumStrNumberFieldSpec to their initial or zero
// values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all pre-existing internal member
// variable data values in the current instance of
// NumStrNumberFieldSpec.
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
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) Empty() {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	new(numStrNumberFieldSpecAtom).empty(
		nStrNumberFieldSpec)

	nStrNumberFieldSpec.lock.Unlock()

	nStrNumberFieldSpec.lock = nil

}

// Equal - Receives a pointer to another instance of
// NumStrNumberFieldSpec and proceeds to compare its internal
// member variables to those of the current
// NumStrNumberFieldSpec instance in order to determine if
// they are equivalent.
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
//	incomingNStrNumFieldSpec	*NumStrNumberFieldSpec
//		A pointer to an external instance of
//		NumStrNumberFieldSpec. The internal member variable
//		data values in this instance will be compared to those
//		in the current instance of NumStrNumberFieldSpec. The
//		results of this comparison will be returned to the
//		calling function as a boolean value.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//		If the internal member variable data values contained in
//		input parameter 'incomingNStrNumFieldSpec' are equivalent
//		in all respects to those contained in the current
//		instance of 'NumStrNumberFieldSpec', this return value
//		will be set to 'true'.
//
//		Otherwise, this method will return 'false'.
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) Equal(
	incomingNStrNumFieldSpec *NumStrNumberFieldSpec) bool {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	return new(numStrNumberFieldSpecAtom).equal(
		nStrNumberFieldSpec,
		incomingNStrNumFieldSpec)
}

// GetNumFieldLength - Returns the field length value contained in
// the current instance of NumStrNumberFieldSpec.
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) GetNumFieldLength() int {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	return nStrNumberFieldSpec.fieldLength
}

// GetNumFieldJustification - Returns the text justification
// specification for the current instance of
// NumStrNumberFieldSpec.
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) GetNumFieldJustification() TextJustify {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	return nStrNumberFieldSpec.fieldJustification
}

// NewFieldSpec - Creates and returns new instance of
// NumStrNumberFieldSpec.
//
// ----------------------------------------------------------------
//
// Input Parameters
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
//	errorPrefix					interface{}
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
// Return Values
//
//	newNStrNumFieldSpec			NumStrNumberFieldSpec
//		If this method completes successfully, a new instance of
//		NumStrNumberFieldSpec will be created and returned.
//
//	err							error
//		If this method completes successfully, the returned error
//		Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the returned
//		error Type will encapsulate an error message. This
//		returned error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'. The
//		'errorPrefix' text will be attached to the beginning of
//		the error message.
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) NewFieldSpec(
	fieldLength int,
	fieldJustification TextJustify,
	errorPrefix interface{}) (
	newNStrNumFieldSpec NumStrNumberFieldSpec,
	err error) {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberFieldSpec."+
			"NewFieldSpec()",
		"")

	if err != nil {
		return newNStrNumFieldSpec, err
	}

	err = new(numStrNumberFieldSpecNanobot).
		setNStrNumberFieldSpec(
			&newNStrNumFieldSpec,
			fieldLength,
			fieldJustification,
			ePrefix.XCpy(
				"newNStrNumFieldSpec"))

	return newNStrNumFieldSpec, err
}

//	NewNOP
//
//	'NOP' is a computer science term that stands for
//	"No Operation". In this context, it means that
//	method 'NewNOP' is returning an instance of
//	NumStrNumberFieldSpec which will NOT construct
//	a number field within which a number string
//	is displayed. This means the length of the
//	number field will be equal to the length of
//	the number string and no text justification
//	will be applied.
//
// ----------------------------------------------------------------
//
// # Defaults
//
//	newNumberField.fieldLength = -1
//
//	newNumberField.fieldJustification = TxtJustify.None()
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	newNumberField				NumStrNumberFieldSpec
//
//		This Number Field Specification configuration will
//		set the length of a Number String Number Field to
//		the length of the number string.
//
//		Field length and Text Justification are automatically
//		set to the following values:
//
//			newNumberField.fieldLength = -1
//
//			newNumberField.fieldJustification = TxtJustify.None()
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) NewNOP() (newNumberField NumStrNumberFieldSpec) {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	newNumberField.fieldLength = -1

	newNumberField.fieldJustification = TxtJustify.None()

	return newNumberField
}

// SetFieldSpec - Deletes and overwrites all member variable
// data values in the current instance of NumStrNumberFieldSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite all pre-existing data
// values in the current instance of NumStrNumberFieldSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	errorPrefix					interface{}
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
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) SetFieldSpec(
	fieldLength int,
	fieldJustification TextJustify,
	errorPrefix interface{}) error {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberFieldSpec."+
			"SetFieldSpec()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberFieldSpecNanobot).
		setNStrNumberFieldSpec(
			nStrNumberFieldSpec,
			fieldLength,
			fieldJustification,
			ePrefix.XCpy(
				"nStrNumberFieldSpec"))
}

// SetFieldLength - Deletes and resets the member variable
// data value for 'NumStrNumberFieldSpec.fieldLength' in
// the current instance of NumStrNumberFieldSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and reset the pre-existing data
// value for the 'NumStrNumberFieldSpec.fieldLength'
// member variable contained in the current instance of
// NumStrNumberFieldSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
//	 errorPrefix                interface{}
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
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) SetFieldLength(
	fieldLength int,
	errorPrefix interface{}) error {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberFieldSpec."+
			"SetFieldLength()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberFieldSpecAtom).
		setNStrNumberFieldLength(
			nStrNumberFieldSpec,
			fieldLength,
			ePrefix.XCpy(
				"nStrNumberFieldSpec.fieldLength"+
					"<-fieldLength"))
}

// SetFieldJustification - Deletes and resets the member variable
// data value for 'NumStrNumberFieldSpec.fieldJustification' in
// the current instance of NumStrNumberFieldSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and reset the pre-existing data
// value for the 'NumStrNumberFieldSpec.fieldJustification'
// member variable contained in the current instance of
// NumStrNumberFieldSpec.
//
// ----------------------------------------------------------------
//
// # Input Parameters
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
func (nStrNumberFieldSpec *NumStrNumberFieldSpec) SetFieldJustification(
	fieldJustification TextJustify,
	errorPrefix interface{}) error {

	if nStrNumberFieldSpec.lock == nil {
		nStrNumberFieldSpec.lock = new(sync.Mutex)
	}

	nStrNumberFieldSpec.lock.Lock()

	defer nStrNumberFieldSpec.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"NumStrNumberFieldSpec."+
			"SetFieldJustification()",
		"")

	if err != nil {
		return err
	}

	return new(numStrNumberFieldSpecAtom).
		setNStrNumberFieldJustification(
			nStrNumberFieldSpec,
			fieldJustification,
			ePrefix.XCpy(
				"nStrNumberFieldSpec.fieldJustification"+
					"<-fieldJustification"))
}

// numStrNumberFieldSpecNanobot - This type provides
// helper methods for NumStrNumberFieldSpec
type numStrNumberFieldSpecNanobot struct {
	lock *sync.Mutex
}

// copyNStrNumberFieldSpec - Copies all data from input parameter
// 'sourceNumberFieldSpec' to input parameter
// 'destinationNumberFieldSpec'. Both instances are of type
// NumStrNumberFieldSpec.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Be advised that the data fields in
// 'destinationNumberFieldSpec' will be deleted and overwritten.
//
// Also, NO data validation is performed on 'sourceNumberFieldSpec'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationNumberFieldSpec  *NumStrNumberFieldSpec
//	   - A pointer to a NumStrNumberFieldSpec instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceNumberFieldSpec'.
//
//	     'destinationNumberFieldSpec' is the destination for this
//	     copy operation.
//
//
//	sourceNumberFieldSpec       *NumStrNumberFieldSpec
//	   - A pointer to another NumStrNumberFieldSpec
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationNumberFieldSpec'.
//
//	     'sourceNumberFieldSpec' is the source for this copy
//	     operation.
//
//	     No data validation is performed on 'sourceNumberFieldSpec'.
//
//
//	errPrefDto                     *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrNumberFieldSpecNanobot *numStrNumberFieldSpecNanobot) copyNStrNumberFieldSpec(
	destinationNumberFieldSpec *NumStrNumberFieldSpec,
	sourceNumberFieldSpec *NumStrNumberFieldSpec,
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
		"numStrNumberFieldSpecNanobot."+
			"copyNStrNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	if destinationNumberFieldSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationNumberFieldSpec' is invalid!\n"+
			"'destinationNumberFieldSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if sourceNumberFieldSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceNumberFieldSpec' is invalid!\n"+
			"'sourceNumberFieldSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	new(numStrNumberFieldSpecAtom).empty(
		destinationNumberFieldSpec)

	destinationNumberFieldSpec.fieldLength =
		sourceNumberFieldSpec.fieldLength

	destinationNumberFieldSpec.fieldJustification =
		sourceNumberFieldSpec.fieldJustification

	return err
}

// setNStrNumberFieldSpec - Deletes and resets all member
// variable data values contained in the instance of
// NumStrNumberFieldSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	nStrNumFieldSpec           *NumStrNumberFieldSpec
//	    - A pointer to an instance of NumStrNumberFieldSpec.
//	      All the member variable data values in this instance
//	      will be deleted and reset according to the data
//	      extracted from the following input parameters.
//
//	 fieldLength                int
//	    - This parameter defines the length of the text field in
//	      which the numeric value will be displayed within a
//	      number string.
//
//	      If 'fieldLength' is less than the length of the numeric
//	      value string, it will be automatically set equal to the
//	      length of that numeric value string.
//
//	      To automatically set the value of fieldLength to the string
//	      length of the numeric value, set this parameter to a value
//	      of minus one (-1).
//
//	      If this parameter is submitted with a value less than minus
//	      one (-1) or greater than 1-million (1,000,000), an error will
//	      be returned.
//
//	 fieldJustification         TextJustify
//	    - An enumeration which specifies the justification of the
//	      numeric value string within the number field length specified
//	      by input parameter 'fieldLength'.
//
//	      Text justification can only be evaluated in the context of
//	      a number string, field length and a 'textJustification'
//	      object of type TextJustify. This is because number strings
//	      with a field length equal to or less than the length of the
//	      numeric value string never use text justification. In these
//	      cases, text justification is completely ignored.
//
//	      If the field length parameter ('fieldLength') is greater
//	      than the length of the numeric value string, text
//	      justification must be equal to one of these
//	      three valid values:
//	                TextJustify(0).Left()
//	                TextJustify(0).Right()
//	                TextJustify(0).Center()
//
//	      You can also use the abbreviated text justification
//	      enumeration syntax as follows:
//
//	                TxtJustify.Left()
//	                TxtJustify.Right()
//	                TxtJustify.Center()
//
//
//	 errPrefDto                 *ePref.ErrPrefixDto
//	    - This object encapsulates an error prefix string which is
//	      included in all returned error messages. Usually, it
//	      contains the name of the calling method or methods listed
//	      as a function chain.
//
//	      If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//	      Type ErrPrefixDto is included in the 'errpref' software
//	      package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrNumberFieldSpecNanobot *numStrNumberFieldSpecNanobot) setNStrNumberFieldSpec(
	nStrNumFieldSpec *NumStrNumberFieldSpec,
	fieldLength int,
	fieldJustification TextJustify,
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
		"numStrNumberFieldSpecNanobot."+
			"setNStrNumberFieldSpec()",
		"")

	if err != nil {
		return err
	}

	new(numStrNumberFieldSpecAtom).empty(
		nStrNumFieldSpec)

	nStrNumFieldSpecAtom := numStrNumberFieldSpecAtom{}

	err = nStrNumFieldSpecAtom.setNStrNumberFieldLength(
		nStrNumFieldSpec,
		fieldLength,
		ePrefix.XCpy(
			"nStrNumFieldSpec.fieldLength<-"+
				"fieldLength"))

	if err != nil {
		return err
	}

	err = nStrNumFieldSpecAtom.setNStrNumberFieldJustification(
		nStrNumFieldSpec,
		fieldJustification,
		ePrefix.XCpy(
			"nStrNumFieldSpec.fieldJustification<-"+
				"fieldJustification"))

	if err != nil {
		return err
	}

	return err
}

// numStrNumberFieldSpecAtom - This type provides
// helper methods for NumStrNumberFieldSpec
type numStrNumberFieldSpecAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// NumStrNumberFieldSpec and proceeds to reset the
// data values for all member variables to their
// initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'nStrNumFieldSpec' will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	nStrNumFieldSpec           *NumStrNumberFieldSpec
//	   - A pointer to an instance of NumStrNumberFieldSpec.
//	     All the internal member variables contained in this
//	     instance will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (nStrNumFieldSpecAtom *numStrNumberFieldSpecAtom) empty(
	nStrNumFieldSpec *NumStrNumberFieldSpec) {

	if nStrNumFieldSpecAtom.lock == nil {
		nStrNumFieldSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumFieldSpecAtom.lock.Lock()

	defer nStrNumFieldSpecAtom.lock.Unlock()

	if nStrNumFieldSpec == nil {
		return
	}

	nStrNumFieldSpec.fieldLength = -2

	nStrNumFieldSpec.fieldJustification = TxtJustify.None()
}

// equal - Receives a pointer to two instances of
// NumStrNumberFieldSpec and proceeds to compare their
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
//	nStrNumFieldSpec1    *NumStrNumberFieldSpec
//	   - An instance of NumStrNumberFieldSpec.
//	     Internal member variables from 'nStrNumFieldSpec1'
//	     will be compared to those of 'nStrNumFieldSpec2' to
//	     determine if both instances are equivalent.
//
//
//	nStrNumFieldSpec2    *NumStrNumberFieldSpec
//	   - An instance of NumStrNumberFieldSpec.
//	     Internal member variables from 'nStrNumFieldSpec2'
//	     will be compared to those of 'nStrNumFieldSpec1' to
//	     determine if both instances are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'nStrNumFieldSpec1' and
//	     'nStrNumFieldSpec2' shows that all internal member
//	     variables are equivalent, this method will return a
//	     boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling
//	     function.
func (nStrNumFieldSpecAtom *numStrNumberFieldSpecAtom) equal(
	nStrNumFieldSpec1 *NumStrNumberFieldSpec,
	nStrNumFieldSpec2 *NumStrNumberFieldSpec) bool {

	if nStrNumFieldSpecAtom.lock == nil {
		nStrNumFieldSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumFieldSpecAtom.lock.Lock()

	defer nStrNumFieldSpecAtom.lock.Unlock()

	if nStrNumFieldSpec1 == nil ||
		nStrNumFieldSpec2 == nil {
		return false
	}

	if nStrNumFieldSpec1.fieldLength !=
		nStrNumFieldSpec2.fieldLength {

		return false
	}

	if nStrNumFieldSpec1.fieldJustification !=
		nStrNumFieldSpec2.fieldJustification {

		return false
	}

	return true
}

// setNStrNumberFieldLength - Deletes and resets the member
// variable data value for 'NumStrNumberFieldSpec.fieldLength'
// contained in the instance of NumStrNumberFieldSpec passed
// as an input parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and reset the pre-existing data
// value for the 'NumStrNumberFieldSpec.fieldLength'
// member variable contained in the instance of
// NumStrNumberFieldSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	nStrNumFieldSpec           *NumStrNumberFieldSpec
//	    - A pointer to an instance of NumStrNumberFieldSpec.
//	      The 'NumStrNumberFieldSpec.fieldLength' member
//	      variable data values in this instance
//	      will be deleted and reset to the value
//	      provided by input parameter, 'fieldLength'.
//
//	 fieldLength                int
//	    - This parameter defines the length of the text field in
//	      which the numeric value will be displayed within a
//	      number string.
//
//	      If 'fieldLength' is less than the length of the numeric
//	      value string, it will be automatically set equal to the
//	      length of that numeric value string.
//
//	      To automatically set the value of fieldLength to the string
//	      length of the numeric value, set this parameter to a value
//	      of minus one (-1).
//
//	      If this parameter is submitted with a value less than minus
//	      one (-1) or greater than 1-million (1,000,000), an error will
//	      be returned.
//
//	 errPrefDto                 *ePref.ErrPrefixDto
//	    - This object encapsulates an error prefix string which is
//	      included in all returned error messages. Usually, it
//	      contains the name of the calling method or methods listed
//	      as a function chain.
//
//	      If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//	      Type ErrPrefixDto is included in the 'errpref' software
//	      package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrNumFieldSpecAtom *numStrNumberFieldSpecAtom) setNStrNumberFieldLength(
	nStrNumFieldSpec *NumStrNumberFieldSpec,
	fieldLength int,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumFieldSpecAtom.lock == nil {
		nStrNumFieldSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumFieldSpecAtom.lock.Lock()

	defer nStrNumFieldSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberFieldSpecNanobot."+
			"setNStrNumberFieldLength()",
		"")

	if err != nil {
		return err
	}

	if nStrNumFieldSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumFieldSpec' is invalid!\n"+
			"'nStrNumFieldSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if fieldLength < -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLength' is invalid!\n"+
			"'fieldLength' has a value less than minus one (-1).\n"+
			"fieldLength = '%v'\n",
			ePrefix.String(),
			fieldLength)

		return err
	}

	if fieldLength > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLength' is invalid!\n"+
			"'fieldLength' has a value greater than one-million (1,000,000).\n"+
			"fieldLength = '%v'\n",
			ePrefix.String(),
			fieldLength)

		return err
	}

	nStrNumFieldSpec.fieldLength = fieldLength

	return err
}

// setNStrNumberFieldJustification - Deletes and resets the
// member variable data value for
// 'NumStrNumberFieldSpec.fieldJustification' contained in
// the instance of NumStrNumberFieldSpec passed as an input
// parameter.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete and overwrite the pre-existing data
// value for the 'NumStrNumberFieldSpec.fieldJustification'
// member variable contained in the instance of
// NumStrNumberFieldSpec passed as an input parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//		nStrNumFieldSpec           *NumStrNumberFieldSpec
//	    - A pointer to an instance of NumStrNumberFieldSpec.
//	      The 'NumStrNumberFieldSpec.fieldLength' member
//	      variable data values in this instance
//	      will be deleted and reset to the value
//	      provided by input parameter, 'fieldLength'.
//
//		 fieldJustification         TextJustify
//		    - An enumeration which specifies the justification of the
//		      numeric value string within the number field length specified
//		      by input parameter 'fieldLength'.
//
//		      Text justification can only be evaluated in the context of
//		      a number string, field length and a 'textJustification'
//		      object of type TextJustify. This is because number strings
//		      with a field length equal to or less than the length of the
//		      numeric value string never use text justification. In these
//		      cases, text justification is completely ignored.
//
//		      If the field length parameter ('fieldLength') is greater
//		      than the length of the numeric value string, text
//		      justification must be equal to one of these
//		      three valid values:
//		                TextJustify(0).Left()
//		                TextJustify(0).Right()
//		                TextJustify(0).Center()
//
//		      You can also use the abbreviated text justification
//		      enumeration syntax as follows:
//
//		                TxtJustify.Left()
//		                TxtJustify.Right()
//		                TxtJustify.Center()
//
//			 errPrefDto                 *ePref.ErrPrefixDto
//			    - This object encapsulates an error prefix string which is
//			      included in all returned error messages. Usually, it
//			      contains the name of the calling method or methods listed
//			      as a function chain.
//
//			      If no error prefix information is needed, set this
//			      parameter to 'nil'.
//
//			      Type ErrPrefixDto is included in the 'errpref' software
//			      package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (nStrNumFieldSpecAtom *numStrNumberFieldSpecAtom) setNStrNumberFieldJustification(
	nStrNumFieldSpec *NumStrNumberFieldSpec,
	fieldJustification TextJustify,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if nStrNumFieldSpecAtom.lock == nil {
		nStrNumFieldSpecAtom.lock = new(sync.Mutex)
	}

	nStrNumFieldSpecAtom.lock.Lock()

	defer nStrNumFieldSpecAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numStrNumberFieldSpecNanobot."+
			"setNStrNumberFieldJustification()",
		"")

	if err != nil {
		return err
	}

	if nStrNumFieldSpec == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'nStrNumFieldSpec' is invalid!\n"+
			"'nStrNumFieldSpec' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if !fieldJustification.XIsValid() &&
		fieldJustification != TxtJustify.None() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldJustification' is invalid!\n"+
			"'fieldJustification' string value  = '%v'\n"+
			"'fieldJustification' integer value = '%v'\n",
			ePrefix.String(),
			fieldJustification.String(),
			fieldJustification.XValueInt())

		return err

	}

	nStrNumFieldSpec.fieldJustification =
		fieldJustification

	return err
}

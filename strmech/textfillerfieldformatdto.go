package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// TextFillerFieldFormatDto
//
// Field Format Data Transfer Objects (Dto) are used
// to facilitate easy data entry which creating and
// configuring text lines strings for screen display,
// file output or printing.
//
// Type TextFillerFieldFormatDto encapsulates all the
// specification parameters necessary to create and
// construct Text Filler Field strings.
//
// Typically, filler fields are used as margins
// containing multiple white space characters, or line
// separators containing multiple dashes, equal signs
// or underscore characters. Filler fields consist of
// filler characters ('FillerChars') and the filler
// characters repeat count ('FillerCharsRepeatCount').
// A filler field is made up of one or more filler
// characters. These filler characters are repeated one
// or more times in order to construct the complete
// filler field as shown in the following examples:
//
//	Example 1:
//	 Filler Characters = "-"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "---"
//
//	Example 2:
//	 Filler Characters = "-*"
//	 Filler Characters Repeat Count = 3
//	 Formatted Text = "-*-*-*"
//
// The 'FillerCharsRepeatCount' integer value is the
// number times that 'fillerCharacters' is repeated in
// order to construct the Filler Text Field.
//
// Be advised that Filler Text Fields requires a
// 'FillerCharsRepeatCount' value greater than zero.
// 'FillerCharsRepeatCount' values less than or equal
// to zero constitute an error condition.
type TextFillerFieldFormatDto struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for this Text Filler Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this Text
	//	Filler Field.

	FillerChars string
	//	A string containing the text characters which
	//	will be included in the Text Filler Field. The
	//	final Text Filler Field content will be
	//	constructed from ths filler characters repeated
	//	one or more times as specified by the
	//	'FillerCharsRepeatCount' parameter.
	//
	//	The Text Filler Field final formatted text is
	//	equal to:
	//
	//		FillerChars X FillerCharsRepeatCount
	//
	//	        Example: FillerChars = "-*"
	//	                 FillerCharsRepeatCount = 3
	//	                 Final Text Filler Field = "-*-*-*"

	FillerCharsRepeatCount int
	//	Controls the number of times 'FillerChars' is
	//	repeated when constructing the final Text Filler
	//	Field. The actual length of the string which will
	//	populate the completed Text Filler Field is
	//	equal to the length of 'FillerChars' times the
	//	value of 'FillerCharsRepeatCount'.
	//
	//		Text Field Filler Length =
	//			Length of FillerChars X FillerCharsRepeatCount
	//
	//	        Example #1: FillerChars = "-*"
	//	                    FillerCharsRepeatCount = 3
	//	                    Final Text Filler Field = "-*-*-*"
	//
	//	        Example #2: FillerChars = "-"
	//	                    fillerRepeatCount = 3
	//	                    Final Text Filler Field = "---"

	RightMarginStr string
	//	One or more characters used to create a right
	//	margin for this Text Filler Field.
	//
	//	If this parameter is set to an empty string, no
	//	right margin will be configured for this Text
	//	Filler Field.

	lock *sync.Mutex
}

// CopyIn
//
// Copies all the data fields from an incoming instance
// of TextFillerFieldFormatDto
// ('incomingTxtFillerFieldFmtDto') to the corresponding
// data fields of the current TextFillerFieldFormatDto
// instance ('txtFillerFieldFmtDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	current instance of TextFillerFieldFormatDto
//	('txtFillerFieldFmtDto').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTxtFillerFieldFmtDto	*TextFillerFieldFormatDto
//
//		A pointer to an instance of
//		TextFillerFieldFormatDto.
//
//		All the internal data field values in this
//		instance will be copied to corresponding data
//		fields of the current TextFillerFieldFormatDto
//		instance.
//
//		The data fields contained in
//		'incomingTxtFillerFieldFmtDto' will NOT be
//		changed or modified.
//
//		If 'incomingTxtFillerFieldFmtDto' contains
//		invalid data values, an error will be returned.
//
//	errorPrefix						interface{}
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
//		7.	IBasicErrorPrefix
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
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) CopyIn(
	incomingTxtFillerFieldFmtDto *TextFillerFieldFormatDto,
	errorPrefix interface{}) error {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFillerFieldFormatDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textFillerFieldFormatDtoNanobot).copy(
		txtFillerFieldFmtDto,
		incomingTxtFillerFieldFmtDto,
		ePrefix.XCpy(
			"txtFillerFieldFmtDto<-"+
				"incomingTxtFillerFieldFmtDto"))
}

// CopyOut
//
// Returns a deep copy of the current
// TextFillerFieldFormatDto instance.
//
// If the current TextFillerFieldFormatDto instance
// contains invalid member variable data values, this
// method will return an error.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
//		7.	IBasicErrorPrefix
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
//	TextFillerFieldFormatDto
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextFillerFieldFormatDto instance.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) CopyOut(
	errorPrefix interface{}) (
	TextFillerFieldFormatDto,
	error) {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	newTxtFillerFieldFmtDto := TextFillerFieldFormatDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFillerFieldFormatDto."+
			"CopyOut()",
		"")

	if err != nil {
		return newTxtFillerFieldFmtDto, err
	}

	err = new(textFillerFieldFormatDtoNanobot).copy(
		&newTxtFillerFieldFmtDto,
		txtFillerFieldFmtDto,
		ePrefix.XCpy(
			"newTxtFillerFieldFmtDto<-"+
				"txtFillerFieldFmtDto"))

	return newTxtFillerFieldFmtDto, err
}

// CopyOutITextFieldFormat
//
// Returns a deep copy of the current
// TextFillerFieldFormatDto instance cast as an
// ITextFieldFormatDto interface object.
//
// If the current TextFillerFieldFormatDto instance
// contains invalid member variable data values, this
// method will return an error.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
//		7.	IBasicErrorPrefix
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
//	ITextFieldFormatDto
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextFillerFieldFormatDto instance cast as an
//		ITextFieldFormatDto interface object.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) CopyOutITextFieldFormat(
	errorPrefix interface{}) (
	ITextFieldFormatDto,
	error) {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	newTxtFillerFieldFmtDto := TextFillerFieldFormatDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFillerFieldFormatDto."+
			"CopyOutITextFieldFormat()",
		"")

	if err != nil {
		return ITextFieldFormatDto(&newTxtFillerFieldFmtDto), err
	}

	err = new(textFillerFieldFormatDtoNanobot).copy(
		&newTxtFillerFieldFmtDto,
		txtFillerFieldFmtDto,
		ePrefix.XCpy(
			"newTxtFillerFieldFmtDto<-"+
				"txtFillerFieldFmtDto"))

	return ITextFieldFormatDto(&newTxtFillerFieldFmtDto), err
}

// Empty
//
// Resets all internal member variables for the current
// instance of TextFillerFieldFormatDto to their zero or
// uninitialized states. This method will leave the
// current instance of TextFillerFieldFormatDto in an
// invalid state and unavailable for immediate reuse.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data
// values in the current instance of
// TextFillerFieldFormatDto. All member variable data
// values will be reset to their zero or uninitialized
// states.
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
//	NONE
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) Empty() {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	new(textFillerFieldFormatDtoAtom).
		empty(txtFillerFieldFmtDto)

	txtFillerFieldFmtDto.lock.Unlock()

	txtFillerFieldFmtDto.lock = nil

	return
}

// Equal
//
// Receives a pointer to another instance of
// TextFillerFieldFormatDto and proceeds to compare the
// member variables to those contained in the current
// TextFillerFieldFormatDto instance in order to
// determine if they are equivalent.
//
// A boolean flag showing the result of this comparison
// is returned. If the member variables of both instances
// are equal in all respects, this flag is set to 'true'.
// Otherwise, this method returns 'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTxtFieldFillerFmtDto	*TextFillerFieldFormatDto
//
//		A pointer to an incoming instance of
//		TextFillerFieldFormatDto. This method will
//		compare all member variable data values in this
//		instance against those contained in the current
//		instance of TextFillerFieldFormatDto. If the data
//		values in both instances are found to be equal in
//		all respects, this method will return a boolean
//		value of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the member variable data values contained in
//		input parameter 'incomingTxtFieldFillerFmtDto'
//		are equal in all respects to those contained in
//		the current instance of TextFillerFieldFormatDto,
//		this method will return a boolean value of
//		'true'. Otherwise, a value of 'false' will be
//		returned to the calling function.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) Equal(
	incomingTxtFieldFillerFmtDto *TextFillerFieldFormatDto) bool {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	return new(textFillerFieldFormatDtoAtom).equal(
		txtFillerFieldFmtDto,
		incomingTxtFieldFillerFmtDto)
}

// GetFieldContentTextLabel
//
// Converts the current TextFillerFieldFormatDto instance
// member variable, 'FillerChars', to an instance of
// TextFieldSpecLabel.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// The returned TextFieldSpecLabel will only contain
// the member variable 'FillerChars'. It will NOT
// contain the left and right margins.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
//		7.	IBasicErrorPrefix
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
//	TextFieldSpecLabel
//
//		If this method completes successfully, the Text
//		Field Contents extracted from the current
//		instance of TextFillerFieldFormatDto, will be
//		returned as text label of type TextFieldSpecLabel.
//
//		This returned text label will ONLY contain the
//		Text Field Contents ('FillerChars'). It will NOT
//		contain the left and right margin strings.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) GetFieldContentTextLabel(
	errorPrefix interface{}) (
	TextFieldSpecLabel,
	error) {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFillerFieldFormatDto."+
			"GetFieldContentTextLabel()",
		"")

	if err != nil {
		return TextFieldSpecLabel{}, err
	}

	return new(textFillerFieldFormatDtoMolecule).
		getFieldContentTextLabel(
			txtFillerFieldFmtDto,
			ePrefix.XCpy(
				"txtFillerFieldFmtDto"))
}

// GetFieldFormatDtoType
//
// Returns a string containing the name of this type
// ('TextFillerFieldFormatDto').
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) GetFieldFormatDtoType() string {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	return "TextFillerFieldFormatDto"
}

// GetFormattedTextFieldStr
//
// Returns a string containing the formatted text field
// generated from the current instance of
// TextFillerFieldFormatDto.
//
// The returned formatted text field string contains the
// left margin, field contents and right margin.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
//		7.	IBasicErrorPrefix
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
//	string
//
//		If this method completes successfully, the text
//		field specifications contained in the current
//		instance of TextFillerFieldFormatDto will be
//		converted to, and returned as, a formatted text
//		field string.
//
//		The returned text field string will contain the
//		left margin, text field contents and right margin
//		as those elements are defined in the current
//		instance of TextFillerFieldFormatDto.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) GetFormattedTextFieldStr(
	errorPrefix interface{}) (
	string,
	error) {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFillerFieldFormatDto."+
			"GetFormattedTextFieldStr()",
		"")

	if err != nil {
		return "", err
	}

	return new(textFillerFieldFormatDtoNanobot).
		getFormattedTextFieldStr(
			txtFillerFieldFmtDto,
			ePrefix.XCpy(
				"txtFillerFieldFmtDto"))
}

// GetLeftMarginLength
//
// Returns the length of the Left Margin String as an
// integer value.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) GetLeftMarginLength() int {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	return len(txtFillerFieldFmtDto.LeftMarginStr)
}

// GetLeftMarginStr
//
// Returns the Left Margin String.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) GetLeftMarginStr() string {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	return txtFillerFieldFmtDto.LeftMarginStr
}

// GetRightMarginLength
//
// Returns the length of the Right Margin String as an
// integer value.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) GetRightMarginLength() int {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	return len(txtFillerFieldFmtDto.RightMarginStr)
}

// GetRightMarginStr
//
// Returns the Right Margin String.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) GetRightMarginStr() string {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	return txtFillerFieldFmtDto.RightMarginStr
}

// IsValidInstance
//
// Performs a diagnostic review of the data values
// encapsulated in the current TextFillerFieldFormatDto
// instance to determine if they are valid.
//
// If all data elements evaluate as valid, this method
// returns 'true'. If any data element is invalid, this
// method returns 'false'.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isValid						bool
//
//		If all data elements encapsulated by the current
//		instance of TextFillerFieldFormatDto are valid,
//		this returned boolean value is set to 'true'. If
//		any data values are invalid, this return
//		parameter is set to 'false'.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) IsValidInstance() (
	isValid bool) {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	isValid,
		_ = new(textFillerFieldFormatDtoAtom).
		testValidityOfTextFillerFieldFmtDto(
			txtFillerFieldFmtDto,
			nil)

	return isValid
}

// IsValidInstanceError
//
// Performs a diagnostic review of the data values
// encapsulated in the current TextFillerFieldFormatDto
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method
// will return an error.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
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
//		7.	IBasicErrorPrefix
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
//		If any of the internal member data variables
//		contained in the current instance of
//		TextFillerFieldFormatDto are found to be invalid,
//		this method will return an error containing an
//		appropriate error message.
//
//		If an error message is returned, the returned
//		error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the
//		beginning of the error message.
func (txtFillerFieldFmtDto *TextFillerFieldFormatDto) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtFillerFieldFmtDto.lock == nil {
		txtFillerFieldFmtDto.lock = new(sync.Mutex)
	}

	txtFillerFieldFmtDto.lock.Lock()

	defer txtFillerFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFillerFieldFormatDto."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textFillerFieldFormatDtoAtom).
		testValidityOfTextFillerFieldFmtDto(
			txtFillerFieldFmtDto,
			ePrefix.XCpy(
				"txtFillerFieldFmtDto"))

	return err
}

// textFillerFieldFormatDtoNanobot - Provides helper
// methods for TextFillerFieldFormatDto.
type textFillerFieldFormatDtoNanobot struct {
	lock *sync.Mutex
}

// copy
//
// Copies all data from a source instance of
// TextFillerFieldFormatDto to a destination instance of
// TextFillerFieldFormatDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	TextFillerFieldFormatDto instance passed as input
//	parameter 'destinationTxtFieldFillerDto'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationTxtFieldFillerDto	*TextFillerFieldFormatDto
//
//		A pointer to an instance of
//		TextFillerFieldFormatDto.
//
//		Data extracted from input parameter
//		'sourceTxtFieldFillerDto' will be copied to this
//		input parameter, 'destinationTxtFieldFillerDto'.
//
//		'destinationTxtFieldFmtDto' is the destination
//		for this copy operation.
//
//		If this method completes successfully, all member
//		data variables encapsulated in
//		'destinationTxtFieldFmtDto' will be identical to
//		those contained in input parameter,
//		'sourceTxtFieldFmtDto'.
//
//		Be advised that the pre-existing data fields
//		contained within input parameter
//		'destinationTxtFieldFmtDto' will be overwritten
//		and deleted.
//
//	sourceTxtFieldFillerDto			*TextFillerFieldFormatDto
//
//		A pointer to an instance of
//		TextFillerFieldFormatDto.
//
//		All data values in this TextFillerFieldFormatDto
//		instance will be copied to input parameter
//		'destinationTxtFieldFillerDto'.
//
//		'sourceTxtFieldFillerDto' is the source of the
//		copy operation.
//
//		The original member variable data values
//		encapsulated within 'sourceTxtFieldFillerDto'
//		will remain unchanged and unmodified.
//
//		If 'sourceTxtFieldFillerDto' contains invalid
//		member data variables, this method will return
//		an error.
//
//	errPrefDto						*ePref.ErrPrefixDto
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (txtFillerFieldDtoNanobot *textFillerFieldFormatDtoNanobot) copy(
	destinationTxtFieldFillerDto *TextFillerFieldFormatDto,
	sourceTxtFieldFillerDto *TextFillerFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFillerFieldDtoNanobot.lock == nil {
		txtFillerFieldDtoNanobot.lock = new(sync.Mutex)
	}

	txtFillerFieldDtoNanobot.lock.Lock()

	defer txtFillerFieldDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFillerFieldFormatDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if destinationTxtFieldFillerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtFieldFillerDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTxtFieldFillerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtFieldFillerDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtFillerFieldFmtDtoAtom := textFillerFieldFormatDtoAtom{}

	_,
		err = txtFillerFieldFmtDtoAtom.
		testValidityOfTextFillerFieldFmtDto(
			sourceTxtFieldFillerDto,
			ePrefix.XCpy(
				"sourceTxtFieldFillerDto"))

	if err != nil {

		return err
	}

	txtFillerFieldFmtDtoAtom.empty(
		destinationTxtFieldFillerDto)

	destinationTxtFieldFillerDto.LeftMarginStr =
		sourceTxtFieldFillerDto.LeftMarginStr

	destinationTxtFieldFillerDto.FillerChars =
		sourceTxtFieldFillerDto.FillerChars

	destinationTxtFieldFillerDto.FillerCharsRepeatCount =
		sourceTxtFieldFillerDto.FillerCharsRepeatCount

	destinationTxtFieldFillerDto.RightMarginStr =
		sourceTxtFieldFillerDto.RightMarginStr

	return err
}

// getFormattedTextFieldStr
//
// Converts an instance of TextFillerFieldFormatDto to a
// formatted text field string.
//
// This formatted text field string contains the left
// margin, field contents and right margin.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFillerFieldFmtDto		*TextFillerFieldFormatDto
//
//		A pointer to an instance of
//		TextFillerFieldFormatDto.
//
//		The left and right margins as well as the member
//		variable 'FillerChars' will be processed and
//		converted to a formatted text field for use in
//		screen displays, file output and printing.
//
//		If input parameter 'txtFillerFieldFmtDto' is
//		found to contain invalid data values, an error
//		will be returned.
//
//		None of the data values in this instance will be
//		changed or modified.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If this method completes successfully, the input
//		parameter, 'txtFillerFieldFmtDto', will be
//		converted to, and returned as, a formatted string
//		of text.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtFillerFieldDtoNanobot *textFillerFieldFormatDtoNanobot) getFormattedTextFieldStr(
	txtFillerFieldFmtDto *TextFillerFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if txtFillerFieldDtoNanobot.lock == nil {
		txtFillerFieldDtoNanobot.lock = new(sync.Mutex)
	}

	txtFillerFieldDtoNanobot.lock.Lock()

	defer txtFillerFieldDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFillerFieldFormatDtoNanobot."+
			"getFormattedTextFieldStr()",
		"")

	if err != nil {

		return "", err

	}

	if txtFillerFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtFillerFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	strBuilder := new(strings.Builder)

	if len(txtFillerFieldFmtDto.LeftMarginStr) > 0 {

		strBuilder.WriteString(txtFillerFieldFmtDto.LeftMarginStr)

	}

	var textLabel TextFieldSpecLabel

	textLabel,
		err = new(textFillerFieldFormatDtoMolecule).
		getFieldContentTextLabel(
			txtFillerFieldFmtDto,
			ePrefix.XCpy(
				"txtFillerFieldFmtDto"))

	if err != nil {

		return "", err

	}

	strBuilder.WriteString(textLabel.GetTextLabel())

	if len(txtFillerFieldFmtDto.RightMarginStr) > 0 {

		strBuilder.WriteString(txtFillerFieldFmtDto.RightMarginStr)

	}

	return strBuilder.String(), err
}

// textFillerFieldFormatDtoMolecule - Provides helper
// methods for TextFillerFieldFormatDto.
type textFillerFieldFormatDtoMolecule struct {
	lock *sync.Mutex
}

// getFieldContentTextLabel
//
// Converts a TextFillerFieldFormatDto instance member
// variable, 'FillerChars', to an instance of
// TextFieldSpecLabel.
//
// The TextFillerFieldFormatDto instance is passed as
// input parameter, 'txtFieldFillerDto'.
//
// The returned TextFieldSpecLabel will only contain
// the member variable 'FillerChars'. It will NOT contain
// the left and right margins.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If input parameter 'txtFieldFillerDto', an instance of
//	TextFillerFieldFormatDto, is found to be invalid, an
//	error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFillerDto			*TextFillerFieldFormatDto
//
//		A pointer to an instance of
//		TextFillerFieldFormatDto.
//
//		The member variable 'FillerChars' will be
//		converted to a text label of type
//		TextFieldSpecLabel and returned to the calling
//		function.
//
//		None of the data values in this instance will be
//		changed or modified.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	TextFieldSpecLabel
//
//		If this method completes successfully, the Text
//		Filler Field Contents extracted from the input
//		parameter, 'txtFieldFillerDto', will be returned
//		as an instance of TextFieldSpecLabel.
//
//		This returned text label will ONLY contain the
//		Text Filler Field Contents. It will NOT contain
//		the left or right margin strings.
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtFillerFieldDtoMolecule *textFillerFieldFormatDtoMolecule) getFieldContentTextLabel(
	txtFieldFillerDto *TextFillerFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecLabel,
	error) {

	if txtFillerFieldDtoMolecule.lock == nil {
		txtFillerFieldDtoMolecule.lock = new(sync.Mutex)
	}

	txtFillerFieldDtoMolecule.lock.Lock()

	defer txtFillerFieldDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fieldContentsLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtFillerFieldDtoMolecule."+
			"getFieldContentTextLabel()",
		"")

	if err != nil {

		return fieldContentsLabel, err

	}

	if txtFieldFillerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtFieldFillerDto' is a nil pointer!\n",
			ePrefix.String())

		return fieldContentsLabel, err
	}

	_,
		err = new(textFillerFieldFormatDtoAtom).
		testValidityOfTextFillerFieldFmtDto(
			txtFieldFillerDto,
			ePrefix.XCpy(
				"txtFieldFillerDto"))

	if err != nil {

		return fieldContentsLabel, err
	}

	fieldContentsText :=
		strings.Repeat(
			txtFieldFillerDto.FillerChars,
			txtFieldFillerDto.FillerCharsRepeatCount)

	fieldContentsLabel,
		err = TextFieldSpecLabel{}.NewTextLabel(
		fieldContentsText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"fieldContentsLabel<-txtFieldFillerDto"))

	return fieldContentsLabel, err

}

// textFillerFieldFormatDtoAtom - Provides helper
// methods for TextFillerFieldFormatDto.
type textFillerFieldFormatDtoAtom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextLabelFieldFormatDto and proceeds to set all the
// member variables to their zero or uninitialized
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values contained within the
//	TextFillerFieldFormatDto instance passed as input
//	parameter 'txtFieldFmtDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFmtDto				*TextFillerFieldFormatDto
//
//		A pointer to an instance of TextFillerFieldFormatDto.
//		All data values contained within this instance
//		will be deleted and reset to their zero or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtFillerFieldDtoAtom *textFillerFieldFormatDtoAtom) empty(
	txtFieldFillerDto *TextFillerFieldFormatDto) {

	if txtFillerFieldDtoAtom.lock == nil {
		txtFillerFieldDtoAtom.lock = new(sync.Mutex)
	}

	txtFillerFieldDtoAtom.lock.Lock()

	defer txtFillerFieldDtoAtom.lock.Unlock()

	if txtFieldFillerDto == nil {

		return
	}

	txtFieldFillerDto.LeftMarginStr = ""

	txtFieldFillerDto.FillerChars = ""

	txtFieldFillerDto.FillerCharsRepeatCount = 0

	txtFieldFillerDto.RightMarginStr = ""

	return
}

// equal
//
// Compares two instances of TextFillerFieldFormatDto and
// returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextFillerFieldFormatDto are
// equal, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFillerDtoOne		*TextFillerFieldFormatDto
//
//		A pointer to an instance of
//		TextFillerFieldFormatDto. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within a
//		second TextFillerFieldFormatDto instance
//		('txtFieldFillerDtoTwo') in order to determine if
//		they are equivalent.
//
//	txtFieldFillerDtoTwo		*TextFillerFieldFormatDto
//
//		A pointer to the second of two instances of
//		TextFillerFieldFormatDto. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within the
//		first TextFillerFieldFormatDto instance
//		('txtFieldFillerDtoOne') in order to determine if
//		they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all the data values within input parameters
//		'txtFieldFillerDtoOne' and 'txtFieldFillerDtoTwo'
//		are found to be equivalent in all respects, this
//		return parameter will be set to 'true'.
//
//		If the compared data values are NOT equivalent,
//		this method returns 'false'.
func (txtFillerFieldDtoAtom *textFillerFieldFormatDtoAtom) equal(
	txtFieldFillerDtoOne *TextFillerFieldFormatDto,
	txtFieldFillerDtoTwo *TextFillerFieldFormatDto) bool {

	if txtFillerFieldDtoAtom.lock == nil {
		txtFillerFieldDtoAtom.lock = new(sync.Mutex)
	}

	txtFillerFieldDtoAtom.lock.Lock()

	defer txtFillerFieldDtoAtom.lock.Unlock()

	if txtFieldFillerDtoOne == nil ||
		txtFieldFillerDtoTwo == nil {

		return false
	}

	if txtFieldFillerDtoOne.LeftMarginStr !=
		txtFieldFillerDtoTwo.LeftMarginStr {

		return false
	}

	if txtFieldFillerDtoOne.FillerChars !=
		txtFieldFillerDtoTwo.FillerChars {

		return false
	}

	if txtFieldFillerDtoOne.FillerCharsRepeatCount !=
		txtFieldFillerDtoTwo.FillerCharsRepeatCount {

		return false
	}

	if txtFieldFillerDtoOne.RightMarginStr !=
		txtFieldFillerDtoTwo.RightMarginStr {

		return false
	}

	return true
}

// testValidityOfTextFieldFmtDto
//
// Receives a pointer to an instance of
// TextFillerFieldFormatDto and performs a diagnostic
// analysis to determine if the data values contained in
// that instance are valid in all respects.
//
// If the input parameter 'txtFieldFillerDto' is
// determined to be invalid, this method will return a
// boolean flag ('isValid') of 'false'. In addition, an
// instance of type error ('err') will be returned
// configured with an appropriate error message.
//
// If the input parameter 'txtFieldFillerDto' is valid,
// this method will return a boolean flag ('isValid') of
// 'true' and the returned error type ('err') will be set
// to 'nil'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFillerDto			*TextFillerFieldFormatDto
//
//		A pointer to an instance of
//		TextFillerFieldFormatDto.
//
//		The data values contained in this instance will
//		be reviewed and analyzed to determine if they
//		are valid in all respects.
//
//		None of the data values in this instance will be
//		changed or modified.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	isValid						bool
//
//		If all data elements contained within input
//		parameter 'txtFieldFillerDto' are judged to be
//		valid, this returned boolean value will be set to
//		'true'. If any data values are invalid, this
//		return parameter will be set to 'false'.
//
//	err							error
//
//		If this method completes successfully and all the
//		data values contained in input parameter
//		'txtFieldFillerDto' are judged to be valid, the
//		returned error Type will be set equal to 'nil'.
//
//		If the data values contained in input parameter
//		'txtFieldFillerDto' are invalid, 'error' will be
//		non-nil and configured with an appropriate error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtFillerFieldDtoAtom *textFillerFieldFormatDtoAtom) testValidityOfTextFillerFieldFmtDto(
	txtFieldFillerDto *TextFillerFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtFillerFieldDtoAtom.lock == nil {
		txtFillerFieldDtoAtom.lock = new(sync.Mutex)
	}

	txtFillerFieldDtoAtom.lock.Lock()

	defer txtFillerFieldDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFillerFieldFormatDtoAtom."+
			"testValidityOfTextFillerFieldFmtDto()",
		"")

	if err != nil {

		return isValid, err

	}

	if txtFieldFillerDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtFieldFillerDto' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if len(txtFieldFillerDto.FillerChars) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFillerFieldFormatDto parameter 'FillerChars' is INVALID!\n"+
			"txtFieldFillerDto.FillerChars is empty an has a length of zero characters.\n",
			ePrefix.String())

		return isValid, err
	}

	if txtFieldFillerDto.FillerCharsRepeatCount < 1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFillerFieldFormatDto parameter 'FillerCharsRepeatCount' is INVALID!\n"+
			"txtFieldFillerDto.FillerCharsRepeatCount has value less than one (1).\n"+
			"txtFieldFillerDto.FillerCharsRepeatCount = '%v'\n",
			ePrefix.String(),
			txtFieldFillerDto.FillerCharsRepeatCount)

		return isValid, err

	}

	isValid = true

	return isValid, err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// TextFieldFormatDtoLabel
//
// Field Format Data Transfer Objects (Dto) are used
// to facilitate easy data entry which creating and
// configuring text lines strings for screen display,
// file output or printing.
//
// Type TextFieldFormatDtoLabel contains all the format
// parameters necessary format a single text label field.
type TextFieldFormatDtoLabel struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for this Text Label Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this Text
	//	Label Field.

	FieldContents interface{}
	//	This parameter may contain one of several
	//	specific data types. This empty interface type
	//	will be converted to a string and configured as
	//	the text column content within a text line.
	//
	//	Supported types which may be submitted through
	//	this empty interface parameter are listed as
	//	follows:
	//
	//		time.Time (Converted using default format)
	//		string
	//		bool
	//		uint, uint8, uint16, uint32, uint64,
	//		int, int8, int16, int32, int64
	//		float32, float64
	//		*big.Int *big.Float
	//		fmt.Stringer (types that support this interface)
	//		TextInputParamFieldDateTimeDto
	//		       (Converts date time to string)
	//		ITextLineSpecification
	//		ITextFieldSpecification
	//		ITextFieldFormatDto
	//			TextFieldFormatDtoBigFloat
	//			TextFieldFormatDtoDate
	//			TextFieldFormatDtoLabel
	//			TextFillerFieldFormatDto
	//
	//		If the 'emptyIFace' object is not convertible to
	//		one of the supported types, an error will be returned.

	FieldLength int
	//	The length of the text field in which the
	//	'FieldContents' will be displayed. If
	//	'FieldLength' is less than the length of the
	//	'FieldContents' string, it will be automatically
	//	set equal to the 'FieldContents' string length.
	//
	//	To automatically set the value of 'FieldLength'
	//	to the length of 'FieldContents', set this
	//	parameter to a value of minus one (-1).
	//
	//	If this parameter is submitted with a value less
	//	than minus one (-1) or greater than 1-million
	//	(1,000,000), an error will be returned.
	//
	//	Field Length Examples
	//
	//		Example-1
	//          FieldContents String = "Hello World!"
	//			FieldContents String Length = 12
	//			FieldLength = 18
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   Hello World!   "
	//
	//		Example-2
	//          FieldContents = "Hello World!"
	//			FieldContents String Length = 12
	//			FieldLength = 18
	//			FieldJustify = TxtJustify.Left()
	//			Text Field String =
	//				"Hello World!      "
	//
	//		Example-3
	//          FieldContents = "Hello World!"
	//			FieldContents String Length = 12
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"Hello World!"
	//
	//		Example-4
	//          FieldContents = "Hello World!"
	//			FieldContents String Length = 12
	//			FieldLength = 2
	//			FieldJustify = TxtJustify.Center()
	//				Ignored, because FieldLength Less
	//				Than FieldContents String Length.
	//			Text Field String =
	//				"Hello World!"

	FieldJustify TextJustify
	//	An enumeration which specifies the justification
	//	of the 'FieldContents' string within the text
	//	field length specified by 'FieldLength'.
	//
	//	Text justification can only be evaluated in the
	//	context of a text label ('FieldContents'), field
	//	length ('FieldLength') and a Text Justification
	//	object of type TextJustify. This is because text
	//	labels with a field length equal to or less than
	//	the length of the text label string will never
	//	use text justification. In these cases, text
	//	justification is completely ignored.
	//
	//	If the field length is greater than the length of
	//	the text label string, text justification must be
	//	equal to one of these three valid values:
	//
	//	    TextJustify(0).Left()
	//	    TextJustify(0).Right()
	//	    TextJustify(0).Center()
	//
	//	Users can also specify the abbreviated text
	//	justification enumeration syntax as follows:
	//
	//	    TxtJustify.Left()
	//	    TxtJustify.Right()
	//	    TxtJustify.Center()
	//
	//	Text Justification Examples
	//
	//		Example-1
	//          FieldContents String = "Hello World!"
	//			FieldContents String Length = 12
	//			FieldLength = 18
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   Hello World!   "
	//
	//		Example-2
	//          FieldContents = "Hello World!"
	//			FieldContents String Length = 12
	//			FieldLength = 18
	//			FieldJustify = TxtJustify.Left()
	//			Text Field String =
	//				"Hello World!      "
	//
	//		Example-3
	//          FieldContents = "Hello World!"
	//			FieldContents String Length = 12
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"Hello World!"
	//
	//		Example-4
	//          FieldContents = "Hello World!"
	//			FieldContents String Length = 12
	//			FieldLength = 2
	//			FieldJustify = TxtJustify.Center()
	//				Ignored, because FieldLength Less
	//				Than FieldContents String Length.
	//			Text Field String =
	//				"Hello World!"

	RightMarginStr string
	//	One or more characters used to create a right
	//	margin for this Text Label Field.
	//
	//	If this parameter is set to an empty string, no
	//	right margin will be configured for this Text
	//	Label Field.

	lock *sync.Mutex
}

// CopyIn
//
// Copies all the data fields from an incoming instance
// of TextFieldFormatDtoLabel
// ('incomingTxtLabelFieldFmtDto') to the corresponding
// data fields of the current TextFieldFormatDtoLabel
// instance ('textLabelFieldFormatDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	current instance of TextFieldFormatDtoLabel
//	('textLabelFieldFormatDto').
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The original member variable data values encapsulated
//	within 'incomingTxtLabelFieldFmtDto' will remain
//	unchanged with the sole exception of
//	'incomingTxtLabelFieldFmtDto.FieldContents'.
//
//	'incomingTxtLabelFieldFmtDto.FieldContents' will be
//	converted to its equivalent string value and that
//	string value will be saved to
//	'incomingTxtLabelFieldFmtDto.FieldContents'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTxtLabelFieldFmtDto		*TextFieldFormatDtoLabel
//
//		A pointer to an instance of TextFieldFormatDtoLabel.
//
//		All the internal data field values in this
//		instance will be copied to corresponding data
//		fields of the current TextFieldFormatDtoLabel instance.
//
//		The data fields contained in
//		'incomingTxtFieldFmtDto' will NOT be changed or
//		modified.
//
//		If 'incomingTxtFieldFmtDto' contains invalid data
//		values, an error will be returned.
//
//	errorPrefix						interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) CopyIn(
	incomingTxtLabelFieldFmtDto *TextFieldFormatDtoLabel,
	errorPrefix interface{}) error {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoLabel."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textLabelFieldFormatDtoNanobot).copy(
		textLabelFieldFormatDto,
		incomingTxtLabelFieldFmtDto,
		ePrefix.XCpy(
			"textLabelFieldFormatDto<-"+
				"incomingTxtLabelFieldFmtDto"))
}

// CopyOut
//
// Returns a deep copy of the current
// TextFieldFormatDtoLabel instance.
//
// If the current TextFieldFormatDtoLabel instance
// contains invalid member variable data values, this
// method will return an error.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The original member variable data values encapsulated
//	within the current TextFieldFormatDtoLabel instance will
//	remain unchanged with the sole exception of
//	'TextFieldFormatDtoLabel.FieldContents'.
//
//	'TextFieldFormatDtoLabel.FieldContents' will be
//	converted to its equivalent string value and that
//	string value will be saved to
//	'TextFieldFormatDtoLabel.FieldContents'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
//	TextFieldFormatDtoLabel
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextFieldFormatDtoLabel instance.
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
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) CopyOut(
	errorPrefix interface{}) (
	TextFieldFormatDtoLabel,
	error) {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextLabelFieldFormatDto := TextFieldFormatDtoLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoLabel."+
			"CopyOut()",
		"")

	if err != nil {
		return newTextLabelFieldFormatDto, err
	}

	err = new(textLabelFieldFormatDtoNanobot).copy(
		&newTextLabelFieldFormatDto,
		textLabelFieldFormatDto,
		ePrefix.XCpy(
			"newTextLabelFieldFormatDto<-"+
				"textLabelFieldFormatDto"))

	return newTextLabelFieldFormatDto, err
}

// CopyOutITextFieldFormat
//
// Returns a deep copy of the current
// TextFieldFormatDtoLabel instance cast as an
// ITextFieldFormatDto interface object.
//
// If the current TextFieldFormatDtoLabel instance
// contains invalid member variable data values, this
// method will return an error.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The original member variable data values encapsulated
//	within the current TextFieldFormatDtoLabel instance will
//	remain unchanged with the sole exception of
//	'TextFieldFormatDtoLabel.FieldContents'.
//
//	'TextFieldFormatDtoLabel.FieldContents' will be
//	converted to its equivalent string value and that
//	string value will be saved to
//	'TextFieldFormatDtoLabel.FieldContents'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
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
//		TextFieldFormatDtoLabel instance cast as an
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
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) CopyOutITextFieldFormat(
	errorPrefix interface{}) (
	ITextFieldFormatDto,
	error) {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextFieldFormatDto := TextFieldFormatDtoLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoLabel."+
			"CopyOutITextFieldFormat()",
		"")

	if err != nil {
		return ITextFieldFormatDto(&newTextFieldFormatDto), err
	}

	err = new(textLabelFieldFormatDtoNanobot).copy(
		&newTextFieldFormatDto,
		textLabelFieldFormatDto,
		ePrefix.XCpy(
			"newTextFieldFormatDto<-"+
				"textLabelFieldFormatDto"))

	return ITextFieldFormatDto(&newTextFieldFormatDto), err
}

// Empty
//
// Resets all internal member variables for the current
// instance of TextFieldFormatDtoLabel to their zero or
// uninitialized states. This method will leave the
// current instance of TextFieldFormatDtoLabel in an invalid
// state and unavailable for immediate reuse.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data
// values in the current instance of TextFieldFormatDtoLabel.
// All member variable data values will be reset to their
// zero or uninitialized states.
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
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) Empty() {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	new(textLabelFieldFormatDtoAtom).empty(
		textLabelFieldFormatDto)

	textLabelFieldFormatDto.lock.Unlock()

	textLabelFieldFormatDto.lock = nil

	return
}

// Equal
//
// Receives a pointer to another instance of
// TextFieldFormatDtoLabel and proceeds to compare the
// member variables to those contained in the current
// TextFieldFormatDtoLabel instance in order to determine
// if they are equivalent.
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
//	incomingTxtLabelFieldFmtDto		*TextFieldFormatDtoLabel
//
//		A pointer to an incoming instance of
//		TextFieldFormatDtoLabel. This method will compare
//		all member variable data values in this instance
//		against those contained in the current instance
//		of TextFieldFormatDtoLabel. If the data values in
//		both instances are found to be equal in all
//		respects, this method will return a boolean value
//		of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the member variable data values contained in
//		input parameter 'incomingTxtFieldFmtDto' are
//		equal in all respects to those contained in the
//		current instance of TextFieldFormatDtoLabel, this
//		method will return a boolean value of 'true'.
//		Otherwise, a value of 'false' will be returned
//		to the calling function.
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) Equal(
	incomingTxtLabelFieldFmtDto *TextFieldFormatDtoLabel) bool {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	return new(textLabelFieldFormatDtoAtom).equal(
		textLabelFieldFormatDto,
		incomingTxtLabelFieldFmtDto)
}

// GetFieldContentTextLabel
//
// Converts the current TextFieldFormatDtoLabel instance
// member variable, 'FieldContents', to an instance of
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
// the member variable 'FieldContents'. It will NOT
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
//		Usually, it contains the name of the calling
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
//		instance of TextFieldFormatDtoLabel, will be
//		returned as text label of type
//		TextFieldSpecLabel.
//
//		This returned text label will ONLY contain the
//		Text Field Contents ('FieldContents'). It will
//		NOT contain the left and right margin strings.
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
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) GetFieldContentTextLabel(
	errorPrefix interface{}) (
	TextFieldSpecLabel,
	error) {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoLabel."+
			"GetFieldContentTextLabel()",
		"")

	if err != nil {
		return TextFieldSpecLabel{}, err
	}

	return new(textLabelFieldFormatDtoMolecule).
		getFieldContentTextLabel(
			textLabelFieldFormatDto,
			ePrefix.XCpy(
				"textLabelFieldFormatDto"))
}

// GetFieldFormatDtoType
//
// Returns a string containing the name of this type
// ('TextFieldFormatDtoLabel').
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) GetFieldFormatDtoType() string {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	return "TextFieldFormatDtoLabel"
}

// GetFormattedTextFieldStr
//
// Returns a string containing the formatted text field
// generated from the current instance of
// TextFieldFormatDtoLabel.
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
//		Usually, it contains the name of the calling
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
//		instance of TextFieldFormatDtoLabel will be
//		converted to, and returned as, a formatted text
//		field string.
//
//		The returned text field string will contain the
//		left margin, text field contents and right margin
//		as those elements are defined in the current
//		instance of TextFieldFormatDtoLabel.
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
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) GetFormattedTextFieldStr(
	errorPrefix interface{}) (
	string,
	error) {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoLabel."+
			"GetFormattedTextFieldStr()",
		"")

	if err != nil {
		return "", err
	}

	return new(textLabelFieldFormatDtoNanobot).
		getFormattedTextFieldStr(
			textLabelFieldFormatDto,
			ePrefix.XCpy(
				"textLabelFieldFormatDto"))
}

// GetLeftMarginLength
//
// Returns the length of the Left Margin String as an
// integer value.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) GetLeftMarginLength() int {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	return len(textLabelFieldFormatDto.LeftMarginStr)
}

// GetLeftMarginStr
//
// Returns the Left Margin String.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) GetLeftMarginStr() string {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	return textLabelFieldFormatDto.LeftMarginStr
}

// GetRightMarginLength
//
// Returns the length of the Right Margin String as an
// integer value.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) GetRightMarginLength() int {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	return len(textLabelFieldFormatDto.RightMarginStr)
}

// GetRightMarginStr
//
// Returns the Right Margin String.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) GetRightMarginStr() string {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	return textLabelFieldFormatDto.RightMarginStr
}

// IsValidInstance
//
// Performs a diagnostic review of the data values
// encapsulated in the current TextFieldFormatDtoLabel
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
//		instance of TextFieldFormatDtoLabel are valid,
//		this returned boolean value is set to 'true'. If
//		any data values are invalid, this return
//		parameter is set to 'false'.
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) IsValidInstance() (
	isValid bool) {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	isValid,
		_ = new(textLabelFieldFormatDtoAtom).
		testValidityOfTextLabelFieldFmtDto(
			textLabelFieldFormatDto,
			nil)

	return isValid
}

// IsValidInstanceError
//
// Performs a diagnostic review of the data values
// encapsulated in the current TextFieldFormatDtoLabel
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
//		Usually, it contains the name of the calling
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
//		TextFieldFormatDtoLabel are found to be invalid,
//		this method will return an error containing an
//		appropriate error message.
//
//		If an error message is returned, the returned
//		error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the
//		beginning of the error message.
func (textLabelFieldFormatDto *TextFieldFormatDtoLabel) IsValidInstanceError(
	errorPrefix interface{}) error {

	if textLabelFieldFormatDto.lock == nil {
		textLabelFieldFormatDto.lock = new(sync.Mutex)
	}

	textLabelFieldFormatDto.lock.Lock()

	defer textLabelFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoLabel."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textLabelFieldFormatDtoAtom).
		testValidityOfTextLabelFieldFmtDto(
			textLabelFieldFormatDto,
			ePrefix.XCpy(
				"textLabelFieldFormatDto"))

	return err
}

// textLabelFieldFormatDtoNanobot - Provides helper methods for
// TextFieldFormatDtoLabel.
type textLabelFieldFormatDtoNanobot struct {
	lock *sync.Mutex
}

// copy
//
// Copies all data from a source instance of
// TextFieldFormatDtoLabel to a destination instance of
// TextFieldFormatDtoLabel.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	TextFieldFormatDtoLabel instance passed as input
//	parameter 'destinationTxtFieldFmtDto'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The original member variable data values encapsulated
//	within 'sourceTxtFieldFmtDto' will remain unchanged
//	with the sole exception of
//	'sourceTxtFieldFmtDto.FieldContents'.
//
//	'sourceTxtFieldFmtDto.FieldContents' will be converted
//	to its equivalent string value and that string value
//	will be saved to 'sourceTxtFieldFmtDto.FieldContents'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationTxtFieldFmtDto	*TextFieldFormatDtoLabel
//
//		A pointer to an instance of TextFieldFormatDtoLabel.
//
//		Data extracted from input parameter
//		'sourceTxtFieldFmtDto' will be copied to this
//		input parameter, 'destinationTxtFieldFmtDto'.
//
//		'destinationTxtFieldFmtDto' is the destination
//		for the	copy operation.
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
//	sourceTxtFieldFmtDto		*TextFieldFormatDtoLabel
//
//		A pointer to an instance of TextFieldFormatDtoLabel.
//
//		All data values in this TextFieldFormatDtoLabel
//		instance will be copied to input parameter
//		'destinationTxtFieldFmtDto'.
//
//		'sourceTxtFieldFmtDto' is the source of the
//		copy operation.
//
//		The original member variable data values
//		encapsulated within 'sourceTxtFieldFmtDto' will
//		remain unchanged with the sole exception of
//		'sourceTxtFieldFmtDto.FieldContents'.
//
//		'sourceTxtFieldFmtDto.FieldContents' will be
//		converted to its equivalent string and that
//		string will be saved to
//		'sourceTxtFieldFmtDto.FieldContents'
//
//		If 'sourceTxtFieldFmtDto' contains invalid member
//		data variables, this method will return an error.
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
func (txtLabelFieldFmtDtoNanobot *textLabelFieldFormatDtoNanobot) copy(
	destinationTxtLabelFieldFmtDto *TextFieldFormatDtoLabel,
	sourceTxtLabelFieldFmtDto *TextFieldFormatDtoLabel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLabelFieldFmtDtoNanobot.lock == nil {
		txtLabelFieldFmtDtoNanobot.lock = new(sync.Mutex)
	}

	txtLabelFieldFmtDtoNanobot.lock.Lock()

	defer txtLabelFieldFmtDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLabelFieldFormatDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtLabelFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtLabelFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtFieldFmtDtoAtom := textLabelFieldFormatDtoAtom{}

	_,
		err = txtFieldFmtDtoAtom.
		testValidityOfTextLabelFieldFmtDto(
			sourceTxtLabelFieldFmtDto,
			ePrefix.XCpy(
				"sourceTxtLabelFieldFmtDto"))

	if err != nil {

		return err

	}

	if destinationTxtLabelFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtLabelFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtFieldFmtDtoAtom.empty(
		destinationTxtLabelFieldFmtDto)

	destinationTxtLabelFieldFmtDto.LeftMarginStr =
		sourceTxtLabelFieldFmtDto.LeftMarginStr

	var convertedStr string

	convertedStr,
		err = new(textSpecificationAtom).
		convertParamEmptyInterfaceToString(
			sourceTxtLabelFieldFmtDto.FieldContents,
			"sourceTxtLabelFieldFmtDto.FieldContents",
			ePrefix.XCpy(
				"sourceTxtLabelFieldFmtDto.FieldContents"))

	if err != nil {

		return err

	}

	destinationTxtLabelFieldFmtDto.FieldContents =
		convertedStr

	sourceTxtLabelFieldFmtDto.FieldContents =
		convertedStr

	destinationTxtLabelFieldFmtDto.FieldLength =
		sourceTxtLabelFieldFmtDto.FieldLength

	destinationTxtLabelFieldFmtDto.FieldJustify =
		sourceTxtLabelFieldFmtDto.FieldJustify

	destinationTxtLabelFieldFmtDto.RightMarginStr =
		sourceTxtLabelFieldFmtDto.RightMarginStr

	return err
}

// getFormattedTextFieldStr
//
// Converts an instance of TextFieldFormatDtoLabel to a
// formatted text field string.
//
// This formatted text field string contains the left
// margin, field contents and right margin.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLabelFieldFmtDto			*TextFieldFormatDtoLabel
//
//		A pointer to an instance of TextFieldFormatDtoLabel.
//
//		The left and right margins as well as the member
//		variable 'FieldContents' will be processed and
//		converted to a formatted text field for use in
//		screen displays, file output and printing.
//
//		If input parameter 'txtLabelFieldFmtDto' is found
//		to contain invalid data values, an error will be
//		returned
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
//		parameter, 'txtLabelFieldFmtDto', will be
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
func (txtLabelFieldFmtDtoNanobot *textLabelFieldFormatDtoNanobot) getFormattedTextFieldStr(
	txtLabelFieldFmtDto *TextFieldFormatDtoLabel,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if txtLabelFieldFmtDtoNanobot.lock == nil {
		txtLabelFieldFmtDtoNanobot.lock = new(sync.Mutex)
	}

	txtLabelFieldFmtDtoNanobot.lock.Lock()

	defer txtLabelFieldFmtDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLabelFieldFormatDtoNanobot."+
			"getFormattedTextFieldStr()",
		"")

	if err != nil {

		return "", err

	}

	if txtLabelFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtLabelFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	strBuilder := new(strings.Builder)

	if len(txtLabelFieldFmtDto.LeftMarginStr) > 0 {

		strBuilder.WriteString(txtLabelFieldFmtDto.LeftMarginStr)

	}

	var textLabel TextFieldSpecLabel

	textLabel,
		err = new(textLabelFieldFormatDtoMolecule).
		getFieldContentTextLabel(
			txtLabelFieldFmtDto,
			ePrefix.XCpy(
				"txtLabelFieldFmtDto"))

	if err != nil {

		return "", err

	}

	strBuilder.WriteString(textLabel.GetTextLabel())

	if len(txtLabelFieldFmtDto.RightMarginStr) > 0 {

		strBuilder.WriteString(txtLabelFieldFmtDto.RightMarginStr)

	}

	return strBuilder.String(), err
}

// textLabelFieldFormatDtoMolecule - Provides helper methods for
// TextFieldFormatDtoLabel.
type textLabelFieldFormatDtoMolecule struct {
	lock *sync.Mutex
}

// getFieldContentTextLabel
//
// Converts a TextFieldFormatDtoLabel instance member
// variable, 'FieldContents', to an instance of
// TextFieldSpecLabel.
//
// The TextFieldFormatDtoLabel instance is passed as
// input parameter, 'txtLabelFieldFmtDto'.
//
// The returned TextFieldSpecLabel will only contain
// the member variable 'FieldContents'. It will NOT
// contain the left and right margins.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If input parameter 'txtLabelFieldFmtDto', an instance
//	of TextFieldFormatDtoLabel, is found to be invalid,
//	an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLabelFieldFmtDto			*TextFieldFormatDtoLabel
//
//		A pointer to an instance of TextFieldFormatDtoLabel.
//
//		The member variable 'FieldContents' will be
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
//		Field Contents extracted from the input
//		parameter, 'txtLabelFieldFmtDto', will be
//		returned as an instance of TextFieldSpecLabel.
//
//		This returned text label will ONLY contain the
//		Text Field Contents. It will NOT contain the left
//		or right margin strings.
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
func (txtLabelFieldFmtDtoMolecule *textLabelFieldFormatDtoMolecule) getFieldContentTextLabel(
	txtLabelFieldFmtDto *TextFieldFormatDtoLabel,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecLabel,
	error) {

	if txtLabelFieldFmtDtoMolecule.lock == nil {
		txtLabelFieldFmtDtoMolecule.lock = new(sync.Mutex)
	}

	txtLabelFieldFmtDtoMolecule.lock.Lock()

	defer txtLabelFieldFmtDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fieldContentsLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLabelFieldFormatDtoMolecule."+
			"getFieldContentTextLabel()",
		"")

	if err != nil {

		return fieldContentsLabel, err
	}

	if txtLabelFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtLabelFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return fieldContentsLabel, err
	}

	_,
		err = new(textLabelFieldFormatDtoAtom).
		testValidityOfTextLabelFieldFmtDto(
			txtLabelFieldFmtDto,
			ePrefix.XCpy(
				"txtLabelFieldFmtDto"))

	if err != nil {

		return fieldContentsLabel, err
	}

	var fieldContentsText string

	fieldContentsText,
		err = new(textSpecificationAtom).
		convertParamEmptyInterfaceToString(
			txtLabelFieldFmtDto.FieldContents,
			"txtLabelFieldFmtDto.FieldContents",
			ePrefix.XCpy(
				"txtLabelFieldFmtDto.FieldContents"))

	if err != nil {
		return fieldContentsLabel, err
	}

	fieldContentsLabel,
		err = TextFieldSpecLabel{}.NewTextLabel(
		fieldContentsText,
		txtLabelFieldFmtDto.FieldLength,
		txtLabelFieldFmtDto.FieldJustify,
		ePrefix.XCpy(
			"fieldContentsLabel<-txtLabelFieldFmtDto"))

	return fieldContentsLabel, err
}

// textLabelFieldFormatDtoAtom - Provides helper methods for
// TextFieldFormatDtoLabel.
type textLabelFieldFormatDtoAtom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextFieldFormatDtoLabel and proceeds to set all the
// member variables to their zero or uninitialized
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values contained within the
//	TextFieldFormatDtoLabel instance passed as input
//	parameter 'txtFieldFmtDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFmtDto				*TextFieldFormatDtoLabel
//
//		A pointer to an instance of TextFieldFormatDtoLabel.
//		All data values contained within this instance
//		will be deleted and reset to their zero or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLabelFieldFmtDtoAtom *textLabelFieldFormatDtoAtom) empty(
	txtLabelFieldFmtDto *TextFieldFormatDtoLabel) {

	if txtLabelFieldFmtDtoAtom.lock == nil {
		txtLabelFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtLabelFieldFmtDtoAtom.lock.Lock()

	defer txtLabelFieldFmtDtoAtom.lock.Unlock()

	if txtLabelFieldFmtDto == nil {

		return
	}

	txtLabelFieldFmtDto.LeftMarginStr = ""

	txtLabelFieldFmtDto.FieldContents = nil

	txtLabelFieldFmtDto.FieldLength = 0

	txtLabelFieldFmtDto.FieldJustify = TxtJustify.None()

	txtLabelFieldFmtDto.RightMarginStr = ""

	return
}

// equal
//
// Compares two instances of TextFieldFormatDtoLabel and
// returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextFieldFormatDtoLabel are equal,
// this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLabelFieldFmtDtoOne			*TextFieldFormatDtoLabel
//
//		A pointer to an instance of
//		TextFieldFormatDtoLabel.
//
//		The data values contained within this instance
//		will be compared to corresponding data values
//		contained within a second TextFieldFormatDtoLabel
//		instance ('txtLabelFieldFmtDtoTwo') in order to
//		determine if they are equivalent.
//
//	txtLabelFieldFmtDtoTwo			*TextFieldFormatDtoLabel
//
//		A pointer to the second of two instances of
//		TextFieldFormatDtoLabel. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within the
//		first TextFieldFormatDtoLabel instance
//		('txtLabelFieldFmtDtoOne') in order to determine
//		if they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all the data values within input parameters
//		'txtLabelFieldFmtDtoOne' and
//		'txtLabelFieldFmtDtoTwo' are found to be
//		equivalent in all respects, this return parameter
//		will be set to 'true'.
//
//		If the compared data values are NOT equivalent,
//		this method returns 'false'.
func (txtLabelFieldFmtDtoAtom *textLabelFieldFormatDtoAtom) equal(
	txtLabelFieldFmtDtoOne *TextFieldFormatDtoLabel,
	txtLabelFieldFmtDtoTwo *TextFieldFormatDtoLabel) bool {

	if txtLabelFieldFmtDtoAtom.lock == nil {
		txtLabelFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtLabelFieldFmtDtoAtom.lock.Lock()

	defer txtLabelFieldFmtDtoAtom.lock.Unlock()

	if txtLabelFieldFmtDtoOne == nil ||
		txtLabelFieldFmtDtoTwo == nil {

		return false
	}

	if txtLabelFieldFmtDtoOne.LeftMarginStr !=
		txtLabelFieldFmtDtoTwo.LeftMarginStr {

		return false
	}

	if txtLabelFieldFmtDtoOne.FieldContents !=
		txtLabelFieldFmtDtoTwo.FieldContents {

		return false
	}

	if txtLabelFieldFmtDtoOne.FieldLength !=
		txtLabelFieldFmtDtoTwo.FieldLength {

		return false
	}

	if txtLabelFieldFmtDtoOne.FieldJustify !=
		txtLabelFieldFmtDtoTwo.FieldJustify {

		return false
	}

	if txtLabelFieldFmtDtoOne.RightMarginStr !=
		txtLabelFieldFmtDtoTwo.RightMarginStr {

		return false
	}

	return true
}

// testValidityOfTextLabelFieldFmtDto
//
// Receives a pointer to an instance of
// TextFieldFormatDtoLabel and performs a diagnostic
// analysis to determine if the data values contained in
// that instance are valid in all respects.
//
// If the input parameter 'txtFieldFmtDto' is determined
// to be invalid, this method will return a boolean flag
// ('isValid') of 'false'. In addition, an instance of
// type error ('err') will be returned configured with an
// appropriate error message.
//
// If the input parameter 'txtFieldFmtDto' is valid, this
// method will return a boolean flag ('isValid') of
// 'true' and the returned error type ('err') will be set
// to 'nil'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLabelFieldFmtDto			*TextFieldFormatDtoLabel
//
//		A pointer to an instance of TextFieldFormatDtoLabel.
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
//		parameter 'txtLabelFieldFmtDto' are judged to be
//		valid, this returned boolean value will be set to
//		'true'. If any data values are invalid, this
//		return parameter will be set to 'false'.
//
//	error
//
//		If this method completes successfully and all the
//		data values contained in input parameter
//		'txtLabelFieldFmtDto' are judged to be valid,
//		the returned error Type will be set equal to
//		'nil'.
//
//		If the data values contained in input parameter
//		'txtLabelFieldFmtDto' are invalid, the returned
//		'error' will be non-nil and configured with an
//		appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLabelFieldFmtDtoAtom *textLabelFieldFormatDtoAtom) testValidityOfTextLabelFieldFmtDto(
	txtLabelFieldFmtDto *TextFieldFormatDtoLabel,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtLabelFieldFmtDtoAtom.lock == nil {
		txtLabelFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtLabelFieldFmtDtoAtom.lock.Lock()

	defer txtLabelFieldFmtDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLabelFieldFormatDtoAtom."+
			"testValidityOfTextLabelFieldFmtDto()",
		"")

	if err != nil {

		return isValid, err

	}

	if txtLabelFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtLabelFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtLabelFieldFmtDto.FieldContents == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoLabel parameter 'FieldContents' is INVALID!\n"+
			"txtLabelFieldFmtDto.FieldContents has a value of 'nil'.\n",
			ePrefix.String())

		return isValid, err
	}

	if txtLabelFieldFmtDto.FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoLabel parameter 'FieldLength' is INVALID!\n"+
			"txtLabelFieldFmtDto.FieldLength has a value less than minus one (-1)\n"+
			"txtLabelFieldFmtDto.FieldLength = %v\n",
			ePrefix.String(),
			txtLabelFieldFmtDto.FieldLength)

		return isValid, err
	}

	if txtLabelFieldFmtDto.FieldLength > 1000000 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoLabel parameter 'FieldLength' is INVALID!\n"+
			"txtLabelFieldFmtDto.FieldLength has a value greater than one-million (1,000,000)\n"+
			"txtLabelFieldFmtDto.FieldLength = %v\n",
			ePrefix.String(),
			txtLabelFieldFmtDto.FieldLength)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

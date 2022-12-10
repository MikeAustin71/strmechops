package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// TextDateFieldFormatDto
//
// The TextDateFieldFormatDto type encapsulates input
// specifications for a text field populated with a
// formatted Date/Time string created from a type
// time.Time.
//
// Field Format Data Transfer Objects (Dto) are used
// to facilitate easy data entry which creating and
// configuring text lines strings for screen display,
// file output or printing.
type TextDateFieldFormatDto struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for this 'FieldDateTime' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this
	//	'FieldDateTime' Text Field.

	FieldDateTime time.Time
	// This time value will be used to populate a Text
	// Field used for screen display, file output or
	// printing.
	//
	// Be advised that if this value is equal to zero,
	// it constitutes an error condition.

	FieldDateTimeFormat string
	// This string will be used to format the date time
	// value contained in the 'FieldDateTime' data
	// element.
	//
	// If 'FieldDateTime' is set to a value greater than
	// zero and this 'FieldDateTimeFormat' string is
	// empty (has a zero string length), a default
	// Date/Time format string will be applied as
	// follows:
	//     "2006-01-02 15:04:05.000000000 -0700 MST"

	FieldLength int
	//	The length of the text field in which the
	//	'FieldDateTime' string will be displayed. If
	//	'FieldLength' is less than the length of the
	//	'FieldDateTime' string, it will be automatically
	//	set equal to the 'FieldDateTime' string length.
	//
	//	To automatically set the value of 'FieldLength'
	//	to the length of 'FieldDateTime', set this
	//	parameter to a value of minus one (-1).
	//
	//	If this parameter is submitted with a value less
	//	than minus one (-1) or greater than 1-million
	//	(1,000,000), an error will be returned.
	//
	//	Field Length Examples
	//
	//		Example-1
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = 45
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   2006-01-02 15:04:05.000000000 -0700 MST   "
	//
	//		Example-2
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = 45
	//			FieldJustify = TxtJustify.Left()
	//			Text Field String =
	//				"2006-01-02 15:04:05.000000000 -0700 MST      "
	//
	//		Example-3
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"2006-01-02 15:04:05.000000000 -0700 MST"

	FieldJustify TextJustify
	//	An enumeration which specifies the justification
	//	of the 'FieldDateTime' string within the text
	//	field length specified by 'FieldLength'.
	//
	//	Text justification can only be evaluated in the
	//	context of a text label ('FieldDateTime'), field
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
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = 45
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   2006-01-02 15:04:05.000000000 -0700 MST   "
	//
	//		Example-2
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = 45
	//			FieldJustify = TxtJustify.Left()
	//			Text Field String =
	//				"2006-01-02 15:04:05.000000000 -0700 MST      "
	//
	//		Example-3
	//          Date/Time String = "2006-01-02 15:04:05.000000000 -0700 MST"
	//			Date/Time String Length = 39
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"2006-01-02 15:04:05.000000000 -0700 MST"

	RightMarginStr string
	//	One or more characters used to create a right
	//	margin for this 'FieldDateTime' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	right margin will be configured for this
	//	'FieldDateTime' Text Field.

	lock *sync.Mutex
}

// CopyIn
//
// Copies all the data fields from an incoming instance
// of TextDateFieldFormatDto
// ('incomingTxtDateFieldFmtDto') to the corresponding
// data fields of the current TextDateFieldFormatDto
// instance ('textDateFieldFormatDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	current instance of TextDateFieldFormatDto
//	('textDateFieldFormatDto').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTxtDateFieldFmtDto		*TextDateFieldFormatDto
//
//		A pointer to an instance of
//		TextDateFieldFormatDto.
//
//		All the internal data field values in this
//		instance will be copied to corresponding data
//		fields of the current TextDateFieldFormatDto
//		instance.
//
//		The data fields contained in
//		'incomingTxtDateFieldFmtDto' will NOT be
//		changed or modified.
//
//		If 'incomingTxtDateFieldFmtDto' contains
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
func (textDateFieldFormatDto *TextDateFieldFormatDto) CopyIn(
	incomingTxtDateFieldFmtDto *TextDateFieldFormatDto,
	errorPrefix interface{}) error {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextDateFieldFormatDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textDateFieldFormatDtoNanobot).
		copy(textDateFieldFormatDto,
			incomingTxtDateFieldFmtDto,
			ePrefix.XCpy(
				"textDateFieldFormatDto"+
					"<-incomingTxtDateFieldFmtDto"))
}

// CopyOut
//
// Returns a deep copy of the current
// TextDateFieldFormatDto instance.
//
// If the current TextDateFieldFormatDto instance
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
//	TextDateFieldFormatDto
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextDateFieldFormatDto instance.
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
func (textDateFieldFormatDto *TextDateFieldFormatDto) CopyOut(
	errorPrefix interface{}) (
	TextDateFieldFormatDto,
	error) {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	newTxtDateFieldFormatDto := TextDateFieldFormatDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextDateFieldFormatDto."+
			"CopyOut()",
		"")

	if err != nil {
		return newTxtDateFieldFormatDto, err
	}

	err = new(textDateFieldFormatDtoNanobot).
		copy(&newTxtDateFieldFormatDto,
			textDateFieldFormatDto,
			ePrefix.XCpy(
				"newTxtDateFieldFormatDto"+
					"<-textDateFieldFormatDto"))

	return newTxtDateFieldFormatDto, err
}

// Empty
//
// Resets all internal member variables for the current
// instance of TextDateFieldFormatDto to their zero or
// uninitialized states. This method will leave the
// current instance of TextDateFieldFormatDto in an
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
// TextDateFieldFormatDto. All member variable data
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
func (textDateFieldFormatDto *TextDateFieldFormatDto) Empty() {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	new(textDateFieldFormatDtoAtom).empty(
		textDateFieldFormatDto)

	textDateFieldFormatDto.lock.Unlock()

	textDateFieldFormatDto.lock = nil

}

// Equal
//
// Receives a pointer to another instance of
// TextDateFieldFormatDto and proceeds to compare the
// member variables to those contained in the current
// TextDateFieldFormatDto instance in order to
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
//	incomingTxtDateFieldFmtDto		*TextDateFieldFormatDto
//
//		A pointer to an incoming instance of
//		TextDateFieldFormatDto. This method will
//		compare all member variable data values in this
//		instance against those contained in the current
//		instance of TextDateFieldFormatDto. If the data
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
//		input parameter 'incomingTxtDateFieldFmtDto'
//		are equal in all respects to those contained in
//		the current instance of TextDateFieldFormatDto,
//		this method will return a boolean value of
//		'true'. Otherwise, a value of 'false' will be
//		returned to the calling function.
func (textDateFieldFormatDto *TextDateFieldFormatDto) Equal(
	incomingTxtDateFieldFmtDto *TextDateFieldFormatDto) bool {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	return new(textDateFieldFormatDtoAtom).equal(
		textDateFieldFormatDto,
		incomingTxtDateFieldFmtDto)
}

// GetFieldContentTextLabel
//
// Converts the current TextDateFieldFormatDto instance
// member variable, 'FieldDateTime', to an instance of
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
// the member variable 'FieldDateTime'. It will NOT
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
//		instance of TextDateFieldFormatDto, will be
//		returned as text label of type
//		TextFieldSpecLabel.
//
//		This returned text label will ONLY contain the
//		Text Field Contents ('FieldDateTime'). It will
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
func (textDateFieldFormatDto *TextDateFieldFormatDto) GetFieldContentTextLabel(
	errorPrefix interface{}) (
	TextFieldSpecLabel,
	error) {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	textFieldContentsLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextDateFieldFormatDto."+
			"GetFieldContentTextLabel()",
		"")

	if err != nil {
		return textFieldContentsLabel, err
	}

	return new(textDateFieldFormatDtoMolecule).
		getFieldContentTextLabel(
			textDateFieldFormatDto,
			ePrefix.XCpy(
				"textDateFieldFormatDto"))
}

// GetFieldFormatDtoType
//
// Returns a string containing the name of this type
// ('TextDateFieldFormatDto').
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textDateFieldFormatDto *TextDateFieldFormatDto) GetFieldFormatDtoType() string {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	return "TextDateFieldFormatDto"
}

// GetFormattedTextFieldStr
//
// Returns a string containing the formatted text field
// generated from the current instance of
// TextDateFieldFormatDto.
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
//		instance of TextDateFieldFormatDto will be
//		converted to, and returned as, a formatted text
//		field string.
//
//		The returned text field string will contain the
//		left margin, text field contents and right margin
//		as those elements are defined in the current
//		instance of TextDateFieldFormatDto.
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
func (textDateFieldFormatDto *TextDateFieldFormatDto) GetFormattedTextFieldStr(
	errorPrefix interface{}) (
	string,
	error) {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextDateFieldFormatDto."+
			"GetFormattedTextFieldStr()",
		"")

	if err != nil {
		return "", err
	}

	return new(textDateFieldFormatDtoNanobot).
		getFormattedTextFieldStr(
			textDateFieldFormatDto,
			ePrefix.XCpy(
				"textDateFieldFormatDto"))
}

// GetLeftMarginLength
//
// Returns the length of the Left Margin String as an
// integer value.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textDateFieldFormatDto *TextDateFieldFormatDto) GetLeftMarginLength() int {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	return len(textDateFieldFormatDto.LeftMarginStr)
}

// GetLeftMarginStr
//
// Returns the Left Margin String.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textDateFieldFormatDto *TextDateFieldFormatDto) GetLeftMarginStr() string {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	return textDateFieldFormatDto.LeftMarginStr
}

// GetRightMarginLength
//
// Returns the length of the Right Margin String as an
// integer value.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textDateFieldFormatDto *TextDateFieldFormatDto) GetRightMarginLength() int {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	return len(textDateFieldFormatDto.RightMarginStr)
}

// GetRightMarginStr
//
// Returns the Right Margin String.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textDateFieldFormatDto *TextDateFieldFormatDto) GetRightMarginStr() string {

	if textDateFieldFormatDto.lock == nil {
		textDateFieldFormatDto.lock = new(sync.Mutex)
	}

	textDateFieldFormatDto.lock.Lock()

	defer textDateFieldFormatDto.lock.Unlock()

	return textDateFieldFormatDto.RightMarginStr
}

// textDateFieldFormatDtoNanobot - Provides helper
// methods for TextDateFieldFormatDto.
type textDateFieldFormatDtoNanobot struct {
	lock *sync.Mutex
}

// copy
//
// Copies all data from a source instance of
// TextDateFieldFormatDto to a destination instance of
// TextDateFieldFormatDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	TextDateFieldFormatDto instance passed as input
//	parameter 'destinationTxtDateFieldDto'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationTxtDateFieldDto		*TextDateFieldFormatDto
//
//		A pointer to an instance of
//		TextDateFieldFormatDto.
//
//		Data extracted from input parameter
//		'sourceTxtDateFieldDto' will be copied to this
//		input parameter, 'destinationTxtDateFieldDto'.
//
//		'destinationTxtDateFieldDto' is the destination
//		for this copy operation.
//
//		If this method completes successfully, all member
//		data variables encapsulated in
//		'destinationTxtDateFieldDto' will be identical to
//		those contained in input parameter,
//		'sourceTxtDateFieldDto'.
//
//		Be advised that the pre-existing data fields
//		contained within input parameter
//		'destinationTxtDateFieldDto' will be deleted and
//		overwritten.
//
//	sourceTxtDateFieldDto			*TextDateFieldFormatDto
//
//		A pointer to an instance of
//		TextDateFieldFormatDto.
//
//		All data values in this TextDateFieldFormatDto
//		instance will be copied to input parameter
//		'destinationTxtDateFieldDto'.
//
//		'sourceTxtDateFieldDto' is the source of the
//		copy operation.
//
//		The original member variable data values
//		encapsulated within 'sourceTxtDateFieldDto'
//		will remain unchanged and unmodified.
//
//		If 'sourceTxtDateFieldDto' contains invalid
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
func (txtDateFieldDtoNanobot *textDateFieldFormatDtoNanobot) copy(
	destinationTxtDateFieldDto *TextDateFieldFormatDto,
	sourceTxtDateFieldDto *TextDateFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtDateFieldDtoNanobot.lock == nil {
		txtDateFieldDtoNanobot.lock = new(sync.Mutex)
	}

	txtDateFieldDtoNanobot.lock.Lock()

	defer txtDateFieldDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textDateFieldFormatDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if destinationTxtDateFieldDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtDateFieldDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTxtDateFieldDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtDateFieldDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtDateFieldFmtDtoAtom := textDateFieldFormatDtoAtom{}

	_,
		err = txtDateFieldFmtDtoAtom.
		testValidityOfTextDateFieldFormatDto(
			sourceTxtDateFieldDto,
			ePrefix.XCpy(
				"sourceTxtDateFieldDto"))

	if err != nil {

		return err
	}

	txtDateFieldFmtDtoAtom.empty(
		destinationTxtDateFieldDto)

	destinationTxtDateFieldDto.LeftMarginStr =
		sourceTxtDateFieldDto.LeftMarginStr

	destinationTxtDateFieldDto.FieldDateTime =
		sourceTxtDateFieldDto.FieldDateTime

	destinationTxtDateFieldDto.FieldDateTimeFormat =
		sourceTxtDateFieldDto.FieldDateTimeFormat

	destinationTxtDateFieldDto.FieldLength =
		sourceTxtDateFieldDto.FieldLength

	destinationTxtDateFieldDto.FieldJustify =
		sourceTxtDateFieldDto.FieldJustify

	destinationTxtDateFieldDto.RightMarginStr =
		sourceTxtDateFieldDto.RightMarginStr

	return err
}

// getFormattedTextFieldStr
//
// Converts an instance of TextDateFieldFormatDto to a
// formatted text field string.
//
// This formatted text field string contains the left
// margin, field contents and right margin.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtDateFieldDto				*TextDateFieldFormatDto
//
//		A pointer to an instance of
//		TextDateFieldFormatDto.
//
//		The left and right margins as well as the member
//		variable 'FieldDateTime' will be processed and
//		converted to a formatted text field for use in
//		screen displays, file output and printing.
//
//		If input parameter 'txtDateFieldDto' is found to
//		contain invalid data values, an error will be
//		returned.
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
//		parameter, 'txtDateFieldDto', will be converted
//		to, and returned as, a formatted string	of text.
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
func (txtDateFieldDtoNanobot *textDateFieldFormatDtoNanobot) getFormattedTextFieldStr(
	txtDateFieldDto *TextDateFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if txtDateFieldDtoNanobot.lock == nil {
		txtDateFieldDtoNanobot.lock = new(sync.Mutex)
	}

	txtDateFieldDtoNanobot.lock.Lock()

	defer txtDateFieldDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textDateFieldFormatDtoNanobot."+
			"getFormattedTextFieldStr()",
		"")

	if err != nil {

		return "", err

	}

	if txtDateFieldDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtDateFieldDto' is a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	strBuilder := new(strings.Builder)

	if len(txtDateFieldDto.LeftMarginStr) > 0 {

		strBuilder.WriteString(txtDateFieldDto.LeftMarginStr)

	}

	var textLabel TextFieldSpecLabel

	textLabel,
		err = new(textDateFieldFormatDtoMolecule).
		getFieldContentTextLabel(
			txtDateFieldDto,
			ePrefix.XCpy(
				"txtDateFieldDto"))

	if err != nil {

		return "", err

	}

	strBuilder.WriteString(textLabel.GetTextLabel())

	if len(txtDateFieldDto.RightMarginStr) > 0 {

		strBuilder.WriteString(txtDateFieldDto.RightMarginStr)

	}

	return strBuilder.String(), err
}

// textDateFieldFormatDtoMolecule - Provides helper
// methods for TextDateFieldFormatDto.
type textDateFieldFormatDtoMolecule struct {
	lock *sync.Mutex
}

// getFieldContentTextLabel
//
// Converts a TextDateFieldFormatDto instance member
// variable, 'FieldDateTime', to an instance of
// TextFieldSpecLabel.
//
// The TextDateFieldFormatDto instance is passed as
// input parameter, 'txtDateFieldDto'.
//
// The returned TextFieldSpecLabel will only contain
// the member variable 'FieldDateTime'. It will NOT
// contain the left and right margins.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If input parameter 'txtDateFieldDto', an instance
//	of TextDateFieldFormatDto, is found to be invalid,
//	an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtDateFieldDto				*TextDateFieldFormatDto
//
//		A pointer to an instance of TextDateFieldFormatDto.
//
//		The member variable 'FieldDateTime' will be
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
//		parameter, 'txtDateFieldDto', will be
//		returned as an instance of TextFieldSpecLabel.
//
//		This returned text label will ONLY contain the
//		Text Field Date/Time String ('FieldDateTime').
//		It will NOT contain the left or right margin
//		strings.
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
func (txtDateFieldFmtDtoMolecule *textDateFieldFormatDtoMolecule) getFieldContentTextLabel(
	txtDateFieldDto *TextDateFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecLabel,
	error) {

	if txtDateFieldFmtDtoMolecule.lock == nil {
		txtDateFieldFmtDtoMolecule.lock = new(sync.Mutex)
	}

	txtDateFieldFmtDtoMolecule.lock.Lock()

	defer txtDateFieldFmtDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fieldContentsLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textDateFieldFormatDtoMolecule."+
			"getFieldContentTextLabel()",
		"")

	if err != nil {

		return fieldContentsLabel, err

	}

	if txtDateFieldDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtDateFieldDto' is a nil pointer!\n",
			ePrefix.String())

		return fieldContentsLabel, err
	}

	_,
		err = new(textDateFieldFormatDtoAtom).
		testValidityOfTextDateFieldFormatDto(
			txtDateFieldDto,
			ePrefix.XCpy(
				"txtDateFieldDto"))

	if err != nil {

		return fieldContentsLabel, err
	}

	// If txtDateFieldDto.FieldDateTimeFormat was
	// empty, default format string was applied

	fieldContentsText := txtDateFieldDto.
		FieldDateTime.Format(txtDateFieldDto.FieldDateTimeFormat)

	fieldContentsLabel,
		err = TextFieldSpecLabel{}.NewTextLabel(
		fieldContentsText,
		txtDateFieldDto.FieldLength,
		txtDateFieldDto.FieldJustify,
		ePrefix.XCpy(
			"fieldContentsLabel<-txtDateFieldDto"))

	return fieldContentsLabel, err
}

// textDateFieldFormatDtoAtom - Provides helper
// methods for TextDateFieldFormatDto.
type textDateFieldFormatDtoAtom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextDateFieldFormatDto and proceeds to set all the
// member variables to their zero or uninitialized
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values contained within the
//	TextDateFieldFormatDto instance passed as input
//	parameter 'txtDateFieldDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtDateFieldDto				*TextDateFieldFormatDto
//
//		A pointer to an instance of TextDateFieldFormatDto.
//		All data values contained within this instance
//		will be deleted and reset to their zero or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtDateFieldDtoAtom *textDateFieldFormatDtoAtom) empty(
	txtDateFieldDto *TextDateFieldFormatDto) {

	if txtDateFieldDtoAtom.lock == nil {
		txtDateFieldDtoAtom.lock = new(sync.Mutex)
	}

	txtDateFieldDtoAtom.lock.Lock()

	defer txtDateFieldDtoAtom.lock.Unlock()

	if txtDateFieldDto == nil {

		return
	}

	txtDateFieldDto.LeftMarginStr = ""

	txtDateFieldDto.FieldDateTime = time.Time{}

	txtDateFieldDto.FieldDateTimeFormat = ""

	txtDateFieldDto.FieldLength = 0

	txtDateFieldDto.FieldJustify = TxtJustify.None()

	txtDateFieldDto.RightMarginStr = ""

	return
}

// equal
//
// Compares two instances of TextDateFieldFormatDto and
// returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextDateFieldFormatDto are
// equal, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtDateFieldDtoOne			*TextDateFieldFormatDto
//
//		A pointer to an instance of
//		TextDateFieldFormatDto. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within a
//		second TextDateFieldFormatDto instance
//		('txtDateFieldDtoTwo') in order to determine if
//		they are equivalent.
//
//	txtDateFieldDtoTwo			*TextDateFieldFormatDto
//
//		A pointer to the second of two instances of
//		TextDateFieldFormatDto. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within the
//		first TextDateFieldFormatDto instance
//		('txtDateFieldDtoOne') in order to determine if
//		they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all the data values within input parameters
//		'txtDateFieldDtoOne' and 'txtDateFieldDtoTwo'
//		are found to be equivalent in all respects, this
//		return parameter will be set to 'true'.
//
//		If the compared data values are NOT equivalent,
//		this method returns 'false'.
func (txtDateFieldDtoAtom *textDateFieldFormatDtoAtom) equal(
	txtDateFieldDtoOne *TextDateFieldFormatDto,
	txtDateFieldDtoTwo *TextDateFieldFormatDto) bool {

	if txtDateFieldDtoAtom.lock == nil {
		txtDateFieldDtoAtom.lock = new(sync.Mutex)
	}

	txtDateFieldDtoAtom.lock.Lock()

	defer txtDateFieldDtoAtom.lock.Unlock()

	if txtDateFieldDtoOne == nil ||
		txtDateFieldDtoTwo == nil {

		return false
	}

	if txtDateFieldDtoOne.LeftMarginStr !=
		txtDateFieldDtoTwo.LeftMarginStr {

		return false
	}

	if !txtDateFieldDtoOne.FieldDateTime.
		Equal(txtDateFieldDtoTwo.FieldDateTime) {

		return false
	}

	if txtDateFieldDtoOne.FieldDateTimeFormat !=
		txtDateFieldDtoTwo.FieldDateTimeFormat {

		return false
	}

	if txtDateFieldDtoOne.FieldLength !=
		txtDateFieldDtoTwo.FieldLength {

		return false
	}

	if txtDateFieldDtoOne.FieldJustify !=
		txtDateFieldDtoTwo.FieldJustify {

		return false
	}

	if txtDateFieldDtoOne.RightMarginStr !=
		txtDateFieldDtoTwo.RightMarginStr {

		return false
	}

	return true
}

// testValidityOfTextDateFieldFormatDto
//
// Receives a pointer to an instance of
// TextDateFieldFormatDto and performs a diagnostic
// analysis to determine if the data values contained in
// that instance are valid in all respects.
//
// If the input parameter 'txtDateFieldDto' is determined
// to be invalid, this method will return a boolean flag
// ('isValid') of 'false'. In addition, an instance of
// type error ('err') will be returned configured with an
// appropriate error message.
//
// If the input parameter 'txtDateFieldDto' is valid,
// this method will return a boolean flag ('isValid') of
// 'true' and the returned error type ('err') will be set
// to 'nil'.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// If the 'FieldDateTimeFormat' string contained in input
// parameter 'txtDateFieldDto' is empty or has a zero
// string length, this method will automatically default
// this value to the default date time format string of:
//
//	txtDateFieldDto.FieldDateTimeFormat =
//		"2006-01-02 15:04:05.000000000 -0700 MST"
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtDateFieldDto				*TextDateFieldFormatDto
//
//		A pointer to an instance of TextDateFieldFormatDto.
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
//		parameter 'txtDateFieldDto' are judged to be
//		valid, this returned boolean value will be set to
//		'true'. If any data values are invalid, this
//		return parameter will be set to 'false'.
//
//	err							error
//
//		If this method completes successfully and all the
//		data values contained in input parameter
//		'txtDateFieldDto' are judged to be valid, the
//		returned error Type will be set equal to 'nil'.
//
//		If the data values contained in input parameter
//		'txtDateFieldDto' are invalid, the returned
//		'error' will be non-nil and configured with an
//		appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtDateFieldDtoAtom *textDateFieldFormatDtoAtom) testValidityOfTextDateFieldFormatDto(
	txtDateFieldDto *TextDateFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtDateFieldDtoAtom.lock == nil {
		txtDateFieldDtoAtom.lock = new(sync.Mutex)
	}

	txtDateFieldDtoAtom.lock.Lock()

	defer txtDateFieldDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textDateFieldFormatDtoAtom."+
			"testValidityOfTextDateFieldFormatDto()",
		"")

	if err != nil {

		return isValid, err

	}

	if txtDateFieldDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtDateFieldDto' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtDateFieldDto.FieldDateTime.IsZero() {

		err = fmt.Errorf("%v\n"+
			"Error: TextDateFieldFormatDto parameter 'FieldDateTime' is INVALID!\n"+
			"txtDateFieldDto.FieldDateTime has a value of zero (0).\n",
			ePrefix.String())

		return isValid, err

	}

	if len(txtDateFieldDto.FieldDateTimeFormat) == 0 {

		// Default = "2006-01-02 15:04:05.000000000 -0700 MST"
		txtDateFieldDto.FieldDateTimeFormat =
			new(textSpecificationMolecule).getDefaultDateTimeFormat()

	}

	if txtDateFieldDto.FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: TextDateFieldFormatDto parameter 'FieldLength' is INVALID!\n"+
			"txtDateFieldDto.FieldLength has a value less than minus one (-1)\n"+
			"txtDateFieldDto.FieldLength = %v\n",
			ePrefix.String(),
			txtDateFieldDto.FieldLength)

		return isValid, err
	}

	if txtDateFieldDto.FieldLength > 1000000 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextDateFieldFormatDto parameter 'FieldLength' is INVALID!\n"+
			"txtDateFieldDto.FieldLength has a value greater than one-million (1,000,000)\n"+
			"txtDateFieldDto.FieldLength = %v\n",
			ePrefix.String(),
			txtDateFieldDto.FieldLength)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

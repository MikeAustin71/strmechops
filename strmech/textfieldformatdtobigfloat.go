package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"sync"
)

// TextFieldFormatDtoBigFloat
//
// The TextFieldFormatDtoBigFloat type encapsulates input
// specifications for a text field populated with a
// big.Float floating point value formatted as a number
// string.
//
// Field Format Data Transfer Objects (Dto) are used
// to facilitate easy data entry which creating and
// configuring text lines strings for screen display,
// file output or printing.
//
// This type implements the ITextFieldFormatDto
// interface.
type TextFieldFormatDtoBigFloat struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for the 'BigFloatNum' Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this
	//	'BigFloatNum' Text Field.

	BigFloatNum big.Float
	// The big.Float floating point number to
	// be formatted for output as a text string.

	RoundingMode big.RoundingMode
	// The rounding mode used to round 'BigFloatNum'
	// to the number of fractional digits specified
	// by parameter, 'NumOfFractionalDigits'.
	//
	// Rounding Modes are defined in Golang as follows:
	//
	//	ToNearestEven RoundingMode == IEEE 754-2008 roundTiesToEven
	//	ToNearestAway == IEEE 754-2008 roundTiesToAway
	//	ToZero        == IEEE 754-2008 roundTowardZero
	//	AwayFromZero  == no IEEE 754-2008 equivalent
	//	ToNegativeInf == IEEE 754-2008 roundTowardNegative
	//	ToPositiveInf == IEEE 754-2008 roundTowardPositive

	NumOfFractionalDigits int
	// The number of digits to the right of the radix
	// point (a.k.a. decimal point) which will be
	// displayed in the formatted text string for the
	// big.Float floating point number, 'BigFloatNum'.
	//
	// If this value is set to minus one (-1), all
	// available fractional digits to the right of the
	// decimal point will be displayed

	FieldLength int
	//	The length of the text field in which the
	//	'BigFloatNum' string will be displayed. If
	//	'FieldLength' is less than the length of the
	//	'BigFloatNum' string, it will be automatically
	//	set equal to the 'BigFloatNum' string length.
	//
	//	To automatically set the value of 'FieldLength'
	//	to the length of the 'BigFloatNum' string, set
	//	this parameter to a value of minus one (-1).
	//
	//	If this parameter is submitted with a value less
	//	than minus one (-1) or greater than 1-million
	//	(1,000,000), an error will be returned.
	//
	//	Field Length Examples
	//
	//		Example-1
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = 25
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   5672.12345678901234   "
	//
	//		Example-2
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = 25
	//			FieldJustify = TxtJustify.Right()
	//			Text Field String =
	//				"      5672.12345678901234"
	//
	//		Example-3
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"5672.12345678901234"

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
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = 25
	//			FieldJustify = TxtJustify.Center()
	//			Text Field String =
	//				"   5672.12345678901234   "
	//
	//		Example-2
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = 25
	//			FieldJustify = TxtJustify.Right()
	//			Text Field String =
	//				"      5672.12345678901234"
	//
	//		Example-3
	//          BigFloatNum String = "5672.12345678901234"
	//			BigFloatNum String Length = 19
	//			FieldLength = -1
	//			FieldJustify = TxtJustify.Center() // Ignored
	//			Text Field String =
	//				"5672.12345678901234"

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
// of TextFieldFormatDtoBigFloat
// ('incomingTxtBigFloatFieldFmtDto') to the corresponding
// data fields of the current TextFieldFormatDtoBigFloat
// instance ('textBigFloatFieldFmtDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	current instance of TextFieldFormatDtoBigFloat
//	('textBigFloatFieldFmtDto').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTxtBigFloatFieldFmtDto		*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of
//		TextFieldFormatDtoBigFloat.
//
//		All the internal data field values in this
//		instance will be copied to corresponding data
//		fields of the current TextFieldFormatDtoBigFloat
//		instance.
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
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) CopyIn(
	incomingTxtBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat,
	errorPrefix interface{}) error {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoBigFloat."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textBigFloatFieldFormatDtoNanobot).copy(
		textBigFloatFieldFmtDto,
		incomingTxtBigFloatFieldFmtDto,
		ePrefix.XCpy(
			"textBigFloatFieldFmtDto<-"+
				"incomingTxtBigFloatFieldFmtDto"))
}

// CopyOut
//
// Returns a deep copy of the current
// TextFieldFormatDtoBigFloat instance.
//
// If the current TextFieldFormatDtoBigFloat instance
// contains invalid member variable data values, this
// method will return an error.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The original member variable data values encapsulated
//	within the current TextFieldFormatDtoBigFloat
//	instance will NOT BE changed or modified.
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
//	TextFieldFormatDtoBigFloat
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextFieldFormatDtoBigFloat instance.
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
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) CopyOut(
	errorPrefix interface{}) (
	TextFieldFormatDtoBigFloat,
	error) {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtBigFloatFieldFmtDto :=
		TextFieldFormatDtoBigFloat{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoBigFloat."+
			"CopyOut()",
		"")

	if err != nil {
		return newTxtBigFloatFieldFmtDto, err
	}

	err = new(textBigFloatFieldFormatDtoNanobot).copy(
		&newTxtBigFloatFieldFmtDto,
		textBigFloatFieldFmtDto,
		ePrefix.XCpy(
			"newTxtBigFloatFieldFmtDto<-"+
				"textBigFloatFieldFmtDto"))

	return newTxtBigFloatFieldFmtDto, err
}

// CopyOutITextFieldFormat
//
// Returns a deep copy of the current
// TextFieldFormatDtoBigFloat instance cast as an
// ITextFieldFormatDto interface object.
//
// If the current TextFieldFormatDtoBigFloat instance
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
//	within the current TextFieldFormatDtoBigFloat
//	instance will NOT BE changed or modified.
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
//		TextFieldFormatDtoBigFloat instance cast as an
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
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) CopyOutITextFieldFormat(
	errorPrefix interface{}) (
	ITextFieldFormatDto,
	error) {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtBigFloatFieldFmtDto :=
		TextFieldFormatDtoBigFloat{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoBigFloat."+
			"CopyOut()",
		"")

	if err != nil {
		return ITextFieldFormatDto(&newTxtBigFloatFieldFmtDto),
			err
	}

	err = new(textBigFloatFieldFormatDtoNanobot).copy(
		&newTxtBigFloatFieldFmtDto,
		textBigFloatFieldFmtDto,
		ePrefix.XCpy(
			"newTxtBigFloatFieldFmtDto<-"+
				"textBigFloatFieldFmtDto"))

	return ITextFieldFormatDto(&newTxtBigFloatFieldFmtDto),
		err
}

// Empty
//
// Resets all internal member variables for the current
// instance of TextFieldFormatDtoBigFloat to their zero
// or uninitialized states. This method will leave the
// current instance of TextFieldFormatDtoBigFloat in an
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
// TextFieldFormatDtoBigFloat. All member variable data
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
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) Empty() {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	new(textFieldFormatDtoBigFloatAtom).empty(
		textBigFloatFieldFmtDto)

	textBigFloatFieldFmtDto.lock.Unlock()

	textBigFloatFieldFmtDto.lock = nil

}

// Equal
//
// Receives a pointer to another instance of
// TextFieldFormatDtoBigFloat and proceeds to compare the
// member variables to those contained in the current
// TextFieldFormatDtoBigFloat instance in order to
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
//	incomingTxtLabelFieldFmtDto		*TextFieldFormatDtoBigFloat
//
//		A pointer to an incoming instance of
//		TextFieldFormatDtoBigFloat. This method will
//		compare all member variable data values in this
//		instance against those contained in the current
//		instance of TextFieldFormatDtoBigFloat. If the
//		data values in both instances are found to be
//		equal in all respects, this method will return a
//		boolean value of 'true'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If the member variable data values contained in
//		input parameter 'incomingBigFloatFieldFmtDto' are
//		equal in all respects to those contained in the
//		current instance of TextFieldFormatDtoBigFloat,
//		this method will return a boolean value of
//		'true'. Otherwise, a value of 'false' will be
//		returned to the calling function.
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) Equal(
	incomingBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) bool {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	return new(textFieldFormatDtoBigFloatAtom).
		equal(
			textBigFloatFieldFmtDto,
			incomingBigFloatFieldFmtDto)
}

// GetFieldContentTextLabel
//
// Converts the current TextFieldFormatDtoBigFloat
// instance member variable, 'BigFloatNum', to an
// instance of TextFieldSpecLabel.
//
// The returned TextFieldSpecLabel will only contain
// the member variable 'BigFloatNum'. It will NOT
// contain the left and right margins. In addition, the
// returned TextFieldSpecLabel will format the
// 'BigFloatNum' numeric value as a pure number string
// generated from a native number to string conversion.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// The returned TextFieldSpecLabel will only contain
// the member variable 'BigFloatNum'. It will NOT
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
//		instance of TextFieldFormatDtoBigFloat, will be
//		returned as text label of type
//		TextFieldSpecLabel.
//
//		This returned text label will ONLY contain the
//		Text Field Contents ('BigFloatNum'). It will
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
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) GetFieldContentTextLabel(
	errorPrefix interface{}) (
	TextFieldSpecLabel,
	error) {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoBigFloat."+
			"GetFieldContentTextLabel()",
		"")

	if err != nil {
		return TextFieldSpecLabel{}, err
	}

	return new(textBigFloatFieldFormatDtoMolecule).
		getFieldContentTextLabel(
			textBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"textBigFloatFieldFmtDto"))
}

// GetFieldFormatDtoType
//
// Returns a string containing the name of this type
// ('TextFieldFormatDtoBigFloat').
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) GetFieldFormatDtoType() string {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	return "TextFieldFormatDtoBigFloat"
}

// GetFormattedTextFieldStr
//
// Returns a string containing the formatted text field
// generated from the current instance of
// TextFieldFormatDtoBigFloat.
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
//		instance of TextFieldFormatDtoBigFloat will be
//		converted to, and returned as, a formatted text
//		field string.
//
//		The returned text field string will contain the
//		left margin, text field contents and right margin
//		as those elements are defined in the current
//		instance of TextFieldFormatDtoBigFloat.
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
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) GetFormattedTextFieldStr(
	errorPrefix interface{}) (
	string,
	error) {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoBigFloat."+
			"GetFormattedTextFieldStr()",
		"")

	if err != nil {
		return "", err
	}

	return new(textBigFloatFieldFormatDtoNanobot).
		getFormattedTextFieldStr(
			textBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"textBigFloatFieldFmtDto"))
}

// GetLeftMarginLength
//
// Returns the length of the Left Margin String as an
// integer value.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) GetLeftMarginLength() int {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	return len(textBigFloatFieldFmtDto.LeftMarginStr)
}

// GetLeftMarginStr
//
// Returns the Left Margin String.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) GetLeftMarginStr() string {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	return textBigFloatFieldFmtDto.LeftMarginStr
}

// GetPureNumberStr
//
// Returns a pure number string representing the floating
// point numeric value specified by the current instance
// of TextFieldFormatDtoBigFloat.
//
// The floating point pure number string returned by
// this method will:
//
//  1. Consist entirely of numeric digit characters.
//
//  2. Separate integer and fractional digits with a
//     decimal point ('.').
//
//  3. Designate negative values with a leading minus
//     sign ('-').
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
// If the current instance of TextFieldFormatDtoBigFloat
// contains invalid data elements, an error will be
// returned.
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
//		If this method completes successfully, this
//		string parameter will return a pure number string
//		representation of the big.Float floating point
//		value specified by the current instance of
//		TextFieldFormatDtoBigFloat.
//
//		The returned floating point pure number string
//		will:
//
//		1.	Consist entirely of numeric digit characters.
//
//		2.	Separate integer and fractional digits with a
//			decimal point ('.').
//
//		3.	Designate negative values with a leading minus
//			sign ('-').
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
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) GetPureNumberStr(
	errorPrefix interface{}) (
	string,
	error) {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoFiller."+
			"GetPureNumberStr()",
		"")

	if err != nil {
		return "", err
	}

	return new(textFieldFormatDtoBigFloatElectron).
		getNativeBigFloatPureNumStr(
			textBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"textBigFloatFieldFmtDto"))
}

// GetRightMarginLength
//
// Returns the length of the Right Margin String as an
// integer value.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) GetRightMarginLength() int {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	return len(textBigFloatFieldFmtDto.RightMarginStr)
}

// GetRightMarginStr
//
// Returns the Right Margin String.
//
// This method is required in order to implement the
// ITextFieldFormatDto interface.
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) GetRightMarginStr() string {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	return textBigFloatFieldFmtDto.RightMarginStr
}

// IsValidInstance
//
// Performs a diagnostic review of the data values
// encapsulated in the current TextFieldFormatDtoBigFloat
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
//		instance of TextFieldFormatDtoBigFloat are valid,
//		this returned boolean value is set to 'true'. If
//		any data values are invalid, this return
//		parameter is set to 'false'.
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) IsValidInstance() (
	isValid bool) {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	isValid,
		_ = new(textFieldFormatDtoBigFloatAtom).
		testValidityOfTxtFieldFmtDtoBigFloat(
			textBigFloatFieldFmtDto,
			nil)

	return isValid
}

// IsValidInstanceError
//
// Performs a diagnostic review of the data values
// encapsulated in the current TextFieldFormatDtoBigFloat
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
//		TextFieldFormatDtoBigFloat are found to be
//		invalid, this method will return an error
//		containing an appropriate error message.
//
//		If an error message is returned, the returned
//		error message will incorporate the method chain
//		and text passed by input parameter, 'errorPrefix'.
//		The 'errorPrefix' text will be attached to the
//		beginning of the error message.
func (textBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) IsValidInstanceError(
	errorPrefix interface{}) error {

	if textBigFloatFieldFmtDto.lock == nil {
		textBigFloatFieldFmtDto.lock = new(sync.Mutex)
	}

	textBigFloatFieldFmtDto.lock.Lock()

	defer textBigFloatFieldFmtDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDtoBigFloat."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	_,
		err = new(textFieldFormatDtoBigFloatAtom).
		testValidityOfTxtFieldFmtDtoBigFloat(
			textBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"textBigFloatFieldFmtDto"))

	return err
}

// textBigFloatFieldFormatDtoNanobot
//
// Provides helper methods for TextFieldFormatDtoBigFloat.
type textBigFloatFieldFormatDtoNanobot struct {
	lock *sync.Mutex
}

// copy
//
// Copies all data from a source instance of
// TextFieldFormatDtoBigFloat to a destination instance of
// TextFieldFormatDtoBigFloat.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	TextFieldFormatDtoBigFloat instance passed as input
//	parameter 'destinationTxtBigFloatFieldFmtDto'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationTxtBigFloatFmtDto	*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of TextFieldFormatDtoBigFloat.
//
//		Data extracted from input parameter
//		'sourceTxtBigFloatFmtDto' will be copied to this
//		input parameter, 'destinationTxtBigFloatFmtDto'.
//
//		'destinationTxtFieldFmtDto' is the destination
//		for this copy operation.
//
//		If this method completes successfully, all member
//		data variables encapsulated in
//		'destinationTxtBigFloatFmtDto' will be identical to
//		those contained in input parameter,
//		'sourceTxtBigFloatFmtDto'.
//
//		Be advised that the pre-existing data fields
//		contained within input parameter
//		'destinationTxtFieldFmtDto' will be overwritten
//		and deleted.
//
//	sourceTxtBigFloatFmtDto			*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of
//		TextFieldFormatDtoBigFloat.
//
//		All data values in this TextFieldFormatDtoBigFloat
//		instance will be copied to input parameter
//		'destinationTxtBigFloatFmtDto'.
//
//		'sourceTxtBigFloatFmtDto' is the source of
//		the copy operation.
//
//		If 'sourceTxtBigFloatFmtDto' contains
//		invalid member data variables, an error will be
//		returned.
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
func (txtBigFloatFieldFmtDtoNanobot *textBigFloatFieldFormatDtoNanobot) copy(
	destinationTxtBigFloatFmtDto *TextFieldFormatDtoBigFloat,
	sourceTxtBigFloatFmtDto *TextFieldFormatDtoBigFloat,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtBigFloatFieldFmtDtoNanobot.lock == nil {
		txtBigFloatFieldFmtDtoNanobot.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoNanobot.lock.Lock()

	defer txtBigFloatFieldFmtDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textBigFloatFieldFormatDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err
	}

	if destinationTxtBigFloatFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtBigFloatFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTxtBigFloatFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtBigFloatFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtBigFloatFieldFmtAtom := textFieldFormatDtoBigFloatAtom{}

	_,
		err = txtBigFloatFieldFmtAtom.
		testValidityOfTxtFieldFmtDtoBigFloat(
			sourceTxtBigFloatFmtDto,
			ePrefix.XCpy(
				"sourceTxtBigFloatFmtDto"))

	if err != nil {

		return err
	}

	txtBigFloatFieldFmtAtom.empty(
		destinationTxtBigFloatFmtDto)

	destinationTxtBigFloatFmtDto.LeftMarginStr =
		sourceTxtBigFloatFmtDto.LeftMarginStr

	destinationTxtBigFloatFmtDto.BigFloatNum.
		Copy(&sourceTxtBigFloatFmtDto.BigFloatNum)

	destinationTxtBigFloatFmtDto.RoundingMode =
		sourceTxtBigFloatFmtDto.RoundingMode

	destinationTxtBigFloatFmtDto.NumOfFractionalDigits =
		sourceTxtBigFloatFmtDto.NumOfFractionalDigits

	destinationTxtBigFloatFmtDto.FieldLength =
		sourceTxtBigFloatFmtDto.FieldLength

	destinationTxtBigFloatFmtDto.FieldJustify =
		sourceTxtBigFloatFmtDto.FieldJustify

	destinationTxtBigFloatFmtDto.RightMarginStr =
		sourceTxtBigFloatFmtDto.RightMarginStr

	return err
}

// getFormattedTextFieldStr
//
// Converts an instance of TextFieldFormatDtoBigFloat to a
// formatted text field string.
//
// This formatted text field string contains the left
// margin, field contents and right margin.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of
//		TextFieldFormatDtoBigFloat.
//
//		The left and right margins as well as the member
//		variable 'BigFloatNum' will be processed and
//		converted to a formatted text field for use in
//		screen displays, file output and printing.
//
//		If input parameter 'txtBigFloatFieldFmtDto' is
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
//		parameter, 'txtBigFloatFieldFmtDto', will be
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
func (txtBigFloatFieldFmtDtoNanobot *textBigFloatFieldFormatDtoNanobot) getFormattedTextFieldStr(
	txtBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if txtBigFloatFieldFmtDtoNanobot.lock == nil {
		txtBigFloatFieldFmtDtoNanobot.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoNanobot.lock.Lock()

	defer txtBigFloatFieldFmtDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textBigFloatFieldFormatDtoNanobot."+
			"getFormattedTextFieldStr()",
		"")

	if err != nil {

		return "", err
	}

	if txtBigFloatFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtLabelFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	strBuilder := new(strings.Builder)

	if len(txtBigFloatFieldFmtDto.LeftMarginStr) > 0 {

		strBuilder.WriteString(txtBigFloatFieldFmtDto.LeftMarginStr)

	}

	var textLabel TextFieldSpecLabel

	textLabel,
		err = new(textBigFloatFieldFormatDtoMolecule).
		getFieldContentTextLabel(
			txtBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"txtBigFloatFieldFmtDto"))

	if err != nil {

		return "", err
	}

	strBuilder.WriteString(textLabel.GetTextLabel())

	if len(txtBigFloatFieldFmtDto.RightMarginStr) > 0 {

		strBuilder.WriteString(txtBigFloatFieldFmtDto.RightMarginStr)

	}

	return strBuilder.String(), err
}

// textBigFloatFieldFormatDtoMolecule - Provides helper methods for
// TextFieldFormatDtoBigFloat.
type textBigFloatFieldFormatDtoMolecule struct {
	lock *sync.Mutex
}

// getFieldContentTextLabel
//
// Converts a TextFieldFormatDtoBigFloat instance member
// variable, 'BigFloatNum', to an instance of
// TextFieldSpecLabel.
//
// The TextFieldFormatDtoBigFloat instance is passed as
// input parameter, 'txtBigFloatFieldFmtDto'.
//
// The returned TextFieldSpecLabel will only contain
// the member variable 'BigFloatNum'. It will NOT
// contain the left and right margins. The returned
// TextFieldSpecLabel will format the 'BigFloatNum'
// numeric value as a pure number string generated
// from a native number to string conversion.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If input parameter 'txtBigFloatFieldFmtDto', an
//	instance of TextFieldFormatDtoBigFloat, is found to
//	be invalid, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of TextFieldFormatDtoBigFloat.
//
//		The member variable 'BigFloatNum' will be
//		converted to a text label of type
//		TextFieldSpecLabel and returned to the calling
//		function.
//
//		None of the data values in this instance will be
//		changed or modified.
//
//		If this instance of TextFieldFormatDtoBigFloat
//		contains invalid data elements, an error will be
//		returned.
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
//		parameter, 'txtBigFloatFieldFmtDto', will be
//		returned as an instance of TextFieldSpecLabel.
//
//		This returned text label will ONLY contain the
//		Big Float numeric value ('BigFloatNum'). It will
//		NOT contain the left or right margin strings.
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
func (txtBigFloatFieldFmtDtoMolecule *textBigFloatFieldFormatDtoMolecule) getFieldContentTextLabel(
	txtBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat,
	errPrefDto *ePref.ErrPrefixDto) (
	TextFieldSpecLabel,
	error) {

	if txtBigFloatFieldFmtDtoMolecule.lock == nil {
		txtBigFloatFieldFmtDtoMolecule.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoMolecule.lock.Lock()

	defer txtBigFloatFieldFmtDtoMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	fieldContentsLabel := TextFieldSpecLabel{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textBigFloatFieldFormatDtoMolecule."+
			"getFieldContentTextLabel()",
		"")

	if err != nil {

		return fieldContentsLabel, err
	}

	if txtBigFloatFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtBigFloatFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return fieldContentsLabel, err
	}

	var pureNumStr string

	pureNumStr,
		err = new(textFieldFormatDtoBigFloatElectron).
		getNativeBigFloatPureNumStr(
			txtBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"txtBigFloatFieldFmtDto"))

	if err != nil {

		return fieldContentsLabel, err
	}

	fieldContentsLabel,
		err = TextFieldSpecLabel{}.NewTextLabel(
		pureNumStr,
		txtBigFloatFieldFmtDto.FieldLength,
		txtBigFloatFieldFmtDto.FieldJustify,
		ePrefix.XCpy(
			"fieldContentsLabel<-txtBigFloatFieldFmtDto"))

	return fieldContentsLabel, err
}

// textFieldFormatDtoBigFloatAtom - Provides helper methods for
// TextFieldFormatDtoBigFloat.
type textFieldFormatDtoBigFloatAtom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextFieldFormatDtoBigFloat and proceeds to set all the
// member variables to their zero or uninitialized
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values contained within the
//	TextFieldFormatDtoBigFloat instance passed as input
//	parameter 'txtBigFloatFieldFmtDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of TextFieldFormatDtoBigFloat.
//		All data values contained within this instance
//		will be deleted and reset to their zero or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtFieldFmtDtoBigFloatAtom *textFieldFormatDtoBigFloatAtom) empty(
	txtBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat) {

	if txtFieldFmtDtoBigFloatAtom.lock == nil {
		txtFieldFmtDtoBigFloatAtom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoBigFloatAtom.lock.Lock()

	defer txtFieldFmtDtoBigFloatAtom.lock.Unlock()

	if txtBigFloatFieldFmtDto == nil {

		return
	}

	txtBigFloatFieldFmtDto.LeftMarginStr = ""

	txtBigFloatFieldFmtDto.BigFloatNum.SetInt64(0)

	txtBigFloatFieldFmtDto.RoundingMode =
		big.ToNearestEven

	txtBigFloatFieldFmtDto.NumOfFractionalDigits = 0

	txtBigFloatFieldFmtDto.FieldLength = 0

	txtBigFloatFieldFmtDto.FieldJustify = TxtJustify.None()

	txtBigFloatFieldFmtDto.RightMarginStr = ""

	return
}

// equal
//
// Compares two instances of TextFieldFormatDtoBigFloat
// and returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextFieldFormatDtoBigFloat are
// equal, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDtoOne		*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of
//		TextFieldFormatDtoBigFloat.
//
//		The data values contained within this instance
//		will be compared to corresponding data values
//		contained within a second
//		TextFieldFormatDtoBigFloat instance
//		('txtBigFloatFieldFmtDtoTwo') in order to
//		determine if they are equivalent.
//
//	txtBigFloatFieldFmtDtoTwo		*TextFieldFormatDtoBigFloat
//
//		A pointer to the second of two instances of
//		TextFieldFormatDtoBigFloat. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within the
//		first TextFieldFormatDtoBigFloat instance
//		('txtBigFloatFieldFmtDtoOne') in order to
//		determine if they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all the data values within input parameters
//		'txtBigFloatFieldFmtDtoOne' and
//		'txtBigFloatFieldFmtDtoTwo' are found to be
//		equivalent in all respects, this return parameter
//		will be set to 'true'.
//
//		If the compared data values are NOT equivalent,
//		this method returns 'false'.
func (txtFieldFmtDtoBigFloatAtom *textFieldFormatDtoBigFloatAtom) equal(
	txtBigFloatFieldFmtDtoOne *TextFieldFormatDtoBigFloat,
	txtBigFloatFieldFmtDtoTwo *TextFieldFormatDtoBigFloat) bool {

	if txtFieldFmtDtoBigFloatAtom.lock == nil {
		txtFieldFmtDtoBigFloatAtom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoBigFloatAtom.lock.Lock()

	defer txtFieldFmtDtoBigFloatAtom.lock.Unlock()

	if txtBigFloatFieldFmtDtoOne == nil ||
		txtBigFloatFieldFmtDtoTwo == nil {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.LeftMarginStr !=
		txtBigFloatFieldFmtDtoTwo.LeftMarginStr {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.RoundingMode !=
		txtBigFloatFieldFmtDtoTwo.RoundingMode {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.NumOfFractionalDigits !=
		txtBigFloatFieldFmtDtoTwo.NumOfFractionalDigits {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.FieldLength !=
		txtBigFloatFieldFmtDtoTwo.FieldLength {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.FieldJustify !=
		txtBigFloatFieldFmtDtoTwo.FieldJustify {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.RightMarginStr !=
		txtBigFloatFieldFmtDtoTwo.RightMarginStr {

		return false
	}

	var bFloatNumStrOne, bFloatNumStrTwo string
	var tempBFloatOne, tempBFloatTwo big.Float

	tempBFloatOne.Copy(
		&txtBigFloatFieldFmtDtoOne.BigFloatNum)

	tempBFloatOne.SetMode(
		txtBigFloatFieldFmtDtoOne.RoundingMode)

	bFloatNumStrOne = tempBFloatOne.Text(
		'f',
		txtBigFloatFieldFmtDtoOne.NumOfFractionalDigits)

	tempBFloatTwo.Copy(
		&txtBigFloatFieldFmtDtoTwo.BigFloatNum)

	tempBFloatTwo.SetMode(
		txtBigFloatFieldFmtDtoTwo.RoundingMode)

	bFloatNumStrTwo = tempBFloatTwo.Text(
		'f',
		txtBigFloatFieldFmtDtoTwo.NumOfFractionalDigits)

	if bFloatNumStrOne != bFloatNumStrTwo {
		return false
	}

	return true
}

// testValidityOfTextLabelFieldFmtDto
//
// Receives a pointer to an instance of
// TextFieldFormatDtoBigFloat and performs a diagnostic
// analysis to determine if the data values contained in
// that instance are valid in all respects.
//
// If the input parameter 'txtBigFloatFieldFmtDto' is
// determined to be invalid, this method will return a
// boolean flag ('isValid') of 'false'. In addition, an
// instance of type error ('err') will be returned
// configured with an appropriate error message.
//
// If the input parameter 'txtBigFloatFieldFmtDto' is
// valid, this method will return a boolean flag
// ('isValid') of 'true' and the returned error type
// ('err') will be set to 'nil'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of
//		TextFieldFormatDtoBigFloat.
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
//		parameter 'txtBigFloatFieldFmtDto' are judged to
//		be valid, this returned boolean value will be set
//		to 'true'. If any data values are invalid, this
//		return parameter will be set to 'false'.
//
//	error
//
//		If this method completes successfully and all the
//		data values contained in input parameter
//		'txtBigFloatFieldFmtDto' are judged to be valid,
//		the returned error Type will be set equal to
//		'nil'.
//
//		If the data values contained in input parameter
//		'txtBigFloatFieldFmtDto' are invalid, the
//		returned 'error' will be non-nil and configured
//		with an appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtFieldFmtDtoBigFloatAtom *textFieldFormatDtoBigFloatAtom) testValidityOfTxtFieldFmtDtoBigFloat(
	txtBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtFieldFmtDtoBigFloatAtom.lock == nil {
		txtFieldFmtDtoBigFloatAtom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoBigFloatAtom.lock.Lock()

	defer txtFieldFmtDtoBigFloatAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldFormatDtoBigFloatAtom."+
			"testValidityOfTxtFieldFmtDtoBigFloat()",
		"")

	if err != nil {

		return isValid, err

	}

	if txtBigFloatFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtLabelFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtBigFloatFieldFmtDto.FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoBigFloat parameter 'FieldLength' is INVALID!\n"+
			"txtBigFloatFieldFmtDto.FieldLength has a value less than minus one (-1)\n"+
			"txtBigFloatFieldFmtDto.FieldLength = %v\n",
			ePrefix.String(),
			txtBigFloatFieldFmtDto.FieldLength)

		return isValid, err
	}

	if txtBigFloatFieldFmtDto.FieldLength > 1000000 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoBigFloat parameter 'FieldLength' is INVALID!\n"+
			"txtBigFloatFieldFmtDto.FieldLength has a value greater than one-million (1,000,000)\n"+
			"txtBigFloatFieldFmtDto.FieldLength = %v\n",
			ePrefix.String(),
			txtBigFloatFieldFmtDto.FieldLength)

		return isValid, err
	}

	if txtBigFloatFieldFmtDto.NumOfFractionalDigits < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDtoBigFloat parameter 'NumOfFractionalDigits' is INVALID!\n"+
			"txtBigFloatFieldFmtDto.NumOfFractionalDigits has a value less than minus one (-1)\n"+
			"txtBigFloatFieldFmtDto.NumOfFractionalDigits = %v\n",
			ePrefix.String(),
			txtBigFloatFieldFmtDto.NumOfFractionalDigits)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// textFieldFormatDtoBigFloatElectron - Provides helper
// methods for TextFieldFormatDtoBigFloat.
type textFieldFormatDtoBigFloatElectron struct {
	lock *sync.Mutex
}

// getNativeBigFloatPureNumStr
//
// Receives a pointer to an instance of
// TextFieldFormatDtoBigFloat and extracts the
// specifications necessary to format and return a
// floating, pure number string.
//
// The floating point pure number string returned by
// this method will:
//
//  1. Consist entirely of numeric digit characters.
//
//  2. Separate integer and fractional digits with a
//     decimal point ('.').
//
//  3. Designate negative values with a leading minus
//     sign ('-').
//
//  4. NOT include integer separators such as commas
//     (',') to separate integer digits by thousands.
//
//     NOT THIS: 1,000,000
//     Pure Number String: 1000000
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	Pure number strings Do NOT include integer separators
//	(i.e. commas ',') to separate integer number strings
//	into thousands.
//
//					  NOT THIS: 1,000,000
//			Pure Number String: 1000000
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextFieldFormatDtoBigFloat
//
//		A pointer to an instance of
//		TextFieldFormatDtoBigFloat.
//
//		This instance of TextFieldFormatDtoBigFloat will
//		be converted, formatted and returned as a
//		floating point pure number string.
//
//		If this instance of TextFieldFormatDtoBigFloat
//		contains invalid data elements, an error will
//		be returned.
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
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If this method completes successfully, this
//		string parameter will return a floating point
//		pure number string representation of the
//		big.Float value passed by input paramter,
//		'txtBigFloatFieldFmtDto'.
//
//		The returned floating point pure number string
//		will:
//
//		1.	Consist entirely of numeric digit characters.
//
//		2.	Separate integer and fractional digits with a
//			decimal point ('.').
//
//		3.	Designate negative values with a leading minus
//			sign ('-').
//
//		4.	NOT include integer separators such as commas
//			(',') to separate integer digits by thousands.
//
//						  NOT THIS: 1,000,000
//				Pure Number String: 1000000
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
func (txtFieldFmtDtoBigFloatElectron *textFieldFormatDtoBigFloatElectron) getNativeBigFloatPureNumStr(
	txtBigFloatFieldFmtDto *TextFieldFormatDtoBigFloat,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if txtFieldFmtDtoBigFloatElectron.lock == nil {
		txtFieldFmtDtoBigFloatElectron.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoBigFloatElectron.lock.Lock()

	defer txtFieldFmtDtoBigFloatElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldFormatDtoBigFloatElectron."+
			"getNativeBigFloatPureNumStr()",
		"")

	if err != nil {

		return "", err

	}

	if txtBigFloatFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtBigFloatFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return "", err
	}

	_,
		err = new(textFieldFormatDtoBigFloatAtom).
		testValidityOfTxtFieldFmtDtoBigFloat(
			txtBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"txtBigFloatFieldFmtDto"))

	if err != nil {

		return "", err

	}

	txtBigFloatFieldFmtDto.BigFloatNum.SetMode(
		txtBigFloatFieldFmtDto.RoundingMode)

	numStr := txtBigFloatFieldFmtDto.BigFloatNum.Text(
		'f',
		txtBigFloatFieldFmtDto.NumOfFractionalDigits)

	return numStr, err
}

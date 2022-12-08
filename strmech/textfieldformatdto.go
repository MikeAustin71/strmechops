package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextFieldFormatDto
//
// Used to specify content and format specifications for
// a Text Field formatted for screen display, file
// output and printing.
//
// This Data Transfer Object (Dto) contains all the
// format parameters necessary format a single text
// field.
type TextFieldFormatDto struct {
	LeftMarginStr string
	//	One or more characters used to create a left
	//	margin for this Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	left margin will be configured for this Text
	//	Field.

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
	//		BigFloatTextFormatDto - Formats big.Float numbers
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

	RightMarginStr string
	//	One or more characters used to create a right
	//	margin for this Text Field.
	//
	//	If this parameter is set to an empty string, no
	//	right margin will be configured for this Text
	//	Field.

	lock *sync.Mutex
}

// CopyIn
//
// Copies all the data fields from an incoming instance
// of TextFieldFormatDto ('incomingTxtFieldFmtDto') to
// the corresponding data fields of the current
// TextFieldFormatDto instance ('textFieldFormatDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	current instance of TextFieldFormatDto
//	('textFieldFormatDto').
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The original member variable data values encapsulated
//	within 'incomingTxtFieldFmtDto' will remain unchanged
//	with the sole exception of
//	'incomingTxtFieldFmtDto.FieldContents'.
//
//	'incomingTxtFieldFmtDto.FieldContents' will be
//	converted to its equivalent string value and that
//	string value will be saved to
//	'incomingTxtFieldFmtDto.FieldContents'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTxtFieldFmtDto		*TextFieldFormatDto
//
//		A pointer to an instance of TextFieldFormatDto.
//
//		All the internal data field values in this
//		instance will be copied to corresponding data
//		fields of the current TextFieldFormatDto instance.
//
//		The data fields contained in
//		'incomingTxtFieldFmtDto' will NOT be changed or
//		modified.
//
//		If 'incomingTxtFieldFmtDto' contains invalid data
//		value, an error will be returned.
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
func (textFieldFormatDto *TextFieldFormatDto) CopyIn(
	incomingTxtFieldFmtDto *TextFieldFormatDto,
	errorPrefix interface{}) error {

	if textFieldFormatDto.lock == nil {
		textFieldFormatDto.lock = new(sync.Mutex)
	}

	textFieldFormatDto.lock.Lock()

	defer textFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(textFieldFormatDtoNanobot).copy(
		textFieldFormatDto,
		incomingTxtFieldFmtDto,
		ePrefix.XCpy(
			"textFieldFormatDto<-"+
				"incomingTxtFieldFmtDto"))
}

// CopyOut
//
// Returns a deep copy of the current TextFieldFormatDto
// instance.
//
// If the current TextFieldFormatDto instance contains
// invalid member variable dat values, this method will
// return an error.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	The original member variable data values encapsulated
//	within the current TextFieldFormatDto instance will
//	remain unchanged with the sole exception of
//	'TextFieldFormatDto.FieldContents'.
//
//	'TextFieldFormatDto.FieldContents' will be
//	converted to its equivalent string value and that
//	string value will be saved to
//	'TextFieldFormatDto.FieldContents'.
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
//	TextFieldFormatDto
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextFieldFormatDto instance.
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
func (textFieldFormatDto *TextFieldFormatDto) CopyOut(
	errorPrefix interface{}) (
	TextFieldFormatDto,
	error) {

	if textFieldFormatDto.lock == nil {
		textFieldFormatDto.lock = new(sync.Mutex)
	}

	textFieldFormatDto.lock.Lock()

	defer textFieldFormatDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTextFieldFormatDto := TextFieldFormatDto{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextFieldFormatDto."+
			"CopyOut()",
		"")

	if err != nil {
		return newTextFieldFormatDto, err
	}

	err = new(textFieldFormatDtoNanobot).copy(
		&newTextFieldFormatDto,
		textFieldFormatDto,
		ePrefix.XCpy(
			"newTextFieldFormatDto<-"+
				"textFieldFormatDto"))

	return newTextFieldFormatDto, err
}

// Empty
//
// Resets all internal member variables for the current
// instance of TextFieldFormatDto to their zero or
// uninitialized states. This method will leave the
// current instance of TextFieldFormatDto in an invalid
// state and unavailable for immediate reuse.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// This method will delete all member variable data
// values in the current instance of TextFieldFormatDto.
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
func (textFieldFormatDto *TextFieldFormatDto) Empty() {

	if textFieldFormatDto.lock == nil {
		textFieldFormatDto.lock = new(sync.Mutex)
	}

	textFieldFormatDto.lock.Lock()

	new(textFieldFormatDtoAtom).empty(
		textFieldFormatDto)

	textFieldFormatDto.lock.Unlock()

	textFieldFormatDto.lock = nil

	return
}

// Equal
//
// Receives a pointer to another instance of
// TextFieldFormatDto and proceeds to compare the member
// variables to those contained in the current
// TextFieldFormatDto instance in order to determine if
// they are equivalent.
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
//	incomingTxtFieldFmtDto		*TextFieldFormatDto
//
//		A pointer to an incoming instance of
//		TextFieldFormatDto. This method will compare all
//		member variable data values in this instance
//		against those contained in the current instance
//		of TextFieldFormatDto. If the data values in both
//		instances are found to be equal in all respects,
//		this method will return a boolean value of 'true'.
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
//		current instance of TextFieldFormatDto, this
//		method will return a boolean value of 'true'.
//		Otherwise, a value of 'false' will be returned
//		to the calling function.
func (textFieldFormatDto *TextFieldFormatDto) Equal(
	incomingTxtFieldFmtDto *TextFieldFormatDto) bool {

	if textFieldFormatDto.lock == nil {
		textFieldFormatDto.lock = new(sync.Mutex)
	}

	textFieldFormatDto.lock.Lock()

	defer textFieldFormatDto.lock.Unlock()

	return new(textFieldFormatDtoAtom).equal(
		textFieldFormatDto,
		incomingTxtFieldFmtDto)
}

// textFieldFormatDtoNanobot - Provides helper methods for
// TextFieldFormatDto.
type textFieldFormatDtoNanobot struct {
	lock *sync.Mutex
}

// copy
//
// Copies all data from a source instance of
// TextFieldFormatDto to a destination instance of
// TextFieldFormatDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	TextFieldFormatDto instance passed as input
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
//	destinationTxtFieldFmtDto	*TextFieldFormatDto
//
//		A pointer to an instance of TextFieldFormatDto.
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
//	sourceTxtFieldFmtDto		*TextFieldFormatDto
//
//		A pointer to an instance of TextFieldFormatDto.
//
//		All data values in this TextFieldFormatDto
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
func (txtFieldFmtDtoNanobot *textFieldFormatDtoNanobot) copy(
	destinationTxtFieldFmtDto *TextFieldFormatDto,
	sourceTxtFieldFmtDto *TextFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtFieldFmtDtoNanobot.lock == nil {
		txtFieldFmtDtoNanobot.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoNanobot.lock.Lock()

	defer txtFieldFmtDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldFormatDtoNanobot."+
			"copy()",
		"")

	if err != nil {

		return err

	}

	if sourceTxtFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtFieldFmtDtoAtom := textFieldFormatDtoAtom{}

	_,
		err = txtFieldFmtDtoAtom.
		testValidityOfTextFieldFmtDto(
			sourceTxtFieldFmtDto,
			ePrefix.XCpy(
				"sourceTxtFieldFmtDto"))

	if err != nil {

		return err

	}

	if destinationTxtFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtFieldFmtDtoAtom.empty(
		destinationTxtFieldFmtDto)

	destinationTxtFieldFmtDto.LeftMarginStr =
		sourceTxtFieldFmtDto.LeftMarginStr

	var convertedStr string

	convertedStr,
		err = new(textSpecificationAtom).
		convertParamEmptyInterfaceToString(
			sourceTxtFieldFmtDto.FieldContents,
			"sourceTxtFieldFmtDto.FieldContents",
			ePrefix.XCpy(
				"sourceTxtFieldFmtDto.FieldContents"))

	if err != nil {

		return err

	}

	destinationTxtFieldFmtDto.FieldContents =
		convertedStr

	sourceTxtFieldFmtDto.FieldContents =
		convertedStr

	destinationTxtFieldFmtDto.FieldLength =
		sourceTxtFieldFmtDto.FieldLength

	destinationTxtFieldFmtDto.FieldJustify =
		sourceTxtFieldFmtDto.FieldJustify

	destinationTxtFieldFmtDto.RightMarginStr =
		sourceTxtFieldFmtDto.RightMarginStr

	return err
}

// textFieldFormatDtoAtom - Provides helper methods for
// TextFieldFormatDto.
type textFieldFormatDtoAtom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextFieldFormatDto and proceeds to set all the member
// variables to their zero or uninitialized states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete reset all pre-existing data
//	values contained within the TextFieldFormatDto
//	instance passed as input
//	parameter 'txtFieldFmtDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFmtDto				*TextFieldFormatDto
//
//		A pointer to an instance of TextFieldFormatDto.
//		All data values contained within this instance
//		will be deleted and reset to their zero or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtFieldFmtDtoAtom *textFieldFormatDtoAtom) empty(
	txtFieldFmtDto *TextFieldFormatDto) {

	if txtFieldFmtDtoAtom.lock == nil {
		txtFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoAtom.lock.Lock()

	defer txtFieldFmtDtoAtom.lock.Unlock()

	if txtFieldFmtDto == nil {

		return
	}

	txtFieldFmtDto.LeftMarginStr = ""

	txtFieldFmtDto.FieldContents = nil

	txtFieldFmtDto.FieldLength = 0

	txtFieldFmtDto.FieldJustify = TxtJustify.None()

	txtFieldFmtDto.RightMarginStr = ""

	return
}

// equal
//
// Compares two instances of TextFieldFormatDto and
// returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextFieldFormatDto are equal,
// this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtFieldFmtDtoOne			*TextFieldFormatDto
//
//		A pointer to an instance of TextFieldFormatDto.
//		The data values contained within this instance
//		will be compared to corresponding data values
//		contained within a second TextFieldFormatDto
//		instance ('txtFieldFmtDtoTwo') in order to
//		determine if they are equivalent.
//
//	txtFieldFmtDtoTwo			*TextFieldFormatDto
//
//		A pointer to the second of two instances of
//		TextFieldFormatDto. The data values contained
//		within this instance will be compared to
//		corresponding data values contained within the
//		first TextFieldFormatDto instance
//		('txtFieldFmtDtoOne') in order to determine if
//		they are equivalent.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If all the data values within input parameters
//		'txtFieldFmtDtoOne' and 'txtFieldFmtDtoOne' are
//		found to be equivalent in all respects, this
//		return parameter will be set to 'true'.
//
//		If the compared data values are NOT equivalent,
//		this method returns 'false'.
func (txtFieldFmtDtoAtom *textFieldFormatDtoAtom) equal(
	txtFieldFmtDtoOne *TextFieldFormatDto,
	txtFieldFmtDtoTwo *TextFieldFormatDto) bool {

	if txtFieldFmtDtoAtom.lock == nil {
		txtFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoAtom.lock.Lock()

	defer txtFieldFmtDtoAtom.lock.Unlock()

	if txtFieldFmtDtoOne == nil ||
		txtFieldFmtDtoTwo == nil {

		return false
	}

	if txtFieldFmtDtoOne.LeftMarginStr !=
		txtFieldFmtDtoTwo.LeftMarginStr {

		return false
	}

	if txtFieldFmtDtoOne.FieldContents !=
		txtFieldFmtDtoTwo.FieldContents {

		return false
	}

	if txtFieldFmtDtoOne.FieldLength !=
		txtFieldFmtDtoTwo.FieldLength {

		return false
	}

	if txtFieldFmtDtoOne.FieldJustify !=
		txtFieldFmtDtoTwo.FieldJustify {

		return false
	}

	if txtFieldFmtDtoOne.RightMarginStr !=
		txtFieldFmtDtoTwo.RightMarginStr {

		return false
	}

	return true
}

// testValidityOfTextFieldFmtDto
//
// Receives a pointer to an instance of
// TextFieldFormatDto and performs a diagnostic analysis
// to determine if the data values contained in that
// instance are valid in all respects.
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
//	txtFieldFmtDto				*TextFieldFormatDto
//
//		A pointer to an instance of TextFieldFormatDto.
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
//	error
//
//		If this method completes successfully and all the
//		data values contained in input parameter '' are judged
//		to be valid, this returned error Type is set equal to 'nil'.
//
//		If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtFieldFmtDtoAtom *textFieldFormatDtoAtom) testValidityOfTextFieldFmtDto(
	txtFieldFmtDto *TextFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtFieldFmtDtoAtom.lock == nil {
		txtFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtFieldFmtDtoAtom.lock.Lock()

	defer txtFieldFmtDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textFieldFormatDtoAtom."+
			"testValidityOfTextFieldFmtDto()",
		"")

	if err != nil {

		return isValid, err

	}

	if txtFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'txtFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtFieldFmtDto.FieldContents == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDto parameter 'FieldContents' is INVALID!\n"+
			"txtFieldFmtDto.FieldContents has a value of 'nil'.\n",
			ePrefix.String())

		return isValid, err
	}

	if txtFieldFmtDto.FieldLength < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextFieldFormatDto parameter 'FieldLength' is INVALID!\n"+
			"txtFieldFmtDto.FieldLength has a value less than minus none (-1)\n"+
			"txtFieldFmtDto.FieldLength = %v\n",
			ePrefix.String(),
			txtFieldFmtDto.FieldLength)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

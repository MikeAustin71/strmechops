package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"sync"
)

// TextBigFloatFieldFormatDto
//
// The TextBigFloatFieldFormatDto type encapsulates input
// specifications for a text field populated with a
// big.Float floating point value formatted as a number
// string.
//
// Field Format Data Transfer Objects (Dto) are used
// to facilitate easy data entry which creating and
// configuring text lines strings for screen display,
// file output or printing.
type TextBigFloatFieldFormatDto struct {
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
// of TextBigFloatFieldFormatDto
// ('incomingTxtBigFloatFieldFmtDto') to the corresponding
// data fields of the current TextBigFloatFieldFormatDto
// instance ('textBigFloatFieldFmtDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	current instance of TextBigFloatFieldFormatDto
//	('textBigFloatFieldFmtDto').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingTxtBigFloatFieldFmtDto		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of
//		TextBigFloatFieldFormatDto.
//
//		All the internal data field values in this
//		instance will be copied to corresponding data
//		fields of the current TextBigFloatFieldFormatDto
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
func (textBigFloatFieldFmtDto *TextBigFloatFieldFormatDto) CopyIn(
	incomingTxtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto,
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
		"TextFillerFieldFormatDto."+
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

// GetPureNumberStr
//
// Returns a pure number string representing the floating
// point numeric value specified by the current instance
// of TextBigFloatFieldFormatDto.
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
// If the current instance of TextBigFloatFieldFormatDto
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
//		TextBigFloatFieldFormatDto.
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
func (textBigFloatFieldFmtDto *TextBigFloatFieldFormatDto) GetPureNumberStr(
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
		"TextFillerFieldFormatDto."+
			"GetPureNumberStr()",
		"")

	if err != nil {
		return "", err
	}

	return new(textBigFloatFieldFormatDtoElectron).
		getBigFloatPureNumberStr(
			textBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"textBigFloatFieldFmtDto"))
}

// textBigFloatFieldFormatDtoNanobot
//
// Provides helper methods for TextBigFloatFieldFormatDto.
type textBigFloatFieldFormatDtoNanobot struct {
	lock *sync.Mutex
}

// copy
//
// Copies all data from a source instance of
// TextBigFloatFieldFormatDto to a destination instance of
// TextBigFloatFieldFormatDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all
//	pre-existing data values contained within the
//	TextBigFloatFieldFormatDto instance passed as input
//	parameter 'destinationTxtBigFloatFieldFmtDto'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationTxtBigFloatFieldFmtDto	*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of TextBigFloatFieldFormatDto.
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
//	sourceTxtBigFloatFieldFmtDto		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of TextBigFloatFieldFormatDto.
//
//		All data values in this TextBigFloatFieldFormatDto
//		instance will be copied to input parameter
//		'destinationTxtBigFloatFieldFmtDto'.
//
//		'sourceTxtBigFloatFieldFmtDto' is the source of
//		the copy operation.
//
//		If 'sourceTxtBigFloatFieldFmtDto' contains
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
	destinationTxtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto,
	sourceTxtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto,
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

	if sourceTxtBigFloatFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTxtBigFloatFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtBigFloatFieldFmtAtom := textBigFloatFieldFormatDtoAtom{}

	_,
		err = txtBigFloatFieldFmtAtom.
		testValidityOfTxtBigFloatFieldFmtDto(
			sourceTxtBigFloatFieldFmtDto,
			ePrefix.XCpy(
				"sourceTxtBigFloatFieldFmtDto"))

	if err != nil {

		return err
	}

	if destinationTxtBigFloatFieldFmtDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTxtBigFloatFieldFmtDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	txtBigFloatFieldFmtAtom.empty(
		destinationTxtBigFloatFieldFmtDto)

	destinationTxtBigFloatFieldFmtDto.LeftMarginStr =
		sourceTxtBigFloatFieldFmtDto.LeftMarginStr

	destinationTxtBigFloatFieldFmtDto.BigFloatNum.
		Copy(&sourceTxtBigFloatFieldFmtDto.BigFloatNum)

	destinationTxtBigFloatFieldFmtDto.RoundingMode =
		sourceTxtBigFloatFieldFmtDto.RoundingMode

	destinationTxtBigFloatFieldFmtDto.NumOfFractionalDigits =
		sourceTxtBigFloatFieldFmtDto.NumOfFractionalDigits

	destinationTxtBigFloatFieldFmtDto.FieldLength =
		sourceTxtBigFloatFieldFmtDto.FieldLength

	destinationTxtBigFloatFieldFmtDto.FieldJustify =
		sourceTxtBigFloatFieldFmtDto.FieldJustify

	destinationTxtBigFloatFieldFmtDto.RightMarginStr =
		sourceTxtBigFloatFieldFmtDto.RightMarginStr

	return err
}

// getFormattedTextFieldStr
//
// Converts an instance of TextBigFloatFieldFormatDto to a
// formatted text field string.
//
// This formatted text field string contains the left
// margin, field contents and right margin.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of
//		TextBigFloatFieldFormatDto.
//
//		The left and right margins as well as the member
//		variable 'BigFloatNum' will be processed and
//		converted to a formatted text field for use in
//		screen displays, file output and printing.
//
//		If input parameter 'txtBigFloatFieldFmtDto' is
//		found to contain invalid data values, an error
//		will be returned
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
	txtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto,
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
// TextBigFloatFieldFormatDto.
type textBigFloatFieldFormatDtoMolecule struct {
	lock *sync.Mutex
}

// getFieldContentTextLabel
//
// Converts a TextBigFloatFieldFormatDto instance member
// variable, 'BigFloatNum', to an instance of
// TextFieldSpecLabel.
//
// The TextBigFloatFieldFormatDto instance is passed as
// input parameter, 'txtBigFloatFieldFmtDto'.
//
// The returned TextFieldSpecLabel will only contain
// the member variable 'BigFloatNum'. It will NOT
// contain the left and right margins.
//
// ----------------------------------------------------------------
//
// # BE ADVISED
//
//	If input parameter 'txtBigFloatFieldFmtDto', an
//	instance of TextBigFloatFieldFormatDto, is found to
//	be invalid, an error will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of TextBigFloatFieldFormatDto.
//
//		The member variable 'BigFloatNum' will be
//		converted to a text label of type
//		TextFieldSpecLabel and returned to the calling
//		function.
//
//		None of the data values in this instance will be
//		changed or modified.
//
//		If this instance of TextBigFloatFieldFormatDto
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
	txtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto,
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
		err = new(textBigFloatFieldFormatDtoElectron).
		getBigFloatPureNumberStr(
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

// textBigFloatFieldFormatDtoAtom - Provides helper methods for
// TextBigFloatFieldFormatDto.
type textBigFloatFieldFormatDtoAtom struct {
	lock *sync.Mutex
}

// empty
//
// Receives a pointer to an instance of
// TextBigFloatFieldFormatDto and proceeds to set all the
// member variables to their zero or uninitialized
// states.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and reset all pre-existing
//	data values contained within the
//	TextBigFloatFieldFormatDto instance passed as input
//	parameter 'txtBigFloatFieldFmtDto' to their zero or
//	uninitialized states.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of TextBigFloatFieldFormatDto.
//		All data values contained within this instance
//		will be deleted and reset to their zero or
//		uninitialized states.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtBigFloatFieldFmtDtoAtom *textBigFloatFieldFormatDtoAtom) empty(
	txtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto) {

	if txtBigFloatFieldFmtDtoAtom.lock == nil {
		txtBigFloatFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoAtom.lock.Lock()

	defer txtBigFloatFieldFmtDtoAtom.lock.Unlock()

	if txtBigFloatFieldFmtDto == nil {

		return
	}

	txtBigFloatFieldFmtDto.LeftMarginStr = ""

	txtBigFloatFieldFmtDto.BigFloatNum.SetInt64(0)

	txtBigFloatFieldFmtDto.RoundingMode =
		big.ToNearestEven

	txtBigFloatFieldFmtDto.NumOfFractionalDigits = 0

	txtBigFloatFieldFmtDto.FieldLength = 0

	txtBigFloatFieldFmtDto.FieldJustify = 0

	txtBigFloatFieldFmtDto.RightMarginStr = ""

	return
}

// equal
//
// Compares two instances of TextBigFloatFieldFormatDto
// and returns a boolean value signaling whether the two
// instances are equivalent in all respects.
//
// If the two instances of TextBigFloatFieldFormatDto are
// equal, this method returns 'true'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDtoOne		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of
//		TextBigFloatFieldFormatDto.
//
//		The data values contained within this instance
//		will be compared to corresponding data values
//		contained within a second
//		TextBigFloatFieldFormatDto instance
//		('txtBigFloatFieldFmtDtoTwo') in order to
//		determine if they are equivalent.
//
//	txtBigFloatFieldFmtDtoTwo		*TextBigFloatFieldFormatDto
//
//		A pointer to the second of two instances of
//		TextBigFloatFieldFormatDto. The data values
//		contained within this instance will be compared
//		to corresponding data values contained within the
//		first TextBigFloatFieldFormatDto instance
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
func (txtBigFloatFieldFmtDtoAtom *textBigFloatFieldFormatDtoAtom) equal(
	txtBigFloatFieldFmtDtoOne *TextBigFloatFieldFormatDto,
	txtBigFloatFieldFmtDtoTwo *TextBigFloatFieldFormatDto) bool {

	if txtBigFloatFieldFmtDtoAtom.lock == nil {
		txtBigFloatFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoAtom.lock.Lock()

	defer txtBigFloatFieldFmtDtoAtom.lock.Unlock()

	if txtBigFloatFieldFmtDtoOne == nil ||
		txtBigFloatFieldFmtDtoTwo == nil {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.LeftMarginStr !=
		txtBigFloatFieldFmtDtoOne.LeftMarginStr {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.RoundingMode !=
		txtBigFloatFieldFmtDtoOne.RoundingMode {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.NumOfFractionalDigits !=
		txtBigFloatFieldFmtDtoOne.NumOfFractionalDigits {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.FieldLength !=
		txtBigFloatFieldFmtDtoOne.FieldLength {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.FieldJustify !=
		txtBigFloatFieldFmtDtoOne.FieldJustify {

		return false
	}

	if txtBigFloatFieldFmtDtoOne.RightMarginStr !=
		txtBigFloatFieldFmtDtoOne.RightMarginStr {

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
// TextBigFloatFieldFormatDto and performs a diagnostic
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
//	txtBigFloatFieldFmtDto		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of
//		TextBigFloatFieldFormatDto.
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
func (txtBigFloatFieldFmtDtoAtom *textBigFloatFieldFormatDtoAtom) testValidityOfTxtBigFloatFieldFmtDto(
	txtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtBigFloatFieldFmtDtoAtom.lock == nil {
		txtBigFloatFieldFmtDtoAtom.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoAtom.lock.Lock()

	defer txtBigFloatFieldFmtDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textBigFloatFieldFormatDtoAtom."+
			"testValidityOfTxtBigFloatFieldFmtDto()",
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
			"ERROR: TextBigFloatFieldFormatDto parameter 'FieldLength' is INVALID!\n"+
			"txtBigFloatFieldFmtDto.FieldLength has a value less than minus one (-1)\n"+
			"txtBigFloatFieldFmtDto.FieldLength = %v\n",
			ePrefix.String(),
			txtBigFloatFieldFmtDto.FieldLength)

		return isValid, err
	}

	if txtBigFloatFieldFmtDto.FieldLength > 1000000 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextBigFloatFieldFormatDto parameter 'FieldLength' is INVALID!\n"+
			"txtBigFloatFieldFmtDto.FieldLength has a value greater than one-million (1,000,000)\n"+
			"txtBigFloatFieldFmtDto.FieldLength = %v\n",
			ePrefix.String(),
			txtBigFloatFieldFmtDto.FieldLength)

		return isValid, err
	}

	if txtBigFloatFieldFmtDto.NumOfFractionalDigits < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: TextBigFloatFieldFormatDto parameter 'NumOfFractionalDigits' is INVALID!\n"+
			"txtBigFloatFieldFmtDto.NumOfFractionalDigits has a value less than minus one (-1)\n"+
			"txtBigFloatFieldFmtDto.NumOfFractionalDigits = %v\n",
			ePrefix.String(),
			txtBigFloatFieldFmtDto.NumOfFractionalDigits)

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// textBigFloatFieldFormatDtoElectron - Provides helper
// methods for TextBigFloatFieldFormatDto.
type textBigFloatFieldFormatDtoElectron struct {
	lock *sync.Mutex
}

// getBigFloatPureNumberStr
//
// Receives a pointer to an instance of
// TextBigFloatFieldFormatDto and extracts the
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
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtBigFloatFieldFmtDto		*TextBigFloatFieldFormatDto
//
//		A pointer to an instance of
//		TextBigFloatFieldFormatDto.
//
//		This instance of TextBigFloatFieldFormatDto will
//		be converted, formatted and returned as a
//		floating point pure number string.
//
//		If this instance of TextBigFloatFieldFormatDto
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
func (txtBigFloatFieldFmtDtoElectron *textBigFloatFieldFormatDtoElectron) getBigFloatPureNumberStr(
	txtBigFloatFieldFmtDto *TextBigFloatFieldFormatDto,
	errPrefDto *ePref.ErrPrefixDto) (
	string,
	error) {

	if txtBigFloatFieldFmtDtoElectron.lock == nil {
		txtBigFloatFieldFmtDtoElectron.lock = new(sync.Mutex)
	}

	txtBigFloatFieldFmtDtoElectron.lock.Lock()

	defer txtBigFloatFieldFmtDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textBigFloatFieldFormatDtoElectron."+
			"getBigFloatPureNumberStr()",
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
		err = new(textBigFloatFieldFormatDtoAtom).
		testValidityOfTxtBigFloatFieldFmtDto(
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

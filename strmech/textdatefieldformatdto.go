package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
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

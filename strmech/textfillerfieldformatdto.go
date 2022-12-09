package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

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
// /
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
//	This method will delete reset all pre-existing data
//	values contained within the TextFillerFieldFormatDto
//	instance passed as input parameter 'txtFieldFmtDto'
//	to their zero or uninitialized states.
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
//	error
//
//		If this method completes successfully and all the
//		data values contained in input parameter
//		'txtFieldFillerDto' are judged to be valid, this
//		returned error Type is set equal to 'nil'.
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

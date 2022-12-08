package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textLineSpecTitleMarqueeElectron
//
// Provides helper methods for type
// TextLineSpecTitleMarquee
type textLineSpecTitleMarqueeElectron struct {
	lock *sync.Mutex
}

//	empty
//
//	Receives a pointer to an instance of
//	TextLineSpecTitleMarquee and proceeds to reset the
//	data values for member variables to their initial
//	or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values contained in
//	input parameter 'txtLineTitleMarquee' will be
//	deleted and reset to their zero values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLineTitleMarquee			*TextLineSpecTitleMarquee
//
//		A pointer to an instance of
//		TextLineSpecTitleMarquee. All the internal member
//		variables contained in this instance will be
//		deleted and reset to their zero values.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (txtLineTitleMarqueeElectron *textLineSpecTitleMarqueeElectron) empty(
	txtLineTitleMarquee *TextLineSpecTitleMarquee) {

	if txtLineTitleMarqueeElectron.lock == nil {
		txtLineTitleMarqueeElectron.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeElectron.lock.Lock()

	defer txtLineTitleMarqueeElectron.lock.Unlock()

	if txtLineTitleMarquee == nil {
		return
	}

	txtLineTitleMarquee.standardTitleLeftMargin = ""

	txtLineTitleMarquee.standardTitleRightMargin = ""

	txtLineTitleMarquee.standardMaxLineLen = 0

	txtLineTitleMarquee.standardTextFieldLen = 0

	txtLineTitleMarquee.standardTextFieldJustification =
		TxtJustify.None()

	txtLineTitleMarquee.leadingMarqueeLines.Empty()

	txtLineTitleMarquee.titleLines.Empty()

	txtLineTitleMarquee.trailingMarqueeLines.Empty()

	txtLineTitleMarquee.textLineReader = nil

	return
}

// equal
//
// Receives a pointer to two instances of
// TextLineSpecTitleMarquee and proceeds to compare
// their member variables in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison
// is returned. If the member variables for both
// instances are equal in all respects, this flag is set
// to 'true'. Otherwise, this method returns 'false'.
func (txtLineTitleMarqueeElectron *textLineSpecTitleMarqueeElectron) equal(
	txtLineTitleOne *TextLineSpecTitleMarquee,
	txtLineTitleTwo *TextLineSpecTitleMarquee) bool {

	if txtLineTitleMarqueeElectron.lock == nil {
		txtLineTitleMarqueeElectron.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeElectron.lock.Lock()

	defer txtLineTitleMarqueeElectron.lock.Unlock()

	if txtLineTitleOne.standardTitleLeftMargin !=
		txtLineTitleTwo.standardTitleLeftMargin {

		return false
	}

	if txtLineTitleOne.standardTitleRightMargin !=
		txtLineTitleTwo.standardTitleRightMargin {

		return false
	}

	if txtLineTitleOne.standardMaxLineLen !=
		txtLineTitleTwo.standardMaxLineLen {

		return false
	}

	if txtLineTitleOne.standardTextFieldLen !=
		txtLineTitleTwo.standardTextFieldLen {

		return false
	}

	if txtLineTitleOne.standardTextFieldJustification !=
		txtLineTitleTwo.standardTextFieldJustification {

		return false
	}

	if !txtLineTitleOne.leadingMarqueeLines.Equal(
		&txtLineTitleTwo.leadingMarqueeLines) {

		return false
	}

	if !txtLineTitleOne.titleLines.Equal(
		&txtLineTitleTwo.titleLines) {

		return false
	}

	if !txtLineTitleOne.trailingMarqueeLines.Equal(
		&txtLineTitleTwo.trailingMarqueeLines) {

		return false
	}

	return true
}

//	testValidityTitleMarquee
//
//	Receives a pointer to an instance of
//	TextLineSpecTitleMarquee and performs a diagnostic
//	analysis to determine if that instance is valid in
//	all respects.
//
//	If the input parameter 'txtLineTitleMarquee' is
//	determined to be invalid, this method will return a
//	boolean flag ('isValid') of 'false'. In addition, an
//	instance of type error ('err') will be returned
//	configured with an appropriate error message.
//
//	If the input parameter 'txtLineTitleMarquee' is valid,
//	this method will return a boolean flag ('isValid') of
//	'true' and the returned error type ('err') will be
//	set to 'nil'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtLineTitleMarquee			*TextLineSpecTitleMarquee
//
//		A pointer to an instance of TextLineSpecTitleMarquee.
//		No data elements in this instance will be modified.
//
//		The internal member data elements contained in this
//		instance will be analyzed to determine if they are
//		valid in all respects.
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
//	isValid                    bool
//
//		If input parameter 'txtLineTitleMarquee' is judged
//		to be valid in all respects, this return parameter
//		will be set to 'true'.
//
//		If input parameter 'txtLineTitleMarquee' is found to
//		be invalid, this return parameter will be set to
//		'false'.
//
//	err							error
//
//		If input parameter 'TextLineSpecTitleMarquee' is
//		judged to be valid in all respects, the returned
//		error Type is set equal to 'nil'.
//
//		If input parameter 'TextLineSpecTitleMarquee' is
//		found to be invalid, the returned error Type will
//		encapsulate an appropriate error message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLineTitleMarqueeElectron *textLineSpecTitleMarqueeElectron) testValidityTitleMarquee(
	txtLineTitleMarquee *TextLineSpecTitleMarquee,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if txtLineTitleMarqueeElectron.lock == nil {
		txtLineTitleMarqueeElectron.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeElectron.lock.Lock()

	defer txtLineTitleMarqueeElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtLineTitleMarqueeElectron."+
			"testValidityTitleMarquee()",
		"")

	if err != nil {
		return isValid, err
	}

	if txtLineTitleMarquee == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineTitleMarquee' is a nil pointer!\n",
			ePrefix.String())

		return isValid, err
	}

	if txtLineTitleMarquee.standardMaxLineLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineSpecTitleMarquee object contains invalid data values!\n"+
			"'standardMaxLineLen' has NOT been properly configured.\n"+
			"'standardMaxLineLen' has a value less than one (1).\n"+
			"standardMaxLineLen = %v\n",
			ePrefix.String(),
			txtLineTitleMarquee.standardMaxLineLen)

		return isValid, err
	}

	if txtLineTitleMarquee.standardTextFieldLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineSpecTitleMarquee object contains invalid data values!\n"+
			"'standardTextFieldLen' has NOT been properly configured.\n"+
			"'standardTextFieldLen' has a value less than one (1).\n"+
			"standardTextFieldLen = %v\n",
			ePrefix.String(),
			txtLineTitleMarquee.standardTextFieldLen)

		return isValid, err
	}

	maxAvailableTextFieldLen :=
		txtLineTitleMarquee.standardMaxLineLen -
			1 -
			len(txtLineTitleMarquee.standardTitleLeftMargin) -
			len(txtLineTitleMarquee.standardTitleRightMargin)

	if maxAvailableTextFieldLen < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineSpecTitleMarquee object contains invalid data values!\n"+
			"Maximum Available Text Field Length produces a value less than one (1).\n"+
			"maxAvailableTextFieldLen = \n"+
			"	standardMaxLineLen - Left Margin Length - Right Margin Length -1\n"+
			"standardMaxLineLen is probably invalid."+
			"standardMaxLineLen  = %v\n"+
			"Left Margin Length  = %v\n"+
			"Right Margin Length = %v\n",
			ePrefix.String(),
			txtLineTitleMarquee.standardMaxLineLen,
			len(txtLineTitleMarquee.standardTitleLeftMargin),
			len(txtLineTitleMarquee.standardTitleRightMargin))

		return isValid, err
	}

	if txtLineTitleMarquee.standardTextFieldLen > maxAvailableTextFieldLen {

		txtLineTitleMarquee.standardTextFieldLen =
			maxAvailableTextFieldLen
	}

	if txtLineTitleMarquee.leadingMarqueeLines.GetNumberOfTextLines() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineSpecTitleMarquee object contains invalid data values!\n"+
			"txtLineTitleMarquee.leadingMarqueeLines is empty.\n"+
			"Configure leadingMarqueeLines before continuing.\n",
			ePrefix.String())

		return isValid, err

	}

	err = txtLineTitleMarquee.leadingMarqueeLines.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLineTitleMarquee.leadingMarqueeLines invalid."))

	if err != nil {

		return isValid, err
	}

	if txtLineTitleMarquee.titleLines.GetNumberOfTextLines() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineSpecTitleMarquee object contains invalid data values!\n"+
			"No Text Title Lines are configured for this TextLineSpecTitleMarquee object.\n"+
			"The Text Title Lines array is empty.\n",
			ePrefix.String())

		return isValid, err

	}

	err = txtLineTitleMarquee.titleLines.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLineTitleMarquee.titleLines invalid."))

	if err != nil {

		return isValid, err
	}

	if txtLineTitleMarquee.trailingMarqueeLines.GetNumberOfTextLines() == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The TextLineSpecTitleMarquee object contains invalid data values!\n"+
			"txtLineTitleMarquee.trailingMarqueeLines is empty.\n"+
			"Configure trailingMarqueeLines before continuing.\n",
			ePrefix.String())

		return isValid, err

	}

	err = txtLineTitleMarquee.trailingMarqueeLines.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLineTitleMarquee.trailingMarqueeLines invalid."))

	if err != nil {

		return isValid, err
	}

	isValid = true

	return isValid, err
}

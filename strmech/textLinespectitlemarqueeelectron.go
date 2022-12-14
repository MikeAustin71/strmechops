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

	txtLinesColAtom := textLineSpecLinesCollectionAtom{}

	txtLinesColAtom.
		emptyCollection(
			&txtLineTitleMarquee.leadingMarqueeLines)

	txtLinesColAtom.
		emptyCollection(
			&txtLineTitleMarquee.titleLines)

	txtLinesColAtom.
		emptyCollection(
			&txtLineTitleMarquee.trailingMarqueeLines)

	txtLineTitleMarquee.textLineReader = nil

	return
}

// emptyOneMarqueeLinesCollection
//
// Empties or deletes a single Marquee Lines Collection.
//
// Type TextLineSpecTitleMarquee encapsulates three types
// of text lines used in generating title marquees:
//
//  1. Leading Marquee Lines
//     Usually consists of leading blank lines
//     and solid lines.
//
//  2. Title Lines
//     Consists entirely of text strings functioning
//     as the main title lines.
//
//  3. Trailing Marquee Lines
//     Usually consists of trailing blank lines
//     and solid lines.
//
// This method will delete all Leading Marquee Lines
// member elements in one of the three collections
// detailed above. The specific collection to be
// deleted is designated by input parameter
// titleMarqueeLineType.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLineTitleMarquee			*TextLineSpecTitleMarquee
//
//		A pointer to an instance of
//		TextLineSpecTitleMarquee. All member elements in
//		one of the three Marquee Text Lines Collections
//		will be deleted form this instance of
//		TextLineSpecTitleMarquee. The specific Marquee
//	 	Text Lines Collection to deleted is designated by
//		input parameter 'titleMarqueeLineType'.
//
//	titleMarqueeLineType		TextTileLineType
//
//		Type TextTileLineType is an enumeration of
//		Title Marquee Text Line Types. This parameter
//		determines which text line collection will
//		be deleted.
//
//		If this parameter is not set to one of the
//		following valid values, an error will be
//		returned.
//
//		Formal TextTileLineType Syntax
//
//			TextTileLineType(0).LeadingMarqueeLine()
//			TextTileLineType(0).TitleLine()
//			TextTileLineType(0).TrailingMarqueeLine()
//
//		Abbreviated TextTileLineType Syntax
//
//			TitleLineType.LeadingMarqueeLine()
//			TitleLineType.TitleLine()
//			TitleLineTypeTrailingMarqueeLine()
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
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLineTitleMarqueeElectron *textLineSpecTitleMarqueeElectron) emptyOneMarqueeLinesCollection(
	txtLineTitleMarquee *TextLineSpecTitleMarquee,
	titleMarqueeLineType TextTileLineType,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLineTitleMarqueeElectron.lock == nil {
		txtLineTitleMarqueeElectron.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeElectron.lock.Lock()

	defer txtLineTitleMarqueeElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"txtLineTitleMarqueeElectron."+
			"emptyOneMarqueeLinesCollection()",
		"")

	if err != nil {
		return err
	}

	var txtLineCollection *TextLineSpecLinesCollection

	switch titleMarqueeLineType {

	case TitleLineType.LeadingMarqueeLine():

		txtLineCollection = &txtLineTitleMarquee.leadingMarqueeLines

	case TitleLineType.TitleLine():

		txtLineCollection = &txtLineTitleMarquee.titleLines

	case TitleLineType.TrailingMarqueeLine():

		txtLineCollection = &txtLineTitleMarquee.trailingMarqueeLines

	default:

		err := fmt.Errorf("%v\n"+
			"Error: Input parameter 'titleMarqueeLineType' is invalid!\n"+
			" titleMarqueeLineType string value = '%v'\n"+
			"titleMarqueeLineType integer value = '%v'\n",
			ePrefix.String(),
			titleMarqueeLineType.String(),
			titleMarqueeLineType.XValueInt())

		return err
	}

	new(textLineSpecLinesCollectionAtom).
		emptyCollection(txtLineCollection)

	return err
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

	numOfTxtLines :=
		txtLineTitleMarquee.leadingMarqueeLines.GetNumberOfTextLines()

	if numOfTxtLines > 0 {

		err = txtLineTitleMarquee.leadingMarqueeLines.IsValidInstanceError(
			ePrefix.XCpy(
				"txtLineTitleMarquee.leadingMarqueeLines"))

		if err != nil {
			return isValid, err
		}
	}

	numOfTxtLines =
		txtLineTitleMarquee.titleLines.GetNumberOfTextLines()

	if numOfTxtLines > 0 {

		err = txtLineTitleMarquee.titleLines.IsValidInstanceError(
			ePrefix.XCpy(
				"txtLineTitleMarquee.titleLines"))

		if err != nil {
			return isValid, err
		}

	}

	numOfTxtLines =
		txtLineTitleMarquee.trailingMarqueeLines.GetNumberOfTextLines()

	if numOfTxtLines > 0 {

		err = txtLineTitleMarquee.trailingMarqueeLines.IsValidInstanceError(
			ePrefix.XCpy(
				"txtLineTitleMarquee.trailingMarqueeLines"))

		if err != nil {
			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}

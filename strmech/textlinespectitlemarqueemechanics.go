package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textLineSpecTitleMarqueeMechanics
//
// Provides helper methods for type
// TextLineSpecTitleMarquee
type textLineSpecTitleMarqueeMechanics struct {
	lock *sync.Mutex
}

//	setTxtLineTitleMarquee
//
//	Receives Text Line Title configuration specifications
//	and proceeds to copy those specifications to an
//	instance of TextLineSpecTitleMarquee passed as input
//	parameter 'txtLineTitleMarquee'.
//
//	NOTE:	If the 'configSpecs' TitleLines array is empty,
//			no error will be returned.
//
//			The user is thereafter responsible for
//			configuring and adding Title Lines as needed.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete and overwrite all pre-existing
//	data values in the TextLineTitleMarqueeDto instance
//	'txtLineTitleMarquee' passed as an input parameter.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	txtLineTitleMarquee 		*TextLineSpecTitleMarquee
//
//		A pointer to an instance of
//		TextLineSpecTitleMarquee. This instance will be
//		configured with the Text Line Title specifications
//		contained in input parameter 'configSpecs'.
//
//		All pre-existing data value in
//		'txtLineTitleMarquee' will be deleted and reset
//		to the configuration specifications contained in
//		'configSpecs'.
//
//	configSpecs					TextLineTitleMarqueeDto
//
//		An instance of TextLineTitleMarqueeDto containing
//		configuration specifications for a
//		TextLineSpecTitleMarquee object.
//
//		The TextLineTitleMarqueeDto data transfer object
//		(DTO) is designed to store and transfer all
//		specifications necessary to produce a Title
//		Marquee for file output or text display.
//
//		NOTE: If the TitleLines array is empty, no error
//		will be returned. The user is thereafter
//		responsible for configuring and adding Title
//		Lines.
//
//		type TextLineTitleMarqueeDto struct {
//
//			StandardTitleLeftMargin		string
//				The standard left margin characters applied
//				to all Text Title Lines in the 'TitleLines'
//				array.
//
//			StandardTitleRightMargin		string
//				The standard left margin characters applied
//				to all Text Title Lines in the 'TitleLines'
//				array.
//
//			StandardMaxLineLen					int
//				The maximum number of characters allowed on
//				a text title line.
//
//				If this line length is less than 5, an
//				error will be returned.
//
//			StandardTextFieldLen		int
//				The standard field length applied to all
//				Text Title Lines in the 'TitleLines' array.
//
//				If this field length is less than 5, it
//				will be automatically defaulted to a value
//				of 'StandardMaxLineLen' - 1.
//
//			NumLeadingBlankLines    		int
//				The number of blank lines or 'new lines'
//				inserted above the Leading Solid Line.
//
//			LeadingSolidLineChar 		string
//				The character used to create the Leading
//				Solid Line displayed above the Title
//				Lines.
//
//			NumLeadingSolidLines  			int
//				The Number of Leading Solid Lines to
//				Display above the Title Lines.
//
//			NumTopTitleBlankLines    		int
//				The number of blank lines or 'new lines' to
//				insert immediately above the Title Lines
//				Display.
//
//			TitleLines            		[]TextLineSpecStandardLine
//				An array of TextLineSpecStandardLine objects
//				containing all specifications necessary to
//				display the Text Title Lines.
//
//			NumBottomTitleBlankLines 		int
//				The number of blank lines or 'new lines' to
//				insert immediately below the Title Lines
//				Display.
//
//			TrailingSolidLineChar 		string
//				The character used to create the Trailing
//				Solid Line displayed below the Title
//				Lines.
//
//			NumTrailingSolidLines 			int
//				The Number of Trailing Solid Lines to
//				Display below the Title Lines.
//
//			NumTrailingBlankLines 			int
//				The number of blank lines or 'new lines'
//				inserted after the Trailing Solid Line.
//		}
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
func (txtLineTitleMarqueeMech *textLineSpecTitleMarqueeMechanics) setTxtLineTitleMarquee(
	txtLineTitleMarquee *TextLineSpecTitleMarquee,
	configSpecs TextLineTitleMarqueeDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLineTitleMarqueeMech.lock == nil {
		txtLineTitleMarqueeMech.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeMech.lock.Lock()

	defer txtLineTitleMarqueeMech.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTitleMarqueeMechanics."+
			"setTxtLineTitleMarquee()",
		"")

	if err != nil {
		return err
	}

	if txtLineTitleMarquee == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineTitleMarquee' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if configSpecs.StandardMaxLineLen < 5 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'configSpecs' is INVALID!\n"+
			"configSpecs.StandardMaxLineLen is less than five (5)\n"+
			"configSpecs.StandardMaxLineLen = %v\n",
			ePrefix.String(),
			configSpecs.StandardMaxLineLen)

		return err
	}

	standardTextFieldLength :=
		configSpecs.StandardTextFieldLen

	if standardTextFieldLength < 5 {
		standardTextFieldLength =
			configSpecs.StandardMaxLineLen - 1 -
				len(configSpecs.StandardTitleLeftMargin) -
				len(configSpecs.StandardTitleRightMargin)

		if standardTextFieldLength < 5 {

			err = fmt.Errorf("%v\n"+
				"Error: Input parameter 'configSpecs' is INVALID!\n"+
				"After attempting to set a default value for configSpecs.StandardTextFieldLen,\n"+
				"configSpecs.StandardTextFieldLen now has a value less than five (5).\n"+
				"configSpecs.StandardTextFieldLen = %v\n",
				ePrefix.String(),
				configSpecs.StandardTextFieldLen)

			return err

		}
	}

	new(textLineSpecTitleMarqueeElectron).empty(
		txtLineTitleMarquee)

	txtLineTitleMarquee.standardTitleLeftMargin =
		configSpecs.StandardTitleLeftMargin

	txtLineTitleMarquee.standardTitleRightMargin =
		configSpecs.StandardTitleRightMargin

	txtLineTitleMarquee.standardMaxLineLen =
		configSpecs.StandardMaxLineLen

	txtLineTitleMarquee.standardTextFieldLen =
		standardTextFieldLength

	if configSpecs.NumLeadingBlankLines > 0 {

		err = txtLineTitleMarquee.leadingBlankLines.
			SetSpecBlankLines(
				configSpecs.NumLeadingBlankLines,
				"\n",
				ePrefix.XCpy("txtLineTitleMarquee.leadingBlankLines"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumLeadingSolidLines > 0 {

		err = txtLineTitleMarquee.leadingSolidLines.
			SetSolidLine(
				configSpecs.StandardTitleLeftMargin,
				configSpecs.StandardTitleRightMargin,
				configSpecs.LeadingSolidLineChar,
				configSpecs.NumLeadingSolidLines,
				ePrefix.XCpy("txtLineTitleMarquee.leadingSolidLines<-"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumTopTitleBlankLines > 0 {

		err = txtLineTitleMarquee.topTitleBlankLines.
			SetSpecBlankLines(
				configSpecs.NumTopTitleBlankLines,
				"\n",
				ePrefix.XCpy("txtLineTitleMarquee.topTitleBlankLines"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumBottomTitleBlankLines > 0 {

		err = txtLineTitleMarquee.bottomTitleBlankLines.
			SetSpecBlankLines(
				configSpecs.NumBottomTitleBlankLines,
				"\n",
				ePrefix.XCpy("txtLineTitleMarquee.bottomTitleBlankLines"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumTrailingSolidLines > 0 {

		err = txtLineTitleMarquee.trailingSolidLines.
			SetSolidLine(
				configSpecs.StandardTitleLeftMargin,
				configSpecs.StandardTitleRightMargin,
				configSpecs.TrailingSolidLineChar,
				configSpecs.NumTrailingBlankLines,
				ePrefix.XCpy("txtLineTitleMarquee.trailingSolidLines<-"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumTrailingBlankLines > 0 {

		err = txtLineTitleMarquee.trailingBlankLines.
			SetSpecBlankLines(
				configSpecs.NumTrailingBlankLines,
				"\n",
				ePrefix.XCpy("txtLineTitleMarquee.trailingBlankLines"))

		if err != nil {
			return err
		}

	}

	lenTitleLines := len(configSpecs.TitleLines)

	if lenTitleLines == 0 {
		return err
	}

	txtLineTitleMarquee.titleLines =
		make([]TextLineSpecStandardLine, lenTitleLines)

	for i := 0; i < lenTitleLines; i++ {

		err = textLineSpecStandardLineNanobot{}.ptr().
			copyIn(
				&txtLineTitleMarquee.titleLines[i],
				&configSpecs.TitleLines[i],
				ePrefix.XCpy(
					fmt.Sprintf("txtLineTitleMarquee.titleLines[%v]",
						i)))

		if err != nil {
			return err
		}
	}

	return err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// textLineSpecTitleMarqueeMechanics
//
// Provides helper methods for type
// TextLineSpecTitleMarquee
type textLineSpecTitleMarqueeMechanics struct {
	lock *sync.Mutex
}

//	setTxtLineTitleMarqueeDto
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
//	configSpecs					*TextLineTitleMarqueeDto
//
//		A pointer to an instance of
//		TextLineTitleMarqueeDto. This instance contains
//		all the text specifications required for
//		configuring a TextLineSpecTitleMarquee object.
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
func (txtLineTitleMarqueeMech *textLineSpecTitleMarqueeMechanics) setTxtLineTitleMarqueeDto(
	txtLineTitleMarquee *TextLineSpecTitleMarquee,
	configSpecs *TextLineTitleMarqueeDto,
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
			"setTxtLineTitleMarqueeDto()",
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

	if configSpecs == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'configSpecs' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(textLineTitleMarqueeDtoAtom).
		testValidityOfTitleMarqueeDto(
			configSpecs,
			ePrefix.XCpy(
				"configSpecs"))

	if err != nil {
		return err
	}

	new(textLineSpecTitleMarqueeElectron).empty(
		txtLineTitleMarquee)

	if configSpecs.TitleLines.GetNumberOfTextLines() > 0 {

		err = txtLineTitleMarquee.titleLines.CopyIn(
			&configSpecs.TitleLines,
			ePrefix.XCpy("configSpecs.TitleLines"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumLeadingBlankLines > 0 {

		err = txtLineTitleMarquee.leadingMarqueeLines.AddBlankLine(
			configSpecs.NumLeadingBlankLines,
			ePrefix.XCpy(
				"configSpecs.NumLeadingBlankLines"))

		if err != nil {
			return err
		}

	}

	var solidLineCharStr string
	var textFieldSpecs []ITextFieldSpecification
	var solidLeftMarginLabel, solidRightMarginLabel,
		txtFieldLabel TextFieldSpecLabel

	var stdLine TextLineSpecStandardLine

	if len(configSpecs.StandardSolidLineLeftMargin) > 0 {

		solidLeftMarginLabel,
			err = TextFieldSpecLabel{}.NewTextLabel(
			configSpecs.StandardSolidLineLeftMargin,
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy(
				"solidLeftMarginLabel-leading solid line"))

		if err != nil {
			return err
		}
	}

	if len(configSpecs.StandardSolidLineRightMargin) > 0 {

		solidRightMarginLabel,
			err = TextFieldSpecLabel{}.NewTextLabel(
			configSpecs.StandardSolidLineRightMargin,
			-1,
			TxtJustify.Right(),
			ePrefix.XCpy(
				"solidRightMarginLabel-leading solid line"))

		if err != nil {
			return err
		}
	}

	if len(configSpecs.LeadingSolidLineChar) > 0 &&
		configSpecs.NumLeadingSolidLines > 0 {
		// If we have left and/or right margins
		// Filler Lines are required.

		solidLineCharStr = strings.Repeat(
			configSpecs.LeadingSolidLineChar,
			configSpecs.StandardTextFieldLen)

		if len(solidLineCharStr) > configSpecs.StandardTextFieldLen {
			solidLineCharStr = solidLineCharStr[0:configSpecs.StandardTextFieldLen]
		}

		txtFieldLabel,
			err = TextFieldSpecLabel{}.
			NewTextLabel(
				solidLineCharStr,
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					"txtFieldLabel-Leading Solid Line"))

		if err != nil {
			return err
		}

		textFieldSpecs = nil

		if len(configSpecs.StandardSolidLineLeftMargin) > 0 {
			textFieldSpecs = append(
				textFieldSpecs,
				&solidLeftMarginLabel)
		}

		textFieldSpecs = append(
			textFieldSpecs,
			&txtFieldLabel)

		if len(configSpecs.StandardSolidLineRightMargin) > 0 {
			textFieldSpecs = append(
				textFieldSpecs,
				&solidRightMarginLabel)
		}

		stdLine,
			err = TextLineSpecStandardLine{}.NewStandardLineAllParms(
			configSpecs.NumLeadingSolidLines,
			textFieldSpecs,
			[]rune{'\n'},
			false,
			ePrefix.XCpy(
				"configSpecs.NumLeadingSolidLines"))

		if err != nil {
			return err
		}

		err = txtLineTitleMarquee.leadingMarqueeLines.
			AddTextLineSpec(
				&stdLine,
				ePrefix.XCpy(""+
					" txtLineTitleMarquee.leadingMarqueeLines<-LeadingSolidLine"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumTopTitleBlankLines > 0 {

		solidLineCharStr = strings.Repeat(
			" ",
			configSpecs.StandardTextFieldLen)

		txtFieldLabel,
			err = TextFieldSpecLabel{}.
			NewTextLabel(
				solidLineCharStr,
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					"txtFieldLabel-NumTopTitleBlankLines"))

		if err != nil {
			return err
		}

		textFieldSpecs = nil

		if len(configSpecs.StandardSolidLineLeftMargin) > 0 {
			textFieldSpecs = append(
				textFieldSpecs,
				&solidLeftMarginLabel)
		}

		textFieldSpecs = append(
			textFieldSpecs,
			&txtFieldLabel)

		if len(configSpecs.StandardSolidLineRightMargin) > 0 {
			textFieldSpecs = append(
				textFieldSpecs,
				&solidRightMarginLabel)
		}

		stdLine,
			err = TextLineSpecStandardLine{}.NewStandardLineAllParms(
			configSpecs.NumTopTitleBlankLines,
			textFieldSpecs,
			[]rune{'\n'},
			false,
			ePrefix.XCpy(
				"configSpecs.NumTopTitleBlankLines"))

		if err != nil {
			return err
		}

		err = txtLineTitleMarquee.leadingMarqueeLines.
			AddTextLineSpec(
				&stdLine,
				ePrefix.XCpy(""+
					" txtLineTitleMarquee.leadingMarqueeLines<-LeadingSolidLine"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumBottomTitleBlankLines > 0 {

		solidLineCharStr = strings.Repeat(
			" ",
			configSpecs.StandardTextFieldLen)

		txtFieldLabel,
			err = TextFieldSpecLabel{}.
			NewTextLabel(
				solidLineCharStr,
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					"txtFieldLabel-NumBottomTitleBlankLines"))

		if err != nil {
			return err
		}

		textFieldSpecs = nil

		if len(configSpecs.StandardSolidLineLeftMargin) > 0 {
			textFieldSpecs = append(
				textFieldSpecs,
				&solidLeftMarginLabel)
		}

		textFieldSpecs = append(
			textFieldSpecs,
			&txtFieldLabel)

		if len(configSpecs.StandardSolidLineRightMargin) > 0 {
			textFieldSpecs = append(
				textFieldSpecs,
				&solidRightMarginLabel)
		}

		stdLine,
			err = TextLineSpecStandardLine{}.NewStandardLineAllParms(
			configSpecs.NumBottomTitleBlankLines,
			textFieldSpecs,
			[]rune{'\n'},
			false,
			ePrefix.XCpy(
				"configSpecs.NumTopTitleBlankLines"))

		if err != nil {
			return err
		}

		err = txtLineTitleMarquee.trailingMarqueeLines.
			AddTextLineSpec(
				&stdLine,
				ePrefix.XCpy(""+
					" txtLineTitleMarquee.trailingMarqueeLines<-LeadingSolidLine"))

		if err != nil {
			return err
		}
	}

	if len(configSpecs.TrailingSolidLineChar) > 0 &&
		configSpecs.NumTrailingSolidLines > 0 {

		solidLineCharStr = strings.Repeat(
			configSpecs.TrailingSolidLineChar,
			configSpecs.StandardTextFieldLen)

		if len(solidLineCharStr) > configSpecs.StandardTextFieldLen {
			solidLineCharStr = solidLineCharStr[0:configSpecs.StandardTextFieldLen]
		}

		txtFieldLabel,
			err = TextFieldSpecLabel{}.
			NewTextLabel(
				solidLineCharStr,
				-1,
				TxtJustify.Left(),
				ePrefix.XCpy(
					"txtFieldLabel-Trailing Solid Line"))

		if err != nil {
			return err
		}

		textFieldSpecs = nil

		if len(configSpecs.StandardSolidLineLeftMargin) > 0 {
			textFieldSpecs = append(
				textFieldSpecs,
				&solidLeftMarginLabel)
		}

		textFieldSpecs = append(
			textFieldSpecs,
			&txtFieldLabel)

		if len(configSpecs.StandardSolidLineRightMargin) > 0 {
			textFieldSpecs = append(
				textFieldSpecs,
				&solidRightMarginLabel)
		}

		stdLine,
			err = TextLineSpecStandardLine{}.NewStandardLineAllParms(
			configSpecs.NumTrailingSolidLines,
			textFieldSpecs,
			[]rune{'\n'},
			false,
			ePrefix.XCpy(
				"configSpecs.NumTrailingSolidLines"))

		if err != nil {
			return err
		}

		err = txtLineTitleMarquee.trailingMarqueeLines.
			AddTextLineSpec(
				&stdLine,
				ePrefix.XCpy(""+
					" txtLineTitleMarquee.trailingMarqueeLines<-TrailingSolidLine"))

		if err != nil {
			return err
		}

	}

	if configSpecs.NumTrailingBlankLines > 0 {

		err = txtLineTitleMarquee.trailingMarqueeLines.AddBlankLine(
			configSpecs.NumTrailingBlankLines,
			ePrefix.XCpy(
				"configSpecs.NumTrailingBlankLines"))

		if err != nil {
			return err
		}
	}

	return err
}

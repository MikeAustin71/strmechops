package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// textLineSpecTitleMarqueeNanobot
//
// Provides helper methods for type
// TextLineSpecTitleMarquee
type textLineSpecTitleMarqueeNanobot struct {
	lock *sync.Mutex
}

//	copy
//
//	Copies the data fields from input parameter
//	'sourceTitleMarquee' to input parameter
//	'destinationTitleMarquee'.
//
//	If input parameter 'sourceTitleMarquee' is judged to
//	be invalid, this method will return an error.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Be advised that all data fields in
//	'destinationTitleMarquee' will be deleted and
//	overwritten by this method.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	destinationTitleMarquee		*TextLineSpecTitleMarquee
//
//		A pointer to a TextLineSpecTitleMarquee instance.
//		All the member variable data fields in this
//		object will be replaced by data values copied
//		from input parameter 'sourceNStrIntSeparator'.
//
//		'destinationPosNumSignSpec' is the destination for
//		this copy operation.
//
//
//	sourceTitleMarquee			*TextLineSpecTitleMarquee
//
//		A pointer to another TextLineSpecTitleMarquee
//		instance. All the member variable data values
//		from this object will be copied to corresponding
//		member variables in 'destinationTitleMarquee'.
//
//		'sourceTitleMarquee' is the source for this copy
//		operation.
//
//		If 'sourceTitleMarquee' is found to be invalid,
//		an error will be returned.
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
func (txtLineTitleMarqueeNanobot *textLineSpecTitleMarqueeNanobot) copy(
	destinationTitleMarquee *TextLineSpecTitleMarquee,
	sourceTitleMarquee *TextLineSpecTitleMarquee,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLineTitleMarqueeNanobot.lock == nil {
		txtLineTitleMarqueeNanobot.lock = new(sync.Mutex)
	}

	txtLineTitleMarqueeNanobot.lock.Lock()

	defer txtLineTitleMarqueeNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTitleMarqueeNanobot."+
			"copy()",
		"")

	if err != nil {
		return err
	}

	if destinationTitleMarquee == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationTitleMarquee' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTitleMarquee == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationTitleMarquee' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(textLineSpecTitleMarqueeElectron).empty(
		destinationTitleMarquee)

	destinationTitleMarquee.standardTitleLeftMargin =
		sourceTitleMarquee.standardTitleLeftMargin

	destinationTitleMarquee.standardTitleRightMargin =
		sourceTitleMarquee.standardTitleRightMargin

	destinationTitleMarquee.standardMaxLineLen =
		sourceTitleMarquee.standardMaxLineLen

	destinationTitleMarquee.standardTextFieldLen =
		sourceTitleMarquee.standardTextFieldLen

	if sourceTitleMarquee.leadingBlankLines.numBlankLines > 0 {
		err = destinationTitleMarquee.leadingBlankLines.
			CopyIn(
				&sourceTitleMarquee.leadingBlankLines,
				ePrefix.XCpy("sourceTitleMarquee."+
					"leadingBlankLines"))

		if err != nil {
			return err
		}
	}

	if sourceTitleMarquee.leadingSolidLines.solidLineCharsRepeatCount > 0 {

		err = destinationTitleMarquee.leadingSolidLines.
			CopyIn(&sourceTitleMarquee.leadingSolidLines,
				ePrefix.XCpy(
					"<-sourceTitleMarquee.leadingSolidLines"))

		if err != nil {
			return err
		}
	}

	if sourceTitleMarquee.topTitleBlankLines.numBlankLines > 0 {
		err = destinationTitleMarquee.topTitleBlankLines.
			CopyIn(
				&sourceTitleMarquee.topTitleBlankLines,
				ePrefix.XCpy("sourceTitleMarquee."+
					"topTitleBlankLines"))

		if err != nil {
			return err
		}
	}

	if sourceTitleMarquee.titleLines.GetNumberOfTextLines() > 0 {

		err = destinationTitleMarquee.titleLines.CopyIn(
			&sourceTitleMarquee.titleLines,
			ePrefix.XCpy(
				"<-sourceTitleMarquee.titleLines"))

		if err != nil {
			return err
		}
	}

	if sourceTitleMarquee.bottomTitleBlankLines.numBlankLines > 0 {
		err = destinationTitleMarquee.bottomTitleBlankLines.
			CopyIn(
				&sourceTitleMarquee.bottomTitleBlankLines,
				ePrefix.XCpy("bottomTitleBlankLines."+
					"topTitleBlankLines"))

		if err != nil {
			return err
		}
	}

	if sourceTitleMarquee.trailingSolidLines.solidLineCharsRepeatCount > 0 {

		err = destinationTitleMarquee.trailingSolidLines.
			CopyIn(&sourceTitleMarquee.trailingSolidLines,
				ePrefix.XCpy(
					"<-sourceTitleMarquee.trailingSolidLines"))

		if err != nil {
			return err
		}
	}

	if sourceTitleMarquee.trailingBlankLines.numBlankLines > 0 {
		err = destinationTitleMarquee.trailingBlankLines.
			CopyIn(
				&sourceTitleMarquee.trailingBlankLines,
				ePrefix.XCpy("trailingBlankLines."+
					"topTitleBlankLines"))

	}

	return err
}

package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type TextLineSpecTitleMarquee struct {
	standardTitleLeftMargin  string
	standardTitleRightMargin string
	standardMaxLineLen       int
	standardTextFieldLen     int
	leadingBlankLines        TextLineSpecBlankLines
	leadingSolidLines        TextLineSpecSolidLine
	topTitleBlankLines       TextLineSpecBlankLines
	titleLines               TextLineSpecLinesCollection
	bottomTitleBlankLines    TextLineSpecBlankLines
	trailingSolidLines       TextLineSpecSolidLine
	trailingBlankLines       TextLineSpecBlankLines

	lock *sync.Mutex
}

//	NewAllParams
//
//	Creates and returns a new instance of
//	TextLineSpecTitleMarquee.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	standardTitleLeftMargin		string
//
//		The standard left margin characters applied
//		to all Text Title Lines in the Title Lines
//		array (TextLineSpecTitleMarquee.titleLines).
//
//	standardTitleRightMargin		string
//
//		The standard right margin characters applied
//		to all Text Title Lines in the Title Lines
//		array (TextLineSpecTitleMarquee.titleLines).
//
//	standardMaxLineLen					int
//
//		The maximum number of characters allowed on
//		a text title line. This maximum limit will be
//		applied to the length of all text lines generated
//		by the returned instance of
//		TextLineSpecTitleMarquee.
//
//	standardTextFieldLen		int
//
//		The standard field length applied to Text
//		Title Lines in the 'TitleLines' array unless
//		overridden by user customizations.
//
//		If the standardTextFieldLen exceeds the value of
//		the Maximum Available Text Field Length, it will
//		be reset and defaulted to the Maximum Available
//		Text Field Length.
//
//		The Maximum Available Text Field Length is
//		calculated as follows:
//
//		Maximum Available Text Field Length =
//			TextLineTitleMarqueeDto.StandardMaxLineLen -
//			1 -
//			len(TextLineTitleMarqueeDto.StandardTitleLeftMargin) -
//			len(TextLineTitleMarqueeDto.StandardTitleRightMargin)
//
//	numLeadingBlankLines		int
//
//		The number of blank lines or 'new lines'
//		inserted above the Leading Solid Line.
//
//	leadingSolidLineChar		string
//
//		The character used to create the Leading
//		Solid Line displayed above the Title
//		Lines.
//
//	numLeadingSolidLines		int
//
//		The Number of Leading Solid Lines to display
//		above the Title Lines.
//
//	numTopTitleBlankLines		int
//
//		The number of blank lines or 'new lines' to
//		insert immediately above the Title Lines
//		Display.
//
//	titleLines					TextLineSpecLinesCollection
//
//		A collection of text line objects containing all
//		specifications necessary to display the Text
//		Title Lines.
//
//	numBottomTitleBlankLines	int
//
//		The number of blank lines or 'new lines' to
//		insert immediately below the Title Lines Display.
//
//	trailingSolidLineChar		string
//
//		The character used to create the Trailing Solid
//		Line displayed below the Title Lines.
//
//	numTrailingSolidLines		int
//
//		The Number of Trailing Solid Lines to display
//		below the Title Lines.
//
//	numTrailingBlankLines		int
//
//		The number of blank lines or 'new lines' inserted
//		after the Trailing Solid Line.
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
//	TextLineSpecTitleMarquee
//
//		If this method completes successfully, a new
//		instance of TextLineSpecTitleMarquee will be
//		returned.
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
func (txtLineSpecTitleMarquee *TextLineSpecTitleMarquee) NewAllParams(
	standardTitleLeftMargin string,
	standardTitleRightMargin string,
	standardMaxLineLen int,
	standardTextFieldLen int,
	numLeadingBlankLines int,
	leadingSolidLineChar string,
	numLeadingSolidLines int,
	numTopTitleBlankLines int,
	titleLines TextLineSpecLinesCollection,
	numBottomTitleBlankLines int,
	trailingSolidLineChar string,
	numTrailingSolidLines int,
	numTrailingBlankLines int,
	errorPrefix interface{}) (
	TextLineSpecTitleMarquee,
	error) {

	if txtLineSpecTitleMarquee.lock == nil {
		txtLineSpecTitleMarquee.lock = new(sync.Mutex)
	}

	txtLineSpecTitleMarquee.lock.Lock()

	defer txtLineSpecTitleMarquee.lock.Unlock()

	var newTxtLineTitle TextLineSpecTitleMarquee
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecTitleMarquee."+
			"NewAllParams()",
		"")

	if err != nil {
		return newTxtLineTitle, err
	}

	titleMarqueeDto := TextLineTitleMarqueeDto{
		StandardTitleLeftMargin:  standardTitleLeftMargin,
		StandardTitleRightMargin: standardTitleRightMargin,
		StandardMaxLineLen:       standardMaxLineLen,
		StandardTextFieldLen:     standardTextFieldLen,
		NumLeadingBlankLines:     numLeadingBlankLines,
		LeadingSolidLineChar:     leadingSolidLineChar,
		NumLeadingSolidLines:     numLeadingSolidLines,
		NumTopTitleBlankLines:    numTopTitleBlankLines,
		NumBottomTitleBlankLines: numBottomTitleBlankLines,
		TrailingSolidLineChar:    trailingSolidLineChar,
		NumTrailingSolidLines:    numTrailingSolidLines,
		NumTrailingBlankLines:    numTrailingBlankLines,
	}

	if titleLines.GetNumberOfTextLines() == 0 {

		titleMarqueeDto.TitleLines.Empty()

	} else {

		err = titleMarqueeDto.TitleLines.CopyIn(
			&titleLines,
			ePrefix.XCpy("<-titleLines"))

		if err != nil {
			return newTxtLineTitle, err
		}

	}

	err = new(textLineSpecTitleMarqueeMechanics).
		setTxtLineTitleMarquee(
			&newTxtLineTitle,
			titleMarqueeDto,
			ePrefix.XCpy(
				"newTxtLineTitle"))

	return newTxtLineTitle, err
}

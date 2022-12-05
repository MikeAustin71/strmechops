package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// textLineSpecTitleMarqueeMolecule
//
// Provides helper methods for type
// TextLineSpecTitleMarquee
type textLineSpecTitleMarqueeMolecule struct {
	lock *sync.Mutex
}

//	getFormattedText
//
//	Returns the formatted text generated by the Text Line
//	Specification Title Marquee instance passed as input
//	parameter, 'txtLineTitleMarquee'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	 strBuilder					*strings.Builder
//
//			A pointer to an instance of *strings.Builder. The
//			formatted text characters produced by this method
//			will be written to this instance of
//			strings.Builder.
//
//	txtLineTitleMarquee			*TextLineSpecTitleMarquee
//
//			A pointer to an instance of
//			TextLineSpecTitleMarquee. The member variables
//			encapsulated by this instance will be used to
//			generate formatted text for text display, file
//			output and printing.
//
//			If this instance of TextLineSpecTitleMarquee
//			is judged to be invalid, an error will be
//			returned.
//
//		errPrefDto					*ePref.ErrPrefixDto
//
//			This object encapsulates an error prefix string
//			which is included in all returned error
//			messages. Usually, it contains the name of the
//			calling method or methods listed as a function
//			chain.
//
//			If no error prefix information is needed, set
//			this parameter to 'nil'.
//
//			Type ErrPrefixDto is included in the 'errpref'
//			software package:
//				"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	maxLineLength				int
//
//		The maximum length calculated from all
//		the text lines produced by this method.
//
//	totalStrLength				int
//
//		The total string length of all strings written
//		to the strings.Builder instance passed by input
//		parameter 'strBuilder'.
//
//	err							error
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
func (txtLineTitleMolecule *textLineSpecTitleMarqueeMolecule) getFormattedText(
	strBuilder *strings.Builder,
	txtLineTitleMarquee *TextLineSpecTitleMarquee,
	errPrefDto *ePref.ErrPrefixDto) (
	maxLineLength int,
	totalStrLength int,
	err error) {

	if txtLineTitleMolecule.lock == nil {
		txtLineTitleMolecule.lock = new(sync.Mutex)
	}

	txtLineTitleMolecule.lock.Lock()

	defer txtLineTitleMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTitleMarqueeMolecule."+
			"getFormattedText()",
		"")

	if err != nil {
		return maxLineLength, totalStrLength, err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is a nil pointer!\n",
			ePrefix.String())

		return maxLineLength, totalStrLength, err
	}

	if txtLineTitleMarquee == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineTitleMarquee' is a nil pointer!\n",
			ePrefix.String())

		return maxLineLength, totalStrLength, err
	}

	_,
		err = new(textLineSpecTitleMarqueeElectron).
		testValidityTitleMarquee(
			txtLineTitleMarquee,
			ePrefix.XCpy(
				"testValidityTitleMarquee"))

	if err != nil {

		return maxLineLength, totalStrLength, err

	}

	var str string
	var lenStr int

	if txtLineTitleMarquee.leadingBlankLines.numBlankLines > 0 {

		str,
			err = txtLineTitleMarquee.leadingBlankLines.GetFormattedText(
			ePrefix.XCpy(
				"txtLineTitleMarquee.leadingBlankLines"))

		if err != nil {

			return maxLineLength, totalStrLength, err

		}

		lenStr = len(str)

		if lenStr > maxLineLength {
			maxLineLength = lenStr
		}

		strBuilder.WriteString(str)
	}

	if txtLineTitleMarquee.leadingSolidLines.solidLineCharsRepeatCount > 0 {

		str,
			err = txtLineTitleMarquee.leadingSolidLines.GetFormattedText(
			ePrefix.XCpy(
				"txtLineTitleMarquee.leadingSolidLines"))

		if err != nil {

			return maxLineLength, totalStrLength, err

		}

		lenStr = len(str)

		if lenStr > maxLineLength {
			maxLineLength = lenStr
		}

		strBuilder.WriteString(str)
	}

	if txtLineTitleMarquee.topTitleBlankLines.numBlankLines > 0 {

		str,
			err = txtLineTitleMarquee.topTitleBlankLines.GetFormattedText(
			ePrefix.XCpy(
				"txtLineTitleMarquee.topTitleBlankLines"))

		if err != nil {

			return maxLineLength, totalStrLength, err

		}

		lenStr = len(str)

		if lenStr > maxLineLength {
			maxLineLength = lenStr
		}

		strBuilder.WriteString(str)
	}

	numTitleLines := txtLineTitleMarquee.titleLines.GetNumberOfTextLines()

	var iTextLine ITextLineSpecification

	for i := 0; i < numTitleLines; i++ {

		iTextLine,
			err =
			txtLineTitleMarquee.titleLines.PeekAtTextLine(
				i,
				ePrefix.XCpy(
					"txtLineTitleMarquee."+
						"titleLines"))

		if err != nil {

			return maxLineLength, totalStrLength, err

		}

		str,
			err = iTextLine.GetFormattedText(
			ePrefix.XCpy(
				fmt.Sprintf(
					"iTextLine[%v]",
					i)))

		if err != nil {

			return maxLineLength, totalStrLength, err

		}

		lenStr = len(str)

		if lenStr > maxLineLength {
			maxLineLength = lenStr
		}

		strBuilder.WriteString(str)

	}

	if txtLineTitleMarquee.bottomTitleBlankLines.numBlankLines > 0 {

		str,
			err = txtLineTitleMarquee.bottomTitleBlankLines.GetFormattedText(
			ePrefix.XCpy(
				"txtLineTitleMarquee.bottomTitleBlankLines"))

		if err != nil {

			return maxLineLength, totalStrLength, err

		}

		lenStr = len(str)

		if lenStr > maxLineLength {
			maxLineLength = lenStr
		}

		strBuilder.WriteString(str)

	}

	if txtLineTitleMarquee.trailingSolidLines.solidLineCharsRepeatCount > 0 {

		str,
			err = txtLineTitleMarquee.trailingSolidLines.GetFormattedText(
			ePrefix.XCpy(
				"trailingSolidLines.trailingSolidLines"))

		if err != nil {

			return maxLineLength, totalStrLength, err

		}

		lenStr = len(str)

		if lenStr > maxLineLength {
			maxLineLength = lenStr
		}

		strBuilder.WriteString(str)

	}

	if txtLineTitleMarquee.trailingBlankLines.numBlankLines > 0 {

		str,
			err = txtLineTitleMarquee.trailingBlankLines.GetFormattedText(
			ePrefix.XCpy(
				"txtLineTitleMarquee.trailingBlankLines"))

		if err != nil {

			return maxLineLength, totalStrLength, err

		}

		lenStr = len(str)

		if lenStr > maxLineLength {
			maxLineLength = lenStr
		}

		strBuilder.WriteString(str)

	}

	totalStrLength = strBuilder.Len()

	return maxLineLength, totalStrLength, err
}

//	readBytes
//
//	This helper method is designed to support the
//	io.Reader interface.
//
//	This method reads up to len(p) bytes into p. It
//	returns the number of bytes read (0 <= n <= len(p))
//	and any error encountered. Even if 'read' returns
//	n < len(p), it may use all of p as scratch space
//	during the call.
//
//	If some data is available but not len(p) bytes,
//	readBytes() conventionally returns what is available
//	instead of waiting for more.
//
//	When this method encounters an error or end-of-file
//	condition after successfully reading n > 0 bytes, it
//	returns the number of bytes read. It may return the
//	(non-nil) error from the same call or return the
//	error (and n == 0) from a subsequent call.
//
//	An instance of this general case is that a Reader
//	returning a non-zero number of bytes at the end of
//	the input stream may return either err == EOF or
//	err == nil. The next Read operation should return 0,
//	EOF.
//
//	Callers should always process the n > 0 bytes
//	returned before considering the error err. Doing so
//	correctly handles I/O errors that happen after
//	reading some bytes and also both of the allowed EOF
//	behaviors.
//
//	This method supports buffered 'read' operations.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	textReader					*strings.Reader
//
//		A pointer to an instance of strings.Reader. The
//		text string encapsulated by 'textReader' will be
//		written to the byte buffer 'p'.
//
//	p							[]byte
//
//		The byte buffer into which the text string
//		encapsulated by parameter 'textReader' will be
//		written.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//		n							int
//	    - The number of bytes written to byte buffer 'p'.
//
//
//		err							error
//
//			If this method completes successfully, this
//			returned error Type is set equal to 'nil'. If
//			errors are encountered during processing, the
//			returned error Type will encapsulate an error
//			message.
//
//			readBytes() reads up to len(p) bytes into p. It
//			returns the number of bytes read
//			(0 <= n <= len(p)) and any error encountered.
//			Even if readBytes() returns n < len(p), it may
//			use all of p as scratch space during the call. If
//			some data is available but not len(p) bytes,
//			readBytes() conventionally returns what is
//			available instead of waiting for more.
//
//			When readBytes() encounters an error or
//			end-of-file condition after successfully reading
//			n > 0 bytes, it returns the number of bytes read.
//			It may return the (non-nil) error from the same
//			call or return the error (and n == 0) from a
//			subsequent call. An instance of this general case
//			is that a Reader returning a non-zero number of
//			bytes at the end of the input stream may return
//			either err == EOF or err == nil. The next Read
//			operation should return 0, EOF.
//
//			If an error message is returned and the error is
//			NOT equal to io.EOF, the text value for input
//			parameter 'errPrefDto' (error prefix) will be
//			prefixed or attached at the beginning of the
//			error message.
func (txtLineTitleMolecule *textLineSpecTitleMarqueeMolecule) readBytes(
	textReader *strings.Reader,
	p []byte,
	errPrefDto *ePref.ErrPrefixDto) (
	n int,
	err error) {

	if txtLineTitleMolecule.lock == nil {
		txtLineTitleMolecule.lock = new(sync.Mutex)
	}

	txtLineTitleMolecule.lock.Lock()

	defer txtLineTitleMolecule.lock.Unlock()

	n = 0

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecTitleMarqueeMolecule."+
			"readBytes()",
		"")

	if err != nil {
		return n, err
	}

	if textReader == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textReader' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return n, err
	}

	pLen := len(p)

	if pLen == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input byte array 'p' has zero length!\n",
			ePrefix)

		return n, err
	}

	n,
		err = textReader.Read(p)

	return n, err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

type textSpecificationNanobot struct {
	lock *sync.Mutex
}

// buildFormattedMarqueeTitle - Builds formatted output text for a
// marquee with title.
//
// The marquee format consists of:
//
//    Blank Line
//    [SPACE]Solid Line[SPACE]
//    Title 1 Text
//    Title 2 Text - Optional
//    DateTime Text - Optional
//    [SPACE]Solid Line[SPACE]
//    Blank Line
//
// Example Text Output Format:
//
//    Blank Line
//    [SPACE]==============================================[SPACE]
//                           Title 1
//                     Title 2 (If available)
//           Monday 2006-01-02 15:04:05.000000000 -0700 MST
//    [SPACE]==============================================[SPACE]
//    Blank Line
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  solidLineChar              string
//     - The text character or text characters which will be used
//       to construct the two solid lines.
//
//
//  title1                     string
//     - The Primary Title String. If this string is empty with a
//       length of zero (0), an error will be returned.
//
//
//  title1Justify              TextJustify
//     - An enumeration which specifies the justification of the
//       'title1' string within the field specified by 'lineLen'.
//
//       Text justification can only be evaluated in the context of
//       a text string, field length and a text justification
//       object of type TextJustify. This is because text strings
//       with a field length equal to or less than the length of
//       that text string never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the display field length ('lineLen') is
//       greater than the length of the text string ('title1'),
//       text justification must be equal to one of these three
//       valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
//
//  title2                     string
//     - Optional. The second title string. If this string has a
//       string length greater than zero (0), it will be displayed
//       on the line below 'title1'.
//
//       This parameter is optional. If 'title2' is empty with a
//       string length of zero (0), it will be skipped and only
//       the 'title1' string will be included in the formatted
//       text output. In the case of an empty 'title2' string, no
//       error will be returned.
//
//
//  title2Justify              TextJustify
//     - An enumeration which specifies the justification of the
//       'title2' string within the field specified by 'lineLen'.
//
//       Text justification can only be evaluated in the context of
//       a text string, field length and a text justification
//       object of type TextJustify. This is because text strings
//       with a field length equal to or less than the length of
//       that text string never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the display field length ('lineLen') is
//       greater than the length of the text string ('title2'),
//       text justification must be equal to one of these three
//       valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
//
//  dateTime                   time.Time
//     - If dateTime.IsZero() == true, no text line will be
//       produced.
//
//       If dateTime.IsZero() == false, a formatted line of text
//       will be generated just below 'title2'.
//
//
//  dateTimeFormat             string
//     - If 'dateTime' is populated, a format string will be
//       applied to the date in order to produce a date time
//       text string.
//
//       If 'dateTimeFormat' is submitted as an empty string or if
//       the 'dateTimeFormat' string is an invalid format, it will
//       be defaulted to the following format:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  dateTimeJustify            TextJustify
//     - An enumeration which specifies the justification of the
//       formatted 'dateTime' string within the field specified by
//       'lineLen'.
//
//       Text justification can only be evaluated in the context of
//       a text string, field length and a text justification
//       object of type TextJustify. This is because text strings
//       with a field length equal to or less than the length of
//       that text string never use text justification. In these
//       cases, text justification is completely ignored.
//
//       If the display field length ('lineLen') is
//       greater than the length of the formatted date time string
//       ('dateTime'), text justification must be equal to one of
//       these three valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
//
//  lineLen                    int
//     - The total length of each line of text produced in the
//       title marquee. If 'lineLen' has a value less than one (1),
//       an error will be returned.
//
//
//  strBuilder                 *strings.Builder
//     - A pointer to a type strings.Builder. All the formatted
//       text produced by this method will be written to this
//       instance of strings.Builder.
//
//
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of TextFieldSpecLabel are found to be
//       invalid, this method will return an error. If the member
//       data variables are determined to be valid, this error
//       return parameter will be set to 'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (textSpecNanobot *textSpecificationNanobot) buildFormattedMarqueeTitle(
	solidLineChar string,
	title1 string,
	title1Justify TextJustify,
	title2 string,
	title2Justify TextJustify,
	dateTime time.Time,
	dateTimeFormat string,
	dateTimeJustify TextJustify,
	lineLen int,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if textSpecNanobot.lock == nil {
		textSpecNanobot.lock = new(sync.Mutex)
	}

	textSpecNanobot.lock.Lock()

	defer textSpecNanobot.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textSpecificationNanobot."+
			"buildFormattedMarqueeTitle()",
		"")

	if err != nil {

		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(title1) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'title1' is invalid!\n"+
			"'title1' is empty and has a string length of zero.\n",
			ePrefix.String())

		return err
	}

	if len(solidLineChar) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'solidLineChar' is invalid!\n"+
			"'solidLineChar' is empty and has a string length of zero.\n",
			ePrefix.String())

		return err
	}

	if lineLen < 1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'lineLen' is invalid!\n"+
			"'lineLen' should be set to the total length of\n"+
			"all text lines.\n"+
			"'lineLen' has a value less than one (1).\n"+
			"lineLen = '%v'\n",
			ePrefix.String(),
			lineLen)

		return err

	}

	var solidLine TextLineSpecSolidLine

	// Title Marquee
	solidLine,
		err = TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		1,
		1,
		solidLineChar,
		lineLen-2,
		"\n",
		ePrefix.XCpy(
			"solidLine"))

	if err != nil {

		return err
	}

	var blankLine *TextLineSpecBlankLines

	blankLine,
		err = TextLineSpecBlankLines{}.NewPtrBlankLines(
		1,
		"\n",
		ePrefix.XCpy(
			"blankLine"))

	if err != nil {

		return err
	}

	err = blankLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"blankLine#1"))

	if err != nil {
		return err
	}

	err = solidLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"solidLine#1"))

	if err != nil {

		return err
	}

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	_,
		err = stdLine.AddTextFieldLabel(
		title1,
		lineLen,
		title1Justify,
		ePrefix.XCpy(
			"stdLine<-title1"))

	if err != nil {

		return err
	}

	err = stdLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"title1-stdLine->strBuilder"))

	if err != nil {

		return err
	}

	stdLine.EmptyTextFields()

	if len(title2) > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			title2,
			lineLen,
			title2Justify,
			ePrefix.XCpy(""+
				"stdLine<-title2Label"))

		if err != nil {

			return err
		}

		err = stdLine.TextBuilder(
			strBuilder,
			ePrefix.XCpy(
				"title2-stdLine->strBuilder"))

		if err != nil {

			return err
		}

		stdLine.EmptyTextFields()

	}

	if !dateTime.IsZero() {

		if len(dateTimeFormat) == 0 {
			dateTimeFormat =
				textSpecificationMolecule{}.ptr().
					getDefaultDateTimeFormat()
		}

		_,
			err = stdLine.AddTextFieldDateTime(
			dateTime,
			lineLen,
			dateTimeFormat,
			dateTimeJustify,
			ePrefix.XCpy(""+
				"stdLine<-title2Label"))

		if err != nil {

			return err
		}

		err = stdLine.TextBuilder(
			strBuilder,
			ePrefix.XCpy(
				"dateTime-stdLine->strBuilder"))

		if err != nil {

			return err
		}

		stdLine.EmptyTextFields()
	}

	err = solidLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"solidLine#1"))

	if err != nil {

		return err
	}

	err = blankLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"blankLine#1"))

	return err
}

// buildFormattedSingleParameterText - Configures formatted text
// output for a single text label and a single parameter value on
// one line of text.
//
// The formatted parameter output consists of the following text
// fields located on a single text line:
//   [LABEL LEFT MARGINS SPACES]
//   Text Label
//   [LABEL RIGHT MARGIN TEXT]
//   Parameter Value
//   [PARAM RIGHT MARGIN TEXT]
//   [LINE TERMINATOR]
//
//
// Example Text Output Format:
//
// [" "]["Target String Length:][": "]["47"][" "]['\n']
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  paramLabelLeftMargin      string
//     - This string defines the text which will be placed on the
//       left side of the 'paramLabel' string.
//
//       If no left margin is required, submit this parameter as a
//       zero length or empty string and no left margin will be
//       created.
//
//
//  paramLabel                 string
//     - This string contains text describing the parameter value
//       which will follow this label in the formatted text output
//       string.
//
//
// 	paramLabelFieldLen         int
//     - The length of the text field in which the 'paramLabel'
//       will be displayed. If 'paramLabelFieldLen' is less than
//       the length of the 'paramLabel' string, it will be
//       automatically set equal to the 'paramLabel' string length.
//
//       To automatically set the value of 'paramLabelFieldLen' to
//       the length of 'paramLabel', set this parameter to a value
//       of minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  paramLabelJustify          TextJustify
//     - An enumeration which specifies the justification of the
//       'paramLabel' string within the field specified by
//       'paramLabelFieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a Text Justification object
//       of type TextJustify. This is because text labels with a
//       field length equal to or less than the length of the text
//       label never use text justification. In these cases, text
//       justification is completely ignored.
//
//       If the display field length ('paramLabelFieldLen') is
//       greater than the length of the text label ('paramLabel'),
//       text justification must be equal to one of these three
//       valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
//
//  paramLabelRightMargin      string
//     - The contents of the string will be used as the right
//       margin for the 'paramLabel' text string. This parameter is
//       commonly used to specify a space or colon-space ": "
//       positioned to the right of 'paramLabel' and immediately
//       prior to 'paramValueStr'.
//
//       If no right margin is required, set
//       'paramLabelRightMargin' to a zero length or empty string
//       and no right margin will be created.
//
//
//  paramValueStr              string
//     - The value of the parameter which will be displayed as
//       text.
//
//
//  paramValueFieldLen         int
//     - The length of the text field in which the 'paramValueStr'
//       will be displayed. If 'paramValueFieldLen' is less than
//       the length of the 'paramValueStr' string, it will be
//       automatically set equal to the 'paramValueStr' string
//       length.
//
//       To automatically set the value of 'paramValueFieldLen' to
//       the length of 'paramValueStr', set this parameter to a
//       value of minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  paramValueJustify          TextJustify
//     - An enumeration which specifies the justification of the
//       'paramValueStr' string within the field specified by
//       'paramValueFieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a Text Justification object
//       of type TextJustify. This is because text labels with a
//       field length equal to or less than the length of the text
//       label never use text justification. In these cases, text
//       justification is completely ignored.
//
//       If the display field length ('paramValueFieldLen') is
//       greater than the length of the text string ('paramValueStr'),
//       text justification must be equal to one of these three
//       valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
//
//  paramValueRightMargin      string
//     - The contents of the string will be used as the right
//       margin for the 'paramValueStr' text string.
//
//       If no right margin is required, set
//       'paramValueRightMargin' to a zero length or empty string
//       and no right margin will be created.
//
//
//  lineTerminator             string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text output string
//       created by this method. The most common usage sets this
//       string to a new line character ("\n").
//
//       If no Line Terminator is required, set 'lineTerminator' to
//       a zero length or empty string and no line termination
//       characters will be created.
//
//
//  strBuilder                *strings.Builder
//  - A pointer to an instance of strings.Builder. The formatted
//    text output generated by this method will be written to
//    parameter 'strBuilder' and returned to the calling
//    function.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of TextFieldSpecLabel are found to be
//       invalid, this method will return an error. If the member
//       data variables are determined to be valid, this error
//       return parameter will be set to 'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (textSpecNanobot *textSpecificationNanobot) buildFormattedSingleParameterText(
	paramLabelLeftMargin string,
	paramLabel string,
	paramLabelFieldLen int,
	paramLabelJustify TextJustify,
	paramLabelRightMargin string,
	paramValueStr string,
	paramValueFieldLen int,
	paramValueJustify TextJustify,
	paramValueRightMargin string,
	lineTerminator string,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if textSpecNanobot.lock == nil {
		textSpecNanobot.lock = new(sync.Mutex)
	}

	textSpecNanobot.lock.Lock()

	defer textSpecNanobot.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textSpecificationNanobot."+
			"buildFormattedSingleParameterText()",
		"")

	if err != nil {

		return err
	}

	if len(paramLabel) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'parameterLabel' is invalid!\n"+
			"'paramLabel' is empty and has a string length of zero (0).\n",
			ePrefix.String())

		return err
	}

	if len(paramValueStr) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'parameterValueStr' is invalid!\n"+
			"'parameterValueStr' is empty and has a string length of zero.\n",
			ePrefix.String())

		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if paramLabelFieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'paramLabelFieldLen' is invalid!\n"+
			"'paramLabelFieldLen' has a value less than minus one (-1).\n"+
			"paramLabelFieldLen = '%v'\n",
			ePrefix.String(),
			paramLabelFieldLen)

		return err

	}

	if paramValueFieldLen < -1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'paramValueFieldLen' is invalid!\n"+
			"'paramValueFieldLen' has a value less than minus one (-1).\n"+
			"paramValueFieldLen = '%v'\n",
			ePrefix.String(),
			paramValueFieldLen)

		return err

	}

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	if len(paramLabelLeftMargin) > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			paramLabelLeftMargin,
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy("stdLine<-"+
				"paramLabelLeftMargin"))

		if err != nil {
			return err
		}

	}

	_,
		err = stdLine.AddTextFieldLabel(
		paramLabel,
		paramLabelFieldLen,
		paramLabelJustify,
		ePrefix.XCpy(
			"stdLine<-paramLabel"))

	if err != nil {

		return err
	}

	if len(paramLabelRightMargin) > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			paramLabelRightMargin,
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy("stdLine<-"+
				"paramLabelRightMargin"))

		if err != nil {
			return err
		}

	}

	_,
		err = stdLine.AddTextFieldLabel(
		paramValueStr,
		paramValueFieldLen,
		paramValueJustify,
		ePrefix.XCpy(
			"stdLine<-paramValueStr"))

	if err != nil {

		return err
	}

	if len(paramValueRightMargin) > 0 {
		_,
			err = stdLine.AddTextFieldLabel(
			paramValueRightMargin,
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy("stdLine<-"+
				"paramValueRightMargin"))

		if err != nil {
			return err
		}

	}

	if len(lineTerminator) > 0 {

		err = stdLine.SetNewLineChars(
			lineTerminator,
			ePrefix.XCpy(
				"stdLine<-lineTerminator"))

		if err != nil {
			return err
		}

	} else {

		// len(lineTerminator) == 0
		stdLine.TurnAutoLineTerminationOff()

	}

	err = stdLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"strBuilder<-stdLine"))

	return err
}

// buildFormattedSingleTextField - Formats a single line of text
// consisting of a single text label.
//
// The formatted single text field output consists of the
// following text fields located on a single text line:
//   [TEXT LABEL LEFT MARGIN]
//   Text Label
//   [TEXT LABEL RIGHT MARGIN]
//   [LINE TERMINATOR]
//
//
// Example Text Output Format:
//
// [" "]["Hello World!"][" "]['\n']
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  textLabelLeftMargin        string
//     - This string defines the text which will be placed on the
//       left side of the 'textLabel' string.
//
//       If no left margin is required, submit this parameter as a
//       zero length or empty string and no left margin will be
//       created.
//
//
//  textLabel                  string
//     - This string contains text which will be formatted as
//       the only text label field inserted in the output string.
//
//       If this parameter is submitted as an empty or zero length
//       string, an error will be returned.
//
//
// 	textLabelFieldLen          int
//     - The length of the text field in which the 'textLabel' will
//       be displayed. If 'textLabelFieldLen' is less than the
//       length of the 'textLabel' string, it will be automatically
//       set equal to the 'textLabel' string length.
//
//       To automatically set the value of 'textLabelFieldLen' to
//       the length of 'textLabel', set this parameter to a value
//       of minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  textLabelJustify           TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' string within the field specified by
//       'textLabelFieldLen'.
//
//       Text justification can only be evaluated in the context of
//       a text label, field length and a Text Justification object
//       of type TextJustify. This is because text labels with a
//       field length equal to or less than the length of the text
//       label never use text justification. In these cases, text
//       justification is completely ignored.
//
//       If the display field length ('paramLabelFieldLen') is
//       greater than the length of the text label ('paramLabel'),
//       text justification must be equal to one of these three
//       valid values:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       You can also use the abbreviated text justification
//       enumeration syntax as follows:
//
//           TxtJustify.Left()
//           TxtJustify.Right()
//           TxtJustify.Center()
//
//
//  textLabelRightMargin       string
//     - The contents of the string will be used as the right
//       margin for the 'textLabel' text string. This parameter is
//       commonly used to specify one or more white space
//       characters positioned to the right of 'textLabel'.
//
//       If no right margin is required, set 'textLabelRightMargin'
//       to a zero length or empty string, and no right margin will
//       be created.
//
//
//  lineTerminator             string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text output string
//       created by this method. The most common usage sets this
//       string to a new line character ("\n").
//
//       If no Line Terminator is required, set 'lineTerminator' to
//       a zero length or empty string and no line termination
//       characters will be created.
//
//
//  strBuilder                *strings.Builder
//  - A pointer to an instance of strings.Builder. The formatted
//    text output generated by this method will be written to
//    parameter 'strBuilder' and returned to the calling
//    function.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of TextFieldSpecLabel are found to be
//       invalid, this method will return an error. If the member
//       data variables are determined to be valid, this error
//       return parameter will be set to 'nil'.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (textSpecNanobot *textSpecificationNanobot) buildFormattedSingleTextField(
	textLabelLeftMargin string,
	textLabel string,
	textLabelFieldLen int,
	textLabelJustify TextJustify,
	textLabelRightMargin string,
	lineTerminator string,
	strBuilder *strings.Builder,
	errPrefDto *ePref.ErrPrefixDto) error {

	if textSpecNanobot.lock == nil {
		textSpecNanobot.lock = new(sync.Mutex)
	}

	textSpecNanobot.lock.Lock()

	defer textSpecNanobot.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textSpecificationNanobot."+
			"buildFormattedSingleTextField()",
		"")

	if err != nil {

		return err
	}

	if len(textLabelLeftMargin) == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'textLabelLeftMargin' is invalid!\n"+
			"'textLabelLeftMargin' is empty and has a string length of zero (0).\n",
			ePrefix.String())

		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if textLabelFieldLen < -1 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'textLabelFieldLen' is invalid!\n"+
			"'textLabelFieldLen' has a value less than minus one (-1).\n"+
			"textLabelFieldLen = '%v'\n",
			ePrefix.String(),
			textLabelFieldLen)

		return err

	}

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	if len(textLabelLeftMargin) > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			textLabelLeftMargin,
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy("stdLine<-"+
				"textLabelLeftMargin"))

		if err != nil {
			return err
		}

	}

	_,
		err = stdLine.AddTextFieldLabel(
		textLabel,
		textLabelFieldLen,
		textLabelJustify,
		ePrefix.XCpy(
			"stdLine<-textLabel"))

	if err != nil {

		return err
	}

	if len(textLabelRightMargin) > 0 {

		_,
			err = stdLine.AddTextFieldLabel(
			textLabelRightMargin,
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy("stdLine<-"+
				"textLabelRightMargin"))

		if err != nil {
			return err
		}

	}

	if len(lineTerminator) > 0 {

		err = stdLine.SetNewLineChars(
			lineTerminator,
			ePrefix.XCpy(
				"stdLine<-lineTerminator"))

		if err != nil {
			return err
		}

	} else {

		// len(lineTerminator) == 0
		stdLine.TurnAutoLineTerminationOff()

	}

	err = stdLine.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"strBuilder<-stdLine"))

	return err
}

// ptr - Returns a pointer to a new instance of
// textSpecificationNanobot.
//
func (textSpecNanobot textSpecificationNanobot) ptr() *textSpecificationNanobot {

	if textSpecNanobot.lock == nil {
		textSpecNanobot.lock = new(sync.Mutex)
	}

	textSpecNanobot.lock.Lock()

	defer textSpecNanobot.lock.Unlock()

	return &textSpecificationNanobot{
		lock: new(sync.Mutex),
	}
}

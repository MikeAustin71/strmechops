package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// textStrBuilderAtom - Provides helper methods for type
// TextStrBuilder.
//
type textStrBuilderAtom struct {
	lock *sync.Mutex
}

// fieldDateTimeWithMargins - Is designed to produce three text
// elements consolidated and formatted as a single text field.
//
// The three text elements consist of a left margin string, a
// date/time text field and a right margin string.
//
// These three text elements can be configured as a complete line
// of text depending on the value applied to input parameter
// 'lineTerminator'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder                 *strings.Builder
//     - A pointer to an instance of strings.Builder. A formatted
//       string of text characters created by this method will be
//       written to this instance of strings.Builder.
//
//
//  leftMarginStr              string
//     - The contents of the string will be used as the left margin
//       for the 'dateTime' field.
//
//       If no left margin is required, set 'LeftMarginStr' to a
//       zero length or empty string, and no left margin will be
//       created.
//
//
//  dateTime                   time.Time
//     - The date/time value which will be formatted as a text
//       string.
//
//       If this parameter is set equal to zero, no error will
//       be generated.
//
//
//  dateTimeFieldLength        int
//     - Used to format Date/Time Text Fields. This is the length
//       of the text field in which the formatted 'dateTime' string
//       will be displayed. If 'dateTimeFieldLength' is less than
//       the length of the 'dateTime' string, it will be
//       automatically set equal to the 'dateTime' string length.
//
//       To automatically set the value of 'dateTimeFieldLength' to
//       the length of 'dateTime', set this parameter to a value of
//       minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  dateTimeFormat             string
//    - This string will be used to format the date/time value
//      'dateTime' as a text string.
//
//       If this 'dateTimeFormat' string is empty (has a zero
//       length), a default Date/Time format string will be applied
//       as follows:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  dateTimeTextJustify        TextJustify
//      An enumeration value specifying the justification of the
//      'dateTime' string within the text field specified by
//      'dateTimeFieldLength'.
//
//      Text justification can only be evaluated in the context of
//      a text label, field length and a Text Justification object
//      of type TextJustify. This is because text labels with a
//      field length equal to or less than the length of the text
//      label never use text justification. In these cases, text
//      justification is completely ignored.
//
//      If the field length is greater than the length of the text
//      label, text justification must be equal to one of these
//      three valid values:
//          TextJustify(0).Left()
//          TextJustify(0).Right()
//          TextJustify(0).Center()
//
//      You can also use the abbreviated text justification
//      enumeration syntax as follows:
//
//          TxtJustify.Left()
//          TxtJustify.Right()
//          TxtJustify.Center()
//
//
//  rightMarginStr             string
//     - The contents of the string will be used as the right
//       margin for the 'dateTime' field.
//
//       If no right margin is required, set 'RightMarginStr' to a
//       zero length or empty string, and no right margin will be
//       created.
//
//
//  lineTerminator             string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text thereby converting
//       this text element into a valid line of text.
//
//       If a text line is required, setting this string to include
//       a new line character ('\n') will ensure that the three
//       text elements formatted by this method as single text
//       field will constitute a single line of text.
//
//       The most common usage sets this string to a new line
//       character ("\n").
//
//       If Line Termination is NOT required, set 'lineTerminator'
//       to a zero length or empty string and no line termination
//       characters will be created.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtBuilderAtom *textStrBuilderAtom) fieldDateTimeWithMargins(
	strBuilder *strings.Builder,
	leftMarginStr string,
	dateTime time.Time,
	dateTimeFieldLength int,
	dateTimeFormat string,
	dateTimeTextJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldLabelWithMargins()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' has a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(dateTimeFormat) == 0 {
		dateTimeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultDateTimeFormat()
	}

	if len(leftMarginStr) > 0 {
		strBuilder.WriteString(leftMarginStr)
	}

	var txtDateTimeField TextFieldSpecDateTime

	txtDateTimeField,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		dateTimeFieldLength,
		dateTimeFormat,
		dateTimeTextJustify,
		ePrefix.XCpy(
			"txtDateTimeField<-dateTime"))

	if err != nil {
		return err
	}

	err = txtDateTimeField.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"strBuilder<-txtDateTimeField"))

	if err != nil {
		return err
	}

	if len(rightMarginStr) > 0 {
		strBuilder.WriteString(rightMarginStr)
	}

	if len(lineTerminator) > 0 {
		strBuilder.WriteString(lineTerminator)
	}

	return err
}

// FieldsSingleFiller - Designed to produce three text elements
// consolidated and formatted as a single text field.
//
// The three text elements consist of a left margin string, a Text
// Filler Field and a right margin string.
//
// These three text elements can be configured as a complete line
// of text depending on the value applied to input parameter
// 'lineTerminator'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder                 *strings.Builder
//     - A pointer to an instance of strings.Builder. A formatted
//       string of text characters created by this method will be
//       written to this instance of strings.Builder.
//
//
//  leftMarginStr              string
//     - The contents of the string will be used as the left margin
//       for 'labelText field.
//
//       If no left margin is required, set 'LeftMarginStr' to a
//       zero length or empty string, and no left margin will be
//       created.
//
//
//  fillerCharacters           string
//     - A string containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from the filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharacters' is submitted as an empty or zero
//       length string, it will be defaulted to a single white
//       space character.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
//
//
//  rightMarginStr             string
//     - The contents of the string will be used as the right
//       margin for the Text Filler Field.
//
//       If no right margin is required, set 'RightMarginStr' to a
//       zero length or empty string, and no right margin will be
//       created.
//
//
//  lineTerminator             string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text thereby converting
//       this text element into a valid line of text.
//
//       If a text line is required, setting this string to include
//       a new line character ('\n') will ensure that the text line
//       consists of the text label field and no other text
//       elements.
//
//       The most common usage sets this string to a new line
//       character ("\n").
//
//       If Line Termination is NOT required, set 'lineTerminator'
//       to a zero length or empty string and no line termination
//       characters will be created.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtBuilderAtom *textStrBuilderAtom) fieldFillerWithMargins(
	strBuilder *strings.Builder,
	leftMarginStr string,
	fillerCharacters string,
	fillerCharsRepeatCount int,
	rightMarginStr string,
	lineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldFillerWithMargins()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' has a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(fillerCharacters) == 0 {

		fillerCharacters = " "

	}

	if len(leftMarginStr) > 0 {
		strBuilder.WriteString(leftMarginStr)
	}

	var txtFillerFieldSpec TextFieldSpecFiller

	txtFillerFieldSpec,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerCharacters,
		fillerCharsRepeatCount,
		ePrefix.XCpy(
			"txtFillerFieldSpec"))

	if err != nil {
		return err
	}

	err = txtFillerFieldSpec.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"strBuilder<-txtFillerFieldSpec"))

	if err != nil {
		return err
	}

	if len(rightMarginStr) > 0 {
		strBuilder.WriteString(leftMarginStr)
	}

	if len(lineTerminator) > 0 {
		strBuilder.WriteString(lineTerminator)
	}

	return err
}

// fieldLabelWithMargins - Formats a single text label and writes
// the output string to an instance of strings.Builder passed as an
// input parameter by the calling function.
//
// If the Left and Right Margin Strings contain characters, they
// will also be written to the strings.Builder instance
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder                 *strings.Builder
//     - A pointer to an instance of strings.Builder. A formatted
//       text label string created by this method will be written
//       to this instance of strings.Builder.
//
//
//  leftMarginStr              string
//     - The contents of the string will be used as the left margin
//       for 'labelText field.
//
//       If no left margin is required, set 'LeftMarginStr' to a
//       zero length or empty string, and no left margin will be
//       created.
//
//
//  labelText                  string
//     - This strings holds the text characters which will be
//       formatted as a text label.
//
//       If 'labelText' is submitted as a zero length or empty
//       string it will automatically be defaulted to a single
//       white space character, " ".
//
//
//  labelFieldLength           int
//     - Used to format Text Label Fields. This is the length of
//       the text field in which the formatted 'labelText' string
//       will be displayed. If 'labelFieldLength' is less than the
//       length of the 'labelText' string, it will be automatically
//       set equal to the 'labelText' string length.
//
//       To automatically set the value of 'labelFieldLength' to
//       the length of 'labelText', set this parameter to a value
//       of  minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  labelTextJustify           TextJustify
//      An enumeration value specifying the justification of the
//      'labelText' string within the text field specified by
//      'labelFieldLength'.
//
//      Text justification can only be evaluated in the context of
//      a text label, field length and a Text Justification object
//      of type TextJustify. This is because text labels with a
//      field length equal to or less than the length of the text
//      label never use text justification. In these cases, text
//      justification is completely ignored.
//
//      If the field length is greater than the length of the text
//      label, text justification must be equal to one of these
//      three valid values:
//          TextJustify(0).Left()
//          TextJustify(0).Right()
//          TextJustify(0).Center()
//
//      You can also use the abbreviated text justification
//      enumeration syntax as follows:
//
//          TxtJustify.Left()
//          TxtJustify.Right()
//          TxtJustify.Center()
//
//
//  rightMarginStr             string
//     - The contents of the string will be used as the right
//       margin for the 'labelText' field.
//
//       If no right margin is required, set 'RightMarginStr' to a
//       zero length or empty string, and no right margin will be
//       created.
//
//
//  lineTerminator             string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text thereby converting
//       this text element into a valid line of text.
//
//       If a text line is required, setting this string to include
//       a new line character ('\n') will ensure that the text line
//       consists of the text label field and no other text
//       elements.
//
//       The most common usage sets this string to a new line
//       character ("\n").
//
//       If Line Termination is NOT required, set 'lineTerminator'
//       to a zero length or empty string and no line termination
//       characters will be created.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtBuilderAtom *textStrBuilderAtom) fieldLabelWithMargins(
	strBuilder *strings.Builder,
	leftMarginStr string,
	labelText string,
	labelFieldLength int,
	labelTextJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldLabelWithMargins()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' has a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	if len(labelText) == 0 {

		labelText = " "

	}

	if len(leftMarginStr) > 0 {
		strBuilder.WriteString(leftMarginStr)
	}

	var txtLabelSpec TextFieldSpecLabel

	txtLabelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		labelText,
		labelFieldLength,
		labelTextJustify,
		ePrefix.XCpy(
			"txtLabelSpec<-labelText"))

	if err != nil {
		return err
	}

	err = txtLabelSpec.TextBuilder(
		strBuilder,
		ePrefix.XCpy(
			"strBuilder<-txtLabelSpec"))

	if err != nil {
		return err
	}

	if len(rightMarginStr) > 0 {
		strBuilder.WriteString(rightMarginStr)
	}

	if len(lineTerminator) > 0 {
		strBuilder.WriteString(lineTerminator)
	}

	return err
}

// ptr - Returns a pointer to a new instance of
// textStrBuilderAtom.
//
func (txtBuilderAtom textStrBuilderAtom) ptr() *textStrBuilderAtom {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	return &textStrBuilderAtom{
		lock: new(sync.Mutex),
	}
}

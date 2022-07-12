package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// textStrBuilderElectron - Provides helper methods for type
// TextStrBuilder.
//
type textStrBuilderElectron struct {
	lock *sync.Mutex
}

// fieldSpacerWithMargins - Creates a string consisting of white
// space characters (" "). The number of white space characters
// included in the text field is determined by input parameter,
// 'fieldLength'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder                 strings.Builder
//     - An instance of strings.Builder. A formatted text label
//       string created by this method will be written to this
//       instance of strings.Builder.
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
//  fieldLength                int
//     - An integer value greater than zero and less than 1,000,001
//       which is used to specify the number of white space
//       characters in the Text Spacer Field.
//
//       Examples:
//          fieldLen = 1 produces text field " "
//          fieldLen = 2 produces text field "  "
//          fieldLen = 5 produces text field "     "
//
//        If the value of 'fieldLength' is less than 1 or greater
//        than one-million, an error will be returned.
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
func (txtBuilderElectron *textStrBuilderElectron) fieldSpacerWithMargins(
	strBuilder strings.Builder,
	leftMarginStr string,
	fieldLength int,
	rightMarginStr string,
	lineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtBuilderElectron.lock == nil {
		txtBuilderElectron.lock = new(sync.Mutex)
	}

	txtBuilderElectron.lock.Lock()

	defer txtBuilderElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderElectron."+
			"lineBlank()",
		"")

	if err != nil {
		return err
	}

	if fieldLength < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLength' is invalid!\n"+
			"'fieldLength' has a value less than one (+1).\n"+
			"fieldLength = '%v'\n",
			ePrefix.String(),
			fieldLength)

		return err
	}

	if fieldLength > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fieldLength' is invalid!\n"+
			"'fieldLength' has a value greater than one-million (1,000,000).\n"+
			"fieldLength = '%v'\n",
			ePrefix.String(),
			fieldLength)

		return err
	}

	if len(leftMarginStr) > 0 {
		strBuilder.WriteString(leftMarginStr)
	}

	strBuilder.WriteString(strings.Repeat(" ", fieldLength))

	if len(rightMarginStr) > 0 {
		strBuilder.WriteString(rightMarginStr)
	}

	if len(lineTerminator) > 0 {
		strBuilder.WriteString(lineTerminator)
	}

	return err
}

// lineBlank - Designed to produce one or more blank or empty lines
// of text. Each blank line will consist of a single new line
// character, '/n'.
//
// Consequently, the terms "blank lines" and "new lines" are
// synonymous as used here.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder                 strings.Builder
//     - An instance of strings.Builder. A formatted string of text
//       characters created by this method will be written to this
//       instance of strings.Builder.
//
//
//  numOfBlankLines            int
//     - The number of blank lines which will be generated by this
//       method.
//
//       If this value is less than one (+1), an error will be
//       returned. Likewise, if this value is greater than
//       one-million (1,000,000), an error will be returned.
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
//  err                             error
//     - If this method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtBuilderElectron *textStrBuilderElectron) lineBlank(
	strBuilder strings.Builder,
	numOfBlankLines int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtBuilderElectron.lock == nil {
		txtBuilderElectron.lock = new(sync.Mutex)
	}

	txtBuilderElectron.lock.Lock()

	defer txtBuilderElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderElectron."+
			"lineBlank()",
		"")

	if err != nil {
		return err
	}

	if numOfBlankLines < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfBlankLines' is invalid!\n"+
			"'numOfBlankLines' has a value less than one (+1).\n"+
			"numOfBlankLines = '%v'\n",
			ePrefix.String(),
			numOfBlankLines)

		return err
	}

	if numOfBlankLines > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfBlankLines' is invalid!\n"+
			"'numOfBlankLines' has a value greater than one-million (1,000,000).\n"+
			"numOfBlankLines = '%v'\n",
			ePrefix.String(),
			numOfBlankLines)

		return err
	}

	for i := 0; i < numOfBlankLines; i++ {

		strBuilder.WriteString("\n")

	}

	return err
}

// ptr - Returns a pointer to a new instance of
// textStrBuilderElectron.
//
func (txtBuilderElectron textStrBuilderElectron) ptr() *textStrBuilderElectron {

	if txtBuilderElectron.lock == nil {
		txtBuilderElectron.lock = new(sync.Mutex)
	}

	txtBuilderElectron.lock.Lock()

	defer txtBuilderElectron.lock.Unlock()

	return &textStrBuilderElectron{
		lock: new(sync.Mutex),
	}
}

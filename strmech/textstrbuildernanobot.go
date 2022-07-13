package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// textStrBuilderNanobot - Provides helper methods for type
// TextStrBuilder.
//
type textStrBuilderNanobot struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// textStrBuilderNanobot.
//
func (txtBuilderNanobot textStrBuilderNanobot) ptr() *textStrBuilderNanobot {

	if txtBuilderNanobot.lock == nil {
		txtBuilderNanobot.lock = new(sync.Mutex)
	}

	txtBuilderNanobot.lock.Lock()

	defer txtBuilderNanobot.lock.Unlock()

	return &textStrBuilderNanobot{
		lock: new(sync.Mutex),
	}
}

// LineSolid - Designed to produce one or more separate lines of
// text.
//
// Each line consists of three text elements: a left margin
// string, a Text Filler Field, and a right margin strings.
//
// These three text elements can be configured as independent
// lines of text or concatenated together depending on the value
// applied to input parameters 'interiorLineTerminator' and
// 'finalLineTerminator'.
//
// This method is similar to method:
//   TextStrBuilder.FieldsSingleFiller()
//
// However, this method is capable of producing multiple lines
// of filler text.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  strBuilder                 strings.Builder
//     - An instance of strings.Builder. A formatted string of
//       text characters created by this method will be written to
//       this instance of strings.Builder.
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
//       length string, this method will automatically set
//       'fillerCharacters' to a single white space character,
//       (" ").
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
//  interiorLineTerminator     string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text thereby converting
//       this text element into a valid line of text.
//
//       Be sure to coordinate 'interiorLineTerminator' with input
//       parameter 'finalLineTerminator'. 'interiorLineTerminator'
//       is applied after each line of text is generated.
//       'finalLineTerminator' is applied after all lines of text
//       have been generated.
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
//  numOfLines                 int
//     - The number of times the combination of left margin string,
//       Text Filler Field, right margin string and interior line
//       terminator string will be repeated.
//
//       Essentially, this the repeat count for the Text Filler
//       Lines.
//
//       If this value is less than one (+1), an error will be
//       returned. Likewise, if this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  finalLineTerminator        string
//     - After all the text lines have been generated according to
//       input parameter 'numOfLines', this line termination
//       sequence will be applied to the final text output string.
//
//       Be sure to coordinate 'finalLineTerminator' with input
//       parameter 'interiorLineTerminator'.
//
//       'interiorLineTerminator' is applied after each line of
//       text is generated. 'finalLineTerminator' is applied after
//       all lines of text have been generated.
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
func (txtBuilderNanobot *textStrBuilderNanobot) lineSolidWithMargins(
	leftMarginStr string,
	fillerCharacters string,
	fillerCharsRepeatCount int,
	rightMarginStr string,
	interiorLineTerminator string,
	numOfLines int,
	finalLineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if txtBuilderNanobot.lock == nil {
		txtBuilderNanobot.lock = new(sync.Mutex)
	}

	txtBuilderNanobot.lock.Lock()

	defer txtBuilderNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var strBuilder strings.Builder

	strBuilder.Grow(128)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldLabelWithMargins()",
		"")

	if err != nil {
		return strBuilder, err
	}

	if numOfLines < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfLines' is invalid!\n"+
			"'numOfLines' has a value less than one (+1).\n"+
			"numOfLines = '%v'\n",
			ePrefix.String(),
			numOfLines)

		return strBuilder, err
	}

	if numOfLines > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfLines' is invalid!\n"+
			"'numOfLines' has a value greater than one-million (1,000,000).\n"+
			"numOfLines = '%v'\n",
			ePrefix.String(),
			numOfLines)

		return strBuilder, err
	}

	txtBuilderAtom := textStrBuilderAtom{}
	var strBuilder2 strings.Builder

	for i := 0; i < numOfLines; i++ {

		strBuilder2,
			err = txtBuilderAtom.fieldFillerWithMargins(
			leftMarginStr,
			fillerCharacters,
			fillerCharsRepeatCount,
			rightMarginStr,
			interiorLineTerminator,
			ePrefix.XCpy(
				fmt.Sprintf(
					"strBuilder<-fillerCharacters[%v]",
					i)))

		if err != nil {
			return strBuilder, err
		}

		strBuilder.WriteString(strBuilder2.String())

		strBuilder2.Reset()
	}

	if len(finalLineTerminator) > 0 {
		strBuilder.WriteString(finalLineTerminator)
	}

	return strBuilder, err
}

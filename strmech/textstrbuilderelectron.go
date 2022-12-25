package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// textStrBuilderElectron - Provides helper methods for type
// TextStrBuilder.
type textStrBuilderElectron struct {
	lock *sync.Mutex
}

// writeLeftMargin - Writes left margin string to the formatted
// text stream.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strBuilder                 *strings.Builder
//	   - A pointer to an instance of strings.Builder. The left
//	     margin characters will be written to this string builder.
//
//
//	maxLineLength              int
//	   - The total length of the output line.
//
//
//	currentLineLength          int
//	   - The number of characters previously written to the current
//	     line.
//
//
//	leftMarginStr              string
//	   - The left margin string which will be written to
//	     'strBuilder'.
//
//
//	lineTerminatorStr          string
//	   - The string of characters used to terminate a line of text.
//
//
//	turnAutoLineBreaksOn       bool
//	- When this parameter is set to 'true' it signals that
//	  automatic line breaks are engaged and will be applied. This
//	  means that if text exceeds the maximum line length, line
//	  termination characters will be automatically inserted.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newCurrentLineLength       int
//	   - The number of characters written to the current output
//	     line.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtBuilderElectron *textStrBuilderElectron) writeLeftMargin(
	strBuilder *strings.Builder,
	maxLineLength int,
	currentLineLength int,
	leftMarginStr string,
	lineTerminatorStr string,
	turnAutoLineBreaksOn bool,
	errPrefDto *ePref.ErrPrefixDto) (
	newCurrentLineLength int,
	err error) {

	if txtBuilderElectron.lock == nil {
		txtBuilderElectron.lock = new(sync.Mutex)
	}

	txtBuilderElectron.lock.Lock()

	defer txtBuilderElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderElectron."+
			"writeLeftMargin()",
		"")

	if err != nil {
		return newCurrentLineLength, err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return newCurrentLineLength, err
	}

	lenLeftMarginStr := len(leftMarginStr)

	if lenLeftMarginStr == 0 {

		newCurrentLineLength = currentLineLength
		// Nothing to do. There is no Left Margin
		// string.
		return newCurrentLineLength, err
	}

	lenLineTerminatorStr := len(lineTerminatorStr)

	if turnAutoLineBreaksOn == true &&
		lenLineTerminatorStr == 0 {

		lineTerminatorStr = "\n"
		lenLineTerminatorStr = 1
	}

	if turnAutoLineBreaksOn == false {
		// No Line Breaks required.
		// Just write the left margin string
		strBuilder.WriteString(leftMarginStr)

		newCurrentLineLength = lenLeftMarginStr +
			currentLineLength

		return newCurrentLineLength, err

	}

	// MUST BE
	// turnAutoLineBreaksOn == true
	// lenLineTerminatorStr > 0

	netMaximumLength := maxLineLength -
		currentLineLength -
		lenLineTerminatorStr

	if netMaximumLength < 0 {
		// Deal with this error condition
		strBuilder.WriteString(leftMarginStr + lineTerminatorStr)

		newCurrentLineLength = 0

		return newCurrentLineLength, err

	}

	if lenLeftMarginStr < netMaximumLength {

		strBuilder.WriteString(leftMarginStr)

		newCurrentLineLength = lenLeftMarginStr + currentLineLength

		return newCurrentLineLength, err

	}

	if lenLeftMarginStr == netMaximumLength {

		strBuilder.WriteString(leftMarginStr + lineTerminatorStr)

		newCurrentLineLength = 0

		return newCurrentLineLength, err

	}

	// MUST BE
	// lenLeftMarginStr > netMaximumLength

	sMecElectron := strMechElectron{}
	var writeLeftMarginStr string

	for lenLeftMarginStr > 0 {

		if lenLeftMarginStr > netMaximumLength {

			writeLeftMarginStr,
				leftMarginStr,
				lenLeftMarginStr,
				err =
				sMecElectron.cutStringAtIndex(
					leftMarginStr,
					netMaximumLength+1,
					ePrefix.XCpy(
						"Cut leftMarginStr"))

			if err != nil {

				newCurrentLineLength = 0
				return newCurrentLineLength, err
			}

			strBuilder.WriteString(
				writeLeftMarginStr + lineTerminatorStr)

			netMaximumLength = maxLineLength - lenLineTerminatorStr

			currentLineLength = 0

			continue

		} else if lenLeftMarginStr == netMaximumLength {

			strBuilder.WriteString(leftMarginStr + lineTerminatorStr)

			newCurrentLineLength = 0

			return newCurrentLineLength, err

		} else {
			// MUST BE
			// lenLeftMarginStr < netMaximumLength

			strBuilder.WriteString(leftMarginStr)

			newCurrentLineLength = lenLeftMarginStr

			return newCurrentLineLength, err

		}

	}

	newCurrentLineLength = 0
	return newCurrentLineLength, err
}

// writeRightMargin - Writes right margin string to the formatted
// text stream.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strBuilder                 *strings.Builder
//	   - A pointer to an instance of strings.Builder. The right
//	     margin characters will be written to this string builder.
//
//
//	maxLineLength              int
//	   - The total length of the output line.
//
//
//	currentLineLength          int
//	   - The number of characters previously written to the current
//	     line.
//
//
//	rightMarginStr             string
//	   - The right margin string which will be written to
//	     'strBuilder'.
//
//
//	lineTerminatorStr          string
//	   - The string of characters used to terminate a line of text.
//
//
//	turnAutoLineBreaksOn       bool
//	- When this parameter is set to 'true' it signals that
//	  automatic line breaks are engaged and will be applied. This
//	  means that if text exceeds the maximum line length, line
//	  termination characters will be automatically inserted.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtBuilderElectron *textStrBuilderElectron) writeRightMargin(
	strBuilder *strings.Builder,
	maxLineLength int,
	currentLineLength int,
	rightMarginStr string,
	lineTerminatorStr string,
	turnAutoLineBreaksOn bool,
	errPrefDto *ePref.ErrPrefixDto) (
	newCurrentLineLength int,
	lastWriteWasLineTerminator bool,
	err error) {

	if txtBuilderElectron.lock == nil {
		txtBuilderElectron.lock = new(sync.Mutex)
	}

	txtBuilderElectron.lock.Lock()

	defer txtBuilderElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lastWriteWasLineTerminator = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderElectron."+
			"writeRightMargin()",
		"")

	if err != nil {

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	lenRightMarginStr := len(rightMarginStr)

	if lenRightMarginStr == 0 {

		newCurrentLineLength = currentLineLength
		// Nothing to do. There is no Right Margin
		// string.

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	lenLineTerminatorStr := len(lineTerminatorStr)

	if turnAutoLineBreaksOn == true &&
		lenLineTerminatorStr == 0 {

		lineTerminatorStr = "\n"
		lenLineTerminatorStr = 1
	}

	if turnAutoLineBreaksOn == false {
		// No Line Breaks required.
		// Just write the right margin string
		strBuilder.WriteString(rightMarginStr)

		newCurrentLineLength =
			lenRightMarginStr +
				currentLineLength

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err

	}

	// MUST BE
	// turnAutoLineBreaksOn == true
	// lenLineTerminatorStr > 0

	netMaximumLength := maxLineLength -
		currentLineLength -
		lenLineTerminatorStr

	if netMaximumLength < 0 {
		// Deal with this error condition
		strBuilder.WriteString(rightMarginStr + lineTerminatorStr)

		newCurrentLineLength = 0

		lastWriteWasLineTerminator = true

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	if lenRightMarginStr < netMaximumLength {

		strBuilder.WriteString(rightMarginStr)

		newCurrentLineLength = lenRightMarginStr + currentLineLength

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	if lenRightMarginStr == netMaximumLength {

		strBuilder.WriteString(rightMarginStr + lineTerminatorStr)

		newCurrentLineLength = 0

		lastWriteWasLineTerminator = true

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	// MUST BE
	// lenRightMarginStr > netMaximumLength

	sMecElectron := strMechElectron{}
	var writeRightMarginStr string

	for lenRightMarginStr > 0 {

		if lenRightMarginStr > netMaximumLength {

			writeRightMarginStr,
				rightMarginStr,
				lenRightMarginStr,
				err =
				sMecElectron.cutStringAtIndex(
					rightMarginStr,
					netMaximumLength+1,
					ePrefix.XCpy(
						"Cut rightMarginStr"))

			if err != nil {

				newCurrentLineLength = 0

				return newCurrentLineLength,
					lastWriteWasLineTerminator,
					err
			}

			strBuilder.WriteString(
				writeRightMarginStr + lineTerminatorStr)

			netMaximumLength = maxLineLength - lenLineTerminatorStr

			currentLineLength = 0

			lastWriteWasLineTerminator = true

			continue

		} else if lenRightMarginStr == netMaximumLength {

			strBuilder.WriteString(rightMarginStr + lineTerminatorStr)

			newCurrentLineLength = 0

			lastWriteWasLineTerminator = true

			return newCurrentLineLength,
				lastWriteWasLineTerminator,
				err

		} else {
			// MUST BE
			// lenRightMarginStr < netMaximumLength

			strBuilder.WriteString(rightMarginStr)

			newCurrentLineLength = lenRightMarginStr

			lastWriteWasLineTerminator = false

			return newCurrentLineLength,
				lastWriteWasLineTerminator,
				err
		}

	} // for lenRightMarginStr > 0

	newCurrentLineLength = 0

	return newCurrentLineLength,
		lastWriteWasLineTerminator,
		err

}

// writeText - Writes the text string to the formatted text
// stream.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strBuilder                 *strings.Builder
//	   - A pointer to an instance of strings.Builder. The text
//	     characters will be written to this string builder.
//
//
//	maxLineLength              int
//	   - The total length of the output line.
//
//	     It is assumed that maxLineLength has been previously
//	     validated.
//
//
//	currentLineLength          int
//	   - The number of characters previously written to the current
//	     line.
//
//
//	leftMarginStr              string
//	   - The left margin string associated with this text.
//	     'leftMarginStr' is optional and may be passed as an empty
//	     string.
//
//
//	textStr                    string
//	   - The text string which will be written to 'strBuilder'.
//
//
//	lineTerminatorStr          string
//	   - The string of characters used to terminate a line of text.
//
//
//	rightMarginStr             string
//	   - The right margin string associated with this line of text.
//	     'rightMarginStr' is optional and may be passed as an empty
//	     string.
//
//
//	turnAutoLineBreaksOn       bool
//	- When this parameter is set to 'true' it signals that
//	  automatic line breaks are engaged and will be applied. This
//	  means that if text exceeds the maximum line length, line
//	  termination characters will be automatically inserted.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	newCurrentLineLength       int
//	   - The number of characters written to the current output
//	     line.
//
//
//	lastWriteWasLineTerminator bool
//	   - When set to 'true', this signals that the last characters
//	     written to 'strBuilder' were line termination characters.
//
//
//	err                        error
//	   - If this method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (txtBuilderElectron *textStrBuilderElectron) writeText(
	strBuilder *strings.Builder,
	maxLineLength int,
	currentLineLength int,
	leftMarginStr string,
	textStr string,
	rightMarginStr string,
	lineTerminatorStr string,
	turnAutoLineBreaksOn bool,
	errPrefDto *ePref.ErrPrefixDto) (
	newCurrentLineLength int,
	lastWriteWasLineTerminator bool,
	err error) {

	if txtBuilderElectron.lock == nil {
		txtBuilderElectron.lock = new(sync.Mutex)
	}

	txtBuilderElectron.lock.Lock()

	defer txtBuilderElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	lastWriteWasLineTerminator = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderElectron."+
			"writeText()",
		"")

	if err != nil {

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	lenTextStr := len(textStr)

	if lenTextStr == 0 {

		newCurrentLineLength = currentLineLength
		// Nothing to do. There is no Text string.

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	lenLineTerminatorStr := len(lineTerminatorStr)

	if turnAutoLineBreaksOn == true &&
		lenLineTerminatorStr == 0 {

		lineTerminatorStr = "\n"
		lenLineTerminatorStr = 1
	}

	if turnAutoLineBreaksOn == false {
		// No Line Breaks required.
		// Just write the right margin string
		strBuilder.WriteString(textStr)

		newCurrentLineLength =
			lenTextStr +
				currentLineLength

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err

	}

	// MUST BE
	// turnAutoLineBreaksOn == true
	// lenLineTerminatorStr > 0

	currLineNetMaxLen := maxLineLength -
		currentLineLength -
		lenLineTerminatorStr

	nextLineNetMaxLen := maxLineLength -
		lenLineTerminatorStr

	lenLeftMarginStr := len(leftMarginStr)
	lenRightMarginStr := len(rightMarginStr)

	longLineMaxLen := maxLineLength -
		lenLineTerminatorStr -
		lenLeftMarginStr -
		lenRightMarginStr

	if currLineNetMaxLen < 0 {
		// Deal with this error condition
		strBuilder.WriteString(textStr + lineTerminatorStr)

		newCurrentLineLength = 0

		lastWriteWasLineTerminator = true

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	isTextStrEmptyOrWhiteSpace :=
		new(strMechQuark).isEmptyOrWhiteSpace(
			textStr)

	if lenTextStr > (currLineNetMaxLen+nextLineNetMaxLen) &&
		isTextStrEmptyOrWhiteSpace == false {

		strBuilder.WriteString(
			lineTerminatorStr +
				leftMarginStr)

		textStr =
			strings.Replace(
				textStr,
				"\n",
				"",
				-1)

		textStr,
			err =
			new(strMechAtom).breakTextAtLineLength(
				textStr,
				longLineMaxLen,
				'\n',
				ePrefix.XCpy(
					fmt.Sprintf("textStr nextLineNetMaxLen='%v'",
						nextLineNetMaxLen)))

		if err != nil {

			return newCurrentLineLength,
				lastWriteWasLineTerminator,
				err
		}

		textStr = strings.Replace(
			textStr,
			"\n",
			rightMarginStr+"\n"+leftMarginStr,
			-1)

		strBuilder.WriteString(
			textStr +
				lineTerminatorStr)

		newCurrentLineLength = 0
		lastWriteWasLineTerminator = true

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err

	}

	if lenTextStr < currLineNetMaxLen {

		strBuilder.WriteString(textStr)

		newCurrentLineLength = lenTextStr + currentLineLength

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	if lenTextStr == currLineNetMaxLen {

		strBuilder.WriteString(textStr + lineTerminatorStr)

		newCurrentLineLength = 0

		lastWriteWasLineTerminator = true

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err
	}

	// MUST BE
	// lenTextStr > currLineNetMaxLen &&
	// lenTextStr <= (currLineNetMaxLen+nextLineNetMaxLen)

	maxEffectiveLineLen :=
		maxLineLength - lenLineTerminatorStr

	if lenTextStr > currLineNetMaxLen &&
		lenTextStr < maxEffectiveLineLen {

		// If text fits on a single line,
		// just terminate the current partial line
		// and place on a separate line below.

		strBuilder.WriteString(lineTerminatorStr)

		strBuilder.WriteString(textStr)

		lastWriteWasLineTerminator = false

		newCurrentLineLength = lenTextStr

		return newCurrentLineLength,
			lastWriteWasLineTerminator,
			err

	}

	sMecElectron := strMechElectron{}
	var writeTextStr string

	// MUST BE lenTextStr > currLineNetMaxLen

	for lenTextStr > 0 {

		if lenTextStr > currLineNetMaxLen {

			writeTextStr,
				textStr,
				lenTextStr,
				err =
				sMecElectron.cutStringAtIndex(
					textStr,
					currLineNetMaxLen,
					ePrefix.XCpy(
						"Cut textStr"))

			if err != nil {

				newCurrentLineLength = 0

				return newCurrentLineLength,
					lastWriteWasLineTerminator,
					err
			}

			strBuilder.WriteString(
				writeTextStr + lineTerminatorStr)

			currLineNetMaxLen = maxLineLength - lenLineTerminatorStr

			lastWriteWasLineTerminator = true

			currentLineLength = 0

			continue

		} else if lenTextStr == currLineNetMaxLen {

			strBuilder.WriteString(textStr + lineTerminatorStr)

			newCurrentLineLength = 0

			lastWriteWasLineTerminator = true

			return newCurrentLineLength,
				lastWriteWasLineTerminator,
				err

		} else {

			// MUST BE
			// lenTextStr < currLineNetMaxLen

			strBuilder.WriteString(textStr)

			newCurrentLineLength = lenTextStr

			lastWriteWasLineTerminator = false

			return newCurrentLineLength,
				lastWriteWasLineTerminator,
				err
		}

	} // for lenTextStr > 0

	return newCurrentLineLength,
		lastWriteWasLineTerminator,
		err
}

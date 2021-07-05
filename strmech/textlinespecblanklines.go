package strmech

import "sync"

// TextLineSpecBlankLines - This type is a specialized form of
// text line specification which is used to create one or more
// blank lines of text.
//
// By default, the blank line(s) are terminated with a standard
// new line character '\n'. However, users have the option to
// substitute an array of runes and apply any character or group
// of characters to terminate the line.
//
type TextLineSpecBlankLines struct {
	numBlankLines int
	newLineChars  []rune
	lock          *sync.Mutex
}

// CopyOut - Returns a deep copy of the current
// TextLineSpecBlankLines instance.
//
func (blkLines *TextLineSpecBlankLines) CopyOut() TextLineSpecBlankLines {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	if len(blkLines.newLineChars) == 0 {
		blkLines.newLineChars = []rune{'\n'}
	}

	newBlankLinesSpec := TextLineSpecBlankLines{}

	newBlankLinesSpec.numBlankLines = blkLines.numBlankLines

	lenBlkLineChars := len(blkLines.newLineChars)

	newBlankLinesSpec.newLineChars = make([]rune, lenBlkLineChars)

	copy(newBlankLinesSpec.newLineChars,
		blkLines.newLineChars)

	return newBlankLinesSpec
}

// GetLineTerminationChars - Returns the Line Termination character
// or characters configured for this current instance of
// TextLineSpecBlankLines.
//
func (blkLines *TextLineSpecBlankLines) GetLineTerminationChars() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	if len(blkLines.newLineChars) == 0 {
		blkLines.newLineChars = []rune{'\n'}
	}

	return string(blkLines.newLineChars)
}

// GetNumOfBlankLines - Returns an integer value signifying the
// number of blank lines which will be produced by this
// TextLineSpecBlankLines instance.
//
func (blkLines *TextLineSpecBlankLines) GetNumOfBlankLines() int {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	if len(blkLines.newLineChars) == 0 {
		blkLines.newLineChars = []rune{'\n'}
	}

	return blkLines.numBlankLines
}

// NewPtr - Returns a pointer to a new instance of
// TextLineSpecBlankLines. The number of blank lines configured in
// this instance will be taken from input parameter,
// 'numOfBlankLines'.
//
// By default, the blank line(s) are terminated with a standard
// new line character '\n'. However, users have the option to
// substitute an array of runes and apply any character or group
// of characters for the line terminating character.
//
// This method will apply the default line termination character,
// '\n'. To apply a different line termination sequence, use the
// method:
//  TextLineSpecBlankLines.SetLineTermination()
//
func (blkLines TextLineSpecBlankLines) NewPtr(
	numOfBlankLines int) *TextLineSpecBlankLines {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	newBlankLinesSpec := TextLineSpecBlankLines{}

	newBlankLinesSpec.numBlankLines = numOfBlankLines

	newBlankLinesSpec.newLineChars = []rune{'\n'}

	return &newBlankLinesSpec
}

// SetLineTermination - By default, the line termination character
// for blank line produced by this text line specification is the
// the new line character, '\n'. However, users have the option
// of substituting and character or series of characters for the
// the line termination sequence.
//
// This method will receive a string as input and apply the
// characters in that string as the line termination sequence for
// this instance of TextLineSpecBlankLines.
//
// If input parameter 'lineTerminationChars' is submitted as an
// empty string, this method will take no action and exit.
//
func (blkLines *TextLineSpecBlankLines) SetLineTermination(
	lineTerminationChars string) {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	if len(blkLines.newLineChars) == 0 {
		blkLines.newLineChars = []rune{'\n'}
	}

	if len(lineTerminationChars) == 0 {
		return
	}

	blkLines.newLineChars = []rune(lineTerminationChars)
}

// SetNumberOfBlankLines - Sets the number of blank lines produced
// by this instance of TextLineSpecBlankLines.
//
// If input parameter 'numOfBlankLines' is less than zero, this
// method will take no action and exit. If 'numOfBlankLines' is
// set to zero, no blank lines will be produced by this text line
// specification.
//
func (blkLines *TextLineSpecBlankLines) SetNumberOfBlankLines(
	numOfBlankLines int) {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	if len(blkLines.newLineChars) == 0 {
		blkLines.newLineChars = []rune{'\n'}
	}

	if numOfBlankLines < 0 {
		return
	}

	blkLines.numBlankLines = numOfBlankLines

}

// TextTypeName - returns a string specifying the type
// of Text Line specification. This method fulfills
// requirements of ITextSpecification interface.
//
func (blkLines TextLineSpecBlankLines) TextTypeName() string {

	return "TextLineSpecBlankLines"
}

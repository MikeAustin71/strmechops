package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextLineSpecBlankLines - This type is a specialized form of
// text line specification which is used to create one or more
// blank lines of text.
//
// By default, the blank line(s) are terminated with a standard
// new line character '\n'. However, users have the option to
// substitute an array of runes and apply any character or group
// of characters to terminate the line.
//
// To override the default line termination character
// (newline '\n'), see method:
//
//      TextLineSpecBlankLines.SetLineTerminationChars()
//
type TextLineSpecBlankLines struct {
	numBlankLines int
	newLineChars  []rune
	lock          *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextLineSpecBlankLines ('incomingBlkLines') to the data fields
// of the current TextLineSpecBlankLines instance ('blkLines').
//
// IMPORTANT
// All of the data fields in current TextLineSpecBlankLines
// instance ('blkLines') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingBlkLines    *TextLineSpecBlankLines
//     - A pointer to an instance of TextLineSpecBlankLines. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
//
//       All data values in this TextLineSpecBlankLines instance
//       will be copied to current TextLineSpecBlankLines
//       instance ('blkLines').
//
//       If 'incomingBlkLines' contains invalid member data
//       variables, this method will return an error.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (blkLines *TextLineSpecBlankLines) CopyIn(
	incomingBlkLines *TextLineSpecBlankLines,
	errorPrefix interface{}) error {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecBlankLines.CopyIn()",
		"")

	if err != nil {
		return err
	}

	return textLineSpecBlankLinesMolecule{}.ptr().
		copyIn(
			blkLines,
			incomingBlkLines,
			ePrefix)
}

// CopyOut - Returns a deep copy of the current
// TextLineSpecBlankLines instance.
//
// If the current TextLineSpecBlankLines instance contains invalid
// member variables, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  TextLineSpecBlankLines
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLineSpecBlankLines instance.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (blkLines *TextLineSpecBlankLines) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecBlankLines,
	error) {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecBlankLines.CopyOut()",
		"")

	if err != nil {
		return TextLineSpecBlankLines{}, err
	}

	var newBlankLinesSpec TextLineSpecBlankLines

	newBlankLinesSpec,
		err = textLineSpecBlankLinesMolecule{}.ptr().
		copyOut(
			blkLines,
			ePrefix.XCtx("blkLines"))

	return newBlankLinesSpec, err
}

// CopyOutITextLine - Returns a deep copy of the current
// TextLineSpecBlankLines instance cast as a type
// ITextLineSpecification.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  ITextLineSpecification
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLineSpecBlankLines instance cast as an
//       ITextLineSpecification object.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (blkLines *TextLineSpecBlankLines) CopyOutITextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecBlankLines.CopyOut()",
		"")

	if err != nil {
		return ITextLineSpecification(&TextLineSpecBlankLines{}),
			err
	}

	var newBlankLinesSpec TextLineSpecBlankLines

	newBlankLinesSpec,
		err = textLineSpecBlankLinesMolecule{}.ptr().
		copyOut(
			blkLines,
			ePrefix)

	return ITextLineSpecification(&newBlankLinesSpec), err
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextLineSpecBlankLines instance.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *TextLineSpecBlankLines
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a pointer to a
//       deep copy of the current TextLineSpecBlankLines instance.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (blkLines *TextLineSpecBlankLines) CopyOutPtr(
	errorPrefix interface{}) (
	*TextLineSpecBlankLines,
	error) {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecBlankLines.CopyOutPtr()",
		"")

	if err != nil {
		return &TextLineSpecBlankLines{}, err
	}

	var newBlankLinesSpec TextLineSpecBlankLines

	newBlankLinesSpec,
		err = textLineSpecBlankLinesMolecule{}.ptr().
		copyOut(
			blkLines,
			ePrefix.XCtx("blkLines"))

	return &newBlankLinesSpec, err
}

// Empty - Resets all internal member variables to their initial
// or zero states.
//
func (blkLines *TextLineSpecBlankLines) Empty() {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()
	textLineSpecBlankLinesMolecule{}.ptr().
		empty(blkLines)

	blkLines.lock.Unlock()

	blkLines.lock = nil
}

// Equal - Receives a pointer to another instance of
// TextLineSpecBlankLines and proceeds to compare the member
// variables to those of the current TextLineSpecBlankLines
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
func (blkLines *TextLineSpecBlankLines) Equal(
	incomingBlkLines *TextLineSpecBlankLines) bool {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	return textLineSpecBlankLinesMolecule{}.ptr().equal(
		blkLines,
		incomingBlkLines)
}

// EqualITextLine - Receives an object implementing the
// ITextLineSpecification interface and proceeds to compare
// the member variables to those of the current
// TextLineSpecBlankLines instance in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables from both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
// This method is required by interface ITextLineSpecification.
//
func (blkLines *TextLineSpecBlankLines) EqualITextLine(
	iTextLine ITextLineSpecification) bool {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	txtBlkLine, ok := iTextLine.(*TextLineSpecBlankLines)

	if !ok {
		return false
	}

	return textLineSpecBlankLinesMolecule{}.ptr().equal(
		blkLines,
		txtBlkLine)
}

// GetFormattedText - Returns the formatted text for output and
// printing.
//
// The value of 'blkLines.newLineChars' will be replicated
// multiple times as specified by 'blkLines.numBlankLines'.
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is identical in function to
// TextLineSpecBlankLines.String()
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (blkLines *TextLineSpecBlankLines) GetFormattedText() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecBlankLines.GetFormattedText()",
		"")

	formattedText,
		err := textLineSpecBlankLinesMolecule{}.ptr().
		getFormattedText(
			blkLines,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedText
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

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecStandardLine
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
// If the number of standard lines for the current
// TextLineSpecBlankLines is set to a value less than one, this
// method will consider the current TextLineSpecBlankLines instance
// invalid and return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix         interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If any of the internal member data variables contained in
//       the current instance of TextLineSpecBlankLines are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (blkLines *TextLineSpecBlankLines) IsValidInstanceError(
	errorPrefix interface{}) error {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecBlankLines.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textLineSpecBlankLinesAtom{}.ptr().
		testValidityOfTextLineSpecBlankLines(
			blkLines,
			ePrefix.XCtx("blkLines"))

	return err
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
//  TextLineSpecBlankLines.SetLineTerminationChars()
//
// If input parameter 'numOfBlankLines' is less than zero, this
// method will automatically reset the value to zero.
//
func (blkLines TextLineSpecBlankLines) NewPtr(
	numOfBlankLines int) *TextLineSpecBlankLines {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	if numOfBlankLines < 0 {
		numOfBlankLines = 0
	}

	newBlankLinesSpec := TextLineSpecBlankLines{}

	newBlankLinesSpec.numBlankLines = numOfBlankLines

	newBlankLinesSpec.newLineChars = []rune{'\n'}

	return &newBlankLinesSpec
}

// SetLineTerminationChars - By default, the line termination
// character for blank line produced and applied by this text line
// specification is the the new line character, '\n'. However,
// users have the option of substituting a character or series
// of characters for the the line termination sequence using this
// method.
//
// This method will receive a string as input and apply the
// characters in that string as the line termination sequence for
// this instance of TextLineSpecBlankLines.
//
// If input parameter 'lineTerminationChars' is submitted as an
// empty string, this method will take no action and exit.
//
func (blkLines *TextLineSpecBlankLines) SetLineTerminationChars(
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

// String - Returns the formatted text for output and
// printing.
//
// The value of 'blkLines.newLineChars' will be replicated
// multiple times as specified by 'blkLines.numBlankLines'.
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is identical in function to
// TextLineSpecBlankLines.GetFormattedText()
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (blkLines TextLineSpecBlankLines) String() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecBlankLines.GetFormattedText()",
		"")

	formattedText,
		err := textLineSpecBlankLinesMolecule{}.ptr().
		getFormattedText(
			&blkLines,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedText
}

// TextLineSpecName - returns Text Line Specification Name. This
// method fulfills requirements of ITextSpecification interface.
//
func (blkLines TextLineSpecBlankLines) TextLineSpecName() string {

	return "TextLineSpecBlankLines"
}

// TextTypeName - returns a string specifying the type
// of Text Line specification. This method fulfills
// requirements of ITextSpecification interface.
//
func (blkLines TextLineSpecBlankLines) TextTypeName() string {

	return "TextLineSpecBlankLines"
}

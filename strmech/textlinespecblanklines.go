package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
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
	numBlankLines  int
	newLineChars   []rune
	textLineReader *strings.Reader
	lock           *sync.Mutex
}

// CopyIn - Copies the data fields from an incoming instance of
// TextLineSpecBlankLines ('incomingBlkLines') to the data fields
// of the current TextLineSpecBlankLines instance ('blkLines').
//
// IMPORTANT
// All the data fields in current TextLineSpecBlankLines
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
// This method fulfills requirements of ITextLineSpecification
// interface.
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
// This method is similar to TextLineSpecBlankLines.String() with
// the sole difference being that this method returns an error.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
// Methods which return formatted text are listed as follows:
//  TextLineSpecBlankLines.String()
//  TextLineSpecBlankLines.TextBuilder()
//  TextLineSpecBlankLines.GetFormattedText()
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
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
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
//       4. [][2]string A two-dimensional slice of strings
//                      containing error prefix and error context
//                      information.
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
//       the 'errpref' software package,
//       "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  string
//     - The formatted text line generated by the current instance
//       of TextLineSpecBlankLines.
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
func (blkLines *TextLineSpecBlankLines) GetFormattedText(
	errorPrefix interface{}) (
	string,
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
		"TextLineSpecBlankLines.GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	return textLineSpecBlankLinesMolecule{}.ptr().
		getFormattedText(
			blkLines,
			ePrefix)
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
// This method fulfills requirements of ITextLineSpecification
// interface.
//
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

	defer blkLines.lock.Unlock()

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

// NewBlankLines - Returns a pointer to a new instance of
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
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numOfBlankLines            int
//     - The number of blank lines which will be generated by this
//       instance of TextLineSpecBlankLines
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
//     - If this method completes successfully, a valid and fully
//       populated instance of TextLineSpecBlankLines will be
//       returned.
//
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
//
func (blkLines TextLineSpecBlankLines) NewBlankLines(
	numOfBlankLines int,
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

	newBlankLinesSpec := TextLineSpecBlankLines{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecBlankLines.NewBlankLines()",
		"")

	if err != nil {
		return newBlankLinesSpec, err
	}

	if numOfBlankLines < 1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numOfBlankLines' is invalid!\n"+
			"'numOfBlankLines' is less than one (1).\n",
			ePrefix.String())

		return newBlankLinesSpec, err
	}

	if numOfBlankLines > 1000000 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numOfBlankLines' is invalid!\n"+
			"'numOfBlankLines' is greater than one-million (1,000,000).\n",
			ePrefix.String())

		return newBlankLinesSpec, err
	}

	newBlankLinesSpec.numBlankLines = numOfBlankLines

	newBlankLinesSpec.newLineChars = []rune{'\n'}

	return newBlankLinesSpec, err
}

// NewPtrBlankLines - Returns a pointer to a new instance of
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
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numOfBlankLines            int
//     - The number of blank lines which will be generated by this
//       instance of TextLineSpecBlankLines
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
//     - If this method completes successfully, a pointer to a
//       valid and fully populated instance of
//       TextLineSpecBlankLines will be returned.
//
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
//
func (blkLines TextLineSpecBlankLines) NewPtrBlankLines(
	numOfBlankLines int,
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

	newBlankLinesSpec := TextLineSpecBlankLines{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecBlankLines.NewPtrBlankLines()",
		"")

	if err != nil {
		return &newBlankLinesSpec, err
	}

	if numOfBlankLines < 1 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numOfBlankLines' is invalid!\n"+
			"'numOfBlankLines' is less than one (1).\n",
			ePrefix.String())

		return &newBlankLinesSpec, err
	}

	if numOfBlankLines > 1000000 {
		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numOfBlankLines' is invalid!\n"+
			"'numOfBlankLines' is greater than one-million (1,000,000).\n",
			ePrefix.String())

		return &newBlankLinesSpec, err
	}

	newBlankLinesSpec.numBlankLines = numOfBlankLines

	newBlankLinesSpec.newLineChars = []rune{'\n'}

	return &newBlankLinesSpec, err
}

// Read - Implements the io.Reader interface for type
// TextLineSpecBlankLines.
//
// The formatted text line string generated by the current
// instance of TextLineSpecBlankLines will be written to the
// byte buffer 'p'. If the length of 'p' is less than the length
// of the formatted text line string, multiple calls to this method
// will write the remaining unread characters to the byte buffer
// 'p'.
//
// Read() supports buffered 'read' operations.
//
// This method reads up to len(p) bytes into p. It returns the
// number of bytes read (0 <= n <= len(p)) and any error
// encountered. Even if read returns n < len(p), it may use all
// of p as scratch space during the call.
//
// If some data is available but not len(p) bytes, readBytes()
// conventionally returns what is available instead of waiting
// for more.
//
// When this method encounters an error or end-of-file condition
// after successfully reading n > 0 bytes, it returns the number
// of bytes read. It may return the (non-nil) error from the same
// call or return the error (and n == 0) from a subsequent call.
// An instance of this general case is that a Reader returning
// a non-zero number of bytes at the end of the input stream may
// return either err == EOF or err == nil. The next read operation
// should return 0, EOF.
//
// Callers should always process the n > 0 bytes returned before
// considering the error err. Doing so correctly handles I/O errors
// that happen after reading some bytes and also both of the
// allowed EOF behaviors.
//
// The last read operation performed on the formatted text string
// will always return n==0 and err==io.EOF.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  p                          []byte
//     - The byte buffer into which the formatted text line string
//       generated by the current TextLineSpecBlankLines instance
//       will be written.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  n                          int
//     - The number of bytes written to byte buffer 'p'.
//
//       Read() reads up to len(p) bytes into p. It returns
//       the number of bytes read (0 <= n <= len(p)) and any error
//       encountered. Even if Read() returns n < len(p), it may use
//       all of 'p' as scratch space during the call. If some
//       data is available but not len(p) bytes, Read()
//       conventionally returns what is available instead of
//       waiting for more.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered
//       during processing, the returned error Type will
//       encapsulate an error message.
//
//       When Read() encounters an error or end-of-file condition
//       after successfully reading n > 0 bytes, it returns the
//       number of bytes read. It may return the (non-nil) error
//       from the same call or return the error (and n == 0) from
//       a subsequent call. An instance of this general case is
//       that a Reader returning a non-zero number of bytes at the
//       end of the input stream may return either err == EOF or
//       err == nil. The next read operation should return 0, EOF.
//
//
// ------------------------------------------------------------------------
//
// Usage Examples:
//
//  Example # 1
//
//  p := make([]byte, 50)
//
//  var n, readBytesCnt int
//  sb := strings.Builder{}
//
//  for {
//
//    n,
//    err = blkLines01.Read(p)
//
//    if n == 0 {
//      break
//    }
//
//    sb.Write(p[:n])
//    readBytesCnt += n
//  }
//
//  if err != nil &&
//    err != io.EOF {
//     return fmt.Errorf(
//      "Error Returned From blkLines01.Read(p)\n"+
//      "Error = \n%v\n",
//       err.Error())
//  }
//
//  fmt.Printf("Text Line String: %s\n",
//                sb.String())
//
//  fmt.Printf("Number of bytes Read: %v\n",
//                readBytesCnt)
//
//  Example # 2
//
//  p := make([]byte, 50)
//
//  var n, readBytesCnt int
//  var actualStr string
//
//  for {
//
//    n,
//    err = blkLines01.Read(p)
//
//    if n == 0 {
//      break
//    }
//
//    actualStr += string(p[:n])
//    readBytesCnt += n
//  }
//
//  if err != nil &&
//    err != io.EOF {
//     return fmt.Errorf(
//      "Error Returned From blkLines01.Read(p)\n"+
//      "Error = \n%v\n",
//       err.Error())
//  }
//
//  fmt.Printf("Text Line String: %v\n",
//                actualStr)
//
//  fmt.Printf("Number of bytes Read: %v\n",
//                readBytesCnt)
//
func (blkLines *TextLineSpecBlankLines) Read(
	p []byte) (
	n int,
	err error) {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecBlankLines.Read()",
		"")

	if blkLines.textLineReader == nil {

		var formattedText string

		formattedText,
			err = textLineSpecBlankLinesMolecule{}.ptr().
			getFormattedText(
				blkLines,
				ePrefix.XCtx("blkLines"))

		if err != nil {
			return n, err
		}

		blkLines.textLineReader =
			strings.NewReader(formattedText)

		if blkLines.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(formattedText)\n"+
				"returned a nil pointer.\n"+
				"blkLines.textLineReader == nil\n",
				ePrefix.XCtxEmpty().String())

			return n, err
		}
	}

	n,
		err = textSpecificationAtom{}.ptr().
		readBytes(
			blkLines.textLineReader,
			p,
			ePrefix.XCtx(
				"p -> blkLines.textLineReader"))

	if err == io.EOF {

		blkLines.textLineReader = nil

	}

	return n, err
}

// ReaderInitialize - This method will reset the internal member
// variable 'TextLineSpecBlankLines.textLineReader' to its initial
// zero state of 'nil'.
//
// This method is rarely used. It provides a means of
// reinitializing the internal strings.Reader in case an error
// occurs during a read operation initiated by method
// TextLineSpecBlankLines.Read().
//
// Calling this method cleans up the residue from an aborted read
// operation and allows the calling function to start a new read
// operation.
//
// If any errors are returned by method
// TextLineSpecBlankLines.Read() which are NOT equal to io.EOF, call
// this method, TextLineSpecBlankLines.ReaderInitialize(), to reset
// the internal reader for future read operations.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (blkLines *TextLineSpecBlankLines) ReaderInitialize() {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	blkLines.textLineReader = nil

	return
}

// SetLineTerminationChars - By default, the line termination
// character for blank line produced and applied by this text line
// specification is the new line character, '\n'. However,
// users have the option of substituting a character or series
// of characters for the line termination sequence using this
// method.
//
// This method will receive a string as input and apply the
// characters in that string as the line termination sequence for
// this instance of TextLineSpecBlankLines.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  lineTerminationChars       string
//     - The line termination character or characters which will be
//       applied to every blank line produced by this instance of
//       TextLineSpecBlankLines.
//
//       If this parameter is submitted as a zero length string, an
//       error will be returned.
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
func (blkLines *TextLineSpecBlankLines) SetLineTerminationChars(
	lineTerminationChars string,
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
		"TextLineSpecBlankLines.SetLineTerminationChars()",
		"")

	if err != nil {
		return err
	}

	if len(blkLines.newLineChars) == 0 {
		blkLines.newLineChars = []rune{'\n'}
	}

	if len(lineTerminationChars) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'lineTerminationChars' is invalid!\n"+
			"'lineTerminationChars' is an empty string.",
			ePrefix.String())

		return err
	}

	sMechPreon := strMechPreon{}

	newLineTermRunes := []rune(lineTerminationChars)

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		newLineTermRunes,
		ePrefix.XCtx(
			"lineTerminationChars->newLineTermRunes"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&blkLines.newLineChars,
		&newLineTermRunes,
		true,
		ePrefix.XCtx(
			"lineTerminationChars->newLineTermRunes"+
				"->blkLines.newLineChars"))

	return err
}

// SetNumberOfBlankLines - Sets the number of blank lines produced
// by this instance of TextLineSpecBlankLines.
//
// If input parameter 'numOfBlankLines' is less than zero, this
// method will take no action and exit.
//
// If 'numOfBlankLines' is set to zero, no blank lines will be
// produced by this text line specification.
//
func (blkLines *TextLineSpecBlankLines) SetNumberOfBlankLines(
	numOfBlankLines int,
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
		"TextLineSpecBlankLines.SetNumberOfBlankLines()",
		"")

	if err != nil {
		return err
	}

	return textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			blkLines,
			numOfBlankLines,
			[]rune{'\n'},
			ePrefix.XCtx(
				"->blkLines"))
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
// This method is similar to
// TextLineSpecBlankLines.GetFormattedText() with the sole
// difference being that this method does not return an error.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
// This method also fulfills the requirements of the 'Stringer'
// interface defined in the Golang package 'fmt'. Reference:
//   https://pkg.go.dev/fmt#Stringer
//
// Methods which return formatted text are listed as follows:
//  TextLineSpecBlankLines.String()
//  TextLineSpecBlankLines.TextBuilder()
//  TextLineSpecBlankLines.GetFormattedText()
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

// TextBuilder - Configures the line of text produced by this
// instance of TextLineSpecBlankLines, and writes it to an instance
// of strings.Builder.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
// Methods which return formatted text are listed as follows:
//  TextLineSpecBlankLines.String()
//  TextLineSpecBlankLines.GetFormattedText()
//  TextLineSpecBlankLines.TextBuilder()
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  sBuilder                   *strings.Builder
//    - A pointer to an instance of strings.Builder. The line of
//      text produced by the current instance of
//      TextLineSpecBlankLines and writes that text to 'sBuilder'.
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
func (blkLines *TextLineSpecBlankLines) TextBuilder(
	sBuilder *strings.Builder,
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
		"TextLineSpecBlankLines.TextBuilder()",
		"")

	if err != nil {
		return err
	}

	if sBuilder == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sBuilder' (strings.Builder)\n"+
			"is invalid! 'sBuilder' is a 'nil' pointer.\n",
			ePrefix.String())

		return err
	}

	var formattedTxtStr string

	formattedTxtStr,
		err = textLineSpecBlankLinesMolecule{}.ptr().
		getFormattedText(
			blkLines,
			ePrefix.XCtx("blkLines"))

	if err != nil {
		return err
	}

	var err2 error

	_,
		err2 = sBuilder.WriteString(formattedTxtStr)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by sBuilder.WriteString(formattedTxtStr)\n"+
			"%v\n",
			ePrefix.XCtxEmpty().String(),
			err2.Error())
	}

	return err
}

// TextLineSpecName - returns Text Line Specification Name.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
func (blkLines TextLineSpecBlankLines) TextLineSpecName() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	return "TextLineSpecBlankLines"
}

// TextTypeName - returns a string specifying the type
// of Text Line specification.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
func (blkLines TextLineSpecBlankLines) TextTypeName() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	return "TextLineSpecBlankLines"
}

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
// Blank Lines typically consist of one or more new line
// characters ('\n') and nothing more. However, users have the
// option to provide alternate or custom line termination
// characters which will be applied instead.
//
// To override the default line termination character
// (newline '\n'), see method:
//
//	TextLineSpecBlankLines.SetNewLineChars()
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	incomingBlkLines    *TextLineSpecBlankLines
//	   - A pointer to an instance of TextLineSpecBlankLines. This
//	     method will NOT change the values of internal member
//	     variables contained in this instance.
//
//	     All data values in this TextLineSpecBlankLines instance
//	     will be copied to current TextLineSpecBlankLines
//	     instance ('blkLines').
//
//	     If 'incomingBlkLines' contains invalid member data
//	     variables, this method will return an error.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecBlankLines
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current TextLineSpecBlankLines instance.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message occurs, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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
			ePrefix.XCpy("blkLines"))

	return newBlankLinesSpec, err
}

//	CopyOutITextLine
//
//	Returns a deep copy of the current
//	TextLineSpecBlankLines instance cast as a type
//	ITextLineSpecification.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
//	ITextLineSpecification
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current
//		TextLineSpecBlankLines instance cast as an
//		ITextLineSpecification object.
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	*TextLineSpecBlankLines
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a pointer to a
//	     deep copy of the current TextLineSpecBlankLines instance.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error occurs, the text value of input parameter
//	     'errorPrefix' will be inserted or prefixed at the
//	     beginning of the error message.
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
			ePrefix.XCpy("blkLines"))

	return &newBlankLinesSpec, err
}

// Empty - Resets all internal member variables to their initial
// or zero states.
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

// EqualITextLine
//
// Receives an object implementing the
// ITextLineSpecification interface and proceeds to
// compare the member variables to those of the current
// TextLineSpecBlankLines instance in order to determine
// if they are equivalent.
//
// A boolean flag showing the result of this comparison
// is returned. If the member variables from both
// instances are equal in all respects, this flag is set
// to 'true'. Otherwise, this method returns 'false'.
//
// This method is required by interface
// ITextLineSpecification.
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

//	GetFormattedText
//
//	Returns the calculated average time for the subject
//	operation as formatted text for screen display, file
//	output or printing.
//
//	This method is similar to method:
//		TextLineSpecBlankLines.String()
//
//	The sole difference being that method and this is the
//	return type. This method returns an error.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
//	Methods which return formatted text are listed as
//	follows:
//
//		TextLineSpecBlankLines.String()
//		TextLineSpecBlankLines.TextBuilder()
//		TextLineSpecBlankLines.GetFormattedText()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
//	string
//
//		The formatted text line output generated by the
//		current instance of TextLineSpecBlankLines.
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

// GetNewLineChars - Returns a string representing the character or
// characters used as the Line termination sequence for the current
// instance of TextLineSpecBlankLines.
//
// All blank lines produced by the current instance of
// TextLineSpecBlankLines will be terminated with this character
// string.
func (blkLines *TextLineSpecBlankLines) GetNewLineChars() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	if len(blkLines.newLineChars) == 0 {
		return ""
	}

	return string(blkLines.newLineChars)
}

// GetNewLineRunes - Returns an array of runes representing the
// character or characters used as the Line termination sequence
// for the current instance of TextLineSpecBlankLines.
//
// All blank lines produced by the current instance of
// TextLineSpecBlankLines will be terminated with the character
// or characters contained in this rune array.
func (blkLines *TextLineSpecBlankLines) GetNewLineRunes() []rune {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	if len(blkLines.newLineChars) == 0 {
		return nil
	}

	return blkLines.newLineChars
}

// GetNumOfBlankLines - Returns an integer value signifying the
// number of blank lines which will be produced by this
// TextLineSpecBlankLines instance.
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

//	IsValidInstance
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current TextLineSpecBlankLines
//	instance to determine if they are valid.
//
//	If any data element evaluates as invalid, this method
//	will return a boolean value of 'false'.
//
//	If all data elements are determined to be valid, this
//	method returns a boolean value of 'true'.
//
//	This method is functionally equivalent to method:
//
//		TextLineSpecBlankLines.IsValidInstanceError()
//
//	The difference between that method and this is that
//	this method takes no input parameters and returns a
//	boolean value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	bool
//
//		If any of the internal member data variables
//		contained in the current instance of
//		TextLineSpecBlankLines are found to be invalid,
//		this method will return a boolean value of
//		'false'.
//
//		Conversely, if all internal member data variables
//		contained in the current instance of
//		TextLineSpecBlankLines are found to be valid,
//		this method returns a boolean value of 'true'.
func (blkLines *TextLineSpecBlankLines) IsValidInstance() bool {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	isValid,
		_ := textLineSpecBlankLinesAtom{}.ptr().
		testValidityOfTextLineSpecBlankLines(
			blkLines,
			nil)

	return isValid
}

//	IsValidInstanceError
//
//	Performs a diagnostic review of the data values
//	encapsulated in the current instance of
//	TextLineSpecStandardLine to determine if they are
//	valid.
//
//	If any data element evaluates as invalid, this
//	method will return an error.
//
//	If the number of standard lines for the current
//	TextLineSpecBlankLines is set to a value less than
//	one, this method will consider the current
//	TextLineSpecBlankLines instance invalid and return an
//	error.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
//	error
//
//		If any of the internal member data variables
//		contained in the current instance of
//		TextLineSpecBlankLines are found to be invalid,
//		this method will return an error containing an
//		appropriate error message.
//
//		If an error message is returned, the text value
//		of input parameter 'errorPrefix' (error prefix)
//		will be inserted or prefixed at the beginning of
//		the error message.
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
			ePrefix.XCpy("blkLines"))

	return err
}

// NewBlankLines - Returns a new, concrete instance of
// TextLineSpecBlankLines. The number of blank lines configured for
// this instance will be taken from input parameter,
// 'numOfBlankLines'.
//
// This method allows the user to specify the line termination
// character or characters configured for the returned instance
// of TextLineSpecBlankLines.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by the
//	     returned instance of TextLineSpecBlankLines.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	newLineChars               string
//	   - The line termination character or characters which will be
//	     applied to every blank line produced by the returned
//	     instance of TextLineSpecBlankLines.
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecBlankLines
//	   - If this method completes successfully, a valid and fully
//	     populated concrete instance of TextLineSpecBlankLines will
//	     be returned.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines TextLineSpecBlankLines) NewBlankLines(
	numOfBlankLines int,
	newLineChars string,
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
		"TextLineSpecBlankLines."+
			"NewBlankLines()",
		"")

	if err != nil {
		return newBlankLinesSpec, err
	}

	err = textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			&newBlankLinesSpec,
			numOfBlankLines,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"->newBlankLinesSpec"))

	return newBlankLinesSpec, err
}

// NewBlankLineRunes - Returns a new, concrete instance of
// TextLineSpecBlankLines. The number of blank lines configured for
// this instance will be taken from input parameter,
// 'numOfBlankLines'.
//
// This method allows the user to specify the line termination
// character or characters configured for the returned instance
// of TextLineSpecBlankLines.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by the
//	     returned instance of TextLineSpecBlankLines.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	newLineRunes               []rune
//	   - This rune array holds the text character or characters
//	     which will be used as the line termination sequence for
//	     all blank lines generated by the returned instance of
//	     TextLineSpecBlankLines.
//
//	     If this parameter is submitted as a 'nil' value or a zero
//	     length rune array, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecBlankLines
//	   - If this method completes successfully, a valid, fully
//	     populated, concrete instance of TextLineSpecBlankLines will
//	     be returned.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines TextLineSpecBlankLines) NewBlankLineRunes(
	numOfBlankLines int,
	newLineRunes []rune,
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
		"TextLineSpecBlankLines."+
			"NewBlankLineRunes()",
		"")

	if err != nil {
		return newBlankLinesSpec, err
	}

	err = textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			&newBlankLinesSpec,
			numOfBlankLines,
			newLineRunes,
			ePrefix.XCpy(
				"->newBlankLinesSpec"))

	return newBlankLinesSpec, err
}

// NewDefaultBlankLines - Returns a new, concrete instance of
// TextLineSpecBlankLines. The number of blank lines configured for
// this instance will be taken from input parameter,
// 'numOfBlankLines'.
//
// By default, blank line(s) are terminated with a standard new
// line character '\n'. However, users have the option to
// substitute one or more text characters for use as the line
// termination sequence by calling one of the following methods:
//
//	TextLineSpecBlankLines.NewBlankLines()
//	TextLineSpecBlankLines.NewPtrBlankLines()
//	TextLineSpecBlankLines.NewBlankLineRunes()
//	TextLineSpecBlankLines.SetSpecBlankLines()
//	TextLineSpecBlankLines.SetSpecBlankLineRunes()
//	TextLineSpecBlankLines.SetNewLineChars()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by the
//	     returned instance of TextLineSpecBlankLines.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecBlankLines
//	   - If this method completes successfully, a valid and fully
//	     populated concrete instance of TextLineSpecBlankLines will
//	     be returned.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines TextLineSpecBlankLines) NewDefaultBlankLines(
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
		"TextLineSpecBlankLines."+
			"NewDefaultBlankLines()",
		"")

	if err != nil {
		return newBlankLinesSpec, err
	}

	err = textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			&newBlankLinesSpec,
			numOfBlankLines,
			[]rune{'\n'},
			ePrefix.XCpy(
				"->newBlankLinesSpec"))

	return newBlankLinesSpec, err
}

// NewPtrBlankLines - Returns a pointer to a new instance of
// TextLineSpecBlankLines. The number of blank lines configured for
// this instance will be taken from input parameter,
// 'numOfBlankLines'.
//
// This method allows the user to specify the line termination
// character or characters configured for the returned instance
// of TextLineSpecBlankLines.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by the
//	     returned instance of TextLineSpecBlankLines.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	newLineChars               string
//	   - The line termination character or characters which will be
//	     applied to every blank line produced by the returned
//	     instance of TextLineSpecBlankLines.
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	*TextLineSpecBlankLines
//	   - If this method completes successfully, a pointer to a
//	     valid and fully populated instance of
//	     TextLineSpecBlankLines will be returned.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines TextLineSpecBlankLines) NewPtrBlankLines(
	numOfBlankLines int,
	newLineChars string,
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
		"TextLineSpecBlankLines."+
			"NewPtrBlankLines()",
		"")

	if err != nil {
		return &newBlankLinesSpec, err
	}

	err = textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			&newBlankLinesSpec,
			numOfBlankLines,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"->newBlankLinesSpec"))

	return &newBlankLinesSpec, err
}

// NewPtrBlankLineRunes - Returns a pointer to a new instance of
// TextLineSpecBlankLines. The number of blank lines configured for
// this instance will be taken from input parameter,
// 'numOfBlankLines'.
//
// This method allows the user to specify the line termination
// character or characters configured for the returned instance
// of TextLineSpecBlankLines.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by the
//	     returned instance of TextLineSpecBlankLines.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	newLineRunes               []rune
//	   - This rune array holds the text character or characters
//	     which will be used as the line termination sequence for
//	     input parameter 'txtBlankLines', an instance of
//	     TextLineSpecBlankLines.
//
//	     If this parameter is submitted as a 'nil' value or a zero
//	     length rune array, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	*TextLineSpecBlankLines
//	   - If this method completes successfully, a pointer to a
//	     valid and fully populated instance of
//	     TextLineSpecBlankLines will be returned.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines TextLineSpecBlankLines) NewPtrBlankLineRunes(
	numOfBlankLines int,
	newLineRunes []rune,
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
		"TextLineSpecBlankLines."+
			"NewPtrBlankLineRunes()",
		"")

	if err != nil {
		return &newBlankLinesSpec, err
	}

	err = textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			&newBlankLinesSpec,
			numOfBlankLines,
			newLineRunes,
			ePrefix.XCpy(
				"->newBlankLinesSpec"))

	return &newBlankLinesSpec, err
}

// NewPtrDefaultBlankLines - Returns a pointer to a new instance of
// TextLineSpecBlankLines. The number of blank lines configured for
// this instance will be taken from input parameter,
// 'numOfBlankLines'.
//
// By default, blank line(s) are terminated with a standard new
// line character '\n'. However, users have the option to
// substitute one or more text characters for use as the line
// termination sequence by calling one of the following methods:
//
//	TextLineSpecBlankLines.NewBlankLines()
//	TextLineSpecBlankLines.NewPtrBlankLines()
//	TextLineSpecBlankLines.NewBlankLineRunes()
//	TextLineSpecBlankLines.SetSpecBlankLines()
//	TextLineSpecBlankLines.SetSpecBlankLineRunes()
//	TextLineSpecBlankLines.SetNewLineChars()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by the
//	     returned instance of TextLineSpecBlankLines.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	*TextLineSpecBlankLines
//	   - If this method completes successfully, a pointer to a
//	     valid and fully populated instance of
//	     TextLineSpecBlankLines will be returned.
//
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines TextLineSpecBlankLines) NewPtrDefaultBlankLines(
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
		"TextLineSpecBlankLines."+
			"NewPtrDefaultBlankLines()",
		"")

	if err != nil {
		return &newBlankLinesSpec, err
	}

	err = textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			&newBlankLinesSpec,
			numOfBlankLines,
			[]rune{'\n'},
			ePrefix.XCpy(
				"->newBlankLinesSpec"))

	return &newBlankLinesSpec, err
}

//	Read
//
//	Implements the io.Reader interface for type
//	TextLineSpecBlankLines.
//
//	The formatted text line string generated by the
//	current instance of TextLineSpecBlankLines will be
//	written to the byte buffer 'p'. If the length of 'p'
//	is less than the length of the formatted text line
//	string, multiple calls to this method will write the
//	remaining unread characters to the byte buffer 'p'.
//
//	Read() supports buffered 'read' operations.
//
//	This method reads up to len(p) bytes into p. It
//	returns the number of bytes read (0 <= n <= len(p))
//	and any error encountered. Even if read returns
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
//	the input stream may return either 'err == EOF' or
//	'err == nil'. The next read operation should return
//	'0, EOF'.
//
//	Callers should always process the n > 0 bytes
//	returned before considering the error 'err'. Doing so
//	correctly handles I/O errors that happen after reading
//	some bytes and also both of the allowed EOF behaviors.
//
//	The last read operation performed on the formatted
//	text string will always return 'n==0' and
//	'err==io.EOF'.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	p							[]byte
//
//		The byte buffer into which the formatted text line string
//		generated by the current TextLineSpecBlankLines instance
//		will be written.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	n							int
//
//		The number of bytes written to byte buffer 'p'.
//
//		Read() reads up to len(p) bytes into p. It returns
//		the number of bytes read (0 <= n <= len(p)) and any error
//		encountered. Even if Read() returns n < len(p), it may use
//		all of 'p' as scratch space during the call. If some
//		data is available but not len(p) bytes, Read()
//		conventionally returns what is available instead of
//		waiting for more.
//
//
//	err							error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered
//		during processing, the returned error Type will
//		encapsulate an error message.
//
//		When Read() encounters an error or end-of-file condition
//		after successfully reading n > 0 bytes, it returns the
//		number of bytes read. It may return the (non-nil) error
//		from the same call or return the error (and n == 0) from
//		a subsequent call. An instance of this general case is
//		that a Reader returning a non-zero number of bytes at the
//		end of the input stream may return either err == EOF or
//		err == nil. The next read operation should return 0, EOF.
//
// ----------------------------------------------------------------
//
// # Usage
//
//	Example # 1
//
//		p := make([]byte, 50)
//
//		var n, readBytesCnt int
//		sb := strings.Builder{}
//
//		for {
//
//		  n,
//		  err = blkLines01.Read(p)
//
//		  if n == 0 {
//		    break
//		  }
//
//		  sb.Write(p[:n])
//		  readBytesCnt += n
//		}
//
//		if err != nil &&
//		  err != io.EOF {
//		   return fmt.Errorf(
//		    "Error Returned From blkLines01.Read(p)\n"+
//		    "Error = \n%v\n",
//		     err.Error())
//		}
//
//		fmt.Printf("Text Line String: %s\n",
//		              sb.String())
//
//		fmt.Printf("Number of bytes Read: %v\n",
//		              readBytesCnt)
//
//	Example # 2
//
//		p := make([]byte, 50)
//
//		var n, readBytesCnt int
//		var actualStr string
//
//		for {
//
//		  n,
//		  err = blkLines01.Read(p)
//
//		  if n == 0 {
//		    break
//		  }
//
//		  actualStr += string(p[:n])
//		  readBytesCnt += n
//		}
//
//		if err != nil &&
//		  err != io.EOF {
//		   return fmt.Errorf(
//		    "Error Returned From blkLines01.Read(p)\n"+
//		    "Error = \n%v\n",
//		     err.Error())
//		}
//
//		fmt.Printf("Text Line String: %v\n",
//		              actualStr)
//
//		fmt.Printf("Number of bytes Read: %v\n",
//		              readBytesCnt)
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
				ePrefix.XCpy("blkLines"))

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
				ePrefix.String())

			return n, err
		}
	}

	n,
		err = new(textSpecificationAtom).
		readBytes(
			blkLines.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> blkLines.textLineReader"))

	if err == io.EOF {

		blkLines.textLineReader = nil

	}

	return n, err
}

// ReaderInitialize
//
// This method will reset the internal member variable
// 'TextLineSpecBlankLines.textLineReader' to its
// initial zero state of 'nil'.
//
// This method is rarely used. It provides a means of
// reinitializing the internal strings.Reader in case an
// error occurs during a read operation initiated by
// method TextLineSpecBlankLines.Read().
//
// Calling this method cleans up the residue from an
// aborted read operation and allows the calling
// function to start a new read operation.
//
// If any errors are returned by method
// TextLineSpecBlankLines.Read() which are NOT equal to
// io.EOF, call this method,
// TextLineSpecBlankLines.ReaderInitialize(), to reset
// the internal reader for future read operations.
//
// This method fulfills requirements of interface
// ITextLineSpecification.
func (blkLines *TextLineSpecBlankLines) ReaderInitialize() {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	blkLines.textLineReader = nil

	return
}

// SetNewLineChars - By default, the "new line characters"
// used for the line termination sequence produced by the current
// instance of TextLineSpecBlankLines is the standard new line
// character, '\n'. However, this method allows users the option of
// substituting another character or series of characters for the line
// termination sequence.
//
// This method will receive a string as input and apply the
// characters in that string as the line termination sequence for
// the current instance of TextLineSpecBlankLines.
//
// Only the internal member variable controlling the line
// termination characters (TextLineSpecBlankLines.newLineChars')
// will be modified. All other internal member data variables will
// remain unchanged. Specifically, the "number of blank of lines"
// member variable ('TextLineSpecBlankLines.numBlankLines') WILL
// NOT be modified.
//
// This method is similar to method
// TextLineSpecBlankLines.SetNewLineRunes() with the sole exception
// being that this method receives a string as an input parameter
// while TextLineSpecBlankLines.SetNewLineRunes() receives a rune
// array as an input parameter.
//
// # IMPORTANT
//
// This method will overwrite and delete the existing "new line
// characters" value ('TextLineSpecBlankLines.newLineChars') for
// the current TextLineSpecBlankLines instance (blkLines).
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	newLineChars               string
//	   - The line termination character or characters which will be
//	     applied to every blank line produced by the current
//	     instance of TextLineSpecBlankLines.
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines *TextLineSpecBlankLines) SetNewLineChars(
	newLineChars string,
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
		"TextLineSpecBlankLines."+
			"SetNewLineChars()",
		"")

	if err != nil {
		return err
	}

	newLineTermRunes := []rune(newLineChars)

	err = textLineSpecBlankLinesElectron{}.ptr().
		testValidityNewLinesChars(
			[]rune(newLineChars),
			ePrefix.XCpy(
				"newLineChars"))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		newLineTermRunes,
		ePrefix.XCpy(
			"newLineChars->newLineTermRunes"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&blkLines.newLineChars,
		&newLineTermRunes,
		true,
		ePrefix.XCpy(
			"newLineChars->newLineTermRunes"+
				"->blkLines.newLineChars"))

	return err
}

// SetNewLineRunes - By default, the "new line characters"
// used for the line termination sequence produced by instances of
// TextLineSpecBlankLines is the standard new line character, '\n'.
// However, this method allows users the option of substituting
// another character or series of characters for the line
// termination sequence.
//
// This method will receive a rune array as input and apply the
// character or characters in that array as the line termination
// sequence for the current instance of TextLineSpecBlankLines.
//
// Only the internal member variable controlling the line
// termination characters (TextLineSpecBlankLines.newLineChars')
// will be modified. All other internal member data variables will
// remain unchanged. Specifically, the "number of blank of lines"
// member variable ('TextLineSpecBlankLines.numBlankLines') WILL
// NOT be modified.
//
// This method is similar to method
// TextLineSpecBlankLines.SetNewLineChars() with the sole exception
// being that this method receives a rune array input parameter
// while TextLineSpecBlankLines.SetNewLineChars() receives a string
// as an input parameter.
//
// # IMPORTANT
//
// This method will overwrite and delete the existing "new line
// characters" value ('TextLineSpecBlankLines.newLineChars') for
// the current TextLineSpecBlankLines instance (blkLines).
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	newLineRunes               []rune
//	   - This rune array holds the line termination character or
//	     characters which will be applied to every blank line
//	     produced by the current instance of TextLineSpecBlankLines.
//
//	     If this parameter is submitted as a nil value or a zero
//	     length array, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines *TextLineSpecBlankLines) SetNewLineRunes(
	newLineRunes []rune,
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
		"TextLineSpecBlankLines."+
			"SetNewLineRunes()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecBlankLinesElectron{}.ptr().
		testValidityNewLinesChars(
			newLineRunes,
			ePrefix.XCpy(
				"newLineRunes"))

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		newLineRunes,
		ePrefix.XCpy(
			"newLineRunes"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&blkLines.newLineChars,
		&newLineRunes,
		true,
		ePrefix.XCpy(
			"newLineRunes"+
				"->blkLines.newLineChars"))

	return err
}

// SetNumberOfBlankLines - Sets the number of blank lines produced
// by the current instance of TextLineSpecBlankLines.
//
// Only the internal member variable controlling the "number of
// blank lines" will be modified
// ('TextLineSpecBlankLines.numBlankLines').
//
// All other internal member data variables will remain unchanged.
// Specifically, the "new line characters" member variable
// ('TextLineSpecBlankLines.newLineChars') WILL NOT be modified.
// The "new line characters" member variable controls the line
// termination sequence for each blank line produced by the current
// instance of TextLineSpecBlankLines.
//
// # IMPORTANT
//
// This method will overwrite and delete the existing "number of
// blank lines" value for the current TextLineSpecBlankLines instance
// (blkLines).
//
// If input parameter 'numOfBlankLines' is less than one (1), it is
// invalid and an error will be returned.
//
// If input parameter 'numOfBlankLines' is greater than one-million
// (1,000,000), it is invalid and an error will be returned.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by an
//	     instance of TextLineSpecBlankLines. This value will be
//	     tested for validity.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If input parameter 'numOfBlankLines' is evaluated as a
//	     valid value, this returned error Type will be set equal to
//	     'nil'. If input parameter 'numOfBlankLines' is found to be
//	     invalid, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
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
		"TextLineSpecBlankLines."+
			"SetNumberOfBlankLines()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecBlankLinesElectron{}.ptr().
		testValidityNumOfBlankLines(
			numOfBlankLines,
			ePrefix.XCpy(
				"numOfBlankLines"))

	if err != nil {
		return err
	}

	blkLines.numBlankLines = numOfBlankLines

	return err
}

// SetSpecBlankLines - Reconfigures the current instance of
// TextLineSpecBlankLines using the input parameters,
// 'numOfBlankLines' and 'newLineChars'.
//
// This method allows the user to specify the line termination
// character or characters configured for the current instance
// of TextLineSpecBlankLines.
//
// This method differs from
// TextLineSpecBlankLines.SetSpecBlankLineRunes() in that the
// "new line characters" parameter is submitted as string.
//
// # IMPORTANT
//
// This method will overwrite and delete all the existing internal
// member data values for the current TextLineSpecBlankLines
// instance (blkLines).
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by an
//	     instance of TextLineSpecBlankLines. This value will be
//	     tested for validity.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	newLineChars               string
//	   - The line termination character or characters which will be
//	     applied to every blank line generated by current instance
//	     of TextLineSpecBlankLines.
//
//	     If this parameter is submitted as a zero length string, an
//	     error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines *TextLineSpecBlankLines) SetSpecBlankLines(
	numOfBlankLines int,
	newLineChars string,
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
		"TextLineSpecBlankLines."+
			"SetSpecBlankLines()",
		"")

	if err != nil {
		return err
	}

	return textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			blkLines,
			numOfBlankLines,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"->blkLines"))
}

// SetSpecDefaultBlankLines - Sets the number of blank lines
// produced by the current instance of TextLineSpecBlankLines as
// well as the "new-line characters" used to generate those blank
// lines.
//
// The "new-line characters" parameter is a default value which
// will be automatically configured as the standard new-line
// character '\n'. This default new line character, or line
// termination sequence, is applied to every blank line produced by
// the current instance of TextLineSpecBlankLines.
//
// Users have the option to substitute one or more text characters
// for use as the line termination sequence by calling one of the
// following methods:
//
//	TextLineSpecBlankLines.NewBlankLines()
//	TextLineSpecBlankLines.NewPtrBlankLines()
//	TextLineSpecBlankLines.NewBlankLineRunes()
//	TextLineSpecBlankLines.SetSpecBlankLines()
//	TextLineSpecBlankLines.SetSpecBlankLineRunes()
//	TextLineSpecBlankLines.SetNewLineChars()
//
// # IMPORTANT
//
// This method will overwrite and delete all the existing internal
// member data values for the current TextLineSpecBlankLines
// instance (blkLines).
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by an
//	     instance of TextLineSpecBlankLines. This value will be
//	     tested for validity.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
func (blkLines *TextLineSpecBlankLines) SetSpecDefaultBlankLines(
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
		"TextLineSpecBlankLines."+
			"SetSpecDefaultBlankLines()",
		"")

	if err != nil {
		return err
	}

	return textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			blkLines,
			numOfBlankLines,
			[]rune{'\n'},
			ePrefix.XCpy(
				"->blkLines"))
}

// SetSpecBlankLineRunes - Reconfigures the current instance of
// TextLineSpecBlankLines using the input parameters,
// 'numOfBlankLines' and 'newLineRunes'.
//
// This method allows the user to specify the line termination
// character or characters configured for the current instance
// of TextLineSpecBlankLines.
//
// This method differs from
// TextLineSpecBlankLines.SetSpecBlankLines() in that the
// "new line characters" parameter is submitted as an array of
// runes.
//
// # IMPORTANT
//
// This method will overwrite and delete all the existing internal
// member data values for the current TextLineSpecBlankLines
// instance (blkLines).
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	numOfBlankLines            int
//	   - The number of blank lines which will be generated by an
//	     instance of TextLineSpecBlankLines. This value will be
//	     tested for validity.
//
//	     If input parameter 'numOfBlankLines' is less than one (1),
//	     it is invalid and an error will be returned.
//
//	     If input parameter 'numOfBlankLines' is greater than
//	     one-million (1,000,000), it is invalid and an error will
//	     be returned.
//
//
//	newLineRunes               []rune
//	   - This rune array holds the text character or characters
//	     which will be used as the line termination sequence for
//	     all blank lines generated by the current instance of
//	     TextLineSpecBlankLines.
//
//	     If this parameter is submitted as a 'nil' value or a zero
//	     length rune array, an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	              collection of error prefix and error context
//	              information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	                 error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings containing
//	                    error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//	                       ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	                        ErrorPrefixInfo from this object will be
//	                       copied to 'errPrefDto'.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	                            a two-dimensional slice of strings
//	                            containing error prefix and error
//	                            context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
func (blkLines *TextLineSpecBlankLines) SetSpecBlankLineRunes(
	numOfBlankLines int,
	newLineRunes []rune,
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
		"TextLineSpecBlankLines."+
			"SetSpecBlankLineRunes()",
		"")

	if err != nil {
		return err
	}

	return textLineSpecBlankLinesMolecule{}.ptr().
		setTextLinesSpecBlankLines(
			blkLines,
			numOfBlankLines,
			newLineRunes,
			ePrefix.XCpy(
				"->blkLines"))
}

//	String
//
//	Returns the formatted text for output and printing.
//
//	The value of 'blkLines.newLineChars' will be
//	replicated multiple times as specified by
//	'blkLines.numBlankLines'.
//
//	If an error occurs, the error message will be
//	included in the returned string.
//
//	This method is similar to
//	TextLineSpecBlankLines.GetFormattedText() with the
//	sole difference being that this method does not
//	return an error.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
//	This method also fulfills the requirements of the
//	'Stringer' interface defined in the Golang package
//	'fmt'. Reference:
//
//		https://pkg.go.dev/fmt#Stringer
//
//	Methods which return formatted text are listed as
//	follows:
//
//		TextLineSpecBlankLines.String()
//		TextLineSpecBlankLines.TextBuilder()
//		TextLineSpecBlankLines.GetFormattedText()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	--- NONE ---
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		If this method completes successfully, this
//		string will contain the formatted text generated
//		by the current instance of
//		TextLineSpecBlankLines.
//
//		If an error condition is encountered, this string
//		will contain an appropriate error message. This
//		error message will contain the word 'Error'.
func (blkLines TextLineSpecBlankLines) String() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecBlankLines.String()",
		"")

	formattedText,
		err := new(textLineSpecBlankLinesMolecule).
		getFormattedText(
			&blkLines,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedText
}

//	TextBuilder
//
//	Configures the line of text produced by this instance
//	of TextLineSpecBlankLines, and writes it to an
//	instance of strings.Builder.
//
//	This method fulfills requirements of interface
//	ITextLineSpecification.
//
//	Methods which return formatted text are listed as
//	follows:
//
//		TextLineSpecBlankLines.String()
//		TextLineSpecBlankLines.GetFormattedText()
//		TextLineSpecBlankLines.TextBuilder()
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of *strings.Builder. The
//		formatted text characters produced by this method
//		will be written to this instance of
//		strings.Builder.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
func (blkLines *TextLineSpecBlankLines) TextBuilder(
	strBuilder *strings.Builder,
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

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a nil pointer.\n",
			ePrefix.String())

		return err
	}

	var formattedTxtStr string

	formattedTxtStr,
		err = textLineSpecBlankLinesMolecule{}.ptr().
		getFormattedText(
			blkLines,
			ePrefix.XCpy("blkLines"))

	if err != nil {
		return err
	}

	lenFormattedText := len(formattedTxtStr)

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	requiredCapacity :=
		lenFormattedText - netCapacityStrBuilder

	if requiredCapacity > 0 {

		strBuilder.Grow(requiredCapacity + 16)
	}

	var err2 error

	_,
		err2 = strBuilder.WriteString(formattedTxtStr)

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error returned by sBuilder.WriteString(formattedTxtStr)\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())
	}

	return err
}

// TextLineSpecName
//
// Returns Text Line Specification Name.
//
// This method fulfills requirements of interface
// ITextLineSpecification.
func (blkLines TextLineSpecBlankLines) TextLineSpecName() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	return "BlankLines"
}

// TextTypeName
//
// Returns a string specifying the type of Text Line
// specification.
//
// This method fulfills requirements of interface
// ITextLineSpecification.
func (blkLines TextLineSpecBlankLines) TextTypeName() string {

	if blkLines.lock == nil {
		blkLines.lock = new(sync.Mutex)
	}

	blkLines.lock.Lock()

	defer blkLines.lock.Unlock()

	return "TextLineSpecBlankLines"
}

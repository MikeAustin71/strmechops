package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"sync"
)

// TextLineSpecSolidLine - Used to create a solid line of text
// characters for screen display, file output or printing.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// By default, all solid lines are terminated by a new line
// character ('\n'). This means that each instance of a solid line
// will be formatted as a single line of text. If another new line
// character or characters is required, reference method:
//
//	TextLineSpecSolidLine.SetNewLineChars()
//
//	   Example:
//	     solidLineChars = "*"
//	     solidLineCharsRepeatCount = 5
//	     leftMargin = 3
//	     rightMargin = 3
//	     Solid line = "   *****   "
type TextLineSpecSolidLine struct {
	leftMarginChars           []rune
	rightMarginChars          []rune
	solidLineChars            []rune
	solidLineCharsRepeatCount int
	newLineChars              []rune
	turnLineTerminatorOff     bool
	textLineReader            *strings.Reader
	lock                      *sync.Mutex
}

// CopyIn - Copies all the data fields from an incoming instance of
// TextLineSpecSolidLine ('incomingTxtSolidLine') to the data fields
// of the current TextLineSpecSolidLine instance,
// ('txtSpecSolidLine').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the data fields in current TextLineSpecSolidLine instance
// ('txtSpecSolidLine') will be modified and overwritten.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 incomingTxtSolidLine       *TextLineSpecSolidLine
//	    - A pointer to an instance of TextLineSpecSolidLine. This
//	      method will NOT change the member variable data values
//	      contained in this instance.
//
//	      All data values in this TextLineSpecSolidLine instance
//	      will be copied to current TextLineSpecSolidLine
//	      instance ('txtSpecSolidLine').
//
//	      If 'incomingTxtSolidLine' contains invalid member data
//	      variables, this method will return an error.
//
//
//	 errorPrefix                interface{}
//	    - This object encapsulates error prefix text which is
//	      included in all returned error messages. Usually, it
//	      contains the name of the calling method or methods
//	      listed as a method or function chain of execution.
//
//	      If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//	      This empty interface must be convertible to one of the
//	      following types:
//
//
//	      1. nil - A nil value is valid and generates an empty
//	               collection of error prefix and error context
//	               information.
//
//	      2. string - A string containing error prefix information.
//
//	      3. []string A one-dimensional slice of strings containing
//	                  error prefix information
//
//	      4. [][2]string A two-dimensional slice of strings
//	         containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		                       from this object will be copied for use
//		                       int error and informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of
//		                        ErrPrefixDto. Information from this
//		                        object will be copied for use in error
//		                        and informational messages.
//
//	      7. IBasicErrorPrefix - An interface to a method generating
//	                             a two-dimensional slice of strings
//	                             containing error prefix and error
//	                             context information.
//
//	      If parameter 'errorPrefix' is NOT convertible to one of
//	      the valid types listed above, it will be considered
//	      invalid and trigger the return of an error.
//
//	      Types ErrPrefixDto and IBasicErrorPrefix are included in
//	      the 'errpref' software package,
//	      "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
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
func (txtSpecSolidLine *TextLineSpecSolidLine) CopyIn(
	incomingTxtSolidLine *TextLineSpecSolidLine,
	errorPrefix interface{}) error {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.CopyIn()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		copyIn(
			txtSpecSolidLine,
			incomingTxtSolidLine,
			ePrefix)

	return err
}

// CopyOut - Returns a deep copy of the current
// TextLineSpecSolidLine instance.
//
// If the current TextLineSpecSolidLine instance contains invalid
// member variable data values, this method will return an error.
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
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	    1. nil - A nil value is valid and generates an empty
//	       collection of error prefix and error context
//	       information.
//
//	    2. string - A string containing error prefix information.
//
//	    3. []string A one-dimensional slice of strings containing
//	       error prefix information
//
//	    4. [][2]string A two-dimensional slice of strings
//	       containing error prefix and error context information.
//
//	    5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	       from this object will be copied for use in error and
//	       informational messages.
//
//	    6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	       Information from this object will be copied for use in
//	       error and informational messages.
//
//	    7. IBasicErrorPrefix - An interface to a method generating
//	       a two-dimensional slice of strings containing error
//	       prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	TextLineSpecSolidLine
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current TextLineSpecSolidLine instance.
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
func (txtSpecSolidLine *TextLineSpecSolidLine) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecSolidLine,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtSolidLine := TextLineSpecSolidLine{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.CopyOut()",
		"")

	if err != nil {
		return newTxtSolidLine, err
	}

	newTxtSolidLine,
		err = textLineSpecSolidLineMolecule{}.ptr().
		copyOut(
			txtSpecSolidLine,
			ePrefix)

	return newTxtSolidLine, err
}

// CopyOutITextLine - Returns a deep copy of the current
// TextLineSpecSolidLine instance cast as a type
// ITextLineSpecification.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
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
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	ITextLineSpecification
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a deep copy of the
//	     current TextLineSpecSolidLine instance cast as an
//	     ITextLineSpecification object.
//
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered, this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error occurs, the text value of input parameter
//	     'errorPrefix' will be inserted or prefixed at the
//	     beginning of the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) CopyOutITextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.CopyOutITextLine()",
		"")

	if err != nil {
		return &TextLineSpecSolidLine{}, err
	}

	var newTxtSolidLine TextLineSpecSolidLine

	newTxtSolidLine,
		err = textLineSpecSolidLineMolecule{}.ptr().
		copyOut(
			txtSpecSolidLine,
			ePrefix)

	return ITextLineSpecification(&newTxtSolidLine), err
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextLineSpecSolidLine instance.
//
// If the current TextLineSpecSolidLine instance contains invalid
// member variable data values, this method will return an error.
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
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	*TextLineSpecSolidLine
//	   - If this method completes successfully and no errors are
//	     encountered, this parameter will return a pointer to a
//	     deep copy of the current TextLineSpecSolidLine
//	     instance.
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
func (txtSpecSolidLine *TextLineSpecSolidLine) CopyOutPtr(
	errorPrefix interface{}) (
	*TextLineSpecSolidLine,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newTxtSolidLine := TextLineSpecSolidLine{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.CopyOutPtr()",
		"")

	if err != nil {
		return &newTxtSolidLine, err
	}

	newTxtSolidLine,
		err = textLineSpecSolidLineMolecule{}.ptr().
		copyOut(
			txtSpecSolidLine,
			ePrefix)

	return &newTxtSolidLine, err
}

// Empty - Resets all internal member variables to their initial
// or zero states.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
func (txtSpecSolidLine *TextLineSpecSolidLine) Empty() {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	textLineSpecSolidLineAtom{}.ptr().
		empty(txtSpecSolidLine)

	txtSpecSolidLine.lock.Unlock()

	txtSpecSolidLine.lock = nil
}

// Equal - Receives a pointer to another instance of
// *TextLineSpecSolidLine and proceeds to compare the member
// variables to those of the current TextLineSpecSolidLine
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
func (txtSpecSolidLine *TextLineSpecSolidLine) Equal(
	incomingTxtSpecSolidLine *TextLineSpecSolidLine) bool {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return textLineSpecSolidLineAtom{}.ptr().
		equal(
			txtSpecSolidLine,
			incomingTxtSpecSolidLine)
}

// EqualITextLine - Receives an object implementing the
// ITextLineSpecification interface and proceeds to compare
// the member variables to those of the current
// TextLineSpecSolidLine instance in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables from both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
// This method is required by interface ITextLineSpecification.
func (txtSpecSolidLine *TextLineSpecSolidLine) EqualITextLine(
	iTextLine ITextLineSpecification) bool {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	incomingTxtSpecSolidLine, ok :=
		iTextLine.(*TextLineSpecSolidLine)

	if !ok {
		return false
	}

	return textLineSpecSolidLineAtom{}.ptr().
		equal(
			txtSpecSolidLine,
			incomingTxtSpecSolidLine)
}

// GetFormattedText - Returns the formatted text generated by this
// Text Line Specification for screen display output and printing.
//
// The current TextLineSpecSolidLine instance will provide the
// necessary information for generating text output which describes
// a solid line of text.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is similar to TextLineSpecSolidLine.String() with
// the sole difference being that this method returns an error.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
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
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	string
//	   - The formatted text line generated by the current instance
//	     of TextLineSpecSolidLine.
//
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
func (txtSpecSolidLine *TextLineSpecSolidLine) GetFormattedText(
	errorPrefix interface{}) (
	string,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.GetFormattedText()",
		"")

	if err != nil {
		return "", err
	}

	return textLineSpecSolidLineMolecule{}.ptr().
		getFormattedText(
			txtSpecSolidLine,
			ePrefix)
}

// GetLeftMarginLength - Returns the currently configured length
// for the solid line, left margin. This integer value represents
// the number of spaces which will be positioned on the left side
// of the solid line produced by the current instance of
// TextLineSpecSolidLine.
func (txtSpecSolidLine *TextLineSpecSolidLine) GetLeftMarginLength() int {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return len(txtSpecSolidLine.leftMarginChars)
}

// GetNewLineChars - Returns the solid line termination sequence
// as a string. New line characters ('newLineChars') are also known
// as line termination characters.
//
// By default, solid line text produced by an instance of
// TextLineSpecSolidLine will be terminated by a single new line
// character ('\n'). However, users have the option to specify
// custom line termination sequences using one or more new line
// characters as specified in one of the following methods:
//
//	TextLineSpecSolidLine.NewFullSolidLineConfig()
//	TextLineSpecSolidLine.NewFullSolidLineRunesConfig()
//	TextLineSpecSolidLine.NewPtrFullSolidLineConfig()
//	TextLineSpecSolidLine.NewPtrFullSolidLineRunesConfig()
//	TextLineSpecSolidLine.SetFullSolidLineConfig()
//	TextLineSpecSolidLine.SetFullSolidLineRunesConfig()
//	TextLineSpecSolidLine.SetNewLineChars()
//
// Example:
//
//	solidLineChars = "-*"
//	solidLineRepeatCount = 4
//	newLineChars = "\n\n"
//	generated solid line = "-*-*-*-*\n\n"
//
// This method returns the line termination character or characters
// as a string.
func (txtSpecSolidLine *TextLineSpecSolidLine) GetNewLineChars() string {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	if len(txtSpecSolidLine.newLineChars) == 0 {
		return ""
	}

	return string(txtSpecSolidLine.newLineChars)
}

// GetNewLineRunes - Returns the solid line termination sequence
// as an array of runes. New line runes or characters
// ('newLineChars') are also known as line termination characters.
//
// By default, solid line text produced by an instance of
// TextLineSpecSolidLine will be terminated by a single new line
// character ('\n'). However, users have the option to specify
// custom line termination sequences using one or more text
// characters as specified in the following methods:
//
//	TextLineSpecSolidLine.NewFullSolidLineConfig()
//	TextLineSpecSolidLine.NewFullSolidLineRunesConfig()
//	TextLineSpecSolidLine.NewPtrFullSolidLineConfig()
//	TextLineSpecSolidLine.NewPtrFullSolidLineRunesConfig()
//	TextLineSpecSolidLine.SetFullSolidLineConfig()
//	TextLineSpecSolidLine.SetFullSolidLineRunesConfig()
//	TextLineSpecSolidLine.SetNewLineChars()
//
// Example:
//
//	solidLineChars = "-*"
//	solidLineRepeatCount = 4
//	newLineChars = "\n\n"
//	generated solid line = "-*-*-*-*\n\n"
//
// This method returns the line termination character or characters
// as a string.
func (txtSpecSolidLine *TextLineSpecSolidLine) GetNewLineRunes() []rune {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	if len(txtSpecSolidLine.newLineChars) == 0 {
		return nil
	}

	return txtSpecSolidLine.newLineChars
}

// GetRightMarginLength - Returns the currently configured length
// for the solid line, left margin. This integer value represents
// the number of spaces which will be positioned on the left side
// of the solid line produced by the current instance of
// TextLineSpecSolidLine.
func (txtSpecSolidLine *TextLineSpecSolidLine) GetRightMarginLength() int {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return len(txtSpecSolidLine.rightMarginChars)
}

// GetSolidLineChars - Returns a string containing the character or
// characters which comprise the solid line produced by this
// instance of TextLineSpecSolidLine.
//
// The solid line characters will be repeated as specified by the
// solid line repeat count to produce the entire solid line text.
//
// Example:
//
//	solidLineChars = "-*"
//	solidLineRepeatCount = 4
//	generated solid line = "-*-*-*-*"
//
// This method returns the solid line character or characters as a
// string.
func (txtSpecSolidLine *TextLineSpecSolidLine) GetSolidLineChars() string {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	if len(txtSpecSolidLine.solidLineChars) == 0 {
		return ""
	}

	return string(txtSpecSolidLine.solidLineChars)
}

// GetSolidLineCharsRepeatCount - Returns an integer value which
// specifies the number of times that the solid line characters
// will be repeated in order to produce the final solid line text
// generated by this instance of TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// number of times the solid line character set will be repeated.
// This is referred to as the 'solidLineCharsRepeatCount'.
//
// Example:
//
//	solidLineChars = "-*"
//	solidLineRepeatCount = 4
//	generated solid line = "-*-*-*-*"
//
// This method returns the currently configured
// 'solidLineCharsRepeatCount'.
func (txtSpecSolidLine *TextLineSpecSolidLine) GetSolidLineCharsRepeatCount() int {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return txtSpecSolidLine.solidLineCharsRepeatCount
}

// GetSolidLineRunes - Returns an array of runes containing the
// character, or characters, which comprise the solid line produced
// by this instance of TextLineSpecSolidLine.
//
// The solid line characters will be repeated as specified by the
// solid line repeat count to produce the entire solid line text.
//
// Example:
//
//	solidLineChars = "-*"
//	solidLineRepeatCount = 4
//	generated solid line = "-*-*-*-*"
//
// This method returns the solid line character, or characters, as
// an array of runes.
//
// If the array of solid line characters has a length of zero, a
// value of 'nil' will be returned.
func (txtSpecSolidLine *TextLineSpecSolidLine) GetSolidLineRunes() []rune {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	if len(txtSpecSolidLine.solidLineChars) == 0 {
		return nil
	}

	return txtSpecSolidLine.solidLineChars
}

// GetTurnLineTerminatorOff - Returns the internal member variable
// 'turnLineTerminatorOff' as a boolean value.
//
// The 'turnLineTerminatorOff' flag controls whether a line
// termination character or characters will be automatically
// appended to each line of text produced by
// TextLineSpecSolidLine.
//
// When the boolean flag 'turnLineTerminatorOff' is set to 'false',
// line terminators as defined by member variable 'newLineChars'
// WILL BE applied as a line termination sequence for the line of
// text produced by TextLineSpecSolidLine. The default value for
// 'turnLineTerminatorOff' is 'false'.
//
// When this boolean value is set to 'true', it turns off or
// cancels the automatic generation of line terminators for each
// line of text produced by TextLineSpecSolidLine.
//
// The default line terminator is the new line character ('\n')
// which is defined by member variable 'newLineChars'. However,
// this value is subject to user control and may be overridden
// by one or more characters.
//
// Again, the default value for 'turnLineTerminatorOff' is 'false'.
func (txtSpecSolidLine *TextLineSpecSolidLine) GetTurnLineTerminatorOff() bool {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return txtSpecSolidLine.turnLineTerminatorOff
}

// IsValidInstance - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecSolidLine
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return a boolean value of 'false'.
//
// If all data elements are determined to be valid, this method
// returns a boolean value of 'true'.
//
// This method is functionally equivalent to
// TextLineSpecSolidLine.IsValidInstanceError() with the sole
// exceptions being that this method takes no input parameters and
// returns a boolean value.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	-- NONE --
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If any of the internal member data variables contained in
//	     the current instance of TextLineSpecSolidLine are found
//	     to be invalid, this method will return a boolean value of
//	     'false'.
//
//	     If all internal member data variables contained in the
//	     current instance of TextLineSpecSolidLine are found to be
//	     valid, this method returns a boolean value of 'true'.
func (txtSpecSolidLine *TextLineSpecSolidLine) IsValidInstance() bool {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	isValid,
		_ := textLineSpecSolidLineAtom{}.ptr().
		testValidityOfTextSpecSolidLine(
			txtSpecSolidLine,
			nil)

	return isValid
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecSolidLine
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If any of the internal member data variables contained in
//	     the current instance of TextLineSpecSolidLine are found
//	     to be invalid, this method will return an error.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' (error prefix) will be inserted or
//	     prefixed at the beginning of the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) IsValidInstanceError(
	errorPrefix interface{}) error {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textLineSpecSolidLineAtom{}.ptr().
		testValidityOfTextSpecSolidLine(
			txtSpecSolidLine,
			ePrefix)

	return err
}

// NewDefaultSolidLine - Creates and returns a new instance of
// TextLineSpecSolidLine using the minimal number of input
// parameters necessary to create a solid line of text. To
// accomplish this, the left margin and right margin are, by
// default, set to zero (0). The line termination character will
// be defaulted to the new-line character ('\n'). This means that
// the resulting solid line specification will have no white space
// on either side of the solid line, and that line will be
// terminated by a single new line character ('\n').
//
//	Example:
//	  solidLineChars = "*"
//	  solidLineCharsRepeatCount = 5
//	  Solid line = "*****"
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// By default, all solid lines of text are terminated with a new
// line character ('\n'). This means that each instance of a solid
// will be formatted as a single line of text. If another new line
// character or characters are required, reference method:
//
//	TextLineSpecSolidLine.SetNewLineChars()
//
// ------------------------------------------------------------------------
//
// # Default Values
//
// This method will automatically set the following default values:
//
//	leftMarginChars             - Defaults to zero (0) length rune
//	                              array.
//
//	rightMarginChars            - Defaults to zero (0) length rune
//	                              array.
//
//	line termination character  - Defaults to new line character
//	                              ('\n')
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = '*'
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 TextLineSpecSolidLine
//	    - If this method completes successfully, it will create and
//	      return a new instance of TextLineSpecSolidLine which is
//	      fully configured with all the parameters necessary to
//	      format a solid line of characters for screen display output
//	      or printing.
//
//
//	 error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewDefaultSolidLine(
	solidLineChars string,
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	TextLineSpecSolidLine,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error
	txtSolidLine := TextLineSpecSolidLine{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.NewDefaultSolidLine()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			[]rune{}, // Empty Rune Array - No Left Margin
			[]rune{}, // Empty Rune Array - No Right Margin
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			false,
			ePrefix.XCpy("txtSolidLine"))

	return txtSolidLine, err
}

// NewFullSolidLineConfig - Creates and returns a new instance of
// TextLineSpecSolidLine. This method requires more input
// parameters than other similar methods, but in return, it allows
// the user to exercise control over all facets of the
// TextLineSpecSolidLine configuration.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	 newLineChars               string
//	    - This string contains one or more characters which will
//	      be used to terminate the solid text line.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        newLineChars = "??\n\n"
//	        Solid line = "*****??\n\n"
//
//	      If this parameter is submitted as a zero length rune
//	      array, 'newLineChars' will be set to the default new
//	      line character ('\n').
//
//	      If this parameter is submitted with a string length
//	      greater than one-million (1,000,000), 'newLineChars' will
//	      be set to the default new line character ('\n').
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 txtSolidLine               TextLineSpecSolidLine
//	    - If this method completes successfully, it will create and
//	      return a new instance of TextLineSpecSolidLine which is
//	      fully configured with all the parameters necessary to
//	      format a solid line of characters for screen display output
//	      or printing.
//
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewFullSolidLineConfig(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	newLineChars string,
	errorPrefix interface{}) (
	txtSolidLine TextLineSpecSolidLine,
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewFullSolidLineConfig()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			false,
			ePrefix.XCpy("txtSolidLine"))

	return txtSolidLine, err
}

// NewFullSolidLineRunesConfig - Creates and returns a new instance of
// TextLineSpecSolidLine. This method requires more input
// parameters than other similar methods, but in return, it allows
// the user to exercise control over all facets of the
// TextLineSpecSolidLine configuration.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// This method is identical to
// TextLineSpecSolidLine.NewFullSolidLineRunesConfig() with the
// sole exception that this method accepts input parameters defined
// as rune arrays instead of strings.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginChars            []rune
//	   - An array of runes containing the text characters to be
//	     positioned on the left side of the Solid Line.
//
//	     If no left margin is required, set this parameter to an
//	     empty rune array.
//
//	     Example:
//	       solidLineChars = []rune{'*'}
//	       solidLineCharsRepeatCount = 5
//	       leftMarginChars = []rune{' ',' ', ' '} // 3-spaces
//	       rightMargin = []rune{} // Empty rune array
//	       Solid line = "   *****"
//
//
//	     If the 'leftMarginChars' rune array length is greater than
//	     one-million (1,000,000), an error will be returned.
//
//
//	rightMarginChars           []rune
//	   - An array of runes containing the text characters to be
//	     positioned on the right side of the Solid Line.
//
//	     If no right margin is required, set this parameter to an
//	     empty rune array.
//
//	     Example:
//	       solidLineChars = []rune{'*'}
//	       solidLineCharsRepeatCount = 5
//	       leftMarginChars = []rune{} // Empty rune array
//	       rightMargin = []rune{' ',' ', ' '} // 3-spaces
//	       Solid line = "*****   "
//
//	     If the 'rightMarginChars' rune array length is greater
//	     than one-million (1,000,000), an error will be returned.
//
//
//	solidLineChars             []rune
//	   - This rune array specifies the character or characters
//	     which will comprise the solid line output for screen display
//	     or printing.
//
//	     Example:
//	       solidLineChars = []rune{'*'}
//	       solidLineCharsRepeatCount = 5
//	       leftMargin = 0
//	       rightMargin = 0
//
// /        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length rune
//	      array, an error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        leftMargin = 0
//	        rightMargin = 0
//	        Solid line = "*****"
//
//
//	 newLineChars               []rune
//	    - This rune array contains one or more characters which will
//	      be used to terminate the solid text line.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        leftMargin = 0
//	        rightMargin = 0
//	        newLineChars =  []rune{'?','?','\n','\n'}
//	        Solid line = "*****??\n\n"
//
//	      If this parameter is submitted as a zero length rune
//	      array, 'newLineChars' will be set to the default new
//	      line character ('\n').
//
//	      If this parameter is submitted with an array length
//	      greater than one-million (1,000,000), 'newLineChars' will
//	      be set to the default new line character ('\n').
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 txtSolidLine               TextLineSpecSolidLine
//	    - If this method completes successfully, it will create and
//	      return a new instance of TextLineSpecSolidLine which is
//	      fully configured with all the parameters necessary to
//	      format a solid line of characters for screen display output
//	      or printing.
//
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewFullSolidLineRunesConfig(
	leftMarginChars []rune,
	rightMarginChars []rune,
	solidLineChars []rune,
	solidLineCharsRepeatCount int,
	newLineChars []rune,
	errorPrefix interface{}) (
	txtSolidLine TextLineSpecSolidLine,
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewFullSolidLineRunesConfig()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			leftMarginChars,
			rightMarginChars,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			false,
			ePrefix.XCpy("txtSolidLine"))

	return txtSolidLine, err
}

// NewSolidLine - Creates and returns a new instance of
// TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// By default, all solid lines of text are terminated with a new
// line character ('\n'). This means that each instance of a solid
// will be formatted as a single line of text. If another new line
// character or characters are required, reference method:
//
//	TextLineSpecSolidLine.SetNewLineChars()
//
// ------------------------------------------------------------------------
//
// # Default Values
//
// This method will automatically set the following default values:
//
//	line termination character  - Defaults to new line character
//	                              ('\n')
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = '*'
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 txtSolidLine               TextLineSpecSolidLine
//	    - If this method completes successfully, it will create and
//	      return a new instance of TextLineSpecSolidLine which is
//	      fully configured with all the parameters necessary to
//	      format a solid line of characters for screen display output
//	      or printing.
//
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewSolidLine(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	txtSolidLine TextLineSpecSolidLine,
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewSolidLine()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			false,
			ePrefix.XCpy("txtSolidLine"))

	return txtSolidLine, err
}

// NewSolidLineAllParms - Creates and returns a new instance of
// TextLineSpecSolidLine. This method requires more input
// parameters than other similar methods, but in return, it allows
// the user to exercise control over all facets of the
// TextLineSpecSolidLine configuration. In addition, this method
// allows the user to exercise full control over the addition of
// line termination characters at the end of the solid line of
// text produced by this instance of TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//		 leftMarginStr              string
//		    - A string containing the text characters to be positioned
//		      on the left side of the Solid Line.
//
//		      If no left margin is required, set this parameter to an
//		      empty string.
//
//		      Example:
//		        solidLineChars = "*"
//		        solidLineCharsRepeatCount = 5
//		        leftMarginStr  = "   " // 3-spaces
//		        rightMarginStr = "" // Empty string
//		        Solid line = "   *****"
//
//		      If the 'leftMarginStr' string length is greater than
//		      one-million (1,000,000), an error will be returned.
//
//
//		 rightMarginStr             string
//		    - A string containing the text characters to be positioned
//	       on the right side of the Solid Line.
//
//		      If no right margin is required, set this parameter to an
//		      empty string.
//
//		      Example:
//		        solidLineChars = "*"
//		        solidLineCharsRepeatCount = 5
//		        leftMarginStr = "" // Empty string
//		        rightMarginStr = "   " // 3-spaces
//		        Solid line = "*****   "
//
//		      If the 'rightMarginStr' string length is greater than
//		      one-million (1,000,000), an error will be returned.
//
//
//		 solidLineChars               string
//		    - This string specifies the character or characters which
//		      will comprise the solid line output for screen display or
//		      printing.
//
//		      Example:
//		        solidLineChars = "*"
//		        solidLineCharsRepeatCount = 5
//		        Solid line = "*****"
//
//		      If this parameter is submitted as a zero length string, an
//		      error will be returned.
//
//
//		 solidLineCharsRepeatCount    int
//		    - This integer value specifies the number of times that
//		      parameter 'solidLineChars' will be repeated in
//		      constructing the solid line.
//
//		      If this parameter is submitted with a value less than one
//		      (1), an error will be returned.
//
//		      Example:
//		        solidLineChars = "*"
//		        solidLineCharsRepeatCount = 5
//		        Solid line = "*****"
//
//
//		 newLineChars                 string
//		    - This string contains one or more characters which will
//		      be used to terminate the solid text line.
//
//		      Example:
//		        solidLineChars = "*"
//		        solidLineCharsRepeatCount = 5
//		        newLineChars = "??\n\n"
//		        Solid line = "*****??\n\n"
//
//		      If this parameter is submitted as a zero length string,
//		      'newLineChars' will be set to the default new line
//		      character ("\n").
//
//		      If this parameter is submitted with a string length
//		      greater than one-million (1,000,000), 'newLineChars' will
//		      be set to the default new line character ('\n').
//
//
//		 turnAutoLineTerminationOff   bool
//		    - When set to 'true', line termination characters
//		      ('newLineChars') will NOT be added to the end of the solid
//		      line text produced by this instance of
//		      TextLineSpecSolidLine.
//
//		      When set to 'false', line termination characters
//		      ('newLineChars') WILL BE added  to the end of the solid
//		      line text produced by this instance of
//		      TextLineSpecSolidLine.
//
//
//		errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//		     parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//		     1. nil - A nil value is valid and generates an empty
//		        collection of error prefix and error context
//		        information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		        error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//		        from this object will be copied for use in error and
//		        informational messages.
//
//		     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//		        Information from this object will be copied for use in
//		        error and informational messages.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		        a two-dimensional slice of strings containing error
//		        prefix and error context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 txtSolidLine               TextLineSpecSolidLine
//	    - If this method completes successfully, it will create and
//	      return a new instance of TextLineSpecSolidLine which is
//	      fully configured with all the parameters necessary to
//	      format a solid line of characters for screen display output
//	      or printing.
//
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewSolidLineAllParms(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	newLineChars string,
	turnAutoLineTerminationOff bool,
	errorPrefix interface{}) (
	txtSolidLine TextLineSpecSolidLine,
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewSolidLineAllParms()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			turnAutoLineTerminationOff,
			ePrefix.XCpy("txtSolidLine"))

	return txtSolidLine, err
}

// NewPtrDefaultSolidLine - Returns a pointer to a new instance of
// TextLineSpecSolidLine using the minimal number of input
// parameters necessary to create a solid line of text. To
// accomplish this, the left margin and right margin are, by
// default, set to zero (0). The line termination character will
// be defaulted to the new-line character ('\n'). This means that
// the resulting solid line specification will have no white space
// on either side of the solid line, and that line will be
// terminated by a single new line character ('\n').
//
//	Example:
//	  solidLineChars = "*"
//	  solidLineCharsRepeatCount = 5
//	  Solid line = "*****"
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// By default, all solid lines of text are terminated with a new
// line character ('\n'). This means that each instance of a solid
// will be formatted as a single line of text. If another new line
// character or characters are required, reference method:
//
//	TextLineSpecSolidLine.SetNewLineChars()
//
// ------------------------------------------------------------------------
//
// # Default Values
//
// This method will automatically set the following default values:
//
//	leftMarginChars             - Defaults to zero (0) length or
//	                              empty rune array. No left margin
//	                              is configured.
//
//	rightMargin                 - Defaults to zero (0) length or
//	                              empty rune array. No right margin
//	                              is configured.
//
//
//	line termination character  - Defaults to new line character
//	                              ('\n')
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = '*'
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 *TextLineSpecSolidLine
//	    - If this method completes successfully, it will return a
//	      pointer to an instance of TextLineSpecSolidLine which is
//	      fully configured with all parameters necessary to
//	      format a solid line of characters for screen display, file
//	      output, or printing.
//
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewPtrDefaultSolidLine(
	solidLineChars string,
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	*TextLineSpecSolidLine,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	txtSolidLine := TextLineSpecSolidLine{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewPtrDefaultSolidLine()",
		"")

	if err != nil {
		return &txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			[]rune{},
			[]rune{},
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			false,
			ePrefix.XCpy("txtSolidLine"))

	return &txtSolidLine, err
}

// NewPtrFullSolidLineConfig - Creates a new instance of
// TextLineSpecSolidLine and returns a pointer to that instance.
// This method requires more input parameters than other similar
// methods, but in return, it allows the user to exercise granular
// control over all facets of the TextLineSpecSolidLine
// configuration.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	 newLineChars               string
//	    - This string contains one or more characters which will
//	      be used to terminate the solid text line.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        newLineChars = "??\n\n"
//	        Solid line = "*****??\n\n"
//
//	      If this parameter is submitted as a zero length rune
//	      array, 'newLineChars' will be set to the default new
//	      line character ('\n').
//
//	      If this parameter is submitted with a string length
//	      greater than one-million (1,000,000), 'newLineChars' will
//	      be set to the default new line character ('\n').
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 *TextLineSpecSolidLine
//	    - If this method completes successfully, it will create a
//	      new instance of TextLineSpecSolidLine and return a pointer
//	      to that instance.  The new TextLineSpecSolidLine instance
//	      will be fully configured with all the parameters necessary
//	      to format a solid line of characters for screen display,
//	      file output or printing.
//
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewPtrFullSolidLineConfig(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	newLineChars string,
	errorPrefix interface{}) (
	*TextLineSpecSolidLine,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	txtSolidLine := TextLineSpecSolidLine{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewPtrFullSolidLineConfig()",
		"")

	if err != nil {
		return &txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			false,
			ePrefix.XCpy("txtSolidLine"))

	return &txtSolidLine, err
}

// NewPtrFullSolidLineRunesConfig - Creates a new instance of
// TextLineSpecSolidLine and returns a pointer to that instance.
// This method requires more input parameters than other similar
// methods, but in return, it allows the user to exercise granular
// control over all facets of the TextLineSpecSolidLine
// configuration.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// This method is identical to
// TextLineSpecSolidLine.NewFullSolidLineRunesConfig() with the
// sole exception that this method accepts input parameters defined
// as rune arrays instead of strings.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	leftMarginChars            []rune
//	   - An array of runes containing the text characters to be
//	     positioned on the left side of the Solid Line.
//
//	     If no left margin is required, set this parameter to an
//	     empty rune array.
//
//	     Example:
//	       solidLineChars = []rune{'*'}
//	       solidLineCharsRepeatCount = 5
//	       leftMarginChars = []rune{' ',' ', ' '} // 3-spaces
//	       rightMargin = []rune{} // Empty rune array
//	       Solid line = "   *****"
//
//
//	     If the 'leftMarginChars' rune array length is greater than
//	     one-million (1,000,000), an error will be returned.
//
//
//	rightMarginChars           []rune
//	   - An array of runes containing the text characters to be
//	     positioned on the right side of the Solid Line.
//
//	     If no right margin is required, set this parameter to an
//	     empty rune array.
//
//	     Example:
//	       solidLineChars = []rune{'*'}
//	       solidLineCharsRepeatCount = 5
//	       leftMarginChars = []rune{} // Empty rune array
//	       rightMargin = []rune{' ',' ', ' '} // 3-spaces
//	       Solid line = "*****   "
//
//	     If the 'rightMarginChars' rune array length is greater
//	     than one-million (1,000,000), an error will be returned.
//
//
//	solidLineChars             []rune
//	   - This rune array specifies the character or characters
//	     which will comprise the solid line output for screen display
//	     or printing.
//
//	     Example:
//	       solidLineChars = []rune{'*'}
//	       solidLineCharsRepeatCount = 5
//	       leftMargin = 0
//	       rightMargin = 0
//
// /        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length rune
//	      array, an error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        leftMargin = 0
//	        rightMargin = 0
//	        Solid line = "*****"
//
//
//	 newLineChars               []rune
//	    - This rune array contains one or more characters which will
//	      be used to terminate the solid text line.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        leftMargin = 0
//	        rightMargin = 0
//	        newLineChars =  []rune{'?','?','\n','\n'}
//	        Solid line = "*****??\n\n"
//
//	      If this parameter is submitted as a zero length rune
//	      array, 'newLineChars' will be set to the default new
//	      line character ('\n').
//
//	      If this parameter is submitted with an array length
//	      greater than one-million (1,000,000), 'newLineChars' will
//	      be set to the default new line character ('\n').
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 *TextLineSpecSolidLine
//	    - If this method completes successfully, it will create a
//	      new instance of TextLineSpecSolidLine and return a pointer
//	      to that instance. The new TextLineSpecSolidLine instance
//	      will be fully configured with all the parameters necessary
//	      to format a solid line of characters for screen display,
//	      file output or printing.
//
//
//	 error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewPtrFullSolidLineRunesConfig(
	leftMarginChars []rune,
	rightMarginChars []rune,
	solidLineChars []rune,
	solidLineCharsRepeatCount int,
	newLineChars []rune,
	errorPrefix interface{}) (
	*TextLineSpecSolidLine,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	txtSolidLine := TextLineSpecSolidLine{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewPtrFullSolidLineRunesConfig()",
		"")

	if err != nil {
		return &txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			leftMarginChars,
			rightMarginChars,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			false,
			ePrefix.XCpy("txtSolidLine"))

	return &txtSolidLine, err
}

// NewPtrSolidLine - Creates a new instance of
// TextLineSpecSolidLine and returns a pointer to that instance.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// By default, all solid lines of text are terminated with a new
// line character ('\n'). This means that each instance of a solid
// will be formatted as a single line of text. If another new line
// character or characters are required, reference method:
//
//	TextLineSpecSolidLine.SetNewLineChars()
//
// ------------------------------------------------------------------------
//
// # Default Values
//
// This method will automatically set the following default values:
//
//	line termination character  - Defaults to new line character
//	                              ('\n')
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = '*'
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 *TextLineSpecSolidLine
//	    - If this method completes successfully, it will create a
//	      new instance of TextLineSpecSolidLine and return a
//	      pointer to that instance. The new TextLineSpecSolidLine
//	      instance will be fully configured with all the parameters
//	      necessary to format a solid line of characters for text
//	      display output or printing.
//
//
//	 error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewPtrSolidLine(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	*TextLineSpecSolidLine,
	error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	txtSolidLine := TextLineSpecSolidLine{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewPtrSolidLine()",
		"")

	if err != nil {
		return &txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			false,
			ePrefix.XCpy("txtSolidLine"))

	return &txtSolidLine, err
}

// NewPtrSolidLineAllParms - Creates a new instance of
// TextLineSpecSolidLine and returns a pointer to that instance.
// This method requires more input parameters than other similar
// methods, but in return, it allows/ the user to exercise control
// over all facets of the TextLineSpecSolidLine configuration.
// In addition, this method allows the user to exercise full
// control over the addition of line termination characters at the
// end of the solid line of text produced by this instance of
// TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars               string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount    int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	 newLineChars                 string
//	    - This string contains one or more characters which will
//	      be used to terminate the solid text line.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        newLineChars = "??\n\n"
//	        Solid line = "*****??\n\n"
//
//	      If this parameter is submitted as a zero length rune
//	      array, 'newLineChars' will be set to the default new
//	      line character ('\n').
//
//	      If this parameter is submitted with a string length
//	      greater than one-million (1,000,000), 'newLineChars' will
//	      be set to the default new line character ('\n').
//
//
//	 turnAutoLineTerminationOff   bool
//	    - When set to 'true', line termination characters
//	      ('newLineChars') will NOT be added to the end of the solid
//	      line text produced by this instance of
//	      TextLineSpecSolidLine.
//
//	      When set to 'false', line termination characters
//	      ('newLineChars') WILL BE added  to the end of the solid
//	      line text produced by this instance of
//	      TextLineSpecSolidLine.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 txtSolidLine               *TextLineSpecSolidLine
//	    - If this method completes successfully, it will create a
//	      new instance of TextLineSpecSolidLine and return a pointer
//	      to that instance. The new TextLineSpecSolidLine instance
//	      will be fully configured with all the parameters necessary
//	      to format a solid line of characters for screen display,
//	      file output or printing.
//
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine TextLineSpecSolidLine) NewPtrSolidLineAllParms(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	newLineChars string,
	turnAutoLineTerminationOff bool,
	errorPrefix interface{}) (
	txtSolidLine *TextLineSpecSolidLine,
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	txtSolidLine = &TextLineSpecSolidLine{}

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"NewPtrSolidLineAllParms()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			txtSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			turnAutoLineTerminationOff,
			ePrefix.XCpy("txtSolidLine"))

	return txtSolidLine, err
}

// Read - Implements the io.Reader interface for type
// TextLineSpecSolidLine.
//
// The formatted text line string generated by the current
// instance of TextLineSpecSolidLine will be written to the
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	p                          []byte
//	   - The byte buffer into which the formatted text line string
//	     generated by the current TextLineSpecSolidLine instance
//	     will be written.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	n                          int
//	   - The number of bytes written to byte buffer 'p'.
//
//	     Read() reads up to len(p) bytes into p. It returns
//	     the number of bytes read (0 <= n <= len(p)) and any error
//	     encountered. Even if Read() returns n < len(p), it may use
//	     all of 'p' as scratch space during the call. If some
//	     data is available but not len(p) bytes, Read()
//	     conventionally returns what is available instead of
//	     waiting for more.
//
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered
//	     during processing, the returned error Type will
//	     encapsulate an error message.
//
//	     When Read() encounters an error or end-of-file condition
//	     after successfully reading n > 0 bytes, it returns the
//	     number of bytes read. It may return the (non-nil) error
//	     from the same call or return the error (and n == 0) from
//	     a subsequent call. An instance of this general case is
//	     that a Reader returning a non-zero number of bytes at the
//	     end of the input stream may return either err == EOF or
//	     err == nil. The next read operation should return 0, EOF.
//
// ------------------------------------------------------------------------
//
// Usage Examples:
//
//	Example # 1
//
//	p := make([]byte, 50)
//
//	var n, readBytesCnt int
//	sb := strings.Builder{}
//
//	for {
//
//	  n,
//	  err = txtSpecSolidLine01.Read(p)
//
//	  if n == 0 {
//	    break
//	  }
//
//	  sb.Write(p[:n])
//	  readBytesCnt += n
//	}
//
//	if err != nil &&
//	  err != io.EOF {
//	   return fmt.Error(
//	    "Error Returned From txtSpecSolidLine01.Read(p)\n"+
//	    "Error = \n%v\n",
//	     err.Error())
//	}
//
//	fmt.Printf("Text Line String: %s\n",
//	              sb.String())
//
//	fmt.Printf("Number of bytes Read: %v\n",
//	              readBytesCnt)
//
//	Example # 2
//
//	p := make([]byte, 50)
//
//	var n, readBytesCnt int
//	var actualStr string
//
//	for {
//
//	  n,
//	  err = txtSpecSolidLine01.Read(p)
//
//	  if n == 0 {
//	    break
//	  }
//
//	  actualStr += string(p[:n])
//	  readBytesCnt += n
//	}
//
//	if err != nil &&
//	  err != io.EOF {
//	   return fmt.Error(
//	    "Error Returned From txtSpecSolidLine01.Read(p)\n"+
//	    "Error = \n%v\n",
//	     err.Error())
//	}
//
//	fmt.Printf("Text Line String: %v\n",
//	              actualStr)
//
//	fmt.Printf("Number of bytes Read: %v\n",
//	              readBytesCnt)
func (txtSpecSolidLine *TextLineSpecSolidLine) Read(
	p []byte) (
	n int,
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecSolidLine.Read()",
		"")

	if txtSpecSolidLine.textLineReader == nil {

		var formattedText string

		formattedText,
			err = textLineSpecSolidLineMolecule{}.ptr().
			getFormattedText(
				txtSpecSolidLine,
				ePrefix.XCpy("txtSpecSolidLine->formattedText"))

		if err != nil {
			return n, err
		}

		txtSpecSolidLine.textLineReader =
			strings.NewReader(formattedText)

		if txtSpecSolidLine.textLineReader == nil {
			err = fmt.Errorf("%v\n"+
				"Error: strings.NewReader(formattedText)\n"+
				"returned a nil pointer.\n"+
				"txtSpecSolidLine.textLineReader == nil\n",
				ePrefix.String())

			return n, err
		}
	}

	n,
		err = new(textSpecificationAtom).
		readBytes(
			txtSpecSolidLine.textLineReader,
			p,
			ePrefix.XCpy(
				"p -> txtSpecSolidLine.textLineReader"))

	if err == io.EOF {

		txtSpecSolidLine.textLineReader = nil

	}

	return n, err
}

// ReaderInitialize - This method will reset the internal member
// variable 'TextLineSpecSolidLine.textLineReader' to its initial
// zero state of 'nil'. Effectively, this resets the internal
// strings.Reader object for use in future read operations.
//
// This method is rarely used or needed. It provides a means of
// reinitializing the internal strings.Reader object in case an
// error occurs during a read operation initiated by method
// TextLineSpecSolidLine.Read().
//
// Calling this method cleans up the residue from an aborted read
// operation and prepares the strings.Reader object for future read
// operations.
//
// If any errors are returned by method
// TextLineSpecSolidLine.Read() which are NOT equal to io.EOF, call
// this method, TextLineSpecSolidLine.ReaderInitialize(), to reset
// and prepare the internal reader for future read operations.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
func (txtSpecSolidLine *TextLineSpecSolidLine) ReaderInitialize() {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	txtSpecSolidLine.textLineReader = nil

	return
}

// SetFullSolidLineConfig - This method configures the current
// TextLineSpecSolidLine instance with new solid line parameters.
// All internal member variable data values for the current
// TextLineSpecSolidLine instance will be reset and overwritten.
// This method requires more input parameters than other similar
// methods, but in return, it allows the user to exercise control
// over all facets of the TextLineSpecSolidLine configuration.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// ------------------------------------------------------------------------
//
// # IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecSolidLine instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	 newLineChars               string
//	    - This string contains one or more characters which will be
//	      used to terminate the solid text line.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        newLineChars = "??\n\n"
//	        Solid line = "*****??\n\n"
//
//	      If this parameter is submitted as a zero length string
//	      it will be set to the default new line character ('\n').
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetFullSolidLineConfig(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	newLineChars string,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"SetFullSolidLineConfig()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			txtSpecSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			false,
			ePrefix.XCpy("txtSpecSolidLine"))

	return err
}

// SetFullSolidLineRunesConfig - This method configures the current
// TextLineSpecSolidLine instance with new solid line parameters.
// All internal member variable data values for the current
// TextLineSpecSolidLine instance will be reset and overwritten.
// This method requires more input parameters than other similar
// methods, but in return, it allows the user to exercise control
// over all facets of the TextLineSpecSolidLine configuration.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// ------------------------------------------------------------------------
//
// # IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecSolidLine instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginChars            []rune
//	    - An array of runes containing the text characters to be
//	      positioned on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty rune array.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        leftMarginChars = []rune{' ',' ', ' '} // 3-spaces
//	        rightMargin = []rune{} // Empty rune array
//	        Solid line = "   *****"
//
//
//	      If the 'leftMarginChars' rune array length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginChars           []rune
//	    - An array of runes containing the text characters to be
//	      positioned on the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty rune array.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        leftMarginChars = []rune{} // Empty rune array
//	        rightMargin = []rune{' ',' ', ' '} // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginChars' rune array length is greater
//	      than one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars             []rune
//	    - This rune array specifies the character or characters
//	      which will comprise the solid line output for screen display
//	      or printing.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length rune
//	      array, an error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	 newLineChars               []rune
//	    - This rune array contains one or more characters which will
//	      be used to terminate the solid text line.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        newLineChars =  []rune{'?','?','\n',\n'}
//	        Solid line = "*****??\n\n"
//
//	      If this parameter is submitted as a zero length rune
//	      array, 'newLineChars' will be set to the default new
//	      line character ('\n').
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetFullSolidLineRunesConfig(
	leftMarginChars []rune,
	rightMarginChars []rune,
	solidLineChars []rune,
	solidLineCharsRepeatCount int,
	newLineChars []rune,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"SetFullSolidLineRunesConfig()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			txtSpecSolidLine,
			leftMarginChars,
			rightMarginChars,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			false,
			ePrefix.XCpy("txtSpecSolidLine"))

	return err
}

// SetLeftMargin - Sets the left margin for the current instance
// of TextLineSpecSolidLine.
//
// The left margin defines the text characters which will be
// positioned on the left side of the solid line.
//
//	Example:
//	  solidLineChars = "*"
//	  solidLineCharsRepeatCount = 5
//	  leftMarginStr  = "   " // 3-spaces
//	  rightMarginStr = "" // Empty string
//	  Solid line = "   *****"
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetLeftMargin(
	leftMarginStr string,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.SetLeftMargin()",
		"")

	if err != nil {
		return err
	}

	if len(leftMarginStr) > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leftMarginStr' is invalid!\n"+
			"The string length of 'leftMarginStr' is greater than 1,000,000.\n"+
			"leftMarginStr length ='%v'\n",
			ePrefix.String(),
			len(leftMarginStr))

		return err
	}

	leftMarginChars := []rune(leftMarginStr)

	sMechPreon := strMechPreon{}

	var err2 error
	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		leftMarginChars,
		ePrefix.XCpy(
			"Error: Input parameter 'leftMarginStr'"+
				" is invalid!"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leftMarginStr' is invalid!\n"+
			"When converted to runes, the following error was returned:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	txtSpecSolidLine.leftMarginChars = nil

	err = sMechPreon.copyRuneArrays(
		&txtSpecSolidLine.leftMarginChars,
		&leftMarginChars,
		true,
		ePrefix.XCpy(
			"Input Parameter leftMarginStr->"+
				"txtSpecSolidLine.leftMarginChars"))

	return err
}

// SetNewLineChars - Sets the new line character or
// characters which will be used to terminate the solid line
// generated by this instance of TextLineSpecSolidLine.
//
// By default, a line termination sequence (usually a new line
// character '\n') is automatically appended to the end of each
// solid line of text produced by an instance of
// TextLineSpecSolidLine.
//
// The new line character or characters may be customized by
// calling this method.
//
// If 'newLineChars' is submitted with a string length greater
// than one-million (1,000,000), an error will be returned.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 newLineChars               string
//	    - This string contains one or more characters which will
//	      be used to terminate the solid text line.
//
//	      This example creates a custom line termination sequence
//	      which includes two new line characters:
//	        solidLineChars = '*'
//	        solidLineCharsRepeatCount = 5
//	        newLineChars = "??\n\n"
//	        Solid line = "*****??\n"
//	                     "\n"
//
//	      If this parameter is submitted as a zero length string,
//	      'newLineChars' will be set to the default new
//	      line character ('\n').
//
//	      If this parameter is submitted as a string length greater
//	      than one-million (1,000,000), an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetNewLineChars(
	newLineChars string,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.SetNewLineChars()",
		"")

	if err != nil {
		return err
	}

	lenNewLineChars := len(newLineChars)

	if lenNewLineChars > 1000000 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Parameter 'newLineChars' is invalid!\n"+
			"The length of 'newLineChars' is greater than one-million (1,000,000).\n"+
			"txtSolidLine.newLineChars length ='%v'\n",
			ePrefix.String(),
			lenNewLineChars)

		return err
	}

	if lenNewLineChars == 0 {
		newLineChars = "\n"
	}

	sMechPreon := strMechPreon{}

	newLineRunes := []rune(newLineChars)

	err = sMechPreon.copyRuneArrays(
		&txtSpecSolidLine.newLineChars,
		&newLineRunes,
		true,
		ePrefix.XCpy("newLineRunes->txtSpecSolidLine.newLineChars"))

	return err
}

// SetNewLineRunes - Sets the new line character, or characters,
// which will be used to terminate the solid line generated by this
// instance of TextLineSpecSolidLine.
//
// By default, all solid lines of text are terminated by a new line
// character ('\n). This means that each instance of a solid will
// be formatted as a single line of text.
//
// This line termination sequence may be customized by providing an
// array of runes which will be substituted for the default line
// termination sequence.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	newLineRunes               []rune
//	   - This rune array contains one or more characters which will
//	     be used to terminate the solid text line.
//
//	     This example creates a custom line termination sequence
//	     which includes two new line characters:
//	       solidLineChars = '*'
//	       solidLineCharsRepeatCount = 5
//	       newLineChars = "??\n\n"
//	       Solid line = "*****??\n"
//	                    "\n"
//
//	     If this parameter is submitted as a zero length rune
//	     array, 'newLineChars' will be set to the default new
//	     line character ('\n').
//
//	     If this parameter is submitted with an array length
//	     greater than one-million (1,000,000), an error will be
//	     returned.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	--- NONE ---
func (txtSpecSolidLine *TextLineSpecSolidLine) SetNewLineRunes(
	newLineRunes []rune,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"SetNewLineRunes()",
		"")

	if err != nil {
		return err
	}

	lenNewLineRunes := len(newLineRunes)

	if lenNewLineRunes < 1 {

		txtSpecSolidLine.newLineChars = []rune{'\n'}

		return err
	}

	if lenNewLineRunes > 1000000 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Parameter 'newLineRunes' is invalid!\n"+
			"The length of 'newLineRunes' is greater than one-million (1,000,000).\n"+
			"newLineRunes length ='%v'\n",
			ePrefix.String(),
			lenNewLineRunes)

		return err
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		newLineRunes,
		ePrefix.XCpy(
			"newLineRunes invalid!"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&txtSpecSolidLine.newLineChars,
		&newLineRunes,
		true,
		nil)

	return err
}

// SetSolidLineAllParms - This method configures the current
// TextLineSpecSolidLine instance with new solid line parameters.
// All internal member variable data values for the current
// TextLineSpecSolidLine instance will be reset and overwritten.
// This method requires more input parameters than other similar
// methods, but in return, it allows the user to exercise control
// over all facets of the TextLineSpecSolidLine configuration. In
// addition, this method allows the user to exercise full control
// over the addition of line termination characters at the end of
// the solid line of text produced by this instance of
// TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// ------------------------------------------------------------------------
//
// # IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecSolidLine instance.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	 newLineChars               string
//	    - This string contains one or more characters which will be
//	      used to terminate the solid text line.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        newLineChars = "??\n\n"
//	        Solid line = "*****??\n\n"
//
//	      If this parameter is submitted as a zero length string
//	      it will be set to the default new line character ('\n').
//
//
//	 turnAutoLineTerminationOff   bool
//	    - When set to 'true', line termination characters
//	      ('newLineChars') will NOT be added to the end of the solid
//	      line text produced by this instance of
//	      TextLineSpecSolidLine.
//
//	      When set to 'false', line termination characters
//	      ('newLineChars') WILL BE added  to the end of the solid
//	      line text produced by this instance of
//	      TextLineSpecSolidLine.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetSolidLineAllParms(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	newLineChars string,
	turnAutoLineTerminationOff bool,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"SetSolidLineAllParms()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			txtSpecSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			turnAutoLineTerminationOff,
			ePrefix.XCpy("txtSpecSolidLine"))

	return err
}

// SetSolidLineCharsRepeatCount - Sets the number of times that
// solid line characters will be repeated when generating the final
// solid line text for this instance of TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// number of times the solid line character set will be repeated.
// This is referred to as the 'solidLineCharsRepeatCount'.
//
// This method takes two input parameters which configure the solid
// character(s) and the number of times that solid character(s) is
// repeated to construct the entire solid line of text.
//
//	Example:
//	  solidLineChars = "*"
//	  solidLineCharsRepeatCount = 5
//	  Solid line = "*****"
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1) or greater than one-million (1,000,000), an error will
//	      be returned.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetSolidLineCharsRepeatCount(
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine."+
			"SetSolidLineCharsRepeatCount()",
		"")

	if err != nil {
		return err
	}

	if solidLineCharsRepeatCount < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"'solidLineCharsRepeatCount' has a value less than one (1).\n"+
			"solidLineCharsRepeatCount = '%v'\n",
			ePrefix.String(),
			solidLineCharsRepeatCount)

		return err
	}

	if solidLineCharsRepeatCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"'solidLineCharsRepeatCount' has a value greater than one-million (1,000,000).\n"+
			"solidLineCharsRepeatCount = '%v'\n",
			ePrefix.String(),
			solidLineCharsRepeatCount)

		return err
	}

	txtSpecSolidLine.solidLineCharsRepeatCount =
		solidLineCharsRepeatCount

	return
}

// SetRightMargin - Sets the right margin for the current instance
// of TextLineSpecSolidLine.
//
// The right margin defines the text characters which will be
// appended to the end, or right side, of the solid line.
//
//	Example:
//	  solidLineChars = "*"
//	  solidLineCharsRepeatCount = 5
//	  leftMarginStr = "" // Empty string
//	  rightMarginStr = "   " // 3-spaces
//	  Solid line = "*****   "
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetRightMargin(
	rightMarginStr string,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.SetRightMargin()",
		"")

	if err != nil {
		return err
	}

	if len(rightMarginStr) > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'rightMarginStr' is invalid!\n"+
			"The string length of 'rightMarginStr' is greater than 1,000,000.\n"+
			"rightMarginStr length ='%v'\n",
			ePrefix.String(),
			len(rightMarginStr))

		return err
	}

	rightMarginChars := []rune(rightMarginStr)

	sMechPreon := strMechPreon{}

	var err2 error
	_,
		err2 = sMechPreon.testValidityOfRuneCharArray(
		rightMarginChars,
		ePrefix.XCpy(
			"Error: Input parameter 'rightMarginStr'"+
				" is invalid!"))

	if err2 != nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'rightMarginStr' is invalid!\n"+
			"When converted to runes, the following error was returned:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	txtSpecSolidLine.rightMarginChars = nil

	err = sMechPreon.copyRuneArrays(
		&txtSpecSolidLine.rightMarginChars,
		&rightMarginChars,
		true,
		ePrefix.XCpy(
			"Input Parameter rightMarginStr->"+
				"txtSpecSolidLine.rightMarginChars"))

	return err
}

// SetSolidLine - This method configures the current
// TextLineSpecSolidLine instance with new solid line parameters.
// All internal member variable data values for the current
// TextLineSpecSolidLine instance will be reset and overwritten.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// By default, all solid lines of text are terminated with a new
// line character ('\n'). This means that each instance of a solid
// will be formatted as a single line of text. If another new line
// character or characters are required, reference method:
//
//	TextLineSpecSolidLine.SetNewLineChars()
//
// ------------------------------------------------------------------------
//
// # IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecSolidLine instance.
//
// ------------------------------------------------------------------------
//
// # Default Values
//
// This method will automatically set the following default values:
//
//	line termination character  - Defaults to new line character
//	                              ('\n')
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 leftMarginStr              string
//	    - A string containing the text characters to be positioned
//	      on the left side of the Solid Line.
//
//	      If no left margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr  = "   " // 3-spaces
//	        rightMarginStr = "" // Empty string
//	        Solid line = "   *****"
//
//	      If the 'leftMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 rightMarginStr             string
//	    - A string containing the text characters to positioned on
//	      the right side of the Solid Line.
//
//	      If no right margin is required, set this parameter to an
//	      empty string.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        leftMarginStr = "" // Empty string
//	        rightMarginStr = "   " // 3-spaces
//	        Solid line = "*****   "
//
//	      If the 'rightMarginStr' string length is greater than
//	      one-million (1,000,000), an error will be returned.
//
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If this parameter is submitted as a zero length string, an
//	      error will be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1), an error will be returned.
//
//	      Example:
//	        solidLineChars = '*'
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetSolidLine(
	leftMarginStr string,
	rightMarginStr string,
	solidLineChars string,
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.SetSolidLine()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			txtSpecSolidLine,
			[]rune(leftMarginStr),
			[]rune(rightMarginStr),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			false,
			ePrefix.XCpy("txtSpecSolidLine"))

	return err
}

// SetSolidLineStrChars - Sets the solid line character or
// characters for the current instance of TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// number of times the solid line character set will be repeated.
// This is referred to as the 'solidLineCharsRepeatCount'.
//
// This method takes two input parameters which configure the solid
// character(s) and the number of times that solid character(s) is
// repeated to construct the entire solid line of text.
//
// This method is identical to
// TextLineSpecSolidLine.SetSolidLineRuneChars with the sole
// exception that this method receives solid line characters as a
// string instead of a rune array.
//
// Be advised that this method does alter the current values for
// left margin and right margin.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	 solidLineChars             string
//	    - This string specifies the character or characters which
//	      will comprise the solid line output for screen display or
//	      printing.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//	      If the string length of 'solidLineChars' is less than one
//	      (1) or greater than one-million (1,000,000), an error will
//	      be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1) or greater than one-million (1,000,000), an error will
//	      be returned.
//
//	      Example:
//	        solidLineChars = "*"
//	        solidLineCharsRepeatCount = 5
//	        Solid line = "*****"
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetSolidLineStrChars(
	solidLineChars string,
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.SetSolidLineStrChars()",
		"")

	if err != nil {
		return err
	}

	lenSolidLineChars := len(solidLineChars)

	if lenSolidLineChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineChars' is invalid!\n"+
			"'solidLineChars' is a zero length string.\n",
			ePrefix.String())

		return err
	}

	if lenSolidLineChars > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineChars' is invalid!\n"+
			"'solidLineChars' has a string length greater than "+
			"one-million (1,000,000).\n",
			ePrefix.String())

		return err
	}

	if solidLineCharsRepeatCount == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"'solidLineCharsRepeatCount' has a value less of zero (0).\n",
			ePrefix.String())

		return err
	}

	if solidLineCharsRepeatCount < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"'solidLineCharsRepeatCount' has a value less than zero (0).\n"+
			"solidLineCharsRepeatCount = '%v'\n",
			ePrefix.String(),
			solidLineCharsRepeatCount)

		return err
	}

	if solidLineCharsRepeatCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"'solidLineCharsRepeatCount' has a value greater than one-million (1,000,000).\n"+
			"solidLineCharsRepeatCount = '%v'\n",
			ePrefix.String(),
			solidLineCharsRepeatCount)

		return err
	}

	solidLineRunes := []rune(solidLineChars)

	err = strMechPreon{}.ptr().copyRuneArrays(
		&txtSpecSolidLine.solidLineChars,
		&solidLineRunes,
		true,
		ePrefix.XCpy("solidLineChars->"+
			"txtSpecSolidLine.solidLineChars"))

	if err != nil {
		return err
	}

	txtSpecSolidLine.solidLineCharsRepeatCount =
		solidLineCharsRepeatCount

	return err
}

// SetSolidLineRuneChars - Sets the solid line character or
// characters for the current instance of TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// number of times the solid line character set will be repeated.
// This is referred to as the 'solidLineCharsRepeatCount'.
//
// This method takes two input parameters which configure the solid
// character(s) and the number of times that solid character(s) is
// repeated to construct the entire solid line of text.
//
// This method is identical to
// TextLineSpecSolidLine.SetSolidLineStrChars with the sole
// exception that this method receives solid line characters as a
// rune array instead of a string.
//
// Be advised that this method does alter the current values for
// left margin and right margin.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	solidLineChars             []rune
//	   - This rune specifies the character or characters which
//	     will comprise the solid line formatted for screen display,
//	     file output or printing.
//
//	     Example:
//	       solidLineChars = []rune{'*'}
//	       solidLineCharsRepeatCount = 5
//	       leftMargin = 0
//	       rightMargin = 0
//
// /         Solid line = "*****"
//
//	      If the array length of 'solidLineChars' is less than one
//	      (1) or greater than one-million (1,000,000), an error will
//	      be returned.
//
//
//	 solidLineCharsRepeatCount  int
//	    - This integer value specifies the number of times that
//	      parameter 'solidLineChars' will be repeated in
//	      constructing the solid line.
//
//	      If this parameter is submitted with a value less than one
//	      (1) or greater than one-million (1,000,000), an error will
//	      be returned.
//
//	      Example:
//	        solidLineChars = []rune{'*'}
//	        solidLineCharsRepeatCount = 5
//	        leftMargin = 0
//	        rightMargin = 0
//	        Solid line = "*****"
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	 err                        error
//		   - If this method completes successfully, the returned error
//		     Type is set equal to 'nil'.
//
//		     If errors are encountered during processing, the returned
//		     error Type will encapsulate an error message. This
//		     returned error message will incorporate the method chain
//		     and text passed by input parameter, 'errorPrefix'. The
//		     'errorPrefix' text will be attached to the beginning of
//		     the error message.
func (txtSpecSolidLine *TextLineSpecSolidLine) SetSolidLineRuneChars(
	solidLineChars []rune,
	solidLineCharsRepeatCount int,
	errorPrefix interface{}) (
	err error) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.SetSolidLineRuneChars()",
		"")

	if err != nil {
		return err
	}

	lenSolidLineChars := len(solidLineChars)

	if lenSolidLineChars == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineChars' is invalid!\n"+
			"'solidLineChars' is a zero length rune array.\n",
			ePrefix.String())

		return err
	}

	if lenSolidLineChars > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineChars' is invalid!\n"+
			"'solidLineChars' has a rune array length greater than "+
			"one-million (1,000,000).\n",
			ePrefix.String())

		return err
	}

	if solidLineCharsRepeatCount == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"'solidLineCharsRepeatCount' has a value of zero (0).\n",
			ePrefix.String())

		return err
	}

	if solidLineCharsRepeatCount < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"'solidLineCharsRepeatCount' has a value less than zero (0).\n"+
			"solidLineCharsRepeatCount = '%v'\n",
			ePrefix.String(),
			solidLineCharsRepeatCount)

		return err
	}

	if solidLineCharsRepeatCount > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineCharsRepeatCount' is invalid!\n"+
			"'solidLineCharsRepeatCount' has a value greater than one-million (1,000,000).\n"+
			"solidLineCharsRepeatCount = '%v'\n",
			ePrefix.String(),
			solidLineCharsRepeatCount)

		return err
	}

	sMechPreon := strMechPreon{}

	_,
		err = sMechPreon.testValidityOfRuneCharArray(
		solidLineChars,
		ePrefix.XCpy(
			"solidLineChars is Invalid!"))

	if err != nil {
		return err
	}

	err = sMechPreon.copyRuneArrays(
		&txtSpecSolidLine.solidLineChars,
		&solidLineChars,
		true,
		ePrefix.XCpy("solidLineChars->"+
			"txtSpecSolidLine.solidLineChars"))

	if err != nil {
		return err
	}

	txtSpecSolidLine.solidLineCharsRepeatCount =
		solidLineCharsRepeatCount

	return err
}

// String - Returns the formatted text generated by this
// Text Line Specification for screen display output and printing.
//
// The current TextLineSpecSolidLine instance will provide the
// necessary information for generating text output which describes
// a solid line of text.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for screen display, file
// output and printing. The use of these types of solid lines with
// text presentations are often referred to as "line breaks" or
// "breaking lines".
//
// A solid line, as used here, consists of a single character or
// multiple characters used in a repeating sequence to construct
// a solid line. Typically, solid lines consist of dashes ("---"),
// underscore characters ("____"), equal signs ("====="), asterisks
// ("*****") and other similar line break presentations. Multiple
// characters may be used to produce more complex line sequences
// ("--*--*--*"). The length of a solid line is specified by the
// calling function using input parameter,
// 'solidLineCharsRepeatCount'.
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is similar to
// TextLineSpecSolidLine.GetFormattedText() with the sole
// difference being that this method does not return an error.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
func (txtSpecSolidLine TextLineSpecSolidLine) String() string {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecSolidLine.String()",
		"")

	formattedText,
		err := textLineSpecSolidLineMolecule{}.ptr().
		getFormattedText(
			&txtSpecSolidLine,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedText
}

// TextBuilder - Configures the line of text produced by this
// instance of TextLineSpecSolidLine, and writes it to an instance
// of strings.Builder.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 strBuilder                 *strings.Builder
//	    - A pointer to an instance of *strings.Builder. The
//	      formatted text characters produced by this method will be
//	      written to this instance of strings.Builder.
//
//
//	errorPrefix                interface{}
//	   - This object encapsulates error prefix text which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods
//	     listed as a method or function chain of execution.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     This empty interface must be convertible to one of the
//	     following types:
//
//	     1. nil - A nil value is valid and generates an empty
//	        collection of error prefix and error context
//	        information.
//
//	     2. string - A string containing error prefix information.
//
//	     3. []string A one-dimensional slice of strings containing
//	        error prefix information
//
//	     4. [][2]string A two-dimensional slice of strings
//	        containing error prefix and error context information.
//
//	     5. ErrPrefixDto - An instance of ErrPrefixDto. Information
//	        from this object will be copied for use in error and
//	        informational messages.
//
//	     6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//	        Information from this object will be copied for use in
//	        error and informational messages.
//
//	     7. IBasicErrorPrefix - An interface to a method generating
//	        a two-dimensional slice of strings containing error
//	        prefix and error context information.
//
//	     If parameter 'errorPrefix' is NOT convertible to one of
//	     the valid types listed above, it will be considered
//	     invalid and trigger the return of an error.
//
//	     Types ErrPrefixDto and IBasicErrorPrefix are included in
//	     the 'errpref' software package,
//	     "github.com/MikeAustin71/errpref".
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
func (txtSpecSolidLine *TextLineSpecSolidLine) TextBuilder(
	strBuilder *strings.Builder,
	errorPrefix interface{}) error {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecSolidLine.TextBuilder()",
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
		err = textLineSpecSolidLineMolecule{}.ptr().
		getFormattedText(
			txtSpecSolidLine,
			ePrefix.XCpy("txtSpecSolidLine"))

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

// TextLineSpecName - returns a string specifying the name
// of this Text Line Specification.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
func (txtSpecSolidLine *TextLineSpecSolidLine) TextLineSpecName() string {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return "SolidLine"
}

// TextTypeName - returns a string specifying the type
// of Text Line Specification.
//
// This method fulfills requirements of ITextSpecification
// interface.
func (txtSpecSolidLine *TextLineSpecSolidLine) TextTypeName() string {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return "TextLineSpecSolidLine"
}

// TurnAutoLineTerminationOff - Turns off or cancels the automatic
// generation of new line terminators for the solid line of text
// produced by this instance of TextLineSpecSolidLine.
//
// By default, a line termination sequence (usually a new line
// character '\n') is automatically appended to the end of each
// solid line of text produced by an instance of
// TextLineSpecSolidLine.
//
// By calling this method (TurnAutoLineTerminationOff()) the line
// termination sequenced is cancelled and will NOT be applied to
// the solid line of text generated by this instance of
// TextLineSpecSolidLine.
//
// Calling this method will set the member variable
// 'turnLineTerminatorOff' to 'true'. Consequently, line
// terminators as defined by member variable 'newLineChars' WILL
// NOT be applied as a line termination sequence for the line of
// text produced by this instance of TextLineSpecPlainText.
//
// In contrast, the method
// TextLineSpecPlainText.TurnAutoLineTerminationOn() will always
// set member variable 'turnLineTerminatorOff' to 'false'. This
// means that the automatic generation of line terminators is
// 'turned on' and a line termination sequence WILL BE applied for
// the line of text produced by TextLineSpecPlainText.
//
// The default line terminator is the new line character ('\n')
// which is defined by member variable 'newLineChars'. However,
// this value is subject to user control and may be overridden by
// one or more characters supplied by the user.
//
// Again, this method (TurnAutoLineTerminationOff()) will always
// set 'turnLineTerminatorOff' to 'true'.  This means that line
// terminators WILL NOT BE applied to each line of text produced
// by this instance of TextLineSpecPlainText.
func (txtSpecSolidLine *TextLineSpecSolidLine) TurnAutoLineTerminationOff() {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	txtSpecSolidLine.turnLineTerminatorOff = true
}

// TurnAutoLineTerminationOn - Turns on or engages the automatic
// generation of new line terminators for the solid line of text
// produced by this instance of TextLineSpecSolidLine.
//
// This means that a new line termination sequence with characters
// defined by member variable 'newLineChars' WILL BE appended to
// the end of the text line produced by the current instance of
// TextLineSpecSolidLine.
//
// This method will set the member variable 'turnLineTerminatorOff'
// is set to 'false'. Consequently, line terminators as defined by
// member variable 'newLineChars' WILL BE applied as a line
// termination sequence for each line of text produced by
// TextLineSpecSolidLine.
//
// In contrast, method
// TextLineSpecPlainText.TurnAutoLineTerminationOff() will always
// set member variable 'turnLineTerminatorOff' to 'true'. This
// means that the automatic generation of line terminators is
// 'turned off' and a line termination sequence WILL NOT BE applied
// for the line of text produced by TextLineSpecPlainText.
//
// The default line terminator is the new line character ('\n')
// which is defined by member variable 'newLineChars'. However,
// this value is subject to user control and may be overridden by
// one or more characters supplied by the user.
//
// Again, this method will always set 'turnLineTerminatorOff' to
// 'false' meaning that line terminators WILL BE applied to
// the line of text produced by this instance of TextLineSpecPlainText.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// Since the default behavior for creation of TextLineSpecSolidLine
// instances is "Turn Auto Line Termination On", you will never
// need to call this method unless you previously called the
// method: TextLineSpecSolidLine.TurnAutoLineTerminationOff().
func (txtSpecSolidLine *TextLineSpecSolidLine) TurnAutoLineTerminationOn() {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	txtSpecSolidLine.turnLineTerminatorOff = false
}

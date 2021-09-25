package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextLineSpecSolidLine - Used to create a solid line of text
// characters for text display, file output or printing.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
// character ('\n'). This means that each instance of a solid will
// be formatted as a single line of text. If another new line
// character or characters is required, reference method:
//    TextLineSpecSolidLine.SetNewLineCharsOverride()
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         rightMargin = 3
//         Solid line = "   *****   "
//
type TextLineSpecSolidLine struct {
	leftMargin                int
	rightMargin               int
	solidLineChars            []rune
	solidLineCharsRepeatCount int
	newLineChars              []rune
	lock                      *sync.Mutex
}

// CopyIn - Copies all the data fields from an incoming instance of
// TextLineSpecSolidLine ('incomingTxtSolidLine') to the data fields
// of the current TextLineSpecSolidLine instance
// ('txtSpecSolidLine').
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the data fields in current TextLineSpecSolidLine instance
// ('txtSpecSolidLine') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingTxtSolidLine       *TextLineSpecSolidLine
//     - A pointer to an instance of TextLineSpecSolidLine. This
//       method will NOT change the member variable data values
//       contained in this instance.
//
//       All data values in this TextLineSpecSolidLine instance
//       will be copied to current TextLineSpecSolidLine
//       instance ('txtSpecSolidLine').
//
//       If 'incomingTxtSolidLine' contains invalid member data
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
//  TextLineSpecSolidLine
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLineSpecSolidLine instance.
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
//       current TextLineSpecSolidLine instance cast as an
//       ITextLineSpecification object.
//
//
//  error
//     - If the method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
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

	return ITextLineSpecification(&newTxtSolidLine), nil
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextLineSpecSolidLine instance.
//
// If the current TextLineSpecSolidLine instance contains invalid
// member variable data values, this method will return an error.
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
//  *TextLineSpecSolidLine
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a pointer to a
//       deep copy of the current TextLineSpecSolidLine
//       instance.
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
//
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
//
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
//
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
// Text Line Specification for text display output and printing.
//
// The current TextLineSpecSolidLine instance will provide the
// necessary information for generating text output which describes
// a solid line of text.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
//       of TextLineSpecSolidLine.
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
//       the current instance of TextLineSpecSolidLine are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
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
		"TextLineSpecSolidLine.IsValidInstanceError()",
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

// NewSolidLine - Creates and returns a new instance of
// TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
//    TextLineSpecSolidLine.SetNewLineCharsOverride()
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  line termination character  - Defaults to new line character
//                                ('\n')
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//       If this value is less than zero (0), 'leftMargin' will be
//       set to a default value of zero (0).
//
//       If the 'leftMargin' value is greater than one-million
//       (1,000,000), an error will be returned.
//
//
//  rightMargin                 int
//     - The number of white space characters appended to the
//       end, or right side, of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//       If this value is less than zero (0), 'rightMargin' will be
//       set to a default value of zero (0).
//
//       If the 'rightMargin' value is greater than one-million
//       (1,000,000), an error will be returned.
//
//
//  solidLineChars             string
//     - This string specifies the character or characters which
//       will comprise the solid line output for text display or
//       printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//       If this parameter is submitted as a zero length string, an
//       error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
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
//  TextLineSpecSolidLine
//     - If this method completes successfully, it will create and
//       return a new instance of TextLineSpecSolidLine which is
//       fully configured with all the parameters necessary to
//       format a solid line of characters for text display output
//       or printing.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine TextLineSpecSolidLine) NewSolidLine(
	leftMargin int,
	rightMargin int,
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
		"TextLineSpecSolidLine.NewSolidLine()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			leftMargin,
			rightMargin,
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			ePrefix.XCtx("txtSolidLine"))

	return txtSolidLine, err
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
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
//    TextLineSpecSolidLine.SetNewLineCharsOverride()
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  leftMargin                  - Defaults to zero (0)
//
//  rightMargin                 - Defaults to zero (0)
//
//  line termination character  - Defaults to new line character
//                                ('\n')
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  solidLineChars             string
//     - This string specifies the character or characters which
//       will comprise the solid line output for text display or
//       printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//       If this parameter is submitted as a zero length string, an
//       error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
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
//  TextLineSpecSolidLine
//     - If this method completes successfully, it will create and
//       return a new instance of TextLineSpecSolidLine which is
//       fully configured with all the parameters necessary to
//       format a solid line of characters for text display output
//       or printing.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine TextLineSpecSolidLine) NewDefaultSolidLine(
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
		"TextLineSpecSolidLine.NewDefaultSolidLine()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			0,
			0,
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			ePrefix.XCtx("txtSolidLine"))

	return txtSolidLine, err
}

// NewFullSolidLineConfig - Creates and returns a new instance of
// TextLineSpecSolidLine. This method requires more input
// parameters than other similar methods, but in return, it allows
// the user to exercise control over all facets of the
// TextLineSpecSolidLine configuration.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//       If this value is less than zero (0), 'leftMargin' will be
//       set to a default value of zero (0).
//
//       If the 'leftMargin' value is greater than one-million
//       (1,000,000), an error will be returned.
//
//
//  rightMargin                 int
//     - The number of white space characters appended to the
//       end, or right side, of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//       If this value is less than zero (0), 'rightMargin' will be
//       set to a default value of zero (0).
//
//       If the 'rightMargin' value is greater than one-million
//       (1,000,000), an error will be returned.
//
//
//  solidLineChars             string
//     - This string specifies the character or characters which
//       will comprise the solid line output for text display or
//       printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//       If this parameter is submitted as a zero length string, an
//       error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//
//  newLineChars               string
//     - This string contains one or more characters which will
//       be used to terminate the solid text line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         newLineChars = "??\n\n"
//         Solid line = "*****??\n\n"
//
//       If this parameter is submitted as a zero length rune
//       array, 'newLineChars' will be set to the default new
//       line character ('\n').
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
//  TextLineSpecSolidLine
//     - If this method completes successfully, it will create and
//       return a new instance of TextLineSpecSolidLine which is
//       fully configured with all the parameters necessary to
//       format a solid line of characters for text display output
//       or printing.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine TextLineSpecSolidLine) NewFullSolidLineConfig(
	leftMargin int,
	rightMargin int,
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
		"TextLineSpecSolidLine.NewFullSolidLineConfig()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			leftMargin,
			rightMargin,
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCtx("txtSolidLine"))

	return txtSolidLine, err
}

// NewFullSolidLineRunesConfig - Creates and returns a new instance of
// TextLineSpecSolidLine. This method requires more input
// parameters than other similar methods, but in return, it allows
// the user to exercise control over all facets of the
// TextLineSpecSolidLine configuration.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         rightMargin = 0
//         Solid line = "   *****"
//
//       If this value is less than zero (0), 'leftMargin' will be
//       set to a default value of zero (0).
//
//       If this value is greater than one-million (1,000,000), an
//       error will be returned.
//
//
//  rightMargin                 int
//     - The number of white space characters appended to the
//       end, or right side, of the solid line.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//       If this value is less than zero (0), 'rightMargin' will be
//       set to a default value of zero (0).
//
//       If this value is greater than one-million (1,000,000), an
//       error will be returned.
//
//
//  solidLineChars             []rune
//     - This rune array specifies the character or characters
//       which will comprise the solid line output for text display
//       or printing.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 0
///        Solid line = "*****"
//
//       If this parameter is submitted as a zero length rune
//       array, an error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 0
//         Solid line = "*****"
//
//
//  newLineChars               []rune
//     - This rune array contains one or more characters which will
//       be used to terminate the solid text line.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 0
//         newLineChars =  []rune{'?','?','\n','\n'}
//         Solid line = "*****??\n\n"
//
//       If this parameter is submitted as a zero length rune
//       array, 'newLineChars' will be set to the default new
//       line character ('\n').
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
//  TextLineSpecSolidLine
//     - If this method completes successfully, it will create and
//       return a new instance of TextLineSpecSolidLine which is
//       fully configured with all the parameters necessary to
//       format a solid line of characters for text display output
//       or printing.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine TextLineSpecSolidLine) NewFullSolidLineRunesConfig(
	leftMargin int,
	rightMargin int,
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
		"TextLineSpecSolidLine.NewFullSolidLineRunesConfig()",
		"")

	if err != nil {
		return txtSolidLine, err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			&txtSolidLine,
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCtx("txtSolidLine"))

	return txtSolidLine, err
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
// specifications for solid lines output for text display, file
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
// IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecSolidLine instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0). If this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  rightMargin                 int
//     - The number of white space characters appended to the
//       end, or right side, of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0). If this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  solidLineChars             string
//     - This string specifies the character or characters which
//       will comprise the solid line output for text display or
//       printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//       If this parameter is submitted as a zero length string, an
//       error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//
//  newLineChars               string
//     - This string contains one or more characters which will be
//       used to terminate the solid text line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         newLineChars = "??\n\n"
//         Solid line = "*****??\n\n"
//
//       If this parameter is submitted as a zero length string
//       it will be set to the default new line character ('\n').
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
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine *TextLineSpecSolidLine) SetFullSolidLineConfig(
	leftMargin int,
	rightMargin int,
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
		"TextLineSpecSolidLine.SetFullSolidLineConfig()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			txtSpecSolidLine,
			leftMargin,
			rightMargin,
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCtx("txtSpecSolidLine"))

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
// specifications for solid lines output for text display, file
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
// IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecSolidLine instance.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0). If this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  rightMargin                 int
//     - The number of white space characters appended to the
//       end, or right side, of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0). If this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  solidLineChars             []rune
//     - This rune array specifies the character or characters
//       which will comprise the solid line output for text display
//       or printing.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//       If this parameter is submitted as a zero length rune
//       array, an error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//
//  newLineChars               []rune
//     - This rune array contains one or more characters which will
//       be used to terminate the solid text line.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         newLineChars =  []rune{'?','?','\n',\n'}
//         Solid line = "*****??\n\n"
//
//       If this parameter is submitted as a zero length rune
//       array, 'newLineChars' will be set to the default new
//       line character ('\n').
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
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine *TextLineSpecSolidLine) SetFullSolidLineRunesConfig(
	leftMargin int,
	rightMargin int,
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
		"TextLineSpecSolidLine.SetFullSolidLineRunesConfig()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecSolidLineMolecule{}.ptr().
		setTxtSolidLine(
			txtSpecSolidLine,
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCtx("txtSpecSolidLine"))

	return err
}

// SetLeftMargin - Sets the left margin for the current instance
// of TextLineSpecSolidLine.
//
// The left margin defines the number of white space characters
// which will be inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//       If this value is less than zero (0), 'leftMargin' will be
//       set to a default value of zero (0).
//
//       If the 'leftMargin' value is greater than one-million
//       (1,000,000), an error will be returned.
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
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine *TextLineSpecSolidLine) SetLeftMargin(
	leftMargin int,
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

	if leftMargin < 0 {
		leftMargin = 0
	}

	if leftMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'leftMargin' is invalid!\n"+
			"The integer value of 'leftMargin' is greater than 1,000,000.\n"+
			"leftMargin='%v'\n",
			ePrefix.String(),
			leftMargin)

		return err
	}

	txtSpecSolidLine.leftMargin = leftMargin

	return
}

// SetNewLineCharsOverride - Sets the new line character or
// characters which will be used to terminate the solid line
// generated by this instance of TextLineSpecSolidLine.
//
// By default, all solid lines of text are terminated by a new line
// character ('\n). This means that each instance of a solid will
// be formatted as a single line of text.
//
// The new line character or characters may be customized by
// calling this method.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  newLineChars               string
//     - This rune array contains one or more characters which will
//       be used to terminate the solid text line.
//
//       This example creates a line termination followed by a
//       second blank line:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         newLineChars = "??\n\n"
//         Solid line = "*****??\n"
//                      "\n"
//
//       If this parameter is submitted as a zero length rune
//       array, 'newLineChars' will be set to the default new
//       line character ('\n').
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  --- NONE ---
//
func (txtSpecSolidLine *TextLineSpecSolidLine) SetNewLineCharsOverride(
	newLineChars string) {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	if len(newLineChars) == 0 {
		newLineChars = "\n"
	}

	sMechPreon := strMechPreon{}

	newLineRunes := []rune(newLineChars)

	_ = sMechPreon.copyRuneArrays(
		&txtSpecSolidLine.newLineChars,
		&newLineRunes,
		true,
		nil)
}

// SetRightMargin - Sets the right margin for the current instance
// of TextLineSpecSolidLine.
//
// The right margin defines the number of white space characters
// which will be appended to the end, or right side, of the solid
// line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  rightMargin                 int
//     - The number of white space characters appended to the
//       end, or right side, of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//       If this value is less than zero (0), 'rightMargin' will be
//       set to a default value of zero (0).
//
//       If the 'rightMargin' value is greater than one-million
//       (1,000,000), an error will be returned.
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
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine *TextLineSpecSolidLine) SetRightMargin(
	rightMargin int,
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

	if rightMargin < 0 {
		rightMargin = 0
	}

	if rightMargin > 1000000 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'rightMargin' is invalid!\n"+
			"The integer value of 'rightMargin' is greater than 1,000,000.\n"+
			"leftMargin='%v'\n",
			ePrefix.String(),
			rightMargin)

		return err
	}

	txtSpecSolidLine.rightMargin = rightMargin

	return
}

// SetSolidLine - This method configures the current
// TextLineSpecSolidLine instance with new solid line parameters.
// All internal member variable data values for the current
// TextLineSpecSolidLine instance will be reset and overwritten.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
//    TextLineSpecSolidLine.SetNewLineCharsOverride()
//
// ------------------------------------------------------------------------
//
// IMPORTANT
//
// This method will overwrite and reset all the data values
// encapsulated by the current TextLineSpecSolidLine instance.
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  line termination character  - Defaults to new line character
//                                ('\n')
//
//
// ------------------------------------------------------------------------
//
//
// Input Parameters
//
//  leftMargin                 int
//     - The number of white space characters which will be
//       inserted on the left side of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 3
//         Solid line = "   *****"
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0). If this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  rightMargin                 int
//     - The number of white space characters appended to the
//       end, or right side, of the solid line.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 3
//         Solid line = "*****   "
//
//       If this value is less than zero (0), it will be set to a
//       default value of zero (0). If this value is greater than
//       one-million (1,000,000), an error will be returned.
//
//
//  solidLineChars             string
//     - This string specifies the character or characters which
//       will comprise the solid line output for text display or
//       printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//       If this parameter is submitted as a zero length string, an
//       error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = '*'
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
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
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (txtSpecSolidLine *TextLineSpecSolidLine) SetSolidLine(
	leftMargin int,
	rightMargin int,
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
			leftMargin,
			rightMargin,
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune{'\n'},
			ePrefix.XCtx("txtSpecSolidLine"))

	return err
}

// SetSolidLineStrChars - Sets the solid line character or
// characters for the current instance of TextLineSpecSolidLine.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  solidLineChars             string
//     - This string specifies the character or characters which
//       will comprise the solid line output for text display or
//       printing.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
//
//       If this parameter is submitted as a zero length string, an
//       error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = "*"
//         solidLineCharsRepeatCount = 5
//         Solid line = "*****"
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
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
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

	if len(solidLineChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineChars' is invalid!\n"+
			"'solidLineChars' is a zero length string.\n",
			ePrefix.String())

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

	solidLineRunes := []rune(solidLineChars)

	err = strMechPreon{}.ptr().copyRuneArrays(
		&txtSpecSolidLine.solidLineChars,
		&solidLineRunes,
		true,
		ePrefix.XCtx("solidLineChars->"+
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
// specifications for solid lines output for text display, file
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
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  solidLineChars             []rune
//     - This rune specifies the character or characters which
//       will comprise the solid line formatted for text display,
//       file output or printing.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 0
///         Solid line = "*****"
//
//       If this parameter is submitted as a zero length string, an
//       error will be returned.
//
//
//  solidLineCharsRepeatCount  int
//     - This integer value specifies the number of times that
//       parameter 'solidLineChars' will be repeated in
//       constructing the solid line.
//
//       If this parameter is submitted with a value less than one
//       (1), an error will be returned.
//
//       Example:
//         solidLineChars = []rune{'*'}
//         solidLineCharsRepeatCount = 5
//         leftMargin = 0
//         rightMargin = 0
//         Solid line = "*****"
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
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
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

	if len(solidLineChars) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineChars' is invalid!\n"+
			"'solidLineChars' is a zero length rune array.\n",
			ePrefix.String())

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

	err = strMechPreon{}.ptr().copyRuneArrays(
		&txtSpecSolidLine.solidLineChars,
		&solidLineChars,
		true,
		ePrefix.XCtx("solidLineChars->"+
			"txtSpecSolidLine.solidLineChars"))

	if err != nil {
		return err
	}

	txtSpecSolidLine.solidLineCharsRepeatCount =
		solidLineCharsRepeatCount

	return err
}

// String - Returns the formatted text generated by this
// Text Line Specification for text display output and printing.
//
// The current TextLineSpecSolidLine instance will provide the
// necessary information for generating text output which describes
// a solid line of text.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display, file
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
//
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

// TextTypeName - returns a string specifying the type
// of Text Line Specification.
//
// This method fulfills requirements of ITextSpecification
// interface.
//
func (txtSpecSolidLine TextLineSpecSolidLine) TextTypeName() string {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return "TextLineSpecSolidLine"
}

// TextLineSpecName - returns a string specifying the name
// of this Text Line Specification.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
func (txtSpecSolidLine TextLineSpecSolidLine) TextLineSpecName() string {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	return "TextLineSpecStandardLine"
}

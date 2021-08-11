package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextLineSpecSolidLine - Used to create a solid line of text
// characters for text display or printing.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display and
// printing. A solid line consists of a single character or
// multiple characters which constitute a solid line and are often
// used for line breaks. Typically, solid lines consist of dashes
// ("---"), underscore characters ("____"), equal signs ("=====")
// or asterisks ("*****"). Multiple characters may be used to
// produce different line sequences ("--*--*--*"). The length of a
// solid is specified by the calling function using input
// parameter 'solidLineCharsRepeatCount'.
//
// By default, all solid lines are terminated by a new line
// character. This means that each instance of a solid will be
// formatted as a single line of text. If another new line
// character or characters is required, reference method:
//    TextLineSpecSolidLine.SetNewLineCharsOverride()
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
// specifications for solid lines output for text display and
// printing. A solid line consists of a single character or
// multiple characters which constitute a solid line and are often
// used for line breaks. Typically, solid lines consist of dashes
// ("---"), underscore characters ("____"), equal signs ("=====")
// or asterisks ("*****"). Multiple characters may be used to
// produce different line sequences ("--*--*--*"). The length of a
// solid is specified by the calling function using input
// parameter 'solidLineCharsRepeatCount'. The use of these types of
// solid lines with text display are often described as
// "line breaks" or "breaking lines".
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is identical in function to
// TextLineSpecSolidLine.String()
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (txtSpecSolidLine *TextLineSpecSolidLine) GetFormattedText() string {

	if txtSpecSolidLine.lock == nil {
		txtSpecSolidLine.lock = new(sync.Mutex)
	}

	txtSpecSolidLine.lock.Lock()

	defer txtSpecSolidLine.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecSolidLine.GetFormattedText()",
		"")

	formattedText,
		err := textLineSpecSolidLineMolecule{}.ptr().
		getFormattedText(
			txtSpecSolidLine,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedText
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
// specifications for solid lines output for text display and
// printing. A solid line consists of a single character or
// multiple characters which constitute a solid line and are often
// used for line breaks. Typically, solid lines consist of dashes
// ("---"), underscore characters ("____"), equal signs ("=====")
// or asterisks ("*****"). Multiple characters may be used to
// produce different line sequences ("--*--*--*"). The length of a
// solid is specified by the calling function using input
// parameter 'solidLineCharsRepeatCount'. The use of these types of
// solid lines with text display are often described as
// "line breaks" or "breaking lines".
//
// By default, all solid lines of text are terminated with a new
// line character ('\n'). This means that each instance of a solid
// will be formatted as a single line of text. If another new line
// character or characters are required, reference method:
//    TextLineSpecSolidLine.SetNewLineCharsOverride()
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
//       If this parameter is submitted as a zero length rune array,
//       or if 'newLineChars' contains invalid characters, it will
//       be set to the default new line character ('\n').
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
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
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
//       If this parameter is submitted as a zero length rune array,
//       it will by default be set to a new line character ('\n').
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

// String - Returns the formatted text generated by this
// Text Line Specification for text display output and printing.
//
// The current TextLineSpecSolidLine instance will provide the
// necessary information for generating text output which describes
// a solid line of text.
//
// The TextLineSpecSolidLine type provides formatting
// specifications for solid lines output for text display and
// printing. A solid line consists of a single character or
// multiple characters which constitute a solid line and are often
// used for line breaks. Typically, solid lines consist of dashes
// ("---"), underscore characters ("____"), equal signs ("=====")
// or asterisks ("*****"). Multiple characters may be used to
// produce different line sequences ("--*--*--*"). The length of a
// solid is specified by the calling function using input
// parameter 'solidLineCharsRepeatCount'. The use of these types of
// solid lines with text display are often described as
// "line breaks" or "breaking lines".
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is identical in function to
// TextLineSpecSolidLine.GetFormattedText()
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

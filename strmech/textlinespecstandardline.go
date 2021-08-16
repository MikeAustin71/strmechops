package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// TextLineSpecStandardLine - This type is a text specification for
// a standard line of text. It encapsulates an array of
// ITextFieldSpecification objects which are used to format text
// fields within a single line of text. Essentially, a standard
// text line is a collection of text fields which implement the
// ITextFieldSpecification interface.
//
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// By default, the TextLineSpecStandardLine type will generate a
// line of text with a "new line" ('\n') line termination
// character. The application of this "new line" ('\n') character
// is automatic and by default.
//
// However, users may control and override the application of line
// termination characters through the following methods:
//
//  TextLineSpecStandardLine.SetLineTerminationChars()
//  TextLineSpecStandardLine.TurnAutoLineTerminationOff()
//  TextLineSpecStandardLine.TurnAutoLineTerminationOn()
//
//
// The TextLineSpecStandardLine type implements the
// ITextLineSpecification interface.
//
//
// Member Variables
//
// ----------------------------------------------------------------
//
//
//
type TextLineSpecStandardLine struct {
	textFields            []ITextFieldSpecification
	numOfStdLines         int
	turnLineTerminatorOff bool
	newLineChars          []rune
	lock                  *sync.Mutex
}

// AddTextField - This method will append a text field object to
// the end of the current array of text field objects maintained by
// the current instance of TextLineSpecStandardLine. The object
// actually appended to the array is a deep copy of the input
// parameter, 'textField'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textField                  ITextFieldSpecification
//     - A text field object which implements the
//       ITextFieldSpecification interface. A deep copy of this
//       object will be added to the text field collection
//       maintained by this instance of TextLineSpecStandardLine.
//
//       If member variable data values contained in this
//       'textField' parameter are found to be invalid, an error
//       will be returned.
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
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (stdLine *TextLineSpecStandardLine) AddTextField(
	textField ITextFieldSpecification,
	errorPrefix interface{}) (
	err error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.AddTextField()",
		"")

	if err != nil {
		return err
	}

	err = textField.IsValidInstanceError(
		ePrefix.XCtx("textField"))

	if err != nil {
		return err
	}

	var newTextField ITextFieldSpecification

	newTextField,
		err = textField.CopyOutITextField(
		ePrefix.XCtx("textField->newTextField"))

	if err != nil {
		return err
	}

	stdLine.textFields = append(stdLine.textFields,
		newTextField)

	return err
}

// CopyIn - Copies the data fields from an incoming instance of
// TextLineSpecStandardLine ('incomingStdLine') to the data fields
// of the current TextLineSpecStandardLine instance ('stdLine').
//
// IMPORTANT
// All the data fields in current TextLineSpecStandardLine instance
// ('stdLine') will be modified and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingStdLine     *TextLineSpecStandardLine
//     - A pointer to an instance of TextLineSpecStandardLine. This
//       method will NOT change the values of internal member
//       variables contained in this instance.
//
//       All data values in this TextLineSpecStandardLine instance
//       will be copied to current TextLineSpecStandardLine
//       instance ('stdLine').
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
func (stdLine *TextLineSpecStandardLine) CopyIn(
	incomingStdLine *TextLineSpecStandardLine,
	errorPrefix interface{}) error {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.CopyIn()",
		"")

	if err != nil {
		return err
	}

	err = textLineSpecStandardLineMolecule{}.ptr().
		copyIn(
			stdLine,
			incomingStdLine,
			ePrefix)

	return err
}

// CopyOut - Returns a deep copy of the current
// TextLineSpecStandardLine instance.
//
// If the current TextLineSpecStandardLine instance contains invalid
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
//  TextLineSpecStandardLine
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current TextLineSpecStandardLine instance.
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
func (stdLine *TextLineSpecStandardLine) CopyOut(
	errorPrefix interface{}) (
	TextLineSpecStandardLine,
	error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.CopyOut()",
		"")

	if err != nil {
		return TextLineSpecStandardLine{}, err
	}

	var newStdLine TextLineSpecStandardLine

	newStdLine,
		err = textLineSpecStandardLineMolecule{}.ptr().
		copyOut(
			stdLine,
			ePrefix.XCtx("stdLine"))

	return newStdLine, err
}

// CopyOutITextLine - Returns a deep copy of the current
// TextLineSpecStandardLine instance cast as a type
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
//       current TextLineSpecStandardLine instance cast as an
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
func (stdLine *TextLineSpecStandardLine) CopyOutITextLine(
	errorPrefix interface{}) (
	ITextLineSpecification,
	error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.CopyOutITextLine()",
		"")

	if err != nil {
		return ITextLineSpecification(&TextLineSpecStandardLine{}), err
	}

	var newStdLine TextLineSpecStandardLine

	newStdLine,
		err = textLineSpecStandardLineMolecule{}.ptr().
		copyOut(
			stdLine,
			ePrefix.XCtx("stdLine"))

	return ITextLineSpecification(&newStdLine), err
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextLineSpecStandardLine instance.
//
// If the current TextLineSpecStandardLine instance contains invalid
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
//  *TextLineSpecStandardLine
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a pointer to a
//       deep copy of the current TextLineSpecStandardLine
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
func (stdLine *TextLineSpecStandardLine) CopyOutPtr(
	errorPrefix interface{}) (
	*TextLineSpecStandardLine,
	error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.CopyOutPtr()",
		"")

	if err != nil {
		return &TextLineSpecStandardLine{}, err
	}

	var newStdLine TextLineSpecStandardLine

	newStdLine,
		err = textLineSpecStandardLineMolecule{}.ptr().
		copyOut(
			stdLine,
			ePrefix.XCtx("stdLine"))

	return &newStdLine, err
}

// Empty - Deletes all the text fields stored as an array of
// ITextFieldSpecification pointers within the current
// TextLineSpecStandardLine instance. In addition, this method
// will set 'numOfStdLines', 'turnLIneTerminatorOff' and
// 'newLineChars' will be set to their initial or zero values.
//
// After calling 'Empty', the caller CAN NOT reuse this instance of
// TextLineSpecStandardLine.
//
// To empty and reuse this TextLineSpecStandardLine instance,
// reference method TextLineSpecStandardLine.EmptyTextFields.
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (stdLine *TextLineSpecStandardLine) Empty() {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	textLineSpecStandardLineMolecule{}.ptr().
		empty(
			stdLine)

	stdLine.lock.Unlock()

	stdLine.lock = nil

	return
}

// EmptyTextFields - Deletes all the text fields stored as an array of
// ITextFieldSpecification pointers within the current
// TextLineSpecStandardLine instance.
//
// This method will allow the user to clear and reuse the current
// instance of TextLineSpecStandardLine because it only deletes the
// existing text field collection. It does not overwrite member
// variables 'numOfStdLines', 'turnLIneTerminatorOff' or
// 'newLineChars'.
//
func (stdLine *TextLineSpecStandardLine) EmptyTextFields() {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	textLineSpecStandardLineElectron{}.ptr().
		emptyTextFields(stdLine)

	return
}

// Equal - Receives a pointer to another instance of
// TextLineSpecStandardLine and proceeds to compare the member
// variables to those of the current TextLineSpecStandardLine
// instance in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables of both instances are equal in
// all respects, this flag is set to 'true'. Otherwise, this method
// returns 'false'.
//
func (stdLine *TextLineSpecStandardLine) Equal(
	incomingStdLine *TextLineSpecStandardLine) bool {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	return textLineSpecStandardLineMolecule{}.ptr().
		equal(
			stdLine,
			incomingStdLine)
}

// EqualITextLine - Receives an object implementing the
// ITextLineSpecification interface and proceeds to compare
// the member variables to those of the current
// TextLineSpecStandardLine instance in order to determine if
// they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables from both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
// This method is required by interface ITextLineSpecification.
//
func (stdLine *TextLineSpecStandardLine) EqualITextLine(
	iTextLine ITextLineSpecification) bool {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	if stdLine.numOfStdLines < 0 {
		stdLine.numOfStdLines = 0
	}

	stdLineTwo, ok := iTextLine.(*TextLineSpecStandardLine)

	if !ok {
		return false
	}

	return textLineSpecStandardLineMolecule{}.ptr().
		equal(
			stdLine,
			stdLineTwo)
}

// GetFormattedText - Returns the formatted text generated by this
// Text Line Specification  for text display output and printing.
//
// The standard line may be replicated multiple times if the
// value of internal member variable' stdLine.numOfStdLines' is
// greater than one ('1').
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is identical in function to
// TextLineSpecStandardLine.String()
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (stdLine *TextLineSpecStandardLine) GetFormattedText() string {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecStandardLine.GetFormattedText()",
		"")

	var formattedText string
	var err error

	formattedText,
		err = textLineSpecStandardLineMolecule{}.ptr().
		getFormattedText(
			stdLine,
			&ePrefix)

	if err != nil {
		formattedText = fmt.Sprintf("%v\n",
			err.Error())
	}

	return formattedText
}

// GetNumOfStdLines - Returns the number of repetitions for this
// instance of TextLineSpecStandardLine. The number of standard
// lines is the number of times this standard line will be output
// printed.
//
func (stdLine *TextLineSpecStandardLine) GetNumOfStdLines() int {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	return stdLine.numOfStdLines
}

// GetNumOfTextFields - Returns the number of text fields
// encapsulated by the current TextLineSpecStandardLine instance.
//
// Text Fields constitute the granular elements of a standard text
// line. Analyzing the number of text fields in the collection
// provides verification that text fields exist and are ready for
// formatting. Once properly formatted, text fields may be
// presented for text display, file output or printing.
//
func (stdLine *TextLineSpecStandardLine) GetNumOfTextFields() int {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	return len(stdLine.textFields)
}

// GetTextFields - Returns a deep copy of the text fields contained
// in the current TextLineSpecStandardLine instance.
//
// These text fields are returned in an array of
// ITextFieldSpecification objects.
//
// If the text field collection maintained by the current
// TextLineSpecStandardLine instance is empty (contains zero
// elements), an error will be returned.
//
// If any of the text fields within the collection maintained by
// the current TextLineSpecStandardLine instance are invalid,
// an error will be returned.
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
//  []ITextFieldSpecification
//     - If this method completes successfully, a deep copy of the
//       text field collection maintained by the current
//       TextLineSpecStandardLine instance will be returned. These
//       text fields are returned as an array of objects
//       implementing the ITextFieldSpecification interface.
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
func (stdLine *TextLineSpecStandardLine) GetTextFields(
	errorPrefix interface{}) (
	[]ITextFieldSpecification,
	error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.GetTextFields()",
		"")

	if err != nil {
		return nil, err
	}

	lenTxtFields := len(stdLine.textFields)

	if lenTxtFields == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The text fields collection is empty!\n"+
			"TextLineSpecStandardLine.newTextFields contains zero text field objects!\n",
			ePrefix.String())

		return nil, err
	}

	newTextFields := make([]ITextFieldSpecification, lenTxtFields)

	for i := 0; i < lenTxtFields; i++ {

		if stdLine.textFields[i] == nil {
			err = fmt.Errorf("%v\n"+
				"Error: Text Field element stdLine.textFields[%v]\n"+
				"has a 'nil' value!\n",
				ePrefix.String())

			return nil, err
		}

		err = stdLine.textFields[i].IsValidInstanceError(
			ePrefix.XCtx(
				fmt.Sprintf(
					"stdLine.newTextFields[%v] invalid",
					i)))

		if err != nil {
			return nil, err
		}

		newTextField,
			err2 := stdLine.textFields[i].CopyOutITextField(
			ePrefix.XCtx(
				fmt.Sprintf(
					"stdLine.newTextFields[%v] copy error",
					i)))

		if err2 != nil {
			return nil, err2
		}

		newTextFields[i] = newTextField
	}

	return newTextFields, err
}

// IsValidInstanceError - Performs a diagnostic review of the data
// values encapsulated in the current TextLineSpecStandardLine
// instance to determine if they are valid.
//
// If any data element evaluates as invalid, this method will
// return an error.
//
// If the number of standard lines for the current
// TextLineSpecStandardLine is set to a value less than one, this
// method will consider the current TextLineSpecStandardLine
// instance invalid and return an error.
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
//       the current instance of TextLineSpecStandardLine are found
//       to be invalid, this method will return an error.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' (error prefix) will be inserted or
//       prefixed at the beginning of the error message.
//
func (stdLine *TextLineSpecStandardLine) IsValidInstanceError(
	errorPrefix interface{}) error {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.IsValidInstanceError()",
		"")

	if err != nil {
		return err
	}

	_,
		err = textLineSpecStandardLineElectron{}.ptr().
		testValidityOfTextLineSpecStdLine(
			stdLine,
			ePrefix.XCtx("stdLine"))

	return err
}

// New - Returns a new and empty, or unpopulated, concrete instance
// of TextLineSpecStandardLine.
//
// The default number of lines repetitions is set to one ('1').
// This means that the returned standard line instance will only
// be output or printed once.
//
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  numOfStdLines  - Defaults number of standard lines to one (1).
//
//  newLineChars   - Defaults new line character to '\n'.
//                   To change the value of the new line character,
//                   or characters, see methods:
//                   TextLineSpecStandardLine.SetLineTerminationChars()
//                   TextLineSpecStandardLine.SetLineTerminationRunes()
//
func (stdLine TextLineSpecStandardLine) New() TextLineSpecStandardLine {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	newStdLine := TextLineSpecStandardLine{}

	newStdLine.lock = new(sync.Mutex)

	newStdLine.numOfStdLines = 1

	newStdLine.newLineChars = []rune{'\n'}

	newStdLine.turnLineTerminatorOff = false

	return newStdLine
}

// NewPtr - This method returns a pointer to an empty or
// unpopulated instance of TextLineSpecStandardLine.
//
// The default number of lines repetitions is set to one ('1').
// This means that the returned standard line instance will only
// be output or printed once.
//
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  numOfStdLines  - Defaults number of standard lines to one (1).
//
//  newLineChars   - Defaults new line character to '\n'.
//                   To change the value of the new line character,
//                   or characters, see methods:
//                   TextLineSpecStandardLine.SetLineTerminationChars()
//                   TextLineSpecStandardLine.SetLineTerminationRunes()
//
func (stdLine TextLineSpecStandardLine) NewPtr() *TextLineSpecStandardLine {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	newStdLine := TextLineSpecStandardLine{}

	newStdLine.lock = new(sync.Mutex)

	newStdLine.numOfStdLines = 1

	newStdLine.newLineChars = []rune{'\n'}

	newStdLine.turnLineTerminatorOff = false

	return &newStdLine
}

// NewWithFieldArray - Returns a pointer to a new instance of
// TextLineSpecStandardLine. The returned new instance is
// constructed from input parameters 'numOfStdLines' and
// 'textFields'.
//
// Input parameter 'textFields' is an array of
// ITextFieldSpecification objects which will be formatted on a
// single line of text.
//
//
// ------------------------------------------------------------------------
//
// Default Values
//
// This method will automatically set the following default values:
//
//  newLineChars   - Defaults new line character to '\n'.
//                   To change the value of the new line character,
//                   or characters, see methods:
//                   TextLineSpecStandardLine.SetLineTerminationChars()
//                   TextLineSpecStandardLine.SetLineTerminationRunes()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numOfStdLines              int
//     - An integer value specifying the number of repetitions for
//       a standard line text formatted for text display, file
//       output or printing.
//
//       A 'numOfStdLines' value of 1 means the line will be output
//       once, a value of 2 signals the line will be repeated or
//       output twice, a value of '3' signals the line will be output
//       3-times and so on.
//
//       If the 'numOfStdLines' value is set to zero, no text line
//       will be formatted for text display, file output or printing.
//
//       If this value is set to a value less than zero, it will be
//       automatically reset to a value of one ('1').
//
//       The following examples illustrate the use of
//       'numOfStdLines':
//         Example #1:
//          Standard Line Text = "Hello World"
//          numOfStdLines = 1
//          Text Output:
//            "Hello World"
//
//         Example #2:
//          Standard Line Text = "Hello World"
//          numOfStdLines = 3
//          Text Output:
//            "Hello World"
//            "Hello World"
//            "Hello World"
//
//
//  textFields                 []ITextFieldSpecification
//     - 'textFields' is a collection of objects implementing the
//       ITextLineSpecification interface. These text fields are
//       assembled by the TextLineSpecStandardLine type and formatted
//       as a single line of text. This single line of text is
//       is output one or more times as specified by input
//       parameter, 'numOfStdLines'.
//
//       If this parameter is submitted as a 'nil' value or a zero
//       length array, an error will be returned.
//
//       If any of the objects contained in this collection are
//       invalid, an error will be returned.
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
//  *TextLineSpecStandardLine
//     - If this method completes successfully, it will create and
//       return a pointer to a new instance of
//       TextLineSpecStandardLine which is fully configured with all
//       the parameters necessary to format one or more standard lines
//       of text for text display, file output or printing.
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
func (stdLine TextLineSpecStandardLine) NewWithFieldArray(
	numOfStdLines int,
	textFields []ITextFieldSpecification,
	errorPrefix interface{}) (
	*TextLineSpecStandardLine,
	error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newStdLine := TextLineSpecStandardLine{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.NewWithFieldArray()",
		"")

	if err != nil {
		return &newStdLine, err
	}

	newStdLine.lock = new(sync.Mutex)

	if textFields == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is invalid!\n"+
			"'textFields' is a 'nil' value.\n",
			ePrefix.String())

		return &newStdLine, err
	}

	lenTxtFields := len(textFields)

	if lenTxtFields == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textFields' is invalid!\n"+
			"'textFields' is a zero length array.\n",
			ePrefix.String())

		return &newStdLine, err
	}

	newStdLine.textFields = make([]ITextFieldSpecification,
		lenTxtFields)

	for i := 0; i < lenTxtFields; i++ {

		err = textFields[i].IsValidInstanceError(
			ePrefix.XCtx(
				fmt.Sprintf(
					"textFields[%v] Error",
					i)))

		if err != nil {
			return &newStdLine, err
		}

	}

	var newTxtField ITextFieldSpecification

	for j := 0; j < lenTxtFields; j++ {

		newTxtField,
			err = textFields[j].
			CopyOutITextField(
				ePrefix.XCtx(
					fmt.Sprintf(
						"textFields[%v] Error", j)))

		newStdLine.textFields[j] = newTxtField
	}

	if numOfStdLines < 0 {
		numOfStdLines = 1
	}

	newStdLine.numOfStdLines = numOfStdLines

	newStdLine.newLineChars = []rune{'\n'}

	newStdLine.turnLineTerminatorOff = false

	return &newStdLine, err
}

// ReplaceTextField - Receives an object which implements the
// ITextFieldSpecification interface. A deep copy of this object
// will replace an existing text field object within the text
// fields collection maintained by this TextLineSpecStandardLine
// instance.
//
// The text field object to be replaced must exist at the index
// specified by input parameter, 'replaceAtIndex'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textField                  ITextFieldSpecification
//     - A text line object which implements the
//       ITextFieldSpecification interface. A deep copy of this
//       object will replace an existing element within the
//       text fields collection maintained by this instance of
//       TextLineSpecStandardLine. The text line object to
//       be replaced is identified by the collection element index
//       supplied by input parameter 'replaceAtIndex'.
//
//       If member variable data values contained in this
//       'textField' parameter are found to be invalid, an error
//       will be returned.
//
//
//  replaceAtIndex             int
//     - The index of an element within the text fields collection
//       maintained by the current TextLineSpecStandardLine
//       instance which will be replaced by input parameter
//       'textField'.
//
//       Remember that the text fields collection maintained by
//       the current TextLineSpecStandardLine instance is a zero
//       based array. Therefore, the first index in the collection
//       is zero (0).
//
//       If 'replaceAtIndex' proves to be an invalid index, an error
//       will be returned.
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
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (stdLine *TextLineSpecStandardLine) ReplaceTextField(
	textField ITextFieldSpecification,
	replaceAtIndex int,
	errorPrefix interface{}) (
	err error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine.ReplaceTextField()",
		"")

	if err != nil {
		return err
	}

	lenOfTextFieldsCol := len(stdLine.textFields)

	if lenOfTextFieldsCol == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: The text fields collection is empty and contains no text fields!\n"+
			"First add some text fields before trying to replace a text field.\n",
			ePrefix.String())

		return err
	}

	err = textField.IsValidInstanceError(
		ePrefix.XCtx("Input Parameter: 'textField' is invalid"))

	if err != nil {
		return err
	}

	if replaceAtIndex < 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is invalid!\n"+
			"Index is out of range. 'replaceAtIndex' is less than zero (0).\n"+
			"replaceAtIndex = '%v'\n",
			ePrefix.String(),
			replaceAtIndex)
	}

	lenOfTextFieldsCol--

	if replaceAtIndex > lenOfTextFieldsCol {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'replaceAtIndex' is out of range and invalid!\n"+
			"'replaceAtIndex' is greater than the maximum collection index.\n"+
			"The last element in the text fields collection is index '%v'.\n"+
			"Input parameter 'replaceAtIndex' = '%v'\n",
			ePrefix.String(),
			lenOfTextFieldsCol,
			replaceAtIndex)

		return err
	}

	var newTextField ITextFieldSpecification

	newTextField,
		err = textField.CopyOutITextField(
		ePrefix.XCtx(
			"newTextField"))

	if err != nil {
		return err
	}

	if stdLine.textFields[replaceAtIndex] != nil {

		stdLine.textFields[replaceAtIndex].Empty()
		stdLine.textFields[replaceAtIndex] = nil

	}

	stdLine.textFields[replaceAtIndex] = newTextField

	return err
}

// SetNumOfStdLines - Sets the number of repetitions for this
// instance of TextLineSpecStandardLine. The number of standard
// lines is the number of times this standard line will be output
// printed.
//
// If the input parameter 'numOfStdLines' is less than zero, this
// method will automatically reset the 'numOfStdLines' value to
// one ('1').
//
// If input parameter 'numOfStdLines' is set to a zero value, it
// means that no text will be output or printed for this instance
// of TextLineSpecStandardLine.
//
//       The following examples illustrate the use of
//       'numOfStdLines':
//         Example #1:
//          Standard Line Text = "Hello World"
//          numOfStdLines = 1
//          Text Output:
//            "Hello World"
//
//         Example #2:
//          Standard Line Text = "Hello World"
//          numOfStdLines = 3
//          Text Output:
//            "Hello World"
//            "Hello World"
//            "Hello World"
//
//
func (stdLine *TextLineSpecStandardLine) SetNumOfStdLines(
	numOfStdLines int) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	if numOfStdLines < 0 {
		numOfStdLines = 1
	}

	stdLine.numOfStdLines = numOfStdLines
}

// SetLineTerminationChars - By default, the line termination
// character for standard lines produced and applied by this text
// line specification is the new line character, '\n'. However,
// users have the option of substituting a character or series
// of characters for the line termination sequence using this
// method.
//
// This method will receive a string as input and apply the
// characters in that string as the line termination sequence for
// this instance of TextLineSpecStandardLine.
//
// If input parameter 'lineTerminationChars' is submitted as an
// empty string, this method will take no action and exit.
//
// By default, the TextLineSpecStandardLine type applies the line
// termination characters to each standard line of text generated.
// However, users have the option of turning off the automatic
// generation of text line terminators by calling method:
//
//   TextLineSpecStandardLine.TurnAutoLineTerminationOff()
//
func (stdLine *TextLineSpecStandardLine) SetLineTerminationChars(
	lineTerminationChars string) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	if len(lineTerminationChars) == 0 {
		return
	}

	stdLine.newLineChars = []rune(lineTerminationChars)
}

// SetLineTerminationRunes - By default, the line termination
// character for standard lines produced and applied by this text
// line specification is the new line character, '\n'. However,
// users have the option of substituting a character, or series
// of characters, for the line termination sequence using this
// method.
//
// This method will receive an array of runes as input and applies
// the characters in this array as the line termination sequence
// for this instance of TextLineSpecStandardLine.
//
// If input parameter 'lineTerminationChars' is submitted as a zero
// length array, or if any of the rune values in this array have a
// zero value, this method will return an error.
//
// By default, the TextLineSpecStandardLine type applies the line
// termination characters to each standard line of text generated.
// However, users have the option of turning off the automatic
// generation of text line terminators by calling method:
//
//   TextLineSpecStandardLine.TurnAutoLineTerminationOff()
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  lineTerminationChars       []rune
//     - An array of runes containing the character or characters
//       used to terminate each standard text line generated by
//       the current instance of TextLineSpecStandardLine.
//
//       If this parameter is submitted as a zero length or empty
//       rune array, or if the rune array contains invalid zero
//       rune values, an error will be returned.
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
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (stdLine *TextLineSpecStandardLine) SetLineTerminationRunes(
	lineTerminationChars []rune,
	errorPrefix interface{}) (
	err error) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextLineSpecStandardLine."+
			"SetLineTerminationRunes()",
		"")

	if err != nil {
		return err
	}

	sMechPreon := strMechPreon{}

	_,
		err =
		sMechPreon.testValidityOfRuneCharArray(
			lineTerminationChars,
			ePrefix.XCtx(
				"lineTerminationChars Error"))

	if err != nil {
		return err
	}

	err =
		sMechPreon.copyRuneArrays(
			&stdLine.newLineChars,
			&lineTerminationChars,
			true,
			ePrefix.XCtx(
				"lineTerminationChars->"+
					"stdLine.newLineChars"))

	return err
}

// SetTextFields - Replaces the existing array of text fields for
// the current TextLineSpecStandardLine instance.
//
// If input parameter 'textFields' is nil or zero length, the internal
// array of text fields will be emptied and set to nil.
//
func (stdLine *TextLineSpecStandardLine) SetTextFields(
	textFields []ITextFieldSpecification) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	lenTextFields := len(textFields)

	if lenTextFields == 0 {
		stdLine.textFields = nil
		return
	}

	stdLine.textFields = make([]ITextFieldSpecification,
		lenTextFields)

	copy(stdLine.textFields, textFields)

	return
}

// String - Returns the formatted text generated by this Text Line
// Specification for output display and printing.
//
// The standard line may be replicated multiple times if the
// value of internal member variable' stdLine.numOfStdLines' is
// greater than one ('1').
//
// If an error occurs, the error message will be included in the
// returned string.
//
// This method is identical in function to
// TextLineSpecStandardLine.GetFormattedText()
//
// This method fulfills requirements of the ITextLineSpecification
// interface.
//
func (stdLine TextLineSpecStandardLine) String() string {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TextLineSpecStandardLine.GetFormattedText()",
		"")

	var formattedText string
	var err error

	formattedText,
		err = textLineSpecStandardLineMolecule{}.ptr().
		getFormattedText(
			&stdLine,
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
func (stdLine TextLineSpecStandardLine) TextTypeName() string {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	return "TextLineSpecStandardLine"
}

// TextLineSpecName - returns a string specifying the name
// of this Text Line Specification.
//
// This method fulfills requirements of ITextLineSpecification
// interface.
//
func (stdLine TextLineSpecStandardLine) TextLineSpecName() string {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	return "TextLineSpecStandardLine"
}

// TurnAutoLineTerminationOff - Turns off or cancels the automatic
// generation of new line terminators for each line of text
// produced by this instance of TextLineSpecStandardLine.
//
// By default, the TextLineSpecStandardLine type will generate a
// line of text with a "new line" ('\n') line termination
// character. The application of this "new line" ('\n') character
// is automatic and by default.
//
// However, users may turn off the automatic application of line
// termination characters by calling this method.
//
func (stdLine *TextLineSpecStandardLine) TurnAutoLineTerminationOff() {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	stdLine.turnLineTerminatorOff = true
}

// TurnAutoLineTerminationOn - Turns on or engages the automatic
// generation of new line terminators for each line of text
// produced by this instance of TextLineSpecStandardLine.
//
// By default, the TextLineSpecStandardLine type will generate a
// line of text with a "new line" ('\n') line termination
// character. The application of this "new line" ('\n') character
// is automatic and by default.
//
// However, users may control and override this default behavior
// by calling the method:
//     TextLineSpecStandardLine.TurnAutoLineTerminationOff()
//
// This method will turn on or re-engage the default behavior which
// will automatically apply line termination characters at the end
// of each line of text generated by this instance of
// TextLineSpecStandardLine.
//
func (stdLine *TextLineSpecStandardLine) TurnAutoLineTerminationOn() {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	stdLine.turnLineTerminatorOff = false
}

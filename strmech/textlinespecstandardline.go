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
// ----------------------------------------------------------------
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
// ----------------------------------------------------------------
//
// The TextLineSpecStandardLine type implements the
// ITextLineSpecification interface.
//
type TextLineSpecStandardLine struct {
	textFields            []ITextFieldSpecification
	numOfStdLines         int
	turnLineTerminatorOff bool
	newLineChars          []rune
	lock                  *sync.Mutex
}

// AddTextField - This method will append a text field object to
// the end of the current array of text field objects.
//
func (stdLine *TextLineSpecStandardLine) AddTextField(
	textField ITextFieldSpecification) {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	stdLine.textFields = append(stdLine.textFields,
		textField)
}

// CopyIn - Copies the data fields from an incoming instance of
// TextLineSpecStandardLine ('incomingStdLine') to the data fields
// of the current TextLineSpecStandardLine instance ('stdLine').
//
// IMPORTANT
// All of the data fields in current TextLineSpecStandardLine
// instance ('stdLine') will be modified and overwritten.
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
func (stdLine *TextLineSpecStandardLine) CopyOut() TextLineSpecStandardLine {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	newStdLine,
		_ := textLineSpecStandardLineMolecule{}.ptr().
		copyOut(
			stdLine,
			nil)

	return newStdLine
}

// CopyOutITextLine - Returns a deep copy of the current
// TextLineSpecStandardLine instance cast as an
// ITextLineSpecification object.
//
func (stdLine *TextLineSpecStandardLine) CopyOutITextLine() ITextLineSpecification {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	newStdLine,
		_ := textLineSpecStandardLineMolecule{}.ptr().
		copyOut(
			stdLine,
			nil)

	return ITextLineSpecification(&newStdLine)
}

// CopyOutPtr - Returns a pointer to a deep copy of the current
// TextLineSpecStandardLine instance.
//
func (stdLine *TextLineSpecStandardLine) CopyOutPtr() *TextLineSpecStandardLine {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	newStdLine,
		_ := textLineSpecStandardLineMolecule{}.ptr().
		copyOut(
			stdLine,
			nil)

	return &newStdLine
}

// Empty - Deletes all of the text fields stored as an array of
// ITextFieldSpecification pointers within the current
// TextLineSpecStandardLine instance.
//
func (stdLine *TextLineSpecStandardLine) Empty() {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	stdLine.numOfStdLines = 0

	stdLine.turnLineTerminatorOff = false

	stdLine.newLineChars = nil

	for i := 0; i < len(stdLine.textFields); i++ {
		stdLine.textFields[i].Empty()
		stdLine.textFields[i] = nil
	}

	stdLine.textFields = nil

	stdLine.lock.Unlock()

	stdLine.lock = nil

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
// Text Line Specification for output and printing.
//
// The standard line may be replicated multiple times if the
// value of internal member variable' stdLine.numOfStdLines' is
// greater than one ('1').
//
func (stdLine *TextLineSpecStandardLine) GetFormattedText() string {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	if stdLine.numOfStdLines < 0 {
		stdLine.numOfStdLines = 0
	}

	lenTextFields := len(stdLine.textFields)

	if stdLine.numOfStdLines == 0 ||
		lenTextFields == 0 {
		return ""
	}

	var result, lineStr string

	for i := 0; i < lenTextFields; i++ {
		lineStr += stdLine.textFields[i].GetFormattedText()
	}

	for j := 0; j < stdLine.numOfStdLines; j++ {
		result += lineStr
	}

	if stdLine.turnLineTerminatorOff == true {
		return result
	}

	if len(stdLine.newLineChars) == 0 {
		stdLine.newLineChars = []rune{'\n'}
	}

	result += string(stdLine.newLineChars)

	return result
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

// GetTextFields - Returns a deep copy of the text fields contained
// in the current TextLineSpecStandardLine instance.
//
func (stdLine *TextLineSpecStandardLine) GetTextFields() []ITextFieldSpecification {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	lenTxtFields := len(stdLine.textFields)

	if lenTxtFields == 0 {
		return nil
	}

	textFields := make([]ITextFieldSpecification, lenTxtFields)

	copy(textFields, stdLine.textFields)

	return textFields
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
//       the current instance of TextFieldSpecLabel are found to be
//       invalid, this method will return an error.
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

	if stdLine.numOfStdLines < 1 {
		err = fmt.Errorf("%v\n"+
			"Error: The number of standard lines is less than one ('1')!\n"+
			"This means that no lines will be generate by this specification.\n",
			ePrefix.String())
		return err
	}

	lenTextFields := len(stdLine.textFields)

	if lenTextFields == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: No Text Fields have been configured for\n"+
			"this standard line specification!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	for i := 0; i < lenTextFields; i++ {

		err2 = stdLine.textFields[i].IsValidInstanceError(ePrefix)

		if err2 != nil {
			err = fmt.Errorf("%v\n"+
				"Error: Text Field Element [%v] is invalid!\n"+
				"Text Field Element Error = \n%v\n",
				ePrefix.String(),
				i,
				err2.Error())

			return err
		}

	}

	return nil
}

// New - Returns a new and empty, or unpopulated, concrete instance
// of TextLineSpecStandardLine.
//
// The default number of lines repetitions is set to one ('1').
// This means that the returned standard line instance will only
// be output or printed once.
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
// 'numOfStdLines' specifies the number of repetitions for
// standard line output or printing. If this value is set to zero,
// no text line will be output or printed. If 'numOfStdLines' is
// set to value less than zero, it will be automatically reset to
// one ('1').
//
// A 'numOfStdLines' value of 1 means the line will be output once,
// a value of 2 signals the line will be repeated or output twice,
// a value of '3' signals the line will be output 3-times and so
// on.
//
// Input parameter 'textFields' is an array of
// ITextFieldSpecification objects which will be formatted on a
// single line of text.
//
func (stdLine TextLineSpecStandardLine) NewWithFieldArray(
	numOfStdLines int,
	textFields []ITextFieldSpecification) *TextLineSpecStandardLine {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	newStdLine := TextLineSpecStandardLine{}

	newStdLine.lock = new(sync.Mutex)

	lenTxtFields := len(textFields)

	if lenTxtFields > 0 {

		newStdLine.textFields = make([]ITextFieldSpecification,
			lenTxtFields)

		copy(newStdLine.textFields,
			textFields)
	}

	if numOfStdLines < 0 {
		numOfStdLines = 1
	}

	newStdLine.numOfStdLines = numOfStdLines

	newStdLine.newLineChars = []rune{'\n'}

	newStdLine.turnLineTerminatorOff = false

	return &newStdLine
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
// character for blank line produced and applied by this text line
// specification is the the new line character, '\n'. However,
// users have the option of substituting a character or series
// of characters for the the line termination sequence using this
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

// SetTextFields - Replaces the existing array of text fields for
// the current TextLineSpecStandardLine instance.
//
// If input parameter 'textFields' is nil or zero length, the internal
// array of text fields will be emptied an set to nil.
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

// TextTypeName - returns a string specifying the type
// of Text Line Specification. This method fulfills
// requirements of ITextSpecification interface.
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
// of this Text Line Specification. This method fulfills
// requirements of ITextLineSpecification interface.
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

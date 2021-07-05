package strmech

import "sync"

// TextLineSpecStandardLine - This type is a text specification for
// a standard line of text. It encapsulates an array of
// ITextFieldSpecification objects which are used to format text
// fields within a single line of text.
//
type TextLineSpecStandardLine struct {
	textFields    []ITextFieldSpecification
	numOfStdLines int
	lock          *sync.Mutex
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

// CopyOut - Returns a deep copy of the current
// TextLineSpecStandardLine instance.
//
func (stdLine *TextLineSpecStandardLine) CopyOut() *TextLineSpecStandardLine {

	if stdLine.lock == nil {
		stdLine.lock = new(sync.Mutex)
	}

	stdLine.lock.Lock()

	defer stdLine.lock.Unlock()

	newStdLine := new(TextLineSpecStandardLine)

	lenTxtFields := len(stdLine.textFields)

	if lenTxtFields > 0 {

		newStdLine.textFields = make([]ITextFieldSpecification,
			lenTxtFields)

		copy(newStdLine.textFields,
			stdLine.textFields)
	}

	newStdLine.numOfStdLines = stdLine.numOfStdLines

	return newStdLine
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

	newStdLine.numOfStdLines = 1

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
// of Text Field specification. This method fulfills
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

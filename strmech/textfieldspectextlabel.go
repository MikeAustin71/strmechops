package strmech

import "sync"

type TextFieldSpecLabel struct {
	textLabel         string
	fieldLen          int
	textJustification TextJustify
	lock              *sync.Mutex
}

// CopyOut - Returns a deep copy of the current
// TextFieldSpecLabel instance.
//
func (txtFieldLabel *TextFieldSpecLabel) CopyOut() TextFieldSpecLabel {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	newTxtFieldLabel := TextFieldSpecLabel{}

	newTxtFieldLabel.textLabel =
		txtFieldLabel.textLabel

	newTxtFieldLabel.fieldLen =
		txtFieldLabel.fieldLen

	newTxtFieldLabel.textJustification =
		txtFieldLabel.textJustification

	return newTxtFieldLabel
}

// NewPtr - Returns a pointer to a new instance of
// TextFieldSpecLabel.
//
// The new returned instance of TextFieldSpecLabel is constructed
// from input parameters, 'textLabel', 'fieldLen' and
// 'textJustification'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  textLabel                  string
//     - The string content to be displayed within the label. If
//       this parameter is submitted as a zero length string it
//       will be automatically converted to a string consisting of
//       white space (space characters) with a length equal to that
//       of input parameter 'fieldLen'.
//
//  fieldLen                   int
//     - The length of the field in which the 'textLabel' will be
//       displayed. If 'fieldLen' is less than the length of the
//       'textLabel' string, it will be automatically set equal to
//       the 'textLabel' string length.
//
//
//  textJustification          TextJustify
//     - An enumeration which specifies the justification of the
//       'textLabel' within the field specified by 'fieldLen'.
//       Options for 'textJustification' are:
//           TextJustify(0).Left()
//           TextJustify(0).Right()
//           TextJustify(0).Center()
//
//       If input parameter 'textJustification' is invalid, it will
//       be automatically set to TextJustify(0).Left().
//
//       If 'fieldLen' is less than or equal to the 'textLabel'
//       string length, text justification will not apply.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  *TextFieldSpecLabel
//     - This method will return a pointer to a new instance of
//       TextFieldSpecLabel constructed from the information
//       provided by the input parameters.
//
func (txtFieldLabel TextFieldSpecLabel) NewPtr(
	textLabel string,
	fieldLen int,
	textJustification TextJustify) *TextFieldSpecLabel {

	if txtFieldLabel.lock == nil {
		txtFieldLabel.lock = new(sync.Mutex)
	}

	txtFieldLabel.lock.Lock()

	defer txtFieldLabel.lock.Unlock()

	newTextLabel := TextFieldSpecLabel{}

	newTextLabel.textLabel = textLabel

	if fieldLen < len(textLabel) {
		fieldLen = len(textLabel)
	}

	newTextLabel.fieldLen = fieldLen

	if !textJustification.XIsValid() {
		textJustification = TextJustify(0).Left()
	}

	newTextLabel.textJustification = textJustification

	return &newTextLabel
}

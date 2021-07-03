package strmech

import "sync"

type TextFieldString struct {
	LeftMargin           TextLineSpecMargin
	LeftSpacer           TextLineSpecMargin
	StrTxtSpec           TextLineSpecString
	RightSpacer          TextLineSpecMargin
	TerminateWithNewLine TextLineSpecNewLine
	lock                 *sync.Mutex
}

func (txtFieldStr *TextFieldString) CopyOut() TextFieldString {

	if txtFieldStr.lock == nil {
		txtFieldStr.lock = new(sync.Mutex)
	}

	txtFieldStr.lock.Lock()

	defer txtFieldStr.lock.Unlock()

	newStrField := TextFieldString{}

	newStrField.LeftMargin = txtFieldStr.LeftMargin.CopyOut()
	newStrField.LeftSpacer = txtFieldStr.LeftSpacer.CopyOut()
	newStrField.StrTxtSpec = txtFieldStr.StrTxtSpec.CopyOut()
	newStrField.RightSpacer = txtFieldStr.RightSpacer.CopyOut()
	newStrField.TerminateWithNewLine = txtFieldStr.TerminateWithNewLine

	return newStrField
}

// NewPtr - Returns a pointer to a new TextFieldString instance.
func (txtFieldStr TextFieldString) NewPtr() *TextFieldString {

	if txtFieldStr.lock == nil {
		txtFieldStr.lock = new(sync.Mutex)
	}

	txtFieldStr.lock.Lock()

	defer txtFieldStr.lock.Unlock()

	newStrFieldSpec := TextFieldString{}

	return &newStrFieldSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtFieldStr TextFieldString) TextTypeName() string {

	if txtFieldStr.lock == nil {
		txtFieldStr.lock = new(sync.Mutex)
	}

	txtFieldStr.lock.Lock()

	defer txtFieldStr.lock.Unlock()

	return "TextFieldString"
}

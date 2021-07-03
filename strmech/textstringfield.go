package strmech

import "sync"

type TextStringField struct {
	LeftMargin           MarginSpec
	LeftSpacer           MarginSpec
	StrTxtSpec           StringSpec
	RightSpacer          MarginSpec
	TerminateWithNewLine NewLineSpec
	lock                 *sync.Mutex
}

func (txtStrField *TextStringField) CopyOut() TextStringField {

	if txtStrField.lock == nil {
		txtStrField.lock = new(sync.Mutex)
	}

	txtStrField.lock.Lock()

	defer txtStrField.lock.Unlock()

	newStrField := TextStringField{}

	newStrField.LeftMargin = txtStrField.LeftMargin.CopyOut()
	newStrField.LeftSpacer = txtStrField.LeftSpacer.CopyOut()
	newStrField.StrTxtSpec = txtStrField.StrTxtSpec.CopyOut()
	newStrField.RightSpacer = txtStrField.RightSpacer.CopyOut()
	newStrField.TerminateWithNewLine = txtStrField.TerminateWithNewLine

	return newStrField
}

// NewPtr - Returns a pointer to a new TextStringField instance.
func (txtStrField TextStringField) NewPtr() *TextStringField {

	if txtStrField.lock == nil {
		txtStrField.lock = new(sync.Mutex)
	}

	txtStrField.lock.Lock()

	defer txtStrField.lock.Unlock()

	newStrFieldSpec := TextStringField{}

	return &newStrFieldSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtStrField TextStringField) TextTypeName() string {

	if txtStrField.lock == nil {
		txtStrField.lock = new(sync.Mutex)
	}

	txtStrField.lock.Lock()

	defer txtStrField.lock.Unlock()

	return "TextStringField"
}

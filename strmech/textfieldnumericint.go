package strmech

import "sync"

type TextFieldNumericInt struct {
	LeftMargin           TextLineSpecMargin
	LeftSpacer           TextLineSpecMargin
	NumberSpec           TextLineSpecInteger
	RightSpacer          TextLineSpecMargin
	TerminateWithNewLine TextLineSpecNewLine
	lock                 *sync.Mutex
}

// CopyOut - Returns a deep copy of the current TextFieldNumericInt
// instance.
func (txtFieldNumInt *TextFieldNumericInt) CopyOut() TextFieldNumericInt {

	if txtFieldNumInt.lock == nil {
		txtFieldNumInt.lock = new(sync.Mutex)
	}

	txtFieldNumInt.lock.Lock()

	defer txtFieldNumInt.lock.Unlock()

	newNumField := TextFieldNumericInt{}

	newNumField.LeftMargin = txtFieldNumInt.LeftMargin.CopyOut()
	newNumField.LeftSpacer = txtFieldNumInt.LeftSpacer.CopyOut()
	newNumField.NumberSpec = txtFieldNumInt.NumberSpec.CopyOut()
	newNumField.RightSpacer = txtFieldNumInt.RightSpacer.CopyOut()
	newNumField.TerminateWithNewLine = txtFieldNumInt.TerminateWithNewLine.CopyOut()

	return newNumField
}

// NewPtr - Returns a pointer to a new TextFieldNumericInt instance.
func (txtFieldNumInt TextFieldNumericInt) NewPtr() *TextFieldNumericInt {

	if txtFieldNumInt.lock == nil {
		txtFieldNumInt.lock = new(sync.Mutex)
	}

	txtFieldNumInt.lock.Lock()

	defer txtFieldNumInt.lock.Unlock()

	newNumIntSpec := TextFieldNumericInt{}

	return &newNumIntSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtFieldNumInt TextFieldNumericInt) TextTypeName() string {

	if txtFieldNumInt.lock == nil {
		txtFieldNumInt.lock = new(sync.Mutex)
	}

	txtFieldNumInt.lock.Lock()

	defer txtFieldNumInt.lock.Unlock()

	return "TextFieldNumericInt"
}

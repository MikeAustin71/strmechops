package strmech

import "sync"

type TextNumericIntField struct {
	LeftMargin           TextLineSpecMargin
	LeftSpacer           TextLineSpecMargin
	NumberSpec           TextLineSpecInteger
	RightSpacer          TextLineSpecMargin
	TerminateWithNewLine TextLineSpecNewLine
	lock                 *sync.Mutex
}

// CopyOut - Returns a deep copy of the current TextNumericIntField
// instance.
func (txtNumIntField *TextNumericIntField) CopyOut() TextNumericIntField {

	if txtNumIntField.lock == nil {
		txtNumIntField.lock = new(sync.Mutex)
	}

	txtNumIntField.lock.Lock()

	defer txtNumIntField.lock.Unlock()

	newNumField := TextNumericIntField{}

	newNumField.LeftMargin = txtNumIntField.LeftMargin.CopyOut()
	newNumField.LeftSpacer = txtNumIntField.LeftSpacer.CopyOut()
	newNumField.NumberSpec = txtNumIntField.NumberSpec.CopyOut()
	newNumField.RightSpacer = txtNumIntField.RightSpacer.CopyOut()
	newNumField.TerminateWithNewLine = txtNumIntField.TerminateWithNewLine.CopyOut()

	return newNumField
}

// NewPtr - Returns a pointer to a new TextNumericIntField instance.
func (txtNumIntField TextNumericIntField) NewPtr() *TextNumericIntField {

	if txtNumIntField.lock == nil {
		txtNumIntField.lock = new(sync.Mutex)
	}

	txtNumIntField.lock.Lock()

	defer txtNumIntField.lock.Unlock()

	newNumIntSpec := TextNumericIntField{}

	return &newNumIntSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtNumIntField TextNumericIntField) TextTypeName() string {

	if txtNumIntField.lock == nil {
		txtNumIntField.lock = new(sync.Mutex)
	}

	txtNumIntField.lock.Lock()

	defer txtNumIntField.lock.Unlock()

	return "TextNumericIntField"
}

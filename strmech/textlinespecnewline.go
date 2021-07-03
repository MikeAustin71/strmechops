package strmech

import "sync"

type TextLineSpecNewLine struct {
	AddNewLine bool
	lock       *sync.Mutex
}

func (txtLineSpecNewLine *TextLineSpecNewLine) CopyOut() TextLineSpecNewLine {

	if txtLineSpecNewLine.lock == nil {
		txtLineSpecNewLine.lock = new(sync.Mutex)
	}

	txtLineSpecNewLine.lock.Lock()

	defer txtLineSpecNewLine.lock.Unlock()

	newNewLine := TextLineSpecNewLine{}

	newNewLine.AddNewLine = txtLineSpecNewLine.AddNewLine

	return newNewLine
}

// NewPtr - Returns a pointer to a new TextLineSpecNewLine
// instance.
func (txtLineSpecNewLine TextLineSpecNewLine) NewPtr() *TextLineSpecNewLine {

	if txtLineSpecNewLine.lock == nil {
		txtLineSpecNewLine.lock = new(sync.Mutex)
	}

	txtLineSpecNewLine.lock.Lock()

	defer txtLineSpecNewLine.lock.Unlock()

	newLineSpec := TextLineSpecNewLine{}

	return &newLineSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtLineSpecNewLine TextLineSpecNewLine) TextTypeName() string {

	if txtLineSpecNewLine.lock == nil {
		txtLineSpecNewLine.lock = new(sync.Mutex)
	}

	txtLineSpecNewLine.lock.Lock()

	defer txtLineSpecNewLine.lock.Unlock()

	return "TextLineSpecNewLine"
}

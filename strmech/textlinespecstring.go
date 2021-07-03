package strmech

import "sync"

type TextLineSpecString struct {
	StrValue       string
	StrFieldLength int
	StrPadChar     rune
	StrPosition    FieldPositionSpec
	lock           *sync.Mutex
}

// CopyOut - Returns a deep copy of the current TextLineSpecString
// instance.
func (txtLineSpecString *TextLineSpecString) CopyOut() TextLineSpecString {

	if txtLineSpecString.lock == nil {
		txtLineSpecString.lock = new(sync.Mutex)
	}

	txtLineSpecString.lock.Lock()

	defer txtLineSpecString.lock.Unlock()

	newStrSpec := TextLineSpecString{}
	newStrSpec.StrValue = txtLineSpecString.StrValue
	newStrSpec.StrFieldLength = txtLineSpecString.StrFieldLength
	newStrSpec.StrPadChar = txtLineSpecString.StrPadChar

	return newStrSpec
}

// NewPtr - Returns a pointer to a new TextLineSpecString instance.
func (txtLineSpecString TextLineSpecString) NewPtr() *TextLineSpecString {

	if txtLineSpecString.lock == nil {
		txtLineSpecString.lock = new(sync.Mutex)
	}

	txtLineSpecString.lock.Lock()

	defer txtLineSpecString.lock.Unlock()

	newStrTxtSpec := TextLineSpecString{}

	return &newStrTxtSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtLineSpecString TextLineSpecString) TextTypeName() string {

	if txtLineSpecString.lock == nil {
		txtLineSpecString.lock = new(sync.Mutex)
	}

	txtLineSpecString.lock.Lock()

	defer txtLineSpecString.lock.Unlock()

	return "TextLineSpecString"
}

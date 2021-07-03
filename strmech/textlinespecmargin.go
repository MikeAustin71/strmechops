package strmech

import "sync"

type TextLineSpecMargin struct {
	MarginStr    string
	MarginLength int
	MarginChar   rune
	lock         *sync.Mutex
}

func (txtLineMargin *TextLineSpecMargin) CopyOut() TextLineSpecMargin {

	if txtLineMargin.lock == nil {
		txtLineMargin.lock = new(sync.Mutex)
	}

	txtLineMargin.lock.Lock()

	defer txtLineMargin.lock.Unlock()

	newMargin := TextLineSpecMargin{}

	newMargin.MarginStr = txtLineMargin.MarginStr
	newMargin.MarginChar = txtLineMargin.MarginChar
	newMargin.MarginLength = txtLineMargin.MarginLength

	return newMargin
}

// NewPtr - Returns a pointer to a new TextLineSpecMargin
// instance.
func (txtLineMargin TextLineSpecMargin) NewPtr() *TextLineSpecMargin {

	if txtLineMargin.lock == nil {
		txtLineMargin.lock = new(sync.Mutex)
	}

	txtLineMargin.lock.Lock()

	defer txtLineMargin.lock.Unlock()

	newMarginField := TextLineSpecMargin{}

	return &newMarginField
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtLineMargin TextLineSpecMargin) TextTypeName() string {

	if txtLineMargin.lock == nil {
		txtLineMargin.lock = new(sync.Mutex)
	}

	txtLineMargin.lock.Lock()

	defer txtLineMargin.lock.Unlock()

	return "TextLineSpecMargin"
}

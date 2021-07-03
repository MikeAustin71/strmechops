package strmech

import "sync"

type TextLineSpecBlankLines struct {
	NumBlankLines int
	lock          *sync.Mutex
}

func (lineSpecBlkLine *TextLineSpecBlankLines) CopyOut() TextLineSpecBlankLines {

	if lineSpecBlkLine.lock == nil {
		lineSpecBlkLine.lock = new(sync.Mutex)
	}

	lineSpecBlkLine.lock.Lock()

	defer lineSpecBlkLine.lock.Unlock()

	newBlankLineSpec := TextLineSpecBlankLines{}

	newBlankLineSpec.NumBlankLines = lineSpecBlkLine.NumBlankLines

	return newBlankLineSpec
}

// NewPtr - Returns a pointer to a new instance of 'TextLineSpecBlankLines'
func (lineSpecBlkLine TextLineSpecBlankLines) NewPtr() *TextLineSpecBlankLines {

	if lineSpecBlkLine.lock == nil {
		lineSpecBlkLine.lock = new(sync.Mutex)
	}

	lineSpecBlkLine.lock.Lock()

	defer lineSpecBlkLine.lock.Unlock()

	newBlkLine := TextLineSpecBlankLines{}

	return &newBlkLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (lineSpecBlkLine TextLineSpecBlankLines) TextTypeName() string {

	if lineSpecBlkLine.lock == nil {
		lineSpecBlkLine.lock = new(sync.Mutex)
	}

	lineSpecBlkLine.lock.Lock()

	defer lineSpecBlkLine.lock.Unlock()

	return "TextLineSpecBlankLines"
}

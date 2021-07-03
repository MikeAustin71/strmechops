package strmech

import "sync"

type TextLineSpecRepeatChar struct {
	LineChar         rune
	LineLength       int
	LineFieldLength  int
	LineFieldPadChar rune
	LinePosition     FieldPositionSpec
	lock             *sync.Mutex
}

// CopyOut - Returns a deep copy of TextLineSpecRepeatChar
// data fields.
func (lineSpecRepeatChar *TextLineSpecRepeatChar) CopyOut() TextLineSpecRepeatChar {

	if lineSpecRepeatChar.lock == nil {
		lineSpecRepeatChar.lock = new(sync.Mutex)
	}

	lineSpecRepeatChar.lock.Lock()

	defer lineSpecRepeatChar.lock.Unlock()

	newLineSpec := TextLineSpecRepeatChar{
		LineChar:         lineSpecRepeatChar.LineChar,
		LineLength:       lineSpecRepeatChar.LineLength,
		LineFieldLength:  lineSpecRepeatChar.LineFieldLength,
		LineFieldPadChar: lineSpecRepeatChar.LineFieldPadChar,
		LinePosition:     lineSpecRepeatChar.LinePosition,
	}

	return newLineSpec
}

// NewPtr - Returns a pointer to a new TextLineSpecRepeatChar
// instance.
func (lineSpecRepeatChar TextLineSpecRepeatChar) NewPtr() *TextLineSpecRepeatChar {

	if lineSpecRepeatChar.lock == nil {
		lineSpecRepeatChar.lock = new(sync.Mutex)
	}

	lineSpecRepeatChar.lock.Lock()

	defer lineSpecRepeatChar.lock.Unlock()

	anotherLineSpec := TextLineSpecRepeatChar{}

	return &anotherLineSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (lineSpecRepeatChar TextLineSpecRepeatChar) TextTypeName() string {

	if lineSpecRepeatChar.lock == nil {
		lineSpecRepeatChar.lock = new(sync.Mutex)
	}

	lineSpecRepeatChar.lock.Lock()

	defer lineSpecRepeatChar.lock.Unlock()

	return "TextLineSpecRepeatChar"
}

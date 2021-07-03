package strmech

import "sync"

type TextLineSpecInteger struct {
	NumericValue       int
	NumericFieldSpec   string
	NumericFieldLength int
	NumericPadChar     rune
	NumericPosition    FieldPositionSpec
	lock               *sync.Mutex
}

func (lineSpecInt *TextLineSpecInteger) CopyOut() TextLineSpecInteger {

	if lineSpecInt.lock == nil {
		lineSpecInt.lock = new(sync.Mutex)
	}

	lineSpecInt.lock.Lock()

	defer lineSpecInt.lock.Unlock()

	newInt := TextLineSpecInteger{}
	newInt.NumericValue = lineSpecInt.NumericValue
	newInt.NumericFieldSpec = lineSpecInt.NumericFieldSpec
	newInt.NumericFieldLength = lineSpecInt.NumericFieldLength
	newInt.NumericPadChar = lineSpecInt.NumericPadChar

	return newInt
}

// NewPtr - Returns a pointer to a new TextLineSpecRepeatChar
// instance.
func (lineSpecInt TextLineSpecInteger) NewPtr() *TextLineSpecInteger {

	if lineSpecInt.lock == nil {
		lineSpecInt.lock = new(sync.Mutex)
	}

	lineSpecInt.lock.Lock()

	defer lineSpecInt.lock.Unlock()

	anotherIntTxtSpec := TextLineSpecInteger{}

	return &anotherIntTxtSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (lineSpecInt TextLineSpecInteger) TextTypeName() string {

	if lineSpecInt.lock == nil {
		lineSpecInt.lock = new(sync.Mutex)
	}

	lineSpecInt.lock.Lock()

	defer lineSpecInt.lock.Unlock()

	return "TextLineSpecInteger"
}

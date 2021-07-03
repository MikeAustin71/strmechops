package strmech

import "sync"

type TextFieldOneLabelOneInt struct {
	LeadingBlankLines  TextLineSpecBlankLines
	TopLineBreak       TextLineBreakField
	Label1             TextFieldString
	Number1            TextFieldNumericInt
	BottomLineBreak    TextLineBreakField
	TrailingBlankLines TextLineSpecBlankLines
	lock               *sync.Mutex
}

// CopyOut - Returns a deep copy of the current TextFieldOneLabelOneInt
// instance.
func (txtLabelIntLine *TextFieldOneLabelOneInt) CopyOut() TextFieldOneLabelOneInt {

	if txtLabelIntLine.lock == nil {
		txtLabelIntLine.lock = new(sync.Mutex)
	}

	txtLabelIntLine.lock.Lock()

	defer txtLabelIntLine.lock.Unlock()

	newLabelInt := TextFieldOneLabelOneInt{}

	newLabelInt.LeadingBlankLines = txtLabelIntLine.LeadingBlankLines
	newLabelInt.TopLineBreak = txtLabelIntLine.TopLineBreak.CopyOut()
	newLabelInt.Label1 = txtLabelIntLine.Label1.CopyOut()
	newLabelInt.Number1 = txtLabelIntLine.Number1.CopyOut()
	newLabelInt.BottomLineBreak = txtLabelIntLine.BottomLineBreak.CopyOut()
	newLabelInt.TrailingBlankLines = txtLabelIntLine.TrailingBlankLines
	return newLabelInt
}

// NewPtr - Returns a pointer to a new TextFieldOneLabelOneInt instance.
func (txtLabelIntLine TextFieldOneLabelOneInt) NewPtr() *TextFieldOneLabelOneInt {

	if txtLabelIntLine.lock == nil {
		txtLabelIntLine.lock = new(sync.Mutex)
	}

	txtLabelIntLine.lock.Lock()

	defer txtLabelIntLine.lock.Unlock()

	newOneLabelOneIntLine := TextFieldOneLabelOneInt{}

	return &newOneLabelOneIntLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtLabelIntLine TextFieldOneLabelOneInt) TextTypeName() string {

	if txtLabelIntLine.lock == nil {
		txtLabelIntLine.lock = new(sync.Mutex)
	}

	txtLabelIntLine.lock.Lock()

	defer txtLabelIntLine.lock.Unlock()

	return "TextFieldOneLabelOneInt"
}

package strmech

import "sync"

type TextOneLabelOneIntField struct {
	LeadingBlankLines  BlankLinesSpec
	TopLineBreak       TextLineBreakField
	Label1             TextStringField
	Number1            TextNumericIntField
	BottomLineBreak    TextLineBreakField
	TrailingBlankLines BlankLinesSpec
	lock               *sync.Mutex
}

// CopyOut - Returns a deep copy of the current TextOneLabelOneIntField
// instance.
func (txtLabelIntLine *TextOneLabelOneIntField) CopyOut() TextOneLabelOneIntField {

	if txtLabelIntLine.lock == nil {
		txtLabelIntLine.lock = new(sync.Mutex)
	}

	txtLabelIntLine.lock.Lock()

	defer txtLabelIntLine.lock.Unlock()

	newLabelInt := TextOneLabelOneIntField{}

	newLabelInt.LeadingBlankLines = txtLabelIntLine.LeadingBlankLines
	newLabelInt.TopLineBreak = txtLabelIntLine.TopLineBreak.CopyOut()
	newLabelInt.Label1 = txtLabelIntLine.Label1.CopyOut()
	newLabelInt.Number1 = txtLabelIntLine.Number1.CopyOut()
	newLabelInt.BottomLineBreak = txtLabelIntLine.BottomLineBreak.CopyOut()
	newLabelInt.TrailingBlankLines = txtLabelIntLine.TrailingBlankLines
	return newLabelInt
}

// NewPtr - Returns a pointer to a new TextOneLabelOneIntField instance.
func (txtLabelIntLine TextOneLabelOneIntField) NewPtr() *TextOneLabelOneIntField {

	if txtLabelIntLine.lock == nil {
		txtLabelIntLine.lock = new(sync.Mutex)
	}

	txtLabelIntLine.lock.Lock()

	defer txtLabelIntLine.lock.Unlock()

	newOneLabelOneIntLine := TextOneLabelOneIntField{}

	return &newOneLabelOneIntLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtLabelIntLine TextOneLabelOneIntField) TextTypeName() string {

	if txtLabelIntLine.lock == nil {
		txtLabelIntLine.lock = new(sync.Mutex)
	}

	txtLabelIntLine.lock.Lock()

	defer txtLabelIntLine.lock.Unlock()

	return "TextOneLabelOneIntField"
}

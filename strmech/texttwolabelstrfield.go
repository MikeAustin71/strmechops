package strmech

import "sync"

type TextTwoLabelStrField struct {
	LeadingBlankLines  BlankLinesSpec
	TopLineBreak       TextLineBreakField
	Label1             TextStringField
	Label2             TextStringField
	BottomLineBreak    TextLineBreakField
	TrailingBlankLines BlankLinesSpec
	lock               *sync.Mutex
}

// CopyOut - Returns a deep copy of the current TextTwoLabelStrField
// instance.
func (txtTwoLabelLine *TextTwoLabelStrField) CopyOut() TextTwoLabelStrField {

	if txtTwoLabelLine.lock == nil {
		txtTwoLabelLine.lock = new(sync.Mutex)
	}

	txtTwoLabelLine.lock.Lock()

	defer txtTwoLabelLine.lock.Unlock()

	newTwoLabelLine := TextTwoLabelStrField{}

	newTwoLabelLine.LeadingBlankLines = txtTwoLabelLine.LeadingBlankLines.CopyOut()
	newTwoLabelLine.TopLineBreak = txtTwoLabelLine.TopLineBreak.CopyOut()
	newTwoLabelLine.Label1 = txtTwoLabelLine.Label1.CopyOut()
	newTwoLabelLine.Label2 = txtTwoLabelLine.Label2.CopyOut()
	newTwoLabelLine.BottomLineBreak = txtTwoLabelLine.BottomLineBreak.CopyOut()
	newTwoLabelLine.TrailingBlankLines = txtTwoLabelLine.TrailingBlankLines.CopyOut()

	return newTwoLabelLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtTwoLabelLine TextTwoLabelStrField) TextTypeName() string {

	if txtTwoLabelLine.lock == nil {
		txtTwoLabelLine.lock = new(sync.Mutex)
	}

	txtTwoLabelLine.lock.Lock()

	defer txtTwoLabelLine.lock.Unlock()

	return "TextTwoLabelStrField"
}

// NewPtr - Returns a pointer to a new TextTwoLabelStrField instance.
func (txtTwoLabelLine TextTwoLabelStrField) NewPtr() *TextTwoLabelStrField {

	if txtTwoLabelLine.lock == nil {
		txtTwoLabelLine.lock = new(sync.Mutex)
	}

	txtTwoLabelLine.lock.Lock()

	defer txtTwoLabelLine.lock.Unlock()

	newTwoLabelLine := TextTwoLabelStrField{}

	return &newTwoLabelLine
}

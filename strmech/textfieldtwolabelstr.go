package strmech

import "sync"

type TextFieldTwoLabelStr struct {
	LeadingBlankLines  TextLineSpecBlankLines
	TopLineBreak       TextLineBreakField
	Label1             TextFieldString
	Label2             TextFieldString
	BottomLineBreak    TextLineBreakField
	TrailingBlankLines TextLineSpecBlankLines
	lock               *sync.Mutex
}

// CopyOut - Returns a deep copy of the current TextFieldTwoLabelStr
// instance.
func (txtFieldTwoLabelStr *TextFieldTwoLabelStr) CopyOut() TextFieldTwoLabelStr {

	if txtFieldTwoLabelStr.lock == nil {
		txtFieldTwoLabelStr.lock = new(sync.Mutex)
	}

	txtFieldTwoLabelStr.lock.Lock()

	defer txtFieldTwoLabelStr.lock.Unlock()

	newTwoLabelLine := TextFieldTwoLabelStr{}

	newTwoLabelLine.LeadingBlankLines = txtFieldTwoLabelStr.LeadingBlankLines.CopyOut()
	newTwoLabelLine.TopLineBreak = txtFieldTwoLabelStr.TopLineBreak.CopyOut()
	newTwoLabelLine.Label1 = txtFieldTwoLabelStr.Label1.CopyOut()
	newTwoLabelLine.Label2 = txtFieldTwoLabelStr.Label2.CopyOut()
	newTwoLabelLine.BottomLineBreak = txtFieldTwoLabelStr.BottomLineBreak.CopyOut()
	newTwoLabelLine.TrailingBlankLines = txtFieldTwoLabelStr.TrailingBlankLines.CopyOut()

	return newTwoLabelLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtFieldTwoLabelStr TextFieldTwoLabelStr) TextTypeName() string {

	if txtFieldTwoLabelStr.lock == nil {
		txtFieldTwoLabelStr.lock = new(sync.Mutex)
	}

	txtFieldTwoLabelStr.lock.Lock()

	defer txtFieldTwoLabelStr.lock.Unlock()

	return "TextFieldTwoLabelStr"
}

// NewPtr - Returns a pointer to a new TextFieldTwoLabelStr instance.
func (txtFieldTwoLabelStr TextFieldTwoLabelStr) NewPtr() *TextFieldTwoLabelStr {

	if txtFieldTwoLabelStr.lock == nil {
		txtFieldTwoLabelStr.lock = new(sync.Mutex)
	}

	txtFieldTwoLabelStr.lock.Lock()

	defer txtFieldTwoLabelStr.lock.Unlock()

	newTwoLabelLine := TextFieldTwoLabelStr{}

	return &newTwoLabelLine
}

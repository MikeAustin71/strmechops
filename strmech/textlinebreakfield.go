package strmech

import "sync"

type TextLineBreakField struct {
	CreateLineBreak      bool
	LeadingBlankLines    TextLineSpecBlankLines
	LeftMargin           TextLineSpecMargin
	LeftSpacer           TextLineSpecMargin
	LineSpec             TextLineSpecRepeatChar
	RightSpacer          TextLineSpecMargin
	TerminateWithNewLine TextLineSpecNewLine
	TrailingBlankLines   TextLineSpecBlankLines
	lock                 *sync.Mutex
}

func (txtLineBrk *TextLineBreakField) CopyOut() TextLineBreakField {

	if txtLineBrk.lock == nil {
		txtLineBrk.lock = new(sync.Mutex)
	}

	txtLineBrk.lock.Lock()

	defer txtLineBrk.lock.Unlock()

	newLineBreak := TextLineBreakField{}

	newLineBreak.CreateLineBreak = txtLineBrk.CreateLineBreak
	newLineBreak.LeadingBlankLines = txtLineBrk.LeadingBlankLines.CopyOut()
	newLineBreak.LeftMargin = txtLineBrk.LeftMargin.CopyOut()
	newLineBreak.LeftSpacer = txtLineBrk.LeftSpacer.CopyOut()
	newLineBreak.LineSpec = txtLineBrk.LineSpec.CopyOut()
	newLineBreak.RightSpacer = txtLineBrk.RightSpacer.CopyOut()
	newLineBreak.TerminateWithNewLine = txtLineBrk.TerminateWithNewLine.CopyOut()
	newLineBreak.TrailingBlankLines = txtLineBrk.TrailingBlankLines.CopyOut()

	return newLineBreak
}

// NewPtr - Returns a pointer to a new TextLineBreakField
// instance.
func (txtLineBrk TextLineBreakField) NewPtr() *TextLineBreakField {

	if txtLineBrk.lock == nil {
		txtLineBrk.lock = new(sync.Mutex)
	}

	txtLineBrk.lock.Lock()

	defer txtLineBrk.lock.Unlock()

	newLineBreak := TextLineBreakField{}

	return &newLineBreak
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (txtLineBrk TextLineBreakField) TextTypeName() string {

	if txtLineBrk.lock == nil {
		txtLineBrk.lock = new(sync.Mutex)
	}

	txtLineBrk.lock.Lock()

	defer txtLineBrk.lock.Unlock()

	return "TextLineBreakField"
}

package strmech

type LineBreakField struct {
	CreateLineBreak      bool
	LeadingBlankLines    BlankLinesSpec
	LeftMargin           MarginSpec
	LeftSpacer           MarginSpec
	LineSpec             LineSpec
	RightSpacer          MarginSpec
	TerminateWithNewLine NewLineSpec
	TrailingBlankLines   BlankLinesSpec
}

func (lineBrk *LineBreakField) CopyOut() LineBreakField {

	newLineBreak := LineBreakField{}

	newLineBreak.CreateLineBreak = lineBrk.CreateLineBreak
	newLineBreak.LeadingBlankLines = lineBrk.LeadingBlankLines.CopyOut()
	newLineBreak.LeftMargin = lineBrk.LeftMargin.CopyOut()
	newLineBreak.LeftSpacer = lineBrk.LeftSpacer.CopyOut()
	newLineBreak.LineSpec = lineBrk.LineSpec.CopyOut()
	newLineBreak.RightSpacer = lineBrk.RightSpacer.CopyOut()
	newLineBreak.TerminateWithNewLine = lineBrk.TerminateWithNewLine.CopyOut()
	newLineBreak.TrailingBlankLines = lineBrk.TrailingBlankLines.CopyOut()

	return newLineBreak
}

// NewPtr - Returns a pointer to a new LineBreakField
// instance.
func (lineBrk LineBreakField) NewPtr() *LineBreakField {

	newLineBreak := LineBreakField{}

	return &newLineBreak
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (lineBrk LineBreakField) TextTypeName() string {
	return "LineBreakField"
}

type NumericIntField struct {
	LeftMargin           MarginSpec
	LeftSpacer           MarginSpec
	NumberSpec           IntegerSpec
	RightSpacer          MarginSpec
	TerminateWithNewLine NewLineSpec
}

// CopyOut - Returns a deep copy of the current NumericIntField
// instance.
func (numIntField *NumericIntField) CopyOut() NumericIntField {

	newNumField := NumericIntField{}

	newNumField.LeftMargin = numIntField.LeftMargin.CopyOut()
	newNumField.LeftSpacer = numIntField.LeftSpacer.CopyOut()
	newNumField.NumberSpec = numIntField.NumberSpec.CopyOut()
	newNumField.RightSpacer = numIntField.RightSpacer.CopyOut()
	newNumField.TerminateWithNewLine = numIntField.TerminateWithNewLine.CopyOut()

	return newNumField
}

// NewPtr - Returns a pointer to a new NumericIntField instance.
func (numIntField NumericIntField) NewPtr() *NumericIntField {

	newNumIntSpec := NumericIntField{}

	return &newNumIntSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (numIntField NumericIntField) TextTypeName() string {
	return "NumericIntField"
}

type OneLabelOneIntField struct {
	LeadingBlankLines  BlankLinesSpec
	TopLineBreak       LineBreakField
	Label1             StringField
	Number1            NumericIntField
	BottomLineBreak    LineBreakField
	TrailingBlankLines BlankLinesSpec
}

// CopyOut - Returns a deep copy of the current OneLabelOneIntField
// instance.
func (labelIntLine *OneLabelOneIntField) CopyOut() OneLabelOneIntField {

	newLabelInt := OneLabelOneIntField{}

	newLabelInt.LeadingBlankLines = labelIntLine.LeadingBlankLines
	newLabelInt.TopLineBreak = labelIntLine.TopLineBreak.CopyOut()
	newLabelInt.Label1 = labelIntLine.Label1.CopyOut()
	newLabelInt.Number1 = labelIntLine.Number1.CopyOut()
	newLabelInt.BottomLineBreak = labelIntLine.BottomLineBreak.CopyOut()
	newLabelInt.TrailingBlankLines = labelIntLine.TrailingBlankLines
	return newLabelInt
}

// NewPtr - Returns a pointer to a new OneLabelOneIntField instance.
func (labelIntLine OneLabelOneIntField) NewPtr() *OneLabelOneIntField {

	newOneLabelOneIntLine := OneLabelOneIntField{}

	return &newOneLabelOneIntLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (labelIntLine OneLabelOneIntField) TextTypeName() string {
	return "OneLabelOneIntField"
}

type StringField struct {
	LeftMargin           MarginSpec
	LeftSpacer           MarginSpec
	StrTxtSpec           StringSpec
	RightSpacer          MarginSpec
	TerminateWithNewLine NewLineSpec
}

func (strField *StringField) CopyOut() StringField {

	newStrField := StringField{}

	newStrField.LeftMargin = strField.LeftMargin.CopyOut()
	newStrField.LeftSpacer = strField.LeftSpacer.CopyOut()
	newStrField.StrTxtSpec = strField.StrTxtSpec.CopyOut()
	newStrField.RightSpacer = strField.RightSpacer.CopyOut()
	newStrField.TerminateWithNewLine = strField.TerminateWithNewLine

	return newStrField
}

// NewPtr - Returns a pointer to a new StringField instance.
func (strField StringField) NewPtr() *StringField {

	newStrFieldSpec := StringField{}

	return &newStrFieldSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (strField StringField) TextTypeName() string {
	return "StringField"
}

type TwoLabelStrField struct {
	LeadingBlankLines  BlankLinesSpec
	TopLineBreak       LineBreakField
	Label1             StringField
	Label2             StringField
	BottomLineBreak    LineBreakField
	TrailingBlankLines BlankLinesSpec
}

// CopyOut - Returns a deep copy of the current TwoLabelStrField
// instance.
func (twoLabelLine *TwoLabelStrField) CopyOut() TwoLabelStrField {

	newTwoLabelLine := TwoLabelStrField{}

	newTwoLabelLine.LeadingBlankLines = twoLabelLine.LeadingBlankLines.CopyOut()
	newTwoLabelLine.TopLineBreak = twoLabelLine.TopLineBreak.CopyOut()
	newTwoLabelLine.Label1 = twoLabelLine.Label1.CopyOut()
	newTwoLabelLine.Label2 = twoLabelLine.Label2.CopyOut()
	newTwoLabelLine.BottomLineBreak = twoLabelLine.BottomLineBreak.CopyOut()
	newTwoLabelLine.TrailingBlankLines = twoLabelLine.TrailingBlankLines.CopyOut()

	return newTwoLabelLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (twoLabelLine TwoLabelStrField) TextTypeName() string {
	return "TwoLabelStrField"
}

// NewPtr - Returns a pointer to a new TwoLabelStrField instance.
func (twoLabelLine TwoLabelStrField) NewPtr() *TwoLabelStrField {

	newTwoLabelLine := TwoLabelStrField{}

	return &newTwoLabelLine
}

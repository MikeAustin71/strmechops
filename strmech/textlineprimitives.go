package strmech

type BlankLinesSpec struct {
	NumBlankLines int
}

func (blkLine *BlankLinesSpec) CopyOut() BlankLinesSpec {
	newBlankLineSpec := BlankLinesSpec{}

	newBlankLineSpec.NumBlankLines = blkLine.NumBlankLines

	return newBlankLineSpec
}

// NewPtr - Returns a pointer to a new instance of 'BlankLinesSpec'
func (blkLine BlankLinesSpec) NewPtr() *BlankLinesSpec {

	newBlkLine := BlankLinesSpec{}

	return &newBlkLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (blkLine BlankLinesSpec) TextTypeName() string {
	return "BlankLinesSpec"
}

type IntegerSpec struct {
	NumericValue       int
	NumericFieldSpec   string
	NumericFieldLength int
	NumericPadChar     rune
	NumericPosition    FieldPositionSpec
}

func (intSpec *IntegerSpec) CopyOut() IntegerSpec {

	newInt := IntegerSpec{}
	newInt.NumericValue = intSpec.NumericValue
	newInt.NumericFieldSpec = intSpec.NumericFieldSpec
	newInt.NumericFieldLength = intSpec.NumericFieldLength
	newInt.NumericPadChar = intSpec.NumericPadChar

	return newInt
}

// NewPtr - Returns a pointer to a new LineSpec
// instance.
func (intSpec IntegerSpec) NewPtr() *IntegerSpec {

	anotherIntTxtSpec := IntegerSpec{}

	return &anotherIntTxtSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (intSpec IntegerSpec) TextTypeName() string {
	return "IntegerSpec"
}

type LineSpec struct {
	LineChar         rune
	LineLength       int
	LineFieldLength  int
	LineFieldPadChar rune
	LinePosition     FieldPositionSpec
}

// CopyOut - Returns a deep copy of LineSpec
// data fields.
func (lineSpec *LineSpec) CopyOut() LineSpec {

	newLineSpec := LineSpec{
		LineChar:         lineSpec.LineChar,
		LineLength:       lineSpec.LineLength,
		LineFieldLength:  lineSpec.LineFieldLength,
		LineFieldPadChar: lineSpec.LineFieldPadChar,
		LinePosition:     lineSpec.LinePosition,
	}

	return newLineSpec
}

// NewPtr - Returns a pointer to a new LineSpec
// instance.
func (lineSpec LineSpec) NewPtr() *LineSpec {

	anotherLineSpec := LineSpec{}

	return &anotherLineSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (lineSpec LineSpec) TextTypeName() string {
	return "LineSpec"
}

type MarginSpec struct {
	MarginStr    string
	MarginLength int
	MarginChar   rune
}

func (margin *MarginSpec) CopyOut() MarginSpec {

	newMargin := MarginSpec{}

	newMargin.MarginStr = margin.MarginStr
	newMargin.MarginChar = margin.MarginChar
	newMargin.MarginLength = margin.MarginLength

	return newMargin
}

// NewPtr - Returns a pointer to a new MarginSpec
// instance.
func (margin MarginSpec) NewPtr() *MarginSpec {

	newMarginField := MarginSpec{}

	return &newMarginField
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (margin MarginSpec) TextTypeName() string {
	return "MarginSpec"
}

type NewLineSpec struct {
	AddNewLine bool
}

func (newLine *NewLineSpec) CopyOut() NewLineSpec {
	newNewLine := NewLineSpec{}

	newNewLine.AddNewLine = newLine.AddNewLine

	return newNewLine
}

// NewPtr - Returns a pointer to a new NewLineSpec
// instance.
func (newLine NewLineSpec) NewPtr() *NewLineSpec {

	newLineSpec := NewLineSpec{}

	return &newLineSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (newLine NewLineSpec) TextTypeName() string {
	return "NewLineSpec"
}

type StringSpec struct {
	StrValue       string
	StrFieldLength int
	StrPadChar     rune
	StrPosition    FieldPositionSpec
}

// CopyOut - Returns a deep copy of the current StringSpec
// instance.
func (strTxtSpec *StringSpec) CopyOut() StringSpec {

	newStrSpec := StringSpec{}
	newStrSpec.StrValue = strTxtSpec.StrValue
	newStrSpec.StrFieldLength = strTxtSpec.StrFieldLength
	newStrSpec.StrPadChar = strTxtSpec.StrPadChar

	return newStrSpec
}

// NewPtr - Returns a pointer to a new StringSpec instance.
func (strTxtSpec StringSpec) NewPtr() *StringSpec {

	newStrTxtSpec := StringSpec{}

	return &newStrTxtSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextSpec interface
func (strTxtSpec StringSpec) TextTypeName() string {
	return "StringSpec"
}

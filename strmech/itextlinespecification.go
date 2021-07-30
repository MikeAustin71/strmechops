package strmech

type ITextLineSpecification interface {
	TextTypeName() string

	TextLineSpecName() string

	GetFormattedText() string

	CopyOutITextLine() ITextLineSpecification

	Empty()

	EqualITextLine(
		iTextLine ITextLineSpecification) bool

	String() string
}

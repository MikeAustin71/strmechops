package strmech

type ITextLineSpecification interface {
	TextLineSpecName() string

	GetFormattedText() string

	CopyOutITextLine() ITextLineSpecification

	Empty()
}

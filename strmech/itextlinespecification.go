package strmech

type ITextLineSpecification interface {
	TextLineSpecName() string

	GetFormattedText() string

	CopyOUtITextLine() ITextLineSpecification

	Empty()
}

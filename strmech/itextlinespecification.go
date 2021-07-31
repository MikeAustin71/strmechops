package strmech

type ITextLineSpecification interface {
	TextTypeName() string

	TextLineSpecName() string

	GetFormattedText() string

	CopyOutITextLine(interface{}) (
		ITextLineSpecification,
		error)

	Empty()

	EqualITextLine(
		iTextLine ITextLineSpecification) bool

	IsValidInstanceError(
		errorPrefix interface{}) error

	String() string
}

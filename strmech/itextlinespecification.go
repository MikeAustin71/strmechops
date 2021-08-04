package strmech

type ITextLineSpecification interface {
	GetFormattedText() string

	CopyOutITextLine(
		errorPrefix interface{}) (
		ITextLineSpecification,
		error)

	Empty()

	EqualITextLine(
		iTextLine ITextLineSpecification) bool

	IsValidInstanceError(
		errorPrefix interface{}) error

	String() string

	TextTypeName() string

	TextLineSpecName() string
}

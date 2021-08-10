package strmech

type ITextLineSpecification interface {
	CopyOutITextLine(
		errorPrefix interface{}) (
		ITextLineSpecification,
		error)

	Empty()

	EqualITextLine(
		iTextLine ITextLineSpecification) bool

	GetFormattedText() string

	IsValidInstanceError(
		errorPrefix interface{}) error

	String() string

	TextTypeName() string

	TextLineSpecName() string
}

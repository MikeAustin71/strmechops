package strmech

type ITextFieldSpecification interface {
	TextTypeName() string

	TextFieldName() string

	GetFormattedText() string

	CopyOutITextField(
		errorPrefix interface{}) (
		ITextFieldSpecification,
		error)

	Empty()

	EqualITextField(iTextField ITextFieldSpecification) bool

	IsValidInstanceError(errorPrefix interface{}) error

	String() string
}

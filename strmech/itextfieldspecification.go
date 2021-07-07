package strmech

type ITextFieldSpecification interface {
	TextTypeName() string

	TextFieldName() string

	GetFormattedText() string

	CopyOutITextField() ITextFieldSpecification

	Empty()

	EqualITextField(iTextField ITextFieldSpecification) bool
}

package strmech

type ITextFieldSpecification interface {
	TextTypeName() string

	TextFieldName() string

	GetFormattedText() string
}

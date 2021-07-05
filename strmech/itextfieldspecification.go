package strmech

type ITextFieldSpecification interface {
	IsValidInstanceError(interface{}) error
	GetFormattedText() string
	TextTypeName() string
	TextFieldName() string
}

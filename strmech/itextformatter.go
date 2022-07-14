package strmech

type ITTextFormatter interface {
	GetTextFieldType() TextFieldType
	SetTextFieldType(errorPrefix interface{}) error
}

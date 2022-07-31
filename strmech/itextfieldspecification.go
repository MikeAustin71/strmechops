package strmech

import "strings"

type ITextFieldSpecification interface {
	CopyOutITextField(
		errorPrefix interface{}) (
		ITextFieldSpecification,
		error)

	Empty()

	EqualITextField(
		iTextField ITextFieldSpecification) bool

	GetFormattedStrLength() int

	GetFormattedText(
		errorPrefix interface{}) (string, error)

	IsValidInstanceError(
		errorPrefix interface{}) error

	Read(p []byte) (n int, err error)

	ReaderInitialize()

	String() string

	TextTypeName() string

	TextFieldName() string

	TextBuilder(
		strBuilder *strings.Builder,
		errorPrefix interface{}) error
}

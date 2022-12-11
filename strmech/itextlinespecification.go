package strmech

import "strings"

// ITextLineSpecification - Types that support this interface
// also support the io.Reader interface.
type ITextLineSpecification interface {
	CopyOutITextLine(
		errorPrefix interface{}) (
		ITextLineSpecification,
		error)

	Empty()

	EqualITextLine(
		iTextLine ITextLineSpecification) bool

	GetFormattedText(
		errorPrefix interface{}) (string, error)

	IsValidInstanceError(
		errorPrefix interface{}) error

	Read(p []byte) (n int, err error)

	ReaderInitialize()

	String() string

	TextBuilder(
		strBuilder *strings.Builder,
		errorPrefix interface{}) error

	TextLineSpecName() string

	TextTypeName() string
}

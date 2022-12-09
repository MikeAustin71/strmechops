package strmech

// ITextFieldFormatDto
//
// This interface is implemented by types configured
// as Text Field Format Data Transfer Objects. These
// types are used to facilitate data entry for creation
// of text line formatting objects.
type ITextFieldFormatDto interface {
	CopyOutITextFieldFormat(
		errorPrefix interface{}) (
		ITextFieldFormatDto,
		error)

	Empty()

	GetFieldContentTextLabel(
		errorPrefix interface{}) (
		TextFieldSpecLabel,
		error)

	GetFieldFormatDtoType() string

	GetFormattedTextFieldStr(
		errorPrefix interface{}) (
		string,
		error)

	GetLeftMarginLength() int

	GetLeftMarginStr() string

	GetRightMarginLength() int

	GetRightMarginStr() string

	IsValidInstanceError(
		errorPrefix interface{}) error

	IsValidInstance() (isValid bool)
}

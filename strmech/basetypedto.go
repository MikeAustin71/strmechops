package strmech

// BaseTypeDto
//
// A data transfer object used to encapsulate elementary
// base types and transmit them between methods as
// required.
type BaseTypeDto struct {
	IsAByteArrayDto     bool
	AByteArrayDto       ByteArrayDto
	AByteArrayDtoLength int

	IsARuneArrayDto     bool
	ARuneArrayDto       RuneArrayDto
	ARuneArrayDtoLength int

	IsAString     bool
	AString       string
	AStringLength int

	IsAStringArrayDto     bool
	AStringArrayDto       StringArrayDto
	AStringArrayDtoLength int
}

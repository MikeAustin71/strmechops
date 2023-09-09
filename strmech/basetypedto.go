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
	AByteArrayDtoDesc1  string
	AByteArrayDtoDesc2  string

	IsARuneArrayDto     bool
	ARuneArrayDto       RuneArrayDto
	ARuneArrayDtoLength int
	ARuneArrayDtoDesc1  string
	ARuneArrayDtoDesc2  string

	IsAString     bool
	AString       string
	AStringLength int
	AStringDesc1  string
	AStringDesc2  string

	IsAStringArrayDto     bool
	AStringArrayDto       StringArrayDto
	AStringArrayDtoLength int
	AStringArrayDtoDesc1  string
	AStringArrayDtoDesc2  string
}

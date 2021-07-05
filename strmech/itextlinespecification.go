package strmech

type ITextLineSpecification interface {
	IsValidInstanceError(interface{}) error
	GetFormattedText() string
	TextTypeName() string
	TextLineName() string
}

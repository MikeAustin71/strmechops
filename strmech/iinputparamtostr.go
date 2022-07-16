package strmech

import "math/big"

type IInputParamToStr interface {
	*big.Int | *big.Float | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64 | ~string | ~bool
}

type IAddColumns interface {
	AddLine1Col(
		column1FieldText IInputParamToStr,
		errorPrefix interface{}) error
}

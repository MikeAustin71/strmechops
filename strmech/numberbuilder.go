package strmech

import "sync"

type NumberBuilder struct {
	numberType        NumStrFormatTypeCode
	integerElement    []rune
	fractionalElement []rune
	decimalSeparator  []rune
	currencySymbol    []rune
	numberSignSymbol  []rune
	integerSeparator  IntegerSeparatorDto
	numberFieldLen    int
	lock              *sync.Mutex
}

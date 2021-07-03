package strmech

type NumberBuilder struct {
	integerElement    []rune
	fractionalElement []rune
	decimalSeparator  FractionalSeparatorDto
	currencySymbol    FractionalSeparatorDto
	numberSignSymbol  NumberSignSymbolDto
	integerSeparator  IntegerSeparatorDto
}

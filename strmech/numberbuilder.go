package strmech

type NumberBuilder struct {
	integerElement    []rune
	fractionalElement []rune
	decimalSeparator  DecimalSeparatorDto
	currencySymbol    DecimalSeparatorDto
	numberSignSymbol  NumberSignSymbolDto
	integerSeparator  IntegerSeparatorDto
}

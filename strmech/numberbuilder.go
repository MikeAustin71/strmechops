package strmech

type NumberBuilder struct {
	integerElement    []rune
	fractionalElement []rune
	decimalSeparator  NumericSymbolDto
	currencySymbol    NumericSymbolDto
	numberSignSymbol  NumberSignSymbolDto
}

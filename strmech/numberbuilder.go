package strmech

type NumberBuilder struct {
	integerElement    []rune
	fractionalElement []rune
	decimalSeparator  NumericSymbolDto
	numberSignSymbol  NumberSignSymbolDto
}

package strmech

import "sync"

// FractionalSeparatorDto - A fractional separator is a text
// character or characters used to separate the integer and
// fractional components of a floating point number. The text
// character or characters which comprise the fractional separator
// are used in number strings to separate integer and fractional
// digits.
//
// In the USA, the fractional separator is the period character
// which is termed the decimal point ('.').
//   Example 123.34
//
type FractionalSeparatorDto struct {
	numericSymbol           []rune
	numericSymbolClass      NumericSymbolClass
	numericSymbolLocation   NumericSymbolLocation
	numericSymbolFound      bool
	numericSymbolFoundIndex int
	lock                    *sync.Mutex
}

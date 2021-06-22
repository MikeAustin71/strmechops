package strmech

import "sync"

type NumericSymbolDto struct {
	numericSymbol           []rune
	numericSymbolClass      NumericSymbolClass
	numericSymbolLocation   NumericSymbolLocation
	numericSymbolFound      bool
	numericSymbolFoundIndex int
	lock                    *sync.Mutex
}

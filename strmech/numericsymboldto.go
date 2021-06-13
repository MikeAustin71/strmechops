package strmech

import "sync"

type NumericSymbolDto struct {
	numericSymbol           []rune
	numericSymbolClass      NumericSymbolClass
	numericSymbolFound      bool
	numericSymbolFoundIndex int
	lock                    *sync.Mutex
}

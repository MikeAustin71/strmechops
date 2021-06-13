package strmech

import "sync"

type NumericSymbolDto struct {
	numericSymbol           []rune
	numericSymbolClass      string
	numericSymbolFound      bool
	numericSymbolFoundIndex int
	lock                    *sync.Mutex
}

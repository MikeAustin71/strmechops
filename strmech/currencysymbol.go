package strmech

import "sync"

type CurrencySymbolDto struct {
	currencyCode          string
	currencyCodeNo        string
	majorCurrencyName     string
	majorCurrencySymbol   []rune
	standardDecimalDigits uint
	minorCurrencyName     string
	minorCurrencySymbol   []rune
	lock                  *sync.Mutex
}

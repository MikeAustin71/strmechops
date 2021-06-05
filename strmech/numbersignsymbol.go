package strmech

import "sync"

// NumberSignSymbol - Defines the text or character value of a numeric
// sign contained in a number string. The most common examples of
// number sign text values are the plus sign ('+') and the minus
// sign ('-').
//
// The number sign text is usually positioned before the numeric
// value ('USA Example: +25') or after the numeric value ('EU
// Example 25-). However, there are cases where the number sign
// is positioned before and after the negative value. As an
// example, in the USA opening and closing parentheses are used to
// designate a negative number "(55)".
//
// Generally, number signs consist of a single text character,
// however there may be cases where multiple characters are used
// to designate positive or negative values.
//
// Since Go does not directly support enumerations, the 'TextJustify'
// type has been adapted to function in a manner similar to classic
// enumerations. 'TextJustify' is declared as a type 'int'. The
// method names effectively represent an enumeration of text
// justification formats. These methods are listed as follows:
//
//
type NumberSignSymbol struct {
	leadingNumSignChars  []rune
	trailingNumSignChars []rune
	lock                 *sync.Mutex
}

// GetLeadingNumSignChars - Returns a deep copy of the leading
// number sign characters contained in this instance of
// NumberSignSymbol.
//
func (nSignSymbol *NumberSignSymbol) GetLeadingNumSignChars() []rune {

	if nSignSymbol.lock == nil {
		nSignSymbol.lock = new(sync.Mutex)
	}

	nSignSymbol.lock.Lock()

	defer nSignSymbol.lock.Unlock()

	var leadingNumSignChars []rune

	lenNumSignChars := len(nSignSymbol.leadingNumSignChars)

	if lenNumSignChars == 0 {
		return leadingNumSignChars
	}

	leadingNumSignChars = make([]rune, lenNumSignChars)

	copy(
		leadingNumSignChars,
		nSignSymbol.leadingNumSignChars)

	return leadingNumSignChars
}

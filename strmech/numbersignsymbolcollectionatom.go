package strmech

import "sync"

type numberSignSymbolCollectionAtom struct {
	lock *sync.Mutex
}

// isTrailingNSignSymbolFound - Receives an array of
// NumberSignSymbolDto objects and returns a boolean flag signaling
// whether a trailing number sign symbol has been located.
//
// If a trailing number sign symbol has been located, the return
// parameter 'isTrailingNumberSignFound' is set to 'true' and the
// array index of the located NumberSignSymbolDto is returned in
// parameter 'trailingNumberSignFoundIndex'.
//
func (nSignSymColAtom *numberSignSymbolCollectionAtom) isTrailingNSignSymbolFound(
	nSignSymbols []NumberSignSymbolDto) (
	isTrailingNumberSignFound bool,
	trailingNumberSignFoundIndex int) {

	if nSignSymColAtom.lock == nil {
		nSignSymColAtom.lock = new(sync.Mutex)
	}

	nSignSymColAtom.lock.Lock()

	defer nSignSymColAtom.lock.Unlock()

	isTrailingNumberSignFound = false
	trailingNumberSignFoundIndex = -1

	lenSymCol := len(nSignSymbols)

	if lenSymCol == 0 {
		return isTrailingNumberSignFound, trailingNumberSignFoundIndex
	}

	for i := 0; i < lenSymCol; i++ {
		if nSignSymbols[i].trailingNumSignFoundInNumber {
			isTrailingNumberSignFound = true
			trailingNumberSignFoundIndex = i
			return isTrailingNumberSignFound, trailingNumberSignFoundIndex
		}
	}

	return isTrailingNumberSignFound, trailingNumberSignFoundIndex
}

// ptr - Returns a pointer to a new instance of
// numberSignSymbolCollectionAtom.
//
func (nSignSymColAtom numberSignSymbolCollectionAtom) ptr() *numberSignSymbolCollectionAtom {

	if nSignSymColAtom.lock == nil {
		nSignSymColAtom.lock = new(sync.Mutex)
	}

	nSignSymColAtom.lock.Lock()

	defer nSignSymColAtom.lock.Unlock()

	return &numberSignSymbolCollectionAtom{
		lock: new(sync.Mutex),
	}
}

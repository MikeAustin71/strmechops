package strmech

import (
	"sync"
)

type numberSignSymbolCollectionAtom struct {
	lock *sync.Mutex
}

// getLeadingNSignSymbolFound - Receives an array of
// NumberSignSymbolDto objects and returns a boolean flag signaling
// whether a leading number sign symbol has been located in the
// array.
//
// If a leading number sign symbol has been located, the return
// parameter 'isLeadingNumberSignFound' is set to 'true' and the
// 'nSignSymbols' array index of the located NumberSignSymbolDto
// is returned in parameter 'leadingNumberSignFoundIndex'.
//
// If multiple leading number symbols have been located, only the
// leading number symbol with the highest index in the host runes
// array will be selected and returned.
//
func (nSignSymColAtom *numberSignSymbolCollectionAtom) getLeadingNSignSymbolFound(
	nSignSymbols []NumberSignSymbolDto) (
	isLeadingNumberSignFound bool,
	leadingNumberSignFoundIndex int) {

	if nSignSymColAtom.lock == nil {
		nSignSymColAtom.lock = new(sync.Mutex)
	}

	nSignSymColAtom.lock.Lock()

	defer nSignSymColAtom.lock.Unlock()

	isLeadingNumberSignFound = false
	leadingNumberSignFoundIndex = -1

	lenSymCol := len(nSignSymbols)

	if lenSymCol == 0 {
		return isLeadingNumberSignFound, leadingNumberSignFoundIndex
	}

	highestIndex := -1
	collectionIndex := -1

	for i := 0; i < lenSymCol; i++ {

		if nSignSymbols[i].leadingNumSignFoundInNumber &&
			nSignSymbols[i].leadingNumSignFoundIndex > highestIndex &&
			(nSignSymbols[i].numSignPosition == NumSymPos.Before() ||
				nSignSymbols[i].numSignPosition == NumSymPos.BeforeAndAfter()) {

			if nSignSymbols[i].numSignPosition == NumSymPos.BeforeAndAfter() &&
				!nSignSymbols[i].trailingNumSignFoundInNumber {

				continue
			}

			highestIndex = nSignSymbols[i].leadingNumSignFoundIndex
			collectionIndex = i
			isLeadingNumberSignFound = true

		}
	}

	if isLeadingNumberSignFound {
		leadingNumberSignFoundIndex = collectionIndex
	}

	return isLeadingNumberSignFound, leadingNumberSignFoundIndex
}

// getTrailingNSignSymbolFound - Receives an array of
// NumberSignSymbolDto objects and returns a boolean flag signaling
// whether a trailing number sign symbol has been located in the
// array.
//
// If a trailing number sign symbol has been located, the return
// parameter 'isTrailingNumberSignFound' is set to 'true' and the
// 'nSignSymbols' array index of the located NumberSignSymbolDto is
// returned in parameter 'trailingNumberSignFoundIndex'.
//
// If multiple trailing number symbols have been located, only the
// first trailing number symbol or the trailing number symbol with
// the lowest index in the 'nSignSymbols' array will be selected
// and returned.
//
func (nSignSymColAtom *numberSignSymbolCollectionAtom) getTrailingNSignSymbolFound(
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

		if nSignSymbols[i].trailingNumSignFoundInNumber &&
			(nSignSymbols[i].numSignPosition == NumSymPos.After() ||
				nSignSymbols[i].numSignPosition == NumSymPos.BeforeAndAfter()) {

			if nSignSymbols[i].numSignPosition == NumSymPos.BeforeAndAfter() &&
				!nSignSymbols[i].leadingNumSignFoundInNumber {

				continue
			}

			isTrailingNumberSignFound = true
			trailingNumberSignFoundIndex = i

			return isTrailingNumberSignFound, trailingNumberSignFoundIndex
		}
	}

	return isTrailingNumberSignFound, trailingNumberSignFoundIndex
}

// isTrailingNSignSymbolFound - Receives an array of
// NumberSignSymbolDto objects and returns a boolean flag signaling
// whether a trailing number sign symbol has been located in the
// array.
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

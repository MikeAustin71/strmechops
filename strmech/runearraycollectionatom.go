package strmech

import "sync"

// runeArrayCollectionAtom - Provides helper methods for type
// RuneArrayCollection.
//
type runeArrayCollectionAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// RuneArrayCollection and proceeds to reset the data values
// for member values to their initial or zero values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'runeArrayCol' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  runeArrayCol           *RuneArrayCollection
//     - A pointer to an instance of RuneArrayCollection. All
//       the internal member variables contained in this instance
//       will be deleted and reset to their zero values.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (runeArrayColAtom *runeArrayCollectionAtom) empty(
	runeArrayCol *RuneArrayCollection) {

	if runeArrayColAtom.lock == nil {
		runeArrayColAtom.lock = new(sync.Mutex)
	}

	runeArrayColAtom.lock.Lock()

	defer runeArrayColAtom.lock.Unlock()

	if runeArrayCol == nil {
		return
	}

	lenRuneArrayCol := len(runeArrayCol.runeArrayDtoCol)

	if lenRuneArrayCol == 0 {

		runeArrayCol.runeArrayDtoCol = nil

		return
	}

	for i := 0; i < lenRuneArrayCol; i++ {

		runeArrayCol.runeArrayDtoCol[i].Empty()

	}

	runeArrayCol = nil

	return
}

// ptr - Returns a pointer to a new instance of
// runeArrayCollectionAtom.
//
func (runeArrayColAtom runeArrayCollectionAtom) ptr() *runeArrayCollectionAtom {

	if runeArrayColAtom.lock == nil {
		runeArrayColAtom.lock = new(sync.Mutex)
	}

	runeArrayColAtom.lock.Lock()

	defer runeArrayColAtom.lock.Unlock()

	return &runeArrayCollectionAtom{
		lock: new(sync.Mutex),
	}
}

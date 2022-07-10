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

// equal - Receives a pointer to two instances of
// RuneArrayCollection and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  runeArrayCol1  *RuneArrayCollection
//     - A pointer to an instance of RuneArrayCollection. Internal
//       member variables from 'runeArrayCol1' will be compared to
//       those of 'runeArrayCol2' to determine if both instances
//       are equivalent.
//
//
//  runeArrayCol2  *RuneArrayCollection
//     - A pointer to an instance of RuneArrayCollection. Internal
//       member variables from 'runeArrayCol2' will be compared to
//       those of 'runeArrayCol1' to determine if both instances
//       are equivalent.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of 'runeArrayCol1' and 'runeArrayCol2'
//       shows that all internal member variables are equivalent,
//       this method will return a boolean value of 'true'.
//
//       If the two instances are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (runeArrayColAtom *runeArrayCollectionAtom) equal(
	runeArrayCol1 *RuneArrayCollection,
	runeArrayCol2 *RuneArrayCollection) bool {

	if runeArrayColAtom.lock == nil {
		runeArrayColAtom.lock = new(sync.Mutex)
	}

	runeArrayColAtom.lock.Lock()

	defer runeArrayColAtom.lock.Unlock()

	if runeArrayCol1 == nil ||
		runeArrayCol2 == nil {

		return false
	}

	lenOfRuneArrayDtoCol := len(runeArrayCol1.runeArrayDtoCol)

	if lenOfRuneArrayDtoCol !=
		len(runeArrayCol2.runeArrayDtoCol) {

		return false
	}

	// Collection Lengths are Equal!
	if lenOfRuneArrayDtoCol == 0 {
		return true
	}

	for i := 0; i < lenOfRuneArrayDtoCol; i++ {

		if !runeArrayCol1.runeArrayDtoCol[i].Equal(
			&runeArrayCol2.runeArrayDtoCol[i]) {

			return false
		}
	}

	return true
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

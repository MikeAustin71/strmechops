package strmech

import "sync"

// textFormatterCollectionMolecule - Provides helper methods for
// TextFormatterCollection.
type textFormatterCollectionMolecule struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// TextFormatterCollection and proceeds to set all the internal
// member variables to their zero or uninitialized states.
//
// This method will therefore delete all data currently held
// by this instance of TextFormatterCollection.
func (textFmtCollectionMolecule *textFormatterCollectionMolecule) empty(
	txtFormatterCollection *TextFormatterCollection) {

	if textFmtCollectionMolecule.lock == nil {
		textFmtCollectionMolecule.lock = new(sync.Mutex)
	}

	textFmtCollectionMolecule.lock.Lock()

	defer textFmtCollectionMolecule.lock.Unlock()

	if txtFormatterCollection == nil {
		return
	}

	txtFmtColAtom := textFormatterCollectionAtom{}

	txtFmtColAtom.emptyFormatterCollection(
		txtFormatterCollection)

	txtFmtColAtom.emptyLineParamCollection(
		txtFormatterCollection)

	return
}

// equal - Receives pointers to two instances of TextFormatterCollection
// and proceeds to compare all the member data variables for both
// instances.
//
// If the two instances of TextFormatterCollection are found to be equal
// in all respects, this method will return a boolean value of
// 'true'.
func (textFmtCollectionMolecule *textFormatterCollectionMolecule) equal(
	txtFormatterCollection1 *TextFormatterCollection,
	txtFormatterCollection2 *TextFormatterCollection) bool {

	if textFmtCollectionMolecule.lock == nil {
		textFmtCollectionMolecule.lock = new(sync.Mutex)
	}

	textFmtCollectionMolecule.lock.Lock()

	defer textFmtCollectionMolecule.lock.Unlock()

	if txtFormatterCollection1 == nil ||
		txtFormatterCollection2 == nil {
		return false
	}

	lenItems1 := len(txtFormatterCollection1.fmtCollection)

	lenItems2 := len(txtFormatterCollection2.fmtCollection)

	if lenItems1 != lenItems2 {
		return false
	}

	for i := 0; i < lenItems1; i++ {

		if !txtFormatterCollection1.fmtCollection[i].Equal(
			txtFormatterCollection2.fmtCollection[i]) {

			return false
		}

	}

	lenItems1 = len(txtFormatterCollection1.stdTextLineParamCollection)

	lenItems2 = len(txtFormatterCollection2.stdTextLineParamCollection)

	if lenItems1 != lenItems2 {
		return false
	}

	for i := 0; i < lenItems1; i++ {

		if !txtFormatterCollection1.
			stdTextLineParamCollection[i].Equal(
			txtFormatterCollection2.
				stdTextLineParamCollection[i]) {

			return false
		}

	}

	return true
}

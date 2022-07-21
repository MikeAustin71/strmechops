package strmech

import "sync"

// textFormatterCollectionAtom - Provides helper methods for
// TextFormatterCollection.
type textFormatterCollectionAtom struct {
	lock *sync.Mutex
}

// emptyFormatterCollection - Receives a pointer to an instance of
// TextFormatterCollection and proceeds to delete all member
// elements of the Text Formatter Collection.
//
// Internal member variable
//		TextFormatterCollection.fmtCollection
// will be set to 'nil'.
//
// The TextFormatterCollection member variable that holds the Text
// Formatter Collection is defined as follows:
//
//   fmtCollection []TextFormatterDto
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  txtFormatterCollection     *TextFormatterCollection
//     - All member elements in the
//        txtFormatterCollection.fmtCollection will be deleted.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (textFmtCollectionAtom *textFormatterCollectionAtom) emptyFormatterCollection(
	txtFormatterCollection *TextFormatterCollection) {

	if textFmtCollectionAtom.lock == nil {
		textFmtCollectionAtom.lock = new(sync.Mutex)
	}

	textFmtCollectionAtom.lock.Lock()

	defer textFmtCollectionAtom.lock.Unlock()

	if txtFormatterCollection == nil {
		return
	}

	lenItems := len(txtFormatterCollection.fmtCollection)

	for i := 0; i < lenItems; i++ {

		txtFormatterCollection.fmtCollection[i].Empty()
	}

	txtFormatterCollection.fmtCollection = nil

	return
}

// emptyLineParamCollection - Receives a pointer to an instance of
// TextFormatterCollection and proceeds to delete all member
// elements in the Standard Text Line Parameters Collection.
//
// Internal member variable
//		TextFormatterCollection.stdTextLineParamCollection
// will be set to 'nil'.
//
// The internal member variable that holds the Standard Text Line
// Parameters Collection is defined as follows:
//
//   stdTextLineParamCollection []TextFmtParamsLineColumnsDto
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  txtFormatterCollection     *TextFormatterCollection
//     - All member elements in the
//        txtFormatterCollection.stdTextLineParamCollection will be
//        deleted.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (textFmtCollectionAtom *textFormatterCollectionAtom) emptyLineParamCollection(
	txtFormatterCollection *TextFormatterCollection) {

	if textFmtCollectionAtom.lock == nil {
		textFmtCollectionAtom.lock = new(sync.Mutex)
	}

	textFmtCollectionAtom.lock.Lock()

	defer textFmtCollectionAtom.lock.Unlock()

	if txtFormatterCollection == nil {
		return
	}

	lenItems :=
		len(txtFormatterCollection.stdTextLineParamCollection)

	for i := 0; i < lenItems; i++ {

		txtFormatterCollection.stdTextLineParamCollection[i].Empty()
	}

	txtFormatterCollection.stdTextLineParamCollection = nil

	return
}

// equal - Receives pointers to two instances of TextFormatterCollection
// and proceeds to compare all the member data variables for both
// instances.
//
// If the two instances of TextFormatterCollection are found to be equal
// in all respects, this method will return a boolean value of
// 'true'.
//
func (textFmtCollectionAtom *textFormatterCollectionAtom) equal(
	txtFormatterCollection1 *TextFormatterCollection,
	txtFormatterCollection2 *TextFormatterCollection) bool {

	if textFmtCollectionAtom.lock == nil {
		textFmtCollectionAtom.lock = new(sync.Mutex)
	}

	textFmtCollectionAtom.lock.Lock()

	defer textFmtCollectionAtom.lock.Unlock()

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

// ptr - Returns a pointer to a new instance of
// textFormatterCollectionAtom.
//
func (textFmtCollectionAtom textFormatterCollectionAtom) ptr() *textFormatterCollectionAtom {

	if textFmtCollectionAtom.lock == nil {
		textFmtCollectionAtom.lock = new(sync.Mutex)
	}

	textFmtCollectionAtom.lock.Lock()

	defer textFmtCollectionAtom.lock.Unlock()

	return &textFormatterCollectionAtom{
		lock: new(sync.Mutex),
	}
}

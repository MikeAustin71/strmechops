package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// runeArrayCollectionAtom - Provides helper methods for type
// RuneArrayCollection.
type runeArrayCollectionAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// RuneArrayCollection and proceeds to reset the data values
// for member values to their initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'runeArrayCol' will be deleted and reset to their zero
// values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayCol           *RuneArrayCollection
//	   - A pointer to an instance of RuneArrayCollection. All
//	     the internal member variables contained in this instance
//	     will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayCol1  *RuneArrayCollection
//	   - A pointer to an instance of RuneArrayCollection. Internal
//	     member variables from 'runeArrayCol1' will be compared to
//	     those of 'runeArrayCol2' to determine if both instances
//	     are equivalent.
//
//
//	runeArrayCol2  *RuneArrayCollection
//	   - A pointer to an instance of RuneArrayCollection. Internal
//	     member variables from 'runeArrayCol2' will be compared to
//	     those of 'runeArrayCol1' to determine if both instances
//	     are equivalent.
//
// ----------------------------------------------------------------
//
// Return Values
//
//	bool
//	   - If the comparison of 'runeArrayCol1' and 'runeArrayCol2'
//	     shows that all internal member variables are equivalent,
//	     this method will return a boolean value of 'true'.
//
//	     If the two instances are NOT equal, this method will
//	     return a boolean value of 'false' to the calling function.
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

// peekPopTextLine - Performs either a 'Peek' or 'Pop' operation
// on an array element in the Text Lines Collection as specified
// by the input parameter, 'popTextLine'.
//
// A 'Pop' operation returns a deep copy of the designated Text
// Line element in the Text Lines Collection and then DELETES that
// designated array element. The designated array element is
// specified by input parameter, 'zeroBasedIndex'.
//
// On the other hand, a 'Peek' operation will return a deep copy of
// the designated Text Line in the Text lines Collection and WILL
// NOT delete that array element. The designated array element
// therefore remains in the collection after the 'Peek' operation
// is completed.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayCol               *RuneArrayCollection
//	   - A pointer to an instance of RuneArrayCollection. A deep
//	     copy of the designated Rune Array Dto in the Rune Array
//	     Dto Collection for this instance of RuneArrayCollection
//	     will be returned to the calling function. The returned
//	     RuneArrayDto object is designated by input parameter,
//	     'zeroBasedIndex'.
//
//	     Depending on the value of input parameter,
//	     'popCollectionElement', either a 'Peek' or 'Pop' operation
//	     will be performed on the designated RuneArrayDto object in
//	     the Rune Arrays Dto Collection
//	     ('RuneArrayCollection.runeArrayDtoCol').
//
//
//	zeroBasedIndex             int
//	   - The index number of the array element in the Rune Arrays
//	     Collection on which the 'Pop' or 'Peek' operation will be
//	     performed.
//
//
//	popCollectionElement       bool
//	   - If this parameter is set to 'true', it signals that a
//	     'Pop' operation will be performed on the designated Rune
//	     Array Dto object in the Rune Arrays Collection
//	     encapsulated in parameter 'runeArrayCol'.
//
//	     A 'Pop' operation will DELETE the designated Text Field
//	     from the Rune Arrays Collection.
//
//	     If this parameter is set to 'false', it signals that a
//	     'Peek' operation will be performed on the designated
//	     RuneArrayDto object in the Rune Arrays Collection
//	     encapsulated in parameter 'runeArrayCol'. A 'Peek'
//	     operation means that the designated RuneArrayDto element
//	     in the Rune Arrays Collection WILL NOT be deleted and will
//	     remain in the collection.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	targetRuneArrayDto         RuneArrayCollection
//	   - If this method completes successfully, a deep copy of
//	     if the RuneArrayDto object specified by array index
//	     'zeroBasedIndex' in the Rune Arrays Collection of input
//	     parameter 'runeArrayCol' will be returned to the calling
//	     function.
//
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (runeArrayColAtom *runeArrayCollectionAtom) peekPopRuneArrayCol(
	runeArrayCol *RuneArrayCollection,
	zeroBasedIndex int,
	popCollectionElement bool,
	errPrefDto *ePref.ErrPrefixDto) (
	targetRuneArrayDto RuneArrayDto,
	err error) {

	if runeArrayColAtom.lock == nil {
		runeArrayColAtom.lock = new(sync.Mutex)
	}

	runeArrayColAtom.lock.Lock()

	defer runeArrayColAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayCollectionNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return targetRuneArrayDto, err

	}

	if runeArrayCol == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeArrayCol' is a nil pointer!\n",
			ePrefix.String())

		return targetRuneArrayDto, err
	}

	lenRuneArrayCol := len(runeArrayCol.runeArrayDtoCol)

	if lenRuneArrayCol == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"The Rune Array Dto Collection, 'runeArrayCol.runeArrayDtoCol' is EMPTY!\n",
			ePrefix.String())

		return targetRuneArrayDto, err
	}

	if zeroBasedIndex < 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'zeroBasedIndex' is invalid!\n"+
			"'zeroBasedIndex' is less than zero.\n"+
			"zeroBasedIndex = '%v'\n",
			ePrefix.String(),
			zeroBasedIndex)

		return targetRuneArrayDto, err
	}

	if zeroBasedIndex >= lenRuneArrayCol {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'zeroBasedIndex' is invalid!\n"+
			"'zeroBasedIndex' is greater than the last index\n"+
			"in the Rune Array Dto Collection.\n"+
			"Last index in collection = '%v'\n"+
			"zeroBasedIndex = '%v'\n",
			ePrefix.String(),
			lenRuneArrayCol-1,
			zeroBasedIndex)

		return targetRuneArrayDto, err
	}

	targetRuneArrayDto,
		err = runeArrayCol.runeArrayDtoCol[zeroBasedIndex].CopyOut(
		ePrefix.XCpy(
			fmt.Sprintf("targetRuneArrayDto"+
				"<-runeArrayCol.runeArrayDtoCol[%v]",
				zeroBasedIndex)))

	if err != nil {

		return targetRuneArrayDto, err

	}

	if !popCollectionElement {

		return targetRuneArrayDto, err

	}

	// popCollectionElement == true
	// Now, Delete Array Element at
	// Index == zeroBasedIndex
	err = runeArrayCollectionQuark{}.ptr().deleteCollectionElement(
		runeArrayCol,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf("Delete runeArrayCol."+
				"runeArrayDtoCol[%v]",
				zeroBasedIndex)))

	return targetRuneArrayDto, err
}

// ptr - Returns a pointer to a new instance of
// runeArrayCollectionAtom.
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

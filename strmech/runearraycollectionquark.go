package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// runeArrayCollectionQuark - Provides helper methods for type
// RuneArrayCollection.
//
type runeArrayCollectionQuark struct {
	lock *sync.Mutex
}

// deleteTextLineElement - Deletes a member of the Text Lines
// Collection. The array element to be deleted is designated by
// input parameter 'zeroBasedIndex'.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  runeArrayCol               *RuneArrayCollection
//     - A pointer to a RuneArrayCollection instance which
//       encapsulates the Rune Array Collection. The collection
//       element identified by input parameter
//       'deleteZeroBasedIndex' WILL BE DELETED.
//
//
//  deleteZeroBasedIndex       int
//     - The index number of the array element in the Rune Array
//       Collection ('runeArrayCol') which will be deleted.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil' signaling that the designated
//       Text Line element in the Text Lines Collection has been
//       deleted. If errors are encountered during processing, the
//       returned error Type will encapsulate an error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (runeArrayColQuark *runeArrayCollectionQuark) deleteCollectionElement(
	runeArrayCol *RuneArrayCollection,
	deleteZeroBasedIndex int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if runeArrayColQuark.lock == nil {
		runeArrayColQuark.lock = new(sync.Mutex)
	}

	runeArrayColQuark.lock.Lock()

	defer runeArrayColQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayCollectionQuark."+
			"deleteCollectionElement()",
		"")

	if err != nil {
		return err
	}

	if runeArrayCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeArrayCol' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenRuneArrayCol := len(runeArrayCol.runeArrayDtoCol)

	if lenRuneArrayCol == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The Rune Array Dto Collection is empty\n"+
			"and contains zero elements!\n",
			ePrefix.String())

		return err

	}

	if deleteZeroBasedIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'deleteZeroBasedIndex' is invalid.\n"+
			"'deleteZeroBasedIndex' is less than zero!\n"+
			"deleteZeroBasedIndex = '%v'\n",
			ePrefix.String(),
			deleteZeroBasedIndex)

		return err
	}

	lastIndex := lenRuneArrayCol - 1

	if deleteZeroBasedIndex > lastIndex {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'deleteZeroBasedIndex' is invalid.\n"+
			"The value of 'deleteZeroBasedIndex' is greater than the last\n"+
			"index in the Rune Array Dto Collection!\n"+
			"Last Collection Index = '%v'\n"+
			"deleteZeroBasedIndex  = '%v'\n",
			ePrefix.String(),
			lastIndex,
			deleteZeroBasedIndex)

		return err
	}

	runeArrayCol.runeArrayDtoCol[deleteZeroBasedIndex].Empty()

	if deleteZeroBasedIndex == 0 {
		// deleteZeroBasedIndex == 0
		runeArrayCol.runeArrayDtoCol =
			runeArrayCol.runeArrayDtoCol[1:]

	} else if deleteZeroBasedIndex == lastIndex {
		// deleteZeroBasedIndex == Last Index

		runeArrayCol.runeArrayDtoCol =
			runeArrayCol.runeArrayDtoCol[0:lastIndex]
	} else {
		// deleteZeroBasedIndex > 0 AND
		//  deleteZeroBasedIndex < lastIn

		runeArrayCol.runeArrayDtoCol = append(
			runeArrayCol.runeArrayDtoCol[0:deleteZeroBasedIndex],
			runeArrayCol.runeArrayDtoCol[deleteZeroBasedIndex+1:]...)
	}

	return err
}

// ptr - Returns a pointer to a new instance of
// runeArrayCollectionQuark.
//
func (runeArrayColQuark runeArrayCollectionQuark) ptr() *runeArrayCollectionQuark {

	if runeArrayColQuark.lock == nil {
		runeArrayColQuark.lock = new(sync.Mutex)
	}

	runeArrayColQuark.lock.Lock()

	defer runeArrayColQuark.lock.Unlock()

	return &runeArrayCollectionQuark{
		lock: new(sync.Mutex),
	}
}

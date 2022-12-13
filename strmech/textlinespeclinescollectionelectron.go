package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type textLineSpecLinesCollectionElectron struct {
	lock *sync.Mutex
}

// deleteTextLineElement - Deletes a member of the Text Lines
// Collection. The array element to be deleted is designated by
// input parameter 'zeroBasedIndex'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	textLinesCol               *TextLineSpecLinesCollection
//	   - A pointer to an instance of TextLineSpecLinesCollection
//	     which encapsulates the Text Lines Collection. The member
//	     of this collection designated by parameter,
//	     'zeroBasedIndex' WILL BE DELETED.
//
//
//	zeroBasedIndex             int
//	   - The index number of the array element in the Text Lines
//	     Collection which will be deleted.
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
// ------------------------------------------------------------------------
//
// Return Values
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil' signaling that the designated
//	     Text Line element in the Text Lines Collection has been
//	     deleted. If errors are encountered during processing, the
//	     returned error Type will encapsulate an error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (txtLinesColElectron textLineSpecLinesCollectionElectron) deleteTextLineElement(
	textLinesCol *TextLineSpecLinesCollection,
	zeroBasedIndex int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLinesColElectron.lock == nil {
		txtLinesColElectron.lock = new(sync.Mutex)
	}

	txtLinesColElectron.lock.Lock()

	defer txtLinesColElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecLinesCollectionElectron."+
			"deleteTextLineElement()",
		"")

	if err != nil {
		return err
	}

	if textLinesCol == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'textLinesCol' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenTextLines := len(textLinesCol.textLines)

	if lenTextLines == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The Text Lines Collection is empty\n"+
			"and contains zero elements!\n",
			ePrefix.String())

		return err

	}

	if zeroBasedIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'zeroBasedIndex' is invalid.\n"+
			"'zeroBasedIndex' is less than zero!\n"+
			"zeroBasedIndex = '%v'\n",
			ePrefix.String(),
			zeroBasedIndex)

		return err
	}

	lastIndex := lenTextLines - 1

	if zeroBasedIndex > lastIndex {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'zeroBasedIndex' is invalid.\n"+
			"The value of 'zeroBasedIndex' is greater than the last\n"+
			"index in the Text Lines Collection!\n"+
			"Last Collection Index = '%v'\n"+
			"zeroBasedIndex        = '%v'\n",
			ePrefix.String(),
			lastIndex,
			zeroBasedIndex)

		return err
	}

	if zeroBasedIndex == 0 {

		if textLinesCol.textLines[0] != nil {

			textLinesCol.textLines[0].Empty()

			textLinesCol.textLines[0] = nil

		}

		textLinesCol.textLines = textLinesCol.textLines[1:]

	} else if zeroBasedIndex == lastIndex {

		if textLinesCol.textLines[lastIndex] != nil {

			textLinesCol.textLines[lastIndex].Empty()

			textLinesCol.textLines[lastIndex] = nil

		}

		textLinesCol.textLines = textLinesCol.textLines[0:lastIndex]

	} else {

		if textLinesCol.textLines[zeroBasedIndex] != nil {

			textLinesCol.textLines[zeroBasedIndex].Empty()

			textLinesCol.textLines[zeroBasedIndex] = nil

		}

		textLinesCol.textLines = append(
			textLinesCol.textLines[0:zeroBasedIndex],
			textLinesCol.textLines[zeroBasedIndex+1:]...)

	}

	return err
}

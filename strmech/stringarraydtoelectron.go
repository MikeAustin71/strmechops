package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type stringArrayDtoElectron struct {
	lock *sync.Mutex
}

// deleteStringArrayElement - Deletes a member of the
// StringArrayDto string array. The array element to be deleted is
// designated by input parameter 'zeroBasedIndex'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strArray                   *StringArrayDto
//	   - A pointer to an instance of StringArrayDto which
//	     encapsulates the target string array. A member element
//	     of this string array designated by parameter,
//	     'zeroBasedIndex', WILL BE DELETED.
//
//
//	zeroBasedIndex             int
//	   - The index number of the array element in the
//	     StringArrayDto string array which will be deleted.
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
func (strArrayDtoElectron *stringArrayDtoElectron) deleteStringArrayElement(
	strArray *StringArrayDto,
	zeroBasedIndex int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if strArrayDtoElectron.lock == nil {
		strArrayDtoElectron.lock = new(sync.Mutex)
	}

	strArrayDtoElectron.lock.Lock()

	defer strArrayDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"stringArrayDtoElectron."+
			"deleteStringArrayElement()",
		"")

	if err != nil {
		return err
	}

	if strArray == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strArray' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenStrArray := len(strArray.StrArray)

	if lenStrArray == 0 {
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

	lastIndex := lenStrArray - 1

	if zeroBasedIndex > lastIndex {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'zeroBasedIndex' is invalid.\n"+
			"The value of 'zeroBasedIndex' is greater than the last\n"+
			"index in the String Array!\n"+
			"Last String Array Index = '%v'\n"+
			"zeroBasedIndex          = '%v'\n",
			ePrefix.String(),
			lastIndex,
			zeroBasedIndex)

		return err
	}

	if zeroBasedIndex == 0 {

		strArray.StrArray[0] = ""

		strArray.StrArray = strArray.StrArray[1:]

	} else if zeroBasedIndex == lastIndex {

		strArray.StrArray[lastIndex] = ""

		strArray.StrArray = strArray.StrArray[0:lastIndex]

	} else {

		strArray.StrArray[zeroBasedIndex] = ""

		strArray.StrArray = append(
			strArray.StrArray[0:zeroBasedIndex],
			strArray.StrArray[zeroBasedIndex+1:]...)

	}

	return err

}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type stringArrayDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// TextFormatterCollection and proceeds to set all the internal
// member variables to their zero or uninitialized states.
//
// This method will therefore delete all data currently held
// by this instance of TextFormatterCollection.
func (strArrayDtoAtom *stringArrayDtoAtom) empty(
	strArray *StringArrayDto) {

	if strArrayDtoAtom.lock == nil {
		strArrayDtoAtom.lock = new(sync.Mutex)
	}

	strArrayDtoAtom.lock.Lock()

	defer strArrayDtoAtom.lock.Unlock()

	if strArray == nil {
		return
	}

	strArray.StrArray = nil

	strArray.Description1 = ""

	strArray.Description2 = ""

}

// equal - Receives pointers to two instances of StringArrayDto
// and proceeds to compare all the member data variables for both
// instances.
//
// If the two instances of StringArrayDto are found to be equal
// in all respects, this method will return a boolean value of
// 'true'.
func (strArrayDtoAtom *stringArrayDtoAtom) equal(
	strArray1 *StringArrayDto,
	strArray2 *StringArrayDto) bool {

	if strArrayDtoAtom.lock == nil {
		strArrayDtoAtom.lock = new(sync.Mutex)
	}

	strArrayDtoAtom.lock.Lock()

	defer strArrayDtoAtom.lock.Unlock()

	lenStrArray1 := len(strArray1.StrArray)

	if lenStrArray1 != len(strArray2.StrArray) {

		return false
	}

	if lenStrArray1 > 0 {

		for i := 0; i < lenStrArray1; i++ {

			if strArray1.StrArray[i] !=
				strArray2.StrArray[i] {

				return false
			}

		}

	}

	if strArray1.Description1 !=
		strArray2.Description1 {

		return false
	}

	if strArray1.Description2 !=
		strArray2.Description2 {

		return false
	}

	return true
}

// peekPopStringArray - Performs either a 'Peek' or 'Pop' operation
// on a string array element in a String Array contained in an
// instance of StringArrayDto as specified by the input parameter,
// 'popStringArrayElement'.
//
// A 'Pop' operations returns a deep copy of the designated string
// array element in the StringArrayDto instance. 'Pop' operations
// will then DELETE the designated array element. That designated
// array element is specified by input parameter, 'zeroBasedIndex'.
//
// On the other hand, a 'Peek' operation will also return a deep
// copy of the designated string array element in the
// StringArrayDto instance. However, the 'Peek' operation WILL
// NOT delete the designated array element. The designated array
// element will therefore remain in the string array after the
// 'Peek' operation is completed.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strArrayDto                *StringArrayDto
//	   - A pointer to an instance of StringArrayDto. A deep copy of
//	     the designated string in the string array for this
//	     instance of StringArrayDto will be returned to the calling
//	     function. The returned string array element is designated
//	     by input parameter, 'zeroBasedIndex'.
//
//	     Depending on the value of input parameter,
//	     'popStringArrayElement', either a 'Peek' or 'Pop'
//	     operation will be performed on the designated string array
//	     contained the StringArrayDto parameter, 'strArrayDto'.
//
//
//	zeroBasedIndex             int
//	   - The index number of the array element in the 'strArrayDto'
//	     string array on which the 'Pop' or 'Peek' operation will
//	     be performed.
//
//
//	popTextLine                bool
//	   - If this parameter is set to 'true', it signals that a
//	     'Pop' operation will be performed on the designated string
//	     array in the StringArrayDto string array encapsulated in
//	     parameter 'strArrayDto'. A 'Pop' operation will DELETE the
//	     designated string array element from the string array
//	     maintained by the StringArrayDto instance, 'strArrayDto'.
//
//	     If this parameter is set to 'false', it signals that a
//	     'Peek' operation will be performed on the designated
//	     string array encapsulated in StringArrayDto parameter
//	     'strArrayDto'. A 'Peek' operation means that the
//	     designated string array element in the StringArrayDto
//	     string array WILL NOT be deleted and will remain in the
//	     string array.
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
//	targetStr                  string
//	   - If this method completes successfully, a copy of the
//	     string designated by the 'zeroBasedIndex' string array
//	     element in the 'strArrayDto' string array will be returned
//	     to the calling function.
//
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered
//	     during processing, the returned error Type will
//	     encapsulate an error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (strArrayDtoAtom *stringArrayDtoAtom) peekPopStringArray(
	strArrayDto *StringArrayDto,
	zeroBasedIndex int,
	popStringArrayElement bool,
	errPrefDto *ePref.ErrPrefixDto) (
	targetStr string,
	err error) {

	if strArrayDtoAtom.lock == nil {
		strArrayDtoAtom.lock = new(sync.Mutex)
	}

	strArrayDtoAtom.lock.Lock()

	defer strArrayDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"stringArrayDtoAtom."+
			"peekPopStringArray()",
		"")

	if err != nil {
		return targetStr, err
	}

	if strArrayDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strArrayDto' is a nil pointer!\n",
			ePrefix.String())

		return targetStr, err
	}

	lenStrArray := len(strArrayDto.StrArray)

	if lenStrArray == 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"The String Array, 'strArrayDto.StrArray' is EMPTY!\n",
			ePrefix.String())

		return targetStr, err
	}

	if zeroBasedIndex < 0 {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'zeroBasedIndex' is invalid!\n"+
			"'zeroBasedIndex' is less than zero.\n"+
			"zeroBasedIndex = '%v'\n",
			ePrefix.String(),
			zeroBasedIndex)

		return targetStr, err
	}

	lastIdx := lenStrArray - 1

	if zeroBasedIndex > lastIdx {

		err = fmt.Errorf("%v - ERROR\n"+
			"Input parameter 'zeroBasedIndex' is invalid!\n"+
			"'zeroBasedIndex' is greater than the last index\n"+
			"in the String Array.\n"+
			"Last index in String Array = '%v'\n"+
			"zeroBasedIndex = '%v'\n",
			ePrefix.String(),
			lastIdx,
			zeroBasedIndex)

		return targetStr, err
	}

	targetStr = strArrayDto.StrArray[zeroBasedIndex]

	if !popStringArrayElement {

		return targetStr, err

	}

	err = new(stringArrayDtoElectron).deleteStringArrayElement(
		strArrayDto,
		zeroBasedIndex,
		ePrefix.XCpy(
			fmt.Sprintf("delete StrArray[%v]",
				zeroBasedIndex)))

	return targetStr, err
}

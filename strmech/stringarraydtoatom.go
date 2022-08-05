package strmech

import (
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

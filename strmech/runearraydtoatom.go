package strmech

import "sync"

type runeArrayDtoAtom struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// RuneArrayDto and proceeds to reset the data values
// for member variables to their initial or zero values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'runeArrayDto' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  runeArrayDto               *RuneArrayDto
//     - A pointer to an instance of RuneArrayDto. All
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
func (runeDtoAtom *runeArrayDtoAtom) empty(
	runeArrayDto *RuneArrayDto) {

	if runeDtoAtom.lock == nil {
		runeDtoAtom.lock = new(sync.Mutex)
	}

	runeDtoAtom.lock.Lock()

	defer runeDtoAtom.lock.Unlock()

	if runeArrayDto == nil {
		return
	}

	new(runeArrayDtoElectron).
		emptyCharsArray(runeArrayDto)

	runeArrayDto.Description1 = ""

	runeArrayDto.Description2 = ""

	runeArrayDto.charSearchType = CharSearchType.LinearTargetStartingIndex()

	return
}

// equal - Receives a pointer to two instances of
// RuneArrayDto and proceeds to compare their member variables in
// order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (runeDtoAtom *runeArrayDtoAtom) equal(
	runeArrayDto1 *RuneArrayDto,
	runeArrayDto2 *RuneArrayDto) bool {

	if runeDtoAtom.lock == nil {
		runeDtoAtom.lock = new(sync.Mutex)
	}

	runeDtoAtom.lock.Lock()

	defer runeDtoAtom.lock.Unlock()

	if runeArrayDto1 == nil ||
		runeArrayDto2 == nil {

		return false
	}

	areEqual := new(runeArrayDtoElectron).
		equalCharArrays(
			runeArrayDto1,
			runeArrayDto2)

	if !areEqual {
		return false
	}

	if runeArrayDto1.Description1 !=
		runeArrayDto2.Description1 {
		return false
	}

	if runeArrayDto1.Description2 !=
		runeArrayDto2.Description2 {
		return false
	}

	if runeArrayDto1.charSearchType !=
		runeArrayDto2.charSearchType {
		return false
	}

	return true
}

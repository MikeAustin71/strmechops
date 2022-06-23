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
func (runeDtoAtom runeArrayDtoAtom) empty(
	runeArrayDto *RuneArrayDto) {

	if runeDtoAtom.lock == nil {
		runeDtoAtom.lock = new(sync.Mutex)
	}

	runeDtoAtom.lock.Lock()

	defer runeDtoAtom.lock.Unlock()

	if runeArrayDto == nil {
		return
	}

	runeArrayDtoElectron{}.ptr().
		emptyCharsArray(runeArrayDto)

	runeArrayDto.Description1 = ""

	runeArrayDto.Description2 = ""

	runeArrayDto.charSearchType = CharSearchType.LinearTargetStartingIndex()

	return
}

// ptr - Returns a pointer to a new instance of
// runeArrayDtoAtom.
//
func (runeDtoAtom runeArrayDtoAtom) ptr() *runeArrayDtoAtom {

	if runeDtoAtom.lock == nil {
		runeDtoAtom.lock = new(sync.Mutex)
	}

	runeDtoAtom.lock.Lock()

	defer runeDtoAtom.lock.Unlock()

	return &runeArrayDtoAtom{
		lock: new(sync.Mutex),
	}
}

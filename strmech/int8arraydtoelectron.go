package strmech

import "sync"

// int8ArrayDtoElectron - Provides helper methods for type
// Int8ArrayDto.
type int8ArrayDtoElectron struct {
	lock *sync.Mutex
}

//	emptyCharsArray
//
//	Empties or deletes all the characters residing in the
//	internal member variable 'IntsArray' contained in the
//	instance of Int8ArrayDto passed as input parameter
//	'i8ArrayDto'.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	i8ArrayDto			*Int8ArrayDto
//
//		A pointer to an instance of Int8ArrayDto. All the
//		integers contained in the internal member variable
//		int8 array 'IntsArray', will be deleted. This
//		internal member variable int8 array will then be
//		set to a value of 'nil'.
//
// ------------------------------------------------------------------------
//
// #Return Values
//
//	NONE
func (i8ArrayDtoElectron *int8ArrayDtoElectron) emptyIntsArray(
	i8ArrayDto *Int8ArrayDto) {

	if i8ArrayDtoElectron.lock == nil {
		i8ArrayDtoElectron.lock = new(sync.Mutex)
	}

	i8ArrayDtoElectron.lock.Lock()

	defer i8ArrayDtoElectron.lock.Unlock()

	if i8ArrayDto == nil {
		return
	}

	i8ArrayDto.IntsArray = nil

	return
}

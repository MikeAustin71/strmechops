package strmech

import (
	"sync"
)

// charSearchTargetInputParametersDtoElectron - Provides helper methods for type
// CharSearchTargetInputParametersDto.
//
type charSearchTargetInputParametersDtoElectron struct {
	lock *sync.Mutex
}

// emptyTargetStrings - Receives a pointer to an instance of
// CharSearchTargetInputParametersDto and proceeds to reset the
// internal member variable 'TargetString' to its zero or
// uninitialized state.
//
// Since member variable 'TargetString' is a pointer to an
// instance of RuneArrayDto, the 'TargetString' pointer is reset to
// 'nil'. This step is taken in order to preserve data sharing in
// case another object also holds a pointer to that same instance
// of RuneArrayDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms           *CharSearchTargetInputParametersDto
//     - Pointer to an instance of
//       CharSearchTargetInputParametersDto. The internal member
//       variable, 'TargetString', will be reset to a value of
//       'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchTargetInputParmsElectron *charSearchTargetInputParametersDtoElectron) emptyTargetStrings(
	targetInputParms *CharSearchTargetInputParametersDto) {

	if searchTargetInputParmsElectron.lock == nil {
		searchTargetInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTargetInputParmsElectron.lock.Lock()

	defer searchTargetInputParmsElectron.lock.Unlock()

	if targetInputParms == nil {
		return
	}

	targetInputParms.TargetString = nil

	return
}

// equalTargetStrings - Compares the Target Strings from two
// instances of CharSearchTargetInputParametersDto to determine if
// they are equivalent.
//
// If both TargetString member variables are 'nil' pointers, this
// method classifies them as equivalent.
//
// If both TargetString member variables are equal in all respects,
// this method returns a boolean value of 'true'. Otherwise, a
// value of 'false' is returned to the calling routine.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms1          *CharSearchTargetInputParametersDto
//     - An instance of CharSearchTargetInputParametersDto. If the
//       internal member variable, 'TargetString', will be compared
//       to the same internal member variable ('TargetString') in
//       parameter 'targetInputParms2' to determine if the two
//       target strings are equivalent.
//
//
//  targetInputParms2          *CharSearchTargetInputParametersDto
//     - An instance of CharSearchTargetInputParametersDto. If the
//       internal member variable, 'TargetString', will be compared
//       to the same internal member variable ('TargetString') in
//       parameter 'targetInputParms1' to determine if the two
//       target strings are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of target strings in parameters
//       'targetInputParms1' and 'targetInputParms2' show that the
//       internal member variables 'TargetString' are equivalent,
//       this method will return a boolean value of 'true'.
//
//       If the two target strings are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchTargetInputParmsElectron *charSearchTargetInputParametersDtoElectron) equalTargetStrings(
	targetInputParms1 *CharSearchTargetInputParametersDto,
	targetInputParms2 *CharSearchTargetInputParametersDto) bool {

	if searchTargetInputParmsElectron.lock == nil {
		searchTargetInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTargetInputParmsElectron.lock.Lock()

	defer searchTargetInputParmsElectron.lock.Unlock()

	if targetInputParms1 == nil ||
		targetInputParms2 == nil {
		return false
	}

	if targetInputParms1.TargetString == nil &&
		targetInputParms2.TargetString == nil {

		return true
	}

	if targetInputParms1.TargetString == nil &&
		targetInputParms2.TargetString != nil {

		return false
	}

	if targetInputParms1.TargetString != nil &&
		targetInputParms2.TargetString == nil {

		return false
	}

	return targetInputParms1.TargetString.Equal(
		targetInputParms2.TargetString)
}

// ptr - Returns a pointer to a new instance of
// charSearchTargetInputParametersDtoElectron.
//
func (searchTargetInputParmsElectron charSearchTargetInputParametersDtoElectron) ptr() *charSearchTargetInputParametersDtoElectron {

	if searchTargetInputParmsElectron.lock == nil {
		searchTargetInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTargetInputParmsElectron.lock.Lock()

	defer searchTargetInputParmsElectron.lock.Unlock()

	return &charSearchTargetInputParametersDtoElectron{
		lock: new(sync.Mutex),
	}
}

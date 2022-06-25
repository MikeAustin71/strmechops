package strmech

import "sync"

// charSearchTestInputParametersDtoElectron - Provides helper methods for type
// CharSearchTestInputParametersDto.
//
type charSearchTestInputParametersDtoElectron struct {
	lock *sync.Mutex
}

// emptyTestStrings - Receives a pointer to an instance of
// CharSearchTestInputParametersDto and proceeds to reset the
// internal member variable 'TestString' to its zero or
// uninitialized state.
//
// Since member variable 'TestString' is a pointer to an instance
// of RuneArrayDto, the 'TestString' pointer is reset to 'nil'.
// This step is taken in order to preserve data sharing in case
// another object also holds a pointer to that same instance of
// RuneArrayDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms           *CharSearchTestInputParametersDto
//     - Pointer to an instance of
//       CharSearchTestInputParametersDto. The internal member
//       variable, 'TestString', will be reset to a value of
//       'nil'.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchTestInputParmsElectron *charSearchTestInputParametersDtoElectron) emptyTestStrings(
	testInputParms *CharSearchTestInputParametersDto) {

	if searchTestInputParmsElectron.lock == nil {
		searchTestInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTestInputParmsElectron.lock.Lock()

	defer searchTestInputParmsElectron.lock.Unlock()

	if testInputParms == nil {
		return
	}

	testInputParms.TestString = nil

	return
}

// equalTestStrings - Compares the Test Strings from two
// instances of CharSearchTestInputParametersDto to determine if
// they are equivalent.
//
// If both TestString member variables are 'nil' pointers, this
// method classifies them as equivalent.
//
// If both TestString member variables are equal in all respects,
// this method returns a boolean value of 'true'. Otherwise, a
// value of 'false' is returned to the calling routine.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  testInputParms1            *CharSearchTestInputParametersDto
//     - An instance of CharSearchTestInputParametersDto. If the
//       internal member variable, 'TestString', will be compared
//       to the same internal member variable ('TestString') in
//       parameter 'testInputParms2' to determine if the two
//       test strings are equivalent.
//
//
//  testInputParms2            *CharSearchTestInputParametersDto
//     - An instance of CharSearchTestInputParametersDto. If the
//       internal member variable, 'TestString', will be compared
//       to the same internal member variable ('TestString') in
//       parameter 'testInputParms1' to determine if the two
//       test strings are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of test strings in parameters
//       'testInputParms1' and 'testInputParms2' show that the
//       internal member variables 'TestString' are equivalent,
//       this method will return a boolean value of 'true'.
//
//       If the two test strings are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchTestInputParmsElectron *charSearchTestInputParametersDtoElectron) equalTestStrings(
	testInputParms1 *CharSearchTestInputParametersDto,
	testInputParms2 *CharSearchTestInputParametersDto) bool {

	if searchTestInputParmsElectron.lock == nil {
		searchTestInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTestInputParmsElectron.lock.Lock()

	defer searchTestInputParmsElectron.lock.Unlock()

	if testInputParms1 == nil ||
		testInputParms2 == nil {
		return false
	}

	if testInputParms1.TestString == nil &&
		testInputParms2.TestString == nil {

		return true
	}

	if testInputParms1.TestString == nil &&
		testInputParms2.TestString != nil {

		return false
	}

	if testInputParms1.TestString != nil &&
		testInputParms2.TestString == nil {

		return false
	}

	return testInputParms1.TestString.Equal(
		testInputParms2.TestString)
}

// ptr - Returns a pointer to a new instance of
// charSearchTestInputParametersDtoElectron.
//
func (searchTestInputParmsElectron charSearchTestInputParametersDtoElectron) ptr() *charSearchTestInputParametersDtoElectron {

	if searchTestInputParmsElectron.lock == nil {
		searchTestInputParmsElectron.lock = new(sync.Mutex)
	}

	searchTestInputParmsElectron.lock.Lock()

	defer searchTestInputParmsElectron.lock.Unlock()

	return &charSearchTestInputParametersDtoElectron{
		lock: new(sync.Mutex),
	}
}

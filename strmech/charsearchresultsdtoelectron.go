package strmech

import "sync"

// charSearchResultsDtoElectron - Provides helper methods for type
// CharSearchResultsDto.
//
type charSearchResultsDtoElectron struct {
	lock *sync.Mutex
}

// emptyRemainderStrings - Receives a pointer to an instance of
// CharSearchResultsDto and proceeds to reset the internal member
// variable 'RemainderString' to its zero or uninitialized state.
//
// Since member variable 'RemainderString' is a pointer to an
// instance of RuneArrayDto, the 'RemainderString' pointer is
// reset to 'nil'. This step is taken in order to preserve data
// sharing in case another object also holds a pointer to that same
// instance of RuneArrayDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto           *CharSearchResultsDto
//     - Pointer to an instance of CharSearchResultsDto. The
//       internal member variable, 'RemainderString', will be
//       reset to a value of 'nil'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchResultsDtoElectron charSearchResultsDtoElectron) emptyRemainderStrings(
	searchResultsDto *CharSearchResultsDto) {

	if searchResultsDtoElectron.lock == nil {
		searchResultsDtoElectron.lock = new(sync.Mutex)
	}

	searchResultsDtoElectron.lock.Lock()

	defer searchResultsDtoElectron.lock.Unlock()

	if searchResultsDto == nil {
		return
	}

	searchResultsDto.RemainderString.Empty()

	return
}

// emptyReplacementStrings - Receives a pointer to an instance of
// CharSearchResultsDto and proceeds to reset the internal member
// variable 'ReplacementString' to its zero or uninitialized state.
//
// Since member variable 'ReplacementString' is a pointer to an
// instance of RuneArrayDto, the 'ReplacementString' pointer is
// reset to 'nil'. This step is taken in order to preserve data
// sharing in case another object also holds a pointer to that same
// instance of RuneArrayDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto           *CharSearchResultsDto
//     - Pointer to an instance of CharSearchResultsDto. The
//       internal member variable, 'ReplacementString', will be
//       reset to a value of 'nil'.
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchResultsDtoElectron charSearchResultsDtoElectron) emptyReplacementStrings(
	searchResultsDto *CharSearchResultsDto) {

	if searchResultsDtoElectron.lock == nil {
		searchResultsDtoElectron.lock = new(sync.Mutex)
	}

	searchResultsDtoElectron.lock.Lock()

	defer searchResultsDtoElectron.lock.Unlock()

	if searchResultsDto == nil {
		return
	}

	searchResultsDto.ReplacementString.Empty()

	return
}

// equalRemainderStrings - Compares the Remainder Strings from
// two instances of CharSearchResultsDto to determine if they are
// equivalent.
//
// If both 'RemainderString' member variables are 'nil' pointers,
// this method classifies them as equivalent.
//
// If both 'RemainderString' member variables are equal in all
// respects, this method returns a boolean value of 'true'.
// Otherwise, a value of 'false' is returned to the calling
// routine.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto1          *CharSearchResultsDto
//     - An instance of CharSearchResultsDto. If the
//       internal member variable, 'RemainderString', will be
//       compared to the same internal member variable
//       ('RemainderString') in parameter 'searchResultsDto2' to
//       determine if the two Remainder Strings are equivalent.
//
//
//  searchResultsDto2          *CharSearchResultsDto
//     - An instance of CharSearchResultsDto. If the internal
//       member variable, 'RemainderString', will be compared to
//       the same internal member variable ('RemainderString') in
//       parameter 'searchResultsDto1' to determine if the two
//       Remainder Strings are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of target strings in parameters
//       'searchResultsDto1' and 'searchResultsDto2' show that the
//       internal member variables 'RemainderString' are
//       equivalent, this method will return a boolean value of
//       'true'.
//
//       If the two target strings are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchResultsDtoElectron charSearchResultsDtoElectron) equalRemainderStrings(
	searchResultsDto1 *CharSearchResultsDto,
	searchResultsDto2 *CharSearchResultsDto) bool {

	if searchResultsDtoElectron.lock == nil {
		searchResultsDtoElectron.lock = new(sync.Mutex)
	}

	searchResultsDtoElectron.lock.Lock()

	defer searchResultsDtoElectron.lock.Unlock()

	if searchResultsDto1 == nil ||
		searchResultsDto2 == nil {
		return false
	}

	return searchResultsDto1.RemainderString.Equal(
		&searchResultsDto2.RemainderString)

}

// equalReplacementStrings - Compares the Replacement Strings from
// two instances of CharSearchResultsDto to determine if they are
// equivalent.
//
// If both 'ReplacementString' member variables are 'nil' pointers,
// this method classifies them as equivalent.
//
// If both 'ReplacementString' member variables are equal in all
// respects, this method returns a boolean value of 'true'.
// Otherwise, a value of 'false' is returned to the calling
// routine.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto1          *CharSearchResultsDto
//     - An instance of CharSearchResultsDto. If the
//       internal member variable, 'ReplacementString', will be
//       compared to the same internal member variable
//       ('ReplacementString') in parameter 'searchResultsDto2' to
//       determine if the two Replacement Strings are equivalent.
//
//
//  searchResultsDto2          *CharSearchResultsDto
//     - An instance of CharSearchResultsDto. If the internal
//       member variable, 'ReplacementString', will be compared to
//       the same internal member variable ('ReplacementString') in
//       parameter 'searchResultsDto1' to determine if the two
//       Replacement Strings are equivalent.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  bool
//     - If the comparison of target strings in parameters
//       'searchResultsDto1' and 'searchResultsDto2' show that the
//       internal member variables 'ReplacementString' are
//       equivalent, this method will return a boolean value of
//       'true'.
//
//       If the two target strings are NOT equal, this method will
//       return a boolean value of 'false' to the calling function.
//
func (searchResultsDtoElectron charSearchResultsDtoElectron) equalReplacementStrings(
	searchResultsDto1 *CharSearchResultsDto,
	searchResultsDto2 *CharSearchResultsDto) bool {

	if searchResultsDtoElectron.lock == nil {
		searchResultsDtoElectron.lock = new(sync.Mutex)
	}

	searchResultsDtoElectron.lock.Lock()

	defer searchResultsDtoElectron.lock.Unlock()

	if searchResultsDto1 == nil ||
		searchResultsDto2 == nil {
		return false
	}

	return searchResultsDto1.ReplacementString.Equal(
		&searchResultsDto2.ReplacementString)

}

// ptr - Returns a pointer to a new instance of
// charSearchResultsDtoElectron.
//
func (searchResultsDtoElectron charSearchResultsDtoElectron) ptr() *charSearchResultsDtoElectron {

	if searchResultsDtoElectron.lock == nil {
		searchResultsDtoElectron.lock = new(sync.Mutex)
	}

	searchResultsDtoElectron.lock.Lock()

	defer searchResultsDtoElectron.lock.Unlock()

	return &charSearchResultsDtoElectron{
		lock: new(sync.Mutex),
	}
}

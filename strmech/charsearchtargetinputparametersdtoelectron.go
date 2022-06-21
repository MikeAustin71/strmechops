package strmech

import "sync"

// charSearchTargetInputParametersDtoElectron - Provides helper methods for type
// CharSearchTargetInputParametersDto.
//
type charSearchTargetInputParametersDtoElectron struct {
	lock *sync.Mutex
}

// empty - Receives a pointer to an instance of
// CharSearchTargetInputParametersDto and proceeds to reset the
// data values for member variables to their initial or zero
// values.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All the member variable data values contained in input parameter
// 'targetInputParms' will be deleted and reset to their zero
// values.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms           *CharSearchTargetInputParametersDto
//     - A pointer to an instance of
//       CharSearchTargetInputParametersDto. All the internal
//       member variables contained in this instance will be
//       deleted and reset to their zero values.
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (searchTargetInputParmsElectron *charSearchTargetInputParametersDtoElectron) empty(
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

	targetInputParms.TargetStringName = ""

	targetInputParms.TargetStringLength = -1

	targetInputParms.TargetStringLengthName = ""

	targetInputParms.TargetStringStartingSearchIndex = -1

	targetInputParms.TargetStringStartingSearchIndexName = ""

	targetInputParms.TargetStringSearchLength = -2

	targetInputParms.TargetStringSearchLengthName = ""

	targetInputParms.TargetStringAdjustedSearchLength = -1

	targetInputParms.TargetStringDescription1 = ""

	targetInputParms.TargetStringDescription2 = ""

	targetInputParms.FoundFirstNumericDigitInNumStr = false

	targetInputParms.TextCharSearchType = CharSearchType.None()

	return
}

// equal - Receives a pointer to two instances of
// charSearchTargetInputParametersDtoElectron and proceeds to
// compare their member variables in order to determine if they are
// equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
//
func (searchTargetInputParmsElectron *charSearchTargetInputParametersDtoElectron) equal(
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
		targetInputParms2.TargetString != nil {

		return false
	}

	if targetInputParms1.TargetString != nil &&
		targetInputParms2.TargetString == nil {

		return false
	}

	if targetInputParms1.TargetString != nil &&
		targetInputParms2.TargetString != nil {

		if !targetInputParms1.TargetString.Equal(
			targetInputParms2.TargetString) {
			return false
		}
	}

	// Target Strings are equal. However, both
	// pointers may be nil

	if targetInputParms1.TargetStringName !=
		targetInputParms2.TargetStringName {

		return false
	}

	if targetInputParms1.TargetStringLength !=
		targetInputParms2.TargetStringLength {

		return false
	}

	if targetInputParms1.TargetStringLengthName !=
		targetInputParms2.TargetStringLengthName {

		return false
	}

	if targetInputParms1.TargetStringStartingSearchIndex !=
		targetInputParms2.TargetStringStartingSearchIndex {

		return false
	}

	if targetInputParms1.TargetStringStartingSearchIndexName !=
		targetInputParms2.TargetStringStartingSearchIndexName {

		return false
	}

	if targetInputParms1.TargetStringSearchLength !=
		targetInputParms2.TargetStringSearchLength {

		return false
	}

	if targetInputParms1.TargetStringSearchLengthName !=
		targetInputParms2.TargetStringSearchLengthName {

		return false
	}

	if targetInputParms1.TargetStringAdjustedSearchLength !=
		targetInputParms2.TargetStringAdjustedSearchLength {

		return false
	}

	if targetInputParms1.TargetStringDescription1 !=
		targetInputParms2.TargetStringDescription1 {

		return false
	}

	if targetInputParms1.TargetStringDescription2 !=
		targetInputParms2.TargetStringDescription2 {

		return false
	}

	if targetInputParms1.FoundFirstNumericDigitInNumStr !=
		targetInputParms2.FoundFirstNumericDigitInNumStr {

		return false
	}

	if targetInputParms1.TextCharSearchType !=
		targetInputParms2.TextCharSearchType {

		return false
	}

	return true
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelElectron.
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

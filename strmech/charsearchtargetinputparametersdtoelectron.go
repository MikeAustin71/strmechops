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

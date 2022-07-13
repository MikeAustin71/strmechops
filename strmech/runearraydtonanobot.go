package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// runeArrayDtoNanobot - Provides helper methods for type
// RuneArrayDto.
//
type runeArrayDtoNanobot struct {
	lock *sync.Mutex
}

func (runeDtoNanobot *runeArrayDtoNanobot) characterSearchExecutor(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchRuneArrayResultsDto,
	error) {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	searchResults := CharSearchRuneArrayResultsDto{}

	searchResults.Empty()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoNanobot."+
			"characterSearchExecutor()",
		"")

	if err != nil {

		return searchResults, err

	}

	err = testInputParms.ValidateTestParameters(
		ePrefix)

	if err != nil {

		return searchResults, err

	}

	err = targetInputParms.ValidateTargetParameters(
		ePrefix)

	if err != nil {

		return searchResults, err

	}

	runeDtoElectron := runeArrayDtoElectron{}

	if testInputParms.TextCharSearchType == CharSearchType.LinearEndOfString() {

		return runeDtoElectron.linearEndOfStringSearch(
			targetInputParms,
			testInputParms,
			ePrefix)

	}

	if testInputParms.TextCharSearchType == CharSearchType.LinearTargetStartingIndex() {

		return runeDtoElectron.linearTargetStartingIndexSearch(
			targetInputParms,
			testInputParms,
			ePrefix)
	}

	if testInputParms.TextCharSearchType == CharSearchType.SingleTargetChar() {

		return runeDtoElectron.singleCharacterSearch(
			targetInputParms,
			testInputParms,
			ePrefix)

	}

	return searchResults, err
}

// ptr - Returns a pointer to a new instance of
// runeArrayDtoNanobot.
//
func (runeDtoNanobot runeArrayDtoNanobot) ptr() *runeArrayDtoNanobot {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	return &runeArrayDtoNanobot{
		lock: new(sync.Mutex),
	}
}

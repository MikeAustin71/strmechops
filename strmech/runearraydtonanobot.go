package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoNanobot struct {
	lock *sync.Mutex
}

func (runeDtoNanobot *runeArrayDtoNanobot) characterSearchExecutor(
	targetInputParms CharSearchTargetInputParametersDto,
	testInputParms CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (CharSearchResultsDto, error) {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	searchResults := CharSearchResultsDto{}

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

	if searchResults.CharSearchType == CharSearchType.LinearEndOfString() {

		return runeDtoElectron.linearEndOfStringSearch(
			targetInputParms,
			testInputParms,
			ePrefix)

	}

	if searchResults.CharSearchType == CharSearchType.LinearTargetStartingIndex() {

		return runeDtoElectron.linearTargetStartingIndexSearch(
			targetInputParms,
			testInputParms,
			ePrefix)
	}

	if searchResults.CharSearchType == CharSearchType.SingleTargetChar() {

		return searchResults, err
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

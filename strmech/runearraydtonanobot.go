package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoNanobot struct {
	lock *sync.Mutex
}

func (runeDtoNanobot *runeArrayDtoNanobot) characterSearchExecutor(
	testSearchString *RuneArrayDto,
	testSearchStringName string,
	targetSearchString *RuneArrayDto,
	targetSearchStringName string,
	targetStartingSearchIndex int,
	targetStartingSearchIndexName string,
	targetSearchLength int,
	targetSearchLengthName string,
	searchType CharacterSearchType,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchResultsDto,
	error) {

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

	if len(testSearchStringName) == 0 {
		testSearchStringName = "testSearchString"
	}

	if testSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			testSearchStringName)

		return searchResults, err
	}

	if !searchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charSearchType' is invalid!\n"+
			"'charSearchType' must be set to one of three enumerations:\n"+
			"CharacterSearchType(0).LinearTargetStartingIndex()\n OR"+
			"CharacterSearchType(0).SingleTargetChar()\n"+
			"CharacterSearchType(0).LinearEndOfString()\n"+
			"'charSearchType' string  value = '%v'\n"+
			"'charSearchType' integer value = '%v'\n",
			ePrefix.String(),
			searchType.String(),
			searchType.XValueInt())

		return searchResults, err
	}

	searchResults.SearchType = searchType

	searchResults.TestStrDescription1 =
		testSearchString.Description1

	searchResults.TestStrDescription2 =
		testSearchString.Description2

	searchResults.TestStrLength =
		len(testSearchString.CharsArray)

	if searchResults.TestStrLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			testSearchStringName,
			testSearchStringName,
			testSearchStringName)

		return searchResults, err
	}

	if len(targetSearchStringName) == 0 {
		targetSearchStringName = "targetSearchString"
	}

	if targetSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			targetSearchStringName)

		return searchResults, err
	}

	searchResults.TargetStringDescription1 =
		targetSearchString.Description1

	searchResults.TargetStringDescription2 =
		targetSearchString.Description2

	searchResults.TargetStringLength =
		len(targetSearchString.CharsArray)

	if searchResults.TargetStringLength == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			targetSearchLengthName,
			targetSearchStringName,
			targetSearchStringName)

		searchResults.Empty()

		return searchResults, err
	}

	if len(targetStartingSearchIndexName) == 0 {
		targetStartingSearchIndexName = "targetStartingSearchIndex"
	}

	if targetStartingSearchIndex < 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter %v is invalid!\n"+
			"%v is less than zero (0)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetStartingSearchIndexName,
			targetStartingSearchIndexName,
			targetStartingSearchIndexName,
			targetStartingSearchIndex)

		return searchResults, err
	}

	if targetStartingSearchIndex >= searchResults.TargetStringLength {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value greater than the last\n"+
			"index in '%v.CharsArray'.\n"+
			"Last Index in %v.CharsArray = '%v'\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetStartingSearchIndexName,
			targetStartingSearchIndexName,
			targetSearchStringName,
			targetSearchStringName,
			searchResults.TargetStringSearchLength-1,
			targetStartingSearchIndexName,
			targetStartingSearchIndex)

		return searchResults, err
	}

	searchResults.TargetStringStartingSearchIndex =
		targetStartingSearchIndex

	if len(targetSearchLengthName) == 0 {
		targetSearchLengthName = "targetSearchLength"
	}

	if targetSearchLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value less than minus one (-1)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetSearchLengthName,
			targetSearchLengthName,
			targetSearchStringName,
			targetSearchLength)

		return searchResults, err
	}

	if targetSearchLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value of Zero (0)\n",
			ePrefix.String(),
			targetSearchLengthName,
			targetSearchLengthName)

		return searchResults, err
	}

	if targetSearchLength == -1 {

		targetSearchLength = searchResults.TargetStringLength
	}

	searchResults.TargetStringSearchLength = targetSearchLength

	searchResults.TargetStringSearchLength =
		searchResults.TargetStringStartingSearchIndex +
			searchResults.TargetStringSearchLength

	if searchResults.TargetStringSearchLength > searchResults.TargetStringLength {

		searchResults.TargetStringSearchLength =
			searchResults.TargetStringLength

	}

	runeDtoElectron := runeArrayDtoElectron{}

	if searchResults.SearchType == CharSearchType.LinearEndOfString() {

		return runeDtoElectron.linearEndOfStringSearch(
			testSearchString,
			testSearchStringName,
			targetSearchString,
			targetSearchStringName,
			searchResults,
			ePrefix)

	}

	if searchResults.SearchType == CharSearchType.LinearTargetStartingIndex() {

		return runeDtoElectron.linearTargetStartingIndexSearch(
			testSearchString,
			testSearchStringName,
			targetSearchString,
			targetSearchStringName,
			searchResults,
			ePrefix)
	}

	if searchResults.SearchType == CharSearchType.SingleTargetChar() {

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

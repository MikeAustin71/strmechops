package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoNanobot struct {
	lock *sync.Mutex
}

func (runeDtoNanobot *runeArrayDtoNanobot) linearTargetStartingIndexSearch(
	testSearchString *RuneArrayDto,
	testSearchStringName string,
	targetSearchString *RuneArrayDto,
	targetSearchStringName string,
	targetStartingSearchIndex int,
	targetStartingSearchIndexName string,
	targetSearchLength int,
	targetSearchLengthName string,
	errPrefDto *ePref.ErrPrefixDto) (
	foundRuneArrayDtoChars bool,
	lastTargetSearchIndex int,
	lastTestStingIndex int,
	searchType CharacterSearchType,
	err error) {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundRuneArrayDtoChars = false
	lastTargetSearchIndex = targetStartingSearchIndex
	lastTestStingIndex = -1
	searchType = CharSearchType.LinearTargetStartingIndex()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoNanobot."+
			"linearTargetStartingIndexSearch()",
		"")

	if err != nil {

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if len(testSearchStringName) == 0 {
		testSearchStringName = "testSearchString"
	}

	if testSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			testSearchStringName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	lenTestSearchString := len(testSearchString.CharsArray)

	if lenTestSearchString == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			testSearchStringName,
			testSearchStringName,
			testSearchStringName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if len(targetSearchStringName) == 0 {
		targetSearchStringName = "targetSearchString"
	}

	if targetSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			targetSearchStringName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if len(targetStartingSearchIndexName) == 0 {
		targetStartingSearchIndexName = "targetStartingSearchIndex"
	}

	if len(targetSearchLengthName) == 0 {
		targetStartingSearchIndexName = "targetSearchLength"
	}

	actualLenTargetSearchString := len(targetSearchString.CharsArray)

	if actualLenTargetSearchString == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			targetSearchStringName,
			targetSearchStringName,
			targetSearchStringName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

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

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if targetStartingSearchIndex >= actualLenTargetSearchString {

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
			actualLenTargetSearchString-1,
			targetStartingSearchIndexName,
			targetStartingSearchIndex)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

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
			targetStartingSearchIndex)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if targetSearchLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value of Zero (0)\n",
			ePrefix.String(),
			targetSearchLengthName,
			targetSearchLengthName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if targetSearchLength == -1 {

		targetSearchLength = actualLenTargetSearchString
	}

	adjustedCharSearchLength :=
		targetStartingSearchIndex + targetSearchLength

	if adjustedCharSearchLength > actualLenTargetSearchString {
		adjustedCharSearchLength = actualLenTargetSearchString
	}

	j := 0

	for i := targetStartingSearchIndex; i < adjustedCharSearchLength; i++ {

		if testSearchString.CharsArray[j] !=
			targetSearchString.CharsArray[i] {

			// Search Failed. No Match!
			// Exit Here!
			return foundRuneArrayDtoChars,
				lastTargetSearchIndex,
				lastTestStingIndex,
				searchType,
				err

		}

		// We found a matching char
		j++

		if j == lenTestSearchString {
			// Search Was SUCCESSFUL!
			// All characters found!
			// EXIT HERE!

			foundRuneArrayDtoChars = true
			lastTargetSearchIndex = i
			lastTestStingIndex = j - 1
			searchType = CharSearchType.LinearTargetStartingIndex()

			return foundRuneArrayDtoChars,
				lastTargetSearchIndex,
				lastTestStingIndex,
				searchType,
				err

		}

	}

	return foundRuneArrayDtoChars,
		lastTargetSearchIndex,
		lastTestStingIndex,
		searchType,
		err
}

func (runeDtoNanobot *runeArrayDtoNanobot) linearEndOfStringSearch(
	testSearchString *RuneArrayDto,
	testSearchStringName string,
	targetSearchString *RuneArrayDto,
	targetSearchStringName string,
	targetStartingSearchIndex int,
	targetStartingSearchIndexName string,
	targetSearchLength int,
	targetSearchLengthName string,
	errPrefDto *ePref.ErrPrefixDto) (
	searchResults CharSearchResultsDto,
	err error) {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	searchResults.Empty()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoNanobot."+
			"linearEndOfStringSearch()",
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

	searchResults.SearchType =
		CharSearchType.LinearEndOfString()

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

	if len(targetStartingSearchIndexName) == 0 {
		targetStartingSearchIndexName = "targetStartingSearchIndex"
	}

	if len(targetSearchLengthName) == 0 {
		targetStartingSearchIndexName = "targetSearchLength"
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
			targetSearchStringName,
			targetSearchStringName,
			targetSearchStringName)

		return searchResults, err
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

	if targetSearchLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value less than minus one (-1)\n"+
			"%v = '%v'\n",
			ePrefix.String(),
			targetSearchLengthName,
			targetSearchLengthName,
			targetSearchStringName,
			targetStartingSearchIndex)

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

	j := 0

	searchResults.TestStrStartingIndex = j

	for i := searchResults.TargetStringStartingSearchIndex; i < searchResults.TargetStringSearchLength; i++ {

		if testSearchString.CharsArray[j] !=
			targetSearchString.CharsArray[i] {

			// Search Failed. No Match!
			// Exit Here!
			return searchResults, err
		}

		if searchResults.TargetStringFirstFoundIndex < 0 {
			searchResults.TargetStringFirstFoundIndex = i
		}

		// We found a matching char
		if searchResults.TestStringFirstFoundIndex < 0 {

			searchResults.TestStringFirstFoundIndex = j
		}

		j++

		if j == searchResults.TestStrLength {
			// Search Was SUCCESSFUL!
			// All characters found!
			// EXIT HERE!

			searchResults.FoundSearchTarget = true
			searchResults.TargetStringLastFoundIndex = i
			searchResults.TestStrLastFoundIndex = j - 1

			return searchResults, err
		}

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

func (runeDtoNanobot *runeArrayDtoNanobot) singleCharacterSearch(
	testSearchString *RuneArrayDto,
	testSearchStringName string,
	targetSearchString *RuneArrayDto,
	targetSearchStringName string,
	targetStartingSearchIndex int,
	targetStartingSearchIndexName string,
	targetSearchLength int,
	targetSearchLengthName string,
	errPrefDto *ePref.ErrPrefixDto) (
	foundRuneArrayDtoChars bool,
	lastTargetSearchIndex int,
	lastTestStingIndex int,
	searchType CharacterSearchType,
	err error) {

	if runeDtoNanobot.lock == nil {
		runeDtoNanobot.lock = new(sync.Mutex)
	}

	runeDtoNanobot.lock.Lock()

	defer runeDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	foundRuneArrayDtoChars = false
	lastTargetSearchIndex = targetStartingSearchIndex
	lastTestStingIndex = -1
	searchType = CharSearchType.SingleTargetChar()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoNanobot."+
			"singleCharacterSearch()",
		"")

	if err != nil {

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if len(testSearchStringName) == 0 {
		testSearchStringName = "testSearchString"
	}

	if testSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			testSearchStringName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	lenTestSearchString := len(testSearchString.CharsArray)

	if lenTestSearchString == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			testSearchStringName,
			testSearchStringName,
			testSearchStringName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if len(targetSearchStringName) == 0 {
		targetSearchStringName = "targetSearchString"
	}

	if targetSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			targetSearchStringName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if len(targetStartingSearchIndexName) == 0 {
		targetStartingSearchIndexName = "targetStartingSearchIndex"
	}

	if len(targetSearchLengthName) == 0 {
		targetStartingSearchIndexName = "targetSearchLength"
	}

	actualLenTargetSearchString := len(targetSearchString.CharsArray)

	if actualLenTargetSearchString == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is invalid!\n"+
			"The rune array encapsulated by '%v' is empty\n"+
			"Length of %v.CharsArray is Zero (0).\n",
			ePrefix.String(),
			targetSearchStringName,
			targetSearchStringName,
			targetSearchStringName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

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

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if targetStartingSearchIndex >= actualLenTargetSearchString {

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
			actualLenTargetSearchString-1,
			targetStartingSearchIndexName,
			targetStartingSearchIndex)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

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
			targetStartingSearchIndex)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if targetSearchLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter %v is invalid!\n"+
			"%v has a value of Zero (0)\n",
			ePrefix.String(),
			targetSearchLengthName,
			targetSearchLengthName)

		return foundRuneArrayDtoChars,
			lastTargetSearchIndex,
			lastTestStingIndex,
			searchType,
			err

	}

	if targetSearchLength == -1 {

		targetSearchLength = actualLenTargetSearchString
	}

	adjustedCharSearchLength :=
		targetStartingSearchIndex + targetSearchLength

	if adjustedCharSearchLength > actualLenTargetSearchString {
		adjustedCharSearchLength = actualLenTargetSearchString
	}

	for i := targetStartingSearchIndex; i < adjustedCharSearchLength; i++ {

		for j := 0; j < lenTestSearchString; j++ {

			if testSearchString.CharsArray[j] ==
				targetSearchString.CharsArray[i] {

				// Search SUCCESSFUL! SINGLE CHARACTER MATCH!
				// Exit Here!

				foundRuneArrayDtoChars = true
				lastTargetSearchIndex = i
				lastTestStingIndex = j
				searchType = CharSearchType.SingleTargetChar()

				return foundRuneArrayDtoChars,
					lastTargetSearchIndex,
					lastTestStingIndex,
					searchType,
					err

			}

		}

	}

	// Search FAILED!
	// EXIT HERE!

	return foundRuneArrayDtoChars,
		lastTargetSearchIndex,
		lastTestStingIndex,
		searchType,
		err
}

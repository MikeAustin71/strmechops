package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoElectron struct {
	lock *sync.Mutex
}

// linearEndOfStringSearch - This low level function executes a
// string search operation proceeds through the entire length of
// Target Search String.
//
// As a low level function, very little validation is performed on
// the input parameters for this function. It is assumed that
// input parameter validation was previously performed by a higher
// level function.
//
func (runeDtoElectron *runeArrayDtoElectron) linearEndOfStringSearch(
	testSearchString *RuneArrayDto,
	testSearchStringName string,
	targetSearchString *RuneArrayDto,
	targetSearchStringName string,
	searchResults CharSearchResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoElectron."+
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

	if len(targetSearchStringName) == 0 {
		targetSearchStringName = "targetSearchStringName"
	}

	if targetSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			targetSearchStringName)

		return searchResults, err
	}

	j := 0
	k := 0

	searchResults.TestStrStartingIndex = 0

	for i := searchResults.TargetStringStartingSearchIndex; i < searchResults.TargetStringSearchLength; i++ {

		k = i
		j = 0

		for true {

			if testSearchString.CharsArray[j] !=
				targetSearchString.CharsArray[k] {

				break
			}

			j++

			if j == searchResults.TestStrLength {

				// Search Was SUCCESSFUL!
				// All Test characters found!
				// EXIT HERE!

				searchResults.FoundSearchTarget = true

				searchResults.TestStrStartingIndex = 0

				searchResults.TestStringFirstFoundIndex = 0

				searchResults.TestStrLastFoundIndex =
					searchResults.TestStrLength - 1

				searchResults.TargetStringLastFoundIndex = k

				searchResults.TargetStringFirstFoundIndex =
					searchResults.TargetStringLastFoundIndex -
						searchResults.TestStrLength +
						1

				return searchResults, err
			}

			k++

			if k == searchResults.TargetStringSearchLength {
				break
			}
		}

	}

	return searchResults, err
}

// linearTargetStartingIndexSearch
func (runeDtoElectron *runeArrayDtoElectron) linearTargetStartingIndexSearch(
	testSearchString *RuneArrayDto,
	testSearchStringName string,
	targetSearchString *RuneArrayDto,
	targetSearchStringName string,
	searchResults CharSearchResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoElectron."+
			"linearTargetStartingIndexSearch()",
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

	if len(targetSearchStringName) == 0 {
		targetSearchStringName = "targetSearchStringName"
	}

	if targetSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			targetSearchStringName)

		return searchResults, err
	}

	j := 0

	for i := searchResults.TargetStringStartingSearchIndex; i < searchResults.TargetStringSearchLength; i++ {

		if testSearchString.CharsArray[j] !=
			targetSearchString.CharsArray[i] {

			searchResults.FoundSearchTarget = false

			// Search Failed. No Match!
			// Exit Here!
			return searchResults, err
		}

		j++

		if j == searchResults.TestStrLength {

			// Search Was SUCCESSFUL!
			// All Test Characters found!
			// EXIT HERE!

			searchResults.FoundSearchTarget = true

			searchResults.TestStrStartingIndex = 0

			searchResults.TestStringFirstFoundIndex = 0

			searchResults.TestStrLastFoundIndex =
				searchResults.TestStrLength - 1

			searchResults.TargetStringLastFoundIndex = i

			searchResults.TargetStringFirstFoundIndex =
				searchResults.TargetStringLastFoundIndex -
					searchResults.TestStrLength +
					1

			return searchResults, err
		}

	}

	return searchResults, err
}

// ptr - Returns a pointer to a new instance of
// runeArrayDtoElectron.
//
func (runeDtoElectron runeArrayDtoElectron) ptr() *runeArrayDtoElectron {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	return &runeArrayDtoElectron{
		lock: new(sync.Mutex),
	}
}

func (runeDtoElectron *runeArrayDtoElectron) singleCharacterSearch(
	testSearchString *RuneArrayDto,
	testSearchStringName string,
	targetSearchString *RuneArrayDto,
	targetSearchStringName string,
	searchResults CharSearchResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	CharSearchResultsDto,
	error) {

	if runeDtoElectron.lock == nil {
		runeDtoElectron.lock = new(sync.Mutex)
	}

	runeDtoElectron.lock.Lock()

	defer runeDtoElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoElectron."+
			"linearTargetStartingIndexSearch()",
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

	if len(targetSearchStringName) == 0 {
		targetSearchStringName = "targetSearchStringName"
	}

	if targetSearchString == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter '%v' is a nil pointer!\n",
			ePrefix.String(),
			targetSearchStringName)

		return searchResults, err
	}

	targetChar :=
		targetSearchString.CharsArray[searchResults.TargetStringStartingSearchIndex]

	for j := 0; j < searchResults.TestStrLength; j++ {

		if testSearchString.CharsArray[j] == targetChar {
			// Search SUCCESSFUL! SINGLE CHARACTER MATCH!
			// Exit Here!

			searchResults.FoundSearchTarget = true

			searchResults.TestStrStartingIndex = 0

			searchResults.TestStringFirstFoundIndex = j

			searchResults.TestStrLastFoundIndex = j

			searchResults.TargetStringLastFoundIndex =
				searchResults.TargetStringStartingSearchIndex

			searchResults.TargetStringFirstFoundIndex =
				searchResults.TargetStringLastFoundIndex

			return searchResults, err
		}
	}

	return searchResults, err
}

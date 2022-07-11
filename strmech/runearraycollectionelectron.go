package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// runeArrayCollectionElectron - Provides helper methods for type
// RuneArrayCollection.
//
type runeArrayCollectionElectron struct {
	lock *sync.Mutex
}

// ptr - Returns a pointer to a new instance of
// runeArrayCollectionElectron.
//
func (runeArrayColElectron runeArrayCollectionElectron) ptr() *runeArrayCollectionElectron {

	if runeArrayColElectron.lock == nil {
		runeArrayColElectron.lock = new(sync.Mutex)
	}

	runeArrayColElectron.lock.Lock()

	defer runeArrayColElectron.lock.Unlock()

	return &runeArrayCollectionElectron{
		lock: new(sync.Mutex),
	}
}

// testValidityRuneArrayCollection - Tests the validity of a
// collection of RuneArrayDto objects.
//
// If the Rune Array collection has a length of zero, an error is
// returned.
//
// If any member of the collection has an empty characters array
// (Number of characters equal zero), an error is returned.
//
// If a member of the collection has an invalid Character Search
// Type, an error is returned.
//
func (runeArrayColElectron runeArrayCollectionElectron) testValidityRuneArrayCollection(
	arrayRuneArrayDtos []RuneArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if runeArrayColElectron.lock == nil {
		runeArrayColElectron.lock = new(sync.Mutex)
	}

	runeArrayColElectron.lock.Lock()

	defer runeArrayColElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayCollectionElectron."+
			"runeArrayCollectionElectron()",
		"")

	if err != nil {

		return isValid, err

	}

	lenArrayRuneArrayDtos := len(arrayRuneArrayDtos)

	if lenArrayRuneArrayDtos == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'arrayRuneArrayDtos' is \n"+
			"empty and has a length of zero!\n",
			ePrefix.String())

		return isValid, err

	}

	for i := 0; i < lenArrayRuneArrayDtos; i++ {

		if len(arrayRuneArrayDtos[i].CharsArray) == 0 {

			err = fmt.Errorf("%v\n"+
				"Error: RuneArrayDtoCollection[%v] has an empty\n"+
				"character array. It contains zero characters!\n",
				ePrefix.String(),
				i)

			return isValid, err
		}

		if !arrayRuneArrayDtos[i].charSearchType.XIsValid() {

			err = fmt.Errorf("%v\n"+
				"Error: RuneArrayDtoCollection[%v] has an invalid\n"+
				"Character Search Type ('charSearchType')!\n"+
				"'charSearchType' must be set to one of three enumerations:\n"+
				"(1) CharacterSearchType(0).LinearTargetStartingIndex()\n"+
				"(2) CharacterSearchType(0).SingleTargetChar()\n"+
				"(3) CharacterSearchType(0).LinearEndOfString()\n"+
				"'charSearchType' string  value = '%v'\n"+
				"'charSearchType' integer value = '%v'\n",
				ePrefix.String(),
				i,
				arrayRuneArrayDtos[i].charSearchType.String(),
				arrayRuneArrayDtos[i].charSearchType.XValueInt())

			return isValid, err

		}

	}

	isValid = true

	return isValid, err
}

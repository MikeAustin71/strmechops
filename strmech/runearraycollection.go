package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// RuneArrayCollection - A collection of Rune Array Dto objects.
//
// Esentially, this is a collection or an array of rune arrays.
//
type RuneArrayCollection struct {
	RuneArrayDtoCol []RuneArrayDto

	lock *sync.Mutex
}

// AddRuneArray - Receives an instance of RuneArrayDto and appends
// that instance to the RuneArrayDto Collection.
//
// This differs from method:
//  RuneArrayCollection.AddRuneArrayDeepCopy()
//
// This method appends the passed RuneArrayDto instance to the
// collection. The Deep Copy method appends a copy of the
// RuneArrayDto to the collection.
//
func (runeArrayCol *RuneArrayCollection) AddRuneArray(
	runeArrayDto RuneArrayDto,
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRuneArray()",
		"")

	if err != nil {
		return err
	}

	lenRuneDtoChars := len(runeArrayDto.CharsArray)

	if lenRuneDtoChars == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeArrayDto' is invalid!\n"+
			"runeArrayDto.CharsArray has a length of zero",
			ePrefix.String())

		return err

	}

	runeArrayCol.RuneArrayDtoCol =
		append(runeArrayCol.RuneArrayDtoCol, runeArrayDto)

	return err
}

// AddRuneArrayDeepCopy - Receives an instance of RuneArrayDto and
// appends a deep copy of that instance to the RuneArrayDto
// Collection.
//
// This differs from method:
//  RuneArrayCollection.AddRuneArray()
//
// This method appends a deep copy of the passed RuneArrayDto
// instance to the collection. The 'AddRuneArray()' method appends
// the actual RuneArrayDto instance to the collection.
//
func (runeArrayCol *RuneArrayCollection) AddRuneArrayDeepCopy(
	runeArrayDto RuneArrayDto,
	errorPrefix interface{}) (
	err error) {

	if runeArrayCol.lock == nil {
		runeArrayCol.lock = new(sync.Mutex)
	}

	runeArrayCol.lock.Lock()

	defer runeArrayCol.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayCollection."+
			"AddRuneArray()",
		"")

	if err != nil {
		return err
	}

	lenRuneDtoChars := len(runeArrayDto.CharsArray)

	if lenRuneDtoChars == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeArrayDto' is invalid!\n"+
			"runeArrayDto.CharsArray has a length of zero",
			ePrefix.String())

		return err

	}

	var deepCopyRuneArrayDto RuneArrayDto

	deepCopyRuneArrayDto,
		err = runeArrayDto.CopyOut(
		ePrefix.XCpy(
			"deepCopyRuneArrayDto<-runeArrayDto"))

	if err != nil {
		return err
	}

	runeArrayCol.RuneArrayDtoCol =
		append(runeArrayCol.RuneArrayDtoCol, deepCopyRuneArrayDto)

	return err
}

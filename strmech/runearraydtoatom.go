package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoAtom struct {
	lock *sync.Mutex
}

//	addRunes
//
//	Will add another rune array to the existing rune
//	array contained in the RuneArrayDto passed as an
//	input parameter.
//
//	The additional rune characters may either be added as
//	leading runes, added at the beginning of the existing
//	rune array, or as trailing runes added at the end of
//	the existing rune array.
//
//	The name of the rune array member variable which will
//	be modified by this method is:
//
//			RuneArrayDto.CharsArray
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	runeArrayDto				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		rune array contained in this RuneArrayDto
//		instance will be modified. Additional rune
//		characters will be added to this rune array:
//
//			runeArrayDto.CharsArray
//
//	charsToAdd					[]rune
//
//		The text characters contained in this rune
//		array will be added to the existing rune array
//		contained in input parameter, 'runeArrayDto'.
//
//		If this parameter is set to 'nil' or has a
//		length of zero, an error will be returned.
//
//	addTrailingChars			bool
//
//		This parameter determines whether 'charsToAdd'
//		will be added to 'runeArrayDto.CharsArray' as
//		trailing characters or as leading characters.
//
//		If 'addTrailingChars' is set to 'true',
//		'charsToAdd' will be added to the end of the
//		existing rune array as trailing characters.
//
//		If 'addTrailingChars' is set to 'false',
//		'charsToAdd' will be added to the beginning of
//		the existing rune array as leading characters.
func (runeDtoAtom *runeArrayDtoAtom) addRunes(
	runeArrayDto *RuneArrayDto,
	charsToAdd []rune,
	addTrailingChars bool,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if runeDtoAtom.lock == nil {
		runeDtoAtom.lock = new(sync.Mutex)
	}

	runeDtoAtom.lock.Lock()

	defer runeDtoAtom.lock.Unlock()
	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoNanobot."+
			"copyRuneArrayDto()",
		"")

	if err != nil {

		return err

	}

	if runeArrayDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeArrayDto' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if len(charsToAdd) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charsToAdd' is invalid!\n"+
			"'charsToAdd' is empty!\n",
			ePrefix.String())

		return err
	}

	if addTrailingChars == true {

		runeArrayDto.CharsArray = append(
			charsToAdd,
			runeArrayDto.CharsArray...)

	} else {

		runeArrayDto.CharsArray = append(
			runeArrayDto.CharsArray,
			charsToAdd...)

	}

	return err
}

// empty - Receives a pointer to an instance of
// RuneArrayDto and proceeds to reset the data values
// for member variables to their initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'runeArrayDto' will be deleted and reset to their zero
// values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto               *RuneArrayDto
//	   - A pointer to an instance of RuneArrayDto. All
//	     the internal member variables contained in this instance
//	     will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	NONE
func (runeDtoAtom *runeArrayDtoAtom) empty(
	runeArrayDto *RuneArrayDto) {

	if runeDtoAtom.lock == nil {
		runeDtoAtom.lock = new(sync.Mutex)
	}

	runeDtoAtom.lock.Lock()

	defer runeDtoAtom.lock.Unlock()

	if runeArrayDto == nil {
		return
	}

	new(runeArrayDtoElectron).
		emptyCharsArray(runeArrayDto)

	runeArrayDto.Description1 = ""

	runeArrayDto.Description2 = ""

	runeArrayDto.charSearchType = CharSearchType.LinearTargetStartingIndex()

	return
}

// equal - Receives a pointer to two instances of
// RuneArrayDto and proceeds to compare their member variables in
// order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
func (runeDtoAtom *runeArrayDtoAtom) equal(
	runeArrayDto1 *RuneArrayDto,
	runeArrayDto2 *RuneArrayDto) bool {

	if runeDtoAtom.lock == nil {
		runeDtoAtom.lock = new(sync.Mutex)
	}

	runeDtoAtom.lock.Lock()

	defer runeDtoAtom.lock.Unlock()

	if runeArrayDto1 == nil ||
		runeArrayDto2 == nil {

		return false
	}

	areEqual := new(runeArrayDtoElectron).
		equalCharArrays(
			runeArrayDto1,
			runeArrayDto2)

	if !areEqual {
		return false
	}

	if runeArrayDto1.Description1 !=
		runeArrayDto2.Description1 {
		return false
	}

	if runeArrayDto1.Description2 !=
		runeArrayDto2.Description2 {
		return false
	}

	if runeArrayDto1.charSearchType !=
		runeArrayDto2.charSearchType {
		return false
	}

	return true
}

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
//			runeArrayDto.CharsArray
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

//	deleteRunes
//
//	Receives a pointer to a rune array and proceeds to
//	delete rune text characters from the rune array
//	contained in the RuneArrayDto input parameter,
//	'runeArrayDto'. The name of the member variable
//	rune array which will be modified by this method is:
//
//			runeArrayDto.CharsArray
//
//	The rune characters deleted will either be leading
//	characters or trailing characters depending on the
//	setting for input parameter, 'deleteTrailingChars'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	runeArrayDto				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		rune array contained in this RuneArrayDto
//		instance will be modified in that rune characters
//		will be deleted. The name of the member variable
//		which will be modified is:
//
//			runeArrayDto.CharsArray
//
//	numOfRunesToDelete			uint64
//
//		This uint64 parameter specifies the number of
//		rune characters which will be deleted from the
//		rune array contained in parameter
//		'runeArrayDto'. These runes will be deleted
//		from rune array 'runeArrayDto.CharsArray'.
//
//		If this parameter is set to zero, no rune
//		characters will be deleted and no error will be
//		returned.
//
//		If this parameter is set to a value greater than
//		or equal to the length of the rune array, the
//		rune array will be set to 'nil' and no error will
//		be returned.
//
//	deleteTrailingChars			bool
//
//		This parameter determines whether the rune
//		characters deleted from the rune array will be
//		trailing characters or leading characters.
//
//		If this parameter is set to 'true', trailing
//		characters at the end of the rune array will be
//		deleted.
//
//		If this parameter is set to 'false', leading
//		characters at the beginning of the rune array
//		will be deleted.
func (runeDtoAtom *runeArrayDtoAtom) deleteRunes(
	runeArrayDto *RuneArrayDto,
	numOfRunesToDelete uint64,
	deleteTrailingChars bool,
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

	lenCharArray := uint64(len(runeArrayDto.CharsArray))

	if numOfRunesToDelete == 0 ||
		lenCharArray == 0 {
		return err
	}

	if numOfRunesToDelete >=
		lenCharArray {

		runeArrayDto.CharsArray = nil

		return err
	}

	// MUST BE: numOfRunesToDelete < lenCharArray

	// Slice Examples
	//arr := []int{1,2,3,4,5}
	//
	// Length = 5
	// Last Index = 4
	//
	//fmt.Println(arr[:2])        // [1,2]
	//
	//fmt.Println(arr[2:])        // [3,4,5]
	//
	//fmt.Println(arr[2:3])        // [3]
	//
	//fmt.Println(arr[:])            // [1,2,3,4,5]

	var targetIdx uint64

	if deleteTrailingChars == true {

		//  2     =  5     -      3
		targetIdx =
			lenCharArray - numOfRunesToDelete

		runeArrayDto.CharsArray =
			runeArrayDto.CharsArray[:targetIdx]

	} else {
		// MUST BE: Delete Leading Chars

		//   3    =  5   -   3  + 1
		targetIdx =
			lenCharArray - numOfRunesToDelete + 1

		runeArrayDto.CharsArray =
			runeArrayDto.CharsArray[targetIdx:]

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

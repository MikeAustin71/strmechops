package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoAtom struct {
	lock *sync.Mutex
}

//	addChar
//
//	Adds a single rune text character to the rune array
//	contained in the RuneArrayDto input parameter,
//	'runeArrayDto'.
//
//	The rune text character will be added either to
//	the beginning of the rune array or to the end of
//	the rune array depending on the setting for input
//	parameter 'addTrailingChar'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	runeArrayDto				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		rune array contained in this RuneArrayDto
//		instance will be modified. A single rune text
//		character will either be added to the beginning,
//		or the end, of this rune array:
//
//			runeArrayDto.CharsArray
//
//	charToAdd					rune
//
//		A single rune text character which will be
//		prepended or appended to the rune array contained
//		in the RuneArrayDto input parameter,
//		'runeArrayDto'.
//
//	addTrailingChar				bool
//
//		This parameter determines whether 'charToAdd'
//		will be added to the existing rune array,
//		'runeArrayDto.CharsArray', as a trailing
//		character or as a leading character.
//
//		If 'addTrailingChars' is set to 'true',
//		'charsToAdd' will be added to the end of the
//		existing rune array as a trailing character.
//
//		If 'addTrailingChars' is set to 'false',
//		'charToAdd' will be added to the beginning of
//		the existing rune array as leading character.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (runeDtoAtom *runeArrayDtoAtom) addChar(
	runeArrayDto *RuneArrayDto,
	charToAdd rune,
	addTrailingChar bool,
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
		"runeArrayDtoAtom."+
			"addChar()",
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

	if charToAdd == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charToAdd' is invalid!\n"+
			"'charToAdd' has a zero value!\n",
			ePrefix.String())

		return err
	}

	if addTrailingChar == true {
		// Add trailing suffix Character

		runeArrayDto.CharsArray = append(
			runeArrayDto.CharsArray,
			charToAdd)

	} else {
		// MUST BE - Add leading Prefix Character
		runeArrayDto.CharsArray = append(
			[]rune{charToAdd},
			runeArrayDto.CharsArray...)

	}

	return err
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
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
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
		"runeArrayDtoAtom."+
			"addRunes()",
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
			runeArrayDto.CharsArray,
			charsToAdd...)

	} else {

		runeArrayDto.CharsArray = append(
			charsToAdd,
			runeArrayDto.CharsArray...)

	}

	return err
}

//	addRuneArrayDtos
//
//	Receives a series of RuneArrayDto objects and proceeds
//	to add the rune arrays contained in these objects to
//	the existing rune array contained in the RuneArrayDto
//	passed as an input parameter ('runeArrayDto').
//
//	The name of the rune array member variable which will
//	be modified by this method is:
//
//			runeArrayDto.CharsArray
//
//	The additional rune arrays are contained in a series
//	of RuneArrayDto objects passed as a variadic
//	argument.
//
//	Each additional rune array will be appended in
//	sequence to the end of the 'runeArrayDto' existing
//	rune array in sequence.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	runeArrayDto				*RuneArrayDto
//
//		A pointer to an instance of RuneArrayDto. The
//		rune array contained in this RuneArrayDto
//		instance will be modified. Additional rune arrays
//		will be appended to the end of rune array:
//
//			runeArrayDto.CharsArray
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
//	runeArrayDtosToAdd			... RuneArrayDto
//
//		This is a variadic argument consisting of a
//		variable number of RuneArrayDto objects.
//
//		The rune arrays contained in these RuneArrayDto
//		objects will be appended in sequence to the end
//		of the existing rune array contained in input
//		parameter, 'runeArrayDto'.
//
//			runeArrayDto.CharsArray
//
//		If this parameter has a length of zero, an error
//		will be returned.
func (runeDtoAtom *runeArrayDtoAtom) addRuneArrayDtos(
	runeArrayDto *RuneArrayDto,
	errPrefDto *ePref.ErrPrefixDto,
	runeArrayDtosToAdd ...RuneArrayDto) (
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
		"runeArrayDtoAtom."+
			"addRuneArrayDtos()",
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

	if len(runeArrayDtosToAdd) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeArrayDtosToAdd' is invalid!\n"+
			"'runeArrayDtosToAdd' is empty!\n",
			ePrefix.String())

		return err
	}

	for _, runeArrayDtoToAdd := range runeArrayDtosToAdd {

		runeArrayDto.CharsArray = append(
			runeArrayDto.CharsArray,
			runeArrayDtoToAdd.CharsArray...)

	}

	return err
}

//	addRuneArrays
//
//	Adds a series of rune arrays to the existing rune
//	array contained in the RuneArrayDto passed as an
//	input parameter ('runeArrayDto').
//
//	The additional rune arrays will be appended to the
//	end of the existing rune array.
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
//		instance will be modified. Additional rune arrays
//		will be appended to the end of rune array:
//
//			runeArrayDto.CharsArray
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
//	charArraysToAdd				... []rune
//
//		This is a variadic argument consisting of a
//		variable number of rune arrays.
//
//		The rune arrays will be appended in sequence
//		to the end of the existing rune array contained
//		in input parameter, 'runeArrayDto'.
//
//			runeArrayDto.CharsArray
//
//		If this parameter has a length of zero, an error
//		will be returned.
func (runeDtoAtom *runeArrayDtoAtom) addRuneArrays(
	runeArrayDto *RuneArrayDto,
	errPrefDto *ePref.ErrPrefixDto,
	charArraysToAdd ...[]rune) (
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
		"runeArrayDtoAtom."+
			"addRuneArrays()",
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

	if len(charArraysToAdd) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charArraysToAdd' is invalid!\n"+
			"'charArraysToAdd' is empty!\n",
			ePrefix.String())

		return err
	}

	for _, charArray := range charArraysToAdd {

		runeArrayDto.CharsArray = append(
			runeArrayDto.CharsArray,
			charArray...)

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
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
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
		"runeArrayDtoAtom."+
			"deleteRunes()",
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
			numOfRunesToDelete

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

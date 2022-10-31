package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoQuark struct {
	lock *sync.Mutex
}

//	isRuneArrayAllNumericDigits
//
//	Receives a pointer to a RuneArrayDto and proceeds to
//	examine the internal rune array to determine if the
//	array contains all numeric character digits in range
//	of '0' through '9', inclusive.
//
//	The name of the member variable rune array is:
//
//			RuneArrayDto.CharsArray
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	runeArrayDto		*RuneArrayDto
//
//		A pointer to a RuneArrayDto object containing a
//		rune array. The contents of this rune array will
//		be examined to determine if the member elements
//		consist exclusively of numeric character digits
//		in the range '0' through '9', inclusive.
//
//		If the rune array is 'nil', or has a length of
//		zero, this method will return 'false' and no
//		error will be generated.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isAllNumericDigits			bool
//
//		If the rune array contained within input
//		paramter, 'runeArrayDto' contains all numeric
//		characters in the range '0' through '9',
//		inclusive, this	method will return 'true'.
//
//		Otherwise, this return parameter is set to
//		'false'.
func (runeDtoQuark *runeArrayDtoQuark) isRuneArrayAllNumericDigits(
	runeArrayDto *RuneArrayDto) (
	isAllNumericDigits bool) {

	if runeDtoQuark.lock == nil {
		runeDtoQuark.lock = new(sync.Mutex)
	}

	runeDtoQuark.lock.Lock()

	defer runeDtoQuark.lock.Unlock()

	isAllNumericDigits = false

	if runeArrayDto == nil {

		return isAllNumericDigits

	}

	lenTargetAry := len(runeArrayDto.CharsArray)

	if lenTargetAry == 0 {

		return isAllNumericDigits
	}

	for i := 0; i < lenTargetAry; i++ {
		if runeArrayDto.CharsArray[i] < '0' ||
			runeArrayDto.CharsArray[i] > '9' {

			return isAllNumericDigits
		}
	}

	isAllNumericDigits = true

	return isAllNumericDigits
}

//	isRuneArrayAllNumericZeros
//
//	Receives a pointer to a RuneArrayDto and proceeds to
//	examine the internal rune array to determine if the
//	array consists entirely of the numeric character
//	digit, zero ('0').
//
//	The name of the member variable rune array tested is:
//
//			RuneArrayDto.CharsArray
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	runeArrayDto				*RuneArrayDto
//
//		A pointer to a RuneArrayDto object containing a
//		rune array. The contents of this rune array will
//		be examined to determine if all the member
//		elements contain the numeric character zero
//		('0').
//
//		If all the rune array member elements consist of
//		the zero numeric character, '0', this method
//		returns true.
//
//		If the rune array is nil or a length of zero,
//		this method returns 'false'. Likewise, if any of
//		the rune array member elements are not equal to
//		the zero numeric character ('0'), this method
//		will return 'false'.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isAllNumericZeros			bool
//
//		If the rune array contained within input
//		paramter, 'runeArrayDto', contains only zero
//		numeric characters ('0'), this method will return
//		'true'.
//
//		Otherwise, this return parameter is set to
//		'false'.
func (runeDtoQuark *runeArrayDtoQuark) isRuneArrayAllNumericZeros(
	runeArrayDto *RuneArrayDto) (
	isAllZeros bool) {

	if runeDtoQuark.lock == nil {
		runeDtoQuark.lock = new(sync.Mutex)
	}

	runeDtoQuark.lock.Lock()

	defer runeDtoQuark.lock.Unlock()

	isAllZeros = false

	if runeArrayDto == nil {

		return isAllZeros
	}

	lenTargetAry := len(runeArrayDto.CharsArray)

	if lenTargetAry == 0 {

		return isAllZeros
	}

	for i := 0; i < lenTargetAry; i++ {

		if runeArrayDto.CharsArray[i] != '0' {

			return isAllZeros
		}
	}

	isAllZeros = true

	return isAllZeros
}

// isValidCharacterArray - Tests the validity of the Character
// Array contained in an instance of RuneArrayDto.
//
// The RuneArrayDto character array is stored in member variable
// 'CharsArray'.
//
// If the character array has a length of zero, it is classified as
// invalid.  If the array length is greater than zero, the
// character array is classified as valid.
//
// The Character Array member variable is styled as:
//
//	runeArrayDto.CharsArray
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto               *RuneArrayDto
//	   - A pointer to an instance of RuneArrayDto. The internal
//	     member variable, 'CharsArray' (Character Array),
//	     will be tested for validity.
//
//	     If the character array has a length of zero, it is
//	     classified as invalid.  If the array length is greater
//	     than zero, the character array is classified as valid.
//
//
//	errPrefDto                   *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	isValid                    bool
//	   - If 'runeArrayDto' member variable 'CharsArray' (Character
//	     Array) is judged to be valid in all respects, this return
//	     parameter will be set to 'true'.
//
//	     If 'runeArrayDto' member variable 'CharsArray' (Character
//	     Array) is found to be invalid, this return parameter will
//	     be set to 'false'.
//
//
//	err                        error
//	   - If 'runeArrayDto' member variable 'CharsArray' (Character
//	     Array) is judged to be valid in all respects, this return
//	     parameter will be set to 'nil'.
//
//	     If 'runeArrayDto' member variable 'CharsArray' (Character
//	     Array) is found to be invalid, this return parameter will
//	     be configured with an appropriate error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (runeDtoQuark *runeArrayDtoQuark) isValidCharacterArray(
	runeArrayDto *RuneArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if runeDtoQuark.lock == nil {
		runeDtoQuark.lock = new(sync.Mutex)
	}

	runeDtoQuark.lock.Lock()

	defer runeDtoQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoQuark."+
			"isValidCharacterArray()",
		"")

	if err != nil {

		return isValid, err

	}

	if runeArrayDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeArrayDto' is invalid!\n"+
			"'runeArrayDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err

	}

	if len(runeArrayDto.CharsArray) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: The Character Array for this instance of\n"+
			"RuneArrayDto is invalid!\n"+
			"RuneArrayDto.CharsArray is empty, has a length\n"+
			"of zero and therefore contains zero characters.\n",
			ePrefix.String())

		return isValid, err
	}

	isValid = true

	return isValid, err
}

// isValidCharacterSearchType - Tests the validity of the Character
// Search Type member variable ('charSearchType') contained in a
// RuneArrayDto instance.
//
// The Character Search Type variable is styled as:
//
//	runeArrayDto.charSearchType
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	runeArrayDto               *RuneArrayDto
//	   - A pointer to an instance of RuneArrayDto. The internal
//	     member variable, 'charSearchType' (Character Search Type),
//	     will be tested for validity.
//
//
//	errPrefDto                   *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	isValid                    bool
//	   - If 'runeArrayDto' member variable 'charSearchType'
//	     (Character Search Type) is judged to be valid in all
//	     respects, this return parameter will be set to 'true'.
//
//	     If 'runeArrayDto' member variable 'charSearchType'
//	     (Character Search Type) is found to be invalid,
//	     this return parameter will be set to 'false'.
//
//
//	err                        error
//	   - If 'runeArrayDto' member variable 'charSearchType'
//	     (Character Search Type) is judged to be valid in
//	     all respects, this return parameter will be set to 'nil'.
//
//	     If 'runeArrayDto' member variable 'charSearchType'
//	     (Character Search Type) is found to be invalid, this
//	     return parameter will be configured with an appropriate
//	     error message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (runeDtoQuark *runeArrayDtoQuark) isValidCharacterSearchType(
	runeArrayDto *RuneArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if runeDtoQuark.lock == nil {
		runeDtoQuark.lock = new(sync.Mutex)
	}

	runeDtoQuark.lock.Lock()

	defer runeDtoQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"runeArrayDtoQuark."+
			"isValidCharacterSearchType()",
		"")

	if err != nil {

		return isValid, err

	}

	if runeArrayDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'runeArrayDto' is invalid!\n"+
			"'runeArrayDto' is a 'nil' pointer!\n",
			ePrefix.String())

		return isValid, err

	}

	if !runeArrayDto.charSearchType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: RuneArrayDto instance has an invalid\n"+
			"Character Search Type ('charSearchType')!\n"+
			"'charSearchType' must be set to one of three enumerations:\n"+
			"(1) CharacterSearchType(0).LinearTargetStartingIndex()\n"+
			"(2) CharacterSearchType(0).SingleTargetChar()\n"+
			"(3) CharacterSearchType(0).LinearEndOfString()\n"+
			"'charSearchType' string  value = '%v'\n"+
			"'charSearchType' integer value = '%v'\n",
			ePrefix.String(),
			runeArrayDto.charSearchType.String(),
			runeArrayDto.charSearchType.XValueInt())

		return isValid, err

	}

	isValid = true

	return isValid, err
}

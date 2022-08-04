package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type runeArrayDtoQuark struct {
	lock *sync.Mutex
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
//   runeArrayDto.CharsArray
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  runeArrayDto               *RuneArrayDto
//     - A pointer to an instance of RuneArrayDto. The internal
//       member variable, 'CharsArray' (Character Array),
//       will be tested for validity.
//
//       If the character array has a length of zero, it is
//       classified as invalid.  If the array length is greater
//       than zero, the character array is classified as valid.
//
//
//  errPrefDto                   *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  isValid                    bool
//     - If 'runeArrayDto' member variable 'CharsArray' (Character
//       Array) is judged to be valid in all respects, this return
//       parameter will be set to 'true'.
//
//       If 'runeArrayDto' member variable 'CharsArray' (Character
//       Array) is found to be invalid, this return parameter will
//       be set to 'false'.
//
//
//  err                        error
//     - If 'runeArrayDto' member variable 'CharsArray' (Character
//       Array) is judged to be valid in all respects, this return
//       parameter will be set to 'nil'.
//
//       If 'runeArrayDto' member variable 'CharsArray' (Character
//       Array) is found to be invalid, this return parameter will
//       be configured with an appropriate error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
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
//   runeArrayDto.charSearchType
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  runeArrayDto               *RuneArrayDto
//     - A pointer to an instance of RuneArrayDto. The internal
//       member variable, 'charSearchType' (Character Search Type),
//       will be tested for validity.
//
//
//  errPrefDto                   *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  isValid                    bool
//     - If 'runeArrayDto' member variable 'charSearchType'
//       (Character Search Type) is judged to be valid in all
//       respects, this return parameter will be set to 'true'.
//
//       If 'runeArrayDto' member variable 'charSearchType'
//       (Character Search Type) is found to be invalid,
//       this return parameter will be set to 'false'.
//
//
//  err                        error
//     - If 'runeArrayDto' member variable 'charSearchType'
//       (Character Search Type) is judged to be valid in
//       all respects, this return parameter will be set to 'nil'.
//
//       If 'runeArrayDto' member variable 'charSearchType'
//       (Character Search Type) is found to be invalid, this
//       return parameter will be configured with an appropriate
//       error message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
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

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// RuneArrayDto - A Rune Array Data Transfer Object. This type is
// designed to transfer a rune array between various functions and
// types.
//
// This type serves as a wrapper for a rune array and allows methods
// to pass pointers to RuneArrayDto and then perform operations on
// the internal rune array, 'CharsArray'.
//
// For example, this rune array can be used to carry out text
// character searches by passing pointers to the RuneArrayDto.
//
type RuneArrayDto struct {
	CharsArray []rune

	lock *sync.Mutex
}

// CopyIn - Copies the internal rune array from an incoming
// instance of RuneArrayDto ('incomingChars') to the internal
// rune array of the current RuneArrayDto instance
// ('charsArrayDto').
//
// IMPORTANT
// The internal rune array for the current RuneArrayDto instance
// ('charsArrayDto') will be deleted and overwritten.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  incomingChars              *RuneArrayDto
//     - A pointer to an instance of RuneArrayDto. This method will
//       NOT change the values of internal rune array contained in
//       this instance.
//
//       All data values in this RuneArrayDto instance will be
//       copied to current RuneArrayDto instance ('charsArrayDto').
//
//       If parameter 'incomingChars' is 'nil' pointer or if its
//       internal rune array has a length of zero, an error will be
//       returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto *RuneArrayDto) CopyIn(
	incomingChars *RuneArrayDto,
	errorPrefix interface{}) (
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	if incomingChars == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingChars' is a 'nil' pointer!\n",
			ePrefix.String())

		return err

	}

	lenIncomingCharsArray := len(incomingChars.CharsArray)

	if lenIncomingCharsArray == 0 {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingChars' is invalid!\n"+
			"The length of 'incomingChars.CharsArray' is zero.\n"+
			"The 'incomingChars' internal Rune Array is empty.",
			ePrefix.String())

		return err

	}

	charsArrayDto.CharsArray =
		make([]rune, lenIncomingCharsArray)

	for i := 0; i < lenIncomingCharsArray; i++ {
		charsArrayDto.CharsArray[i] =
			incomingChars.CharsArray[i]
	}

	return err
}

// CopyOut - Returns a deep copy of the current RuneArrayDto
// instance.
//
// If the current RuneArrayDto instance contains an empty rune
// array, this method will return an error.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  copyOfRuneArrayDto         copyOfRuneArrayDto
//     - If this method completes successfully and no errors are
//       encountered, this parameter will return a deep copy of the
//       current RuneArrayDto instance.
//
//
//  err                        error
//     - If the method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto *RuneArrayDto) CopyOut(
	errorPrefix interface{}) (
	copyOfRuneArrayDto RuneArrayDto,
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"CopyOut()",
		"")

	if err != nil {
		return copyOfRuneArrayDto, err
	}

	lenOfLocalCharsArray := len(charsArrayDto.CharsArray)

	if lenOfLocalCharsArray == 0 {

		err = fmt.Errorf("%v\n" +
			"ERROR: The current instance of RuneArrayDto has an empty rune array!\n" +
			"The length of 'charsArrayDto.CharsArray' is zero.\n" +
			ePrefix.String())

		return copyOfRuneArrayDto, err
	}

	copyOfRuneArrayDto.CharsArray =
		make([]rune, lenOfLocalCharsArray)

	for i := 0; i < lenOfLocalCharsArray; i++ {
		copyOfRuneArrayDto.CharsArray[i] =
			charsArrayDto.CharsArray[i]
	}

	return copyOfRuneArrayDto, err
}

// Empty - Resets the internal member variable rune array to 'nil'.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// This method will delete the pre-existing characters in the
// internal rune array for the current instance of RuneArrayDto.
//
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  NONE
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  NONE
//
func (charsArrayDto *RuneArrayDto) Empty() {

	charsArrayDto.lock.Lock()

	charsArrayDto.CharsArray = nil

	charsArrayDto.lock.Unlock()

	charsArrayDto.lock = nil
}

// GetRuneArrayLength - Returns the length of the internal
// rune array, 'CharsArray' as an integer value.
//
func (charsArrayDto *RuneArrayDto) GetRuneArrayLength() int {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	return len(charsArrayDto.CharsArray)
}

// GetRuneArray - Returns a deep copy of the internal rune array,
// 'CharsArray' configured for the current instance of
// RuneArrayDto.
//
// If the length of the internal rune array 'CharsArray' is zero,
// a value of 'nil' will be returned.
//
func (charsArrayDto *RuneArrayDto) GetRuneArray() []rune {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	lenCharsArray := len(charsArrayDto.CharsArray)

	if lenCharsArray == 0 {
		return nil
	}

	newCharsArray := make([]rune, lenCharsArray)

	for i := 0; i < lenCharsArray; i++ {
		newCharsArray[i] =
			charsArrayDto.CharsArray[i]
	}

	return newCharsArray
}

// NewRuneArray - Receives an array of runes and returns a new
// instance of RuneArrayDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  charArray                  []rune
//     - An array of runes used to populate a new instance of
//       RuneArrayDto which is returned to the calling function.
//
//       If this array is empty or has a zero length, an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  newRuneArrayDto            RuneArrayDto
//     - If this method completes successfully a new, fully
//       populated instance of RuneArrayDto will be returned to the
//       calling function.
//
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto RuneArrayDto) NewRuneArray(
	charArray []rune,
	errorPrefix interface{}) (
	newRuneArrayDto RuneArrayDto,
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"NewRuneArray()",
		"")

	if err != nil {
		return newRuneArrayDto, err
	}

	lenOfCharArray := len(charArray)

	if lenOfCharArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charArray' is invalid!\n"+
			"'charArray' is empty and has a length of zero.\n",
			ePrefix.String())

		return newRuneArrayDto, err
	}

	newRuneArrayDto.CharsArray =
		make([]rune, lenOfCharArray)

	for i := 0; i < lenOfCharArray; i++ {
		newRuneArrayDto.CharsArray[i] =
			charArray[i]
	}

	return newRuneArrayDto, err
}

// SetRuneArray - Deletes the internal rune array for the current
// instance of RuneArrayDto and repopulates it with a deep copy
// extracted from input parameter, 'charArray'.
//
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// All pre-existing data in the current instance of RuneArrayDto
// will be deleted by this method.
//
// The pre-existing data in the internal rune array member
// variable, 'CharsArray' will be deleted and overwritten with new
// data.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  charArray                  []rune
//     - An array of runes used to populate a new instance of
//       RuneArrayDto which is returned to the calling function.
//
//       If this array is empty or has a zero length, an error will
//       be returned.
//
//
//  errorPrefix                interface{}
//     - This object encapsulates error prefix text which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods
//       listed as a method or function chain of execution.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       This empty interface must be convertible to one of the
//       following types:
//
//
//       1. nil - A nil value is valid and generates an empty
//                collection of error prefix and error context
//                information.
//
//       2. string - A string containing error prefix information.
//
//       3. []string A one-dimensional slice of strings containing
//                   error prefix information
//
//       4. [][2]string A two-dimensional slice of strings containing
//                      error prefix and error context information.
//
//       5. ErrPrefixDto - An instance of ErrPrefixDto. The
//                         ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       6. *ErrPrefixDto - A pointer to an instance of ErrPrefixDto.
//                          ErrorPrefixInfo from this object will be
//                         copied to 'errPrefDto'.
//
//       7. IBasicErrorPrefix - An interface to a method generating
//                              a two-dimensional slice of strings
//                              containing error prefix and error
//                              context information.
//
//       If parameter 'errorPrefix' is NOT convertible to one of
//       the valid types listed above, it will be considered
//       invalid and trigger the return of an error.
//
//       Types ErrPrefixDto and IBasicErrorPrefix are included in
//       the 'errpref' software package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (charsArrayDto *RuneArrayDto) SetRuneArray(
	charArray []rune,
	errorPrefix interface{}) (
	err error) {

	if charsArrayDto.lock == nil {
		charsArrayDto.lock = new(sync.Mutex)
	}

	charsArrayDto.lock.Lock()

	defer charsArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"SetRuneArray()",
		"")

	if err != nil {
		return err
	}

	lenOfCharArray := len(charArray)

	if lenOfCharArray == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'charArray' is invalid!\n"+
			"'charArray' is empty and has a length of zero.\n",
			ePrefix.String())

		return err
	}

	charsArrayDto.CharsArray =
		make([]rune, lenOfCharArray)

	for i := 0; i < lenOfCharArray; i++ {
		charsArrayDto.CharsArray[i] =
			charArray[i]
	}

	return err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// ByteArrayDto
//
// The Byte Array Data Transfer Object is a wrapper for
// a standard byte array.
type ByteArrayDto struct {
	ByteArray []byte

	Description1 string
	// Optional. A name, label or narrative text used to describe
	// the current instance of ByteArrayDto.

	Description2 string
	// Optional. A name, label or narrative text used to describe
	// the current instance of ByteArrayDto.

	lock *sync.Mutex
}

// AddRunes
//
// This method will add one or more runes to the byte
// array encapsulated by the current instance of
// ByteArrayDto.
//
// Depending on the value of input parameter
// 'addTrailingRunes', the rune(s) to be added will
// either be prefixed or suffixed to the current byte
// array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	addTrailingRunes			bool
//
//		If this boolean value is set to 'true', the
//		runes contained in parameter 'runesToAdd' will
//		be appended to the end of the byte array
//		contained in the current instance of
//		ByteArrayDto.
//
//		If 'addTrailingRunes' is set to 'false', the
//		runes to be added will be prefixed to the
//		beginning of the byte array contained in the
//		current instance of ByteArrayDto.
//
//	runesToAdd 					...rune
//
//		'runesToAdd' is a variadic parameter which
//		will accept one or more runes for addition to
//		the byte array maintained by the current instance
//		of ByteArrayDto.
//
//		Input parameter 'addTrailingRunes' controls
//		whether 'runesToAdd' will be prefixed or suffixed
//		to the byte array.
//
//		If 'runesToAdd' is empty and contains zero runes,
//		this method will take no action and exit.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (byteArrayDto *ByteArrayDto) AddRunes(
	addTrailingRunes bool,
	runesToAdd ...rune) {

	if byteArrayDto.lock == nil {
		byteArrayDto.lock = new(sync.Mutex)
	}

	byteArrayDto.lock.Lock()

	defer byteArrayDto.lock.Unlock()

	var newByteArray []byte

	for _, singleRune := range runesToAdd {

		newByteArray = append(
			newByteArray,
			[]byte(string(singleRune))...)

	}

	if len(newByteArray) == 0 {

		return
	}

	if len(byteArrayDto.ByteArray) == 0 ||
		addTrailingRunes == true {

		byteArrayDto.ByteArray = append(
			byteArrayDto.ByteArray, newByteArray...)

	} else {
		// len(byteArrayDto.ByteArray) > 0
		//         AND
		// addTrailingRune == false
		byteArrayDto.ByteArray = append(
			newByteArray,
			byteArrayDto.ByteArray...)

	}

	return
}

// AddStrings
//
// Adds one or more strings to the byte array
// encapsulated by the current instance of ByteArray.
//
// Depending on the value of input parameter
// 'addTrailingStrings', the string(s) to be added will
// either be prefixed or suffixed to the current byte
// array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	addTrailingStrings			bool
//
//		If this boolean value is set to 'true', the
//		strings contained in parameter 'stringsToAdd'
//		will be appended to the end of the byte array
//		contained in the current instance of
//		ByteArrayDto.
//
//		If 'addTrailingStrings' is set to 'false', the
//		strings to be added will be prefixed to the
//		beginning of the byte array contained in the
//		current instance of ByteArrayDto.
//
//	stringsToAdd 				...string
//
//		'stringsToAdd' is a variadic parameter which
//		will accept one or more strings for addition to
//		the byte array maintained by the current instance
//		of ByteArrayDto.
//
//		Input parameter 'addTrailingStrings' controls
//		whether 'stringsToAdd' will be prefixed or
//		suffixed to the byte array.
//
//		If 'stringsToAdd' is empty and contains zero
//		strings, this method will take no action and
//		exit.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (byteArrayDto *ByteArrayDto) AddStrings(
	addTrailingStrings bool,
	stringsToAdd ...string) {

	if byteArrayDto.lock == nil {
		byteArrayDto.lock = new(sync.Mutex)
	}

	byteArrayDto.lock.Lock()

	defer byteArrayDto.lock.Unlock()

	var newByteArray []byte

	for _, singleString := range stringsToAdd {

		newByteArray = append(
			newByteArray,
			[]byte(singleString)...)

	}

	if len(newByteArray) == 0 {

		return
	}

	if len(byteArrayDto.ByteArray) == 0 ||
		addTrailingStrings == true {

		byteArrayDto.ByteArray = append(
			byteArrayDto.ByteArray, newByteArray...)

	} else {
		// len(byteArrayDto.ByteArray) > 0
		//         AND
		// addTrailingRune == false
		byteArrayDto.ByteArray = append(
			newByteArray,
			byteArrayDto.ByteArray...)

	}

	return
}

// Clear
//
// Resets each byte in the internal byte array contained
// in the current instance of ByteArray to a zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing data
//	values in the internal byte array maintained by
//	the current instance of ByteArrayDto. The array
//	length will remain unchanged. However, the data
//	value of each element in the array will be reset
//	to zero ('0').
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (byteArrayDto *ByteArrayDto) Clear() {

	if byteArrayDto.lock == nil {
		byteArrayDto.lock = new(sync.Mutex)
	}

	byteArrayDto.lock.Lock()

	defer byteArrayDto.lock.Unlock()

	if len(byteArrayDto.ByteArray) == 0 {

		return
	}

	clear(byteArrayDto.ByteArray)

	return
}

//	CopyIn
//
//	Copies the internal byte array from an incoming
//	instance of ByteArrayDto ('incomingBArrayDto') to the
//	internal rune array of the current ByteArrayDto
//	instance. This copy operation is a 'deep' copy.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The internal rune array for the current ByteArrayDto
//	instance ('byteArrayDto.ByteArray') will be deleted
//	and overwritten with new data copied from
//	'incomingBArrayDto'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingBArrayDto			*ByteArrayDto
//
//		A pointer to an incoming instance of
//		ByteArrayDto. This method will NOT change the
//		values of the internal byte array contained in
//		this instance.
//
//		All data values in this ByteArrayDto instance
//		will be copied to current ByteArrayDto instance
//		('byteArrayDto.ByteArray').
//
//		If parameter 'incomingChars' is a 'nil' pointer
//		or if its internal rune array has a length of
//		zero, an error will be returned.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it	contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		This empty interface must be convertible to one of
//		the following types:
//
//		1.	nil
//				A nil value is valid and generates an
//				empty collection of error prefix and
//				error context information.
//
//		2.	string
//				A string containing error prefix
//				information.
//
//		3.	[]string
//				A one-dimensional slice of strings
//				containing error prefix information.
//
//		4.	[][2]string
//				A two-dimensional slice of strings
//		   		containing error prefix and error
//		   		context information.
//
//		5.	ErrPrefixDto
//				An instance of ErrPrefixDto.
//				Information from this object will
//				be copied for use in error and
//				informational messages.
//
//		6.	*ErrPrefixDto
//				A pointer to an instance of
//				ErrPrefixDto. Information from
//				this object will be copied for use
//				in error and informational messages.
//
//		7.	IBasicErrorPrefix
//				An interface to a method
//				generating a two-dimensional slice
//				of strings containing error prefix
//				and error context information.
//
//		If parameter 'errorPrefix' is NOT convertible
//		to one of the valid types listed above, it will
//		be considered invalid and trigger the return of
//		an error.
//
//		Types ErrPrefixDto and IBasicErrorPrefix are
//		included in the 'errpref' software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message. This returned error message will
//		incorporate the method chain and text passed by
//		input parameter, 'errorPrefix'. The 'errorPrefix'
//		text will be attached to the beginning of the
//		error message.
func (byteArrayDto *ByteArrayDto) CopyIn(
	incomingBArrayDto *ByteArrayDto,
	errorPrefix interface{}) error {

	if byteArrayDto.lock == nil {
		byteArrayDto.lock = new(sync.Mutex)
	}

	byteArrayDto.lock.Lock()

	defer byteArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"RuneArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	lenIncomingBArray := len(incomingBArrayDto.ByteArray)

	if lenIncomingBArray == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'incomingBArrayDto' is invalid!\n"+
			"'incomingBArrayDto' contains an empty or zero length\n"+
			"byte array.\n",
			ePrefix.String())

		return err
	}

	byteArrayDto.ByteArray = make([]byte, lenIncomingBArray)

	for i := 0; i < lenIncomingBArray; i++ {

		byteArrayDto.ByteArray[i] =
			incomingBArrayDto.ByteArray[i]

	}

	return err
}

// CopyOut
//
// Returns a deep copy of the internal byte array
// encapsulated by the current instance of ByteArrayDto.
//
// If the internal byte array is empty or has a zero
// length, this method will return a zero length byte
// array.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	[]byte
//
//		This method returns a deep copy of the internal
//		byte array maintained by the current instance of
//		ByteArrayDto.
func (byteArrayDto *ByteArrayDto) CopyOut() []byte {

	if byteArrayDto.lock == nil {
		byteArrayDto.lock = new(sync.Mutex)
	}

	byteArrayDto.lock.Lock()

	defer byteArrayDto.lock.Unlock()

	lenCurrByteArray := len(byteArrayDto.ByteArray)

	if lenCurrByteArray == 0 {

		return make([]byte, 0)
	}

	newByteArray := make([]byte, lenCurrByteArray)

	for i := 0; i < lenCurrByteArray; i++ {

		newByteArray[i] = byteArrayDto.ByteArray[i]

	}

	return newByteArray
}

// Empty
//
// Resets the internal byte array contained in the
// current instance of ByteArray to value of 'nil'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing data
//	values in the internal byte array maintained by
//	the current instance of ByteArrayDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (byteArrayDto *ByteArrayDto) Empty() {

	if byteArrayDto.lock == nil {
		byteArrayDto.lock = new(sync.Mutex)
	}

	byteArrayDto.lock.Lock()

	byteArrayDto.ByteArray = nil

	byteArrayDto.lock.Unlock()

	byteArrayDto.lock = nil
}

// String
//
// Returns the internal byte array maintained by the
// current instance of ByteArrayDto as a string.
//
// This method implements the Stringer interface.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	-- NONE --
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	string
//
//		This method returns a string of characters
//		extracted from the internal byte array maintained
//		by the current instance of ByteArrayDto.
func (byteArrayDto *ByteArrayDto) String() string {

	if byteArrayDto.lock == nil {
		byteArrayDto.lock = new(sync.Mutex)
	}

	byteArrayDto.lock.Lock()

	defer byteArrayDto.lock.Unlock()

	if len(byteArrayDto.ByteArray) == 0 {

		return ""
	}

	return string(byteArrayDto.ByteArray)
}

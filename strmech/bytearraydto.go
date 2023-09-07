package strmech

import "sync"

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

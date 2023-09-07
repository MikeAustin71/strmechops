package strmech

import "sync"

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

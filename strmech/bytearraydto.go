package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// ByteArrayDto
//
// The Byte Array Data Transfer Object is a wrapper for
// a standard byte array or more accurately, a slice of
// bytes.
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
// in the current instance of ByteArrayDto to a zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete all pre-existing data
//		values in the internal byte array maintained by
//		the current instance of ByteArrayDto.
//
//	(2)	The length of the internal byte array will
//		remain unchanged, however each element in that
//		internal byte array will be set to a zero ('0')
//		value.
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

	new(byteArrayDtoElectron).clear(byteArrayDto)

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
//	copyDescriptionData			bool
//
//		If this parameter is set to 'true', the
//		description data contained in 'incomingBArrayDto'
//		will also be copied to the current instance of
//		ByteArrayDto.
//
//		Description data is contained in two internal
//		member variables:
//
//			ByteArrayDto.Description1
//			ByteArrayDto.Description2
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
	copyDescriptionData bool,
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

	return new(byteArrayDtoAtom).
		copy(
			incomingBArrayDto,
			"incomingBArrayDto",
			byteArrayDto,
			"byteArrayDto",
			copyDescriptionData,
			ePrefix)
}

// CopyOut
//
// Returns a new instance of ByteArrayDto containing a
// deep copy of the member variable data values
// encapsulated by the current ByteArrayDto instance.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method is designed to return a new instance
//		of ByteArrayDto representing a 'deep' copy of the
//		internal byte array encapsulated in the current
//		ByteArrayDto instance.
//
//			ByteArrayDto.ByteArray
//
//	(2) If the current instance of ByteArrayDto contains
//		a zero length internal byte array, no error will
//		be returned and the returned new ByteArrayDto
//		instance will contain a zero length byte array.
//
//	(3)	If the input parameter 'copyDescriptionData' is
//		set to 'true', the description data contained
//		in the current ByteArrayDto instance will also be
//		returned in the new ByteArrayDto instance.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	copyDescriptionData			bool
//
//		If this parameter is set to 'true', the
//		description data contained in the current
//		instance of ByteArrayDto will be copied to
//		the returned instance ByteArrayDto.
//
//		Description data is contained in two internal
//		member variables:
//
//			ByteArrayDto.Description1
//			ByteArrayDto.Description2
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	ByteArrayDto
//
//		This method returns a new instance of ByteArrayDto
//		containing a deep copy of the internal byte array
//		encapsulated in the current instance of
//		ByteArrayDto.
//
//		If input parameter 'copyDescriptionData' is set
//		to 'true', the description data contained in the
//		current instance of ByteArrayDto will also be
//		copied to the returned ByteArrayDto instance.
//
//	errorPrefix					interface{}
//
//		This object encapsulates error prefix text which
//		is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a method or function
//		chain of execution.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		This empty interface must be convertible to one
//		of the following types:
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
func (byteArrayDto *ByteArrayDto) CopyOut(
	copyDescriptionData bool,
	errorPrefix interface{}) (
	ByteArrayDto,
	error) {

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
		"ByteArrayDto."+
			"GetEntryTypeComponent()",
		"")

	var newByteArrayDto ByteArrayDto

	err = new(byteArrayDtoAtom).
		copy(
			byteArrayDto,
			"byteArrayDto",
			&newByteArrayDto,
			"newByteArrayDto",
			copyDescriptionData,
			ePrefix)

	return newByteArrayDto, err
}

// DeleteLeadingBytes
//
// This method will delete a specified number of bytes
// from the leading edge of an internal byte array
// encapsulated by the current instance of ByteArrayDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete pre-existing data values
//		in the internal byte array encapsulated by the
//		current instance of ByteArrayDto.
//
//	(2)	This method is designed to delete leading bytes
//		from the internal byte array contained within the
//		current instance of ByteArrayDto.
//
//			ByteArrayDto.ByteArray
//
//	(3)	If the ByteArrayDto internal byte array is empty
//		with a length of zero ('0'), this method will
//		take no action, return no error and exit.
//
//		If the number of leading bytes to be deleted
//		('numOfLeadingBytesToDelete') is greater than or
//		equal to the current length of the internal byte
//		array, this method will set the internal byte
//		array to 'nil' (a zero length array) and no error
//		will be returned.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numOfLeadingBytesToDelete	int
//
//		The number of bytes which will be deleted from
//		the leading edge of the internal byte array
//		encapsulated by the current instance of
//		ByteArrayDto.
//
//		If the internal byte array is empty with a
//		length of zero ('0'), this method will take no
//		action, return no error and exit.
//
//		If the number of leading bytes to be deleted
//		('numOfLeadingBytesToDelete') is greater than or
//		equal to the current length of the internal byte
//		array, this method will set the internal byte
//		array to 'nil' (a zero length array) and no error
//		will be returned.
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (byteArrayDto *ByteArrayDto) DeleteLeadingBytes(
	numOfLeadingBytesToDelete int,
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
		"FilePermissionConfig."+
			"GetEntryTypeComponent()",
		"")

	if err != nil {
		return err
	}

	return new(byteArrayDtoAtom).
		deleteLeadingBytes(
			byteArrayDto,
			"byteArrayDto",
			numOfLeadingBytesToDelete,
			ePrefix)
}

// DeleteTrailingBytes
//
// This method will delete a specified number of bytes
// from the trailing edge of an internal byte array
// encapsulated by the current instance of ByteArrayDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete pre-existing data values
//		in the internal byte array encapsulated by the
//		current instance of ByteArrayDto.
//
//	(2)	This method is designed to delete trailing bytes
//		from the internal byte array contained within the
//		current instance of ByteArrayDto.
//
//			ByteArrayDto.ByteArray
//
//	(3)	If the ByteArrayDto internal byte array is empty
//		with a length of zero ('0'), this method will
//		take no action, return no error and exit.
//
//		If the number of trailing bytes to be deleted
//		('numOfTrailingBytesToDelete') is greater than or
//		equal to the current length of the internal byte
//		array, this method will set the internal byte
//		array to 'nil' (a zero length array) and no error
//		will be returned.
//
//	(4)	If the number of trailing bytes to delete
//		('numOfTrailingBytesToDelete') is set to zero
//		('0'), this method will take no action, return no
//		error and exit.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numOfTrailingBytesToDelete	int
//
//		The number of bytes which will be deleted from
//		the trailing edge of the internal byte array
//		encapsulated by the current instance of
//		ByteArrayDto.
//
//		If the number of trailing bytes to delete
//		('numOfTrailingBytesToDelete') is set to zero
//		('0'), this method will take no action, return
//		no error and exit.
//
//		If the internal byte array is empty with a
//		length of zero ('0'), this method will take no
//		action, return no error and exit.
//
//		If the number of trailing bytes to be deleted
//		('numOfTrailingBytesToDelete') is greater than or
//		equal to the current length of the internal byte
//		array, this method will set the internal byte
//		array to 'nil' (a zero length array) and no error
//		will be returned.
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (byteArrayDto *ByteArrayDto) DeleteTrailingBytes(
	numOfTrailingBytesToDelete int,
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
		"FilePermissionConfig."+
			"GetEntryTypeComponent()",
		"")

	if err != nil {
		return err
	}

	return new(byteArrayDtoAtom).
		deleteTrailingBytes(
			byteArrayDto,
			"byteArrayDto",
			numOfTrailingBytesToDelete,
			ePrefix)
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

	new(byteArrayDtoElectron).empty(
		byteArrayDto)

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

type byteArrayDtoAtom struct {
	lock *sync.Mutex
}

/*

package main

import (
	"fmt"
)


func deleteLeadingBytesMech() {


	var oldByteArray , newByteArray []byte

	oldByteArray = []byte("xxxHello")

	fmt.Println("-------------------------")
	fmt.Println(" Deleting Leading Bytes  ")
	fmt.Println("-------------------------")

	fmt.Printf("Old Byte Array: %v\n",
		string(oldByteArray))

	newByteArray = oldByteArray [3:]


	fmt.Printf("New Byte Array: %v\n\n",
	string(newByteArray))

}


func deleteTrailingBytesMech() {


	var oldByteArray , newByteArray []byte

	oldByteArray = []byte("Helloxxx")

	fmt.Println("-------------------------")
	fmt.Println(" Deleting Trailing Bytes   ")
	fmt.Println("-------------------------")

	fmt.Printf("Old Byte Array: %v\n",
		string(oldByteArray))

	lenOfOldByteArray := len(oldByteArray)

	lenOfBytesToKeep := lenOfOldByteArray - 3

	newByteArray = oldByteArray [0:lenOfBytesToKeep]


	fmt.Printf("New Byte Array: %v\n",
	string(newByteArray))

}

func main() {

	deleteLeadingBytesMech()
	deleteTrailingBytesMech()
}



*/

// copy
//
//	Copies the internal byte array from an incoming
//	instance of ByteArrayDto ('sourceBArray') to the
//	internal rune array of the destination ByteArrayDto
//	instance ('destinationBArray').
//
//	This copy operation is a 'deep' copy.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	sourceBArray				*ByteArrayDto
//
//		A pointer to an instance of ByteArrayDto. The
//		internal byte array contained in this instance
//		will be copied to the corresponding internal byte
//		array contained in 'destinationBArray'.
//
//		If 'copyDescriptionData' is set to true, the
//		source description data will also be copied to
//		the destination ByteArrayDto instance.
//
//	sourceBArrayLabel			string
//
//		The name or label associated with input parameter
//		'sourceBArray' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "sourceBArray" will be
//		automatically applied.
//
//	destinationBArray			*ByteArrayDto
//
//		A pointer to an instance of ByteArrayDto.
//
//		The internal byte array contained in
//		'sourceBArray' will be copied to the
//		corresponding internal byte array contained in
//		this instance, 'destinationBArray'.
//
//		If 'copyDescriptionData' is set to true, the
//		source description data will also be copied to
//		'destinationBArray'.
//
//	destinationBArrayLabel		string
//
//		The name or label associated with input parameter
//		'destinationBArray' which will be used in error
//		messages returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "destinationBArray"
//		will be automatically applied.
//
//	copyDescriptionData			bool
//
//		If this parameter is set to 'true', the
//		description data contained in 'sourceBArray' will
//		also be copied to the destination ByteArrayDto
//		instance, 'destinationBArray'.
//
//		Description data is contained in two internal
//		member variables:
//
//			ByteArrayDto.Description1
//			ByteArrayDto.Description2
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (bArrayDtoAtom *byteArrayDtoAtom) copy(
	sourceBArray *ByteArrayDto,
	sourceBArrayLabel string,
	destinationBArray *ByteArrayDto,
	destinationBArrayLabel string,
	copyDescriptionData bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bArrayDtoAtom.lock == nil {
		bArrayDtoAtom.lock = new(sync.Mutex)
	}

	bArrayDtoAtom.lock.Lock()

	defer bArrayDtoAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"byteArrayDtoAtom."+
			"copy()",
		"")

	if err != nil {
		return err
	}

	if len(sourceBArrayLabel) == 0 {

		sourceBArrayLabel = "sourceBArray"
	}

	if len(destinationBArrayLabel) == 0 {

		destinationBArrayLabel = "destinationBArray"
	}

	if sourceBArray == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a nil pointer!\n",
			ePrefix.String(),
			sourceBArrayLabel,
			sourceBArrayLabel)

		return err
	}

	if destinationBArray == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a nil pointer!\n",
			ePrefix.String(),
			destinationBArrayLabel,
			destinationBArrayLabel)

		return err
	}

	lenSourceBArray := len(sourceBArray.ByteArray)

	if lenSourceBArray == 0 {

		destinationBArray.ByteArray = nil

	} else {

		destinationBArray.ByteArray =
			make([]byte, lenSourceBArray)

		for i := 0; i < lenSourceBArray; i++ {

			destinationBArray.ByteArray[i] =
				sourceBArray.ByteArray[i]
		}
	}

	if copyDescriptionData == true {

		destinationBArray.Description1 =
			sourceBArray.Description1

		destinationBArray.Description2 =
			sourceBArray.Description2

	}

	return err
}

// deleteLeadingBytes
//
// This method will delete a specified number of bytes
// from the leading edge of an internal byte array
// encapsulated by a ByteArrayDto instance passed as
// input parameter 'bArrayDto'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete pre-existing data values
//		in the internal byte array encapsulated by the
//		ByteArrayDto instance passed as input parameter
//		'bArrayDto'.
//
//	(2)	This method is designed to delete leading bytes
//		from the internal byte array passed by input
//		parameter 'bArrayDto'.
//
//	(3)	If the internal byte array passed by 'bArrayDto'
//		is empty with a length of zero ('0'), this method
//		will take no action, return no error and exit.
//
//		If the number of leading bytes to be deleted
//		('numOfLeadingBytesToDelete') is greater than or
//		equal to the current length of the internal byte
//		array, this method will set the internal byte
//		array to 'nil' (a zero length array) and no error
//		will be returned.
//
//	(4)	If the number of leading bytes to delete
//		('numOfLeadingBytesToDelete') is set to zero
//		('0'), this method will take no action, return no
//		error and exit.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bArrayDto					*ByteArrayDto
//
//		A pointer to an instance of ByteArrayDto.
//
//		This will method will delete bytes from the
//		leading edge of the internal byte array
//		encapsulated by this ByteArrayDto instance.
//
//	bArrayDtoLabel				string
//
//		The name or label associated with input parameter
//		'bArrayDto' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "bArrayDto" will be
//		automatically applied.
//
//	numOfLeadingBytesToDelete	int
//
//		The number of bytes which will be deleted from
//		the leading edge of the internal byte array
//		encapsulated by 'bArrayDto'.
//
//		If the number of leading bytes to delete
//		('numOfLeadingBytesToDelete') is set to zero
//		('0'), this method will take no action, return
//		no error and exit.
//
//		If the internal byte array is empty with a
//		length of zero ('0'), this method will take no
//		action, return no error and exit.
//
//		If the number of leading bytes to be deleted
//		('numOfLeadingBytesToDelete') is greater than or
//		equal to the current length of the internal byte
//		array, this method will set the internal byte
//		array to 'nil' (a zero length array) and no error
//		will be returned.
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (bArrayDtoAtom *byteArrayDtoAtom) deleteLeadingBytes(
	bArrayDto *ByteArrayDto,
	bArrayDtoLabel string,
	numOfLeadingBytesToDelete int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bArrayDtoAtom.lock == nil {
		bArrayDtoAtom.lock = new(sync.Mutex)
	}

	bArrayDtoAtom.lock.Lock()

	defer bArrayDtoAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"byteArrayDtoAtom."+
			"deleteLeadingBytes()",
		"")

	if err != nil {
		return err
	}

	if len(bArrayDtoLabel) == 0 {

		bArrayDtoLabel = "bArrayDto"
	}

	if bArrayDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a nil pointer!\n",
			ePrefix.String(),
			bArrayDtoLabel,
			bArrayDtoLabel)

		return err
	}

	if numOfLeadingBytesToDelete < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfLeadingBytesToDelete' is invalid!\n"+
			"'numOfLeadingBytesToDelete' has a value less than zero ('0')!\n"+
			"numOfLeadingBytesToDelete= '%v'\n",
			ePrefix.String(),
			numOfLeadingBytesToDelete)

		return err
	}

	if numOfLeadingBytesToDelete == 0 {

		return err
	}

	lenOrigByteArray := len(bArrayDto.ByteArray)

	if lenOrigByteArray == 0 {

		return err
	}

	if numOfLeadingBytesToDelete >= lenOrigByteArray {

		bArrayDto.ByteArray = nil

		return err
	}

	bArrayDto.ByteArray =
		bArrayDto.ByteArray[numOfLeadingBytesToDelete:]

	return err
}

// deleteTrailingBytes
//
// This method will delete a specified number of bytes
// from the trailing edge of an internal byte array
// encapsulated by a ByteArrayDto instance passed as
// input parameter 'bArrayDto'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete pre-existing data values
//		in the internal byte array encapsulated by the
//		ByteArrayDto instance passed as input parameter
//		'bArrayDto'.
//
//	(2)	This method is designed to delete trailing bytes
//		from the internal byte array passed by input
//		parameter 'bArrayDto'.
//
//	(3)	If the internal byte array passed by 'bArrayDto'
//		is empty with a length of zero ('0'), this method
//		will take no action, return no error and exit.
//
//		If the number of trailing bytes to be deleted
//		('numOfTrailingBytesToDelete') is greater than or
//		equal to the current length of the internal byte
//		array, this method will set the internal byte
//		array to 'nil' (a zero length array) and no error
//		will be returned.
//
//	(4)	If the number of trailing bytes to delete
//		('numOfTrailingBytesToDelete') is set to zero
//		('0'), this method will take no action, return no
//		error and exit.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bArrayDto					*ByteArrayDto
//
//		A pointer to an instance of ByteArrayDto.
//
//		This will method will delete bytes from the
//		trailing edge of the internal byte array
//		encapsulated by this ByteArrayDto instance.
//
//	bArrayDtoLabel				string
//
//		The name or label associated with input parameter
//		'bArrayDto' which will be used in error messages
//		returned by this method.
//
//		If this parameter is submitted as an empty
//		string, a default value of "bArrayDto" will be
//		automatically applied.
//
//	numOfTrailingBytesToDelete	int
//
//		The number of bytes which will be deleted from
//		the trailing edge of the internal byte array
//		encapsulated by 'bArrayDto'.
//
//		If the number of trailing bytes to delete
//		('numOfTrailingBytesToDelete') is set to zero
//		('0'), this method will take no action, return
//		no error and exit.
//
//		If the internal byte array is empty with a
//		length of zero ('0'), this method will take no
//		action, return no error and exit.
//
//		If the number of trailing bytes to be deleted
//		('numOfTrailingBytesToDelete') is greater than or
//		equal to the current length of the internal byte
//		array, this method will set the internal byte
//		array to 'nil' (a zero length array) and no error
//		will be returned.
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
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing, the
//		returned error Type will encapsulate an
//		appropriate error message. This returned error
//	 	message will incorporate the method chain and
//	 	text passed by input parameter, 'errPrefDto'.
//	 	The 'errPrefDto' text will be prefixed or
//	 	attached to the	beginning of the error message.
func (bArrayDtoAtom *byteArrayDtoAtom) deleteTrailingBytes(
	bArrayDto *ByteArrayDto,
	bArrayDtoLabel string,
	numOfTrailingBytesToDelete int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if bArrayDtoAtom.lock == nil {
		bArrayDtoAtom.lock = new(sync.Mutex)
	}

	bArrayDtoAtom.lock.Lock()

	defer bArrayDtoAtom.lock.Unlock()

	var err error

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"byteArrayDtoAtom."+
			"deleteTrailingBytes()",
		"")

	if err != nil {
		return err
	}

	if len(bArrayDtoLabel) == 0 {

		bArrayDtoLabel = "bArrayDto"
	}

	if bArrayDto == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter '%v' is invalid!\n"+
			"'%v' is a nil pointer!\n",
			ePrefix.String(),
			bArrayDtoLabel,
			bArrayDtoLabel)

		return err
	}

	if numOfTrailingBytesToDelete < 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numOfTrailingBytesToDelete' is invalid!\n"+
			"'numOfTrailingBytesToDelete' has a value less than zero ('0')!\n"+
			"numOfTrailingBytesToDelete= '%v'\n",
			ePrefix.String(),
			numOfTrailingBytesToDelete)

		return err
	}

	lenOfOrigByteArray := len(bArrayDto.ByteArray)

	if lenOfOrigByteArray == 0 {

		return err
	}

	lenOfBytesToKeep :=
		lenOfOrigByteArray -
			numOfTrailingBytesToDelete

	if lenOfBytesToKeep <= 0 {

		bArrayDto.ByteArray = nil

		return err
	}

	bArrayDto.ByteArray =
		bArrayDto.ByteArray[0:lenOfBytesToKeep]

	return err
}

type byteArrayDtoElectron struct {
	lock *sync.Mutex
}

// clear
//
// Receives an instance of ByteArrayDto passed as input
// parameter 'bArrayDto'. This method then proceeds to
// reset each byte in the 'bArrayDto' internal byte array
// to a zero value.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	(1)	This method will delete all pre-existing data
//		values in the internal byte array maintained by
//		the instance of ByteArrayDto passed as input
//		parameter 'bArrayDto'.
//
//	(2)	The length of the internal byte array will
//		remain unchanged, however each element in that
//		internal byte array will be set to a zero ('0')
//		value.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bArrayDto					*ByteArrayDto
//
//		A pointer to an instance of ByteArrayDto.
//
//		This will method will delete all pre-existing
//		data maintained by the internal byte array
//		encapsulated by this ByteArrayDto instance.
//
//		Upon exit, the internal byte array contained in
//		'bArrayDto' will be set to 'nil', a zero length
//		array
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (bArrayDtoElectron *byteArrayDtoElectron) clear(
	bArrayDto *ByteArrayDto) {

	if bArrayDtoElectron.lock == nil {
		bArrayDtoElectron.lock = new(sync.Mutex)
	}

	bArrayDtoElectron.lock.Lock()

	defer bArrayDtoElectron.lock.Unlock()

	if bArrayDto == nil {

		return
	}

	if len(bArrayDto.ByteArray) == 0 {

		return
	}

	clear(bArrayDto.ByteArray)

	return
}

// empty
//
// Receives an instance of ByteArrayDto and sets the
// internal to a value of 'nil'.
//
// Effectively, this resets the internal byte array to
// an empty or zero length array.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing data
//	values in the internal byte array maintained by
//	the instance of ByteArrayDto passed as input
//	parameter 'ByteArrayDto'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	bArrayDto					*ByteArrayDto
//
//		A pointer to an instance of ByteArrayDto.
//
//		This will method will delete all pre-existing
//		data maintained by the internal byte array
//		encapsulated by this ByteArrayDto instance.
//
//		Upon exit, the internal byte array contained in
//		'bArrayDto' will be set to 'nil', a zero length
//		array
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	-- NONE --
func (bArrayDtoElectron *byteArrayDtoElectron) empty(
	bArrayDto *ByteArrayDto) {

	if bArrayDtoElectron.lock == nil {
		bArrayDtoElectron.lock = new(sync.Mutex)
	}

	bArrayDtoElectron.lock.Lock()

	defer bArrayDtoElectron.lock.Unlock()

	if bArrayDto == nil {

		return
	}

	bArrayDto.ByteArray = nil

	bArrayDto.Description1 = ""

	bArrayDto.Description2 = ""

	return
}

package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// Int8ArrayDto
//
// This type is designed to store and transport an
// integer array containing numeric digits.
//
// The integer array consists of an array of type
// int8.
//
// The value of each element in the Int8Array will
// from zero (0) to nine (9) inclusive.
type Int8ArrayDto struct {
	Int8Array []int8
	//	An array of int8 types. Each element of this
	//	array is designed hold a numeric digit with a
	//	value equal to or greater than zero and less
	//	than or equal to nine (0-9 inclusive).

	Description1 string
	//	Optional. A name, label or narrative text used to
	//	describe the current instance of Int8ArrayDto.

	Description2 string
	//	Optional. A name, label or narrative text used to
	//	describe the current instance of Int8ArrayDto.

	lock *sync.Mutex
}

//	CopyIn
//
//	Copies the internal rune array from an incoming
//	instance of Int8ArrayDto ('incomingI8Array') to the
//	internal int8 array of the current Int8ArrayDto
//	instance ('i8ArrayDto').
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The internal member variable data values for the
//	current Int8ArrayDto instance ('i8ArrayDto') will
//	be deleted and overwritten.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingI8Array			*Int8ArrayDto
//
//		A pointer to an instance of Int8ArrayDto. This
//		method will NOT change the values of internal
//		member variable data values contained in this
//		instance.
//
//		All data values in this Int8ArrayDto instance
//		will be copied to current Int8ArrayDto instance
//		('i8ArrayDto').
//
//		If parameter 'incomingI8Array' data values are
//		found to be invalid, an error will be returned.
//
//	 errorPrefix                interface{}
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
//	error
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
func (i8ArrayDto *Int8ArrayDto) CopyIn(
	incomingI8Array *Int8ArrayDto,
	errorPrefix interface{}) error {

	if i8ArrayDto.lock == nil {
		i8ArrayDto.lock = new(sync.Mutex)
	}

	i8ArrayDto.lock.Lock()

	defer i8ArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"Int8ArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return err
	}

	return new(int8ArrayDtoNanobot).copyInt8ArrayDto(
		i8ArrayDto,
		incomingI8Array,
		ePrefix.XCpy(
			"i8ArrayDto<-incomingI8Array"))
}

//	CopyOut
//
//	Returns a deep copy of the current Int8ArrayDto
//	instance.
//
//	If the current Int8ArrayDto instance is found to be
//	invalid, this method will return an error.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	 errorPrefix                interface{}
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
//	deepCopyI8ArrayDto			Int8ArrayDto
//
//		If this method completes successfully and no
//		errors are encountered, this parameter will
//		return a deep copy of the current Int8ArrayDto
//		instance.
//
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
func (i8ArrayDto *Int8ArrayDto) CopyOut(
	errorPrefix interface{}) (
	deepCopyI8ArrayDto Int8ArrayDto,
	err error) {

	if i8ArrayDto.lock == nil {
		i8ArrayDto.lock = new(sync.Mutex)
	}

	i8ArrayDto.lock.Lock()

	defer i8ArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"Int8ArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return deepCopyI8ArrayDto, err
	}

	err = new(int8ArrayDtoNanobot).copyInt8ArrayDto(
		&deepCopyI8ArrayDto,
		i8ArrayDto,
		ePrefix.XCpy(
			"deepCopyI8ArrayDto<-i8ArrayDto"))

	return deepCopyI8ArrayDto, err
}

//	Empty
//
//	Resets all internal member variables for the current
//	instance of Int8ArrayDto to their initial or zero
//	values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method will delete all pre-existing internal
//	member variable data values in the current instance
//	of Int8ArrayDto.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (i8ArrayDto *Int8ArrayDto) Empty() {

	if i8ArrayDto.lock == nil {
		i8ArrayDto.lock = new(sync.Mutex)
	}

	i8ArrayDto.lock.Lock()

	new(int8ArrayDtoAtom).empty(
		i8ArrayDto)

	i8ArrayDto.lock.Unlock()

	i8ArrayDto.lock = nil
}

//	EmptyIntegersArray
//
//	Deletes the internal integers array. Member variable
//	Int8ArrayDto.Int8Array is set to nil.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	For the current instance of Int8ArrayDto, all data
//	values in the internal Integers array
//	(Int8ArrayDto.Int8Array) will be deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (i8ArrayDto *Int8ArrayDto) EmptyIntegersArray() {

	if i8ArrayDto.lock == nil {
		i8ArrayDto.lock = new(sync.Mutex)
	}

	i8ArrayDto.lock.Lock()

	defer i8ArrayDto.lock.Unlock()

	new(int8ArrayDtoElectron).emptyIntsArray(
		i8ArrayDto)

	return
}

//	Equal
//
//	Receives a pointer to an instance of Int8ArrayDto
//	and proceeds to compare all their member variable
//	data values to the corresponding data values in the
//	current Int8ArrayDto instance.
//
//	A boolean flag showing the result of this comparison
//	is returned. If the member variables for both
//	instances are equal in all respects, this flag
//	is set to 'true'. Otherwise, this method returns
//	'false'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingI8Array				*Int8ArrayDto
//
//		A pointer to an instance of Int8ArrayDto.
//		All internal member variable data values
//		contained in this instance will be compared to
//		the corresponding data values in the current
//		Int8ArrayDto instance to determine if they are
//		equivalent in all respects.
//
//	 errorPrefix                interface{}
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
//	areEqual					bool
//
//		This method compares the data values contained in
//		input parameter 'incomingI8Array' to those
//		contained in the current Int8ArrayDto instance.
//
//		If all data elements are found to be equal in all
//		respects, this parameter will be set to 'true'.
//
//		Otherwise, this parameter returns 'false'.
//
//	err							error
//
//		If this method completes successfully, and
//		all data values in parameter 'incomingI8Array'
//		are equivalent to those contained in the current
//		Int8ArrayDto instance, the returned error Type
//		is set equal to 'nil'.
//
//		If errors are encountered during processing or
//		the data elements are not equivalent, the
//		returned error Type will encapsulate an
//		appropriate error message.
//
//		This returned error message will incorporate the
//		method chain and text passed by input parameter,
//		'errorPrefix'. The 'errorPrefix' text will be
//		attached to the beginning of the error message.
func (i8ArrayDto *Int8ArrayDto) Equal(
	incomingI8Array *Int8ArrayDto,
	errorPrefix interface{}) (
	areEqual bool,
	err error) {

	if i8ArrayDto.lock == nil {
		i8ArrayDto.lock = new(sync.Mutex)
	}

	i8ArrayDto.lock.Lock()

	defer i8ArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	areEqual = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"Int8ArrayDto."+
			"CopyIn()",
		"")

	if err != nil {
		return areEqual, err
	}

	areEqual,
		err = new(int8ArrayDtoAtom).equal(
		i8ArrayDto,
		incomingI8Array,
		ePrefix.XCpy(
			"i8ArrayDto!=incomingI8Array"))

	return areEqual, err
}

//	int8ArraysAreEqual
//
//	Compares the integer arrays from the current
//	Int8ArrayDto instance and input parameter
//	'incomingI8Array' to determine if they are
//	equal in all respects.
//
//	The name of the internal member variable being
//	evaluated is:
//
//		Int8ArrayDto.Int8Array
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	incomingI8Array		*Int8ArrayDto
//
//		A pointer to the first of Int8ArrayDto instances.
//
//		The internal integer array data values contained
//		in this instance will be compared to the
//		corresponding integer array data values in
//		the current instance of incomingI8Array to
//		determine if the arrays are equivalent in all
//		respects.
//
//	 errorPrefix		interface{}
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
//	areEqual			bool
//
//		This method will evaluate the integer arrays
//		contained in the input parameter
//		'incomingI8Array' and the current Int8ArrayDto
//		instance to determine if the arrays are
//		equivalent.
//
//		If all array data values are equivalent, this
//		parameter will be set to 'true'.
//
//		Otherwise, this parameter will return 'false'.
//
//	err					error
//
//		If this method completes successfully, and
//		integer array data values in parameter
//		'incomingI8Array' are equivalent to those
//		contained in the current Int8ArrayDto instance,
//		the returned error Type is set equal to 'nil'.
//
//		If errors are encountered during processing or
//		the array data elements are not equivalent, the
//		returned error Type will encapsulate an
//		appropriate error message.
//
//		This returned error message will incorporate the
//		method chain and text passed by input parameter,
//		'errorPrefix'. The 'errorPrefix' text will be
//		attached to the beginning of the error message.
func (i8ArrayDto *Int8ArrayDto) EqualIntegerArrays(
	incomingI8Array *Int8ArrayDto,
	errorPrefix interface{}) (
	areEqual bool,
	err error) {

	if i8ArrayDto.lock == nil {
		i8ArrayDto.lock = new(sync.Mutex)
	}

	i8ArrayDto.lock.Lock()

	defer i8ArrayDto.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	areEqual = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"Int8ArrayDto."+
			"EqualIntegerArrays()",
		"")

	if err != nil {
		return areEqual, err
	}

	areEqual,
		err = new(int8ArrayDtoPreon).int8ArraysAreEqual(
		i8ArrayDto,
		incomingI8Array,
		ePrefix.XCpy(
			"i8ArrayDto!=incomingI8Array"))

	return areEqual, err
}

//	GetCopyInt8Array
//
//	Returns a deep copy of the int8 array of numerical
//	digits encapsulated within the current instance of
//	Int8ArrayDto.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	 NONE
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	[]int8
//
//		An array of numerical digits comprising the
//		Int8Array for the current instance of
//		Int8ArrayDto.
func (i8ArrayDto *Int8ArrayDto) GetCopyInt8Array() []int8 {

	if i8ArrayDto.lock == nil {
		i8ArrayDto.lock = new(sync.Mutex)
	}

	i8ArrayDto.lock.Lock()

	defer i8ArrayDto.lock.Unlock()

	lenI8Array := len(i8ArrayDto.Int8Array)

	if lenI8Array == 0 {

		return nil
	}

	newI8Array := make([]int8, lenI8Array)

	for i := 0; i < lenI8Array; i++ {

		newI8Array[i] = i8ArrayDto.Int8Array[i]
	}

	return newI8Array
}

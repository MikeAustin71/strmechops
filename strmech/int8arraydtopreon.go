package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// int8ArrayDtoPreon - Provides helper methods for type
// Int8ArrayDto.
type int8ArrayDtoPreon struct {
	lock *sync.Mutex
}

//	int8ArraysAreEqual
//
//	Compares the integer arrays from two instances of
//	Int8ArrayDto to determine if they are equal.
//
//	The name of the internal member variable being
//	evaluated is:
//
//		Int8ArrayDto.IntsArray
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	i8ArrayDto1			*Int8ArrayDto
//
//		A pointer to the first of two Int8ArrayDto
//		instances.
//
//		The internal integer array data values contained
//		in this instance will be compared to the
//		corresponding integer array data values in
//		parameter 'i8ArrayDto2' to determine if they are
//		equivalent in all respects.
//
//	i8ArrayDto2			*Int8ArrayDto
//
//		A pointer to the second of two Int8ArrayDto
//		instances.
//
//		The internal integer array data values contained
//		in this instance will be compared to the
//		corresponding integer array data values in
//		parameter 'i8ArrayDto1' to determine if they are
//		equivalent in all respects.
//
//	errPrefDto			*ePref.ErrPrefixDto
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
//	areEqual			bool
//
//		This method will evaluate the integer arrays
//		contained in the input parameters 'i8ArrayDto1'
//		and 'i8ArrayDto2' to determine if they are
//		equivalent.
//
//		If all data values are equivalent, this
//		parameter will be set to 'true'.
//
//		Otherwise, the parameter will return 'false'.
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
func (i8ArrayDtoPreon *int8ArrayDtoPreon) int8ArraysAreEqual(
	i8ArrayDto1 *Int8ArrayDto,
	i8ArrayDto2 *Int8ArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	areEqual bool,
	err error) {

	if i8ArrayDtoPreon.lock == nil {
		i8ArrayDtoPreon.lock = new(sync.Mutex)
	}

	i8ArrayDtoPreon.lock.Lock()

	defer i8ArrayDtoPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	areEqual = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"int8ArrayDtoPreon."+
			"int8ArraysAreEqual()",
		"")

	if err != nil {

		return areEqual, err
	}

	lenI8Array1 := len(i8ArrayDto1.IntsArray)

	if lenI8Array1 != len(i8ArrayDto2.IntsArray) {

		err = fmt.Errorf("%v\n"+
			"Error: Integer Arrays have unequal lengths!\n",
			ePrefix.String())

		return areEqual, err
	}

	for i := 0; i < lenI8Array1; i++ {

		if i8ArrayDto1.IntsArray[i] !=
			i8ArrayDto2.IntsArray[i] {

			err = fmt.Errorf("%v\n"+
				"Error: Integer Arrays have unequal values!\n"+
				"i8ArrayDto1.IntsArray[%v]='%v'\n"+
				"i8ArrayDto2.IntsArray[%v]='%v'\n",
				ePrefix.String(),
				i,
				i8ArrayDto1.IntsArray[i],
				i,
				i8ArrayDto2.IntsArray[i])

			return areEqual, err

		}

	}

	areEqual = true

	return areEqual, err
}

//	isValidI8ArrayDto
//
//	This method receives a pointer to an instance of
//	Int8ArrayDto and proceeds to test the member variable
//	data values. If any values are found to be invalid,
//	this method returns a boolean value of false and an
//	appropriate error message.
//
// ------------------------------------------------------------------------
//
// # Input Parameters
//
//	i8ArrayDto			*Int8ArrayDto
//
//		A pointer to an instance of Int8ArrayDto.
//
//		This method will evaluate the member variables
//		contained in this instance to determine if they
//		are valid in all respects.
//
//	errPrefDto			*ePref.ErrPrefixDto
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
//	isValid				bool
//
//		This method will evaluate the member variables
//		contained in the input parameter 'i8ArrayDto' to
//		determine if they are valid in all respects.
//
//		If all data values are valid, this parameter will
//		be set to 'true'.
//
//		If any data values are found to be invalid, this
//		parameter is set to 'false'.
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
func (i8ArrayDtoPreon *int8ArrayDtoPreon) isValidI8ArrayDto(
	i8ArrayDto *Int8ArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	isValid bool,
	err error) {

	if i8ArrayDtoPreon.lock == nil {
		i8ArrayDtoPreon.lock = new(sync.Mutex)
	}

	i8ArrayDtoPreon.lock.Lock()

	defer i8ArrayDtoPreon.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	isValid = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"int8ArrayDtoPreon."+
			"isValidI8ArrayDto()",
		"")

	if err != nil {

		return isValid, err
	}

	lenI8Array := len(i8ArrayDto.IntsArray)

	if lenI8Array == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter i8ArrayDto.IntsArray is invalid.\n"+
			"i8ArrayDto.IntsArray is empty and has a length of zero!\n",
			ePrefix.String())

		return isValid, err
	}

	for i := 0; i < lenI8Array; i++ {

		if i8ArrayDto.IntsArray[i] < 0 ||
			i8ArrayDto.IntsArray[i] > 9 {

			err = fmt.Errorf("%v\n"+
				"Error: in8 Array Element is invalid!\n"+
				"i8ArrayDto.IntsArray[%v] = %v\n",
				ePrefix.String(),
				i,
				i8ArrayDto.IntsArray[i])

			return isValid, err
		}

	}

	isValid = true

	return isValid, err
}

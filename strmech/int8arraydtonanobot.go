package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// int8ArrayDtoNanobot - Provides helper methods for type
// Int8ArrayDto.
type int8ArrayDtoNanobot struct {
	lock *sync.Mutex
}

//	copyInt8ArrayDto
//
//	Copies all data from input parameter 'sourceI8Array'
//	to input parameter 'destinationI8Array'.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	The pre-existing data fields for input parameter
//	'destinationI8Array' will be overwritten and deleted.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	destinationI8Array		*Int8ArrayDto
//
//		A pointer to an instance of Int8ArrayDto. All
//		data contained in the internal member variables
//		of input parameter 'sourceI8Array' will be copied
//		to the corresponding member variables of this
//		input parameter, 'destinationI8Array'.
//
//		If this method completes successfully, all member
//		data variable data values encapsulated in
//		'destinationI8Array' will be identical to those
//		contained in input parameter, 'sourceI8Array'.
//
//		Be advised that the pre-existing data fields in
//		input parameter 'destinationI8Array' will be
//		overwritten and deleted.
//
//
//	sourceI8Array			*Int8ArrayDto
//
//		A pointer to an instance of Int8ArrayDto.
//
//		All data contained in the member variables of
//		this Int8ArrayDto instance will be copied to
//		corresponding member variables contained within
//		input parameter 'destinationI8Array'.
//
//		The original member variable data values
//		encapsulated in 'sourceI8Array' will remain
//		unchanged and will NOT be modified.
//
//		If 'sourceI8Array' contains invalid member data
//		variables, this method will return an error.
//
//
//	errPrefDto				*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling
//		method or methods listed as a function chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// # Return Values
//
//	error
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
func (i8ArrayNanobot *int8ArrayDtoNanobot) copyInt8ArrayDto(
	destinationI8Array *Int8ArrayDto,
	sourceI8Array *Int8ArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if i8ArrayNanobot.lock == nil {
		i8ArrayNanobot.lock = new(sync.Mutex)
	}

	i8ArrayNanobot.lock.Lock()

	defer i8ArrayNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"int8ArrayDtoNanobot."+
			"copyInt8ArrayDto()",
		"")

	if err != nil {

		return err

	}

	if sourceI8Array == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'sourceI8Array' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if destinationI8Array == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'destinationRunesDto' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return err
	}

	_,
		err = new(int8ArrayDtoPreon).isValidI8ArrayDto(
		sourceI8Array,
		ePrefix.XCpy(
			"sourceI8Array"))

	if err != nil {

		return err
	}

	new(int8ArrayDtoAtom).empty(
		destinationI8Array)

	lenSourceArray := len(sourceI8Array.significand)

	destinationI8Array.significand = make([]int8, lenSourceArray)

	for i := 0; i < lenSourceArray; i++ {
		destinationI8Array.significand[i] =
			sourceI8Array.significand[i]
	}

	destinationI8Array.exponent =
		sourceI8Array.exponent

	destinationI8Array.numberSign =
		sourceI8Array.numberSign

	destinationI8Array.Description1 =
		sourceI8Array.Description1

	destinationI8Array.Description2 =
		sourceI8Array.Description2

	return err
}

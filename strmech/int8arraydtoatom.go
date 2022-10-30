package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// int8ArrayDtoAtom - Provides helper methods for type
// Int8ArrayDto.
type int8ArrayDtoAtom struct {
	lock *sync.Mutex
}

//	empty
//
//	Receives a pointer to an instance of Int8ArrayDto and
//	proceeds to reset the data values for member
//	variables to their initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values contained in
//	input parameter 'i8ArrayDto' will be deleted and
//	reset to their initial or zero values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	i8ArrayDto			*Int8ArrayDto
//
//		A pointer to an instance of Int8ArrayDto. All
//		internal member variable data values contained in
//		this instance will be deleted and reset to their
//		initial or zero values.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (i8ArrayDtoAtom *int8ArrayDtoAtom) empty(
	i8ArrayDto *Int8ArrayDto) {

	if i8ArrayDtoAtom.lock == nil {
		i8ArrayDtoAtom.lock = new(sync.Mutex)
	}

	i8ArrayDtoAtom.lock.Lock()

	defer i8ArrayDtoAtom.lock.Unlock()

	if i8ArrayDto == nil {

		return
	}

	new(int8ArrayDtoElectron).emptyIntsArray(
		i8ArrayDto)

	i8ArrayDto.Description1 = ""

	i8ArrayDto.Description2 = ""

	return
}

//	equal
//
//	Receives a pointer to two instances of Int8ArrayDto
//	and proceeds to compare all their member variable
//	data values in order to determine if they are
//	equivalent.
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
//	i8ArrayDto1			*Int8ArrayDto
//
//		A pointer to the first of two Int8ArrayDto
//		instances. All internal member variable data
//		values contained in this instance will be
//		compared to the corresponding data values in
//		parameter 'i8ArrayDto2' to determine if they are
//		equivalent in all respects.
//
//	i8ArrayDto2			*Int8ArrayDto
//
//		A pointer to the second of two Int8ArrayDto
//		instances. All internal member variable data
//		values contained in this instance will be
//		compared to the corresponding data values in
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
//		This method compares the data values contained in
//		two instances of Int8ArrayDto, 'i8ArrayDto1' and
//		'i8ArrayDto2'.
//
//		If all data elements are found to be equal in all
//		respects, this parameter will be set to 'true'.
//
//		Otherwise, this parameter returns 'false'.
//
//	err						error
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
func (i8ArrayDtoAtom *int8ArrayDtoAtom) equal(
	i8ArrayDto1 *Int8ArrayDto,
	i8ArrayDto2 *Int8ArrayDto,
	errPrefDto *ePref.ErrPrefixDto) (
	areEqual bool,
	err error) {

	if i8ArrayDtoAtom.lock == nil {
		i8ArrayDtoAtom.lock = new(sync.Mutex)
	}

	i8ArrayDtoAtom.lock.Lock()

	defer i8ArrayDtoAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	areEqual = false

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"int8ArrayDtoAtom."+
			"equal()",
		"")

	if err != nil {

		return areEqual, err
	}

	if i8ArrayDto1 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'i8ArrayDto1' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return areEqual, err
	}

	if i8ArrayDto2 == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'i8ArrayDto2' is "+
			"a nil pointer!\n",
			ePrefix.String())

		return areEqual, err
	}

	areEqual,
		err = new(int8ArrayDtoPreon).int8ArraysAreEqual(
		i8ArrayDto1,
		i8ArrayDto2,
		ePrefix.XCpy(
			"i8ArrayDto1!=i8ArrayDto2"))

	if err != nil {
		return areEqual, err
	}

	if i8ArrayDto1.Description1 !=
		i8ArrayDto2.Description1 {

		err = fmt.Errorf("%v\n"+
			"Error: Description1 values are NOT Equal!\n",
			ePrefix.String())

		return areEqual, err

	}

	if i8ArrayDto1.Description2 !=
		i8ArrayDto2.Description2 {

		err = fmt.Errorf("%v\n"+
			"Error: Description2 values are NOT Equal!\n",
			ePrefix.String())

		return areEqual, err

	}

	areEqual = true

	return areEqual, err
}

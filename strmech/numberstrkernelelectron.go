package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelElectron - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelElectron struct {
	lock *sync.Mutex
}

//	empty
//
//	Receives a pointer to an instance of NumberStrKernel and
//	proceeds to reset the data values for member variables
//	to their initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	All the member variable data values contained in input
//	parameter 'numStrKernel' will be deleted and reset to
//	their zero values.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. All
//		the internal member variables contained in this
//		instance will be deleted and reset to their zero
//		values.
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	NONE
func (numStrKernelElectron *numberStrKernelElectron) empty(
	numStrKernel *NumberStrKernel) {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	if numStrKernel == nil {
		return
	}

	numStrKernel.integerDigits.Empty()

	numStrKernel.integerDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	numStrKernel.fractionalDigits.Empty()

	numStrKernel.fractionalDigits.charSearchType =
		CharSearchType.LinearTargetStartingIndex()

	numStrKernel.numberValueType =
		NumValType.None()

	numStrKernel.numberSign = NumSignVal.None()

	numStrKernel.isNonZeroValue = false

	numStrKernel.numStrFormatSpec.Empty()
}

// equal - Receives a pointer to two instances of
// NumberStrKernel and proceeds to compare their member
// variables in order to determine if they are equivalent.
//
// A boolean flag showing the result of this comparison is
// returned. If the member variables for both instances are equal
// in all respects, this flag is set to 'true'. Otherwise, this
// method returns 'false'.
func (numStrKernelElectron *numberStrKernelElectron) equal(
	numStrKernel1 *NumberStrKernel,
	numStrKernel2 *NumberStrKernel) bool {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	if numStrKernel1 == nil ||
		numStrKernel2 == nil {

		return false
	}

	if !numStrKernel1.integerDigits.Equal(
		&numStrKernel2.integerDigits) {

		return false
	}

	if !numStrKernel1.fractionalDigits.Equal(
		&numStrKernel2.fractionalDigits) {

		return false
	}

	if numStrKernel1.numberValueType !=
		numStrKernel2.numberValueType {

		return false
	}

	if numStrKernel1.numberSign !=
		numStrKernel2.numberSign {

		return false
	}

	if numStrKernel1.isNonZeroValue !=
		numStrKernel2.isNonZeroValue {

		return false
	}

	if !numStrKernel1.numStrFormatSpec.Equal(
		&numStrKernel2.numStrFormatSpec) {

		return false
	}

	return true
}

//	equalizeNumStrDigitsLengths
//
//	Receives pointers to two instances of NumberStrKernel,
//	'numStrKernel01' and 'numStrKernel02'. This method
//	will ensure that both the integer digit arrays and
//	fractional digit arrays contained in both instances
//	have equal array lengths.
//
//	If the integer arrays do not have equal array
//	lengths, leading zero characters ('0') will be added
//	to configure their array lengths as equal.
//
//	If the fractional arrays do not have equal array
//	lengths, trailing zero characters ('0') will be added
//	to configure their array lengths as equal.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	This method may modify the internal integer digit and
//	fractional digit rune arrays contained within input
//	parameters 'numStrKernel01' and 'numStrKernel02'.
//
// ----------------------------------------------------------------
//
//	# Input Parameters
//
//	numStrKernel01				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal integer and fractional array lengths for
//		this instance will be compared to those of input
//		parameter, 'numStrKernel02'.
//
//		If the 'numStrKernel01' integer digit array has a
//		shorter array length, leading zero characters
//		('0') will be added to the 'numStrKernel01'
//		integer array to achieve array length
//		equivalency.
//
//		If the 'numStrKernel01' fractional digit array
//		has a shorter array length, trailing zero
//		characters ('0') will be added to the
//		'numStrKernel01' fractional array to achieve
//		array length equivalency.
//
//
//	numStrKernel02				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The
//		internal integer and fractional array lengths for
//		this instance will be compared to those of input
//		parameter, 'numStrKernel01'.
//
//		If the 'numStrKernel02' integer digit array has a
//		shorter array length, leading zero characters
//		('0') will be added to the 'numStrKernel02'
//		integer array to achieve array length
//		equivalency.
//
//		If the 'numStrKernel02' fractional digit array
//		has a shorter array length, trailing zero
//		characters ('0') will be added to the
//		'numStrKernel02' fractional array to achieve
//		array length equivalency.
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
func (numStrKernelElectron *numberStrKernelElectron) equalizeNumStrDigitsLengths(
	numStrKernel01 *NumberStrKernel,
	numStrKernel02 *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelElectron."+
			"equalizeNumStrDigitsLengths()",
		"")

	if err != nil {

		return err
	}

	if numStrKernel01 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel01' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if numStrKernel02 == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel02' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	lenArray01 :=
		len(numStrKernel01.integerDigits.CharsArray)

	lenArray02 :=
		len(numStrKernel02.integerDigits.CharsArray)

	numStrKernelQuark := numberStrKernelQuark{}

	if lenArray01 != lenArray02 {

		err = numStrKernelQuark.
			equalizeNStrIntDigitsLengths(
				numStrKernel01,
				numStrKernel02,
				ePrefix.XCpy(
					"numStrKernel01 - numStrKernel02"))

		if err != nil {

			return err
		}

		lenArray01 =
			len(numStrKernel01.integerDigits.CharsArray)

		lenArray02 =
			len(numStrKernel02.integerDigits.CharsArray)

		if lenArray01 != lenArray02 {

			err = fmt.Errorf("%v\n"+
				"SYSTEM ERROR: Integer Digit Array Lengths\n"+
				"ARE NOT EQUAL!\n"+
				"lenIntDigitArray01 = '%v'\n"+
				"lenIntDigitArray02 = '%v'\n",
				ePrefix.String(),
				lenArray01,
				lenArray02)

			return err
		}
	}

	lenArray01 =
		len(numStrKernel01.fractionalDigits.CharsArray)

	lenArray02 =
		len(numStrKernel02.fractionalDigits.CharsArray)

	if lenArray01 != lenArray02 {

		err = numStrKernelQuark.
			equalizeNStrFracDigitsLengths(
				numStrKernel01,
				numStrKernel02,
				ePrefix.XCpy(
					"numStrKernel01 - numStrKernel02"))

		if err != nil {

			return err
		}

		lenArray01 =
			len(numStrKernel01.fractionalDigits.CharsArray)

		lenArray02 =
			len(numStrKernel02.fractionalDigits.CharsArray)

		if lenArray01 != lenArray02 {

			err = fmt.Errorf("%v\n"+
				"SYSTEM ERROR: Fractional Digit Array Lengths\n"+
				"ARE NOT EQUAL!\n"+
				"lenFracDigitArray01 = '%v'\n"+
				"lenFracDigitArray02 = '%v'\n",
				ePrefix.String(),
				lenArray01,
				lenArray02)

			return err
		}
	}

	return err
}

//	getSetIsNonZeroValue
//
//	Receives a pointer to an instance of
//	NumberStrKernel and proceeds to determine if
//	the numeric value contained within this
//	instance is non-zero.
//
//	If this method returns 'true', it means that
//	the numeric value of the NumberStrKernel
//	instance is non-zero. A non-zero numeric value
//	signals that the numeric value is less than or
//	greater than zero (0).
//
//	If this method returns 'false' it means that the
//	numeric value of the NumberStrKernel instance
//	is zero ('0').
//
//	If the NumberStrKernel is found to have a zero,
//	this method will reset the Number Sign and Numeric
//	Value Type accordingly.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		method will determine whether the numeric value
//		of this instance is non-zero.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	isNonZeroValue				bool
//
//		This method will examine the NumberStrKernel
//		instance passed as input parameter 'numStrKernel'
//		to determine if it contains a non-zero numeric
//		value.
//
//		If this method returns 'true', it means that
//		the numeric value of 'numStrKernel' is non-zero.
//		A non-zero numeric value signals that the numeric
//		value is less than or greater than zero (0).
//
//		If this method returns 'false' it means that the
//		numeric value of 'numStrKernel' is zero ('0').
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If errors
//		are	encountered during processing, the returned
//		error Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelElectron *numberStrKernelElectron) getSetIsNonZeroValue(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	isNonZeroValue bool,
	err error) {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelElectron."+
			"getSetIsNonZeroValue()",
		"")

	if err != nil {

		return isNonZeroValue, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return isNonZeroValue, err
	}

	lenArray := numStrKernel.
		integerDigits.
		GetRuneArrayLength()

	numStrKernelQuark := numberStrKernelQuark{}

	for i := 0; i < lenArray; i++ {

		if numStrKernel.integerDigits.CharsArray[i] > '0' &&
			numStrKernel.integerDigits.CharsArray[i] <= '9' {

			numStrKernel.isNonZeroValue = true

			isNonZeroValue = true

			_,
				_ = numStrKernelQuark.getSetNumValueType(
				numStrKernel,
				nil)

			return isNonZeroValue, err
		}
	}

	lenArray = numStrKernel.
		fractionalDigits.
		GetRuneArrayLength()

	for i := 0; i < lenArray; i++ {

		if numStrKernel.fractionalDigits.CharsArray[i] > '0' &&
			numStrKernel.fractionalDigits.CharsArray[i] <= '9' {

			numStrKernel.isNonZeroValue = true

			isNonZeroValue = true

			_,
				_ = numStrKernelQuark.getSetNumValueType(
				numStrKernel,
				nil)

			return isNonZeroValue, err
		}
	}

	numStrKernel.isNonZeroValue = false

	isNonZeroValue = false

	_,
		_ = numStrKernelQuark.getSetNumValueType(
		numStrKernel,
		nil)

	numStrKernel.numberSign = NumSignVal.Zero()

	return isNonZeroValue, err
}

//	rationalizeFractionalIntegerDigits
//
//	If fractional digits are present in this instance of
//	NumberStrKernel, this method will ensure that integer
//	digits are also present.
//
//	If fractional digits are present and no integer digits
//	are found, this method will configure a zero ('0') in
//	the integer digits rune array.
//
//	Example:
//
//		.752 will be converted to 0.752
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. If
//		the fractional digits array contains numeric
//		digits and the integer digits array contains
//		zero digits, the integer digits array will be
//		configured with a zero ('0').
//
//			Example:
//				.752 will be converted to 0.752
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
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
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If errors
//		are	encountered during processing, the returned
//		error Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelElectron *numberStrKernelElectron) rationalizeFractionalIntegerDigits(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelElectron."+
			"rationalizeFractionalIntegerDigits()",
		"")

	if err != nil {

		return err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if numStrKernel.fractionalDigits.GetRuneArrayLength() == 0 {
		return err
	}

	// Fractional Digits exist!
	if numStrKernel.integerDigits.GetRuneArrayLength() == 0 {

		numStrKernel.integerDigits.CharsArray =
			make([]rune, 1)

		numStrKernel.integerDigits.CharsArray[0] = '0'
	}

	return err
}

//	setUninitializedKernelToZero
//
//	Receives a pointer to an instance of NumberStrKernel
//	and proceeds to test the internal member variables
//	for uninitialized values.
//
//	If the NumberStrKernel numeric value is uninitialized,
//	this method will set the numeric value to zero.
//
//	An uninitialized numeric value is defined as the case
//	where both integer digits and fractional digits rune
//	arrays are both empty and equal to nil. In this case
//	this method will configure the integer rune array with
//	a single zero ('0') character.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
//	Modification of internal data values for input
//	parameter 'numStrKernel' may occur.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		method will be examined to determine whether the
//		numeric value is uninitialized. An uninitialized
//		numeric value is defined as a case where both the
//		integer and fractional digits rune arrays are
//		equal to zero. In this case, this method will set
//		the integer array to a single zero ('0') character.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error messages.
//		Usually, it contains the name of the calling method
//		or methods listed as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	err							error
//
//		If this method completes successfully, this
//		returned error Type is set equal to 'nil'. If errors
//		are	encountered during processing, the returned
//		error Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelElectron *numberStrKernelElectron) setUninitializedKernelToZero(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelElectron."+
			"setUninitializedKernelToZero()",
		"")

	if err != nil {

		return err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if numStrKernel.integerDigits.GetRuneArrayLength() == 0 &&
		numStrKernel.fractionalDigits.GetRuneArrayLength() == 0 {

		numStrKernel.integerDigits.CharsArray =
			make([]rune, 1)

		numStrKernel.integerDigits.CharsArray[0] = '0'

	}

	return err
}

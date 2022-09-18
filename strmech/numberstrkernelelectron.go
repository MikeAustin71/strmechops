package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"sync"
)

// numberStrKernelElectron - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelElectron struct {
	lock *sync.Mutex
}

//	convertKernelToBigInt
//
//	Converts an instance of NumberStrKernel to an integer value of
//	type *big.Int.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The numeric
//		value contained in this instance will be converted to
//		an integer of type *big.Int.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this parameter
//		to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	*big.Int
//
//		If this method completes successfully, the numeric
//		value represented by the current instance of
//		NumberStrKernel will be returned as a type *big.Int.
//
//	error
//
//		If this method completes successfully, this returned error
//		Type is set equal to 'nil'. If errors are encountered during
//		processing, the returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value for input
//		parameter 'errPrefDto' (error prefix) will be prefixed or
//		attached at the beginning of the error message.
func (numStrKernelElectron *numberStrKernelElectron) convertKernelToBigInt(
	numStrKernel *NumberStrKernel,
	roundingType NumberRoundingType,
	errPrefDto *ePref.ErrPrefixDto) (
	*big.Int,
	error) {

	if numStrKernelElectron.lock == nil {
		numStrKernelElectron.lock = new(sync.Mutex)
	}

	numStrKernelElectron.lock.Lock()

	defer numStrKernelElectron.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	bigIntValue := big.NewInt(0)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelElectron."+
			"convertKernelToBigInt()",
		"")

	if err != nil {

		return bigIntValue, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return bigIntValue, err
	}

	if !roundingType.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrRoundingSpec Rounding Type' is invalid!\n"+
			"'roundingType' string  value = '%v'\n"+
			"'roundingType' integer value = '%v'\n",
			ePrefix.String(),
			roundingType.String(),
			roundingType.XValueInt())

		return bigIntValue, err

	}

	var ok bool

	if roundingType == NumRoundType.NoRounding() {

		_,
			ok = bigIntValue.SetString(
			numStrKernel.integerDigits.GetCharacterString(),
			10)

		if !ok {
			err = fmt.Errorf("%v\n"+
				"Error Converting Integer string to *big.Int!\n"+
				"The following integerDigits string generated an error.\n"+
				"numStrKernel.integerDigits = '%v'\n",
				ePrefix.String(),
				numStrKernel.integerDigits.GetCharacterString())
		}

		return bigIntValue, err
	}

	var copyNStrKernel NumberStrKernel

	copyNStrKernel,
		err = numStrKernel.CopyOut(
		ePrefix.XCpy(
			"copyNStrKernel<-numStrKernel"))

	if err != nil {

		return bigIntValue, err

	}

	var numStrRoundingSpec NumStrRoundingSpec

	numStrRoundingSpec,
		err =
		new(NumStrRoundingSpec).NewRoundingSpec(
			roundingType,
			0,
			ePrefix)

	if err != nil {
		return bigIntValue, err
	}

	err = new(numStrMathRoundingNanobot).roundNumStrKernel(
		&copyNStrKernel,
		numStrRoundingSpec,
		ePrefix)

	if err != nil {
		return bigIntValue, err
	}

	_,
		ok = bigIntValue.SetString(
		copyNStrKernel.integerDigits.GetCharacterString(),
		10)

	if !ok {
		err = fmt.Errorf("%v\n"+
			"Error Converting Rounded Integer string to *big.Int!\n"+
			"The following integerDigits string generated an error.\n"+
			"numStrKernel.integerDigits = '%v'\n",
			ePrefix.String(),
			copyNStrKernel.integerDigits.GetCharacterString())
	}

	return bigIntValue, err
}

// empty - Receives a pointer to an instance of
// NumberStrKernel and proceeds to reset the data values
// for member variables to their initial or zero values.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// All the member variable data values contained in input parameter
// 'numStrKernel' will be deleted and reset to their zero
// values.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	numStrKernel               *NumberStrKernel
//	   - A pointer to an instance of NumberStrKernel. All
//	     the internal member variables contained in this instance
//	     will be deleted and reset to their zero values.
//
// ------------------------------------------------------------------------
//
// Return Values
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

	numStrKernel.fractionalDigits.Empty()

	numStrKernel.numericValueType =
		NumValType.None()

	numStrKernel.numberSign = NumSignVal.None()

	numStrKernel.isNonZeroValue = false
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

	if numStrKernel1.numericValueType !=
		numStrKernel2.numericValueType {

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

	return true
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelQuark - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelQuark struct {
	lock *sync.Mutex
}

//	getSetNumValueType
//
//	Sets and returns the current NumericValueType for the
//	instance of NumberStrKernel passed as an input
//	parameter.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. This
//		method will examine the internal member variables
//		contained in this instance and set the correct
//		value for Numeric Value Type.
//
//		NumericValueType is an enumeration value specifying
//		the type of numeric value contained in the
//		'numStrKernel' instance.
//
//		Possible NumericValueType enumeration values are
//		listed as follows:
//			NumValType.None()
//			NumValType.FloatingPoint()
//			NumValType.Integer()
//
//		The internal variable contained in 'numStrKernel'
//		which will be configured is:
//
//			NumberStrKernel.numericValueType
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
// ------------------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, this returned
//		error Type is set equal to 'nil'. If errors are
//		encountered during processing, the returned error
//		Type will encapsulate an error message.
//
//		If an error message is returned, the text value for
//		input parameter 'errPrefDto' (error prefix) will be
//		prefixed or attached at the beginning of the error
//		message.
func (numStrKernelQuark *numberStrKernelQuark) getSetNumValueType(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	NumericValueType,
	error) {

	if numStrKernelQuark.lock == nil {
		numStrKernelQuark.lock = new(sync.Mutex)
	}

	numStrKernelQuark.lock.Lock()

	defer numStrKernelQuark.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	newNumericValueType := NumValType.None()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"negNumSearchSpecAtom."+
			"emptyIntegerDigits()",
		"")

	if err != nil {

		return newNumericValueType, err
	}

	if numStrKernel == nil {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return newNumericValueType, err
	}

	lenIntegerDigits :=
		numStrKernel.integerDigits.GetRuneArrayLength()

	lenFracDigits :=
		numStrKernel.fractionalDigits.GetRuneArrayLength()

	if lenIntegerDigits == 0 &&
		lenFracDigits == 0 {
		numStrKernel.numericValueType = newNumericValueType

		return newNumericValueType, err
	}

	if lenIntegerDigits > 0 &&
		lenFracDigits == 0 {

		newNumericValueType = NumValType.Integer()

		return newNumericValueType, err
	}

	if lenFracDigits > 0 {

		newNumericValueType = NumValType.FloatingPoint()

		return newNumericValueType, err
	}

	newNumericValueType = NumValType.None()

	return newNumericValueType, err
}

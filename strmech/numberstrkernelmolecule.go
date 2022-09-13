package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

// numberStrKernelMolecule - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelMolecule struct {
	lock *sync.Mutex
}

// convertIntToKernel
//
// Receives an empty interface which is assumed to be an
// integer numeric value configured as one of the following
// types:
//
//	int8
//	int16
//	int32
//	int	(equivalent to int32)
//	int64
//
// This integer numeric value is then converted to a
// type of 'NumberStrKernel' and returned to the calling
// function.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	intNumericValue	interface{}
//
//		This empty interface is assumed to encapsulate an integer
//		numeric value comprised of one of the following types:
//		This numeric value will be used to populate the instance
//		of NumberStrKernel passed by parameter, 'numStrKernel'.
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//
//		If the object passed by this empty interface is NOT one of
//		the types listed above, an error will be returned.
//
//		No Data Validation is performed on 'intNumericValue'. It
//		is assumed to be one of the valid types identified above.
//
//	numberSign					NumericSignValueType
//
//		The Number Sign is specified by means of a
//		NumericSignValueType enumeration value.
//
//		Possible values are listed as follows:
//
//			NumSignVal.None()     = -2 - Infer From Number
//			NumSignVal.Negative() = -1 - Valid Value
//			NumSignVal.Zero()     =  0 - Valid Value
//			NumSignVal.Positive() =  1 - Valid Value
//
//		If 'numberSign' is set to 'NumSignVal.None()', the
//		number sign will be inferred from the integer numeric
//		value and generated as a default.
//
//		If 'numberSign' is set to NumSignVal.Negative() or
//		NumSignVal.Positive(), it will override the default
//		number sign associated with 'intNumericValue'.
//
//		If 'numberSign' is set to an invalid value or if
//		'numberSign' is set to NumSignVal.Zero() for a
//		non-zero numeric value, an error will be returned.
func (numStrKernelMolecule *numberStrKernelMolecule) convertIntToKernel(
	numStrKernel *NumberStrKernel,
	numberSign NumericSignValueType,
	intNumericValue interface{},
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelMolecule.lock == nil {
		numStrKernelMolecule.lock = new(sync.Mutex)
	}

	numStrKernelMolecule.lock.Lock()

	defer numStrKernelMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelNanobot."+
			"convertIntToKernel()",
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

	if !numberSign.XIsValid() {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numberSign' is invalid!\n"+
			"'numberSign' string value  = '%v\n"+
			"'numberSign' integer value = '%v\n",
			ePrefix.String(),
			numberSign.String(),
			numberSign.XValueInt())

		return err

	}

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		intNumericValue)

	decimalSeparatorSpec := DecimalSeparatorSpec{}
	numParsingTerminators := RuneArrayCollection{}
	negativeNumSearchSpecs := NegNumSearchSpecCollection{}

	err = decimalSeparatorSpec.SetDecimalSeparatorStr(
		".",
		ePrefix.XCpy("Decimal Point '.'"))

	if err != nil {
		return err
	}

	err = negativeNumSearchSpecs.AddLeadingNegNumSearchStr(
		"-",
		ePrefix.XCpy("Leading Minus Sign '-'"))

	if err != nil {
		return err
	}

	runeArrayDto := RuneArrayDto{
		CharsArray:   []rune(numberStr),
		Description1: "",
		Description2: "",
	}

	_,
		*numStrKernel,
		err = new(numStrBuilderElectron).extractNumRunes(
		runeArrayDto,
		"integerNumberStr",
		0,
		-1,
		negativeNumSearchSpecs,
		decimalSeparatorSpec,
		numParsingTerminators,
		false,
		ePrefix.XCpy(
			numberStr))

	if err != nil {
		return err
	}

	if numStrKernel.numberSign == NumSignVal.Zero() {

		return err
	}

	// Override number sign
	if numberSign == NumSignVal.Positive() ||
		numberSign == NumSignVal.Negative() {

		numStrKernel.numberSign = numberSign
	}

	return err
}

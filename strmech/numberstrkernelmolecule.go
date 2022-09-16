package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"sync"
)

// numberStrKernelMolecule - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelMolecule struct {
	lock *sync.Mutex
}

//	convertSignedIntToKernel
//
//	Receives an empty interface which is assumed to be an
//	integer numeric value configured as one of the following
//	types:
//
//		int8
//		int16
//		int32
//		int	(equivalent to int32)
//		int64
//
//	This integer numeric value is then converted to a
//	type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	signedIntNumericValue		interface{}
//
//		This empty interface is assumed to encapsulate a signed
//		integer	numeric value comprised of one of the following
//		types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//
//		This numeric value will be used to populate the instance
//		of NumberStrKernel passed by parameter, 'numStrKernel'.
//
//		If the object passed by this empty interface is NOT one
//		of the types listed above, an error will be returned.
//
//	errPrefDto          *ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
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
func (numStrKernelMolecule *numberStrKernelMolecule) convertSignedIntToKernel(
	numStrKernel *NumberStrKernel,
	signedIntNumericValue interface{},
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
		"numberStrKernelMolecule."+
			"convertSignedIntToKernel()",
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

	switch signedIntNumericValue.(type) {

	case int8, int16, int, int32, int64:

		goto intNumericValueProcessing

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'signedIntNumericValue' is an invalid type!\n"+
			"Only signed integer types are suppored.\n"+
			"'signedIntNumericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", signedIntNumericValue))

		return err

	}

intNumericValueProcessing:

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		signedIntNumericValue)

	var searchResults CharSearchNumStrParseResultsDto
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

	searchResults,
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

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	return err
}

//	convertBigIntToKernel
//
//	Receives a parameter of type empty interface which is
//	assumed to be a type *big.Int. If the empty interface
//	is NOT convertible to a type *big.Int an error will be
//	returned.
//
//	This *big.Int integer numeric value is then converted to
//	a type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	bigFloatValue				*big.Float
//
//		The numeric value this method will use to configure
//		parameter 'numStrKernel'.
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
func (numStrKernelMolecule *numberStrKernelMolecule) convertBigFloatToKernel(
	numStrKernel *NumberStrKernel,
	bigFloatValue *big.Float,
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
		"numberStrKernelMolecule."+
			"convertBigFloatToKernel()",
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

	new(numberStrKernelElectron).empty(
		numStrKernel)

	precision := bigFloatValue.Prec()

	if precision > uint(math.MaxInt) {

		err = fmt.Errorf("%v\n"+
			"Error: Precision Out-Of-Range!\n"+
			"The precision specified by parameter 'bigFloatValue'\n"+
			"exceeds the maximum value for an integer and therefore\n"+
			"cannot be converted to a string value.\n"+
			"'bigFloatValue' precision = '%v'\n"+
			"Maximum allowed conversion precision = '%v'\n",
			ePrefix.String(),
			precision,
			math.MaxInt)

		return err
	}

	numberStr := fmt.Sprintf("%v",
		bigFloatValue.Text('f', int(precision)))

	var searchResults CharSearchNumStrParseResultsDto
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

	searchResults,
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

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	return err
}

//	convertBigIntToKernel
//
//	Receives a parameter of type empty interface which is
//	assumed to be a type *big.Int. If the empty interface
//	is NOT convertible to a type *big.Int an error will be
//	returned.
//
//	This *big.Int integer numeric value is then converted to
//	a type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	bigIntValue					*big.Int
//
//		The numeric value this method will use to configure
//		parameter 'numStrKernel'.
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
func (numStrKernelMolecule *numberStrKernelMolecule) convertBigIntToKernel(
	numStrKernel *NumberStrKernel,
	bigIntValue *big.Int,
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
		"numberStrKernelMolecule."+
			"convertBigIntToKernel()",
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

	if bigIntValue == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'bigIntValue' is a nil pointer!\n",
			ePrefix.String())

		return err

	}

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		bigIntValue.Text(10))

	var searchResults CharSearchNumStrParseResultsDto
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

	searchResults,
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

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	return err
}

//	convertFloatToKernel
//
//	Receives an empty interface which is assumed to be a
//	floating point numeric value configured as one of the
//	following types:
//
//		float32
//		float64
//
//	This floating point numeric value is then converted to
//	a type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	floatNumericValue			interface{}
//
//		This empty interface is assumed to encapsulate a
//		floating point numeric value comprised of one of the
//		following types:
//
//		float32
//		float64
//
//		This numeric value will be used to populate the instance
//		of NumberStrKernel passed by parameter, 'numStrKernel'.
//
//		If the object passed by this empty interface is NOT one
//		of the types listed above, an error will be returned.
//
//	errPrefDto          *ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string which is
//		included in all returned error messages. Usually, it
//		contains the name of the calling method or methods listed
//		as a function chain.
//
//		If no error prefix information is needed, set this
//		parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref' software
//		package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
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
func (numStrKernelMolecule *numberStrKernelMolecule) convertFloatToKernel(
	numStrKernel *NumberStrKernel,
	floatNumericValue interface{},
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
		"numberStrKernelMolecule."+
			"convertFloatToKernel()",
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

	switch floatNumericValue.(type) {

	case float32, float64:

		goto floatNumericValueProcessing

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'floatNumericValue' is an invalid type!\n"+
			"Only floating point types are suppored (float32, float64).\n"+
			"'floatNumericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", floatNumericValue))

		return err

	}

floatNumericValueProcessing:

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		floatNumericValue)

	var searchResults CharSearchNumStrParseResultsDto
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

	searchResults,
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

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	return err
}

//	convertUnsignedInteger
//
//	Receives an empty interface which is assumed to be an
//	unsigned integer numeric value configured as one of the
//	following types:
//
//		uint8
//		uint16
//		uint32
//		uint	(equivalent to uint32)
//		uint64
//
//	This unsigned integer numeric value is then converted to
//	a type of 'NumberStrKernel' and returned to the calling
//	function.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and reset to new values.
//
//	intNumericValue	interface{}
//
//		This empty interface is assumed to encapsulate an unsigned
//		integer numeric value comprised of one of the following
//		types:
//			uint8
//			uint16
//			uint32
//			uint (equivalent to uint32)
//			uint64
//
//		If the object passed by this empty interface is NOT one of
//		the types listed above, an error will be returned.
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
//		Unsigned integer values are be default converted as
//		positive numeric values. If this parameter is set
//		to NumSignVal.Negative(), the numeric value returned
//		through parameter 'numStrKernel' will be classified
//		as a negative value.
//
//		If 'numberSign' is set to any value other than
//		NumSignVal.Negative(), it will be ignored.
//
//	errPrefDto          *ePref.ErrPrefixDto
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
func (numStrKernelMolecule *numberStrKernelMolecule) convertUnsignedInteger(
	numStrKernel *NumberStrKernel,
	unsignedIntValue interface{},
	numberSign NumericSignValueType,
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
		"numberStrKernelMolecule."+
			"convertUnsignedInteger()",
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

	switch unsignedIntValue.(type) {

	case uint8, uint16, uint, uint32, uint64:

		goto unsignedIntProcessing

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'unsignedIntValue' is an invalid type!\n"+
			"Only unsigned integer types are suppored.\n"+
			"'unsignedIntValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", unsignedIntValue))

		return err

	}

unsignedIntProcessing:

	new(numberStrKernelElectron).empty(
		numStrKernel)

	numberStr := fmt.Sprintf("%v",
		unsignedIntValue)

	var searchResults CharSearchNumStrParseResultsDto
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

	searchResults,
		*numStrKernel,
		err = new(numStrBuilderElectron).extractNumRunes(
		runeArrayDto,
		"unsignedIntegerNumberStr",
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

	if !searchResults.FoundNumericDigits {
		err = fmt.Errorf("%v\n"+
			"Error: No Numeric Digits Found in 'numberStr'!\n",
			ePrefix.String())

	}

	if numStrKernel.numberSign == NumSignVal.Zero() {

		return err
	}

	if numberSign == NumSignVal.Negative() {
		numStrKernel.numberSign = NumSignVal.Negative()
	}

	return err
}

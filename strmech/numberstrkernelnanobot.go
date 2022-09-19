package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"sync"
	"time"
)

// numberStrKernelNanobot - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingNumStrKernel' to input parameter 'targetNumStrKernel'.
// Both instances are of type NumberStrKernel.
//
// # IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'targetNumStrKernel' will be
// overwritten.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	targetNumStrKernel        *NumberStrKernel
//	   - A pointer to a NumberStrKernel instance. All the
//	     member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'incomingNumStrKernel'.
//
//	     'targetNumStrKernel' is the target of this copy
//	     operation.
//
//	incomingNumStrKernel      *NumberStrKernel
//	   - A pointer to another NumberStrKernel instance. All
//	     the member variable data values from this object will
//	     be copied to corresponding member variables in
//	     'targetNumStrKernel'.
//
//	     'incomingNumStrKernel' is the source for this copy
//	     operation.
//
//	     If 'incomingNumStrKernel' is determined to be invalid,
//	     an error will be returned.
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (numStrKernelNanobot *numberStrKernelNanobot) copyIn(
	targetNumStrKernel *NumberStrKernel,
	incomingNumStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetNumStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingNumStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	var err2 error

	numStrKernelAtom := numberStrKernelAtom{}

	_,
		err2 =
		numStrKernelAtom.testValidityOfNumStrKernel(
			incomingNumStrKernel,
			nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'incomingNumStrKernel' failed!\n"+
			"This instance of NumberStrKernel is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	new(numberStrKernelElectron).empty(
		targetNumStrKernel)

	err = targetNumStrKernel.integerDigits.CopyIn(
		&incomingNumStrKernel.integerDigits,
		ePrefix.XCpy(
			"targetNumStrKernel.integerDigits"+
				"<-incomingNumStrKernel.integerDigits"))

	if err != nil {
		return err
	}

	err = targetNumStrKernel.fractionalDigits.CopyIn(
		&incomingNumStrKernel.fractionalDigits,
		ePrefix.XCpy(
			"targetNumStrKernel.fractionalDigits"+
				"<-incomingNumStrKernel.fractionalDigits"))

	if err != nil {
		return err
	}

	targetNumStrKernel.numericValueType =
		incomingNumStrKernel.numericValueType

	targetNumStrKernel.numberSign =
		incomingNumStrKernel.numberSign

	targetNumStrKernel.isNonZeroValue =
		incomingNumStrKernel.isNonZeroValue

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'numStrKernel'. a pointer to an instance of
// NumberStrKernel.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// The input parameter 'numStrKernel' is determined to be invalid,
// this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	numStrKernel           *NumberStrKernel
//	   - A pointer to an instance of NumberStrKernel. A
//	     deep copy of the internal member variables will be created
//	     and returned in a new instance of NumberStrKernel.
//
//	     If the member variable data values encapsulated by
//	     'numStrKernel' are found to be invalid, this method will
//	     return an error
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	deepCopyNumStrKernel       NumberStrKernel
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'numStrKernel' will be created and
//	     returned in a new instance of NumberStrKernel.
//
//
//	err                        error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (numStrKernelNanobot *numberStrKernelNanobot) copyOut(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyNumStrKernel NumberStrKernel,
	err error) {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	new(numberStrKernelElectron).
		empty(&deepCopyNumStrKernel)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyNumStrKernel, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyNumStrKernel, err
	}

	var err2 error

	_,
		err2 = numberStrKernelAtom{}.ptr().
		testValidityOfNumStrKernel(
			numStrKernel,
			ePrefix.XCpy(
				"numStrKernel"))

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'numStrKernel' failed!\n"+
			"This instance of NumberStrKernel is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return deepCopyNumStrKernel, err
	}

	err = deepCopyNumStrKernel.integerDigits.CopyIn(
		&numStrKernel.integerDigits,
		ePrefix.XCpy(
			"deepCopyNumStrKernel.integerDigits"+
				"<-numStrKernel.integerDigits"))

	if err != nil {
		return deepCopyNumStrKernel, err
	}

	err = deepCopyNumStrKernel.fractionalDigits.CopyIn(
		&numStrKernel.fractionalDigits,
		ePrefix.XCpy(
			"deepCopyNumStrKernel.fractionalDigits"+
				"<-numStrKernel.fractionalDigits"))

	if err != nil {
		return deepCopyNumStrKernel, err
	}

	deepCopyNumStrKernel.numericValueType =
		numStrKernel.numericValueType

	deepCopyNumStrKernel.numberSign =
		numStrKernel.numberSign

	deepCopyNumStrKernel.isNonZeroValue =
		numStrKernel.isNonZeroValue

	return deepCopyNumStrKernel, err
}

// getParameterTextListing - Returns formatted text output
// listing the member variable names and corresponding values
// contained in the 'numStrKernel' instance of NumberStrKernel.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 strBuilder                 *strings.Builder
//	    - A pointer to an instance of *strings.Builder. The
//	      formatted text characters produced by this method will be
//	      written to this instance of strings.Builder.
//
//
//		displayFunctionChain       bool
//		   - Set 'displayFunctionChain' to 'true' and a list of the
//		     functions which led to this result will be included in
//		     the formatted text output.
//
//
//	 numStrKernel               *NumberStrKernel
//	    - A pointer to an instance of NumberStrKernel instance.
//	      Formatted text output will be generated listing the member
//	      variable names and their corresponding values. The
//	      formatted text can then be used for screen displays, file
//	      output or printing.
//
//	      No data validation is performed on this instance of
//	      CharSearchDecimalSeparatorResultsDto.
//
//
//	 errPrefDto                 *ePref.ErrPrefixDto
//	    - This object encapsulates an error prefix string which is
//	      included in all returned error messages. Usually, it
//	      contains the name of the calling method or methods listed
//	      as a function chain.
//
//	      If no error prefix information is needed, set this
//	      parameter to 'nil'.
//
//	      Type ErrPrefixDto is included in the 'errpref' software
//	      package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (numStrKernelNanobot *numberStrKernelNanobot) getParameterTextListing(
	strBuilder *strings.Builder,
	displayFunctionChain bool,
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) error {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	requiredCapacity :=
		256 - netCapacityStrBuilder

	if requiredCapacity > 0 {

		strBuilder.Grow(requiredCapacity + 16)
	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	// Total available Length of Output Line
	const maxLineLen = 79

	// Max Label Field Length = 38
	const maxLabelFieldLen = 38

	txtFormatCol := TextFormatterCollection{}

	solidLineDto := TextLineSolidDto{
		FormatType:                 TxtFieldType.SolidLine(),
		LeftMarginStr:              " ",
		SolidLineChars:             "=",
		SolidLineCharRepeatCount:   maxLineLen - 2,
		RightMarginStr:             " ",
		TurnLineTerminationOff:     false,
		LineTerminator:             "",
		MaxLineLength:              -1,
		TurnAutoLineLengthBreaksOn: false,
		lock:                       nil,
	}

	err = txtFormatCol.SetStdFormatParamsLine1Col(
		"",
		maxLineLen,
		TxtJustify.Center(),
		"",
		false,
		"",
		-1,
		false,
		ePrefix.XCpy(
			"1-Column Setup"))

	if err != nil {
		return err
	}

	// Leading Title Marquee

	// Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Title Line 1
	err = txtFormatCol.AddLine1Col(
		"NumberStrKernel",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return err
	}

	// Title Line 2
	err = txtFormatCol.AddLine1Col(
		"Parameter Listing",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return err
	}

	// Title Line  3 Date/Time

	err = txtFormatCol.AddLine1Col(
		TextInputParamFieldDateTimeDto{
			FieldDateTime:       time.Now(),
			FieldDateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		},
		ePrefix.XCpy(
			"Top-Title Line 2"))

	if err != nil {
		return err
	}

	// Filler Line '========='
	// Marquee Bottom
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Trailing Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// End Of Marquee

	// Begin Label Parameter Pairs

	colonSpace := ": "

	var txtStrParam, txtStrLabel string

	// Build NumStrKernelFunctionChain

	if displayFunctionChain == true {

		txtStrLabel := "Number String Kernel Function Chain"

		txtStrParam = ePrefix.String()

		if len(txtStrParam) == 0 {

			txtStrParam = "Number String Kernel Function Chain is EMPTY!"

			txtFormatCol.AddFieldLabel(
				" ",
				txtStrLabel,
				maxLabelFieldLen,
				TxtJustify.Right(),
				colonSpace,
				"",
				-1,
				false)

			txtFormatCol.AddFieldLabel(
				"",
				txtStrParam,
				-1,
				TxtJustify.Left(),
				"",
				"\n",
				-1,
				false)

		} else {

			txtFormatCol.AddFieldLabel(
				" ",
				txtStrLabel,
				maxLabelFieldLen,
				TxtJustify.Right(),
				colonSpace,
				"\n",
				-1,
				false)

			spacer := strings.Repeat(" ", 16)

			txtStrParam = strings.Replace(
				txtStrParam,
				"\n",
				"\n"+spacer,
				-1)

			txtStrParam = "\n" + spacer + txtStrParam

			txtFormatCol.AddFieldLabel(
				"",
				txtStrParam,
				-1,
				TxtJustify.Left(),
				"",
				"\n",
				-1,
				false)

		}

		txtFormatCol.AddLineBlank(1, "")
	}

	// Set up 2-Column Parameters

	// Build Integer Digits

	txtStrLabel = "Integer Digits"

	txtStrParam =
		numStrKernel.integerDigits.GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "Integer Digits is EMPTY!"
	}

	err = txtFormatCol.CfgLine2Col(
		" ",
		txtStrLabel,
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		txtStrParam,
		-1,
		TxtJustify.Left(),
		"",
		false,
		"",
		maxLineLen,
		true,
		true,
		ePrefix.XCpy(
			"Set 2-Column Params Integer Digits"))

	if err != nil {
		return err
	}

	// Build Fractional Digits

	txtStrLabel = "Fractional Digits"

	txtStrParam =
		numStrKernel.fractionalDigits.GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "Fractional Digits is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"Fractional Digits"))

	if err != nil {
		return err
	}

	// Build NumericValueType

	txtStrLabel = "Numeric Value Type"

	if !numStrKernel.numericValueType.XIsValid() {
		numStrKernel.numericValueType = NumValType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		numStrKernel.numericValueType,
		ePrefix.XCpy(
			"Numeric Value Type"))

	if err != nil {
		return err
	}

	// Build NumberSign

	txtStrLabel = "Number Sign"

	if !numStrKernel.numberSign.XIsValid() {
		numStrKernel.numberSign = NumSignVal.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		numStrKernel.numberSign,
		ePrefix.XCpy(
			"Number Sign"))

	if err != nil {
		return err
	}

	// Build Is Non Zero Value

	txtStrLabel = "Is Non Zero Value"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		numStrKernel.isNonZeroValue,
		ePrefix.XCpy(
			"Is Non Zero Value"))

	if err != nil {
		return err
	}

	// Trailing Title Marquee
	// Top Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Title # 1
	err = txtFormatCol.AddLine1Col(
		"NumberStrKernel",
		ePrefix.XCpy(
			"Bottom-Title Line 1"))

	if err != nil {
		return err
	}

	// Title # 2
	err = txtFormatCol.AddLine1Col(
		"End of Parameter Listing",
		ePrefix.XCpy(
			"Top-Title Line 3"))

	if err != nil {
		return err
	}

	// Filler =======
	// Marquee Bottom
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Blank Line
	txtFormatCol.AddLineBlank(
		2,
		"")

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Text Output"))

	return err
}

// setWithNumber
//
// Sets the internal member variable values for an instance
// of NumberStrKernel passed as an input parameter.
//
// This method assumes that the numeric value passed as
// an empty interface object is one of the following types:
//
//	int8
//	int16
//	int32
//	int	(equivalent to int32)
//	int64
//	uint8
//	uint16
//	uint32
//	uint (equivalent to uint32)
//	uint64
//	*big.Int
//	float32
//	float64
//	*big.Float
//
// If the empty interface object is not one of the types
// listed above, an error will be returned.
//
// # IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'numStrKernel' will be
// deleted and set to new values.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	numStrKernel				*NumberStrKernel
//
//		A pointer to an instance of NumberStrKernel. The data
//		values for all internal member variables contained in
//		this instance will be deleted and set to new values.
//
//	numericValue				interface{}
//
//		A numeric value passed by means of an empty interface.
//		This numeric value will be used to populate the instance
//		of NumberStrKernel passed by parameter, 'numStrKernel'.
//
//		The object passed by this empty interface must be one of
//		the following types:
//
//			int8
//			int16
//			int32
//			int	(equivalent to int32)
//			int64
//			uint8
//			uint16
//			uint32
//			uint (equivalent to uint32)
//			uint64
//			*big.Int
//			float32
//			float64
//			*big.Float
//
//	If the object passed by this parameter is not one of the types
//	listed above, an error will be returned.
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
//		number sign will be inferred from the numeric value
//		and generated as a default. Unsigned integer are
//		assigned a positive number sign, signed integers and
//		floating point values are assigned according to their
//		inherent numbers signs.
//
//		A valid 'numberSign' setting will override the default
//		number sign associated with 'numericValue'.
//
//		If 'numberSign' is set to an invalid value or if
//		'numberSign' is set to NumSignVal.Zero() for a
//		non-zero numeric value, an error will be returned.
//
//		If parameter 'numericValue' includes a signed numeric
//		value, parameter 'numberSign' will override that
//		designation.
func (numStrKernelNanobot *numberStrKernelNanobot) setWithNumber(
	numStrKernel *NumberStrKernel,
	numericValue interface{},
	numberSign NumericSignValueType,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelNanobot."+
			"setWithNumber()",
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

	var nStrKernelMolecule numberStrKernelMolecule

	var ok bool

	switch numericValue.(type) {

	case int8, int16, int32, int, int64:

		return nStrKernelMolecule.convertSignedIntToKernel(
			numStrKernel,
			numericValue,
			ePrefix)

	case uint8, uint16, uint, uint32, uint64:

		return nStrKernelMolecule.convertUnsignedInteger(
			numStrKernel,
			numericValue,
			numberSign,
			ePrefix)

	case *big.Int:

		var bigIntNum *big.Int

		bigIntNum, ok = numericValue.(*big.Int)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Int cast to 'bigIntNum' failed!\n",
				ePrefix.String())

			return err

		}

		return nStrKernelMolecule.convertBigIntToKernel(
			numStrKernel,
			bigIntNum,
			ePrefix)

	case float32, float64:

		return nStrKernelMolecule.convertFloatToKernel(
			numStrKernel,
			numericValue,
			ePrefix)

	case *big.Float:

		var bigFloatNum *big.Float

		bigFloatNum, ok = numericValue.(*big.Float)

		if !ok {

			err = fmt.Errorf("%v\n"+
				"ERROR: *big.Float cast to 'bigFloatNum' failed!\n",
				ePrefix.String())

			return err

		}

		return nStrKernelMolecule.convertBigFloatToKernel(
			numStrKernel,
			bigFloatNum,
			ePrefix)

	default:

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numericValue' is an invalid type!\n"+
			"'numericValue' is unsupported type '%v'\n",
			ePrefix.String(),
			fmt.Sprintf("%T", numericValue))

	}

	return err
}

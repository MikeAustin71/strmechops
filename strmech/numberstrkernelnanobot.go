package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// numberStrKernelNanobot - Provides helper methods for type
// NumberStrKernel.
type numberStrKernelNanobot struct {
	lock *sync.Mutex
}

// copy
//
// Copies all data from input parameter
// 'sourceNumStrKernel' to input parameter
// 'destinationNumStrKernel'. Both instances are of type
// NumberStrKernel.
//
// # IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in
// 'destinationNumStrKernel' will be deleted and
// overwritten.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	destinationNumStrKernel        *NumberStrKernel
//	   - A pointer to a NumberStrKernel instance. All the
//	     member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceNumStrKernel'.
//
//	     'destinationNumStrKernel' is the target of this copy
//	     operation.
//
//	sourceNumStrKernel      *NumberStrKernel
//	   - A pointer to another NumberStrKernel instance. All
//	     the member variable data values from this object will
//	     be copied to corresponding member variables in
//	     'destinationNumStrKernel'.
//
//	     'sourceNumStrKernel' is the source for this copy
//	     operation.
//
//	     If 'sourceNumStrKernel' is determined to be invalid,
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
func (numStrKernelNanobot *numberStrKernelNanobot) copy(
	destinationNumStrKernel *NumberStrKernel,
	sourceNumStrKernel *NumberStrKernel,
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
			"copy()",
		"")

	if err != nil {

		return err

	}

	if destinationNumStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetNegNumSearchSpec' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceNumStrKernel == nil {

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
			sourceNumStrKernel,
			nil)

	if err2 != nil {

		err = fmt.Errorf("%v\n"+
			"Error: Validation of input parameter 'sourceNumStrKernel' failed!\n"+
			"This instance of NumberStrKernel is invalid.\n"+
			"Validation error message reads as follows:\n"+
			"%v\n",
			ePrefix.String(),
			err2.Error())

		return err
	}

	new(numberStrKernelElectron).empty(
		destinationNumStrKernel)

	err = destinationNumStrKernel.integerDigits.CopyIn(
		&sourceNumStrKernel.integerDigits,
		ePrefix.XCpy(
			"destinationNumStrKernel.integerDigits"+
				"<-sourceNumStrKernel.integerDigits"))

	if err != nil {
		return err
	}

	err = destinationNumStrKernel.fractionalDigits.CopyIn(
		&sourceNumStrKernel.fractionalDigits,
		ePrefix.XCpy(
			"destinationNumStrKernel.fractionalDigits"+
				"<-sourceNumStrKernel.fractionalDigits"))

	if err != nil {
		return err
	}

	destinationNumStrKernel.numberValueType =
		sourceNumStrKernel.numberValueType

	destinationNumStrKernel.numberSign =
		sourceNumStrKernel.numberSign

	destinationNumStrKernel.isNonZeroValue =
		sourceNumStrKernel.isNonZeroValue

	if sourceNumStrKernel.numStrFormatSpec.IsValidInstance() {

		err = destinationNumStrKernel.numStrFormatSpec.CopyIn(
			&sourceNumStrKernel.numStrFormatSpec,
			ePrefix.XCpy(
				"destinationNumStrKernel<-"+
					"sourceNumStrKernel.numStrFormatSpec"))
	} else {

		destinationNumStrKernel.numStrFormatSpec,
			err = new(NumStrFormatSpec).NewSignedNumFmtUS(
			NumStrNumberFieldSpec{
				fieldLength:        -1,
				fieldJustification: TxtJustify.Right(),
			},
			ePrefix.XCpy("destinationNumStrKernel "+
				"Default NumStrNumberFieldSpec"))

		if err != nil {
			return err
		}

		sourceNumStrKernel.numStrFormatSpec,
			err = new(NumStrFormatSpec).NewSignedNumFmtUS(
			NumStrNumberFieldSpec{
				fieldLength:        -1,
				fieldJustification: TxtJustify.Right(),
			},
			ePrefix.XCpy("sourceNumStrKernel "+
				"Default NumStrNumberFieldSpec"))

	}

	return err
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

	if !numStrKernel.numberValueType.XIsValid() {
		numStrKernel.numberValueType = NumValType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		numStrKernel.numberValueType,
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

//	setWithRunes
//
//	Deletes and resets all the internal member variable
//	values for an instance of NumberStrKernel passed as
//	an input parameter. Integer and fractional digits
//	are configured from rune array parameters passed by
//	the calling function.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
//	Be advised that all the data fields within input
//	parameter 'numStrKernel' will be deleted and reset
//	to new values.
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
//	integerDigitRunes			[]rune
//
//		A rune array used to configure the integer digits
//		array contained within the NumberStrKernel instance,
//		'numStrKernel'.
//
//	fractionalDigitRunes		[]rune
//
//		A rune array used to configure the fractional
//		digits array contained within the NumberStrKernel
//		instance, 'numStrKernel'.
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
func (numStrKernelNanobot *numberStrKernelNanobot) setWithRunes(
	numStrKernel *NumberStrKernel,
	integerDigitRunes []rune,
	fractionalDigitRunes []rune,
	numberSign NumericSignValueType,
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
			"setWithRunes()",
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

	nStrKernelElectron := numberStrKernelElectron{}

	nStrKernelElectron.empty(numStrKernel)

	lenIntDigits := len(integerDigitRunes)

	lenFracDigits := len(fractionalDigitRunes)

	if lenIntDigits == 0 &&
		lenFracDigits == 0 {

		numStrKernel.integerDigits.CharsArray =
			make([]rune, 1)

		numStrKernel.integerDigits.CharsArray[0] =
			'0'

		numStrKernel.numberValueType = NumValType.Integer()

		numStrKernel.numberSign = NumSignVal.Zero()

		numStrKernel.isNonZeroValue = false

		return err
	}

	nStrKernelAtom := numberStrKernelAtom{}

	for i := 0; i < lenIntDigits; i++ {

		err = nStrKernelAtom.addIntegerDigit(
			numStrKernel,
			integerDigitRunes[i],
			ePrefix.XCpy(
				fmt.Sprintf(
					"integerDigitRunes[%v]=%v",
					i,
					integerDigitRunes[i])))

		if err != nil {
			return err
		}
	}

	for j := 0; j < lenFracDigits; j++ {

		err = nStrKernelAtom.addFractionalDigit(
			numStrKernel,
			fractionalDigitRunes[j],
			ePrefix.XCpy(
				fmt.Sprintf(
					"fractionalDigitRunes[%v]=%v",
					j,
					fractionalDigitRunes[j])))

		if err != nil {
			return err
		}
	}

	if numStrKernel.isNonZeroValue == false {

		numStrKernel.numberSign = NumSignVal.Zero()

	} else {

		numStrKernel.numberSign = NumSignVal.Positive()

		if numberSign == NumSignVal.Negative() {

			numStrKernel.numberSign = NumSignVal.Negative()

		}
	}

	err = nStrKernelElectron.rationalizeFractionalIntegerDigits(
		numStrKernel,
		ePrefix)

	return err
}

//	setWithRuneArrayDto
//
//	Deletes and resets all the internal member variable
//	values for an instance of NumberStrKernel passed as
//	an input parameter. Integer and fractional digits
//	are configured from rune array data transfer object
//	(Dto) parameters passed by the calling function.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
//	Be advised that all the data fields within input
//	parameter 'numStrKernel' will be deleted and reset
//	to new values.
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
//	integerDigits				*RuneArrayDto
//
//		A pointer to a rune array data transfer object used to
//		configure the integer digits array contained within
//		the NumberStrKernel instance, 'numStrKernel'.
//
//	fractionalDigits			*RuneArrayDto
//
//		A pointer to a rune array data transfer object used to
//		configure the fractional digits array contained within
//		the NumberStrKernel instance, 'numStrKernel'.
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
func (numStrKernelNanobot *numberStrKernelNanobot) setWithRuneArrayDto(
	numStrKernel *NumberStrKernel,
	integerDigits *RuneArrayDto,
	fractionalDigits *RuneArrayDto,
	numberSign NumericSignValueType,
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
			"setWithRuneArrayDto()",
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

	if integerDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'integerDigits' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if fractionalDigits == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'fractionalDigits' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	nStrKernelElectron := numberStrKernelElectron{}

	nStrKernelElectron.empty(numStrKernel)

	lenIntDigits := len(integerDigits.CharsArray)

	lenFracDigits := len(fractionalDigits.CharsArray)

	if lenIntDigits == 0 &&
		lenFracDigits == 0 {

		numStrKernel.integerDigits.CharsArray =
			make([]rune, 1)

		numStrKernel.integerDigits.CharsArray[0] =
			'0'

		numStrKernel.numberValueType = NumValType.Integer()

		numStrKernel.numberSign = NumSignVal.Zero()

		numStrKernel.isNonZeroValue = false

		return err
	}

	nStrKernelAtom := numberStrKernelAtom{}

	for i := 0; i < lenIntDigits; i++ {

		err = nStrKernelAtom.addIntegerDigit(
			numStrKernel,
			integerDigits.CharsArray[i],
			ePrefix.XCpy(
				fmt.Sprintf(
					"integerDigits.CharsArray[%v]=%v",
					i,
					integerDigits.CharsArray[i])))

		if err != nil {
			return err
		}
	}

	for j := 0; j < lenFracDigits; j++ {

		err = nStrKernelAtom.addFractionalDigit(
			numStrKernel,
			fractionalDigits.CharsArray[j],
			ePrefix.XCpy(
				fmt.Sprintf(
					"fractionalDigits.CharsArray[%v]=%v",
					j,
					fractionalDigits.CharsArray[j])))

		if err != nil {
			return err
		}
	}

	if numStrKernel.isNonZeroValue == false {

		numStrKernel.numberSign = NumSignVal.Zero()

	} else {

		numStrKernel.numberSign = NumSignVal.Positive()

		if numberSign == NumSignVal.Negative() {

			numStrKernel.numberSign = NumSignVal.Negative()

		}
	}

	err = nStrKernelElectron.rationalizeFractionalIntegerDigits(
		numStrKernel,
		ePrefix)

	return err
}

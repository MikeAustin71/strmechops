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
//
type numberStrKernelNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingNumStrKernel' to input parameter 'targetNumStrKernel'.
// Both instances are of type NumberStrKernel.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'targetNumStrKernel' will be
// overwritten.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  targetNumStrKernel        *NumberStrKernel
//     - A pointer to a NumberStrKernel instance. All the
//       member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'incomingNumStrKernel'.
//
//       'targetNumStrKernel' is the target of this copy
//       operation.
//
//
//  incomingNumStrKernel      *NumberStrKernel
//     - A pointer to another NumberStrKernel instance. All
//       the member variable data values from this object will
//       be copied to corresponding member variables in
//       'targetNumStrKernel'.
//
//       'incomingNumStrKernel' is the source for this copy
//       operation.
//
//       If 'incomingNumStrKernel' is determined to be invalid,
//       an error will be returned.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
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

	numStrKernelElectron := numberStrKernelElectron{}

	numStrKernelElectron.empty(
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
// IMPORTANT
//
// The input parameter 'numStrKernel' is determined to be invalid,
// this method will return an error.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  numStrKernel           *NumberStrKernel
//     - A pointer to an instance of NumberStrKernel. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of NumberStrKernel.
//
//       If the member variable data values encapsulated by
//       'numStrKernel' are found to be invalid, this method will
//       return an error
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  deepCopyNumStrKernel       NumberStrKernel
//     - If this method completes successfully, a deep copy of
//       input parameter 'numStrKernel' will be created and
//       returned in a new instance of NumberStrKernel.
//
//
//  err                        error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
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

	numberStrKernelElectron{}.ptr().
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
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrKernel               *NumberStrKernel
//     - A pointer to an instance of NumberStrKernel instance.
//       Formatted text output will be generated listing the member
//       variable names and their corresponding values. The
//       formatted text can then be used for screen displays, file
//       output or printing.
//
//       No data validation is performed on this instance of
//       CharSearchDecimalSeparatorResultsDto.
//
//
//  errPrefDto                 *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this
//       parameter to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//  strings.Builder
//     - If this method completes successfully, an instance of
//       strings.Builder will be returned. This instance contains
//       the formatted text output listing the member variable
//       names and their corresponding values for input parameter
//       'numStrKernel' . This formatted text can them be used for
//       screen displays, file output or printing.
//
//
//  error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (numStrKernelNanobot *numberStrKernelNanobot) getParameterTextListing(
	numStrKernel *NumberStrKernel,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"numberStrKernelNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {

		return strBuilder, err

	}

	if numStrKernel == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrKernel' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	// Total available Length of Output Line
	const maxFieldLen = 70

	// Max Label Field Length = 38
	const maxLabelFieldLen = 38

	txtBuilder := TextStrBuilder{}

	// Leading Title Marquee
	var fmtrs []TextFormatterDto

	// Blank Line
	txtFmt := TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.BlankLine()
	txtFmt.BlankLine.NumOfBlankLines = 1
	fmtrs = append(fmtrs, txtFmt)

	// Filler =======
	// Marquee Top
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Filler()
	txtFmt.Filler.LeftMarginStr = " "
	txtFmt.Filler.FillerCharacters = "="
	txtFmt.Filler.FillerCharsRepeatCount = maxFieldLen - 2
	txtFmt.Filler.RightMarginStr = " "
	txtFmt.Filler.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title Line 1
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = "NumberStrKernel"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title Line 2
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = "Parameter Listing"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title Line  3 Date/Time
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.DateTime()
	txtFmt.DateTime.LeftMarginStr = ""
	txtFmt.DateTime.FieldDateTime = time.Now()
	txtFmt.DateTime.FieldLength = maxFieldLen
	txtFmt.DateTime.FieldJustify = TxtJustify.Center()
	txtFmt.DateTime.FieldDateTimeFormat =
		"Monday 2006-01-02 15:04:05.000000000 -0700 MST"
	txtFmt.DateTime.RightMarginStr = ""
	txtFmt.DateTime.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Filler Line '========='
	// Marquee Bottom
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Filler()
	txtFmt.Filler.LeftMarginStr = " "
	txtFmt.Filler.FillerCharacters = "="
	txtFmt.Filler.FillerCharsRepeatCount = maxFieldLen - 2
	txtFmt.Filler.RightMarginStr = " "
	txtFmt.Filler.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Trailing Blank Line
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.BlankLine()
	txtFmt.BlankLine.NumOfBlankLines = 1
	fmtrs = append(fmtrs, txtFmt)

	var strBuilder2 strings.Builder

	strBuilder2,
		err = txtBuilder.BuildTextFormatters(
		fmtrs,
		ePrefix.XCpy(
			"strBuilder<-Marquee Top"))

	if err != nil {

		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	fmtrs = nil

	// End Of Marquee

	// Begin Label Parameter Pairs

	colonSpace := ": "

	var labelParams []TextLabelValueStrings

	// Build Integer Digits
	labelParam := TextLabelValueStrings{}

	labelParam.ParamLabel = "Integer Digits"

	labelParam.ParamValue =
		numStrKernel.integerDigits.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "Integer Digits is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build Fractional Digits
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "Fractional Digits"

	labelParam.ParamValue =
		numStrKernel.fractionalDigits.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "Fractional Digits is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build NumericValueType
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "Numeric Value Type"

	if !numStrKernel.numericValueType.XIsValid() {
		numStrKernel.numericValueType = NumValType.None()
	}

	labelParam.ParamValue =
		numStrKernel.numericValueType.String()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "Numeric Value Type is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build NumberSign
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "Number Sign"

	if !numStrKernel.numberSign.XIsValid() {
		numStrKernel.numberSign = NumSignVal.None()
	}

	labelParam.ParamValue =
		numStrKernel.numericValueType.String()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "Number Sign Type is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build Is Non Zero Value
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "Is Non Zero Value"

	labelParam.ParamValue =
		fmt.Sprintf("%v",
			numStrKernel.isNonZeroValue)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "Is Non Zero Value is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Write Label/Parameter Values to String Builder
	strBuilder2,
		err = txtBuilder.BuildLabelsValues(
		labelParams,
		" ",
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		-1,
		TxtJustify.Left(),
		" ",
		"\n",
		ePrefix.XCpy(
			"labelParams #2"))

	labelParams = nil

	if err != nil {

		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	// Trailing Title Marquee
	// Top Blank Line
	fmtrs = nil
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.BlankLine()
	txtFmt.BlankLine.NumOfBlankLines = 1
	fmtrs = append(fmtrs, txtFmt)

	// Filler =======
	// Marquee Top
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Filler()
	txtFmt.Filler.LeftMarginStr = " "
	txtFmt.Filler.FillerCharacters = "="
	txtFmt.Filler.FillerCharsRepeatCount = maxFieldLen - 2
	txtFmt.Filler.RightMarginStr = " "
	txtFmt.Filler.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title # 1
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = "NumberStrKernel"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title # 2
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = "End of Parameter Listing"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Filler =======
	// Marquee Bottom
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Filler()
	txtFmt.Filler.LeftMarginStr = " "
	txtFmt.Filler.FillerCharacters = "="
	txtFmt.Filler.FillerCharsRepeatCount = maxFieldLen - 2
	txtFmt.Filler.RightMarginStr = " "
	txtFmt.Filler.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Blank Line
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.BlankLine()
	txtFmt.BlankLine.NumOfBlankLines = 2
	fmtrs = append(fmtrs, txtFmt)

	strBuilder2,
		err = txtBuilder.BuildTextFormatters(
		fmtrs,
		ePrefix.XCpy(
			"Marquee-Bottom"))
	fmtrs = nil

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	return strBuilder, err
}

// ptr - Returns a pointer to a new instance of
// numberStrKernelNanobot.
//
func (numStrKernelNanobot numberStrKernelNanobot) ptr() *numberStrKernelNanobot {

	if numStrKernelNanobot.lock == nil {
		numStrKernelNanobot.lock = new(sync.Mutex)
	}

	numStrKernelNanobot.lock.Lock()

	defer numStrKernelNanobot.lock.Unlock()

	return &numberStrKernelNanobot{
		lock: new(sync.Mutex),
	}
}

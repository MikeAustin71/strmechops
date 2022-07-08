package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchTestConfigDtoNanobot - Provides helper methods for type
// CharSearchTestConfigDto.
//
type charSearchTestConfigDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceTestCfgDto' to input parameter
// 'destinationTestCfgDto'. Both instances are of type
// CharSearchTestConfigDto.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationTestCfgDto'
// will be deleted and overwritten.
//
// Also, NO validation is performed on 'sourceTestCfgDto'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  destinationTestCfgDto     *CharSearchTestConfigDto
//     - A pointer to a CharSearchTestConfigDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceTestCfgDto'.
//
//       'destinationTestCfgDto' is the destination for this
//       copy operation.
//
//
//  sourceTestCfgDto          *CharSearchTestConfigDto
//     - A pointer to another CharSearchTestConfigDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationTestCfgDto'.
//
//       'sourceTestCfgDto' is the source for this copy
//       operation.
//
//       No data validation is performed on
//       'sourceTestCfgDto'.
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
func (searchTestConfigNanobot *charSearchTestConfigDtoNanobot) copyIn(
	destinationTestCfgDto *CharSearchTestConfigDto,
	sourceTestCfgDto *CharSearchTestConfigDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestConfigDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchTestConfigDtoAtom{}.ptr().empty(
		destinationTestCfgDto)

	destinationTestCfgDto.TestInputParametersName =
		sourceTestCfgDto.TestInputParametersName

	destinationTestCfgDto.TestStringName =
		sourceTestCfgDto.TestStringName

	destinationTestCfgDto.TestStringLengthName =
		sourceTestCfgDto.TestStringLengthName

	destinationTestCfgDto.TestStringStartingIndex =
		sourceTestCfgDto.TestStringStartingIndex

	destinationTestCfgDto.TestStringStartingIndexName =
		sourceTestCfgDto.TestStringStartingIndexName

	destinationTestCfgDto.TestStringDescription1 =
		sourceTestCfgDto.TestStringDescription1

	destinationTestCfgDto.TestStringDescription2 =
		sourceTestCfgDto.TestStringDescription2

	destinationTestCfgDto.CollectionTestObjIndex =
		sourceTestCfgDto.CollectionTestObjIndex

	destinationTestCfgDto.NumValueType =
		sourceTestCfgDto.NumValueType

	destinationTestCfgDto.NumStrFormatType =
		sourceTestCfgDto.NumStrFormatType

	destinationTestCfgDto.NumSymbolLocation =
		sourceTestCfgDto.NumSymbolLocation

	destinationTestCfgDto.NumSymbolClass =
		sourceTestCfgDto.NumSymbolClass

	destinationTestCfgDto.NumSignValue =
		sourceTestCfgDto.NumSignValue

	destinationTestCfgDto.PrimaryNumSignPosition =
		sourceTestCfgDto.PrimaryNumSignPosition

	destinationTestCfgDto.SecondaryNumSignPosition =
		sourceTestCfgDto.SecondaryNumSignPosition

	destinationTestCfgDto.TextCharSearchType =
		sourceTestCfgDto.TextCharSearchType

	destinationTestCfgDto.RequestFoundTestCharacters =
		sourceTestCfgDto.RequestFoundTestCharacters

	destinationTestCfgDto.RequestRemainderString =
		sourceTestCfgDto.RequestRemainderString

	destinationTestCfgDto.RequestReplacementString =
		sourceTestCfgDto.RequestReplacementString

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'searchTestCfgDto', a pointer to an instance of
// CharSearchTestConfigDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
//
// NO validation is performed on 'searchTestCfgDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  searchTestCfgDto                *CharSearchTestConfigDto
//     - A pointer to an instance of CharSearchTestConfigDto. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of CharSearchTestConfigDto.
//
//       No data validation is performed on 'searchTestCfgDto'.
//
//
//  errPrefDto                      *ePref.ErrPrefixDto
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
//  deepCopySearchTestCfgDto        CharSearchTestConfigDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'searchTestCfgDto' will be created and
//       returned in a new instance of CharSearchTestConfigDto.
//
//
//  err                             error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchTestConfigNanobot *charSearchTestConfigDtoNanobot) copyOut(
	searchTestCfgDto *CharSearchTestConfigDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopySearchTestCfgDto CharSearchTestConfigDto,
	err error) {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestConfigDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopySearchTestCfgDto, err

	}

	if searchTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return deepCopySearchTestCfgDto, err
	}

	charSearchTestConfigDtoAtom{}.ptr().empty(
		&deepCopySearchTestCfgDto)

	deepCopySearchTestCfgDto.TestInputParametersName =
		searchTestCfgDto.TestInputParametersName

	deepCopySearchTestCfgDto.TestStringName =
		searchTestCfgDto.TestStringName

	deepCopySearchTestCfgDto.TestStringLengthName =
		searchTestCfgDto.TestStringLengthName

	deepCopySearchTestCfgDto.TestStringStartingIndex =
		searchTestCfgDto.TestStringStartingIndex

	deepCopySearchTestCfgDto.TestStringStartingIndexName =
		searchTestCfgDto.TestStringStartingIndexName

	deepCopySearchTestCfgDto.TestStringDescription1 =
		searchTestCfgDto.TestStringDescription1

	deepCopySearchTestCfgDto.TestStringDescription2 =
		searchTestCfgDto.TestStringDescription2

	deepCopySearchTestCfgDto.CollectionTestObjIndex =
		searchTestCfgDto.CollectionTestObjIndex

	deepCopySearchTestCfgDto.NumValueType =
		searchTestCfgDto.NumValueType

	deepCopySearchTestCfgDto.NumStrFormatType =
		searchTestCfgDto.NumStrFormatType

	deepCopySearchTestCfgDto.NumSymbolLocation =
		searchTestCfgDto.NumSymbolLocation

	deepCopySearchTestCfgDto.NumSymbolClass =
		searchTestCfgDto.NumSymbolClass

	deepCopySearchTestCfgDto.NumSignValue =
		searchTestCfgDto.NumSignValue

	deepCopySearchTestCfgDto.PrimaryNumSignPosition =
		searchTestCfgDto.PrimaryNumSignPosition

	deepCopySearchTestCfgDto.SecondaryNumSignPosition =
		searchTestCfgDto.SecondaryNumSignPosition

	deepCopySearchTestCfgDto.TextCharSearchType =
		searchTestCfgDto.TextCharSearchType

	deepCopySearchTestCfgDto.RequestFoundTestCharacters =
		searchTestCfgDto.RequestFoundTestCharacters

	deepCopySearchTestCfgDto.RequestRemainderString =
		searchTestCfgDto.RequestRemainderString

	deepCopySearchTestCfgDto.RequestReplacementString =
		searchTestCfgDto.RequestReplacementString

	return deepCopySearchTestCfgDto, err
}

// getParameterTextListing - Returns formatted text output
// detailing the member variable values contained in the
// 'searchTestCfgDto' instance of CharSearchTestConfigDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchTestCfgDto           *CharSearchTestInputParametersDto
//     - A pointer to an instance of
//       CharSearchTestInputParametersDto instance. Formatted
//       text output will be generated listing the member variable
//       names and their corresponding values. The formatted text
//       can then be used for text displays, file output or
//       printing.
//
//       No data validation is performed on this instance of
//       CharSearchTestInputParametersDto.
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
//       'searchTestCfgDto' . This formatted text can then be used
//       for text displays, file output or printing.
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
func (searchTestConfigNanobot *charSearchTestConfigDtoNanobot) getParameterTextListing(
	searchTestCfgDto *CharSearchTestConfigDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestConfigDtoNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {

		return strBuilder, err

	}

	if searchTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	// Total available Length of Output Line
	const maxFieldLen = 70

	// Max Label Field Length = 36
	const maxLabelFieldLen = 36

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
	txtFmt.Label.FieldText = "CharSearchTestConfigDto"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	txtStrParam :=
		searchTestCfgDto.TestInputParametersName

	if len(txtStrParam) > 0 {

		// Title Line 2
		txtFmt = TextFormatterDto{}
		txtFmt.FormatType = TxtFieldType.Label()
		txtFmt.Label.LeftMarginStr = ""
		txtFmt.Label.FieldText = txtStrParam
		txtFmt.Label.FieldLength = maxFieldLen
		txtFmt.Label.FieldJustify = TxtJustify.Center()
		txtFmt.Label.RightMarginStr = ""
		txtFmt.Label.LineTerminator = "\n"
		fmtrs = append(fmtrs, txtFmt)

	}

	// Title Line 3
	txtFmt = TextFormatterDto{}
	txtFmt.FormatType = TxtFieldType.Label()
	txtFmt.Label.LeftMarginStr = ""
	txtFmt.Label.FieldText = "Parameter Listing"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	// Title Line 4 Date/Time
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

	err = txtBuilder.BuildTextFormatters(
		&strBuilder,
		fmtrs,
		ePrefix.XCpy(
			"strBuilder<-Marquee Top"))

	if err != nil {

		return strBuilder, err
	}

	fmtrs = nil

	// End Of Marquee

	colonSpace := ": "

	var labelParams []TextLabelValueStrings

	// Build TestInputParametersName
	labelParam := TextLabelValueStrings{}

	labelParam.ParamLabel = "TestInputParametersName"

	labelParam.ParamValue =
		searchTestCfgDto.TestInputParametersName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestInputParametersName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringName"

	labelParam.ParamValue =
		searchTestCfgDto.TestStringName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringLengthName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLengthName"

	labelParam.ParamValue =
		searchTestCfgDto.TestStringLengthName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringLengthName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchTestCfgDto.TestStringStartingIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndexName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndexName"

	labelParam.ParamValue =
		searchTestCfgDto.TestStringStartingIndexName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringStartingIndexName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription1"

	labelParam.ParamValue =
		searchTestCfgDto.TestStringDescription1

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringDescription1 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription2"

	labelParam.ParamValue =
		searchTestCfgDto.TestStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringDescription2 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build CollectionTestObjIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "CollectionTestObjIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchTestCfgDto.CollectionTestObjIndex)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "CollectionTestObjIndex is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build NumValueType

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumValueType"

	labelParam.ParamValue =
		searchTestCfgDto.NumValueType.
			XReturnNoneIfInvalid().String()

	labelParams = append(labelParams, labelParam)

	// Build NumStrFormatType

	labelParam = TextLabelValueStrings{}

	if !searchTestCfgDto.NumStrFormatType.XIsValid() {
		searchTestCfgDto.NumStrFormatType = NumStrFmtType.None()
	}

	labelParam.ParamLabel = "NumStrFormatType"

	labelParam.ParamValue =
		searchTestCfgDto.NumStrFormatType.String()

	labelParams = append(labelParams, labelParam)

	// Build NumSymbolLocation

	labelParam = TextLabelValueStrings{}

	if !searchTestCfgDto.NumSymbolLocation.XIsValid() {
		searchTestCfgDto.NumSymbolLocation = NumSymLocation.None()
	}

	labelParam.ParamLabel = "NumSymbolLocation"

	labelParam.ParamValue =
		searchTestCfgDto.NumSymbolLocation.String()

	labelParams = append(labelParams, labelParam)

	// Build NumSymbolClass

	labelParam = TextLabelValueStrings{}

	if !searchTestCfgDto.NumSymbolClass.XIsValid() {
		searchTestCfgDto.NumSymbolClass = NumSymClass.None()
	}

	labelParam.ParamLabel = "NumSymbolClass"

	labelParam.ParamValue =
		searchTestCfgDto.NumSymbolClass.String()

	labelParams = append(labelParams, labelParam)

	// Build NumSignValue

	labelParam = TextLabelValueStrings{}

	if !searchTestCfgDto.NumSignValue.XIsValid() {
		searchTestCfgDto.NumSignValue = NumSignVal.None()
	}

	labelParam.ParamLabel = "NumSignValue"

	labelParam.ParamValue =
		searchTestCfgDto.NumSignValue.String()

	labelParams = append(labelParams, labelParam)

	// Build PrimaryNumSignPosition

	labelParam = TextLabelValueStrings{}

	if !searchTestCfgDto.PrimaryNumSignPosition.XIsValid() {

		searchTestCfgDto.PrimaryNumSignPosition =
			NumSignSymPos.None()

	}

	labelParam.ParamLabel = "PrimaryNumSignPosition"

	labelParam.ParamValue =
		searchTestCfgDto.PrimaryNumSignPosition.String()

	labelParams = append(labelParams, labelParam)

	// Build SecondaryNumSignPosition

	labelParam = TextLabelValueStrings{}

	if !searchTestCfgDto.SecondaryNumSignPosition.XIsValid() {

		searchTestCfgDto.SecondaryNumSignPosition =
			NumSignSymPos.None()

	}

	labelParam.ParamLabel = "SecondaryNumSignPosition"

	labelParam.ParamValue =
		searchTestCfgDto.SecondaryNumSignPosition.String()

	labelParams = append(labelParams, labelParam)

	// Build TextCharSearchType

	labelParam = TextLabelValueStrings{}

	if !searchTestCfgDto.TextCharSearchType.XIsValid() {

		searchTestCfgDto.TextCharSearchType =
			CharSearchType.None()

	}
	labelParam.ParamLabel = "TextCharSearchType"

	labelParam.ParamValue =
		searchTestCfgDto.TextCharSearchType.String()

	labelParams = append(labelParams, labelParam)

	// Build RequestFoundTestCharacters
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "RequestFoundTestCharacters"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchTestCfgDto.RequestFoundTestCharacters)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "RequestFoundTestCharacters is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build RequestRemainderString
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "RequestRemainderString"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchTestCfgDto.RequestRemainderString)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "RequestRemainderString is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build RequestReplacementString
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "RequestReplacementString"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchTestCfgDto.RequestReplacementString)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "RequestReplacementString is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Label and Parameter values to String Builder
	err = txtBuilder.BuildLabelsValues(
		&strBuilder,
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
	txtFmt.Label.FieldText = "CharSearchTestConfigDto"
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

	err = txtBuilder.BuildTextFormatters(
		&strBuilder,
		fmtrs,
		ePrefix.XCpy(
			"Marquee-Bottom"))
	fmtrs = nil

	return strBuilder, err
}

// ptr - Returns a pointer to a new instance of
// charSearchTestConfigDtoNanobot.
//
func (searchTestConfigNanobot charSearchTestConfigDtoNanobot) ptr() *charSearchTestConfigDtoNanobot {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	return &charSearchTestConfigDtoNanobot{
		lock: new(sync.Mutex),
	}
}

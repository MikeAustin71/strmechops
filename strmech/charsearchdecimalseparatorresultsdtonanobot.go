package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchDecimalSeparatorResultsDtoNanobot - Provides helper
// methods for type CharSearchDecimalSeparatorResultsDto.
//
type charSearchDecimalSeparatorResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceDecSepResults' to input parameter
// 'destinationDecSepResults'. Both instances are of type
// CharSearchDecimalSeparatorResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationDecSepResults'
// will be overwritten.
//
// Also, NO data validation is performed on 'sourceDecSepResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  destinationDecSepRes   *CharSearchDecimalSeparatorResultsDto
//     - A pointer to a CharSearchDecimalSeparatorResultsDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceDecSepResults'.
//
//       'destinationDecSepResults' is the destination for this
//       copy operation.
//
//
//  sourceDecSepResults    *CharSearchDecimalSeparatorResultsDto
//     - A pointer to another CharSearchDecimalSeparatorResultsDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationDecSepResults'.
//
//       'sourceDecSepResults' is the source for this copy
//       operation.
//
//       No data validation is performed on 'sourceDecSepResults'.
//
//
//  errPrefDto             *ePref.ErrPrefixDto
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
func (searchDecimalSepResultsNanobot *charSearchDecimalSeparatorResultsDtoNanobot) copyIn(
	destinationDecSepResults *CharSearchDecimalSeparatorResultsDto,
	sourceDecSepResults *CharSearchDecimalSeparatorResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchDecimalSepResultsNanobot.lock == nil {
		searchDecimalSepResultsNanobot.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsNanobot.lock.Lock()

	defer searchDecimalSepResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchDecimalSeparatorResultsDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationDecSepResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationDecSepResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceDecSepResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceDecSepResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchDecimalSeparatorResultsDtoAtom{}.ptr().empty(
		destinationDecSepResults)

	destinationDecSepResults.SearchResultsName =
		sourceDecSepResults.SearchResultsName

	destinationDecSepResults.SearchResultsFunctionChain =
		sourceDecSepResults.SearchResultsFunctionChain

	destinationDecSepResults.FoundDecimalSeparatorSymbols =
		sourceDecSepResults.FoundDecimalSeparatorSymbols

	destinationDecSepResults.FoundDecimalSepSymbolsOnPreviousSearch =
		sourceDecSepResults.FoundDecimalSeparatorSymbols

	destinationDecSepResults.FoundFirstNumericDigitInNumStr =
		sourceDecSepResults.FoundFirstNumericDigitInNumStr

	destinationDecSepResults.FoundNonZeroValue =
		sourceDecSepResults.FoundNonZeroValue

	destinationDecSepResults.TargetInputParametersName =
		sourceDecSepResults.TargetInputParametersName

	destinationDecSepResults.TargetStringLength =
		sourceDecSepResults.TargetStringLength

	destinationDecSepResults.TargetStringSearchLength =
		sourceDecSepResults.TargetStringSearchLength

	destinationDecSepResults.TargetStringAdjustedSearchLength =
		sourceDecSepResults.TargetStringAdjustedSearchLength

	destinationDecSepResults.TargetStringStartingSearchIndex =
		sourceDecSepResults.TargetStringStartingSearchIndex

	destinationDecSepResults.TargetStringCurrentSearchIndex =
		sourceDecSepResults.TargetStringCurrentSearchIndex

	destinationDecSepResults.TargetStringFirstFoundIndex =
		sourceDecSepResults.TargetStringFirstFoundIndex

	destinationDecSepResults.TargetStringLastFoundIndex =
		sourceDecSepResults.TargetStringLastFoundIndex

	destinationDecSepResults.TargetStringLastSearchIndex =
		sourceDecSepResults.TargetStringLastSearchIndex

	destinationDecSepResults.TargetStringNextSearchIndex =
		sourceDecSepResults.TargetStringNextSearchIndex

	destinationDecSepResults.TargetStringDescription1 =
		sourceDecSepResults.TargetStringDescription1

	destinationDecSepResults.TargetStringDescription2 =
		sourceDecSepResults.TargetStringDescription2

	destinationDecSepResults.TestInputParametersName =
		sourceDecSepResults.TestInputParametersName

	destinationDecSepResults.TestStringName =
		sourceDecSepResults.TestStringName

	destinationDecSepResults.TestStringLength =
		sourceDecSepResults.TestStringLength

	destinationDecSepResults.TestStringLengthName =
		sourceDecSepResults.TestStringLengthName

	destinationDecSepResults.TestStringStartingIndex =
		sourceDecSepResults.TestStringStartingIndex

	destinationDecSepResults.TestStringStartingIndexName =
		sourceDecSepResults.TestStringStartingIndexName

	destinationDecSepResults.TestStringFirstFoundIndex =
		sourceDecSepResults.TestStringFirstFoundIndex

	destinationDecSepResults.TestStringLastFoundIndex =
		sourceDecSepResults.TestStringLastFoundIndex

	destinationDecSepResults.TestStringDescription1 =
		sourceDecSepResults.TestStringDescription1

	destinationDecSepResults.TestStringDescription2 =
		sourceDecSepResults.TestStringDescription2

	destinationDecSepResults.CollectionTestObjIndex =
		sourceDecSepResults.CollectionTestObjIndex

	destinationDecSepResults.NumValueType =
		sourceDecSepResults.NumValueType

	destinationDecSepResults.NumSymbolLocation =
		sourceDecSepResults.NumSymbolLocation

	destinationDecSepResults.TextCharSearchType =
		sourceDecSepResults.TextCharSearchType

	err = destinationDecSepResults.DecimalSeparatorSymbolsSpec.CopyIn(
		&sourceDecSepResults.DecimalSeparatorSymbolsSpec,
		ePrefix.XCpy(
			"destinationDecSepResults."+
				"DecimalSeparatorSymbolsSpec<-sourceDecSepResults"))

	if err != nil {
		return err
	}

	err = destinationDecSepResults.ReplacementString.CopyIn(
		&sourceDecSepResults.ReplacementString,
		ePrefix.XCpy(
			"destinationDecSepResults.ReplacementString"+
				"<-sourceDecSepResults"))

	if err != nil {
		return err
	}

	err = destinationDecSepResults.RemainderString.CopyIn(
		&sourceDecSepResults.RemainderString,
		ePrefix.XCpy(
			"destinationDecSepResults.RemainderString"+
				"<-sourceDecSepResults"))

	if err != nil {
		return err
	}

	err = destinationDecSepResults.FoundRuneArrayChars.CopyIn(
		&sourceDecSepResults.FoundRuneArrayChars,
		ePrefix.XCpy(
			"destinationDecSepResults.FoundRuneArrayChars"+
				"<-sourceDecSepResults"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'decimalSeparatorResults', a pointer to an instance of
// CharSearchDecimalSeparatorResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'decimalSeparatorResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorResults  *CharSearchDecimalSeparatorResultsDto
//     - A pointer to an instance of
//       CharSearchDecimalSeparatorResultsDto. A deep copy of the
//       internal member variables contained in this instance will
//       be created and returned in a new instance of
//       CharSearchDecimalSeparatorResultsDto.
//
//       No data validation is performed on
//       'decimalSeparatorResults'.
//
//
//  errPrefDto               *ePref.ErrPrefixDto
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
//  deepCopyNegNumResultsDto CharSearchDecimalSeparatorResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'decimalSeparatorResults' will be created
//       and returned in a new instance of
//       CharSearchDecimalSeparatorResultsDto.
//
//
//  err                      error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchDecimalSepResultsNanobot *charSearchDecimalSeparatorResultsDtoNanobot) copyOut(
	decimalSeparatorResults *CharSearchDecimalSeparatorResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyDecSepResults CharSearchDecimalSeparatorResultsDto,
	err error) {

	if searchDecimalSepResultsNanobot.lock == nil {
		searchDecimalSepResultsNanobot.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsNanobot.lock.Lock()

	defer searchDecimalSepResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchDecimalSeparatorResultsDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyDecSepResults, err

	}

	if decimalSeparatorResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decimalSeparatorResults' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyDecSepResults, err
	}

	charSearchDecimalSeparatorResultsDtoAtom{}.ptr().empty(
		&deepCopyDecSepResults)

	deepCopyDecSepResults.SearchResultsName =
		decimalSeparatorResults.SearchResultsName

	deepCopyDecSepResults.SearchResultsFunctionChain =
		decimalSeparatorResults.SearchResultsFunctionChain

	deepCopyDecSepResults.FoundDecimalSeparatorSymbols =
		decimalSeparatorResults.FoundDecimalSeparatorSymbols

	deepCopyDecSepResults.FoundDecimalSepSymbolsOnPreviousSearch =
		decimalSeparatorResults.FoundDecimalSeparatorSymbols

	deepCopyDecSepResults.FoundFirstNumericDigitInNumStr =
		decimalSeparatorResults.FoundFirstNumericDigitInNumStr

	deepCopyDecSepResults.FoundNonZeroValue =
		decimalSeparatorResults.FoundNonZeroValue

	deepCopyDecSepResults.TargetInputParametersName =
		decimalSeparatorResults.TargetInputParametersName

	deepCopyDecSepResults.TargetStringLength =
		decimalSeparatorResults.TargetStringLength

	deepCopyDecSepResults.TargetStringSearchLength =
		decimalSeparatorResults.TargetStringSearchLength

	deepCopyDecSepResults.TargetStringAdjustedSearchLength =
		decimalSeparatorResults.TargetStringAdjustedSearchLength

	deepCopyDecSepResults.TargetStringStartingSearchIndex =
		decimalSeparatorResults.TargetStringStartingSearchIndex

	deepCopyDecSepResults.TargetStringCurrentSearchIndex =
		decimalSeparatorResults.TargetStringCurrentSearchIndex

	deepCopyDecSepResults.TargetStringFirstFoundIndex =
		decimalSeparatorResults.TargetStringFirstFoundIndex

	deepCopyDecSepResults.TargetStringLastFoundIndex =
		decimalSeparatorResults.TargetStringLastFoundIndex

	deepCopyDecSepResults.TargetStringLastSearchIndex =
		decimalSeparatorResults.TargetStringLastSearchIndex

	deepCopyDecSepResults.TargetStringNextSearchIndex =
		decimalSeparatorResults.TargetStringNextSearchIndex

	deepCopyDecSepResults.TargetStringDescription1 =
		decimalSeparatorResults.TargetStringDescription1

	deepCopyDecSepResults.TargetStringDescription2 =
		decimalSeparatorResults.TargetStringDescription2

	deepCopyDecSepResults.TestInputParametersName =
		decimalSeparatorResults.TestInputParametersName

	deepCopyDecSepResults.TestStringName =
		decimalSeparatorResults.TestStringName

	deepCopyDecSepResults.TestStringLength =
		decimalSeparatorResults.TestStringLength

	deepCopyDecSepResults.TestStringLengthName =
		decimalSeparatorResults.TestStringLengthName

	deepCopyDecSepResults.TestStringStartingIndex =
		decimalSeparatorResults.TestStringStartingIndex

	deepCopyDecSepResults.TestStringStartingIndexName =
		decimalSeparatorResults.TestStringStartingIndexName

	deepCopyDecSepResults.TestStringFirstFoundIndex =
		decimalSeparatorResults.TestStringFirstFoundIndex

	deepCopyDecSepResults.TestStringLastFoundIndex =
		decimalSeparatorResults.TestStringLastFoundIndex

	deepCopyDecSepResults.TestStringDescription1 =
		decimalSeparatorResults.TestStringDescription1

	deepCopyDecSepResults.TestStringDescription2 =
		decimalSeparatorResults.TestStringDescription2

	deepCopyDecSepResults.CollectionTestObjIndex =
		decimalSeparatorResults.CollectionTestObjIndex

	deepCopyDecSepResults.NumValueType =
		decimalSeparatorResults.NumValueType

	deepCopyDecSepResults.NumSymbolLocation =
		decimalSeparatorResults.NumSymbolLocation

	deepCopyDecSepResults.TextCharSearchType =
		decimalSeparatorResults.TextCharSearchType

	err = deepCopyDecSepResults.DecimalSeparatorSymbolsSpec.CopyIn(
		&decimalSeparatorResults.DecimalSeparatorSymbolsSpec,
		ePrefix.XCpy(
			"deepCopyDecSepResults<-decimalSeparatorResults"))

	if err != nil {
		return deepCopyDecSepResults, err
	}

	err = deepCopyDecSepResults.ReplacementString.CopyIn(
		&decimalSeparatorResults.ReplacementString,
		ePrefix.XCpy(
			"deepCopyDecSepResults.ReplacementString"+
				"<-decimalSeparatorResults"))

	if err != nil {
		return deepCopyDecSepResults, err
	}

	err = deepCopyDecSepResults.RemainderString.CopyIn(
		&decimalSeparatorResults.RemainderString,
		ePrefix.XCpy(
			"deepCopyDecSepResults.RemainderString"+
				"<-decimalSeparatorResults"))

	if err != nil {
		return deepCopyDecSepResults, err
	}

	err = deepCopyDecSepResults.FoundRuneArrayChars.CopyIn(
		&decimalSeparatorResults.FoundRuneArrayChars,
		ePrefix.XCpy(
			"deepCopyDecSepResults.FoundRuneArrayChars"+
				"<-decimalSeparatorResults"))

	return deepCopyDecSepResults, err
}

// getParameterTextListing - Returns formatted text output
// listing the member variable names and corresponding values
// contained in the 'decimalSeparatorResults' instance of
// CharSearchDecimalSeparatorResultsDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  decimalSeparatorResults *CharSearchDecimalSeparatorResultsDto
//     - A pointer to an instance of
//       CharSearchDecimalSeparatorResultsDto instance. Formatted
//       text output will be generated listing the member variable
//       names and their corresponding values. The formatted text
//       can then be used for text displays, file output or
//       printing.
//
//       No data validation is performed on this instance of
//       CharSearchDecimalSeparatorResultsDto.
//
//
//  errPrefDto              *ePref.ErrPrefixDto
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
//       'decimalSeparatorResults' . This formatted text can them be used
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
func (searchDecimalSepResultsNanobot *charSearchDecimalSeparatorResultsDtoNanobot) getParameterTextListing(
	decimalSeparatorResults *CharSearchDecimalSeparatorResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchDecimalSepResultsNanobot.lock == nil {
		searchDecimalSepResultsNanobot.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsNanobot.lock.Lock()

	defer searchDecimalSepResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchDecimalSeparatorResultsDtoNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {

		return strBuilder, err

	}

	if decimalSeparatorResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decimalSeparatorResults' is a nil pointer!\n",
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
	txtFmt.Label.FieldText = "CharSearchDecimalSeparatorResultsDto"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	txtStrParam :=
		decimalSeparatorResults.SearchResultsName

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

	// Title Line  4 Date/Time
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

	// Build SearchResultsName Name
	labelParam := TextLabelValueStrings{}

	labelParam.ParamLabel = "SearchResultsName"

	labelParam.ParamValue =
		decimalSeparatorResults.SearchResultsName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SearchResultsName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build SearchResultsFunctionChain
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "SearchResultsFunctionChain"

	labelParam.ParamValue =
		decimalSeparatorResults.SearchResultsFunctionChain

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SearchResultsFunctionChain is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build FoundDecimalSeparatorSymbols
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundDecimalSeparatorSymbols"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.FoundDecimalSeparatorSymbols)

	labelParams = append(labelParams, labelParam)

	// Build FoundDecimalSepSymbolsOnPreviousSearch
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundDecimalSepSymbolsOnPreviousSearch"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.FoundDecimalSepSymbolsOnPreviousSearch)

	labelParams = append(labelParams, labelParam)

	// Build FoundFirstNumericDigitInNumStr
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundFirstNumericDigitInNumStr"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.FoundFirstNumericDigitInNumStr)

	labelParams = append(labelParams, labelParam)

	// Build FoundNonZeroValue
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundNonZeroValue"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.FoundNonZeroValue)

	labelParams = append(labelParams, labelParam)

	// Build TargetInputParametersName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetInputParametersName"

	labelParam.ParamValue =
		decimalSeparatorResults.TargetInputParametersName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetInputParametersName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringLength)

	labelParams = append(labelParams, labelParam)
	// Build TargetStringSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringSearchLength)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringSearchLength is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringAdjustedSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringAdjustedSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringAdjustedSearchLength)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringAdjustedSearchLength is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringStartingSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringStartingSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringStartingSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringCurrentSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringCurrentSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringCurrentSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringFirstFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringFirstFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringFirstFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLastFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLastFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringLastFoundIndex)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringLastFoundIndex is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLastSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLastSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringLastSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringNextSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringNextSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TargetStringNextSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription1"

	labelParam.ParamValue =
		decimalSeparatorResults.TargetStringDescription1

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringDescription1 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription2"

	labelParam.ParamValue =
		decimalSeparatorResults.TargetStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringDescription2 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestInputParametersName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestInputParametersName"

	labelParam.ParamValue =
		decimalSeparatorResults.TargetStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestInputParametersName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringName"

	labelParam.ParamValue =
		decimalSeparatorResults.TestStringName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TestStringLength)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringLength is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringLengthName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLengthName"

	labelParam.ParamValue =
		decimalSeparatorResults.TestStringLengthName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringLengthName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TestStringStartingIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndexName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndexName"

	labelParam.ParamValue =
		decimalSeparatorResults.TestStringStartingIndexName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringStartingIndexName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringFirstFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringFirstFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TestStringFirstFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringLastFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLastFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TestStringLastFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription1"

	labelParam.ParamValue =
		decimalSeparatorResults.TestStringDescription1

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringDescription1 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription2"

	labelParam.ParamValue =
		decimalSeparatorResults.TestStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringDescription2 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build CollectionTestObjIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "CollectionTestObjIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.CollectionTestObjIndex)

	labelParams = append(labelParams, labelParam)

	// Build NumValueType
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumValueType"

	if !decimalSeparatorResults.NumValueType.XIsValid() {
		decimalSeparatorResults.NumValueType =
			NumValType.None()
	}

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.NumValueType.String())

	labelParams = append(labelParams, labelParam)

	// Build NumSymbolLocation
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumSymbolLocation"

	if !decimalSeparatorResults.NumSymbolLocation.XIsValid() {
		decimalSeparatorResults.NumSymbolLocation =
			NumSymLocation.None()
	}

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.NumValueType.String())

	labelParams = append(labelParams, labelParam)

	// Build TextCharSearchType
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TextCharSearchType"

	if !decimalSeparatorResults.TextCharSearchType.XIsValid() {
		decimalSeparatorResults.TextCharSearchType =
			CharSearchType.None()
	}

	labelParam.ParamValue = fmt.Sprintf("%v",
		decimalSeparatorResults.TextCharSearchType.String())

	labelParams = append(labelParams, labelParam)

	// Build DecimalSeparatorSymbolsSpec
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "DecimalSeparatorSymbolsSpec"

	labelParam.ParamValue =
		decimalSeparatorResults.DecimalSeparatorSymbolsSpec.
			GetDecimalSeparatorStr()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "DecimalSeparatorSymbolsSpec is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build ReplacementString
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "ReplacementString"

	labelParam.ParamValue =
		decimalSeparatorResults.ReplacementString.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "ReplacementString is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build RemainderString
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "RemainderString"

	labelParam.ParamValue =
		decimalSeparatorResults.RemainderString.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "RemainderString is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build FoundRuneArrayChars
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundRuneArrayChars"

	labelParam.ParamValue =
		decimalSeparatorResults.FoundRuneArrayChars.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "FoundRuneArrayChars is EMPTY!"
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

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

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
	txtFmt.Label.FieldText = "CharSearchDecimalSeparatorResultsDto"
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
// charSearchDecimalSeparatorResultsDtoNanobot.
//
func (searchDecimalSepResultsNanobot charSearchDecimalSeparatorResultsDtoNanobot) ptr() *charSearchDecimalSeparatorResultsDtoNanobot {

	if searchDecimalSepResultsNanobot.lock == nil {
		searchDecimalSepResultsNanobot.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsNanobot.lock.Lock()

	defer searchDecimalSepResultsNanobot.lock.Unlock()

	return &charSearchDecimalSeparatorResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

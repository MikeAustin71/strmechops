package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchRuneArrayResultsDtoNanobot - Provides helper methods for
// type charSearchRuneArrayResultsDtoNanobot.
//
type charSearchRuneArrayResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceRuneSearchResults' to input parameter
// 'destinationRuneSearchResults'. Both instances are of type
// CharSearchRuneArrayResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationRuneSearchResults'
// will be overwritten.
//
// Also, NO data validation is performed on 'sourceRuneSearchResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  destinationRuneSearchResults     *CharSearchRuneArrayResultsDto
//     - A pointer to a CharSearchRuneArrayResultsDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceRuneSearchResults'.
//
//       'destinationRuneSearchResults' is the destination for this
//       copy operation.
//
//
//  sourceRuneSearchResults          *CharSearchRuneArrayResultsDto
//     - A pointer to another CharSearchRuneArrayResultsDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationRuneSearchResults'.
//
//       'sourceRuneSearchResults' is the source for this copy
//       operation.
//
//       No data validation is performed on 'sourceRuneSearchResults'.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
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
func (searchRunesResultsDtoNanobot charSearchRuneArrayResultsDtoNanobot) copyIn(
	destinationRuneSearchResults *CharSearchRuneArrayResultsDto,
	sourceRuneSearchResults *CharSearchRuneArrayResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchRunesResultsDtoNanobot.lock == nil {
		searchRunesResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchRunesResultsDtoNanobot.lock.Lock()

	defer searchRunesResultsDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchRuneArrayResultsDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationRuneSearchResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationRuneSearchResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceRuneSearchResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceRuneSearchResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchRuneArrayResultsDtoAtom{}.ptr().
		empty(destinationRuneSearchResults)

	destinationRuneSearchResults.SearchResultsName =
		sourceRuneSearchResults.SearchResultsName

	destinationRuneSearchResults.SearchResultsFunctionChain =
		sourceRuneSearchResults.SearchResultsFunctionChain

	destinationRuneSearchResults.FoundSearchTarget =
		sourceRuneSearchResults.FoundSearchTarget

	destinationRuneSearchResults.FoundFirstNumericDigitInNumStr =
		sourceRuneSearchResults.FoundFirstNumericDigitInNumStr

	destinationRuneSearchResults.FoundDecimalSeparatorSymbols =
		sourceRuneSearchResults.FoundDecimalSeparatorSymbols

	destinationRuneSearchResults.FoundNonZeroValue =
		sourceRuneSearchResults.FoundNonZeroValue

	destinationRuneSearchResults.TargetInputParametersName =
		sourceRuneSearchResults.TargetInputParametersName

	destinationRuneSearchResults.TargetStringLength =
		sourceRuneSearchResults.TargetStringLength

	destinationRuneSearchResults.TargetStringSearchLength =
		sourceRuneSearchResults.TargetStringSearchLength

	destinationRuneSearchResults.TargetStringAdjustedSearchLength =
		sourceRuneSearchResults.TargetStringAdjustedSearchLength

	destinationRuneSearchResults.TargetStringStartingSearchIndex =
		sourceRuneSearchResults.TargetStringStartingSearchIndex

	destinationRuneSearchResults.TargetStringCurrentSearchIndex =
		sourceRuneSearchResults.TargetStringCurrentSearchIndex

	destinationRuneSearchResults.TargetStringFirstFoundIndex =
		sourceRuneSearchResults.TargetStringFirstFoundIndex

	destinationRuneSearchResults.TargetStringLastFoundIndex =
		sourceRuneSearchResults.TargetStringLastFoundIndex

	destinationRuneSearchResults.TargetStringLastSearchIndex =
		sourceRuneSearchResults.TargetStringLastSearchIndex

	destinationRuneSearchResults.TargetStringNextSearchIndex =
		sourceRuneSearchResults.TargetStringNextSearchIndex

	destinationRuneSearchResults.TargetStringDescription1 =
		sourceRuneSearchResults.TargetStringDescription1

	destinationRuneSearchResults.TargetStringDescription2 =
		sourceRuneSearchResults.TargetStringDescription2

	destinationRuneSearchResults.TestInputParametersName =
		sourceRuneSearchResults.TestInputParametersName

	destinationRuneSearchResults.TestStringName =
		sourceRuneSearchResults.TestStringName

	destinationRuneSearchResults.TestStringLength =
		sourceRuneSearchResults.TestStringLength

	destinationRuneSearchResults.TestStringLengthName =
		sourceRuneSearchResults.TestStringLengthName

	destinationRuneSearchResults.TestStringStartingIndex =
		sourceRuneSearchResults.TestStringStartingIndex

	destinationRuneSearchResults.TestStringStartingIndexName =
		sourceRuneSearchResults.TestStringStartingIndexName

	destinationRuneSearchResults.TestStringFirstFoundIndex =
		sourceRuneSearchResults.TestStringFirstFoundIndex

	destinationRuneSearchResults.TestStringLastFoundIndex =
		sourceRuneSearchResults.TestStringLastFoundIndex

	destinationRuneSearchResults.TestStringDescription1 =
		sourceRuneSearchResults.TestStringDescription1

	destinationRuneSearchResults.TestStringDescription2 =
		sourceRuneSearchResults.TestStringDescription2

	destinationRuneSearchResults.CollectionTestObjIndex =
		sourceRuneSearchResults.CollectionTestObjIndex

	destinationRuneSearchResults.TextCharSearchType =
		sourceRuneSearchResults.TextCharSearchType

	err = destinationRuneSearchResults.ReplacementString.CopyIn(
		&sourceRuneSearchResults.ReplacementString,
		ePrefix.XCpy(
			"destinationRuneSearchResults.ReplacementString<-"+
				"sourceRuneSearchResults"))

	if err != nil {
		return err
	}

	err = destinationRuneSearchResults.RemainderString.CopyIn(
		&sourceRuneSearchResults.RemainderString,
		ePrefix.XCpy(
			"destinationRuneSearchResults.RemainderString<-"+
				"sourceRuneSearchResults"))

	if err != nil {
		return err
	}

	err = destinationRuneSearchResults.FoundCharacters.CopyIn(
		&sourceRuneSearchResults.FoundCharacters,
		ePrefix.XCpy(
			"destinationRuneSearchResults.FoundCharacters<-"+
				"sourceRuneSearchResults"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'runeSearchResultsDto', a pointer to an instance of
// CharSearchRuneArrayResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'runeSearchResultsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  runeSearchResultsDto       *CharSearchRuneArrayResultsDto
//     - A pointer to an instance of CharSearchRuneArrayResultsDto.
//       A deep copy of the internal member variables will be
//       created and returned in a new instance of
//       CharSearchRuneArrayResultsDto.
//
//       No data validation is performed on 'runeSearchResultsDto'.
//
//
//  errPrefDto                      *ePref.ErrPrefixDto
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
//  deepCopyRuneSearchResults  CharSearchRuneArrayResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'runeSearchResultsDto' will be created and
//       returned in a new instance of
//       CharSearchRuneArrayResultsDto.
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
func (searchRunesResultsDtoNanobot charSearchRuneArrayResultsDtoNanobot) copyOut(
	runeSearchResultsDto *CharSearchRuneArrayResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyRuneSearchResults CharSearchRuneArrayResultsDto,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchRuneArrayResultsDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyRuneSearchResults, err

	}

	if runeSearchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyRuneSearchResults, err
	}

	charSearchRuneArrayResultsDtoAtom{}.ptr().
		empty(&deepCopyRuneSearchResults)

	deepCopyRuneSearchResults.SearchResultsName =
		runeSearchResultsDto.SearchResultsName

	deepCopyRuneSearchResults.SearchResultsFunctionChain =
		runeSearchResultsDto.SearchResultsFunctionChain

	deepCopyRuneSearchResults.FoundSearchTarget =
		runeSearchResultsDto.FoundSearchTarget

	deepCopyRuneSearchResults.FoundFirstNumericDigitInNumStr =
		runeSearchResultsDto.FoundFirstNumericDigitInNumStr

	deepCopyRuneSearchResults.FoundDecimalSeparatorSymbols =
		runeSearchResultsDto.FoundDecimalSeparatorSymbols

	deepCopyRuneSearchResults.FoundNonZeroValue =
		runeSearchResultsDto.FoundNonZeroValue

	deepCopyRuneSearchResults.TargetInputParametersName =
		runeSearchResultsDto.TargetInputParametersName

	deepCopyRuneSearchResults.TargetStringLength =
		runeSearchResultsDto.TargetStringLength

	deepCopyRuneSearchResults.TargetStringSearchLength =
		runeSearchResultsDto.TargetStringSearchLength

	deepCopyRuneSearchResults.TargetStringAdjustedSearchLength =
		runeSearchResultsDto.TargetStringAdjustedSearchLength

	deepCopyRuneSearchResults.TargetStringStartingSearchIndex =
		runeSearchResultsDto.TargetStringStartingSearchIndex

	deepCopyRuneSearchResults.TargetStringCurrentSearchIndex =
		runeSearchResultsDto.TargetStringCurrentSearchIndex

	deepCopyRuneSearchResults.TargetStringFirstFoundIndex =
		runeSearchResultsDto.TargetStringFirstFoundIndex

	deepCopyRuneSearchResults.TargetStringLastFoundIndex =
		runeSearchResultsDto.TargetStringLastFoundIndex

	deepCopyRuneSearchResults.TargetStringLastSearchIndex =
		runeSearchResultsDto.TargetStringLastSearchIndex

	deepCopyRuneSearchResults.TargetStringNextSearchIndex =
		runeSearchResultsDto.TargetStringNextSearchIndex

	deepCopyRuneSearchResults.TargetStringDescription1 =
		runeSearchResultsDto.TargetStringDescription1

	deepCopyRuneSearchResults.TargetStringDescription2 =
		runeSearchResultsDto.TargetStringDescription2

	deepCopyRuneSearchResults.TestInputParametersName =
		runeSearchResultsDto.TestInputParametersName

	deepCopyRuneSearchResults.TestStringName =
		runeSearchResultsDto.TestStringName

	deepCopyRuneSearchResults.TestStringLength =
		runeSearchResultsDto.TestStringLength

	deepCopyRuneSearchResults.TestStringLengthName =
		runeSearchResultsDto.TestStringLengthName

	deepCopyRuneSearchResults.TestStringStartingIndex =
		runeSearchResultsDto.TestStringStartingIndex

	deepCopyRuneSearchResults.TestStringStartingIndexName =
		runeSearchResultsDto.TestStringStartingIndexName

	deepCopyRuneSearchResults.TestStringFirstFoundIndex =
		runeSearchResultsDto.TestStringFirstFoundIndex

	deepCopyRuneSearchResults.TestStringLastFoundIndex =
		runeSearchResultsDto.TestStringLastFoundIndex

	deepCopyRuneSearchResults.TestStringDescription1 =
		runeSearchResultsDto.TestStringDescription1

	deepCopyRuneSearchResults.TestStringDescription2 =
		runeSearchResultsDto.TestStringDescription2

	deepCopyRuneSearchResults.CollectionTestObjIndex =
		runeSearchResultsDto.CollectionTestObjIndex

	deepCopyRuneSearchResults.TextCharSearchType =
		runeSearchResultsDto.TextCharSearchType

	err = deepCopyRuneSearchResults.ReplacementString.CopyIn(
		&runeSearchResultsDto.ReplacementString,
		ePrefix.XCpy(
			"deepCopyRuneSearchResults.ReplacementString<-"+
				"runeSearchResultsDto"))

	if err != nil {
		return deepCopyRuneSearchResults, err
	}

	err = deepCopyRuneSearchResults.RemainderString.CopyIn(
		&runeSearchResultsDto.RemainderString,
		ePrefix.XCpy(
			"deepCopyRuneSearchResults.RemainderString<-"+
				"runeSearchResultsDto"))

	if err != nil {
		return deepCopyRuneSearchResults, err
	}

	err = deepCopyRuneSearchResults.FoundCharacters.CopyIn(
		&runeSearchResultsDto.FoundCharacters,
		ePrefix.XCpy(
			"deepCopyRuneSearchResults.FoundCharacters<-"+
				"runeSearchResultsDto"))

	return deepCopyRuneSearchResults, err
}

// getParameterTextListing - Returns formatted text output
// detailing the member variable names and corresponding values
// contained in the 'runeSearchResultsDto' instance of
// CharSearchRuneArrayResultsDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  runeSearchResultsDto       *CharSearchRuneArrayResultsDto
//     - A pointer to an instance of CharSearchRuneArrayResultsDto. Formatted
//       text output will be generated listing the member variable
//       names and their corresponding values. The formatted text
//       can then be used for text displays, file output or
//       printing.
//
//       No data validation is performed on this instance of
//       CharSearchRuneArrayResultsDto.
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
//       'runeSearchResultsDto' . This formatted text can them be used
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
func (searchRunesResultsDtoNanobot charSearchRuneArrayResultsDtoNanobot) getParameterTextListing(
	runeSearchResultsDto *CharSearchRuneArrayResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchRuneArrayResultsDtoNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {

		return strBuilder, err

	}

	if runeSearchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	// Total available Length of Output Line
	const maxFieldLen = 70

	// Max Label Field Length = 24
	const maxLabelFieldLen = 24

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
	txtFmt.Label.FieldText = "CharSearchRuneArrayResultsDto"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	txtStrParam :=
		runeSearchResultsDto.SearchResultsName

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

	// Begin Label Parameter Pairs

	colonSpace := ": "

	var labelParams []TextLabelValueStrings

	// Build SearchResultsName Name
	labelParam := TextLabelValueStrings{}

	labelParam.ParamLabel = "SearchResultsName"

	labelParam.ParamValue =
		runeSearchResultsDto.SearchResultsName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SearchResultsName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build SearchResultsFunctionChain
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "SearchResultsFunctionChain"

	labelParam.ParamValue =
		runeSearchResultsDto.SearchResultsFunctionChain

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SearchResultsFunctionChain is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build FoundSearchTarget
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundSearchTarget"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.FoundSearchTarget)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "FoundSearchTarget is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build FoundFirstNumericDigitInNumStr
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundFirstNumericDigitInNumStr"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.FoundFirstNumericDigitInNumStr)

	labelParams = append(labelParams, labelParam)

	// Build FoundDecimalSeparatorSymbols
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundDecimalSeparatorSymbols"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.FoundDecimalSeparatorSymbols)

	labelParams = append(labelParams, labelParam)

	// Build FoundNonZeroValue
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundNonZeroValue"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.FoundNonZeroValue)

	labelParams = append(labelParams, labelParam)

	// Build TargetInputParametersName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetInputParametersName"

	labelParam.ParamValue =
		runeSearchResultsDto.TargetInputParametersName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetInputParametersName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TargetStringLength)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TargetStringSearchLength)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringAdjustedSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringAdjustedSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TargetStringAdjustedSearchLength)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringStartingSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringStartingSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TargetStringStartingSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringCurrentSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringCurrentSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TargetStringCurrentSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringFirstFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringFirstFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TargetStringFirstFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLastSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLastSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TargetStringLastSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringNextSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringNextSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TargetStringNextSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription1"

	labelParam.ParamValue =
		runeSearchResultsDto.TargetStringDescription1

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringDescription1 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription2"

	labelParam.ParamValue =
		runeSearchResultsDto.TargetStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringDescription2 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestInputParametersName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestInputParametersName"

	labelParam.ParamValue =
		runeSearchResultsDto.TestInputParametersName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestInputParametersName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringName"

	labelParam.ParamValue =
		runeSearchResultsDto.TestStringName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TestStringLength)

	labelParams = append(labelParams, labelParam)

	// Build TestStringLengthName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLengthName"

	labelParam.ParamValue =
		runeSearchResultsDto.TestStringLengthName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringLengthName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TestStringStartingIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndexName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndexName"

	labelParam.ParamValue =
		runeSearchResultsDto.TestStringStartingIndexName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringStartingIndexName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringFirstFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringFirstFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TestStringFirstFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringLastFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLastFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TestStringLastFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription1"

	labelParam.ParamValue =
		runeSearchResultsDto.TestStringDescription1

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringDescription1 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription2"

	labelParam.ParamValue =
		runeSearchResultsDto.TestStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringDescription2 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build CollectionTestObjIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "CollectionTestObjIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.CollectionTestObjIndex)

	labelParams = append(labelParams, labelParam)

	// Build TextCharSearchType
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TextCharSearchType"

	if !runeSearchResultsDto.TextCharSearchType.XIsValid() {
		runeSearchResultsDto.TextCharSearchType =
			CharSearchType.None()
	}

	labelParam.ParamValue = fmt.Sprintf("%v",
		runeSearchResultsDto.TextCharSearchType.String())

	labelParams = append(labelParams, labelParam)

	// Build ReplacementString
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "ReplacementString"

	labelParam.ParamValue =
		runeSearchResultsDto.ReplacementString.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "ReplacementString is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build RemainderString
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "RemainderString"

	labelParam.ParamValue =
		runeSearchResultsDto.RemainderString.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "RemainderString is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build FoundCharacters
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundCharacters"

	labelParam.ParamValue =
		runeSearchResultsDto.FoundCharacters.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "FoundCharacters is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Write Label/Parameter Values to String Builder
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
	txtFmt.Label.FieldText = "CharSearchRuneArrayResultsDto"
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
// charSearchResultsDtoNanobot.
//
func (searchRunesResultsDtoNanobot charSearchRuneArrayResultsDtoNanobot) ptr() *charSearchRuneArrayResultsDtoNanobot {

	if searchRunesResultsDtoNanobot.lock == nil {
		searchRunesResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchRunesResultsDtoNanobot.lock.Lock()

	defer searchRunesResultsDtoNanobot.lock.Unlock()

	return &charSearchRuneArrayResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

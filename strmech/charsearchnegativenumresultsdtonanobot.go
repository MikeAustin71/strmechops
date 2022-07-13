package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchNegNumResultsDtoNanobot - Provides helper
// methods for type CharSearchNegativeNumberResultsDto.
//
type charSearchNegNumResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceNegNumResults' to input parameter
// 'destinationNegNumResults'. Both instances are of type
// CharSearchNegativeNumberResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationNegNumResults'
// will be overwritten.
//
// Also, NO data validation is performed on 'sourceNegNumResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  destinationNegNumResults     *CharSearchNegativeNumberResultsDto
//     - A pointer to a CharSearchNegativeNumberResultsDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceNegNumResults'.
//
//       'destinationNegNumResults' is the destination for this
//       copy operation.
//
//
//  sourceNegNumResults          *CharSearchNegativeNumberResultsDto
//     - A pointer to another CharSearchNegativeNumberResultsDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationNegNumResults'.
//
//       'sourceNegNumResults' is the source for this copy
//       operation.
//
//       No data validation is performed on 'sourceNegNumResults'.
//
//
//  errPrefDto                   *ePref.ErrPrefixDto
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
func (searchNegNumResultsNanobot *charSearchNegNumResultsDtoNanobot) copyIn(
	destinationNegNumResults *CharSearchNegativeNumberResultsDto,
	sourceNegNumResults *CharSearchNegativeNumberResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchNegNumResultsNanobot.lock == nil {
		searchNegNumResultsNanobot.lock = new(sync.Mutex)
	}

	searchNegNumResultsNanobot.lock.Lock()

	defer searchNegNumResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNegNumResultsDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationNegNumResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationNegNumResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceNegNumResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceNegNumResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchNegativeNumberResultsDtoAtom{}.ptr().
		empty(destinationNegNumResults)

	destinationNegNumResults.SearchResultsName =
		sourceNegNumResults.SearchResultsName

	destinationNegNumResults.SearchResultsFunctionChain =
		sourceNegNumResults.SearchResultsFunctionChain

	destinationNegNumResults.FoundNegativeNumberSymbols =
		sourceNegNumResults.FoundNegativeNumberSymbols

	destinationNegNumResults.FoundNegNumSymbolsOnPreviousSearch =
		sourceNegNumResults.FoundNegNumSymbolsOnPreviousSearch

	destinationNegNumResults.FoundLeadingNegNumSymbols =
		sourceNegNumResults.FoundLeadingNegNumSymbols

	destinationNegNumResults.FoundTrailingNegNumSymbols =
		sourceNegNumResults.FoundTrailingNegNumSymbols

	destinationNegNumResults.FoundFirstNumericDigitInNumStr =
		sourceNegNumResults.FoundFirstNumericDigitInNumStr

	destinationNegNumResults.FoundDecimalSeparatorSymbols =
		sourceNegNumResults.FoundDecimalSeparatorSymbols

	destinationNegNumResults.FoundNonZeroValue =
		sourceNegNumResults.FoundNonZeroValue

	destinationNegNumResults.TargetInputParametersName =
		sourceNegNumResults.TargetInputParametersName

	destinationNegNumResults.TargetStringLength =
		sourceNegNumResults.TargetStringLength

	destinationNegNumResults.TargetStringSearchLength =
		sourceNegNumResults.TargetStringSearchLength

	destinationNegNumResults.TargetStringAdjustedSearchLength =
		sourceNegNumResults.TargetStringAdjustedSearchLength

	destinationNegNumResults.TargetStringStartingSearchIndex =
		sourceNegNumResults.TargetStringStartingSearchIndex

	destinationNegNumResults.TargetStringCurrentSearchIndex =
		sourceNegNumResults.TargetStringCurrentSearchIndex

	destinationNegNumResults.TargetStringFirstFoundIndex =
		sourceNegNumResults.TargetStringFirstFoundIndex

	destinationNegNumResults.TargetStringLastFoundIndex =
		sourceNegNumResults.TargetStringLastFoundIndex

	destinationNegNumResults.TargetStringLastSearchIndex =
		sourceNegNumResults.TargetStringLastSearchIndex

	destinationNegNumResults.TargetStringNextSearchIndex =
		sourceNegNumResults.TargetStringNextSearchIndex

	destinationNegNumResults.TargetStringDescription1 =
		sourceNegNumResults.TargetStringDescription1

	destinationNegNumResults.TargetStringDescription2 =
		sourceNegNumResults.TargetStringDescription2

	destinationNegNumResults.TestInputParametersName =
		sourceNegNumResults.TestInputParametersName

	destinationNegNumResults.TestStringName =
		sourceNegNumResults.TestStringName

	destinationNegNumResults.TestStringLength =
		sourceNegNumResults.TestStringLength

	destinationNegNumResults.TestStringLengthName =
		sourceNegNumResults.TestStringLengthName

	destinationNegNumResults.TestStringStartingIndex =
		sourceNegNumResults.TestStringStartingIndex

	destinationNegNumResults.TestStringStartingIndexName =
		sourceNegNumResults.TestStringStartingIndexName

	destinationNegNumResults.TestStringFirstFoundIndex =
		sourceNegNumResults.TestStringFirstFoundIndex

	destinationNegNumResults.TestStringLastFoundIndex =
		sourceNegNumResults.TestStringLastFoundIndex

	destinationNegNumResults.TestStringDescription1 =
		sourceNegNumResults.TestStringDescription1

	destinationNegNumResults.TestStringDescription2 =
		sourceNegNumResults.TestStringDescription2

	destinationNegNumResults.CollectionTestObjIndex =
		sourceNegNumResults.CollectionTestObjIndex

	destinationNegNumResults.NumSignValue =
		sourceNegNumResults.NumSignValue

	destinationNegNumResults.PrimaryNumSignPosition =
		sourceNegNumResults.PrimaryNumSignPosition

	destinationNegNumResults.SecondaryNumSignPosition =
		sourceNegNumResults.SecondaryNumSignPosition

	destinationNegNumResults.TextCharSearchType =
		sourceNegNumResults.TextCharSearchType

	err = destinationNegNumResults.NegativeNumberSymbolsSpec.CopyIn(
		&sourceNegNumResults.NegativeNumberSymbolsSpec,
		ePrefix.XCpy(
			"destinationNegNumResults."+
				"NegativeNumberSymbolsSpec<-"+
				"sourceNegNumResults"))

	if err != nil {
		return err
	}

	err = destinationNegNumResults.ReplacementString.CopyIn(
		&sourceNegNumResults.ReplacementString,
		ePrefix.XCpy(
			"destinationNegNumResults."+
				"ReplacementString<-"+
				"sourceNegNumResults"))

	if err != nil {
		return err
	}

	err = destinationNegNumResults.RemainderString.CopyIn(
		&sourceNegNumResults.RemainderString,
		ePrefix.XCpy(
			"destinationNegNumResults."+
				"RemainderString<-"+
				"sourceNegNumResults"))

	if err != nil {
		return err
	}

	err = destinationNegNumResults.FoundRuneArrayChars.CopyIn(
		&sourceNegNumResults.FoundRuneArrayChars,
		ePrefix.XCpy(
			"destinationNegNumResults."+
				"RemainderString<-"+
				"sourceNegNumResults"))

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'searchNegNumResults', a pointer to an instance of
// CharSearchNegativeNumberResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'searchNegNumResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchNegNumResults        *CharSearchNegativeNumberResultsDto
//     - A pointer to an instance of
//       CharSearchNegativeNumberResultsDto. A deep copy of the
//       internal member variables contained in this instance will
//       be created and returned in a new instance of
//       CharSearchNegativeNumberResultsDto.
//
//       No data validation is performed on 'searchNegNumResults'.
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
//  deepCopyNegNumResultsDto        CharSearchNegativeNumberResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'searchNegNumResults' will be created and
//       returned in a new instance of
//       CharSearchNegativeNumberResultsDto.
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
func (searchNegNumResultsNanobot *charSearchNegNumResultsDtoNanobot) copyOut(
	searchNegNumResults *CharSearchNegativeNumberResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyNegNumResultsDto CharSearchNegativeNumberResultsDto,
	err error) {

	if searchNegNumResultsNanobot.lock == nil {
		searchNegNumResultsNanobot.lock = new(sync.Mutex)
	}

	searchNegNumResultsNanobot.lock.Lock()

	defer searchNegNumResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNegNumResultsDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyNegNumResultsDto, err

	}

	if searchNegNumResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchNegNumResults' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyNegNumResultsDto, err
	}

	charSearchNegativeNumberResultsDtoAtom{}.ptr().
		empty(&deepCopyNegNumResultsDto)

	deepCopyNegNumResultsDto.SearchResultsName =
		searchNegNumResults.SearchResultsName

	deepCopyNegNumResultsDto.SearchResultsFunctionChain =
		searchNegNumResults.SearchResultsFunctionChain

	deepCopyNegNumResultsDto.FoundNegativeNumberSymbols =
		searchNegNumResults.FoundNegativeNumberSymbols

	deepCopyNegNumResultsDto.FoundNegNumSymbolsOnPreviousSearch =
		searchNegNumResults.FoundNegNumSymbolsOnPreviousSearch

	deepCopyNegNumResultsDto.FoundLeadingNegNumSymbols =
		searchNegNumResults.FoundLeadingNegNumSymbols

	deepCopyNegNumResultsDto.FoundTrailingNegNumSymbols =
		searchNegNumResults.FoundTrailingNegNumSymbols

	deepCopyNegNumResultsDto.FoundFirstNumericDigitInNumStr =
		searchNegNumResults.FoundFirstNumericDigitInNumStr

	deepCopyNegNumResultsDto.FoundDecimalSeparatorSymbols =
		searchNegNumResults.FoundDecimalSeparatorSymbols

	deepCopyNegNumResultsDto.FoundNonZeroValue =
		searchNegNumResults.FoundNonZeroValue

	deepCopyNegNumResultsDto.TargetInputParametersName =
		searchNegNumResults.TargetInputParametersName

	deepCopyNegNumResultsDto.TargetStringLength =
		searchNegNumResults.TargetStringLength

	deepCopyNegNumResultsDto.TargetStringSearchLength =
		searchNegNumResults.TargetStringSearchLength

	deepCopyNegNumResultsDto.TargetStringAdjustedSearchLength =
		searchNegNumResults.TargetStringAdjustedSearchLength

	deepCopyNegNumResultsDto.TargetStringStartingSearchIndex =
		searchNegNumResults.TargetStringStartingSearchIndex

	deepCopyNegNumResultsDto.TargetStringCurrentSearchIndex =
		searchNegNumResults.TargetStringCurrentSearchIndex

	deepCopyNegNumResultsDto.TargetStringFirstFoundIndex =
		searchNegNumResults.TargetStringFirstFoundIndex

	deepCopyNegNumResultsDto.TargetStringLastFoundIndex =
		searchNegNumResults.TargetStringLastFoundIndex

	deepCopyNegNumResultsDto.TargetStringLastSearchIndex =
		searchNegNumResults.TargetStringLastSearchIndex

	deepCopyNegNumResultsDto.TargetStringNextSearchIndex =
		searchNegNumResults.TargetStringNextSearchIndex

	deepCopyNegNumResultsDto.TargetStringDescription1 =
		searchNegNumResults.TargetStringDescription1

	deepCopyNegNumResultsDto.TargetStringDescription2 =
		searchNegNumResults.TargetStringDescription2

	deepCopyNegNumResultsDto.TestInputParametersName =
		searchNegNumResults.TestInputParametersName

	deepCopyNegNumResultsDto.TestStringName =
		searchNegNumResults.TestStringName

	deepCopyNegNumResultsDto.TestStringLength =
		searchNegNumResults.TestStringLength

	deepCopyNegNumResultsDto.TestStringLengthName =
		searchNegNumResults.TestStringLengthName

	deepCopyNegNumResultsDto.TestStringStartingIndex =
		searchNegNumResults.TestStringStartingIndex

	deepCopyNegNumResultsDto.TestStringStartingIndexName =
		searchNegNumResults.TestStringStartingIndexName

	deepCopyNegNumResultsDto.TestStringFirstFoundIndex =
		searchNegNumResults.TestStringFirstFoundIndex

	deepCopyNegNumResultsDto.TestStringLastFoundIndex =
		searchNegNumResults.TestStringLastFoundIndex

	deepCopyNegNumResultsDto.TestStringDescription1 =
		searchNegNumResults.TestStringDescription1

	deepCopyNegNumResultsDto.TestStringDescription2 =
		searchNegNumResults.TestStringDescription2

	deepCopyNegNumResultsDto.CollectionTestObjIndex =
		searchNegNumResults.CollectionTestObjIndex

	deepCopyNegNumResultsDto.NumSignValue =
		searchNegNumResults.NumSignValue

	deepCopyNegNumResultsDto.PrimaryNumSignPosition =
		searchNegNumResults.PrimaryNumSignPosition

	deepCopyNegNumResultsDto.SecondaryNumSignPosition =
		searchNegNumResults.SecondaryNumSignPosition

	deepCopyNegNumResultsDto.TextCharSearchType =
		searchNegNumResults.TextCharSearchType

	err = deepCopyNegNumResultsDto.NegativeNumberSymbolsSpec.CopyIn(
		&searchNegNumResults.NegativeNumberSymbolsSpec,
		ePrefix.XCpy(
			"deepCopyNegNumResultsDto."+
				"NegativeNumberSymbolsSpec<-"+
				"searchNegNumResults"))

	if err != nil {

		return deepCopyNegNumResultsDto, err
	}

	err = deepCopyNegNumResultsDto.ReplacementString.CopyIn(
		&searchNegNumResults.ReplacementString,
		ePrefix.XCpy(
			"deepCopyNegNumResultsDto."+
				"ReplacementString<-"+
				"searchNegNumResults"))

	if err != nil {

		return deepCopyNegNumResultsDto, err
	}

	err = deepCopyNegNumResultsDto.RemainderString.CopyIn(
		&searchNegNumResults.RemainderString,
		ePrefix.XCpy(
			"deepCopyNegNumResultsDto."+
				"RemainderString<-"+
				"searchNegNumResults"))

	if err != nil {

		return deepCopyNegNumResultsDto, err
	}

	err = deepCopyNegNumResultsDto.FoundRuneArrayChars.CopyIn(
		&searchNegNumResults.FoundRuneArrayChars,
		ePrefix.XCpy(
			"deepCopyNegNumResultsDto."+
				"FoundRuneArrayChars<-"+
				"searchNegNumResults"))

	if err != nil {

		return deepCopyNegNumResultsDto, err
	}

	return deepCopyNegNumResultsDto, err
}

// getParameterTextListing - Returns formatted text output
// listing the member variable names and corresponding values
// contained in the 'searchNegNumResults' instance of
// CharSearchNegativeNumberResultsDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchNegNumResults *CharSearchNegativeNumberResultsDto
//     - A pointer to an instance of
//       CharSearchNegativeNumberResultsDto instance. Formatted
//       text output will be generated listing the member variable
//       names and their corresponding values. The formatted text
//       can then be used for text displays, file output or
//       printing.
//
//       No data validation is performed on this instance of
//       CharSearchNegativeNumberResultsDto.
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
//  strings.Builder
//     - If this method completes successfully, an instance of
//       strings.Builder will be returned. This instance contains
//       the formatted text output listing the member variable
//       names and their corresponding values for input parameter
//       'searchNegNumResults' . This formatted text can them be used
//       for screen displays, file output or printing.
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
func (searchNegNumResultsNanobot *charSearchNegNumResultsDtoNanobot) getParameterTextListing(
	searchNegNumResults *CharSearchNegativeNumberResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchNegNumResultsNanobot.lock == nil {
		searchNegNumResultsNanobot.lock = new(sync.Mutex)
	}

	searchNegNumResultsNanobot.lock.Lock()

	defer searchNegNumResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNegNumResultsDtoNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {

		return strBuilder, err

	}

	if searchNegNumResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchNegNumResults' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	// Total available Length of Output Line
	const maxFieldLen = 79

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
	txtFmt.Label.FieldText = "CharSearchNegativeNumberResultsDto"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	txtStrParam :=
		searchNegNumResults.SearchResultsName

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
		searchNegNumResults.SearchResultsName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SearchResultsName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build SearchResultsFunctionChain
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "SearchResultsFunctionChain"

	labelParam.ParamValue =
		searchNegNumResults.SearchResultsFunctionChain

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SearchResultsFunctionChain is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build FoundNegativeNumberSymbols
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundNegativeNumberSymbols"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.FoundNegativeNumberSymbols)

	labelParams = append(labelParams, labelParam)

	// Build FoundNegNumSymbolsOnPreviousSearch
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundNegNumSymbolsOnPreviousSearch"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.FoundNegNumSymbolsOnPreviousSearch)

	labelParams = append(labelParams, labelParam)

	// Build FoundLeadingNegNumSymbols
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundLeadingNegNumSymbols"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.FoundLeadingNegNumSymbols)

	labelParams = append(labelParams, labelParam)

	// Build FoundTrailingNegNumSymbols
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundTrailingNegNumSymbols"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.FoundTrailingNegNumSymbols)

	labelParams = append(labelParams, labelParam)

	// Build FoundFirstNumericDigitInNumStr
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundFirstNumericDigitInNumStr"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.FoundFirstNumericDigitInNumStr)

	labelParams = append(labelParams, labelParam)

	// Build FoundDecimalSeparatorSymbols
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundDecimalSeparatorSymbols"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.FoundDecimalSeparatorSymbols)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "FoundDecimalSeparatorSymbols is EMPTY!"
	}

	// Build FoundNonZeroValue
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundNonZeroValue"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.FoundNonZeroValue)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "FoundNonZeroValue is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetInputParametersName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetInputParametersName"

	labelParam.ParamValue =
		searchNegNumResults.TargetInputParametersName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetInputParametersName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringLength)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringSearchLength)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringAdjustedSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringAdjustedSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringAdjustedSearchLength)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringAdjustedSearchLength is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringStartingSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringStartingSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringStartingSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringCurrentSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringCurrentSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringCurrentSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringFirstFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringFirstFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringFirstFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLastFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLastFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringLastFoundIndex)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringLastFoundIndex is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLastSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLastSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringLastSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringNextSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringNextSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TargetStringNextSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription1"

	labelParam.ParamValue =
		searchNegNumResults.TargetStringDescription1

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringDescription1 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TargetStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription2"

	labelParam.ParamValue =
		searchNegNumResults.TargetStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TargetStringDescription2 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestInputParametersName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestInputParametersName"

	labelParam.ParamValue =
		searchNegNumResults.TargetStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestInputParametersName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringName"

	labelParam.ParamValue =
		searchNegNumResults.TestStringName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TestStringLength)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringLength is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringLengthName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLengthName"

	labelParam.ParamValue =
		searchNegNumResults.TestStringLengthName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringLengthName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TestStringStartingIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndexName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndexName"

	labelParam.ParamValue =
		searchNegNumResults.TestStringStartingIndexName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringStartingIndexName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringFirstFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringFirstFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TestStringFirstFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringLastFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLastFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TestStringLastFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription1"

	labelParam.ParamValue =
		searchNegNumResults.TestStringDescription1

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringDescription1 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription2"

	labelParam.ParamValue =
		searchNegNumResults.TestStringDescription2

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestStringDescription2 is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build CollectionTestObjIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "CollectionTestObjIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.CollectionTestObjIndex)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "CollectionTestObjIndex is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build NumSignValue
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumSignValue"

	if !searchNegNumResults.NumSignValue.XIsValid() {
		searchNegNumResults.NumSignValue =
			NumSignVal.None()
	}

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.NumSignValue.String())

	labelParams = append(labelParams, labelParam)

	// Build PrimaryNumSignPosition
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "PrimaryNumSignPosition"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.PrimaryNumSignPosition)

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "PrimaryNumSignPosition is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build SecondaryNumSignPosition
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "SecondaryNumSignPosition"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.SecondaryNumSignPosition.String())

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SecondaryNumSignPosition is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build TextCharSearchType
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TextCharSearchType"

	if !searchNegNumResults.TextCharSearchType.XIsValid() {
		searchNegNumResults.TextCharSearchType =
			CharSearchType.None()
	}

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchNegNumResults.TextCharSearchType.String())

	labelParams = append(labelParams, labelParam)

	// Build NegativeNumberSymbolsSpec
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NegativeNumberSymbolsSpec"

	if searchNegNumResults.NegativeNumberSymbolsSpec.
		IsValidInstance() {
		labelParam.ParamValue = "NegativeNumberSymbolsSpec is populated and valid."
	} else {
		labelParam.ParamValue = "NegativeNumberSymbolsSpec is invalid and empty."
	}

	labelParams = append(labelParams, labelParam)

	// Build ReplacementString
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "ReplacementString"

	labelParam.ParamValue =
		searchNegNumResults.ReplacementString.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "ReplacementString is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build RemainderString
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "RemainderString"

	labelParam.ParamValue =
		searchNegNumResults.RemainderString.GetCharacterString()

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "RemainderString is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build FoundRuneArrayChars
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundRuneArrayChars"

	labelParam.ParamValue =
		searchNegNumResults.FoundRuneArrayChars.GetCharacterString()

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
	txtFmt.Label.FieldText = "CharSearchNegativeNumberResultsDto"
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

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	fmtrs = nil

	return strBuilder, err
}

// ptr - Returns a pointer to a new instance of
// charSearchNegNumResultsDtoNanobot.
//
func (searchNegNumResultsNanobot charSearchNegNumResultsDtoNanobot) ptr() *charSearchNegNumResultsDtoNanobot {

	if searchNegNumResultsNanobot.lock == nil {
		searchNegNumResultsNanobot.lock = new(sync.Mutex)
	}

	searchNegNumResultsNanobot.lock.Lock()

	defer searchNegNumResultsNanobot.lock.Unlock()

	return &charSearchNegNumResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

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

	destinationRuneSearchResults.IsNOP =
		sourceRuneSearchResults.IsNOP

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

	err = destinationRuneSearchResults.FoundRuneArrayChars.CopyIn(
		&sourceRuneSearchResults.FoundRuneArrayChars,
		ePrefix.XCpy(
			"destinationRuneSearchResults.FoundRuneArrayChars<-"+
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

	deepCopyRuneSearchResults.IsNOP =
		runeSearchResultsDto.IsNOP

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

	err = deepCopyRuneSearchResults.FoundRuneArrayChars.CopyIn(
		&runeSearchResultsDto.FoundRuneArrayChars,
		ePrefix.XCpy(
			"deepCopyRuneSearchResults.FoundRuneArrayChars<-"+
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
//  strBuilder                 *strings.Builder
//     - A pointer to an instance of *strings.Builder. The
//       formatted text characters produced by this method will be
//       written to this instance of strings.Builder.
//
//
//  runeSearchResultsDto       *CharSearchRuneArrayResultsDto
//     - A pointer to an instance of CharSearchRuneArrayResultsDto.
//       Formatted text output will be generated listing the member
//       variable names and their corresponding values. The
//       formatted text can then be used for screen displays, file
//       output or printing.
//
//       No data validation is performed on this instance of
//       CharSearchRuneArrayResultsDto.
//
//
//  displayFunctionChain       bool
//     - Set 'displayFunctionChain' to 'true' and a list of the
//       functions which led to this result will be included in
//       the text output.
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
	strBuilder *strings.Builder,
	runeSearchResultsDto *CharSearchRuneArrayResultsDto,
	displayFunctionChain bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchRuneArrayResultsDtoNanobot."+
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

	if runeSearchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'runeSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	strBuilder.Grow(256)

	// Total available Length of Output Line
	const maxLineLen = 78

	// Max Label Field Length = 33
	const maxLabelFieldLen = 33

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
		"CharSearchRuneArrayResultsDto",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return err
	}

	txtStrParam :=
		runeSearchResultsDto.SearchResultsName

	if len(txtStrParam) > 0 {

		// Title Line 2
		err = txtFormatCol.AddLine1Col(
			txtStrParam,
			ePrefix.XCpy(
				"Top-Title Line 2"))

		if err != nil {
			return err
		}

	}

	// Title Line 3
	err = txtFormatCol.AddLine1Col(
		"Parameter Listing",
		ePrefix.XCpy(
			"Top-Title Line 3"))

	if err != nil {
		return err
	}

	// Title Line  4 Date/Time

	err = txtFormatCol.AddLine1Col(
		TextInputParamFieldDateTimeDto{
			FieldDateTime:       time.Now(),
			FieldDateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		},
		ePrefix.XCpy(
			"Top-Title Line 4"))

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

	// Set up 2-Column Parameters
	err = txtFormatCol.SetStdFormatParamsLine2Col(
		" ",
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		-1,
		TxtJustify.Left(),
		"",
		false,
		"",
		maxLineLen,
		true,
		ePrefix.XCpy(
			"Set 2-Column Params"))

	if err != nil {
		return err
	}

	// Build SearchResultsName Name

	txtStrLabel := "SearchResultsName"

	txtStrParam =
		runeSearchResultsDto.SearchResultsName

	if len(txtStrParam) == 0 {
		txtStrParam = "SearchResultsName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"SearchResultsName"))

	if err != nil {
		return err
	}

	// Build SearchResultsFunctionChain

	if displayFunctionChain {

		txtStrLabel = "SearchResultsFunctionChain"

		txtStrParam = runeSearchResultsDto.SearchResultsFunctionChain

		if len(txtStrParam) == 0 {

			txtStrParam = "SearchResultsFunctionChain is EMPTY!"

			err = txtFormatCol.AddLine2Col(
				txtStrLabel,
				txtStrParam,
				ePrefix.XCpy(
					""))

			if err != nil {
				return err
			}

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

			txtFormatCol.AddLineBlank(
				1,
				"")

		}

	}

	// Build IsNOP Name

	txtStrLabel = "IsNOP - Is No Operation"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.IsNOP,
		ePrefix.XCpy(
			"SearchResultsName"))

	if err != nil {
		return err
	}

	if runeSearchResultsDto.IsNOP {

		spacer := strings.Repeat(" ", maxLabelFieldLen)

		txtStrLabel =
			"This entity is a NOP or No Operation. " +
				"It is configured as an empty placeholder " +
				"and played no role in the most recent " +
				"search operation."

		txtFormatCol.AddAdHocText(
			spacer,
			txtStrLabel,
			"",
			false,
			"\n",
			maxLineLen,
			true)

		goto exitMethodTrailer

	}

	// Build FoundSearchTarget

	txtStrLabel = "FoundSearchTarget"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.FoundSearchTarget,
		ePrefix.XCpy(
			"FoundSearchTarget"))

	if err != nil {
		return err
	}

	// Build FoundFirstNumericDigitInNumStr

	txtStrLabel = "FoundFirstNumericDigitInNumStr"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.FoundFirstNumericDigitInNumStr,
		ePrefix.XCpy(
			"FoundFirstNumericDigitInNumStr"))

	if err != nil {
		return err
	}

	// Build FoundDecimalSeparatorSymbols

	txtStrLabel = "FoundDecimalSeparatorSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.FoundDecimalSeparatorSymbols,
		ePrefix.XCpy(
			"FoundDecimalSeparatorSymbols"))

	if err != nil {
		return err
	}

	// Build FoundNonZeroValue

	txtStrLabel = "FoundNonZeroValue"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.FoundNonZeroValue,
		ePrefix.XCpy(
			"FoundNonZeroValue"))

	if err != nil {
		return err
	}

	// Build TargetInputParametersName

	txtStrLabel = "TargetInputParametersName"

	txtStrParam =
		runeSearchResultsDto.TargetInputParametersName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetInputParametersName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"FoundNonZeroValue"))

	if err != nil {
		return err
	}

	// Build TargetStringLength

	txtStrLabel = "TargetStringLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TargetStringLength,
		ePrefix.XCpy(
			"TargetStringLength"))

	if err != nil {
		return err
	}

	// Build TargetStringSearchLength

	txtStrLabel = "TargetStringSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TargetStringSearchLength,
		ePrefix.XCpy(
			"TargetStringSearchLength"))

	if err != nil {
		return err
	}

	// Build TargetStringAdjustedSearchLength

	txtStrLabel = "TargetStringAdjustedSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TargetStringAdjustedSearchLength,
		ePrefix.XCpy(
			"TargetStringAdjustedSearchLength"))

	if err != nil {
		return err
	}

	// Build TargetStringStartingSearchIndex

	txtStrLabel = "TargetStringStartingSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TargetStringStartingSearchIndex,
		ePrefix.XCpy(
			"TargetStringStartingSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringCurrentSearchIndex

	txtStrLabel = "TargetStringCurrentSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TargetStringCurrentSearchIndex,
		ePrefix.XCpy(
			"TargetStringCurrentSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringFirstFoundIndex

	txtStrLabel = "TargetStringFirstFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TargetStringFirstFoundIndex,
		ePrefix.XCpy(
			"TargetStringFirstFoundIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringLastSearchIndex

	txtStrLabel = "TargetStringLastSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TargetStringLastSearchIndex,
		ePrefix.XCpy(
			"TargetStringLastSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringNextSearchIndex

	txtStrLabel = "TargetStringNextSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TargetStringNextSearchIndex,
		ePrefix.XCpy(
			"TargetStringNextSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringDescription1

	txtStrLabel = "TargetStringDescription1"

	txtStrParam =
		runeSearchResultsDto.TargetStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringDescription1 is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringDescription1"))

	if err != nil {
		return err
	}

	// Build TargetStringDescription2

	txtStrLabel = "TargetStringDescription2"

	txtStrParam =
		runeSearchResultsDto.TargetStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringDescription2 is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringDescription2"))

	if err != nil {
		return err
	}

	// Build TestInputParametersName

	txtStrLabel = "TestInputParametersName"

	txtStrParam =
		runeSearchResultsDto.TestInputParametersName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestInputParametersName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestInputParametersName"))

	if err != nil {
		return err
	}

	// Build TestStringName

	txtStrLabel = "TestStringName"

	txtStrParam =
		runeSearchResultsDto.TestStringName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringName"))

	if err != nil {
		return err
	}

	// Build TestStringLength

	txtStrLabel = "TestStringLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TestStringLength,
		ePrefix.XCpy(
			"TestStringLength"))

	if err != nil {
		return err
	}

	// Build TestStringLengthName

	txtStrLabel = "TestStringLengthName"

	txtStrParam =
		runeSearchResultsDto.TestStringLengthName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringLengthName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringLengthName"))

	if err != nil {
		return err
	}

	// Build TestStringStartingIndex

	txtStrLabel = "TestStringStartingIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TestStringStartingIndex,
		ePrefix.XCpy(
			"TestStringStartingIndex"))

	if err != nil {
		return err
	}

	// Build TestStringStartingIndexName

	txtStrLabel = "TestStringStartingIndexName"

	txtStrParam =
		runeSearchResultsDto.TestStringStartingIndexName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringStartingIndexName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringStartingIndexName"))

	if err != nil {
		return err
	}

	// Build TestStringFirstFoundIndex

	txtStrLabel = "TestStringFirstFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TestStringFirstFoundIndex,
		ePrefix.XCpy(
			"TestStringFirstFoundIndex"))

	if err != nil {
		return err
	}

	// Build TestStringLastFoundIndex

	txtStrLabel = "TestStringLastFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TestStringLastFoundIndex,
		ePrefix.XCpy(
			"TestStringLastFoundIndex"))

	if err != nil {
		return err
	}

	// Build TestStringDescription1

	txtStrLabel = "TestStringDescription1"

	txtStrParam =
		runeSearchResultsDto.TestStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringDescription1 is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringDescription1"))

	if err != nil {
		return err
	}

	// Build TestStringDescription2

	txtStrLabel = "TestStringDescription2"

	txtStrParam =
		runeSearchResultsDto.TestStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringDescription2 is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringDescription2"))

	if err != nil {
		return err
	}

	// Build CollectionTestObjIndex

	txtStrLabel = "CollectionTestObjIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.CollectionTestObjIndex,
		ePrefix.XCpy(
			"CollectionTestObjIndex"))

	if err != nil {
		return err
	}

	// Build TextCharSearchType

	txtStrLabel = "TextCharSearchType"

	if !runeSearchResultsDto.TextCharSearchType.XIsValid() {
		runeSearchResultsDto.TextCharSearchType =
			CharSearchType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		runeSearchResultsDto.TextCharSearchType,
		ePrefix.XCpy(
			"TextCharSearchType"))

	if err != nil {
		return err
	}

	// Build ReplacementString

	txtStrLabel = "ReplacementString"

	txtStrParam =
		runeSearchResultsDto.ReplacementString.GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "ReplacementString is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"ReplacementString"))

	if err != nil {
		return err
	}

	// Build RemainderString

	txtStrLabel = "RemainderString"

	txtStrParam =
		runeSearchResultsDto.RemainderString.GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "RemainderString is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"RemainderString"))

	if err != nil {
		return err
	}

	// Build FoundRuneArrayChars

	txtStrLabel = "FoundRuneArrayChars"

	txtStrParam =
		runeSearchResultsDto.FoundRuneArrayChars.GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "FoundRuneArrayChars is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"FoundRuneArrayChars"))

	if err != nil {
		return err
	}

exitMethodTrailer:

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
		"CharSearchRuneArrayResultsDto",
		ePrefix.XCpy(
			"Bottom-Title Line 1"))

	if err != nil {
		return err
	}

	txtStrParam =
		runeSearchResultsDto.SearchResultsName

	if len(txtStrParam) > 0 {

		// Title Line 2
		err = txtFormatCol.AddLine1Col(
			txtStrParam,
			ePrefix.XCpy(
				"Bottom-Title Line 2"))

		if err != nil {
			return err
		}

	}

	// Title # 3
	err = txtFormatCol.AddLine1Col(
		"End of Parameter Listing",
		ePrefix.XCpy(
			"Bottom-Title Line 2"))

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

	if err != nil {
		return err
	}

	return err
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

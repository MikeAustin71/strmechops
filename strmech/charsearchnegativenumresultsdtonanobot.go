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

	err = destinationNegNumResults.FoundLeadingNegNumSymbol.CopyIn(
		&sourceNegNumResults.FoundLeadingNegNumSymbol,
		ePrefix.XCpy(
			"destinationNegNumResults."+
				"FoundLeadingNegNumSymbol<-"+
				"sourceNegNumResults"))

	if err != nil {
		return err
	}

	err = destinationNegNumResults.FoundTrailingNegNumSymbol.CopyIn(
		&sourceNegNumResults.FoundTrailingNegNumSymbol,
		ePrefix.XCpy(
			"destinationNegNumResults."+
				"FoundTrailingNegNumSymbol<-"+
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

	err = deepCopyNegNumResultsDto.FoundLeadingNegNumSymbol.CopyIn(
		&searchNegNumResults.FoundLeadingNegNumSymbol,
		ePrefix.XCpy(
			"deepCopyNegNumResultsDto."+
				"FoundLeadingNegNumSymbol<-"+
				"searchNegNumResults"))

	if err != nil {

		return deepCopyNegNumResultsDto, err
	}

	err = deepCopyNegNumResultsDto.FoundTrailingNegNumSymbol.CopyIn(
		&searchNegNumResults.FoundTrailingNegNumSymbol,
		ePrefix.XCpy(
			"deepCopyNegNumResultsDto."+
				"FoundTrailingNegNumSymbol<-"+
				"searchNegNumResults"))

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
	strBuilder *strings.Builder,
	searchNegNumResults *CharSearchNegativeNumberResultsDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if searchNegNumResultsNanobot.lock == nil {
		searchNegNumResultsNanobot.lock = new(sync.Mutex)
	}

	searchNegNumResultsNanobot.lock.Lock()

	defer searchNegNumResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNegNumResultsDtoNanobot."+
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

	strBuilder.Grow(512)

	if searchNegNumResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchNegNumResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	// Total available Length of Output Line
	const maxLineLen = 79

	// Max Label Field Length = 24
	const maxLabelFieldLen = 35

	txtFormatCol := TextFormatterCollection{}

	// Leading Title Marquee

	// Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		maxLineLen-2,
		" ",
		false,
		"",
		-1,
		false)

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

	// Title Line 1
	err = txtFormatCol.AddLine1Col(
		"CharSearchNegativeNumberResultsDto",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return err
	}

	txtStrParam :=
		searchNegNumResults.SearchResultsName

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
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		maxLineLen-2,
		" ",
		false,
		"",
		-1,
		false)

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
		searchNegNumResults.SearchResultsName

	if len(txtStrParam) == 0 {
		txtStrParam = "SearchResultsName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			""))

	if err != nil {
		return err
	}

	// Build SearchResultsFunctionChain

	txtStrLabel = "SearchResultsFunctionChain"

	txtStrParam = searchNegNumResults.SearchResultsFunctionChain

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

	}

	// Build FoundNegativeNumberSymbols

	txtStrLabel = "FoundNegativeNumberSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.FoundNegativeNumberSymbols,
		ePrefix.XCpy(
			"FoundNegativeNumberSymbols"))

	if err != nil {
		return err
	}

	// Build FoundNegNumSymbolsOnPreviousSearch

	txtStrLabel = "FoundNegNumSymbolsOnPreviousSearch"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.FoundNegNumSymbolsOnPreviousSearch,
		ePrefix.XCpy(
			"FoundNegNumSymbolsOnPreviousSearch"))

	if err != nil {
		return err
	}

	// Build FoundLeadingNegNumSymbols

	txtStrLabel = "FoundLeadingNegNumSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.FoundLeadingNegNumSymbols,
		ePrefix.XCpy(
			"FoundLeadingNegNumSymbols"))

	if err != nil {
		return err
	}

	// Build FoundTrailingNegNumSymbols

	txtStrLabel = "FoundTrailingNegNumSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.FoundTrailingNegNumSymbols,
		ePrefix.XCpy(
			"FoundTrailingNegNumSymbols"))

	if err != nil {
		return err
	}

	// Build FoundFirstNumericDigitInNumStr

	txtStrLabel = "FoundFirstNumericDigitInNumStr"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.FoundFirstNumericDigitInNumStr,
		ePrefix.XCpy(
			"FoundFirstNumericDigitInNumStr"))

	if err != nil {
		return err
	}

	// Build FoundDecimalSeparatorSymbols

	txtStrLabel = "FoundDecimalSeparatorSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.FoundDecimalSeparatorSymbols,
		ePrefix.XCpy(
			"FoundDecimalSeparatorSymbols"))

	if err != nil {
		return err
	}

	// Build FoundNonZeroValue

	txtStrLabel = "FoundNonZeroValue"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.FoundNonZeroValue,
		ePrefix.XCpy(
			"FoundNonZeroValue"))

	if err != nil {
		return err
	}

	// Build TargetInputParametersName

	txtStrLabel = "TargetInputParametersName"

	txtStrParam =
		searchNegNumResults.TargetInputParametersName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetInputParametersName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetInputParametersName"))

	if err != nil {
		return err
	}

	// Build TargetStringLength

	txtStrLabel = "TargetStringLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringLength,
		ePrefix.XCpy(
			"TargetStringLength"))

	if err != nil {
		return err
	}

	// Build TargetStringSearchLength

	txtStrLabel = "TargetStringSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringSearchLength,
		ePrefix.XCpy(
			"TargetStringSearchLength"))

	if err != nil {
		return err
	}

	// Build TargetStringAdjustedSearchLength

	txtStrLabel = "TargetStringAdjustedSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringAdjustedSearchLength,
		ePrefix.XCpy(
			"TargetStringAdjustedSearchLength"))

	if err != nil {
		return err
	}

	// Build TargetStringStartingSearchIndex

	txtStrLabel = "TargetStringStartingSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringStartingSearchIndex,
		ePrefix.XCpy(
			"TargetStringStartingSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringCurrentSearchIndex

	txtStrLabel = "TargetStringCurrentSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringCurrentSearchIndex,
		ePrefix.XCpy(
			"TargetStringCurrentSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringFirstFoundIndex

	txtStrLabel = "TargetStringFirstFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringFirstFoundIndex,
		ePrefix.XCpy(
			"TargetStringFirstFoundIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringLastFoundIndex

	txtStrLabel = "TargetStringLastFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringLastFoundIndex,
		ePrefix.XCpy(
			"TargetStringLastFoundIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringLastSearchIndex

	txtStrLabel = "TargetStringLastSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringLastSearchIndex,
		ePrefix.XCpy(
			"TargetStringLastSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringNextSearchIndex

	txtStrLabel = "TargetStringNextSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TargetStringNextSearchIndex,
		ePrefix.XCpy(
			"TargetStringNextSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringDescription1

	txtStrLabel = "TargetStringDescription1"

	txtStrParam =
		searchNegNumResults.TargetStringDescription1

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
		searchNegNumResults.TargetStringDescription2

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
		searchNegNumResults.TestInputParametersName

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
		searchNegNumResults.TestStringName

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
		searchNegNumResults.TestStringLength,
		ePrefix.XCpy(
			"TestStringLength"))

	if err != nil {
		return err
	}

	// Build TestStringLengthName

	txtStrLabel = "TestStringLengthName"

	txtStrParam =
		searchNegNumResults.TestStringLengthName

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
		searchNegNumResults.TestStringStartingIndex,
		ePrefix.XCpy(
			"TestStringStartingIndex"))

	if err != nil {
		return err
	}

	// Build TestStringStartingIndexName

	txtStrLabel = "TestStringStartingIndexName"

	txtStrParam =
		searchNegNumResults.TestStringStartingIndexName

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
		searchNegNumResults.TestStringFirstFoundIndex,
		ePrefix.XCpy(
			"TestStringFirstFoundIndex"))

	if err != nil {
		return err
	}

	// Build TestStringLastFoundIndex

	txtStrLabel = "TestStringLastFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TestStringLastFoundIndex,
		ePrefix.XCpy(
			"TestStringLastFoundIndex"))

	if err != nil {
		return err
	}

	// Build TestStringDescription1

	txtStrLabel = "TestStringDescription1"

	txtStrParam =
		searchNegNumResults.TestStringDescription1

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
		searchNegNumResults.TestStringDescription2

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
		searchNegNumResults.CollectionTestObjIndex,
		ePrefix.XCpy(
			"CollectionTestObjIndex"))

	if err != nil {
		return err
	}

	// Build NumSignValue

	txtStrLabel = "NumSignValue"

	if !searchNegNumResults.NumSignValue.XIsValid() {
		searchNegNumResults.NumSignValue =
			NumSignVal.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.NumSignValue,
		ePrefix.XCpy(
			"NumSignValue"))

	if err != nil {
		return err
	}

	// Build PrimaryNumSignPosition

	txtStrLabel = "PrimaryNumSignPosition"

	if !searchNegNumResults.PrimaryNumSignPosition.XIsValid() {
		searchNegNumResults.PrimaryNumSignPosition =
			NumSignSymPos.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.PrimaryNumSignPosition,
		ePrefix.XCpy(
			"PrimaryNumSignPosition"))

	if err != nil {
		return err
	}

	// Build SecondaryNumSignPosition

	txtStrLabel = "SecondaryNumSignPosition"

	if !searchNegNumResults.SecondaryNumSignPosition.XIsValid() {
		searchNegNumResults.SecondaryNumSignPosition =
			NumSignSymPos.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.SecondaryNumSignPosition,
		ePrefix.XCpy(
			"SecondaryNumSignPosition"))

	if err != nil {
		return err
	}

	// Build TextCharSearchType

	txtStrLabel = "TextCharSearchType"

	if !searchNegNumResults.TextCharSearchType.XIsValid() {
		searchNegNumResults.TextCharSearchType =
			CharSearchType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchNegNumResults.TextCharSearchType,
		ePrefix.XCpy(
			"TextCharSearchType"))

	if err != nil {
		return err
	}

	// Build NegativeNumberSymbolsSpec

	txtStrLabel = "NegativeNumberSymbolsSpec"

	if !searchNegNumResults.NegativeNumberSymbolsSpec.
		IsValidInstance() {
		txtStrParam = "NegativeNumberSymbolsSpec is populated and valid."
	} else {
		txtStrParam = "NegativeNumberSymbolsSpec is invalid and empty."
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"NegativeNumberSymbolsSpec"))

	if err != nil {
		return err
	}

	// Build ReplacementString

	txtStrLabel = "ReplacementString"

	txtStrParam =
		searchNegNumResults.ReplacementString.GetCharacterString()

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
		searchNegNumResults.RemainderString.GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "RemainderString is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"ReplacementString"))

	if err != nil {
		return err
	}

	// Build FoundLeadingNegNumSymbol Characters

	txtStrParam =
		searchNegNumResults.FoundLeadingNegNumSymbol.GetCharacterString()

	lenTxtStrParam := len(txtStrParam)

	txtStrLabel = "FoundLeadingNegNumSymbol Chars"

	if lenTxtStrParam == 0 {
		txtStrParam = "FoundLeadingNegNumSymbol is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"FoundLeadingNegNumSymbol"))

	if err != nil {
		return err
	}

	if lenTxtStrParam > 0 {

		txtStrLabel = "FoundLeadingNegNumSymbol Chars Length"

		err = txtFormatCol.AddLine2Col(
			txtStrLabel,
			txtStrParam,
			ePrefix.XCpy(
				"FoundLeadingNegNumSymbol Chars Length"))

		if err != nil {
			return err
		}

	}

	// Build FoundTrailingNegNumSymbol Characters

	txtStrParam =
		searchNegNumResults.FoundTrailingNegNumSymbol.
			GetCharacterString()

	lenTxtStrParam = len(txtStrParam)

	txtStrLabel = "FoundTrailingNegNumSymbol Chars"

	if lenTxtStrParam == 0 {
		txtStrParam = "FoundTrailingNegNumSymbol is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"FoundTrailingNegNumSymbol"))

	if err != nil {
		return err
	}

	if lenTxtStrParam > 0 {

		txtStrLabel = "FoundTrailingNegNumSymbol Chars Length"

		err = txtFormatCol.AddLine2Col(
			txtStrLabel,
			txtStrParam,
			ePrefix.XCpy(
				"FoundTrailingNegNumSymbol Chars Length"))

		if err != nil {
			return err
		}

	}

	// Trailing Title Marquee
	// Top Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		maxLineLen-2,
		" ",
		false,
		"",
		-1,
		false)

	// Title # 1
	err = txtFormatCol.AddLine1Col(
		"CharSearchNegativeNumberResultsDto",
		ePrefix.XCpy(
			"Bottom-Title Line 1"))

	if err != nil {
		return err
	}

	// Title # 2
	err = txtFormatCol.AddLine1Col(
		"End of Parameter Listing",
		ePrefix.XCpy(
			"Bottom-Title Line 2"))

	if err != nil {
		return err
	}

	// Filler =======
	// Marquee Bottom
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		maxLineLen-2,
		" ",
		false,
		"",
		-1,
		false)

	// Blank Line
	txtFormatCol.AddLineBlank(
		2,
		"")

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Marquee-Bottom"))

	return err
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

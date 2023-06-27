package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// charSearchDecimalSeparatorResultsDtoNanobot - Provides helper
// methods for type CharSearchDecimalSeparatorResultsDto.
type charSearchDecimalSeparatorResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceDecSepResults' to input parameter
// 'destinationDecSepResults'. Both instances are of type
// CharSearchDecimalSeparatorResultsDto.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationDecSepResults'
// will be overwritten.
//
// Also, NO data validation is performed on 'sourceDecSepResults'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationDecSepRes   *CharSearchDecimalSeparatorResultsDto
//	   - A pointer to a CharSearchDecimalSeparatorResultsDto instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceDecSepResults'.
//
//	     'destinationDecSepResults' is the destination for this
//	     copy operation.
//
//
//	sourceDecSepResults    *CharSearchDecimalSeparatorResultsDto
//	   - A pointer to another CharSearchDecimalSeparatorResultsDto
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationDecSepResults'.
//
//	     'sourceDecSepResults' is the source for this copy
//	     operation.
//
//	     No data validation is performed on 'sourceDecSepResults'.
//
//
//	errPrefDto             *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
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

	new(charSearchDecimalSeparatorResultsDtoAtom).empty(
		destinationDecSepResults)

	destinationDecSepResults.SearchResultsName =
		sourceDecSepResults.SearchResultsName

	destinationDecSepResults.SearchResultsFunctionChain =
		sourceDecSepResults.SearchResultsFunctionChain

	destinationDecSepResults.IsNOP =
		sourceDecSepResults.IsNOP

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
// # IMPORTANT
//
// NO validation is performed on 'decimalSeparatorResults'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	decimalSeparatorResults  *CharSearchDecimalSeparatorResultsDto
//	   - A pointer to an instance of
//	     CharSearchDecimalSeparatorResultsDto. A deep copy of the
//	     internal member variables contained in this instance will
//	     be created and returned in a new instance of
//	     CharSearchDecimalSeparatorResultsDto.
//
//	     No data validation is performed on
//	     'decimalSeparatorResults'.
//
//
//	errPrefDto               *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	deepCopyNegNumResultsDto CharSearchDecimalSeparatorResultsDto
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'decimalSeparatorResults' will be created
//	     and returned in a new instance of
//	     CharSearchDecimalSeparatorResultsDto.
//
//
//	err                      error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
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

	new(charSearchDecimalSeparatorResultsDtoAtom).empty(
		&deepCopyDecSepResults)

	deepCopyDecSepResults.SearchResultsName =
		decimalSeparatorResults.SearchResultsName

	deepCopyDecSepResults.SearchResultsFunctionChain =
		decimalSeparatorResults.SearchResultsFunctionChain

	deepCopyDecSepResults.IsNOP =
		decimalSeparatorResults.IsNOP

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
// ----------------------------------------------------------------
//
// Input Parameters
//
//	strBuilder                 *strings.Builder
//	   - A pointer to an instance of *strings.Builder. The
//	     formatted text characters produced by this method will be
//	     written to this instance of strings.Builder.
//
//
//	decimalSeparatorResults *CharSearchDecimalSeparatorResultsDto
//	   - A pointer to an instance of
//	     CharSearchDecimalSeparatorResultsDto instance. Formatted
//	     text output will be generated listing the member variable
//	     names and their corresponding values. The formatted text
//	     can then be used for text displays, file output or
//	     printing.
//
//	     No data validation is performed on this instance of
//	     CharSearchDecimalSeparatorResultsDto.
//
//
//	displayFunctionChain       bool
//	   - Set 'displayFunctionChain' to 'true' and a list of the
//	     functions which led to this result will be included in
//	     the text output.
//
//
//	errPrefDto              *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this
//	     parameter to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
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
func (searchDecimalSepResultsNanobot *charSearchDecimalSeparatorResultsDtoNanobot) getParameterTextListing(
	strBuilder *strings.Builder,
	decimalSeparatorResults *CharSearchDecimalSeparatorResultsDto,
	displayFunctionChain bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if searchDecimalSepResultsNanobot.lock == nil {
		searchDecimalSepResultsNanobot.lock = new(sync.Mutex)
	}

	searchDecimalSepResultsNanobot.lock.Lock()

	defer searchDecimalSepResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchDecimalSeparatorResultsDtoNanobot."+
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

	if decimalSeparatorResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'decimalSeparatorResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	// Total available Length of Output Line
	const maxLineLen = 79

	// Max Label Field Length = 38
	const maxLabelFieldLen = 38

	txtFormatCol := TextFormatterCollection{}

	// Leading Title Marquee

	txtFmtParams := TextFmtParamsLineColumnsDto{
		FormatType: TxtFieldType.LineColumns(),
		FieldFormatParams: []TextFieldFmtParamsDto{
			{
				LeftMarginStr:  " ",
				FieldLength:    maxLineLen,
				FieldJustify:   TxtJustify.Center(),
				DateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
				RightMarginStr: "",
			},
		},
		TurnLineTerminationOff:     false,
		LineTerminator:             "\n",
		MaxLineLength:              maxLineLen,
		TurnAutoLineLengthBreaksOn: false,
	}

	titles := StringArrayDto{}

	titles.AddString(
		"CharSearchDecimalSeparatorResultsDto")

	txtStrParam :=
		decimalSeparatorResults.SearchResultsName

	if len(txtStrParam) > 0 {

		titles.AddString(
			txtStrParam)
	}

	titles.AddString(
		"Parameter Listing")

	err = new(TextUtility).BuildOneColLeadingMarquee(
		strBuilder,
		true,
		"=",
		"=",
		true,
		titles,
		txtFmtParams,
		true,
		ePrefix.XCpy(
			"strBuilder<-txtFmtParams"))

	if err != nil {

		return err

	}

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
		" ",
		ePrefix.XCpy(
			"Set 2-Column Params"))

	if err != nil {
		return err
	}

	// Build SearchResultsName Name

	txtStrLabel := "SearchResultsName"

	txtStrParam =
		decimalSeparatorResults.SearchResultsName

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

	if displayFunctionChain {

		txtStrLabel = "SearchResultsFunctionChain"

		txtStrParam = decimalSeparatorResults.SearchResultsFunctionChain

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

			txtFormatCol.AddLineBlank(
				1,
				"")

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

	var lenTxtStrParam int

	// Build IsNOP Name

	txtStrLabel = "IsNOP - Is No Operation"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.IsNOP,
		ePrefix.XCpy(
			"SearchResultsName"))

	if err != nil {
		return err
	}

	if decimalSeparatorResults.IsNOP {

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

	// Build FoundDecimalSeparatorSymbols

	txtStrLabel = "FoundDecimalSeparatorSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.FoundDecimalSeparatorSymbols,
		ePrefix.XCpy(
			"FoundDecimalSeparatorSymbols"))

	if err != nil {
		return err
	}

	// Build FoundDecimalSepSymbolsOnPreviousSearch

	txtStrLabel = "FoundDecimalSepSymbolsOnPreviousSearch"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.FoundDecimalSepSymbolsOnPreviousSearch,
		ePrefix.XCpy(
			"FoundDecimalSepSymbolsOnPreviousSearch"))

	if err != nil {
		return err
	}

	// Build FoundFirstNumericDigitInNumStr

	txtStrLabel = "FoundFirstNumericDigitInNumStr"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.FoundFirstNumericDigitInNumStr,
		ePrefix.XCpy(
			"FoundFirstNumericDigitInNumStr"))

	if err != nil {
		return err
	}

	// Build FoundNonZeroValue

	txtStrLabel = "FoundNonZeroValue"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.FoundNonZeroValue,
		ePrefix.XCpy(
			"FoundNonZeroValue"))

	if err != nil {
		return err
	}

	// Build TargetInputParametersName

	txtStrLabel = "TargetInputParametersName"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetInputParametersName,
		ePrefix.XCpy(
			"TargetInputParametersName"))

	if err != nil {
		return err
	}

	// Build TargetStringLength

	txtStrLabel = "TargetStringLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringLength,
		ePrefix.XCpy(
			"TargetStringLength"))

	if err != nil {
		return err
	}

	// Build TargetStringSearchLength

	txtStrLabel = "TargetStringSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringSearchLength,
		ePrefix.XCpy(
			"TargetStringSearchLength"))

	if err != nil {
		return err
	}

	// Build TargetStringAdjustedSearchLength

	txtStrLabel = "TargetStringAdjustedSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringAdjustedSearchLength,
		ePrefix.XCpy(
			"TargetStringAdjustedSearchLength"))

	if err != nil {
		return err
	}

	// Build TargetStringStartingSearchIndex

	txtStrLabel = "TargetStringStartingSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringStartingSearchIndex,
		ePrefix.XCpy(
			"TargetStringStartingSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringCurrentSearchIndex

	txtStrLabel = "TargetStringCurrentSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringCurrentSearchIndex,
		ePrefix.XCpy(
			"TargetStringCurrentSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringFirstFoundIndex

	txtStrLabel = "TargetStringFirstFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringFirstFoundIndex,
		ePrefix.XCpy(
			"TargetStringFirstFoundIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringLastFoundIndex

	txtStrLabel = "TargetStringLastFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringLastFoundIndex,
		ePrefix.XCpy(
			"TargetStringLastFoundIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringLastSearchIndex

	txtStrLabel = "TargetStringLastSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringLastSearchIndex,
		ePrefix.XCpy(
			"TargetStringLastSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringNextSearchIndex

	txtStrLabel = "TargetStringNextSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TargetStringNextSearchIndex,
		ePrefix.XCpy(
			"TargetStringNextSearchIndex"))

	if err != nil {
		return err
	}

	// Build TargetStringDescription1

	txtStrLabel = "TargetStringDescription1"

	txtStrParam = decimalSeparatorResults.TargetStringDescription1

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

	txtStrParam = decimalSeparatorResults.TargetStringDescription2

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

	txtStrParam = decimalSeparatorResults.TestInputParametersName

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

	txtStrParam = decimalSeparatorResults.TestStringName

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
		decimalSeparatorResults.TestStringLength,
		ePrefix.XCpy(
			"TestStringLength"))

	if err != nil {
		return err
	}

	// Build TestStringLengthName

	txtStrLabel = "TestStringLengthName"

	txtStrParam = decimalSeparatorResults.TestStringLengthName

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
		decimalSeparatorResults.TestStringStartingIndex,
		ePrefix.XCpy(
			"TestStringStartingIndex"))

	if err != nil {
		return err
	}

	// Build TestStringStartingIndexName

	txtStrLabel = "TestStringStartingIndexName"

	txtStrParam = decimalSeparatorResults.TestStringStartingIndexName

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
		decimalSeparatorResults.TestStringFirstFoundIndex,
		ePrefix.XCpy(
			"TestStringFirstFoundIndex"))

	if err != nil {
		return err
	}

	// Build TestStringLastFoundIndex

	txtStrLabel = "TestStringLastFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TestStringLastFoundIndex,
		ePrefix.XCpy(
			"TestStringLastFoundIndex"))

	if err != nil {
		return err
	}

	// Build TestStringDescription1

	txtStrLabel = "TestStringDescription1"

	txtStrParam = decimalSeparatorResults.TestStringDescription1

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

	txtStrParam = decimalSeparatorResults.TestStringDescription2

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
		decimalSeparatorResults.CollectionTestObjIndex,
		ePrefix.XCpy(
			"CollectionTestObjIndex"))

	if err != nil {
		return err
	}

	// Build NumValueType

	txtStrLabel = "NumValueType"

	if !decimalSeparatorResults.NumValueType.XIsValid() {
		decimalSeparatorResults.NumValueType =
			NumValType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.NumValueType,
		ePrefix.XCpy(
			"NumValueType"))

	if err != nil {
		return err
	}

	// Build NumSymbolLocation

	txtStrLabel = "NumSymbolLocation"

	if !decimalSeparatorResults.NumSymbolLocation.XIsValid() {
		decimalSeparatorResults.NumSymbolLocation =
			NumSymLocation.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.NumSymbolLocation,
		ePrefix.XCpy(
			"NumSymbolLocation"))

	if err != nil {
		return err
	}

	// Build TextCharSearchType

	txtStrLabel = "TextCharSearchType"

	if !decimalSeparatorResults.TextCharSearchType.XIsValid() {
		decimalSeparatorResults.TextCharSearchType =
			CharSearchType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		decimalSeparatorResults.TextCharSearchType,
		ePrefix.XCpy(
			"TextCharSearchType"))

	if err != nil {
		return err
	}

	// Build DecimalSeparatorSymbolsSpec

	txtStrParam = decimalSeparatorResults.
		DecimalSeparatorSymbolsSpec.
		GetDecimalSeparatorStr()

	lenTxtStrParam = len(txtStrParam)

	txtStrLabel = "DecimalSeparatorSymbolsSpec"

	if lenTxtStrParam == 0 {

		txtStrParam = "DecimalSeparatorSymbolsSpec is EMPTY!"

	} else {

		txtStrParam = "\"" + txtStrParam + "\""

	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"DecimalSeparatorSymbolsSpec"))

	if err != nil {
		return err
	}

	if lenTxtStrParam > 0 {
		txtStrLabel = "DecimalSeparatorSymbols Length"

		err = txtFormatCol.AddLine2Col(
			txtStrLabel,
			lenTxtStrParam,
			ePrefix.XCpy(
				"DecimalSeparatorSymbols Length"))

		if err != nil {
			return err
		}

	}

	// Build ReplacementString

	txtStrLabel = "ReplacementString"

	txtStrParam = decimalSeparatorResults.
		ReplacementString.
		GetCharacterString()

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

	txtStrParam = decimalSeparatorResults.
		RemainderString.
		GetCharacterString()

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

	txtStrParam = decimalSeparatorResults.
		FoundRuneArrayChars.
		GetCharacterString()

	lenTxtStrParam = len(txtStrParam)

	txtStrLabel = "FoundRuneArrayChars"

	txtStrParam = decimalSeparatorResults.
		FoundRuneArrayChars.
		GetCharacterString()

	if lenTxtStrParam == 0 {
		txtStrParam = "FoundRuneArrayChars is EMPTY!"
	} else {
		txtStrParam = "\"" + txtStrParam + "\""
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"FoundRuneArrayChars"))

	if err != nil {
		return err
	}

	if lenTxtStrParam > 0 {

		txtStrLabel = "FoundRuneArrayChars Length"

		err = txtFormatCol.AddLine2Col(
			txtStrLabel,
			lenTxtStrParam,
			ePrefix.XCpy(
				"FoundRuneArrayChars Length"))

		if err != nil {
			return err
		}

	}

exitMethodTrailer:

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Output"))

	// Trailing Title Marquee
	// Top Blank Line

	titles.Empty()

	titles.AddString(
		"CharSearchNumStrParseResultsDto")

	txtStrParam =
		decimalSeparatorResults.SearchResultsName

	if len(txtStrParam) > 0 {

		titles.AddString(
			txtStrParam)
	}

	titles.AddString(
		"End Of Parameter Listing")

	err = new(TextUtility).BuildOneColTrailingMarquee(
		strBuilder,
		true,
		"=",
		"=",
		true,
		titles,
		txtFmtParams,
		ePrefix.XCpy(
			"strBuilder<-txtFmtParams"))

	if err != nil {
		return err
	}

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Output"))

	return err
}

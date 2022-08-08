package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

// charSearchNumStrParseResultsDtoNanobot - Provides helper
// methods for type CharSearchNumStrParseResultsDto.
type charSearchNumStrParseResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyNumStrParseResultsDto - Copies all data from input parameter
// 'sourceNumStrParseResults' to input parameter
// 'destinationNumStrParseResults'. Both instances are of type
// CharSearchNumStrParseResultsDto.
//
// # IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in
// 'destinationNumStrParseResults' will be overwritten.
//
// Also, NO data validation is performed on
// 'sourceNumStrParseResults'.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	destinationNumStrParseResults  *CharSearchNumStrParseResultsDto
//	   - A pointer to a CharSearchNumStrParseResultsDto instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceNumStrParseResults'.
//
//	     'destinationNumStrParseResults' is the destination for this
//	     copy operation.
//
//
//	sourceNumStrParseResults       *CharSearchNumStrParseResultsDto
//	   - A pointer to another CharSearchNumStrParseResultsDto
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationNumStrParseResults'.
//
//	     'sourceNumStrParseResults' is the source for this copy
//	     operation.
//
//	     No data validation is performed on 'sourceNumStrParseResults'.
//
//
//	errPrefDto                     *ePref.ErrPrefixDto
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
func (searchNumStrParseResultsNanobot *charSearchNumStrParseResultsDtoNanobot) copyNumStrParseResultsDto(
	destinationNumStrParseResults *CharSearchNumStrParseResultsDto,
	sourceNumStrParseResults *CharSearchNumStrParseResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchNumStrParseResultsNanobot.lock == nil {
		searchNumStrParseResultsNanobot.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsNanobot.lock.Lock()

	defer searchNumStrParseResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNumStrParseResultsDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationNumStrParseResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationNumStrParseResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceNumStrParseResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceNumStrParseResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	new(charSearchNumStrParseResultsDtoAtom).
		empty(destinationNumStrParseResults)

	destinationNumStrParseResults.SearchResultsName =
		sourceNumStrParseResults.SearchResultsName

	destinationNumStrParseResults.SearchResultsFunctionChain =
		sourceNumStrParseResults.SearchResultsFunctionChain

	err = destinationNumStrParseResults.TargetSearchString.CopyIn(
		&sourceNumStrParseResults.TargetSearchString,
		ePrefix.XCpy(
			"destinationNumStrParseResults"+
				"<-sourceNumStrParseResults.TargetString"))

	if err != nil {
		return err
	}

	destinationNumStrParseResults.TargetStringSearchLength =
		sourceNumStrParseResults.TargetStringSearchLength

	destinationNumStrParseResults.TargetStringAdjustedSearchLength =
		sourceNumStrParseResults.TargetStringAdjustedSearchLength

	destinationNumStrParseResults.TargetStringStartingSearchIndex =
		sourceNumStrParseResults.TargetStringStartingSearchIndex

	destinationNumStrParseResults.ReasonForSearchTermination =
		sourceNumStrParseResults.ReasonForSearchTermination

	destinationNumStrParseResults.FoundNumericDigits =
		sourceNumStrParseResults.FoundNumericDigits

	destinationNumStrParseResults.FoundNonZeroValue =
		sourceNumStrParseResults.FoundNonZeroValue

	destinationNumStrParseResults.FoundDecimalSeparatorSymbols =
		sourceNumStrParseResults.FoundDecimalSeparatorSymbols

	destinationNumStrParseResults.FoundIntegerDigits =
		sourceNumStrParseResults.FoundIntegerDigits

	destinationNumStrParseResults.IdentifiedIntegerDigits =
		sourceNumStrParseResults.IdentifiedIntegerDigits

	destinationNumStrParseResults.FoundDecimalDigits =
		sourceNumStrParseResults.FoundDecimalDigits

	destinationNumStrParseResults.IdentifiedFractionalDigits =
		sourceNumStrParseResults.IdentifiedFractionalDigits

	destinationNumStrParseResults.NumSignValue =
		sourceNumStrParseResults.NumSignValue

	destinationNumStrParseResults.NumValueType =
		sourceNumStrParseResults.NumValueType

	err = destinationNumStrParseResults.RemainderString.
		CopyIn(
			&sourceNumStrParseResults.RemainderString,
			ePrefix.XCpy(
				"destinationNumStrParseResults"+
					"<-sourceNumStrParseResults"))

	if err != nil {
		return err
	}

	err = destinationNumStrParseResults.DecimalSeparatorSearchResults.
		CopyIn(
			&sourceNumStrParseResults.DecimalSeparatorSearchResults,
			ePrefix.XCpy(
				"destinationNumStrParseResults"+
					"<-sourceNumStrParseResults"))

	if err != nil {
		return err
	}

	err = destinationNumStrParseResults.NegativeNumberSymbolSearchResults.
		CopyIn(
			&sourceNumStrParseResults.NegativeNumberSymbolSearchResults,
			ePrefix.XCpy(
				"destinationNumStrParseResults"+
					"<-sourceNumStrParseResults"))

	if err != nil {
		return err
	}

	err = destinationNumStrParseResults.ParsingTerminatorSearchResults.
		CopyIn(
			&sourceNumStrParseResults.ParsingTerminatorSearchResults,
			ePrefix.XCpy(
				"destinationNumStrParseResults"+
					"<-sourceNumStrParseResults"))

	return err
}

// getParameterTextListing - Returns formatted text output
// listing the member variable names and corresponding values
// contained in the 'numStrParseResults' instance of
// CharSearchNumStrParseResultsDto.
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
//	numStrParseResults  *CharSearchNumStrParseResultsDto
//	   - A pointer to an instance of
//	     CharSearchNumStrParseResultsDto instance. Formatted
//	     text output will be generated listing the member variable
//	     names and their corresponding values. The formatted text
//	     can then be used for screen displays, file output or
//	     printing.
//
//	     No data validation is performed on this instance of
//	     CharSearchNumStrParseResultsDto.
//
//
//	displayFunctionChain       bool
//	   - Set 'displayFunctionChain' to 'true' and a list of the
//	     functions which led to this result will be included in
//	     the text output.
//
//
//	printDetail         bool
//	   - If this parameter is set to 'true', detail information for
//	     subsidiary types RemainderString,
//	     DecimalSeparatorSearchResults,
//	     NegativeNumberSymbolSearchResults and
//	     ParsingTerminatorSearchResults will be included in the
//	     formatted text output.
//
//	errPrefDto          *ePref.ErrPrefixDto
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
func (searchNumStrParseResultsNanobot *charSearchNumStrParseResultsDtoNanobot) getParameterTextListing(
	strBuilder *strings.Builder,
	numStrParseResults *CharSearchNumStrParseResultsDto,
	displayFunctionChain bool,
	printDetail bool,
	errPrefDto *ePref.ErrPrefixDto) error {

	if searchNumStrParseResultsNanobot.lock == nil {
		searchNumStrParseResultsNanobot.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsNanobot.lock.Lock()

	defer searchNumStrParseResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNumStrParseResultsDtoNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {
		return err
	}

	if numStrParseResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrParseResults' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'strBuilder' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	strBuilder.Grow(512)

	// Total available Length of Output Line
	const maxLineLen = 79

	// Max Label Field Length = 35
	const maxLabelFieldLen = 35
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
		"CharSearchNumStrParseResultsDto")

	txtStrParam :=
		numStrParseResults.SearchResultsName

	if len(txtStrParam) > 0 {

		titles.AddString(
			txtStrParam)
	}

	titles.AddString(
		"Parameter Listing")

	err = new(TextUtility).BuildOneColLeadingMarquee(
		strBuilder,
		titles,
		txtFmtParams,
		ePrefix.XCpy(
			"strBuilder<-txtFmtParams"))

	if err != nil {

		return err

	}

	// End Of Marquee

	// Begin Label Parameter Pairs

	colonSpace := ": "

	// Build SearchResultsName Name

	txtStrParam = numStrParseResults.SearchResultsName

	if len(txtStrParam) == 0 {
		txtStrParam = "SearchResultsName is EMPTY!"
	}

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
	err = txtFormatCol.AddLine2Col(
		"SearchResultsName",
		txtStrParam,
		ePrefix.XCpy(
			"Build FoundNumericDigits"))

	if err != nil {
		return err
	}

	txtFormatCol.AddLineBlank(1, "")

	// Build SearchResultsFunctionChain

	if displayFunctionChain == true {

		txtStrLabel := "SearchResultsFunctionChain"

		txtStrParam = numStrParseResults.SearchResultsFunctionChain

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

		txtFormatCol.AddLineBlank(1, "")

	}

	// Build Target Search String

	txtStrParam =
		numStrParseResults.TargetSearchString.GetCharacterString()

	lenTxtStrParam := len(txtStrParam)

	if lenTxtStrParam == 0 {
		txtStrParam = "TargetSearchString is EMPTY!"
	} else {

		txtStrParam = "\"" + txtStrParam + "\""
	}

	err = txtFormatCol.AddLine2Col(
		"Original TargetSearchString",
		txtStrParam,
		ePrefix.XCpy(
			"Build TargetSearchString"))

	if err != nil {
		return err
	}

	// Build Target Search String Length

	err = txtFormatCol.AddLine2Col(
		"Original TargetSearchString Length",
		lenTxtStrParam,
		ePrefix.XCpy(
			"Build TargetSearchString Length"))

	if err != nil {
		return err
	}

	// Build TargetStringSearchLength
	err = txtFormatCol.AddLine2Col(
		"TargetStringSearchLength",
		numStrParseResults.TargetStringSearchLength,
		ePrefix.XCpy(
			"Build TargetStringSearchLength"))

	if err != nil {
		return err
	}

	// Build TargetStringAdjustedSearchLength
	err = txtFormatCol.AddLine2Col(
		"TargetStringAdjustedSearchLength",
		numStrParseResults.TargetStringAdjustedSearchLength,
		ePrefix.XCpy(
			"Build TargetStringAdjustedSearchLength"))

	if err != nil {
		return err
	}

	// Build TargetStringStartingSearchIndex
	err = txtFormatCol.AddLine2Col(
		"TargetStringStartingSearchIndex",
		numStrParseResults.TargetStringStartingSearchIndex,
		ePrefix.XCpy(
			"Build TargetStringStartingSearchIndex"))

	if err != nil {
		return err
	}

	// Build ReasonForSearchTermination
	txtStrParam =
		numStrParseResults.
			ReasonForSearchTermination.
			XReturnNoneIfInvalid().XDescription()

	err = txtFormatCol.AddLine2Col(
		"Reason For Search Termination",
		txtStrParam,
		ePrefix.XCpy(
			"Build ReasonForSearchTermination"))

	if err != nil {
		return err
	}

	// Build FoundNumericDigits
	err = txtFormatCol.AddLine2Col(
		"FoundNumericDigits",
		numStrParseResults.FoundNumericDigits,
		ePrefix.XCpy(
			"Build FoundNumericDigits"))

	if err != nil {
		return err
	}

	// Build FoundDecimalSeparatorSymbols
	err = txtFormatCol.AddLine2Col(
		"FoundDecimalSeparatorSymbols",
		numStrParseResults.FoundDecimalSeparatorSymbols,
		ePrefix.XCpy(
			"Build FoundDecimalSeparatorSymbols"))

	if err != nil {
		return err
	}

	// Build FoundIntegerDigits
	err = txtFormatCol.AddLine2Col(
		"FoundIntegerDigits",
		numStrParseResults.FoundIntegerDigits,
		ePrefix.XCpy(
			"Build FoundIntegerDigits"))

	if err != nil {
		return err
	}

	// Build IdentifiedIntegerDigits

	txtStrParam =
		numStrParseResults.IdentifiedIntegerDigits

	if len(txtStrParam) == 0 {

		txtStrParam = "Identified Integer Digits is EMPTY!"

	} else {

		txtStrParam = "\"" + txtStrParam + "\""

	}

	err = txtFormatCol.AddLine2Col(
		"IdentifiedIntegerDigits",
		txtStrParam,
		ePrefix.XCpy(
			"Build IdentifiedIntegerDigits"))

	if err != nil {
		return err
	}

	// Build FoundDecimalDigits
	err = txtFormatCol.AddLine2Col(
		"FoundDecimalDigits",
		numStrParseResults.FoundDecimalDigits,
		ePrefix.XCpy(
			"Build FoundDecimalDigits"))

	if err != nil {
		return err
	}

	// Build IdentifiedFractionalDigits

	txtStrParam =
		numStrParseResults.IdentifiedFractionalDigits

	if len(txtStrParam) == 0 {

		txtStrParam = "Identified Fractional Digits is EMPTY!"

	} else {

		txtStrParam = "\"" + txtStrParam + "\""

	}

	err = txtFormatCol.AddLine2Col(
		"IdentifiedFractionalDigits",
		txtStrParam,
		ePrefix.XCpy(
			"Build IdentifiedFractionalDigits"))

	if err != nil {
		return err
	}

	// Build NumSignValue
	if !numStrParseResults.NumSignValue.XIsValid() {
		numStrParseResults.NumSignValue =
			NumSignVal.None()
	}

	err = txtFormatCol.AddLine2Col(
		"NumSignValue",
		numStrParseResults.NumSignValue,
		ePrefix.XCpy(
			"Build NumSignValue"))

	if err != nil {
		return err
	}

	// Build NumValueType

	if !numStrParseResults.NumValueType.XIsValid() {
		numStrParseResults.NumValueType =
			NumValType.None()
	}

	err = txtFormatCol.AddLine2Col(
		"NumValueType",
		numStrParseResults.NumValueType,
		ePrefix.XCpy(
			"Build NumValueType"))

	if err != nil {
		return err
	}

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Marquee-Bottom"))

	if err != nil {
		return err
	}

	// Trailing Title Marquee

	titles.Empty()

	titles.AddManyStrings(
		"CharSearchNumStrParseResultsDto",
		"End of Parameter Listing")

	err = new(TextUtility).BuildOneColTrailingMarquee(
		strBuilder,
		titles,
		txtFmtParams,
		ePrefix.XCpy(
			"strBuilder<-txtFmtParams"))

	if err != nil {
		return err
	}

	if printDetail {

		// Build RemainderString

		strLabel := "RemainderString"

		txtStrParam =
			numStrParseResults.RemainderString.GetCharacterString()

		if len(txtStrParam) == 0 {
			txtStrParam = "RemainderString is EMPTY!"
		}

		err = txtFormatCol.AddLine2Col(
			strLabel,
			txtStrParam,
			ePrefix.XCpy(
				"Build RemainderString"))

		if err != nil {
			return err
		}

		// Build DecimalSeparatorSearchResults

		err = numStrParseResults.DecimalSeparatorSearchResults.
			GetParameterTextListing(
				strBuilder,
				true,
				ePrefix.XCpy(
					"numStrParseResults"))

		if err != nil {

			return err
		}

		// Build NegativeNumberSymbolSearchResults
		err = numStrParseResults.NegativeNumberSymbolSearchResults.
			GetParameterTextListing(
				strBuilder,
				true,
				ePrefix.XCpy(
					"numStrParseResults"))

		if err != nil {

			return err
		}

		// Build ParsingTerminatorSearchResults
		err = numStrParseResults.ParsingTerminatorSearchResults.
			GetParameterTextListing(
				strBuilder,
				true,
				ePrefix.XCpy(
					"numStrParseResults"))

		if err != nil {

			return err
		}

	}

	return err
}

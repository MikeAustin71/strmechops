package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchNumStrParseResultsDtoNanobot - Provides helper
// methods for type CharSearchNumStrParseResultsDto.
//
type charSearchNumStrParseResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceNumStrParseResults' to input parameter
// 'destinationNumStrParseResults'. Both instances are of type
// CharSearchNumStrParseResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in
// 'destinationNumStrParseResults' will be overwritten.
//
// Also, NO data validation is performed on
// 'sourceNumStrParseResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  destinationNumStrParseResults  *CharSearchNumStrParseResultsDto
//     - A pointer to a CharSearchNumStrParseResultsDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceNumStrParseResults'.
//
//       'destinationNumStrParseResults' is the destination for this
//       copy operation.
//
//
//  sourceNumStrParseResults       *CharSearchNumStrParseResultsDto
//     - A pointer to another CharSearchNumStrParseResultsDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationNumStrParseResults'.
//
//       'sourceNumStrParseResults' is the source for this copy
//       operation.
//
//       No data validation is performed on 'sourceNumStrParseResults'.
//
//
//  errPrefDto                     *ePref.ErrPrefixDto
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
func (searchNumStrParseResultsNanobot *charSearchNumStrParseResultsDtoNanobot) copyIn(
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

	charSearchNumStrParseResultsDtoAtom{}.ptr().
		empty(destinationNumStrParseResults)

	destinationNumStrParseResults.SearchResultsName =
		sourceNumStrParseResults.SearchResultsName

	destinationNumStrParseResults.SearchResultsFunctionChain =
		sourceNumStrParseResults.SearchResultsFunctionChain

	destinationNumStrParseResults.FoundNumericDigits =
		sourceNumStrParseResults.FoundNumericDigits

	destinationNumStrParseResults.FoundNonZeroValue =
		sourceNumStrParseResults.FoundNonZeroValue

	destinationNumStrParseResults.FoundDecimalSeparatorSymbols =
		sourceNumStrParseResults.FoundDecimalSeparatorSymbols

	destinationNumStrParseResults.FoundIntegerDigits =
		sourceNumStrParseResults.FoundIntegerDigits

	destinationNumStrParseResults.FoundDecimalDigits =
		sourceNumStrParseResults.FoundDecimalDigits

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

// copyOut - Returns a deep copy of the input parameter
// 'numStrParseResults', a pointer to an instance of
// CharSearchNumStrParseResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'numStrParseResults'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrParseResults            *CharSearchNumStrParseResultsDto
//     - A pointer to an instance of
//       CharSearchNumStrParseResultsDto. A deep copy of the
//       internal member variables contained in this instance will
//       be created and returned in a new instance of
//       CharSearchNumStrParseResultsDto.
//
//       No data validation is performed on 'numStrParseResults'.
//
//
//  errPrefDto                    *ePref.ErrPrefixDto
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
//  deepCopyNumStrParseResults    CharSearchNumStrParseResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'numStrParseResults' will be created and
//       returned in a new instance of
//       CharSearchNumStrParseResultsDto.
//
//
//  err                           error
//     - If this method completes successfully, this returned error
//       Type is set equal to 'nil'. If errors are encountered during
//       processing, the returned error Type will encapsulate an error
//       message.
//
//       If an error message is returned, the text value for input
//       parameter 'errPrefDto' (error prefix) will be prefixed or
//       attached at the beginning of the error message.
//
func (searchNumStrParseResultsNanobot *charSearchNumStrParseResultsDtoNanobot) copyOut(
	numStrParseResults *CharSearchNumStrParseResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyNumStrParseResults CharSearchNumStrParseResultsDto,
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
			"copyOut()",
		"")

	if err != nil {

		return deepCopyNumStrParseResults, err

	}

	if numStrParseResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrParseResults' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyNumStrParseResults, err
	}

	charSearchNumStrParseResultsDtoAtom{}.ptr().
		empty(&deepCopyNumStrParseResults)

	deepCopyNumStrParseResults.SearchResultsName =
		numStrParseResults.SearchResultsName

	deepCopyNumStrParseResults.SearchResultsFunctionChain =
		numStrParseResults.SearchResultsFunctionChain

	deepCopyNumStrParseResults.FoundNumericDigits =
		numStrParseResults.FoundNumericDigits

	deepCopyNumStrParseResults.FoundNonZeroValue =
		numStrParseResults.FoundNonZeroValue

	deepCopyNumStrParseResults.FoundDecimalSeparatorSymbols =
		numStrParseResults.FoundDecimalSeparatorSymbols

	deepCopyNumStrParseResults.FoundIntegerDigits =
		numStrParseResults.FoundIntegerDigits

	deepCopyNumStrParseResults.FoundDecimalDigits =
		numStrParseResults.FoundDecimalDigits

	deepCopyNumStrParseResults.NumSignValue =
		numStrParseResults.NumSignValue

	deepCopyNumStrParseResults.NumValueType =
		numStrParseResults.NumValueType

	err = deepCopyNumStrParseResults.RemainderString.
		CopyIn(
			&numStrParseResults.RemainderString,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-numStrParseResults"))

	if err != nil {
		return deepCopyNumStrParseResults, err
	}

	err = deepCopyNumStrParseResults.DecimalSeparatorSearchResults.
		CopyIn(
			&numStrParseResults.DecimalSeparatorSearchResults,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-numStrParseResults"))

	if err != nil {
		return deepCopyNumStrParseResults, err
	}

	err = deepCopyNumStrParseResults.NegativeNumberSymbolSearchResults.
		CopyIn(
			&numStrParseResults.NegativeNumberSymbolSearchResults,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-numStrParseResults"))

	if err != nil {
		return deepCopyNumStrParseResults, err
	}

	err = deepCopyNumStrParseResults.ParsingTerminatorSearchResults.
		CopyIn(
			&numStrParseResults.ParsingTerminatorSearchResults,
			ePrefix.XCpy(
				"deepCopyNumStrParseResults"+
					"<-numStrParseResults"))

	return deepCopyNumStrParseResults, err
}

// getParameterTextListing - Returns formatted text output
// listing the member variable names and corresponding values
// contained in the 'numStrParseResults' instance of
// CharSearchNumStrParseResultsDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  numStrParseResults  *CharSearchNumStrParseResultsDto
//     - A pointer to an instance of
//       CharSearchNumStrParseResultsDto instance. Formatted
//       text output will be generated listing the member variable
//       names and their corresponding values. The formatted text
//       can then be used for screen displays, file output or
//       printing.
//
//       No data validation is performed on this instance of
//       CharSearchNumStrParseResultsDto.
//
//
//  printDetail         bool
//     - If this parameter is set to 'true', detail information for
//       subsidiary types RemainderString,
//       DecimalSeparatorSearchResults,
//       NegativeNumberSymbolSearchResults and
//       ParsingTerminatorSearchResults will be included in the
//       formatted text output.
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
//       'numStrParseResults' . This formatted text can them be used
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
func (searchNumStrParseResultsNanobot *charSearchNumStrParseResultsDtoNanobot) getParameterTextListing(
	numStrParseResults *CharSearchNumStrParseResultsDto,
	printDetail bool,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchNumStrParseResultsNanobot.lock == nil {
		searchNumStrParseResultsNanobot.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsNanobot.lock.Lock()

	defer searchNumStrParseResultsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchNumStrParseResultsDtoNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {

		return strBuilder, err

	}

	if numStrParseResults == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'numStrParseResults' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	// Total available Length of Output Line
	const maxLineLen = 79

	// Max Label Field Length = 33
	const maxLabelFieldLen = 33

	// Leading Title Marquee
	txtFormatCol := TextFormatterCollection{}

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

	// Title Line 1
	txtFormatCol.AddFieldLabel(
		"",
		"CharSearchNumStrParseResultsDto",
		maxLineLen,
		TxtJustify.Center(),
		"",
		"\n",
		-1,
		false)

	txtStrParam :=
		numStrParseResults.SearchResultsName

	if len(txtStrParam) > 0 {

		// Title Line 2
		txtFormatCol.AddFieldLabel(
			"",
			txtStrParam,
			maxLineLen,
			TxtJustify.Center(),
			"",
			"\n",
			-1,
			false)

	}

	// Title Line 3
	txtFormatCol.AddFieldLabel(
		"",
		"Parameter Listing",
		maxLineLen,
		TxtJustify.Center(),
		"",
		"\n",
		-1,
		false)

	// Title Line  4 Date/Time
	txtFormatCol.AddFieldDateTime(
		"",
		time.Now(),
		"Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		maxLineLen,
		TxtJustify.Center(),
		"",
		"\n",
		-1,
		false)

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

	// Build SearchResultsName Name

	txtStrParam = numStrParseResults.SearchResultsName

	if len(txtStrParam) == 0 {
		txtStrParam = "SearchResultsName is EMPTY!"
	}

	err = txtFormatCol.CfgLine2Col(
		" ",
		"SearchResultsName",
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
		ePrefix.XCpy("Initial 2-Col Setup"))

	if err != nil {
		return strBuilder, err
	}

	txtFormatCol.AddLineBlank(1, "")

	// Build SearchResultsFunctionChain
	txtFormatCol.AddFieldLabel(
		" ",
		"SearchResultsFunctionChain",
		maxLabelFieldLen,
		TxtJustify.Right(),
		colonSpace,
		"\n",
		-1,
		false)

	strParam := numStrParseResults.SearchResultsFunctionChain

	spacer := strings.Repeat(" ", 16)

	strParam = strings.Replace(
		strParam,
		"\n",
		"\n"+spacer,
		-1)

	strParam = "\n" + spacer + strParam

	txtFormatCol.AddFieldLabel(
		"",
		strParam,
		-1,
		TxtJustify.Left(),
		"",
		"\n",
		-1,
		false)

	// Build FoundNumericDigits
	err = txtFormatCol.AddLine2Col(
		"FoundNumericDigits",
		numStrParseResults.FoundNumericDigits,
		ePrefix.XCpy(
			"Build FoundNumericDigits"))

	if err != nil {
		return strBuilder, err
	}

	// Build FoundDecimalSeparatorSymbols
	err = txtFormatCol.AddLine2Col(
		"FoundDecimalSeparatorSymbols",
		numStrParseResults.FoundDecimalSeparatorSymbols,
		ePrefix.XCpy(
			"Build FoundDecimalSeparatorSymbols"))

	if err != nil {
		return strBuilder, err
	}

	// Build FoundIntegerDigits
	err = txtFormatCol.AddLine2Col(
		"FoundIntegerDigits",
		numStrParseResults.FoundIntegerDigits,
		ePrefix.XCpy(
			"Build FoundIntegerDigits"))

	if err != nil {
		return strBuilder, err
	}

	// Build FoundDecimalDigits
	err = txtFormatCol.AddLine2Col(
		"FoundDecimalDigits",
		numStrParseResults.FoundDecimalDigits,
		ePrefix.XCpy(
			"Build FoundDecimalDigits"))

	if err != nil {
		return strBuilder, err
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
		return strBuilder, err
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
		return strBuilder, err
	}

	var strBuilder2 strings.Builder

	if printDetail {

		// Build RemainderString

		strLabel := "RemainderString"

		strParam =
			numStrParseResults.RemainderString.GetCharacterString()

		if len(strParam) == 0 {
			strParam = "RemainderString is EMPTY!"
		}

		err = txtFormatCol.AddLine2Col(
			strLabel,
			strParam,
			ePrefix.XCpy(
				"Build RemainderString"))

		if err != nil {
			return strBuilder, err
		}

		// Build DecimalSeparatorSearchResults

		strBuilder2,
			err = txtFormatCol.BuildText(
			ePrefix.XCpy("Break Build DecimalSeparatorSearchResults"))

		if err != nil {
			return strBuilder, err
		}

		strBuilder.WriteString(strBuilder2.String())
		strBuilder2.Reset()
		txtFormatCol.EmptyFormatterCollection()

		strBuilder2,
			err = numStrParseResults.DecimalSeparatorSearchResults.
			GetParameterTextListing(
				ePrefix.XCpy(
					"numStrParseResults"))

		if err != nil {

			return strBuilder, err
		}

		strBuilder.WriteString(strBuilder2.String())

		// Reset to Empty
		strBuilder2.Reset()

		// Build NegativeNumberSymbolSearchResults
		strBuilder2,
			err = numStrParseResults.NegativeNumberSymbolSearchResults.
			GetParameterTextListing(
				ePrefix.XCpy(
					"numStrParseResults"))

		if err != nil {

			return strBuilder, err
		}

		strBuilder.WriteString(strBuilder2.String())

		// Reset to Empty
		strBuilder2.Reset()

		// Build ParsingTerminatorSearchResults
		strBuilder2,
			err = numStrParseResults.ParsingTerminatorSearchResults.
			GetParameterTextListing(
				ePrefix.XCpy(
					"numStrParseResults"))

		if err != nil {

			return strBuilder, err
		}

		strBuilder.WriteString(strBuilder2.String())

		// Reset to Empty
		strBuilder2.Reset()

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
	txtFormatCol.AddFieldLabel(
		"",
		"CharSearchNumStrParseResultsDto",
		maxLineLen,
		TxtJustify.Center(),
		"",
		"\n",
		-1,
		false)

	// Title # 2
	txtFormatCol.AddFieldLabel(
		"",
		"End of Parameter Listing",
		maxLineLen,
		TxtJustify.Center(),
		"",
		"\n",
		-1,
		false)

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

	strBuilder2,
		err = txtFormatCol.BuildText(
		ePrefix.XCpy(
			"Marquee-Bottom"))

	if err != nil {
		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	return strBuilder, err
}

// ptr - Returns a pointer to a new instance of
// charSearchNegNumResultsDtoNanobot.
//
func (searchNumStrParseResultsNanobot charSearchNumStrParseResultsDtoNanobot) ptr() *charSearchNumStrParseResultsDtoNanobot {

	if searchNumStrParseResultsNanobot.lock == nil {
		searchNumStrParseResultsNanobot.lock = new(sync.Mutex)
	}

	searchNumStrParseResultsNanobot.lock.Lock()

	defer searchNumStrParseResultsNanobot.lock.Unlock()

	return &charSearchNumStrParseResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

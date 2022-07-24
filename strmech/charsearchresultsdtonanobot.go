package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchResultsDtoNanobot - Provides helper methods for type
// CharSearchResultsDto.
//
type charSearchResultsDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'incomingSearchResultsDto' to input parameter
// 'targetSearchResultsDto'. Both instances are of type
// CharSearchResultsDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'targetSearchResultsDto'
// will be overwritten.
//
// Also, NO data validation is performed on
// 'incomingSearchResultsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  targetSearchResultsDto     *CharSearchResultsDto
//     - A pointer to a CharSearchResultsDto instance. All the
//       member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'incomingSearchResultsDto'.
//
//       'targetSearchResultsDto' is the destination for this
//       copy operation.
//
//
//  incomingSearchResultsDto          *CharSearchResultsDto
//     - A pointer to another CharSearchResultsDto instance. All
//       the member variable data values from this object will be
//       copied to corresponding member variables in
//       'targetSearchResultsDto'.
//
//       'incomingSearchResultsDto' is the source for this copy
//       operation.
//
//       If 'incomingSearchResultsDto' is determined to be invalid,
//       an error will be returned.
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
func (searchResultsDtoNanobot *charSearchResultsDtoNanobot) copyIn(
	targetSearchResultsDto *CharSearchResultsDto,
	incomingSearchResultsDto *CharSearchResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchResultsDtoNanobot.lock == nil {
		searchResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchResultsDtoNanobot.lock.Lock()

	defer searchResultsDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if targetSearchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if incomingSearchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'incomingSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchResultsDtoAtom{}.ptr().empty(
		targetSearchResultsDto)

	targetSearchResultsDto.SearchResultsName =
		incomingSearchResultsDto.SearchResultsName

	targetSearchResultsDto.SearchResultsFunctionChain =
		incomingSearchResultsDto.SearchResultsFunctionChain

	targetSearchResultsDto.FoundSearchTarget =
		incomingSearchResultsDto.FoundSearchTarget

	targetSearchResultsDto.FoundSearchTargetOnPreviousSearch =
		incomingSearchResultsDto.FoundSearchTargetOnPreviousSearch

	targetSearchResultsDto.FoundFirstNumericDigitInNumStr =
		incomingSearchResultsDto.FoundFirstNumericDigitInNumStr

	targetSearchResultsDto.FoundDecimalSeparatorSymbols =
		incomingSearchResultsDto.FoundDecimalSeparatorSymbols

	targetSearchResultsDto.FoundNonZeroValue =
		incomingSearchResultsDto.FoundNonZeroValue

	targetSearchResultsDto.TargetInputParametersName =
		incomingSearchResultsDto.TargetInputParametersName

	targetSearchResultsDto.TargetStringLength =
		incomingSearchResultsDto.TargetStringLength

	targetSearchResultsDto.TargetStringSearchLength =
		incomingSearchResultsDto.TargetStringSearchLength

	targetSearchResultsDto.TargetStringAdjustedSearchLength =
		incomingSearchResultsDto.TargetStringAdjustedSearchLength

	targetSearchResultsDto.TargetStringStartingSearchIndex =
		incomingSearchResultsDto.TargetStringStartingSearchIndex

	targetSearchResultsDto.TargetStringCurrentSearchIndex =
		incomingSearchResultsDto.TargetStringCurrentSearchIndex

	targetSearchResultsDto.TargetStringFirstFoundIndex =
		incomingSearchResultsDto.TargetStringFirstFoundIndex

	targetSearchResultsDto.TargetStringLastFoundIndex =
		incomingSearchResultsDto.TargetStringLastFoundIndex

	targetSearchResultsDto.TargetStringLastSearchIndex =
		incomingSearchResultsDto.TargetStringLastSearchIndex

	targetSearchResultsDto.TargetStringNextSearchIndex =
		incomingSearchResultsDto.TargetStringNextSearchIndex

	targetSearchResultsDto.TargetStringDescription1 =
		incomingSearchResultsDto.TargetStringDescription1

	targetSearchResultsDto.TargetStringDescription2 =
		incomingSearchResultsDto.TargetStringDescription2

	targetSearchResultsDto.TestInputParametersName =
		incomingSearchResultsDto.TestInputParametersName

	targetSearchResultsDto.TestStringName =
		incomingSearchResultsDto.TestStringName

	targetSearchResultsDto.TestStringLength =
		incomingSearchResultsDto.TestStringLength

	targetSearchResultsDto.TestStringLengthName =
		incomingSearchResultsDto.TestStringLengthName

	targetSearchResultsDto.TestStringStartingIndex =
		incomingSearchResultsDto.TestStringStartingIndex

	targetSearchResultsDto.TestStringStartingIndexName =
		incomingSearchResultsDto.TestStringStartingIndexName

	targetSearchResultsDto.TestStringFirstFoundIndex =
		incomingSearchResultsDto.TestStringFirstFoundIndex

	targetSearchResultsDto.TestStringLastFoundIndex =
		incomingSearchResultsDto.TestStringLastFoundIndex

	targetSearchResultsDto.TestStringDescription1 =
		incomingSearchResultsDto.TestStringDescription1

	targetSearchResultsDto.TestStringDescription2 =
		incomingSearchResultsDto.TestStringDescription2

	targetSearchResultsDto.CollectionTestObjIndex =
		incomingSearchResultsDto.CollectionTestObjIndex

	if incomingSearchResultsDto.RemainderString.GetRuneArrayLength() > 0 {

		err = targetSearchResultsDto.RemainderString.CopyIn(
			&incomingSearchResultsDto.RemainderString,
			ePrefix.XCpy(
				"targetSearchResultsDto.RemainderString"+
					"<-incomingSearchResultsDto.RemainderString"))

		if err != nil {

			return err

		}

	}

	if incomingSearchResultsDto.ReplacementString.GetRuneArrayLength() > 0 {

		err = targetSearchResultsDto.ReplacementString.CopyIn(
			&incomingSearchResultsDto.ReplacementString,
			ePrefix.XCpy(
				"targetSearchResultsDto.ReplacementString"+
					"<-incomingSearchResultsDto.ReplacementString"))

		if err != nil {

			return err

		}

	}

	targetSearchResultsDto.NumValueType =
		incomingSearchResultsDto.NumValueType

	targetSearchResultsDto.NumStrFormatType =
		incomingSearchResultsDto.NumStrFormatType

	targetSearchResultsDto.NumSymbolLocation =
		incomingSearchResultsDto.NumSymbolLocation

	targetSearchResultsDto.NumSymbolClass =
		incomingSearchResultsDto.NumSymbolClass

	targetSearchResultsDto.NumSignValue =
		incomingSearchResultsDto.NumSignValue

	targetSearchResultsDto.PrimaryNumSignPosition =
		incomingSearchResultsDto.PrimaryNumSignPosition

	targetSearchResultsDto.SecondaryNumSignPosition =
		incomingSearchResultsDto.SecondaryNumSignPosition

	targetSearchResultsDto.TextCharSearchType =
		incomingSearchResultsDto.TextCharSearchType

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'searchResultsDto', a pointer to an instance of
// CharSearchResultsDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'searchResultsDto'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto           *CharSearchResultsDto
//     - A pointer to an instance of CharSearchResultsDto. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of CharSearchResultsDto.
//
//       No data validation is performed on 'searchResultsDto'.
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
//  copySearchResultsDto       CharSearchResultsDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'searchResultsDto' will be created and
//       returned in a new instance of
//       CharSearchResultsDto.
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
func (searchResultsDtoNanobot *charSearchResultsDtoNanobot) copyOut(
	searchResultsDto *CharSearchResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	copySearchResultsDto CharSearchResultsDto,
	err error) {

	if searchResultsDtoNanobot.lock == nil {
		searchResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchResultsDtoNanobot.lock.Lock()

	defer searchResultsDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return copySearchResultsDto, err

	}

	if searchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetSearchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return copySearchResultsDto, err
	}

	charSearchResultsDtoAtom{}.ptr().empty(
		&copySearchResultsDto)

	copySearchResultsDto.SearchResultsName =
		searchResultsDto.SearchResultsName

	copySearchResultsDto.SearchResultsFunctionChain =
		searchResultsDto.SearchResultsFunctionChain

	copySearchResultsDto.FoundSearchTarget =
		searchResultsDto.FoundSearchTarget

	copySearchResultsDto.FoundSearchTargetOnPreviousSearch =
		searchResultsDto.FoundSearchTargetOnPreviousSearch

	copySearchResultsDto.FoundFirstNumericDigitInNumStr =
		searchResultsDto.FoundFirstNumericDigitInNumStr

	copySearchResultsDto.FoundDecimalSeparatorSymbols =
		searchResultsDto.FoundDecimalSeparatorSymbols

	copySearchResultsDto.FoundNonZeroValue =
		searchResultsDto.FoundNonZeroValue

	copySearchResultsDto.TargetInputParametersName =
		searchResultsDto.TargetInputParametersName

	copySearchResultsDto.TargetStringLength =
		searchResultsDto.TargetStringLength

	copySearchResultsDto.TargetStringSearchLength =
		searchResultsDto.TargetStringSearchLength

	copySearchResultsDto.TargetStringAdjustedSearchLength =
		searchResultsDto.TargetStringAdjustedSearchLength

	copySearchResultsDto.TargetStringStartingSearchIndex =
		searchResultsDto.TargetStringStartingSearchIndex

	copySearchResultsDto.TargetStringCurrentSearchIndex =
		searchResultsDto.TargetStringCurrentSearchIndex

	copySearchResultsDto.TargetStringFirstFoundIndex =
		searchResultsDto.TargetStringFirstFoundIndex

	copySearchResultsDto.TargetStringLastFoundIndex =
		searchResultsDto.TargetStringLastFoundIndex

	copySearchResultsDto.TargetStringLastSearchIndex =
		searchResultsDto.TargetStringLastSearchIndex

	copySearchResultsDto.TargetStringNextSearchIndex =
		searchResultsDto.TargetStringNextSearchIndex

	copySearchResultsDto.TargetStringDescription1 =
		searchResultsDto.TargetStringDescription1

	copySearchResultsDto.TargetStringDescription2 =
		searchResultsDto.TargetStringDescription2

	copySearchResultsDto.TestInputParametersName =
		searchResultsDto.TestInputParametersName

	copySearchResultsDto.TestStringName =
		searchResultsDto.TestStringName

	copySearchResultsDto.TestStringLength =
		searchResultsDto.TestStringLength

	copySearchResultsDto.TestStringLengthName =
		searchResultsDto.TestStringLengthName

	copySearchResultsDto.TestStringStartingIndex =
		searchResultsDto.TestStringStartingIndex

	copySearchResultsDto.TestStringStartingIndexName =
		searchResultsDto.TestStringStartingIndexName

	copySearchResultsDto.TestStringFirstFoundIndex =
		searchResultsDto.TestStringFirstFoundIndex

	copySearchResultsDto.TestStringLastFoundIndex =
		searchResultsDto.TestStringLastFoundIndex

	copySearchResultsDto.TestStringDescription1 =
		searchResultsDto.TestStringDescription1

	copySearchResultsDto.TestStringDescription2 =
		searchResultsDto.TestStringDescription2

	copySearchResultsDto.CollectionTestObjIndex =
		searchResultsDto.CollectionTestObjIndex

	if searchResultsDto.RemainderString.GetRuneArrayLength() > 0 {

		err = copySearchResultsDto.RemainderString.
			CopyIn(
				&searchResultsDto.RemainderString,
				ePrefix.XCpy(
					"copySearchResultsDto.RemainderString<-"+
						"searchResultsDto.RemainderString"))

		if err != nil {

			return copySearchResultsDto, err

		}

	}

	if searchResultsDto.ReplacementString.GetRuneArrayLength() > 0 {

		err = copySearchResultsDto.ReplacementString.
			CopyIn(
				&searchResultsDto.ReplacementString,
				ePrefix.XCpy(
					"copySearchResultsDto.ReplacementString<-"+
						"searchResultsDto.ReplacementString"))

		if err != nil {

			return copySearchResultsDto, err

		}

	}

	copySearchResultsDto.NumValueType =
		searchResultsDto.NumValueType

	copySearchResultsDto.NumStrFormatType =
		searchResultsDto.NumStrFormatType

	copySearchResultsDto.NumSymbolLocation =
		searchResultsDto.NumSymbolLocation

	copySearchResultsDto.NumSymbolClass =
		searchResultsDto.NumSymbolClass

	copySearchResultsDto.NumSignValue =
		searchResultsDto.NumSignValue

	copySearchResultsDto.PrimaryNumSignPosition =
		searchResultsDto.PrimaryNumSignPosition

	copySearchResultsDto.SecondaryNumSignPosition =
		searchResultsDto.SecondaryNumSignPosition

	copySearchResultsDto.TextCharSearchType =
		searchResultsDto.TextCharSearchType

	return copySearchResultsDto, err
}

// getParameterTextListing - Returns formatted text output
// detailing the member variable names and their corresponding
// values contained in the 'searchResultsDto' instance of
// CharSearchResultsDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  searchResultsDto           *CharSearchResultsDto
//     - A pointer to an instance of CharSearchResultsDto.
//       Formatted text output will be generated listing the member
//       variable names and their corresponding values. The
//       formatted text can then be used for text displays, file
//       output or printing.
//
//       No data validation is performed on this instance of
//       CharSearchResultsDto.
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
//       'searchResultsDto' . This formatted text can them be used
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
func (searchResultsDtoNanobot *charSearchResultsDtoNanobot) getParameterTextListing(
	searchResultsDto *CharSearchResultsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchResultsDtoNanobot.lock == nil {
		searchResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchResultsDtoNanobot.lock.Lock()

	defer searchResultsDtoNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"getParameterTextListing()",
		"")

	if err != nil {

		return strBuilder, err

	}

	if searchResultsDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchResultsDto' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	const maxLineLen = 79

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
		return strBuilder, err
	}

	// Leading Title Marquee

	// Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolidDto(
		solidLineDto)

	// Title # 1
	err = txtFormatCol.AddLine1Col(
		"CharSearchResultsDto",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return strBuilder, err
	}

	txtStrParam :=
		searchResultsDto.SearchResultsName

	if len(txtStrParam) > 0 {
		// Title # 2
		err = txtFormatCol.AddLine1Col(
			txtStrParam,
			ePrefix.XCpy(
				"Top-Title Line 2"))

	}

	// Title Line 3
	err = txtFormatCol.AddLine1Col(
		"Parameter Listing",
		ePrefix.XCpy(
			"Top-Title Line 3"))

	if err != nil {
		return strBuilder, err
	}

	// Text Line 4 Date Time

	err = txtFormatCol.AddLine1Col(
		TextInputParamFieldDateTimeDto{
			FieldDateTime:       time.Now(),
			FieldDateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		},
		ePrefix.XCpy(
			"Top-Title Line 4"))

	if err != nil {
		return strBuilder, err
	}

	// Filler Line '========='
	// Marquee Bottom
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Trailing Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Configure Parameter Label-Value Pairs

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
		return strBuilder, err
	}

	// Build SearchResultsName Parameter

	txtStrLabel := "SearchResultsName"

	txtStrParam = searchResultsDto.SearchResultsName

	if len(txtStrParam) == 0 {
		txtStrParam = "SearchResultsName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			""))

	if err != nil {
		return strBuilder, err
	}

	// Build SearchResultsFunctionChain

	txtStrLabel = "SearchResultsFunctionChain"

	txtStrParam = searchResultsDto.SearchResultsFunctionChain

	if len(txtStrParam) == 0 {

		txtStrParam = "SearchResultsFunctionChain is EMPTY!"

		err = txtFormatCol.AddLine2Col(
			txtStrLabel,
			txtStrParam,
			ePrefix.XCpy(
				""))

		if err != nil {
			return strBuilder, err
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

	// Build FoundSearchTarget Parameter

	txtStrLabel = "FoundSearchTarget"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.FoundSearchTarget,
		ePrefix.XCpy(
			"FoundSearchTarget"))

	if err != nil {
		return strBuilder, err
	}

	// Build FoundSearchTargetOnPreviousSearch

	txtStrLabel = "FoundSearchTargetOnPreviousSearch"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.FoundSearchTargetOnPreviousSearch,
		ePrefix.XCpy(
			"FoundSearchTargetOnPreviousSearch"))

	if err != nil {
		return strBuilder, err
	}

	// Build FoundFirstNumericDigitInNumStr

	txtStrLabel = "FoundFirstNumericDigitInNumStr"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.FoundFirstNumericDigitInNumStr,
		ePrefix.XCpy(
			"FoundFirstNumericDigitInNumStr"))

	if err != nil {
		return strBuilder, err
	}

	// Build FoundDecimalSeparatorSymbols

	txtStrLabel = "FoundDecimalSeparatorSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.FoundDecimalSeparatorSymbols,
		ePrefix.XCpy(
			"FoundDecimalSeparatorSymbols"))

	if err != nil {
		return strBuilder, err
	}

	// Build FoundNonZeroValue

	txtStrLabel = "FoundNonZeroValue"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.FoundNonZeroValue,
		ePrefix.XCpy(
			"FoundNonZeroValue"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetInputParametersName

	txtStrLabel = "TargetInputParametersName"

	txtStrParam =
		searchResultsDto.TargetInputParametersName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetInputParametersName is Empty!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetInputParametersName"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringLength

	txtStrLabel = "TargetStringLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringLength,
		ePrefix.XCpy(
			"TargetStringLength"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringSearchLength

	txtStrLabel = "TargetStringSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringSearchLength,
		ePrefix.XCpy(
			"TargetStringSearchLength"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringAdjustedSearchLength

	txtStrLabel = "TargetStringAdjustedSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringAdjustedSearchLength,
		ePrefix.XCpy(
			"TargetStringAdjustedSearchLength"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringStartingSearchIndex

	txtStrLabel = "TargetStringStartingSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringStartingSearchIndex,
		ePrefix.XCpy(
			"TargetStringStartingSearchIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringCurrentSearchIndex

	txtStrLabel = "TargetStringCurrentSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringCurrentSearchIndex,
		ePrefix.XCpy(
			"TargetStringCurrentSearchIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringFirstFoundIndex

	txtStrLabel = "TargetStringFirstFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringFirstFoundIndex,
		ePrefix.XCpy(
			"TargetStringFirstFoundIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringLastFoundIndex

	txtStrLabel = "TargetStringLastFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringLastFoundIndex,
		ePrefix.XCpy(
			"TargetStringLastFoundIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringLastSearchIndex

	txtStrLabel = "TargetStringLastSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringLastSearchIndex,
		ePrefix.XCpy(
			"TargetStringLastSearchIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringNextSearchIndex

	txtStrLabel = "TargetStringNextSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TargetStringNextSearchIndex,
		ePrefix.XCpy(
			"TargetStringNextSearchIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringDescription1

	txtStrLabel = "TargetStringDescription1"

	txtStrParam =
		searchResultsDto.TargetStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringDescription1 is Empty."
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringDescription1"))

	if err != nil {
		return strBuilder, err
	}

	// Build TargetStringDescription2

	txtStrLabel = "TargetStringDescription2"

	txtStrParam =
		searchResultsDto.TargetStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringDescription2 is Empty."
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringDescription2"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringLength

	txtStrLabel = "TestStringLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TestStringLength,
		ePrefix.XCpy(
			"TestStringLength"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringStartingIndex

	txtStrLabel = "TestStringStartingIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TestStringStartingIndex,
		ePrefix.XCpy(
			"TestStringStartingIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringFirstFoundIndex

	txtStrLabel = "TestStringFirstFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TestStringFirstFoundIndex,
		ePrefix.XCpy(
			"TestStringFirstFoundIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringLastFoundIndex

	txtStrLabel = "TestStringLastFoundIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TestStringLastFoundIndex,
		ePrefix.XCpy(
			"TestStringLastFoundIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringDescription1

	txtStrLabel = "TestStringDescription1"

	txtStrParam =
		searchResultsDto.TestStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringDescription1 is Empty."
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringDescription1"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringDescription2

	txtStrLabel = "TestStringDescription2"

	txtStrParam =
		searchResultsDto.TestStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringDescription2 is Empty."
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringDescription2"))

	if err != nil {
		return strBuilder, err
	}

	// Build CollectionTestObjIndex

	txtStrLabel = "CollectionTestObjIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.CollectionTestObjIndex,
		ePrefix.XCpy(
			"CollectionTestObjIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build RemainderString

	txtStrLabel = "RemainderString"

	txtStrParam =
		searchResultsDto.RemainderString.GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "RemainderString is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"RemainderString"))

	if err != nil {
		return strBuilder, err
	}

	// Build ReplacementString

	txtStrLabel = "ReplacementString"

	txtStrParam =
		searchResultsDto.ReplacementString.GetCharacterString()

	if len(txtStrParam) == 0 {
		txtStrParam = "ReplacementString is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"ReplacementString"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumValueType

	txtStrLabel = "NumValueType"

	if !searchResultsDto.NumValueType.XIsValid() {
		searchResultsDto.NumValueType =
			NumValType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.NumValueType,
		ePrefix.XCpy(
			"NumValueType"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumStrFormatType

	txtStrLabel = "NumStrFormatType"

	if !searchResultsDto.NumStrFormatType.XIsValid() {
		searchResultsDto.NumStrFormatType =
			NumStrFmtType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.NumStrFormatType,
		ePrefix.XCpy(
			"NumStrFormatType"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumSymLocation

	txtStrLabel = "NumSymLocation"

	if !searchResultsDto.NumSymbolLocation.XIsValid() {
		searchResultsDto.NumSymbolLocation =
			NumSymLocation.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.NumSymbolLocation,
		ePrefix.XCpy(
			"NumSymbolLocation"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumSymbolClass

	txtStrLabel = "NumSymbolClass"

	if !searchResultsDto.NumSymbolClass.XIsValid() {
		searchResultsDto.NumSymbolClass =
			NumSymClass.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.NumSymbolClass,
		ePrefix.XCpy(
			"NumSymbolClass"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumSignValue

	txtStrLabel = "NumSignValue"

	if !searchResultsDto.NumSignValue.XIsValid() {
		searchResultsDto.NumSignValue =
			NumSignVal.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.NumSignValue,
		ePrefix.XCpy(
			"NumSignValue"))

	if err != nil {
		return strBuilder, err
	}

	// Build PrimaryNumSignPosition

	txtStrLabel = "PrimaryNumSignPosition"

	if !searchResultsDto.PrimaryNumSignPosition.XIsValid() {
		searchResultsDto.PrimaryNumSignPosition =
			NumSignSymPos.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.PrimaryNumSignPosition,
		ePrefix.XCpy(
			"PrimaryNumSignPosition"))

	if err != nil {
		return strBuilder, err
	}

	// Build SecondaryNumSignPosition

	txtStrLabel = "SecondaryNumSignPosition"

	if !searchResultsDto.SecondaryNumSignPosition.XIsValid() {
		searchResultsDto.SecondaryNumSignPosition =
			NumSignSymPos.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.SecondaryNumSignPosition,
		ePrefix.XCpy(
			"SecondaryNumSignPosition"))

	if err != nil {
		return strBuilder, err
	}

	// Build TextCharSearchType

	txtStrLabel = "TextCharSearchType"

	if !searchResultsDto.TextCharSearchType.XIsValid() {
		searchResultsDto.TextCharSearchType =
			CharSearchType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchResultsDto.TextCharSearchType,
		ePrefix.XCpy(
			"TextCharSearchType"))

	if err != nil {
		return strBuilder, err
	}

	// Trailing Title Marquee
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolidDto(
		solidLineDto)

	// Title # 1
	err = txtFormatCol.AddLine1Col(
		"CharSearchResultsDto",
		ePrefix.XCpy(
			"Bottom-Title Line 1"))

	if err != nil {
		return strBuilder, err
	}

	// Title # 2
	err = txtFormatCol.AddLine1Col(
		"End of Parameter Listing",
		ePrefix.XCpy(
			"Bottom-Title Line 2"))

	if err != nil {
		return strBuilder, err
	}

	// Filler =======
	// Marquee Bottom
	txtFormatCol.AddLineSolidDto(
		solidLineDto)

	// Blank Line
	txtFormatCol.AddLineBlank(
		2,
		"")

	var strBuilder2 strings.Builder

	strBuilder2,
		err = txtFormatCol.BuildText(
		ePrefix.XCpy(
			"Final Text Output"))

	if err != nil {
		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	return strBuilder, err
}

// ptr - Returns a pointer to a new instance of
// charSearchResultsDtoNanobot.
//
func (searchResultsDtoNanobot charSearchResultsDtoNanobot) ptr() *charSearchResultsDtoNanobot {

	if searchResultsDtoNanobot.lock == nil {
		searchResultsDtoNanobot.lock = new(sync.Mutex)
	}

	searchResultsDtoNanobot.lock.Lock()

	defer searchResultsDtoNanobot.lock.Unlock()

	return &charSearchResultsDtoNanobot{
		lock: new(sync.Mutex),
	}
}

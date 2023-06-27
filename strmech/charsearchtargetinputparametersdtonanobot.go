package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchTargetInputParametersDtoNanobot - Provides helper methods for type
// CharSearchTargetInputParametersDto.
type charSearchTargetInputParametersDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceTargetInputParms' to input parameter
// 'destinationTargetInputParms'. Both instances are of type
// CharSearchTargetInputParametersDto.
//
// # IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationTargetInputParms'
// will be overwritten.
//
// Also, NO validation is performed on 'sourceTargetInputParms'.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	destinationTargetInputParms     *CharSearchTargetInputParametersDto
//	   - A pointer to a CharSearchTargetInputParametersDto instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceTargetInputParms'.
//
//	     'destinationTargetInputParms' is the destination for this
//	     copy operation.
//
//
//	sourceTargetInputParms          *CharSearchTargetInputParametersDto
//	   - A pointer to another CharSearchTargetInputParametersDto
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationTargetInputParms'.
//
//	     'sourceTargetInputParms' is the source for this copy
//	     operation.
//
//	     No data validation is performed on
//	     'sourceTargetInputParms'.
//
//
//	errPrefDto          *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
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
func (searchTargetInputParmsNanobot *charSearchTargetInputParametersDtoNanobot) copyIn(
	destinationTargetInputParms *CharSearchTargetInputParametersDto,
	sourceTargetInputParms *CharSearchTargetInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationTargetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTargetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTargetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTargetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	targetInputParmsAtom := charSearchTargetInputParametersDtoAtom{}

	targetInputParmsAtom.empty(
		destinationTargetInputParms)

	if sourceTargetInputParms.TargetString != nil {

		destinationTargetInputParms.TargetString =
			&RuneArrayDto{}

		err = destinationTargetInputParms.TargetString.CopyIn(
			sourceTargetInputParms.TargetString,
			ePrefix.XCpy("destinationTargetInputParms<-"+
				"sourceTargetInputParms"))

		if err != nil {
			return err
		}
	}

	destinationTargetInputParms.TargetInputParametersName =
		sourceTargetInputParms.TargetInputParametersName

	destinationTargetInputParms.TargetStringName =
		sourceTargetInputParms.TargetStringName

	destinationTargetInputParms.TargetStringLength =
		sourceTargetInputParms.TargetStringLength

	destinationTargetInputParms.TargetStringLengthName =
		sourceTargetInputParms.TargetStringLengthName

	destinationTargetInputParms.TargetStringStartingSearchIndex =
		sourceTargetInputParms.TargetStringStartingSearchIndex

	destinationTargetInputParms.TargetStringCurrentSearchIndex =
		sourceTargetInputParms.TargetStringCurrentSearchIndex

	destinationTargetInputParms.TargetStringNextSearchIndex =
		sourceTargetInputParms.TargetStringNextSearchIndex

	destinationTargetInputParms.TargetStringStartingSearchIndexName =
		sourceTargetInputParms.TargetStringStartingSearchIndexName

	destinationTargetInputParms.TargetStringSearchLength =
		sourceTargetInputParms.TargetStringSearchLength

	destinationTargetInputParms.TargetStringSearchLengthName =
		sourceTargetInputParms.TargetStringSearchLengthName

	destinationTargetInputParms.TargetStringAdjustedSearchLength =
		sourceTargetInputParms.TargetStringAdjustedSearchLength

	destinationTargetInputParms.TargetStringDescription1 =
		sourceTargetInputParms.TargetStringDescription1

	destinationTargetInputParms.TargetStringDescription2 =
		sourceTargetInputParms.TargetStringDescription2

	destinationTargetInputParms.FoundFirstNumericDigitInNumStr =
		sourceTargetInputParms.FoundFirstNumericDigitInNumStr

	destinationTargetInputParms.FoundDecimalSeparatorSymbols =
		sourceTargetInputParms.FoundDecimalSeparatorSymbols

	destinationTargetInputParms.FoundNonZeroValue =
		sourceTargetInputParms.FoundNonZeroValue

	destinationTargetInputParms.TextCharSearchType =
		sourceTargetInputParms.TextCharSearchType

	destinationTargetInputParms.RequestFoundTestCharacters =
		sourceTargetInputParms.RequestFoundTestCharacters

	destinationTargetInputParms.RequestRemainderString =
		sourceTargetInputParms.RequestRemainderString

	destinationTargetInputParms.RequestReplacementString =
		sourceTargetInputParms.RequestReplacementString

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'targetInputParms', a pointer to an instance of
// CharSearchTargetInputParametersDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// NO validation is performed on 'targetInputParms'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	targetInputParms                *CharSearchTargetInputParametersDto
//	   - A pointer to an instance of CharSearchTargetInputParametersDto. A
//	     deep copy of the internal member variables will be created
//	     and returned in a new instance of CharSearchTargetInputParametersDto.
//
//	     No data validation is performed on 'targetInputParms'.
//
//
//	errPrefDto                      *ePref.ErrPrefixDto
//	   - This object encapsulates an error prefix string which is
//	     included in all returned error messages. Usually, it
//	     contains the name of the calling method or methods listed
//	     as a function chain.
//
//	     If no error prefix information is needed, set this parameter
//	     to 'nil'.
//
//	     Type ErrPrefixDto is included in the 'errpref' software
//	     package, "github.com/MikeAustin71/errpref".
//
// ------------------------------------------------------------------------
//
// Return Values
//
//	deepCopyTargetInputParms        CharSearchTargetInputParametersDto
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'targetInputParms' will be created and
//	     returned in a new instance of CharSearchTargetInputParametersDto.
//
//
//	err                             error
//	   - If this method completes successfully, this returned error
//	     Type is set equal to 'nil'. If errors are encountered during
//	     processing, the returned error Type will encapsulate an error
//	     message.
//
//	     If an error message is returned, the text value for input
//	     parameter 'errPrefDto' (error prefix) will be prefixed or
//	     attached at the beginning of the error message.
func (searchTargetInputParmsNanobot *charSearchTargetInputParametersDtoNanobot) copyOut(
	targetInputParms *CharSearchTargetInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyTargetInputParms CharSearchTargetInputParametersDto,
	err error) {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	deepCopyTargetInputParms.Empty()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyTargetInputParms, err
	}

	if targetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyTargetInputParms, err
	}

	if targetInputParms.TargetString != nil {

		deepCopyTargetInputParms.TargetString =
			&RuneArrayDto{}

		err = deepCopyTargetInputParms.TargetString.CopyIn(
			targetInputParms.TargetString,
			ePrefix.XCpy(
				"deepCopyTargetInputParms.TargetString<-"+
					"targetInputParms.TargetString"))

		if err != nil {
			return deepCopyTargetInputParms, err
		}

	}

	deepCopyTargetInputParms.TargetInputParametersName =
		targetInputParms.TargetInputParametersName

	deepCopyTargetInputParms.TargetStringName =
		targetInputParms.TargetStringName

	deepCopyTargetInputParms.TargetStringLength =
		targetInputParms.TargetStringLength

	deepCopyTargetInputParms.TargetStringLengthName =
		targetInputParms.TargetStringLengthName

	deepCopyTargetInputParms.TargetStringStartingSearchIndex =
		targetInputParms.TargetStringStartingSearchIndex

	deepCopyTargetInputParms.TargetStringCurrentSearchIndex =
		targetInputParms.TargetStringCurrentSearchIndex

	deepCopyTargetInputParms.TargetStringNextSearchIndex =
		targetInputParms.TargetStringNextSearchIndex

	deepCopyTargetInputParms.TargetStringStartingSearchIndexName =
		targetInputParms.TargetStringStartingSearchIndexName

	deepCopyTargetInputParms.TargetStringSearchLength =
		targetInputParms.TargetStringSearchLength

	deepCopyTargetInputParms.TargetStringSearchLengthName =
		targetInputParms.TargetStringSearchLengthName

	deepCopyTargetInputParms.TargetStringAdjustedSearchLength =
		targetInputParms.TargetStringAdjustedSearchLength

	deepCopyTargetInputParms.TargetStringDescription1 =
		targetInputParms.TargetStringDescription1

	deepCopyTargetInputParms.TargetStringDescription2 =
		targetInputParms.TargetStringDescription2

	deepCopyTargetInputParms.FoundFirstNumericDigitInNumStr =
		targetInputParms.FoundFirstNumericDigitInNumStr

	deepCopyTargetInputParms.FoundDecimalSeparatorSymbols =
		targetInputParms.FoundDecimalSeparatorSymbols

	deepCopyTargetInputParms.FoundNonZeroValue =
		targetInputParms.FoundNonZeroValue

	deepCopyTargetInputParms.TextCharSearchType =
		targetInputParms.TextCharSearchType

	deepCopyTargetInputParms.RequestFoundTestCharacters =
		targetInputParms.RequestFoundTestCharacters

	deepCopyTargetInputParms.RequestRemainderString =
		targetInputParms.RequestRemainderString

	deepCopyTargetInputParms.RequestReplacementString =
		targetInputParms.RequestReplacementString

	return deepCopyTargetInputParms, err
}

// getParameterTextListing - Returns formatted text output
// listing the member variable names and corresponding values
// contained in the 'targetInputParms' instance of
// CharSearchTargetInputParametersDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	targetInputParms           *CharSearchTargetInputParametersDto
//	   - A pointer to an instance of
//	     CharSearchTargetInputParametersDto instance. Formatted
//	     text output will be generated listing the member variable
//	     names and their corresponding values. The formatted text
//	     can then be used for text displays, file output or
//	     printing.
//
//	     No data validation is performed on this instance of
//	     CharSearchTargetInputParametersDto.
//
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
//	strings.Builder
//	   - If this method completes successfully, an instance of
//	     strings.Builder will be returned. This instance contains
//	     the formatted text output listing the member variable
//	     names and their corresponding values for input parameter
//	     'targetInputParms' . This formatted text can them be used
//	     for text displays, file output or printing.
//
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
func (searchTargetInputParmsNanobot charSearchTargetInputParametersDtoNanobot) getParameterTextListing(
	strBuilder *strings.Builder,
	targetInputParms *CharSearchTargetInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
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

	if targetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	// Maximum Available Output Line Length
	const maxLineLen = 78

	// Max Label Field Length = 36
	const maxLabelFieldLen = 36

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
		"",
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
		"CharSearchTargetInputParametersDto",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return err
	}

	txtStrParam :=
		targetInputParms.TargetInputParametersName

	if len(txtStrParam) > 0 {

		// Title Line 2

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

	// Title Line 4 Date/Time

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

	// Build TargetInputParametersName

	txtStrLabel := "TargetInputParametersName"

	txtStrParam =
		targetInputParms.TargetInputParametersName

	if len(txtStrParam) == 0 {
		txtStrParam =
			"TargetInputParametersName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetInputParametersName"))

	if err != nil {
		return err
	}

	// Build Formatted Target String

	txtStrLabel = "TargetString"

	if targetInputParms.TargetString == nil {

		txtStrParam =
			"Not Set. TargetString is a nil pointer"

		err = txtFormatCol.AddLine2Col(
			txtStrLabel,
			txtStrParam,
			ePrefix.XCpy(
				""))

		if err != nil {
			return err
		}

	} else {

		txtStrParam = targetInputParms.TargetString.GetCharacterString()

		if len(txtStrParam) == 0 {
			txtStrParam = "TargetString is NOT NIL and is an Empty string."
		}

		err = txtFormatCol.AddLine2Col(
			txtStrLabel,
			txtStrParam,
			ePrefix.XCpy(
				"TargetString"))

		if err != nil {
			return err
		}

	}

	// TargetStringName

	txtStrLabel = "TargetStringName"

	txtStrParam =
		targetInputParms.TargetStringName

	if len(txtStrParam) == 0 {
		txtStrParam =
			"TargetStringName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringName"))

	if err != nil {
		return err
	}

	// TargetStringLength

	txtStrLabel = "TargetStringLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.TargetStringLength,
		ePrefix.XCpy(
			"TargetStringLength"))

	if err != nil {
		return err
	}

	// TargetStringLengthName

	txtStrLabel = "TargetStringLengthName"

	txtStrParam =
		targetInputParms.TargetStringLengthName

	if len(txtStrParam) == 0 {
		txtStrParam =
			"TargetStringLengthName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringLengthName"))

	if err != nil {
		return err
	}

	// TargetStringStartingSearchIndex

	txtStrLabel = "TargetStringStartingSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.TargetStringStartingSearchIndex,
		ePrefix.XCpy(
			"TargetStringStartingSearchIndex"))

	if err != nil {
		return err
	}

	// TargetStringStartingSearchIndexName

	txtStrLabel = "TargetStringStartingSearchIndexName"

	txtStrParam =
		targetInputParms.TargetStringStartingSearchIndexName

	if len(txtStrParam) == 0 {
		txtStrParam =
			"TargetStringStartingSearchIndexName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringStartingSearchIndexName"))

	if err != nil {
		return err
	}

	// TargetStringCurrentSearchIndex
	txtStrLabel = "TargetStringCurrentSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.TargetStringCurrentSearchIndex,
		ePrefix.XCpy(
			"TargetStringCurrentSearchIndex"))

	if err != nil {
		return err
	}

	// TargetStringNextSearchIndex

	txtStrLabel = "TargetStringNextSearchIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.TargetStringNextSearchIndex,
		ePrefix.XCpy(
			"TargetStringNextSearchIndex"))

	if err != nil {
		return err
	}

	// TargetStringSearchLength

	txtStrLabel = "TargetStringSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.TargetStringSearchLength,
		ePrefix.XCpy(
			"TargetStringSearchLength"))

	if err != nil {
		return err
	}

	// TargetStringSearchLengthName

	txtStrLabel = "TargetStringSearchLengthName"

	txtStrParam =
		targetInputParms.TargetStringSearchLengthName

	if len(txtStrParam) == 0 {
		txtStrParam =
			"TargetStringSearchLengthName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringSearchLength"))

	if err != nil {
		return err
	}

	// TargetStringAdjustedSearchLength

	txtStrLabel = "TargetStringAdjustedSearchLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.TargetStringAdjustedSearchLength,
		ePrefix.XCpy(
			"TargetStringAdjustedSearchLength"))

	if err != nil {
		return err
	}

	// TargetStringDescription1

	txtStrLabel = "TargetStringDescription1"

	txtStrParam =
		targetInputParms.TargetStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam =
			"TargetStringDescription1 is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringDescription1"))

	if err != nil {
		return err
	}

	// TargetStringDescription2

	txtStrLabel = "TargetStringDescription2"

	txtStrParam =
		targetInputParms.TargetStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam =
			"TargetStringDescription2 is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TargetStringDescription2"))

	if err != nil {
		return err
	}

	// FoundFirstNumericDigitInNumStr

	txtStrLabel = "FoundFirstNumericDigitInNumStr"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.FoundFirstNumericDigitInNumStr,
		ePrefix.XCpy(
			"FoundFirstNumericDigitInNumStr"))

	if err != nil {
		return err
	}

	// FoundDecimalSeparatorSymbols

	txtStrLabel = "FoundDecimalSeparatorSymbols"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.FoundDecimalSeparatorSymbols,
		ePrefix.XCpy(
			"FoundDecimalSeparatorSymbols"))

	if err != nil {
		return err
	}

	// FoundNonZeroValue

	txtStrLabel = "FoundNonZeroValue"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.FoundNonZeroValue,
		ePrefix.XCpy(
			"FoundNonZeroValue"))

	if err != nil {
		return err
	}

	// TextCharSearchType

	txtStrLabel = "TextCharSearchType"

	if !targetInputParms.TextCharSearchType.XIsValid() {
		targetInputParms.TextCharSearchType = CharSearchType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.TextCharSearchType,
		ePrefix.XCpy(
			"TextCharSearchType"))

	if err != nil {
		return err
	}

	// RequestFoundTestCharacters

	txtStrLabel = "RequestFoundTestCharacters"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.RequestFoundTestCharacters,
		ePrefix.XCpy(
			"RequestFoundTestCharacters"))

	if err != nil {
		return err
	}

	// RequestRemainderString

	txtStrLabel = "RequestRemainderString"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.RequestRemainderString,
		ePrefix.XCpy(
			"RequestRemainderString"))

	if err != nil {
		return err
	}

	// RequestReplacementString

	txtStrLabel = "RequestReplacementString"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		targetInputParms.RequestReplacementString,
		ePrefix.XCpy(
			"RequestReplacementString"))

	if err != nil {
		return err
	}

	// Trailing Title Marquee
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Title # 1
	err = txtFormatCol.AddLine1Col(
		"CharSearchTargetInputParametersDto",
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
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Text Output"))

	return err
}

// ptr - Returns a pointer to a new instance of
// charSearchTargetInputParametersDtoNanobot.
func (searchTargetInputParmsNanobot charSearchTargetInputParametersDtoNanobot) ptr() *charSearchTargetInputParametersDtoNanobot {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	return &charSearchTargetInputParametersDtoNanobot{
		lock: new(sync.Mutex),
	}
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchTestConfigDtoNanobot - Provides helper methods for type
// CharSearchTestConfigDto.
type charSearchTestConfigDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceTestCfgDto' to input parameter
// 'destinationTestCfgDto'. Both instances are of type
// CharSearchTestConfigDto.
//
// # IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationTestCfgDto'
// will be deleted and overwritten.
//
// Also, NO validation is performed on 'sourceTestCfgDto'.
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//	destinationTestCfgDto     *CharSearchTestConfigDto
//	   - A pointer to a CharSearchTestConfigDto instance.
//	     All the member variable data fields in this object will be
//	     replaced by data values copied from input parameter
//	     'sourceTestCfgDto'.
//
//	     'destinationTestCfgDto' is the destination for this
//	     copy operation.
//
//
//	sourceTestCfgDto          *CharSearchTestConfigDto
//	   - A pointer to another CharSearchTestConfigDto
//	     instance. All the member variable data values from this
//	     object will be copied to corresponding member variables in
//	     'destinationTestCfgDto'.
//
//	     'sourceTestCfgDto' is the source for this copy
//	     operation.
//
//	     No data validation is performed on
//	     'sourceTestCfgDto'.
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
func (searchTestConfigNanobot *charSearchTestConfigDtoNanobot) copyIn(
	destinationTestCfgDto *CharSearchTestConfigDto,
	sourceTestCfgDto *CharSearchTestConfigDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestConfigDtoNanobot."+
			"copyIn()",
		"")

	if err != nil {

		return err

	}

	if destinationTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchTestConfigDtoAtom{}.ptr().empty(
		destinationTestCfgDto)

	destinationTestCfgDto.TestInputParametersName =
		sourceTestCfgDto.TestInputParametersName

	destinationTestCfgDto.TestStringName =
		sourceTestCfgDto.TestStringName

	destinationTestCfgDto.TestStringLengthName =
		sourceTestCfgDto.TestStringLengthName

	destinationTestCfgDto.TestStringStartingIndex =
		sourceTestCfgDto.TestStringStartingIndex

	destinationTestCfgDto.TestStringStartingIndexName =
		sourceTestCfgDto.TestStringStartingIndexName

	destinationTestCfgDto.TestStringDescription1 =
		sourceTestCfgDto.TestStringDescription1

	destinationTestCfgDto.TestStringDescription2 =
		sourceTestCfgDto.TestStringDescription2

	destinationTestCfgDto.CollectionTestObjIndex =
		sourceTestCfgDto.CollectionTestObjIndex

	destinationTestCfgDto.NumValueType =
		sourceTestCfgDto.NumValueType

	destinationTestCfgDto.NumStrFormatType =
		sourceTestCfgDto.NumStrFormatType

	destinationTestCfgDto.NumSymbolLocation =
		sourceTestCfgDto.NumSymbolLocation

	destinationTestCfgDto.NumSymbolClass =
		sourceTestCfgDto.NumSymbolClass

	destinationTestCfgDto.NumSignValue =
		sourceTestCfgDto.NumSignValue

	destinationTestCfgDto.PrimaryNumSignPosition =
		sourceTestCfgDto.PrimaryNumSignPosition

	destinationTestCfgDto.SecondaryNumSignPosition =
		sourceTestCfgDto.SecondaryNumSignPosition

	destinationTestCfgDto.TextCharSearchType =
		sourceTestCfgDto.TextCharSearchType

	destinationTestCfgDto.RequestFoundTestCharacters =
		sourceTestCfgDto.RequestFoundTestCharacters

	destinationTestCfgDto.RequestRemainderString =
		sourceTestCfgDto.RequestRemainderString

	destinationTestCfgDto.RequestReplacementString =
		sourceTestCfgDto.RequestReplacementString

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'searchTestCfgDto', a pointer to an instance of
// CharSearchTestConfigDto.
//
// ----------------------------------------------------------------
//
// # IMPORTANT
//
// NO validation is performed on 'searchTestCfgDto'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//	searchTestCfgDto                *CharSearchTestConfigDto
//	   - A pointer to an instance of CharSearchTestConfigDto. A
//	     deep copy of the internal member variables will be created
//	     and returned in a new instance of CharSearchTestConfigDto.
//
//	     No data validation is performed on 'searchTestCfgDto'.
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
//	deepCopySearchTestCfgDto        CharSearchTestConfigDto
//	   - If this method completes successfully, a deep copy of
//	     input parameter 'searchTestCfgDto' will be created and
//	     returned in a new instance of CharSearchTestConfigDto.
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
func (searchTestConfigNanobot *charSearchTestConfigDtoNanobot) copyOut(
	searchTestCfgDto *CharSearchTestConfigDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopySearchTestCfgDto CharSearchTestConfigDto,
	err error) {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestConfigDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopySearchTestCfgDto, err

	}

	if searchTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return deepCopySearchTestCfgDto, err
	}

	charSearchTestConfigDtoAtom{}.ptr().empty(
		&deepCopySearchTestCfgDto)

	deepCopySearchTestCfgDto.TestInputParametersName =
		searchTestCfgDto.TestInputParametersName

	deepCopySearchTestCfgDto.TestStringName =
		searchTestCfgDto.TestStringName

	deepCopySearchTestCfgDto.TestStringLengthName =
		searchTestCfgDto.TestStringLengthName

	deepCopySearchTestCfgDto.TestStringStartingIndex =
		searchTestCfgDto.TestStringStartingIndex

	deepCopySearchTestCfgDto.TestStringStartingIndexName =
		searchTestCfgDto.TestStringStartingIndexName

	deepCopySearchTestCfgDto.TestStringDescription1 =
		searchTestCfgDto.TestStringDescription1

	deepCopySearchTestCfgDto.TestStringDescription2 =
		searchTestCfgDto.TestStringDescription2

	deepCopySearchTestCfgDto.CollectionTestObjIndex =
		searchTestCfgDto.CollectionTestObjIndex

	deepCopySearchTestCfgDto.NumValueType =
		searchTestCfgDto.NumValueType

	deepCopySearchTestCfgDto.NumStrFormatType =
		searchTestCfgDto.NumStrFormatType

	deepCopySearchTestCfgDto.NumSymbolLocation =
		searchTestCfgDto.NumSymbolLocation

	deepCopySearchTestCfgDto.NumSymbolClass =
		searchTestCfgDto.NumSymbolClass

	deepCopySearchTestCfgDto.NumSignValue =
		searchTestCfgDto.NumSignValue

	deepCopySearchTestCfgDto.PrimaryNumSignPosition =
		searchTestCfgDto.PrimaryNumSignPosition

	deepCopySearchTestCfgDto.SecondaryNumSignPosition =
		searchTestCfgDto.SecondaryNumSignPosition

	deepCopySearchTestCfgDto.TextCharSearchType =
		searchTestCfgDto.TextCharSearchType

	deepCopySearchTestCfgDto.RequestFoundTestCharacters =
		searchTestCfgDto.RequestFoundTestCharacters

	deepCopySearchTestCfgDto.RequestRemainderString =
		searchTestCfgDto.RequestRemainderString

	deepCopySearchTestCfgDto.RequestReplacementString =
		searchTestCfgDto.RequestReplacementString

	return deepCopySearchTestCfgDto, err
}

// getParameterTextListing - Returns formatted text output
// detailing the member variable values contained in the
// 'searchTestCfgDto' instance of CharSearchTestConfigDto.
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	searchTestCfgDto           *CharSearchTestInputParametersDto
//	   - A pointer to an instance of
//	     CharSearchTestInputParametersDto instance. Formatted
//	     text output will be generated listing the member variable
//	     names and their corresponding values. The formatted text
//	     can then be used for text displays, file output or
//	     printing.
//
//	     No data validation is performed on this instance of
//	     CharSearchTestInputParametersDto.
//
//
//	errPrefDto                 *ePref.ErrPrefixDto
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
//	     'searchTestCfgDto' . This formatted text can then be used
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
func (searchTestConfigNanobot *charSearchTestConfigDtoNanobot) getParameterTextListing(
	strBuilder *strings.Builder,
	searchTestCfgDto *CharSearchTestConfigDto,
	errPrefDto *ePref.ErrPrefixDto) error {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestConfigDtoNanobot."+
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

	if searchTestCfgDto == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'searchTestCfgDto' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	// Total available Length of Output Line
	const maxLineLen = 79

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
		"CharSearchTestConfigDto",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return err
	}

	txtStrParam :=
		searchTestCfgDto.TestInputParametersName

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

	// Build TestInputParametersName

	txtStrLabel := "TestInputParametersName"

	txtStrParam =
		searchTestCfgDto.TestInputParametersName

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
		searchTestCfgDto.TestStringName

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

	// Build TestStringLengthName
	txtStrLabel = "TestStringLengthName"

	txtStrParam =
		searchTestCfgDto.TestStringLengthName

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
		searchTestCfgDto.TestStringStartingIndex,
		ePrefix.XCpy(
			"TestStringStartingIndex"))

	if err != nil {
		return err
	}

	// Build TestStringStartingIndexName
	txtStrLabel = "TestStringStartingIndexName"

	txtStrParam =
		searchTestCfgDto.TestStringStartingIndexName

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

	// Build TestStringDescription1
	txtStrLabel = "TestStringDescription1"

	txtStrParam =
		searchTestCfgDto.TestStringDescription1

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
		searchTestCfgDto.TestStringDescription2

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
		searchTestCfgDto.CollectionTestObjIndex,
		ePrefix.XCpy(
			"CollectionTestObjIndex"))

	if err != nil {
		return err
	}

	// Build NumValueType

	txtStrLabel = "NumValueType"

	txtStrParam =
		searchTestCfgDto.NumValueType.
			XReturnNoneIfInvalid().String()

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"NumValueType"))

	if err != nil {
		return err
	}

	// Build NumStrFormatType

	txtStrLabel = "NumStrFormatType"

	if !searchTestCfgDto.NumStrFormatType.XIsValid() {
		searchTestCfgDto.NumStrFormatType = NumStrFmtType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.NumStrFormatType,
		ePrefix.XCpy(
			"NumStrFormatType"))

	if err != nil {
		return err
	}

	// Build NumSymbolLocation

	txtStrLabel = "NumSymbolLocation"

	if !searchTestCfgDto.NumSymbolLocation.XIsValid() {
		searchTestCfgDto.NumSymbolLocation = NumSymLocation.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.NumSymbolLocation,
		ePrefix.XCpy(
			"NumSymbolLocation"))

	if err != nil {
		return err
	}

	// Build NumSymbolClass

	txtStrLabel = "NumSymbolClass"

	if !searchTestCfgDto.NumSymbolClass.XIsValid() {
		searchTestCfgDto.NumSymbolClass = NumSymClass.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.NumSymbolClass,
		ePrefix.XCpy(
			"NumSymbolClass"))

	if err != nil {
		return err
	}

	// Build NumSignValue

	txtStrLabel = "NumSignValue"

	if !searchTestCfgDto.NumSignValue.XIsValid() {
		searchTestCfgDto.NumSignValue = NumSignVal.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.NumSignValue,
		ePrefix.XCpy(
			"NumSignValue"))

	if err != nil {
		return err
	}

	// Build PrimaryNumSignPosition

	txtStrLabel = "PrimaryNumSignPosition"

	if !searchTestCfgDto.PrimaryNumSignPosition.XIsValid() {

		searchTestCfgDto.PrimaryNumSignPosition =
			NumSignSymPos.None()

	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.PrimaryNumSignPosition,
		ePrefix.XCpy(
			"PrimaryNumSignPosition"))

	if err != nil {
		return err
	}

	// Build SecondaryNumSignPosition

	txtStrLabel = "SecondaryNumSignPosition"

	if !searchTestCfgDto.SecondaryNumSignPosition.XIsValid() {

		searchTestCfgDto.SecondaryNumSignPosition =
			NumSignSymPos.None()

	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.SecondaryNumSignPosition,
		ePrefix.XCpy(
			"SecondaryNumSignPosition"))

	if err != nil {
		return err
	}

	// Build TextCharSearchType

	txtStrLabel = "TextCharSearchType"

	if !searchTestCfgDto.TextCharSearchType.XIsValid() {

		searchTestCfgDto.TextCharSearchType =
			CharSearchType.None()

	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.TextCharSearchType,
		ePrefix.XCpy(
			"TextCharSearchType"))

	if err != nil {
		return err
	}

	// Build RequestFoundTestCharacters
	txtStrLabel = "RequestFoundTestCharacters"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.RequestFoundTestCharacters,
		ePrefix.XCpy(
			"RequestFoundTestCharacters"))

	if err != nil {
		return err
	}

	// Build RequestRemainderString
	txtStrLabel = "RequestRemainderString"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.RequestRemainderString,
		ePrefix.XCpy(
			"RequestRemainderString"))

	if err != nil {
		return err
	}

	// Build RequestReplacementString
	txtStrLabel = "RequestReplacementString"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		searchTestCfgDto.RequestReplacementString,
		ePrefix.XCpy(
			"RequestReplacementString"))

	if err != nil {
		return err
	}

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
		"CharSearchTestConfigDto",
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
		2,
		"")

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"Final Text Output"))

	return err
}

// ptr - Returns a pointer to a new instance of
// charSearchTestConfigDtoNanobot.
func (searchTestConfigNanobot charSearchTestConfigDtoNanobot) ptr() *charSearchTestConfigDtoNanobot {

	if searchTestConfigNanobot.lock == nil {
		searchTestConfigNanobot.lock = new(sync.Mutex)
	}

	searchTestConfigNanobot.lock.Lock()

	defer searchTestConfigNanobot.lock.Unlock()

	return &charSearchTestConfigDtoNanobot{
		lock: new(sync.Mutex),
	}
}

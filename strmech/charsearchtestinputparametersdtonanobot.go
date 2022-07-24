package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// charSearchTestInputParametersDtoNanobot - Provides helper methods for type
// CharSearchTestInputParametersDto.
//
type charSearchTestInputParametersDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceTestInputParms' to input parameter
// 'destinationTestInputParms'. Both instances are of type
// CharSearchTestInputParametersDto.
//
// IMPORTANT
//
// ----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationTestInputParms'
// will be overwritten.
//
// Also, NO data validation is performed on 'sourceTestInputParms'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  destinationTestInputParms     *CharSearchTestInputParametersDto
//     - A pointer to a CharSearchTestInputParametersDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceTestInputParms'.
//
//       'destinationTestInputParms' is the destination for this
//       copy operation.
//
//
//  sourceTestInputParms          *CharSearchTestInputParametersDto
//     - A pointer to another CharSearchTestInputParametersDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationTestInputParms'.
//
//       'sourceTestInputParms' is the source for this copy
//       operation.
//
//       No data validation is performed on 'sourceTestInputParms'.
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
func (searchTestInputParmsNanobot *charSearchTestInputParametersDtoNanobot) copyIn(
	destinationTestInputParms *CharSearchTestInputParametersDto,
	sourceTestInputParms *CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

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

	if destinationTestInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'destinationTestInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	if sourceTestInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'sourceTestInputParms' is a nil pointer!\n",
			ePrefix.String())

		return err
	}

	charSearchTestInputParametersDtoAtom{}.ptr().empty(
		destinationTestInputParms)

	if sourceTestInputParms.TestString != nil {

		destinationTestInputParms.TestString =
			&RuneArrayDto{}

		err = destinationTestInputParms.TestString.CopyIn(
			sourceTestInputParms.TestString,
			ePrefix.XCpy("destinationTestInputParms<-"+
				"sourceTestInputParms"))

		if err != nil {
			return err
		}
	}

	destinationTestInputParms.TestInputParametersName =
		sourceTestInputParms.TestInputParametersName

	destinationTestInputParms.TestStringName =
		sourceTestInputParms.TestStringName

	destinationTestInputParms.TestStringLength =
		sourceTestInputParms.TestStringLength

	destinationTestInputParms.TestStringLengthName =
		sourceTestInputParms.TestStringLengthName

	destinationTestInputParms.TestStringStartingIndex =
		sourceTestInputParms.TestStringStartingIndex

	destinationTestInputParms.TestStringStartingIndexName =
		sourceTestInputParms.TestStringStartingIndexName

	destinationTestInputParms.TestStringDescription1 =
		sourceTestInputParms.TestStringDescription1

	destinationTestInputParms.TestStringDescription2 =
		sourceTestInputParms.TestStringDescription2

	destinationTestInputParms.CollectionTestObjIndex =
		sourceTestInputParms.CollectionTestObjIndex

	destinationTestInputParms.NumValueType =
		sourceTestInputParms.NumValueType

	destinationTestInputParms.NumStrFormatType =
		sourceTestInputParms.NumStrFormatType

	destinationTestInputParms.NumSymbolLocation =
		sourceTestInputParms.NumSymbolLocation

	destinationTestInputParms.NumSymbolClass =
		sourceTestInputParms.NumSymbolClass

	destinationTestInputParms.NumSignValue =
		sourceTestInputParms.NumSignValue

	destinationTestInputParms.PrimaryNumSignPosition =
		sourceTestInputParms.PrimaryNumSignPosition

	destinationTestInputParms.SecondaryNumSignPosition =
		sourceTestInputParms.SecondaryNumSignPosition

	destinationTestInputParms.TextCharSearchType =
		sourceTestInputParms.TextCharSearchType

	destinationTestInputParms.RequestFoundTestCharacters =
		sourceTestInputParms.RequestFoundTestCharacters

	destinationTestInputParms.RequestRemainderString =
		sourceTestInputParms.RequestRemainderString

	destinationTestInputParms.RequestReplacementString =
		sourceTestInputParms.RequestReplacementString

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'testInputParms', a pointer to an instance of
// CharSearchTestInputParametersDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
// NO validation is performed on 'testInputParms'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  testInputParms                *CharSearchTestInputParametersDto
//     - A pointer to an instance of CharSearchTestInputParametersDto. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of CharSearchTestInputParametersDto.
//
//       No data validation is performed on 'testInputParms'.
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
//  deepCopyTestInputParms        CharSearchTestInputParametersDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'testInputParms' will be created and
//       returned in a new instance of
//       CharSearchTestInputParametersDto.
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
func (searchTestInputParmsNanobot *charSearchTestInputParametersDtoNanobot) copyOut(
	testInputParms *CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	deepCopyTestInputParms CharSearchTestInputParametersDto,
	err error) {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	charSearchTestInputParametersDtoAtom{}.ptr().empty(
		&deepCopyTestInputParms)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTestInputParametersDtoNanobot."+
			"copyOut()",
		"")

	if err != nil {

		return deepCopyTestInputParms, err

	}

	if testInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'testInputParms' is a nil pointer!\n",
			ePrefix.String())

		return deepCopyTestInputParms, err
	}

	if testInputParms.TestString != nil {

		deepCopyTestInputParms.TestString =
			&RuneArrayDto{}

		err = deepCopyTestInputParms.TestString.CopyIn(
			testInputParms.TestString,
			ePrefix.XCpy("deepCopyTestInputParms<-"+
				"testInputParms"))

		if err != nil {
			return deepCopyTestInputParms, err
		}
	}

	deepCopyTestInputParms.TestInputParametersName =
		testInputParms.TestInputParametersName

	deepCopyTestInputParms.TestStringName =
		testInputParms.TestStringName

	deepCopyTestInputParms.TestStringLength =
		testInputParms.TestStringLength

	deepCopyTestInputParms.TestStringLengthName =
		testInputParms.TestStringLengthName

	deepCopyTestInputParms.TestStringStartingIndex =
		testInputParms.TestStringStartingIndex

	deepCopyTestInputParms.TestStringStartingIndexName =
		testInputParms.TestStringStartingIndexName

	deepCopyTestInputParms.TestStringDescription1 =
		testInputParms.TestStringDescription1

	deepCopyTestInputParms.TestStringDescription2 =
		testInputParms.TestStringDescription2

	deepCopyTestInputParms.CollectionTestObjIndex =
		testInputParms.CollectionTestObjIndex

	deepCopyTestInputParms.NumValueType =
		testInputParms.NumValueType

	deepCopyTestInputParms.NumStrFormatType =
		testInputParms.NumStrFormatType

	deepCopyTestInputParms.NumSymbolLocation =
		testInputParms.NumSymbolLocation

	deepCopyTestInputParms.NumSymbolClass =
		testInputParms.NumSymbolClass

	deepCopyTestInputParms.NumSignValue =
		testInputParms.NumSignValue

	deepCopyTestInputParms.PrimaryNumSignPosition =
		testInputParms.PrimaryNumSignPosition

	deepCopyTestInputParms.SecondaryNumSignPosition =
		testInputParms.SecondaryNumSignPosition

	deepCopyTestInputParms.TextCharSearchType =
		testInputParms.TextCharSearchType

	deepCopyTestInputParms.RequestFoundTestCharacters =
		testInputParms.RequestFoundTestCharacters

	deepCopyTestInputParms.RequestRemainderString =
		testInputParms.RequestRemainderString

	deepCopyTestInputParms.RequestReplacementString =
		testInputParms.RequestReplacementString

	return deepCopyTestInputParms, err
}

// getParameterTextListing - Returns formatted text output
// detailing the member variable names and corresponding values
// contained in the 'testInputParms' instance of CharSearchTestInputParametersDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  testInputParms           *CharSearchTestInputParametersDto
//     - A pointer to an instance of
//       CharSearchTestInputParametersDto. Formatted text output
//       will be generated listing the member variable names and
//       their corresponding values. The formatted text can then
//       be used for text displays, file output or printing.
//
//       No data validation is performed on this instance of
//       CharSearchTestInputParametersDto.
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
//       'testInputParms' . This formatted text can them be used
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
func (searchTestInputParmsNanobot *charSearchTestInputParametersDtoNanobot) getParameterTextListing(
	testInputParms *CharSearchTestInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

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

	if testInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'testInputParms' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	// Total available Length of Output Line
	const maxLineLen = 78

	// Max Label Field Length = 24
	const maxLabelFieldLen = 24

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
	txtFormatCol.AddLineSolidDto(solidLineDto)

	// Title Line 1
	err = txtFormatCol.AddLine1Col(
		"CharSearchTestInputParametersDto",
		ePrefix.XCpy(
			"Top-Title Line 1"))

	if err != nil {
		return strBuilder, err
	}

	txtStrParam :=
		testInputParms.TestInputParametersName

	if len(txtStrParam) > 0 {

		// Title Line 2
		err = txtFormatCol.AddLine1Col(
			txtStrParam,
			ePrefix.XCpy(
				"Top-Title Line 2"))

		if err != nil {
			return strBuilder, err
		}

	}

	// Title Line 3
	err = txtFormatCol.AddLine1Col(
		"Parameter Listing",
		ePrefix.XCpy(
			"Top-Title Line 3"))

	if err != nil {
		return strBuilder, err
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
		return strBuilder, err
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
		return strBuilder, err
	}

	// Build TestInputParametersName

	txtStrLabel := "TestInputParametersName"

	txtStrParam =
		testInputParms.TestInputParametersName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestInputParametersName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestInputParametersName"))

	if err != nil {
		return strBuilder, err
	}

	// Build Formatted Test String

	txtStrLabel = "TestString"

	if testInputParms.TestString == nil {

		txtStrParam =
			"Not Set. TestString is a nil pointer"

	} else {

		txtStrParam =
			testInputParms.TestString.GetCharacterString()

		if len(txtStrParam) == 0 {

			txtStrParam = "TestString is NOT NIL and Empty."

		}

	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestString"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringName

	txtStrLabel = "TestStringName"

	txtStrParam =
		testInputParms.TestStringName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringName"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringLength

	txtStrLabel = "TestStringLength"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.TestStringLength,
		ePrefix.XCpy(
			"TestStringLength"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringLengthName

	txtStrLabel = "TestStringLengthName"

	txtStrParam =
		testInputParms.TestStringLengthName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringLengthName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringLengthName"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringStartingIndex

	txtStrLabel = "TestStringStartingIndex"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.TestStringStartingIndex,
		ePrefix.XCpy(
			"TestStringStartingIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringStartingIndexName

	txtStrLabel = "TestStringStartingIndexName"

	txtStrParam =
		testInputParms.TestStringStartingIndexName

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringStartingIndexName is EMPTY!"
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		txtStrParam,
		ePrefix.XCpy(
			"TestStringStartingIndexName"))

	if err != nil {
		return strBuilder, err
	}

	// Build TestStringDescription1

	txtStrLabel = "TestStringDescription1"

	txtStrParam =
		testInputParms.TestStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringDescription1 is EMPTY!"
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
		testInputParms.TestStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam = "TestStringDescription2 is EMPTY!"
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
		testInputParms.CollectionTestObjIndex,
		ePrefix.XCpy(
			"CollectionTestObjIndex"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumValueType

	txtStrLabel = "NumValueType"

	if !testInputParms.NumValueType.XIsValid() {
		testInputParms.NumValueType = NumValType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.NumValueType,
		ePrefix.XCpy(
			"NumValueType"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumStrFormatType

	txtStrLabel = "NumStrFormatType"

	if !testInputParms.NumStrFormatType.XIsValid() {
		testInputParms.NumStrFormatType = NumStrFmtType.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.NumStrFormatType,
		ePrefix.XCpy(
			"NumStrFormatType"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumSymbolLocation

	txtStrLabel = "NumSymbolLocation"

	if !testInputParms.NumSymbolLocation.XIsValid() {
		testInputParms.NumSymbolLocation = NumSymLocation.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.NumSymbolLocation,
		ePrefix.XCpy(
			"NumSymbolLocation"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumSymbolClass

	txtStrLabel = "NumSymbolClass"

	if !testInputParms.NumSymbolClass.XIsValid() {
		testInputParms.NumSymbolClass = NumSymClass.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.NumSymbolClass,
		ePrefix.XCpy(
			"NumSymbolClass"))

	if err != nil {
		return strBuilder, err
	}

	// Build NumSignValue

	txtStrLabel = "NumSignValue"

	if !testInputParms.NumSignValue.XIsValid() {
		testInputParms.NumSignValue = NumSignVal.None()
	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.NumSignValue,
		ePrefix.XCpy(
			"NumSignValue"))

	if err != nil {
		return strBuilder, err
	}

	// Build PrimaryNumSignPosition

	txtStrLabel = "PrimaryNumSignPosition"

	if !testInputParms.PrimaryNumSignPosition.XIsValid() {

		testInputParms.PrimaryNumSignPosition =
			NumSignSymPos.None()

	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.PrimaryNumSignPosition,
		ePrefix.XCpy(
			"PrimaryNumSignPosition"))

	if err != nil {
		return strBuilder, err
	}

	// Build SecondaryNumSignPosition

	txtStrLabel = "SecondaryNumSignPosition"

	if !testInputParms.SecondaryNumSignPosition.XIsValid() {

		testInputParms.SecondaryNumSignPosition =
			NumSignSymPos.None()

	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.SecondaryNumSignPosition,
		ePrefix.XCpy(
			"SecondaryNumSignPosition"))

	if err != nil {
		return strBuilder, err
	}

	// Build TextCharSearchType

	txtStrLabel = "TextCharSearchType"

	if !testInputParms.TextCharSearchType.XIsValid() {

		testInputParms.TextCharSearchType =
			CharSearchType.None()

	}

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.TextCharSearchType,
		ePrefix.XCpy(
			"TextCharSearchType"))

	if err != nil {
		return strBuilder, err
	}

	// Build RequestFoundTestCharacters

	txtStrLabel = "RequestFoundTestCharacters"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.RequestFoundTestCharacters,
		ePrefix.XCpy(
			"RequestFoundTestCharacters"))

	if err != nil {
		return strBuilder, err
	}

	// Build RequestRemainderString

	txtStrLabel = "RequestRemainderString"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.RequestRemainderString,
		ePrefix.XCpy(
			"RequestRemainderString"))

	if err != nil {
		return strBuilder, err
	}

	// Build RequestReplacementString

	txtStrLabel = "RequestReplacementString"

	err = txtFormatCol.AddLine2Col(
		txtStrLabel,
		testInputParms.RequestReplacementString,
		ePrefix.XCpy(
			"RequestReplacementString"))

	if err != nil {
		return strBuilder, err
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
		"CharSearchTestInputParametersDto",
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
	txtFormatCol.AddLineSolidDto(solidLineDto)

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
// charSearchTestInputParametersDtoNanobot.
//
func (searchTestInputParmsNanobot charSearchTestInputParametersDtoNanobot) ptr() *charSearchTestInputParametersDtoNanobot {

	if searchTestInputParmsNanobot.lock == nil {
		searchTestInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTestInputParmsNanobot.lock.Lock()

	defer searchTestInputParmsNanobot.lock.Unlock()

	return &charSearchTestInputParametersDtoNanobot{
		lock: new(sync.Mutex),
	}
}

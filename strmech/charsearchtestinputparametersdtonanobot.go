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

	destinationTestInputParms.NumSymLocation =
		sourceTestInputParms.NumSymLocation

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
			"copyIn()",
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

	deepCopyTestInputParms.NumSymLocation =
		testInputParms.NumSymLocation

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

	return deepCopyTestInputParms, err
}

// getFormattedText - Returns formatted text output detailing the
// member variable values contained in the 'testInputParms'
// instance of CharSearchTestInputParametersDto.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  testInputParms           *CharSearchTestInputParametersDto
//     - A pointer to an instance of
//       CharSearchTestInputParametersDto instance. Formatted
//       text output will be generated listing the member variable
//       names and their corresponding values. The formatted text
//       can then be used for text displays, file output or
//       printing.
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
func (searchTestInputParmsNanobot *charSearchTestInputParametersDtoNanobot) getFormattedText(
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
			"getFormattedText()",
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

	const maxLineLen = 78

	// Total available Length of Output Line
	const maxFieldLen = 70

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
	txtFmt.Label.FieldText = "CharSearchTestInputParametersDto"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	txtStrParam :=
		testInputParms.TestInputParametersName

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

	err = txtBuilder.BuildTextFormatters(
		&strBuilder,
		fmtrs,
		ePrefix.XCpy(
			"strBuilder<-Marquee Top"))

	if err != nil {

		return strBuilder, err
	}

	fmtrs = nil

	// End Of Marquee

	// Begin Label Parameter Pairs

	colonSpace := ": "

	var labelParams []TextLabelValueStrings

	// TestInputParametersName
	labelParam := TextLabelValueStrings{}

	labelParam.ParamLabel = "TestInputParametersName"

	labelParam.ParamValue =
		testInputParms.TestInputParametersName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "TestInputParametersName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build Formatted Test String

	if testInputParms.TestString == nil {
		labelParam = TextLabelValueStrings{}

		labelParam.ParamLabel = "TestString"

		labelParam.ParamValue =
			"Not Set. TestString is a nil pointer"

		labelParams = append(labelParams, labelParam)

	} else if testInputParms.TestString.
		GetRuneArrayLength() <=
		(maxLineLen - maxLabelFieldLen - 3) {

		labelParam = TextLabelValueStrings{}

		labelParam.ParamLabel = "TestString"

		labelParam.ParamValue =
			testInputParms.TestString.
				GetCharacterString()

		if len(labelParam.ParamValue) == 0 {
			labelParam.ParamValue = "TestString is Empty. Length==Zero."
		}

		labelParams = append(labelParams, labelParam)

	} else {
		// Must be
		//  testInputParms.TestString.
		//  	GetRuneArrayLength() >
		//  	(maxLineLen - maxLabelFieldLen -3)

		// Write existing Parameter Label-Value Pairs
		// to string (strBuilder)
		err = txtBuilder.BuildLabelsValues(
			&strBuilder,
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
				"labelParams #1"))

		if err != nil {

			return strBuilder, err
		}

		labelParams = nil

		err = txtBuilder.FieldsSingleLabel(
			&strBuilder,
			" ",
			"TestString",
			maxLabelFieldLen,
			TxtJustify.Right(),
			colonSpace,
			"\n",
			ePrefix.XCpy(
				"TestString Label"))

		if err != nil {

			return strBuilder, err
		}

		err = txtBuilder.FieldsSingleLabel(
			&strBuilder,
			"  ",
			testInputParms.TestString.GetCharacterString(),
			-1,
			TxtJustify.Left(),
			"",
			"\n",
			ePrefix.XCpy(
				"TestString Param Value"))

		if err != nil {

			return strBuilder, err
		}

		err = txtBuilder.LineBlank(
			&strBuilder,
			1,
			ePrefix.XCpy(
				"Blank Line After ReplacementString"))

		if err != nil {

			return strBuilder, err
		}

	}

	// TestStringName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringName"

	labelParam.ParamValue =
		testInputParms.TestStringName

	labelParams = append(labelParams, labelParam)

	// TestStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		testInputParms.TestStringLength)

	labelParams = append(labelParams, labelParam)

	// TestStringLengthName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLengthName"

	labelParam.ParamValue =
		testInputParms.TestStringLengthName

	labelParams = append(labelParams, labelParam)

	// TestStringStartingIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		testInputParms.TestStringStartingIndex)

	labelParams = append(labelParams, labelParam)

	// TestStringStartingIndexName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndexName"

	labelParam.ParamValue =
		testInputParms.TestStringStartingIndexName

	labelParams = append(labelParams, labelParam)

	// TestStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription1"

	labelParam.ParamValue =
		testInputParms.TestStringDescription1

	labelParams = append(labelParams, labelParam)

	// TestStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription2"

	labelParam.ParamValue =
		testInputParms.TestStringDescription2

	labelParams = append(labelParams, labelParam)

	// CollectionTestObjIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "CollectionTestObjIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		testInputParms.CollectionTestObjIndex)

	labelParams = append(labelParams, labelParam)

	// NumValueType

	labelParam = TextLabelValueStrings{}

	if !testInputParms.NumValueType.XIsValid() {
		testInputParms.NumValueType = NumValType.None()
	}

	labelParam.ParamLabel = "NumValueType"

	labelParam.ParamValue =
		testInputParms.NumValueType.String()

	labelParams = append(labelParams, labelParam)

	// NumStrFormatType

	labelParam = TextLabelValueStrings{}

	if !testInputParms.NumStrFormatType.XIsValid() {
		testInputParms.NumStrFormatType = NumStrFmtType.None()
	}

	labelParam.ParamLabel = "NumStrFormatType"

	labelParam.ParamValue =
		testInputParms.NumStrFormatType.String()

	labelParams = append(labelParams, labelParam)

	// NumSymLocation

	labelParam = TextLabelValueStrings{}

	if !testInputParms.NumSymLocation.XIsValid() {
		testInputParms.NumSymLocation = NumSymLocation.None()
	}

	labelParam.ParamLabel = "NumSymLocation"

	labelParam.ParamValue =
		testInputParms.NumSymLocation.String()

	labelParams = append(labelParams, labelParam)

	// NumSymbolClass

	labelParam = TextLabelValueStrings{}

	if !testInputParms.NumSymbolClass.XIsValid() {
		testInputParms.NumSymbolClass = NumSymClass.None()
	}

	labelParam.ParamLabel = "NumSymbolClass"

	labelParam.ParamValue =
		testInputParms.NumSymbolClass.String()

	labelParams = append(labelParams, labelParam)

	// NumSignValue

	labelParam = TextLabelValueStrings{}

	if !testInputParms.NumSignValue.XIsValid() {
		testInputParms.NumSignValue = NumSignVal.None()
	}

	labelParam.ParamLabel = "NumSignValue"

	labelParam.ParamValue =
		testInputParms.NumSignValue.String()

	labelParams = append(labelParams, labelParam)

	// PrimaryNumSignPosition

	labelParam = TextLabelValueStrings{}

	if !testInputParms.PrimaryNumSignPosition.XIsValid() {

		testInputParms.PrimaryNumSignPosition =
			NumSignSymPos.None()

	}

	labelParam.ParamLabel = "PrimaryNumSignPosition"

	labelParam.ParamValue =
		testInputParms.PrimaryNumSignPosition.String()

	labelParams = append(labelParams, labelParam)

	// SecondaryNumSignPosition

	labelParam = TextLabelValueStrings{}

	if !testInputParms.SecondaryNumSignPosition.XIsValid() {

		testInputParms.SecondaryNumSignPosition =
			NumSignSymPos.None()

	}

	labelParam.ParamLabel = "SecondaryNumSignPosition"

	labelParam.ParamValue =
		testInputParms.SecondaryNumSignPosition.String()

	labelParams = append(labelParams, labelParam)

	// TextCharSearchType

	labelParam = TextLabelValueStrings{}

	if !testInputParms.TextCharSearchType.XIsValid() {

		testInputParms.TextCharSearchType =
			CharSearchType.None()

	}

	labelParam.ParamLabel = "TextCharSearchType"

	labelParam.ParamValue =
		testInputParms.TextCharSearchType.String()

	labelParams = append(labelParams, labelParam)

	err = txtBuilder.BuildLabelsValues(
		&strBuilder,
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
	txtFmt.Label.FieldText = "CharSearchTestInputParametersDto"
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

	err = txtBuilder.BuildTextFormatters(
		&strBuilder,
		fmtrs,
		ePrefix.XCpy(
			"Marquee-Bottom"))
	fmtrs = nil

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

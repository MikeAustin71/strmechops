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
//
type charSearchTargetInputParametersDtoNanobot struct {
	lock *sync.Mutex
}

// copyIn - Copies all data from input parameter
// 'sourceTargetInputParms' to input parameter
// 'destinationTargetInputParms'. Both instances are of type
// CharSearchTargetInputParametersDto.
//
// IMPORTANT
//
// -----------------------------------------------------------------
//
// Be advised that the data fields in 'destinationTargetInputParms'
// will be overwritten.
//
// Also, NO validation is performed on 'sourceTargetInputParms'.
//
//
// -----------------------------------------------------------------
//
// Input Parameters
//
//  destinationTargetInputParms     *CharSearchTargetInputParametersDto
//     - A pointer to a CharSearchTargetInputParametersDto instance.
//       All the member variable data fields in this object will be
//       replaced by data values copied from input parameter
//       'sourceTargetInputParms'.
//
//       'destinationTargetInputParms' is the destination for this
//       copy operation.
//
//
//  sourceTargetInputParms          *CharSearchTargetInputParametersDto
//     - A pointer to another CharSearchTargetInputParametersDto
//       instance. All the member variable data values from this
//       object will be copied to corresponding member variables in
//       'destinationTargetInputParms'.
//
//       'sourceTargetInputParms' is the source for this copy
//       operation.
//
//       No data validation is performed on
//       'sourceTargetInputParms'.
//
//
//  errPrefDto          *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
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

	destinationTargetInputParms.TextCharSearchType =
		sourceTargetInputParms.TextCharSearchType

	return err
}

// copyOut - Returns a deep copy of the input parameter
// 'targetInputParms', a pointer to an instance of
// CharSearchTargetInputParametersDto.
//
// ----------------------------------------------------------------
//
// IMPORTANT
//
//
// NO validation is performed on 'targetInputParms'.
//
// ------------------------------------------------------------------------
//
// Input Parameters
//
//  targetInputParms                *CharSearchTargetInputParametersDto
//     - A pointer to an instance of CharSearchTargetInputParametersDto. A
//       deep copy of the internal member variables will be created
//       and returned in a new instance of CharSearchTargetInputParametersDto.
//
//       No data validation is performed on 'targetInputParms'.
//
//
//  errPrefDto                      *ePref.ErrPrefixDto
//     - This object encapsulates an error prefix string which is
//       included in all returned error messages. Usually, it
//       contains the name of the calling method or methods listed
//       as a function chain.
//
//       If no error prefix information is needed, set this parameter
//       to 'nil'.
//
//       Type ErrPrefixDto is included in the 'errpref' software
//       package, "github.com/MikeAustin71/errpref".
//
//
// ------------------------------------------------------------------------
//
// Return Values
//
//  deepCopyTargetInputParms        CharSearchTargetInputParametersDto
//     - If this method completes successfully, a deep copy of
//       input parameter 'targetInputParms' will be created and
//       returned in a new instance of CharSearchTargetInputParametersDto.
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

	deepCopyTargetInputParms.TextCharSearchType =
		targetInputParms.TextCharSearchType

	return deepCopyTargetInputParms, err
}

// getFormattedText - Returns formatted text output detailing the
// member variable values contained in the 'targetInputParms'
// instance of CharSearchTargetInputParametersDto.
//
func (searchTargetInputParmsNanobot charSearchTargetInputParametersDtoNanobot) getFormattedText(
	targetInputParms *CharSearchTargetInputParametersDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strings.Builder,
	error) {

	if searchTargetInputParmsNanobot.lock == nil {
		searchTargetInputParmsNanobot.lock = new(sync.Mutex)
	}

	searchTargetInputParmsNanobot.lock.Lock()

	defer searchTargetInputParmsNanobot.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	strBuilder := strings.Builder{}

	strBuilder.Grow(1024)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"charSearchTargetInputParametersDtoNanobot."+
			"getFormattedText()",
		"")

	if err != nil {

		return strBuilder, err
	}

	if targetInputParms == nil {

		err = fmt.Errorf("%v\n"+
			"ERROR: Input parameter 'targetInputParms' is a nil pointer!\n",
			ePrefix.String())

		return strBuilder, err
	}

	const maxLineLen = 78

	// Total available Length of Output Line
	const maxFieldLen = 70

	// Max Label Field Length = 36
	const maxLabelFieldLen = 36

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
	txtFmt.Label.FieldText = "CharSearchTargetInputParametersDto"
	txtFmt.Label.FieldLength = maxFieldLen
	txtFmt.Label.FieldJustify = TxtJustify.Center()
	txtFmt.Label.RightMarginStr = ""
	txtFmt.Label.LineTerminator = "\n"
	fmtrs = append(fmtrs, txtFmt)

	txtStrParam :=
		targetInputParms.TargetInputParametersName

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

	// Title Line  3 Date/Time
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

	// Label Parameter Pairs

	colonSpace := ": "

	var labelParams []TextLabelValueStrings

	// TargetInputParametersName
	labelParam := TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetInputParametersName"

	labelParam.ParamValue =
		targetInputParms.TargetInputParametersName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SearchResultsName is EMPTY!"
	}

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
			"labelParams #1"))

	if err != nil {

		return strBuilder, err
	}

	labelParams = nil

	// Build Formatted Target String

	if targetInputParms.TargetString == nil {
		labelParam = TextLabelValueStrings{}

		labelParam.ParamLabel = "TargetString"

		labelParam.ParamValue =
			"Not Set. TargetString is a nil pointer"

		labelParams = append(labelParams, labelParam)

	} else if targetInputParms.TargetString.
		GetRuneArrayLength() == 0 {

		labelParam = TextLabelValueStrings{}

		labelParam.ParamLabel = "TargetString"

		labelParam.ParamValue =
			"TargetString is empty - zero characters"

		labelParams = append(labelParams, labelParam)

	} else if targetInputParms.TargetString.
		GetRuneArrayLength() <=
		(maxLineLen - maxLabelFieldLen - 3) {

		labelParam = TextLabelValueStrings{}

		labelParam.ParamLabel = "TargetString"

		labelParam.ParamValue =
			targetInputParms.TargetString.
				GetCharacterString()

		labelParams = append(labelParams, labelParam)

	} else {
		// Must be
		//  targetInputParms.TargetString.
		//  	GetRuneArrayLength() >
		//  	(maxLineLen - maxLabelFieldLen -3)

		err = txtBuilder.FieldsSingleLabel(
			&strBuilder,
			" ",
			"TargetString",
			maxLabelFieldLen,
			TxtJustify.Right(),
			colonSpace,
			"\n",
			ePrefix.XCpy(
				"TargetString Label"))

		if err != nil {

			return strBuilder, err
		}

		err = txtBuilder.FieldsSingleLabel(
			&strBuilder,
			"  ",
			targetInputParms.TargetString.GetCharacterString(),
			-1,
			TxtJustify.Left(),
			"",
			"\n",
			ePrefix.XCpy(
				"TargetString Param Value"))

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

	// TargetStringName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringName"

	labelParam.ParamValue =
		targetInputParms.TargetStringName

	labelParams = append(labelParams, labelParam)

	// TargetStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundSearchTarget"

	labelParam.ParamValue = fmt.Sprintf("%v",
		targetInputParms.TargetStringLength)

	labelParams = append(labelParams, labelParam)

	// TargetStringLengthName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLengthName"

	labelParam.ParamValue =
		targetInputParms.TargetStringLengthName

	labelParams = append(labelParams, labelParam)

	// TargetStringStartingSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringStartingSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		targetInputParms.TargetStringStartingSearchIndex)

	labelParams = append(labelParams, labelParam)

	// TargetStringStartingSearchIndexName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringStartingSearchIndexName"

	labelParam.ParamValue =
		targetInputParms.TargetStringStartingSearchIndexName

	labelParams = append(labelParams, labelParam)

	// TargetStringSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		targetInputParms.TargetStringSearchLength)

	labelParams = append(labelParams, labelParam)

	// TargetStringSearchLengthName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringSearchLengthName"

	labelParam.ParamValue =
		targetInputParms.TargetStringSearchLengthName

	labelParams = append(labelParams, labelParam)

	// TargetStringAdjustedSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringAdjustedSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		targetInputParms.TargetStringAdjustedSearchLength)

	labelParams = append(labelParams, labelParam)

	// TargetStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription1"

	labelParam.ParamValue =
		targetInputParms.TargetStringDescription1

	labelParams = append(labelParams, labelParam)

	// TargetStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription2"

	labelParam.ParamValue =
		targetInputParms.TargetStringDescription2

	labelParams = append(labelParams, labelParam)

	// FoundFirstNumericDigitInNumStr

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundFirstNumericDigitInNumStr"

	labelParam.ParamValue = fmt.Sprintf("%v",
		targetInputParms.FoundFirstNumericDigitInNumStr)

	labelParams = append(labelParams, labelParam)

	// TextCharSearchType

	labelParam = TextLabelValueStrings{}

	if !targetInputParms.TextCharSearchType.XIsValid() {
		targetInputParms.TextCharSearchType = CharSearchType.None()
	}

	labelParam.ParamLabel = "TextCharSearchType"

	labelParam.ParamValue =
		targetInputParms.TextCharSearchType.String()

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
	txtFmt.Label.FieldText = "CharSearchTargetInputParametersDto"
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
// charSearchTargetInputParametersDtoNanobot.
//
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

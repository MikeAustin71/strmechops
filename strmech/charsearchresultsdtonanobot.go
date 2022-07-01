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

	targetSearchResultsDto.FoundSearchTarget =
		incomingSearchResultsDto.FoundSearchTarget

	targetSearchResultsDto.FoundSearchTargetOnPreviousSearch =
		incomingSearchResultsDto.FoundSearchTargetOnPreviousSearch

	targetSearchResultsDto.FoundFirstNumericDigitInNumStr =
		incomingSearchResultsDto.FoundFirstNumericDigitInNumStr

	targetSearchResultsDto.TargetInputParametersName =
		incomingSearchResultsDto.TargetInputParametersName

	targetSearchResultsDto.TargetStringLength =
		incomingSearchResultsDto.TargetStringLength

	targetSearchResultsDto.TargetStringSearchLength =
		incomingSearchResultsDto.TargetStringSearchLength

	targetSearchResultsDto.TargetStringStartingSearchIndex =
		incomingSearchResultsDto.TargetStringStartingSearchIndex

	targetSearchResultsDto.TargetStringFirstFoundIndex =
		incomingSearchResultsDto.TargetStringFirstFoundIndex

	targetSearchResultsDto.TargetStringLastFoundIndex =
		incomingSearchResultsDto.TargetStringLastFoundIndex

	targetSearchResultsDto.TargetStringLastSearchIndex =
		incomingSearchResultsDto.TargetStringLastSearchIndex

	targetSearchResultsDto.TargetStringDescription1 =
		incomingSearchResultsDto.TargetStringDescription1

	targetSearchResultsDto.TargetStringDescription2 =
		incomingSearchResultsDto.TargetStringDescription2

	targetSearchResultsDto.TestInputParametersName =
		incomingSearchResultsDto.TestInputParametersName

	targetSearchResultsDto.TestStringLength =
		incomingSearchResultsDto.TestStringLength

	targetSearchResultsDto.TestStringStartingIndex =
		incomingSearchResultsDto.TestStringStartingIndex

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

	if incomingSearchResultsDto.ReplacementString != nil {

		targetSearchResultsDto.ReplacementString =
			&RuneArrayDto{}

		err = targetSearchResultsDto.ReplacementString.CopyIn(
			incomingSearchResultsDto.ReplacementString,
			ePrefix.XCpy(
				"targetSearchResultsDto.ReplacementString"+
					"<-incomingSearchResultsDto.ReplacementString"))
	}

	targetSearchResultsDto.NumValueType =
		incomingSearchResultsDto.NumValueType

	targetSearchResultsDto.NumStrFormatType =
		incomingSearchResultsDto.NumStrFormatType

	targetSearchResultsDto.NumSymLocation =
		incomingSearchResultsDto.NumSymLocation

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

	copySearchResultsDto.FoundSearchTarget =
		searchResultsDto.FoundSearchTarget

	copySearchResultsDto.FoundSearchTargetOnPreviousSearch =
		searchResultsDto.FoundSearchTargetOnPreviousSearch

	copySearchResultsDto.FoundFirstNumericDigitInNumStr =
		searchResultsDto.FoundFirstNumericDigitInNumStr

	copySearchResultsDto.TargetInputParametersName =
		searchResultsDto.TargetInputParametersName

	copySearchResultsDto.TargetStringLength =
		searchResultsDto.TargetStringLength

	copySearchResultsDto.TargetStringSearchLength =
		searchResultsDto.TargetStringSearchLength

	copySearchResultsDto.TargetStringStartingSearchIndex =
		searchResultsDto.TargetStringStartingSearchIndex

	copySearchResultsDto.TargetStringFirstFoundIndex =
		searchResultsDto.TargetStringFirstFoundIndex

	copySearchResultsDto.TargetStringLastFoundIndex =
		searchResultsDto.TargetStringLastFoundIndex

	copySearchResultsDto.TargetStringLastSearchIndex =
		searchResultsDto.TargetStringLastSearchIndex

	copySearchResultsDto.TargetStringDescription1 =
		searchResultsDto.TargetStringDescription1

	copySearchResultsDto.TargetStringDescription2 =
		searchResultsDto.TargetStringDescription2

	copySearchResultsDto.TestInputParametersName =
		searchResultsDto.TestInputParametersName

	copySearchResultsDto.TestStringLength =
		searchResultsDto.TestStringLength

	copySearchResultsDto.TestStringStartingIndex =
		searchResultsDto.TestStringStartingIndex

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

	copySearchResultsDto.NumValueType =
		searchResultsDto.NumValueType

	copySearchResultsDto.NumStrFormatType =
		searchResultsDto.NumStrFormatType

	copySearchResultsDto.NumSymLocation =
		searchResultsDto.NumSymLocation

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

// getFormattedText - Returns formatted text output detailing the
// member variable values contained in the 'searchResultsDto'
// instance of CharSearchResultsDto.
//
func (searchResultsDtoNanobot *charSearchResultsDtoNanobot) getFormattedText(
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
			"copyIn()",
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

	const maxLineLen = 78

	// Total available Length of Marquee Output Line
	const maxFieldLen = 70

	// Max Label Field Length = 33
	const maxLabelFieldLen = 33

	txtBuilder := TextStrBuilder{}

	// Leading Title Marquee

	err = txtBuilder.LineBlank(
		&strBuilder,
		1,
		ePrefix.XCpy(
			"1-Leading Blank Line"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.FieldsSingleFiller(
		&strBuilder,
		" ",
		"=",
		maxFieldLen-2,
		" ",
		"\n",
		ePrefix.XCpy(
			"Marquee-Top"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.FieldLabel(
		&strBuilder,
		"CharSearchResultsDto",
		maxFieldLen,
		TxtJustify.Center(),
		"\n",
		ePrefix.XCpy(
			"Title-1"))

	if err != nil {

		return strBuilder, err
	}

	if len(searchResultsDto.SearchResultsName) > 0 {

		err = txtBuilder.FieldLabel(
			&strBuilder,
			searchResultsDto.SearchResultsName,
			maxFieldLen,
			TxtJustify.Center(),
			"\n",
			ePrefix.XCpy(
				"Title-2"))

		if err != nil {

			return strBuilder, err
		}

	}

	err = txtBuilder.FieldDateTime(
		&strBuilder,
		time.Now(),
		maxFieldLen,
		"Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		TxtJustify.Center(),
		"\n",
		ePrefix.XCpy(
			"Title-3"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.FieldsSingleFiller(
		&strBuilder,
		" ",
		"=",
		maxFieldLen-2,
		" ",
		"\n",
		ePrefix.XCpy(
			"Marquee-Bottom"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.LineBlank(
		&strBuilder,
		1,
		ePrefix.XCpy(
			"1-Trailing Blank Line"))

	if err != nil {

		return strBuilder, err
	}

	colonSpace := ": "

	var labelParams []TextLabelValueStrings

	// Build SearchResultsName Parameter
	labelParam := TextLabelValueStrings{}

	labelParam.ParamLabel = "SearchResultsName"

	labelParam.ParamValue = searchResultsDto.SearchResultsName

	if len(labelParam.ParamValue) == 0 {
		labelParam.ParamValue = "SearchResultsName is EMPTY!"
	}

	labelParams = append(labelParams, labelParam)

	// Build FoundSearchTarget Parameter
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundSearchTarget"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.FoundSearchTarget)

	labelParams = append(labelParams, labelParam)

	// Build FoundSearchTargetOnPreviousSearch
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundSearchTargetOnPreviousSearch"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.FoundSearchTargetOnPreviousSearch)

	labelParams = append(labelParams, labelParam)

	// Build FoundFirstNumericDigitInNumStr
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "FoundFirstNumericDigitInNumStr"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.FoundFirstNumericDigitInNumStr)

	labelParams = append(labelParams, labelParam)

	// Build TargetInputParametersName
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetInputParametersName"

	labelParam.ParamValue =
		searchResultsDto.TargetInputParametersName

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TargetStringLength)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringSearchLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringSearchLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TargetStringSearchLength)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringStartingSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringStartingSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TargetStringStartingSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringFirstFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringFirstFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TargetStringFirstFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLastFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLastFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TargetStringLastFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringLastSearchIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringLastSearchIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TargetStringLastSearchIndex)

	labelParams = append(labelParams, labelParam)

	// Build TargetStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription1"

	labelParam.ParamValue =
		searchResultsDto.TargetStringDescription1

	labelParams = append(labelParams, labelParam)

	// Build TargetStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TargetStringDescription2"

	labelParam.ParamValue =
		searchResultsDto.TargetStringDescription2

	labelParams = append(labelParams, labelParam)

	// Build TestStringLength
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLength"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TestStringLength)

	labelParams = append(labelParams, labelParam)

	// Build TestStringStartingIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringStartingIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TestStringStartingIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringFirstFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringFirstFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TestStringFirstFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringLastFoundIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringLastFoundIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.TestStringLastFoundIndex)

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription1
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription1"

	labelParam.ParamValue =
		searchResultsDto.TestStringDescription1

	labelParams = append(labelParams, labelParam)

	// Build TestStringDescription2
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TestStringDescription2"

	labelParam.ParamValue =
		searchResultsDto.TestStringDescription2

	labelParams = append(labelParams, labelParam)

	// Build CollectionTestObjIndex
	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "CollectionTestObjIndex"

	labelParam.ParamValue = fmt.Sprintf("%v",
		searchResultsDto.CollectionTestObjIndex)

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

	if searchResultsDto.ReplacementString == nil {
		// Build ReplacementString nil message
		labelParam = TextLabelValueStrings{}

		labelParam.ParamLabel = "ReplacementString"

		labelParam.ParamValue =
			"Not Set. ReplacementString is a nil pointer"

		labelParams = append(labelParams, labelParam)

	} else if searchResultsDto.ReplacementString.
		GetRuneArrayLength() == 0 {

		labelParam = TextLabelValueStrings{}

		labelParam.ParamLabel = "ReplacementString"

		labelParam.ParamValue =
			"ReplacementString is empty - zero characters"

		labelParams = append(labelParams, labelParam)

	} else if searchResultsDto.ReplacementString.
		GetRuneArrayLength() <=
		(maxLineLen - maxLabelFieldLen - 3) {

		labelParam = TextLabelValueStrings{}

		labelParam.ParamLabel = "ReplacementString"

		labelParam.ParamValue =
			searchResultsDto.ReplacementString.
				GetCharacterString()

		labelParams = append(labelParams, labelParam)

	} else {
		// Must be
		//  searchResultsDto.ReplacementString.
		//  	GetRuneArrayLength() >
		//  	(maxLineLen - maxLabelFieldLen -3)

		err = txtBuilder.FieldsSingleLabel(
			&strBuilder,
			" ",
			"ReplacementString",
			maxLabelFieldLen,
			TxtJustify.Right(),
			colonSpace,
			"\n",
			ePrefix.XCpy(
				"ReplacementString"))

		if err != nil {

			return strBuilder, err
		}

		err = txtBuilder.FieldsSingleLabel(
			&strBuilder,
			"  ",
			searchResultsDto.ReplacementString.GetCharacterString(),
			-1,
			TxtJustify.Left(),
			"",
			"\n",
			ePrefix.XCpy(
				"ReplacementString"))

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

	// Build NumValueType

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumValueType"

	if !searchResultsDto.NumValueType.XIsValid() {
		searchResultsDto.NumValueType =
			NumValType.None()
	}

	labelParam.ParamValue =
		searchResultsDto.NumValueType.String()

	labelParams = append(labelParams, labelParam)

	// Build NumStrFormatType

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumStrFormatType"

	if !searchResultsDto.NumStrFormatType.XIsValid() {
		searchResultsDto.NumStrFormatType =
			NumStrFmtType.None()
	}

	labelParam.ParamValue =
		searchResultsDto.NumStrFormatType.String()

	labelParams = append(labelParams, labelParam)

	// Build NumSymLocation

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumSymLocation"

	if !searchResultsDto.NumSymLocation.XIsValid() {
		searchResultsDto.NumSymLocation =
			NumSymLocation.None()
	}

	labelParam.ParamValue =
		searchResultsDto.NumSymLocation.String()

	labelParams = append(labelParams, labelParam)

	// Build NumSymbolClass

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumSymbolClass"

	if !searchResultsDto.NumSymbolClass.XIsValid() {
		searchResultsDto.NumSymbolClass =
			NumSymClass.None()
	}

	labelParam.ParamValue =
		searchResultsDto.NumSymbolClass.String()

	labelParams = append(labelParams, labelParam)

	// Build NumSignValue

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "NumSignValue"

	if !searchResultsDto.NumSignValue.XIsValid() {
		searchResultsDto.NumSignValue =
			NumSignVal.None()
	}

	labelParam.ParamValue =
		searchResultsDto.NumSignValue.String()

	labelParams = append(labelParams, labelParam)

	// Build PrimaryNumSignPosition

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "PrimaryNumSignPosition"

	if !searchResultsDto.PrimaryNumSignPosition.XIsValid() {
		searchResultsDto.PrimaryNumSignPosition =
			NumSignSymPos.None()
	}

	labelParam.ParamValue =
		searchResultsDto.PrimaryNumSignPosition.String()

	labelParams = append(labelParams, labelParam)

	// Build SecondaryNumSignPosition

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "SecondaryNumSignPosition"

	if !searchResultsDto.SecondaryNumSignPosition.XIsValid() {
		searchResultsDto.SecondaryNumSignPosition =
			NumSignSymPos.None()
	}

	labelParam.ParamValue =
		searchResultsDto.SecondaryNumSignPosition.String()

	labelParams = append(labelParams, labelParam)

	// Build TextCharSearchType

	labelParam = TextLabelValueStrings{}

	labelParam.ParamLabel = "TextCharSearchType"

	if !searchResultsDto.TextCharSearchType.XIsValid() {
		searchResultsDto.TextCharSearchType =
			CharSearchType.None()
	}

	labelParam.ParamValue =
		searchResultsDto.TextCharSearchType.String()

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

	// Trailing Title Marquee

	err = txtBuilder.LineBlank(
		&strBuilder,
		1,
		ePrefix.XCpy(
			"1-Trailing Blank Line"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.FieldsSingleFiller(
		&strBuilder,
		" ",
		"=",
		maxFieldLen-2,
		" ",
		"\n",
		ePrefix.XCpy(
			"Trailing Marquee-Top"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.FieldLabel(
		&strBuilder,
		"CharSearchResultsDto",
		maxFieldLen,
		TxtJustify.Center(),
		"\n",
		ePrefix.XCpy(
			"Title-1"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.FieldLabel(
		&strBuilder,
		"End of Parameter Listing",
		maxFieldLen,
		TxtJustify.Center(),
		"\n",
		ePrefix.XCpy(
			"Title-2"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.FieldsSingleFiller(
		&strBuilder,
		" ",
		"=",
		maxFieldLen-2,
		" ",
		"\n",
		ePrefix.XCpy(
			"Trailing Marquee-Bottom"))

	if err != nil {

		return strBuilder, err
	}

	err = txtBuilder.LineBlank(
		&strBuilder,
		1,
		ePrefix.XCpy(
			"2-Trailing Blank Line"))

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
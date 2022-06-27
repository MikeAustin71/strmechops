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
//       If 'sourceTargetInputParms' is determined to be invalid,
//       an error will be returned.
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
//       If the member variable data values encapsulated by
//       'targetInputParms' are found to be invalid, this method will
//       return an error
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

		err = deepCopyTargetInputParms.TargetString.CopyIn(
			targetInputParms.TargetString,
			ePrefix.XCpy("deepCopyTargetInputParms<-"+
				"targetInputParms"))

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

	// Total available Length of Output Line
	const maxFieldLen = 70

	// Max Label Field Length = 36
	const maxLabelFieldLen = 36

	// Title Marquee

	txtStdLine := TextLineSpecStandardLine{}

	var txtFieldDtos []TextFieldDto

	var txtFieldDto TextFieldDto

	txtFieldDto = TextFieldDto{
		FieldLength: 1,
		FieldType:   TxtFieldType.BlankLine(),
	}

	txtFieldDtos = append(txtFieldDtos, txtFieldDto)

	// Filler Line '========='
	txtFieldDto = TextFieldDto{
		LeftMarginStr:  " ",
		FieldDateTime:  time.Time{},
		DateTimeFormat: "",
		FieldText:      "=",
		FieldLength:    maxFieldLen - 2,
		FieldJustify:   TxtJustify.Left(),
		FieldType:      TxtFieldType.Filler(),
		RightMarginStr: " ",
		LineTerminator: "\n",
	}

	txtFieldDtos = append(txtFieldDtos, txtFieldDto)

	// Title Line 1
	txtFieldDto = TextFieldDto{
		FieldText:      "CharSearchTargetInputParametersDto",
		FieldLength:    maxFieldLen,
		FieldJustify:   TxtJustify.Center(),
		FieldType:      TxtFieldType.Label(),
		LineTerminator: "\n",
	}

	txtFieldDtos = append(txtFieldDtos, txtFieldDto)

	txtStrParam :=
		targetInputParms.TargetInputParametersName

	if len(txtStrParam) > 0 {

		// Title Line 2
		txtFieldDto = TextFieldDto{
			FieldText:      txtStrParam,
			FieldLength:    maxFieldLen,
			FieldJustify:   TxtJustify.Center(),
			FieldType:      TxtFieldType.Label(),
			LineTerminator: "\n",
		}

		txtFieldDtos = append(txtFieldDtos, txtFieldDto)

	}

	// Title Line 3 Date/Time
	txtFieldDto = TextFieldDto{
		FieldDateTime:  time.Now(),
		DateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		FieldLength:    maxFieldLen,
		FieldJustify:   TxtJustify.Center(),
		FieldType:      TxtFieldType.DateTime(),
	}

	txtFieldDtos = append(txtFieldDtos, txtFieldDto)

	// Filler Line '========='
	txtFieldDto = TextFieldDto{
		LeftMarginStr:  " ",
		FieldDateTime:  time.Time{},
		DateTimeFormat: "",
		FieldText:      "=",
		FieldLength:    maxFieldLen - 2,
		FieldJustify:   TxtJustify.Left(),
		FieldType:      TxtFieldType.Filler(),
		RightMarginStr: " ",
		LineTerminator: "\n",
	}

	txtFieldDtos = append(txtFieldDtos, txtFieldDto)

	txtFieldDto = TextFieldDto{
		FieldLength: 1,
		FieldType:   TxtFieldType.BlankLine(),
	}

	txtFieldDtos = append(txtFieldDtos, txtFieldDto)

	err = txtStdLine.BuildTextFieldLines(
		&strBuilder,
		txtFieldDtos,
		ePrefix)

	if err != nil {

		return strBuilder, err
	}

	txtFieldDtos = make([]TextFieldDto, 0)

	// End Of Marquee

	colonSpace := ": "

	// Build Formatted Target String

	txtStrLabel := "Target String"

	lenTxtStrParam := 0

	if targetInputParms.TargetString != nil {

		if targetInputParms.TargetString.GetRuneArrayLength() > 0 {
			txtStrParam =
				targetInputParms.TargetString.GetCharacterString()
		} else {

			txtStrParam = "Target String is EMPTY!"

		}
	} else {
		// targetInputParms.TargetString == nil
		txtStrParam = "Target String has a nil pointer (Not Set)!"
	}

	lenTxtStrParam = len(txtStrParam)

	var txtLabelParamField TextLabelParameterValueFieldDto

	var txtLabelParamFields []TextLabelParameterValueFieldDto

	if lenTxtStrParam >= (maxFieldLen - maxLabelFieldLen - 3) {
		// We need two Lines of Text

		txtFieldDto = TextFieldDto{
			LeftMarginStr:  " ",
			FieldDateTime:  time.Time{},
			DateTimeFormat: "",
			FieldText:      txtStrLabel,
			FieldLength:    maxLabelFieldLen,
			FieldJustify:   TxtJustify.Right(),
			FieldType:      TxtFieldType.Label(),
			RightMarginStr: colonSpace,
			LineTerminator: "\n",
		}

		txtFieldDtos = append(txtFieldDtos, txtFieldDto)

		txtFieldDto = TextFieldDto{
			LeftMarginStr:  " ",
			FieldDateTime:  time.Time{},
			DateTimeFormat: "",
			FieldText:      txtStrParam,
			FieldLength:    -1,
			FieldJustify:   TxtJustify.Left(),
			FieldType:      TxtFieldType.Label(),
			RightMarginStr: "",
			LineTerminator: "\n",
		}

		txtFieldDtos = append(txtFieldDtos, txtFieldDto)

		err = txtStdLine.BuildTextFieldLines(
			&strBuilder,
			txtFieldDtos,
			ePrefix)

		if err != nil {

			return strBuilder, err
		}

		txtFieldDtos = make([]TextFieldDto, 0)

	} else {

		txtLabelParamField = TextLabelParameterValueFieldDto{
			LeftMarginStr:            " ",
			ParamLabelStr:            txtStrLabel,
			ParamLabelLength:         maxLabelFieldLen,
			ParamLabelJustify:        TxtJustify.Right(),
			ParamLabelRightMarginStr: colonSpace,
			ParamValueDateTime:       time.Time{},
			DateTimeFormat:           "",
			ParamValueStr:            txtStrParam,
			ParamValueLength:         -1,
			ParamValueJustify:        TxtJustify.Left(),
			ParamRightMarginStr:      "",
			LineTerminator:           "\n",
		}

		txtLabelParamFields = append(
			txtLabelParamFields,
			txtLabelParamField)

	}

	// TargetInputParametersName
	txtStrLabel = "TargetInputParametersName"

	txtStrParam = targetInputParms.TargetInputParametersName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetInputParametersName is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringName

	txtStrLabel = "TargetStringName"

	txtStrParam = targetInputParms.TargetStringName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringName is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringLength

	txtStrLabel = "TargetStringLength"

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.TargetStringLength)

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringLengthName

	txtStrLabel = "TargetStringLengthName"

	txtStrParam = targetInputParms.TargetStringLengthName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringLengthName is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringStartingSearchIndex
	txtStrLabel = "TargetStringStartingSearchIndex"

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.TargetStringStartingSearchIndex)

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringStartingSearchIndex is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringStartingSearchIndexName
	txtStrLabel = "TargetStringStartingSearchIndexName"

	txtStrParam =
		targetInputParms.TargetStringStartingSearchIndexName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringStartingSearchIndexName is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringSearchLength
	txtStrLabel = "TargetStringSearchLength"

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.TargetStringSearchLength)

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringSearchLength is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringSearchLengthName
	txtStrLabel = "TargetStringSearchLengthName"

	txtStrParam =
		targetInputParms.TargetStringSearchLengthName

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringSearchLengthName is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringAdjustedSearchLength

	txtStrLabel = "TargetStringAdjustedSearchLength"

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.TargetStringAdjustedSearchLength)

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringAdjustedSearchLength is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringDescription1
	txtStrLabel = "TargetStringDescription1"

	txtStrParam = targetInputParms.TargetStringDescription1

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringDescription1 is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TargetStringDescription2

	txtStrLabel = "TargetStringDescription2"

	txtStrParam = targetInputParms.TargetStringDescription2

	if len(txtStrParam) == 0 {
		txtStrParam = "TargetStringDescription2 is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// FoundFirstNumericDigitInNumStr
	txtStrLabel = "FoundFirstNumericDigitInNumStr"

	txtStrParam =
		fmt.Sprintf("%v",
			targetInputParms.FoundFirstNumericDigitInNumStr)

	if len(txtStrParam) == 0 {
		txtStrParam = "FoundFirstNumericDigitInNumStr is EMPTY!"
	}

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	// TextCharSearchType
	txtStrLabel = "TextCharSearchType"

	searchType := targetInputParms.TextCharSearchType

	if !searchType.XIsValid() {
		searchType = CharSearchType.None()
	}

	txtStrParam =
		fmt.Sprintf("%v",
			searchType.String())

	txtLabelParamField = TextLabelParameterValueFieldDto{
		LeftMarginStr:            " ",
		ParamLabelStr:            txtStrLabel,
		ParamLabelLength:         maxLabelFieldLen,
		ParamLabelJustify:        TxtJustify.Right(),
		ParamLabelRightMarginStr: colonSpace,
		ParamValueDateTime:       time.Time{},
		DateTimeFormat:           "",
		ParamValueStr:            txtStrParam,
		ParamValueLength:         -1,
		ParamValueJustify:        TxtJustify.Left(),
		ParamRightMarginStr:      "",
		LineTerminator:           "\n",
	}

	txtLabelParamFields = append(
		txtLabelParamFields,
		txtLabelParamField)

	err = txtStdLine.BuildTextLabelParameterLines(
		&strBuilder,
		txtLabelParamFields,
		ePrefix)

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

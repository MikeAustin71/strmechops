package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
)

type textStrBuilderParamsDto struct {
	strBuilder                 *strings.Builder
	leftMarginStr              string
	lenLeftMarginStr           int
	textStr                    string
	lenTextStr                 int
	rightMarginStr             string
	lenRightMarginStr          int
	turnLineTerminationOff     bool
	lineTerminatorStr          string
	lenLineTerminatorStr       int
	maxLineLength              int
	currentLineLength          int
	turnAutoLineLengthBreaksOn bool
	lastWriteWasLineTerminator bool
	sourceTag                  string
	sourceDtoTag               string
	errPrefDto                 *ePref.ErrPrefixDto
}

// textStrBuilderMolecule - Provides helper methods for type
// TextStrBuilder.
//
type textStrBuilderMolecule struct {
	lock *sync.Mutex
}

func (txtBuilderMolecule *textStrBuilderMolecule) buildFieldDateTimeWithDto(
	strBuilder *strings.Builder,
	dateTimeFieldDto TextFieldDateTimeDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderMolecule.lock == nil {
		txtBuilderMolecule.lock = new(sync.Mutex)
	}

	txtBuilderMolecule.lock.Lock()

	defer txtBuilderMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderMolecule."+
			"buildFieldDateTimeWithDto()",
		"")

	if err != nil {
		return err
	}

	if dateTimeFieldDto.FormatType !=
		TxtFieldType.DateTime() {

		err = fmt.Errorf("%v\n"+
			"Error: 'dateTimeFieldDto.FormatType' is invalid!\n"+
			"'dateTimeFieldDto.FormatType' should be set to \n"+
			"TxtFieldType.DateTime(). It is NOT!\n"+
			"'dateTimeFieldDto.FormatType' String Value  = '%v'\n"+
			"'dateTimeFieldDto.FormatType' Integer Value = '%v'\n",
			ePrefix.String(),
			dateTimeFieldDto.FormatType.String(),
			dateTimeFieldDto.FormatType.XValueInt())

		return err
	}

	if dateTimeFieldDto.FieldDateTime.IsZero() {

		err = fmt.Errorf("%v\n"+
			"Error: Format Parameter Date Time is invalid!\n"+
			"'dateTimeFieldDto.FieldDateTime' has a time value of zero.\n",
			ePrefix.String())

		return err

	}

	dateTimeFormat := dateTimeFieldDto.FieldDateTimeFormat

	if len(dateTimeFormat) == 0 {
		dateTimeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultDateTimeFormat()
	}

	var txtDateTimeField TextFieldSpecDateTime

	txtDateTimeField,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTimeFieldDto.FieldDateTime,
		dateTimeFieldDto.FieldLength,
		dateTimeFormat,
		dateTimeFieldDto.FieldJustify,
		ePrefix.XCpy(
			"txtDateTimeField<-dateTime"))

	if err != nil {
		return err
	}

	var dateTimeStr string

	dateTimeStr,
		err = txtDateTimeField.GetFormattedText(
		ePrefix.XCpy(
			"txtDateTimeField"))

	if err != nil {
		return err
	}

	txtBuilderParams := textStrBuilderParamsDto{
		strBuilder:                 strBuilder,
		leftMarginStr:              dateTimeFieldDto.LeftMarginStr,
		lenLeftMarginStr:           len(dateTimeFieldDto.LeftMarginStr),
		textStr:                    dateTimeStr,
		lenTextStr:                 len(dateTimeStr),
		rightMarginStr:             dateTimeFieldDto.RightMarginStr,
		lenRightMarginStr:          len(dateTimeFieldDto.RightMarginStr),
		turnLineTerminationOff:     true,
		lineTerminatorStr:          dateTimeFieldDto.LineTerminator,
		lenLineTerminatorStr:       0,
		maxLineLength:              dateTimeFieldDto.MaxLineLength,
		currentLineLength:          0,
		turnAutoLineLengthBreaksOn: dateTimeFieldDto.TurnAutoLineLengthBreaksOn,
		lastWriteWasLineTerminator: false,
		sourceTag:                  "DateTime",
		sourceDtoTag:               "dateTimeFieldDto",
		errPrefDto:                 ePrefix,
	}

	err = new(textStrBuilderAtom).preBuildScreening(
		&txtBuilderParams)

	return err
}

func (txtBuilderMolecule *textStrBuilderMolecule) buildFieldFillerWithDto(
	strBuilder *strings.Builder,
	fillerFieldDto TextFieldFillerDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderMolecule.lock == nil {
		txtBuilderMolecule.lock = new(sync.Mutex)
	}

	txtBuilderMolecule.lock.Lock()

	defer txtBuilderMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderMolecule."+
			"buildFieldFillerWithDto()",
		"")

	if err != nil {
		return err
	}

	if fillerFieldDto.FormatType !=
		TxtFieldType.Filler() {

		err = fmt.Errorf("%v\n"+
			"Error: 'fillerFieldDto.FormatType' is invalid!\n"+
			"'fillerFieldDto.FormatType' should be set to \n"+
			"TxtFieldType.Filler(). It is NOT!\n"+
			"'fillerFieldDto.FormatType' String Value  = '%v'\n"+
			"'fillerFieldDto.FormatType' Integer Value = '%v'\n",
			ePrefix.String(),
			fillerFieldDto.FormatType.String(),
			fillerFieldDto.FormatType.XValueInt())

		return err
	}

	fillerCharacters := fillerFieldDto.FillerCharacters

	lenFillerChars := len(fillerCharacters)

	if lenFillerChars == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerFieldDto.FillerCharacters' is invalid!\n"+
			"'fillerFieldDto.FillerCharacters' is an empty string.\n",
			ePrefix.String())

		if err != nil {
			return err
		}

	}

	if fillerFieldDto.FillerCharsRepeatCount < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'fillerFieldDto.FillerCharsRepeatCount' is invalid!\n"+
			"'fillerFieldDto.FillerCharsRepeatCount' is an empty string.\n",
			ePrefix.String())

		if err != nil {
			return err
		}

	}

	var txtFillerFieldSpec TextFieldSpecFiller

	txtFillerFieldSpec,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerCharacters,
		fillerFieldDto.FillerCharsRepeatCount,
		ePrefix.XCpy(
			"txtFillerFieldSpec"))

	if err != nil {
		return err
	}

	var fillerCharsStr string

	fillerCharsStr,
		err = txtFillerFieldSpec.GetFormattedText(
		ePrefix.XCpy(
			"txtFillerFieldSpec"))

	if err != nil {
		return err
	}

	txtBuilderParams := textStrBuilderParamsDto{
		strBuilder:                 strBuilder,
		leftMarginStr:              fillerFieldDto.LeftMarginStr,
		lenLeftMarginStr:           len(fillerFieldDto.LeftMarginStr),
		textStr:                    fillerCharsStr,
		lenTextStr:                 len(fillerCharsStr),
		rightMarginStr:             fillerFieldDto.RightMarginStr,
		lenRightMarginStr:          len(fillerFieldDto.RightMarginStr),
		turnLineTerminationOff:     true,
		lineTerminatorStr:          fillerFieldDto.LineTerminator,
		lenLineTerminatorStr:       0,
		maxLineLength:              fillerFieldDto.MaxLineLength,
		currentLineLength:          0,
		turnAutoLineLengthBreaksOn: fillerFieldDto.TurnAutoLineLengthBreaksOn,
		lastWriteWasLineTerminator: false,
		sourceTag:                  "Filler Field",
		sourceDtoTag:               "fillerFieldDto",
		errPrefDto:                 ePrefix,
	}

	return new(textStrBuilderAtom).preBuildScreening(
		&txtBuilderParams)
}

func (txtBuilderMolecule *textStrBuilderMolecule) buildFieldLabelWithDto(
	strBuilder *strings.Builder,
	labelFieldDto TextFieldLabelDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderMolecule.lock == nil {
		txtBuilderMolecule.lock = new(sync.Mutex)
	}

	txtBuilderMolecule.lock.Lock()

	defer txtBuilderMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderMolecule."+
			"buildFieldLabelWithDto()",
		"")

	if err != nil {
		return err
	}

	if labelFieldDto.FormatType !=
		TxtFieldType.Label() {

		err = fmt.Errorf("%v\n"+
			"Error: 'labelFieldDto.FormatType' is invalid!\n"+
			"'labelFieldDto.FormatType' should be set to \n"+
			"TxtFieldType.Label(). It is NOT!\n"+
			"'labelFieldDto.FormatType' String Value  = '%v'\n"+
			"'labelFieldDto.FormatType' Integer Value = '%v'\n",
			ePrefix.String(),
			labelFieldDto.FormatType.String(),
			labelFieldDto.FormatType.XValueInt())

		return err
	}

	if len(labelFieldDto.FieldText) == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'labelFieldDto.FieldText' is invalid!\n"+
			"'labelFieldDto.FieldText' is an empty string.\n",
			ePrefix.String())

		if err != nil {
			return err
		}

	}

	var txtLabelSpec TextFieldSpecLabel

	txtLabelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		labelFieldDto.FieldText,
		labelFieldDto.FieldLength,
		labelFieldDto.FieldJustify,
		ePrefix.XCpy(
			"txtLabelSpec<-labelText"))

	if err != nil {
		return err
	}

	var labelText string

	labelText,
		err = txtLabelSpec.GetFormattedText(
		ePrefix.XCpy(
			"txtLabelSpec"))

	if err != nil {
		return err
	}

	txtBuilderParams := textStrBuilderParamsDto{
		strBuilder:                 strBuilder,
		leftMarginStr:              labelFieldDto.LeftMarginStr,
		lenLeftMarginStr:           len(labelFieldDto.LeftMarginStr),
		textStr:                    labelText,
		lenTextStr:                 len(labelText),
		rightMarginStr:             labelFieldDto.RightMarginStr,
		lenRightMarginStr:          len(labelFieldDto.RightMarginStr),
		turnLineTerminationOff:     true,
		lineTerminatorStr:          labelFieldDto.LineTerminator,
		lenLineTerminatorStr:       0,
		maxLineLength:              labelFieldDto.MaxLineLength,
		currentLineLength:          0,
		turnAutoLineLengthBreaksOn: labelFieldDto.TurnAutoLineLengthBreaksOn,
		lastWriteWasLineTerminator: false,
		sourceTag:                  "Label Field",
		sourceDtoTag:               "labelFieldDto",
		errPrefDto:                 ePrefix,
	}

	return new(textStrBuilderAtom).preBuildScreening(
		&txtBuilderParams)
}

func (txtBuilderMolecule *textStrBuilderMolecule) buildFieldSpacerWithDto(
	strBuilder *strings.Builder,
	spacerFieldDto TextFieldSpacerDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderMolecule.lock == nil {
		txtBuilderMolecule.lock = new(sync.Mutex)
	}

	txtBuilderMolecule.lock.Lock()

	defer txtBuilderMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderMolecule."+
			"buildFieldSpacerWithDto()",
		"")

	if err != nil {
		return err
	}

	if spacerFieldDto.FormatType !=
		TxtFieldType.Spacer() {

		err = fmt.Errorf("%v\n"+
			"Error: 'labelFieldDto.FormatType' is invalid!\n"+
			"'spacerFieldDto.FormatType' should be set to \n"+
			"TxtFieldType.Spacer(). It is NOT!\n"+
			"'spacerFieldDto.FormatType' String Value  = '%v'\n"+
			"'spacerFieldDto.FormatType' Integer Value = '%v'\n",
			ePrefix.String(),
			spacerFieldDto.FormatType.String(),
			spacerFieldDto.FormatType.XValueInt())

		return err
	}

	if spacerFieldDto.FieldLength < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'spacerFieldDto.FieldLength' is invalid!\n"+
			"'spacerFieldDto.FieldLength' has a value less than one (1).\n",
			ePrefix.String())

		if err != nil {
			return err
		}

	}

	var txtFieldSpecSpacer TextFieldSpecSpacer

	txtFieldSpecSpacer,
		err = TextFieldSpecSpacer{}.NewSpacer(
		spacerFieldDto.FieldLength,
		ePrefix.XCpy(
			fmt.Sprintf("spacerFieldDto.FieldLength='%v'",
				spacerFieldDto.FieldLength)))

	if err != nil {
		return err
	}

	var spacerFieldText string

	spacerFieldText,
		err = txtFieldSpecSpacer.GetFormattedText(
		ePrefix.XCpy(
			"spacerFieldText<-txtFieldSpecSpacer"))

	if err != nil {
		return err
	}

	txtBuilderParams := textStrBuilderParamsDto{
		strBuilder:                 strBuilder,
		leftMarginStr:              spacerFieldDto.LeftMarginStr,
		lenLeftMarginStr:           len(spacerFieldDto.LeftMarginStr),
		textStr:                    spacerFieldText,
		lenTextStr:                 len(spacerFieldText),
		rightMarginStr:             spacerFieldDto.RightMarginStr,
		lenRightMarginStr:          len(spacerFieldDto.RightMarginStr),
		turnLineTerminationOff:     true,
		lineTerminatorStr:          spacerFieldDto.LineTerminator,
		lenLineTerminatorStr:       0,
		maxLineLength:              spacerFieldDto.MaxLineLength,
		currentLineLength:          0,
		turnAutoLineLengthBreaksOn: spacerFieldDto.TurnAutoLineLengthBreaksOn,
		lastWriteWasLineTerminator: false,
		sourceTag:                  "Spacer Field",
		sourceDtoTag:               "spacerFieldDto",
		errPrefDto:                 ePrefix,
	}

	return new(textStrBuilderAtom).preBuildScreening(
		&txtBuilderParams)
}

func (txtBuilderMolecule *textStrBuilderMolecule) buildLineAdHocTextWithDto(
	strBuilder *strings.Builder,
	txtAdHocDto TextAdHocDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderMolecule.lock == nil {
		txtBuilderMolecule.lock = new(sync.Mutex)
	}

	txtBuilderMolecule.lock.Lock()

	defer txtBuilderMolecule.lock.Unlock()

	if txtBuilderMolecule.lock == nil {
		txtBuilderMolecule.lock = new(sync.Mutex)
	}

	txtBuilderMolecule.lock.Lock()

	defer txtBuilderMolecule.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderMolecule."+
			"buildLineAdHocTextWithDto()",
		"")

	if err != nil {
		return err
	}

	if txtAdHocDto.FormatType !=
		TxtFieldType.TextAdHoc() {

		err = fmt.Errorf("%v\n"+
			"Error: 'txtAdHocDto.FormatType' is invalid!\n"+
			"'txtAdHocDto.FormatType' should be set to \n"+
			"TxtFieldType.TextAdHoc(). It is NOT!\n"+
			"'txtAdHocDto.FormatType' String Value  = '%v'\n"+
			"'txtAdHocDto.FormatType' Integer Value = '%v'\n",
			ePrefix.String(),
			txtAdHocDto.FormatType.String(),
			txtAdHocDto.FormatType.XValueInt())

		return err
	}

	lenAdHocText := len(txtAdHocDto.AdHocText)

	if lenAdHocText == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtAdHocDto.AdHocText' is invalid!\n"+
			"'txtAdHocDto.AdHocText' is empty an contains zero (0) characters.\n",
			ePrefix.String())

		return err
	}

	txtBuilderParams := textStrBuilderParamsDto{
		strBuilder:                 strBuilder,
		leftMarginStr:              txtAdHocDto.LeftMarginStr,
		lenLeftMarginStr:           len(txtAdHocDto.LeftMarginStr),
		textStr:                    txtAdHocDto.AdHocText,
		lenTextStr:                 len(txtAdHocDto.AdHocText),
		rightMarginStr:             txtAdHocDto.RightMarginStr,
		lenRightMarginStr:          len(txtAdHocDto.RightMarginStr),
		turnLineTerminationOff:     true,
		lineTerminatorStr:          txtAdHocDto.LineTerminator,
		lenLineTerminatorStr:       0,
		maxLineLength:              txtAdHocDto.MaxLineLength,
		currentLineLength:          0,
		turnAutoLineLengthBreaksOn: txtAdHocDto.TurnAutoLineLengthBreaksOn,
		lastWriteWasLineTerminator: false,
		sourceTag:                  "Ad Hoc Text",
		sourceDtoTag:               "txtAdHocDto",
		errPrefDto:                 ePrefix,
	}

	return new(textStrBuilderAtom).preBuildScreening(
		&txtBuilderParams)
}

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

func (txtBuilderMolecule *textStrBuilderMolecule) buildDateTimeFieldWithDto(
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
			"buildDateTimeFieldWithDto()",
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

func (txtBuilderMolecule *textStrBuilderMolecule) buildLabelFieldWithDto(
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
			"buildLabelFieldWithDto()",
		"")

	if err != nil {
		return err
	}

	labelText := labelFieldDto.FieldText

	if len(labelText) == 0 {

		labelText = " "

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

	err = new(textStrBuilderAtom).preBuildScreening(
		&txtBuilderParams)

	return err
}

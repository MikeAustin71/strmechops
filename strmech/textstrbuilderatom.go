package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

// textStrBuilderAtom - Provides helper methods for type
// TextStrBuilder.
//
type textStrBuilderAtom struct {
	lock *sync.Mutex
}

func (txtBuilderAtom *textStrBuilderAtom) preBuildScreening(
	textParams *textStrBuilderParamsDto) error {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		textParams.errPrefDto,
		"textStrBuilderAtom."+
			"preBuildScreening()",
		"")

	if err != nil {
		return err
	}

	if textParams.strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: %v Text Builder\n"+
			"Input parameter 'strBuilder' is invalid!\n"+
			"'strBuilder' is a nil pointer.\n",
			ePrefix.String(),
			textParams.sourceTag)

		if err != nil {
			return err
		}

	}

	if textParams.lenLineTerminatorStr > 0 {
		// Terminator String Length > 0
		// There is a command for line termination
		textParams.turnLineTerminationOff = false
	}

	if textParams.turnLineTerminationOff == true {

		textParams.lenLineTerminatorStr = 0
		textParams.lineTerminatorStr = ""

	} else {
		// MUST BE
		// textParams.turnLineTerminationOff == false

		if textParams.lenLineTerminatorStr == 0 {
			textParams.lenLineTerminatorStr = 1
			textParams.lineTerminatorStr = "\n"
		}
	}

	var adjustedMaximumLineLength = textParams.maxLineLength -
		textParams.lenLeftMarginStr -
		textParams.lenRightMarginStr -
		textParams.lenLineTerminatorStr

	if textParams.turnAutoLineLengthBreaksOn == false {

		textParams.maxLineLength = 900000

	} else {
		// turnAutoLineLengthBreaksOn == true

		if textParams.maxLineLength < 1 {
			err = fmt.Errorf("%v\n"+
				"Error: Maximum Line Length is invalid!\n"+
				"'%v.TurnAutoLineLengthBreaksOn' == 'true'\n"+
				"Maximum Line Length has a value less than one (1)\n"+
				"When Automatic Line Breaks are engaged, Maximum\n"+
				"Line Length must be set to a meaningful number.\n"+
				"Maximum Line Length = '%v'\n",
				ePrefix.String(),
				textParams.sourceDtoTag,
				textParams.maxLineLength)

			return err
		}

		if adjustedMaximumLineLength < 5 {

			err = fmt.Errorf("%v\n"+
				"Error: Adjusted Maximum Line Length is invalid!\n"+
				"This %v Text Line has turned on automatic line breaks. This\n"+
				"means that Line Termination Character(s) will be automatically inserted\n"+
				"when a string of text exceeds the Adjusted Maximum Line Length.\n"+
				"'%v.TurnAutoLineLengthBreaksOn' == 'true'\n"+
				"Adjusted Maximum Line Length is equal to Maximum Line Length\n"+
				"minus Length of Left Margin minus Length of Right Margin\n"+
				"minus Length of Line Termination Characters. With automatic line"+
				"breaks engaged, Adjusted Maximum Line Length MUST BE greater than\n"+
				"or equal to five (5).\n"+
				"SET Maximum Line Length to a meaningful number.\n"+
				"               Maximum Line Length =  %v \n"+
				"               Left Margin Length  = (%v)\n"+
				"               Right Margin Length = (%v)\n"+
				"Line Termination Characters Length = (%v)\n"+
				"                                   ------\n"+
				"      Adjusted Maximum Line Length = '%v'\n",
				ePrefix.String(),
				textParams.sourceTag,
				textParams.sourceDtoTag,
				textParams.maxLineLength,
				textParams.lenLeftMarginStr,
				textParams.lenRightMarginStr,
				textParams.lenLineTerminatorStr,
				adjustedMaximumLineLength)

			return err
		}
	}

	totalStrLen := textParams.lenLineTerminatorStr +
		textParams.lenTextStr +
		textParams.lenRightMarginStr +
		textParams.lenLineTerminatorStr +
		(textParams.lenTextStr / adjustedMaximumLineLength)

	netCapacityStrBuilder :=
		textParams.strBuilder.Cap() -
			textParams.strBuilder.Len()

	requiredCapacity :=
		totalStrLen - netCapacityStrBuilder

	if requiredCapacity > 0 {

		textParams.strBuilder.Grow(requiredCapacity + 10)
	}

	txtBuilderElectron := textStrBuilderElectron{}

	textParams.currentLineLength,
		err =
		txtBuilderElectron.writeLeftMargin(
			textParams.strBuilder,
			textParams.maxLineLength,
			textParams.currentLineLength,
			textParams.leftMarginStr,
			textParams.lineTerminatorStr,
			textParams.turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				fmt.Sprintf("%v.LeftMarginStr",
					textParams.sourceDtoTag)))

	if err != nil {
		return err
	}

	textParams.currentLineLength,
		textParams.lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeText(
			textParams.strBuilder,
			textParams.maxLineLength,
			textParams.currentLineLength,
			textParams.leftMarginStr,
			textParams.textStr,
			textParams.rightMarginStr,
			textParams.lineTerminatorStr,
			textParams.turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				fmt.Sprintf("%v",
					textParams.sourceTag)))

	if err != nil {
		return err
	}

	textParams.currentLineLength,
		textParams.lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeRightMargin(
			textParams.strBuilder,
			textParams.maxLineLength,
			textParams.currentLineLength,
			textParams.rightMarginStr,
			textParams.lineTerminatorStr,
			textParams.turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				fmt.Sprintf("%v.RightMarginStr",
					textParams.sourceDtoTag)))

	if err != nil {
		return err
	}

	if !textParams.lastWriteWasLineTerminator &&
		textParams.lenLineTerminatorStr > 0 {
		textParams.strBuilder.WriteString(
			textParams.lineTerminatorStr)
	}

	return err
}

func (txtBuilderAtom *textStrBuilderAtom) buildAdHocTextWithDto(
	strBuilder *strings.Builder,
	txtAdHocDto TextAdHocDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"buildDateTimeFieldWithDto()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'' is a nil pointer.\n",
			ePrefix.String())

		if err != nil {
			return err
		}

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

	var maximumLineLength = txtAdHocDto.MaxLineLength

	var turnAutoLineLengthBreaksOn = txtAdHocDto.TurnAutoLineLengthBreaksOn

	var turnLineTerminationOff = txtAdHocDto.TurnLineTerminationOff

	var lineTerminatorStr = txtAdHocDto.LineTerminator

	lenLeftMargin := len(txtAdHocDto.LeftMarginStr)
	lenRightMargin := len(txtAdHocDto.RightMarginStr)
	lenLineTerminator := len(lineTerminatorStr)

	if turnLineTerminationOff == true {

		lenLineTerminator = 0
		lineTerminatorStr = ""

	} else {
		// turnLineTerminationOff == false
		if lenLineTerminator == 0 {
			lenLineTerminator = 1
			lineTerminatorStr = "\n"
		}
	}

	totalStrLen := lenLeftMargin +
		lenAdHocText +
		lenRightMargin +
		lenLineTerminator +
		(lenAdHocText / maximumLineLength)

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	requiredCapacity :=
		totalStrLen - netCapacityStrBuilder

	if requiredCapacity > 0 {

		strBuilder.Grow(requiredCapacity + 10)
	}

	var adjustedMaximumLineLength = maximumLineLength -
		lenLeftMargin -
		lenRightMargin -
		lenLineTerminator

	if turnAutoLineLengthBreaksOn == false {

		maximumLineLength = 900000

	} else {
		// turnAutoLineLengthBreaksOn == true
		if adjustedMaximumLineLength < 5 {

			err = fmt.Errorf("%v\n"+
				"Error: Adjusted Maximum Line Length is invalid!\n"+
				"This Ad Hoc Text Line has turned on automatic line breaks. This\n"+
				"means that Line Termination Character(s) will be automatically inserted\n"+
				"when a string of text exceeds the Adjusted Maximum Line Length.\n"+
				"'txtAdHocDto.TurnAutoLineLengthBreaksOn' == 'true'\n"+
				"Adjusted Maximum Line Length is equal to Maximum Line Length\n"+
				" minus Length of Left Margin minus Length of Right Margin\n"+
				"minus Length of Line Termination Characters. With automatic"+
				"line breaks engaged, Adjusted Maximum Line Length MUST BE greater\n"+
				"than or equal to five (5).\n"+
				"SET Maximum Line Length to a meaningful number.\n"+
				"               Maximum Line Length =  %v \n"+
				"               Left Margin Length  = (%v)\n"+
				"               Right Margin Length = (%v)\n"+
				"Line Termination Characters Length = (%v)\n"+
				"                                   ------\n"+
				"      Adjusted Maximum Line Length = '%v'\n",
				ePrefix.String(),
				maximumLineLength,
				lenLeftMargin,
				lenRightMargin,
				lenLineTerminator,
				adjustedMaximumLineLength)

			return err

		}
	}

	txtBuilderElectron := textStrBuilderElectron{}

	var currentLineLength = 0

	currentLineLength,
		err =
		txtBuilderElectron.writeLeftMargin(
			strBuilder,
			maximumLineLength,
			currentLineLength,
			txtAdHocDto.LeftMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"txtAdHocDto.LeftMarginStr"))

	if err != nil {
		return err
	}

	var lastWriteWasLineTerminator bool

	currentLineLength,
		lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeText(
			strBuilder,
			maximumLineLength,
			currentLineLength,
			txtAdHocDto.LeftMarginStr,
			txtAdHocDto.AdHocText,
			txtAdHocDto.RightMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"txtAdHocDto.AdHocText"))

	if err != nil {
		return err
	}

	currentLineLength,
		lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeRightMargin(
			strBuilder,
			maximumLineLength,
			currentLineLength,
			txtAdHocDto.RightMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"txtAdHocDto.RightMarginStr"))

	if err != nil {
		return err
	}

	if !lastWriteWasLineTerminator &&
		len(lineTerminatorStr) > 0 {
		strBuilder.WriteString(lineTerminatorStr)
	}

	return err
}

// buildDateTimeFieldWithDto - Receives a Date Time Field Data
// Transfer Objects and generates a Date Time Text string which
// is returned in a strings.Builder instance.
func (txtBuilderAtom *textStrBuilderAtom) buildDateTimeFieldWithDto(
	strBuilder *strings.Builder,
	dateTimeFieldDto TextFieldDateTimeDto,
	errPrefDto *ePref.ErrPrefixDto) (
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"buildDateTimeFieldWithDto()",
		"")

	if err != nil {
		return err
	}

	if strBuilder == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'strBuilder' is invalid!\n"+
			"'' is a nil pointer.\n",
			ePrefix.String())

		if err != nil {
			return err
		}

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

	var maximumLineLength = dateTimeFieldDto.MaxLineLength

	var turnAutoLineLengthBreaksOn = dateTimeFieldDto.TurnAutoLineLengthBreaksOn

	var turnLineTerminationOff = true // This a text field
	// Assumption is there is no line terminator

	var lineTerminatorStr = dateTimeFieldDto.LineTerminator

	lenLineTerminator := len(lineTerminatorStr)
	lenLeftMargin := len(dateTimeFieldDto.LeftMarginStr)
	lenRightMargin := len(dateTimeFieldDto.RightMarginStr)

	if lenLineTerminator > 0 {
		// Terminator String Length > 0
		// There is a command for line termination
		turnLineTerminationOff = false
	}

	if turnLineTerminationOff == true {

		lenLineTerminator = 0
		lineTerminatorStr = ""

	} else {
		// MUST BE
		// turnLineTerminationOff == false

		if lenLineTerminator == 0 {
			lenLineTerminator = 1
			lineTerminatorStr = "\n"
		}
	}

	var adjustedMaximumLineLength = maximumLineLength -
		lenLeftMargin -
		lenRightMargin -
		lenLineTerminator

	if turnAutoLineLengthBreaksOn == false {

		maximumLineLength = 900000

	} else {
		// turnAutoLineLengthBreaksOn == true
		if adjustedMaximumLineLength < 5 {

			err = fmt.Errorf("%v\n"+
				"Error: Adjusted Maximum Line Length is invalid!\n"+
				"This Date Time Text Line has turned on automatic line breaks. This\n"+
				"means that Line Termination Character(s) will be automatically inserted\n"+
				"when a string of text exceeds the Adjusted Maximum Line Length.\n"+
				"'dateTimeFieldDto.TurnAutoLineLengthBreaksOn' == 'true'\n"+
				"Adjusted Maximum Line Length is equal to Maximum Line Length\n"+
				"minus Length of Left Margin minus Length of Right Margin\n"+
				"minus Length of Line Termination Characters. With automatic line"+
				"breaks engaged, Adjusted Maximum Line Length MUST BE greater than\n"+
				"or equal to five (5).\n"+
				"SET Maximum Line Length to a meaningful number.\n"+
				"               Maximum Line Length =  %v \n"+
				"               Left Margin Length  = (%v)\n"+
				"               Right Margin Length = (%v)\n"+
				"Line Termination Characters Length = (%v)\n"+
				"                                   ------\n"+
				"      Adjusted Maximum Line Length = '%v'\n",
				ePrefix.String(),
				maximumLineLength,
				lenLeftMargin,
				lenRightMargin,
				lenLineTerminator,
				adjustedMaximumLineLength)

			return err

		}
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

	lenDateTimeStr := len(dateTimeStr)

	totalStrLen := lenLeftMargin +
		lenDateTimeStr +
		lenRightMargin +
		lenLineTerminator +
		(lenDateTimeStr / maximumLineLength)

	netCapacityStrBuilder :=
		strBuilder.Cap() -
			strBuilder.Len()

	requiredCapacity :=
		totalStrLen - netCapacityStrBuilder

	if requiredCapacity > 0 {

		strBuilder.Grow(requiredCapacity + 10)
	}

	var currentLineLength = 0

	txtBuilderElectron := textStrBuilderElectron{}

	currentLineLength,
		err =
		txtBuilderElectron.writeLeftMargin(
			strBuilder,
			maximumLineLength,
			currentLineLength,
			dateTimeFieldDto.LeftMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"dateTimeFieldDto.LeftMarginStr"))

	if err != nil {
		return err
	}

	var lastWriteWasLineTerminator bool

	currentLineLength,
		lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeText(
			strBuilder,
			maximumLineLength,
			currentLineLength,
			dateTimeFieldDto.LeftMarginStr,
			dateTimeStr,
			dateTimeFieldDto.RightMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"dateTimeStr"))

	if err != nil {
		return err
	}

	currentLineLength,
		lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeRightMargin(
			strBuilder,
			maximumLineLength,
			currentLineLength,
			dateTimeFieldDto.RightMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"dateTimeFieldDto.RightMarginStr"))

	if err != nil {
		return err
	}

	if !lastWriteWasLineTerminator &&
		len(lineTerminatorStr) > 0 {
		strBuilder.WriteString(lineTerminatorStr)
	}

	return err
}

func (txtBuilderAtom *textStrBuilderAtom) buildFillerFieldWithDto(
	fillerFieldDto TextFieldFillerDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldFillerWithMargins()",
		"")

	if err != nil {
		return strBuilder, err
	}

	fillerCharacters := fillerFieldDto.FillerCharacters

	lenFillerChars := len(fillerCharacters)

	if lenFillerChars == 0 {

		fillerCharacters = " "

		lenFillerChars = 1
	}

	var maximumLineLength = fillerFieldDto.MaxLineLength

	var turnAutoLineLengthBreaksOn = fillerFieldDto.TurnAutoLineLengthBreaksOn

	var turnLineTerminationOff = true // This a text field
	// Assumption is there is no line terminator

	var lineTerminatorStr = fillerFieldDto.LineTerminator

	lenLineTerminator := len(lineTerminatorStr)

	if lenLineTerminator > 0 {
		// Terminator String Length > 0
		// There is a command for line termination
		turnLineTerminationOff = false
	}

	lenLeftMargin := len(fillerFieldDto.LeftMarginStr)
	lenRightMargin := len(fillerFieldDto.RightMarginStr)

	if turnLineTerminationOff == true {

		lenLineTerminator = 0
		lineTerminatorStr = ""

	} else {
		// turnLineTerminationOff == false

		if lenLineTerminator == 0 {
			lenLineTerminator = 1
			lineTerminatorStr = "\n"
		}
	}

	var adjustedMaximumLineLength = maximumLineLength -
		lenLeftMargin -
		lenRightMargin -
		lenLineTerminator

	if turnAutoLineLengthBreaksOn == false {

		maximumLineLength = 900000

	} else {
		// turnAutoLineLengthBreaksOn == true
		if adjustedMaximumLineLength < 5 {

			err = fmt.Errorf("%v\n"+
				"Error: Adjusted Maximum Line Length is invalid!\n"+
				"'dateTimeFieldDto.MaxLineLength' has a value less than five (5).\n"+
				"'dateTimeFieldDto.TurnAutoLineLengthBreaksOn' == 'true'\n"+
				"Adjusted Maximum Line Length is equal to Maximum Line Length\n"+
				" minus Length of Left Margin minus Length of Right Margin\n"+
				"minus Length of Line Termination Characters.\n"+
				"With automatic line breaks engaged, set Maximum Line Length\n"+
				"to a meaninful number greater than or equal to five (5).\n"+
				"               Maximum Line Length = '%v'\n"+
				"               Left Margin Length  = '%v'\n"+
				"               Right Margin Length = '%v'\n"+
				"Line Termination Characters Length = '%v'\n"+
				"      Adjusted Maximum Line Length = '%v'\n",
				ePrefix.String(),
				maximumLineLength,
				lenLeftMargin,
				lenRightMargin,
				lenLineTerminator,
				adjustedMaximumLineLength)

			return strBuilder, err

		}
	}

	var currentLineLength = 0

	txtBuilderElectron := textStrBuilderElectron{}

	currentLineLength,
		err =
		txtBuilderElectron.writeLeftMargin(
			&strBuilder,
			maximumLineLength,
			currentLineLength,
			fillerFieldDto.LeftMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"fillerFieldDto.LeftMarginStr"))

	if err != nil {
		return strBuilder, err
	}

	var txtFillerFieldSpec TextFieldSpecFiller

	txtFillerFieldSpec,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerCharacters,
		fillerFieldDto.FillerCharsRepeatCount,
		ePrefix.XCpy(
			"txtFillerFieldSpec"))

	if err != nil {
		return strBuilder, err
	}

	var fillerCharsStr string

	fillerCharsStr,
		err = txtFillerFieldSpec.GetFormattedText(
		ePrefix.XCpy(
			"txtFillerFieldSpec"))

	if err != nil {
		return strBuilder, err
	}

	currentLineLength,
		_,
		err =
		txtBuilderElectron.writeText(
			&strBuilder,
			maximumLineLength,
			currentLineLength,
			fillerFieldDto.LeftMarginStr,
			fillerCharsStr,
			fillerFieldDto.RightMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"fillerCharsStr"))

	if err != nil {
		return strBuilder, err
	}

	var lastWriteWasLineTerminator bool

	currentLineLength,
		lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeRightMargin(
			&strBuilder,
			maximumLineLength,
			currentLineLength,
			fillerFieldDto.RightMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"fillerFieldDto.RightMarginStr"))

	if err != nil {
		return strBuilder, err
	}

	if !lastWriteWasLineTerminator &&
		len(lineTerminatorStr) > 0 {
		strBuilder.WriteString(lineTerminatorStr)
	}

	return strBuilder, err
}

func (txtBuilderAtom *textStrBuilderAtom) buildLabelFieldWithDto(
	labelFieldDto TextFieldLabelDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	strBuilder.Grow(256)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"buildLabelFieldWithDto()",
		"")

	if err != nil {
		return strBuilder, err
	}

	var maximumLineLength = labelFieldDto.MaxLineLength

	var turnAutoLineLengthBreaksOn = labelFieldDto.TurnAutoLineLengthBreaksOn

	var turnLineTerminationOff = true // This a text field
	// Assumption is there is no line terminator

	var lineTerminatorStr = labelFieldDto.LineTerminator

	lenLineTerminator := len(lineTerminatorStr)

	if lenLineTerminator > 0 {
		// Terminator String Length > 0
		// There is a command for line termination
		turnLineTerminationOff = false
	}

	lenLeftMargin := len(labelFieldDto.LeftMarginStr)
	lenRightMargin := len(labelFieldDto.RightMarginStr)

	if turnLineTerminationOff == true {

		lenLineTerminator = 0
		lineTerminatorStr = ""

	} else {
		// turnLineTerminationOff == false

		if lenLineTerminator == 0 {
			lenLineTerminator = 1
			lineTerminatorStr = "\n"
		}
	}

	var adjustedMaximumLineLength = maximumLineLength -
		lenLeftMargin -
		lenRightMargin -
		lenLineTerminator

	if turnAutoLineLengthBreaksOn == false {

		maximumLineLength = 900000

	} else {
		// turnAutoLineLengthBreaksOn == true
		if adjustedMaximumLineLength < 5 {

			err = fmt.Errorf("%v\n"+
				"Error: Adjusted Maximum Line Length is invalid!\n"+
				"'dateTimeFieldDto.MaxLineLength' has a value less than five (5).\n"+
				"'dateTimeFieldDto.TurnAutoLineLengthBreaksOn' == 'true'\n"+
				"Adjusted Maximum Line Length is equal to Maximum Line Length\n"+
				" minus Length of Left Margin minus Length of Right Margin\n"+
				"minus Length of Line Termination Characters.\n"+
				"With automatic line breaks engaged, set Maximum Line Length\n"+
				"to a meaninful number greater than or equal to five (5).\n"+
				"               Maximum Line Length = '%v'\n"+
				"               Left Margin Length  = '%v'\n"+
				"               Right Margin Length = '%v'\n"+
				"Line Termination Characters Length = '%v'\n"+
				"      Adjusted Maximum Line Length = '%v'\n",
				ePrefix.String(),
				maximumLineLength,
				lenLeftMargin,
				lenRightMargin,
				lenLineTerminator,
				adjustedMaximumLineLength)

			return strBuilder, err

		}
	}

	labelText := labelFieldDto.FieldText

	if len(labelText) == 0 {

		labelText = " "

	}

	var currentLineLength = 0

	txtBuilderElectron := textStrBuilderElectron{}

	currentLineLength,
		err =
		txtBuilderElectron.writeLeftMargin(
			&strBuilder,
			maximumLineLength,
			currentLineLength,
			labelFieldDto.LeftMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"labelFieldDto.LeftMarginStr"))

	if err != nil {
		return strBuilder, err
	}

	var txtLabelSpec TextFieldSpecLabel

	txtLabelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		labelText,
		labelFieldDto.FieldLength,
		labelFieldDto.FieldJustify,
		ePrefix.XCpy(
			"txtLabelSpec<-labelFieldDto.FieldText"))

	if err != nil {
		return strBuilder, err
	}

	var labelStr string

	labelStr,
		err = txtLabelSpec.GetFormattedText(
		ePrefix.XCpy(
			"labelStr<-txtLabelSpec"))

	if err != nil {
		return strBuilder, err
	}

	currentLineLength,
		_,
		err =
		txtBuilderElectron.writeText(
			&strBuilder,
			maximumLineLength,
			currentLineLength,
			labelFieldDto.LeftMarginStr,
			labelStr,
			labelFieldDto.RightMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"labelStr"))

	if err != nil {
		return strBuilder, err
	}

	var lastWriteWasLineTerminator bool

	currentLineLength,
		lastWriteWasLineTerminator,
		err =
		txtBuilderElectron.writeRightMargin(
			&strBuilder,
			maximumLineLength,
			currentLineLength,
			labelFieldDto.RightMarginStr,
			lineTerminatorStr,
			turnAutoLineLengthBreaksOn,
			ePrefix.XCpy(
				"labelFieldDto.RightMarginStr"))

	if err != nil {
		return strBuilder, err
	}

	if !lastWriteWasLineTerminator &&
		len(lineTerminatorStr) > 0 {
		strBuilder.WriteString(lineTerminatorStr)
	}

	return strBuilder, err
}

func (txtBuilderAtom *textStrBuilderAtom) buildSpacerFieldWithDto(
	spacerFieldDto TextFieldSpacerDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"buildSpacerFieldWithDto()",
		"")

	if err != nil {
		return strBuilder, err
	}

	stdLine := TextLineSpecStandardLine{}

	if len(spacerFieldDto.LeftMarginStr) > 0 {
		_,
			err = stdLine.AddTextFieldLabel(
			spacerFieldDto.LeftMarginStr,
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy("spacerFieldDto.LeftMarginStr"))

		if err != nil {
			return strBuilder, err
		}

	}

	_,
		err = stdLine.AddTextFieldSpacer(
		spacerFieldDto.FieldLength,
		ePrefix.XCpy(
			fmt.Sprintf(
				"spacerFieldDto.FieldLength=%v",
				spacerFieldDto.FieldLength)))

	if err != nil {
		return strBuilder, err
	}

	if len(spacerFieldDto.RightMarginStr) > 0 {
		_,
			err = stdLine.AddTextFieldLabel(
			spacerFieldDto.RightMarginStr,
			-1,
			TxtJustify.Left(),
			ePrefix.XCpy("spacerFieldDto.RightMarginStr"))

		if err != nil {
			return strBuilder, err
		}

	}

	if len(spacerFieldDto.LineTerminator) == 0 {

		stdLine.TurnAutoLineTerminationOff()

	} else {

		stdLine.TurnAutoLineTerminationOn()

		err = stdLine.SetNewLineChars(
			spacerFieldDto.LineTerminator,
			ePrefix.XCpy(
				"spacerFieldDto.LineTerminator"))

		if err != nil {
			return strBuilder, err
		}
	}

	var strBuilder2 strings.Builder

	strBuilder2,
		err = stdLine.TextBuilder(
		ePrefix.XCpy(
			"strBuilder2<-stdLine"))

	if err != nil {
		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	return strBuilder, err
}

func (txtBuilderAtom *textStrBuilderAtom) buildTextLineBlankWithDto(
	blankLineDto TextLineBlankDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"buildTextLineBlankWithDto()",
		"")

	if err != nil {
		return strBuilder, err
	}

	if blankLineDto.NumOfBlankLines < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'blankLineDto.NumOfBlankLines' is invalid!\n"+
			"'blankLineDto.NumOfBlankLines' has a value less than one (+1).\n"+
			"blankLineDto.NumOfBlankLines = '%v'\n",
			ePrefix.String(),
			blankLineDto.NumOfBlankLines)

		return strBuilder, err
	}

	if blankLineDto.NumOfBlankLines > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'blankLineDto.NumOfBlankLines' is invalid!\n"+
			"'blankLineDto.NumOfBlankLines' has a value greater than one-million (1,000,000).\n"+
			"blankLineDto.NumOfBlankLines = '%v'\n",
			ePrefix.String(),
			blankLineDto.NumOfBlankLines)

		return strBuilder, err
	}

	var blankLinesSpec TextLineSpecBlankLines

	if len(blankLineDto.LineTerminator) == 0 {

		blankLinesSpec,
			err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
			blankLineDto.NumOfBlankLines,
			ePrefix.XCpy(
				fmt.Sprintf(
					"blankLineDto.NumOfBlankLines='%v'\n",
					blankLineDto.NumOfBlankLines)))

	} else {

		blankLinesSpec,
			err = TextLineSpecBlankLines{}.NewBlankLines(
			blankLineDto.NumOfBlankLines,
			blankLineDto.LineTerminator,
			ePrefix.XCpy(
				fmt.Sprintf(
					"blankLineDto.NumOfBlankLines='%v'\n",
					blankLineDto.NumOfBlankLines)))

	}

	if err != nil {
		return strBuilder, err
	}

	var strBuilder2 strings.Builder

	strBuilder2,
		err = blankLinesSpec.TextBuilder(
		ePrefix.XCpy("" +
			"blankLinesSpec"))

	if err != nil {
		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	return strBuilder, err
}

// buildTextLineColumns - This method is used to generate text for
// Text Line Column configurations. A Text Line Column
// specification describes a single line of text consisting of
// one or more columns containing formatted text. Typically,
// each Text Line Column is terminated with default or custom
// line termination characters. However, line termination is
// optional.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  lineCols                   TextLineColumnsDto
//     - An instance of TextLineColumnsDto which contains the
//       specifications for generating a single line of text
//       containing one or more text columns or text fields.
//
//
//
//
// ----------------------------------------------------------------
//
// Return Values
//
//
//  strBuilder                 strings.Builder
//     - If this method completes successfully, an instance of
//       strings.Builder will be returned containing the formatted
//       string of text characters generated from the Text Line
//       Column specification passed as input parameter,
//       'lineCols'.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtBuilderAtom *textStrBuilderAtom) buildTextLineColumns(
	lineCols TextLineColumnsDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	strBuilder.Grow(512)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"buildTextLineColumns()",
		"")

	if err != nil {
		return strBuilder, err
	}

	if lineCols.FormatType != TxtFieldType.LineColumns() {
		err = fmt.Errorf("%v\n"+
			"Error: 'lineCols.FormatType' is invalid!\n"+
			"'lineCols.FormatType' should be set to \n"+
			"TxtFieldType.LineColumns(). It is NOT!\n"+
			"'lineCols.FormatType' String Value  = '%v'\n"+
			"'lineCols.FormatType' Integer Value = '%v'\n",
			ePrefix.String(),
			lineCols.FormatType.String(),
			lineCols.FormatType.XValueInt())

		return strBuilder, err
	}

	numOfTextFields := len(lineCols.TextFieldsContent)

	if numOfTextFields == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'lineCols.TextFieldsContent' is invalid!\n"+
			"'lineCols.TextFieldsContent' is empty an contains zero (0) elements.\n",
			ePrefix.String())

		return strBuilder, err
	}

	lenFmtParams := len(lineCols.FmtParameters.FieldFormatParams)

	if lenFmtParams != numOfTextFields {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'lineCols is invalid!\n"+
			"The number of Field Format Parameters does\n"+
			"NOT MATCH the number text fields.\n"+
			"'lineCols.TextFieldsContent' Length = '%v'\n"+
			"'lineCols.FmtParameters.FieldFormatParams' Length = '%v'\n",
			ePrefix.String(),
			numOfTextFields,
			lenFmtParams)

		return strBuilder, err

	}

	var columnText, actualDateTimeFormat string

	var maximumLineLength = lineCols.FmtParameters.MaxLineLength

	var turnAutoLineLengthBreaksOn = lineCols.FmtParameters.TurnAutoLineLengthBreaksOn

	if maximumLineLength == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Format Parameter Maximum Line Length is invalid!\n"+
			"'lineCols.FmtParameters.MaxLineLength' has a value of zero (0).\n",
			ePrefix.String())

		return strBuilder, err

	}

	if maximumLineLength < -1 {

		err = fmt.Errorf("%v\n"+
			"Error: Format Parameter Maximum Line Length is invalid!\n"+
			"'lineCols.FmtParameters.MaxLineLength' has a value less"+
			"than minus one (-1).\n"+
			"lineCols.FmtParameters.MaxLineLength= '%v'\n",
			ePrefix.String(),
			maximumLineLength)

		return strBuilder, err

	}

	if turnAutoLineLengthBreaksOn == false {

		maximumLineLength = 900000

	} else {
		// turnAutoLineLengthBreaksOn == true
		if maximumLineLength < 5 {
			turnAutoLineLengthBreaksOn = false
		}
	}

	var currentLineLength = 0
	var lengthOfNextItem = 0
	var nextCurrentLineLength = 0

	defaultDateTimeFormat := textSpecificationMolecule{}.ptr().
		getDefaultDateTimeFormat()

	var txtLabelSpec TextFieldSpecLabel
	var strBuilder2 strings.Builder

	for i := 0; i < numOfTextFields; i++ {

		if !lineCols.TextFieldsContent[i].TextFieldDateTime.IsZero() {
			// Extract DateTime or String for Column Text
			actualDateTimeFormat = lineCols.FmtParameters.FieldFormatParams[i].DateTimeFormat

			if len(actualDateTimeFormat) == 0 {
				actualDateTimeFormat = defaultDateTimeFormat
			}

			columnText =
				lineCols.TextFieldsContent[i].TextFieldDateTime.Format(
					actualDateTimeFormat)
		} else {

			columnText = lineCols.TextFieldsContent[i].TextFieldString

		}

		// Build Left Margin Str
		lengthOfNextItem =
			len(lineCols.FmtParameters.FieldFormatParams[i].LeftMarginStr)

		if lengthOfNextItem > 0 {

			if turnAutoLineLengthBreaksOn {

				nextCurrentLineLength = currentLineLength + lengthOfNextItem

				if nextCurrentLineLength > maximumLineLength {
					currentLineLength = 0
					strBuilder.WriteString("\n")
				}

				currentLineLength += lengthOfNextItem

			}

			strBuilder.WriteString(
				lineCols.FmtParameters.FieldFormatParams[i].LeftMarginStr)

		}

		// Build Column Field Text

		if len(columnText) == 0 {
			columnText = " "
		}

		txtLabelSpec,
			err = TextFieldSpecLabel{}.NewTextLabel(
			columnText,
			lineCols.FmtParameters.FieldFormatParams[i].FieldLength,
			lineCols.FmtParameters.FieldFormatParams[i].FieldJustify,
			ePrefix.XCpy(fmt.Sprintf(
				"columnText[%v]",
				i)))

		if err != nil {
			return strBuilder, err
		}

		strBuilder2,
			err = txtLabelSpec.TextBuilder(
			ePrefix.XCpy(fmt.Sprintf(
				"lineCols.FmtParameters.FieldFormatParams[%v].LeftMarginStr",
				i)))

		if err != nil {
			return strBuilder, err
		}

		lengthOfNextItem = strBuilder2.Len()

		if turnAutoLineLengthBreaksOn {

			nextCurrentLineLength = currentLineLength + lengthOfNextItem

			if nextCurrentLineLength > maximumLineLength {
				currentLineLength = 0
				strBuilder.WriteString("\n")
			}

			if lengthOfNextItem > maximumLineLength {
				var newColStr string
				newColStr,
					err = strMechAtom{}.ptr().
					breakTextAtLineLength(
						strBuilder2.String(),
						maximumLineLength,
						'\n',
						ePrefix.XCpy(
							"txtLabelSpec Very Long Column"))

				if err != nil {

					strBuilder2.Reset()
					return strBuilder, err
				}

				strBuilder2.Reset()
				strBuilder.WriteString(newColStr)
				lengthOfNextItem = 0
				currentLineLength = 0

				continue
			}

			currentLineLength += lengthOfNextItem
		}

		if lengthOfNextItem > 0 {

			strBuilder.WriteString(strBuilder2.String())

			strBuilder2.Reset()

		}

		// Build Right Margin Str

		lengthOfNextItem =
			len(lineCols.FmtParameters.FieldFormatParams[i].RightMarginStr)

		if lengthOfNextItem > 0 {

			if turnAutoLineLengthBreaksOn {

				nextCurrentLineLength = currentLineLength + lengthOfNextItem

				if nextCurrentLineLength > maximumLineLength {
					currentLineLength = 0
					strBuilder.WriteString("\n")
				}

				currentLineLength += lengthOfNextItem
			}

			strBuilder.WriteString(
				lineCols.FmtParameters.FieldFormatParams[i].RightMarginStr)

		}

	} // End of Text Field Loop

	if lineCols.FmtParameters.TurnLineTerminationOff == true {
		// Line Termination is OFF
		// We are commanded to skip Line Terminator
		return strBuilder, err

	}

	// Line Termination is ON
	// We are REQUIRED to provide a Line Terminator
	if len(lineCols.FmtParameters.LineTerminator) > 0 {

		strBuilder.WriteString(
			lineCols.FmtParameters.LineTerminator)

	} else {
		strBuilder.WriteString("\n")
	}

	return strBuilder, err
}

func (txtBuilderAtom *textStrBuilderAtom) buildTextLineSolidWithDto(
	solidLineDto TextLineSolidDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldFillerWithMargins()",
		"")

	if err != nil {
		return strBuilder, err
	}

	if solidLineDto.SolidLineCharRepeatCount < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineDto.SolidLineCharRepeatCount' is invalid!\n"+
			"'solidLineDto.SolidLineCharRepeatCount' has a value less than one (+1).\n"+
			"This means that no Solid Line Characters will be generated.\n"+
			"solidLineDto.SolidLineCharRepeatCount = '%v'\n",
			ePrefix.String(),
			solidLineDto.SolidLineCharRepeatCount)

		return strBuilder, err

	}

	if solidLineDto.SolidLineCharRepeatCount > 1000000 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'solidLineDto.SolidLineCharRepeatCount' is invalid!\n"+
			"'solidLineDto.SolidLineCharRepeatCount' has a value greater than one-million (1,000,000).\n"+
			"solidLineDto.SolidLineCharRepeatCount = '%v'\n",
			ePrefix.String(),
			solidLineDto.SolidLineCharRepeatCount)

		return strBuilder, err
	}

	var txtSpecSolidLine TextLineSpecSolidLine

	txtSpecSolidLine,
		err = TextLineSpecSolidLine{}.NewFullSolidLineConfig(
		solidLineDto.LeftMarginStr,
		solidLineDto.RightMarginStr,
		solidLineDto.SolidLineChars,
		solidLineDto.SolidLineCharRepeatCount,
		solidLineDto.LineTerminator,
		ePrefix.XCpy(
			"txtSpecSolidLine"))

	if err != nil {
		return strBuilder, err
	}

	var strBuilder2 strings.Builder

	strBuilder2,
		err = txtSpecSolidLine.TextBuilder(
		ePrefix.XCpy(
			"txtSpecSolidLine"))

	if err != nil {
		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	return strBuilder, err
}

func (txtBuilderAtom *textStrBuilderAtom) buildTextLinesTimerStartStopWithDto(
	timerStartStopDto TextLineTimerStartStopDto,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"buildTextLinesTimerStartStopWithDto()",
		"")

	if err != nil {
		return strBuilder, err
	}

	var timerLinesSpec *TextLineSpecTimerLines

	timerLinesSpec,
		err = TextLineSpecTimerLines{}.NewFullTimerEvent(
		timerStartStopDto.LeftMarginStr,
		timerStartStopDto.StartTimeLabel,
		timerStartStopDto.StartTime,
		timerStartStopDto.EndTimeLabel,
		timerStartStopDto.EndTime,
		timerStartStopDto.TimeFormat,
		timerStartStopDto.TimeDurationLabel,
		timerStartStopDto.TextLabelFieldLength,
		timerStartStopDto.TextLabelJustification,
		timerStartStopDto.RightMarginStr,
		ePrefix.XCpy(
			"timerLinesSpec<-timerStartStopDto"))

	if err != nil {
		return strBuilder, err
	}

	return timerLinesSpec.TextBuilder(
		ePrefix.XCpy(
			"strBuilder<-timerLinesSpec"))

}

// fieldDateTimeWithMargins - Is designed to produce three text
// elements consolidated and formatted as a single text field.
//
// The three text elements consist of a left margin string, a
// date/time text field and a right margin string.
//
// These three text elements can be configured as a complete line
// of text depending on the value applied to input parameter
// 'lineTerminator'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftMarginStr              string
//     - The contents of the string will be used as the left margin
//       for the 'dateTime' field.
//
//       If no left margin is required, set 'LeftMarginStr' to a
//       zero length or empty string, and no left margin will be
//       created.
//
//
//  dateTime                   time.Time
//     - The date/time value which will be formatted as a text
//       string.
//
//       If this parameter is set equal to zero, no error will
//       be generated.
//
//
//  dateTimeFieldLength        int
//     - Used to format Date/Time Text Fields. This is the length
//       of the text field in which the formatted 'dateTime' string
//       will be displayed. If 'dateTimeFieldLength' is less than
//       the length of the 'dateTime' string, it will be
//       automatically set equal to the 'dateTime' string length.
//
//       To automatically set the value of 'dateTimeFieldLength' to
//       the length of 'dateTime', set this parameter to a value of
//       minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  dateTimeFormat             string
//    - This string will be used to format the date/time value
//      'dateTime' as a text string.
//
//       If this 'dateTimeFormat' string is empty (has a zero
//       length), a default Date/Time format string will be applied
//       as follows:
//         "2006-01-02 15:04:05.000000000 -0700 MST"
//
//
//  dateTimeTextJustify        TextJustify
//      An enumeration value specifying the justification of the
//      'dateTime' string within the text field specified by
//      'dateTimeFieldLength'.
//
//      Text justification can only be evaluated in the context of
//      a text label, field length and a Text Justification object
//      of type TextJustify. This is because text labels with a
//      field length equal to or less than the length of the text
//      label never use text justification. In these cases, text
//      justification is completely ignored.
//
//      If the field length is greater than the length of the text
//      label, text justification must be equal to one of these
//      three valid values:
//          TextJustify(0).Left()
//          TextJustify(0).Right()
//          TextJustify(0).Center()
//
//      You can also use the abbreviated text justification
//      enumeration syntax as follows:
//
//          TxtJustify.Left()
//          TxtJustify.Right()
//          TxtJustify.Center()
//
//
//  rightMarginStr             string
//     - The contents of the string will be used as the right
//       margin for the 'dateTime' field.
//
//       If no right margin is required, set 'RightMarginStr' to a
//       zero length or empty string, and no right margin will be
//       created.
//
//
//  lineTerminator             string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text thereby converting
//       this text element into a valid line of text.
//
//       If a text line is required, setting this string to include
//       a new line character ('\n') will ensure that the three
//       text elements formatted by this method as single text
//       field will constitute a single line of text.
//
//       The most common usage sets this string to a new line
//       character ("\n").
//
//       If Line Termination is NOT required, set 'lineTerminator'
//       to a zero length or empty string and no line termination
//       characters will be created.
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
//  strBuilder                 strings.Builder
//     - If this method completes successfully, it will return an
//       instance of strings.Builder containing a formatted string
//       of text characters.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtBuilderAtom *textStrBuilderAtom) fieldDateTimeWithMargins(
	leftMarginStr string,
	dateTime time.Time,
	dateTimeFieldLength int,
	dateTimeFormat string,
	dateTimeTextJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto
	strBuilder.Grow(256)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldLabelWithMargins()",
		"")

	if err != nil {
		return strBuilder, err
	}

	if len(dateTimeFormat) == 0 {
		dateTimeFormat =
			textSpecificationMolecule{}.ptr().
				getDefaultDateTimeFormat()
	}

	if len(leftMarginStr) > 0 {
		strBuilder.WriteString(leftMarginStr)
	}

	var txtDateTimeField TextFieldSpecDateTime

	txtDateTimeField,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		dateTimeFieldLength,
		dateTimeFormat,
		dateTimeTextJustify,
		ePrefix.XCpy(
			"txtDateTimeField<-dateTime"))

	if err != nil {
		return strBuilder, err
	}

	var strBuilder2 strings.Builder

	strBuilder2,
		err = txtDateTimeField.TextBuilder(
		ePrefix.XCpy(
			"strBuilder<-txtDateTimeField"))

	if err != nil {
		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	if len(rightMarginStr) > 0 {
		strBuilder.WriteString(rightMarginStr)
	}

	if len(lineTerminator) > 0 {
		strBuilder.WriteString(lineTerminator)
	}

	return strBuilder, err
}

// FieldsSingleFiller - Designed to produce three text elements
// consolidated and formatted as a single text field.
//
// The three text elements consist of a left margin string, a Text
// Filler Field and a right margin string.
//
// These three text elements can be configured as a complete line
// of text depending on the value applied to input parameter
// 'lineTerminator'.
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftMarginStr              string
//     - The contents of the string will be used as the left margin
//       for 'labelText' field.
//
//       If no left margin is required, set 'LeftMarginStr' to a
//       zero length or empty string, and no left margin will be
//       created.
//
//
//  fillerCharacters           string
//     - A string containing the text characters which will be
//       included in the Text Filler Field. The final Text Filler
//       Field will be constructed from the filler characters
//       repeated one or more times as specified by the
//       'fillerCharsRepeatCount' parameter.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharacters' is submitted as an empty or zero
//       length string, it will be defaulted to a single white
//       space character.
//
//
//  fillerCharsRepeatCount     int
//     - Controls the number of times 'fillerCharacters' is
//       repeated when constructing the final Text Filler Field
//       returned by this method. The actual length of the string
//       which will populated the completed Text Filler Field is
//       equal to the length of 'fillerCharacters' times the value
//       of 'fillerCharsRepeatCount'.
//
//        Text Field Filler Length =
//          Length of fillerCharacters X fillerCharsRepeatCount
//
//          Example #1: fillerCharacters = "-*"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "-*-*-*"
//
//          Example #2: fillerCharacters = "-"
//                      fillerRepeatCount = 3
//                      Final Text Filler Field = "---"
//
//       If 'fillerCharsRepeatCount' has a value less than one (1) or
//       greater than one-million (1,000,000), an error will be
//       returned.
//
//
//  rightMarginStr             string
//     - The contents of the string will be used as the right
//       margin for the Text Filler Field.
//
//       If no right margin is required, set 'RightMarginStr' to a
//       zero length or empty string, and no right margin will be
//       created.
//
//
//  lineTerminator             string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text thereby converting
//       this text element into a valid line of text.
//
//       If a text line is required, setting this string to include
//       a new line character ('\n') will ensure that the text line
//       consists of the text label field and no other text
//       elements.
//
//       The most common usage sets this string to a new line
//       character ("\n").
//
//       If Line Termination is NOT required, set 'lineTerminator'
//       to a zero length or empty string and no line termination
//       characters will be created.
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
//  strBuilder                 strings.Builder
//     - If this method completes successfully, an instance of
//       strings.Builder will be returned containing a formatted
//       string of text characters.
//
//
//  error
//     - If this method completes successfully and no errors are
//       encountered this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtBuilderAtom *textStrBuilderAtom) fieldFillerWithMargins(
	leftMarginStr string,
	fillerCharacters string,
	fillerCharsRepeatCount int,
	rightMarginStr string,
	lineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	strBuilder.Grow(256)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldFillerWithMargins()",
		"")

	if err != nil {
		return strBuilder, err
	}

	if len(fillerCharacters) == 0 {

		fillerCharacters = " "

	}

	if len(leftMarginStr) > 0 {
		strBuilder.WriteString(leftMarginStr)
	}

	var txtFillerFieldSpec TextFieldSpecFiller

	txtFillerFieldSpec,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerCharacters,
		fillerCharsRepeatCount,
		ePrefix.XCpy(
			"txtFillerFieldSpec"))

	if err != nil {
		return strBuilder, err
	}

	var strBuilder2 strings.Builder

	strBuilder2,
		err = txtFillerFieldSpec.TextBuilder(
		ePrefix.XCpy(
			"strBuilder<-txtFillerFieldSpec"))

	if err != nil {
		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	if len(rightMarginStr) > 0 {
		strBuilder.WriteString(rightMarginStr)
	}

	if len(lineTerminator) > 0 {
		strBuilder.WriteString(lineTerminator)
	}

	return strBuilder, err
}

// fieldLabelWithMargins - Formats a single text label and writes
// the output string to an instance of strings.Builder passed as an
// input parameter by the calling function.
//
// If the Left and Right Margin Strings contain characters, they
// will also be written to the strings.Builder instance
//
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//  leftMarginStr              string
//     - The contents of the string will be used as the left margin
//       for 'labelText' field.
//
//       If no left margin is required, set 'LeftMarginStr' to a
//       zero length or empty string, and no left margin will be
//       created.
//
//
//  labelText                  string
//     - This strings holds the text characters which will be
//       formatted as a text label.
//
//       If 'labelText' is submitted as a zero length or empty
//       string it will automatically be defaulted to a single
//       white space character, " ".
//
//
//  labelFieldLength           int
//     - Used to format Text Label Fields. This is the length of
//       the text field in which the formatted 'labelText' string
//       will be displayed. If 'labelFieldLength' is less than the
//       length of the 'labelText' string, it will be automatically
//       set equal to the 'labelText' string length.
//
//       To automatically set the value of 'labelFieldLength' to
//       the length of 'labelText', set this parameter to a value
//       of  minus one (-1).
//
//       If this parameter is submitted with a value less than
//       minus one (-1) or greater than 1-million (1,000,000), an
//       error will be returned.
//
//
//  labelTextJustify           TextJustify
//      An enumeration value specifying the justification of the
//      'labelText' string within the text field specified by
//      'labelFieldLength'.
//
//      Text justification can only be evaluated in the context of
//      a text label, field length and a Text Justification object
//      of type TextJustify. This is because text labels with a
//      field length equal to or less than the length of the text
//      label never use text justification. In these cases, text
//      justification is completely ignored.
//
//      If the field length is greater than the length of the text
//      label, text justification must be equal to one of these
//      three valid values:
//          TextJustify(0).Left()
//          TextJustify(0).Right()
//          TextJustify(0).Center()
//
//      You can also use the abbreviated text justification
//      enumeration syntax as follows:
//
//          TxtJustify.Left()
//          TxtJustify.Right()
//          TxtJustify.Center()
//
//
//  rightMarginStr             string
//     - The contents of the string will be used as the right
//       margin for the 'labelText' field.
//
//       If no right margin is required, set 'RightMarginStr' to a
//       zero length or empty string, and no right margin will be
//       created.
//
//
//  lineTerminator             string
//     - This string holds the character or characters which will
//       be used to terminate the formatted text thereby converting
//       this text element into a valid line of text.
//
//       If a text line is required, setting this string to include
//       a new line character ('\n') will ensure that the text line
//       consists of the text label field and no other text
//       elements.
//
//       The most common usage sets this string to a new line
//       character ("\n").
//
//       If Line Termination is NOT required, set 'lineTerminator'
//       to a zero length or empty string and no line termination
//       characters will be created.
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
//  strBuilder                 strings.Builder
//     - If this method completes successfully, it will return an
//       instance of strings.Builder containing a formatted text label.
//
//
//  err                        error
//     - If this method completes successfully and no errors are
//       encountered, this return value is set to 'nil'. Otherwise,
//       if errors are encountered, this return value will contain
//       an appropriate error message.
//
//       If an error message is returned, the text value of input
//       parameter 'errorPrefix' will be inserted or prefixed at
//       the beginning of the error message.
//
func (txtBuilderAtom *textStrBuilderAtom) fieldLabelWithMargins(
	leftMarginStr string,
	labelText string,
	labelFieldLength int,
	labelTextJustify TextJustify,
	rightMarginStr string,
	lineTerminator string,
	errPrefDto *ePref.ErrPrefixDto) (
	strBuilder strings.Builder,
	err error) {

	if txtBuilderAtom.lock == nil {
		txtBuilderAtom.lock = new(sync.Mutex)
	}

	txtBuilderAtom.lock.Lock()

	defer txtBuilderAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	strBuilder.Grow(256)

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textStrBuilderAtom."+
			"fieldLabelWithMargins()",
		"")

	if err != nil {
		return strBuilder, err
	}

	if len(labelText) == 0 {

		labelText = " "

	}

	if len(leftMarginStr) > 0 {
		strBuilder.WriteString(leftMarginStr)
	}

	var txtLabelSpec TextFieldSpecLabel

	txtLabelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		labelText,
		labelFieldLength,
		labelTextJustify,
		ePrefix.XCpy(
			"txtLabelSpec<-labelText"))

	if err != nil {
		return strBuilder, err
	}

	var strBuilder2 strings.Builder

	strBuilder2,
		err = txtLabelSpec.TextBuilder(
		ePrefix.XCpy(
			"strBuilder<-txtLabelSpec"))

	if err != nil {
		return strBuilder, err
	}

	strBuilder.WriteString(strBuilder2.String())

	strBuilder2.Reset()

	if len(rightMarginStr) > 0 {
		strBuilder.WriteString(rightMarginStr)
	}

	if len(lineTerminator) > 0 {
		strBuilder.WriteString(lineTerminator)
	}

	return strBuilder, err
}

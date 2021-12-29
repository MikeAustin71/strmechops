package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
	"time"
)

func TestTextLineSpecStandardLine_AddTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextField_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.New()

	labelTxt,
		err := TextFieldSpecLabel{}.NewTextLabel(
		"Hello World",
		-1,
		TxtJustify.Left(),
		ePrefix.XCtx(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var indexId int

	indexId,
		err = stdLine01.AddTextField(
		&labelTxt,
		ePrefix.XCtx(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}.New()

	_,
		err = stdLine02.AddTextField(
		&labelTxt,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine02{}."+
			"AddTextField()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03 := TextLineSpecStandardLine{}.New()

	labelTxt.textLabel = nil
	labelTxt.fieldLen = -95

	_,
		err = stdLine03.AddTextField(
		&labelTxt,
		ePrefix.XCtx(
			"stdLine03<-labelTxt"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextField()\n"+
			"because 'labelTxt' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine04 := TextLineSpecStandardLine{}

	labelTxt.textLabel = []rune("Hello World")
	labelTxt.fieldLen = -1

	_,
		err = stdLine04.AddTextField(
		&labelTxt,
		ePrefix.XCtx(
			"stdLine04<-labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine05 := TextLineSpecStandardLine{}

	_,
		err = stdLine05.AddTextField(
		nil,
		ePrefix.XCtx(
			"stdLine05<-labelTxt"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine05{}."+
			"AddTextField()\n"+
			"because 'textField' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFieldDateTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldDateTime_000100()",
		"")

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(12),
		27,
		15,
		28,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	stdLine01 := TextLineSpecStandardLine{}.New()

	var indexId int

	indexId,
		err = stdLine01.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldDateTime() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	// Add a second Date Time Field
	indexId,
		err = stdLine01.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test # 2\n"+
			"stdLine01.AddTextFieldDateTime() should have\n"+
			"returned 'indexId' = 1\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	indexId,
		err = stdLine02.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine02.AddTextFieldDateTime() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	stdLine03 := TextLineSpecStandardLine{}

	dateTimeBad := time.Time{}
	fieldLenBad := -97

	_,
		err = stdLine03.AddTextFieldDateTime(
		dateTimeBad,
		fieldLenBad,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx(
			"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldDateTime()\n"+
			"because 'dateTimeBad' and 'fieldLenBad' are invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine04 := TextLineSpecStandardLine{}

	_,
		err = stdLine04.AddTextFieldDateTime(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine04{}."+
			"AddTextFieldDateTime()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFieldFiller_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldFiller_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.New()

	fillerCharacters := " "
	fillerCharsRepeatCount := 5

	indexId,
		err := stdLine01.AddTextFieldFiller(
		fillerCharacters,
		fillerCharsRepeatCount,
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {

		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldFiller() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	// Add a second Text Filler Field
	indexId,
		err = stdLine01.AddTextFieldFiller(
		fillerCharacters,
		fillerCharsRepeatCount,
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 1 {

		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldFiller() should have\n"+
			"returned 'indexId' = 1\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	_,
		err = stdLine02.AddTextFieldFiller(
		fillerCharacters,
		fillerCharsRepeatCount,
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	badFillerChars := ""

	_,
		err = stdLine02.AddTextFieldFiller(
		badFillerChars,
		fillerCharsRepeatCount,
		ePrefix.XCtx(
			"stdLine02 - badFillerChars"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldDateTime()\n"+
			"because 'badFillerChars' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}
	_,
		err = stdLine02.AddTextFieldFiller(
		badFillerChars,
		fillerCharsRepeatCount,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldDateTime()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFieldLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldLabel_000100()",
		"")

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	expectedStdLineText :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4) +
			"\n"

	stdLine01 := TextLineSpecStandardLine{}.New()

	indexId,
		err := stdLine01.AddTextFieldLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"stdLine01 - valid label"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldLabel() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	_,
		err = stdLine02.AddTextFieldLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"stdLine02 is empty."))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}.NewPtr()

	_,
		err = stdLine03.AddTextFieldLabel(
		label,
		fieldLen,
		txtJustify,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldLabel()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	badFieldLen := -97

	_,
		err = stdLine03.AddTextFieldLabel(
		label,
		badFieldLen,
		txtJustify,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldLabel()\n"+
			"because 'badFieldLen' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_AddTextFieldSpacer_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_AddTextFieldSpacer_000100()",
		"")

	expectedFieldLen := 4

	expectedStdLineText :=
		strings.Repeat(" ", expectedFieldLen) +
			"\n"

	stdLine01 := TextLineSpecStandardLine{}.New()

	indexId,
		err := stdLine01.AddTextFieldSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"stdLine01 - valid Spacer"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextFieldSpacer() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	_,
		err = stdLine02.AddTextFieldSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"stdLine02 is empty."))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}.NewPtr()

	_,
		err = stdLine03.AddTextFieldSpacer(
		expectedFieldLen,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldSpacer()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	badFieldLen := -97

	_,
		err = stdLine03.AddTextFieldSpacer(
		badFieldLen,
		ePrefix.XCtx(
			"stdLine03 - badFieldLen invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"AddTextFieldSpacer()\n"+
			"because 'badFieldLen' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_CopyIn_000100()",
		"")

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	numOfStdLines := stdLine.GetNumOfStdLines()

	if numOfStdLines != 1 {
		t.Errorf("%v\n"+
			"Error: Expected Number of Standard Lines = '1'.\n"+
			"Instead, Number of Standard Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			numOfStdLines)

		return
	}

	numOfTextFields := stdLine.GetNumOfTextFields()

	if numOfTextFields != 0 {
		t.Errorf("%v\n"+
			"Error: Expected Number of Standard Lines = '0'.\n"+
			"Instead, Number of Standard Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			numOfTextFields)

		return
	}

	newLineChars := stdLine.GetNewLineChars()

	if newLineChars != "\n" {
		t.Errorf("%v\n"+
			"Error: Expected newLineChars = \"\\n\".\n"+
			"Instead, newLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			[]rune(newLineChars))

		return
	}

	newLineRunes := stdLine.GetNewLineRunes()
	expectedNewLineRunes := "\n"

	if string(newLineRunes) != expectedNewLineRunes {
		t.Errorf("%v\n"+
			"Error: Expected newLineRunes = \"\\n\".\n"+
			"Instead, newLineRunes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			[]rune(newLineChars))

		return

	}

	turnLineTerminatorOff :=
		stdLine.GetTurnLineTerminatorOff()

	if turnLineTerminatorOff != false {
		t.Errorf("%v\n"+
			"Error: Expected turnLineTerminatorOff = 'false'.\n"+
			"Instead, turnLineTerminatorOff = 'true'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	leftMargin,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		" ",
		2,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by leftMargin := TextFieldSpecFiller{}.NewTextFiller().\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	var rightMargin *TextFieldSpecFiller

	rightMargin,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		" ",
		2,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by rightMargin = TextFieldSpecFiller{}.NewTextFiller().\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	labelTxt := "Hello World!"

	expectedFinalTxt := "  " + labelTxt + "  \\n"

	var centerLabel *TextFieldSpecLabel

	centerLabel,
		err = TextFieldSpecLabel{}.NewPtrTextLabel(
		labelTxt,
		-1,
		TxtJustify.Left(),
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by centerLabel = TextFieldSpecLabel{}.NewTextLabel().\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	_,
		err = stdLine.AddTextField(
		leftMargin,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(leftMargin).\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	_,
		err = stdLine.AddTextField(
		centerLabel,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(centerLabel).\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	_,
		err = stdLine.AddTextField(
		rightMargin,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(rightMargin).\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	err = stdLine.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.IsValidInstanceError().\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	stdLineTwo := TextLineSpecStandardLine{}

	err = stdLineTwo.CopyIn(
		stdLine,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLineTwo.CopyIn(stdLine).\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	err = stdLineTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLineTwo.IsValidInstanceError().\n"+
			"Error =\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if !stdLineTwo.Equal(stdLine) {
		t.Errorf("%v\n"+
			"Error: Expected stdLineTwo.Equal(stdLine) == 'true'.\n"+
			"Instead, stdLineTwo.Equal(stdLine) == 'false'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var rawOutput string

	rawOutput,
		err = stdLine.GetFormattedText(
		ePrefix.XCtx("stdLine->rawOutput"))

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	sMech := StrMech{}

	actualStr := sMech.ConvertNonPrintableChars(
		[]rune(rawOutput),
		false)

	if expectedFinalTxt != actualStr {
		t.Errorf("%v\n"+
			"Error: Expected stdLine final text output = '%v'.\n"+
			"Instead, stdLine final text output        = '%v'!\n",
			ePrefix.XCtxEmpty().String(),
			expectedFinalTxt,
			actualStr)

		return
	}

	rawOutput,
		err = stdLineTwo.GetFormattedText(
		"stdLineTwo->rawOutput")

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	actualStr = sMech.ConvertNonPrintableChars([]rune(rawOutput), false)

	if expectedFinalTxt != actualStr {
		t.Errorf("%v\n"+
			"Error: Expected stdLineTwo final text output = '%v'.\n"+
			"Instead, stdLineTwo final text output        = '%v'!\n",
			ePrefix.XCtxEmpty().String(),
			expectedFinalTxt,
			actualStr)

		return
	}

	stdLineThree := TextLineSpecStandardLine{}.New()

	err = stdLineThree.CopyIn(
		&stdLineTwo,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecStandardLine{}."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_CopyIn_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCtx(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCtx(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStdLineText :=
		strings.Repeat(" ", leftMarginLen) +
			"  " + label + "  " +
			strings.Repeat(" ", rightMarginLen) +
			"\n"

	var indexId int

	indexId,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 0 {
		t.Errorf("%v - ERROR\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 0\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	collectionCount :=
		stdLine01.GetNumOfTextFields()

	if collectionCount != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 1\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.XCtxEmpty().String(),
			collectionCount)

		return
	}

	indexId,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCtx(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 1 {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 1\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	collectionCount =
		stdLine01.GetNumOfTextFields()

	if collectionCount != 2 {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 2\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.XCtxEmpty().String(),
			collectionCount)

		return
	}

	indexId,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if indexId != 2 {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"stdLine01.AddTextField() should have\n"+
			"returned 'indexId' = 2\n"+
			"HOWEVER, indexId = %v\n",
			ePrefix.XCtxEmpty().String(),
			indexId)

		return
	}

	collectionCount =
		stdLine01.GetNumOfTextFields()

	if collectionCount != 3 {
		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"stdLine01.GetNumOfTextFields() should have\n"+
			"returned 'collectionCount' = 3\n"+
			"HOWEVER, 'collectionCount' = %v\n",
			ePrefix.XCtxEmpty().String(),
			collectionCount)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCtx(
			"stdLine01"))

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = stdLine01.CopyOut(
		ePrefix.XCtx(
			"stdLine01->stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v - ERROR\n"+
			"stdLine02.Equal(stdLine01)\n"+
			"Expected that stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	actualStdLineText,
		err = stdLine02.GetFormattedText(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	_,
		err = stdLine01.CopyOut(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine01{}."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03 := TextLineSpecStandardLine{}

	_,
		err = stdLine03.CopyOut(
		ePrefix.XCtx(
			"stdLine03 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"CopyOut()\n"+
			"because 'stdLine03' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_CopyOutITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_CopyOutITextLine_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCtx(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCtx(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStdLineText :=
		strings.Repeat(" ", leftMarginLen) +
			"  " + label + "  " +
			strings.Repeat(" ", rightMarginLen) +
			"\n"

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCtx(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCtx(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var iTextLine ITextLineSpecification

	iTextLine, err =
		stdLine01.CopyOutITextLine(
			ePrefix.XCtx(
				"stdLine01->iTextLine"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02, ok := iTextLine.(*TextLineSpecStandardLine)

	if !ok {
		t.Errorf("%v\n"+
			"Error: iTextLine.(*TextLineSpecStandardLine)\n"+
			"Could not convert 'iTextLine' to TextLineSpecStandardLine\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v\n"+
			"Error: Expected stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStdLineText),
			true)

	var actualStdLineText string

	actualStdLineText,
		err = stdLine01.GetFormattedText(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1 - stdLine01"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 01.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	actualStdLineText,
		err = stdLine02.GetFormattedText(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStdLineText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2 - stdLine02"+
			"Expected Formatted Text String 02 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	_,
		err =
		stdLine02.CopyOutITextLine(
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine02{}."+
			"CopyOutITextLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03 := TextLineSpecStandardLine{}

	_,
		err = stdLine03.CopyOutITextLine(
		ePrefix.XCtx(
			"stdLine03 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03{}."+
			"CopyOutITextLine()\n"+
			"because 'stdLine03' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

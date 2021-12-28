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
			"AddTextFieldDateTime()\n"+
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

	newLineChars := stdLine.GetLineTerminationChars()

	if newLineChars != "\n" {
		t.Errorf("%v\n"+
			"Error: Expected newLineChars = \"\\n\".\n"+
			"Instead, newLineChars = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			[]rune(newLineChars))

		return
	}

	newLineRunes := stdLine.GetLineTerminationRunes()
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

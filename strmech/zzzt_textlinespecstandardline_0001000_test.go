package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecStandardLine_AddTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_CopyIn_000100()",
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

	_,
		err = stdLine01.AddTextField(
		&labelTxt,
		ePrefix.XCtx(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
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

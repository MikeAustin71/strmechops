package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecStandardLine_CopyIn_000100(t *testing.T) {

	funcName := "TestTextLineSpecStandardLine_CopyIn_000100()"

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		funcName,
		"")

	stdLine := TextLineSpecStandardLine{}.NewPtr()

	numOfStdLines := stdLine.GetNumOfStdLines()

	if numOfStdLines != 1 {
		t.Errorf("%v\n"+
			"Error: Expected Number of Standard Lines = '1'.\n"+
			"Instead, Number of Standard Lines = '%v'\n",
			funcName,
			numOfStdLines)

		return
	}

	numOfTextFields := stdLine.GetNumOfTextFields()

	if numOfTextFields != 0 {
		t.Errorf("%v\n"+
			"Error: Expected Number of Standard Lines = '0'.\n"+
			"Instead, Number of Standard Lines = '%v'\n",
			funcName,
			numOfTextFields)

		return
	}

	newLineChars := stdLine.GetLineTerminationChars()

	if newLineChars != "\n" {
		t.Errorf("%v\n"+
			"Error: Expected newLineChars = \"\\n\".\n"+
			"Instead, newLineChars = '%v'\n",
			funcName,
			[]rune(newLineChars))

		return
	}

	newLineRunes := stdLine.GetLineTerminationRunes()
	expectedNewLineRunes := "\n"

	if string(newLineRunes) != expectedNewLineRunes {
		t.Errorf("%v\n"+
			"Error: Expected newLineRunes = \"\\n\".\n"+
			"Instead, newLineRunes = '%v'\n",
			funcName,
			[]rune(newLineChars))

		return

	}

	turnLineTerminatorOff :=
		stdLine.GetTurnLineTerminatorOff()

	if turnLineTerminatorOff != false {
		t.Errorf("%v\n"+
			"Error: Expected turnLineTerminatorOff = 'false'.\n"+
			"Instead, turnLineTerminatorOff = 'true'\n",
			funcName)

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
			funcName,
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
			funcName,
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
			funcName,
			err.Error())

		return
	}

	err = stdLine.AddTextField(
		leftMargin,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(leftMargin).\n"+
			"Error =\n%v\n",
			funcName,
			err.Error())

		return
	}

	err = stdLine.AddTextField(
		centerLabel,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(centerLabel).\n"+
			"Error =\n%v\n",
			funcName,
			err.Error())

		return
	}

	err = stdLine.AddTextField(
		rightMargin,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.AddTextField(rightMargin).\n"+
			"Error =\n%v\n",
			funcName,
			err.Error())

		return
	}

	err = stdLine.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLine.IsValidInstanceError().\n"+
			"Error =\n%v\n",
			funcName,
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
			funcName,
			err.Error())

		return
	}

	err = stdLineTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by stdLineTwo.IsValidInstanceError().\n"+
			"Error =\n%v\n",
			funcName,
			err.Error())

		return
	}

	if !stdLineTwo.Equal(stdLine) {
		t.Errorf("%v\n"+
			"Error: Expected stdLineTwo.Equal(stdLine) == 'true'.\n"+
			"Instead, stdLineTwo.Equal(stdLine) == 'false'!\n",
			funcName)

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

	actualStr := sMech.ConvertNonPrintableChars([]rune(rawOutput), false)

	if expectedFinalTxt != actualStr {
		t.Errorf("%v\n"+
			"Error: Expected stdLine final text output = '%v'.\n"+
			"Instead, stdLine final text output        = '%v'!\n",
			funcName,
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
			funcName,
			expectedFinalTxt,
			actualStr)
	}

	return
}

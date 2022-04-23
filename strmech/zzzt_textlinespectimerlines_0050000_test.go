package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecTimerLinesAtom_testValidityOfTxtSpecTimerLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesAtom_testValidityOfTxtSpecTimerLines_000100()",
		"")

	txtTimerLinesAtom := textLineSpecTimerLinesAtom{}

	_,
		err := txtTimerLinesAtom.testValidityOfTxtSpecTimerLines(
		nil,
		ePrefix.XCpy(
			"txtTimerLines == nil"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = timerLines02.IsValidInstanceError()\n"+
			"Expected an error return because input\n"+
			"parameter 'txtTimerLines' is 'nil'.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLinesNanobot_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesNanobot_copyIn_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	_,
		timerLines02,
		err = createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtTimerLinesNanobot := textLineSpecTimerLinesNanobot{}

	err =
		txtTimerLinesNanobot.copyIn(
			timerLines02,
			nil,
			ePrefix.XCpy(
				"timerLines02<-nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'incomingTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtTimerLinesNanobot.copyIn(
			nil,
			timerLines01,
			ePrefix.XCpy(
				"nil<-timerLines01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'targetTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err =
		txtTimerLinesNanobot.copyIn(
			timerLines02,
			&timerLines03,
			ePrefix.XCpy(
				"timerLines02<-timerLines03 (invalid)"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'incomingTimerLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLinesNanobot_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesNanobot_copyOut_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	txtTimerLinesNanobot := textLineSpecTimerLinesNanobot{}

	timerLines02,
		err = txtTimerLinesNanobot.copyOut(
		timerLines01,
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines02.GetFormattedText(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines02.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return
	}

	_,
		err = txtTimerLinesNanobot.copyOut(
		nil,
		ePrefix.XCpy(
			"txtTimerLines is nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyOut()\n"+
			"because 'txtTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	_,
		err = txtTimerLinesNanobot.copyOut(
		&timerLines03,
		ePrefix.XCpy(
			"timerLines03 is invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyOut()\n"+
			"because 'txtTimerLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLinesMolecule_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesMolecule_getFormattedText_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtTimerLinesMolecule := textLineSpecTimerLinesMolecule{}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		timerLines01,
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		nil,
		ePrefix.XCpy(
			"timerLines01-Test#2"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesMolecule."+
			"getFormattedText()\n"+
			"because 'txtTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		&timerLines02,
		ePrefix.XCpy(
			"timerLines02-Test#1"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesMolecule."+
			"getFormattedText()\n"+
			"because 'txtTimerLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

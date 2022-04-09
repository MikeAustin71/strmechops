package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecTimerLines_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetFormattedText_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCtx(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines01.GetFormattedText(
		ePrefix.XCtx(
			"timerLines01"))

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
			"timerLines01.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			outputStr,
			printableActualStr)

		return

	}

	return
}

func TestTextLineSpecTimerLines_GetFormattedText_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTESTSERIES_TESTMETHOD_000200()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCtx(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines01.GetFormattedText(
		ePrefix.XCtx(
			"timerLines01"))

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
			"timerLines01.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			outputStr,
			printableActualStr)

		return

	}

	return
}

func TestTextLineSpecTimerLines_GetFormattedText_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLines_GetFormattedText_000300()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCtx(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = timerLines01.GetFormattedText(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from timerLines01."+
			"GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

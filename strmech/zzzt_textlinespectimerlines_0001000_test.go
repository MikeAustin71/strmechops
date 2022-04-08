package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecTimerLines_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTESTSERIES_TESTMETHOD_000100()",
		"")

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
		"timerLines01")

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

	printableExpectedStr :=
		"[SPACE][SPACE]Start[SPACE]Time:[SPACE]2022-04-05[SPACE]10:00:00.000000000[SPACE]-0500[SPACE]CDT\\n" +
			"[SPACE][SPACE][SPACE][SPACE]End[SPACE]Time:[SPACE]2022-04-05[SPACE]10:00:00.000005999[SPACE]" +
			"-0500[SPACE]CDT\\nElapsed[SPACE]Time:[SPACE]5[SPACE]Microseconds[SPACE]999[SPACE]Nanoseconds\\n" +
			"[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]" +
			"Total[SPACE]Elapsed[SPACE]Nanoseconds:[SPACE]5,999\\n"

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines01.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return

	}

	return
}

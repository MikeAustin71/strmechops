package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestNumStrMathRoundingRoundHalfAwayFromZero_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfAwayFromZero_000100()",
		"")

	intRunes,
		err := RuneArrayDto{}.NewRunes(
		[]rune("1"),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fracRunes RuneArrayDto

	fracRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune("23456789"),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	nStrMathRoundAtom := numStrMathRoundingAtom{}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		3,
		ePrefix.XCpy(
			"roundToFracDigits=3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr := intRunes.GetCharacterString() +
		"." +
		fracRunes.GetCharacterString()

	expectedStr := "1.235"

	if expectedStr != actualStr {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero()\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	expectedIntDigits := "1"

	expectedFracDigits := "9999999"

	intRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune(expectedIntDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fracRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune(expectedFracDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		3,
		ePrefix.XCpy(
			"roundToFracDigits=3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr = intRunes.GetCharacterString() +
		"." +
		fracRunes.GetCharacterString()

	expectedStr = "2.000"

	if expectedStr != actualStr {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() #2\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	expectedIntDigits = "9"

	expectedFracDigits = "9999999"

	intRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune(expectedIntDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fracRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune(expectedFracDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		3,
		ePrefix.XCpy(
			"roundToFracDigits=3"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr = intRunes.GetCharacterString() +
		"." +
		fracRunes.GetCharacterString()

	expectedStr = "10.000"

	if expectedStr != actualStr {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() #2\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		nil,
		&fracRunes,
		3,
		ePrefix.XCpy(
			"Error Test #1"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() Error #1\n"+
			"Expected an error return because 'integerDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		nil,
		3,
		ePrefix.XCpy(
			"Error Test #2"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() Error #2\n"+
			"Expected an error return because 'fractionalDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		-1,
		ePrefix.XCpy(
			"Error Test #3"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() Error #2\n"+
			"Expected an error return because 'roundToFractionalDigits' is minus one (-1)\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	return
}

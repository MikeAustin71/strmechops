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

	return
}

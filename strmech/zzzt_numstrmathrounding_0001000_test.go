package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestNumStrMathRoundingRoundHalfUpWithNegNums_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfUpWithNegNums_000100()",
		"")

	// Test # 1
	expectedIntDigits := "7"

	expectedFracDigits := "5"

	expectedStr := "8"

	roundToFractionalDigits := 0

	intRunes,
		err := RuneArrayDto{}.NewRunes(
		[]rune(expectedIntDigits),
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
		[]rune(expectedFracDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	nStrMathRoundAtom := numStrMathRoundingAtom{}

	err = nStrMathRoundAtom.roundHalfUpWithNegNums(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Positive(),
		ePrefix.XCpy(
			"roundToFracDigits=0 Test#1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr := intRunes.GetCharacterString()

	if fracRunes.GetRuneArrayLength() > 0 {

		actualStr += "." +
			fracRunes.GetCharacterString()

	}

	if expectedStr != actualStr {

		t.Errorf("%v - Test #1\n"+
			"Error: nStrMathRoundAtom.roundHalfUpWithNegNums()\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	// Test # 2
	expectedIntDigits = "7"

	expectedFracDigits = "5"

	expectedStr = "7"

	roundToFractionalDigits = 0

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
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfUpWithNegNums(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"roundToFracDigits=0 Test#2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr = intRunes.GetCharacterString()

	if fracRunes.GetRuneArrayLength() > 0 {

		actualStr += "." +
			fracRunes.GetCharacterString()

	}

	if expectedStr != actualStr {

		t.Errorf("%v - Test #2\n"+
			"Error: nStrMathRoundAtom.roundHalfUpWithNegNums()\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

}

func TestNumStrMathRoundingRoundHalfAwayFromZero_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfAwayFromZero_000100()",
		"")

	// Test # 1
	expectedIntDigits := "1"

	expectedFracDigits := "23456789"

	roundToFractionalDigits := 3

	intRunes,
		err := RuneArrayDto{}.NewRunes(
		[]rune(expectedIntDigits),
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
		[]rune(expectedFracDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	nStrMathRoundAtom := numStrMathRoundingAtom{}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Positive(),
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

	// Test # 2
	expectedIntDigits = "1"

	expectedFracDigits = "9999999"

	roundToFractionalDigits = 3

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
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Positive(),
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

	// Test # 3

	expectedIntDigits = "9"

	expectedFracDigits = "9999999"

	roundToFractionalDigits = 3

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
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Positive(),
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
		NumSignVal.Positive(),
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
		NumSignVal.Positive(),
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
		NumSignVal.Positive(),
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

func TestNumStrMathRoundingRoundHalfAwayFromZero_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfAwayFromZero_000200()",
		"")

	// Test # 1
	expectedIntDigits := "1"

	expectedFracDigits := "23456789"

	roundToFractionalDigits := 3

	intRunes,
		err := RuneArrayDto{}.NewRunes(
		[]rune(expectedIntDigits),
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
		[]rune(expectedFracDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	nStrMathRoundAtom := numStrMathRoundingAtom{}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Negative(),
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

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
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() #1\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	// Test # 2
	expectedIntDigits = "1"

	expectedFracDigits = "9999999"

	roundToFractionalDigits = 3

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
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Negative(),
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

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

	// Test # 3
	expectedIntDigits = "9"

	expectedFracDigits = "9999999"

	roundToFractionalDigits = 3

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
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Negative(),
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

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
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() #3\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	// Test #4
	expectedIntDigits = "0"

	expectedFracDigits = "00"

	roundToFractionalDigits = 5

	expectedStr = "0.00000"

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
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Zero(),
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr = intRunes.GetCharacterString() +
		"." +
		fracRunes.GetCharacterString()

	if expectedStr != actualStr {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() #4\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	// Test #5
	expectedIntDigits = "1"

	expectedFracDigits = "5"

	roundToFractionalDigits = 0

	expectedStr = "2"

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
			"fracRunes"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrMathRoundAtom.roundHalfAwayFromZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		NumSignVal.Zero(),
		ePrefix.XCpy(
			fmt.Sprintf("Test #5 roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualStr = intRunes.GetCharacterString()

	if fracRunes.GetRuneArrayLength() > 0 {

		actualStr += "." +
			fracRunes.GetCharacterString()

	}

	if expectedStr != actualStr {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() #5\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	return
}

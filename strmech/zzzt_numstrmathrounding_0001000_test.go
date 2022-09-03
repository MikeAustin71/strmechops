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

func TestNumStrMathRoundingRoundHalfDownWithNegNums_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfDownWithNegNums_000100()",
		"")

	const arrayLen = 12

	var expectedIntDigits, expectedFracDigits, expectedNumStr [arrayLen]string

	var roundToFractionalDigits [arrayLen]int

	var numberSigns [arrayLen]NumericSignValueType

	expectedIntDigits[0] = "7"
	expectedFracDigits[0] = "6"
	expectedNumStr[0] = "8"
	roundToFractionalDigits[0] = 0
	numberSigns[0] = NumSignVal.Positive()

	expectedIntDigits[1] = "7"
	expectedFracDigits[1] = "5"
	expectedNumStr[1] = "7"
	roundToFractionalDigits[1] = 0
	numberSigns[1] = NumSignVal.Positive()

	expectedIntDigits[2] = "7"
	expectedFracDigits[2] = "4"
	expectedNumStr[2] = "7"
	roundToFractionalDigits[2] = 0
	numberSigns[2] = NumSignVal.Positive()

	expectedIntDigits[3] = "7"
	expectedFracDigits[3] = "4"
	expectedNumStr[3] = "7"
	roundToFractionalDigits[3] = 0
	numberSigns[3] = NumSignVal.Negative()

	expectedIntDigits[4] = "7"
	expectedFracDigits[4] = "5"
	expectedNumStr[4] = "8"
	roundToFractionalDigits[4] = 0
	numberSigns[4] = NumSignVal.Negative()

	expectedIntDigits[5] = "7"
	expectedFracDigits[5] = "6"
	expectedNumStr[5] = "8"
	roundToFractionalDigits[5] = 0
	numberSigns[5] = NumSignVal.Negative()

	expectedIntDigits[6] = "7"
	expectedFracDigits[6] = "1236"
	expectedNumStr[6] = "7.124"
	roundToFractionalDigits[6] = 3
	numberSigns[6] = NumSignVal.Negative()

	expectedIntDigits[7] = "7"
	expectedFracDigits[7] = "1235"
	expectedNumStr[7] = "7.124"
	roundToFractionalDigits[7] = 3
	numberSigns[7] = NumSignVal.Negative()

	expectedIntDigits[8] = "7"
	expectedFracDigits[8] = "1234"
	expectedNumStr[8] = "7.123"
	roundToFractionalDigits[8] = 3
	numberSigns[8] = NumSignVal.Negative()

	expectedIntDigits[9] = "7"
	expectedFracDigits[9] = "1236"
	expectedNumStr[9] = "7.124"
	roundToFractionalDigits[9] = 3
	numberSigns[9] = NumSignVal.Positive()

	expectedIntDigits[10] = "7"
	expectedFracDigits[10] = "1235"
	expectedNumStr[10] = "7.123"
	roundToFractionalDigits[10] = 3
	numberSigns[10] = NumSignVal.Positive()

	expectedIntDigits[11] = "7"
	expectedFracDigits[11] = "1234"
	expectedNumStr[11] = "7.123"
	roundToFractionalDigits[11] = 3
	numberSigns[11] = NumSignVal.Positive()

	var actualStr string

	var intRunes, fracRunes RuneArrayDto
	var err error
	nStrMathRoundAtom := numStrMathRoundingAtom{}

	for i := 0; i < arrayLen; i++ {

		intRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedIntDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"intRunes[%v]=%v",
					i,
					expectedIntDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		fracRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedFracDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"fracRunes[%v]=%v",
					i,
					expectedFracDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		err = nStrMathRoundAtom.roundHalfDownWithNegNums(
			&intRunes,
			&fracRunes,
			roundToFractionalDigits[i],
			numberSigns[i],
			ePrefix.XCpy(
				fmt.Sprintf(
					"roundToFracDigits[%v]=%v",
					i,
					roundToFractionalDigits[i])))

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

		if expectedNumStr[i] != actualStr {

			t.Errorf("%v - Test #1\n"+
				"Error: nStrMathRoundAtom.roundHalfDownWithNegNums()\n"+
				"Cycle Number    = '%v'\n"+
				"Integer Base    = '%v'\n"+
				"Fractional Base = '%v'\n"+
				"Round To Digit  = '%v'\n"+
				"Number Sign     = '%v'\n"+
				"Expected String = '%v'\n"+
				"  Actual String = '%v'\n",
				ePrefix.String(),
				i,
				expectedIntDigits[i],
				expectedFracDigits[i],
				roundToFractionalDigits[i],
				numberSigns[i].String(),
				expectedNumStr[i],
				actualStr)

			return
		}

	}

	return
}

func TestNumStrMathRoundingRoundHalfDownWithNegNums_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfDownWithNegNums_000200()",
		"")

	// Test # 1
	expectedIntDigits := "1"

	expectedFracDigits := "23456789"

	roundToFractionalDigits := 3

	numberSign := NumSignVal.Negative()

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

	err = nStrMathRoundAtom.roundHalfDownWithNegNums(
		nil,
		&fracRunes,
		roundToFractionalDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfDownWithNegNums() Error #1\n"+
			"Expected an error return because 'integerDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfDownWithNegNums(
		&intRunes,
		nil,
		3,
		numberSign,
		ePrefix.XCpy(
			"Error Test #2"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfDownWithNegNums() Error #2\n"+
			"Expected an error return because 'fractionalDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfDownWithNegNums(
		&intRunes,
		&fracRunes,
		-1,
		numberSign,
		ePrefix.XCpy(
			"Error Test #3"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfDownWithNegNums() Error #2\n"+
			"Expected an error return because 'roundToFractionalDigits' is minus one (-1)\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfDownWithNegNums(
		&intRunes,
		&fracRunes,
		3,
		NumSignVal.None(),
		ePrefix.XCpy(
			"Error Test #4"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfDownWithNegNums() Error #2\n"+
			"Expected an error return because 'Number Sign' is 'None'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

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

func TestNumStrMathRoundingRoundHalfTowardsZero_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfTowardsZero_000100()",
		"")

	const arrayLen = 12

	var expectedIntDigits, expectedFracDigits, expectedNumStr [arrayLen]string

	var roundToFractionalDigits [arrayLen]int

	var numberSigns [arrayLen]NumericSignValueType

	expectedIntDigits[0] = "7"
	expectedFracDigits[0] = "6"
	expectedNumStr[0] = "8"
	roundToFractionalDigits[0] = 0
	numberSigns[0] = NumSignVal.Positive()

	expectedIntDigits[1] = "7"
	expectedFracDigits[1] = "5"
	expectedNumStr[1] = "7"
	roundToFractionalDigits[1] = 0
	numberSigns[1] = NumSignVal.Positive()

	expectedIntDigits[2] = "7"
	expectedFracDigits[2] = "4"
	expectedNumStr[2] = "7"
	roundToFractionalDigits[2] = 0
	numberSigns[2] = NumSignVal.Positive()

	expectedIntDigits[3] = "7"
	expectedFracDigits[3] = "4"
	expectedNumStr[3] = "7"
	roundToFractionalDigits[3] = 0
	numberSigns[3] = NumSignVal.Negative()

	expectedIntDigits[4] = "7"
	expectedFracDigits[4] = "5"
	expectedNumStr[4] = "7"
	roundToFractionalDigits[4] = 0
	numberSigns[4] = NumSignVal.Negative()

	expectedIntDigits[5] = "7"
	expectedFracDigits[5] = "6"
	expectedNumStr[5] = "8"
	roundToFractionalDigits[5] = 0
	numberSigns[5] = NumSignVal.Negative()

	expectedIntDigits[6] = "7"
	expectedFracDigits[6] = "1236"
	expectedNumStr[6] = "7.124"
	roundToFractionalDigits[6] = 3
	numberSigns[6] = NumSignVal.Negative()

	expectedIntDigits[7] = "7"
	expectedFracDigits[7] = "1235"
	expectedNumStr[7] = "7.123"
	roundToFractionalDigits[7] = 3
	numberSigns[7] = NumSignVal.Negative()

	expectedIntDigits[8] = "7"
	expectedFracDigits[8] = "1234"
	expectedNumStr[8] = "7.123"
	roundToFractionalDigits[8] = 3
	numberSigns[8] = NumSignVal.Negative()

	expectedIntDigits[9] = "7"
	expectedFracDigits[9] = "1236"
	expectedNumStr[9] = "7.124"
	roundToFractionalDigits[9] = 3
	numberSigns[9] = NumSignVal.Positive()

	expectedIntDigits[10] = "7"
	expectedFracDigits[10] = "1235"
	expectedNumStr[10] = "7.123"
	roundToFractionalDigits[10] = 3
	numberSigns[10] = NumSignVal.Positive()

	expectedIntDigits[11] = "7"
	expectedFracDigits[11] = "1234"
	expectedNumStr[11] = "7.123"
	roundToFractionalDigits[11] = 3
	numberSigns[11] = NumSignVal.Positive()

	var actualStr string

	var intRunes, fracRunes RuneArrayDto
	var err error
	nStrMathRoundAtom := numStrMathRoundingAtom{}

	for i := 0; i < arrayLen; i++ {

		intRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedIntDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"intRunes[%v]=%v",
					i,
					expectedIntDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		fracRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedFracDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"fracRunes[%v]=%v",
					i,
					expectedFracDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		err = nStrMathRoundAtom.roundHalfTowardsZero(
			&intRunes,
			&fracRunes,
			roundToFractionalDigits[i],
			numberSigns[i],
			ePrefix.XCpy(
				fmt.Sprintf(
					"roundToFracDigits[%v]=%v",
					i,
					roundToFractionalDigits[i])))

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

		if expectedNumStr[i] != actualStr {

			t.Errorf("%v - Test #1\n"+
				"Error: nStrMathRoundAtom.roundHalfTowardsZero()\n"+
				"Cycle Number    = '%v'\n"+
				"Integer Base    = '%v'\n"+
				"Fractional Base = '%v'\n"+
				"Round To Digit  = '%v'\n"+
				"Number Sign     = '%v'\n"+
				"Expected String = '%v'\n"+
				"  Actual String = '%v'\n",
				ePrefix.String(),
				i,
				expectedIntDigits[i],
				expectedFracDigits[i],
				roundToFractionalDigits[i],
				numberSigns[i].String(),
				expectedNumStr[i],
				actualStr)

			return
		}

	}

	return
}

func TestNumStrMathRoundingRoundHalfTowardsZero_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfTowardsZero_000200()",
		"")

	// Test # 1
	expectedIntDigits := "1"

	expectedFracDigits := "23456789"

	roundToFractionalDigits := 3

	numberSign := NumSignVal.Negative()

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

	err = nStrMathRoundAtom.roundHalfTowardsZero(
		nil,
		&fracRunes,
		roundToFractionalDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfTowardsZero() Error #1\n"+
			"Expected an error return because 'integerDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfTowardsZero(
		&intRunes,
		nil,
		3,
		numberSign,
		ePrefix.XCpy(
			"Error Test #2"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfTowardsZero() Error #2\n"+
			"Expected an error return because 'fractionalDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfTowardsZero(
		&intRunes,
		&fracRunes,
		-1,
		numberSign,
		ePrefix.XCpy(
			"Error Test #3"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfTowardsZero() Error #2\n"+
			"Expected an error return because 'roundToFractionalDigits' is minus one (-1)\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfTowardsZero(
		&intRunes,
		&fracRunes,
		3,
		NumSignVal.None(),
		ePrefix.XCpy(
			"Error Test #4"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfTowardsZero() Error #2\n"+
			"Expected an error return because 'Number Sign' is 'None'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

}

func TestNumStrMathRoundingRoundHalfToEven_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfToEven_000100()",
		"")

	const arrayLen = 16

	var expectedIntDigits, expectedFracDigits, expectedNumStr [arrayLen]string

	var roundToFractionalDigits [arrayLen]int

	var numberSigns [arrayLen]NumericSignValueType

	expectedIntDigits[0] = "7"
	expectedFracDigits[0] = "6"
	expectedNumStr[0] = "8"
	roundToFractionalDigits[0] = 0
	numberSigns[0] = NumSignVal.Positive()

	expectedIntDigits[1] = "7"
	expectedFracDigits[1] = "5"
	expectedNumStr[1] = "8"
	roundToFractionalDigits[1] = 0
	numberSigns[1] = NumSignVal.Positive()

	expectedIntDigits[2] = "7"
	expectedFracDigits[2] = "4"
	expectedNumStr[2] = "7"
	roundToFractionalDigits[2] = 0
	numberSigns[2] = NumSignVal.Positive()

	expectedIntDigits[3] = "7"
	expectedFracDigits[3] = "4"
	expectedNumStr[3] = "7"
	roundToFractionalDigits[3] = 0
	numberSigns[3] = NumSignVal.Negative()

	expectedIntDigits[4] = "7"
	expectedFracDigits[4] = "5"
	expectedNumStr[4] = "8"
	roundToFractionalDigits[4] = 0
	numberSigns[4] = NumSignVal.Negative()

	expectedIntDigits[5] = "7"
	expectedFracDigits[5] = "6"
	expectedNumStr[5] = "8"
	roundToFractionalDigits[5] = 0
	numberSigns[5] = NumSignVal.Negative()

	expectedIntDigits[6] = "7"
	expectedFracDigits[6] = "1236"
	expectedNumStr[6] = "7.124"
	roundToFractionalDigits[6] = 3
	numberSigns[6] = NumSignVal.Negative()

	expectedIntDigits[7] = "7"
	expectedFracDigits[7] = "1235"
	expectedNumStr[7] = "7.124"
	roundToFractionalDigits[7] = 3
	numberSigns[7] = NumSignVal.Negative()

	expectedIntDigits[8] = "7"
	expectedFracDigits[8] = "1225"
	expectedNumStr[8] = "7.122"
	roundToFractionalDigits[8] = 3
	numberSigns[8] = NumSignVal.Negative()

	expectedIntDigits[9] = "7"
	expectedFracDigits[9] = "1236"
	expectedNumStr[9] = "7.124"
	roundToFractionalDigits[9] = 3
	numberSigns[9] = NumSignVal.Positive()

	expectedIntDigits[10] = "7"
	expectedFracDigits[10] = "1225"
	expectedNumStr[10] = "7.122"
	roundToFractionalDigits[10] = 3
	numberSigns[10] = NumSignVal.Positive()

	expectedIntDigits[11] = "7"
	expectedFracDigits[11] = "1234"
	expectedNumStr[11] = "7.123"
	roundToFractionalDigits[11] = 3
	numberSigns[11] = NumSignVal.Positive()

	expectedIntDigits[12] = "23"
	expectedFracDigits[12] = "5"
	expectedNumStr[12] = "24"
	roundToFractionalDigits[12] = 0
	numberSigns[12] = NumSignVal.Negative()

	expectedIntDigits[13] = "24"
	expectedFracDigits[13] = "5"
	expectedNumStr[13] = "24"
	roundToFractionalDigits[13] = 0
	numberSigns[13] = NumSignVal.Negative()

	expectedIntDigits[14] = "7"
	expectedFracDigits[14] = "123"
	expectedNumStr[14] = "7.123000"
	roundToFractionalDigits[14] = 6
	numberSigns[14] = NumSignVal.Positive()

	expectedIntDigits[15] = "7"
	expectedFracDigits[15] = "123"
	expectedNumStr[15] = "7.123000"
	roundToFractionalDigits[15] = 6
	numberSigns[15] = NumSignVal.Negative()

	var actualStr string

	var intRunes, fracRunes RuneArrayDto
	var err error
	nStrMathRoundAtom := numStrMathRoundingAtom{}

	for i := 0; i < arrayLen; i++ {

		intRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedIntDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"intRunes[%v]=%v",
					i,
					expectedIntDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		fracRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedFracDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"fracRunes[%v]=%v",
					i,
					expectedFracDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		err = nStrMathRoundAtom.roundHalfToEven(
			&intRunes,
			&fracRunes,
			roundToFractionalDigits[i],
			numberSigns[i],
			ePrefix.XCpy(
				fmt.Sprintf(
					"roundToFracDigits[%v]=%v",
					i,
					roundToFractionalDigits[i])))

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

		if expectedNumStr[i] != actualStr {

			t.Errorf("%v - Test #1\n"+
				"Error: nStrMathRoundAtom.roundHalfToEven()\n"+
				"Cycle Number    = '%v'\n"+
				"Integer Base    = '%v'\n"+
				"Fractional Base = '%v'\n"+
				"Round To Digit  = '%v'\n"+
				"Number Sign     = '%v'\n"+
				"Expected String = '%v'\n"+
				"  Actual String = '%v'\n",
				ePrefix.String(),
				i,
				expectedIntDigits[i],
				expectedFracDigits[i],
				roundToFractionalDigits[i],
				numberSigns[i].String(),
				expectedNumStr[i],
				actualStr)

			return
		}

	}

	return
}

func TestNumStrMathRoundingRoundHalfToEven_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfToEven_000200()",
		"")

	// Test # 1
	expectedIntDigits := "1"

	expectedFracDigits := "23456789"

	roundToFractionalDigits := 3

	numberSign := NumSignVal.Negative()

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

	err = nStrMathRoundAtom.roundHalfToEven(
		nil,
		&fracRunes,
		roundToFractionalDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfToEven() Error #1\n"+
			"Expected an error return because 'integerDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfToEven(
		&intRunes,
		nil,
		3,
		numberSign,
		ePrefix.XCpy(
			"Error Test #2"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfToEven() Error #2\n"+
			"Expected an error return because 'fractionalDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfToEven(
		&intRunes,
		&fracRunes,
		-1,
		numberSign,
		ePrefix.XCpy(
			"Error Test #3"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfToEven() Error #2\n"+
			"Expected an error return because 'roundToFractionalDigits' is minus one (-1)\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfToEven(
		&intRunes,
		&fracRunes,
		3,
		NumSignVal.None(),
		ePrefix.XCpy(
			"Error Test #4"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfToEven() Error #2\n"+
			"Expected an error return because 'Number Sign' is 'None'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

}

func TestNumStrMathRoundingRoundHalfToOdd_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfToOdd_000100()",
		"")

	const arrayLen = 16

	var expectedIntDigits, expectedFracDigits, expectedNumStr [arrayLen]string

	var roundToFractionalDigits [arrayLen]int

	var numberSigns [arrayLen]NumericSignValueType

	expectedIntDigits[0] = "7"
	expectedFracDigits[0] = "6"
	expectedNumStr[0] = "8"
	roundToFractionalDigits[0] = 0
	numberSigns[0] = NumSignVal.Positive()

	expectedIntDigits[1] = "7"
	expectedFracDigits[1] = "5"
	expectedNumStr[1] = "7"
	roundToFractionalDigits[1] = 0
	numberSigns[1] = NumSignVal.Positive()

	expectedIntDigits[2] = "7"
	expectedFracDigits[2] = "4"
	expectedNumStr[2] = "7"
	roundToFractionalDigits[2] = 0
	numberSigns[2] = NumSignVal.Positive()

	expectedIntDigits[3] = "7"
	expectedFracDigits[3] = "4"
	expectedNumStr[3] = "7"
	roundToFractionalDigits[3] = 0
	numberSigns[3] = NumSignVal.Negative()

	expectedIntDigits[4] = "7"
	expectedFracDigits[4] = "5"
	expectedNumStr[4] = "7"
	roundToFractionalDigits[4] = 0
	numberSigns[4] = NumSignVal.Negative()

	expectedIntDigits[5] = "7"
	expectedFracDigits[5] = "6"
	expectedNumStr[5] = "8"
	roundToFractionalDigits[5] = 0
	numberSigns[5] = NumSignVal.Negative()

	expectedIntDigits[6] = "7"
	expectedFracDigits[6] = "1236"
	expectedNumStr[6] = "7.124"
	roundToFractionalDigits[6] = 3
	numberSigns[6] = NumSignVal.Negative()

	expectedIntDigits[7] = "7"
	expectedFracDigits[7] = "1235"
	expectedNumStr[7] = "7.123"
	roundToFractionalDigits[7] = 3
	numberSigns[7] = NumSignVal.Negative()

	expectedIntDigits[8] = "7"
	expectedFracDigits[8] = "1225"
	expectedNumStr[8] = "7.123"
	roundToFractionalDigits[8] = 3
	numberSigns[8] = NumSignVal.Negative()

	expectedIntDigits[9] = "7"
	expectedFracDigits[9] = "1236"
	expectedNumStr[9] = "7.124"
	roundToFractionalDigits[9] = 3
	numberSigns[9] = NumSignVal.Positive()

	expectedIntDigits[10] = "7"
	expectedFracDigits[10] = "1225"
	expectedNumStr[10] = "7.123"
	roundToFractionalDigits[10] = 3
	numberSigns[10] = NumSignVal.Positive()

	expectedIntDigits[11] = "7"
	expectedFracDigits[11] = "1234"
	expectedNumStr[11] = "7.123"
	roundToFractionalDigits[11] = 3
	numberSigns[11] = NumSignVal.Positive()

	expectedIntDigits[12] = "23"
	expectedFracDigits[12] = "5"
	expectedNumStr[12] = "23"
	roundToFractionalDigits[12] = 0
	numberSigns[12] = NumSignVal.Negative()

	expectedIntDigits[13] = "24"
	expectedFracDigits[13] = "5"
	expectedNumStr[13] = "25"
	roundToFractionalDigits[13] = 0
	numberSigns[13] = NumSignVal.Negative()

	expectedIntDigits[14] = "7"
	expectedFracDigits[14] = "123"
	expectedNumStr[14] = "7.123000"
	roundToFractionalDigits[14] = 6
	numberSigns[14] = NumSignVal.Positive()

	expectedIntDigits[15] = "7"
	expectedFracDigits[15] = "123"
	expectedNumStr[15] = "7.123000"
	roundToFractionalDigits[15] = 6
	numberSigns[15] = NumSignVal.Negative()

	var actualStr string

	var intRunes, fracRunes RuneArrayDto
	var err error
	nStrMathRoundAtom := numStrMathRoundingAtom{}

	for i := 0; i < arrayLen; i++ {

		intRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedIntDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"intRunes[%v]=%v",
					i,
					expectedIntDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		fracRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedFracDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"fracRunes[%v]=%v",
					i,
					expectedFracDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		err = nStrMathRoundAtom.roundHalfToOdd(
			&intRunes,
			&fracRunes,
			roundToFractionalDigits[i],
			numberSigns[i],
			ePrefix.XCpy(
				fmt.Sprintf(
					"roundToFracDigits[%v]=%v",
					i,
					roundToFractionalDigits[i])))

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

		if expectedNumStr[i] != actualStr {

			t.Errorf("%v - Test #1\n"+
				"Error: nStrMathRoundAtom.roundHalfToOdd()\n"+
				"Cycle Number    = '%v'\n"+
				"Integer Base    = '%v'\n"+
				"Fractional Base = '%v'\n"+
				"Round To Digit  = '%v'\n"+
				"Number Sign     = '%v'\n"+
				"Expected String = '%v'\n"+
				"  Actual String = '%v'\n",
				ePrefix.String(),
				i,
				expectedIntDigits[i],
				expectedFracDigits[i],
				roundToFractionalDigits[i],
				numberSigns[i].String(),
				expectedNumStr[i],
				actualStr)

			return
		}

	}

	return
}

func TestNumStrMathRoundingRoundHalfToOdd_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundHalfToOdd_000200()",
		"")

	// Test # 1
	expectedIntDigits := "1"

	expectedFracDigits := "23456789"

	roundToFractionalDigits := 3

	numberSign := NumSignVal.Negative()

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

	err = nStrMathRoundAtom.roundHalfToOdd(
		nil,
		&fracRunes,
		roundToFractionalDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfToOdd() Error #1\n"+
			"Expected an error return because 'integerDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfToOdd(
		&intRunes,
		nil,
		3,
		numberSign,
		ePrefix.XCpy(
			"Error Test #2"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfToOdd() Error #2\n"+
			"Expected an error return because 'fractionalDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfToOdd(
		&intRunes,
		&fracRunes,
		-1,
		numberSign,
		ePrefix.XCpy(
			"Error Test #3"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfToOdd() Error #2\n"+
			"Expected an error return because 'roundToFractionalDigits' is minus one (-1)\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundHalfToOdd(
		&intRunes,
		&fracRunes,
		3,
		NumSignVal.None(),
		ePrefix.XCpy(
			"Error Test #4"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfToOdd() Error #2\n"+
			"Expected an error return because 'Number Sign' is 'None'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

}

func TestNumStrMathRoundingRoundRandomly_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundRandomly_000100()",
		"")

	const arrayLen = 16

	var expectedIntDigits, expectedFracDigits, expectedNumStr [arrayLen]string

	var roundToFractionalDigits [arrayLen]int

	var numberSigns [arrayLen]NumericSignValueType

	expectedIntDigits[0] = "7"
	expectedFracDigits[0] = "7"
	expectedNumStr[0] = "8"
	roundToFractionalDigits[0] = 0
	numberSigns[0] = NumSignVal.Positive()

	expectedIntDigits[1] = "7"
	expectedFracDigits[1] = "6"
	expectedNumStr[1] = "8"
	roundToFractionalDigits[1] = 0
	numberSigns[1] = NumSignVal.Positive()

	expectedIntDigits[2] = "7"
	expectedFracDigits[2] = "4"
	expectedNumStr[2] = "7"
	roundToFractionalDigits[2] = 0
	numberSigns[2] = NumSignVal.Positive()

	expectedIntDigits[3] = "7"
	expectedFracDigits[3] = "4"
	expectedNumStr[3] = "7"
	roundToFractionalDigits[3] = 0
	numberSigns[3] = NumSignVal.Negative()

	expectedIntDigits[4] = "7"
	expectedFracDigits[4] = "6"
	expectedNumStr[4] = "8"
	roundToFractionalDigits[4] = 0
	numberSigns[4] = NumSignVal.Negative()

	expectedIntDigits[5] = "7"
	expectedFracDigits[5] = "7"
	expectedNumStr[5] = "8"
	roundToFractionalDigits[5] = 0
	numberSigns[5] = NumSignVal.Negative()

	expectedIntDigits[6] = "7"
	expectedFracDigits[6] = "1236"
	expectedNumStr[6] = "7.124"
	roundToFractionalDigits[6] = 3
	numberSigns[6] = NumSignVal.Negative()

	expectedIntDigits[7] = "7"
	expectedFracDigits[7] = "1234"
	expectedNumStr[7] = "7.123"
	roundToFractionalDigits[7] = 3
	numberSigns[7] = NumSignVal.Negative()

	expectedIntDigits[8] = "7"
	expectedFracDigits[8] = "1226"
	expectedNumStr[8] = "7.123"
	roundToFractionalDigits[8] = 3
	numberSigns[8] = NumSignVal.Negative()

	expectedIntDigits[9] = "7"
	expectedFracDigits[9] = "1237"
	expectedNumStr[9] = "7.124"
	roundToFractionalDigits[9] = 3
	numberSigns[9] = NumSignVal.Positive()

	expectedIntDigits[10] = "7"
	expectedFracDigits[10] = "1224"
	expectedNumStr[10] = "7.122"
	roundToFractionalDigits[10] = 3
	numberSigns[10] = NumSignVal.Positive()

	expectedIntDigits[11] = "7"
	expectedFracDigits[11] = "1233"
	expectedNumStr[11] = "7.123"
	roundToFractionalDigits[11] = 3
	numberSigns[11] = NumSignVal.Positive()

	expectedIntDigits[12] = "23"
	expectedFracDigits[12] = "6"
	expectedNumStr[12] = "24"
	roundToFractionalDigits[12] = 0
	numberSigns[12] = NumSignVal.Negative()

	expectedIntDigits[13] = "24"
	expectedFracDigits[13] = "4"
	expectedNumStr[13] = "24"
	roundToFractionalDigits[13] = 0
	numberSigns[13] = NumSignVal.Negative()

	expectedIntDigits[14] = "7"
	expectedFracDigits[14] = "123"
	expectedNumStr[14] = "7.123000"
	roundToFractionalDigits[14] = 6
	numberSigns[14] = NumSignVal.Positive()

	expectedIntDigits[15] = "7"
	expectedFracDigits[15] = "123"
	expectedNumStr[15] = "7.123000"
	roundToFractionalDigits[15] = 6
	numberSigns[15] = NumSignVal.Negative()

	var actualStr string

	var intRunes, fracRunes RuneArrayDto
	var err error
	nStrMathRoundAtom := numStrMathRoundingAtom{}

	for i := 0; i < arrayLen; i++ {

		intRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedIntDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"intRunes[%v]=%v",
					i,
					expectedIntDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		fracRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(expectedFracDigits[i]),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"fracRunes[%v]=%v",
					i,
					expectedFracDigits[i])))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		err = nStrMathRoundAtom.roundRandomly(
			&intRunes,
			&fracRunes,
			roundToFractionalDigits[i],
			numberSigns[i],
			ePrefix.XCpy(
				fmt.Sprintf(
					"roundToFracDigits[%v]=%v",
					i,
					roundToFractionalDigits[i])))

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

		if expectedNumStr[i] != actualStr {

			t.Errorf("%v - Test #1\n"+
				"Error: nStrMathRoundAtom.roundRandomly()\n"+
				"Cycle Number    = '%v'\n"+
				"Integer Base    = '%v'\n"+
				"Fractional Base = '%v'\n"+
				"Round To Digit  = '%v'\n"+
				"Number Sign     = '%v'\n"+
				"Expected String = '%v'\n"+
				"  Actual String = '%v'\n",
				ePrefix.String(),
				i,
				expectedIntDigits[i],
				expectedFracDigits[i],
				roundToFractionalDigits[i],
				numberSigns[i].String(),
				expectedNumStr[i],
				actualStr)

			return
		}

	}

	intVal := 600
	var fracStr string

	numberSign := NumSignVal.Positive()
	roundAtFractionalDigits := 2

	for j := 0; j < 500; j++ {

		intVal += j + 57

		intRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(fmt.Sprintf("%v", intVal)),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"intVal[%v]=%v",
					j,
					intVal)))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		fracStr = fmt.Sprintf("%v", j)

		fracStr += "5"

		fracRunes,
			err = RuneArrayDto{}.NewRunes(
			[]rune(fracStr),
			CharSearchType.LinearTargetStartingIndex(),
			ePrefix.XCpy(
				fmt.Sprintf(
					"fracRunes[%v]=%v",
					j,
					fracStr)))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

		if j%2 == 0 {
			numberSign = NumSignVal.Positive()
		} else {
			numberSign = NumSignVal.Negative()
		}

		roundAtFractionalDigits = len(fracStr) - 1

		err = nStrMathRoundAtom.roundRandomly(
			&intRunes,
			&fracRunes,
			roundAtFractionalDigits,
			numberSign,
			ePrefix.XCpy(
				fmt.Sprintf(
					"roundAtFractionalDigits[%v]=%v",
					j,
					roundAtFractionalDigits)))

		if err != nil {
			t.Errorf("%v\n",
				err.Error())
			return
		}

	}

	return
}

func TestNumStrMathRoundingRoundRandomly_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathRoundingRoundRandomly_000200()",
		"")

	// Test # 1
	expectedIntDigits := "1"

	expectedFracDigits := "23456789"

	roundToFractionalDigits := 3

	numberSign := NumSignVal.Negative()

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

	err = nStrMathRoundAtom.roundRandomly(
		nil,
		&fracRunes,
		roundToFractionalDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf("roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundRandomly() Error #1\n"+
			"Expected an error return because 'integerDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundRandomly(
		&intRunes,
		nil,
		3,
		numberSign,
		ePrefix.XCpy(
			"Error Test #2"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundRandomly() Error #2\n"+
			"Expected an error return because 'fractionalDigits' is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundRandomly(
		&intRunes,
		&fracRunes,
		-1,
		numberSign,
		ePrefix.XCpy(
			"Error Test #3"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundRandomly() Error #2\n"+
			"Expected an error return because 'roundToFractionalDigits' is minus one (-1)\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	err = nStrMathRoundAtom.roundRandomly(
		&intRunes,
		&fracRunes,
		3,
		NumSignVal.None(),
		ePrefix.XCpy(
			"Error Test #4"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: nStrMathRoundAtom.roundRandomly() Error #2\n"+
			"Expected an error return because 'Number Sign' is 'None'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

}

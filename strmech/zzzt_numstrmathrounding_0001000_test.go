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

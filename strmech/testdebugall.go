package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"sync"
)

type TestDebugAll struct {
	lock *sync.Mutex
}

func (testDebug *TestDebugAll) MathRounding(
	errorPrefix interface{}) {

	if testDebug.lock == nil {
		testDebug.lock = new(sync.Mutex)
	}

	testDebug.lock.Lock()

	defer testDebug.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDebugAll."+
			"MathRounding()",
		"")

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	expectedIntDigits := "7"

	expectedFracDigits := "5"

	roundToFractionalDigits := 0

	expectedStr := "7"

	numberSign := NumSignVal.Negative()

	var intRunes RuneArrayDto

	intRunes,
		err = RuneArrayDto{}.NewRunes(
		[]rune(expectedIntDigits),
		CharSearchType.LinearTargetStartingIndex(),
		ePrefix.XCpy(
			"intRunes"))

	if err != nil {

		fmt.Printf("%v\n",
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

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	nStrMathRoundAtom := numStrMathRoundingAtom{}

	err = nStrMathRoundAtom.roundHalfTowardsZero(
		&intRunes,
		&fracRunes,
		roundToFractionalDigits,
		numberSign,
		ePrefix.XCpy(
			fmt.Sprintf("Test #5 roundToFracDigits=%v",
				roundToFractionalDigits)))

	if err != nil {

		fmt.Printf("%v\n",
			err.Error())

		return
	}

	actualStr := intRunes.GetCharacterString()

	if fracRunes.GetRuneArrayLength() > 0 {

		actualStr += "." +
			fracRunes.GetCharacterString()

	}

	if expectedStr != actualStr {

		fmt.Printf("%v\n"+
			"Error: nStrMathRoundAtom.roundHalfAwayFromZero() #5\n"+
			"Expected String = '%v'\n"+
			"  Actual String = '%v'\n",
			ePrefix.String(),
			expectedStr,
			actualStr)

		return
	}

	fmt.Printf("%v\n"+
		"No Errors!\n"+
		"Expected String = '%v'\n"+
		"  Actual String = '%v'\n",
		ePrefix.String(),
		expectedStr,
		actualStr)

	return
}

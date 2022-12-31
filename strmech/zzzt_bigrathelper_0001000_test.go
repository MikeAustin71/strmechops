package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"testing"
)

func TestMathBigRatHelper_RatToBigFloat_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathBigRatHelper_RatToBigFloat_000100()",
		"")

	newRat := big.NewRat(1, 3)

	expectedBigFloatNum := "0.3333333333"

	var bigFloatNum *big.Float

	var err error

	bigFloatNum,
		err = new(MathBigRatHelper).BigRatToBigFloat(
		newRat,
		10,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualBigFloatNum := bigFloatNum.Text('f', -1)

	if actualBigFloatNum != expectedBigFloatNum {

		t.Errorf("\n%v\n"+
			"Test#1\n"+
			"Error: actualBigFloatNum != expectedBigFloatNum\n"+
			"actualBigFloatNum   = '%v'\n"+
			"expectedBigFloatNum = '%v'\n",
			ePrefix.String(),
			actualBigFloatNum,
			expectedBigFloatNum)

		return

	}

	return
}

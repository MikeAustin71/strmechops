package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"math/big"
	"strings"
)

type MainTest03 struct {
	input string
}

func (MainTest03) RoundBigFloat01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.RoundBigFloat01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	requiredDigits := int64(150)

	floatHelper := strmech.MathFloatHelper{}

	requiredPrecision,
		err := floatHelper.
		DigitsToPrecisionEstimate(
			requiredDigits,
			ePrefix.XCpy(
				""))

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	if requiredPrecision != 504 {

		fmt.Printf("\n%v\n"+
			"Error: Expected 504 bits of precision from "+
			"150 digits.\n"+
			"Instead, requiredPrecision = '%v'\n",
			ePrefix.String(),
			requiredPrecision)

		return
	}

	test01Float :=
		new(big.Float).
			SetPrec(requiredPrecision).
			SetMode(big.AwayFromZero).
			SetInt64(0)

	test02Float :=
		new(big.Float).
			SetPrec(requiredPrecision).
			SetMode(big.AwayFromZero).
			SetInt64(0)

	originalValueStr := "123456.78125"

	var ok bool

	_,
		ok = test01Float.SetString(originalValueStr)

	if !ok {

		fmt.Printf("\n%v\n"+
			"Error:\n"+
			"test01Float.SetString(originalValueStr) FAILED!\n"+
			"originalValueStr = '%v'\n",
			ePrefix.String(),
			originalValueStr)

		return
	}

	err = floatHelper.RoundBigFloat(
		test01Float,
		test02Float,
		strmech.NumRoundType.HalfAwayFromZero(),
		4,
		ePrefix.XCpy(
			"test02Float<- to 4-Frac digits"))

	actualRoundedValueStr :=
		test02Float.Text('f', 5)

	expectedRoundedValueStr := "123456.78130"

	if actualRoundedValueStr != expectedRoundedValueStr {

		fmt.Printf("\n%v\n"+
			"Error: originalValueStr = '%v'\n"+
			"Expected a rounded value      = '%v'\n"+
			"Instead, actual rounded value = '%v'\n\n",
			ePrefix.String(),
			originalValueStr,
			expectedRoundedValueStr,
			actualRoundedValueStr)

		return
	}

	fmt.Printf("\n\n%v\n"+
		"Successful Rounding Operation!\n"+
		"originalValueStr = '%v'\n"+
		"Expected a rounded value of   = '%v'\n"+
		"Matching Actual rounded value = '%v'\n\n",
		ePrefix.String(),
		originalValueStr,
		expectedRoundedValueStr,
		actualRoundedValueStr)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	return
}

func (MainTest03) EstimatePrecisionBits01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.EstimatePrecisionBits01()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	requiredDigits := int64(150)

	floatHelper := strmech.MathFloatHelper{}

	requiredPrecision,
		err := floatHelper.
		DigitsToPrecisionEstimate(
			requiredDigits,
			ePrefix.XCpy(
				""))

	if err != nil {
		fmt.Printf("%v\n",
			err.Error())
		return
	}

	if requiredPrecision != 504 {

		fmt.Printf("\n%v\n"+
			"Error: Expected 504 bits of precision from "+
			"150 digits.\n"+
			"Instead, requiredPrecision = '%v'\n",
			ePrefix.String(),
			requiredPrecision)

		return
	}

	fmt.Printf("requiredPrecision = %v\n"+
		"expected required precision = %v\n"+
		"SUCCESS!\n\n",
		requiredPrecision,
		504)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

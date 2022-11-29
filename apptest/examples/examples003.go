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

func (MainTest03) RaiseToExponent02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.RaiseToExponent02()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	// int digits 4
	// frac digits 12
	baseStr := "5084.987654321000"

	exponent := int64(4)

	// int digits = 15
	// frac digits = 17
	expectedResultStr := "668589591687777.75101222860206783"
	expectedFracDigits := 17

	floatHelper := strmech.MathFloatHelper{}

	var bFloatDto strmech.BigFloatDto
	var err error

	bFloatDto,
		err = floatHelper.BigFloatFromPureNumStr(
		baseStr,
		".",
		true,
		10,
		0,
		ePrefix)

	fmt.Printf("\nInitialization\n"+
		"----- RaiseToIntPositiveExponent() -----\n"+
		"baseStr        = %v\n"+
		"baseFloat      = %v\n"+
		"base Precision = %v\n"+
		"base Mode		= %v\n"+
		"base Accuracy  = %v\n"+
		"exponent       = %v\n\n",
		baseStr,
		bFloatDto.Value.Text('f', 12),
		bFloatDto.Value.Prec(),
		bFloatDto.Value.Mode(),
		bFloatDto.Value.Acc(),
		exponent)

	var raisedToIntExponent *big.Float

	raisedToIntExponent,
		err = floatHelper.RaiseToIntPositiveExponent(
		&bFloatDto.Value,
		exponent,
		ePrefix.XCpy(
			"raisedToIntExponent"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	fmt.Printf("After Calculation Of Raise To Power\n"+
		"----- RaiseToIntPositiveExponent() -----\n"+
		"Actual raisedToIntExponent    = %v\n"+
		"Expected raisedToIntExponent  = %v\n"+
		"raisedToIntExponent Precision = %v\n"+
		"raisedToIntExponent Mode      = %v\n"+
		"raisedToIntExponent Accuracy  = %v\n"+
		"exponent                   = %v\n\n\n",
		raisedToIntExponent.Text('f', expectedFracDigits),
		expectedResultStr,
		raisedToIntExponent.Prec(),
		raisedToIntExponent.Mode(),
		raisedToIntExponent.Acc(),
		exponent)

	fmt.Printf("\nInitialization\n"+
		"----- RaiseToFloatPositiveExponent() -----\n"+
		"baseStr        = %v\n"+
		"baseFloat      = %v\n"+
		"base Precision = %v\n"+
		"base Mode		= %v\n"+
		"base Accuracy  = %v\n"+
		"exponent       = %v\n\n\n",
		baseStr,
		bFloatDto.Value.Text('f', -1),
		bFloatDto.Value.Prec(),
		bFloatDto.Value.Mode(),
		bFloatDto.Value.Acc(),
		exponent)

	var raisedToFloatExponent *big.Float

	raisedToFloatExponent,
		err = floatHelper.RaiseToFloatPositiveExponent(
		&bFloatDto.Value,
		exponent,
		200,
		0,
		ePrefix)

	fmt.Printf("After Calculation Of Raise To Power\n"+
		"----- RaiseToFloatPositiveExponent() -----\n"+
		"Actual raisedToFloatExponent -1 = %v\n"+
		"Actual raisedToFloatExponent    = %v\n"+
		"Expected raisedToFloatExponent  = %v\n"+
		"raisedToFloatExponent Precision = %v\n"+
		"raisedToFloatExponent Mode      = %v\n"+
		"raisedToFloatExponent Accuracy  = %v\n"+
		"exponent                   = %v\n\n",
		raisedToFloatExponent.Text('f', -1),
		raisedToFloatExponent.Text('f', expectedFracDigits),
		expectedResultStr,
		raisedToFloatExponent.Prec(),
		raisedToFloatExponent.Mode(),
		raisedToFloatExponent.Acc(),
		exponent)

	fmt.Printf("Results Comparison\n"+
		"RaiseToIntPositiveExponent vs. RaiseToFloatPositiveExponent\n"+
		"Actual raisedToFloatExponent -1 = %v\n"+
		"Actual raisedToIntExponent -1   = %v\n"+
		"Actual raisedToFloatExponent    = %v\n"+
		"Actual raisedToIntExponent      = %v\n"+
		"Expected raisedToIntExponent    = %v\n"+
		"raisedToIntExponent Precision   = %v\n"+
		"raisedToIntExponent Mode        = %v\n"+
		"raisedToIntExponent Accuracy    = %v\n"+
		"raisedToFloatExponent Precision = %v\n"+
		"raisedToFloatExponent Mode      = %v\n"+
		"raisedToFloatExponent Accuracy  = %v\n"+
		"exponent                        = %v\n\n",
		raisedToFloatExponent.Text('f', -1),
		raisedToIntExponent.Text('f', -1),
		raisedToFloatExponent.Text('f', expectedFracDigits),
		raisedToIntExponent.Text('f', expectedFracDigits),
		expectedResultStr,
		raisedToIntExponent.Prec(),
		raisedToIntExponent.Mode(),
		raisedToIntExponent.Acc(),
		raisedToFloatExponent.Prec(),
		raisedToFloatExponent.Mode(),
		raisedToFloatExponent.Acc(),
		exponent)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (MainTest03) RaiseToExponent01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest03.RaiseToExponent01()",
		"")

	breakStr := strings.Repeat("=", 200)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	// int digits 4
	// frac digits 12
	baseStr := "5084.987654321000"
	//baseStr := "2.24"

	exponent := int64(4)

	// int digits = 15
	// frac digits = 17
	expectedResultStr := "668589591687777.75101222860206783"
	//expectedResultStr := "25.17630976"
	expectedFracDigits := 17

	floatHelper := strmech.MathFloatHelper{}

	var bFloatDto strmech.BigFloatDto
	var err error

	bFloatDto,
		err = floatHelper.BigFloatFromPureNumStr(
		baseStr,
		".",
		true,
		10,
		0,
		ePrefix)

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	fmt.Printf("\nInitialization\n"+
		"baseStr        = %v\n"+
		"baseFloat      = %v\n"+
		"base Precision = %v\n"+
		"base Accuracy  = %v\n"+
		"base Round Mode= %v\n"+
		"exponent       = %v\n\n",
		baseStr,
		bFloatDto.Value.Text('f', -1),
		bFloatDto.Value.Prec(),
		bFloatDto.Value.Acc(),
		bFloatDto.Value.Mode(),
		exponent)

	var raisedToExponent *big.Float

	raisedToExponent,
		err = floatHelper.RaiseToFloatPositiveExponent(
		&bFloatDto.Value,
		exponent,
		200,
		0,
		ePrefix.XCpy(
			"raisedToExponent"))

	if err != nil {
		fmt.Printf("\n%v\n",
			err.Error())
		return
	}

	raisedToExponentStr :=
		raisedToExponent.Text('f', expectedFracDigits)

	fmt.Printf("After Calculation Of Raise To Power\n"+
		"Actual raisedToExponent    = %v\n"+
		"Expected raisedToExponent  = %v\n"+
		"raisedToExponent Precision = %v\n"+
		"raisedToExponent Accuracy  = %v\n"+
		"raisedToExponent Mode      = %v\n",
		raisedToExponentStr,
		expectedResultStr,
		raisedToExponent.Prec(),
		raisedToExponent.Acc(),
		raisedToExponent.Mode())

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

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

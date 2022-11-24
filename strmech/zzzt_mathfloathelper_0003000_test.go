package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"testing"
)

func TestMathFloatHelper_RoundBigFloat_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathFloatHelper_RoundBigFloat_000100()",
		"")

	requiredDigits := int64(150)

	floatHelper := MathFloatHelper{}

	requiredPrecision,
		err := floatHelper.
		DigitsToPrecisionEstimate(
			requiredDigits,
			ePrefix.XCpy(
				""))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if requiredPrecision != 504 {

		t.Errorf("\n%v\n"+
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

		t.Errorf("\n%v\n"+
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
		NumRoundType.HalfAwayFromZero(),
		4,
		ePrefix.XCpy(
			"test02Float<- to 4-Frac digits"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualRoundedValueStr :=
		test02Float.Text('f', 5)

	expectedRoundedValueStr := "123456.78130"

	if actualRoundedValueStr != expectedRoundedValueStr {

		t.Errorf("\n%v\n"+
			"Error: originalValueStr       = '%v'\n"+
			"Expected a rounded value      = '%v'\n"+
			"Instead, actual rounded value = '%v'\n\n",
			ePrefix.String(),
			originalValueStr,
			expectedRoundedValueStr,
			actualRoundedValueStr)

		return
	}

	return
}

func TestMathFloatHelper_PiTo20k_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathFloatHelper_PiTo20k_000100()",
		"")

	roundToFracDigits := 12

	piTo20kFloat,
		err := new(MathFloatHelper).
		PiTo20k(
			NumRoundType.HalfAwayFromZero(),
			roundToFracDigits,
			ePrefix.XCpy(
				""))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	// 3.1415926535897932384626433832795028841971693993751
	expectedPiStr := "3.141592653590"

	actualPiStr := piTo20kFloat.Text('f', roundToFracDigits)

	if expectedPiStr != actualPiStr {

		t.Errorf("\n%v\n"+
			"Pi to 50 digits: 3.1415926535897 932384626433832795028841971693993751"+
			"Error:\n"+
			"Expected Pi value        = '%v'\n"+
			"Instead, actual Pi value = '%v'\n\n",
			ePrefix.String(),
			expectedPiStr,
			actualPiStr)

		return
	}

	return
}

func TestMathFloatHelper_PrecisionToDigitsEstimate_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathFloatHelper_PrecisionToDigitsEstimate_000100()",
		"")

	originalPrecisionBits := uint(504)

	expectedNumberOfDigits := int64(150)

	actualNumOfNumericalDigits,
		err := new(MathFloatHelper).PrecisionToDigitsEstimate(
		originalPrecisionBits,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	maxNumOfDigits := expectedNumberOfDigits + 3
	minNumOfDigits := expectedNumberOfDigits - 3

	if actualNumOfNumericalDigits > maxNumOfDigits &&
		actualNumOfNumericalDigits < minNumOfDigits {

		t.Errorf("\n%v\n"+
			"Error: originalPrecisionBits         = '%v'\n"+
			"Expected number of numerical digits  = '%v'\n"+
			"Actual number of numerical digits    = '%v'\n"+
			"Remember there is a margin of error of plus\n"+
			"or minus three (+ or - 3)\n\n",
			ePrefix.String(),
			originalPrecisionBits,
			expectedNumberOfDigits,
			actualNumOfNumericalDigits)

	}

	return
}

func TestMathFloatHelper_PrecisionBitsFromRequiredDigits_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathFloatHelper_PrecisionBitsFromRequiredDigits_000100()",
		"")

	testStr01 := "9586000000000.87912345678901234567890000000000000000000000000000"

	idx := strings.Index(testStr01, ".")

	if idx == -1 {
		t.Errorf("\n%v\n"+
			"Error: Failed to find decimal point in testStr01!\n",
			ePrefix)

		return
	}

	lenTestStr := len(testStr01)

	var requiredIntegerDigits,
		requiredFractionalDigits,
		requestedBufferDigits int64

	requiredIntegerDigits = int64(idx)
	requiredFractionalDigits = int64(lenTestStr-1) - requiredIntegerDigits
	requestedBufferDigits = 100

	actualPrecision,
		err := new(MathFloatHelper).PrecisionBitsFromRequiredDigits(
		requiredIntegerDigits,
		requiredFractionalDigits,
		requestedBufferDigits,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedPrecision := uint(544)

	if actualPrecision != expectedPrecision {

		t.Errorf("\n%v\n"+
			"Error: Expected Does NOT Match Actual Precision Bits!\n"+
			"Expected Precision Bits = '%v'\n"+
			"Actual Precision Bits   = '%v'\n",
			ePrefix.String(),
			expectedPrecision,
			actualPrecision)

		return

	}

	var ok bool

	test01Float,
		ok :=
		new(big.Float).
			SetMode(big.AwayFromZero).
			SetInt64(0).SetString(
			testStr01)

	if !ok {

		t.Errorf("\n%v\n"+
			"Error:\n"+
			"test02Float.SetString(testStr01) FAILED!\n"+
			"testStr01 = '%v'\n",
			ePrefix.String(),
			testStr01)

		return
	}

	resultsStr01 :=
		test01Float.Text('f', int(requiredFractionalDigits))

	var test02Float *big.Float

	test02Float,
		ok =
		new(big.Float).
			SetMode(big.AwayFromZero).
			SetInt64(0).SetString(
			testStr01)

	if !ok {

		t.Errorf("\n%v\n"+
			"Error:\n"+
			"test02Float.SetString(testStr01) FAILED!\n"+
			"testStr01 = '%v'\n",
			ePrefix.String(),
			testStr01)

		return
	}

	test02Float.SetPrec(actualPrecision)

	resultsStr02 :=
		test02Float.Text('f', int(requiredFractionalDigits))

	if resultsStr01 != resultsStr02 {

		t.Errorf("\n%v\n"+
			"Error: Expected output string Does NOT Match Actual output string!\n"+
			"resultsStr01 was configured with original Precision   = %v\n"+
			"resultsStr02 was configured with calculated Precision = %v\n"+
			"resultsStr01 = '%v'\n"+
			"resultsStr02  = '%v'\n",
			ePrefix.String(),
			test01Float.Prec(),
			actualPrecision,
			resultsStr01,
			resultsStr02)

		return

	}

	return
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"testing"
)

func getRaiseToExponentTestDataSeries01() (
	baseStrs []string,
	exponents []int64,
	expectedResults []string,
	expectedResultFracDigits []int) {

	// ------------------------------------------
	// # 0
	baseStrs = append(
		baseStrs, "5084.987654321000")

	exponents = append(
		exponents, 4)

	expectedResults = append(
		expectedResults, "668589591687777.75101222860206783")

	// ------------------------------------------
	// # 1
	baseStrs = append(
		baseStrs, "32")

	exponents = append(
		exponents, 8)

	expectedResults = append(
		expectedResults, "1099511627776")

	// ------------------------------------------
	// # 2
	baseStrs = append(
		baseStrs, "0.0257")

	exponents = append(
		exponents, 7)

	expectedResults = append(
		expectedResults, "0.0000000000074051159531521793")

	// ------------------------------------------
	// # 3
	baseStrs = append(
		baseStrs, "-51.264")

	exponents = append(
		exponents, 12)

	expectedResults = append(
		expectedResults, "329419937773680086383.95166647701")

	// ------------------------------------------
	// # 4
	baseStrs = append(
		baseStrs, "-51.264")

	exponents = append(
		exponents, 12)

	expectedResults = append(
		expectedResults, "329419937773680086383.95166647701")

	// ------------------------------------------
	// # 5
	baseStrs = append(
		baseStrs, "8.045678917")

	exponents = append(
		exponents, 28)

	expectedResults = append(
		expectedResults, "22685873975896914224185022.181615")

	// ------------------------------------------
	// # 6
	baseStrs = append(
		baseStrs, "-14.239")

	exponents = append(
		exponents, 11)

	expectedResults = append(
		expectedResults, "-4878366112834.239885577386540692")

	// ------------------------------------------
	// # 7
	baseStrs = append(
		baseStrs, "0.2031004789618437")

	exponents = append(
		exponents, 4)

	expectedResults = append(
		expectedResults, "0.00170154637578694995110718021402")

	// ------------------------------------------
	// # 8
	baseStrs = append(
		baseStrs, "0.0587243")

	exponents = append(
		exponents, 3)

	expectedResults = append(
		expectedResults, "0.000202513297800637907")

	// ------------------------------------------
	// # 9
	baseStrs = append(
		baseStrs, "76.234567991")

	exponents = append(
		exponents, 10)

	expectedResults = append(
		expectedResults, "6630090181354170396.3198929680763")

	// ------------------------------------------
	// # 10
	baseStrs = append(
		baseStrs, "-16.7653333333218")

	exponents = append(
		exponents, 15)

	expectedResults = append(
		expectedResults, "-2323715149200499897.1442509387987")

	// ------------------------------------------
	// # 11
	baseStrs = append(
		baseStrs, "-16.7653333333218")

	exponents = append(
		exponents, 0)

	expectedResults = append(
		expectedResults, "1")

	// ------------------------------------------
	// # 12
	baseStrs = append(
		baseStrs, "-16.7653333333218")

	exponents = append(
		exponents, 1)

	expectedResults = append(
		expectedResults, "-16.7653333333218")

	// ------------------------------------------
	// # 13
	baseStrs = append(
		baseStrs, "8.324159")

	exponents = append(
		exponents, 11)

	expectedResults = append(
		expectedResults, "13296705133.607966468215693356424")

	// ------------------------------------------
	// # 14
	baseStrs = append(
		baseStrs, "51.35698745233")

	exponents = append(
		exponents, 7)

	expectedResults = append(
		expectedResults, "942316474194.46866272292874238307")

	// ------------------------------------------
	// # 15
	baseStrs = append(
		baseStrs, "-1.012597111")

	exponents = append(
		exponents, 14)

	expectedResults = append(
		expectedResults, "1.1915535798808993682549439073204")

	// =====================================================

	lenExpectedResults := len(expectedResults)

	var idx, lenStr int

	for i := 0; i < lenExpectedResults; i++ {

		idx = strings.Index(expectedResults[i], ".")

		lenStr = len(expectedResults[i])

		if idx == -1 ||
			(lenStr-1) == idx {

			expectedResultFracDigits =
				append(expectedResultFracDigits, 0)

		} else {
			// Len = 6
			// 012345
			// 123.45

			expectedResultFracDigits =
				append(expectedResultFracDigits,
					lenStr-(idx+1))

		}

	}

	return baseStrs,
		exponents,
		expectedResults,
		expectedResultFracDigits
}

func TestMathFloatHelper_RaiseToFloatPositiveExponent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathFloatHelper_RaiseToFloatPositiveExponent_000100()",
		"")

	// int digits 4
	// frac digits 12
	baseStr := "5084.987654321000"

	exponent := int64(4)

	// int digits = 15
	// frac digits = 17
	expectedResultStr := "668589591687777.75101222860206783"
	expectedFracDigits := 17

	floatHelper := MathFloatHelper{}

	var bFloatDto BigFloatDto
	var err error

	bFloatDto,
		err = floatHelper.BigFloatFromPureNumStr(
		baseStr,
		".",
		true,
		100,
		0,
		big.AwayFromZero,
		&ePrefix)

	var raisedToExponent *big.Float

	raisedToExponent,
		err = floatHelper.RaiseToFloatPositiveExponent(
		&bFloatDto.Value,
		exponent,
		200,
		0,
		big.AwayFromZero,
		ePrefix.XCpy(
			"raisedToExponent"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	raisedToExponentStr :=
		raisedToExponent.Text('f', expectedFracDigits)

	if expectedResultStr != raisedToExponentStr {

		t.Errorf("\n%v\n"+
			"Test Series # 1\n"+
			"floatHelper.RaiseToFloatPositiveExponent()\n"+
			"Error: Expected Result Does NOT Match Actual Result!\n"+
			"Expected Result = '%v'\n"+
			"Actual Result   = '%v'\n",
			ePrefix.String(),
			expectedResultStr,
			raisedToExponentStr)

		return

	}

	exponent = 0

	raisedToExponent,
		err = floatHelper.RaiseToFloatPositiveExponent(
		&bFloatDto.Value,
		exponent,
		200,
		0,
		big.AwayFromZero,
		ePrefix.XCpy(
			"raisedToExponent"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	raisedToExponentStr =
		raisedToExponent.Text('f', 4)

	expectedResultStr = "1.0000"

	if expectedResultStr != raisedToExponentStr {

		t.Errorf("\n%v\n"+
			"Test Series # 2\n"+
			"floatHelper.RaiseToFloatPositiveExponent()\n"+
			"Error: Expected Result Does NOT Match Actual Result!\n"+
			"Expected Result = '%v'\n"+
			"Actual Result   = '%v'\n",
			ePrefix.String(),
			expectedResultStr,
			raisedToExponentStr)

		return

	}

	exponent = 1

	raisedToExponent,
		err = floatHelper.RaiseToFloatPositiveExponent(
		&bFloatDto.Value,
		exponent,
		200,
		0,
		big.AwayFromZero,
		ePrefix.XCpy(
			"raisedToExponent"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	raisedToExponentStr =
		raisedToExponent.Text('f', 9)

	expectedResultStr = "5084.987654321"

	if expectedResultStr != raisedToExponentStr {

		t.Errorf("\n%v\n"+
			"Test Series # 3\n"+
			"floatHelper.RaiseToFloatPositiveExponent()\n"+
			"Error: Expected Result Does NOT Match Actual Result!\n"+
			"Expected Result = '%v'\n"+
			"Actual Result   = '%v'\n",
			ePrefix.String(),
			expectedResultStr,
			raisedToExponentStr)

		return

	}

	return
}

func TestMathFloatHelper_RaiseToIntPositiveExponent_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathFloatHelper_RaiseToFloatPositiveExponent_000100()",
		"")

	baseStrs,
		exponents,
		expectedResults,
		expectedResultFracDigits := getRaiseToExponentTestDataSeries01()

	floatHelper := MathFloatHelper{}
	var bFloatDto BigFloatDto
	var err error
	var raisedToExponent *big.Float
	var raisedToExponentStr string
	var numOfExtraDigitsBuffer int64
	var precisionBitsOverride uint

	numOfExtraDigitsBuffer = 50
	precisionBitsOverride = 0

	lenBaseStrs := len(baseStrs)

	for i := 0; i < lenBaseStrs; i++ {

		bFloatDto,
			err = floatHelper.BigFloatFromPureNumStr(
			baseStrs[i],
			".",
			true,
			numOfExtraDigitsBuffer,
			precisionBitsOverride,
			big.AwayFromZero,
			&ePrefix)

		raisedToExponent,
			err = floatHelper.RaiseToFloatPositiveExponent(
			&bFloatDto.Value,
			exponents[i],
			numOfExtraDigitsBuffer,
			precisionBitsOverride,
			big.AwayFromZero,
			ePrefix.XCpy(
				fmt.Sprintf("raisedToExponent<- index=%v",
					i)))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		raisedToExponentStr =
			raisedToExponent.Text('f', expectedResultFracDigits[i])

		if expectedResults[i] != raisedToExponentStr {

			t.Errorf("\n%v\n"+
				"Test Series Index = %v\n"+
				"floatHelper.RaiseToFloatPositiveExponent()\n"+
				"Error: Expected Result Does NOT Match Actual Result!\n"+
				"Expected Result = '%v'\n"+
				"Actual Result   = '%v'\n"+
				"Base Str        = '%v'\n"+
				"Exponent        = '%v'\n",
				ePrefix.String(),
				i,
				expectedResults[i],
				raisedToExponentStr,
				baseStrs[i],
				exponents[i])

			return

		}

	}

	return
}

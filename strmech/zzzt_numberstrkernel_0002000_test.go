package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strings"
	"testing"
)

func TestNumberStrKernel_RoundNoRounding_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundNoRounding_000100",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "7.5"

	var err error
	var baseVal, nStr01, nStr02 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if baseVal.IsZeroValue() {

		t.Errorf("%v\n"+
			"Setup Test#1 - baseVal.IsZeroValue()\n"+
			"Error: IsZeroValue returned 'true'.\n"+
			"However baseVal = '%v'\n",
			ePrefix.String(),
			expectedNumberStr)

		return
	}

	if !baseVal.IsFloatingPointValue() {

		t.Errorf("%v\n"+
			"Setup Test#2 - baseVal.IsFloatingPointValue()\n"+
			"Error: IsZeroValue returned 'false'.\n"+
			"However baseVal = '%v'\n",
			ePrefix.String(),
			expectedNumberStr)

		return
	}

	var intNumSign = 0

	intNumSign,
		err = baseVal.GetNumberSignAsInt(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if intNumSign != 1 {

		t.Errorf("%v\n"+
			"Setup Test#3 - baseVal.GetNumberSignAsInt()\n"+
			"Error: Integer Number Sign should be '1'.\n"+
			"However intNumSign = '%v'\n",
			ePrefix.String(),
			intNumSign)

		return
	}

	strBuilder := &strings.Builder{}

	err = baseVal.GetParameterTextListing(
		strBuilder,
		false,
		ePrefix.XCpy(
			""))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if strBuilder.Len() < 50 {

		t.Errorf("%v\n"+
			"Setup Test#4 - baseVal.GetParameterTextListing()\n"+
			"Error: Returned string length should be greater than 50.\n"+
			"However string length = '%v'\n",
			ePrefix.String(),
			strBuilder.Len())

		return
	}

	var numOfFracDigits, numOfIntDigits, numOfDigits int

	numOfFracDigits = baseVal.GetNumberOfFractionalDigits()

	if numOfFracDigits != 1 {

		t.Errorf("%v\n"+
			"Setup Test#5 - baseVal.GetNumberOfFractionalDigits()\n"+
			"Error: Returned number of fractional digits should be '1'.\n"+
			"However actual number of fractional digits = '%v'\n",
			ePrefix.String(),
			numOfFracDigits)

		return

	}

	numOfIntDigits = baseVal.GetNumberOfIntegerDigits()

	if numOfIntDigits != 1 {

		t.Errorf("%v\n"+
			"Setup Test#6 - baseVal.GetNumberOfIntegerDigits()\n"+
			"Error: Returned number of integer digits should be '1'.\n"+
			"However actual number of integer digits = '%v'\n",
			ePrefix.String(),
			numOfIntDigits)

		return

	}

	numOfDigits = baseVal.GetNumberOfNumericDigits()

	if numOfDigits != 2 {

		t.Errorf("%v\n"+
			"Setup Test#6 - baseVal.GetNumberOfIntegerDigits()\n"+
			"Error: Returned total number of numeric digits should be '2'.\n"+
			"However actual total number of numeric digits = '%v'\n",
			ePrefix.String(),
			numOfDigits)

		return

	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01-NoRounding"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string
	var pureNumStrComponents PureNumberStrComponents

	actualNumberStr,
		pureNumStrComponents,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1A - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	if pureNumStrComponents.NativeNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1B - Positive Test\n"+
			"pureNumStrComponents\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"pureNumStrComponents.NativeNumStr = '%v'\n"+
			"              expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			pureNumStrComponents.NativeNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02-NoRounding"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-7.5"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfUpWithNegNums_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfUpWithNegNums_000200",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "8"

	var err error
	var nStr01, nStr02 NumberStrKernel

	nStr01,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr01<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&nStr01,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfUpWithNegNums(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfUpWithNegNums"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfUpWithNegNums(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfUpWithNegNums"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-7"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfDownWithNegNums_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfDownWithNegNums_000300()",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "7"

	var err error
	var baseVal, nStr01, nStr02 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfDownWithNegNums(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfDownWithNegNums"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfDownWithNegNums(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfDownWithNegNums"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-8"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfAwayFromZero_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfAwayFromZero_000400",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "8"

	var err error
	var baseVal, nStr01, nStr02 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfAwayFromZero(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfAwayFromZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfAwayFromZero(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfAwayFromZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-8"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfTowardsZero_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfTowardsZero_000500",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "7"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfTowardsZero(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfTowardsZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfTowardsZero(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfTowardsZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	expectedNumberStr = "-7"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-8"
	inputNumberStr = "-7.6"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.HalfTowardsZero(),
		0,
		ePrefix.XCpy(
			"nStr03-HalfTowardsZero"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfToEven_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfToEven_000600",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "8"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03, nStr04 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfToEven(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfToEven"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfToEven(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfToEven"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-8"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "6"
	inputNumberStr = "6.5"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.HalfToEven(),
		0,
		ePrefix.XCpy(
			"nStr03-HalfToEven"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Postive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-6"
	inputNumberStr = "-6.5"

	nStr04,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr04<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr04.Round(
		NumRoundType.HalfToEven(),
		0,
		ePrefix.XCpy(
			"nStr04-HalfToEven"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr04.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#4 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundHalfToOdd_000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundHalfToOdd_000700",
		"")

	inputNumberStr := "7.5"
	expectedNumberStr := "7"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03, nStr04 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.HalfToOdd(),
		0,
		ePrefix.XCpy(
			"nStr01-HalfToOdd"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.HalfToOdd(),
		0,
		ePrefix.XCpy(
			"nStr02-HalfToOdd"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-7"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "7"
	inputNumberStr = "6.5"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.HalfToOdd(),
		0,
		ePrefix.XCpy(
			"nStr03-HalfToOdd"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Postive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-6"
	inputNumberStr = "-6.4"

	nStr04,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr04<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr04.Round(
		NumRoundType.HalfToOdd(),
		0,
		ePrefix.XCpy(
			"nStr04-HalfToOdd"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr04.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#4 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundRandomly_000800(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundRandomly_000800()",
		"")

	inputNumberIntStr := "7"
	inputNumberFracStr := "5"
	var numberSign = NumSignVal.Positive()
	var loopCount = 50

	upperVal := "8"
	lowerVal := "7"
	var actualNumberStr string

	var upperValCount = 0
	var lowerValCount = 0

	var err error

	for i := 0; i < loopCount; i++ {

		var nStr01 NumberStrKernel

		nStr01,
			err = new(NumberStrKernel).NewFromStringDigits(
			inputNumberIntStr,
			inputNumberFracStr,
			numberSign,
			ePrefix.XCpy(
				"nStr01<-"))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		err = nStr01.Round(
			NumRoundType.Randomly(),
			0,
			ePrefix.XCpy(
				"nStr01-Randomly"))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		actualNumberStr,
			_,
			err = nStr01.FmtNumStrPure(
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr01"))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		if actualNumberStr == upperVal {
			upperValCount++
		}

		if actualNumberStr == lowerVal {
			lowerValCount++
		}

	}

	if upperValCount == loopCount {

		t.Errorf("%v\n"+
			"Test#1 - Random Rounding Test\n"+
			"Error: upperValCount == loopCount\n"+
			"Upper Value Rounding Result = '%v'"+
			"Number of roundings = '%v'\n"+
			"Number of results equal to Upper Value = '%v'\n",
			ePrefix.String(),
			upperVal,
			loopCount,
			upperValCount)

		return

	}

	if lowerValCount == loopCount {

		t.Errorf("%v\n"+
			"Test#1 - Random Rounding Test\n"+
			"Error: lowerValCount == loopCount\n"+
			"Lower Value Rounding Result = '%v'"+
			"Number of roundings = '%v'\n"+
			"Number of results equal to Lower Value = '%v'\n",
			ePrefix.String(),
			lowerVal,
			loopCount,
			lowerValCount)

		return

	}

	return
}

func TestNumberStrKernel_RoundFloor_000900(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundFloor_000900",
		"")

	inputNumberStr := "2.9"
	expectedNumberStr := "2"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03, nStr04 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.Floor(),
		0,
		ePrefix.XCpy(
			"nStr01-Floor"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.Floor(),
		0,
		ePrefix.XCpy(
			"nStr02-Floor"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-3"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "2"
	inputNumberStr = "2.4"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.Floor(),
		0,
		ePrefix.XCpy(
			"nStr03-Floor"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Postive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-2"
	inputNumberStr = "-2"

	nStr04,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr04<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr04.Round(
		NumRoundType.Floor(),
		0,
		ePrefix.XCpy(
			"nStr04-Floor"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr04.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#4 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_SetDefaultNumStrFmtSpecElements_001000(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_SetDefaultNumStrFmtSpecElements_001000()",
		"")

	origNumStrValue := "-1234456.78"

	var err error
	var numStrKernel NumberStrKernel

	numStrKernel,
		_,
		err = new(NumberStrKernel).NewParseNativeNumberStr(
		origNumStrValue,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			""))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr := origNumStrValue

	var actualNumberStr string

	actualNumberStr,
		_,
		err = numStrKernel.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"numStrKernel Test#1"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := fmt.Sprintf("Test #1 Original Number String Setup Verification\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}

	testName = fmt.Sprintf("Test #2 SetDefaultNumStrFmtSpecElements()\n"+
		"Default Format has NO Currency Symbols.\n"+
		"Original Native Number String Value = %v\n",
		origNumStrValue)

	expectedNumberStr = "  1.234.456,78-"

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        15,
		fieldJustification: TxtJustify.Right(),
		lock:               nil,
	}

	var decSepSpec DecimalSeparatorSpec

	decSepSpec,
		err = new(DecimalSeparatorSpec).NewStr(
		",",
		ePrefix.XCpy(
			"decSepSpec<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var intSepSpec IntegerSeparatorSpec

	intSepSpec,
		err = new(IntegerSeparatorSpec).NewThousands(
		".",
		ePrefix.XCpy(
			"intSepSpec"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		currencySymbol NumStrNumberSymbolSpec

	negativeNumberSign,
		err = new(NumStrNumberSymbolSpec).
		NewNumberSignTrailingSymbol(
			"-",
			NumFieldSymPos.InsideNumField(),
			ePrefix.XCpy(
				"negativeNumberSign<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	positiveNumberSign = new(NumStrNumberSymbolSpec).NewNOP()

	zeroNumberSign = new(NumStrNumberSymbolSpec).NewNOP()

	currencySymbol = new(NumStrNumberSymbolSpec).NewNOP()

	err = numStrKernel.SetDefaultNumStrFmtSpecElements(
		decSepSpec,
		intSepSpec,
		negativeNumberSign,
		positiveNumberSign,
		zeroNumberSign,
		currencySymbol,
		numberFieldSpec,
		ePrefix.XCpy(
			"numStrKernel02<-Test #3"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr = numStrKernel.String()

	if expectedNumberStr != actualNumberStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_RoundCeiling_001000(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_RoundCeiling_001000",
		"")

	inputNumberStr := "2.9"
	expectedNumberStr := "3"

	var err error
	var baseVal, nStr01, nStr02,
		nStr03, nStr04 NumberStrKernel

	baseVal,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseVal<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseVal.IsValidInstanceError(
		ePrefix.XCpy(
			"baseVal"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr01<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.CopyIn(
		&baseVal,
		ePrefix.XCpy(
			"nStr02<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr02.SetNumberSign(
		NumSignVal.Negative(),
		ePrefix.XCpy(
			"nStr02-NumSignVal.Negative()"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr01.Round(
		NumRoundType.Ceiling(),
		0,
		ePrefix.XCpy(
			"nStr01-Ceiling"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumberStr string

	actualNumberStr,
		_,
		err = nStr01.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#1 - Positive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	err = nStr02.Round(
		NumRoundType.Ceiling(),
		0,
		ePrefix.XCpy(
			"nStr02-Ceiling"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr02.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "-2"

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#2 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "3"
	inputNumberStr = "2.4"

	nStr03,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr03<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr03.Round(
		NumRoundType.Ceiling(),
		0,
		ePrefix.XCpy(
			"nStr03-Ceiling"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr03.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#3 - Postive Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	expectedNumberStr = "-2"
	inputNumberStr = "-2"

	nStr04,
		_,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			inputNumberStr,
			".",
			true,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"nStr04<-inputNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = nStr04.Round(
		NumRoundType.Ceiling(),
		0,
		ePrefix.XCpy(
			"nStr04-Ceiling"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumberStr,
		_,
		err = nStr04.FmtNumStrPure(
		".",
		true,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStr01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualNumberStr {

		t.Errorf("%v\n"+
			"Test#4 - Negative Test\n"+
			"Error: actualNumberStr NOT EQUAL TO expectedNumberStr\n"+
			"    actualNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			actualNumberStr,
			expectedNumberStr)

		return
	}

	return
}

func TestNumberStrKernel_SetFromDirtyNumberStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_SetFromDirtyNumberStr_000100()",
		"")

	dirtyNumberStr := "1.123.456,78 €"

	expectedStr := "1123456.78"

	var nStrKernel01 NumberStrKernel

	var numStats NumberStrStatsDto
	var err error

	numStats,
		err = nStrKernel01.SetFromDirtyNumberStr(
		dirtyNumberStr,
		",",
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"nStrKernel01<-dirtyNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !nStrKernel01.IsValidInstance() {

		err = nStrKernel01.IsValidInstanceError(
			ePrefix.XCpy(
				"nStrKernel01"))

		if err != nil {

			t.Errorf("Error return from nStrKernel01.IsValidInstanceError()\n"+
				"Error= \n%v\n",
				err.Error())

			return
		}
	}

	testName := "Test #1 Number Stats Fractional Digits Test"

	if numStats.NumOfFractionalDigits != 2 ||
		numStats.NumOfSignificantFractionalDigits != 2 {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: Number Stats Fractional Digits Not Equal to '2'\n"+
			"           numStats.NumOfFractionalDigits = %v\n"+
			"numStats.NumOfSignificantFractionalDigits = %v\n",
			ePrefix.String(),
			testName,
			numStats.NumOfFractionalDigits,
			numStats.NumOfSignificantFractionalDigits)

		return
	}

	var actualFmtStr string

	actualFmtStr,
		numStats,
		err = nStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"actualFmtStr<-StrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test #2 Native Number String Comparison"

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	return
}

func TestNumberStrKernel_SetFromNativeNumberStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_SetFromNativeNumberStr_000100()",
		"")

	originalNativeNumberStr := "1123456.775"

	expectedStr := "1123456.78"

	var nStrKernel01 NumberStrKernel

	var numStats NumberStrStatsDto
	var err error

	numStats,
		err = nStrKernel01.SetFromNativeNumberStr(
		originalNativeNumberStr,
		NumRoundType.HalfAwayFromZero(),
		2,
		ePrefix.XCpy(
			"nStrKernel01<-originalNativeNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !nStrKernel01.IsValidInstance() {

		err = nStrKernel01.IsValidInstanceError(
			ePrefix.XCpy(
				"nStrKernel01"))

		if err != nil {

			t.Errorf("Error return from nStrKernel01.IsValidInstanceError()\n"+
				"Error= \n%v\n",
				err.Error())

			return
		}
	}

	testName := "Test #1 Number Stats Fractional Digits Test"

	if numStats.NumOfFractionalDigits != 2 ||
		numStats.NumOfSignificantFractionalDigits != 2 {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: Number Stats Fractional Digits Not Equal to '2'\n"+
			"           numStats.NumOfFractionalDigits = %v\n"+
			"numStats.NumOfSignificantFractionalDigits = %v\n",
			ePrefix.String(),
			testName,
			numStats.NumOfFractionalDigits,
			numStats.NumOfSignificantFractionalDigits)

		return
	}

	var actualFmtStr string

	actualFmtStr,
		numStats,
		err = nStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"actualFmtStr<-StrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test #2 Native Number String Comparison"

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	originalNativeNumberStr = "-1123456.775"

	expectedStr = "-1123456.78"

	var nStrKernel02 NumberStrKernel

	numStats,
		err = nStrKernel02.SetFromNativeNumberStr(
		originalNativeNumberStr,
		NumRoundType.HalfAwayFromZero(),
		2,
		ePrefix.XCpy(
			"nStrKernel02<-originalNativeNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !nStrKernel02.IsValidInstance() {

		err = nStrKernel02.IsValidInstanceError(
			ePrefix.XCpy(
				"nStrKernel02"))

		if err != nil {

			t.Errorf("Error return from nStrKernel02.IsValidInstanceError()\n"+
				"Error= \n%v\n",
				err.Error())

			return
		}
	}

	testName = "Test #3 Number Stats Fractional Digits Test"

	if numStats.NumOfFractionalDigits != 2 ||
		numStats.NumOfSignificantFractionalDigits != 2 {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: Number Stats Fractional Digits Not Equal to '2'\n"+
			"           numStats.NumOfFractionalDigits = %v\n"+
			"numStats.NumOfSignificantFractionalDigits = %v\n",
			ePrefix.String(),
			testName,
			numStats.NumOfFractionalDigits,
			numStats.NumOfSignificantFractionalDigits)

		return
	}

	actualFmtStr,
		numStats,
		err = nStrKernel02.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"actualFmtStr<-StrKernel02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName = "Test #4 Negative Native Number String Comparison"

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	return
}

func TestNumberStrKernel_SetFromNumericValue_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_SetFromNumericValue_000100()",
		"")

	var actualFmtStr, expectedStr, testName string
	var numStats NumberStrStatsDto

	inputStr := "Hello World (1123456.775) Goodbye World"

	var nStrKernel01 NumberStrKernel

	startSearchIndex := 11
	characterSearchLength := 26

	negativeNumSignSearchSpecs := NegNumSearchSpecCollection{}

	var err error

	err =
		negativeNumSignSearchSpecs.AddLeadingAndTrailingNegNumSearchStr(
			"(",
			")",
			ePrefix.XCpy(
				"()"))

	var decimalSeparator DecimalSeparatorSpec

	decimalSeparator,
		err = new(DecimalSeparatorSpec).NewRunes(
		[]rune{'.'},
		ePrefix)

	numParsingTerminators := []string{
		"oo",
		"bye",
	}

	var requestRemainderString bool

	requestRemainderString = true

	var numberStrSearchResults CharSearchNumStrParseResultsDto

	/*
	   func (numStrKernel *NumberStrKernel) NewParseCustomNumberStr(
	   	rawNumStr string,
	   	startSearchIndex int,
	   	characterSearchLength int,
	   	negativeNumSignSearchSpecs NegNumSearchSpecCollection,
	   	decimalSeparator DecimalSeparatorSpec,
	   	numParsingTerminators []string,
	   	requestRemainderString bool,
	   	errorPrefix interface{}) (
	   	numberStrSearchResults CharSearchNumStrParseResultsDto,
	   	numberStrKernel NumberStrKernel,
	   	err error) {


	*/

	numberStrSearchResults,
		nStrKernel01,
		err = new(NumberStrKernel).NewParseCustomNumberStr(
		inputStr,
		startSearchIndex,
		characterSearchLength,
		negativeNumSignSearchSpecs,
		decimalSeparator,
		numParsingTerminators,
		requestRemainderString,
		ePrefix.XCpy(
			"nStrKernel01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !nStrKernel01.IsValidInstance() {

		err = nStrKernel01.IsValidInstanceError(
			ePrefix.XCpy(
				"NewParseCustomNumberStr - nStrKernel01"))

		if err != nil {

			t.Errorf("Error return from nStrKernel01.IsValidInstanceError()\n"+
				"Error= \n%v\n",
				err.Error())

			return
		}
	}

	testName = "Test #1-A numberStrSearchResults.FoundNumericDigits"

	if !numberStrSearchResults.FoundNumericDigits {

		t.Errorf("%v\n"+
			"%v\n"+
			"Expected 'FoundNumericDigits' = 'true'\n"+
			"Instead 'FoundNumericDigits' = 'false'\n",
			ePrefix.String(),
			testName)

		return
	}

	testName = "Test #1-B numberStrSearchResults.RemainderString"

	if numberStrSearchResults.RemainderString.GetRuneArrayLength() == 0 {

		t.Errorf("%v\n"+
			"%v\n"+
			"Expected 'RemainderString' Length > 0 \n"+
			"Instead 'RemainderString' Length == 0\n",
			ePrefix.String(),
			testName)

		return
	}

	testName = "Test #1-C numberStrSearchResults.NumSignValue"

	if numberStrSearchResults.NumSignValue != NumSignVal.Negative() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Expected 'NumSignValue' == NumSignVal.Negative() \n"+
			"Instead 'NumSignValue' String = %v\n"+
			" 'NumSignValue' Integer Value = %v\n",
			ePrefix.String(),
			testName,
			numberStrSearchResults.NumSignValue.String(),
			numberStrSearchResults.NumSignValue.XValueInt())

		return
	}

	testName = "Test #1-D NumberStrStatsDtoNumberSign Should Be NumSignVal.Negative()"

	actualFmtStr,
		numStats,
		err = nStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"1-D actualFmtStr<-"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if numStats.NumberSign != NumSignVal.Negative() {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: numStats.NumberSign != NumSignVal.Negative()\n"+
			"Expected numStats.NumberSign = %v\n"+
			"  Actual numStats.NumberSign = %v\n",
			ePrefix.String(),
			testName,
			NumSignVal.Negative().String(),
			numStats.NumberSign.String())

		return

	}

	testName = "Test #1-E nStrKernel01.FmtNumStrNative()"

	expectedStr = "-1123456.775"

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	testName = "Test #2-A nStrKernel01.FmtSignedNumStrDefaultsGermany()"

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        -1,
		fieldJustification: TxtJustify.Right(),
	}

	roundingSpec := NumStrRoundingSpec{
		roundingType:            NumRoundType.NoRounding(),
		roundToFractionalDigits: 0,
		lock:                    nil,
	}

	actualFmtStr,
		err = nStrKernel01.FmtSignedNumStrDefaultsGermany(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy(
			"2-A nStrKernel01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr = "1.123.456,775-"

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	testName = "Test #3-A nStrKernel01.FmtSignedNumStrDefaultsGermany()"

	expectedStr = "576123.45678"

	//inputStr := "1123456.775) Goodbye World"
	var inputBigFloat big.Float

	inputBigFloat,
		err = new(MathFloatHelper).NativeNumStrToBigFloat(
		expectedStr,
		ePrefix.XCpy(
			"3-A"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = nStrKernel01.SetFromNumericValue(
		inputBigFloat,
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"3-A nStrKernel01<-inputBigFloat"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !nStrKernel01.IsValidInstance() {

		err = nStrKernel01.IsValidInstanceError(
			ePrefix.XCpy(
				"nStrKernel01"))

		if err != nil {

			t.Errorf("Error return from nStrKernel01.IsValidInstanceError()\n"+
				"Error= \n%v\n",
				err.Error())

			return
		}
	}

	actualFmtStr,
		_,
		err = nStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"3-A actualFmtStr<-nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	nStrKernel01.Empty()

	return
}

func TestNumberStrKernel_SetFromPureNumberStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_String_000100()",
		"")

	var actualFmtStr, expectedStr, testName string
	var numStats1, numStats2 NumberStrStatsDto
	var err error

	inputStr := "$1 123 456,775 -"

	var nStrKernel01 NumberStrKernel

	nStrKernel01,
		numStats1,
		err = new(NumberStrKernel).
		NewParseDirtyNumberStr(
			inputStr,
			",",
			NumRoundType.HalfAwayFromZero(),
			2,
			ePrefix.XCpy(
				"1-A nStrKernel01<-"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !nStrKernel01.IsValidInstance() {

		err = nStrKernel01.IsValidInstanceError(
			ePrefix.XCpy(
				"1-A nStrKernel01"))

		if err != nil {

			t.Errorf("Error 1-A return from nStrKernel01.IsValidInstanceError()\n"+
				"Error= \n%v\n",
				err.Error())

			return
		}
	}

	testName = "Test 1-A NumberStrKernel.NewParseDirtyNumberStr() NumberStrStatsDto"

	numStats2,
		err = nStrKernel01.GetNumericValueStats(
		ePrefix.XCpy(
			""))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if numStats1 != numStats2 {

		t.Errorf("%v\n"+
			"%v\n"+
			"ERROR: numStats1 != numStats2 "+
			"numStats1 = \n%v\n\n"+
			"numStats2 = \n%v\n\n",
			ePrefix.String(),
			testName,
			numStats1.String(),
			numStats2.String())

		return

	}

	testName = "Test #1-B nStrKernel01.FmtNumStrNative()"

	expectedStr = "-1123456.78"

	actualFmtStr,
		numStats1,
		err = nStrKernel01.FmtNumStrNative(
		NumRoundType.NoRounding(),
		0,
		ePrefix.XCpy(
			"1-D actualFmtStr<-"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	testName = "Test #2-A nStrKernel01.SetFromPureNumberStr()\n" +
		"French Currency Format"

	nStrKernel01.Empty()

	inputStr = "1234,9876-"
	expectedStr = "-1234.988"

	_,
		err = nStrKernel01.SetFromPureNumberStr(
		inputStr,
		",",
		false,
		NumRoundType.HalfAwayFromZero(),
		3,
		ePrefix.XCpy(
			"2-A nStrKernel01<-inputStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !nStrKernel01.IsValidInstance() {

		err = nStrKernel01.IsValidInstanceError(
			ePrefix.XCpy(
				"2-A - nStrKernel01"))

		if err != nil {

			t.Errorf("Error return from nStrKernel01.IsValidInstanceError()\n"+
				"Test 2-A"+
				"Error= \n%v\n",
				err.Error())

			return
		}
	}

	numberFieldSpec := NumStrNumberFieldSpec{
		fieldLength:        -1,
		fieldJustification: TxtJustify.Right(),
	}

	roundingSpec := NumStrRoundingSpec{
		roundingType:            NumRoundType.NoRounding(),
		roundToFractionalDigits: 0,
	}

	actualFmtStr,
		err = nStrKernel01.FmtCurrencyDefaultsFrance(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy("2-A nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	// expectedStr = "-1234.988"
	expectedStr = "-1 234,988 €"

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	testName = "Test #3-A nStrKernel01.FmtCurrencyDefaultsGermany()\n" +
		"German Currency Format"

	// expectedStr = "-1234.988"
	expectedStr = "1.234,988- €"

	actualFmtStr,
		err = nStrKernel01.FmtCurrencyDefaultsGermany(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy("3-A nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	testName = "Test #4-A nStrKernel01.FmtCurrencyDefaultsUKMinusOutside()\n" +
		"UK Currency Format"

	// expectedStr = "-1234.988"
	expectedStr = "-£ 1,234.988"

	actualFmtStr,
		err = nStrKernel01.FmtCurrencyDefaultsUKMinusOutside(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy("4-A nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	testName = "Test #5-A nStrKernel01.FmtCurrencyDefaultsUSParen()\n" +
		"US Currency Format"

	// expectedStr = "-1234.988"
	expectedStr = "$ (1,234.988)"

	actualFmtStr,
		err = nStrKernel01.FmtCurrencyDefaultsUSParen(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy("5-A nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	testName = "Test #6-A nStrKernel01.FmtCurrencyDefaultsUSMinus()\n" +
		"US Currency Format"

	// expectedStr = "-1234.988"
	expectedStr = "$ -1,234.99"

	err = roundingSpec.SetRoundingSpec(
		NumRoundType.HalfAwayFromZero(),
		2,
		ePrefix.XCpy(
			""))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualFmtStr,
		err = nStrKernel01.FmtCurrencyDefaultsUSMinus(
		roundingSpec,
		numberFieldSpec,
		ePrefix.XCpy("6-A nStrKernel01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedStr != actualFmtStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: expectedStr != actualFmtStr\n"+
			" expectedStr = %v\n"+
			"actualFmtStr = %v\n",
			ePrefix.String(),
			testName,
			expectedStr,
			actualFmtStr)

		return
	}

	return
}

func TestNumberStrKernel_String_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumberStrKernel_String_000100()",
		"")

	origNumberStr := "1234.5678"

	expectedNumberStr := "1,234.5678"

	var err error
	var baseValueNStr NumberStrKernel

	baseValueNStr,
		_,
		err = new(NumberStrKernel).
		NewParseNativeNumberStr(
			origNumberStr,
			NumRoundType.NoRounding(),
			0,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualFmtNumberStr string

	actualFmtNumberStr = baseValueNStr.String()

	testName := "Test #1 - Default Format Spec"

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualFmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualFmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	testName = "Test #2 - Rounded Before Default Format Spec"

	err = baseValueNStr.Round(
		NumRoundType.HalfAwayFromZero(),
		2,
		ePrefix.XCpy("baseValueNStr<-2digits"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumberStr = "1,234.57"

	actualFmtNumberStr = baseValueNStr.String()

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualFmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualFmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	testName = "Test #3 - Rounded Before Default Format Spec"

	actualFmtNumberStr,
		err = baseValueNStr.FmtNumStrDefault(
		ePrefix.XCpy(
			"actualFmtNumberStr<=baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualFmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualFmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	testName = "Test #3 - Rounded During Default Format Spec"

	expectedNumberStr = "1,235"

	roundingSpec := NumStrRoundingSpec{
		roundingType:            NumRoundType.HalfAwayFromZero(),
		roundToFractionalDigits: 0,
	}

	actualFmtNumberStr,
		err = baseValueNStr.FmtNumStrDefaultRound(
		roundingSpec,
		ePrefix.XCpy(
			"actualFmtNumberStr<=baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumberStr != actualFmtNumberStr {

		t.Errorf("%v\n"+
			"%v\n"+
			"Error: actualFmtNumberStr NOT EQUAL TO expectedNumberStr\n"+
			" actualFmtNumberStr = '%v'\n"+
			"expectedNumberStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualFmtNumberStr,
			expectedNumberStr)

		return
	}

	return
}

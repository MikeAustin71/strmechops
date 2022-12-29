package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math/big"
	"strconv"
	"testing"
)

func TestMathHelper_NumericValueToNativeNumStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000100()",
		"")

	expectedStr := "12345.123456"

	var bigFloatNum big.Float
	var ok bool
	var err error
	_,
		ok = bigFloatNum.
		SetMode(big.AwayFromZero).
		SetString(expectedStr)

	if !ok {

		t.Errorf("\n%v\n"+
			"Error: bigFloatNum.SetString(expectedStr) Failed!\n"+
			"expectedStr= %v\n",
			ePrefix.String(),
			expectedStr)

		return
	}

	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		bigFloatNum,
		ePrefix.XCpy(
			"actualNumStr<-bigFloatNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"Test#1 \n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000200()",
		"")

	expectedStr := "12345.12345"

	var float64Num float64
	var err error

	float64Num,
		err = strconv.ParseFloat(expectedStr, 64)

	if err != nil {

		t.Errorf("\n%v\n"+
			"Error: strconv.ParseFloat(expectedStr,64) Failed!\n"+
			"expectedStr= %v\n"+
			"Error = \n%v\n",
			ePrefix.String(),
			expectedStr,
			err.Error())

		return
	}

	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		float64Num,
		ePrefix.XCpy(
			"actualNumStr<-float64Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"Test#1 \n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			actualNumStr,
			expectedStr)

		return

	}

	return

}

func TestMathHelper_NumericValueToNativeNumStr_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000200()",
		"")

	origIntStr := "12345"
	origFracStr := "12345"

	expectedStr := origIntStr +
		"." +
		origFracStr

	var err error
	var baseValueNStr NumberStrKernel

	baseValueNStr,
		err = new(NumberStrKernel).
		NewParsePureNumberStr(
			expectedStr,
			".",
			true,
			ePrefix.XCpy(
				"baseValueNStr<-origNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}
	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		baseValueNStr,
		ePrefix.XCpy(
			"actualNumStr<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"Test#1 \n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000400()",
		"")

	testName := "Test #1 - int8(123)"

	expectedStr := "123"

	int8Num := int8(123)

	var err error
	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		int8Num,
		ePrefix.XCpy(
			"actualNumStr<-int8Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #2 - int16(137)"
	expectedStr = "137"

	int16Num := int16(137)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		int16Num,
		ePrefix.XCpy(
			"actualNumStr<-int16Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #3 - int(152)"

	expectedStr = "152"

	intNum := 152

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		intNum,
		ePrefix.XCpy(
			"actualNumStr<-intNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #4 - int32(1921)"

	expectedStr = "1921"

	int32Num := int32(1921)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		int32Num,
		ePrefix.XCpy(
			"actualNumStr<-int32Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #5 - int64(220005197)"

	expectedStr = "220005197"

	int64Num := int64(220005197)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		int64Num,
		ePrefix.XCpy(
			"actualNumStr<-int64Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #6 - bigIntNum.SetString(5961234567890, 10)"

	expectedStr = "5961234567890"

	var ptrBigIntNum *big.Int
	var bigIntNum big.Int
	var ok bool

	ptrBigIntNum,
		ok = big.NewInt(0).SetString(expectedStr, 10)

	if !ok {

		t.Errorf("\n%v\n"+
			"Error: bigIntNum.SetString(expectedStr, 10) Failed!\n"+
			"expectedStr= %v\n",
			ePrefix.String(),
			expectedStr)

		return
	}

	bigIntNum.Set(ptrBigIntNum)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		bigIntNum,
		ePrefix.XCpy(
			"actualNumStr<-bigIntNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #7 - ptrBigIntNum.SetString(6239874128, 10)"

	expectedStr = "6239874128"

	_,
		ok = ptrBigIntNum.SetString(expectedStr, 10)

	if !ok {

		t.Errorf("\n%v\n"+
			"Error: bigIntNum.SetString(expectedStr, 10) Failed!\n"+
			"expectedStr= %v\n",
			ePrefix.String(),
			expectedStr)

		return
	}

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrBigIntNum,
		ePrefix.XCpy(
			"actualNumStr<-ptrBigIntNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000500()",
		"")

	testName := "Test #1 - *int8 ptrInt8Num(123)"

	expectedStr := "123"

	int8Num := int8(123)

	var ptrInt8Num *int8

	ptrInt8Num = &int8Num

	mathHelperUtil := MathHelper{}

	var err error
	var actualNumStr string

	actualNumStr,
		err = mathHelperUtil.NumericValueToNativeNumStr(
		ptrInt8Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrInt8Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #2 - *int16 ptrInt16Num(137)"
	expectedStr = "137"

	int16Num := int16(137)

	var ptrInt16Num *int16

	ptrInt16Num = &int16Num

	actualNumStr,
		err = mathHelperUtil.NumericValueToNativeNumStr(
		ptrInt16Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrInt16Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #3 - *int ptrIntNum(152)"

	expectedStr = "152"

	intNum := 152

	var ptrIntNum *int
	ptrIntNum = &intNum

	actualNumStr,
		err = mathHelperUtil.NumericValueToNativeNumStr(
		ptrIntNum,
		ePrefix.XCpy(
			"actualNumStr<-ptrIntNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #4 - *int32 ptrInt32Num(1921)"

	expectedStr = "1921"

	int32Num := int32(1921)

	var ptrInt32Num *int32

	ptrInt32Num = &int32Num

	actualNumStr,
		err = mathHelperUtil.NumericValueToNativeNumStr(
		ptrInt32Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrInt32Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #5 - *int64 ptrInt64Num(220005197)"

	expectedStr = "220005197"

	int64Num := int64(220005197)

	var ptrInt64Num *int64

	ptrInt64Num = &int64Num

	actualNumStr,
		err = mathHelperUtil.NumericValueToNativeNumStr(
		ptrInt64Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrInt64Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	ptrInt64Num = nil

	actualNumStr,
		err = mathHelperUtil.NumericValueToNativeNumStr(
		ptrInt64Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrInt64Num"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Test#6 - Did NOT Receive Expected Error\n"+
			"Expected to receive an error from:\n"+
			"mathHelperUtil.NumericValueToNativeNumStr(ptrInt64Num)\n"+
			"because ptrInt64Num is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000600()",
		"")

	testName := "Test #1 - uint8(123)"

	expectedStr := "123"

	uint8Num := uint8(123)

	var err error
	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		uint8Num,
		ePrefix.XCpy(
			"actualNumStr<-uint8Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #2 - uint16(137)"
	expectedStr = "137"

	uint16Num := uint16(137)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		uint16Num,
		ePrefix.XCpy(
			"actualNumStr<-uint16Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #3 - uint(152)"

	expectedStr = "152"

	uintNum := uint(152)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		uintNum,
		ePrefix.XCpy(
			"actualNumStr<-uintNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #4 - uint32(1921)"

	expectedStr = "1921"

	uint32Num := uint32(1921)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		uint32Num,
		ePrefix.XCpy(
			"actualNumStr<-uint32Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #5 - uint64(220005197)"

	expectedStr = "220005197"

	uint64Num := uint64(220005197)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		uint64Num,
		ePrefix.XCpy(
			"actualNumStr<-uint64Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000700()",
		"")

	testName := "Test #1 - *uint8 ptrUint8Num(123)"

	expectedStr := "123"

	uint8Num := uint8(123)

	var ptrUint8Num *uint8

	ptrUint8Num = &uint8Num

	var err error
	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrUint8Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrUint8Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #2 - *uint16 ptrUint16Num(137)"
	expectedStr = "137"

	uint16Num := uint16(137)

	var ptrUint16Num *uint16

	ptrUint16Num = &uint16Num

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrUint16Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrUint16Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #3 - *uint  ptrUintNum(152)"

	expectedStr = "152"

	uintNum := uint(152)

	var ptrUintNum *uint

	ptrUintNum = &uintNum

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrUintNum,
		ePrefix.XCpy(
			"actualNumStr<-ptrUintNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #4 - *uint32 ptrUint32Num(1921)"

	expectedStr = "1921"

	uint32Num := uint32(1921)

	var ptrUint32Num *uint32

	ptrUint32Num = &uint32Num

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrUint32Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrUint32Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #5 - *uint64  ptrUint64Num(220005197)"

	expectedStr = "220005197"

	uint64Num := uint64(220005197)

	var ptrUint64Num *uint64

	ptrUint64Num = &uint64Num

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrUint64Num,
		ePrefix.XCpy(
			"actualNumStr<-ptrUint64Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_000800(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000800()",
		"")

	testName := "Test #1 - int8(-123)"

	expectedStr := "-123"

	int8Num := int8(-123)

	var err error
	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		int8Num,
		ePrefix.XCpy(
			"actualNumStr<-int8Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #2 - int16(-137)"
	expectedStr = "-137"

	int16Num := int16(-137)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		int16Num,
		ePrefix.XCpy(
			"actualNumStr<-int16Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #3 - int(-152)"

	expectedStr = "-152"

	intNum := -152

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		intNum,
		ePrefix.XCpy(
			"actualNumStr<-intNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #4 - int32(-1921)"

	expectedStr = "-1921"

	int32Num := int32(-1921)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		int32Num,
		ePrefix.XCpy(
			"actualNumStr<-int32Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #5 - int64(-220005197)"

	expectedStr = "-220005197"

	int64Num := int64(-220005197)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		int64Num,
		ePrefix.XCpy(
			"actualNumStr<-int64Num"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #6 - bigIntNum.SetString(-5961234567890, 10)"

	expectedStr = "-5961234567890"

	var ptrBigIntNum *big.Int
	var bigIntNum big.Int
	var ok bool

	ptrBigIntNum,
		ok = big.NewInt(0).SetString(expectedStr, 10)

	if !ok {

		t.Errorf("\n%v\n"+
			"Error: bigIntNum.SetString(expectedStr, 10) Failed!\n"+
			"expectedStr= %v\n",
			ePrefix.String(),
			expectedStr)

		return
	}

	bigIntNum.Set(ptrBigIntNum)

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		bigIntNum,
		ePrefix.XCpy(
			"actualNumStr<-bigIntNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	testName = "Test #7 - ptrBigIntNum.SetString(-6239874128, 10)"

	expectedStr = "-6239874128"

	_,
		ok = ptrBigIntNum.SetString(expectedStr, 10)

	if !ok {

		t.Errorf("\n%v\n"+
			"Error: bigIntNum.SetString(expectedStr, 10) Failed!\n"+
			"expectedStr= %v\n",
			ePrefix.String(),
			expectedStr)

		return
	}

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrBigIntNum,
		ePrefix.XCpy(
			"actualNumStr<-ptrBigIntNum"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_000900(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_000900()",
		"")

	expectedStr := "-1234.5678"

	var err error
	var baseValueNStr NumberStrKernel

	numberStrKernel := NumberStrKernel{}

	baseValueNStr,
		err = numberStrKernel.
		NewParsePureNumberStr(
			expectedStr,
			".",
			true,
			ePrefix.XCpy(
				"baseValueNStr<-expectedNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = baseValueNStr.IsValidInstanceError(
		ePrefix.XCpy(
			"baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	testName := fmt.Sprintf("Test #1 - NumberStrKernel baseValueNStr(%v)",
		expectedStr)

	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		baseValueNStr,
		ePrefix.XCpy(
			"actualNumStr<-baseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-543210"

	testName = fmt.Sprintf("Test #2 - *NumberStrKernel ptrBaseValueNStr(%v)",
		expectedStr)

	baseValueNStr,
		err = numberStrKernel.
		NewParsePureNumberStr(
			expectedStr,
			".",
			true,
			ePrefix.XCpy(
				"baseValueNStr<-expectedNumberStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var ptrBaseValueNStr *NumberStrKernel

	ptrBaseValueNStr = &baseValueNStr

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrBaseValueNStr,
		ePrefix.XCpy(
			"actualNumStr<-ptrBaseValueNStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

}

func TestMathHelper_NumericValueToNativeNumStr_001000(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_001000()",
		"")

	expectedStr := "1234.5678"

	testName := fmt.Sprintf("Test #1 TextFieldFormatDtoBigFloat - txtFieldFmtDtoBFloat(%v)\n",
		expectedStr)

	txtFieldFmtDtoBFloat := TextFieldFormatDtoBigFloat{
		LeftMarginStr:         "",
		BigFloatNum:           big.Float{},
		LeadingMinusSign:      false,
		NativeRoundingMode:    big.AwayFromZero,
		RoundingType:          NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var err error

	err = txtFieldFmtDtoBFloat.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoBFloat<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		txtFieldFmtDtoBFloat,
		ePrefix.XCpy(
			"actualNumStr<-txtFieldFmtDtoBFloat"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-87654321.12345678"

	testName = fmt.Sprintf("Test #2 *TextFieldFormatDtoBigFloat - ptrTxtFieldFmtDtoBFloat2(%v)\n",
		expectedStr)

	txtFieldFmtDtoBFloat2 := TextFieldFormatDtoBigFloat{
		LeftMarginStr:         "",
		BigFloatNum:           big.Float{},
		LeadingMinusSign:      false,
		NativeRoundingMode:    big.AwayFromZero,
		RoundingType:          NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          TxtJustify.Right(),
		RightMarginStr:        "",
	}

	err = txtFieldFmtDtoBFloat2.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoBFloat2<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		&txtFieldFmtDtoBFloat2,
		ePrefix.XCpy(
			"actualNumStr<-ptrTxtFieldFmtDtoBFloat2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_001100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_001100()",
		"")

	expectedStr := "1234.5678"

	testName := fmt.Sprintf("Test #1 TextFieldFormatDtoFloat64 - txtFieldFmtDtoBFloat(%v)\n",
		expectedStr)

	txtFieldFmtDtoF64 := TextFieldFormatDtoFloat64{
		LeftMarginStr:         "",
		Float64Num:            0,
		LeadingMinusSign:      true,
		RoundingType:          NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var err error

	err = txtFieldFmtDtoF64.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoF64<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		txtFieldFmtDtoF64,
		ePrefix.XCpy(
			"actualNumStr<-txtFieldFmtDtoF64"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-8761.123456"

	testName = fmt.Sprintf("Test #2 *TextFieldFormatDtoFloat64 - ptrTxtFieldFmtDtoF64(%v)\n",
		expectedStr)

	txtFieldFmtDtoF642 := TextFieldFormatDtoFloat64{
		LeftMarginStr:         "",
		Float64Num:            0,
		LeadingMinusSign:      true,
		RoundingType:          NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var ptrTxtFieldFmtDtoF64 *TextFieldFormatDtoFloat64

	err = txtFieldFmtDtoF642.SetFromNativeNumStr(
		expectedStr,
		ePrefix.XCpy(
			"txtFieldFmtDtoF642<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	ptrTxtFieldFmtDtoF64 = &txtFieldFmtDtoF642

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrTxtFieldFmtDtoF64,
		ePrefix.XCpy(
			"actualNumStr<-txtFieldFmtDtoF64"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_001200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_001100()",
		"")

	expectedStr := "12345678.12345678902"

	testName := fmt.Sprintf("Test #1 BigFloatDto - bigFloatDto(%v)\n",
		expectedStr)

	var bigFloatDto BigFloatDto
	var err error

	bigFloatDto,
		err = new(MathFloatHelper).PureNumStrToBigFloatDto(
		expectedStr,
		".",
		true,
		2,
		0,
		big.AwayFromZero,
		ePrefix.XCpy(
			"bigFloatDto<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		bigFloatDto,
		ePrefix.XCpy(
			"actualNumStr<-bigFloatDto"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-378421.123456789012"

	testName = fmt.Sprintf("Test #2 *BigFloatDto - ptrBigFloatDto(%v)\n",
		expectedStr)

	var bigFloatDto2 BigFloatDto

	bigFloatDto2,
		err = new(MathFloatHelper).PureNumStrToBigFloatDto(
		expectedStr,
		".",
		true,
		15,
		0,
		big.AwayFromZero,
		ePrefix.XCpy(
			"bigFloatDto2<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var ptrBigFloatDto *BigFloatDto

	ptrBigFloatDto = &bigFloatDto2

	actualNumStr,
		err = new(MathHelper).NumericValueToNativeNumStr(
		ptrBigFloatDto,
		ePrefix.XCpy(
			"actualNumStr<-ptrBigFloatDto"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	return
}

func TestMathHelper_NumericValueToNativeNumStr_001400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NumericValueToNativeNumStr_001400()",
		"")

	testProfile := []testNativeNumStrProfile{
		{
			IntegerDigits:    1,
			FractionalDigits: 0,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    1,
			FractionalDigits: 0,
			IsNegativeValue:  true,
		},
		{
			IntegerDigits:    8,
			FractionalDigits: 12,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    8,
			FractionalDigits: 12,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    5,
			FractionalDigits: 19,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    5,
			FractionalDigits: 19,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    50,
			FractionalDigits: 0,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    50,
			FractionalDigits: 0,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    14,
			FractionalDigits: 157,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    14,
			FractionalDigits: 157,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    1,
			FractionalDigits: 283,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    1,
			FractionalDigits: 283,
			IsNegativeValue:  true,
		},
		{
			IntegerDigits:    276,
			FractionalDigits: 14,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    276,
			FractionalDigits: 14,
			IsNegativeValue:  true,
		},
		{
			IntegerDigits:    383,
			FractionalDigits: 512,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    383,
			FractionalDigits: 512,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    9,
			FractionalDigits: 902,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    9,
			FractionalDigits: 902,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    23,
			FractionalDigits: 704,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    23,
			FractionalDigits: 705,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    2,
			FractionalDigits: 1,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    2,
			FractionalDigits: 1,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    57,
			FractionalDigits: 0,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    57,
			FractionalDigits: 0,
			IsNegativeValue:  true,
		},
		{
			IntegerDigits:    0,
			FractionalDigits: 1004,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    0,
			FractionalDigits: 1004,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    75,
			FractionalDigits: 1050,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    75,
			FractionalDigits: 1050,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 2,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 2,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 0,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 0,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 1050,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 1050,
			IsNegativeValue:  true,
		},
	}

	var expectedStr, testName,
		actualNumStr, numberSignDescription string

	lenTestProfile := len(testProfile)

	var err error

	for i := 0; i < lenTestProfile; i++ {

		expectedStr = getTestNativeNumStr(
			testProfile[i].IntegerDigits,
			testProfile[i].FractionalDigits,
			testProfile[i].IsNegativeValue)

		if testProfile[i].IsNegativeValue == true {

			numberSignDescription = "Negative Number"
		} else {
			numberSignDescription = "Positive Number"
		}

		testName = fmt.Sprintf("Test #%v TextFieldFormatDtoBigFloat\n"+
			"InputStr = (%v)\n"+
			"%v Test\n",
			i+1,
			expectedStr,
			numberSignDescription)

		txtFieldFmtDtoBFloat := TextFieldFormatDtoBigFloat{
			LeftMarginStr:         "",
			BigFloatNum:           big.Float{},
			LeadingMinusSign:      true,
			NativeRoundingMode:    big.AwayFromZero,
			RoundingType:          NumRoundType.HalfAwayFromZero(),
			NumOfFractionalDigits: -1,
			DefaultNumStrFmt:      NumStrFormatSpec{},
			FieldLength:           -1,
			FieldJustify:          TxtJustify.Right(),
			RightMarginStr:        "",
			lock:                  nil,
		}

		err = txtFieldFmtDtoBFloat.SetFromNativeNumStr(
			expectedStr,
			ePrefix.XCpy(
				"txtFieldFmtDtoBFloat<-expectedStr"))

		if err != nil {
			fmt.Printf("\n%v\n",
				err.Error())
			return
		}

		actualNumStr,
			err = new(MathHelper).NumericValueToNativeNumStr(
			txtFieldFmtDtoBFloat,
			ePrefix.XCpy(
				"actualNumStr<-txtFieldFmtDtoBFloat"))

		if err != nil {
			fmt.Printf("\n%v\n",
				err.Error())
			return
		}

		if actualNumStr != expectedStr {

			fmt.Printf("\n%v\n"+
				"%v\n"+
				"Error: actualNumStr != expectedStr\n"+
				"actualNumStr  = '%v'\n"+
				"expectedStr   = '%v'\n",
				ePrefix.String(),
				testName,
				actualNumStr,
				expectedStr)

			return

		}

	}
}

func TestMathHelper_NativeNumStrToNumericValue_0000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000100()",
		"")

	expectedStr := "12345678.12345678902"

	testName := fmt.Sprintf("Test #1 big.Float - bigFloat = (%v)\n",
		expectedStr)

	bigFloat := new(big.Float)

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		bigFloat,
		ePrefix.XCpy(
			"bigFloat<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr := bigFloat.Text('f', -1)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-12345678.12345678902"

	testName = fmt.Sprintf("Test #2 Negative Number big.Float - bigFloat = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		bigFloat,
		ePrefix.XCpy(
			"bigFloat<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = bigFloat.Text('f', -1)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	bigFloatConcrete := big.Float{}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		bigFloatConcrete,
		ePrefix.XCpy(
			"bigFloat<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter bigFloatConcrete is a concrete\n"+
			"value and NOT a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000200()",
		"")

	expectedStr := "123.456"

	testName := fmt.Sprintf("Test #1 float32 - float32Num = (%v)\n",
		expectedStr)

	var float32Num float32

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&float32Num,
		ePrefix.XCpy(
			"float32Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr := strconv.FormatFloat(
		float64(float32Num),
		'f',
		-1,
		32)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-123.456"

	testName = fmt.Sprintf("Test #2 Negative Number float32 - float32Num = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&float32Num,
		ePrefix.XCpy(
			"#2 float32Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = strconv.FormatFloat(
		float64(float32Num),
		'f',
		-1,
		32)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		float32Num,
		ePrefix.XCpy(
			"#2 float32Num<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter float32Num is a concrete\n"+
			"value and NOT a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000300()",
		"")

	expectedStr := "1234.56789012"

	testName := fmt.Sprintf("Test #1 float64 - float64Num = (%v)\n",
		expectedStr)

	var float64Num float64

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&float64Num,
		ePrefix.XCpy(
			"float64Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr := strconv.FormatFloat(
		float64Num,
		'f',
		-1,
		64)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1234.56789012"

	testName = fmt.Sprintf("Test #2 Negative Number float64 - float64Num = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&float64Num,
		ePrefix.XCpy(
			"#2 float64Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = strconv.FormatFloat(
		float64Num,
		'f',
		-1,
		64)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		float64Num,
		ePrefix.XCpy(
			"#2 float64Num<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter float64Num is a concrete\n"+
			"value and NOT a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000400()",
		"")

	expectedStr := "1234.56789012"

	testName := fmt.Sprintf("Test #1 BigFloatDto - bigFloatDto = (%v)\n",
		expectedStr)

	var bigFloatDto BigFloatDto

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&bigFloatDto,
		ePrefix.XCpy(
			"#1 bigFloatDto<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr :=
		bigFloatDto.Value.Text('f', -1)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1234.567890123456"

	testName = fmt.Sprintf("Test #2 Negative Number BigFloatDto - bigFloatDto = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&bigFloatDto,
		ePrefix.XCpy(
			"#2 bigFloatDto<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr =
		bigFloatDto.Value.Text('f', -1)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		bigFloatDto,
		ePrefix.XCpy(
			"#2 bigFloatDto<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter bigFloatDto is a concrete\n"+
			"value and NOT a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000500()",
		"")

	expectedStr := "1234.56789012"

	testName := fmt.Sprintf("Test #1 TextFieldFormatDtoFloat64 - txtFieldFmtDtoF64 = (%v)\n",
		expectedStr)

	txtFieldFmtDtoF64 := TextFieldFormatDtoFloat64{
		LeftMarginStr:         "",
		Float64Num:            0,
		LeadingMinusSign:      true,
		RoundingType:          NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          TxtJustify.Right(),
		RightMarginStr:        "",
	}

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&txtFieldFmtDtoF64,
		ePrefix.XCpy(
			"#1 txtFieldFmtDtoF64<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = txtFieldFmtDtoF64.
		FmtNumStrNative(
			ePrefix.XCpy(
				"#1 actualNumStr<-txtFieldFmtDtoF64"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1234.567890123456"

	testName = fmt.Sprintf("Test #2 Negative Number TextFieldFormatDtoFloat64 - txtFieldFmtDtoF64 = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&txtFieldFmtDtoF64,
		ePrefix.XCpy(
			"#2 txtFieldFmtDtoF64<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr,
		err = txtFieldFmtDtoF64.
		FmtNumStrNative(
			ePrefix.XCpy(
				"#2 actualNumStr<-txtFieldFmtDtoF64"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		txtFieldFmtDtoF64,
		ePrefix.XCpy(
			"#3 txtFieldFmtDtoF64<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter txtFieldFmtDtoF64 is a concrete\n"+
			"value and NOT a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000600()",
		"")

	expectedStr := "1234567.890123456789992345678"

	testName := fmt.Sprintf("Test #1 TextFieldFormatDtoBigFloat - txtFieldFmtDtoBFloat = (%v)\n",
		expectedStr)

	txtFieldFmtDtoBFloat := TextFieldFormatDtoBigFloat{
		LeftMarginStr:         "",
		BigFloatNum:           big.Float{},
		LeadingMinusSign:      true,
		NativeRoundingMode:    big.AwayFromZero,
		RoundingType:          NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: -1,
		DefaultNumStrFmt:      NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          TxtJustify.Right(),
		RightMarginStr:        "",
		lock:                  nil,
	}

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&txtFieldFmtDtoBFloat,
		ePrefix.XCpy(
			"#1 txtFieldFmtDtoBFloat<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr,
		err = txtFieldFmtDtoBFloat.
		FmtNumStrNative(
			ePrefix.XCpy(
				"#1 actualNumStr<-txtFieldFmtDtoBFloat"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1234567.890123456789012345678"

	testName = fmt.Sprintf("Test #2 Negative Number TextFieldFormatDtoBigFloat - "+
		"txtFieldFmtDtoBFloat = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&txtFieldFmtDtoBFloat,
		ePrefix.XCpy(
			"#2 txtFieldFmtDtoBFloat<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr,
		err = txtFieldFmtDtoBFloat.
		FmtNumStrNative(
			ePrefix.XCpy(
				"#2 actualNumStr<-txtFieldFmtDtoBFloat"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		txtFieldFmtDtoBFloat,
		ePrefix.XCpy(
			"#3 txtFieldFmtDtoBFloat<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter txtFieldFmtDtoBFloat is a concrete\n"+
			"value and NOT a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000700()",
		"")

	testProfile := []testNativeNumStrProfile{
		{
			IntegerDigits:    1,
			FractionalDigits: 0,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    1,
			FractionalDigits: 0,
			IsNegativeValue:  true,
		},
		{
			IntegerDigits:    8,
			FractionalDigits: 12,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    8,
			FractionalDigits: 12,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    5,
			FractionalDigits: 19,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    5,
			FractionalDigits: 19,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    50,
			FractionalDigits: 0,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    50,
			FractionalDigits: 0,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    14,
			FractionalDigits: 157,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    14,
			FractionalDigits: 157,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    1,
			FractionalDigits: 283,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    1,
			FractionalDigits: 283,
			IsNegativeValue:  true,
		},
		{
			IntegerDigits:    276,
			FractionalDigits: 14,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    276,
			FractionalDigits: 14,
			IsNegativeValue:  true,
		},
		{
			IntegerDigits:    383,
			FractionalDigits: 512,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    383,
			FractionalDigits: 512,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    9,
			FractionalDigits: 902,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    9,
			FractionalDigits: 902,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    23,
			FractionalDigits: 704,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    23,
			FractionalDigits: 705,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    2,
			FractionalDigits: 1,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    2,
			FractionalDigits: 1,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    57,
			FractionalDigits: 0,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    57,
			FractionalDigits: 0,
			IsNegativeValue:  true,
		},
		{
			IntegerDigits:    0,
			FractionalDigits: 1004,
			IsNegativeValue:  false,
		},
		{
			IntegerDigits:    0,
			FractionalDigits: 1004,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    75,
			FractionalDigits: 1050,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    75,
			FractionalDigits: 1050,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 2,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 2,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 0,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 0,
			IsNegativeValue:  true,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 1050,
			IsNegativeValue:  false,
		}, {
			IntegerDigits:    1050,
			FractionalDigits: 1050,
			IsNegativeValue:  true,
		},
	}

	var expectedStr, testName,
		actualNumStr, numberSignDescription string

	lenTestProfile := len(testProfile)

	var err error

	for i := 0; i < lenTestProfile; i++ {

		expectedStr = getTestNativeNumStr(
			testProfile[i].IntegerDigits,
			testProfile[i].FractionalDigits,
			testProfile[i].IsNegativeValue)

		if testProfile[i].IsNegativeValue == true {

			numberSignDescription = "Negative Number"
		} else {
			numberSignDescription = "Positive Number"
		}

		testName = fmt.Sprintf("Test #%v TextFieldFormatDtoBigFloat\n"+
			"InputStr = (%v)\n"+
			"%v Test\n",
			i+1,
			expectedStr,
			numberSignDescription)

		txtFieldFmtDtoBFloat := TextFieldFormatDtoBigFloat{
			LeftMarginStr:         "",
			BigFloatNum:           big.Float{},
			LeadingMinusSign:      true,
			NativeRoundingMode:    big.AwayFromZero,
			RoundingType:          NumRoundType.HalfAwayFromZero(),
			NumOfFractionalDigits: -1,
			DefaultNumStrFmt:      NumStrFormatSpec{},
			FieldLength:           -1,
			FieldJustify:          TxtJustify.Right(),
			RightMarginStr:        "",
			lock:                  nil,
		}

		err = new(MathHelper).NativeNumStrToNumericValue(
			expectedStr,
			&txtFieldFmtDtoBFloat,
			ePrefix.XCpy(
				fmt.Sprintf("idx=%v txtFieldFmtDtoBFloat<-expectedStr",
					i)))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		actualNumStr,
			err = txtFieldFmtDtoBFloat.
			FmtNumStrNative(
				ePrefix.XCpy(
					"#1 actualNumStr<-txtFieldFmtDtoBFloat"))

		if err != nil {
			t.Errorf("\n%v\n",
				err.Error())
			return
		}

		if actualNumStr != expectedStr {

			t.Errorf("\n%v\n"+
				"%v\n"+
				"Error: actualNumStr != expectedStr\n"+
				"actualNumStr  = '%v'\n"+
				"expectedStr   = '%v'\n",
				ePrefix.String(),
				testName,
				actualNumStr,
				expectedStr)

			return

		}

	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0000800(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000800()",
		"")

	expectedStr := "122"

	testName := fmt.Sprintf("Test #1 int8 - int8Num = (%v)\n",
		expectedStr)

	int8Num := int8(0)

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&int8Num,
		ePrefix.XCpy(
			"#1 int8Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr = strconv.FormatInt(int64(int8Num), 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-122"
	int8Num = 0

	testName = fmt.Sprintf("Test #2 Negative Number int8 - "+
		"int8Num = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&int8Num,
		ePrefix.XCpy(
			"#2 int8Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = strconv.FormatInt(int64(int8Num), 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		int8Num,
		ePrefix.XCpy(
			"#3 int8Num<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter int8Num is a concrete value and NOT\n"+
			"a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0000900(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0000900()",
		"")

	expectedStr := "1225"

	testName := fmt.Sprintf("Test #1 int16 - int16Num = (%v)\n",
		expectedStr)

	int16Num := int16(0)

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&int16Num,
		ePrefix.XCpy(
			"#1 int16Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr = strconv.FormatInt(int64(int16Num), 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1225"
	int16Num = 0

	testName = fmt.Sprintf("Test #2 Negative Number int8 - "+
		"int16Num = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&int16Num,
		ePrefix.XCpy(
			"#2 int16Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = strconv.FormatInt(int64(int16Num), 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		int16Num,
		ePrefix.XCpy(
			"#3 int16Num<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter int16Num is a concrete value and NOT\n"+
			"a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0001000(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0001000()",
		"")

	expectedStr := "1225397"

	testName := fmt.Sprintf("Test #1 int - intNum = (%v)\n",
		expectedStr)

	intNum := 0

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&intNum,
		ePrefix.XCpy(
			"#1 intNum<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr = strconv.FormatInt(int64(intNum), 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1225397"
	intNum = 0

	testName = fmt.Sprintf("Test #2 Negative Number int8 - "+
		"intNum = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&intNum,
		ePrefix.XCpy(
			"#2 intNum<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = strconv.FormatInt(int64(intNum), 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		intNum,
		ePrefix.XCpy(
			"#3 intNum<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter intNum is a concrete value and NOT\n"+
			"a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0001100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0001100()",
		"")

	expectedStr := "1225897"

	testName := fmt.Sprintf("Test #1 int32 - int32Num = (%v)\n",
		expectedStr)

	int32Num := int32(0)

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&int32Num,
		ePrefix.XCpy(
			"#1 int32Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr = strconv.FormatInt(int64(int32Num), 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1225897"
	int32Num = 0

	testName = fmt.Sprintf("Test #2 Negative Number int8 - "+
		"int32Num = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&int32Num,
		ePrefix.XCpy(
			"#2 int32Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = strconv.FormatInt(int64(int32Num), 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		int32Num,
		ePrefix.XCpy(
			"#3 int32Num<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter int32Num is a concrete value and NOT\n"+
			"a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0001200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0001200()",
		"")

	expectedStr := "1225897"

	testName := fmt.Sprintf("Test #1 int64 - int64Num = (%v)\n",
		expectedStr)

	int64Num := int64(0)

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&int64Num,
		ePrefix.XCpy(
			"#1 int64Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr = strconv.FormatInt(int64Num, 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1225897"
	int64Num = 0

	testName = fmt.Sprintf("Test #2 Negative Number int8 - "+
		"int64Num = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&int64Num,
		ePrefix.XCpy(
			"#2 int64Num<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = strconv.FormatInt(int64Num, 10)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		int64Num,
		ePrefix.XCpy(
			"#3 int64Num<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter int64Num is a concrete value and NOT\n"+
			"a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestMathHelper_NativeNumStrToNumericValue_0001400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestMathHelper_NativeNumStrToNumericValue_0001400()",
		"")

	expectedStr := "1225897123456789"

	testName := fmt.Sprintf("Test #1 big.Int - bigIntNum = (%v)\n",
		expectedStr)

	bigIntNum := big.Int{}

	var err error

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&bigIntNum,
		ePrefix.XCpy(
			"#1 bigIntNum<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumStr string

	actualNumStr = bigIntNum.Text(10)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	expectedStr = "-1225897123456789"

	bigIntNum = big.Int{}

	testName = fmt.Sprintf("Test #2 Negative Number int8 - "+
		"bigIntNum = (%v)\n",
		expectedStr)

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		&bigIntNum,
		ePrefix.XCpy(
			"#2 bigIntNum<-expectedStr"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumStr = bigIntNum.Text(10)

	if actualNumStr != expectedStr {

		t.Errorf("\n%v\n"+
			"%v\n"+
			"Error: actualNumStr != expectedStr\n"+
			"actualNumStr  = '%v'\n"+
			"expectedStr   = '%v'\n",
			ePrefix.String(),
			testName,
			actualNumStr,
			expectedStr)

		return

	}

	err = new(MathHelper).NativeNumStrToNumericValue(
		expectedStr,
		bigIntNum,
		ePrefix.XCpy(
			"#3 bigIntNum<-expectedStr"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Expected an error return from NativeNumStrToNumericValue()\n"+
			"because input parameter bigIntNum is a concrete value and NOT\n"+
			"a pointer. HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

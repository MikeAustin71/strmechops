package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextFieldFormatDtoFloat64_FmtNumStr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldFormatDtoFloat64_FmtNumStr_000100",
		"")

	txtFmtFloat64 := TextFieldFormatDtoFloat64{
		LeftMarginStr:         "",
		Float64Num:            12345.123456789,
		LeadingMinusSign:      true,
		RoundingType:          NumRoundType.HalfAwayFromZero(),
		NumOfFractionalDigits: 6,
		DefaultNumStrFmt:      NumStrFormatSpec{},
		FieldLength:           -1,
		FieldJustify:          TxtJustify.Right(),
		RightMarginStr:        "",
		lock:                  nil,
	}

	numStrFmtSpec,
		err := new(NumStrFormatSpec).NewNumFmtParams(
		".",
		",",
		IntGroupingType.Thousands(),
		"",
		"",
		NumFieldSymPos.InsideNumField(),
		"-",
		"",
		NumFieldSymPos.InsideNumField(),
		"",
		"",
		NumFieldSymPos.InsideNumField(),
		23,
		TxtJustify.Center(),
		ePrefix.XCpy(
			"numStrFmtSpec<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedStr := "     12,345.123457     "

	sMech := StrMech{}

	convertedExpectedStr := sMech.ConvertNonPrintableString(
		expectedStr,
		true)

	err = txtFmtFloat64.SetDefaultNumStrFmt(
		numStrFmtSpec,
		ePrefix.XCpy(
			"txtFmtFloat64<-numStrFmtSpec"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = txtFmtFloat64.FmtDefaultNumStr(
		ePrefix.XCpy(
			"txtFmtFloat64"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	convertedActualStr := sMech.ConvertNonPrintableString(
		actualStr,
		true)

	if convertedActualStr != convertedExpectedStr {

		t.Errorf("%v\n"+
			"Test#1 - txtFmtFloat64.SetDefaultNumStrFmt()\n"+
			"Error: convertedActualStr NOT EQUAL TO convertedExpectedStr\n"+
			"    convertedActualStr = '%v'\n"+
			"convertedExpectedStr   = '%v'\n",
			ePrefix.String(),
			convertedActualStr,
			convertedExpectedStr)

		return

	}

	return
}

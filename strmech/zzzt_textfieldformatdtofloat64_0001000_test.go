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
		"",
		"",
		false,
		NumFieldSymPos.None(),
		23,
		TxtJustify.Center(),
		ePrefix.XCpy(
			"Test#1 numStrFmtSpec<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualNumFieldSpec NumStrNumberFieldSpec

	actualNumFieldSpec,
		err = numStrFmtSpec.GetNumberFieldSpec(
		ePrefix.XCpy("actualNumFieldSpec<-numStrFmtSpec"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumFieldSpec.GetNumFieldLength() != 23 {

		t.Errorf("\n%v\n"+
			"Test#1-A\n"+
			"Error: NumStrNumberFieldSpec.fieldLength is NOT EQUAL TO '23'\n"+
			"NumStrNumberFieldSpec.fieldLength = %v\n",
			ePrefix.String(),
			actualNumFieldSpec.GetNumFieldLength())

		return
	}

	if actualNumFieldSpec.GetNumFieldJustification() !=
		TxtJustify.Center() {

		t.Errorf("\n%v\n"+
			"Test#1-B"+
			"Error: NumStrNumberFieldSpec.fieldJustification is NOT EQUAL TO 'Center'\n"+
			"NumStrNumberFieldSpec.fieldJustification = %v\n",
			ePrefix.String(),
			actualNumFieldSpec.GetNumFieldJustification().XValueInt())

		return

	}

	expectedStr := "     12,345.123457     "

	sMech := StrMech{}

	convertedExpectedStr := sMech.ConvertNonPrintableString(
		expectedStr,
		true)

	err = txtFmtFloat64.SetNumStrFmtDefault(
		numStrFmtSpec,
		ePrefix.XCpy(
			"Test #2 txtFmtFloat64<-numStrFmtSpec"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualDefaultFmt NumStrFormatSpec

	actualDefaultFmt,
		err = txtFmtFloat64.GetDefaultNumStrFmtSpec(
		ePrefix.XCpy("Test #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumFieldSpec,
		err = actualDefaultFmt.GetNumberFieldSpec(
		ePrefix.XCpy("Test #2-A"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if actualNumFieldSpec.GetNumFieldLength() != 23 {

		t.Errorf("\n%v\n"+
			"Test#2-B\n"+
			"Error: NumStrNumberFieldSpec.fieldLength is NOT EQUAL TO '23'\n"+
			"NumStrNumberFieldSpec.fieldLength = %v\n",
			ePrefix.String(),
			actualNumFieldSpec.GetNumFieldLength())

		return
	}

	if actualNumFieldSpec.GetNumFieldJustification() !=
		TxtJustify.Center() {

		t.Errorf("\n%v\n"+
			"Test#2-C"+
			"Error: NumStrNumberFieldSpec.fieldJustification is NOT EQUAL TO 'Center'\n"+
			"NumStrNumberFieldSpec.fieldJustification = %v\n",
			ePrefix.String(),
			actualNumFieldSpec.GetNumFieldJustification().XValueInt())

		return

	}

	var actualStr string

	actualStr,
		err = txtFmtFloat64.FmtNumStrDefault(
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
			"Test#1 - txtFmtFloat64.SetNumStrFmtDefault()\n"+
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

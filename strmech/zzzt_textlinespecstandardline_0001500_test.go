package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecStandardLine_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	stdLine02,
		err := createTestTextLineSpecStandardLine04(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine01 != stdLine02.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

}

func TestTextLineSpecStandardLine_Equal_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	stdLine02.newLineChars = []rune{'\n', '\n'}

	areEqual = stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 \n"+
			"Expected stdLine01 != stdLine02\n"+
			"because 'newLineChars' are different.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000300()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	stdLine02.numOfStdLines = 2

	areEqual = stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 \n"+
			"Expected stdLine01 != stdLine02\n"+
			"because 'numOfStdLines' are different.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000400()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	stdLine02.turnLineTerminatorOff = true

	areEqual = stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 \n"+
			"Expected stdLine01 != stdLine02\n"+
			"because 'turnLineTerminatorOff' \n"+
			"values are different.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000500()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.Equal(
		&stdLine02)

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	err = stdLine02.DeleteAtIndex(
		2,
		ePrefix.XCtx(
			"stdLine02"))

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	areEqual = stdLine01.Equal(
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #2 \n"+
			"Expected stdLine01 != stdLine02\n"+
			"because 'Text Field Arrays' \n"+
			"are different.\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_EqualITextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_EqualITextLine_000100()",
		"")

	expectedLeftMarginChars := []rune{' ', ' ', ' '}
	expectedRightMarginChars := []rune{' ', ' ', ' '}
	expectedNewLineChars := []rune{'\n', '\n'}

	expectedTextString := "How now brown cow!"

	plainTextLine01,
		err := TextLineSpecPlainText{}.NewPtrPlainText(
		expectedLeftMarginChars,
		expectedRightMarginChars,
		expectedTextString,
		expectedNewLineChars,
		false,
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine01.IsValidInstanceError(
		ePrefix.XCtx("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var stdLine01 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine01.EqualITextLine(
		plainTextLine01)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: stdLine01.EqualITextLine()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'plainTextLine01' is the wrong type.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual = stdLine01.EqualITextLine(
		&stdLine02)

	if areEqual == false {

		t.Errorf("%v\n"+
			"Error: stdLine01.EqualITextLine()\n"+
			"Expected areEqual == true because input parameter\n"+
			"'stdLine02' is identical to 'stdLine01'.\n"+
			"HOWEVER, areEqual == false\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_EqualITextLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTESTSERIES_TESTMETHOD_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	areEqual := stdLine01.EqualITextLine(
		nil)

	if areEqual == true {

		t.Errorf("%v\n"+
			"Error: stdLine01.EqualITextLine()\n"+
			"Expected areEqual == false because input\n"+
			"parameter 'iTextLine' is nil.\n"+
			"HOWEVER, areEqual == true\n",
			ePrefix.XCtxEmpty().String())

		return

	}
}

func TestTextLineSpecStandardLine_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetFormattedText_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		err := stdLine01.GetFormattedText(
		ePrefix.XCtx(
			"stdLine01"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: stdLine01.GetFormattedText()\n"+
			"Expected an error return because parameter\n"+
			"'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

}

func TestTextLineSpecStandardLine_GetFormattedText_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetFormattedText_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.GetFormattedText(
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from\n"+
			"TextLineSpecStandardLine{}.GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetNewLineChars_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetNewLineChars_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_ =
		stdLine01.GetNewLineChars()

	stdLine02,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNewLineChars := "\n\n"

	err = stdLine02.SetNewLineChars(
		expectedNewLineChars,
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNewLineChars :=
		stdLine02.GetNewLineChars()

	areEqual := expectedNewLineChars == actualNewLineChars

	if !areEqual {

		t.Errorf("%v\n"+
			"Error: stdLine02.GetNewLineChars()\n"+
			"expectedNewLineChars != actualNewLineChars\n"+
			"Expected areEqual == true\n"+
			"HOWEVER, areEqual == false\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetNewLineRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetNewLineRunes_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_ =
		stdLine01.GetNewLineRunes()

	expectedNewLineRunes := []rune{'\n', '\n'}

	stdLine02,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.SetNewLineRunes(
		expectedNewLineRunes,
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNewLineRunes :=
		stdLine02.GetNewLineRunes()

	areEqual := StrMech{}.Ptr().
		EqualRuneArrays(
			expectedNewLineRunes,
			actualNewLineRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"Error: stdLine02.GetNewLineRunes()\n"+
			"expectedNewLineRunes != actualNewLineRunes\n"+
			"Expected areEqual == true\n"+
			"HOWEVER, areEqual == false\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetNumOfStdLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetNumOfStdLines_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_ =
		stdLine01.GetNumOfStdLines()

	stdLine02,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumOfStdLines := 3

	stdLine02.SetNumOfStdLines(expectedNumOfStdLines)

	actualNumOfStdLines :=
		stdLine02.GetNumOfStdLines()

	if actualNumOfStdLines != expectedNumOfStdLines {

		t.Errorf("%v\n"+
			"Error: stdLine02.GetNumOfStdLines()\n"+
			"Expected expectedNumOfStdLines == actualNumOfStdLines\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n"+
			"Expected Number of Stdandard Lines = '%v'\n"+
			"  Actual Number of Stdandard Lines = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedNumOfStdLines,
			actualNumOfStdLines)

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetTextFields_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		err := stdLine01.GetTextFields(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textLines []ITextFieldSpecification

	textLines,
		err = stdLine02.GetTextFields(
		ePrefix.XCtx(
			"textLines<-stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}.New()

	err = stdLine03.SetTextFields(
		textLines,
		ePrefix.XCtx(
			"textLines->stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := stdLine02.Equal(&stdLine03)

	if !areEqual {

		t.Errorf("%v\n"+
			"Error: stdLine02.Equal(&stdLine03)\n"+
			"Expected stdLine02==stdLine03\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	_,
		err = stdLine03.GetTextFields(
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecStandardLine{}."+
			"GetTextFields()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetTextFields_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetTextFields_000200()",
		"")

	stdLine02,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02.textFields[2] = nil

	_,
		err = stdLine02.GetTextFields(
		ePrefix.XCtx(
			"textLines<-stdLine02"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLine02.GetTextFields()\n"+
			"Expected an error return.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

//func TestTESTSERIES_TESTMETHOD_000100(t *testing.T) {
//
//	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
//		"TestTESTSERIES_TESTMETHOD_000100()",
//		"")
//
//
//
//}

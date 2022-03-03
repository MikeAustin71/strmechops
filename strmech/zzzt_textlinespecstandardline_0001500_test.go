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
			"Expected an error return because\n"+
			"stdLine02.textFields[2] = nil .\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetTurnLineTerminatorOff_000100(t *testing.T) {

	stdLine01 := TextLineSpecStandardLine{}

	stdLine01.GetTurnLineTerminatorOff()

}

func TestTextLineSpecStandardLine_IsValidInstance_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_IsValidInstance_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	isValid :=
		stdLine01.IsValidInstance()

	stdLine02,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	isValid = stdLine02.IsValidInstance()

	if !isValid {

		t.Errorf("%v\n"+
			"Error: isValid = stdLine02.IsValidInstance()\n"+
			"Expected 'isValid' == true.\n"+
			"HOWEVER, 'isValid' == false !\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	stdLine02.textFields[2] = nil

	isValid = stdLine02.IsValidInstance()

	if isValid {

		t.Errorf("%v\n"+
			"Error: isValid = stdLine02.IsValidInstance()\n"+
			"Expected 'isValid' == false because \n"+
			"stdLine02.textFields[2] = nil .\n"+
			"HOWEVER, 'isValid' == true !\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	return
}

func TestTextLineSpecStandardLine_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_IsValidInstanceError_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	err :=
		stdLine01.IsValidInstanceError(
			ePrefix.XCtx(
				"stdLine01"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLine01.IsValidInstanceError()\n"+
			"Expected an error return because \n"+
			"stdLine01 is empty.\n"+
			"HOWEVER, 'isValid' == true !\n",
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

	err =
		stdLine02.IsValidInstanceError(
			textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLine02.IsValidInstanceError()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		stdLine02.IsValidInstanceError(
			ePrefix.XCtx(
				"stdLine02"))

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	stdLine02.textFields[3] = nil

	err =
		stdLine02.IsValidInstanceError(
			ePrefix.XCtx(
				"stdLine02.textFields[3]=nil"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLine02.IsValidInstanceError()\n"+
			"Expected an error return because\n"+
			"'stdLine02' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_NewStandardLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_NewStandardLine_000100()",
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

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			1,
			textFields,
			ePrefix.XCtx(
				"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine03 TextLineSpecStandardLine

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			1,
			nil,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLine()\n"+
			"Expected an error return because\n"+
			"'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textFields02 := make([]ITextFieldSpecification, 0)

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			1,
			textFields02,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLine()\n"+
			"Expected an error return because\n"+
			"'textFields' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			0,
			textFields02,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLine()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			-5,
			textFields,
			ePrefix.XCtx(
				"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecStandardLine_NewStandardLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_NewStandardLine_000200()",
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

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			1,
			textFields,
			textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"TextLineSpecStandardLine{}.NewStandardLine()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_NewPtrStandardLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_NewPtrStandardLine_000100()",
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

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 *TextLineSpecStandardLine

	stdLine02,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			1,
			textFields,
			ePrefix.XCtx(
				"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine03 *TextLineSpecStandardLine

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			1,
			nil,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLine()\n"+
			"Expected an error return because\n"+
			"'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textFields02 := make([]ITextFieldSpecification, 0)

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			1,
			textFields02,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLine()\n"+
			"Expected an error return because\n"+
			"'textFields02' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			0,
			textFields,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLine()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			-5,
			textFields,
			ePrefix.XCtx(
				"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecStandardLine_NewPtrStandardLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_NewPtrStandardLine_000200()",
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

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}
	_,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			1,
			textFields,
			textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLine()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_NewStandardLineAllParms_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_NewStandardLineAllParms_000100()",
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

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	newLineChars := []rune{'\n'}

	stdLine02,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			1,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine03 TextLineSpecStandardLine

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			1,
			nil,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			0,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textFields02 := make([]ITextFieldSpecification, 0)

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			1,
			textFields02,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'textFields02' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			0,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			-5,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineChars = []rune{}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			1,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"newLineChars is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'newLineChars' is an empty array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var stdLine04 TextLineSpecStandardLine

	newLineChars = []rune{'\n'}

	stdLine04,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			1,
			textFields,
			newLineChars,
			true,
			ePrefix.XCtx(
				"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine04.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	turnLineTerminatorOff :=
		stdLine04.GetTurnLineTerminatorOff()

	if turnLineTerminatorOff == false {
		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected 'turnLineTerminatorOff' == 'true'\n"+
			"HOWEVER, THE RETURN VALUE WAS 'false'!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	return
}

func TestTextLineSpecStandardLine_NewStandardLineAllParms_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_NewStandardLineAllParms_000200()",
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

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineChars := []rune{'\n'}

	_,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			1,
			textFields,
			newLineChars,
			false,
			textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_NewPtrStandardLineAllParms_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_NewPtrStandardLineAllParms_000100()",
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

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 *TextLineSpecStandardLine

	newLineChars := []rune{'\n'}

	stdLine02,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			1,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine03 *TextLineSpecStandardLine

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			1,
			nil,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			0,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	textFields02 := make([]ITextFieldSpecification, 0)

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			1,
			textFields02,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'textFields02' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			0,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			-5,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineChars = []rune{}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			1,
			textFields,
			newLineChars,
			false,
			ePrefix.XCtx(
				"newLineChars is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'newLineChars' is an empty array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var stdLine04 *TextLineSpecStandardLine

	newLineChars = []rune{'\n'}

	stdLine04,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			1,
			textFields,
			newLineChars,
			true,
			ePrefix.XCtx(
				"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine04.IsValidInstanceError(
		ePrefix.XCtx(
			"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	turnLineTerminatorOff :=
		stdLine04.GetTurnLineTerminatorOff()

	if turnLineTerminatorOff == false {
		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected 'turnLineTerminatorOff' == 'true'\n"+
			"HOWEVER, THE RETURN VALUE WAS 'false'!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	return
}

func TestTextLineSpecStandardLine_NewPtrStandardLineAllParms_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_NewPtrStandardLineAllParms_000200()",
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

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFields(
		ePrefix.XCtx(
			"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLineChars := []rune{'\n'}

	_,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			1,
			textFields,
			newLineChars,
			false,
			textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_PeekAtFirstTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PeekAtFirstTextField_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var iTxtFieldSpec ITextFieldSpecification

	iTxtFieldSpec,
		err = stdLine01.PeekAtFirstTextField(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecSpacer'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	spacerLen := spacerField.GetFieldLength()

	if spacerLen != 6 {

		t.Errorf("%v - Error\n"+
			"Expected the first index to be a type\n"+
			"TextFieldSpecSpacer with a length of '6' characters.\n"+
			"HOWEVER, A DIFFERENT LENGTH WAS RETURNED!\n"+
			"spacerLen='%v'\n",
			ePrefix.XCtxEmpty().String(),
			spacerLen)

		return
	}

	return
}

func TestTextLineSpecStandardLine_PeekAtFirstTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PeekAtFirstTextField_000200()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		err := stdLine01.PeekAtFirstTextField(
		ePrefix.XCtx(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.PeekAtFirstTextField()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
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

	_,
		err = stdLine02.PeekAtFirstTextField(
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"stdLine02.PeekAtFirstTextField()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_PeekAtLastTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PeekAtLastTextField_000100()",
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

	var iTxtFieldSpec ITextFieldSpecification

	iTxtFieldSpec,
		err = stdLine01.PeekAtLastTextField(
		ePrefix.XCtx(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecSpacer'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	spacerLen := spacerField.GetFieldLength()

	if spacerLen != 3 {

		t.Errorf("%v - Error\n"+
			"Expected the last index to be a type\n"+
			"TextFieldSpecSpacer with a length of '3' characters.\n"+
			"HOWEVER, A DIFFERENT LENGTH WAS RETURNED!\n"+
			"spacerLen='%v'\n",
			ePrefix.XCtxEmpty().String(),
			spacerLen)

		return
	}

	return
}

func TestTextLineSpecStandardLine_PeekAtLastTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PeekAtLastTextField_000200()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		err := stdLine01.PeekAtLastTextField(
		ePrefix.XCtx(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.PeekAtLastTextField()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
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

	_,
		err = stdLine02.PeekAtLastTextField(
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"stdLine02.PeekAtLastTextField()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

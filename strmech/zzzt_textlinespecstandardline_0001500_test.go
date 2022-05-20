package strmech

import (
	"fmt"
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
			ePrefix.String())

		return

	}

}

func TestTextLineSpecStandardLine_Equal_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

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
			ePrefix.String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000300()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

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
			ePrefix.String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000400()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

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
			ePrefix.String())

		return

	}
}

func TestTextLineSpecStandardLine_Equal_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_Equal_000500()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

		return

	}

	err = stdLine02.DeleteAtIndex(
		2,
		ePrefix.XCpy(
			"stdLine02"))

	if areEqual != true {

		t.Errorf("\n%v - ERROR\n"+
			"Test #1 \n"+
			"Expected stdLine01 == stdLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

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
			ePrefix.String())

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
		ePrefix.XCpy("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = plainTextLine01.IsValidInstanceError(
		ePrefix.XCpy("plainTextLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var stdLine01 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: stdLine01.GetFormattedText()\n"+
			"Expected an error return because parameter\n"+
			"'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

}

func TestTextLineSpecStandardLine_GetFormattedText_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetFormattedText_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNewLineChars := "\n\n"

	err = stdLine02.SetNewLineChars(
		expectedNewLineChars,
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.SetNewLineRunes(
		expectedNewLineRunes,
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
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
			ePrefix.String(),
			expectedNumOfStdLines,
			actualNumOfStdLines)

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetTextField_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		err := stdLine01.GetTextField(
		0,
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine01.GetTextField()\n"+
			"because the Text Fields collection is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var iTextField ITextFieldSpecification

	iTextField,
		err = stdLine02.GetTextField(
		1,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	fieldName := iTextField.TextFieldName()

	if fieldName != "Label" {

		t.Errorf("%v - ERROR\n"+
			"Expected Field Name for returned field index 1\n"+
			"would be 'Label'\n"+
			"HOWEVER, A DIFFERNT FIELD NAME WAS RETURNED!\n"+
			"Expected Field Name = 'Label'\n"+
			"  Actual Field Name = '%v'",
			ePrefix.String(),
			fieldName)

		return
	}

	var stdLine03 TextLineSpecStandardLine

	stdLine03,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine03.GetTextField(
		1,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03."+
			"GetTextField()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err = stdLine03.GetTextField(
		-1,
		ePrefix.XCpy(
			"stdLine03 index=-1"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03."+
			"GetTextField()\n"+
			"because input parameter 'zeroBasedIndex' has\n"+
			"a value of minus one (-1) and is therefore invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err = stdLine03.GetTextField(
		999,
		ePrefix.XCpy(
			"stdLine03 index=999"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from stdLine03."+
			"GetTextField()\n"+
			"because input parameter 'zeroBasedIndex' has\n"+
			"a value of 999 and is therefore invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetTextFieldCollection_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_GetTextFieldCollection_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		err := stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textLines []ITextFieldSpecification

	textLines,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textLines<-stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine03 := TextLineSpecStandardLine{}.New()

	err = stdLine03.SetTextFields(
		textLines,
		ePrefix.XCpy(
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
			ePrefix.String())

		return
	}

	_,
		err = stdLine03.GetTextFieldCollection(
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextLineSpecStandardLine{}."+
			"GetTextFieldCollection()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_TestTextLineSpecStandardLine_GetTextFieldCollection_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_TestTextLineSpecStandardLine_GetTextFieldCollection_000200()",
		"")

	stdLine02,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02.textFields[2] = nil

	_,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textLines<-stdLine02"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLine02.GetTextFieldCollection()\n"+
			"Expected an error return because\n"+
			"stdLine02.textFields[2] = nil .\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_GetTurnLineTerminatorOff_000100(t *testing.T) {

	stdLine01 := TextLineSpecStandardLine{}

	stdLine01.GetTurnLineTerminatorOff()

}

func TestTextLineSpecStandardLine_InsertTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_InsertTextField_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedLabelText := "Xray7 where are?"

	var labelTxt TextFieldSpecLabel

	labelTxt,
		err = TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedLastIndex :=
		stdLine01.GetNumOfTextFields()

	var lastIndexId int

	lastIndexId,
		err = stdLine01.InsertTextField(
		&labelTxt,
		2,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01 after insert"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedLastIndex != lastIndexId {

		t.Errorf("%v - Error\n"+
			"stdLine01.InsertTextField()\n"+
			"Expected Last Index Id did NOT match\n"+
			"Actual Last Index Id.\n"+
			"Expected Last Index Id = '%v'\n"+
			"  Actual Last Index Id = '%v'\n",
			ePrefix.String(),
			expectedLastIndex,
			lastIndexId)

		return
	}

	// var labelSpec TextFieldSpecSpacer
	var iTxtFieldSpec ITextFieldSpecification

	iTxtFieldSpec,
		err = stdLine01.PeekAtTextFieldAtIndex(
		2,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelField,
		ok := iTxtFieldSpec.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecLabel'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	err = actualLabelField.IsValidInstanceError(
		ePrefix.XCpy(
			"labelField"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelText := actualLabelField.GetTextLabel()

	if expectedLabelText != actualLabelText {

		t.Errorf("%v - Error\n"+
			"Expected expectedLabelText==actualLabelText\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedLabelText,
			actualLabelText)

		return
	}

	expectedLastIndex++

	expectedTxtFieldLen := expectedLastIndex

	actualTxtFieldLen :=
		stdLine01.GetNumOfTextFields()

	if expectedTxtFieldLen != actualTxtFieldLen {

		t.Errorf("%v - Error\n"+
			"Expected expectedTxtFieldLen==actualTxtFieldLen\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedTxtFieldLen= '%v'\n"+
			"  actualTxtFieldLen= '%v'\n",
			ePrefix.String(),
			expectedTxtFieldLen,
			actualTxtFieldLen)

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lastIdx := stdLine02.GetNumOfTextFields() - 1

	lastIndexId,
		err = stdLine02.InsertTextField(
		&labelTxt,
		lastIdx,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtField ITextFieldSpecification

	txtField,
		err = stdLine02.GetTextField(
		lastIdx,
		ePrefix.XCpy(
			"stdLine02"))

	actualLabelField,
		ok = txtField.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"Text #2\n"+
			"actualLabelField, ok := txtField.(*TextFieldSpecLabel)\n"+
			"Expected return of type 'TextFieldSpecLabel'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	err = actualLabelField.IsValidInstanceError(
		ePrefix.XCpy(
			"labelField #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelText = actualLabelField.GetTextLabel()

	if expectedLabelText != actualLabelText {

		t.Errorf("%v - Error\n"+
			"Expected expectedLabelText==actualLabelText\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedLabelText,
			actualLabelText)

		return
	}

	return
}

func TestTextLineSpecStandardLine_InsertTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_InsertTextField_000200()",
		"")

	expectedLabelText := "Xray7 where are?"

	var labelTxt TextFieldSpecLabel

	labelTxt,
		err := TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine01 := TextLineSpecStandardLine{}

	_,
		err = stdLine01.InsertTextField(
		&labelTxt,
		2,
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.InsertTextField()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine02.InsertTextField(
		&labelTxt,
		2,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"stdLine02.InsertTextField()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_InsertTextField_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_InsertTextField_000300()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine05(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	expectedLabelText := "Car 54 where are you?"

	var expectedTxtLabelField TextFieldSpecLabel

	expectedTxtLabelField,
		err = TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"expectedTxtLabelField"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedNumTxtFields :=
		stdLine01.GetNumOfTextFields() + 1

	initialTargetIndex := 0

	var iOldTxtField ITextFieldSpecification

	iOldTxtField,
		err = stdLine01.GetTextField(
		initialTargetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"iOldTxtField<-stdLine01[%v]",
				initialTargetIndex)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.InsertTextField(
		&expectedTxtLabelField,
		initialTargetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"stdLine01[%v]<-expectedTxtLabelField",
				initialTargetIndex)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01 after insertion"))

	actualNumTxtFields :=
		stdLine01.GetNumOfTextFields()

	if expectedNumTxtFields != actualNumTxtFields {

		t.Errorf("%v - Error\n"+
			"stdLine01.InsertTextField()\n"+
			"Expected Number Of Text Fields\n"+
			"DID NOT MATCH the Actual Number\n"+
			"of Text Fields.\n"+
			"Expected Number of Text Fields = '%v'\n"+
			"  Actual Number of Text Fields = '%v'\n",
			ePrefix.String(),
			expectedNumTxtFields,
			actualNumTxtFields)

		return
	}

	var iTargetTxtField ITextFieldSpecification

	iTargetTxtField,
		err = stdLine01.GetTextField(
		initialTargetIndex,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelField,
		ok := iTargetTxtField.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTargetTxtField.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecLabel'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	if !expectedTxtLabelField.Equal(actualLabelField) {

		t.Errorf("%v - Error\n"+
			"Expected expectedTxtLabelField==actualLabelField\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedTxtLabelField.GetTextLabel(),
			actualLabelField.GetTextLabel())

		return

	}

	var iOldTxtField2 ITextFieldSpecification

	iOldTxtField2,
		err = stdLine01.GetTextField(
		initialTargetIndex+1,
		ePrefix.XCpy(
			fmt.Sprintf(
				"iOldTxtField<-stdLine01[%v]",
				initialTargetIndex+1)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !iOldTxtField.EqualITextField(
		iOldTxtField2) {

		t.Errorf("%v - Error\n"+
			"Expected original stdLine01[%v] \n"+
			"would EQUAL stdLine01[%v] after insertion.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String(),
			initialTargetIndex,
			initialTargetIndex+1)

		return

	}

	return
}

func TestTextLineSpecStandardLine_InsertTextField_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_InsertTextField_000400()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine05(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	expectedLabelText := "The answer is 42!"

	var expectedTxtLabelField TextFieldSpecLabel

	expectedTxtLabelField,
		err = TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"expectedTxtLabelField"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedNumTxtFields :=
		stdLine01.GetNumOfTextFields() + 1

	initialTargetIndex := 4

	var iOldTxtField ITextFieldSpecification

	iOldTxtField,
		err = stdLine01.GetTextField(
		initialTargetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"iOldTxtField<-stdLine01[%v]",
				initialTargetIndex)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.InsertTextField(
		&expectedTxtLabelField,
		initialTargetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"stdLine01[%v]<-expectedTxtLabelField",
				initialTargetIndex)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01 after insertion"))

	actualNumTxtFields :=
		stdLine01.GetNumOfTextFields()

	if expectedNumTxtFields != actualNumTxtFields {

		t.Errorf("%v - Error\n"+
			"stdLine01.InsertTextField()\n"+
			"Expected Number Of Text Fields\n"+
			"DID NOT MATCH the Actual Number\n"+
			"of Text Fields.\n"+
			"Expected Number of Text Fields = '%v'\n"+
			"  Actual Number of Text Fields = '%v'\n",
			ePrefix.String(),
			expectedNumTxtFields,
			actualNumTxtFields)

		return
	}

	var iTargetTxtField ITextFieldSpecification

	iTargetTxtField,
		err = stdLine01.GetTextField(
		initialTargetIndex,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelField,
		ok := iTargetTxtField.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTargetTxtField.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecLabel'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	if !expectedTxtLabelField.Equal(actualLabelField) {

		t.Errorf("%v - Error\n"+
			"Expected expectedTxtLabelField==actualLabelField\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedTxtLabelField.GetTextLabel(),
			actualLabelField.GetTextLabel())

		return

	}

	var iOldTxtField2 ITextFieldSpecification

	iOldTxtField2,
		err = stdLine01.GetTextField(
		initialTargetIndex+1,
		ePrefix.XCpy(
			fmt.Sprintf(
				"iOldTxtField<-stdLine01[%v]",
				initialTargetIndex+1)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !iOldTxtField.EqualITextField(
		iOldTxtField2) {

		t.Errorf("%v - Error\n"+
			"Expected original stdLine01[%v] \n"+
			"would EQUAL stdLine01[%v] after insertion.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String(),
			initialTargetIndex,
			initialTargetIndex+1)

		return

	}

	return
}

func TestTextLineSpecStandardLine_InsertTextField_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_InsertTextField_000500()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine05(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	expectedLabelText := "Shiver me timbers!"

	var expectedTxtLabelField TextFieldSpecLabel

	expectedTxtLabelField,
		err = TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"expectedTxtLabelField"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedNumTxtFields :=
		stdLine01.GetNumOfTextFields() + 1

	initialTargetIndex := stdLine01.GetNumOfTextFields() - 1

	var iOldTxtField ITextFieldSpecification

	iOldTxtField,
		err = stdLine01.GetTextField(
		initialTargetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"iOldTxtField<-stdLine01[%v]",
				initialTargetIndex)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.InsertTextField(
		&expectedTxtLabelField,
		initialTargetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"stdLine01[%v]<-expectedTxtLabelField",
				initialTargetIndex)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01 after insertion"))

	actualNumTxtFields :=
		stdLine01.GetNumOfTextFields()

	if expectedNumTxtFields != actualNumTxtFields {

		t.Errorf("%v - Error\n"+
			"stdLine01.InsertTextField()\n"+
			"Expected Number Of Text Fields\n"+
			"DID NOT MATCH the Actual Number\n"+
			"of Text Fields.\n"+
			"Expected Number of Text Fields = '%v'\n"+
			"  Actual Number of Text Fields = '%v'\n",
			ePrefix.String(),
			expectedNumTxtFields,
			actualNumTxtFields)

		return
	}

	var iTargetTxtField ITextFieldSpecification

	iTargetTxtField,
		err = stdLine01.GetTextField(
		initialTargetIndex,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelField,
		ok := iTargetTxtField.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTargetTxtField.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecLabel'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	if !expectedTxtLabelField.Equal(actualLabelField) {

		t.Errorf("%v - Error\n"+
			"Expected expectedTxtLabelField==actualLabelField\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedTxtLabelField.GetTextLabel(),
			actualLabelField.GetTextLabel())

		return

	}

	var iOldTxtField2 ITextFieldSpecification

	iOldTxtField2,
		err = stdLine01.GetTextField(
		initialTargetIndex+1,
		ePrefix.XCpy(
			fmt.Sprintf(
				"iOldTxtField<-stdLine01[%v]",
				initialTargetIndex+1)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !iOldTxtField.EqualITextField(
		iOldTxtField2) {

		t.Errorf("%v - Error\n"+
			"Expected original stdLine01[%v] \n"+
			"would EQUAL stdLine01[%v] after insertion.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String(),
			initialTargetIndex,
			initialTargetIndex+1)

		return

	}

	return
}

func TestTextLineSpecStandardLine_InsertTextField_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_InsertTextField_000600()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine05(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	expectedLabelText := "Say it ain't so Joe!"

	var expectedTxtLabelField TextFieldSpecLabel

	expectedTxtLabelField,
		err = TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"expectedTxtLabelField"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedNumTxtFields :=
		stdLine01.GetNumOfTextFields() + 1

	initialTargetIndex := stdLine01.GetNumOfTextFields()

	var iOldTxtField ITextFieldSpecification

	iOldTxtField,
		err = stdLine01.GetTextField(
		initialTargetIndex-1,
		ePrefix.XCpy(
			fmt.Sprintf(
				"iOldTxtField<-stdLine01[%v]",
				initialTargetIndex-1)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.InsertTextField(
		&expectedTxtLabelField,
		initialTargetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"stdLine01[%v]<-expectedTxtLabelField",
				initialTargetIndex)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01 after insertion"))

	actualNumTxtFields :=
		stdLine01.GetNumOfTextFields()

	if expectedNumTxtFields != actualNumTxtFields {

		t.Errorf("%v - Error\n"+
			"stdLine01.InsertTextField()\n"+
			"Expected Number Of Text Fields\n"+
			"DID NOT MATCH the Actual Number\n"+
			"of Text Fields.\n"+
			"Expected Number of Text Fields = '%v'\n"+
			"  Actual Number of Text Fields = '%v'\n",
			ePrefix.String(),
			expectedNumTxtFields,
			actualNumTxtFields)

		return
	}

	var iTargetTxtField ITextFieldSpecification

	iTargetTxtField,
		err = stdLine01.GetTextField(
		initialTargetIndex,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelField,
		ok := iTargetTxtField.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTargetTxtField.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecLabel'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	if !expectedTxtLabelField.Equal(actualLabelField) {

		t.Errorf("%v - Error\n"+
			"Expected expectedTxtLabelField==actualLabelField\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedTxtLabelField.GetTextLabel(),
			actualLabelField.GetTextLabel())

		return

	}

	var iOldTxtField2 ITextFieldSpecification

	iOldTxtField2,
		err = stdLine01.GetTextField(
		initialTargetIndex-1,
		ePrefix.XCpy(
			fmt.Sprintf(
				"iOldTxtField<-stdLine01[%v]",
				initialTargetIndex-1)))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !iOldTxtField.EqualITextField(
		iOldTxtField2) {

		t.Errorf("%v - Error\n"+
			"Expected original stdLine01[%v] \n"+
			"would EQUAL stdLine01[%v] after insertion.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String(),
			initialTargetIndex,
			initialTargetIndex-1)

		return

	}

	return
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
		ePrefix.XCpy(
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
			ePrefix.String())

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
			ePrefix.String())

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
			ePrefix.XCpy(
				"stdLine01"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLine01.IsValidInstanceError()\n"+
			"Expected an error return because \n"+
			"stdLine01 is empty.\n"+
			"HOWEVER, 'isValid' == true !\n",
			ePrefix.String())

		return

	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

		return
	}

	err =
		stdLine02.IsValidInstanceError(
			ePrefix.XCpy(
				"stdLine02"))

	if err != nil {
		t.Errorf("%v",
			err.Error())

		return
	}

	stdLine02.textFields[3] = nil

	err =
		stdLine02.IsValidInstanceError(
			ePrefix.XCpy(
				"stdLine02.textFields[3]=nil"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLine02.IsValidInstanceError()\n"+
			"Expected an error return because\n"+
			"'stdLine02' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLine()\n"+
			"Expected an error return because\n"+
			"'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	textFields02 := make([]ITextFieldSpecification, 0)

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			1,
			textFields02,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLine()\n"+
			"Expected an error return because\n"+
			"'textFields' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			0,
			textFields02,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLine()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLine(
			-5,
			textFields,
			ePrefix.XCpy(
				"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCpy(
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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLine()\n"+
			"Expected an error return because\n"+
			"'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	textFields02 := make([]ITextFieldSpecification, 0)

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			1,
			textFields02,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLine()\n"+
			"Expected an error return because\n"+
			"'textFields02' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			0,
			textFields,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLine()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLine(
			-5,
			textFields,
			ePrefix.XCpy(
				"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCpy(
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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			0,
			textFields,
			newLineChars,
			false,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

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
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'textFields02' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			0,
			textFields,
			newLineChars,
			false,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewStandardLineAllParms(
			-5,
			textFields,
			newLineChars,
			false,
			ePrefix.XCpy(
				"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"newLineChars is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'newLineChars' is an empty array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

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
			ePrefix.XCpy(
				"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine04.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			0,
			textFields,
			newLineChars,
			false,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

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
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'textFields02' is a zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			0,
			textFields,
			newLineChars,
			false,
			ePrefix.XCpy(
				"stdLine03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'numOfStdLines' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

		return
	}

	stdLine03,
		err = TextLineSpecStandardLine{}.
		NewPtrStandardLineAllParms(
			-5,
			textFields,
			newLineChars,
			false,
			ePrefix.XCpy(
				"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"newLineChars is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"TextLineSpecStandardLine{}.NewPtrStandardLineAllParms()\n"+
			"Expected an error return because\n"+
			"'newLineChars' is an empty array.\n"+
			"HOWEVER, NO ERROR WAS RETURN!\n",
			ePrefix.String())

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
			ePrefix.XCpy(
				"stdLine04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine04.IsValidInstanceError(
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var iTxtFieldSpec ITextFieldSpecification

	iTxtFieldSpec,
		err = stdLine01.PeekAtFirstTextField(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var spacerField *TextFieldSpecSpacer

	var ok bool

	spacerField, ok = iTxtFieldSpec.(*TextFieldSpecSpacer)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecSpacer'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	spacerLen := spacerField.GetFieldLength()

	if spacerLen != 6 {

		t.Errorf("%v - Error\n"+
			"Expected the first index to be a type\n"+
			"TextFieldSpecSpacer with a length of '6' characters.\n"+
			"HOWEVER, A DIFFERENT LENGTH WAS RETURNED!\n"+
			"spacerLen='%v'\n",
			ePrefix.String(),
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
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.PeekAtFirstTextField()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

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
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var iTxtFieldSpec ITextFieldSpecification

	iTxtFieldSpec,
		err = stdLine01.PeekAtLastTextField(
		ePrefix.XCpy(
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
			ePrefix.String())

		return
	}

	spacerLen := spacerField.GetFieldLength()

	if spacerLen != 3 {

		t.Errorf("%v - Error\n"+
			"Expected the last index to be a type\n"+
			"TextFieldSpecSpacer with a length of '3' characters.\n"+
			"HOWEVER, A DIFFERENT LENGTH WAS RETURNED!\n"+
			"spacerLen='%v'\n",
			ePrefix.String(),
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
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.PeekAtLastTextField()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
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
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_PeekAtTextFieldAtIndex_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PeekAtTextFieldAtIndex_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	// var labelSpec TextFieldSpecSpacer
	var iTxtFieldSpec ITextFieldSpecification

	iTxtFieldSpec,
		err = stdLine01.PeekAtTextFieldAtIndex(
		4,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedLabelText := "Ain't gonna make war no more!"

	labelField, ok := iTxtFieldSpec.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecSpacer'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	err = labelField.IsValidInstanceError(
		ePrefix.XCpy(
			"labelField"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelText := labelField.GetTextLabel()

	if expectedLabelText != actualLabelText {

		t.Errorf("%v - Error\n"+
			"Expected expectedLabelText==actualLabelText\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedLabelText,
			actualLabelText)

		return
	}

	return
}

func TestTextLineSpecStandardLine_PeekAtTextFieldAtIndex_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PeekAtTextFieldAtIndex_000200()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		err := stdLine01.PeekAtTextFieldAtIndex(
		0,
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.PeekAtTextFieldAtIndex()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine02.PeekAtTextFieldAtIndex(
		2,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"stdLine02.PeekAtTextFieldAtIndex()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_PopFirstTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PopFirstTextField_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumOfTxtFields :=
		stdLine01.GetNumOfTextFields()

	expectedNumOfTxtFields--

	var iTxtFieldSpec ITextFieldSpecification
	var remainingNumOfTxtFields int

	iTxtFieldSpec,
		remainingNumOfTxtFields,
		err = stdLine01.PopFirstTextField(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumOfTxtFields != remainingNumOfTxtFields {

		t.Errorf("%v - Error\n"+
			"stdLine01.PopFirstTextField()\n"+
			"Expected Number of Remaining Text Fields is\n"+
			"NOT Equal to the Actual Number of Remaining\n"+
			"Text Fields!\n"+
			"Expected Number of Remaining Text Fields = '%v'\n"+
			"  Actual Number of Remaining Text Fields = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtFields,
			remainingNumOfTxtFields)

		return
	}

	spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)

	if !ok {

		t.Errorf("%v - Error\n"+
			"stdLine01.PopFirstTextField()\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecSpacer'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	spacerLen := spacerField.GetFieldLength()

	if spacerLen != 6 {

		t.Errorf("%v - Error\n"+
			"Expected the first index to be a type\n"+
			"TextFieldSpecSpacer with a length of '6' characters.\n"+
			"HOWEVER, A DIFFERENT LENGTH WAS RETURNED!\n"+
			"spacerLen='%v'\n",
			ePrefix.String(),
			spacerLen)

		return
	}

	return
}

func TestTextLineSpecStandardLine_PopFirstTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PopFirstTextField_000200()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		_,
		err := stdLine01.PopFirstTextField(
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.PopFirstTextField()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		_,
		err = stdLine02.PopFirstTextField(
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"stdLine02.PopFirstTextField()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_PopLastTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PopLastTextField_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumOfTxtFields :=
		stdLine01.GetNumOfTextFields()

	expectedNumOfTxtFields--

	var iTxtFieldSpec ITextFieldSpecification
	var remainingNumOfTxtFields int

	iTxtFieldSpec,
		remainingNumOfTxtFields,
		err = stdLine01.PopLastTextField(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedNumOfTxtFields != remainingNumOfTxtFields {

		t.Errorf("%v - Error\n"+
			"stdLine01.PopLastTextField()\n"+
			"Expected Number of Remaining Text Fields is\n"+
			"NOT Equal to the Actual Number of Remaining\n"+
			"Text Fields!\n"+
			"Expected Number of Remaining Text Fields = '%v'\n"+
			"  Actual Number of Remaining Text Fields = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtFields,
			remainingNumOfTxtFields)

		return
	}

	spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)

	if !ok {

		t.Errorf("%v - Error\n"+
			"stdLine01.PopLastTextField()\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecSpacer'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	spacerLen := spacerField.GetFieldLength()

	if spacerLen != 3 {

		t.Errorf("%v - Error\n"+
			"Expected the last index to be a type\n"+
			"TextFieldSpecSpacer with a length of '3' characters.\n"+
			"HOWEVER, A DIFFERENT LENGTH WAS RETURNED!\n"+
			"spacerLen='%v'\n",
			ePrefix.String(),
			spacerLen)

		return
	}

	return
}

func TestTextLineSpecStandardLine_PopLastTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PopLastTextField_000200()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		_,
		err := stdLine01.PopLastTextField(
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.PopLastTextField()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		_,
		err = stdLine02.PopLastTextField(
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"stdLine02.PopLastTextField()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLine_PopTextFieldAtIndex_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PopTextFieldAtIndex_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	// var labelSpec TextFieldSpecSpacer
	var iTxtFieldSpec ITextFieldSpecification
	var expectedNumOfTextFields,
		remainingNumOfTxtFields int

	expectedNumOfTextFields =
		stdLine01.GetNumOfTextFields()

	expectedNumOfTextFields--

	iTxtFieldSpec,
		remainingNumOfTxtFields,
		err = stdLine01.PopTextFieldAtIndex(
		4,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedLabelText := "Ain't gonna make war no more!"

	labelField, ok := iTxtFieldSpec.(*TextFieldSpecLabel)

	if !ok {

		t.Errorf("%v - Error\n"+
			"spacerField, ok := iTxtFieldSpec.(*TextFieldSpecSpacer)\n"+
			"Expected return of type 'TextFieldSpecSpacer'.\n"+
			"HOWEVER, THAT TYPE WAS NOT RETURNED!\n",
			ePrefix.String())

		return
	}

	err = labelField.IsValidInstanceError(
		ePrefix.XCpy(
			"labelField"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualLabelText := labelField.GetTextLabel()

	if expectedLabelText != actualLabelText {

		t.Errorf("%v - Error\n"+
			"stdLine01.PopTextFieldAtIndex()\n"+
			"Expected expectedLabelText==actualLabelText\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelText= '%v'\n"+
			"  actualLabelText= '%v'\n",
			ePrefix.String(),
			expectedLabelText,
			actualLabelText)

		return
	}

	if expectedNumOfTextFields != remainingNumOfTxtFields {

		t.Errorf("%v - Error\n"+
			"stdLine01.PopFirstTextField()\n"+
			"Expected Number of Remaining Text Fields is\n"+
			"NOT Equal to the Actual Number of Remaining\n"+
			"Text Fields!\n"+
			"Expected Number of Remaining Text Fields = '%v'\n"+
			"  Actual Number of Remaining Text Fields = '%v'\n",
			ePrefix.String(),
			expectedNumOfTextFields,
			remainingNumOfTxtFields)

		return
	}

	return
}

func TestTextLineSpecStandardLine_PopTextFieldAtIndex_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLine_PopTextFieldAtIndex_000200()",
		"")

	stdLine01 := TextLineSpecStandardLine{}

	_,
		_,
		err := stdLine01.PopTextFieldAtIndex(
		0,
		ePrefix.XCpy(
			"stdLine01"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"Expected an error return from stdLine01.PopTextFieldAtIndex()\n"+
			"because 'stdLine01' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		_,
		err = stdLine02.PopTextFieldAtIndex(
		2,
		textLineSpecStandardLineElectron{})

	if err == nil {

		t.Errorf("%v - Error\n"+
			"stdLine02.PopTextFieldAtIndex()\n"+
			"Expected an error return because\n"+
			"'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecLinesCollection_PeekAtFirstTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_PeekAtFirstTextLine_000100()",
		"")

	// Index 1
	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var txtLinesCol01 TextLineSpecLinesCollection

	_,
		txtLinesCol01,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var iTxtLineSpec ITextLineSpecification

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtFirstTextLine(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		&stdLine01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PeekAtFirstTextLine()\n"+
			"Expected ITextLineSpecifications object would be\n"+
			"Equal to Actual ITextLineSpecifications object."+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	err = txtLinesCol02.CopyIn(
		&txtLinesCol01,
		ePrefix.XCpy(
			"txtLinesCol02<=txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = txtLinesCol02.PeekAtFirstTextLine(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"PeekAtFirstTextLine()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	_,
		err = txtLinesCol03.PeekAtFirstTextLine(
		ePrefix.XCpy(
			"txtLinesCol03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"PeekAtFirstTextLine()\n"+
			"because 'txtLinesCol03' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol04 TextLineSpecLinesCollection

	_,
		txtLinesCol04,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol04.textLines[0] = nil

	_,
		err = txtLinesCol04.PeekAtFirstTextLine(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"PeekAtFirstTextLine()\n"+
			"because 'txtLinesCol04.textLines[0]' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol05 TextLineSpecLinesCollection

	_,
		txtLinesCol05,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol05.textLines[0].Empty()

	_,
		err = txtLinesCol05.PeekAtFirstTextLine(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol05."+
			"PeekAtFirstTextLine()\n"+
			"because 'txtLinesCol05.textLines[0]' is invalid.\n"+
			"txtLinesCol05.textLines[0].Empty() was previously called.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_PeekAtLastTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_PeekAtLastTextLine_000100()",
		"")

	lastIndexId,
		txtLinesCol01,
		err := createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lastIndexId--

	origITxtLineSpec := txtLinesCol01.textLines[lastIndexId]

	var iTxtLineSpec ITextLineSpecification

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		origITxtLineSpec) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PeekAtLastTextLine()\n"+
			"Expected ITextLineSpecifications object would be\n"+
			"Equal to Actual ITextLineSpecifications object."+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	err = txtLinesCol02.CopyIn(
		&txtLinesCol01,
		ePrefix.XCpy(
			"txtLinesCol02<=txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = txtLinesCol02.PeekAtLastTextLine(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"PeekAtLastTextLine()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	_,
		err = txtLinesCol03.PeekAtLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"PeekAtLastTextLine()\n"+
			"because 'txtLinesCol03' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol04 TextLineSpecLinesCollection

	_,
		txtLinesCol04,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol04.textLines[lastIndexId] = nil

	_,
		err = txtLinesCol04.PeekAtLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"PeekAtLastTextLine()\n"+
			"because 'txtLinesCol04.textLines[%v]' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			lastIndexId)

		return
	}

	var txtLinesCol05 TextLineSpecLinesCollection

	_,
		txtLinesCol05,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol05.textLines[lastIndexId].Empty()

	_,
		err = txtLinesCol05.PeekAtLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol05."+
			"PeekAtLastTextLine()\n"+
			"because 'txtLinesCol05.textLines[%v]' is invalid.\n"+
			"txtLinesCol05.textLines[%v].Empty() was previously called.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			lastIndexId,
			lastIndexId)

		return
	}

	return
}

func TestTextLineSpecLinesCollection_PeekAtTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_PeekAtTextLine_000100()",
		"")

	// Index 1
	stdLine02,
		err := createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine02"))

	var txtLinesCol01 TextLineSpecLinesCollection

	_,
		txtLinesCol01,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var iTxtLineSpec ITextLineSpecification

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtTextLine(
		1,
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		&stdLine02) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PeekAtTextLine(1)\n"+
			"Expected ITextLineSpecifications object would be\n"+
			"Equal to Actual ITextLineSpecifications object."+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	_,
		err = txtLinesCol02.PeekAtTextLine(
		-1,
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"PeekAtTextLine()\n"+
			"because input parameter 'zeroBaseIndex' has\n"+
			"a value of minus one (-1).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	err = txtLinesCol03.CopyIn(
		&txtLinesCol01,
		ePrefix.XCpy(
			"txtLinesCol03<=txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = txtLinesCol03.PeekAtTextLine(
		1,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"PeekAtTextLine()\n"+
			"because input parameter 'errorPrefix' has\n"+
			"is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol04 := TextLineSpecLinesCollection{}

	_,
		txtLinesCol04,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = txtLinesCol04.PeekAtTextLine(
		999,
		ePrefix.XCpy(
			"txtLinesCol04[999]"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"PeekAtTextLine()\n"+
			"because input parameter 'zeroBasedIndex' has\n"+
			"a value of '999' and is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_PeekAtTextLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_PeekAtTextLine_000100()",
		"")

	expectedNumOfTxtLines,
		txtLinesCol01,
		err := createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumOfTxtLines :=
		txtLinesCol01.GetNumberOfTextLines()

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.GetNumberOfTextLines()\n"+
			"Expected Text Lines Don't Match Actual Text Lines\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"  Acutal Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return
	}

	var iTxtLineArray []ITextLineSpecification

	iTxtLineArray,
		err = txtLinesCol01.GetTextLineCollection(
		ePrefix.XCpy(
			"txtLinesCol01->iTxtLineArray"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumOfTxtLines = len(iTxtLineArray)

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PeekAtTextLineCollection()\n"+
			"Expected Text Lines Don't Match Actual Text Lines\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"  Acutal Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return
	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	err = txtLinesCol02.SetTextLineCollection(
		iTxtLineArray,
		ePrefix.XCpy(
			"txtLinesCol02<-iTxtLineArray"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !txtLinesCol02.Equal(&txtLinesCol01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.Equal(&txtLinesCol01)\n"+
			"Expected 'txtLinesCol02' EQUALS 'txtLinesCol01'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!!\n",
			ePrefix.String())

		return

	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	_,
		err = txtLinesCol03.PeekAtTextLine(
		0,
		ePrefix.XCpy(
			"txtLinesCol03 is EMPTY"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol03.PeekAtTextLine(0)\n"+
			"Expected an error return because 'txtLinesCol03'\n"+
			"is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol04 TextLineSpecLinesCollection

	_,
		txtLinesCol04,
		err = createTestTextLineSpecCollection02(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol04.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = txtLinesCol04.PeekAtTextLine(
		0,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol04.PeekAtTextLine()\n"+
			"Expected an error return because input\n"+
			"parameter 'errorPrefix is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol05 TextLineSpecLinesCollection

	_,
		txtLinesCol05,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines01 *TextLineSpecTimerLines

	_,
		timerLines01,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var actualTextTimerLine ITextLineSpecification

	actualTextTimerLine,
		err = txtLinesCol05.PeekAtTextLine(
		8,
		ePrefix.XCpy(
			"actualTextTimerLine<-txtLinesCol05[8]"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !actualTextTimerLine.EqualITextLine(
		timerLines01) {

		t.Errorf("%v - ERROR\n"+
			"Expected actualTextTimerLine would be equal to\n"+
			"timerLines01 from txtLinesCol05.PeekAtTextLine(8)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	txtLinesCol06 := TextLineSpecLinesCollection{}

	_,
		txtLinesCol06,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	leftMargin := 3
	rightMargin := 3
	textString := "The cow jumped over the moon!"

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCpy(
			"plainTextLine01"))

	var actualPlainTextLine ITextLineSpecification

	actualPlainTextLine,
		err = txtLinesCol06.PeekAtTextLine(
		0,
		ePrefix.XCpy(
			"actualPlainTextLine<-txtLinesCol06[0]"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !actualPlainTextLine.EqualITextLine(
		&plainTextLine01) {

		t.Errorf("%v - ERROR\n"+
			"Expected actualPlainTextLine would be equal to\n"+
			"plainTextLine01 from txtLinesCol06.PeekAtTextLine(0)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_ReplaceTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_ReplaceTextLine_000100()",
		"")

	_,
		txtLinesCol01,
		err := createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	leftMargin := 3
	rightMargin := 3
	textString := "There is a red moon in the sky!"

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCpy(
			"plainTextLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol01.ReplaceTextLine(
		&plainTextLine01,
		1,
		ePrefix.XCpy(
			"txtLinesCol01[1]<-plainTextLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var idx1ITxtLineSpec ITextLineSpecification

	idx1ITxtLineSpec,
		err = txtLinesCol01.PeekAtTextLine(
		1,
		ePrefix.XCpy(
			"idx1ITxtLineSpec<-txtLinesCol01[1]"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !idx1ITxtLineSpec.EqualITextLine(
		&plainTextLine01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.ReplaceTextLine()\n"+
			"Expected 'idx1ITxtLineSpec' EQUALS 'plainTextLine01'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!!\n",
			ePrefix.String())

		return
	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	err = txtLinesCol02.ReplaceTextLine(
		&plainTextLine01,
		1,
		ePrefix.XCpy(
			"txtLinesCol02 is EMPTY"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol02.ReplaceTextLine()\n"+
			"Expected an error return because 'txtLinesCol02'\n"+
			"is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol03 TextLineSpecLinesCollection

	_,
		txtLinesCol03,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol03.ReplaceTextLine(
		&plainTextLine01,
		-1,
		ePrefix.XCpy(
			"txtLinesCol03[-1]<-plainTextLine01"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol03.ReplaceTextLine()\n"+
			"Expected an error return because the target\n"+
			"index is less than zero.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol04 TextLineSpecLinesCollection

	_,
		txtLinesCol04,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol04.ReplaceTextLine(
		&plainTextLine01,
		999,
		ePrefix.XCpy(
			"txtLinesCol04[999]<-plainTextLine01"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol04.ReplaceTextLine()\n"+
			"Expected an error return because the target\n"+
			"index is greater than the last array index.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol05 TextLineSpecLinesCollection

	_,
		txtLinesCol05,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol05.ReplaceTextLine(
		&plainTextLine01,
		1,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol05.ReplaceTextLine()\n"+
			"Expected an error return because input\n"+
			"parameter 'errorPrefix is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol06 TextLineSpecLinesCollection

	_,
		txtLinesCol06,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol06.ReplaceTextLine(
		nil,
		1,
		ePrefix.XCpy(
			"textLine is 'nil'"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol06.ReplaceTextLine()\n"+
			"Expected an error return because input\n"+
			"parameter 'textLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

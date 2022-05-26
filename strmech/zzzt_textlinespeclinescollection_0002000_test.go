package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecLinesCollection_NewTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_NewTextLine_000100()",
		"")

	var err error
	var stdLine01 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var txtLinesCol01 TextLineSpecLinesCollection

	txtLinesCol01,
		err = TextLineSpecLinesCollection{}.
		NewTextLine(
			&stdLine01,
			ePrefix.XCpy(
				"txtLinesCol01<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol01.IsValidInstanceError(
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
			"iTxtLineSpec<-txtLinesCol01.textLines[0]"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		&stdLine01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PeekAtFirstTextLine()\n"+
			"Expected 'iTxtLineSpec' object would be\n"+
			"Equal to 'stdLine01' object."+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
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

	stdLine02.Empty()

	_,
		err = TextLineSpecLinesCollection{}.
		NewTextLine(
			&stdLine02,
			ePrefix.XCpy(
				"stdLine02 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from\n"+
			"TextLineSpecLinesCollection{}."+
			"NewTextLine()\n"+
			"because 'stdLine02' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err = TextLineSpecLinesCollection{}.
		NewTextLine(
			nil,
			ePrefix.XCpy(
				"textLine is nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from\n"+
			"TextLineSpecLinesCollection{}."+
			"NewTextLine()\n"+
			"because 'textLine' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

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
		err = TextLineSpecLinesCollection{}.
		NewTextLine(
			&stdLine03,
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from\n"+
			"TextLineSpecLinesCollection{}."+
			"NewTextLine()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

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

func TestTextLineSpecLinesCollection_PopFirstTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_PopFirstTextLine_000100()",
		"")

	// Index 0

	leftMargin := 3
	rightMargin := 3
	textString := "The cow jumped over the moon!"

	var plainTextLine01 TextLineSpecPlainText
	var err error

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCpy(
			"plainTextLine01"))

	var txtLinesCol01 TextLineSpecLinesCollection
	expectedCollectionLen := 0

	expectedCollectionLen,
		txtLinesCol01,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var iTxtLineSpec ITextLineSpecification
	actualCollectionLen := 0

	iTxtLineSpec,
		actualCollectionLen,
		err = txtLinesCol01.PopFirstTextLine(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		&plainTextLine01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PopFirstTextLine()\n"+
			"Expected ITextLineSpecifications object would be\n"+
			"Equal to 'plainTextLine01' object."+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	expectedCollectionLen--

	if expectedCollectionLen != actualCollectionLen {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PopFirstTextLine()\n"+
			"Expected Collection Length is NOT EQUAL TO\n"+
			"Actual Collection Length!!\n"+
			"Expected Collection Length = '%v'\n"+
			"  Actual Collection Length = '%v'\n",
			ePrefix.String(),
			expectedCollectionLen,
			actualCollectionLen)

		return

	}

	targetIndex := 0

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol01[%v]",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if iTxtLineSpec.EqualITextLine(
		&plainTextLine01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PopFirstTextLine()\n"+
			"Expected ITextLineSpecifications object would NOT be\n"+
			"Equal to 'plainTextLine01' object."+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.String())

		return
	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	expectedCollectionLen,
		txtLinesCol02,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		_,
		err = txtLinesCol02.PopFirstTextLine(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"PopFirstTextLine()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	_,
		_,
		err = txtLinesCol03.PopFirstTextLine(
		ePrefix.XCpy(
			"txtLinesCol03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"PopFirstTextLine()\n"+
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
		_,
		err = txtLinesCol04.PopFirstTextLine(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"PopFirstTextLine()\n"+
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
		_,
		err = txtLinesCol05.PopFirstTextLine(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol05."+
			"PopFirstTextLine()\n"+
			"because 'txtLinesCol05.textLines[0]' is invalid.\n"+
			"txtLinesCol05.textLines[0].Empty() was previously called.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_PopLastTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_PopLastTextLine_000100()",
		"")

	// Index 8

	var timerLines01 *TextLineSpecTimerLines
	var err error

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

	var txtLinesCol01 TextLineSpecLinesCollection
	expectedCollectionLen := 0

	expectedCollectionLen,
		txtLinesCol01,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lastIndex := expectedCollectionLen - 1

	var iTxtLineSpec ITextLineSpecification
	actualCollectionLen := 0

	iTxtLineSpec,
		actualCollectionLen,
		err = txtLinesCol01.PopLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		timerLines01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PopLastTextLine()\n"+
			"Expected ITextLineSpecifications object would be\n"+
			"Equal to Actual ITextLineSpecifications object."+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	expectedCollectionLen--

	if expectedCollectionLen != actualCollectionLen {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PopLastTextLine()\n"+
			"Expected Collection Length is NOT EQUAL TO\n"+
			"Actual Collection Length!!\n"+
			"Expected Collection Length = '%v'\n"+
			"  Actual Collection Length = '%v'\n",
			ePrefix.String(),
			expectedCollectionLen,
			actualCollectionLen)

		return

	}

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtTextLine(
		lastIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol01[%v]",
				lastIndex)))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol01.PeekAtTextLine(%v)\n"+
			"Expected an error return because txtLinesCol03[%v]\n"+
			"should have been deleted.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			lastIndex,
			lastIndex)

		return
	}

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtLastTextLine(
		ePrefix.XCpy("txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if iTxtLineSpec.EqualITextLine(
		timerLines01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PeekAtLastTextLine()\n"+
			"Expected ITextLineSpecifications object would NOT be\n"+
			"Equal to Actual ITextLineSpecifications object."+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.String())

		return

	}

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	_,
		txtLinesCol02,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		_,
		err = txtLinesCol02.PopLastTextLine(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"PopLastTextLine()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	_,
		_,
		err = txtLinesCol03.PopLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"PopLastTextLine()\n"+
			"because 'txtLinesCol03' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol04 TextLineSpecLinesCollection

	expectedCollectionLen,
		txtLinesCol04,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lastIndex = expectedCollectionLen - 1

	txtLinesCol04.textLines[lastIndex] = nil

	_,
		_,
		err = txtLinesCol04.PopLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"PopLastTextLine()\n"+
			"because 'txtLinesCol04.textLines[%v]' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			lastIndex)

		return
	}

	var txtLinesCol05 TextLineSpecLinesCollection

	lastIndex,
		txtLinesCol05,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	lastIndex--

	txtLinesCol05.textLines[lastIndex].Empty()

	_,
		_,
		err = txtLinesCol05.PopLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol05."+
			"PopLastTextLine()\n"+
			"because 'txtLinesCol05.textLines[0]' is invalid.\n"+
			"txtLinesCol05.textLines[%v].Empty() was previously called.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			lastIndex)

		return
	}

	return
}

func TestTextLineSpecLinesCollection_PopTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_PopTextLine_000100()",
		"")

	expectedNumOfTxtLines,
		txtLinesCol01,
		err := createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var stdLine05 TextLineSpecStandardLine

	stdLine05,
		err = createTestTextLineSpecStandardLine05(
		ePrefix.XCpy(
			"stdLine05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumOfTxtLines := 0

	var iTxtLineSpec ITextLineSpecification

	targetIndex := 5

	iTxtLineSpec,
		actualNumOfTxtLines,
		err =
		txtLinesCol01.PopTextLine(
			targetIndex,
			ePrefix.XCpy(
				fmt.Sprintf("txtLinesCol01[%v]",
					targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumOfTxtLines--

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PopTextLine()\n"+
			"Expected Text Lines Don't Match Actual Text Lines\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"  Acutal Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return
	}

	if !iTxtLineSpec.EqualITextLine(&stdLine05) {

		t.Errorf("%v - ERROR\n"+
			"Expected iTxtLineSpec would be equal to\n"+
			"stdLine05 from txtLinesCol01.PopTextLine(%v)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String(),
			targetIndex)

		return
	}

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol01[%v]",
				targetIndex)))

	if iTxtLineSpec.EqualITextLine(&stdLine05) {

		t.Errorf("%v - ERROR\n"+
			"Expected txtLinesCol01.PopTextLine(%v) would\n"+
			"delete 'stdLine05' from the array.\n"+
			"HOWEVER, IT IS STILL PRESENT AT THE SAME INDEX!\n",
			ePrefix.String(),
			targetIndex)

		return
	}

	var txtLinesCol02 TextLineSpecLinesCollection

	expectedNumOfTxtLines,
		txtLinesCol02,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex = 0

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

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	iTxtLineSpec,
		actualNumOfTxtLines,
		err = txtLinesCol02.PopTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol02[%v]",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(&plainTextLine01) {

		t.Errorf("%v - ERROR\n"+
			"Expected iTxtLineSpec would be equal to\n"+
			"plainTextLine01 from txtLinesCol02.PopTextLine(%v)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String(),
			targetIndex)

		return
	}

	expectedNumOfTxtLines--

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.PopTextLine()\n"+
			"Expected Text Lines Don't Match Actual Text Lines\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"  Acutal Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return
	}

	iTxtLineSpec,
		err = txtLinesCol02.PeekAtTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol02[%v]",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if iTxtLineSpec.EqualITextLine(&plainTextLine01) {

		t.Errorf("%v - ERROR\n"+
			"Expected txtLinesCol02.PopTextLine(%v) would\n"+
			"delete 'plainTextLine01' from the array.\n"+
			"HOWEVER, IT IS STILL PRESENT AT THE SAME INDEX!\n",
			ePrefix.String(),
			targetIndex)

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	expectedNumOfTxtLines,
		txtLinesCol03,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol03"))

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

	lastIndex := expectedNumOfTxtLines - 1

	iTxtLineSpec,
		actualNumOfTxtLines,
		err = txtLinesCol03.PopTextLine(
		lastIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol03[%v]",
				lastIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	expectedNumOfTxtLines--

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol03.PopTextLine()\n"+
			"Expected Text Lines Don't Match Actual Text Lines\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"  Acutal Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return
	}

	if !iTxtLineSpec.EqualITextLine(timerLines01) {

		t.Errorf("%v - ERROR\n"+
			"Expected iTxtLineSpec would be equal to\n"+
			"timerLines01 from txtLinesCol01.PopTextLine(%v)\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String(),
			lastIndex)

		return
	}

	_,
		err = txtLinesCol03.PeekAtTextLine(
		lastIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol03[%v]",
				targetIndex)))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol03.PeekAtTextLine(%v)\n"+
			"Expected an error return because txtLinesCol03[%v]\n"+
			"should have been deleted.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			lastIndex,
			lastIndex)

		return
	}

	iTxtLineSpec,
		err = txtLinesCol03.PeekAtLastTextLine(
		ePrefix.XCpy(
			"txtLinesCol03"))

	if iTxtLineSpec.EqualITextLine(timerLines01) {

		t.Errorf("%v - ERROR\n"+
			"Expected iTxtLineSpec would NOT be equal to\n"+
			"timerLines01 from txtLinesCol01.PopAtLastTextLine()\n"+
			"HOWEVER, THEY ARE EQUAL!\n",
			ePrefix.String())

		return
	}

	txtLinesCol04 := TextLineSpecLinesCollection{}
	targetIndex = 0

	_,
		_,
		err = txtLinesCol04.PopTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol04[%v]",
				targetIndex)))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"PopTextLine()\n"+
			"because input parameter 'txtLinesCol04' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol05 := TextLineSpecLinesCollection{}

	targetIndex = 0

	expectedNumOfTxtLines,
		txtLinesCol05,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		_,
		err = txtLinesCol05.PopTextLine(
		targetIndex,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol05."+
			"PopTextLine()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol06 := TextLineSpecLinesCollection{}

	expectedNumOfTxtLines,
		txtLinesCol06,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol06"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex = -1

	_,
		_,
		err = txtLinesCol06.PopTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol06[%v]",
				targetIndex)))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol06."+
			"PopTextLine() because\n"+
			"input parameter zeroBasedIndex = '%v' and is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!",
			ePrefix.String(),
			targetIndex)

		return
	}

	txtLinesCol07 := TextLineSpecLinesCollection{}

	expectedNumOfTxtLines,
		txtLinesCol07,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol07"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex = 999

	_,
		_,
		err = txtLinesCol07.PopTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol07[%v]",
				targetIndex)))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol07."+
			"PopTextLine() because\n"+
			"input parameter zeroBasedIndex = '%v' and is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!",
			ePrefix.String(),
			targetIndex)

		return
	}

	txtLinesCol08 := TextLineSpecLinesCollection{}

	expectedNumOfTxtLines,
		txtLinesCol08,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol08"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex = 5

	txtLinesCol08.textLines[targetIndex] = nil

	_,
		_,
		err = txtLinesCol08.PopTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol08[%v]",
				targetIndex)))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol08."+
			"PopTextLine() because\n"+
			"txtLinesCol08.textLines[%v] = nil and is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!",
			ePrefix.String(),
			targetIndex)

		return
	}

	txtLinesCol09 := TextLineSpecLinesCollection{}

	expectedNumOfTxtLines,
		txtLinesCol09,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol09"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex = 5

	txtLinesCol09.textLines[targetIndex].Empty()

	_,
		_,
		err = txtLinesCol09.PopTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf(
				"txtLinesCol09[%v]",
				targetIndex)))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol05."+
			"PopFirstTextLine()\n"+
			"because 'txtLinesCol09.textLines[%v]' is invalid.\n"+
			"txtLinesCol09.textLines[%v].Empty() was previously called.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			targetIndex,
			targetIndex)

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

	targetIndex := 1

	err = txtLinesCol01.ReplaceTextLine(
		&plainTextLine01,
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf("txtLinesCol01[%v]<-plainTextLine01",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var iTxtLineSpec ITextLineSpecification

	iTxtLineSpec,
		err = txtLinesCol01.PeekAtTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf("iTxtLineSpec<-txtLinesCol01[%v]",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		&plainTextLine01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.PeekAtTextLine(%v)\n"+
			"Expected 'iTxtLineSpec' EQUALS 'plainTextLine01'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!!\n",
			ePrefix.String(),
			targetIndex)

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

	var txtLinesCol07 TextLineSpecLinesCollection

	_,
		txtLinesCol07,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol07"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex = 0

	err = txtLinesCol07.ReplaceTextLine(
		&plainTextLine01,
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf("txtLinesCol07[%v] <- plainTextLine01",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	iTxtLineSpec,
		err = txtLinesCol07.PeekAtTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf("iTxtLineSpec<-txtLinesCol07[%v]",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		&plainTextLine01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol07.PeekAtTextLine(%v)\n"+
			"Expected 'iTxtLineSpec' EQUALS 'plainTextLine01'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!!\n",
			ePrefix.String(),
			targetIndex)

		return
	}

	var txtLinesCol08 TextLineSpecLinesCollection
	numOfTxtLines := 0

	numOfTxtLines,
		txtLinesCol08,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol08"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex = numOfTxtLines - 1

	err = txtLinesCol08.ReplaceTextLine(
		&plainTextLine01,
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf("txtLinesCol08[%v] <- plainTextLine01",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	iTxtLineSpec,
		err = txtLinesCol08.PeekAtTextLine(
		targetIndex,
		ePrefix.XCpy(
			fmt.Sprintf("iTxtLineSpec<-txtLinesCol08[%v]",
				targetIndex)))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !iTxtLineSpec.EqualITextLine(
		&plainTextLine01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol08.PeekAtTextLine(%v)\n"+
			"Expected 'iTxtLineSpec' EQUALS 'plainTextLine01'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!!\n",
			ePrefix.String(),
			targetIndex)

		return
	}

	plainTextLine02 := TextLineSpecPlainText{}

	var txtLinesCol09 TextLineSpecLinesCollection

	_,
		txtLinesCol09,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol09"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol09.ReplaceTextLine(
		&plainTextLine02,
		0,
		ePrefix.XCpy(
			"plainTextLine02 is empty"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol09.ReplaceTextLine()\n"+
			"Expected an error return because input\n"+
			"parameter 'textLine' (plainTextLine02) is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_SetTextLineCollection_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_SetTextLineCollection_000100()",
		"")

	_,
		txtLinesCol01,
		err := createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var txtLineCol []ITextLineSpecification

	txtLineCol,
		err = txtLinesCol01.GetTextLineCollection(
		ePrefix.XCpy(
			"txtLineCol<-txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newNumOfTxtLines := len(txtLineCol)

	var txtLinesCol02 TextLineSpecLinesCollection

	oldNumOfTxtLines := 0

	oldNumOfTxtLines,
		txtLinesCol02,
		err = createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol02.SetTextLineCollection(
		txtLineCol,
		ePrefix.XCpy(
			"txtLinesCol02<-txtLineCol"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !txtLinesCol02.Equal(&txtLinesCol01) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.SetTextLineCollection()\n"+
			"Expected txtLinesCol01 == txtLinesCol02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n",
			ePrefix.String())

		return
	}

	currentNumOfTxtLines :=
		txtLinesCol02.GetNumberOfTextLines()

	if currentNumOfTxtLines != newNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.SetTextLineCollection()\n"+
			"Expected Text Lines Don't Match Actual Text Lines\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"  Acutal Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			newNumOfTxtLines,
			currentNumOfTxtLines)

		return
	}

	if currentNumOfTxtLines == oldNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.SetTextLineCollection()\n"+
			"Expected Text Lines Don't Match Actual Text Lines\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"  Acutal Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			newNumOfTxtLines,
			currentNumOfTxtLines)

		return
	}

	var txtLinesCol03 TextLineSpecLinesCollection

	_,
		txtLinesCol03,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol03.SetTextLineCollection(
		nil,
		ePrefix.XCpy(
			"txtLinesCol03<-nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"SetTextLineCollection()\n"+
			"because input parameter 'newTextLineCol' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol04 TextLineSpecLinesCollection

	newNumOfTxtLines,
		txtLinesCol04,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLineCol,
		err = txtLinesCol04.GetTextLineCollection(
		ePrefix.XCpy(
			"txtLinesCol04"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex := newNumOfTxtLines - 1

	txtLineCol[targetIndex].Empty()

	err = txtLinesCol04.SetTextLineCollection(
		txtLineCol,
		ePrefix.XCpy(
			"txtLinesCol04<-nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"SetTextLineCollection()\n"+
			"because input parameter 'txtLineCol[%v]' is invalid.\n"+
			"txtLineCol[%v].Empty() was previously called.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			targetIndex,
			targetIndex)

		return
	}

	var txtLinesCol05 TextLineSpecLinesCollection

	newNumOfTxtLines,
		txtLinesCol05,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLineCol,
		err = txtLinesCol05.GetTextLineCollection(
		ePrefix.XCpy(
			"txtLinesCol05"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	targetIndex = newNumOfTxtLines - 1

	txtLineCol[targetIndex] = nil

	err = txtLinesCol05.SetTextLineCollection(
		txtLineCol,
		ePrefix.XCpy(
			"txtLinesCol05<-nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol05."+
			"SetTextLineCollection()\n"+
			"because input parameter 'txtLineCol[%v]' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String(),
			targetIndex)

		return
	}

	_,
		txtLinesCol05,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol05 #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLineCol,
		err = txtLinesCol05.GetTextLineCollection(
		ePrefix.XCpy(
			"txtLinesCol05 #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol06 := TextLineSpecLinesCollection{}

	err = txtLinesCol06.SetTextLineCollection(
		txtLineCol,
		ePrefix.XCpy(
			"txtLinesCol06<-txtLineCol"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var txtLinesCol07 TextLineSpecLinesCollection

	newNumOfTxtLines,
		txtLinesCol07,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"txtLinesCol07"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLineCol,
		err = txtLinesCol07.GetTextLineCollection(
		ePrefix.XCpy(
			"txtLinesCol07"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol07.SetTextLineCollection(
		txtLineCol,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol07."+
			"SetTextLineCollection()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

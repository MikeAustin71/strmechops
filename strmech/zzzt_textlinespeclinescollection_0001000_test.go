package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func TestTextLineSpecLinesCollection_AddTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_AddTextLine_000100()",
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

	if !txtLinesCol01.IsValidInstance() {
		t.Errorf("%v - ERROR\n"+
			"The 'txtLinesCol1' is invalid!\n"+
			"createTestTextLineSpecCollection01()\n"+
			"returned an invalid collection!\n",
			ePrefix.String())

		return

	}

	actualNumOfTxtLines :=
		txtLinesCol01.GetNumberOfTextLines()

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.GetNumberOfTextLines()\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"Instead  Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return

	}

	var stdLine01 TextLineSpecStandardLine
	txtLinesCol02 := TextLineSpecLinesCollection{}

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol02.AddTextLine(
		&stdLine01,
		ePrefix.XCpy(
			"txtLinesCol02<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	actualNumOfTxtLines = txtLinesCol02.GetNumberOfTextLines()
	expectedNumOfTxtLines = 1

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.GetNumberOfTextLines()\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"Instead  Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return

	}

	var stdLine02 TextLineSpecStandardLine
	txtLinesCol03 := TextLineSpecLinesCollection{}

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol03.AddTextLine(
		&stdLine02,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"AddTextLine()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_CopyIn_000100()",
		"")

	expectedNumOfTxtLines,
		txtLinesCol01,
		err := createTestTextLineSpecCollection02(
		ePrefix.XCpy(
			"txtLinesCol01<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	err =
		txtLinesCol02.CopyIn(
			&txtLinesCol01,
			ePrefix.XCpy(
				"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !txtLinesCol02.IsValidInstance() {
		t.Errorf("%v - ERROR\n"+
			"The 'txtLinesCol02' is invalid!\n"+
			"txtLinesCol02.CopyIn(&txtLinesCol01)\n"+
			"returned an invalid collection!\n",
			ePrefix.String())

		return
	}

	actualNumOfTxtLines := txtLinesCol02.GetNumberOfTextLines()

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.GetNumberOfTextLines()\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"Instead  Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return

	}

	if !txtLinesCol01.Equal(&txtLinesCol02) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.Equal(&txtLinesCol02)\n"+
			"Expected 'txtLinesCol01' == 'txtLinesCol02'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	err =
		txtLinesCol03.CopyIn(
			nil,
			ePrefix.XCpy(
				"txtLinesCol03"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"CopyIn()\n"+
			"because 'incomingTxtLinesCol' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol04 := TextLineSpecLinesCollection{}

	err =
		txtLinesCol04.CopyIn(
			&txtLinesCol01,
			StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_CopyOut_000100()",
		"")

	expectedNumOfTxtLines,
		txtLinesCol01,
		err := createTestTextLineSpecCollection02(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !txtLinesCol01.IsValidInstance() {
		t.Errorf("%v - ERROR\n"+
			"The 'txtLinesCol01' is invalid!\n"+
			"createTestTextLineSpecCollection02()\n"+
			"returned an invalid collection!\n",
			ePrefix.String())

		return
	}

	actualNumOfTxtLines := txtLinesCol01.GetNumberOfTextLines()

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.GetNumberOfTextLines()\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"Instead  Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return

	}

	var txtLinesCol02 TextLineSpecLinesCollection

	txtLinesCol02,
		err = txtLinesCol01.CopyOut(
		ePrefix.XCpy(
			"txtLinesCol02<-txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !txtLinesCol02.IsValidInstance() {
		t.Errorf("%v - ERROR\n"+
			"The 'txtLinesCol02' is invalid!\n"+
			"txtLinesCol02.CopyOut() returned\n"+
			"an invalid collection!\n",
			ePrefix.String())

		return
	}

	actualNumOfTxtLines = txtLinesCol02.GetNumberOfTextLines()

	if expectedNumOfTxtLines != actualNumOfTxtLines {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.GetNumberOfTextLines()\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"Instead  Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return

	}

	if !txtLinesCol01.Equal(&txtLinesCol02) {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol01.Equal(&txtLinesCol02)\n"+
			"Expected 'txtLinesCol01' == 'txtLinesCol02'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	_,
		err = txtLinesCol03.CopyOut(
		ePrefix.XCpy(
			"txtLinesCol03 is Empty!"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"CopyOut()\n"+
			"because 'txtLinesCol03' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol04 TextLineSpecLinesCollection

	txtLinesCol04,
		err = txtLinesCol01.CopyOut(
		ePrefix.XCpy(
			"txtLinesCol04<-txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		err = txtLinesCol04.CopyOut(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_DeleteTextLineMember_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_DeleteTextLineMember_000100()",
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

	expectedNumOfTxtLines--

	err = txtLinesCol01.DeleteTextLineMember(
		2,
		ePrefix.XCpy(
			"txtLinesCol01 index=2"))

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

	var iTxtLine ITextLineSpecification

	iTxtLine,
		err = txtLinesCol01.GetTextLine(
		2,
		ePrefix.XCpy(
			"txtLinesCol01 index=2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	typeName := iTxtLine.TextLineSpecName()

	if typeName != "BlankLines" {

		t.Errorf("\n%v\n"+
			"Error: iTxtLine.TextLineSpecName()\n"+
			"Expected Type Name would equal 'BlankLines'\n"+
			"after deleting index #2. However, the Expected\n"+
			"Type Name DOES NOT EQUAL the Actual Type Name\n"+
			"Expected Type Name = 'Blank Lines' \n"+
			"  Acutal Type Name = '%v'\n",
			ePrefix.String(),
			typeName)

		return
	}

	var txtLinesCol02 TextLineSpecLinesCollection

	_,
		txtLinesCol02,
		err = createTestTextLineSpecCollection02(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol02.DeleteTextLineMember(
		-1,
		ePrefix.XCpy(
			"txtLinesCol02 index=-1"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"DeleteTextLineMember()\n"+
			"because input parameter 'zeroBaseIndex' has\n"+
			"a value of minus one (-1).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtLinesCol02.DeleteTextLineMember(
		99,
		ePrefix.XCpy(
			"txtLinesCol02 index=99"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"DeleteTextLineMember()\n"+
			"because input parameter 'zeroBaseIndex'\n"+
			"has a value of 99.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtLinesCol02.DeleteTextLineMember(
		2,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"DeleteTextLineMember()\n"+
			"because input parameter 'errorPrefix is invalid.'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_Empty_000100()",
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

	err = txtLinesCol01.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol01.Empty()

	actualNumOfTextLines := txtLinesCol01.GetNumberOfTextLines()

	if actualNumOfTextLines != 0 {

		t.Errorf("%v - ERROR\n"+
			"After calling txtLinesCol01.Empty()\n"+
			"it was expected that all Text Lines would be deleted.\n"+
			"HOWEVER, TEXT LINES STILL REMAIN IN THE COLLECTION!\n"+
			"Actual Number of Text Lines in Collection = '%v'\n",
			ePrefix.String(),
			actualNumOfTextLines)

		return

	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	txtLinesCol02.Empty()

	err = txtLinesCol02.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol02 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"IsValidInstanceError()\n"+
			"because 'txtLinesCol02.Empty()' was called\n"+
			"immediately prior.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_EmptyTextLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_EmptyTextLines_000100()",
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

	err = txtLinesCol01.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol01.EmptyTextLines()

	actualNumOfTextLines := txtLinesCol01.GetNumberOfTextLines()

	if actualNumOfTextLines != 0 {

		t.Errorf("%v - ERROR\n"+
			"After calling txtLinesCol01.EmptyTextLines()\n"+
			"it was expected that all Text Lines would be deleted.\n"+
			"HOWEVER, TEXT LINES STILL REMAIN IN THE COLLECTION!\n"+
			"Actual Number of Text Lines in Collection = '%v'\n",
			ePrefix.String(),
			actualNumOfTextLines)

		return

	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	txtLinesCol02.EmptyTextLines()

	err = txtLinesCol02.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol02 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"IsValidInstanceError()\n"+
			"because 'txtLinesCol02.EmptyTextLines()' was called\n"+
			"immediately prior.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_Equal_000100()",
		"")

	_,
		txtLinesCol01,
		err := createTestTextLineSpecCollection02(
		ePrefix.XCpy(
			"txtLinesCol01<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	areEqual := txtLinesCol02.Equal(
		&txtLinesCol01)

	if areEqual == true {

		t.Errorf("%v - ERROR\n"+
			"Expected a return of 'false' from txtLinesCol02."+
			"Equal()\n"+
			"because 'txtLinesCol02' is empty.\n"+
			"HOWEVER, A VALUE OF 'true' WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtLinesCol02.CopyIn(
			&txtLinesCol01,
			ePrefix.XCpy(
				"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol02.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if !txtLinesCol01.Equal(&txtLinesCol02) {

		t.Errorf("\n%v\n"+
			"Test #2 \n"+
			"Error: txtLinesCol01.Equal(&txtLinesCol02)\n"+
			"Expected 'txtLinesCol01' == 'txtLinesCol02'.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	err = txtLinesCol02.DeleteTextLineMember(
		1,
		ePrefix.XCpy(
			"txtLinesCol02 Delete Index 1"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if txtLinesCol01.Equal(&txtLinesCol02) {

		t.Errorf("\n%v\n"+
			"Test #3 \n"+
			"Error: txtLinesCol01.Equal(&txtLinesCol02)\n"+
			"Expected 'txtLinesCol01' would NOT be equal to 'txtLinesCol02'"+
			"because txtLinesCol02 index #1 has been deleted.\n"+
			"HOWEVER, THEY ARE EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_GetNumberOfTextLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_GetNumberOfTextLines_000100()",
		"")

	expectedNumOfTxtLines,
		txtLinesCol01,
		err := createTestTextLineSpecCollection02(
		ePrefix.XCpy(
			"txtLinesCol01<-"))

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
			"Expected Number Of Text Lines = '%v' \n"+
			"Instead  Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return

	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	actualNumOfTxtLines =
		txtLinesCol02.GetNumberOfTextLines()

	if actualNumOfTxtLines != 0 {

		t.Errorf("\n%v\n"+
			"Error: txtLinesCol02.GetNumberOfTextLines()\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"Instead  Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			0,
			actualNumOfTxtLines)

		return

	}

	return
}

func TestTextLineSpecLinesCollection_GetTextLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_GetTextLine_000100()",
		"")

	// Index 1
	stdLine02,
		err := createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine01"))

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
		err = txtLinesCol01.GetTextLine(
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
			"Error: txtLinesCol01.GetTextLine(1)\n"+
			"Expected ITextLineSpecifications object would be\n"+
			"Equal to Actual ITextLineSpecifications object."+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return

	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	_,
		err = txtLinesCol02.GetTextLine(
		1,
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"DeleteTextLineMember()\n"+
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
		err = txtLinesCol03.GetTextLine(
		1,
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"DeleteTextLineMember()\n"+
			"because input parameter 'zeroBaseIndex' has\n"+
			"a value of minus one (-1).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTestTextLineSpecLinesCollection_GetTextLineCollection_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTestTextLineSpecLinesCollection_GetTextLineCollection_000100()",
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
			"Error: txtLinesCol01.GetTextLineCollection()\n"+
			"Expected Text Lines Don't Match Actual Text Lines\n"+
			"Expected Number Of Text Lines = '%v' \n"+
			"  Acutal Number Of Text Lines = '%v'\n",
			ePrefix.String(),
			expectedNumOfTxtLines,
			actualNumOfTxtLines)

		return
	}

	txtLinesCol02 := TextLineSpecLinesCollection{}

	txtLinesCol02.textLines = iTxtLineArray

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
		err = txtLinesCol03.GetTextLineCollection(
		ePrefix.XCpy(
			"txtLinesCol03 is EMPTY"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol03.GetTextLineCollection()\n"+
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
			"txtLinesCol03"))

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
		err = txtLinesCol04.GetTextLineCollection(
		StrMech{})

	if err == nil {

		t.Errorf("\n%v\n"+
			"ERROR: txtLinesCol04.GetTextLineCollection()\n"+
			"Expected an error return because input\n"+
			"parameter 'errorPrefix is invalid.'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_IsValidInstance_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_IsValidInstance_000100()",
		"")

	txtLinesCol01 := TextLineSpecLinesCollection{}

	if txtLinesCol01.IsValidInstance() {

		t.Errorf("%v - ERROR\n"+
			"Expected a value of 'false' from txtLinesCol01."+
			"IsValidInstance()\n"+
			"because 'txtLinesCol01' is empty.\n"+
			"HOWEVER, A VALUE OF 'true' WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecLinesCollection_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_IsValidInstanceError_000100()",
		"")

	_,
		txtLinesCol01,
		err := createTestTextLineSpecCollection02(
		ePrefix.XCpy(
			"txtLinesCol01"))

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

	var txtLinesCol02 TextLineSpecLinesCollection

	txtLinesCol02,
		err = txtLinesCol01.CopyOut(
		ePrefix.XCpy(
			"txtLinesCol02<-txtLinesCol01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = txtLinesCol02.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtLinesCol02.textLines[2].Empty()

	err = txtLinesCol02.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol02-invalid array element"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol02."+
			"IsValidInstanceError()\n"+
			"because 'txtLinesCol02.textLines[2]' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLinesCol03 := TextLineSpecLinesCollection{}

	err = txtLinesCol03.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLinesCol03 is empty"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol03."+
			"IsValidInstanceError()\n"+
			"because 'txtLinesCol03' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	var txtLinesCol04 TextLineSpecLinesCollection

	txtLinesCol04,
		err = txtLinesCol01.CopyOut(
		ePrefix.XCpy(
			"txtLinesCol04<-txtLinesCol01"))

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

	err = txtLinesCol04.IsValidInstanceError(
		StrMech{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtLinesCol04."+
			"IsValidInstanceError()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

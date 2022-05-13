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

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

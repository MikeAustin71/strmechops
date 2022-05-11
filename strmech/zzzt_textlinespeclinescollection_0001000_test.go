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
			"stdLinesCol01"))

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

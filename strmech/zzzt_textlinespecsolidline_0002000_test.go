package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func TestTextLineSpecSolidLine_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_CopyIn_000100()",
		"")

	solidLineChars := "-"
	solidLineCharsRepeatCount := 35

	expectedSolidLineStr :=
		strings.Repeat(
			solidLineChars,
			solidLineCharsRepeatCount) +
			"\n"

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrDefaultSolidLine(
		solidLineChars,
		solidLineCharsRepeatCount,
		ePrefix.XCtx(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	err =
		txtSolidLineMolecule.copyIn(
			&txtSolidLine02,
			txtSolidLine01,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCtx(
			"txtSolidLine01->txtSolidLine01FmtText"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCtx(
			"txtSolidLine02->txtSolidLine02FmtText"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine02FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	printableExpectedStr =
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Formatted Text String DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if !txtSolidLine02.Equal(txtSolidLine01) {
		t.Errorf("%v - ERROR\n"+
			"Test #3"+
			"Expected Text Solid Line Object #1 would be\n"+
			"Equal to Text Solid Line Object #2.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtSolidLine01.solidLineChars = nil

	txtSolidLine03 := TextLineSpecSolidLine{}

	err =
		txtSolidLineMolecule.copyIn(
			&txtSolidLine03,
			txtSolidLine01,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'txtSolidLine01.solidLineChars' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtSolidLineMolecule.copyIn(
			nil,
			&txtSolidLine02,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'targetTxtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		txtSolidLineMolecule.copyIn(
			&txtSolidLine03,
			nil,
			ePrefix.XCtx(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'incomingTxtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

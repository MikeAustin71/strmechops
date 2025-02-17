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
		ePrefix.XCpy(
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
			ePrefix.XCpy(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine01FmtText"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCpy(
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
			ePrefix.String(),
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
			ePrefix.String(),
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
			ePrefix.String())

		return
	}

	txtSolidLine01.solidLineChars = nil

	txtSolidLine03 := TextLineSpecSolidLine{}

	err =
		txtSolidLineMolecule.copyIn(
			&txtSolidLine03,
			txtSolidLine01,
			ePrefix.XCpy(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'txtSolidLine01.solidLineChars' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtSolidLineMolecule.copyIn(
			nil,
			&txtSolidLine02,
			ePrefix.XCpy(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'targetTxtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtSolidLineMolecule.copyIn(
			&txtSolidLine03,
			nil,
			ePrefix.XCpy(
				"txtSolidLine02<-txtSolidLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'incomingTxtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_copyOut_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.NewPtrFullSolidLineConfig(
		leftMargin,
		rightMargin,
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	txtSolidLine02,
		err =
		txtSolidLineMolecule.copyOut(
			txtSolidLine01,
			ePrefix.XCpy(
				"txtSolidLine02<-txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText, txtSolidLine02FmtText string

	txtSolidLine01FmtText,
		err = txtSolidLine01.GetFormattedText(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine01FmtText"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02FmtText,
		err = txtSolidLine02.GetFormattedText(
		ePrefix.XCpy(
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
			ePrefix.String(),
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
			ePrefix.String(),
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
			ePrefix.String())

		return
	}

	txtSolidLine01.solidLineChars = nil

	_,
		err =
		txtSolidLineMolecule.copyOut(
			txtSolidLine01,
			ePrefix.XCpy(
				"txtSolidLine01 invalid test #1"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyOut()\n"+
			"because 'txtSolidLine01.solidLineChars' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err =
		txtSolidLineMolecule.copyOut(
			nil,
			ePrefix.XCpy(
				"txtSolidLine == nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule."+
			"copyIn()\n"+
			"because 'targetTxtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_empty_000100(t *testing.T) {

	txtSolidLineAtom := textLineSpecSolidLineAtom{}

	txtSolidLineAtom.empty(
		nil)

}

func TestTextLineSpecSolidLine_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_equal_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n-\n"

	txtSolidLine01 := TextLineSpecSolidLine{}

	err :=
		txtSolidLine01.SetFullSolidLineConfig(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtSolidLine01.IsValidInstanceError(
		ePrefix.XCpy("txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLine02,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLineAtom := textLineSpecSolidLineAtom{}

	areEqual :=
		txtSolidLineAtom.equal(
			&txtSolidLine01,
			&txtSolidLine02)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected txtSolidLine01==txtSolidLine02.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	areEqual =
		txtSolidLineAtom.equal(
			nil,
			&txtSolidLine02)

	if areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected areEqual WOULD BE 'false'\n"+
			"because 'txtSolidLineOne' has a 'nil' value.\n"+
			"HOWEVER, 'areEqual' is 'true'!!\n",
			ePrefix.String())

		return
	}

	areEqual =
		txtSolidLineAtom.equal(
			&txtSolidLine01,
			nil)

	if areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #3\n"+
			"Expected areEqual WOULD BE 'false'\n"+
			"because 'txtSolidLineTwo' has a 'nil' value.\n"+
			"HOWEVER, 'areEqual' is 'true'!!\n",
			ePrefix.String())

		return
	}

	solidLineChars = "-"
	solidLineCharsRepeatCount = 35
	newLineChars = "\n-\n"

	txtSolidLine03 := TextLineSpecSolidLine{}

	err =
		txtSolidLine03.SetFullSolidLineConfig(
			leftMargin,
			rightMargin,
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			ePrefix.XCpy(
				"txtSolidLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtSolidLine03.IsValidInstance() {
		t.Errorf("%v - ERROR\n"+
			"txtSolidLine03.IsValidInstance()=='false'\n",
			ePrefix.String())
	}

	var txtSolidLine04 TextLineSpecSolidLine

	txtSolidLine04,
		err = txtSolidLine03.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine03->txtSolidLine04"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual =
		txtSolidLineAtom.equal(
			&txtSolidLine03,
			&txtSolidLine04)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #4\n"+
			"Expected areEqual WOULD BE 'true'\n"+
			"because 'txtSolidLine03' and 'txtSolidLine04' are equivalent.\n"+
			"HOWEVER, 'areEqual' is 'false'!!\n",
			ePrefix.String())

		return
	}

	txtSolidLine04.turnLineTerminatorOff = true

	areEqual =
		txtSolidLineAtom.equal(
			&txtSolidLine03,
			&txtSolidLine04)

	if areEqual {

		t.Errorf("%v - ERROR\n"+
			"Test #4\n"+
			"Expected areEqual WOULD BE 'false'\n"+
			"because 'txtSolidLine03' and 'txtSolidLine04' are NOT equal.\n"+
			"txtSolidLine04.turnLineTerminatorOff = true\n"+
			"HOWEVER, 'areEqual' is 'false'!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_getFormattedText_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n"

	expectedSolidLineStr :=
		leftMargin +
			strings.Repeat(
				solidLineChars,
				solidLineCharsRepeatCount) +
			rightMargin +
			newLineChars

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.
		NewPtrFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine01FmtText string

	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	txtSolidLine01FmtText,
		err = txtSolidLineMolecule.getFormattedText(
		txtSolidLine01,
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine01FmtText"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedSolidLineStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(txtSolidLine01FmtText),
			true)

	if printableExpectedStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Formatted Text String 01 DOES NOT match\n"+
			"Actual Formatted Text String 02.\n"+
			"Expected Formatted Text String = '%v'\n"+
			"Instead, Formatted Text String = '%v'\n",
			ePrefix.String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	txtSolidLineMolecule2 := textLineSpecSolidLineMolecule{}

	_,
		err =
		txtSolidLineMolecule2.getFormattedText(
			nil,
			ePrefix.XCpy(
				"txtSolidLine=='nil'"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule2."+
			"getFormattedText()\n"+
			"because 'txtSolidLine' has a 'nil' value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtSolidLine01.rightMarginChars =
		[]rune(strings.Repeat(" ", 1000001))

	_,
		err =
		txtSolidLineMolecule2.getFormattedText(
			txtSolidLine01,
			ePrefix.XCpy(
				"txtSolidLine01 is invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtSolidLineMolecule2."+
			"getFormattedText()\n"+
			"because 'txtSolidLine01.rightMargin = 1000001'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_testValidityOfTextSpecSolidLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_testValidityOfTextSpecSolidLine_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := "-"
	solidLineCharsRepeatCount := 35
	newLineChars := "\n\n"

	txtSolidLine01,
		err := TextLineSpecSolidLine{}.
		NewPtrFullSolidLineRunesConfig(
			[]rune(leftMargin),
			[]rune(rightMargin),
			[]rune(solidLineChars),
			solidLineCharsRepeatCount,
			[]rune(newLineChars),
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine02 TextLineSpecSolidLine

	txtSolidLine02,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02.solidLineCharsRepeatCount = -97

	var isValid bool
	txtSolidLineAtom := textLineSpecSolidLineAtom{}

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine02,
		ePrefix.XCpy(
			"Testing 'txtSolidLine02'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine02.solidLineCharsRepeatCount = -97.\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine02.solidLineCharsRepeatCount = -97\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	_,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine03"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtSolidLine05 TextLineSpecSolidLine

	txtSolidLine05,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine05"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine05.leftMarginChars =
		[]rune(strings.Repeat(" ", 1000001))

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine05,
		ePrefix.XCpy(
			"Testing 'txtSolidLine05'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine05.leftMargin = 1000001.\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine05.leftMargin = 1000001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine06 TextLineSpecSolidLine

	txtSolidLine06,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine06"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine06.rightMarginChars =
		[]rune(strings.Repeat(" ", 1000001))

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine06,
		ePrefix.XCpy(
			"Testing 'txtSolidLine06'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine06.rightMargin = 1000001.\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine06.rightMargin = 1000001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine07 TextLineSpecSolidLine

	txtSolidLine07,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine06"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine07.newLineChars = nil

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine07,
		ePrefix.XCpy(
			"Testing 'txtSolidLine07'"))

	if isValid == false {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'true' because\n"+
			"txtSolidLine07 is a valid instance\n"+
			"HOWEVER, 'isValid' IS 'false'!!!\n",
			ePrefix.String())

		return
	}

	if err != nil {
		t.Errorf("%v - ERROR\n"+
			"Expected no errors from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine07 is a valid instance.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!!!\n"+
			"Error='%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		nil,
		ePrefix.XCpy(
			"'txtSolidLine'=='nil'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine =='nil'.\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine =='nil'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine08 TextLineSpecSolidLine

	txtSolidLine08,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine08"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine08.newLineChars = []rune{'\n', 0, 0, '\n'}

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine08,
		ePrefix.XCpy(
			"Testing 'txtSolidLine08'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine08.newLineChars = []rune{'\\n',0,0,'\\n'}\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine08.newLineChars = []rune{'\\n',0,0,'\\n'}\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	var txtSolidLine09 TextLineSpecSolidLine

	txtSolidLine09,
		err = txtSolidLine01.CopyOut(
		ePrefix.XCpy(
			"txtSolidLine01->txtSolidLine09"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine09.solidLineChars = []rune{'-', 0, 0, '-'}

	isValid,
		err = txtSolidLineAtom.testValidityOfTextSpecSolidLine(
		&txtSolidLine09,
		ePrefix.XCpy(
			"Testing 'txtSolidLine09'"))

	if isValid == true {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' equal 'false' because\n"+
			"txtSolidLine09.solidLineChars = []rune{'-', 0, 0, '-'}\n"+
			"HOWEVER, 'isValid' IS 'true'!!!\n",
			ePrefix.String())

		return
	}

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from \n"+
			"txtSolidLineAtom.testValidityOfTextSpecSolidLine() \n"+
			"because txtSolidLine09.solidLineChars = []rune{'-', 0, 0, '-'}\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecSolidLine_setTxtSolidLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_setTxtSolidLine_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := []rune{'-'}
	solidLineCharsRepeatCount := 35
	newLineChars := []rune{'\n', '\n'}
	var turnLineTerminatorOff bool

	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	turnLineTerminatorOff = false

	err := txtSolidLineMolecule.setTxtSolidLine(
		nil,
		[]rune(leftMargin),
		[]rune(rightMargin),
		solidLineChars,
		solidLineCharsRepeatCount,
		newLineChars,
		turnLineTerminatorOff,
		ePrefix.XCpy(
			"txtSolidLine is 'nil'"))

	if err == nil {
		t.Errorf("%v - ERROR \n"+
			"Expected an error return from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because 'txtSolidLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	txtSolidLine01 := TextLineSpecSolidLine{}

	solidLineChars = nil
	turnLineTerminatorOff = true

	err =
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"solidLineChars length == 0"))

	if err == nil {
		t.Errorf("%v - ERROR \n"+
			"Expected an error return from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because the length of 'solidLineChars' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	solidLineChars = []rune{'-', '*', 0, '-', 0, '*'}

	err =
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"solidLineChars length == 0"))

	if err == nil {
		t.Errorf("%v - ERROR \n"+
			"Expected an error return from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because 'solidLineChars' contains invalid runes.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	solidLineChars = []rune{'-', '-'}

	solidLineCharsRepeatCount = -1

	err =
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"solidLineChars length == 0"))

	if err == nil {
		t.Errorf("%v - ERROR \n"+
			"Expected an error return from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because 'solidLineCharsRepeatCount' is less than one (1).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	solidLineCharsRepeatCount = 5

	leftMargin =
		strings.Repeat(" ", 1000001)

	err =
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"solidLineChars length == 0"))

	if err == nil {
		t.Errorf("%v - ERROR \n"+
			"Expected an error return from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because 'leftMargin' is greater than 1-million (1,000,000).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	leftMargin = "   "
	rightMargin =
		strings.Repeat(" ", 1000001)

	err =
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"solidLineChars length == 0"))

	if err == nil {
		t.Errorf("%v - ERROR \n"+
			"Expected an error return from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because 'rightMargin' is greater than one-million (1,000,000).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	rightMargin = strings.Repeat(" ", 3)

	sMechPreon := strMechPreon{}

	err = sMechPreon.setRepeatRuneChar(
		&newLineChars,
		1000001,
		'X',
		ePrefix.XCpy(
			"newLineChars x 1,000,001"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"newLineChars length == 1,000,001"))

	if err != nil {
		t.Errorf("%v - ERROR \n"+
			"Expected no errors from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because 'newLineChars' is greater than one-million (1,000,000).\n"+
			"HOWEVER, AN ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	newLineChars = []rune{'\n', 0, '\n', 0, '\n'}

	err =
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"solidLineChars length == 0"))

	if err == nil {
		t.Errorf("%v - ERROR \n"+
			"Expected an error return from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because 'newLineChars'contains invalid runes.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecSolidLine_setTxtSolidLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecSolidLine_setTxtSolidLine_000100()",
		"")

	leftMargin := "  "
	rightMargin := "  "
	solidLineChars := []rune{'-'}
	solidLineCharsRepeatCount := 35
	newLineChars := []rune{'\n', '\n'}
	var turnLineTerminatorOff bool

	txtSolidLine01 := TextLineSpecSolidLine{}

	txtSolidLineMolecule := textLineSpecSolidLineMolecule{}

	err :=
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine01,
			[]rune(leftMargin),
			[]rune(rightMargin),
			solidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"txtSolidLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSolidLine02 := TextLineSpecSolidLine{}

	badSolidLineChars := []rune{'-', 0, '*', 0, '?'}

	err =
		txtSolidLineMolecule.setTxtSolidLine(
			&txtSolidLine02,
			[]rune(leftMargin),
			[]rune(rightMargin),
			badSolidLineChars,
			solidLineCharsRepeatCount,
			newLineChars,
			turnLineTerminatorOff,
			ePrefix.XCpy(
				"txtSolidLine01 + badSolidLineChars"))

	if err == nil {

		t.Errorf("%v - ERROR \n"+
			"Expected an error return from txtSolidLineMolecule."+
			"setTxtSolidLine()\n"+
			"because 'badSolidLineChars' is an invalid rune array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	return
}

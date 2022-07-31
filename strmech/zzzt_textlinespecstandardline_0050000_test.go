package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func TestTextLineSpecStandardLineAtom_copyTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineAtom_copyTextFields_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
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

	var textFields01, textFields02 []ITextFieldSpecification

	textFields01,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	textFields02,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields02<-stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLineAtom := textLineSpecStandardLineAtom{}

	var lenSourceTxtFieldArray, lenTargetTxtFieldArray int

	lenSourceTxtFieldArray = len(textFields01)

	lenTargetTxtFieldArray,
		err = stdLineAtom.copyTextFields(
		&textFields02,
		&textFields01,
		ePrefix.XCpy(
			"textFields02<-textFields01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if lenSourceTxtFieldArray != lenTargetTxtFieldArray {

		t.Errorf("\n%v - ERROR\n"+
			"Expected length of source array == "+
			"length of target array\n"+
			"after stdLineAtom.copyTextFields()"+
			"HOWEVER, THE ARRAY LENGTHS ARE NOT EQUAL!\n"+
			"Source Text Field Array Length = '%v'\n"+
			"Target Text Field Array Length = '%v'\n",
			ePrefix.String(),
			lenSourceTxtFieldArray,
			lenTargetTxtFieldArray)

		return
	}

	stdLine02.EmptyTextFields()

	_,
		err = stdLine02.AddTextFields(
		&textFields02,
		ePrefix.XCpy(
			"stdLine02<-textFields02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLineMolecule := textLineSpecStandardLineMolecule{}

	areEqual := stdLineMolecule.equal(
		&stdLine01,
		&stdLine02)

	if !areEqual {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine01 == stdLine02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLineAtom_copyTextFields_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineAtom_copyTextFields_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var textFields01, textFields02 []ITextFieldSpecification

	textFields01,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"->stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	textFields02,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields02<-stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLineAtom := textLineSpecStandardLineAtom{}

	_,
		err = stdLineAtom.copyTextFields(
		nil,
		&textFields01,
		ePrefix.XCpy(
			"targetTextFields = nil pointer"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineAtom.copyTextFields()\n"+
			"because 'targetTextFields' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err = stdLineAtom.copyTextFields(
		&textFields02,
		nil,
		ePrefix.XCpy(
			"sourceTextFields = nil pointer"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineAtom.copyTextFields()\n"+
			"because 'sourceTextFields' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	textFields01 = nil

	_,
		err = stdLineAtom.copyTextFields(
		&textFields02,
		&textFields01,
		ePrefix.XCpy(
			"sourceTextFields array is nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineAtom.copyTextFields()\n"+
			"because 'sourceTextFields' array is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLineAtom_copyTextFields_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineAtom_copyTextFields_000300()",
		"")

	var stdLine01, stdLine02 TextLineSpecStandardLine
	var err error

	stdLine01,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01<-"))

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
	var textFields01, textFields02 []ITextFieldSpecification

	textFields01,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine02<-"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields02,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields02<-stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if len(textFields01) == len(textFields02) {

		t.Errorf("\n%v - ERROR\n"+
			"len(textFields01) == len(textFields02)\n"+
			"The length of these Text Fields should be different.\n"+
			"HOWEVER, THE LENGTHS ARE EQUAL!\n",
			ePrefix.String())

		return

	}

	stdLineAtom := textLineSpecStandardLineAtom{}

	_,
		err = stdLineAtom.copyTextFields(
		&textFields02,
		&textFields01,
		ePrefix.XCpy(
			"textFields01 -> textFields02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	areEqual := textLineSpecStandardLineElectron{}.ptr().
		equalTextFieldArrays(
			&textFields01,
			&textFields02)

	if !areEqual {

		t.Errorf("\n%v - ERROR\n"+
			"stdLineAtom.copyTextFields() Failed!\n"+
			"textFields01 is NOT EQUAL to textFields02.\n",
			ePrefix.String())

		return
	}

	stdLine02.EmptyTextFields()

	_,
		err = stdLine02.AddTextFields(
		&textFields01,
		ePrefix.XCpy(
			"stdLine02<-textFields01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLineMolecule := textLineSpecStandardLineMolecule{}

	areEqual = stdLineMolecule.equal(
		&stdLine01,
		&stdLine02)

	if !areEqual {

		t.Errorf("\n%v - ERROR\n"+
			"Expected stdLine01 == stdLine02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	numOfTextFields := stdLine02.GetNumOfTextFields()

	if numOfTextFields != 6 {

		t.Errorf("\n%v - ERROR\n"+
			"'stdLine02' should contain 6 Text Fields\n"+
			"Instead, it contains '%v' Text Fields!\n",
			ePrefix.String(),
			numOfTextFields)

		return

	}

	return
}

func TestTextLineSpecStandardLineAtom_peekPopTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineAtom_peekPopTextField_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
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

	txtLineAtom := textLineSpecStandardLineAtom{}

	_,
		err = txtLineAtom.peekPopTextField(
		nil,
		5,
		false,
		ePrefix.XCpy(
			"txtStdLine is 'nil'"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLineAtom.peekPopTextField()\n"+
			"because 'txtStdLine' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	_,
		err = txtLineAtom.peekPopTextField(
		&stdLine01,
		6,
		false,
		ePrefix.XCpy(
			"indexId is 6. Last Index is 5."))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLineAtom.peekPopTextField()\n"+
			"because 'indexId' is exceeds last index of array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	_,
		err = txtLineAtom.peekPopTextField(
		&stdLine01,
		-1,
		false,
		ePrefix.XCpy(
			"indexId is -1"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLineAtom.peekPopTextField()\n"+
			"because 'indexId' is less than zero.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	stdLine01.textFields = make([]ITextFieldSpecification, 0)

	_,
		err = txtLineAtom.peekPopTextField(
		&stdLine01,
		1,
		false,
		ePrefix.XCpy(
			"txtStdLine is 'nil'"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtLineAtom.peekPopTextField()\n"+
			"because 'stdLine01.textFields' has a length of zero.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTextLineSpecStandardLineAtom_testValidityOfTextLineSpecStdLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineAtom_testValidityOfTextLineSpecStdLine_000100()",
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

	txtStdLineAtom := textLineSpecStandardLineAtom{}

	_,
		_ = txtStdLineAtom.testValidityOfTextLineSpecStdLine(
		&stdLine01,
		true,
		ePrefix.XCpy(
			"stdLine01"))

	var isValid bool

	var stdLine02 *TextLineSpecStandardLine

	isValid,
		err = txtStdLineAtom.testValidityOfTextLineSpecStdLine(
		stdLine02,
		true,
		ePrefix.XCpy(
			"stdLine02"))

	if isValid {

		t.Errorf("\n%v - ERROR\n"+
			"isValid = stdLine02.testValidityOfTextLineSpecStdLine()\n"+
			"Expected 'isValid' equal 'false' \n"+
			"because input parameter 'txtStdLine' is 'nil'.\n"+
			"HOWEVER, 'isValid' == true !\n",
			ePrefix.String())

		return

	}

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLine02.testValidityOfTextLineSpecStdLine()\n"+
			"because input parameter 'txtStdLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLineAtom_testValidityOfTextLineSpecStdLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineAtom_testValidityOfTextLineSpecStdLine_000200()",
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

	stdLine01.newLineChars = nil

	txtStdLineAtom := textLineSpecStandardLineAtom{}

	var isValid bool

	isValid,
		err = txtStdLineAtom.testValidityOfTextLineSpecStdLine(
		&stdLine01,
		true,
		ePrefix.XCpy(
			"stdLine02"))

	if isValid == false {

		t.Errorf("\n%v - ERROR\n"+
			"isValid = txtStdLineAtom.testValidityOfTextLineSpecStdLine()\n"+
			"Expected 'isValid' equal 'true' \n"+
			"because input parameter 'stdLine01' is valid.\n"+
			"HOWEVER, 'isValid' == false !\n",
			ePrefix.String())

		return

	}

	if err != nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected no errors from txtStdLineAtom.testValidityOfTextLineSpecStdLine()\n"+
			"because input parameter 'stdLine01' is valid.\n"+
			"HOWEVER, AN ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLineElectron_deleteTextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_deleteTextField_000100()",
		"")

	indexId := 2

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	err := txtStdLineElectron.deleteTextField(
		nil,
		indexId,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error from txtStdLineElectron.deleteTextField()\n"+
			"because txtStdLine is a 'nil' pointer!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtStdLineElectron02 := textLineSpecStandardLineElectron{}

	indexId = 18

	err = txtStdLineElectron02.deleteTextField(
		stdLine01,
		indexId,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error from txtStdLineElectron02.deleteTextField()\n"+
			"because indexId is invalid!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	indexId = -2

	err = txtStdLineElectron02.deleteTextField(
		stdLine01,
		indexId,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error from txtStdLineElectron02.deleteTextField()\n"+
			"because indexId is less than zero!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	indexId = 0

	stdLine01.textFields = nil

	err = txtStdLineElectron02.deleteTextField(
		stdLine01,
		indexId,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error from txtStdLineElectron02.deleteTextField()\n"+
			"because stdLine01.textFields = nil!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLineElectron_deleteTextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_deleteTextField_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine02(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		textLineSpecStandardLineElectron{}.ptr().deleteTextField(
			&stdLine01,
			1,
			ePrefix.XCpy(
				"stdLine01 delete index 1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecStandardLineElectron_emptyTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_emptyTextFields_000100()",
		"")

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	var nilPtr *[]ITextFieldSpecification

	_ = txtStdLineElectron.emptyTextFields(
		nilPtr,
		ePrefix.XCpy(
			""))

}

func TestTextLineSpecStandardLineElectron_equalTextFieldArrays_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_equalTextFieldArrays_000100()",
		"")

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	var nilPtr01, nilPtr02 []ITextFieldSpecification

	_ =
		txtStdLineElectron.equalTextFieldArrays(
			&nilPtr01,
			&nilPtr02)

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields01, textFields02 []ITextFieldSpecification

	textFields01,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine01"))

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

	textFields02,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields02<-stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields02[2] = nil

	_ =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	textFields01[2] = nil

	_ =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01 #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields01,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine01 #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine02 #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields02,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine02 #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_ =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	textFields01 = nil

	_ =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	_ =
		txtStdLineElectron.equalTextFieldArrays(
			nil,
			&textFields02)

	textFields02 = nil

	_ =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01 #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields01,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine01 #3"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02 #2"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields02,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields02<-stdLine02 #3"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	newLabelText := "Xray97 where are?"

	var labelTxt TextFieldSpecLabel

	labelTxt,
		err = TextFieldSpecLabel{}.NewTextLabel(
		newLabelText,
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine02.ReplaceTextField(
		&labelTxt,
		1,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	textFields02,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine02 #4"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	compareResult :=
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	if compareResult == true {

		t.Errorf("%v - Error\n"+
			"Test #1 \n"+
			"compareResult = txtStdLineElectron.equalTextFieldArrays()\n "+
			"Expected compareResult = false\n"+
			"HOWEVER, compareResult = true!\n",
			ePrefix.String())

		return

	}

	textFields01 = make([]ITextFieldSpecification, 0)

	compareResult =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	if compareResult == true {

		t.Errorf("%v - Error\n"+
			"Test #2 \n"+
			"compareResult = txtStdLineElectron.equalTextFieldArrays()\n "+
			"Expected compareResult = false\n"+
			"HOWEVER, compareResult = true!\n",
			ePrefix.String())

		return

	}

	textFields01,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine01 #4"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields02 = make([]ITextFieldSpecification, 0)

	compareResult =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	if compareResult == true {

		t.Errorf("%v - Error\n"+
			"Test #3 \n"+
			"compareResult = txtStdLineElectron.equalTextFieldArrays()\n "+
			"Expected compareResult = false\n"+
			"HOWEVER, compareResult = true!\n",
			ePrefix.String())

		return
	}

	textFields01 = make([]ITextFieldSpecification, 0)

	textFields02,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine02 #4"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	compareResult =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	if compareResult == true {

		t.Errorf("%v - Error\n"+
			"Test #4 \n"+
			"compareResult = txtStdLineElectron.equalTextFieldArrays()\n "+
			"Expected compareResult = false\n"+
			"HOWEVER, compareResult = true!\n",
			ePrefix.String())

		return
	}

	textFields02 = make([]ITextFieldSpecification, 0)

	compareResult =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	if compareResult == false {

		t.Errorf("%v - Error\n"+
			"Test #5 \n"+
			"compareResult = txtStdLineElectron.equalTextFieldArrays()\n "+
			"Expected compareResult = true\n"+
			"HOWEVER, compareResult = false!\n",
			ePrefix.String())

		return
	}

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01 #5"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields01,
		err = stdLine01.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine01 #5"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	stdLine02,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine02 #5"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields02,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields01<-stdLine02 #5"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields01[1] = nil

	compareResult =
		txtStdLineElectron.equalTextFieldArrays(
			&textFields01,
			&textFields02)

	if compareResult == true {

		t.Errorf("%v - Error\n"+
			"Test #6 \n"+
			"compareResult = txtStdLineElectron.equalTextFieldArrays()\n "+
			"Expected compareResult = false\n"+
			"HOWEVER, compareResult = true!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000100()",
		"")

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	_,
		err := txtStdLineElectron.testValidityOfTextFields(
		nil,
		false, // allowZeroLengthTextFieldsArray
		ePrefix.XCpy(
			"txtFields is empty"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtStdLineElectron.emptyStandardLine()\n"+
			"because 'txtFields' is  'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err =
		stdLine01.GetTextFieldCollection(
			ePrefix.XCpy(
				"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	textFields[1] = nil

	_,
		err = txtStdLineElectron.testValidityOfTextFields(
		&textFields,
		false, // allowZeroLengthTextFieldsArray
		ePrefix.XCpy(
			"textFields[1] = nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtStdLineElectron.emptyStandardLine()\n"+
			"because 'txtFields[1]' is  'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	textFields,
		err =
		stdLine01.GetTextFieldCollection(
			ePrefix.XCpy(
				"#2 textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	textFields[1].Empty()

	_,
		err = txtStdLineElectron.testValidityOfTextFields(
		&textFields,
		false, // allowZeroLengthTextFieldsArray
		ePrefix.XCpy(
			"txtFields[1] is invalid"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtStdLineElectron.emptyStandardLine()\n"+
			"because 'txtFields[1]' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineElectron_testValidityOfTextFields_000300()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err =
		stdLine01.GetTextFieldCollection(
			ePrefix.XCpy(
				"textFields<-stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var expectedTextFieldsLength = len(textFields)

	var actualTextFieldsLength int

	txtStdLineElectron := textLineSpecStandardLineElectron{}

	actualTextFieldsLength,
		err = txtStdLineElectron.testValidityOfTextFields(
		&textFields,
		false, // allowZeroLengthTextFieldsArray
		ePrefix.XCpy(
			"textFields[1] = nil"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if expectedTextFieldsLength != actualTextFieldsLength {

		t.Errorf("\n%v - ERROR\n"+
			"expectedTextFieldsLength != actualTextFieldsLength\n"+
			"txtStdLineElectron.emptyStandardLine() returned\n"+
			"an invalid Text Fields Array Length.\n"+
			"expectedTextFieldsLength = '%v'\n"+
			"  actualTextFieldsLength = '%v'\n",
			ePrefix.String(),
			expectedTextFieldsLength,
			actualTextFieldsLength)

		return

	}

	return
}

func TestTextLineSpecStandardLineMolecule_emptyStandardLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineMolecule_emptyStandardLine_000100()",
		"")

	txtStdLineMolecule := textLineSpecStandardLineMolecule{}

	err := txtStdLineMolecule.emptyStandardLine(
		nil,
		ePrefix.XCpy(
			"txtStdLine is 'nil'"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from txtStdLineMolecule.emptyStandardLine()\n"+
			"because 'txtStdLine' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLineMolecule_emptyStandardLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineMolecule_emptyStandardLine_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtStdLineMolecule := textLineSpecStandardLineMolecule{}

	stdLine01.textFields[1] = nil

	err = txtStdLineMolecule.emptyStandardLine(
		&stdLine01,
		ePrefix.XCpy(
			"stdLine01.textFields[1] = nil"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecStandardLineMolecule_emptyStdLineTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineMolecule_emptyStdLineTextFields_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtStdLineMolecule := textLineSpecStandardLineMolecule{}

	txtStdLineMolecule.emptyStdLineTextFields(
		nil)

	txtStdLineMolecule.emptyStdLineTextFields(
		&stdLine01)

}

func TestTextLineSpecStandardLineMolecule_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineMolecule_equal_000100()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtStdLineMolecule := textLineSpecStandardLineMolecule{}

	areEqual := txtStdLineMolecule.equal(
		&stdLine01,
		nil)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Expected txtStdLineMolecule.equal() return false.\n"+
			"because 'stdLine02' is a 'nil' pointer.\n"+
			"HOWEVER, THE RETURN VALUE IS 'true'!!!\n",
			ePrefix.String())

		return

	}

	var stdLine02 TextLineSpecStandardLine

	stdLine02,
		err = createTestTextLineSpecStandardLine04(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual = txtStdLineMolecule.equal(
		nil,
		&stdLine02)

	if areEqual == true {

		t.Errorf("\n%v - ERROR\n"+
			"Expected txtStdLineMolecule.equal() return false.\n"+
			"because 'stdLine01' is a 'nil' pointer.\n"+
			"HOWEVER, THE RETURN VALUE IS 'true'!!!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTextLineSpecStandardLineMolecule_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineMolecule_getFormattedText_000100()",
		"")

	stdLineMolecule := textLineSpecStandardLineMolecule{}

	sb := strings.Builder{}

	_,
		_,
		err := stdLineMolecule.getFormattedText(
		&sb,
		nil,
		ePrefix.XCpy(
			"txtStdLine is 'nil'"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLineMolecule.getFormattedText()\n"+
			"Expected an error return because parameter\n"+
			"'txtStdLine' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	var stdLine01 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine04(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine01.numOfStdLines = -5

	sb.Reset()

	_,
		_,
		err = stdLineMolecule.getFormattedText(
		&sb,
		&stdLine01,
		ePrefix.XCpy(
			"txtStdLine is 'nil'"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLineMolecule.getFormattedText()\n"+
			"Expected an error return because parameter\n"+
			"'stdLine01.numOfStdLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTextLineSpecStandardLineMolecule_getFormattedText_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineMolecule_getFormattedText_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLineMolecule := textLineSpecStandardLineMolecule{}

	stdLine01.textFields = nil

	sb := strings.Builder{}

	_,
		_,
		err = stdLineMolecule.getFormattedText(
		&sb,
		&stdLine01,
		ePrefix.XCpy(
			"stdLine01 is invalid"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLineMolecule.getFormattedText()\n"+
			"Expected an error return because parameter\n"+
			"stdLine01.textFields = nil.\n"+
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
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02.textFields[1] = nil

	sb.Reset()

	_,
		_,
		err = stdLineMolecule.getFormattedText(
		&sb,
		&stdLine01,
		ePrefix.XCpy(
			"stdLine02 is invalid"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: stdLineMolecule.getFormattedText()\n"+
			"Expected an error return because parameter\n"+
			"stdLine02.textFields[1] = nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTestTextLineSpecStandardLineNanobot_addTextFields_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTestTextLineSpecStandardLineNanobot_addTextFields_000100()",
		"")

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine02"))

	stdLineNanobot := textLineSpecStandardLineNanobot{}

	_,
		err =
		stdLineNanobot.addTextFields(
			nil,
			&textFields,
			ePrefix.XCpy(
				"txtStdLine is nil pointer"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.addTextFields()\n"+
			"because 'txtStdLine' input parameter is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return

	}

	stdLine01 := TextLineSpecStandardLine{}.New()

	_,
		err =
		stdLineNanobot.addTextFields(
			&stdLine01,
			nil,
			ePrefix.XCpy(
				"textFields is nil pointer"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.addTextFields()\n"+
			"because 'textFields' input parameter is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return

	}

}

func TestTestTextLineSpecStandardLineNanobot_addTextFields_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTestTextLineSpecStandardLineNanobot_addTextFields_000200()",
		"")

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine02"))

	stdLineNanobot := textLineSpecStandardLineNanobot{}

	textFields[1].Empty()

	stdLine01 := TextLineSpecStandardLine{}.New()

	_,
		err =
		stdLineNanobot.addTextFields(
			&stdLine01,
			&textFields,
			ePrefix.XCpy(
				"textFields[1] is invalid"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.addTextFields()\n"+
			"because 'textFields[1]' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return

	}

	return
}

func TestTextLineSpecStandardLineNanobot_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineNanobot_copyIn_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	txtStdLineNanobot := textLineSpecStandardLineNanobot{}

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			stdLine01,
			ePrefix.XCpy(
				"stdLine02<-stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v\n"+
			"Error: Expected stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			nil,
			ePrefix.XCpy(
				"incomingStdLine==nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'incomingStdLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtStdLineNanobot.copyIn(
			nil,
			stdLine01,
			ePrefix.XCpy(
				"nil<-stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'targetStdLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine01.numOfStdLines = -47

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			stdLine01,
			ePrefix.XCpy(
				"stdLine02<-invalid stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'targetStdLine' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine01.numOfStdLines = 1

	stdLine01.newLineChars = nil

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			stdLine01,
			ePrefix.XCpy(
				"stdLine02<-stdLine01 newLineChars==nil"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine01.newLineChars = []rune{'\n', 0, '\n', 0}

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			stdLine01,
			ePrefix.XCpy(
				"stdLine02<-stdLine01 newLineChars invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'stdLine01.newLineChars' contains invalid characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLineNanobot_copyIn_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineNanobot_copyIn_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}.New()

	stdLine01.newLineChars = []rune{'\n', 0, '\n', 0}

	txtStdLineNanobot := textLineSpecStandardLineNanobot{}

	err =
		txtStdLineNanobot.copyIn(
			&stdLine02,
			&stdLine01,
			ePrefix.XCpy(
				"stdLine02<-stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyIn()\n"+
			"because 'stdLine01.newLineChars' are invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTestTextLineSpecStandardLineNanobot_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTestTextLineSpecStandardLineNanobot_copyOut_000100()",
		"")

	stdLine01 := TextLineSpecStandardLine{}.NewPtr()

	rightMarginLen := 5

	rightMarginSpec,
		err := TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var leftMarginSpec TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := TxtJustify.Center()

	var labelSpec TextFieldSpecLabel

	labelSpec,
		err = TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine02 := TextLineSpecStandardLine{}

	txtStdLineNanobot := textLineSpecStandardLineNanobot{}

	stdLine02,
		err =
		txtStdLineNanobot.copyOut(
			stdLine01,
			ePrefix.XCpy(
				"stdLine02<-stdLine01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !stdLine02.Equal(stdLine01) {
		t.Errorf("%v\n"+
			"Error: Expected stdLine02 == stdLine01\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	_,
		err =
		txtStdLineNanobot.copyOut(
			nil,
			ePrefix.XCpy(
				"txtStdLine==nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyOut()\n"+
			"because 'txtStdLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine01.numOfStdLines = -47
	_,
		err =
		txtStdLineNanobot.copyOut(
			stdLine01,
			ePrefix.XCpy(
				"stdLine01 is invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyOut()\n"+
			"because 'stdLine01.numOfStdLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine01.numOfStdLines = 1
	stdLine01.newLineChars = nil
	// nil newLinChars defaults to '\n'

	_,
		err =
		txtStdLineNanobot.copyOut(
			stdLine01,
			ePrefix.XCpy(
				"stdLine01.newLineChars==nil"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLine01.newLineChars = []rune{'\n', 0, '\n', 0}
	_,
		err =
		txtStdLineNanobot.copyOut(
			stdLine01,
			ePrefix.XCpy(
				"stdLine01 is invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyOut()\n"+
			"because 'stdLine01.newLineChars' contains invalid characters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTestTextLineSpecStandardLineNanobot_copyOut_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTestTextLineSpecStandardLineNanobot_copyOut_000200()",
		"")

	stdLine01,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtStdLineNanobot := textLineSpecStandardLineNanobot{}

	stdLine01.newLineChars = []rune{'\n', 0, '\n'}

	_,
		err = txtStdLineNanobot.copyOut(
		&stdLine01,
		ePrefix.XCpy(
			"stdLine02<-stdLine01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtStdLineNanobot{}."+
			"copyOut()\n"+
			"because 'stdLine01.newLineChars' are invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecStandardLineNanobot_insertTextFieldAtIndex_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineNanobot_insertTextFieldAtIndex_000100()",
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
		err = stdLine01.PeekAtTextField(
		2,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	txtStdLineNanobot := textLineSpecStandardLineNanobot{}

	_,
		err = txtStdLineNanobot.insertTextFieldAtIndex(
		nil,
		iTxtFieldSpec,
		2,
		ePrefix.XCpy(
			"txtStdLine==nil"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"txtStdLineNanobot.insertTextFieldAtIndex()\n"+
			"Expected an error return because\n"+
			"input parameter 'txtStdLine' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	_,
		err = txtStdLineNanobot.insertTextFieldAtIndex(
		&stdLine01,
		nil,
		2,
		ePrefix.XCpy(
			"iTextField==nil"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"txtStdLineNanobot.insertTextFieldAtIndex()\n"+
			"Expected an error return because\n"+
			"input parameter 'iTextField' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	var labelTxt TextFieldSpecLabel

	labelTxt,
		err = TextFieldSpecLabel{}.NewTextLabel(
		"Xray7 where are?",
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	labelTxt.textLabel = nil

	labelTxt.fieldLen = -542

	_,
		err = txtStdLineNanobot.insertTextFieldAtIndex(
		&stdLine01,
		&labelTxt,
		2,
		ePrefix.XCpy(
			"labelTxt is invalid"))

	if err == nil {

		t.Errorf("%v - Error\n"+
			"txtStdLineNanobot.insertTextFieldAtIndex()\n"+
			"Expected an error return because\n"+
			"input parameter 'labelTxt' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	labelTxt,
		err = TextFieldSpecLabel{}.NewTextLabel(
		"Xray7 where are?",
		-1,
		TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtStdLineNanobot.insertTextFieldAtIndex(
		&stdLine01,
		&labelTxt,
		972,
		ePrefix.XCpy(
			"indexId is out of range"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtStdLineNanobot.insertTextFieldAtIndex(
		&stdLine01,
		&labelTxt,
		-92,
		ePrefix.XCpy(
			"labelTxt is invalid"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecStandardLineNanobot_setTxtSpecStandardLine_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTESTSERIES_FUNCNAME_000100()",
		"")

	numOfStdLines := 1

	newLineChars := []rune{'\n'}

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	stdLineNanobot := textLineSpecStandardLineNanobot{}

	err = stdLineNanobot.setTxtSpecStandardLine(
		nil,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCpy(
			"txtStdLine==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'txtStdLine' is  'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	stdLine01 := TextLineSpecStandardLine{}.New()

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		nil,
		newLineChars,
		false,
		ePrefix.XCpy(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'textFields' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	numOfStdLines = -1

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCpy(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'numOfStdLines' is  '-1'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	numOfStdLines = 1000001

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCpy(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'numOfStdLines' is  '1000001'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	numOfStdLines = 1
	newLineChars = nil

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCpy(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'newLineChars' is  'nill'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	numOfStdLines = 1
	newLineChars = []rune{'\n', 0, '\n', 0}

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCpy(
			"textFields[]==nil"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'newLineChars' is  invalid'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextLineSpecStandardLineNanobot_setTxtSpecStandardLine_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecStandardLineNanobot_setTxtSpecStandardLine_000200()",
		"")

	stdLine02,
		err := createTestTextLineSpecStandardLine01(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var textFields []ITextFieldSpecification

	textFields,
		err = stdLine02.GetTextFieldCollection(
		ePrefix.XCpy(
			"textFields<-stdLine02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	numOfStdLines := 1

	newLineChars := []rune{'\n'}

	textFields[1].Empty()

	stdLineNanobot := textLineSpecStandardLineNanobot{}

	stdLine01 := TextLineSpecStandardLine{}.New()

	err = stdLineNanobot.setTxtSpecStandardLine(
		&stdLine01,
		numOfStdLines,
		textFields,
		newLineChars,
		false,
		ePrefix.XCpy(
			"textFields[1] invalid"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from stdLineNanobot.setTxtSpecStandardLine()\n"+
			"because 'textFields[1]' is invalid'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

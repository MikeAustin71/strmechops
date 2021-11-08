package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

func TestTextFieldSpecSpacer_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyIn_000100()",
		"")

	expectedFieldLen := 4

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("txtFieldSpacerOne - Error\n"+
			"%v\n",
			err.Error())
		return
	}

	err = txtFieldSpacerOne.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualFieldLen := txtFieldSpacerOne.GetFieldLength()

	if expectedFieldLen != actualFieldLen {

		t.Errorf("%v\n"+
			"Error: Invalid Field Length returned by\n"+
			"txtFieldSpacerOne.GetFieldLength()\n"+
			"Expected Field Length = '%v'\n"+
			"Instead, Actual Field Length = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedFieldLen,
			actualFieldLen)

		return
	}

	var txtFieldSpacerTwo TextFieldSpecSpacer

	txtFieldSpacerTwo,
		err = TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldSpacerTwo.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: txtFieldSpacerTwo.IsValidInstance()\n"+
			"returned a value of 'false'!\n",
			ePrefix.String())

		return
	}

	if !txtFieldSpacerTwo.Equal(txtFieldSpacerOne) {
		t.Errorf("%v\n"+
			"Error: txtFieldSpacerOne should be equal to txtFieldSpacerTwo\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	badFieldLen := -9
	txtSpacerElectron := textFieldSpecSpacerElectron{}

	var isValid bool

	isValid,
		err =
		txtSpacerElectron.isFieldLenValidError(
			badFieldLen,
			&ePrefix)

	if err == nil {

		t.Errorf("%v\n"+
			"Error: Expected an error return from textFieldSpecSpacerElectron.isFieldLenValidError()\n"+
			"because field length = -9\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	if isValid {

		t.Errorf("%v\n"+
			"Error: Expected textFieldSpecSpacerElectron.isFieldLenValidError()\n"+
			"to return 'false' because field length = -9\n"+
			"HOWEVER, THE RETURN VALUE IS 'true'!\n",
			ePrefix.String())

		return
	}

	txtFieldSpacerTwo.fieldLen = badFieldLen

	isValid =
		txtFieldSpacerTwo.IsValidInstance()

	if isValid {

		t.Errorf("%v\n"+
			"Error: Expected txtFieldSpacerTwo.IsValidInstance()\n"+
			"to return 'false' because txtFieldSpacerTwo.fieldLen = -9\n"+
			"HOWEVER, THE RETURN VALUE IS 'true'!\n",
			ePrefix.String())

		return
	}

	err =
		txtFieldSpacerTwo.IsValidInstanceError(
			ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldSpacerTwo.IsValidInstanceError()\n"+
			"to return an error because field length = -9\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtFieldSpacerThree := &TextFieldSpecSpacer{}

	txtFieldNanobot := textFieldSpecSpacerNanobot{}

	err =
		txtFieldNanobot.copyIn(
			txtFieldSpacerThree,
			&txtFieldSpacerTwo,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldNanobot.copyIn()\n"+
			"to return an error because txtFieldSpacerTwo.fieldLen = -9\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtFieldSpacerTwo.CopyIn(
			txtFieldSpacerOne,
			ePrefix.XCtx(
				"txtFieldSpacerOne->txtFieldSpacerTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldSpacerTwo.Equal(txtFieldSpacerOne) {
		t.Errorf("%v Test #2\n"+
			"Error: txtFieldSpacerOne should be equal to txtFieldSpacerTwo\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	err =
		txtFieldSpacerTwo.CopyIn(
			txtFieldSpacerOne,
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerTwo."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_CopyIn_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyIn_000200()",
		"")

	expectedFieldLen := 4

	txtFieldSpacerOne := TextFieldSpecSpacer{}
	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	err :=
		txtFieldSpacerTwo.SetFieldLen(
			expectedFieldLen,
			ePrefix.XCtx(
				"expectedFieldLen=4"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		txtFieldSpacerOne.CopyIn(
			&txtFieldSpacerTwo,
			ePrefix.XCtx(
				"txtFieldSpacerTwo->txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldSpacerOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("txtFieldSpacerOne Error\n"+
			"%v\n",
			err.Error())
		return
	}

	txtSpacerNanobot2 := textFieldSpecSpacerNanobot{}

	err =
		txtSpacerNanobot2.copyIn(
			nil,
			&txtFieldSpacerOne,
			ePrefix.XCtx(
				"txtFieldSpacerOne->nil"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtSpacerNanobot2.copyIn()\n"+
			"because 'targetTxtFieldSpacer' is a nil pointer!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

	}

	txtSpacerNanobot3 := textFieldSpecSpacerNanobot{}

	err =
		txtSpacerNanobot3.copyIn(
			&txtFieldSpacerOne,
			nil,
			ePrefix.XCtx(
				"nil->txtFieldSpacerOne"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtSpacerNanobot2.copyIn()\n"+
			"because 'targetTxtFieldSpacer' is a nil pointer!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

	}

	return
}

func TestTextFieldSpecSpacer_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyOut_000100()",
		"")

	expectedFieldLen := 9

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtFieldSpacerTwo TextFieldSpecSpacer

	txtFieldSpacerTwo,
		err = txtFieldSpacerOne.CopyOut(
		ePrefix.XCtx("txtFieldSpacerOne->txtFieldSpacerTwo"))

	var txtFieldSpacerThree TextFieldSpecSpacer

	txtFieldSpacerNanobot :=
		textFieldSpecSpacerNanobot{}

	txtFieldSpacerThree,
		err =
		txtFieldSpacerNanobot.copyOut(
			nil,
			ePrefix.XCtx(
				"nil -> txtFieldSpacerThree"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldSpacerNanobot.copyOut()\n"+
			"because 'targetTxtFieldSpacer' is a nil pointer!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

	}

	txtFieldSpacerTwo.fieldLen = 1000001

	txtFieldSpacerThree,
		err =
		txtFieldSpacerNanobot.copyOut(
			&txtFieldSpacerTwo,
			ePrefix.XCtx(
				"txtFieldSpacerTwo->txtFieldSpacerThree"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldSpacerNanobot.copyOut()\n"+
			"because txtFieldSpacerTwo.fieldLen = -59!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtFieldSpacerThree,
		err =
		txtFieldSpacerOne.CopyOut(
			ePrefix.XCtx(
				"txtFieldSpacerOne->txtFieldSpacerThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldSpacerThree.Equal(&txtFieldSpacerOne) {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldSpacerThree would equal txtFieldSpacerOne.\n"+
			"HOWEVER, THE ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	if txtFieldSpacerThree.Equal(&txtFieldSpacerTwo) {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldSpacerThree would NOT equal txtFieldSpacerTwo.\n"+
			"HOWEVER, THE ARE EQUAL!\n",
			ePrefix.String())

		return
	}

	_,
		err =
		txtFieldSpacerTwo.CopyOut(
			ePrefix.XCtx(
				"txtFieldSpacerTwo->"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldSpacerTwo.CopyOut()\n"+
			"because txtFieldSpacerTwo.fieldLen = -59!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		err =
		txtFieldSpacerOne.CopyOut(
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerOne."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_CopyOut_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyOut_000200()",
		"")

	txtFieldSpacerOne := TextFieldSpecSpacer{}

	_,
		err :=
		txtFieldSpacerOne.CopyOut(
			ePrefix.XCtx(
				"txtFieldSpacerOne"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldSpacerOne.CopyOut()\n"+
			"because 'txtFieldSpacerOne' is empty and uninitialized!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_CopyOutITextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyOutITextField_000100()",
		"")

	expectedFieldLen := 9

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var iTextField ITextFieldSpecification

	iTextField,
		err =
		txtFieldSpacerOne.CopyOutITextField(
			ePrefix.XCtx("iTextField"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSpacer, ok := iTextField.(*TextFieldSpecSpacer)

	if !ok {
		t.Errorf("%v\n"+
			"Error: Expected iTextField.(*TextFieldSpecSpacer)\n"+
			"would return an object of type *TextFieldSpecSpacer!\n"+
			"HOWEVER, the type cast FAILED!\n",
			ePrefix.String())
	}

	actualFieldLen := txtSpacer.GetFieldLength()

	if expectedFieldLen != actualFieldLen {

		t.Errorf("%v\n"+
			"Error: Invalid Field Length returned by\n"+
			"txtSpacer.GetFieldLength()\n"+
			"Expected Field Length = '%v'\n"+
			"Instead, Actual Field Length = '%v'\n",
			ePrefix.String(),
			expectedFieldLen,
			actualFieldLen)

		return
	}

	_,
		err =
		txtFieldSpacerOne.CopyOutITextField(
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerOne."+
			"CopyOutITextField()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	_,
		err =
		txtFieldSpacerTwo.CopyOutITextField(
			ePrefix.XCtx("Empty txtFieldSpacerTwo"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldSpacerTwo.CopyOutITextField()\n"+
			"because field length = 0\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyOutPtr_000100()",
		"")

	expectedFieldLen := 14

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtFieldSpacerTwo *TextFieldSpecSpacer

	txtFieldSpacerTwo,
		err = txtFieldSpacerOne.CopyOutPtr(
		ePrefix.XCtx(
			"txtFieldSpacerOne->" +
				"txtFieldSpacerTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldSpacerOne.Equal(txtFieldSpacerTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldSpacerOne should be equal to txtFieldSpacerTwo\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	txtFieldSpacerTwo.fieldLen = -99

	_,
		err = txtFieldSpacerTwo.CopyOutPtr(
		ePrefix.XCtx(
			"txtFieldSpacerTwo is INVALID!"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldSpacerTwo.CopyOutPtr()\n"+
			"because txtFieldSpacerTwo.fieldLen = -99\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	txtFieldSpacerThree := TextFieldSpecSpacer{}

	_,
		err = txtFieldSpacerThree.CopyOutPtr(
		ePrefix.XCtx(
			"txtFieldSpacerThree is INVALID!"))

	if err == nil {

		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldSpacerTwo.CopyOutPtr()\n"+
			"because txtFieldSpacerThree.fieldLen = 0\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return

	}

	_,
		err = txtFieldSpacerOne.CopyOutPtr(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerOne."+
			"CopyOutPtr()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_empty_000100()",
		"")

	expectedFieldLen := 14

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSpacerNanobot := textFieldSpecSpacerNanobot{}

	txtSpacerNanobot.empty(nil)

	txtSpacerNanobot2 := textFieldSpecSpacerNanobot{}

	txtSpacerNanobot2.empty(&txtFieldSpacerOne)

	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	if !txtFieldSpacerTwo.Equal(&txtFieldSpacerOne) {

		t.Errorf("%v Test #1\n"+
			"Error: txtFieldSpacerOne should be equal to txtFieldSpacerTwo\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_Empty_000100()",
		"")

	expectedFieldLen := 14

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	txtFieldSpacerOne.Empty()

	if !txtFieldSpacerTwo.Equal(&txtFieldSpacerOne) {

		t.Errorf("%v Test #1\n"+
			"Error: txtFieldSpacerOne should be equal to txtFieldSpacerTwo\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	txtFieldSpacerThree := TextFieldSpecSpacer{}

	txtFieldSpacerThree.Empty()

	if !txtFieldSpacerTwo.Equal(&txtFieldSpacerThree) {

		t.Errorf("%v Test #1\n"+
			"Error: txtFieldSpacerTwo should be equal to txtFieldSpacerThree\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_Empty_000100()",
		"")

	expectedFieldLen := 14

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtSpacerNanobot := textFieldSpecSpacerNanobot{}

	areEqual :=
		txtSpacerNanobot.equal(
			nil,
			&txtFieldSpacerOne)

	if areEqual == true {
		t.Errorf("%v Test #1\n"+
			"Error: areEqual should be 'false'"+
			"because 'txtFieldSpacer' is 'nil'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	txtSpacerNanobot2 := textFieldSpecSpacerNanobot{}

	areEqual =
		txtSpacerNanobot2.equal(
			&txtFieldSpacerOne,
			nil)

	if areEqual == true {
		t.Errorf("%v Test #1\n"+
			"Error: areEqual should be 'false'"+
			"because 'txtFieldSpacer' is 'nil'\n"+
			"HOWEVER, THE RETURN VALUE IS 'true'!\n",
			ePrefix.String())

		return
	}

	expectedFieldLen = 37
	var txtFieldSpacerThree *TextFieldSpecSpacer

	txtFieldSpacerThree,
		err = TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual =
		txtSpacerNanobot2.equal(
			&txtFieldSpacerOne,
			txtFieldSpacerThree)

	if areEqual == true {
		t.Errorf("%v Test #2\n"+
			"Error: areEqual should be 'false'"+
			"because 'txtFieldSpacerOne' and 'txtFieldSpacerThree'\n"+
			"are not equal\n"+
			"HOWEVER, THE RETURN VALUE IS 'true'!\n",
			ePrefix.String())

		return
	}

	var txtFieldSpacerFour *TextFieldSpecSpacer

	txtFieldSpacerFour,
		err = txtFieldSpacerOne.CopyOutPtr(
		ePrefix.XCtx(
			"txtFieldSpacerOne->txtFieldSpacerFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual =
		txtSpacerNanobot2.equal(
			&txtFieldSpacerOne,
			txtFieldSpacerFour)

	if areEqual == false {
		t.Errorf("%v Test #3\n"+
			"Error: areEqual should be 'true'\n"+
			"because 'txtFieldSpacerOne' and 'txtFieldSpacerFour'\n"+
			"are equal. HOWEVER, THE RETURN VALUE IS 'false'!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_Equal_000100()",
		"")

	var txtFieldSpacerOne, txtFieldSpacerTwo TextFieldSpecSpacer
	var err error

	expectedFieldLen := 4

	txtFieldSpacerOne,
		err = TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if txtFieldSpacerTwo.Equal(&txtFieldSpacerOne) {
		t.Errorf("%v\n"+
			"Error: 'txtFieldSpacerOne' should NOT be equal to"+
			"'txtFieldSpacerTwo' HOWEVER, THEY ARE SHOWING AS EQUAL!\n",
			ePrefix.String())

		return
	}

	txtFieldSpacerTwo.fieldLen = expectedFieldLen

	if !txtFieldSpacerTwo.Equal(&txtFieldSpacerOne) {
		t.Errorf("%v\n"+
			"Error: 'txtFieldSpacerOne' should be equal to\n"+
			"'txtFieldSpacerTwo' HOWEVER, THEY ARE SHOWING AS NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_EqualITextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_Equal_000100()",
		"")

	var txtFieldSpacerOne TextFieldSpecSpacer
	var err error

	expectedFieldLen := 4

	txtFieldSpacerOne,
		err = TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	areEqual := txtFieldSpacerTwo.EqualITextField(
		nil)

	if areEqual == true {
		t.Errorf("%v Test #1\n"+
			"Error: areEqual should be 'false'\n"+
			"because 'txtFieldSpacerTwo' is empty!\n"+
			"are equal. HOWEVER, THE RETURN VALUE IS 'true'!\n",
			ePrefix.String())

		return
	}

	txtFiller := TextFieldSpecFiller{}

	areEqual = txtFieldSpacerOne.EqualITextField(
		&txtFiller)

	if areEqual == true {
		t.Errorf("%v Test #2\n"+
			"Error: areEqual should be 'false'\n"+
			"because 'txtFiller' is of type 'TextFieldSpecFiller'!\n"+
			"are equal. HOWEVER, THE RETURN VALUE IS 'true'!\n",
			ePrefix.String())

		return
	}

	txtFieldSpacerTwo,
		err = TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	areEqual = txtFieldSpacerOne.EqualITextField(
		&txtFieldSpacerTwo)

	if areEqual == false {
		t.Errorf("%v Test #2\n"+
			"Error: areEqual should be 'true'\n"+
			"because 'txtFieldSpacerOne' and 'txtFieldSpacerTwo'\n"+
			"are equal. HOWEVER, THE RETURN VALUE IS 'false'!\n",
			ePrefix.String())

		return
	}

	txtFieldSpacerTwo.fieldLen = 42

	areEqual = txtFieldSpacerOne.EqualITextField(
		&txtFieldSpacerTwo)

	if areEqual == true {
		t.Errorf("%v Test #3\n"+
			"Error: areEqual should be 'false'\n"+
			"because 'txtFieldSpacerOne' and 'txtFieldSpacerTwo'\n"+
			"are NOT equal. HOWEVER, THE RETURN VALUE IS 'true'!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_GetFormattedStrLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_GetFormattedStrLength_000100()",
		"")

	fieldLen := 4

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		fieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualFormattedText string

	actualFormattedText,
		err =
		txtFieldSpacerOne.GetFormattedText(
			ePrefix.XCtx(
				"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStrLen := len(actualFormattedText)

	actualStrLen := txtFieldSpacerOne.
		GetFormattedStrLength()

	if expectedStrLen !=
		actualStrLen {
		t.Errorf("%v - ERROR\n"+
			"Expected Formatted String Length = '%v'\n"+
			"Instead, Formatted String Length = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedStrLen,
			actualStrLen)

		return
	}

	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	actualStrLen =
		txtFieldSpacerTwo.GetFormattedStrLength()

	if actualStrLen != -1 {
		t.Errorf("%v - ERROR\n"+
			"'txtFieldSpacerTwo' is INVALID!\n"+
			"Expected Formatted String Length = '-1'\n"+
			"Instead, Formatted String Length = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualStrLen)

		return
	}

	return
}

func TestTextFieldSpecSpacer_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_getFormattedText_000100()",
		"")

	expectedFieldLen := 4

	expectedStr := strings.Repeat(" ", expectedFieldLen)

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldSpacerNanobot := textFieldSpecSpacerNanobot{}

	_,
		err =
		txtFieldSpacerNanobot.getFormattedText(
			nil,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldSpacerNanobot.getFormattedText()\n"+
			"because input parameter 'txtFieldSpacer' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtx("txtFieldSpacer = nil"))

		return
	}

	txtFieldSpacerNanobot2 := textFieldSpecSpacerNanobot{}

	var actualStr string

	actualStr,
		err =
		txtFieldSpacerNanobot2.getFormattedText(
			&txtFieldSpacerOne,
			ePrefix.XCtx(
				"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("%v\n"+
			"Error: Expected formtted string = '%v'\n"+
			"Instead,       formatted string = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedStr,
			actualStr)

		return
	}

	var txtFieldSpacerTwo TextFieldSpecSpacer

	txtFieldSpacerTwo,
		err = TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldSpacerTwo.fieldLen = -99

	_,
		err =
		txtFieldSpacerNanobot.getFormattedText(
			&txtFieldSpacerTwo,
			ePrefix.XCtx("txtFieldSpacerTwo.fieldLen = -99"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldSpacerNanobot.getFormattedText()\n"+
			"because input parameter txtFieldSpacer.fieldLen is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_GetFormattedText_000100()",
		"")

	expectedFieldLen := 14

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := strings.Repeat(" ", expectedFieldLen)

	var actualStr string

	actualStr,
		err =
		txtFieldSpacerOne.GetFormattedText(
			ePrefix.XCtx(
				"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("%v\n"+
			"Error: Expected formtted string = '%v'\n"+
			"Instead,       formatted string = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedStr,
			actualStr)

		return
	}

	_,
		err =
		txtFieldSpacerOne.GetFormattedText(
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerOne."+
			"GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	txtFieldSpacerTwo.fieldLen = -99

	actualStr,
		err =
		txtFieldSpacerTwo.GetFormattedText(
			ePrefix.XCtx(
				"txtFieldSpacerTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected an error return from txtFieldSpacerTwo.GetFormattedText()\n"+
			"because txtFieldSpacerTwo.fieldLen = -99 and is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			err.Error())
		return
	}

	txtFieldSpacerThree := TextFieldSpecSpacer{}
	txtFieldSpacerThree.fieldLen = 20
	expectedStr = strings.Repeat(" ", 20)

	actualStr,
		err = txtFieldSpacerThree.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldSpacerThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedStr != actualStr {
		t.Errorf("%v Test #2\n"+
			"Error: Expected formtted string = '%v'\n"+
			"Instead,       formatted string = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedStr,
			actualStr)

		return
	}

}

func TestTextFieldSpecSpacer_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyIn_000100()",
		"")

	expectedFieldLen := 4

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix)

	if err != nil {
		t.Errorf("txtFieldSpacerOne - Error\n"+
			"%v\n",
			err.Error())
		return
	}

	err = txtFieldSpacerOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldSpacerOne.IsValidInstanceError(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerOne."+
			"IsValidInstanceError()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_NewPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyIn_000200()",
		"")

	txtFieldSpacerOne := TextFieldSpecSpacer{}.NewPtr()

	if txtFieldSpacerOne.textLineReader != nil {
		t.Errorf("%v - ERROR\n"+
			"TextFieldSpecSpacer{}.NewPtrBlankLines()\n"+
			"'txtFieldSpacerOne.textLineReader' != 'nil' !\n",
			ePrefix.String())

	}

	return
}

func TestTextFieldSpecSpacer_NewPtrSpacer_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_NewPtrSpacer_000100()",
		"")

	expectedFieldLen := 1000001

	_,
		err := TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx("Invalid Field Length == 1000001"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecSpacer{}.NewPtrSpacer()\n"+
			"because Field Length = '1000001' and is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
	}

	expectedFieldLen = 4
	var txtFieldSpacerOne *TextFieldSpecSpacer

	txtFieldSpacerOne,
		err = TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("txtFieldSpacerOne - Error\n"+
			"%v\n",
			err.Error())
		return
	}

	err = txtFieldSpacerOne.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextFieldSpecSpacer{}."+
			"NewPtrSpacer()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_NewSpacer_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_NewSpacer_000100()",
		"")

	expectedFieldLen := -2

	_,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx("Field Length invalid!"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecSpacer{}.NewSpacer()\n"+
			"because Field Length = '-2' and is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
	}

	expectedFieldLen = 4

	_,
		err = TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextFieldSpecSpacer{}."+
			"NewPtrSpacer()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_Read_000100()",
		"")

	p := make([]byte, 500)

	var n, readBytesCnt int

	expectedFieldLen := 14

	txtFieldSpacerZero,
		err := TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldSpacerZero.fieldLen = -999
	txtFieldSpacerZero.textLineReader = nil

	_,
		err = txtFieldSpacerZero.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerZero{}."+
			"Read()\n"+
			"because 'txtFieldSpacerZero' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtFieldSpacerOne *TextFieldSpecSpacer

	txtFieldSpacerOne,
		err = TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := strings.Repeat(" ", expectedFieldLen)

	var actualStr string

	for {

		n,
			err = txtFieldSpacerOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtFieldSpacerOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if txtFieldSpacerOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"fillerTxtFieldOne.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != expectedFieldLen {
		t.Errorf("%v\n"+
			"Byte Length Error: fillerTxtFieldOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedFieldLen,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)
	}

	return
}

func TestTextFieldSpecSpacer_ReadInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_ReadInitialize_000100()",
		"")

	txtFieldSpacerZero := TextFieldSpecSpacer{}

	txtFieldSpacerZero.ReaderInitialize()

	expectedFieldLen := 14

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := strings.Repeat(" ", expectedFieldLen)

	p := make([]byte, 5)

	var n int

	n,
		err = txtFieldSpacerOne.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldSpacerOne.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: txtFieldSpacerOne.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.XCtxEmpty().String(),
			n)

		return
	}

	p = make([]byte, 100)

	txtFieldSpacerOne.ReaderInitialize()

	var readBytesCnt int
	var actualStr string
	n = 0

	for {

		n,
			err = txtFieldSpacerOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtFieldSpacerOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if txtFieldSpacerOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"fillerTxtFieldOne.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != expectedFieldLen {
		t.Errorf("%v\n"+
			"Byte Length Error: txtFieldSpacerOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read Length = '%v'\n"+
			"       Actual Bytes Length = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedFieldLen,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedStr),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if txtFieldSpacerOne.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but txtFieldSpacerOne.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	p = make([]byte, 100)
	n = 0
	actualStr = ""
	readBytesCnt = 0

	for {

		n,
			err = txtFieldSpacerOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error Returned From txtFieldSpacerOne.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v Test #2\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if txtFieldSpacerOne.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but txtFieldSpacerOne.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_setTextFieldSpacer_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_setTextFieldSpacer_000100()",
		"")

	expectedFieldLen := 14

	txtFieldSpacerOne := TextFieldSpecSpacer{}

	txtFieldSpacerNanobot := textFieldSpecSpacerNanobot{}

	err := txtFieldSpacerNanobot.setTextFieldSpacer(
		&txtFieldSpacerOne,
		expectedFieldLen,
		ePrefix.XCtx(
			"->txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	err = txtFieldSpacerNanobot.setTextFieldSpacer(
		&txtFieldSpacerTwo,
		-999,
		ePrefix.XCtx(
			"fieldLen is invalid!"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from  txtFieldSpacerNanobot."+
			"setTextFieldSpacer()\n"+
			"because 'fieldLen' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtFieldSpacerNanobot.setTextFieldSpacer(
		nil,
		expectedFieldLen,
		ePrefix.XCtx(
			"txtFieldSpacer is nil!"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from  txtFieldSpacerNanobot."+
			"setTextFieldSpacer()\n"+
			"because 'txtFieldSpacer' is 'nill' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_SetFieldLen_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_SetFieldLen_000100()",
		"")

	expectedFieldLen := 9

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldSpacerOne.SetFieldLen(
		24,
		ePrefix.XCtx(
			"txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldSpacerOne.SetFieldLen(
		expectedFieldLen,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerOne."+
			"SetFieldLen()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecSpacer_String_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_String_000100()",
		"")

	expectedFieldLen := 14

	txtFieldSpacerOne,
		err := TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedStr := strings.Repeat(" ", expectedFieldLen)

	actualStr :=
		txtFieldSpacerOne.String()

	if expectedStr != actualStr {
		t.Errorf("%v\n"+
			"Error: Expected formtted string = '%v'\n"+
			"Instead,       formatted string = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedStr,
			actualStr)

		return
	}

	txtFieldSpacerTwo := TextFieldSpecSpacer{}

	txtFieldSpacerTwo.fieldLen = -99

	actualStr =
		txtFieldSpacerTwo.String()

	if !strings.Contains(actualStr, "Error") {
		t.Errorf("%v\n"+
			"Error: Expected 'actualStr' to contain the word 'Error'.\n"+
			"because txtFieldSpacerTwo.fieldLen = -99\n"+
			"HOWEVER, the returned 'actualStr' did NOT contain the "+
			"word 'Error'\n"+
			"'actualStr' = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualStr)

		return
	}

	txtFieldSpacerThree := TextFieldSpecSpacer{}
	txtFieldSpacerThree.fieldLen = 20
	expectedStr = strings.Repeat(" ", 20)

	actualStr = txtFieldSpacerThree.String()

	if expectedStr != actualStr {
		t.Errorf("%v Test #2\n"+
			"Error: Expected formtted string = '%v'\n"+
			"Instead,       formatted string = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedStr,
			actualStr)

		return
	}

	return
}

func TestTextFieldSpecSpacer_TextBuilder_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_CopyOut_000100()",
		"")

	txtFieldSpacerZero := TextFieldSpecSpacer{}

	err :=
		txtFieldSpacerZero.TextBuilder(
			nil,
			ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerZero."+
			"TextBuilder()\n"+
			"because input parameter 'sBuilder' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	expectedFieldLen := 9

	expectedFormattedText := strings.Repeat(" ", expectedFieldLen)

	var txtFieldSpacerOne TextFieldSpecSpacer

	txtFieldSpacerOne,
		err = TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldSpacerOne.IsValidInstanceError(
		ePrefix.XCtx("txtFieldSpacerOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sb := strings.Builder{}

	err =
		txtFieldSpacerOne.TextBuilder(
			&sb,
			ePrefix.XCtx(
				"txtFieldSpacerOne->sb"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldSpacerOne.TextBuilder()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFormattedText),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(sb.String()),
			true)

	if printableExpectedStr !=
		printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"txtFieldSpacerOne.TextBuilder()\n"+
			"Expected Formatted Text = '%v'\n"+
			"Instead, Formatted Text = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	sb.Reset()

	err =
		txtFieldSpacerOne.TextBuilder(
			&sb,
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldOne{}."+
			"TextBuilder()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtFieldSpacerTwo TextFieldSpecSpacer

	txtFieldSpacerTwo,
		err = TextFieldSpecSpacer{}.NewSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecSpacer{}.NewSpacer()\n"+
			"Error:\n'%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	err =
		txtFieldSpacerTwo.TextBuilder(
			nil,
			ePrefix.XCtx(
				"txtFieldSpacerTwo"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerTwo."+
			"TextBuilder()\n"+
			"because input parameter 'sBuilder' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtFieldSpacerThree *TextFieldSpecSpacer

	txtFieldSpacerThree,
		err = TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerThree"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecSpacer{}.NewPtrSpacer()\n"+
			"Error:\n'%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	sb.Reset()

	err =
		txtFieldSpacerThree.TextBuilder(
			&sb,
			TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerThree."+
			"TextBuilder()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	sb.Reset()

	txtFieldSpacerThree.fieldLen = -9999

	err =
		txtFieldSpacerThree.TextBuilder(
			&sb,
			ePrefix.XCtx(
				"txtFieldSpacerThree is invalid!"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldSpacerThree."+
			"TextBuilder()\n"+
			"because 'txtFieldSpacerThree.fieldLen' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var txtFieldSpacerFour *TextFieldSpecSpacer

	txtFieldSpacerFour,
		err = TextFieldSpecSpacer{}.NewPtrSpacer(
		expectedFieldLen,
		ePrefix.XCtx("txtFieldSpacerFour"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecSpacer{}.NewPtrSpacer()\n"+
			"Error:\n'%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	sb.Reset()

	err =
		txtFieldSpacerFour.TextBuilder(
			&sb,
			ePrefix.XCtx(
				"txtFieldSpacerFour->sb"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualFmtStr := sb.String()

	if expectedFormattedText !=
		actualFmtStr {
		t.Errorf("%v - ERROR\n"+
			"Expected String Builder Text = '%v'\n"+
			"Instead, String Builder Text = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedFormattedText,
			actualFmtStr)

		return
	}

	var formattedString string

	formattedString,
		err = txtFieldSpacerFour.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldSpacerFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if formattedString !=
		actualFmtStr {
		t.Errorf("%v - ERROR\n"+
			"txtFieldSpacerFour.GetFormattedText()"+
			"Expected String Builder Text = '%v'\n"+
			"Instead, String Builder Text = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			formattedString,
			actualFmtStr)

		return
	}

	return
}

func TestTextFieldSpecSpacer_TextFieldName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_TextFieldName_000100()",
		"")

	txtFieldSpacerOne := TextFieldSpecSpacer{}.New()

	fieldName := txtFieldSpacerOne.TextFieldName()

	if fieldName != "TextFieldSpecSpacer" {

		t.Errorf("%v\n"+
			"Expected txtFieldSpacerOne.TextFieldName()\n"+
			"to return 'TextFieldSpecSpacer'.\n"+
			"Instead, Field Name = '%v'\n",
			ePrefix.String(),
			fieldName)

	}

	return
}

func TestTextFieldSpecSpacer_TextTypeName_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecSpacer_TextTypeName_000100()",
		"")

	txtFieldSpacerOne := TextFieldSpecSpacer{}.New()

	fieldName := txtFieldSpacerOne.TextTypeName()

	if fieldName != "TextFieldSpecSpacer" {

		t.Errorf("%v\n"+
			"Expected txtFieldSpacerOne.TextTypeName()\n"+
			"to return 'TextFieldSpecSpacer'.\n"+
			"Instead, Field Name = '%v'\n",
			ePrefix.String(),
			fieldName)

	}

	return
}

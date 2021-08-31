package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
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

	actualFieldLen := txtFieldSpacerOne.GetFieldLength()

	if expectedFieldLen != actualFieldLen {

		t.Errorf("%v\n"+
			"Error: Invalid Field Length returned by\n"+
			"txtFieldSpacerOne.GetFieldLength()\n"+
			"Expected Field Length = '%v'\n"+
			"Instead, Actual Field Length = '%v'\n",
			ePrefix.String(),
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

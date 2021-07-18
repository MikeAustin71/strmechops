package strmech

import (
	"strings"
	"testing"
)

func TestTextFieldSpecLabel_CopyIn_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyIn_000100() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()
	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewConstructor(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldLabelTwo := TextFieldSpecLabel{}

	err = txtFieldLabelTwo.CopyIn(
		txtFieldLabelOne,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(&txtFieldLabelTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelTwo!\n",
			ePrefix)
		return
	}

	actualLabel := txtFieldLabelOne.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
	}

	return
}

func TestTextFieldSpecLabel_CopyIn_000200(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyIn_000200() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewConstructor(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldLabelTwo := TextFieldSpecLabel{}

	err = txtFieldLabelOne.CopyIn(
		&txtFieldLabelTwo,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelOne.CopyIn(txtFieldLabelTwo)\n"+
			"because 'txtFieldLabelTwo' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!!\n",
			ePrefix)
		return
	}

}

func TestTextFieldSpecLabel_CopyOut_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyOut_000100() "

	labelRunes := []rune("12345")
	fieldLen := 14
	txtJustify := TxtJustify.Left()
	expectedTextLabel :=
		string(labelRunes) +
			strings.Repeat(" ", 9)

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewConstructorRunes(
		labelRunes,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtFieldLabelTwo TextFieldSpecLabel

	txtFieldLabelTwo,
		err = txtFieldLabelOne.CopyOut(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(&txtFieldLabelTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelTwo!\n",
			ePrefix)
		return
	}

	actualLabel := txtFieldLabelOne.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
	}

	return
}

func TestTextFieldSpecLabel_CopyOut_000200(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyOut_000200() "

	labelText := "12345"
	fieldLen := 14
	txtJustify := TxtJustify.Right()
	expectedTextLabel :=
		strings.Repeat(" ", 9) +
			labelText

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewTextLabel(
		labelText,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("txtFieldLabelOne\n"+
			"%v\n",
			err.Error())
		return
	}

	var txtFieldLabelTwo TextFieldSpecLabel

	txtFieldLabelTwo,
		err = txtFieldLabelOne.CopyOut(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("txtFieldLabelTwo\n"+
			"%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(&txtFieldLabelTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelTwo!\n",
			ePrefix)
		return
	}

	var txtFieldLabelThree TextFieldSpecLabel

	txtFieldLabelThree = txtFieldLabelTwo

	err = txtFieldLabelThree.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("Error: txtFieldLabelThree\n"+
			"%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(&txtFieldLabelThree) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelThree!\n",
			ePrefix)
		return
	}

	actualLabel := txtFieldLabelThree.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
	}

	return
}

func TestTextFieldSpecLabel_CopyOut_000300(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyOut_000300() "

	txtFieldLabelOne := TextFieldSpecLabel{}

	_,
		err := txtFieldLabelOne.CopyOut(
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelOne.CopyOut()\n"+
			"because txtFieldLabelOne is an invalid Text Field Label.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

}

func TestTextFieldSpecLabel_Empty_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyIn_000100() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewConstructor(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: Tested validity of 'txtFieldLabelOne'.\n"+
			"txtFieldLabelOne.IsValidInstance()\n"+
			"'txtFieldLabelOne' is INVALID!\n",
			ePrefix)

		return
	}

	txtFieldLabelOne.Empty()

	textLabel := txtFieldLabelOne.GetFormattedText()

	if !strings.Contains(textLabel, "Error") {
		t.Errorf("%v\n"+
			"Error: Expected an error string from txtFieldLabelOne.GetFormattedText()\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)
		return
	}

	if txtFieldLabelOne.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldLabelOne.IsValidInstance()\n"+
			"to return 'false' INSTEAD, IT RETURNED 'true'!!\n",
			ePrefix)
	}
}

func TestTextFieldSpecLabel_NewTextLabel_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_NewTextLabel_000100() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtFieldLabelTwo *TextFieldSpecLabel

	txtFieldLabelTwo,
		err = TextFieldSpecLabel{}.NewConstructorRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(txtFieldLabelTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelTwo!\n",
			ePrefix)
		return
	}

	actualLabel := txtFieldLabelOne.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
	}

	return
}

func TestTextFieldSpecLabel_NewTextLabel_000200(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_NewTextLabel_000200() "

	label := "12345"
	fieldLen := -1
	txtJustify := TxtJustify.None()

	expectedTextLabel := label

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtFieldLabelTwo *TextFieldSpecLabel

	txtFieldLabelTwo,
		err = TextFieldSpecLabel{}.NewConstructorRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(txtFieldLabelTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelTwo!\n",
			ePrefix)
		return
	}

	actualLabel := txtFieldLabelOne.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
	}

	return
}

func TestTextFieldSpecLabel_NewTextLabel_000300(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_NewTextLabel_000300() "

	label := "12345"
	fieldLen := -99
	txtJustify := TxtJustify.None()

	_,
		err := TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecLabel{}.NewTextLabel()\n"+
			"because 'fieldLen' is less than one (1).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecLabel_NewTextLabel_000400(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_NewTextLabel_000400() "

	label := "12345"
	fieldLen := 15
	txtJustify := TxtJustify.None()

	_,
		err := TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecLabel{}.NewTextLabel()\n"+
			"because 'txtJustify' ==  TxtJustify.None().\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecLabel_SetTextLabel_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_SetTextLabel_000100() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()
	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewConstructor(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldLabelTwo := TextFieldSpecLabel{}

	err = txtFieldLabelTwo.SetTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(&txtFieldLabelTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelTwo!\n"+
			"txtFieldLabelTwo.SetTextLabel() did NOT produce an identical copy!\n",
			ePrefix)
		return
	}

	actualLabel := txtFieldLabelTwo.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
	}

	return
}

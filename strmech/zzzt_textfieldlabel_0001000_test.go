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

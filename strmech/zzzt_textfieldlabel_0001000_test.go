package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
)

func TestTextFieldSpecLabel_CopyIn_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyIn_000100()"

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()
	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
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
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
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
		err := TextFieldSpecLabel{}.NewPtrTextLabelRunes(
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

func TestTextFieldSpecLabel_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_CopyOutPtr_000100()",
		"")

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

	var txtFieldLabelTwo *TextFieldSpecLabel

	txtFieldLabelTwo,
		err = txtFieldLabelOne.CopyOutPtr(
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

	if !txtFieldLabelOne.Equal(txtFieldLabelTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelTwo!\n",
			ePrefix)
		return
	}

	var txtFieldLabelThree TextFieldSpecLabel

	txtFieldLabelThree = *txtFieldLabelTwo

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

		return
	}

	if txtFieldLabelOne.GetFieldLength() !=
		txtFieldLabelThree.GetFieldLength() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelThree\n"+
			"field lengths are NOT equal!\n"+
			" txtFieldLabelOne.GetFieldLength()   == '%v'\n"+
			" txtFieldLabelThree.GetFieldLength() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetFieldLength(),
			txtFieldLabelThree.GetFieldLength())
		return
	}

	if txtFieldLabelOne.GetTextJustification() !=
		txtFieldLabelThree.GetTextJustification() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelThree\n"+
			"text justification values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextJustification()   == '%v'\n"+
			" txtFieldLabelThree.GetTextJustification() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextJustification().String(),
			txtFieldLabelThree.GetTextJustification().String())
		return
	}

	if txtFieldLabelOne.GetTextLabel() !=
		txtFieldLabelThree.GetTextLabel() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelThree\n"+
			"text label values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextLabel()   == '%v'\n"+
			" txtFieldLabelThree.GetTextLabel() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextLabel(),
			txtFieldLabelThree.GetTextLabel())
	}

	return
}

func TestTextFieldSpecLabel_Empty_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_CopyIn_000100() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
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

func TestTextFieldSpecLabel_NewPtrTextLabelRunes_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_NewPtrTextLabelRunes_000100() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		[]rune(label),
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
		t.Errorf("%v - txtFieldLabelOne\n",
			err.Error())
		return
	}

	txtFieldLabelTwo := TextFieldSpecLabel{}

	err = txtFieldLabelTwo.SetTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	err = txtFieldLabelTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v - txtFieldLabelTwo\n",
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
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel = txtFieldLabelTwo.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelTwo\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	if txtFieldLabelOne.GetFieldLength() !=
		txtFieldLabelTwo.GetFieldLength() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"field lengths are NOT equal!\n"+
			" txtFieldLabelOne.GetFieldLength() == '%v'\n"+
			" txtFieldLabelTwo.GetFieldLength() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetFieldLength(),
			txtFieldLabelTwo.GetFieldLength())
		return
	}

	if txtFieldLabelOne.GetTextJustification() !=
		txtFieldLabelTwo.GetTextJustification() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text justification values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextJustification() == '%v'\n"+
			" txtFieldLabelTwo.GetTextJustification() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextJustification().String(),
			txtFieldLabelTwo.GetTextJustification().String())
		return
	}

	if txtFieldLabelOne.GetTextLabel() !=
		txtFieldLabelTwo.GetTextLabel() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text label values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextLabel() == '%v'\n"+
			" txtFieldLabelTwo.GetTextLabel() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextLabel(),
			txtFieldLabelTwo.GetTextLabel())
	}

	return
}

func TestTextFieldSpecLabel_NewPtrTextLabelRunes_000200(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_NewPtrTextLabelRunes_000200() "

	label := "12345"
	fieldLen := 6
	txtJustify := TxtJustify.Center()

	expectedTextLabel :=
		label + " "

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		[]rune(label),
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
		t.Errorf("%v - txtFieldLabelOne\n",
			err.Error())
		return
	}

	txtFieldLabelTwo := TextFieldSpecLabel{}

	err = txtFieldLabelTwo.SetTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v - txtFieldLabelTwo\n",
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
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel = txtFieldLabelTwo.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelTwo\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	if txtFieldLabelOne.GetFieldLength() !=
		txtFieldLabelTwo.GetFieldLength() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"field lengths are NOT equal!\n"+
			" txtFieldLabelOne.GetFieldLength() == '%v'\n"+
			" txtFieldLabelTwo.GetFieldLength() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetFieldLength(),
			txtFieldLabelTwo.GetFieldLength())
		return
	}

	if txtFieldLabelOne.GetTextJustification() !=
		txtFieldLabelTwo.GetTextJustification() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text justification values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextJustification() == '%v'\n"+
			" txtFieldLabelTwo.GetTextJustification() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextJustification().String(),
			txtFieldLabelTwo.GetTextJustification().String())
		return
	}

	if txtFieldLabelOne.GetTextLabel() !=
		txtFieldLabelTwo.GetTextLabel() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text label values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextLabel() == '%v'\n"+
			" txtFieldLabelTwo.GetTextLabel() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextLabel(),
			txtFieldLabelTwo.GetTextLabel())
	}

	return
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
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
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
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
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
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
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

func TestTextFieldSpecLabel_SetTextLabel_000200(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_SetTextLabel_000200() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()
	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
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

	txtFieldLabelThree := TextFieldSpecLabel{}

	err = txtFieldLabelThree.SetTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(&txtFieldLabelThree) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelThree!\n"+
			"txtFieldLabelTwo.SetTextLabelRunes() did NOT produce an identical copy!\n",
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

func TestTextFieldSpecLabel_SetTextLabel_000300(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_SetTextLabel_000300() "

	label := "12345"
	fieldLen := -1
	txtJustify := TxtJustify.None()
	expectedTextLabel := label

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
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

	txtFieldLabelThree := TextFieldSpecLabel{}

	err = txtFieldLabelThree.SetTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldLabelOne.Equal(&txtFieldLabelThree) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelThree!\n"+
			"txtFieldLabelTwo.SetTextLabelRunes() did NOT produce an identical copy!\n",
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

func TestTextFieldSpecLabel_SetTextLabel_000400(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_SetTextLabel_000400() "

	label := "12345"
	fieldLen := -99
	txtJustify := TxtJustify.None()

	txtFieldLabelOne := TextFieldSpecLabel{}

	err := txtFieldLabelOne.SetTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelOne.SetTextLabel()\n"+
			"because 'fieldLen' == '-99' and is therefore INVALID!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecLabel_SetTextLabel_000500(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_SetTextLabel_000500() "

	label := "12345"
	fieldLen := 99
	txtJustify := TxtJustify.None()

	txtFieldLabelOne := TextFieldSpecLabel{}

	err := txtFieldLabelOne.SetTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelOne.SetTextLabel()\n"+
			"because 'txtJustify' == TxtJustify.None() and is therefore INVALID!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecLabel_SetTextLabelRunes_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TestTextFieldSpecLabel_SetTextLabelRunes_000100() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	txtFieldLabelOne := TextFieldSpecLabel{}.NewEmpty()

	err := txtFieldLabelOne.SetTextLabelRunes(
		[]rune(label),
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
		t.Errorf("%v - txtFieldLabelOne\n",
			err.Error())
		return
	}

	var txtFieldLabelTwo *TextFieldSpecLabel

	txtFieldLabelTwo,
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v - txtFieldLabelTwo\n",
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
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel = txtFieldLabelTwo.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelTwo\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	if txtFieldLabelOne.GetFieldLength() !=
		txtFieldLabelTwo.GetFieldLength() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"field lengths are NOT equal!\n"+
			" txtFieldLabelOne.GetFieldLength() == '%v'\n"+
			" txtFieldLabelTwo.GetFieldLength() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetFieldLength(),
			txtFieldLabelTwo.GetFieldLength())
		return
	}

	if txtFieldLabelOne.GetTextJustification() !=
		txtFieldLabelTwo.GetTextJustification() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text justification values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextJustification() == '%v'\n"+
			" txtFieldLabelTwo.GetTextJustification() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextJustification().String(),
			txtFieldLabelTwo.GetTextJustification().String())
		return
	}

	if txtFieldLabelOne.GetTextLabel() !=
		txtFieldLabelTwo.GetTextLabel() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text label values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextLabel() == '%v'\n"+
			" txtFieldLabelTwo.GetTextLabel() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextLabel(),
			txtFieldLabelTwo.GetTextLabel())
	}

	return
}

func TestTextFieldSpecLabel_SetTextLabelRunes_000200(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TestTextFieldSpecLabel_SetTextLabelRunes_000200() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Left()

	expectedTextLabel :=
		label +
			strings.Repeat(" ", 8)

	txtFieldLabelOne := TextFieldSpecLabel{}.NewEmpty()

	err := txtFieldLabelOne.SetTextLabelRunes(
		[]rune(label),
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
		t.Errorf("%v - txtFieldLabelOne\n",
			err.Error())
		return
	}

	var txtFieldLabelTwo *TextFieldSpecLabel

	txtFieldLabelTwo,
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v - txtFieldLabelTwo\n",
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
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel = txtFieldLabelTwo.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelTwo\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	if txtFieldLabelOne.GetFieldLength() !=
		txtFieldLabelTwo.GetFieldLength() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"field lengths are NOT equal!\n"+
			" txtFieldLabelOne.GetFieldLength() == '%v'\n"+
			" txtFieldLabelTwo.GetFieldLength() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetFieldLength(),
			txtFieldLabelTwo.GetFieldLength())
		return
	}

	if txtFieldLabelOne.GetTextJustification() !=
		txtFieldLabelTwo.GetTextJustification() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text justification values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextJustification() == '%v'\n"+
			" txtFieldLabelTwo.GetTextJustification() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextJustification().String(),
			txtFieldLabelTwo.GetTextJustification().String())
		return
	}

	if txtFieldLabelOne.GetTextLabel() !=
		txtFieldLabelTwo.GetTextLabel() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text label values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextLabel() == '%v'\n"+
			" txtFieldLabelTwo.GetTextLabel() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextLabel(),
			txtFieldLabelTwo.GetTextLabel())
	}

	return
}

func TestTextFieldSpecLabel_SetTextLabelRunes_000300(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TestTextFieldSpecLabel_SetTextLabelRunes_000300() "

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Right()

	expectedTextLabel :=
		strings.Repeat(" ", 8) +
			label

	txtFieldLabelOne := TextFieldSpecLabel{}.NewEmpty()

	err := txtFieldLabelOne.SetTextLabelRunes(
		[]rune(label),
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
		t.Errorf("%v - txtFieldLabelOne\n",
			err.Error())
		return
	}

	var txtFieldLabelTwo *TextFieldSpecLabel

	txtFieldLabelTwo,
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v - txtFieldLabelTwo\n",
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
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel = txtFieldLabelTwo.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelTwo\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	if txtFieldLabelOne.GetFieldLength() !=
		txtFieldLabelTwo.GetFieldLength() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"field lengths are NOT equal!\n"+
			" txtFieldLabelOne.GetFieldLength() == '%v'\n"+
			" txtFieldLabelTwo.GetFieldLength() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetFieldLength(),
			txtFieldLabelTwo.GetFieldLength())
		return
	}

	if txtFieldLabelOne.GetTextJustification() !=
		txtFieldLabelTwo.GetTextJustification() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text justification values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextJustification() == '%v'\n"+
			" txtFieldLabelTwo.GetTextJustification() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextJustification().String(),
			txtFieldLabelTwo.GetTextJustification().String())
		return
	}

	if txtFieldLabelOne.GetTextLabel() !=
		txtFieldLabelTwo.GetTextLabel() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text label values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextLabel() == '%v'\n"+
			" txtFieldLabelTwo.GetTextLabel() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextLabel(),
			txtFieldLabelTwo.GetTextLabel())
	}

	return
}

func TestTextFieldSpecLabel_SetTextLabelRunes_000400(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TestTextFieldSpecLabel_SetTextLabelRunes_000400() "

	label := "12345"
	fieldLen := 1000001
	txtJustify := TxtJustify.Right()

	txtFieldLabelOne := TextFieldSpecLabel{}.NewEmpty()

	err := txtFieldLabelOne.SetTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelOne.SetTextLabelRunes()\n"+
			"because field length is '1,000,001'!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix)

	if err == nil {
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Expected an error return from txtFieldLabelOne."+
			"IsValidInstanceError()\n"+
			"because field length is '1,000,001'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecLabel_SetTextLabelRunes_000500(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TestTextFieldSpecLabel_SetTextLabelRunes_000500() "

	label := "12345"
	fieldLen := 45
	txtJustify := TxtJustify.None()

	txtFieldLabelOne := TextFieldSpecLabel{}.NewEmpty()

	err := txtFieldLabelOne.SetTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelOne.SetTextLabelRunes()\n"+
			"because 'txtJustify' == TxtJustify.None()!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix)

	if err == nil {
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Expected an error return from txtFieldLabelOne."+
			"IsValidInstanceError()\n"+
			"because 'txtJustify' == TxtJustify.None()\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecLabel_SetTextLabelRunes_000600(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TestTextFieldSpecLabel_SetTextLabelRunes_000600() "

	label := "12345"
	fieldLen := 5
	txtJustify := TxtJustify.None()

	expectedTextLabel := label

	txtFieldLabelOne := TextFieldSpecLabel{}.NewEmpty()

	err := txtFieldLabelOne.SetTextLabelRunes(
		[]rune(label),
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
		t.Errorf("%v - txtFieldLabelOne\n",
			err.Error())
		return
	}

	var txtFieldLabelTwo *TextFieldSpecLabel

	txtFieldLabelTwo,
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v - txtFieldLabelTwo\n",
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
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel = txtFieldLabelTwo.GetFormattedText()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelTwo\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	if txtFieldLabelOne.GetFieldLength() !=
		txtFieldLabelTwo.GetFieldLength() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"field lengths are NOT equal!\n"+
			" txtFieldLabelOne.GetFieldLength() == '%v'\n"+
			" txtFieldLabelTwo.GetFieldLength() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetFieldLength(),
			txtFieldLabelTwo.GetFieldLength())
		return
	}

	if txtFieldLabelOne.GetTextJustification() !=
		txtFieldLabelTwo.GetTextJustification() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text justification values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextJustification() == '%v'\n"+
			" txtFieldLabelTwo.GetTextJustification() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextJustification().String(),
			txtFieldLabelTwo.GetTextJustification().String())
		return
	}

	if txtFieldLabelOne.GetTextLabel() !=
		txtFieldLabelTwo.GetTextLabel() {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne and txtFieldLabelTwo\n"+
			"text label values are NOT equal!\n"+
			" txtFieldLabelOne.GetTextLabel() == '%v'\n"+
			" txtFieldLabelTwo.GetTextLabel() == '%v'\n",
			ePrefix,
			txtFieldLabelOne.GetTextLabel(),
			txtFieldLabelTwo.GetTextLabel())
	}

	return
}

func TestTextFieldSpecLabel_TextFieldName_0001000(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TextFieldName_0001000()"

	expectedTxtFieldName := "TextFieldSpecLabel"

	actualTxtFieldName := TextFieldSpecLabel{}.TextFieldName()

	if expectedTxtFieldName != actualTxtFieldName {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected Text Field Name = '%v'\n"+
			"Instead, Actual Text Field Name = '%v'\n",
			ePrefix,
			expectedTxtFieldName,
			actualTxtFieldName)

		return
	}

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Left()

	txtFieldLabelOne := TextFieldSpecLabel{}.NewEmpty()

	err := txtFieldLabelOne.SetTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualTxtFieldName = txtFieldLabelOne.TextFieldName()

	if expectedTxtFieldName != actualTxtFieldName {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected Text Field Name = '%v'\n"+
			"Instead, Actual Text Field Name = '%v'\n",
			ePrefix,
			expectedTxtFieldName,
			actualTxtFieldName)

		return
	}

	txtFieldLabelTwo := TextFieldSpecLabel{}

	actualTxtFieldName = txtFieldLabelTwo.TextFieldName()

	if expectedTxtFieldName != actualTxtFieldName {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected Text Field Name = '%v'\n"+
			"Instead, Actual Text Field Name = '%v'\n",
			ePrefix,
			expectedTxtFieldName,
			actualTxtFieldName)
	}

	return
}

func TestTextFieldSpecLabel_TextTypeName_0001000(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TextTypeName_0001000()"

	expectedTxtTypeName := "TextFieldSpecLabel"

	actualTxtTypeName := TextFieldSpecLabel{}.TextTypeName()

	if expectedTxtTypeName != actualTxtTypeName {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected Text Type Name = '%v'\n"+
			"Instead, Actual Text Type Name = '%v'\n",
			ePrefix,
			expectedTxtTypeName,
			actualTxtTypeName)

		return
	}

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Left()

	txtFieldLabelOne := TextFieldSpecLabel{}.NewEmpty()

	err := txtFieldLabelOne.SetTextLabelRunes(
		[]rune(label),
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualTxtTypeName = txtFieldLabelOne.TextFieldName()

	if expectedTxtTypeName != actualTxtTypeName {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected Text Type Name = '%v'\n"+
			"Instead, Actual Text Type Name = '%v'\n",
			ePrefix,
			expectedTxtTypeName,
			actualTxtTypeName)

		return
	}

	txtFieldLabelTwo := TextFieldSpecLabel{}

	actualTxtTypeName = txtFieldLabelTwo.TextFieldName()

	if expectedTxtTypeName != actualTxtTypeName {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected Text Type Name = '%v'\n"+
			"Instead, Actual Text Type Name = '%v'\n",
			ePrefix,
			expectedTxtTypeName,
			actualTxtTypeName)
	}

	return
}

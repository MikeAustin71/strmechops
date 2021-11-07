package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
	"time"
)

func TestTextFieldSpecLabel_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_CopyIn_000100()",
		"")

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
		ePrefix.XCtx(
			"txtFieldLabelOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelOne.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldLabelOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtFieldLabelTwo := TextFieldSpecLabel{}

	err = txtFieldLabelTwo.CopyIn(
		txtFieldLabelOne,
		ePrefix.XCtx(
			"txtFieldLabelTwo"))

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
	}

	txtFieldLabelThree := TextFieldSpecLabel{}

	err = txtFieldLabelThree.CopyIn(
		txtFieldLabelOne,
		ePrefix.XCtx(
			"txtFieldLabelOne->txtFieldLabelThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelThree.CopyIn(
		&txtFieldLabelTwo,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldLabelThree."+
			"CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
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

func TestTextFieldSpecLabel_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_copyIn_000100()",
		"")

	txtFieldLabelMolecule := textFieldSpecLabelMolecule{}

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

	err =
		txtFieldLabelMolecule.copyIn(
			txtFieldLabelOne,
			nil,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected error return from txtFieldLabelMolecule.copyIn()\n"+
			"because 'incomingTxtFieldLabel' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtFieldLabelMolecule.copyIn(
			nil,
			txtFieldLabelOne,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected error return from txtFieldLabelMolecule.copyIn()\n"+
			"because 'targetTxtFieldLabel' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtFieldLabelOne.textLabel = nil

	txtFieldLabelTwo := TextFieldSpecLabel{}

	err =
		txtFieldLabelMolecule.copyIn(
			&txtFieldLabelTwo,
			txtFieldLabelOne,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected error return from txtFieldLabelMolecule.copyIn()\n"+
			"because 'targetTxtFieldLabel' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextFieldSpecLabel_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_CopyOut_000100()",
		"")

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
		ePrefix.XCtx(
			"txtFieldLabelOne"))

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
		ePrefix.XCtx(
			"txtFieldLabelOne->txtFieldLabelTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	ePrefix.XCtxEmpty()

	if !txtFieldLabelOne.Equal(&txtFieldLabelTwo) {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne IS NOT EQUAL to txtFieldLabelTwo!\n",
			ePrefix.String())
		return
	}

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldLabelOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	ePrefix.XCtxEmpty()

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	_,
		err = txtFieldLabelTwo.CopyOut(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldLabelThree."+
			"CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelThree.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

func TestTextFieldSpecLabel_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_copyOut_000100()",
		"")

	labelRunes := []rune("12345")
	fieldLen := 14
	txtJustify := TxtJustify.Left()

	_,
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

	txtFieldLabelMolecule := textFieldSpecLabelMolecule{}

	_,
		err =
		txtFieldLabelMolecule.copyOut(
			nil,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected error return from txtFieldLabelMolecule.copyIn()\n"+
			"because 'txtFieldLabel' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextFieldSpecLabel_CopyOutITextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_CopyOutITextField_000100()",
		"")

	txtFieldLabel := TextFieldSpecLabel{}

	_,
		err := txtFieldLabel.CopyOutITextField(
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtFieldLabel.CopyOutITextField()\n"+
			"because 'txtFieldLabel' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETUNRED!\n",
			ePrefix.String())

		return
	}

	labelText := "12345"
	fieldLen := 14
	txtJustify := TxtJustify.Right()

	var txtFieldLabelTwo TextFieldSpecLabel

	txtFieldLabelTwo,
		err = TextFieldSpecLabel{}.NewTextLabel(
		labelText,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"txtFieldLabelTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldLabelTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldLabelTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtFieldSpec ITextFieldSpecification

	txtFieldSpec,
		err = txtFieldLabelTwo.CopyOutITextField(
		ePrefix.XCtx(
			"txtFieldLabelTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var ok bool
	var txtFieldLabelThree *TextFieldSpecLabel

	txtFieldLabelThree, ok =
		txtFieldSpec.(*TextFieldSpecLabel)

	if !ok {
		t.Errorf("%v\n"+
			"Error: Could not convert 'txtFieldSpec' to "+
			"'*TextFieldSpecLabel'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err = txtFieldLabelThree.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldLabelThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtFieldLabelThree.CopyOutITextField(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldLabelThree."+
			"CopyOutITextField()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelThree.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	expectedTextLabelLen := len(expectedTextLabel)

	actualTextLabelLen :=
		txtFieldLabelThree.GetFormattedStrLength()

	if expectedTextLabelLen != actualTextLabelLen {
		t.Errorf("%v\n"+
			"Error: Expected Label Length = '%v'\n"+
			"Instead, Actual Label Length = '%v'\n",
			ePrefix,
			expectedTextLabelLen,
			actualTextLabelLen)

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

	_,
		err = txtFieldLabelOne.CopyOutPtr(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldLabelOne."+
			"txtFieldLabelOne()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecLabel_CopyOutPtr_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_CopyOutITextField_000100()",
		"")

	txtFieldLabel := TextFieldSpecLabel{}

	_,
		err := txtFieldLabel.CopyOutPtr(
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtFieldLabel.CopyOutPtr()\n"+
			"because 'txtFieldLabel' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETUNRED!\n",
			ePrefix.String())

		return
	}

}

func TestTextFieldSpecLabel_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_Empty_000100()",
		"")

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

	_,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix.XCtx("txtFieldLabelOne"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error string from txtFieldLabelOne.GetFormattedText()\n"+
			"because txtFieldLabelOne is empty and txtFieldLabelOne.textLabel is nil.\n"+
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

	txtFieldLabelTwo := TextFieldSpecLabel{}

	txtFieldLabelTwo.Empty()
}

func TestTextFieldSpecLabel_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_equal_000100()",
		"")

	labelRunes := []rune("12345")
	fieldLen := 14
	txtJustify := TxtJustify.Left()

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

	txtFieldLabelMolecule := textFieldSpecLabelMolecule{}

	result :=
		txtFieldLabelMolecule.equal(
			txtFieldLabelOne,
			nil)

	if result == true {
		t.Errorf("%v Test #1\n"+
			"Error: Expected result = 'false'\n"+
			"Instead, result = 'true'\n",
			ePrefix.String())

		return
	}

	result =
		txtFieldLabelMolecule.equal(
			nil,
			txtFieldLabelOne)

	if result == true {
		t.Errorf("%v Test #2\n"+
			"Error: Expected result = 'false'\n"+
			"Instead, result = 'true'\n",
			ePrefix.String())

		return
	}

	labelRunes = []rune("67890")
	fieldLen = 20
	txtJustify = TxtJustify.Right()

	var txtFieldLabelTwo *TextFieldSpecLabel

	txtFieldLabelTwo,
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		labelRunes,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	result =
		txtFieldLabelMolecule.equal(
			txtFieldLabelTwo,
			txtFieldLabelOne)

	if result == true {
		t.Errorf("%v Test #3\n"+
			"Error: Expected result = 'false'\n"+
			"Instead, result = 'true'\n",
			ePrefix.String())

		return
	}

	labelRunes = []rune("12345")
	fieldLen = 19
	txtJustify = TxtJustify.Left()

	var txtFieldLabelThree *TextFieldSpecLabel

	txtFieldLabelThree,
		err = TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		labelRunes,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	result =
		txtFieldLabelMolecule.equal(
			txtFieldLabelThree,
			txtFieldLabelOne)

	if result == true {
		t.Errorf("%v Test #4\n"+
			"Error: Expected result = 'false'\n"+
			"Instead, result = 'true'\n",
			ePrefix.String())

		return
	}

}

func TestTextFieldSpecLabel_EqualITextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_EqualITextField_000100()",
		"")

	txtFieldLabelOne := TextFieldSpecLabel{}

	var txtFieldLabelTwo *TextFieldSpecLabel

	areEqual :=
		txtFieldLabelOne.EqualITextField(txtFieldLabelTwo)

	if areEqual == true {
		t.Errorf("%v - ERROR\n"+
			"areEqual = txtFieldLabelOne.EqualITextField(txtFieldLabelTwo)\n"+
			"Expected areEqual == false\n"+
			"because 'txtFieldLabelTwo' is a nil pointer.\n"+
			"HOWEVER, areEqual == true!\n",
			ePrefix.String())

		return
	}

	txtFieldLabelThree,
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
		"Hello World",
		24,
		TxtJustify.Left(),
		ePrefix.XCtx(
			"txtFieldLabelThree"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecLabel{}.NewPtrTextLabel()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	timeZoneName := "America/Chicago"

	var tzLocPtr *time.Location

	tzLocPtr,
		err = time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		14,
		15,
		28,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTime *TextFieldSpecDateTime

	txtFieldDateTime,
		err = TextFieldSpecDateTime{}.NewPtrDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTime"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewPtrDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	areEqual =
		txtFieldLabelThree.EqualITextField(txtFieldDateTime)

	if areEqual == true {
		t.Errorf("%v - ERROR\n"+
			"areEqual = txtFieldLabelThree.EqualITextField(txtFieldDateTime)\n"+
			"Expected areEqual == 'false'\n"+
			"because 'txtFieldDateTime' is of type 'TextFieldSpecDateTime'.\n"+
			"HOWEVER, areEqual == 'true'!\n",
			ePrefix.String())

		return
	}

	var txtITextFieldSpecLabel ITextFieldSpecification

	txtITextFieldSpecLabel,
		err = txtFieldLabelThree.CopyOutITextField(
		ePrefix.XCtx(
			"txtITextFieldSpecLabel"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldLabelThree.CopyOutITextField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	areEqual = txtFieldLabelThree.EqualITextField(
		txtITextFieldSpecLabel)

	if areEqual == false {
		t.Errorf("%v - ERROR\n"+
			"areEqual = txtFieldLabelThree.EqualITextField(txtITextFieldSpecLabel)\n"+
			"Expected areEqual == 'true'\n"+
			"HOWEVER, areEqual == 'false'!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecLabel_empty_000100(t *testing.T) {

	txtFieldLabelMolecule := textFieldSpecLabelMolecule{}

	txtFieldLabelMolecule.empty(nil)
}

func TestTextFieldSpecLabel_EqualITextField_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_EqualITextField_000100()",
		"")

	txtFieldLabelOne := TextFieldSpecLabel{}

	var txtFieldFiller *TextFieldSpecLabel

	areEqual :=
		txtFieldLabelOne.EqualITextField(txtFieldFiller)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtFieldLabelOne.EqualITextField(txtFieldLabelTwo)\n"+
			"because 'txtFieldLabelTwo' is a nil pointer.\n"+
			"HOWEVER, areEqual == true !\n",
			ePrefix.String())

	}

	return
}

func TestTextFieldSpecLabel_EqualITextField_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_EqualITextField_000300()",
		"")

	txtFieldLabelOne := TextFieldSpecLabel{}

	areEqual :=
		txtFieldLabelOne.EqualITextField(nil)

	if areEqual == true {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtFieldLabelOne.EqualITextField(txtFieldLabelTwo)\n"+
			"because the input is 'nil'.\n"+
			"HOWEVER, EqualITextField(nil) == 'true'!\n",
			ePrefix.String())

	}

	return
}

func TestTextFieldSpecLabel_GetFieldLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_GetFieldLength_000100()",
		"")

	txtFieldLabelOne := TextFieldSpecLabel{}

	actualFieldLen := txtFieldLabelOne.GetFieldLength()

	if actualFieldLen != 0 {
		t.Errorf("%v\n"+
			"Error: Expected actualFieldLen == 0 because\n"+
			"txtFieldLabelOne is empty.\n"+
			"HOWEVER 'actualFieldLen' IS A NON-ZERO VALUE!\n"+
			"actualFieldLen = '%v'\n",
			ePrefix.String(),
			actualFieldLen)
	}

}

func TestTextFieldSpecLabel_GetFormattedStrLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_GetFormattedStrLength_000100()",
		"")

	txtFieldLabelOne := TextFieldSpecLabel{}

	formattedStrLen :=
		txtFieldLabelOne.GetFormattedStrLength()

	if formattedStrLen != -1 {
		t.Errorf("%v - ERROR\n"+
			"Expected Formatted String Length = -1\n"+
			"Instead, Formatted String Length = %v\n",
			ePrefix.String(),
			formattedStrLen)

		return
	}

	return
}

func TestTextFieldSpecLabel_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_CopyOut_000100()",
		"")

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
		ePrefix.XCtx(
			"txtFieldLabelOne"))

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

	var formattedText string

	formattedText,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != formattedText {
		t.Errorf("%v - ERROR\n"+
			"Expected Formatted Text Label = '%v'\n"+
			"Instead, Formatted Text Label = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedTextLabel,
			formattedText)

		return
	}

	formattedText,
		err = txtFieldLabelOne.GetFormattedText(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldLabelOne."+
			"GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecLabel_GetTextJustification_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_GetFieldLength_000100()",
		"")

	txtFieldLabelOne := TextFieldSpecLabel{}

	txtJustification :=
		txtFieldLabelOne.GetTextJustification()

	if txtJustification != TxtJustify.None() {
		t.Errorf("%v\n"+
			"Error: Expected txtJustification == TxtJustify.None()\n"+
			"because txtFieldLabelOne is empty.\n"+
			"Instead, txtJustification string = '%v'\n"+
			"  txtJustification integer value = '%v'\n",
			ePrefix.String(),
			txtJustification.String(),
			txtJustification.XValueInt())
	}

}

func TestTextFieldSpecLabel_GetTextLabel_000100(t *testing.T) {

	txtLabelOne := TextFieldSpecLabel{}

	strResult := txtLabelOne.GetTextLabel()

	if len(strResult) > 0 {
		t.Errorf("TestTextFieldSpecLabel_GetTextLabel_000100()\n"+
			"Error: Expected txtLabelOne.GetTextLabel() to yield an empty string.\n"+
			"HOWEVER, THE STRING IS NOT EMPTY!\n"+
			"strResult = '%v'\n",
			strResult)
	}

}

func TestTextFieldSpecLabel_GetTextLabelRunes_000100(t *testing.T) {

	txtLabelOne := TextFieldSpecLabel{}

	runeArray := txtLabelOne.GetTextLabelRunes()

	if runeArray != nil {
		t.Errorf("TestTextFieldSpecLabel_GetTextLabelRunes_000100()\n"+
			"Error: Expected txtLabelOne.GetTextLabelRunes() to yield a 'nil' value.\n"+
			"HOWEVER, THE RESULT IS NOT 'NIL'!\n"+
			"runeArray = '%v'\n",
			runeArray)
	}

}

func TestTextFieldSpecLabel_GetTextLabelRunes_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_GetTextLabelRunes_000200()",
		"")

	expectedLabelRunes := []rune("12345")
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	txtLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		expectedLabelRunes,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualRuneArray := txtLabelOne.GetTextLabelRunes()

	areEqual := strMechPreon{}.ptr().
		equalRuneArrays(
			expectedLabelRunes,
			actualRuneArray)

	if !areEqual {
		t.Errorf("%v\n"+
			"Error: Expected expectedLabelRunes == actualRuneArray\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedLabelRunes = '%v'\n"+
			"actualRuneArray    = '%v'\n",
			ePrefix.String(),
			string(expectedLabelRunes),
			string(actualRuneArray))
	}

}

func TestTextFieldSpecLabel_isValidTextFieldLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_isValidTextFieldLabel_000100()",
		"")

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

	txtFieldLabelAtom := textFieldSpecLabelAtom{}

	_,
		err = txtFieldLabelAtom.isValidTextFieldLabel(
		nil,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelAtom.isValidTextFieldLabel()\n"+
			"because 'txtFieldLabel' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtFieldLabelOne.fieldLen = -99

	_,
		err = txtFieldLabelAtom.isValidTextFieldLabel(
		txtFieldLabelOne,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelAtom.isValidTextFieldLabel()\n"+
			"because 'txtFieldLabelOne' field length = -99.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtFieldLabelOne.fieldLen = 13

	txtFieldLabelOne.textJustification = TxtJustify.None()

	_,
		err = txtFieldLabelAtom.isValidTextFieldLabel(
		txtFieldLabelOne,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelAtom.isValidTextFieldLabel()\n"+
			"because 'txtFieldLabelOne' text justification = 'NONE'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextFieldSpecLabel_isTextLabelValid_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_isTextLabelValid_000100()",
		"")

	textLabel := make([]rune, 0)

	txtFieldLabelElectron := textFieldSpecLabelElectron{}

	_,
		err :=
		txtFieldLabelElectron.isTextLabelValid(
			textLabel,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelElectron.isTextLabelValid()\n"+
			"because 'textLabel' is an empty rune array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	textLabel = []rune{'H', 'e', 0, 'l', 'l', 'o'}

	_,
		err =
		txtFieldLabelElectron.isTextLabelValid(
			textLabel,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelElectron.isTextLabelValid()\n"+
			"because 'textLabel' contains a rune with a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextFieldSpecLabel_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_CopyOut_000100()",
		"")

	labelRunes := []rune("12345")
	fieldLen := 14
	txtJustify := TxtJustify.Left()

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabelRunes(
		labelRunes,
		fieldLen,
		txtJustify,
		ePrefix.XCtx(
			"txtFieldLabelOne"))

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

	err = txtFieldLabelOne.IsValidInstanceError(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldLabelOne."+
			"IsValidInstanceError()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel,
		err = txtFieldLabelTwo.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel,
		err = txtFieldLabelTwo.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

func TestTextFieldSpecLabel_NewTextTextLabelRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_TextTextLabelRunes_000100()",
		"")

	labelTextRunes := []rune("12345")
	fieldLen := 11

	expectedOutput :=
		strings.Repeat(" ", 6) +
			string(labelTextRunes)

	txtFieldLabel,
		err := TextFieldSpecLabel{}.NewTextLabelRunes(
		labelTextRunes,
		fieldLen,
		TxtJustify.Right(),
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	actualOutput := txtFieldLabel.String()

	if expectedOutput != actualOutput {
		t.Errorf("%v\n"+
			"Error: Expected output string = '%v'\n"+
			"         Actual output string = '%v'\n",
			ePrefix.String(),
			expectedOutput,
			actualOutput)

	}

	return
}

func TestTextFieldSpecLabel_NewTextTextLabelRunes_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_TextTextLabelRunes_000200()",
		"")

	labelTextRunes := []rune("12345")
	fieldLen := 11

	expectedOutput :=
		strings.Repeat(" ", 3) +
			string(labelTextRunes) +
			strings.Repeat(" ", 3)

	txtFieldLabel,
		err := TextFieldSpecLabel{}.NewTextLabelRunes(
		labelTextRunes,
		fieldLen,
		TxtJustify.Center(),
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	actualOutput := txtFieldLabel.String()

	if expectedOutput != actualOutput {
		t.Errorf("%v\n"+
			"Error: Expected output string = '%v'\n"+
			"         Actual output string = '%v'\n",
			ePrefix.String(),
			expectedOutput,
			actualOutput)

	}

	return
}

func TestTextFieldSpecLabel_NewTextTextLabelRunes_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_TextTextLabelRunes_000300()",
		"")

	labelTextRunes := []rune("12345")
	fieldLen := 11

	_,
		err := TextFieldSpecLabel{}.NewTextLabelRunes(
		labelTextRunes,
		fieldLen,
		TxtJustify.None(),
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from NewTextLabelRunes()\n"+
			"because Text Justification was passed as 'None'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecLabel_NewTextTextLabelRunes_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_TextTextLabelRunes_000400()",
		"")

	labelTextRunes := []rune("12345")
	fieldLen := 1000001

	_,
		err := TextFieldSpecLabel{}.NewTextLabelRunes(
		labelTextRunes,
		fieldLen,
		TxtJustify.Right(),
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from NewTextLabelRunes()\n"+
			"because Field Length == 1,000,001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return

}

func TestTextFieldSpecLabel_NewTextTextLabelRunes_000500(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_TextTextLabelRunes_000500()",
		"")

	labelTextRunes := make([]rune, 0)
	fieldLen := 11

	_,
		err := TextFieldSpecLabel{}.NewTextLabelRunes(
		labelTextRunes,
		fieldLen,
		TxtJustify.Right(),
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from NewTextLabelRunes()\n"+
			"because labelTextRunes is a zero length array\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecLabel_NewTextTextLabelRunes_000600(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_TextTextLabelRunes_000600()",
		"")

	var labelTextRunes []rune
	fieldLen := 11

	_,
		err := TextFieldSpecLabel{}.NewTextLabelRunes(
		labelTextRunes,
		fieldLen,
		TxtJustify.Right(),
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from NewTextLabelRunes()\n"+
			"because labelTextRunes is nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return

}

func TestTextFieldSpecLabel_NewTextTextLabelRunes_000700(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_NewTextTextLabelRunes_000700()",
		"")

	labelTextRunes := []rune("12345")
	fieldLen := -2

	_,
		err := TextFieldSpecLabel{}.NewTextLabelRunes(
		labelTextRunes,
		fieldLen,
		TxtJustify.Right(),
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from NewTextLabelRunes()\n"+
			"because fieldLen has a value of -2\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecLabel_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_Read_000100()",
		"")

	label := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	lenExpectedTextLabel :=
		len(expectedTextLabel)

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

	p := make([]byte, 500)

	var n, readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtFieldLabelOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtFieldLabelOne.Read(p)\n"+
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

	if txtFieldLabelOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtFieldLabelOne.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedTextLabel {
		t.Errorf("%v\n"+
			"Byte Length Error: txtFieldLabelOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedTextLabel,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextLabel),
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

	return
}

func TestTextFieldSpecLabel_ReadInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_Read_000100()",
		"")

	label := "1234567890"
	fieldLen := 18
	txtJustify := TxtJustify.Center()

	expectedTextLabel :=
		strings.Repeat(" ", 4) +
			label +
			strings.Repeat(" ", 4)

	lenExpectedTextLabel :=
		len(expectedTextLabel)

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

	p := make([]byte, 5)

	var n int

	n,
		err = txtFieldLabelOne.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldLabelOne.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: txtFieldLabelOne.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.XCtxEmpty().String(),
			n)

		return
	}

	p = make([]byte, 100)

	txtFieldLabelOne.ReaderInitialize()

	var readBytesCnt int
	var actualStr string
	n = 0

	for {

		n,
			err = txtFieldLabelOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtFieldLabelOne.Read(p)\n"+
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

	if txtFieldLabelOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtFieldLabelOne.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedTextLabel {
		t.Errorf("%v\n"+
			"Byte Length Error: txtFieldLabelOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedTextLabel,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedTextLabel),
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

	if txtFieldLabelOne.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but txtFieldLabelOne.textLineReader\n"+
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
			err = txtFieldLabelOne.Read(p)

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
			"Error Returned From txtFieldLabelOne.Read(p)\n"+
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

	if txtFieldLabelOne.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but txtFieldLabelOne.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecLabel_SetFieldLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetFieldLength_000100()",
		"")

	newFieldLen := 22

	txtFieldLabelOne := TextFieldSpecLabel{}

	err := txtFieldLabelOne.SetFieldLength(
		newFieldLen,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualFieldLength := txtFieldLabelOne.GetFieldLength()

	if newFieldLen != actualFieldLength {
		t.Errorf("%v\n"+
			"Error: Expected Field Length = '%v'\n"+
			"         Actual Field Length = '%v'\n"+
			"THEY ARE NOT EQUAL!\n",
			ePrefix,
			newFieldLen,
			actualFieldLength)
	}

}

func TestTextFieldSpecLabel_SetFieldLength_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetFieldLength_000200()",
		"")

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

	newFieldLen := 15

	err = txtFieldLabelOne.SetFieldLength(
		newFieldLen,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())

		return
	}

	actualFieldLength := txtFieldLabelOne.GetFieldLength()

	if newFieldLen != actualFieldLength {
		t.Errorf("%v\n"+
			"Error: Expected Field Length = '%v'\n"+
			"         Actual Field Length = '%v'\n"+
			"THEY ARE NOT EQUAL!\n",
			ePrefix,
			newFieldLen,
			actualFieldLength)
	}

}

func TestTextFieldSpecLabel_SetFieldLength_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetFieldLength_000300()",
		"")

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

	newFieldLen := -2

	err = txtFieldLabelOne.SetFieldLength(
		newFieldLen,
		ePrefix)

	if err == nil {

		t.Errorf("%v\n"+
			"Expected an error return from txtFieldLabelOne.SetFieldLength()\n"+
			"because 'newFieldLen' = -2\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

	}

	return
}

func TestTextFieldSpecLabel_SetText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetText_000100()",
		"")

	expectedLabel := "67890"

	initialLabel := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
		initialLabel,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLabel := txtFieldLabelOne.GetTextLabel()

	if actualLabel != initialLabel {
		t.Errorf("%v Test #1\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			initialLabel,
			actualLabel)

		return
	}

	err =
		txtFieldLabelOne.SetText(
			expectedLabel,
			&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLabel = txtFieldLabelOne.GetTextLabel()

	if actualLabel != expectedLabel {
		t.Errorf("%v Test #2\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedLabel,
			actualLabel)

	}

	return
}

func TestTextFieldSpecLabel_SetText_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetText_000200()",
		"")

	txtLabelOne := TextFieldSpecLabel{}

	err :=
		txtLabelOne.SetText(
			"",
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtLabelOne.SetText()\n"+
			"because the input label is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLabelTwo := TextFieldSpecLabel{}

	expectedLabel := "12345"

	err =
		txtLabelTwo.SetText(
			expectedLabel,
			&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLabel := txtLabelTwo.GetTextLabel()

	if expectedLabel != actualLabel {
		t.Errorf("%v Test #1\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedLabel,
			actualLabel)

	}

	return
}

func TestTextFieldSpecLabel_SetTextJustification_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetTextJustification_000100()",
		"")

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldLabelOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v Test #1\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	txtJustify = TxtJustify.Left()

	expectedTextLabel =
		label +
			strings.Repeat(" ", 8)

	err =
		txtFieldLabelOne.SetTextJustification(
			txtJustify,
			ePrefix)

	if err != nil {
		t.Errorf("%v\n Test #2",
			err.Error())
		return
	}

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldLabelOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v Test #2\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
		return
	}

	txtJustify = TxtJustify.Right()

	expectedTextLabel =
		strings.Repeat(" ", 8) +
			label

	err =
		txtFieldLabelOne.SetTextJustification(
			txtJustify,
			ePrefix)

	if err != nil {
		t.Errorf("%v Test #3\n",
			err.Error())
		return
	}

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldLabelOne #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v Test #3\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)
		return
	}

}

func TestTextFieldSpecLabel_SetTextJustification_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetTextJustification_000200()",
		"")

	txtLabelOne := TextFieldSpecLabel{}

	err :=
		txtLabelOne.SetTextJustification(
			TxtJustify.None(),
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtLabelOne.SetTextJustification()\n"+
			"because input justification = TxtJustify.None() .\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelTwo.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelThree.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelThree.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

func TestTextFieldSpecLabel_setTextLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetText_000100()",
		"")

	txtFieldLabelNanobot := textFieldSpecLabelNanobot{}

	textLabel := []rune("12345")
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	err :=
		txtFieldLabelNanobot.setTextFieldLabel(
			nil,
			textLabel,
			fieldLen,
			txtJustify,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected error return from txtFieldLabelNanobot.setTextFieldLabel()\n"+
			"because 'textFieldLabel' is a 'nil' pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	textFieldLabel := TextFieldSpecLabel{}

	textLabel = []rune{'H', 'e', 'l', 0, 'l', 'o'}

	err =
		txtFieldLabelNanobot.setTextFieldLabel(
			&textFieldLabel,
			textLabel,
			fieldLen,
			txtJustify,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected error return from txtFieldLabelNanobot.setTextFieldLabel()\n"+
			"because 'textLabel' contains a rune with a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

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

	txtFieldLabelOne := TextFieldSpecLabel{}.New()

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel,
		err = txtFieldLabelTwo.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	txtFieldLabelOne := TextFieldSpecLabel{}.New()

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel,
		err = txtFieldLabelTwo.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	txtFieldLabelOne := TextFieldSpecLabel{}.New()

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel,
		err = txtFieldLabelTwo.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

	txtFieldLabelOne := TextFieldSpecLabel{}.New()

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

	txtFieldLabelOne := TextFieldSpecLabel{}.New()

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

	txtFieldLabelOne := TextFieldSpecLabel{}.New()

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

	var actualLabel string

	actualLabel,
		err = txtFieldLabelOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedTextLabel != actualLabel {
		t.Errorf("%v - txtFieldLabelOne\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedTextLabel,
			actualLabel)

		return
	}

	actualLabel,
		err = txtFieldLabelTwo.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

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

func TestTextFieldSpecLabel_SetTextRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetTextRunes_000100()",
		"")

	expectedLabel := "67890"

	initialLabel := "12345"
	fieldLen := 13
	txtJustify := TxtJustify.Center()

	txtFieldLabelOne,
		err := TextFieldSpecLabel{}.NewPtrTextLabel(
		initialLabel,
		fieldLen,
		txtJustify,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLabel := txtFieldLabelOne.GetTextLabel()

	if actualLabel != initialLabel {
		t.Errorf("%v Test #1\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			initialLabel,
			actualLabel)

		return
	}

	err =
		txtFieldLabelOne.SetTextRunes(
			[]rune(expectedLabel),
			&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLabel = txtFieldLabelOne.GetTextLabel()

	if actualLabel != expectedLabel {
		t.Errorf("%v Test #2\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedLabel,
			actualLabel)

	}

	return
}

func TestTextFieldSpecLabel_SetTextRunes_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_SetTextRunes_000200()",
		"")

	txtLabelOne := TextFieldSpecLabel{}

	var txtRunes []rune

	err :=
		txtLabelOne.SetTextRunes(
			txtRunes,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtLabelOne.SetText()\n"+
			"because the input label is an empty rune array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtLabelTwo := TextFieldSpecLabel{}

	expectedLabel := "12345"

	err =
		txtLabelTwo.SetTextRunes(
			[]rune(expectedLabel),
			&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualLabel := txtLabelTwo.GetTextLabel()

	if expectedLabel != actualLabel {
		t.Errorf("%v Test #1\n"+
			"Error: Expected Label = '%v'\n"+
			"Instead, Actual Label = '%v'\n",
			ePrefix,
			expectedLabel,
			actualLabel)

		return
	}

	badLabel := []rune{'6', '7', 0, '8', '9'}

	err =
		txtLabelTwo.SetTextRunes(
			badLabel,
			&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from txtLabelOne.SetText()\n"+
			"because the input label rune array contains a zero element value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

	}

	return
}

func TestTextFieldSpecLabel_String_000100(t *testing.T) {

	txtLabelOne := TextFieldSpecLabel{}

	strResult := txtLabelOne.String()

	if !strings.Contains(strResult, "Error") {
		t.Errorf("TestTextFieldSpecLabel_String_000100()\n"+
			"Error: Expected txtLabelOne.String() to yield an error.\n"+
			"HOWEVER, THE STRING DID NOT CONTAIN THE TEXT 'Error'!\n"+
			"strResult = '%v'\n",
			strResult)
	}
}

func TestTextFieldSpecLabel_TextFieldName_0001000(t *testing.T) {

	ePrefix := "TestTextFieldSpecLabel_TextFieldName_0001000()"

	expectedTxtFieldName := "TextFieldSpecLabel"

	actualTxtFieldName := TextFieldSpecLabel{}.
		NewPtr().TextFieldName()

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

	txtFieldLabelOne := TextFieldSpecLabel{}.New()

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

	actualTxtTypeName := TextFieldSpecLabel{}.
		NewPtr().TextTypeName()

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

	txtFieldLabelOne := TextFieldSpecLabel{}.New()

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

func TestTextFieldSpecLabel_textFieldSpecLabelElectron_000100(t *testing.T) {
	// Edge Cases

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecLabel_textFieldSpecLabelElectron_000100()",
		"")

	txtFieldLabelElectron1 := textFieldSpecLabelElectron{}

	_ =
		txtFieldLabelElectron1.isFieldLengthValid(
			37,
			&ePrefix)

	txtFieldLabelElectron2 := textFieldSpecLabelElectron{}

	_ =
		txtFieldLabelElectron2.isTextJustificationValid(
			[]rune("Hello"),
			25,
			TxtJustify.Right(),
			&ePrefix)

}

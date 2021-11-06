package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
)

func TestTextFieldSpecFiller_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyIn_000100()",
		"")

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	err = fillerTxtFieldTwo.CopyIn(
		fillerTxtFieldOne,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected Filler Text = '%v'\n"+
			"Instead, Actual Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected Filler Text = '%v'\n"+
			"Instead, Actual Filler Text = '%v'\n",
			ePrefix.String(),
			expectedFillerText,
			actualFillerText)
		return
	}

	err = fillerTxtFieldTwo.CopyIn(
		fillerTxtFieldOne,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldTwo.CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var fillerTxtFieldThree *TextFieldSpecFiller

	fillerTxtFieldThree,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		rune(fillerChars[0]),
		fillerRepeatCnt,
		ePrefix.String())

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree.IsValidInstance()\n"+
			"returned 'false'!\n",
			ePrefix.String())

		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldThree) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldThree!\n",
			ePrefix.String())
		return
	}

	fillerTxtFieldThree.Empty()

	err = fillerTxtFieldOne.CopyIn(
		fillerTxtFieldThree,
		ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from fillerTxtFieldOne.CopyIn()\n"+
			"because 'fillerTxtFieldThree' is invalid!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())

		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v - Final Check\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecFiller_CopyIn_000200(t *testing.T) {

	ePrefix := "TestTextFieldSpecFiller_CopyIn_000200()"

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'-'},
		5,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v - Series #1\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	fillerTxtFieldTwo.fillerCharacters = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because fillerTxtFieldTwo.fillerCharacters = 'nil'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)

		return
	}

	err = fillerTxtFieldOne.CopyIn(
		fillerTxtFieldTwo,
		ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from fillerTxtFieldOne.CopyIn(fillerTxtFieldTwo)\n"+
			"because fillerTxtFieldTwo.fillerCharacters = 'nil'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)

		return
	}

	if fillerTxtFieldTwo.IsValidInstance() {
		t.Errorf("%v - Error\n"+
			"Expected a return value of 'false' from fillerTxtFieldTwo.IsValidInstance()\n"+
			"because fillerTxtFieldTwo.fillerCharacters = 'nil'\n"+
			"INSTEAD, IsValidInstance() returned 'true'!\n",
			ePrefix)

		return
	}

	err = fillerTxtFieldTwo.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldTwo.IsValidInstance() {
		t.Errorf("%v - Error\n"+
			"Expected a return value of 'true' from fillerTxtFieldTwo.IsValidInstance()\n"+
			"because fillerTxtFieldTwo was repaired by the call to\n"+
			"fillerTxtFieldTwo.SetTextFiller()\n"+
			"INSTEAD, IsValidInstance() returned 'false'!\n",
			ePrefix)

		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v - Series #3\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.CopyIn(
		fillerTxtFieldTwo,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldThree) {
		t.Errorf("%v - Series #4\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldThree!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected Filler Text = '%v'\n"+
			"Instead, Actual Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldThree.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldThree.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldThree.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldThree.GetFillerCharsRepeatCount())

	}

	return
}

func TestTextFieldSpecFiller_CopyIn_000300(t *testing.T) {

	ePrefix := "TestTextFieldSpecFiller_CopyIn_000300()"

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'-'},
		5,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v - Series #1\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	fillerTxtFieldTwo.fillerCharsRepeatCount = 0

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because fillerTxtFieldTwo.fillerCharsRepeatCount = '0'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)

		return
	}

	err = fillerTxtFieldOne.CopyIn(
		fillerTxtFieldTwo,
		ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from fillerTxtFieldOne.CopyIn(fillerTxtFieldTwo)\n"+
			"because fillerTxtFieldTwo.fillerCharsRepeatCount = '0'\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix)

		return
	}

	if fillerTxtFieldTwo.IsValidInstance() {
		t.Errorf("%v - Error\n"+
			"Expected a return value of 'false' from fillerTxtFieldTwo.IsValidInstance()\n"+
			"because fillerTxtFieldTwo.fillerCharsRepeatCount = '0'\n"+
			"INSTEAD, IsValidInstance() returned 'true'!\n",
			ePrefix)

		return
	}

	err = fillerTxtFieldTwo.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldTwo.IsValidInstance() {
		t.Errorf("%v - Error\n"+
			"Expected a return value of 'true' from fillerTxtFieldTwo.IsValidInstance()\n"+
			"because fillerTxtFieldTwo was repaired by the call to\n"+
			"fillerTxtFieldTwo.SetTextFiller()\n"+
			"INSTEAD, IsValidInstance() returned 'false'!\n",
			ePrefix)

		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v - Series #3\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.CopyIn(
		fillerTxtFieldTwo,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldThree) {
		t.Errorf("%v - Series #4\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldThree!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected Filler Text = '%v'\n"+
			"Instead, Actual Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldThree.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldThree.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldThree.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldThree.GetFillerCharsRepeatCount())

	}

	return
}

func TestTextFieldSpecFiller_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_copyIn_000100()",
		"")

	fillerChars := "-"
	fillerRepeatCnt := 5
	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	txtFieldFillerMolecule := textFieldSpecFillerMolecule{}

	err = txtFieldFillerMolecule.copyIn(
		&fillerTxtFieldTwo,
		fillerTxtFieldOne,
		&ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedFillerText,
			actualFillerText)
		return
	}

	return
}

func TestTextFieldSpecFiller_copyIn_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_copyIn_000200()",
		"")

	var fillerTxtFieldOne *TextFieldSpecFiller

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	txtFieldFillerMolecule := textFieldSpecFillerMolecule{}

	err := txtFieldFillerMolecule.copyIn(
		&fillerTxtFieldTwo,
		fillerTxtFieldOne,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldFillerMolecule.copyIn()\n"+
			"because fillerTxtFieldOne is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	fillerChars := "-"
	fillerRepeatCnt := 5

	fillerTxtFieldOne,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldThree *TextFieldSpecFiller

	err = txtFieldFillerMolecule.copyIn(
		fillerTxtFieldThree,
		fillerTxtFieldOne,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldFillerMolecule.copyIn()\n"+
			"because fillerTxtFieldThree is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	fillerTxtFieldFour := TextFieldSpecFiller{}

	fillerTxtFieldOne.fillerCharacters = nil

	err = txtFieldFillerMolecule.copyIn(
		&fillerTxtFieldFour,
		fillerTxtFieldOne,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldFillerMolecule.copyIn()\n"+
			"because fillerTxtFieldOne.fillerCharacters = nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	fillerTxtFieldOne,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldOne.fillerCharacters =
		make([]rune, 0)

	err = txtFieldFillerMolecule.copyIn(
		&fillerTxtFieldFour,
		fillerTxtFieldOne,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldFillerMolecule.copyIn()\n"+
			"because fillerTxtFieldOne.fillerCharacters = zero length array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecFiller_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyOut_000100()",
		"")

	fillerChar := '-'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(string(fillerChar), fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = fillerTxtFieldOne.CopyOut(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx("fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())

	}

	_,
		err = fillerTxtFieldOne.CopyOut(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldOne.CopyOut()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecFiller_CopyOut_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyOut_000200()",
		"")

	fillerChars := []rune{'*'}

	fillerCharsStr := string(fillerChars)
	fillerRepeatCnt := 25

	expectedFillerText :=
		strings.Repeat(string(fillerChars), fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldOne.fillerCharsRepeatCount = -7

	var fillerTxtFieldTwoPtr *TextFieldSpecFiller

	fillerTxtFieldTwoPtr,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		string(fillerChars),
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldOne,
		err = fillerTxtFieldTwoPtr.CopyOutPtr(
		ePrefix.XCtx("Repairing fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldTwoPtr.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldTwoPtr IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	var fillerTxtFieldThree TextFieldSpecFiller

	fillerTxtFieldThree,
		err = fillerTxtFieldOne.CopyOut(
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerCharsStr != fillerTxtFieldThree.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Characters = '%v'\n",
			ePrefix,
			fillerCharsStr,
			fillerTxtFieldThree.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldThree.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldThree.GetFillerCharsRepeatCount())

	}

	return
}

func TestTextFieldSpecFiller_CopyOut_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyOut_000300()",
		"")

	fillerChar := '&'

	fillerCharsStr := string(fillerChar)
	fillerRepeatCnt := 25

	expectedFillerText :=
		strings.Repeat(fillerCharsStr, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldOne.fillerCharacters = []rune{0}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = fillerTxtFieldOne.CopyOut(
		ePrefix.XCtx("fillerTxtFieldTwo - #1"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from fillerTxtFieldOne.CopyOut()\n"+
			"because fillerTxtFieldOne.fillerCharsRepeatCount = -7\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtx("fillerTxtFieldOne - invalid"))
		return
	}

	err = fillerTxtFieldOne.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("Check repaired fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo = TextFieldSpecFiller{}

	fillerTxtFieldTwo,
		err = fillerTxtFieldOne.CopyOut(
		ePrefix.XCtx("fillerTxtFieldTwo - #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("Check new fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldTwo.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldTwo IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	var fillerTxtFieldThree TextFieldSpecFiller

	fillerTxtFieldThree,
		err = fillerTxtFieldTwo.CopyOut(
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerCharsStr != fillerTxtFieldThree.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Characters = '%v'\n",
			ePrefix,
			fillerCharsStr,
			fillerTxtFieldThree.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldThree.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldThree.GetFillerCharsRepeatCount())

	}

	return
}

func TestTextFieldSpecFiller_CopyOut_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyOut_000400()",
		"")

	fillerTxtFieldOne := TextFieldSpecFiller{}

	var err error

	_,
		err = fillerTxtFieldOne.CopyOut(ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from fillerTxtFieldOne.CopyOut()\n"+
			"because fillerTxtFieldOne is empty.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
	}

	return
}

func TestTextFieldSpecFiller_CopyOutITextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyOutITextField_000100()",
		"")

	fillerTxtFieldOne := TextFieldSpecFiller{}

	_,
		err := fillerTxtFieldOne.CopyOutITextField(
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from fillerTxtFieldOne.CopyOutITextField()\n"+
			"because fillerTxtFieldOne is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			err.Error())
	}

	fillerChars := "-"
	fillerRepeatCnt := 5

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecFiller{}.NewTextFiller()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	var txtFieldSpec ITextFieldSpecification

	txtFieldSpec,
		err = fillerTxtFieldTwo.CopyOutITextField(
		ePrefix)

	var ok bool
	var fillerTxtFieldThree *TextFieldSpecFiller

	fillerTxtFieldThree, ok =
		txtFieldSpec.(*TextFieldSpecFiller)

	if !ok {
		t.Errorf("%v\n"+
			"Error: Could not convert 'txtFieldSpec' to "+
			"'*TextFieldSpecFiller'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err =
		fillerTxtFieldThree.IsValidInstanceError(
			ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by fillerTxtFieldThree.IsValidInstanceError()\n"+
			"'fillerTxtFieldThree' is INVALID!\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	_,
		err = fillerTxtFieldThree.CopyOutITextField(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldThree.CopyOutITextField()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecFiller_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyOutPtr_000100()",
		"")

	fillerChar := '&'

	fillerCharsStr := string(fillerChar)
	fillerRepeatCnt := 25

	expectedFillerText :=
		strings.Repeat(fillerCharsStr, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldOne.fillerCharacters = []rune{0}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = fillerTxtFieldOne.CopyOutPtr(
		ePrefix.XCtx("fillerTxtFieldTwo - #1"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from fillerTxtFieldOne.CopyOut()\n"+
			"because fillerTxtFieldOne.fillerCharsRepeatCount = -7\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtx("fillerTxtFieldOne - invalid"))
		return
	}

	err = fillerTxtFieldOne.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("Check repaired fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo = &TextFieldSpecFiller{}

	fillerTxtFieldTwo,
		err = fillerTxtFieldOne.CopyOutPtr(
		ePrefix.XCtx("fillerTxtFieldTwo - #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("Check new fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldTwo.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldTwo IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	var fillerTxtFieldThree TextFieldSpecFiller

	fillerTxtFieldThree,
		err = fillerTxtFieldTwo.CopyOut(
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerCharsStr != fillerTxtFieldThree.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Characters = '%v'\n",
			ePrefix,
			fillerCharsStr,
			fillerTxtFieldThree.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldThree.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldThree.GetFillerCharsRepeatCount())

		return
	}

	var fillerTxtFieldFour TextFieldSpecFiller

	fillerTxtFieldFour,
		err = fillerTxtFieldThree.CopyOut(
		ePrefix.XCtx("fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)

		return
	}

	_,
		err = fillerTxtFieldThree.CopyOutPtr(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldThree.CopyOutPtr()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecFiller_CopyOutPtr_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyOutPtr_000200()",
		"")

	fillerChar := '&'

	fillerCharsStr := string(fillerChar)
	fillerRepeatCnt := 25

	expectedFillerText :=
		strings.Repeat(fillerCharsStr, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldOne.fillerCharsRepeatCount = 1000001

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = fillerTxtFieldOne.CopyOutPtr(
		ePrefix.XCtx("fillerTxtFieldTwo - #1"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from fillerTxtFieldOne.CopyOut()\n"+
			"because fillerTxtFieldOne.fillerCharsRepeatCount = -7\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.XCtx("fillerTxtFieldOne - invalid"))
		return
	}

	err = fillerTxtFieldOne.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("Check repaired fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo = &TextFieldSpecFiller{}

	fillerTxtFieldTwo,
		err = fillerTxtFieldOne.CopyOutPtr(
		ePrefix.XCtx("fillerTxtFieldTwo - #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("Check new fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldTwo.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldTwo IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	var fillerTxtFieldThree TextFieldSpecFiller

	fillerTxtFieldThree,
		err = fillerTxtFieldTwo.CopyOut(
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerCharsStr != fillerTxtFieldThree.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Characters = '%v'\n",
			ePrefix,
			fillerCharsStr,
			fillerTxtFieldThree.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldThree.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldThree Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldThree.GetFillerCharsRepeatCount())

		return
	}

	var fillerTxtFieldFour TextFieldSpecFiller

	fillerTxtFieldFour,
		err = fillerTxtFieldThree.CopyOut(
		ePrefix.XCtx("fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecFiller_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_copyOut_000100()",
		"")

	var fillerTxtFieldSource *TextFieldSpecFiller

	txtFieldFillerMolecule := textFieldSpecFillerMolecule{}

	_,
		err := txtFieldFillerMolecule.copyOut(
		fillerTxtFieldSource,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldFillerMolecule.copyOut()\n"+
			"because 'fillerTxtFieldSource' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

}

func TestTextFieldSpecFiller_empty_000100(t *testing.T) {

	var fillerTxtFieldOne *TextFieldSpecFiller

	txtFieldFillerMolecule := textFieldSpecFillerMolecule{}

	txtFieldFillerMolecule.empty(fillerTxtFieldOne)

}

func TestTextFieldSpecFiller_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_Empty_000100()",
		"")

	fillerChar := '-'
	fillerRepeatCnt := 5

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldOne.Empty()

	if fillerTxtFieldOne.IsValidInstance() {
		t.Errorf("%v - ERROR\n"+
			"Expected fillerTxtFieldOne.IsValidInstance() to"+
			"return 'false' because 'fillerTxtFieldOne' is empty.\n"+
			"HOWEVER, THE RETURN VALUE WAS 'true'!\n",
			ePrefix.String())

		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	fillerTxtFieldTwo.Empty()

	return
}

func TestTextFieldSpecFiller_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_equal_000100()",
		"")

	var fillerTxtFieldOne *TextFieldSpecFiller
	var fillerTxtFieldTwo *TextFieldSpecFiller

	txtFieldFillerMolecule := textFieldSpecFillerMolecule{}

	isEqual :=
		txtFieldFillerMolecule.equal(
			fillerTxtFieldOne,
			fillerTxtFieldTwo)

	if isEqual == true {
		t.Errorf("%v\n"+
			"Error: Expected isEqual == 'false' because\n"+
			"fillerTxtFieldOne is is a nil pointer.\n"+
			"HOWEVER, isEqual == 'true'!!\n",
			ePrefix)

		return
	}

}

func TestTextFieldSpecFiller_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_Equal_000100()",
		"")

	fillerChar := '-'
	fillerRepeatCnt := 5

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	isEqual := fillerTxtFieldTwo.Equal(fillerTxtFieldOne)

	if isEqual == true {
		t.Errorf("%v\n"+
			"Error: Expected isEqual == 'false' because\n"+
			"fillerTxtFieldTwo is empty.\n"+
			"HOWEVER, isEqual == 'true'!!\n",
			ePrefix)

		return
	}

}

func TestTextFieldSpecFiller_EqualITextField(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_EqualITextField()",
		"")

	var txtFieldLabel *TextFieldSpecLabel

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	isEqual :=
		fillerTxtFieldTwo.EqualITextField(
			txtFieldLabel)

	if isEqual == true {
		t.Errorf("%v\n"+
			"Error: Expected isEqual == 'false' because\n"+
			"txtFieldLabel is type TextFieldSpecLabel.\n"+
			"HOWEVER, isEqual == 'true'!!\n",
			ePrefix)

		return
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}
	var err error

	fillerChars := "-"
	fillerRepeatCnt := 5

	fillerTxtFieldThree,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err =
		fillerTxtFieldThree.IsValidInstanceError(
			ePrefix.XCtx(
				"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	isEqual =
		fillerTxtFieldThree.EqualITextField(
			nil)

	if isEqual == true {
		t.Errorf("%v - ERROR\n"+
			"Expected isEqual == 'false' because\n"+
			"input parameter 'iTextField' is 'nil'.\n"+
			"HOWEVER, 'isEqual' is 'true'!\n",
			ePrefix.String())
		return
	}

	return
}

func TestTextFieldSpecFiller_GetFormattedStrLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_GetFormattedText_000100()",
		"")

	fillerChars := "-"
	fillerRepeatCnt := 5

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedFillerTextLen := len(actualFillerText)

	actualFillerTextLen :=
		fillerTxtFieldOne.GetFormattedStrLength()

	if expectedFillerTextLen != actualFillerTextLen {
		t.Errorf("%v - ERROR\n"+
			"Expected 'expectedFillerTextLen' == 'actualFillerTextLen'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	actualFillerTextLen =
		fillerTxtFieldTwo.GetFormattedStrLength()

	if actualFillerTextLen != -1 {
		t.Errorf("%v - ERROR\n"+
			"Expected 'actualFillerTextLen' to equal minus one (-1).\n"+
			"HOWEVER, THE RETURN VALUE IS NOT MINUS ONE (-1)!\n"+
			"actualFillerTextLen= '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualFillerTextLen)

		return
	}

	return
}

func TestTextFieldSpecFiller_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_GetFormattedText_000100()",
		"")

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected Filler Text = '%v'\n"+
			"Instead, Actual Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	fillerTxtFieldOne.fillerCharacters = nil

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix)

	if err == nil {
		t.Errorf("%v - Test #1\n"+
			"Error:\n"+
			"Expected an error return from fillerTxtFieldOne.GetFormattedText()\n"+
			"because fillerTxtFieldOne.fillerCharacters == nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)
		return
	}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.
		IsValidInstanceError(
			ePrefix.XCtx("fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = fillerTxtFieldTwo.GetFormattedText(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldTwo.GetFormattedText()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecFiller_GetFillerRunes_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_GetFillerRunes_000100()",
		"")

	txtFillerField := TextFieldSpecFiller{}

	actualFillerRunes := txtFillerField.GetFillerRunes()

	if actualFillerRunes != nil {

		t.Errorf("%v\n"+
			"Error: Expected a return value of 'nil' from txtFillerField.GetFillerRunes()\n"+
			"because 'txtFillerField' is empty.\n"+
			"HOWEVER, the return value is NOT 'nil'\n"+
			"actualFillerRunes = '%v'\n",
			ePrefix.String(),
			string(actualFillerRunes))

		return
	}

	fillerChars := []rune{'-', '*'}
	fillerRepeatCnt := 3

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualFillerRunes =
		fillerTxtFieldOne.GetFillerRunes()

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		fillerChars,
		actualFillerRunes)

	if areEqual != true {

		t.Errorf("%v\n"+
			"Error: Expected a fillerChars==actualFillerRunes\n"+
			"because 'fillerChars' is the original input value.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"      fillerChars = '%v'\n"+
			"actualFillerRunes = '%v'\n",
			ePrefix.String(),
			string(fillerChars),
			string(actualFillerRunes))

		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	runeArray :=
		fillerTxtFieldTwo.GetFillerRunes()

	if runeArray != nil {
		t.Errorf("%v - ERROR\n"+
			"Expected 'runeArray' == 'nil'."+
			"HOWEVER, 'runeArray' != 'nil'!!!\n",
			ePrefix.XCtx(
				"runeArray"))

		return
	}

	return
}

func TestTextFieldSpecFiller_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_CopyOutPtr_000200()",
		"")

	txtFieldFillerMolecule := textFieldSpecFillerMolecule{}

	var txtFillerField *TextFieldSpecFiller

	_,
		err := txtFieldFillerMolecule.getFormattedText(
		txtFillerField,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtFieldFillerMolecule.getFormattedText()\n"+
			"because 'txtFillerField' is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

	}

}

func TestTextFieldSpecFiller_isFillerCharsRepeatCountValid_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_isFillerCharsRepeatCountValid_000100()",
		"")

	txtFieldFillerElectron := textFieldSpecFillerElectron{}

	err := txtFieldFillerElectron.isFillerCharsRepeatCountValid(
		0,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtFieldFillerElectron.isFillerCharsRepeatCountValid()\n"+
			"because 'fillerCharsRepeatCount' is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

	}

}

func TestTextFieldSpecFiller_IsValidInstanceError_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_IsValidInstanceError_000100()",
		"")

	fillerTxtFieldThree := TextFieldSpecFiller{}
	var err error

	fillerChars := "-"
	fillerRepeatCnt := 5

	fillerTxtFieldThree,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldThree.IsValidInstanceError()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecFiller_isValidTextFieldSpecFiller_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_isValidTextFieldSpecFiller_000100()",
		"")

	var txtFillerField *TextFieldSpecFiller

	txtFieldFillerAtom := textFieldSpecFillerAtom{}

	_,
		err := txtFieldFillerAtom.isValidTextFieldSpecFiller(
		txtFillerField,
		&ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFieldFillerAtom.isValidTextFieldSpecFiller()\n"+
			"because txtFillerField is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
	}

}

func TestTextFieldSpecFiller_New_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_New_000100()",
		"")

	txtFillerOne := TextFieldSpecFiller{}.New()

	err := txtFillerOne.IsValidInstanceError(
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFillerOne.IsValidInstanceError()\n"+
			"because 'txtFillerOne' member variables contain zero values.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecFiller_NewPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtr_000100()",
		"")

	txtFillerOne := TextFieldSpecFiller{}.NewPtr()

	err := txtFillerOne.IsValidInstanceError(
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected an error return from txtFillerOne.IsValidInstanceError()\n"+
			"because 'txtFillerOne' member variables contain zero values.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFiller_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFiller_000100()",
		"")

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		"",
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler chararacters parameter is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	_,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldTwo.CopyIn()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFiller_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFiller_000200()",
		"")

	fillerChars := "_"
	fillerRepeatCnt := 25

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler chararacters Repeat Count is '-32'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFiller_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFiller_000300()",
		"")

	fillerChar := '_'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 25

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler chararacters Repeat Count is '1,000,001'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFiller_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFiller_000400()",
		"")

	fillerCharRuneArray := []rune{'%'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 15

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		"",
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler chararacters Filler Characters is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFillerRune_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFillerRune_000100()",
		"")

	fillerChar := '-'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		0,
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFillerRune()\n"+
			"because filler chararacters parameter is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	_,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFillerRune()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		'&',
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFillerRune_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFillerRune_000200()",
		"")

	fillerChar := '?'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 15

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFillerRune()\n"+
			"because filler characters repeat count parameter is -32.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		'*',
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFillerRune_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFillerRune_000300()",
		"")

	fillerChar := 'X'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 21

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFillerRune()\n"+
			"because filler characters repeat count parameter is 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		'&',
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFillerRuneArray_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFillerRuneArray_000100()",
		"")

	fillerCharRuneArray := []rune{'-'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from "+
			"TextFieldSpecFiller{}.NewPtrTextFillerRuneArray()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{0},
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFillerRuneArray()\n"+
			"because filler chararacters parameter is zero ([]rune{0}).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx("fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFillerRuneArray_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFillerRuneArray_000200()",
		"")

	fillerCharRuneArray := []rune{'?'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 15

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFillerRuneArray()\n"+
			"because filler characters repeat count parameter is -32.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewPtrTextFillerRuneArray_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewPtrTextFillerRuneArray_000300()",
		"")

	fillerCharRuneArray := []rune{'X'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 21

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFillerRuneArray()\n"+
			"because filler characters repeat count parameter is 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldTwo == nil {
		t.Errorf("%v\n"+
			"'fillerTxtFieldTwo' is a nil pointer.\n"+
			"Test cannot continue!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewTextFiller_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFiller_000100()",
		"")

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		"",
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler chararacters parameter is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	_,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from "+
			"TextFieldSpecFiller{}.NewTextFiller()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		"&&",
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewTextFiller_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFiller_000200()",
		"")

	fillerChars := "?"
	fillerRepeatCnt := 15

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler characters repeat count parameter is -32.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewTextFiller(
		"&&",
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(&fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewTextFiller_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFiller_000300()",
		"")

	fillerChars := "X"
	fillerRepeatCnt := 21

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler characters repeat count parameter is 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewTextFillerRune_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFillerRune_000100()",
		"")

	fillerChar := '-'
	fillerRepeatCnt := 5
	fillerCharsStr := string(fillerChar)
	expectedFillerText :=
		strings.Repeat(fillerCharsStr, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		"",
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler chararacters parameter is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	_,
		err = TextFieldSpecFiller{}.NewTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from "+
			"TextFieldSpecFiller{}.NewTextFillerRune()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerCharsStr != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerCharsStr,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		"&&",
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerCharsStr != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerCharsStr,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewTextFillerRune_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFillerRune_000200()",
		"")

	var fillerChar rune = 0
	fillerRepeatCnt := 5

	_,
		err := TextFieldSpecFiller{}.NewPtrTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from "+
			"TextFieldSpecFiller{}.NewPtrTextFillerRune()\n"+
			"because 'fillerChar' has a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())
	}

	return
}

func TestTextFieldSpecFiller_NewTextFillerRuneArray_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFillerRuneArray_000100()",
		"")

	fillerCharRuneArray := []rune{'-'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		TextFieldSpecDateTime{})

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from "+
			"TextFieldSpecFiller{}.NewTextFillerRuneArray()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		[]rune{0},
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFillerRuneArray()\n"+
			"because filler chararacters parameter is zero ([]rune{0}).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewTextFillerRuneArray_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFillerRuneArray_000200()",
		"")

	fillerCharRuneArray := []rune{'?'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 15

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewTextFillerRuneArray()\n"+
			"because filler characters repeat count parameter is -32.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(&fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_NewTextFillerRuneArray_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFillerRuneArray_000300()",
		"")

	fillerCharRuneArray := []rune{'X'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 21

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewTextFillerRuneArray()\n"+
			"because filler characters repeat count parameter is 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(&fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_newEmpty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_newEmpty_000100()",
		"")

	var txtFillerField TextFieldSpecFiller

	txtFieldFillerMolecule := textFieldSpecFillerMolecule{}

	txtFillerField =
		txtFieldFillerMolecule.newEmpty()

	if txtFillerField.fillerCharacters != nil {
		t.Errorf("%v\n"+
			"Error: txtFillerField.fillerCharacters IS NOT EQUAL to nil\n"+
			"txtFillerField.fillerCharacters = '%v'\n",
			ePrefix.String(),
			string(txtFillerField.fillerCharacters))
	}

	return
}

func TestTextFieldSpecFiller_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_Read_000100()",
		"")

	var err error

	p := make([]byte, 500)

	fillerChars := "-"
	fillerRepeatCnt := 5
	fillerTxtFieldZero := TextFieldSpecFiller{}

	fillerTxtFieldZero,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldZero"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var n, readBytesCnt int
	var actualStr string

	fillerTxtFieldZero.fillerCharacters = nil
	fillerTxtFieldZero.fillerCharsRepeatCount = -999

	_,
		err = fillerTxtFieldZero.Read(p)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from fillerTxtFieldZero.Read(p)\n"+
			"because fillerTxtFieldZero is invalid!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	lenExpectedFillerText := len(expectedFillerText)

	var fillerTxtFieldOne *TextFieldSpecFiller

	fillerTxtFieldOne,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	for {

		n,
			err = fillerTxtFieldOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From fillerTxtFieldOne.Read(p)\n"+
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

	if fillerTxtFieldOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"fillerTxtFieldOne.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedFillerText {
		t.Errorf("%v\n"+
			"Byte Length Error: fillerTxtFieldOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedFillerText,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFillerText),
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

func TestTextFieldSpecFiller_ReadInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_Read_000100()",
		"")

	fillerTxtFieldZero := TextFieldSpecFiller{}

	fillerTxtFieldZero.ReaderInitialize()

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	lenExpectedFillerText := len(expectedFillerText)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	p := make([]byte, 5)

	var n int

	n,
		err = fillerTxtFieldOne.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by fillerTxtFieldOne.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.XCtxEmpty().String(),
			n)

		return
	}

	p = make([]byte, 100)

	fillerTxtFieldOne.ReaderInitialize()

	var readBytesCnt int
	var actualStr string
	n = 0

	for {

		n,
			err = fillerTxtFieldOne.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From fillerTxtFieldOne.Read(p)\n"+
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

	if fillerTxtFieldOne.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"fillerTxtFieldOne.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedFillerText {
		t.Errorf("%v\n"+
			"Byte Length Error: fillerTxtFieldOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedFillerText,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedFillerText),
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

	if fillerTxtFieldOne.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but fillerTxtFieldOne.textLineReader\n"+
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
			err = fillerTxtFieldOne.Read(p)

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
			"Error Returned From fillerTxtFieldOne.Read(p)\n"+
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

	if fillerTxtFieldOne.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but fillerTxtFieldOne.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecFiller_setTxtFieldSpecFiller_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFillerRuneArray_000100()",
		"")

	fillerCharRuneArray := []rune{'-'}
	fillerRepeatCnt := 5

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	txtFieldFillerNanobot := textFieldSpecFillerNanobot{}

	err = txtFieldFillerNanobot.
		setTxtFieldSpecFiller(
			&fillerTxtFieldTwo,
			fillerCharRuneArray,
			fillerRepeatCnt,
			ePrefix.XCtx(
				"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	return
}

func TestTextFieldSpecFiller_setTxtFieldSpecFiller_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewTextFillerRuneArray_000200()",
		"")

	fillerCharRuneArray := []rune{'-'}
	fillerRepeatCnt := 1000001

	txtFieldFillerNanobot1 := textFieldSpecFillerNanobot{}
	fillerTxtFieldOne := TextFieldSpecFiller{}

	err := txtFieldFillerNanobot1.
		setTxtFieldSpecFiller(
			&fillerTxtFieldOne,
			fillerCharRuneArray,
			fillerRepeatCnt,
			ePrefix.XCtx(
				"fillerRepeatCnt := 1000001"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from setTxtFieldSpecFiller()\n"+
			"because fillerRepeatCnt := 1000001\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			err.Error())
		return
	}

	fillerCharRuneArray = []rune{'-'}
	fillerRepeatCnt = -99

	err = txtFieldFillerNanobot1.
		setTxtFieldSpecFiller(
			&fillerTxtFieldOne,
			fillerCharRuneArray,
			fillerRepeatCnt,
			ePrefix.XCtx(
				"fillerRepeatCnt := -99"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from setTxtFieldSpecFiller()\n"+
			"because fillerRepeatCnt = -99\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			err.Error())
		return
	}

	fillerCharRuneArray = nil
	fillerRepeatCnt = 5
	txtFieldFillerNanobot2 := textFieldSpecFillerNanobot{}

	err = txtFieldFillerNanobot2.
		setTxtFieldSpecFiller(
			&fillerTxtFieldOne,
			fillerCharRuneArray,
			fillerRepeatCnt,
			ePrefix.XCtx(
				"fillerCharRuneArray = nil"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from setTxtFieldSpecFiller()\n"+
			"because fillerCharRuneArray = nil\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			err.Error())
		return
	}

	fillerCharRuneArray = []rune{'H', 'e', 'l', 0, 'l', 'o', '!'}
	fillerRepeatCnt = 5
	txtFieldFillerNanobot3 := textFieldSpecFillerNanobot{}

	err = txtFieldFillerNanobot3.
		setTxtFieldSpecFiller(
			&fillerTxtFieldOne,
			fillerCharRuneArray,
			fillerRepeatCnt,
			ePrefix.XCtx(
				"fillerCharRuneArray contains zero value"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from setTxtFieldSpecFiller()\n"+
			"because fillerCharRuneArray contains a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			err.Error())
		return
	}

	fillerCharRuneArray = []rune{'H', 'e', 'l', 'l', 'o', '!'}
	fillerRepeatCnt = 5

	var fillerTxtFieldThree *TextFieldSpecFiller

	err = txtFieldFillerNanobot3.
		setTxtFieldSpecFiller(
			fillerTxtFieldThree,
			fillerCharRuneArray,
			fillerRepeatCnt,
			ePrefix.XCtx(
				"fillerTxtFieldThree is nil pointer"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from setTxtFieldSpecFiller()\n"+
			"because fillerTxtFieldThree is a nil pointer.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			err.Error())
		return
	}

	return
}

func TestTextFieldSpecFiller_SetFillerCharsRepeatCount_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetFillerCharsRepeatCount_000100()",
		"")

	fillerChars := "-"
	fillerRepeatCnt := 5
	newFillerRepeatCnt := 15

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.SetFillerCharsRepeatCount(
		newFillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualRepeatCount := fillerTxtFieldOne.GetFillerCharsRepeatCount()

	if newFillerRepeatCnt != actualRepeatCount {
		t.Errorf("%v\n"+
			"Error: Expected Character Repeat Count = '%v'\n"+
			"Instead- actual Character Repeat Count = '%v'\n",
			ePrefix.String(),
			newFillerRepeatCnt,
			actualRepeatCount)

		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	newFillerRepeatCnt = 25

	err = fillerTxtFieldTwo.SetFillerCharsRepeatCount(
		newFillerRepeatCnt,
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualRepeatCount = fillerTxtFieldTwo.GetFillerCharsRepeatCount()

	if newFillerRepeatCnt != actualRepeatCount {
		t.Errorf("%v - Text #2\n"+
			"Error: Expected Character Repeat Count = '%v'\n"+
			"Instead- actual Character Repeat Count = '%v'\n",
			ePrefix.String(),
			newFillerRepeatCnt,
			actualRepeatCount)

		return
	}

	return
}

func TestTextFieldSpecFiller_SetFillerCharsRepeatCount_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetFillerCharsRepeatCount_000200()",
		"")

	fillerTxtFieldOne := TextFieldSpecFiller{}

	newFillerCharsRepeatCnt := fillerTxtFieldOne.GetFillerCharsRepeatCount()

	if newFillerCharsRepeatCnt != 0 {
		t.Errorf("%v\n"+
			"Error: Expected newFillerCharsRepeatCnt== 0\n"+
			"However:       newFillerCharsRepeatCnt == %v\n",
			ePrefix.String(),
			newFillerCharsRepeatCnt)

		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	newFillerChars := fillerTxtFieldTwo.GetFillerChars()

	if len(newFillerChars) != 0 {
		t.Errorf("%v\n"+
			"Error: Expected len(newFillerChars) == 0\n"+
			"However:        len(newFillerChars) == %v\n",
			ePrefix.String(),
			len(newFillerChars))

		return
	}

}

func TestTextFieldSpecFiller_SetTextFillerRune_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetTextFillerRune_000100()",
		"")

	fillerChar := '-'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne := TextFieldSpecFiller{}

	err := fillerTxtFieldOne.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	err = fillerTxtFieldTwo.SetTextFillerRune(
		0,
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.SetTextFillerRune()\n"+
			"because filler chararacters parameter is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	err = fillerTxtFieldTwo.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		'&',
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_SetTextFillerRune_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetTextFillerRune_000200()",
		"")

	fillerChar := '-'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne := TextFieldSpecFiller{}

	err := fillerTxtFieldOne.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	err = fillerTxtFieldTwo.SetTextFillerRune(
		fillerChar,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.SetTextFillerRune()\n"+
			"because filler characters repeat count parameter is -32.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	err = fillerTxtFieldTwo.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		'&',
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_SetTextFillerRune_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetTextFillerRune_00030()",
		"")

	fillerChar := '-'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne := TextFieldSpecFiller{}

	err := fillerTxtFieldOne.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	err = fillerTxtFieldTwo.SetTextFillerRune(
		fillerChar,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.SetTextFillerRune()\n"+
			"because filler characters repeat count parameter is 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	err = fillerTxtFieldTwo.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRune(
		'&',
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRune(
		fillerChar,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_SetTextFillerRuneArray_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetTextFillerRuneArray_000100()",
		"")

	fillerCharRuneArray := []rune{'-'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne := TextFieldSpecFiller{}

	err := fillerTxtFieldOne.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	err = fillerTxtFieldTwo.SetTextFillerRuneArray(
		[]rune{0},
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.SetTextFillerRune()\n"+
			"because filler chararacters parameter is zero.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	err = fillerTxtFieldTwo.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_SetTextFillerRuneArray_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetTextFillerRuneArray_000200()",
		"")

	fillerCharRuneArray := []rune{'?'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 15

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne := TextFieldSpecFiller{}

	err := fillerTxtFieldOne.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	err = fillerTxtFieldTwo.SetTextFillerRuneArray(
		fillerCharRuneArray,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.SetTextFillerRune()\n"+
			"because filler characters repeat count parameter is -32.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	err = fillerTxtFieldTwo.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldFive := &TextFieldSpecFiller{}

	fillerTxtFieldFive,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFive.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFive - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldSix := TextFieldSpecFiller{}

	err = fillerTxtFieldSix.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldSix.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldSix - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFive.Equal(&fillerTxtFieldSix) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFive IS NOT EQUAL to fillerTxtFieldSix!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecFiller_SetTextFillerRuneArray_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetTextFillerRuneArray_000300()",
		"")

	fillerCharRuneArray := []rune{'X'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 21

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne := TextFieldSpecFiller{}

	err := fillerTxtFieldOne.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	err = fillerTxtFieldTwo.SetTextFillerRuneArray(
		fillerCharRuneArray,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.SetTextFillerRune()\n"+
			"because filler characters repeat count parameter is 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	err = fillerTxtFieldTwo.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		[]rune{'&', '&'},
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(&fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldFive := &TextFieldSpecFiller{}

	fillerTxtFieldFive,
		err = TextFieldSpecFiller{}.NewPtrTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFive.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFive - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldSix := TextFieldSpecFiller{}

	err = fillerTxtFieldSix.SetTextFillerRuneArray(
		fillerCharRuneArray,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldSix.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldSix - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFive.Equal(&fillerTxtFieldSix) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFive IS NOT EQUAL to fillerTxtFieldSix!\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecFiller_SetTextFiller_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_SetTextFiller_000100()",
		"")

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	var fillerTxtFieldOne *TextFieldSpecFiller
	var err error

	fillerTxtFieldOne,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldOne.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		"",
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewPtrTextFiller()\n"+
			"because filler chararacters parameter is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!!\n",
			ePrefix.String())
		return
	}

	err = nil

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from fillerTxtFieldTwo.IsValidInstanceError()\n"+
			"because 'fillerTxtFieldTwo' object is empty.\n",
			ePrefix.XCtxEmpty())
		return
	}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldTwo.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldTwo - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldTwo) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
		return
	}

	var actualFillerText string

	actualFillerText,
		err = fillerTxtFieldOne.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldOne"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText,
		err = fillerTxtFieldTwo.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected fillerTxtFieldTwo Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldTwo.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldTwo.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldTwo.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldTwo Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldTwo Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldTwo.GetFillerCharsRepeatCount())
	}

	fillerTxtFieldThree := TextFieldSpecFiller{}

	err = fillerTxtFieldThree.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldThree.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldThree - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldThree.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldFour *TextFieldSpecFiller

	fillerTxtFieldFour,
		err = TextFieldSpecFiller{}.NewPtrTextFiller(
		"&&",
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: 'fillerTxtFieldFour' is INVAlID!\n",
			ePrefix.String())
		return
	}

	if fillerTxtFieldOne.Equal(fillerTxtFieldFour) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS EQUAL to fillerTxtFieldFour!\n"+
			"This wrong. fillerTxtFieldFour has a different filler character.\n",
			ePrefix)
		return
	}

	err = fillerTxtFieldFour.SetTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix.XCtx("Repairing fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = fillerTxtFieldFour.IsValidInstanceError(
		ePrefix.XCtx("fillerTxtFieldFour - Test #2"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldFour.Equal(fillerTxtFieldOne) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldFour IS NOT EQUAL to fillerTxtFieldOne!\n",
			ePrefix)
		return
	}

	actualFillerText,
		err = fillerTxtFieldFour.GetFormattedText(
		ePrefix.XCtx(
			"fillerTxtFieldFour"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - fillerTxtFieldFour - Test #2\n"+
			"Error: Expected fillerTxtFieldFour Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	if fillerChars != fillerTxtFieldFour.GetFillerChars() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Characters = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Characters = '%v'\n",
			ePrefix,
			fillerChars,
			fillerTxtFieldFour.GetFillerChars())

		return
	}

	if fillerRepeatCnt != fillerTxtFieldFour.GetFillerCharsRepeatCount() {
		t.Errorf("%v\n"+
			"Error: Expected fillerTxtFieldFour Filler Repeat Count = '%v'\n"+
			"Instead, Actual fillerTxtFieldFour Filler Repeat Count = '%v'\n",
			ePrefix,
			fillerRepeatCnt,
			fillerTxtFieldFour.GetFillerCharsRepeatCount())
	}

	return
}

func TestTextFieldSpecFiller_String_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecFiller_String_000100()"

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualFillerText := fillerTxtFieldOne.String()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected Filler Text = '%v'\n"+
			"Instead, Actual Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	fillerTxtFieldOne.fillerCharacters = nil

	actualFillerText = fillerTxtFieldOne.String()

	if !strings.Contains(actualFillerText, "Error") {
		t.Errorf("%v - Test #1\n"+
			"Expected 'actualFillerText' to contain word 'Error'.\n"+
			"HOWEVER, There is no word 'Error' in the returned text.\n",
			ePrefix)
	}

	return
}

func TestTextFieldSpecFiller_TextFieldName_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecFiller_TextFieldName_000100()"

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFieldName := "TextFieldSpecFiller"

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualFieldName := fillerTxtFieldOne.TextFieldName()

	if expectedFieldName != actualFieldName {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected Field Name = '%v'\n"+
			"Instead, Actual Field Name = '%v'\n",
			ePrefix,
			expectedFieldName,
			actualFieldName)
		return
	}

	actualFieldName = "X Y Z A B C"

	actualFieldName = TextFieldSpecFiller{}.NewPtr().TextFieldName()

	if expectedFieldName != actualFieldName {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected Field Name = '%v'\n"+
			"Instead, Actual Field Name = '%v'\n",
			ePrefix,
			expectedFieldName,
			actualFieldName)
	}

	return
}

func TestTextFieldSpecFiller_TextTypeName_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecFiller_TextTypeName_000100()"

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedTextTypeName := "TextFieldSpecFiller"

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewPtrTextFiller(
		fillerChars,
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualTextTypeName := fillerTxtFieldOne.TextTypeName()

	if expectedTextTypeName != actualTextTypeName {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected Text Type Name = '%v'\n"+
			"Instead, Actual Text Type Name = '%v'\n",
			ePrefix,
			expectedTextTypeName,
			actualTextTypeName)

		return
	}

	actualTextTypeName = "A B C D E F X Y Z"

	actualTextTypeName = TextFieldSpecFiller{}.NewPtr().TextFieldName()

	if expectedTextTypeName != actualTextTypeName {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected Text Type Name = '%v'\n"+
			"Instead, Actual Text Type Name = '%v'\n",
			ePrefix,
			expectedTextTypeName,
			actualTextTypeName)

	}

	return
}

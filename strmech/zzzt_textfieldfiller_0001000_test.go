package strmech

import (
	"strings"
	"testing"
)

func TestTextFieldSpecFiller_CopyIn_000100(t *testing.T) {

	ePrefix := "TestTextFieldSpecFiller_CopyIn_000100()"

	fillerChars := "-"
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewConstructor(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected Filler Text = '%v'\n"+
			"Instead, Actual Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #2\n"+
			"Error: Expected Filler Text = '%v'\n"+
			"Instead, Actual Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	var fillerTxtFieldThree *TextFieldSpecFiller

	fillerTxtFieldThree,
		err = TextFieldSpecFiller{}.NewConstructorRune(
		rune(fillerChars[0]),
		fillerRepeatCnt,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !fillerTxtFieldThree.IsValidInstance() {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldThree.IsValidInstance()\n"+
			"returned 'false'!\n",
			ePrefix)

		return
	}

	if !fillerTxtFieldOne.Equal(fillerTxtFieldThree) {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldThree!\n",
			ePrefix)
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
			ePrefix)

		return
	}

	if !fillerTxtFieldOne.Equal(&fillerTxtFieldTwo) {
		t.Errorf("%v - Final Check\n"+
			"Error: fillerTxtFieldOne IS NOT EQUAL to fillerTxtFieldTwo!\n",
			ePrefix)
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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText := fillerTxtFieldThree.GetFormattedText()

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

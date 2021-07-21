package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
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
		err := TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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
		err := TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		string(fillerChars),
		fillerRepeatCnt,
		ePrefix.XCtx("fillerTxtFieldTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	fillerTxtFieldOne,
		err = fillerTxtFieldTwo.CopyOutPtr(
		ePrefix.XCtx("Repairing fillerTxtFieldOne"))

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
		err = fillerTxtFieldOne.CopyOut(
		ePrefix.XCtx("fillerTxtFieldThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	actualFillerText := fillerTxtFieldThree.GetFormattedText()

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
		err := TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText := fillerTxtFieldThree.GetFormattedText()

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
		err := TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText := fillerTxtFieldThree.GetFormattedText()

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
		err := TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText := fillerTxtFieldThree.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructor_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructor_000100()",
		"")

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
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewConstructor(
		"",
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructor(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #3\n"+
			"Error: Expected fillerTxtFieldThree Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldThree Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	return
}

func TestTextFieldSpecFiller_NewConstructor_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructor_000200()",
		"")

	fillerChars := "_"
	fillerRepeatCnt := 25

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
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewConstructor(
		fillerChars,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructor(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructor_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructor_000300()",
		"")

	fillerChar := '_'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 25

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
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewConstructor(
		fillerChars,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructor(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructor_000400(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructor_000400()",
		"")

	fillerCharRuneArray := []rune{'%'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 15

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
		ePrefix.XCtx("fillerTxtFieldOne - Test #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var fillerTxtFieldTwo *TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewConstructor(
		"",
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructor(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructorRune_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructorRune_000100()",
		"")

	fillerChar := '-'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewConstructorRune(
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
		err = TextFieldSpecFiller{}.NewConstructorRune(
		0,
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructorRune_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructorRune_000200()",
		"")

	fillerChar := '?'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 15

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewConstructorRune(
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
		err = TextFieldSpecFiller{}.NewConstructorRune(
		fillerChar,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructorRune_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructorRune_000300()",
		"")

	fillerChar := 'X'
	fillerChars := string(fillerChar)
	fillerRepeatCnt := 21

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewConstructorRune(
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
		err = TextFieldSpecFiller{}.NewConstructorRune(
		fillerChar,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructorRuneArray_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructorRuneArray_000100()",
		"")

	fillerCharRuneArray := []rune{'-'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 5

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewConstructorRuneArray(
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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
		[]rune{0},
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructorRuneArray_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructorRuneArray_000200()",
		"")

	fillerCharRuneArray := []rune{'?'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 15

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewConstructorRuneArray(
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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
		fillerCharRuneArray,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

func TestTextFieldSpecFiller_NewConstructorRuneArray_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecFiller_NewConstructorRuneArray_000300()",
		"")

	fillerCharRuneArray := []rune{'X'}
	fillerChars := string(fillerCharRuneArray)
	fillerRepeatCnt := 21

	expectedFillerText :=
		strings.Repeat(fillerChars, fillerRepeatCnt)

	fillerTxtFieldOne,
		err := TextFieldSpecFiller{}.NewConstructorRuneArray(
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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
		fillerCharRuneArray,
		1000001,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

	fillerTxtFieldTwo := TextFieldSpecFiller{}

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		"",
		fillerRepeatCnt,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

	var fillerTxtFieldTwo TextFieldSpecFiller

	fillerTxtFieldTwo,
		err = TextFieldSpecFiller{}.NewTextFiller(
		fillerChars,
		-32,
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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
			"Expected an error return from TextFieldSpecFiller{}.NewConstructor()\n"+
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRune(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText := fillerTxtFieldOne.GetFormattedText()

	if expectedFillerText != actualFillerText {
		t.Errorf("%v - Test #1\n"+
			"Error: Expected fillerTxtFieldOne Filler Text = '%v'\n"+
			"Instead, Actual fillerTxtFieldOne Filler Text = '%v'\n",
			ePrefix,
			expectedFillerText,
			actualFillerText)
		return
	}

	actualFillerText = fillerTxtFieldTwo.GetFormattedText()

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

	actualFillerText = fillerTxtFieldThree.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

	actualFillerText = fillerTxtFieldFour.GetFormattedText()

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
		err = TextFieldSpecFiller{}.NewConstructorRuneArray(
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

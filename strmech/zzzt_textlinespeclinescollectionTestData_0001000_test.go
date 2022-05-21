package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
)

func createTestTextLineSpecCollection01(
	errorPrefix interface{}) (
	numOfLines int,
	txtLineCollection TextLineSpecLinesCollection,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	txtLineCollection = TextLineSpecLinesCollection{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecCollection01()",
		"")

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	var stdLine01, stdLine02, stdLine03 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 0
	err = txtLineCollection.AddTextLine(
		&stdLine01,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 1
	err = txtLineCollection.AddTextLine(
		&stdLine02,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	stdLine03,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 2
	err = txtLineCollection.AddTextLine(
		&stdLine03,
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	numOfBlankLines := 1

	newLineRunes := []rune{'\n'}

	var blankLines01 TextLineSpecBlankLines

	blankLines01,
		err = TextLineSpecBlankLines{}.NewBlankLineRunes(
		numOfBlankLines,
		newLineRunes,
		ePrefix.XCpy(
			"blankLines01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	err = txtLineCollection.AddTextLine(
		&blankLines01,
		ePrefix.XCpy(
			"blankLines01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	leftMargin := 3
	rightMargin := 3
	textString := "How now brown cow!"

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCpy(
			"plainTextLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	err = txtLineCollection.AddTextLine(
		&plainTextLine01,
		ePrefix.XCpy(
			"plainTextLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	leftMargin = 2
	rightMargin = 2
	textString = "The cow jumped over the moon!"

	var plainTextLine02 TextLineSpecPlainText

	plainTextLine02,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCpy(
			"plainTextLine02"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	err = txtLineCollection.AddTextLine(
		&plainTextLine02,
		ePrefix.XCpy(
			"plainTextLine02"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	var timerLines01 *TextLineSpecTimerLines
	_,
		timerLines01,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	err = txtLineCollection.AddTextLine(
		timerLines01,
		ePrefix.XCpy(
			"timerLines01"))

	numOfLines++

	return numOfLines, txtLineCollection, err
}

func createTestTextLineSpecCollection02(
	errorPrefix interface{}) (
	numOfLines int,
	txtLineCollection TextLineSpecLinesCollection,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	txtLineCollection = TextLineSpecLinesCollection{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecCollection02()",
		"")

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	leftMargin := 3
	rightMargin := 3
	textString := "The cow jumped over the moon!"

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCpy(
			"plainTextLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	err = txtLineCollection.AddTextLine(
		&plainTextLine01,
		ePrefix.XCpy(
			"plainTextLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	var stdLine01, stdLine02 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	err = txtLineCollection.AddTextLine(
		&stdLine01,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	err = txtLineCollection.AddTextLine(
		&stdLine02,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	return numOfLines, txtLineCollection, err
}

func createTestTextLineSpecCollection03(
	errorPrefix interface{}) (
	numOfLines int,
	txtLineCollection TextLineSpecLinesCollection,
	err error) {

	var ePrefix *ePref.ErrPrefixDto

	txtLineCollection = TextLineSpecLinesCollection{}

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecCollection03()",
		"")

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	leftMargin := 3
	rightMargin := 3
	textString := "The cow jumped over the moon!"

	var plainTextLine01 TextLineSpecPlainText

	plainTextLine01,
		err = TextLineSpecPlainText{}.NewDefault(
		leftMargin,
		rightMargin,
		textString,
		ePrefix.XCpy(
			"plainTextLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 0
	err = txtLineCollection.AddTextLine(
		&plainTextLine01,
		ePrefix.XCpy(
			"plainTextLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	var stdLine01, stdLine02 TextLineSpecStandardLine

	stdLine01,
		err = createTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 1
	err = txtLineCollection.AddTextLine(
		&stdLine01,
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	stdLine02,
		err = createTestTextLineSpecStandardLine02(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 2
	err = txtLineCollection.AddTextLine(
		&stdLine02,
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	var stdLine03, stdLine04,
		stdLine05 TextLineSpecStandardLine

	stdLine03,
		err = createTestTextLineSpecStandardLine03(
		3,
		"Helo World! Have a GREAT DAY!",
		-1,
		TxtJustify.Left(),
		2,
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 3
	err = txtLineCollection.AddTextLine(
		&stdLine03,
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	stdLine04,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine04"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 4
	err = txtLineCollection.AddTextLine(
		&stdLine04,
		ePrefix.XCpy(
			"stdLine04"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	stdLine05,
		err = createTestTextLineSpecStandardLine05(
		ePrefix.XCpy(
			"stdLine05"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 5
	err = txtLineCollection.AddTextLine(
		&stdLine05,
		ePrefix.XCpy(
			"stdLine05"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	var txtSolidLine01 TextLineSpecSolidLine

	txtSolidLine01,
		err = TextLineSpecSolidLine{}.NewDefaultSolidLine(
		"-",
		22,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 6
	err = txtLineCollection.AddTextLine(
		&txtSolidLine01,
		ePrefix.XCpy(
			"txtSolidLine01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	var blankLines01 TextLineSpecBlankLines

	blankLines01,
		err = TextLineSpecBlankLines{}.NewDefaultBlankLines(
		2,
		ePrefix.XCpy(
			"blankLines01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 7
	err = txtLineCollection.AddTextLine(
		&blankLines01,
		ePrefix.XCpy(
			"blankLines01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	var timerLines01 *TextLineSpecTimerLines

	_,
		timerLines01,
		err = createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	// Index 8

	err = txtLineCollection.AddTextLine(
		timerLines01,
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

	numOfLines++

	err = txtLineCollection.IsValidInstanceError(
		ePrefix.XCpy(
			"txtLineCollection Final Check"))

	return numOfLines, txtLineCollection, err
}

func TestTextLineSpecLinesCollection_IsValidInstanceError_090100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecLinesCollection_IsValidInstanceError_090100()",
		"")

	_,
		stdLine01,
		err := createTestTextLineSpecCollection01(
		ePrefix.XCpy(
			"stdLine01"))

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

	var stdLine02,
		stdLine03 TextLineSpecLinesCollection

	_,
		stdLine02,
		err = createTestTextLineSpecCollection02(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine02.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine02"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		stdLine03,
		err = createTestTextLineSpecCollection03(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = stdLine03.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

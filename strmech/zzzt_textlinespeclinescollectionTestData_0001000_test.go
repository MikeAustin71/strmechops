package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
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

	stdLine03,
		err = createTestTextLineSpecStandardLine04(
		ePrefix.XCpy(
			"stdLine03"))

	if err != nil {
		return numOfLines, txtLineCollection, err
	}

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

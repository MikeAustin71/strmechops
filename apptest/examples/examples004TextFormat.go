package examples

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"github.com/MikeAustin71/strmechops/strmech"
	"strings"
	"time"
)

type MainTextFormatTest04 struct {
	input string
}

func (mTxt04 MainTextFormatTest04) TxtLineMarquee001() {
	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTextFormatTest04.TxtLineMarquee001()",
		"")

	breakStr := strings.Repeat("=", 50)

	fmt.Printf("\n\n" + breakStr + "\n")

	fmt.Printf("\n Starting Run!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

	marqueeDto := strmech.TextLineTitleMarqueeDto{
		StandardTitleLeftMargin:   " *",
		StandardTitleRightMargin:  "*",
		StandardMaxLineLen:        80,
		StandardTextFieldLen:      99,
		StandardTextJustification: strmech.TxtJustify.Center(),
		NumLeadingBlankLines:      1,
		LeadingSolidLineChar:      "*",
		NumLeadingSolidLines:      2,
		NumTopTitleBlankLines:     1,
		TitleLines:                strmech.TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:  1,
		TrailingSolidLineChar:     "*",
		NumTrailingSolidLines:     2,
		NumTrailingBlankLines:     1,
	}

	strArrayDto := strmech.StringArrayDto{}

	strArrayDto.AddString("Title Line #1")

	strArrayDto.AddString("Title Line #2")

	strArrayDto.AddString("Title Line #3")

	err := marqueeDto.AddTitleLineStrArrayDto(
		strArrayDto,
		ePrefix.XCpy("strArrayDto"))

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	var errTxt string

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		fmt.Println(errTxt)

		return
	}

	dateTime := time.Date(
		2022,
		time.Month(12),
		16,
		14,
		19,
		3,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	err = marqueeDto.AddTitleLineDateTimeStr(
		dateTime,
		dateTimeFormat,
		ePrefix.XCpy(
			"Time Stamp"))

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	/*
		var titleLinesCollection strmech.TextLineSpecLinesCollection
		var numOfTitleLines, maxTitleLineLen int

		numOfTitleLines,
			titleLinesCollection,
			err = marqueeDto.GetTitleLines(
			ePrefix.XCpy("titleLinesCollection"))

		if err != nil {

			fmt.Printf("\n\n%v\n"+
				"%v\n",
				ePrefix.String(),
				err.Error())

			return

		}

		fmt.Printf("\n\nDto Title Lines\n"+
			"Number Of Title Lines: %v\n"+
			"Title Lines:\n\n",
			numOfTitleLines)

		strBuilder := strings.Builder{}

		maxTitleLineLen,
			err = titleLinesCollection.GetFormattedText(
			&strBuilder,
			ePrefix.XCpy(
				"strBuilder<-titleLinesCollection"))

		if err != nil {

			fmt.Printf("\n\n%v\n"+
				"%v\n",
				ePrefix.String(),
				err.Error())

			return

		}

		fmt.Printf("Max Title Line Length: %v\n\n%v\n\n\n",
			maxTitleLineLen,
			strBuilder.String())
	*/

	var titleLineMarquee strmech.TextLineSpecTitleMarquee

	titleLineMarquee,
		err = new(strmech.TextLineSpecTitleMarquee).
		NewMarqueeDto(
			&marqueeDto,
			ePrefix.XCpy(
				"titleLineMarquee<-marqueeDto"))

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

	var marqueeStr string

	marqueeStr,
		err = titleLineMarquee.GetFormattedText(
		ePrefix.XCpy("marqueeStr<-titleLineMarquee"))

	fmt.Printf("** OUTPUT FROM Title Line MARQUEE\n\n")

	fmt.Printf(marqueeStr)

	fmt.Printf("\n\nConverted Print String\n\n")

	sMech := strmech.StrMech{}

	convertedMarqueeText := sMech.ConvertNonPrintableChars(
		[]rune(marqueeStr),
		true)

	fmt.Printf(convertedMarqueeText)

	/*
		   Printed Output

		   	******************************************************************************
		   	******************************************************************************
		   	*                                                                            *
		   	*                               Title Line #1                                *
		   	*                               Title Line #2                                *
		   	*                               Title Line #3                                *
		   	*                  2022-12-16 14:19:03.000000000 -0600 CST                   *
		   	*                                                                            *
		   	******************************************************************************
		   	******************************************************************************

		Converted Output
		\n[SPACE]******************************************************************************\n[SPACE]******************************************************************************\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Title[SPACE]Line[SPACE]#1[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Title[SPACE]Line[SPACE]#2[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Title[SPACE]Line[SPACE]#3[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]2022-12-16[SPACE]14:19:03.000000000[SPACE]-0600[SPACE]CST[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\n[SPACE]******************************************************************************\n[SPACE]******************************************************************************\n\n

	*/

	fmt.Printf("\n\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

func (mTxt04 MainTextFormatTest04) TextLineSpecStandardLine01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecStandardLine01()",
		"")

	label := "12345"
	fieldLen := -37
	txtJustify := strmech.TxtJustify.Center()
	var output string
	var err error

	stdLine02 := strmech.TextLineSpecStandardLine{}

	_,
		err = stdLine02.AddTextFieldLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"stdLine02 - fieldLen invalid."))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(output)
		return
	}

	output =
		fmt.Sprintf("%v\n"+
			"SUCCESSFUL COMPLETION!!!\n\n",
			ePrefix.String())

	fmt.Println(output)
}

func (mTxt04 MainTextFormatTest04) TextLineSpecStandardLine02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecStandardLine02()",
		"")

	var output string

	stdLine01 := strmech.TextLineSpecStandardLine{}

	rightMarginLen := 5

	rightMarginSpec,
		err := strmech.TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Print(output)

		return
	}

	var leftMarginSpec strmech.TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = strmech.TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Print(output)

		return
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := strmech.TxtJustify.Center()

	var labelSpec strmech.TextFieldSpecLabel

	labelSpec,
		err = strmech.TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Print(output)

		return
	}

	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Print(output)

		return
	}

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Print(output)

		return
	}

	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Print(output)

		return
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Print(output)

		return
	}

	stdLine02 := strmech.TextLineSpecStandardLine{}

	err =
		stdLine02.CopyIn(
			&stdLine01,
			ePrefix.XCpy(
				"stdLine02<-stdLine01 invalid newLinesChars"))

	if err != nil {
		output = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Print(output)

		return
	}

	output = fmt.Sprintf("\n%v\n"+
		"SUCCESSFUL COMPLETION\n",
		ePrefix.String())

	fmt.Println(output)
}

func (mTxt04 MainTextFormatTest04) TextLineSpecStandardLine03() error {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecStandardLine03()",
		"")

	lineBar := strings.Repeat("-", 60)

	testUtil := TestUtilities{}

	stdLine01,
		err := testUtil.CreateTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return err
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return err
	}

	lenTxtFields := stdLine01.GetNumOfTextFields()

	initialLastIdx := lenTxtFields - 1

	fmt.Printf("\n%v\n"+
		"Running Function - %v\n"+
		"%v\n\n",
		lineBar,
		ePrefix.String(),
		lineBar)

	fmt.Printf("Initial Length of\n"+
		"stdLine01 Text Fields ='%v'\n"+
		"Initial Last Index = '%v'\n\n",
		lenTxtFields,
		initialLastIdx)

	expectedLabelText := "Xray7 where are you?"

	var labelTxt strmech.TextFieldSpecLabel

	labelTxt,
		err = strmech.TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		strmech.TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		return err
	}

	var newLastIndexId int

	newLastIndexId,
		err = stdLine01.InsertTextField(
		&labelTxt,
		initialLastIdx,
		ePrefix.XCpy(
			fmt.Sprintf(
				"stdLine01[%v]<-labeText",
				initialLastIdx)))

	if err != nil {
		return err
	}

	lenTxtFields = stdLine01.GetNumOfTextFields()

	fmt.Printf("After Insertion Length of\n"+
		"stdLine01 Text Fields ='%v'\n"+
		"New Last Index ID ='%v'\n\n",
		lenTxtFields,
		newLastIndexId)

	var lastITxtField strmech.ITextFieldSpecification

	lastITxtField,
		err =
		stdLine01.GetTextField(
			newLastIndexId,
			ePrefix.XCpy(
				fmt.Sprintf(
					"stdLine01[%v]",
					newLastIndexId)))

	if err != nil {
		return err
	}

	err = lastITxtField.IsValidInstanceError(
		ePrefix.XCpy(
			"lastITxtField"))

	if err != nil {
		return err
	}

	var iTxtFieldSpec strmech.ITextFieldSpecification

	iTxtFieldSpec,
		err = stdLine01.GetTextField(
		initialLastIdx,
		ePrefix.XCpy(
			fmt.Sprintf(
				"stdLine01[%v]",
				initialLastIdx)))

	if err != nil {
		return err
	}

	err = iTxtFieldSpec.IsValidInstanceError(
		ePrefix.XCpy(
			"iTxtFieldSpec - Inserted Field"))

	if err != nil {
		return err
	}

	actualLabelField,
		ok := iTxtFieldSpec.(*strmech.TextFieldSpecLabel)

	if !ok {
		err = fmt.Errorf("\n%v\n"+
			"Error converting stdLine01[%v] to a\n"+
			"label Field!\n",
			ePrefix.String(),
			initialLastIdx)

		return err
	}

	if !actualLabelField.Equal(&labelTxt) {

		err = fmt.Errorf("\n%v\n"+
			"Error!\n"+
			"Expected that stdLine01[%v] would\n"+
			"equal the inserted Label Field.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String(),
			initialLastIdx)

		return err
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			fmt.Sprintf(
				"After stdLine[%v] label insertion",
				initialLastIdx)))

	if err != nil {
		return err
	}

	fmt.Printf("\n%v\n"+
		"Successful Completion - %v\n"+
		"%v\n\n",
		lineBar,
		ePrefix.String(),
		lineBar)

	return err
}

func (mTxt04 MainTextFormatTest04) TextLineSpecStandardLine04() error {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecStandardLine04()",
		"")

	lineBar := strings.Repeat("-", 60)

	testUtil := TestUtilities{}

	stdLine01,
		err := testUtil.CreateTestTextLineSpecStandardLine01(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return err
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return err
	}

	lenTxtFields := stdLine01.GetNumOfTextFields()

	initialTargetIdx := 1

	fmt.Printf("\n%v\n"+
		"Running Function - %v\n"+
		"%v\n\n",
		lineBar,
		ePrefix.String(),
		lineBar)

	fmt.Printf("Initial Length of\n"+
		"stdLine01 Text Fields ='%v'\n"+
		"Initial Target Index = '%v'\n\n",
		lenTxtFields,
		initialTargetIdx)

	expectedLabelText := "Xray7 where are you?"

	var labelTxt strmech.TextFieldSpecLabel

	labelTxt,
		err = strmech.TextFieldSpecLabel{}.NewTextLabel(
		expectedLabelText,
		-1,
		strmech.TxtJustify.Left(),
		ePrefix.XCpy(
			"labelTxt"))

	if err != nil {
		return err
	}

	var newLastIndexId int

	newLastIndexId,
		err = stdLine01.InsertTextField(
		&labelTxt,
		initialTargetIdx,
		ePrefix.XCpy(
			fmt.Sprintf(
				"stdLine01[%v]<-labeText",
				initialTargetIdx)))

	if err != nil {
		return err
	}

	lenTxtFields = stdLine01.GetNumOfTextFields()

	targetIdxPlus1 := initialTargetIdx + 1

	fmt.Printf("After Insertion Length of\n"+
		"stdLine01 Text Fields ='%v'\n"+
		"New Last Index ID ='%v'\n\n",
		lenTxtFields,
		newLastIndexId)

	var targetITxtField strmech.ITextFieldSpecification

	targetITxtField,
		err =
		stdLine01.GetTextField(
			initialTargetIdx,
			ePrefix.XCpy(
				fmt.Sprintf(
					"stdLine01[%v]",
					initialTargetIdx)))

	if err != nil {
		return err
	}

	err = targetITxtField.IsValidInstanceError(
		ePrefix.XCpy(
			"targetITxtField"))

	if err != nil {
		return err
	}

	var targetPlus1ITxtField strmech.ITextFieldSpecification

	targetPlus1ITxtField,
		err =
		stdLine01.GetTextField(
			targetIdxPlus1,
			ePrefix.XCpy(
				fmt.Sprintf(
					"stdLine01[%v]",
					targetIdxPlus1)))

	if err != nil {
		return err
	}

	err = targetPlus1ITxtField.IsValidInstanceError(
		ePrefix.XCpy(
			"targetPlus1ITxtField"))

	if err != nil {
		return err
	}

	var initialITargetField strmech.ITextFieldSpecification

	initialITargetField,
		err = stdLine01.GetTextField(
		initialTargetIdx,
		ePrefix.XCpy(
			fmt.Sprintf(
				"stdLine01[%v]",
				initialTargetIdx)))

	if err != nil {
		return err
	}

	err = initialITargetField.IsValidInstanceError(
		ePrefix.XCpy(
			"initialITargetField"))

	if err != nil {
		return err
	}

	actualLabelField,
		ok := initialITargetField.(*strmech.TextFieldSpecLabel)

	if !ok {
		err = fmt.Errorf("\n%v\n"+
			"Error converting stdLine01[%v] to a\n"+
			"label Field!\n",
			ePrefix.String(),
			initialTargetIdx)

		return err
	}

	if !actualLabelField.Equal(&labelTxt) {

		err = fmt.Errorf("\n%v\n"+
			"Error!\n"+
			"Expected that stdLine01[%v] would\n"+
			"equal the inserted Label Field.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String(),
			initialTargetIdx)

		return err
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			fmt.Sprintf(
				"After stdLine[%v] label insertion",
				initialTargetIdx)))

	if err != nil {
		return err
	}

	fmt.Printf("\n%v\n"+
		"Successful Completion - %v\n"+
		"%v\n\n",
		lineBar,
		ePrefix.String(),
		lineBar)

	return err
}

func (mTxt04 MainTextFormatTest04) TextLineSpecSolidLine01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecSolidLine01()",
		"")

	txtSolidLine03 := strmech.TextLineSpecSolidLine{}

	_,
		err :=
		txtSolidLine03.CopyOutITextLine(
			ePrefix.XCpy(
				"txtSolidLine03->_"))

	if err == nil {

		errorStr := fmt.Sprintf("%v\n"+
			"Error: txtSolidLine03.CopyOutITextLine()\n"+
			"Expected an error return because 'txtSolidLine03' is empty!\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		fmt.Printf(errorStr + "\n")

	}

	return

}

func (mTxt04 MainTextFormatTest04) TextLineSpecPlainText01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecBlkLines01()",
		"")

	leftMarginChars := "  "
	rightMarginChars := "  "
	textString := "The cow jumped over the moon!"
	newLineChars := "\n"
	outStr := ""

	_,
		err := strmech.TextLineSpecPlainText{}.NewPtrPlainTextStrings(
		leftMarginChars,
		rightMarginChars,
		textString,
		newLineChars,
		false,
		strmech.StrMech{})

	if err != nil {
		outStr = fmt.Sprintf("%v\n"+
			"Error returned from: NewPtrPlainTextStrings()\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())
	} else {
		outStr = fmt.Sprintf("%v\n"+
			"TEST FAILED!! No Error was returned from\n"+
			"NewPtrPlainTextStrings()\n",
			ePrefix.String())
	}

	fmt.Printf(outStr)

	return
}

func (mTxt04 MainTextFormatTest04) TextLineSpecBlkLines01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecBlkLines01()",
		"")

	var errText string

	blankLinesAlpha,
		err := strmech.TextLineSpecBlankLines{}.NewDefaultBlankLines(
		3,
		ePrefix.XCpy(
			"blankLinesAlpha"))

	if err != nil {
		errText = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(errText)

		return
	}

	err = blankLinesAlpha.IsValidInstanceError(
		ePrefix.XCpy(
			"blankLinesAlpha"))

	if err != nil {
		errText = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(errText)

		return
	}

	blankLinesBravo := strmech.TextLineSpecBlankLines{}

	err =
		blankLinesBravo.CopyIn(
			&blankLinesAlpha,
			ePrefix.XCpy(
				"blankLinesAlpha->blankLinesBravo"))

	if err != nil {
		errText = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(errText)

		return
	}

	err = blankLinesBravo.IsValidInstanceError(
		ePrefix.XCpy(
			"blankLinesBravo"))

	if err != nil {
		errText = fmt.Sprintf("%v\n",
			err.Error())

		fmt.Println(errText)

		return
	}

	errText = fmt.Sprintf("%v\nSuccessful Completion\n",
		ePrefix.String())

	fmt.Println(errText)

	return
}

func (mTxt04 MainTextFormatTest04) DateTimeTextField01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.DateTimeTextField01()",
		"")

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	var errTxt string

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		fmt.Println(errTxt)

		return
	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		21,
		14,
		19,
		3,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen := 39 + 4

	textJustification := strmech.TxtJustify.Center()

	var txtFieldDateTimeOne strmech.TextFieldSpecDateTime

	txtFieldDateTimeOne,
		err = strmech.TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy("txtFieldDateTimeOne"))

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	var fmtText, convertedFmtText string

	fmtText,
		err = txtFieldDateTimeOne.GetFormattedText(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by txtFieldDateTimeOne.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	fmt.Printf("Formatted string = \n"+
		"          \"%v\"\n",
		fmtText)
	fmt.Println()

	sMech := strmech.StrMech{}

	convertedFmtText = sMech.ConvertNonPrintableChars(
		[]rune("\""+fmtText+"\""),
		false)

	fmt.Printf("Printable string = \n"+
		"          %v\n",
		convertedFmtText)

	return
}

func (mTxt04 MainTextFormatTest04) DateTimeTextField02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.DateTimeTextField02()",
		"")

	timeZoneName := "America/Los_Angeles"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	var errTxt string

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		fmt.Println(errTxt)

		return
	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		10,
		20,
		13,
		34,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen := 39 + 4

	textJustification := strmech.TxtJustify.Center()

	var txtFieldDateTimeOne strmech.TextFieldSpecDateTime

	txtFieldDateTimeOne,
		err = strmech.TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy("txtFieldDateTimeOne"))

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	var fmtText, convertedFmtText string

	fmtText,
		err = txtFieldDateTimeOne.GetFormattedText(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by txtFieldDateTimeOne.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	fmt.Printf("Formatted string = \n"+
		"          \"%v\"\n",
		fmtText)
	fmt.Println()

	sMech := strmech.StrMech{}

	convertedFmtText = sMech.ConvertNonPrintableChars(
		[]rune("\""+fmtText+"\""),
		false)

	fmt.Printf("Printable string = \n"+
		"          %v\n",
		convertedFmtText)

	return
}

func (mTxt04 MainTextFormatTest04) DateTimeTextField03() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.DateTimeTextField03()",
		"")

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	var errTxt string

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		fmt.Println(errTxt)

		return
	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		21,
		14,
		19,
		3,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := 52 + 6

	textJustification := strmech.TxtJustify.Center()

	var txtFieldDateTimeOne strmech.TextFieldSpecDateTime

	txtFieldDateTimeOne,
		err = strmech.TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy("txtFieldDateTimeOne"))

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	var fmtText, convertedFmtText string

	fmtText,
		err = txtFieldDateTimeOne.GetFormattedText(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf("%v\n"+
			"Error returned by txtFieldDateTimeOne.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		fmt.Println(errTxt)

		return
	}

	fmt.Printf("Formatted string = \n"+
		"          \"%v\"\n",
		fmtText)
	fmt.Println()

	sMech := strmech.StrMech{}

	convertedFmtText = sMech.ConvertNonPrintableChars(
		[]rune("\""+fmtText+"\""),
		false)

	fmt.Printf("Printable string = \n"+
		"          %v\n",
		convertedFmtText)

	return
}

func (mTxt04 MainTextFormatTest04) TimerEventText01() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TimerEventText01()",
		"")

	var errTxt string

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n"+
				"%v\n",
			"Error - time.LoadLocation()",
			err.Error())
		fmt.Println(errTxt)
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	elapsedTime := (time.Microsecond * 5) + 999
	endTime := startTime.Add(elapsedTime)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var timerLines *strmech.TextLineSpecTimerLines

	timerLines,
		err = strmech.TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	fmt.Println("Time Lines Output")
	fmt.Println()
	fmt.Printf(timerLines.String())

	/*
		  Start Time: 2022-04-05 10:00:00.000000000 -0500 CDT
		    End Time: 2022-04-05 10:00:00.000005999 -0500 CDT
		Elapsed Time: 5 Microseconds 999 Nanoseconds
		              Total Elapsed Nanoseconds: 5,999
	*/

}

func (mTxt04 MainTextFormatTest04) TimerEventText02() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TimerEventText01()",
		"")

	var errTxt string

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n"+
				"%v\n",
			"Error - time.LoadLocation()",
			err.Error())
		fmt.Println(errTxt)
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	elapsedTime := (time.Microsecond * 5) + 999
	endTime := startTime.Add(elapsedTime)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	timerLines,
		err := strmech.TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	sMech := strmech.StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(timerLines.String()),
			true)

	fmt.Println("Printable Time Lines Output")
	fmt.Println()
	fmt.Printf(printableExpectedStr)

	/*
		"[SPACE][SPACE]Start[SPACE]Time:[SPACE]2022-04-05[SPACE]10:00:00.000000000[SPACE]-0500[SPACE]CDT\n[SPACE][SPACE][SPACE][SPACE]End[SPACE]Time:[SPACE]2022-04-05[SPACE]10:00:00.000005999[SPACE]-0500[SPACE]CDT\nElapsed[SPACE]Time:[SPACE]5[SPACE]Microseconds[SPACE]999[SPACE]Nanoseconds\n[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Total[SPACE]Elapsed[SPACE]Nanoseconds:[SPACE]5,999\n"
	*/

}

func (mTxt04 MainTextFormatTest04) TimerEventText03() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TimerEventText01()",
		"")

	var errTxt string

	loc,
		err := time.LoadLocation(
		"America/Los_Angeles")

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n"+
				"%v\n",
			"Error - time.LoadLocation()",
			err.Error())
		fmt.Println(errTxt)
		return
	}

	startTime := time.Date(
		2021,
		4,
		20,
		18,
		30,
		0,
		0,
		loc)

	elapsedTime := (time.Millisecond * 200) +
		(time.Microsecond * 5) + 355

	endTime := startTime.Add(elapsedTime)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var timerLines *strmech.TextLineSpecTimerLines

	timerLines,
		err = strmech.TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	sMech := strmech.StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(timerLines.String()),
			true)

	fmt.Println("Printable Time Lines Output")
	fmt.Println()
	fmt.Printf(printableExpectedStr)

	/*
		"[SPACE][SPACE]Start[SPACE]Time:[SPACE]2021-04-20[SPACE]18:30:00.000000000[SPACE]-0700[SPACE]PDT\n[SPACE][SPACE][SPACE][SPACE]End[SPACE]Time:[SPACE]2021-04-20[SPACE]18:30:00.200005355[SPACE]-0700[SPACE]PDT\nElapsed[SPACE]Time:[SPACE]200[SPACE]Milliseconds[SPACE]5[SPACE]Microseconds[SPACE]355[SPACE]Nanoseconds\n[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Total[SPACE]Elapsed[SPACE]Nanoseconds:[SPACE]200,005,355\n"
	*/

}

func (mTxt04 MainTextFormatTest04) TimerEventText04() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.DateTimeTextField04()",
		"")

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var err error
	var timerLines01, timerLines02 *strmech.TextLineSpecTimerLines

	var errTxt, outputStr01, outputStr02 string

	tUtil := TestUtilities{}

	outputStr01,
		timerLines01,
		err =
		tUtil.CreateTestTextLineSpecTimerLines01(
			ePrefix.XCpy(
				"timerLines01"))

	if err != nil {
		errTxt = fmt.Sprintf(
			"\n%v - ERROR\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		errTxt = fmt.Sprintf(
			"\n%v - ERROR\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	outputStr02,
		timerLines02,
		err =
		tUtil.CreateTestTextLineSpecTimerLines01(
			ePrefix.XCpy(
				"timerLines01"))

	err = timerLines02.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		errTxt = fmt.Sprintf(
			"\n%v - ERROR\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	if outputStr01 != outputStr02 {
		errTxt = fmt.Sprintf(
			"\n%v - ERROR\n"+
				"Expected outputStr01 == outputStr02\n"+
				"HOWEVER, THEY ARE NOT EQUAL!\n"+
				"outputStr01 = '%v'\n"+
				"outputStr02 = '%v'\n",
			ePrefix.String(),
			outputStr01,
			outputStr02)
		fmt.Println(errTxt)
		return
	}

	timeFtmStr := "2006-01-02 15:04:05.000000000 -0700 MST"

	startTime01 := timerLines01.GetStartTime()
	startTime01Str := startTime01.Format(timeFtmStr)

	startTime02 := timerLines02.GetStartTime()
	startTime02Str := startTime02.Format(timeFtmStr)

	fmt.Println()
	fmt.Println("---------------------------------")
	errTxt = fmt.Sprintf(
		"StartTime01: %v\n"+
			"StartTime02: %v\n",
		startTime01Str,
		startTime02Str)

	fmt.Println(errTxt)
	fmt.Println("---------------------------------")
	fmt.Println()

	if !startTime01.Equal(startTime02) {

		errTxt = fmt.Sprintf(
			"\n%v - ERROR\n"+
				"startTime01.Equal(startTime02)\n"+
				"Expected timerLines01 == timerLines02\n"+
				"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		fmt.Println(errTxt)
		return

	} else {

		fmt.Println()
		fmt.Println("---------------------------------")
		fmt.Println("startTime01 and startTime02 ARE equal!")
		fmt.Println("---------------------------------")
		fmt.Println()

	}

	if !timerLines01.Equal(timerLines02) {

		errTxt = fmt.Sprintf(
			"\n%v - ERROR\n"+
				"timerLines01.Equal(timerLines02)\n"+
				"Expected timerLines01 == timerLines02\n"+
				"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())

		fmt.Println(errTxt)
		return
	}

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println("**    Successful Completion    **")
	fmt.Println("---------------------------------")
	fmt.Println()

}

func (mTxt04 MainTextFormatTest04) TimerEventText05() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TimerEventText05()",
		"")

	var errTxt string

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n"+
				"%v\n",
			"Error - time.LoadLocation()",
			err.Error())
		fmt.Println(errTxt)
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := time.Date(
		2022,
		5,
		5,
		10,
		30,
		45,
		9582,
		loc)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var timerLines *strmech.TextLineSpecTimerLines

	timerLines,
		err = strmech.TextLineSpecTimerLines{}.NewFullTimerEvent(
		" ",
		"Start Time",
		startTime,
		"End Time",
		endTime,
		"2006-01-02 15:04:05.000000000 -0700 MST",
		"Elapsed Time",
		len("Elapsed Time"),
		strmech.TxtJustify.Left(),
		": ",
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	fmt.Println("Time Lines Output")
	fmt.Println()
	fmt.Printf(timerLines.String())

	/*

	   Start Time  : 2022-04-05 10:00:00.000000000 -0500 CDT
	   End Time    : 2022-05-05 10:30:45.000009582 -0500 CDT
	   Elapsed Time: 30 Days 0 Hours 30 Minutes 45 Seconds 0 Milliseconds
	                 9 Microseconds 582 Nanoseconds
	                 Total Elapsed Nanoseconds: 2,593,845,000,009,582
	*/

}

func (mTxt04 MainTextFormatTest04) TimerEventText06() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TimerEventText06()",
		"")

	var errTxt string

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n"+
				"%v\n",
			"Error - time.LoadLocation()",
			err.Error())
		fmt.Println(errTxt)
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := time.Date(
		2022,
		5,
		5,
		10,
		30,
		45,
		999999998,
		loc)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var timerLines *strmech.TextLineSpecTimerLines

	timerLines,
		err = strmech.TextLineSpecTimerLines{}.NewFullTimerEvent(
		" ",
		"Start Time",
		startTime,
		"End Time",
		endTime,
		"2006-01-02 15:04:05.000000000 -0700 MST",
		"Elapsed Time",
		6,
		strmech.TxtJustify.Left(),
		": ",
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	fmt.Println("Time Lines Output")
	fmt.Println()
	fmt.Printf(timerLines.String())

	/*

	   Start Time  : 2022-04-05 10:00:00.000000000 -0500 CDT
	   End Time    : 2022-05-05 10:30:45.000009582 -0500 CDT
	   Elapsed Time: 30 Days 0 Hours 30 Minutes 45 Seconds 0 Milliseconds
	                 9 Microseconds 582 Nanoseconds
	                 Total Elapsed Nanoseconds: 2,593,845,000,009,582
	*/

}

func (mTxt04 MainTextFormatTest04) TextLineSpecTimerLinesSetEndTimeLabel07() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TextLineSpecTimerLinesSetEndTimeLabel07()",
		"")

	var errTxt string

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n"+
				"%v\n",
			"Error - time.LoadLocation()",
			err.Error())
		fmt.Println(errTxt)
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	elapsedTime := (time.Microsecond * 5) + 999
	endTime := startTime.Add(elapsedTime)

	fmt.Println()
	fmt.Println("---------------------------------")
	fmt.Println(ePrefix.String())
	fmt.Println("---------------------------------")
	fmt.Println()

	var timerLines *strmech.TextLineSpecTimerLines

	timerLines,
		err = strmech.TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	maximumTextLabelLength :=
		timerLines.GetMaximumTextLabelLength()

	totalTextLabelLength :=
		timerLines.GetLengthTotalLabel()

	adjustedTotalTextLabelLength :=
		totalTextLabelLength -
			timerLines.GetLengthLongestTextLabel()

	newEndTimeLabelLength :=
		maximumTextLabelLength - adjustedTotalTextLabelLength + 1

	newEndTimeLabel :=
		strings.Repeat("X", newEndTimeLabelLength)

	err = timerLines.SetEndTimeLabel(
		newEndTimeLabel,
		ePrefix.XCpy(
			"timerLines"))

	if err == nil {

		errTxt = fmt.Sprintf(
			"\n%v\n"+
				"ERROR timerLines.SetEndTimeLabel()\n"+
				"Expected an error return because input\n"+
				"parameter 'endTimeLabel' exceeds maximum length!\n"+
				"HOWEVER NO ERROR WAS RETURNED!\n"+
				"Maximum Text Label Length   = '%v'\n"+
				"Original Total Label Length = '%v'\n"+
				"New End Time Label Length   = '%v'\n",
			ePrefix.String(),
			maximumTextLabelLength,
			totalTextLabelLength,
			len(newEndTimeLabel))

		fmt.Println(errTxt)

		return
	}

	fmt.Println()
	fmt.Println("    SUCCESSFUL COMPLETION!")
	fmt.Println("=================================")
	fmt.Println(ePrefix.String())
	fmt.Println("=================================")
	fmt.Println()

}

func (mTxt04 MainTextFormatTest04) TimerEventText08() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TimerEventText08()",
		"")

	var errTxt string

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n"+
				"%v\n",
			"Error - time.LoadLocation()",
			err.Error())
		fmt.Println(errTxt)
		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := time.Date(
		2022,
		5,
		5,
		10,
		30,
		45,
		999999998,
		loc)

	var lineBreakSpec strmech.TextLineSpecSolidLine

	lineBreakSpec,
		err = strmech.TextLineSpecSolidLine{}.NewSolidLine(
		"",
		"",
		"-",
		75,
		ePrefix.XCpy(
			"lineBreak"))

	lineBreakSpec.TurnAutoLineTerminationOff()

	lineBreak := lineBreakSpec.String()

	fmt.Println()
	fmt.Println(lineBreak)
	fmt.Println(ePrefix.String())
	fmt.Println(lineBreak)
	fmt.Println()

	var timerLines *strmech.TextLineSpecTimerLines

	timerLines,
		err = strmech.TextLineSpecTimerLines{}.NewFullTimerEvent(
		" ",
		"Start Time",
		startTime,
		"End Time",
		endTime,
		"2006-01-02 15:04:05.000000000 -0700 MST",
		"Elapsed Time",
		6,
		strmech.TxtJustify.Left(),
		": ",
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	timerLinesText := timerLines.String()

	sMech := strmech.StrMech{}

	var timerLineStrs []string
	numOfTxtLines := 0
	var remainderStr string
	eolDelimiters := []string{"\n"}

	timerLineStrs,
		numOfTxtLines,
		remainderStr,
		err = sMech.ExtractTextLines(
		timerLinesText,
		eolDelimiters,
		true,
		ePrefix.XCpy(
			"timerLineStrs"))

	if err != nil {
		errTxt = fmt.Sprintf(
			"%v\n", err.Error())
		fmt.Println(errTxt)
		return
	}

	if numOfTxtLines <= 0 {
		errTxt = fmt.Sprintf(
			"%v\n"+
				"sMech.ExtractTextLines() FAILED!\n"+
				"Error: 'numOfTxtLines' is zero.\n",
			ePrefix.String())
		return
	}

	printableRemainderStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(remainderStr),
			true)

	if len(remainderStr) != 0 {
		fmt.Println(lineBreak)
		fmt.Println()
		fmt.Printf("Remainder Strings:\n"+
			"%v\n",
			printableRemainderStr)
		fmt.Println()
		fmt.Println(lineBreak)
		fmt.Println()
	}

	fmt.Println(lineBreak)
	fmt.Println()
	fmt.Println("Timer Line Strings")
	fmt.Println()
	for i := 0; i < numOfTxtLines; i++ {
		fmt.Printf(timerLineStrs[i])
	}
	fmt.Println(lineBreak)

	fmt.Println()
	fmt.Println()
	fmt.Println(lineBreak)
	fmt.Println("SUCCESSFUL COMPLETION")
	fmt.Println(ePrefix.String())
	fmt.Println("TextLineSpecTimerLines Output")
	fmt.Println()
	fmt.Printf(timerLinesText)
	fmt.Println()
	fmt.Println(lineBreak)

	/*

	   Start Time  : 2022-04-05 10:00:00.000000000 -0500 CDT
	   End Time    : 2022-05-05 10:30:45.000009582 -0500 CDT
	   Elapsed Time: 30 Days 0 Hours 30 Minutes 45 Seconds 0 Milliseconds
	                 9 Microseconds 582 Nanoseconds
	                 Total Elapsed Nanoseconds: 2,593,845,000,009,582
	*/

}

func (mTxt04 MainTextFormatTest04) TimerEventText09() {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"MainTest.TimerEventText01()",
		"")

	var errTxt string

	var loc *time.Location

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {

		err2 := fmt.Errorf(
			"\n%v - ERROR\n"+
				"time.LoadLocation(\"America/Chicago\")"+
				"%v\n",
			ePrefix.String(),
			err.Error())

		errTxt = fmt.Sprintf("%v",
			err2.Error())

		fmt.Println(errTxt)

		return
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := startTime.Add((time.Microsecond * 5) + 999)

	var timerLines01 *strmech.TextLineSpecTimerLines

	timerLines01,
		err = strmech.TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {

		errTxt = fmt.Sprintf("%v",
			err.Error())

		fmt.Println(errTxt)

		return
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {

		errTxt = fmt.Sprintf("%v",
			err.Error())

		fmt.Println(errTxt)

		return
	}

	var expectedText string

	expectedText,
		err = timerLines01.GetFormattedText(
		ePrefix.XCpy(
			"expectedText<-timerLines01"))

	if err != nil {

		errTxt = fmt.Sprintf("%v",
			err.Error())

		fmt.Println(errTxt)

		return
	}

	fmt.Printf("\ntimerLines01 - Formatted Text\n")
	fmt.Printf("%v\n\n", expectedText)

	sMech := strmech.StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedText),
			true)

	fmt.Printf("\ntimerLines01 - Printable Text\n")
	fmt.Printf("%v\n\n",
		printableExpectedStr)

	fmt.Printf("\n\nSuccessful Completion!\n")
	fmt.Printf("%v\n\n", ePrefix.String())

}

type TestUtilities struct {
	input string
}

func (tUtil TestUtilities) CreateTestTextLineSpecTimerLines01(
	errorPrefix interface{}) (
	string,
	*strmech.TextLineSpecTimerLines,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var outputStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecStandardLine01()",
		"")

	if err != nil {
		return outputStr, &strmech.TextLineSpecTimerLines{}, err
	}

	var loc *time.Location

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		return outputStr, &strmech.TextLineSpecTimerLines{}, err
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := startTime.Add((time.Microsecond * 5) + 999)

	var timerLines01 *strmech.TextLineSpecTimerLines

	timerLines01,
		err = strmech.TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		return outputStr, &strmech.TextLineSpecTimerLines{}, err
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		return outputStr, &strmech.TextLineSpecTimerLines{}, err
	}

	/*
		"[SPACE][SPACE]Start[SPACE]Time:[SPACE]2022-04-05[SPACE]10:00:00.000000000[SPACE]-0500[SPACE]CDT\n[SPACE][SPACE][SPACE][SPACE]End[SPACE]Time:[SPACE]2022-04-05[SPACE]10:00:00.000005999[SPACE]-0500[SPACE]CDT\nElapsed[SPACE]Time:[SPACE]5[SPACE]Microseconds[SPACE]999[SPACE]Nanoseconds\n[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Total[SPACE]Elapsed[SPACE]Nanoseconds:[SPACE]5,999\n"
	*/

	outputStr =
		"[SPACE][SPACE]Start[SPACE]Time:[SPACE]2022-04-05[SPACE]" +
			"10:00:00.000000000[SPACE]-0500[SPACE]CDT\\n[SPACE][SPACE]" +
			"[SPACE][SPACE]End[SPACE]Time:[SPACE]2022-04-05[SPACE]" +
			"10:00:00.000005999[SPACE]-0500[SPACE]CDT\\nElapsed[SPACE]" +
			"Time:[SPACE]5[SPACE]Microseconds[SPACE]999[SPACE]" +
			"Nanoseconds\\n[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]" +
			"[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]" +
			"[SPACE]Total[SPACE]Elapsed[SPACE]" +
			"Nanoseconds:[SPACE]5,999\\n"

	return outputStr, timerLines01, err
}

func (tUtil TestUtilities) CreateTestTextLineSpecStandardLine01(
	errorPrefix interface{}) (
	strmech.TextLineSpecStandardLine,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error
	stdLine01 := strmech.TextLineSpecStandardLine{}.New()

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestUtilities."+
			"CreateTestTextLineSpecStandardLine01()",
		"")

	if err != nil {
		return stdLine01, err
	}

	rightMarginLen := 5
	var rightMarginSpec strmech.TextFieldSpecSpacer

	rightMarginSpec,
		err = strmech.TextFieldSpecSpacer{}.NewSpacer(
		rightMarginLen,
		ePrefix.XCpy(
			"rightMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	var leftMarginSpec strmech.TextFieldSpecSpacer

	leftMarginLen := 6

	leftMarginSpec,
		err = strmech.TextFieldSpecSpacer{}.NewSpacer(
		leftMarginLen,
		ePrefix.XCpy(
			"leftMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	label := "How Now Brown Cow!"
	fieldLen := len(label) + 4
	txtJustify := strmech.TxtJustify.Center()

	var labelSpec strmech.TextFieldSpecLabel

	labelSpec,
		err = strmech.TextFieldSpecLabel{}.NewTextLabel(
		label,
		fieldLen,
		txtJustify,
		ePrefix.XCpy(
			"labelSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 0
	_,
		err = stdLine01.AddTextField(
		&leftMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-leftMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 1

	_,
		err = stdLine01.AddTextField(
		&labelSpec,
		ePrefix.XCpy(
			"stdLine01<-labelSpec"))

	if err != nil {
		return stdLine01, err
	}

	//Index 2
	_,
		err = stdLine01.AddTextField(
		&rightMarginSpec,
		ePrefix.XCpy(
			"stdLine01<-rightMarginSpec"))

	if err != nil {
		return stdLine01, err
	}

	err = stdLine01.IsValidInstanceError(
		ePrefix.XCpy(
			"stdLine01"))

	if err != nil {
		return stdLine01, err
	}

	return stdLine01, err
}

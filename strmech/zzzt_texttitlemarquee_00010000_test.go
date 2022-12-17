package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
	"time"
)

func TestTextLineTitleMarqueeDto_AddTitleLineStrArrayDtos_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestNumStrMathQuarkExtendRunes_000100()",
		"")

	marqueeDto := TextLineTitleMarqueeDto{
		StandardTitleLeftMargin:   " *",
		StandardTitleRightMargin:  "*",
		StandardMaxLineLen:        80,
		StandardTextFieldLen:      99,
		StandardTextJustification: TxtJustify.Center(),
		NumLeadingBlankLines:      1,
		LeadingSolidLineChar:      "*",
		NumLeadingSolidLines:      2,
		NumTopTitleBlankLines:     1,
		TitleLines:                TextLineSpecLinesCollection{},
		NumBottomTitleBlankLines:  1,
		TrailingSolidLineChar:     "*",
		NumTrailingSolidLines:     2,
		NumTrailingBlankLines:     1,
	}

	strArrayDto := StringArrayDto{}

	strArrayDto.AddString("Title Line #1")

	strArrayDto.AddString("Title Line #2")

	strArrayDto.AddString("Title Line #3")

	err := marqueeDto.AddTitleLineStrArrayDto(
		strArrayDto,
		ePrefix.XCpy("strArrayDto"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

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
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var titleLineMarquee TextLineSpecTitleMarquee

	titleLineMarquee,
		err = new(TextLineSpecTitleMarquee).
		NewMarqueeDto(
			&marqueeDto,
			ePrefix.XCpy(
				"titleLineMarquee<-marqueeDto"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var marqueeStr string

	marqueeStr,
		err = titleLineMarquee.GetFormattedText(
		ePrefix.XCpy("marqueeStr<-titleLineMarquee"))

	sMech := StrMech{}

	convetedMarqueeStr := sMech.ConvertNonPrintableChars(
		[]rune(marqueeStr),
		true)

	expectedMarqueeStr :=
		"\\n[SPACE]******************************************************************************\\n[SPACE]******************************************************************************\\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Title[SPACE]Line[SPACE]#1[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Title[SPACE]Line[SPACE]#2[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Title[SPACE]Line[SPACE]#3[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]2022-12-16[SPACE]14:19:03.000000000[SPACE]-0600[SPACE]CST[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\\n[SPACE]*[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]*\\n[SPACE]******************************************************************************\\n[SPACE]******************************************************************************\\n\\n"

	if convetedMarqueeStr != expectedMarqueeStr {
		t.Errorf("\n\n%v\n"+
			"Error: Expected Marquee String is NOT\n"+
			"equal to Actual Marquee String!\n"+
			"Converted Expected Marquee String\n%v\n\n"+
			"Converted Actual Marquee String\n%v\n\n"+
			"Readable Actual Marquee String\n%v\n\n",
			ePrefix.String(),
			expectedMarqueeStr,
			convetedMarqueeStr,
			marqueeStr)

		return
	}

}

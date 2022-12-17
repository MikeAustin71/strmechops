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

func (MainTextFormatTest04) TxtLineMarquee001() {
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

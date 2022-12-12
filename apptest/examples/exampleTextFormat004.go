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
		StandardTitleLeftMargin:   "",
		StandardTitleRightMargin:  "",
		StandardMaxLineLen:        80,
		StandardTextFieldLen:      70,
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

	err = marqueeDto.AddTitleLineDateTimeStr(
		time.Now(),
		"",
		ePrefix.XCpy(
			"Time Stamp"))

	if err != nil {

		fmt.Printf("\n\n%v\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return

	}

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

	fmt.Printf("\n\n\n" + breakStr + "\n")

	fmt.Printf("\n Successful Completion!\n"+
		"Function: %v\n",
		ePrefix.String())

	fmt.Printf("\n" + breakStr + "\n")

}

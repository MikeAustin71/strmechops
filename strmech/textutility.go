package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"sync"
	"time"
)

type TextUtility struct {
	lock *sync.Mutex
}

func (txtUtility *TextUtility) BuildOneColTopMarquee(
	strBuilder *strings.Builder,
	maxLineLen int,
	title1 string,
	turnAutoLineLengthBreaksOn bool,
	txtFmtParams TextFmtParamsLineColumnsDto,
	errorPrefix interface{}) (
	err error) {

	if txtUtility.lock == nil {
		txtUtility.lock = new(sync.Mutex)
	}

	txtUtility.lock.Lock()

	defer txtUtility.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TextUtility."+
			"BuildOneColTopMarquee()",
		"")

	if err != nil {
		return err
	}

	if len(title1) == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'title1' is a zero length\n"+
			"Empty string!\n",
			ePrefix.String())

		return err
	}

	txtFormatCol := TextFormatterCollection{}

	if txtFmtParams.GetNumOfFieldFmtParams() != 1 {
		// txtFmtParams is invalid,
		// apply default standard format parameters
		// for 1-Column text displays.

		txtFmtParams.FieldFormatParams = nil

		txtFmtParams := TextFmtParamsLineColumnsDto{
			FormatType: TxtFieldType.LineColumns(),
			FieldFormatParams: []TextFieldFmtParamsDto{
				{
					LeftMarginStr:  " ",
					FieldLength:    maxLineLen,
					FieldJustify:   TxtJustify.Center(),
					DateTimeFormat: string(new(textSpecificationAtom).getDefaultNewLineChars()),
					RightMarginStr: "",
					lock:           nil,
				},
			},
			TurnLineTerminationOff:     false,
			LineTerminator:             "\n",
			MaxLineLength:              maxLineLen,
			TurnAutoLineLengthBreaksOn: turnAutoLineLengthBreaksOn,
			lock:                       nil,
		}

		err = txtFormatCol.
			SetStdFormatParams(
				txtFmtParams,
				ePrefix.XCpy(
					"Default txtFmtParams 1-Column"))

		if err != nil {
			return err
		}

	}
	// Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// Filler =======
	// Marquee Top
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		maxLineLen-2,
		" ",
		false,
		"",
		-1,
		false)

	// Title Line 1
	err = txtFormatCol.AddLine1Col(
		"CharSearchNumStrParseResultsDto",
		ePrefix.XCpy(
			"Title Line 1"))

	if err != nil {
		return err
	}
	// Title Line  4 Date/Time

	err = txtFormatCol.AddLine1Col(
		TextInputParamFieldDateTimeDto{
			FieldDateTime:       time.Now(),
			FieldDateTimeFormat: "Monday 2006-01-02 15:04:05.000000000 -0700 MST",
		},
		ePrefix.XCpy(
			"Top-Title Line 4"))

	if err != nil {
		return err
	}
	// Filler Line '========='
	// Marquee Bottom
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		maxLineLen-2,
		" ",
		false,
		"",
		-1,
		false)

	// Trailing Blank Line
	txtFormatCol.AddLineBlank(
		1,
		"")

	// End Of Marquee

	err = txtFormatCol.BuildText(
		strBuilder,
		ePrefix.XCpy(
			"strBuilder<-txtFormatCol"))

	return err
}

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

func (txtUtility *TextUtility) BuildOneColLeadingMarquee(
	strBuilder *strings.Builder,
	titles StringArrayDto,
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
			"BuildOneColLeadingMarquee()",
		"")

	if err != nil {
		return err
	}

	lenTitles := titles.GetStringArrayLength()

	if lenTitles == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titles' is a zero length\n"+
			"string array!\n",
			ePrefix.String())

		return err
	}

	if txtFmtParams.MaxLineLength < 5 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtFmtParams.MaxLineLength' is invalid!\n"+
			"The minimum value for 'txtFmtParams.MaxLineLength' is five (5).\n"+
			"maxLineLen = '%v'\n",
			ePrefix.String(),
			txtFmtParams.MaxLineLength)

		return err
	}

	cntEmptyStrs := 0

	for i := 0; i < lenTitles; i++ {
		if len(titles.StrArray[i]) == 0 {
			cntEmptyStrs++
		}
	}

	if cntEmptyStrs == lenTitles {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titles' is invalid!\n"+
			"All the strings in string array 'titles' are\n"+
			"zero length or empty strings.\n",
			ePrefix.String())

		return err

	}

	txtFormatCol := TextFormatterCollection{}

	if txtFmtParams.GetNumOfFieldFmtParams() != 1 {
		// txtFmtParams is invalid,
		// apply default standard format parameters
		// for 1-Column text displays.

		txtFmtParams.FieldFormatParams = nil

		maxLineLen := txtFmtParams.MaxLineLength
		turnAutoLineLengthBreaksOn :=
			txtFmtParams.TurnAutoLineLengthBreaksOn

		txtFmtParams = TextFmtParamsLineColumnsDto{
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

	}

	err = txtFormatCol.
		SetStdFormatParams(
			txtFmtParams,
			ePrefix.XCpy(
				"Default txtFmtParams 1-Column"))

	if err != nil {
		return err
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
		txtFmtParams.MaxLineLength-2,
		" ",
		false,
		"",
		-1,
		false)

	// Title Line 1
	for i := 0; i < lenTitles; i++ {

		title := titles.StrArray[i]

		if len(title) > 0 {

			err = txtFormatCol.AddLine1Col(
				title,
				ePrefix.XCpy(
					fmt.Sprintf(
						"titles.StrArray[i]")))

			if err != nil {
				return err
			}

		}

	}
	// Last Line Date/Time

	err = txtFormatCol.AddLine1Col(
		TextInputParamFieldDateTimeDto{
			FieldDateTime:       time.Now(),
			FieldDateTimeFormat: txtFmtParams.FieldFormatParams[0].DateTimeFormat,
		},
		ePrefix.XCpy(
			"Last Line Date/Time"))

	if err != nil {
		return err
	}
	// Filler Line '========='
	// Marquee Bottom
	txtFormatCol.AddLineSolid(
		" ",
		"=",
		txtFmtParams.MaxLineLength-2,
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

func (txtUtility *TextUtility) BuildOneColTrailingMarquee(
	strBuilder *strings.Builder,
	titles StringArrayDto,
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
			"BuildOneColTrailingMarquee()",
		"")

	if err != nil {
		return err
	}

	lenTitles := titles.GetStringArrayLength()

	if lenTitles == 0 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titles' is a zero length\n"+
			"string array!\n",
			ePrefix.String())

		return err
	}

	turnAutoLineLengthBreaksOn :=
		txtFmtParams.TurnAutoLineLengthBreaksOn

	maxLineLen :=
		txtFmtParams.MaxLineLength

	if maxLineLen < 5 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLen' is invalid!\n"+
			"The minimum value for 'maxLineLen' is five (5).\n"+
			"maxLineLen = '%v'\n",
			ePrefix.String(),
			maxLineLen)

		return err
	}

	cntEmptyStrs := 0

	for i := 0; i < lenTitles; i++ {
		if len(titles.StrArray[i]) == 0 {
			cntEmptyStrs++
		}
	}

	if cntEmptyStrs == lenTitles {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'titles' is invalid!\n"+
			"All the strings in string array 'titles' are\n"+
			"zero length or empty strings.\n",
			ePrefix.String())

		return err

	}

	txtFormatCol := TextFormatterCollection{}

	if txtFmtParams.GetNumOfFieldFmtParams() != 1 {
		// txtFmtParams is invalid,
		// apply default standard format parameters
		// for 1-Column text displays.

		txtFmtParams.FieldFormatParams = nil

		txtFmtParams = TextFmtParamsLineColumnsDto{
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
	}

	err = txtFormatCol.
		SetStdFormatParams(
			txtFmtParams,
			ePrefix.XCpy(
				"Default txtFmtParams 1-Column"))

	if err != nil {
		return err
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
	for i := 0; i < lenTitles; i++ {

		title := titles.StrArray[i]

		if len(title) > 0 {

			err = txtFormatCol.AddLine1Col(
				title,
				ePrefix.XCpy(
					fmt.Sprintf(
						"titles.StrArray[i]")))

			if err != nil {
				return err
			}

		}

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

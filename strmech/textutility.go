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

// BuildOneColLeadingMarquee - Builds a one-column top title
// display or 'Marquee' and writes the formatted text to a
// strings.Builder instance passed as an input parameter.
//
// Be advised. This is the top half of a title Marquee. The bottom
// half of the title marquee is configured with a call to method:
//
//	TextUtility.BuildOneColTrailingMarquee()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 strBuilder                 *strings.Builder
//	    - A pointer to an instance of strings.Builder. The formatted
//	      text created by this method will be written to this
//	      instance of string builder.
//
//	 titles                     StringArrayDto
//	    - This String Array Data Transfer Object encapsulates a
//	      string array which will supply the Marquee titles. Each
//	      title will be treated as a 1-Column display and placed
//	      on a separate text line.
//
//	      If the first element in this array is an empty string,
//	      an error will be returned. Empty strings in the remaining
//	      array elements will be skipped and ignored.
//
//	 txtFmtParams               TextFmtParamsLineColumnsDto
//	    - This instance of TextFmtParamsLineColumnsDto should
//	      be configured for a 1-Column Display. The format
//	      parameters contained in this instance will be used
//	      to format the text for all title lines formatted for
//	      text output.
//
//	      If this instance of TextFmtParamsLineColumnsDto is invalid,
//	      default values will be used as shown below:
//
//	      txtFmtParams = TextFmtParamsLineColumnsDto{
//	        FormatType: TxtFieldType.LineColumns(),
//	        FieldFormatParams: []TextFieldFmtParamsDto{
//	           {
//	            LeftMarginStr:  " ",
//	            FieldLength:    79,
//	            FieldJustify:   TxtJustify.Center(),
//	            DateTimeFormat: "2006-01-02 15:04:05.000000000 -0700 MST",
//	            RightMarginStr: "",
//	           },
//	        },
//	        TurnLineTerminationOff:     false,
//	        LineTerminator:             "\n",
//	        MaxLineLength:              79,
//	        TurnAutoLineLengthBreaksOn: true,
//	      }
//
//
//		errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//		     parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//
//		     1. nil - A nil value is valid and generates an empty
//		              collection of error prefix and error context
//		              information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		                 error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//		                       ErrorPrefixInfo from this object will be
//		                       copied to 'errPrefDto'.
//
//		     6. *ErrPrefixDto - A pointer to an instance of
//		                        ErrPrefixDto. ErrorPrefixInfo from this
//		                        object will be copied to 'errPrefDto'.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		                            a two-dimensional slice of strings
//		                            containing error prefix and error
//		                            context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

		txtFmtParams = TextFmtParamsLineColumnsDto{
			FormatType: TxtFieldType.LineColumns(),
			FieldFormatParams: []TextFieldFmtParamsDto{
				{
					LeftMarginStr: " ",
					FieldLength:   79,
					FieldJustify:  TxtJustify.Center(),
					DateTimeFormat: new(textSpecificationMolecule).
						getDefaultDateTimeFormat(),
					RightMarginStr: "",
					lock:           nil,
				},
			},
			TurnLineTerminationOff:     false,
			LineTerminator:             "\n",
			MaxLineLength:              79,
			TurnAutoLineLengthBreaksOn: true,
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

// BuildOneColTrailingMarquee - Builds a one-column bottom title
// display or 'Marquee' and writes the formatted text to a
// strings.Builder instance passed as an input parameter.
//
// Be advised. This is the bottom half of a title Marquee. The top
// half of the title marquee is configured with a call to method:
//
//	TextUtility.BuildOneColLeadingMarquee()
//
// ----------------------------------------------------------------
//
// Input Parameters
//
//	 strBuilder                 *strings.Builder
//	    - A pointer to an instance of strings.Builder. The formatted
//	      text created by this method will be written to this
//	      instance of string builder.
//
//	 titles                     StringArrayDto
//	    - This String Array Data Transfer Object encapsulates a
//	      string array which will supply the Marquee titles. Each
//	      title will be treated as a 1-Column display and placed
//	      on a separate text line.
//
//	      If the first element in this array is an empty string,
//	      an error will be returned. Empty strings in the remaining
//	      array elements will be skipped and ignored.
//
//	 txtFmtParams               TextFmtParamsLineColumnsDto
//	    - This instance of TextFmtParamsLineColumnsDto should
//	      be configured for a 1-Column Display. The format
//	      parameters contained in this instance will be used
//	      to format the text for all title lines formatted for
//	      text output.
//
//	      If this instance of TextFmtParamsLineColumnsDto is invalid,
//	      default values will be used as shown below:
//
//	      txtFmtParams = TextFmtParamsLineColumnsDto{
//	        FormatType: TxtFieldType.LineColumns(),
//	        FieldFormatParams: []TextFieldFmtParamsDto{
//	           {
//	            LeftMarginStr:  " ",
//	            FieldLength:    79,
//	            FieldJustify:   TxtJustify.Center(),
//	            DateTimeFormat: "2006-01-02 15:04:05.000000000 -0700 MST",
//	            RightMarginStr: "",
//	           },
//	        },
//	        TurnLineTerminationOff:     false,
//	        LineTerminator:             "\n",
//	        MaxLineLength:              79,
//	        TurnAutoLineLengthBreaksOn: true,
//	      }
//
//
//		errorPrefix                interface{}
//		   - This object encapsulates error prefix text which is
//		     included in all returned error messages. Usually, it
//		     contains the name of the calling method or methods
//		     listed as a method or function chain of execution.
//
//		     If no error prefix information is needed, set this
//		     parameter to 'nil'.
//
//		     This empty interface must be convertible to one of the
//		     following types:
//
//
//		     1. nil - A nil value is valid and generates an empty
//		              collection of error prefix and error context
//		              information.
//
//		     2. string - A string containing error prefix information.
//
//		     3. []string A one-dimensional slice of strings containing
//		                 error prefix information
//
//		     4. [][2]string A two-dimensional slice of strings
//		        containing error prefix and error context information.
//
//		     5. ErrPrefixDto - An instance of ErrPrefixDto. The
//		                       ErrorPrefixInfo from this object will be
//		                       copied to 'errPrefDto'.
//
//		     6. *ErrPrefixDto - A pointer to an instance of
//		                        ErrPrefixDto. ErrorPrefixInfo from this
//		                        object will be copied to 'errPrefDto'.
//
//		     7. IBasicErrorPrefix - An interface to a method generating
//		                            a two-dimensional slice of strings
//		                            containing error prefix and error
//		                            context information.
//
//		     If parameter 'errorPrefix' is NOT convertible to one of
//		     the valid types listed above, it will be considered
//		     invalid and trigger the return of an error.
//
//		     Types ErrPrefixDto and IBasicErrorPrefix are included in
//		     the 'errpref' software package,
//		     "github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// Return Values
//
//	error
//	   - If the method completes successfully and no errors are
//	     encountered this return value is set to 'nil'. Otherwise,
//	     if errors are encountered, this return value will contain
//	     an appropriate error message.
//
//	     If an error message is returned, the text value of input
//	     parameter 'errorPrefix' will be inserted or prefixed at
//	     the beginning of the error message.
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

		txtFmtParams = TextFmtParamsLineColumnsDto{
			FormatType: TxtFieldType.LineColumns(),
			FieldFormatParams: []TextFieldFmtParamsDto{
				{
					LeftMarginStr: " ",
					FieldLength:   79,
					FieldJustify:  TxtJustify.Center(),
					DateTimeFormat: new(textSpecificationMolecule).
						getDefaultDateTimeFormat(),
					RightMarginStr: "",
					lock:           nil,
				},
			},
			TurnLineTerminationOff:     false,
			LineTerminator:             "\n",
			MaxLineLength:              79,
			TurnAutoLineLengthBreaksOn: true,
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

	// Format Title Lines
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

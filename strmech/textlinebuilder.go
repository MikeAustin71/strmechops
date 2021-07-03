package strmech

import (
	"errors"
	"fmt"
	"strings"
)

type TextLineBuilder struct {
	Input  string
	Output string
}

func (txtBuilder TextLineBuilder) ArrayBuild(
	b *strings.Builder,
	specs []TextSpec,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.ArrayBuild() "

	return txtBuilder.Build(b, ePrefix, specs...)
}

func (txtBuilder TextLineBuilder) Build(
	b *strings.Builder,
	ePrefix string,
	specs ...TextSpec) error {

	ePrefix += "TextLineBuilder.Build() "
	var err error

	for i, txtSpec := range specs {

		switch t := txtSpec.(type) {

		case BlankLinesSpec:
			err = txtBuilder.CreateBlankLinesSpec(t, b, ePrefix)
		case IntegerSpec:
			err = txtBuilder.CreateIntegerSpec(t, b, ePrefix)
		case LineSpec:
			err = txtBuilder.CreateLineSpec(t, b, ePrefix)
		case MarginSpec:
			err = txtBuilder.CreateMarginSpec(t, b, ePrefix)
		case NewLineSpec:
			err = txtBuilder.CreateNewLineSpec(t, b, ePrefix)
		case StringSpec:
			err = txtBuilder.CreateStringSpec(t, b, ePrefix)
		case TextLineBreakField:
			err = txtBuilder.CreateLineBreakField(t, b, ePrefix)
		case TextNumericIntField:
			err = txtBuilder.CreateNumericIntField(t, b, ePrefix)
		case TextOneLabelOneIntField:
			err = txtBuilder.CreateOneLabelOneIntField(t, b, ePrefix)
		case TextStringField:
			err = txtBuilder.CreateStringField(t, b, ePrefix)
		case TextTwoLabelStrField:
			err = txtBuilder.CreateTwoLabelStrField(t, b, ePrefix)
		default:
			err = fmt.Errorf(ePrefix+
				"\nError: Text Specification at index #%v is invalid!\n", i)
		}
	}

	return err
}

func (txtBuilder TextLineBuilder) CenterInField(
	strValue string,
	strFieldLen int,
	padChar rune,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CenterInField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix+
			"\nError: Input parameter 'strValue' length exceeds field length parameter, 'strFieldLen'.\n"+
			"strValue='%v'\n"+
			"strValue length='%v'\n"+
			"strFieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	if lenStr == strFieldLen {

		_, err = b.WriteString(strValue)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteString(strValue)\n"+
				"strValue string length == strFieldLen\n"+
				"strValue='%v'\n"+
				"Error='%v'\n", strValue, err.Error())
		}
	}

	if padChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'padChar' is ZERO!\n")
	}

	grossPad := strFieldLen - lenStr
	var leftPad, rightPad int

	// Pad example.
	// Length of pad = 7
	// 7/2 = 3
	// Left Pad = 3     Right Pad = 4
	// Extra 1-Pad char added to Right Pad
	leftPad = grossPad / 2
	rightPad = grossPad - leftPad

	for i := 0; i < leftPad; i++ {

		_, err = b.WriteRune(padChar)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteRune(left padChar)\n"+
				"padChar='%v'   i='%v'\n"+
				"Error='%v'\n", padChar, i, err.Error())
		}
	}

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned after left pad by b.WriteString(strValue)\n"+
			"strValue='%v'\n"+
			"Error='%v'\n", strValue, err.Error())
	}

	for j := 0; j < rightPad; j++ {

		_, err = b.WriteRune(padChar)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteRune(right padChar)\n"+
				"padChar='%v'   j='%v'\n"+
				"Error='%v'\n", padChar, j, err.Error())
		}
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateBlankLinesSpec(
	blankLines BlankLinesSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateBlankLinesSpec() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	if blankLines.NumBlankLines < 1 {
		return nil
	}

	var err error

	for i := 0; i < blankLines.NumBlankLines; i++ {

		_, err = b.WriteRune('\n')

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteRune(newLine).\n"+
				"i='%v'\n"+
				"Error='%v'\n", i, err.Error())
		}
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateIntegerSpec(
	intTxtSpec IntegerSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateIntegerSpec() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	numFmt := ""

	if len(intTxtSpec.NumericFieldSpec) == 0 {
		numFmt = "%d"
	} else {
		numFmt = intTxtSpec.NumericFieldSpec
	}

	numStr := fmt.Sprintf(numFmt, intTxtSpec.NumericValue)

	lenNumStr := len(numStr)

	if lenNumStr > intTxtSpec.NumericFieldLength {
		return fmt.Errorf(ePrefix+
			"\nError: Length of number string is greater than NumericFieldLength!\n"+
			"Number String='%v'\n"+
			"Number String Length='%v'\n"+
			"Numeric Field Length='%v'\n",
			numStr,
			lenNumStr,
			intTxtSpec.NumericFieldLength)
	}

	var err error

	if lenNumStr == intTxtSpec.NumericFieldLength {

		_, err = b.WriteString(numStr)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteString(numStr).\n"+
				"numStr='%v'\n"+
				"Error='%v'\n", numStr, err.Error())
		}
	}

	switch intTxtSpec.NumericPosition {
	case FieldPos.LeftJustify():

		err = txtBuilder.LeftJustifyField(
			numStr,
			intTxtSpec.NumericFieldLength,
			intTxtSpec.NumericPadChar,
			b,
			ePrefix)

	case FieldPos.RightJustify():

		err = txtBuilder.RightJustifyField(
			numStr,
			intTxtSpec.NumericFieldLength,
			intTxtSpec.NumericPadChar,
			b,
			ePrefix)

	case FieldPos.Center():

		err = txtBuilder.CenterInField(
			numStr,
			intTxtSpec.NumericFieldLength,
			intTxtSpec.NumericPadChar,
			b,
			ePrefix)

	default:

		err = errors.New(ePrefix +
			"\nError: Invalid parameter intTxtSpec.NumericPosition\n")
	}

	if err != nil {
		return err
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateLineBreakField(
	lineBreak TextLineBreakField,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateLineBreakField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	if !lineBreak.CreateLineBreak {
		// Nothing to do
		return nil
	}

	var err error

	err = txtBuilder.CreateBlankLinesSpec(
		lineBreak.LeadingBlankLines,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginSpec(
		lineBreak.LeftMargin,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginSpec(
		lineBreak.LeftSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineSpec(lineBreak.LineSpec, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginSpec(
		lineBreak.RightSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateNewLineSpec(
		lineBreak.TerminateWithNewLine,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateBlankLinesSpec(
		lineBreak.TrailingBlankLines,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	return nil
}

// CreateLineSpec - Creates a string from the Create Line
// Specification and writes it to a strings.Builder.
func (txtBuilder TextLineBuilder) CreateLineSpec(
	lineSpec LineSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateLineSpec() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!\n")
	}

	if lineSpec.LineLength < 1 {
		return nil
	}

	if lineSpec.LineChar == 0 {
		return nil
	}

	if lineSpec.LineFieldLength < lineSpec.LineLength {
		return fmt.Errorf(ePrefix+
			"\nError: lineSpec.LineFieldLength IS LESS THAN lineSpec.LineLength!\n"+
			"lineSpec.LineFieldLength='%v'   lineSpec.LineLength='%v'\n",
			lineSpec.LineFieldLength, lineSpec.LineLength)
	}

	xb := strings.Builder{}

	xb.Grow(lineSpec.LineLength + 5)

	var err error

	for i := 0; i < lineSpec.LineLength; i++ {

		_, err = xb.WriteRune(lineSpec.LineChar)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by xb.WriteRune(lineSpec.LineChar).\n"+
				"lineSpec.LineChar='%v'\n"+
				"Error='%v'\n", lineSpec.LineChar, err.Error())
		}
	}

	if lineSpec.LineLength == lineSpec.LineFieldLength {

		_, err = b.WriteString(xb.String())

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteString(xb.String()).\n"+
				"lineSpec.LineChar='%v'\n"+
				"lineSpec.LineLength='%v'   lineSpec.LineFieldLength='%v'\n"+
				"Error='%v'\n",
				lineSpec.LineChar,
				lineSpec.LineLength,
				lineSpec.LineFieldLength,
				err.Error())
		}
	}

	if lineSpec.LineFieldPadChar == 0 {
		return errors.New(ePrefix +
			"\nError: lineSpec.LineFieldPadChar is ZERO!\n")
	}

	switch lineSpec.LinePosition {

	case FieldPos.LeftJustify():

		err = txtBuilder.LeftJustifyField(
			xb.String(),
			lineSpec.LineFieldLength,
			lineSpec.LineFieldPadChar,
			b,
			ePrefix)

	case FieldPos.RightJustify():

		err = txtBuilder.RightJustifyField(
			xb.String(),
			lineSpec.LineFieldLength,
			lineSpec.LineFieldPadChar,
			b,
			ePrefix)

	case FieldPos.Center():

		err = txtBuilder.CenterInField(
			xb.String(),
			lineSpec.LineFieldLength,
			lineSpec.LineFieldPadChar,
			b,
			ePrefix)

	default:
		err = fmt.Errorf(ePrefix+
			"\nError: Input parameter lineSpec.LinePosition is invalid!\n"+
			"lineSpec.LinePosition value='%v'\n", lineSpec.LinePosition.UtilityValue())
	}

	return err
}

func (txtBuilder TextLineBuilder) CreateMarginSpec(
	margin MarginSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateMarginSpec() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	if len(margin.MarginStr) > 0 {
		_, err = b.WriteString(margin.MarginStr)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteString(margin.MarginStr)\n"+
				"margin.MarginStr='%v'\n"+
				"Error='%v'\n", margin.MarginStr, err.Error())
		}
	}

	if margin.MarginLength == 0 {
		// Nothing to do
		return nil
	}

	if margin.MarginChar == 0 {
		return errors.New(ePrefix +
			"\nError: margin.MarginChar = 0\n")
	}

	xb := strings.Builder{}
	xb.Grow(margin.MarginLength + 2)

	for i := 0; i < margin.MarginLength; i++ {
		_, err = xb.WriteRune(margin.MarginChar)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by xb.WriteRune(margin.MarginChar)\n"+
				"margin.MarginChar='%v'\n"+
				"Error='%v'\n", margin.MarginChar, err.Error())
		}
	}

	_, err = b.WriteString(xb.String())

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by b.WriteString(xb.String())\n"+
			"xb.String()='%v'\n"+
			"Error='%v'\n", xb.String(), err.Error())
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateNewLineSpec(
	newLineSpec NewLineSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateNewLineSpec() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	if !newLineSpec.AddNewLine {
		// Nothing to do
		return nil
	}

	var err error

	_, err = b.WriteRune('\n')

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by b.WriteRune(newline)\n"+
			"Error='%v'\n", err.Error())
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateNumericIntField(
	numSpec TextNumericIntField,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateNumericIntField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = txtBuilder.CreateMarginSpec(
		numSpec.LeftMargin,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginSpec(
		numSpec.LeftSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateIntegerSpec(numSpec.NumberSpec, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginSpec(
		numSpec.RightSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateNewLineSpec(numSpec.TerminateWithNewLine, b, ePrefix)

	return err
}

func (txtBuilder TextLineBuilder) CreateOneLabelOneIntField(
	oneLabelOneInt TextOneLabelOneIntField,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzLogOps.CreateOneLabelOneIntField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = txtBuilder.CreateBlankLinesSpec(oneLabelOneInt.LeadingBlankLines, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineBreakField(oneLabelOneInt.TopLineBreak, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateStringField(oneLabelOneInt.Label1, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateNumericIntField(oneLabelOneInt.Number1, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineBreakField(oneLabelOneInt.BottomLineBreak, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateBlankLinesSpec(oneLabelOneInt.TrailingBlankLines, b, ePrefix)

	if err != nil {
		return err
	}

	return nil
}

func (txtBuilder TextLineBuilder) CreateTwoLabelStrField(
	twoLabelStrLine TextTwoLabelStrField,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateTwoLabelStrField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = txtBuilder.CreateBlankLinesSpec(twoLabelStrLine.LeadingBlankLines, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineBreakField(twoLabelStrLine.TopLineBreak, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateStringField(twoLabelStrLine.Label1, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateStringField(twoLabelStrLine.Label2, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateLineBreakField(twoLabelStrLine.BottomLineBreak, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateBlankLinesSpec(twoLabelStrLine.TrailingBlankLines, b, ePrefix)

	if err != nil {
		return err
	}

	return nil
}

// CreateStringField - Designed to handle TextStringField specifications.
func (txtBuilder TextLineBuilder) CreateStringField(
	strField TextStringField,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.CreateStringField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	var err error

	err = txtBuilder.CreateMarginSpec(
		strField.LeftMargin,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginSpec(
		strField.LeftSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateStringSpec(strField.StrTxtSpec, b, ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateMarginSpec(
		strField.RightSpacer,
		b,
		ePrefix)

	if err != nil {
		return err
	}

	err = txtBuilder.CreateNewLineSpec(strField.TerminateWithNewLine, b, ePrefix)

	if err != nil {
		return err
	}

	return nil
}

// CreateStringSpec - Receives 'StringSpec' type and then proceeds
// to create the string and write it to the string builder.
func (txtBuilder TextLineBuilder) CreateStringSpec(
	strSpec StringSpec,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder) CreateStringSpec() "

	var err error

	switch strSpec.StrPosition {

	case FieldPos.LeftJustify():

		err = txtBuilder.LeftJustifyField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	case FieldPos.RightJustify():

		err = txtBuilder.RightJustifyField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	case FieldPos.Center():

		err = txtBuilder.CenterInField(
			strSpec.StrValue,
			strSpec.StrFieldLength,
			strSpec.StrPadChar,
			b,
			ePrefix)

	default:
		err = fmt.Errorf(ePrefix+
			"\nError: strSpec.StrPosition is invalid!\n"+
			"StrPosition Value='%v'\n", strSpec.StrPosition.UtilityValue())
	}

	if err != nil {
		return err
	}

	return nil
}

// LeftJustifyField - Left Justifies parameter 'str' in a field of length 'fieldLen'.
func (txtBuilder TextLineBuilder) LeftJustifyField(
	strValue string,
	strFieldLen int,
	trailingPadChar rune,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TzStrFmt.LeftJustifyField() "

	if b == nil {
		return errors.New(ePrefix +
			"\nError: Input parameter b *strings.Builder is nil!")
	}

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix+
			"\nError: Input parameter 'strValue' length exceeds field length parameter, 'strFieldLen'.\n"+
			"strValue='%v'\n"+
			"strValue length='%v'\n"+
			"strFieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"\nError returned by b.WriteString(strValue)\n"+
			"strValue='%v'\n"+
			"Error='%v'\n", strValue, err.Error())
	}

	if lenStr == strFieldLen {
		return nil
	}

	if trailingPadChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'trailingPadChar' is Zero!")
	}

	padLen := strFieldLen - lenStr

	for i := 0; i < padLen; i++ {

		_, err = b.WriteRune(trailingPadChar)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteRune(trailingPadChar)\n"+
				"trailingPadChar='%v'   i='%v'\n"+
				"Error='%v'\n", trailingPadChar, i, err.Error())
		}
	}

	return nil
}

func (txtBuilder TextLineBuilder) RightJustifyField(
	strValue string,
	strFieldLen int,
	leadingPadChar rune,
	b *strings.Builder,
	ePrefix string) error {

	ePrefix += "TextLineBuilder.RightJustifyField() "

	lenStr := len(strValue)

	if lenStr > strFieldLen {
		return fmt.Errorf(ePrefix+
			"\nError: Input Parameter 'strValue' length exceeds 'fieldLen'.\n"+
			"strValue='%v'"+
			"strValue length='%v'\n"+
			"fieldLen='%v'\n", strValue, lenStr, strFieldLen)
	}

	var err error

	if lenStr == strFieldLen {

		_, err = b.WriteString(strValue)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"Error returned by b.WriteString(strValue)\n"+
				"strValue='%v'\n"+
				"Error='%v'\n", strValue, err.Error())
		}

		return nil
	}

	if leadingPadChar == 0 {
		return errors.New(ePrefix +
			"\nError: Input parameter 'leadingPadChar' is ZERO!\n")
	}

	padLen := strFieldLen - lenStr

	for i := 0; i < padLen; i++ {

		_, err = b.WriteRune(leadingPadChar)

		if err != nil {
			return fmt.Errorf(ePrefix+
				"\nError returned by b.WriteRune(leadingPadChar)\n"+
				"leadingPadChar='%v'   i='%v'\n"+
				"Error='%v'\n", leadingPadChar, i, err.Error())
		}
	}

	_, err = b.WriteString(strValue)

	if err != nil {
		return fmt.Errorf(ePrefix+
			"Error returned after padChar by b.WriteString(strValue)\n"+
			"strValue='%v'\n"+
			"Error='%v'\n", strValue, err.Error())
	}

	return nil
}

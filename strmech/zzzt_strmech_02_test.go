package strmech

import (
	"strings"
	"testing"
)

func TestStrMech_ExtractDataField_01(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_01() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Link:\t US/Central\t\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:", "Link:"}
	expectedDataFieldStr := "America/Chicago"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\t"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()
	expectedLeadingWordDelimiter := leadingKeyWordDelimiters[0]
	expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, expectedLeadingWordDelimiter)
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1

	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	sMech := StrMech{}

	datDto,
		err := sMech.ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedDataFieldTrailingDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.DataFieldTrailingDelimiter), true))
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrMech_ExtractDataField_02(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_02() "

	endOfLineDelimiters := []string{"\n"}

	commentDelimiters := []string{"#"}

	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Link:", "Duplicate:", "Zone:"}
	expectedLeadingKeyWordDelimiter := leadingKeyWordDelimiters[2]
	expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, expectedLeadingKeyWordDelimiter)
	expectedDataFieldStr := "America/Chicago"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := " "
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1

	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	datDto,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrMech_ExtractDataField_03(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_03() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}
	targetStr := " America/Chicago Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:", "Link:", "Duplicate:"}
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\n"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfLine()
	expectedLeadingKeyWordDelimiter := ""
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1
	expectedNextTargetIdx := -1

	datDto,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrMech_ExtractDataField_04(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_04() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago Good morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	startIdx := 6
	// leadingKeyWordDelimiters consisting if a zero length
	// array or an array of empty strings are ignored.
	leadingKeyWordDelimiters := make([]string, 0)
	expectedLeadingKeyWordDelimiter := ""
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedDataFieldStr := "America/Chicago"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := " "
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfField()
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1
	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	datDto,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}

}

func TestStrMech_ExtractDataField_05(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_05() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t #America/Chicago\t Link:\tUS/Pacific\tGood morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.Index(targetStr, "#")
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Link:", "Duplicate:", "Zone:"}
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "#"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.Comment()
	expectedLeadingKeyWordDelimiter := leadingKeyWordDelimiters[2]
	expectedLeadingKeyWordDelimiterIndex := strings.Index(targetStr, expectedLeadingKeyWordDelimiter)
	expectedNextTargetIdx := -1
	expectedEndOfLineDelimiter := "\n"
	expectedEndOfLineDelimiterIdx := strings.Index(targetStr, "\n")

	expectedCommentDelimiter := "#"
	expectedCommentDelimiterIndex := strings.Index(targetStr, "#")

	datDto,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}
}

func TestStrMech_ExtractDataField_06(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_06() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " #Zone:\t America/Chicago\t Duplicate:\tGood morning America!\n"
	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := strings.LastIndex(targetStr, "#")
	expectedLastGoodIdx--
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Link:", "Duplicate:", "Zone:"}
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "#"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.Comment()
	expectedLeadingKeyWordDelimiters := ""
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedNextTargetIdx := -1
	expectedEndOfLineDelimiter := "\n"
	expectedEndOfLineDelimiterIdx := strings.Index(targetStr, "\n")
	expectedCommentDelimiter := "#"
	expectedCommentDelimiterIndex := strings.Index(targetStr, "#")

	datDto,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingKeyWordDelimiters != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiters, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}
}

func TestStrMech_ExtractDataField_07(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_07() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := "\tZone:\tAmerica/Chicago\t\tLink:\tAmerica/New_York\t\tZone:\tAmerica/Los_Angeles\n"
	expectedLastGoodIdx := strings.LastIndex(targetStr, "\n")
	expectedEndOfLineDelimiterIdx := expectedLastGoodIdx
	expectedLastGoodIdx--
	lenTargetStr := len(targetStr)
	startIdx := 0
	expectedStartIdx := 46
	leadingKeyWordDelimiters := []string{"Zone:", "Link:", "Duplicate:"}
	expectedDataFieldStr := "America/Los_Angeles"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "\n"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfLine()
	expectedLeadingKeyWordDelimiter := "Zone:"
	expectedLeadingKeyWordDelimiterIndex := strings.LastIndex(targetStr, expectedLeadingKeyWordDelimiter)
	expectedEndOfLineDelimiter := "\n"
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1
	expectedNextTargetIdx := expectedDataFieldIdx + expectedDataFieldLength

	if expectedNextTargetIdx > expectedLastGoodIdx {
		expectedNextTargetIdx = -1
	}

	var datDto DataFieldProfileDto
	var err error

	for i := 0; i < 3; i++ {

		datDto,
			err = StrMech{}.Ptr().ExtractDataField(
			targetStr,
			leadingKeyWordDelimiters,
			startIdx,
			leadingFieldDelimiters,
			trailingFieldDelimiters,
			commentDelimiters,
			endOfLineDelimiters,
			ePrefix)

		if err != nil {
			t.Errorf("Error returned by StrMech{}.Ptr().ExtractDataField()\n"+
				"Cycle No='%v'\n"+
				"targetStr='%v'\tstartIdx='%v'\n"+
				"Error='%v'\n", i, targetStr, startIdx, err.Error())
			return
		}

		startIdx = datDto.NextTargetStrIndex
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if expectedStartIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			expectedStartIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}
}

func TestStrMech_ExtractDataField_08(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_08() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := "\tZone:\tAmerica/Chicago\t\t#Zone:\tAmerica/New_York\t\tLink:\tAmerica/Los_Angeles\n"
	expectedLastGoodIdx := strings.LastIndex(targetStr, "#")
	expectedLastGoodIdx--
	lenTargetStr := len(targetStr)
	startIdx := 3
	expectedStartIdx := 3
	leadingKeyWordDelimiters := []string{"Zone:", "Link:"}
	expectedLeadingKeyWordDelimiter := ""
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedDataFieldStr := ""
	expectedDataFieldIdx := -1
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := "#"
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.Comment()
	expectedEndOfLineDelimiter := "\n"
	expectedEndOfLineDelimiterIdx := strings.Index(targetStr, "\n")
	expectedCommentDelimiter := "#"
	expectedCommentDelimiterIndex := strings.Index(targetStr, "#")
	expectedNextTargetIdx := -1

	var datDto DataFieldProfileDto
	var err error

	datDto,
		err = StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if expectedStartIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			expectedStartIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			leadingKeyWordDelimiters, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			expectedDataFieldTrailingDelimiter, datDto.DataFieldTrailingDelimiter)
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}
}

func TestStrMech_ExtractDataField_09(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_09() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := make([]string, 0)

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\tLink:\t US/Central/t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Link:", "Zone:", "Duplicate:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for  StrMech{}.Ptr().ExtractDataField()\n" +
			"because input parameter 'leadingFieldDelimiters' is a zero length string array.\n" +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestStrMech_ExtractDataField_10(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_10() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := make([]string, 0)

	targetStr := " Zone:\t America/Chicago\tLink:\tGood morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Link:", "Zone:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for StrMech{}.Ptr().ExtractDataField(...)\n" +
			"because input parameter 'trailingFieldDelimiters' is a zero length string array.\n" +
			"However, NO ERROR WAS RETURNED!")
	}
}

func TestStrMech_ExtractDataField_11(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_11() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := ""
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for StrMech{}.Ptr().ExtractDataField(...)\n" +
			"because input parameter 'targetStr' is an empty string.\n" +
			"However, NO ERROR WAS RETURNED!")
	}

}

func TestStrMech_ExtractDataField_12(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_12() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	startIdx := -1
	leadingKeyWordDelimiters := []string{"Zone:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for X\n" +
			"because input parameter 'startIdx' is less than zero.\n" +
			"However, NO ERROR WAS RETURNED!")
	}
}

func TestStrMech_ExtractDataField_13(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_13() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	startIdx := 999
	leadingKeyWordDelimiters := []string{"Zone:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for StrMech{}.Ptr().ExtractDataField()\n" +
			"because input parameter 'startIdx' is exceeds the outer boundary of 'targetStr'.\n" +
			"However, NO ERROR WAS RETURNED!")
	}
}

func TestStrMech_ExtractDataField_14(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_14() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v"}

	targetStr := "Good morning America!"

	lenTargetStr := len(targetStr)
	expectedLastGoodIdx := lenTargetStr - 1
	expectedEndOfLineDelimiterIdx := -1
	startIdx := 0
	leadingKeyWordDelimiters := make([]string, 0)
	expectedDataFieldStr := "Good morning America!"
	expectedDataFieldIdx := strings.Index(targetStr, expectedDataFieldStr)
	expectedDataFieldLength := len(expectedDataFieldStr)
	expectedDataFieldTrailingDelimiter := ""
	expectedDataFieldTrailingDelimiterType := DfTrailDelimiter.EndOfString()
	expectedLeadingKeyWordDelimiter := ""
	expectedLeadingKeyWordDelimiterIndex := -1
	expectedEndOfLineDelimiter := ""
	expectedCommentDelimiter := ""
	expectedCommentDelimiterIndex := -1

	expectedNextTargetIdx := -1

	datDto,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractDataField()\n"+
			"targetStr='%v'\tstartIdx='%v'\n"+
			"Error='%v'\n", targetStr, startIdx, err.Error())
		return
	}

	if targetStr != datDto.TargetStr {
		t.Errorf("ERROR: Expected datDto.TargetStr='%v'.\n"+
			"Instead, datDto.TargetStr='%v'.\n",
			targetStr, datDto.TargetStr)
	}

	if lenTargetStr != datDto.TargetStrLength {
		t.Errorf("ERROR: Expected datDto.TargetStrLength='%v'.\n"+
			"Instead, datDto.TargetStrLength='%v'.\n",
			lenTargetStr, datDto.TargetStrLength)
	}

	if startIdx != datDto.TargetStrStartIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrStartIndex='%v'.\n"+
			"Instead, datDto.TargetStrStartIndex='%v'.\n",
			startIdx, datDto.TargetStrStartIndex)
	}

	if expectedLeadingKeyWordDelimiter != datDto.LeadingKeyWordDelimiter {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiter='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiter='%v'.\n",
			expectedLeadingKeyWordDelimiter, datDto.LeadingKeyWordDelimiter)
	}

	if expectedLeadingKeyWordDelimiterIndex != datDto.LeadingKeyWordDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.LeadingKeyWordDelimiterIndex='%v'.\n"+
			"Instead, datDto.LeadingKeyWordDelimiterIndex='%v'.\n",
			expectedLeadingKeyWordDelimiterIndex, datDto.LeadingKeyWordDelimiterIndex)
	}

	if expectedDataFieldStr != datDto.DataFieldStr {
		t.Errorf("ERROR: Expected datDto.DataFieldStr='%v'.\n"+
			"Instead, datDto.DataFieldStr='%v'.\n",
			expectedDataFieldStr, datDto.DataFieldStr)
	}

	if expectedDataFieldLength != datDto.DataFieldLength {
		t.Errorf("ERROR: Expected datDto.DataFieldLength='%v'.\n"+
			"Instead, datDto.DataFieldLength='%v'.\n",
			expectedDataFieldLength, datDto.DataFieldLength)
	}

	if expectedDataFieldIdx != datDto.DataFieldIndex {
		t.Errorf("ERROR: Expected datDto.DataFieldIndex='%v'.\n"+
			"Instead, datDto.DataFieldIndex='%v'.\n",
			expectedDataFieldIdx, datDto.DataFieldIndex)
	}

	if expectedDataFieldTrailingDelimiter != datDto.DataFieldTrailingDelimiter {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiter='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedDataFieldTrailingDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.DataFieldTrailingDelimiter), true))
	}

	if expectedDataFieldTrailingDelimiterType != datDto.DataFieldTrailingDelimiterType {
		t.Errorf("ERROR: Expected datDto.DataFieldTrailingDelimiterType='%v'.\n"+
			"Instead, datDto.DataFieldTrailingDelimiterType='%v'.\n",
			expectedDataFieldTrailingDelimiterType.String(), datDto.DataFieldTrailingDelimiterType.String())
	}

	if expectedLastGoodIdx != datDto.TargetStrLastGoodIndex {
		t.Errorf("ERROR: Expected datDto.TargetStrLastGoodIndex='%v'.\n"+
			"Instead, datDto.TargetStrLastGoodIndex='%v'.\n",
			expectedLastGoodIdx, datDto.TargetStrLastGoodIndex)
	}

	if expectedNextTargetIdx != datDto.NextTargetStrIndex {
		t.Errorf("ERROR: Expected datDto.NextTargetStrIndex='%v'.\n"+
			"Instead, datDto.NextTargetStrIndex='%v'.\n",
			expectedNextTargetIdx, datDto.NextTargetStrIndex)
	}

	if expectedEndOfLineDelimiter != datDto.EndOfLineDelimiter {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiter='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedEndOfLineDelimiter), false),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.EndOfLineDelimiter), false))
	}

	if expectedEndOfLineDelimiterIdx != datDto.EndOfLineDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.EndOfLineDelimiterIndex='%v'.\n"+
			"Instead, datDto.EndOfLineDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.EndOfLineDelimiterIndex)
	}

	if expectedCommentDelimiter != datDto.CommentDelimiter {
		t.Errorf("ERROR: Expected datDto.CommentDelimiter='%v'.\n"+
			"Instead, datDto.CommentDelimiter='%v'.\n",
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(expectedCommentDelimiter), true),
			StrMech{}.NewPtr().ConvertNonPrintableChars([]rune(datDto.CommentDelimiter), true))
	}

	if expectedCommentDelimiterIndex != datDto.CommentDelimiterIndex {
		t.Errorf("ERROR: Expected datDto.CommentDelimiterIndex='%v'.\n"+
			"Instead, datDto.CommentDelimiterIndex='%v'.\n",
			expectedEndOfLineDelimiterIdx, datDto.CommentDelimiterIndex)
	}
}

func TestStrMech_ExtractDataField_15(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_15() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"",
		"",
		"",
		"",
		""}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for StrMech{}.Ptr().ExtractDataField()\n" +
			"because input parameter 'leadingFieldDelimiters' array consists entirely\n" +
			"of empty strings. \n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestStrMech_ExtractDataField_16(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_16() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"",
		"",
		"",
		""}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for StrMech{}.Ptr().ExtractDataField()\n" +
			"because input parameter 'trailingFieldDelimiters' array consists entirely\n" +
			"of empty strings. \n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestStrMech_ExtractDataField_17(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_17() "

	endOfLineDelimiters := []string{"", "", ""}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for StrMech{}.Ptr().ExtractDataField()\n" +
			"because input parameter 'endOfLineDelimiters' array consists entirely\n" +
			"of empty strings. \n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestStrMech_ExtractDataField_18(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_18() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"", "", ""}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:"}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for StrMech{}.Ptr().ExtractDataField()\n" +
			"because input parameter 'commentDelimiters' array consists entirely\n" +
			"of empty strings. \n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestStrMech_ExtractDataField_19(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_19() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := "#Zone:\t America/Chicago\tLink:\t US/Central\t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:", "Link:"}

	fieldDto,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractDataField()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if fieldDto.TargetStrLastGoodIndex != -1 {
		t.Errorf("Expected TargetStrLastGoodIndex==-1.\n"+
			"Instead, TargetStrLastGoodIndex='%v'\n", fieldDto.TargetStrLastGoodIndex)
	}
}

func TestStrMech_ExtractDataField_20(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_20() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := "\nZone:\t America/Chicago\tLink:\t US/Central\t Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{"Zone:", "Link:"}

	fieldDto,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractDataField()\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if fieldDto.TargetStrLastGoodIndex != -1 {
		t.Errorf("Expected TargetStrLastGoodIndex==-1.\n"+
			"Instead, TargetStrLastGoodIndex='%v'\n", fieldDto.TargetStrLastGoodIndex)
	}
}

func TestStrMech_ExtractDataField_21(t *testing.T) {

	ePrefix := "TestStrMech_ExtractDataField_21() "

	endOfLineDelimiters := []string{"\n"}
	commentDelimiters := []string{"#"}
	leadingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	trailingFieldDelimiters := []string{
		"\t",
		"\r",
		"\f",
		"\v",
		" "}

	targetStr := " Zone:\t America/Chicago\t#Good morning America!\n"
	startIdx := 0
	leadingKeyWordDelimiters := []string{
		"",
		"",
		""}

	_,
		err := StrMech{}.Ptr().ExtractDataField(
		targetStr,
		leadingKeyWordDelimiters,
		startIdx,
		leadingFieldDelimiters,
		trailingFieldDelimiters,
		commentDelimiters,
		endOfLineDelimiters,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return for StrMech{}.Ptr().ExtractDataField()\n" +
			"because input parameter 'leadingKeyWordDelimiters' array consists entirely\n" +
			"of empty strings. \n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

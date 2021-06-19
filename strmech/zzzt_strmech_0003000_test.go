package strmech

import (
	"strings"
	"testing"
)

func TestStrMech_ExtractNumericDigits_01(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_01() "

	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "12"
	expectedNumStrLen := len(expectedNumStr)
	expectedLeadingSignChar := ""
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedLeadingSignCharIndex := -1
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	sMech := StrMech{}

	nStrDto,
		err := sMech.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_02(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_02() "

	targetStr := "Etc/GMT+11"
	startIndex := 0
	keepLeadingChars := "+"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "+11"
	expectedNumStrLen := len(expectedNumStr)
	expectedLeadingSignChar := "+"
	expectedLeadingSignCharIndex := 0
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)

	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_03(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_03() "

	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	startIndex := 23
	keepLeadingChars := "+-"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "+0000"
	expectedNumStrLen := len(expectedNumStr)
	expectedLeadingSignChar := "+"
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedLeadingSignCharIndex := 0
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_04(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_04() "

	targetStr := "2016 1:6:3pm +0000 UTC"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "2016"

	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if "" != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			"", nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_05(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_05() "

	targetStr := "2016"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "2016"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_06(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_06() "

	targetStr := "Hello World! Your bank account =$(1,250,364.33).44 What do you think?"
	startIndex := 0
	keepLeadingChars := "$("
	keepInteriorChars := ",."
	keepTrailingChars := ")"

	expectedNumStr := "$(1,250,364.33)"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_07(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_07() "

	targetStr := "Hello World! The time zone here is 'Etc/GMT+11'. What do you think?"
	startIndex := 0
	keepLeadingChars := "+-"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "+11"
	expectedLeadingSignChar := "+"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_08(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_08() "

	targetStr := "Etc/GMT-4"
	startIndex := 0
	keepLeadingChars := "+-"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "-4"
	expectedLeadingSignChar := "-"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_09(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_09() "

	targetStr := "+$697,621,911.77"
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := "+$697,621,911.77"
	expectedLeadingSignChar := "+"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_10(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_10() "

	targetStr := "Hello World\t+-$697,621,911.77\n"
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := "-$697,621,911.77"
	expectedLeadingSignChar := "-"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_11(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_11() "

	targetStr := "Hello World\t\n"
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := ""
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := -1
	expectedNextTargetStrIndex := -1

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_12(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_12() "

	targetStr := ""
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	_,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return from StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n" +
			"because input parameter 'targetStr' is an empty string.\n" +
			"However, NO ERROR WAS RETURNED!!")
		return
	}
}

func TestStrMech_ExtractNumericDigits_13(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_13() "

	targetStr := "Hello World7Have a great day!"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "7"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_14(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_14() "

	targetStr := "7Hello World Have a great day!"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "7"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_15(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_15() "

	targetStr := "Hello World Have a great day!7"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "7"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_16(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_16() "

	targetStr := "Hello World -7\t6 Have a great day!"
	startIndex := 0
	keepLeadingChars := "+-"
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "-7"
	expectedLeadingSignChar := "-"
	expectedLeadingSignCharIndex := 0
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_17(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_17() "

	targetStr := "Hello World.\t+$-697,621,911.77.\nHow are you.\n"
	startIndex := 0
	keepLeadingChars := "+-$"
	keepInteriorChars := ",."
	keepTrailingChars := ""

	expectedNumStr := "$-697,621,911.77"
	expectedLeadingSignChar := "-"
	expectedLeadingSignCharIndex := 1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

func TestStrMech_ExtractNumericDigits_18(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_19() "

	targetStr := ""
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	_,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return from StrMech{}.Ptr().ExtractNumericDigits()\n" +
			"because input parameter 'targetStr' is an empty string.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
	}
}

func TestStrMech_ExtractNumericDigits_19(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_19() "

	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	startIndex := -1
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	_,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return from StrMech{}.Ptr().ExtractNumericDigits()\n" +
			"because input parameter 'startIdx' is less than zero.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
	}
}

func TestStrMech_ExtractNumericDigits_20(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_20() "

	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	startIndex := 999
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	_,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return from StrMech{}.Ptr().ExtractNumericDigits()\n" +
			"because input parameter 'startIdx' is beyond the maximum boundary for 'targetStr'.\n" +
			"However, NO ERROR WAS RETURNED!!\n")
	}
}

func TestStrMech_ExtractNumericDigits_21(t *testing.T) {

	ePrefix := "TestStrMech_ExtractNumericDigits_21() "

	targetStr := "Etc/GMT-4"
	startIndex := 0
	keepLeadingChars := ""
	keepInteriorChars := ""
	keepTrailingChars := ""

	expectedNumStr := "4"
	expectedLeadingSignChar := ""
	expectedLeadingSignCharIndex := -1
	expectedNumStrLen := len(expectedNumStr)
	expectedNumIdx := strings.Index(targetStr, expectedNumStr)
	expectedNextTargetStrIndex := expectedNumIdx + expectedNumStrLen

	if expectedNextTargetStrIndex >= len(targetStr) {
		expectedNextTargetStrIndex = -1
	}

	nStrDto,
		err := StrMech{}.Ptr().ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ExtractNumericDigits(targetStr, 0)\n"+
			"targetStr='%v'\nError='%v'\n", targetStr, err.Error())
		return
	}

	if expectedNumIdx != nStrDto.FirstNumCharIndex {
		t.Errorf("Expected starting numeric index='%v'\n"+
			"Instead, staring numeric index='%v'\n",
			expectedNumIdx, nStrDto.FirstNumCharIndex)
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)
	}
}

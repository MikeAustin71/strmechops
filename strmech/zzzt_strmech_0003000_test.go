package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"strings"
	"testing"
	"time"
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

		return
	}

	if expectedNumStr != nStrDto.NumStr {
		t.Errorf("Expected number string ='%v'\n"+
			"Instead, number string ='%v'\n",
			expectedNumStr, nStrDto.NumStr)

		return
	}

	if expectedNumStrLen != nStrDto.NumStrLen {
		t.Errorf("Expected number string length ='%v'\n"+
			"Instead, number string length ='%v'\n",
			expectedNumStrLen, nStrDto.NumStrLen)

		return
	}

	if expectedLeadingSignChar != nStrDto.LeadingSignChar {
		t.Errorf("Expected leading sign char ='%v'\n"+
			"Instead, leading sign char ='%v'\n",
			expectedLeadingSignChar, nStrDto.LeadingSignChar)

		return
	}

	if expectedLeadingSignCharIndex != nStrDto.LeadingSignIndex {
		t.Errorf("Expected leading sign char index ='%v'\n"+
			"Instead, leading sign char index ='%v'\n",
			expectedLeadingSignCharIndex, nStrDto.LeadingSignIndex)

		return
	}

	if expectedNextTargetStrIndex != nStrDto.NextTargetStrIndex {
		t.Errorf("Expected next target index after number string ='%v'\n"+
			"Instead, next target string index ='%v'\n",
			expectedNextTargetStrIndex, nStrDto.NextTargetStrIndex)

		return
	}

	_,
		err = sMech.ExtractNumericDigits(
		targetStr,
		startIndex,
		keepLeadingChars,
		keepInteriorChars,
		keepTrailingChars,
		strMechMolecule{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractNumericDigits()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

	return
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

/*


func TestStrMech_ExtractNumberRunes_01(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractNumberRunes_01()",
		"")
	rawNumStrRunes := []rune("1234.5678")
	leadingNegativeSignChars := []rune{'-'}
	var trailingNegativeSignChars []rune
	decimalSeparatorChars := []rune{'.'}

	expectedIntRunes := []rune{'1', '2', '3', '4'}
	expectedFractionalRunes := []rune{'5', '6', '7', '8'}
	expectedNumSign := 1
	expectedDigitsFound := 8

	sMech := StrMech{}

	integerDigits,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedIntRunes,
		integerDigits)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes()\n"+
			"Expected Integer Runes DO NOT MATCH\n"+
			"Actual Integer Runes.\n"+
			"Expected Integer Runes = '%v'\n"+
			"  Actual Integer Runes = '%v'\n",
			ePrefix.String(),
			string(expectedIntRunes),
			string(integerDigits))

		return

	}

	areEqual = sMechPreon.equalRuneArrays(
		expectedFractionalRunes,
		fracRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Fractional Runes\n"+
			"Expected Fractional Runes DO NOT MATCH\n"+
			"Actual Fractional Runes.\n"+
			"Expected Fractional Runes = '%v'\n"+
			"  Actual Fractional Runes = '%v'\n",
			ePrefix.String(),
			string(expectedFractionalRunes),
			string(fracRunes))

		return

	}

	if expectedNumSign != numberSign {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Number Sign\n"+
			"Expected Number Sign DOES NOT MATCH\n"+
			"Actual Number Sign.\n"+
			"Expected Number Sign = '%v'\n"+
			"  Actual Number Sign = '%v'\n",
			ePrefix.String(),
			expectedNumSign,
			numberSign)

		return

	}

	if expectedDigitsFound != digitsFound {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Digits Found\n"+
			"Expected Digits Found DOES NOT MATCH\n"+
			"Actual Digits Found.\n"+
			"Expected Digits Found = '%v'\n"+
			"  Actual Digits Found = '%v'\n",
			expectedDigitsFound,
			digitsFound,
			ePrefix.String())

		return

	}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		strMechMolecule{})

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractNumberRunes()\n"+
			"because 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		nil,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractNumberRunes()\n"+
			"because 'rawNumStrRunes' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		rawNumStrRunes,
		nil,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		nil,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		nil,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	badRawNumStrRunes :=
		[]rune{'1', '2', '3', '4', 0, '.', '5', '6', '7', '8'}
	badLeadingNegativeSignChars := []rune{'-', 0}
	badTrailingNegativeSignChars := []rune{0}
	badDecimalSeparatorChars := []rune{'.', 0}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		badRawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractNumberRunes()\n"+
			"because 'rawNumStrRunes' is an invalid character array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		rawNumStrRunes,
		badLeadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractNumberRunes()\n"+
			"because 'leadingNegativeSignChars' is an invalid character array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		badTrailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractNumberRunes()\n"+
			"because 'trailingNegativeSignChars' is an invalid character array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

	_,
		_,
		_,
		_,
		err = sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		badDecimalSeparatorChars,
		ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractNumberRunes()\n"+
			"because 'decimalSeparatorChars' is an invalid character array.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix)

		return
	}

	return
}

func TestStrMech_ExtractNumberRunes_02(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractNumberRunes_02()",
		"")
	rawNumStrRunes := []rune("-1234.5678")
	leadingNegativeSignChars := []rune{'-'}
	var trailingNegativeSignChars []rune
	decimalSeparatorChars := []rune{'.'}

	expectedIntRunes := []rune{'1', '2', '3', '4'}
	expectedFractionalRunes := []rune{'5', '6', '7', '8'}
	expectedNumSign := -1
	expectedDigitsFound := 8

	sMech := StrMech{}

	integerDigits,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedIntRunes,
		integerDigits)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes()\n"+
			"Expected Integer Runes DO NOT MATCH\n"+
			"Actual Integer Runes.\n"+
			"Expected Integer Runes = '%v'\n"+
			"  Actual Integer Runes = '%v'\n",
			ePrefix.String(),
			string(expectedIntRunes),
			string(integerDigits))

		return

	}

	areEqual = sMechPreon.equalRuneArrays(
		expectedFractionalRunes,
		fracRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Fractional Runes\n"+
			"Expected Fractional Runes DO NOT MATCH\n"+
			"Actual Fractional Runes.\n"+
			"Expected Fractional Runes = '%v'\n"+
			"  Actual Fractional Runes = '%v'\n",
			ePrefix.String(),
			string(expectedFractionalRunes),
			string(fracRunes))

		return

	}

	if expectedNumSign != numberSign {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Number Sign\n"+
			"Expected Number Sign DOES NOT MATCH\n"+
			"Actual Number Sign.\n"+
			"Expected Number Sign = '%v'\n"+
			"  Actual Number Sign = '%v'\n",
			ePrefix.String(),
			expectedNumSign,
			numberSign)

		return

	}

	if expectedDigitsFound != digitsFound {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Digits Found\n"+
			"Expected Digits Found DOES NOT MATCH\n"+
			"Actual Digits Found.\n"+
			"Expected Digits Found = '%v'\n"+
			"  Actual Digits Found = '%v'\n",
			expectedDigitsFound,
			digitsFound,
			ePrefix.String())

		return

	}

	return
}

func TestStrMech_ExtractNumberRunes_03(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractNumberRunes_03()",
		"")
	rawNumStrRunes := []rune("-1234")
	leadingNegativeSignChars := []rune{'-'}
	var trailingNegativeSignChars []rune
	decimalSeparatorChars := []rune{'.'}

	expectedIntRunes := []rune{'1', '2', '3', '4'}
	var expectedFractionalRunes []rune
	expectedNumSign := -1
	expectedDigitsFound := 4

	sMech := StrMech{}

	integerDigits,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedIntRunes,
		integerDigits)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes()\n"+
			"Expected Integer Runes DO NOT MATCH\n"+
			"Actual Integer Runes.\n"+
			"Expected Integer Runes = '%v'\n"+
			"  Actual Integer Runes = '%v'\n",
			ePrefix.String(),
			string(expectedIntRunes),
			string(integerDigits))

		return

	}

	areEqual = sMechPreon.equalRuneArrays(
		expectedFractionalRunes,
		fracRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Fractional Runes\n"+
			"Expected Fractional Runes DO NOT MATCH\n"+
			"Actual Fractional Runes.\n"+
			"Expected Fractional Runes = '%v'\n"+
			"  Actual Fractional Runes = '%v'\n",
			ePrefix.String(),
			string(expectedFractionalRunes),
			string(fracRunes))

		return

	}

	if expectedNumSign != numberSign {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Number Sign\n"+
			"Expected Number Sign DOES NOT MATCH\n"+
			"Actual Number Sign.\n"+
			"Expected Number Sign = '%v'\n"+
			"  Actual Number Sign = '%v'\n",
			ePrefix.String(),
			expectedNumSign,
			numberSign)

		return

	}

	if expectedDigitsFound != digitsFound {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Digits Found\n"+
			"Expected Digits Found DOES NOT MATCH\n"+
			"Actual Digits Found.\n"+
			"Expected Digits Found = '%v'\n"+
			"  Actual Digits Found = '%v'\n",
			expectedDigitsFound,
			digitsFound,
			ePrefix.String())

		return

	}

	return
}

func TestStrMech_ExtractNumberRunes_04(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractNumberRunes_04()",
		"")
	rawNumStrRunes := []rune("1234")
	leadingNegativeSignChars := []rune{'-'}
	var trailingNegativeSignChars []rune
	decimalSeparatorChars := []rune{'.'}

	expectedIntRunes := []rune{'1', '2', '3', '4'}
	var expectedFractionalRunes []rune
	expectedNumSign := 1
	expectedDigitsFound := 4

	sMech := StrMech{}

	integerDigits,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedIntRunes,
		integerDigits)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes()\n"+
			"Expected Integer Runes DO NOT MATCH\n"+
			"Actual Integer Runes.\n"+
			"Expected Integer Runes = '%v'\n"+
			"  Actual Integer Runes = '%v'\n",
			ePrefix.String(),
			string(expectedIntRunes),
			string(integerDigits))

		return

	}

	areEqual = sMechPreon.equalRuneArrays(
		expectedFractionalRunes,
		fracRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Fractional Runes\n"+
			"Expected Fractional Runes DO NOT MATCH\n"+
			"Actual Fractional Runes.\n"+
			"Expected Fractional Runes = '%v'\n"+
			"  Actual Fractional Runes = '%v'\n",
			ePrefix.String(),
			string(expectedFractionalRunes),
			string(fracRunes))

		return

	}

	if expectedNumSign != numberSign {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Number Sign\n"+
			"Expected Number Sign DOES NOT MATCH\n"+
			"Actual Number Sign.\n"+
			"Expected Number Sign = '%v'\n"+
			"  Actual Number Sign = '%v'\n",
			ePrefix.String(),
			expectedNumSign,
			numberSign)

		return

	}

	if expectedDigitsFound != digitsFound {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Digits Found\n"+
			"Expected Digits Found DOES NOT MATCH\n"+
			"Actual Digits Found.\n"+
			"Expected Digits Found = '%v'\n"+
			"  Actual Digits Found = '%v'\n",
			expectedDigitsFound,
			digitsFound,
			ePrefix.String())

		return

	}

	return
}

func TestStrMech_ExtractNumberRunes_05(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractNumberRunes_05()",
		"")
	rawNumStrRunes := []rune(".1234")
	leadingNegativeSignChars := []rune{'-'}
	var trailingNegativeSignChars []rune
	decimalSeparatorChars := []rune{'.'}

	expectedIntRunes := []rune{'0'}
	expectedFractionalRunes := []rune{'1', '2', '3', '4'}
	expectedNumSign := 1
	expectedDigitsFound := 4

	sMech := StrMech{}

	integerDigits,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedIntRunes,
		integerDigits)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes()\n"+
			"Expected Integer Runes DO NOT MATCH\n"+
			"Actual Integer Runes.\n"+
			"Expected Integer Runes = '%v'\n"+
			"  Actual Integer Runes = '%v'\n",
			ePrefix.String(),
			string(expectedIntRunes),
			string(integerDigits))

		return

	}

	areEqual = sMechPreon.equalRuneArrays(
		expectedFractionalRunes,
		fracRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Fractional Runes\n"+
			"Expected Fractional Runes DO NOT MATCH\n"+
			"Actual Fractional Runes.\n"+
			"Expected Fractional Runes = '%v'\n"+
			"  Actual Fractional Runes = '%v'\n",
			ePrefix.String(),
			string(expectedFractionalRunes),
			string(fracRunes))

		return

	}

	if expectedNumSign != numberSign {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Number Sign\n"+
			"Expected Number Sign DOES NOT MATCH\n"+
			"Actual Number Sign.\n"+
			"Expected Number Sign = '%v'\n"+
			"  Actual Number Sign = '%v'\n",
			ePrefix.String(),
			expectedNumSign,
			numberSign)

		return

	}

	if expectedDigitsFound != digitsFound {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Digits Found\n"+
			"Expected Digits Found DOES NOT MATCH\n"+
			"Actual Digits Found.\n"+
			"Expected Digits Found = '%v'\n"+
			"  Actual Digits Found = '%v'\n",
			expectedDigitsFound,
			digitsFound,
			ePrefix.String())

		return

	}

	return
}

func TestStrMech_ExtractNumberRunes_06(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractNumberRunes_06()",
		"")
	rawNumStrRunes := []rune("-.1234")
	leadingNegativeSignChars := []rune{'-'}
	var trailingNegativeSignChars []rune
	decimalSeparatorChars := []rune{'.'}

	expectedIntRunes := []rune{'0'}
	expectedFractionalRunes := []rune{'1', '2', '3', '4'}
	expectedNumSign := -1
	expectedDigitsFound := 4

	sMech := StrMech{}

	integerDigits,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedIntRunes,
		integerDigits)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes()\n"+
			"Expected Integer Runes DO NOT MATCH\n"+
			"Actual Integer Runes.\n"+
			"Expected Integer Runes = '%v'\n"+
			"  Actual Integer Runes = '%v'\n",
			ePrefix.String(),
			string(expectedIntRunes),
			string(integerDigits))

		return

	}

	areEqual = sMechPreon.equalRuneArrays(
		expectedFractionalRunes,
		fracRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Fractional Runes\n"+
			"Expected Fractional Runes DO NOT MATCH\n"+
			"Actual Fractional Runes.\n"+
			"Expected Fractional Runes = '%v'\n"+
			"  Actual Fractional Runes = '%v'\n",
			ePrefix.String(),
			string(expectedFractionalRunes),
			string(fracRunes))

		return

	}

	if expectedNumSign != numberSign {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Number Sign\n"+
			"Expected Number Sign DOES NOT MATCH\n"+
			"Actual Number Sign.\n"+
			"Expected Number Sign = '%v'\n"+
			"  Actual Number Sign = '%v'\n",
			ePrefix.String(),
			expectedNumSign,
			numberSign)

		return

	}

	if expectedDigitsFound != digitsFound {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Digits Found\n"+
			"Expected Digits Found DOES NOT MATCH\n"+
			"Actual Digits Found.\n"+
			"Expected Digits Found = '%v'\n"+
			"  Actual Digits Found = '%v'\n",
			expectedDigitsFound,
			digitsFound,
			ePrefix.String())

		return

	}

	return
}

func TestStrMech_ExtractNumberRunes_07(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractNumberRunes_07()",
		"")
	rawNumStrRunes := []rune("0.1234")
	leadingNegativeSignChars := []rune{'-'}
	var trailingNegativeSignChars []rune
	decimalSeparatorChars := []rune{'.'}

	expectedIntRunes := []rune{'0'}
	expectedFractionalRunes := []rune{'1', '2', '3', '4'}
	expectedNumSign := 1
	expectedDigitsFound := 5

	sMech := StrMech{}

	integerDigits,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedIntRunes,
		integerDigits)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes()\n"+
			"Expected Integer Runes DO NOT MATCH\n"+
			"Actual Integer Runes.\n"+
			"Expected Integer Runes = '%v'\n"+
			"  Actual Integer Runes = '%v'\n",
			ePrefix.String(),
			string(expectedIntRunes),
			string(integerDigits))

		return

	}

	areEqual = sMechPreon.equalRuneArrays(
		expectedFractionalRunes,
		fracRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Fractional Runes\n"+
			"Expected Fractional Runes DO NOT MATCH\n"+
			"Actual Fractional Runes.\n"+
			"Expected Fractional Runes = '%v'\n"+
			"  Actual Fractional Runes = '%v'\n",
			ePrefix.String(),
			string(expectedFractionalRunes),
			string(fracRunes))

		return

	}

	if expectedNumSign != numberSign {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Number Sign\n"+
			"Expected Number Sign DOES NOT MATCH\n"+
			"Actual Number Sign.\n"+
			"Expected Number Sign = '%v'\n"+
			"  Actual Number Sign = '%v'\n",
			ePrefix.String(),
			expectedNumSign,
			numberSign)

		return

	}

	if expectedDigitsFound != digitsFound {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Digits Found\n"+
			"Expected Digits Found DOES NOT MATCH\n"+
			"Actual Digits Found.\n"+
			"Expected Digits Found = '%v'\n"+
			"  Actual Digits Found = '%v'\n",
			expectedDigitsFound,
			digitsFound,
			ePrefix.String())

		return

	}

	return
}

func TestStrMech_ExtractNumberRunes_08(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractNumberRunes_08()",
		"")
	rawNumStrRunes := []rune("-0.1234")
	leadingNegativeSignChars := []rune{'-'}
	var trailingNegativeSignChars []rune
	decimalSeparatorChars := []rune{'.'}

	expectedIntRunes := []rune{'0'}
	expectedFractionalRunes := []rune{'1', '2', '3', '4'}
	expectedNumSign := -1
	expectedDigitsFound := 5

	sMech := StrMech{}

	integerDigits,
		fracRunes,
		numberSign,
		digitsFound,
		err := sMech.ExtractNumberRunes(
		rawNumStrRunes,
		leadingNegativeSignChars,
		trailingNegativeSignChars,
		decimalSeparatorChars,
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	sMechPreon := strMechPreon{}

	areEqual := sMechPreon.equalRuneArrays(
		expectedIntRunes,
		integerDigits)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes()\n"+
			"Expected Integer Runes DO NOT MATCH\n"+
			"Actual Integer Runes.\n"+
			"Expected Integer Runes = '%v'\n"+
			"  Actual Integer Runes = '%v'\n",
			ePrefix.String(),
			string(expectedIntRunes),
			string(integerDigits))

		return

	}

	areEqual = sMechPreon.equalRuneArrays(
		expectedFractionalRunes,
		fracRunes)

	if !areEqual {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Fractional Runes\n"+
			"Expected Fractional Runes DO NOT MATCH\n"+
			"Actual Fractional Runes.\n"+
			"Expected Fractional Runes = '%v'\n"+
			"  Actual Fractional Runes = '%v'\n",
			ePrefix.String(),
			string(expectedFractionalRunes),
			string(fracRunes))

		return

	}

	if expectedNumSign != numberSign {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Number Sign\n"+
			"Expected Number Sign DOES NOT MATCH\n"+
			"Actual Number Sign.\n"+
			"Expected Number Sign = '%v'\n"+
			"  Actual Number Sign = '%v'\n",
			ePrefix.String(),
			expectedNumSign,
			numberSign)

		return

	}

	if expectedDigitsFound != digitsFound {

		t.Errorf("%v\n"+
			"ERROR: sMech.ExtractNumberRunes() - Digits Found\n"+
			"Expected Digits Found DOES NOT MATCH\n"+
			"Actual Digits Found.\n"+
			"Expected Digits Found = '%v'\n"+
			"  Actual Digits Found = '%v'\n",
			expectedDigitsFound,
			digitsFound,
			ePrefix.String())

		return

	}

	return
}

func TestStrMech_ExtractTextLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractTextLines_000100()",
		"")

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {

		t.Errorf(
			"%v\n"+
				"Error - time.LoadLocation()\n"+
				"%v\n",
			ePrefix.String(),
			err.Error())

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

	var timerLines *TextLineSpecTimerLines

	timerLines,
		err = TextLineSpecTimerLines{}.NewFullTimerEvent(
		" ",
		"Start Time",
		startTime,
		"End Time",
		endTime,
		"2006-01-02 15:04:05.000000000 -0700 MST",
		"Elapsed Time",
		6,
		TxtJustify.Left(),
		": ",
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	timerLinesText := timerLines.String()

	sMech := StrMech{}

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
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if len(remainderStr) > 0 {

		printableRemainderStr :=
			sMech.ConvertNonPrintableString(
				remainderStr,
				true)

		t.Errorf(
			"%v\n"+
				"Error: sMech.ExtractTextLines()\n"+
				"Text Line Extraction Failed!\n"+
				"Remainder String =\n     '%v'\n",
			ePrefix.String(),
			printableRemainderStr)

	}

	if numOfTxtLines <= 0 {

		t.Errorf(
			"%v\n"+
				"sMech.ExtractTextLines() FAILED!\n"+
				"Error: 'numOfTxtLines' is zero.\n",
			ePrefix.String())

		return
	}

	expectedMaxStrLen := 0
	lenTimerLine := 0
	expectedTotalLinesLength := 0

	for i := 0; i < numOfTxtLines; i++ {

		lenTimerLine = len(timerLineStrs[i])

		if lenTimerLine > expectedMaxStrLen {
			expectedMaxStrLen = lenTimerLine
		}

		expectedTotalLinesLength += lenTimerLine
	}

	actualMaxStrLen := timerLines.GetSingleLineLength()

	if expectedMaxStrLen != actualMaxStrLen {

		t.Errorf(
			"%v\n"+
				"Error:"+
				"Expected Single Line Length is NOT EQUAL TO"+
				"Actual Single Line Length!\n"+
				"Expected Single Line Length = '%v'\n"+
				"  Actual Single Line Length = '%v'\n",
			ePrefix.String(),
			expectedMaxStrLen,
			actualMaxStrLen)

		return

	}

	actualTotalLinesLen := timerLines.GetTotalLinesLength()

	if expectedTotalLinesLength != actualTotalLinesLen {

		t.Errorf(
			"%v\n"+
				"Error:"+
				"Expected Total Lines Length is NOT EQUAL TO"+
				"Actual Total Lines Length!\n"+
				"Expected Total Lines Length = '%v'\n"+
				"  Actual Total Lines Length = '%v'\n",
			ePrefix.String(),
			expectedTotalLinesLength,
			actualTotalLinesLen)

		return

	}



	   //Start Time  : 2022-04-05 10:00:00.000000000 -0500 CDT
	   //End Time    : 2022-05-05 10:30:45.000009582 -0500 CDT
	   //Elapsed Time: 30 Days 0 Hours 30 Minutes 45 Seconds 0 Milliseconds
	   //              9 Microseconds 582 Nanoseconds
	   //              Total Elapsed Nanoseconds: 2,593,845,000,009,582


	return
}

*/

// TestStrMech_ExtractTextLines_000200
func TestStrMech_ExtractTextLines_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_ExtractTextLines_000200()",
		"")

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {

		t.Errorf(
			"%v\n"+
				"Error - time.LoadLocation()\n"+
				"%v\n",
			ePrefix.String(),
			err.Error())

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

	var timerLines *TextLineSpecTimerLines

	timerLines,
		err = TextLineSpecTimerLines{}.NewFullTimerEvent(
		" ",
		"Start Time",
		startTime,
		"End Time",
		endTime,
		"2006-01-02 15:04:05.000000000 -0700 MST",
		"Elapsed Time",
		6,
		TxtJustify.Left(),
		": ",
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	err = timerLines.IsValidInstanceError(
		ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	timerLinesText := timerLines.String()

	sMech := StrMech{}

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
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	if len(remainderStr) > 0 {

		printableRemainderStr :=
			sMech.ConvertNonPrintableString(
				remainderStr,
				true)

		t.Errorf(
			"%v\n"+
				"Error: sMech.ExtractTextLines()\n"+
				"Text Line Extraction Failed!\n"+
				"Remainder String =\n     '%v'\n",
			ePrefix.String(),
			printableRemainderStr)

	}

	if numOfTxtLines <= 0 {

		t.Errorf(
			"%v\n"+
				"sMech.ExtractTextLines() FAILED!\n"+
				"Error: 'numOfTxtLines' is zero.\n",
			ePrefix.String())

		return
	}

	expectedMaxStrLen := 0
	lenTimerLine := 0
	expectedTotalLinesLength := 0

	for i := 0; i < numOfTxtLines; i++ {

		lenTimerLine = len(timerLineStrs[i])

		if lenTimerLine > expectedMaxStrLen {
			expectedMaxStrLen = lenTimerLine
		}

		expectedTotalLinesLength += lenTimerLine
	}

	actualMaxStrLen := timerLines.GetSingleLineLength()

	if expectedMaxStrLen != actualMaxStrLen {

		t.Errorf(
			"%v\n"+
				"Error:"+
				"Expected Single Line Length is NOT EQUAL TO"+
				"Actual Single Line Length!\n"+
				"Expected Single Line Length = '%v'\n"+
				"  Actual Single Line Length = '%v'\n",
			ePrefix.String(),
			expectedMaxStrLen,
			actualMaxStrLen)

		return

	}

	actualTotalLinesLen := timerLines.GetTotalLinesLength()

	if expectedTotalLinesLength != actualTotalLinesLen {

		t.Errorf(
			"%v\n"+
				"Error:"+
				"Expected Total Lines Length is NOT EQUAL TO"+
				"Actual Total Lines Length!\n"+
				"Expected Total Lines Length = '%v'\n"+
				"  Actual Total Lines Length = '%v'\n",
			ePrefix.String(),
			expectedTotalLinesLength,
			actualTotalLinesLen)

		return

	}

	_,
		_,
		_,
		err = sMech.ExtractTextLines(
		"",
		eolDelimiters,
		true,
		ePrefix.XCpy(
			"timerLineStrs"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractTextLines()\n"+
			"because input parameter 'targetStr' is empty and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	//timerLineStrs,
	//	numOfTxtLines,
	//	remainderStr,
	//	err = sMech.ExtractTextLines(
	//	timerLinesText,
	//	eolDelimiters,
	//	true,
	//	ePrefix.XCpy(
	//		"timerLineStrs"))

	_,
		_,
		_,
		err = sMech.ExtractTextLines(
		timerLinesText,
		nil,
		true,
		ePrefix.XCpy(
			"timerLineStrs"))

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractTextLines()\n"+
			"because input parameter 'endOfLineDelimiters' is 'nil' and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	_,
		_,
		_,
		err = sMech.ExtractTextLines(
		timerLinesText,
		eolDelimiters,
		true,
		StrMech{})

	if err == nil {

		t.Errorf("\n%v - ERROR\n"+
			"Expected an error return from sMech."+
			"ExtractTextLines()\n"+
			"because input parameter 'errorPrefix' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	testStr := "Hello World! How are you. Have a GREAT Day!"

	timerLineStrs,
		numOfTxtLines,
		remainderStr,
		err = sMech.ExtractTextLines(
		testStr,
		eolDelimiters,
		true,
		ePrefix.XCpy(
			"timerLineStrs"))

	if err != nil {

		t.Errorf("\n%v\n"+
			"sMech.ExtractTextLines()\n"+
			"Did not expect an error because input parameter\n"+
			"'targetStr' did NOT contain any End-Of-Line Characters\n"+
			"HOWEVER, AN ERROR WAS RETURNED!\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if testStr != remainderStr {

		t.Errorf("\n%v\n"+
			"sMech.ExtractTextLines()\n"+
			"Expected that 'testStr' would equal 'remainderStr' because input\n"+
			"parameter 'testStr' did NOT contain any End-Of-Line Characters\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"     testStr = '%v'\n"+
			"remainderStr = '%v'\n",
			ePrefix.String(),
			testStr,
			remainderStr)

		return

	}

	testStr = "Good Morning!\nWake up sleepy head!\nHello World! How are you.\nHave a GREAT Day!"

	expectedStr := "Have a GREAT Day!"

	timerLineStrs,
		numOfTxtLines,
		remainderStr,
		err = sMech.ExtractTextLines(
		testStr,
		eolDelimiters,
		true,
		ePrefix.XCpy(
			"timerLineStrs"))

	if expectedStr != remainderStr {

		t.Errorf("\n%v\n"+
			"ERROR: sMech.ExtractTextLines()\n"+
			"Expected 'remainderStr' DID NOT EQUAL Actual 'remainderStr'.\n"+
			"Expected 'remainderStr' = '%v'\n"+
			"  Actual 'remainderStr' = '%v'\n",
			ePrefix.String(),
			expectedStr,
			remainderStr)

		return
	}

	if numOfTxtLines != 3 {

		t.Errorf("\n%v\n"+
			"ERROR: sMech.ExtractTextLines()\n"+
			"Expected 'numOfTxtLines' DID NOT EQUAL '3'.\n"+
			"Expected 'numOfTxtLines' = '3'\n"+
			"  Actual 'numOfTxtLines' = '%v'\n",
			ePrefix.String(),
			numOfTxtLines)

		return
	}

	if len(timerLineStrs) != 3 {

		t.Errorf("\n%v\n"+
			"ERROR: sMech.ExtractTextLines()\n"+
			"Expected length of 'timerLineStrs' DID NOT EQUAL '3'.\n"+
			"Expected length of 'timerLineStrs' = '3'\n"+
			"  Actual length of 'timerLineStrs' = '%v'\n",
			ePrefix.String(),
			numOfTxtLines)

		return
	}

	return
}

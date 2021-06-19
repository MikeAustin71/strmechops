package strmech

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestStrMech_MakeSingleCharString_01(t *testing.T) {

	ePrefix := "TestStrMech_MakeSingleCharString_01() "

	sUtil := StrMech{}
	requestedLen := 20

	charRune := '*'

	outputStr, err := sUtil.MakeSingleCharString(
		charRune,
		requestedLen,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by sUtil.MakeSingleCharString(charRune, 10). "+
			"Error='%v' ", err.Error())
		return
	}

	outputStrLen := len(outputStr)

	if requestedLen != outputStrLen {
		t.Errorf("Error: Expected outputStr length='%v'. Instead, string length='%v'",
			requestedLen, outputStrLen)
	}

	for i := 0; i < outputStrLen; i++ {
		if rune(outputStr[i]) != charRune {
			t.Errorf("Error: outputStr rune at index='%v' DOES NOT MATCH "+
				"specified rune '%v'. Actual rune='%v' ", i, charRune, rune(outputStr[i]))
		}

	}

}

func TestStrMech_MakeSingleCharString_02(t *testing.T) {

	ePrefix := "TestStrMech_MakeSingleCharString_02() "

	sUtil := StrMech{}
	requestedLen := 100

	charRune := '='

	outputStr, err := sUtil.MakeSingleCharString(
		charRune,
		requestedLen,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by sUtil.MakeSingleCharString(charRune, 10). "+
			"Error='%v' ", err.Error())
		return
	}

	outputStrLen := len(outputStr)

	if requestedLen != outputStrLen {
		t.Errorf("Error: Expected outputStr length='%v'. Instead, string length='%v'",
			requestedLen, outputStrLen)
	}

	for i := 0; i < outputStrLen; i++ {
		if rune(outputStr[i]) != charRune {
			t.Errorf("Error: outputStr rune at index='%v' DOES NOT MATCH "+
				"specified rune '%v'. Actual rune='%v' ", i, charRune, rune(outputStr[i]))
		}

	}

}

func TestStrMech_MakeSingleCharString_03(t *testing.T) {

	ePrefix := "TestStrMech_MakeSingleCharString_03() "

	sUtil := StrMech{}
	requestedLen := 20

	charRune := rune(0)

	_, err := sUtil.MakeSingleCharString(
		charRune,
		requestedLen,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return from sUtil.MakeSingleCharString(charRune, 10)\n" +
			"because 'charRune'==0.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestStrMech_MakeSingleCharString_04(t *testing.T) {

	ePrefix := "TestStrMech_MakeSingleCharString_04() "

	sUtil := StrMech{}
	requestedLen := 0

	charRune := 'x'

	_, err := sUtil.MakeSingleCharString(
		charRune,
		requestedLen,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return from sUtil.MakeSingleCharString(charRune, 10)\n" +
			"because 'requestedLen'< 1.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestStrMech_Read_01(t *testing.T) {

	expected := "original base string"
	lenExpected := len(expected)

	p := make([]byte, 100)

	s1 := StrMech{}
	s1.SetStringData(expected)

	n, err := s1.Read(p)

	if err != nil && err != io.EOF {
		t.Errorf("Error returned by s1.Read(p). Error='%v' ", err.Error())
	}

	actualStr := string(p[:n])

	if expected != actualStr {
		t.Errorf("Error: Expected StrOut='%v'. Instead, StrOut='%v' ",
			expected, actualStr)
	}

	if lenExpected != n {
		t.Errorf("Error: Expected bytes read n='%v'. Instead, n='%v' ",
			lenExpected, n)
	}
}

func TestStrMech_Read_02(t *testing.T) {

	expected := "Original sops1 base string"
	lenExpected := len(expected)

	p := make([]byte, 5, 15)

	s1 := StrMech{}.NewPtr()
	s1.SetStringData(expected)
	n := 0
	var err error
	err = nil
	b := strings.Builder{}
	b.Grow(len(expected) + 150)

	for err != io.EOF {

		n, err = s1.Read(p)

		if err != nil && err != io.EOF {
			fmt.Printf("Error returned by s1.Read(p). "+
				"Error='%v' \n", err.Error())
			return
		}

		b.Write(p[:n])

		for i := 0; i < len(p); i++ {
			p[i] = byte(0)
		}

	}

	actualStr := b.String()

	if expected != actualStr {
		t.Errorf("Error: Expected StrOut='%v'. Instead, StrOut='%v' ",
			expected, actualStr)
	}

	lenActual := len(actualStr)

	if lenExpected != lenActual {
		t.Errorf("Error: Expected bytes read ='%v'. Instead, bytes read='%v' ",
			lenExpected, lenActual)
	}
}

func TestStrMech_Read_03(t *testing.T) {

	expected := "Original sops1 base string"
	lenExpected := int64(len(expected))

	s1 := StrMech{}.NewPtr()
	s1.SetStringData(expected)

	s2 := StrMech{}.NewPtr()

	n, err := io.Copy(s2, s1)

	if err != nil {
		fmt.Printf("Error returned by io.Copy(sops2, sops1). "+
			"Error='%v' \n", err.Error())
		return
	}

	actualData := s2.GetStringData()

	if expected != actualData {
		t.Errorf("Error: Expected StrOut='%v'. Instead, String Data='%v' ",
			expected, actualData)
	}

	if lenExpected != n {
		t.Errorf("Error: Expected bytes read ='%v'. Instead, bytes read='%v' ",
			lenExpected, n)
	}
}

func TestStrMech_Read_04(t *testing.T) {
	originalStr := "Hello World"

	sops1 := StrMech{}.NewPtr()
	sops1.SetStringData(originalStr)
	p := make([]byte, 0)

	_, err := sops1.Read(p)

	if err == nil {
		t.Error("Error: Expected error return. NO ERROR WAS RETURNED!")
	}

}

func TestStrMech_ReadStringFromBytes_01(t *testing.T) {

	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\n',
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\n',
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r'}

	expectedStr := "Hello World"
	expectedNextIdx := 13

	sMech := StrMech{}

	result, nextStartIdx := sMech.ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_02(t *testing.T) {

	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\n',
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\n',
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r'}

	expectedStr := "Does your program run?"
	expectedNextIdx := -1

	var result string

	_, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	_, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	result, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_03(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11   12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\n',
		//
		//13  14  15  16  17  18  19  20  21  22  23  24  25
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\n',
		//26  27  28  29  30  31  32  33  34  35  36  37  38  39  40  41  42  43  44  45  46  47   48
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r'}

	expectedStr := "How are you?"
	expectedNextIdx := 26
	var result string

	_, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	result, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_04(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r',
		//
		//12  13  14  15  16  17  18  19  20  21  22  23  24
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\n',
		//25  26  27  28  29  30  31  32  33  34  35  36  34  38  39  40  41  42  43  44  45  46   47
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r'}

	expectedStr := "How are you?"
	expectedNextIdx := 25
	var result string

	_, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	result, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'.\nInstead, result='%v'\n",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'.\nInstead, nextStartIdx='%v'\n",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_05(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', ',', ' ',
		//
		//13  14  15  16  17  18  19  20  21  22  23  24
		'h', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?'}

	expectedStr := "Hello World, how are you?"
	expectedNextIdx := -1

	result, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_06(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', ',', ' ',
		//
		//13  14  15  16  17  18  19  20  21  22  23  24
		'h', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?'}

	expectedStr := ""
	expectedNextIdx := -1
	var result string

	_, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	result, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_07(t *testing.T) {

	var bytes []byte

	expectedStr := ""
	expectedNextIdx := -1

	result, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_08(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11   12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\v',
		//
		//13  14  15  16  17  18  19  20  21  22  23  24   25   26
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\r', '\n',
		//27  28  29  30  31  32  33  34  35  36  34  38  39  40  41  42  43  44  45  46   47 48   49   50
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r', '\n'}

	expectedStr := "Does your program run?"
	expectedNextIdx := -1
	var result string

	_, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	_, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	result, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'.\n Instead, result='%v'\n",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'.\nInstead, nextStartIdx='%v'\n",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_09(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11   12   13
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\v', '\n',
		//
		//14  15  16  17  18  19  20  21  22  23  24   25   26  27   28
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?', '\r', '\v', '\n',
		//29  30  31  32  33  34  35  36  34  38  39  40  41  42  43  44  45  46  47  48  49  50   51   52
		'D', 'o', 'e', 's', ' ', 'y', 'o', 'u', 'r', ' ', 'p', 'r', 'o', 'g', 'r', 'a', 'm', ' ', 'r', 'u', 'n', '?', '\r', '\n'}

	expectedStr := "Does your program run?"
	expectedNextIdx := -1
	var result string

	_, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	_, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	result, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_10(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10   11   12   13
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '\r', '\v', '\n',
		//
		//14  15  16  17  18  19  20  21  22  23  24   25
		'H', 'o', 'w', ' ', 'a', 'r', 'e', ' ', 'y', 'o', 'u', '?'}

	expectedStr := "How are you?"
	expectedNextIdx := -1
	var result string

	_, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	result, nextStartIdx = StrMech{}.Ptr().ReadStringFromBytes(bytes, nextStartIdx)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_11(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}

	expectedStr := "Hello World!"
	expectedNextIdx := -1

	result, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_12(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!', '\n'}

	expectedStr := "Hello World!"
	expectedNextIdx := -1

	result, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_13(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!', '\r'}

	expectedStr := "Hello World!"
	expectedNextIdx := -1

	result, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_ReadStringFromBytes_14(t *testing.T) {

	//               0   1   2   3   4   5   6   7   8   9  10  11  12
	bytes := []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!', '\v'}

	expectedStr := "Hello World!"
	expectedNextIdx := -1

	result, nextStartIdx := StrMech{}.Ptr().ReadStringFromBytes(bytes, 0)

	if expectedStr != result {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, result)
	}

	if expectedNextIdx != nextStartIdx {
		t.Errorf("Error: Expected nextStartIdx='%v'. Instead, nextStartIdx='%v'",
			expectedNextIdx, nextStartIdx)
	}

}

func TestStrMech_RemoveStringChar_01(t *testing.T) {

	ePrefix := "TestStrMech_RemoveStringChar_01() "
	helloWorld := "XHelloX World!XX HowXX the hellXX are you?X"
	expectedDeletions := 9
	expectedStr := "Hello World! How the hell are you?"

	sops := StrMech{}

	actualStr,
		actualNoDeletions,
		err := sops.RemoveStringChar(
		helloWorld,
		'X',
		-1,
		ePrefix+"\n"+
			"Removing 'X'\n")

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if actualStr != expectedStr {
		t.Errorf("Error:\n"+
			"Expected string='%v'\n"+
			"  Actual string='%v'\n",
			expectedStr, actualStr)
	}

	if actualNoDeletions != expectedDeletions {
		t.Errorf("Error: Expected %v-deletions.\n"+
			"Instead, there were %v-deletions.\n",
			expectedDeletions,
			actualNoDeletions)
	}

}

func TestStrMech_RemoveStringChar_02(t *testing.T) {

	ePrefix := "TestStrMech_RemoveStringChar_02() "
	helloWorld := "XHelloX World!XX HowXX the hellXX are you?X"
	expectedDeletions := 8
	expectedStr := "Hello World! How the hell are you?X"

	sops := StrMech{}

	actualStr,
		actualNoDeletions,
		err := sops.RemoveStringChar(
		helloWorld,
		'X',
		8,
		ePrefix+"\n"+
			"Removing 'X' 8-times\n")

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if actualStr != expectedStr {
		t.Errorf("Error:\n"+
			"Expected string='%v'\n"+
			"  Actual string='%v'\n",
			expectedStr, actualStr)
	}

	if actualNoDeletions != expectedDeletions {
		t.Errorf("Error: Expected %v-deletions.\n"+
			"Instead, there were %v-deletions.\n",
			expectedDeletions,
			actualNoDeletions)
	}

}

func TestStrMech_RemoveStringChar_03(t *testing.T) {

	ePrefix := "TestStrMech_RemoveStringChar_03() "
	helloWorld := "XHelloX World!XX HowXX the hellXX are you?X"

	sops := StrMech{}

	_,
		_,
		err := sops.RemoveStringChar(
		helloWorld,
		'X',
		0,
		ePrefix+"\n"+
			"Error Test: 'maxNumOfCharDeletions' == 0\n")

	if err == nil {
		t.Error("Expected an error return from sops.RemoveStringChar()\n" +
			"because the 'maxNumOfCharDeletions' is set to zero.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestStrMech_RemoveStringChar_04(t *testing.T) {

	ePrefix := "TestStrMech_RemoveStringChar_04() "
	helloWorld := "XHelloX World!XX HowXX the hellXX are you?X"

	sops := StrMech{}

	_,
		_,
		err := sops.RemoveStringChar(
		helloWorld,
		0,
		0,
		ePrefix+"\n"+
			"Error Test: 'charToRemove' == 0\n")

	if err == nil {
		t.Error("Expected an error return from sops.RemoveStringChar()\n" +
			"because the 'charToRemove' is set to zero.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestStrMech_RemoveStringChar_05(t *testing.T) {

	ePrefix := "TestStrMech_RemoveStringChar_05() "
	helloWorld := ""

	sops := StrMech{}

	_,
		_,
		err := sops.RemoveStringChar(
		helloWorld,
		'X',
		-1,
		ePrefix+"\n"+
			"Error Test: 'targetStr' is empty string!\n")

	if err == nil {
		t.Error("Expected an error return from sops.RemoveStringChar()\n" +
			"because the 'charToRemove' is set to zero.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestStrMech_ReplaceBytes_01(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_01() "

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "1A2B3C4D5E6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'a'
	replaceBytes[0][1] = 'A'

	replaceBytes[1][0] = 'b'
	replaceBytes[1][1] = 'B'

	replaceBytes[2][0] = 'c'
	replaceBytes[2][1] = 'C'

	replaceBytes[3][0] = 'd'
	replaceBytes[3][1] = 'D'

	replaceBytes[4][0] = 'e'
	replaceBytes[4][1] = 'E'

	sMech := StrMech{}

	actualRunes, err := sMech.ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_ReplaceBytes_02(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_02() "

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "1A23C45E6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'a'
	replaceBytes[0][1] = 'A'

	replaceBytes[1][0] = 'b'
	replaceBytes[1][1] = 0

	replaceBytes[2][0] = 'c'
	replaceBytes[2][1] = 'C'

	replaceBytes[3][0] = 'd'
	replaceBytes[3][1] = 0

	replaceBytes[4][0] = 'e'
	replaceBytes[4][1] = 'E'

	actualRunes, err := StrMech{}.Ptr().ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_ReplaceBytes_03(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_03() "

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "1a2b3c4d5e6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'z'
	replaceBytes[0][1] = 'Z'

	replaceBytes[1][0] = 'y'
	replaceBytes[1][1] = 'Y'

	replaceBytes[2][0] = 'x'
	replaceBytes[2][1] = 'X'

	replaceBytes[3][0] = 'w'
	replaceBytes[3][1] = 'W'

	replaceBytes[4][0] = 'v'
	replaceBytes[4][1] = 'V'

	actualRunes, err := StrMech{}.Ptr().ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrMech_ReplaceBytes_04(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_04() "

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "3a4b5c6d7e6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = '1'
	replaceBytes[0][1] = '3'

	replaceBytes[1][0] = '2'
	replaceBytes[1][1] = '4'

	replaceBytes[2][0] = '3'
	replaceBytes[2][1] = '5'

	replaceBytes[3][0] = '4'
	replaceBytes[3][1] = '6'

	replaceBytes[4][0] = '5'
	replaceBytes[4][1] = '7'

	actualRunes, err := StrMech{}.Ptr().ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrMech_ReplaceBytes_05(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_05() "

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "1a23c4d5e6"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'z'
	replaceBytes[0][1] = 'Z'

	replaceBytes[1][0] = 'y'
	replaceBytes[1][1] = 'Y'

	replaceBytes[2][0] = 'x'
	replaceBytes[2][1] = 'X'

	replaceBytes[3][0] = 'w'
	replaceBytes[3][1] = 'W'

	replaceBytes[4][0] = 'b'
	replaceBytes[4][1] = 0

	actualRunes, err := StrMech{}.Ptr().ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrMech_ReplaceBytes_06(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_06() "

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	expected := "123456"

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'a'
	replaceBytes[0][1] = 0

	replaceBytes[1][0] = 'b'
	replaceBytes[1][1] = 0

	replaceBytes[2][0] = 'c'
	replaceBytes[2][1] = 0

	replaceBytes[3][0] = 'd'
	replaceBytes[3][1] = 0

	replaceBytes[4][0] = 'e'
	replaceBytes[4][1] = 0

	actualRunes, err := StrMech{}.Ptr().ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceBytes(testBytes, replaceBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected actual result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrMech_ReplaceBytes_07(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_07() "

	testBytes := make([]byte, 0, 0)

	replaceBytes := make([][]byte, 5, 10)

	for i := 0; i < 5; i++ {
		replaceBytes[i] = make([]byte, 2, 5)
	}

	replaceBytes[0][0] = 'a'
	replaceBytes[0][1] = 0

	replaceBytes[1][0] = 'b'
	replaceBytes[1][1] = 0

	replaceBytes[2][0] = 'c'
	replaceBytes[2][1] = 0

	replaceBytes[3][0] = 'd'
	replaceBytes[3][1] = 0

	replaceBytes[4][0] = 'e'
	replaceBytes[4][1] = 0

	_, err := StrMech{}.Ptr().ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err == nil {
		t.Error("Error: Expected an error return. NO ERROR RETURNED!. ")
	}
}

func TestStrMech_ReplaceBytes_08(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_08() "

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	replaceBytes := make([][]byte, 0, 0)

	_, err := StrMech{}.Ptr().ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err == nil {
		t.Error("Error: Expected an error return. NO ERROR WAS RETURNED! ")
	}

}

func TestStrMech_ReplaceBytes_09(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceBytes_09() "

	testStr := "1a2b3c4d5e6"
	testBytes := []byte(testStr)

	replaceBytes := make([][]byte, 5, 10)

	_, err := StrMech{}.Ptr().ReplaceBytes(
		testBytes,
		replaceBytes,
		ePrefix)

	if err == nil {
		t.Errorf("Error: Expected error return. NO ERROR WAS RETURNED!")
	}
}

func TestStrMech_ReplaceMultipleStrs_01(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceMultipleStrs_01() "

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "Hello World"

	rStrs[0][0] = "o"
	rStrs[0][1] = "x"
	rStrs[1][0] = " "
	rStrs[1][1] = "J"
	rStrs[2][0] = "l"
	rStrs[2][1] = "F"

	expectedStr := "HeFFxJWxrFd"

	sMech := StrMech{}

	actualStr, err := sMech.ReplaceMultipleStrs(
		testStr,
		rStrs,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrMech_ReplaceMultipleStrs_02(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceMultipleStrs_02() "

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "Hello World"

	rStrs[0][0] = "o"
	rStrs[0][1] = ""
	rStrs[1][0] = " "
	rStrs[1][1] = ""
	rStrs[2][0] = "l"
	rStrs[2][1] = ""

	expectedStr := "HeWrd"

	actualStr, err := StrMech{}.Ptr().ReplaceMultipleStrs(
		testStr,
		rStrs,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrMech_ReplaceMultipleStrs_03(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceMultipleStrs_03() "

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "Hello World"

	rStrs[0][0] = "f"
	rStrs[0][1] = " "
	rStrs[1][0] = "j"
	rStrs[1][1] = "r"
	rStrs[2][0] = "M"
	rStrs[2][1] = "x"

	expectedStr := "Hello World"

	actualStr, err := StrMech{}.Ptr().ReplaceMultipleStrs(
		testStr,
		rStrs,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrMech_ReplaceMultipleStrs_04(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceMultipleStrs_04() "

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "Hello World Hello World"

	rStrs[0][0] = "o"
	rStrs[0][1] = "x"
	rStrs[1][0] = " "
	rStrs[1][1] = "J"
	rStrs[2][0] = "l"
	rStrs[2][1] = "F"

	expectedStr := "HeFFxJWxrFdJHeFFxJWxrFd"

	actualStr, err := StrMech{}.Ptr().ReplaceMultipleStrs(
		testStr,
		rStrs,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrMech_ReplaceMultipleStrs_05(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceMultipleStrs_05() "

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := ""

	rStrs[0][0] = "o"
	rStrs[0][1] = "x"
	rStrs[1][0] = " "
	rStrs[1][1] = "J"
	rStrs[2][0] = "l"
	rStrs[2][1] = "F"

	_, err := StrMech{}.Ptr().ReplaceMultipleStrs(
		testStr,
		rStrs,
		ePrefix)

	if err == nil {
		t.Error("Expected error return from StrMech{}.Ptr().ReplaceMultipleStrs(testStr, rStrs)\n" +
			"because 'testStr' is a zero length string.\n" +
			"However, NO ERROR WAS RETURNED!\n")
	}
}

func TestStrMech_ReplaceMultipleStrs_06(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceMultipleStrs_06() "

	testStr := "Hello World"

	rStrs := make([][]string, 0)

	_, err := StrMech{}.Ptr().ReplaceMultipleStrs(
		testStr,
		rStrs,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return from StrMech{}.Ptr().ReplaceMultipleStrs(testStr, rStrs)\n" +
			"because 'rStrs' is a zero length array.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestStrMech_ReplaceMultipleStrs_07(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceMultipleStrs_07() "

	testStr := "Hello World"

	rStrs := make([][]string, 35)

	_, err := StrMech{}.Ptr().ReplaceMultipleStrs(
		testStr,
		rStrs,
		ePrefix)

	if err == nil {
		t.Error("Expected an error return from StrMech{}.Ptr().ReplaceMultipleStrs(testStr, rStrs)\n" +
			"because 'rStrs' is a 1-dimensional array.\n" +
			"However, NO ERROR WAS RETURNED!!!\n")
	}
}

func TestStrMech_ReplaceMultipleStrs_08(t *testing.T) {

	ePrefix := "TestStrMech_ReplaceMultipleStrs_08() "

	rStrs := make([][]string, 3, 5)

	for i := 0; i < 3; i++ {
		rStrs[i] = make([]string, 2, 5)
	}

	testStr := "HeFFxJWxrFd"

	rStrs[0][0] = "x"
	rStrs[0][1] = "o"
	rStrs[1][0] = "J"
	rStrs[1][1] = " "
	rStrs[2][0] = "F"
	rStrs[2][1] = "l"

	expectedStr := "Hello World"

	actualStr, err := StrMech{}.Ptr().ReplaceMultipleStrs(
		testStr,
		rStrs,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.Ptr().ReplaceMultipleStrs(testStr, rStrs). "+
			"Error='%v' ", err.Error())
	}

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expectedStr, actualStr)
	}

}

func TestStrMech_ReplaceNewLines_01(t *testing.T) {

	testStr := "Hello\nWorld"
	replaceStr := " "
	expectedStr := "Hello World"

	actualStr := StrMech{}.Ptr().ReplaceNewLines(
		testStr,
		replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrMech_ReplaceNewLines_02(t *testing.T) {

	testStr := "Hello World"
	replaceStr := " "
	expectedStr := "Hello World"

	actualStr := StrMech{}.Ptr().ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

}

func TestStrMech_ReplaceNewLines_03(t *testing.T) {

	testStr := "\n\nHello\nWorld\n\n\n"
	replaceStr := ""
	expectedStr := "HelloWorld"
	lenExpectedStr := len(expectedStr)

	actualStr := StrMech{}.Ptr().ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrMech_ReplaceNewLines_04(t *testing.T) {

	testStr := "\n\nHello World"
	replaceStr := ""
	expectedStr := "Hello World"

	actualStr := StrMech{}.Ptr().ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrMech_ReplaceNewLines_05(t *testing.T) {

	testStr := "Hello World\n\n"
	replaceStr := ""
	expectedStr := "Hello World"

	actualStr := StrMech{}.Ptr().ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrMech_ReplaceNewLines_06(t *testing.T) {

	testStr := "Hello World\n"
	replaceStr := ""
	expectedStr := "Hello World"

	actualStr := StrMech{}.Ptr().ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrMech_ReplaceNewLines_07(t *testing.T) {

	testStr := "\nHello World"
	replaceStr := ""
	expectedStr := "Hello World"

	actualStr := StrMech{}.Ptr().ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrMech_ReplaceNewLines_08(t *testing.T) {

	testStr := "\tHello World"
	replaceStr := ""
	expectedStr := "\tHello World"

	sMech := StrMech{}

	actualStr := sMech.ReplaceNewLines(testStr, replaceStr)

	if expectedStr != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'",
			expectedStr, actualStr)
	}

	lenExpectedStr := len(expectedStr)

	lenActualStr := len(actualStr)

	if lenExpectedStr != lenActualStr {
		t.Errorf("Error: Expected actual length='%v'. Instead, actual length='%v'",
			lenExpectedStr, lenActualStr)
	}

}

func TestStrMech_ReplaceNewLines_09(t *testing.T) {

	testStr := ""
	replaceStr := "XX"

	sMech := StrMech{}

	actualStr := sMech.ReplaceNewLines(testStr, replaceStr)

	lenActualStr := len(actualStr)

	if lenActualStr != 0 {
		t.Errorf("Error: Expected actual length='0' because\n"+
			"'testStr' is an empty string.\n"+
			"Instead, actual length='%v'\n",
			lenActualStr)
	}

}

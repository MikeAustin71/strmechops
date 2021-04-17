package strmech

import (
	"io"
	"testing"
)

func TestStrOps_FindRegExIndex_01(t *testing.T) {

	regex := "\\d:\\d:\\d"
	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	expected := "1:6:3"
	su := StrOps{}

	idx := su.FindRegExIndex(targetStr, regex)

	if idx == nil {
		t.Errorf("Error: Did not locate Regular Expression,'%v', in 'targetStr', '%v'.",
			regex, targetStr)
		return
	}

	sExtract := targetStr[idx[0]:idx[1]]

	if expected != sExtract {
		t.Errorf("Error: Expected regular expression match on string='%v'. "+
			"Instead, matched string='%v'. ", expected, sExtract)
	}
}

func TestStrOps_GetReader_01(t *testing.T) {
	originalStr := "Now is the time for all good men to come to the aid of their country."
	s1 := StrOps{}.NewPtr()
	s1.SetStringData(originalStr)
	s2 := StrOps{}.NewPtr()
	rdr := s1.GetReader()
	n, err := io.Copy(s2, rdr)

	if err != nil {
		t.Errorf("Error returned by io.Copy(s2, s1.GetReader()). "+
			"Error='%v' ", err.Error())
	}

	actualStr := s2.GetStringData()

	if originalStr != actualStr {
		t.Errorf("Error: Expected actualStr='%v'. Instead, actualStr='%v'",
			originalStr, actualStr)
	}

	if int64(len(originalStr)) != n {
		t.Errorf("Error: Expected characters read='%v'. Instead, "+
			"characters read='%v' ",
			len(originalStr), n)
	}

}

func TestStrOps_GetReader_02(t *testing.T) {
	originalStr := "xx"
	s1 := StrOps{}.NewPtr()
	s1.SetStringData(originalStr)
	s2 := StrOps{}.NewPtr()
	rdr := s1.GetReader()
	n, err := io.Copy(s2, rdr)

	if err != nil {
		t.Errorf("Error returned by io.Copy(s2, s1.GetReader()). "+
			"Error='%v' ", err.Error())
	}

	actualStr := s2.GetStringData()

	if originalStr != actualStr {
		t.Errorf("Error: Expected actualStr='%v'. Instead, actualStr='%v'",
			originalStr, actualStr)
	}

	if int64(len(originalStr)) != n {
		t.Errorf("Error: Expected characters read='%v'. Instead, "+
			"characters read='%v' ",
			len(originalStr), n)
	}

}

func TestStrOps_GetSoftwareVersion_01(t *testing.T) {
	expectedStr := "3.0.0"
	actualStr := StrOps{}.GetSoftwareVersion()

	if expectedStr != actualStr {
		t.Errorf("Error: Expected Software Version string='%v'.\n"+
			"Instead, Software Version string='%v'.\n",
			expectedStr, actualStr)
	}
}

func TestStrOps_GetValidBytes_01(t *testing.T) {

	ePrefix := "TestStrOps_GetValidBytes_01() "

	validBytes := []byte{'v', 'a', 'l', 'i', 'd'}

	testBytes := []byte{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}

	expected := "valid"

	actualBytes, err := StrOps{}.Ptr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_02(t *testing.T) {

	ePrefix := "TestStrOps_GetValidBytes_02() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', '1', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "1355"

	actualBytes, err := StrOps{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_03(t *testing.T) {

	ePrefix := "TestStrOps_GetValidBytes_03() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "355"

	actualBytes, err := StrOps{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_04(t *testing.T) {

	ePrefix := "TestStrOps_GetValidBytes_04() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', 'J', 'm', '!', 'a', 'J', '%', 'Z', 'i', 'F', 'd', '^'}

	expected := ""

	actualBytes, err := StrOps{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_05(t *testing.T) {

	ePrefix := "TestStrOps_GetValidBytes_05() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	expected := ""

	actualBytes, err := StrOps{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidBytes_06(t *testing.T) {

	ePrefix := "TestStrOps_GetValidBytes_06() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := make([]byte, 0, 5)

	_, err := StrOps{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testBytes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidBytes_07(t *testing.T) {

	ePrefix := "TestStrOps_GetValidBytes_07() "

	validBytes := make([]byte, 0, 5)

	testBytes := []byte{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	_, err := StrOps{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err == nil {
		t.Error("Expected Error return due to empty 'validBytes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidRunes_01(t *testing.T) {

	ePrefix := "TestStrOps_GetValidRunes_01() "

	validRunes := []rune{'v', 'a', 'l', 'i', 'd'}

	testRunes := []rune{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}

	expected := "valid"

	actualRunes, err := StrOps{}.NewPtr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_02(t *testing.T) {

	ePrefix := "TestStrOps_GetValidRunes_02() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', '1', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "1355"

	actualRunes, err := StrOps{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_03(t *testing.T) {

	ePrefix := "TestStrOps_GetValidRunes_03() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "355"

	actualRunes, err := StrOps{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_04(t *testing.T) {

	ePrefix := "TestStrOps_GetValidRunes_04() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', 'J', 'm', '!', 'a', 'J', '%', 'Z', 'i', 'F', 'd', '^'}

	expected := ""

	actualRunes, err := StrOps{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_05(t *testing.T) {

	ePrefix := "TestStrOps_GetValidRunes_05() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	expected := ""

	actualRunes, err := StrOps{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidRunes_06(t *testing.T) {

	ePrefix := "TestStrOps_GetValidRunes_06() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := make([]rune, 0, 5)

	_, err := StrOps{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testRunes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidRunes_07(t *testing.T) {

	ePrefix := "TestStrOps_GetValidRunes_01() "

	validRunes := make([]rune, 0, 5)

	testRunes := []rune{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	_, err := StrOps{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err == nil {
		t.Error("Expected Error return due to empty 'validRunes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidString_01(t *testing.T) {

	ePrefix := "TestStrOps_GetValidString_01() "

	validRunes := []rune{'v', 'a', 'l', 'i', 'd'}

	testStr := "xjvmRaJlZiFdS"

	expected := "valid"

	actualStr, err := StrOps{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_GetValidString_02(t *testing.T) {

	ePrefix := "TestStrOps_GetValidString_02() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "x13m5aJ7ZiFd5"

	expected := "1355"

	actualStr, err := StrOps{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidString_03(t *testing.T) {

	ePrefix := "TestStrOps_GetValidString_03() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xz3m5aJ7ZiFd5"

	expected := "355"

	actualStr, err := StrOps{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidString_04(t *testing.T) {

	ePrefix := "TestStrOps_GetValidString_04() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xzJm!aJ%ZiFd^"

	expected := ""

	actualStr, err := StrOps{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrOps_GetValidString_05(t *testing.T) {

	ePrefix := "TestStrOps_GetValidString_05() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xzUmMaJ9ZiFd&"

	expected := ""

	actualStr, err := StrOps{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrOps{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrOps_GetValidString_06(t *testing.T) {

	ePrefix := "TestStrOps_GetValidString_06() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := ""

	_, err := StrOps{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testStr'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrOps_GetValidString_07(t *testing.T) {

	ePrefix := "TestStrOps_GetValidString_07() "

	validRunes := make([]rune, 0, 5)

	testStr := "xzUmMaJ9ZiFd&"

	_, err := StrOps{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err == nil {
		t.Error("Expected Error return due to empty 'validRunes'. " +
			"NO ERROR WAS RETURNED!")
	}
}

func TestStrOps_IsEmptyOrWhiteSpace_01(t *testing.T) {

	testStr := "       "

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != true {
		t.Error("Error: Expected result='true'. Instead, result='false'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_02(t *testing.T) {

	testStr := ""

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != true {
		t.Error("Error: Expected result='true'. Instead, result='false'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_03(t *testing.T) {

	testStr := " xyz "

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_04(t *testing.T) {

	testStr := "xyz"

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_05(t *testing.T) {

	testStr := "/t"

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_06(t *testing.T) {

	testStr := "/n           "

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_07(t *testing.T) {

	testStr := "  /n"

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_08(t *testing.T) {

	testStr := "  x"

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_IsEmptyOrWhiteSpace_09(t *testing.T) {

	testStr := "x   "

	result := StrOps{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrOps_LowerCaseFirstLetter_01(t *testing.T) {

	testStr := "Now is the time for all good men to come to the aid of their country."

	expected := "now is the time for all good men to come to the aid of their country."

	actualStr := StrOps{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_LowerCaseFirstLetter_02(t *testing.T) {

	testStr := "  Now is the time for all good men to come to the aid of their country."

	expected := "  now is the time for all good men to come to the aid of their country."

	actualStr := StrOps{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_LowerCaseFirstLetter_03(t *testing.T) {

	testStr := "now is the time for all good men to come to the aid of their country."

	expected := "now is the time for all good men to come to the aid of their country."

	actualStr := StrOps{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_LowerCaseFirstLetter_04(t *testing.T) {

	testStr := "  now is the time for all good men to come to the aid of their country."

	expected := "  now is the time for all good men to come to the aid of their country."

	actualStr := StrOps{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrOps_LowerCaseFirstLetter_05(t *testing.T) {

	testStr := ""

	expected := ""

	actualStr := StrOps{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

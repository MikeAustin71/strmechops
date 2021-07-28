package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"testing"
)

func TestStrMech_FindRegExIndex_01(t *testing.T) {

	regex := "\\d:\\d:\\d"
	targetStr := "November 12, 2016 1:6:3pm +0000 UTC"
	expected := "1:6:3"
	su := StrMech{}

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

func TestStrMech_FindRunesInRunes_000100(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000100()"

	hostRunes := []rune("Now is the time for all good men" +
		"to come to the aid of their country.")

	targetRunes := []rune("is the time")

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		-1,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if foundIndex != 4 {
		t.Errorf("ERROR: Expected foundIndex==4\n"+
			"Instead foundIndex=='%v'\n",
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_000200(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000200()"

	hostRunes := []rune("Now is the time for all good men" +
		"to come to the aid of their country. If the dare to do so.")

	targetRunes := []rune("If the dare to do so.")

	expectedFoundIndex := len(hostRunes) - len(targetRunes)

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		10,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_000300(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000300()"

	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("cow")

	expectedFoundIndex := 14

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		10,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_000400(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000400()"

	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("How")

	expectedFoundIndex := 0

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		-1,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_000500(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000500()"

	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("How now brown cow!")

	expectedFoundIndex := -1

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		0,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_000600(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000600()"

	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	expectedFoundIndex := -1

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		5,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_000700(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000700()"

	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	expectedFoundIndex := 4

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		4,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_000800(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000800()"

	// Len 17
	var hostRunes []rune

	targetRunes := []rune("now")

	su := StrMech{}

	_,
		err := su.FindRunesInRunes(
		hostRunes,
		0,
		targetRunes,
		funcName)

	if err == nil {
		t.Error("Expected an error return from su.FindRunesInRunes()\n" +
			"because 'hostRunes' is an empty array.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestStrMech_FindRunesInRunes_000900(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_000900()"

	// Len 17
	hostRunes := []rune("How now brown cow!")

	var targetRunes []rune

	su := StrMech{}

	_,
		err := su.FindRunesInRunes(
		hostRunes,
		0,
		targetRunes,
		funcName)

	if err == nil {
		t.Error("Expected an error return from su.FindRunesInRunes()\n" +
			"because 'targetRunes' is an empty array.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestStrMech_FindRunesInRunes_001000(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_001000()"

	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	expectedFoundIndex := 4

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		-99,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_001100(t *testing.T) {

	funcName := "TestStrMech_FindRunesInRunes_001100()"

	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	expectedFoundIndex := -1

	su := StrMech{}

	foundIndex,
		err := su.FindRunesInRunes(
		hostRunes,
		99,
		targetRunes,
		funcName)

	if err != nil {
		t.Errorf("Error returned by su.FindRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_FindRunesInRunes_001200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestStrMech_FindRunesInRunes_001200()",
		"")

	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	expectedFoundIndex := 4

	sMechPreon := strMechPreon{}

	foundIndex,
		err := sMechPreon.findRunesInRunes(
		hostRunes,
		0,
		targetRunes,
		&ePrefix)

	if err != nil {
		t.Errorf("Error returned by sMechPreon.findRunesInRunes()\n"+
			"Error=\n%v\n",
			err.Error())
		return
	}

	if expectedFoundIndex != foundIndex {
		t.Errorf("ERROR: Expected foundIndex=='%v'\n"+
			"Instead foundIndex=='%v'\n",
			expectedFoundIndex,
			foundIndex)
	}

}

func TestStrMech_GetReader_01(t *testing.T) {
	originalStr := "Now is the time for all good men to come to the aid of their country."
	s1 := StrMech{}
	s1.SetStringData(originalStr)
	s2 := StrMech{}.NewPtr()
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

func TestStrMech_GetReader_02(t *testing.T) {
	originalStr := "xx"
	s1 := StrMech{}
	s1.SetStringData(originalStr)
	s2 := StrMech{}
	rdr := s1.GetReader()
	n, err := io.Copy(&s2, rdr)

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

func TestStrMech_GetValidBytes_01(t *testing.T) {

	ePrefix := "TestStrMech_GetValidBytes_01() "

	validBytes := []byte{'v', 'a', 'l', 'i', 'd'}

	testBytes := []byte{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}

	expected := "valid"

	sMech := StrMech{}
	actualBytes, err := sMech.GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidBytes_02(t *testing.T) {

	ePrefix := "TestStrMech_GetValidBytes_02() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', '1', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "1355"

	actualBytes, err := StrMech{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidBytes_03(t *testing.T) {

	ePrefix := "TestStrMech_GetValidBytes_03() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "355"

	actualBytes, err := StrMech{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidBytes_04(t *testing.T) {

	ePrefix := "TestStrMech_GetValidBytes_04() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', 'J', 'm', '!', 'a', 'J', '%', 'Z', 'i', 'F', 'd', '^'}

	expected := ""

	actualBytes, err := StrMech{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidBytes_05(t *testing.T) {

	ePrefix := "TestStrMech_GetValidBytes_05() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := []byte{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	expected := ""

	actualBytes, err := StrMech{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidBytes(testBytes, validBytes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualBytes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidBytes_06(t *testing.T) {

	ePrefix := "TestStrMech_GetValidBytes_06() "

	validBytes := []byte{'1', '2', '3', '4', '5'}

	testBytes := make([]byte, 0, 5)

	_, err := StrMech{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testBytes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrMech_GetValidBytes_07(t *testing.T) {

	ePrefix := "TestStrMech_GetValidBytes_07() "

	validBytes := make([]byte, 0, 5)

	testBytes := []byte{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	_, err := StrMech{}.NewPtr().GetValidBytes(
		testBytes,
		validBytes,
		ePrefix)

	if err == nil {
		t.Error("Expected Error return due to empty 'validBytes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrMech_GetValidRunes_01(t *testing.T) {

	ePrefix := "TestStrMech_GetValidRunes_01() "

	validRunes := []rune{'v', 'a', 'l', 'i', 'd'}

	testRunes := []rune{'x', 'j', 'v', 'm', 'R', 'a', 'J', 'l', 'Z', 'i', 'F', 'd', 'S'}

	expected := "valid"

	sMech := StrMech{}

	actualRunes, err := sMech.GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidRunes_02(t *testing.T) {

	ePrefix := "TestStrMech_GetValidRunes_02() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', '1', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "1355"

	actualRunes, err := StrMech{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidRunes_03(t *testing.T) {

	ePrefix := "TestStrMech_GetValidRunes_03() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', '3', 'm', '5', 'a', 'J', '7', 'Z', 'i', 'F', 'd', '5'}

	expected := "355"

	actualRunes, err := StrMech{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidRunes_04(t *testing.T) {

	ePrefix := "TestStrMech_GetValidRunes_04() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', 'J', 'm', '!', 'a', 'J', '%', 'Z', 'i', 'F', 'd', '^'}

	expected := ""

	actualRunes, err := StrMech{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidRunes_05(t *testing.T) {

	ePrefix := "TestStrMech_GetValidRunes_05() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := []rune{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	expected := ""

	actualRunes, err := StrMech{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidRunes(testRunes, validRunes). "+
			"Error='%v' ", err.Error())
	}

	actualStr := string(actualRunes)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidRunes_06(t *testing.T) {

	ePrefix := "TestStrMech_GetValidRunes_06() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testRunes := make([]rune, 0, 5)

	_, err := StrMech{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testRunes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrMech_GetValidRunes_07(t *testing.T) {

	ePrefix := "TestStrMech_GetValidRunes_01() "

	validRunes := make([]rune, 0, 5)

	testRunes := []rune{'x', 'z', 'U', 'm', 'M', 'a', 'J', '9', 'Z', 'i', 'F', 'd', '&'}

	_, err := StrMech{}.Ptr().GetValidRunes(
		testRunes,
		validRunes,
		ePrefix)

	if err == nil {
		t.Error("Expected Error return due to empty 'validRunes'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrMech_GetValidString_01(t *testing.T) {

	ePrefix := "TestStrMech_GetValidString_01() "

	validRunes := []rune{'v', 'a', 'l', 'i', 'd'}

	testStr := "xjvmRaJlZiFdS"

	expected := "valid"

	sMech := StrMech{}

	actualStr, err := sMech.GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrMech_GetValidString_02(t *testing.T) {

	ePrefix := "TestStrMech_GetValidString_02() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "x13m5aJ7ZiFd5"

	expected := "1355"

	actualStr, err := StrMech{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidString_03(t *testing.T) {

	ePrefix := "TestStrMech_GetValidString_03() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xz3m5aJ7ZiFd5"

	expected := "355"

	actualStr, err := StrMech{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidString_04(t *testing.T) {

	ePrefix := "TestStrMech_GetValidString_04() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xzJm!aJ%ZiFd^"

	expected := ""

	actualStr, err := StrMech{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}
}

func TestStrMech_GetValidString_05(t *testing.T) {

	ePrefix := "TestStrMech_GetValidString_05() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := "xzUmMaJ9ZiFd&"

	expected := ""

	actualStr, err := StrMech{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err != nil {
		t.Errorf("Error returned by StrMech{}.GetValidString(testStr, validRunes). "+
			"Error='%v' ", err.Error())
	}

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v'. ",
			expected, actualStr)
	}

}

func TestStrMech_GetValidString_06(t *testing.T) {

	ePrefix := "TestStrMech_GetValidString_06() "

	validRunes := []rune{'1', '2', '3', '4', '5'}

	testStr := ""

	_, err := StrMech{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err == nil {
		t.Error("Expected an Error Return due to empty 'testStr'. " +
			"NO ERROR WAS RETURNED!")
	}

}

func TestStrMech_GetValidString_07(t *testing.T) {

	ePrefix := "TestStrMech_GetValidString_07() "

	validRunes := make([]rune, 0, 5)

	testStr := "xzUmMaJ9ZiFd&"

	_, err := StrMech{}.NewPtr().GetValidString(
		testStr,
		validRunes,
		ePrefix)

	if err == nil {
		t.Error("Expected Error return due to empty 'validRunes'. " +
			"NO ERROR WAS RETURNED!")
	}
}

func TestStrMech_IsEmptyOrWhiteSpace_01(t *testing.T) {

	testStr := "       "

	sMech := StrMech{}

	result := sMech.IsEmptyOrWhiteSpace(testStr)

	if result != true {
		t.Error("Error: Expected result='true'. Instead, result='false'")
	}

}

func TestStrMech_IsEmptyOrWhiteSpace_02(t *testing.T) {

	testStr := ""

	result := StrMech{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != true {
		t.Error("Error: Expected result='true'. Instead, result='false'")
	}

}

func TestStrMech_IsEmptyOrWhiteSpace_03(t *testing.T) {

	testStr := " xyz "

	result := StrMech{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrMech_IsEmptyOrWhiteSpace_04(t *testing.T) {

	testStr := "xyz"

	result := StrMech{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrMech_IsEmptyOrWhiteSpace_05(t *testing.T) {

	testStr := "/t"

	result := StrMech{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrMech_IsEmptyOrWhiteSpace_06(t *testing.T) {

	testStr := "/n           "

	result := StrMech{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrMech_IsEmptyOrWhiteSpace_07(t *testing.T) {

	testStr := "  /n"

	result := StrMech{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrMech_IsEmptyOrWhiteSpace_08(t *testing.T) {

	testStr := "  x"

	result := StrMech{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrMech_IsEmptyOrWhiteSpace_09(t *testing.T) {

	testStr := "x   "

	result := StrMech{}.NewPtr().IsEmptyOrWhiteSpace(testStr)

	if result != false {
		t.Error("Error: Expected result='false'. Instead, result='true'")
	}

}

func TestStrMech_IsTargetRunesIndex_000100(t *testing.T) {
	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	sMech := StrMech{}

	hostStartIndex := 4

	foundTargetRunes := sMech.IsTargetRunesIndex(
		hostRunes,
		hostStartIndex,
		targetRunes)

	if foundTargetRunes == false {
		t.Errorf("Error:\n"+
			"Expected foundTargetRunes=='true' at index %v .\n"+
			"However, the search routine returned 'false'.\n",
			hostStartIndex)
	}

}

func TestStrMech_IsTargetRunesIndex_000200(t *testing.T) {
	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	sMech := StrMech{}

	hostStartIndex := 3

	foundTargetRunes := sMech.IsTargetRunesIndex(
		hostRunes,
		hostStartIndex,
		targetRunes)

	if foundTargetRunes == true {
		t.Errorf("Error:\n"+
			"Expected foundTargetRunes=='false' at index %v .\n"+
			"However, the search routine returned 'true'.\n",
			hostStartIndex)
	}

}

func TestStrMech_IsTargetRunesIndex_000300(t *testing.T) {
	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	sMech := StrMech{}

	hostStartIndex := 99

	foundTargetRunes := sMech.IsTargetRunesIndex(
		hostRunes,
		hostStartIndex,
		targetRunes)

	if foundTargetRunes == true {
		t.Errorf("Error:\n"+
			"Expected foundTargetRunes=='false' at index %v .\n"+
			"However, the search routine returned 'true'.\n",
			hostStartIndex)
	}
}

func TestStrMech_IsTargetRunesIndex_000400(t *testing.T) {
	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	sMech := StrMech{}

	hostStartIndex := -99

	foundTargetRunes := sMech.IsTargetRunesIndex(
		hostRunes,
		hostStartIndex,
		targetRunes)

	if foundTargetRunes == true {
		t.Errorf("Error:\n"+
			"Expected foundTargetRunes=='false' at index %v .\n"+
			"However, the search routine returned 'true'.\n",
			hostStartIndex)
	}
}

func TestStrMech_IsTargetRunesIndex_000500(t *testing.T) {
	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("How")

	sMech := StrMech{}

	hostStartIndex := -99

	foundTargetRunes := sMech.IsTargetRunesIndex(
		hostRunes,
		hostStartIndex,
		targetRunes)

	if foundTargetRunes == false {
		t.Errorf("Error:\n"+
			"Expected foundTargetRunes=='true' at index %v .\n"+
			"However, the search routine returned 'false'.\n",
			hostStartIndex)
	}
}

func TestStrMech_IsTargetRunesIndex_000600(t *testing.T) {
	// Len 17
	var hostRunes []rune

	targetRunes := []rune("How")

	sMech := StrMech{}

	hostStartIndex := 0

	foundTargetRunes := sMech.IsTargetRunesIndex(
		hostRunes,
		hostStartIndex,
		targetRunes)

	if foundTargetRunes == true {
		t.Errorf("Error:\n"+
			"Expected foundTargetRunes=='false' at index %v .\n"+
			"However, the search routine returned 'true'.\n",
			hostStartIndex)
	}
}

func TestStrMech_IsTargetRunesIndex_000700(t *testing.T) {
	// Len 17
	hostRunes := []rune("How now brown cow")

	var targetRunes []rune

	sMech := StrMech{}

	hostStartIndex := 0

	foundTargetRunes := sMech.IsTargetRunesIndex(
		hostRunes,
		hostStartIndex,
		targetRunes)

	if foundTargetRunes == true {
		t.Errorf("Error:\n"+
			"Expected foundTargetRunes=='false' at index %v .\n"+
			"However, the search routine returned 'true'.\n",
			hostStartIndex)
	}
}

func TestStrMech_IsTargetRunesIndex_000800(t *testing.T) {
	// Len 17
	hostRunes := []rune("How now brown cow")

	targetRunes := []rune("now")

	sMechPreon := strMechPreon{}

	hostStartIndex := 4

	foundTargetRunes := sMechPreon.isTargetRunesIndex(
		hostRunes,
		hostStartIndex,
		targetRunes)

	if foundTargetRunes == false {
		t.Errorf("Error:\n"+
			"Expected foundTargetRunes=='true' at index %v .\n"+
			"However, the search routine returned 'false'.\n",
			hostStartIndex)
	}

}

func TestStrMech_IsValidRuneCharArray_000100(t *testing.T) {

	ePrefix := "TestStrMech_IsValidRuneCharArray_000100()"

	charAry := []rune("abcdefghijklmnopqrstuvwxyz 1234567890 ABCDEFGHIJKLMNOPQRSTUVWXYZ!&*()@#$%^*_-+=-{}?.")

	sMech := StrMech{}

	isValid,
		err := sMech.IsValidRuneCharArray(
		charAry,
		ePrefix)

	if isValid == false {
		t.Errorf("%v\n"+
			"'isValid' returned from sMech.IsValidRuneCharArray()\n"+
			"is 'false'. IT SHOULD BE 'true'!\n",
			ePrefix)
	}

	if err != nil {
		t.Errorf("%v\n"+
			"'err' returned from sMech.IsValidRuneCharArray()\n"+
			"is NOT 'nil'. IT SHOULD BE 'nil'!\n"+
			"Error= '%v'\n",
			ePrefix,
			err.Error())
	}

	return
}

func TestStrMech_IsValidRuneCharArray_000200(t *testing.T) {

	ePrefix := "TestStrMech_IsValidRuneCharArray_000200()"

	var charAry []rune = nil

	sMech := StrMech{}

	isValid,
		err := sMech.IsValidRuneCharArray(
		charAry,
		ePrefix)

	if isValid == true {
		t.Errorf("%v\n"+
			"'isValid' returned from sMech.IsValidRuneCharArray()\n"+
			"is 'true'. IT SHOULD BE 'false' because\n"+
			"'charAry' is 'nil' !\n",
			ePrefix)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"'err' returned from sMech.IsValidRuneCharArray()\n"+
			"is 'nil'. IT SHOULD BE NOT 'nil' because\n"+
			"'charAry' is 'nil' !\n",
			ePrefix)
	}

	return
}

func TestStrMech_IsValidRuneCharArray_000300(t *testing.T) {

	ePrefix := "TestStrMech_IsValidRuneCharArray_000300()"

	charAry := make([]rune, 0)

	sMech := StrMech{}

	isValid,
		err := sMech.IsValidRuneCharArray(
		charAry,
		ePrefix)

	if isValid == true {
		t.Errorf("%v\n"+
			"'isValid' returned from sMech.IsValidRuneCharArray()\n"+
			"is 'true'. IT SHOULD BE 'false' because\n"+
			"'charAry' is a zero length array !\n",
			ePrefix)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"'err' returned from sMech.IsValidRuneCharArray()\n"+
			"is 'nil'. IT SHOULD BE NOT 'nil' because\n"+
			"'charAry' is a zero length array !\n",
			ePrefix)
	}

	return
}

func TestStrMech_IsValidRuneCharArray_000400(t *testing.T) {

	ePrefix := "TestStrMech_IsValidRuneCharArray_000400()"

	charAry := []rune{'H', 'e', 'l', 'l', 'o', 0, 'W', 'o', 'r', 'l', 'd'}

	sMech := StrMech{}

	isValid,
		err := sMech.IsValidRuneCharArray(
		charAry,
		ePrefix)

	if isValid == true {
		t.Errorf("%v\n"+
			"'isValid' returned from sMech.IsValidRuneCharArray()\n"+
			"is 'true'. IT SHOULD BE 'false' because\n"+
			"'charAry' contains an array element with a zero\n"+
			"value!\n",
			ePrefix)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"'err' returned from sMech.IsValidRuneCharArray()\n"+
			"is 'nil'. IT SHOULD BE NOT 'nil' because\n"+
			"'charAry' contains an array element with a zero\n"+
			"value!\n",
			ePrefix)
	}

	return
}

func TestStrMech_IsValidRuneIntArray_000100(t *testing.T) {

	ePrefix := "TestStrMech_IsValidRuneIntArray_000100()"

	intAry := []rune("1234567890")

	sMech := StrMech{}

	isValid,
		err := sMech.IsValidRuneIntArray(
		intAry,
		ePrefix)

	if isValid == false {
		t.Errorf("%v\n"+
			"'isValid' returned from sMech.IsValidRuneIntArray()\n"+
			"is 'false'. IT SHOULD BE 'true'!\n",
			ePrefix)
	}

	if err != nil {
		t.Errorf("%v\n"+
			"'err' returned from sMech.IsValidRuneIntArray()\n"+
			"is NOT 'nil'. IT SHOULD BE 'nil'!\n"+
			"Error= '%v'\n",
			ePrefix,
			err.Error())
	}

	return
}

func TestStrMech_IsValidRuneIntArray_000200(t *testing.T) {

	ePrefix := "TestStrMech_IsValidRuneIntArray_000200()"

	var intAry []rune = nil

	sMech := StrMech{}

	isValid,
		err := sMech.IsValidRuneIntArray(
		intAry,
		ePrefix)

	if isValid == true {
		t.Errorf("%v\n"+
			"'isValid' returned from sMech.IsValidRuneIntArray()\n"+
			"is 'true'. IT SHOULD BE 'false' because\n"+
			"'intAry' is 'nil' !\n",
			ePrefix)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"'err' returned from sMech.IsValidRuneIntArray()\n"+
			"is 'nil'. IT SHOULD BE NOT 'nil' because\n"+
			"'intAry' is 'nil' !\n",
			ePrefix)
	}

	return
}

func TestStrMech_IsValidRuneIntArray_000300(t *testing.T) {

	ePrefix := "TestStrMech_IsValidRuneIntArray_000300()"

	intAry := make([]rune, 0)

	sMech := StrMech{}

	isValid,
		err := sMech.IsValidRuneIntArray(
		intAry,
		ePrefix)

	if isValid == true {
		t.Errorf("%v\n"+
			"'isValid' returned from sMech.IsValidRuneIntArray()\n"+
			"is 'true'. IT SHOULD BE 'false' because\n"+
			"'intAry' is a zero length array !\n",
			ePrefix)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"'err' returned from sMech.IsValidRuneIntArray()\n"+
			"is 'nil'. IT SHOULD BE NOT 'nil' because\n"+
			"'intAry' is a zero length array !\n",
			ePrefix)
	}

	return
}

func TestStrMech_IsValidRuneIntArray_000400(t *testing.T) {

	ePrefix := "TestStrMech_IsValidRuneIntArray_000400()"

	intAry := []rune{'1', '2', '3', '4', '5', 0, '6', '7', '8', '9', '0'}

	sMech := StrMech{}

	isValid,
		err := sMech.IsValidRuneIntArray(
		intAry,
		ePrefix)

	if isValid == true {
		t.Errorf("%v\n"+
			"'isValid' returned from sMech.IsValidRuneIntArray()\n"+
			"is 'true'. IT SHOULD BE 'false' because\n"+
			"'intAry' contains an array element with a zero\n"+
			"value!\n",
			ePrefix)
	}

	if err == nil {
		t.Errorf("%v\n"+
			"'err' returned from sMech.IsValidRuneCharArray()\n"+
			"is 'nil'. IT SHOULD BE NOT 'nil' because\n"+
			"'intAry' contains an array element with a zero\n"+
			"value!\n",
			ePrefix)
	}

	return
}

func TestStrMech_LowerCaseFirstLetter_01(t *testing.T) {

	testStr := "Now is the time for all good men to come to the aid of their country."

	expected := "now is the time for all good men to come to the aid of their country."

	sMech := StrMech{}

	actualStr := sMech.LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrMech_LowerCaseFirstLetter_02(t *testing.T) {

	testStr := "  Now is the time for all good men to come to the aid of their country."

	expected := "  now is the time for all good men to come to the aid of their country."

	actualStr := StrMech{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrMech_LowerCaseFirstLetter_03(t *testing.T) {

	testStr := "now is the time for all good men to come to the aid of their country."

	expected := "now is the time for all good men to come to the aid of their country."

	actualStr := StrMech{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrMech_LowerCaseFirstLetter_04(t *testing.T) {

	testStr := "  now is the time for all good men to come to the aid of their country."

	expected := "  now is the time for all good men to come to the aid of their country."

	actualStr := StrMech{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

func TestStrMech_LowerCaseFirstLetter_05(t *testing.T) {

	testStr := ""

	expected := ""

	actualStr := StrMech{}.Ptr().LowerCaseFirstLetter(testStr)

	if expected != actualStr {
		t.Errorf("Error: Expected result='%v'. Instead, result='%v' ",
			expected, actualStr)
	}

}

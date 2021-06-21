package strmech

import (
	"strings"
	"testing"
)

func TestNumSignSymbolDisplayMode_XParseString_000100(t *testing.T) {

	testStr := "Explicit"

	nSignSymDisMode,
		err := NSignSymDisplayMode.XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NSignSymDisplayMode.XParseString(testStr, true).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if nSignSymDisMode != NumSignSymbolDisplayMode(0).Explicit() {
		t.Errorf("Error: Expected return of object='nSignSymDisMode.Explicit()'.\n"+
			"Instead, object integer value = '%v'\n",
			nSignSymDisMode.XValueInt())
	}

}

func TestNumSignSymbolDisplayMode_XParseString_000200(t *testing.T) {

	testStr := "Implicit"

	nSignSymDisMode,
		err := NSignSymDisplayMode.XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NSignSymDisplayMode.XParseString(testStr, true).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if nSignSymDisMode != NumSignSymbolDisplayMode(0).Implicit() {
		t.Errorf("Error: Expected return of object='nSignSymDisMode.Implicit()'.\n"+
			"Instead, object integer value = '%v'\n",
			nSignSymDisMode.XValueInt())
	}

}

func TestNumSignSymbolDisplayMode_XParseString_000300(t *testing.T) {

	testStr := "None"

	nSignSymDisMode,
		err := NSignSymDisplayMode.XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NSignSymDisplayMode.XParseString(testStr, true).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if nSignSymDisMode != NumSignSymbolDisplayMode(0).None() {
		t.Errorf("Error: Expected return of object='nSignSymDisMode.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			nSignSymDisMode.XValueInt())
	}

}

func TestNumSignSymbolDisplayMode_XParseString_000400(t *testing.T) {

	testStr := "explicit"

	nSignSymDisMode,
		err := NSignSymDisplayMode.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NSignSymDisplayMode.XParseString(testStr, false).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if nSignSymDisMode != NumSignSymbolDisplayMode(0).Explicit() {
		t.Errorf("Error: Expected return of object='nSignSymDisMode.Explicit()'.\n"+
			"Instead, object integer value = '%v'\n",
			nSignSymDisMode.XValueInt())
	}

}

func TestNumSignSymbolDisplayMode_XParseString_000500(t *testing.T) {

	testStr := "implicit"

	nSignSymDisMode,
		err := NSignSymDisplayMode.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NSignSymDisplayMode.XParseString(testStr, false).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if nSignSymDisMode != NumSignSymbolDisplayMode(0).Implicit() {
		t.Errorf("Error: Expected return of object='nSignSymDisMode.Implicit()'.\n"+
			"Instead, object integer value = '%v'\n",
			nSignSymDisMode.XValueInt())
	}

}

func TestNumSignSymbolDisplayMode_XParseString_000600(t *testing.T) {

	testStr := "none"

	nSignSymDisMode,
		err := NSignSymDisplayMode.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NSignSymDisplayMode.XParseString(testStr, false).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if nSignSymDisMode != NumSignSymbolDisplayMode(0).None() {
		t.Errorf("Error: Expected return of object='nSignSymDisMode.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			nSignSymDisMode.XValueInt())
	}

}

func TestNumSignSymbolDisplayMode_XParseString_000700(t *testing.T) {

	testStr := "none"

	_,
		err := NSignSymDisplayMode.XParseString(testStr, true)

	if err == nil {
		t.Error("Error\n" +
			"Expected an error return from NSignSymDisplayMode.XParseString(testStr, true).\n" +
			"because 'none' should fail a case sensitive search.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}

func TestNumSignSymbolDisplayMode_XParseString_000800(t *testing.T) {

	testStr := "none"

	_,
		err := NSignSymDisplayMode.XParseString(testStr, true)

	if err == nil {
		t.Error("Error\n" +
			"Expected an error return from NSignSymDisplayMode.XParseString(testStr, true).\n" +
			"because 'none' should fail a case sensitive search.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}

func TestNumSignSymbolDisplayMode_XParseString_000900(t *testing.T) {

	testStr := "no"

	_,
		err := NSignSymDisplayMode.XParseString(testStr, false)

	if err == nil {
		t.Error("Error\n" +
			"Expected an error return from NSignSymDisplayMode.XParseString(testStr, true).\n" +
			"because 'no' should fail a string length test.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}

func TestNumSignSymbolDisplayMode_XParseString_001000(t *testing.T) {

	testStr := "SomethingElse"

	_,
		err := NSignSymDisplayMode.XParseString(testStr, true)

	if err == nil {
		t.Error("Error\n" +
			"Expected an error return from NSignSymDisplayMode.XParseString(testStr, true).\n" +
			"because 'SomethingElse' should fail a map lookup test.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}

func TestNumSignSymbolDisplayMode_XParseString_001100(t *testing.T) {

	testStr := "SomethingElse"

	_,
		err := NSignSymDisplayMode.XParseString(testStr, false)

	if err == nil {
		t.Error("Error\n" +
			"Expected an error return from NSignSymDisplayMode.XParseString(testStr, true).\n" +
			"because 'SomethingElse' should fail a map lookup test.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")

		return
	}

}

func TestNumSignSymbolDisplayMode_XValue_000100(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).None()

	curValue := nSignSymDisMode.XValue()

	if curValue != NSignSymDisplayMode.None() {
		t.Errorf("Error: Expected return of object='NumSignSymbolDisplayMode(0).None()'.\n"+
			"Instead, object integer value = '%v'\n",
			curValue.XValueInt())

	}
}

func TestNumSignSymbolDisplayMode_XValue_000200(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).Explicit()

	curValue := nSignSymDisMode.XValue()

	if curValue != NSignSymDisplayMode.Explicit() {
		t.Errorf("Error: Expected return of object='NumSignSymbolDisplayMode(0).Explicit()'.\n"+
			"Instead, object integer value = '%v'\n",
			curValue.XValueInt())
	}
}

func TestNumSignSymbolDisplayMode_XValue_000300(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).Implicit()

	curValue := nSignSymDisMode.XValue()

	if curValue != NSignSymDisplayMode.Implicit() {
		t.Errorf("Error: Expected return of object='NumSignSymbolDisplayMode(0).Implicit()'.\n"+
			"Instead, object integer value = '%v'\n",
			curValue.XValueInt())
	}
}

func TestNumSignSymbolDisplayMode_XValueInt_000100(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).None()

	intValue := nSignSymDisMode.XValueInt()

	if intValue != 0 {
		t.Errorf("Error:"+
			"Expected NumSignSymbolDisplayMode(0).None()\n"+
			"would yield an integer value of '0'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumSignSymbolDisplayMode_XValueInt_000200(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).Explicit()

	intValue := nSignSymDisMode.XValueInt()

	if intValue != 1 {
		t.Errorf("Error:"+
			"Expected NumSignSymbolDisplayMode(0).Explicit()\n"+
			"would yield an integer value of '1'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumSignSymbolDisplayMode_XValueInt_000300(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).Implicit()

	intValue := nSignSymDisMode.XValueInt()

	if intValue != 2 {
		t.Errorf("Error:"+
			"Expected NumSignSymbolDisplayMode(0).Implicit()\n"+
			"would yield an integer value of '2'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumSignSymbolDisplayMode_String_000100(t *testing.T) {

	testStr := "None"

	nSignSymDisMode := NSignSymDisplayMode.None()

	actualStr := nSignSymDisMode.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= '%v'.\n"+
			"Instead, object string value = '%v'\n",
			testStr,
			actualStr)
	}
}

func TestNumSignSymbolDisplayMode_String_000200(t *testing.T) {

	testStr := "Explicit"

	nSignSymDisMode := NSignSymDisplayMode.Explicit()

	actualStr := nSignSymDisMode.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= '%v'.\n"+
			"Instead, object string value = '%v'\n",
			testStr,
			actualStr)
	}
}

func TestNumSignSymbolDisplayMode_String_000300(t *testing.T) {

	testStr := "Implicit"

	nSignSymDisMode := NSignSymDisplayMode.Implicit()

	actualStr := nSignSymDisMode.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= '%v'.\n"+
			"Instead, object string value = '%v'\n",
			testStr,
			actualStr)
	}
}

func TestNumSignSymbolDisplayMode_String_000400(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(999)

	actualStr := nSignSymDisMode.String()

	if !strings.Contains(actualStr, "Error") {
		t.Error("Error:\n" +
			"Expected NumSignSymbolDisplayMode(999) to generate\n" +
			"a string value containing the word 'Error'.\n" +
			"HOWEVER, NO ERROR WAS DETECTED!\n")
	}

}

func TestNumSignSymbolDisplayMode_XIsValid_000100(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).None()

	isValid := nSignSymDisMode.XIsValid()

	if isValid == true {
		t.Error("Error:\n" +
			"Expected NumSignSymbolDisplayMode(0).None()\n" +
			"to generate an XIsValid value of 'false'.\n" +
			"HOWEVER, XIsValid() returned 'true'!\n")
	}

}

func TestNumSignSymbolDisplayMode_XIsValid_000200(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).Explicit()

	isValid := nSignSymDisMode.XIsValid()

	if isValid == false {
		t.Error("Error:\n" +
			"Expected NumSignSymbolDisplayMode(0).Explicit()\n" +
			"to generate an XIsValid value of 'true'.\n" +
			"HOWEVER, XIsValid() returned 'false'!\n")
	}

}

func TestNumSignSymbolDisplayMode_XIsValid_000300(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(0).Implicit()

	isValid := nSignSymDisMode.XIsValid()

	if isValid == false {
		t.Error("Error:\n" +
			"Expected NumSignSymbolDisplayMode(0).Implicit()\n" +
			"to generate an XIsValid value of 'true'.\n" +
			"HOWEVER, XIsValid() returned 'false'!\n")
	}

}

func TestNumSignSymbolDisplayMode_XIsValid_000400(t *testing.T) {

	nSignSymDisMode := NumSignSymbolDisplayMode(999)

	isValid := nSignSymDisMode.XIsValid()

	if isValid == true {
		t.Error("Error:\n" +
			"Expected NumSignSymbolDisplayMode(999)\n" +
			"to generate an XIsValid value of 'false'.\n" +
			"HOWEVER, XIsValid() returned 'true'!\n")
	}

}

package strmech

import (
	"strings"
	"testing"
)

func TestNumSignSymbolPosition_XParseString000100(t *testing.T) {

	testStr := "Before"

	nStrSymPos,
		err := NumSignSymPos.XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumSignSymPos.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).Before() {
		t.Errorf("Error: Expected return of object='nStrSymPos.Before()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000110(t *testing.T) {

	testStr := "Before"

	nStrSymPos,
		err := NumSignSymbolPosition(0).XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumSignSymbolPosition(0).XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).Before() {
		t.Errorf("Error: Expected return of object='nStrSymPos.Before()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000200(t *testing.T) {

	testStr := "After"

	nStrSymPos,
		err := NumSignSymbolPosition(0).XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumSignSymbolPosition(0).XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).After() {
		t.Errorf("Error: Expected return of object='nStrSymPos.After()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000300(t *testing.T) {

	testStr := "BeforeAndAfter"

	nStrSymPos,
		err := NumSignSymbolPosition(0).XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumSignSymbolPosition(0).XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).BeforeAndAfter() {
		t.Errorf("Error: Expected return of object='nStrSymPos.BeforeAndAfter()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000400(t *testing.T) {

	testStr := "None"

	nStrSymPos,
		err := NumSignSymbolPosition(0).XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumSignSymbolPosition(0).XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).None() {
		t.Errorf("Error: Expected return of object='nStrSymPos.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000500(t *testing.T) {

	testStr := "before"

	nStrSymPos,
		err := NumSignSymPos.XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumSignSymPos.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).Before() {
		t.Errorf("Error: Expected return of object='nStrSymPos.Before()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000510(t *testing.T) {

	testStr := "before"

	nStrSymPos,
		err := NumSignSymbolPosition(0).XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumSignSymbolPosition(0).XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).Before() {
		t.Errorf("Error: Expected return of object='nStrSymPos.Before()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000600(t *testing.T) {

	testStr := "after"

	nStrSymPos,
		err := NumSignSymbolPosition(0).XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumSignSymbolPosition(0).XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).After() {
		t.Errorf("Error: Expected return of object='nStrSymPos.After()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000700(t *testing.T) {

	testStr := "beforeandafter"

	nStrSymPos,
		err := NumSignSymbolPosition(0).XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumSignSymbolPosition(0).XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).BeforeAndAfter() {
		t.Errorf("Error: Expected return of object='nStrSymPos.BeforeAndAfter()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000800(t *testing.T) {

	testStr := "none"

	nStrSymPos,
		err := NumSignSymbolPosition(0).XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumSignSymbolPosition(0).XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if nStrSymPos != NumSignSymbolPosition(0).None() {
		t.Errorf("Error: Expected return of object='nStrSymPos.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			nStrSymPos.XValueInt())
	}

}

func TestNumSignSymbolPosition_XParseString000900(t *testing.T) {

	testStr := "xxxx"

	_,
		err := NumSignSymbolPosition(0).XParseString(testStr, false)

	if err == nil {
		t.Error("Expected an error return from NumSignSymbolPosition(0).XParseString()\n" +
			"because testStr= \"xxxx\", a non-existent enum value.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}
}

func TestNumSignSymbolPosition_XParseString001000(t *testing.T) {

	testStr := "XXXX"

	_,
		err := NumSignSymbolPosition(0).XParseString(testStr, true)

	if err == nil {
		t.Error("Expected an error return from NumSignSymbolPosition(0).XParseString()\n" +
			"because testStr= \"xxxx\", a non-existent enum value.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}
}

func TestNumSignSymbolPosition_XValue_000100(t *testing.T) {

	nStrSymPos := NumSignSymbolPosition(0).Before()

	currValue := nStrSymPos.XValue()

	if currValue != NumSignSymPos.Before() {
		t.Errorf("Error: Expected return of object='NumSignSymbolPosition(0).Before()'.\n"+
			"Instead, object integer value = '%v'\n",
			currValue.XValueInt())

	}

}

func TestNumSignSymbolPosition_XValueInt_000100(t *testing.T) {

	nSgnSymPos := NumSignSymbolPosition(0).After()

	intValue := nSgnSymPos.XValueInt()

	if intValue != 2 {
		t.Errorf("Error: Expected return of object integer = '2'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumSignSymbolPosition_String_000100(t *testing.T) {

	testStr := "BeforeAndAfter"

	nSgnSymPos := NumSignSymbolPosition(0).BeforeAndAfter()

	actualStr := nSgnSymPos.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"BeforeAndAfter\".\n"+
			"Instead, object string value = '%v'\n",
			actualStr)
	}
}

func TestNumSignSymbolPosition_String_000200(t *testing.T) {

	testStr := "After"

	nSgnSymPos := NumSignSymbolPosition(0).After()

	actualStr := nSgnSymPos.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"After\".\n"+
			"Instead, object string value = '%v'\n",
			actualStr)
	}
}

func TestNumSignSymbolPosition_String_000300(t *testing.T) {

	testStr := "Before"

	nSgnSymPos := NumSignSymbolPosition(0).Before()

	actualStr := nSgnSymPos.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"Before\".\n"+
			"Instead, object string value = '%v'\n",
			actualStr)
	}
}

func TestNumSignSymbolPosition_String_000400(t *testing.T) {

	testStr := "None"

	nSgnSymPos := NumSignSymbolPosition(0).None()

	actualStr := nSgnSymPos.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"None\".\n"+
			"Instead, object string value = '%v'\n",
			actualStr)
	}
}

func TestNumSignSymbolPosition_String_000500(t *testing.T) {

	nSgnSymPos := NumSignSymbolPosition(999)

	actualStr := nSgnSymPos.String()

	if !strings.Contains(actualStr, "Error") {
		t.Error("Error:\n" +
			"Expected NumSignSymbolPosition(999) to generate\n" +
			"a string value containing the word 'Error'.\n" +
			"HOWEVER, NO ERROR WAS DETECTED!\n")
	}

}

func TestNumSignSymbolPosition_XIsValid_000100(t *testing.T) {

	nSgnSymPos := NumSignSymbolPosition(0).None()

	isValid := nSgnSymPos.XIsValid()

	if isValid == true {
		t.Error("Error: Expected NumSignSymbolPosition(0).None() to yield an INVALID enum.\n" +
			"It did NOT! Return value from nSgnSymPos.XIsValid()=='true'\n")
	}

}

func TestNumSignSymbolPosition_XIsValid_000200(t *testing.T) {

	nSgnSymPos := NumSignSymbolPosition(0).Before()

	isValid := nSgnSymPos.XIsValid()

	if isValid == false {
		t.Error("Error: Expected NumSignSymbolPosition(0).Before() to yield a VALID enum.\n" +
			"It did NOT! Return value from nSgnSymPos.XIsValid()='false'\n")
	}
}

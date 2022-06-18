package strmech

import (
	"strings"
	"testing"
)

func TestNumericSymbolLocation_XParseString_000100(t *testing.T) {

	testStr := "None"

	numSymLocation,
		err := NumericSymbolLocation(0).XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumericSymbolLocation.XParseString(testStr, true).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if numSymLocation != NumSymLocation.None() {
		t.Errorf("Error: Expected return of object='numSymLocation.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSymLocation.XValueInt())
	}

}

func TestNumericSymbolLocation_XParseString_000200(t *testing.T) {

	testStr := "Before"

	numSymLocation,
		err := NumericSymbolLocation(0).XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumericSymbolLocation.XParseString(testStr, true).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if numSymLocation != NumSymLocation.Before() {
		t.Errorf("Error: Expected return of object='numSymLocation.Before()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSymLocation.XValueInt())
	}

}

func TestNumericSymbolLocation_XParseString_000300(t *testing.T) {

	testStr := "Interior"

	numSymLocation,
		err := NumericSymbolLocation(0).XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumericSymbolLocation.XParseString(testStr, true).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if numSymLocation != NumSymLocation.Interior() {
		t.Errorf("Error: Expected return of object='numSymLocation.Interior()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSymLocation.XValueInt())
	}

}

func TestNumericSymbolLocation_XParseString_000400(t *testing.T) {

	testStr := "After"

	numSymLocation,
		err := NumericSymbolLocation(0).XParseString(testStr, true)

	if err != nil {
		t.Errorf("Error returned by NumericSymbolLocation.XParseString(testStr, true).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if numSymLocation != NumSymLocation.After() {
		t.Errorf("Error: Expected return of object='numSymLocation.After()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSymLocation.XValueInt())
	}

}

func TestNumericSymbolLocation_XParseString_000500(t *testing.T) {

	testStr := "none"

	numSymLocation,
		err := NumericSymbolLocation(0).XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumericSymbolLocation.XParseString(testStr, false).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if numSymLocation != NumSymLocation.None() {
		t.Errorf("Error - Case Insensitive Search\n"+
			"Expected return of object='numSymLocation.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSymLocation.XValueInt())
	}

}

func TestNumericSymbolLocation_XParseString_000600(t *testing.T) {

	testStr := "before"

	numSymLocation,
		err := NumericSymbolLocation(0).XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumericSymbolLocation.XParseString(testStr, false).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if numSymLocation != NumericSymbolLocation(0).Before() {
		t.Errorf("Error - Case Insensitive Search\n"+
			"Expected return of object='numSymLocation.Before()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSymLocation.XValueInt())
	}

}

func TestNumericSymbolLocation_XParseString_000700(t *testing.T) {

	testStr := "interior"

	numSymLocation,
		err := NumericSymbolLocation(0).XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumericSymbolLocation.XParseString(testStr, false).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if numSymLocation != NumericSymbolLocation(0).Interior() {
		t.Errorf("Error - Case Insensitive Search\n"+
			"Expected return of object='numSymLocation.Interior()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSymLocation.XValueInt())
	}

}

func TestNumericSymbolLocation_XParseString_000800(t *testing.T) {

	testStr := "after"

	numSymLocation,
		err := NumericSymbolLocation(0).XParseString(testStr, false)

	if err != nil {
		t.Errorf("Error returned by NumericSymbolLocation.XParseString(testStr, false).\n"+
			"testStr='%v'\n"+
			"Error='%v'\n",
			testStr,
			err.Error())

		return
	}

	if numSymLocation != NumericSymbolLocation(0).After() {
		t.Errorf("Error - Case Insensitive Search\n"+
			"Expected return of object='numSymLocation.After()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSymLocation.XValueInt())
	}

}

func TestNumericSymbolLocation_XParseString_000900(t *testing.T) {

	testStr := "X"

	_,
		err := NumericSymbolLocation(0).XParseString(testStr, true)

	if err == nil {
		t.Errorf("Error - Case Sensitive Search\n"+
			"Expected an error return from NumericSymbolLocation(0).XParseString(testStr, true).\n"+
			"because 'testStr' consists of only one character.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n"+
			"testStr='%v'\n",
			testStr)

		return
	}

}

func TestNumericSymbolLocation_XParseString_001000(t *testing.T) {

	testStr := "none"

	_,
		err := NumericSymbolLocation(0).XParseString(testStr, true)

	if err == nil {
		t.Errorf("Error - Case Sensitive Search\n"+
			"Expected an error return from NumericSymbolLocation(0).XParseString(testStr, true).\n"+
			"because 'testStr' consists of all lower case letters.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n"+
			"testStr='%v'\n",
			testStr)

		return
	}

}

func TestNumericSymbolLocation_XParseString_001100(t *testing.T) {

	testStr := "Xray"

	_,
		err := NumericSymbolLocation(0).XParseString(testStr, true)

	if err == nil {
		t.Errorf("Error - Case Sensitive Search\n"+
			"Expected an error return from NumericSymbolLocation(0).XParseString(testStr, true).\n"+
			"because 'testStr' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n"+
			"testStr='%v'\n",
			testStr)

		return
	}

}

func TestNumericSymbolLocation_XParseString_001200(t *testing.T) {

	testStr := "x"

	_,
		err := NumericSymbolLocation(0).XParseString(testStr, false)

	if err == nil {
		t.Errorf("Error - Case Insensitive Search\n"+
			"Expected an error return from NumericSymbolLocation(0).XParseString(testStr, false).\n"+
			"because 'testStr' consists of only one character.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n"+
			"testStr='%v'\n",
			testStr)

		return
	}

}

func TestNumericSymbolLocation_XParseString_001300(t *testing.T) {

	testStr := "X"

	_,
		err := NumericSymbolLocation(0).XParseString(testStr, false)

	if err == nil {
		t.Errorf("Error - Case Insensitive Search\n"+
			"Expected an error return from NumericSymbolLocation(0).XParseString(testStr, false).\n"+
			"because 'testStr' consists of only one character.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n"+
			"testStr='%v'\n",
			testStr)

		return
	}

}

func TestNumericSymbolLocation_XParseString_001400(t *testing.T) {

	testStr := "Xray"

	_,
		err := NumericSymbolLocation(0).XParseString(testStr, false)

	if err == nil {
		t.Errorf("Error - Case Insensitive Search\n"+
			"Expected an error return from NumericSymbolLocation(0).XParseString(testStr, false).\n"+
			"because 'testStr' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n"+
			"testStr='%v'\n",
			testStr)

		return
	}

}

func TestNumericSymbolLocation_XValue_000100(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).None()

	curValue := nSignSymDisMode.XValue()

	if curValue != NumSymLocation.None() {
		t.Errorf("Error:\n"+
			"Expected return of object='NumericSymbolLocation(0).None()'.\n"+
			"Instead, object integer value = '%v'\n",
			curValue.XValueInt())
	}

}

func TestNumericSymbolLocation_XValue_000200(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).Before()

	curValue := nSignSymDisMode.XValue()

	if curValue != NumSymLocation.Before() {
		t.Errorf("Error:\n"+
			"Expected return of object='NumericSymbolLocation(0).Before()'.\n"+
			"Instead, object integer value = '%v'\n",
			curValue.XValueInt())
	}

}

func TestNumericSymbolLocation_XValue_000300(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).Interior()

	curValue := nSignSymDisMode.XValue()

	if curValue != NumSymLocation.Interior() {
		t.Errorf("Error:\n"+
			"Expected return of object='NumericSymbolLocation(0).Interior()'.\n"+
			"Instead, object integer value = '%v'\n",
			curValue.XValueInt())
	}

}

func TestNumericSymbolLocation_XValue_000400(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).After()

	curValue := nSignSymDisMode.XValue()

	if curValue != NumSymLocation.After() {
		t.Errorf("Error:\n"+
			"Expected return of object='NumericSymbolLocation(0).After()'.\n"+
			"Instead, object integer value = '%v'\n",
			curValue.XValueInt())
	}

}

func TestNumericSymbolLocation_XValueInt_000100(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).None()

	intValue := nSignSymDisMode.XValueInt()

	if intValue != 0 {
		t.Errorf("Error:"+
			"Expected NumericSymbolLocation(0).None()\n"+
			"would yield an integer value of '0'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSymbolLocation_XValueInt_000200(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).Before()

	intValue := nSignSymDisMode.XValueInt()

	if intValue != 1 {
		t.Errorf("Error:"+
			"Expected NumericSymbolLocation(0).Before()\n"+
			"would yield an integer value of '1'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSymbolLocation_XValueInt_000300(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).Interior()

	intValue := nSignSymDisMode.XValueInt()

	if intValue != 2 {
		t.Errorf("Error:"+
			"Expected NumericSymbolLocation(0).Interior()\n"+
			"would yield an integer value of '2'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSymbolLocation_XValueInt_000400(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).After()

	intValue := nSignSymDisMode.XValueInt()

	if intValue != 3 {
		t.Errorf("Error:"+
			"Expected NumericSymbolLocation(0).After()\n"+
			"would yield an integer value of '3'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSymbolLocation_XValueInt_000500(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(-99)

	intValue := nSignSymDisMode.XValueInt()

	if intValue != -99 {
		t.Errorf("Error:"+
			"Expected NumericSymbolLocation(-99)\n"+
			"would yield an integer value of '-99'.\n"+
			"Instead, object integer value = '%v'\n",
			intValue)
	}
}

func TestNumericSymbolLocation_String_000100(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).None()

	testStr := nSignSymDisMode.String()

	if testStr != "None" {
		t.Errorf("Error:\n"+
			"Expected nSignSymDisMode.String()=='None'\n"+
			"Instead, nSignSymDisMode.String()== '%v'\n",
			testStr)
	}

}

func TestNumericSymbolLocation_String_000200(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).Before()

	testStr := nSignSymDisMode.String()

	if testStr != "Before" {
		t.Errorf("Error:\n"+
			"Expected nSignSymDisMode.String()=='Before'\n"+
			"Instead, nSignSymDisMode.String()== '%v'\n",
			testStr)
	}

}

func TestNumericSymbolLocation_String_000300(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).Interior()

	testStr := nSignSymDisMode.String()

	if testStr != "Interior" {
		t.Errorf("Error:\n"+
			"Expected nSignSymDisMode.String()=='Interior'\n"+
			"Instead, nSignSymDisMode.String()== '%v'\n",
			testStr)
	}

}

func TestNumericSymbolLocation_String_000400(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).After()

	testStr := nSignSymDisMode.String()

	if testStr != "After" {
		t.Errorf("Error:\n"+
			"Expected nSignSymDisMode.String()=='After'\n"+
			"Instead, nSignSymDisMode.String()== '%v'\n",
			testStr)
	}

}

func TestNumericSymbolLocation_String_000500(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(9999)

	testStr := nSignSymDisMode.String()

	if !strings.Contains(testStr, "Error") {
		t.Error("Error:\n" +
			"Expected nSignSymDisMode.String() to return an Error\n" +
			"because nSignSymDisMode is invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
	}

}

func TestNumericSymbolLocation_String_000600(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(-1)

	testStr := nSignSymDisMode.String()

	if !strings.Contains(testStr, "Error") {
		t.Error("Error:\n" +
			"Expected nSignSymDisMode.String() to return an Error\n" +
			"because nSignSymDisMode == -1 and is therefore invalid.\n" +
			"HOWEVER, NO ERROR WAS RETURNED!\n")
	}
}

func TestNumericSymbolLocation_XIsValid_000100(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).None()

	isValid := nSignSymDisMode.XIsValid()

	if isValid == true {
		t.Error("Error:\n" +
			"Expected isValid == 'false' because\n" +
			"nSignSymDisMode == NumericSymbolLocation(0).None()\n" +
			"HOWEVER, isValid == 'true'!\n")
	}

}

func TestNumericSymbolLocation_XIsValid_000200(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).Before()

	isValid := nSignSymDisMode.XIsValid()

	if isValid == false {
		t.Error("Error:\n" +
			"Expected isValid == 'true' because\n" +
			"nSignSymDisMode == NumericSymbolLocation(0).Before()\n" +
			"HOWEVER, isValid == 'false'!\n")
	}

}

func TestNumericSymbolLocation_XIsValid_000300(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).Interior()

	isValid := nSignSymDisMode.XIsValid()

	if isValid == false {
		t.Error("Error:\n" +
			"Expected isValid == 'true' because\n" +
			"nSignSymDisMode == NumericSymbolLocation(0).Interior()\n" +
			"HOWEVER, isValid == 'false'!\n")
	}

}

func TestNumericSymbolLocation_XIsValid_000400(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(0).After()

	isValid := nSignSymDisMode.XIsValid()

	if isValid == false {
		t.Error("Error:\n" +
			"Expected isValid == 'true' because\n" +
			"nSignSymDisMode == NumericSymbolLocation(0).After()\n" +
			"HOWEVER, isValid == 'false'!\n")
	}

}

func TestNumericSymbolLocation_XIsValid_000500(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(4)

	isValid := nSignSymDisMode.XIsValid()

	if isValid == true {
		t.Error("Error:\n" +
			"Expected isValid == 'false' because\n" +
			"nSignSymDisMode == NumericSymbolLocation(4)\n" +
			"HOWEVER, isValid == 'true'!\n")
	}

}

func TestNumericSymbolLocation_XIsValid_000600(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(-99)

	isValid := nSignSymDisMode.XIsValid()

	if isValid == true {
		t.Error("Error:\n" +
			"Expected isValid == 'false' because\n" +
			"nSignSymDisMode == NumericSymbolLocation(-99)\n" +
			"HOWEVER, isValid == 'true'!\n")
	}

}

func TestNumericSymbolLocation_XIsValid_000700(t *testing.T) {

	nSignSymDisMode := NumericSymbolLocation(99)

	isValid := nSignSymDisMode.XIsValid()

	if isValid == true {
		t.Error("Error:\n" +
			"Expected isValid == 'false' because\n" +
			"nSignSymDisMode == NumericSymbolLocation(99)\n" +
			"HOWEVER, isValid == 'true'!\n")
	}

}

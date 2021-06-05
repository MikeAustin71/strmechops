package strmech

import "testing"

func TestNumericSignValueType_XParseString_000100(t *testing.T) {

	testStr := "None"

	numSignValType,
		err := NumSignVal.XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).None() {
		t.Errorf("Error: Expected return of object='numSignValType.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000200(t *testing.T) {

	testStr := "None"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).None() {
		t.Errorf("Error: Expected return of object='numSignValType.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000300(t *testing.T) {

	testStr := "Negative"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Negative() {
		t.Errorf("Error: Expected return of object='numSignValType.Negative()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000400(t *testing.T) {

	testStr := "Zero"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Zero() {
		t.Errorf("Error: Expected return of object='numSignValType.Zero()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000500(t *testing.T) {

	testStr := "Positive"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		true)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Positive() {
		t.Errorf("Error: Expected return of object='numSignValType.Positive()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000600(t *testing.T) {

	testStr := "none"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		false)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).None() {
		t.Errorf("Error: Expected return of object='numSignValType.None()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000700(t *testing.T) {

	testStr := "negative"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		false)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Negative() {
		t.Errorf("Error: Expected return of object='numSignValType.Negative()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}

}

func TestNumericSignValueType_XParseString_000800(t *testing.T) {

	testStr := "zero"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		false)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Zero() {
		t.Errorf("Error: Expected return of object='numSignValType.Zero()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}
}

func TestNumericSignValueType_XParseString_000900(t *testing.T) {

	testStr := "positive"

	numSignValType,
		err := NumericSignValueType(0).XParseString(
		testStr,
		false)

	if err != nil {
		t.Errorf("Error returned by NumSignVal.XParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if numSignValType != NumericSignValueType(0).Positive() {
		t.Errorf("Error: Expected return of object='numSignValType.Positive()'.\n"+
			"Instead, object integer value = '%v'\n",
			numSignValType.XValueInt())
	}
}

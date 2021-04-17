package strmech

import "testing"

func TestStrOpsTextJustify_ParseString_01(t *testing.T) {

	testStr := "Center"

	txtJustify,
		err := TextJustify(0).XParseString(
		testStr, true)

	if err != nil {
		t.Errorf("Error returned by  TextJustify(0).ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if txtJustify != TxtJustify.Center() {
		t.Errorf("Error: Expected return of object='TxtJustify.Center()'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}
}

func TestStrOpsTextJustify_ParseString_02(t *testing.T) {

	testStr := "Right"

	txtJustify,
		err := TextJustify(0).XParseString(
		testStr, true)

	if err != nil {
		t.Errorf("Error returned by  TextJustify(0).ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if txtJustify != TxtJustify.Right() {
		t.Errorf("Error: Expected return of object='TxtJustify.Right()'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}
}

func TestStrOpsTextJustify_ParseString_03(t *testing.T) {

	testStr := "Left"

	txtJustify,
		err := TextJustify(0).XParseString(
		testStr, true)

	if err != nil {
		t.Errorf("Error returned by  TextJustify(0).ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if txtJustify != TxtJustify.Left() {
		t.Errorf("Error: Expected return of object='TxtJustify.Left()'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}

}

func TestStrOpsTextJustify_ParseString_04(t *testing.T) {

	testStr := "center"

	txtJustify,
		err := TextJustify(0).XParseString(
		testStr, false)

	if err != nil {
		t.Errorf("Error returned by  TextJustify(0).ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if txtJustify != TxtJustify.Center() {
		t.Errorf("Error: Expected return of object='TxtJustify.Center()'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}
}

func TestStrOpsTextJustify_ParseString_05(t *testing.T) {

	testStr := "right"

	txtJustify,
		err := TextJustify(0).XParseString(
		testStr, false)

	if err != nil {
		t.Errorf("Error returned by  TextJustify(0).ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if txtJustify != TxtJustify.Right() {
		t.Errorf("Error: Expected return of object='TxtJustify.Right()'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}
}

func TestStrOpsTextJustify_ParseString_06(t *testing.T) {

	testStr := "left"

	txtJustify,
		err := TextJustify(0).XParseString(
		testStr, false)

	if err != nil {
		t.Errorf("Error returned by  TextJustify(0).ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if txtJustify != TxtJustify.Left() {
		t.Errorf("Error: Expected return of object='TxtJustify.Left()'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}

}

func TestStrOpsTextJustify_ParseString_07(t *testing.T) {

	testStr := "xxxx"

	_,
		err := TextJustify(0).XParseString(
		testStr, false)

	if err == nil {
		t.Error("Expected an error return from TextJustify(0).XParseString()\n" +
			"because testStr= \"xxxx\", a non-existent enum value.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestStrOpsTextJustify_ParseString_08(t *testing.T) {

	testStr := "XXXX"

	_,
		err := TextJustify(0).XParseString(
		testStr, true)

	if err == nil {
		t.Error("Expected an error return from TextJustify(0).XParseString()\n" +
			"because testStr= \"XXXX\", a non-existent enum value.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}

}

func TestStrOpsTextJustify_ParseString_09(t *testing.T) {

	testStr := "Left"

	txtJustify,
		err := TextJustify(0).XParseString(
		testStr, false)

	if err != nil {
		t.Errorf("Error returned by  TextJustify(0).ParseString(testStr, true).\n"+
			"Error='%v'\n", err.Error())
		return
	}

	if txtJustify != TxtJustify.Left() {
		t.Errorf("Error: Expected return of object='TxtJustify.Left()'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}
}

func TestStrOpsTextJustify_XValue_01(t *testing.T) {

	txtJustify := TextJustify(0).Center()

	currValue := txtJustify.XValue()

	if currValue != TxtJustify.Center() {
		t.Errorf("Error: Expected return of object='TxtJustify.Center()'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}
}

func TestStrOpsTextJustify_XValueInt_01(t *testing.T) {

	txtJustify := TextJustify(0).Center()

	currValue := txtJustify.XValueInt()

	if currValue != 3 {
		t.Errorf("Error: Expected return of object integer = '3'.\n"+
			"Instead, object integer value = '%v'\n",
			txtJustify.XValueInt())
	}
}

func TestStrOpsTextJustify_String_01(t *testing.T) {

	testStr := "Center"

	txtJustify := TextJustify(0).Center()

	actualStr := txtJustify.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"Center\".\n"+
			"Instead, object string value = '%v'\n",
			testStr)
	}

}

func TestStrOpsTextJustify_String_02(t *testing.T) {

	testStr := "Left"

	txtJustify := TextJustify(0).Left()

	actualStr := txtJustify.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"Left\".\n"+
			"Instead, object string value = '%v'\n",
			testStr)
	}

}

func TestStrOpsTextJustify_String_03(t *testing.T) {

	testStr := "Right"

	txtJustify := TextJustify(0).Right()

	actualStr := txtJustify.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"Right\".\n"+
			"Instead, object string value = '%v'\n",
			testStr)
	}

}

func TestStrOpsTextJustify_String_04(t *testing.T) {

	testStr := "None"

	txtJustify := TextJustify(0).None()

	actualStr := txtJustify.String()

	if actualStr != testStr {
		t.Errorf("Error: Expected return of object string value= \"None\".\n"+
			"Instead, object string value = '%v'\n",
			testStr)
	}

}

func TestStrOpsTextJustify_XIsValid_01(t *testing.T) {

	txtJustify := TextJustify(0).Center()

	isValid := txtJustify.XIsValid()

	if isValid != true {
		t.Error("Error: Expected TextJustify(0).Center() to yield a valid enum.\n" +
			"It did NOT! Return value from txtJustify.XIsValid()='false'\n")
	}

}

func TestStrOpsTextJustify_XIsValid_02(t *testing.T) {

	txtJustify := TextJustify(0).None()

	isValid := txtJustify.XIsValid()

	if isValid != false {
		t.Error("Error: Expected TextJustify(0).None() to yield an INVALID enum.\n" +
			"It did NOT! Return value from txtJustify.XIsValid()='true'\n")
	}

}

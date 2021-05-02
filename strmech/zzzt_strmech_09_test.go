package strmech

import (
	"strings"
	"testing"
)

func TestStrMech_JustifyTextInStrField_01(t *testing.T) {
	ePrefix := "TestStrMech_JustifyTextInStrField_01() "
	strToJustify := "12345"
	fieldLen := 15
	expectedTotalLen := fieldLen
	exLeftPad := strings.Repeat(" ", 10)
	expectedStr := exLeftPad + strToJustify

	su := StrMech{}

	actualStr, err := su.JustifyTextInStrField(
		strToJustify,
		fieldLen,
		TxtJustify.Right(),
		ePrefix+
			"\nTesting TxtJustify.Right()\n")

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if actualStr != expectedStr {
		t.Errorf("Strings did not match.\n"+
			"Expected formatted string='%v'\n"+
			"Instead, string='%v'\n",
			expectedStr, actualStr)
		return
	}

	l1 := su.StrGetRuneCnt(actualStr)

	if l1 != expectedTotalLen {
		t.Errorf("Expected total actualStr length='%v'.\n"+
			"Instead, actualStr length='%v'.\n",
			expectedTotalLen, l1)
	}
}

func TestStrMech_JustifyTextInStrField_02(t *testing.T) {
	ePrefix := "TestStrMech_JustifyTextInStrField_02() "
	strToJustify := "12345"
	fieldLen := 15
	expectedTotalLen := fieldLen
	exRightPad := strings.Repeat(" ", 10)
	expectedStr := strToJustify + exRightPad

	su := StrMech{}

	actualStr, err := su.JustifyTextInStrField(
		strToJustify,
		fieldLen,
		TxtJustify.Left(),
		ePrefix+
			"\nTesting TxtJustify.Left()\n")

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if actualStr != expectedStr {
		t.Errorf("Strings did not match.\n"+
			"Expected formatted string='%v'\n"+
			"Instead, string='%v'\n",
			expectedStr, actualStr)
		return
	}

	l1 := su.StrGetRuneCnt(actualStr)

	if l1 != expectedTotalLen {
		t.Errorf("Expected total actualStr length='%v'.\n"+
			"Instead, actualStr length='%v'.\n",
			expectedTotalLen, l1)
	}
}

func TestStrMech_JustifyTextInStrField_03(t *testing.T) {
	ePrefix := "TestStrMech_JustifyTextInStrField_03() "
	strToJustify := "12345"
	fieldLen := 15
	expectedTotalLen := fieldLen
	exRightPad := strings.Repeat(" ", 5)
	exLeftPad := strings.Repeat(" ", 5)
	expectedStr := exLeftPad + strToJustify + exRightPad

	su := StrMech{}

	actualStr, err := su.JustifyTextInStrField(
		strToJustify,
		fieldLen,
		TxtJustify.Center(),
		ePrefix+
			"\nTesting TxtJustify.Center()\n")

	if err != nil {
		t.Errorf("%v\n", err.Error())
		return
	}

	if actualStr != expectedStr {
		t.Errorf("Strings did not match.\n"+
			"Expected formatted string='%v'\n"+
			"Instead, string='%v'\n",
			expectedStr, actualStr)
		return
	}

	l1 := su.StrGetRuneCnt(actualStr)

	if l1 != expectedTotalLen {
		t.Errorf("Expected total actualStr length='%v'.\n"+
			"Instead, actualStr length='%v'.\n",
			expectedTotalLen, l1)
	}
}

func TestStrMech_JustifyTextInStrField_04(t *testing.T) {
	ePrefix := "TestStrMech_JustifyTextInStrField_04() "
	strToJustify := ""
	fieldLen := -1

	su := StrMech{}

	_, err := su.JustifyTextInStrField(
		strToJustify,
		fieldLen,
		TxtJustify.Center(),
		ePrefix+
			"\nTesting invalid input parameters.\n")

	if err == nil {
		t.Error("Expected an error return from StrMech{}.JustifyTextInStrField()\n" +
			"because 'strToJustify' == \"\" and 'fieldLen' == -1.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}
}

func TestStrMech_JustifyTextInStrField_05(t *testing.T) {
	ePrefix := "TestStrMech_JustifyTextInStrField_05() "
	strToJustify := "12345"
	fieldLen := 15

	su := StrMech{}

	_, err := su.JustifyTextInStrField(
		strToJustify,
		fieldLen,
		TxtJustify.None(),
		ePrefix+
			"\nTesting invalid TxtJustify value.\n")

	if err == nil {
		t.Error("Expected an error return from StrMech{}.JustifyTextInStrField()\n" +
			"because 'TxtJustify.None()' is an invalid setting for text justification.\n" +
			"However, NO ERROR WAS RETURNED!\n")
		return
	}
}

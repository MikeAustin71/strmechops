package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
	"time"
)

func TestTextFieldSpecDateTimeAtom_empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeAtom_empty_000100()",
		"")

	txtFieldDateTime := TextFieldSpecDateTime{}

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	txtFieldDateTime.dateTime = time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	txtFieldDateTime.dateTimeFormat =
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	txtFieldDateTime.fieldLen = -1

	txtFieldDateTime.textJustification = TxtJustify.Left()

	txtFieldDateTimeAtom := textFieldSpecDateTimeAtom{}

	txtFieldDateTimeAtom.empty(
		nil)

	txtFieldDateTimeAtom.empty(
		&txtFieldDateTime)

	if !txtFieldDateTime.dateTime.IsZero() {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldDateTime.dateTime would be ZERO.\n"+
			"Instead, txtFieldDateTime.dateTime = '%v'\n",
			ePrefix.String(),
			txtFieldDateTime.dateTime.Format("2006-01-02 15:04:05.000000000 -0700 MST"))

		return
	}

	if txtFieldDateTime.fieldLen != 0 {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldDateTime.fieldLen would be ZERO.\n"+
			"Instead, txtFieldDateTime.fieldLen = '%v'\n",
			ePrefix.String(),
			txtFieldDateTime.fieldLen)

		return
	}

	if txtFieldDateTime.dateTimeFormat != "" {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldDateTime.fieldLen would be "+
			"a ZERO length string.\n"+
			"Instead, txtFieldDateTime.fieldLen = '%v'\n",
			ePrefix.String(),
			txtFieldDateTime.fieldLen)

		return
	}

	if txtFieldDateTime.textJustification != TxtJustify.None() {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldDateTime.textJustification\n"+
			"would be equal to TxtJustify.None().\n"+
			"Instead, txtFieldDateTime.textJustification = '%v'\n"+
			"txtFieldDateTime.textJustification integer  = '%v'\n",
			ePrefix.String(),
			txtFieldDateTime.textJustification.String(),
			txtFieldDateTime.textJustification.XValueInt())

		return
	}

	return
}

func TestTextFieldSpecDateTimeAtom_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeAtom_equal_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	txtFieldDateTimeOne.dateTime = time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	txtFieldDateTimeOne.dateTimeFormat =
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	txtFieldDateTimeOne.fieldLen = -1

	txtFieldDateTimeOne.textJustification = TxtJustify.Left()

	txtFieldDateTimeAtom := textFieldSpecDateTimeAtom{}

	areEqual := txtFieldDateTimeAtom.equal(
		nil,
		&txtFieldDateTimeOne)

	if areEqual == true {
		t.Errorf("%v Series #1\n"+
			"Error: Expected txtFieldDateTimeAtom.equal() to\n"+
			"return 'false'.\n"+
			"Instead, txtFieldDateTimeAtom.equal() returned\n"+
			"'true'\n",
			ePrefix.String())

		return

	}

	areEqual = txtFieldDateTimeAtom.equal(
		&txtFieldDateTimeOne,
		nil)

	if areEqual == true {
		t.Errorf("%v Series #2\n"+
			"Error: Expected txtFieldDateTimeAtom.equal() to\n"+
			"return 'false'.\n"+
			"Instead, txtFieldDateTimeAtom.equal() returned\n"+
			"'true'\n",
			ePrefix.String())

		return
	}

	areEqual = txtFieldDateTimeAtom.equal(
		nil,
		nil)

	if areEqual == true {
		t.Errorf("%v Series #3\n"+
			"Error: Expected txtFieldDateTimeAtom.equal() to\n"+
			"return 'false'.\n"+
			"Instead, txtFieldDateTimeAtom.equal() returned\n"+
			"'true'\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeTwo := TextFieldSpecDateTime{}

	timeZoneName = "America/Chicago"

	tzLocPtr, err = time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v Series #4\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	txtFieldDateTimeTwo.dateTime = time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	txtFieldDateTimeTwo.dateTimeFormat =
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	txtFieldDateTimeTwo.fieldLen = -1

	txtFieldDateTimeTwo.textJustification = TxtJustify.Left()

	txtFieldDateTimeAtom2 := textFieldSpecDateTimeAtom{}

	areEqual = txtFieldDateTimeAtom2.equal(
		&txtFieldDateTimeOne,
		&txtFieldDateTimeTwo)

	if areEqual == false {
		t.Errorf("%v Series #5\n"+
			"Error: Expected txtFieldDateTimeAtom.equal() to\n"+
			"return 'true'.\n"+
			"Instead, txtFieldDateTimeAtom.equal() returned\n"+
			"'false'\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeThree := TextFieldSpecDateTime{}

	txtFieldDateTimeThree.dateTime =
		txtFieldDateTimeTwo.dateTime

	txtFieldDateTimeThree.dateTimeFormat =
		txtFieldDateTimeTwo.dateTimeFormat

	txtFieldDateTimeThree.fieldLen =
		txtFieldDateTimeTwo.fieldLen

	txtFieldDateTimeThree.textJustification =
		TxtJustify.None()

	areEqual = txtFieldDateTimeAtom2.equal(
		&txtFieldDateTimeTwo,
		&txtFieldDateTimeThree)

	if areEqual == true {
		t.Errorf("%v Series #6\n"+
			"Error: Expected txtFieldDateTimeAtom.equal() to return\n"+
			"'false' because Text Justifcations are NOT equal.\n"+
			"Instead, txtFieldDateTimeAtom.equal() returned\n"+
			"'true'\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeThree.textJustification =
		txtFieldDateTimeTwo.textJustification

	txtFieldDateTimeThree.fieldLen = 927

	areEqual = txtFieldDateTimeAtom2.equal(
		&txtFieldDateTimeTwo,
		&txtFieldDateTimeThree)

	if areEqual == true {
		t.Errorf("%v Series #7\n"+
			"Error: Expected txtFieldDateTimeAtom.equal() to return\n"+
			"'false' because Field Lengths are NOT equal.\n"+
			"Instead, txtFieldDateTimeAtom.equal() returned\n"+
			"'true'\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeThree.fieldLen =
		txtFieldDateTimeTwo.fieldLen

	txtFieldDateTimeThree.dateTimeFormat = ""

	areEqual = txtFieldDateTimeAtom2.equal(
		&txtFieldDateTimeTwo,
		&txtFieldDateTimeThree)

	if areEqual == true {
		t.Errorf("%v Series #7\n"+
			"Error: Expected txtFieldDateTimeAtom.equal() to return\n"+
			"'false' because Date Time Formats are NOT equal.\n"+
			"Instead, txtFieldDateTimeAtom.equal() returned\n"+
			"'true'\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeThree.dateTimeFormat =
		txtFieldDateTimeTwo.dateTimeFormat

	txtFieldDateTimeThree.dateTime = time.Date(
		2021,
		time.Month(10),
		8,
		22,
		24,
		0,
		0,
		tzLocPtr)

	areEqual = txtFieldDateTimeAtom2.equal(
		&txtFieldDateTimeTwo,
		&txtFieldDateTimeThree)

	if areEqual == true {
		t.Errorf("%v Series #7\n"+
			"Error: Expected txtFieldDateTimeAtom.equal() to return\n"+
			"'false' because Date Time Values are NOT equal.\n"+
			"Instead, txtFieldDateTimeAtom.equal() returned\n"+
			"'true'\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTimeAtom_isValidTextFieldDateTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeAtom_isValidTextFieldDateTime_000100()",
		"")

	txtFieldDateTimeAtom := textFieldSpecDateTimeAtom{}

	isValid,
		err := txtFieldDateTimeAtom.isValidTextFieldDateTime(
		nil,
		ePrefix.XCpy(
			"dateTimeTxtField is 'nil'!"))

	if isValid == true {
		t.Errorf("%v\n"+
			"Error: Expected 'isValid' equals 'false'\n"+
			"because 'dateTimeTxtField' is 'nil'.)\n"+
			"HOWEVER, 'isValid' EQUALS 'true'!\n",
			ePrefix.String())
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected 'err' NOT EQUAL TO 'nil'\n"+
			"because 'dateTimeTxtField' is 'nil'.)\n"+
			"HOWEVER, 'err' IS EQUAL TO 'nil'!\n",
			ePrefix.String())
	}

	return
}

func TestTextFieldSpecDateTimeAtom_isValidTextFieldDateTime_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeAtom_isValidTextFieldDateTime_000200()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}.New()

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	err = txtFieldDateTimeOne.SetDateTimeValue(
		dateTime,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetDateTimeValue()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	err = txtFieldDateTimeOne.SetDateTimeFormat(
		dateTimeFormat,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetDateTimeFormat()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	fieldLen :=
		txtFieldDateTimeOne.GetDateTimeRawStrLen()

	fieldLen += 4

	err = txtFieldDateTimeOne.SetFieldLength(
		fieldLen,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetFieldLength()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	err = txtFieldDateTimeOne.SetTextJustification(
		TxtJustify.Center(),
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetTextJustification()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	txtFieldDateTimeAtom := textFieldSpecDateTimeAtom{}

	var isValid bool

	isValid,
		err = txtFieldDateTimeAtom.isValidTextFieldDateTime(
		&txtFieldDateTimeOne,
		ePrefix.XCpy(
			"txtFieldDateTimeOne"))

	if isValid == false {
		t.Errorf("%v\n"+
			"Error: Expected 'isValid' equals 'true'\n"+
			"because 'txtFieldDateTimeOne' is valid.)\n"+
			"HOWEVER, 'isValid' EQUALS 'false'!\n",
			ePrefix.String())
	}

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeAtom.isValidTextFieldDateTime()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())
	}

	return
}

func TestTextFieldSpecDateTimeAtom_isValidTextFieldDateTime_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeAtom_isValidTextFieldDateTime_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	err = txtFieldDateTimeOne.SetDateTimeValue(
		dateTime,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetDateTimeValue()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	err = txtFieldDateTimeOne.SetDateTimeFormat(
		dateTimeFormat,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetDateTimeFormat()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	fieldLen :=
		txtFieldDateTimeOne.GetDateTimeRawStrLen()

	fieldLen += 4

	err = txtFieldDateTimeOne.SetFieldLength(
		fieldLen,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetFieldLength()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	textJustification := TxtJustify.Center()

	err = txtFieldDateTimeOne.SetTextJustification(
		textJustification,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetTextJustification()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	txtFieldDateTimeAtom := textFieldSpecDateTimeAtom{}

	var isValid bool

	txtFieldDateTimeOne.fieldLen = -375

	isValid,
		err = txtFieldDateTimeAtom.isValidTextFieldDateTime(
		&txtFieldDateTimeOne,
		ePrefix.XCpy(
			"txtFieldDateTimeOne"))

	if isValid == true {
		t.Errorf("%v\n"+
			"Error: Expected 'isValid' equals 'false'\n"+
			"because txtFieldDateTimeOne.fieldLen = -375.)\n"+
			"HOWEVER, 'isValid' EQUALS 'true'!\n",
			ePrefix.String())
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected err != nil\n"+
			"because txtFieldDateTimeOne.fieldLen = -375.)\n"+
			"HOWEVER, err == nil!\n",
			ePrefix.String())
	}

	if isValid == true ||
		err == nil {
		return
	}

	txtFieldDateTimeOne.fieldLen = fieldLen

	txtFieldDateTimeOne.dateTime = time.Time{}

	isValid,
		err = txtFieldDateTimeAtom.isValidTextFieldDateTime(
		&txtFieldDateTimeOne,
		ePrefix.XCpy(
			"txtFieldDateTimeOne"))

	if isValid == true {
		t.Errorf("%v\n"+
			"Error: Expected 'isValid' equals 'false'\n"+
			"because txtFieldDateTimeOne.dateTime = time.Time{}.)\n"+
			"HOWEVER, 'isValid' EQUALS 'true'!\n",
			ePrefix.String())
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected err != nil\n"+
			"because txtFieldDateTimeOne.dateTime = time.Time{}.)\n"+
			"HOWEVER, err == nil!\n",
			ePrefix.String())
	}

	if isValid == true ||
		err == nil {
		return
	}

	txtFieldDateTimeOne.dateTime = dateTime

	txtFieldDateTimeOne.textJustification =
		TxtJustify.None()

	isValid,
		err = txtFieldDateTimeAtom.isValidTextFieldDateTime(
		&txtFieldDateTimeOne,
		ePrefix.XCpy(
			"txtFieldDateTimeOne"))

	if isValid == true {
		t.Errorf("%v\n"+
			"Error: Expected 'isValid' equals 'false'\n"+
			"because txtFieldDateTimeOne.textJustification is invalid.)\n"+
			"HOWEVER, 'isValid' EQUALS 'true'!\n",
			ePrefix.String())
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected err != nil\n"+
			"because txtFieldDateTimeOne.textJustification is invalid.)\n"+
			"HOWEVER, err == nil!\n",
			ePrefix.String())
	}

	if isValid == true ||
		err == nil {
		return
	}

	txtFieldDateTimeOne.textJustification =
		textJustification

	txtFieldDateTimeOne.dateTimeFormat = ""

	isValid,
		err = txtFieldDateTimeAtom.isValidTextFieldDateTime(
		&txtFieldDateTimeOne,
		ePrefix.XCpy(
			"txtFieldDateTimeOne"))

	if isValid == true {
		t.Errorf("%v\n"+
			"Error: Expected 'isValid' equals 'false'\n"+
			"because txtFieldDateTimeOne.dateTimeFormat is invalid.)\n"+
			"HOWEVER, 'isValid' EQUALS 'true'!\n",
			ePrefix.String())
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected err != nil\n"+
			"because txtFieldDateTimeOne.dateTimeFormat is invalid.)\n"+
			"HOWEVER, err == nil!\n",
			ePrefix.String())
	}

	return
}

func TestTextFieldSpecDateTimeMechanics_setTextFieldDateTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeMechanics_setTextFieldDateTime_000100()",
		"")

	txtFieldDateTimeMech := textFieldSpecDateTimeMechanics{}

	txtFieldDateTime := TextFieldSpecDateTime{}

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		14,
		15,
		28,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := 56

	textJustification := TxtJustify.Left()

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		nil,
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from setTextFieldDateTime()\n"+
			"because parameter 'dateTimeTxtField' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	badDateTime := time.Time{}

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		&txtFieldDateTime,
		badDateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from setTextFieldDateTime()\n"+
			"because parameter 'dateTime' has a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		&txtFieldDateTime,
		dateTime,
		-2,
		dateTimeFormat,
		textJustification,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from setTextFieldDateTime()\n"+
			"because parameter 'fieldLen' has a value of minus two (-2).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		&txtFieldDateTime,
		dateTime,
		fieldLen,
		"",
		textJustification,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from setTextFieldDateTime()\n"+
			"because parameter 'dateTimeFormat' is a zero length empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		&txtFieldDateTime,
		dateTime,
		fieldLen,
		"",
		textJustification,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from setTextFieldDateTime()\n"+
			"because parameter 'dateTimeFormat' is a zero length empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		&txtFieldDateTime,
		dateTime,
		fieldLen,
		dateTimeFormat,
		TxtJustify.None(),
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from setTextFieldDateTime()\n"+
			"because parameter 'textJustification' is TxtJustify.None().\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		&txtFieldDateTime,
		dateTime,
		fieldLen,
		dateTimeFormat,
		-9004,
		&ePrefix)

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from setTextFieldDateTime()\n"+
			"because parameter 'textJustification' is '-9004'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		&txtFieldDateTime,
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		&ePrefix)

	if err != nil {
		t.Errorf("%v Test #1\n"+
			"Error: setTextFieldDateTime() returned the\n"+
			"following error when using valid input parameters.\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	txtFieldDateTimeTwo := TextFieldSpecDateTime{}

	txtFieldDateTimeTwo.dateTime = time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	txtFieldDateTimeTwo.dateTimeFormat =
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	txtFieldDateTimeTwo.fieldLen = 60

	txtFieldDateTimeTwo.textJustification = TxtJustify.Center()

	dateTime = time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	dateTimeFormat =
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen = 60

	textJustification = TxtJustify.Center()

	txtFieldDateTimeThree := TextFieldSpecDateTime{}

	err = txtFieldDateTimeMech.setTextFieldDateTime(
		&txtFieldDateTimeThree,
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		&ePrefix)

	if err != nil {
		t.Errorf("%v Test #2\n"+
			"Error: setTextFieldDateTime() returned the\n"+
			"following error when using valid input parameters.\n"+
			"%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	areEqual := textFieldSpecDateTimeAtom{}.ptr().equal(
		&txtFieldDateTimeTwo,
		&txtFieldDateTimeThree)

	if !areEqual {
		t.Errorf("%v Test #3\n"+
			"Error: Expected txtFieldDateTimeTwo and txtFieldDateTimeThree\n"+
			"to have equal member variable data values.\n"+
			"HOWERVER, THE DATA VALUES ARE NOT EQUAL!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTimeMechanics_setTextFieldDateTime_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeMechanics_setTextFieldDateTime_000200()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}.New()

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		14,
		15,
		28,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen := -1

	textJustification := TxtJustify.Left()

	txtFieldDateTimeMech := textFieldSpecDateTimeMechanics{}

	err =
		txtFieldDateTimeMech.setTextFieldDateTime(
			&txtFieldDateTimeOne,
			dateTime,
			fieldLen,
			dateTimeFormat,
			textJustification,
			ePrefix.XCpy(
				"txtFieldDateTimeOne"))

	if err != nil {
		t.Errorf("%v Series-1\n"+
			"Error returned by txtFieldDateTimeMech.setTextFieldDateTime()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	txtFieldDateTimeTwo := TextFieldSpecDateTime{}.NewPtr()

	fieldLen = 3

	err =
		txtFieldDateTimeMech.setTextFieldDateTime(
			txtFieldDateTimeTwo,
			dateTime,
			fieldLen,
			dateTimeFormat,
			textJustification,
			ePrefix.XCpy(
				"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v Series-2\n"+
			"Error returned by txtFieldDateTimeMech.setTextFieldDateTime()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if !txtFieldDateTimeOne.Equal(
		txtFieldDateTimeTwo) {
		t.Errorf("%v\n"+
			"Error: Expected txtFieldDateTimeOne==txtFieldDateTimeTwo.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.String())
	}

	return
}

func TestTextFieldSpecDateTimeNanobot_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeNanobot_copyIn_000100()",
		"")

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		14,
		15,
		28,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var incomingTxtFieldDateTime TextFieldSpecDateTime

	incomingTxtFieldDateTime,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy("incomingTxtFieldDateTime"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	targetTxtFieldDateTime := TextFieldSpecDateTime{}.NewPtr()

	txtFieldDateTimeNanobot := textFieldSpecDateTimeNanobot{}

	err =
		txtFieldDateTimeNanobot.copyIn(
			targetTxtFieldDateTime,
			&incomingTxtFieldDateTime,
			ePrefix.XCtxEmpty())

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeNanobot.copyIn()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if !targetTxtFieldDateTime.Equal(
		&incomingTxtFieldDateTime) {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected 'targetTxtFieldDateTime' to equal 'incomingTxtFieldDateTime'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeNanobotTwo := textFieldSpecDateTimeNanobot{}

	err =
		txtFieldDateTimeNanobotTwo.copyIn(
			targetTxtFieldDateTime,
			nil,
			ePrefix.XCtxEmpty())

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeNanobotTwo.copyIn()\n"+
			"because parameter 'incomingDateTimeTxtField' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtFieldDateTimeNanobotTwo.copyIn(
			nil,
			&incomingTxtFieldDateTime,
			ePrefix.XCtxEmpty())

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeNanobotTwo.copyIn()\n"+
			"because parameter 'targetDateTimeTxtField' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	incomingTxtFieldDateTime.fieldLen = -904

	err =
		txtFieldDateTimeNanobotTwo.copyIn(
			targetTxtFieldDateTime,
			&incomingTxtFieldDateTime,
			ePrefix.XCtxEmpty())

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeNanobotTwo.copyIn()\n"+
			"because parameter 'incomingDateTimeTxtField' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTimeNanobot_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeNanobot_copyOut_000100()",
		"")

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		14,
		15,
		28,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var oldTxtFieldDateTime TextFieldSpecDateTime

	oldTxtFieldDateTime,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCpy("oldTxtFieldDateTime"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	newTxtFieldDateTime := TextFieldSpecDateTime{}.New()

	txtFieldDateTimeNanobot := textFieldSpecDateTimeNanobot{}

	newTxtFieldDateTime,
		err = txtFieldDateTimeNanobot.copyOut(
		&oldTxtFieldDateTime,
		ePrefix.XCpy("oldTxtFieldDateTime"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeNanobot.copyOut()\n"+
			"Error:\n'%v'\n",
			ePrefix.XCpy("newTxtFieldDateTime").String(),
			err.Error())

		return
	}

	if !newTxtFieldDateTime.Equal(
		&oldTxtFieldDateTime) {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected 'newTxtFieldDateTime' to equal 'oldTxtFieldDateTime'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeNanobotTwo := textFieldSpecDateTimeNanobot{}

	_,
		err =
		txtFieldDateTimeNanobotTwo.copyOut(
			nil,
			ePrefix.XCtxEmpty())

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeNanobotTwo.copyOut()\n"+
			"because parameter 'dateTimeTxtField' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	oldTxtFieldDateTime.fieldLen = -9004

	_,
		err =
		txtFieldDateTimeNanobotTwo.copyOut(
			&oldTxtFieldDateTime,
			ePrefix.XCtxEmpty())

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeNanobotTwo.copyOut()\n"+
			"because parameter 'oldTxtFieldDateTime' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTimeNanobot_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeNanobot_getFormattedText_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}.New()

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	fieldLen := -1

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	textJustification := TxtJustify.Center()

	err = txtFieldDateTimeOne.SetDateTimeFieldSpec(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetDateTimeFieldSpec()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	expectedDateTimeStr :=
		txtFieldDateTimeOne.GetDateTimeRawString()

	txtFieldDateTimeNanobotOne := textFieldSpecDateTimeNanobot{}

	var actualFormattedDateTimeStr string

	actualFormattedDateTimeStr,
		err =
		txtFieldDateTimeNanobotOne.getFormattedText(
			&txtFieldDateTimeOne,
			ePrefix.XCpy(
				"txtFieldDateTimeOne"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeNanobotOne.getFormattedText()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if expectedDateTimeStr != actualFormattedDateTimeStr {
		t.Errorf("%v - Error\n"+
			"Expected expectedDateTimeStr == actualFormattedDateTimeStr.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedDateTimeStr        = '%v'\n"+
			"actualFormattedDateTimeStr = '%v'\n",
			ePrefix.String(),
			expectedDateTimeStr,
			actualFormattedDateTimeStr)
	}

	return
}

func TestTextFieldSpecDateTimeNanobot_getFormattedText_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeNanobot_getFormattedText_000200()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}.New()

	timeZoneName := "America/Chicago"

	tzLocPtr, err := time.LoadLocation(timeZoneName)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by time.LoadLocation(timeZoneName)\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	dateTime := time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	fieldLen := -1

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	textJustification := TxtJustify.Center()

	err = txtFieldDateTimeOne.SetDateTimeFieldSpec(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetDateTimeFieldSpec()\n"+
			"Error=\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	txtFieldDateTimeOne.fieldLen = 1000001

	txtFieldDateTimeNanobotOne := textFieldSpecDateTimeNanobot{}

	_,
		err =
		txtFieldDateTimeNanobotOne.getFormattedText(
			&txtFieldDateTimeOne,
			ePrefix.XCpy(
				"txtFieldDateTimeOne"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeNanobotOne.getFormattedText()\n"+
			"because 'txtFieldDateTimeOne.fieldLen' == 1,000,001.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTimeNanobot_getFormattedText_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeNanobot_getFormattedText_000200()",
		"")

	txtFieldDateTimeNanobotOne := textFieldSpecDateTimeNanobot{}

	_,
		err :=
		txtFieldDateTimeNanobotOne.getFormattedText(
			nil,
			ePrefix.XCpy(
				"dateTimeTxtField==nil"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeNanobotOne.getFormattedText()\n"+
			"because 'dateTimeTxtField' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

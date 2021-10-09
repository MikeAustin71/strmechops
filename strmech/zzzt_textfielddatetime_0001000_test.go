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

func TestTextFieldSpecDateTime_Empty_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_Empty_000100()",
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

	txtFieldDateTime.Empty()

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

	txtFieldDateTime2 := TextFieldSpecDateTime{}

	txtFieldDateTime2.Empty()

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

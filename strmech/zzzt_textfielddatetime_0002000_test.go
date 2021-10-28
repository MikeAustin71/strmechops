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

func TestTextFieldSpecDateTimeAtom_isValidTextFieldDateTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTimeAtom_isValidTextFieldDateTime_000100()",
		"")

	txtFieldDateTimeAtom := textFieldSpecDateTimeAtom{}

	isValid,
		err := txtFieldDateTimeAtom.isValidTextFieldDateTime(
		nil,
		ePrefix.XCtx(
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
		ePrefix.XCtx(
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
		ePrefix.XCtx(
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
		ePrefix.XCtx(
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
		ePrefix.XCtx(
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
		ePrefix.XCtx(
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

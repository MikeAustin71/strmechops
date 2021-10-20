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

func TestTextFieldSpecDateTime_CopyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_CopyIn_000100()",
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
		ePrefix.XCtx("incomingTxtFieldDateTime"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	targetTxtFieldDateTime := TextFieldSpecDateTime{}.NewPtr()

	err = targetTxtFieldDateTime.CopyIn(
		&incomingTxtFieldDateTime,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by targetTxtFieldDateTime.CopyIn()\n"+
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

	targetTxtFieldDateTime2 := TextFieldSpecDateTime{}.New()

	err =
		targetTxtFieldDateTime2.CopyIn(
			nil,
			ePrefix.XCtx("incomingTxtFieldDateTime==nil"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from targetTxtFieldDateTime2.CopyIn()\n"+
			"because parameter 'incomingTxtFieldDateTime' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	incomingTxtFieldDateTime.fieldLen = -9009

	err =
		targetTxtFieldDateTime2.CopyIn(
			&incomingTxtFieldDateTime,
			ePrefix.XCtx("incomingTxtFieldDateTime==nil"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from targetTxtFieldDateTime2.CopyIn()\n"+
			"because parameter 'incomingTxtFieldDateTime' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	incomingTxtFieldDateTime.fieldLen =
		len(dateTimeFormat) + 8

	timeZoneName = "America/Los_Angeles"

	tzLocPtr, err = time.LoadLocation(timeZoneName)

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

	dateTime = time.Date(
		2021,
		time.Month(10),
		20,
		13,
		34,
		0,
		0,
		tzLocPtr)

	dateTimeFormat =
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen = len(dateTimeFormat) + 8

	textJustification = TxtJustify.Left()

	var targetTxtFieldDateTime3 *TextFieldSpecDateTime

	targetTxtFieldDateTime3,
		err = TextFieldSpecDateTime{}.NewPtrDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("targetTxtFieldDateTime3"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewPtrDateTimeField()\n"+
			"timeZoneName='%v'\n"+
			"Error='%v'\n",
			ePrefix.String(),
			timeZoneName,
			err.Error())

		return

	}

	err =
		targetTxtFieldDateTime3.CopyIn(
			targetTxtFieldDateTime,
			ePrefix.XCtx("targetTxtFieldDateTime"))

	if !targetTxtFieldDateTime3.Equal(
		targetTxtFieldDateTime) {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected 'targetTxtFieldDateTime3' to equal 'targetTxtFieldDateTime'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

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

func TestTextFieldSpecDateTime_Equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_Equal_000100()",
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

	areEqual := txtFieldDateTimeOne.Equal(
		&txtFieldDateTimeTwo)

	if areEqual == false {
		t.Errorf("%v Series #2\n"+
			"Error: Expected txtFieldDateTimeOne.Equal() to\n"+
			"return 'true' because 'txtFieldDateTimeOne' and\n"+
			"'txtFieldDateTimeTwo' are equal to each other.\n"+
			"Instead, txtFieldDateTimeOne.Equal() returned\n"+
			"'false'\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeThree := TextFieldSpecDateTime{}

	areEqual = txtFieldDateTimeThree.Equal(
		&txtFieldDateTimeTwo)

	if areEqual == true {
		t.Errorf("%v Series #3\n"+
			"Error: Expected txtFieldDateTimeThree.Equal() to\n"+
			"return 'false' because 'txtFieldDateTimeThree' and\n"+
			"'txtFieldDateTimeTwo' are NOT equal to each other.\n"+
			"Instead, txtFieldDateTimeThree.Equal() returned\n"+
			"'true'\n",
			ePrefix.String())

		return
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
		ePrefix.XCtx("incomingTxtFieldDateTime"))

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
		ePrefix.XCtx("oldTxtFieldDateTime"))

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
		ePrefix.XCtx("oldTxtFieldDateTime"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeNanobot.copyOut()\n"+
			"Error:\n'%v'\n",
			ePrefix.XCtx("newTxtFieldDateTime").String(),
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

func TestTextFieldSpecDateTime_NewPtrDateTimeField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_NewPtrDateTimeField_000100()",
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

	txtFieldDateTimeOne.fieldLen =
		len(txtFieldDateTimeOne.dateTimeFormat) + 8

	txtFieldDateTimeOne.textJustification = TxtJustify.Center()

	var txtFieldDateTimeTwo *TextFieldSpecDateTime

	dateTime := time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewPtrDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewPtrDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if !txtFieldDateTimeTwo.Equal(
		&txtFieldDateTimeOne) {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected 'txtFieldDateTimeOne' to equal to 'txtFieldDateTimeTwo'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTime_NewDateTimeField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_NewDateTimeField_000100()",
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

	txtFieldDateTimeOne.fieldLen =
		len(txtFieldDateTimeOne.dateTimeFormat) + 8

	txtFieldDateTimeOne.textJustification = TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	dateTime := time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if !txtFieldDateTimeTwo.Equal(
		&txtFieldDateTimeOne) {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected 'txtFieldDateTimeOne' to equal 'txtFieldDateTimeTwo'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTime_SetDateTimeField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_SetDateTimeField_000100()",
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

	txtFieldDateTimeOne.fieldLen =
		len(txtFieldDateTimeOne.dateTimeFormat) + 8

	txtFieldDateTimeOne.textJustification = TxtJustify.Center()

	txtFieldDateTimeTwo := TextFieldSpecDateTime{}.NewPtr()

	dateTime := time.Date(
		2021,
		time.Month(10),
		6,
		23,
		55,
		0,
		0,
		tzLocPtr)

	dateTimeFormat :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewPtrDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewPtrDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	err =
		txtFieldDateTimeTwo.SetDateTimeField(
			dateTime,
			fieldLen,
			dateTimeFormat,
			textJustification,
			ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.SetDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if !txtFieldDateTimeTwo.Equal(
		&txtFieldDateTimeOne) {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected 'txtFieldDateTimeOne' to equal to 'txtFieldDateTimeTwo'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

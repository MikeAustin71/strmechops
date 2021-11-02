package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"io"
	"strings"
	"testing"
	"time"
)

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
			"Error='%v'\n",
			ePrefix.String(),
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

func TestTextFieldSpecDateTime_CopyIn_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_CopyIn_000200()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	err :=
		txtFieldDateTimeOne.CopyIn(
			nil,
			ePrefix.XCtx(
				"incomingDateTimeTxtField='nil'"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeOne.CopyIn()\n"+
			"because 'incomingDateTimeTxtField' is nil.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

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

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.
		NewDateTimeField(
			dateTime,
			fieldLen,
			dateTimeFormat,
			textJustification,
			ePrefix.XCtx(
				"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error='%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	dateTimeRawString :=
		txtFieldDateTimeTwo.GetDateTimeRawString()

	fieldLen = len(dateTimeRawString) + 4

	err =
		txtFieldDateTimeTwo.SetFieldLength(
			fieldLen,
			ePrefix.XCtx(
				"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.SetFieldLength()\n"+
			"Error='%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	var dateTimeTwoFmtStr, dateTimeThreeFmtStr string

	dateTimeTwoFmtStr,
		err =
		txtFieldDateTimeTwo.GetFormattedText(
			ePrefix.XCtx(
				"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.GetFormattedText()\n"+
			"Error='%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	txtFieldDateTimeThree := TextFieldSpecDateTime{}

	err =
		txtFieldDateTimeThree.CopyIn(
			&txtFieldDateTimeTwo,
			ePrefix.XCtx(
				"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.CopyIn()\n"+
			"Error='%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if !txtFieldDateTimeTwo.Equal(&txtFieldDateTimeThree) {
		t.Errorf("%v - Error\n"+
			"Expected txtFieldDateTimeTwo==txtFieldDateTimeThree.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	dateTimeThreeFmtStr,
		err =
		txtFieldDateTimeThree.GetFormattedText(
			ePrefix.XCtx(
				"txtFieldDateTimeThree"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeThree.GetFormattedText()\n"+
			"Error='%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if dateTimeTwoFmtStr != dateTimeThreeFmtStr {
		t.Errorf("%v - Error\n"+
			"Expected dateTimeTwoFmtStr==dateTimeThreeFmtStr.\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecDateTime_CopyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_CopyOut_000100()",
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

	var txtFieldDateTimeOne TextFieldSpecDateTime

	txtFieldDateTimeOne,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeOne"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = txtFieldDateTimeOne.CopyOut(
		ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.CopyOut()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if !txtFieldDateTimeTwo.Equal(
		&txtFieldDateTimeOne) {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected 'txtFieldDateTimeTwo' to equal 'txtFieldDateTimeOne'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeTwo.fieldLen = -9009

	_,
		err =
		txtFieldDateTimeTwo.CopyOut(
			ePrefix.XCtx("txtFieldDateTimeTwo is invalid!"))

	if err == nil {
		t.Errorf("%v - Error\n"+
			"Expected an error return from txtFieldDateTimeTwo.CopyOut()\n"+
			"because parameter 'txtFieldDateTimeTwo.fieldLen' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	txtFieldDateTimeTwo.fieldLen =
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

	var txtFieldDateTimeThree TextFieldSpecDateTime

	txtFieldDateTimeThree,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeThree"))

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

	txtFieldDateTimeThree,
		err =
		txtFieldDateTimeOne.CopyOut(
			ePrefix.XCtx("txtFieldDateTimeOne->txtFieldDateTimeThree"))

	if !txtFieldDateTimeOne.Equal(
		&txtFieldDateTimeThree) {

		t.Errorf("%v\n"+
			"Error:\n"+
			"Expected 'txtFieldDateTimeThree' to equal 'txtFieldDateTimeOne'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTime_CopyOutITextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_CopyOutITextField_000100()",
		"")

	txtFieldDateTime := TextFieldSpecDateTime{}

	_,
		err := txtFieldDateTime.CopyOutITextField(
		ePrefix.XCtx(
			"txtFieldDateTime"))

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtFieldLabel.CopyOutITextField()\n"+
			"because 'txtFieldDateTime' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETUNRED!\n",
			ePrefix.String())

		return
	}

	timeZoneName := "America/Chicago"

	var tzLocPtr *time.Location

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

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	err = txtFieldDateTimeTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var txtFieldSpec ITextFieldSpecification

	txtFieldSpec,
		err = txtFieldDateTimeTwo.CopyOutITextField(
		ePrefix.XCtx(
			"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var ok bool
	var txtFieldDateTimeThree *TextFieldSpecDateTime

	txtFieldDateTimeThree, ok =
		txtFieldSpec.(*TextFieldSpecDateTime)

	if !ok {
		t.Errorf("%v\n"+
			"Error: Could not convert 'txtFieldSpec' to "+
			"'*TextFieldSpecDateTime'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	err = txtFieldDateTimeThree.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldDateTimeThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldDateTimeThree.Equal(
		&txtFieldDateTimeTwo) {
		t.Errorf("%v\n"+
			"Error: Expected 'txtFieldDateTimeTwo'==txtFieldDateTimeThree'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	return
}

func TestTextFieldSpecDateTime_CopyOutPtr_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_CopyOutITextField_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	_,
		err := txtFieldDateTimeOne.CopyOutPtr(
		ePrefix)

	if err == nil {
		t.Errorf("%v\n"+
			"Error: Expected error return from txtFieldLabel.CopyOutITextField()\n"+
			"because 'txtFieldDateTimeOne' is empty.\n"+
			"HOWEVER, NO ERROR WAS RETUNRED!\n",
			ePrefix.String())

		return
	}

	timeZoneName := "America/Chicago"

	var tzLocPtr *time.Location

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

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Right()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	err = txtFieldDateTimeTwo.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	expectedDateTimeStr :=
		txtFieldDateTimeTwo.GetDateTimeRawString()

	expectedDateTimeStr =
		strings.Repeat(" ", 8) +
			expectedDateTimeStr

	var txtFieldDateTimeThree *TextFieldSpecDateTime

	txtFieldDateTimeThree,
		err = txtFieldDateTimeTwo.CopyOutPtr(
		ePrefix.XCtx(
			"txtFieldDateTimeThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	err = txtFieldDateTimeThree.IsValidInstanceError(
		ePrefix.XCtx(
			"txtFieldDateTimeThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if !txtFieldDateTimeThree.Equal(
		&txtFieldDateTimeTwo) {
		t.Errorf("%v\n"+
			"Error: Expected 'txtFieldDateTimeTwo'=='txtFieldDateTimeThree'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	var actualFmtStr string

	actualFmtStr,
		err =
		txtFieldDateTimeThree.GetFormattedText(
			ePrefix.XCtx(
				"txtFieldDateTimeThree"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	if actualFmtStr != expectedDateTimeStr {
		t.Errorf("%v - ERROR\n"+
			"Error: Expected 'actualFmtStr'=='expectedDateTimeStr'\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"expectedDateTimeStr= '%v'\n"+
			"actualFmtStr       = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			expectedDateTimeStr,
			actualFmtStr)

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

func TestTextFieldSpecDateTime_EqualITextField_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_EqualITextField_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	var txtFieldDateTimeTwo *TextFieldSpecDateTime

	areEqual :=
		txtFieldDateTimeOne.EqualITextField(txtFieldDateTimeTwo)

	if areEqual == true {
		t.Errorf("%v - ERROR\n"+
			"areEqual = txtFieldDateTimeOne.EqualITextField(txtFieldDateTimeTwo)\n"+
			"Expected areEqual == false\n"+
			"because 'txtFieldDateTimeTwo' is a nil pointer.\n"+
			"HOWEVER, areEqual == true!\n",
			ePrefix.String())

		return
	}

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

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeThree TextFieldSpecDateTime

	txtFieldDateTimeThree,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeThree"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	var txtFieldLabel *TextFieldSpecLabel

	txtFieldLabel,
		err = TextFieldSpecLabel{}.NewPtrTextLabel(
		"Hello World",
		24,
		TxtJustify.Left(),
		ePrefix.XCtx(
			"txtFieldLabel"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecLabel{}.NewPtrTextLabel()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	areEqual = txtFieldDateTimeThree.EqualITextField(
		txtFieldLabel)

	if areEqual == true {
		t.Errorf("%v - ERROR\n"+
			"areEqual = txtFieldDateTimeThree.EqualITextField(txtFieldLabel)\n"+
			"Expected areEqual == false\n"+
			"because 'txtFieldLabel' is of type 'TextFieldSpecLabel'.\n"+
			"HOWEVER, areEqual == true!\n",
			ePrefix.String())

		return
	}

	var txtITextFieldSpecDateTime ITextFieldSpecification

	txtITextFieldSpecDateTime,
		err = txtFieldDateTimeThree.CopyOutITextField(
		ePrefix.XCtx(
			"txtITextFieldSpecDateTime"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeThree.CopyOutITextField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	areEqual = txtFieldDateTimeThree.EqualITextField(
		txtITextFieldSpecDateTime)

	if areEqual == false {
		t.Errorf("%v - ERROR\n"+
			"areEqual = txtFieldDateTimeThree.EqualITextField(txtITextFieldSpecDateTime)\n"+
			"Expected areEqual == 'true'\n"+
			"HOWEVER, areEqual == 'false'!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTime_GetDateTime_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_GetDateTime_000100()",
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
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	actualDateTime :=
		txtFieldDateTimeOne.GetDateTime()

	if !actualDateTime.IsZero() {
		t.Errorf("%v - ERROR\n"+
			"Expected 'actualDateTime' == ZERO.\n"+
			"HOWEVER, 'actualDateTime' IS NOT ZERO!\n"+
			"actualDateTime = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualDateTime.Format(dateTimeFormat))

		return
	}

	txtFieldDateTimeOne,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeOne"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	actualDateTime =
		txtFieldDateTimeOne.GetDateTime()

	if dateTime != actualDateTime {
		t.Errorf("%v - ERROR\n"+
			"Expected dateTime == actualDateTime\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"dateTime       = '%v'\n"+
			"actualDateTime = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			dateTime.Format(dateTimeFormat),
			actualDateTime.Format(dateTimeFormat))

		return
	}

	return
}

func TestTextFieldSpecDateTime_GetDateTimeFormat_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_GetDateTime_000100()",
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
		"2006-01-02 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	actualDateTimeFormat :=
		txtFieldDateTimeOne.GetDateTimeFormat()

	if actualDateTimeFormat != "" {
		t.Errorf("%v - ERROR\n"+
			"Expected 'actualDateTimeFormat' == empty string.\n"+
			"HOWEVER, 'actualDateTimeFormat' IS NOT AN EMPTY STRING!\n"+
			"actualDateTimeFormat = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualDateTimeFormat)

		return
	}

	txtFieldDateTimeOne,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeOne"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	actualDateTimeFormat =
		txtFieldDateTimeOne.GetDateTimeFormat()

	if dateTimeFormat != actualDateTimeFormat {
		t.Errorf("%v - ERROR\n"+
			"Expected dateTimeFormat == actualDateTimeFormat\n"+
			"HOWEVER, THEY ARE NOT EQUAL!\n"+
			"dateTimeFormat       = '%v'\n"+
			"actualDateTimeFormat = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			dateTimeFormat,
			actualDateTimeFormat)

		return
	}

	return
}

func TestTextFieldSpecDateTime_GetDateTimeRawStrLen_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_GetDateTime_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	expectedDateTimeRawStringLen :=
		len(time.Time{}.Format(""))

	actualDateTimeRawStringLen :=
		txtFieldDateTimeOne.GetDateTimeRawStrLen()

	if expectedDateTimeRawStringLen !=
		actualDateTimeRawStringLen {

		t.Errorf("%v - ERROR\n"+
			"Test # 1"+
			"Expected Date Time Raw String Length is NOT\n"+
			"EQUAL to the Actual Date Time Raw String Length!\n"+
			"Expected Date Time Raw String Length = '%v'\n"+
			"Actual Date Time Raw String Length   = '%v'\n",
			ePrefix.String(),
			expectedDateTimeRawStringLen,
			actualDateTimeRawStringLen)

		return
	}

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

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	rawDateTimeStr :=
		txtFieldDateTimeTwo.GetDateTimeRawString()

	expectedDateTimeRawStringLen =
		len(rawDateTimeStr)

	actualDateTimeRawStringLen =
		txtFieldDateTimeTwo.GetDateTimeRawStrLen()

	if expectedDateTimeRawStringLen !=
		actualDateTimeRawStringLen {

		t.Errorf("%v - ERROR\n"+
			"Test # 2"+
			"Expected Date Time Raw String Length is NOT\n"+
			"EQUAL to the Actual Date Time Raw String Length!\n"+
			"Expected Date Time Raw String Length = '%v'\n"+
			"Actual Date Time Raw String Length   = '%v'\n",
			ePrefix.String(),
			expectedDateTimeRawStringLen,
			actualDateTimeRawStringLen)

		return
	}

	return
}

func TestTextFieldSpecDateTime_GetFieldLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_GetFieldLength_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	expectedFieldLength := 0

	actualFieldLength :=
		txtFieldDateTimeOne.GetFieldLength()

	if expectedFieldLength !=
		actualFieldLength {
		t.Errorf("%v - ERROR\n"+
			"Test #1"+
			"Expected Field Length is NOT EQUAL to Actual Field Length!\n"+
			"Expected Field Length = '%v'\n"+
			"Actual Field Length   = '%v'\n",
			ePrefix.String(),
			expectedFieldLength,
			actualFieldLength)

		return
	}

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

	expectedFieldLength = len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		expectedFieldLength,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	actualFieldLength =
		txtFieldDateTimeTwo.GetFieldLength()

	if expectedFieldLength !=
		actualFieldLength {
		t.Errorf("%v - ERROR\n"+
			"Test #2"+
			"Expected Field Length is NOT EQUAL to Actual Field Length!\n"+
			"Expected Field Length = '%v'\n"+
			"Actual Field Length   = '%v'\n",
			ePrefix.String(),
			expectedFieldLength,
			actualFieldLength)

		return
	}

	return
}

func TestTextFieldSpecDateTime_GetFormattedStrLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_GetFormattedStrLength_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	expectedFormattedStringLength := -1

	actualFormattedStringLength :=
		txtFieldDateTimeOne.GetFormattedStrLength()

	if expectedFormattedStringLength !=
		actualFormattedStringLength {
		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Formatted String Length is NOT EQUAL to\n"+
			"Actual Formatted String Length!\n"+
			"Expected Formatted String Length = '%v'\n"+
			"Actual Formatted String Length   = '%v'\n",
			ePrefix.String(),
			expectedFormattedStringLength,
			actualFormattedStringLength)

		return
	}

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

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	var formattedString string

	formattedString,
		err = txtFieldDateTimeTwo.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	expectedFormattedStringLength =
		len(formattedString)

	actualFormattedStringLength =
		txtFieldDateTimeTwo.GetFormattedStrLength()

	if expectedFormattedStringLength !=
		actualFormattedStringLength {
		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected Formatted String Length is NOT EQUAL to\n"+
			"Actual Formatted String Length!\n"+
			"Expected Formatted String Length = '%v'\n"+
			"Actual Formatted String Length   = '%v'\n",
			ePrefix.String(),
			expectedFormattedStringLength,
			actualFormattedStringLength)

		return
	}

	return
}

func TestTextFieldSpecDateTime_GetFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_GetFormattedText_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	_,
		err := txtFieldDateTimeOne.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldDateTimeOne"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from "+
			"txtFieldDateTimeOne.GetFormattedText()\n"+
			"because 'txtFieldDateTimeOne' is EMPTY.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	timeZoneName := "America/Chicago"

	var tzLocPtr *time.Location

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

	fieldLen := len(dateTimeFormat) + 4

	textJustification := TxtJustify.Left()

	expectedFormattedText :=
		dateTime.Format(dateTimeFormat) +
			strings.Repeat(" ", 4)

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	var actualFormattedText string

	actualFormattedText,
		err = txtFieldDateTimeTwo.GetFormattedText(
		ePrefix.XCtx(
			"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if expectedFormattedText !=
		actualFormattedText {
		t.Errorf("%v - ERROR\n"+
			"Expected Formatted Text is NOT EQUAL to\n"+
			"Actual Formatted Text!\n"+
			"Expected Formatted Text = '%v'\n"+
			"Actual Formatted Text   = '%v'\n",
			ePrefix.String(),
			expectedFormattedText,
			actualFormattedText)

		return
	}

	return
}

func TestTextFieldSpecDateTime_GetTextJustification_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_GetTextJustification_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	expectedTextJustification := TxtJustify.None()

	actualTextJustification :=
		txtFieldDateTimeOne.GetTextJustification()

	if expectedTextJustification !=
		actualTextJustification {

		t.Errorf("%v - ERROR\n"+
			"Test #1\n"+
			"Expected Text Justification is NOT EQUAL to\n"+
			"Actual Text Justification!\n"+
			"Expected Text Justification = '%v'\n"+
			"Actual Text Justification   = '%v'\n",
			ePrefix.String(),
			expectedTextJustification,
			actualTextJustification)

		return
	}

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

	fieldLen := len(dateTimeFormat) + 4

	expectedTextJustification = TxtJustify.Left()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		expectedTextJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	actualTextJustification =
		txtFieldDateTimeTwo.GetTextJustification()

	if expectedTextJustification !=
		actualTextJustification {

		t.Errorf("%v - ERROR\n"+
			"Test #2\n"+
			"Expected Text Justification is NOT EQUAL to\n"+
			"Actual Text Justification!\n"+
			"Expected Text Justification = '%v'\n"+
			"Actual Text Justification   = '%v'\n",
			ePrefix.String(),
			expectedTextJustification,
			actualTextJustification)

		return
	}

	return
}

func TestTextFieldSpecDateTime_IsValidInstance_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_IsValidInstance_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	actualIsValid := txtFieldDateTimeOne.IsValidInstance()

	if actualIsValid {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' == 'false' because\n"+
			"'txtFieldDateTimeOne' is an empty object.\n"+
			"HOWEVER, 'isValid' == 'true'!\n",
			ePrefix.String())

		return
	}

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

	fieldLen := len(dateTimeFormat) + 4

	textJustification := TxtJustify.Left()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	actualIsValid = txtFieldDateTimeTwo.IsValidInstance()

	if !actualIsValid {
		t.Errorf("%v - ERROR\n"+
			"Expected 'isValid' == 'true' because\n"+
			"'txtFieldDateTimeTwo' is valid.\n"+
			"HOWEVER, 'isValid' == 'false'!\n",
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

func TestTextFieldSpecDateTime_Read_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_Read_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	pxByte := make([]byte, 300)

	_,
		err :=
		txtFieldDateTimeOne.Read(pxByte)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldDateTimeOne.Read(pxByte)\n"+
			"because 'txtFieldDateTimeOne' is an empty object and invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timeZoneName := "America/Chicago"

	var tzLocPtr *time.Location

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

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	err = txtFieldDateTimeTwo.IsValidInstanceError(
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.IsValidInstanceError()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	lenExpectedDateTimeText :=
		txtFieldDateTimeTwo.GetFormattedStrLength()

	var expectedDateTimeText string

	expectedDateTimeText,
		err =
		txtFieldDateTimeTwo.GetFormattedText(
			ePrefix.XCtx(
				"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	p := make([]byte, 500)

	var n, readBytesCnt int
	var actualStr string

	for {

		n,
			err = txtFieldDateTimeTwo.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtFieldDateTimeTwo.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if txtFieldDateTimeTwo.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtFieldDateTimeTwo.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedDateTimeText {
		t.Errorf("%v\n"+
			"Byte Length Error: txtFieldLabelOne.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedDateTimeText,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedDateTimeText),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	return
}

func TestTextFieldSpecDateTime_ReadInitialize_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_ReadInitialize_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	txtFieldDateTimeOne.ReaderInitialize()

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

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	err = txtFieldDateTimeTwo.IsValidInstanceError(
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.IsValidInstanceError()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	lenExpectedDateTimeText :=
		txtFieldDateTimeTwo.GetFormattedStrLength()

	var expectedDateTimeText string

	expectedDateTimeText,
		err =
		txtFieldDateTimeTwo.GetFormattedText(
			ePrefix.XCtx(
				"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.GetFormattedText()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	p := make([]byte, 5)

	var n int

	n,
		err = txtFieldDateTimeTwo.Read(p)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.Read(p)\n"+
			"Error:\n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if n != 5 {
		t.Errorf("%v\n"+
			"Error: fillerTxtFieldOne.Read(p)\n"+
			"Expected n == 5\n"+
			"Instead, n == %v\n",
			ePrefix.XCtxEmpty().String(),
			n)

		return
	}

	p = make([]byte, 100)

	txtFieldDateTimeTwo.ReaderInitialize()

	var readBytesCnt int
	var actualStr string
	n = 0

	for {

		n,
			err = txtFieldDateTimeTwo.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Error Returned From txtFieldDateTimeTwo.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if txtFieldDateTimeTwo.textLineReader != nil {
		t.Errorf("%v\n"+
			"Error: After completing Read Operation\n"+
			"txtFieldDateTimeTwo.textLineReader != 'nil'\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if readBytesCnt != lenExpectedDateTimeText {
		t.Errorf("%v\n"+
			"Byte Length Error: txtFieldDateTimeTwo.Read(p)\n"+
			"The actual length of bytes read\n"+
			"does NOT match the expected length.\n"+
			"Expected Bytes Read = '%v'\n"+
			"       Actual Bytes = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			lenExpectedDateTimeText,
			readBytesCnt)

		return
	}

	sMech := StrMech{}

	printableExpectedStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(expectedDateTimeText),
			true)

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if txtFieldDateTimeTwo.textLineReader != nil {
		t.Errorf("%v Test #1\n"+
			"Completed Read Operation but txtFieldLabelOne.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	p = make([]byte, 100)
	n = 0
	actualStr = ""
	readBytesCnt = 0

	for {

		n,
			err = txtFieldDateTimeTwo.Read(p)

		if n == 0 {
			break
		}

		actualStr += string(p[:n])
		readBytesCnt += n
	}

	if err != nil &&
		err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error Returned From txtFieldDateTimeTwo.Read(p)\n"+
			"Error = \n%v\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return
	}

	if err == nil {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error == nil!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	if err != io.EOF {
		t.Errorf("%v\n"+
			"Test # 2"+
			"Error: After completing Read Operation\n"+
			"the returned error should equal io.EOF.\n"+
			"HOWEVER, returned error is NOT equal io.EOF!\n",
			ePrefix.XCtxEmpty().String())

		return
	}

	printableActualStr =
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if printableExpectedStr != printableActualStr {
		t.Errorf("%v Test #2\n"+
			"Error: Expected Text String DOES NOT match\n"+
			"Actual Text String.\n"+
			"Expected Text String = '%v'\n"+
			"Instead, Text String = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			printableExpectedStr,
			printableActualStr)

		return
	}

	if txtFieldDateTimeTwo.textLineReader != nil {
		t.Errorf("%v Test #2\n"+
			"Completed Read Operation but txtFieldLabelOne.textLineReader\n"+
			"is NOT equal to 'nil'!\n",
			ePrefix.XCtxEmpty().String())

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
		txtFieldDateTimeTwo.SetDateTimeFieldSpec(
			dateTime,
			fieldLen,
			dateTimeFormat,
			textJustification,
			ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.SetDateTimeFieldSpec()\n"+
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

func TestTextFieldSpecDateTime_SetDateTimeFieldSpec_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_SetDateTimeFieldSpec_000200()",
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
		txtFieldDateTimeTwo.SetDateTimeFieldSpec(
			dateTime,
			fieldLen,
			dateTimeFormat,
			textJustification,
			ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.SetDateTimeFieldSpec()\n"+
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

func TestTextFieldSpecDateTime_SetDateTimeFieldSpec_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_SetDateTimeFieldSpec_000300()",
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

	txtFieldDateTimeTwo := TextFieldSpecDateTime{}

	err =
		txtFieldDateTimeTwo.SetDateTimeFieldSpec(
			dateTime,
			fieldLen,
			dateTimeFormat,
			textJustification,
			ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.SetDateTimeFieldSpec()\n"+
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

func TestTextFieldSpecDateTime_SetDateTimeFormat_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_SetDateTimeFormat_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	dateTimeFormat :=
		"2006-01-02 15:04:05.000000000 -0700 MST"

	err :=
		txtFieldDateTimeOne.SetDateTimeFormat(
			dateTimeFormat,
			ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetDateTimeFormat()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if dateTimeFormat !=
		txtFieldDateTimeOne.dateTimeFormat {
		t.Errorf("%v - ERROR\n"+
			"Expected Date Time Format = '%v'\n"+
			"Instead, Date Time Format = '%v'\n",
			ePrefix.String(),
			dateTimeFormat,
			txtFieldDateTimeOne.dateTimeFormat)

		return
	}

	timeZoneName := "America/Chicago"

	var tzLocPtr *time.Location

	tzLocPtr, err =
		time.LoadLocation(timeZoneName)

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

	dateTimeFormatTwo :=
		"Monday January 2, 2006 15:04:05.000000000 -0700 MST"

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormatTwo,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if dateTimeFormatTwo !=
		txtFieldDateTimeTwo.dateTimeFormat {
		t.Errorf("%v - ERROR\n"+
			"Expected Date Time Format Two = '%v'\n"+
			"Instead, Date Time Format Two = '%v'\n",
			ePrefix.String(),
			dateTimeFormatTwo,
			txtFieldDateTimeTwo.dateTimeFormat)

		return
	}

	err =
		txtFieldDateTimeTwo.SetDateTimeFormat(
			dateTimeFormat,
			ePrefix)

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeOne.SetDateTimeFormat()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	if dateTimeFormat !=
		txtFieldDateTimeTwo.dateTimeFormat {
		t.Errorf("%v - ERROR\n"+
			"Test #2 - Reset Format for 'txtFieldDateTimeTwo'\n"+
			"Expected Date Time Format = '%v'\n"+
			"Instead, Date Time Format = '%v'\n",
			ePrefix.String(),
			dateTimeFormat,
			txtFieldDateTimeTwo.dateTimeFormat)

		return
	}

	err =
		txtFieldDateTimeTwo.SetDateTimeFormat(
			"",
			ePrefix.XCtx(
				"dateTimeFormat==Empty String!"))

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldDateTimeTwo.SetDateTimeFormat()\n"+
			"because input parameter 'dateTimeFormat' is an empty string.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextFieldSpecDateTime_SetDateTimeValue_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_SetDateTimeValue_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	err :=
		txtFieldDateTimeOne.SetDateTimeValue(
			time.Time{},
			ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldDateTimeOne.SetDateTimeValue()\n"+
			"because input parameter 'dateTime' has a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timeZoneName := "America/Chicago"

	var tzLocPtr *time.Location

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

	fieldLen := len(dateTimeFormat) + 8

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

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

	dateTimeTwo := time.Date(
		2020,
		time.Month(1),
		1,
		14,
		30,
		0,
		0,
		tzLocPtr)

	err =
		txtFieldDateTimeTwo.SetDateTimeValue(
			dateTimeTwo,
			ePrefix.XCtx(
				"txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.SetDateTimeValue()\n"+
			"Error='%v'\n",
			ePrefix.XCtxEmpty().String(),
			err.Error())

		return

	}

	actualDateTime :=
		txtFieldDateTimeTwo.GetDateTime()

	if actualDateTime !=
		dateTimeTwo {
		t.Errorf("%v - ERROR\n"+
			"txtFieldDateTimeTwo.GetDateTime()\n"+
			"Expected Date Time = '%v'\n"+
			"Instead, Date Time = '%v'\n",
			ePrefix.XCtxEmpty().String(),
			actualDateTime.Format(dateTimeFormat),
			dateTimeTwo.Format(dateTimeFormat))

		return
	}

}

func TestTextFieldSpecDateTime_SetFieldLength_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextFieldSpecDateTime_SetFieldLength_000100()",
		"")

	txtFieldDateTimeOne := TextFieldSpecDateTime{}

	err :=
		txtFieldDateTimeOne.SetFieldLength(
			1000001,
			ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldDateTimeOne.SetFieldLength()\n"+
			"because input parameter 'fieldLen' has an invalid value (1000001).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtFieldDateTimeOne.SetFieldLength(
			-2,
			ePrefix)

	if err == nil {
		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtFieldDateTimeOne.SetFieldLength()\n"+
			"because input parameter 'fieldLen' has an invalid value (-2).\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timeZoneName := "America/Chicago"

	var tzLocPtr *time.Location

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

	textJustification := TxtJustify.Center()

	var txtFieldDateTimeTwo TextFieldSpecDateTime

	txtFieldDateTimeTwo,
		err = TextFieldSpecDateTime{}.NewDateTimeField(
		dateTime,
		fieldLen,
		dateTimeFormat,
		textJustification,
		ePrefix.XCtx("txtFieldDateTimeTwo"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by TextFieldSpecDateTime{}.NewDateTimeField()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	expectedFieldLen := 40

	err =
		txtFieldDateTimeTwo.SetFieldLength(
			expectedFieldLen,
			ePrefix.XCtx(
				"fieldLen==40"))

	if err != nil {
		t.Errorf("%v\n"+
			"Error returned by txtFieldDateTimeTwo.SetFieldLength()\n"+
			"Error:\n'%v'\n",
			ePrefix.String(),
			err.Error())

		return
	}

	actualFieldLen :=
		txtFieldDateTimeTwo.GetFieldLength()

	if expectedFieldLen != actualFieldLen {
		t.Errorf("%v - ERROR\n"+
			"txtFieldDateTimeTwo.GetFieldLength()\n"+
			"Expected Field Length = '%v'\n"+
			"Instead, Field Length = '%v'\n",
			ePrefix.String(),
			expectedFieldLen,
			actualFieldLen)

		return
	}

	return
}

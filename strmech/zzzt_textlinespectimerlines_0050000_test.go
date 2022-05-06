package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"testing"
	"time"
)

func TestTextLineSpecTimerLinesAtom_equal_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesAtom_equal_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines01(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	timerLines02,
		err = timerLines01.CopyOutPtr(
		ePrefix.XCpy(
			"timerLines02<-timerLines01 #1"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtTimerLinesAtom := textLineSpecTimerLinesAtom{}

	areEqual := txtTimerLinesAtom.equal(
		timerLines01,
		timerLines02)

	if !areEqual {

		t.Errorf("%v - ERROR\n"+
			"txtTimerLinesAtom.equal(timerLines02, timerLines01)\n"+
			"Expected timerLines01 == timerLines02\n"+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	areEqual = txtTimerLinesAtom.equal(
		nil,
		timerLines02)

	if areEqual {

		t.Errorf("%v - ERROR Test#2\n"+
			"txtTimerLinesAtom.equal(timerLines02, timerLines01)\n"+
			"Expected timerLines01 NOT EQUAL to timerLines02\n"+
			"because 'txtTimerLinesOne input parameter is 'nil'."+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	areEqual = txtTimerLinesAtom.equal(
		timerLines01,
		nil)

	if areEqual {

		t.Errorf("%v - ERROR Test#3\n"+
			"txtTimerLinesAtom.equal(timerLines02, timerLines01)\n"+
			"Expected timerLines01 NOT EQUAL to timerLines02\n"+
			"because 'txtTimerLinesTwo input parameter is 'nil'."+
			"HOWEVER, THEY ARE NOT EQUAL!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLinesAtom_testValidityOfTxtSpecTimerLines_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesAtom_testValidityOfTxtSpecTimerLines_000100()",
		"")

	txtTimerLinesAtom := textLineSpecTimerLinesAtom{}

	_,
		err := txtTimerLinesAtom.testValidityOfTxtSpecTimerLines(
		nil,
		ePrefix.XCpy(
			"txtTimerLines == nil"))

	if err == nil {

		t.Errorf("\n%v\n"+
			"Error: err = txtTimerLinesAtom.testValidityOfTxtSpecTimerLines()\n"+
			"Expected an error return because input\n"+
			"parameter 'txtTimerLines' is 'nil'.\n"+
			"HOWEVER NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLinesElectron_computeTimeDuration_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesElectron_computeTimeDuration_000100()",
		"")

	var loc *time.Location

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {

		t.Errorf(
			"\n%v - ERROR\n"+
				"time.LoadLocation(\"America/Chicago\")"+
				"%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	startTime := time.Date(
		2022,
		2,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := time.Date(
		2022,
		9,
		5,
		10,
		32,
		16,
		9000,
		loc)

	timerLinesElectron := textLineSpecTimerLinesElectron{}

	_,
		err = timerLinesElectron.computeTimeDuration(
		startTime,
		endTime,
		5,
		&ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecTimerLinesElectron_computeTimeDuration_000200(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesElectron_computeTimeDuration_000200()",
		"")

	var loc *time.Location

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {

		t.Errorf(
			"\n%v - ERROR\n"+
				"time.LoadLocation(\"America/Chicago\")"+
				"%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	startTime := time.Date(
		2022,
		2,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := time.Date(
		2022,
		9,
		5,
		10,
		32,
		16,
		9000,
		loc)

	timerLinesElectron := textLineSpecTimerLinesElectron{}

	_,
		err = timerLinesElectron.computeTimeDuration(
		time.Time{},
		endTime,
		5,
		&ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from\n"+
			"timerLinesElectron.computeTimeDuration()\n"+
			"because 'startTime' has a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	_,
		err = timerLinesElectron.computeTimeDuration(
		startTime,
		time.Time{},
		5,
		&ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from\n"+
			"timerLinesElectron.computeTimeDuration()\n"+
			"because 'endTime' has a zero value.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	_,
		err = timerLinesElectron.computeTimeDuration(
		endTime,
		startTime,
		5,
		&ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from\n"+
			"timerLinesElectron.computeTimeDuration()\n"+
			"because 'endTime' occurrs before 'startTime.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	_,
		err = timerLinesElectron.computeTimeDuration(
		startTime,
		endTime,
		-1,
		&ePrefix)

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from\n"+
			"timerLinesElectron.computeTimeDuration()\n"+
			"because 'summaryTextLineLeftMargin' has a value of -1.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLinesElectron_computeTimeDuration_000300(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesElectron_computeTimeDuration_000300()",
		"")

	var loc *time.Location

	loc,
		err := time.LoadLocation(
		"America/Chicago")

	if err != nil {

		t.Errorf(
			"\n%v - ERROR\n"+
				"time.LoadLocation(\"America/Chicago\")"+
				"%v\n",
			ePrefix.String(),
			err.Error())

		return
	}

	startTime := time.Date(
		2022,
		2,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := time.Date(
		2022,
		9,
		5,
		22,
		58,
		47,
		999999989,
		loc)

	timerLinesElectron := textLineSpecTimerLinesElectron{}

	_,
		err = timerLinesElectron.computeTimeDuration(
		startTime,
		endTime,
		55,
		&ePrefix)

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	return
}

func TestTextLineSpecTimerLinesElectron_empty_000100(t *testing.T) {

	timerLinesElectron := textLineSpecTimerLinesElectron{}

	timerLinesElectron.empty(
		nil)

}

func TestTextLineSpecTimerLinesElectron_getLengthOfLongestLabel_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTESTSERIES_getLengthOfLongestLabel_000100()",
		"")

	timerLinesElectron := textLineSpecTimerLinesElectron{}

	longestLabel := "A Very Grand End Time Label"

	expectedLabelLen := len(longestLabel)

	actualLongestLabelLen :=
		timerLinesElectron.getLengthOfLongestLabel(
			[]rune("startTime"),
			[]rune(longestLabel),
			[]rune("Time Duration"))

	if expectedLabelLen != actualLongestLabelLen {

		t.Errorf("\n%v\n"+
			"Error: timerLinesElectron.getLengthOfLongestLabel()\n"+
			"Expected Max Label Length != Actual Max Label Length\n"+
			"Expected Max Label Length = '%v'\n"+
			"  Actual Max Label Length = '%v'\n",
			ePrefix.String(),
			expectedLabelLen,
			actualLongestLabelLen)

		return
	}

	return
}

func TestTextLineSpecTimerLinesMolecule_getFormattedText_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesMolecule_getFormattedText_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtTimerLinesMolecule := textLineSpecTimerLinesMolecule{}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		timerLines01,
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		nil,
		ePrefix.XCpy(
			"timerLines01-Test#2"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesMolecule."+
			"getFormattedText()\n"+
			"because 'txtTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	_,
		err = txtTimerLinesMolecule.getFormattedText(
		&timerLines02,
		ePrefix.XCpy(
			"timerLines02-Test#1"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesMolecule."+
			"getFormattedText()\n"+
			"because 'txtTimerLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLinesNanobot_copyIn_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesNanobot_copyIn_000100()",
		"")

	_,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("\n%v\n",
			err.Error())
		return
	}

	var timerLines02 *TextLineSpecTimerLines

	_,
		timerLines02,
		err = createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	txtTimerLinesNanobot := textLineSpecTimerLinesNanobot{}

	err =
		txtTimerLinesNanobot.copyIn(
			timerLines02,
			nil,
			ePrefix.XCpy(
				"timerLines02<-nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'incomingTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	err =
		txtTimerLinesNanobot.copyIn(
			nil,
			timerLines01,
			ePrefix.XCpy(
				"nil<-timerLines01"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'targetTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	err =
		txtTimerLinesNanobot.copyIn(
			timerLines02,
			&timerLines03,
			ePrefix.XCpy(
				"timerLines02<-timerLines03 (invalid)"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyIn()\n"+
			"because 'incomingTimerLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

func TestTextLineSpecTimerLinesNanobot_copyOut_000100(t *testing.T) {

	ePrefix := ePref.ErrPrefixDto{}.NewEPrefCtx(
		"TestTextLineSpecTimerLinesNanobot_copyOut_000100()",
		"")

	outputStr,
		timerLines01,
		err := createTestTextLineSpecTimerLines02(
		ePrefix.XCpy(
			"timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	timerLines02 := TextLineSpecTimerLines{}

	txtTimerLinesNanobot := textLineSpecTimerLinesNanobot{}

	timerLines02,
		err = txtTimerLinesNanobot.copyOut(
		timerLines01,
		ePrefix.XCpy(
			"timerLines02<-timerLines01"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	var actualStr string

	actualStr,
		err = timerLines02.GetFormattedText(
		ePrefix.XCpy(
			"timerLines02"))

	if err != nil {
		t.Errorf("%v\n",
			err.Error())
		return
	}

	sMech := StrMech{}

	printableActualStr :=
		sMech.ConvertNonPrintableChars(
			[]rune(actualStr),
			true)

	if outputStr != printableActualStr {

		t.Errorf("%v - ERROR\n"+
			"timerLines02.GetFormattedText()"+
			"Expected string DOES NOT match Actual string\n"+
			"Expected string = '%v'\n"+
			"  Actual string = '%v'\n",
			ePrefix.String(),
			outputStr,
			printableActualStr)

		return
	}

	_,
		err = txtTimerLinesNanobot.copyOut(
		nil,
		ePrefix.XCpy(
			"txtTimerLines is nil"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyOut()\n"+
			"because 'txtTimerLines' is 'nil'.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	timerLines03 := TextLineSpecTimerLines{}

	_,
		err = txtTimerLinesNanobot.copyOut(
		&timerLines03,
		ePrefix.XCpy(
			"timerLines03 is invalid"))

	if err == nil {

		t.Errorf("%v - ERROR\n"+
			"Expected an error return from txtTimerLinesNanobot."+
			"copyOut()\n"+
			"because 'txtTimerLines' is invalid.\n"+
			"HOWEVER, NO ERROR WAS RETURNED!\n",
			ePrefix.String())

		return
	}

	return
}

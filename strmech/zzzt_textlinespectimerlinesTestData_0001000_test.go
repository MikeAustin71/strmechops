package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"time"
)

func createTestTextLineSpecTimerLines01(
	errorPrefix interface{}) (
	*TextLineSpecTimerLines,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecStandardLine01()",
		"")

	if err != nil {
		return &TextLineSpecTimerLines{}, err
	}

	var loc *time.Location

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		return &TextLineSpecTimerLines{}, err
	}

	startTime := time.Date(
		2022,
		4,
		5,
		10,
		0,
		0,
		0,
		loc)

	endTime := startTime.Add((time.Microsecond * 5) + 999)

	var timerLines01 *TextLineSpecTimerLines

	timerLines01,
		err = TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		return &TextLineSpecTimerLines{}, err
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCtx(
			"timerLines01"))

	if err != nil {
		return &TextLineSpecTimerLines{}, err
	}

	return timerLines01, err
}

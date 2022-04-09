package strmech

import (
	ePref "github.com/MikeAustin71/errpref"
	"time"
)

func createTestTextLineSpecTimerLines01(
	errorPrefix interface{}) (
	string,
	*TextLineSpecTimerLines,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var outputStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecStandardLine01()",
		"")

	if err != nil {
		return outputStr, &TextLineSpecTimerLines{}, err
	}

	var loc *time.Location

	loc,
		err = time.LoadLocation(
		"America/Chicago")

	if err != nil {
		return outputStr, &TextLineSpecTimerLines{}, err
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
		return outputStr, &TextLineSpecTimerLines{}, err
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCtx(
			"timerLines01"))

	if err != nil {
		return outputStr, &TextLineSpecTimerLines{}, err
	}

	/*
		"[SPACE][SPACE]Start[SPACE]Time:[SPACE]2022-04-05[SPACE]10:00:00.000000000[SPACE]-0500[SPACE]CDT\n[SPACE][SPACE][SPACE][SPACE]End[SPACE]Time:[SPACE]2022-04-05[SPACE]10:00:00.000005999[SPACE]-0500[SPACE]CDT\nElapsed[SPACE]Time:[SPACE]5[SPACE]Microseconds[SPACE]999[SPACE]Nanoseconds\n[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Total[SPACE]Elapsed[SPACE]Nanoseconds:[SPACE]5,999\n"
	*/

	outputStr =
		"[SPACE][SPACE]Start[SPACE]Time:[SPACE]2022-04-05[SPACE]" +
			"10:00:00.000000000[SPACE]-0500[SPACE]CDT\\n[SPACE][SPACE]" +
			"[SPACE][SPACE]End[SPACE]Time:[SPACE]2022-04-05[SPACE]" +
			"10:00:00.000005999[SPACE]-0500[SPACE]CDT\\nElapsed[SPACE]" +
			"Time:[SPACE]5[SPACE]Microseconds[SPACE]999[SPACE]" +
			"Nanoseconds\\n[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]" +
			"[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]" +
			"[SPACE]Total[SPACE]Elapsed[SPACE]" +
			"Nanoseconds:[SPACE]5,999\\n"

	return outputStr, timerLines01, err
}

func createTestTextLineSpecTimerLines02(
	errorPrefix interface{}) (
	string,
	*TextLineSpecTimerLines,
	error) {

	var ePrefix *ePref.ErrPrefixDto
	var err error
	var outputStr string

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewIEmpty(
		errorPrefix,
		"TestDataGeneration - "+
			"createTestTextLineSpecStandardLine02()",
		"")

	if err != nil {
		return outputStr, &TextLineSpecTimerLines{}, err
	}

	var loc *time.Location

	loc,
		err = time.LoadLocation(
		"America/Los_Angeles")

	if err != nil {
		return outputStr, &TextLineSpecTimerLines{}, err
	}

	startTime := time.Date(
		2021,
		4,
		20,
		18,
		30,
		0,
		0,
		loc)

	elapsedTime := (time.Millisecond * 200) +
		(time.Microsecond * 5) + 355

	endTime := startTime.Add(elapsedTime)

	var timerLines01 *TextLineSpecTimerLines

	timerLines01,
		err = TextLineSpecTimerLines{}.NewDefaultFullTimerEvent(
		startTime,
		endTime,
		ePrefix)

	if err != nil {
		return outputStr, &TextLineSpecTimerLines{}, err
	}

	err = timerLines01.IsValidInstanceError(
		ePrefix.XCtx(
			"timerLines01"))

	if err != nil {
		return outputStr, &TextLineSpecTimerLines{}, err
	}

	/*
		"[SPACE][SPACE]Start[SPACE]Time:[SPACE]2021-04-20[SPACE]18:30:00.000000000[SPACE]-0700[SPACE]PDT\n[SPACE][SPACE][SPACE][SPACE]End[SPACE]Time:[SPACE]2021-04-20[SPACE]18:30:00.200005355[SPACE]-0700[SPACE]PDT\nElapsed[SPACE]Time:[SPACE]200[SPACE]Milliseconds[SPACE]5[SPACE]Microseconds[SPACE]355[SPACE]Nanoseconds\n[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]Total[SPACE]Elapsed[SPACE]Nanoseconds:[SPACE]200,005,355\n"
	*/

	outputStr = "[SPACE][SPACE]Start[SPACE]Time:[SPACE]2021-04-20" +
		"[SPACE]18:30:00.000000000[SPACE]-0700[SPACE]PDT\\n[SPACE]" +
		"[SPACE][SPACE][SPACE]End[SPACE]Time:[SPACE]2021-04-20" +
		"[SPACE]18:30:00.200005355[SPACE]-0700[SPACE]PDT\\nElapsed" +
		"[SPACE]Time:[SPACE]200[SPACE]Milliseconds[SPACE]5[SPACE]" +
		"Microseconds[SPACE]355[SPACE]Nanoseconds\\n[SPACE][SPACE]" +
		"[SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE][SPACE]" +
		"[SPACE][SPACE][SPACE][SPACE]Total[SPACE]Elapsed[SPACE]" +
		"Nanoseconds:[SPACE]200,005,355\\n"

	return outputStr, timerLines01, err
}

package strmech

import (
	"fmt"
	ePref "github.com/MikeAustin71/errpref"
	"math"
	"math/big"
	"strings"
	"sync"
	"time"
)

type textLineSpecAverageTimeAtom struct {
	lock *sync.Mutex
}

//	calcAverageDuration
//
//	Calculates the average duration of a timer event
//	series encapsulated in an instance of
//	TextLineSpecAverageTime passed as input parameter
//	'txtLineAvgTimer'.
//
//	Average time duration is calculated by dividing the
//	total time duration by the number of separate timing
//	events.
//
//	In addition to Average Duration, this method also
//	returns the minimum and maximum time duration for
//	all timing events included in the average time
//	calculation.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLineAvgTimer				*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecAverageTime.
//		The internal average time counters maintained by
//		this instance will be used to produce an average
//		time duration for all recorded timing events.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	avgDuration					int64
//
//		If this method completes successfully, this
//		parameter will return the average time duration
//		for all recorded timing events. Average duration
//		is calculated by dividing the total duration by
//		the number of separate timing events.
//
//	maximumTimeDuration			int64
//
//		If this method completes successfully, this
//		parameter will return the maximum time duration
//		from all recorded timing events.
//
//	minimumTimeDuration			int64
//
//		If this method completes successfully, this
//		parameter will return the minimum time duration
//		from all recorded timing events.
//
//	err							error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLineAvgTimeAtom *textLineSpecAverageTimeAtom) calcAverageDuration(
	txtLineAvgTimer *TextLineSpecAverageTime,
	errPrefDto *ePref.ErrPrefixDto) (
	avgDuration int64,
	maximumTimeDuration int64,
	minimumTimeDuration int64,
	err error) {

	if txtLineAvgTimeAtom.lock == nil {
		txtLineAvgTimeAtom.lock = new(sync.Mutex)
	}

	txtLineAvgTimeAtom.lock.Lock()

	defer txtLineAvgTimeAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeAtom."+
			"calcAverageDuration()",
		"")

	if err != nil {

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	if txtLineAvgTimer == nil {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is a nil pointer!\n",
			ePrefix.String())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	var err2 error

	_,
		err2 = new(textLineSpecAverageTimeElectron).
		testValidityOfTxtLineAvgTimer(
			txtLineAvgTimer,
			ePrefix.XCpy("txtLineAvgTimer"))

	if err2 != nil {

		err = fmt.Errorf("calcAverageDuration()\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"\n%v\n",
			err2.Error())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	bigZero := big.NewInt(0)

	cmpValue := txtLineAvgTimer.numberOfDurationEvents.Cmp(
		bigZero)

	if cmpValue == 0 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"No timer events have been recorded.\n"+
			"'txtLineAvgTimer.numberOfDurationEvents' is equal to zero!\n",
			ePrefix.String())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	cmpValue = txtLineAvgTimer.totalDurationNanoSecs.Cmp(
		bigZero)

	if cmpValue == -1 {
		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'txtLineAvgTimer' is invalid!\n"+
			"'txtLineAvgTimer.totalDurationNanoSecs' is less than zero!\n",
			ePrefix.String())

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	if cmpValue == 0 {

		avgDuration = 0

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	cmpValue = txtLineAvgTimer.numberOfDurationEvents.Cmp(
		&txtLineAvgTimer.totalDurationNanoSecs)

	if cmpValue == 0 {

		avgDuration = 1

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	if cmpValue == 1 {

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	// Number of Duration Events is
	// less than total duration

	bigZero.Quo(
		&txtLineAvgTimer.totalDurationNanoSecs,
		&txtLineAvgTimer.numberOfDurationEvents)

	maxInt64 := big.NewInt(0).SetInt64(
		math.MaxInt64)

	cmpValue = bigZero.Cmp(maxInt64)

	if cmpValue == 1 {
		err = fmt.Errorf("%v\n"+
			"Calculation Error: Average Duration exceeds\n"+
			"Maximum int64 value.\n"+
			"Average Duration = %v\n",
			ePrefix.String(),
			bigZero.Text(10))

		return avgDuration,
			maximumTimeDuration,
			minimumTimeDuration,
			err
	}

	avgDuration = bigZero.Int64()

	minimumTimeDuration =
		txtLineAvgTimer.minimumTimeDuration.Int64()

	maximumTimeDuration =
		txtLineAvgTimer.maximumTimeDuration.Int64()

	return avgDuration,
		maximumTimeDuration,
		minimumTimeDuration,
		err
}

//	getDurationElementReport
//
//	Receives time duration allocation data broken down by
//	days, hours, minutes, seconds, milliseconds,
//	microseconds and nanoseconds.
//
//	This method is used to produce reports for average
//	time duration, maximum time duration and minimum time
//	duration.
//
//	The text lines contained in the time duration report
//	created by this method will be stored in input
//	parameter, 'strBuilder'.
//
// ----------------------------------------------------------------
//
// # Input Parameters
//
//	txtLineAvgTimer				*TextLineSpecAverageTime
//
//		A pointer to an instance of TextLineSpecAverageTime.
//		The internal average time counters maintained by
//		this instance will be used to produce an average
//		time duration for all recorded timing events.
//
//	strBuilder					*strings.Builder
//
//		A pointer to an instance of *strings.Builder. The
//		formatted text characters produced by this method
//		will be written to this instance of
//		strings.Builder.
//
//	allocatedDuration			TimeDurationDto
//
//		Allocated time duration values. Type
//		TimeDurationDto contains time duration numerical
//		values broken down by Days, Hours, Minutes,
//		Seconds, Milliseconds, Microseconds and
//		Nanoseconds. These items are referred to as time
//		duration elements. The duration element values
//		are converted to text strings for the creation
//		of time duration reports.
//
//		type TimeDurationDto struct {
//
//			TotalNanoseconds int64
//				The total number of Nanoseconds to
//				be allocated.
//
//			NumberOfDays int64
//				The number Days represented by
//				'TotalNanoseconds'.
//
//			NumberOfHours int64
//				The number Hours represented by
//				'TotalNanoseconds'.
//
//			NumberOfMinutes int64
//				The number Minutes represented by
//				'TotalNanoseconds'.
//
//			NumberOfSeconds int64
//				The number Seconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfMilliseconds int64
//				The number Milliseconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfMicroseconds int64
//				The number Microseconds represented by
//				'TotalNanoseconds'.
//
//			NumberOfNanoseconds int64
//				The number Nanoseconds remaining after
//				the allocation of Microseconds.
//		}
//
//	isAllocatedAvgDuration		bool
//
//		This boolean value selects the report format to
//		be applied.
//
//		When set to 'true', this parameter specifies the
//		Average Duration Report Format which includes
//		the number of timing cycles used to compute the
//		average time duration.
//
//		When set to 'false', this parameter specifies the
//		Standard Time Duration report. This report format
//		details an absolute time duration value as
//		opposed to an 'average' of time durations. This
//		format is used to report maximum and minimum time
//		duration values.
//
//	timeDurationTitle			string
//
//		The title or label applied to this time duration
//		report.
//
//	maxLineLength				int
//
//		This integer value specifies the maximum length of
//		the line in which the report text will be formatted.
//
//	errPrefDto					*ePref.ErrPrefixDto
//
//		This object encapsulates an error prefix string
//		which is included in all returned error
//		messages. Usually, it contains the name of the
//		calling method or methods listed as a function
//		chain.
//
//		If no error prefix information is needed, set
//		this parameter to 'nil'.
//
//		Type ErrPrefixDto is included in the 'errpref'
//		software package:
//			"github.com/MikeAustin71/errpref".
//
// ----------------------------------------------------------------
//
// # Return Values
//
//	error
//
//		If this method completes successfully, the
//		returned error Type is set equal to 'nil'. If
//		errors are encountered during processing, the
//		returned error Type will encapsulate an error
//		message.
//
//		If an error message is returned, the text value
//		for input parameter 'errPrefDto' (error prefix)
//		will be prefixed or attached at the beginning of
//		the error message.
func (txtLineAvgTimeAtom *textLineSpecAverageTimeAtom) getDurationElementReport(
	txtLineAvgTimer *TextLineSpecAverageTime,
	strBuilder *strings.Builder,
	allocatedDuration TimeDurationDto,
	isAllocatedAvgDuration bool,
	timeDurationTitle string,
	maxLineLength int,
	errPrefDto *ePref.ErrPrefixDto) error {

	if txtLineAvgTimeAtom.lock == nil {
		txtLineAvgTimeAtom.lock = new(sync.Mutex)
	}

	txtLineAvgTimeAtom.lock.Lock()

	defer txtLineAvgTimeAtom.lock.Unlock()

	var ePrefix *ePref.ErrPrefixDto

	var err error

	ePrefix,
		err = ePref.ErrPrefixDto{}.NewFromErrPrefDto(
		errPrefDto,
		"textLineSpecAverageTimeAtom."+
			"getDurationElementReport()",
		"")

	if err != nil {

		return err
	}
	solidLineLeftMargin := " "
	titleLineLeftMargin := "  "

	originalMaxLineLength := maxLineLength
	lenTitleLineLeftMargin := len(titleLineLeftMargin)
	maxLineLength -= lenTitleLineLeftMargin

	if maxLineLength < 1 {

		err = fmt.Errorf("%v\n"+
			"Error: Input parameter 'maxLineLength' is invalid!\n"+
			"'maxLineLength' has a length less than one ('1').\n"+
			"Original maxLineLength = %v\n"+
			"Adjusted Maximum Line Length = maxLineLength - %v \n"+
			"Adjusted Maximum Line Length = %v\n",
			ePrefix.String(),
			originalMaxLineLength,
			lenTitleLineLeftMargin,
			maxLineLength)

		return err
	}

	txtLineCollection := new(TextLineSpecLinesCollection).New()

	err = txtLineCollection.AddBlankLine(
		2,
		ePrefix.XCpy("Top Line #1"))

	if err != nil {
		return err
	}

	err = txtLineCollection.AddSolidLine(
		solidLineLeftMargin,
		"=",
		maxLineLength,
		"",
		"",
		false,
		1,
		ePrefix.XCpy("Top Solid Line #1"))

	if err != nil {
		return err
	}

	err = txtLineCollection.AddPlainTextLine(
		titleLineLeftMargin,
		"",
		timeDurationTitle,
		maxLineLength,
		TxtJustify.Center(),
		ePrefix.XCpy(
			"Text Title Top Line #1"))

	if err != nil {
		return err
	}

	err = txtLineCollection.AddDateTimeLine(
		titleLineLeftMargin,
		"",
		time.Now(),
		new(DateTimeHelper).GetDateTimeFormat(4),
		maxLineLength,
		TxtJustify.Center(),
		ePrefix)

	if err != nil {
		return err
	}

	err = txtLineCollection.AddSolidLine(
		solidLineLeftMargin,
		"-",
		maxLineLength,
		"",
		"",
		false,
		1,
		ePrefix.XCpy("Leading After Date Solid Line"))

	if err != nil {
		return err
	}

	err = txtLineCollection.AddBlankLine(
		1,
		ePrefix.XCpy("Before Data Blank Line"))

	if err != nil {
		return err
	}

	var allocatedDurationStrs StringArrayDto

	allocatedDurationStrs,
		err = new(DateTimeHelper).GetFmtAllocatedDurationText(
		allocatedDuration,
		maxLineLength,
		ePrefix.XCpy("allocatedDurationStrs<-"))

	if err != nil {
		return err
	}

	var summaryTotalStr string
	/*
		fmt.Printf("Initial String Array Length: '%v'\n\n",
			allocatedDurationStrs.GetStringArrayLength())
	*/
	summaryTotalStr,
		_,
		err = allocatedDurationStrs.PopLastStr(
		ePrefix.XCpy("summaryTotalStr<-"))

	if err != nil {
		return err
	}

	/*
		fmt.Printf("\nThis is summaryTotalStr: %v\n\n",
			summaryTotalStr)

		fmt.Printf("Final String Array Length: '%v'\n\n",
			allocatedDurationStrs.GetStringArrayLength())

	*/

	err = txtLineCollection.AddPlainTextStrArray(
		titleLineLeftMargin+"  ",
		"",
		allocatedDurationStrs,
		maxLineLength-2,
		TxtJustify.Center(),
		ePrefix)

	if err != nil {
		return err
	}

	err = txtLineCollection.AddSolidLine(
		solidLineLeftMargin,
		"-",
		maxLineLength,
		"",
		"",
		false,
		1,
		ePrefix.XCpy("Summary Total Solid Line"))

	if err != nil {
		return err
	}

	err = txtLineCollection.AddPlainTextLine(
		titleLineLeftMargin,
		"",
		summaryTotalStr,
		maxLineLength,
		TxtJustify.Left(),
		ePrefix.XCpy("summaryTotalStr"))

	if err != nil {
		return err
	}

	if isAllocatedAvgDuration == true {

		var nStrIntSeparator IntegerSeparatorSpec

		nStrIntSeparator,
			err = new(IntegerSeparatorSpec).NewUnitedStatesDefaults(
			ePrefix.XCpy(
				"nStrIntSeparator<-"))

		if err != nil {
			return err
		}

		var tLine string

		tLine,
			err = nStrIntSeparator.GetFmtIntSeparatedNumStr(
			txtLineAvgTimer.numberOfDurationEvents.Text(10),
			ePrefix.XCpy("tLine<-"))

		if err != nil {
			return err
		}

		err = txtLineCollection.AddPlainTextLine(
			titleLineLeftMargin+"  ",
			"",
			fmt.Sprintf("Total Timing Cycles: %v",
				tLine),
			maxLineLength,
			TxtJustify.Left(),
			ePrefix.XCpy("summaryTotalStr"))

		if err != nil {
			return err
		}

	}

	err = txtLineCollection.AddSolidLine(
		solidLineLeftMargin,
		"=",
		maxLineLength,
		"",
		"",
		false,
		1,
		ePrefix.XCpy("Summary Total Solid Line"))

	if err != nil {
		return err
	}

	err = txtLineCollection.AddBlankLine(
		2,
		ePrefix.XCpy("Bottom-Ending Blank Lines"))

	if err != nil {
		return err
	}

	_,
		err = txtLineCollection.TextBuilder(
		strBuilder,
		ePrefix.XCpy("strBuilder<-"))

	return err
}
